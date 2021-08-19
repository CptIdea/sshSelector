package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sshSelector/internal/executor"
	ssh "sshSelector/internal/ssh_manager"
	"sshSelector/internal/tui"
)

func main() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path.Join(configDir,"ssh_selector"))
	manager, err := ssh.NewManagerFromFile(path.Join(configDir,"ssh_selector"))
	if err != nil {
		log.Fatal(err)
	}

	ans, err := tui.RunTUI(manager.GetList())
	if err != nil {
		log.Fatal(err)
	}

	if ans == ""{
		os.Exit(0)
	}

	err = executor.Connect(ans)
	if err != nil {
		log.Fatal(err)
	}
}
