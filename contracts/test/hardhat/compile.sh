mkdir contracts
cp -r ./../../groth16 ./contracts
cp -r ./../../volte.sol ./contracts
npx hardhat clean && npx hardhat compile