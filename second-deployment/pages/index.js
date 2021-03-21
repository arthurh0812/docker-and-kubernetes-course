import Head from "next/head";
import Link from "next/link";
import Home from "../components/home";
import styles from "../styles/Home.module.css";
import Documentation from "./documentation";
import Learn from "./learn";
import Examples from "./examples";
import Signup from "./signup";

export default function Index() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Tradaccs</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <Link href="/">
        <Home />
      </Link>

      <footer className={styles.footer}>
        <p>
          Powered by <b>onebean</b>
        </p>
      </footer>
    </div>
  );
}
