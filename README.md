Server: (https://viblo.asia/p/tao-web-api-crud-don-gian-voi-golang-va-echo-framework-Do754kYJlM6)
1. Tạo folder chứa project:
{{{
	$ cd $GOPATH/src
	$ mkdir echo-example-crud
}}}

* $GOPATH mặc định ở C:/User/<user_name>/go

2. Cài đặt echo framework
- Sử dụng go get:
{{{
	$ go get -u github.com/labstack/echo/...
}}}

- Sử dụng go dep:
{{{
	$ cd <PROJECT IN $GOPATH>
	$ dep ensure -add github.com/labstack/echo@^3.1
}}}

3. Xây dựng service:

echo-example-crud
	|__ README.md
	|__ service.go
	|__ server
		|__ server.go

4. Lệnh chạy server 
{{{
	$ go run server/server.go
}}}

Client:

1. Tại thư mục : echo-example-crud chạy lệnh tạo webpack
{{{
	vue init webpack client
}}}

* Nếu chưa hiểu vue cli thì chạy lệnh cài vue-clie

{{{
	npm install -g vue-cli
}}}

2. Cấu trúc cây thư mục sẽ là:
echo-example-crud
	|__ README.md
	|__ service.go
	|__ server
		|__ server.go
	|__ client
- Truy cập client và chạy lệnh install
{{{
	cd client
	npm install
}}}

3. Chạy lệnh run client
{{{
	npm run dev
}}}

================================================= QUIZ ================================================================
|																													  |
|- Phần server: Sử dụng golang echo framework kết nối với một db bất kì. Viết các api đảm bảo CRUD 1 table của SQL	  |
|- Phần client: Sử dụng vuejs, element UI để thực hiện apply source code từ phần server.							  |
|																													  |
=======================================================================================================================

____ *** NOTE ***____
1. Cài đặt mysql: go get github.com/go-sql-driver/mysql
2. Element UI
- Cài đặt element UI:
{{{
	npm i element-ui -S
}}}
- Áp dụng vào source:
{{{
	import ElementUI from 'element-ui'
	import 'element-ui/lib/theme-chalk/index.css'
	Vue.use(ElementUI)
}}}

3. Test API bằng Postman
Ví dụ:
	{{{
		user struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
	}}}
	POSTMAN: POST localhost:<port>/link | BODY: id .. name

