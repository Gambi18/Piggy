# Piggy - Personal Finance Management Application

A modern web application for tracking personal finances, including income, expenses, savings goals, and budgets.

## 🚀 Features

- **User Authentication**: Secure account creation and login with JWT-based authentication
- **Transaction Management**: Add and track income/expenses with detailed categorization
- **Balance Tracking**: Real-time balance updates and transaction history
- **Savings Goals**: Set and monitor personal savings objectives
- **Budget Management**: Create and track budget limits
- **Responsive Design**: Mobile-friendly interface built with modern UI components

## 🛠 Tech Stack

### Frontend
- **Framework**: Next.js 16.2.3 with React 19.2.4
- **Language**: TypeScript
- **Styling**: TailwindCSS v4
- **HTTP Client**: Axios
- **Notifications**: React Toastify

### Backend
- **Language**: Go 1.26.0
- **Web Framework**: Gin (HTTP router)
- **Database**: PostgreSQL 17
- **Authentication**: JWT (JSON Web Tokens)
- **ORM**: SQLC for type-safe SQL operations
- **Migration**: golang-migrate
- **Logging**: zerolog

### Infrastructure
- **Containerization**: Docker & Docker Compose
- **Database**: PostgreSQL 17-alpine

## 📋 Prerequisites

- Go 1.26.0 or higher
- Node.js 18 or higher
- Docker and Docker Compose
- PostgreSQL (if running locally without Docker)

## 🚀 Quick Start

### Using Docker Compose (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd Piggy
   ```

2. **Start the application**
   ```bash
   docker-compose up -d
   ```

3. **Access the application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8081
   - Database: localhost:5433

### Manual Setup

#### Backend Setup

1. **Navigate to backend directory**
   ```bash
   cd backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

4. **Start database with Docker**
   ```bash
   docker-compose up -d db
   ```

5. **Run database migrations**
   ```bash
   make migrate-up
   ```

6. **Start the backend server**
   ```bash
   go run cmd/main.go
   ```

#### Frontend Setup

1. **Navigate to frontend directory**
   ```bash
   cd frontend
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Start the development server**
   ```bash
   npm run dev
   ```

## 📚 API Documentation

### Authentication Endpoints

#### POST /api/v1/signup
Create a new user account.

**Request Body:**
```json
{
  "username": "johndoe",
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

#### POST /api/v1/login
Authenticate user and receive JWT token.

**Request Body:**
```json
{
  "username": "johndoe",
  "password": "securepassword123"
}
```

**Response:**
```json
{
  "user": {
    "id": "user-uuid",
    "username": "johndoe",
    "name": "John Doe",
    "email": "john@example.com",
    "balance": 1000
  },
  "token": "jwt-token-here"
}
```

### Protected Endpoints (Require Authorization Header)

All protected endpoints require a Bearer token in the Authorization header:
```
Authorization: Bearer <jwt-token>
```

#### POST /api/v1/transactions
Create a new transaction.

**Request Body:**
```json
{
  "amount": 500,
  "type": "saving",
  "reason": "Monthly savings"
}
```

#### GET /api/v1/transactions
Get all transactions for the authenticated user.

#### GET /api/v1/balance
Get current balance and financial summary for the authenticated user.

## 🔧 Configuration

### Environment Variables

#### Backend (.env)
```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5433
DB_NAME=piggy
DB_USER=admin
DB_PASSWORD=2323

# Server Configuration
PORT=8081

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-change-in-production
```

#### Frontend
The frontend automatically connects to `http://localhost:8081` for API calls.

## 🗂 Project Structure

```
Piggy/
├── backend/
│   ├── cmd/                 # Application entry point
│   ├── internal/
│   │   ├── db/             # Database layer (migrations, repo, sqlc)
│   │   ├── handlers/       # HTTP handlers
│   │   ├── middleware/     # Authentication middleware
│   │   ├── models/         # Data models
│   │   └── piggyservice/   # Business logic
│   ├── docker-compose.yaml # Docker configuration
│   └── go.mod              # Go modules
├── frontend/
│   ├── app/                # Next.js app directory
│   ├── components/         # Reusable React components
│   ├── api/                # API client functions
│   └── package.json        # Node.js dependencies
└── README.md               # This file
```

## 🧪 Development

### Database Migrations

To run database migrations:
```bash
cd backend
make migrate-up
```

To rollback migrations:
```bash
make migrate-down
```

### Code Generation

Generate SQLC code after modifying SQL queries:
```bash
cd backend
make sqlc
```

### Testing

Run backend tests:
```bash
cd backend
go test ./...
```

Run frontend tests:
```bash
cd frontend
npm test
```

## 🔒 Security Features

- **JWT Authentication**: Secure token-based authentication
- **Password Hashing**: Bcrypt for secure password storage
- **CORS Protection**: Configurable cross-origin resource sharing
- **Input Validation**: Request payload validation
- **SQL Injection Prevention**: Type-safe SQL operations with SQLC

## 🚀 Deployment

### Production Deployment

1. **Set production environment variables**
2. **Build the frontend**
   ```bash
   cd frontend
   npm run build
   ```
3. **Build the backend**
   ```bash
   cd backend
   go build -o piggy cmd/main.go
   ```
4. **Use production Docker Compose**
   ```bash
   docker-compose -f docker-compose.prod.yaml up -d
   ```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🐛 Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Ensure PostgreSQL is running
   - Check database credentials in `.env`
   - Verify database exists

2. **JWT Token Invalid**
   - Check JWT_SECRET environment variable
   - Ensure token is not expired (24-hour expiry)

3. **CORS Errors**
   - Verify frontend URL is in CORS allowed origins
   - Check API base URL in frontend configuration

### Getting Help

- Check the [Issues](../../issues) page for known problems
- Create a new issue for bugs or feature requests
- Review the API documentation above for endpoint usage

## 🔄 Version History

- **v1.0.0** - Initial release with basic financial tracking features
- Authentication system with JWT
- Transaction management
- Balance tracking
- Docker support
