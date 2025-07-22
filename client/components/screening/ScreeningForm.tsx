"use client"
import { useEffect, useRef, useState } from "react"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import * as z from "zod"
import { toast } from "sonner"
import { apiGet } from "@/lib/utils"
import { LoadingSkeleton } from "./LoadingSkeleton"
import { Button } from "../ui/button"
import { Stepper } from "./Stepper"
import { DataPendakiForm } from "./DataPendakiForm"
import { ScreeningQuestions } from "./ScreeningQuestions"
import { ConfirmDialog } from "./ConfirmDialog"
import { FormData, Question } from "@/types/screening"
import { Form } from "@/components/ui/form"

const genderOptions = ["Laki-laki", "Perempuan"]
const religionOptions = ["Islam", "Kristen", "Katolik", "Hindu", "Buddha", "Konghucu"]
const maritalStatusOptions = ["Belum Menikah", "Menikah", "Cerai Hidup", "Cerai Mati"]
const bloodTypeOptions = ["A", "B", "AB", "O", "A+", "A-", "B+", "B-", "AB+", "AB-", "O+", "O-"]

const formSchema = z.object({
  nik: z.string().length(16, "NIK harus 16 digit").regex(/^\d+$/, "NIK harus berupa angka"),
  name: z.string().min(2, "Nama minimal 2 karakter"),
  email: z.string().email("Format email tidak valid"),
  age: z.number().min(1, "Usia minimal 1 tahun").max(120, "Usia maksimal 120 tahun"),
  gender: z.string().optional(),
  contact: z.string().optional(),
  place_of_birth: z.string().optional(),
  date_of_birth: z.date().nullable().refine(val => val !== null, { message: "Tanggal Lahir wajib diisi" }),
  address: z.string().optional(),
  rt_rw: z.string().optional(),
  village: z.string().optional(),
  district: z.string().optional(),
  religion: z.string().optional(),
  marital_status: z.string().optional(),
  occupation: z.string().optional(),
  nationality: z.string().optional(),
  valid_until: z.string().optional(),
  blood_type: z.string().optional(),
  ktp_images: z.any().optional(),
  tinggi_badan: z.string().optional(),
  berat_badan: z.string().optional(),
}).passthrough()

