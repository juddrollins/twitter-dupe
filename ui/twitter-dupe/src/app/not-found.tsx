'use client'
import Link from 'next/link'
import { usePathname, useRouter } from 'next/navigation';
import { useEffect } from 'react';
import isRouteValid from '../../lib/routes';
 
export default function NotFound() {

  const pathname = usePathname();
  const validRoute = isRouteValid(pathname);
  const router = useRouter();

  console.log(validRoute)



  
  
  return (
    <div>
      <h2>Not Found</h2>
      <p>Could not find requested resource</p>
      <Link href="/">Return Home</Link>
    </div>
  )
}