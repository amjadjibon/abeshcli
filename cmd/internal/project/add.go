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

var repoAddIgnores = []string{
	".git", ".github", "api", "README.md", "LICENSE", "go.mod", "go.sum", "third_party",
}

func (p *Project) Add(ctx context.Context, dir string, layout string, branch string, mod string) error {
	to := path.Join(dir, p.Path)

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
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

	fmt.Printf("🚀 Add service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)

	repo := base.NewRepo(layout, branch)

	if len(modulePath) == 0 {
		modulePath = p.Path
	}

	if err := repo.CopyToV2(ctx, to, path.Join(mod, modulePath), repoAddIgnores, []string{path.Join(p.Path, "api"), "api"}); err != nil {
		return err
	}

	e := os.Rename(
		path.Join(to, "cmd", "server"),
		path.Join(to, "cmd", p.Name),
	)
	if e != nil {
		return e
	}

	base.Tree(to, dir)

	fmt.Printf("\n🍺 Project creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to start the project 👇:\n\n")

	fmt.Println(color.WhiteString("$ cd %s", p.Name))
	fmt.Println(color.WhiteString("$ go mod download"))
	fmt.Println(color.WhiteString("$ go mod tidy"))
	fmt.Println(color.WhiteString("$ abesh run"))
	fmt.Println("Thanks for using abesh")
	return nil
}
