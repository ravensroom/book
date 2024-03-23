import { Book } from "@/models/book";

type BookProps = {
  book: Book;
};

export function Book({ book }: BookProps) {
  return (
    <div className="w-52 h-64 bg-blue-100 px-6 py-4 hover:bg-blue-200">
      <h1 className="font-bold">{book.title}</h1>
      <p>{book.description}</p>
    </div>
  );
}
