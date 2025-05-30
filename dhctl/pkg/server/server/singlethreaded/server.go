// Copyright 2024 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package singlethreaded

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/deckhouse/deckhouse/dhctl/pkg/config"
	dhctllog "github.com/deckhouse/deckhouse/dhctl/pkg/log"
	pbdhctl "github.com/deckhouse/deckhouse/dhctl/pkg/server/pb/dhctl"
	"github.com/deckhouse/deckhouse/dhctl/pkg/server/pkg/interceptors"
	"github.com/deckhouse/deckhouse/dhctl/pkg/server/pkg/logger"
	"github.com/deckhouse/deckhouse/dhctl/pkg/server/rpc/dhctl"
	"github.com/deckhouse/deckhouse/dhctl/pkg/util/tomb"
)

// Serve starts GRPC server
func Serve(network, address string) error {
	dhctllog.InitLoggerWithOptions("silent", dhctllog.LoggerOptions{})
	lvl := &slog.LevelVar{}
	lvl.Set(slog.LevelDebug)
	log := logger.NewLogger(lvl).With(slog.String("component", "singlethreaded_server"))

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	defer close(done)

	// set concurrency limit of 1 for all rpcs
	sem := make(chan struct{}, 1)
	limiterPrefix := ""

	podName := os.Getenv("HOSTNAME")

	tomb.RegisterOnShutdown("server", func() {
		log.Info("stopping grpc server")
		cancel()
		<-done
		log.Info("grpc server stopped")
	})

	cacheDir, err := cacheDirectory()
	if err != nil {
		return fmt.Errorf("failed to init grpc server: %w", err)
	}

	log.Info(
		"starting grpc server",
		slog.String("network", network),
		slog.String("address", address),
		slog.String("cache directory", cacheDir),
	)

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Error("failed to listen", logger.Err(err))
		return err
	}
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.UnaryLogger(log),
			logging.UnaryServerInterceptor(interceptors.Logger()),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(interceptors.PanicRecoveryHandler())),
			interceptors.UnaryParallelTasksLimiter(sem, limiterPrefix),
		),
		grpc.ChainStreamInterceptor(
			interceptors.StreamLogger(log),
			logging.StreamServerInterceptor(interceptors.Logger()),
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandlerContext(interceptors.PanicRecoveryHandler())),
			interceptors.StreamParallelTasksLimiter(sem, limiterPrefix),
		),
	)

	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-a-grpc-liveness-probe
	healthService := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthService)

	// grpcurl -plaintext host:port describe
	reflection.Register(s)

	// init services
	dhctlService := dhctl.New(podName, cacheDir, config.NewSchemaStore())

	// register services
	pbdhctl.RegisterDHCTLServer(s, dhctlService)

	go func() {
		<-ctx.Done()

		s.GracefulStop()
	}()

	if err = s.Serve(listener); err != nil {
		log.Error("failed to serve", logger.Err(err))
		return err
	}
	return nil
}

func cacheDirectory() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("creating uuid for cache directory")
	}

	path := filepath.Join(os.TempDir(), "dhctl", "cache_"+id.String())

	return path, nil
}
