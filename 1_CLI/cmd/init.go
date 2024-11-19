package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "add",
	Short: "short",
	Long:  "long",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run init begin ")
		fmt.Println(
			//cmd.Flags().Lookup("source").Value.String(),
			cmd.Flags().Lookup("License").Value.String(),
			cmd.Flags().Lookup("viper").Value.String(),
			viper.GetString("author"),
			cmd.Flags().Lookup("config").Value.String(),
			cmd.Parent().Flags().Lookup("source").Value,
		)
		fmt.Println("run init end ")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
