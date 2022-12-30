<p align="center">
 <img src="misc/img/logo2.png" width="350">
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

## Build

Builds API binary and APP archive in `build/` folder

```bash

make build
```

## Run locally

Starts Redis using `docker-compose` and two processes: API and APP.

```bash
make run -j2

// cleanup after:
make clean
```
