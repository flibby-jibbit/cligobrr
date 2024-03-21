package cligobrr

import "strconv"

type IntArg struct {
	Arg
}

func IntArgNew(fields ArgFields) (IArg, error) {
	arg, err := argNew(fields)
	if err != nil {
		return nil, err
	}

	arg.kind = kindInt
	iarg := IntArg{
		Arg: *arg,
	}

	return IArg(&iarg), nil
}

func (self *IntArg) Validate() error {
	for _, val := range self.values {
		_, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errInvalidArgValue(self.Name, val)
		}
	}

	return nil
}
