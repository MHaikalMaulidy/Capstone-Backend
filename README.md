#  Backend Summarizer

API backend untuk meringkas teks Bahasa Indonesia secara otomatis menggunakan model Python (TextRank) dan menyimpan hasil ringkasan ke PostgreSQL (Supabase).

---

## 📁 Struktur Proyek

- **main.go**: Entry point aplikasi
- **controller/**: Handler HTTP
- **service/**: Logika pemrosesan ringkasan
- **database/**: Koneksi & operasi DB
- **model/**: Definisi model data
- **router/**: Setup routing

##  Cara Menjalankan 

```bash
git clone https://github.com/username/backend-summarizer.git
cd backend-summarizer

docker build -t backend-summarizer .
docker run -p 8080:8080 backend-summarizer
```
## Cara Menggunakan Endpoint

# Request JSON:
```bash
{
  "text": "Teks yang ingin diringkas...",
  "n": 2
}
```
