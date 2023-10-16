# rpiGoMtaBuses

Goal is to pick two closest bus lines and the status of the closes outward and inward bound buses on each line.  MTA bus lines picked are those nearest to your gelocation on app start


## Build
```
go build
./rpigomtabuses
```
run the binary to see the output, recompile after updating the env!

I also included an example service file for systemd, can be used by copying to 
`/etc/systemd/system/bus-routes-app.service`
then
```
sudo systemctl daemon-reload
sudo systemctl enable bus-routes-app
sudo systemctl start bus-routes-app
```

## Getting started
uses golang!

## dev
on a 64 bit arch use
```
air
``

on 32 bit, like raspberry pi, use
```
npx nodemon --exec "go run" . --ext "go,json"  --signal SIGTERM
```


### Tooling
To use golang in nvim and with ALE only, so its fast on a headless environment, you can use gopls, gofumpt/gofmt, and goimports to handle automatic imports and opinionated formatting. 

For gopls and everything to work, make sure the gopath is set up right


```
go install golang.org/x/tools/gopls@latest
go install mvdan.cc/gofumpt@latest
go install golang.org/x/tools/cmd/goimports@latest
```

``` bash
# in .profile or .**rc
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```
