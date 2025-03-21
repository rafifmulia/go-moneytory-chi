# I don't get it
- Gunanya server-side code generation buat apa kalau ujung-ujungnya harus membuat struct response json sendiri?
- Buat apa kalau ujung-ujungnya http request juga harus implementasi sendiri?

## Sample code to use openapi server-side code generation & your own implementation
```go
r := initRouter()
h := api.HandlerFromMux(api.ServerImpl{}, r)
err = http.ListenAndServe("127.0.0.1:8080", h)
```

## Rerences
- https://i4o.dev/blog/oapi-codegen-with-chi-router
- https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#impl-chi
