# Deneme AAMCLI Repository

Bu repository, AAMCLI için oluşturulmuş bir test repository'sidir. İçerisinde basit bir Go REST API örneği bulunmaktadır.

## API Özellikleri

- `/` endpoint'i: "Hello, World!" mesajını JSON formatında döndürür
- `/hello/{isim}` endpoint'i: Verilen isme özel karşılama mesajını JSON formatında döndürür

## Kurulum ve Çalıştırma

1. Repository'yi klonlayın:
```bash
git clone https://github.com/oilik/deneme-aamcli.git
cd deneme-aamcli
```

2. Uygulamayı çalıştırın:
```bash
go run main.go
```

3. API'yi test edin:
```bash
# Ana endpoint'i test etme
curl http://localhost:8080/

# İsme özel karşılama
curl http://localhost:8080/hello/Ahmet
```

## API Yanıtları

Ana endpoint (`/`):
```json
{
    "message": "Hello, World!"
}
```

İsme özel endpoint (`/hello/{isim}`):
```json
{
    "message": "Merhaba, Ahmet!"
}
```

## Gereksinimler

- Go 1.21 veya üzeri

## Katkıda Bulunma

1. Bu repository'yi fork edin
2. Feature branch'i oluşturun (`git checkout -b feature/yeniOzellik`)
3. Değişikliklerinizi commit edin (`git commit -am 'Yeni özellik eklendi'`)
4. Branch'inizi push edin (`git push origin feature/yeniOzellik`)
5. Pull Request oluşturun

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır.