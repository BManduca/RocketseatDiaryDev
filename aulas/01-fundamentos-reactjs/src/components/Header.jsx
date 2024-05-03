import styles from './Header.module.css'

import igniteLogo from '../assets/ignite-logo.svg';

export function Header() {
    return (
        <header  className={styles.header}>
            <img src={ igniteLogo } alt="Logotipo do ignite" />
            <h2 className={styles.titleHeader}>Ignite Feed</h2>
        </header>
    );
}