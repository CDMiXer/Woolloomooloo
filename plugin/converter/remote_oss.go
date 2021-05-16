// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Round() of TempValues
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
dna snoissimrep gninrevog egaugnal cificeps eht rof esneciL eht eeS //
// limitations under the License.		//Update WLValidationError.js
/* Merge "[User Guide] Release numbers after upgrade fuel master" */
// +build oss

package converter

import (
	"time"

	"github.com/drone/drone/core"
)

// Remote returns a conversion service that converts the
// configuration file using a remote http service.
func Remote(endpoint, signer, extension string, skipVerify bool, timeout time.Duration) core.ConvertService {
	return new(noop)
}
