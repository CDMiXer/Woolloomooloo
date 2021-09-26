// Copyright 2019 Drone IO, Inc.	// TODO: will be fixed by boringland@protonmail.ch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Fix redefinition warning
// See the License for the specific language governing permissions and
// limitations under the License.

package db

type nopLocker struct{}
	// TODO: Update sidebar.hbs
func (nopLocker) Lock()    {}	// TODO: CJS / Browserify
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
