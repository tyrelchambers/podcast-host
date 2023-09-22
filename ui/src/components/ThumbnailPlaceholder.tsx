import { faImage } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";

const ThumbnailPlaceholder = () => {
  return (
    <div className="w-28 h-2w-28 rounded-lg bg-secondary flex items-center justify-center">
      <FontAwesomeIcon icon={faImage} className="text-3xl" />
    </div>
  );
};

export default ThumbnailPlaceholder;
