<p align="center">
<img src="https://github.com/ashleymcnamara/gophers/blob/master/TEACHING_GOPHER.png?raw=true" width="250"/>
</p>

<h1 align="center">Mercado Fresco Gopher Rangers</h1>

<p align="center">
 <img alt="Github top language" src="https://img.shields.io/github/languages/top/Gopher-Rangers/mercadofresco-gopherrangers?color=56BEB8">
  <img alt="Github language count" src="https://img.shields.io/github/languages/count/Gopher-Rangers/mercadofresco-gopherrangers?color=56BEB8">
  <img alt="Repository size" src="https://img.shields.io/github/repo-size/Gopher-Rangers/mercadofresco-gopherrangers?color=56BEB8">
  <img alt="License" src="https://img.shields.io/github/license/Gopher-Rangers/mercadofresco-gopherrangers?color=56BEB8">
  <img alt="Github issues" src="https://img.shields.io/github/issues/Gopher-Rangers/mercadofresco-gopherrangers?color=56BEB8" />
  <img alt="Github forks" src="https://img.shields.io/github/forks/Gopher-Rangers/mercadofresco-gopherrangers?color=56BEB8" />
  <img alt="Github stars" src="https://img.shields.io/github/stars/Gopher-Rangers/mercadofresco-gopherrangers?color=56BEB8" />
</p>

