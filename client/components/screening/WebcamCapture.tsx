import React, { useRef, useState } from "react"
import ReactWebcam from "react-webcam"
import { Button } from "@/components/ui/button"
import { Camera, Repeat } from "lucide-react"

interface WebcamCaptureProps {
  onCapture: (img: string) => void
  isActive: boolean
  setIsActive: (v: boolean) => void
}

export function WebcamCapture({ onCapture, isActive, setIsActive }: WebcamCaptureProps) {
  const webcamRef = useRef<any>(null)
  const [facingMode, setFacingMode] = useState("environment")
  if (!isActive) return null
  return (
    <div className="relative">
      <ReactWebcam
        audio={false}
        ref={webcamRef}
        screenshotFormat="image/jpeg"
        className="w-full rounded-lg"
        videoConstraints={{ facingMode }}
      />
      <div className="absolute bottom-4 left-1/2 -translate-x-1/2 flex gap-2">
        <Button type="button" onClick={() => {
          const imageSrc = webcamRef.current?.getScreenshot()
          if (imageSrc) {
            onCapture(imageSrc)
            setIsActive(false)
          }
        }}>
          <Camera className="mr-2 h-4 w-4" /> Capture
        </Button>
        <Button type="button" variant="outline" onClick={() => setFacingMode(facingMode === "user" ? "environment" : "user") }>
          <Repeat className="mr-2 h-4 w-4" /> Ganti Kamera
        </Button>
      </div>
    </div>
  )
} 