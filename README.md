# chiken
The icon of kotaoue

## Usage
```
# basic
go run main.go

# walk style
go run main.go -s=walk
```

## Result
|args|theme|style|size|background|image|
|---|---|---|---|---|---|
||white|basic|32*32|transparent|![white](img/white.png)|
|-s=walk|white|walk|32*32|transparent|![white_walk](img/white_walk.png)|
|-f=gif -s=basic-walk|white|basic-walk|32*32|transparent|![white_basic-walk](img/white_basic-walk.gif)|
|-f=gif -s=basic-walk -d=16|white|basic-walk|32*32|transparent|![white_basic-walk_delay16](img/white_basic-walk_delay16.gif)|
|-t=black|black|basic|32*32|transparent|![black](img/black.png)|
|-b=#ffffff|white|basic|32*32|#ffffff|![white_ffffff](img/white_ffffff.png)|
|-m=2|white|basic|64*64|transparent|![white_2](img/white_2.png)|
|-m=3|white|basic|96*96|transparent|![white_3](img/white_3.png)|