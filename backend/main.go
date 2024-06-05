package main

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

const jwtSecret = "khomkrit"

func main() {

	var err error
	dsn := "root:OakXzVBqNvuHpqViugacwxyAHSQKJfKv@tcp(viaduct.proxy.rlwy.net:18566)/railway?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Middleware Authentication
	app := fiber.New()
	authentication := jwtware.New(jwtware.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte(jwtSecret),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.ErrUnauthorized
		},
	})
	// db.AutoMigrate(User{}, Book{})
	app.Post("/signup", Signup)
	app.Post("/login", Login)

	securedRoutes := app.Group("/", authentication)
	securedRoutes.Get("/getbooks", GetBooks)
	securedRoutes.Get("/getbook/:id", GetBook)
	securedRoutes.Post("/createbook", CreateBook)
	securedRoutes.Put("/updatebook/:id", Updatebook)
	securedRoutes.Delete("/deletebook/:id", DeleteBook)

	app.Listen(":8000")
}

func Signup(c *fiber.Ctx) error {
	request := SignupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}
	user := User{Username: request.Username, Password: string(password)}
	tx := db.Create(&user)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Can't create user")
	}
	response := User{
		Id:       user.Id,
		Username: user.Username,
		Password: string(password),
	}
	return c.Status(fiber.StatusCreated).JSON(response)
}
func Login(c *fiber.Ctx) error {
	request := LoginRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}
	user := User{}
	if err := db.Where("username = ?", request.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Database error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Incorrect username or password")
	}

	cliams := jwt.StandardClaims{
		Issuer:    strconv.Itoa((user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{
		"jwtToken": token,
	})
}
func GetBooks(c *fiber.Ctx) error {
	books := []Book{}
	tx := db.Order("id").Find(&books).Error
	if tx != nil {
		return fiber.NewError(fiber.StatusNotFound, "Failed to fetch books")
	}
	return c.Status(fiber.StatusOK).JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	book := Book{}
	tx := db.Where("id = ?", id).First(&book).Error
	if tx != nil {
		return fiber.NewError(fiber.StatusNotFound, "Failed to fetch books")
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func CreateBook(c *fiber.Ctx) error {
	request := BookRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}
	// Note
	//This code for decode jwt token will get user_id to create book but can't working because version of jwt, I will default 1
	// claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	// userID := int(claims["user_id"].(float64))

	book := Book{Title: request.Title, Author: request.Author, Genre: request.Genre, Year: request.Year, UserID: 1}
	tx := db.Create(&book)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Can't create user")
	}
	return c.Status(fiber.StatusCreated).JSON(book)
}

func Updatebook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	request := BookRequest{}
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	book := Book{}
	tx := db.First(&book, id)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Book not found")
	}
	book.Title = request.Title
	book.Author = request.Author
	book.Genre = request.Genre
	book.Year = request.Year
	tx = db.Save(&book)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Could not update book")
	}
	return c.Status(fiber.StatusOK).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.ErrBadRequest
	}
	tx := db.Delete(&Book{}, id)
	if tx.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Book not found")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

type User struct {
	gorm.Model
	Id       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type Book struct {
	gorm.Model
	Title  string `db:"title" json:"title"`
	Author string `db:"author" json:"author"`
	Genre  string `db:"genre" json:"genre"`
	Year   int    `db:"year" json:"year"`
	UserID int    `db:"user_id" json:"user_id"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
}
