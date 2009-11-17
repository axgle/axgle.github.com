<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/syscall_linux.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/syscall_linux.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Linux system calls.</span>
<a id="L6"></a><span class="comment">// This file is compiled as ordinary Go code,</span>
<a id="L7"></a><span class="comment">// but it is also input to mksyscall,</span>
<a id="L8"></a><span class="comment">// which parses the //sys lines and generates system call stubs.</span>
<a id="L9"></a><span class="comment">// Note that sometimes we use a lowercase //sys name and</span>
<a id="L10"></a><span class="comment">// wrap it in our own nicer implementation.</span>

<a id="L12"></a>package syscall

<a id="L14"></a>import &#34;unsafe&#34;

<a id="L16"></a>const OS = &#34;linux&#34;

<a id="L18"></a><span class="comment">/*</span>
<a id="L19"></a><span class="comment"> * Wrapped</span>
<a id="L20"></a><span class="comment"> */</span>

<a id="L22"></a><span class="comment">//sys	pipe(p *[2]_C_int) (errno int)</span>
<a id="L23"></a>func Pipe(p []int) (errno int) {
    <a id="L24"></a>if len(p) != 2 {
        <a id="L25"></a>return EINVAL
    <a id="L26"></a>}
    <a id="L27"></a>var pp [2]_C_int;
    <a id="L28"></a>errno = pipe(&amp;pp);
    <a id="L29"></a>p[0] = int(pp[0]);
    <a id="L30"></a>p[1] = int(pp[1]);
    <a id="L31"></a>return;
<a id="L32"></a>}

<a id="L34"></a><span class="comment">//sys	utimes(path string, times *[2]Timeval) (errno int)</span>
<a id="L35"></a>func Utimes(path string, tv []Timeval) (errno int) {
    <a id="L36"></a>if len(tv) != 2 {
        <a id="L37"></a>return EINVAL
    <a id="L38"></a>}
    <a id="L39"></a>return utimes(path, (*[2]Timeval)(unsafe.Pointer(&amp;tv[0])));
<a id="L40"></a>}

<a id="L42"></a><span class="comment">//sys	futimesat(dirfd int, path string, times *[2]Timeval) (errno int)</span>
<a id="L43"></a>func Futimesat(dirfd int, path string, tv []Timeval) (errno int) {
    <a id="L44"></a>if len(tv) != 2 {
        <a id="L45"></a>return EINVAL
    <a id="L46"></a>}
    <a id="L47"></a>return futimesat(dirfd, path, (*[2]Timeval)(unsafe.Pointer(&amp;tv[0])));
<a id="L48"></a>}

<a id="L50"></a>const ImplementsGetwd = true

<a id="L52"></a><span class="comment">//sys	Getcwd(buf []byte) (n int, errno int)</span>
<a id="L53"></a>func Getwd() (wd string, errno int) {
    <a id="L54"></a>var buf [PathMax]byte;
    <a id="L55"></a>n, err := Getcwd(&amp;buf);
    <a id="L56"></a>if err != 0 {
        <a id="L57"></a>return &#34;&#34;, err
    <a id="L58"></a>}
    <a id="L59"></a><span class="comment">// Getcwd returns the number of bytes written to buf, including the NUL.</span>
    <a id="L60"></a>if n &lt; 1 || n &gt; len(buf) || buf[n-1] != 0 {
        <a id="L61"></a>return &#34;&#34;, EINVAL
    <a id="L62"></a>}
    <a id="L63"></a>return string(buf[0 : n-1]), 0;
<a id="L64"></a>}

<a id="L66"></a>func Getgroups() (gids []int, errno int) {
    <a id="L67"></a>n, err := getgroups(0, nil);
    <a id="L68"></a>if err != 0 {
        <a id="L69"></a>return nil, errno
    <a id="L70"></a>}
    <a id="L71"></a>if n == 0 {
        <a id="L72"></a>return nil, 0
    <a id="L73"></a>}

    <a id="L75"></a><span class="comment">// Sanity check group count.  Max is 1&lt;&lt;16 on Linux.</span>
    <a id="L76"></a>if n &lt; 0 || n &gt; 1&lt;&lt;20 {
        <a id="L77"></a>return nil, EINVAL
    <a id="L78"></a>}

    <a id="L80"></a>a := make([]_Gid_t, n);
    <a id="L81"></a>n, err = getgroups(n, &amp;a[0]);
    <a id="L82"></a>if err != 0 {
        <a id="L83"></a>return nil, errno
    <a id="L84"></a>}
    <a id="L85"></a>gids = make([]int, n);
    <a id="L86"></a>for i, v := range a[0:n] {
        <a id="L87"></a>gids[i] = int(v)
    <a id="L88"></a>}
    <a id="L89"></a>return;
<a id="L90"></a>}

<a id="L92"></a>func Setgroups(gids []int) (errno int) {
    <a id="L93"></a>if len(gids) == 0 {
        <a id="L94"></a>return setgroups(0, nil)
    <a id="L95"></a>}

    <a id="L97"></a>a := make([]_Gid_t, len(gids));
    <a id="L98"></a>for i, v := range gids {
        <a id="L99"></a>a[i] = _Gid_t(v)
    <a id="L100"></a>}
    <a id="L101"></a>return setgroups(len(a), &amp;a[0]);
<a id="L102"></a>}

<a id="L104"></a>type WaitStatus uint32

<a id="L106"></a><span class="comment">// Wait status is 7 bits at bottom, either 0 (exited),</span>
<a id="L107"></a><span class="comment">// 0x7F (stopped), or a signal number that caused an exit.</span>
<a id="L108"></a><span class="comment">// The 0x80 bit is whether there was a core dump.</span>
<a id="L109"></a><span class="comment">// An extra number (exit code, signal causing a stop)</span>
<a id="L110"></a><span class="comment">// is in the high bits.  At least that&#39;s the idea.</span>
<a id="L111"></a><span class="comment">// There are various irregularities.  For example, the</span>
<a id="L112"></a><span class="comment">// &#34;continued&#34; status is 0xFFFF, distinguishing itself</span>
<a id="L113"></a><span class="comment">// from stopped via the core dump bit.</span>

