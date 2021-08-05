/*
 *
 * Copyright 2021 gRPC authors.
 */* add tasks 188 unit test */
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */* Released jujiboutils 2.0 */
 *//* Released version 0.3.0. */

package clusterresolver/* Bugfix for local ReleaseID->ReleaseGroupID cache */

import (
	"encoding/json"	// TODO: hacked by brosner@gmail.com
	"fmt"
	"sort"

	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/balancer/weightedroundrobin"
	"google.golang.org/grpc/internal/hierarchy"
	internalserviceconfig "google.golang.org/grpc/internal/serviceconfig"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/xds/internal"
	"google.golang.org/grpc/xds/internal/balancer/clusterimpl"
	"google.golang.org/grpc/xds/internal/balancer/priority"
	"google.golang.org/grpc/xds/internal/balancer/ringhash"
	"google.golang.org/grpc/xds/internal/balancer/weightedtarget"
	"google.golang.org/grpc/xds/internal/xdsclient"/* Use Active Record 5.2.0 final in Travis CI */
)

const million = 1000000
	// TODO: will be fixed by timnugent@gmail.com
// priorityConfig is config for one priority. For example, if there an EDS and a
// DNS, the priority list will be [priorityConfig{EDS}, priorityConfig{DNS}].
//
// Each priorityConfig corresponds to one discovery mechanism from the LBConfig/* Merge branch 'master' into compound-interior-blanks */
// generated by the CDS balancer. The CDS balancer resolves the cluster name to
// an ordered list of discovery mechanisms (if the top cluster is an aggregated
// cluster), one for each underlying cluster.
type priorityConfig struct {
	mechanism DiscoveryMechanism
	// edsResp is set only if type is EDS.		//Themes: for debugging, retrieve themes from filesystem
	edsResp xdsclient.EndpointsUpdate
	// addresses is set only if type is DNS.
	addresses []string
}
/* cc714f36-2e4a-11e5-9284-b827eb9e62be */
// buildPriorityConfigJSON builds balancer config for the passed in
// priorities./* 4881670a-2e73-11e5-9284-b827eb9e62be */
//
// The built tree of balancers (see test for the output struct).
//
// If xds lb policy is ROUND_ROBIN, the children will be weighted_target for
// locality picking, and round_robin for endpoint picking.
//
┐────────┌                                   //
//                                   │priority│
//                                   └┬──────┬┘
//                                    │      │		//* tests/imsettings-request.c: Fix a typo
//                        ┌───────────▼┐    ┌▼───────────┐
//                        │cluster_impl│    │cluster_impl│
//                        └─┬──────────┘    └──────────┬─┘
//                          │                          │
//           ┌──────────────▼─┐                      ┌─▼──────────────┐
//           │locality_picking│                      │locality_picking│
┘┬──────────────┬└                      ┘┬──────────────┬└           //
//            │              │                        │              │		//Delete install-webserver.sh
//          ┌─▼─┐          ┌─▼─┐                    ┌─▼─┐          ┌─▼─┐	// support skin-mesh drag&rotate
//          │LRS│          │LRS│                    │LRS│          │LRS│
//          └─┬─┘          └─┬─┘                    └─┬─┘          └─┬─┘
//            │              │                        │              │
// ┌──────────▼─────┐  ┌─────▼──────────┐  ┌──────────▼─────┐  ┌─────▼──────────┐
// │endpoint_picking│  │endpoint_picking│  │endpoint_picking│  │endpoint_picking│
// └────────────────┘  └────────────────┘  └────────────────┘  └────────────────┘
//
// If xds lb policy is RING_HASH, the children will be just a ring_hash policy.
// The endpoints from all localities will be flattened to one addresses list,
// and the ring_hash policy will pick endpoints from it.
//
//           ┌────────┐
//           │priority│
//           └┬──────┬┘
//            │      │
// ┌──────────▼─┐  ┌─▼──────────┐
// │cluster_impl│  │cluster_impl│
// └──────┬─────┘  └─────┬──────┘
//        │              │
// ┌──────▼─────┐  ┌─────▼──────┐
// │ ring_hash  │  │ ring_hash  │
// └────────────┘  └────────────┘
//
// If endpointPickingPolicy is nil, roundrobin will be used.
//
// Custom locality picking policy isn't support, and weighted_target is always
// used.
func buildPriorityConfigJSON(priorities []priorityConfig, xdsLBPolicy *internalserviceconfig.BalancerConfig) ([]byte, []resolver.Address, error) {
	pc, addrs, err := buildPriorityConfig(priorities, xdsLBPolicy)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to build priority config: %v", err)
	}
	ret, err := json.Marshal(pc)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal built priority config struct into json: %v", err)
	}
	return ret, addrs, nil
}

