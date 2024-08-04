"use server";

import { auth } from "@/lib/auth";

// Wrapper to fetch with token as all the api routes (besides login/register) are protected
export const fetchWithToken = async (urlSufix, options = {}) => {
  const session = await auth();

  if (!session) throw new Error("No authentication session found");

  const token = session?.user.accessToken;
  if (!token) throw new Error("No authentication token found");

  // API Base Url
  const apiUrl = process.env.NEXT_PUBLIC_BASE_API;
  const url = `${apiUrl}${urlSufix}`;

  const headers = {
    ...options.headers,
    Authorization: `Bearer ${token}`,
    "Content-Type": "application/json",
  };

  const response = await fetch(url, {
    ...options,
    headers,
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.message || "Something went wrong");
  }

  return data;
};
