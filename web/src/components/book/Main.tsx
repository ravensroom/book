"use client";

import { Book } from "@/models/book";
import { Sidebar } from "./Sidebar";
import { useState } from "react";
import { ChapterView } from "../chapter/ChapterView";
import { Chapter } from "@/models/chapter";

type MainProps = {
  book: Book;
};

export function Main(props: MainProps) {
  const { book } = props;
  const [selectedChapter, setSelectedChapter] = useState<Chapter | null>(null);

  return (
    <div className="flex">
      <Sidebar
        book={book}
        onClickBook={() => {
          setSelectedChapter(null);
        }}
        onClickChapter={(chapter) => {
          setSelectedChapter(chapter);
        }}
        onClickAddChapter={() => {
          console.log("Add Chapter");
        }}
      />
      <MainContainer>
        {selectedChapter ? (
          <ChapterView chapter={selectedChapter} />
        ) : (
          <BookView book={book} />
        )}
      </MainContainer>
    </div>
  );
}

type BookViewProps = {
  book: Book;
};

function BookView(props: BookViewProps) {
  const { book } = props;
  return (
    <>
      <h1 className="text-2xl font-semibold">{book.title}</h1>
      <p className="text-lg">{book.description}</p>
    </>
  );
}

type MainContainerProps = {
  className?: string;
  children: React.ReactNode;
};
function MainContainer(props: MainContainerProps) {
  const { className, children } = props;
  return (
    <div
      className={`flex-1 flex flex-col min-h-screen p-4 bg-gray-100 ${className}`}
    >
      {children}
    </div>
  );
}
