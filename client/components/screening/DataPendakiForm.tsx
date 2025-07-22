import { FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { Alert, AlertDescription } from "@/components/ui/alert"
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs"
import { Upload, Loader2 } from "lucide-react"
import { PhoneInput } from "@/components/ui/phone-input"
import { Calendar as CalendarIcon } from "lucide-react"
import { cn } from "@/lib/utils"
import { Calendar } from "@/components/ui/calendar"
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover"
import { Badge } from "@/components/ui/badge"
import { WebcamCapture } from "./WebcamCapture"
import { FormData } from "@/types/screening"
import { RefObject } from "react"

interface DataPendakiFormProps {
  form: any
  entryMethod: string
  setEntryMethod: (v: string) => void
  imageFile: File | null
  isAnalyzing: boolean
  analysisError: string | null
  fileInputRef: RefObject<HTMLInputElement>
  handleFileChange: (e: React.ChangeEvent<HTMLInputElement>) => void
  isCameraActive: boolean
  setIsCameraActive: (v: boolean) => void
  handleCameraCapture: (img: string) => void
}

export function DataPendakiForm({
  form,
  entryMethod,
  setEntryMethod,
  imageFile,
  isAnalyzing,
  analysisError,
  fileInputRef,
  handleFileChange,
  isCameraActive,
  setIsCameraActive,
  handleCameraCapture,
}: DataPendakiFormProps) {
  return (
    <>
      <Alert variant="default" className="mb-4">
        <AlertDescription>
          <b>Perhatian:</b> Jika menggunakan AI (scan/upload KTP), mohon cek ulang data yang terisi otomatis. Data hasil AI kadang tidak sesuai, pastikan semua data sudah benar sebelum lanjut.
        </AlertDescription>
      </Alert>
      <Tabs value={entryMethod} onValueChange={setEntryMethod} className="mb-4">
        <TabsList className="grid grid-cols-3 w-full">
          <TabsTrigger value="manual">Input Manual</TabsTrigger>
          <TabsTrigger value="upload">Upload KTP</TabsTrigger>
          <TabsTrigger value="camera">Scan KTP</TabsTrigger>
        </TabsList>
        <TabsContent value="upload">
          <div className="space-y-4">
            <div className="flex justify-center items-center w-full">
              <label htmlFor="ktp-upload" className="flex flex-col justify-center items-center w-full h-48 bg-gray-50 rounded-lg border-2 border-dashed cursor-pointer">
                <Upload className="mb-2 w-8 h-8 text-gray-500" />
                <span className="text-sm text-gray-500">Klik untuk unggah atau seret foto KTP di sini</span>
                <input id="ktp-upload" type="file" accept="image/*" className="hidden" ref={fileInputRef} onChange={handleFileChange} />
              </label>
            </div>
            {imageFile && <div className="mt-2 text-sm text-gray-500">File terpilih: {imageFile.name}</div>}
            <Button onClick={() => fileInputRef.current?.click()} disabled={isAnalyzing} className="w-full">
              {isAnalyzing ? <><Loader2 className="mr-2 h-4 w-4 animate-spin" /> Menganalisis...</> : "Pilih Foto KTP"}
            </Button>
            {analysisError && <Alert variant="destructive"><AlertDescription>{analysisError}</AlertDescription></Alert>}
          </div>
        </TabsContent>
        <TabsContent value="camera">
          <div className="space-y-4">
            <Button type="button" onClick={() => setIsCameraActive(!isCameraActive)} className="w-full mb-2">
              {isCameraActive ? "Hentikan Kamera" : "Mulai Kamera"}
            </Button>
            {isCameraActive && <WebcamCapture onCapture={handleCameraCapture} isActive={isCameraActive} setIsActive={setIsCameraActive} />}
            {analysisError && <Alert variant="destructive"><AlertDescription>{analysisError}</AlertDescription></Alert>}
          </div>
        </TabsContent>
      </Tabs>
      {/* Grid 2 kolom, urutan dan label sesuai PatientInfoForm.jsx */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {/* NIK */}
        <FormField control={form.control} name="nik" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="nik">NIK</FormLabel>
            <FormControl>
              <Input id="nik" placeholder="NIK" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Nama Lengkap */}
        <FormField control={form.control} name="name" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="name">Nama Lengkap</FormLabel>
            <FormControl>
              <Input id="name" placeholder="Nama Lengkap" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Tempat Lahir */}
        <FormField control={form.control} name="place_of_birth" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="place_of_birth">Tempat Lahir</FormLabel>
            <FormControl>
              <Input id="place_of_birth" placeholder="Tempat Lahir" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Tanggal Lahir */}
        <FormField control={form.control} name="date_of_birth" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="date_of_birth">Tanggal Lahir</FormLabel>
            <Popover>
              <PopoverTrigger asChild>
                <FormControl>
                  <Button
                    variant={"outline"}
                    className={cn(
                      "w-full pl-3 text-left font-normal",
                      !field.value && "text-muted-foreground"
                    )}
                  >
                    {field.value ? format(field.value, "PPP") : <span>Pilih tanggal</span>}
                    <CalendarIcon className="ml-auto h-4 w-4 opacity-50" />
                  </Button>
                </FormControl>
              </PopoverTrigger>
              <PopoverContent className="w-auto p-0" align="start">
                <Calendar
                  mode="single"
                  selected={field.value ?? undefined}
                  onSelect={field.onChange}
                  disabled={(date) => date > new Date() || date < new Date("1900-01-01")}
                  captionLayout="dropdown"
                />
              </PopoverContent>
            </Popover>
            <FormMessage />
          </FormItem>
        )} />
        {/* Jenis Kelamin */}
        <FormField control={form.control} name="gender" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="gender">Jenis Kelamin</FormLabel>
            <FormControl>
              <Tabs value={field.value || ""} onValueChange={field.onChange}>
                <TabsList className="w-full">
                  <TabsTrigger value="laki-laki">Laki-Laki</TabsTrigger>
                  <TabsTrigger value="perempuan">Perempuan</TabsTrigger>
                  <TabsTrigger value="lainnya">Lainnya</TabsTrigger>
                </TabsList>
              </Tabs>
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Alamat */}
        <FormField control={form.control} name="address" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="address">Alamat</FormLabel>
            <FormControl>
              <Input id="address" placeholder="Alamat" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* RT/RW */}
        <FormField control={form.control} name="rt_rw" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="rt_rw">RT/RW</FormLabel>
            <FormControl>
              <Input id="rt_rw" placeholder="002/014" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Kel/Desa */}
        <FormField control={form.control} name="village" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="village">Kel/Desa</FormLabel>
            <FormControl>
              <Input id="village" placeholder="Kel/Desa" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Kecamatan */}
        <FormField control={form.control} name="district" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="district">Kecamatan</FormLabel>
            <FormControl>
              <Input id="district" placeholder="Kecamatan" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Agama */}
        <FormField control={form.control} name="religion" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="religion">Agama</FormLabel>
            <FormControl>
              <Input id="religion" placeholder="Agama" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Status Perkawinan */}
        <FormField control={form.control} name="marital_status" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="marital_status">Status Perkawinan</FormLabel>
            <FormControl>
              <Input id="marital_status" placeholder="Status Perkawinan" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Pekerjaan */}
        <FormField control={form.control} name="occupation" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="occupation">Pekerjaan</FormLabel>
            <FormControl>
              <Input id="occupation" placeholder="Pekerjaan" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Kewarganegaraan */}
        <FormField control={form.control} name="nationality" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="nationality">Kewarganegaraan</FormLabel>
            <FormControl>
              <Input id="nationality" placeholder="Kewarganegaraan" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Berlaku Hingga */}
        <FormField control={form.control} name="valid_until" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="valid_until">Berlaku Hingga</FormLabel>
            <FormControl>
              <Input id="valid_until" placeholder="Berlaku Hingga" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Golongan Darah */}
        <FormField control={form.control} name="blood_type" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="blood_type">Golongan Darah</FormLabel>
            <FormControl>
              <Input id="blood_type" placeholder="Contoh: A, B, AB, O (- jika tidak diketahui)" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Tinggi Badan */}
        <FormField control={form.control} name="tinggi_badan" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="tinggi_badan">Tinggi Badan (cm)</FormLabel>
            <FormControl>
              <Input id="tinggi_badan" placeholder="Tinggi Badan dalam cm" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Berat Badan */}
        <FormField control={form.control} name="berat_badan" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="berat_badan">Berat Badan (kg)</FormLabel>
            <FormControl>
              <Input id="berat_badan" placeholder="Berat Badan dalam kg" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Umur */}
        <FormField control={form.control} name="age" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="age">Umur</FormLabel>
            <FormControl>
              <Input id="age" placeholder="Umur" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Email */}
        <FormField control={form.control} name="email" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="email">Email</FormLabel>
            <FormControl>
              <Input id="email" type="email" placeholder="email@example.com" {...field} />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
        {/* Nomor Telepon */}
        <FormField control={form.control} name="contact" render={({ field }) => (
          <FormItem className="w-full">
            <FormLabel htmlFor="contact">Nomor Telepon</FormLabel>
            <FormControl>
              <PhoneInput
                id="contact"
                value={field.value || ""}
                onChange={field.onChange}
                defaultCountry="ID"
                international
                countryCallingCodeEditable={false}
                placeholder="08xxxxxxxxxx"
              />
            </FormControl>
            <FormMessage />
          </FormItem>
        )} />
      </div>
    </>
  )
} 