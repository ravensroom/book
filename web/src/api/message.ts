import { messages } from "./mockDB/messages";

export function getMessagesBySourceId(sourceId: string) {
  return messages.filter((message) => message.sourceId === sourceId);
}
