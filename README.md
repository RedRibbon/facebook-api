# facebook api

## 설치

`revel` 설치

```
$ brew install hg
$ go get github.com/revel/cmd/revel
```

facebok api 소스 clone

```
$ go get github.com/RedRibbon/facebook-api
$ cd $GOPATH
$ cd src
$ git clone https://github.com/RedRibbon/facebook-api.git
$ revel run facebook-api
```

다음 주소로 api 동작 확인 

```
http://localhost:9000/v1/most-commented-feeds
```
