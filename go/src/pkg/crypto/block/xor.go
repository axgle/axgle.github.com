<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/xor.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/xor.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Encrypt/decrypt data by xor with a pseudo-random data stream.</span>

<a id="L7"></a>package block

<a id="L9"></a>import (
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;os&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// A dataStream is an interface to an unending stream of data,</span>
<a id="L15"></a><span class="comment">// used by XorReader and XorWriter to model a pseudo-random generator.</span>
<a id="L16"></a><span class="comment">// Calls to Next() return sequential blocks of data from the stream.</span>
<a id="L17"></a><span class="comment">// Each call must return at least one byte: there is no EOF.</span>
<a id="L18"></a>type dataStream interface {
    <a id="L19"></a>Next() []byte;
<a id="L20"></a>}

<a id="L22"></a>type xorReader struct {
    <a id="L23"></a>r    io.Reader;
    <a id="L24"></a>rand dataStream; <span class="comment">// pseudo-random</span>
    <a id="L25"></a>buf  []byte;     <span class="comment">// data available from last call to rand</span>
<a id="L26"></a>}

<a id="L28"></a>func newXorReader(rand dataStream, r io.Reader) io.Reader {
    <a id="L29"></a>x := new(xorReader);
    <a id="L30"></a>x.r = r;
    <a id="L31"></a>x.rand = rand;
    <a id="L32"></a>return x;
<a id="L33"></a>}

<a id="L35"></a>func (x *xorReader) Read(p []byte) (n int, err os.Error) {
    <a id="L36"></a>n, err = x.r.Read(p);

    <a id="L38"></a><span class="comment">// xor input with stream.</span>
    <a id="L39"></a>bp := 0;
    <a id="L40"></a>buf := x.buf;
    <a id="L41"></a>for i := 0; i &lt; n; i++ {
        <a id="L42"></a>if bp &gt;= len(buf) {
            <a id="L43"></a>buf = x.rand.Next();
            <a id="L44"></a>bp = 0;
        <a id="L45"></a>}
        <a id="L46"></a>p[i] ^= buf[bp];
        <a id="L47"></a>bp++;
    <a id="L48"></a>}
    <a id="L49"></a>x.buf = buf[bp:len(buf)];
    <a id="L50"></a>return n, err;
<a id="L51"></a>}

<a id="L53"></a>type xorWriter struct {
    <a id="L54"></a>w     io.Writer;
    <a id="L55"></a>rand  dataStream; <span class="comment">// pseudo-random</span>
    <a id="L56"></a>buf   []byte;     <span class="comment">// last buffer returned by rand</span>
    <a id="L57"></a>extra []byte;     <span class="comment">// extra random data (use before buf)</span>
    <a id="L58"></a>work  []byte;     <span class="comment">// work space</span>
<a id="L59"></a>}

<a id="L61"></a>func newXorWriter(rand dataStream, w io.Writer) io.Writer {
    <a id="L62"></a>x := new(xorWriter);
    <a id="L63"></a>x.w = w;
    <a id="L64"></a>x.rand = rand;
    <a id="L65"></a>x.work = make([]byte, 4096);
    <a id="L66"></a>return x;
<a id="L67"></a>}

<a id="L69"></a>func (x *xorWriter) Write(p []byte) (n int, err os.Error) {
    <a id="L70"></a>for len(p) &gt; 0 {
        <a id="L71"></a><span class="comment">// Determine next chunk of random data</span>
        <a id="L72"></a><span class="comment">// and xor with p into x.work.</span>
        <a id="L73"></a>var chunk []byte;
        <a id="L74"></a>m := len(p);
        <a id="L75"></a>if nn := len(x.extra); nn &gt; 0 {
            <a id="L76"></a><span class="comment">// extra points into work, so edit directly</span>
            <a id="L77"></a>if m &gt; nn {
                <a id="L78"></a>m = nn
            <a id="L79"></a>}
            <a id="L80"></a>for i := 0; i &lt; m; i++ {
                <a id="L81"></a>x.extra[i] ^= p[i]
            <a id="L82"></a>}
            <a id="L83"></a>chunk = x.extra[0:m];
        <a id="L84"></a>} else {
            <a id="L85"></a><span class="comment">// xor p ^ buf into work, refreshing buf as needed</span>
            <a id="L86"></a>if nn := len(x.work); m &gt; nn {
                <a id="L87"></a>m = nn
            <a id="L88"></a>}
            <a id="L89"></a>bp := 0;
            <a id="L90"></a>buf := x.buf;
            <a id="L91"></a>for i := 0; i &lt; m; i++ {
                <a id="L92"></a>if bp &gt;= len(buf) {
                    <a id="L93"></a>buf = x.rand.Next();
                    <a id="L94"></a>bp = 0;
                <a id="L95"></a>}
                <a id="L96"></a>x.work[i] = buf[bp] ^ p[i];
                <a id="L97"></a>bp++;
            <a id="L98"></a>}
            <a id="L99"></a>x.buf = buf[bp:len(buf)];
            <a id="L100"></a>chunk = x.work[0:m];
        <a id="L101"></a>}

        <a id="L103"></a><span class="comment">// Write chunk.</span>
        <a id="L104"></a>var nn int;
        <a id="L105"></a>nn, err = x.w.Write(chunk);
        <a id="L106"></a>if nn != len(chunk) &amp;&amp; err == nil {
            <a id="L107"></a>err = io.ErrShortWrite
        <a id="L108"></a>}
        <a id="L109"></a>if nn &lt; len(chunk) {
            <a id="L110"></a><span class="comment">// Reconstruct the random bits from the unwritten</span>
            <a id="L111"></a><span class="comment">// data and save them for next time.</span>
            <a id="L112"></a>for i := nn; i &lt; m; i++ {
                <a id="L113"></a>chunk[i] ^= p[i]
            <a id="L114"></a>}
            <a id="L115"></a>x.extra = chunk[nn:len(chunk)];
        <a id="L116"></a>}
        <a id="L117"></a>n += nn;
        <a id="L118"></a>if err != nil {
            <a id="L119"></a>return
        <a id="L120"></a>}
        <a id="L121"></a>p = p[m:len(p)];
    <a id="L122"></a>}
    <a id="L123"></a>return;
<a id="L124"></a>}
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
