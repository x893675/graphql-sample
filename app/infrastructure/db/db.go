package db

import (
	"github.com/x893675/graphql-sample/app/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormschema "gorm.io/gorm/schema"
	"os"
)

func OpenDB(dsn string) *gorm.DB {

	databaseDriver := os.Getenv("DATABASE_DRIVER")

	switch databaseDriver {
	case "mysql":
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: gormschema.NamingStrategy{
				TablePrefix:   "g_",
				SingularTable: true,
			},
			//SkipDefaultTransaction: true,
			//DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			panic(err)
		}
		return db
	case "postgres":
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: false,
		}), &gorm.Config{
			NamingStrategy: gormschema.NamingStrategy{
				TablePrefix:   "g_",
				SingularTable: true,
			},
		})
		if err != nil {
			panic(err)
		}
		return db
	default:
		panic("unknown database type")
	}
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Question{}, &models.QuestionOption{}, &models.Answer{})
}
