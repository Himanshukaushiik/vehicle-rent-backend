# 🚗 Vehicle Rent Backend (Go)

A backend system built using:
- Go (Golang)
- Gin Framework
- GORM ORM
- Google Wire (Dependency Injection)
- Clean Architecture

## 📌 Features
- User Management
- Vehicle CRUD APIs
- Rent system
- Clean layered architecture

## 🚀 API Endpoints

### Users
- POST /users
- GET /users

### Vehicles
- POST /vehicles
- GET /vehicles
- GET /vehicles/:id
- DELETE /vehicles/:id

### Rents
- POST /rents
- GET /rents

## 🛠️ Run Project

```bash
go mod tidy
wire
go run main.go
