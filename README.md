# Elelden Backend

Modern, scalable marketplace backend API built with Go, following Clean Architecture principles.

## 🚀 Features

- **Clean Architecture** - Well-structured, maintainable codebase
- **JWT Authentication** - Secure user authentication and authorization
- **Category-based Listings** - Support for Emlak, Araç, and İkinci El categories
- **RESTful API** - Standard HTTP methods and status codes
- **MySQL Integration** - Robust database operations with GORM
- **Docker Support** - Containerized development and deployment
- **Migration System** - Database schema versioning
- **Comprehensive Validation** - Input validation and error handling

## 🛠 Tech Stack

- **Language**: Go 1.21+
- **Framework**: Gin Web Framework
- **ORM**: GORM
- **Database**: MySQL 8.0
- **Authentication**: JWT
- **Containerization**: Docker & Docker Compose
- **Password Hashing**: bcrypt

## 📁 Project Structure

```
backend/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/            # HTTP request handlers
│   │   ├── middleware/          # Middleware functions
│   │   │   └── auth.go         # JWT authentication middleware
│   │   └── routes/             # Route definitions
│   │       └── routes.go
│   ├── config/
│   │   └── config.go           # Configuration management
│   ├── domain/
│   │   ├── entities/           # Domain models
│   │   │   └── models.go
│   │   └── services/           # Business logic services
│   └── repository/
│       └── mysql/              # Data access layer
├── migrations/
│   └── 001_initial_schema.sql  # Database migrations
├── pkg/
│   ├── auth/
│   │   └── jwt.go              # JWT utilities
│   └── utils/
│       └── password.go         # Password utilities
├── .env.example                # Environment variables template
├── Dockerfile                  # Docker configuration
├── go.mod                      # Go module definition
└── go.sum                      # Go module checksums
```

## 🗄️ Database Schema

### Users Table
```sql
CREATE TABLE users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### Categories Table
```sql
CREATE TABLE categories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Listings Table
```sql
CREATE TABLE listings (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    title VARCHAR(200) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    location VARCHAR(100),
    details JSON,                    -- Category-specific fields
    images JSON,                     -- Array of image URLs
    status ENUM('active', 'sold', 'inactive') DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
```

## 📋 API Endpoints

### Authentication
```
POST   /api/v1/auth/register     # User registration
POST   /api/v1/auth/login        # User login
POST   /api/v1/auth/refresh      # Refresh JWT token
GET    /api/v1/auth/profile      # Get user profile (Protected)
PUT    /api/v1/auth/profile      # Update user profile (Protected)
```

### Categories
```
GET    /api/v1/categories        # List all categories
POST   /api/v1/categories        # Create category (Admin)
GET    /api/v1/categories/:id    # Get category details
PUT    /api/v1/categories/:id    # Update category (Admin)
DELETE /api/v1/categories/:id    # Delete category (Admin)
```

### Listings
```
GET    /api/v1/listings          # List listings (with filters)
POST   /api/v1/listings          # Create listing (Protected)
GET    /api/v1/listings/:id      # Get listing details
PUT    /api/v1/listings/:id      # Update listing (Protected, Owner only)
DELETE /api/v1/listings/:id      # Delete listing (Protected, Owner only)
GET    /api/v1/listings/user/:id # Get user's listings
```

### Query Parameters for Listings
```
?category_id=1           # Filter by category
?location=istanbul       # Filter by location
?min_price=1000         # Minimum price filter
?max_price=5000         # Maximum price filter
?page=1                 # Pagination
?limit=20               # Items per page
?sort=price_asc         # Sort by price ascending
?sort=price_desc        # Sort by price descending
?sort=date_desc         # Sort by creation date (default)
```

## 🏗️ Category-Specific Fields

### Emlak (Real Estate)
```json
{
  "konut_tipi": "daire|villa|mustakil_ev",
  "oda_sayisi": "1+1|2+1|3+1|4+1|5+1",
  "metrekare": 120,
  "kat": 3,
  "bina_yasi": 5,
  "iskan": "var|yok",
  "site_icerisinde": true,
  "balkon": true,
  "asansor": true,
  "otopark": true
}
```

