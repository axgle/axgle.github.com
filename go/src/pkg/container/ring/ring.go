<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/ring/ring.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/container/ring/ring.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The ring package implements operations on circular lists.</span>
<a id="L6"></a>package ring

<a id="L8"></a><span class="comment">// A Ring is an element of a circular list, or ring.</span>
<a id="L9"></a><span class="comment">// Rings do not have a beginning or end; a pointer to any ring element</span>
<a id="L10"></a><span class="comment">// serves as reference to the entire ring. Empty rings are represented</span>
<a id="L11"></a><span class="comment">// as nil Ring pointers. The zero value for a Ring is a one-element</span>
<a id="L12"></a><span class="comment">// ring with a nil Value.</span>
<a id="L13"></a><span class="comment">//</span>
<a id="L14"></a>type Ring struct {
    <a id="L15"></a>next, prev *Ring;
    <a id="L16"></a>Value      interface{}; <span class="comment">// for use by client; untouched by this library</span>
<a id="L17"></a>}


<a id="L20"></a>func (r *Ring) init() *Ring {
    <a id="L21"></a>r.next = r;
    <a id="L22"></a>r.prev = r;
    <a id="L23"></a>return r;
<a id="L24"></a>}


<a id="L27"></a><span class="comment">// Next returns the next ring element. r must not be empty.</span>
<a id="L28"></a>func (r *Ring) Next() *Ring {
    <a id="L29"></a>if r.next == nil {
        <a id="L30"></a>return r.init()
    <a id="L31"></a>}
    <a id="L32"></a>return r.next;
<a id="L33"></a>}


<a id="L36"></a><span class="comment">// Prev returns the previous ring element. r must not be empty.</span>
<a id="L37"></a>func (r *Ring) Prev() *Ring {
    <a id="L38"></a>if r.next == nil {
        <a id="L39"></a>return r.init()
    <a id="L40"></a>}
    <a id="L41"></a>return r.prev;
<a id="L42"></a>}


<a id="L45"></a><span class="comment">// Move moves n % r.Len() elements backward (n &lt; 0) or forward (n &gt;= 0)</span>
<a id="L46"></a><span class="comment">// in the ring and returns that ring element. r must not be empty.</span>
<a id="L47"></a><span class="comment">//</span>
<a id="L48"></a>func (r *Ring) Move(n int) *Ring {
    <a id="L49"></a>if r.next == nil {
        <a id="L50"></a>return r.init()
    <a id="L51"></a>}
    <a id="L52"></a>switch {
    <a id="L53"></a>case n &lt; 0:
        <a id="L54"></a>for ; n &lt; 0; n++ {
            <a id="L55"></a>r = r.prev
        <a id="L56"></a>}
    <a id="L57"></a>case n &gt; 0:
        <a id="L58"></a>for ; n &gt; 0; n-- {
            <a id="L59"></a>r = r.next
        <a id="L60"></a>}
    <a id="L61"></a>}
    <a id="L62"></a>return r;
<a id="L63"></a>}


<a id="L66"></a><span class="comment">// New creates a ring of n elements.</span>
<a id="L67"></a>func New(n int) *Ring {
    <a id="L68"></a>if n &lt;= 0 {
        <a id="L69"></a>return nil
    <a id="L70"></a>}
    <a id="L71"></a>r := new(Ring);
    <a id="L72"></a>p := r;
    <a id="L73"></a>for i := 1; i &lt; n; i++ {
        <a id="L74"></a>p.next = &amp;Ring{prev: p};
        <a id="L75"></a>p = p.next;
    <a id="L76"></a>}
    <a id="L77"></a>p.next = r;
    <a id="L78"></a>r.prev = p;
    <a id="L79"></a>return r;
<a id="L80"></a>}


<a id="L83"></a><span class="comment">// Link connects ring r with with ring s such that r.Next()</span>
<a id="L84"></a><span class="comment">// becomes s and returns the original value for r.Next().</span>
<a id="L85"></a><span class="comment">// r must not be empty.</span>
<a id="L86"></a><span class="comment">//</span>
<a id="L87"></a><span class="comment">// If r and s point to the same ring, linking</span>
<a id="L88"></a><span class="comment">// them removes the elements between r and s from the ring.</span>
<a id="L89"></a><span class="comment">// The removed elements form a subring and the result is a</span>
<a id="L90"></a><span class="comment">// reference to that subring (if no elements were removed,</span>
<a id="L91"></a><span class="comment">// the result is still the original value for r.Next(),</span>
<a id="L92"></a><span class="comment">// and not nil).</span>
<a id="L93"></a><span class="comment">//</span>
<a id="L94"></a><span class="comment">// If r and s point to different rings, linking</span>
<a id="L95"></a><span class="comment">// them creates a single ring with the elements of s inserted</span>
<a id="L96"></a><span class="comment">// after r. The result points to the element following the</span>
<a id="L97"></a><span class="comment">// last element of s after insertion.</span>
<a id="L98"></a><span class="comment">//</span>
<a id="L99"></a>func (r *Ring) Link(s *Ring) *Ring {
    <a id="L100"></a>n := r.Next();
    <a id="L101"></a>if s != nil {
        <a id="L102"></a>p := s.Prev();
        <a id="L103"></a><span class="comment">// Note: Cannot use multiple assignment because</span>
        <a id="L104"></a><span class="comment">// evaluation order of LHS is not specified.</span>
        <a id="L105"></a>r.next = s;
        <a id="L106"></a>s.prev = r;
        <a id="L107"></a>n.prev = p;
        <a id="L108"></a>p.next = n;
    <a id="L109"></a>}
    <a id="L110"></a>return n;
<a id="L111"></a>}


<a id="L114"></a><span class="comment">// Unlink removes n % r.Len() elements from the ring r, starting</span>
<a id="L115"></a><span class="comment">// at r.Next(). If n % r.Len() == 0, r remains unchanged.</span>
<a id="L116"></a><span class="comment">// The result is the removed subring. r must not be empty.</span>
<a id="L117"></a><span class="comment">//</span>
<a id="L118"></a>func (r *Ring) Unlink(n int) *Ring {
    <a id="L119"></a>if n &lt;= 0 {
        <a id="L120"></a>return nil
    <a id="L121"></a>}
    <a id="L122"></a>return r.Link(r.Move(n + 1));
<a id="L123"></a>}


<a id="L126"></a><span class="comment">// Len computes the number of elements in ring r.</span>
<a id="L127"></a><span class="comment">// It executes in time proportional to the number of elements.</span>
<a id="L128"></a><span class="comment">//</span>
<a id="L129"></a>func (r *Ring) Len() int {
    <a id="L130"></a>n := 0;
    <a id="L131"></a>if r != nil {
        <a id="L132"></a>n = 1;
        <a id="L133"></a>for p := r.Next(); p != r; p = p.next {
            <a id="L134"></a>n++
        <a id="L135"></a>}
    <a id="L136"></a>}
    <a id="L137"></a>return n;
<a id="L138"></a>}


<a id="L141"></a>func (r *Ring) Iter() &lt;-chan interface{} {
    <a id="L142"></a>c := make(chan interface{});
    <a id="L143"></a>go func() {
        <a id="L144"></a>if r != nil {
            <a id="L145"></a>c &lt;- r.Value;
            <a id="L146"></a>for p := r.Next(); p != r; p = p.next {
                <a id="L147"></a>c &lt;- p.Value
            <a id="L148"></a>}
        <a id="L149"></a>}
        <a id="L150"></a>close(c);
    <a id="L151"></a>}();
    <a id="L152"></a>return c;
<a id="L153"></a>}
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
