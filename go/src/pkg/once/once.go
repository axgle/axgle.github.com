<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/once/once.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/once/once.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package provides a single function, Do, to run a function</span>
<a id="L6"></a><span class="comment">// exactly once, usually used as part of initialization.</span>
<a id="L7"></a>package once

<a id="L9"></a>import &#34;sync&#34;

<a id="L11"></a>type job struct {
    <a id="L12"></a>done        bool;
    <a id="L13"></a>sync.Mutex; <span class="comment">// should probably be sync.Notification or some such</span>
<a id="L14"></a>}

<a id="L16"></a>var jobs = make(map[func()]*job)
<a id="L17"></a>var joblock sync.Mutex

<a id="L19"></a><span class="comment">// Do is the the only exported piece of the package.</span>
<a id="L20"></a><span class="comment">// For one-time initialization that is not done during init,</span>
<a id="L21"></a><span class="comment">// wrap the initialization in a niladic function f() and call</span>
<a id="L22"></a><span class="comment">//	Do(f)</span>
<a id="L23"></a><span class="comment">// If multiple processes call Do(f) simultaneously</span>
<a id="L24"></a><span class="comment">// with the same f argument, only one will call f, and the</span>
<a id="L25"></a><span class="comment">// others will block until f finishes running.</span>
<a id="L26"></a><span class="comment">//</span>
<a id="L27"></a><span class="comment">// Since a func() expression typically evaluates to a differerent</span>
<a id="L28"></a><span class="comment">// function value each time it is evaluated, it is incorrect to</span>
<a id="L29"></a><span class="comment">// pass such values to Do.  For example,</span>
<a id="L30"></a><span class="comment">//	func f(x int) {</span>
<a id="L31"></a><span class="comment">//		Do(func() { fmt.Println(x) })</span>
<a id="L32"></a><span class="comment">//	}</span>
<a id="L33"></a><span class="comment">// behaves the same as</span>
<a id="L34"></a><span class="comment">//	func f(x int) {</span>
<a id="L35"></a><span class="comment">//		fmt.Println(x)</span>
<a id="L36"></a><span class="comment">//	}</span>
<a id="L37"></a><span class="comment">// because the func() expression in the first creates a new</span>
<a id="L38"></a><span class="comment">// func each time f runs, and each of those funcs is run once.</span>
<a id="L39"></a>func Do(f func()) {
    <a id="L40"></a>joblock.Lock();
    <a id="L41"></a>j, present := jobs[f];
    <a id="L42"></a>if !present {
        <a id="L43"></a><span class="comment">// run it</span>
        <a id="L44"></a>j = new(job);
        <a id="L45"></a>j.Lock();
        <a id="L46"></a>jobs[f] = j;
        <a id="L47"></a>joblock.Unlock();
        <a id="L48"></a>f();
        <a id="L49"></a>j.done = true;
        <a id="L50"></a>j.Unlock();
    <a id="L51"></a>} else {
        <a id="L52"></a><span class="comment">// wait for it</span>
        <a id="L53"></a>joblock.Unlock();
        <a id="L54"></a>if j.done != true {
            <a id="L55"></a>j.Lock();
            <a id="L56"></a>j.Unlock();
        <a id="L57"></a>}
    <a id="L58"></a>}
<a id="L59"></a>}
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
