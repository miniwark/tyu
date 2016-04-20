[![license](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://opensource.org/licenses/MIT)

tyu
===
`tyu` is a small command line utility to display basic system informations.

`tyu` display system informations like, OS type, CPUs, memory usage, disk usage and network trafic in real time. `tyu` is intended to be a basic tool to check main informations in a quick and clear way. It **is not** intended to be another `top`, `iftop`, `iotop` (and alternatives) with a detailed monitoring.

Install
-------
`tyu` is written with the [Go language](https://golang.org/), you need to have it available on your system.
Then you can install `tyu` with the `go get` command:

```
go get github.com/miniwark/tyu
```

Building
--------
You can also build `tyu` from sources:

Firstly install [gopsutils](https://github.com/shirou/gopsutil) and [termui](https://github.com/gizak/termui) depedencies:

```
go get github.com/shirou/gopsutil
go get github.com/gizak/termui
```

Then download and build `tyu`:
```
git clone https://github.com/miniwark/tyu.git
cd tyu
go install
```

Usage
-----
Simply lauch `tyu` in the terminal:
```
tyu
```

To quit the program type `q`.

Screenshot
----------
TODO
