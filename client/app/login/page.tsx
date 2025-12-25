"use client"

import type React from "react"

import { useState } from "react"
import Link from "next/link"
import { useRouter } from 'next/navigation'
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useToast } from "@/hooks/use-toast"
import { Shield, AlertTriangle } from 'lucide-react'
import {toCommitment, generatePrivateKey} from "@/lib/auth"

export default function LoginPage() {
  const router = useRouter()
  const { toast } = useToast()
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [rawPrivateKey, setRawPrivateKey] = useState("")
  const [isLoading, setIsLoading] = useState(false)

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsLoading(true)

    try {
      const hashedPrivateKey = await toCommitment(rawPrivateKey)
      console.log("The commitment : ", hashedPrivateKey)

      console.log(JSON.stringify({
        username: username,
        password: password,
        commitment: hashedPrivateKey,
      }))

      const response = await fetch("http://localhost:8000/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: username,
          password: password,
          commitment: hashedPrivateKey,
        }),
        credentials: "include",
      })

      if (!response.ok) {
        throw new Error("Invalid credentials")
      }

      const data = await response.json()

      localStorage.setItem("currentUser", JSON.stringify(data))

      toast({
        title: "Welcome back!",
        description: "You have successfully logged in.",
      })

      router.push("/dashboard")
    } catch (err) {
      toast({
        title: "Login failed",
        description: "Invalid username, password, or private key.",
        variant: "destructive",
      })
    } finally {
      setIsLoading(false)
    }
  }

  return (
      <div className="min-h-screen bg-background flex items-center justify-center p-4">
        <div className="w-full max-w-md space-y-6">
          <Link href="/" className="flex items-center justify-center gap-2 mb-8">
            <Shield className="h-8 w-8 text-primary" />
            <span className="font-semibold text-2xl">ZK Vote</span>
          </Link>

          <Card>
            <CardHeader className="space-y-1">
              <CardTitle className="text-2xl">Log in</CardTitle>
              <CardDescription>Enter your credentials to access your account</CardDescription>
            </CardHeader>
            <CardContent>
              <form onSubmit={handleLogin} className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="username">Username</Label>
                  <Input
                      id="username"
                      type="username"
                      placeholder=""
                      value={username}
                      onChange={(e) => setUsername(e.target.value)}
                      required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="password">Password</Label>
                  <Input
                      id="password"
                      type="password"
                      placeholder="••••••••"
                      value={password}
                      onChange={(e) => setPassword(e.target.value)}
                      required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="privateKey">Private Key</Label>
                  <Input
                      id="privateKey"
                      type="password"
                      placeholder="Enter your private key"
                      value={rawPrivateKey}
                      onChange={(e) => setRawPrivateKey(e.target.value)}
                      required
                  />

                  <p className="text-xs text-muted-foreground">
                    Enter the raw private key you saved during signup.
                  </p>

                  {/* Attention / security note */}
                  <div className="flex items-start gap-2 rounded-md bg-yellow-100 dark:bg-yellow-900/40 p-2">
                    <AlertTriangle className="h-4 w-4 text-yellow-700 dark:text-yellow-400 mt-0.5"/>
                    <p className="text-xs text-yellow-700 dark:text-yellow-400">
                      This is only used for verification. Your private key is hashed locally and never exposed to the
                      internet.
                    </p>
                  </div>
                </div>
                <Button type="submit" className="w-full" disabled={isLoading}>
                  {isLoading ? "Logging in..." : "Log in"}
                </Button>
              </form>

              <div className="mt-4 text-center text-sm">
                <span className="text-muted-foreground">Don't have an account? </span>
                <Link href="/signup" className="text-primary hover:underline">
                  Sign up
                </Link>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
  )
}
