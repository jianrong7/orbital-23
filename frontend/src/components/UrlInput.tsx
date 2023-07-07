"use client";

import useRequestStore from "@/stores/request.store";

const UrlInput = () => {
  const {
    url,
    setUrl,
    sendRequest,
    isSendingRequest,
    isValidUrl,
    sentRequest,
  } = useRequestStore((state) => ({
    url: state.url,
    setUrl: state.setUrl,
    sendRequest: state.sendRequest,
    isSendingRequest: state.isSendingRequest,
    isValidUrl: state.isValidUrl,
    sentRequest: state.sentRequest,
  }));

  return (
    <div className="flex justify-between gap-4">
      <input
        type="text"
        id="url"
        name="url"
        className={`text-black w-full rounded p-2 ${
          sentRequest && !isValidUrl
            ? "border-red-500 border-2"
            : "border-black"
        }`}
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />
      <button
        type="submit"
        className="bg-blue-500 p-2 rounded"
        onClick={sendRequest}
      >
        {isSendingRequest ? "Loading..." : "Send POST Request"}
      </button>
    </div>
  );
};

export default UrlInput;
