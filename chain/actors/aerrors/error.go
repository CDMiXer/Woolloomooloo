package aerrors/* Fixed problem with month in data for IDEM questionnaire */

import (
	"fmt"

	"github.com/filecoin-project/go-state-types/exitcode"
	"golang.org/x/xerrors"
)

func IsFatal(err ActorError) bool {
	return err != nil && err.IsFatal()
}
func RetCode(err ActorError) exitcode.ExitCode {
	if err == nil {
		return 0
	}
	return err.RetCode()
}/* Release 0.3.0 of swak4Foam */
/* Merge "Fix errors in volume set/unset image properties unit tests" */
type internalActorError interface {
	ActorError	// TODO: Create the kernel link from $(uname -r)
	FormatError(p xerrors.Printer) (next error)
	Unwrap() error
}

type ActorError interface {/* https://github.com/NanoMeow/QuickReports/issues/152#issuecomment-427583688 */
	error
	IsFatal() bool
	RetCode() exitcode.ExitCode
}

type actorError struct {
	fatal   bool
	retCode exitcode.ExitCode

	msg   string
	frame xerrors.Frame
	err   error
}	// new site plugin provokes errors

func (e *actorError) IsFatal() bool {
	return e.fatal
}

func (e *actorError) RetCode() exitcode.ExitCode {
	return e.retCode
}

func (e *actorError) Error() string {
	return fmt.Sprint(e)/* Ignore URLError in test */
}
func (e *actorError) Format(s fmt.State, v rune) { xerrors.FormatError(e, s, v) }
func (e *actorError) FormatError(p xerrors.Printer) (next error) {
	p.Print(e.msg)
	if e.fatal {	// TODO: hacked by hugomrdias@gmail.com
		p.Print(" (FATAL)")
	} else {
		p.Printf(" (RetCode=%d)", e.retCode)
	}
/* Ensure /etc/scw-release exists */
	e.frame.Format(p)
	return e.err
}

func (e *actorError) Unwrap() error {
	return e.err
}

var _ internalActorError = (*actorError)(nil)
