package functions

import (
	"fmt"
	"log"
	"os/exec"
)

//Menu is the main menu
func Menu() {

	fmt.Println("MAIN MENU")

	var mainM = make(map[int]string)
	mainM[1] = "1.- List active processes"
	mainM[2] = "2.- Kill process"
	mainM[3] = "3.- Log file"
	mainM[4] = "4.- Documents files"
	mainM[5] = "5.- Remove File"

	var i int
	for i = 0; i <= len(mainM); i++ {
		fmt.Println(mainM[i])
	}

}

//ListProcess show all actives process
func ListProcess() {
	var user string
	var ip string
	var address string
	fmt.Println("What host do you want to work in")
	fmt.Print("Remote host's username:")
	fmt.Scanln(&user)
	fmt.Print("Please enter host's IP:")
	fmt.Scanln(&ip)
	address = user + "@" + ip
	fmt.Println("")
	cmd, err := exec.Command("ssh", address, "ps", "aux").Output()
	// cmd, err := exec.Command("ssh", address, "cd ~/log", "ps", ">>", "log.txt", "|", "cat log.txt").Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(cmd))
	fmt.Println("Process successfully completed")
}

//KillProcess kill selected process
func KillProcess() {
	var user string
	var ip string
	var address string
	var proc string
	fmt.Println("What host do you want to work in")
	fmt.Print("Remote host's username:")
	fmt.Scanln(&user)
	fmt.Print("Please enter host's IP:")
	fmt.Scanln(&ip)
	fmt.Println("What process do you want to finish")
	fmt.Scanln(&proc)
	address = user + "@" + ip
	fmt.Println("")
	cmd, err := exec.Command("ssh", address, "pkill", proc).Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(cmd))
	fmt.Println("Process successfully completed")
}

//ShowSyslog show the logsys file
func ShowSyslog() {
	var user string
	var ip string
	var address string
	fmt.Println("What host do you want to work in")
	fmt.Print("Remote host's username:")
	fmt.Scanln(&user)
	fmt.Print("Please enter host's IP:")
	fmt.Scanln(&ip)
	address = user + "@" + ip
	fmt.Println("")
	//cmd, err := exec.Command("ssh", address, "ps -A").Output()
	//cmd, err := exec.Command("ssh", address, "cd ~/log", "ps aux", ">>", "log.txt", "|", "cat log.txt").Output()
	cmd, err := exec.Command("ssh", address, "less ", "/var/log/syslog").Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(cmd))
	fmt.Println("Process successfully completed")
}

//DocumentsLs show the files inside the Documents folder
func DocumentsLs() {
	var user string
	var ip string
	var address string
	fmt.Println("What host do you want to work in")
	fmt.Print("Remote host's username:")
	fmt.Scanln(&user)
	fmt.Print("Please enter host's IP:")
	fmt.Scanln(&ip)
	address = user + "@" + ip
	fmt.Println("")
	cmd, err := exec.Command("ssh", address, "ls ", "/home/agent1/Documents").Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(cmd))
	fmt.Println("Process successfully completed")
}

//PassFilesLH transfer files from local to host
func PassFilesLH() {
	var user string
	var ip string
	var address string
	var folder string
	var file string
	var ruta2 string
	cmd1, err := exec.Command("ls", "/home/jeinostroza/Documents").Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(cmd1))
	fmt.Println("What host do you want to work in")
	fmt.Print("Remote host's username:")
	fmt.Scanln(&user)
	fmt.Print("Please enter host's IP:")
	fmt.Scanln(&ip)
	address = user + "@" + ip
	fmt.Println("")
	fmt.Println("Please enter the folder")
	fmt.Scanln(&folder)
	fmt.Println("Please enter the file name")
	fmt.Scanln(&file)
	//ruta = " go/src/github.com/jeinostroza/" + folder + "/" + file
	ruta2 = " " + address + ":/home/agent1/Documents"
	// fmt.Println(ruta)
	// fmt.Println(ruta2)
	//cmd, err := exec.Command("ssh", address, "ps -A").Output()
	//cmd, err := exec.Command("ssh", address, "cd ~/log", "ps aux", ">>", "log.txt", "|", "cat log.txt").Output()
	cmd, err := exec.Command("scp", file, ruta2).Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(cmd))
	fmt.Println("Process successfully completed")
}

//ssh agent1@192.168.56.102 touch /home/agent1/Documents/parapapan.txt

//CreateFile create a file in documents folder
func CreateFile() {
	var user string
	var ip string
	var address string
	fmt.Println("What host do you want to work in")
	fmt.Print("Remote host's username:")
	fmt.Scanln(&user)
	fmt.Print("Please enter host's IP:")
	fmt.Scanln(&ip)
	address = user + "@" + ip
	fmt.Println("")
	var file string
	fmt.Println("Please enter the file name")
	fmt.Scanln(&file)

	cmd, err := exec.Command("ssh " + address + "touch " + "/home/agent1/Documents/" + file).Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(cmd))
	fmt.Println("Process successfully completed")
}

//RemoveFile delete a file
func RemoveFile() {
	var user string
	var ip string
	var address string
	var file string
	fmt.Println("What host do you want to work in")
	fmt.Print("Remote host's username:")
	fmt.Scanln(&user)
	fmt.Print("Please enter host's IP:")
	fmt.Scanln(&ip)
	address = user + "@" + ip
	fmt.Println("")
	cmd, err := exec.Command("ssh", address, "ls ", "/home/agent1/Documents").Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(string(cmd))
	fmt.Println("What file do you want to delete")
	fmt.Scanln(&file)
	fmt.Println(address)
	fmt.Println(file)
	cmd1, err := exec.Command("ssh", address, "rm ", "/home/agent1/Documents/"+file).Output()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Println(string(cmd1))

	fmt.Println("Process successfully completed")
}

//ssh agent1@192.168.56.102 rm /home/agent1/Documents/averquetal.txt
