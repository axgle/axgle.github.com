<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/heap/heap.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/container/heap/heap.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package provides heap operations for any type that implements</span>
<a id="L6"></a><span class="comment">// heap.Interface.</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a>package heap

<a id="L10"></a>import &#34;sort&#34;

<a id="L12"></a><span class="comment">// Any type that implements heap.Interface may be used as a</span>
<a id="L13"></a><span class="comment">// heap with the following invariants (established after Init</span>
<a id="L14"></a><span class="comment">// has been called):</span>
<a id="L15"></a><span class="comment">//</span>
<a id="L16"></a><span class="comment">//	h.Less(i, j) for 0 &lt;= i &lt; h.Len() and j = 2*i+1 or 2*i+2 and j &lt; h.Len()</span>
<a id="L17"></a><span class="comment">//</span>
<a id="L18"></a>type Interface interface {
    <a id="L19"></a>sort.Interface;
    <a id="L20"></a>Push(x interface{});
    <a id="L21"></a>Pop() interface{};
<a id="L22"></a>}


<a id="L25"></a><span class="comment">// A heaper must be initialized before any of the heap operations</span>
<a id="L26"></a><span class="comment">// can be used. Init is idempotent with respect to the heap invariants</span>
<a id="L27"></a><span class="comment">// and may be called whenever the heap invariants may have been invalidated.</span>
<a id="L28"></a><span class="comment">// Its complexity is O(n*log(n)) where n = h.Len().</span>
<a id="L29"></a><span class="comment">//</span>
<a id="L30"></a>func Init(h Interface) { sort.Sort(h) }


<a id="L33"></a><span class="comment">// Push pushes the element x onto the heap. The complexity is</span>
<a id="L34"></a><span class="comment">// O(log(n)) where n = h.Len().</span>
<a id="L35"></a><span class="comment">//</span>
<a id="L36"></a>func Push(h Interface, x interface{}) {
    <a id="L37"></a>h.Push(x);
    <a id="L38"></a>up(h, h.Len()-1);
<a id="L39"></a>}


<a id="L42"></a><span class="comment">// Pop removes the minimum element (according to Less) from the heap</span>
<a id="L43"></a><span class="comment">// and returns it. The complexity is O(log(n)) where n = h.Len().</span>
<a id="L44"></a><span class="comment">//</span>
<a id="L45"></a>func Pop(h Interface) interface{} {
    <a id="L46"></a>n := h.Len() - 1;
    <a id="L47"></a>h.Swap(0, n);
    <a id="L48"></a>down(h, 0, n);
    <a id="L49"></a>return h.Pop();
<a id="L50"></a>}


<a id="L53"></a><span class="comment">// Remove removes the element at index i from the heap.</span>
<a id="L54"></a><span class="comment">// The complexity is O(log(n)) where n = h.Len().</span>
<a id="L55"></a><span class="comment">//</span>
<a id="L56"></a>func Remove(h Interface, i int) interface{} {
    <a id="L57"></a>n := h.Len() - 1;
    <a id="L58"></a>if n != i {
        <a id="L59"></a>h.Swap(n, i);
        <a id="L60"></a>down(h, i, n);
        <a id="L61"></a>up(h, i);
    <a id="L62"></a>}
    <a id="L63"></a>return h.Pop();
<a id="L64"></a>}


<a id="L67"></a>func up(h Interface, j int) {
    <a id="L68"></a>for {
        <a id="L69"></a>i := (j - 1) / 2;
        <a id="L70"></a>if i == j || h.Less(i, j) {
            <a id="L71"></a>break
        <a id="L72"></a>}
        <a id="L73"></a>h.Swap(i, j);
        <a id="L74"></a>j = i;
    <a id="L75"></a>}
<a id="L76"></a>}


<a id="L79"></a>func down(h Interface, i, n int) {
    <a id="L80"></a>for {
        <a id="L81"></a>j := 2*i + 1;
        <a id="L82"></a>if j &gt;= n {
            <a id="L83"></a>break
        <a id="L84"></a>}
        <a id="L85"></a>if j1 := j + 1; j1 &lt; n &amp;&amp; !h.Less(j, j1) {
            <a id="L86"></a>j = j1 <span class="comment">// = 2*i + 2</span>
        <a id="L87"></a>}
        <a id="L88"></a>if h.Less(i, j) {
            <a id="L89"></a>break
        <a id="L90"></a>}
        <a id="L91"></a>h.Swap(i, j);
        <a id="L92"></a>i = j;
    <a id="L93"></a>}
<a id="L94"></a>}
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
