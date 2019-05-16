package gomodel

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Model struct {
	tableName string
	modelDb   *sqlx.DB
	query     Query
	entity    interface{}
}

type Query struct {
	whereClause string
	whereParams []interface{}
	selectAttr  string
	limiter     string
}

func NewModel(tableNameParam string, entity interface{}) Model {
	var temp []interface{}
	return Model{
		tableName: tableNameParam,
		query: Query{
			whereClause: "",
			whereParams: append(temp, ""),
			selectAttr:  "*",
			limiter:     "",
		},
		modelDb: db,
		entity:  entity,
	}
}

func SetDb(dbParam *sqlx.DB) {
	db = dbParam
}

func (m *Model) GetEntity() interface{} {
	return m.entity
}

func (m *Model) Find(ID int64) error {
	m.Where("id = ?", ID)
	err := m.modelDb.QueryRowx(m.getFullQuery(), m.query.whereParams...).StructScan(m.entity)
	return err
}

func (m *Model) Exec() ([]interface{}, error) {
	rows, err := m.modelDb.Queryx(m.getFullQuery(), m.query.whereParams...)
	res := []interface{}{}
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var elem = reflect.New(reflect.ValueOf(m.entity).Elem().Type()).Interface()

		if err := rows.StructScan(elem); err != nil {
			return res, err
		}
		res = append(res, elem)
	}
	return res, err
}

func (m *Model) Count() (int, error) {
	rows, err := m.modelDb.Queryx(m.getCountQuery(), m.query.whereParams...)
	res := 0
	if err != nil {
		return res, err
	}
	for rows.Next() {
		rows.Scan(&res)
	}
	return res, err
}

func (m *Model) Where(query string, args ...interface{}) *Model {
	m.query.whereClause = "where"
	m.query.whereClause += " " + query
	m.query.whereParams = args
	return m
}

func (m *Model) Select(attributes ...string) *Model {
	if len(attributes) == 0 {
		return m
	}
	res := ""
	for _, t := range attributes {
		res += t + " "
	}
	m.query.selectAttr = strings.Trim(res, " ")

	return m
}

func (m *Model) Limit(limiter int) *Model {
	if limiter < 1 {
		return m
	}
	m.query.limiter = strconv.Itoa(limiter)
	return m
}

func (m *Model) getFullQuery() string {
	return "select " + m.query.selectAttr + m.getSelectSource()
}

func (m *Model) getCountQuery() string {
	return "select count(*)" + m.getSelectSource()
}

func (m *Model) getSelectSource() string {
	res := " from " + m.tableName + " " + m.query.whereClause
	if m.query.limiter != "" {
		res += " limit " + m.query.limiter
	}
	return res
}