<a id="L115"></a>const (
    <a id="L116"></a>mask    = 0x7F;
    <a id="L117"></a>core    = 0x80;
    <a id="L118"></a>exited  = 0x00;
    <a id="L119"></a>stopped = 0x7F;
    <a id="L120"></a>shift   = 8;
<a id="L121"></a>)

<a id="L123"></a>func (w WaitStatus) Exited() bool { return w&amp;mask == exited }

<a id="L125"></a>func (w WaitStatus) Signaled() bool { return w&amp;mask != stopped &amp;&amp; w&amp;mask != exited }

<a id="L127"></a>func (w WaitStatus) Stopped() bool { return w&amp;0xFF == stopped }

<a id="L129"></a>func (w WaitStatus) Continued() bool { return w == 0xFFFF }

<a id="L131"></a>func (w WaitStatus) CoreDump() bool { return w.Signaled() &amp;&amp; w&amp;core != 0 }

<a id="L133"></a>func (w WaitStatus) ExitStatus() int {
    <a id="L134"></a>if !w.Exited() {
        <a id="L135"></a>return -1
    <a id="L136"></a>}
    <a id="L137"></a>return int(w&gt;&gt;shift) &amp; 0xFF;
<a id="L138"></a>}

<a id="L140"></a>func (w WaitStatus) Signal() int {
    <a id="L141"></a>if !w.Signaled() {
        <a id="L142"></a>return -1
    <a id="L143"></a>}
    <a id="L144"></a>return int(w &amp; mask);
<a id="L145"></a>}

<a id="L147"></a>func (w WaitStatus) StopSignal() int {
    <a id="L148"></a>if !w.Stopped() {
        <a id="L149"></a>return -1
    <a id="L150"></a>}
    <a id="L151"></a>return int(w&gt;&gt;shift) &amp; 0xFF;
<a id="L152"></a>}

<a id="L154"></a>func (w WaitStatus) TrapCause() int {
    <a id="L155"></a>if w.StopSignal() != SIGTRAP {
        <a id="L156"></a>return -1
    <a id="L157"></a>}
    <a id="L158"></a>return int(w&gt;&gt;shift) &gt;&gt; 8;
<a id="L159"></a>}

<a id="L161"></a><span class="comment">//sys	wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, errno int)</span>
<a id="L162"></a>func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, errno int) {
    <a id="L163"></a>var status _C_int;
    <a id="L164"></a>wpid, errno = wait4(pid, &amp;status, options, rusage);
    <a id="L165"></a>if wstatus != nil {
        <a id="L166"></a>*wstatus = WaitStatus(status)
    <a id="L167"></a>}
    <a id="L168"></a>return;
<a id="L169"></a>}

<a id="L171"></a>func Sleep(nsec int64) (errno int) {
    <a id="L172"></a>tv := NsecToTimeval(nsec);
    <a id="L173"></a>_, err := Select(0, nil, nil, nil, &amp;tv);
    <a id="L174"></a>return err;
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// For testing: clients can set this flag to force</span>
<a id="L178"></a><span class="comment">// creation of IPv6 sockets to return EAFNOSUPPORT.</span>
<a id="L179"></a>var SocketDisableIPv6 bool

<a id="L181"></a>type Sockaddr interface {
    <a id="L182"></a>sockaddr() (ptr uintptr, len _Socklen, errno int); <span class="comment">// lowercase; only we can define Sockaddrs</span>
<a id="L183"></a>}

<a id="L185"></a>type SockaddrInet4 struct {
    <a id="L186"></a>Port int;
    <a id="L187"></a>Addr [4]byte;
    <a id="L188"></a>raw  RawSockaddrInet4;
<a id="L189"></a>}

<a id="L191"></a>func (sa *SockaddrInet4) sockaddr() (uintptr, _Socklen, int) {
    <a id="L192"></a>if sa.Port &lt; 0 || sa.Port &gt; 0xFFFF {
        <a id="L193"></a>return 0, 0, EINVAL
    <a id="L194"></a>}
    <a id="L195"></a>sa.raw.Family = AF_INET;
    <a id="L196"></a>p := (*[2]byte)(unsafe.Pointer(&amp;sa.raw.Port));
    <a id="L197"></a>p[0] = byte(sa.Port &gt;&gt; 8);
    <a id="L198"></a>p[1] = byte(sa.Port);
    <a id="L199"></a>for i := 0; i &lt; len(sa.Addr); i++ {
        <a id="L200"></a>sa.raw.Addr[i] = sa.Addr[i]
    <a id="L201"></a>}
    <a id="L202"></a>return uintptr(unsafe.Pointer(&amp;sa.raw)), SizeofSockaddrInet4, 0;
<a id="L203"></a>}

<a id="L205"></a>type SockaddrInet6 struct {
    <a id="L206"></a>Port int;
    <a id="L207"></a>Addr [16]byte;
    <a id="L208"></a>raw  RawSockaddrInet6;
<a id="L209"></a>}

<a id="L211"></a>func (sa *SockaddrInet6) sockaddr() (uintptr, _Socklen, int) {
    <a id="L212"></a>if sa.Port &lt; 0 || sa.Port &gt; 0xFFFF {
        <a id="L213"></a>return 0, 0, EINVAL
    <a id="L214"></a>}
    <a id="L215"></a>sa.raw.Family = AF_INET6;
    <a id="L216"></a>p := (*[2]byte)(unsafe.Pointer(&amp;sa.raw.Port));
    <a id="L217"></a>p[0] = byte(sa.Port &gt;&gt; 8);
    <a id="L218"></a>p[1] = byte(sa.Port);
    <a id="L219"></a>for i := 0; i &lt; len(sa.Addr); i++ {
        <a id="L220"></a>sa.raw.Addr[i] = sa.Addr[i]
    <a id="L221"></a>}
    <a id="L222"></a>return uintptr(unsafe.Pointer(&amp;sa.raw)), SizeofSockaddrInet6, 0;
<a id="L223"></a>}

<a id="L225"></a>type SockaddrUnix struct {
    <a id="L226"></a>Name string;
    <a id="L227"></a>raw  RawSockaddrUnix;
<a id="L228"></a>}

<a id="L230"></a>func (sa *SockaddrUnix) sockaddr() (uintptr, _Socklen, int) {
    <a id="L231"></a>name := sa.Name;
    <a id="L232"></a>n := len(name);
    <a id="L233"></a>if n &gt;= len(sa.raw.Path) || n == 0 {
        <a id="L234"></a>return 0, 0, EINVAL
    <a id="L235"></a>}
    <a id="L236"></a>sa.raw.Family = AF_UNIX;
    <a id="L237"></a>for i := 0; i &lt; n; i++ {
        <a id="L238"></a>sa.raw.Path[i] = int8(name[i])
    <a id="L239"></a>}
    <a id="L240"></a>if sa.raw.Path[0] == &#39;@&#39; {
        <a id="L241"></a>sa.raw.Path[0] = 0
    <a id="L242"></a>}

    <a id="L244"></a><span class="comment">// length is family, name, NUL.</span>
    <a id="L245"></a>return uintptr(unsafe.Pointer(&amp;sa.raw)), 1 + _Socklen(n) + 1, 0;
<a id="L246"></a>}

