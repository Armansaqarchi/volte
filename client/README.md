# ZK Vote - Zero-Knowledge Voting Application

A privacy-preserving voting application built with Next.js and designed for Groth16 zero-knowledge proof integration.

## Features

- **User Authentication**: Sign up and login with email/password
- **Event Management**: Create voting events with custom options
- **Private Voting**: Vote using zero-knowledge proofs (Groth16)
- **Real-time Results**: View voting results while maintaining voter anonymity

## Getting Started

### Prerequisites

- Node.js 18+ 
- Your Groth16 circuit compiled and ready

### Installation

1. Clone the repository
2. Install dependencies: `npm install`
3. Run the development server: `npm run dev`
4. Open [http://localhost:3000](http://localhost:3000)

## ZK Proof Integration

This application includes **clear integration points** for your Groth16 zero-knowledge proof system. You need to implement the actual cryptographic operations.

### Integration Points

#### 1. **Proof Generation** (`lib/zk-proof.ts`)
Replace the `generateGroth16Proof()` function with your actual proof generation logic:

\`\`\`typescript
export async function generateGroth16Proof(params: ProofGenerationParams): Promise<ZKProof> {
  // Your snarkjs/circom proof generation here
  const { proof, publicSignals } = await snarkjs.groth16.fullProve(
    { /* your inputs */ },
    'circuit.wasm',
    'proving_key.zkey'
  )
  return { proof, publicSignals, timestamp: Date.now() }
}
\`\`\`

#### 2. **Proof Verification** (`lib/zk-proof.ts`)
Replace the `verifyGroth16Proof()` function with your verification logic:

\`\`\`typescript
export async function verifyGroth16Proof(params: ProofVerificationParams): Promise<boolean> {
  // Your snarkjs verification here
  return await snarkjs.groth16.verify(
    params.verificationKey,
    params.publicSignals,
    params.proof
  )
}
\`\`\`

#### 3. **Voting Interface** (`components/voting-interface.tsx`)
The voting component has two main functions you can customize:
- `handleGenerateProof()` - Called when user generates their ZK proof
- `handleSubmitVote()` - Called when user submits their vote with proof

Check the console logs for debugging information during development.

### What the ZK Proof Should Demonstrate

Your Groth16 circuit should prove:

1. ✅ User is eligible to vote (e.g., merkle proof of inclusion)
2. ✅ User hasn't voted before (using nullifiers)
3. ✅ Vote is for a valid option in the event

**Without revealing:**
- ❌ Which specific user is voting
- ❌ Which option they selected

### Recommended Circuit Structure

**Public Inputs:**
- Event ID
- Nullifier (prevents double voting)
- Merkle root (of eligible voters)

**Private Inputs:**
- User ID / Secret
- Vote selection
- Eligibility proof
- Randomness

## Current Implementation

The app currently uses **localStorage** for demo purposes. In production, you should:

1. Replace auth system with proper backend (e.g., JWT, OAuth)
2. Store events and votes in a database
3. Implement backend verification of ZK proofs
4. Add proper nullifier checking to prevent double voting
5. Use merkle trees for efficient eligibility verification

## Project Structure

\`\`\`
app/
├── page.tsx                    # Landing page
├── login/page.tsx             # Login page
├── signup/page.tsx            # Signup page
├── dashboard/
│   ├── page.tsx               # Dashboard overview
│   └── events/
│       ├── create/page.tsx    # Create event form
│       └── [id]/page.tsx      # Event detail & voting
components/
├── voting-interface.tsx       # Main voting UI (ZK integration here)
├── dashboard-nav.tsx          # Navigation component
└── ui/                        # UI components
lib/
├── auth.ts                    # Auth utilities
├── events.ts                  # Event management
├── votes.ts                   # Vote recording
└── zk-proof.ts               # 🔐 ZK PROOF INTEGRATION MODULE
\`\`\`

## Tech Stack

- **Framework**: Next.js 16 (App Router)
- **Styling**: Tailwind CSS v4
- **UI Components**: shadcn/ui
- **Icons**: Lucide React
- **ZK Proofs**: (Your Groth16 implementation)

## Development Notes

- Check browser console for `[v0]` and `[ZK]` debug logs
- The app shows helpful developer notes in the UI during development
- All ZK integration points are clearly marked with ⚠️ comments

## Next Steps

1. Implement your Groth16 circuit
2. Compile your circuit and generate proving/verification keys
3. Replace mock functions in `lib/zk-proof.ts`
4. Test proof generation and verification
5. Deploy with proper backend infrastructure

## Security Considerations

- Never store plain passwords in production
- Verify all ZK proofs on the backend
- Implement proper rate limiting
- Use HTTPS in production
- Secure your proving keys appropriately
- Implement proper session management

---

Built with v0 by Vercel
