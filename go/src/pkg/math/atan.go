<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/math/atan.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/math/atan.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package math


<a id="L8"></a><span class="comment">/*</span>
<a id="L9"></a><span class="comment"> *	floating-point arctangent</span>
<a id="L10"></a><span class="comment"> *</span>
<a id="L11"></a><span class="comment"> *	atan returns the value of the arctangent of its</span>
<a id="L12"></a><span class="comment"> *	argument in the range [-pi/2,pi/2].</span>
<a id="L13"></a><span class="comment"> *	there are no error returns.</span>
<a id="L14"></a><span class="comment"> *	coefficients are #5077 from Hart &amp; Cheney. (19.56D)</span>
<a id="L15"></a><span class="comment"> */</span>

<a id="L17"></a><span class="comment">/*</span>
<a id="L18"></a><span class="comment"> *	xatan evaluates a series valid in the</span>
<a id="L19"></a><span class="comment"> *	range [-0.414...,+0.414...]. (tan(pi/8))</span>
<a id="L20"></a><span class="comment"> */</span>
<a id="L21"></a>func xatan(arg float64) float64 {
    <a id="L22"></a>const (
        <a id="L23"></a>P4  = .161536412982230228262e2;
        <a id="L24"></a>P3  = .26842548195503973794141e3;
        <a id="L25"></a>P2  = .11530293515404850115428136e4;
        <a id="L26"></a>P1  = .178040631643319697105464587e4;
        <a id="L27"></a>P0  = .89678597403663861959987488e3;
        <a id="L28"></a>Q4  = .5895697050844462222791e2;
        <a id="L29"></a>Q3  = .536265374031215315104235e3;
        <a id="L30"></a>Q2  = .16667838148816337184521798e4;
        <a id="L31"></a>Q1  = .207933497444540981287275926e4;
        <a id="L32"></a>Q0  = .89678597403663861962481162e3;
    <a id="L33"></a>)
    <a id="L34"></a>sq := arg * arg;
    <a id="L35"></a>value := ((((P4*sq+P3)*sq+P2)*sq+P1)*sq + P0);
    <a id="L36"></a>value = value / (((((sq+Q4)*sq+Q3)*sq+Q2)*sq+Q1)*sq + Q0);
    <a id="L37"></a>return value * arg;
<a id="L38"></a>}

<a id="L40"></a><span class="comment">/*</span>
<a id="L41"></a><span class="comment"> *	satan reduces its argument (known to be positive)</span>
<a id="L42"></a><span class="comment"> *	to the range [0,0.414...] and calls xatan.</span>
<a id="L43"></a><span class="comment"> */</span>
<a id="L44"></a>func satan(arg float64) float64 {
    <a id="L45"></a>if arg &lt; Sqrt2-1 {
        <a id="L46"></a>return xatan(arg)
    <a id="L47"></a>}
    <a id="L48"></a>if arg &gt; Sqrt2+1 {
        <a id="L49"></a>return Pi/2 - xatan(1/arg)
    <a id="L50"></a>}
    <a id="L51"></a>return Pi/4 + xatan((arg-1)/(arg+1));
<a id="L52"></a>}

<a id="L54"></a><span class="comment">/*</span>
<a id="L55"></a><span class="comment"> *	atan makes its argument positive and</span>
<a id="L56"></a><span class="comment"> *	calls the inner routine satan.</span>
<a id="L57"></a><span class="comment"> */</span>

<a id="L59"></a><span class="comment">// Atan returns the arc tangent of x.</span>
<a id="L60"></a>func Atan(x float64) float64 {
    <a id="L61"></a>if x &gt; 0 {
        <a id="L62"></a>return satan(x)
    <a id="L63"></a>}
    <a id="L64"></a>return -satan(-x);
<a id="L65"></a>}
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
