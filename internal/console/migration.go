package console

import (
	"database/sql"
	"log"

	"github.com/kodinggo/gb-2-api-story-service/internal/config"
	"github.com/kodinggo/gb-2-api-story-service/internal/helper"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var (
	direction string
	step      int = 1
)

func init() {
	rootCmd.AddCommand(migrationCMD)

	migrationCMD.Flags().StringVarP(&direction, "direction", "d", "up", "Migration direction")
	migrationCMD.Flags().IntVarP(&step, "step", "s", 1, "Migration step")
}

var migrationCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Run:   migrateDB,
}

func migrateDB(cmd *cobra.Command, args []string) {
	// TODO: implement
	config.LoadWithViper()

	connDB, err := sql.Open("mysql", helper.GetConnectionString())
	if err != nil {
		log.Panicf("Error connecting to database, %s", err.Error())
	}

	defer connDB.Close()

	migrations := &migrate.FileMigrationSource{Dir: "./db/migrations"}

	var n int
	if direction == "down" {
		n, err = migrate.ExecMax(connDB, "mysql", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(connDB, "mysql", migrations, migrate.Up, step)
	}

	if err != nil {
		log.Panicf("Error migrating database, %s", err.Error())
	}

	log.Printf("Successfully Applied %d migrations", n)
}
