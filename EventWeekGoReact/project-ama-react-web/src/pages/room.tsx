import { useParams } from "react-router-dom"
import amaLogo from '../assets/ama-logo.svg'
import { ArrowRight, Share2 } from 'lucide-react'
import { toast } from "sonner"
import { Message } from "../components/message"

export function Room() {

    const { roomID } = useParams()

    function handleShareRoom() {
        const url = window.location.href.toString()

        if (navigator.share != undefined && navigator.canShare()) {
            navigator.share({ url })
        } else {
            // não puder compartilhar por falta de compartibilidade ou permissão
            navigator.clipboard.writeText(url)

            toast.info('O link da sala foi copiado com sucesso para a área de transferência!')
        }

    }

    return (
        <div className="mx-auto max-w-[640px] flex flex-col gap-6 py-10 px-4">
            <div className="flex items-center gap-3 px-3">
                <img src={amaLogo} alt="Logo da aplicação AMA" className="h-5"/>
                <span className="text-sm text-zinc-500 truncate">
                    Código da sala: <span className="text-zinc-300">{ roomID }</span>
                </span>

                <button
                    type="submit"
                    onClick={handleShareRoom}
                    className="ml-auto bg-zinc-800 text-zinc-300 px-3 py-1.5 gap-1.5 flex items-center rounded-lg font-medium text-sm transition-colors hover:bg-zinc-700"
                >
                    Compartilhar
                    <Share2 className="size-4"/>
                </button>
            </div>

            <div className="h-px w-full bg-zinc-900"/>

            <form
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

            <ol className="list-decimal list-outside px-3 space-y-8">
                <Message text="O que é GoLang e quais são suas principais vantagens em comparação com outras linguagens de programação como Python, Java ou C++?" amountOfReactions={100} answered />
                <Message text="Como funcionam as goroutines em GoLang e por que elas são importantes para a concorrência e paralelismo?" amountOfReactions={50} />
                <Message text="Quais são as melhores práticas para organizar o código em um projeto GoLang, incluindo pacotes, módulos e a estrutura de diretórios?" amountOfReactions={10} />
            </ol>
        </div>
    )
}