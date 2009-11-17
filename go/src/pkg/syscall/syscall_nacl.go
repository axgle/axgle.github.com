<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/syscall_nacl.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/syscall_nacl.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Native Client system calls.</span>

<a id="L7"></a>package syscall

<a id="L9"></a>const OS = &#34;nacl&#34;

<a id="L11"></a><span class="comment">// Auto-generated</span>

<a id="L13"></a><span class="comment">//sys	Chmod(path string, mode int) (errno int)</span>
<a id="L14"></a><span class="comment">//sys	Clock() (clock int)</span>
<a id="L15"></a><span class="comment">//sys	Close(fd int) (errno int)</span>
<a id="L16"></a><span class="comment">//sys	Exit(code int)</span>
<a id="L17"></a><span class="comment">//sys	Fstat(fd int, stat *Stat_t) (errno int)</span>
<a id="L18"></a><span class="comment">//sys	Getdents(fd int, buf []byte) (n int, errno int)</span>
<a id="L19"></a><span class="comment">//sys	Getpid() (pid int)</span>
<a id="L20"></a><span class="comment">//sys	Gettimeofday(tv *Timeval) (errno int)</span>
<a id="L21"></a><span class="comment">//sys	Open(path string, mode int, perm int) (fd int, errno int)</span>
<a id="L22"></a><span class="comment">//sys	Read(fd int, p []byte) (n int, errno int)</span>
<a id="L23"></a><span class="comment">//sys	read(fd int, buf *byte, nbuf int) (n int, errno int)</span>
<a id="L24"></a><span class="comment">//sys	Stat(path string, stat *Stat_t) (errno int)</span>
<a id="L25"></a><span class="comment">//sys	Write(fd int, p []byte) (n int, errno int)</span>

<a id="L27"></a><span class="comment">//sys	MultimediaInit(subsys int) (errno int)</span>
<a id="L28"></a><span class="comment">//sys	MultimediaShutdown() (errno int)</span>

<a id="L30"></a><span class="comment">//sys	CondCreate() (cv int, errno int)</span>
<a id="L31"></a><span class="comment">//sys	CondWait(cv int, mutex int) (errno int)</span>
<a id="L32"></a><span class="comment">//sys	CondSignal(cv int) (errno int)</span>
<a id="L33"></a><span class="comment">//sys	CondBroadcast(cv int) (errno int)</span>
<a id="L34"></a><span class="comment">//sys	CondTimedWaitAbs(cv int, mutex int, abstime *Timespec) (errno int)</span>
<a id="L35"></a><span class="comment">//sys	MutexCreate() (mutex int, errno int)</span>
<a id="L36"></a><span class="comment">//sys	MutexLock(mutex int) (errno int)</span>
<a id="L37"></a><span class="comment">//sys	MutexUnlock(mutex int) (errno int)</span>
<a id="L38"></a><span class="comment">//sys	MutexTryLock(mutex int) (errno int) = SYS_MUTEX_TRYLOCK</span>
<a id="L39"></a><span class="comment">//sys	SemCreate() (sema int, errno int)</span>
<a id="L40"></a><span class="comment">//sys	SemWait(sema int) (errno int)</span>
<a id="L41"></a><span class="comment">//sys	SemPost(sema int) (errno int)</span>
<a id="L42"></a><span class="comment">//sys	VideoInit(dx int, dy int) (errno int)</span>
<a id="L43"></a><span class="comment">//sys	VideoUpdate(data *uint32) (errno int)</span>
<a id="L44"></a><span class="comment">//sys	VideoPollEvent(ev *byte) (errno int)</span>
<a id="L45"></a><span class="comment">//sys	VideoShutdown() (errno int)</span>
<a id="L46"></a><span class="comment">//sys	AudioInit(fmt int, nreq int, data *int) (errno int)</span>
<a id="L47"></a><span class="comment">//sys	AudioShutdown() (errno int)</span>
<a id="L48"></a><span class="comment">//sys	AudioStream(data *uint16, size *uintptr) (errno int)</span>

<a id="L50"></a><span class="comment">// Hand-written</span>

