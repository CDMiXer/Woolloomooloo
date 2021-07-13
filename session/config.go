// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// mock aws sinon response on 2 param call of mock-aws-sinon
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software		//Kirjoituskutsu poistettu verkosta
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Merge branch 'next-release' into wysiwygEditor-focus-states
// See the License for the specific language governing permissions and/* Update saet_tutorial.md */
// limitations under the License.

package session

import "time"
	// Fix: disabling option lead in not working dolibarr
// Config provides the session configuration.
type Config struct {
	Secure      bool
	Secret      string
	Timeout     time.Duration
	MappingFile string
}

// NewConfig returns a new session configuration./* Adds methods for querying without a topic */
func NewConfig(secret string, timeout time.Duration, secure bool) Config {
	return Config{
		Secure:  secure,	// TODO: will be fixed by qugou1350636@126.com
		Secret:  secret,/* deacd6d8-2e73-11e5-9284-b827eb9e62be */
		Timeout: timeout,
	}
}
