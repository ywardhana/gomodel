package gomodel_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql" // for sqlx
	"github.com/stretchr/testify/assert"
	"github.com/ywardhana/gomodel"
)

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

func TestCount(t *testing.T) {
	db := newMysqlClient()
	db.MustExec(createTableTest)
	db.MustExec(insertTableTest)
	gomodel.SetDb(db)
	entity := Test{}
	model := gomodel.NewModel(dbName, &entity)
	res, err := model.Where("id = ? ", 1).Count()
	if err != nil {
		assert.Nil(t, err)
		return
	}

	assert.Equal(t, 1, res)
	db.MustExec(dropTable)
}
