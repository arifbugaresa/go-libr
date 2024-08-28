# Libr
#### The Smart Way to Manage Your Library.

## Summary
Libr is a library management application designed to simplify the management of your library's collections and operations. The application offers essential features, including Book Management for organizing book inventories, Author Management for centralized author data management, Book Categories Management to categorize and organize collections, and Book Stock Management to ensure book availability.

## key Feature
- **Book Management**: Add, update, delete, and view book details.
- **Author Management**: Manage author information.
- **Book Categories**: Organize and manage book categories.
- **Borrowing & Returning Book**: Facilitate the process of borrowing and returning books.

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
Once the application is running, the API will be available at http://localhost:8080. You can access the API endpoints using tools like cURL or Postman or you can use swagger docs.

## List Endpoint
### Category
`GET` `localhost:8080/api/categories` - Retrieve a list of books.

## Architecture

## Entity Relation Diagram

## Contributing
If you would like to contribute, please fork this repository and create a pull request. For major changes, please open an issue first to discuss what you would like to change.

## Changelog
- v1.0.0 - Initial release with basic library management features.

## License
This project does not have an official license yet.

## Author
Developed by [Arif Yuniarto Fajar Bugaresa](https://www.linkedin.com/in/arifbugaresa/)