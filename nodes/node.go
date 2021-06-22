// Copyright 2020 Nokia
// Licensed under the BSD 3-Clause License.
// SPDX-License-Identifier: BSD-3-Clause

package nodes

import (
	"context"

	"github.com/srl-labs/containerlab/runtime"
	"github.com/srl-labs/containerlab/types"
)

const (
	// default connection mode for vrnetlab based containers
	VrDefConnMode = "tc"
)

type Node interface {
	Init(*types.NodeConfig, ...NodeOption) error
	Config() *types.NodeConfig
	PreDeploy(configName, labCADir, labCARoot string) error
	Deploy(context.Context, runtime.ContainerRuntime) error
	PostDeploy(context.Context, runtime.ContainerRuntime, map[string]Node) error
	WithMgmtNet(*types.MgmtNet)
}

var Nodes = map[string]Initializer{}

type Initializer func() Node

func Register(name string, initFn Initializer) {
	Nodes[name] = initFn
}

type NodeOption func(Node)

func WithMgmtNet(mgmt *types.MgmtNet) NodeOption {
	return func(n Node) {
		n.WithMgmtNet(mgmt)
	}
}

var DefaultConfigTemplates = map[string]string{
	"srl":     "/etc/containerlab/templates/srl/srlconfig.tpl",
	"ceos":    "/etc/containerlab/templates/arista/ceos.cfg.tpl",
	"crpd":    "/etc/containerlab/templates/crpd/juniper.conf",
	"vr-sros": "",
}
