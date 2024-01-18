"use client";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";
import Image from "next/image";
import isRouteValid from "../../lib/routes";
import notFound from "../../public/assets/funny-error-404-background-design_1167-219.png";

export default function NotFound() {
  const pathname = usePathname();
  const validRoute = isRouteValid(pathname);
  const router = useRouter();

  useEffect(() => {
    if (validRoute) {
      router.push(pathname);
    }
  }, [pathname, router, validRoute]);


  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <Image src={notFound} alt="404 Not Found" width={500} height={400} />
      <h1 style={{ fontSize: "3rem", margin: "20px 0" }}>
        Oops! Page Not Found
      </h1>
      <Link
        style={{
          fontSize: "1.5rem",
          color: "#0070f3",
          textDecoration: "underline",
        }}
        href="/"
      >
        GO HOME
      </Link>
    </div>
  );
}
