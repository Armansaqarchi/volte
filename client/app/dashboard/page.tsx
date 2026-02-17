"use client"

import { useEffect, useState } from "react"
import { useRouter } from "next/navigation"
import Link from "next/link"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Plus, Calendar, Vote } from "lucide-react"
import { getCurrentUser } from "@/lib/auth"
import {getEvents, type Event, IsEventActive} from "@/lib/events"
import { DashboardNav } from "@/components/dashboard-nav"

export default function DashboardPage() {
  const router = useRouter()
  const [user, setUser] = useState(() => getCurrentUser())
  const [events, setEvents] = useState<Event[]>([])

  useEffect(() => {
    async function loadEvents() {
      if (!user) {
        router.push("/login")
        return
      }

      const events = await getEvents()
      console.log(events)
      setEvents(events)
    }
    loadEvents()
  }, [user, router])

  if (!user) return null

  const myEvents = events.filter((e) => e.admin == user.commitment)
  const activeEvents = events.filter((e) => IsEventActive(e))
  return (
    <div className="min-h-screen bg-background">
      <DashboardNav user={user} />

      <main className="container mx-auto px-4 py-8">
        <div className="flex items-center justify-between mb-8">
          <div>
            <h1 className="text-3xl font-bold">Dashboard</h1>
            <p className="text-muted-foreground">Manage your voting events</p>
          </div>
          <Link href="/dashboard/events/create">
            <Button>
              <Plus className="h-4 w-4 mr-2" />
              Create Event
            </Button>
          </Link>
        </div>

        {/* Stats */}
        <div className="grid md:grid-cols-3 gap-6 mb-8">
          <Card>
            <CardHeader className="flex flex-row items-center justify-between pb-2">
              <CardTitle className="text-sm font-medium">My Events</CardTitle>
              <Calendar className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{myEvents.length}</div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader className="flex flex-row items-center justify-between pb-2">
              <CardTitle className="text-sm font-medium">Active Events</CardTitle>
              <Vote className="h-4 w-4 text-muted-foreground" />
            </CardHeader>
            <CardContent>
              <div className="text-2xl font-bold">{activeEvents.length}</div>
            </CardContent>
          </Card>
        </div>

        {/* My Events */}
        <div className="mb-8">
          <h2 className="text-xl font-semibold mb-4">My Events</h2>
          {myEvents.length === 0 ? (
            <Card>
              <CardContent className="py-12 text-center">
                <p className="text-muted-foreground mb-4">You haven't created any events yet.</p>
                <Link href="/dashboard/events/create">
                  <Button>Create Your First Event</Button>
                </Link>
              </CardContent>
            </Card>
          ) : (
            <div className="grid md:grid-cols-2 gap-6">
              {myEvents.map((event) => (
                  <EventCard key={event.id} event={event} isOwner/>
              ))}
            </div>
          )}
        </div>

        {/* Active Events */}
        <div>
          <h2 className="text-xl font-semibold mb-4">Active Events</h2>
          {activeEvents.length === 0 ? (
            <Card>
              <CardContent className="py-12 text-center">
                <p className="text-muted-foreground">No active events at the moment.</p>
              </CardContent>
            </Card>
          ) : (
            <div className="grid md:grid-cols-2 gap-6">
              {activeEvents.map((event) => (
                <EventCard key={event.id} event={event} isOwner={event.admin === user.commitment} />
              ))}
            </div>
          )}
        </div>
      </main>
    </div>
  )
}

function EventCard({ event, isOwner }: { event: Event; isOwner: boolean }) {
  return (
    <Link href={`/dashboard/events/${event.id}`}>
      <Card className="hover:border-primary transition-colors cursor-pointer">
        <CardHeader>
          <div className="flex items-start justify-between">
            <div className="space-y-1">
              <CardTitle className="text-lg">{event.name}</CardTitle>
            </div>
            <Badge variant={IsEventActive(event) ? "default" : "secondary"}>{IsEventActive(event) ? "active" : "inactive"}</Badge>
          </div>
        </CardHeader>
        <CardContent>
          <div className="flex items-center gap-4 text-sm text-muted-foreground">
            <div className="flex items-center gap-1">
              {event.startTime && (
                <>
                  <Calendar className="h-4 w-4"/>
                  <span>
                  {"Ends " + new Date(new Date(event.startTime).getTime() + event.duration).toLocaleDateString()}
                  </span>
                </>
              )}
            </div>
          </div>
          {isOwner && (
            <Badge variant="outline" className="mt-3">
              Owner
            </Badge>
          )}
        </CardContent>
      </Card>
    </Link>
  )
}
