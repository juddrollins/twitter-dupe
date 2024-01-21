
export default function ErrorBanner({ message }: { message: string }) {
  return (
    <div className="bg-red-500 text-white p-4 text-center m-10">
      <p className="text-lg font-semibold">{message}</p>
    </div>
  );
}
