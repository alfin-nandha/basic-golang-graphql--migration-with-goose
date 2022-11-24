package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	godotenv.Load(".env")
}

func DBConn() *gorm.DB {
	config := map[string]interface{}{
		"DB_Username": GetEnv("DB_USERNAME", "root"),
		"DB_Password": GetEnv("DB_PASSWORD", "yasfin123"),
		"DB_Port":     GetEnv("DB_PORT", "3306"),
		"DB_Host":     GetEnv("DB_HOST", "127.0.0.1"),
		"DB_Name":     GetEnv("DB_NAME", "basicGraphql"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Printf("Error Db Connection %s", err.Error())
		return nil
	}
	return db
}

func GetEnv(key, defValue string) string {
	val, exist := os.LookupEnv(key)
	if !exist {
		return defValue
	}
	return val
}
