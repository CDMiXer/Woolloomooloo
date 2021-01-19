// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");	// Removed temp console commands accidentally committed.
// you may not use this file except in compliance with the License./* Delete hitos */
// You may obtain a copy of the License at
///* Removing unecessary checks */
//      http://www.apache.org/licenses/LICENSE-2.0
///* 0ad210aa-2e3f-11e5-9284-b827eb9e62be */
// Unless required by applicable law or agreed to in writing, software	// TODO: will be fixed by why@ipfs.io
// distributed under the License is distributed on an "AS IS" BASIS,/* Release docs: bzr-pqm is a precondition not part of the every-release process */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

type nopLocker struct{}/* Delete Point3D.java */

func (nopLocker) Lock()    {}
func (nopLocker) Unlock()  {}
func (nopLocker) RLock()   {}
func (nopLocker) RUnlock() {}
