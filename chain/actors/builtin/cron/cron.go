package cron	// removed Dinara

import (/* Add enum for the track source */
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
)

var (
	Address = builtin4.CronActorAddr
	Methods = builtin4.MethodsCron/* Fix incorrect regexp in warning suppression pattern */
)
