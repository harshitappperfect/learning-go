package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // PostgreSQL driver
)

// User represents a user in the database
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Database connection variable
var db *gorm.DB
var err error

// Initialize the database connection
func init() {
	// Database connection string (update this with your credentials)
	dsn := "host=localhost user=postgres dbname=go_crud password=your_password sslmode=disable"
	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	fmt.Println("Database connected successfully!")
}

//Explantion of the above lines

// Let's break down this code snippet line by line and explain it in simple terms.

// ### Code Snippet:
// ```go
// // Database connection variable
// var db *gorm.DB
// var err error
// ```

// ### 1. **Declaring Database Variables**
// ```go
// var db *gorm.DB
// var err error
// ```
// - `var db *gorm.DB`:
//   This line declares a variable `db` that will hold the connection to the database using the **GORM** package. The `*gorm.DB` type means that `db` will point to an instance of the `gorm.DB` struct, which is the main object used to interact with the database.

// - `var err error`:
//   This line declares a variable `err` of type `error`. This will hold any error messages that might occur during database operations, such as connecting to the database or querying it.

// ---

// ### 2. **Initializing the Database Connection in the `init` Function**
// ```go
// func init() {
//     // Database connection string (update this with your credentials)
//     dsn := "host=localhost user=postgres dbname=go_crud password=your_password sslmode=disable"
// ```

// - `func init()`:
//   This is the `init` function in Go. The `init` function is special because it is automatically executed when the program starts, even before the `main` function. It's commonly used for setup tasks, such as database connections, loading configuration files, or initializing global variables.

// - `dsn := "host=localhost user=postgres dbname=go_crud password=your_password sslmode=disable"`:
//   This line creates a string `dsn` that contains the **Data Source Name (DSN)** for connecting to the PostgreSQL database. This string provides all the information required to establish the connection:
//   - `host=localhost`: This specifies that the PostgreSQL server is running on your local machine.
//   - `user=postgres`: This is the username for connecting to the database. Here, it's set to `postgres`, which is the default PostgreSQL username.
//   - `dbname=go_crud`: The name of the database you want to connect to. In this case, it’s `go_crud`.
//   - `password=your_password`: This is the password for the PostgreSQL user (`postgres` in this case). You should replace `your_password` with the actual password you set during PostgreSQL installation.
//   - `sslmode=disable`: This disables SSL encryption for the connection. It’s common to use this when working in development environments where SSL is not required.

// ---

// ### 3. **Opening the Database Connection**
// ```go
// db, err = gorm.Open("postgres", dsn)
// ```

// - `db, err = gorm.Open("postgres", dsn)`:
//   This line tries to open a connection to the PostgreSQL database using the `gorm.Open()` function. Here's how it works:
//   - `"postgres"`: This specifies that we are using PostgreSQL as the database type (since we are using GORM’s PostgreSQL dialect).
//   - `dsn`: This is the connection string we defined earlier. It provides all the necessary details to establish the connection.
//   - `gorm.Open("postgres", dsn)`: The `gorm.Open()` function returns two values:
//     - `db`: This is the **GORM database connection** object that you’ll use to interact with the database.
//     - `err`: If there is an error during the connection (for example, if the database is not available or the credentials are wrong), it will be captured here.

// ---

// ### 4. **Handling Connection Errors**
// ```go
// if err != nil {
//     log.Fatalf("Error connecting to the database: %v", err)
// }
// ```

// - `if err != nil`:
//   This checks if there was an error while trying to open the connection to the database. In Go, `err` is set to `nil` if there was no error, and a non-`nil` value (error message) is returned if there was an error.

// - `log.Fatalf("Error connecting to the database: %v", err)`:
//   If there is an error (`err != nil`), this line logs the error message to the console using `log.Fatalf()`. The `%v` is a placeholder for printing the actual error message that was returned.
//   - `log.Fatalf()` not only prints the error message but also stops the program execution. This is useful because if the database connection fails, continuing the program wouldn't make sense.

// ---

// ### 5. **Automatic Database Migration**
// ```go
// db.AutoMigrate(&User{})
// ```

// - `db.AutoMigrate(&User{})`:
//   This line tells GORM to automatically migrate the `User` struct to the database. In simple terms, it:
//   - **Creates** the necessary table in the database if it doesn’t exist.
//   - **Updates** the table schema if changes were made to the struct (e.g., adding a new field).
//   - **Ensures** the database schema is consistent with the `User` struct, which represents the structure of a user in the database.

//   Here’s the breakdown of the arguments:
//   - `&User{}`: This is a pointer to the `User` struct. GORM uses this struct to figure out what columns should exist in the database table. For example, if your struct has `Name`, `Email`, and `Age` fields, GORM will create corresponding columns in the `users` table.

//   **Note**: `AutoMigrate` is useful for development and prototyping, but in production, it is better to handle schema migrations manually to avoid unintended changes.

// ---

// ### 6. **Print Confirmation**
// ```go
// fmt.Println("Database connected successfully!")
// ```

// - `fmt.Println("Database connected successfully!")`:
//   This line prints a confirmation message to the console to let the developer know that the database connection was successful. It’s useful for debugging and ensuring everything is working properly when the program starts.

// ---

