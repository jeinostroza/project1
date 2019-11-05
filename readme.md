pass files from local to host
scp go/src/github.com/jeinostroza/server/main.go agent1@192.168.56.102:/home/agent1/Documents

pass file from localhost to remote host
scp otherfile.txt agent1@192.168.56.102:/home/agent1/Documents


//ProcessLog show all actives process
// func ProcessLog() {
// 	var user string
// 	var ip string
// 	var address string
// 	fmt.Println("What host do you want to work in")
// 	fmt.Print("Please enter de user for the computer:")
// 	fmt.Scanln(&user)
// 	fmt.Print("Please enter the IP:")
// 	fmt.Scanln(&ip)
// 	address = user + "@" + ip
// 	fmt.Println("")
// 	//cmd, err := exec.Command("ssh", address, "ps -A").Output()
// 	cmd, err := exec.Command("ssh", address, "cd ~/log", "ps", ">>", "log.txt", "|", "cat log.txt").Output()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}
// 	fmt.Println(string(cmd))
// 	fmt.Println("Process successfully completed")
// }

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