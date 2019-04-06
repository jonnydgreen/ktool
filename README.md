# Toolketes CLI

[![Greenkeeper badge](https://badges.greenkeeper.io/projectjudge/toolketes.svg)](https://greenkeeper.io/)

</h3>
<p align="center">
  A CLI for making kubernetes a piece of caketes.
</p>
<p align="center">
  <a href="https://github.com/projectjudge/toolketes/blob/master/LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="Toolketes is released under the MIT license." />
  </a>
  <a href="https://travis-ci.org/projectjudge/toolketes">
    <img src="https://travis-ci.org/projectjudge/toolketes.svg?branch=master" alt="Current TravisCI build status." />
  </a>
  <a >
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs welcome!" />
  </a>
  <a href="https://twitter.com/intent/follow?screen_name=gaussianfrog">
    <img src="https://img.shields.io/twitter/follow/gaussianfrog.svg?label=Follow%20@gaussianfrog" alt="Follow @gaussianfrog" />
  </a>
</p>

## Usage

At the moment, toolketes is not published to npm yet so you'll need to clone the repo to get started at this alpha stage.

```shell
$ git clone https://github.com/projectjudge/toolketes.git
$ cd toolketes
$ npm i
$ npm link
```

Open a new pane/terminal and run:

```shell
$ toolketes
$ toolketes -h
$ toolketes context
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
