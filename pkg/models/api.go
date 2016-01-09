package models

import (
	"fmt"
	"github.com/asyoume/Auth_ser/pkg/log"
	"github.com/asyoume/Auth_ser/pkg/pulic_type"
	db_driver "github.com/asyoume/postgres"
)

var DB *db_driver.DB

func Init(db, logser pulic_type.MicroSerType) {
	//初始化数据库
	sqlurl := db.Addr
	loger := log.KafakLoger{}
	DB = db_driver.NewDB()
	TypeDecorat(DB)
	err := DB.Open(sqlurl, &loger)
	if err != nil {
		fmt.Println("初始化postgresql数据库失败")
	} else {
		fmt.Println("初始化postgresql数据库成功")
	}

	/*test = Test1{Test2: "adadad22"}
	err = DB.Update("test1", "", &test, []string{"test1"})
	fmt.Println(err)

	test = Test1{}
	err = DB.One("test1", "", &test, []string{"test1", "test2"})
	fmt.Println(err)
	fmt.Println(test)

	tests := []Test1{}
	err = DB.All("test1", "", &tests, []string{"test1", "test2"}, "", 1, 7)
	fmt.Println(err)
	fmt.Println(tests)*/
}
