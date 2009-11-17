<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/zsyscall_darwin_amd64.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/syscall/zsyscall_darwin_amd64.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// mksyscall.sh syscall_darwin.go syscall_darwin_amd64.go</span>
<a id="L2"></a><span class="comment">// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT</span>

<a id="L4"></a>package syscall

<a id="L6"></a>import &#34;unsafe&#34;

<a id="L8"></a>func getgroups(ngid int, gid *_Gid_t) (n int, errno int) {
    <a id="L9"></a>r0, _, e1 := Syscall(SYS_GETGROUPS, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0);
    <a id="L10"></a>n = int(r0);
    <a id="L11"></a>errno = int(e1);
    <a id="L12"></a>return;
<a id="L13"></a>}

<a id="L15"></a>func setgroups(ngid int, gid *_Gid_t) (errno int) {
    <a id="L16"></a>_, _, e1 := Syscall(SYS_SETGROUPS, uintptr(ngid), uintptr(unsafe.Pointer(gid)), 0);
    <a id="L17"></a>errno = int(e1);
    <a id="L18"></a>return;
<a id="L19"></a>}

<a id="L21"></a>func wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, errno int) {
    <a id="L22"></a>r0, _, e1 := Syscall6(SYS_WAIT4, uintptr(pid), uintptr(unsafe.Pointer(wstatus)), uintptr(options), uintptr(unsafe.Pointer(rusage)), 0, 0);
    <a id="L23"></a>wpid = int(r0);
    <a id="L24"></a>errno = int(e1);
    <a id="L25"></a>return;
<a id="L26"></a>}

<a id="L28"></a>func pipe() (r int, w int, errno int) {
    <a id="L29"></a>r0, r1, e1 := Syscall(SYS_PIPE, 0, 0, 0);
    <a id="L30"></a>r = int(r0);
    <a id="L31"></a>w = int(r1);
    <a id="L32"></a>errno = int(e1);
    <a id="L33"></a>return;
<a id="L34"></a>}

<a id="L36"></a>func accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, errno int) {
    <a id="L37"></a>r0, _, e1 := Syscall(SYS_ACCEPT, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
    <a id="L38"></a>fd = int(r0);
    <a id="L39"></a>errno = int(e1);
    <a id="L40"></a>return;
<a id="L41"></a>}

<a id="L43"></a>func bind(s int, addr uintptr, addrlen _Socklen) (errno int) {
    <a id="L44"></a>_, _, e1 := Syscall(SYS_BIND, uintptr(s), uintptr(addr), uintptr(addrlen));
    <a id="L45"></a>errno = int(e1);
    <a id="L46"></a>return;
<a id="L47"></a>}

<a id="L49"></a>func connect(s int, addr uintptr, addrlen _Socklen) (errno int) {
    <a id="L50"></a>_, _, e1 := Syscall(SYS_CONNECT, uintptr(s), uintptr(addr), uintptr(addrlen));
    <a id="L51"></a>errno = int(e1);
    <a id="L52"></a>return;
<a id="L53"></a>}

<a id="L55"></a>func socket(domain int, typ int, proto int) (fd int, errno int) {
    <a id="L56"></a>r0, _, e1 := Syscall(SYS_SOCKET, uintptr(domain), uintptr(typ), uintptr(proto));
    <a id="L57"></a>fd = int(r0);
    <a id="L58"></a>errno = int(e1);
    <a id="L59"></a>return;
<a id="L60"></a>}

<a id="L62"></a>func setsockopt(s int, level int, name int, val uintptr, vallen int) (errno int) {
    <a id="L63"></a>_, _, e1 := Syscall6(SYS_SETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(vallen), 0);
    <a id="L64"></a>errno = int(e1);
    <a id="L65"></a>return;
<a id="L66"></a>}

<a id="L68"></a>func getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int) {
    <a id="L69"></a>_, _, e1 := Syscall(SYS_GETPEERNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
    <a id="L70"></a>errno = int(e1);
    <a id="L71"></a>return;
<a id="L72"></a>}

<a id="L74"></a>func getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int) {
    <a id="L75"></a>_, _, e1 := Syscall(SYS_GETSOCKNAME, uintptr(fd), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)));
    <a id="L76"></a>errno = int(e1);
    <a id="L77"></a>return;
<a id="L78"></a>}

