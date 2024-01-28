"use client";
import Login from "@/components/login";
import { useRouter } from "next/navigation";
import {
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useState,
} from "react";
import useDataApi from "../lib/fetchData";
import callLogin from "../api/login";
import Loading from "@/components/loading";

type AuthProviderProps = {
  children: ReactNode;
};

type userInfo = {
  username: string;
  password: string;
};

const AuthContext = createContext({
  user: null,
  login: (userInfo: userInfo) => {},
  logout: () => {},
});

export const useAuth = () => {
  return useContext(AuthContext);
};

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [user, setUser] = useState(null);
  const router = useRouter();
  const { data, error, loading, fetchData } = useDataApi(callLogin);

  useEffect(() => {
    if (data?.token) {
      const base64Url = data.token.split(".")[1];
      const decodedUserString = atob(base64Url);
      const decodedUser = JSON.parse(decodedUserString);
      setUser(decodedUser);
      console.log(decodedUser)
      router.push("/home");
    }
  }, [data]);

  const login = (userInfo: userInfo) => {
    console.log("logging in");
    fetchData(userInfo.username, userInfo.password);
  };

  const logout = () => {
    setUser(null);
  };

  if (loading) return <Loading />;
  if (error) return <div>error: {error.message}</div>;

  if (!user) {
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
