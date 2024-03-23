"use client";

import { getChapterById } from "@/api/chapter";
import { Book } from "@/models/book";
import { Chapter } from "@/models/chapter";

type SidebarProps = {
  book: Book;
  onClickBook: () => void;
  onClickChapter: (chapter: Chapter) => void;
  onClickAddChapter: () => void;
};

export function Sidebar(props: SidebarProps) {
  const { book, onClickBook, onClickChapter, onClickAddChapter } = props;

  return (
    <div className="min-h-screen w-56 bg-white">
      <SidebarItem
        title={book.title}
        className="hover:bg-gray-100 font-semibold"
        onClick={onClickBook}
      />
      {book.chapterIds.map((chapterId) => {
        const chapter = getChapterById(chapterId);
        if (!chapter) {
          return null;
        }
        return (
          <SidebarItem
            key={chapter.id}
            title={chapter.title}
            className="hover:bg-gray-100"
            onClick={() => {
              onClickChapter(chapter);
            }}
          />
        );
      })}
      <SidebarItem
        title="Add Chapter"
        className="bg-indigo-100 hover:bg-indigo-200"
        onClick={onClickAddChapter}
      />
    </div>
  );
}

type SidebarItemProps = {
  title: string;
  className: string;
  onClick: () => void;
};

export function SidebarItem(props: SidebarItemProps) {
  const { title, className, onClick } = props;
  return (
    <div
      className={`p-3 hover:cursor-pointer border-solid border-t-2 border-gray-100 ${className}`}
      onClick={onClick}
    >
      {title}
    </div>
  );
}
