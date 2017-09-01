# Klotski

![effect](http://oid1xlj7h.bkt.clouddn.com/image/jpg/Screen%20Shot%202017-09-01%20at%204.21.55%20PM.png)

cd to `application`

and build ：

```shell
go build
```

execute it：

```shell
./application
```

If you want to load your game file：

```shell
./application file-to-path.kls
```

1.There is some character in game and the same name in different pane represent one character who cover multi pane. As you can see `caocao(cc)` cover 4 pane.

2.You can only move one character one step at one times to empty pane.

3.It`s easy to find that there only two pane is empty.

4.If you want move character `zu` to right, you just input `zu 1 0` and press enter.

5.If you want to move one character to up input `name 0 -1`, to down `name 0 1`, to left `name -1 0`.

6.If you want to save you schedule just input `save`, One there is a new file named like `klotski-2017-****.kls`.

7.When you wand load your game file which you saved before,run game with `./application filename.kls`.

8.Once you move the Boss `cc` to the bottom of ChessBoard, you win.

9.Input `q` for quit, please insure your game schedule is saved.

Have a fun!

