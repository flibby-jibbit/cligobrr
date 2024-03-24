package cligobrr

const (
	// Kinds
	kindBool      = "bool"
	kindFloat     = "float"
	kindInt       = "int"
	kindString    = "string"
	kindUndefined = "undefined"

	// Separators
	separatorDefault = ","
	separatorNone    = ""

	// Errors
	msgArgHasNoValues         = "Argument has no values: %s."
	msgInvalidArgValue        = "Invalid argument value: %s=%s."
	msgMissingArgValue        = "Missing argument value: %s."
	msgMissingRequiredArg     = "Required argument missing: %s."
	msgUnexpectedArgValue     = "Unexpected argument value: %s."
	msgNameRequired           = "Name is required."
	msgUnexpectedArg          = "Unexpected argument: %s."
	msgUnexpectedCmd          = "Unexpected command: %s."
	msgDefaultNotAValidChoice = "Default value is not a valid choice: %s."
	msgTableColsRequired      = "Table columns is required."
	msgTableRowIncorrectCols  = "Table row must contain %d columns."

	// Tables
	tablePadDefault = uint8(4)

	// Tests
	testAppName    string = "myApp"
	testAppDesc    string = "My App"
	testAppVersion string = "0.1.0"
	testCmdName    string = "myCmd"
	testCmdAlias   string = "mc"
	testCmdDesc    string = "My Command"
)

// Golang doesn't allow complex types as constants, so this is
// essentially our only alternative. :(
func truthyValues() []string {
	return []string{"true", "t", "yes", "y", "on", "1"}
}

func falseyValues() []string {
	return []string{"false", "f", "no", "n", "off", "0"}
}

// For tests.
func testCmdExec(args Args) {
}
