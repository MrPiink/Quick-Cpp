package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qcpp",
	Short: "Quick C++",
	Long:  "Quick C++ is a tool to help you get started with your C++ project on Windows without the need of Visual Studio.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops, An error while executing Quick C++: %s\n", err)
		os.Exit(1)
	}
}
