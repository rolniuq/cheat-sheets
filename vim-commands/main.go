package main

import (
	"flag"
	"fmt"
	"vim-syntax/cursor"
	cutpaste "vim-syntax/cut_paste"
	"vim-syntax/edit"
	"vim-syntax/exiting"
	"vim-syntax/global"
	"vim-syntax/help"
	indenttext "vim-syntax/indent_text"
	"vim-syntax/insert"
	"vim-syntax/macro"
	maketext "vim-syntax/make_text"
	markposition "vim-syntax/mark_position"
	"vim-syntax/register"
	searchmul "vim-syntax/search_mul"
	searchreplace "vim-syntax/search_replace"
	"vim-syntax/tab"
	viusalcmd "vim-syntax/viusal_cmd"
)

type CommandsForLife interface {
	GetCommands() map[string]string
}

func main() {
	command := flag.String("c", "", "Command to execute")
	flag.StringVar(command, "command", "", "Command to execute")
	flag.Parse()

	var c CommandsForLife

	switch *command {
	case global.GlobalName:
		c = global.NewGlobal()
	case cursor.CursorName:
		c = cursor.NewCursor()
	case insert.InsertName:
		c = insert.NewInsert()
	case edit.EditName:
		c = edit.NewEdit()
	case maketext.MakeTextName:
		c = maketext.NewMakeText()
	case viusalcmd.VisualCmdName:
		c = viusalcmd.NewVisualCmd()
	case register.RegisterName:
		c = register.NewRegister()
	case markposition.MarkPositionName:
		c = markposition.NewMarkPosition()
	case macro.MacroName:
		c = macro.NewMacro()
	case cutpaste.CutPasteName:
		c = cutpaste.NewCutPaste()
	case indenttext.IndentTextName:
		c = indenttext.NewIndentText()
	case exiting.ExitingName:
		c = exiting.NewExiting()
	case searchreplace.SearchReplaceName:
		c = searchreplace.NewSearchReplace()
	case searchmul.SearchMulName:
		c = searchmul.NewSearchMul()
	case tab.TabName:
		c = tab.NewTab()
	default:
		c = help.NewHelp()
	}

	commands := c.GetCommands()
	for k, command := range commands {
		fmt.Println(fmt.Sprintf("With %s", k), fmt.Sprintf("-- you can %s", command))
	}
}
