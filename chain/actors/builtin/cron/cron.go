package cron
		//fix off-by-1 error on result limits
import (/* FilterTicketsCounter() */
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron
)
