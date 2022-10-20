package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	cmd := &cobra.Command{
		Use:     "csibuilder",
		Long:    "",
		Example: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	createCmd := newCreateCmd()
	cmd.AddCommand(createCmd)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
