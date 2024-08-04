"use client";

import { useState, useEffect } from "react";
import Styles from "./CookieConsent.module.css";

const CookieConsent = () => {
  const [showBanner, setShowBanner] = useState(false);

  useEffect(() => {
    const consent = localStorage.getItem("cookieConsent");
    if (!consent) {
      setShowBanner(true);
    }
  }, []);

  const handleAccept = () => {
    localStorage.setItem("cookieConsent", "true");
    setShowBanner(false);
  };

  if (!showBanner) {
    return null;
  }

  return (
    <div className={Styles.container}>
      <p className={Styles.txt}>
        We use cookies to ensure you get the best experience on our website. By
        continuing to use our site, you accept our use of cookies.
      </p>
      <button onClick={handleAccept} className={Styles.button}>
        Accept
      </button>
    </div>
  );
};

export default CookieConsent;
