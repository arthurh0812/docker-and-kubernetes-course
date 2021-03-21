import Link from "next/link";
import styles from "../../styles/Home.module.css";

export default function Documentation() {
  return (
    <>
      <div className={styles.description}>
        This is the documentation about Tradaccs and its secure workflows that
        happen behind the scenes...
      </div>

      <Link href="/">
        <div className={styles.grid}>
          <a className={styles.link}>Back to the Homepage</a>
        </div>
      </Link>
    </>
  );
}
