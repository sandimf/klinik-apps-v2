import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

// Gunakan variabel lingkungan untuk BASE_URL di Next.js
// Pastikan NEXT_PUBLIC_API_BASE_URL didefinisikan di .env
export const BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || "http://127.0.0.1:8080/api/v1"

export class ApiError extends Error {
  constructor(
    message: string,
    public status: number,
    public statusText: string,
    public response?: any // Pertimbangkan untuk menggunakan Response | null
  ) {
    super(message)
    this.name = 'ApiError'
    // Mengatur prototype secara eksplisit untuk kompatibilitas
    // terutama saat menggunakan transpiler seperti Babel
    Object.setPrototypeOf(this, ApiError.prototype);
  }
}

interface ApiErrorResponse {
  error?: string
  message?: string
  details?: string
  [key: string]: any
}

async function extractErrorMessage(response: Response): Promise<string> {
  const contentType = response.headers.get('content-type')

  try {
    if (contentType?.includes('application/json')) {
      const errorData: ApiErrorResponse = await response.json()
      return errorData.error || errorData.message || errorData.details || `Request failed with status ${response.status}`
    } else {
      const textError = await response.text()
      return textError || `HTTP ${response.status}: ${response.statusText}`
    }
  } catch (parseError) {
    console.error("Failed to parse error response body:", parseError); // Log parsing error
    return `HTTP ${response.status}: ${response.statusText}`
  }
}

// Generic API request function
async function apiRequest<T>(
  url: string,
  options: RequestInit = {},
  timeout = 10000
): Promise<T> {
  const fullUrl = url.startsWith("http") ? url : BASE_URL + url

  // Setup abort controller untuk timeout
  const controller = new AbortController()
  const timeoutId = setTimeout(() => controller.abort(), timeout)

  try {
    const defaultHeaders: HeadersInit = {
      'Content-Type': 'application/json',
    }

    const response = await fetch(fullUrl, {
      signal: controller.signal,
      ...options,
      headers: {
        ...defaultHeaders,
        ...options.headers,
      },
    })

    clearTimeout(timeoutId) // Clear timeout segera setelah response diterima

    // Tangani respons non-OK (status code 4xx atau 5xx)
    if (!response.ok) {
      const errorMessage = await extractErrorMessage(response)
      throw new ApiError(
        errorMessage,
        response.status,
        response.statusText,
        response
      )
    }

    // Tangani respons OK (status code 2xx)
    const contentType = response.headers.get('content-type')
    if (response.status === 204) { // 204 No Content
      return undefined as T; // Return undefined atau null jika tidak ada konten yang diharapkan
    }
    if (!contentType?.includes('application/json')) {
      // Jika respons sukses tapi bukan JSON, dan bukan 204, Anda bisa throw error
      // atau mengembalikan objek kosong/undefined tergantung ekspektasi
      console.warn(`API returned successful non-JSON response for ${fullUrl}. Content-Type: ${contentType}`);
      // Pilihan: throw new ApiError(`Unexpected non-JSON response`, response.status, response.statusText, response);
      return {} as T; // Default ke objek kosong jika tidak ada JSON
    }

    return await response.json()

  } catch (error) {
    clearTimeout(timeoutId) // Pastikan timeout dihapus meskipun ada error sebelum fetch selesai

    if (error instanceof ApiError) {
      throw error // Re-throw ApiError yang sudah ditangkap dan dikustomisasi
    }

    if (typeof error === 'object' && error !== null && 'name' in error && (error as any).name === 'AbortError') {
      throw new ApiError('Request timeout', 408, 'Request Timeout', null) // Tambahkan null untuk response
    }

    // Periksa apakah ini Type Error yang berasal dari masalah jaringan
    // Error jaringan biasanya akan menghasilkan TypeError dengan message tertentu
    if (error instanceof TypeError && (error.message.includes('Failed to fetch') || error.message.includes('NetworkError'))) {
        throw new ApiError('Network error - please check your connection', 0, 'Network Error', null);
    }


    // Fallback untuk error yang tidak teridentifikasi
    // Pastikan `error` di-cast ke `Error` sebelum mengakses `.message`
    const unknownError = error as Error;
    throw new ApiError(
      unknownError.message || 'Unknown error occurred',
      0,
      'Unknown Error',
      null
    )
  }
}

export async function apiPost<T>(
  url: string,
  data: any,
  options: Omit<RequestInit, 'method' | 'body'> = {}
): Promise<T> {
  return apiRequest<T>(url, {
    method: 'POST',
    body: JSON.stringify(data),
    ...options,
  })
}

export async function apiGet<T>(
  url: string,
  options: Omit<RequestInit, 'method'> = {}
): Promise<T> {
  // Jika T bisa undefined/null untuk kasus 204 No Content, sesuaikan di pemanggil
  // Contoh: apiGet<MyDataType | undefined>('/data')
  return apiRequest<T>(url, {
    method: 'GET',
    ...options,
  })
}

export async function apiPut<T>(
  url: string,
  data: any,
  options: Omit<RequestInit, 'method' | 'body'> = {}
): Promise<T> {
  return apiRequest<T>(url, {
    method: 'PUT',
    body: JSON.stringify(data),
    ...options,
  })
}

export async function apiDelete<T>(
  url: string,
  options: Omit<RequestInit, 'method'> = {}
): Promise<T> {
  return apiRequest<T>(url, {
    method: 'DELETE',
    ...options,
  })
}

export async function apiPatch<T>(
  url: string,
  data: any,
  options: Omit<RequestInit, 'method' | 'body'> = {}
): Promise<T> {
  return apiRequest<T>(url, {
    method: 'PATCH',
    body: JSON.stringify(data),
    ...options,
  })
}

export async function withLoading<T>(
  apiCall: () => Promise<T>,
  setLoading?: (loading: boolean) => void
): Promise<T> {
  try {
    setLoading?.(true)
    return await apiCall()
  } finally {
    setLoading?.(false)
  }
}

// Hook-like function untuk retry logic
export async function withRetry<T>(
  apiCall: () => Promise<T>,
  maxRetries = 3,
  delay = 1000
): Promise<T> {
  // Inisialisasi lastError dengan Error default
  let lastError: Error = new Error("Max retries exceeded and no successful response.");

  for (let i = 0; i <= maxRetries; i++) {
    try {
      return await apiCall()
    } catch (error: unknown) { // Tangkap sebagai unknown
      lastError = error as Error; // Type assertion

      if (i === maxRetries) break; // Jangan retry jika sudah mencapai batas

      // Jangan retry untuk client errors (4xx) atau error yang bukan ApiError
      if (error instanceof ApiError && error.status >= 400 && error.status < 500) {
        break;
      }
      // Jika error bukan ApiError (misal TypeError jaringan), tetap bisa di-retry
      // jika tidak masuk ke kondisi di atas.

      // Exponential backoff
      await new Promise(resolve => setTimeout(resolve, delay * Math.pow(2, i)))
    }
  }

  throw lastError; // Akan selalu bertipe Error karena inisialisasi dan type assertion
}