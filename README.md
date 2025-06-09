# goany
**goany** is a lightweight utility that brings JavaScript-like JSON handling to Go.  
It allows you to work with dynamic request and response bodies without defining structs, making your API code cleaner and more flexible.

Use it when building lightweight APIs, rapid prototypes, or dealing with partially-known or flexible JSON schemas.

> 📥 `req.Get("user").String()` to access request data  
> 📤 `res.Set("message", "hello")` to build a response

## Support Middlewares

| Framework | Adapter               |
|----------|------------------------|
| `net/http` | `goanyhttp.WithAny`    |
| `chi`      | `goanychi.WithAny`     |
| `gin`      | `goanygin.WithAny`     |
| `echo`     | `goanyecho.WithAny`    |
| `fiber`    | `goanyfiber.WithAny`   |

## Usage

```go
import "github.com/rabbitprincess/goany/goany"

func AnyHandler(req *goany.Request, res *goany.Response) error {
    name := req.Path("user.name").String()
    age := req.Get("user").Get("age").Int()

    res.Set("message", "Hello " + name).Set("age", age)
    return nil
}
```

### Using net/http

```go
import "github.com/rabbitprincess/goany/middleware/goanyhttp"

mux := http.NewServeMux()
mux.HandleFunc("/hello", goanyhttp.WithAny(AnyHandler))
```

### Using chi

```go
import "github.com/rabbitprincess/goany/middleware/goanychi"

r := chi.NewRouter()
r.Post("/hello", goanychi.WithAny(AnyHandler))
```

### Using gin

```go
import "github.com/rabbitprincess/goany/middleware/goanygin"

r := gin.Default()
r.POST("/hello", goanygin.WithAny(AnyHandler))
```

### Using echo

```go
import "github.com/rabbitprincess/goany/middleware/goanyecho"

e := echo.New()
e.POST("/hello", goanyecho.WithAny(AnyHandler))
```

### Using fiber

```go
import "github.com/rabbitprincess/goany/middleware/goanyfiber"

app := fiber.New()
app.Post("/hello", goanyfiber.WithAny(AnyHandler))
```