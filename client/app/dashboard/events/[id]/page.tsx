"use client"

import { useEffect, useState } from "react"
import { useRouter, useParams } from "next/navigation"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { Input } from "@/components/ui/input"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog"
import {
    ArrowLeft,
    Calendar,
    Plus,
    Users,
    MessageCircle,
    CheckCircle2,
    Clock,
    Trash2,
    Play,
    AlertTriangle,
    X,
    BarChart3,
} from "lucide-react"
import {getCommitment, getCurrentUser, MimC7Hash, type User} from "@/lib/auth"
import { getEvent, type Event, IsEventActive, IsEventStarted } from "@/lib/events"
import { DashboardNav } from "@/components/dashboard-nav"
import { toast } from "@/hooks/use-toast"
import { LoadingSpinner } from "@/components/ui/loading-spinner"
import { runWasm } from "@/lib/zkproof/go/main"
import { ChartContainer, ChartTooltip } from "@/components/ui/chart"
import { PieChart, Pie, Cell, ResponsiveContainer } from "recharts"
import generateBallotProof from "@/lib/zkproof/circom/ballot/generate_proof.mjs"
import generateMerkleProof from "@/lib/zkproof/circom/merkletree/generator.mjs";
import generateNullifier from "@/lib/zkproof/circom/nullifier/generateNullifier.mjs";

type VoteResults = {
    option: string
    votes: number
    percentage: number
    color: string
}

const CHART_COLORS = [
    "hsl(221, 83%, 53%)", // Blue
    "hsl(142, 76%, 36%)", // Green
    "hsl(262, 83%, 58%)", // Purple
    "hsl(31, 97%, 52%)", // Orange
    "hsl(346, 87%, 49%)", // Red
]