### Araç (Vehicles)
```json
{
  "marka": "Toyota",
  "model": "Corolla",
  "yil": 2020,
  "kilometre": 50000,
  "motor_hacmi": "1.6",
  "yakit_tipi": "benzin|dizel|lpg|elektrik|hibrit",
  "vites": "manuel|otomatik",
  "motor_gucu": 132,
  "renk": "beyaz",
  "hasar_kaydi": "yok|var",
  "degisen_parca": "yok|var",
  "tramer_kaydi": "yok|var"
}
```

### İkinci El (Second Hand)
```json
{
  "alt_kategori": "elektronik|giyim|ev_esyasi|kitap_dergi|spor_outdoor",
  "marka": "Apple",
  "model": "iPhone 13",
  "kullanim_durumu": "sifir|cok_az_kullanilmis|az_kullanilmis|orta|cok_kullanilmis",
  "garanti": "var|yok",
  "garanti_suresi": 12,
  "orijinal_kutu": true,
  "aksesuar": true
}
```

## 🔧 Installation & Setup

### Prerequisites
- Go 1.21 or higher
- MySQL 8.0
- Docker (optional)

### Local Development

1. **Clone the repository**
```bash
git clone https://github.com/ByGultekin-tak/elelden-backend.git
cd elelden-backend
```

2. **Install dependencies**
```bash
go mod tidy
```

3. **Set up environment variables**
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. **Environment Variables**
```env
# Server
PORT=8080
GIN_MODE=release

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=elelden_db

# JWT
JWT_SECRET=your_super_secret_jwt_key
JWT_EXPIRE_HOURS=24

# CORS
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

5. **Set up MySQL database**
```bash
# Create database
mysql -u root -p
CREATE DATABASE elelden_db;
```

6. **Run migrations**
```bash
# Execute the migration file
mysql -u root -p elelden_db < migrations/001_initial_schema.sql
```

7. **Run the application**
```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

### Docker Development

1. **Using Docker Compose**
```bash
# Start all services (MySQL + API)
docker-compose -f docker-compose.dev.yml up -d

# View logs
docker-compose -f docker-compose.dev.yml logs -f api

# Stop services
docker-compose -f docker-compose.dev.yml down
```

2. **Building Docker image manually**
```bash
# Build image
docker build -t elelden-backend .

# Run container
docker run -p 8080:8080 --env-file .env elelden-backend
```

## 🧪 Testing the API

### Using curl

**Register a new user:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "full_name": "Test User",
    "phone": "+90555123456"
  }'
```

**Login:**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

**Create a listing:**
```bash
curl -X POST http://localhost:8080/api/v1/listings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "category_id": 1,
    "title": "Satılık Daire",
    "description": "Merkezi konumda satılık daire",
    "price": 250000,
    "location": "İstanbul, Kadıköy",
    "details": {
      "konut_tipi": "daire",
      "oda_sayisi": "2+1",
      "metrekare": 90,
      "kat": 3
    }
  }'
```

**Get listings:**
```bash
curl http://localhost:8080/api/v1/listings?category_id=1&page=1&limit=10
```

## 📚 Development Guidelines

### Code Structure
- Follow Clean Architecture principles
- Keep business logic in `internal/domain/services`
- Database operations in `internal/repository`
- HTTP handlers in `internal/api/handlers`

### Error Handling
```go
// Use structured error responses
type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message"`
    Code    int    `json:"code"`
}
```

### Database Transactions
```go
// Use transactions for multiple operations
tx := db.Begin()
defer func() {
    if r := recover(); r != nil {
        tx.Rollback()
    }
}()

if err := tx.Create(&listing).Error; err != nil {
    tx.Rollback()
    return err
}

return tx.Commit().Error
```

## 🐳 Production Deployment

### Docker Production
```bash
# Build production image
docker build -f Dockerfile.prod -t elelden-backend:prod .

# Run with production settings
docker run -d \
  --name elelden-backend \
  -p 8080:8080 \
  --env-file .env.prod \
  elelden-backend:prod
```

### Environment Variables for Production
```env
GIN_MODE=release
DB_HOST=your_production_db_host
JWT_SECRET=your_super_secure_production_secret
CORS_ALLOWED_ORIGINS=https://yourdomain.com
```

## 🔍 Monitoring & Logging

The application uses structured logging with different levels:
- `ERROR`: Critical errors that need immediate attention
- `WARN`: Warning messages for potential issues
- `INFO`: General application flow information
- `DEBUG`: Detailed debugging information (development only)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Support

For support and questions:
- Create an issue on GitHub
- Contact: [your-email@example.com]

---

**Built with ❤️ using Go and Clean Architecture**
