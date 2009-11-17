<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/testing/iotest/reader.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Source file /src/pkg/testing/iotest/reader.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The iotest package implements Readers and Writers</span>
<a id="L6"></a><span class="comment">// useful only for testing.</span>
<a id="L7"></a>package iotest

<a id="L9"></a>import (
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;bytes&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// OneByteReader returns a Reader that implements</span>
<a id="L16"></a><span class="comment">// each non-empty Read by reading one byte from r.</span>
<a id="L17"></a>func OneByteReader(r io.Reader) io.Reader { return &amp;oneByteReader{r} }

<a id="L19"></a>type oneByteReader struct {
    <a id="L20"></a>r io.Reader;
<a id="L21"></a>}

<a id="L23"></a>func (r *oneByteReader) Read(p []byte) (int, os.Error) {
    <a id="L24"></a>if len(p) == 0 {
        <a id="L25"></a>return 0, nil
    <a id="L26"></a>}
    <a id="L27"></a>return r.r.Read(p[0:1]);
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// HalfReader returns a Reader that implements Read</span>
<a id="L31"></a><span class="comment">// by reading half as many requested bytes from r.</span>
<a id="L32"></a>func HalfReader(r io.Reader) io.Reader { return &amp;halfReader{r} }

<a id="L34"></a>type halfReader struct {
    <a id="L35"></a>r io.Reader;
<a id="L36"></a>}

<a id="L38"></a>func (r *halfReader) Read(p []byte) (int, os.Error) {
    <a id="L39"></a>return r.r.Read(p[0 : (len(p)+1)/2])
<a id="L40"></a>}


<a id="L43"></a><span class="comment">// DataErrReader returns a Reader that returns the final</span>
<a id="L44"></a><span class="comment">// error with the last data read, instead of by itself with</span>
<a id="L45"></a><span class="comment">// zero bytes of data.</span>
<a id="L46"></a>func DataErrReader(r io.Reader) io.Reader { return &amp;dataErrReader{r, nil, make([]byte, 1024)} }

<a id="L48"></a>type dataErrReader struct {
    <a id="L49"></a>r      io.Reader;
    <a id="L50"></a>unread []byte;
    <a id="L51"></a>data   []byte;
<a id="L52"></a>}

<a id="L54"></a>func (r *dataErrReader) Read(p []byte) (n int, err os.Error) {
    <a id="L55"></a><span class="comment">// loop because first call needs two reads:</span>
    <a id="L56"></a><span class="comment">// one to get data and a second to look for an error.</span>
    <a id="L57"></a>for {
        <a id="L58"></a>if len(r.unread) == 0 {
            <a id="L59"></a>n1, err1 := r.r.Read(r.data);
            <a id="L60"></a>r.unread = r.data[0:n1];
            <a id="L61"></a>err = err1;
        <a id="L62"></a>}
        <a id="L63"></a>if n &gt; 0 {
            <a id="L64"></a>break
        <a id="L65"></a>}
        <a id="L66"></a>n = bytes.Copy(p, r.unread);
        <a id="L67"></a>r.unread = r.unread[n:len(r.unread)];
    <a id="L68"></a>}
    <a id="L69"></a>return;
<a id="L70"></a>}
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
