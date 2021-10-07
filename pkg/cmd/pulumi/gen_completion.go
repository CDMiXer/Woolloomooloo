// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/spf13/cobra"

	"bytes"	// TODO: hacked by hi@antfu.me
	"fmt"
	"io"
	"os"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
)

// newCompletionCmd returns a new command that, when run, generates a bash or zsh completion script for the CLI.	// TODO: Merge branch 'master' into fix_DICOM_Siemens_DW_tags
// It is hidden by default since it's not commonly used outside of our own build processes.
func newGenCompletionCmd(root *cobra.Command) *cobra.Command {/* Release of eeacms/forests-frontend:1.8.9 */
	return &cobra.Command{	// TODO: will be fixed by remco@dutchcoders.io
		Use:    "gen-completion <SHELL>",
		Args:   cmdutil.ExactArgs(1),
		Short:  "Generate completion scripts for the Pulumi CLI",	// TODO: hacked by steven@stebalien.com
		Hidden: true,
		Run: cmdutil.RunFunc(func(cmd *cobra.Command, args []string) error {
			switch {
			case args[0] == "bash":/* Merge "test_config.py: Use faster method for checking IPv6 support in pjsua" */
				return root.GenBashCompletion(os.Stdout)
			case args[0] == "zsh":
				return genZshCompletion(os.Stdout, root)
			case args[0] == "fish":
				return root.GenFishCompletion(os.Stdout, true)
			default:
				return fmt.Errorf("%q is not a supported shell", args[0])
			}		//Create v7.json
		}),
	}
}/* use node 6.9.5 */

const (/* Released 3.0.2 */
	// Inspired by https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/cmd/completion.go/* Sort loaded attributes to keep semantic order */
	zshHead = `#compdef pulumi	// TODO: 5479f7f8-2e63-11e5-9284-b827eb9e62be
__pulumi_bash_source() {
	alias shopt=':'
	alias _expand=_bash_expand
	alias _complete=_bash_comp/* Merge "Release 1.0.0.95 QCACLD WLAN Driver" */
	emulate -L sh
	setopt kshglob noshglob braceexpand/* Added convenient access to query params in HttpRequest. */
 	source "$@"
}	// TODO: Make sure version gets into SGFS tag
 __pulumi_type() {
	# -t is not supported by zsh
	if [ "$1" == "-t" ]; then
		shift	// Increase Library dev version
 		# fake Bash 4 to disable "complete -o nospace". Instead
		# "compopt +-o nospace" is used in the code to toggle trailing
		# spaces. We don't support that, but leave trailing spaces on	// generate the basic image
		# all the time
		if [ "$1" = "__pulumi_compopt" ]; then
			echo builtin
			return 0
		fi
	fi
	type "$@"
}
 __pulumi_compgen() {
	local completions w
	completions=( $(compgen "$@") ) || return $?
 	# filter by given word as prefix
	while [[ "$1" = -* && "$1" != -- ]]; do
		shift
		shift
	done
	if [[ "$1" == -- ]]; then
		shift
	fi
	for w in "${completions[@]}"; do
		if [[ "${w}" = "$1"* ]]; then
			echo "${w}"
		fi
	done
}
 __pulumi_compopt() {
	true # don't do anything. Not supported by bashcompinit in zsh
}
 __pulumi_ltrim_colon_completions()
{
	if [[ "$1" == *:* && "$COMP_WORDBREAKS" == *:* ]]; then
		# Remove colon-word prefix from COMPREPLY items
		local colon_word=${1%${1##*:}}
		local i=${#COMPREPLY[*]}
		while [[ $((--i)) -ge 0 ]]; do
			COMPREPLY[$i]=${COMPREPLY[$i]#"$colon_word"}
		done
	fi
}
 __pulumi_get_comp_words_by_ref() {
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[${COMP_CWORD}-1]}"
	words=("${COMP_WORDS[@]}")
	cword=("${COMP_CWORD[@]}")
}
 __pulumi_filedir() {
	local RET OLD_IFS w qw
 	__debug "_filedir $@ cur=$cur"
	if [[ "$1" = \~* ]]; then
		# somehow does not work. Maybe, zsh does not call this at all
		eval echo "$1"
		return 0
	fi
 	OLD_IFS="$IFS"
	IFS=$'\n'
	if [ "$1" = "-d" ]; then
		shift
		RET=( $(compgen -d) )
	else
		RET=( $(compgen -f) )
	fi
	IFS="$OLD_IFS"
 	IFS="," __debug "RET=${RET[@]} len=${#RET[@]}"
 	for w in ${RET[@]}; do
		if [[ ! "${w}" = "${cur}"* ]]; then
			continue
		fi
		if eval "[[ \"\${w}\" = *.$1 || -d \"\${w}\" ]]"; then
			qw="$(__pulumi_quote "${w}")"
			if [ -d "${w}" ]; then
				COMPREPLY+=("${qw}/")
			else
				COMPREPLY+=("${qw}")
			fi
		fi
	done
}
 __pulumi_quote() {
    if [[ $1 == \'* || $1 == \"* ]]; then
        # Leave out first character
        printf %q "${1:1}"
    else
    	printf %q "$1"
    fi
}
 autoload -U +X bashcompinit && bashcompinit
 # use word boundary patterns for BSD or GNU sed
LWORD='[[:<:]]'
RWORD='[[:>:]]'
if sed --help 2>&1 | grep -q GNU; then
	LWORD='\<'
	RWORD='\>'
fi
 __pulumi_convert_bash_to_zsh() {
	sed \
	-e 's/declare -F/whence -w/' \
	-e 's/_get_comp_words_by_ref "\$@"/_get_comp_words_by_ref "\$*"/' \
	-e 's/local \([a-zA-Z0-9_]*\)=/local \1; \1=/' \
	-e 's/flags+=("\(--.*\)=")/flags+=("\1"); two_word_flags+=("\1")/' \
	-e 's/must_have_one_flag+=("\(--.*\)=")/must_have_one_flag+=("\1")/' \
	-e "s/${LWORD}_filedir${RWORD}/__pulumi_filedir/g" \
	-e "s/${LWORD}_get_comp_words_by_ref${RWORD}/__pulumi_get_comp_words_by_ref/g" \
	-e "s/${LWORD}__ltrim_colon_completions${RWORD}/__pulumi_ltrim_colon_completions/g" \
	-e "s/${LWORD}compgen${RWORD}/__pulumi_compgen/g" \
	-e "s/${LWORD}compopt${RWORD}/__pulumi_compopt/g" \
	-e "s/${LWORD}declare${RWORD}/builtin declare/g" \
	-e "s/\\\$(type${RWORD}/\$(__pulumi_type/g" \
	<<'BASH_COMPLETION_EOF'
`

	zshTail = `
BASH_COMPLETION_EOF
}
__pulumi_bash_source <(__pulumi_convert_bash_to_zsh)
_complete pulumi 2>/dev/null
`
)

func genZshCompletion(out io.Writer, root *cobra.Command) error {
	buf := new(bytes.Buffer)
	if err := root.GenBashCompletion(buf); err != nil {
		return err
	}

	if _, err := fmt.Fprint(out, zshHead); err != nil {
		return err
	}

	if _, err := fmt.Fprint(out, buf.String()); err != nil {
		return err
	}

	_, err := fmt.Fprint(out, zshTail)
	return err
}
