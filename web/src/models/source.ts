export type Source = {
    id: string;
    chapterId: string;
    title: string;
    contexts: Array<Context>;
    messageIds: string[];
}

export type Context = {
    id: string;
    content: string;
    description: string;
    assetId?: string;
    sourceId: string;
}

export type Message = {
    id: string;
    content: string;
    sourceId: string;
    assetId?: string;
}

export type Asset = {
    id: string;
    type: string;
    url: string;
    content: string;
    sourceId?: string;
    messageId?: string;
}
