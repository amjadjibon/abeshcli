package add

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/amjadjibon/abeshcli/cmd/internal/base"
	"github.com/amjadjibon/abeshcli/template"
	"github.com/amjadjibon/abeshcli/utils"
)

// CmdAddService run project command.
var CmdAddService = &cobra.Command{
	Use:   "service",
	Short: "add service capability",
	Long:  "add service capability",
	Run: func(cmd *cobra.Command, args []string) {
		if len(service) == 0 {
			fmt.Println("service name can not be empty")
			os.Exit(0)
		}

		if !utils.FileExists("go.mod") {
			fmt.Println("go.mod file does not exist")
			fmt.Println("use: `abesh new` command")
			os.Exit(0)
		}

		if !utils.FolderExists("capability") {
			fmt.Println("capability folder is does not exist")
			fmt.Println("use: `abesh new` command")
			os.Exit(0)
		}

		var filePath = "capability/" + service
		err := os.Mkdir("capability/"+service, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		nameFile, err := os.Create(filePath + "/name.go")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		defer nameFile.Close()

		err = template.NameTemplate.ExecuteTemplate(nameFile, "name", struct {
			TimeStamp   string
			PackageName string
		}{
			TimeStamp:   time.Now().Format(time.RFC850),
			PackageName: strings.ToLower(service),
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		contractIdFile, err := os.Create(filePath + "/contractid.go")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		defer contractIdFile.Close()

		err = template.ContractIdTemplate.ExecuteTemplate(contractIdFile, "contractid", struct {
			TimeStamp   string
			PackageName string
		}{
			TimeStamp:   time.Now().Format(time.RFC850),
			PackageName: strings.ToLower(service),
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		categoryFile, err := os.Create(filePath + "/category.go")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		defer categoryFile.Close()

		err = template.CategoryTemplate.ExecuteTemplate(categoryFile, "category", struct {
			TimeStamp   string
			PackageName string
		}{
			TimeStamp:   time.Now().Format(time.RFC850),
			PackageName: strings.ToLower(service),
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		filePath = filePath + "/" + service
		f, err := os.Create(filePath + ".go")
		if err != nil {
			panic(err)
		}

		modulePath, err := base.ModulePath("go.mod")
		if err != nil {
			panic(err)
		}

		err = template.ServiceTemplate.ExecuteTemplate(f, "service", struct {
			TimeStamp     string
			PackageName   string
			InterfaceName string
			InterfaceRune string
			ServiceName   string
			ContractID    string
			ModulePath    string
		}{
			TimeStamp:     time.Now().Format(time.RFC850),
			PackageName:   strings.ToLower(service),
			InterfaceName: strings.ToUpper(string(service[0])) + service[1:],
			InterfaceRune: string(service[0]),
			ServiceName:   "abesh_" + service,
			ContractID:    "abesh:" + service,
			ModulePath:    modulePath,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	},
}

var (
	service string
)

func init() {
	CmdAddService.Flags().StringVarP(&service, "service", "s", service, "Service Name, Ex: echo")
	_ = CmdAddService.MarkFlagRequired("service")
}
