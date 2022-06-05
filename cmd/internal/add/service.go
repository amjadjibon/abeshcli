package add

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/amjadjibon/abeshcli/cmd/internal/base"
	"github.com/amjadjibon/abeshcli/model"
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

		if len(path) == 0 {
			path = "/abesh/v1/echo"
		}

		method = strings.ToUpper(method)
		if len(method) == 0 {
			path = "GET"
		}

		if !IsInAllowedMethodList(method) {
			fmt.Println("method not allowed")
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
		defer func() {
			_ = nameFile.Close()
		}()

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
		defer func() {
			_ = contractIdFile.Close()
		}()

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
		defer func() {
			_ = categoryFile.Close()
		}()

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
			fmt.Println(err)
			os.Exit(0)
		}

		modulePath, err := base.ModulePath("go.mod")
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
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

		var manifestFilePath = "main/mono/manifest.yaml"
		yamlFile, err := ioutil.ReadFile(manifestFilePath)
		if err != nil {
			panic(err)
		}

		var manifestModel = &model.Manifest{}
		err = yaml.Unmarshal(yamlFile, manifestModel)
		if err != nil {
			panic(err)
		}

		manifestModel.Capabilities = append(manifestModel.Capabilities,
			&model.CapabilityManifest{
				ContractId: "abesh:" + service,
			},
		)

		var triggerValues = make(model.ConfigMap)
		triggerValues["method"] = method
		triggerValues["path"] = path

		manifestModel.Triggers = append(manifestModel.Triggers, &model.TriggerManifest{
			Trigger:       "abesh:httpserver2",
			TriggerValues: triggerValues,
			Service:       "abesh:" + service,
		})

		manifestYaml, err := yaml.Marshal(manifestModel)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(manifestFilePath, manifestYaml, 0644)
		if err != nil {
			panic(err)
		}

		var mainFilePath = "main/mono/main.go"
		var importStr = "	_ \"github.com/amjadjibon/hello/capability/echo\""

		err = utils.InsertStringToFile2(mainFilePath, importStr)
		if err != nil {
			panic(err)
		}
	},
}

var (
	service string
	method  string
	path    string
)

func init() {
	CmdAddService.Flags().StringVarP(&service, "service", "s", service, "Service Name, Ex: echo")
	CmdAddService.Flags().StringVarP(&method, "method", "m", service, "Service Method, Ex: GET")
	CmdAddService.Flags().StringVarP(&path, "path", "p", service, "Service Path, Ex: /abesh/v1/echo")
	_ = CmdAddService.MarkFlagRequired("service")
}

func IsInAllowedMethodList(method string) bool {
	var methodList = []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}
	for _, m := range methodList {
		if m == method {
			return true
		}
	}
	return false
}
