interface CreateRoomRequest {
  theme: string // tema da sala
}

export async function createRoom({ theme }: CreateRoomRequest) {
  const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/rooms`, {
    method: 'POST',
    body: JSON.stringify({
      theme,
    })
  })

  const data: { id: string } = await response.json()

  // so acontecerá o retorno caso a request não falhar
  return { roomID: data.id }
}