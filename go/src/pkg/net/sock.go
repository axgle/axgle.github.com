<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/sock.go</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/net/sock.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Sockets</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;reflect&#34;;
    <a id="L12"></a>&#34;syscall&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// Boolean to int.</span>
<a id="L16"></a>func boolint(b bool) int {
    <a id="L17"></a>if b {
        <a id="L18"></a>return 1
    <a id="L19"></a>}
    <a id="L20"></a>return 0;
<a id="L21"></a>}

<a id="L23"></a><span class="comment">// Generic socket creation.</span>
<a id="L24"></a>func socket(net string, f, p, t int, la, ra syscall.Sockaddr, toAddr func(syscall.Sockaddr) Addr) (fd *netFD, err os.Error) {
    <a id="L25"></a><span class="comment">// See ../syscall/exec.go for description of ForkLock.</span>
    <a id="L26"></a>syscall.ForkLock.RLock();
    <a id="L27"></a>s, e := syscall.Socket(f, p, t);
    <a id="L28"></a>if e != 0 {
        <a id="L29"></a>syscall.ForkLock.RUnlock();
        <a id="L30"></a>return nil, os.Errno(e);
    <a id="L31"></a>}
    <a id="L32"></a>syscall.CloseOnExec(s);
    <a id="L33"></a>syscall.ForkLock.RUnlock();

    <a id="L35"></a><span class="comment">// Allow reuse of recently-used addresses.</span>
    <a id="L36"></a>syscall.SetsockoptInt(s, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1);

    <a id="L38"></a>if la != nil {
        <a id="L39"></a>e = syscall.Bind(s, la);
        <a id="L40"></a>if e != 0 {
            <a id="L41"></a>syscall.Close(s);
            <a id="L42"></a>return nil, os.Errno(e);
        <a id="L43"></a>}
    <a id="L44"></a>}

    <a id="L46"></a>if ra != nil {
        <a id="L47"></a>e = syscall.Connect(s, ra);
        <a id="L48"></a>if e != 0 {
            <a id="L49"></a>syscall.Close(s);
            <a id="L50"></a>return nil, os.Errno(e);
        <a id="L51"></a>}
    <a id="L52"></a>}

    <a id="L54"></a>sa, _ := syscall.Getsockname(s);
    <a id="L55"></a>laddr := toAddr(sa);
    <a id="L56"></a>sa, _ = syscall.Getpeername(s);
    <a id="L57"></a>raddr := toAddr(sa);

    <a id="L59"></a>fd, err = newFD(s, f, p, net, laddr, raddr);
    <a id="L60"></a>if err != nil {
        <a id="L61"></a>syscall.Close(s);
        <a id="L62"></a>return nil, err;
    <a id="L63"></a>}

    <a id="L65"></a>return fd, nil;
<a id="L66"></a>}

