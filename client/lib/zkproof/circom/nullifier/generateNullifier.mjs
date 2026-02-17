import * as snarkjs from "snarkjs";
import wcBuilder from "./witness_calculator.js";

async function fetchArrayBuffer(path) {
    const res = await fetch(path);
    if (!res.ok) throw new Error(`Failed to fetch ${path}: ${res.status} ${res.statusText}`);
    return await res.arrayBuffer();
}

async function generateNullifier(input){

    const wasmPath = "/zk/nullifier/nullifier.wasm";
    const zkeyPath = "/zk/nullifier/nullifier_final.zkey";

    const wasmBuffer = await fetchArrayBuffer(wasmPath);
    const zkeyBuffer = await fetchArrayBuffer(zkeyPath);

    const witnessCalculator = await wcBuilder(wasmBuffer);

// 0 = no sanity check (fastest)
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

export default generateNullifier