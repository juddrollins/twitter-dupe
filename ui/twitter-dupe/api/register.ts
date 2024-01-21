const baseUrl = "https://tr68zpl5d5.execute-api.us-east-1.amazonaws.com";

type RegisterResponse = {
  success: boolean;
};

const callRegister = async (username: string, password: string) => {
  const response = await fetch(`${baseUrl}/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });

  if (!response.ok) throw new Error("Login failed: " + response.statusText);

  console.log(response);

  return { success: true } as RegisterResponse;
};

export default callRegister;
