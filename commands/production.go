package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	utils "github.com/JasnRathore/glide/utils"
)


func RunBuildForFrontend(packageManager string) {
	switch packageManager {
		case "npm":
			err := utils.RunCommand("npm", "run", "build")
			utils.Check(err)
		case "pnpm":
			err := utils.RunCommand("pnpm", "run", "build")
			utils.Check(err)
		case "yarn":
			err := utils.RunCommand("yarn", "build")
			utils.Check(err)
		case "bun":
			err := utils.RunCommand("bun", "run", "build")
			utils.Check(err)
		case "deno":
			err := utils.RunCommand("deno", "task", "build")
			utils.Check(err)
		default:
			fmt.Println("default")
	}
}

func CopyDir(src string, dst string) error {
	// Get properties of source dir
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	// Create the destination directory
	err = os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return err
	}

	// Read the source directory
	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			// Recursively copy subdirectories
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			// Copy files
			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CopyFile copies a single file from src to dst
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the file contents
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// Copy file permissions
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	return os.Chmod(dst, srcInfo.Mode())
}


func ensureDir(path string) error {
	// Check if the directory exists
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		// Doesn't exist, create it
		err := os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	} else if err != nil {
		// Some other error
		return err
	} else if !info.IsDir() {
		// Path exists but is not a directory
		return fmt.Errorf("%s exists but is not a directory", path)
	}	
	return nil
}

func RunProdBuild() {
	sourceDir := "dist"
	destinationDir := "src-glide"
	
	fmt.Println("Running Production Build")
	
	data, err := utils.ReadProjectDetails()
	utils.Check(err)
	
	RunBuildForFrontend(data.PackageManager)
	
	err = CopyDir(sourceDir, filepath.Join(destinationDir, filepath.Base(sourceDir)))
	utils.Check(err)
	
	//go into src-glide
	err = os.Chdir(destinationDir )
	utils.Check(err)
	err = ensureDir("target")
	utils.Check(err)
	
	//go build -o name.exe build.go
	exeName := fmt.Sprintf("./target/%s.exe", data.Name)
	utils.RunCommand("go", "build", "-o", exeName,"build.go")
	
	err = os.RemoveAll("dist")
	utils.Check(err)
	
}

