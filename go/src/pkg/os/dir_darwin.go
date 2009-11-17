<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/os/dir_darwin.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/os/dir_darwin.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package os

<a id="L7"></a>import (
    <a id="L8"></a>&#34;syscall&#34;;
    <a id="L9"></a>&#34;unsafe&#34;;
<a id="L10"></a>)

<a id="L12"></a>const (
    <a id="L13"></a>blockSize = 4096; <span class="comment">// TODO(r): use statfs</span>
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// Readdirnames reads the contents of the directory associated with file and</span>
<a id="L17"></a><span class="comment">// returns an array of up to count names, in directory order.  Subsequent</span>
<a id="L18"></a><span class="comment">// calls on the same file will yield further names.</span>
<a id="L19"></a><span class="comment">// A negative count means to read until EOF.</span>
<a id="L20"></a><span class="comment">// Readdirnames returns the array and an Error, if any.</span>
<a id="L21"></a>func (file *File) Readdirnames(count int) (names []string, err Error) {
    <a id="L22"></a><span class="comment">// If this file has no dirinfo, create one.</span>
    <a id="L23"></a>if file.dirinfo == nil {
        <a id="L24"></a>file.dirinfo = new(dirInfo);
        <a id="L25"></a><span class="comment">// The buffer must be at least a block long.</span>
        <a id="L26"></a><span class="comment">// TODO(r): use fstatfs to find fs block size.</span>
        <a id="L27"></a>file.dirinfo.buf = make([]byte, blockSize);
    <a id="L28"></a>}
    <a id="L29"></a>d := file.dirinfo;
    <a id="L30"></a>size := count;
    <a id="L31"></a>if size &lt; 0 {
        <a id="L32"></a>size = 100
    <a id="L33"></a>}
    <a id="L34"></a>names = make([]string, 0, size); <span class="comment">// Empty with room to grow.</span>
    <a id="L35"></a>for count != 0 {
        <a id="L36"></a><span class="comment">// Refill the buffer if necessary</span>
        <a id="L37"></a>if d.bufp &gt;= d.nbuf {
            <a id="L38"></a>var errno int;
            <a id="L39"></a>d.bufp = 0;
            <a id="L40"></a><span class="comment">// Final argument is (basep *uintptr) and the syscall doesn&#39;t take nil.</span>
            <a id="L41"></a>d.nbuf, errno = syscall.Getdirentries(file.fd, d.buf, new(uintptr));
            <a id="L42"></a>if errno != 0 {
                <a id="L43"></a>d.nbuf = 0;
                <a id="L44"></a>return names, NewSyscallError(&#34;getdirentries&#34;, errno);
            <a id="L45"></a>}
            <a id="L46"></a>if d.nbuf &lt;= 0 {
                <a id="L47"></a>break <span class="comment">// EOF</span>
            <a id="L48"></a>}
        <a id="L49"></a>}
        <a id="L50"></a><span class="comment">// Drain the buffer</span>
        <a id="L51"></a>for count != 0 &amp;&amp; d.bufp &lt; d.nbuf {
            <a id="L52"></a>dirent := (*syscall.Dirent)(unsafe.Pointer(&amp;d.buf[d.bufp]));
            <a id="L53"></a>if dirent.Reclen == 0 {
                <a id="L54"></a>d.bufp = d.nbuf;
                <a id="L55"></a>break;
            <a id="L56"></a>}
            <a id="L57"></a>d.bufp += int(dirent.Reclen);
            <a id="L58"></a>if dirent.Ino == 0 { <span class="comment">// File absent in directory.</span>
                <a id="L59"></a>continue
            <a id="L60"></a>}
            <a id="L61"></a>bytes := (*[len(dirent.Name)]byte)(unsafe.Pointer(&amp;dirent.Name[0]));
            <a id="L62"></a>var name = string(bytes[0:dirent.Namlen]);
            <a id="L63"></a>if name == &#34;.&#34; || name == &#34;..&#34; { <span class="comment">// Useless names</span>
                <a id="L64"></a>continue
            <a id="L65"></a>}
            <a id="L66"></a>count--;
            <a id="L67"></a>if len(names) == cap(names) {
                <a id="L68"></a>nnames := make([]string, len(names), 2*len(names));
                <a id="L69"></a>for i := 0; i &lt; len(names); i++ {
                    <a id="L70"></a>nnames[i] = names[i]
                <a id="L71"></a>}
                <a id="L72"></a>names = nnames;
            <a id="L73"></a>}
            <a id="L74"></a>names = names[0 : len(names)+1];
            <a id="L75"></a>names[len(names)-1] = name;
        <a id="L76"></a>}
    <a id="L77"></a>}
    <a id="L78"></a>return names, nil;
<a id="L79"></a>}
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
