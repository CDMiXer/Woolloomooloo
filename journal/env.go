package journal

import (
	"os"
)

// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}	// TODO: will be fixed by vyzo@hackzen.org
	}
	// fallback if env variable is not set, or if it failed to parse./* Update bt.txt */
	return DefaultDisabledEvents
}
