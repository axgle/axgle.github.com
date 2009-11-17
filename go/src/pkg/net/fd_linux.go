<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/fd_linux.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/fd_linux.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Waiting for FDs via epoll(7).</span>

<a id="L7"></a>package net

<a id="L9"></a>import (
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
<a id="L12"></a>)

<a id="L14"></a>const (
    <a id="L15"></a>readFlags  = syscall.EPOLLIN | syscall.EPOLLRDHUP;
    <a id="L16"></a>writeFlags = syscall.EPOLLOUT;
<a id="L17"></a>)

<a id="L19"></a>type pollster struct {
    <a id="L20"></a>epfd int;

    <a id="L22"></a><span class="comment">// Events we&#39;re already waiting for</span>
    <a id="L23"></a>events map[int]uint32;
<a id="L24"></a>}

<a id="L26"></a>func newpollster() (p *pollster, err os.Error) {
    <a id="L27"></a>p = new(pollster);
    <a id="L28"></a>var e int;

    <a id="L30"></a><span class="comment">// The arg to epoll_create is a hint to the kernel</span>
    <a id="L31"></a><span class="comment">// about the number of FDs we will care about.</span>
    <a id="L32"></a><span class="comment">// We don&#39;t know.</span>
    <a id="L33"></a>if p.epfd, e = syscall.EpollCreate(16); e != 0 {
        <a id="L34"></a>return nil, os.NewSyscallError(&#34;epoll_create&#34;, e)
    <a id="L35"></a>}
    <a id="L36"></a>p.events = make(map[int]uint32);
    <a id="L37"></a>return p, nil;
<a id="L38"></a>}

<a id="L40"></a>func (p *pollster) AddFD(fd int, mode int, repeat bool) os.Error {
    <a id="L41"></a>var ev syscall.EpollEvent;
    <a id="L42"></a>var already bool;
    <a id="L43"></a>ev.Fd = int32(fd);
    <a id="L44"></a>ev.Events, already = p.events[fd];
    <a id="L45"></a>if !repeat {
        <a id="L46"></a>ev.Events |= syscall.EPOLLONESHOT
    <a id="L47"></a>}
    <a id="L48"></a>if mode == &#39;r&#39; {
        <a id="L49"></a>ev.Events |= readFlags
    <a id="L50"></a>} else {
        <a id="L51"></a>ev.Events |= writeFlags
    <a id="L52"></a>}

    <a id="L54"></a>var op int;
    <a id="L55"></a>if already {
        <a id="L56"></a>op = syscall.EPOLL_CTL_MOD
    <a id="L57"></a>} else {
        <a id="L58"></a>op = syscall.EPOLL_CTL_ADD
    <a id="L59"></a>}
    <a id="L60"></a>if e := syscall.EpollCtl(p.epfd, op, fd, &amp;ev); e != 0 {
        <a id="L61"></a>return os.NewSyscallError(&#34;epoll_ctl&#34;, e)
    <a id="L62"></a>}
    <a id="L63"></a>p.events[fd] = ev.Events;
    <a id="L64"></a>return nil;
<a id="L65"></a>}

<a id="L67"></a>func (p *pollster) StopWaiting(fd int, bits uint) {
    <a id="L68"></a>events, already := p.events[fd];
    <a id="L69"></a>if !already {
        <a id="L70"></a>print(&#34;Epoll unexpected fd=&#34;, fd, &#34;\n&#34;);
        <a id="L71"></a>return;
    <a id="L72"></a>}

    <a id="L74"></a><span class="comment">// If syscall.EPOLLONESHOT is not set, the wait</span>
    <a id="L75"></a><span class="comment">// is a repeating wait, so don&#39;t change it.</span>
    <a id="L76"></a>if events&amp;syscall.EPOLLONESHOT == 0 {
        <a id="L77"></a>return
    <a id="L78"></a>}

    <a id="L80"></a><span class="comment">// Disable the given bits.</span>
    <a id="L81"></a><span class="comment">// If we&#39;re still waiting for other events, modify the fd</span>
    <a id="L82"></a><span class="comment">// event in the kernel.  Otherwise, delete it.</span>
    <a id="L83"></a>events &amp;= ^uint32(bits);
    <a id="L84"></a>if int32(events)&amp;^syscall.EPOLLONESHOT != 0 {
        <a id="L85"></a>var ev syscall.EpollEvent;
        <a id="L86"></a>ev.Fd = int32(fd);
        <a id="L87"></a>ev.Events = events;
        <a id="L88"></a>if e := syscall.EpollCtl(p.epfd, syscall.EPOLL_CTL_MOD, fd, &amp;ev); e != 0 {
            <a id="L89"></a>print(&#34;Epoll modify fd=&#34;, fd, &#34;: &#34;, os.Errno(e).String(), &#34;\n&#34;)
        <a id="L90"></a>}
        <a id="L91"></a>p.events[fd] = events;
    <a id="L92"></a>} else {
        <a id="L93"></a>if e := syscall.EpollCtl(p.epfd, syscall.EPOLL_CTL_DEL, fd, nil); e != 0 {
            <a id="L94"></a>print(&#34;Epoll delete fd=&#34;, fd, &#34;: &#34;, os.Errno(e).String(), &#34;\n&#34;)
        <a id="L95"></a>}
        <a id="L96"></a>p.events[fd] = 0, false;
    <a id="L97"></a>}
<a id="L98"></a>}

<a id="L100"></a>func (p *pollster) DelFD(fd int, mode int) {
    <a id="L101"></a>if mode == &#39;r&#39; {
        <a id="L102"></a>p.StopWaiting(fd, readFlags)
    <a id="L103"></a>} else {
        <a id="L104"></a>p.StopWaiting(fd, writeFlags)
    <a id="L105"></a>}
<a id="L106"></a>}

<a id="L108"></a>func (p *pollster) WaitFD(nsec int64) (fd int, mode int, err os.Error) {
    <a id="L109"></a><span class="comment">// Get an event.</span>
    <a id="L110"></a>var evarray [1]syscall.EpollEvent;
    <a id="L111"></a>ev := &amp;evarray[0];
    <a id="L112"></a>var msec int = -1;
    <a id="L113"></a>if nsec &gt; 0 {
        <a id="L114"></a>msec = int((nsec + 1e6 - 1) / 1e6)
    <a id="L115"></a>}
    <a id="L116"></a>n, e := syscall.EpollWait(p.epfd, &amp;evarray, msec);
    <a id="L117"></a>for e == syscall.EAGAIN || e == syscall.EINTR {
        <a id="L118"></a>n, e = syscall.EpollWait(p.epfd, &amp;evarray, msec)
    <a id="L119"></a>}
    <a id="L120"></a>if e != 0 {
        <a id="L121"></a>return -1, 0, os.NewSyscallError(&#34;epoll_wait&#34;, e)
    <a id="L122"></a>}
    <a id="L123"></a>if n == 0 {
        <a id="L124"></a>return -1, 0, nil
    <a id="L125"></a>}
    <a id="L126"></a>fd = int(ev.Fd);

    <a id="L128"></a>if ev.Events&amp;writeFlags != 0 {
        <a id="L129"></a>p.StopWaiting(fd, writeFlags);
        <a id="L130"></a>return fd, &#39;w&#39;, nil;
    <a id="L131"></a>}
    <a id="L132"></a>if ev.Events&amp;readFlags != 0 {
        <a id="L133"></a>p.StopWaiting(fd, readFlags);
        <a id="L134"></a>return fd, &#39;r&#39;, nil;
    <a id="L135"></a>}

    <a id="L137"></a><span class="comment">// Other events are error conditions - wake whoever is waiting.</span>
    <a id="L138"></a>events, _ := p.events[fd];
    <a id="L139"></a>if events&amp;writeFlags != 0 {
        <a id="L140"></a>p.StopWaiting(fd, writeFlags);
        <a id="L141"></a>return fd, &#39;w&#39;, nil;
    <a id="L142"></a>}
    <a id="L143"></a>p.StopWaiting(fd, readFlags);
    <a id="L144"></a>return fd, &#39;r&#39;, nil;
<a id="L145"></a>}

<a id="L147"></a>func (p *pollster) Close() os.Error {
    <a id="L148"></a>return os.NewSyscallError(&#34;close&#34;, syscall.Close(p.epfd))
<a id="L149"></a>}
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
