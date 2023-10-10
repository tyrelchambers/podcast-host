import { Button } from "@/components/ui/button";
import EditPodcast from "@/forms/EditPodcast";
import React from "react";

const ShowSettings = () => {
  return (
    <div>
      <EditPodcast />

      <footer className="mt-6">
        <div className="flex gap-4 justify-end">
          <Button variant="outlineSecondary">Transfer ownership</Button>

          <Button variant="destructive">Delete podcast</Button>
        </div>
      </footer>
    </div>
  );
};

export default ShowSettings;
