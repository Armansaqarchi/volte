pragma circom 2.1.0;

include "circomlib/circuits/bitify.circom";
include "circomlib/circuits/babyjub.circom";
include "circomlib/circuits/escalarmulfix.circom";

// Same functionality as your gnark circuit, but over BabyJubJub:
//
// Public:  C1, C2
// Private: M, K
//
// Constraints:
// 1) C1 == K*G
// 2) C2 == K*Y + M*G
// 3) M is boolean

template BallotCircuit(
    // Compile-time constants (just like your gnark flags/meta)
    Gx, Gy,
    Yx, Yy
) {
    // Public inputs (ciphertext points)
    signal input C1x;
    signal input C1y;
    signal input C2x;
    signal input C2y;

    // Private witness
    signal input M; // vote bit
    signal input K; // randomness scalar

    // (Optional but recommended) sanity-check that ciphertext points are on-curve
    component c1Check = BabyCheck();
    c1Check.x <== C1x;
    c1Check.y <== C1y;

    component c2Check = BabyCheck();
    c2Check.x <== C2x;
    c2Check.y <== C2y;

    // 3) Enforce M is boolean: M*(M-1)=0
    M * (M - 1) === 0;

    // Convert K to bits once and reuse
    // 253 is the typical scalar bitlength used with BabyJubJub tooling.
    component kBits = Num2Bits(253);
    kBits.in <== K;

    // 1) C1 == K*G  (fixed-base mul)
    component kG = EscalarMulFix(253, [Gx, Gy]);
    for (var i = 0; i < 253; i++) {
        kG.e[i] <== kBits.out[i];
    }

    kG.out[0] === C1x;
    kG.out[1] === C1y;

    // 2) Ky == K*Y  (fixed-base mul)
    component kY = EscalarMulFix(253, [Yx, Yy]);
    for (var j = 0; j < 253; j++) {
        kY.e[j] <== kBits.out[j];
    }

    // KyPlusG = Ky + G
    component kyPlusG = BabyAdd();
    kyPlusG.x1 <== kY.out[0];
    kyPlusG.y1 <== kY.out[1];
    kyPlusG.x2 <== Gx;
    kyPlusG.y2 <== Gy;

    // Select expected C2 based on M, without branching:
    // exp = Ky + M*(KyPlusG - Ky)
    signal expX;
    signal expY;

    expX <== kY.out[0] + M * (kyPlusG.xout - kY.out[0]);
    expY <== kY.out[1] + M * (kyPlusG.yout - kY.out[1]);
``
    expX === C2x;
    expY === C2y;
}

component main { public [C1x, C1y, C2x, C2y] } = BallotCircuit(
    // You set these at compile time (same role as your gnark flags)
    // Example placeholders:
    5299619240641551281634865583518297030282874472190772894086521144482721001553,
    16950150798460657717958625567821834550301663161624707787222815936182638968203,
    18524760469487540272086982072248352918977679699605098074565248706868593560314,
    21825033186726430338019907128549959920138456209872775056787107207311558509214
);
