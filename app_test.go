package cligobrr

import "fmt"
import "testing"
import "github.com/stretchr/testify/assert"

func TestAppNew(t *testing.T) {
	assert := assert.New(t)

	fields := AppFields{
		Name:        testAppName,
		Description: testAppDesc,
		Version:     testAppVersion,
	}

	app := AppNew(fields)
	assert.Equal(testAppName, app.Name)
	assert.Equal(testAppDesc, app.Description)
	assert.Equal(testAppVersion, app.Version)
}

func TestAppNewWithDefaults(t *testing.T) {
	assert := assert.New(t)

	fields := AppFields{
		Name: testAppName,
	}

	app := AppNew(fields)
	assert.Equal(testAppName, app.Name)
	assert.Empty(app.Description)
	assert.Empty(app.Version)
}

func TestAppNewAddsHelp(t *testing.T) {
	assert := assert.New(t)

	fields := AppFields{
		Name: testAppName,
	}

	app := AppNew(fields)
	assert.NotNil(app.Cmds.get("help"))
	assert.Nil(app.Cmds.get("h"))
}

func TestAppNewAddsVersion(t *testing.T) {
	assert := assert.New(t)

	fields := AppFields{
		Name: testAppName,
	}

	app := AppNew(fields)
	assert.NotNil(app.Cmds.get("version"))
	assert.Nil(app.Cmds.get("v"))
}

func TestAppParseArg(t *testing.T) {
	assert := assert.New(t)

	appFields := AppFields{
		Name: testAppName,
	}

	app := AppNew(appFields)

	argFields := ArgFields{
		Name:        testArgName,
		Alias:       testArgAlias,
		Description: testArgDesc,
	}

	arg, err := StringArgNew(argFields)
	assert.NotNil(arg)
	assert.Nil(err)

	app.Args.Add(arg)

	// Parse by name.
	argFragment := fmt.Sprintf("%s=val", arg.GetName())
	cmdToExec, err := app.Parse([]string{testAppName, argFragment})
	assert.NotNil(cmdToExec)
	assert.Nil(err)
	assert.Equal("help", cmdToExec.Name)

	argVal, err := app.Args.AsString(arg.GetName())
	assert.Nil(err)
	assert.Equal("val", argVal)

	// Parse by alias.
	argFragment = fmt.Sprintf("%s=val", arg.GetAlias())
	cmdToExec, err = app.Parse([]string{testAppName, argFragment})
	assert.NotNil(cmdToExec)
	assert.Nil(err)
	assert.Equal("help", cmdToExec.Name)

	argVal, err = app.Args.AsString(arg.GetAlias())
	assert.Nil(err)
	assert.Equal("val", argVal)
}

func TestAppParseCommand(t *testing.T) {
	assert := assert.New(t)

	appFields := AppFields{
		Name: testAppName,
	}

	app := AppNew(appFields)

	cmdFields := CmdFields{
		Name:         testCmdName,
		Alias:        testCmdAlias,
		Exec:         testCmdExec,
		ExecWithArgs: testCmdExecWithArgs,
	}

	cmd, err := CmdNew(cmdFields)
	assert.NotNil(cmd)
	assert.Nil(err)

	app.Cmds.Add(cmd)

	// Parse by name.
	cmdToExec, err := app.Parse([]string{testAppName, cmd.Name})
	assert.NotNil(cmdToExec)
	assert.Nil(err)
	assert.Equal(cmd.Name, cmdToExec.Name)

	// Parse by alias.
	cmdToExec, err = app.Parse([]string{testAppName, cmd.Alias})
	assert.NotNil(cmdToExec)
	assert.Nil(err)
	assert.Equal(cmd.Name, cmdToExec.Name)
}

func TestAppParseCommandDoesNotExist(t *testing.T) {
	assert := assert.New(t)

	fields := AppFields{
		Name: testAppName,
	}

	app := AppNew(fields)
	cmdToExec, err := app.Parse([]string{testAppName, "does-not-exist"})
	assert.Nil(cmdToExec)
	assert.NotNil(err)
}

func TestAppParseNoArgsWithDefaultCommand(t *testing.T) {
	assert := assert.New(t)

	appFields := AppFields{
		Name: testAppName,
	}

	app := AppNew(appFields)

	cmdFields := CmdFields{
		Name:    testCmdName,
		Default: true,
	}

	cmd, err := CmdNew(cmdFields)
	assert.NotNil(cmd)
	assert.Nil(err)

	app.Cmds.Add(cmd)

	cmdToExec, err := app.Parse([]string{testAppName})
	assert.NotNil(cmdToExec)
	assert.Nil(err)
	assert.Equal(cmd.Name, cmdToExec.Name)
}