export default function EventDetailPage() {
    const router = useRouter()
    const params = useParams<{ id: string }>()
    const id = params.id

    const [user, setUser] = useState<User | null>(null)
    const [event, setEvent] = useState<Event | null>(null)
    const [addedUserCommitment, setAddedUserCommitment] = useState("")
    const [eventMembers, setEventMembers] = useState<string[]>([])
    const [isAddingMember, setIsAddingMember] = useState(false)
    const [isSubmittingVote, setIsSubmittingVote] = useState(false)
    const [IsEventLoading, setIsEventLoading] = useState(true)
    const [isDeletingMember, setIsDeletingMember] = useState<string | null>(null)
    const [isDeletingEvent, setIsDeletingEvent] = useState(false)
    const [isStartingEvent, setIsStartingEvent] = useState(false)
    const [errorMessage, setErrorMessage] = useState<string | null>(null)
    const [isResultsDialogOpen, setIsResultsDialogOpen] = useState(false)
    const [voteResults, setVoteResults] = useState<VoteResults[] | null>(null)
    const [isLoadingResults, setIsLoadingResults] = useState(false)
    const [startConfirmation, setStartConfirmation] = useState<{ open: boolean; confirmText: string; error?: string }>({
        open: false,
        confirmText: "",
        error: undefined,
    })
    const [deleteConfirmation, setDeleteConfirmation] = useState<{ open: boolean; confirmText: string; error?: string }>({
        open: false,
        confirmText: "",
        error: undefined,
    })
    const [voteConfirmation, setVoteConfirmation] = useState<{
        open: boolean
        selectedOption: number | null
        confirmText: string
        error?: string
    }>({
        open: false,
        selectedOption: null,
        confirmText: "",
        error: undefined,
    })
    const [addMemberError, setAddMemberError] = useState<string | null>(null)
    const [isAddMemberDialogOpen, setIsAddMemberDialogOpen] = useState(false)
    const [isForceEndingEvent, setIsForceEndingEvent] = useState(false)
    const [forceEndConfirmation, setForceEndConfirmation] = useState<{ open: boolean; confirmText: string; error?: string }>({
        open: false,
        confirmText: "",
        error: undefined,
    })

    useEffect(() => {
        if (!id) return

        const u = getCurrentUser()
        if (!u) {
            router.push("/login")
            return
        }
        setUser(u)

        async function loadEvent() {
            const loaded = await getEvent(id)
            if (!loaded) {
                router.push("/dashboard")
                return
            }
            setEvent(loaded)
            setIsEventLoading(false)
        }

        loadEvent()
    }, [id, router])

    const handleViewResults = async () => {
        setIsResultsDialogOpen(true)
        setIsLoadingResults(true)

        try {
            const response = await fetch(`http://localhost:8000/event/${event?.id}/tally`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            })
            if (!response.ok) {
                const errorData = await response.json()
                const error = errorData.error || "Failed to fetch results"
                setErrorMessage(error)
                toast({ title: "Failed to load results" })
                return
            }
            const data = await response.json()

            const score = Number(data.score)
            const total = Number(data.total)
            const secondOptionVotes = total - score

            const transformedResults: VoteResults[] = [
                {
                    option: event?.voteOptions?.[0] || "Option 1",
                    votes: score,
                    percentage: total > 0 ? Math.round((score / total) * 100) : 0,
                    color: CHART_COLORS[0],
                },
                {
                    option: event?.voteOptions?.[1] || "Option 2",
                    votes: secondOptionVotes,
                    percentage: total > 0 ? Math.round((secondOptionVotes / total) * 100) : 0,
                    color: CHART_COLORS[1],
                },
            ]

            setVoteResults(transformedResults)
        } catch (error) {
            const message = error instanceof Error ? error.message : "An error occurred"
            setErrorMessage(message)
            toast({ title: "Error loading results" })
        } finally {
            setIsLoadingResults(false)
        }
    }

    const handleAddMember = async () => {
        if (!addedUserCommitment.trim()) return

        setIsAddingMember(true)
        setAddMemberError(null)

        try {
            const response = await fetch(`http://localhost:8000/event/${event?.id}/members/${addedUserCommitment}`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            })

            if (!response.ok) {
                if (response.status === 404) {
                    const error = "User not found. Please check the User ID and try again."
                    setAddMemberError(error)
                    setErrorMessage(error)
                    toast({
                        title: "User not found",
                        description: "The User ID you entered does not exist.",
                        variant: "destructive",
                    })
                    return
                }

                const errorData = await response.json()
                const error = errorData.error || "Failed to add member"
                setAddMemberError(error)
                setErrorMessage(error)
                toast({ title: "Failed to add member" })
                return
            }

            if (response.status === 208) {
                const error = "User is already a member of the event."
                setAddMemberError(error)
                setErrorMessage(error)
                toast({
                    title: "User is a member of the event",
                    description: "The commitment belongs to a user that is already a member of this event",
                })
                return
            }

            toast({ title: "User added" })
            setEventMembers((prev) => [...prev, addedUserCommitment])
            setAddedUserCommitment("")
            setIsAddMemberDialogOpen(false)
            setAddMemberError(null)
        } catch (error) {
            const message = error instanceof Error ? error.message : "An error occurred"
            setAddMemberError(message)
            setErrorMessage(message)
            toast({ title: "Error adding member" })
        } finally {
            setIsAddingMember(false)
        }
    }

    const handleDeleteMember = async (memberToDelete: string) => {
        setIsDeletingMember(memberToDelete)

        try {
            const response = await fetch(`http://localhost:8000/event/${event?.id}/members/${memberToDelete}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            })

            if (!response.ok) {
                toast({ title: "Failed to delete member" })
                throw new Error("Failed to delete member")
            }

            toast({ title: "Member removed" })
            setEvent((prev) => {
                if (!prev) return prev
                return {
                    ...prev,
                    voteMembers: prev.voteMembers.filter((m) => m !== memberToDelete),
                }
            })
        } catch (error) {
            toast({ title: "Error deleting member" })
        } finally {
            setIsDeletingMember(null)
        }
    }

    const handleDeleteEvent = async () => {
        setIsDeletingEvent(true)

        try {
            const response = await fetch(`http://localhost:8000/event/${event?.id}`, {
                method: "DELETE",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            })

            console.log(await response.json())

            if (!response.ok) {
                const errorData = await response.json()
                const error = errorData.error || "Failed to delete event"
                setDeleteConfirmation((prev) => ({ ...prev, error }))
                setErrorMessage(error)
                toast({ title: "Failed to delete event" })
                return
            }

            toast({ title: "Event deleted successfully" })
            router.push("/dashboard")
        } catch (error) {
            const message = error instanceof Error ? error.message : "An error occurred"
            setDeleteConfirmation((prev) => ({ ...prev, error: message }))
            setErrorMessage(message)
            toast({ title: "Error deleting event" })
            setIsDeletingEvent(false)
        }
    }

    function merklePathToBits(path: number, depth: number): string[] {
        const bits: string[] = new Array(depth)
        for (let i = 0; i < depth; i++) {
            bits[i] = String(((path >> i) & 1) ^ 1 /* toggle the bit */)
        }
        return bits
    }

    function toCommaSeparated<T>(arr: T[]): string {
        return arr.map(String).join(",")
    }

    function uuidToBigInt(input: string): bigint {
        const cleaned = input.replace(/[-/]/g, "").trim()
        if (cleaned.length === 0) {
            throw new Error("Empty id after removing separators")
        }
        if (!/^[0-9a-fA-F]+$/.test(cleaned)) {
            throw new Error("ID must be hex after removing separators")
        }
        return BigInt("0x" + cleaned)
    }

    async function handleVote(option: number) {
        try {
            let response = await fetch(`http://localhost:8000/event/${event?.id}/membership/merkle`, {
                method: "GET",
                credentials: "include",
            })

            if (!response.ok) {
                const errorData = await response.json()
                const error = errorData.error || "Failed to retrieve membership proof"
                setErrorMessage(error)
                toast({ title: "Failed to retrieve membership proof" })
                return
            }

            const jsonBody = await response.json()
            console.log(jsonBody)

            const membershipPath: string[] = []
            jsonBody.data.proof.Siblings.forEach((sibling: string) => membershipPath.push((atob(sibling))))
            const paths = merklePathToBits(jsonBody.data.proof.Path, jsonBody.data.proof.Siblings.length)
            console.log(paths)
            console.log(uuidToBigInt(event?.id as string))
            console.log("root : ", atob(jsonBody.data.root).toString())

            const argv: string[] = [
                "--Yx=1527465159374431915328497116935179161014331322368960485951268517950184093102",
                "--Yy=17274044707157828649723710289902216429715848248207037129568326237800068062774",
                `--user_secret_key=${localStorage.getItem("privateKey")}`,
                `--event_id=${uuidToBigInt(event?.id as string)}`,
                `--membership_merkle_path=${toCommaSeparated(membershipPath)}`,
                `--membership_path_positions=${paths}`,
                `--membership_merkle_root=${atob(jsonBody.data.root)}`,
                `--msg=${option}`,
                "--Gx=1",
                "--Gy=2",
            ]

            console.log(argv)

            // const merkle = await runWasm("runMerkle", argv)
            // const nullifier = await runWasm("runNullifier", argv)
            const m = await generateBallotProof(option)
            console.log( {
                "MerkleRoot": `${atob(jsonBody.data.root)}`,
                "LeafValue": getCommitment(),
                "MerklePath": membershipPath,
                "PathPositions": paths,
                "SecretKey": `${localStorage.getItem("privateKey")}`
            },)
            const m2 = await generateMerkleProof(
                {
                    "MerkleRoot": `${atob(jsonBody.data.root)}`,
                    "LeafValue": getCommitment(),
                    "MerklePath": membershipPath,
                    "PathPositions": paths,
                    "SecretKey": `${localStorage.getItem("privateKey")}`
                },
            )
            console.log("generated : ", m2)

            const nullifier = await MimC7Hash([BigInt(localStorage.getItem("privateKey")), uuidToBigInt(event?.id as string)])
            console.log("nullifier : ", nullifier)
            console.log({
                "Commitment": getCommitment(),
                "EventID": `${uuidToBigInt(event?.id as string)}`,
                "Nullifier": nullifier,
                "SecretKey": `${localStorage.getItem("privateKey")}`
            })
            const m3 = await generateNullifier({
                "Commitment": getCommitment(),
                "EventID": `${uuidToBigInt(event?.id as string)}`,
                "Nullifier": nullifier,
                "SecretKey": `${localStorage.getItem("privateKey")}`
            })
            console.log("m3 : ", m3)



            const proof = {
                ballot: m,
                membership: m2,
                nullifier: m3
            }

            console.log(JSON.stringify({
                "EventID": event?.id,
                "Proofs": proof
            }))

            response = await fetch(`http://localhost:8000/event/${event?.id}/vote`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    "EventID": event?.id,
                    "Proofs": proof
                }),
                credentials: "include",
            })

            console.log(response.json())
        } catch (error) {
            const message = error instanceof Error ? error.message : "An error occurred"
            setErrorMessage(message)
        }
    }

    function handleStartEventClick() {
        setStartConfirmation({ open: true, confirmText: "", error: undefined })
    }

    async function submitStartEventAfterConfirmation() {
        if (startConfirmation.confirmText.toLowerCase() !== "start") {
            setStartConfirmation((prev) => ({ ...prev, error: "Please type 'start' to proceed" }))
            return
        }
        await handleConfirmStartEvent()
    }

    async function handleConfirmStartEvent() {
        setIsStartingEvent(true)

        try {
            const response = await fetch(`http://localhost:8000/event/${event?.id}/start`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            })

            if (!response.ok) {
                const errorData = await response.json()
                const error = errorData.error || "Failed to start event"
                setStartConfirmation((prev) => ({ ...prev, error }))
                setErrorMessage(error)
                toast({ title: "Failed to start event" })
                setIsStartingEvent(false)
                return
            }

            const data: { start: string } = await response.json()
            setEvent((prev) =>
                prev
                    ? {
                        ...prev,
                        startTime: new Date(data.start),
                    }
                    : prev,
            )

            toast({ title: "Event successfully started." })
            setStartConfirmation({ open: false, confirmText: "", error: undefined })
            setIsStartingEvent(false)
        } catch (error) {
            const message = error instanceof Error ? error.message : "An error occurred"
            setStartConfirmation((prev) => ({ ...prev, error: message }))
            setErrorMessage(message)
            toast({ title: "Failed to start event" })
            setIsStartingEvent(false)
        }
    }

    async function submitVoteAfterConfirmation() {
        if (voteConfirmation.confirmText.toLowerCase() !== "confirm") {
            setVoteConfirmation((prev) => ({ ...prev, error: "Please type 'confirm' to proceed" }))
            return
        }

        setIsSubmittingVote(true)

        try {
            if (
                voteConfirmation.selectedOption != null &&
                event &&
                event.voteOptions &&
                voteConfirmation.selectedOption >= 0 &&
                voteConfirmation.selectedOption < event?.voteOptions.length
            ) {

                await handleVote(voteConfirmation.selectedOption)
            } else {
                setVoteConfirmation((prev) => ({ ...prev, error: "Option not selected or is invalid" }))
            }
        } catch (error) {
        } finally {
            setIsSubmittingVote(false)
        }
    }

    async function submitDeleteEventAfterConfirmation() {
        if (deleteConfirmation.confirmText.toLowerCase() !== "delete") {
            setDeleteConfirmation((prev) => ({ ...prev, error: "Please type 'delete' to proceed" }))
            return
        }
        await handleDeleteEvent()
    }

    function handleForceEndEventClick() {
        setForceEndConfirmation({ open: true, confirmText: "", error: undefined })
    }

    async function submitForceEndEventAfterConfirmation() {
        if (forceEndConfirmation.confirmText.toLowerCase() !== "end") {
            setForceEndConfirmation((prev) => ({ ...prev, error: "Please type 'end' to proceed" }))
            return
        }
        await handleConfirmForceEndEvent()
    }

    async function handleConfirmForceEndEvent() {
        setIsForceEndingEvent(true)

        try {
            const response = await fetch(`http://localhost:8000/event/${event?.id}/end`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
            })

            if (!response.ok) {
                const errorData = await response.json()
                const error = errorData.error || "Failed to force end event"
                setForceEndConfirmation((prev) => ({ ...prev, error }))
                setErrorMessage(error)
                toast({ title: "Failed to force end event" })
                setIsForceEndingEvent(false)
                return
            }

            const data: { forceEnd: boolean } = await response.json()
            setEvent((prev) =>
                prev
                    ? {
                        ...prev,
                        forceEnd: data.forceEnd,
                    }
                    : prev,
            )

            toast({ title: "Event successfully ended." })
            setForceEndConfirmation({ open: false, confirmText: "", error: undefined })
            setIsForceEndingEvent(false)

            // Refresh the page to recalculate isFinished and update all components
            setTimeout(() => {
                window.location.reload()
            }, 1000)
        } catch (error) {
            const message = error instanceof Error ? error.message : "An error occurred"
            setForceEndConfirmation((prev) => ({ ...prev, error: message }))
            setErrorMessage(message)
            toast({ title: "Error force ending event" })
            setIsForceEndingEvent(false)
        }
    }

    const isEventFinished = () => {
        if (!event?.startTime) return false
        const now = new Date().getTime()
        const startTime = new Date(event.startTime).getTime()
        const endTime = startTime + event.duration * 1000
        return now >= endTime || event.forceEnd
    }

    if (IsEventLoading) {
        return (
            <div className="min-h-screen flex flex-col items-center justify-center gap-4">
                <LoadingSpinner size="lg" text="Loading event…" />
            </div>
        )
    }

    if (!event || !user) {
        router.push("/dashboard")
        return
    }

    const isOwner = event?.admin === user?.commitment
    const isActive = IsEventActive(event)
    const isStarted = IsEventStarted(event)
    const isMember = event?.voteMembers?.includes(user?.commitment)
    const isFinished = isEventFinished()

    return (
        <div className="min-h-screen bg-background">
            <DashboardNav user={user} />

            <main className="container mx-auto px-4 py-6 max-w-7xl">
                {errorMessage && (
                    <div className="mb-6 rounded-lg border border-red-200 dark:border-red-800 bg-red-50 dark:bg-red-950/20 p-4 flex items-start gap-3">
                        <AlertTriangle className="h-5 w-5 text-red-600 dark:text-red-400 flex-shrink-0 mt-0.5" />
                        <div className="flex-1">
                            <p className="text-sm font-semibold text-red-800 dark:text-red-200">Error</p>
                            <p className="text-sm text-red-700 dark:text-red-300">{errorMessage}</p>
                        </div>
                        <button
                            onClick={() => setErrorMessage(null)}
                            className="text-red-600 dark:text-red-400 hover:text-red-700 dark:hover:text-red-300 flex-shrink-0"
                        >
                            <X className="h-4 w-4" />
                        </button>
                    </div>
                )}

                <Button variant="ghost" onClick={() => router.push("/dashboard")} className="mb-6 -ml-2 h-8">
                    <ArrowLeft className="h-4 w-4 mr-2" />
                    Back
                </Button>

                <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                    <div className="lg:col-span-2 space-y-5">
                        <div className="rounded-lg border border-border bg-card p-6">
                            <div className="flex items-start justify-between gap-4 mb-4">
                                <div className="flex-1 min-w-0">
                                    <div className="flex items-center gap-3 flex-wrap mb-3">
                                        <h1 className="text-3xl font-bold text-foreground">{event.name}</h1>
                                        <Badge variant={IsEventActive(event) ? "default" : "secondary"} className="h-fit">
                                            {IsEventActive(event) ? "Active" : "Inactive"}
                                        </Badge>
                                    </div>
                                    {isOwner && (
                                        <Badge variant="outline" className="w-fit text-xs">
                                            Creator
                                        </Badge>
                                    )}
                                </div>

                                <div className="flex flex-col gap-3 flex-shrink-0 w-full sm:w-auto">
                                    {isFinished && (
                                        <button
                                            onClick={handleViewResults}
                                            className="group relative inline-flex items-center justify-center gap-2 px-4 py-2 rounded-lg font-semibold text-white bg-gradient-to-r from-blue-500 to-indigo-600 hover:from-blue-600 hover:to-indigo-700 active:scale-95 transition-all duration-200 shadow-lg hover:shadow-xl"
                                        >
                                            <BarChart3 className="h-4 w-4" />
                                            <span className="text-sm">View Results</span>
                                            <div className="absolute inset-0 rounded-lg bg-white opacity-0 group-hover:opacity-10 transition-opacity duration-200" />
                                        </button>
                                    )}

                                    {isOwner && !isStarted && (
                                        <button
                                            onClick={handleStartEventClick}
                                            disabled={isStartingEvent}
                                            className="group relative inline-flex items-center justify-center gap-2 px-4 py-2 rounded-lg font-semibold text-white bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 active:scale-95 transition-all duration-200 shadow-lg hover:shadow-xl disabled:opacity-50 disabled:cursor-not-allowed"
                                        >
                                            {isStartingEvent ? (
                                                <LoadingSpinner size="sm" />
                                            ) : (
                                                <>
                                                    <Play className="h-4 w-4" />
                                                    <span className="text-sm">Start</span>
                                                </>
                                            )}
                                            <div className="absolute inset-0 rounded-lg bg-white opacity-0 group-hover:opacity-10 transition-opacity duration-200" />
                                        </button>
                                    )}

                                    {isOwner && isStarted && !isFinished && (
                                        <button
                                            onClick={handleForceEndEventClick}
                                            disabled={isForceEndingEvent}
                                            className="group relative inline-flex items-center justify-center gap-2 px-4 py-2 rounded-lg font-semibold text-white bg-gradient-to-r from-orange-500 to-red-600 hover:from-orange-600 hover:to-red-700 active:scale-95 transition-all duration-200 shadow-lg hover:shadow-xl disabled:opacity-50 disabled:cursor-not-allowed"
                                        >
                                            {isForceEndingEvent ? (
                                                <LoadingSpinner size="sm" />
                                            ) : (
                                                <>
                                                    <X className="h-4 w-4" />
                                                    <span className="text-sm">Force End</span>
                                                </>
                                            )}
                                            <div className="absolute inset-0 rounded-lg bg-white opacity-0 group-hover:opacity-10 transition-opacity duration-200" />
                                        </button>
                                    )}

                                    {isOwner && (
                                        <>
                                            <button
                                                onClick={() => setDeleteConfirmation({ open: true, confirmText: "", error: undefined })}
                                                disabled={isDeletingEvent}
                                                className="group relative inline-flex items-center justify-center gap-2 px-4 py-2 rounded-lg font-semibold text-white bg-gradient-to-r from-red-500 to-rose-600 hover:from-red-600 hover:to-rose-700 active:scale-95 transition-all duration-200 shadow-lg hover:shadow-xl disabled:opacity-50 disabled:cursor-not-allowed"
                                            >
                                                {isDeletingEvent ? (
                                                    <LoadingSpinner size="sm" />
                                                ) : (
                                                    <>
                                                        <Trash2 className="h-4 w-4" />
                                                        <span className="text-sm">Delete</span>
                                                    </>
                                                )}
                                                <div className="absolute inset-0 rounded-lg bg-white opacity-0 group-hover:opacity-10 transition-opacity duration-200" />
                                            </button>
                                        </>
                                    )}
                                </div>
                            </div>

                            <div className="flex flex-col sm:flex-row gap-4 text-sm">
                                {isStarted && event.startTime && (
                                    <div className="flex items-center gap-2 text-muted-foreground">
                                        <Clock className="h-4 w-4 flex-shrink-0" />
                                        <span>Started {new Date(event.startTime).toLocaleDateString()}</span>
                                    </div>
                                )}
                                {isStarted && event.startTime && (
                                    <div className="flex items-center gap-2 text-muted-foreground">
                                        <Calendar className="h-4 w-4 flex-shrink-0" />
                                        <span>
                      {isActive ? "Ends " : "Ended "}
                                            {new Date(new Date(event.startTime).getTime() + event.duration).toLocaleDateString()}
                    </span>
                                    </div>
                                )}
                            </div>
                        </div>

                        {event.question && (
                            <div className="rounded-lg border border-primary/20 bg-primary/5 p-6">
                                <div className="flex gap-4">
                                    <MessageCircle className="h-5 w-5 text-primary flex-shrink-0 mt-0.5" />
                                    <div className="flex-1 min-w-0">
                                        <p className="text-xs font-semibold text-muted-foreground mb-2 uppercase tracking-wide">
                                            The Question
                                        </p>
                                        <p className="text-base font-medium text-foreground leading-relaxed">{event.question}</p>
                                    </div>
                                </div>
                            </div>
                        )}

                        {isStarted && (
                            <div className="rounded-lg border border-border bg-card p-6">
                                <h2 className="text-lg font-semibold text-foreground mb-4">Cast Your Vote</h2>
                                <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
                                    {event.voteOptions.map((option, index) => (
                                        <button
                                            key={index}
                                            onClick={() => setVoteConfirmation({ open: true, selectedOption: index, confirmText: "" })}
                                            disabled={isSubmittingVote}
                                            className="p-4 rounded-lg border-2 border-border bg-background hover:border-primary hover:bg-muted transition-all active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed text-left font-medium text-foreground text-sm"
                                        >
                                            {isSubmittingVote ? <LoadingSpinner size="sm" text={option} /> : option}
                                        </button>
                                    ))}
                                </div>
                            </div>
                        )}
                    </div>

                    <div className="lg:col-span-1">
                        <div className="rounded-lg border border-border bg-card p-6 sticky top-6">
                            <div className="flex items-center justify-between gap-3 mb-5">
                                <div className="flex items-center gap-2">
                                    <Users className="h-5 w-5 text-primary" />
                                    <h2 className="font-semibold text-foreground">Members</h2>
                                    <Badge variant="secondary" className="ml-auto text-xs">
                                        {event.voteMembers.length}
                                    </Badge>
                                </div>
                            </div>

                            {isOwner && !isStarted && (
                                <Dialog open={isAddMemberDialogOpen} onOpenChange={setIsAddMemberDialogOpen}>
                                    <DialogTrigger asChild>
                                        <Button
                                            size="sm"
                                            variant="outline"
                                            className="w-full mb-4 bg-gradient-to-r from-blue-50 to-indigo-50 dark:from-blue-950/30 dark:to-indigo-950/30 hover:from-blue-100 hover:to-indigo-100 dark:hover:from-blue-900/50 dark:hover:to-indigo-900/50 border-blue-200 dark:border-blue-800"
                                        >
                                            <Plus className="h-4 w-4 mr-2" />
                                            Add Member
                                        </Button>
                                    </DialogTrigger>
                                    <DialogContent className="sm:max-w-md">
                                        <DialogHeader>
                                            <div className="flex items-center gap-3">
                                                <div className="p-2 bg-blue-100 dark:bg-blue-950 rounded-lg">
                                                    <Plus className="h-5 w-5 text-blue-600 dark:text-blue-400" />
                                                </div>
                                                <DialogTitle>Add Member</DialogTitle>
                                            </div>
                                        </DialogHeader>
                                        <div className="space-y-4">
                                            <p className="text-sm text-muted-foreground">
                                                Invite a member to participate in this event by entering their user ID.
                                            </p>
                                            <div>
                                                <label className="text-xs font-semibold text-muted-foreground mb-2 block uppercase tracking-wide">
                                                    User ID
                                                </label>
                                                <Input
                                                    placeholder="Enter user ID"
                                                    value={addedUserCommitment}
                                                    onChange={(e) => {
                                                        setAddedUserCommitment(e.target.value)
                                                        setAddMemberError(null)
                                                    }}
                                                    onKeyUp={(e) => {
                                                        if (e.key === "Enter" && addedUserCommitment.trim()) {
                                                            handleAddMember()
                                                        }
                                                    }}
                                                    disabled={isAddingMember}
                                                    className="font-mono text-sm"
                                                />
                                            </div>
                                            {addMemberError && (
                                                <div className="rounded-lg bg-red-50 dark:bg-red-950/20 border border-red-200 dark:border-red-800 p-3">
                                                    <p className="text-xs text-red-800 dark:text-red-200">
                                                        <span className="font-semibold">Error:</span> {addMemberError}
                                                    </p>
                                                </div>
                                            )}
                                            <Button
                                                onClick={handleAddMember}
                                                disabled={!addedUserCommitment.trim() || isAddingMember}
                                                className="w-full bg-gradient-to-r from-blue-500 to-indigo-600 hover:from-blue-600 hover:to-indigo-700 text-white flex items-center justify-center disabled:opacity-100 disabled:bg-gradient-to-r disabled:from-blue-500 disabled:to-indigo-600"
                                            >
                                                {isAddingMember ? <LoadingSpinner size="sm" /> : "Add Member"}
                                            </Button>
                                        </div>
                                    </DialogContent>
                                </Dialog>
                            )}

                            <div className="space-y-2 max-h-96 overflow-y-auto">
                                {event.voteMembers?.map((member) => (
                                    <div
                                        key={member}
                                        className="flex items-center justify-between p-3 rounded-md bg-muted/50 hover:bg-muted transition-colors"
                                    >
                                        <span className="font-medium text-sm truncate">{member}</span>
                                        <div className="flex items-center gap-2">
                                            <CheckCircle2 className="h-4 w-4 text-primary flex-shrink-0" />
                                            {isOwner && !isStarted && (
                                                <button
                                                    onClick={() => handleDeleteMember(member)}
                                                    disabled={isDeletingMember === member}
                                                    className="p-1 hover:bg-destructive/20 rounded transition-colors disabled:opacity-50"
                                                    title="Delete member"
                                                >
                                                    {isDeletingMember === member ? (
                                                        <LoadingSpinner size="sm" />
                                                    ) : (
                                                        <Trash2 className="h-3.5 w-3.5 text-destructive" />
                                                    )}
                                                </button>
                                            )}
                                        </div>
                                    </div>
                                ))}
                            </div>

                            {eventMembers.length > 0 && (
                                <div className="border-t border-border mt-4 pt-4">
                                    <p className="text-xs text-muted-foreground mb-3 font-semibold uppercase tracking-wide">
                                        Recently Added
                                    </p>
                                    <div className="space-y-2">
                                        {eventMembers?.map((member) => (
                                            <div
                                                key={member}
                                                className="flex items-center justify-between p-3 rounded-md bg-primary/5 hover:bg-primary/10 transition-colors"
                                            >
                                                <span className="font-medium text-sm truncate text-primary">{member}</span>
                                                <div className="flex items-center gap-2">
                                                    <Badge variant="secondary" className="text-xs flex-shrink-0">
                                                        New
                                                    </Badge>
                                                    {isOwner && !isStarted && (
                                                        <button
                                                            onClick={() => handleDeleteMember(member)}
                                                            disabled={isDeletingMember === member}
                                                            className="p-1 hover:bg-destructive/20 rounded transition-colors disabled:opacity-50"
                                                            title="Delete member"
                                                        >
                                                            {isDeletingMember === member ? (
                                                                <LoadingSpinner size="sm" />
                                                            ) : (
                                                                <Trash2 className="h-3.5 w-3.5 text-destructive" />
                                                            )}
                                                        </button>
                                                    )}
                                                </div>
                                            </div>
                                        ))}
                                    </div>
                                </div>
                            )}
                        </div>
                    </div>
                </div>
            </main>

            <Dialog open={isResultsDialogOpen} onOpenChange={setIsResultsDialogOpen}>
                <DialogContent className="sm:max-w-2xl">
                    <DialogHeader>
                        <div className="flex items-center gap-3">
                            <div className="p-2 bg-blue-100 dark:bg-blue-950 rounded-lg">
                                <BarChart3 className="h-5 w-5 text-blue-600 dark:text-blue-400" />
                            </div>
                            <DialogTitle>Vote Results</DialogTitle>
                        </div>
                    </DialogHeader>
                    {isLoadingResults ? (
                        <div className="py-12 flex flex-col items-center justify-center gap-4">
                            <LoadingSpinner size="lg" text="Loading results..." />
                        </div>
                    ) : voteResults ? (
                        <div className="space-y-6">
                            <div className="rounded-xl bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-blue-950/30 dark:to-indigo-950/30 border border-blue-200 dark:border-blue-800 p-6">
                                <p className="text-sm font-medium text-muted-foreground mb-2">Total Votes</p>
                                <p className="text-4xl font-bold bg-gradient-to-r from-blue-600 to-indigo-600 bg-clip-text text-transparent">
                                    {voteResults.reduce((sum, r) => sum + r.votes, 0)}
                                </p>
                            </div>

                            <div className="h-80 flex items-center justify-center">
                                <ChartContainer
                                    config={{
                                        votes: {
                                            label: "Votes",
                                        },
                                    }}
                                    className="w-full h-full"
                                >
                                    <ResponsiveContainer width="100%" height="100%">
                                        <PieChart>
                                            <Pie
                                                data={voteResults}
                                                cx="50%"
                                                cy="50%"
                                                labelLine={false}
                                                label={({ option, percentage }) => `${option}: ${percentage}%`}
                                                outerRadius={120}
                                                fill="#8884d8"
                                                dataKey="votes"
                                            >
                                                {voteResults.map((entry, index) => (
                                                    <Cell key={`cell-${index}`} fill={entry.color} />
                                                ))}
                                            </Pie>
                                            <ChartTooltip
                                                content={({ active, payload }) => {
                                                    if (active && payload && payload.length) {
                                                        const data = payload[0].payload as VoteResults
                                                        return (
                                                            <div className="rounded-lg border bg-background p-3 shadow-lg">
                                                                <div className="flex items-center gap-2 mb-1">
                                                                    <div className="w-3 h-3 rounded-full" style={{ backgroundColor: data.color }} />
                                                                    <span className="font-semibold text-sm">{data.option}</span>
                                                                </div>
                                                                <p className="text-sm text-muted-foreground">
                                                                    {data.votes} votes ({data.percentage}%)
                                                                </p>
                                                            </div>
                                                        )
                                                    }
                                                    return null
                                                }}
                                            />
                                        </PieChart>
                                    </ResponsiveContainer>
                                </ChartContainer>
                            </div>

                            <div className="space-y-4">
                                <p className="text-sm font-semibold text-muted-foreground uppercase tracking-wide">Breakdown</p>
                                {voteResults.map((result, index) => (
                                    <div key={index} className="space-y-2 p-4 rounded-lg bg-muted/30 hover:bg-muted/50 transition-colors">
                                        <div className="flex items-center justify-between">
                                            <div className="flex items-center gap-3">
                                                <div className="w-4 h-4 rounded-full flex-shrink-0" style={{ backgroundColor: result.color }} />
                                                <span className="text-base font-semibold text-foreground">{result.option}</span>
                                            </div>
                                            <div className="flex items-center gap-4">
                                                <span className="text-sm text-muted-foreground">{result.votes} votes</span>
                                                <span className="text-lg font-bold text-foreground min-w-[4rem] text-right">
                          {result.percentage}%
                        </span>
                                            </div>
                                        </div>
                                        <div className="h-3 bg-muted rounded-full overflow-hidden">
                                            <div
                                                className="h-full rounded-full transition-all duration-500"
                                                style={{
                                                    width: `${result.percentage}%`,
                                                    backgroundColor: result.color,
                                                }}
                                            />
                                        </div>
                                    </div>
                                ))}
                            </div>

                            <Button
                                variant="outline"
                                className="w-full bg-transparent hover:bg-muted"
                                onClick={() => setIsResultsDialogOpen(false)}
                            >
                                Close
                            </Button>
                        </div>
                    ) : (
                        <div className="py-12 text-center">
                            <p className="text-muted-foreground">No results available</p>
                        </div>
                    )}
                </DialogContent>
            </Dialog>

            <Dialog
                open={voteConfirmation.open}
                onOpenChange={(open) => {
                    if (!open) {
                        setVoteConfirmation({ open: false, selectedOption: null, confirmText: "", error: undefined })
                    }
                }}
            >
                <DialogContent>
                    <DialogHeader>
                        <DialogTitle>Confirm Your Vote</DialogTitle>
                    </DialogHeader>
                    <div className="space-y-4">
                        <div className="rounded-lg bg-muted p-4">
                            <p className="text-xs text-muted-foreground mb-2 uppercase tracking-wide font-semibold">
                                Selected Option
                            </p>
                            <p className="font-semibold text-lg text-foreground">
                                {voteConfirmation.selectedOption !== null && event?.voteOptions[voteConfirmation.selectedOption]}
                            </p>
                        </div>
                        {voteConfirmation.error && (
                            <div className="rounded-lg bg-red-50 dark:bg-red-950/20 border border-red-200 dark:border-red-800 p-3">
                                <p className="text-xs text-red-800 dark:text-red-200">
                                    <span className="font-semibold">Error:</span> {voteConfirmation.error}
                                </p>
                            </div>
                        )}
                        <div>
                            <p className="text-sm text-muted-foreground mb-3">
                                Type <span className="font-semibold text-foreground">"confirm"</span> to proceed:
                            </p>
                            <Input
                                placeholder="confirm"
                                value={voteConfirmation.confirmText}
                                onChange={(e) => {
                                    setVoteConfirmation((prev) => ({ ...prev, confirmText: e.target.value }))
                                    setVoteConfirmation((prev) => ({ ...prev, error: undefined }))
                                }}
                                onKeyUp={(e) => {
                                    if (e.key === "Enter" && voteConfirmation.confirmText.toLowerCase() === "confirm") {
                                        submitVoteAfterConfirmation()
                                    }
                                }}
                                disabled={isSubmittingVote}
                                className="font-mono"
                            />
                        </div>
                        <div className="flex gap-3 pt-2">
                            <Button
                                variant="outline"
                                onClick={() =>
                                    setVoteConfirmation({ open: false, selectedOption: null, confirmText: "", error: undefined })
                                }
                                disabled={isSubmittingVote}
                                className="flex-1"
                            >
                                Cancel
                            </Button>
                            <Button
                                onClick={submitVoteAfterConfirmation}
                                disabled={voteConfirmation.confirmText.toLowerCase() !== "confirm" || isSubmittingVote}
                                className="flex-1"
                            >
                                {isSubmittingVote ? <LoadingSpinner size="sm" text="Submitting..." /> : "Submit Vote"}
                            </Button>
                        </div>
                    </div>
                </DialogContent>
            </Dialog>

            <Dialog
                open={startConfirmation.open}
                onOpenChange={(open) => {
                    if (!open) {
                        setStartConfirmation({ open: false, confirmText: "", error: undefined })
                    }
                }}
            >
                <DialogContent className="sm:max-w-md">
                    <DialogHeader>
                        <div className="flex items-center gap-3">
                            <div className="p-2 bg-green-100 dark:bg-green-950 rounded-lg">
                                <Play className="h-5 w-5 text-green-600 dark:text-green-400" />
                            </div>
                            <DialogTitle>Start Event?</DialogTitle>
                        </div>
                    </DialogHeader>
                    <div className="space-y-4">
                        <p className="text-sm text-muted-foreground">
                            This will start the event <span className="font-semibold text-foreground">"{event?.name}"</span> and
                            enable voting for all members.
                        </p>
                        <div className="bg-blue-50 dark:bg-blue-950/20 border border-blue-200 dark:border-blue-800 rounded-lg p-3">
                            <p className="text-xs text-blue-800 dark:text-blue-200">
                                <span className="font-semibold">Info:</span> Members will be able to cast their votes once the event
                                starts.
                            </p>
                        </div>
                        {startConfirmation.error && (
                            <div className="rounded-lg bg-red-50 dark:bg-red-950/20 border border-red-200 dark:border-red-800 p-3">
                                <p className="text-xs text-red-800 dark:text-red-200">
                                    <span className="font-semibold">Error:</span> {startConfirmation.error}
                                </p>
                            </div>
                        )}
                        <div>
                            <p className="text-sm text-muted-foreground mb-3">
                                Type <span className="font-semibold text-foreground">"start"</span> to proceed:
                            </p>
                            <Input
                                placeholder="start"
                                value={startConfirmation.confirmText}
                                onChange={(e) => {
                                    setStartConfirmation((prev) => ({ ...prev, confirmText: e.target.value }))
                                    setStartConfirmation((prev) => ({ ...prev, error: undefined }))
                                }}
                                onKeyUp={(e) => {
                                    if (e.key === "Enter" && startConfirmation.confirmText.toLowerCase() === "start") {
                                        submitStartEventAfterConfirmation()
                                    }
                                }}
                                disabled={isStartingEvent}
                                className="font-mono"
                            />
                        </div>
                    </div>
                    <div className="flex gap-3 pt-4">
                        <Button
                            variant="outline"
                            className="flex-1 bg-transparent"
                            onClick={() => setStartConfirmation({ open: false, confirmText: "", error: undefined })}
                            disabled={isStartingEvent}
                        >
                            Cancel
                        </Button>
                        <Button
                            className="flex-1 bg-gradient-to-r from-green-500 to-emerald-600 hover:from-green-600 hover:to-emerald-700 text-white"
                            onClick={submitStartEventAfterConfirmation}
                            disabled={startConfirmation.confirmText.toLowerCase() !== "start" || isStartingEvent}
                        >
                            {isStartingEvent ? <LoadingSpinner size="sm" text="Starting..." /> : "Start Event"}
                        </Button>
                    </div>
                </DialogContent>
            </Dialog>

            <Dialog
                open={deleteConfirmation.open}
                onOpenChange={(open) => {
                    if (!open) {
                        setDeleteConfirmation({ open: false, confirmText: "", error: undefined })
                    }
                }}
            >
                <DialogContent className="sm:max-w-md">
                    <DialogHeader>
                        <div className="flex items-center gap-3">
                            <div className="p-2 bg-red-100 dark:bg-red-950 rounded-lg">
                                <AlertTriangle className="h-5 w-5 text-red-600 dark:text-red-400" />
                            </div>
                            <DialogTitle>Delete Event?</DialogTitle>
                        </div>
                    </DialogHeader>
                    <div className="space-y-4">
                        <p className="text-sm text-muted-foreground">
                            This action cannot be undone. The event{" "}
                            <span className="font-semibold text-foreground">"{event.name}"</span> and all associated data will be
                            permanently deleted.
                        </p>
                        <div className="bg-yellow-50 dark:bg-yellow-950/20 border border-yellow-200 dark:border-yellow-800 rounded-lg p-3">
                            <p className="text-xs text-yellow-800 dark:text-yellow-200">
                                <span className="font-semibold">Warning:</span> This cannot be reversed.
                            </p>
                        </div>
                        {deleteConfirmation.error && (
                            <div className="rounded-lg bg-red-50 dark:bg-red-950/20 border border-red-200 dark:border-red-800 p-3">
                                <p className="text-xs text-red-800 dark:text-red-200">
                                    <span className="font-semibold">Error:</span> {deleteConfirmation.error}
                                </p>
                            </div>
                        )}
                        <div>
                            <p className="text-sm text-muted-foreground mb-3">
                                Type <span className="font-semibold text-foreground">"delete"</span> to proceed:
                            </p>
                            <Input
                                placeholder="delete"
                                value={deleteConfirmation.confirmText}
                                onChange={(e) => {
                                    setDeleteConfirmation((prev) => ({ ...prev, confirmText: e.target.value }))
                                    setDeleteConfirmation((prev) => ({ ...prev, error: undefined }))
                                }}
                                onKeyUp={(e) => {
                                    if (e.key === "Enter" && deleteConfirmation.confirmText.toLowerCase() === "delete") {
                                        submitDeleteEventAfterConfirmation()
                                    }
                                }}
                                disabled={isDeletingEvent}
                                className="font-mono"
                            />
                        </div>
                    </div>
                    <div className="flex gap-3 pt-4">
                        <Button
                            variant="outline"
                            className="flex-1 bg-transparent"
                            onClick={() => setDeleteConfirmation({ open: false, confirmText: "", error: undefined })}
                            disabled={isDeletingEvent}
                        >
                            Cancel
                        </Button>
                        <Button
                            className="flex-1 bg-gradient-to-r from-red-500 to-rose-600 hover:from-red-600 hover:to-rose-700 text-white"
                            onClick={submitDeleteEventAfterConfirmation}
                            disabled={deleteConfirmation.confirmText.toLowerCase() !== "delete" || isDeletingEvent}
                        >
                            {isDeletingEvent ? <LoadingSpinner size="sm" text="Deleting..." /> : "Delete Event"}
                        </Button>
                    </div>
                </DialogContent>
            </Dialog>

            <Dialog
                open={forceEndConfirmation.open}
                onOpenChange={(open) => {
                    if (!open) {
                        setForceEndConfirmation({ open: false, confirmText: "", error: undefined })
                    }
                }}
            >
                <DialogContent className="sm:max-w-md">
                    <DialogHeader>
                        <div className="flex items-center gap-3">
                            <div className="p-2 bg-orange-100 dark:bg-orange-950 rounded-lg">
                                <X className="h-5 w-5 text-orange-600 dark:text-orange-400" />
                            </div>
                            <DialogTitle>Force End Event?</DialogTitle>
                        </div>
                    </DialogHeader>
                    <div className="space-y-4">
                        <p className="text-sm text-muted-foreground">
                            This will immediately end the event <span className="font-semibold text-foreground">"{event?.name}"</span> before
                            its scheduled end time.
                        </p>
                        <div className="bg-yellow-50 dark:bg-yellow-950/20 border border-yellow-200 dark:border-yellow-800 rounded-lg p-3">
                            <p className="text-xs text-yellow-800 dark:text-yellow-200">
                                <span className="font-semibold">Warning:</span> Members will no longer be able to cast votes after the event is
                                force ended.
                            </p>
                        </div>
                        {forceEndConfirmation.error && (
                            <div className="rounded-lg bg-red-50 dark:bg-red-950/20 border border-red-200 dark:border-red-800 p-3">
                                <p className="text-xs text-red-800 dark:text-red-200">
                                    <span className="font-semibold">Error:</span> {forceEndConfirmation.error}
                                </p>
                            </div>
                        )}
                        <div>
                            <p className="text-sm text-muted-foreground mb-3">
                                Type <span className="font-semibold text-foreground">"end"</span> to proceed:
                            </p>
                            <Input
                                placeholder="end"
                                value={forceEndConfirmation.confirmText}
                                onChange={(e) => {
                                    setForceEndConfirmation((prev) => ({ ...prev, confirmText: e.target.value }))
                                    setForceEndConfirmation((prev) => ({ ...prev, error: undefined }))
                                }}
                                onKeyUp={(e) => {
                                    if (e.key === "Enter" && forceEndConfirmation.confirmText.toLowerCase() === "end") {
                                        submitForceEndEventAfterConfirmation()
                                    }
                                }}
                                disabled={isForceEndingEvent}
                                className="font-mono"
                            />
                        </div>
                    </div>
                    <div className="flex gap-3 pt-4">
                        <Button
                            variant="outline"
                            className="flex-1 bg-transparent"
                            onClick={() => setForceEndConfirmation({ open: false, confirmText: "", error: undefined })}
                            disabled={isForceEndingEvent}
                        >
                            Cancel
                        </Button>
                        <Button
                            className="flex-1 bg-gradient-to-r from-orange-500 to-red-600 hover:from-orange-600 hover:to-red-700 text-white"
                            onClick={submitForceEndEventAfterConfirmation}
                            disabled={forceEndConfirmation.confirmText.toLowerCase() !== "end" || isForceEndingEvent}
                        >
                            {isForceEndingEvent ? <LoadingSpinner size="sm" text="Ending..." /> : "Force End Event"}
                        </Button>
                    </div>
                </DialogContent>
            </Dialog>
        </div>
    )
}
