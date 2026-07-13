package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite" // Pure Go driver, no CGO required!
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dbPath := getDBPath()
	log.Printf("using database at: %s", dbPath)

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return DB
}

func getDBPath() string {
	const dockerDataDir = "/data"
	if info, err := os.Stat(dockerDataDir); err == nil && info.IsDir() {
		return filepath.Join(dockerDataDir, "vouchers.db")
	}
	return "vouchers.db"
}

func MigrateDB(db *gorm.DB, models ...any) {
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}
	log.Println("database migration completed")
}
