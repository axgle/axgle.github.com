<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/4s/4s.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/4s/4s.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This is a simple demo of Go running under Native Client.</span>
<a id="L6"></a><span class="comment">// It is a tetris clone built on top of the exp/nacl/av and exp/draw</span>
<a id="L7"></a><span class="comment">// packages.</span>
<a id="L8"></a>package main

<a id="L10"></a>import (
    <a id="L11"></a>&#34;exp/nacl/av&#34;;
    <a id="L12"></a>&#34;exp/nacl/srpc&#34;;
    <a id="L13"></a>&#34;log&#34;;
    <a id="L14"></a>&#34;runtime&#34;;
    <a id="L15"></a>&#34;os&#34;;
<a id="L16"></a>)

<a id="L18"></a>var sndc chan []uint16

<a id="L20"></a>func main() {
    <a id="L21"></a><span class="comment">// Native Client requires that some calls are issued</span>
    <a id="L22"></a><span class="comment">// consistently by the same OS thread.</span>
    <a id="L23"></a>runtime.LockOSThread();

    <a id="L25"></a>if srpc.Enabled() {
        <a id="L26"></a>go srpc.ServeRuntime()
    <a id="L27"></a>}

    <a id="L29"></a>args := os.Args;
    <a id="L30"></a>p := pieces4;
    <a id="L31"></a>if len(args) &gt; 1 &amp;&amp; args[1] == &#34;-5&#34; {
        <a id="L32"></a>p = pieces5
    <a id="L33"></a>}
    <a id="L34"></a>dx, dy := 500, 500;
    <a id="L35"></a>w, err := av.Init(av.SubsystemVideo|av.SubsystemAudio, dx, dy);
    <a id="L36"></a>if err != nil {
        <a id="L37"></a>log.Exit(err)
    <a id="L38"></a>}

    <a id="L40"></a>sndc = make(chan []uint16, 10);
    <a id="L41"></a>go audioServer();
    <a id="L42"></a>Play(p, w);
<a id="L43"></a>}

<a id="L45"></a>func audioServer() {
    <a id="L46"></a><span class="comment">// Native Client requires that all audio calls</span>
    <a id="L47"></a><span class="comment">// original from a single OS thread.</span>
    <a id="L48"></a>runtime.LockOSThread();

    <a id="L50"></a>n, err := av.AudioStream(nil);
    <a id="L51"></a>if err != nil {
        <a id="L52"></a>log.Exit(err)
    <a id="L53"></a>}
    <a id="L54"></a>for {
        <a id="L55"></a>b := &lt;-sndc;
        <a id="L56"></a>for len(b)*2 &gt;= n {
            <a id="L57"></a>var a []uint16;
            <a id="L58"></a>a, b = b[0:n/2], b[n/2:len(b)];
            <a id="L59"></a>n, err = av.AudioStream(a);
            <a id="L60"></a>if err != nil {
                <a id="L61"></a>log.Exit(err)
            <a id="L62"></a>}
            <a id="L63"></a>println(n, len(b)*2);
        <a id="L64"></a>}
        <a id="L65"></a>a := make([]uint16, n/2);
        <a id="L66"></a>for i := range b {
            <a id="L67"></a>a[i] = b[i]
        <a id="L68"></a>}
        <a id="L69"></a>n, err = av.AudioStream(a);
    <a id="L70"></a>}
<a id="L71"></a>}

<a id="L73"></a>func PlaySound(b []uint16) { sndc &lt;- b }

<a id="L75"></a>var whoosh = []uint16{
<a id="L76"></a><span class="comment">// Insert your favorite sound samples here.</span>
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
