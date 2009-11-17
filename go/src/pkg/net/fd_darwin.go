<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/fd_darwin.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/fd_darwin.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Waiting for FDs via kqueue/kevent.</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
<a id="L12"></a>)

<a id="L14"></a>type pollster struct {
    <a id="L15"></a>kq       int;
    <a id="L16"></a>eventbuf [10]syscall.Kevent_t;
    <a id="L17"></a>events   []syscall.Kevent_t;
<a id="L18"></a>}

<a id="L20"></a>func newpollster() (p *pollster, err os.Error) {
    <a id="L21"></a>p = new(pollster);
    <a id="L22"></a>var e int;
    <a id="L23"></a>if p.kq, e = syscall.Kqueue(); e != 0 {
        <a id="L24"></a>return nil, os.NewSyscallError(&#34;kqueue&#34;, e)
    <a id="L25"></a>}
    <a id="L26"></a>p.events = p.eventbuf[0:0];
    <a id="L27"></a>return p, nil;
<a id="L28"></a>}

<a id="L30"></a>func (p *pollster) AddFD(fd int, mode int, repeat bool) os.Error {
    <a id="L31"></a>var kmode int;
    <a id="L32"></a>if mode == &#39;r&#39; {
        <a id="L33"></a>kmode = syscall.EVFILT_READ
    <a id="L34"></a>} else {
        <a id="L35"></a>kmode = syscall.EVFILT_WRITE
    <a id="L36"></a>}
    <a id="L37"></a>var events [1]syscall.Kevent_t;
    <a id="L38"></a>ev := &amp;events[0];
    <a id="L39"></a><span class="comment">// EV_ADD - add event to kqueue list</span>
    <a id="L40"></a><span class="comment">// EV_RECEIPT - generate fake EV_ERROR as result of add,</span>
    <a id="L41"></a><span class="comment">//	rather than waiting for real event</span>
    <a id="L42"></a><span class="comment">// EV_ONESHOT - delete the event the first time it triggers</span>
    <a id="L43"></a>flags := syscall.EV_ADD | syscall.EV_RECEIPT;
    <a id="L44"></a>if !repeat {
        <a id="L45"></a>flags |= syscall.EV_ONESHOT
    <a id="L46"></a>}
    <a id="L47"></a>syscall.SetKevent(ev, fd, kmode, flags);

    <a id="L49"></a>n, e := syscall.Kevent(p.kq, &amp;events, &amp;events, nil);
    <a id="L50"></a>if e != 0 {
        <a id="L51"></a>return os.NewSyscallError(&#34;kevent&#34;, e)
    <a id="L52"></a>}
    <a id="L53"></a>if n != 1 || (ev.Flags&amp;syscall.EV_ERROR) == 0 || int(ev.Ident) != fd || int(ev.Filter) != kmode {
        <a id="L54"></a>return os.ErrorString(&#34;kqueue phase error&#34;)
    <a id="L55"></a>}
    <a id="L56"></a>if ev.Data != 0 {
        <a id="L57"></a>return os.Errno(int(ev.Data))
    <a id="L58"></a>}
    <a id="L59"></a>return nil;
<a id="L60"></a>}

<a id="L62"></a>func (p *pollster) DelFD(fd int, mode int) {
    <a id="L63"></a>var kmode int;
    <a id="L64"></a>if mode == &#39;r&#39; {
        <a id="L65"></a>kmode = syscall.EVFILT_READ
    <a id="L66"></a>} else {
        <a id="L67"></a>kmode = syscall.EVFILT_WRITE
    <a id="L68"></a>}
    <a id="L69"></a>var events [1]syscall.Kevent_t;
    <a id="L70"></a>ev := &amp;events[0];
    <a id="L71"></a><span class="comment">// EV_DELETE - delete event from kqueue list</span>
    <a id="L72"></a><span class="comment">// EV_RECEIPT - generate fake EV_ERROR as result of add,</span>
    <a id="L73"></a><span class="comment">//	rather than waiting for real event</span>
    <a id="L74"></a>syscall.SetKevent(ev, fd, kmode, syscall.EV_DELETE|syscall.EV_RECEIPT);
    <a id="L75"></a>syscall.Kevent(p.kq, &amp;events, &amp;events, nil);
<a id="L76"></a>}

<a id="L78"></a>func (p *pollster) WaitFD(nsec int64) (fd int, mode int, err os.Error) {
    <a id="L79"></a>var t *syscall.Timespec;
    <a id="L80"></a>for len(p.events) == 0 {
        <a id="L81"></a>if nsec &gt; 0 {
            <a id="L82"></a>if t == nil {
                <a id="L83"></a>t = new(syscall.Timespec)
            <a id="L84"></a>}
            <a id="L85"></a>*t = syscall.NsecToTimespec(nsec);
        <a id="L86"></a>}
        <a id="L87"></a>nn, e := syscall.Kevent(p.kq, nil, &amp;p.eventbuf, t);
        <a id="L88"></a>if e != 0 {
            <a id="L89"></a>if e == syscall.EINTR {
                <a id="L90"></a>continue
            <a id="L91"></a>}
            <a id="L92"></a>return -1, 0, os.NewSyscallError(&#34;kevent&#34;, e);
        <a id="L93"></a>}
        <a id="L94"></a>if nn == 0 {
            <a id="L95"></a>return -1, 0, nil
        <a id="L96"></a>}
        <a id="L97"></a>p.events = p.eventbuf[0:nn];
    <a id="L98"></a>}
    <a id="L99"></a>ev := &amp;p.events[0];
    <a id="L100"></a>p.events = p.events[1:len(p.events)];
    <a id="L101"></a>fd = int(ev.Ident);
    <a id="L102"></a>if ev.Filter == syscall.EVFILT_READ {
        <a id="L103"></a>mode = &#39;r&#39;
    <a id="L104"></a>} else {
        <a id="L105"></a>mode = &#39;w&#39;
    <a id="L106"></a>}
    <a id="L107"></a>return fd, mode, nil;
<a id="L108"></a>}

<a id="L110"></a>func (p *pollster) Close() os.Error { return os.NewSyscallError(&#34;close&#34;, syscall.Close(p.kq)) }
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
