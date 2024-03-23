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
    <div className="flex flex-1 min-h-full">
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
    <div className="flex flex-col flex-1">
      <h1 className="text-2xl font-semibold">{book.title}</h1>
      <p className="text-lg">{book.description}</p>
    </div>
  );
}

type MainContainerProps = {
  className?: string;
  children: React.ReactNode;
};

function MainContainer(props: MainContainerProps) {
  const { className, children } = props;
  return <div className={`flex flex-1 p-4 ${className}`}>{children}</div>;
}
