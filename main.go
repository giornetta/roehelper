package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/shirou/gopsutil/process"
)

var welcomestr = `
=================================================================
Welcome to the ROE Helper!
This tool will generate a batch file that you can use to play
Ring Of Elysium in English!
Make sure to start the game from Garena before running this tool,
and also make sure to run this in Administrator Mode!


source code for this tool lives at github.com/giornetta/roehelper
=================================================================
`

func main() {
	fmt.Println(welcomestr)
	var cmd string

	procs, err := process.Processes()
	if err != nil {
		fmt.Printf("Could not get running processes: %v", err)
		quit(1)
	}

	var found bool
	fmt.Println("Trying to find Ring Of Elysium process...")
	for _, p := range procs {
		e, _ := p.Exe()
		exe := strings.Split(e, "\\")
		if exe[len(exe)-1] != "Europa_Client.exe" {
			continue
		}

		cmd, _ = p.Cmdline()
		found = true
		break
	}

	if !found {
		fmt.Println("Could not find ROE's process... Are you sure it's running? If it is, run this in Administrator mode!")
		quit(1)
	}

	fmt.Println("Finally found it!")

	sl := strings.Split(cmd, "-")

	for i, s := range sl {
		if strings.Contains(s, "language=th") {
			fmt.Println("Changing game's language to English...")
			sl[i] = "language=en "
		}
	}
	sl = append(sl, "-toggleapp=32838")

	fmt.Println("Generating batch file...")
	f, _ := os.Create("ROE.bat")
	defer f.Close()
	str := "@echo off\n\n"
	for _, s := range sl {
		str += s
	}
	f.Write([]byte(str))

	fmt.Println("Done! You can know use the generated ROE.bat file to run the game!")
	quit(0)
}

func quit(code int) {
	fmt.Scanln()
	os.Exit(code)
}
