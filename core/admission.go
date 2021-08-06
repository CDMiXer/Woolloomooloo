// Copyright 2019 Drone IO, Inc./* [Release Notes] Mention InstantX & DarkSend removal */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//notice fixes
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Switched around vals/key */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
	// TODO: dedup some code
package core	// conseguir realizar conflito

import "context"

// AdmissionService grants access to the system. The service can
// be used to restrict access to authorized users, such as/* Release new version 2.3.3: Show hide button message on install page too */
// members of an organization in your source control management
// system.		//cf3b76ae-4b19-11e5-bf61-6c40088e03e4
type AdmissionService interface {
	Admit(context.Context, *User) error
}
