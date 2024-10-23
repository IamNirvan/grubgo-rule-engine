package data

import (
	"sync"

	"github.com/IamNirvan/grubgo-rule-engine/internal/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func New(config *config.Config) *gorm.DB {
	once.Do(func() {
		dsn := config.GetConnectionString()
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to establish connection with database: %s", err.Error())
		}
		log.Debug("established database connection")
		instance = db
	})
	return instance
}
