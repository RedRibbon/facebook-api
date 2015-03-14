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

## API

### GET most-commented-feeds

댓글 많이 달린 게시글 조회

#### Resouce URL

```
http://yhbyun.redribbon.io:9200/v1/most-commented-feeds
```

#### Parameters

* from : 통계시작일 (YYYY, YYYYMM, YYYYMMDD)
* to : 통계종료일 (YYYY, YYYYMM, YYYYMMDD)
* limit : 조회할 레코드 수 (defult: 20)

> `from`이 생략되면 처음부터 `to`까지
> `to`가 생략되면 `from`에서 끝까지
> `from`, `to`가 모두 생략되면 전체에서 통계를 구한다.

### GET most-posted-persons

사용자별 게시글 등록 수 조회


#### Resource URL

```
http://yhbyun.redribbon.io:9200/v1/most-posted-persons
```

#### Parameters


* from : 통계시작일 (YYYY, YYYYMM, YYYYMMDD)
* to : 통계종료일 (YYYY, YYYYMM, YYYYMMDD)
* limit : 조회할 레코드 수 (default:20)
