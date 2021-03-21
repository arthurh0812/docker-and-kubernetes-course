import Link from "next/link";
import styles from "../../styles/Home.module.css";

export default function Examples() {
  return (
    <>
      <div className={styles.description}>
        Here are some amazing Accounts, waiting for you to inspect them...
      </div>
      <Link href="/">
        <div className={styles.grid}>
          <a className={styles.link}>Back to the Homepage</a>
        </div>
      </Link>
    </>
  );
}
