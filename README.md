load-balancer/
├── go.mod
├── go.sum
├── main.go                 # Entry point, starts the load balancer
├── balancer/               # Contains core load balancer logic
│   ├── balancer.go         # Manages traffic using selected algorithm
│   ├── round_robin.go      # Round Robin implementation
│   ├── least_conn.go       # Least Connections implementation
│   ├── random.go           # Random algorithm
│   ├── ip_hash.go          # IP Hash algorithm
│   ├── weighted_rr.go      # Weighted Round Robin
│   ├── custom_algo.go      # Sixth (custom) algorithm
├── health/                 # Health checks for backend servers
│   └── healthcheck.go
├── config/                 # Configs for servers, ports, etc.
│   └── config.go
├── utils/                  # Common utility functions
│   └── logger.go
├── server/                 # Backend test server (FastAPI not included here, just example Go server if needed)
│   └── mock_backend.go
├── test/                   # Load testing scripts or commands
│   ├── hey_test.sh
│   └── wrk_test.sh
├── Dockerfile              # Dockerfile for the load balancer
├── docker-compose.yml      # Compose file to run balancer + FastAPI backends
└── README.md



| File/Folder                         | Description                                                   |
| ----------------------------------- | ------------------------------------------------------------- |
| `main.go`                           | Initializes config, health checks, and starts HTTP server.    |
| `balancer.go`                       | Routes requests by calling one of the algorithms dynamically. |
| `round_robin.go` → `custom_algo.go` | Each file implements one load balancing strategy.             |
| `healthcheck.go`                    | Pings `/health` endpoint on backends regularly.               |
| `config.go`                         | Parses list of backends and other settings.                   |
| `logger.go`                         | Optional logger utility to format logs.                       |
| `Dockerfile`                        | Builds the load balancer Go image.                            |
| `docker-compose.yml`                | Runs the load balancer and backend FastAPI containers.        |


# 🚀 FastAPI Project

This is a FastAPI-based project. Follow the instructions below to set up and run the project locally.

---

## 📦 Requirements

- uvicorn
- fastapi

---

## ⚙️ Setup Instructions

## Setup FastAPI server

### 1. Create a virtual environment

Go to the server directory, cd inside server directory.

#### On Windows:

```bash
python -m venv venv
venv\Scripts\activate
```

#### On macOS/Linux:

```bash
python3 -m venv venv
source venv/bin/activate
```

### 3. Install dependencies

```bash
pip install -r requirements.txt
```

---

## 🚀 Run the FastAPI Server

Use **uvicorn** to run the project on `localhost` at **port 8001**.

```bash
uvicorn main:app --reload --port 8001
```
To run another server at 8002 port, replace 8001 with 8002 in new CLI

- Replace `main:app` with the correct filename and app instance if different.
  - `main` = the Python file without `.py`
  - `app` = your FastAPI instance (e.g., `app = FastAPI()`)

### Example:

If your file is named `app.py` and contains:
```python
app = FastAPI()
```

Then run:
```bash
uvicorn app:app --reload --port 8001
```

---


## 🧪 Deactivate the Virtual Environment

When you're done:

```bash
deactivate
```
