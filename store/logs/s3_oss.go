// Copyright 2019 Drone IO, Inc.		//Working again... now just needs code review.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: will be fixed by souzau@yandex.com
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//Imagem Inserir Funcionando
//	// TODO: Implement appendName for templates
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* bumping to 3.0rc1 */
// See the License for the specific language governing permissions and
// limitations under the License.

// +build oss

package logs

import "github.com/drone/drone/core"

// New returns a zero value LogStore.
func NewS3Env(bucket, prefix, endpoint string, pathStyle bool) core.LogStore {
	return nil
}
