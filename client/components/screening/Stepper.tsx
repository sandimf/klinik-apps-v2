import { CheckCircle2 } from "lucide-react"

interface StepperProps {
  step: number
}

export function Stepper({ step }: StepperProps) {
  const steps = [
    { label: "Data Pendaki", status: step === 1 ? "current" : "completed" },
    { label: "Screening", status: step === 2 ? "current" : step === 1 ? "pending" : "completed" },
  ]
  return (
    <div className="flex items-center justify-center gap-8 mb-8">
      {steps.map((s, i) => (
        <div key={i} className="flex flex-col items-center flex-1 min-w-0">
          <div className={`flex items-center justify-center w-7 h-7 rounded-full border-2 ${s.status === "completed" ? "bg-green-500 border-green-500 text-white" : s.status === "current" ? "border-blue-500 text-blue-500 bg-background" : "border-muted-foreground text-muted-foreground bg-muted"} transition-colors`}>
            {s.status === "completed" ? (
              <CheckCircle2 className="w-4 h-4" />
            ) : (
              <span className="font-bold text-sm">{i + 1}</span>
            )}
          </div>
          <div className="mt-1 text-xs font-semibold text-center" style={{ color: s.status === "current" ? "#2563eb" : undefined }}>{s.label}</div>
        </div>
      ))}
    </div>
  )
} 