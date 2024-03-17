package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// dsn := "username:password@tcp(localhost:port)/databasename?parseTime=True"
	dsn := ("DATABASE_URL")

	dial := mysql.Open(dsn)
	var err error
	db, err = gorm.Open(dial)
	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(Gender{}, Test{})
	// CreateCustomer("Nook", 2)
	// RawQueryGetCustomer()
}

// Get all
func GetGenders() {
	genders := []Gender{}
	tx := db.Find(&genders)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(genders)
}

// Get one by id
func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

// Get one by name
func GetGenderByName(name string) {
	gender := Gender{}
	tx := db.Where("name=?", name).Find(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

// Update
func UpdateGender(id uint, name string) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

// Update2
func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=?", id).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

// Delete for real
func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(id)
}

// Create
func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(gender)
}

// Create test
func CreateTest(code uint, name string) {
	test := Test{Code: code, Name: name}
	tx := db.Create(&test)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(test)
}

// Get all test
func GetTests() {
	tests := []Test{}
	tx := db.Find(&tests)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(tests)
}

// Delete test
// if use gorm.Model. it will soft delete
// if want to delete for real, use Unscoped => db.Unscoped().Delete(&Test{}, id)
func DeleteTest(id uint) {
	tx := db.Delete(&Test{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetTests()
}

// Create customer
func CreateCustomer(name string, genderID uint) {
	customer := Customer{Name: name, GenderID: genderID}
	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

// Get all customer and join table for find Gender
func GetCustomers() {
	customers := []Customer{}
	// Preload คือการ join table
	// another way is get all table before join => db.Prelasd(clause.Associations).Find(&customers)
	tx := db.Preload("Gender").Find(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, customer := range customers {
		fmt.Printf("%v/%v/%v/n", customer.ID, customer.Name, customer)

	}
}

// Gorm can use Raw or Exec
func RawQueryGetCustomer() {
	// สามารถใช้ struct เดิมหรือสร้าง struct ใหม่ก็ได้
	customers := []Customer{}
	tx := db.Raw(`SELECT customers.*, genders.Name as GenderName 
	FROM customers 
	LEFT JOIN genders ON customers.Gender_id = genders.ID`).Scan(&customers)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	for _, customer := range customers {
		fmt.Printf("%v/%v/%v/n", customer.ID, customer.Name, customer)

	}
}

type Gender struct {
	ID   uint
	Name string
}

// if use gorm.Model, it will create id, created_at, updated_at, deleted_at
// when delete, it will not delete, but set deleted_at
// soft delete
type Test struct {
	gorm.Model
	Code uint
	Name string
}

type Customer struct {
	ID   uint
	Name string
	//Gender อ้างถึงตาราง Gender
	Gender Gender
	// Foreign key reference gerder_id
	GenderID uint
}
