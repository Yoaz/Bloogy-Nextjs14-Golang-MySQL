import { Inter } from "next/font/google";
import "./globals.css";
import Header from "@/components/header/Header";
import Footer from "@/components/footer/Footer";
import CookieConsent from "@/components/cookieConsent/CookieConsent";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: {
    default: "Blooogy Homepage",
  },
  description: "Your simple blog",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="container">
          <Header />
          {children}
          <Footer />
        </div>
        <CookieConsent />
      </body>
    </html>
  );
}
