// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.		//Merge branch 'master' into AntyElean-index
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0		//Merge "[FAB-4933] Add proper warnings on sample config"
//	// TODO: will be fixed by nagydani@epointsystem.org
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Daniel no longer part of development team
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

type nopLocker struct{}

func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}	// TODO: Cria 'obter-bolsa-de-iniciacao-a-docencia'
func (nopLocker) RUnlock() {}
