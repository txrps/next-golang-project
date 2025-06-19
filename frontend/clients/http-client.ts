/* eslint-disable @typescript-eslint/ban-ts-comment */
/* eslint-disable @typescript-eslint/no-explicit-any */
import { getCookie } from 'cookies-next'
import https from 'https'

class HttpClient {
  private readonly baseURL: string
  private readonly httpsAgent: https.Agent

  constructor() {
    this.baseURL = '/api'
    this.httpsAgent = new https.Agent({ rejectUnauthorized: false }) // Only used server-side
  }

  async request<T>(
    input: string,
    options: RequestInit = {},
  ): Promise<SuccessResponse<T>> {
    const url = `${this.baseURL}${input}`

    const rawHeaders: Record<string, string> = {
      'Content-Type': 'application/json',
      'redirect_proxy': 'true',
    }

    const token = getCookie("jwt_token")
    if (token) {
      rawHeaders['Authorization'] = `Bearer ${token}`

      if (
        options.method &&
        ['POST', 'PUT', 'PATCH', 'DELETE'].includes(options.method.toUpperCase())
      ) {
        const body = options.body ? JSON.parse(options.body.toString()) : {}
        options.body = JSON.stringify({
          ...body,
        })
      }
    }

    const headers: HeadersInit = rawHeaders

    const response = await fetch(url, {
      ...options,
      headers,
      // @ts-expect-error
      agent: typeof window === 'undefined' ? this.httpsAgent : undefined, // only server-side
    })

    if (!response.ok) {
      const error = await response.json().catch(() => ({}))
      throw new Error(error.message)
    }

    return response.json()
  }

  get<T>(url: string) {
    return this.request<T>(url, { method: 'GET' })
  }

  post<T>(url: string, data?: any) {
    return this.request<T>(url, {
      method: 'POST',
      body: JSON.stringify(data ?? {}),
    })
  }

  put<T>(url: string, data?: any) {
    return this.request<T>(url, {
      method: 'PUT',
      body: JSON.stringify(data ?? {}),
    })
  }

  delete<T>(url: string, data?: any) {
    return this.request<T>(url, {
      method: 'DELETE',
      body: JSON.stringify(data ?? {}),
    })
  }
}

export const httpClient = new HttpClient()

export type SuccessResponse<T> = {
  code: number
  data: T
}
