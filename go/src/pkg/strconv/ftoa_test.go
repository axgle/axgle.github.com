<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/ftoa_test.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strconv/ftoa_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package strconv_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;math&#34;;
    <a id="L9"></a>. &#34;strconv&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)

<a id="L13"></a>type ftoaTest struct {
    <a id="L14"></a>f    float64;
    <a id="L15"></a>fmt  byte;
    <a id="L16"></a>prec int;
    <a id="L17"></a>s    string;
<a id="L18"></a>}

<a id="L20"></a>func fdiv(a, b float64) float64 { return a / b } <span class="comment">// keep compiler in the dark</span>

<a id="L22"></a>const (
    <a id="L23"></a>below1e23 = 99999999999999974834176;
    <a id="L24"></a>above1e23 = 100000000000000008388608;
<a id="L25"></a>)

<a id="L27"></a>var ftoatests = []ftoaTest{
    <a id="L28"></a>ftoaTest{1, &#39;e&#39;, 5, &#34;1.00000e+00&#34;},
    <a id="L29"></a>ftoaTest{1, &#39;f&#39;, 5, &#34;1.00000&#34;},
    <a id="L30"></a>ftoaTest{1, &#39;g&#39;, 5, &#34;1&#34;},
    <a id="L31"></a>ftoaTest{1, &#39;g&#39;, -1, &#34;1&#34;},
    <a id="L32"></a>ftoaTest{20, &#39;g&#39;, -1, &#34;20&#34;},
    <a id="L33"></a>ftoaTest{1234567.8, &#39;g&#39;, -1, &#34;1.2345678e+06&#34;},
    <a id="L34"></a>ftoaTest{200000, &#39;g&#39;, -1, &#34;200000&#34;},
    <a id="L35"></a>ftoaTest{2000000, &#39;g&#39;, -1, &#34;2e+06&#34;},

    <a id="L37"></a>ftoaTest{0, &#39;e&#39;, 5, &#34;0.00000e+00&#34;},
    <a id="L38"></a>ftoaTest{0, &#39;f&#39;, 5, &#34;0.00000&#34;},
    <a id="L39"></a>ftoaTest{0, &#39;g&#39;, 5, &#34;0&#34;},
    <a id="L40"></a>ftoaTest{0, &#39;g&#39;, -1, &#34;0&#34;},

    <a id="L42"></a>ftoaTest{-1, &#39;e&#39;, 5, &#34;-1.00000e+00&#34;},
    <a id="L43"></a>ftoaTest{-1, &#39;f&#39;, 5, &#34;-1.00000&#34;},
    <a id="L44"></a>ftoaTest{-1, &#39;g&#39;, 5, &#34;-1&#34;},
    <a id="L45"></a>ftoaTest{-1, &#39;g&#39;, -1, &#34;-1&#34;},

    <a id="L47"></a>ftoaTest{12, &#39;e&#39;, 5, &#34;1.20000e+01&#34;},
    <a id="L48"></a>ftoaTest{12, &#39;f&#39;, 5, &#34;12.00000&#34;},
    <a id="L49"></a>ftoaTest{12, &#39;g&#39;, 5, &#34;12&#34;},
    <a id="L50"></a>ftoaTest{12, &#39;g&#39;, -1, &#34;12&#34;},

    <a id="L52"></a>ftoaTest{123456700, &#39;e&#39;, 5, &#34;1.23457e+08&#34;},
    <a id="L53"></a>ftoaTest{123456700, &#39;f&#39;, 5, &#34;123456700.00000&#34;},
    <a id="L54"></a>ftoaTest{123456700, &#39;g&#39;, 5, &#34;1.2346e+08&#34;},
    <a id="L55"></a>ftoaTest{123456700, &#39;g&#39;, -1, &#34;1.234567e+08&#34;},

    <a id="L57"></a>ftoaTest{1.2345e6, &#39;e&#39;, 5, &#34;1.23450e+06&#34;},
    <a id="L58"></a>ftoaTest{1.2345e6, &#39;f&#39;, 5, &#34;1234500.00000&#34;},
    <a id="L59"></a>ftoaTest{1.2345e6, &#39;g&#39;, 5, &#34;1.2345e+06&#34;},

    <a id="L61"></a>ftoaTest{1e23, &#39;e&#39;, 17, &#34;9.99999999999999916e+22&#34;},
    <a id="L62"></a>ftoaTest{1e23, &#39;f&#39;, 17, &#34;99999999999999991611392.00000000000000000&#34;},
    <a id="L63"></a>ftoaTest{1e23, &#39;g&#39;, 17, &#34;9.9999999999999992e+22&#34;},

    <a id="L65"></a>ftoaTest{1e23, &#39;e&#39;, -1, &#34;1e+23&#34;},
    <a id="L66"></a>ftoaTest{1e23, &#39;f&#39;, -1, &#34;100000000000000000000000&#34;},
    <a id="L67"></a>ftoaTest{1e23, &#39;g&#39;, -1, &#34;1e+23&#34;},

    <a id="L69"></a>ftoaTest{below1e23, &#39;e&#39;, 17, &#34;9.99999999999999748e+22&#34;},
    <a id="L70"></a>ftoaTest{below1e23, &#39;f&#39;, 17, &#34;99999999999999974834176.00000000000000000&#34;},
    <a id="L71"></a>ftoaTest{below1e23, &#39;g&#39;, 17, &#34;9.9999999999999975e+22&#34;},

    <a id="L73"></a>ftoaTest{below1e23, &#39;e&#39;, -1, &#34;9.999999999999997e+22&#34;},
    <a id="L74"></a>ftoaTest{below1e23, &#39;f&#39;, -1, &#34;99999999999999970000000&#34;},
    <a id="L75"></a>ftoaTest{below1e23, &#39;g&#39;, -1, &#34;9.999999999999997e+22&#34;},

    <a id="L77"></a>ftoaTest{above1e23, &#39;e&#39;, 17, &#34;1.00000000000000008e+23&#34;},
    <a id="L78"></a>ftoaTest{above1e23, &#39;f&#39;, 17, &#34;100000000000000008388608.00000000000000000&#34;},
    <a id="L79"></a>ftoaTest{above1e23, &#39;g&#39;, 17, &#34;1.0000000000000001e+23&#34;},

    <a id="L81"></a>ftoaTest{above1e23, &#39;e&#39;, -1, &#34;1.0000000000000001e+23&#34;},
    <a id="L82"></a>ftoaTest{above1e23, &#39;f&#39;, -1, &#34;100000000000000010000000&#34;},
    <a id="L83"></a>ftoaTest{above1e23, &#39;g&#39;, -1, &#34;1.0000000000000001e+23&#34;},

    <a id="L85"></a>ftoaTest{fdiv(5e-304, 1e20), &#39;g&#39;, -1, &#34;5e-324&#34;},
    <a id="L86"></a>ftoaTest{fdiv(-5e-304, 1e20), &#39;g&#39;, -1, &#34;-5e-324&#34;},

    <a id="L88"></a>ftoaTest{32, &#39;g&#39;, -1, &#34;32&#34;},
    <a id="L89"></a>ftoaTest{32, &#39;g&#39;, 0, &#34;3e+01&#34;},

    <a id="L91"></a>ftoaTest{100, &#39;x&#39;, -1, &#34;%x&#34;},

    <a id="L93"></a>ftoaTest{math.NaN(), &#39;g&#39;, -1, &#34;NaN&#34;},
    <a id="L94"></a>ftoaTest{-math.NaN(), &#39;g&#39;, -1, &#34;NaN&#34;},
    <a id="L95"></a>ftoaTest{math.Inf(0), &#39;g&#39;, -1, &#34;+Inf&#34;},
    <a id="L96"></a>ftoaTest{math.Inf(-1), &#39;g&#39;, -1, &#34;-Inf&#34;},
    <a id="L97"></a>ftoaTest{-math.Inf(0), &#39;g&#39;, -1, &#34;-Inf&#34;},

    <a id="L99"></a>ftoaTest{-1, &#39;b&#39;, -1, &#34;-4503599627370496p-52&#34;},
<a id="L100"></a>}

<a id="L102"></a>func TestFtoa(t *testing.T) {
    <a id="L103"></a>if FloatSize != 32 {
        <a id="L104"></a>panic(&#34;floatsize: &#34;, FloatSize)
    <a id="L105"></a>}
    <a id="L106"></a>for i := 0; i &lt; len(ftoatests); i++ {
        <a id="L107"></a>test := &amp;ftoatests[i];
        <a id="L108"></a>s := Ftoa64(test.f, test.fmt, test.prec);
        <a id="L109"></a>if s != test.s {
            <a id="L110"></a>t.Error(&#34;test&#34;, test.f, string(test.fmt), test.prec, &#34;want&#34;, test.s, &#34;got&#34;, s)
        <a id="L111"></a>}
        <a id="L112"></a>if float64(float32(test.f)) == test.f &amp;&amp; test.fmt != &#39;b&#39; {
            <a id="L113"></a>s := Ftoa32(float32(test.f), test.fmt, test.prec);
            <a id="L114"></a>if s != test.s {
                <a id="L115"></a>t.Error(&#34;test32&#34;, test.f, string(test.fmt), test.prec, &#34;want&#34;, test.s, &#34;got&#34;, s)
            <a id="L116"></a>}
        <a id="L117"></a>}
    <a id="L118"></a>}
<a id="L119"></a>}
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
