package gen

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var (
	fileName string
	StartCmd = &cobra.Command{
		Use:     "generate",
		Short:   "Generate code skeleton",
		Long:    "Use when you need to generate sample code for your data model",
		Example: "generate -f sample.json",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "Generate source code files from json")
}

func run() {
	fmt.Println(`Generating code skeletons...`)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("can not read from %s, %v\n", fileName, err)
		os.Exit(-1)
	}

	var tab = SysTables{}
	if err := json.Unmarshal(data, &tab); err != nil {
		fmt.Printf("can not parse the json file, %v\n", err)
		os.Exit(-2)
	}

	var gen = Gen{}
	gen.GenCode(&tab)
}
