import PostCard from "@/components/postCard/postCard";
import styles from "./blog.module.css";
import { getPosts } from "@/lib/data";
import { isEmptyArray } from "@/lib/helpers";

const BlogPage = async () => {
  const data = await getPosts();
  const posts = data.data.posts;

  return (
    <div className={styles.container}>
      {isEmptyArray(posts) && (
        <h1 className={styles.emptyBlog}>
          Create you first blog post! What are you waiting for?!
        </h1>
      )}
      {posts.map((post) => (
        <div className={styles.post} key={post.ID}>
          <PostCard post={post} />
        </div>
      ))}
    </div>
  );
};

export default BlogPage;
