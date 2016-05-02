[![license](https://img.shields.io/badge/license-MIT_License-blue.svg?style=flat-square)](https://github.com/miniwark/tyu/blob/master/LICENSE)
[![release](https://img.shields.io/badge/release-v_0.2-blue.svg?style=flat-square)](https://github.com/miniwark/tyu/releases)
[![Build Status](https://drone.io/github.com/miniwark/tyu/status.png)](https://drone.io/github.com/miniwark/tyu/latest)
[![Test coverage](https://img.shields.io/codecov/c/github/miniwark/tyu.svg?style=flat-square)](https://codecov.io/gh/miniwark/tyu)
[![Go Report Card](https://goreportcard.com/badge/github.com/miniwark/tyu?style=flat-square)](https://goreportcard.com/report/github.com/miniwark/tyu)

tyu
===
`tyu` is a small command line utility to display basic system informations.

`tyu` display system informations like, OS type, CPUs, memory usage, disk usage and network trafic in real time. `tyu` is intended to be a basic tool to check main informations in a quick and clear way. It **is not** intended to be another `top`, `iftop`, `iotop` (and alternatives) with a detailed monitoring.


Installation
------------
`tyu` is written with the [Go language](https://golang.org/), you need to have it available on your system.
Then you can install `tyu` with the `go get` command:

    go get github.com/miniwark/tyu


Usage
-----
Simply lauch `tyu` in the terminal:

    tyu

To quit the program type `q`.


Screenshot
----------
![tyu_screenshot](https://cloud.githubusercontent.com/assets/301895/14684766/c8ead42a-0731-11e6-9091-d5dcf4a11a8d.png)


Contribute
----------
The source code and issue tracker are available at [Github](https://github.com/miniwark/tyu).

if your are planing to contribute, please read our [guildelines](https://github.com/miniwark/tyu/blob/master/CONTRIBUTING.md).


Building
--------
You can build `tyu` from sources:

Firstly install [gopsutils](https://github.com/shirou/gopsutil) and [termui](https://github.com/gizak/termui) depedencies:

    go get github.com/shirou/gopsutil
    go get github.com/gizak/termui
    go get github.com/ogier/pflag

then download and build `tyu`:

    git clone https://github.com/miniwark/tyu.git
    cd tyu
    go install

License
-------
`tyu` is licensed under the [MIT license](https://github.com/miniwark/tyu/blob/master/LICENSE).
