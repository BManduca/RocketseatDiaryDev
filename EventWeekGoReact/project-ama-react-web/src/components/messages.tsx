import { useParams } from "react-router-dom";
import { Message } from "./message";
import { getRoomMessages } from "../http/get-room-messages";
import { useSuspenseQuery } from "@tanstack/react-query";

export function Messages() {
  const { roomID } = useParams()

  if (!roomID) {
    throw new Error('Messages components must be used within room page!')
  }

  const { data } = useSuspenseQuery({
    queryKey: ['messages', roomID],
    queryFn: () => getRoomMessages({ roomID })
  })

  console.log(data)

  return (
    <ol className="list-decimal list-outside px-3 space-y-8">
        {data.messages.map((message, index) => {
          return (
            <Message
              key={index}
              id={message.id}
              text={message.text}
              amountOfReactions={message.amountOfReactions}
              answered={message.answered}
            />
          )
        })}
    </ol>
  )
}


