# Đề bài

- Phần server: Sử dụng golang echo framework kết nối với một db bất kì. Viết các api đảm bảo CRUD 1 table của SQL

- Phần client: Sử dụng vuejs, element UI để thực hiện apply source code từ phần server

## TO DO LIST

- [x] Kiến trúc server
- [x] Kết nối database
- [x] API Create
- [x] API Read
- [x] API Update
- [x] API Delete
- [x] Kiến trúc client
- [ ] API Create
- [x] API Read
- [ ] API Update
- [ ] API Delete


## Server

> Tham khảo tại link: <https://viblo.asia/p/tao-web-api-crud-don-gian-voi-golang-va-echo-framework-Do754kYJlM6>

### Tạo folder chứa project

```bash
cd $GOPATH/src
mkdir echo-example-crud
```

> $GOPATH mặc định ở C:/User/<user_name>/go

### Cài đặt echo framework

- Sử dụng go get:

```bash
go get -u github.com/labstack/echo/...
```

- Sử dụng go dep:

```bash
cd <PROJECT IN $GOPATH>
dep ensure -add github.com/labstack/echo@^3.1
```

### Xây dựng cây thư mục

```bash
echo-example-crud
├── README.md
├── server
    ├── controller
        ├── controller.go
    ├── models
├── employee.go
    ├── database
        ├── db_connection.go
    ├── server.go
```

### Lệnh chạy server

```bash
go run server/server.go
```

## Client

### Tạo webpack

```bash
cd echo-example-crud
vue init webpack client
```

> Nếu chưa hiểu vue cli thì chạy lệnh cài vue-clie

```bash
npm install -g vue-cli
```

### Cấu trúc cây thư mục

```bash
echo-example-crud
├── README.md
├── server
        ├── ...
├── client
```

- Truy cập client và chạy lệnh install

```bash
cd client
npm install
```

### Chạy lệnh run client

```bash
npm run dev
```

## NOTE

### Cài đặt mysql: go get github.com/go-sql-driver/mysql

### Element UI

- Cài đặt element UI:

```bash
npm i element-ui -S
```

- Áp dụng vào source:

```go
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)
```

### Fix lỗi cors

> Access-Control-Allow-Origin' header is present on the requested resource.

```go
// Filename: server.go
e := echo.New()

e.Use(middleware.Logger())
e.Use(middleware.Recover())

// CORS default
// Allows requests from any origin wth GET, HEAD, PUT, POST or DELETE method.
// e.Use(middleware.CORS())

// CORS restricted
// Allows requests from any `http://localhost:8080` origin
// wth GET, PUT, POST or DELETE method.
e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"http://localhost:8080"},
    AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
}))
```
