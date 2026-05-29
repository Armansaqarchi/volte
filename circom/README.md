# Circom ZK-SNARK Circuits for Privacy-Preserving Voting

This repository contains three Circom circuits implementing zero-knowledge proofs for privacy-preserving voting systems using the Groth16 proving system on the BN128 (BN254) elliptic curve.

## 🔐 Circuits Overview

### 1. Ballot Circuit (`ballot/ballot.circom`)
Implements ElGamal encryption proof for encrypted voting on the BabyJubJub curve.

**Public Inputs:**
- `C1x, C1y`: First ciphertext point (ephemeral public key)
- `C2x, C2y`: Second ciphertext point (encrypted vote)

**Private Inputs:**
- `M`: Vote bit (0 or 1)
- `K`: Random scalar (encryption randomness)

**Constraints:**
1. `C1 = K·G` (ephemeral key generation)
2. `C2 = K·Y + M·G` (vote encryption)
3. `M ∈ {0, 1}` (vote validity)

**Use Case:** Proves that an encrypted vote is valid (0 or 1) without revealing the vote itself.

---

### 2. Merkle Circuit (`merkle/merkle.circom`)
Implements Merkle tree membership proof using MiMC hash function.

**Public Inputs:**
- `MerkleRoot`: Root hash of the Merkle tree
- `LeafValue`: Commitment to the secret key

**Private Inputs:**
- `SecretKey`: User's secret key
- `MerklePath[8]`: Sibling hashes along the path (depth 8)
- `PathPositions[8]`: Binary path indicators (0=left, 1=right)

**Constraints:**
1. `LeafValue = MiMC(SecretKey)`
2. Merkle path verification from leaf to root
3. `ComputedRoot = MerkleRoot`

**Use Case:** Proves membership in a voter registry without revealing which voter you are.

---

### 3. Nullifier Circuit (`nullifier/nullifier.circom`)
Implements double-voting prevention using nullifiers.

**Public Inputs:**
- `Commitment`: User's public commitment
- `EventID`: Unique identifier for the voting event
- `Nullifier`: Unique nullifier for this vote

**Private Inputs:**
- `SecretKey`: User's secret key

**Constraints:**
1. `Commitment = MiMC(SecretKey)`
2. `Nullifier = MiMC(SecretKey, EventID)`

**Use Case:** Proves you have the right to vote (via commitment) and generates a unique nullifier to prevent double-voting, without revealing your identity.

---

## 🛠️ Prerequisites

