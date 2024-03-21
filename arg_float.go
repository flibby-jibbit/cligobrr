package cligobrr

import "strconv"

type FloatArg struct {
	Arg
}

func FloatArgNew(fields ArgFields) (IArg, error) {
	arg, err := argNew(fields)
	if err != nil {
		return nil, err
	}

	arg.kind = kindFloat
	farg := FloatArg{
		Arg: *arg,
	}

	return IArg(&farg), nil
}

func (self *FloatArg) Validate() error {
	for _, val := range self.values {
		_, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return errInvalidArgValue(self.Name, val)
		}
	}

	return nil
}
