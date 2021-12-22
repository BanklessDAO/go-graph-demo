package database

import (
	"fmt"
	"go-graph-demo/app/models"
	"os"
	"time"

	gorm_logrus "github.com/onrik/gorm-logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Database struct {
// 	*gorm.DB
// }

// InitRepositories creates a connection to the database & keycloak.
// Returns the client/context to the services
func OpenDB() (db *gorm.DB, err error) {
	db, err = dialectSelector()
	if err != nil {
		return nil, err
	}

	// Seed database tables
	db.AutoMigrate(
		&models.Bounty{},
		&models.User{},
	)

	return db, nil
}

// dialectSelector selects and connects to database engine
func dialectSelector() (db *gorm.DB, err error) {
	dialect := os.Getenv("DATABASE_DIALECT")
	switch dialect {
	case "mysql":
		log.Infoln("connecting to mysql...")
		return gorm.Open(mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("MYSQL_USERNAME"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DATABASE"),
		)), &gorm.Config{
			FullSaveAssociations: true,
			Logger:               gorm_logrus.New(),
		})
	case "postgres":
		log.Infoln("connecting to postgress...")
		return gorm.Open(postgres.New(postgres.Config{
			DSN: fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
				os.Getenv("POSTGRES_HOST"),
				os.Getenv("POSTGRES_USERNAME"),
				os.Getenv("POSTGRES_PASSWORD"),
				os.Getenv("POSTGRES_DATABASE"),
				os.Getenv("POSTGRES_PORT"),
				os.Getenv("POSTGRES_TIMEZONE"),
			),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{
			FullSaveAssociations: true,
			Logger:               gorm_logrus.New(),
			NowFunc: func() time.Time {
				loc, _ := time.LoadLocation("UTC")
				timeFormatted := time.Now().In(loc).Format("2006-01-02T15:04:05")
				t, _ := time.Parse("2006-01-02T15:04:05", timeFormatted)
				return t
			},
		})
	default:
		return nil, fmt.Errorf("%s not a valid database dialect", dialect)
	}
}

// Close database connection
func Close(db *gorm.DB) {
	sql, _ := db.DB()
	sql.Close()
	log.Debugln("sql connection closed")
}
