export type Chapter = {
  id: string;
  bookId: string;
  title: string;
  samples: Array<Sample>;
  sourceIds: Array<string>;
};

export type Sample = {
  id: string;
  version: string;
  content: string;
  isPrimary: boolean;
};
