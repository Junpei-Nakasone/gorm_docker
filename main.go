package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

// MySQL接続に必要な情報を定数で定義する場合の書き方
// const (
// 	// データベース
// 	Dialect = "mysql”
// 	// ユーザー名
// 	DBUser = "root"
// 	// パスワード
// 	DBPass = "root"
// 	// プロトコル
// 	DBProtocol = "tcp(127.0.0.1:3306)"
// 	// DB名
// 	DBName = "go_sample"
// )

// モデルはテーブルを構造体で表現したもの
// そのままDBテーブルになる
type User struct {
	// gorm.goで定義してるID他を構造体Userに注入
	Name  string
	Email string
	// カラム名を指定できる
	Age int    `gorm:"column:test_columnName"`
	Sex string `gorm:"size:255"`
}

// // MySQLとGORMを繋ぐ関数
func connectGorm() *gorm.DB {
	// 変数を使ってMySQLと接続する場合
	// connectTemplate := "%s:%s@%s/%s"
	// connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)

	// ハードコーディングで接続
	// tcp(127.0.0.1:3306)なしでも接続できたので改良できるかも
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gorm_practice")

	if err != nil {
		log.Println(err.Error())
	}

	return db
}

func allUsers(c echo.Context) error {
	// GORMとMySQLを繋いで、変数dbに格納
	db := connectGorm()
	// 常にdbはクローズする
	defer db.Close()

	var users []User
	db.Find(&users)

	return c.JSON(http.StatusOK, &users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	// GORMとMySQLを繋いで、変数dbbに格納
	db := connectGorm()
	// 常にdbはクローズする
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(email)

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// GORMとMySQLを繋いで、変数dbbに格納
	db := connectGorm()
	// 常にdbはクローズする
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	var user User
	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// GORMとMySQLを繋いで、変数dbbに格納
	db := connectGorm()
	// 常にdbはクローズする
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

func main() {
	// GORMとMySQLを繋いで、変数dbbに格納
	db := connectGorm()
	// 常にdbはクローズする
	defer db.Close()

	// テーブル名を単数形にする設定。AutoMigrateより後にあるとエラーになる
	// db.SingularTable(true)
	// テーブルが存在しない場合に対象のテーブルを作成する
	// テーブルなどの生成は行うが削除はできない
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})

	e := echo.New()
	e.GET("/users", allUsers)

	e.Start(":8080")
}
