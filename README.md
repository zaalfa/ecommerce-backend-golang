# 🛒 E-Commerce Backend API (Golang)

Backend API untuk sistem e-commerce dengan fitur shopping cart, checkout, dan order management. Dibangun menggunakan **Clean Architecture** untuk maintainability dan scalability.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-14+-336791?style=flat&logo=postgresql)
![License](https://img.shields.io/badge/License-MIT-green.svg)

---

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Architecture](#-architecture)
- [Quick Start](#-quick-start)
- [API Documentation](#-api-documentation)
- [Testing](#-testing)
- [Project Structure](#-project-structure)
- [Future Improvements](#-future-improvements)

---

## ✨ Features

### 🔐 Authentication & Authorization
- User registration dengan email validation
- JWT-based authentication
- Role-based access control (User & Admin)
- Password hashing dengan bcrypt

### 🛍️ Product Management
- CRUD operations untuk products
- Stock management
- Product listing dengan detail lengkap
- Admin-only product creation

### 🛒 Shopping Cart
- Add/update/remove items
- Real-time stock validation
- Quantity adjustment
- Multi-product cart support

### 📦 Order Management
- Seamless checkout dari cart
- Automatic stock reduction
- Order history tracking
- Order status management (pending, paid, shipped, delivered, cancelled)
- Transaction-safe checkout (rollback on error)

### 👨‍💼 Admin Features
- Product creation & management
- Order status updates
- Protected dengan role-based middleware

---

## 🛠️ Tech Stack

### Backend Framework
- **Go 1.21+** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM library
- **PostgreSQL** - Database

### Security & Authentication
- **JWT** - Token-based authentication
- **bcrypt** - Password hashing
- **golang.org/x/crypto** - Cryptographic functions

### Architecture Pattern
- **Clean Architecture** - Separation of concerns
- **Repository Pattern** - Data access abstraction
- **Service Layer** - Business logic isolation
- **Dependency Injection** - Loose coupling

--- 

## 🏗️ Architecture

Project ini menggunakan **Clean Architecture** dengan layer separation:
```
┌─────────────────────────────────────────┐
│           HTTP Handler Layer            │
│            (Controllers)                │
├─────────────────────────────────────────┤
│          Business Logic Layer           │
│             (Services)                  │
├─────────────────────────────────────────┤
│         Data Access Layer               │
│          (Repositories)                 │
├─────────────────────────────────────────┤
│              Database                   │
│            (PostgreSQL)                 │
└─────────────────────────────────────────┘
```


**Benefits:**
- High testability
- Maintainable and scalable codebase
- Framework & database independence



## 🚀 Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 14+
- Git

### Installation

```bash
git clone https://github.com/zaalfa/ecommerce-backend-golang.git
cd ecommerce-backend-golang
go mod download
go mod tidy


## 📚 API Documentation

### Base URL
```
http://localhost:8080
```

### Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/auth/register` | Register user baru | ❌ |
| POST | `/auth/login` | Login & get JWT token | ❌ |

### Product Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/products` | Get all products | ❌ |
| GET | `/products/:id` | Get product by ID | ❌ |
| POST | `/admin/products` | Create new product | ✅ Admin |

### Cart Endpoints (Protected)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/cart` | Get user's cart | ✅ User |
| POST | `/cart/items` | Add item to cart | ✅ User |
| PUT | `/cart/items/:id` | Update item quantity | ✅ User |
| DELETE | `/cart/items/:id` | Remove item from cart | ✅ User |
| DELETE | `/cart` | Clear entire cart | ✅ User |

### Order Endpoints (Protected)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/orders` | Checkout (create order) | ✅ User |
| GET | `/orders` | Get order history | ✅ User |
| GET | `/orders/:id` | Get order detail | ✅ User |
| PUT | `/admin/orders/:id/status` | Update order status | ✅ Admin |

### Example Request/Response

**Register User:**
```bash
POST /auth/register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "message": "register success"
}
```

**Add to Cart:**
```bash
POST /cart/items
Authorization: Bearer {your_token}
Content-Type: application/json

{
  "product_id": 1,
  "quantity": 2
}
```

**Response:**
```json
{
  "message": "item added to cart"
}
```

📖 **Full API Documentation:** [View detailed API docs](docs/API.md)

---

## 🧪 Testing

### What's Being Tested

✅ **Authentication Flow**
- User registration dengan validasi email
- Login & JWT token generation
- Token validation pada protected routes

✅ **Authorization & Security**
- JWT middleware protection
- Role-based access control (user vs admin)
- Invalid/expired token handling

✅ **Cart Management**
- Add items dengan stock validation
- Update quantity dengan boundary testing
- Remove items & cart clearing
- Multi-product cart scenarios

✅ **Checkout & Order**
- End-to-end checkout process
- Stock reduction verification
- Transaction rollback pada error
- Empty cart validation

✅ **Error Handling**
- Insufficient stock scenarios
- Invalid product ID
- Unauthorized access attempts
- Missing required fields

### Testing Tools
- **Postman** - API testing collection included
- **cURL** - Command-line testing examples
- **Manual Testing** - Step-by-step guide

### Quick Test
```bash
# Test authentication
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'

# Test cart (requires token)
curl -X GET http://localhost:8080/cart \
  -H "Authorization: Bearer YOUR_TOKEN"
```

🧪 **Complete Testing Guide:** [View testing documentation](docs/TESTING.md)

---

## 📁 Project Structure
```
ecommerce-backend-golang/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── database.go         # Database configuration
│   ├── controllers/            # HTTP handlers
│   │   ├── auth_controller.go
│   │   ├── cart_controller.go
│   │   ├── order_controller.go
│   │   ├── product_controller.go
│   │   └── user_controller.go
│   ├── middleware/             # HTTP middlewares
│   │   ├── auth_middleware.go
│   │   └── admin_middleware.go
│   ├── models/                 # Domain models
│   │   ├── user.go
│   │   ├── product.go
│   │   ├── cart.go
│   │   └── order.go
│   ├── repositories/           # Data access layer
│   │   ├── user_repository.go
│   │   ├── product_repository.go
│   │   ├── cart_repository.go
│   │   └── order_repository.go
│   ├── services/               # Business logic layer
│   │   ├── auth_service.go
│   │   ├── product_service.go
│   │   ├── cart_service.go
│   │   └── order_service.go
│   └── routes/
│       └── router.go           # Route definitions
├── pkg/
│   └── utils/
│       └── jwt.go              # JWT utilities
├── docs/                       # Documentation
│   ├── API.md                  # API documentation
│   └── TESTING.md              # Testing guide
├── .env.example                # Environment variables template
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

**Design Principles:**
- **Separation of Concerns** - Each layer has specific responsibility
- **Dependency Injection** - Dependencies injected via constructors
- **Single Responsibility** - Each file/struct has one clear purpose
- **DRY (Don't Repeat Yourself)** - Reusable components

---

## 🔮 Future Improvements

### Phase 1 - Core Features
- [ ] Update & Delete products (admin)
- [ ] Pagination untuk product listing
- [ ] Search & filter products
- [ ] Product categories

### Phase 2 - Enhanced Features
- [ ] Image upload untuk products
- [ ] Product reviews & ratings
- [ ] Wishlist functionality
- [ ] User profile management

### Phase 3 - Payment & Notifications
- [ ] Payment gateway integration (Midtrans/Stripe)
- [ ] Email notifications (order confirmation, shipping updates)
- [ ] Order tracking system
- [ ] Invoice generation

### Phase 4 - Advanced Features
- [ ] Discount & promo codes
- [ ] Inventory management
- [ ] Sales analytics dashboard
- [ ] Multi-language support

### Technical Improvements
- [ ] Unit tests dengan testify
- [ ] Integration tests
- [ ] API rate limiting
- [ ] Redis caching
- [ ] Docker containerization
- [ ] CI/CD pipeline
- [ ] API documentation dengan Swagger
- [ ] Logging dengan structured logger (zap/logrus)

---

## 👨‍💻 Developer

**Zalfa** - [GitHub](https://github.com/zaalfa)

---

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [JWT-Go](https://github.com/golang-jwt/jwt)
- Clean Architecture principles by Robert C. Martin

---

⭐ **If you find this project helpful, please give it a star!**