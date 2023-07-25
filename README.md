## Cli
Is a cli to help you interact with ms-recommendation-plat-admin

### Install

```shell
go install github.com/antunes-pp/cli/cmd/admin@latest
```

### Commands

- Register new user
```shell
admin register // --dev or -d to use DEV enviroment
```

Example
```shell
➜  admin register -d
🚀 Running in DEV mode
[?] Customer name: Jhon Doe
[?] Customer email: Jhon.doe@Picpay.com
[?] Choose squad to register this customer:  [Use arrows to move, space to select, <right> to all, <left> to none, type to filter, ? for more help]
  [ ]  promoverso
  [ ]  universovendedores
  [ ]  engajamentopf
  [ ]  personalizacaohome
  [ ]  storedesign
  [ ]  busca
> [x]  recomendacao

[?] Do you really want to register this new customer? (y/N) y
```