<a id="L80"></a>func recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, errno int) {
    <a id="L81"></a>var _p0 *byte;
    <a id="L82"></a>if len(p) &gt; 0 {
        <a id="L83"></a>_p0 = &amp;p[0]
    <a id="L84"></a>}
    <a id="L85"></a>r0, _, e1 := Syscall6(SYS_RECVFROM, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)));
    <a id="L86"></a>n = int(r0);
    <a id="L87"></a>errno = int(e1);
    <a id="L88"></a>return;
<a id="L89"></a>}

<a id="L91"></a>func sendto(s int, buf []byte, flags int, to uintptr, addrlen _Socklen) (errno int) {
    <a id="L92"></a>var _p0 *byte;
    <a id="L93"></a>if len(buf) &gt; 0 {
        <a id="L94"></a>_p0 = &amp;buf[0]
    <a id="L95"></a>}
    <a id="L96"></a>_, _, e1 := Syscall6(SYS_SENDTO, uintptr(s), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags), uintptr(to), uintptr(addrlen));
    <a id="L97"></a>errno = int(e1);
    <a id="L98"></a>return;
<a id="L99"></a>}

<a id="L101"></a>func kevent(kq int, change uintptr, nchange int, event uintptr, nevent int, timeout *Timespec) (n int, errno int) {
    <a id="L102"></a>r0, _, e1 := Syscall6(SYS_KEVENT, uintptr(kq), uintptr(change), uintptr(nchange), uintptr(event), uintptr(nevent), uintptr(unsafe.Pointer(timeout)));
    <a id="L103"></a>n = int(r0);
    <a id="L104"></a>errno = int(e1);
    <a id="L105"></a>return;
<a id="L106"></a>}

<a id="L108"></a>func sysctl(mib []_C_int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (errno int) {
    <a id="L109"></a>var _p0 *_C_int;
    <a id="L110"></a>if len(mib) &gt; 0 {
        <a id="L111"></a>_p0 = &amp;mib[0]
    <a id="L112"></a>}
    <a id="L113"></a>_, _, e1 := Syscall6(SYS___SYSCTL, uintptr(unsafe.Pointer(_p0)), uintptr(len(mib)), uintptr(unsafe.Pointer(old)), uintptr(unsafe.Pointer(oldlen)), uintptr(unsafe.Pointer(new)), uintptr(newlen));
    <a id="L114"></a>errno = int(e1);
    <a id="L115"></a>return;
<a id="L116"></a>}

<a id="L118"></a>func fcntl(fd int, cmd int, arg int) (val int, errno int) {
    <a id="L119"></a>r0, _, e1 := Syscall(SYS_FCNTL, uintptr(fd), uintptr(cmd), uintptr(arg));
    <a id="L120"></a>val = int(r0);
    <a id="L121"></a>errno = int(e1);
    <a id="L122"></a>return;
<a id="L123"></a>}

<a id="L125"></a>func Access(path string, flags int) (errno int) {
    <a id="L126"></a>_, _, e1 := Syscall(SYS_ACCESS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), 0);
    <a id="L127"></a>errno = int(e1);
    <a id="L128"></a>return;
<a id="L129"></a>}

<a id="L131"></a>func Adjtime(delta *Timeval, olddelta *Timeval) (errno int) {
    <a id="L132"></a>_, _, e1 := Syscall(SYS_ADJTIME, uintptr(unsafe.Pointer(delta)), uintptr(unsafe.Pointer(olddelta)), 0);
    <a id="L133"></a>errno = int(e1);
    <a id="L134"></a>return;
<a id="L135"></a>}

<a id="L137"></a>func Chdir(path string) (errno int) {
    <a id="L138"></a>_, _, e1 := Syscall(SYS_CHDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L139"></a>errno = int(e1);
    <a id="L140"></a>return;
<a id="L141"></a>}

<a id="L143"></a>func Chflags(path string, flags int) (errno int) {
    <a id="L144"></a>_, _, e1 := Syscall(SYS_CHFLAGS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), 0);
    <a id="L145"></a>errno = int(e1);
    <a id="L146"></a>return;
<a id="L147"></a>}

<a id="L149"></a>func Chmod(path string, mode int) (errno int) {
    <a id="L150"></a>_, _, e1 := Syscall(SYS_CHMOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
    <a id="L151"></a>errno = int(e1);
    <a id="L152"></a>return;
<a id="L153"></a>}

<a id="L155"></a>func Chown(path string, uid int, gid int) (errno int) {
    <a id="L156"></a>_, _, e1 := Syscall(SYS_CHOWN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid));
    <a id="L157"></a>errno = int(e1);
    <a id="L158"></a>return;
<a id="L159"></a>}

