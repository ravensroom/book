import { books } from "@/api/mockDB/books";

export function getAllBooks() {
  return books;
}

export function getBookById(id: string) {
  return books.find((book) => book.id === id);
}
