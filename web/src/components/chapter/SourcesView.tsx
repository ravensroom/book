"use client";

import { Source } from "@/models/source";
import { ContextsView } from "../source/ContextsView";
import { getContextsBySourceId } from "@/api/context";
import { useState } from "react";
import { ChatView } from "../source/ChatView";
import { getMessagesBySourceId } from "@/api/message";

type SourcesViewProps = {
  sources: Array<Source>;
};

export function SourcesView(props: SourcesViewProps) {
  const { sources } = props;
  const [selectedSource, setSelectedSource] = useState<Source | undefined>(
    sources[0]
  );

  return (
    <div className="flex flex-col flex-1 bg-gray-200">
      <ul className="flex">
        {sources.map((source) => {
          const isSelected = selectedSource?.id === source.id;
          return (
            <li
              key={source.id}
              onClick={() => setSelectedSource(source)}
              className={`flex flex-col bg-${
                isSelected ? "gray-200" : "gray-100"
              } p-2 border-r-2 border-gray-300 hover:cursor-pointer hover:bg-gray-300`}
            >
              <h3>{source.title}</h3>
            </li>
          );
        })}
      </ul>
      {selectedSource && <SourceView source={selectedSource} />}
    </div>
  );
}

type SourceViewProps = {
  source: Source;
};

function SourceView(props: SourceViewProps) {
  const { source } = props;
  const contexts = getContextsBySourceId(source.id);
  const messages = getMessagesBySourceId(source.id);
  return (
    <div className="flex flex-col flex-1 gap-4 p-2 bg-gray-200">
      <ContextsView contexts={contexts} />
      <ChatView messages={messages} />
    </div>
  );
}
