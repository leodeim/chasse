<p align="center">
 <img src="img/logo2.png" width="350">
</p>
<div align="center">

  <a href="">![Tests](https://github.com/leonidasdeim/app-chessboard/actions/workflows/go.yml/badge.svg)</a>
  <a href="">![Code Scanning](https://github.com/leonidasdeim/app-chessboard/actions/workflows/codeql.yml/badge.svg)</a>
  <a href="">![Release](https://badgen.net/github/release/leonidasdeim/app-chessboard/)</a>
  <a href="">![Releases](https://badgen.net/github/releases/leonidasdeim/app-chessboard)</a>
  <a href="">![Contributors](https://badgen.net/github/contributors/leonidasdeim/app-chessboard)</a>
  
</div>

# chasse

Just a simple chessboard - without timers, rules etc. Play just like you do it OTB.

<https://chasse.fun>

## Build an API

```bash
cd chasse-api/
GOOS=linux GOARCH=amd64 go build -o build/chasse-api main.go
```

## Build front-end

```bash
cd chasse-app/
npm i
npm run build
```
