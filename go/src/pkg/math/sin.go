<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/math/sin.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/math/sin.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package math


<a id="L8"></a>func sinus(x float64, quad int) float64 {
    <a id="L9"></a><span class="comment">// Coefficients are #3370 from Hart &amp; Cheney (18.80D).</span>
    <a id="L10"></a>const (
        <a id="L11"></a>P0  = .1357884097877375669092680e8;
        <a id="L12"></a>P1  = -.4942908100902844161158627e7;
        <a id="L13"></a>P2  = .4401030535375266501944918e6;
        <a id="L14"></a>P3  = -.1384727249982452873054457e5;
        <a id="L15"></a>P4  = .1459688406665768722226959e3;
        <a id="L16"></a>Q0  = .8644558652922534429915149e7;
        <a id="L17"></a>Q1  = .4081792252343299749395779e6;
        <a id="L18"></a>Q2  = .9463096101538208180571257e4;
        <a id="L19"></a>Q3  = .1326534908786136358911494e3;
    <a id="L20"></a>)
    <a id="L21"></a>if x &lt; 0 {
        <a id="L22"></a>x = -x;
        <a id="L23"></a>quad = quad + 2;
    <a id="L24"></a>}
    <a id="L25"></a>x = x * (2 / Pi); <span class="comment">/* underflow? */</span>
    <a id="L26"></a>var y float64;
    <a id="L27"></a>if x &gt; 32764 {
        <a id="L28"></a>var e float64;
        <a id="L29"></a>e, y = Modf(x);
        <a id="L30"></a>e = e + float64(quad);
        <a id="L31"></a>_, f := Modf(0.25 * e);
        <a id="L32"></a>quad = int(e - 4*f);
    <a id="L33"></a>} else {
        <a id="L34"></a>k := int32(x);
        <a id="L35"></a>y = x - float64(k);
        <a id="L36"></a>quad = (quad + int(k)) &amp; 3;
    <a id="L37"></a>}

    <a id="L39"></a>if quad&amp;1 != 0 {
        <a id="L40"></a>y = 1 - y
    <a id="L41"></a>}
    <a id="L42"></a>if quad &gt; 1 {
        <a id="L43"></a>y = -y
    <a id="L44"></a>}

    <a id="L46"></a>yy := y * y;
    <a id="L47"></a>temp1 := ((((P4*yy+P3)*yy+P2)*yy+P1)*yy + P0) * y;
    <a id="L48"></a>temp2 := ((((yy+Q3)*yy+Q2)*yy+Q1)*yy + Q0);
    <a id="L49"></a>return temp1 / temp2;
<a id="L50"></a>}

<a id="L52"></a><span class="comment">// Cos returns the cosine of x.</span>
<a id="L53"></a>func Cos(x float64) float64 {
    <a id="L54"></a>if x &lt; 0 {
        <a id="L55"></a>x = -x
    <a id="L56"></a>}
    <a id="L57"></a>return sinus(x, 1);
<a id="L58"></a>}

<a id="L60"></a><span class="comment">// Sin returns the sine of x.</span>
<a id="L61"></a>func Sin(x float64) float64 { return sinus(x, 0) }
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
