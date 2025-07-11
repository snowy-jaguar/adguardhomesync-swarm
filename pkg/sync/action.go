// Copyright 2025 snowy-jaguar
// Contact: @snowyjaguar (Discord)
// Contact: contact@snowyjaguar.xyz (Email)
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

package sync

import (
	"go.uber.org/zap"

	"github.com/snowy-jaguar/adguardhomesync-swarm/pkg/client"
	"github.com/snowy-jaguar/adguardhomesync-swarm/pkg/client/model"
	"github.com/snowy-jaguar/adguardhomesync-swarm/pkg/types"
)

func setupActions(cfg *types.Config) (actions []syncAction) {
	if cfg.Features.GeneralSettings {
		actions = append(actions,
			action("profile info", actionProfileInfo),
			action("protection", actionProtection),
			action("parental", actionParental),
			action("safe search config", actionSafeSearchConfig),
			action("safe browsing", actionSafeBrowsing),
		)
	}
	if cfg.Features.DNS.ServerConfig {
		actions = append(actions,
			action("DNS server config", actionDNSServerConfig),
		)
	}
	if cfg.Features.QueryLogConfig {
		actions = append(actions,
			action("query log config", actionQueryLogConfig),
		)
	}
	if cfg.Features.StatsConfig {
		actions = append(actions,
			action("stats config", actionStatsConfig),
		)
	}
	if cfg.Features.DNS.Rewrites {
		actions = append(actions,
			action("DNS rewrites", actionDNSRewrites),
		)
	}
	if cfg.Features.Filters {
		actions = append(actions,
			action("actionFilters", actionFilters),
		)
	}
	if cfg.Features.Services {
		actions = append(actions,
			action("blocked services schedule", actionBlockedServicesSchedule),
		)
	}
	if cfg.Features.ClientSettings {
		actions = append(actions,
			action("client settings", actionClientSettings),
		)
	}
	if cfg.Features.DNS.AccessLists {
		actions = append(actions,
			action("DNS access lists", actionDNSAccessLists),
		)
	}
	if cfg.Features.DHCP.ServerConfig {
		actions = append(actions,
			action("DHCP server config", actionDHCPServerConfig),
		)
	}
	if cfg.Features.DHCP.StaticLeases {
		actions = append(actions,
			action("DHCP static leases", actionDHCPStaticLeases),
		)
	}
	return actions
}

type syncAction interface {
	sync(ac *actionContext) error
	name() string
}

type actionContext struct {
	rl            *zap.SugaredLogger
	origin        *origin
	client        client.Client
	replicaStatus *model.ServerStatus
	replica       types.AdGuardInstance
	cfg           *types.Config
}

type defaultAction struct {
	myName string
	doSync func(ac *actionContext) error
}

func action(name string, f func(ac *actionContext) error) syncAction {
	return &defaultAction{myName: name, doSync: f}
}

func (d *defaultAction) sync(ac *actionContext) error {
	return d.doSync(ac)
}

func (d *defaultAction) name() string {
	return d.myName
}
