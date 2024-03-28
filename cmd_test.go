package cligobrr

import "testing"
import "github.com/stretchr/testify/assert"

func TestCmdNew(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name:         testCmdName,
		Alias:        testCmdAlias,
		Description:  testCmdDesc,
		Exec:         testCmdExec,
		ExecWithArgs: testCmdExecWithArgs,
	}

	cmd, err := CmdNew(fields)
	assert.Equal(testCmdName, cmd.Name)
	assert.Equal(testCmdAlias, cmd.Alias)
	assert.Equal(testCmdDesc, cmd.Description)
	assert.NotNil(cmd.Exec)
	assert.NotNil(cmd.ExecWithArgs)
	assert.Nil(err)
}

func TestCmdNewWithDefaults(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: testCmdName,
	}

	cmd, err := CmdNew(fields)
	assert.Equal(testCmdName, cmd.Name)
	assert.Empty(cmd.Alias)
	assert.Empty(cmd.Description)
	assert.Nil(cmd.Exec)
	assert.Nil(cmd.ExecWithArgs)
	assert.Nil(err)
}

func TestCmdNewRequiresName(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{}

	cmd, err := CmdNew(fields)
	assert.Nil(cmd)
	assert.NotNil(err)
}

func TestCmdsAdd(t *testing.T) {
	assert := assert.New(t)

	fields1 := CmdFields{
		Name: "cmd1",
	}

	cmd1, err := CmdNew(fields1)
	assert.NotNil(cmd1)
	assert.Nil(err)

	fields2 := CmdFields{
		Name:  testCmdName,
		Alias: testCmdAlias,
	}

	cmd2, err := CmdNew(fields2)
	assert.NotNil(cmd2)
	assert.Nil(err)

	cmd1.Cmds.Add(cmd2)
	assert.NotNil(cmd1.Cmds.get(testCmdName))
	assert.NotNil(cmd1.Cmds.get(testCmdAlias))
}

func TestCmdParseWithDefaultCommand(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: testCmdName,
	}

	cmd, err := CmdNew(fields)
	assert.NotNil(cmd)
	assert.Nil(err)

	defaultCmdFields := CmdFields{
		Name:    "defaultCmd",
		Default: true,
	}

	defaultCmd, err := CmdNew(defaultCmdFields)
	assert.NotNil(defaultCmd)
	assert.Nil(err)

	notDefaultCmdFields := CmdFields{
		Name:    "notDefaultCmd",
		Default: true,
	}

	notDefaultCmd, err := CmdNew(notDefaultCmdFields)
	assert.NotNil(notDefaultCmd)
	assert.Nil(err)

	cmd.Cmds.Add(defaultCmd)
	cmd.Cmds.Add(notDefaultCmd)

	cmdToExecute, err := cmd.Parse([]string{})
	assert.NotNil(cmdToExecute)
	assert.Nil(err)
	assert.Equal(defaultCmd.Name, cmdToExecute.Name)
}

func TestSubCmdParseWithArgs(t *testing.T) {
	assert := assert.New(t)

	cmd1Fields := CmdFields{
		Name: "cmd1",
	}

	cmd1, err := CmdNew(cmd1Fields)
	assert.NotNil(cmd1)
	assert.Nil(err)

	cmd2Fields := CmdFields{
		Name: "cmd2",
	}

	cmd2, err := CmdNew(cmd2Fields)
	assert.NotNil(cmd2)
	assert.Nil(err)

	arg1Fields := ArgFields{
		Name: "arg1",
	}

	arg1, err := BoolArgNew(arg1Fields)
	assert.NotNil(arg1)
	assert.Nil(err)

	cmd2.Args.Add(arg1)

	arg2Fields := ArgFields{
		Name: "arg2",
	}

	arg2, err := StringArgNew(arg2Fields)
	assert.NotNil(arg2)
	assert.Nil(err)

	cmd2.Args.Add(arg2)
	cmd1.Cmds.Add(cmd2)

	cmdToExec, err := cmd1.Parse([]string{"cmd2", "arg1=true", "arg2=wonky"})
	assert.Nil(err)
	assert.Equal("cmd2", cmdToExec.Name)
	assert.True(cmdToExec.Args.AsBool("arg1"))
	val2, err := cmdToExec.Args.AsString("arg2")
	assert.Nil(err)
	assert.Equal("wonky", val2)
}

func TestCmdNewAddsHelp(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: testCmdName,
	}

	cmd, err := CmdNew(fields)
	assert.NotNil(cmd)
	assert.Nil(err)

	assert.NotNil(cmd.Cmds.get("help"))
	assert.Nil(cmd.Cmds.get("h"))
}

func TestCmdNewDoesNotAddVersion(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: testCmdName,
	}

	cmd, err := CmdNew(fields)
	assert.NotNil(cmd)
	assert.Nil(err)
	assert.Nil(cmd.Cmds.get("version"))
}

func TestCmdNewDoesNotAddHelpToHelp(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: "help",
	}

	cmd, err := CmdNew(fields)
	assert.NotNil(cmd)
	assert.Nil(err)
	assert.Nil(cmd.Cmds.get("help"))
}

func TestCmdNewDoesNotAddHelpToVersion(t *testing.T) {
	assert := assert.New(t)

	fields := CmdFields{
		Name: "version",
	}

	cmd, err := CmdNew(fields)
	assert.NotNil(cmd)
	assert.Nil(err)
	assert.Nil(cmd.Cmds.get("help"))
}
