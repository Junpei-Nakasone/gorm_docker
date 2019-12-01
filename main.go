package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
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
	gorm.Model
	// GORM`gorm:`でモデル宣言時に任意でタグを使用可能。Nameフィールドは
	// 文字列のサイズを255に指定されている。
	Name string `gorm:"size:255"`
	// カラム名を指定できる
	Age  int	`gorm:"column:test_columnName"`
	Sex  string `gorm:"size:255"`
}

func (u User) String() string {
	return fmt.Sprintf("%s(%d)", u.Name, u.Age)
}

// MySQLとGORMを繋ぐ関数
func connectGorm() *gorm.DB {
	// 変数を使ってMySQLと接続する場合
	// connectTemplate := "%s:%s@%s/%s"
	// connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName)

	// ハードコーディングで接続
	// tcp(127.0.0.1:3306)なしでも接続できたので改良できるかも
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_sample")

	if err != nil {
		log.Println(err.Error())
	}

	return db
}

// 構造体usersにデータを入れる関数
func insert(users []User, db *gorm.DB) {
	for _, user := range users {
		db.NewRecord(user)
		db.Create(&user)
	}
}

func main() {
	// GORMとMySQLを繋いで、変数dbbに格納
	db := connectGorm()
	// 常にdbはクローズする
	defer db.Close()

	// テーブル名を単数形にする設定。AutoMigrateより後にあるとエラーになる
	db.SingularTable(true)
	// テーブルが存在しない場合に対象のテーブルを作成する
	// テーブルなどの生成は行うが削除はできない
	db.Set("gorm:table_options", "ENGINE = InnoDB").AutoMigrate(&User{})

	// 以下、GORMが動いてるか確かめる処理
	user1 := User{Name: "yamada", Age: 25, Sex: "male"}
	user2 := User{Name: "tanaka", Age: 22, Sex: "felame"}
	insertUsers := []User{user1, user2}
	insert(insertUsers, db)

	
}
