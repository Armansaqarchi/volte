// Authentication utilities
import { buildMimc7 } from "circomlibjs";


export interface User {
  commitment: string
  username: string
}

export function getCurrentUser(): User | null {
  if (typeof window === "undefined") return null

  const userStr = localStorage.getItem("currentUser")
  if (!userStr) return null
  return userStr ? JSON.parse(userStr) : null
}

export function logout() {
  localStorage.removeItem("currentUser")
}

export function isAuthenticated(): boolean {
  return getCurrentUser() !== null
}

export function generateRandomBn254(): string {
  // random 32 bytes (browser)
  const bytes = new Uint8Array(32);
  crypto.getRandomValues(bytes);

  const BN254_FR = BigInt("21888242871839275222246405745257275088548364400416034343698204186575808495617");

  let x = BigInt(0);
  for (const b of bytes) x = (x << BigInt(8)) + BigInt(b);

  // make it a valid BN254 Fr private key: 1 <= x < BN254_FR
  // (rejection sampling to avoid modulo bias)
  if (x === BigInt(0) || x >= BN254_FR) return generateRandomBn254();

  return x.toString();
}

let _mimc7Promise: ReturnType<typeof buildMimc7> | null = null;

async function getMimc7() {
  if (!_mimc7Promise) _mimc7Promise = buildMimc7();
  return _mimc7Promise;
}

export async function mimc7Hash(inputs: Array<bigint>, key: bigint = 0n): Promise<bigint> {
  const mimc7 = await getMimc7();

  // Normalize into the field (helps avoid negative / oversized inputs)
  const F = mimc7.F;
  const inF = inputs.map((x) => F.e(x));

  // circomlibjs exposes multiHash([...], key)
  // (commonly used with key = 0)
  const out = mimc7.multiHash(inF, F.e(key));

  // multiHash returns an F element; convert to bigint
  return F.toObject(out) as bigint;
}

/** Convenience: return 0x-prefixed hex (32 bytes-ish, not fixed-width) */
export async function MimC7Hash(inputs: Array<bigint>): Promise<string> {
  const h = await mimc7Hash(inputs, 0n);
  return h.toString(10);
}

export function getCommitment() {
  const currentUserData = localStorage.getItem("currentUser")
  if (!currentUserData) {
    return null
  }

  return JSON.parse(currentUserData).commitment
}
