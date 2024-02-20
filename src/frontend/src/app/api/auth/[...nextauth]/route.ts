import { apiBaseUrl } from "@/constants";
import authService from "@/services/auth";
import userService from "@/services/user";
import axios from "axios";
import NextAuth from "next-auth";
import { AuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { parse } from "cookie";

export const authOptions: AuthOptions = {
  providers: [
    CredentialsProvider({
      // The name to display on the sign in form (e.g. "Sign in with...")
      name: "Credentials",
      // `credentials` is used to generate a form on the sign in page.
      // You can specify which fields should be submitted, by adding keys to the `credentials` object.
      // e.g. domain, username, password, 2FA token, etc.
      // You can pass any HTML attribute to the <input> tag through the object.
      credentials: {
        email: { label: "Email", type: "email", placeholder: "email" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials, req) {
        if (!credentials) return null;

        if (credentials.email && credentials.password) {
          const user = await authService.login({
            email: credentials.email,
            password: credentials.password,
          });

          if (user) {
            return user;
          }
          // Attempt to authenticate using a cookie if credentials are not provided
        } else {
          const cookies = parse(req.headers?.cookie || ""); // Safely parse cookies
          const authToken = cookies["token"];

          if (authToken) {
            // Retrieve user information
            const axiosInstance = axios.create({
              baseURL: apiBaseUrl,
              headers: { Authorization: `Bearer ${authToken}` },
            });
            try {
              const userProfile = await userService.getMyUserInfo(
                axiosInstance
              );
              console.log(userProfile);
              return userProfile ? userProfile : null;
            } catch (error) {
              return null;
            }
          }
        }
        // If neither method succeeds, return null to indicate authentication failure
        return null;
      },
    }),
  ],
  pages: {
    signIn: "/auth/login",
  },
  session: { strategy: "jwt" },
  callbacks: {
    async jwt({ token, user }) {
      return { ...token, ...user };
    },
    async session({ session, token, user }) {
      session.user = token as any;
      return session;
    },
  },
};

const handler = NextAuth(authOptions);
export { handler as GET, handler as POST };
