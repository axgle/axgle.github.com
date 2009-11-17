<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/ip_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/ip_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package net

<a id="L7"></a>import (
    <a id="L8"></a>&#34;testing&#34;;
<a id="L9"></a>)

<a id="L11"></a>func isEqual(a, b IP) bool {
    <a id="L12"></a>if a == nil &amp;&amp; b == nil {
        <a id="L13"></a>return true
    <a id="L14"></a>}
    <a id="L15"></a>if a == nil || b == nil || len(a) != len(b) {
        <a id="L16"></a>return false
    <a id="L17"></a>}
    <a id="L18"></a>for i := 0; i &lt; len(a); i++ {
        <a id="L19"></a>if a[i] != b[i] {
            <a id="L20"></a>return false
        <a id="L21"></a>}
    <a id="L22"></a>}
    <a id="L23"></a>return true;
<a id="L24"></a>}

<a id="L26"></a>type parseIPTest struct {
    <a id="L27"></a>in  string;
    <a id="L28"></a>out IP;
<a id="L29"></a>}

<a id="L31"></a>var parseiptests = []parseIPTest{
    <a id="L32"></a>parseIPTest{&#34;127.0.1.2&#34;, IPv4(127, 0, 1, 2)},
    <a id="L33"></a>parseIPTest{&#34;127.0.0.1&#34;, IPv4(127, 0, 0, 1)},
    <a id="L34"></a>parseIPTest{&#34;127.0.0.256&#34;, nil},
    <a id="L35"></a>parseIPTest{&#34;abc&#34;, nil},
    <a id="L36"></a>parseIPTest{&#34;::ffff:127.0.0.1&#34;, IPv4(127, 0, 0, 1)},
    <a id="L37"></a>parseIPTest{&#34;2001:4860:0:2001::68&#34;,
        <a id="L38"></a>IP{0x20, 0x01, 0x48, 0x60, 0, 0, 0x20, 0x01,
            <a id="L39"></a>0, 0, 0, 0, 0, 0, 0x00, 0x68,
        <a id="L40"></a>},
    <a id="L41"></a>},
    <a id="L42"></a>parseIPTest{&#34;::ffff:4a7d:1363&#34;, IPv4(74, 125, 19, 99)},
<a id="L43"></a>}

<a id="L45"></a>func TestParseIP(t *testing.T) {
    <a id="L46"></a>for i := 0; i &lt; len(parseiptests); i++ {
        <a id="L47"></a>tt := parseiptests[i];
        <a id="L48"></a>if out := ParseIP(tt.in); !isEqual(out, tt.out) {
            <a id="L49"></a>t.Errorf(&#34;ParseIP(%#q) = %v, want %v&#34;, tt.in, out, tt.out)
        <a id="L50"></a>}
    <a id="L51"></a>}
<a id="L52"></a>}
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
