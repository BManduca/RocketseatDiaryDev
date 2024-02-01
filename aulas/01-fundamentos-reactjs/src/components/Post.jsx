import { Avatar } from './Avatar'
import { Comment } from './Comment'
import styles from './Post.module.css'

export function Post(props) {
    return (
        <article className={styles.post}>
            <header>
                <div className={styles.author}>
                    <Avatar src={props.author.avatarUrl} />
                    <div className={styles.authorInfo}>
                        <strong> {props.author.name} </strong>
                        <span> {props.author.role} </span>
                    </div>
                </div>

                <time title='29 de Janeiro às 14:15h' dateTime='2024-01-29 14:15:35'>Publicado há 1h.</time>
            </header>

            <div className={styles.content}>
            </div>

            <form className={styles.commentForm}>
                <strong>Deixe seu feedback:</strong>
                <textarea  placeholder='Deixe um comentário: ' />

                <footer>
                    <button type='submit'> Publicar </button>
                </footer>
            </form>

            <div className={styles.commentList}>
                <Comment />
                <Comment />
                <Comment />
            </div>
        </article>
    )
}