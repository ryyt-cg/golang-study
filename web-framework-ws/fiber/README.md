







## App
### Static
Use the Static method to serve static files as images, CSS and Javascript.

```go
func (app *App) Static(prefix string, root string, config ...StaticConfig) Router
```

### Route Handlers
Use the Route method to register a route handler for a specific HTTP method and path.

```go
func (app *App) Get(path string, handlers ...Handler) Router
func (app *App) Post(path string, handlers ...Handler) Router
func (app *App) Put(path string, handlers ...Handler) Router
```

Examples;
```go
// Simple GET handler
app.Get("/api/list", func(c *fiber.Ctx) error {
  return c.SendString("I'm a GET request!")
})

// Simple POST handler
app.Post("/api/register", func(c *fiber.Ctx) error {
  return c.SendString("I'm a POST request!")
})
```

### Group
Use the Group method to create a group of routes with a common prefix.

```go
func (app *App) Group(prefix string, handlers ...Handler) Router
```

### Route
You can define routes with a common prefix inside the common function

Signature:
```go
func (app *App) Route(prefix string, fn func(router Router), name ...string) Router
```

### ListenLTS
ListenTLS serves HTTPs requests from the given address using certFile and keyFile paths to as TLS certificate and key file.
```go
func (app *App) ListenTLS(addr, certFile, keyFile string) error
```
### ListenTLSWithCertificate
```go
func (app *App) ListenTLS(addr string, cert tls.Certificate) error
```

## Ctx
There are several methods available in the Ctx struct to handle requests and responses.  For detailed information about the Ctx struct, refer to the [Ctx documentation](https://pkg.go.dev/github.com/gofiber/fiber/v2#Ctx).  I list some of the most commonly used methods below.

### Params
Use the Params method to retrieve a route parameter by its key. You could pass an optional default value that will be returned if the param key does not exist.

Signature:	
```go
func (c *Ctx) Params(key string, defaultValue ...string) string
````


### ParamInt
Use the ParamInt method to retrieve a route parameter as an integer.

Signature:	
```go
func (c *Ctx) ParamInt(key string) (int, error)
```
example:
```go
// GET http://example.com/user/123
app.Get("/user/:id", func(c *fiber.Ctx) error {
  id, err := c.ParamsInt("id") // int 123 and no error

  // ...
})
```

### Queries
Queries is a function that returns an object containing a property for each query string parameter in the route.<br/>
Signature:
```go
func (c *Ctx) Queries() map[string]string
```
Example:
```go
// GET http://example.com/?name=alex&want_pizza=false&id=
app.Get("/", func(c *fiber.Ctx) error {
m := c.Queries()
m["name"] // "alex"
m["want_pizza"] // "false"
m["id"] // ""
// ...
})
````

### Query
This property is an object containing a property for each query string parameter in the route, you could pass an optional default value that will be returned if the query key does not exist.

Signature:
```go
func (c *Ctx) Query(key string, defaultValue ...string) string
```
Example:
```go
// GET http://example.com/?order=desc&brand=nike

app.Get("/", func(c *fiber.Ctx) error {
c.Query("order")         // "desc"
c.Query("brand")         // "nike"
c.Query("empty", "nike") // "nike"

// ...
})
````
There are other query methods: QueryInt, QueryBool, QueryFloat.





## Client
Start a http request with http method and url.<br/>
Signature:
```go
func (c *Client) Get(url string) *Agent
func (c *Client) Head(url string) *Agent
func (c *Client) Post(url string) *Agent
func (c *Client) Put(url string) *Agent
func (c *Client) Patch(url string) *Agent
func (c *Client) Delete(url string) *Agent
```


### Agent
Agent is built on top of FastHTTP's HostClient which has lots of convenient helper methods such as dedicated methods for request methods.

 