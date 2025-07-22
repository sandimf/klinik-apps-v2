"use client";
import { useEffect, useState } from "react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { useToast } from "@/hooks/useToast";

export function Navbar() {
  const [user, setUser] = useState<any>(null);
  const toast = useToast();

  useEffect(() => {
    fetch("/api/v1/me", {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    })
      .then((res) => {
        if (!res.ok) throw new Error("Gagal mengambil data user");
        return res.json();
      })
      .then(setUser)
      .catch((err) => toast.error(err.message));
  }, []);

  return (
    <nav className="flex items-center gap-4 p-4 border-b">
      <div className="flex items-center gap-2 ml-auto">
        <Avatar>
          <AvatarImage src={user?.avatar || undefined} alt={user?.email || "avatar"} />
          <AvatarFallback>{user?.email?.[0]?.toUpperCase() || "U"}</AvatarFallback>
        </Avatar>
        <div className="flex flex-col">
          <span className="font-semibold">{user?.email || "-"}</span>
          <span className="text-xs text-muted-foreground">{user?.role || "-"}</span>
        </div>
      </div>
    </nav>
  );
} 