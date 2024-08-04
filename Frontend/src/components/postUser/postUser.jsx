import styles from "./postUser.module.css";
import Image from "next/image";

const PostUser = async ({ author }) => {
  return (
    <div className={styles.container}>
      <Image
        className={styles.avatar}
        src={author.img ? author.img : "/noavatar.png"}
        alt=""
        width={50}
        height={50}
      />
      <div className={styles.texts}>
        <span className={styles.title}>Author</span>
        <span className={styles.author}>{author.email}</span>
      </div>
    </div>
  );
};

export default PostUser;
