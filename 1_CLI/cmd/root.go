package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "short decs",
	Long:  "long decs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root start")
		fmt.Println(
			cmd.Flags().Lookup("source").Value.String(),
			cmd.PersistentFlags().Lookup("License").Value.String(),
			cmd.PersistentFlags().Lookup("viper").Value.String(),
			cmd.PersistentFlags().Lookup("config").Value.String(),
			cmd.Flags().Lookup("author").Value.String(),
		)
		fmt.Println("=====================================")
		fmt.Println(
			viper.GetString("author"),
			viper.GetString("License"),
		)

		fmt.Println("root end")
	},
	TraverseChildren: true,
}

var cfgFile string
var userLicense string

func Execute() {
	rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().Bool("viper", false, "verbose output")
	rootCmd.PersistentFlags().StringP("author", "a", "your name", "log level")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "default", "this is config file")
	rootCmd.Flags().StringP("source", "s", "null", "no resource")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "License", "l", "hahahaLicense", "name of license for the project")

	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.SetDefault("author", "default author")
	viper.SetDefault("License", "default Liscense")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}
