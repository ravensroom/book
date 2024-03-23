import { chapters } from "./mockDB/chapters";

export function getChapterById(id: string) {
  return chapters.find((chapter) => chapter.id === id);
}
