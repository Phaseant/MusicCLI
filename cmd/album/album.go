/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package album

import (
	"github.com/spf13/cobra"
)

// AlbumCmd represents the album command
var AlbumCmd = &cobra.Command{
	Use:   "album",
	Short: "Commands to manage albums",
	Long:  `Commands to manage albums`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// albumCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// albumCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
