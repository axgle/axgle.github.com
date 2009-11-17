<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/syscall_linux_386.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/syscall_linux_386.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package syscall

<a id="L7"></a>import &#34;unsafe&#34;

<a id="L9"></a>func Getpagesize() int { return 4096 }

<a id="L11"></a>func TimespecToNsec(ts Timespec) int64 { return int64(ts.Sec)*1e9 + int64(ts.Nsec) }

<a id="L13"></a>func NsecToTimespec(nsec int64) (ts Timespec) {
    <a id="L14"></a>ts.Sec = int32(nsec / 1e9);
    <a id="L15"></a>ts.Nsec = int32(nsec % 1e9);
    <a id="L16"></a>return;
<a id="L17"></a>}

<a id="L19"></a>func TimevalToNsec(tv Timeval) int64 { return int64(tv.Sec)*1e9 + int64(tv.Usec)*1e3 }

<a id="L21"></a>func NsecToTimeval(nsec int64) (tv Timeval) {
    <a id="L22"></a>nsec += 999; <span class="comment">// round up to microsecond</span>
    <a id="L23"></a>tv.Sec = int32(nsec / 1e9);
    <a id="L24"></a>tv.Usec = int32(nsec % 1e9 / 1e3);
    <a id="L25"></a>return;
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// 64-bit file system and 32-bit uid calls</span>
<a id="L29"></a><span class="comment">// (386 default is 32-bit file system and 16-bit uid).</span>
<a id="L30"></a><span class="comment">//sys	Chown(path string, uid int, gid int) (errno int) = SYS_CHOWN32</span>
<a id="L31"></a><span class="comment">//sys	Fchown(fd int, uid int, gid int) (errno int) = SYS_FCHOWN32</span>
<a id="L32"></a><span class="comment">//sys	Fstat(fd int, stat *Stat_t) (errno int) = SYS_FSTAT64</span>
<a id="L33"></a><span class="comment">//sys	Fstatfs(fd int, buf *Statfs_t) (errno int) = SYS_FSTATFS64</span>
<a id="L34"></a><span class="comment">//sys	Getegid() (egid int) = SYS_GETEGID32</span>
<a id="L35"></a><span class="comment">//sys	Geteuid() (euid int) = SYS_GETEUID32</span>
<a id="L36"></a><span class="comment">//sys	Getgid() (gid int) = SYS_GETGID32</span>
<a id="L37"></a><span class="comment">//sys	Getuid() (uid int) = SYS_GETUID32</span>
<a id="L38"></a><span class="comment">//sys	Ioperm(from int, num int, on int) (errno int)</span>
<a id="L39"></a><span class="comment">//sys	Iopl(level int) (errno int)</span>
<a id="L40"></a><span class="comment">//sys	Lchown(path string, uid int, gid int) (errno int) = SYS_LCHOWN32</span>
<a id="L41"></a><span class="comment">//sys	Lstat(path string, stat *Stat_t) (errno int) = SYS_LSTAT64</span>
<a id="L42"></a><span class="comment">//sys	Setfsgid(gid int) (errno int) = SYS_SETFSGID32</span>
<a id="L43"></a><span class="comment">//sys	Setfsuid(uid int) (errno int) = SYS_SETFSUID32</span>
<a id="L44"></a><span class="comment">//sys	Setgid(gid int) (errno int) = SYS_SETGID32</span>
<a id="L45"></a><span class="comment">//sys	Setregid(rgid int, egid int) (errno int) = SYS_SETREGID32</span>
<a id="L46"></a><span class="comment">//sys	Setresgid(rgid int, egid int, sgid int) (errno int) = SYS_SETRESGID32</span>
<a id="L47"></a><span class="comment">//sys	Setresuid(ruid int, euid int, suid int) (errno int) = SYS_SETRESUID32</span>
<a id="L48"></a><span class="comment">//sys	Setreuid(ruid int, euid int) (errno int) = SYS_SETREUID32</span>
<a id="L49"></a><span class="comment">//sys	Stat(path string, stat *Stat_t) (errno int) = SYS_STAT64</span>
<a id="L50"></a><span class="comment">//sys	Statfs(path string, buf *Statfs_t) (errno int) = SYS_STATFS64</span>
<a id="L51"></a><span class="comment">//sys	SyncFileRange(fd int, off int64, n int64, flags int) (errno int)</span>
<a id="L52"></a><span class="comment">//sys	getgroups(n int, list *_Gid_t) (nn int, errno int) = SYS_GETGROUPS32</span>
<a id="L53"></a><span class="comment">//sys	setgroups(n int, list *_Gid_t) (errno int) = SYS_SETGROUPS32</span>
<a id="L54"></a><span class="comment">//sys	Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, errno int) = SYS__NEWSELECT</span>

<a id="L56"></a><span class="comment">// Underlying system call writes to newoffset via pointer.</span>
<a id="L57"></a><span class="comment">// Implemented in assembly to avoid allocation.</span>
<a id="L58"></a>func Seek(fd int, offset int64, whence int) (newoffset int64, errno int)

<a id="L60"></a><span class="comment">// On x86 Linux, all the socket calls go through an extra indirection,</span>
<a id="L61"></a><span class="comment">// I think because the 5-register system call interface can&#39;t handle</span>
<a id="L62"></a><span class="comment">// the 6-argument calls like sendto and recvfrom.  Instead the</span>
<a id="L63"></a><span class="comment">// arguments to the underlying system call are the number below</span>
<a id="L64"></a><span class="comment">// and a pointer to an array of uintptr.  We hide the pointer in the</span>
<a id="L65"></a><span class="comment">// socketcall assembly to avoid allocation on every system call.</span>

<a id="L67"></a>const (
    <a id="L68"></a><span class="comment">// see linux/net.h</span>
    <a id="L69"></a>_SOCKET      = 1;
    <a id="L70"></a>_BIND        = 2;
    <a id="L71"></a>_CONNECT     = 3;
    <a id="L72"></a>_LISTEN      = 4;
    <a id="L73"></a>_ACCEPT      = 5;
    <a id="L74"></a>_GETSOCKNAME = 6;
    <a id="L75"></a>_GETPEERNAME = 7;
    <a id="L76"></a>_SOCKETPAIR  = 8;
    <a id="L77"></a>_SEND        = 9;
    <a id="L78"></a>_RECV        = 10;
    <a id="L79"></a>_SENDTO      = 11;
    <a id="L80"></a>_RECVFROM    = 12;
    <a id="L81"></a>_SHUTDOWN    = 13;
    <a id="L82"></a>_SETSOCKOPT  = 14;
    <a id="L83"></a>_GETSOCKOPT  = 15;
    <a id="L84"></a>_SENDMSG     = 16;
    <a id="L85"></a>_RECVMSG     = 17;
<a id="L86"></a>)

<a id="L88"></a>func socketcall(call int, a0, a1, a2, a3, a4, a5 uintptr) (n int, errno int)

<a id="L90"></a>func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, errno int) {
    <a id="L91"></a>fd, errno = socketcall(_ACCEPT, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0);
    <a id="L92"></a>return;
<a id="L93"></a>}

