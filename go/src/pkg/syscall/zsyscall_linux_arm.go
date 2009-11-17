<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zsyscall_linux_arm.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zsyscall_linux_arm.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mksyscall.sh -l32 syscall_linux.go syscall_linux_arm.go</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT</span>

<a id="L4"></a>package syscall

<a id="L6"></a>import &#34;unsafe&#34;

<a id="L8"></a>func pipe(p *[2]_C_int) (errno int) {
    <a id="L9"></a>_, _, e1 := Syscall(SYS_PIPE, uintptr(unsafe.Pointer(p)), 0, 0);
    <a id="L10"></a>errno = int(e1);
    <a id="L11"></a>return;
<a id="L12"></a>}

<a id="L14"></a>func utimes(path string, times *[2]Timeval) (errno int) {
    <a id="L15"></a>_, _, e1 := Syscall(SYS_UTIMES, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(times)), 0);
    <a id="L16"></a>errno = int(e1);
    <a id="L17"></a>return;
<a id="L18"></a>}

<a id="L20"></a>func futimesat(dirfd int, path string, times *[2]Timeval) (errno int) {
    <a id="L21"></a>_, _, e1 := Syscall(SYS_FUTIMESAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(times)));
    <a id="L22"></a>errno = int(e1);
    <a id="L23"></a>return;
<a id="L24"></a>}

<a id="L26"></a>func Getcwd(buf []byte) (n int, errno int) {
    <a id="L27"></a>var _p0 *byte;
    <a id="L28"></a>if len(buf) &gt; 0 {
        <a id="L29"></a>_p0 = &amp;buf[0]
    <a id="L30"></a>}
    <a id="L31"></a>r0, _, e1 := Syscall(SYS_GETCWD, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), 0);
    <a id="L32"></a>n = int(r0);
    <a id="L33"></a>errno = int(e1);
    <a id="L34"></a>return;
<a id="L35"></a>}

