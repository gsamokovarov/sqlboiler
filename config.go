package sqlboiler

// Config for the running of the commands
type Config struct {
	DriverName string `toml:"driver_name"`
	PkgName    string `toml:"pkg_name"`
	OutFolder  string `toml:"out_folder"`
	TableName  string `toml:"table_name"`

	Postgres PostgresConfig `toml:"postgres"`
}

// PostgresConfig configures a postgres database
type PostgresConfig struct {
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
	Host   string `toml:"host"`
	Port   int    `toml:"port"`
	DBName string `toml:"dbname"`
}