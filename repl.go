package main

import (
	"nano-cli/repl"

	"github.com/spf13/cobra"
)

var docsRoute string
var fileName string
var prettyJSON bool

// replCmd opens nano REPL tool
var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "starts nano repl tool",
	Long:  `starts nano repl tool`,
	Run: func(cmd *cobra.Command, args []string) {
		repl.Start(docsRoute, fileName, prettyJSON)
	},
}

func init() {
	replCmd.Flags().StringVarP(&docsRoute, "docs", "d", "", "route containing the documentation")
	replCmd.Flags().StringVarP(&fileName, "filename", "f", "", "file containing the commands to run")
	replCmd.Flags().BoolVarP(&prettyJSON, "pretty", "p", false, "print pretty jsons")
	rootCmd.AddCommand(replCmd)
}
