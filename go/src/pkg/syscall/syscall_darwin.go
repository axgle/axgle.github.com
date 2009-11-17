<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/syscall/syscall_darwin.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/syscall/syscall_darwin.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Darwin system calls.</span>
<a id="L6"></a><span class="comment">// This file is compiled as ordinary Go code,</span>
<a id="L7"></a><span class="comment">// but it is also input to mksyscall,</span>
<a id="L8"></a><span class="comment">// which parses the //sys lines and generates system call stubs.</span>
<a id="L9"></a><span class="comment">// Note that sometimes we use a lowercase //sys name and</span>
<a id="L10"></a><span class="comment">// wrap it in our own nicer implementation.</span>

<a id="L12"></a>package syscall

<a id="L14"></a>import &#34;unsafe&#34;

<a id="L16"></a>const OS = &#34;darwin&#34;

<a id="L18"></a><span class="comment">/*</span>
<a id="L19"></a><span class="comment"> * Pseudo-system calls</span>
<a id="L20"></a><span class="comment"> */</span>
<a id="L21"></a><span class="comment">// The const provides a compile-time constant so clients</span>
<a id="L22"></a><span class="comment">// can adjust to whether there is a working Getwd and avoid</span>
<a id="L23"></a><span class="comment">// even linking this function into the binary.  See ../os/getwd.go.</span>
<a id="L24"></a>const ImplementsGetwd = false

