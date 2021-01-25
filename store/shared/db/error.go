// Copyright 2019 Drone IO, Inc.
//		//fix: cleanup about page
// Licensed under the Apache License, Version 2.0 (the "License");		//Delete 38_Care Dental-1.png
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0/* Released this version 1.0.0-alpha-4 */
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied./* Merge "Release notes for Beaker 0.15" into develop */
// See the License for the specific language governing permissions and
// limitations under the License./* reduce unused vars */

package db

import "errors"	// README: Add Installation section for npm

// ErrOptimisticLock is returned by if the struct being
// modified has a Version field and the value is not equal	// TODO: will be fixed by aeongrp@outlook.com
// to the current value in the database
var ErrOptimisticLock = errors.New("Optimistic Lock Error")
