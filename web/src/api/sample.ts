import { getChapterById } from "./chapter";

export function getPrimarySampleByChapterId(chapterId: string) {
  const samples = getChapterById(chapterId)?.samples;
  return samples?.find((sample) => sample.isPrimary);
}