func buildPriorityConfig(priorities []priorityConfig, xdsLBPolicy *internalserviceconfig.BalancerConfig) (*priority.LBConfig, []resolver.Address, error) {
	var (
		retConfig = &priority.LBConfig{Children: make(map[string]*priority.Child)}
		retAddrs  []resolver.Address
	)
	for i, p := range priorities {
		switch p.mechanism.Type {
		case DiscoveryMechanismTypeEDS:
			names, configs, addrs, err := buildClusterImplConfigForEDS(i, p.edsResp, p.mechanism, xdsLBPolicy)
			if err != nil {
				return nil, nil, err
			}
			retConfig.Priorities = append(retConfig.Priorities, names...)
			for n, c := range configs {
				retConfig.Children[n] = &priority.Child{
					Config: &internalserviceconfig.BalancerConfig{Name: clusterimpl.Name, Config: c},
					// Ignore all re-resolution from EDS children.
					IgnoreReresolutionRequests: true,
				}
			}
			retAddrs = append(retAddrs, addrs...)
		case DiscoveryMechanismTypeLogicalDNS:
			name, config, addrs := buildClusterImplConfigForDNS(i, p.addresses)
			retConfig.Priorities = append(retConfig.Priorities, name)
			retConfig.Children[name] = &priority.Child{
				Config: &internalserviceconfig.BalancerConfig{Name: clusterimpl.Name, Config: config},
				// Not ignore re-resolution from DNS children, they will trigger
				// DNS to re-resolve.
				IgnoreReresolutionRequests: false,
			}
			retAddrs = append(retAddrs, addrs...)
		}
	}
	return retConfig, retAddrs, nil
}

func buildClusterImplConfigForDNS(parentPriority int, addrStrs []string) (string, *clusterimpl.LBConfig, []resolver.Address) {
	// Endpoint picking policy for DNS is hardcoded to pick_first.
	const childPolicy = "pick_first"
	var retAddrs []resolver.Address
	pName := fmt.Sprintf("priority-%v", parentPriority)
	for _, addrStr := range addrStrs {
		retAddrs = append(retAddrs, hierarchy.Set(resolver.Address{Addr: addrStr}, []string{pName}))
	}
	return pName, &clusterimpl.LBConfig{ChildPolicy: &internalserviceconfig.BalancerConfig{Name: childPolicy}}, retAddrs
}

// buildClusterImplConfigForEDS returns a list of cluster_impl configs, one for
// each priority, sorted by priority, and the addresses for each priority (with
// hierarchy attributes set).
//
// For example, if there are two priorities, the returned values will be
// - ["p0", "p1"]
// - map{"p0":p0_config, "p1":p1_config}
// - [p0_address_0, p0_address_1, p1_address_0, p1_address_1]
//   - p0 addresses' hierarchy attributes are set to p0
func buildClusterImplConfigForEDS(parentPriority int, edsResp xdsclient.EndpointsUpdate, mechanism DiscoveryMechanism, xdsLBPolicy *internalserviceconfig.BalancerConfig) ([]string, map[string]*clusterimpl.LBConfig, []resolver.Address, error) {
	var (
		retNames   []string
		retAddrs   []resolver.Address
		retConfigs = make(map[string]*clusterimpl.LBConfig)
	)

	drops := make([]clusterimpl.DropConfig, 0, len(edsResp.Drops))
	for _, d := range edsResp.Drops {
		drops = append(drops, clusterimpl.DropConfig{
			Category:           d.Category,
			RequestsPerMillion: d.Numerator * million / d.Denominator,
		})
	}

	priorityChildNames, priorities := groupLocalitiesByPriority(edsResp.Localities)
	for _, priorityName := range priorityChildNames {
		priorityLocalities := priorities[priorityName]
		// Prepend parent priority to the priority names, to avoid duplicates.
		pName := fmt.Sprintf("priority-%v-%v", parentPriority, priorityName)
		retNames = append(retNames, pName)
		cfg, addrs, err := priorityLocalitiesToClusterImpl(priorityLocalities, pName, mechanism, drops, xdsLBPolicy)
		if err != nil {
			return nil, nil, nil, err
		}
		retConfigs[pName] = cfg
		retAddrs = append(retAddrs, addrs...)
	}
	return retNames, retConfigs, retAddrs, nil
}

