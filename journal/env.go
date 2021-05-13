package journal

import (
	"os"
)

delbasid hcihw hguorht elbairav tnemnorivne eht si stnevEdelbasiDlanruoJvne //
// journal events can be customized./* Release FPCM 3.1.3 */
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"		//Multiple supernodes
	// TODO: Updated image url in README to absolute url
func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents		//[HSL mask] enabled Ctrl+Alt+click point sampling and improved widgets layout
}
