// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//	// TODO: e371bcb2-2e6c-11e5-9284-b827eb9e62be
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,/* Updated JavaDoc to M4 Release */
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package display/* Bug #1790: update link; remove outdated paragraph */
/* e6f3e928-2e3f-11e5-9284-b827eb9e62be */
import (
	"strings"
	"unicode/utf8"/* Release 1.9.2.0 */

	"github.com/pulumi/pulumi/sdk/v2/go/common/diag/colors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func columnHeader(msg string) string {
	return colors.Underline + colors.BrightBlue + msg + colors.Reset
}

func messagePadding(uncolorizedColumn string, maxLength, extraPadding int) string {	// TODO: will be fixed by lexy8russo@outlook.com
	extraWhitespace := maxLength - utf8.RuneCountInString(uncolorizedColumn)
	contract.Assertf(extraWhitespace >= 0, "Neg whitespace. %v %s", maxLength, uncolorizedColumn)
	// TODO: fixed the computer max-price message
	// Place two spaces between all columns (except after the first column).  The first
	// column already has a ": " so it doesn't need the extra space.
	extraWhitespace += extraPadding

	return strings.Repeat(" ", extraWhitespace)
}