<a id="L37"></a>func wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, errno int) {
    <a id="L38"></a>r0, _, e1 := Syscall6(SYS_WAIT4, uintptr(pid), uintptr(unsafe.Pointer(wstatus)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0);
    <a id="L39"></a>wpid = int(r0);
    <a id="L40"></a>errno = int(e1);
    <a id="L41"></a>return;
<a id="L42"></a>}

<a id="L44"></a>func ptrace(request int, pid int, addr uintptr, data uintptr) (errno int) {
    <a id="L45"></a>_, _, e1 := Syscall6(SYS_PTRACE, uintptr(request), uintptr(pid), uintptr(addr), uintptr(data), 0, 0);
    <a id="L46"></a>errno = int(e1);
    <a id="L47"></a>return;
<a id="L48"></a>}

<a id="L50"></a>func Access(path string, mode int) (errno int) {
    <a id="L51"></a>_, _, e1 := Syscall(SYS_ACCESS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
    <a id="L52"></a>errno = int(e1);
    <a id="L53"></a>return;
<a id="L54"></a>}

<a id="L56"></a>func Acct(path string) (errno int) {
    <a id="L57"></a>_, _, e1 := Syscall(SYS_ACCT, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L58"></a>errno = int(e1);
    <a id="L59"></a>return;
<a id="L60"></a>}

<a id="L62"></a>func Adjtimex(buf *Timex) (state int, errno int) {
    <a id="L63"></a>r0, _, e1 := Syscall(SYS_ADJTIMEX, uintptr(unsafe.Pointer(buf)), 0, 0);
    <a id="L64"></a>state = int(r0);
    <a id="L65"></a>errno = int(e1);
    <a id="L66"></a>return;
<a id="L67"></a>}

<a id="L69"></a>func Chdir(path string) (errno int) {
    <a id="L70"></a>_, _, e1 := Syscall(SYS_CHDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L71"></a>errno = int(e1);
    <a id="L72"></a>return;
<a id="L73"></a>}

<a id="L75"></a>func Chmod(path string, mode int) (errno int) {
    <a id="L76"></a>_, _, e1 := Syscall(SYS_CHMOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
    <a id="L77"></a>errno = int(e1);
    <a id="L78"></a>return;
<a id="L79"></a>}

<a id="L81"></a>func Chroot(path string) (errno int) {
    <a id="L82"></a>_, _, e1 := Syscall(SYS_CHROOT, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L83"></a>errno = int(e1);
    <a id="L84"></a>return;
<a id="L85"></a>}

<a id="L87"></a>func Close(fd int) (errno int) {
    <a id="L88"></a>_, _, e1 := Syscall(SYS_CLOSE, uintptr(fd), 0, 0);
    <a id="L89"></a>errno = int(e1);
    <a id="L90"></a>return;
<a id="L91"></a>}

<a id="L93"></a>func Creat(path string, mode int) (fd int, errno int) {
    <a id="L94"></a>r0, _, e1 := Syscall(SYS_CREAT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
    <a id="L95"></a>fd = int(r0);
    <a id="L96"></a>errno = int(e1);
    <a id="L97"></a>return;
<a id="L98"></a>}

<a id="L100"></a>func Dup(oldfd int) (fd int, errno int) {
    <a id="L101"></a>r0, _, e1 := Syscall(SYS_DUP, uintptr(oldfd), 0, 0);
    <a id="L102"></a>fd = int(r0);
    <a id="L103"></a>errno = int(e1);
    <a id="L104"></a>return;
<a id="L105"></a>}

<a id="L107"></a>func Dup2(oldfd int, newfd int) (fd int, errno int) {
    <a id="L108"></a>r0, _, e1 := Syscall(SYS_DUP2, uintptr(oldfd), uintptr(newfd), 0);
    <a id="L109"></a>fd = int(r0);
    <a id="L110"></a>errno = int(e1);
    <a id="L111"></a>return;
<a id="L112"></a>}

<a id="L114"></a>func EpollCreate(size int) (fd int, errno int) {
    <a id="L115"></a>r0, _, e1 := Syscall(SYS_EPOLL_CREATE, uintptr(size), 0, 0);
    <a id="L116"></a>fd = int(r0);
    <a id="L117"></a>errno = int(e1);
    <a id="L118"></a>return;
<a id="L119"></a>}

<a id="L121"></a>func EpollCtl(epfd int, op int, fd int, event *EpollEvent) (errno int) {
    <a id="L122"></a>_, _, e1 := Syscall6(SYS_EPOLL_CTL, uintptr(epfd), uintptr(op), uintptr(fd), uintptr(unsafe.Pointer(event)), 0, 0);
    <a id="L123"></a>errno = int(e1);
    <a id="L124"></a>return;
<a id="L125"></a>}

<a id="L127"></a>func EpollWait(epfd int, events []EpollEvent, msec int) (n int, errno int) {
    <a id="L128"></a>var _p0 *EpollEvent;
    <a id="L129"></a>if len(events) &gt; 0 {
        <a id="L130"></a>_p0 = &amp;events[0]
    <a id="L131"></a>}
    <a id="L132"></a>r0, _, e1 := Syscall6(SYS_EPOLL_WAIT, uintptr(epfd), uintptr(unsafe.Pointer(_p0)), uintptr(len(events)), uintptr(msec), 0, 0);
    <a id="L133"></a>n = int(r0);
    <a id="L134"></a>errno = int(e1);
    <a id="L135"></a>return;
<a id="L136"></a>}

<a id="L138"></a>func Exit(code int) {
    <a id="L139"></a>Syscall(SYS_EXIT_GROUP, uintptr(code), 0, 0);
    <a id="L140"></a>return;
<a id="L141"></a>}

<a id="L143"></a>func Faccessat(dirfd int, path string, mode int, flags int) (errno int) {
    <a id="L144"></a>_, _, e1 := Syscall6(SYS_FACCESSAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(flags), 0, 0);
    <a id="L145"></a>errno = int(e1);
    <a id="L146"></a>return;
<a id="L147"></a>}

<a id="L149"></a>func Fallocate(fd int, mode int, off int64, len int64) (errno int) {
    <a id="L150"></a>_, _, e1 := Syscall6(SYS_FALLOCATE, uintptr(fd), uintptr(mode), uintptr(off), uintptr(off&gt;&gt;32), uintptr(len), uintptr(len&gt;&gt;32));
    <a id="L151"></a>errno = int(e1);
    <a id="L152"></a>return;
<a id="L153"></a>}

<a id="L155"></a>func Fchdir(fd int) (errno int) {
    <a id="L156"></a>_, _, e1 := Syscall(SYS_FCHDIR, uintptr(fd), 0, 0);
    <a id="L157"></a>errno = int(e1);
    <a id="L158"></a>return;
<a id="L159"></a>}

<a id="L161"></a>func Fchmod(fd int, mode int) (errno int) {
    <a id="L162"></a>_, _, e1 := Syscall(SYS_FCHMOD, uintptr(fd), uintptr(mode), 0);
    <a id="L163"></a>errno = int(e1);
    <a id="L164"></a>return;
<a id="L165"></a>}

<a id="L167"></a>func Fchmodat(dirfd int, path string, mode int, flags int) (errno int) {
    <a id="L168"></a>_, _, e1 := Syscall6(SYS_FCHMODAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(flags), 0, 0);
    <a id="L169"></a>errno = int(e1);
    <a id="L170"></a>return;
<a id="L171"></a>}

<a id="L173"></a>func Fchownat(dirfd int, path string, uid int, gid int, flags int) (errno int) {
    <a id="L174"></a>_, _, e1 := Syscall6(SYS_FCHOWNAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid), uintptr(flags), 0);
    <a id="L175"></a>errno = int(e1);
    <a id="L176"></a>return;
<a id="L177"></a>}

<a id="L179"></a>func fcntl(fd int, cmd int, arg int) (val int, errno int) {
    <a id="L180"></a>r0, _, e1 := Syscall(SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg));
    <a id="L181"></a>val = int(r0);
    <a id="L182"></a>errno = int(e1);
    <a id="L183"></a>return;
<a id="L184"></a>}

<a id="L186"></a>func Fdatasync(fd int) (errno int) {
    <a id="L187"></a>_, _, e1 := Syscall(SYS_FDATASYNC, uintptr(fd), 0, 0);
    <a id="L188"></a>errno = int(e1);
    <a id="L189"></a>return;
<a id="L190"></a>}

<a id="L192"></a>func Fsync(fd int) (errno int) {
    <a id="L193"></a>_, _, e1 := Syscall(SYS_FSYNC, uintptr(fd), 0, 0);
    <a id="L194"></a>errno = int(e1);
    <a id="L195"></a>return;
<a id="L196"></a>}

<a id="L198"></a>func Ftruncate(fd int, length int64) (errno int) {
    <a id="L199"></a>_, _, e1 := Syscall(SYS_FTRUNCATE, uintptr(fd), uintptr(length), uintptr(length&gt;&gt;32));
    <a id="L200"></a>errno = int(e1);
    <a id="L201"></a>return;
<a id="L202"></a>}

<a id="L204"></a>func Getdents(fd int, buf []byte) (n int, errno int) {
    <a id="L205"></a>var _p0 *byte;
    <a id="L206"></a>if len(buf) &gt; 0 {
        <a id="L207"></a>_p0 = &amp;buf[0]
    <a id="L208"></a>}
    <a id="L209"></a>r0, _, e1 := Syscall(SYS_GETDENTS64, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)));
    <a id="L210"></a>n = int(r0);
    <a id="L211"></a>errno = int(e1);
    <a id="L212"></a>return;
<a id="L213"></a>}

<a id="L215"></a>func Getpgid(pid int) (pgid int, errno int) {
    <a id="L216"></a>r0, _, e1 := Syscall(SYS_GETPGID, uintptr(pid), 0, 0);
    <a id="L217"></a>pgid = int(r0);
    <a id="L218"></a>errno = int(e1);
    <a id="L219"></a>return;
<a id="L220"></a>}

<a id="L222"></a>func Getpgrp() (pid int) {
    <a id="L223"></a>r0, _, _ := Syscall(SYS_GETPGRP, 0, 0, 0);
    <a id="L224"></a>pid = int(r0);
    <a id="L225"></a>return;
<a id="L226"></a>}

<a id="L228"></a>func Getpid() (pid int) {
    <a id="L229"></a>r0, _, _ := Syscall(SYS_GETPID, 0, 0, 0);
    <a id="L230"></a>pid = int(r0);
    <a id="L231"></a>return;
<a id="L232"></a>}

<a id="L234"></a>func Getppid() (ppid int) {
    <a id="L235"></a>r0, _, _ := Syscall(SYS_GETPPID, 0, 0, 0);
    <a id="L236"></a>ppid = int(r0);
    <a id="L237"></a>return;
<a id="L238"></a>}

<a id="L240"></a>func Getrlimit(resource int, rlim *Rlimit) (errno int) {
    <a id="L241"></a>_, _, e1 := Syscall(SYS_GETRLIMIT, uintptr(resource), uintptr(unsafe.Pointer(rlim)), 0);
    <a id="L242"></a>errno = int(e1);
    <a id="L243"></a>return;
<a id="L244"></a>}

<a id="L246"></a>func Getrusage(who int, rusage *Rusage) (errno int) {
    <a id="L247"></a>_, _, e1 := Syscall(SYS_GETRUSAGE, uintptr(who), uintptr(unsafe.Pointer(rusage)), 0);
    <a id="L248"></a>errno = int(e1);
    <a id="L249"></a>return;
<a id="L250"></a>}

<a id="L252"></a>func Gettid() (tid int) {
    <a id="L253"></a>r0, _, _ := Syscall(SYS_GETTID, 0, 0, 0);
    <a id="L254"></a>tid = int(r0);
    <a id="L255"></a>return;
<a id="L256"></a>}

<a id="L258"></a>func Gettimeofday(tv *Timeval) (errno int) {
    <a id="L259"></a>_, _, e1 := Syscall(SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0);
    <a id="L260"></a>errno = int(e1);
    <a id="L261"></a>return;
<a id="L262"></a>}

<a id="L264"></a>func Kill(pid int, sig int) (errno int) {
    <a id="L265"></a>_, _, e1 := Syscall(SYS_KILL, uintptr(pid), uintptr(sig), 0);
    <a id="L266"></a>errno = int(e1);
    <a id="L267"></a>return;
<a id="L268"></a>}

<a id="L270"></a>func Klogctl(typ int, buf []byte) (n int, errno int) {
    <a id="L271"></a>var _p0 *byte;
    <a id="L272"></a>if len(buf) &gt; 0 {
        <a id="L273"></a>_p0 = &amp;buf[0]
    <a id="L274"></a>}
    <a id="L275"></a>r0, _, e1 := Syscall(SYS_SYSLOG, uintptr(typ), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)));
    <a id="L276"></a>n = int(r0);
    <a id="L277"></a>errno = int(e1);
    <a id="L278"></a>return;
<a id="L279"></a>}

<a id="L281"></a>func Link(oldpath string, newpath string) (errno int) {
    <a id="L282"></a>_, _, e1 := Syscall(SYS_LINK, uintptr(unsafe.Pointer(StringBytePtr(oldpath))), uintptr(unsafe.Pointer(StringBytePtr(newpath))), 0);
    <a id="L283"></a>errno = int(e1);
    <a id="L284"></a>return;
<a id="L285"></a>}

<a id="L287"></a>func Mkdir(path string, mode int) (errno int) {
    <a id="L288"></a>_, _, e1 := Syscall(SYS_MKDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
    <a id="L289"></a>errno = int(e1);
    <a id="L290"></a>return;
<a id="L291"></a>}

<a id="L293"></a>func Mkdirat(dirfd int, path string, mode int) (errno int) {
    <a id="L294"></a>_, _, e1 := Syscall(SYS_MKDIRAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode));
    <a id="L295"></a>errno = int(e1);
    <a id="L296"></a>return;
<a id="L297"></a>}

<a id="L299"></a>func Mknod(path string, mode int, dev int) (errno int) {
    <a id="L300"></a>_, _, e1 := Syscall(SYS_MKNOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(dev));
    <a id="L301"></a>errno = int(e1);
    <a id="L302"></a>return;
<a id="L303"></a>}

<a id="L305"></a>func Mknodat(dirfd int, path string, mode int, dev int) (errno int) {
    <a id="L306"></a>_, _, e1 := Syscall6(SYS_MKNODAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(dev), 0, 0);
    <a id="L307"></a>errno = int(e1);
    <a id="L308"></a>return;
<a id="L309"></a>}

<a id="L311"></a>func Nanosleep(time *Timespec, leftover *Timespec) (errno int) {
    <a id="L312"></a>_, _, e1 := Syscall(SYS_NANOSLEEP, uintptr(unsafe.Pointer(time)), uintptr(unsafe.Pointer(leftover)), 0);
    <a id="L313"></a>errno = int(e1);
    <a id="L314"></a>return;
<a id="L315"></a>}

<a id="L317"></a>func Open(path string, mode int, perm int) (fd int, errno int) {
    <a id="L318"></a>r0, _, e1 := Syscall(SYS_OPEN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(perm));
    <a id="L319"></a>fd = int(r0);
    <a id="L320"></a>errno = int(e1);
    <a id="L321"></a>return;
<a id="L322"></a>}

<a id="L324"></a>func Openat(dirfd int, path string, flags int, mode int) (fd int, errno int) {
    <a id="L325"></a>r0, _, e1 := Syscall6(SYS_OPENAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), uintptr(mode), 0, 0);
    <a id="L326"></a>fd = int(r0);
    <a id="L327"></a>errno = int(e1);
    <a id="L328"></a>return;
<a id="L329"></a>}

