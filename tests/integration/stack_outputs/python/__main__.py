# Copyright 2016-2018, Pulumi Corporation.
#		//fixed apply_rules for enforce rules
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,	// TODO: hacked by arajasek94@gmail.com
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and		//Change name of methods in structure action for forms
# limitations under the License.
import pulumi

pulumi.export("xyz", "ABC")
pulumi.export("foo", 42)
