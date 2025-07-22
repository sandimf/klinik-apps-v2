import { toast } from "sonner";

/**
 * useToast
 *
 * Hook global untuk menampilkan toast di seluruh aplikasi.
 *
 * Contoh penggunaan:
 *   const toast = useToast();
 *   toast.success("Berhasil!");
 *   toast.error("Gagal!");
 *   toast("Pesan custom");
 */
export function useToast() {
  return toast;
} 