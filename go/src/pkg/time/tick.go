<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/time/tick.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/time/tick.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package time

<a id="L7"></a><span class="comment">// TODO(rsc): This implementation of Tick is a</span>
<a id="L8"></a><span class="comment">// simple placeholder.  Eventually, there will need to be</span>
<a id="L9"></a><span class="comment">// a single central time server no matter how many tickers</span>
<a id="L10"></a><span class="comment">// are active.</span>
<a id="L11"></a><span class="comment">//</span>
<a id="L12"></a><span class="comment">// Also, if timeouts become part of the select statement,</span>
<a id="L13"></a><span class="comment">// perhaps the Ticker is just:</span>
<a id="L14"></a><span class="comment">//</span>
<a id="L15"></a><span class="comment">//	func Ticker(ns int64, c chan int64) {</span>
<a id="L16"></a><span class="comment">//		for {</span>
<a id="L17"></a><span class="comment">//			select { timeout ns: }</span>
<a id="L18"></a><span class="comment">//			nsec, err := Nanoseconds();</span>
<a id="L19"></a><span class="comment">//			c &lt;- nsec;</span>
<a id="L20"></a><span class="comment">//		}</span>


<a id="L23"></a><span class="comment">// A Ticker holds a synchronous channel that delivers `ticks&#39; of a clock</span>
<a id="L24"></a><span class="comment">// at intervals.</span>
<a id="L25"></a>type Ticker struct {
    <a id="L26"></a>C        &lt;-chan int64; <span class="comment">// The channel on which the ticks are delivered.</span>
    <a id="L27"></a>ns       int64;
    <a id="L28"></a>shutdown bool;
<a id="L29"></a>}

<a id="L31"></a><span class="comment">// Stop turns off a ticker.  After Stop, no more ticks will be delivered.</span>
<a id="L32"></a>func (t *Ticker) Stop() { t.shutdown = true }

<a id="L34"></a>func (t *Ticker) ticker(c chan&lt;- int64) {
    <a id="L35"></a>now := Nanoseconds();
    <a id="L36"></a>when := now;
    <a id="L37"></a>for !t.shutdown {
        <a id="L38"></a>when += t.ns; <span class="comment">// next alarm</span>

        <a id="L40"></a><span class="comment">// if c &lt;- now took too long, skip ahead</span>
        <a id="L41"></a>if when &lt; now {
            <a id="L42"></a><span class="comment">// one big step</span>
            <a id="L43"></a>when += (now - when) / t.ns * t.ns
        <a id="L44"></a>}
        <a id="L45"></a>for when &lt;= now {
            <a id="L46"></a><span class="comment">// little steps until when &gt; now</span>
            <a id="L47"></a>when += t.ns
        <a id="L48"></a>}

        <a id="L50"></a>Sleep(when - now);
        <a id="L51"></a>now = Nanoseconds();
        <a id="L52"></a>if t.shutdown {
            <a id="L53"></a>return
        <a id="L54"></a>}
        <a id="L55"></a>c &lt;- now;
    <a id="L56"></a>}
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// Tick is a convenience wrapper for NewTicker providing access to the ticking</span>
<a id="L60"></a><span class="comment">// channel only.  Useful for clients that have no need to shut down the ticker.</span>
<a id="L61"></a>func Tick(ns int64) &lt;-chan int64 {
    <a id="L62"></a>if ns &lt;= 0 {
        <a id="L63"></a>return nil
    <a id="L64"></a>}
    <a id="L65"></a>return NewTicker(ns).C;
<a id="L66"></a>}

<a id="L68"></a><span class="comment">// Ticker returns a new Ticker containing a synchronous channel that will</span>
<a id="L69"></a><span class="comment">// send the time, in nanoseconds, every ns nanoseconds.  It adjusts the</span>
<a id="L70"></a><span class="comment">// intervals to make up for pauses in delivery of the ticks.</span>
<a id="L71"></a>func NewTicker(ns int64) *Ticker {
    <a id="L72"></a>if ns &lt;= 0 {
        <a id="L73"></a>return nil
    <a id="L74"></a>}
    <a id="L75"></a>c := make(chan int64);
    <a id="L76"></a>t := &amp;Ticker{c, ns, false};
    <a id="L77"></a>go t.ticker(c);
    <a id="L78"></a>return t;
<a id="L79"></a>}
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
