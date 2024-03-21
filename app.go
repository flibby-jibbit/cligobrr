package cligobrr

import "fmt"
import "strings"

type AppFields struct {
	Name        string
	Description string
	Version     string
}

type App struct {
	AppFields
	Cmds Cmds
}

func AppNew(fields AppFields) *App {
	app := App{
		AppFields: fields,
	}

	// No alias on help or version to avoid collisions
	// with user-defined commands.
	helpFields := CmdFields{
		Name:        "help",
		Description: "Display help.",
	}

	helpCmd, _ := CmdNew(helpFields)
	app.Cmds.Add(helpCmd)

	versionFields := CmdFields{
		Name:        "version",
		Description: "Display version.",
	}

	versionCmd, _ := CmdNew(versionFields)
	app.Cmds.Add(versionCmd)

	return &app
}

func (self *App) Parse(args []string) (*Cmd, error) {
	// Remove the app name.
	args = args[1:]

	// Could need this in a couple of places, so let's
	// just grab it now.
	defCmd := self.Cmds.defaultCmd()

	// Parse the CLI input and return the command
	// that needs to be executed.

	if len(args) > 0 {
		token := args[0]

		cmd := self.Cmds.get(token)
		if cmd != nil {
			args = args[1:]

			if cmd.Name == "help" {
				appHelp(self)
				return cmd, nil
			}

			if cmd.Name == "version" {
				self.version()
				return cmd, nil
			}

			// Some other commmand.
			cmd, err := cmd.Parse(args)
			if err != nil {
				return nil, err
			} else {
				return cmd, nil
			}
		} else {
			// No command matching token. If there is a default command,
			// let's assume the input is args for that.
			if defCmd != nil {
				cmd, err := defCmd.Parse(args)
				if err != nil {
					return nil, err
				} else {
					return cmd, nil
				}
			} else {
				// No default command, so let's say we don't know what
				// to do with the input.
				return nil, errUnexpectedCmd(token)
			}
		}
	}

	// There was no input on the CLI. If there is a default command, return it.
	// Otherwise, help.
	if defCmd != nil {
		return defCmd, nil
	}

	appHelp(self)
	return self.Cmds.get("help"), nil
}

func (self *App) version() {
	ver := strings.TrimSpace(self.Version)

	if len(self.Version) == 0 {
		ver = "undefined"
	}

	fmt.Println(self.Name, "version", ver)
}
