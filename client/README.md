# Volte Client

Next.js frontend for Volte, a zero-knowledge voting application. The client handles account onboarding, event management, member administration, vote submission, and result viewing against the Go backend.

## Features

- Account signup and login using a locally generated BN254 private key.
- MiMC7 commitment generation in the browser.
- Dashboard for creating and viewing voting events.
- Event detail pages for adding/removing eligible voters, starting/ending events, voting, and viewing tally results.
- Browser-side Groth16 proof helper code for ballot, membership, and nullifier circuits.
- Tailwind CSS and shadcn-style UI components.

## Requirements

- Node.js 18+.
- npm, or pnpm if you prefer the included `pnpm-lock.yaml`.
- Volte backend running at `http://localhost:8000`.

## Setup

```bash
npm install
npm run dev
```

Open `http://localhost:3000`.

The client currently calls the backend with hard-coded `http://localhost:8000` URLs. Start the backend before using signup, login, event, or voting flows.

## Scripts

```bash
npm run dev      # Start the Next.js dev server
npm run build    # Build for production
npm run start    # Start a production build
npm run lint     # Run ESLint
```

## App Routes

| Route | Purpose |
| --- | --- |
| `/` | Landing page. |
| `/signup` | Generates a private key, registers a user, and stores the current user locally. |
| `/login` | Hashes the provided private key into a commitment and authenticates with the backend. |
| `/dashboard` | Lists the current user's events and active events. |
| `/dashboard/events/create` | Creates a new voting event. |
| `/dashboard/events/[id]` | Manages membership, starts/ends events, submits votes, and displays results. |

## Project Structure

```text
client/
├── app/                    # Next.js App Router pages
├── components/             # Shared UI and feature components
├── hooks/                  # UI hooks
├── lib/
│   ├── auth.ts             # User storage, private key generation, MiMC7 commitments
│   ├── events.ts           # Event API helpers
│   └── zkproof/            # Groth16 proof helpers and WASM runtime support
├── public/                 # Static assets and WASM artifacts
└── styles/                 # Global styles
```

## Backend Integration

The UI uses cookie-backed sessions and sends requests with `credentials: "include"`.

Main backend endpoints used by the client:

- `POST /auth/signup`
- `POST /auth/login`
- `GET /users/events`
- `GET /users/event/:event_id`
- `POST /users/events`
- `POST /event/:id/members/:commitment`
- `DELETE /event/:id/members/:commitment`
- `POST /event/:id/start`
- `POST /event/:id/vote`
- `POST /event/:id/end`
- `GET /event/:id/tally`
- `GET /event/:id/membership/merkle`
- `DELETE /event/:id`

## ZK Proof Notes

Proof-related code lives under `lib/zkproof/`.

- `lib/zkproof/circom/ballot/generate_proof.mjs` builds an encrypted ballot and Groth16 ballot proof.
- `lib/zkproof/circom/merkletree/generator.mjs` builds a membership proof from Merkle path inputs.
- `lib/zkproof/circom/nullifier/generateNullifier.mjs` builds the nullifier proof.
- `lib/zkproof/go/` contains support for loading Go-compiled WASM from `public/proof.wasm`.
- `lib/zkproof/zk-proof.ts` is an older/mock integration surface kept as a reference.

The Circom helpers expect proof artifacts at these public paths:

```text
public/zk/ballot/ballot.wasm
public/zk/ballot/ballot_final.zkey
public/zk/merkletree/merkle.wasm
public/zk/merkletree/merkle_final.zkey
public/zk/nullifier/nullifier.wasm
public/zk/nullifier/nullifier_final.zkey
```

Generate or copy these artifacts from the circuit subproject before using browser-side proof generation.

## Security Notes

- The signup page displays a private key once; users must save it securely.
- The private key is used client-side to derive a MiMC7 commitment.
- `localStorage` is used for current-user convenience and is not a substitute for hardened production auth.
- Backend proof verification and contract checks are required for production safety.
