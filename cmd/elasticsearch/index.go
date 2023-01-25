/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package elasticsearch

import (
	"github.com/spf13/cobra"
)

// indexCmd represents the index command
/*var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(os.Args[1:])
	},
}
*/

func enableCmdIndex() *cobra.Command {
	var indexCmd = &cobra.Command{
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			List()
		},
	}
	return indexCmd

}
