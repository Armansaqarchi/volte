"use client"

import type React from "react"

import { useState } from "react"
import { useRouter } from "next/navigation"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { ArrowLeft, Plus, X, Clock } from "lucide-react"
import { getCurrentUser } from "@/lib/auth"
import { createEvent } from "@/lib/events"
import { useToast } from "@/hooks/use-toast"
import { DashboardNav } from "@/components/dashboard-nav"

export default function CreateEventPage() {
  const router = useRouter()
  const { toast } = useToast()
  const user = getCurrentUser()

  const [name, setName] = useState("")
  const [admin, setAdmin] = useState("")
  const [question, setQuestion] = useState("")
  const [months, setMonths] = useState("")
  const [days, setDays] = useState("")
  const [hours, setHours] = useState("")
  const [seconds, setSeconds] = useState("")
  const [options, setOptions] = useState(["", ""])
  const [isLoading, setIsLoading] = useState(false)

  const addOption = () => {
    setOptions([...options, ""])
  }

  const removeOption = (index: number) => {
    if (options.length > 2) {
      setOptions(options.filter((_, i) => i !== index))
    }
  }

  const updateOption = (index: number, value: string) => {
    const newOptions = [...options]
    newOptions[index] = value
    setOptions(newOptions)
  }

  const calculateTotalSeconds = (): number => {
    const m = Number.parseInt(months, 10) || 0
    const d = Number.parseInt(days, 10) || 0
    const h = Number.parseInt(hours, 10) || 0
    const s = Number.parseInt(seconds, 10) || 0

    // Convert everything to seconds: 1 month = 30 days
    return m * 30 * 24 * 60 * 60 + d * 24 * 60 * 60 + h * 60 * 60 + s
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsLoading(true)

    // Validate
    const validOptions = options.filter((o) => o.trim() !== "")
    if (validOptions.length < 2) {
      toast({
        title: "Invalid options",
        description: "Please provide at least 2 voting options.",
        variant: "destructive",
      })
      setIsLoading(false)
      return
    }
    const durationNum = calculateTotalSeconds()

    if (durationNum <= 0) {
      toast({
        title: "Invalid duration",
        description: "Please provide a valid duration for the event.",
        variant: "destructive",
      })
      setIsLoading(false)
      return
    }

    const event = await createEvent({
      id: "",
      name,
      question,
      admin,
      duration: durationNum,
      voteOptions: validOptions,
    })
    try {
      const user = getCurrentUser()
      console.log(JSON.stringify(event))
      const response = await fetch(`http://localhost:8000/users/events`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          vote_options: event.voteOptions,
          name: event.name,
          question: event.question,
          admin: user?.commitment,
          duration: durationNum,
        }),
        credentials: "include",
      })
      if (!response.ok) {
        throw new Error(`Failed to create event.}`)
      }

      toast({
        title: "Event created!",
        description: "Your voting event is now live.",
      })
      const respData = await response.json()
      setIsLoading(false)
      router.push(`/dashboard/events/${respData.event.id}`)
    } catch (e) {
      toast({
        title: `Failed to create the event`,
      })
    }
  }

  return (
      <div className="min-h-screen bg-background">
        <DashboardNav user={user} />

        <main className="container mx-auto px-4 py-8 max-w-3xl">
          <Button variant="ghost" onClick={() => router.back()} className="mb-6">
            <ArrowLeft className="h-4 w-4 mr-2" />
            Back
          </Button>

          <Card>
            <CardHeader>
              <CardTitle>Create Voting Event</CardTitle>
              <CardDescription>
                Set up a new event where participants can vote using zero-knowledge proofs
              </CardDescription>
            </CardHeader>
            <CardContent>
              <form onSubmit={handleSubmit} className="space-y-6">
                <div className="space-y-2">
                  <Label htmlFor="name">Name</Label>
                  <Input
                      id="name"
                      placeholder="Event name"
                      value={name}
                      onChange={(e) => setName(e.target.value)}
                      required
                  />
                </div>

                <div className="space-y-2">
                  <Label htmlFor="question">Question</Label>
                  <Input
                      id="question"
                      placeholder="Community Decision: Which feature to build next?"
                      value={question}
                      onChange={(e) => setQuestion(e.target.value)}
                      required
                  />
                </div>

                <div className="space-y-3">
                  <div className="flex items-center gap-2">
                    <Clock className="h-4 w-4 text-muted-foreground" />
                    <Label>Event Duration</Label>
                  </div>
                  <p className="text-sm text-muted-foreground">
                    Specify how long the event will remain active after starting
                  </p>

                  <div className="grid grid-cols-2 sm:grid-cols-4 gap-3">
                    <div className="space-y-2">
                      <Label htmlFor="months" className="text-xs text-muted-foreground">
                        Months
                      </Label>
                      <Input
                          id="months"
                          type="number"
                          min="0"
                          placeholder="0"
                          value={months}
                          onChange={(e) => setMonths(e.target.value)}
                          className="text-center"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="days" className="text-xs text-muted-foreground">
                        Days
                      </Label>
                      <Input
                          id="days"
                          type="number"
                          min="0"
                          max="30"
                          placeholder="0"
                          value={days}
                          onChange={(e) => setDays(e.target.value)}
                          className="text-center"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="hours" className="text-xs text-muted-foreground">
                        Hours
                      </Label>
                      <Input
                          id="hours"
                          type="number"
                          min="0"
                          max="23"
                          placeholder="0"
                          value={hours}
                          onChange={(e) => setHours(e.target.value)}
                          className="text-center"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="seconds" className="text-xs text-muted-foreground">
                        Seconds
                      </Label>
                      <Input
                          id="seconds"
                          type="number"
                          min="0"
                          max="59"
                          placeholder="0"
                          value={seconds}
                          onChange={(e) => setSeconds(e.target.value)}
                          className="text-center"
                      />
                    </div>
                  </div>

                  {(months || days || hours || seconds) && (
                      <div className="rounded-lg bg-muted/50 p-3 border border-border">
                        <p className="text-sm text-muted-foreground">
                          Total duration:{" "}
                          <span className="font-semibold text-foreground">{calculateTotalSeconds().toLocaleString()}</span>{" "}
                          seconds
                        </p>
                      </div>
                  )}
                </div>

                <div className="space-y-3">
                  <Label>Voting Options</Label>
                  <p className="text-sm text-muted-foreground">Add at least 2 options for participants to choose from</p>
                  {options.map((option, index) => (
                      <div key={index} className="flex gap-2">
                        <Input
                            placeholder={`Option ${index + 1}`}
                            value={option}
                            onChange={(e) => updateOption(index, e.target.value)}
                            required
                        />
                        {options.length > 2 && (
                            <Button type="button" variant="outline" size="icon" onClick={() => removeOption(index)}>
                              <X className="h-4 w-4" />
                            </Button>
                        )}
                      </div>
                  ))}
                  <Button type="button" variant="outline" onClick={addOption} className="w-full bg-transparent">
                    <Plus className="h-4 w-4 mr-2" />
                    Add Option
                  </Button>
                </div>

                <div className="flex gap-3 pt-4">
                  <Button type="button" variant="outline" onClick={() => router.back()} className="flex-1">
                    Cancel
                  </Button>
                  <Button type="submit" disabled={isLoading} className="flex-1">
                    {isLoading ? "Creating..." : "Create Event"}
                  </Button>
                </div>
              </form>
            </CardContent>
          </Card>
        </main>
      </div>
  )
}
