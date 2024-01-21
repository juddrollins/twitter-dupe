import React from "react";
import Image from "next/image";
import LoadingGIF from "../../public/assets/gif-animations-replace-loading-screen-14.gif";

const Loading = () => {
  return (
    <div className="flex flex-col justify-center items-center h-screen">
      <h1 className="text-4xl font-bold">Loadin</h1>
      <Image src={LoadingGIF} alt="logo" width={600} height={600} />
    </div>
  );
};

export default Loading;
