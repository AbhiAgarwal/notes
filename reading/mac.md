That's indeed a useful list. lsof alone covers a lot of ground!
Mac OS X has a pretty nice set of DTrace scripts built in: http://dtrace.org/blogs/brendan/2011/10/10/top-10-dtrace-scr...
The ones I use the most:
* iosnoop: see all disk I/O. especially useful to find disk-chatty/poll-y apps.
* execsnoop: see new processes being spawned.
* opensnoop: see file opens. especially useful for failed file opens that break an app.
* dtruss: see all system calls. get access to the entire OS interaction history of an process (or app).
* errinfo: trace failing system calls. where there is smoke...
* iotop -- who is using disk
There is just a crazy, crazy list of things available, built-in:
  man -k dtrace
