snarkjs powersoftau new bn128 15 tau/pot15_0000.ptau -v
snarkjs powersoftau contribute tau/pot15_0000.ptau pot15_0001.ptau --name="first contribution" -v
snarkjs powersoftau prepare phase2 tau/pot15_0001.ptau pot15_final.ptau -v
snarkjs powersoftau verify tau/pot15_final.ptau
