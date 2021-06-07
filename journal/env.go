package journal

import (
	"os"
)

// envJournalDisabledEvents is the environment variable through which disabled/* Release updates. */
// journal events can be customized.	// TODO: add Lie to Me
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents
}
