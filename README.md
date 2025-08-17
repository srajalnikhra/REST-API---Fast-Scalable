# Students API - Fast & Scalable | GoLang, Gorilla Mux, SQLite

![Go](https://img.shields.io/badge/Go-1.22+-blue)
![Database](https://img.shields.io/badge/Database-SQLite-green)
![License](https://img.shields.io/badge/License-MIT-yellow)

🚀 **Students API** is a fast and scalable RESTful API built with **GoLang**, **Gorilla Mux**, and **SQLite**.  
It provides an efficient backend for managing student data with clean routing, modular structure, and easy configuration.

---

## 📌 Features
✅ RESTful API with **CRUD operations** for student data 📚  
✅ **Gorilla Mux** for routing and middleware support ⚡  
✅ **SQLite integration** for persistent and lightweight storage 🗄️  
✅ Configurable using `local.yaml` 🔧  
✅ Modular design for scalability and maintainability 🏗️  

---

## ⚙️ Tech Stack
🔹 **Backend:** Golang (Gorilla Mux, net/http)  
🔹 **Database:** SQLite  
🔹 **Configuration:** YAML (`local.yaml`)  
🔹 **Architecture:** REST API with clean separation of concerns (Config, HTTP, Storage, Utils)  

---

## 📂 Project Structure

```
STUDENTS-API/
│── config/
│   └── local.yaml             # App configuration file
│
│── internal/
│   ├── config/
│   │   └── config.go          # Config loader
│   │
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go # Student handlers (CRUD)
│   │
│   ├── storage/
│   │   ├── sqlite/
│   │   │   └── sqlite.go      # SQLite DB setup
│   │   ├── types/
│   │   │   └── types.go       # Data models
│   │   └── storage.go         # Storage interface
│   │
│   └── utils/
│       └── response/
│           └── response.go    # Response helpers
│
│── storage/
│   └── storage.db             # SQLite database file
│
│── .gitignore
│── go.mod
│── go.sum
│── main.go                    # Entry point
```



---

## 🚀 Getting Started

### Prerequisites
- Go 1.20+  
- SQLite installed  

### Setup & Run
```bash
# Clone the repository
git clone https://github.com/srajalnikhra/REST-API---Fast-Scalable.git
cd students-api

# Install dependencies
go mod tidy

# Run the API
go run main.go

💡 Future Enhancements

🔹 Add authentication & JWT support 🔐
🔹 Migrate to PostgreSQL/MongoDB for scalability 🗄️
🔹 Add Docker support for containerized setup 🐳
🔹 Unit & integration testing 🧪

📜 License

This project is licensed under the MIT License.


---

This README is **GitHub ready** — with structure, features, API endpoints, and future enhancements.  

Do you want me to also add **badges** (like Go version, build status, license) at the top for a more polished GitHub profile?
