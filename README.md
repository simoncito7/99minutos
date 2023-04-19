# 99minutos

## Setup
To run this project, you will need to follow these steps:

## Prerequisites
Docker
Go (at least version 1.16)
Postman or some tool to hit endpoints.

## Installation
Clone the repository.
Start the Postgres container by running `make postgres`.
Create the database by running `make createdb`.
Run the migrations by running `make migrateup`.

## Usage
To start the server, execute `make run`. This will start the server on port `8080` (you can set the port number).

Endpoints
The following endpoints are available:

GET /test - This is a test endpoint to check if the server is running.
POST /client - Create a new client.
POST /client/login - Login a client.
POST /order - Create a new order.
GET /order/:id - Get an order by ID.
GET /orders - Get all orders.
PUT /order/update - Update an order's status.
DELETE /order/:id - Cancel an order.

Testing
To run tests, execute `make test`. This will run all tests and show the coverage.

