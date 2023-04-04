package postgres

import (
	"NewProUser/utils"
	"fmt"
	"sync"

	_ "github.com/lib/pq" //pq for connection

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instanceG *gorm.DB
	once      sync.Once
)

func DB() *gorm.DB {
	once.Do(func() {
		dsn, err := utils.ConnectionURLBuilder("postgres")
		if err != nil {
			panic(err)
		}
		instanceG, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Printf("GORM connections: %v", err.Error())
		}
	})
	return instanceG

}