<a id="L248"></a>func anyToSockaddr(rsa *RawSockaddrAny) (Sockaddr, int) {
    <a id="L249"></a>switch rsa.Addr.Family {
    <a id="L250"></a>case AF_UNIX:
        <a id="L251"></a>pp := (*RawSockaddrUnix)(unsafe.Pointer(rsa));
        <a id="L252"></a>sa := new(SockaddrUnix);
        <a id="L253"></a>if pp.Path[0] == 0 {
            <a id="L254"></a><span class="comment">// &#34;Abstract&#34; Unix domain socket.</span>
            <a id="L255"></a><span class="comment">// Rewrite leading NUL as @ for textual display.</span>
            <a id="L256"></a><span class="comment">// (This is the standard convention.)</span>
            <a id="L257"></a><span class="comment">// Not friendly to overwrite in place,</span>
            <a id="L258"></a><span class="comment">// but the callers below don&#39;t care.</span>
            <a id="L259"></a>pp.Path[0] = &#39;@&#39;
        <a id="L260"></a>}

        <a id="L262"></a><span class="comment">// Assume path ends at NUL.</span>
        <a id="L263"></a><span class="comment">// This is not technically the Linux semantics for</span>
        <a id="L264"></a><span class="comment">// abstract Unix domain sockets--they are supposed</span>
        <a id="L265"></a><span class="comment">// to be uninterpreted fixed-size binary blobs--but</span>
        <a id="L266"></a><span class="comment">// everyone uses this convention.</span>
        <a id="L267"></a>n := 0;
        <a id="L268"></a>for n &lt; len(pp.Path) &amp;&amp; pp.Path[n] != 0 {
            <a id="L269"></a>n++
        <a id="L270"></a>}
        <a id="L271"></a>bytes := (*[len(pp.Path)]byte)(unsafe.Pointer(&amp;pp.Path[0]));
        <a id="L272"></a>sa.Name = string(bytes[0:n]);
        <a id="L273"></a>return sa, 0;

    <a id="L275"></a>case AF_INET:
        <a id="L276"></a>pp := (*RawSockaddrInet4)(unsafe.Pointer(rsa));
        <a id="L277"></a>sa := new(SockaddrInet4);
        <a id="L278"></a>p := (*[2]byte)(unsafe.Pointer(&amp;pp.Port));
        <a id="L279"></a>sa.Port = int(p[0])&lt;&lt;8 + int(p[1]);
        <a id="L280"></a>for i := 0; i &lt; len(sa.Addr); i++ {
            <a id="L281"></a>sa.Addr[i] = pp.Addr[i]
        <a id="L282"></a>}
        <a id="L283"></a>return sa, 0;

    <a id="L285"></a>case AF_INET6:
        <a id="L286"></a>pp := (*RawSockaddrInet6)(unsafe.Pointer(rsa));
        <a id="L287"></a>sa := new(SockaddrInet6);
        <a id="L288"></a>p := (*[2]byte)(unsafe.Pointer(&amp;pp.Port));
        <a id="L289"></a>sa.Port = int(p[0])&lt;&lt;8 + int(p[1]);
        <a id="L290"></a>for i := 0; i &lt; len(sa.Addr); i++ {
            <a id="L291"></a>sa.Addr[i] = pp.Addr[i]
        <a id="L292"></a>}
        <a id="L293"></a>return sa, 0;
    <a id="L294"></a>}
    <a id="L295"></a>return nil, EAFNOSUPPORT;
<a id="L296"></a>}

<a id="L298"></a>func Accept(fd int) (nfd int, sa Sockaddr, errno int) {
    <a id="L299"></a>var rsa RawSockaddrAny;
    <a id="L300"></a>var len _Socklen = SizeofSockaddrAny;
    <a id="L301"></a>nfd, errno = accept(fd, &amp;rsa, &amp;len);
    <a id="L302"></a>if errno != 0 {
        <a id="L303"></a>return
    <a id="L304"></a>}
    <a id="L305"></a>sa, errno = anyToSockaddr(&amp;rsa);
    <a id="L306"></a>if errno != 0 {
        <a id="L307"></a>Close(nfd);
        <a id="L308"></a>nfd = 0;
    <a id="L309"></a>}
    <a id="L310"></a>return;
<a id="L311"></a>}

<a id="L313"></a>func Getsockname(fd int) (sa Sockaddr, errno int) {
    <a id="L314"></a>var rsa RawSockaddrAny;
    <a id="L315"></a>var len _Socklen = SizeofSockaddrAny;
    <a id="L316"></a>if errno = getsockname(fd, &amp;rsa, &amp;len); errno != 0 {
        <a id="L317"></a>return
    <a id="L318"></a>}
    <a id="L319"></a>return anyToSockaddr(&amp;rsa);
<a id="L320"></a>}

<a id="L322"></a>func Getpeername(fd int) (sa Sockaddr, errno int) {
    <a id="L323"></a>var rsa RawSockaddrAny;
    <a id="L324"></a>var len _Socklen = SizeofSockaddrAny;
    <a id="L325"></a>if errno = getpeername(fd, &amp;rsa, &amp;len); errno != 0 {
        <a id="L326"></a>return
    <a id="L327"></a>}
    <a id="L328"></a>return anyToSockaddr(&amp;rsa);
<a id="L329"></a>}

