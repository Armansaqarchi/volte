# Privacy-Preserving Voting System with Zero-Knowledge Proofs

A complete implementation of privacy-preserving voting using Circom ZK-SNARKs, featuring three core circuits for encrypted voting, voter registry membership, and double-voting prevention.

## 🎯 Project Overview

This repository implements a cryptographically secure voting system where:
- **Votes remain encrypted** end-to-end
- **Voter identity is anonymous** but verifiable
- **Double-voting is prevented** without revealing who voted
- **All proofs are verifiable on-chain** via Solidity smart contracts

### System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         Voting System                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐      │
│  │   Ballot     │    │   Merkle     │    │  Nullifier   │      │
│  │   Circuit    │    │   Circuit    │    │   Circuit    │      │
│  └──────────────┘    └──────────────┘    └──────────────┘      │
│        │                    │                    │               │
│        ▼                    ▼                    ▼               │
│  Encrypted Vote      Voter Registry      Double-Vote            │
│  Validation          Membership          Prevention             │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘
         │                    │                    │
         └────────────────────┴────────────────────┘
                              │
                              ▼
                    ┌──────────────────┐
                    │  Proof Server    │
                    │  (Optional)      │
                    └──────────────────┘
                              │
                              ▼
                    ┌──────────────────┐
                    │  Smart Contract  │
                    │  Verifiers       │
                    └──────────────────┘
```

## 📦 Repository Structure

```
.
├── README.md                          # Detailed circuit documentation
├── PROJECT_README.md                  # This file (project overview)
├── Makefile                           # Build automation
├── package.json                       # Node.js dependencies
│
├── ballot/                            # Encrypted voting circuit
│   ├── ballot.circom                  # ElGamal encryption proof
│   ├── build/                         # Compiled artifacts
│   │   ├── ballot.wasm                # Witness generator
│   │   ├── ballot_final.zkey          # Proving key
│   │   ├── verification_key.json      # Verification key
│   │   └── ballot.sol                 # Solidity verifier
│   └── circom-pairing/
│       └── demo/server/               # Proof generation server
│           ├── README.md              # Server documentation
│           ├── index.js               # Express API
│           └── exec.sh                # Proof generation script
│
├── merkle/                            # Voter registry circuit
│   ├── merkle.circom                  # Merkle tree membership proof
│   └── build/                         # Compiled artifacts
│       ├── merkle.wasm
│       ├── merkle_final.zkey
│       ├── verification_key.json
│       └── membership.sol
│
├── nullifier/                         # Double-voting prevention
│   ├── nullifier.circom               # Nullifier generation proof
│   └── build/                         # Compiled artifacts
│       ├── nullifier.wasm
│       ├── nullifier_final.zkey
│       ├── verification_key.json
│       └── nullifier.sol
│
└── tau/                               # Powers of Tau ceremony files
    ├── pot15_0000.ptau
    ├── pot15_0001.ptau
    └── pot15_final.ptau
```

## 🚀 Quick Start

### Prerequisites

```bash
# Install Circom compiler
cargo install --git https://github.com/iden3/circom.git circom

# Install Node.js dependencies
npm install
```

### Build All Circuits

```bash
# Generate trusted setup + compile all circuits
make all
```

This creates:
- `.wasm` files for witness generation
- `.zkey` files for proof generation
- `.sol` files for on-chain verification
- `verification_key.json` for off-chain verification

**Build time:** ~2-5 minutes

## 🔐 Circuit Details

### 1. Ballot Circuit - Encrypted Vote Validation

**Purpose:** Prove that an encrypted vote is valid (0 or 1) without revealing the vote.

**Cryptography:** ElGamal encryption on BabyJubJub curve

**Public Inputs:**
- `C1x, C1y` - Ephemeral public key
- `C2x, C2y` - Encrypted vote

**Private Inputs:**
- `M` - Vote bit (0 or 1)
- `K` - Encryption randomness

**Constraints:** 1,234 (example)

**Use Case:**
```javascript
// Voter encrypts their vote
const vote = 1; // or 0
const randomness = generateRandom();
const ciphertext = elgamalEncrypt(vote, publicKey, randomness);

// Generate proof that ciphertext encrypts a valid vote
const proof = await generateBallotProof({
  C1x: ciphertext.C1.x,
  C1y: ciphertext.C1.y,
  C2x: ciphertext.C2.x,
  C2y: ciphertext.C2.y,
  M: vote,
  K: randomness
});

