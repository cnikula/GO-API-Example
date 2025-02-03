# GO-API-Example
An example of an API written in GO 

This Go REST API provides the following endpoints:

1. GET /products - Retrieve a list of all products
2. GET /products/:id - Retrieve a product by its ProductID
3. DELETE /products/:id - Delete a product by its ProductID
4. PUT /products/:id - Update a product by its ProductID

To run this API:

1. Make sure you have Go installed on your system.
2. Install the required dependencies:

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql


3. Update the database connection string in `main.go` with your Northwind database credentials.
4. Run the application:

The API will start running on `http://localhost:8080`.