export function ScreeningForm() {
  const [questions, setQuestions] = useState<Question[] | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(true)
  const [isSubmitting, setIsSubmitting] = useState(false)
  const [step, setStep] = useState(1)
  const [entryMethod, setEntryMethod] = useState<'manual' | 'upload' | 'camera'>("manual")
  const [imageFile, setImageFile] = useState<File | null>(null)
  const [isAnalyzing, setIsAnalyzing] = useState(false)
  const [analysisError, setAnalysisError] = useState<string | null>(null)
  const [isCameraActive, setIsCameraActive] = useState(false)
  const fileInputRef = useRef<HTMLInputElement>(null)
  const [showConfirmDialog, setShowConfirmDialog] = useState(false)

  const form = useForm<FormData>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      nik: "",
      name: "",
      email: "",
      age: 0,
      gender: "",
      contact: "",
      place_of_birth: "",
      date_of_birth: null,
      address: "",
      rt_rw: "",
      village: "",
      district: "",
      religion: "",
      marital_status: "",
      occupation: "",
      nationality: "Indonesia",
      valid_until: "",
      blood_type: "",
      ktp_images: null,
      tinggi_badan: "",
      berat_badan: "",
    },
  })

  useEffect(() => {
    const fetchQuestions = async () => {
      try {
        setIsLoading(true)
        const data = await apiGet<Question[]>("/screening/questions")
        setQuestions(data)
        setError(null)
      } catch (err) {
        setError(err instanceof Error ? err.message : "Terjadi kesalahan saat memuat pertanyaan")
        console.error("Error fetching questions:", err)
      } finally {
        setIsLoading(false)
      }
    }
    fetchQuestions()
  }, [])

  useEffect(() => {
    const dob = form.watch("date_of_birth")
    if (dob instanceof Date && !isNaN(dob.getTime())) {
      const today = new Date()
      let age = today.getFullYear() - dob.getFullYear()
      const m = today.getMonth() - dob.getMonth()
      if (m < 0 || (m === 0 && today.getDate() < dob.getDate())) {
        age--
      }
      form.setValue("age", age)
    }
  }, [form.watch("date_of_birth")])

  async function onSubmit(values: FormData) {
    try {
      setIsSubmitting(true)
      toast.success("Data berhasil dikirim!", {
        description: "Terima kasih telah mengisi formulir screening.",
      })
      form.reset()
    } catch (error) {
      console.error("Submit error:", error)
      toast.error("Terjadi kesalahan saat mengirim data")
    } finally {
      setIsSubmitting(false)
    }
  }

  async function analyzeKtpImage(fileOrBase64: File | string): Promise<Partial<FormData>> {
    await new Promise(r => setTimeout(r, 2000))
    return {
      nik: "1234567890123456",
      name: "Nama AI",
      place_of_birth: "Kota AI",
      date_of_birth: new Date("1990-01-01"),
      gender: "Laki-laki",
      address: "Jalan AI No. 1",
      rt_rw: "001/002",
      village: "Desa AI",
      district: "Kecamatan AI",
      religion: "Islam",
      marital_status: "Belum Menikah",
      occupation: "Programmer",
      nationality: "Indonesia",
      valid_until: "2025-12-31",
      blood_type: "O",
    }
  }

  // AI analyze handler
  const handleAnalyzeImage = async (fileOrBase64: File | string) => {
    setIsAnalyzing(true)
    setAnalysisError(null)
    try {
      const aiData = await analyzeKtpImage(fileOrBase64)
      Object.entries(aiData).forEach(([key, value]) => {
        if (value) form.setValue(key as string, value as any)
      })
      toast.success("Data berhasil diisi otomatis dari KTP!")
    } catch (err: any) {
      setAnalysisError("Gagal menganalisis gambar KTP. Coba lagi.")
      toast.error("Gagal menganalisis gambar KTP.")
    } finally {
      setIsAnalyzing(false)
    }
  }

  const handleFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (file) {
      setImageFile(file)
      await handleAnalyzeImage(file)
    }
  }

  const handleCameraCapture = async (imgBase64: string) => {
    setImageFile(null)
    await handleAnalyzeImage(imgBase64)
  }

  if (error) {
    return (
      <div className="container mx-auto p-6 max-w-2xl">
        <div className="text-red-600 font-semibold">{error}</div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-background">
      <div className="container mx-auto p-6 max-w-4xl">
        <div className="text-center mb-8">
          <h1 className="text-4xl font-bold tracking-tight">Formulir Screening</h1>
          <p className="text-muted-foreground mt-2">
            Mohon lengkapi data berikut dengan benar dan sesuai dokumen resmi
          </p>
        </div>
        <Stepper step={step} />
        {isLoading ? (
          <LoadingSkeleton />
        ) : questions && questions.length === 0 ? (
          <div className="text-center text-muted-foreground">
            Belum ada data pertanyaan screening tersedia.
          </div>) : (
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
              {step === 1 && (
                <DataPendakiForm
                  form={form}
                  entryMethod={entryMethod}
                  setEntryMethod={(v: string) => setEntryMethod(v as 'manual' | 'upload' | 'camera')}
                  imageFile={imageFile}
                  isAnalyzing={isAnalyzing}
                  analysisError={analysisError}
                  fileInputRef={fileInputRef as React.RefObject<HTMLInputElement>}
                  handleFileChange={handleFileChange}
                  isCameraActive={isCameraActive}
                  setIsCameraActive={setIsCameraActive}
                  handleCameraCapture={handleCameraCapture}
                />
              )}
              <ScreeningQuestions questions={questions || []} form={form} step={step} />
              {step === 1 && (
                <div className="flex justify-end mt-6">
                  <Button type="button" className="w-full h-12 text-base font-semibold" onClick={() => {
                    // Validasi field wajib
                    let missing: string[] = [];
                    const requiredFields: string[] = ["nik", "name", "gender", "date_of_birth", "age"];
                    const values = { ...form.getValues() };
                    for (const f of requiredFields) {
                      if (!values[f]) {
                        missing.push(f)
                      }
                    }
                    if (missing.length > 0) {
                      toast.error("Silakan lengkapi data pasien terlebih dahulu.")
                      return
                    }
                    setShowConfirmDialog(true)
                  }}>
                    Lanjut ke Screening
                  </Button>
                </div>
              )}
              <ConfirmDialog
                open={showConfirmDialog}
                onOpenChange={setShowConfirmDialog}
                onConfirm={() => {
                  setShowConfirmDialog(false)
                  setStep(2)
                  window.scrollTo({ top: 0, behavior: "smooth" })
                }}
                onCancel={() => setShowConfirmDialog(false)}
              />
              {step === 2 && (
                <div className="flex flex-col items-center space-y-4">
                  <Button
                    type="submit"
                    size="lg"
                    className="w-full max-w-md"
                    disabled={isSubmitting}
                  >
                    {isSubmitting ? "Mengirim Data..." : "Kirim Formulir"}
                  </Button>
                  <p className="text-sm text-muted-foreground text-center">
                    Dengan mengirim formulir ini, Anda menyetujui bahwa data yang diberikan adalah benar dan dapat dipertanggungjawabkan.
                  </p>
                </div>
              )}
            </form>
          </Form>
        )}
      </div>
    </div>
  )
}
