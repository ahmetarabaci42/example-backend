package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite sürücüsü
)

// Global değişken: Diğer dosyalardan veritabanına ulaşmak için kullanacağız
var DB *sql.DB

// User yapısı: Senin istediğin "parçalar" burada tanımlanıyor
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"` // Kullanıcı Adı
	Email    string `json:"email"`    // E-posta
	Password string `json:"password"` // Şifre
}

// InitDB: Veritabanını başlatan fonksiyon
func InitDB() {
	var err error
	// 1. Veritabanı dosyasını aç (yoksa oluşturur)
	DB, err = sql.Open("sqlite3", "./scheduler.db")
	if err != nil {
		log.Fatal("Veritabanı açılamadı:", err)
	}

	// 2. Tablo oluşturma sorgusu (SQL)
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		email TEXT,
		password TEXT
	);`

	// 3. Tabloyu oluştur
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Tablo oluşturulamadı:", err)
	}
}

func AddUser(username, email, password string) error {
	// SQL Injection saldırılarını önlemek için "?" (placeholder) kullanıyoruz
	stmt, err := DB.Prepare("INSERT INTO users(username, email, password) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Verileri yerlerine koy ve çalıştır
	_, err = stmt.Exec(username, email, password)
	return err
}


/*func LoginUser(username string, password string) (User, error) {
	stmt, err := DB.Prepare("SELECT id, username, email, password FROM users WHERE username = ? AND password = ?")
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password)
	return User{}, err
}*/

func LoginUser(username string, password string) (User, error) {
	var user User
	
	// SELECT sorgusu için QueryRow kullanılmalı (Exec değil)
	err := DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = ? AND password = ?", username, password).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("kullanıcı bulunamadı")
		}
		return User{}, err
	}
	
	return user, nil
}