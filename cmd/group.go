/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var username string

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Gerrit group command",
	Long: `Commands to operate on Gerrit groups.
Current supportted:
	* List groups
	* Get groups for specific user

`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().Changed("username") {
			fmt.Println(username)
		} else {
			fmt.Println(args)
			cmd.Help()
			os.Exit(0)
		}
	},
}

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List Gerrit groups",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list gerrit groups")
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)
	groupCmd.AddCommand(groupListCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	groupCmd.Flags().StringVarP(&username, "username", "u", "", "Gerrit username(short name)")

}
