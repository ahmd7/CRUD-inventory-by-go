# myInventory API

A RESTful inventory management system built with Go that allows you to track and manage products.

## Features

- Create, read, update, and delete products
- MySQL database integration
- Simple and easy-to-use RESTful API

## Prerequisites

- Go 1.24+
- MySQL database
- Git (optional)

## Installation

1. Clone the repository (or download the source code):
   ```
   git clone [[repository-url]](https://github.com/ahmd7/CRUD-inventory-by-go.git)
   cd myInventory
   ```

2. Configure your database connection:
   Edit the `constants.go` file to update your database credentials if needed.

3. Install dependencies:
   ```
   go mod download
   ```

4. Build the application:
   ```
   go build
   ```

## Running the Application

Execute the compiled binary:
```
./myInventory
```

By default, the server runs on `localhost:10000`.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/products` | Get all products |
| GET | `/products/{id}` | Get a single product by ID |
| POST | `/product` | Create a new product |
| PUT | `/products/{id}` | Update a product by ID |
| DELETE | `/products/{id}` | Delete a product by ID |

### Sample Request Bodies

**Create/Update Product (POST/PUT):**
```json
{
  "name": "Product Name",
  "quantity": 100,
  "price": 19.99
}
```

## Database Structure

The application expects a MySQL database with a `products` table:

```sql
CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  quantity INT NOT NULL,
  price DECIMAL(10, 2) NOT NULL
);
```

## Testing

Run tests with:
```
go test
```


