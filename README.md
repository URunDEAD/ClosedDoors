Description:
===
Centralized security system made in GO and arduino for an university project.

Intended use:
-
One would scan an NFC/RFID Tag, the system would encrypt it and send it to the server which will validate it.

How to build the go API:
===
```
make build
```
Easy deploy with docker:
===
```
docker-compose -f build/docker/docker-compose.yaml up service --build
```