<a id="L331"></a>func Pause() (errno int) {
    <a id="L332"></a>_, _, e1 := Syscall(SYS_PAUSE, 0, 0, 0);
    <a id="L333"></a>errno = int(e1);
    <a id="L334"></a>return;
<a id="L335"></a>}

<a id="L337"></a>func PivotRoot(newroot string, putold string) (errno int) {
    <a id="L338"></a>_, _, e1 := Syscall(SYS_PIVOT_ROOT, uintptr(unsafe.Pointer(StringBytePtr(newroot))), uintptr(unsafe.Pointer(StringBytePtr(putold))), 0);
    <a id="L339"></a>errno = int(e1);
    <a id="L340"></a>return;
<a id="L341"></a>}

<a id="L343"></a>func Pread(fd int, p []byte, offset int64) (n int, errno int) {
    <a id="L344"></a>var _p0 *byte;
    <a id="L345"></a>if len(p) &gt; 0 {
        <a id="L346"></a>_p0 = &amp;p[0]
    <a id="L347"></a>}
    <a id="L348"></a>r0, _, e1 := Syscall6(SYS_PREAD64, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), uintptr(offset&gt;&gt;32), 0);
    <a id="L349"></a>n = int(r0);
    <a id="L350"></a>errno = int(e1);
    <a id="L351"></a>return;
