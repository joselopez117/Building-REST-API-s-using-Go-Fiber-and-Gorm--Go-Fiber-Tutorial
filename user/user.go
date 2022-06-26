package user //similar to a namespace in c#/package in java (moreso like java)

//fmt for strings/print(f/ln) stuff

//gofiber be gofiber

//mysql for database access

//gorm is to allow for object relation mapping,
//in other words allows programming language to more seamlessly
//add data to database, without having user implement all the database logic
import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB //databse object
var err error   //for error messages

//DNS address is the address to the database
const DNS = "root:j053$3rv3r@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

//remember the attributes are capital letters
//to be able to export these to other packages

//struct that will be the object to send to the database
type User struct {
	gorm.Model //adds the primary key of our database(ID) and includes some common attributes
	//for adding objects to database
	//the attributes users will interact with generally
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

//creates the database that will contain our user data
func InitialMigration() {
	//gorm.open opens the connection to database
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	//error message if connection to DB fails
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	//will send the schema that was declared in user struct
	//to the table if a table is not already created
	DB.AutoMigrate(&User{})
}

//called when going to "/users"
//creates array(slice in golang) to hold all the users JSONs
//and then returns them
func GetUsers(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)
	return c.JSON(&users)
}

//called when going to "/user/:id"
//
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.Find(&user, id)
	return c.JSON(&user)
}

//called when going to "/user/:id"
func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Create(&user)
	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}

	DB.Delete(&user)
	return c.SendString("User is deleted!")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User not available")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&user)
	return c.JSON(&user)
}
