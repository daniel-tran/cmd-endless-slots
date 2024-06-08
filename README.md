# cmd-endless-slots
A simple but never-ending command line-based slot machine game

This was mainly put together in a vain attempt to learn more about the Go programming language, but is partly inspired by all those people I've seen on public transport playing slot machine games on their phone.

# How to play
Press enter to make the current slot column stop.

# How to customise
The game currently takes a single command that controls how fast the game plays. For example, the following command runs the game in a "slow mode" that has the slot columns change every second:

```bash
cmd-endless-slots.exe 1000
```

You can also modify variables such as `items` and `rows` to change how the game plays.

# How to quit
Press Control + C on your keyboard.

# How to build locally
Run the following command in a Command Prompt shell on Windows:

```go
go build cmd-endless-slots.go
```

# Limitations
This game will only run on Windows machines due to some implementation details being platform-specific.

If you want to run this on other operating systems, you will need to modify the `cls()` function to replicate the same functionality for your target operating system.
