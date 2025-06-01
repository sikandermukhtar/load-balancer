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
