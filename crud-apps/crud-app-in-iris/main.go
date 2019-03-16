package main

import (
	"fmt"
	"time"
	"github.com/kataras/iris"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// Create database model for the application
type User struct {
	ID        bson.ObjectId  `bson:"_id,omitempty"`
	Firstname string         `json: "firstname"`
	Lastname  string         `json: "lastname"`
	Age       int            `json: "age"`
	Phonenumber    string         `json: "phonenumber"`
	InsertedAt time.Time     `json: "inserted_at" bson: "inserted_at"`
	LastUpdate time.Time     `json: "last_update" bson: "last_update"`
}

func main() {
	// Create server 
	app := iris.New()
	app.Logger().SetLevel("debug")
	// adding two built in handlers 
	// to recover from http-relative panics
	// and log requests in terminal
	app.Use(recover.New())
	app.Use(logger.New())

	session, err := mgo.Dial("127.0.0.1")
	if nil != err {
		// defer, panic, and recover.
		// Panic is a built-in function that stops the ordinary flow of control and begins panicking
		// Recover is a built-in function that regains control of a panicking goroutine.
		panic(err)
	}
	//  defer statement pushes a function call onto a list.
	// clean up actions 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("usergo").C("profiles")

	// Index
	index := mgo.Index{
		Key:   []string{"phonenumber"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}

	err = c.EnsureIndex(index)
	if err !=  nil {
		panic(err)
	}

	// create a default base url (GET Request)
	app.Handle("GET", "/", func(ctx context.Context) {
		ctx.JSON(context.Map{"message": "Welcome to first CRUD app in Go Lang using Iris"})
	})

	// Get all users
	app.Handle("GET", "/users", func(ctx context.Context) {
		userlists := []User{}

		err := c.Find(nil).All(&userlists)
		if err != nil {
			// Return the error
			ctx.JSON(context.Map{"error": err.Error()})
		} else {
			fmt.Println(userlists)
		}
		if len(userlists) < 1 {
			// TODO
			// Change status code to 404
			ctx.JSON(context.Map{"message": "no user available at this time", "status": 404})
		} else {
			ctx.JSON(context.Map{"users": userlists})
		}
	})

	// Get a single user by its phonenumber
	app.Handle("GET", "/users/{phonenumber: string}", func(ctx context.Context) {
		phonenumber := ctx.Params().Get("phonenumber")
		fmt.Println(phonenumber)
		if phonenumber == "" {
			ctx.JSON(context.Map{"message": "please pass a valid phonenumber"})
		}
		user := User{}
		err = c.Find(bson.M{"phonenumber": phonenumber}).One(&user)
		if err != nil {
			checkerr := err.Error()
			ctx.JSON(context.Map{"error": checkerr})
		}
		ctx.JSON(context.Map{"user": user})
	})

	// Create a new user method: POST
	app.Handle("POST", "/users", func(ctx context.Context) {
		params := &User{}
		err := ctx.ReadJSON(params)
		if err != nil {
			ctx.JSON(context.Map{"error": err.Error()})
		} else {
			params.LastUpdate = time.Now()
			err := c.Insert(params)
			if err != nil {
				ctx.JSON(context.Map{"error": err.Error()})
			}
		}
		ctx.JSON(context.Map{"user": params})
	})

	// Update user details method: PATCH
	app.Handle("PATCH", "/users/{phonenumber: string}", func(ctx context.Context) {
		phonenumber := ctx.Params().Get("phonenumber")
		if phonenumber == "" {
			ctx.JSON(context.Map{"error": "please provide a valid phonenumber"})
		}
		params := &User{}
		err := ctx.ReadJSON(params)
		if err != nil {
			ctx.JSON(context.Map{"error": err.Error()})
		} else {
			params.InsertedAt = time.Now()
			query := bson.M{"phonenumber": phonenumber}
			err = c.Update(query, params)
			if err != nil {
				ctx.JSON(context.Map{"error": err.Error()})
			} else {
				user := User{}
				err = c.Find(bson.M{"phonenumber": params.Phonenumber}).One(&user)
				if err != nil {
					ctx.JSON(context.Map{"error": err.Error()})
				}
				ctx.JSON(context.Map{"message": "user details successfully updated", "user": user})
			}
		}
	})

	// Delete a user functionality
	app.Handle("DELETE", "/users/{phonenumber: string}", func(ctx context.Context) {
		phonenumber := ctx.Params().Get("phonenumber")
		if phonenumber == "" {
			ctx.JSON(context.Map{"error": "please provide a valid phonenumber"})
		}
		params := &User{}
		err := ctx.ReadJSON(params)
		params.InsertedAt = time.Now()
		query := bson.M{"phonenumber": phonenumber}
		err = c.Remove(query)
		if err != nil {
			ctx.JSON(context.Map{"error": err.Error()})
		} else {
			ctx.JSON(context.Map{"message": "user details deleted successfully"})
			}
	})

	// Running the application in port 8080
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	
}
