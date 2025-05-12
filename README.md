# SB-Go-Quiz-Nabiel API

This is a RESTful API built with Golang using the Gin framework and PostgreSQL for data storage. It manages **books** and **categories**, and includes middleware for authentication.

## 🧰 Tech Stack

- Go (Golang)
- Gin Web Framework
- PostgreSQL
- `.env` Configuration with `godotenv`

## 📦 Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/KonekkoNekko/sb-go-quiz-nabiel.git
   cd sb-go-quiz-nabiel
   ```

2. Set up your `.env` file inside `config/.env`:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_user
   DB_PASSWORD=your_password
   DB_NAME=your_database
   ```

3. Run the application:
   ```bash
   go run .
   ```

---

## 📘 API Documentation

All endpoints are prefixed with `/api` and require authentication using middleware.

### 📚 Books Endpoints

| Method | Endpoint              | Description            |
|--------|-----------------------|------------------------|
| GET    | `/api/books`          | Get all books          |
| GET    | `/api/books/:id`      | Get a book by ID       |
| POST   | `/api/books`          | Create a new book      |
| PUT    | `/api/books/:id`      | Update a book by ID    |
| DELETE | `/api/books/:id`      | Delete a book by ID    |

### 🗂️ Categories Endpoints

| Method | Endpoint                          | Description                            |
|--------|-----------------------------------|----------------------------------------|
| GET    | `/api/categories`                 | Get all categories                     |
| GET    | `/api/categories/:id`             | Get a category by ID                   |
| GET    | `/api/categories/:id/books`       | Get books under a specific category    |
| POST   | `/api/categories`                 | Create a new category                  |
| PUT    | `/api/categories/:id`             | Update a category by ID                |
| DELETE | `/api/categories/:id`             | Delete a category by ID                |

---

---

## 📄 License

This project is licensed under the MIT License.
