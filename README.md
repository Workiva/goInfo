# Introduction:
  A simple package to fetch information about the OS.  Currently supports Linux, Windows, and Darwin
  
## Version:

version:0.0.1  
  
## Install:
```sh
  go get github.com/OneCloudInc/goInfo
  go build
```

## Struct:
```sh
type GoInfoObject struct {
    GoOS string
    Kernel string
    Core string
    Platform string
    OS string
    Hostname string
    CPUs int
    Name string // linux only
    Distribution string // linux only
}
```

## Example:

```sh   
   package main

   import (
	   "github.com/matishsiao/goInfo"
   )

   func main() {
		gi, err := goInfo.GetInfo()
		if err != nil {
		    gi.VarDump()
		}
	 }

```

On Linux:
```sh
   GoOS: linux
   Kernel: Linux
   Core: 3.10.0-693.2.1.el7.x86_64
   Platform: x86_64
   OS: GNU/Linux
   Hostname: localhost.localdomain
   CPUs: 1
   Distribution version: 7
   Distribution name: centos
```
On Windows:
```sh
   GoOS: windows
   Kernel: windows
   Core: 10.0.15063 N/A Build 15063
   Platform: x64-based PC
   OS: Microsoft Windows 10 Pro
   Hostname: windows10
   CPUs: 4
```

##License and Copyright
This software is Copyright 2012-2014 Matis Hsiao.
