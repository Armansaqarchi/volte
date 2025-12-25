"use client"

import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
    DialogFooter,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { useState } from "react"

interface ConfirmProps {
    buttonValue: string
    message: string
    onConfirm: () => void
    onCancel: () => void
    icon?: React.ReactNode // optional icon support
}

export default function ConfirmButton({
                                          buttonValue,
                                          message,
                                          onConfirm,
                                          onCancel,
                                          icon,
                                      }: ConfirmProps) {
    const [open, setOpen] = useState(false)

    const handleConfirm = () => {
        setOpen(false)
        onConfirm()
    }

    const handleCancel = () => {
        setOpen(false)
        onCancel()
    }

    return (
        <>
            <Button
                onClick={() => setOpen(true)}
                className="items-center gap-2"
            >
                {buttonValue}
            </Button>

            <Dialog open={open} onOpenChange={setOpen}>
                <DialogContent className="rounded-lg">
                    <DialogHeader>
                        <DialogTitle className="text-lg font-semibold">
                            {message}
                        </DialogTitle>
                    </DialogHeader>

                    <DialogFooter className="flex justify-end gap-3">
                        <Button variant="outline" onClick={handleCancel}>
                            Cancel
                        </Button>
                        <Button onClick={handleConfirm}>
                            Confirm
                        </Button>
                    </DialogFooter>
                </DialogContent>
            </Dialog>
        </>
    )
}
