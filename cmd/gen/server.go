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
	bOverwrite bool
	StartCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string {"g", "gen"},
		Short:   "Generate code skeleton from JSON file",
		Long:    "Use when you need to generate sample code for your data model",
		Example: "generate sample.json -o",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fileName = args[0]
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().BoolVarP(&bOverwrite, "overwrite", "o", false, "Overwrite when the file is already there")
}

func run() {
	fmt.Println(`Generating code skeletons...`)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("can not read from json file '%s', %v\n", fileName, err)
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
