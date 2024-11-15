import Image from "next/image";
import styles from "./postCard.module.css";
import Link from "next/link";

const PostCard = ({ post }) => {
  return (
    <div className={styles.container}>
      <div className={styles.top}>
        <div className={styles.imgContainer}>
          {post.img ? (
            <Image src={post.img} alt="" fill className={styles.img} />
          ) : (
            <Image
              src="/No-Image-Placeholder.svg"
              alt="Post Image Placeholder"
              fill
              className={styles.img}
            />
          )}
        </div>
        <span className={styles.date}>
          {post.createdAt?.toString().slice(4, 16)}
        </span>
      </div>
      <div className={styles.bottom}>
        <h1 className={styles.title}>{post.title}</h1>
        <p className={styles.desc}>{post.body}</p>
        <Link className={styles.link} href={`/blog/${post.ID}`}>
          READ MORE
        </Link>
      </div>
    </div>
  );
};

export default PostCard;
