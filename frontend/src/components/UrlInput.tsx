"use client";

import axios from "axios";
import { useState } from "react";
import useRequestStore from "@/stores/request.store";

const UrlInput = () => {
  const { url, setUrl, sendRequest } = useRequestStore((state) => ({
    url: state.url,
    setUrl: state.setUrl,
    sendRequest: state.sendRequest,
  }));

  return (
    <div className="flex justify-between gap-4">
      <input
        type="text"
        id="url"
        name="url"
        className="text-black w-full rounded p-2"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />
      <button
        type="submit"
        className="bg-blue-500 p-2 rounded"
        onClick={sendRequest}
      >
        Send POST Request
      </button>
    </div>
  );
};

export default UrlInput;
