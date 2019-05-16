# Gomodel
Simple golang model-orm like- module. Implemented for Mysql adapter, using sqlx(`github.com/jmoiron/sqlx`) library.

## Recent Changes
Features:
* Where method
* Find method
* Limit method
* Count method

## Installation
```go get github.com/ywardhana/gomodel```

## Usage
Some examples are available below

### Initialization
```go
db := newMysqlClient() // Create new mysql client using sqlx
gomodel.SetDb(db)
model := gomodel.NewModel(yourTableName, &yourStruct)
```

### Find based on ID
```go
///// initialization step /////
model.Find(yourDesiredID)

/// the result is available at `yourStruct`

```

### Where
```go
///// initialization step /////
model.Where("id = ?", yourDesiredParameter)

/// multiparameter is supported, so you can use where such as model.Where("id = ? and column1 = ?", param1, param2)

/// if you want the result, execute your query with `Exec()` method after `Where`

```

### Limit
```go
///// initialization step /////
model.Where("column1 = ?", yourDesiredParameter).Limit(yourDesiredLimit)

/// if you want the result, execute your query with `Exec()` method after `Limit`

```

### Count
```go
///// initialization step /////
countRes, err := model.Where("column1 = ?", yourDesiredParameter).count

/// the countRes is in integer

```
