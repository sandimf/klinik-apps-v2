package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ScreeningQuestion struct {
	ID      uuid.UUID
	Label   string
	Type    string
	Options []string
}

func main() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_DSN env required")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	questions := []ScreeningQuestion{
		{ID: uuid.New(), Label: "Tanggal Rencana Pendakian", Type: "date"},
		{ID: uuid.New(), Label: "Jumlah Pendakian Sebelumnya (di atas 2.000 mdpl)", Type: "text"},
		{ID: uuid.New(), Label: "Apakah Anda memiliki riwayat penyakit berikut ini?", Type: "checkbox", Options: []string{"Penyakit jantung", "Asma", "Hipertensi (tekanan darah tinggi)", "Hipotensi (tekanan darah rendah)", "Diabetes", "Masalah paru-paru lainnya", "Cedera sendi/lutut/pergelangan kaki", "Tidak ada dari yang disebutkan"}},
		{ID: uuid.New(), Label: "Kapan terakhir kali Anda melakukan pemeriksaan kesehatan umum?", Type: "select", Options: []string{"Kurang dari 6 bulan yang lalu", "6 bulan - 1 tahun yang lalu", "Lebih dari 1 tahun yang lalu", "Belum pernah melakukan"}},
		{ID: uuid.New(), Label: "Apakah Anda memiliki masalah dengan:", Type: "checkbox", Options: []string{"Pernapasan saat berolahraga berat", "Daya tahan tubuh saat melakukan aktivitas fisik", "Tidak ada masalah di atas"}},
		{ID: uuid.New(), Label: "Apakah Anda sedang dalam pengobatan rutin atau menggunakan obat tertentu? Jika ya, sebutkan:", Type: "checkbox_textarea", Options: []string{"Ya", "Tidak"}},
		{ID: uuid.New(), Label: "Bagaimana Anda menilai kondisi fisik Anda saat ini untuk pendakian (misal: kekuatan otot, keseimbangan, stamina)?", Type: "select", Options: []string{"Sangat baik", "Baik", "Cukup", "Buruk"}},
		{ID: uuid.New(), Label: "Apakah Anda memiliki alergi (terhadap makanan, obat, atau lainnya)? jika Ya, sebutkan:", Type: "checkbox_textarea", Options: []string{"Ya", "Tidak"}},
	}

	for _, q := range questions {
		_, err := db.Exec(ctx, `INSERT INTO screening_questions (id, label, type, options) VALUES ($1, $2, $3, $4) ON CONFLICT (id) DO NOTHING`, q.ID, q.Label, q.Type, q.Options)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Seeder screening_questions selesai!")
}
