/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	root := newRootCMD()
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newRootCMD() *cobra.Command {
	loadConfig()
	rootCmd := &cobra.Command{
		Use:   "restic",
		Short: "restic",
		Long:  `restic`,
	}

	rootCmd.AddCommand(
		backupCmd,
		initRepoCmd,
	)
	rootCmd.PersistentFlags().StringP("repo", "r", "", "Restic repository")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Restic password")
	return rootCmd
}

func loadConfig() {
	viper.SetConfigName(".restic.config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	log.Println("Using config file:", viper.ConfigFileUsed())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myRestic.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// Read password from the configuration file using Viper

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
