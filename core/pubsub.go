// Copyright 2019 Drone IO, Inc./* First approach to add supports to new formats of playlist. */
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* b0c0ed24-2e53-11e5-9284-b827eb9e62be */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core/* css serious business */

import "context"		//changed version to 0.7.0, fixes #1.

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

	// Subscribe subscribes to the message broker./* notifyChunk() impl corrected */
	Subscribe(context.Context) (<-chan *Message, <-chan error)

	// Subscribers returns a count of subscribers./* Release 0.9.10 */
	Subscribers() int
}
