package main

import (
	"os"
	"fmt"
	cmd "glide/commands"
)

func main() {
	size := len(os.Args)
	if (size < 2) {
		fmt.Println("Docs")
		return
	}
	args := os.Args[1:]
	switch args[0] {
		case "init":
			cmd.InitProject();
		case "dev":
			cmd.RunDevBuild();
		case "build":
			cmd.RunProdBuild();
		default:
			fmt.Println("Invalid Arguments")
	}
}
