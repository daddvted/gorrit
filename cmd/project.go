/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Gerrit project xxx",
	Long: `
	Available commands:
		* list project
		* create project
`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("project called")
	// },
}

var projectListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Gerrit projects",
	Long:  `List Gerrit projects, blah blah ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list gerrit project")
	},
}

var projectCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Gerrit projects",
	Long:  `Create Gerrit projects, blah blah ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("creating gerrit project")
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectListCmd)
	projectCmd.AddCommand(projectCreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// projectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// projectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
