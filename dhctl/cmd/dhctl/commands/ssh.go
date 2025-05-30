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

package commands

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"path"
	"strings"

	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"github.com/deckhouse/deckhouse/dhctl/pkg/app"
	"github.com/deckhouse/deckhouse/dhctl/pkg/config"
	"github.com/deckhouse/deckhouse/dhctl/pkg/log"
	"github.com/deckhouse/deckhouse/dhctl/pkg/system/node/ssh"
)

func DefineTestSSHConnectionCommand(cmd *kingpin.CmdClause) *kingpin.CmdClause {
	app.DefineSSHFlags(cmd, config.ConnectionConfigParser{})
	app.DefineBecomeFlags(cmd)

	cmd.Action(func(c *kingpin.ParseContext) error {
		sshCl, err := ssh.NewClientFromFlagsWithHosts()
		if err != nil {
			return err
		}
		sshCl, err = sshCl.Start()
		if err != nil {
			return err
		}

		err = sshCl.Check().AwaitAvailability(context.Background())

		if err != nil {
			return fmt.Errorf("check connection: %v", err)
		}

		TestCommandDelay()

		return nil
	})
	return cmd
}

func DefineTestSCPCommand(cmd *kingpin.CmdClause) *kingpin.CmdClause {
	var SrcPath string
	var DstPath string
	var Data string
	var Direction string

	app.DefineSSHFlags(cmd, config.ConnectionConfigParser{})
	app.DefineBecomeFlags(cmd)

	cmd.Flag("src", "source path").Short('s').StringVar(&SrcPath)
	cmd.Flag("dst", "destination path").Short('d').StringVar(&DstPath)
	cmd.Flag("data", "data to test uploadbytes method").StringVar(&Data)
	cmd.Flag("way", "transfer direction: 'up' to upload to remote or 'down' to download from remote").Short('w').StringVar(&Direction)
	cmd.Action(func(c *kingpin.ParseContext) error {
		log.DebugLn("scp: start ssh-agent")
		sshCl, err := ssh.NewClientFromFlagsWithHosts()
		if err != nil {
			return err
		}
		sshCl, err = sshCl.Start()
		if err != nil {
			return err
		}

		log.DebugLn("scp: start")

		success := false
		if Direction == "up" {
			if Data != "" {
				log.InfoF("upload bytes to '%s' on remote\n", DstPath)
				err = sshCl.File().UploadBytes(context.Background(), []byte(Data), DstPath)
			} else {
				log.InfoF("upload local '%s' to '%s' on remote\n", SrcPath, DstPath)
				err = sshCl.File().Upload(context.Background(), SrcPath, DstPath)
			}
			if err != nil {
				return err
			}
			success = true
		} else {
			if DstPath == "stdout" {
				log.InfoF("download bytes from remote '%s'\n", SrcPath)
				data, err := sshCl.File().DownloadBytes(context.Background(), SrcPath)
				if err != nil {
					return err
				}
				log.InfoLn(string(data))
				success = true
			} else {
				log.InfoF("download bytes from remote '%s' to local '%s'\n", SrcPath, DstPath)
				err = sshCl.File().Download(context.Background(), SrcPath, DstPath)
				if err != nil {
					return err
				}
				success = true
			}
		}

		if !success {
			log.InfoLn("unrecognized flags")
		}

		return nil
	})

	return cmd
}

func DefineTestUploadExecCommand(cmd *kingpin.CmdClause) *kingpin.CmdClause {
	var ScriptPath string
	var Sudo bool

	app.DefineSSHFlags(cmd, config.ConnectionConfigParser{})
	app.DefineBecomeFlags(cmd)
	cmd.Flag("script", "source path").
		StringVar(&ScriptPath)
	cmd.Flag("sudo", "source path").Short('s').
		BoolVar(&Sudo)

	cmd.Action(func(c *kingpin.ParseContext) error {
		sshClient, err := ssh.NewInitClientFromFlagsWithHosts(true)
		if err != nil {
			return nil
		}

		cmd := sshClient.UploadScript(ScriptPath)
		if Sudo {
			cmd.Sudo()
		}
		var stdout []byte
		stdout, err = cmd.Execute(context.Background())
		if err != nil {
			var ee *exec.ExitError
			if errors.As(err, &ee) {
				return fmt.Errorf("script '%s' error: %w stderr: %s", ScriptPath, err, string(ee.Stderr))
			}
			return fmt.Errorf("script '%s' error: %w", ScriptPath, err)
		}
		log.InfoF("stdout: %s\n", strings.Trim(string(stdout), "\n "))
		log.InfoF("Got %d symbols\n", len(stdout))
		return nil
	})

	return cmd
}

func DefineTestBundle(cmd *kingpin.CmdClause) *kingpin.CmdClause {
	var ScriptName string
	var BundleDir string

	app.DefineSSHFlags(cmd, config.ConnectionConfigParser{})
	app.DefineBecomeFlags(cmd)
	cmd.Flag("bundle-dir", "path of a bundle root directory").
		Short('d').
		StringVar(&BundleDir)
	cmd.Flag("bundle-script", "path of a bundle main script").
		Short('s').
		StringVar(&ScriptName)

	cmd.Action(func(c *kingpin.ParseContext) error {
		sshClient, err := ssh.NewInitClientFromFlagsWithHosts(true)
		if err != nil {
			return nil
		}

		cmd := sshClient.UploadScript(ScriptName)
		cmd.Sudo()
		parentDir := path.Dir(BundleDir)
		bundleDir := path.Base(BundleDir)
		stdout, err := cmd.ExecuteBundle(context.Background(), parentDir, bundleDir)
		if err != nil {
			var ee *exec.ExitError
			if errors.As(err, &ee) {
				return fmt.Errorf("bundle '%s' error: %w\nstderr: %s\n", bundleDir, err, string(ee.Stderr))
			}
			return fmt.Errorf("bundle '%s' error: %w", bundleDir, err)
		}
		log.InfoF("Got %d symbols\n", len(stdout))
		return nil
	})

	return cmd
}