<a id="L26"></a>func Getwd() (string, int) { return &#34;&#34;, ENOTSUP }


<a id="L29"></a><span class="comment">/*</span>
<a id="L30"></a><span class="comment"> * Wrapped</span>
<a id="L31"></a><span class="comment"> */</span>

<a id="L33"></a><span class="comment">//sys	getgroups(ngid int, gid *_Gid_t) (n int, errno int)</span>
<a id="L34"></a><span class="comment">//sys	setgroups(ngid int, gid *_Gid_t) (errno int)</span>

<a id="L36"></a>func Getgroups() (gids []int, errno int) {
    <a id="L37"></a>n, err := getgroups(0, nil);
    <a id="L38"></a>if err != 0 {
        <a id="L39"></a>return nil, errno
    <a id="L40"></a>}
    <a id="L41"></a>if n == 0 {
        <a id="L42"></a>return nil, 0
    <a id="L43"></a>}

    <a id="L45"></a><span class="comment">// Sanity check group count.  Max is 16 on BSD.</span>
    <a id="L46"></a>if n &lt; 0 || n &gt; 1000 {
        <a id="L47"></a>return nil, EINVAL
    <a id="L48"></a>}

    <a id="L50"></a>a := make([]_Gid_t, n);
    <a id="L51"></a>n, err = getgroups(n, &amp;a[0]);
    <a id="L52"></a>if err != 0 {
        <a id="L53"></a>return nil, errno
    <a id="L54"></a>}
    <a id="L55"></a>gids = make([]int, n);
    <a id="L56"></a>for i, v := range a[0:n] {
        <a id="L57"></a>gids[i] = int(v)
    <a id="L58"></a>}
    <a id="L59"></a>return;
<a id="L60"></a>}

<a id="L62"></a>func Setgroups(gids []int) (errno int) {
    <a id="L63"></a>if len(gids) == 0 {
        <a id="L64"></a>return setgroups(0, nil)
    <a id="L65"></a>}

    <a id="L67"></a>a := make([]_Gid_t, len(gids));
    <a id="L68"></a>for i, v := range gids {
        <a id="L69"></a>a[i] = _Gid_t(v)
    <a id="L70"></a>}
    <a id="L71"></a>return setgroups(len(a), &amp;a[0]);
<a id="L72"></a>}

<a id="L74"></a><span class="comment">// Wait status is 7 bits at bottom, either 0 (exited),</span>
<a id="L75"></a><span class="comment">// 0x7F (stopped), or a signal number that caused an exit.</span>
<a id="L76"></a><span class="comment">// The 0x80 bit is whether there was a core dump.</span>
<a id="L77"></a><span class="comment">// An extra number (exit code, signal causing a stop)</span>
<a id="L78"></a><span class="comment">// is in the high bits.</span>

<a id="L80"></a>type WaitStatus uint32

<a id="L82"></a>const (
    <a id="L83"></a>mask  = 0x7F;
    <a id="L84"></a>core  = 0x80;
    <a id="L85"></a>shift = 8;

    <a id="L87"></a>exited  = 0;
    <a id="L88"></a>stopped = 0x7F;
<a id="L89"></a>)

<a id="L91"></a>func (w WaitStatus) Exited() bool { return w&amp;mask == exited }

<a id="L93"></a>func (w WaitStatus) ExitStatus() int {
    <a id="L94"></a>if w&amp;mask != exited {
        <a id="L95"></a>return -1
    <a id="L96"></a>}
    <a id="L97"></a>return int(w &gt;&gt; shift);
<a id="L98"></a>}

<a id="L100"></a>func (w WaitStatus) Signaled() bool { return w&amp;mask != stopped &amp;&amp; w&amp;mask != 0 }

<a id="L102"></a>func (w WaitStatus) Signal() int {
    <a id="L103"></a>sig := int(w &amp; mask);
    <a id="L104"></a>if sig == stopped || sig == 0 {
        <a id="L105"></a>return -1
    <a id="L106"></a>}
    <a id="L107"></a>return sig;
<a id="L108"></a>}

<a id="L110"></a>func (w WaitStatus) CoreDump() bool { return w.Signaled() &amp;&amp; w&amp;core != 0 }

<a id="L112"></a>func (w WaitStatus) Stopped() bool { return w&amp;mask == stopped &amp;&amp; w&gt;&gt;shift != SIGSTOP }

<a id="L114"></a>func (w WaitStatus) Continued() bool { return w&amp;mask == stopped &amp;&amp; w&gt;&gt;shift == SIGSTOP }

<a id="L116"></a>func (w WaitStatus) StopSignal() int {
    <a id="L117"></a>if !w.Stopped() {
        <a id="L118"></a>return -1
    <a id="L119"></a>}
    <a id="L120"></a>return int(w&gt;&gt;shift) &amp; 0xFF;
<a id="L121"></a>}

<a id="L123"></a>func (w WaitStatus) TrapCause() int {
    <a id="L124"></a><span class="comment">// Darwin doesn&#39;t have trap causes</span>
    <a id="L125"></a>return -1
<a id="L126"></a>}

<a id="L128"></a><span class="comment">//sys	wait4(pid int, wstatus *_C_int, options int, rusage *Rusage) (wpid int, errno int)</span>

<a id="L130"></a>func Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (wpid int, errno int) {
    <a id="L131"></a>var status _C_int;
    <a id="L132"></a>wpid, errno = wait4(pid, &amp;status, options, rusage);
    <a id="L133"></a>if wstatus != nil {
        <a id="L134"></a>*wstatus = WaitStatus(status)
    <a id="L135"></a>}
    <a id="L136"></a>return;
<a id="L137"></a>}

<a id="L139"></a><span class="comment">//sys	pipe() (r int, w int, errno int)</span>

<a id="L141"></a>func Pipe(p []int) (errno int) {
    <a id="L142"></a>if len(p) != 2 {
        <a id="L143"></a>return EINVAL
    <a id="L144"></a>}
    <a id="L145"></a>p[0], p[1], errno = pipe();
    <a id="L146"></a>return;
<a id="L147"></a>}

<a id="L149"></a>func Sleep(ns int64) (errno int) {
    <a id="L150"></a>tv := NsecToTimeval(ns);
    <a id="L151"></a>return Select(0, nil, nil, nil, &amp;tv);
<a id="L152"></a>}

<a id="L154"></a><span class="comment">//sys	accept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, errno int)</span>
<a id="L155"></a><span class="comment">//sys	bind(s int, addr uintptr, addrlen _Socklen) (errno int)</span>
<a id="L156"></a><span class="comment">//sys	connect(s int, addr uintptr, addrlen _Socklen) (errno int)</span>
<a id="L157"></a><span class="comment">//sys	socket(domain int, typ int, proto int) (fd int, errno int)</span>
<a id="L158"></a><span class="comment">//sys	setsockopt(s int, level int, name int, val uintptr, vallen int) (errno int)</span>
<a id="L159"></a><span class="comment">//sys	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int)</span>
<a id="L160"></a><span class="comment">//sys	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (errno int)</span>

<a id="L162"></a><span class="comment">// For testing: clients can set this flag to force</span>
<a id="L163"></a><span class="comment">// creation of IPv6 sockets to return EAFNOSUPPORT.</span>
<a id="L164"></a>var SocketDisableIPv6 bool

<a id="L166"></a>type Sockaddr interface {
    <a id="L167"></a>sockaddr() (ptr uintptr, len _Socklen, errno int); <span class="comment">// lowercase; only we can define Sockaddrs</span>
<a id="L168"></a>}

<a id="L170"></a>type SockaddrInet4 struct {
    <a id="L171"></a>Port int;
    <a id="L172"></a>Addr [4]byte;
    <a id="L173"></a>raw  RawSockaddrInet4;
<a id="L174"></a>}

<a id="L176"></a>func (sa *SockaddrInet4) sockaddr() (uintptr, _Socklen, int) {
    <a id="L177"></a>if sa.Port &lt; 0 || sa.Port &gt; 0xFFFF {
        <a id="L178"></a>return 0, 0, EINVAL
    <a id="L179"></a>}
    <a id="L180"></a>sa.raw.Len = SizeofSockaddrInet4;
    <a id="L181"></a>sa.raw.Family = AF_INET;
    <a id="L182"></a>p := (*[2]byte)(unsafe.Pointer(&amp;sa.raw.Port));
    <a id="L183"></a>p[0] = byte(sa.Port &gt;&gt; 8);
    <a id="L184"></a>p[1] = byte(sa.Port);
    <a id="L185"></a>for i := 0; i &lt; len(sa.Addr); i++ {
        <a id="L186"></a>sa.raw.Addr[i] = sa.Addr[i]
    <a id="L187"></a>}
    <a id="L188"></a>return uintptr(unsafe.Pointer(&amp;sa.raw)), _Socklen(sa.raw.Len), 0;
<a id="L189"></a>}

<a id="L191"></a>type SockaddrInet6 struct {
    <a id="L192"></a>Port int;
    <a id="L193"></a>Addr [16]byte;
    <a id="L194"></a>raw  RawSockaddrInet6;
<a id="L195"></a>}

<a id="L197"></a>func (sa *SockaddrInet6) sockaddr() (uintptr, _Socklen, int) {
    <a id="L198"></a>if sa.Port &lt; 0 || sa.Port &gt; 0xFFFF {
        <a id="L199"></a>return 0, 0, EINVAL
    <a id="L200"></a>}
    <a id="L201"></a>sa.raw.Len = SizeofSockaddrInet6;
    <a id="L202"></a>sa.raw.Family = AF_INET6;
    <a id="L203"></a>p := (*[2]byte)(unsafe.Pointer(&amp;sa.raw.Port));
    <a id="L204"></a>p[0] = byte(sa.Port &gt;&gt; 8);
    <a id="L205"></a>p[1] = byte(sa.Port);
    <a id="L206"></a>for i := 0; i &lt; len(sa.Addr); i++ {
        <a id="L207"></a>sa.raw.Addr[i] = sa.Addr[i]
    <a id="L208"></a>}
    <a id="L209"></a>return uintptr(unsafe.Pointer(&amp;sa.raw)), _Socklen(sa.raw.Len), 0;
<a id="L210"></a>}

<a id="L212"></a>type SockaddrUnix struct {
    <a id="L213"></a>Name string;
    <a id="L214"></a>raw  RawSockaddrUnix;
<a id="L215"></a>}

<a id="L217"></a>func (sa *SockaddrUnix) sockaddr() (uintptr, _Socklen, int) {
    <a id="L218"></a>name := sa.Name;
    <a id="L219"></a>n := len(name);
    <a id="L220"></a>if n &gt;= len(sa.raw.Path) || n == 0 {
        <a id="L221"></a>return 0, 0, EINVAL
    <a id="L222"></a>}
    <a id="L223"></a>sa.raw.Len = byte(3 + n); <span class="comment">// 2 for Family, Len; 1 for NUL</span>
    <a id="L224"></a>sa.raw.Family = AF_UNIX;
    <a id="L225"></a>for i := 0; i &lt; n; i++ {
        <a id="L226"></a>sa.raw.Path[i] = int8(name[i])
    <a id="L227"></a>}
    <a id="L228"></a>return uintptr(unsafe.Pointer(&amp;sa.raw)), _Socklen(sa.raw.Len), 0;
<a id="L229"></a>}

<a id="L231"></a>func anyToSockaddr(rsa *RawSockaddrAny) (Sockaddr, int) {
    <a id="L232"></a>switch rsa.Addr.Family {
    <a id="L233"></a>case AF_UNIX:
        <a id="L234"></a>pp := (*RawSockaddrUnix)(unsafe.Pointer(rsa));
        <a id="L235"></a>if pp.Len &lt; 3 || pp.Len &gt; SizeofSockaddrUnix {
            <a id="L236"></a>return nil, EINVAL
        <a id="L237"></a>}
        <a id="L238"></a>sa := new(SockaddrUnix);
        <a id="L239"></a>n := int(pp.Len) - 3; <span class="comment">// subtract leading Family, Len, terminating NUL</span>
        <a id="L240"></a>for i := 0; i &lt; n; i++ {
            <a id="L241"></a>if pp.Path[i] == 0 {
                <a id="L242"></a><span class="comment">// found early NUL; assume Len is overestimating</span>
                <a id="L243"></a>n = i;
                <a id="L244"></a>break;
            <a id="L245"></a>}
        <a id="L246"></a>}
        <a id="L247"></a>bytes := (*[len(pp.Path)]byte)(unsafe.Pointer(&amp;pp.Path[0]));
        <a id="L248"></a>sa.Name = string(bytes[0:n]);
        <a id="L249"></a>return sa, 0;

    <a id="L251"></a>case AF_INET:
        <a id="L252"></a>pp := (*RawSockaddrInet4)(unsafe.Pointer(rsa));
        <a id="L253"></a>sa := new(SockaddrInet4);
        <a id="L254"></a>p := (*[2]byte)(unsafe.Pointer(&amp;pp.Port));
        <a id="L255"></a>sa.Port = int(p[0])&lt;&lt;8 + int(p[1]);
        <a id="L256"></a>for i := 0; i &lt; len(sa.Addr); i++ {
            <a id="L257"></a>sa.Addr[i] = pp.Addr[i]
        <a id="L258"></a>}
        <a id="L259"></a>return sa, 0;

    <a id="L261"></a>case AF_INET6:
        <a id="L262"></a>pp := (*RawSockaddrInet6)(unsafe.Pointer(rsa));
        <a id="L263"></a>sa := new(SockaddrInet6);
        <a id="L264"></a>p := (*[2]byte)(unsafe.Pointer(&amp;pp.Port));
        <a id="L265"></a>sa.Port = int(p[0])&lt;&lt;8 + int(p[1]);
        <a id="L266"></a>for i := 0; i &lt; len(sa.Addr); i++ {
            <a id="L267"></a>sa.Addr[i] = pp.Addr[i]
        <a id="L268"></a>}
        <a id="L269"></a>return sa, 0;
    <a id="L270"></a>}
    <a id="L271"></a>return nil, EAFNOSUPPORT;
<a id="L272"></a>}

<a id="L274"></a>func Accept(fd int) (nfd int, sa Sockaddr, errno int) {
    <a id="L275"></a>var rsa RawSockaddrAny;
    <a id="L276"></a>var len _Socklen = SizeofSockaddrAny;
    <a id="L277"></a>nfd, errno = accept(fd, &amp;rsa, &amp;len);
    <a id="L278"></a>if errno != 0 {
        <a id="L279"></a>return
    <a id="L280"></a>}
    <a id="L281"></a>sa, errno = anyToSockaddr(&amp;rsa);
    <a id="L282"></a>if errno != 0 {
        <a id="L283"></a>Close(nfd);
        <a id="L284"></a>nfd = 0;
    <a id="L285"></a>}
    <a id="L286"></a>return;
<a id="L287"></a>}

<a id="L289"></a>func Getsockname(fd int) (sa Sockaddr, errno int) {
    <a id="L290"></a>var rsa RawSockaddrAny;
    <a id="L291"></a>var len _Socklen = SizeofSockaddrAny;
    <a id="L292"></a>if errno = getsockname(fd, &amp;rsa, &amp;len); errno != 0 {
        <a id="L293"></a>return
    <a id="L294"></a>}
    <a id="L295"></a>return anyToSockaddr(&amp;rsa);
<a id="L296"></a>}

<a id="L298"></a>func Getpeername(fd int) (sa Sockaddr, errno int) {
    <a id="L299"></a>var rsa RawSockaddrAny;
    <a id="L300"></a>var len _Socklen = SizeofSockaddrAny;
    <a id="L301"></a>if errno = getpeername(fd, &amp;rsa, &amp;len); errno != 0 {
        <a id="L302"></a>return
    <a id="L303"></a>}
    <a id="L304"></a>return anyToSockaddr(&amp;rsa);
<a id="L305"></a>}

<a id="L307"></a>func Bind(fd int, sa Sockaddr) (errno int) {
    <a id="L308"></a>ptr, n, err := sa.sockaddr();
    <a id="L309"></a>if err != 0 {
        <a id="L310"></a>return err
    <a id="L311"></a>}
    <a id="L312"></a>return bind(fd, ptr, n);
<a id="L313"></a>}

<a id="L315"></a>func Connect(fd int, sa Sockaddr) (errno int) {
    <a id="L316"></a>ptr, n, err := sa.sockaddr();
    <a id="L317"></a>if err != 0 {
        <a id="L318"></a>return err
    <a id="L319"></a>}
    <a id="L320"></a>return connect(fd, ptr, n);
<a id="L321"></a>}

<a id="L323"></a>func Socket(domain, typ, proto int) (fd, errno int) {
    <a id="L324"></a>if domain == AF_INET6 &amp;&amp; SocketDisableIPv6 {
        <a id="L325"></a>return -1, EAFNOSUPPORT
    <a id="L326"></a>}
    <a id="L327"></a>fd, errno = socket(domain, typ, proto);
    <a id="L328"></a>return;
<a id="L329"></a>}

<a id="L331"></a>func SetsockoptInt(fd, level, opt int, value int) (errno int) {
    <a id="L332"></a>var n = int32(value);
    <a id="L333"></a>return setsockopt(fd, level, opt, uintptr(unsafe.Pointer(&amp;n)), 4);
<a id="L334"></a>}

<a id="L336"></a>func SetsockoptTimeval(fd, level, opt int, tv *Timeval) (errno int) {
    <a id="L337"></a>return setsockopt(fd, level, opt, uintptr(unsafe.Pointer(tv)), unsafe.Sizeof(*tv))
<a id="L338"></a>}

<a id="L340"></a>func SetsockoptLinger(fd, level, opt int, l *Linger) (errno int) {
    <a id="L341"></a>return setsockopt(fd, level, opt, uintptr(unsafe.Pointer(l)), unsafe.Sizeof(*l))
<a id="L342"></a>}


<a id="L345"></a><span class="comment">//sys recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, errno int)</span>

<a id="L347"></a>func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, errno int) {
    <a id="L348"></a>var rsa RawSockaddrAny;
    <a id="L349"></a>var len _Socklen = SizeofSockaddrAny;
    <a id="L350"></a>if n, errno = recvfrom(fd, p, flags, &amp;rsa, &amp;len); errno != 0 {
        <a id="L351"></a>return
    <a id="L352"></a>}
    <a id="L353"></a>from, errno = anyToSockaddr(&amp;rsa);
    <a id="L354"></a>return;
<a id="L355"></a>}

