"use client";
import Login from "@/components/login";
import { useRouter } from "next/navigation";
import { ReactNode, createContext, useContext, useState } from "react";

type AuthProviderProps = {
  children: ReactNode;
};

type userInfo = {
  username: string;
  password: string;
};

const AuthContext = createContext({
  user: "",
  login: (userInfo: userInfo) => {},
  logout: () => {},
});

export const useAuth = () => {
  return useContext(AuthContext);
};

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [user, setUser] = useState("");
  const router = useRouter()

  const login = (userInfo: userInfo) => {
    console.log("logging in");
    // You may want to perform additional validation here
    setUser(userInfo.username);
    router.push('/home')
  };

  const logout = () => {
    setUser("null");
  };

  if (user === "") {
    return (
      <AuthContext.Provider value={{ user, login, logout }}>
        <Login />
      </AuthContext.Provider>
    );
  }

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {user && children}
    </AuthContext.Provider>
  );
};