<a id="L352"></a>}

<a id="L354"></a>func Pwrite(fd int, p []byte, offset int64) (n int, errno int) {
    <a id="L355"></a>var _p0 *byte;
    <a id="L356"></a>if len(p) &gt; 0 {
        <a id="L357"></a>_p0 = &amp;p[0]
    <a id="L358"></a>}
    <a id="L359"></a>r0, _, e1 := Syscall6(SYS_PWRITE64, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), uintptr(offset&gt;&gt;32), 0);
    <a id="L360"></a>n = int(r0);
    <a id="L361"></a>errno = int(e1);
    <a id="L362"></a>return;
<a id="L363"></a>}

<a id="L365"></a>func Read(fd int, p []byte) (n int, errno int) {
    <a id="L366"></a>var _p0 *byte;
    <a id="L367"></a>if len(p) &gt; 0 {
        <a id="L368"></a>_p0 = &amp;p[0]
    <a id="L369"></a>}
    <a id="L370"></a>r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)));
    <a id="L371"></a>n = int(r0);
    <a id="L372"></a>errno = int(e1);
    <a id="L373"></a>return;
<a id="L374"></a>}

<a id="L376"></a>func Readlink(path string, buf []byte) (n int, errno int) {
    <a id="L377"></a>var _p0 *byte;
    <a id="L378"></a>if len(buf) &gt; 0 {
        <a id="L379"></a>_p0 = &amp;buf[0]
    <a id="L380"></a>}
    <a id="L381"></a>r0, _, e1 := Syscall(SYS_READLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)));
    <a id="L382"></a>n = int(r0);
    <a id="L383"></a>errno = int(e1);
    <a id="L384"></a>return;
<a id="L385"></a>}

<a id="L387"></a>func Rename(oldpath string, newpath string) (errno int) {
    <a id="L388"></a>_, _, e1 := Syscall(SYS_RENAME, uintptr(unsafe.Pointer(StringBytePtr(oldpath))), uintptr(unsafe.Pointer(StringBytePtr(newpath))), 0);
    <a id="L389"></a>errno = int(e1);
    <a id="L390"></a>return;
<a id="L391"></a>}

<a id="L393"></a>func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (errno int) {
    <a id="L394"></a>_, _, e1 := Syscall6(SYS_RENAMEAT, uintptr(olddirfd), uintptr(unsafe.Pointer(StringBytePtr(oldpath))), uintptr(newdirfd), uintptr(unsafe.Pointer(StringBytePtr(newpath))), 0, 0);
    <a id="L395"></a>errno = int(e1);
    <a id="L396"></a>return;
<a id="L397"></a>}

