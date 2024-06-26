package cligobrr

import "fmt"
import "strings"
import "strconv"

func appHelp(app *App, input []string) error {
	return help(
		app.Name,
		app.Description,
		app.Cmds,
		app.Args,
		input,
	)
}

func cmdHelp(cmd *Cmd, input []string) error {
	return help(
		cmd.Name,
		cmd.Description,
		cmd.Cmds,
		cmd.Args,
		input,
	)
}

func help(
	name string,
	description string,
	commands Cmds,
	arguments Args,
	input []string,
) error {
	helpHeader(name, description)

	if len(input) > 0 {
		token := input[0]
		arg := arguments.get(token)
		if arg == nil {
			return errUnexpectedArg(token)
		}

		helpUsage(name, []IArg{*arg})
		helpSingleArg(*arg)
	} else {
		args := arguments.args
		if len(args) > 0 {
			helpUsage(name, args)
			helpAllArgs(name, args)
		}

		helpCmds(commands.cmds)
	}

	return nil
}

func helpHeader(name, description string) {
	tableFields := TableFields{
		Cols: 2,
	}

	table, _ := tableNew(tableFields)
	table.Add([]string{"Name:", name})
	if len(description) > 0 {
		table.Add([]string{"Description:", description})
	}

	fmt.Println(table.ToString())
	fmt.Println("")
}

func helpUsage(name string, args []IArg) {
	output := []string{"Usage:", ""}
	cmdLine := []string{name}

	for _, arg := range args {
		fragment := fmt.Sprintf("%s=%s", arg.GetName(), arg.GetKind())

		if !arg.GetRequired() {
			fragment = fmt.Sprintf("[%s]", fragment)
		}

		cmdLine = append(cmdLine, fragment)
	}

	output = append(output, strings.Join(cmdLine, " "))
	fmt.Println(strings.Join(output, "\n"))
	fmt.Println("")
}

func helpSingleArg(arg IArg) {
	tableFields := TableFields{
		Cols: 2,
	}

	table, _ := tableNew(tableFields)
	table.Add([]string{"Name:", arg.GetName()})
	table.Add([]string{"Alias:", arg.GetAlias()})
	table.Add([]string{"Description:", arg.GetDescription()})
	table.Add([]string{"Kind:", arg.GetKind()})
	table.Add([]string{"Multiple:", strconv.FormatBool(arg.GetMultiple())})
	table.Add([]string{"Required:", strconv.FormatBool(arg.GetRequired())})
	table.Add([]string{"Default:", arg.GetDefault()})
	table.Add([]string{"Choices:", strings.Join(arg.GetChoices(), arg.GetSeparator())})
	fmt.Println(table.ToString())
	fmt.Println("")
}

func helpAllArgs(name string, args []IArg) {

	output := []string{
		"Arguments:",
		"",
	}

	tableFields := TableFields{
		Cols: 2,
	}

	table, _ := tableNew(tableFields)
	table.Add([]string{"Name", "Description"})
	table.Add([]string{"----", "-----------"})

	for _, arg := range args {
		table.Add([]string{arg.GetName(), arg.GetDescription()})
	}

	output = append(output, table.ToString())
	output = append(output, "")
	output = append(output, fmt.Sprintf("`%s help arg` for more information.", name))
	fmt.Println(strings.Join(output, "\n"))
	fmt.Println("")
}

func helpCmds(cmds []Cmd) {
	output := []string{
		"Commands:",
		"",
	}

	tableFields := TableFields{
		Cols: 3,
	}

	table, _ := tableNew(tableFields)
	table.Add([]string{"Name", "Alias", "Description"})
	table.Add([]string{"----", "-----", "-----------"})

	hasDefault := false

	for _, cmd := range cmds {
		name := cmd.Name

		if cmd.Default {
			name = fmt.Sprintf("%s*", name)
			hasDefault = true
		}

		table.Add([]string{
			name,
			cmd.Alias,
			cmd.Description,
		})
	}

	output = append(output, table.ToString())

	if hasDefault {
		output = append(output, "")
		output = append(output, "* indicates default command")
	}

	fmt.Println(strings.Join(output, "\n"))
}
