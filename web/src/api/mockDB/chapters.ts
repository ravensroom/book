import { Chapter } from "@/models/chapter";

export const chapters: Array<Chapter> = [
  {
    id: "a",
    bookId: "1",
    title: "Introduction",
    samples: [
      {
        id: "a1",
        version: "1",
        content: "This is the first sample of chapter a",
        isPrimary: true,
      },
      {
        id: "a2",
        version: "2",
        content: "This is the second sample of chapter a",
        isPrimary: false,
      },
    ],
    sourceIds: ["a1", "a2"],
  },
  {
    id: "b",
    bookId: "1",
    title: "Methodology",
    samples: [
      {
        id: "b1",
        version: "1",
        content: "This is the first sample of chapter b",
        isPrimary: true,
      },
      {
        id: "b2",
        version: "2",
        content: "This is the second sample of chapter b",
        isPrimary: false,
      },
    ],
    sourceIds: ["b1"],
  },
  {
    id: "i",
    bookId: "2",
    title: "Introduction",
    samples: [
      {
        id: "i1",
        version: "1",
        content: "This is the first sample of chapter i",
        isPrimary: true,
      },
      {
        id: "i2",
        version: "2",
        content: "This is the second sample of chapter i",
        isPrimary: false,
      },
    ],
    sourceIds: ["i1"],
  },
  {
    id: "j",
    bookId: "2",
    title: "Methodology",
    samples: [
      {
        id: "j1",
        version: "1",
        content: "This is the first sample of chapter j",
        isPrimary: true,
      },
      {
        id: "j2",
        version: "2",
        content: "This is the second sample of chapter j",
        isPrimary: false,
      },
    ],
    sourceIds: ["j1"],
  },
  {
    id: "p",
    bookId: "3",
    title: "Introduction",
    samples: [
      {
        id: "p1",
        version: "1",
        content: "This is the first sample of chapter p",
        isPrimary: true,
      },
      {
        id: "p2",
        version: "2",
        content: "This is the second sample of chapter p",
        isPrimary: false,
      },
    ],
    sourceIds: ["p1"],
  },
  {
    id: "q",
    bookId: "3",
    title: "Methodology",
    samples: [
      {
        id: "q1",
        version: "1",
        content: "This is the first sample of chapter q",
        isPrimary: true,
      },
      {
        id: "q2",
        version: "2",
        content: "This is the second sample of chapter q",
        isPrimary: false,
      },
    ],
    sourceIds: ["q1"],
  },
];
