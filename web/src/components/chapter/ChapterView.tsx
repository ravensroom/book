import { Chapter } from "@/models/chapter";
import { EditorView } from "./EditorView";
import { SourcesView } from "./SourcesView";
import { getSourcesByChapterId } from "@/api/source";

type ChapterViewProps = {
  chapter: Chapter;
};

export function ChapterView(props: ChapterViewProps) {
  const { chapter } = props;
  const sample = chapter.samples.find((sample) => sample.isPrimary);
  const sources = getSourcesByChapterId(chapter.id);
  return (
    <div className="flex gap-4 flex-1">
      <EditorView sample={sample} />
      <SourcesView sources={sources} />
    </div>
  );
}
