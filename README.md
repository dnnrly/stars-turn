# Stars Turn

An API for getting the next several turns for a Stars! game.

This effort is very much inspired by the efforts of Rick of [TotalHost](https://github.com/ricks03/TotalHost) fame.

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/dnnrly/stars-turn)](https://github.com/dnnrly/stars-turn/releases/latest)
[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/dnnrly/stars-turn/Release%20workflow)](https://github.com/dnnrly/stars-turn/actions?query=workflow%3A%22Release+workflow%22)
[![report card](https://goreportcard.com/badge/github.com/dnnrly/stars-turn)](https://goreportcard.com/report/github.com/dnnrly/stars-turn)
[![godoc](https://godoc.org/github.com/dnnrly/stars-turn?status.svg)](http://godoc.org/github.com/dnnrly/stars-turn)

![GitHub watchers](https://img.shields.io/github/watchers/dnnrly/stars-turn?style=social)
![GitHub stars](https://img.shields.io/github/stars/dnnrly/stars-turn?style=social)
[![Twitter URL](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fdnnrly%2Fstars-turn)](https://twitter.com/intent/tweet?url=https://github.com/dnnrly/stars-turn)


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

To get this API working, there are a few important dependencies that you will need to install. Pay attention, there are a couple of surprising steps. And please note, these instructions are likely to be incorrect in the very short term. If you spot an error or have any difficulty in general then feel free to raise an issue with all of the gory details.

The first important note is that this assumes that you are running on Linux. I'm hoping to get this working on Windows eventually but 1 thing at a time...

The dependencies:

1. Wine (the Windows compatability layer)
2. Stars! - for this you will need a fully licensed copy. I won't tell you how to acquire a license, you'll have to use your google-fu to do that


### Installing

```bash
$ git clone http://github.com/dnnrly/stars-turn.git
$ cd stars-turn
$ make install
```

### Running Unit Tests

```bash
$ make test
```

### Running Acceptance tests

```bash
$ make deps
$ make build acceptance-test
```

## Important `make` targets

* `deps` - downloads all of the deps you need to build, test, and release
* `install` - installs your application
* `build` - builds your application
* `test` - runs unit tests
* `ci-test` - run tests for CI validation
* `acceptance-test` - run the acceptance tests
* `lint` -  run linting
* `update` - update Go dependencies
* `clean` - clean project dependencies
* `clean-deps` - remove all of the build dependencies too


## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/dnnrly/stars-turn/tags). 

## Authors

* **Pascal Dennerly** - *Initial work* - [dnnrly](https://github.com/dnnrly)

See also the list of [contributors](https://github.com/dnnrly/stars-turn/contributors) who participated in this project.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* [Rick](https://github.com/ricks03) has coded up a fairly comprehensive Stars! hosting service
