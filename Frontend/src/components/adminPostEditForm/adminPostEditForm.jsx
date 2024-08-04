"use client";

import styles from "./adminPostEditForm.module.css";
import { useFormState } from "react-dom";
import { editPost } from "@/lib/action";
import { getPosts, getUserPosts } from "@/lib/data";
import { useState, useEffect } from "react";

const AdminPostEditForm = ({ session }) => {
  const [state, formAction] = useFormState(editPost, undefined);
  const [posts, setPosts] = useState([]);
  const isAdmin = session?.user.isAdmin;
  const [selectedPost, setSelectedPost] = useState(null);
  const [formData, setFormData] = useState({
    title: "",
    img: "",
    body: "",
  });

  useEffect(() => {
    const fetchPosts = async () => {
      const data = isAdmin ? await getPosts() : await getUserPosts();
      setPosts(data.data.posts);
    };

    fetchPosts();
  }, [isAdmin, session]);

  const handlePostChange = (event) => {
    const postId = event.target.value;
    const post = posts.find((p) => p.ID == postId);
    if (post) {
      setSelectedPost(post);
      setFormData({
        title: post.title,
        img: "",
        body: post.body,
      });
    } else {
      setSelectedPost(null);
      setFormData({
        title: "",
        img: "",
        body: "",
      });
    }
  };

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  return (
    <form action={formAction} className={styles.container}>
      <h1 className={styles.title}>Edit Post</h1>
      <select
        name="postId"
        onChange={handlePostChange}
        value={selectedPost?.id || ""}
      >
        <option value="">Select Post</option>
        {posts.map((post) => (
          <option key={post.ID} value={post.ID}>
            {post.title}
          </option>
        ))}
      </select>
      <input type="hidden" name="postId" value={selectedPost?.ID || ""} />
      <input
        type="text"
        name="title"
        placeholder="Title"
        value={formData.title}
        onChange={handleInputChange}
        disabled={!selectedPost}
      />
      <input
        type="text"
        name="img"
        placeholder="Image URL (optional)"
        value={formData.img}
        onChange={handleInputChange}
        disabled={!selectedPost}
      />
      <textarea
        name="body"
        placeholder="Description"
        rows={10}
        value={formData.body}
        onChange={handleInputChange}
        disabled={!selectedPost}
      />
      <button className={styles.button} type="submit" disabled={!selectedPost}>
        Edit
      </button>
      {state?.error && <div className={styles.error}>{state.error}</div>}
    </form>
  );
};

export default AdminPostEditForm;
