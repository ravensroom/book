import { getAllBooks, getBookById } from "@/api/book";
import { Main } from "@/components/book/Main";
import { notFound } from "next/navigation";

type Props = {
  params: {
    slug: string;
  };
};

export default function Page(props: Props) {
  const { params } = props;
  const book = getBookById(params.slug);
  if (!book) {
    return notFound();
  }
  return <Main book={book} />;
}

export async function generateStaticParams() {
  const books = getAllBooks();
  return books.map((book) => ({
    slug: book.id,
  }));
}
