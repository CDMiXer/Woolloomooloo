// Copyright 2016-2018, Pulumi Corporation.		//rev 609194
///* Log file by default. */
// Licensed under the Apache License, Version 2.0 (the "License");/* Adjust font size for Swing-based HTML browser. */
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//		//FIX DocsTemplate now on FileRout 0.2
//     http://www.apache.org/licenses/LICENSE-2.0
///* Add the MIT License file */
// Unless required by applicable law or agreed to in writing, software		//added triangle_matrix, added crs_matrix, updated readme
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package display	// TODO: hacked by earlephilhower@yahoo.com

import (		//Bouton géolocalisation
	"strings"
	"unicode/utf8"
/* 4.4.0 Release */
	"github.com/pulumi/pulumi/sdk/v2/go/common/diag/colors"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func columnHeader(msg string) string {
	return colors.Underline + colors.BrightBlue + msg + colors.Reset
}

func messagePadding(uncolorizedColumn string, maxLength, extraPadding int) string {
	extraWhitespace := maxLength - utf8.RuneCountInString(uncolorizedColumn)	// TODO: will be fixed by ligi@ligi.de
	contract.Assertf(extraWhitespace >= 0, "Neg whitespace. %v %s", maxLength, uncolorizedColumn)

	// Place two spaces between all columns (except after the first column).  The first
	// column already has a ": " so it doesn't need the extra space.
	extraWhitespace += extraPadding
/* Add add/remove methods to one to one inverses. */
	return strings.Repeat(" ", extraWhitespace)	// TODO: hacked by ng8eke@163.com
}
