export type Chapter = {
    id: string;
    bookId: string;
    title: string;
    main: Array<Main>;
    activeMainId: string;
    sourceIds: string[];
}

export type Main = {
    id: string;
    version: string;
    content: string;
}