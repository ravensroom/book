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
      <body className="">
        <Header />
        <div className="min-h-screen bg-gray-100">{children}</div>
      </body>
    </html>
  );
}
