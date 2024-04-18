package cligobrr

import "slices"
import "strings"
import "strconv"

type IArg interface {
	GetKind() string
	GetName() string
	GetAlias() string
	GetDescription() string
	GetSeparator() string
	GetMultiple() bool
	GetRequired() bool
	GetDefault() string
	GetChoices() []string
	AsBool() (bool, error)
	AsBools() ([]bool, error)
	AsFloat() (float64, error)
	AsFloats() ([]float64, error)
	AsInt() (int64, error)
	AsInts() ([]int64, error)
	AsString() (string, error)
	AsStrings() ([]string, error)
	Validate() error
	Parse(string)
	Store([]string)
	Stored() []string
}

type ArgFields struct {
	Name        string
	Alias       string
	Description string
	Separator   string
	Multiple    bool
	Required    bool
	Default     string
	Choices     []string
}

type Arg struct {
	ArgFields
	kind   string
	values []string
}

func argNew(fields ArgFields) (*Arg, error) {
	fields.Name = strings.TrimSpace(fields.Name)

	if len(fields.Name) == 0 {
		return nil, errNameRequired()
	}

	// Make sure everything is nice and tidy.
	fields.Alias = strings.TrimSpace(fields.Alias)
	fields.Description = strings.TrimSpace(fields.Description)
	fields.Separator = strings.TrimSpace(fields.Separator)
	fields.Default = strings.TrimSpace(fields.Default)

	if len(fields.Choices) > 0 {
		var choices []string
		for _, choice := range fields.Choices {
			choice = strings.TrimSpace(choice)
			if len(choice) > 0 {
				choices = append(choices, choice)
			}
		}
		fields.Choices = choices
	}

	if fields.Multiple || len(fields.Choices) > 0 {
		if len(fields.Separator) == 0 {
			fields.Separator = separatorDefault
		}
	} else {
		fields.Separator = separatorNone
	}

	if len(fields.Default) > 0 && len(fields.Choices) > 0 {
		if !slices.Contains(fields.Choices, fields.Default) {
			return nil, errDefaultNotAValidChoice(fields.Name)
		}
	}

	arg := Arg{
		ArgFields: fields,
		kind:      kindUndefined,
	}

	return &arg, nil
}

func (self *Arg) GetKind() string {
	return self.kind
}

func (self *Arg) GetName() string {
	return self.Name
}

func (self *Arg) GetAlias() string {
	return self.Alias
}

func (self *Arg) GetDescription() string {
	return self.Description
}

func (self *Arg) GetSeparator() string {
	return self.Separator
}

func (self *Arg) GetMultiple() bool {
	return self.Multiple
}

func (self *Arg) GetRequired() bool {
	return self.Required
}

func (self *Arg) GetDefault() string {
	return self.Default
}

func (self *Arg) GetChoices() []string {
	return self.Choices
}

func (self *Arg) Store(values []string) {
	self.values = values
}

func (self *Arg) Stored() []string {
	return self.values
}

func (self *Arg) Validate() error {
	// To validate means that each stored valued is
	// checked against a rule to make sure it is valid
	// input. For the base Arg, anything is valid because
	// each sub-type must define what is and is not valid.
	return nil
}

func (self *Arg) Parse(input string) {
	var vals []string

	if self.Multiple {
		vals = strings.Split(input, self.Separator)
	} else {
		vals = []string{input}
	}

	// Make sure we start with empty values in the event
	// that Parse gets called repeatedly (like in testing).
	newValues := []string{}

	for _, val := range vals {
		val = strings.TrimSpace(val)
		if len(val) > 0 {
			newValues = append(newValues, val)
		}
	}

	self.Store(newValues)
}

func (self *Arg) AsBool() (bool, error) {
	stored := self.Stored()
	if len(stored) == 0 {
		return false, errArgHasNoValues(self.GetName())
	}

	truthy := slices.Contains(truthyValues(), stored[0])
	return truthy, nil
}

func (self *Arg) AsBools() ([]bool, error) {
	stored := self.Stored()
	if len(stored) == 0 {
		return nil, errArgHasNoValues(self.GetName())
	}

	var truthy []bool
	for _, val := range stored {
		truthy = append(truthy, slices.Contains(truthyValues(), val))
	}

	return truthy, nil
}

func (self *Arg) AsFloat() (float64, error) {
	stored := self.Stored()
	if len(stored) == 0 {
		return 0.0, errArgHasNoValues(self.GetName())
	}

	return strconv.ParseFloat(stored[0], 64)
}

func (self *Arg) AsFloats() ([]float64, error) {
	stored := self.Stored()
	if len(stored) == 0 {
		return nil, errArgHasNoValues(self.GetName())
	}

	var vals []float64
	for _, val := range stored {
		// Values have already been validated,
		// so err can be ignored.
		v, _ := strconv.ParseFloat(val, 64)
		vals = append(vals, v)
	}

	return vals, nil
}

func (self *Arg) AsInt() (int64, error) {
	stored := self.Stored()
	if len(stored) == 0 {
		return 0, errArgHasNoValues(self.GetName())
	}

	return strconv.ParseInt(stored[0], 10, 64)
}

func (self *Arg) AsInts() ([]int64, error) {
	stored := self.Stored()
	if len(stored) == 0 {
		return nil, errArgHasNoValues(self.GetName())
	}

	var vals []int64
	for _, val := range stored {
		// Values have already been validated,
		// so err can be ignored.
		v, _ := strconv.ParseInt(val, 10, 64)
		vals = append(vals, v)
	}

	return vals, nil
}

func (self *Arg) AsString() (string, error) {
	stored := self.Stored()
	if len(stored) == 0 {
		return "", errArgHasNoValues(self.GetName())
	}

	return stored[0], nil
}

func (self *Arg) AsStrings() ([]string, error) {
	stored := self.Stored()
	if len(stored) == 0 {
		return nil, errArgHasNoValues(self.GetName())
	}

	return stored, nil
}