<a id="L331"></a>func Bind(fd int, sa Sockaddr) (errno int) {
    <a id="L332"></a>ptr, n, err := sa.sockaddr();
    <a id="L333"></a>if err != 0 {
        <a id="L334"></a>return err
    <a id="L335"></a>}
    <a id="L336"></a>return bind(fd, ptr, n);
<a id="L337"></a>}

<a id="L339"></a>func Connect(fd int, sa Sockaddr) (errno int) {
    <a id="L340"></a>ptr, n, err := sa.sockaddr();
    <a id="L341"></a>if err != 0 {
        <a id="L342"></a>return err
    <a id="L343"></a>}
    <a id="L344"></a>return connect(fd, ptr, n);
<a id="L345"></a>}

<a id="L347"></a>func Socket(domain, typ, proto int) (fd, errno int) {
    <a id="L348"></a>if domain == AF_INET6 &amp;&amp; SocketDisableIPv6 {
        <a id="L349"></a>return -1, EAFNOSUPPORT
    <a id="L350"></a>}
    <a id="L351"></a>fd, errno = socket(domain, typ, proto);
    <a id="L352"></a>return;
<a id="L353"></a>}

<a id="L355"></a>func SetsockoptInt(fd, level, opt int, value int) (errno int) {
    <a id="L356"></a>var n = int32(value);
    <a id="L357"></a>return setsockopt(fd, level, opt, uintptr(unsafe.Pointer(&amp;n)), 4);
<a id="L358"></a>}

<a id="L360"></a>func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (errno int) {
    <a id="L361"></a>return setsockopt(fd, level, opt, uintptr(unsafe.Pointer(tv)), unsafe.Sizeof(*tv))
<a id="L362"></a>}

<a id="L364"></a>func SetsockoptLinger(fd, level, opt int, l *Linger) (errno int) {
    <a id="L365"></a>return setsockopt(fd, level, opt, uintptr(unsafe.Pointer(l)), unsafe.Sizeof(*l))
<a id="L366"></a>}

<a id="L368"></a>func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, errno int) {
    <a id="L369"></a>var rsa RawSockaddrAny;
    <a id="L370"></a>var len _Socklen = SizeofSockaddrAny;
    <a id="L371"></a>if n, errno = recvfrom(fd, p, flags, &amp;rsa, &amp;len); errno != 0 {
        <a id="L372"></a>return
    <a id="L373"></a>}
    <a id="L374"></a>from, errno = anyToSockaddr(&amp;rsa);
    <a id="L375"></a>return;
<a id="L376"></a>}

<a id="L378"></a>func Sendto(fd int, p []byte, flags int, to Sockaddr) (errno int) {
    <a id="L379"></a>ptr, n, err := to.sockaddr();
    <a id="L380"></a>if err != 0 {
        <a id="L381"></a>return err
    <a id="L382"></a>}
    <a id="L383"></a>return sendto(fd, p, flags, ptr, n);
<a id="L384"></a>}

<a id="L386"></a><span class="comment">//sys	ptrace(request int, pid int, addr uintptr, data uintptr) (errno int)</span>

<a id="L388"></a><span class="comment">// See bytes.Copy.</span>
<a id="L389"></a>func bytesCopy(dst, src []byte) int {
    <a id="L390"></a>if len(src) &gt; len(dst) {
        <a id="L391"></a>src = src[0:len(dst)]
    <a id="L392"></a>}
    <a id="L393"></a>for i, x := range src {
        <a id="L394"></a>dst[i] = x
    <a id="L395"></a>}
    <a id="L396"></a>return len(src);
<a id="L397"></a>}

<a id="L399"></a>func ptracePeek(req int, pid int, addr uintptr, out []byte) (count int, errno int) {
    <a id="L400"></a><span class="comment">// The peek requests are machine-size oriented, so we wrap it</span>
    <a id="L401"></a><span class="comment">// to retrieve arbitrary-length data.</span>

    <a id="L403"></a><span class="comment">// The ptrace syscall differs from glibc&#39;s ptrace.</span>
    <a id="L404"></a><span class="comment">// Peeks returns the word in *data, not as the return value.</span>

    <a id="L406"></a>var buf [sizeofPtr]byte;

    <a id="L408"></a><span class="comment">// Leading edge.  PEEKTEXT/PEEKDATA don&#39;t require aligned</span>
    <a id="L409"></a><span class="comment">// access (PEEKUSER warns that it might), but if we don&#39;t</span>
    <a id="L410"></a><span class="comment">// align our reads, we might straddle an unmapped page</span>
    <a id="L411"></a><span class="comment">// boundary and not get the bytes leading up to the page</span>
    <a id="L412"></a><span class="comment">// boundary.</span>
    <a id="L413"></a>n := 0;
    <a id="L414"></a>if addr%sizeofPtr != 0 {
        <a id="L415"></a>errno = ptrace(req, pid, addr-addr%sizeofPtr, uintptr(unsafe.Pointer(&amp;buf[0])));
        <a id="L416"></a>if errno != 0 {
            <a id="L417"></a>return 0, errno
        <a id="L418"></a>}
        <a id="L419"></a>n += bytesCopy(out, buf[addr%sizeofPtr:len(buf)]);
        <a id="L420"></a>out = out[n:len(out)];
    <a id="L421"></a>}

    <a id="L423"></a><span class="comment">// Remainder.</span>
    <a id="L424"></a>for len(out) &gt; 0 {
        <a id="L425"></a><span class="comment">// We use an internal buffer to gaurantee alignment.</span>
        <a id="L426"></a><span class="comment">// It&#39;s not documented if this is necessary, but we&#39;re paranoid.</span>
        <a id="L427"></a>errno = ptrace(req, pid, addr+uintptr(n), uintptr(unsafe.Pointer(&amp;buf[0])));
        <a id="L428"></a>if errno != 0 {
            <a id="L429"></a>return n, errno
        <a id="L430"></a>}
        <a id="L431"></a>copied := bytesCopy(out, &amp;buf);
        <a id="L432"></a>n += copied;
        <a id="L433"></a>out = out[copied:len(out)];
    <a id="L434"></a>}

    <a id="L436"></a>return n, 0;
<a id="L437"></a>}

<a id="L439"></a>func PtracePeekText(pid int, addr uintptr, out []byte) (count int, errno int) {
    <a id="L440"></a>return ptracePeek(PTRACE_PEEKTEXT, pid, addr, out)
<a id="L441"></a>}

