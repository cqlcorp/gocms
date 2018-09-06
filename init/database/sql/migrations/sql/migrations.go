package migrations

import (
	"github.com/rubenv/sql-migrate"
)

func Default() *migrate.MemoryMigrationSource {

	migrationsList := migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			CreateInitial(),
			AddLastModified(),
			AddACL(),
			AddExternalPlugin(),
			MigrateToRSAKeys(),
			AddDocumentationToggle(),
			ErrorReportingMigration(),
		},
	}
	return &migrationsList
}
