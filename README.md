# 🚀 Fiber User Auth Platform (Go + Fiber)

A fast and lightweight user authentication platform built with **Go Fiber**, supporting:

- 🔐 User Signup & Login
- ✅ OTP-based User Verification via SMS
- 🧑‍💼 Role management for Buyer/Seller

---

## 📦 Features

- 📝 User Signup with email, password, and phone number
- 🔑 Secure Login with hashed passwords (bcrypt)
- 📲 OTP Verification (via SMS API like Twilio)
- 👥 Role Assignment: Users can become **Buyers** or **Sellers**
- 🧾 Basic JWT-based session/token support
- 🧪 Structured modular code for extensibility

---

## 🧱 Tech Stack

- **Go (Golang)** – Language
- **Fiber** – Fast HTTP web framework
- **GORM** – ORM for PostgreSQL/MySQL (optional)
- **JWT** – JSON Web Tokens for auth
- **Bcrypt** – Password hashing
- **SMS Gateway API** – For OTP delivery (e.g., Twilio, Fast2SMS)

---

## 📁 Project Structure
```
go-fiber-user-auth-microservice/
├── main.go
├── model/
│ └── user.go
├── handler/
│ └── auth.go
├── routes/
│ └── routes.go
├── utils/
│ ├── jwt.go
│ └── otp.go
├── config/
│ └── db.go
└── go.mod```


---

## 🚀 Getting Started

### 1. Clone the Repo
```bash
git clone https://github.com/manupanand-freelence-developer/go-fiber-user-auth-microservice.git
cd go-fiber-user-auth-microservice

2. Initialize Go Modules

go mod tidy

3. Run the App

go run main.go

🛠️ Environment Variables (or .env)

PORT=3000
JWT_SECRET=supersecretkey
DB_URL=postgres://user:pass@localhost:5432/mydb
SMS_API_KEY=your_sms_api_key

Use godotenv if needed for .env loading.
🔐 API Endpoints (Planned)
Method	Endpoint	Description
POST	/signup	Register new user
POST	/login	Login user and get token
POST	/verify	Verify OTP sent via SMS
POST	/role/buyer	Set user as buyer
POST	/role/seller	Set user as seller
🧪 Future Enhancements

    Email + SMS fallback OTP support

    Admin role and dashboard

    Password reset via OTP

    Rate limiting and brute-force protection



MIT License © 2025 [Manu P Anand]
💬 Contact

[www.linkedin.co.in.manupanand]
