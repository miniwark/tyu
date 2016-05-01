
Version 0.1
-----------
Implement all this infos and gauges:

#### Host
- [x] Hostname
- [x] Domain
- [x] OS
- [x] OS version
- [x] Platform
- [x] Platform version
- [x] Architecture
- [x] Uptime

#### CPUs
- [x] CPUs
- [x] Vendor
- [x] Model
- [x] Frequency
- [ ] ~~Temperature~~
- [x] CPU usage gauge

The temperature info is only a display placeholder for now. It will be implemented for Raspbian in v0.4 and later for other platforms.

#### BIOS / Motherboard
- [x] Motherboard
- [x] Vendor
- [x] BIOS
- [x] Version
- [x] Date

#### Memory
- [x] RAM usage gauge
- [x] Swap usage gauge

#### Disks
- [x] Display up to 3 disks usage gauges

#### Net
- [x] Net up
- [x] Net down

#### Processes
- [x] Nb of processes
- [x] Nb of running


Version 0.2
-----------
- [x] Add tests
- [x] Continuous integration
- [x] Refactor variables and function names
- [x] Refactor error checking
- [x] Add command line options
- [x] Change the display depending of the number of mounted disk


Version 0.3
-----------
- [ ] Improve testing: try to get rid of the var func() wrappers
- [ ] Improve documentation
- [ ] Publish on sites like go-search.org, golanglibs.com, gopkg.in, gobuild.io...
- [ ] Improved bytes display. Display KiB if < 1 Mib and Mib if < 1 Gib

Version 0.4
-----------
#### CPUs
- [ ] Temperature (Raspbian only)

- [ ] Add an extended help text in a tab
- [ ] Maybe launch the help tab if `tyu -h`

Version 0.5
-----------
- [ ] Improve BSD, OsX & Windows support

Version 0.6
-----------
- [ ] Add a tab with timed series gauges; ram, swap, diskio, cpu, net

Version 0.7
-----------
Version for bugs catching, more tests, suggestions, small improvements.

version 0.8
-----------

#### CPUs
- [ ] Temperature (lm-sebsors)

Future
------
- [ ] change the gauges and values colors to red for high values
- [ ] use something like psutil.process_iter() instead of process.Pids() in process.go
- [ ] maybe a proc list in a tab like `top`...
