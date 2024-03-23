import { Context } from "@/models/source";

type ContextViewProps = {
  contexts: Array<Context>;
};

export function ContextsView(props: ContextViewProps) {
  const { contexts } = props;
  return (
    <div className="flex flex-wrap gap-2 bg-white p-4 rounded-md">
      {contexts.map((context) => (
        <div
          key={context.id}
          className="p-2 h-28 w-28 rounded-md bg-blue-100 hover:bg-blue-200"
        >
          <h4>{context.content}</h4>
        </div>
      ))}
    </div>
  );
}
