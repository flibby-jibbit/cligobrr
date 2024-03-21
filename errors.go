package cligobrr

import "fmt"
import "errors"

func errDefaultNotAValidChoice(name string) error {
	msg := fmt.Sprintf(msgDefaultNotAValidChoice, name)
	return errors.New(msg)
}

func errInvalidArgValue(name string, value string) error {
	msg := fmt.Sprintf(msgInvalidArgValue, name, value)
	return errors.New(msg)
}

func errMissingArgValue(name string) error {
	msg := fmt.Sprintf(msgMissingArgValue, name)
	return errors.New(msg)
}

func errMissingRequiredArg(name string) error {
	msg := fmt.Sprintf(msgMissingRequiredArg, name)
	return errors.New(msg)
}

func errUnexpectedArgValue(name string) error {
	msg := fmt.Sprintf(msgUnexpectedArgValue, name)
	return errors.New(msg)
}

func errNameRequired() error {
	return errors.New(msgNameRequired)
}

func errUnexpectedArg(token string) error {
	msg := fmt.Sprintf(msgUnexpectedArg, token)
	return errors.New(msg)
}

func errUnexpectedCmd(token string) error {
	msg := fmt.Sprintf(msgUnexpectedCmd, token)
	return errors.New(msg)
}

func errArgHasNoValues(name string) error {
	msg := fmt.Sprintf(msgArgHasNoValues, name)
	return errors.New(msg)
}

func errTableColsRequired() error {
	return errors.New(msgTableColsRequired)
}

func errTableRowIncorrectCols(cols uint8) error {
	msg := fmt.Sprintf(msgTableRowIncorrectCols, cols)
	return errors.New(msg)
}
