<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/math/bits.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/math/bits.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package math

<a id="L7"></a>const (
    <a id="L8"></a>uvnan    = 0x7FF0000000000001;
    <a id="L9"></a>uvinf    = 0x7FF0000000000000;
    <a id="L10"></a>uvneginf = 0xFFF0000000000000;
    <a id="L11"></a>mask     = 0x7FF;
    <a id="L12"></a>shift    = 64 - 11 - 1;
    <a id="L13"></a>bias     = 1022;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// Inf returns positive infinity if sign &gt;= 0, negative infinity if sign &lt; 0.</span>
<a id="L17"></a>func Inf(sign int) float64 {
    <a id="L18"></a>var v uint64;
    <a id="L19"></a>if sign &gt;= 0 {
        <a id="L20"></a>v = uvinf
    <a id="L21"></a>} else {
        <a id="L22"></a>v = uvneginf
    <a id="L23"></a>}
    <a id="L24"></a>return Float64frombits(v);
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// NaN returns an IEEE 754 ``not-a-number&#39;&#39; value.</span>
<a id="L28"></a>func NaN() float64 { return Float64frombits(uvnan) }

<a id="L30"></a><span class="comment">// IsNaN returns whether f is an IEEE 754 ``not-a-number&#39;&#39; value.</span>
<a id="L31"></a>func IsNaN(f float64) (is bool) {
    <a id="L32"></a>x := Float64bits(f);
    <a id="L33"></a>return uint32(x&gt;&gt;shift)&amp;mask == mask &amp;&amp; x != uvinf &amp;&amp; x != uvneginf;
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// IsInf returns whether f is an infinity, according to sign.</span>
<a id="L37"></a><span class="comment">// If sign &gt; 0, IsInf returns whether f is positive infinity.</span>
<a id="L38"></a><span class="comment">// If sign &lt; 0, IsInf returns whether f is negative infinity.</span>
<a id="L39"></a><span class="comment">// If sign == 0, IsInf returns whether f is either infinity.</span>
<a id="L40"></a>func IsInf(f float64, sign int) bool {
    <a id="L41"></a>x := Float64bits(f);
    <a id="L42"></a>return sign &gt;= 0 &amp;&amp; x == uvinf || sign &lt;= 0 &amp;&amp; x == uvneginf;
<a id="L43"></a>}

<a id="L45"></a><span class="comment">// Frexp breaks f into a normalized fraction</span>
<a id="L46"></a><span class="comment">// and an integral power of two.</span>
<a id="L47"></a><span class="comment">// It returns frac and exp satisfying f == frac × 2&lt;sup&gt;exp&lt;/sup&gt;,</span>
<a id="L48"></a><span class="comment">// with the absolute value of frac in the interval [½, 1).</span>
<a id="L49"></a>func Frexp(f float64) (frac float64, exp int) {
    <a id="L50"></a>if f == 0 {
        <a id="L51"></a>return
    <a id="L52"></a>}
    <a id="L53"></a>x := Float64bits(f);
    <a id="L54"></a>exp = int((x&gt;&gt;shift)&amp;mask) - bias;
    <a id="L55"></a>x &amp;^= mask &lt;&lt; shift;
    <a id="L56"></a>x |= bias &lt;&lt; shift;
    <a id="L57"></a>frac = Float64frombits(x);
    <a id="L58"></a>return;
<a id="L59"></a>}

<a id="L61"></a><span class="comment">// Ldexp is the inverse of Frexp.</span>
<a id="L62"></a><span class="comment">// It returns frac × 2&lt;sup&gt;exp&lt;/sup&gt;.</span>
<a id="L63"></a>func Ldexp(frac float64, exp int) float64 {
    <a id="L64"></a>x := Float64bits(frac);
    <a id="L65"></a>exp += int(x&gt;&gt;shift) &amp; mask;
    <a id="L66"></a>if exp &lt;= 0 {
        <a id="L67"></a>return 0 <span class="comment">// underflow</span>
    <a id="L68"></a>}
    <a id="L69"></a>if exp &gt;= mask { <span class="comment">// overflow</span>
        <a id="L70"></a>if frac &lt; 0 {
            <a id="L71"></a>return Inf(-1)
        <a id="L72"></a>}
        <a id="L73"></a>return Inf(1);
    <a id="L74"></a>}
    <a id="L75"></a>x &amp;^= mask &lt;&lt; shift;
    <a id="L76"></a>x |= uint64(exp) &lt;&lt; shift;
    <a id="L77"></a>return Float64frombits(x);
<a id="L78"></a>}

<a id="L80"></a><span class="comment">// Modf returns integer and fractional floating-point numbers</span>
<a id="L81"></a><span class="comment">// that sum to f.</span>
<a id="L82"></a><span class="comment">// Integer and frac have the same sign as f.</span>
<a id="L83"></a>func Modf(f float64) (int float64, frac float64) {
    <a id="L84"></a>if f &lt; 1 {
        <a id="L85"></a>if f &lt; 0 {
            <a id="L86"></a>int, frac = Modf(-f);
            <a id="L87"></a>return -int, -frac;
        <a id="L88"></a>}
        <a id="L89"></a>return 0, f;
    <a id="L90"></a>}

    <a id="L92"></a>x := Float64bits(f);
    <a id="L93"></a>e := uint(x&gt;&gt;shift)&amp;mask - bias;

    <a id="L95"></a><span class="comment">// Keep the top 11+e bits, the integer part; clear the rest.</span>
    <a id="L96"></a>if e &lt; 64-11 {
        <a id="L97"></a>x &amp;^= 1&lt;&lt;(64-11-e) - 1
    <a id="L98"></a>}
    <a id="L99"></a>int = Float64frombits(x);
    <a id="L100"></a>frac = f - int;
    <a id="L101"></a>return;
<a id="L102"></a>}
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
