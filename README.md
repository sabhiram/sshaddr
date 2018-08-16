# sshaddr

golang ssh &amp; scp connection string parser

## Install

```
go get github.com/sabhiram/sshaddr
```

## Test

```
cd ${GOPATH}/src/github.com/sabhiram/sshaddr
go test ./...
```

## Usage

```
addr, err := sshaddr.Parse("user:password@host:port:/tmp/destination")
```

At this point, `addr` contains:

```
addr.User()             // string - username    (required)
addr.Pass()             // string - password    (optional) : ""
addr.Host()             // string - hostname    (required)
addr.Port()             // int    - port        (optional) : 22
addr.Destination()      // string - destination (optional) : ""
```
