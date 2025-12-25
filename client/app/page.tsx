import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Card } from "@/components/ui/card"
import { Shield, Vote, Lock } from "lucide-react"

export default function HomePage() {
  return (
    <div className="min-h-screen bg-background">
      {/* Header */}
      <header className="border-b border-border">
        <div className="container mx-auto px-4 py-4 flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Shield className="h-6 w-6 text-primary" />
            <span className="font-semibold text-xl">ZK Vote</span>
          </div>
          <div className="flex items-center gap-3">
            <Link href="/login">
              <Button variant="ghost">Log in</Button>
            </Link>
            <Link href="/signup">
              <Button>Sign up</Button>
            </Link>
          </div>
        </div>
      </header>

      {/* Hero Section */}
      <main className="container mx-auto px-4 py-16">
        <div className="max-w-4xl mx-auto text-center space-y-6">
          <div className="inline-flex items-center gap-2 px-4 py-2 bg-primary/10 rounded-full text-sm font-medium text-primary mb-4">
            <Lock className="h-4 w-4" />
            Powered by Zero-Knowledge Proofs
          </div>

          <h1 className="text-5xl font-bold tracking-tight text-balance">Secure Voting with Privacy</h1>

          <p className="text-xl text-muted-foreground text-balance max-w-2xl mx-auto">
            Create and participate in verifiable elections while maintaining complete voter anonymity using Groth16
            zero-knowledge proofs.
          </p>

          <div className="flex items-center justify-center gap-4 pt-4">
            <Link href="/signup">
              <Button size="lg">Get Started</Button>
            </Link>
            <Link href="/login">
              <Button size="lg" variant="outline">
                View Events
              </Button>
            </Link>
          </div>
        </div>

        {/* Features */}
        <div className="grid md:grid-cols-3 gap-6 mt-20 max-w-5xl mx-auto">
          <Card className="p-6 space-y-3">
            <div className="h-12 w-12 rounded-lg bg-primary/10 flex items-center justify-center">
              <Shield className="h-6 w-6 text-primary" />
            </div>
            <h3 className="font-semibold text-lg">Zero-Knowledge Privacy</h3>
            <p className="text-muted-foreground text-sm leading-relaxed">
              Vote without revealing your identity. Groth16 proofs ensure your vote is valid while keeping you
              anonymous.
            </p>
          </Card>

          <Card className="p-6 space-y-3">
            <div className="h-12 w-12 rounded-lg bg-primary/10 flex items-center justify-center">
              <Vote className="h-6 w-6 text-primary" />
            </div>
            <h3 className="font-semibold text-lg">Create Events</h3>
            <p className="text-muted-foreground text-sm leading-relaxed">
              Launch voting events with custom options. Set eligibility criteria and manage your elections.
            </p>
          </Card>

          <Card className="p-6 space-y-3">
            <div className="h-12 w-12 rounded-lg bg-primary/10 flex items-center justify-center">
              <Lock className="h-6 w-6 text-primary" />
            </div>
            <h3 className="font-semibold text-lg">Verifiable Results</h3>
            <p className="text-muted-foreground text-sm leading-relaxed">
              Every vote is cryptographically verified. Results are tamper-proof and publicly auditable.
            </p>
          </Card>
        </div>
      </main>
    </div>
  )
}
