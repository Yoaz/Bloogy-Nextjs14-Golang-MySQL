import Image from "next/image";
import styles from "./about.module.css";

export const metadata = {
  title: "About Blooogy",
  description: "About description",
};

const AboutPage = () => {
  return (
    <div className={styles.container}>
      <div className={styles.textContainer}>
        <h1 className={styles.title}>
          Blooogy is a simple free online blog, fast and with the latest
          technology. Have Fun!
        </h1>
        <p className={styles.desc}>
          With dedication for details, fastest technology, clean UI and very
          friendly approach. Blooogy is your right choice, it&apos;s free and
          always would be!
        </p>
        <div className={styles.boxes}>
          <div className={styles.box}>
            <h1>100 K+</h1>
            <p>Daily New Posts</p>
          </div>
          <div className={styles.box}>
            <h1>10 K+</h1>
            <p>Daily Joining Users</p>
          </div>
          <div className={styles.box}>
            <h1>5 K+</h1>
            <p>Daily New Comments</p>
          </div>
        </div>
      </div>
      <div className={styles.imgContainer}>
        <Image src="/about.png" alt="About Image" fill className={styles.img} />
      </div>
    </div>
  );
};

export default AboutPage;
