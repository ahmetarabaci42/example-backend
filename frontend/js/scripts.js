// ...existing code...
// API Sunucu Adresi - Go sunucusunun çalıştığı IP ve port
// Bu değeri kendi local IP adresinizle güncelleyin (örn: 'http://192.168.1.24:3000')
const API_BASE_URL = 'YOUR_IP:3000';

// Sayfa tamamen yüklendiğinde çalışması için kodu DOMContentLoaded içine alıyoruz
document.addEventListener('DOMContentLoaded', function() {

    // --- LOGIN SAYFASI İŞLEMLERİ ---
    const loginForm = document.getElementById('loginForm');
    const loginBtn = document.getElementById('loginBtn');

    // Form submitini yakalayıp Go sunucusuna gönder
    if (loginForm) {
        loginForm.addEventListener('submit', function(e) {
            e.preventDefault(); // Sayfanın yenilenmesini engelle

            // 1. Form verilerini al
            const username = document.getElementById('username').value.trim();
            const password = document.getElementById('password').value;

            // 2. Boş alan kontrolü
            if (!username || !password) {
                alert('Lütfen kullanıcı adı ve şifrenizi girin.');
                return;
            }

            // 3. Login verisini JSON formatına hazırla
            const loginData = {
                username: username,
                password: password
            };

            // 4. Backend'e (Go sunucusuna) gönder
            console.log('Login isteği gönderiliyor:', API_BASE_URL + '/login');
            console.log('Gönderilen veri:', loginData);

            fetch(API_BASE_URL + '/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(loginData)
            })
            .then(response => {
                console.log('Sunucu yanıtı:', response.status, response.statusText);
                if (response.ok) {
                    // Giriş başarılı, index.html'ye yönlendir
                    alert('Giriş başarılı!');
                    window.location.href = 'index.html';
                } else {
                    // Sunucudan gelen hata mesajını al ve göster
                    response.text().then(errorMessage => {
                        alert('Giriş başarısız: ' + errorMessage);
                    });
                }
            })
            .catch(error => {
                console.error('Hata detayı:', error);
                alert('Sunucuya bağlanılamadı!\n\nHata: ' + error.message + '\n\nKontrol edin:\n1. Go sunucusu çalışıyor mu?\n2. IP adresi doğru mu? (' + API_BASE_URL + ')');
            });
        });
    }

    // --- KAYIT SAYFASINA YÖNLENDİRME ---
    const goToRegisterBtn = document.getElementById('goToRegisterBtn');
    if (goToRegisterBtn) {
        goToRegisterBtn.addEventListener('click', function() {
            window.location.href = 'register.html';
        });
    }

// --- REGISTER SAYFASI İŞLEMLERİ ---
    const registerForm = document.getElementById('registerForm');
    
    if (registerForm) {
        registerForm.addEventListener('submit', function(e) {
            e.preventDefault(); // Sayfanın yenilenmesini engelle

            // 1. Form verilerini al ve trim (boşlukları temizle)
            const username = document.getElementById('reg_username').value.trim();
            const email = document.getElementById('reg_email').value.trim();
            const password = document.getElementById('reg_password').value;
            const passwordConfirm = document.getElementById('reg_password_confirm').value;

            // 2. Validasyonlar - Her şey uygun mu kontrol et
            // 2.1. Boş alan kontrolü
            if (!username || !email || !password || !passwordConfirm) {
                alert('Lütfen tüm alanları doldurun.');
                return;
            }

            // 2.2. Kullanıcı adı uzunluk kontrolü
            if (username.length < 3) {
                alert('Kullanıcı adı en az 3 karakter olmalıdır.');
                return;
            }

            // 2.3. Email format kontrolü
            const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
            if (!emailRegex.test(email)) {
                alert('Lütfen geçerli bir e-posta adresi girin.');
                return;
            }

            // 2.4. Şifre uzunluk kontrolü
            if (password.length < 6) {
                alert('Şifre en az 6 karakter olmalıdır.');
                return;
            }

            // 2.5. Şifre eşleşme kontrolü
            if (password !== passwordConfirm) {
                alert('Şifreler uyuşmuyor.');
                return;
            }

            // 3. Her şey uygunsa userData'yı JSON formatına hazırla
            const userData = {
                username: username,
                email: email,
                password: password
            };

            // 4. Backend'e (Go sunucusuna) gönder
            fetch(API_BASE_URL + '/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(userData)
            })
            .then(response => {
                if (response.ok) {
                    alert('Kayıt işleminiz başarıyla tamamlandı!');
                    window.location.href = 'login.html'; // Başarılıysa yönlendir
                } else {
                    // Sunucudan gelen hata mesajını al ve göster
                    response.text().then(errorMessage => {
                        alert('Kayıt sırasında bir hata oluştu: ' + errorMessage);
                    });
                }
            })
            .catch(error => {
                console.error('Hata:', error);
                alert('Sunucuya bağlanılamadı: ' + error.message);
            });
        });
    }

});
// ...existing code...