import { Message } from "@/models/source";

type ChatViewProps = {
  messages: Array<Message>;
};

export function ChatView(props: ChatViewProps) {
  const { messages } = props;
  return (
    <div className="flex flex-col flex-1 bg-white p-4 rounded-md">
      <div className="flex-1 flex flex-col gap-2">
        {messages.map((message) => {
          const isUserMessage = message.sender === "USER";
          return (
            <div
              key={message.id}
              className={`flex gap-2 ${
                isUserMessage ? "justify-end" : "justify-start"
              }`}
            >
              <div className={`p-2 rounded-md bg-gray-100`}>
                {message.content}
              </div>
            </div>
          );
        })}
      </div>
      <Input />
    </div>
  );
}

function Input() {
  return (
    <div className="flex gap-2">
      <input
        type="text"
        className="flex-1 p-2 border-2 border-gray-300 rounded-md focus:outline-none focus:border-blue-500"
        placeholder="Type a message..."
      />
      <button className="p-2 bg-blue-500 text-white rounded-md hover:bg-blue-400">
        Send
      </button>
    </div>
  );
}
