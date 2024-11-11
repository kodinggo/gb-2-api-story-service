package console

import (
	"github.com/kodinggo/gb-2-api-story-service/db"
	"github.com/kodinggo/gb-2-api-story-service/internal/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCMD)
}

var serverCMD = &cobra.Command{
	Use:   "httpsrv",
	Short: "Start server",
	Run:   httpServer,
}

func httpServer(cmd *cobra.Command, args []string) {
	// TODO: implement
	config.LoadWithViper()

	mysql := db.NewMysql()
	defer mysql.Close()

}
