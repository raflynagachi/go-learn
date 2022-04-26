package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return &DatabasePostgreSQL{
		Name: "PostgreSQL",
	}
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return &DatabaseMongoDB{
		Name: "MongoDB",
	}
}

type DatabaseRepository struct {
	databasePostgreSQL *DatabasePostgreSQL
	databaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		databasePostgreSQL: postgreSQL,
		databaseMongoDB:    mongoDB,
	}
}
