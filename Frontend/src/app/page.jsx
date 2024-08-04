import Image from "next/image";
import styles from "./home.module.css";
import Link from "next/link";
import { auth } from "@/lib/auth";

const Home = async () => {
  const session = await auth();

  return (
    <div className={styles.container}>
      <div className={styles.textContainer}>
        <h1 className={styles.title}>Your Simple Blog</h1>
        <p className={styles.desc}>
          Create an instance for-any-use-case blog for free and start
          Blooooogy... :)
        </p>
        <div className={styles.buttons}>
          {session?.user ? (
            <>
              <Link className={styles.Link} href="/dashboard">
                <button className={styles.button}>Create Post</button>
              </Link>
              <Link href="/blog">
                <button className={styles.button}>Check Blog</button>
              </Link>{" "}
            </>
          ) : (
            <>
              <Link className={styles.Link} href="/about">
                <button className={styles.button}>Learn More</button>
              </Link>
              <Link href="/login">
                <button className={styles.button}>Start Now!</button>
              </Link>{" "}
            </>
          )}
        </div>
      </div>
      <div className={styles.imgContainer}>
        <Image src="/home.gif" alt="" fill className={styles.heroImg} />
      </div>
    </div>
  );
};

export default Home;
