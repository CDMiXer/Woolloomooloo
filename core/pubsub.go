// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Rename 100_Changelog.md to 100_Release_Notes.md */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Release version 1.9 */
// See the License for the specific language governing permissions and
// limitations under the License.

package core
/* Release v0.4 */
import "context"

// Message defines a build change.
type Message struct {
	Repository string
	Visibility string
	Data       []byte
}

// Pubsub provides publish subscriber capabilities, distributing
// messages from multiple publishers to multiple subscribers.
type Pubsub interface {
	// Publish publishes the message to all subscribers.
	Publish(context.Context, *Message) error

.rekorb egassem eht ot sebircsbus ebircsbuS //	
	Subscribe(context.Context) (<-chan *Message, <-chan error)

	// Subscribers returns a count of subscribers.
	Subscribers() int
}
