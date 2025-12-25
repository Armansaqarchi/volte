"use client"

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { RadioGroup, RadioGroupItem } from "@/components/ui/radio-group"
import { Label } from "@/components/ui/label"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Shield, Loader2, CheckCircle2 } from "lucide-react"
import type { Event } from "@/lib/events"
import { castVote } from "@/lib/votes"
import { useToast } from "@/hooks/use-toast"

interface VotingInterfaceProps {
  event: Event
  userId: string
  onVoteComplete: () => void
}

export function VotingInterface({ event, userId, onVoteComplete }: VotingInterfaceProps) {
  const { toast } = useToast()
  const [selectedOption, setSelectedOption] = useState<string>("")
  const [isGeneratingProof, setIsGeneratingProof] = useState(false)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [zkProof, setZkProof] = useState<any>(null)

  // ============================================
  // ZK PROOF INTEGRATION POINT #1: PROOF GENERATION
  // ============================================
  // This function is called when the user clicks "Generate Proof"
  // Replace this with your actual Groth16 proof generation logic
  const handleGenerateProof = async () => {
    if (!selectedOption) {
      toast({
        title: "No option selected",
        description: "Please select an option before generating proof.",
        variant: "destructive",
      })
      return
    }

    setIsGeneratingProof(true)

    try {
      console.log("[v0] Starting ZK proof generation...")
      console.log("[v0] Event ID:", event.id)
      console.log("[v0] User ID:", userId)
      console.log("[v0] Selected Option:", selectedOption)

      // ⚠️ IMPLEMENT YOUR GROTH16 PROOF GENERATION HERE ⚠️
      //
      // Example of what you might do:
      // const proof = await generateGroth16Proof({
      //   eventId: event.id,
      //   userId: userId,
      //   vote: selectedOption,
      //   eligibilityToken: userEligibilityToken,
      // })
      //
      // The proof should demonstrate:
      // 1. User is eligible to vote in this event
      // 2. User hasn't voted before
      // 3. Vote is for a valid option
      // Without revealing which user cast which vote

      // Simulate proof generation (remove this in production)
      await new Promise((resolve) => setTimeout(resolve, 2000))

      const mockProof = {
        proof: {
          pi_a: ["mock_a_1", "mock_a_2"],
          pi_b: [
            ["mock_b_1", "mock_b_2"],
            ["mock_b_3", "mock_b_4"],
          ],
          pi_c: ["mock_c_1", "mock_c_2"],
        },
        publicSignals: [event.id, "commitment_hash"],
        timestamp: Date.now(),
      }

      setZkProof(mockProof)

      console.log("[v0] ZK proof generated successfully!")
      console.log("[v0] Proof:", mockProof)

      toast({
        title: "Proof generated",
        description: "Your zero-knowledge proof is ready to submit.",
      })
    } catch (error) {
      console.error("[v0] Proof generation failed:", error)
      toast({
        title: "Proof generation failed",
        description: "Please try again or contact support.",
        variant: "destructive",
      })
    } finally {
      setIsGeneratingProof(false)
    }
  }

  // ============================================
  // ZK PROOF INTEGRATION POINT #2: PROOF SUBMISSION
  // ============================================
  // This function is called when the user submits their vote with the proof
  // Replace this with your actual proof verification and vote recording logic
  const handleSubmitVote = async () => {
    if (!zkProof) return

    setIsSubmitting(true)

    try {
      console.log("[v0] Submitting vote with ZK proof...")
      console.log("[v0] Proof:", zkProof)

      // ⚠️ IMPLEMENT YOUR PROOF VERIFICATION HERE ⚠️
      //
      // Example of what you might do:
      // const isValid = await verifyGroth16Proof({
      //   proof: zkProof.proof,
      //   publicSignals: zkProof.publicSignals,
      //   verificationKey: eventVerificationKey,
      // })
      //
      // if (!isValid) {
      //   throw new Error('Proof verification failed')
      // }
      //
      // Then submit the vote to your backend:
      // await submitVote({
      //   eventId: event.id,
      //   proof: zkProof,
      //   // Note: Don't send selectedOption or userId to backend
      //   // The proof itself contains the commitment
      // })

      // Simulate submission (remove this in production)
      await new Promise((resolve) => setTimeout(resolve, 1500))

      // Record vote locally (this would be done on your backend in production)
      castVote(event.id, userId, selectedOption)

      console.log("[v0] Vote submitted and verified successfully!")

      toast({
        title: "Vote recorded!",
        description: "Your vote has been verified and counted.",
      })

      onVoteComplete()
    } catch (error) {
      console.error("[v0] Vote submission failed:", error)
      toast({
        title: "Submission failed",
        description: "Please try again or contact support.",
        variant: "destructive",
      })
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle>Cast Your Vote</CardTitle>
        <CardDescription>Select an option and generate a zero-knowledge proof to vote anonymously</CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        {/* Option Selection */}
        <RadioGroup value={selectedOption} onValueChange={setSelectedOption}>
          <div className="space-y-3">
            {event.options.map((option) => (
              <div key={option} className="flex items-center space-x-3 space-y-0">
                <RadioGroupItem value={option} id={option} />
                <Label htmlFor={option} className="font-normal cursor-pointer flex-1 py-3">
                  {option}
                </Label>
              </div>
            ))}
          </div>
        </RadioGroup>

        {/* ZK Proof Info */}
        <Alert>
          <Shield className="h-4 w-4" />
          <AlertDescription className="text-sm">
            Your vote will be encrypted using Groth16 zero-knowledge proofs. This ensures your choice remains private
            while proving you're eligible to vote.
          </AlertDescription>
        </Alert>

        {/* Proof Status */}
        {zkProof && (
          <div className="p-4 bg-primary/5 border border-primary/20 rounded-lg space-y-2">
            <div className="flex items-center gap-2 text-primary">
              <CheckCircle2 className="h-4 w-4" />
              <span className="font-medium text-sm">Zero-Knowledge Proof Generated</span>
            </div>
            <p className="text-xs text-muted-foreground font-mono">
              Proof Hash: {zkProof.publicSignals[1].substring(0, 32)}...
            </p>
          </div>
        )}

        {/* Action Buttons */}
        <div className="flex flex-col sm:flex-row gap-3">
          {!zkProof ? (
            <Button onClick={handleGenerateProof} disabled={!selectedOption || isGeneratingProof} className="flex-1">
              {isGeneratingProof ? (
                <>
                  <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                  Generating Proof...
                </>
              ) : (
                <>
                  <Shield className="h-4 w-4 mr-2" />
                  Generate Proof
                </>
              )}
            </Button>
          ) : (
            <>
              <Button
                variant="outline"
                onClick={() => {
                  setZkProof(null)
                  setSelectedOption("")
                }}
                className="flex-1"
              >
                Cancel
              </Button>
              <Button onClick={handleSubmitVote} disabled={isSubmitting} className="flex-1">
                {isSubmitting ? (
                  <>
                    <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                    Submitting...
                  </>
                ) : (
                  "Submit Vote"
                )}
              </Button>
            </>
          )}
        </div>

        {/* Developer Note */}
        <div className="pt-4 border-t text-xs text-muted-foreground">
          <p className="font-medium mb-1">Developer Integration Points:</p>
          <ul className="list-disc list-inside space-y-1">
            <li>
              Proof generation in <code className="text-xs bg-muted px-1 py-0.5 rounded">handleGenerateProof()</code>
            </li>
            <li>
              Proof verification in <code className="text-xs bg-muted px-1 py-0.5 rounded">handleSubmitVote()</code>
            </li>
            <li>Check console logs for integration details</li>
          </ul>
        </div>
      </CardContent>
    </Card>
  )
}
