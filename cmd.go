package cligobrr

import "fmt"
import "strings"

type FuncCmdExec func(args Args)

type CmdFields struct {
	Name        string
	Alias       string
	Description string
	Default     bool
	Exec        FuncCmdExec
}

type Cmd struct {
	CmdFields
	Cmds Cmds
	Args Args
}

type Cmds struct {
	cmds []Cmd
}

func CmdNew(fields CmdFields) (*Cmd, error) {
	if len(strings.TrimSpace(fields.Name)) == 0 {
		return nil, errNameRequired()
	}

	cmd := Cmd{
		CmdFields: fields,
	}

	// We don't want to add 'help' to the 'help' and
	// 'version' commands that get added automatically.
	if fields.Name != "help" && fields.Name != "version" {
		helpFields := CmdFields{
			Name:        "help",
			Description: fmt.Sprintf("Display help for %s.", cmd.Name),
		}

		helpCmd, _ := CmdNew(helpFields)
		cmd.Cmds.Add(helpCmd)
	}

	return &cmd, nil
}

func (self *Cmds) Add(cmd *Cmd) {
	alias := strings.TrimSpace(cmd.Alias)

	nameExists := self.get(cmd.Name) != nil
	aliasExists := len(alias) > 0 && self.get(alias) != nil

	if nameExists || aliasExists {
		return
	}

	// If any other command is already default, this one
	// can't be. The first one wins.
	defCmd := self.defaultCmd()
	if cmd.Default && defCmd != nil {
		cmd.Default = false
	}

	self.cmds = append(self.cmds, *cmd)
}

func (self *Cmds) get(identifier string) *Cmd {
	finder := func(cmd Cmd) bool {
		return cmd.Name == identifier || cmd.Alias == identifier
	}

	return self.find(finder)
}

func (self *Cmds) defaultCmd() *Cmd {
	finder := func(cmd Cmd) bool { return cmd.Default }
	return self.find(finder)
}

func (self *Cmds) find(finder func(Cmd) bool) *Cmd {
	for _, cmd := range self.cmds {
		if finder(cmd) {
			return &cmd
		}
	}

	return nil
}

func (self *Cmd) Parse(args []string) (*Cmd, error) {
	if len(args) > 0 {
		token := args[0]

		cmd := self.Cmds.get(token)
		if cmd != nil {
			args = args[1:]

			if cmd.Name == "help" {
				err := cmdHelp(self, args)
				if err != nil {
					return nil, err
				} else {
					return cmd, nil
				}
			}

			cmd, err := cmd.Parse(args)
			if err != nil {
				return nil, err
			} else {
				return cmd, nil
			}
		} else {
			// Didn't find a command, so assume
			// input is args for self.
			err := self.Args.Parse(args)
			if err != nil {
				return nil, err
			}

			return self, nil
		}
	}

	if self.Exec == nil {
		// If there is nothing to execute, we need to see if
		// there is a default command.

		defCmd := self.Cmds.defaultCmd()
		if defCmd != nil {
			return defCmd, nil
		}

		// No default command, so let's display help.
		err := cmdHelp(self, args)
		if err != nil {
			return nil, err
		} else {
			return self.Cmds.get("help"), nil
		}
	} else {
		// Even though there isn't any input, we need to let
		// Args.Parse check for required args.
		err := self.Args.Parse(args)
		if err != nil {
			return nil, err
		}
	}

	return self, nil
}