<a id="L443"></a>func PtracePeekData(pid int, addr uintptr, out []byte) (count int, errno int) {
    <a id="L444"></a>return ptracePeek(PTRACE_PEEKDATA, pid, addr, out)
<a id="L445"></a>}

<a id="L447"></a>func ptracePoke(pokeReq int, peekReq int, pid int, addr uintptr, data []byte) (count int, errno int) {
    <a id="L448"></a><span class="comment">// As for ptracePeek, we need to align our accesses to deal</span>
    <a id="L449"></a><span class="comment">// with the possibility of straddling an invalid page.</span>

    <a id="L451"></a><span class="comment">// Leading edge.</span>
    <a id="L452"></a>n := 0;
    <a id="L453"></a>if addr%sizeofPtr != 0 {
        <a id="L454"></a>var buf [sizeofPtr]byte;
        <a id="L455"></a>errno = ptrace(peekReq, pid, addr-addr%sizeofPtr, uintptr(unsafe.Pointer(&amp;buf[0])));
        <a id="L456"></a>if errno != 0 {
            <a id="L457"></a>return 0, errno
        <a id="L458"></a>}
        <a id="L459"></a>n += bytesCopy(buf[addr%sizeofPtr:len(buf)], data);
        <a id="L460"></a>word := *((*uintptr)(unsafe.Pointer(&amp;buf[0])));
        <a id="L461"></a>errno = ptrace(pokeReq, pid, addr-addr%sizeofPtr, word);
        <a id="L462"></a>if errno != 0 {
            <a id="L463"></a>return 0, errno
        <a id="L464"></a>}
        <a id="L465"></a>data = data[n:len(data)];
    <a id="L466"></a>}

    <a id="L468"></a><span class="comment">// Interior.</span>
    <a id="L469"></a>for len(data) &gt; sizeofPtr {
        <a id="L470"></a>word := *((*uintptr)(unsafe.Pointer(&amp;data[0])));
        <a id="L471"></a>errno = ptrace(pokeReq, pid, addr+uintptr(n), word);
        <a id="L472"></a>if errno != 0 {
            <a id="L473"></a>return n, errno
        <a id="L474"></a>}
        <a id="L475"></a>n += sizeofPtr;
        <a id="L476"></a>data = data[sizeofPtr:len(data)];
    <a id="L477"></a>}

    <a id="L479"></a><span class="comment">// Trailing edge.</span>
    <a id="L480"></a>if len(data) &gt; 0 {
        <a id="L481"></a>var buf [sizeofPtr]byte;
        <a id="L482"></a>errno = ptrace(peekReq, pid, addr+uintptr(n), uintptr(unsafe.Pointer(&amp;buf[0])));
        <a id="L483"></a>if errno != 0 {
            <a id="L484"></a>return n, errno
        <a id="L485"></a>}
        <a id="L486"></a>bytesCopy(&amp;buf, data);
        <a id="L487"></a>word := *((*uintptr)(unsafe.Pointer(&amp;buf[0])));
        <a id="L488"></a>errno = ptrace(pokeReq, pid, addr+uintptr(n), word);
        <a id="L489"></a>if errno != 0 {
            <a id="L490"></a>return n, errno
        <a id="L491"></a>}
        <a id="L492"></a>n += len(data);
    <a id="L493"></a>}

    <a id="L495"></a>return n, 0;
<a id="L496"></a>}

<a id="L498"></a>func PtracePokeText(pid int, addr uintptr, data []byte) (count int, errno int) {
    <a id="L499"></a>return ptracePoke(PTRACE_POKETEXT, PTRACE_PEEKTEXT, pid, addr, data)
<a id="L500"></a>}

<a id="L502"></a>func PtracePokeData(pid int, addr uintptr, data []byte) (count int, errno int) {
    <a id="L503"></a>return ptracePoke(PTRACE_POKEDATA, PTRACE_PEEKDATA, pid, addr, data)
<a id="L504"></a>}

<a id="L506"></a>func PtraceGetRegs(pid int, regsout *PtraceRegs) (errno int) {
    <a id="L507"></a>return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
<a id="L508"></a>}

<a id="L510"></a>func PtraceSetRegs(pid int, regs *PtraceRegs) (errno int) {
    <a id="L511"></a>return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
<a id="L512"></a>}

<a id="L514"></a>func PtraceSetOptions(pid int, options int) (errno int) {
    <a id="L515"></a>return ptrace(PTRACE_SETOPTIONS, pid, 0, uintptr(options))
<a id="L516"></a>}

<a id="L518"></a>func PtraceGetEventMsg(pid int) (msg uint, errno int) {
    <a id="L519"></a>var data _C_long;
    <a id="L520"></a>errno = ptrace(PTRACE_GETEVENTMSG, pid, 0, uintptr(unsafe.Pointer(&amp;data)));
    <a id="L521"></a>msg = uint(data);
    <a id="L522"></a>return;
<a id="L523"></a>}

<a id="L525"></a>func PtraceCont(pid int, signal int) (errno int) {
    <a id="L526"></a>return ptrace(PTRACE_CONT, pid, 0, uintptr(signal))
<a id="L527"></a>}

<a id="L529"></a>func PtraceSingleStep(pid int) (errno int) { return ptrace(PTRACE_SINGLESTEP, pid, 0, 0) }

<a id="L531"></a>func PtraceAttach(pid int) (errno int) { return ptrace(PTRACE_ATTACH, pid, 0, 0) }

<a id="L533"></a>func PtraceDetach(pid int) (errno int) { return ptrace(PTRACE_DETACH, pid, 0, 0) }

<a id="L535"></a><span class="comment">// Sendto</span>
<a id="L536"></a><span class="comment">// Recvfrom</span>
<a id="L537"></a><span class="comment">// Sendmsg</span>
<a id="L538"></a><span class="comment">// Recvmsg</span>
<a id="L539"></a><span class="comment">// Socketpair</span>
<a id="L540"></a><span class="comment">// Getsockopt</span>

