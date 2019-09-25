[![Coverage Status](https://coveralls.io/repos/github/axetroy/s4/badge.svg?branch=master)](https://coveralls.io/github/axetroy/s4?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/axetroy/s4)](https://goreportcard.com/report/github.com/axetroy/s4)
![License](https://img.shields.io/github/license/axetroy/s4.svg)
![Repo Size](https://img.shields.io/github/repo-size/axetroy/s4.svg)

## tcp

Check if the tcp port is available

### Usage

```bash
tcp host:port
```

### Installation

Download the executable file for your platform at [release page](https://github.com/axetroy/tcp/releases)

Then set the environment variable.

eg, the executable file is in the `~/bin` directory.

```bash
# ~/.bash_profile
export PATH="$PATH:~/bin"
```

finally, try it out.

```bash
tcp 192.168.0.1:8080
```

### Build from source code

```bash
> go get -v -u github.com/axetroy/tcp
> cd $GOPATH/src/github.com/axetroy/tcp
> make build
```

### Test

```bash
make test
```

### License

The [MIT License](https://github.com/axetroy/tcp/blob/master/LICENSE)