<a id="L161"></a>func Chroot(path string) (errno int) {
    <a id="L162"></a>_, _, e1 := Syscall(SYS_CHROOT, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L163"></a>errno = int(e1);
    <a id="L164"></a>return;
<a id="L165"></a>}

<a id="L167"></a>func Close(fd int) (errno int) {
    <a id="L168"></a>_, _, e1 := Syscall(SYS_CLOSE, uintptr(fd), 0, 0);
    <a id="L169"></a>errno = int(e1);
    <a id="L170"></a>return;
<a id="L171"></a>}

<a id="L173"></a>func Dup(fd int) (nfd int, errno int) {
    <a id="L174"></a>r0, _, e1 := Syscall(SYS_DUP, uintptr(fd), 0, 0);
    <a id="L175"></a>nfd = int(r0);
    <a id="L176"></a>errno = int(e1);
    <a id="L177"></a>return;
<a id="L178"></a>}

<a id="L180"></a>func Dup2(from int, to int) (errno int) {
    <a id="L181"></a>_, _, e1 := Syscall(SYS_DUP2, uintptr(from), uintptr(to), 0);
    <a id="L182"></a>errno = int(e1);
    <a id="L183"></a>return;
<a id="L184"></a>}

<a id="L186"></a>func Exchangedata(path1 string, path2 string, options int) (errno int) {
    <a id="L187"></a>_, _, e1 := Syscall(SYS_EXCHANGEDATA, uintptr(unsafe.Pointer(StringBytePtr(path1))), uintptr(unsafe.Pointer(StringBytePtr(path2))), uintptr(options));
    <a id="L188"></a>errno = int(e1);
    <a id="L189"></a>return;
<a id="L190"></a>}

<a id="L192"></a>func Exit(code int) {
    <a id="L193"></a>Syscall(SYS_EXIT, uintptr(code), 0, 0);
    <a id="L194"></a>return;
<a id="L195"></a>}

<a id="L197"></a>func Fchdir(fd int) (errno int) {
    <a id="L198"></a>_, _, e1 := Syscall(SYS_FCHDIR, uintptr(fd), 0, 0);
    <a id="L199"></a>errno = int(e1);
    <a id="L200"></a>return;
<a id="L201"></a>}

<a id="L203"></a>func Fchflags(path string, flags int) (errno int) {
    <a id="L204"></a>_, _, e1 := Syscall(SYS_FCHFLAGS, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), 0);
    <a id="L205"></a>errno = int(e1);
    <a id="L206"></a>return;
<a id="L207"></a>}

<a id="L209"></a>func Fchmod(fd int, mode int) (errno int) {
    <a id="L210"></a>_, _, e1 := Syscall(SYS_FCHMOD, uintptr(fd), uintptr(mode), 0);
    <a id="L211"></a>errno = int(e1);
    <a id="L212"></a>return;
<a id="L213"></a>}

<a id="L215"></a>func Fchown(fd int, uid int, gid int) (errno int) {
    <a id="L216"></a>_, _, e1 := Syscall(SYS_FCHOWN, uintptr(fd), uintptr(uid), uintptr(gid));
    <a id="L217"></a>errno = int(e1);
    <a id="L218"></a>return;
<a id="L219"></a>}

<a id="L221"></a>func Flock(fd int, how int) (errno int) {
    <a id="L222"></a>_, _, e1 := Syscall(SYS_FLOCK, uintptr(fd), uintptr(how), 0);
    <a id="L223"></a>errno = int(e1);
    <a id="L224"></a>return;
<a id="L225"></a>}

<a id="L227"></a>func Fpathconf(fd int, name int) (val int, errno int) {
    <a id="L228"></a>r0, _, e1 := Syscall(SYS_FPATHCONF, uintptr(fd), uintptr(name), 0);
    <a id="L229"></a>val = int(r0);
    <a id="L230"></a>errno = int(e1);
    <a id="L231"></a>return;
<a id="L232"></a>}

