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
- **Dashboard**: Analytics for receptionists (e.g., total patients, new patients today) and doctors (e.g., assigned patients).
- **File Upload**: Doctors can upload medical documents (e.g., lab results) with size and type validation.
- **Security**:
  - Password hashing with bcrypt.
  - Rate limiting on login endpoint to prevent brute force attacks.
  - Input validation to prevent SQL injection and XSS.
- **Testing**: Unit and integration tests with ~80% coverage.

## Tech Stack
- **Backend**: Go 1.20, Gorilla Mux (routing), GORM (ORM)
- **Database**: PostgreSQL
- **Authentication**: JWT (golang-jwt/jwt), bcrypt
- **Documentation**: Swagger (swaggo/swag)
- **Testing**: testify, Go's testing package
- **Logging**: Logrus (structured logging)
- **Others**: godotenv (environment config), go-playground/validator (input validation)
- **Migration**: rubenv/sql-migrate (database migration tool)

## Prerequisites
- **Go**: Version 1.20 or later
- **PostgreSQL**: For local database (optional if using Docker)
- **Postman**: For testing API endpoints

## Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/berrylradianh/makerble-golang-coding-assesment.git
cd makerble-golang-coding-assesment
```

### 2. Configure Environment Variables
Copy the example `.env.json` file and update the values:
```bash
cp .env.example.json .env.json
```

### 3. Running Migration
  1. Make sure you have PostgreSQL running and create a database named `clinic-portal-makerble-golang-test` with public schema.
  2. Make sure you already installed `sql-migrate` by running `go install github.com/rubenv/sql-migrate@latest`.
  3. Run the following command to run migration:
     ```bash
     sql-migrate up -config=dbconfig.yml -env=development
     ```

### 4. Run Locally (Without Docker)
  1. Install dependencies:
      ```bash
      go mod tidy
      ```
  2. Run the application:
      ```bash
      go run .
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
| PUT    | `/api/v1/patients/{id}/medical-records/{record_id}` | Doctor | Update medical record     |
| GET    | `/api/v1/dashboard`          | Receptionist, Doctor | View analytics dashboard           |

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