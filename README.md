# Gomodel
Simple golang model-orm like- module. Implemented for Mysql adapter, using sqlx(`github.com/jmoiron/sqlx`) library.

## Recent Changes
### v0.0.1
Features:
* Where method
* Find method
* Limit method

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

/// the result is available at `yourStruct`

```

### Limit
```go
///// initialization step /////
model.Where("column1 = ?", yourDesiredParameter).Limit(yourDesiredLimit)

/// the result is available at `yourStruct`

```

