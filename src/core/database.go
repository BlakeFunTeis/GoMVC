package core

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "log"
    "os"
    "strconv"
    "time"
)

var (
    DB *gorm.DB
)

func GetInstance() {
    var err error
    DB, err = gorm.Open(os.Getenv("database_type"), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
        os.Getenv("mysql_user"),
        os.Getenv("mysql_password"),
        os.Getenv("mysql_host"),
        os.Getenv("mysql_name"),
    ))

    if err != nil {
       log.Fatal(err)
    }

    maxIdle, _ := strconv.Atoi(os.Getenv("mysql__max_idle_conns"))
    maxOpen, _ := strconv.Atoi(os.Getenv("mysql__max_idle_conns"))
    DB.DB().SetConnMaxLifetime(time.Minute*5)
    DB.DB().SetMaxIdleConns(maxIdle)
    DB.DB().SetMaxOpenConns(maxOpen)
}

func ChangeDatabase(database string)  {
    CloseDatabase()
    err := os.Setenv("mysql_name", database)
    if err == nil {
        GetInstance()
        return
    }

    log.Println("切換資料庫失敗")
}

func CloseDatabase() {
    _ = DB.Close()
}