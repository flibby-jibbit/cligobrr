package cligobrr

import "slices"
import "strings"

type BoolArg struct {
	Arg
}

func BoolArgNew(fields ArgFields) (IArg, error) {
	arg, err := argNew(fields)
	if err != nil {
		return nil, err
	}

	arg.kind = kindBool
	barg := BoolArg{
		Arg: *arg,
	}

	return IArg(&barg), nil
}

func (self *BoolArg) Validate() error {
	acceptableValues := append(truthyValues()[:], falseyValues()...)

	for _, val := range self.values {
		if !slices.Contains(acceptableValues, strings.ToLower(val)) {
			return errInvalidArgValue(self.Name, val)
		}
	}

	return nil
}
