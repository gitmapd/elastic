/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package elasticsearch

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func execs() {
	cmd := exec.Command("ls", "-lr")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

// indexCmd represents the index command
func indexCmdCreate() *cobra.Command {
	var indexCmd = &cobra.Command{
		Use:   "index",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			execs()
		},
	}
	return indexCmd
}
