package cmd

import (
	"Quick-Cpp/internal"

	"github.com/spf13/cobra"

	"fmt"
)

var revertCmd = &cobra.Command{
	Use:   "revert <project name> [--msys2]",
	Short: "Revert back to before Quick C++ was used",
	Long:  "Revert back to before Quick C++ created the project with the option to remove MSYS2",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Removing NoVS project files")
		msys2, _ := cmd.Flags().GetBool("msys2")
		internal.Revert(args[0], msys2)
	},
}

func init() {
	rootCmd.AddCommand(revertCmd)
	revertCmd.Flags().Bool("msys2", false, "Remove MSYS2")
}