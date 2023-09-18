# KP UNJUK KETERAMPILAN
    Membuat simple API dengan arsitektur MVC dengan go, gin and mysql  

# CARA MENJALANKAN :
## 1. Buatlah database dengan nama kp-golang-restful-api
## 2. Konfigurasi DB
    Pada file db (config/db) Sesuaikan dengan konfigurasi database server anda :
    db, err := gorm.Open("mysql", "root:@/restful?charset=utf8&parseTime=True&loc=Local")
    
## Tambahkan library
Buka terminal kemudian masukan perintah dibawah ini:

	go get github.com/gin-gonic/gin

	go get golang.org/x/crypto/bcrypt

	go get github.com/go-sql-driver/mysql
	go get github.com/kataras/go-sessions

### Dan terakhir masukan peritah dibawah
	go run main.go