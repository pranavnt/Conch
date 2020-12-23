import Head from "next/head";
import styles from "../styles/Home.module.css";

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Conch</title>
        <link rel="icon" href="/conch.png" />
      </Head>

      <h1 className={styles.title} align="center">
        Conch
      </h1>
      <p align="center">
        <img src="/conch.png" width="40%" />
      </p>
      <h2 align="center">
        Store your shell scripts in the cloud and run them anywhere!
      </h2>
      <hr />
      <p className={styles.description}>
        To run a script:{" "}
        <code className={styles.code}>conch run scriptName</code>{" "}
      </p>
      <p className={styles.description}>
        To upload a script:{" "}
        <code className={styles.code}>conch upload filePath scriptName</code>{" "}
      </p>

      <footer className={styles.footer}>
        <a
          href="https://github.com/pranavnt/Conch"
          target="_blank"
          rel="noopener noreferrer"
        >
          Made with ❤️ by @pranavnt
        </a>
      </footer>
    </div>
  );
}
