package bootstrap

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewPostgresDatabase(env *Env) *gorm.DB {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		env.DbHost,
		env.DbUser,
		env.DbPass,
		env.DbName,
		env.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	// test := db.WithContext(ctx)
	// fmt.Println(*test)

	return db
}

func CloseConnection(db *gorm.DB) {
	if db == nil {
		return
	}

	if db != nil {
		db, _ := db.DB()
		db.Close()
	}
}
