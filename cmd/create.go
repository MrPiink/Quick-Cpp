package cmd

import (
	"fmt"

	"Quick-Cpp/internal"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <project name> [--full] [--skip-msys] [--no-files]",
	Short: "Create a new project",
	Long:  "Create a new project with the specified layout",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Creating Project: %s\n", args[0])
		full, _ := cmd.Flags().GetBool("full")
		skipMsys, _ := cmd.Flags().GetBool("skip-msys")
		noFiles, _ := cmd.Flags().GetBool("no-files")
		internal.Create(args[0], full, skipMsys, noFiles)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().Bool("full", false, "Include optional top-level pitchfork directories")
	createCmd.Flags().Bool("skip-msys", false, "Skip MSYS2 installation")
	createCmd.Flags().Bool("no-files", false, "No file initialization")
}