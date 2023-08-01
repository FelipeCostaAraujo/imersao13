import FlowbiteContext from "./components/FlowbiteContex";
import DefaultNavbar from "./components/Navbar";
import "./globals.css";

export const metadata = {
  title: "Home Broker",
  description: "Imersao fullcycle 13",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="dark">
      <body className="bg-gray-900 h-screen flex flex-col">
        <DefaultNavbar />
        <FlowbiteContext>{children}</FlowbiteContext>
      </body>
    </html>
  );
}
