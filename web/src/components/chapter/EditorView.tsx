import { Sample } from "@/models/chapter";

type EditorViewProps = {
  sample: Sample | undefined;
};

export function EditorView(props: EditorViewProps) {
  const { sample } = props;
  if (!sample) return null;

  return (
    <div
      contentEditable="true"
      className="w-[500px] py-2 px-3 bg-white border border-gray-300 rounded-md focus:outline-none focus:border-blue-500 focus:ring focus:ring-blue-100"
    >
      {sample.content}
    </div>
  );
}
