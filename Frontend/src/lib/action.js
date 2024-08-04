"use server";

import { fetchWithToken } from "./utils";
import { auth, signIn, signOut } from "@/lib/auth";

/* -------------------------------- POST/EDIT/DELETE ----------------------------------*/

// Add post hook
export const addPost = async (prevState, formData) => {
  // const { title, desc, userId } = Object.fromEntries(formData);
  const data = Object.fromEntries(formData);
  const url = `${process.env.NEXT_PUBLIC_EDIT_POST}`;

  // temporary --> no img implementation
  delete data.img; // temporary --> no img implementation

  return await fetchWithToken(url, {
    method: "POST",
    body: JSON.stringify(data),
  });
};

export const editPost = async (prevState, formData) => {
  const data = Object.fromEntries(formData);
  const url = `${process.env.NEXT_PUBLIC_EDIT_POST}${data.postId}`;

  // temporary --> no img implementation
  delete data.img; // temporary --> no img implementation

  return await fetchWithToken(url, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

// Delete post hook
export const deletePost = async (id) => {
  console.log("Deleting post with ID:", id);
  const url = `${process.env.NEXT_PUBLIC_DELETE_POST}${id}`;
  console.log("From deletePost ", url);
  return await fetchWithToken(url, { method: "DELETE" });
};

// Add user hook
export const addUser = async (prevState, formData) => {
  const data = Object.fromEntries(formData);
  // const { email, password, userId } = Object.fromEntries(formData);

  // Transform isAdmin value to role type: "user/admin"
  data.role = data.isAdmin === "true" ? "admin" : "user";
  delete data.isAdmin;
  delete data.img; // temporary --> no img implematation

  const url = `${process.env.NEXT_PUBLIC_REGISTER}`;
  return await fetchWithToken(url, {
    method: "POST",
    body: JSON.stringify(data),
  });
};

// Edit user hook
export const editUser = async (formData) => {
  const data = Object.fromEntries(formData.entries());
  const url = `${process.env.NEXT_PUBLIC_EDIT_USER}${data.id}`;
  return await fetchWithToken(url, {
    method: "PUT",
    body: JSON.stringify(data),
  });
};

// Delete user hook
export const deleteUser = async (id) => {
  console.log("Deleting user with ID:", id);
  const url = `${process.env.NEXT_PUBLIC_DELETE_USER}${id}`;
  const session = await auth();
  const isAdmin = session.user.isAdmin;

  if (!isAdmin) {
    throw new Error("Not authorized to delete users");
  }

  return await fetchWithToken(url, { method: "DELETE" });
};

/* -------------------------------- AUTH ----------------------------------*/

// Register hook
export const register = async (prevState, formData) => {
  const data = Object.fromEntries(formData);

  if (data.password !== data.passwordRepeat) {
    return { error: "Passwords do not match" };
  }

  // temporary --> no img implementation
  delete data.img; // temporary --> no img implementation
  delete data.passwordRepeat; // not required in payload

  const url = `${process.env.NEXT_PUBLIC_BASE_API}${process.env.NEXT_PUBLIC_REGISTER}`;

  const response = await fetch(url, {
    method: "POST",
    body: JSON.stringify({ ...data, role: "user" }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  const responseData = await response.json();

  if (!response.ok) {
    throw new Error(responseData.message || "Something went wrong");
  }

  // Call handleLogin to directly log in the user
  // Currently no email vertification implementation
  const loginData = new FormData();
  loginData.append("email", data.email);
  loginData.append("password", data.password);

  return await handleLogin(null, loginData);
};

// Handle login hook
export const handleLogin = async (prevState, formData) => {
  const { email, password } = Object.fromEntries(formData);

  try {
    await signIn("credentials", { email, password });
  } catch (err) {
    if (err.message.includes("CredentialsSignin")) {
      return { error: "Invalid username or password" };
    }

    if (err.message.includes("Invalid passowrd")) {
      return { error: "Invalid username or password" };
    }

    throw err;
  }
};

// Handle logout hook
export const handleLogout = async () => {
  "use server";
  await signOut();
};
