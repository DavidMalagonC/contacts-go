
# Contacts API

## Project Description

This project is a Contacts API built using Go. It allows users to perform CRUD operations (Create, Read, Update, Delete) on a contact list via a REST API. The data is stored in a MySQL database. The API is containerized with Docker.

### Design Decisions

- **Go**: The API is written in Go for its simplicity and performance.
- **REST API**: A RESTful API was chosen for easy integration with front-end applications or other services.
- **MySQL**: MySQL is used as the database to persist contact information.
- **Docker**: The API and MySQL are containerized using Docker, making it easy to run and deploy in different environments.
- **SQL Initialization**: The database schema is automatically created at startup via an SQL initialization script.

## Technologies Used

- **Go**: For the main application logic.
- **MySQL**: As the database to store contacts.
- **Docker**: For containerization.
- **net/http**: For handling HTTP requests in the Go application.

## Execution Instructions

1. **Clone the repository**:

2. **Build and run the application using Docker**:

   docker-compose up --build

This will build the Go application and start both the Go server and MySQL in Docker containers.

## How to Test the Application

### Get All Contacts

To retrieve all contacts, use the following `curl` command:
curl -X GET http://localhost:8080/contact

### Create a New Contact

To create a new contact, use the following `curl` command:
curl -X POST -H "Content-Type: application/json" \
-d '{"name":"John Doe", "email":"john.doe@example.com", "phone":"1234567890"}' \
http://localhost:8080/contact

### Update an Existing Contact

To update an existing contact, use the following `curl` command. Replace `1` with the ID of the contact you want to update:
curl -X PUT -H "Content-Type: application/json" \
-d '{"name":"John Smith", "email":"john.smith@example.com", "phone":"0987654321"}' \
http://localhost:8080/contacts/

### Delete a Contact

To delete a contact, use the following `curl` command. Replace `1` with the ID of the contact you want to delete:
curl -X DELETE http://localhost:8080/contacts/

## Dockerfile

The project includes a Dockerfile for containerizing the Go application. You can build and run the application in a Docker container using the following commands:

1. **Build the Docker image**:

   docker build --tag contacts-api .

2. **Run the Docker container**:

   docker run -p 8080:8080 contacts-api

## SQL Initialization

The MySQL database is initialized with a table called `contacts` using an `init.sql` file located in the `config` directory. The script automatically creates the table and inserts a few test records if they don't already exist.
