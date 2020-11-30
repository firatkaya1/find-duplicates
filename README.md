# Nedir ? 
find-duplicates terminal üzerinde çalışan, belirtildiği hedef alandaki tekrarlı dosyaları bulup silinmesini sağlayan basit bir komuttur. 

# Nasıl Kurulur ?

### Adım 1
Github repositoriden proje linkini kopyalayın ve kendi bilgisayarınıza indirin.  
```
git clone https://github.com/firatkaya1/find-duplicates
```
### Adım 2

Klonlanan dosyayı açın. 
```
cd find-duplicates
```
### Adım 3 
Projeyi build edin.
```
go build -o find-duplicates
```
#### Dipnot
Eğer makinanız  **go commond not found** gibi bir hata alırsanız aşağıdaki komutu çalıştırın. Ardından projeyi tekrar build edin. 
```
sudo apt install golang
```
### Adım 4 
Projeyi çalıştırın. 
```
./find-duplicates
```
Örnek bir çıktısı : 
```
A best CLI for delete duplicate files.

Usage:
  findDuplicate [command]

Available Commands:
  author      This command will return author of this project.
  duplicate   Find duplicate and delete easily
  help        Help about any command
  license     Return license type.
  version     This command will return project version.

Flags:
  -h, --help   help for findDuplicate

Use "findDuplicate [command] --help" for more information about a command.

```
### Adım 5 (Opsiyonel)
Her defasında komuta ulaşabilmek için build ettiğimiz dosyayı *./find-duplicates* diyerek çalıştırmak zahmetli gelebilir. Linux kullanıcıları için elde ettiğiniz build dosyasını /bin dosyasının içine kopyalamanız gerekmektedir.
```
sudo cp find-duplicates /bin  
```
### Bitti. 
Artık terminal'den dosyanın kendisi olmadan çalıştırabilir ve kullanabilirsiniz. 

# Nasıl Kullanılır ? 

Bir kullanıcı senaryosuna göre yanlışlıkla aynı dosyadan birden fazla kopyaladığınızı gördünüz. Sonuç olarak bu kopyalanmış dosyaları silmek istiyorsunuz. 
Dosyanızın adı images olsun. 
### Adım 1 
Aşağıda gördüğünüz **-delete** değeri, kopyalanan dosyaların direk silineceğini mi yoksa sadece gösterileceğinimi belirtir. Default değeri false'dir.
Eğer ki sadece dosyaların gösterilmesini istiyorsak, aşağıdaki şekilde terminal ekranına yapıştırın. */home/kaya/Desktop/images* benim kendi path'imi ifade ediyor.
Bunu kendinize uygun düzenleyin.
```
find-duplicates duplicate /home/kaya/Desktop/images  -delete=false
```
Örnek çıktı : 

```
Duplicate File : /home/kaya/Desktop/images/instagram (4th copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (3rd copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (5th copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (another copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (5th copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (6th copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (3rd copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (4th copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (6th copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (3rd copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (another copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (another copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (6th copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (4th copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (5th copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (5th copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (4th copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (6th copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (another copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (3rd copy).png
Total Scanned Files : 28
Total Unique Files : 4
Total Duplicate Files : 24

```
Görebildiğiniz üzere en altta ufak bir sonuç çıkarmaktadır. Herhangibir dosya silme işlemi olmadığı için Silinen toplam dosya sayısı istatistiği bulunmaz.   
**Total Scanned Files** : Toplam taranan dosyayıyı ifade eder.   
**Total Unique Files** : Aslında olması gereken dosya sayısını ifade eder.  
**Total Duplicate Files** : Orjinal dosyaların dışında kalan kopyalanmış dosyaları ifade eder.   
Burada 28 dosya taranmış, sadece 4 orjinal dosya bulunmuş ve geri kalanların hepsi de varolan 4 orjinal dosyanın kopyasıdır.   

Eğer ki dosyaları silmek istiyorsanız aşağıdaki komutu çalıştırın. Bir önceki komuttan tek farkı, delete değerini true vermemizdir. Alternatif olarak 
**-d=true** da tercih edebilirsiniz.

```
find-duplicates duplicate /home/kaya/Desktop/images  -delete=true

```
Örnek çıktı : 
```
Duplicate File : /home/kaya/Desktop/images/facebook (3rd copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (4th copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (another copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (5th copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (3rd copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (5th copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (6th copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (4th copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (5th copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (4th copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (another copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (6th copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (another copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (3rd copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (5th copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (copy).png
Duplicate File : /home/kaya/Desktop/images/facebook (4th copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (3rd copy).png
Duplicate File : /home/kaya/Desktop/images/instagram (6th copy).png
Duplicate File : /home/kaya/Desktop/images/twitter (copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (6th copy).png
Duplicate File : /home/kaya/Desktop/images/whatsapp (another copy).png
Total Scanned Files : 28
Total Unique Files : 4
Total Duplicate Files : 24
Total Deleted Files : 24
```
Eski çıktıdan tek farkı burada toplam silinen dosya sayısının da bulunmasıdır.  
**Total Deleted Files**: Silinen kopya dosyaları ifade eder. 

---
#### Author

Mevcut author bilgisini görebilmek istiyorsanız, aşağıdaki komutu çalıştırın. 
```
find-duplicates author
```

#### Version

Mevcut version bilgisini görebilmek istiyorsanız, aşağıdaki komutu çalıştırın. 
```
find-duplicates version
```

#### License

Mevcut lisans bilgisini görebilmek istiyorsanız, aşağıdaki komutu çalıştırın. 
```
find-duplicates license
```

#### Help

Mevcut tüm komutları görebilmek istiyorsanız help komutunu kullanabilirsiniz.
```
find-duplicates help
```

# Komutu bilgisayarımdan nasıl kaldırabilirim ? 
Build ettiğimiz dosyayı /bin klasörüne kopyalamıştık. Yapmanız gereken tek şey /bin klasörü altındaki find-duplicates dosyasını bulup silmektir. Bunu terminal ekranından aşağıdaki komut ile yapabilirsiniz.

```
sudo rm /bin/find-duplicates
```
---
Bir problemle karşılaşırsanız bana yazabilirsiniz. 

[me@kayafirat.com](mailto:me@kayafirat.com?subject=[GitHub]%20find-duplicate)

