package cligobrr

import "testing"
import "github.com/stretchr/testify/assert"

func TestArgNew(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:        "arg",
		Alias:       "a",
		Description: "An undefined arg.",
		Multiple:    true,
		Required:    true,
		Default:     "two",
		Choices:     []string{"one", "two", "three"},
	}

	arg, err := argNew(fields)
	assert.Nil(err)
	assert.Equal(kindUndefined, arg.GetKind())
	assert.Equal(fields.Name, arg.GetName())
	assert.Equal(fields.Alias, arg.GetAlias())
	assert.Equal(fields.Description, arg.GetDescription())
	assert.Equal(separatorDefault, arg.GetSeparator())
	assert.Equal(fields.Multiple, arg.GetMultiple())
	assert.Equal(fields.Required, arg.GetRequired())
	assert.Equal(fields.Default, arg.GetDefault())
	assert.Equal(fields.Choices, arg.GetChoices())
}

func TestArgNewCleanup(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:        "arg  ",
		Alias:       "  a",
		Description: " An undefined arg. ",
		Separator:   "             ^                ",
		Multiple:    true,
		Required:    true,
		Default:     "two ",
		Choices:     []string{" one", "two ", "  three      "},
	}

	arg, err := argNew(fields)
	assert.Nil(err)
	assert.Equal("arg", arg.GetName())
	assert.Equal("a", arg.GetAlias())
	assert.Equal("An undefined arg.", arg.GetDescription())
	assert.Equal("^", arg.GetSeparator())
	assert.Equal("two", arg.GetDefault())
	assert.Equal([]string{"one", "two", "three"}, arg.GetChoices())
}

func TestArgNewNoName(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{}

	arg, err := argNew(fields)
	assert.Nil(arg)
	assert.NotNil(err)
}

func TestArgNewDefaultNotAValidChoice(t *testing.T) {
	assert := assert.New(t)

	// All good scenario.
	fields := ArgFields{
		Name:    "arg",
		Default: "two",
		Choices: []string{"one", "two", "three"},
	}

	_, err := argNew(fields)

	assert.Nil(err)

	// No Default scenario.
	fields.Default = ""
	_, err = argNew(fields)
	assert.Nil(err)

	// No Choices scenario.
	fields.Default = "two"
	fields.Choices = []string{}
	_, err = argNew(fields)
	assert.Nil(err)

	// Failure scenario.
	fields.Default = "two"
	fields.Choices = []string{"four", "five", "six"}
	_, err = argNew(fields)
	assert.NotNil(err)
}

func TestArgParseSingle(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: false,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("one,two,three")

	stored := arg.Stored()
	assert.Equal(1, len(stored))
	assert.Equal("one,two,three", stored[0])
}

func TestArgParseMultiple(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("one,two,three")

	stored := arg.Stored()
	assert.Equal(3, len(stored))
	assert.Equal("one", stored[0])
	assert.Equal("two", stored[1])
	assert.Equal("three", stored[2])
}

func TestArgParseMultipleDifferentSeparator(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:      "arg",
		Separator: "@",
		Multiple:  true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("one@two@three")

	stored := arg.Stored()
	assert.Equal(3, len(stored))
	assert.Equal("one", stored[0])
	assert.Equal("two", stored[1])
	assert.Equal("three", stored[2])
}

func TestArgParseEmpty(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("")
	stored := arg.Stored()
	assert.Equal(0, len(stored))
}

func TestArgAsBool(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("true")
	truthy, err := arg.AsBool()
	assert.True(truthy)
	assert.Nil(err)

	arg.Parse("false")
	truthy, err = arg.AsBool()
	assert.False(truthy)
	assert.Nil(err)
}

func TestArgAsBoolNoValue(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("")
	_, err = arg.AsBool()
	assert.NotNil(err)
}

func TestArgAsBools(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("true,false,true")
	truthy, err := arg.AsBools()
	assert.Equal([]bool{true, false, true}, truthy)
	assert.Nil(err)
}

func TestArgAsBoolsNoValue(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse(",,")
	_, err = arg.AsBools()
	assert.NotNil(err)
}

func TestArgAsFloat(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("1.23")
	val, err := arg.AsFloat()
	assert.Equal(1.23, val)
	assert.Nil(err)
}

func TestArgAsFloatNoValue(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("")
	_, err = arg.AsFloat()
	assert.NotNil(err)
}

func TestArgAsFloats(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("1.23,4.56,7.89")
	vals, err := arg.AsFloats()
	assert.Equal([]float64{1.23, 4.56, 7.89}, vals)
	assert.Nil(err)
}

func TestArgAsFloatsNoValue(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse(",,")
	_, err = arg.AsFloats()
	assert.NotNil(err)
}

func TestArgAsInt(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("123")
	val, err := arg.AsInt()
	assert.Equal(int64(123), val)
	assert.Nil(err)
}

func TestArgAsIntNoValue(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("")
	_, err = arg.AsInt()
	assert.NotNil(err)
}

func TestArgAsInts(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("123,456,789")
	vals, err := arg.AsInts()
	assert.Equal([]int64{123, 456, 789}, vals)
	assert.Nil(err)
}

func TestArgAsIntsNoValue(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse(",,")
	_, err = arg.AsInts()
	assert.NotNil(err)
}

func TestArgAsString(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("123")
	val, err := arg.AsString()
	assert.Equal("123", val)
	assert.Nil(err)
}

func TestArgAsStringNoValue(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("")
	_, err = arg.AsString()
	assert.NotNil(err)
}

func TestArgAsStrings(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse("123,4.56,another")
	vals, err := arg.AsStrings()
	assert.Equal([]string{"123", "4.56", "another"}, vals)
	assert.Nil(err)
}

func TestArgAsStringsNoValue(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(fields)
	assert.Nil(err)

	arg.Parse(",,")
	_, err = arg.AsStrings()
	assert.NotNil(err)
}
