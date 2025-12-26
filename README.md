# timer

This is a simple countdown timer.

## To use it:

`go run cmd/baseTimer/baseTimer.go -H <hours> -M <minutes> -S <seconds>`

Example, run a 1 hour, 3 minute, and 4 second timer:

`go run cmd/baseTimer/baseTimer.go -H 1 -M 3 -S 4`

## To build it:

`go build cmd/baseTimer/baseTimer.go`

# To run the binary

Run it just like the `go run ...` example above:

`./baseTimer -H <hours> -M <minutes> -S <seconds>`

Example, run a 1 hour, 3 minute, and 4 second timer:

`./baseTimer -H 1 -M 3 -S 4`

## Stop the timer
Just use ctrl-c I guess.