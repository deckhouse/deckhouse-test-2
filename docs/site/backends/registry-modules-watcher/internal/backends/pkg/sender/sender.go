// Copyright 2023 Flant JSC
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

package sender

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	url2 "net/url"
	"registry-modules-watcher/internal/backends"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/deckhouse/deckhouse/pkg/log"
)

const maxElapsedTime = 15 // minutes
const maxRetries = 10

type sender struct {
	client *http.Client

	logger *log.Logger
}

// New
func New(logger *log.Logger) *sender {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr}

	return &sender{
		client: client,
		logger: logger,
	}
}

func (s *sender) Send(ctx context.Context, listBackends map[string]struct{}, versions []backends.Version) error {
	syncChan := make(chan struct{}, 10)
	for backend := range listBackends {
		syncChan <- struct{}{}

		go func(backend string) {
			for _, version := range versions {
				if version.ToDelete {
					err := s.delete(ctx, backend, version.Module, version.ReleaseChannels)
					if err != nil {
						s.logger.Error("send delete", log.Err(err))
					}

					continue
				}

				url := "http://" + backend + "/api/v1/doc/" + version.Module + "/" + version.Version + "?channels=" + strings.Join(version.ReleaseChannels, ",")

				err := s.loadDocArchive(ctx, url, version.TarFile)
				if err != nil {
					s.logger.Error("send docs", log.Err(err))
				}
			}

			url := "http://" + backend + "/api/v1/build"

			err := s.build(ctx, url)
			if err != nil {
				s.logger.Error("build docs", log.Err(err))
			}
			<-syncChan
		}(backend)
	}

	return nil
}

func (s *sender) delete(_ context.Context, backend, moduleName string, releaseChannels []string) error {
	url := fmt.Sprintf("http://%s/api/v1/doc/%s", backend, moduleName)
	if len(releaseChannels) > 0 {
		params := url2.Values{}
		params.Add("channels", strings.Join(releaseChannels, ","))
		url += "?" + params.Encode()
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("client: could not create request: %s", err)
	}

	operation := func() error {
		resp, err := s.client.Do(req)
		if err != nil {
			return fmt.Errorf("client: error making http request: %s", err)
		}

		// TODO: add >=300 status code handling
		s.logger.Info("delete response", slog.Int("status_code", resp.StatusCode))

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = maxElapsedTime * time.Minute
	err = backoff.Retry(operation, backoff.WithMaxRetries(b, maxRetries))
	if err != nil {
		return err
	}

	return nil
}

func (s *sender) loadDocArchive(_ context.Context, url string, tarFile []byte) error {
	s.logger.Info("send loadDoc", slog.String("url", url))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(tarFile))
	if err != nil {
		return fmt.Errorf("client: could not create request: %s", err)
	}

	req.Header.Set("Content-Type", "application/tar")

	operation := func() error {
		resp, err := s.client.Do(req)
		if err != nil {
			return fmt.Errorf("client: error making http request: %s", err)
		}

		// TODO: add >=300 status code handling
		s.logger.Info("send archive response", slog.Int("status_code", resp.StatusCode))

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = maxElapsedTime * time.Minute
	err = backoff.Retry(operation, backoff.WithMaxRetries(b, maxRetries))
	if err != nil {
		return err
	}

	return nil
}

func (s *sender) build(_ context.Context, url string) error {
	s.logger.Info("send build", slog.String("url", url))

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("client: could not create request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	operation := func() error {
		resp, err := s.client.Do(req)
		if err != nil {
			return fmt.Errorf("client: error making http request: %s", err)
		}

		// TODO: add >=300 status code handling
		s.logger.Info("send build response", slog.Int("status_code", resp.StatusCode))

		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = maxElapsedTime * time.Minute
	err = backoff.Retry(operation, backoff.WithMaxRetries(b, maxRetries))
	if err != nil {
		return err
	}

	return nil
}
