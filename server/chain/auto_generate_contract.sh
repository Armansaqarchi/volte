solc --abi --bin ./contracts/volte.sol -o ./contracts/build/volte --overwrite
solc --abi --bin ./contracts/groth16/ballot.sol -o ./contracts/build/groth16/ballot --overwrite
solc --abi --bin ./contracts/groth16/membership.sol -o ./contracts/build/groth16/membership --overwrite
solc --abi --bin ./contracts/groth16/nullifier.sol -o ./contracts/build/groth16/nullifier --overwrite
abigen --bin=./contracts/build/groth16/ballot/BallotVerifier.bin --abi=./contracts/build/groth16/ballot/BallotVerifier.abi --pkg=ballot --out=./server/chain/contracts/ballot/ballot.go --type=volte
abigen --bin=./contracts/build/groth16/membership/MembershipVerifier.bin --abi=./contracts/build/groth16/membership/MembershipVerifier.abi --pkg=membership --out=./server/chain/contracts/membership/membership.go --type=volte
abigen --bin=./contracts/build/groth16/nullifier/NullifierVerifier.bin --abi=./contracts/build/groth16/nullifier/NullifierVerifier.abi --pkg=nullifier --out=./server/chain/contracts/nullifier/nullifier.go --type=volte
abigen --bin=./contracts/build/volte/VolteContract.bin --abi=./contracts/build/volte/VolteContract.abi --pkg=contracts --out=./server/chain/contracts/contract.go --type=volte