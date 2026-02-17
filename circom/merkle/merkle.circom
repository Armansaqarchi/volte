pragma circom 2.1.0;

include "circomlib/circuits/mimc.circom";

/*
Public:
  MerkleRoot
  LeafValue

Secret:
  MerklePath[depth]
  PathPositions[depth]  (0 => current is left, 1 => current is right)
  SecretKey

Constraints:
  LeafValue == MiMC(SecretKey) with k=0
  current_0 = LeafValue
  For each i:
    right_i = isRight_i ? current_i : sibling_i
    left_i  = isRight_i ? sibling_i : current_i
    current_{i+1} = MiMC(left_i, right_i) with k=0
  current_depth == MerkleRoot
*/
template MerkleCircuit(depth) {
    // --- public inputs ---
    signal input MerkleRoot;
    signal input LeafValue;

    // --- secret inputs ---
    signal input MerklePath[depth];
    signal input PathPositions[depth];
    signal input SecretKey;

    // circomlib MiMC7 rounds (BN254 standard)
    var N_ROUNDS = 91;

    // 1) Leaf check: LeafValue == MiMC(SecretKey) with k=0
    component leafHash = MultiMiMC7(1, N_ROUNDS);
    leafHash.k <== 0;
    leafHash.in[0] <== SecretKey;
    leafHash.out === LeafValue;

    // 2) Merkle folding
    signal currents[depth + 1];
    signal lefts[depth];
    signal rights[depth];

    component nodeHash[depth];

    currents[0] <== LeafValue;

    for (var i = 0; i < depth; i++) {
        // enforce PathPositions[i] is a bit
        PathPositions[i] * (PathPositions[i] - 1) === 0;

        // right = Select(isRight, current, sibling)
        // right = sibling + isRight*(current - sibling)
        rights[i] <== MerklePath[i] + PathPositions[i] * (currents[i] - MerklePath[i]);

        // left = Select(isRight, sibling, current)
        // left = current + isRight*(sibling - current)
        lefts[i] <== currents[i] + PathPositions[i] * (MerklePath[i] - currents[i]);

        nodeHash[i] = MultiMiMC7(2, N_ROUNDS);
        nodeHash[i].k <== 0;
        nodeHash[i].in[0] <== lefts[i];
        nodeHash[i].in[1] <== rights[i];

        currents[i + 1] <== nodeHash[i].out;
    }

    // 3) Root equality
    currents[depth] === MerkleRoot;
}

// Set depth to match your witness lengths
component main { public [MerkleRoot, LeafValue] } = MerkleCircuit(8);
