package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (	// TODO: rev 863984
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron/* Expose some methods for endpoint state as public for ease of use. */
)
