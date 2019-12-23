# Go Go Stardew Cal

This is a command-line tool written in Go to keep track of
things you want to do each day in your Stardew Valley game, such
as whose birthday it is so you can give them a gift, when the
Night Market arrives, when the traveling merchant arrives, etc.

![Screenshot](screenshot1.png)

## How to run

I built this app using Go version 1.13.4.

```sh
make
bin/stardewcal
```

It will prompt you for the current season and day in your Stardew Valley
game, then tell you if any birthdays, festivals, or other notable events
are happening that day. You can keep going forward a day at a time as
you play the game.

If you're modifying the app and want to build it and run it in one step,
do:

```sh
go run cmd/stardewcal/main.go
```
