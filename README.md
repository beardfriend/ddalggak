Golang Fiber Starterkit with Boilerplate Code Generation


This Boilerplate is designed to start easier.
In long term, you'll have to create many repeated code.
It will help you from boring stuff.

[![Video Label](http://img.youtube.com/vi/9A8xONwJoFE/0.jpg)](https://youtu.be/9A8xONwJoFE)

# How to Generate Code

```bash
# bin/bash
go install github.com/beardfriend/ddalggak-gen@v1.0.3
```
install https://github.com/ent/ent
write orm code in schema directory

```bash
# bin/bash
make ent-init && make ent-generate
```

```bash
# bin/bash
go run -mod=mod github.com/beardfriend/ddalggak-gen -s {schemaFile NameOnly} -m {module path}
```

example
```bash
# bin/bash
go run -mod=mod github.com/beardfriend/ddalggak-gen -s "product" -m "./internal"
```

code will be generated

# Generate Mock File

install https://github.com/vektra/mockery

```bash
# bin/bash
mockery
```

# Generate Swagger docs

install https://github.com/swaggo/swag

```bash
# bin/bash
make swag
```


### used

- fiber/v2 (web framework)
- ent (ORM)
- swaggo/swag (docs)
- air (hot rereload)
- wire (dependency injection)
- jwt 
- mock