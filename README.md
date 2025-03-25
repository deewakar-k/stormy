## stormy

minimal cli tool for fetching weather, inspired by [rainy](https://github.com/liveslol/rainy).

### preview
![Screenshot From 2025-03-24 21-10-34](https://github.com/user-attachments/assets/4d6b8a2e-78a6-49d1-a0c8-4dae9af91bbc)

![image](https://github.com/user-attachments/assets/762358a6-46f4-44c3-84cb-c906f1048110)


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

