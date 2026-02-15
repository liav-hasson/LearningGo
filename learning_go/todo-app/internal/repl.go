package internal

import (
	"bufio"
	"os"
	"strings"
)

func InitTasks() {
	getExistingTaskList()
}

func StartRepl() {
	printHeader()

	for {
		println("	Enter a command to execute:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		rawCmd := scanner.Text()

		parsedCmd := strings.ToLower(rawCmd)

		switch parsedCmd {
		case "new", "n":
			createTask()
			printTask()
		case "edit", "e":
			chooseTask()
			editTask()
			printTask()
		case "delete", "d":
			chooseTask()
			deleteTask()
			printTask()
		case "print", "p":
			printAllTasks()
		case "help", "h":
			printHelp()
		case "save", "s":
			saveTasks()
		default:
			println("Invalid argument, try again.")
			printHelp()
		}
	}

}
