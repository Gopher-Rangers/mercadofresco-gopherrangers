<p align="center">
<img src="https://github.com/ashleymcnamara/gophers/blob/master/TEACHING_GOPHER.png?raw=true" width="250"/>
</p>

<h1 align="center">Mercado Fresco Gopher Rangers</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/Gopher-Rangers/mercadofresco-gopherrangers?color=3de069">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/Gopher-Rangers/mercadofresco-gopherrangers?color=3de069">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/Gopher-Rangers/mercadofresco-gopherrangers?color=3de069">

  <img alt="License" src="https://img.shields.io/github/license/Gopher-Rangers/mercadofresco-gopherrangers?color=3de069">

  <img alt="Github stars" src="https://img.shields.io/github/stars/Gopher-Rangers/mercadofresco-gopherrangers?color=3de069" />

  <a href="https://github.com/Gopher-Rangers/mercadofresco-gopherrangers/actions/workflows/test.yml">
    <img src="https://github.com/Gopher-Rangers/mercadofresco-gopherrangers/actions/workflows/test.yml/badge.svg">
  </a>

  <a href="https://codecov.io/gh/Gopher-Rangers/mercadofresco-gopherrangers"> 
    <img src="https://codecov.io/gh/Gopher-Rangers/mercadofresco-gopherrangers/branch/main/graph/badge.svg?token=NUUR12FFLR"> 
  </a>

</p>

<h4 align="center"> 
	Mercado Fresco Gopher Rangers 🚀 Under construction...
</h4> 

<hr>

<p align="center">
  <a href="#about">About</a> &#xa0; | &#xa0; 
  <a href="#features">Features</a> &#xa0; | &#xa0;
  <a href="#technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#starting">Starting</a> &#xa0; | &#xa0;
  <a href="#license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/Gopher-Rangers" target="_blank">Author</a>
</p>

<br>

## About ##

This API Handles Mercado Fresco Sellers, Warehouses, Sections, Products, Employees and Buyers

It was made for the Mercado Livre's Go Bootcamp

## Features ##

1. Sellers:<br>
- endpoint: /sellers
- /sellers `[POST]`: Create a Seller
- /sellers `[GET]`: List all Seller (READ)
- /sellers/:id `[GET]`: List a Seller (READ)
- /sellers/:id `[PUT]`: Modify Seller (UPDATE)
- /sellers/:id `[DELETE]`: Delete Seller (DELETE)

2. Warehouses:<br>
- endpoint: /warehouses
- /warehouses `[POST]`: Create a Warehouse
- /warehouses `[GET]`: List all Warehouse (READ)
- /warehouses/:id `[GET]`: List a Warehouse (READ)
- /warehouses/:id `[PATCH]`: Modify Warehouse (UPDATE)
- /warehouses/:id `[DELETE]`: Delete Warehouse (DELETE)

3. Sections:<br>
- endpoint: /sections
- /sections `[POST]`: Create a Section
- /sections `[GET]`: List all Section (READ)
- /sections/:id `[GET]`: List a Section (READ)
- /sections/:id `[PATCH]`: Modify Section (UPDATE)
- /sections/:id `[DELETE]`: Delete Section (DELETE)

4. Products:<br>
- endpoint: /products
- /products `[POST]`: Create a Product
- /products `[GET]`: List all Product (READ)
- /products/:id `[GET]`: List a Product (READ)
- /products/:id `[PATCH]`: Modify Product (UPDATE)
- /products/:id `[DELETE]`: Delete Product (DELETE)

5. Employees:<br>
- endpoint: /employees
- /employees `[POST]`: Create an Employee
- /employees `[GET]`: List all Employee (READ)
- /employees/:id `[GET]`: List an Employee (READ)
- /employees/:id `[PATCH]`: Modify Employee (UPDATE)
- /employees/:id `[DELETE]`: Delete Employee (DELETE)

6. Buyers:<br>
- endpoint: /buyers
- /buyers `[POST]`: Create a Buyer
- /buyers `[GET]`: List all Buyers (READ)
- /buyers/:id `[GET]`: List a Buyer (READ)
- /buyers/:id `[PUT]`: Modify Buyers (UPDATE)
- /buyers/:id `[DELETE]`: Delete Buyer (DELETE)

## Technologies ##

The following tools were used in this project:

- [Go](https://go.dev/)
- [Gin](https://gin-gonic.com/)
- [Swagger](https://swagger.io/)
- [Validator](https://pkg.go.dev/github.com/go-playground/validator/v10)
- [UUID](https://pkg.go.dev/github.com/google/UUID)
- [Godotenv](https://github.com/joho/godotenv)
- [Testify](https://github.com/stretchr/testify)
- [Mockery](https://github.com/vektra/mockery)

## Requirements ##

Before starting, you need to have [Git](https://git-scm.com), [Go](https://go.dev/) and [Swagger](https://swagger.io/) installed.

## Starting ##

```bash
# Clone this project
git clone https://github.com/Gopher-Rangers/mercadofresco-gopherrangers

# Access
cd mercadofresco-gopherrangers

# Install requirements
go get -u

# Access the server folder
cd /cmd/server

# Run the project
go run main.go

# To run the tests
cd ../..
go test ./... 

# To see the tests coverage
go test -coverprofile=coverage.out ./...

# To see the tests coverage in each function
go tool cover -func=coverage.out 

# The server will initialize in the <http://localhost:8080/api/v1/>
# To see the documentation go to <http://localhost:8080/docs/index.html#/>
```

## License ##

This project is under license from Apache 2.0. For more details, see the [LICENSE](LICENSE) file.

&#xa0;

<a href="#top">Back to top</a>
