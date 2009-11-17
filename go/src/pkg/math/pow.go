<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/math/pow.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/math/pow.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package math


<a id="L8"></a><span class="comment">// Pow returns x**y, the base-x exponential of y.</span>
<a id="L9"></a>func Pow(x, y float64) float64 {
    <a id="L10"></a><span class="comment">// TODO: x or y NaN, ±Inf, maybe ±0.</span>
    <a id="L11"></a>switch {
    <a id="L12"></a>case y == 0:
        <a id="L13"></a>return 1
    <a id="L14"></a>case y == 1:
        <a id="L15"></a>return x
    <a id="L16"></a>case x == 0 &amp;&amp; y &gt; 0:
        <a id="L17"></a>return 0
    <a id="L18"></a>case x == 0 &amp;&amp; y &lt; 0:
        <a id="L19"></a>return Inf(1)
    <a id="L20"></a>case y == 0.5:
        <a id="L21"></a>return Sqrt(x)
    <a id="L22"></a>case y == -0.5:
        <a id="L23"></a>return 1 / Sqrt(x)
    <a id="L24"></a>}

    <a id="L26"></a>absy := y;
    <a id="L27"></a>flip := false;
    <a id="L28"></a>if absy &lt; 0 {
        <a id="L29"></a>absy = -absy;
        <a id="L30"></a>flip = true;
    <a id="L31"></a>}
    <a id="L32"></a>yi, yf := Modf(absy);
    <a id="L33"></a>if yf != 0 &amp;&amp; x &lt; 0 {
        <a id="L34"></a>return NaN()
    <a id="L35"></a>}
    <a id="L36"></a>if yi &gt;= 1&lt;&lt;63 {
        <a id="L37"></a>return Exp(y * Log(x))
    <a id="L38"></a>}

    <a id="L40"></a><span class="comment">// ans = a1 * 2^ae (= 1 for now).</span>
    <a id="L41"></a>a1 := float64(1);
    <a id="L42"></a>ae := 0;

    <a id="L44"></a><span class="comment">// ans *= x^yf</span>
    <a id="L45"></a>if yf != 0 {
        <a id="L46"></a>if yf &gt; 0.5 {
            <a id="L47"></a>yf--;
            <a id="L48"></a>yi++;
        <a id="L49"></a>}
        <a id="L50"></a>a1 = Exp(yf * Log(x));
    <a id="L51"></a>}

    <a id="L53"></a><span class="comment">// ans *= x^yi</span>
    <a id="L54"></a><span class="comment">// by multiplying in successive squarings</span>
    <a id="L55"></a><span class="comment">// of x according to bits of yi.</span>
    <a id="L56"></a><span class="comment">// accumulate powers of two into exp.</span>
    <a id="L57"></a>x1, xe := Frexp(x);
    <a id="L58"></a>for i := int64(yi); i != 0; i &gt;&gt;= 1 {
        <a id="L59"></a>if i&amp;1 == 1 {
            <a id="L60"></a>a1 *= x1;
            <a id="L61"></a>ae += xe;
        <a id="L62"></a>}
        <a id="L63"></a>x1 *= x1;
        <a id="L64"></a>xe &lt;&lt;= 1;
        <a id="L65"></a>if x1 &lt; .5 {
            <a id="L66"></a>x1 += x1;
            <a id="L67"></a>xe--;
        <a id="L68"></a>}
    <a id="L69"></a>}

    <a id="L71"></a><span class="comment">// ans = a1*2^ae</span>
    <a id="L72"></a><span class="comment">// if flip { ans = 1 / ans }</span>
    <a id="L73"></a><span class="comment">// but in the opposite order</span>
    <a id="L74"></a>if flip {
        <a id="L75"></a>a1 = 1 / a1;
        <a id="L76"></a>ae = -ae;
    <a id="L77"></a>}
    <a id="L78"></a>return Ldexp(a1, ae);
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
