package database

func Migrate(models ...interface{}) error {
	return DB.AutoMigrate(models...)
}
