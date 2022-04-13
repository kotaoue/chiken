# chiken
The icon of kotaoue

## Usage
```
# basic
go run main.go

# walk style
go run main.go -s=walk
```

## Args Example
|args|theme|style|effect|size|background|image|
|---|---|---|---|---|---|---|
||white|basic||32*32|transparent|![white](img/white.png)|
|-s=walk|white|walk||32*32|transparent|![white_walk](img/white_walk.png)|
|-s=wide|white|wide||32*32|transparent|![white_wide](img/white_wide.png)|
|-s=tiptoe|white|tiptoe||32*32|transparent|![white_tiptoe](img/white_tiptoe.png)|
|-s=jump|white|jump||32*32|transparent|![white_jump](img/white_jump.png)|
|-s=sleep|white|sleep||32*32|transparent|![white_sleep](img/white_sleep.png)|
|-s=deepSleep|white|deepSleep||32*32|transparent|![white_deepSleep](img/white_deepSleep.png)|
|-s=wake|white|wake||32*32|transparent|![white_wake](img/white_wake.png)|
|-f=gif -s=basic-walk|white|basic-walk||32*32|transparent|![white_basic-walk](img/white_basic-walk.gif)|
|-f=gif -s=basic-walk -d=16|white|basic-walk||32*32|transparent|![white_basic-walk_delay16](img/white_basic-walk_delay16.gif)|
|-f=gif -s=basic-tiptoe -d=16|white|basic-tiptoe||32*32|transparent|![white_basic-tiptoe_delay16](img/white_basic-tiptoe_delay16.gif)|
|-f=gif -s=basic-jump -d=16|white|basic-jump||32*32|transparent|![white_basic-jump_delay16](img/white_basic-jump_delay16.gif)|
|-t=panda|panda|basic||32*32|transparent|![panda](img/panda.png)|
|-t=brown|brown|basic||32*32|transparent|![brown](img/brown.png)|
|-t=brownBlack|brownBlack|basic||32*32|transparent|![brownBlack](img/brownBlack.png)|
|-t=black|black|basic||32*32|transparent|![black](img/black.png)|
|-t=yellow|yellow|basic||32*32|transparent|![yellow](img/yellow.png)|
|-t=green|green|basic||32*32|transparent|![green](img/green.png)|
|-t=mossGreen|mossGreen|basic||32*32|transparent|![mossGreen](img/mossGreen.png)|
|-t=lightBlue|lightBlue|basic||32*32|transparent|![lightBlue](img/lightBlue.png)|
|-t=blue|blue|basic||32*32|transparent|![blue](img/blue.png)|
|-t=bluePurple|bluePurple|basic||32*32|transparent|![bluePurple](img/bluePurple.png)|
|-t=purple|purple|basic||32*32|transparent|![purple](img/purple.png)|
|-t=pinkPurple|pinkPurple|basic||32*32|transparent|![pinkPurple](img/pinkPurple.png)|
|-t=pink|pink|basic||32*32|transparent|![pink](img/pink.png)|
|-t=red|red|basic||32*32|transparent|![red](img/red.png)|
|-t=orange|orange|basic||32*32|transparent|![orange](img/orange.png)|
|-t=gray|gray|basic||32*32|transparent|![gray](img/gray.png)|
|-t=player2|player2|basic||32*32|transparent|![player2](img/player2.png)|
|-t=player3|player3|basic||32*32|transparent|![player3](img/player3.png)|
|-t=player4|player4|basic||32*32|transparent|![player4](img/player4.png)|
|-t=player5|player5|basic||32*32|transparent|![player5](img/player5.png)|
|-t=vivid|vivid|basic||32*32|transparent|![vivid](img/vivid.png)|
|-t=random|random|basic||32*32|transparent|![random](img/random.png)|
|-t=random -n=random1|random|basic||32*32|transparent|![random1](img/random1.png)|
|-t=random -n=random2|random|basic||32*32|transparent|![random2](img/random2.png)|
|-e=negative|white|basic|negative|32*32|transparent|![white_negative](img/white_negative.png)|
|-e=grayscale|white|basic|grayscale|32*32|transparent|![white_grayscale](img/white_grayscale.png)|
|-e=negative-grayscale|white|basic|negative-grayscale|32*32|transparent|![white_negative-grayscale](img/white_negative-grayscale.png)|
|-f=gif -e=rightLoop1|white|basic|rightLoop1|32*32|transparent|![white_rightLoop1](img/white_rightLoop1.gif)|
|-f=gif -e=rightLoop4|white|basic|rightLoop4|32*32|transparent|![white_rightLoop4](img/white_rightLoop4.gif)|
|-f=gif -e=leftLoop1|white|basic|leftLoop1|32*32|transparent|![white_leftLoop1](img/white_leftLoop1.gif)|
|-e=mirror|white|basic|mirror|32*32|transparent|![white_mirror](img/white_mirror.png)|
|-f=gif -t=white-player2-player3-player4|white-player2-player3-player4|basic||32*32|transparent|![white-player2-player3-player4](img/white-player2-player3-player4.gif)|
|-t=party1|party1|basic||32*32|transparent|![party1](img/party1.png)|
|-f=gif -t=party8|party8|basic||32*32|transparent|![party8](img/party8.gif)|
|-f=gif -t=party16|party16|basic||32*32|transparent|![party16](img/party16.gif)|
|-f=gif -t=party32|party32|basic||32*32|transparent|![party32](img/party32.gif)|
|-f=gif -t=party32 -s=basic-tiptoe|party32|basic-tiptoe||32*32|transparent|![party32_basic-tiptoe](img/party32_basic-tiptoe.gif)|
|-b=#ffffff|white|basic||32*32|#ffffff|![white_ffffff](img/white_ffffff.png)|
|-m=2|white|basic||64*64|transparent|![white_2](img/white_2.png)|
|-f=gif -s=basic-tiptoe-basic-tiptoe-basic-jump -d=64 -m=3|white|basic-tiptoe-basic-tiptoe-basic-jump||96*96|transparent|![white_basic-tiptoe-basic-tiptoe-basic-jump_3_delay64](img/white_basic-tiptoe-basic-tiptoe-basic-jump_3_delay64.gif)|