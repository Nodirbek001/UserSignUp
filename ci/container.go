package ci

import (
	"NewProUser/configs"
	"NewProUser/utils"
	"log"
	"strings"
	"sync"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	once sync.Once

	cfg = configs.Load()
)

func MigrationsUp(){
	url, err:=utils.ConnectionURLBuilder("migration")
	if err!=nil {
		log.Println("Error generating migration url: ", err.Error())
	}
	m, err:=migrate.New("file://migrations", url)
	if err!=nil {
		log.Fatal("error in creating migrations:", err.Error())
	}

	if err:=m.Up(); err!=nil {
		if !strings.Contains(err.Error(),"no change") {
			log.Println("Error in migrating", err.Error())
		}
	}
}
							