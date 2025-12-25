'use client'

export function LoadingOverlay({ isVisible = false, message = 'Loading...' }: { isVisible?: boolean; message?: string }) {
    if (!isVisible) return null

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center">
            <div className="absolute inset-0 bg-black/50 backdrop-blur-sm animate-in fade-in duration-300" />

            <div className="relative flex flex-col items-center gap-4 animate-in zoom-in-50 fade-in duration-500">
                <div className="relative h-12 w-12">
                    <div className="absolute inset-0 rounded-full border-4 border-transparent border-t-primary border-r-primary border-r-opacity-50 animate-spin" />
                    {/* Inner fading ring for color fade effect */}
                    <div className="absolute inset-2 rounded-full border-3 border-transparent border-b-primary/30 opacity-75 animate-spin [animation-direction:reverse]" />
                </div>

                {message && (
                    <p className="text-sm font-medium text-foreground/80 animate-pulse">
                        {message}
                    </p>
                )}
            </div>
        </div>
    )
}
