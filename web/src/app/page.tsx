import { getAllBooks } from "@/api/book";
import { Book } from "@/components/home/Book";
import Link from "next/link";

export default function Home() {
  const books = getAllBooks();

  return (
    <main className="p-12">
      <ul className="flex flex-wrap gap-5">
        {books.map((book) => (
          <Link key={book.id} href={`/book/${book.id}`}>
            <Book book={book} />
          </Link>
        ))}
      </ul>
    </main>
  );
}
