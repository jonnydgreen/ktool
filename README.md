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
  <a >
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs welcome!" />
  </a>
  <a href="https://twitter.com/intent/follow?screen_name=gaussianfrog">
    <img src="https://img.shields.io/twitter/follow/gaussianfrog.svg?label=Follow%20@gaussianfrog" alt="Follow @gaussianfrog" />
  </a>
</p>

## Install

```bash
git clone https://github.com/projectjudge/ktool.git
cp ktool/bin/ktool <somewhere_in_your_path>/ktool
```

Or

```bash
go get github.com/projectjudge/ktool
```

### Devving

Single build:

```bash
./build.sh
```

Watcher build:

```bash
gomon **/*.go -- go build -o bin/ktool .
```

Open a new pane/terminal and run:

```shell
$ ktool -h
Ktool is a simplified kubectl

Usage:
  ktool [flags]
  ktool [command]

Available Commands:
  a           Watch all the pods in a k8s cluster
  help        Help about any command
  l           Watch a pods logs

Flags:
  -h, --help   help for ktool

Use "ktool [command] --help" for more information about a command.
