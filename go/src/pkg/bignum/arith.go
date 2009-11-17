<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bignum/arith.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/bignum/arith.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Fast versions of the routines in this file are in fast.arith.s.</span>
<a id="L6"></a><span class="comment">// Simply replace this file with arith.s (renamed from fast.arith.s)</span>
<a id="L7"></a><span class="comment">// and the bignum package will build and run on a platform that</span>
<a id="L8"></a><span class="comment">// supports the assembly routines.</span>

<a id="L10"></a>package bignum

<a id="L12"></a>import &#34;unsafe&#34;

<a id="L14"></a><span class="comment">// z1&lt;&lt;64 + z0 = x*y</span>
<a id="L15"></a>func Mul128(x, y uint64) (z1, z0 uint64) {
    <a id="L16"></a><span class="comment">// Split x and y into 2 halfwords each, multiply</span>
    <a id="L17"></a><span class="comment">// the halfwords separately while avoiding overflow,</span>
    <a id="L18"></a><span class="comment">// and return the product as 2 words.</span>

    <a id="L20"></a>const (
        <a id="L21"></a>W   = uint(unsafe.Sizeof(x)) * 8;
        <a id="L22"></a>W2  = W / 2;
        <a id="L23"></a>B2  = 1 &lt;&lt; W2;
        <a id="L24"></a>M2  = B2 - 1;
    <a id="L25"></a>)

    <a id="L27"></a>if x &lt; y {
        <a id="L28"></a>x, y = y, x
    <a id="L29"></a>}

    <a id="L31"></a>if x &lt; B2 {
        <a id="L32"></a><span class="comment">// y &lt; B2 because y &lt;= x</span>
        <a id="L33"></a><span class="comment">// sub-digits of x and y are (0, x) and (0, y)</span>
        <a id="L34"></a><span class="comment">// z = z[0] = x*y</span>
        <a id="L35"></a>z0 = x * y;
        <a id="L36"></a>return;
    <a id="L37"></a>}

    <a id="L39"></a>if y &lt; B2 {
        <a id="L40"></a><span class="comment">// sub-digits of x and y are (x1, x0) and (0, y)</span>
        <a id="L41"></a><span class="comment">// x = (x1*B2 + x0)</span>
        <a id="L42"></a><span class="comment">// y = (y1*B2 + y0)</span>
        <a id="L43"></a>x1, x0 := x&gt;&gt;W2, x&amp;M2;

        <a id="L45"></a><span class="comment">// x*y = t2*B2*B2 + t1*B2 + t0</span>
        <a id="L46"></a>t0 := x0 * y;
        <a id="L47"></a>t1 := x1 * y;

        <a id="L49"></a><span class="comment">// compute result digits but avoid overflow</span>
        <a id="L50"></a><span class="comment">// z = z[1]*B + z[0] = x*y</span>
        <a id="L51"></a>z0 = t1&lt;&lt;W2 + t0;
        <a id="L52"></a>z1 = (t1 + t0&gt;&gt;W2) &gt;&gt; W2;
        <a id="L53"></a>return;
    <a id="L54"></a>}

    <a id="L56"></a><span class="comment">// general case</span>
    <a id="L57"></a><span class="comment">// sub-digits of x and y are (x1, x0) and (y1, y0)</span>
    <a id="L58"></a><span class="comment">// x = (x1*B2 + x0)</span>
    <a id="L59"></a><span class="comment">// y = (y1*B2 + y0)</span>
    <a id="L60"></a>x1, x0 := x&gt;&gt;W2, x&amp;M2;
    <a id="L61"></a>y1, y0 := y&gt;&gt;W2, y&amp;M2;

    <a id="L63"></a><span class="comment">// x*y = t2*B2*B2 + t1*B2 + t0</span>
    <a id="L64"></a>t0 := x0 * y0;
    <a id="L65"></a>t1 := x1*y0 + x0*y1;
    <a id="L66"></a>t2 := x1 * y1;

    <a id="L68"></a><span class="comment">// compute result digits but avoid overflow</span>
    <a id="L69"></a><span class="comment">// z = z[1]*B + z[0] = x*y</span>
    <a id="L70"></a>z0 = t1&lt;&lt;W2 + t0;
    <a id="L71"></a>z1 = t2 + (t1+t0&gt;&gt;W2)&gt;&gt;W2;
    <a id="L72"></a>return;
<a id="L73"></a>}


<a id="L76"></a><span class="comment">// z1&lt;&lt;64 + z0 = x*y + c</span>
<a id="L77"></a>func MulAdd128(x, y, c uint64) (z1, z0 uint64) {
    <a id="L78"></a><span class="comment">// Split x and y into 2 halfwords each, multiply</span>
    <a id="L79"></a><span class="comment">// the halfwords separately while avoiding overflow,</span>
    <a id="L80"></a><span class="comment">// and return the product as 2 words.</span>

    <a id="L82"></a>const (
        <a id="L83"></a>W   = uint(unsafe.Sizeof(x)) * 8;
        <a id="L84"></a>W2  = W / 2;
        <a id="L85"></a>B2  = 1 &lt;&lt; W2;
        <a id="L86"></a>M2  = B2 - 1;
    <a id="L87"></a>)

    <a id="L89"></a><span class="comment">// TODO(gri) Should implement special cases for faster execution.</span>

    <a id="L91"></a><span class="comment">// general case</span>
    <a id="L92"></a><span class="comment">// sub-digits of x, y, and c are (x1, x0), (y1, y0), (c1, c0)</span>
    <a id="L93"></a><span class="comment">// x = (x1*B2 + x0)</span>
    <a id="L94"></a><span class="comment">// y = (y1*B2 + y0)</span>
    <a id="L95"></a>x1, x0 := x&gt;&gt;W2, x&amp;M2;
    <a id="L96"></a>y1, y0 := y&gt;&gt;W2, y&amp;M2;
    <a id="L97"></a>c1, c0 := c&gt;&gt;W2, c&amp;M2;

    <a id="L99"></a><span class="comment">// x*y + c = t2*B2*B2 + t1*B2 + t0</span>
    <a id="L100"></a>t0 := x0*y0 + c0;
    <a id="L101"></a>t1 := x1*y0 + x0*y1 + c1;
    <a id="L102"></a>t2 := x1 * y1;

    <a id="L104"></a><span class="comment">// compute result digits but avoid overflow</span>
    <a id="L105"></a><span class="comment">// z = z[1]*B + z[0] = x*y</span>
    <a id="L106"></a>z0 = t1&lt;&lt;W2 + t0;
    <a id="L107"></a>z1 = t2 + (t1+t0&gt;&gt;W2)&gt;&gt;W2;
    <a id="L108"></a>return;
<a id="L109"></a>}


<a id="L112"></a><span class="comment">// q = (x1&lt;&lt;64 + x0)/y + r</span>
<a id="L113"></a>func Div128(x1, x0, y uint64) (q, r uint64) {
    <a id="L114"></a>if x1 == 0 {
        <a id="L115"></a>q, r = x0/y, x0%y;
        <a id="L116"></a>return;
    <a id="L117"></a>}

    <a id="L119"></a><span class="comment">// TODO(gri) Implement general case.</span>
    <a id="L120"></a>panic(&#34;Div128 not implemented for x &gt; 1&lt;&lt;64-1&#34;);
<a id="L121"></a>}
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
