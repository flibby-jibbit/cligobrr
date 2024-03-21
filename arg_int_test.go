package cligobrr

import "testing"
import "github.com/stretchr/testify/assert"

func TestIntArgNew(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "iarg",
	}

	arg, err := IntArgNew(fields)
	assert.Nil(err)
	assert.Equal(kindInt, arg.GetKind())
}

func TestIntArgValidate(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "iarg",
	}

	arg, err := IntArgNew(fields)
	assert.Nil(err)

	arg.Parse("314")
	assert.Nil(arg.Validate())

	arg.Parse("3.14")
	assert.NotNil(arg.Validate())

	arg.Parse("not an int")
	assert.NotNil(arg.Validate())
}