<a id="L52"></a>func Seek(fd int, offset int64, whence int) (newoffset int64, errno int) {
    <a id="L53"></a><span class="comment">// Offset passed to system call is 32 bits.  Failure of vision by NaCl.</span>
    <a id="L54"></a>if int64(int32(offset)) != offset {
        <a id="L55"></a>return 0, ERANGE
    <a id="L56"></a>}
    <a id="L57"></a>o, _, e := Syscall(SYS_LSEEK, uintptr(fd), uintptr(offset), uintptr(whence));
    <a id="L58"></a>return int64(o), int(e);
<a id="L59"></a>}

<a id="L61"></a><span class="comment">// Sleep by waiting on a condition variable that will never be signaled.</span>
<a id="L62"></a><span class="comment">// TODO(rsc): Replace when NaCl adds a proper sleep system call.</span>
<a id="L63"></a>var tcv, tmu int

<a id="L65"></a>func init() {
    <a id="L66"></a>tmu, _ = MutexCreate();
    <a id="L67"></a>tcv, _ = CondCreate();
<a id="L68"></a>}

<a id="L70"></a>func Sleep(ns int64) (errno int) {
    <a id="L71"></a>ts := NsecToTimespec(ns);
    <a id="L72"></a>var tv Timeval;
    <a id="L73"></a>if errno = Gettimeofday(&amp;tv); errno != 0 {
        <a id="L74"></a>return
    <a id="L75"></a>}
    <a id="L76"></a>ts.Sec += tv.Sec;
    <a id="L77"></a>ts.Nsec += tv.Usec * 1000;
    <a id="L78"></a>switch {
    <a id="L79"></a>case ts.Nsec &gt;= 1e9:
        <a id="L80"></a>ts.Nsec -= 1e9;
        <a id="L81"></a>ts.Sec++;
    <a id="L82"></a>case ts.Nsec &lt;= -1e9:
        <a id="L83"></a>ts.Nsec += 1e9;
        <a id="L84"></a>ts.Sec--;
    <a id="L85"></a>}
    <a id="L86"></a>if errno = MutexLock(tmu); errno != 0 {
        <a id="L87"></a>return
    <a id="L88"></a>}
    <a id="L89"></a>errno = CondTimedWaitAbs(tcv, tmu, &amp;ts);
    <a id="L90"></a>if e := MutexUnlock(tmu); e != 0 &amp;&amp; errno == 0 {
        <a id="L91"></a>errno = e
    <a id="L92"></a>}
    <a id="L93"></a>return;
<a id="L94"></a>}

<a id="L96"></a><span class="comment">// Implemented in NaCl but not here; maybe later:</span>
<a id="L97"></a><span class="comment">//	SYS_IOCTL</span>
<a id="L98"></a><span class="comment">//	SYS_IMC_*</span>
<a id="L99"></a><span class="comment">//	SYS_MMAP ???</span>
<a id="L100"></a><span class="comment">//	SYS_SRPC_*</span>
<a id="L101"></a><span class="comment">//	SYS_SYSCONF</span>

<a id="L103"></a><span class="comment">// Implemented in NaCl but not here; used by runtime instead:</span>
<a id="L104"></a><span class="comment">//	SYS_SYSBRK</span>
<a id="L105"></a><span class="comment">//	SYS_MMAP</span>
<a id="L106"></a><span class="comment">//	SYS_MUNMAP</span>
<a id="L107"></a><span class="comment">//	SYS_THREAD_*</span>
<a id="L108"></a><span class="comment">//	SYS_TLS_*</span>
<a id="L109"></a><span class="comment">//	SYS_SCHED_YIELD</span>

<a id="L111"></a><span class="comment">// Not implemented in NaCl but needed to compile other packages.</span>

<a id="L113"></a>const (
    <a id="L114"></a>SIGTRAP = 5;
<a id="L115"></a>)

<a id="L117"></a>func Pipe(p []int) (errno int) { return ENACL }

<a id="L119"></a>func fcntl(fd, cmd, arg int) (val int, errno int) {
    <a id="L120"></a>return 0, ENACL
<a id="L121"></a>}

<a id="L123"></a>func Pread(fd int, p []byte, offset int64) (n int, errno int) {
    <a id="L124"></a>return 0, ENACL
<a id="L125"></a>}

