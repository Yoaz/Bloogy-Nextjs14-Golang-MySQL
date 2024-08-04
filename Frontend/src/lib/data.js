"use server";

import { fetchWithToken } from "./utils";

/* -------------------------------- GET ----------------------------------*/

// Get all posts hook
export const getPosts = async () => {
  const url = `${process.env.NEXT_PUBLIC_GET_ALL_POSTS}`;
  return await fetchWithToken(url, { method: "GET" });
};

// Get user's all posts hook
export const getUserPosts = async () => {
  const url = `${process.env.NEXT_PUBLIC_GET_USER_ALL_POSTS}`;
  return await fetchWithToken(url, { method: "GET" });
};

// Get single post hook
export const getPost = async (id) => {
  const url = `${process.env.NEXT_PUBLIC_GET_SINGLE_POST}${id}`;
  return await fetchWithToken(url, { method: "GET" });
};

// Get all users hook
export const getUsers = async () => {
  const url = `${process.env.NEXT_PUBLIC_GET_ALL_USERS}`;
  return await fetchWithToken(url, { method: "GET" });
};

// Get single user hook
export const getUser = async (id) => {
  const url = `${process.env.NEXT_PUBLIC_GET_SINGLE_USER}${id}`;
  return await fetchWithToken(url, { method: "GET" });
};
