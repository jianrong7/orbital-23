"use client";

import { lowlight } from "lowlight/lib/core.js";
import json from "highlight.js/lib/languages/json";
import CodeBlockLowlight from "@tiptap/extension-code-block-lowlight";
import Document from "@tiptap/extension-document";
import Paragraph from "@tiptap/extension-paragraph";
import Text from "@tiptap/extension-text";
import { EditorContent, useEditor } from "@tiptap/react";
import { useEffect } from "react";

import useRequestStore from "@/stores/request.store";

import "highlight.js/styles/github.css";
import { getRandomInt } from "@/utils";

lowlight.registerLanguage("json", json);

const BodyRequest = () => {
  const { setJsonBody, isValidJsonBody, sentRequest } = useRequestStore(
    (state) => ({
      setJsonBody: state.setJsonBody,
      isValidJsonBody: state.isValidJsonBody,
      sentRequest: state.sentRequest,
    })
  );

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
  });

  const generateRandomJsonBody = () => {
    const randomJsonBody = JSON.stringify(
      {
        first: getRandomInt(1, 1000),
        second: getRandomInt(1, 1000),
      },
      null,
      2
    );
    editor?.commands.setContent(randomJsonBody);
  };

  useEffect(() => {
    setJsonBody(editor?.getText());
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [editor?.getText()]);

  if (!editor) return null;

  return (
    <div className="flex flex-col gap-2">
      <div className="flex justify-between items-center">
        <h2>JSON Body</h2>
        <button
          className="bg-blue-500 p-2 rounded"
          onClick={generateRandomJsonBody}
        >
          Generate random valid JSON body
        </button>
      </div>
      <div>
        <EditorContent
          editor={editor}
          className={`bg-white text-black p-2 max-h-96 rounded border-black ${
            sentRequest && !isValidJsonBody
              ? "border-red-500 border-2"
              : "border-black"
          }`}
        />
      </div>
    </div>
  );
};

export default BodyRequest;
