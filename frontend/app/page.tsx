"use client";

import { FormEvent, useState } from "react";

export default function Home() {
  const [loading, setLoading] = useState(false);
  const [link, setLink] = useState("");
  const [error, setError] = useState("");

  const onSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);
    const formData = new FormData(e.currentTarget);
    const longUrl = formData.get("long-url") as string;
    const BASE_URL = process.env.NEXT_PUBLIC_BASE_BACKEND;
    try {
      setError("");
      const res = await fetch(`${BASE_URL}/create-short-url`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          original_url: longUrl,
          user_id: "fd7ba4bc-6a61-4387-b596-e577507c8dd4",
        }),
      });

      if (!res.ok) {
        const text = await res.text();
        throw new Error(`HTTP ${res.status}: ${text}`);
      }

      const data = await res.json();
      setLoading(false);
      setLink(data.short_url);
    } catch (err) {
      console.error("Error while fetching:", err);
      setError("Failed to create short URL");
    }
  };

  return (
    <div className="h-screen flex items-center justify-center">
      <div>
        <h1 className="text-5xl text-center">create short links</h1>
        <form onSubmit={onSubmit} className="flex justify-center mt-4 gap-x-4">
          <input
            className="px-4 py-2.5 border-2 rounded-4xl min-w-md"
            placeholder="enter long url"
            type="url"
            name="long-url"
          />
          <button className="px-6 py-2.5 bg-blue-600 rounded-4xl cursor-pointer">
            submit
          </button>
        </form>

        {link !== "" && (
          <div className="w-full flex justify-between items-center p-4 mt-8 mx-auto rounded-2xl bg-green-600">
            <a className="text-lg" href={link}>
              {link}
            </a>

            <button
              className="bg-white text-black rounded-xl px-4 py-2 text-md cursor-pointer"
              onClick={() => {
                navigator.clipboard.writeText(link);
              }}
            >
              Copy
            </button>
          </div>
        )}
        {error !== "" && (
          <div className="w-full p-4 mt-8 mx-auto rounded-2xl bg-red-500">
            Error
          </div>
        )}
      </div>
      <div className="absolute bottom-4">
        <p className="text-lg">
          made by artem kim |{" "}
          <a
            className="text-blue-400 hover:text-blue-300 underline"
            href="https://github.com/Kapitar/url-shortener"
          >
            github
          </a>
        </p>
      </div>
    </div>
  );
}
