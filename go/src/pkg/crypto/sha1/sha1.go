<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/sha1/sha1.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/sha1/sha1.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements the SHA1 hash algorithm as defined in RFC 3174.</span>
<a id="L6"></a>package sha1

<a id="L8"></a>import (
    <a id="L9"></a>&#34;hash&#34;;
    <a id="L10"></a>&#34;os&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// The size of a SHA1 checksum in bytes.</span>
<a id="L14"></a>const Size = 20

<a id="L16"></a>const (
    <a id="L17"></a>_Chunk = 64;
    <a id="L18"></a>_Init0 = 0x67452301;
    <a id="L19"></a>_Init1 = 0xEFCDAB89;
    <a id="L20"></a>_Init2 = 0x98BADCFE;
    <a id="L21"></a>_Init3 = 0x10325476;
    <a id="L22"></a>_Init4 = 0xC3D2E1F0;
<a id="L23"></a>)

<a id="L25"></a><span class="comment">// digest represents the partial evaluation of a checksum.</span>
<a id="L26"></a>type digest struct {
    <a id="L27"></a>h   [5]uint32;
    <a id="L28"></a>x   [_Chunk]byte;
    <a id="L29"></a>nx  int;
    <a id="L30"></a>len uint64;
<a id="L31"></a>}

<a id="L33"></a>func (d *digest) Reset() {
    <a id="L34"></a>d.h[0] = _Init0;
    <a id="L35"></a>d.h[1] = _Init1;
    <a id="L36"></a>d.h[2] = _Init2;
    <a id="L37"></a>d.h[3] = _Init3;
    <a id="L38"></a>d.h[4] = _Init4;
    <a id="L39"></a>d.nx = 0;
    <a id="L40"></a>d.len = 0;
<a id="L41"></a>}

<a id="L43"></a><span class="comment">// New returns a new hash.Hash computing the SHA1 checksum.</span>
<a id="L44"></a>func New() hash.Hash {
    <a id="L45"></a>d := new(digest);
    <a id="L46"></a>d.Reset();
    <a id="L47"></a>return d;
<a id="L48"></a>}

<a id="L50"></a>func (d *digest) Size() int { return Size }

<a id="L52"></a>func (d *digest) Write(p []byte) (nn int, err os.Error) {
    <a id="L53"></a>nn = len(p);
    <a id="L54"></a>d.len += uint64(nn);
    <a id="L55"></a>if d.nx &gt; 0 {
        <a id="L56"></a>n := len(p);
        <a id="L57"></a>if n &gt; _Chunk-d.nx {
            <a id="L58"></a>n = _Chunk - d.nx
        <a id="L59"></a>}
        <a id="L60"></a>for i := 0; i &lt; n; i++ {
            <a id="L61"></a>d.x[d.nx+i] = p[i]
        <a id="L62"></a>}
        <a id="L63"></a>d.nx += n;
        <a id="L64"></a>if d.nx == _Chunk {
            <a id="L65"></a>_Block(d, &amp;d.x);
            <a id="L66"></a>d.nx = 0;
        <a id="L67"></a>}
        <a id="L68"></a>p = p[n:len(p)];
    <a id="L69"></a>}
    <a id="L70"></a>n := _Block(d, p);
    <a id="L71"></a>p = p[n:len(p)];
    <a id="L72"></a>if len(p) &gt; 0 {
        <a id="L73"></a>for i := 0; i &lt; len(p); i++ {
            <a id="L74"></a>d.x[i] = p[i]
        <a id="L75"></a>}
        <a id="L76"></a>d.nx = len(p);
    <a id="L77"></a>}
    <a id="L78"></a>return;
<a id="L79"></a>}

<a id="L81"></a>func (d *digest) Sum() []byte {
    <a id="L82"></a><span class="comment">// Padding.  Add a 1 bit and 0 bits until 56 bytes mod 64.</span>
    <a id="L83"></a>len := d.len;
    <a id="L84"></a>var tmp [64]byte;
    <a id="L85"></a>tmp[0] = 0x80;
    <a id="L86"></a>if len%64 &lt; 56 {
        <a id="L87"></a>d.Write(tmp[0 : 56-len%64])
    <a id="L88"></a>} else {
        <a id="L89"></a>d.Write(tmp[0 : 64+56-len%64])
    <a id="L90"></a>}

    <a id="L92"></a><span class="comment">// Length in bits.</span>
    <a id="L93"></a>len &lt;&lt;= 3;
    <a id="L94"></a>for i := uint(0); i &lt; 8; i++ {
        <a id="L95"></a>tmp[i] = byte(len &gt;&gt; (56 - 8*i))
    <a id="L96"></a>}
    <a id="L97"></a>d.Write(tmp[0:8]);

    <a id="L99"></a>if d.nx != 0 {
        <a id="L100"></a>panicln(&#34;oops&#34;)
    <a id="L101"></a>}

    <a id="L103"></a>p := make([]byte, 20);
    <a id="L104"></a>j := 0;
    <a id="L105"></a>for i := 0; i &lt; 5; i++ {
        <a id="L106"></a>s := d.h[i];
        <a id="L107"></a>p[j] = byte(s &gt;&gt; 24);
        <a id="L108"></a>j++;
        <a id="L109"></a>p[j] = byte(s &gt;&gt; 16);
        <a id="L110"></a>j++;
        <a id="L111"></a>p[j] = byte(s &gt;&gt; 8);
        <a id="L112"></a>j++;
        <a id="L113"></a>p[j] = byte(s);
        <a id="L114"></a>j++;
    <a id="L115"></a>}
    <a id="L116"></a>return p;
<a id="L117"></a>}
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
