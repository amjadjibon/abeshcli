package build

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// CmdBuild build project command.
var CmdBuild = &cobra.Command{
	Use:   "build",
	Short: "build project",
	Long:  "build project. Example: kratos run",
	Run:   Build,
}

var (
	mainFilePath  string
	buildFilePath string
)

func init() {
	CmdBuild.Flags().StringVarP(&mainFilePath, "main-file-path", "m", mainFilePath, "main file path")
	CmdBuild.Flags().StringVarP(&buildFilePath, "build-file-path", "b", buildFilePath, "build file path")
}

// Build project.
func Build(cmd *cobra.Command, args []string) {
	var dir string
	if len(args) > 0 {
		dir = args[0]
	}
	base, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
		return
	}
	if dir == "" {
		// find the directory containing the cmd/*
		cmdPath, err := findCMD(base)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
			return
		}
		switch len(cmdPath) {
		case 0:
			fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", "The cmd directory cannot be found in the current directory")
			return
		case 1:
			for _, v := range cmdPath {
				dir = v
			}
		default:
			var cmdPaths []string
			for k := range cmdPath {
				cmdPaths = append(cmdPaths, k)
			}
			prompt := &survey.Select{
				Message:  "Which directory do you want to run?",
				Options:  cmdPaths,
				PageSize: 10,
			}
			e := survey.AskOne(prompt, &dir)
			if e != nil || dir == "" {
				return
			}
			dir = cmdPath[dir]
		}
	}

	if len(mainFilePath) == 0 {
		mainFilePath = "main/mono/main.go"
	}

	if len(buildFilePath) == 0 {
		mainFilePath = "bin/mono"
	}

	fd := exec.Command("go", "build", "-v", "-o", buildFilePath, mainFilePath)
	fd.Stdout = os.Stdout
	fd.Stderr = os.Stderr
	fd.Dir = dir
	if err := fd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err.Error())
		return
	}
}

func findCMD(base string) (map[string]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(wd, "/") {
		wd += "/"
	}
	var root bool
	next := func(dir string) (map[string]string, error) {
		cmdPath := make(map[string]string)
		err := filepath.Walk(dir, func(walkPath string, info os.FileInfo, err error) error {
			// multi level directory is not allowed under the cmdPath directory, so it is judged that the path ends with cmdPath.
			if strings.HasSuffix(walkPath, "cmd") {
				paths, err := os.ReadDir(walkPath)
				if err != nil {
					return err
				}
				for _, fileInfo := range paths {
					if fileInfo.IsDir() {
						abs := path.Join(walkPath, fileInfo.Name())
						cmdPath[strings.TrimPrefix(abs, wd)] = abs
					}
				}
				return nil
			}
			if info.Name() == "go.mod" {
				root = true
			}
			return nil
		})
		return cmdPath, err
	}
	for i := 0; i < 5; i++ {
		tmp := base
		cmd, err := next(tmp)
		if err != nil {
			return nil, err
		}
		if len(cmd) > 0 {
			return cmd, nil
		}
		if root {
			break
		}
		_ = filepath.Join(base, "..")
	}
	return map[string]string{"": base}, nil
}
