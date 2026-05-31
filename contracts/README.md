# Volte Contracts

Solidity contracts and local chain tooling for Volte. This subproject contains the main voting contract, generated Groth16 verifier contracts, compiled artifacts, a local geth development chain, and Hardhat test/compile scaffolding.

## Structure

```text
contracts/
├── volte.sol            # Main Volte voting contract
├── groth16/             # Generated Solidity verifiers for ballot, membership, and nullifier proofs
├── build/               # Compiled ABI/bin artifacts
├── chain/               # Local geth dev chain configuration
└── test/
    └── hardhat/         # Hardhat project for compiling/testing contracts
```

## Contract Overview

`volte.sol` defines `VolteContract`, which coordinates the on-chain voting checks.

- Stores event metadata hashes and vote Merkle roots by event ID.
- Verifies three Groth16 proofs for each vote:
  - ballot proof for encrypted vote validity,
  - membership proof for eligible voter inclusion,
  - nullifier proof for double-vote prevention inputs.
- Accumulates encrypted tally points using BabyJubJub point addition.
- Exposes event hash, Merkle root, encrypted tally, and total vote getters.

The main contract imports generated verifier contracts from:

- `groth16/ballot.sol`
- `groth16/membership.sol`
- `groth16/nullifier.sol`

## Requirements

- Node.js and npm for Hardhat tooling.
- Docker for the local geth development chain.
- Generated verifier contracts from the server or circuit tooling when regenerating proof systems.

## Local Chain

Start the local geth chain:

```bash
cd chain
docker compose up
```

The development RPC is exposed at `http://localhost:8545`, with WebSocket RPC at `ws://localhost:8546`.

The dev chain imports the standard Hardhat-style private key configured in `chain/start-geth.sh`.

## Hardhat Compile/Test Project

The Hardhat project lives in `test/hardhat`.

```bash
cd test/hardhat
npm install
./compile.sh
npx hardhat test
```

`compile.sh` copies `../../groth16` and `../../volte.sol` into the Hardhat `contracts/` directory before running `npx hardhat clean && npx hardhat compile`.

## Regenerating Verifiers

The server can export Solidity verifiers from generated Groth16 keys:

```bash
cd ../server
go run main.go generate-keys
go run main.go export-soliidity
```

## Deployment

Contract deployment is currently driven from the server CLI:

```bash
cd ../server
go run main.go deploy-contract $(cat contracts.env | tr '\n' ' ')
```

The deployment command deploys the ballot, membership, and nullifier verifier contracts first, then deploys `VolteContract` with those verifier addresses.

## Build Artifacts

Compiled ABIs and bytecode are stored under `build/`:

- `build/volte/` contains `VolteContract` and verifier artifacts.
- `build/groth16/` contains generated verifier artifacts grouped by proof type.

Regenerate these artifacts after changing `volte.sol` or any verifier contract.
