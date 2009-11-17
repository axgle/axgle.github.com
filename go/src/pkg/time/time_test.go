<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/time/time_test.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/time/time_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package time_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;testing&#34;;
    <a id="L10"></a>. &#34;time&#34;;
<a id="L11"></a>)

<a id="L13"></a>func init() {
    <a id="L14"></a><span class="comment">// Force US Pacific time for daylight-savings</span>
    <a id="L15"></a><span class="comment">// tests below (localtests).  Needs to be set</span>
    <a id="L16"></a><span class="comment">// before the first call into the time library.</span>
    <a id="L17"></a>os.Setenv(&#34;TZ&#34;, &#34;US/Pacific&#34;)
<a id="L18"></a>}

<a id="L20"></a>type TimeTest struct {
    <a id="L21"></a>seconds int64;
    <a id="L22"></a>golden  Time;
<a id="L23"></a>}

<a id="L25"></a>var utctests = []TimeTest{
    <a id="L26"></a>TimeTest{0, Time{1970, 1, 1, 0, 0, 0, Thursday, 0, &#34;UTC&#34;}},
    <a id="L27"></a>TimeTest{1221681866, Time{2008, 9, 17, 20, 4, 26, Wednesday, 0, &#34;UTC&#34;}},
    <a id="L28"></a>TimeTest{-1221681866, Time{1931, 4, 16, 3, 55, 34, Thursday, 0, &#34;UTC&#34;}},
    <a id="L29"></a>TimeTest{1e18, Time{31688740476, 10, 23, 1, 46, 40, Friday, 0, &#34;UTC&#34;}},
    <a id="L30"></a>TimeTest{-1e18, Time{-31688736537, 3, 10, 22, 13, 20, Tuesday, 0, &#34;UTC&#34;}},
    <a id="L31"></a>TimeTest{0x7fffffffffffffff, Time{292277026596, 12, 4, 15, 30, 7, Sunday, 0, &#34;UTC&#34;}},
    <a id="L32"></a>TimeTest{-0x8000000000000000, Time{-292277022657, 1, 27, 8, 29, 52, Sunday, 0, &#34;UTC&#34;}},
<a id="L33"></a>}

<a id="L35"></a>var localtests = []TimeTest{
    <a id="L36"></a>TimeTest{0, Time{1969, 12, 31, 16, 0, 0, Wednesday, -8 * 60 * 60, &#34;PST&#34;}},
    <a id="L37"></a>TimeTest{1221681866, Time{2008, 9, 17, 13, 4, 26, Wednesday, -7 * 60 * 60, &#34;PDT&#34;}},
<a id="L38"></a>}

<a id="L40"></a>func same(t, u *Time) bool {
    <a id="L41"></a>return t.Year == u.Year &amp;&amp;
        <a id="L42"></a>t.Month == u.Month &amp;&amp;
        <a id="L43"></a>t.Day == u.Day &amp;&amp;
        <a id="L44"></a>t.Hour == u.Hour &amp;&amp;
        <a id="L45"></a>t.Minute == u.Minute &amp;&amp;
        <a id="L46"></a>t.Second == u.Second &amp;&amp;
        <a id="L47"></a>t.Weekday == u.Weekday &amp;&amp;
        <a id="L48"></a>t.ZoneOffset == u.ZoneOffset &amp;&amp;
        <a id="L49"></a>t.Zone == u.Zone
<a id="L50"></a>}

<a id="L52"></a>func TestSecondsToUTC(t *testing.T) {
    <a id="L53"></a>for i := 0; i &lt; len(utctests); i++ {
        <a id="L54"></a>sec := utctests[i].seconds;
        <a id="L55"></a>golden := &amp;utctests[i].golden;
        <a id="L56"></a>tm := SecondsToUTC(sec);
        <a id="L57"></a>newsec := tm.Seconds();
        <a id="L58"></a>if newsec != sec {
            <a id="L59"></a>t.Errorf(&#34;SecondsToUTC(%d).Seconds() = %d&#34;, sec, newsec)
        <a id="L60"></a>}
        <a id="L61"></a>if !same(tm, golden) {
            <a id="L62"></a>t.Errorf(&#34;SecondsToUTC(%d):&#34;, sec);
            <a id="L63"></a>t.Errorf(&#34;  want=%+v&#34;, *golden);
            <a id="L64"></a>t.Errorf(&#34;  have=%+v&#34;, *tm);
        <a id="L65"></a>}
    <a id="L66"></a>}
<a id="L67"></a>}

<a id="L69"></a>func TestSecondsToLocalTime(t *testing.T) {
    <a id="L70"></a>for i := 0; i &lt; len(localtests); i++ {
        <a id="L71"></a>sec := localtests[i].seconds;
        <a id="L72"></a>golden := &amp;localtests[i].golden;
        <a id="L73"></a>tm := SecondsToLocalTime(sec);
        <a id="L74"></a>newsec := tm.Seconds();
        <a id="L75"></a>if newsec != sec {
            <a id="L76"></a>t.Errorf(&#34;SecondsToLocalTime(%d).Seconds() = %d&#34;, sec, newsec)
        <a id="L77"></a>}
        <a id="L78"></a>if !same(tm, golden) {
            <a id="L79"></a>t.Errorf(&#34;SecondsToLocalTime(%d):&#34;, sec);
            <a id="L80"></a>t.Errorf(&#34;  want=%+v&#34;, *golden);
            <a id="L81"></a>t.Errorf(&#34;  have=%+v&#34;, *tm);
        <a id="L82"></a>}
    <a id="L83"></a>}
<a id="L84"></a>}
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
