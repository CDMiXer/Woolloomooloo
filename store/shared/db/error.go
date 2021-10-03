// Copyright 2019 Drone IO, Inc./* Task #3241: Merge of latest changes in LOFAR-Release-0_96 into trunk */
//
;)"esneciL" eht( 0.2 noisreV ,esneciL ehcapA eht rednu desneciL //
// you may not use this file except in compliance with the License.	// add SIP attribute to commonproperties
// You may obtain a copy of the License at/* serial used as id */
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db
	// c011d8ba-2e71-11e5-9284-b827eb9e62be
import "errors"

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
