# Libr
#### The Smart Way to Manage Your Library.

## Summary
Libr is a library management application designed to simplify the management of your library's collections and operations. The application offers essential features, including Book Management for organizing book inventories, Author Management for centralized author data management, Book Categories Management to categorize and organize collections, and Book Stock Management to ensure book availability.

## key Feature
- **Book Management**: Add, update, delete, and view book details. [ON PROGRESS]
- **Author Management**: Manage author information. [ON PROGRESS]
- **Book Categories**: Organize and manage book categories. [DONE]
- **Borrowing & Returning Book**: Facilitate the process of borrowing and returning books. [ON PROGRESS]

## Installation
### Prerequisites

- Go version 1.20+
- PostgreSQL
- Redis
- RabbitMq
- Docker & Docker Compose

### Installation Steps

1. **Clone this repository:**

   ```bash
   git clone https://github.com/arifbugaresa/go-libr.git

2. **Install dependencies:**

   ```bash
   go mod tidy
   go mod vendor
   
## Usage
### Running the API
Once the application is running, the API will be available at http://localhost:8080. You can access the API endpoints using tools like Postman or you can use swagger docs (http://localhost:8080/swagger/index.html)

## List Endpoint
### Category
Service Category: https://github.com/arifbugaresa/go-libr-category
- `GET` `localhost:8080/api/categories` - Retrieve a list of books.
- `POST` `localhost:8080/api/categories` - Insert data book by id.
- `PUT` `localhost:8080/api/categories/:id` - Update data book by id.


## Architecture
<img width="777" alt="Screenshot 2024-08-28 at 11 21 16" src="https://github.com/user-attachments/assets/e65db8ef-d648-49d0-a8f6-0ebefe411fa3">

## Entity Relation Diagram
<img width="777" alt="Screenshot 2024-08-27 at 13 15 28" src="https://github.com/user-attachments/assets/96a7d75c-260d-4a33-b004-b3404c12ea24">

## Changelog
- v1.0.0 - Initial release with basic library management features.

## License
This project does not have an official license yet.

## Author
Developed by [Arif Yuniarto Fajar Bugaresa](https://www.linkedin.com/in/arifbugaresa/)

