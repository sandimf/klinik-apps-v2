import { Dialog, DialogContent, DialogHeader, DialogFooter, DialogTitle, DialogDescription } from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"

interface ConfirmDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  onConfirm: () => void
  onCancel?: () => void
}

export function ConfirmDialog({ open, onOpenChange, onConfirm, onCancel }: ConfirmDialogProps) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Konfirmasi Data Pendaki</DialogTitle>
          <DialogDescription>
            Apakah Anda sudah yakin data pendaki sudah benar? Pastikan data sudah dicek sebelum melanjutkan ke screening.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <Button variant="outline" onClick={() => { onOpenChange(false); onCancel?.() }}>
            Batal
          </Button>
          <Button onClick={() => { onOpenChange(false); onConfirm() }}>
            Yakin & Lanjut
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
} 