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
|-s=wide|white|wide|32*32|transparent|![white_wide](img/white_wide.png)|
|-s=tiptoe|white|tiptoe|32*32|transparent|![white_tiptoe](img/white_tiptoe.png)|
|-s=jump|white|jump|32*32|transparent|![white_jump](img/white_jump.png)|
|-f=gif -s=basic-walk|white|basic-walk|32*32|transparent|![white_basic-walk](img/white_basic-walk.gif)|
|-f=gif -s=basic-walk -d=16|white|basic-walk|32*32|transparent|![white_basic-walk_delay16](img/white_basic-walk_delay16.gif)|
|-f=gif -s=basic-tiptoe -d=16|white|basic-tiptoe|32*32|transparent|![white_basic-tiptoe_delay16](img/white_basic-tiptoe_delay16.gif)|
|-f=gif -s=basic-jump -d=16|white|basic-jump|32*32|transparent|![white_basic-jump_delay16](img/white_basic-jump_delay16.gif)|
|-t=black|black|basic|32*32|transparent|![black](img/black.png)|
|-b=#ffffff|white|basic|32*32|#ffffff|![white_ffffff](img/white_ffffff.png)|
|-m=2|white|basic|64*64|transparent|![white_2](img/white_2.png)|
|-m=3|white|basic|96*96|transparent|![white_3](img/white_3.png)|
|-f=gif -s=basic-tiptoe-basic-tiptoe-basic-jump -d=64 -m=3|white|basic-tiptoe-basic-tiptoe-basic-jump|96*96|transparent|![white_basic-tiptoe-basic-tiptoe-basic-jump_3_delay64](img/white_basic-tiptoe-basic-tiptoe-basic-jump_3_delay64.gif)|
