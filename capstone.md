# Capstone Project: Food Delivery API

## Project Overview

The **Food Delivery API** is a scalable backend service for managing food orders, users, restaurants, and delivery operations. It leverages **Golang**, **Gin** (for HTTP routing), **GORM** (ORM for PostgreSQL), and **Goroutines** for concurrent processing.

## Tech Stack

- **Golang**: Backend service
- **Gin**: Web framework for RESTful APIs
- **GORM**: ORM for PostgreSQL
- **PostgreSQL**: Database
- **Goroutines & Channels**: Handling concurrent operations


## Features

### 1. User Management

- User Registration & Login (middleware for Username and password check to proceed or not based on the username and password )
- Role-based access (Admin, User, Restaurant Owner, Delivery Person, included in the middleware)

### 2. Restaurant & Food Management

- Add, update, delete restaurants
- Manage food items (CRUD operations)
- Search restaurants by location & category

### 3. Order & Delivery

- Place orders with multiple food items
- Order status tracking (Pending, Preparing, Out for Delivery, Delivered)
- Assigning delivery partners (Goroutines for order processing)

### 4. Payment & Reviews

- Payment processing simulation
- User reviews & ratings for food items

### 5. Audit Logs & Background Jobs

- Goroutines for logging user activities asynchronously

## Project Structure (You can follow your own structure)

```
food-delivery-api/
├── main.go
├── config/
│   ├── config.go
├── models/
│   ├── user.go
│   ├── restaurant.go
│   ├── food.go
│   ├── order.go
├── controllers/
│   ├── auth_controller.go
│   ├── restaurant_controller.go
│   ├── order_controller.go
├── repositories/
│   ├── user_repository.go
│   ├── order_repository.go
├── routes/
│   ├── routes.go
├── db/
│   ├── migrations/
│   ├── seed.go
├── utils/
│   ├── response.go
│   ├── logger.go
```

## Database Schema (PostgreSQL)

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    password_hash TEXT,
    role VARCHAR(50) CHECK (role IN ('admin', 'user', 'restaurant_owner', 'delivery_person')),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE restaurants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    location TEXT,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE food_items (
    id SERIAL PRIMARY KEY,
    restaurant_id INT REFERENCES restaurants(id),
    name VARCHAR(255),
    price DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    status VARCHAR(50) CHECK (status IN ('pending', 'preparing', 'out_for_delivery', 'delivered')),
    created_at TIMESTAMP DEFAULT now()
);
```

## Goroutine Implementation (Order Processing)

```go 
// order id can be sent to a channel and that channel reads the order data from the database and process it
// or send full order data to the channel 
func ProcessOrder(orderID chan uint) {
    go func(id uint) {
        // based on the orderID fetch details from the database
        time.Sleep(5 * time.Second) // Simulating order processing time
        fmt.Printf("Order %d has been processed\n", id)
    }(orderID)
}
```

## API Endpoints

| Method | Endpoint     | Description           |
| ------ | ------------ | --------------------- |
| POST   | /register    | User registration     |
| POST   | /login       | User authentication   |
| GET    | /restaurants | Get all restaurants   |
| GET    | /food/{id}   | Get food item details |
| POST   | /order       | Place an order        |
| GET    | /order/{id}  | Get order status      |

## Running the Project

```sh
git clone https://github.com/yourrepo/food-delivery-api.git
cd food-delivery-api
go mod tidy
go run main.go
```

## Conclusion

This **Golang Capstone Project** demonstrates building a production-ready backend service for food delivery. It efficiently handles concurrent operations, secure authentication, and scalable database management using PostgreSQL.


