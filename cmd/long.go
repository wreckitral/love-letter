/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// longCmd represents the long command
var longCmd = &cobra.Command{
	Use:   "long",
    Short: "My heartfelt for you",
	Run: func(cmd *cobra.Command, args []string) {
        l, err := GetLongLetter()
        if err != nil {
            log.Fatalf("The program have a bug call the big muscular depa to fix it: %s", err.Error())
        }

        fmt.Println(l)
	},
}

func init() {
	rootCmd.AddCommand(longCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// longCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// longCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
