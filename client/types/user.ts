export type UserRole = "admin" | "dokter" | "paramedis" | "kasir" | "pasien"

export interface User {
  id: string
  email: string
  name?: string
  role: UserRole
  // Add other fields as needed
} 