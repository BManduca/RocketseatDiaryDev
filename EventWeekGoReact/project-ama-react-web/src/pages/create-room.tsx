import { ArrowRight } from 'lucide-react'

import amaLogo from '../assets/ama-logo.svg'

export function CreateRoom() {
    return (
        <main className='h-screen flex items-center justify-center px-4'>
            <div className='max-w-[450px] flex flex-col gap-6'>
                <img src={amaLogo} alt="Logo AMA application" className='h-10' />

                <p className='leading-relaxed text-zinc-300 text-center'>
                    Crie uma sala pública de AMA (Ask me anything) e priorize as perguntas mais importantes para a comunidade.
                </p>

                <form className='flex items-center gap-2 bg-zinc-900 p-2 rounded-xl border border-zinc-800'>
                    <input
                        type="text"
                        name='theme'
                        placeholder='Nome da sala'
                    />

                    <button type='submit' className='bg-orange-400 text-orange-950 px-3 py-1.5 gap-1.5 flex items-center rounded-lg font-medium text-sm'>
                        Criar sala <ArrowRight className='size-4' /> 
                        {/* setando o tamanho do ícone como 1rem tanto de width e height*/}
                    </button>
                </form>
            </div>
        </main>
    )
}