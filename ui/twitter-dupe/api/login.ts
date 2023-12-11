import { type } from "os";
import config from "../tailwind.config";

const baseUrl = "https://tr68zpl5d5.execute-api.us-east-1.amazonaws.com";

type LoginResponse = {
  token: string;
};

const callLogin = async (username: string, password: string) => {
  const response = await fetch(`${baseUrl}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });

  if (!response.ok) throw new Error("Login failed: " + response.statusText);

  const jsonResponse = await response.json();

  return {token: jsonResponse.token} as LoginResponse ;
};

export default callLogin;
