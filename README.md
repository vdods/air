# AIR ROAR

An implementation of golang's error interface which allows accumulation of error messages in order to lessen the problem of error repackaging.

# Usage

The function `air.Roar` appends a message to the given error by constructing an `air.AirRoar` structure
containing the given error and the message to append.  This can be called repeatedly, tagging each
error handling with a message indicating what was being processed when the error occurred.  Then,
when the error string is retrieved (via the error interface method `Error () string`), the messages
are joined together, separated by newlines, giving a rich description of the error and avoiding
the need to explicitly repackage error messages at each instance of error handling.

See [`air_test.go`](air_test.go) for example code (the unit tests).

# Testing

To run the unit tests, simply run `go test` from the project root.

# Notes

-   Related library that may be a useful complement: https://github.com/dagoof/failure

# To-dos

-   Perhaps add some runtime information to each added error message (such as source code location, or current function name).
-   Maybe figure out some way to call air.Roar once per function (perhaps using defer and a named error return value), since
    it's likely that the added message will be the same for each function.
-   Add more realistic test cases, showing deeply nested error messages, printing error messages nicely.