<a id="L95"></a>func getsockname(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int) {
    <a id="L96"></a>_, errno = socketcall(_GETSOCKNAME, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0);
    <a id="L97"></a>return;
<a id="L98"></a>}

<a id="L100"></a>func getpeername(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int) {
    <a id="L101"></a>_, errno = socketcall(_GETPEERNAME, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0);
    <a id="L102"></a>return;
<a id="L103"></a>}

<a id="L105"></a>func bind(s int, addr uintptr, addrlen _Socklen) (errno int) {
    <a id="L106"></a>_, errno = socketcall(_BIND, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0);
    <a id="L107"></a>return;
<a id="L108"></a>}

<a id="L110"></a>func connect(s int, addr uintptr, addrlen _Socklen) (errno int) {
    <a id="L111"></a>_, errno = socketcall(_CONNECT, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0);
    <a id="L112"></a>return;
<a id="L113"></a>}

<a id="L115"></a>func socket(domain int, typ int, proto int) (fd int, errno int) {
    <a id="L116"></a>fd, errno = socketcall(_SOCKET, uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0);
    <a id="L117"></a>return;
<a id="L118"></a>}

<a id="L120"></a>func setsockopt(s int, level int, name int, val uintptr, vallen int) (errno int) {
    <a id="L121"></a>_, errno = socketcall(_SETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0);
    <a id="L122"></a>return;
<a id="L123"></a>}

<a id="L125"></a>func recvfrom(s int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, errno int) {
    <a id="L126"></a>var base uintptr;
    <a id="L127"></a>if len(p) &gt; 0 {
        <a id="L128"></a>base = uintptr(unsafe.Pointer(&amp;p))
    <a id="L129"></a>}
    <a id="L130"></a>n, errno = socketcall(_RECVFROM, uintptr(s), base, uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)));
    <a id="L131"></a>return;
<a id="L132"></a>}

<a id="L134"></a>func sendto(s int, p []byte, flags int, to uintptr, addrlen _Socklen) (errno int) {
    <a id="L135"></a>var base uintptr;
    <a id="L136"></a>if len(p) &gt; 0 {
        <a id="L137"></a>base = uintptr(unsafe.Pointer(&amp;p))
    <a id="L138"></a>}
    <a id="L139"></a>_, errno = socketcall(_SENDTO, uintptr(s), base, uintptr(len(p)), uintptr(flags), to, uintptr(addrlen));
    <a id="L140"></a>return;
<a id="L141"></a>}

<a id="L143"></a>func Listen(s int, n int) (errno int) {
    <a id="L144"></a>_, errno = socketcall(_LISTEN, uintptr(s), uintptr(n), 0, 0, 0, 0);
    <a id="L145"></a>return;
<a id="L146"></a>}

<a id="L148"></a>func (r *PtraceRegs) PC() uint64 { return uint64(uint32(r.Eip)) }

<a id="L150"></a>func (r *PtraceRegs) SetPC(pc uint64) { r.Eip = int32(pc) }
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
