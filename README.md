This project is the source for http://godoc.org/

[![GoDoc](https://godoc.org/github.com/zhuahrev/s?status.svg)](http://godoc.org/github.com/zhuahrev/s)

# s
golang url shortener

## Install

```
go get github.com/zhuharev/s/...
```

## Usage

```
srv, e := shortener.New("cnf")
	if e != nil {
		panic(e)
	}

	http.Handle("/s/", srv)
	e = http.ListenAndServe(":8089", nil)
	if e != nil {
		panic(e)
	}
```