<a id="L357"></a><span class="comment">//sys sendto(s int, buf []byte, flags int, to uintptr, addrlen _Socklen) (errno int)</span>

<a id="L359"></a>func Sendto(fd int, p []byte, flags int, to Sockaddr) (errno int) {
    <a id="L360"></a>ptr, n, err := to.sockaddr();
    <a id="L361"></a>if err != 0 {
        <a id="L362"></a>return err
    <a id="L363"></a>}
    <a id="L364"></a>return sendto(fd, p, flags, ptr, n);
<a id="L365"></a>}

<a id="L367"></a><span class="comment">//sys	kevent(kq int, change uintptr, nchange int, event uintptr, nevent int, timeout *Timespec) (n int, errno int)</span>

<a id="L369"></a>func Kevent(kq int, changes, events []Kevent_t, timeout *Timespec) (n int, errno int) {
    <a id="L370"></a>var change, event uintptr;
    <a id="L371"></a>if len(changes) &gt; 0 {
        <a id="L372"></a>change = uintptr(unsafe.Pointer(&amp;changes[0]))
    <a id="L373"></a>}
    <a id="L374"></a>if len(events) &gt; 0 {
        <a id="L375"></a>event = uintptr(unsafe.Pointer(&amp;events[0]))
    <a id="L376"></a>}
    <a id="L377"></a>return kevent(kq, change, len(changes), event, len(events), timeout);
<a id="L378"></a>}

