## stormy

minimal cli app for checking the weather, inspired by [rainy](https://github.com/liveslol/rainy).

### preview
![Screenshot From 2025-03-24 21-10-34](https://github.com/user-attachments/assets/4d6b8a2e-78a6-49d1-a0c8-4dae9af91bbc)


![Screenshot From 2025-03-24 21-12-56](https://github.com/user-attachments/assets/b4c1b9bf-9f3f-44a0-8a0e-db42ae20c745)

### installation

```bash
go install github.com/deewakar-k/stormy@latest
```

### build from source

```bash
git clone https://github.com/deewakar-k/stormy.git
cd stormy
go build -o stormy .
```

### usage

add your api key directly in `main.go` before building:

```go
const apiKey = "your_api_key"
```

run directly:

```bash
stormy city_name
```

### notes
- requires an api key from your preferred weather service.
- supports major city names.
- simple and clean output.

---