### Required Tools
- **Circom compiler** (v2.1.0+): Install from [circom documentation](https://docs.circom.io/getting-started/installation/)
  ```bash
  cargo install --git https://github.com/iden3/circom.git circom
  ```
- **snarkjs** (v0.7.6+): Installed via npm
  ```bash
  npm install
  ```

### System Requirements
- Node.js 16+
- 8GB+ RAM (for larger circuits)
- ~500MB disk space for build artifacts

---

## 🚀 Quick Start

### 1. Install Dependencies
```bash
npm install
```

### 2. Build All Circuits
```bash
make all
```

This command will:
1. Generate Powers of Tau ceremony files (trusted setup)
2. Compile all three circuits
3. Generate proving keys (`.zkey` files)
4. Generate verification keys (`.json` files)
5. Export Solidity verifiers (`.sol` files)

**Build time:** ~2-5 minutes depending on your system.

---

## 📦 Build Artifacts

After running `make all`, each circuit directory will contain a `build/` folder:

```
ballot/build/
├── ballot.wasm              # WebAssembly witness generator
├── ballot_final.zkey        # Proving key (4.2 MB)
├── verification_key.json    # Verification key
├── ballot.sol               # Solidity verifier contract
├── ballot.r1cs              # R1CS constraint system
└── ballot_js/
    ├── ballot.wasm          # WASM file for witness calculation
    ├── witness_calculator.js # Witness calculator module
    └── generate_witness.js  # CLI witness generator

merkle/build/
├── merkle.wasm
├── merkle_final.zkey
├── verification_key.json
├── membership.sol
└── merkle_js/...

nullifier/build/
├── nullifier.wasm
├── nullifier_final.zkey
├── verification_key.json
├── nullifier.sol
└── nullifier_js/...
```

---

## 🌐 Client-Side Proof Generation

### Browser Usage

The `.wasm` and `_final.zkey` files are designed for client-side proof generation in web browsers.

#### Example: Generating a Ballot Proof

```html
<!DOCTYPE html>
<html>
<head>
    <script src="https://cdn.jsdelivr.net/npm/snarkjs@0.7.6/build/snarkjs.min.js"></script>
</head>
<body>
<script type="module">
    // 1. Load the WASM and zkey files
    const wasmPath = '/ballot/build/ballot_js/ballot.wasm';
    const zkeyPath = '/ballot/build/ballot_final.zkey';

    // 2. Prepare circuit inputs
    const input = {
        // Public inputs (ciphertext)
        C1x: "12345...",  // Ephemeral public key X
        C1y: "67890...",  // Ephemeral public key Y
        C2x: "11111...",  // Encrypted vote X
        C2y: "22222...",  // Encrypted vote Y
        
        // Private inputs (witness)
        M: "1",           // Vote: 0 or 1
        K: "98765..."     // Random scalar
    };

    // 3. Generate witness
    const { proof, publicSignals } = await snarkjs.groth16.fullProve(
        input,
        wasmPath,
        zkeyPath
    );

    console.log("Proof:", proof);
    console.log("Public Signals:", publicSignals);

    // 4. Verify proof locally (optional)
    const vKey = await fetch('/ballot/build/verification_key.json')
        .then(r => r.json());
    
    const verified = await snarkjs.groth16.verify(
        vKey,
        publicSignals,
        proof
    );
    
    console.log("Proof valid:", verified);

    // 5. Export proof for Solidity verification
    const calldata = await snarkjs.groth16.exportSolidityCallData(
        proof,
        publicSignals
    );
    
    console.log("Solidity calldata:", calldata);
</script>
</body>
</html>
```

#### Example: Merkle Membership Proof

```javascript
import * as snarkjs from 'snarkjs';

async function proveMembership(secretKey, merklePath, pathPositions, merkleRoot) {
    // Calculate leaf value (commitment)
    const leafValue = mimcHash(secretKey);  // Use your MiMC implementation
    
    const input = {
        // Public
        MerkleRoot: merkleRoot,
        LeafValue: leafValue,
        
        // Private
        SecretKey: secretKey,
        MerklePath: merklePath,      // Array of 8 sibling hashes
        PathPositions: pathPositions  // Array of 8 bits (0 or 1)
    };

    const { proof, publicSignals } = await snarkjs.groth16.fullProve(
        input,
        'merkle/build/merkle_js/merkle.wasm',
        'merkle/build/merkle_final.zkey'
    );

    return { proof, publicSignals };
}
```

#### Example: Nullifier Generation

```javascript
async function generateNullifier(secretKey, eventID) {
    // Calculate commitment and nullifier
    const commitment = mimcHash(secretKey);
    const nullifier = mimcHash(secretKey, eventID);
    
    const input = {
        // Public
        Commitment: commitment,
        EventID: eventID,
        Nullifier: nullifier,
        
        // Private
        SecretKey: secretKey
    };

    const { proof, publicSignals } = await snarkjs.groth16.fullProve(
        input,
        'nullifier/build/nullifier_js/nullifier.wasm',
        'nullifier/build/nullifier_final.zkey'
    );

    return { proof, publicSignals, nullifier };
}
```

---

### Node.js Usage

```javascript
const snarkjs = require('snarkjs');
const fs = require('fs');

async function generateProof(circuitName, input) {
    const wasmPath = `./${circuitName}/build/${circuitName}_js/${circuitName}.wasm`;
    const zkeyPath = `./${circuitName}/build/${circuitName}_final.zkey`;

    // Generate proof
    const { proof, publicSignals } = await snarkjs.groth16.fullProve(
        input,
        wasmPath,
        zkeyPath
    );

    // Verify proof
    const vKey = JSON.parse(
        fs.readFileSync(`./${circuitName}/build/verification_key.json`)
    );
    
    const verified = await snarkjs.groth16.verify(vKey, publicSignals, proof);
    
    return { proof, publicSignals, verified };
}

// Example usage
const ballotInput = {
    C1x: "123...",
    C1y: "456...",
    C2x: "789...",
    C2y: "012...",
    M: "1",
    K: "345..."
};

generateProof('ballot', ballotInput).then(result => {
    console.log('Proof generated:', result.verified);
});
```

---

## 🔗 On-Chain Verification

### Deploy Verifier Contracts

Each circuit generates a Solidity verifier contract:

```bash
# Deploy to Ethereum/Polygon/etc.
# ballot/build/ballot.sol
# merkle/build/membership.sol
# nullifier/build/nullifier.sol
```

### Verify Proof On-Chain

```solidity
// Example: Verifying a ballot proof
contract VotingSystem {
    Groth16Verifier public ballotVerifier;
    
    constructor(address _verifier) {
        ballotVerifier = Groth16Verifier(_verifier);
    }
    
    function submitVote(
        uint[2] memory a,
        uint[2][2] memory b,
        uint[2] memory c,
        uint[4] memory publicSignals  // [C1x, C1y, C2x, C2y]
    ) public {
        require(
            ballotVerifier.verifyProof(a, b, c, publicSignals),
            "Invalid proof"
        );
        
        // Store encrypted vote (C1, C2)
        // ...
    }
}
```

### Converting snarkjs Proof to Solidity Format

```javascript
// Client-side: Export proof for Solidity
const calldata = await snarkjs.groth16.exportSolidityCallData(proof, publicSignals);

// Parse calldata for contract call
const argv = calldata.split(',').map(x => x.trim());
const a = JSON.parse(argv[0]);
const b = JSON.parse(argv[1]);
const c = JSON.parse(argv[2]);
const publicInputs = JSON.parse(argv[3]);

// Call contract
await votingContract.submitVote(a, b, c, publicInputs);
```

---

## 🧪 Testing

### Generate Test Witness (Node.js)

```bash
# Create input file
echo '{
  "C1x": "123...",
  "C1y": "456...",
  "C2x": "789...",
  "C2y": "012...",
  "M": "1",
  "K": "345..."
}' > input.json

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

---

## 📋 Makefile Commands

```bash
make help       # Show available commands
make tau        # Generate Powers of Tau ceremony (one-time setup)
make build      # Compile all circuits and generate keys
make all        # Run tau + build
make clean      # Remove all build artifacts
```

---

## 🔧 Advanced Configuration

### Modify Circuit Parameters

**Ballot Circuit:** Edit compile-time constants in `ballot/ballot.circom`:
```circom
component main { public [C1x, C1y, C2x, C2y] } = BallotCircuit(
    Gx,  // Generator point X
    Gy,  // Generator point Y
    Yx,  // Public key X
    Yy   // Public key Y
);
```

**Merkle Circuit:** Change tree depth in `merkle/merkle.circom`:
```circom
component main { public [MerkleRoot, LeafValue] } = MerkleCircuit(8);  // depth=8
```

After modifying circuits, rebuild:
```bash
make clean
make all
```

### Increase Powers of Tau

For larger circuits, increase the tau power in `Makefile`:
```makefile
TAU_POWER = 20  # Supports up to 2^20 constraints
```

---

## 📊 Circuit Statistics

| Circuit    | Constraints | Wasm Size | Zkey Size | Proving Time* | Verification Time* |
|------------|-------------|-----------|-----------|---------------|-------------------|
| Ballot     | ~5,000      | 97 KB     | 4.2 MB    | ~1-2s         | ~50ms             |
| Merkle     | ~3,000      | 85 KB     | 3.8 MB    | ~0.8-1.5s     | ~40ms             |
| Nullifier  | ~500        | 45 KB     | 2.1 MB    | ~0.3-0.5s     | ~30ms             |

*Approximate times on modern hardware (browser/Node.js)

---

## 🔐 Security Considerations

### Trusted Setup
- The Powers of Tau ceremony in this repo is for **development/testing only**
- For production, use a multi-party computation (MPC) ceremony
- Consider using [Perpetual Powers of Tau](https://github.com/privacy-scaling-explorations/perpetualpowersoftau)

### Client-Side Security
- Never expose private inputs (M, K, SecretKey) in logs or network requests
- Validate all public inputs before proof generation
- Use secure random number generation for K and SecretKey
- Store zkey files on trusted CDN/IPFS

### Smart Contract Security
- Add access controls to verifier contracts
- Implement nullifier tracking to prevent double-voting
- Consider gas optimization for on-chain verification (~250k-350k gas per proof)

---

## 📚 Resources

- [Circom Documentation](https://docs.circom.io/)
- [snarkjs Documentation](https://github.com/iden3/snarkjs)
- [Groth16 Paper](https://eprint.iacr.org/2016/260.pdf)
- [circomlib Library](https://github.com/iden3/circomlib)
- [ZK-SNARK Explainer](https://z.cash/technology/zksnarks/)

---

## 🤝 Contributing

Contributions are welcome! Please ensure:
- Circuits compile without errors
- Tests pass for all circuits
- Documentation is updated for new features

---

## 📄 License

This project uses circuits and libraries with various licenses:
- Circom circuits: Custom (see individual files)
- circomlib: GPL-3.0
- snarkjs: GPL-3.0

---

## 🙋 FAQ

**Q: Can I use these circuits in production?**  
A: These circuits are functional but require a proper trusted setup ceremony for production use.

**Q: How do I reduce proof generation time?**  
A: Use Web Workers in browsers, optimize circuit constraints, or use faster curves (BLS12-381).

**Q: Can I verify proofs off-chain?**  
A: Yes! Use `snarkjs.groth16.verify()` in Node.js or browsers for free verification.

**Q: What's the difference between `.wasm` in `build/` vs `build/ballot_js/`?**  
A: They're the same file. The `ballot_js/` directory contains the WASM plus helper scripts for witness generation.

**Q: How do I integrate with React/Vue/Angular?**  
A: Import snarkjs as an ES module and load `.wasm`/`.zkey` files as static assets. See browser example above.

---

**Built with ❤️ using Circom and snarkjs**