<a id="L68"></a>func setsockoptInt(fd, level, opt int, value int) os.Error {
    <a id="L69"></a>return os.NewSyscallError(&#34;setsockopt&#34;, syscall.SetsockoptInt(fd, level, opt, value))
<a id="L70"></a>}

<a id="L72"></a>func setsockoptNsec(fd, level, opt int, nsec int64) os.Error {
    <a id="L73"></a>var tv = syscall.NsecToTimeval(nsec);
    <a id="L74"></a>return os.NewSyscallError(&#34;setsockopt&#34;, syscall.SetsockoptTimeval(fd, level, opt, &amp;tv));
<a id="L75"></a>}

<a id="L77"></a>func setReadBuffer(fd *netFD, bytes int) os.Error {
    <a id="L78"></a>return setsockoptInt(fd.fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, bytes)
<a id="L79"></a>}

<a id="L81"></a>func setWriteBuffer(fd *netFD, bytes int) os.Error {
    <a id="L82"></a>return setsockoptInt(fd.fd, syscall.SOL_SOCKET, syscall.SO_SNDBUF, bytes)
<a id="L83"></a>}

<a id="L85"></a>func setReadTimeout(fd *netFD, nsec int64) os.Error {
    <a id="L86"></a>fd.rdeadline_delta = nsec;
    <a id="L87"></a>return nil;
<a id="L88"></a>}

<a id="L90"></a>func setWriteTimeout(fd *netFD, nsec int64) os.Error {
    <a id="L91"></a>fd.wdeadline_delta = nsec;
    <a id="L92"></a>return nil;
<a id="L93"></a>}

<a id="L95"></a>func setTimeout(fd *netFD, nsec int64) os.Error {
    <a id="L96"></a>if e := setReadTimeout(fd, nsec); e != nil {
        <a id="L97"></a>return e
    <a id="L98"></a>}
    <a id="L99"></a>return setWriteTimeout(fd, nsec);
<a id="L100"></a>}

<a id="L102"></a>func setReuseAddr(fd *netFD, reuse bool) os.Error {
    <a id="L103"></a>return setsockoptInt(fd.fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, boolint(reuse))
<a id="L104"></a>}

<a id="L106"></a>func bindToDevice(fd *netFD, dev string) os.Error {
    <a id="L107"></a><span class="comment">// TODO(rsc): call setsockopt with null-terminated string pointer</span>
    <a id="L108"></a>return os.EINVAL
<a id="L109"></a>}

<a id="L111"></a>func setDontRoute(fd *netFD, dontroute bool) os.Error {
    <a id="L112"></a>return setsockoptInt(fd.fd, syscall.SOL_SOCKET, syscall.SO_DONTROUTE, boolint(dontroute))
<a id="L113"></a>}

<a id="L115"></a>func setKeepAlive(fd *netFD, keepalive bool) os.Error {
    <a id="L116"></a>return setsockoptInt(fd.fd, syscall.SOL_SOCKET, syscall.SO_KEEPALIVE, boolint(keepalive))
<a id="L117"></a>}

<a id="L119"></a>func setLinger(fd *netFD, sec int) os.Error {
    <a id="L120"></a>var l syscall.Linger;
    <a id="L121"></a>if sec &gt;= 0 {
        <a id="L122"></a>l.Onoff = 1;
        <a id="L123"></a>l.Linger = int32(sec);
    <a id="L124"></a>} else {
        <a id="L125"></a>l.Onoff = 0;
        <a id="L126"></a>l.Linger = 0;
    <a id="L127"></a>}
    <a id="L128"></a>e := syscall.SetsockoptLinger(fd.fd, syscall.SOL_SOCKET, syscall.SO_LINGER, &amp;l);
    <a id="L129"></a>return os.NewSyscallError(&#34;setsockopt&#34;, e);
<a id="L130"></a>}

<a id="L132"></a>type UnknownSocketError struct {
    <a id="L133"></a>sa syscall.Sockaddr;
<a id="L134"></a>}

<a id="L136"></a>func (e *UnknownSocketError) String() string {
    <a id="L137"></a>return &#34;unknown socket address type &#34; + reflect.Typeof(e.sa).String()
<a id="L138"></a>}

<a id="L140"></a>func sockaddrToString(sa syscall.Sockaddr) (name string, err os.Error) {
    <a id="L141"></a>switch a := sa.(type) {
    <a id="L142"></a>case *syscall.SockaddrInet4:
        <a id="L143"></a>return joinHostPort(IP(&amp;a.Addr).String(), itoa(a.Port)), nil
    <a id="L144"></a>case *syscall.SockaddrInet6:
        <a id="L145"></a>return joinHostPort(IP(&amp;a.Addr).String(), itoa(a.Port)), nil
    <a id="L146"></a>case *syscall.SockaddrUnix:
        <a id="L147"></a>return a.Name, nil
    <a id="L148"></a>}

    <a id="L150"></a>return &#34;&#34;, &amp;UnknownSocketError{sa};
<a id="L151"></a>}
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