<a id="L399"></a>func Rmdir(path string) (errno int) {
    <a id="L400"></a>_, _, e1 := Syscall(SYS_RMDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L401"></a>errno = int(e1);
    <a id="L402"></a>return;
<a id="L403"></a>}

<a id="L405"></a>func Setdomainname(p []byte) (errno int) {
    <a id="L406"></a>var _p0 *byte;
    <a id="L407"></a>if len(p) &gt; 0 {
        <a id="L408"></a>_p0 = &amp;p[0]
    <a id="L409"></a>}
    <a id="L410"></a>_, _, e1 := Syscall(SYS_SETDOMAINNAME, uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0);
    <a id="L411"></a>errno = int(e1);
    <a id="L412"></a>return;
<a id="L413"></a>}

<a id="L415"></a>func Sethostname(p []byte) (errno int) {
    <a id="L416"></a>var _p0 *byte;
    <a id="L417"></a>if len(p) &gt; 0 {
        <a id="L418"></a>_p0 = &amp;p[0]
    <a id="L419"></a>}
    <a id="L420"></a>_, _, e1 := Syscall(SYS_SETHOSTNAME, uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), 0);
    <a id="L421"></a>errno = int(e1);
    <a id="L422"></a>return;
<a id="L423"></a>}

<a id="L425"></a>func Setpgid(pid int, pgid int) (errno int) {
    <a id="L426"></a>_, _, e1 := Syscall(SYS_SETPGID, uintptr(pid), uintptr(pgid), 0);
    <a id="L427"></a>errno = int(e1);
    <a id="L428"></a>return;
<a id="L429"></a>}

<a id="L431"></a>func Setrlimit(resource int, rlim *Rlimit) (errno int) {
    <a id="L432"></a>_, _, e1 := Syscall(SYS_SETRLIMIT, uintptr(resource), uintptr(unsafe.Pointer(rlim)), 0);
    <a id="L433"></a>errno = int(e1);
    <a id="L434"></a>return;
<a id="L435"></a>}

<a id="L437"></a>func Setsid() (pid int) {
    <a id="L438"></a>r0, _, _ := Syscall(SYS_SETSID, 0, 0, 0);
    <a id="L439"></a>pid = int(r0);
    <a id="L440"></a>return;
<a id="L441"></a>}

<a id="L443"></a>func Settimeofday(tv *Timeval) (errno int) {
    <a id="L444"></a>_, _, e1 := Syscall(SYS_SETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0);
    <a id="L445"></a>errno = int(e1);
    <a id="L446"></a>return;
<a id="L447"></a>}

<a id="L449"></a>func Setuid(uid int) (errno int) {
    <a id="L450"></a>_, _, e1 := Syscall(SYS_SETUID, uintptr(uid), 0, 0);
    <a id="L451"></a>errno = int(e1);
    <a id="L452"></a>return;
<a id="L453"></a>}

<a id="L455"></a>func Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, errno int) {
    <a id="L456"></a>r0, r1, _ := Syscall6(SYS_SPLICE, uintptr(rfd), uintptr(unsafe.Pointer(roff)), uintptr(wfd), uintptr(unsafe.Pointer(woff)), uintptr(len), uintptr(flags));
    <a id="L457"></a>n = int64(int64(r1)&lt;&lt;32 | int64(r0));
    <a id="L458"></a>return;
<a id="L459"></a>}

<a id="L461"></a>func Symlink(oldpath string, newpath string) (errno int) {
    <a id="L462"></a>_, _, e1 := Syscall(SYS_SYMLINK, uintptr(unsafe.Pointer(StringBytePtr(oldpath))), uintptr(unsafe.Pointer(StringBytePtr(newpath))), 0);
    <a id="L463"></a>errno = int(e1);
    <a id="L464"></a>return;
<a id="L465"></a>}

<a id="L467"></a>func Sync() {
    <a id="L468"></a>Syscall(SYS_SYNC, 0, 0, 0);
    <a id="L469"></a>return;
<a id="L470"></a>}

<a id="L472"></a>func Sysinfo(info *Sysinfo_t) (errno int) {
    <a id="L473"></a>_, _, e1 := Syscall(SYS_SYSINFO, uintptr(unsafe.Pointer(info)), 0, 0);
    <a id="L474"></a>errno = int(e1);
    <a id="L475"></a>return;
<a id="L476"></a>}

<a id="L478"></a>func Tee(rfd int, wfd int, len int, flags int) (n int64, errno int) {
    <a id="L479"></a>r0, r1, _ := Syscall6(SYS_TEE, uintptr(rfd), uintptr(wfd), uintptr(len), uintptr(flags), 0, 0);
    <a id="L480"></a>n = int64(int64(r1)&lt;&lt;32 | int64(r0));
    <a id="L481"></a>return;
<a id="L482"></a>}

<a id="L484"></a>func Tgkill(tgid int, tid int, sig int) (errno int) {
    <a id="L485"></a>_, _, e1 := Syscall(SYS_TGKILL, uintptr(tgid), uintptr(tid), uintptr(sig));
    <a id="L486"></a>errno = int(e1);
    <a id="L487"></a>return;
<a id="L488"></a>}

