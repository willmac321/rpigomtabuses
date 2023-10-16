# rpiGoMtaBuses

Goal is to pick two closest bus lines and the status of the closest outward and inward bound buses on each line.  MTA bus lines picked are those nearest to the lat long supplied in the env at compile time.


### Future work
I want to add a way to find the closest bus stops by geolocation, however the only available api's are paid or require more raspberry pi hardware.  So going to mark that as future work for now.


## Video!
[link to video in repo](https://gitlab.com/willmac321/rpigomtabuses/-/blob/dfc7521d6eb3dd4ec36884b0ce4d7a0be2ac7fca/output.mp4)


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
