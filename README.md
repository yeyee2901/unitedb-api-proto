# Protobuf Repository

- This protobuf repository is meant to be used for `unitedb-api` (using gRPC)
- You can access the development server at: `203.194.113.23`
  - gRPC port: `40`
  - HTTP port (using JSON-REST contract): `41` (access using `http://203.194.113.23`)
  - There's also [Swagger-UI](http://203.194.113.23/swagger-ui) available for use
- For the development, I won't be enforcing HTTPS yet, I plan to enforce HTTPS when it's production ready.
- The production release is not ready yet, as I'm planning to work on all the
  features mentioned below before deploying it.

# TODO

- ✅ GetBattleItem
- 👉 GetHeldItem
- GetPokemon
- GetPokemonStats
- GetPokemonMoves

# How To Use This Protobuf Repository

## Golang

- Add this as a dependency using `go get`, I have prepared a pre-compiled protobuf for it.

```bash
go get github.com/yeyee2901/unitedb-api-proto@latest

# or if you need specific commit
go get github.com/yeyee2901/unitedb-api-proto@<commit-tag>
```
