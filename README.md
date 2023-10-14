# rpiGoMtaBuses

Goal is to pick two closest bus lines and the status of the closes outward and inward bound buses on each line.  MTA bus lines picked are those nearest to your gelocation on app start



## Getting started
uses golang!

## dev
on a 64 bit arch use
```
air
``

on 32 bit, like raspberry pi, use
```
npx nodemon --exec "go run" ./main.go --signal SIGTERM
```
