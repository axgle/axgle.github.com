<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/syscall/types_linux.c</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/syscall/types_linux.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Input to godefs.  See also mkerrors.sh and mkall.sh
 */

#define _LARGEFILE_SOURCE
#define _LARGEFILE64_SOURCE
#define _FILE_OFFSET_BITS 64
#define _GNU_SOURCE

#include &lt;dirent.h&gt;
#include &lt;fcntl.h&gt;
#include &lt;linux/user.h&gt;
#include &lt;netinet/in.h&gt;
#include &lt;netinet/tcp.h&gt;
#include &lt;signal.h&gt;
#include &lt;stdio.h&gt;
#include &lt;sys/epoll.h&gt;
#include &lt;sys/mman.h&gt;
#include &lt;sys/mount.h&gt;
#include &lt;sys/param.h&gt;
#include &lt;sys/ptrace.h&gt;
#include &lt;sys/resource.h&gt;
#include &lt;sys/select.h&gt;
#include &lt;sys/signal.h&gt;
#include &lt;sys/stat.h&gt;
#include &lt;sys/statfs.h&gt;
#include &lt;sys/sysinfo.h&gt;
#include &lt;sys/time.h&gt;
#include &lt;sys/times.h&gt;
#include &lt;sys/timex.h&gt;
#include &lt;sys/types.h&gt;
#include &lt;sys/un.h&gt;
#include &lt;sys/utsname.h&gt;
#include &lt;sys/wait.h&gt;
#include &lt;time.h&gt;
#include &lt;unistd.h&gt;
#include &lt;ustat.h&gt;
#include &lt;utime.h&gt;

// Machine characteristics; for internal use.

enum
{
	$sizeofPtr = sizeof(void*),
	$sizeofShort = sizeof(short),
	$sizeofInt = sizeof(int),
	$sizeofLong = sizeof(long),
	$sizeofLongLong = sizeof(long long),
	$PathMax = PATH_MAX,
};

// Basic types

typedef short $_C_short;
typedef int $_C_int;
typedef long $_C_long;
typedef long long $_C_long_long;

// Time

typedef struct timespec $Timespec;
typedef struct timeval $Timeval;
typedef struct timex $Timex;
typedef time_t $Time_t;
typedef struct tms $Tms;
typedef struct utimbuf $Utimbuf;

// Processes

typedef struct rusage $Rusage;
typedef struct rlimit $Rlimit;

typedef gid_t $_Gid_t;

// Files

typedef struct stat $Stat_t;
typedef struct statfs $Statfs_t;

typedef struct dirent $Dirent;

// Sockets

union sockaddr_all {
	struct sockaddr s1;	// this one gets used for fields
	struct sockaddr_in s2;	// these pad it out
	struct sockaddr_in6 s3;
	struct sockaddr_un s4;
};

struct sockaddr_any {
	struct sockaddr addr;
	char pad[sizeof(union sockaddr_all) - sizeof(struct sockaddr)];
};

typedef struct sockaddr_in $RawSockaddrInet4;
typedef struct sockaddr_in6 $RawSockaddrInet6;
typedef struct sockaddr_un $RawSockaddrUnix;
typedef struct sockaddr $RawSockaddr;
typedef struct sockaddr_any $RawSockaddrAny;
typedef socklen_t $_Socklen;
typedef struct linger $Linger;
typedef struct iovec $Iovec;
typedef struct msghdr $Msghdr;
typedef struct cmsghdr $Cmsghdr;

enum {
	$SizeofSockaddrInet4 = sizeof(struct sockaddr_in),
	$SizeofSockaddrInet6 = sizeof(struct sockaddr_in6),
	$SizeofSockaddrAny = sizeof(struct sockaddr_any),
	$SizeofSockaddrUnix = sizeof(struct sockaddr_un),
	$SizeofLinger = sizeof(struct linger),
	$SizeofMsghdr = sizeof(struct msghdr),
	$SizeofCmsghdr = sizeof(struct cmsghdr),
};


// Ptrace

// Register structures
typedef struct user_regs_struct $PtraceRegs;

// Misc

typedef fd_set $FdSet;
typedef struct sysinfo $Sysinfo_t;
typedef struct utsname $Utsname;
typedef struct ustat $Ustat_t;

// The real epoll_event is a union, and godefs doesn&#39;t handle it well.
struct my_epoll_event {
	uint32_t events;
	int32_t fd;
	int32_t pad;
};

typedef struct my_epoll_event $EpollEvent;
</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->
