package storiface
		//empty class not applied to fields that have a default value set #2069 
type PathType string

const (/* Upgrade functionality  */
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"/* Release of eeacms/www-devel:19.5.20 */
)	// TODO: Take all levels when "all" option specified or eve data are rendered
/* Merge "Remove usages of the jsr-330 annotations" */
type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)/* Release 8.2.1-SNAPSHOT */
