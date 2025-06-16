import { NextRequest, NextResponse } from 'next/server'

export const runtime = 'edge' // หรือ 'nodejs' ตามที่ต้องการ (ถ้าใช้ฟีเจอร์ Node.js api)

export async function GET(req: NextRequest, { params }: { params: { path: string[] } }) {
  return handleProxy(req, params)
}
export async function POST(req: NextRequest, { params }: { params: { path: string[] } }) {
  return handleProxy(req, params)
}
export async function PUT(req: NextRequest, { params }: { params: { path: string[] } }) {
  return handleProxy(req, params)
}
export async function DELETE(req: NextRequest, { params }: { params: { path: string[] } }) {
  return handleProxy(req, params)
}

async function handleProxy(req: NextRequest, params: { path: string[] }) {
  const urlPath = params.path.join('/')
  const backendUrl = `http://localhost:8080/api/${urlPath}`

  const headers = new Headers(req.headers)
  const cookie = req.cookies.get('jwt_token')
  const csrf_token = req.cookies.get('csrf_token')
  if (cookie) {
    headers.set('Authorization', `Bearer ${cookie.value}`)
  }

  if (csrf_token) {
    headers.set('X-CRSF-Token', csrf_token.value)
  }

  const fetchOptions: RequestInit = {
    method: req.method,
    headers,
    body: ['POST', 'PUT', 'PATCH'].includes(req.method) ? await req.text() : undefined,
  }

  const response = await fetch(backendUrl, fetchOptions)

  const data = await response.arrayBuffer()
  const resHeaders = new Headers(response.headers)

  resHeaders.delete('content-encoding')

  return new NextResponse(Buffer.from(data), {
    status: response.status,
    headers: resHeaders,
  })
}
