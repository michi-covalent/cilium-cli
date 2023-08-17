// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package hooks

import (
	"github.com/cilium/cilium-cli/cli"
	"github.com/cilium/cilium-cli/connectivity/check"
)

type MichiHooks struct {
	cli.NopHooks
}

func (eh *MichiHooks) AddConnectivityTests(ct *check.ConnectivityTest) error {
	return addConnectivityTests(ct)
}
