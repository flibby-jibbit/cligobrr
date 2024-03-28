package cligobrr

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

func TestAppParse(t *testing.T) {
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
