package main //similar to a namespace like in c#/ package in java (moreso like java)

//import the user package from the user folder
//also import the gofiber package
import (
	"main/user"

	"github.com/gofiber/fiber/v2"
)

//This function is retrieved when the website is at the root of the
//website's directory
//c *fiber.Ctx is the context of the routing/HTTP request, as you can see
// with c.Context().String()
func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Jose's website\n" + c.Context().String())
}

//function to specify routes on website
//could add the hello function in here by moving it to the users.go file
//and adding to the Routers function: app.GET("/", user.hello)
func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)         //will create a GET request and grab all user JSONS
	app.Get("/user/:id", user.GetUser)       //will create a GET request and grab a specific user based on their id
	app.Post("/user", user.SaveUser)         //will create a POST request where a new user can be created/saved
	app.Delete("/user/:id", user.DeleteUser) //will create a DELETE request where a specified user is hidden from /users by adding the time it was deleted in the DB
	app.Put("/user/:id", user.UpdateUser)    //PUT request will update user information
}

func main() {
	user.InitialMigration() //opens connection to database + creates the datatable if one is not already there
	app := fiber.New()      //opens new instance of fiber, which allows use create our routes, and implement different requests for specific routes GET,POST,PUT,etc.
	app.Get("/", hello)     //simple get request that calls the hello function, which returns a string when accessing http://localhost:3000/
	Routers(app)            // all the other routes which are: http://localhost:3000/route
	app.Listen(":3000")     //listens for http requests at port number 3000
}
