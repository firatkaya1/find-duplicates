# find-duplicates
A simple CLI for find duplicate files and delete
## Nedir ? 
find-duplicates terminal üzerinde çalışan, belirtildiği hedef alandaki tekrarlı dosyaları bulup silinmesini sağlayan basit bir komuttur. 

## Nasıl Kurulur ?

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

### Nasıl Kullanılır ? 







