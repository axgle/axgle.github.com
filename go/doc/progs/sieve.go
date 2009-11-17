<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /doc/progs/sieve.go</title>

  <link rel="stylesheet" type="text/css" href="../style.css">
  <script type="text/javascript" src="../godocs.js"></script>

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
        <a href="../../index.html"><img src="../logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../go_tutorial.html">Tutorial</a></li>
    <li><a href="../effective_go.html">Effective Go</a></li>
    <li><a href="../go_faq.html">FAQ</a></li>
    <li><a href="../go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../go_spec.html">Language Specification</a></li>
    <li><a href="../go_mem.html">Memory Model</a></li>
    <li><a href="../go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../install.html">Install Go</a></li>
    <li><a href="../contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../src/index.html">Source files</a></li>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /doc/progs/sieve.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package main

<a id="L7"></a>import &#34;fmt&#34;

<a id="L9"></a><span class="comment">// Send the sequence 2, 3, 4, ... to channel &#39;ch&#39;.</span>
<a id="L10"></a>func generate(ch chan int) {
    <a id="L11"></a>for i := 2; ; i++ {
        <a id="L12"></a>ch &lt;- i <span class="comment">// Send &#39;i&#39; to channel &#39;ch&#39;.</span>
    <a id="L13"></a>}
<a id="L14"></a>}

<a id="L16"></a><span class="comment">// Copy the values from channel &#39;in&#39; to channel &#39;out&#39;,</span>
<a id="L17"></a><span class="comment">// removing those divisible by &#39;prime&#39;.</span>
<a id="L18"></a>func filter(in, out chan int, prime int) {
    <a id="L19"></a>for {
        <a id="L20"></a>i := &lt;-in; <span class="comment">// Receive value of new variable &#39;i&#39; from &#39;in&#39;.</span>
        <a id="L21"></a>if i%prime != 0 {
            <a id="L22"></a>out &lt;- i <span class="comment">// Send &#39;i&#39; to channel &#39;out&#39;.</span>
        <a id="L23"></a>}
    <a id="L24"></a>}
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// The prime sieve: Daisy-chain filter processes together.</span>
<a id="L28"></a>func main() {
    <a id="L29"></a>ch := make(chan int); <span class="comment">// Create a new channel.</span>
    <a id="L30"></a>go generate(ch);      <span class="comment">// Start generate() as a goroutine.</span>
    <a id="L31"></a>for {
        <a id="L32"></a>prime := &lt;-ch;
        <a id="L33"></a>fmt.Println(prime);
        <a id="L34"></a>ch1 := make(chan int);
        <a id="L35"></a>go filter(ch, ch1, prime);
        <a id="L36"></a>ch = ch1;
    <a id="L37"></a>}
<a id="L38"></a>}
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
