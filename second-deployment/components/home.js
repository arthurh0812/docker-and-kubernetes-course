import Link from "next/link";
import styles from "../styles/Home.module.css";

export default function Home() {
  return (
    <main className={styles.main}>
      <h1 className={styles.title}>Welcome to Tradaccs!</h1>

      <p className={styles.description}>
        Get started by editing{" "}
        <code className={styles.code}>pages/index.js</code>
      </p>

      <div className={styles.grid}>
        <Link href="/documentation">
          <a className={styles.card}>
            <h3>Documentation &rarr;</h3>
            <p>
              Find in-depth information about Tradaccs Platform and Why it is so
              Secure.
            </p>
          </a>
        </Link>

        <Link href="/learn">
          <a className={styles.card}>
            <h3>Learn &rarr;</h3>
            <p>
              Learn about how to Manage, Register and Trade with Accounts in an
              interactive Demo
            </p>
          </a>
        </Link>

        <Link href="/examples">
          <a className={styles.card}>
            <h3>Examples &rarr;</h3>
            <p>Discover Amazing Accounts of Professional Gamers</p>
          </a>
        </Link>

        <Link href="/signup">
          <a className={styles.card}>
            <h3>Sign Up&rarr;</h3>
            <p>Sign Up with Tradaccs now for free!</p>
          </a>
        </Link>
      </div>
    </main>
  );
}
