package cmd

import (
	"gin-quick-start/internal/configs"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Gin web quick start controller",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.Flags().IntVarP(&configs.GetConfiguration().Port, "port", "", 9000, "Local start port")
	_ = rootCmd.MarkFlagRequired("service-port")
	rootCmd.Flags().StringVarP(&configs.GetConfiguration().Options.RunEnvironment, "env", "", "prod", "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
	configs.InitializateCallback()
}
