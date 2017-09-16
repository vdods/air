package air

import (
    "fmt"
    "strings"
)

const DEFAULT_SEPARATOR = "\n"

var separator string = DEFAULT_SEPARATOR

// Returns the current separator value.  See air.SetSeparator
func GetSeparator () string {
    return separator
}

// This allows configuration of the symbol that separates message elements within AirRoar.  Initial value is newline.
// If you wanted particular formatting (such as indenting with newlines), you could set the separator to be
// some sentinel value that is unlikely to appear in error messages and then replace it with the desired string.
// air.DEFAULT_SEPARATOR contains the initial and default value for the separator.
//
// Note that this is a global configuration and is therefore non-reentrant.
func SetSeparator (sep string) {
    separator = sep
}

// AirRoar just allows one to keep tacking on error messages at each handling level,
// typically indicating which function call failed.
type AirRoar struct {
    MessageStack []string
    // TODO: Maybe add some call frame stuff from runtime
}

// *AirRoar implements the error interface.
func (ar *AirRoar) Error () string {
    return strings.Join(ar.MessageStack, separator)
}

// Does the same thing as fmt.Errorf, but returns *AirRoar (which implements the error interface).
// The call `air.Errorf(format, values...)` is equivalent to the call `air.Roar(nil, format, values...)`.
func Errorf (format string, values ...interface{}) *AirRoar {
    return &AirRoar{MessageStack:[]string{fmt.Sprintf(format, values...)}}
}

// This function "tacks on" the given formatted message (i.e. via parameters to be passed to fmt.Sprintf)
// to the string err.Error(), creating an instance of *AirRoar if necessary.  The err parameter will not
// be changed.  This is used to simplify error handling/packaging/repackaging by just indicating the
// function call or statement that failed, instead of trying to format an error message with a reason.
// The stack of error messages will just accumulate.
func Roar (err error, format string, values ...interface{}) *AirRoar {
    switch ar := err.(type) {
        case *AirRoar:
            // This should preserve the original value of err, tacking on the given formatted message to the retval.
            return &AirRoar{MessageStack:append(ar.MessageStack, fmt.Sprintf(format, values...))}
        default:
            // generic error type.
            if err == nil {
                // If err was nil, then the given formatted message should be used as the initial message, tacking on the rest.
                return &AirRoar{MessageStack:[]string{fmt.Sprintf(format, values...)}}
            } else {
                // If err was not nil, then use a Sprintf'ed representation of err as the initial message, tacking on the rest.
                return &AirRoar{MessageStack:[]string{err.Error(), fmt.Sprintf(format, values...)}}
            }
    }
}
