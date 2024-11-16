package console

import (
	"os"

	"github.com/kodinggo/gb-2-api-story-service/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "story service",
	Short: "Story Service",
	Long:  `Story Service`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	config.LoadWithViper()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	config.SetupLogger()
}