<a id="L490"></a>func Time(t *Time_t) (tt Time_t, errno int) {
    <a id="L491"></a>r0, _, e1 := Syscall(SYS_TIME, uintptr(unsafe.Pointer(t)), 0, 0);
    <a id="L492"></a>tt = Time_t(r0);
    <a id="L493"></a>errno = int(e1);
    <a id="L494"></a>return;
<a id="L495"></a>}

<a id="L497"></a>func Times(tms *Tms) (ticks uintptr, errno int) {
    <a id="L498"></a>r0, _, e1 := Syscall(SYS_TIMES, uintptr(unsafe.Pointer(tms)), 0, 0);
    <a id="L499"></a>ticks = uintptr(r0);
    <a id="L500"></a>errno = int(e1);
    <a id="L501"></a>return;
<a id="L502"></a>}

<a id="L504"></a>func Truncate(path string, length int64) (errno int) {
    <a id="L505"></a>_, _, e1 := Syscall(SYS_TRUNCATE, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(length), uintptr(length&gt;&gt;32));
    <a id="L506"></a>errno = int(e1);
    <a id="L507"></a>return;
<a id="L508"></a>}

<a id="L510"></a>func Umask(mask int) (oldmask int) {
    <a id="L511"></a>r0, _, _ := Syscall(SYS_UMASK, uintptr(mask), 0, 0);
    <a id="L512"></a>oldmask = int(r0);
    <a id="L513"></a>return;
<a id="L514"></a>}

<a id="L516"></a>func Uname(buf *Utsname) (errno int) {
    <a id="L517"></a>_, _, e1 := Syscall(SYS_UNAME, uintptr(unsafe.Pointer(buf)), 0, 0);
    <a id="L518"></a>errno = int(e1);
    <a id="L519"></a>return;
<a id="L520"></a>}

<a id="L522"></a>func Unlink(path string) (errno int) {
    <a id="L523"></a>_, _, e1 := Syscall(SYS_UNLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L524"></a>errno = int(e1);
    <a id="L525"></a>return;
<a id="L526"></a>}

<a id="L528"></a>func Unlinkat(dirfd int, path string) (errno int) {
    <a id="L529"></a>_, _, e1 := Syscall(SYS_UNLINKAT, uintptr(dirfd), uintptr(unsafe.Pointer(StringBytePtr(path))), 0);
    <a id="L530"></a>errno = int(e1);
    <a id="L531"></a>return;
<a id="L532"></a>}

<a id="L534"></a>func Unshare(flags int) (errno int) {
    <a id="L535"></a>_, _, e1 := Syscall(SYS_UNSHARE, uintptr(flags), 0, 0);
    <a id="L536"></a>errno = int(e1);
    <a id="L537"></a>return;
<a id="L538"></a>}

<a id="L540"></a>func Ustat(dev int, ubuf *Ustat_t) (errno int) {
    <a id="L541"></a>_, _, e1 := Syscall(SYS_USTAT, uintptr(dev), uintptr(unsafe.Pointer(ubuf)), 0);
    <a id="L542"></a>errno = int(e1);
    <a id="L543"></a>return;
<a id="L544"></a>}

<a id="L546"></a>func Utime(path string, buf *Utimbuf) (errno int) {
    <a id="L547"></a>_, _, e1 := Syscall(SYS_UTIME, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(buf)), 0);
    <a id="L548"></a>errno = int(e1);
    <a id="L549"></a>return;
<a id="L550"></a>}

<a id="L552"></a>func Write(fd int, p []byte) (n int, errno int) {
    <a id="L553"></a>var _p0 *byte;
    <a id="L554"></a>if len(p) &gt; 0 {
        <a id="L555"></a>_p0 = &amp;p[0]
    <a id="L556"></a>}
    <a id="L557"></a>r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)));
    <a id="L558"></a>n = int(r0);
    <a id="L559"></a>errno = int(e1);
    <a id="L560"></a>return;
<a id="L561"></a>}

<a id="L563"></a>func exitThread(code int) (errno int) {
    <a id="L564"></a>_, _, e1 := Syscall(SYS_EXIT, uintptr(code), 0, 0);
    <a id="L565"></a>errno = int(e1);
    <a id="L566"></a>return;
<a id="L567"></a>}

<a id="L569"></a>func read(fd int, p *byte, np int) (n int, errno int) {
    <a id="L570"></a>r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(np));
    <a id="L571"></a>n = int(r0);
    <a id="L572"></a>errno = int(e1);
    <a id="L573"></a>return;
<a id="L574"></a>}

<a id="L576"></a>func write(fd int, p *byte, np int) (n int, errno int) {
    <a id="L577"></a>r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(np));
    <a id="L578"></a>n = int(r0);
    <a id="L579"></a>errno = int(e1);
    <a id="L580"></a>return;
<a id="L581"></a>}

