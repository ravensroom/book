import type { Metadata } from "next";
import "./globals.css";
import { Header } from "@/components/Header";

export const metadata: Metadata = {
  title: "Book",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <meta name="color-scheme" content="light" />
      <body className="min-h-screen flex flex-col">
        <Header />
        <div className="flex flex-1 bg-gray-100">{children}</div>
      </body>
    </html>
  );
}
