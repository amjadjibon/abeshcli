package project

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"

	"github.com/amjadjibon/abeshcli/cmd/internal/base"
)

// Project is a project template.
type Project struct {
	Name string
	Path string
}

// New a project from remote repo.
func (p *Project) New(ctx context.Context, dir string, layout string, branch string) error {
	to := path.Join(dir, p.Name)
	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("š« %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "š Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		_ = os.RemoveAll(to)
	}
	fmt.Printf("š Creating service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)
	repo := base.NewRepo(layout, branch)

	if len(modulePath) == 0 {
		modulePath = p.Path
	}

	if err := repo.CopyTo(ctx, to, modulePath, []string{".git", ".github"}); err != nil {
		return err
	}

	base.Tree(to, dir)

	fmt.Printf("\nšŗ Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("š» Use the following command to start the project š:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go mod download"))
	fmt.Println(color.WhiteString("$ go mod tidy"))
	fmt.Println(color.WhiteString("$ abesh run"))
	fmt.Println("Thanks for using abesh")
	return nil
}
