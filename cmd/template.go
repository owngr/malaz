/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"text/template"

	"github.com/go-yaml/yaml"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "locally render go templates",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		values := loadValues()
		template := loadTemplate()
		template.Execute(os.Stdout, values)
	},
}

func loadValues() map[string]interface{} {
	data, err := ioutil.ReadFile("values.yaml")
	if err != nil {
		panic(err)
	}

	// Create a map to hold the YAML data
	var config map[string]interface{}

	// Unmarshal the YAML data into the map
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	return config
}

func loadTemplate() *template.Template {
	t, err := template.ParseFiles("templates/hello.go.tmpl")

	if err != nil {
		panic(err)
	}
	return t
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