<a id="L127"></a>func Pwrite(fd int, p []byte, offset int64) (n int, errno int) {
    <a id="L128"></a>return 0, ENACL
<a id="L129"></a>}

<a id="L131"></a>func Mkdir(path string, mode int) (errno int) { return ENACL }

<a id="L133"></a>func Lstat(path string, stat *Stat_t) (errno int) {
    <a id="L134"></a>return ENACL
<a id="L135"></a>}

<a id="L137"></a>func Chdir(path string) (errno int) { return ENACL }

<a id="L139"></a>func Fchdir(fd int) (errno int) { return ENACL }

<a id="L141"></a>func Unlink(path string) (errno int) { return ENACL }

<a id="L143"></a>func Rmdir(path string) (errno int) { return ENACL }

<a id="L145"></a>func Link(oldpath, newpath string) (errno int) {
    <a id="L146"></a>return ENACL
<a id="L147"></a>}

<a id="L149"></a>func Symlink(path, link string) (errno int) { return ENACL }

<a id="L151"></a>func Readlink(path string, buf []byte) (n int, errno int) {
    <a id="L152"></a>return 0, ENACL
<a id="L153"></a>}

<a id="L155"></a>func Fchmod(fd int, mode int) (errno int) { return ENACL }

<a id="L157"></a>func Chown(path string, uid int, gid int) (errno int) {
    <a id="L158"></a>return ENACL
<a id="L159"></a>}

<a id="L161"></a>func Lchown(path string, uid int, gid int) (errno int) {
    <a id="L162"></a>return ENACL
<a id="L163"></a>}

<a id="L165"></a>func Fchown(fd int, uid int, gid int) (errno int) {
    <a id="L166"></a>return ENACL
<a id="L167"></a>}

<a id="L169"></a>func Truncate(name string, size int64) (errno int) {
    <a id="L170"></a>return ENACL
<a id="L171"></a>}

