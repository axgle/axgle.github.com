<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/heap/heap_test.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/container/heap/heap_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package heap

<a id="L7"></a>import (
    <a id="L8"></a>&#34;testing&#34;;
    <a id="L9"></a>&#34;container/vector&#34;;
<a id="L10"></a>)


<a id="L13"></a>type myHeap struct {
    <a id="L14"></a>vector.IntVector;
<a id="L15"></a>}


<a id="L18"></a>func newHeap() *myHeap {
    <a id="L19"></a>var h myHeap;
    <a id="L20"></a>h.IntVector.Init(0);
    <a id="L21"></a>return &amp;h;
<a id="L22"></a>}


<a id="L25"></a>func (h *myHeap) verify(t *testing.T, i int) {
    <a id="L26"></a>n := h.Len();
    <a id="L27"></a>j1 := 2*i + 1;
    <a id="L28"></a>j2 := 2*i + 2;
    <a id="L29"></a>if j1 &lt; n {
        <a id="L30"></a>if h.Less(j1, i) {
            <a id="L31"></a>t.Errorf(&#34;heap invariant invalidated [%d] = %d &gt; [%d] = %d&#34;, i, h.At(i), j1, h.At(j1));
            <a id="L32"></a>return;
        <a id="L33"></a>}
        <a id="L34"></a>h.verify(t, j1);
    <a id="L35"></a>}
    <a id="L36"></a>if j2 &lt; n {
        <a id="L37"></a>if h.Less(j2, i) {
            <a id="L38"></a>t.Errorf(&#34;heap invariant invalidated [%d] = %d &gt; [%d] = %d&#34;, i, h.At(i), j1, h.At(j2));
            <a id="L39"></a>return;
        <a id="L40"></a>}
        <a id="L41"></a>h.verify(t, j2);
    <a id="L42"></a>}
<a id="L43"></a>}


<a id="L46"></a>func (h *myHeap) Push(x interface{}) { h.IntVector.Push(x.(int)) }


<a id="L49"></a>func (h *myHeap) Pop() interface{} { return h.IntVector.Pop() }


<a id="L52"></a>func TestInit(t *testing.T) {
    <a id="L53"></a>h := newHeap();
    <a id="L54"></a>for i := 20; i &gt; 0; i-- {
        <a id="L55"></a>h.Push(i)
    <a id="L56"></a>}
    <a id="L57"></a>Init(h);
    <a id="L58"></a>h.verify(t, 0);

    <a id="L60"></a>for i := 1; h.Len() &gt; 0; i++ {
        <a id="L61"></a>x := Pop(h).(int);
        <a id="L62"></a>h.verify(t, 0);
        <a id="L63"></a>if x != i {
            <a id="L64"></a>t.Errorf(&#34;%d.th pop got %d; want %d&#34;, i, x, i)
        <a id="L65"></a>}
    <a id="L66"></a>}
<a id="L67"></a>}


<a id="L70"></a>func Test(t *testing.T) {
    <a id="L71"></a>h := newHeap();
    <a id="L72"></a>h.verify(t, 0);

    <a id="L74"></a>for i := 20; i &gt; 10; i-- {
        <a id="L75"></a>h.Push(i)
    <a id="L76"></a>}
    <a id="L77"></a>Init(h);
    <a id="L78"></a>h.verify(t, 0);

    <a id="L80"></a>for i := 10; i &gt; 0; i-- {
        <a id="L81"></a>Push(h, i);
        <a id="L82"></a>h.verify(t, 0);
    <a id="L83"></a>}

    <a id="L85"></a>for i := 1; h.Len() &gt; 0; i++ {
        <a id="L86"></a>x := Pop(h).(int);
        <a id="L87"></a>if i &lt; 20 {
            <a id="L88"></a>Push(h, 20+i)
        <a id="L89"></a>}
        <a id="L90"></a>h.verify(t, 0);
        <a id="L91"></a>if x != i {
            <a id="L92"></a>t.Errorf(&#34;%d.th pop got %d; want %d&#34;, i, x, i)
        <a id="L93"></a>}
    <a id="L94"></a>}
<a id="L95"></a>}
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