// ### **Summary of the `init` function**

// 1. **Declare database connection variables**: `db` for the GORM database object and `err` for storing any errors.
// 2. **Create a connection string (DSN)**: The `dsn` string holds the necessary details (host, user, password, etc.) to connect to the PostgreSQL database.
// 3. **Open the database connection**: `gorm.Open()` attempts to connect to the database. If there’s an error, the program stops with a message.
// 4. **Migrate the schema**: `db.AutoMigrate()` ensures the database schema matches the `User` struct by creating/updating the table automatically.
// 5. **Print success message**: Once the database is connected, a success message is printed.

// ---

// This process sets up the database connection for your Go application using GORM and PostgreSQL. The `init()` function is executed automatically before the main program starts, ensuring that the database is ready for use when the application begins handling API requests.

// Create a new user (POST /users)
func createUser(w http.ResponseWriter, r *http.Request) { //w (ResponseWriter): This is used to send the response back to the client.
	// r (Request): This holds the incoming HTTP request that contains the user data we want to save.
	var user User
	err := json.NewDecoder(r.Body).Decode(&user) // Decode the incoming JSON request body:
	// The client sends a JSON body in the request, which contains the user’s data. We need to decode this JSON into a Go struct (user) so we can work with it..

	// The json.NewDecoder(r.Body) creates a JSON decoder that reads the body of the incoming request (r.Body), and the .Decode(&user) method converts the JSON into our user struct.
	// example:

	// JSON to Go struct:
	// user := User{
	// 	Name:  "John Doe",
	// 	Email: "johndoe@example.com",
	// 	Age:   28,
	// }

	if err != nil {
		// If there’s an error while decoding the JSON (e.g., the client sends invalid JSON or the data doesn't match the expected format), we send a response back to the client that something is wrong.

		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Save the user to the database
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// if err := db.Create(&user).Error; err != nil {
	// 	Explanation: This line is using Go's short variable declaration (:=) to declare and assign the variable err. It attempts to create a new user in the database and checks if an error occurred during that process.

	// 	Breakdown:

	// 	db.Create(&user): This is a call to GORM’s Create method, which is used to insert a new record into the database. Here, db is the database connection object (a *gorm.DB), and user is a pointer to the struct that represents the user (likely a model in your application).
	// 	&user: The & is used to pass the address (pointer) of the user object to Create, as GORM expects a pointer to the model when inserting data.
	// 	.Error: GORM provides a Error field in the result of database operations to represent any errors that occur. This field will be nil if the operation is successful and will contain the error message if the operation fails.
	// 	if err != nil: This checks if the Error is not nil, meaning that an error occurred during the database operation.

	// In this example, if there's an error while creating the user in the database (like a constraint violation), the err will contain that error, and the code inside the if block will be executed.

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

	// 	w.WriteHeader(http.StatusCreated): This sends a 201 Created status code to the client. This tells the client that the new user has been successfully created.

	// json.NewEncoder(w).Encode(user): This converts the user struct into a JSON format and sends it back to the client. The client receives the newly created user’s details in the response.
}

// Get all users (GET /users)
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	// This line declares a variable users, which will hold a slice (a dynamically-sized array) of User objects. The User type is assumed to be a struct that represents a user, typically containing fields like ID, Name, Email, etc.
	if err := db.Find(&users).Error; err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//This line sets the Content-Type header of the HTTP response to application/json. It tells the client that the response body will contain JSON data.
	json.NewEncoder(w).Encode(users)

	// 	Putting It All Together:
	// This function is part of a web handler that:

	// Fetches all users from the database.
	// If successful, returns the users in JSON format with a 200 OK status.
	// If there is an error fetching users (e.g., database issue), it returns an HTTP error with status 500 Internal Server Error.
}

// Get a user by ID (GET /users/{id})
func getUser(w http.ResponseWriter, r *http.Request) {

	//mux.Vars(r) is a method from the gorilla/mux router, which extracts the path variables from the incoming request.
	// mux.Vars(r) returns a map of key-value pairs where each key is the name of a path variable and each value is the corresponding value in the URL.
	// Example: If the URL path is /users/{id}, like /users/1, mux.Vars(r) will return a map like:
	// params := map[string]string{
	// 	"id": "1",
	// }

	// This means the variable id in the path is 1.

	params := mux.Vars(r)
	var user User
	id := params["id"] // id will be string: "1"

	if err := db.First(&user, id).Error; err != nil { //query to find the first record that matches the provided id. If a user with that ID exists, it is loaded into the user variable.
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

	//json.NewEncoder(w) creates a new JSON encoder that will write the JSON data to the http.ResponseWriter (i.e., the response that will be sent to the client).
	// .Encode(user) serializes the user struct into a JSON string and writes it to the response.
}

// Update a user (PUT /users/{id})
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	id := params["id"]

	// Find the existing user
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the updated user data from the request
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Update the user in the database
	if err := db.Save(&user).Error; err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Delete a user (DELETE /users/{id})
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	id := params["id"]

	// Find the user to delete
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Delete the user from the database
	if err := db.Delete(&user).Error; err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// Set up routes
	r := mux.NewRouter()

	// Define routes and their corresponding handlers
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Start the server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
