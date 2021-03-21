import Link from "next/link";
import styles from "../../styles/Home.module.css";

export default function Learn() {
  return (
    <>
      <div className={styles.description}>
        It's great to see your interest in learning about Tradaccs!
      </div>
      <Link href="/">
        <div className={styles.grid}>
          <a className={styles.link}>Back to the Homepage</a>
        </div>
      </Link>
    </>
  );
}