<a id="L234"></a>func Fstat(fd int, stat *Stat_t) (errno int) {
    <a id="L235"></a>_, _, e1 := Syscall(SYS_FSTAT64, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L236"></a>errno = int(e1);
    <a id="L237"></a>return;
<a id="L238"></a>}

<a id="L240"></a>func Fstatfs(fd int, stat *Statfs_t) (errno int) {
    <a id="L241"></a>_, _, e1 := Syscall(SYS_FSTATFS64, uintptr(fd), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L242"></a>errno = int(e1);
    <a id="L243"></a>return;
<a id="L244"></a>}

<a id="L246"></a>func Fsync(fd int) (errno int) {
    <a id="L247"></a>_, _, e1 := Syscall(SYS_FSYNC, uintptr(fd), 0, 0);
    <a id="L248"></a>errno = int(e1);
    <a id="L249"></a>return;
<a id="L250"></a>}

<a id="L252"></a>func Ftruncate(fd int, length int64) (errno int) {
    <a id="L253"></a>_, _, e1 := Syscall(SYS_FTRUNCATE, uintptr(fd), uintptr(length), 0);
    <a id="L254"></a>errno = int(e1);
    <a id="L255"></a>return;
<a id="L256"></a>}

<a id="L258"></a>func Getdirentries(fd int, buf []byte, basep *uintptr) (n int, errno int) {
    <a id="L259"></a>var _p0 *byte;
    <a id="L260"></a>if len(buf) &gt; 0 {
        <a id="L261"></a>_p0 = &amp;buf[0]
    <a id="L262"></a>}
    <a id="L263"></a>r0, _, e1 := Syscall6(SYS_GETDIRENTRIES64, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(unsafe.Pointer(basep)), 0, 0);
    <a id="L264"></a>n = int(r0);
    <a id="L265"></a>errno = int(e1);
    <a id="L266"></a>return;
<a id="L267"></a>}

<a id="L269"></a>func Getdtablesize() (size int) {
    <a id="L270"></a>r0, _, _ := Syscall(SYS_GETDTABLESIZE, 0, 0, 0);
    <a id="L271"></a>size = int(r0);
    <a id="L272"></a>return;
<a id="L273"></a>}

<a id="L275"></a>func Getegid() (egid int) {
    <a id="L276"></a>r0, _, _ := Syscall(SYS_GETEGID, 0, 0, 0);
    <a id="L277"></a>egid = int(r0);
    <a id="L278"></a>return;
<a id="L279"></a>}

<a id="L281"></a>func Geteuid() (uid int) {
    <a id="L282"></a>r0, _, _ := Syscall(SYS_GETEUID, 0, 0, 0);
    <a id="L283"></a>uid = int(r0);
    <a id="L284"></a>return;
<a id="L285"></a>}

<a id="L287"></a>func Getfsstat(buf []Statfs_t, flags int) (n int, errno int) {
    <a id="L288"></a>var _p0 *Statfs_t;
    <a id="L289"></a>if len(buf) &gt; 0 {
        <a id="L290"></a>_p0 = &amp;buf[0]
    <a id="L291"></a>}
    <a id="L292"></a>r0, _, e1 := Syscall(SYS_GETFSSTAT64, uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)), uintptr(flags));
    <a id="L293"></a>n = int(r0);
    <a id="L294"></a>errno = int(e1);
    <a id="L295"></a>return;
<a id="L296"></a>}

<a id="L298"></a>func Getgid() (gid int) {
    <a id="L299"></a>r0, _, _ := Syscall(SYS_GETGID, 0, 0, 0);
    <a id="L300"></a>gid = int(r0);
    <a id="L301"></a>return;
<a id="L302"></a>}

<a id="L304"></a>func Getpgid(pid int) (pgid int, errno int) {
    <a id="L305"></a>r0, _, e1 := Syscall(SYS_GETPGID, uintptr(pid), 0, 0);
    <a id="L306"></a>pgid = int(r0);
    <a id="L307"></a>errno = int(e1);
    <a id="L308"></a>return;
<a id="L309"></a>}

<a id="L311"></a>func Getpgrp() (pgrp int) {
    <a id="L312"></a>r0, _, _ := Syscall(SYS_GETPGRP, 0, 0, 0);
    <a id="L313"></a>pgrp = int(r0);
    <a id="L314"></a>return;
<a id="L315"></a>}

<a id="L317"></a>func Getpid() (pid int) {
    <a id="L318"></a>r0, _, _ := Syscall(SYS_GETPID, 0, 0, 0);
    <a id="L319"></a>pid = int(r0);
    <a id="L320"></a>return;
<a id="L321"></a>}

<a id="L323"></a>func Getppid() (ppid int) {
    <a id="L324"></a>r0, _, _ := Syscall(SYS_GETPPID, 0, 0, 0);
    <a id="L325"></a>ppid = int(r0);
    <a id="L326"></a>return;
<a id="L327"></a>}

<a id="L329"></a>func Getpriority(which int, who int) (prio int, errno int) {
    <a id="L330"></a>r0, _, e1 := Syscall(SYS_GETPRIORITY, uintptr(which), uintptr(who), 0);
    <a id="L331"></a>prio = int(r0);
    <a id="L332"></a>errno = int(e1);
    <a id="L333"></a>return;
<a id="L334"></a>}

