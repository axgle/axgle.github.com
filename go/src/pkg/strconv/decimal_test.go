<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/decimal_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/strconv/decimal_test.go</h1>

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
    <a id="L8"></a>. &#34;strconv&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>type shiftTest struct {
    <a id="L13"></a>i     uint64;
    <a id="L14"></a>shift int;
    <a id="L15"></a>out   string;
<a id="L16"></a>}

<a id="L18"></a>var shifttests = []shiftTest{
    <a id="L19"></a>shiftTest{0, -100, &#34;0&#34;},
    <a id="L20"></a>shiftTest{0, 100, &#34;0&#34;},
    <a id="L21"></a>shiftTest{1, 100, &#34;1267650600228229401496703205376&#34;},
    <a id="L22"></a>shiftTest{1, -100,
        <a id="L23"></a>&#34;0.00000000000000000000000000000078886090522101180541&#34;
            <a id="L24"></a>&#34;17285652827862296732064351090230047702789306640625&#34;,
    <a id="L25"></a>},
    <a id="L26"></a>shiftTest{12345678, 8, &#34;3160493568&#34;},
    <a id="L27"></a>shiftTest{12345678, -8, &#34;48225.3046875&#34;},
    <a id="L28"></a>shiftTest{195312, 9, &#34;99999744&#34;},
    <a id="L29"></a>shiftTest{1953125, 9, &#34;1000000000&#34;},
<a id="L30"></a>}

<a id="L32"></a>func TestDecimalShift(t *testing.T) {
    <a id="L33"></a>for i := 0; i &lt; len(shifttests); i++ {
        <a id="L34"></a>test := &amp;shifttests[i];
        <a id="L35"></a>s := NewDecimal(test.i).Shift(test.shift).String();
        <a id="L36"></a>if s != test.out {
            <a id="L37"></a>t.Errorf(&#34;Decimal %v &lt;&lt; %v = %v, want %v\n&#34;,
                <a id="L38"></a>test.i, test.shift, s, test.out)
        <a id="L39"></a>}
    <a id="L40"></a>}
<a id="L41"></a>}

<a id="L43"></a>type roundTest struct {
    <a id="L44"></a>i               uint64;
    <a id="L45"></a>nd              int;
    <a id="L46"></a>down, round, up string;
    <a id="L47"></a>int             uint64;
<a id="L48"></a>}

<a id="L50"></a>var roundtests = []roundTest{
    <a id="L51"></a>roundTest{0, 4, &#34;0&#34;, &#34;0&#34;, &#34;0&#34;, 0},
    <a id="L52"></a>roundTest{12344999, 4, &#34;12340000&#34;, &#34;12340000&#34;, &#34;12350000&#34;, 12340000},
    <a id="L53"></a>roundTest{12345000, 4, &#34;12340000&#34;, &#34;12340000&#34;, &#34;12350000&#34;, 12340000},
    <a id="L54"></a>roundTest{12345001, 4, &#34;12340000&#34;, &#34;12350000&#34;, &#34;12350000&#34;, 12350000},
    <a id="L55"></a>roundTest{23454999, 4, &#34;23450000&#34;, &#34;23450000&#34;, &#34;23460000&#34;, 23450000},
    <a id="L56"></a>roundTest{23455000, 4, &#34;23450000&#34;, &#34;23460000&#34;, &#34;23460000&#34;, 23460000},
    <a id="L57"></a>roundTest{23455001, 4, &#34;23450000&#34;, &#34;23460000&#34;, &#34;23460000&#34;, 23460000},

    <a id="L59"></a>roundTest{99994999, 4, &#34;99990000&#34;, &#34;99990000&#34;, &#34;100000000&#34;, 99990000},
    <a id="L60"></a>roundTest{99995000, 4, &#34;99990000&#34;, &#34;100000000&#34;, &#34;100000000&#34;, 100000000},
    <a id="L61"></a>roundTest{99999999, 4, &#34;99990000&#34;, &#34;100000000&#34;, &#34;100000000&#34;, 100000000},

    <a id="L63"></a>roundTest{12994999, 4, &#34;12990000&#34;, &#34;12990000&#34;, &#34;13000000&#34;, 12990000},
    <a id="L64"></a>roundTest{12995000, 4, &#34;12990000&#34;, &#34;13000000&#34;, &#34;13000000&#34;, 13000000},
    <a id="L65"></a>roundTest{12999999, 4, &#34;12990000&#34;, &#34;13000000&#34;, &#34;13000000&#34;, 13000000},
<a id="L66"></a>}

<a id="L68"></a>func TestDecimalRound(t *testing.T) {
    <a id="L69"></a>for i := 0; i &lt; len(roundtests); i++ {
        <a id="L70"></a>test := &amp;roundtests[i];
        <a id="L71"></a>s := NewDecimal(test.i).RoundDown(test.nd).String();
        <a id="L72"></a>if s != test.down {
            <a id="L73"></a>t.Errorf(&#34;Decimal %v RoundDown %d = %v, want %v\n&#34;,
                <a id="L74"></a>test.i, test.nd, s, test.down)
        <a id="L75"></a>}
        <a id="L76"></a>s = NewDecimal(test.i).Round(test.nd).String();
        <a id="L77"></a>if s != test.round {
            <a id="L78"></a>t.Errorf(&#34;Decimal %v Round %d = %v, want %v\n&#34;,
                <a id="L79"></a>test.i, test.nd, s, test.down)
        <a id="L80"></a>}
        <a id="L81"></a>s = NewDecimal(test.i).RoundUp(test.nd).String();
        <a id="L82"></a>if s != test.up {
            <a id="L83"></a>t.Errorf(&#34;Decimal %v RoundUp %d = %v, want %v\n&#34;,
                <a id="L84"></a>test.i, test.nd, s, test.up)
        <a id="L85"></a>}
    <a id="L86"></a>}
<a id="L87"></a>}

<a id="L89"></a>type roundIntTest struct {
    <a id="L90"></a>i     uint64;
    <a id="L91"></a>shift int;
    <a id="L92"></a>int   uint64;
<a id="L93"></a>}

<a id="L95"></a>var roundinttests = []roundIntTest{
    <a id="L96"></a>roundIntTest{0, 100, 0},
    <a id="L97"></a>roundIntTest{512, -8, 2},
    <a id="L98"></a>roundIntTest{513, -8, 2},
    <a id="L99"></a>roundIntTest{640, -8, 2},
    <a id="L100"></a>roundIntTest{641, -8, 3},
    <a id="L101"></a>roundIntTest{384, -8, 2},
    <a id="L102"></a>roundIntTest{385, -8, 2},
    <a id="L103"></a>roundIntTest{383, -8, 1},
    <a id="L104"></a>roundIntTest{1, 100, 1&lt;&lt;64 - 1},
    <a id="L105"></a>roundIntTest{1000, 0, 1000},
<a id="L106"></a>}

<a id="L108"></a>func TestDecimalRoundedInteger(t *testing.T) {
    <a id="L109"></a>for i := 0; i &lt; len(roundinttests); i++ {
        <a id="L110"></a>test := roundinttests[i];
        <a id="L111"></a>int := NewDecimal(test.i).Shift(test.shift).RoundedInteger();
        <a id="L112"></a>if int != test.int {
            <a id="L113"></a>t.Errorf(&#34;Decimal %v &gt;&gt; %v RoundedInteger = %v, want %v\n&#34;,
                <a id="L114"></a>test.i, test.shift, int, test.int)
        <a id="L115"></a>}
    <a id="L116"></a>}
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
