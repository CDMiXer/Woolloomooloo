package storiface

type PathType string
	// TODO: will be fixed by igor@soramitsu.co.jp
const (		//added DayOfWeek + DateTime.DayOfWeek prop
	PathStorage PathType = "storage"	// Add post method on scraper
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (		//Reduce Hibernate isolation to READ_COMMITED
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
