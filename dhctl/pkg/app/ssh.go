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

package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

const DefaultSSHAgentPrivateKeys = "~/.ssh/id_rsa"

var (
	SSHAgentPrivateKeys = make([]string, 0)
	SSHPrivateKeys      = make([]string, 0)
	SSHBastionHost      = ""
	SSHBastionPort      = ""
	SSHBastionUser      = os.Getenv("USER")
	SSHUser             = os.Getenv("USER")
	SSHHosts            = make([]string, 0)
	SSHPort             = ""
	SSHExtraArgs        = ""

	AskBecomePass = false
	BecomePass    = ""
)

func DefineSSHFlags(cmd *kingpin.CmdClause) {
	cmd.Flag("ssh-agent-private-keys", "Paths to private keys. Those keys will be used to connect to servers and to the bastion. Can be specified multiple times (default: '~/.ssh/id_rsa')").
		Envar(configEnvName("SSH_AGENT_PRIVATE_KEYS")).
		StringsVar(&SSHAgentPrivateKeys)
	cmd.Flag("ssh-bastion-host", "Jumper (bastion) host to connect to servers (will be used both by terraform and ansible). Only IPs or hostnames are supported, name from ssh-config will not work.").
		Envar(configEnvName("SSH_BASTION_HOST")).
		StringVar(&SSHBastionHost)
	cmd.Flag("ssh-bastion-port", "SSH destination port").
		Envar(configEnvName("SSH_BASTION_PORT")).
		StringVar(&SSHBastionPort)
	cmd.Flag("ssh-bastion-user", "User to authenticate under when connecting to bastion (default: $USER)").
		Default(SSHBastionUser).
		Envar(configEnvName("SSH_BASTION_USER")).
		StringVar(&SSHBastionUser)
	cmd.Flag("ssh-user", "User to authenticate under (default: $USER)").
		Envar(configEnvName("SSH_USER")).
		Default(SSHUser).
		StringVar(&SSHUser)
	cmd.Flag("ssh-host", "SSH destination hosts, can be specified multiple times").
		Envar(configEnvName("SSH_HOSTS")).
		StringsVar(&SSHHosts)
	cmd.Flag("ssh-port", "SSH destination port").
		Envar(configEnvName("SSH_PORT")).
		StringVar(&SSHPort)
	cmd.Flag("ssh-extra-args", "extra args for ssh commands (-vvv)").
		Envar(configEnvName("SSH_EXTRA_ARGS")).
		StringVar(&SSHExtraArgs)

	cmd.PreAction(func(c *kingpin.ParseContext) (err error) {
		if len(SSHAgentPrivateKeys) == 0 {
			SSHAgentPrivateKeys = append(SSHAgentPrivateKeys, DefaultSSHAgentPrivateKeys)
		}
		SSHPrivateKeys, err = ParseSSHPrivateKeyPaths(SSHAgentPrivateKeys)
		if err != nil {
			return fmt.Errorf("ssh private keys: %v", err)
		}
		return nil
	})
}

func ParseSSHPrivateKeyPaths(pathSets []string) ([]string, error) {
	res := make([]string, 0)
	if len(pathSets) == 0 || (len(pathSets) == 1 && pathSets[0] == "") {
		return res, nil
	}

	for _, pathSet := range pathSets {
		keys := strings.Split(pathSet, ",")
		for _, k := range keys {
			if strings.HasPrefix(k, "~") {
				home := os.Getenv("HOME")
				if home == "" {
					return nil, fmt.Errorf("HOME is not defined for key '%s'", k)
				}
				k = strings.Replace(k, "~", home, 1)
			}

			keyPath, err := filepath.Abs(k)
			if err != nil {
				return nil, fmt.Errorf("get absolute path for '%s': %v", k, err)
			}
			res = append(res, keyPath)
		}
	}
	return res, nil
}

func DefineBecomeFlags(cmd *kingpin.CmdClause) {
	// Ansible compatible
	cmd.Flag("ask-become-pass", "Ask for sudo password before the installation process.").
		Envar(configEnvName("ASK_BECOME_PASS")).
		Short('K').
		BoolVar(&AskBecomePass)
}