<a id="L336"></a>func Getrlimit(which int, lim *Rlimit) (errno int) {
    <a id="L337"></a>_, _, e1 := Syscall(SYS_GETRLIMIT, uintptr(which), uintptr(unsafe.Pointer(lim)), 0);
    <a id="L338"></a>errno = int(e1);
    <a id="L339"></a>return;
<a id="L340"></a>}

<a id="L342"></a>func Getrusage(who int, rusage *Rusage) (errno int) {
    <a id="L343"></a>_, _, e1 := Syscall(SYS_GETRUSAGE, uintptr(who), uintptr(unsafe.Pointer(rusage)), 0);
    <a id="L344"></a>errno = int(e1);
    <a id="L345"></a>return;
<a id="L346"></a>}

<a id="L348"></a>func Getsid(pid int) (sid int, errno int) {
    <a id="L349"></a>r0, _, e1 := Syscall(SYS_GETSID, uintptr(pid), 0, 0);
    <a id="L350"></a>sid = int(r0);
    <a id="L351"></a>errno = int(e1);
    <a id="L352"></a>return;
<a id="L353"></a>}

<a id="L355"></a>func Getuid() (uid int) {
    <a id="L356"></a>r0, _, _ := Syscall(SYS_GETUID, 0, 0, 0);
    <a id="L357"></a>uid = int(r0);
    <a id="L358"></a>return;
<a id="L359"></a>}

<a id="L361"></a>func Issetugid() (tainted bool) {
    <a id="L362"></a>r0, _, _ := Syscall(SYS_ISSETUGID, 0, 0, 0);
    <a id="L363"></a>tainted = bool(r0 != 0);
    <a id="L364"></a>return;
<a id="L365"></a>}

<a id="L367"></a>func Kill(pid int, signum int, posix int) (errno int) {
    <a id="L368"></a>_, _, e1 := Syscall(SYS_KILL, uintptr(pid), uintptr(signum), uintptr(posix));
    <a id="L369"></a>errno = int(e1);
    <a id="L370"></a>return;
<a id="L371"></a>}

<a id="L373"></a>func Kqueue() (fd int, errno int) {
    <a id="L374"></a>r0, _, e1 := Syscall(SYS_KQUEUE, 0, 0, 0);
    <a id="L375"></a>fd = int(r0);
    <a id="L376"></a>errno = int(e1);
    <a id="L377"></a>return;
<a id="L378"></a>}

<a id="L380"></a>func Lchown(path string, uid int, gid int) (errno int) {
    <a id="L381"></a>_, _, e1 := Syscall(SYS_LCHOWN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(uid), uintptr(gid));
    <a id="L382"></a>errno = int(e1);
    <a id="L383"></a>return;
<a id="L384"></a>}

<a id="L386"></a>func Link(path string, link string) (errno int) {
    <a id="L387"></a>_, _, e1 := Syscall(SYS_LINK, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(StringBytePtr(link))), 0);
    <a id="L388"></a>errno = int(e1);
    <a id="L389"></a>return;
<a id="L390"></a>}

<a id="L392"></a>func Listen(s int, backlog int) (errno int) {
    <a id="L393"></a>_, _, e1 := Syscall(SYS_LISTEN, uintptr(s), uintptr(backlog), 0);
    <a id="L394"></a>errno = int(e1);
    <a id="L395"></a>return;
<a id="L396"></a>}

