package i18n

var IdMessages = Messages{
	RootShort: "",
	RootLong:  "",

	LoginShort:      "Tambah atau autentikasi akun GitHub",
	LoginLong:       "Simpan kredensial akun ke ghx dan atur sebagai akun aktif untuk gh CLI dan git.",
	LoginPrompt:     "Masukkan Personal Access Token GitHub Anda: ",
	LoginSuccess:    "Berhasil login sebagai %s (%s). Akun diatur sebagai aktif.",
	LoginErrorFetch: "Gagal mengambil informasi akun dengan token tersebut. Periksa kembali token Anda.",
	LoginErrorSave:  "Gagal menyimpan konfigurasi akun.",

	AddShort: "Tambah akun GitHub menggunakan token",
	AddLong:  "Simpan kredensial akun ke ghx menggunakan token dan atur sebagai akun aktif.",

	SwitchShort:           "Pindah ke akun GitHub yang tersimpan",
	SwitchLong:            "Pindah kredensial gh CLI dan identitas global git ke akun yang dipilih.",
	SwitchSuccess:         "Berhasil berpindah ke akun '%s' (%s).",
	SwitchAccountNotFound: "Akun '%s' tidak ditemukan. Gunakan 'ghx list' untuk melihat akun yang tersedia.",

	ListShort:          "Tampilkan daftar akun GitHub yang tersimpan",
	ListLong:           "Tampilkan semua akun tersimpan di ghx beserta indikator akun aktif.",
	ListNoAccounts:     "Belum ada akun tersimpan. Gunakan 'ghx login' untuk menambah akun.",
	ListHeaderStatus:   "STATUS",
	ListHeaderUsername: "USERNAME",
	ListHeaderName:     "NAMA",
	ListHeaderEmail:    "EMAIL",

	RemoveShort:           "Hapus akun GitHub dari daftar tersimpan",
	RemoveLong:            "Hapus akun yang ditentukan dari daftar akun ghx.",
	RemoveSuccess:         "Akun '%s' berhasil dihapus.",
	RemoveAccountNotFound: "Akun '%s' tidak ditemukan. Gunakan 'ghx list' untuk melihat akun yang tersedia.",

	EmailShort:   "Lihat atau ubah email akun aktif",
	EmailLong:    "Lihat atau perbarui alamat email untuk akun aktif dan konfigurasi global git.",
	EmailCurrent: "Email git saat ini untuk '%s': %s",
	EmailUpdated: "Berhasil mengubah email '%s' menjadi '%s'.",

	CurrentShort:    "Tampilkan akun yang sedang aktif",
	CurrentLong:     "Tampilkan username, nama, dan email dari akun yang sedang aktif.",
	CurrentActive:   "Akun aktif: %s (%s) <%s>",
	CurrentNoActive: "Belum ada akun yang aktif. Gunakan 'ghx switch <akun>' atau 'ghx login'.",

	LangShort:       "Lihat atau atur bahasa aplikasi",
	LangLong:        "Ganti bahasa antarmuka antara Bahasa Inggris (en) dan Bahasa Indonesia (id).",
	LangCurrent:     "Bahasa saat ini: %s",
	LangUpdated:     "Bahasa berhasil diubah ke %s.",
	LangUnsupported: "Bahasa '%s' tidak didukung. Opsi yang tersedia: en, id",

	ErrGHNotInstalled:  "GitHub CLI (gh) tidak terinstall atau tidak ada di PATH.",
	ErrGitNotInstalled: "Git tidak terinstall atau tidak ada di PATH.",
	ErrConfigLoad:      "Gagal membaca file konfigurasi ghx.",
}
