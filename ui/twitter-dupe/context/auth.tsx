"use client";
import Login from "@/components/login";
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

  const login = (userInfo: userInfo) => {
    console.log("logging in");
    // You may want to perform additional validation here
    setUser(userInfo.username);
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