<a id="L380"></a><span class="comment">//sys	sysctl(mib []_C_int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (errno int) = SYS___SYSCTL</span>

<a id="L382"></a><span class="comment">// Translate &#34;kern.hostname&#34; to []_C_int{0,1,2,3}.</span>
<a id="L383"></a>func nametomib(name string) (mib []_C_int, errno int) {
    <a id="L384"></a>const CTL_MAXNAME = 12;
    <a id="L385"></a>const siz = uintptr(unsafe.Sizeof(mib[0]));

    <a id="L387"></a><span class="comment">// NOTE(rsc): It seems strange to set the buffer to have</span>
    <a id="L388"></a><span class="comment">// size CTL_MAXNAME+2 but use only CTL_MAXNAME</span>
    <a id="L389"></a><span class="comment">// as the size.  I don&#39;t know why the +2 is here, but the</span>
    <a id="L390"></a><span class="comment">// kernel uses +2 for its own implementation of this function.</span>
    <a id="L391"></a><span class="comment">// I am scared that if we don&#39;t include the +2 here, the kernel</span>
    <a id="L392"></a><span class="comment">// will silently write 2 words farther than we specify</span>
    <a id="L393"></a><span class="comment">// and we&#39;ll get memory corruption.</span>
    <a id="L394"></a>var buf [CTL_MAXNAME + 2]_C_int;
    <a id="L395"></a>n := uintptr(CTL_MAXNAME) * siz;

    <a id="L397"></a>p := (*byte)(unsafe.Pointer(&amp;buf[0]));
    <a id="L398"></a>bytes := StringByteSlice(name);

    <a id="L400"></a><span class="comment">// Magic sysctl: &#34;setting&#34; 0.3 to a string name</span>
    <a id="L401"></a><span class="comment">// lets you read back the array of integers form.</span>
    <a id="L402"></a>if errno = sysctl([]_C_int{0, 3}, p, &amp;n, &amp;bytes[0], uintptr(len(name))); errno != 0 {
        <a id="L403"></a>return nil, errno
    <a id="L404"></a>}
    <a id="L405"></a>return buf[0 : n/siz], 0;
<a id="L406"></a>}

<a id="L408"></a>func Sysctl(name string) (value string, errno int) {
    <a id="L409"></a><span class="comment">// Translate name to mib number.</span>
    <a id="L410"></a>mib, errno := nametomib(name);
    <a id="L411"></a>if errno != 0 {
        <a id="L412"></a>return &#34;&#34;, errno
    <a id="L413"></a>}

    <a id="L415"></a><span class="comment">// Find size.</span>
    <a id="L416"></a>n := uintptr(0);
    <a id="L417"></a>if errno = sysctl(mib, nil, &amp;n, nil, 0); errno != 0 {
        <a id="L418"></a>return &#34;&#34;, errno
    <a id="L419"></a>}
    <a id="L420"></a>if n == 0 {
        <a id="L421"></a>return &#34;&#34;, 0
    <a id="L422"></a>}

    <a id="L424"></a><span class="comment">// Read into buffer of that size.</span>
    <a id="L425"></a>buf := make([]byte, n);
    <a id="L426"></a>if errno = sysctl(mib, &amp;buf[0], &amp;n, nil, 0); errno != 0 {
        <a id="L427"></a>return &#34;&#34;, errno
    <a id="L428"></a>}

    <a id="L430"></a><span class="comment">// Throw away terminating NUL.</span>
    <a id="L431"></a>if n &gt; 0 &amp;&amp; buf[n-1] == &#39;\x00&#39; {
        <a id="L432"></a>n--
    <a id="L433"></a>}
    <a id="L434"></a>return string(buf[0:n]), 0;
<a id="L435"></a>}

