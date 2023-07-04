"use client";

import { lowlight } from "lowlight/lib/core";
import json from "highlight.js/lib/languages/json";
import CodeBlockLowlight from "@tiptap/extension-code-block-lowlight";
import Document from "@tiptap/extension-document";
import Paragraph from "@tiptap/extension-paragraph";
import Text from "@tiptap/extension-text";
import { EditorContent, useEditor } from "@tiptap/react";

import useRequestStore from "@/stores/request.store";

import "highlight.js/styles/github.css";

lowlight.registerLanguage("json", json);

const Response = () => {
  const { response } = useRequestStore((state) => ({
    response: state.response,
  }));

  const editor = useEditor({
    extensions: [
      Document,
      Paragraph,
      Text,
      CodeBlockLowlight.configure({
        defaultLanguage: "json",
        lowlight,
      }),
    ],
    content: `<pre><code class="language-json">${JSON.stringify(
      response?.data,
      null,
      4
    )}</code></pre>`,
  });

  const statusColor =
    response?.status >= 200 && response?.status < 300
      ? "text-green-500"
      : "text-red-500";

  if (!response || !editor) return null;

  return (
    <div className="flex flex-col gap-2">
      <div className="flex justify-between">
        <h2>Response</h2>
        <span>
          Status:{" "}
          <span className={`${statusColor} font-semibold`}>
            {response.status} {response.statusText}
          </span>
        </span>
      </div>
      <EditorContent
        editor={editor}
        className="bg-white text-black p-2 max-h-96 rounded border-black"
      />
    </div>
  );
};

export default Response;
