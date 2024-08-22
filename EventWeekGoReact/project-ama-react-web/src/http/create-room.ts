interface CreateRoomRequest {
  theme: string // tema da sala
}

export async function createRoom({ theme }: CreateRoomRequest) {
  const response = await fetch('http://localhost:8080/api/rooms', {
    method: 'POST',
    body: JSON.stringify({
      theme,
    })
  })

  const data: { id: string } = await response.json()

  // so acontecerá o retorno caso a request não falhar
  return { roomID: data.id }
}