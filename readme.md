# Puzzle game

A small project of a web game. It's a muzzle game, you spawn in a random position of the path and you have to find the portal to win.
It's all created randomly, so you can find a path that's linear, or a block or a realy complicated. <br>
The map scale based on your screen dimension, so the bigger your screen is the more the muzzle is.

## If you want to make it start

first clone the repository

```
git clone https://github.com/mattemello/puzzle-game
```

Then you need to have ``` go installed https://go.dev/doc/install ```.<br>
Extract the wasm file ``` cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" THIS_REPOSITORY/assets/js``` (if it's not in $(go env GOROOT)/... you have to search it).

(you need to set a .env file for the port of the server)

Compile the wasm file: 

```
cd src/wasm
GOOS=js GOARCH=wasm go build -o ../../assert/main.wasm
cd ../..
make server
```

And then you can go in the localhost port

## Description

This is little game created by me. It's use golang, javascript and wasm. <br>
The visualizatios it's menaged by the js. <br>
The path, the hero, the arena it's created and menaged by the Golang file. <br>
The server is a Golang server.

For moving you can use vim motion, wasd and the arrow keys.
 
