"use client";

import { useEffect, useState } from "react";
import { getUsers } from "@/lib/data";
import styles from "./adminUsers.module.css";
import Image from "next/image";
import { deleteUser } from "@/lib/action";

const AdminUsers = () => {
  const [users, setUsers] = useState([]);

  const fetchUsers = async () => {
    const data = await getUsers();
    setUsers(data.data.users);
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  const handleDelete = async (id) => {
    await deleteUser(id);
    setUsers((prevUsers) => prevUsers.filter((user) => user.ID !== id));
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Users</h1>
      {users.map((user) => (
        <div className={styles.user} key={user.ID}>
          <div className={styles.detail}>
            <Image
              src={user.img || "/noAvatar.png"}
              alt=""
              width={50}
              height={50}
            />
            <span>{user.email}</span>
          </div>
          <button
            className={styles.userButton}
            onClick={() => handleDelete(user.ID)}
          >
            Delete
          </button>
        </div>
      ))}
    </div>
  );
};

export default AdminUsers;
