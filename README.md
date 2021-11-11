# In memory key-value store olarak çalışan bir REST-API servisi

Standart Kütüphaneler kullanılmıştır

# Özellikler
+ key ’i set etmek için bir endpoint
+ key ’i get etmek için bir endpoint
+ Komple data’yı flush etmek için bir endpoint
+ Belirli bir interval’da (N dakikada bir) dosyaya kaydetmeli
+ Uygulama durup tekrar ayağa kalktığında, eğer kaydedilmiş dosya varsa, tekrar varolan verileri hafızaya yükelemeli ( /temp/data.json)

# Endpointler
+ get
+ add
+ flush

# Kullanım

+ Get
    + İstenilen keyi getirmek için
    + curl --location --request GET 'localhost:8888/get/0'
+ Add
    + Key eklemek için
    + curl --location --request POST 'localhost:8888/api/add?key=enis'
+ Flush
    + Tüm keyleri temizlemek için
    + curl --location --request PUT 'localhost:8888/flush'

# Çalıştırma

+ Standart
   + go run main.go

+ Docker Container
    + docker build -t go-api
    + docker run -p 8888:8888 go-api
