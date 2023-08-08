/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"myRestic/logging"
	"os"
	"os/exec"
	"strings"
)

// initRepoCmd represents the initRepo command
var initRepoCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: initRepoFunc,
}

func initRepoFunc(cmd *cobra.Command, args []string) {
	repoLocation, _ := cmd.Flags().GetString("repo")
	password, _ := cmd.Flags().GetString("password")
	if password == "" {
		password = viper.GetString("password")
	}
	if repoLocation == "" {
		repoLocation = viper.GetString("repo")
		if repoLocation != "" {
			cmdArgs := []string{"init", "--repo", repoLocation}
			// If a password is provided, use stdin pipe to pass it
			if err := runResticCommand(password, cmdArgs...); err != nil {
				fmt.Println("Error initializing repository:", err)
			} else {
				logging.Get().Info("Repository initialized successfully.")
			}
		} else {
			logging.Get().Error("Please mention repo in config/terminal...")
		}
		return
	} else {
		cmdArgs := []string{"init", "--repo", repoLocation}
		if err := runResticCommand(password, cmdArgs...); err != nil {
			fmt.Println("Error initializing repository:", err)
		} else {
			logging.Get().Info("Repository initialized successfully.")
		}
		return
	}

}

func runResticCommand(password string, args ...string) error {
	cmd := exec.Command("restic", args...)
	if password != "" {
		cmd.Stdin = strings.NewReader(password)
	} else {
		cmd.Stdin = os.Stdin
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initRepoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initRepoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
