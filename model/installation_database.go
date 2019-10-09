package model

const (
	// InstallationDatabaseOperator is a database hosted in kubernetes via the operator.
	InstallationDatabaseOperator = "operator"
	// InstallationDatabaseRDS is a database hosted via Amazon RDS.
	InstallationDatabaseRDS = "rds"
)

// IsSupportedDatabase returns true if the given database string is supported.
func IsSupportedDatabase(database string) bool {
	return database == InstallationDatabaseOperator || database == InstallationDatabaseRDS
}
