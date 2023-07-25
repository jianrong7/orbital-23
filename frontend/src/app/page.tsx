"use client";

import UrlInput from "@/components/UrlInput";
import BodyRequest from "@/components/BodyRequest";
import Response from "@/components/Response";
import { useSearchParams } from "next/navigation";
import { useEffect } from "react";
import useRequestStore from "@/stores/request.store";

export default function Home() {
  const { setUrl } = useRequestStore((state) => ({
    setUrl: state.setUrl,
  }));

  const searchParams = useSearchParams();

  useEffect(() => {
    setUrl(searchParams?.get("url") || "");
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

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
