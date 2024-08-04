import { Suspense } from "react";
import styles from "./dashboard.module.css";
import AdminPosts from "@/components/adminPosts/adminPosts";
import AdminPostForm from "@/components/adminPostForm/adminPostForm";
import AdminUsers from "@/components/adminUsers/adminUsers";
import AdminUserForm from "@/components/adminUserForm/adminUserForm";
import AdminPostEditForm from "@/components/adminPostEditForm/adminPostEditForm";
import { auth } from "@/lib/auth";

const DashboardPage = async () => {
  const session = await auth();

  // Render dasboard based user role "admin"/"user"
  const isAdmin = session?.user.isAdmin;

  return (
    <div className={styles.container}>
      <h2 className={styles.title}>
        Hi, <span>{session?.user.email}</span>
      </h2>
      <div className={styles.row}>
        <div className={styles.col}>
          <Suspense fallback={<div>Loading...</div>}>
            <AdminPosts session={session} />
          </Suspense>
        </div>
        <div className={styles.col}>
          <AdminPostForm />
        </div>
      </div>
      {isAdmin && (
        <div className={styles.row}>
          <div className={styles.col}>
            <Suspense fallback={<div>Loading...</div>}>
              <AdminUsers />
            </Suspense>
          </div>
          <div className={styles.col}>
            <AdminUserForm />
          </div>
        </div>
      )}
      <div className={styles.col}>
        <AdminPostEditForm session={session} />
      </div>
    </div>
  );
};

export default DashboardPage;
