# 🧸 Toko Mainan (RESTful API in Golang)

Hey there!  
This is a beginner-friendly RESTful API project built using pure Go (`net/http`) with no external frameworks.

I made this to learn about RESTful API
---

## 📌 What is this?

A simple API to manage toy data — like an online toy catalog.  
It covers core CRUD features using only the standard library in Go.

---

## 💡 Features

| Endpoint        | Method | Description                           |
|----------------|--------|---------------------------------------|
| `/toys`        | GET    | Fetch all toys                        |
| `/toys`        | POST   | Add a new toy                         |
| `/toys/{id}`   | GET    | Fetch toy details by ID               |
| `/toys/{id}`   | DELETE | Delete a toy by ID                    |

---

## 🚀 How to Run

1. Clone this repo
2. Make sure you have Go installed (v1.20+ recommended)
3. Run this command in terminal:

```bash
go run main.go
```
Server will be running at: http://localhost:8080

🧪 Test with CURL
List all toys:
```bash
curl http://localhost:8080/toys
```

Add a toy:
```bash
curl -X POST http://localhost:8080/toys \
  -H "Content-Type: application/json" \
  -d '{"nama": "Digimon Agumon", "brand": "Bandai"}'
```

Get a toy by ID:
```bash
curl http://localhost:8080/toys/3
```

Delete a toy:
```bash
curl -X DELETE http://localhost:8080/toys/2
```

## 💬 Why this repo?
This repo was created as part of my learning journey to understand how RESTful APIs work using Go.
It’s a simple practice project — not built to impress, but to progress.
Just learning one step at a time 🚀

## 📜 License

MIT License — feel free to fork, explore, and remix. just give me credit 😄

## 🧘 Final Words

If you find anything useful here, that’s amazing.
If not, that’s totally fine too — this is just part of my dev journey ✌️
