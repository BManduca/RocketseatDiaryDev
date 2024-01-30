import { ThumbsUp, Trash } from 'phosphor-react';
import styles from './Comment.module.css';

export function Comment() {
    return (
        <div className={styles.comment}>
            <img src="https://github.com/BManduca.png" />

            <div className={styles.commentBox}>
                <div className={styles.commentContent}>
                    <header>
                        <div className={styles.authorAndTime}>
                            <strong>Brunno Manduca</strong>
                            <time title='29 de Janeiro às 14:15h' dateTime='2024-01-29 14:15:35'>Cerca de 1h atrás.</time>
                        </div>
                        <button title='Deletar comentário'>
                            <Trash size={20} />
                        </button>
                    </header>

                    <p>
                        Muito bom Devon, parabéns!! 👏👏
                    </p>
                </div>

                <footer>
                    <button>
                        <ThumbsUp />
                        Aplaudir <span>20</span>
                    </button>
                </footer>
            </div>


        </div>
    )
}