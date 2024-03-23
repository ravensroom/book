import { Chapter } from "@/models/chapter";

type ChapterViewProps = {
  chapter: Chapter;
};

export function ChapterView(props: ChapterViewProps) {
  const { chapter } = props;
  return <div>{chapter.title}</div>;
}
