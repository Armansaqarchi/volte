# Simplified Makefile for Circom ZK-SNARK Circuit Compilation

.PHONY: all tau build clean help

# Default target
all: tau build

# Directories
TAU_DIR = tau
NODE_MODULES = node_modules

# Tau ceremony parameters
TAU_POWER = 15
TAU_CURVE = bn128

# Tau files
TAU_INITIAL = $(TAU_DIR)/pot$(TAU_POWER)_0000.ptau
TAU_CONTRIB = $(TAU_DIR)/pot$(TAU_POWER)_0001.ptau
TAU_FINAL = $(TAU_DIR)/pot$(TAU_POWER)_final.ptau

CIRCOM := $(HOME)/.cargo/bin/circom


help:
	@echo "Circom ZK-SNARK Build System"
	@echo "============================="
	@echo ""
	@echo "Available commands:"
	@echo "  make tau    - Generate powers of tau ceremony files"
	@echo "  make build  - Build all circuits (compile + setup + export)"
	@echo "  make all    - Do both tau and build"
	@echo "  make clean  - Remove all build artifacts"

# ============================================================================
# Powers of Tau Ceremony
# ============================================================================


tau: | $(TAU_DIR)
	@echo "🔧 Generating powers of tau ceremony..."
	@mkdir -p $(TAU_DIR)
	@snarkjs powersoftau new $(TAU_CURVE) $(TAU_POWER) $(TAU_INITIAL) -v
	@snarkjs powersoftau contribute $(TAU_INITIAL) $(TAU_CONTRIB) --name="first contribution" -v
	@snarkjs powersoftau prepare phase2 $(TAU_CONTRIB) $(TAU_FINAL) -v
	@snarkjs powersoftau verify $(TAU_FINAL)
	@echo "✅ Powers of tau ceremony complete!"

# ============================================================================
# Build All Circuits
# ============================================================================

build: $(TAU_FINAL)
	@echo "🔧 Building all circuits..."
	@echo ""
	@echo "📦 Compiling ballot circuit..."
	@mkdir -p ballot/build
	@$(CIRCOM) ballot/ballot.circom -l $(NODE_MODULES) --r1cs --wasm --sym -o ballot/build
	@snarkjs groth16 setup ballot/build/ballot.r1cs $(TAU_FINAL) ballot/build/ballot_0000.zkey
	@snarkjs zkey contribute ballot/build/ballot_0000.zkey ballot/build/ballot_final.zkey --name="1st contributor" -v
	@snarkjs zkey export verificationkey ballot/build/ballot_final.zkey ballot/build/verification_key.json
	@snarkjs zkey export solidityverifier ballot/build/ballot_final.zkey ballot/build/ballot.sol
	@echo "✅ Ballot circuit complete!"
	@echo ""
	@echo "📦 Compiling merkle circuit..."
	@mkdir -p merkle/build
	@$(CIRCOM) merkle/merkle.circom -l $(NODE_MODULES) --r1cs --wasm --sym -o merkle/build
	@snarkjs groth16 setup merkle/build/merkle.r1cs $(TAU_FINAL) merkle/build/merkle_0000.zkey
	@snarkjs zkey contribute merkle/build/merkle_0000.zkey merkle/build/merkle_final.zkey --name="1st contributor" -v
	@snarkjs zkey export verificationkey merkle/build/merkle_final.zkey merkle/build/verification_key.json
	@snarkjs zkey export solidityverifier merkle/build/merkle_final.zkey merkle/build/membership.sol
	@echo "✅ Merkle circuit complete!"
	@echo ""
	@echo "📦 Compiling nullifier circuit..."
	@mkdir -p nullifier/build
	@$(CIRCOM) nullifier/nullifier.circom -l $(NODE_MODULES) --r1cs --wasm --sym -o nullifier/build
	@snarkjs groth16 setup nullifier/build/nullifier.r1cs $(TAU_FINAL) nullifier/build/nullifier_0000.zkey
	@snarkjs zkey contribute nullifier/build/nullifier_0000.zkey nullifier/build/nullifier_final.zkey --name="1st contributor" -v
	@snarkjs zkey export verificationkey nullifier/build/nullifier_final.zkey nullifier/build/verification_key.json
	@snarkjs zkey export solidityverifier nullifier/build/nullifier_final.zkey nullifier/build/nullifier.sol
	@echo "✅ Nullifier circuit complete!"
	@echo ""
	@echo "🎉 All circuits built successfully!"
	@echo ""
	@echo "Build folders created:"
	@echo "  - ballot/build/    (contains ballot.wasm, ballot_final.zkey, verification_key.json, ballot.sol)"
	@echo "  - merkle/build/    (contains merkle.wasm, merkle_final.zkey, verification_key.json, membership.sol)"
	@echo "  - nullifier/build/ (contains nullifier.wasm, nullifier_final.zkey, verification_key.json, nullifier.sol)"

# ============================================================================
# Clean
# ============================================================================

clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf ballot/build merkle/build nullifier/build $(TAU_DIR)
	@echo "✅ Clean complete!"
.PHONY: all setup-tau build clean help
all: setup-tau build
	@echo "  make setup-tau - Generate powers of tau ceremony files"
	@echo "  make all       - Do both setup-tau and build"
setup-tau: $(TAU_DIR)
