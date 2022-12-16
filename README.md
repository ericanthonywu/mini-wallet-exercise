# Mini Wallet Exercise

### Stacks used:
- Go (https://go.dev/doc/install)
- Postgresql (https://www.postgresql.org/download/)
- Redis (https://redis.io/docs/getting-started/installation/)

### Run Instructions
- Make sure to install the stacks
- Copy `.env.example` to the root project with `.env` as the name 
- Set up the backend port (default 8000), database and redis connection at the `.env` file as your pc configuration

### 1. Postgresql
- Install uuid extension by running  `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`

### 2. Go
- Install the dependency by run `go install` 
- To run the project, simply run `go build mini-wallet-exercise`