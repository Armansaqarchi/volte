"use client"

import type React from "react"

import { useState, useEffect } from "react"
import Link from "next/link"
import { useRouter } from 'next/navigation'
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Shield, Copy, Check } from 'lucide-react'
import { useToast } from "@/hooks/use-toast"
import {toCommitment, generateRandomBn254} from "@/lib/auth"

export default function SignupPage() {
  const router = useRouter()
  const { toast } = useToast()
  const [username, setUsername] = useState("") // added username state
  const [privateKey, setPrivateKey] = useState("")
  const [password, setPassword] = useState("")
  const [isLoading, setIsLoading] = useState(false)
  const [copied, setCopied] = useState(false)
  const [savedKey, setSavedKey] = useState(false)

  useEffect(() => {
    setPrivateKey(generateRandomBn254())
  }, [])

  const copyToClipboard = () => {
    navigator.clipboard.writeText(privateKey)
    setCopied(true)
    setTimeout(() => setCopied(false), 2000)
    toast({
      title: "Copied!",
      description: "Private key copied to clipboard.",
    })
  }

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!savedKey) {
      toast({
        title: "Please confirm",
        description: "You must confirm that you've saved your private key.",
        variant: "destructive",
      })
      return
    }

    setIsLoading(true)
    try {
      const response = await fetch("http://localhost:8000/auth/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: username,
          commitment: await toCommitment(privateKey),
          password: password,
        }),
        credentials: "include",
      })

      if (!response.ok) throw new Error("Failed to fetch response.")
      const data = await response.json()
      localStorage.setItem("currentUser", JSON.stringify(data))
      localStorage.setItem("privateKey", privateKey)
      toast({
        title: "Account created!",
        description: "Welcome to ZK Vote.",
      })
      router.push("/dashboard")
    } catch (err) {
      toast({
        title: "Signup failed",
        description: "An error occurred during signup.",
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
              <CardTitle className="text-2xl">Create an account</CardTitle>
              <CardDescription>Enter your information to get started</CardDescription>
            </CardHeader>
            <CardContent>
              <form onSubmit={handleSignup} className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="username">Username</Label>
                  <Input
                      id="username"
                      type="text"
                      placeholder="Choose your username"
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
                      minLength={6}
                  />
                </div>

                <div className="space-y-2">
                  <Label>Your Private Key</Label>
                  <div className="flex gap-2">
                    <div className="flex-1 bg-muted p-3 rounded-md font-mono text-sm break-all flex items-center">
                      {privateKey}
                    </div>
                    <Button
                        type="button"
                        variant="outline"
                        size="icon"
                        onClick={copyToClipboard}
                    >
                      {copied ? (
                          <Check className="h-4 w-4" />
                      ) : (
                          <Copy className="h-4 w-4" />
                      )}
                    </Button>
                  </div>
                  <p className="text-xs text-muted-foreground">
                    ⚠️ Save this private key somewhere safe. You will need it to access your account.
                  </p>
                </div>

                <div className="flex items-center gap-2">
                  <input
                      type="checkbox"
                      id="saved-key"
                      checked={savedKey}
                      onChange={(e) => setSavedKey(e.target.checked)}
                      className="h-4 w-4 rounded border border-input"
                  />
                  <Label htmlFor="saved-key" className="font-normal cursor-pointer">
                    I have saved my private key securely
                  </Label>
                </div>

                <Button type="submit" className="w-full" disabled={isLoading || !savedKey}>
                  {isLoading ? "Creating account..." : "Sign up"}
                </Button>
              </form>

              <div className="mt-4 text-center text-sm">
                <span className="text-muted-foreground">Already have an account? </span>
                <Link href="/login" className="text-primary hover:underline">
                  Log in
                </Link>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
  )
}
