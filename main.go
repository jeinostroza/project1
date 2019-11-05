package main

import (
	"fmt"

	"github.com/jeinostroza/project11/functions"
)

func main() {

	var mainoption int
	var cont = "y"

	fmt.Println("██╗     ██╗███╗   ██╗██╗   ██╗██╗  ██╗    ███████╗██╗   ██╗███████╗████████╗███████╗███╗   ███╗")
	fmt.Println("██║     ██║████╗  ██║██║   ██║╚██╗██╔╝    ██╔════╝╚██╗ ██╔╝██╔════╝╚══██╔══╝██╔════╝████╗ ████║")
	fmt.Println("██║     ██║██╔██╗ ██║██║   ██║ ╚███╔╝     ███████╗ ╚████╔╝ ███████╗   ██║   █████╗  ██╔████╔██║")
	fmt.Println("██║     ██║██║╚██╗██║██║   ██║ ██╔██╗     ╚════██║  ╚██╔╝  ╚════██║   ██║   ██╔══╝  ██║╚██╔╝██║")
	fmt.Println("███████╗██║██║ ╚████║╚██████╔╝██╔╝ ██╗    ███████║   ██║   ███████║   ██║   ███████╗██║ ╚═╝ ██║")
	fmt.Println("╚══════╝╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝  ╚═╝    ╚══════╝   ╚═╝   ╚══════╝   ╚═╝   ╚══════╝╚═╝     ╚═╝")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	for cont == "y" {
		functions.Menu()
		fmt.Println("What do you want to do?")
		fmt.Scanln(&mainoption)

		switch mainoption {
		case 1:
			functions.ListProcess()
		case 2:
			functions.KillProcess()
		case 3:
			functions.ShowSyslog()
		case 4:
			functions.DocumentsLs()
		case 5:
			functions.RemoveFile()

		}

		fmt.Println("Do you want to continue")
		fmt.Scanln(&cont)
	}
}
