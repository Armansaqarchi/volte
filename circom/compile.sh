circom ballot/ballot.circom -l node_modules --r1cs --wasm --sym -o ballot/build
snarkjs groth16 setup ballot/build/ballot.r1cs tau/pot15_final.ptau ballot/build/ballot_0000.zkey
snarkjs zkey contribute ballot/build/ballot_0000.zkey ballot/build/ballot_final.zkey --name="1st contributor" -v
snarkjs zkey export verificationkey ballot/build/ballot_final.zkey ballot/build/verification_key.json
snarkjs zkey export solidityverifier ballot/build/ballot_final.zkey ballot/build/ballot.sol

circom merkle/merkle.circom -l node_modules --r1cs --wasm --sym -o merkle/build
snarkjs groth16 setup merkle/build/merkle.r1cs tau/pot15_final.ptau merkle/build/merkle_0000.zkey
snarkjs zkey contribute merkle/build/merkle_0000.zkey merkle/build/merkle_final.zkey --name="1st contributor" -v
snarkjs zkey export verificationkey merkle/build/merkle_final.zkey merkle/build/verification_key.json
snarkjs zkey export solidityverifier merkle/build/merkle_final.zkey merkle/build/membership.sol

circom nullifier/nullifier.circom -l node_modules --r1cs --wasm --sym -o nullifier/build
snarkjs groth16 setup nullifier/build/nullifier.r1cs tau/pot15_final.ptau nullifier/build/nullifier_0000.zkey
snarkjs zkey contribute nullifier/build/nullifier_0000.zkey nullifier/build/nullifier_final.zkey --name="1st contributor" -v
snarkjs zkey export verificationkey nullifier/build/nullifier_final.zkey nullifier/build/verification_key.json
snarkjs zkey export solidityverifier nullifier/build/nullifier_final.zkey nullifier/build/nullifier.sol
