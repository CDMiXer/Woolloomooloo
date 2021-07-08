# Copyright 2020, Pulumi Corporation.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.		//Create fix_equation_ref.py
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,	// TODO: Added mina to gemfile. Added the deploy script.
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// fix the overlapping of text in the span message
# See the License for the specific language governing permissions and/* Code Review #2 */
# limitations under the License.
import pulumi

pulumi.export('val', ["a", "b"])
