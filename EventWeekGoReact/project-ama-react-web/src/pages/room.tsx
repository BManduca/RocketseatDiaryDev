import { useParams } from "react-router-dom"

export function Room() {

    const roomID = useParams()

    console.log(roomID)

    return (
        <h1>Room</h1>
    )
}