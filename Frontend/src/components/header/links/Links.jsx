"use client";

import { useState, useRef, useEffect } from "react";
import styles from "./links.module.css";
import NavLink from "./navLink/navLink";
import Image from "next/image";
import { handleLogout } from "@/lib/action";
import { links } from "@/lib/info";

const Links = ({ session }) => {
  const [open, setOpen] = useState(false);
  const mobileLinksRef = useRef(null);

  const handleClickOutside = (event) => {
    if (
      mobileLinksRef.current &&
      !mobileLinksRef.current.contains(event.target)
    ) {
      setOpen(false);
    }
  };

  useEffect(() => {
    if (open) {
      document.addEventListener("click", handleClickOutside);
    } else {
      document.removeEventListener("click", handleClickOutside);
    }
    return () => {
      document.removeEventListener("click", handleClickOutside);
    };
  }, [open]);

  return (
    <div className={styles.container}>
      <div className={styles.links}>
        {links.map((link) => (
          <NavLink item={link} key={link.title} />
        ))}
        {session?.user ? (
          <>
            {session?.user && (
              <NavLink item={{ title: "Dashboard", path: "/dashboard" }} />
            )}
            <form action={handleLogout}>
              <button className={styles.logout}>Leave</button>
            </form>{" "}
          </>
        ) : (
          <NavLink item={{ title: "Login", path: "/login" }} />
        )}
      </div>
      <Image
        className={`${styles.menuButton} ${open ? styles.open : ""}`}
        src={open ? "/menu-close.svg" : "/menu-open.svg"}
        alt="Menu"
        width={30}
        height={30}
        onClick={() => setOpen((prev) => !prev)}
      />
      <div
        className={`${styles.mobileLinks} ${open ? styles.open : ""}`}
        ref={mobileLinksRef}
      >
        {links.map((link) => (
          <NavLink
            item={link}
            key={link.title}
            closeMenu={() => setOpen(false)}
          />
        ))}
        {session?.user ? (
          <>
            {session?.user && (
              <NavLink
                item={{ title: "Dashboard", path: "/dashboard" }}
                closeMenu={() => setOpen(false)}
              />
            )}
            <form action={handleLogout}>
              <button className={styles.logout}>Leave</button>
            </form>
          </>
        ) : (
          <NavLink
            item={{ title: "Login", path: "/login" }}
            closeMenu={() => setOpen(false)}
          />
        )}
      </div>
    </div>
  );
};

export default Links;
