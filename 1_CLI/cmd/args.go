package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var cusArgsCheckCmd = &cobra.Command{
	Use: "cusargs",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("at least one arg")
		}
		if len(args) > 2 {
			return errors.New("at most two args")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cusArgsCheckCmd start")

		fmt.Println(args)

		fmt.Println("cusArgsCheckCmd end")
	},
}

var argsCmd = &cobra.Command{
	Use:       "args",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{"abc", "123", "tom"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("argsCmd start")
		fmt.Println(args)
		fmt.Println("argsCmd end")
	},
}

func init() {
	rootCmd.AddCommand(cusArgsCheckCmd)
	rootCmd.AddCommand(argsCmd)
}
