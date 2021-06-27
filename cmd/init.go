package cmd

import (
	"github.com/spf13/cobra"
	"gobrief-tool/codeGen"
)

// initCmd represents the run command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "example: gobrief-tool init <projectName>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}
		codeGen.ProjectName = args[0]
		runStart()
	},
}

func runStart() {
	codeGen.GenCodeFiles("./template/project")
}

func init() {
	rootCmd.AddCommand(initCmd)
}
