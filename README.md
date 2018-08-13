# Go Web-Service Scaffold

To bootstrap new project:

1. Clone or `go get` this repo

```sh
mkdir -p $GOPATH/src/gitlab.inn4science.com/gophers
cd $GOPATH/src/gitlab.inn4science.com/gophers
git clone https://gitlab.inn4science.com/gophers/service-scaffold
### OR
go get gitlab.inn4science.com/gophers/service-scaffold
```

2. Go to scaffold directory and run `./init.sh`
	
```sh
cd $GOPATH/src/gitlab.inn4science.com/gophers/service-scaffold
sh ./init.sh
```

#### Example

```sh
$ cd $GOPATH/src/gitlab.inn4science.com/gophers/service-scaffold
$  sh ./init.sh 
Enter VCS domain (default: gitlab.inn4science.com): gitlab.com
Enter VCS username or group: inn4sci-go
Enter project name: api
```

