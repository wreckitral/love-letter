/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// shortCmd represents the short command
var shortCmd = &cobra.Command{
	Use:   "short",
	Short: "To Brighten your day",
	Run: func(cmd *cobra.Command, args []string) {
        l, err := GetShortLetter()
        if err != nil {
            log.Fatalf("The program have a bug call the big muscular depa to fix it: %s", err.Error())
        }

        fmt.Println(l)
	},
}

func init() {
	rootCmd.AddCommand(shortCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shortCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shortCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