<a id="L173"></a>func Ftruncate(fd int, length int64) (errno int) {
    <a id="L174"></a>return ENACL
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// NaCL doesn&#39;t actually implement Getwd, but it also</span>
<a id="L178"></a><span class="comment">// don&#39;t implement Chdir, so the fallback algorithm</span>
<a id="L179"></a><span class="comment">// fails worse than calling Getwd does.</span>

<a id="L181"></a>const ImplementsGetwd = true

<a id="L183"></a>func Getwd() (wd string, errno int) { return &#34;&#34;, ENACL }

<a id="L185"></a>func Getuid() (uid int) { return -1 }

<a id="L187"></a>func Geteuid() (euid int) { return -1 }

<a id="L189"></a>func Getgid() (gid int) { return -1 }

<a id="L191"></a>func Getegid() (egid int) { return -1 }

<a id="L193"></a>func Getppid() (ppid int) { return -1 }

<a id="L195"></a>func Getgroups() (gids []int, errno int) { return nil, ENACL }

<a id="L197"></a>type Sockaddr interface {
    <a id="L198"></a>sockaddr();
<a id="L199"></a>}

<a id="L201"></a>type SockaddrInet4 struct {
    <a id="L202"></a>Port int;
    <a id="L203"></a>Addr [4]byte;
<a id="L204"></a>}

<a id="L206"></a>func (*SockaddrInet4) sockaddr() {}

<a id="L208"></a>type SockaddrInet6 struct {
    <a id="L209"></a>Port int;
    <a id="L210"></a>Addr [16]byte;
<a id="L211"></a>}

<a id="L213"></a>func (*SockaddrInet6) sockaddr() {}

<a id="L215"></a>type SockaddrUnix struct {
    <a id="L216"></a>Name string;
<a id="L217"></a>}

<a id="L219"></a>func (*SockaddrUnix) sockaddr() {}

<a id="L221"></a>const (
    <a id="L222"></a>AF_INET = 1 + iota;
    <a id="L223"></a>AF_INET6;
    <a id="L224"></a>AF_UNIX;
    <a id="L225"></a>IPPROTO_TCP;
    <a id="L226"></a>SOCK_DGRAM;
    <a id="L227"></a>SOCK_STREAM;
    <a id="L228"></a>SOL_SOCKET;
    <a id="L229"></a>SOMAXCONN;
    <a id="L230"></a>SO_DONTROUTE;
    <a id="L231"></a>SO_KEEPALIVE;
    <a id="L232"></a>SO_LINGER;
    <a id="L233"></a>SO_RCVBUF;
    <a id="L234"></a>SO_REUSEADDR;
    <a id="L235"></a>SO_SNDBUF;
    <a id="L236"></a>TCP_NODELAY;
    <a id="L237"></a>WNOHANG;
    <a id="L238"></a>WSTOPPED;
    <a id="L239"></a>_PTRACE_TRACEME;
<a id="L240"></a>)

<a id="L242"></a>func Accept(fd int) (nfd int, sa Sockaddr, errno int) {
    <a id="L243"></a>return 0, nil, ENACL
<a id="L244"></a>}

<a id="L246"></a>func Getsockname(fd int) (sa Sockaddr, errno int) {
    <a id="L247"></a>return nil, ENACL
<a id="L248"></a>}

<a id="L250"></a>func Getpeername(fd int) (sa Sockaddr, errno int) {
    <a id="L251"></a>return nil, ENACL
<a id="L252"></a>}

<a id="L254"></a>func Bind(fd int, sa Sockaddr) (errno int) { return ENACL }

<a id="L256"></a>func Connect(fd int, sa Sockaddr) (errno int) { return ENACL }

<a id="L258"></a>func Socket(domain, typ, proto int) (fd, errno int) {
    <a id="L259"></a>return 0, ENACL
<a id="L260"></a>}

<a id="L262"></a>func SetsockoptInt(fd, level, opt int, value int) (errno int) {
    <a id="L263"></a>return ENACL
<a id="L264"></a>}

<a id="L266"></a>func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (errno int) {
    <a id="L267"></a>return ENACL
<a id="L268"></a>}

<a id="L270"></a>type Linger struct {
    <a id="L271"></a>Onoff  int32;
    <a id="L272"></a>Linger int32;
<a id="L273"></a>}

<a id="L275"></a>func SetsockoptLinger(fd, level, opt int, l *Linger) (errno int) {
    <a id="L276"></a>return ENACL
<a id="L277"></a>}

<a id="L279"></a>func Listen(s int, n int) (errno int) { return ENACL }

<a id="L281"></a>type Rusage struct {
    <a id="L282"></a>Utime    Timeval;
    <a id="L283"></a>Stime    Timeval;
    <a id="L284"></a>Maxrss   int32;
    <a id="L285"></a>Ixrss    int32;
    <a id="L286"></a>Idrss    int32;
    <a id="L287"></a>Isrss    int32;
    <a id="L288"></a>Minflt   int32;
    <a id="L289"></a>Majflt   int32;
    <a id="L290"></a>Nswap    int32;
    <a id="L291"></a>Inblock  int32;
    <a id="L292"></a>Oublock  int32;
    <a id="L293"></a>Msgsnd   int32;
    <a id="L294"></a>Msgrcv   int32;
    <a id="L295"></a>Nsignals int32;
    <a id="L296"></a>Nvcsw    int32;
    <a id="L297"></a>Nivcsw   int32;
<a id="L298"></a>}

<a id="L300"></a>func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, errno int) {
    <a id="L301"></a>return 0, ENACL
<a id="L302"></a>}

<a id="L304"></a>type WaitStatus uint32

<a id="L306"></a>func (WaitStatus) Exited() bool { return false }

<a id="L308"></a>func (WaitStatus) ExitStatus() int { return -1 }

<a id="L310"></a>func (WaitStatus) Signal() int { return -1 }

<a id="L312"></a>func (WaitStatus) CoreDump() bool { return false }

<a id="L314"></a>func (WaitStatus) Stopped() bool { return false }

<a id="L316"></a>func (WaitStatus) Continued() bool { return false }

<a id="L318"></a>func (WaitStatus) StopSignal() int { return -1 }

<a id="L320"></a>func (WaitStatus) Signaled() bool { return false }

<a id="L322"></a>func (WaitStatus) TrapCause() int { return -1 }
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
