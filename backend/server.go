package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ahmetarabaci42/example-backend/backend/database"
)

func main() {

	// 1. Veritabanını Başlat
	database.InitDB()

	// 2. API Endpoint'leri
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/login", loginHandler)

	// 3. Local IP adresini bul
	localIP := getLocalIP()
	port := ":3000"
	
	// Sunucu bilgilerini yazdır
	fmt.Println("========================================")
	fmt.Println("API Sunucusu başlatılıyor...")
	fmt.Printf("Local IP: %s\n", localIP)
	fmt.Printf("Port: %s\n", port)
	fmt.Println("========================================")
	fmt.Printf("\n🔌 API Endpoint'leri:\n")
	fmt.Printf("   - Register: http://%s%s/register\n", localIP, port)
	fmt.Printf("   - Login: http://%s%s/login\n", localIP, port)
	fmt.Println("\n📝 HTML dosyalarını file:// ile açın")
	fmt.Println("   Frontend'den gelen istekler bu sunucuya gönderilecek")
	fmt.Println("\nSunucu çalışıyor, API isteklerini bekliyor...")
	fmt.Println("========================================\n")

	// Sunucuyu tüm network interface'lerinde dinle (0.0.0.0 = tüm IP'ler)
	log.Fatal(http.ListenAndServe("0.0.0.0"+port, nil))

}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// CORS header'larını ekle (file:// protokolü için gerekli)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// OPTIONS isteğini handle et (preflight request)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Sadece POST isteklerini kabul et
	if r.Method != http.MethodPost {
		http.Error(w, "Sadece POST isteği atılabilir", http.StatusMethodNotAllowed)
		return
	}

	// Gelen JSON verisini karşılayacak geçici yapı
	var newUser database.User

	// JSON verisini oku
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Hatalı veri formatı", http.StatusBadRequest)
		return
	}



	// Veritabanına kaydet (db.go'daki fonksiyonu çağırıyoruz)
	err = database.AddUser(newUser.Username, newUser.Email, newUser.Password)
	if err != nil {
		http.Error(w, "Veritabanı hatası: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Başarılı cevabı dön
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Kayıt Başarılı"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// CORS header'larını ekle (file:// protokolü için gerekli)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// OPTIONS isteğini handle et (preflight request)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Sadece POST isteklerini kabul et
	if r.Method != http.MethodPost {
		http.Error(w, "Sadece POST isteği atılabilir", http.StatusMethodNotAllowed)
		return
	}

	// Gelen JSON verisini karşılayacak geçici yapı
	var loginData database.User

	// JSON verisini oku
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "Hatalı veri formatı", http.StatusBadRequest)
		return
	}

	// Veritabanında kullanıcıyı kontrol et
	user, err := database.LoginUser(loginData.Username, loginData.Password)
	if err != nil {
		// Kullanıcı bulunamadı veya hata oluştu
		http.Error(w, "Kullanıcı adı veya şifre hatalı", http.StatusUnauthorized)
		return
	}

	// Kullanıcı bulundu, başarılı
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Giriş başarılı",
		"user":    user,
	})
}

// getLocalIP: Local IP adresini bulan fonksiyon
func getLocalIP() string {
	// Tüm network interface'lerini al
	interfaces, err := net.Interfaces()
	if err != nil {
		return "localhost"
	}

	// Loopback olmayan, aktif interface'leri bul
	for _, iface := range interfaces {
		// Loopback ve down olan interface'leri atla
		if iface.Flags&net.FlagLoopback != 0 || iface.Flags&net.FlagUp == 0 {
			continue
		}

		// Interface'in IP adreslerini al
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		// IPv4 adresini bul
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// IPv4 ve loopback olmayan adresi döndür
			if ip != nil && ip.To4() != nil && !ip.IsLoopback() {
				return ip.String()
			}
		}
	}

	// Local IP bulunamazsa localhost döndür
	return "localhost"
}
