Reproduction of https://github.com/golang/open2opaque/issues/21

```
$ /tmp/open2opaque-debug rewrite ./..

Resolving Go package names...
rewriting 3 packages:
  o2o-deadlock
  o2o-deadlock/o2o-deadlock
  o2o-deadlock/o2o-deadlock/pb
Starting the Blaze loader
Loading packages (in batches of up to 20)...
PROCESSED 1 packages (total patterns: 3)
	Last package:         o2o-deadlock
	Total time:           173.534333ms
	Package profile:      TOTAL: 1.141541ms | start 250ns main/scheduled 1.125583ms fix/fixed 11.833µs fix/wrotefiles 459ns fix/donestats 291ns main/fixed 2.917µs main/gotresp 208ns done
	Failures:             0 (0.00%)
	Average time:         173.534333ms
	Estimated until done: 347.068666ms
	Estimated done at:    2026-01-02 15:21:43.254466666 +0100 CET m=+0.559100167
	Error:                <nil>

PROCESSED 2 packages (total patterns: 3)
	Last package:         o2o-deadlock/o2o-deadlock/pb.test
	Total time:           174.798708ms
	Package profile:      TOTAL: 2.364834ms | start 167ns main/scheduled 2.158542ms fix/fixed 143.5µs fix/wrotefiles 458ns fix/donestats 625ns main/fixed 60.875µs main/gotresp 667ns done
	Failures:             0 (0.00%)
	Average time:         87.399354ms
	Estimated until done: 87.399354ms
	Estimated done at:    2026-01-02 15:21:42.996080354 +0100 CET m=+0.300712980
	Error:                <nil>

PROCESSED 3 packages (total patterns: 3)
	Last package:         o2o-deadlock/o2o-deadlock/pb [o2o-deadlock/o2o-deadlock/pb.test]
	Total time:           175.911166ms
	Package profile:      TOTAL: 3.4845ms | start 208ns main/scheduled 3.464542ms fix/fixed 12.666µs fix/wrotefiles 417ns fix/donestats 2.458µs main/fixed 4µs main/gotresp 209ns done
	Failures:             0 (0.00%)
	Average time:         58.637055ms
	Estimated until done: 0s
	Estimated done at:    2026-01-02 15:21:42.909778 +0100 CET m=+0.214410876
	Error:                <nil>

PROCESSED 4 packages (total patterns: 3)
	Last package:         o2o-deadlock/o2o-deadlock/pb
	Total time:           182.982541ms
	Package profile:      TOTAL: 10.625291ms | start 250ns main/scheduled 10.600125ms fix/fixed 15.5µs fix/wrotefiles 458ns fix/donestats 2.708µs main/fixed 6µs main/gotresp 250ns done
	Failures:             0 (0.00%)
	Average time:         45.745635ms
	Estimated until done: -45.745635ms
	Estimated done at:    2026-01-02 15:21:42.871098365 +0100 CET m=+0.175731616
	Error:                <nil>

SIGABRT: abort
PC=0x187a913cc m=0 sigcode=0

goroutine 0 gp=0x1037a9280 m=0 mp=0x1037a9d40 [idle]:
runtime.pthread_cond_wait(0x1037aa2a8, 0x1037aa268)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/sys_darwin.go:547 +0x20 fp=0x16d5a64f0 sp=0x16d5a64c0 pc=0x1028b8ef0
runtime.semasleep(0xffffffffffffffff)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/os_darwin.go:72 +0x78 fp=0x16d5a6550 sp=0x16d5a64f0 pc=0x102895ed8
runtime.notesleep(0x1037a9e90)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/lock_sema.go:62 +0xac fp=0x16d5a6580 sp=0x16d5a6550 pc=0x10286c6bc
runtime.mPark()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:1960 +0x20 fp=0x16d5a65a0 sp=0x16d5a6580 pc=0x10289e9c0
runtime.stopm()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:2999 +0x60 fp=0x16d5a65d0 sp=0x16d5a65a0 pc=0x10289ffb0
runtime.findRunnable()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:3770 +0x30 fp=0x16d5a6770 sp=0x16d5a65d0 pc=0x1028a09b0
runtime.schedule()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:4138 +0x9c fp=0x16d5a67b0 sp=0x16d5a6770 pc=0x1028a20ec
runtime.park_m(0x14000371340)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:4266 +0x164 fp=0x16d5a6800 sp=0x16d5a67b0 pc=0x1028a2454
runtime.mcall()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:241 +0x54 fp=0x16d5a6810 sp=0x16d5a6800 pc=0x1028d2884

goroutine 1 gp=0x140000021c0 m=nil [chan receive]:
runtime.gopark(0x10337b6d8, 0x14000358068, 0xe, 0x7, 0x2)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x140002d0ba0 sp=0x140002d0b70 pc=0x1028cd334
runtime.chanrecv(0x14000358000, 0x140002d0fc8, 0x1)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/chan.go:667 +0x2c8 fp=0x140002d0c20 sp=0x140002d0ba0 pc=0x1028673a8
runtime.chanrecv2(0x10311a46a?, 0x105?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/chan.go:514 +0x14 fp=0x140002d0c50 sp=0x140002d0c20 pc=0x1028670c4
google.golang.org/open2opaque/internal/o2o/rewrite.rewrite({0x10338aa70, 0x1037d07c0}, 0x14000340000)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:487 +0x8c0 fp=0x140002d15f0 sp=0x140002d0c50 pc=0x102f50800
google.golang.org/open2opaque/internal/o2o/rewrite.(*Cmd).RewriteTargets(0x14000226b40, {0x10338aa70, 0x1037d07c0}, {0x14000030200, 0x1, 0x1})
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:350 +0x1a88 fp=0x140002d1bb0 sp=0x140002d15f0 pc=0x102f4f5c8
google.golang.org/open2opaque/internal/o2o/rewrite.(*Cmd).rewrite(0x14000226b40, {0x10338aa70, 0x1037d07c0}, 0x14000205030)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:186 +0x12c fp=0x140002d1c80 sp=0x140002d1bb0 pc=0x102f4dafc
google.golang.org/open2opaque/internal/o2o/rewrite.(*Cmd).Execute(0x14000226b40, {0x10338aa70, 0x1037d07c0}, 0x14000205030, {0x0, 0x0, 0x0})
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:156 +0x44 fp=0x140002d1d30 sp=0x140002d1c80 pc=0x102f4d804
github.com/google/subcommands.(*Commander).Execute(0x1400022d900, {0x10338aa70, 0x1037d07c0}, {0x0, 0x0, 0x0})
	/Users/kieran.gorman/go/pkg/mod/github.com/google/subcommands@v1.2.0/subcommands.go:209 +0x3b0 fp=0x140002d1e90 sp=0x140002d1d30 pc=0x102be02e0
main.main()
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/open2opaque.go:70 +0x2d4 fp=0x140002d1f60 sp=0x140002d1e90 pc=0x10304c554
runtime.main()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:285 +0x218 fp=0x140002d1fd0 sp=0x140002d1f60 pc=0x10289b858
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x140002d1fd0 sp=0x140002d1fd0 pc=0x1028d4c84

goroutine 2 gp=0x14000002700 m=nil [force gc (idle)]:
runtime.gopark(0x10337ba00, 0x1037a6ed0, 0x11, 0xa, 0x1)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000088f70 sp=0x14000088f40 pc=0x1028cd334
runtime.goparkunlock(0x1037a6ed0?, 0x0?, 0x0?, 0x0?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:466 +0x34 fp=0x14000088fa0 sp=0x14000088f70 pc=0x10289bcd4
runtime.forcegchelper()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:373 +0xb0 fp=0x14000088fd0 sp=0x14000088fa0 pc=0x10289bba0
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000088fd0 sp=0x14000088fd0 pc=0x1028d4c84
created by runtime.init.7 in goroutine 1
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:361 +0x24

goroutine 3 gp=0x14000002c40 m=nil [GC sweep wait]:
runtime.gopark(0x10337ba00, 0x1037a7a40, 0xc, 0x9, 0x1)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000089740 sp=0x14000089710 pc=0x1028cd334
runtime.goparkunlock(0x1037a7a40?, 0x0?, 0x0?, 0x0?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:466 +0x34 fp=0x14000089770 sp=0x14000089740 pc=0x10289bcd4
runtime.bgsweep(0x140000a0000)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgcsweep.go:323 +0xd8 fp=0x140000897b0 sp=0x14000089770 pc=0x102883b78
runtime.gcenable.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:212 +0x28 fp=0x140000897d0 sp=0x140000897b0 pc=0x102877e78
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x140000897d0 sp=0x140000897d0 pc=0x1028d4c84
created by runtime.gcenable in goroutine 1
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:212 +0x6c

goroutine 4 gp=0x14000002e00 m=nil [GC scavenge wait]:
runtime.gopark(0x10337ba00, 0x1037a8c00, 0xd, 0xa, 0x2)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000480f20 sp=0x14000480ef0 pc=0x1028cd334
runtime.goparkunlock(0x1037a8c00?, 0xb1?, 0x8d?, 0x0?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:466 +0x34 fp=0x14000480f50 sp=0x14000480f20 pc=0x10289bcd4
runtime.(*scavengerState).park(0x1037a8c00)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgcscavenge.go:425 +0x4c fp=0x14000480f80 sp=0x14000480f50 pc=0x10288131c
runtime.bgscavenge(0x140000a0000)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgcscavenge.go:658 +0x60 fp=0x14000480fb0 sp=0x14000480f80 pc=0x102881870
runtime.gcenable.gowrap2()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:213 +0x28 fp=0x14000480fd0 sp=0x14000480fb0 pc=0x102877e18
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000480fd0 sp=0x14000480fd0 pc=0x1028d4c84
created by runtime.gcenable in goroutine 1
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:213 +0xac

goroutine 5 gp=0x14000003340 m=nil [GOMAXPROCS updater (idle)]:
runtime.gopark(0x10337ba00, 0x1037a6e50, 0x12, 0xa, 0x1)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400008a740 sp=0x1400008a710 pc=0x1028cd334
runtime.goparkunlock(0x1037a6e50?, 0x0?, 0x0?, 0x0?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:466 +0x34 fp=0x1400008a770 sp=0x1400008a740 pc=0x10289bcd4
runtime.updateMaxProcsGoroutine()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:6706 +0xe0 fp=0x1400008a7d0 sp=0x1400008a770 pc=0x1028a73e0
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x1400008a7d0 sp=0x1400008a7d0 pc=0x1028d4c84
created by runtime.defaultGOMAXPROCSUpdateEnable in goroutine 1
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:6694 +0x48

goroutine 6 gp=0x14000003500 m=nil [finalizer wait]:
runtime.gopark(0x10337b730, 0x1037d09b8, 0x10, 0xa, 0x1)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000088580 sp=0x14000088550 pc=0x1028cd334
runtime.runFinalizers()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mfinal.go:210 +0xec fp=0x140000887d0 sp=0x14000088580 pc=0x102876fac
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x140000887d0 sp=0x140000887d0 pc=0x1028d4c84
created by runtime.createfing in goroutine 1
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mfinal.go:172 +0x4c

goroutine 7 gp=0x140000036c0 m=nil [select]:
runtime.gopark(0x10337ba70, 0x0, 0x9, 0x3, 0x1)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x140000b2d60 sp=0x140000b2d30 pc=0x1028cd334
runtime.selectgo(0x140000b2f48, 0x140000b2f00, 0x0?, 0x0, 0x0?, 0x1)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/select.go:351 +0x814 fp=0x140000b2ec0 sp=0x140000b2d60 pc=0x1028ad1d4
github.com/golang/glog.(*fileSink).flushDaemon(0x1037a8b18)
	/Users/kieran.gorman/go/pkg/mod/github.com/golang/glog@v1.2.5/glog_file.go:380 +0xec fp=0x140000b2fb0 sp=0x140000b2ec0 pc=0x102bf17cc
github.com/golang/glog.init.1.gowrap1()
	/Users/kieran.gorman/go/pkg/mod/github.com/golang/glog@v1.2.5/glog_file.go:188 +0x2c fp=0x140000b2fd0 sp=0x140000b2fb0 pc=0x102bf023c
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x140000b2fd0 sp=0x140000b2fd0 pc=0x1028d4c84
created by github.com/golang/glog.init.1 in goroutine 1
	/Users/kieran.gorman/go/pkg/mod/github.com/golang/glog@v1.2.5/glog_file.go:188 +0x26c

goroutine 8 gp=0x14000003a40 m=nil [cleanup wait]:
runtime.gopark(0x10337ba00, 0x1037a9240, 0x2e, 0xa, 0x1)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400008b720 sp=0x1400008b6f0 pc=0x1028cd334
runtime.goparkunlock(0x1037a9248?, 0x0?, 0x0?, 0x0?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:466 +0x34 fp=0x1400008b750 sp=0x1400008b720 pc=0x10289bcd4
runtime.(*cleanupQueue).dequeue(0x1037a90c0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mcleanup.go:439 +0xa0 fp=0x1400008b790 sp=0x1400008b750 pc=0x102874240
runtime.runCleanups()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mcleanup.go:635 +0x60 fp=0x1400008b7d0 sp=0x1400008b790 pc=0x102874a10
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x1400008b7d0 sp=0x1400008b7d0 pc=0x1028d4c84
created by runtime.(*cleanupQueue).createGs in goroutine 1
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mcleanup.go:589 +0xb0

goroutine 9 gp=0x14000003c00 m=nil [IO wait]:
runtime.gopark(0x10337b9e0, 0x103d6e820, 0x2, 0x2, 0x5)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x140000c1400 sp=0x140000c13d0 pc=0x1028cd334
runtime.netpollblock(0x103d6e800, 0x72, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/netpoll.go:575 +0xb8 fp=0x140000c1440 sp=0x140000c1400 pc=0x102894f68
internal/poll.runtime_pollWait(0x103d6e800, 0x72)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/netpoll.go:351 +0x48 fp=0x140000c1460 sp=0x140000c1440 pc=0x1028cc838
internal/poll.(*pollDesc).wait(0x1400012e020, 0x72, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_poll_runtime.go:84 +0x84 fp=0x140000c14b0 sp=0x140000c1460 pc=0x102962544
internal/poll.(*pollDesc).waitRead(0x1400012e020, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_poll_runtime.go:89 +0x38 fp=0x140000c1500 sp=0x140000c14b0 pc=0x1029625c8
internal/poll.(*FD).Accept(0x1400012e000)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/internal/poll/fd_unix.go:613 +0x364 fp=0x140000c1770 sp=0x140000c1500 pc=0x102964374
net.(*netFD).accept(0x1400012e000)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/net/fd_unix.go:161 +0x58 fp=0x140000c1910 sp=0x140000c1770 pc=0x102a866e8
net.(*TCPListener).accept(0x14000136000)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/net/tcpsock_posix.go:159 +0x4c fp=0x140000c1a00 sp=0x140000c1910 pc=0x102aa023c
net.(*TCPListener).Accept(0x14000136000)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/net/tcpsock.go:380 +0x58 fp=0x140000c1ab0 sp=0x140000c1a00 pc=0x102a9e838
net/http.(*onceCloseListener).Accept(0x14000124120)
	<autogenerated>:1 +0x60 fp=0x140000c1b30 sp=0x140000c1ab0 pc=0x102bb6020
net/http.(*Server).Serve(0x14000294000, {0x103388e80, 0x14000136000})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/net/http/server.go:3463 +0x3a4 fp=0x140000c1e10 sp=0x140000c1b30 pc=0x102ba0c24
net/http.(*Server).ListenAndServe(0x14000294000)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/net/http/server.go:3389 +0x140 fp=0x140000c1ed0 sp=0x140000c1e10 pc=0x102ba0770
net/http.ListenAndServe({0x1030ee197, 0xe}, {0x0, 0x0})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/net/http/server.go:3704 +0x118 fp=0x140000c1f30 sp=0x140000c1ed0 pc=0x102ba1b48
google.golang.org/open2opaque/internal/o2o/rewrite.(*Cmd).RewriteTargets.func1()
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:230 +0x60 fp=0x140000c1fd0 sp=0x140000c1f30 pc=0x102f4fcf0
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x140000c1fd0 sp=0x140000c1fd0 pc=0x1028d4c84
created by google.golang.org/open2opaque/internal/o2o/rewrite.(*Cmd).RewriteTargets in goroutine 1
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:222 +0x70c

goroutine 33 gp=0x14000003dc0 m=nil [sync.WaitGroup.Wait]:
runtime.gopark(0x10337ba00, 0x1037afea0, 0x19, 0x5, 0x4)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000a82ca0 sp=0x14000a82c70 pc=0x1028cd334
runtime.goparkunlock(0x1037afea0?, 0xdc?, 0xd0?, 0x140001e9aa0?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:466 +0x34 fp=0x14000a82cd0 sp=0x14000a82ca0 pc=0x10289bcd4
runtime.semacquire1(0x1400030ab38, 0x0, 0x1, 0x0, 0x19)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/sema.go:192 +0x148 fp=0x14000a82d20 sp=0x14000a82cd0 pc=0x1028adc58
sync.runtime_SemacquireWaitGroup(0x1400030ab30?, 0x0?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/sema.go:114 +0x38 fp=0x14000a82d60 sp=0x14000a82d20 pc=0x1028ce9b8
sync.(*WaitGroup).Wait(0x1400030ab30)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/sync/waitgroup.go:206 +0x1b4 fp=0x14000a82dd0 sp=0x14000a82d60 pc=0x1028e27d4
google.golang.org/open2opaque/internal/o2o/rewrite.fixPackageBatch({0x10338aa70, 0x1037d07c0}, {{0x1033859a8, 0x1400030c380}, 0x14000226c80, 0x14000226dc0, 0x0, {{0x0, 0x0}, 0x0, ...}}, ...)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:648 +0x3ec fp=0x14000a82ea0 sp=0x14000a82dd0 pc=0x102f5218c
google.golang.org/open2opaque/internal/o2o/rewrite.rewrite.func1()
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:478 +0x13c fp=0x14000a82fd0 sp=0x14000a82ea0 pc=0x102f5151c
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000a82fd0 sp=0x14000a82fd0 pc=0x1028d4c84
created by google.golang.org/open2opaque/internal/o2o/rewrite.rewrite in goroutine 1
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:471 +0x76c

goroutine 43 gp=0x14000370a80 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000302800, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000360f00 sp=0x14000360ed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x14000360fb0 sp=0x14000360f00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x14000360fd0 sp=0x14000360fb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000360fd0 sp=0x14000360fd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 44 gp=0x14000370c40 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000302a00, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000361700 sp=0x140003616d0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x140003617b0 sp=0x14000361700 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x140003617d0 sp=0x140003617b0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x140003617d0 sp=0x140003617d0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 45 gp=0x14000370e00 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000302c00, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000a84f00 sp=0x14000a84ed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x14000a84fb0 sp=0x14000a84f00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x14000a84fd0 sp=0x14000a84fb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000a84fd0 sp=0x14000a84fd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 46 gp=0x14000370fc0 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000302e00, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400035a700 sp=0x1400035a6d0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x1400035a7b0 sp=0x1400035a700 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x1400035a7d0 sp=0x1400035a7b0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x1400035a7d0 sp=0x1400035a7d0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 47 gp=0x14000371180 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000303000, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400035af00 sp=0x1400035aed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x1400035afb0 sp=0x1400035af00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x1400035afd0 sp=0x1400035afb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x1400035afd0 sp=0x1400035afd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 48 gp=0x14000371340 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000303200, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400035b700 sp=0x1400035b6d0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x1400035b7b0 sp=0x1400035b700 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x1400035b7d0 sp=0x1400035b7b0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x1400035b7d0 sp=0x1400035b7d0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 49 gp=0x14000371500 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000303400, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000a80f00 sp=0x14000a80ed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x14000a80fb0 sp=0x14000a80f00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x14000a80fd0 sp=0x14000a80fb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000a80fd0 sp=0x14000a80fd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 50 gp=0x140003716c0 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000303600, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400035c700 sp=0x1400035c6d0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x1400035c7b0 sp=0x1400035c700 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x1400035c7d0 sp=0x1400035c7b0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x1400035c7d0 sp=0x1400035c7d0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 51 gp=0x14000371880 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000303800, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400035cf00 sp=0x1400035ced0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x1400035cfb0 sp=0x1400035cf00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x1400035cfd0 sp=0x1400035cfb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x1400035cfd0 sp=0x1400035cfd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 52 gp=0x14000371a40 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000303a00, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000481f00 sp=0x14000481ed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x14000481fb0 sp=0x14000481f00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x14000481fd0 sp=0x14000481fb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000481fd0 sp=0x14000481fd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 53 gp=0x14000371c00 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000303c00, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000114f00 sp=0x14000114ed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x14000114fb0 sp=0x14000114f00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x14000114fd0 sp=0x14000114fb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000114fd0 sp=0x14000114fd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 54 gp=0x14000371dc0 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x14000303e00, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000564700 sp=0x140005646d0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x140005647b0 sp=0x14000564700 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x140005647d0 sp=0x140005647b0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x140005647d0 sp=0x140005647d0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 55 gp=0x140003fc000 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x140003fe000, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x140000b7f00 sp=0x140000b7ed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x140000b7fb0 sp=0x140000b7f00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x140000b7fd0 sp=0x140000b7fb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x140000b7fd0 sp=0x140000b7fd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 56 gp=0x140003fc1c0 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x140003fe200, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400010ef00 sp=0x1400010eed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x1400010efb0 sp=0x1400010ef00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x1400010efd0 sp=0x1400010efb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x1400010efd0 sp=0x1400010efd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 57 gp=0x140003fc380 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x140003fe400, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000565f00 sp=0x14000565ed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x14000565fb0 sp=0x14000565f00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x14000565fd0 sp=0x14000565fb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000565fd0 sp=0x14000565fd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 58 gp=0x140003fc540 m=nil [GC worker (idle)]:
runtime.gopark(0x10337b750, 0x140003fe600, 0x1c, 0xa, 0x0)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x14000482f00 sp=0x14000482ed0 pc=0x1028cd334
runtime.gcBgMarkWorker(0x14000358310)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1463 +0xdc fp=0x14000482fb0 sp=0x14000482f00 pc=0x102879b6c
runtime.gcBgMarkStartWorkers.gowrap1()
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x28 fp=0x14000482fd0 sp=0x14000482fb0 pc=0x102879a38
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000482fd0 sp=0x14000482fd0 pc=0x1028d4c84
created by runtime.gcBgMarkStartWorkers in goroutine 40
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/mgc.go:1373 +0x6c

goroutine 258 gp=0x1400029ddc0 m=nil [chan send]:
runtime.gopark(0x10337b6d8, 0x14000b004c8, 0xf, 0x6, 0x2)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/proc.go:460 +0xd4 fp=0x1400093d410 sp=0x1400093d3e0 pc=0x1028cd334
runtime.chansend(0x14000b00460, 0x1400093d8f8, 0x1, 0x1400081aa40?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/chan.go:283 +0x1fc fp=0x1400093d480 sp=0x1400093d410 pc=0x10286686c
runtime.chansend1(0x14000793f10?, 0x2?)
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/chan.go:161 +0x18 fp=0x1400093d4b0 sp=0x1400093d480 pc=0x102866658
google.golang.org/open2opaque/internal/o2o/loader.(*BlazeLoader).LoadPackages(0x1400030c380, {0x10338aa70, 0x1037d07c0}, {0x14000b0e0a8, 0x1, 0x1}, 0x14000b00460)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/loader/loaderblaze.go:114 +0xc08 fp=0x1400093da00 sp=0x1400093d4b0 pc=0x102ef30f8
google.golang.org/open2opaque/internal/o2o/loader.LoadOne({0x10338aa70, 0x1037d07c0}, {0x1033859a8, 0x1400030c380}, 0x14000830620)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/loader/loader.go:95 +0xdc fp=0x1400093dae0 sp=0x1400093da00 pc=0x102ef226c
google.golang.org/open2opaque/internal/fix.generateOneofBuilderCases(0x1400098c0c0, {0x0, 0x0, 0x0}, 0x140009881e0, 0x14000340510)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/fix/oneof.go:66 +0x214 fp=0x1400093e140 sp=0x1400093dae0 pc=0x102f2aa34
google.golang.org/open2opaque/internal/fix.updateBuilderElements(0x1400098c0c0, 0x140009881e0)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/fix/build.go:174 +0xb00 fp=0x1400093e550 sp=0x1400093e140 pc=0x102f0cd00
google.golang.org/open2opaque/internal/fix.buildPost(0x1400098c0c0)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/fix/build.go:27 +0x198 fp=0x1400093e940 sp=0x1400093e550 pc=0x102f0b0a8
google.golang.org/open2opaque/internal/fix.makeApplyFn.func1(0x14000b031b0)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/fix/fix.go:520 +0xd8 fp=0x1400093e9f0 sp=0x1400093e940 pc=0x102f19718
github.com/dave/dst/dstutil.(*application).apply(0x14000b031a0, {0x103382b20, 0x14000520380}, {0x1030e93c2, 0x7}, 0x14000b031e8, {0x103382a40, 0x140005203f0})
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:431 +0x36a0 fp=0x1400093f740 sp=0x1400093e9f0 pc=0x102d6f5f0
github.com/dave/dst/dstutil.(*application).applyList(0x14000b031a0, {0x103382b20, 0x14000520380}, {0x1030e93c2, 0x7})
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:461 +0x1f4 fp=0x1400093f870 sp=0x1400093f740 pc=0x102d6fd84
github.com/dave/dst/dstutil.(*application).apply(0x14000b031a0, {0x1033829e0, 0x1400022da80}, {0x1030e6312, 0x4}, 0x14000b031e8, {0x103382b20, 0x14000520380})
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:335 +0x1e2c fp=0x140009405c0 sp=0x1400093f870 pc=0x102d6dd7c
github.com/dave/dst/dstutil.(*application).applyList(0x14000b031a0, {0x1033829e0, 0x1400022da80}, {0x1030e6312, 0x4})
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:461 +0x1f4 fp=0x140009406f0 sp=0x140009405c0 pc=0x102d6fd84
github.com/dave/dst/dstutil.(*application).apply(0x14000b031a0, {0x103382ae0, 0x14000198870}, {0x1030e6272, 0x4}, 0x0, {0x1033829e0, 0x1400022da80})
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:341 +0x2030 fp=0x14000941440 sp=0x140009406f0 pc=0x102d6df80
github.com/dave/dst/dstutil.(*application).apply(0x14000b031a0, {0x103382be0, 0x1400013c4e0}, {0x1030e701d, 0x5}, 0x14000b031e8, {0x103382ae0, 0x14000198870})
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:407 +0x2e14 fp=0x14000942190 sp=0x14000941440 pc=0x102d6ed64
github.com/dave/dst/dstutil.(*application).applyList(0x14000b031a0, {0x103382be0, 0x1400013c4e0}, {0x1030e701d, 0x5})
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:461 +0x1f4 fp=0x140009422c0 sp=0x14000942190 pc=0x102d6fd84
github.com/dave/dst/dstutil.(*application).apply(0x14000b031a0, {0x1033847e0, 0x14000b04380}, {0x1030e630e, 0x4}, 0x0, {0x103382be0, 0x1400013c4e0})
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:412 +0x2b74 fp=0x14000943010 sp=0x140009422c0 pc=0x102d6eac4
github.com/dave/dst/dstutil.Apply({0x103382be0, 0x1400013c4e0}, 0x0, 0x140007a21c8)
	/Users/kieran.gorman/go/pkg/mod/github.com/dave/dst@v0.27.3/dstutil/rewrite.go:54 +0x1c8 fp=0x140009430f0 sp=0x14000943010 pc=0x102d6b148
google.golang.org/open2opaque/internal/fix.(*ConfiguredPackage).Fix(0x14000943c60)
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/fix/fix.go:335 +0xfc0 fp=0x14000943800 sp=0x140009430f0 pc=0x102f175c0
google.golang.org/open2opaque/internal/o2o/rewrite.fixPackage({0x10338a990, 0x140002487b0}, {{0x1033859a8, 0x1400030c380}, 0x14000226c80, 0x14000226dc0, 0x0, {{0x1033859a8, 0x1400030c380}, 0x140007ecde0, ...}})
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:661 +0x94 fp=0x14000943c30 sp=0x14000943800 pc=0x102f52714
google.golang.org/open2opaque/internal/o2o/rewrite.fixPackageBatch.func2()
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:636 +0x200 fp=0x14000943fd0 sp=0x14000943c30 pc=0x102f523c0
runtime.goexit({})
	/opt/homebrew/Cellar/go/1.25.1/libexec/src/runtime/asm_arm64.s:1268 +0x4 fp=0x14000943fd0 sp=0x14000943fd0 pc=0x1028d4c84
created by google.golang.org/open2opaque/internal/o2o/rewrite.fixPackageBatch in goroutine 33
	/Users/kieran.gorman/go/pkg/mod/google.golang.org/open2opaque@v0.1.7/internal/o2o/rewrite/rewrite.go:619 +0x3dc

r0      0x104
r1      0x0
r2      0x1b00
r3      0x0
r4      0x0
r5      0xa0
r6      0x0
r7      0x0
r8      0x16d5a63d8
r9      0x0
r10     0x0
r11     0x2
r12     0x2
r13     0x0
r14     0x0
r15     0x0
r16     0x131
r17     0x1f6a81500
r18     0x0
r19     0x1037aa268
r20     0x1037aa2a8
r21     0x1f5a2a220
r22     0x0
r23     0x0
r24     0x1b00
r25     0x1b01
r26     0x1c00
r27     0x103784000
r28     0x1037a9280
r29     0x16d5a6450
lr      0x187ad009c
sp      0x16d5a63c0
pc      0x187a913cc
fault   0x187a913cc
```
