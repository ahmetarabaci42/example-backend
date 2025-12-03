# E-Commerce Backend Infrastructure with Go and SQLite

## Project Summary

A minimalist web application developed to understand the fundamentals of backend architecture and client-server communication using Go (Golang). The project demonstrates the integration of a static frontend template with a custom-built backend service via HTTP requests.

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: SQLite3
- **Frontend**: HTML/CSS (Template integration)

## Key Features & Workflow

- **Server Configuration**: An HTTP server, built using Go's standard library, handles routing and manages incoming requests.
- **User Registration**: Incoming user data is processed and structurally stored within a users table in an SQLite database.
- **Authentication (Login)**: Login requests trigger a database query to verify credentials. Upon successful validation, the user is redirected to the protected store interface.

## Prerequisites

- Go 1.25.5 or higher
- SQLite3
- A modern web browser

## Installation

1. **Clone the repository** (or navigate to the project directory):
   ```bash
   cd /path/to/example-backend
   ```

2. **Install Go dependencies**:
   ```bash
   go mod download
   ```

3. **Verify SQLite3 driver**:
   The project uses `github.com/mattn/go-sqlite3`. If you encounter issues, ensure CGO is enabled:
   ```bash
   export CGO_ENABLED=1
   ```

## Running the Application

### Step 1: Start the Backend Server

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Run the server:
   ```bash
   go run server.go
   ```

3. The server will:
   - Initialize the SQLite database (`scheduler.db`)
   - Create the `users` table if it doesn't exist
   - Display your local IP address and port information
   - Start listening on `0.0.0.0:3000`

4. You should see output like:
   ```
   ========================================
   API Sunucusu başlatılıyor...
   Local IP: 192.168.1.24
   Port: :3000
   ========================================
   
   🔌 API Endpoint'leri:
      - Register: http://192.168.1.24:3000/register
      - Login: http://192.168.1.24:3000/login
   
   📝 HTML dosyalarını file:// ile açın
      Frontend'den gelen istekler bu sunucuya gönderilecek
   
   Sunucu çalışıyor, API isteklerini bekliyor...
   ========================================
   ```

### Step 2: Configure Frontend

1. Open `frontend/js/scripts.js`
2. Update the `API_BASE_URL` constant with your server's IP address:
   ```javascript
   const API_BASE_URL = 'http://YOUR_IP:3000';
   ```
   Replace `YOUR_IP` with the IP address shown in the server output (e.g., `192.168.1.24`)

### Step 3: Open Frontend Files

Open the HTML files directly in your browser using the `file://` protocol:

- **Login Page**: `file:///path/to/example-backend/frontend/login.html`
- **Register Page**: `file:///path/to/example-backend/frontend/register.html`
- **Home Page**: `file:///path/to/example-backend/frontend/index.html`

## API Endpoints

### POST `/register`

Registers a new user in the database.

**Request Body** (JSON):
```json
{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response** (201 Created):
```
Kayıt Başarılı
```

**Error Responses**:
- `400 Bad Request`: Invalid data format
- `405 Method Not Allowed`: Only POST requests are accepted
- `500 Internal Server Error`: Database error

**Example using cURL**:
```bash
curl -X POST http://YOUR_IP:3000/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","email":"test@example.com","password":"test123"}'
```

### POST `/login`

Authenticates a user and returns user information if credentials are valid.

**Request Body** (JSON):
```json
{
  "username": "john_doe",
  "password": "password123"
}
```

**Success Response** (200 OK):
```json
{
  "message": "Giriş başarılı",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "john@example.com",
    "password": "password123"
  }
}
```

**Error Responses**:
- `400 Bad Request`: Invalid data format
- `401 Unauthorized`: Invalid username or password
- `405 Method Not Allowed`: Only POST requests are accepted

**Example using cURL**:
```bash
curl -X POST http://YOUR_IP:3000/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"test123"}'
```

## Project Structure

```
example-backend/
├── backend/
│   ├── database/
│   │   └── db.go          # Database initialization and operations
│   ├── server.go           # HTTP server and route handlers
│   └── scheduler.db        # SQLite database file (auto-generated)
├── frontend/
│   ├── css/
│   │   ├── loginStyles.css
│   │   └── styles.css
│   ├── js/
│   │   └── scripts.js     # Frontend JavaScript logic
│   ├── assets/
│   ├── index.html          # Main store page
│   ├── login.html          # Login page
│   └── register.html      # Registration page
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
└── README.md              # This file
```

## Database Schema

The `users` table structure:

```sql
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    email TEXT,
    password TEXT
);
```

## Security Considerations

⚠️ **Important**: This is a learning project. The following security improvements are recommended for production use:

1. **Password Hashing**: Currently, passwords are stored in plain text. Use `golang.org/x/crypto/bcrypt` to hash passwords before storing them.

2. **Input Validation**: Add more robust validation for username, email, and password fields.

3. **SQL Injection**: The project uses parameterized queries (prepared statements), which is good, but always validate inputs.

4. **CORS**: The current CORS configuration allows all origins (`*`). Restrict this in production.

5. **HTTPS**: Use HTTPS in production to encrypt data in transit.

## Troubleshooting

### Server won't start
- Ensure port 3000 is not already in use
- Check that Go is properly installed: `go version`
- Verify SQLite3 driver: `go mod download`

### Frontend can't connect to server
- Verify the server is running
- Check that `API_BASE_URL` in `scripts.js` matches the server IP
- Ensure CORS headers are being sent (check browser console)

### Database errors
- Check file permissions for `backend/scheduler.db`
- Verify SQLite3 is installed: `sqlite3 --version`

## Next Steps

💡 **Backend Yolculuğun İçin Bir Sonraki Adım**

Şu anki yapın mantığı kavramak için mükemmel. Ancak bunu teknik olarak "daha düzgün" (secure ve scalable) hale getirmek istersen şu adımı atmalısın:

**Şifreleme (Hashing)**: Şu an muhtemelen şifreleri veritabanına 12345 gibi düz metin (plain text) olarak kaydediyorsun. Backend dünyasında bu büyük bir güvenlik açığıdır.

- **Yapılacak**: Go'nun `golang.org/x/crypto/bcrypt` kütüphanesini kullanarak, şifreyi veritabanına kaydetmeden önce "hash"le (karıştır). Giriş yaparken de kullanıcının girdiği şifreyi bu hash ile karşılaştır.

## License

This project is for educational purposes.

## Author

Developed as a learning project to understand backend architecture fundamentals.
