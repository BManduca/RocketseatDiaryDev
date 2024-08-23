import { ArrowRight } from "lucide-react";
import { useParams } from "react-router-dom";
import { createMessage } from "../http/create-message";
import { toast } from "sonner";

export function CreateMessageForm() {

  const { roomID } = useParams()

  if (!roomID) {
    throw new Error('Messages components must be used within room page!')
  }

  async function createdMessageAction(data: FormData) {
    const message = data.get('message')?.toString()

    if (!message || !roomID) {
      return
    }

    try {
      await createMessage({ message, roomID })

    } catch {
      toast.error('Falha ao enviar pergunta, tente novamente!')
    }
  }

  return (
    <form
        action={createdMessageAction}
        className='flex items-center gap-2 bg-zinc-900 p-2 rounded-xl border border-zinc-800 ring-orange-400 ring-offset-2 ring-offset-zinc-950 focus-within:ring-1'
    >
        <input
            type="text"
            name='message'
            placeholder='Qual a sua pergunta?'
            autoComplete='off'
            className='flex-1 text-sm bg-transparent mx-2 outline-none text-zinc-100 placeholder:text-zinc-500'
            required
        />

        <button
            type='submit'
            className='bg-orange-400 text-orange-950 px-3 py-1.5 gap-1.5 flex items-center rounded-lg font-medium text-sm transition-colors hover:bg-orange-500'
        >
            Criar pergunta <ArrowRight className='size-4' /> 
            {/* setando o tamanho do ícone como 1rem(16px) tanto de width e height*/}
        </button>
    </form>
  )
}