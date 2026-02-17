pragma circom 2.1.0;

include "circomlib/circuits/mimc.circom";

/*
Public:
  Commitment
  EventID
  Nullifier

Secret:
  SecretKey

Constraints:
  Commitment == MiMC(SecretKey) with k=0
  Nullifier  == MiMC(SecretKey, EventID) with k=0
*/
template NullifierCircuit() {
    // --- public inputs ---
    signal input Commitment;
    signal input EventID;
    signal input Nullifier;

    // --- secret input ---
    signal input SecretKey;

    // circomlib MiMC7 rounds (BN254 standard)
    var N_ROUNDS = 91;

    // Commitment = MiMC(SecretKey) with k=0
    component commitHash = MultiMiMC7(1, N_ROUNDS);
    commitHash.k <== 0;
    commitHash.in[0] <== SecretKey;
    commitHash.out === Commitment;

    // Nullifier = MiMC(SecretKey, EventID) with k=0
    component nullHash = MultiMiMC7(2, N_ROUNDS);
    nullHash.k <== 0;
    nullHash.in[0] <== SecretKey;
    nullHash.in[1] <== EventID;
    nullHash.out === Nullifier;
}

component main { public [Commitment, EventID, Nullifier] } = NullifierCircuit();
