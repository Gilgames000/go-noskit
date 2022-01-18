# go-noskit

> NOTE: this repository is currently unmaintained as I don't play the game anymore, but pull requests are accepted nonetheless!

Go NosKit is a collection of packages, completely interoperable through interfaces, that let you interact with the MMORPG NosTale in a headless way, without requiring the game client.

### Installation
```
go get -u github.com/gilgames000/go-noskit
```

### Example usage
A demo implementation of a bot written using this library can be found [here](https://github.com/Gilgames000/go-noskit/blob/master/example/basic_bot/main.go). The bot logs in a character, moves it near the NosBazaar NPC. Then it opens the bazaar and queries it for an item. Finally it prints the result. It also
assumes that the character is in the Marketplace Area.

### Documentation
The documentation can be found [here](https://pkg.go.dev/github.com/gilgames000/go-noskit).

### License
This software is licensed under the GNU GPL v3 license that can be found [here](https://github.com/Gilgames000/go-noskit/blob/master/LICENSE).
