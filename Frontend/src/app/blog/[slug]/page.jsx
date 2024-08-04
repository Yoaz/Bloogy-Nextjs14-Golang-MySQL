import Image from "next/image";
import styles from "./singlePost.module.css";
import PostUser from "@/components/postUser/postUser";
import { Suspense } from "react";
import { getPost } from "@/lib/data";

export const generateMetadata = async ({ params }) => {
  const { slug } = params;

  const post = await getPost(slug);

  return {
    title: post.title,
    description: post.body,
  };
};

const SinglePostPage = async ({ params }) => {
  const { slug } = params;

  const data = await getPost(slug);
  const post = data.data.post;

  return (
    <div className={styles.container}>
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
      <div className={styles.textContainer}>
        <h1 className={styles.title}>{post.title}</h1>
        <div className={styles.detail}>
          {post && (
            <Suspense fallback={<div>Loading...</div>}>
              <PostUser author={post.author} />
            </Suspense>
          )}
          <div className={styles.detailText}>
            <span className={styles.detailTitle}>Published</span>
            <span className={styles.detailValue}>
              {post.CreatedAt.toString().slice(0, 10)}
            </span>
          </div>
        </div>
        <div className={styles.content}>{post.body}</div>
      </div>
    </div>
  );
};

export default SinglePostPage;
