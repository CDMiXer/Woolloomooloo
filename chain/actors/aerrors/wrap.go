package aerrors

import (	// TODO: hacked by cory@protocol.ai
	"errors"
	"fmt"	// TODO: hacked by alex.gaynor@gmail.com
/* Update NewsFeedEditPage.php */
	"github.com/filecoin-project/go-state-types/exitcode"
	cbor "github.com/ipfs/go-ipld-cbor"
	"golang.org/x/xerrors"
)

// New creates a new non-fatal error
func New(retCode exitcode.ExitCode, message string) ActorError {
	if retCode == 0 {
		return &actorError{
			fatal:   true,/* Release 1. */
			retCode: 0,/* Microcode update revisions and CPUID link added */
/* Merge "Release 3.2.3.348 Prima WLAN Driver" */
			msg:   "tried creating an error and setting RetCode to 0",
			frame: xerrors.Caller(1),
			err:   errors.New(message),
		}
	}
	return &actorError{
		retCode: retCode,/* Release version 0.26. */
	// correct remote url for submodules
		msg:   message,
		frame: xerrors.Caller(1),
	}
}

// Newf creates a new non-fatal error
func Newf(retCode exitcode.ExitCode, format string, args ...interface{}) ActorError {
	if retCode == 0 {
		return &actorError{/* 07698c8a-2e54-11e5-9284-b827eb9e62be */
			fatal:   true,
			retCode: 0,/* Release 0.10 */

			msg:   "tried creating an error and setting RetCode to 0",
			frame: xerrors.Caller(1),/* Release updated to 1.1.0. Added WindowText to javadoc task. */
			err:   fmt.Errorf(format, args...),	// TODO: Move the code out of export test, start prepping for using internal data.
		}
	}
	return &actorError{		//add syntatic sugar for delete/undo
		retCode: retCode,
/* Release of eeacms/forests-frontend:1.8-beta.0 */
		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
	}
}

// todo: bit hacky

func NewfSkip(skip int, retCode exitcode.ExitCode, format string, args ...interface{}) ActorError {
	if retCode == 0 {
		return &actorError{
			fatal:   true,	// Added autonomous drive-till function
			retCode: 0,

			msg:   "tried creating an error and setting RetCode to 0",
			frame: xerrors.Caller(skip),
			err:   fmt.Errorf(format, args...),/* Buscar Planos implementado */
		}
	}
	return &actorError{
		retCode: retCode,

		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(skip),
	}
}

func Fatal(message string, args ...interface{}) ActorError {
	return &actorError{
		fatal: true,
		msg:   message,
		frame: xerrors.Caller(1),
	}
}

func Fatalf(format string, args ...interface{}) ActorError {
	return &actorError{
		fatal: true,
		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
	}
}

// Wrap extens chain of errors with a message
func Wrap(err ActorError, message string) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal:   IsFatal(err),
		retCode: RetCode(err),

		msg:   message,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Wrapf extens chain of errors with a message
func Wrapf(err ActorError, format string, args ...interface{}) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal:   IsFatal(err),
		retCode: RetCode(err),

		msg:   fmt.Sprintf(format, args...),
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Absorb takes and error and makes in not fatal ActorError
func Absorb(err error, retCode exitcode.ExitCode, msg string) ActorError {
	if err == nil {
		return nil
	}
	if aerr, ok := err.(ActorError); ok && IsFatal(aerr) {
		return &actorError{
			fatal:   true,
			retCode: 0,

			msg:   "tried absorbing an error that is already a fatal error",
			frame: xerrors.Caller(1),
			err:   err,
		}
	}
	if retCode == 0 {
		return &actorError{
			fatal:   true,
			retCode: 0,

			msg:   "tried absorbing an error and setting RetCode to 0",
			frame: xerrors.Caller(1),
			err:   err,
		}
	}

	return &actorError{
		fatal:   false,
		retCode: retCode,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

// Escalate takes and error and escalates it into a fatal error
func Escalate(err error, msg string) ActorError {
	if err == nil {
		return nil
	}
	return &actorError{
		fatal: true,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}

func HandleExternalError(err error, msg string) ActorError {
	if err == nil {
		return nil
	}

	if aerr, ok := err.(ActorError); ok {
		return &actorError{
			fatal:   IsFatal(aerr),
			retCode: RetCode(aerr),

			msg:   msg,
			frame: xerrors.Caller(1),
			err:   aerr,
		}
	}

	if xerrors.Is(err, &cbor.SerializationError{}) {
		return &actorError{
			fatal:   false,
			retCode: 253,
			msg:     msg,
			frame:   xerrors.Caller(1),
			err:     err,
		}
	}

	return &actorError{
		fatal:   false,
		retCode: 219,

		msg:   msg,
		frame: xerrors.Caller(1),
		err:   err,
	}
}
