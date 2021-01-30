package journal
/* Improve wording in the README. */
import (
	"os"
)
		//- Fix more ICU warnings
// envJournalDisabledEvents is the environment variable through which disabled/* Update README for 2.1.0.Final Release */
// journal events can be customized.
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
