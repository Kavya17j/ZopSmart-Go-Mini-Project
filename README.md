# ZopSmart-Go-Mini-Project
Kavya Jain
Enrollment : 9920103154
Jaypee Institute Of Information Technology

The Go project is a simple CRUD (Create, Read, Update, Delete) API for managing movie data. It includes functionality to interact with a MySQL database for storing and retrieving movie information. The project is structured using the Gorilla Mux router and leverages the standard `database/sql` package for MySQL connectivity.

### Project Components:

1. **Main Go File (`main.go`):**
   - This file is the entry point of the application.
   - It initializes the Gorilla Mux router, defines routes for handling different HTTP methods (GET, POST, PUT, DELETE), and starts the HTTP server.
   - The main function initializes a slice of `Movie` structs and populates it with initial movie data.
   - It also establishes a connection to the MySQL database named `MoviesDB`.

2. **Movie Struct:**
   - Represents the structure of a movie with attributes like `ID`, `Isbn`, `Title`, and a nested `Director` struct.
   - The `Director` struct contains `Firstname` and `Lastname` attributes.

3. **Director Struct:**
   - Represents the director of a movie with attributes `Firstname` and `Lastname`.

4. **API Endpoints:**
   - `/movies` (GET): Retrieves a list of all movies.
   - `/movies/{id}` (GET): Retrieves details of a specific movie by ID.
   - `/movies` (POST): Creates a new movie.
   - `/movies/{id}` (PUT): Updates an existing movie.
   - `/movies/{id}` (DELETE): Deletes a movie by ID.

5. **HTTP Request Handlers:**
   - `getMovies`: Handles GET request for retrieving all movies.
   - `getMovie`: Handles GET request for retrieving a specific movie by ID.
   - `createMovie`: Handles POST request for creating a new movie.
   - `updateMovie`: Handles PUT request for updating an existing movie.
   - `deleteMovie`: Handles DELETE request for deleting a movie by ID.

6. **MySQL Database Connectivity:**
   - The application establishes a connection to a MySQL database named `MoviesDB`.
   - The database connection details are specified in the `sql.Open` function call in the `main` function.
   - The `Movies` table is assumed to exist in the `MoviesDB` database, storing movie data.

7. **Random ID Generation:**
   - The `createMovie` function generates a random ID for a new movie using `rand.Intn(100000000)`.

8. **Error Handling:**
   - Basic error handling is implemented, and error responses are returned in case of issues during JSON decoding, database operations, etc.

### Usage:

1. The API can be interacted with using HTTP requests to the defined endpoints.
2. Movies are stored in memory during the runtime and are not persistent across application restarts.
3. The MySQL database (`MoviesDB`) is assumed to be available for storing and retrieving movie data.

### Potential Improvements:

1. **Persistence:**
   - Implement a more robust persistence layer, such as using an ORM library or a database migration tool, to make movie data persistent.

2. **Validation:**
   - Implement input validation for API requests to ensure the integrity of data.

3. **Security:**
   - Enhance security measures, especially when dealing with database connections and user input.

4. **Logging:**
   - Introduce comprehensive logging to monitor the application's behavior and troubleshoot issues.

5. **Testing:**
   - Expand test coverage to ensure the reliability and correctness of the API's functionality.

This project serves as a starting point for building a more feature-rich and scalable movie management API in Go.




### How to Run the Movie Management API:

**Prerequisites:**
- Ensure you have Go installed on your system. If not, you can download and install it from https://github.com/Kavya17j/ZopSmart-Go-Mini-Project/.

**Steps:**

1. **Clone the Repository:**
   - Clone the repository containing the Go code for the Movie Management API.

 

2. **Install Dependencies:**
   - If your project has any external dependencies, use the following command to install them.

   ```bash
   go get ./...
   ```

3. **Build and Run:**
   - Build and run the `main.go` file.

   ```bash
   go build main.go
   ./main
   ```

   This will start the API server, and it will be accessible at `http://localhost:8000`.

4. **Database Setup:**
   - Make sure you have MySQL installed on your system.
   - Create a MySQL database named `MovieDB`.

   ```sql
   CREATE DATABASE MovieDB;
   ```

   - Ensure your MySQL server is running on `localhost` at the default port (`3306`).

5. **Connect to Database:**
   - Update the database connection details in the `main.go` file if needed.

   ```go
   con, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/MovieDB")
   ```

   Update `root` and `password` with your MySQL username and password.

6. **Testing with Postman:**
   - Install [Postman](https://www.postman.com/) to test the API endpoints.
   - Use Postman to send HTTP requests to the API at `http://localhost:8000/movies` for testing CRUD operations.

### Note:
- This guide assumes you have MySQL installed and running locally.
- Ensure that the required MySQL driver is available (`github.com/go-sql-driver/mysql`).
- Customize the database connection details and API routes based on your setup.

Thank you !!!!!!
