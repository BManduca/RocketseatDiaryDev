import { useState } from 'react'
import { Avatar } from './Avatar'
import { Comment } from './Comment'
import styles from './Post.module.css'

import { format, formatDistanceToNow } from 'date-fns'
import { ptBR } from 'date-fns/locale/pt-BR'

export function Post({ author, publishedAt, content }) {

    const [comments, setComments] = useState([
        'Post muito bacana, hein?'
    ])

    //armazenando o texto digitado no input
    const [newCommentText, setNewCommentText] = useState('')

    const publishedDateFormatted = format(publishedAt, "d 'de' LLLL 'às' HH:mm'h'", {
        locale: ptBR
    })

    const publishedDateRelativeToNow = formatDistanceToNow(publishedAt, {
        locale: ptBR,
        addSuffix: true
    })

    //função para disparar criação de novo comentário com base na ação do usuário
    function handleCreateNewComment() {
        event.preventDefault()

        // adicionando novo comentário
        setComments([...comments, newCommentText]);
        
        //estado que armazena o conteúdo da nossa textarea
        setNewCommentText('')
    }

    function handleNewCommentChange() {
        setNewCommentText(event.target.value)
    }

    function deleteComment(commentToDelete) {

        const commentsWithoutDeleteOne = comments.filter(comment => {

            // const commentsWithoutDeleteOne = comments.filter(comment => comment.id !== commentId)

            return comment !== commentToDelete
        })

        // atualizando a lista de comentários, removendo o que foi deletado
        setComments(commentsWithoutDeleteOne);
    }

    return (
        <article className={styles.post}>
            <header>
                <div className={styles.author}>
                    <Avatar src={author.avatarUrl} />
                    <div className={styles.authorInfo}>
                        <strong> {author.name} </strong>
                        <span> {author.role} </span>
                    </div>
                </div>

                <time title={publishedDateFormatted} dateTime={publishedAt.toISOString()}>
                    {publishedDateRelativeToNow}
                </time>
            </header>

            <div className={styles.content}>
                {content.map(line => {
                    if (line.type === 'paragraph') {
                        return <p key={line.content}>{line.content}</p>
                    }else if (line.type === 'link') {
                        return <p key={line.content}> <a href="#">{line.content}</a> </p>
                    }
                })}
            </div>

            <div className={styles.content}>
            </div>

            <form onSubmit={handleCreateNewComment} className={styles.commentForm}>
                <strong>Deixe seu feedback:</strong>
                <textarea 
                    name="comment" 
                    placeholder='Deixe um comentário: '
                    value={newCommentText}
                    onChange={handleNewCommentChange}
                />

                <footer>
                    <button type='submit'> Publicar </button>
                </footer>
            </form>

            <div className={styles.commentList}>
                { comments.map(comment => {
                    return (
                        <Comment 
                            key={comment} 
                            content={comment} 
                            onDeleteComment={deleteComment} 
                        />
                    )
                }) }
            </div>
        </article>
    )
}