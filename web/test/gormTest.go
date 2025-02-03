package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Id    int
	Code  string    `gorm:"size:100;default:'0'"`
	Price int       `gorm:"not null`
	Join  time.Time `gorm:"type:timedata`
}

type Student struct {
	gorm.Model
	Name string
	Age  int
}

var Students []Student

func main() {
	dbconf := "root:Aa123456@tcp(127.0.0.1:3306)/test?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dbconf), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//创建数据库的表
	//db.AutoMigrate(&Product{})

	// 增
	//db.Create(&Product{Id: 1, Code: "D42", Price: 100})
	//db.Create(&Product{Code: "C22", Price: 666})
	//Save

	//查
	//1.First Last Find
	//2.Select Where
	//var product Product
	//var productC22 Product
	//db.First(&product, 1) //查找主键 1
	//db.Last(&product)
	//var products []Product
	//db.Find(&products)  select * from products;
	//db.First(&productC22, "Code = ?", "C22")
	//db.First(&product, "code = ?", "D42")
	//fmt.Println(product)

	//改
	//db.Model(&product).Where("id = ?", "2").Update("Price", 123)
	//db.First(&product, "price = ?", 123)
	//fmt.Println(product)

	//软删除
	//db.AutoMigrate(&Student{})
	//db.Create(&Student{Name: "zlican3", Age: 18})
	//软查询
	//db.Unscoped().Find(&Students)
	//fmt.Println(Students)
	//db.Where("name = ?", "zlican").Delete(new(Student))

}
