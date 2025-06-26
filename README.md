# Clinic Portal

A modern, secure, and scalable web application built with Go for managing clinic operations. The application provides a **Receptionist Portal** for patient management (CRUD operations) and a **Doctor Portal** for viewing and updating patient medical records, with a single login system for both roles. Designed for a technical test, this project showcases clean architecture, robust security, and user-friendly features.

## Features

### Core Features
- **Single Login System**: Unified authentication for receptionists and doctors using JWT with role-based access control (RBAC).
- **Receptionist Portal**:
  - Register, view, update, and delete patient records.
  - Search and filter patients by name, email, or gender with pagination.
  - Export patient data to CSV for reporting.
  - Audit logging for tracking patient data changes.
- **Doctor Portal**:
  - View detailed patient information, including medical history.
  - Add and update medical records (notes, diagnosis, treatment).
  - Export patient medical records to PDF.
- **RESTful API**: Well-documented endpoints with Swagger for easy testing.

### Additional Features
- **Real-Time Notifications**: WebSocket-based alerts for doctors when new patients are registered.
- **Dashboard**: Analytics for receptionists (e.g., total patients, new patients today) and doctors (e.g., assigned patients).
- **File Upload**: Doctors can upload medical documents (e.g., lab results) with size and type validation.
- **Security**:
  - Password hashing with bcrypt.
  - Rate limiting on login endpoint to prevent brute force attacks.
  - Input validation to prevent SQL injection and XSS.
- **Testing**: Unit and integration tests with ~80% coverage.
- **Deployment**: Dockerized application with `docker-compose` for easy setup.

## Tech Stack
- **Backend**: Go 1.20, Gorilla Mux (routing), GORM (ORM)
- **Database**: PostgreSQL
- **Authentication**: JWT (golang-jwt/jwt), bcrypt
- **WebSocket**: gorilla/websocket
- **Documentation**: Swagger (swaggo/swag)
- **Testing**: testify, Go's testing package
- **Deployment**: Docker, docker-compose
- **Logging**: Logrus (structured logging)
- **Others**: godotenv (environment config), go-playground/validator (input validation)

## Project Structure
```
clinic-portal/
├── cmd/
│   └── app/
│       └── main.go              # Application entry point
├── internal/
│   ├── auth/                    # Authentication logic (login, JWT)
│   ├── patient/                 # Patient CRUD operations
│   ├── doctor/                  # Doctor-specific logic (medical records)
│   └── middleware/              # JWT, rate limiting, logging middleware
├── pkg/
│   ├── config/                  # Environment configuration
│   └── logger/                  # Structured logging setup
├── api/
│   └── docs/                    # Swagger API documentation
├── scripts/
│   └── init.sql                 # Database schema initialization
├── Dockerfile                   # Docker configuration
├── docker-compose.yml           # Docker Compose for app and DB
├── .env.example                 # Example environment variables
├── go.mod                       # Go module dependencies
├── go.sum                       # Go module checksums
└── README.md                    # Project documentation
```

## Prerequisites
- **Go**: Version 1.20 or later
- **Docker**: For running the application and database
- **PostgreSQL**: For local database (optional if using Docker)
- **curl** or **Postman**: For testing API endpoints

## Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/clinic-portal.git
cd clinic-portal
```

### 2. Configure Environment Variables
Copy the example `.env` file and update the values:
```bash
cp .env.example .env
```

Example `.env`:
```plaintext
DB_HOST=localhost
DB_PORT=5439
DB_USER=postgres
DB_PASSWORD=secret
DB_NAME=clinic_db
JWT_SECRET=your_jwt_secret_key
PORT=8080
```

### 3. Run with Docker
Start the application and PostgreSQL database:
```bash
docker-compose up --build
```

- API will be available at `http://localhost:8080`.
- Swagger documentation at `http://localhost:8080/swagger/index.html`.

### 4. Run Locally (Without Docker)
1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Initialize the database (ensure PostgreSQL is running):
   ```bash
   psql -U postgres -d clinic_db -f scripts/init.sql
   ```
3. Run the application:
   ```bash
   go run cmd/app/main.go
   ```

### 5. Seed Initial Data
Insert default users for testing:
- Receptionist: `email: receptionist@clinic.com`, `password: password123`
- Doctor: `email: doctor@clinic.com`, `password: password123`

Run the following SQL in your database:
```sql
INSERT INTO users (id, email, password, role, created_at, updated_at)
VALUES
  ('uuid1', 'receptionist@clinic.com', '$2a$10$...', 'receptionist', NOW(), NOW()),
  ('uuid2', 'doctor@clinic.com', '$2a$10$...', 'doctor', NOW(), NOW());
```

## API Endpoints
Below are key endpoints. Full documentation is available via Swagger at `/swagger/index.html`.

| Method | Endpoint                     | Role           | Description                              |
|--------|------------------------------|----------------|------------------------------------------|
| POST   | `/api/v1/auth/login`         | Public         | Authenticate user and return JWT         |
| POST   | `/api/v1/patients`           | Receptionist   | Register a new patient                   |
| GET    | `/api/v1/patients`           | Receptionist   | List patients with pagination and search |
| GET    | `/api/v1/patients/{id}`      | Receptionist, Doctor | Get patient details                |
| PUT    | `/api/v1/patients/{id}`      | Receptionist   | Update patient details                   |
| DELETE | `/api/v1/patients/{id}`      | Receptionist   | Delete patient (soft delete)             |
| POST   | `/api/v1/patients/{id}/medical-records` | Doctor | Add medical record                    |
| PUT    | `/api/v1/patients/{id}/medical-records/{record_id}` | Doctor | Update medical record           |
| GET    | `/api/v1/dashboard`          | Receptionist, Doctor | View analytics dashboard            |

### Example Request (Login)
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"receptionist@clinic.com","password":"password123"}'
```

Response:
```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "role": "receptionist",
  "user_id": "uuid1"
}
```

## Testing
Run unit and integration tests:
```bash
go test ./... -v -cover
```

Generate coverage report:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Future Improvements
- Add multi-tenant support for multiple clinics.
- Implement two-factor authentication (2FA).
- Integrate with cloud storage (e.g., AWS S3) for file uploads.
- Add internationalization (i18n) for multi-language support.

## Contributing
Feel free to open issues or submit pull requests for improvements or bug fixes.

## License
This project is licensed under the MIT License.

## Contact
For questions or feedback, reach out to [berrylhamesha@gmail.com](mailto:berrylhamesha@gmail.com).