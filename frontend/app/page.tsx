export default function Home() {
  return (
    <div className="h-screen flex items-center justify-center">
      <div>
        <h1 className="text-5xl text-center">create short links</h1>
        <form className="flex justify-center mt-4 gap-x-4">
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
