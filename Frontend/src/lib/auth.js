import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import bcrypt from "bcryptjs";
import { authConfig } from "./auth.config";
import { jwtDecode } from "jwt-decode";

export const login = async (credentials) => {
  const url = `${process.env.NEXT_PUBLIC_BASE_API}${process.env.NEXT_PUBLIC_LOGIN}`;

  try {
    const response = await fetch(url, {
      method: "POST",
      body: JSON.stringify(credentials),
      headers: {
        "Content-Type": "application/json",
      },
    });

    const responseData = await response.json();

    if (!response.ok) {
      throw new Error(responseData.message || "Something went wrong");
    }

    // Decode JWT upon success and return with undecoded JWT
    const data = {
      ...jwtDecode(responseData.data.token),
      accessToken: responseData.data.token,
    };
    console.log("from login at auth.js", data);
    return data;
  } catch (error) {
    console.error("Login error:", error);
    throw error;
  }
};

export const {
  handlers: { GET, POST },
  auth,
  signIn,
  signOut,
} = NextAuth({
  ...authConfig,
  providers: [
    CredentialsProvider({
      async authorize(credentials) {
        try {
          const user = await login(credentials);
          return user;
        } catch (err) {
          return null;
        }
      },
    }),
  ],
  callbacks: {
    ...authConfig.callbacks,
  },
});
