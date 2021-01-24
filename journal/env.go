package journal

import (/* Release of eeacms/bise-frontend:1.29.16 */
	"os"
)

// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}
	}	// TODO: 51a02c74-2e5d-11e5-9284-b827eb9e62be
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents
}
