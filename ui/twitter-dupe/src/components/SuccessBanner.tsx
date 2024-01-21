import Register from "@/components/register";

export default function RegisterPage({ message }: { message: string }) {
  return (
    <div className="bg-green-500 text-white p-4 text-center m-10">
      <p className="text-lg font-semibold">{message}</p>
    </div>
  );
}