<a id="L437"></a>func SysctlUint32(name string) (value uint32, errno int) {
    <a id="L438"></a><span class="comment">// Translate name to mib number.</span>
    <a id="L439"></a>mib, errno := nametomib(name);
    <a id="L440"></a>if errno != 0 {
        <a id="L441"></a>return 0, errno
    <a id="L442"></a>}

    <a id="L444"></a><span class="comment">// Read into buffer of that size.</span>
    <a id="L445"></a>n := uintptr(4);
    <a id="L446"></a>buf := make([]byte, 4);
    <a id="L447"></a>if errno = sysctl(mib, &amp;buf[0], &amp;n, nil, 0); errno != 0 {
        <a id="L448"></a>return 0, errno
    <a id="L449"></a>}
    <a id="L450"></a>if n != 4 {
        <a id="L451"></a>return 0, EIO
    <a id="L452"></a>}
    <a id="L453"></a>return *(*uint32)(unsafe.Pointer(&amp;buf[0])), 0;
<a id="L454"></a>}

<a id="L456"></a><span class="comment">// TODO: wrap</span>
<a id="L457"></a><span class="comment">//	Acct(name nil-string) (errno int)</span>
<a id="L458"></a><span class="comment">//	Futimes(fd int, timeval *Timeval) (errno int)	// Pointer to 2 timevals!</span>
<a id="L459"></a><span class="comment">//	Gethostuuid(uuid *byte, timeout *Timespec) (errno int)</span>
<a id="L460"></a><span class="comment">//	Getsockopt(s int, level int, name int, val *byte, vallen *int) (errno int)</span>
<a id="L461"></a><span class="comment">//	Madvise(addr *byte, len int, behav int) (errno int)</span>
<a id="L462"></a><span class="comment">//	Mprotect(addr *byte, len int, prot int) (errno int)</span>
<a id="L463"></a><span class="comment">//	Msync(addr *byte, len int, flags int) (errno int)</span>
<a id="L464"></a><span class="comment">//	Munmap(addr *byte, len int) (errno int)</span>
<a id="L465"></a><span class="comment">//	Ptrace(req int, pid int, addr uintptr, data int) (ret uintptr, errno int)</span>
<a id="L466"></a><span class="comment">//	Recvmsg(s int, msg *Msghdr, flags int) (n int, errno int)</span>
<a id="L467"></a><span class="comment">//	Sendmsg(s int, msg *Msghdr, flags int) (n int, errno int)</span>
<a id="L468"></a><span class="comment">//	Utimes(path string, timeval *Timeval) (errno int)	// Pointer to 2 timevals!</span>
<a id="L469"></a><span class="comment">//sys	fcntl(fd int, cmd int, arg int) (val int, errno int)</span>


