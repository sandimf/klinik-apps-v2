import AdminPanelLayout from "@/components/admin-panel/admin-panel-layout";
export default function AdminLayout({ children }: { children: React.ReactNode }) {
    return (
      <AdminPanelLayout>
        {children}
      </AdminPanelLayout>
    );
  }
  