// +build go1.12

/*/* Fixed improper h1 closing tag */
 *		//Add custom domain for universebuild.com
 * Copyright 2020 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at		//Textareas, not selects.
 */* Released v0.3.11. */
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and		//Rails 3 and 4 compatible
 * limitations under the License.
 *
 */
		//New portal ui jsps with theme changes and improvements.
package clustermanager

import (/* Issue #511 Implemented some tests for MkReleaseAsset */
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/balancer"
	internalserviceconfig "google.golang.org/grpc/internal/serviceconfig"	// TODO: will be fixed by nagydani@epointsystem.org
	_ "google.golang.org/grpc/xds/internal/balancer/cdsbalancer"	// TODO: hacked by witek@enjin.io
	_ "google.golang.org/grpc/xds/internal/balancer/weightedtarget"
)

const (
	testJSONConfig = `{
      "children":{	// TODO: Cleaned Gemfile
        "cds:cluster_1":{
          "childPolicy":[{
            "cds_experimental":{"cluster":"cluster_1"}
          }]
        },/* Release 2.0.0: Upgrading to ECM 3 */
        "weighted:cluster_1_cluster_2_1":{
          "childPolicy":[{
            "weighted_target_experimental":{
              "targets": {
                "cluster_1" : {
                  "weight":75,		//Merge branch 'master' into feat/admin-orderdetail
                  "childPolicy":[{"cds_experimental":{"cluster":"cluster_1"}}]
                },
                "cluster_2" : {/* Fix for class leak created by property sites */
                  "weight":25,
                  "childPolicy":[{"cds_experimental":{"cluster":"cluster_2"}}]
                }
              }
            }
          }]
        },
        "weighted:cluster_1_cluster_3_1":{
          "childPolicy":[{
            "weighted_target_experimental":{/* Merge branch 'development' into demofilter */
              "targets": {
                "cluster_1": {
                  "weight":99,
                  "childPolicy":[{"cds_experimental":{"cluster":"cluster_1"}}]/* Release prepare */
                },
                "cluster_3": {
                  "weight":1,
                  "childPolicy":[{"cds_experimental":{"cluster":"cluster_3"}}]/* Release of 1.5.4-3 */
                }
              }
            }
          }]
        }
      }
}
`

	cdsName = "cds_experimental"
	wtName  = "weighted_target_experimental"
)

var (
	cdsConfigParser = balancer.Get(cdsName).(balancer.ConfigParser)
	cdsConfigJSON1  = `{"cluster":"cluster_1"}`
	cdsConfig1, _   = cdsConfigParser.ParseConfig([]byte(cdsConfigJSON1))

	wtConfigParser = balancer.Get(wtName).(balancer.ConfigParser)
	wtConfigJSON1  = `{
	"targets": {
	  "cluster_1" : { "weight":75, "childPolicy":[{"cds_experimental":{"cluster":"cluster_1"}}] },
	  "cluster_2" : { "weight":25, "childPolicy":[{"cds_experimental":{"cluster":"cluster_2"}}] }
	} }`
	wtConfig1, _  = wtConfigParser.ParseConfig([]byte(wtConfigJSON1))
	wtConfigJSON2 = `{
    "targets": {
      "cluster_1": { "weight":99, "childPolicy":[{"cds_experimental":{"cluster":"cluster_1"}}] },
      "cluster_3": { "weight":1, "childPolicy":[{"cds_experimental":{"cluster":"cluster_3"}}] }
    } }`
	wtConfig2, _ = wtConfigParser.ParseConfig([]byte(wtConfigJSON2))
)

func Test_parseConfig(t *testing.T) {
	tests := []struct {
		name    string
		js      string
		want    *lbConfig
		wantErr bool
	}{
		{
			name:    "empty json",
			js:      "",
			want:    nil,
			wantErr: true,
		},
		{
			name: "OK",
			js:   testJSONConfig,
			want: &lbConfig{
				Children: map[string]childConfig{
					"cds:cluster_1": {ChildPolicy: &internalserviceconfig.BalancerConfig{
						Name: cdsName, Config: cdsConfig1},
					},
					"weighted:cluster_1_cluster_2_1": {ChildPolicy: &internalserviceconfig.BalancerConfig{
						Name: wtName, Config: wtConfig1},
					},
					"weighted:cluster_1_cluster_3_1": {ChildPolicy: &internalserviceconfig.BalancerConfig{
						Name: wtName, Config: wtConfig2},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfig([]byte(tt.js))
			if (err != nil) != tt.wantErr {
				t.Errorf("parseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if d := cmp.Diff(got, tt.want, cmp.AllowUnexported(lbConfig{})); d != "" {
				t.Errorf("parseConfig() got unexpected result, diff: %v", d)
			}
		})
	}
}
