package handler

import (
	"echo/server"
	"fmt"
	"github.com/labstack/echo"
	_ "mysql-master"
	"net/http"
)

type menu struct {
	Id_menu     string
	Nama_menu   string
	Deskripsi   string
	Jenis       string
	Harga       string
	Url_gambar  string
	Total_order string
}

var data []menu

func BacaData(c echo.Context) error {
	menu_makanan()

	return c.JSON(http.StatusOK, data)

}

func TambahData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var url_gambar = c.FormValue("Url_gambar")
	var jenis = c.FormValue("Jenis")
	var harga = c.FormValue("Harga")

	_, err = db.Exec("insert into tbl_menu values (?,?,?,?,?,?)", nil, nama, deskripsi, url_gambar, jenis, harga)

	if err != nil {
		fmt.Println("Menu Gagal Ditambahkan")
		return c.JSON(http.StatusOK, "Gagal Menambahkan Menu")
	} else {
		fmt.Println("Menu Berhasil Ditambahkan")
		return c.JSON(http.StatusOK, "Berhasil Menambahkan Menu")
	}
}

func UbahData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var url_gambar = c.FormValue("Url_gambar")
	var jenis = c.FormValue("Jenis")
	var harga = c.FormValue("Harga")
	var id = c.FormValue("Id_menu")

	_, err = db.Exec("update tbl_menu set nama_menu = ?, deskripsi = ?, harga = ?, jenis = ?,url_gambar = ? where id_menu = ?", nama, deskripsi, harga, jenis, url_gambar, id)

	if err != nil {
		fmt.Println("Menu Gagal Diubah")
		return c.JSON(http.StatusOK, "Gagal Mengubah Menu")
	} else {
		fmt.Println("Menu Berhasil Diubah")
		return c.JSON(http.StatusOK, "Berhasil Mengubah Menu")
	}
}

func HapusData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var id = c.FormValue("Id_menu")

	_, err = db.Exec("delete from tbl_menu where id_menu = ?", id)

	if err != nil {
		fmt.Println("Menu Gagal Dihapus")
		return c.JSON(http.StatusOK, "Gagal Menghapus Menu")
	} else {
		fmt.Println("Menu Berhasil Dihapus")
		return c.JSON(http.StatusOK, "Berhasil Menghapus Menu")
	}
}

func menu_makanan() {
	data = nil
	db, err := server.Koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from tbl_menu")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)

		fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func InputOrder(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()
	var id = c.FormValue("id")
	var nama_pemesan = c.FormValue("nama_pemesan")
	var nomor_telephone = c.FormValue("nomor_telephone")
	var alamat = c.FormValue("alamat")
	var jumlah = c.FormValue("jumlah")

	_, err = db.Exec("insert into tbl_order values (?,?,?,?,?,?)", nil, id, nama_pemesan, nomor_telephone, alamat, jumlah)

	if err != nil {
		fmt.Println("Pesanan Gagal Dibuat")
		return c.HTML(http.StatusOK, "<strong>Gagal Menambahkan Pesanan</strong>")
	} else {
		fmt.Println("Pesanan Berhasil Dibuat")
		return c.HTML(http.StatusOK, "<script>alert('Berhasil Menambahkan Pesanan, Silahkan menunggu konfirmasi dari kami..Terima Kasih');window.location = 'http://localhost:1323';</script>")
	}
	return c.Redirect(http.StatusSeeOther, "/")
}

func BacaPopuler(c echo.Context) error {
	menu_populer()

	return c.JSON(http.StatusOK, data)

}

func menu_populer() {
	data = nil
	db, err := server.Koneksi()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `vw_totalorder` ORDER BY `vw_totalorder`.`total_older` DESC LIMIT 4;")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga, &each.Total_order)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)

		fmt.Println(data)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}
