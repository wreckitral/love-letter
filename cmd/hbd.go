/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hbdCmd represents the hbd command
var hbdCmd = &cobra.Command{
	Use:   "hbd",
	Short: "For your special birthday",
	Run: func(cmd *cobra.Command, args []string) {
        fmt.Println(Wish)
	},
}

func init() {
	rootCmd.AddCommand(hbdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hbdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hbdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
