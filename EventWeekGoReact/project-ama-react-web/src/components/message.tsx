import { ArrowUp } from "lucide-react";
import { useState } from "react";
import { useParams } from "react-router-dom";
import { createMessageReaction } from "../http/create-message-reaction";
import { toast } from "sonner";
import { removeMessageReaction } from "../http/remove-message-reaction";

interface MessageProps {
  id: string
  text: string
  amountOfReactions: number
  answered?: boolean
}

export function Message({
  id: messageID,
  text,
  amountOfReactions,
  answered = false
}: MessageProps) {
  const { roomID } = useParams()
  const [hasReacted, setHasReacted] = useState(false)

  if (!roomID) {
    throw new Error('Messages components must be used within room page!')
  }

  async function createMessageReactionAction() {
    if (!roomID) {
      return
    }

    try {
      await createMessageReaction({ messageID, roomID })
    } catch {
      toast.error('Falha ao tentar reagir a mensagem, tente novamente!')
    }

    setHasReacted(true)

  }

  async function removeMessageReactionAction() {
    if (!roomID) {
      return
    }

    try {
      await removeMessageReaction({ messageID, roomID })
    } catch {
      toast.error('Falha ao remover a reaçao de curtida, tente novamente!')
    }

    setHasReacted(false)
  }

  return (
    <li  data-answered={answered} className="ml-4 leading-relaxed text-zinc-100 data-[answered=true]:opacity-50 data-[answered=true]:pointer-events-none">
        {text}

        {hasReacted ? (
          <button
            type="button"
            onClick={removeMessageReactionAction}
            className="mt-3 flex items-center gap-2 text-orange-400 text-sm font-medium hover:text-orange-500"
          >
              <ArrowUp className="size-4"/>
              Curtir pergunta ({amountOfReactions})
          </button>
        ) : (
          <button
            type="button"
            onClick={createMessageReactionAction}
            className="mt-3 flex items-center gap-2 text-zinc-400 text-sm font-medium hover:text-zinc-300"
          >
              <ArrowUp className="size-4"/>
              Curtir pergunta ({amountOfReactions})
          </button>
        )}
    </li>
  )
}