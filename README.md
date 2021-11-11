# go-api-deliveryhero-ens
In memory key-value store olarak çalışan bir REST-API
servisi

Standart Kütüphaneler kullanılmıştır


# Özellikler
key ’i set etmek için bir endpoint
key ’i get etmek için bir endpoint
Komple data’yı flush etmek için bir endpoint ?
(zorunlu değil)
Belirli bir interval’da (N dakikada bir) dosyaya
kaydetmeli
Uygulama durup tekrar ayağa kalktığında, eğer
kaydedilmiş dosya varsa, tekrar varolan verileri
hafızaya yükelemeli ( /tmp/TIMESTAMP-data.json ?)
