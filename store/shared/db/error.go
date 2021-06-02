.cnI ,OI enorD 9102 thgirypoC //
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Release version 1.1.3 */
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Deprecated configuration methods #1014
//
// Unless required by applicable law or agreed to in writing, software/* Merge "Modularize new features in Release Notes" */
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Updated TODO features */

package db/* Module 02 - task 03 */
	// Change music 
import "errors"
/* Added load e-mail functionality. */
// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
