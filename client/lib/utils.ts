import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const BASE_URL = "http://127.0.0.1:8080/api/v1"

export class ApiError extends Error {
  constructor(
    message: string,
    public status: number,
    public statusText: string,
    public response?: any
  ) {
    super(message)
    this.name = 'ApiError'
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
      return errorData.error || errorData.message || errorData.details || 'Request failed'
    } else {
      const textError = await response.text()
      return textError || `HTTP ${response.status}: ${response.statusText}`
    }
  } catch {
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

    clearTimeout(timeoutId)

    if (!response.ok) {
      const errorMessage = await extractErrorMessage(response)
      throw new ApiError(
        errorMessage,
        response.status,
        response.statusText,
        response
      )
    }

    const contentType = response.headers.get('content-type')
    if (!contentType?.includes('application/json')) {
      return {} as T
    }

    return await response.json()

  } catch (error) {
    clearTimeout(timeoutId)

    if (error instanceof ApiError) {
      throw error
    }

    if (error.name === 'AbortError') {
      throw new ApiError('Request timeout', 408, 'Request Timeout')
    }

    if (error instanceof TypeError && error.message.includes('fetch')) {
      throw new ApiError('Network error - please check your connection', 0, 'Network Error')
    }

    throw new ApiError(
      error.message || 'Unknown error occurred',
      0,
      'Unknown Error'
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
  let lastError: Error

  for (let i = 0; i <= maxRetries; i++) {
    try {
      return await apiCall()
    } catch (error) {
      lastError = error as Error

      if (i === maxRetries) break

      // Jangan retry untuk client errors (4xx)
      if (error instanceof ApiError && error.status >= 400 && error.status < 500) {
        break
      }

      // Exponential backoff
      await new Promise(resolve => setTimeout(resolve, delay * Math.pow(2, i)))
    }
  }

  throw lastError
}
