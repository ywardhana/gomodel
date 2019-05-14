package temp1

import (
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// type Model struct {
// 	Find (ID int64, *interface{})
// }

var db *sqlx.DB

func main() {
	// m := NewModel("coba",)
	// m.Where("id = ? and idx = ? and name = ?", 5, 10, "haha")
	// m.Find(1, &Query{
	// 	whereClause: "",
	// 	selectAttr:  "",
	// })
}

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

func (m *Model) Find(ID int64) {
	m.Where("id = ?", ID)
	// a := Test{}
	res := m.modelDb.QueryRowx(m.getFullQuery(), m.query.whereParams...).StructScan(m.entity)
	log.Println(m.entity)
	log.Println(res)
	log.Println(reflect.TypeOf(m.entity))
	log.Println(m.getFullQuery())
}

type Test struct {
	ID    int    `db:"id"`
	Value string `db:"value"`
}

// Where(att = ?, val)
func (m *Model) Where(query string, args ...interface{}) *Model {
	m.query.whereClause = "where"
	// for _, t := range parsers {
	// 	var replacer string
	// 	switch t.(type) {
	// 	case string:
	// 		replacer = t.(string)
	// 	case int:
	// 		replacer = strconv.Itoa(t.(int))
	// 	case int64:
	// 		replacer = strconv.FormatInt(t.(int64), 10)
	// 	default:
	// 		panic(errors.New("Unsupported Type"))
	// 	}
	// 	index := strings.Index(query, "?")
	// 	query = query[:index] + replacer + query[index+1:]
	// }
	// m.query.whereClause = m.query.whereClause + " " + query
	log.Println(args...)
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
	res := "select " + m.query.selectAttr + " from " + m.tableName + " " + m.query.whereClause
	if m.query.limiter != "" {
		res += " limit " + m.query.limiter
	}
	return res
}
