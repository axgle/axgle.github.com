<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/net/net_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/net/net_test.go</h1>

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
    <a id="L8"></a>&#34;flag&#34;;
    <a id="L9"></a>&#34;regexp&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)

<a id="L13"></a>var runErrorTest = flag.Bool(&#34;run_error_test&#34;, false, &#34;let TestDialError check for dns errors&#34;)

<a id="L15"></a>type DialErrorTest struct {
    <a id="L16"></a>Net     string;
    <a id="L17"></a>Laddr   string;
    <a id="L18"></a>Raddr   string;
    <a id="L19"></a>Pattern string;
<a id="L20"></a>}

<a id="L22"></a>var dialErrorTests = []DialErrorTest{
    <a id="L23"></a>DialErrorTest{
        <a id="L24"></a>&#34;datakit&#34;, &#34;&#34;, &#34;mh/astro/r70&#34;,
        <a id="L25"></a>&#34;dial datakit mh/astro/r70: unknown network datakit&#34;,
    <a id="L26"></a>},
    <a id="L27"></a>DialErrorTest{
        <a id="L28"></a>&#34;tcp&#34;, &#34;&#34;, &#34;127.0.0.1:☺&#34;,
        <a id="L29"></a>&#34;dial tcp 127.0.0.1:☺: unknown port tcp/☺&#34;,
    <a id="L30"></a>},
    <a id="L31"></a>DialErrorTest{
        <a id="L32"></a>&#34;tcp&#34;, &#34;&#34;, &#34;no-such-name.google.com.:80&#34;,
        <a id="L33"></a>&#34;dial tcp no-such-name.google.com.:80: lookup no-such-name.google.com.( on .*)?: no (.*)&#34;,
    <a id="L34"></a>},
    <a id="L35"></a>DialErrorTest{
        <a id="L36"></a>&#34;tcp&#34;, &#34;&#34;, &#34;no-such-name.no-such-top-level-domain.:80&#34;,
        <a id="L37"></a>&#34;dial tcp no-such-name.no-such-top-level-domain.:80: lookup no-such-name.no-such-top-level-domain.( on .*)?: no (.*)&#34;,
    <a id="L38"></a>},
    <a id="L39"></a>DialErrorTest{
        <a id="L40"></a>&#34;tcp&#34;, &#34;&#34;, &#34;no-such-name:80&#34;,
        <a id="L41"></a>`dial tcp no-such-name:80: lookup no-such-name\.(.*\.)?( on .*)?: no (.*)`,
    <a id="L42"></a>},
    <a id="L43"></a>DialErrorTest{
        <a id="L44"></a>&#34;tcp&#34;, &#34;&#34;, &#34;mh/astro/r70:http&#34;,
        <a id="L45"></a>&#34;dial tcp mh/astro/r70:http: lookup mh/astro/r70: invalid domain name&#34;,
    <a id="L46"></a>},
    <a id="L47"></a>DialErrorTest{
        <a id="L48"></a>&#34;unix&#34;, &#34;&#34;, &#34;/etc/file-not-found&#34;,
        <a id="L49"></a>&#34;dial unix /etc/file-not-found: no such file or directory&#34;,
    <a id="L50"></a>},
    <a id="L51"></a>DialErrorTest{
        <a id="L52"></a>&#34;unix&#34;, &#34;&#34;, &#34;/etc/&#34;,
        <a id="L53"></a>&#34;dial unix /etc/: (permission denied|socket operation on non-socket|connection refused)&#34;,
    <a id="L54"></a>},
<a id="L55"></a>}

<a id="L57"></a>func TestDialError(t *testing.T) {
    <a id="L58"></a>if !*runErrorTest {
        <a id="L59"></a>t.Logf(&#34;test disabled; use --run_error_test to enable&#34;);
        <a id="L60"></a>return;
    <a id="L61"></a>}
    <a id="L62"></a>for i, tt := range dialErrorTests {
        <a id="L63"></a>c, e := Dial(tt.Net, tt.Laddr, tt.Raddr);
        <a id="L64"></a>if c != nil {
            <a id="L65"></a>c.Close()
        <a id="L66"></a>}
        <a id="L67"></a>if e == nil {
            <a id="L68"></a>t.Errorf(&#34;#%d: nil error, want match for %#q&#34;, i, tt.Pattern);
            <a id="L69"></a>continue;
        <a id="L70"></a>}
        <a id="L71"></a>s := e.String();
        <a id="L72"></a>match, _ := regexp.MatchString(tt.Pattern, s);
        <a id="L73"></a>if !match {
            <a id="L74"></a>t.Errorf(&#34;#%d: %q, want match for %#q&#34;, i, s, tt.Pattern)
        <a id="L75"></a>}
    <a id="L76"></a>}
<a id="L77"></a>}
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
