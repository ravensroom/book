import { contexts } from "./mockDB/contexts";

export function getContextsBySourceId(sourceId: string) {
  return contexts.filter((context) => context.sourceId === sourceId);
}
