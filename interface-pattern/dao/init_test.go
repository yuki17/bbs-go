package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/shohhei1126/bbs-go/common/db"
	"github.com/shohhei1126/bbs-go/common/log"
	"github.com/shohhei1126/bbs-go/interface-pattern/model"
	"gopkg.in/gorp.v1"
	"os"
	"testing"
)

var dbMap *gorp.DbMap

func TestMain(m *testing.M) {
	// db
	dbString := os.Getenv("BBSGO_DB_TEST")
	if dbString == "" {
		dbString = "root:@tcp(localhost:3306)/bbs_go?parseTime=true&loc=Local"
	}
	db, err := db.NewMySqlDb(dbString)
	if err != nil {
		panic(err)
	}

	// log
	logLevel := os.Getenv("BBSGO_LOG_LEVEL")
	if logLevel == "" {
		logLevel = "debug"
	}
	err = log.Init(log.Conf{LogLevel: logLevel})
	if err != nil {
		panic(err)
	}

	dbMap = model.Init(db, log.Logger)

	os.Exit(m.Run())
}
