<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/hash/adler32/adler32.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/hash/adler32/adler32.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements the Adler-32 checksum.</span>
<a id="L6"></a><span class="comment">// Defined in RFC 1950:</span>
<a id="L7"></a><span class="comment">//	Adler-32 is composed of two sums accumulated per byte: s1 is</span>
<a id="L8"></a><span class="comment">//	the sum of all bytes, s2 is the sum of all s1 values. Both sums</span>
<a id="L9"></a><span class="comment">//	are done modulo 65521. s1 is initialized to 1, s2 to zero.  The</span>
<a id="L10"></a><span class="comment">//	Adler-32 checksum is stored as s2*65536 + s1 in most-</span>
<a id="L11"></a><span class="comment">//	significant-byte first (network) order.</span>
<a id="L12"></a>package adler32

<a id="L14"></a>import (
    <a id="L15"></a>&#34;hash&#34;;
    <a id="L16"></a>&#34;os&#34;;
<a id="L17"></a>)

<a id="L19"></a>const (
    <a id="L20"></a>mod = 65521;
<a id="L21"></a>)

<a id="L23"></a><span class="comment">// The size of an Adler-32 checksum in bytes.</span>
<a id="L24"></a>const Size = 4

<a id="L26"></a><span class="comment">// digest represents the partial evaluation of a checksum.</span>
<a id="L27"></a>type digest struct {
    <a id="L28"></a><span class="comment">// invariant: (a &lt; mod &amp;&amp; b &lt; mod) || a &lt;= b</span>
    <a id="L29"></a><span class="comment">// invariant: a + b + 255 &lt;= 0xffffffff</span>
    <a id="L30"></a>a, b uint32;
<a id="L31"></a>}

<a id="L33"></a>func (d *digest) Reset() { d.a, d.b = 1, 0 }

<a id="L35"></a><span class="comment">// New returns a new hash.Hash32 computing the Adler-32 checksum.</span>
<a id="L36"></a>func New() hash.Hash32 {
    <a id="L37"></a>d := new(digest);
    <a id="L38"></a>d.Reset();
    <a id="L39"></a>return d;
<a id="L40"></a>}

<a id="L42"></a>func (d *digest) Size() int { return Size }

<a id="L44"></a><span class="comment">// Add p to the running checksum a, b.</span>
<a id="L45"></a>func update(a, b uint32, p []byte) (aa, bb uint32) {
    <a id="L46"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L47"></a>a += uint32(p[i]);
        <a id="L48"></a>b += a;
        <a id="L49"></a><span class="comment">// invariant: a &lt;= b</span>
        <a id="L50"></a>if b &gt; (0xffffffff-255)/2 {
            <a id="L51"></a>a %= mod;
            <a id="L52"></a>b %= mod;
            <a id="L53"></a><span class="comment">// invariant: a &lt; mod &amp;&amp; b &lt; mod</span>
        <a id="L54"></a>} else {
            <a id="L55"></a><span class="comment">// invariant: a + b + 255 &lt;= 2 * b + 255 &lt;= 0xffffffff</span>
        <a id="L56"></a>}
    <a id="L57"></a>}
    <a id="L58"></a>return a, b;
<a id="L59"></a>}

<a id="L61"></a><span class="comment">// Return the 32-bit checksum corresponding to a, b.</span>
<a id="L62"></a>func finish(a, b uint32) uint32 {
    <a id="L63"></a>if b &gt;= mod {
        <a id="L64"></a>a %= mod;
        <a id="L65"></a>b %= mod;
    <a id="L66"></a>}
    <a id="L67"></a>return b&lt;&lt;16 | a;
<a id="L68"></a>}

<a id="L70"></a>func (d *digest) Write(p []byte) (nn int, err os.Error) {
    <a id="L71"></a>d.a, d.b = update(d.a, d.b, p);
    <a id="L72"></a>return len(p), nil;
<a id="L73"></a>}

<a id="L75"></a>func (d *digest) Sum32() uint32 { return finish(d.a, d.b) }

<a id="L77"></a>func (d *digest) Sum() []byte {
    <a id="L78"></a>p := make([]byte, 4);
    <a id="L79"></a>s := d.Sum32();
    <a id="L80"></a>p[0] = byte(s &gt;&gt; 24);
    <a id="L81"></a>p[1] = byte(s &gt;&gt; 16);
    <a id="L82"></a>p[2] = byte(s &gt;&gt; 8);
    <a id="L83"></a>p[3] = byte(s);
    <a id="L84"></a>return p;
<a id="L85"></a>}

<a id="L87"></a><span class="comment">// Checksum returns the Adler-32 checksum of data.</span>
<a id="L88"></a>func Checksum(data []byte) uint32 { return finish(update(1, 0, data)) }
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
