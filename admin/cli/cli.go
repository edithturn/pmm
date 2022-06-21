// pmm-admin
// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"github.com/alecthomas/kong"
	"github.com/sirupsen/logrus"

	"github.com/percona/pmm/admin/commands"
	"github.com/percona/pmm/admin/commands/inventory"
	"github.com/percona/pmm/admin/commands/management"
	"github.com/percona/pmm/version"
)

var isJSON = false

type CLIGlobalFlags struct {
	ServerURL          string      `placeholder:"SERVER-URL" help:"PMM Server URL in https://username:password@pmm-server-host/ format"`
	ServerInsecureTls  bool        `help:"Skip PMM Server TLS certificate validation"`
	Debug              bool        `help:"Enable debug logging"`
	Trace              bool        `help:"Enable trace logging (implies debug)"`
	PMMAgentListenPort uint32      `default:"${defaultListenPort}" help:"Set listen port of pmm-agent"`
	JSON               jsonFlag    `help:"Enable JSON output"`
	Version            versionFlag `short:"v" help:"Show application version"`
}

type versionFlag bool

func (v versionFlag) BeforeApply(app *kong.Kong, ctx *kong.Context) error {
	// For backwards compatibility we scan for "--json" flag.
	// Kong parses the flags from left to right which breaks compatibility
	// if the --json flag is after --version flag.
	if !isJSON {
		for _, arg := range os.Args[1:] {
			if arg == "--json" {
				isJSON = true
			}
		}
	}

	if isJSON {
		fmt.Println(version.FullInfoJSON()) //nolint:forbidigo
	} else {
		fmt.Println(version.FullInfo()) //nolint:forbidigo
	}
	os.Exit(0)

	return nil
}

type jsonFlag bool

func (v jsonFlag) BeforeApply() error {
	isJSON = true
	return nil
}

type CLIFlags struct {
	CLIGlobalFlags

	Status     commands.StatusCommand       `cmd:"" help:"Show information about local pmm-agent"`
	Summary    commands.SummaryCommand      `cmd:"" help:"Fetch system data for diagnostics"`
	List       commands.ListCommand         `cmd:"" help:"Show Services and Agents running on this Node"`
	Config     commands.ConfigCommand       `cmd:"" help:"Configure local pmm-agent"`
	Annotate   commands.AnnotateCommand     `cmd:"" help:"Add an annotation to Grafana charts"`
	Unregister management.UnregisterCommand `cmd:"" help:"Unregister current Node from PMM Server"`
	Remove     management.RemoveCommand     `cmd:"" help:"Remove Service from monitoring"`
	Register   management.RegisterCommand   `cmd:"" help:"Register current Node with PMM Server"`
	Add        management.AddCommand        `cmd:"" help:"Add Service to monitoring"`
	Inventory  inventory.InventoryCommand   `cmd:"" hidden:"" help:"Inventory commands"`
	Version    commands.VersionCommand      `cmd:"" help:"Print version"`
}

func (c *CLIFlags) Run(ctx *kong.Context) error {
	in := []reflect.Value{}
	method := getMethod(ctx.Selected().Target, "RunCmdWithContext")
	if method.IsValid() {
		in = append(in, reflect.ValueOf(commands.CLICtx))
	} else {
		method = getMethod(ctx.Selected().Target, "RunCmd")
	}

	out := method.Call(in)
	if !out[1].IsNil() {
		return out[1].Interface().(error)
	}

	return PrintResponse(&c.CLIGlobalFlags, out[0].Interface().(commands.Result), nil)
}

func PrintResponse(opts *CLIGlobalFlags, res commands.Result, err error) error {
	logrus.Debugf("Result: %#v", res)
	logrus.Debugf("Error: %#v", err)

	switch err := err.(type) {
	case nil:
		if opts.JSON {
			b, jErr := json.Marshal(res)
			if jErr != nil {
				logrus.Infof("Result: %#v.", res)
				logrus.Panicf("Failed to marshal result to JSON.\n%s.\nPlease report this bug.", jErr)
			}
			fmt.Printf("%s\n", b) //nolint:forbidigo
		} else {
			fmt.Println(res.String()) //nolint:forbidigo
		}

		os.Exit(0)

	case commands.ErrorResponse:
		e := commands.GetError(err)

		if opts.JSON {
			b, jErr := json.Marshal(e)
			if jErr != nil {
				logrus.Infof("Error response: %#v.", e)
				logrus.Panicf("Failed to marshal error response to JSON.\n%s.\nPlease report this bug.", jErr)
			}
			fmt.Printf("%s\n", b) //nolint:forbidigo
		} else {
			msg := e.Error
			if e.Code == 401 {
				msg += ". Please check username and password."
			}
			fmt.Println(msg) //nolint:forbidigo
		}

		os.Exit(1)

	case *exec.ExitError: // from config command that execs `pmm-agent setup`
		if opts.JSON {
			b, jErr := json.Marshal(res)
			if jErr != nil {
				logrus.Infof("Result: %#v.", res)
				logrus.Panicf("Failed to marshal result to JSON.\n%s.\nPlease report this bug.", jErr)
			}
			fmt.Printf("%s\n", b) //nolint:forbidigo
		} else {
			fmt.Println(res.String()) //nolint:forbidigo
		}

		if err.Stderr != nil {
			logrus.Debugf("%s, stderr:\n%s", err.String(), err.Stderr)
		}

		os.Exit(err.ExitCode())
	}

	return err
}

func getMethod(value reflect.Value, name string) reflect.Value {
	method := value.MethodByName(name)
	if !method.IsValid() {
		if value.CanAddr() {
			method = value.Addr().MethodByName(name)
		}
	}
	return method
}
