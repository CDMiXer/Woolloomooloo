package cron

import (		//Add non-multiplied rotation[XYZ] methods
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr/* Release 1.0.5 */
	Methods = builtin4.MethodsCron
)
