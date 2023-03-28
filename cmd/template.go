/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
		goFiles, _, dirs := loadTemplate()
		fmt.Println("test")
		destinationDir := recreateDirs(dirs)
		renderTemplates(goFiles, destinationDir, values)
		fmt.Println(destinationDir)
	},
}

func renderTemplates(templates []string, destinationDir string, values map[string]map[string]interface{}) {
	for _, goFile := range templates {
		template, err := template.ParseFiles(goFile)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(filepath.Join(destinationDir, goFile))
		template.Execute(file, values)
	}
}

func recreateDirs(dirs []string) string {
	destinationDir, err := os.MkdirTemp("", "malaz")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		os.MkdirAll(filepath.Join(destinationDir, dir), 0755)
	}

	return destinationDir
}

func loadValues() map[string]map[string]interface{} {
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
	var values = make(map[string]map[string]interface{})
	values["Values"] = config
	return values
}

func loadTemplate() ([]string, []string, []string) {
	root := "templates"
	var goFiles []string
	var normalFiles []string
	var dirs []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() {
			if filepath.Ext(path) == ".tmpl" {
				goFiles = append(goFiles, path)
			} else {
				normalFiles = append(normalFiles, path)
			}
		} else {
			dirs = append(dirs, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return goFiles, normalFiles, dirs
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
