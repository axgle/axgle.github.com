<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/block/cipher.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/block/cipher.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The block package implements standard block cipher modes</span>
<a id="L6"></a><span class="comment">// that can be wrapped around low-level block cipher implementations.</span>
<a id="L7"></a><span class="comment">// See http://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html</span>
<a id="L8"></a><span class="comment">// and NIST Special Publication 800-38A.</span>
<a id="L9"></a>package block

<a id="L11"></a><span class="comment">// A Cipher represents an implementation of block cipher</span>
<a id="L12"></a><span class="comment">// using a given key.  It provides the capability to encrypt</span>
<a id="L13"></a><span class="comment">// or decrypt individual blocks.  The mode implementations</span>
<a id="L14"></a><span class="comment">// extend that capability to streams of blocks.</span>
<a id="L15"></a>type Cipher interface {
    <a id="L16"></a><span class="comment">// BlockSize returns the cipher&#39;s block size.</span>
    <a id="L17"></a>BlockSize() int;

    <a id="L19"></a><span class="comment">// Encrypt encrypts the first block in src into dst.</span>
    <a id="L20"></a><span class="comment">// Src and dst may point at the same memory.</span>
    <a id="L21"></a>Encrypt(src, dst []byte);

    <a id="L23"></a><span class="comment">// Decrypt decrypts the first block in src into dst.</span>
    <a id="L24"></a><span class="comment">// Src and dst may point at the same memory.</span>
    <a id="L25"></a>Decrypt(src, dst []byte);
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// Utility routines</span>

<a id="L30"></a>func shift1(src, dst []byte) byte {
    <a id="L31"></a>var b byte;
    <a id="L32"></a>for i := len(src) - 1; i &gt;= 0; i-- {
        <a id="L33"></a>bb := src[i] &gt;&gt; 7;
        <a id="L34"></a>dst[i] = src[i]&lt;&lt;1 | b;
        <a id="L35"></a>b = bb;
    <a id="L36"></a>}
    <a id="L37"></a>return b;
<a id="L38"></a>}

<a id="L40"></a>func same(p, q []byte) bool {
    <a id="L41"></a>if len(p) != len(q) {
        <a id="L42"></a>return false
    <a id="L43"></a>}
    <a id="L44"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L45"></a>if p[i] != q[i] {
            <a id="L46"></a>return false
        <a id="L47"></a>}
    <a id="L48"></a>}
    <a id="L49"></a>return true;
<a id="L50"></a>}

<a id="L52"></a>func copy(p []byte) []byte {
    <a id="L53"></a>q := make([]byte, len(p));
    <a id="L54"></a>for i, b := range p {
        <a id="L55"></a>q[i] = b
    <a id="L56"></a>}
    <a id="L57"></a>return q;
<a id="L58"></a>}
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
