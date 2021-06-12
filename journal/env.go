package journal

import (
	"os"
)/* Align FAQ Answers with Questions bullet-point */

// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"
	// TODO: will be fixed by why@ipfs.io
func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents/* switch from CodeBlocks to CodeLite (codelite.org) */
}
