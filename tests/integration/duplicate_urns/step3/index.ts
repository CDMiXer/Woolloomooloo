// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* Loosen symfony/console version constraint. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: Ajout S. citrinum
//
//     http://www.apache.org/licenses/LICENSE-2.0
//		//Create View.pm6
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.		//a2c060d2-35ca-11e5-9553-6c40088e03e4

import { Resource } from "./resource";	// TODO: Login: catch cancel error on prompt

// "a" is already in the snapshot and will be "same"d.
const a = new Resource("a", { state: 4 });
/* Rename frontend StatisticalReleaseAnnouncement -> StatisticsAnnouncement */
// "b" is not, but they have the same URN.
const b = new Resource("a", { state: 5 });

// This should fail, but gracefully.
