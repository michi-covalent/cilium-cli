// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cilium/cilium/cilium-cli/hooks"
	gops "github.com/google/gops/agent"

	"github.com/cilium/cilium-cli/cli"
	_ "github.com/cilium/cilium-cli/internal/logging" // necessary to disable unwanted cfssl log messages
)

func main() {
	if err := gops.Listen(gops.Options{}); err != nil {
		log.Printf("Unable to start gops: %s", err)
	}

	if err := cli.NewCiliumCommand(&hooks.MichiHooks{}).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
