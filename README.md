# Go Web-Service Scaffold

To bootstrap new project:

1. Clone or `go get` this repo

```sh
mkdir -p $GOPATH/src/github.com/lancer-kit/armory
cd $GOPATH/src/github.com/lancer-kit/armory
git clone https://github.com/lancer-kit/service-scaffold
### OR
go get github.com/lancer-kit/service-scaffold
```

2. Go to scaffold directory and run `./init.sh`
	
```sh
cd $GOPATH/src/github.com/lancer-kit/service-scaffold
sh ./init.sh
```

#### Example

```sh
$ cd $GOPATH/src/github.com/lancer-kit/service-scaffold
$  sh ./init.sh 
Enter VCS domain (default: gitlab.inn4science.com): gitlab.com
Enter VCS username or group: inn4sci-go
Enter project name: api
```

