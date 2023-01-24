/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package elasticsearch

import (
	"github.com/spf13/cobra"
)

// elasticsearchCmd represents the elasticsearch command
func NewCmdCreate() *cobra.Command {
	var ElasticsearchCmd = &cobra.Command{
		Use:   "elasticsearch",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
	return ElasticsearchCmd
}
func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// elasticsearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// elasticsearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