<a id="L583"></a>func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, errno int) {
    <a id="L584"></a>r0, _, e1 := Syscall(SYS_ACCEPT, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
    <a id="L585"></a>fd = int(r0);
    <a id="L586"></a>errno = int(e1);
    <a id="L587"></a>return;
<a id="L588"></a>}

<a id="L590"></a>func bind(s int, addr uintptr, addrlen _Socklen) (errno int) {
    <a id="L591"></a>_, _, e1 := Syscall(SYS_BIND, uintptr(s), uintptr(addr), uintptr(addrlen));
    <a id="L592"></a>errno = int(e1);
    <a id="L593"></a>return;
<a id="L594"></a>}

<a id="L596"></a>func connect(s int, addr uintptr, addrlen _Socklen) (errno int) {
    <a id="L597"></a>_, _, e1 := Syscall(SYS_CONNECT, uintptr(s), uintptr(addr), uintptr(addrlen));
    <a id="L598"></a>errno = int(e1);
    <a id="L599"></a>return;
<a id="L600"></a>}

<a id="L602"></a>func getgroups(n int, list *_Gid_t) (nn int, errno int) {
    <a id="L603"></a>r0, _, e1 := Syscall(SYS_GETGROUPS32, uintptr(n), uintptr(unsafe.Pointer(list)), 0);
    <a id="L604"></a>nn = int(r0);
    <a id="L605"></a>errno = int(e1);
    <a id="L606"></a>return;
<a id="L607"></a>}

<a id="L609"></a>func setgroups(n int, list *_Gid_t) (errno int) {
    <a id="L610"></a>_, _, e1 := Syscall(SYS_SETGROUPS32, uintptr(n), uintptr(unsafe.Pointer(list)), 0);
    <a id="L611"></a>errno = int(e1);
    <a id="L612"></a>return;
<a id="L613"></a>}

<a id="L615"></a>func setsockopt(s int, level int, name int, val uintptr, vallen int) (errno int) {
    <a id="L616"></a>_, _, e1 := Syscall6(SYS_SETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0);
    <a id="L617"></a>errno = int(e1);
    <a id="L618"></a>return;
<a id="L619"></a>}

<a id="L621"></a>func socket(domain int, typ int, proto int) (fd int, errno int) {
    <a id="L622"></a>r0, _, e1 := Syscall(SYS_SOCKET, uintptr(domain), uintptr(typ), uintptr(proto));
    <a id="L623"></a>fd = int(r0);
    <a id="L624"></a>errno = int(e1);
    <a id="L625"></a>return;
<a id="L626"></a>}

<a id="L628"></a>func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int) {
    <a id="L629"></a>_, _, e1 := Syscall(SYS_GETPEERNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
    <a id="L630"></a>errno = int(e1);
    <a id="L631"></a>return;
<a id="L632"></a>}

<a id="L634"></a>func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int) {
    <a id="L635"></a>_, _, e1 := Syscall(SYS_GETSOCKNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
    <a id="L636"></a>errno = int(e1);
    <a id="L637"></a>return;
<a id="L638"></a>}

<a id="L640"></a>func Chown(path string, uid int, gid int) (errno int) {
    <a id="L641"></a>_, _, e1 := Syscall(SYS_CHOWN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid));
    <a id="L642"></a>errno = int(e1);
    <a id="L643"></a>return;
<a id="L644"></a>}

<a id="L646"></a>func Fchown(fd int, uid int, gid int) (errno int) {
    <a id="L647"></a>_, _, e1 := Syscall(SYS_FCHOWN, uintptr(fd), uintptr(uid), uintptr(gid));
    <a id="L648"></a>errno = int(e1);
    <a id="L649"></a>return;
<a id="L650"></a>}

<a id="L652"></a>func Fstat(fd int, stat *Stat_t) (errno int) {
    <a id="L653"></a>_, _, e1 := Syscall(SYS_FSTAT, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L654"></a>errno = int(e1);
    <a id="L655"></a>return;
<a id="L656"></a>}

<a id="L658"></a>func Fstatfs(fd int, buf *Statfs_t) (errno int) {
    <a id="L659"></a>_, _, e1 := Syscall(SYS_FSTATFS, uintptr(fd), uintptr(unsafe.Pointer(buf)), 0);
    <a id="L660"></a>errno = int(e1);
    <a id="L661"></a>return;
<a id="L662"></a>}

<a id="L664"></a>func Getegid() (egid int) {
    <a id="L665"></a>r0, _, _ := Syscall(SYS_GETEGID, 0, 0, 0);
    <a id="L666"></a>egid = int(r0);
    <a id="L667"></a>return;
<a id="L668"></a>}

<a id="L670"></a>func Geteuid() (euid int) {
    <a id="L671"></a>r0, _, _ := Syscall(SYS_GETEUID, 0, 0, 0);
    <a id="L672"></a>euid = int(r0);
    <a id="L673"></a>return;
<a id="L674"></a>}

<a id="L676"></a>func Getgid() (gid int) {
    <a id="L677"></a>r0, _, _ := Syscall(SYS_GETGID, 0, 0, 0);
    <a id="L678"></a>gid = int(r0);
    <a id="L679"></a>return;
<a id="L680"></a>}

<a id="L682"></a>func Getuid() (uid int) {
    <a id="L683"></a>r0, _, _ := Syscall(SYS_GETUID, 0, 0, 0);
    <a id="L684"></a>uid = int(r0);
    <a id="L685"></a>return;
<a id="L686"></a>}

<a id="L688"></a>func Lchown(path string, uid int, gid int) (errno int) {
    <a id="L689"></a>_, _, e1 := Syscall(SYS_LCHOWN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid));
    <a id="L690"></a>errno = int(e1);
    <a id="L691"></a>return;
<a id="L692"></a>}

<a id="L694"></a>func Listen(s int, n int) (errno int) {
    <a id="L695"></a>_, _, e1 := Syscall(SYS_LISTEN, uintptr(s), uintptr(n), 0);
    <a id="L696"></a>errno = int(e1);
    <a id="L697"></a>return;
<a id="L698"></a>}