<a id="L472"></a><span class="comment">/*</span>
<a id="L473"></a><span class="comment"> * Exposed directly</span>
<a id="L474"></a><span class="comment"> */</span>
<a id="L475"></a><span class="comment">//sys	Access(path string, flags int) (errno int)</span>
<a id="L476"></a><span class="comment">//sys	Adjtime(delta *Timeval, olddelta *Timeval) (errno int)</span>
<a id="L477"></a><span class="comment">//sys	Chdir(path string) (errno int)</span>
<a id="L478"></a><span class="comment">//sys	Chflags(path string, flags int) (errno int)</span>
<a id="L479"></a><span class="comment">//sys	Chmod(path string, mode int) (errno int)</span>
<a id="L480"></a><span class="comment">//sys	Chown(path string, uid int, gid int) (errno int)</span>
<a id="L481"></a><span class="comment">//sys	Chroot(path string) (errno int)</span>
<a id="L482"></a><span class="comment">//sys	Close(fd int) (errno int)</span>
<a id="L483"></a><span class="comment">//sys	Dup(fd int) (nfd int, errno int)</span>
<a id="L484"></a><span class="comment">//sys	Dup2(from int, to int) (errno int)</span>
<a id="L485"></a><span class="comment">//sys	Exchangedata(path1 string, path2 string, options int) (errno int)</span>
<a id="L486"></a><span class="comment">//sys	Exit(code int)</span>
<a id="L487"></a><span class="comment">//sys	Fchdir(fd int) (errno int)</span>
<a id="L488"></a><span class="comment">//sys	Fchflags(path string, flags int) (errno int)</span>
<a id="L489"></a><span class="comment">//sys	Fchmod(fd int, mode int) (errno int)</span>
<a id="L490"></a><span class="comment">//sys	Fchown(fd int, uid int, gid int) (errno int)</span>
<a id="L491"></a><span class="comment">//sys	Flock(fd int, how int) (errno int)</span>
<a id="L492"></a><span class="comment">//sys	Fpathconf(fd int, name int) (val int, errno int)</span>
<a id="L493"></a><span class="comment">//sys	Fstat(fd int, stat *Stat_t) (errno int) = SYS_FSTAT64</span>
<a id="L494"></a><span class="comment">//sys	Fstatfs(fd int, stat *Statfs_t) (errno int) = SYS_FSTATFS64</span>
<a id="L495"></a><span class="comment">//sys	Fsync(fd int) (errno int)</span>
<a id="L496"></a><span class="comment">//sys	Ftruncate(fd int, length int64) (errno int)</span>
<a id="L497"></a><span class="comment">//sys	Getdirentries(fd int, buf []byte, basep *uintptr) (n int, errno int) = SYS_GETDIRENTRIES64</span>
<a id="L498"></a><span class="comment">//sys	Getdtablesize() (size int)</span>
<a id="L499"></a><span class="comment">//sys	Getegid() (egid int)</span>
<a id="L500"></a><span class="comment">//sys	Geteuid() (uid int)</span>
<a id="L501"></a><span class="comment">//sys	Getfsstat(buf []Statfs_t, flags int) (n int, errno int) = SYS_GETFSSTAT64</span>
<a id="L502"></a><span class="comment">//sys	Getgid() (gid int)</span>
<a id="L503"></a><span class="comment">//sys	Getpgid(pid int) (pgid int, errno int)</span>
<a id="L504"></a><span class="comment">//sys	Getpgrp() (pgrp int)</span>
<a id="L505"></a><span class="comment">//sys	Getpid() (pid int)</span>
<a id="L506"></a><span class="comment">//sys	Getppid() (ppid int)</span>
<a id="L507"></a><span class="comment">//sys	Getpriority(which int, who int) (prio int, errno int)</span>
<a id="L508"></a><span class="comment">//sys	Getrlimit(which int, lim *Rlimit) (errno int)</span>
<a id="L509"></a><span class="comment">//sys	Getrusage(who int, rusage *Rusage) (errno int)</span>
<a id="L510"></a><span class="comment">//sys	Getsid(pid int) (sid int, errno int)</span>
<a id="L511"></a><span class="comment">//sys	Getuid() (uid int)</span>
<a id="L512"></a><span class="comment">//sys	Issetugid() (tainted bool)</span>
<a id="L513"></a><span class="comment">//sys	Kill(pid int, signum int, posix int) (errno int)</span>
<a id="L514"></a><span class="comment">//sys	Kqueue() (fd int, errno int)</span>
<a id="L515"></a><span class="comment">//sys	Lchown(path string, uid int, gid int) (errno int)</span>
<a id="L516"></a><span class="comment">//sys	Link(path string, link string) (errno int)</span>
<a id="L517"></a><span class="comment">//sys	Listen(s int, backlog int) (errno int)</span>
<a id="L518"></a><span class="comment">//sys	Lstat(path string, stat *Stat_t) (errno int) = SYS_LSTAT64</span>
<a id="L519"></a><span class="comment">//sys	Mkdir(path string, mode int) (errno int)</span>
<a id="L520"></a><span class="comment">//sys	Mkfifo(path string, mode int) (errno int)</span>
<a id="L521"></a><span class="comment">//sys	Mknod(path string, mode int, dev int) (errno int)</span>
<a id="L522"></a><span class="comment">//sys	Open(path string, mode int, perm int) (fd int, errno int)</span>
<a id="L523"></a><span class="comment">//sys	Pathconf(path string, name int) (val int, errno int)</span>
<a id="L524"></a><span class="comment">//sys	Pread(fd int, p []byte, offset int64) (n int, errno int)</span>
<a id="L525"></a><span class="comment">//sys	Pwrite(fd int, p []byte, offset int64) (n int, errno int)</span>
<a id="L526"></a><span class="comment">//sys	Read(fd int, p []byte) (n int, errno int)</span>
<a id="L527"></a><span class="comment">//sys	Readlink(path string, buf []byte) (n int, errno int)</span>
<a id="L528"></a><span class="comment">//sys	Rename(from string, to string) (errno int)</span>
<a id="L529"></a><span class="comment">//sys	Revoke(path string) (errno int)</span>
<a id="L530"></a><span class="comment">//sys	Rmdir(path string) (errno int)</span>
<a id="L531"></a><span class="comment">//sys	Seek(fd int, offset int64, whence int) (newoffset int64, errno int) = SYS_LSEEK</span>
<a id="L532"></a><span class="comment">//sys	Select(n int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (errno int)</span>
<a id="L533"></a><span class="comment">//sys	Setegid(egid int) (errno int)</span>
<a id="L534"></a><span class="comment">//sys	Seteuid(euid int) (errno int)</span>
<a id="L535"></a><span class="comment">//sys	Setgid(gid int) (errno int)</span>
<a id="L536"></a><span class="comment">//sys	Setlogin(name string) (errno int)</span>
<a id="L537"></a><span class="comment">//sys	Setpgid(pid int, pgid int) (errno int)</span>
<a id="L538"></a><span class="comment">//sys	Setpriority(which int, who int, prio int) (errno int)</span>
<a id="L539"></a><span class="comment">//sys	Setprivexec(flag int) (errno int)</span>
<a id="L540"></a><span class="comment">//sys	Setregid(rgid int, egid int) (errno int)</span>
<a id="L541"></a><span class="comment">//sys	Setreuid(ruid int, euid int) (errno int)</span>
<a id="L542"></a><span class="comment">//sys	Setrlimit(which int, lim *Rlimit) (errno int)</span>
<a id="L543"></a><span class="comment">//sys	Setsid() (pid int, errno int)</span>
<a id="L544"></a><span class="comment">//sys	Settimeofday(tp *Timeval) (errno int)</span>
<a id="L545"></a><span class="comment">//sys	Setuid(uid int) (errno int)</span>
<a id="L546"></a><span class="comment">//sys	Stat(path string, stat *Stat_t) (errno int) = SYS_STAT64</span>
<a id="L547"></a><span class="comment">//sys	Statfs(path string, stat *Statfs_t) (errno int) = SYS_STATFS64</span>
<a id="L548"></a><span class="comment">//sys	Symlink(path string, link string) (errno int)</span>
<a id="L549"></a><span class="comment">//sys	Sync() (errno int)</span>
<a id="L550"></a><span class="comment">//sys	Truncate(path string, length int64) (errno int)</span>
<a id="L551"></a><span class="comment">//sys	Umask(newmask int) (errno int)</span>
<a id="L552"></a><span class="comment">//sys	Undelete(path string) (errno int)</span>
<a id="L553"></a><span class="comment">//sys	Unlink(path string) (errno int)</span>
<a id="L554"></a><span class="comment">//sys	Unmount(path string, flags int) (errno int)</span>
<a id="L555"></a><span class="comment">//sys	Write(fd int, p []byte) (n int, errno int)</span>
<a id="L556"></a><span class="comment">//sys	read(fd int, buf *byte, nbuf int) (n int, errno int)</span>
<a id="L557"></a><span class="comment">//sys	write(fd int, buf *byte, nbuf int) (n int, errno int)</span>


