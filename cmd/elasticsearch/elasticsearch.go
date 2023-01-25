/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package elasticsearch

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func List() {
	cmd := exec.Command("ls", "./")

	// The `Output` method executes the command and
	// collects the output, returning its value
	out, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}
	// otherwise, print the output from running the command
	fmt.Println("Output: ", string(out))

}

var err error

// elasticsearchCmd represents the elasticsearch command
func NewCmdElastic() *cobra.Command {
	var elasticsearchCmd *cobra.Command = &cobra.Command{
		Use:   "elasticsearch",
		Short: "A brief description of your command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err != nil {
				return err
			}
			return nil
		},
	}
	elasticsearchCmd.AddCommand(enableCmdIndex())

	return elasticsearchCmd
}
