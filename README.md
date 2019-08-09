# Go Web-Service Scaffold

To bootstrap new project:

1. Clone or `go get` this repo

```shell script
mkdir -p $GOPATH/src/github.com/lancer-kit/
cd $GOPATH/src/github.com/lancer-kit/
git clone https://github.com/lancer-kit/service-scaffold
### OR
go get github.com/lancer-kit/service-scaffold
```

2. Go to scaffold directory and run `./init.sh`

```shell script
cd $GOPATH/src/github.com/lancer-kit/service-scaffold
sh ./init.sh
```

3. Get `forge` — a tool for code generation:

```shell script
go get -u github.com/lancer-kit/forge
```

#### Example

```shell script
cd $GOPATH/src/github.com/lancer-kit/service-scaffold
sh ./init.sh 
Enter VCS domain (default: github.com): gitlab.com
Enter VCS username or group: inn4sci-go
Enter project name: api
```

