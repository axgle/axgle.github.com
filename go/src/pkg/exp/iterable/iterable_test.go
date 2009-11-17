<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/iterable/iterable_test.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/iterable/iterable_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package iterable

<a id="L7"></a>import (
    <a id="L8"></a>&#34;testing&#34;;
<a id="L9"></a>)

<a id="L11"></a>type IntArray []int

<a id="L13"></a>func (arr IntArray) Iter() &lt;-chan interface{} {
    <a id="L14"></a>ch := make(chan interface{});
    <a id="L15"></a>go func() {
        <a id="L16"></a>for _, x := range arr {
            <a id="L17"></a>ch &lt;- x
        <a id="L18"></a>}
        <a id="L19"></a>close(ch);
    <a id="L20"></a>}();
    <a id="L21"></a>return ch;
<a id="L22"></a>}

<a id="L24"></a>var oneToFive = IntArray{1, 2, 3, 4, 5}

<a id="L26"></a>func isNegative(n interface{}) bool     { return n.(int) &lt; 0 }
<a id="L27"></a>func isPositive(n interface{}) bool     { return n.(int) &gt; 0 }
<a id="L28"></a>func isAbove3(n interface{}) bool       { return n.(int) &gt; 3 }
<a id="L29"></a>func isEven(n interface{}) bool         { return n.(int)%2 == 0 }
<a id="L30"></a>func doubler(n interface{}) interface{} { return n.(int) * 2 }
<a id="L31"></a>func addOne(n interface{}) interface{}  { return n.(int) + 1 }
<a id="L32"></a>func adder(acc interface{}, n interface{}) interface{} {
    <a id="L33"></a>return acc.(int) + n.(int)
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// A stream of the natural numbers: 0, 1, 2, 3, ...</span>
<a id="L37"></a>type integerStream struct{}

<a id="L39"></a>func (i integerStream) Iter() &lt;-chan interface{} {
    <a id="L40"></a>ch := make(chan interface{});
    <a id="L41"></a>go func() {
        <a id="L42"></a>for i := 0; ; i++ {
            <a id="L43"></a>ch &lt;- i
        <a id="L44"></a>}
    <a id="L45"></a>}();
    <a id="L46"></a>return ch;
<a id="L47"></a>}

<a id="L49"></a>func TestAll(t *testing.T) {
    <a id="L50"></a>if !All(oneToFive, isPositive) {
        <a id="L51"></a>t.Error(&#34;All(oneToFive, isPositive) == false&#34;)
    <a id="L52"></a>}
    <a id="L53"></a>if All(oneToFive, isAbove3) {
        <a id="L54"></a>t.Error(&#34;All(oneToFive, isAbove3) == true&#34;)
    <a id="L55"></a>}
<a id="L56"></a>}

<a id="L58"></a>func TestAny(t *testing.T) {
    <a id="L59"></a>if Any(oneToFive, isNegative) {
        <a id="L60"></a>t.Error(&#34;Any(oneToFive, isNegative) == true&#34;)
    <a id="L61"></a>}
    <a id="L62"></a>if !Any(oneToFive, isEven) {
        <a id="L63"></a>t.Error(&#34;Any(oneToFive, isEven) == false&#34;)
    <a id="L64"></a>}
<a id="L65"></a>}

<a id="L67"></a>func assertArraysAreEqual(t *testing.T, res []interface{}, expected []int) {
    <a id="L68"></a>if len(res) != len(expected) {
        <a id="L69"></a>t.Errorf(&#34;len(res) = %v, want %v&#34;, len(res), len(expected));
        <a id="L70"></a>goto missing;
    <a id="L71"></a>}
    <a id="L72"></a>for i := range res {
        <a id="L73"></a>if v := res[i].(int); v != expected[i] {
            <a id="L74"></a>t.Errorf(&#34;res[%v] = %v, want %v&#34;, i, v, expected[i]);
            <a id="L75"></a>goto missing;
        <a id="L76"></a>}
    <a id="L77"></a>}
    <a id="L78"></a>return;
<a id="L79"></a>missing:
    <a id="L80"></a>t.Errorf(&#34;res = %v\nwant  %v&#34;, res, expected);
<a id="L81"></a>}

<a id="L83"></a>func TestFilter(t *testing.T) {
    <a id="L84"></a>ints := integerStream{};
    <a id="L85"></a>moreInts := Filter(ints, isAbove3).Iter();
    <a id="L86"></a>res := make([]interface{}, 3);
    <a id="L87"></a>for i := 0; i &lt; 3; i++ {
        <a id="L88"></a>res[i] = &lt;-moreInts
    <a id="L89"></a>}
    <a id="L90"></a>assertArraysAreEqual(t, res, []int{4, 5, 6});
<a id="L91"></a>}

<a id="L93"></a>func TestFind(t *testing.T) {
    <a id="L94"></a>ints := integerStream{};
    <a id="L95"></a>first := Find(ints, isAbove3);
    <a id="L96"></a>if first.(int) != 4 {
        <a id="L97"></a>t.Errorf(&#34;Find(ints, isAbove3) = %v, want 4&#34;, first)
    <a id="L98"></a>}
<a id="L99"></a>}

<a id="L101"></a>func TestInject(t *testing.T) {
    <a id="L102"></a>res := Inject(oneToFive, 0, adder);
    <a id="L103"></a>if res.(int) != 15 {
        <a id="L104"></a>t.Errorf(&#34;Inject(oneToFive, 0, adder) = %v, want 15&#34;, res)
    <a id="L105"></a>}
<a id="L106"></a>}

<a id="L108"></a>func TestMap(t *testing.T) {
    <a id="L109"></a>res := Data(Map(Map(oneToFive, doubler), addOne));
    <a id="L110"></a>assertArraysAreEqual(t, res, []int{3, 5, 7, 9, 11});
<a id="L111"></a>}

<a id="L113"></a>func TestPartition(t *testing.T) {
    <a id="L114"></a>ti, fi := Partition(oneToFive, isEven);
    <a id="L115"></a>assertArraysAreEqual(t, Data(ti), []int{2, 4});
    <a id="L116"></a>assertArraysAreEqual(t, Data(fi), []int{1, 3, 5});
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
