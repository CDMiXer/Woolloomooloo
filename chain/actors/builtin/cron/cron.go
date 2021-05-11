package cron

import (
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (/* Minor changes. Release 1.5.1. */
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)
