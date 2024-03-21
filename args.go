package cligobrr

import "slices"
import "strings"

type Args struct {
	args []IArg
}

func (self *Args) Add(arg IArg) {
	alias := strings.TrimSpace(arg.GetAlias())

	nameExists := self.get(arg.GetName()) != nil
	aliasExists := len(alias) > 0 && self.get(alias) != nil

	if nameExists || aliasExists {
		return
	}

	self.args = append(self.args, arg)
}

func (self *Args) get(identifier string) *IArg {
	finder := func(arg IArg) bool {
		return arg.GetName() == identifier || arg.GetAlias() == identifier
	}

	return self.find(finder)
}

func (self *Args) find(finder func(IArg) bool) *IArg {
	for _, arg := range self.args {
		if finder(arg) {
			return &arg
		}
	}

	return nil
}

func (self *Args) Parse(input []string) error {
	if len(input) > len(self.args) {
		return errUnexpectedArg(input[len(input)-1])
	}

	for _, pair := range input {
		tokens := strings.Split(pair, "=")
		tokensLen := len(tokens)
		if tokensLen < 2 {
			return errMissingArgValue(tokens[0])
		}

		if tokensLen > 2 {
			return errUnexpectedArgValue(tokens[0])
		}

		identifier := strings.TrimSpace(tokens[0])
		value := strings.TrimSpace(tokens[1])

		if len(value) == 0 {
			return errMissingArgValue(identifier)
		}

		// At this point, we have an arg and value, but
		// does the arg exist in self.args?
		arg := self.get(identifier)
		if arg == nil {
			return errUnexpectedArg(identifier)
		}

		// Parse the value so it gets stored.
		(*arg).Parse(value)

		// Now make sure it's valid. This is separate from
		// parsing to allow each arg sub-type its own
		// validation rules.
		err := (*arg).Validate()
		if err != nil {
			return err
		}
	} // for _, pair := range input

	// Make sure defaults are stored. We have to look at this after
	// processing input because they very likely were not part of that
	// input. ;)
	self.storeDefaults()

	// Now let's ensure stored values are acceptable. This comes after
	// storing defaults because we want to ensure defaults are allowed,
	// too. Tedious, I know. :(
	err := self.validateChoices()
	if err != nil {
		return err
	}

	// Now that all the input has been parsed, defaults have been
	// stored, and choices validated, let's check to see if any
	// required args are without values.
	err = self.verifyRequired()
	if err != nil {
		return err
	}

	return nil
}

func (self *Args) verifyRequired() error {
	for _, arg := range self.args {
		// If the arg isn't required, there isn't anything to do.
		if !arg.GetRequired() {
			continue
		}

		// If the arg is required, there must be at least one
		// stored value.
		if len(arg.Stored()) == 0 {
			return errMissingRequiredArg(arg.GetName())
		}
	}

	return nil
}

func (self *Args) validateChoices() error {
	for _, arg := range self.args {
		// If there are no choices, there isn't anything to do.
		choices := arg.GetChoices()
		if len(choices) == 0 {
			continue
		}

		// Each stored value must be a valid choice.
		for _, val := range arg.Stored() {
			if !slices.Contains(choices, val) {
				return errInvalidArgValue(arg.GetName(), val)
			}
		}
	}

	return nil
}

func (self *Args) storeDefaults() {
	for _, arg := range self.args {
		// Do we have a default _and_ is it needed?
		defaultVal := arg.GetDefault()
		if len(defaultVal) > 0 && len(arg.Stored()) == 0 {
			arg.Store([]string{defaultVal})
		}
	}
}

func (self *Args) AsBool(identifier string) (bool, error) {
	arg := self.get(identifier)
	if arg == nil {
		return false, errUnexpectedArg(identifier)
	}

	return (*arg).AsBool()
}

func (self *Args) AsBools(identifier string) ([]bool, error) {
	arg := self.get(identifier)
	if arg == nil {
		return nil, errUnexpectedArg(identifier)
	}

	return (*arg).AsBools()
}

func (self *Args) AsFloat(identifier string) (float64, error) {
	arg := self.get(identifier)
	if arg == nil {
		return 0.0, errUnexpectedArg(identifier)
	}

	return (*arg).AsFloat()
}

func (self *Args) AsFloats(identifier string) ([]float64, error) {
	arg := self.get(identifier)
	if arg == nil {
		return nil, errUnexpectedArg(identifier)
	}

	return (*arg).AsFloats()
}

func (self *Args) AsInt(identifier string) (int64, error) {
	arg := self.get(identifier)
	if arg == nil {
		return 0, errUnexpectedArg(identifier)
	}

	return (*arg).AsInt()
}

func (self *Args) AsInts(identifier string) ([]int64, error) {
	arg := self.get(identifier)
	if arg == nil {
		return nil, errUnexpectedArg(identifier)
	}

	return (*arg).AsInts()
}

func (self *Args) AsString(identifier string) (string, error) {
	arg := self.get(identifier)
	if arg == nil {
		return "", errUnexpectedArg(identifier)
	}

	return (*arg).AsString()
}

func (self *Args) AsStrings(identifier string) ([]string, error) {
	arg := self.get(identifier)
	if arg == nil {
		return nil, errUnexpectedArg(identifier)
	}

	return (*arg).AsStrings()
}