// Submit ciphertext + proof (vote remains hidden)
await submitVote(ciphertext, proof);
```

---

### 2. Merkle Circuit - Voter Registry Membership

**Purpose:** Prove you're a registered voter without revealing which voter you are.

**Cryptography:** MiMC hash function, Merkle tree (depth 8)

**Public Inputs:**
- `MerkleRoot` - Root of voter registry tree
- `LeafValue` - Your public commitment

**Private Inputs:**
- `SecretKey` - Your secret key
- `MerklePath[8]` - Sibling hashes
- `PathPositions[8]` - Path directions

**Constraints:** 2,456 (example)

**Use Case:**
```javascript
// Voter proves they're in the registry
const secretKey = loadSecretKey();
const commitment = mimcHash(secretKey);
const { path, positions } = getMerklePath(commitment);

const proof = await generateMerkleProof({
  MerkleRoot: registryRoot,
  LeafValue: commitment,
  SecretKey: secretKey,
  MerklePath: path,
  PathPositions: positions
});

// Proof shows membership without revealing identity
await verifyVoterEligibility(proof);
```

---

### 3. Nullifier Circuit - Double-Voting Prevention

**Purpose:** Generate a unique nullifier to prevent voting twice while maintaining anonymity.

**Cryptography:** MiMC hash function

**Public Inputs:**
- `Commitment` - Your public commitment
- `EventID` - Unique voting event identifier
- `Nullifier` - Unique nullifier for this vote

**Private Inputs:**
- `SecretKey` - Your secret key

**Constraints:** 789 (example)

**Use Case:**
```javascript
// Generate nullifier for this election
const secretKey = loadSecretKey();
const commitment = mimcHash(secretKey);
const eventID = "election-2024-president";
const nullifier = mimcHash(secretKey, eventID);

const proof = await generateNullifierProof({
  Commitment: commitment,
  EventID: eventID,
  Nullifier: nullifier,
  SecretKey: secretKey
});

// Smart contract checks if nullifier was used before
await castVote(nullifier, proof); // Reverts if nullifier exists
```

## 🌐 Deployment Options

### Option 1: Client-Side Proof Generation (Recommended)

**Pros:** Maximum privacy, no server trust required  
**Cons:** Slower proof generation (~5-30s in browser)

```html
<script src="https://cdn.jsdelivr.net/npm/snarkjs@0.7.6/build/snarkjs.min.js"></script>
<script>
  const { proof, publicSignals } = await snarkjs.groth16.fullProve(
    input,
    '/ballot/build/ballot_js/ballot.wasm',
    '/ballot/build/ballot_final.zkey'
  );
</script>
```

### Option 2: Server-Side Proof Generation

**Pros:** Faster proofs (1-5s with RapidSNARK)  
**Cons:** Server sees private inputs (requires trust)

See [`ballot/circom-pairing/demo/server/README.md`](ballot/circom-pairing/demo/server/README.md) for server setup.

```bash
cd ballot/circom-pairing/demo/server
npm install
npm start
```

**API Usage:**
```javascript
// Submit proof request
const { id } = await fetch('http://localhost:3000/generate_proof', {
  method: 'POST',
  body: JSON.stringify(input)
}).then(r => r.json());

// Poll for result
const proof = await pollResult(id);
```

### Option 3: Hybrid (Client Witness + Server Proof)

**Pros:** Fast + private (server never sees secrets)  
**Cons:** More complex implementation

```javascript
// 1. Generate witness client-side
const witness = await generateWitness(input);

