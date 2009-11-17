<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/syscall/mkall.sh</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/syscall/mkall.sh</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
#!/bin/sh
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# The syscall package provides access to the raw system call
# interface of the underlying operating system.  Porting Go to
# a new architecture/operating system combination requires
# some manual effort, though there are tools that automate
# much of the process.  The auto-generated files have names
# beginning with z.
#
# This script runs or (given -n) prints suggested commands to generate z files
# for the current system.  Running those commands is not automatic.
# This script is documentation more than anything else.
#
# * asm_${GOOS}_${GOARCH}.s
#
# This hand-written assembly file implements system call dispatch.
# There are three entry points:
#
# 	func Syscall(trap, a1, a2, a3 uintptr) (r1, r2, err uintptr);
# 	func Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2, err uintptr);
# 	func RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2, err uintptr);
#
# The first and second are the standard ones; they differ only in
# how many arguments can be passed to the kernel.
# The third is for low-level use by the ForkExec wrapper;
# unlike the first two, it does not call into the scheduler to
# let it know that a system call is running.
#
# * syscall_${GOOS}.go
#
# This hand-written Go file implements system calls that need
# special handling and lists &#34;//sys&#34; comments giving prototypes
# for ones that can be auto-generated.  Mksyscall reads those
# comments to generate the stubs.
#
# * syscall_${GOOS}_${GOARCH}.go
#
# Same as syscall_${GOOS}.go except that it contains code specific
# to ${GOOS} on one particular architecture.
#
# * types_${GOOS}.c
#
# This hand-written C file includes standard C headers and then
# creates typedef or enum names beginning with a dollar sign
# (use of $ in variable names is a gcc extension).  The hardest
# part about preparing this file is figuring out which headers to
# include and which symbols need to be #defined to get the
# actual data structures that pass through to the kernel system calls.
# Some C libraries present alternate versions for binary compatibility
# and translate them on the way in and out of system calls, but
# there is almost always a #define that can get the real ones.
# See types_darwin.c and types_linux.c for examples.
#
# * zerror_${GOOS}_${GOARCH}.go
#
# This machine-generated file defines the system&#39;s error numbers,
# error strings, and signal numbers.  The generator is &#34;mkerrors.sh&#34;.
# Usually no arguments are needed, but mkerrors.sh will pass its
# arguments on to godefs.
#
# * zsyscall_${GOOS}_${GOARCH}.go
#
# Generated by mksyscall.sh; see syscall_${GOOS}.go above.
#
# * zsysnum_${GOOS}_${GOARCH}.go
#
# Generated by mksysnum_${GOOS}.
#
# * ztypes_${GOOS}_${GOARCH}.go
#
# Generated by godefs; see types_${GOOS}.c above.

GOOSARCH=&#34;${GOOS}_${GOARCH}&#34;

# defaults
mksyscall=&#34;mksyscall.sh&#34;
mkerrors=&#34;mkerrors.sh&#34;
run=&#34;sh&#34;

case &#34;$1&#34; in
-n)
	run=&#34;cat&#34;
	shift
esac

case &#34;$#&#34; in
0)
	;;
*)
	echo &#39;usage: mkall.sh [-n]&#39; 1&gt;&amp;2
	exit 2
esac

case &#34;$GOOSARCH&#34; in
_* | *_ | _)
	echo &#39;undefined $GOOS_$GOARCH:&#39; &#34;$GOOSARCH&#34; 1&gt;&amp;2
	exit 1
	;;
darwin_386)
	mksyscall=&#34;mksyscall.sh -l32&#34;
	mksysnum=&#34;mksysnum_darwin.sh /home/rsc/pub/xnu-1228/bsd/kern/syscalls.master&#34;
	mktypes=&#34;godefs -gsyscall -f-m32&#34;
	;;
darwin_amd64)
	mksysnum=&#34;mksysnum_darwin.sh /home/rsc/pub/xnu-1228/bsd/kern/syscalls.master&#34;
	mktypes=&#34;godefs -gsyscall -f-m64&#34;
	mkerrors=&#34;mkerrors.sh&#34;
	;;
linux_386)
	mksyscall=&#34;mksyscall.sh -l32&#34;
	mksysnum=&#34;mksysnum_linux.sh /usr/include/asm/unistd_32.h&#34;
	mktypes=&#34;godefs -gsyscall -f-m32&#34;
	;;
linux_amd64)
	mksysnum=&#34;mksysnum_linux.sh /usr/include/asm/unistd_64.h&#34;
	mktypes=&#34;godefs -gsyscall -f-m64&#34;
	;;
nacl_386)
	NACL=&#34;/home/rsc/pub/nacl/native_client&#34;
	NACLRUN=&#34;$NACL/src/trusted/service_runtime&#34;
	NACLSDK=&#34;$NACL/src/third_party/nacl_sdk/linux/sdk/nacl-sdk/nacl&#34;
	mksyscall=&#34;mksyscall.sh -l32&#34;
	mksysnum=&#34;mksysnum_nacl.sh $NACLRUN/include/bits/nacl_syscalls.h&#34;
	mktypes=&#34;godefs -gsyscall -f-m32 -f-I$NACLSDK/include -f-I$NACL&#34;
	mkerrors=&#34;mkerrors_nacl.sh $NACLRUN/include/sys/errno.h&#34;
	;;
linux_arm)
	ARM=&#34;/home/kaib/public/linux-2.6.28&#34;
	mksyscall=&#34;mksyscall.sh -l32&#34;
	mksysnum=&#34;mksysnum_linux.sh $ARM/arch/arm/include/asm/unistd.h&#34;
//	mktypes=&#34;godefs -gsyscall -carm-gcc -f-I$ARM/arch/arm/include -f-I$ARM/include -f-D__deprecated=&#39;&#39; -f-I$ARM/arch/arm/mach-at91/include -f-DCONFIG_ARCH_AT91SAM9260 &#34;
	mktypes=&#34;godefs -gsyscall -carm-gcc&#34;
	mkerrors=&#34;mkerrors.sh&#34;
	;;
*)
	echo &#39;unrecognized $GOOS_$GOARCH: &#39; &#34;$GOOSARCH&#34; 1&gt;&amp;2
	exit 1
	;;
esac

(
	echo &#34;$mkerrors |gofmt &gt;zerrors_$GOOSARCH.go&#34;
	echo &#34;$mksyscall syscall_$GOOS.go syscall_$GOOSARCH.go |gofmt &gt;zsyscall_$GOOSARCH.go&#34;
	echo &#34;$mksysnum |gofmt &gt;zsysnum_$GOOSARCH.go&#34;
	echo &#34;$mktypes types_$GOOS.c |gofmt &gt;ztypes_$GOOSARCH.go&#34;
) | $run
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
