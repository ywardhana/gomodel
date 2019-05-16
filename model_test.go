package gomodel_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql" // for sqlx
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
	"github.com/ywardhana/gomodel"
)

var createTableTest = "CREATE TABLE IF NOT EXISTS " + dbName + " (" +
	"`id` INT(11) NOT NULL," +
	"`value` VARCHAR(20) NOT NULL," +
	"PRIMARY KEY (`id`)" +
	") ENGINE = InnoDB DEFAULT CHARSET = utf8"

var idTest = 1
var valTest = "value1"
var dbName = "test"
var insertTableTest = "INSERT INTO " + dbName + " VALUES(" + strconv.Itoa(idTest) + ",\"" + valTest + "\")"
var dropTable = "DROP TABLE " + dbName

func TestFind(t *testing.T) {
	db := newMysqlClient()
	db.MustExec(createTableTest)
	db.MustExec(insertTableTest)
	gomodel.SetDb(db)
	entity := Test{}
	model := gomodel.NewModel(dbName, &entity)
	model.Find(1)
	assert.Equal(t, idTest, entity.ID)
	assert.Equal(t, valTest, entity.Value)
	db.MustExec(dropTable)
}

func TestExec(t *testing.T) {
	db := newMysqlClient()
	db.MustExec(createTableTest)
	db.MustExec(insertTableTest)
	gomodel.SetDb(db)
	entity := Test{}
	model := gomodel.NewModel(dbName, &entity)
	res, err := model.Where("id = ? ", 1).Exec()
	if err != nil {
		assert.Nil(t, err)
		return
	}
	entity = *(res[0].(*Test))
	assert.Equal(t, idTest, entity.ID)
	assert.Equal(t, valTest, entity.Value)
	db.MustExec(dropTable)
}

type Test struct {
	ID    int    `db:"id"`
	Value string `db:"value"`
}

func newMysqlClient() *sqlx.DB {
	gotenv.Load()
	dbPort := os.Getenv("DATABASE_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbPoolSize, _ := strconv.Atoi(os.Getenv("DATABASE_POOL"))
	if dbPoolSize == 0 {
		dbPoolSize = 50
	}

	dataSourceName := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@(" + os.Getenv("DATABASE_HOST") + ":" + (string(dbPort)) + ")/" + os.Getenv("DATABASE_NAME") + "?parseTime=true"
	db, err := sqlx.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 60)
	db.SetMaxIdleConns(dbPoolSize)
	db.SetMaxOpenConns(dbPoolSize)

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}