<a id="L542"></a><span class="comment">/*</span>
<a id="L543"></a><span class="comment"> * Direct access</span>
<a id="L544"></a><span class="comment"> */</span>
<a id="L545"></a><span class="comment">//sys	Access(path string, mode int) (errno int)</span>
<a id="L546"></a><span class="comment">//sys	Acct(path string) (errno int)</span>
<a id="L547"></a><span class="comment">//sys	Adjtimex(buf *Timex) (state int, errno int)</span>
<a id="L548"></a><span class="comment">//sys	Chdir(path string) (errno int)</span>
<a id="L549"></a><span class="comment">//sys	Chmod(path string, mode int) (errno int)</span>
<a id="L550"></a><span class="comment">//sys	Chroot(path string) (errno int)</span>
<a id="L551"></a><span class="comment">//sys	Close(fd int) (errno int)</span>
<a id="L552"></a><span class="comment">//sys	Creat(path string, mode int) (fd int, errno int)</span>
<a id="L553"></a><span class="comment">//sys	Dup(oldfd int) (fd int, errno int)</span>
<a id="L554"></a><span class="comment">//sys	Dup2(oldfd int, newfd int) (fd int, errno int)</span>
<a id="L555"></a><span class="comment">//sys	EpollCreate(size int) (fd int, errno int)</span>
<a id="L556"></a><span class="comment">//sys	EpollCtl(epfd int, op int, fd int, event *EpollEvent) (errno int)</span>
<a id="L557"></a><span class="comment">//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, errno int)</span>
<a id="L558"></a><span class="comment">//sys	Exit(code int) = SYS_EXIT_GROUP</span>
<a id="L559"></a><span class="comment">//sys	Faccessat(dirfd int, path string, mode int, flags int) (errno int)</span>
<a id="L560"></a><span class="comment">//sys	Fallocate(fd int, mode int, off int64, len int64) (errno int)</span>
<a id="L561"></a><span class="comment">//sys	Fchdir(fd int) (errno int)</span>
<a id="L562"></a><span class="comment">//sys	Fchmod(fd int, mode int) (errno int)</span>
<a id="L563"></a><span class="comment">//sys	Fchmodat(dirfd int, path string, mode int, flags int) (errno int)</span>
<a id="L564"></a><span class="comment">//sys	Fchownat(dirfd int, path string, uid int, gid int, flags int) (errno int)</span>
<a id="L565"></a><span class="comment">//sys	fcntl(fd int, cmd int, arg int) (val int, errno int)</span>
<a id="L566"></a><span class="comment">//sys	Fdatasync(fd int) (errno int)</span>
<a id="L567"></a><span class="comment">//sys	Fsync(fd int) (errno int)</span>
<a id="L568"></a><span class="comment">//sys	Ftruncate(fd int, length int64) (errno int)</span>
<a id="L569"></a><span class="comment">//sys	Getdents(fd int, buf []byte) (n int, errno int) = SYS_GETDENTS64</span>
<a id="L570"></a><span class="comment">//sys	Getpgid(pid int) (pgid int, errno int)</span>
<a id="L571"></a><span class="comment">//sys	Getpgrp() (pid int)</span>
<a id="L572"></a><span class="comment">//sys	Getpid() (pid int)</span>
<a id="L573"></a><span class="comment">//sys	Getppid() (ppid int)</span>
<a id="L574"></a><span class="comment">//sys	Getrlimit(resource int, rlim *Rlimit) (errno int)</span>
<a id="L575"></a><span class="comment">//sys	Getrusage(who int, rusage *Rusage) (errno int)</span>
<a id="L576"></a><span class="comment">//sys	Gettid() (tid int)</span>
<a id="L577"></a><span class="comment">//sys	Gettimeofday(tv *Timeval) (errno int)</span>
<a id="L578"></a><span class="comment">//sys	Kill(pid int, sig int) (errno int)</span>
<a id="L579"></a><span class="comment">//sys	Klogctl(typ int, buf []byte) (n int, errno int) = SYS_SYSLOG</span>
<a id="L580"></a><span class="comment">//sys	Link(oldpath string, newpath string) (errno int)</span>
<a id="L581"></a><span class="comment">//sys	Mkdir(path string, mode int) (errno int)</span>
<a id="L582"></a><span class="comment">//sys	Mkdirat(dirfd int, path string, mode int) (errno int)</span>
<a id="L583"></a><span class="comment">//sys	Mknod(path string, mode int, dev int) (errno int)</span>
<a id="L584"></a><span class="comment">//sys	Mknodat(dirfd int, path string, mode int, dev int) (errno int)</span>
<a id="L585"></a><span class="comment">//sys	Nanosleep(time *Timespec, leftover *Timespec) (errno int)</span>
<a id="L586"></a><span class="comment">//sys	Open(path string, mode int, perm int) (fd int, errno int)</span>
<a id="L587"></a><span class="comment">//sys	Openat(dirfd int, path string, flags int, mode int) (fd int, errno int)</span>
<a id="L588"></a><span class="comment">//sys	Pause() (errno int)</span>
<a id="L589"></a><span class="comment">//sys	PivotRoot(newroot string, putold string) (errno int) = SYS_PIVOT_ROOT</span>
<a id="L590"></a><span class="comment">//sys	Pread(fd int, p []byte, offset int64) (n int, errno int) = SYS_PREAD64</span>
<a id="L591"></a><span class="comment">//sys	Pwrite(fd int, p []byte, offset int64) (n int, errno int) = SYS_PWRITE64</span>
<a id="L592"></a><span class="comment">//sys	Read(fd int, p []byte) (n int, errno int)</span>
<a id="L593"></a><span class="comment">//sys	Readlink(path string, buf []byte) (n int, errno int)</span>
<a id="L594"></a><span class="comment">//sys	Rename(oldpath string, newpath string) (errno int)</span>
<a id="L595"></a><span class="comment">//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (errno int)</span>
<a id="L596"></a><span class="comment">//sys	Rmdir(path string) (errno int)</span>
<a id="L597"></a><span class="comment">//sys	Setdomainname(p []byte) (errno int)</span>
<a id="L598"></a><span class="comment">//sys	Sethostname(p []byte) (errno int)</span>
<a id="L599"></a><span class="comment">//sys	Setpgid(pid int, pgid int) (errno int)</span>
<a id="L600"></a><span class="comment">//sys	Setrlimit(resource int, rlim *Rlimit) (errno int)</span>
<a id="L601"></a><span class="comment">//sys	Setsid() (pid int)</span>
<a id="L602"></a><span class="comment">//sys	Settimeofday(tv *Timeval) (errno int)</span>
<a id="L603"></a><span class="comment">//sys	Setuid(uid int) (errno int)</span>
<a id="L604"></a><span class="comment">//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, errno int)</span>
<a id="L605"></a><span class="comment">//sys	Symlink(oldpath string, newpath string) (errno int)</span>
<a id="L606"></a><span class="comment">//sys	Sync()</span>
<a id="L607"></a><span class="comment">//sys	Sysinfo(info *Sysinfo_t) (errno int)</span>
<a id="L608"></a><span class="comment">//sys	Tee(rfd int, wfd int, len int, flags int) (n int64, errno int)</span>
<a id="L609"></a><span class="comment">//sys	Tgkill(tgid int, tid int, sig int) (errno int)</span>
<a id="L610"></a><span class="comment">//sys	Time(t *Time_t) (tt Time_t, errno int)</span>
<a id="L611"></a><span class="comment">//sys	Times(tms *Tms) (ticks uintptr, errno int)</span>
<a id="L612"></a><span class="comment">//sys	Truncate(path string, length int64) (errno int)</span>
<a id="L613"></a><span class="comment">//sys	Umask(mask int) (oldmask int)</span>
<a id="L614"></a><span class="comment">//sys	Uname(buf *Utsname) (errno int)</span>
<a id="L615"></a><span class="comment">//sys	Unlink(path string) (errno int)</span>
<a id="L616"></a><span class="comment">//sys	Unlinkat(dirfd int, path string) (errno int)</span>
<a id="L617"></a><span class="comment">//sys	Unshare(flags int) (errno int)</span>
<a id="L618"></a><span class="comment">//sys	Ustat(dev int, ubuf *Ustat_t) (errno int)</span>
<a id="L619"></a><span class="comment">//sys	Utime(path string, buf *Utimbuf) (errno int)</span>
<a id="L620"></a><span class="comment">//sys	Write(fd int, p []byte) (n int, errno int)</span>
<a id="L621"></a><span class="comment">//sys	exitThread(code int) (errno int) = SYS_EXIT</span>
<a id="L622"></a><span class="comment">//sys	read(fd int, p *byte, np int) (n int, errno int)</span>
<a id="L623"></a><span class="comment">//sys	write(fd int, p *byte, np int) (n int, errno int)</span>