<a id="L560"></a><span class="comment">/*</span>
<a id="L561"></a><span class="comment"> * Unimplemented</span>
<a id="L562"></a><span class="comment"> */</span>
<a id="L563"></a><span class="comment">// Profil</span>
<a id="L564"></a><span class="comment">// Sigaction</span>
<a id="L565"></a><span class="comment">// Sigprocmask</span>
<a id="L566"></a><span class="comment">// Getlogin</span>
<a id="L567"></a><span class="comment">// Sigpending</span>
<a id="L568"></a><span class="comment">// Sigaltstack</span>
<a id="L569"></a><span class="comment">// Ioctl</span>
<a id="L570"></a><span class="comment">// Reboot</span>
<a id="L571"></a><span class="comment">// Execve</span>
<a id="L572"></a><span class="comment">// Vfork</span>
<a id="L573"></a><span class="comment">// Sbrk</span>
<a id="L574"></a><span class="comment">// Sstk</span>
<a id="L575"></a><span class="comment">// Ovadvise</span>
<a id="L576"></a><span class="comment">// Mincore</span>
<a id="L577"></a><span class="comment">// Setitimer</span>
<a id="L578"></a><span class="comment">// Swapon</span>
<a id="L579"></a><span class="comment">// Select</span>
<a id="L580"></a><span class="comment">// Sigsuspend</span>
<a id="L581"></a><span class="comment">// Readv</span>
<a id="L582"></a><span class="comment">// Writev</span>
<a id="L583"></a><span class="comment">// Nfssvc</span>
<a id="L584"></a><span class="comment">// Getfh</span>
<a id="L585"></a><span class="comment">// Quotactl</span>
<a id="L586"></a><span class="comment">// Mount</span>
<a id="L587"></a><span class="comment">// Csops</span>
<a id="L588"></a><span class="comment">// Waitid</span>
<a id="L589"></a><span class="comment">// Add_profil</span>
<a id="L590"></a><span class="comment">// Kdebug_trace</span>
<a id="L591"></a><span class="comment">// Sigreturn</span>
<a id="L592"></a><span class="comment">// Mmap</span>
<a id="L593"></a><span class="comment">// Mlock</span>
<a id="L594"></a><span class="comment">// Munlock</span>
<a id="L595"></a><span class="comment">// Atsocket</span>
<a id="L596"></a><span class="comment">// Kqueue_from_portset_np</span>
<a id="L597"></a><span class="comment">// Kqueue_portset</span>
<a id="L598"></a><span class="comment">// Getattrlist</span>
<a id="L599"></a><span class="comment">// Setattrlist</span>
<a id="L600"></a><span class="comment">// Getdirentriesattr</span>
<a id="L601"></a><span class="comment">// Searchfs</span>
<a id="L602"></a><span class="comment">// Delete</span>
<a id="L603"></a><span class="comment">// Copyfile</span>
<a id="L604"></a><span class="comment">// Poll</span>
<a id="L605"></a><span class="comment">// Watchevent</span>
<a id="L606"></a><span class="comment">// Waitevent</span>
<a id="L607"></a><span class="comment">// Modwatch</span>
<a id="L608"></a><span class="comment">// Getxattr</span>
<a id="L609"></a><span class="comment">// Fgetxattr</span>
<a id="L610"></a><span class="comment">// Setxattr</span>
<a id="L611"></a><span class="comment">// Fsetxattr</span>
<a id="L612"></a><span class="comment">// Removexattr</span>
<a id="L613"></a><span class="comment">// Fremovexattr</span>
<a id="L614"></a><span class="comment">// Listxattr</span>
<a id="L615"></a><span class="comment">// Flistxattr</span>
<a id="L616"></a><span class="comment">// Fsctl</span>
<a id="L617"></a><span class="comment">// Initgroups</span>
<a id="L618"></a><span class="comment">// Posix_spawn</span>
<a id="L619"></a><span class="comment">// Nfsclnt</span>
<a id="L620"></a><span class="comment">// Fhopen</span>
<a id="L621"></a><span class="comment">// Minherit</span>
<a id="L622"></a><span class="comment">// Semsys</span>
<a id="L623"></a><span class="comment">// Msgsys</span>
<a id="L624"></a><span class="comment">// Shmsys</span>
<a id="L625"></a><span class="comment">// Semctl</span>
<a id="L626"></a><span class="comment">// Semget</span>
<a id="L627"></a><span class="comment">// Semop</span>
<a id="L628"></a><span class="comment">// Msgctl</span>
<a id="L629"></a><span class="comment">// Msgget</span>
<a id="L630"></a><span class="comment">// Msgsnd</span>
<a id="L631"></a><span class="comment">// Msgrcv</span>
<a id="L632"></a><span class="comment">// Shmat</span>
<a id="L633"></a><span class="comment">// Shmctl</span>
<a id="L634"></a><span class="comment">// Shmdt</span>
<a id="L635"></a><span class="comment">// Shmget</span>
<a id="L636"></a><span class="comment">// Shm_open</span>
<a id="L637"></a><span class="comment">// Shm_unlink</span>
<a id="L638"></a><span class="comment">// Sem_open</span>
<a id="L639"></a><span class="comment">// Sem_close</span>
<a id="L640"></a><span class="comment">// Sem_unlink</span>
<a id="L641"></a><span class="comment">// Sem_wait</span>
<a id="L642"></a><span class="comment">// Sem_trywait</span>
<a id="L643"></a><span class="comment">// Sem_post</span>
<a id="L644"></a><span class="comment">// Sem_getvalue</span>
<a id="L645"></a><span class="comment">// Sem_init</span>
<a id="L646"></a><span class="comment">// Sem_destroy</span>
<a id="L647"></a><span class="comment">// Open_extended</span>
<a id="L648"></a><span class="comment">// Umask_extended</span>
<a id="L649"></a><span class="comment">// Stat_extended</span>
<a id="L650"></a><span class="comment">// Lstat_extended</span>
<a id="L651"></a><span class="comment">// Fstat_extended</span>
<a id="L652"></a><span class="comment">// Chmod_extended</span>
<a id="L653"></a><span class="comment">// Fchmod_extended</span>
<a id="L654"></a><span class="comment">// Access_extended</span>
<a id="L655"></a><span class="comment">// Settid</span>
<a id="L656"></a><span class="comment">// Gettid</span>
<a id="L657"></a><span class="comment">// Setsgroups</span>
<a id="L658"></a><span class="comment">// Getsgroups</span>
<a id="L659"></a><span class="comment">// Setwgroups</span>
<a id="L660"></a><span class="comment">// Getwgroups</span>
<a id="L661"></a><span class="comment">// Mkfifo_extended</span>
<a id="L662"></a><span class="comment">// Mkdir_extended</span>
<a id="L663"></a><span class="comment">// Identitysvc</span>
<a id="L664"></a><span class="comment">// Shared_region_check_np</span>
<a id="L665"></a><span class="comment">// Shared_region_map_np</span>
<a id="L666"></a><span class="comment">// __pthread_mutex_destroy</span>
<a id="L667"></a><span class="comment">// __pthread_mutex_init</span>
<a id="L668"></a><span class="comment">// __pthread_mutex_lock</span>
<a id="L669"></a><span class="comment">// __pthread_mutex_trylock</span>
<a id="L670"></a><span class="comment">// __pthread_mutex_unlock</span>
<a id="L671"></a><span class="comment">// __pthread_cond_init</span>
<a id="L672"></a><span class="comment">// __pthread_cond_destroy</span>
<a id="L673"></a><span class="comment">// __pthread_cond_broadcast</span>
<a id="L674"></a><span class="comment">// __pthread_cond_signal</span>
<a id="L675"></a><span class="comment">// Setsid_with_pid</span>
<a id="L676"></a><span class="comment">// __pthread_cond_timedwait</span>
<a id="L677"></a><span class="comment">// Aio_fsync</span>
<a id="L678"></a><span class="comment">// Aio_return</span>
<a id="L679"></a><span class="comment">// Aio_suspend</span>
<a id="L680"></a><span class="comment">// Aio_cancel</span>
<a id="L681"></a><span class="comment">// Aio_error</span>
<a id="L682"></a><span class="comment">// Aio_read</span>
<a id="L683"></a><span class="comment">// Aio_write</span>
<a id="L684"></a><span class="comment">// Lio_listio</span>
<a id="L685"></a><span class="comment">// __pthread_cond_wait</span>
<a id="L686"></a><span class="comment">// Iopolicysys</span>
<a id="L687"></a><span class="comment">// Mlockall</span>
<a id="L688"></a><span class="comment">// Munlockall</span>
<a id="L689"></a><span class="comment">// __pthread_kill</span>
<a id="L690"></a><span class="comment">// __pthread_sigmask</span>
<a id="L691"></a><span class="comment">// __sigwait</span>
<a id="L692"></a><span class="comment">// __disable_threadsignal</span>
<a id="L693"></a><span class="comment">// __pthread_markcancel</span>
<a id="L694"></a><span class="comment">// __pthread_canceled</span>
<a id="L695"></a><span class="comment">// __semwait_signal</span>
<a id="L696"></a><span class="comment">// Proc_info</span>
<a id="L697"></a><span class="comment">// Sendfile</span>
<a id="L698"></a><span class="comment">// Stat64_extended</span>
<a id="L699"></a><span class="comment">// Lstat64_extended</span>
<a id="L700"></a><span class="comment">// Fstat64_extended</span>
<a id="L701"></a><span class="comment">// __pthread_chdir</span>
<a id="L702"></a><span class="comment">// __pthread_fchdir</span>
<a id="L703"></a><span class="comment">// Audit</span>
<a id="L704"></a><span class="comment">// Auditon</span>
<a id="L705"></a><span class="comment">// Getauid</span>
<a id="L706"></a><span class="comment">// Setauid</span>
<a id="L707"></a><span class="comment">// Getaudit</span>
<a id="L708"></a><span class="comment">// Setaudit</span>
<a id="L709"></a><span class="comment">// Getaudit_addr</span>
<a id="L710"></a><span class="comment">// Setaudit_addr</span>
<a id="L711"></a><span class="comment">// Auditctl</span>
<a id="L712"></a><span class="comment">// Bsdthread_create</span>
<a id="L713"></a><span class="comment">// Bsdthread_terminate</span>
<a id="L714"></a><span class="comment">// Stack_snapshot</span>
<a id="L715"></a><span class="comment">// Bsdthread_register</span>
<a id="L716"></a><span class="comment">// Workq_open</span>
<a id="L717"></a><span class="comment">// Workq_ops</span>
<a id="L718"></a><span class="comment">// __mac_execve</span>
<a id="L719"></a><span class="comment">// __mac_syscall</span>
<a id="L720"></a><span class="comment">// __mac_get_file</span>
<a id="L721"></a><span class="comment">// __mac_set_file</span>
<a id="L722"></a><span class="comment">// __mac_get_link</span>
<a id="L723"></a><span class="comment">// __mac_set_link</span>
<a id="L724"></a><span class="comment">// __mac_get_proc</span>
<a id="L725"></a><span class="comment">// __mac_set_proc</span>
<a id="L726"></a><span class="comment">// __mac_get_fd</span>
<a id="L727"></a><span class="comment">// __mac_set_fd</span>
<a id="L728"></a><span class="comment">// __mac_get_pid</span>
<a id="L729"></a><span class="comment">// __mac_get_lcid</span>
<a id="L730"></a><span class="comment">// __mac_get_lctx</span>
<a id="L731"></a><span class="comment">// __mac_set_lctx</span>
<a id="L732"></a><span class="comment">// Setlcid</span>
<a id="L733"></a><span class="comment">// Read_nocancel</span>
<a id="L734"></a><span class="comment">// Write_nocancel</span>
<a id="L735"></a><span class="comment">// Open_nocancel</span>
<a id="L736"></a><span class="comment">// Close_nocancel</span>
<a id="L737"></a><span class="comment">// Wait4_nocancel</span>
<a id="L738"></a><span class="comment">// Recvmsg_nocancel</span>
<a id="L739"></a><span class="comment">// Sendmsg_nocancel</span>
<a id="L740"></a><span class="comment">// Recvfrom_nocancel</span>
<a id="L741"></a><span class="comment">// Accept_nocancel</span>
<a id="L742"></a><span class="comment">// Msync_nocancel</span>
<a id="L743"></a><span class="comment">// Fcntl_nocancel</span>
<a id="L744"></a><span class="comment">// Select_nocancel</span>
<a id="L745"></a><span class="comment">// Fsync_nocancel</span>
<a id="L746"></a><span class="comment">// Connect_nocancel</span>
<a id="L747"></a><span class="comment">// Sigsuspend_nocancel</span>
<a id="L748"></a><span class="comment">// Readv_nocancel</span>
<a id="L749"></a><span class="comment">// Writev_nocancel</span>
<a id="L750"></a><span class="comment">// Sendto_nocancel</span>
<a id="L751"></a><span class="comment">// Pread_nocancel</span>
<a id="L752"></a><span class="comment">// Pwrite_nocancel</span>
<a id="L753"></a><span class="comment">// Waitid_nocancel</span>
<a id="L754"></a><span class="comment">// Poll_nocancel</span>
<a id="L755"></a><span class="comment">// Msgsnd_nocancel</span>
<a id="L756"></a><span class="comment">// Msgrcv_nocancel</span>
<a id="L757"></a><span class="comment">// Sem_wait_nocancel</span>
<a id="L758"></a><span class="comment">// Aio_suspend_nocancel</span>
<a id="L759"></a><span class="comment">// __sigwait_nocancel</span>
<a id="L760"></a><span class="comment">// __semwait_signal_nocancel</span>
<a id="L761"></a><span class="comment">// __mac_mount</span>
<a id="L762"></a><span class="comment">// __mac_get_mount</span>
<a id="L763"></a><span class="comment">// __mac_getfsstat</span>
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
