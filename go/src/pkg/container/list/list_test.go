<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/list/list_test.go</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/container/list/list_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package list

<a id="L7"></a>import (
    <a id="L8"></a>&#34;testing&#34;;
<a id="L9"></a>)

<a id="L11"></a>func checkListPointers(t *testing.T, l *List, es []*Element) {
    <a id="L12"></a>if len(es) == 0 {
        <a id="L13"></a>if l.front != nil || l.back != nil {
            <a id="L14"></a>t.Errorf(&#34;l.front/l.back = %v/%v should be nil/nil&#34;, l.front, l.back)
        <a id="L15"></a>}
        <a id="L16"></a>return;
    <a id="L17"></a>}

    <a id="L19"></a>if l.front != es[0] {
        <a id="L20"></a>t.Errorf(&#34;l.front = %v, want %v&#34;, l.front, es[0])
    <a id="L21"></a>}
    <a id="L22"></a>if last := es[len(es)-1]; l.back != last {
        <a id="L23"></a>t.Errorf(&#34;l.back = %v, want %v&#34;, l.back, last)
    <a id="L24"></a>}

    <a id="L26"></a>for i := 0; i &lt; len(es); i++ {
        <a id="L27"></a>e := es[i];
        <a id="L28"></a>var e_prev, e_next *Element = nil, nil;
        <a id="L29"></a>if i &gt; 0 {
            <a id="L30"></a>e_prev = es[i-1]
        <a id="L31"></a>}
        <a id="L32"></a>if i &lt; len(es)-1 {
            <a id="L33"></a>e_next = es[i+1]
        <a id="L34"></a>}
        <a id="L35"></a>if e.prev != e_prev {
            <a id="L36"></a>t.Errorf(&#34;elt #%d (%v) has prev=%v, want %v&#34;, i, e, e.prev, e_prev)
        <a id="L37"></a>}
        <a id="L38"></a>if e.next != e_next {
            <a id="L39"></a>t.Errorf(&#34;elt #%d (%v) has next=%v, want %v&#34;, i, e, e.next, e_next)
        <a id="L40"></a>}
    <a id="L41"></a>}
<a id="L42"></a>}

<a id="L44"></a>func checkListLen(t *testing.T, l *List, n int) {
    <a id="L45"></a>if an := l.Len(); an != n {
        <a id="L46"></a>t.Errorf(&#34;l.Len() = %d, want %d&#34;, an, n)
    <a id="L47"></a>}
<a id="L48"></a>}

<a id="L50"></a>func TestList(t *testing.T) {
    <a id="L51"></a>l := New();
    <a id="L52"></a>checkListPointers(t, l, []*Element{});
    <a id="L53"></a>checkListLen(t, l, 0);

    <a id="L55"></a><span class="comment">// Single element list</span>
    <a id="L56"></a>e := l.PushFront(&#34;a&#34;);
    <a id="L57"></a>checkListLen(t, l, 1);
    <a id="L58"></a>checkListPointers(t, l, []*Element{e});
    <a id="L59"></a>l.MoveToFront(e);
    <a id="L60"></a>checkListPointers(t, l, []*Element{e});
    <a id="L61"></a>l.MoveToBack(e);
    <a id="L62"></a>checkListPointers(t, l, []*Element{e});
    <a id="L63"></a>checkListLen(t, l, 1);
    <a id="L64"></a>l.Remove(e);
    <a id="L65"></a>checkListPointers(t, l, []*Element{});
    <a id="L66"></a>checkListLen(t, l, 0);

    <a id="L68"></a><span class="comment">// Bigger list</span>
    <a id="L69"></a>e2 := l.PushFront(2);
    <a id="L70"></a>e1 := l.PushFront(1);
    <a id="L71"></a>e3 := l.PushBack(3);
    <a id="L72"></a>e4 := l.PushBack(&#34;banana&#34;);
    <a id="L73"></a>checkListPointers(t, l, []*Element{e1, e2, e3, e4});
    <a id="L74"></a>checkListLen(t, l, 4);

    <a id="L76"></a>l.Remove(e2);
    <a id="L77"></a>checkListPointers(t, l, []*Element{e1, e3, e4});
    <a id="L78"></a>checkListLen(t, l, 3);

    <a id="L80"></a>l.MoveToFront(e3); <span class="comment">// move from middle</span>
    <a id="L81"></a>checkListPointers(t, l, []*Element{e3, e1, e4});

    <a id="L83"></a>l.MoveToFront(e1);
    <a id="L84"></a>l.MoveToBack(e3); <span class="comment">// move from middle</span>
    <a id="L85"></a>checkListPointers(t, l, []*Element{e1, e4, e3});

    <a id="L87"></a>l.MoveToFront(e3); <span class="comment">// move from back</span>
    <a id="L88"></a>checkListPointers(t, l, []*Element{e3, e1, e4});
    <a id="L89"></a>l.MoveToFront(e3); <span class="comment">// should be no-op</span>
    <a id="L90"></a>checkListPointers(t, l, []*Element{e3, e1, e4});

    <a id="L92"></a>l.MoveToBack(e3); <span class="comment">// move from front</span>
    <a id="L93"></a>checkListPointers(t, l, []*Element{e1, e4, e3});
    <a id="L94"></a>l.MoveToBack(e3); <span class="comment">// should be no-op</span>
    <a id="L95"></a>checkListPointers(t, l, []*Element{e1, e4, e3});

    <a id="L97"></a>e2 = l.InsertBefore(2, e1); <span class="comment">// insert before front</span>
    <a id="L98"></a>checkListPointers(t, l, []*Element{e2, e1, e4, e3});
    <a id="L99"></a>l.Remove(e2);
    <a id="L100"></a>e2 = l.InsertBefore(2, e4); <span class="comment">// insert before middle</span>
    <a id="L101"></a>checkListPointers(t, l, []*Element{e1, e2, e4, e3});
    <a id="L102"></a>l.Remove(e2);
    <a id="L103"></a>e2 = l.InsertBefore(2, e3); <span class="comment">// insert before back</span>
    <a id="L104"></a>checkListPointers(t, l, []*Element{e1, e4, e2, e3});
    <a id="L105"></a>l.Remove(e2);

    <a id="L107"></a>e2 = l.InsertAfter(2, e1); <span class="comment">// insert after front</span>
    <a id="L108"></a>checkListPointers(t, l, []*Element{e1, e2, e4, e3});
    <a id="L109"></a>l.Remove(e2);
    <a id="L110"></a>e2 = l.InsertAfter(2, e4); <span class="comment">// insert after middle</span>
    <a id="L111"></a>checkListPointers(t, l, []*Element{e1, e4, e2, e3});
    <a id="L112"></a>l.Remove(e2);
    <a id="L113"></a>e2 = l.InsertAfter(2, e3); <span class="comment">// insert after back</span>
    <a id="L114"></a>checkListPointers(t, l, []*Element{e1, e4, e3, e2});
    <a id="L115"></a>l.Remove(e2);

    <a id="L117"></a><span class="comment">// Check standard iteration.</span>
    <a id="L118"></a>sum := 0;
    <a id="L119"></a>for e := range l.Iter() {
        <a id="L120"></a>if i, ok := e.(int); ok {
            <a id="L121"></a>sum += i
        <a id="L122"></a>}
    <a id="L123"></a>}
    <a id="L124"></a>if sum != 4 {
        <a id="L125"></a>t.Errorf(&#34;sum over l.Iter() = %d, want 4&#34;, sum)
    <a id="L126"></a>}

    <a id="L128"></a><span class="comment">// Clear all elements by iterating</span>
    <a id="L129"></a>var next *Element;
    <a id="L130"></a>for e := l.Front(); e != nil; e = next {
        <a id="L131"></a>next = e.Next();
        <a id="L132"></a>l.Remove(e);
    <a id="L133"></a>}
    <a id="L134"></a>checkListPointers(t, l, []*Element{});
    <a id="L135"></a>checkListLen(t, l, 0);
<a id="L136"></a>}
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
