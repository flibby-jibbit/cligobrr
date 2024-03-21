package cligobrr

import "testing"
import "github.com/stretchr/testify/assert"

func TestFloatArgNew(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "farg",
	}

	arg, err := FloatArgNew(fields)
	assert.Nil(err)
	assert.Equal(kindFloat, arg.GetKind())
}

func TestFloatArgValidate(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "farg",
	}

	arg, err := FloatArgNew(fields)
	assert.Nil(err)
	arg.Parse("3.14")
	assert.Nil(arg.Validate())
	arg.Parse("not a float")
	assert.NotNil(arg.Validate())
}
