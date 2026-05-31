# Volte Server

Go backend for Volte. It provides the HTTP API, session-based authentication, MongoDB/Redis persistence, Ethereum contract integration, and Groth16 proof generation helpers for the voting flow.

## Responsibilities

- Manage users, sessions, voting events, event membership, vote submission, and tally retrieval.
- Store application data in MongoDB and short-lived vote/session data in Redis.
- Build event membership Merkle trees and expose membership proof details.
- Generate Groth16 keys/proofs for ballot, membership, and nullifier circuits.
- Deploy and interact with the Volte smart contracts when not using test mode.

## Requirements

- Go `1.24.3` or compatible.
- Docker and Docker Compose for local MongoDB and Redis.
- Ethereum JSON-RPC endpoint, wallet private key, and deployed contract address when running with `--test=false`.

## Setup

```bash
cp .env.example .env
cp contracts.env.example contracts.env
cp circuits.env.example circuits.env
docker compose up -d mongo redis
go mod download
```

Update `.env` with local service values. For the provided Docker Compose services, MongoDB runs on `localhost:27017` with user `root` and password `mongo_pass`, and Redis runs on `localhost:6379`.

## Run

```bash
go run main.go start $(cat .env | tr '\n' ' ')
```

The same command is available through the make target:

```bash
make -f MakeFile run
```

By default, the server listens on `0.0.0.0:8000` and uses the fake contract handler because `--test` defaults to `true`.

To use real contracts, include chain settings in `.env` and start with `--test=false`:

```bash
go run main.go start $(cat .env | tr '\n' ' ') --test=false
```

## Configuration

The server uses Go flags, and the env example files are formatted as command-line flag arguments.

- `.env` configures the API server, sessions, databases, CORS, and contract connection.
- `contracts.env` configures contract deployment/export commands.
- `circuits.env` configures proof generation inputs such as Merkle paths, event IDs, and elliptic curve points.

Common flags:

| Flag | Purpose |
| --- | --- |
| `--host` | API bind host. Defaults to `0.0.0.0`. |
| `--port` | API bind port. Defaults to `8000`. |
| `--allow_origins` | Comma-separated CORS origins. |
| `--session_secret` | Cookie session signing secret. |
| `--mongo_username`, `--mongo_password`, `--mongo_host` | MongoDB connection settings. |
| `--redis_host`, `--redis_port`, `--redis_pass`, `--redis_username`, `--redis_db` | Redis connection settings. |
| `--wallet_private_key`, `--chain_rpc_node_url`, `--contract_address` | Ethereum integration settings. |
| `--test` | Use fake contract handler when `true`; use Ethereum contracts when `false`. |

## API Routes

| Method | Route | Description |
| --- | --- | --- |
| `POST` | `/auth/signup` | Register a user and create a session. |
| `POST` | `/auth/login` | Authenticate a user and create a session. |
| `POST` | `/users/events` | Create an event for the logged-in admin. |
| `GET` | `/users/events` | List events associated with the logged-in user. |
| `GET` | `/users/event/:event_id` | Fetch one user event. |
| `POST` | `/event/:id/members/:commitment` | Add an eligible voter commitment to an event. |
| `DELETE` | `/event/:id/members/:commitment` | Remove an eligible voter commitment from an event. |
| `POST` | `/event/:id/start` | Start an event and build membership proofs. |
| `POST` | `/event/:id/vote` | Submit a vote with ZK proof data. |
| `POST` | `/event/:id/end` | End an event. |
| `GET` | `/event/:id/tally` | Fetch the event tally. |
| `GET` | `/event/:id/membership/merkle` | Fetch membership Merkle proof details. |
| `DELETE` | `/event/:id` | Delete an event. |

## CLI Commands

Run commands from this directory with `go run main.go <command>`.

- `start` starts the HTTP server.
- `generate-keys` creates Groth16 proving and verifying keys under `cmd/proof/keys/groth16/`.
- `export-solidity` exports Solidity verifiers to `../contracts/groth16/`.
- `deploy-contract` deploys the Volte contract stack using `contracts.env` style chain flags.
- `random-proof` generates sample ballot, membership, and nullifier proofs from circuit flags.
- `random-basepoint` derives BabyJubJub public points from `--privkey32`.

Examples:

```bash
go run main.go generate-keys
go run main.go export-solidity
go run main.go deploy-contract $(cat contracts.env | tr '\n' ' ')
go run main.go random-proof $(cat circuits.env | tr '\n' ' ')
```

## Tests

```bash
go test ./...
```

Some tests use Docker-backed dependencies through testcontainers, so Docker should be running before executing the full test suite.
