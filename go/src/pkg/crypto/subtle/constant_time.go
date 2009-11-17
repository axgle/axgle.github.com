<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/subtle/constant_time.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/crypto/subtle/constant_time.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements functions that are often useful in cryptographic</span>
<a id="L6"></a><span class="comment">// code but require careful thought to use correctly.</span>
<a id="L7"></a>package subtle

<a id="L9"></a><span class="comment">// ConstantTimeCompare returns 1 iff the two equal length slices, x</span>
<a id="L10"></a><span class="comment">// and y, have equal contents. The time taken is a function of the length of</span>
<a id="L11"></a><span class="comment">// the slices and is independent of the contents.</span>
<a id="L12"></a>func ConstantTimeCompare(x, y []byte) int {
    <a id="L13"></a>var v byte;

    <a id="L15"></a>for i := 0; i &lt; len(x); i++ {
        <a id="L16"></a>v |= x[i] ^ y[i]
    <a id="L17"></a>}

    <a id="L19"></a>return ConstantTimeByteEq(v, 0);
<a id="L20"></a>}

<a id="L22"></a><span class="comment">// ConstantTimeSelect returns x if v is 1 and y if v is 0.</span>
<a id="L23"></a><span class="comment">// Its behavior is undefined if v takes any other value.</span>
<a id="L24"></a>func ConstantTimeSelect(v, x, y int) int { return ^(v-1)&amp;x | (v-1)&amp;y }

<a id="L26"></a><span class="comment">// ConstantTimeByteEq returns 1 if x == x and 0 otherwise.</span>
<a id="L27"></a>func ConstantTimeByteEq(x, y uint8) int {
    <a id="L28"></a>z := ^(x ^ y);
    <a id="L29"></a>z &amp;= z &gt;&gt; 4;
    <a id="L30"></a>z &amp;= z &gt;&gt; 2;
    <a id="L31"></a>z &amp;= z &gt;&gt; 1;

    <a id="L33"></a>return int(z);
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// ConstantTimeEq returns 1 if x == y and 0 otherwise.</span>
<a id="L37"></a>func ConstantTimeEq(x, y int32) int {
    <a id="L38"></a>z := ^(x ^ y);
    <a id="L39"></a>z &amp;= z &gt;&gt; 16;
    <a id="L40"></a>z &amp;= z &gt;&gt; 8;
    <a id="L41"></a>z &amp;= z &gt;&gt; 4;
    <a id="L42"></a>z &amp;= z &gt;&gt; 2;
    <a id="L43"></a>z &amp;= z &gt;&gt; 1;

    <a id="L45"></a>return int(z &amp; 1);
<a id="L46"></a>}

<a id="L48"></a><span class="comment">// ConstantTimeCopy copies the contents of y into x iff v == 1. If v == 0, x is left unchanged.</span>
<a id="L49"></a><span class="comment">// Its behavior is undefined if v takes any other value.</span>
<a id="L50"></a>func ConstantTimeCopy(v int, x, y []byte) {
    <a id="L51"></a>xmask := byte(v - 1);
    <a id="L52"></a>ymask := byte(^(v - 1));
    <a id="L53"></a>for i := 0; i &lt; len(x); i++ {
        <a id="L54"></a>x[i] = x[i]&amp;xmask | y[i]&amp;ymask
    <a id="L55"></a>}
    <a id="L56"></a>return;
<a id="L57"></a>}
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
