// Authentication utilities
import {runWasm} from "@/lib/zkproof/go/main";

export interface User {
  commitment: string
  username: string
}

export function getCurrentUser(): User | null {
  if (typeof window === "undefined") return null

  const userStr = localStorage.getItem("currentUser")
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

export const toCommitment = async (randomHex: string) => {
  const hash = await runWasm("getMIMCHash", [`--secret=${randomHex}`])
  return hash as string
}
