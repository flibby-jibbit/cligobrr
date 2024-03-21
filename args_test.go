package cligobrr

import "fmt"
import "testing"
import "github.com/stretchr/testify/assert"

func TestArgParse(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(fields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "quarter",
	}

	arg, err := StringArgNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"quarter=Q1"})
	assert.Nil(err)
	stored := arg.Stored()
	assert.Equal(1, len(stored))
	assert.Equal("Q1", stored[0])
}

func TestArgParseMissingRequiredArg(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(fields)
	assert.Nil(err)

	arg1Fields := ArgFields{
		Name:     "quarter",
		Required: true,
	}

	arg1, err := StringArgNew(arg1Fields)
	assert.Nil(err)

	cmd.Args.Add(arg1)

	arg2Fields := ArgFields{
		Name:     "employees",
		Required: true,
	}

	arg2, err := IntArgNew(arg2Fields)
	assert.Nil(err)

	cmd.Args.Add(arg2)

	err = cmd.Args.Parse([]string{"employees=100"})
	assert.NotNil(err)
}

func TestArgParseInvalidChoice(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(fields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:    "quarter",
		Choices: []string{"Q1", "Q2", "Q3", "Q4"},
	}

	arg, err := StringArgNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"quarter=Q5"})
	assert.NotNil(err)
}

func TestArgsParseUnexpectedArg(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(fields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := BoolArgNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	// This tests the scenario where more args are passed
	// than have been added.
	err = cmd.Args.Parse([]string{"bool-arg=true", "any=thing"})
	assert.NotNil(err)

	// This tests the scenario where an arg is passed but
	// has not been added.
	err = cmd.Args.Parse([]string{"any=thing"})
	assert.NotNil(err)
}

func TestArgsParseMissingArgValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := BoolArgNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{argFields.Name})
	assert.NotNil(err)

	err = cmd.Args.Parse([]string{fmt.Sprintf("%s=", argFields.Name)})
	assert.NotNil(err)
}

func TestArgsParseUnexpectedArgValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := BoolArgNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{fmt.Sprintf("%s=%s=%s", argFields.Name, "a", "b")})
	assert.NotNil(err)
}

func TestArgsAsBool(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"arg=true"})
	truthy, err := cmd.Args.AsBool(arg.GetName())
	assert.True(truthy)
	assert.Nil(err)

	err = cmd.Args.Parse([]string{"arg=false"})
	truthy, err = cmd.Args.AsBool(arg.GetName())
	assert.False(truthy)
	assert.Nil(err)
}

func TestArgsAsBoolNoValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{""})
	_, err = cmd.Args.AsBool(arg.GetName())
	assert.NotNil(err)
}

func TestArgsAsBools(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"arg=true,false,true"})
	truthy, err := cmd.Args.AsBools(arg.GetName())
	assert.Equal([]bool{true, false, true}, truthy)
	assert.Nil(err)
}

func TestArgsAsBoolsNoValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{""})
	_, err = cmd.Args.AsBools(arg.GetName())
	assert.NotNil(err)
}

func TestArgsAsFloat(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"arg=1.23"})
	val, err := cmd.Args.AsFloat(arg.GetName())
	assert.Equal(1.23, val)
	assert.Nil(err)
}

func TestArgsAsFloatNoValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{""})
	_, err = cmd.Args.AsFloat(arg.GetName())
	assert.NotNil(err)
}

func TestArgsAsFloats(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"arg=1.23,4.56,7.89"})
	vals, err := cmd.Args.AsFloats(arg.GetName())
	assert.Equal([]float64{1.23, 4.56, 7.89}, vals)
	assert.Nil(err)
}

func TestArgsAsFloatsNoValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{""})
	_, err = cmd.Args.AsFloats(arg.GetName())
	assert.NotNil(err)
}

func TestArgsAsInt(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"arg=123"})
	val, err := cmd.Args.AsInt(arg.GetName())
	assert.Equal(int64(123), val)
	assert.Nil(err)
}

func TestArgsAsIntNoValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{""})
	_, err = cmd.Args.AsInt(arg.GetName())
	assert.NotNil(err)
}

func TestArgsAsInts(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"arg=123,456,789"})
	vals, err := cmd.Args.AsInts(arg.GetName())
	assert.Equal([]int64{123, 456, 789}, vals)
	assert.Nil(err)
}

func TestArgsAsIntsNoValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{""})
	_, err = cmd.Args.AsInts(arg.GetName())
	assert.NotNil(err)
}

func TestArgsAsString(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"arg=123"})
	val, err := cmd.Args.AsString(arg.GetName())
	assert.Equal("123", val)
	assert.Nil(err)
}

func TestArgsAsStringNoValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name: "arg",
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{""})
	_, err = cmd.Args.AsString(arg.GetName())
	assert.NotNil(err)
}

func TestArgsAsStrings(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{"arg=123,4.56,another"})
	vals, err := cmd.Args.AsStrings(arg.GetName())
	assert.Equal([]string{"123", "4.56", "another"}, vals)
	assert.Nil(err)
}

func TestArgsAsStringsNoValue(t *testing.T) {
	assert := assert.New(t)

	cmdFields := CmdFields{
		Name: "cmd",
	}

	cmd, err := CmdNew(cmdFields)
	assert.Nil(err)

	argFields := ArgFields{
		Name:     "arg",
		Multiple: true,
	}

	arg, err := argNew(argFields)
	assert.Nil(err)

	cmd.Args.Add(arg)

	err = cmd.Args.Parse([]string{""})
	_, err = cmd.Args.AsStrings(arg.GetName())
	assert.NotNil(err)
}
