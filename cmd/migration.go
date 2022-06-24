package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/kitabisa/golang-poc/config"
	"github.com/kitabisa/golang-poc/internal/app/appcontext"
	migrate "github.com/rubenv/sql-migrate"
	plog "github.com/kitabisa/perkakas/v2/log"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Up DB golang-poc",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config()
		appCtx := appcontext.NewAppContext(c)
		mSource := getMigrateSource()
		doMigrate(appCtx, mSource, appcontext.DBDialectPostgres, migrate.Up)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate Up DB golang-poc",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config()
		appCtx := appcontext.NewAppContext(c)
		mSource := getMigrateSource()
		doMigrate(appCtx, mSource, appcontext.DBDialectPostgres, migrate.Down)
	},
}

var migrateNewCmd = &cobra.Command{
	Use:   "migratenew [migration name]",
	Short: "Create new migration file",
	Long:  `Create new migration file on folder migrations/sql with timestamp as prefix`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mDir := "migrations/sql/"

		createMigrationFile(mDir, args[0])
	},
}

func init() {
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateNewCmd)
}

func getMigrateSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: "migrations/sql",
	}

	return source
}

func doMigrate(appCtx *appcontext.AppContext, mSource migrate.FileMigrationSource, dbDialect string, direction migrate.MigrationDirection) error {
	db, err := appCtx.GetDBInstance(dbDialect)
	if err != nil {
		plog.Zlogger(context.Background()).Fatal().Err(err).Msg("Error connection to DB")
		return err
	}

	defer db.Db.Close()

	total, err := migrate.Exec(db.Db, dbDialect, mSource, direction)
	if err != nil {
		plog.Zlogger(context.Background()).Err(err).Msg("Fail migration")
		return err
	}

	plog.Zlogger(context.Background()).Info().Msgf("Migrate Success, total migrated: %d", total)
	return nil
}

func createMigrationFile(mDir string, mName string) error {
	var migrationContent = `-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- [your SQL script here]

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- [your SQL script here]
`
	filename := fmt.Sprintf("%d_%s.sql", time.Now().Unix(), mName)
	filepath := fmt.Sprintf("%s%s", mDir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		plog.Zlogger(context.Background()).Err(err).Str("filepath", filepath).Msg("Error create migration file")
		return err
	}
	defer f.Close()

	f.WriteString(migrationContent)
	f.Sync()

	plog.Zlogger(context.Background()).Info().Str("filepath", filepath).Msg("New migration file has been created")
	return nil
}
