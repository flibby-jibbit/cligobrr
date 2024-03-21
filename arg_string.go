package cligobrr

type StringArg struct {
	Arg
}

func StringArgNew(fields ArgFields) (IArg, error) {
	arg, err := argNew(fields)

	if err != nil {
		return nil, err
	}

	arg.kind = kindString
	sarg := StringArg{
		Arg: *arg,
	}

	return IArg(&sarg), nil
}