// 2. Send witness to server for proof generation
const proof = await serverProve(witness);
```

## 🔗 Smart Contract Integration

### Deploy Verifiers

```solidity
// Deploy generated verifier contracts
BallotVerifier ballotVerifier = new BallotVerifier();
MembershipVerifier membershipVerifier = new MembershipVerifier();
NullifierVerifier nullifierVerifier = new NullifierVerifier();
```

### Voting Contract Example

```solidity
contract PrivateVoting {
    BallotVerifier public ballotVerifier;
    MembershipVerifier public membershipVerifier;
    NullifierVerifier public nullifierVerifier;
    
    bytes32 public voterRegistryRoot;
    mapping(bytes32 => bool) public usedNullifiers;
    
    struct EncryptedVote {
        uint256 C1x;
        uint256 C1y;
        uint256 C2x;
        uint256 C2y;
    }
    
    EncryptedVote[] public encryptedVotes;
    
    function castVote(
        EncryptedVote memory vote,
        uint256[2] memory ballotProofA,
        uint256[2][2] memory ballotProofB,
        uint256[2] memory ballotProofC,
        uint256[8] memory ballotPublicSignals,
        
        uint256[2] memory membershipProofA,
        uint256[2][2] memory membershipProofB,
        uint256[2] memory membershipProofC,
        uint256[8] memory membershipPublicSignals,
        
        uint256[2] memory nullifierProofA,
        uint256[2][2] memory nullifierProofB,
        uint256[2] memory nullifierProofC,
        uint256[8] memory nullifierPublicSignals
    ) external {
        // 1. Verify voter is registered
        require(
            membershipVerifier.verifyProof(
                membershipProofA,
                membershipProofB,
                membershipProofC,
                membershipPublicSignals
            ),
            "Invalid membership proof"
        );
        
        // 2. Verify encrypted vote is valid (0 or 1)
        require(
            ballotVerifier.verifyProof(
                ballotProofA,
                ballotProofB,
                ballotProofC,
                ballotPublicSignals
            ),
            "Invalid ballot proof"
        );
        
        // 3. Verify nullifier and prevent double-voting
        bytes32 nullifier = bytes32(nullifierPublicSignals[2]);
        require(!usedNullifiers[nullifier], "Already voted");
        
        require(
            nullifierVerifier.verifyProof(
                nullifierProofA,
                nullifierProofB,
                nullifierProofC,
                nullifierPublicSignals
            ),
            "Invalid nullifier proof"
        );
        
        // 4. Record vote
        usedNullifiers[nullifier] = true;
        encryptedVotes.push(vote);
        
        emit VoteCast(nullifier, vote);
    }
    
    // Homomorphic tally (sum encrypted votes)
    function tallyVotes() external view returns (EncryptedVote memory) {
        EncryptedVote memory sum = encryptedVotes[0];
        for (uint i = 1; i < encryptedVotes.length; i++) {
            sum = addEncryptedVotes(sum, encryptedVotes[i]);
        }
        return sum;
    }
}
```

## 🧪 Testing

### Test Individual Circuits

```bash
# Generate test inputs
echo '{"C1x":"123","C1y":"456","C2x":"789","C2y":"012","M":"1","K":"345"}' > input.json

# Generate witness
node ballot/build/ballot_js/generate_witness.js \
  ballot/build/ballot_js/ballot.wasm \
  input.json \
  witness.wtns

# Generate proof
snarkjs groth16 prove \
  ballot/build/ballot_final.zkey \
  witness.wtns \
  proof.json \
  public.json

# Verify proof
snarkjs groth16 verify \
  ballot/build/verification_key.json \
  public.json \
  proof.json
```

### Integration Testing

```javascript
const { expect } = require('chai');

describe('Voting System', () => {
  it('should accept valid vote from registered voter', async () => {
    const secretKey = generateSecretKey();
    const commitment = mimcHash(secretKey);
    
    // Register voter
    await voterRegistry.addVoter(commitment);
    
    // Cast vote
    const vote = 1;
    const { ciphertext, ballotProof } = await encryptVote(vote);
    const membershipProof = await proveMembership(secretKey);
    const nullifierProof = await generateNullifier(secretKey, eventID);
    
    await votingContract.castVote(
      ciphertext,
      ballotProof,
      membershipProof,
      nullifierProof
    );
    
    expect(await votingContract.voteCount()).to.equal(1);
  });
  
  it('should reject double-voting', async () => {
    // ... cast first vote ...
    
    await expect(
      votingContract.castVote(/* same nullifier */)
    ).to.be.revertedWith('Already voted');
  });
});
```

## 🔒 Security Considerations

### Trusted Setup

⚠️ **The included Powers of Tau ceremony is for TESTING ONLY.**

For production:
1. Use [Perpetual Powers of Tau](https://github.com/privacy-scaling-explorations/perpetualpowersoftau)
2. Or run a multi-party ceremony with your community
3. Verify ceremony transcripts

### Key Management

- **Never reuse** secret keys across different voting events
- Store secret keys in secure enclaves (hardware wallets, TEEs)
- Use key derivation (BIP-39) for backup/recovery

## 🛠️ Development

### Optimizing Circuits

```bash
# Analyze constraint count
circom mycircuit.circom --r1cs --sym --inspect

# Profile witness generation
time node build/mycircuit_js/generate_witness.js ...

# Benchmark proof generation
time snarkjs groth16 prove ...
```

## 🤝 Contributing

Contributions welcome! Please:
- Follow existing code style
- Add tests for new features
- Update documentation
- Run `make all` before submitting PR

## 🙋 FAQ

**Q: Can votes be decrypted?**  
A: Only by the election authority holding the private key. Use threshold encryption for decentralized tallying.

**Q: How many voters can the system support?**  
A: Merkle tree depth 8 supports 256 voters. Increase depth for more (depth 20 = 1M voters).

**Q: What prevents the server from generating fake proofs?**  
A: The server cannot generate valid proofs without the voter's secret key. Use client-side proving for maximum security.

**Q: Can I use this in production?**  
A: After proper trusted setup, security audit, and key management implementation, yes.

**Q: How do I handle voter registration?**  
A: Maintain a Merkle tree of voter commitments. Update root on-chain when adding voters.
