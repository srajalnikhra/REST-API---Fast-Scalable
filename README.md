# Students API – Fast & Scalable | GoLang, SQLite, ServeMux (Go 1.22+)

![Go](https://img.shields.io/badge/Go-1.22+-blue)
![API](https://img.shields.io/badge/Type-REST%20API-green)
![Database](https://img.shields.io/badge/Database-SQLite-green)
![License](https://img.shields.io/badge/License-MIT-yellow)

🚀 **Students API** is a fast and scalable RESTful API built with **GoLang**, **SQLite**, and Go’s modern **`http.ServeMux` (Go 1.22+)**.  
It provides an efficient backend for managing student data with clean routing, modular structure, structured logging, and easy configuration.  

---

## 📌 Features  
- ✅ RESTful API with **CRUD operations** for student data 📚  
- ✅ **ServeMux (Go 1.22+)** for clean routing with path parameters ⚡  
- ✅ **SQLite integration** for lightweight and persistent storage 🗄️  
- ✅ **Structured logging (`slog`)** for better observability 📊  
- ✅ Configurable using `local.yaml` 🔧  
- ✅ Modular design for scalability and maintainability 🏗️  
- ✅ **Graceful server shutdown** with context ⏹️  

---

## ⚙️ Tech Stack  
- **Backend:** GoLang (net/http, ServeMux, slog)  
- **Database:** SQLite  
- **Configuration:** YAML (`local.yaml`)  
- **Architecture:** REST API with clean separation of concerns (Config, HTTP, Storage, Utils)  

---

## 📂 Project Structure  

```plaintext
REST-API---Fast-Scalable/
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
- Go 1.22+  
- SQLite installed  

## Setup & Run
### Clone the repository
```bash
git clone https://github.com/srajalnikhra/REST-API---Fast-Scalable.git
```

```bash
cd REST-API---Fast-Scalable
```

### Install dependencies
```bash
go mod tidy
```

### Run the API
```bash
go run main.go
```

## 💡 Future Enhancements

- 🔹 Add authentication & JWT support 🔐
- 🔹 Migrate to PostgreSQL/MongoDB for scalability 🗄️
- 🔹 Add Docker support for containerized setup 🐳
- 🔹 Unit & integration testing 🧪

## 📜 License

This project is licensed under the **MIT License**.