<p align="center">
  <a href="https://github.com/Gopher-Rangers/mercadofresco-gopherrangers/actions/workflows/test.yml">
    <img src="https://github.com/Gopher-Rangers/mercadofresco-gopherrangers/actions/workflows/test.yml/badge.svg">
  </a>
  <a href="https://codecov.io/gh/Gopher-Rangers/mercadofresco-gopherrangers"> 
    <img src="https://codecov.io/gh/Gopher-Rangers/mercadofresco-gopherrangers/branch/main/graph/badge.svg?token=NUUR12FFLR"> 
  </a>
  <a href="https://goreportcard.com/report/github.com/Gopher-Rangers/mercadofresco-gopherrangers"> 
    <img src="https://goreportcard.com/badge/github.com/Gopher-Rangers/mercadofresco-gopherrangers"> 
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
<table>
  <tr>
    <td>
      1.1. Sellers:<br>
      - /sellers <code>[POST]</code>: Create a Seller (CREATE)<br>
      - /sellers <code>[GET]</code>: List all Sellers (READ)<br>
      - /sellers/:id <code>[GET]</code>: List a Seller (READ)<br>
      - /sellers/:id <code>[PUT]</code>: Modify a Seller (UPDATE)<br>
      - /sellers/:id <code>[DELETE]</code>: Delete a Seller (DELETE)<br>
    </td>
    <td>
      1.2. Warehouses:<br>
      - /warehouses <code>[POST]</code>: Create a Warehouse (CREATE)<br>
      - /warehouses <code>[GET]</code>: List all Warehouses (READ)<br>
      - /warehouses/:id <code>[GET]</code>: List a Warehouse (READ)<br>
      - /warehouses/:id <code>[PATCH]</code>: Modify a Warehouse (UPDATE)<br>
      - /warehouses/:id <code>[DELETE]</code>: Delete a Warehouse (DELETE)<br>
    </td>
  </tr>

  <tr>
    <td>
      1.3. Sections:<br>
      - /sections <code>[POST]</code>: Create a Section (CREATE)<br>
      - /sections <code>[GET]</code>: List all Sections (READ)<br>
      - /sections/:id <code>[GET]</code>: List a Section (READ)<br>
      - /sections/:id <code>[PATCH]</code>: Modify a Section (UPDATE)<br>
      - /sections/:id <code>[DELETE]</code>: Delete a Section (DELETE)<br>
    </td>
    <td>
      1.4. Products:<br>
      - /products <code>[POST]</code>: Create a Product (CREATE)<br>
      - /products <code>[GET]</code>: List all Products (READ)<br>
      - /products/:id <code>[GET]</code>: List a Product (READ)<br>
      - /products/:id <code>[PATCH]</code>: Modify a Product (UPDATE)<br>
      - /products/:id <code>[DELETE]</code>: Delete a Product (DELETE)<br>
    </td>
  </tr>

  <tr>
    <td>
      1.5. Employees:<br>
      - /employees <code>[POST]</code>: Create an Employee (CREATE)<br>
      - /employees <code>[GET]</code>: List all Employeea (READ)<br>
      - /employees/:id <code>[GET]</code>: List an Employee (READ)<br>
      - /employees/:id <code>[PATCH]</code>: Modify an Employee (UPDATE)<br>
      - /employees/:id <code>[DELETE]</code>: Delete an Employee (DELETE)<br>
    </td>
    <td>
      1.6. Buyers:<br>
      - /buyers <code>[POST]</code>: Create a Buyer (CREATE)<br>
      - /buyers <code>[GET]</code>: List all Buyers (READ)<br>
      - /buyers/:id <code>[GET]</code>: List a Buyer (READ)<br>
      - /buyers/:id <code>[PUT]</code>: Modify a Buyer (UPDATE)<br>
      - /buyers/:id <code>[DELETE]</code>: Delete a Buyer (DELETE)<br>
    </td>
  </tr>

  <tr>
    <td>
      2.1. Localities:<br>
      - /localities <code>[POST]</code>: Create a Locality (CREATE)<br>
      - /localities/reportSellers <code>[GET]</code>: List all Localities (READ)<br>
      - /localities/reportSellers?id=some_id <code>[GET]</code>: List a Locality (READ)<br>
      - /sellers <code>[POST]</code>: List all Sellers (READ)<br>
    </td>
    <td>
      2.2. Carries:<br>
      - /carries <code>[POST]</code>: Create a Carry (CREATE)<br>
      - /localities/reportCarries <code>[GET]</code>: List all Carries (READ)<br>
      - /localities/reportCarries?id=some_id <code>[GET]</code>: List a Carrie (READ)<br>
    </td>
  </tr>

 <tr>
    <td>
      2.3. Product Batches:<br>
      - /productBatches <code>[POST]</code>: Create a Product Batch (CREATE)<br>
      - /sections/reportProducts <code>[GET]</code>: List all Product Batches (READ)<br>
      - /sections/reportProducts?id=some_id <code>[GET]</code>: List a Product Batch (READ)<br>
    </td>
    <td>
      2.4. Product Records:<br>
      - /productRecords <code>[POST]</code>: Create a Product Record (CREATE)<br>
      - /products/reportRecords <code>[GET]</code>: List all Product Records (READ)<br>
      - /products/reportRecords?id=some_id <code>[GET]</code>: List a Product Record (READ)<br>
    </td>
  </tr>

  <tr>
    <td>
      2.5. Inbound Orders:<br>
      - /inboundOrders <code>[POST]</code>: Create a Inbound Order (CREATE)<br>
      - /employees/reportInboundOrders <code>[GET]</code>: List all Inbound Orders (READ)<br>
      - /employees/reportInboundOrders?id=some_id <code>[GET]</code>: List a Inbound Order (READ)<br>
    </td>
    <td>
      2.6. Purchase Orders:<br>
      - /purchaseOrders <code>[POST]</code>: Create a Purchase Order (CREATE)<br>
      - /buyers/reportPurchaseOrders <code>[GET]</code>: List all Purchase Orders (READ)<br>
      - /buyers/reportPurchaseOrders?id=some_id <code>[GET]</code>: List a Purchase Order (READ)<br>
    </td>

  </tr>

</table>

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
)
Before starting, you need to have [Git](https://git-scm.com), [Go](https://go.dev/) and [Docker](https://www.docker.com/) installed.

## Starting ##

```bash
# Clone this project
git clone https://github.com/Gopher-Rangers/mercadofresco-gopherrangers

# Access
cd mercadofresco-gopherrangers

# Config the .env file following the .env_example

# Create a mabiadb database
docker-compose up

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
