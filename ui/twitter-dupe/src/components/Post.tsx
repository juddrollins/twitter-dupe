import Avatar from "@mui/material/Avatar";
import React from "react";

const Post = () => {
  return (
    <div className="flex items-center">
      <Avatar
        src="https://placekitten.com/64/64"
        alt="User Avatar"
        className="mr-2"
      />
      <div>
        <p className="font-bold text-md">{"username"}</p>
        <p className="text-gray-800">{"text"}</p>
      </div>
    </div>
  );
};

export default Post;
