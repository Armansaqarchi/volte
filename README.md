# Volte

Volte is a privacy-preserving voting project that combines a Next.js client, a Go backend, zero-knowledge proof circuits, and Ethereum smart contracts. The system is designed around Groth16 proofs for voter eligibility, vote validity, and double-vote prevention.

## Project Structure

```text
.
├── client/      # Next.js voting interface
├── server/      # Go API, proof helpers, database access, and chain integration
├── circom/      # Circom versions of the voting circuits and build tooling
└── contracts/   # Solidity contracts, generated verifiers, and contract tests
```

## Subprojects

- `client/` contains the web application. See `client/README.md` for setup and frontend ZK integration notes.
- `server/` contains the backend service, CLI commands, MongoDB/Redis integration, proof generation utilities, and Ethereum contract bindings. See `server/Readme.md`.
- `circom/` contains Circom implementations of the ballot, membership, and nullifier circuits. See `circom/README.md`.
- `contracts/` contains Solidity contract sources and generated verifier contracts used by the voting flow.

## Requirements

- Node.js and npm for `client/`, `circom/`, and contract test tooling.
- Go `1.24.3` or compatible for `server/`.
- Docker for local MongoDB and Redis services used by the backend.
- Circom/snarkjs tooling when working with the Circom circuit subproject.
- An Ethereum JSON-RPC endpoint and wallet key when running the server against real contracts instead of the fake test contract handler.

## Quick Start

### Frontend

```bash
cd client
npm install
npm run dev
```

The client development server runs at `http://localhost:3000` by default.

### Backend

```bash
cd server
cp .env.example .env
docker compose up -d mongo redis
go run main.go start $(cat .env | tr '\n' ' ')
```

The backend API listens on `0.0.0.0:8000` by default. The server defaults to test contract mode, so it can run with the fake contract handler while you develop locally.

### Circuits

```bash
cd circom
npm install
make all
```

After creating build files, you will need to overwrite zkey and wasm file located in `client/lib/zkproof/circom` subproject.


### Contracts:

Solidity contracts for each circuit is created while building the circom files, located in `circom/<circuit>/build`. <br>
Because volte.sol file located in `contracts/volte.sol` uses all the circuit soldity files, you need to replace the main class 
for each circuit to:

ballot ---> BallotVerifier <br>
membership ---> MembershipVerifier <br>
nullifier ---> NullifierVerifier <br>

Then override previous solidity files in `contracts/groth16`


## Changing ECC base points:

changing ecc base points require to generate new G1 and G2 x and y coordinates.
Assuming you have G1 and G2, you need to place them inside the ballot circom circuit.

1 - Edit `/volte/circom/ballot/ballot.circom`  and replace the old coordinates in the main component with the new ones.

2 - Compile the new circom file from `./circom/MakeFile`

3 - After Creating zkey and wasm files, you need to overwrite those in `client/lib/zkproof/circom`

4- Update base points again in generate_proof.mjs


This builds the Circom circuits and exports proof artifacts. See `circom/README.md` for details.

## Voting Flow

1. Users register and receive a public commitment.
2. An authenticated admin creates a voting event.
3. Eligible user commitments are added to the event.
4. Starting the event creates a Merkle tree of eligible voters.
5. Voters generate and submit zero-knowledge proofs for membership, ballot validity, and nullifier uniqueness.
6. The backend and contracts verify proofs, store votes, and expose event tally data.

## Documentation

Each maintained subproject owns its own README. This root README is only an overview and entry point; use the subproject docs for implementation-specific commands and notes.
