### Prerequisites
* Go 1.13 

### Run the project

Example of execution with only correct urls
```
go run main.go -parallel=5 http://twitter.com http://google.com
```

Example of execution with only correct urls without indicating number of parallel requests
```
go run main.go http://twitter.com http://www.amazon.com
```

Example of execution with one wrong url

```
go run main.go -parallel=5 www.google.com amazon.fr htttp://www.jeuxvideo.fr http://www.twitter.com
```

### Build & run the project

Build 
```
go build -o dist/app main.go
```
Run 
```
dist/app -parallel=5 http://twitter.com http://google.com
```