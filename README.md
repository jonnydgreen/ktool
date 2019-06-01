# ktool CLI

</h3>
<p align="center">
  A CLI for making kubernetes a piece of caketes.
</p>
<p align="center">
  <a href="https://github.com/projectjudge/ktool/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="ktool is released under the MIT license." />
  </a>
  <a href="https://travis-ci.org/projectjudge/ktool">
    <img src="https://travis-ci.org/projectjudge/ktool.svg?branch=master" alt="Current TravisCI build status." />
  </a>
  <a href="https://greenkeeper.io/">
    <img src="https://badges.greenkeeper.io/projectjudge/ktool.svg" alt="Current Greenkeeper status" />
  </a>
  <a >
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs welcome!" />
  </a>
  <a href="https://twitter.com/intent/follow?screen_name=gaussianfrog">
    <img src="https://img.shields.io/twitter/follow/gaussianfrog.svg?label=Follow%20@gaussianfrog" alt="Follow @gaussianfrog" />
  </a>
</p>

## Building

At the moment, ktool is not published to npm yet so you'll need to clone the repo to get started at this alpha stage.

```bash
git clone https://github.com/projectjudge/ktool.git
cd ktool
go build -o bin/ktool ./cmd/ktool
```

Or

```bash
go get github.com/projectjudge/ktool
```

### Devving

```bash
gomon **/*.go -- go build -o bin/ktool ./cmd/ktools
```

Open a new pane/terminal and run:

```shell
$ ktool -h
ktool version 0.0.1

  ktool       ktool CLI welcome screen
  all-pods (a)    Get all pods
  context (c)     Changes the current context of your kubernetes config
  namespace (n)   Changes the current namespace of your kubernetes config
  logs (l)        Shows the logs of a selected pod
  help (h)        -
  version (v)     Output the version number
```

## Publishing to NPM

To package your CLI up for NPM, do this:

```shell
$ npm login
$ npm whoami
$ npm lint
$ npm test
(if typescript, run `npm run build` here)
$ npm publish
```
