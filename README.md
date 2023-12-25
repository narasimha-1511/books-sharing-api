# Books - Sharing API

## Table of Contents
- [Prerequisites](#prerequisites)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [License](#license)
- [Apis](#api-usage-guide)
- [Documentation](#documentation)

## Prerequisites

Before using this repository, make sure you have the following prerequisites installed:

- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Go: [Installation Guide](https://golang.org/doc/install)

## Tech Stack

This project uses the following technologies:

- [Go](https://golang.org/): The backend service is written in Go.
- [Docker](https://www.docker.com/): Docker is used for setting up the database and running the application in a containerized environment.
- [GORM](https://gorm.io/): GORM (Go Object-Relational Mapper) is used for handling database operations.
- [Gin](https://gin-gonic.com/): Gin is a web framework used for routing and handling HTTP requests.

## Getting Started

Follow these steps to get started with the project:

1. Clone the repository: 
    ```bash
    git clone https://github.com/narasimha-1511/zolo-backend.git
    ```

2. Run Docker to setup Postgres:

    ```bash
    sudo docker compose up --build
    ```

    This works because we have a YAML file that contains the configuration of the PostgreSQL.

3. Run the main file:
    To run the main file and start the application, execute the following command in your terminal:

    ```bash
    go run main.go
    ```
## API Usage Guide

Here are the available endpoints and how to use them:

- **Add a book**

  `PUT /api/v1/booky/`

  Parameters:
    - name of the book
    - author of the book
    - title of the book

- **Browse the shared books**

  `GET /api/v1/booky/`

- **Borrow a book for a certain duration of time**

  `PUT /api/v1/booky/<book_id>/borrow`

  Parameters:

  - `book_id`: ID of the Book that has been added for sharing with others
  - `borrow_period`: No of Days You Want To B 

- **View all the borrowed books**

  `GET /api/v1/booky/borrowed`

- **Return a borrowed book**

  `POST /api/v1/booky/<book_id>/borrow/<borrow_id>`

  Parameters:

  - `book_id`: ID of the Book that has been added for sharing with others
  - `borrow_id`: ID of the Borrow operation for Book with book_id

## Documentation

The following APIs are available:

- ```GET``` ``/api/v1/booky`` Get all books available.

    Call Using the folliwing curl command
    ```bash
    curl -X GET localhost:<port>/api/v1/booky
    ```
    
     Example here Normaly Books are Identifed By Its `book_id`
     
     This is Response for the following curl command

    ```json
        "Books": [
            {
            "ID": 1,
            "CreatedAt": "2023-12-25T15:20:19.8477+05:30",
            "UpdatedAt": "2023-12-25T15:20:19.8477+05:30",
            "DeletedAt": null,
            "book_id": 1600532220260512696,
            "name": "The Book Name",
            "title": "The Title",
            "author": "Author",
            "borrowed": false
            },
            {
            "ID": 2,
            "CreatedAt": "2023-12-25T15:20:14.8477+05:30",
            "UpdatedAt": "2023-12-25T15:20:20.8477+05:30",
            "DeletedAt": null,
            "book_id":  684684515656488798,
            "name":     "Narasimha",
            "title":    "Devil",
            "author":   "God",
            "borrowed": false
            }...
        ]
    ```
- ```PUT``` ``/api/v1/booky``: Create a new book.

    To create a new book, use the following form-encoded 
    
    request body:
        "Request Body": name=myName&title=myTitle&author=myAuthor

    Use the Follwing curl command to test it
    ```bash
    curl -X PUT -d "name=myName&title=myTitle&author=myAuthor" http://localhost:<port>/api/v1/booky
    ```
    The Response Here is Similar to the Get Book API.

    ```json
    {
  "book_id": 12345,
  "name":   "myName",
  "title":  "myTitle"
  "author": "myAuthor",
  "status": "Book Created Successfully",
    }
    ```

- `PUT` ``api/v1/booky/<book_id>/borrow``:

    To borrow a book by sending a PUT request to the specified endpoint. The <book_id> parameter should be replaced with the ID of the book to be borrowed.

    Example usage:
    `PUT` `/api/v1/booky/12345/borrow`
    
    You Need to Specify the Time Period for Borrowing 
    If You Won't It Will be 7 Days by Default.

    Testing Code:

    ```bash
    curl -X PUT -d "borrow_period=3" https://localhost:<port>/api/v1/booky/12345/borrow
    ```

    Response:
    - If the book is available and the borrowing is successful, the API will return a success message.
        ```json
            {
            "book_id": 7218268149424024368,
            "borrowed_id": "e2f912cb-9172-4d82-a5bc-23502b46cc6b",
            "start_time": "2023-12-25T18:43:50.689008544+05:30",
            "end_time": "2023-12-28T18:43:50.689008618+05:30",
            "status": "Book Borrowed Successfully"
            "returned": false,
            }
        ```
    - If the book is already borrowed or unavailable, the API will return an error message.
        ```json
            {
                "message": "Book is already borrowed"
            }
        ```

- `GET` `api/v1/booky/borrowed`:

    You Will get the Books Have Been Borrowed Till Now

    use the Statement:

    ```bash
    curl -X GET localhost:<port>/api/v1/booky/borrowed
    ``` 
    
    This is the response for the following curl command
    
    ```json
    "BorrowedBooks": [
        {
            "book_id": 7218268149424024368,
            "borrowed_id": "e2f912cb-9172-4d82-a5bc-23502b46cc6b",
            "start_time": "2023-12-25T18:43:50.689008544+05:30",
            "end_time": "2023-12-28T18:43:50.689008618+05:30",
            "book_name": "Money",
            "returned": false
        },
        {
            "book_id": 1234567890,
            "borrowed_id": "f3a4b5c6-d7e8-f9a0-b1c2-d3e4f5a6b7c8",
            "start_time": "2023-12-25T18:43:50.689008544+05:30",
            "end_time": "2023-12-28T18:43:50.689008618+05:30",
            "book_name": "Dead Boy",
            "returned": false
        }
    ]
    ```
- `POST` `api/v1/<book_id>/borrow/<borrow_id>`: Return a borrowed book.

    To return a borrowed book, send a POST request to the specified endpoint. The `<book_id>` and `<borrow_id>` parameters should be replaced with the ID of the book and the ID of the borrowing respectively.

    Example usage:
    `POST` `/api/v1/12345/borrow/e2f912cb-9172-4d82-a5bc-23502b46cc6b`

    Testing Code:

    ```bash
    curl -X POST https://localhost:<port>/api/v1/12345/borrow/e2f912cb-9172-4d82-a5bc-23502b46cc6b
    ```

    Response:
    - If the book is successfully returned, the API will return a success message.
        ```json
        {
            "book_id": 12345,
            "borrowed_id": "e2f912cb-9172-4d82-a5bc-23502b46cc6b",
            "status": "Book Returned Successfully"
        }
        ```
    - If the book is not found or the borrowing is not valid, the API will return an error message.
        ```json
        {
            "message": "Invalid book or borrowing ID"
        }
        ```

## License

This project is licensed under the [MIT License](LICENSE).

