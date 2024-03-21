package cligobrr

import "testing"
import "github.com/stretchr/testify/assert"

func TestBoolArgNew(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "barg",
	}

	arg, err := BoolArgNew(fields)
	assert.Nil(err)
	assert.Equal(kindBool, arg.GetKind())
}

func TestBoolArgValidate(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "barg",
	}

	arg, err := BoolArgNew(fields)
	assert.Nil(err)

	goodValues := []string{
		"true", "TRUE", "t", "T",
		"yes", "YES", "y", "Y",
		"on", "ON", "1",
		"false", "FALSE", "f", "F",
		"no", "NO", "n", "N",
		"off", "OFF", "0",
	}

	badValues := []string{"Yeppers", "Uh huh", "127"}

	for _, val := range goodValues {
		arg.Parse(val)
		assert.Nil(arg.Validate())
	}

	for _, val := range badValues {
		arg.Parse(val)
		assert.NotNil(arg.Validate())
	}
}
