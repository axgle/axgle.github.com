<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/iterable/iterable.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/iterable/iterable.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The iterable package provides several traversal and searching methods.</span>
<a id="L6"></a><span class="comment">// It can be used on anything that satisfies the Iterable interface,</span>
<a id="L7"></a><span class="comment">// including vector, though certain functions, such as Map, can also be used on</span>
<a id="L8"></a><span class="comment">// something that would produce an infinite amount of data.</span>
<a id="L9"></a>package iterable

<a id="L11"></a>import &#34;container/vector&#34;

<a id="L13"></a>type Iterable interface {
    <a id="L14"></a><span class="comment">// Iter should return a fresh channel each time it is called.</span>
    <a id="L15"></a>Iter() &lt;-chan interface{};
<a id="L16"></a>}

<a id="L18"></a>func not(f func(interface{}) bool) (func(interface{}) bool) {
    <a id="L19"></a>return func(e interface{}) bool { return !f(e) }
<a id="L20"></a>}

<a id="L22"></a><span class="comment">// All tests whether f is true for every element of iter.</span>
<a id="L23"></a>func All(iter Iterable, f func(interface{}) bool) bool {
    <a id="L24"></a>for e := range iter.Iter() {
        <a id="L25"></a>if !f(e) {
            <a id="L26"></a>return false
        <a id="L27"></a>}
    <a id="L28"></a>}
    <a id="L29"></a>return true;
<a id="L30"></a>}

<a id="L32"></a><span class="comment">// Any tests whether f is true for at least one element of iter.</span>
<a id="L33"></a>func Any(iter Iterable, f func(interface{}) bool) bool {
    <a id="L34"></a>return !All(iter, not(f))
<a id="L35"></a>}

<a id="L37"></a><span class="comment">// Data returns a slice containing the elements of iter.</span>
<a id="L38"></a>func Data(iter Iterable) []interface{} {
    <a id="L39"></a>vec := vector.New(0);
    <a id="L40"></a>for e := range iter.Iter() {
        <a id="L41"></a>vec.Push(e)
    <a id="L42"></a>}
    <a id="L43"></a>return vec.Data();
<a id="L44"></a>}

<a id="L46"></a><span class="comment">// filteredIterable is a struct that implements Iterable with each element</span>
<a id="L47"></a><span class="comment">// passed through a filter.</span>
<a id="L48"></a>type filteredIterable struct {
    <a id="L49"></a>it  Iterable;
    <a id="L50"></a>f   func(interface{}) bool;
<a id="L51"></a>}

<a id="L53"></a>func (f *filteredIterable) iterate(out chan&lt;- interface{}) {
    <a id="L54"></a>for e := range f.it.Iter() {
        <a id="L55"></a>if f.f(e) {
            <a id="L56"></a>out &lt;- e
        <a id="L57"></a>}
    <a id="L58"></a>}
    <a id="L59"></a>close(out);
<a id="L60"></a>}

<a id="L62"></a>func (f *filteredIterable) Iter() &lt;-chan interface{} {
    <a id="L63"></a>ch := make(chan interface{});
    <a id="L64"></a>go f.iterate(ch);
    <a id="L65"></a>return ch;
<a id="L66"></a>}

<a id="L68"></a><span class="comment">// Filter returns an Iterable that returns the elements of iter that satisfy f.</span>
<a id="L69"></a>func Filter(iter Iterable, f func(interface{}) bool) Iterable {
    <a id="L70"></a>return &amp;filteredIterable{iter, f}
<a id="L71"></a>}

<a id="L73"></a><span class="comment">// Find returns the first element of iter that satisfies f.</span>
<a id="L74"></a><span class="comment">// Returns nil if no such element is found.</span>
<a id="L75"></a>func Find(iter Iterable, f func(interface{}) bool) interface{} {
    <a id="L76"></a>for e := range Filter(iter, f).Iter() {
        <a id="L77"></a>return e
    <a id="L78"></a>}
    <a id="L79"></a>return nil;
<a id="L80"></a>}

<a id="L82"></a><span class="comment">// Injector is a type representing a function that takes two arguments,</span>
<a id="L83"></a><span class="comment">// an accumulated value and an element, and returns the next accumulated value.</span>
<a id="L84"></a><span class="comment">// See the Inject function.</span>
<a id="L85"></a>type Injector func(interface{}, interface{}) interface{}

<a id="L87"></a><span class="comment">// Inject combines the elements of iter by repeatedly calling f with an</span>
<a id="L88"></a><span class="comment">// accumulated value and each element in order. The starting accumulated value</span>
<a id="L89"></a><span class="comment">// is initial, and after each call the accumulated value is set to the return</span>
<a id="L90"></a><span class="comment">// value of f. For instance, to compute a sum:</span>
<a id="L91"></a><span class="comment">//   var arr IntArray = []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };</span>
<a id="L92"></a><span class="comment">//   sum := iterable.Inject(arr, 0,</span>
<a id="L93"></a><span class="comment">//                          func(ax interface {}, x interface {}) interface {} {</span>
<a id="L94"></a><span class="comment">//                            return ax.(int) + x.(int) }).(int)</span>
<a id="L95"></a>func Inject(iter Iterable, initial interface{}, f Injector) interface{} {
    <a id="L96"></a>acc := initial;
    <a id="L97"></a>for e := range iter.Iter() {
        <a id="L98"></a>acc = f(acc, e)
    <a id="L99"></a>}
    <a id="L100"></a>return acc;
<a id="L101"></a>}

<a id="L103"></a><span class="comment">// mappedIterable is a helper struct that implements Iterable, returned by Map.</span>
<a id="L104"></a>type mappedIterable struct {
    <a id="L105"></a>it  Iterable;
    <a id="L106"></a>f   func(interface{}) interface{};
<a id="L107"></a>}

<a id="L109"></a>func (m *mappedIterable) iterate(out chan&lt;- interface{}) {
    <a id="L110"></a>for e := range m.it.Iter() {
        <a id="L111"></a>out &lt;- m.f(e)
    <a id="L112"></a>}
    <a id="L113"></a>close(out);
<a id="L114"></a>}

<a id="L116"></a>func (m *mappedIterable) Iter() &lt;-chan interface{} {
    <a id="L117"></a>ch := make(chan interface{});
    <a id="L118"></a>go m.iterate(ch);
    <a id="L119"></a>return ch;
<a id="L120"></a>}

<a id="L122"></a><span class="comment">// Map returns an Iterable that returns the result of applying f to each</span>
<a id="L123"></a><span class="comment">// element of iter.</span>
<a id="L124"></a>func Map(iter Iterable, f func(interface{}) interface{}) Iterable {
    <a id="L125"></a>return &amp;mappedIterable{iter, f}
<a id="L126"></a>}

<a id="L128"></a><span class="comment">// Partition(iter, f) returns Filter(iter, f) and Filter(iter, !f).</span>
<a id="L129"></a>func Partition(iter Iterable, f func(interface{}) bool) (Iterable, Iterable) {
    <a id="L130"></a>return Filter(iter, f), Filter(iter, not(f))
<a id="L131"></a>}

<a id="L133"></a><span class="comment">// TODO:</span>
<a id="L134"></a><span class="comment">// - Zip</span>
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
