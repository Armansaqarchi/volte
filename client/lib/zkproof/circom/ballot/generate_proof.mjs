import wcBuilder from "./witness_calculator.js";
import * as snarkjs from "snarkjs";
import * as circomlibjs from "circomlibjs";

// ---- browser-safe randomness ----
function randomBytes(n) {
    const arr = new Uint8Array(n);
    crypto.getRandomValues(arr);
    return arr;
}

function randScalarBelow(n) {
    // uniform random scalar in [1, n-1]
    const nBig = BigInt(n.toString());
    const nBits = nBig.toString(2).length;
    const nBytes = Math.ceil(nBits / 8);

    while (true) {
        const rHex = [...randomBytes(nBytes)]
            .map((b) => b.toString(16).padStart(2, "0"))
            .join("");
        const r = BigInt("0x" + rHex);
        const x = r % nBig;
        if (x !== 0n) return x;
    }
}

async function fetchArrayBuffer(path) {
    const res = await fetch(path);
    if (!res.ok) throw new Error(`Failed to fetch ${path}: ${res.status} ${res.statusText}`);
    return await res.arrayBuffer();
}

async function makeCiphertext({ M, Gx, Gy, Yx, Yy }) {
    const babyjub = await circomlibjs.buildBabyjub();
    const F = babyjub.F;

    const m = BigInt(M);
    if (m !== 0n && m !== 1n) throw new Error("M must be 0 or 1");

    const G = [F.e(BigInt(Gx)), F.e(BigInt(Gy))];
    const Y = [F.e(BigInt(Yx)), F.e(BigInt(Yy))];

    const K = randScalarBelow(babyjub.subOrder);
    const C1 = babyjub.mulPointEscalar(G, K);
    const kY = babyjub.mulPointEscalar(Y, K);
    const C2 = (m === 0n) ? kY : babyjub.addPoint(kY, G);

    return {
        C1x: F.toObject(C1[0]).toString(),
        C1y: F.toObject(C1[1]).toString(),
        C2x: F.toObject(C2[0]).toString(),
        C2y: F.toObject(C2[1]).toString(),
        K: K.toString(), // private witness
        M: m.toString(), // private witness
    };
}

const Gx = "5299619240641551281634865583518297030282874472190772894086521144482721001553";
const Gy = "16950150798460657717958625567821834550301663161624707787222815936182638968203";
const Yx = "18524760469487540272086982072248352918977679699605098074565248706868593560314";
const Yy = "21825033186726430338019907128549959920138456209872775056787107207311558509214";

export async function generateBallotProof(M) {
    const input = await makeCiphertext({ M, Gx, Gy, Yx, Yy });
    const wasmBuffer = await fetchArrayBuffer("/zk/ballot/ballot.wasm");
    const zkeyBuffer = await fetchArrayBuffer("/zk/ballot/ballot_final.zkey");

    const witnessCalculator = await wcBuilder(wasmBuffer);
    const wtnsBuff = await witnessCalculator.calculateWTNSBin(input, 0);
    const { proof, publicSignals } = await snarkjs.groth16.prove(
        new Uint8Array(zkeyBuffer),
        wtnsBuff
    );

    return {
        proof: {
            Arx:  proof.pi_a[0],
            Ary:  proof.pi_a[1],
            Brx0: proof.pi_b[0][0],
            Brx1: proof.pi_b[0][1],
            Bry0: proof.pi_b[1][0],
            Bry1: proof.pi_b[1][1],
            Cx:   proof.pi_c[0],
            Cy:   proof.pi_c[1]
        },
        input: publicSignals
    }
}

export default generateBallotProof