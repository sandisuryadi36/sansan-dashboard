package core

import (
	"database/sql"
	"os"

	featurev1 "github.com/sandisuryadi36/sansan-dashboard/gen/feature/v1"
	rolev1 "github.com/sandisuryadi36/sansan-dashboard/gen/role/v1"
	transactionv1 "github.com/sandisuryadi36/sansan-dashboard/gen/transaction/v1"
	userv1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	"github.com/sandisuryadi36/sansan-dashboard/libs"
	"github.com/sandisuryadi36/sansan-dashboard/libs/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBMain    *gorm.DB
	DBMainSQL *sql.DB
)

func StartDBConnection() {
	logger.Printf("Starting Db Connections...")

	InitDBMain()

}

func InitDBMain() {
	logger.Printf("Main Db - Connecting")
	var err error
	DBMain, err = gorm.Open(postgres.Open(libs.GetEnv("DB_DSN", "")), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed connect to DB main: %v", err)
		os.Exit(1)
		return
	}

	DBMainSQL, err = DBMain.DB()
	if err != nil {
		logger.Fatalf("Error cannot initiate connection to DB main: %v", err)
		os.Exit(1)
		return
	}

	DBMainSQL.SetMaxIdleConns(0)
	DBMainSQL.SetMaxOpenConns(0)

	err = DBMainSQL.Ping()
	if err != nil {
		logger.Fatalf("Cannot ping DB main: %v", err)
		os.Exit(1)
		return
	}

	logger.Printf("Main Db - Connected")
}

func CloseDBMain() {
	logger.Print("Closing DB Main Connection ... ")
	if err := DBMainSQL.Close(); err != nil {
		logger.Fatalf("Error on disconnection with DB Main : %v", err)
	}
	logger.Println("Closing DB Main Success")
}

func MigrateDB() error {
	InitDBMain()
	defer CloseDBMain()

	ormList := []interface{}{
		&rolev1.RoleORM{},
		&userv1.UserORM{},
		&featurev1.FeatureORM{},
		&featurev1.ServiceORM{},
		&featurev1.UserExtraFeatureORM{},
		&featurev1.FeatureTransactionORM{},
		&transactionv1.UserTransactionORM{},
	}

	logger.Println("Migration process begin...")
	if err := DBMain.AutoMigrate(
		// List table from proto gorm
		ormList...,
	); err != nil {
		logger.Fatalf("Migration failed: %v", err)
		os.Exit(1)
	}

	logger.Println("Migration process finished...")

	return nil
}
