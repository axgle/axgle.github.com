<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/math/const.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/math/const.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The math package provides basic constants and mathematical functions.</span>
<a id="L6"></a>package math

<a id="L8"></a><span class="comment">// Mathematical constants.</span>
<a id="L9"></a><span class="comment">// Reference: http://www.research.att.com/~njas/sequences/Axxxxxx</span>
<a id="L10"></a>const (
    <a id="L11"></a>E   = 2.71828182845904523536028747135266249775724709369995957496696763; <span class="comment">// A001113</span>
    <a id="L12"></a>Pi  = 3.14159265358979323846264338327950288419716939937510582097494459; <span class="comment">// A000796</span>
    <a id="L13"></a>Phi = 1.61803398874989484820458683436563811772030917980576286213544862; <span class="comment">// A001622</span>

    <a id="L15"></a>Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974; <span class="comment">// A002193</span>
    <a id="L16"></a>SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931; <span class="comment">// A019774</span>
    <a id="L17"></a>SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779; <span class="comment">// A002161</span>
    <a id="L18"></a>SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038; <span class="comment">// A139339</span>

    <a id="L20"></a>Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009; <span class="comment">// A002162</span>
    <a id="L21"></a>Log2E  = 1 / Ln2;
    <a id="L22"></a>Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790; <span class="comment">// A002392</span>
    <a id="L23"></a>Log10E = 1 / Ln10;
<a id="L24"></a>)

<a id="L26"></a><span class="comment">// Floating-point limit values.</span>
<a id="L27"></a><span class="comment">// Max is the largest finite value representable by the type.</span>
<a id="L28"></a><span class="comment">// Min is the smallest nonzero value representable by the type.</span>
<a id="L29"></a>const (
    <a id="L30"></a>MaxFloat32 = 3.40282346638528859811704183484516925440e+38;  <span class="comment">/* 2^127 * (2^24 - 1) / 2^23 */</span>
    <a id="L31"></a>MinFloat32 = 1.401298464324817070923729583289916131280e-45; <span class="comment">/* 1 / 2^(127 - 1 + 23) */</span>

    <a id="L33"></a>MaxFloat64 = 1.797693134862315708145274237317043567981e+308; <span class="comment">/* 2^1023 * (2^53 - 1) / 2^52 */</span>
    <a id="L34"></a>MinFloat64 = 4.940656458412465441765687928682213723651e-324; <span class="comment">/* 1 / 2^(1023 - 1 + 52) */</span>
<a id="L35"></a>)

<a id="L37"></a><span class="comment">// Integer limit values.</span>
<a id="L38"></a>const (
    <a id="L39"></a>MaxInt8   = 1&lt;&lt;7 - 1;
    <a id="L40"></a>MinInt8   = -1 &lt;&lt; 7;
    <a id="L41"></a>MaxInt16  = 1&lt;&lt;15 - 1;
    <a id="L42"></a>MinInt16  = -1 &lt;&lt; 15;
    <a id="L43"></a>MaxInt32  = 1&lt;&lt;31 - 1;
    <a id="L44"></a>MinInt32  = -1 &lt;&lt; 31;
    <a id="L45"></a>MaxInt64  = 1&lt;&lt;63 - 1;
    <a id="L46"></a>MinInt64  = -1 &lt;&lt; 63;
    <a id="L47"></a>MaxUint8  = 1&lt;&lt;8 - 1;
    <a id="L48"></a>MaxUint16 = 1&lt;&lt;16 - 1;
    <a id="L49"></a>MaxUint32 = 1&lt;&lt;32 - 1;
    <a id="L50"></a>MaxUint64 = 1&lt;&lt;64 - 1;
<a id="L51"></a>)

<a id="L53"></a><span class="comment">// BUG(rsc): The manual should define the special cases for all of these functions.</span>
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
