"use client";

import UrlInput from "@/components/UrlInput";
import BodyRequest from "@/components/BodyRequest";
import Response from "@/components/Response";

export default function Home() {
  return (
    <>
      <header className="px-12 pt-12">
        <h1 className="font-semibold text-xl">Postman Emulator</h1>
      </header>
      <main className="flex flex-col p-12">
        <div className="flex flex-col gap-4">
          <UrlInput />
          <BodyRequest />
          <Response />
        </div>
      </main>
    </>
  );
}