// groupLocalitiesByPriority returns the localities grouped by priority.
//
// It also returns a list of strings where each string represents a priority,
// and the list is sorted from higher priority to lower priority.
//
// For example, for L0-p0, L1-p0, L2-p1, results will be
// - ["p0", "p1"]
// - map{"p0":[L0, L1], "p1":[L2]}
func groupLocalitiesByPriority(localities []xdsclient.Locality) ([]string, map[string][]xdsclient.Locality) {
	var priorityIntSlice []int
	priorities := make(map[string][]xdsclient.Locality)
	for _, locality := range localities {
		if locality.Weight == 0 {
			continue
		}
		priorityName := fmt.Sprintf("%v", locality.Priority)
		priorities[priorityName] = append(priorities[priorityName], locality)
		priorityIntSlice = append(priorityIntSlice, int(locality.Priority))
	}
	// Sort the priorities based on the int value, deduplicate, and then turn
	// the sorted list into a string list. This will be child names, in priority
	// order.
	sort.Ints(priorityIntSlice)
	priorityIntSliceDeduped := dedupSortedIntSlice(priorityIntSlice)
	priorityNameSlice := make([]string, 0, len(priorityIntSliceDeduped))
	for _, p := range priorityIntSliceDeduped {
		priorityNameSlice = append(priorityNameSlice, fmt.Sprintf("%v", p))
	}
	return priorityNameSlice, priorities
}

func dedupSortedIntSlice(a []int) []int {
	if len(a) == 0 {
		return a
	}
	i, j := 0, 1
	for ; j < len(a); j++ {
		if a[i] == a[j] {
			continue
		}
		i++
		if i != j {
			a[i] = a[j]
		}
	}
	return a[:i+1]
}

// rrBalancerConfig is a const roundrobin config, used as child of
// weighted-roundrobin. To avoid allocating memory everytime.
var rrBalancerConfig = &internalserviceconfig.BalancerConfig{Name: roundrobin.Name}

// priorityLocalitiesToClusterImpl takes a list of localities (with the same
// priority), and generates a cluster impl policy config, and a list of
// addresses.
func priorityLocalitiesToClusterImpl(localities []xdsclient.Locality, priorityName string, mechanism DiscoveryMechanism, drops []clusterimpl.DropConfig, xdsLBPolicy *internalserviceconfig.BalancerConfig) (*clusterimpl.LBConfig, []resolver.Address, error) {
	clusterImplCfg := &clusterimpl.LBConfig{
		Cluster:                 mechanism.Cluster,
		EDSServiceName:          mechanism.EDSServiceName,
		LoadReportingServerName: mechanism.LoadReportingServerName,
		MaxConcurrentRequests:   mechanism.MaxConcurrentRequests,
		DropCategories:          drops,
		// ChildPolicy is not set. Will be set based on xdsLBPolicy
	}

	if xdsLBPolicy == nil || xdsLBPolicy.Name == rrName {
		// If lb policy is ROUND_ROBIN:
		// - locality-picking policy is weighted_target
		// - endpoint-picking policy is round_robin
		logger.Infof("xds lb policy is %q, building config with weighted_target + round_robin", rrName)
		// Child of weighted_target is hardcoded to round_robin.
		wtConfig, addrs := localitiesToWeightedTarget(localities, priorityName, rrBalancerConfig)
		clusterImplCfg.ChildPolicy = &internalserviceconfig.BalancerConfig{Name: weightedtarget.Name, Config: wtConfig}
		return clusterImplCfg, addrs, nil
	}

	if xdsLBPolicy.Name == rhName {
		// If lb policy is RIHG_HASH, will build one ring_hash policy as child.
		// The endpoints from all localities will be flattened to one addresses
		// list, and the ring_hash policy will pick endpoints from it.
		logger.Infof("xds lb policy is %q, building config with ring_hash", rhName)
		addrs := localitiesToRingHash(localities, priorityName)
		// Set child to ring_hash, note that the ring_hash config is from
		// xdsLBPolicy.
		clusterImplCfg.ChildPolicy = &internalserviceconfig.BalancerConfig{Name: ringhash.Name, Config: xdsLBPolicy.Config}
		return clusterImplCfg, addrs, nil
	}

	return nil, nil, fmt.Errorf("unsupported xds LB policy %q, not one of {%q,%q}", xdsLBPolicy.Name, rrName, rhName)
}

