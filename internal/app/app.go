package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/viktoriina/databases-lab1/internal/config"
	"github.com/viktoriina/databases-lab1/internal/database"
	"github.com/viktoriina/databases-lab1/internal/message"
)

const replacementLimit = -1 // unlimited

type App struct {
	db       *database.Database
	cfg      config.Config
	commands map[string]func() error
}

func NewApp(cfg config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Start() error {
	a.initCommands()
	if err := a.initDB(); err != nil {
		return err
	}
	a.runConsole()
	return nil
}

func (a *App) initDB() error {
	a.db = database.NewDatabase(a.cfg)
	if err := a.db.Init(); err != nil {
		return err
	}
	a.db.RunIndexesSorting()
	return nil
}

func (a *App) initCommands() {
	a.commands = map[string]func() error{
		"0": func() error { return a.insertBlock() },
		"1": func() error { return a.getBlock() },
		"2": func() error { return a.updateBlock() },
		"3": func() error { return a.deleteBlock() },
		"4": func() error { return a.dumpBlocks() },

		"5": func() error { return a.insertTransaction() },
		"6": func() error { return a.getTransaction() },
		"7": func() error { return a.updateTransaction() },
		"8": func() error { return a.deleteTransaction() },
		"9": func() error { return a.dumpTransactionsPerBlock() },
	}
}

func (a *App) runConsole() {
	fmt.Println(message.Welcome)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(message.Commands)
		fmt.Print(message.CommandSeparator)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.Replace(cmd, "\n", "", replacementLimit)
		cmdHandler, supported := a.commands[cmd]
		if !supported {
			fmt.Println(message.UnsupportedCommand)
		} else {
			if err := cmdHandler(); err != nil {
				fmt.Println()
				log.Printf("Command error: %s", err)
			}
		}
	}
}

func (a *App) Shutdown() error {
	return a.db.Shutdown()
}