<a id="L700"></a>func Lstat(path string, stat *Stat_t) (errno int) {
    <a id="L701"></a>_, _, e1 := Syscall(SYS_LSTAT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L702"></a>errno = int(e1);
    <a id="L703"></a>return;
<a id="L704"></a>}

<a id="L706"></a>func Seek(fd int, offset int64, whence int) (off int64, errno int) {
    <a id="L707"></a>r0, r1, _ := Syscall6(SYS_LSEEK, uintptr(fd), uintptr(offset), uintptr(offset&gt;&gt;32), uintptr(whence), 0, 0);
    <a id="L708"></a>off = int64(int64(r1)&lt;&lt;32 | int64(r0));
    <a id="L709"></a>return;
<a id="L710"></a>}

<a id="L712"></a>func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, errno int) {
    <a id="L713"></a>r0, _, e1 := Syscall6(SYS__NEWSELECT, uintptr(nfd), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0);
    <a id="L714"></a>n = int(r0);
    <a id="L715"></a>errno = int(e1);
    <a id="L716"></a>return;
<a id="L717"></a>}

<a id="L719"></a>func Setfsgid(gid int) (errno int) {
    <a id="L720"></a>_, _, e1 := Syscall(SYS_SETFSGID, uintptr(gid), 0, 0);
    <a id="L721"></a>errno = int(e1);
    <a id="L722"></a>return;
<a id="L723"></a>}

<a id="L725"></a>func Setfsuid(uid int) (errno int) {
    <a id="L726"></a>_, _, e1 := Syscall(SYS_SETFSUID, uintptr(uid), 0, 0);
    <a id="L727"></a>errno = int(e1);
    <a id="L728"></a>return;
<a id="L729"></a>}

<a id="L731"></a>func Setgid(gid int) (errno int) {
    <a id="L732"></a>_, _, e1 := Syscall(SYS_SETGID, uintptr(gid), 0, 0);
    <a id="L733"></a>errno = int(e1);
    <a id="L734"></a>return;
<a id="L735"></a>}

<a id="L737"></a>func Setregid(rgid int, egid int) (errno int) {
    <a id="L738"></a>_, _, e1 := Syscall(SYS_SETREGID, uintptr(rgid), uintptr(egid), 0);
    <a id="L739"></a>errno = int(e1);
    <a id="L740"></a>return;
<a id="L741"></a>}

<a id="L743"></a>func Setresgid(rgid int, egid int, sgid int) (errno int) {
    <a id="L744"></a>_, _, e1 := Syscall(SYS_SETRESGID, uintptr(rgid), uintptr(egid), uintptr(sgid));
    <a id="L745"></a>errno = int(e1);
    <a id="L746"></a>return;
<a id="L747"></a>}

<a id="L749"></a>func Setresuid(ruid int, euid int, suid int) (errno int) {
    <a id="L750"></a>_, _, e1 := Syscall(SYS_SETRESUID, uintptr(ruid), uintptr(euid), uintptr(suid));
    <a id="L751"></a>errno = int(e1);
    <a id="L752"></a>return;
<a id="L753"></a>}

<a id="L755"></a>func Setreuid(ruid int, euid int) (errno int) {
    <a id="L756"></a>_, _, e1 := Syscall(SYS_SETREUID, uintptr(ruid), uintptr(euid), 0);
    <a id="L757"></a>errno = int(e1);
    <a id="L758"></a>return;
<a id="L759"></a>}

<a id="L761"></a>func Shutdown(fd int, how int) (errno int) {
    <a id="L762"></a>_, _, e1 := Syscall(SYS_SHUTDOWN, uintptr(fd), uintptr(how), 0);
    <a id="L763"></a>errno = int(e1);
    <a id="L764"></a>return;
<a id="L765"></a>}

<a id="L767"></a>func Stat(path string, stat *Stat_t) (errno int) {
    <a id="L768"></a>_, _, e1 := Syscall(SYS_STAT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L769"></a>errno = int(e1);
    <a id="L770"></a>return;
<a id="L771"></a>}

<a id="L773"></a>func Statfs(path string, buf *Statfs_t) (errno int) {
    <a id="L774"></a>_, _, e1 := Syscall(SYS_STATFS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(buf)), 0);
    <a id="L775"></a>errno = int(e1);
    <a id="L776"></a>return;
<a id="L777"></a>}

<a id="L779"></a>func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, errno int) {
    <a id="L780"></a>var _p0 *byte;
    <a id="L781"></a>if len(p) &gt; 0 {
        <a id="L782"></a>_p0 = &amp;p[0]
    <a id="L783"></a>}
    <a id="L784"></a>r0, _, e1 := Syscall6(SYS_RECVFROM, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)));
    <a id="L785"></a>n = int(r0);
    <a id="L786"></a>errno = int(e1);
    <a id="L787"></a>return;
<a id="L788"></a>}

<a id="L790"></a>func sendto(s int, buf []byte, flags int, to uintptr, addrlen _Socklen) (errno int) {
    <a id="L791"></a>var _p0 *byte;
    <a id="L792"></a>if len(buf) &gt; 0 {
        <a id="L793"></a>_p0 = &amp;buf[0]
    <a id="L794"></a>}
    <a id="L795"></a>_, _, e1 := Syscall6(SYS_SENDTO, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen));
    <a id="L796"></a>errno = int(e1);
    <a id="L797"></a>return;
<a id="L798"></a>}
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
