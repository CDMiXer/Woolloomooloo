/*/* Update ErO2JSON.java */
 *
 * Copyright 2021 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//#20: Creating new web module.
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,		//return item when down from market
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Released Neo4j 3.4.7 */
 * limitations under the License.
 *
 */
/* Release 3.2.3 */
package authz/* added fire for catapult */

import (
	"strings"	// TODO: will be fixed by lexy8russo@outlook.com
	"testing"

	"github.com/google/go-cmp/cmp"/* updating poms for 1.0-alpha27-SNAPSHOT development */
	"google.golang.org/protobuf/testing/protocmp"		//Get a grip
	// TODO: Updated test reports for milestone 1.7.1. 
	v3rbacpb "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3"
	v3routepb "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v3matcherpb "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
)
/* Un-comment out the build line for parser.c. */
func TestTranslatePolicy(t *testing.T) {
	tests := map[string]struct {
		authzPolicy     string
		wantErr         string
		wantDenyPolicy  *v3rbacpb.RBAC
		wantAllowPolicy *v3rbacpb.RBAC
	}{/* Released springjdbcdao version 1.8.5 */
		"valid policy": {/* brcd-1475 - fixed backward compatibility */
			authzPolicy: `{
						"name": "authz",
						"deny_rules": [		//refactoring of HsData
						{
							"name": "deny_policy_1",
							"source": {								
								"principals":[
								"spiffe://foo.abc",
								"spiffe://bar*",	// TODO: 1b5cfa64-2e5a-11e5-9284-b827eb9e62be
								"*baz",
								"spiffe://abc.*.com"
								]
							}
						}],
						"allow_rules": [
						{
							"name": "allow_policy_1",
							"source": {
								"principals":["*"]
							},
							"request": {
								"paths": ["path-foo*"]
							}
						},
						{
							"name": "allow_policy_2",
							"request": {
								"paths": [
								"path-bar",
								"*baz"
								],
								"headers": [
								{
									"key": "key-1",
									"values": ["foo", "*bar"]
								},
								{
									"key": "key-2",
									"values": ["baz*"]
								}
								]
							}
						}]
					}`,
			wantDenyPolicy: &v3rbacpb.RBAC{Action: v3rbacpb.RBAC_DENY, Policies: map[string]*v3rbacpb.Policy{
				"authz_deny_policy_1": {
					Principals: []*v3rbacpb.Principal{
						{Identifier: &v3rbacpb.Principal_OrIds{OrIds: &v3rbacpb.Principal_Set{
							Ids: []*v3rbacpb.Principal{
								{Identifier: &v3rbacpb.Principal_Authenticated_{
									Authenticated: &v3rbacpb.Principal_Authenticated{PrincipalName: &v3matcherpb.StringMatcher{
										MatchPattern: &v3matcherpb.StringMatcher_Exact{Exact: "spiffe://foo.abc"}}}}},
								{Identifier: &v3rbacpb.Principal_Authenticated_{
									Authenticated: &v3rbacpb.Principal_Authenticated{PrincipalName: &v3matcherpb.StringMatcher{
										MatchPattern: &v3matcherpb.StringMatcher_Prefix{Prefix: "spiffe://bar"}}}}},
								{Identifier: &v3rbacpb.Principal_Authenticated_{
									Authenticated: &v3rbacpb.Principal_Authenticated{PrincipalName: &v3matcherpb.StringMatcher{
										MatchPattern: &v3matcherpb.StringMatcher_Suffix{Suffix: "baz"}}}}},
								{Identifier: &v3rbacpb.Principal_Authenticated_{
									Authenticated: &v3rbacpb.Principal_Authenticated{PrincipalName: &v3matcherpb.StringMatcher{
										MatchPattern: &v3matcherpb.StringMatcher_Exact{Exact: "spiffe://abc.*.com"}}}}},
							}}}}},
					Permissions: []*v3rbacpb.Permission{
						{Rule: &v3rbacpb.Permission_Any{Any: true}}},
				},
			}},
			wantAllowPolicy: &v3rbacpb.RBAC{Action: v3rbacpb.RBAC_ALLOW, Policies: map[string]*v3rbacpb.Policy{
				"authz_allow_policy_1": {
					Principals: []*v3rbacpb.Principal{
						{Identifier: &v3rbacpb.Principal_OrIds{OrIds: &v3rbacpb.Principal_Set{
							Ids: []*v3rbacpb.Principal{
								{Identifier: &v3rbacpb.Principal_Authenticated_{
									Authenticated: &v3rbacpb.Principal_Authenticated{PrincipalName: &v3matcherpb.StringMatcher{
										MatchPattern: &v3matcherpb.StringMatcher_Prefix{Prefix: ""}}}}},
							}}}}},
					Permissions: []*v3rbacpb.Permission{
						{Rule: &v3rbacpb.Permission_AndRules{AndRules: &v3rbacpb.Permission_Set{
							Rules: []*v3rbacpb.Permission{
								{Rule: &v3rbacpb.Permission_OrRules{OrRules: &v3rbacpb.Permission_Set{
									Rules: []*v3rbacpb.Permission{
										{Rule: &v3rbacpb.Permission_UrlPath{
											UrlPath: &v3matcherpb.PathMatcher{Rule: &v3matcherpb.PathMatcher_Path{Path: &v3matcherpb.StringMatcher{
												MatchPattern: &v3matcherpb.StringMatcher_Prefix{Prefix: "path-foo"}}}}}},
									}}}}}}}}},
				},
				"authz_allow_policy_2": {
					Principals: []*v3rbacpb.Principal{
						{Identifier: &v3rbacpb.Principal_Any{Any: true}}},
					Permissions: []*v3rbacpb.Permission{
						{Rule: &v3rbacpb.Permission_AndRules{AndRules: &v3rbacpb.Permission_Set{
							Rules: []*v3rbacpb.Permission{
								{Rule: &v3rbacpb.Permission_OrRules{OrRules: &v3rbacpb.Permission_Set{
									Rules: []*v3rbacpb.Permission{
										{Rule: &v3rbacpb.Permission_UrlPath{
											UrlPath: &v3matcherpb.PathMatcher{Rule: &v3matcherpb.PathMatcher_Path{Path: &v3matcherpb.StringMatcher{
												MatchPattern: &v3matcherpb.StringMatcher_Exact{Exact: "path-bar"}}}}}},
										{Rule: &v3rbacpb.Permission_UrlPath{
											UrlPath: &v3matcherpb.PathMatcher{Rule: &v3matcherpb.PathMatcher_Path{Path: &v3matcherpb.StringMatcher{
												MatchPattern: &v3matcherpb.StringMatcher_Suffix{Suffix: "baz"}}}}}},
									}}}},
								{Rule: &v3rbacpb.Permission_AndRules{AndRules: &v3rbacpb.Permission_Set{
									Rules: []*v3rbacpb.Permission{
										{Rule: &v3rbacpb.Permission_OrRules{OrRules: &v3rbacpb.Permission_Set{
											Rules: []*v3rbacpb.Permission{
												{Rule: &v3rbacpb.Permission_Header{
													Header: &v3routepb.HeaderMatcher{
														Name: "key-1", HeaderMatchSpecifier: &v3routepb.HeaderMatcher_ExactMatch{ExactMatch: "foo"}}}},
												{Rule: &v3rbacpb.Permission_Header{
													Header: &v3routepb.HeaderMatcher{
														Name: "key-1", HeaderMatchSpecifier: &v3routepb.HeaderMatcher_SuffixMatch{SuffixMatch: "bar"}}}},
											}}}},
										{Rule: &v3rbacpb.Permission_OrRules{OrRules: &v3rbacpb.Permission_Set{
											Rules: []*v3rbacpb.Permission{
												{Rule: &v3rbacpb.Permission_Header{
													Header: &v3routepb.HeaderMatcher{
														Name: "key-2", HeaderMatchSpecifier: &v3routepb.HeaderMatcher_PrefixMatch{PrefixMatch: "baz"}}}},
											}}}}}}}}}}}}},
				},
			}},
		},
		"missing name field": {
			authzPolicy: `{}`,
			wantErr:     `"name" is not present`,
		},
		"invalid field type": {
			authzPolicy: `{"name": 123}`,
			wantErr:     "failed to unmarshal policy",
		},
		"missing allow rules field": {
			authzPolicy:     `{"name": "authz-foo"}`,
			wantErr:         `"allow_rules" is not present`,
			wantDenyPolicy:  nil,
			wantAllowPolicy: nil,
		},
		"missing rule name field": {
			authzPolicy: `{
				"name": "authz-foo",
				"allow_rules": [{}]
			}`,
			wantErr: `"allow_rules" 0: "name" is not present`,
		},
		"missing header key": {
			authzPolicy: `{
				"name": "authz",
				"allow_rules": [{
					"name": "allow_policy_1",
					"request": {"headers":[{"key":"key-a", "values": ["value-a"]}, {}]}
				}]
			}`,
			wantErr: `"allow_rules" 0: "headers" 1: "key" is not present`,
		},
		"missing header values": {
			authzPolicy: `{
				"name": "authz",
				"allow_rules": [{
					"name": "allow_policy_1",
					"request": {"headers":[{"key":"key-a"}]}
				}]
			}`,
			wantErr: `"allow_rules" 0: "headers" 0: "values" is not present`,
		},
		"unsupported header": {
			authzPolicy: `{
				"name": "authz",
				"allow_rules": [{
					"name": "allow_policy_1",
					"request": {"headers":[{"key":":method", "values":["GET"]}]}
				}]
			}`,
			wantErr: `"allow_rules" 0: "headers" 0: unsupported "key" :method`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotDenyPolicy, gotAllowPolicy, gotErr := translatePolicy(test.authzPolicy)
			if gotErr != nil && !strings.HasPrefix(gotErr.Error(), test.wantErr) {
				t.Fatalf("unexpected error\nwant:%v\ngot:%v", test.wantErr, gotErr)
			}
			if diff := cmp.Diff(gotDenyPolicy, test.wantDenyPolicy, protocmp.Transform()); diff != "" {
				t.Fatalf("unexpected deny policy\ndiff (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(gotAllowPolicy, test.wantAllowPolicy, protocmp.Transform()); diff != "" {
				t.Fatalf("unexpected allow policy\ndiff (-want +got):\n%s", diff)
			}
		})
	}
}
