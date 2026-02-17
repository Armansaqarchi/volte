"use client"

import { useEffect, useRef, useCallback } from "react"

interface Node {
  x: number
  y: number
  vx: number
  vy: number
  radius: number
  opacity: number
  pulsePhase: number
  pulseSpeed: number
  connected: boolean
}

interface FloatingSymbol {
  x: number
  y: number
  vx: number
  vy: number
  opacity: number
  maxOpacity: number
  fadeSpeed: number
  symbol: string
  size: number
  rotation: number
  rotationSpeed: number
}

export function ZKBackground() {
  const canvasRef = useRef<HTMLCanvasElement>(null)
  const animationRef = useRef<number>(0)
  const nodesRef = useRef<Node[]>([])
  const symbolsRef = useRef<FloatingSymbol[]>([])
  const mouseRef = useRef<{ x: number; y: number }>({ x: -1000, y: -1000 })
  const timeRef = useRef(0)

  const ZK_SYMBOLS = [
    "\u03C0", "\u03A3", "\u03B4", "\u2200", "\u2203",
    "\u2227", "\u2228", "\u22A2", "\u2261", "\u2260",
    "0", "1", "\u03BB", "\u03B1", "\u03B2",
    "\u2295", "\u2297", "\u221E",
  ]

  const initNodes = useCallback((width: number, height: number) => {
    const count = Math.min(Math.floor((width * height) / 18000), 80)
    const nodes: Node[] = []
    for (let i = 0; i < count; i++) {
      nodes.push({
        x: Math.random() * width,
        y: Math.random() * height,
        vx: (Math.random() - 0.5) * 0.3,
        vy: (Math.random() - 0.5) * 0.3,
        radius: Math.random() * 1.5 + 0.5,
        opacity: Math.random() * 0.5 + 0.1,
        pulsePhase: Math.random() * Math.PI * 2,
        pulseSpeed: Math.random() * 0.01 + 0.005,
        connected: Math.random() > 0.3,
      })
    }
    nodesRef.current = nodes
  }, [])

  const initSymbols = useCallback((width: number, height: number) => {
    const count = Math.min(Math.floor((width * height) / 60000), 15)
    const symbols: FloatingSymbol[] = []
    for (let i = 0; i < count; i++) {
      symbols.push({
        x: Math.random() * width,
        y: Math.random() * height,
        vx: (Math.random() - 0.5) * 0.15,
        vy: (Math.random() - 0.5) * 0.15,
        opacity: 0,
        maxOpacity: Math.random() * 0.08 + 0.03,
        fadeSpeed: Math.random() * 0.003 + 0.001,
        symbol: ZK_SYMBOLS[Math.floor(Math.random() * ZK_SYMBOLS.length)],
        size: Math.random() * 14 + 10,
        rotation: Math.random() * Math.PI * 2,
        rotationSpeed: (Math.random() - 0.5) * 0.005,
      })
    }
    symbolsRef.current = symbols
  }, [])

  useEffect(() => {
    const canvas = canvasRef.current
    if (!canvas) return

    const ctx = canvas.getContext("2d", { alpha: true })
    if (!ctx) return

    const handleResize = () => {
      const dpr = window.devicePixelRatio || 1
      canvas.width = window.innerWidth * dpr
      canvas.height = window.innerHeight * dpr
      canvas.style.width = `${window.innerWidth}px`
      canvas.style.height = `${window.innerHeight}px`
      ctx.scale(dpr, dpr)
      initNodes(window.innerWidth, window.innerHeight)
      initSymbols(window.innerWidth, window.innerHeight)
    }

    const handleMouseMove = (e: MouseEvent) => {
      mouseRef.current = { x: e.clientX, y: e.clientY }
    }

    handleResize()
    window.addEventListener("resize", handleResize)
    window.addEventListener("mousemove", handleMouseMove)

    const CONNECTION_DISTANCE = 150
    const MOUSE_RADIUS = 200

    const animate = () => {
      const width = window.innerWidth
      const height = window.innerHeight
      timeRef.current += 1

      ctx.clearRect(0, 0, width, height)

      // Update and draw nodes
      const nodes = nodesRef.current
      for (const node of nodes) {
        node.x += node.vx
        node.y += node.vy

        if (node.x < 0 || node.x > width) node.vx *= -1
        if (node.y < 0 || node.y > height) node.vy *= -1

        node.x = Math.max(0, Math.min(width, node.x))
        node.y = Math.max(0, Math.min(height, node.y))

        // Pulse
        const pulse = Math.sin(timeRef.current * node.pulseSpeed + node.pulsePhase) * 0.3 + 0.7
        const currentOpacity = node.opacity * pulse

        // Mouse interaction - subtle glow near cursor
        const dx = node.x - mouseRef.current.x
        const dy = node.y - mouseRef.current.y
        const distToMouse = Math.sqrt(dx * dx + dy * dy)
        const mouseInfluence = distToMouse < MOUSE_RADIUS ? (1 - distToMouse / MOUSE_RADIUS) * 0.4 : 0

        // Draw node
        ctx.beginPath()
        ctx.arc(node.x, node.y, node.radius + mouseInfluence * 2, 0, Math.PI * 2)
        ctx.fillStyle = `rgba(0, 210, 190, ${currentOpacity + mouseInfluence})`
        ctx.fill()

        // Subtle glow
        if (mouseInfluence > 0.1) {
          ctx.beginPath()
          ctx.arc(node.x, node.y, node.radius + 6, 0, Math.PI * 2)
          ctx.fillStyle = `rgba(0, 210, 190, ${mouseInfluence * 0.15})`
          ctx.fill()
        }
      }

      // Draw connections
      for (let i = 0; i < nodes.length; i++) {
        if (!nodes[i].connected) continue
        for (let j = i + 1; j < nodes.length; j++) {
          if (!nodes[j].connected) continue
          const dx = nodes[i].x - nodes[j].x
          const dy = nodes[i].y - nodes[j].y
          const dist = Math.sqrt(dx * dx + dy * dy)

          if (dist < CONNECTION_DISTANCE) {
            const alpha = (1 - dist / CONNECTION_DISTANCE) * 0.08
            ctx.beginPath()
            ctx.moveTo(nodes[i].x, nodes[i].y)
            ctx.lineTo(nodes[j].x, nodes[j].y)
            ctx.strokeStyle = `rgba(0, 210, 190, ${alpha})`
            ctx.lineWidth = 0.5
            ctx.stroke()
          }
        }
      }

      // Draw and update floating symbols
      const symbols = symbolsRef.current
      for (const sym of symbols) {
        sym.x += sym.vx
        sym.y += sym.vy
        sym.rotation += sym.rotationSpeed

        if (sym.x < -50 || sym.x > width + 50) sym.vx *= -1
        if (sym.y < -50 || sym.y > height + 50) sym.vy *= -1

        // Fade in/out cycle
        sym.opacity += sym.fadeSpeed
        if (sym.opacity >= sym.maxOpacity) {
          sym.fadeSpeed = -Math.abs(sym.fadeSpeed)
        } else if (sym.opacity <= 0) {
          sym.fadeSpeed = Math.abs(sym.fadeSpeed)
          sym.symbol = ZK_SYMBOLS[Math.floor(Math.random() * ZK_SYMBOLS.length)]
          sym.x = Math.random() * width
          sym.y = Math.random() * height
        }

        ctx.save()
        ctx.translate(sym.x, sym.y)
        ctx.rotate(sym.rotation)
        ctx.font = `${sym.size}px "Geist Mono", monospace`
        ctx.fillStyle = `rgba(0, 210, 190, ${Math.max(0, sym.opacity)})`
        ctx.textAlign = "center"
        ctx.textBaseline = "middle"
        ctx.fillText(sym.symbol, 0, 0)
        ctx.restore()
      }

      // Subtle grid overlay (circuit-like)
      const gridSpacing = 80
      const gridOpacity = 0.015
      ctx.strokeStyle = `rgba(0, 210, 190, ${gridOpacity})`
      ctx.lineWidth = 0.5

      for (let x = 0; x < width; x += gridSpacing) {
        ctx.beginPath()
        ctx.moveTo(x, 0)
        ctx.lineTo(x, height)
        ctx.stroke()
      }
      for (let y = 0; y < height; y += gridSpacing) {
        ctx.beginPath()
        ctx.moveTo(0, y)
        ctx.lineTo(width, y)
        ctx.stroke()
      }

      animationRef.current = requestAnimationFrame(animate)
    }

    animate()

    return () => {
      cancelAnimationFrame(animationRef.current)
      window.removeEventListener("resize", handleResize)
      window.removeEventListener("mousemove", handleMouseMove)
    }
  }, [initNodes, initSymbols])

  return (
    <canvas
      ref={canvasRef}
      className="fixed inset-0 pointer-events-none"
      style={{ zIndex: 0 }}
      aria-hidden="true"
    />
  )
}
