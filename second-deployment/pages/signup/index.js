import Link from "next/link";
import styles from "../../styles/Home.module.css";

export default function Signup() {
  return (
    <>
      <div className={styles.description}>
        Signing up: Creating User Account...
      </div>
      <Link href="/">
        <div className={styles.grid}>
          <a className={styles.link}>Back to the Homepage</a>
        </div>
      </Link>
    </>
  );
}
