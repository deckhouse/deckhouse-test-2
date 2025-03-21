// Copyright 2021 Flant JSC
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

package destroy

import (
	"context"

	"github.com/google/uuid"

	"github.com/deckhouse/deckhouse/dhctl/pkg/operations/converge/lock"

	"github.com/deckhouse/deckhouse/dhctl/pkg/kubernetes"
	"github.com/deckhouse/deckhouse/dhctl/pkg/kubernetes/actions/deckhouse"
	"github.com/deckhouse/deckhouse/dhctl/pkg/kubernetes/client"
	"github.com/deckhouse/deckhouse/dhctl/pkg/log"
	"github.com/deckhouse/deckhouse/dhctl/pkg/system/node/ssh"
)

type DeckhouseDestroyerOptions struct {
	CommanderMode bool
	CommanderUUID uuid.UUID
}

type DeckhouseDestroyer struct {
	convergeUnlocker func(fullUnlock bool)
	sshClient        *ssh.Client
	kubeCl           *client.KubernetesClient
	state            *State

	DeckhouseDestroyerOptions
}

func NewDeckhouseDestroyer(sshClient *ssh.Client, state *State, opts DeckhouseDestroyerOptions) *DeckhouseDestroyer {
	return &DeckhouseDestroyer{
		sshClient:                 sshClient,
		state:                     state,
		DeckhouseDestroyerOptions: opts,
	}
}

func (g *DeckhouseDestroyer) UnlockConverge(fullUnlock bool) {
	if g.convergeUnlocker != nil {
		g.convergeUnlocker(fullUnlock)
		g.convergeUnlocker = nil
	}
}

func (g *DeckhouseDestroyer) StopProxy() {
	if g.kubeCl == nil {
		return
	}

	g.kubeCl.KubeProxy.Stop(0)
	g.kubeCl = nil
}

func (g *DeckhouseDestroyer) GetKubeClient() (*client.KubernetesClient, error) {
	if g.kubeCl != nil {
		return g.kubeCl, nil
	}

	// TODO(dhctl-for-commander-cancels): pass ctx
	kubeCl, err := kubernetes.ConnectToKubernetesAPI(context.TODO(), ssh.NewNodeInterfaceWrapper(g.sshClient))
	if err != nil {
		return nil, err
	}
	g.kubeCl = kubeCl

	if !g.CommanderMode {
		unlockConverge, err := lock.LockConverge(kubernetes.NewSimpleKubeClientGetter(kubeCl), "local-destroyer")
		if err != nil {
			return nil, err
		}
		g.convergeUnlocker = unlockConverge
	}

	return kubeCl, err
}

func (g *DeckhouseDestroyer) DeleteResources(cloudType string) error {
	resourcesDestroyed, err := g.state.IsResourcesDestroyed()
	if err != nil {
		return err
	}

	if resourcesDestroyed {
		log.WarnLn("Resources was destroyed. Skip it")
		return nil
	}

	kubeCl, err := g.GetKubeClient()
	if err != nil {
		return err
	}

	return log.Process("common", "Delete resources from the Kubernetes cluster", func() error {
		return g.deleteEntities(kubeCl)
	})
}

func (g *DeckhouseDestroyer) deleteEntities(kubeCl *client.KubernetesClient) error {
	err := deckhouse.DeleteDeckhouseDeployment(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.WaitForDeckhouseDeploymentDeletion(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.DeletePDBs(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.DeleteServices(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.WaitForServicesDeletion(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.DeleteAllD8StorageResources(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.DeleteStorageClasses(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.DeletePVC(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.DeletePods(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.WaitForPVCDeletion(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.DeletePV(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.WaitForPVDeletion(kubeCl)
	if err != nil {
		return err
	}

	err = deckhouse.DeleteMachinesIfResourcesExist(kubeCl)
	if err != nil {
		return err
	}

	return nil
}
