solc --abi --bin ./contracts/volte.sol -o ./contracts/build/volte --overwrite
solc --abi --bin ./contracts/groth16/ballot.sol -o ./contracts/build/groth16/ballot --overwrite
solc --abi --bin ./contracts/groth16/membership.sol -o ./contracts/build/groth16/membership --overwrite
solc --abi --bin ./contracts/groth16/nullifier.sol -o ./contracts/build/groth16/nullifier --overwrite
abigen --bin=./contracts/build/VolteContract.bin --abi=./contracts/build/VolteContract.abi --pkg=contracts --out=./server/chain/contracts/contract.go --type=volte