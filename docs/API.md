# 📚 API Documentation

Complete API reference untuk E-Commerce Backend. Dokumentasi ini mencakup semua endpoints, request/response format, dan authentication requirements.

---

## 📋 Table of Contents

- [Base Information](#base-information)
- [Authentication](#authentication)
- [Products](#products)
- [Cart Management](#cart-management)
- [Orders](#orders)
- [Admin Operations](#admin-operations)
- [Error Responses](#error-responses)

---

## Base Information

### Base URL
```
http://localhost:8080
```

### Authentication
Endpoints yang membutuhkan authentication harus menyertakan JWT token di header:
```
Authorization: Bearer {your_jwt_token}
```

### Response Format
Semua responses menggunakan JSON format.

**Success Response:**
```json
{
  "data": {...},
  "message": "success message"
}
```

**Error Response:**
```json
{
  "error": "error message description"
}
```

---

## Authentication

### Register User

Register user baru ke sistem.

**Endpoint:** `POST /auth/register`

**Authentication:** Not required

**Request Body:**
```json
{
  "name": "string (required)",
  "email": "string (required, valid email)",
  "password": "string (required, min 6 characters)"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Success Response:** `201 Created`
```json
{
  "message": "register success"
}
```

**Error Responses:**

`400 Bad Request` - Validation error
```json
{
  "error": "Key: 'name' Error:Field validation for 'name' failed on the 'required' tag"
}
```

`400 Bad Request` - Email already exists
```json
{
  "error": "email already used"
}
```

---

### Login

Login dan mendapat JWT token untuk authentication.

**Endpoint:** `POST /auth/login`

**Authentication:** Not required

**Request Body:**
```json
{
  "email": "string (required)",
  "password": "string (required)"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Success Response:** `200 OK`
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ5NzI4MDAsInJvbGUiOiJ1c2VyIiwidXNlcl9pZCI6MX0.xxx",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "user",
    "created_at": "2025-01-04T10:00:00Z",
    "updated_at": "2025-01-04T10:00:00Z"
  }
}
```

**Error Response:** `401 Unauthorized`
```json
{
  "error": "Invalid email or password"
}
```

**Notes:**
- Token expires dalam 7 hari
- Token harus disimpan dan digunakan untuk authenticated requests
- Role bisa `user` atau `admin`

---

## Products

### Get All Products

Mendapat list semua products. Endpoint ini public (tidak perlu authentication).

**Endpoint:** `GET /products`

**Authentication:** Not required

**Example Request:**
```bash
curl -X GET http://localhost:8080/products
```

**Success Response:** `200 OK`
```json
[
  {
    "id": 1,
    "name": "iPhone 15 Pro",
    "description": "Latest iPhone with A17 chip",
    "price": 15000000,
    "stock": 10,
    "created_at": "2025-01-04T09:00:00Z",
    "updated_at": "2025-01-04T09:00:00Z"
  },
  {
    "id": 2,
    "name": "MacBook Pro M3",
    "description": "16-inch MacBook Pro",
    "price": 35000000,
    "stock": 5,
    "created_at": "2025-01-04T09:15:00Z",
    "updated_at": "2025-01-04T09:15:00Z"
  }
]
```

---

### Get Product by ID

Mendapat detail product berdasarkan ID.

**Endpoint:** `GET /products/:id`

**Authentication:** Not required

**URL Parameters:**
- `id` (integer) - Product ID

**Example Request:**
```bash
curl -X GET http://localhost:8080/products/1
```

**Success Response:** `200 OK`
```json
{
  "id": 1,
  "name": "iPhone 15 Pro",
  "description": "Latest iPhone with A17 chip",
  "price": 15000000,
  "stock": 10,
  "created_at": "2025-01-04T09:00:00Z",
  "updated_at": "2025-01-04T09:00:00Z"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "product not found"
}
```

---

### Create Product (Admin Only)

Create product baru. Hanya admin yang bisa akses endpoint ini.

**Endpoint:** `POST /admin/products`

**Authentication:** Required (Admin only)

**Headers:**
```
Authorization: Bearer {admin_token}
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "string (required)",
  "description": "string (optional)",
  "price": "integer (required)",
  "stock": "integer (required)"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/admin/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer {admin_token}" \
  -d '{
    "name": "AirPods Pro",
    "description": "Active Noise Cancellation",
    "price": 3500000,
    "stock": 20
  }'
```

**Success Response:** `201 Created`
```json
{
  "id": 3,
  "name": "AirPods Pro",
  "description": "Active Noise Cancellation",
  "price": 3500000,
  "stock": 20,
  "created_at": "2025-01-04T10:30:00Z",
  "updated_at": "2025-01-04T10:30:00Z"
}
```

**Error Responses:**

`401 Unauthorized` - No token atau token invalid
```json
{
  "error": "authorization header missing"
}
```

`403 Forbidden` - User bukan admin
```json
{
  "error": "admin access required"
}
```

`400 Bad Request` - Validation error
```json
{
  "error": "validation error message"
}
```

---

## Cart Management

Semua cart endpoints membutuhkan authentication (JWT token).

### Get Cart

Mendapat cart user yang sedang login beserta items-nya.

**Endpoint:** `GET /cart`

**Authentication:** Required (User)

**Headers:**
```
Authorization: Bearer {user_token}
```

**Example Request:**
```bash
curl -X GET http://localhost:8080/cart \
  -H "Authorization: Bearer {user_token}"
```

**Success Response:** `200 OK`
```json
{
  "id": 1,
  "user_id": 1,
  "items": [
    {
      "id": 1,
      "cart_id": 1,
      "product_id": 1,
      "product": {
        "id": 1,
        "name": "iPhone 15 Pro",
        "description": "Latest iPhone with A17 chip",
        "price": 15000000,
        "stock": 10
      },
      "quantity": 2,
      "created_at": "2025-01-04T10:30:00Z",
      "updated_at": "2025-01-04T10:30:00Z"
    }
  ],
  "created_at": "2025-01-04T10:30:00Z",
  "updated_at": "2025-01-04T10:35:00Z"
}
```

**Notes:**
- Cart dibuat otomatis saat user pertama kali add item
- Setiap user hanya punya 1 cart
- Items include full product details

**Error Response:** `404 Not Found`
```json
{
  "error": "cart not found"
}
```

---

### Add Item to Cart

Menambahkan product ke cart. Jika product sudah ada di cart, quantity akan ditambah.

**Endpoint:** `POST /cart/items`

**Authentication:** Required (User)

**Headers:**
```
Authorization: Bearer {user_token}
Content-Type: application/json
```

**Request Body:**
```json
{
  "product_id": "integer (required)",
  "quantity": "integer (required, min: 1)"
}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/cart/items \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer {user_token}" \
  -d '{
    "product_id": 1,
    "quantity": 2
  }'
```

**Success Response:** `200 OK`
```json
{
  "message": "item added to cart"
}
```

**Error Responses:**

`400 Bad Request` - Product not found
```json
{
  "error": "product not found"
}
```

`400 Bad Request` - Insufficient stock
```json
{
  "error": "insufficient stock"
}
```

`400 Bad Request` - Validation error
```json
{
  "error": "validation error message"
}
```

---

### Update Cart Item Quantity

Update quantity item yang sudah ada di cart.

**Endpoint:** `PUT /cart/items/:id`

**Authentication:** Required (User)

**Headers:**
```
Authorization: Bearer {user_token}
Content-Type: application/json
```

**URL Parameters:**
- `id` (integer) - Cart item ID (bukan product ID)

**Request Body:**
```json
{
  "quantity": "integer (required, min: 1)"
}
```

**Example Request:**
```bash
curl -X PUT http://localhost:8080/cart/items/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer {user_token}" \
  -d '{
    "quantity": 5
  }'
```

**Success Response:** `200 OK`
```json
{
  "message": "cart item updated"
}
```

**Error Responses:**

`400 Bad Request` - Item not found in cart
```json
{
  "error": "item not found in cart"
}
```

`400 Bad Request` - Insufficient stock
```json
{
  "error": "insufficient stock"
}
```

---

### Remove Item from Cart

Menghapus item dari cart.

**Endpoint:** `DELETE /cart/items/:id`

**Authentication:** Required (User)

**Headers:**
```
Authorization: Bearer {user_token}
```

**URL Parameters:**
- `id` (integer) - Cart item ID

**Example Request:**
```bash
curl -X DELETE http://localhost:8080/cart/items/1 \
  -H "Authorization: Bearer {user_token}"
```

**Success Response:** `200 OK`
```json
{
  "message": "item removed from cart"
}
```

**Error Response:** `400 Bad Request`
```json
{
  "error": "item not found in cart"
}
```

---

### Clear Cart

Menghapus semua items dari cart.

**Endpoint:** `DELETE /cart`

**Authentication:** Required (User)

**Headers:**
```
Authorization: Bearer {user_token}
```

**Example Request:**
```bash
curl -X DELETE http://localhost:8080/cart \
  -H "Authorization: Bearer {user_token}"
```

**Success Response:** `200 OK`
```json
{
  "message": "cart cleared"
}
```

---

## Orders

### Create Order (Checkout)

Checkout: membuat order dari items yang ada di cart. Proses ini akan:
- Create order baru
- Mengurangi stock products
- Mengosongkan cart
- Semua dalam 1 transaction (rollback jika ada error)

**Endpoint:** `POST /orders`

**Authentication:** Required (User)

**Headers:**
```
Authorization: Bearer {user_token}
```

**Example Request:**
```bash
curl -X POST http://localhost:8080/orders \
  -H "Authorization: Bearer {user_token}"
```

**Success Response:** `201 Created`
```json
{
  "id": 1,
  "user_id": 1,
  "total_price": 30000000,
  "status": "pending",
  "items": [
    {
      "id": 1,
      "order_id": 1,
      "product_id": 1,
      "product": {
        "id": 1,
        "name": "iPhone 15 Pro",
        "price": 15000000,
        "stock": 8
      },
      "quantity": 2,
      "price": 15000000,
      "created_at": "2025-01-04T11:00:00Z",
      "updated_at": "2025-01-04T11:00:00Z"
    }
  ],
  "created_at": "2025-01-04T11:00:00Z",
  "updated_at": "2025-01-04T11:00:00Z"
}
```

**Notes:**
- `price` di order item adalah price saat pembelian (bisa beda dengan current price)
- `total_price` adalah sum dari semua items
- Default status adalah `pending`

**Error Responses:**

`400 Bad Request` - Cart empty
```json
{
  "error": "cart is empty"
}
```

`400 Bad Request` - Insufficient stock
```json
{
  "error": "insufficient stock for: iPhone 15 Pro"
}
```

`404 Not Found` - Cart not found
```json
{
  "error": "cart not found"
}
```

---

### Get Order History

Mendapat semua orders milik user yang sedang login.

**Endpoint:** `GET /orders`

**Authentication:** Required (User)

**Headers:**
```
Authorization: Bearer {user_token}
```

**Example Request:**
```bash
curl -X GET http://localhost:8080/orders \
  -H "Authorization: Bearer {user_token}"
```

**Success Response:** `200 OK`
```json
[
  {
    "id": 2,
    "user_id": 1,
    "total_price": 38500000,
    "status": "shipped",
    "items": [...],
    "created_at": "2025-01-04T12:00:00Z",
    "updated_at": "2025-01-04T14:00:00Z"
  },
  {
    "id": 1,
    "user_id": 1,
    "total_price": 30000000,
    "status": "delivered",
    "items": [...],
    "created_at": "2025-01-04T11:00:00Z",
    "updated_at": "2025-01-04T15:00:00Z"
  }
]
```

**Notes:**
- Orders sorted by `created_at` descending (newest first)
- Include full order items dengan product details

---

### Get Order Detail

Mendapat detail order spesifik berdasarkan ID.

**Endpoint:** `GET /orders/:id`

**Authentication:** Required (User)

**Headers:**
```
Authorization: Bearer {user_token}
```

**URL Parameters:**
- `id` (integer) - Order ID

**Example Request:**
```bash
curl -X GET http://localhost:8080/orders/1 \
  -H "Authorization: Bearer {user_token}"
```

**Success Response:** `200 OK`
```json
{
  "id": 1,
  "user_id": 1,
  "total_price": 30000000,
  "status": "pending",
  "items": [
    {
      "id": 1,
      "order_id": 1,
      "product_id": 1,
      "product": {
        "id": 1,
        "name": "iPhone 15 Pro",
        "price": 15000000
      },
      "quantity": 2,
      "price": 15000000,
      "created_at": "2025-01-04T11:00:00Z"
    }
  ],
  "created_at": "2025-01-04T11:00:00Z",
  "updated_at": "2025-01-04T11:00:00Z"
}
```

**Error Response:** `404 Not Found`
```json
{
  "error": "order not found"
}
```

---

## Admin Operations

### Update Order Status

Admin bisa update status order. Endpoint ini hanya bisa diakses oleh admin.

**Endpoint:** `PUT /admin/orders/:id/status`

**Authentication:** Required (Admin only)

**Headers:**
```
Authorization: Bearer {admin_token}
Content-Type: application/json
```

**URL Parameters:**
- `id` (integer) - Order ID

**Request Body:**
```json
{
  "status": "string (required)"
}
```

**Valid Status Values:**
- `pending` - Order baru dibuat
- `paid` - Payment confirmed
- `shipped` - Order sedang dikirim
- `delivered` - Order sudah sampai
- `cancelled` - Order dibatalkan

**Example Request:**
```bash
curl -X PUT http://localhost:8080/admin/orders/1/status \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer {admin_token}" \
  -d '{
    "status": "paid"
  }'
```

**Success Response:** `200 OK`
```json
{
  "message": "order status updated"
}
```

**Error Responses:**

`403 Forbidden` - User bukan admin
```json
{
  "error": "admin access required"
}
```

`400 Bad Request` - Invalid status
```json
{
  "error": "invalid status"
}
```

---

## Error Responses

### Standard Error Format

Semua error responses menggunakan format:
```json
{
  "error": "error description"
}
```

### HTTP Status Codes

| Code | Meaning | Description |
|------|---------|-------------|
| 200 | OK | Request berhasil |
| 201 | Created | Resource berhasil dibuat |
| 400 | Bad Request | Request invalid (validation error, dll) |
| 401 | Unauthorized | Authentication required atau token invalid |
| 403 | Forbidden | User tidak punya akses ke resource |
| 404 | Not Found | Resource tidak ditemukan |
| 500 | Internal Server Error | Server error |

### Common Error Messages

**Authentication Errors:**
```json
{"error": "authorization header missing"}
{"error": "invalid authorization format"}
{"error": "invalid token"}
{"error": "Invalid email or password"}
```

**Authorization Errors:**