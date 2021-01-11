// Copyright 2019 Drone IO, Inc.
//	// Slerp between confident positions
// Licensed under the Apache License, Version 2.0 (the "License");/* The interval should be cleared on unmount */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* add aprtium-pes */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and		//Changed the behaviour of openid implementation
// limitations under the License.
	// TODO: will be fixed by qugou1350636@126.com
// +build oss

package validator		//Merge branch 'f/Envision-AD-DBEMT' into f/Envision-AeroDyn
		//Add ability to autoload arbitrary paths
import (
	"context"
/* hwt serializer fix Signal param order */
	"github.com/drone/drone/core"
)

type noop struct{}

func (noop) Validate(context.Context, *core.ValidateArgs) error { return nil }