<a id="L398"></a>func Lstat(path string, stat *Stat_t) (errno int) {
    <a id="L399"></a>_, _, e1 := Syscall(SYS_LSTAT64, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L400"></a>errno = int(e1);
    <a id="L401"></a>return;
<a id="L402"></a>}

<a id="L404"></a>func Mkdir(path string, mode int) (errno int) {
    <a id="L405"></a>_, _, e1 := Syscall(SYS_MKDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
    <a id="L406"></a>errno = int(e1);
    <a id="L407"></a>return;
<a id="L408"></a>}

<a id="L410"></a>func Mkfifo(path string, mode int) (errno int) {
    <a id="L411"></a>_, _, e1 := Syscall(SYS_MKFIFO, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), 0);
    <a id="L412"></a>errno = int(e1);
    <a id="L413"></a>return;
<a id="L414"></a>}

<a id="L416"></a>func Mknod(path string, mode int, dev int) (errno int) {
    <a id="L417"></a>_, _, e1 := Syscall(SYS_MKNOD, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(dev));
    <a id="L418"></a>errno = int(e1);
    <a id="L419"></a>return;
<a id="L420"></a>}

<a id="L422"></a>func Open(path string, mode int, perm int) (fd int, errno int) {
    <a id="L423"></a>r0, _, e1 := Syscall(SYS_OPEN, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(mode), uintptr(perm));
    <a id="L424"></a>fd = int(r0);
    <a id="L425"></a>errno = int(e1);
    <a id="L426"></a>return;
<a id="L427"></a>}

<a id="L429"></a>func Pathconf(path string, name int) (val int, errno int) {
    <a id="L430"></a>r0, _, e1 := Syscall(SYS_PATHCONF, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(name), 0);
    <a id="L431"></a>val = int(r0);
    <a id="L432"></a>errno = int(e1);
    <a id="L433"></a>return;
<a id="L434"></a>}

<a id="L436"></a>func Pread(fd int, p []byte, offset int64) (n int, errno int) {
    <a id="L437"></a>var _p0 *byte;
    <a id="L438"></a>if len(p) &gt; 0 {
        <a id="L439"></a>_p0 = &amp;p[0]
    <a id="L440"></a>}
    <a id="L441"></a>r0, _, e1 := Syscall6(SYS_PREAD, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0);
    <a id="L442"></a>n = int(r0);
    <a id="L443"></a>errno = int(e1);
    <a id="L444"></a>return;
<a id="L445"></a>}

<a id="L447"></a>func Pwrite(fd int, p []byte, offset int64) (n int, errno int) {
    <a id="L448"></a>var _p0 *byte;
    <a id="L449"></a>if len(p) &gt; 0 {
        <a id="L450"></a>_p0 = &amp;p[0]
    <a id="L451"></a>}
    <a id="L452"></a>r0, _, e1 := Syscall6(SYS_PWRITE, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)), uintptr(offset), 0, 0);
    <a id="L453"></a>n = int(r0);
    <a id="L454"></a>errno = int(e1);
    <a id="L455"></a>return;
<a id="L456"></a>}

<a id="L458"></a>func Read(fd int, p []byte) (n int, errno int) {
    <a id="L459"></a>var _p0 *byte;
    <a id="L460"></a>if len(p) &gt; 0 {
        <a id="L461"></a>_p0 = &amp;p[0]
    <a id="L462"></a>}
    <a id="L463"></a>r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)));
    <a id="L464"></a>n = int(r0);
    <a id="L465"></a>errno = int(e1);
    <a id="L466"></a>return;
<a id="L467"></a>}

<a id="L469"></a>func Readlink(path string, buf []byte) (n int, errno int) {
    <a id="L470"></a>var _p0 *byte;
    <a id="L471"></a>if len(buf) &gt; 0 {
        <a id="L472"></a>_p0 = &amp;buf[0]
    <a id="L473"></a>}
    <a id="L474"></a>r0, _, e1 := Syscall(SYS_READLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(_p0)), uintptr(len(buf)));
    <a id="L475"></a>n = int(r0);
    <a id="L476"></a>errno = int(e1);
    <a id="L477"></a>return;
<a id="L478"></a>}

<a id="L480"></a>func Rename(from string, to string) (errno int) {
    <a id="L481"></a>_, _, e1 := Syscall(SYS_RENAME, uintptr(unsafe.Pointer(StringBytePtr(from))), uintptr(unsafe.Pointer(StringBytePtr(to))), 0);
    <a id="L482"></a>errno = int(e1);
    <a id="L483"></a>return;
<a id="L484"></a>}

