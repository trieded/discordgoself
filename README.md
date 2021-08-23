# DiscordGo

[![Go Reference](https://pkg.go.dev/badge/github.com/courtier/discordgo.svg)](https://pkg.go.dev/github.com/courtier/discordgo) [![Go Report Card](https://goreportcard.com/badge/github.com/courtier/discordgo)](https://goreportcard.com/report/github.com/courtier/discordgo)

<img align="right" src="https://github.com/courtier/discordgo/blob/master/docs/img/discordgo.png">

This fork aims to provide Discord client capabilities for use by selfbots.
Currently receving messages works properly.

**For help/discussion, join the [discord](https://discord.gg/tZJxXar7td).**

## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` *will always pull the latest tagged release from the master branch.*

```sh
go get github.com/courtier/discordgo
```

### Usage

Import the package into your project.

```go
import "github.com/courtier/discordgo"
```

Construct a new Discord client which can be used to access the variety of 
Discord API functions and to set callback functions for Discord events.

```go
discord, err := discordgo.New("Bot " + "authentication token")
```

See Documentation and Examples below for more detailed information.


## Documentation

**NOTICE**: This library and the Discord API are unfinished.
Because of that there may be major changes to library in the future.

The DiscordGo code is fairly well documented at this point and is currently
the only documentation available.  Both GoDoc and GoWalker (below) present
that information in a nice format.

- [![Go Reference](https://pkg.go.dev/badge/github.com/courtier/discordgo.svg)](https://pkg.go.dev/github.com/courtier/discordgo)
- [![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/courtier/discordgo) 
- Hand crafted documentation coming eventually.


## Examples
There are currently no examples. Maybe in the examples folder?

## Troubleshooting
Join the [discord](https://discord.gg/tZJxXar7td) and ask away if you have questions,.

## Contributing
Contributions are very welcomed, however please follow the below guidelines.

- First open an issue describing the bug or enhancement so it can be
discussed.  
- Try to match current naming conventions as closely as possible.  
- This package is intended to be a low level direct mapping of the Discord API, 
so please avoid adding enhancements outside of that scope without first 
discussing it.
- Create a Pull Request with your changes against the master branch.
