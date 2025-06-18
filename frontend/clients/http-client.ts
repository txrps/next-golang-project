/* eslint-disable @typescript-eslint/no-explicit-any */
import { getCookie } from 'cookies-next'
import { jwtDecode } from 'jwt-decode'

const BASE_URL = "http://localhost:3000/"

export const decodeKeyCloakToken = (ctx?: any): string | null => {
  const token = getCookie("jwt_token", ctx)

  if (!token) {
    return null
  }

  const tokenString = Array.isArray(token) ? token[0] : token

  if (typeof tokenString !== "string") {
    return null
  }

  const decoded = jwtDecode<{ preferred_username?: string }>(tokenString)
  return decoded.preferred_username ?? null
}


type RequestMethod = 'GET' | 'POST' | 'PUT' | 'DELETE' | 'PATCH'

type RequestOptions = {
  method?: RequestMethod
  headers?: HeadersInit
  body?: any
}

class HttpClient {
  private static _instance: HttpClient

  public static get instance(): HttpClient {
    if (!this._instance) {
      this._instance = new HttpClient()
    }
    return this._instance
  }

  private buildHeaders(method: RequestMethod, body?: any): HeadersInit {
    const headers: HeadersInit = {
      'Content-Type': 'application/json',
      'redirect_proxy': 'true',
    }

    const token = String(getCookie("jwt_token") || '')
    if (token) {
      headers['Authorization'] = `Bearer ${token}`

      if (['POST', 'PUT', 'PATCH', 'DELETE'].includes(method)) {
        const userId = decodeKeyCloakToken(token)
        if (body && typeof body === 'object') {
          body.keyCloakUserId = userId
        }
      }
    }

    return headers
  }

  async request<T>(url: string, options: RequestOptions = {}): Promise<SuccessResponse<T>> {
    const method = (options.method ?? 'GET').toUpperCase() as RequestMethod
    const body = options.body ? JSON.stringify(options.body) : undefined
    const headers = this.buildHeaders(method, options.body)

    const response = await fetch(BASE_URL + url, {
      method,
      headers,
      body,
    })

    const json = await response.json()
    if (!response.ok) {
      throw new Error(json?.message ?? 'API error')
    }

    return json
  }

  get<T>(url: string) {
    return this.request<T>(url, { method: 'GET' })
  }

  post<T>(url: string, body?: any) {
    return this.request<T>(url, { method: 'POST', body })
  }

  put<T>(url: string, body?: any) {
    return this.request<T>(url, { method: 'PUT', body })
  }

  delete<T>(url: string, body?: any) {
    return this.request<T>(url, { method: 'DELETE', body })
  }

  patch<T>(url: string, body?: any) {
    return this.request<T>(url, { method: 'PATCH', body })
  }
}

export const httpClient = HttpClient.instance

export type SuccessResponse<T> = {
  code: number
  data: T
}
