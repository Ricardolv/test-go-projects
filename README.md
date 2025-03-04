# Testing in go

### Projects:
- primeapp


#### Example, commands:
```console
go test .
```
![test...](./data/primeapp/image-1.png)
![test...](./data/primeapp/image-2.png)

```console
go test -v .
```
![test...](./data/primeapp/image-3.png)
![test...](./data/primeapp/image-4.png)

```console
go test -cover .
```
![test...](./data/primeapp/image-5.png)

```console
go test -coverprofile=coverage.out
```
![test...](./data/primeapp/image-6.png)
![test...](./data/primeapp/image-7.png)

```console
go tool cover -html=coverage.out
```
![test...](./data/primeapp/image-8.png)

---