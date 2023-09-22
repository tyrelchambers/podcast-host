import { faImage } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";

const ThumbnailPlaceholder = () => {
  return (
    <div className="w-32 h-32 rounded-lg bg-secondary flex items-center justify-center">
      <FontAwesomeIcon icon={faImage} className="text-3xl" />
    </div>
  );
};

export default ThumbnailPlaceholder;
