"use client";

import { useEffect, useState } from "react";
import { getPosts, getUserPosts } from "@/lib/data";
import styles from "./adminPosts.module.css";
import Image from "next/image";
import { deletePost } from "@/lib/action";
import { isEmptyArray } from "@/lib/helpers";

const AdminPosts = ({ session }) => {
  const [posts, setPosts] = useState([]);
  const isAdmin = session?.user.isAdmin;

  useEffect(() => {
    const fetchPosts = async () => {
      const data = isAdmin ? await getPosts() : await getUserPosts();
      setPosts(data.data.posts);
    };

    fetchPosts();
  }, [isAdmin, session]);

  const handleDelete = async (id) => {
    await deletePost(id);
    setPosts((prevPosts) => prevPosts.filter((post) => post.ID !== id));
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Posts</h1>
      {isEmptyArray(posts) ? (
        <p className={styles.emptyPosts}>
          Create your first post here at <span>Add New Post</span>
        </p>
      ) : (
        <>
          {posts.map((post) => (
            <div className={styles.post} key={post.ID}>
              <div className={styles.detail}>
                <Image
                  src={post.img || "/noAvatar.png"}
                  alt=""
                  width={50}
                  height={50}
                />
                <span className={styles.postTitle}>{post.title}</span>
              </div>
              <button
                className={styles.postButton}
                onClick={() => handleDelete(post.ID)}
              >
                Delete
              </button>
            </div>
          ))}
        </>
      )}
    </div>
  );
};

export default AdminPosts;