<a id="L486"></a>func Revoke(path string) (errno int) {
    <a id="L487"></a>_, _, e1 := Syscall(SYS_REVOKE, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L488"></a>errno = int(e1);
    <a id="L489"></a>return;
<a id="L490"></a>}

<a id="L492"></a>func Rmdir(path string) (errno int) {
    <a id="L493"></a>_, _, e1 := Syscall(SYS_RMDIR, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L494"></a>errno = int(e1);
    <a id="L495"></a>return;
<a id="L496"></a>}

<a id="L498"></a>func Seek(fd int, offset int64, whence int) (newoffset int64, errno int) {
    <a id="L499"></a>r0, _, e1 := Syscall(SYS_LSEEK, uintptr(fd), uintptr(offset), uintptr(whence));
    <a id="L500"></a>newoffset = int64(r0);
    <a id="L501"></a>errno = int(e1);
    <a id="L502"></a>return;
<a id="L503"></a>}

<a id="L505"></a>func Select(n int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (errno int) {
    <a id="L506"></a>_, _, e1 := Syscall6(SYS_SELECT, uintptr(n), uintptr(unsafe.Pointer(r)), uintptr(unsafe.Pointer(w)), uintptr(unsafe.Pointer(e)), uintptr(unsafe.Pointer(timeout)), 0);
    <a id="L507"></a>errno = int(e1);
    <a id="L508"></a>return;
<a id="L509"></a>}

<a id="L511"></a>func Setegid(egid int) (errno int) {
    <a id="L512"></a>_, _, e1 := Syscall(SYS_SETEGID, uintptr(egid), 0, 0);
    <a id="L513"></a>errno = int(e1);
    <a id="L514"></a>return;
<a id="L515"></a>}

<a id="L517"></a>func Seteuid(euid int) (errno int) {
    <a id="L518"></a>_, _, e1 := Syscall(SYS_SETEUID, uintptr(euid), 0, 0);
    <a id="L519"></a>errno = int(e1);
    <a id="L520"></a>return;
<a id="L521"></a>}

<a id="L523"></a>func Setgid(gid int) (errno int) {
    <a id="L524"></a>_, _, e1 := Syscall(SYS_SETGID, uintptr(gid), 0, 0);
    <a id="L525"></a>errno = int(e1);
    <a id="L526"></a>return;
<a id="L527"></a>}

<a id="L529"></a>func Setlogin(name string) (errno int) {
    <a id="L530"></a>_, _, e1 := Syscall(SYS_SETLOGIN, uintptr(unsafe.Pointer(StringBytePtr(name))), 0, 0);
    <a id="L531"></a>errno = int(e1);
    <a id="L532"></a>return;
<a id="L533"></a>}

<a id="L535"></a>func Setpgid(pid int, pgid int) (errno int) {
    <a id="L536"></a>_, _, e1 := Syscall(SYS_SETPGID, uintptr(pid), uintptr(pgid), 0);
    <a id="L537"></a>errno = int(e1);
    <a id="L538"></a>return;
<a id="L539"></a>}

<a id="L541"></a>func Setpriority(which int, who int, prio int) (errno int) {
    <a id="L542"></a>_, _, e1 := Syscall(SYS_SETPRIORITY, uintptr(which), uintptr(who), uintptr(prio));
    <a id="L543"></a>errno = int(e1);
    <a id="L544"></a>return;
<a id="L545"></a>}

<a id="L547"></a>func Setprivexec(flag int) (errno int) {
    <a id="L548"></a>_, _, e1 := Syscall(SYS_SETPRIVEXEC, uintptr(flag), 0, 0);
    <a id="L549"></a>errno = int(e1);
    <a id="L550"></a>return;
<a id="L551"></a>}

<a id="L553"></a>func Setregid(rgid int, egid int) (errno int) {
    <a id="L554"></a>_, _, e1 := Syscall(SYS_SETREGID, uintptr(rgid), uintptr(egid), 0);
    <a id="L555"></a>errno = int(e1);
    <a id="L556"></a>return;
<a id="L557"></a>}

<a id="L559"></a>func Setreuid(ruid int, euid int) (errno int) {
    <a id="L560"></a>_, _, e1 := Syscall(SYS_SETREUID, uintptr(ruid), uintptr(euid), 0);
    <a id="L561"></a>errno = int(e1);
    <a id="L562"></a>return;
<a id="L563"></a>}

<a id="L565"></a>func Setrlimit(which int, lim *Rlimit) (errno int) {
    <a id="L566"></a>_, _, e1 := Syscall(SYS_SETRLIMIT, uintptr(which), uintptr(unsafe.Pointer(lim)), 0);
    <a id="L567"></a>errno = int(e1);
    <a id="L568"></a>return;
<a id="L569"></a>}

<a id="L571"></a>func Setsid() (pid int, errno int) {
    <a id="L572"></a>r0, _, e1 := Syscall(SYS_SETSID, 0, 0, 0);
    <a id="L573"></a>pid = int(r0);
    <a id="L574"></a>errno = int(e1);
    <a id="L575"></a>return;
<a id="L576"></a>}

<a id="L578"></a>func Settimeofday(tp *Timeval) (errno int) {
    <a id="L579"></a>_, _, e1 := Syscall(SYS_SETTIMEOFDAY, uintptr(unsafe.Pointer(tp)), 0, 0);
    <a id="L580"></a>errno = int(e1);
    <a id="L581"></a>return;
<a id="L582"></a>}

<a id="L584"></a>func Setuid(uid int) (errno int) {
    <a id="L585"></a>_, _, e1 := Syscall(SYS_SETUID, uintptr(uid), 0, 0);
    <a id="L586"></a>errno = int(e1);
    <a id="L587"></a>return;
<a id="L588"></a>}

<a id="L590"></a>func Stat(path string, stat *Stat_t) (errno int) {
    <a id="L591"></a>_, _, e1 := Syscall(SYS_STAT64, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L592"></a>errno = int(e1);
    <a id="L593"></a>return;
<a id="L594"></a>}

<a id="L596"></a>func Statfs(path string, stat *Statfs_t) (errno int) {
    <a id="L597"></a>_, _, e1 := Syscall(SYS_STATFS64, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(stat)), 0);
    <a id="L598"></a>errno = int(e1);
    <a id="L599"></a>return;
<a id="L600"></a>}

<a id="L602"></a>func Symlink(path string, link string) (errno int) {
    <a id="L603"></a>_, _, e1 := Syscall(SYS_SYMLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(unsafe.Pointer(StringBytePtr(link))), 0);
    <a id="L604"></a>errno = int(e1);
    <a id="L605"></a>return;
<a id="L606"></a>}

<a id="L608"></a>func Sync() (errno int) {
    <a id="L609"></a>_, _, e1 := Syscall(SYS_SYNC, 0, 0, 0);
    <a id="L610"></a>errno = int(e1);
    <a id="L611"></a>return;
<a id="L612"></a>}

<a id="L614"></a>func Truncate(path string, length int64) (errno int) {
    <a id="L615"></a>_, _, e1 := Syscall(SYS_TRUNCATE, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(length), 0);
    <a id="L616"></a>errno = int(e1);
    <a id="L617"></a>return;
<a id="L618"></a>}

<a id="L620"></a>func Umask(newmask int) (errno int) {
    <a id="L621"></a>_, _, e1 := Syscall(SYS_UMASK, uintptr(newmask), 0, 0);
    <a id="L622"></a>errno = int(e1);
    <a id="L623"></a>return;
<a id="L624"></a>}

<a id="L626"></a>func Undelete(path string) (errno int) {
    <a id="L627"></a>_, _, e1 := Syscall(SYS_UNDELETE, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L628"></a>errno = int(e1);
    <a id="L629"></a>return;
<a id="L630"></a>}

<a id="L632"></a>func Unlink(path string) (errno int) {
    <a id="L633"></a>_, _, e1 := Syscall(SYS_UNLINK, uintptr(unsafe.Pointer(StringBytePtr(path))), 0, 0);
    <a id="L634"></a>errno = int(e1);
    <a id="L635"></a>return;
<a id="L636"></a>}

<a id="L638"></a>func Unmount(path string, flags int) (errno int) {
    <a id="L639"></a>_, _, e1 := Syscall(SYS_UNMOUNT, uintptr(unsafe.Pointer(StringBytePtr(path))), uintptr(flags), 0);
    <a id="L640"></a>errno = int(e1);
    <a id="L641"></a>return;
<a id="L642"></a>}

<a id="L644"></a>func Write(fd int, p []byte) (n int, errno int) {
    <a id="L645"></a>var _p0 *byte;
    <a id="L646"></a>if len(p) &gt; 0 {
        <a id="L647"></a>_p0 = &amp;p[0]
    <a id="L648"></a>}
    <a id="L649"></a>r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(len(p)));
    <a id="L650"></a>n = int(r0);
    <a id="L651"></a>errno = int(e1);
    <a id="L652"></a>return;
<a id="L653"></a>}

<a id="L655"></a>func read(fd int, buf *byte, nbuf int) (n int, errno int) {
    <a id="L656"></a>r0, _, e1 := Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf));
    <a id="L657"></a>n = int(r0);
    <a id="L658"></a>errno = int(e1);
    <a id="L659"></a>return;
<a id="L660"></a>}

<a id="L662"></a>func write(fd int, buf *byte, nbuf int) (n int, errno int) {
    <a id="L663"></a>r0, _, e1 := Syscall(SYS_WRITE, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf));
    <a id="L664"></a>n = int(r0);
    <a id="L665"></a>errno = int(e1);
    <a id="L666"></a>return;
<a id="L667"></a>}

<a id="L669"></a>func gettimeofday(tp *Timeval) (sec int64, usec int32, errno int) {
    <a id="L670"></a>r0, r1, e1 := Syscall(SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tp)), 0, 0);
    <a id="L671"></a>sec = int64(r0);
    <a id="L672"></a>usec = int32(r1);
    <a id="L673"></a>errno = int(e1);
    <a id="L674"></a>return;
<a id="L675"></a>}
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
