# go-signal-bot

go-signal-bot is a bot written in the Go programming language. It is designed to interact with Signal, the secure and privacy-focused messaging platform. It uses OpenAI's GPT4 or GPT3.5's API to respond meaningfully to any messages that contain the bot's name. It can be added to group chats like Meta AI 
to function like an AI assistant.

# Getting Started

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go1.20
- [signal-cli](https://github.com/AsamK/signal-cli) installed and able to run
- [go-chatgpt](https://pkg.go.dev/github.com/ayush6624/go-chatgpt)
- You have a _separate_ signal account with a different number for the bot
- I used the paid version of TextNow, an Android app, to register with Signal
- OpenAI API key

## Installation

1. Clone the repository:
```bash
cd go-signal-bot
```

2. Edit `config.yml` with necessary configurations
3. Run `go run main.go` to start the listener

# Contributing

Fork it! Please feel free to contribute and create a pull request.

## Roadmap
- [ ] improve coverage
- [ ] improve API with additional functionality offered by CLI
- [ ] improve documentation

# Contact
Skanda Kaashyap (skandakk [ 4t ] gmail [ dot ] com)

# Acknowledgements

Thanks a great deal to the following projects and their maintainers:
1. [signal-cli](https://github.com/AsamK/signal-cli)
2. [go-chatgpt](https://pkg.go.dev/github.com/ayush6624/go-chatgpt)
3. OpenAI GPTs
4. [Signal](https://www.signal.org/) and the [Signal Foundation](https://signalfoundation.org/)

# License

Copyright <2024> <Skanda Kaashyap>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
