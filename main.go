package main

import (
	"fmt"

	"github.com/alexbrainman/printer"
)

func main() {

	// Ambil nama default printer di windows
	printerName, _ := printer.Default()

	// Membuka printer
	p, err := printer.Open(printerName)
	if err != nil {
		fmt.Println(err)
	}
	// Menutup printer setelah selesai digunakan
	defer p.Close()

	// Memberikan nama/judul document di queue/antrian yang akan di cetak
	err = p.StartRawDocument("Mencetak Hello World")
	if err != nil {
		fmt.Println(err)
	}
	// Menutup document/file setelah selesai digunakan
	defer p.EndDocument()

	// Memulai halaman untuk dicetak
	err = p.StartPage()
	if err != nil {
		fmt.Println(err)
	}

	// Contoh string  yang akan dicetak
	textPolos := "=============================\n         HELLO WORLD\n=============================\nIni akan dicetak di printer default windows.\nJangan lupa berikan star dan sematkan sumber jika ini membantu anda.\nTERIMA KASIH"

	// Mulai mencetak string ke printer default yang ada di windows
	fmt.Fprintf(p, "%s\r\n", textPolos)
}