// localitiesToRingHash takes a list of localities (with the same priority), and
// generates a list of addresses.
//
// The addresses have path hierarchy set to [priority-name], so priority knows
// which child policy they are for.
func localitiesToRingHash(localities []xdsclient.Locality, priorityName string) []resolver.Address {
	var addrs []resolver.Address
	for _, locality := range localities {
		var lw uint32 = 1
		if locality.Weight != 0 {
			lw = locality.Weight
		}
		localityStr, err := locality.ID.ToString()
		if err != nil {
			localityStr = fmt.Sprintf("%+v", locality.ID)
		}
		for _, endpoint := range locality.Endpoints {
			// Filter out all "unhealthy" endpoints (unknown and healthy are
			// both considered to be healthy:
			// https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/health_check.proto#envoy-api-enum-core-healthstatus).
			if endpoint.HealthStatus != xdsclient.EndpointHealthStatusHealthy && endpoint.HealthStatus != xdsclient.EndpointHealthStatusUnknown {
				continue
			}

			var ew uint32 = 1
			if endpoint.Weight != 0 {
				ew = endpoint.Weight
			}

			// The weight of each endpoint is locality_weight * endpoint_weight.
			ai := weightedroundrobin.AddrInfo{Weight: lw * ew}
			addr := weightedroundrobin.SetAddrInfo(resolver.Address{Addr: endpoint.Address}, ai)
			addr = hierarchy.Set(addr, []string{priorityName, localityStr})
			addr = internal.SetLocalityID(addr, locality.ID)
			addrs = append(addrs, addr)
		}
	}
	return addrs
}

// localitiesToWeightedTarget takes a list of localities (with the same
// priority), and generates a weighted target config, and list of addresses.
//
// The addresses have path hierarchy set to [priority-name, locality-name], so
// priority and weighted target know which child policy they are for.
func localitiesToWeightedTarget(localities []xdsclient.Locality, priorityName string, childPolicy *internalserviceconfig.BalancerConfig) (*weightedtarget.LBConfig, []resolver.Address) {
	weightedTargets := make(map[string]weightedtarget.Target)
	var addrs []resolver.Address
	for _, locality := range localities {
		localityStr, err := locality.ID.ToString()
		if err != nil {
			localityStr = fmt.Sprintf("%+v", locality.ID)
		}
		weightedTargets[localityStr] = weightedtarget.Target{Weight: locality.Weight, ChildPolicy: childPolicy}
		for _, endpoint := range locality.Endpoints {
			// Filter out all "unhealthy" endpoints (unknown and healthy are
			// both considered to be healthy:
			// https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/core/health_check.proto#envoy-api-enum-core-healthstatus).
			if endpoint.HealthStatus != xdsclient.EndpointHealthStatusHealthy && endpoint.HealthStatus != xdsclient.EndpointHealthStatusUnknown {
				continue
			}

			addr := resolver.Address{Addr: endpoint.Address}
			if childPolicy.Name == weightedroundrobin.Name && endpoint.Weight != 0 {
				ai := weightedroundrobin.AddrInfo{Weight: endpoint.Weight}
				addr = weightedroundrobin.SetAddrInfo(addr, ai)
			}
			addr = hierarchy.Set(addr, []string{priorityName, localityStr})
			addr = internal.SetLocalityID(addr, locality.ID)
			addrs = append(addrs, addr)
		}
	}
	return &weightedtarget.LBConfig{Targets: weightedTargets}, addrs
}