<a id="L625"></a><span class="comment">/*</span>
<a id="L626"></a><span class="comment"> * Unimplemented</span>
<a id="L627"></a><span class="comment"> */</span>
<a id="L628"></a><span class="comment">// AddKey</span>
<a id="L629"></a><span class="comment">// AfsSyscall</span>
<a id="L630"></a><span class="comment">// Alarm</span>
<a id="L631"></a><span class="comment">// ArchPrctl</span>
<a id="L632"></a><span class="comment">// Brk</span>
<a id="L633"></a><span class="comment">// Capget</span>
<a id="L634"></a><span class="comment">// Capset</span>
<a id="L635"></a><span class="comment">// ClockGetres</span>
<a id="L636"></a><span class="comment">// ClockGettime</span>
<a id="L637"></a><span class="comment">// ClockNanosleep</span>
<a id="L638"></a><span class="comment">// ClockSettime</span>
<a id="L639"></a><span class="comment">// Clone</span>
<a id="L640"></a><span class="comment">// CreateModule</span>
<a id="L641"></a><span class="comment">// DeleteModule</span>
<a id="L642"></a><span class="comment">// EpollCtlOld</span>
<a id="L643"></a><span class="comment">// EpollPwait</span>
<a id="L644"></a><span class="comment">// EpollWaitOld</span>
<a id="L645"></a><span class="comment">// Eventfd</span>
<a id="L646"></a><span class="comment">// Execve</span>
<a id="L647"></a><span class="comment">// Fadvise64</span>
<a id="L648"></a><span class="comment">// Fgetxattr</span>
<a id="L649"></a><span class="comment">// Flistxattr</span>
<a id="L650"></a><span class="comment">// Flock</span>
<a id="L651"></a><span class="comment">// Fork</span>
<a id="L652"></a><span class="comment">// Fremovexattr</span>
<a id="L653"></a><span class="comment">// Fsetxattr</span>
<a id="L654"></a><span class="comment">// Futex</span>
<a id="L655"></a><span class="comment">// GetKernelSyms</span>
<a id="L656"></a><span class="comment">// GetMempolicy</span>
<a id="L657"></a><span class="comment">// GetRobustList</span>
<a id="L658"></a><span class="comment">// GetThreadArea</span>
<a id="L659"></a><span class="comment">// Getitimer</span>
<a id="L660"></a><span class="comment">// Getpmsg</span>
<a id="L661"></a><span class="comment">// Getpriority</span>
<a id="L662"></a><span class="comment">// Getxattr</span>
<a id="L663"></a><span class="comment">// InotifyAddWatch</span>
<a id="L664"></a><span class="comment">// InotifyInit</span>
<a id="L665"></a><span class="comment">// InotifyRmWatch</span>
<a id="L666"></a><span class="comment">// IoCancel</span>
<a id="L667"></a><span class="comment">// IoDestroy</span>
<a id="L668"></a><span class="comment">// IoGetevents</span>
<a id="L669"></a><span class="comment">// IoSetup</span>
<a id="L670"></a><span class="comment">// IoSubmit</span>
<a id="L671"></a><span class="comment">// Ioctl</span>
<a id="L672"></a><span class="comment">// IoprioGet</span>
<a id="L673"></a><span class="comment">// IoprioSet</span>
<a id="L674"></a><span class="comment">// KexecLoad</span>
<a id="L675"></a><span class="comment">// Keyctl</span>
<a id="L676"></a><span class="comment">// Lgetxattr</span>
<a id="L677"></a><span class="comment">// Listxattr</span>
<a id="L678"></a><span class="comment">// Llistxattr</span>
<a id="L679"></a><span class="comment">// LookupDcookie</span>
<a id="L680"></a><span class="comment">// Lremovexattr</span>
<a id="L681"></a><span class="comment">// Lsetxattr</span>
<a id="L682"></a><span class="comment">// Madvise</span>
<a id="L683"></a><span class="comment">// Mbind</span>
<a id="L684"></a><span class="comment">// MigratePages</span>
<a id="L685"></a><span class="comment">// Mincore</span>
<a id="L686"></a><span class="comment">// Mlock</span>
<a id="L687"></a><span class="comment">// Mmap</span>
<a id="L688"></a><span class="comment">// ModifyLdt</span>
<a id="L689"></a><span class="comment">// Mount</span>
<a id="L690"></a><span class="comment">// MovePages</span>
<a id="L691"></a><span class="comment">// Mprotect</span>
<a id="L692"></a><span class="comment">// MqGetsetattr</span>
<a id="L693"></a><span class="comment">// MqNotify</span>
<a id="L694"></a><span class="comment">// MqOpen</span>
<a id="L695"></a><span class="comment">// MqTimedreceive</span>
<a id="L696"></a><span class="comment">// MqTimedsend</span>
<a id="L697"></a><span class="comment">// MqUnlink</span>
<a id="L698"></a><span class="comment">// Mremap</span>
<a id="L699"></a><span class="comment">// Msgctl</span>
<a id="L700"></a><span class="comment">// Msgget</span>
<a id="L701"></a><span class="comment">// Msgrcv</span>
<a id="L702"></a><span class="comment">// Msgsnd</span>
<a id="L703"></a><span class="comment">// Msync</span>
<a id="L704"></a><span class="comment">// Munlock</span>
<a id="L705"></a><span class="comment">// Munlockall</span>
<a id="L706"></a><span class="comment">// Munmap</span>
<a id="L707"></a><span class="comment">// Newfstatat</span>
<a id="L708"></a><span class="comment">// Nfsservctl</span>
<a id="L709"></a><span class="comment">// Personality</span>
<a id="L710"></a><span class="comment">// Poll</span>
<a id="L711"></a><span class="comment">// Ppoll</span>
<a id="L712"></a><span class="comment">// Prctl</span>
<a id="L713"></a><span class="comment">// Pselect6</span>
<a id="L714"></a><span class="comment">// Ptrace</span>
<a id="L715"></a><span class="comment">// Putpmsg</span>
<a id="L716"></a><span class="comment">// QueryModule</span>
<a id="L717"></a><span class="comment">// Quotactl</span>
<a id="L718"></a><span class="comment">// Readahead</span>
<a id="L719"></a><span class="comment">// Readv</span>
<a id="L720"></a><span class="comment">// Reboot</span>
<a id="L721"></a><span class="comment">// RemapFilePages</span>
<a id="L722"></a><span class="comment">// Removexattr</span>
<a id="L723"></a><span class="comment">// RequestKey</span>
<a id="L724"></a><span class="comment">// RestartSyscall</span>
<a id="L725"></a><span class="comment">// RtSigaction</span>
<a id="L726"></a><span class="comment">// RtSigpending</span>
<a id="L727"></a><span class="comment">// RtSigprocmask</span>
<a id="L728"></a><span class="comment">// RtSigqueueinfo</span>
<a id="L729"></a><span class="comment">// RtSigreturn</span>
<a id="L730"></a><span class="comment">// RtSigsuspend</span>
<a id="L731"></a><span class="comment">// RtSigtimedwait</span>
<a id="L732"></a><span class="comment">// SchedGetPriorityMax</span>
<a id="L733"></a><span class="comment">// SchedGetPriorityMin</span>
<a id="L734"></a><span class="comment">// SchedGetaffinity</span>
<a id="L735"></a><span class="comment">// SchedGetparam</span>
<a id="L736"></a><span class="comment">// SchedGetscheduler</span>
<a id="L737"></a><span class="comment">// SchedRrGetInterval</span>
<a id="L738"></a><span class="comment">// SchedSetaffinity</span>
<a id="L739"></a><span class="comment">// SchedSetparam</span>
<a id="L740"></a><span class="comment">// SchedYield</span>
<a id="L741"></a><span class="comment">// Security</span>
<a id="L742"></a><span class="comment">// Semctl</span>
<a id="L743"></a><span class="comment">// Semget</span>
<a id="L744"></a><span class="comment">// Semop</span>
<a id="L745"></a><span class="comment">// Semtimedop</span>
<a id="L746"></a><span class="comment">// Sendfile</span>
<a id="L747"></a><span class="comment">// SetMempolicy</span>
<a id="L748"></a><span class="comment">// SetRobustList</span>
<a id="L749"></a><span class="comment">// SetThreadArea</span>
<a id="L750"></a><span class="comment">// SetTidAddress</span>
<a id="L751"></a><span class="comment">// Setpriority</span>
<a id="L752"></a><span class="comment">// Setxattr</span>
<a id="L753"></a><span class="comment">// Shmat</span>
<a id="L754"></a><span class="comment">// Shmctl</span>
<a id="L755"></a><span class="comment">// Shmdt</span>
<a id="L756"></a><span class="comment">// Shmget</span>
<a id="L757"></a><span class="comment">// Sigaltstack</span>
<a id="L758"></a><span class="comment">// Signalfd</span>
<a id="L759"></a><span class="comment">// Swapoff</span>
<a id="L760"></a><span class="comment">// Swapon</span>
<a id="L761"></a><span class="comment">// Sysfs</span>
<a id="L762"></a><span class="comment">// TimerCreate</span>
<a id="L763"></a><span class="comment">// TimerDelete</span>
<a id="L764"></a><span class="comment">// TimerGetoverrun</span>
<a id="L765"></a><span class="comment">// TimerGettime</span>
<a id="L766"></a><span class="comment">// TimerSettime</span>
<a id="L767"></a><span class="comment">// Timerfd</span>
<a id="L768"></a><span class="comment">// Tkill (obsolete)</span>
<a id="L769"></a><span class="comment">// Tuxcall</span>
<a id="L770"></a><span class="comment">// Umount2</span>
<a id="L771"></a><span class="comment">// Uselib</span>
<a id="L772"></a><span class="comment">// Utimensat</span>
<a id="L773"></a><span class="comment">// Vfork</span>
<a id="L774"></a><span class="comment">// Vhangup</span>
<a id="L775"></a><span class="comment">// Vmsplice</span>
<a id="L776"></a><span class="comment">// Vserver</span>
<a id="L777"></a><span class="comment">// Waitid</span>
<a id="L778"></a><span class="comment">// Writev</span>
<a id="L779"></a><span class="comment">// _Sysctl</span>
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
