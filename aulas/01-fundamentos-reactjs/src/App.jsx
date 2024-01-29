import { Post } from './components/Post';
import { Header } from './components/Header';
import { Sidebar } from './components/Sidebar'

import styles from './App.module.css';

import './global.css';

export function App() {
  return (
    <>
      <Header />
      <div className={styles.wrapper}>
        <Sidebar />
        <main>
          <Post 
            author="Brunno Manduca" 
            content="Lorem ipsum dolor sit amet consectetur adipisicing elit. Enim officia quo explicabo ab id porro obcaecati sint autem consectetur consequuntur unde tempora dolorem, iusto laboriosam minima harum voluptatem dolore. Aperiam."
          />
          <Post
            author="Gabriel"
            content="Um novo post muito bom!"
          />
        </main>
      </div>

      
    </>
  )
}
