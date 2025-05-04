package commands

import (
	"context"
	"fmt"
	"os"
	utils "glide/utils"
)

func getPackageManagerCommand(pm string) (string, []string) {
	switch pm {
	case "npm":
		return "npm", []string{"run", "dev"}
	case "pnpm":
		return "pnpm", []string{"run", "dev"}
	case "yarn":
		return "yarn", []string{"dev"}
	case "bun":
		return "bun", []string{"run", "dev"}
	case "deno":
		return "deno", []string{"run", "dev"}
	default:
		return "echo", []string{"Unknown package manager"}
	}
}

func RunDevBuild() {
	fmt.Println("Running Dev Build")
	
	data, err := utils.ReadProjectDetails()
	utils.Check(err)
	fmt.Println(data.PackageManager)
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start web app
	cmd, args := getPackageManagerCommand(data.PackageManager)
	webCmd, err := utils.StartCommandWithContext(ctx, cmd, args...)
	utils.Check(err)

	// Change directory
	dirName := "src-glide"
	err = os.Chdir(dirName)
	utils.Check(err)
	
	// Start air
	airCmd, err := utils.StartCommandWithContext(ctx, "air")
	utils.Check(err)

	// Setup signal handler for both commands
	utils.SetupSignalHandler(cancel, webCmd, airCmd)

	// Wait for context cancellation
	<-ctx.Done()
}