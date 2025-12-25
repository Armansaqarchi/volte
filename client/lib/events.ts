// Event management utilities

import {getCurrentUser, User} from "./auth"

export interface Event {
  id          : string
  question    : string
  name        : string
  admin       : string
  duration    : number
  startTime   : Date | null
  voteOptions : string[]
  voteMembers : string[]
  tally       : Map<string, number>
  revoked     : boolean
}

export function IsEventActive(e: Event): boolean{
  const now = new Date()
  return e.startTime? now > e.startTime && now < new Date(e.startTime.getTime() + e.duration) : false
}

export function IsEventStarted(e: Event): boolean{
  const now = new Date()
  return e.startTime? now > e.startTime : false
}

export async function getEvents(): Promise<Event[]> {
  if (typeof window === "undefined") return []
  const user = getCurrentUser()
  console.log("fetched user ", user)
  if (user === null) {
    return []
  }
  const response = await fetch(`http://localhost:8000/users/events`, {
    method: "GET",
    credentials: "include"
  })
  if (!response.ok) throw new Error("Failed to fetch")
  const events = await response.json()
  if (events !== null && events.data !== null && events.data.length > 0) {
    return events.data
  } else{
    return []
  }
}

export async function getEvent(eventId: string): Promise<Event | null> {
  if (typeof window === "undefined") return null
  const user = getCurrentUser()
  if (user === null) {
    return null
  }
  const response = await fetch(`http://localhost:8000/users/event/${eventId}`, {
    method: "GET",
    credentials: "include"
  })
  if (!response.ok) throw new Error("Failed to fetch")
  const data = await response.json()
  if (data.startTime !== null){
    data.startTime = new Date(data.startTime)
  } else{
    data.startTime = null
  }
  return data
}

export async function createEvent(data: {
  id: string
  name: string
  question: string
  admin: string
  duration    : number
  voteOptions : string[]
}): Promise<Event> {
  const events = await getEvents()

  const newEvent: Event = {
    id: data.id,
    name: data.name,
    question: data.question,
    admin: data.admin,
    duration: data.duration,
    startTime: null,
    voteOptions: data.voteOptions,
    voteMembers: [],
    tally: new Map<string, number>(),
    revoked: false
  }

  events.push(newEvent)
  localStorage.setItem("events", JSON.stringify(events))

  return newEvent
}

export async function updateEvent(id: string, updates: Partial<Event>): Promise<Event | null> {
  const events = await getEvents()
  const index = events.findIndex((e) => e.id === id)

  if (index === -1) return null

  events[index] = { ...events[index], ...updates }
  localStorage.setItem("events", JSON.stringify(events))

  return events[index]
}
