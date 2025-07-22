export interface Question {
  id: string
  label: string
  type: "text" | "date" | "select" | "checkbox" | "checkbox_textarea"
  options?: string[]
}

// You may need to adjust this type if you use z.infer elsewhere
export type FormData = {
  nik: string
  name: string
  email: string
  age: number
  gender?: string
  contact?: string
  place_of_birth?: string
  date_of_birth: Date | null
  address?: string
  rt_rw?: string
  village?: string
  district?: string
  religion?: string
  marital_status?: string
  occupation?: string
  nationality?: string
  valid_until?: string
  blood_type?: string
  ktp_images?: any
  tinggi_badan?: string
  berat_badan?: string
  [key: string]: any
} 