import { Source } from "@/models/source";
import { getChapterById } from "./chapter";
import { sources } from "./mockDB/sources";

export function getSourcesByChapterId(chapterId: string) {
  const sources = (getChapterById(chapterId)
    ?.sourceIds.map((sourceId) => getSourceById(sourceId))
    .filter((source) => source !== undefined) ?? []) as Array<Source>;
  return sources;
}

export function getSourceById(id: string) {
  return sources.find((source) => source.id === id);
}
