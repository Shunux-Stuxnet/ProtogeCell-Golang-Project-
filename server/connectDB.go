package server

import (
	"fmt"
	"log"
	"strconv"

	"ProtogeCell/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("golang.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	err = DB.AutoMigrate(&models.User{}, &models.IMEI{}, &models.Device{})
	if err != nil {
		log.Fatal("Failed to migrate database schema")
	}

	fmt.Println("Connected successfully to the database")
}

func InsertData(data models.User) {
	val := DB.Create(data)
	fmt.Println(val)
}

func InsertIMEI(imei models.IMEI) {
	val := DB.Create(imei)
	fmt.Println(val)
}

func InsertDevice(device models.Device) {
	val := DB.Create(device)
	fmt.Println(val)
}

func Query(c *fiber.Ctx) error {
	imeinum, errr := strconv.Atoi(c.FormValue("imei"))
	if errr != nil {
		return c.Render("views/index2.html", fiber.Map{"hel": "qwer"})
	}
	var imeidata models.IMEI
	var devicedata models.Device

	if err := DB.Where("imei = ?", imeinum).First(&devicedata).Error; err != nil {
		if err := DB.Where("imei = ?", imeinum).First(&imeidata).Error; err != nil {
			DB.Create(&models.Device{IMEI: imeinum, Count: 0})
			return c.Render("views/qres.html", fiber.Map{"imeinum": imeinum, "count": devicedata.Count})
		} else {
			DB.Create(&models.Device{IMEI: imeinum, Count: 1})
			return c.Render("views/qres.html", fiber.Map{"imeinum": imeinum, "count": devicedata.Count})
		}
	} else {
		DB.Model(&models.Device{}).Where("IMEI = ?", imeinum).Update("count", gorm.Expr("count + ?", 1))
		return c.Render("views/qres.html", fiber.Map{"imeinum": imeinum, "count": devicedata.Count})
	}
}

func ReportIMEI(c *fiber.Ctx) error {
	num := c.FormValue("imeinum")
	imei, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return c.SendString("invalid imei")
	}
	name := c.FormValue("name")
	email := c.FormValue("email")
	fir := c.FormValue("fir")
	locaion := c.FormValue("location")
	info := c.FormValue("info")
	mobile := c.FormValue("mobile")

	var devicedata models.Device
	if err := DB.Where("imei = ?", num).First(&devicedata).Error; err != nil {
		fmt.Println("1")
		user := models.IMEI{IMEI: imei, Name: name, Email: email, Mobile: mobile, Location: locaion, FIR: fir, Info: info}
		DB.Create(&user)
		fmt.Println("2")
	} else {
		fmt.Println("3")
		//DB.Model(&models.Device{}).Where("IMEI = ?", num).Update("count", gorm.Expr("count + ?", 1))
		user := models.IMEI{IMEI: imei, Name: name, Email: email, Mobile: mobile, Location: locaion, FIR: fir, Info: info}
		DB.Create(&user)
		fmt.Println("4")
	}
	return c.SendString("Your data has been saved, you can check count value if it was queried by anyone.")

}
