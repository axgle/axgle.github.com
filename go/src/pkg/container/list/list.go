<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/list/list.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/container/list/list.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The list package implements a doubly linked list.</span>
<a id="L6"></a>package list

<a id="L8"></a><span class="comment">// Element is an element in the linked list.</span>
<a id="L9"></a>type Element struct {
    <a id="L10"></a><span class="comment">// Next and previous pointers in the doubly-linked list of elements.</span>
    <a id="L11"></a><span class="comment">// The front of the list has prev = nil, and the back has next = nil.</span>
    <a id="L12"></a>next, prev *Element;

    <a id="L14"></a><span class="comment">// A unique ID for the list to which this element belongs.</span>
    <a id="L15"></a>id  *byte;

    <a id="L17"></a><span class="comment">// The contents of this list element.</span>
    <a id="L18"></a>Value interface{};
<a id="L19"></a>}

<a id="L21"></a><span class="comment">// Next returns the next list element or nil.</span>
<a id="L22"></a>func (e *Element) Next() *Element { return e.next }

<a id="L24"></a><span class="comment">// Prev returns the previous list element or nil.</span>
<a id="L25"></a>func (e *Element) Prev() *Element { return e.prev }

<a id="L27"></a><span class="comment">// List represents a doubly linked list.</span>
<a id="L28"></a>type List struct {
    <a id="L29"></a>front, back *Element;
    <a id="L30"></a>len         int;
    <a id="L31"></a>id          *byte;
<a id="L32"></a>}

<a id="L34"></a><span class="comment">// Init initializes or clears a List.</span>
<a id="L35"></a>func (l *List) Init() *List {
    <a id="L36"></a>l.front = nil;
    <a id="L37"></a>l.back = nil;
    <a id="L38"></a>l.len = 0;
    <a id="L39"></a>l.id = new(byte);
    <a id="L40"></a>return l;
<a id="L41"></a>}

<a id="L43"></a><span class="comment">// New returns an initialized list.</span>
<a id="L44"></a>func New() *List { return new(List).Init() }

<a id="L46"></a><span class="comment">// Front returns the first element in the list.</span>
<a id="L47"></a>func (l *List) Front() *Element { return l.front }

<a id="L49"></a><span class="comment">// Back returns the last element in the list.</span>
<a id="L50"></a>func (l *List) Back() *Element { return l.back }

<a id="L52"></a><span class="comment">// Remove removes the element from the list.</span>
<a id="L53"></a>func (l *List) Remove(e *Element) {
    <a id="L54"></a>if e.id != l.id {
        <a id="L55"></a>return
    <a id="L56"></a>}
    <a id="L57"></a>if e.prev == nil {
        <a id="L58"></a>l.front = e.next
    <a id="L59"></a>} else {
        <a id="L60"></a>e.prev.next = e.next
    <a id="L61"></a>}
    <a id="L62"></a>if e.next == nil {
        <a id="L63"></a>l.back = e.prev
    <a id="L64"></a>} else {
        <a id="L65"></a>e.next.prev = e.prev
    <a id="L66"></a>}

    <a id="L68"></a>e.prev = nil;
    <a id="L69"></a>e.next = nil;
    <a id="L70"></a>l.len--;
<a id="L71"></a>}

<a id="L73"></a>func (l *List) insertBefore(e *Element, mark *Element) {
    <a id="L74"></a>if mark.prev == nil {
        <a id="L75"></a><span class="comment">// new front of the list</span>
        <a id="L76"></a>l.front = e
    <a id="L77"></a>} else {
        <a id="L78"></a>mark.prev.next = e
    <a id="L79"></a>}
    <a id="L80"></a>e.prev = mark.prev;
    <a id="L81"></a>mark.prev = e;
    <a id="L82"></a>e.next = mark;
    <a id="L83"></a>l.len++;
<a id="L84"></a>}

<a id="L86"></a>func (l *List) insertAfter(e *Element, mark *Element) {
    <a id="L87"></a>if mark.next == nil {
        <a id="L88"></a><span class="comment">// new back of the list</span>
        <a id="L89"></a>l.back = e
    <a id="L90"></a>} else {
        <a id="L91"></a>mark.next.prev = e
    <a id="L92"></a>}
    <a id="L93"></a>e.next = mark.next;
    <a id="L94"></a>mark.next = e;
    <a id="L95"></a>e.prev = mark;
    <a id="L96"></a>l.len++;
<a id="L97"></a>}

<a id="L99"></a>func (l *List) insertFront(e *Element) {
    <a id="L100"></a>if l.front == nil {
        <a id="L101"></a><span class="comment">// empty list</span>
        <a id="L102"></a>l.front, l.back = e, e;
        <a id="L103"></a>e.prev, e.next = nil, nil;
        <a id="L104"></a>l.len = 1;
        <a id="L105"></a>return;
    <a id="L106"></a>}
    <a id="L107"></a>l.insertBefore(e, l.front);
<a id="L108"></a>}

<a id="L110"></a>func (l *List) insertBack(e *Element) {
    <a id="L111"></a>if l.back == nil {
        <a id="L112"></a><span class="comment">// empty list</span>
        <a id="L113"></a>l.front, l.back = e, e;
        <a id="L114"></a>e.prev, e.next = nil, nil;
        <a id="L115"></a>l.len = 1;
        <a id="L116"></a>return;
    <a id="L117"></a>}
    <a id="L118"></a>l.insertAfter(e, l.back);
<a id="L119"></a>}

<a id="L121"></a><span class="comment">// PushFront inserts the value at the front of the list and returns a new Element containing the value.</span>
<a id="L122"></a>func (l *List) PushFront(value interface{}) *Element {
    <a id="L123"></a>if l.id == nil {
        <a id="L124"></a>l.Init()
    <a id="L125"></a>}
    <a id="L126"></a>e := &amp;Element{nil, nil, l.id, value};
    <a id="L127"></a>l.insertFront(e);
    <a id="L128"></a>return e;
<a id="L129"></a>}

<a id="L131"></a><span class="comment">// PushBack inserts the value at the back of the list and returns a new Element containing the value.</span>
<a id="L132"></a>func (l *List) PushBack(value interface{}) *Element {
    <a id="L133"></a>if l.id == nil {
        <a id="L134"></a>l.Init()
    <a id="L135"></a>}
    <a id="L136"></a>e := &amp;Element{nil, nil, l.id, value};
    <a id="L137"></a>l.insertBack(e);
    <a id="L138"></a>return e;
<a id="L139"></a>}

<a id="L141"></a><span class="comment">// InsertBefore inserts the value immediately before mark and returns a new Element containing the value.</span>
<a id="L142"></a>func (l *List) InsertBefore(value interface{}, mark *Element) *Element {
    <a id="L143"></a>if mark.id != l.id {
        <a id="L144"></a>return nil
    <a id="L145"></a>}
    <a id="L146"></a>e := &amp;Element{nil, nil, l.id, value};
    <a id="L147"></a>l.insertBefore(e, mark);
    <a id="L148"></a>return e;
<a id="L149"></a>}

<a id="L151"></a><span class="comment">// InsertAfter inserts the value immediately after mark and returns a new Element containing the value.</span>
<a id="L152"></a>func (l *List) InsertAfter(value interface{}, mark *Element) *Element {
    <a id="L153"></a>if mark.id != l.id {
        <a id="L154"></a>return nil
    <a id="L155"></a>}
    <a id="L156"></a>e := &amp;Element{nil, nil, l.id, value};
    <a id="L157"></a>l.insertAfter(e, mark);
    <a id="L158"></a>return e;
<a id="L159"></a>}

<a id="L161"></a><span class="comment">// MoveToFront moves the element to the front of the list.</span>
<a id="L162"></a>func (l *List) MoveToFront(e *Element) {
    <a id="L163"></a>if e.id != l.id || l.front == e {
        <a id="L164"></a>return
    <a id="L165"></a>}
    <a id="L166"></a>l.Remove(e);
    <a id="L167"></a>l.insertFront(e);
<a id="L168"></a>}

<a id="L170"></a><span class="comment">// MoveToBack moves the element to the back of the list.</span>
<a id="L171"></a>func (l *List) MoveToBack(e *Element) {
    <a id="L172"></a>if e.id != l.id || l.back == e {
        <a id="L173"></a>return
    <a id="L174"></a>}
    <a id="L175"></a>l.Remove(e);
    <a id="L176"></a>l.insertBack(e);
<a id="L177"></a>}

<a id="L179"></a><span class="comment">// Len returns the number of elements in the list.</span>
<a id="L180"></a>func (l *List) Len() int { return l.len }

<a id="L182"></a>func (l *List) iterate(c chan&lt;- interface{}) {
    <a id="L183"></a>for e := l.front; e != nil; e = e.next {
        <a id="L184"></a>c &lt;- e.Value
    <a id="L185"></a>}
    <a id="L186"></a>close(c);
<a id="L187"></a>}

<a id="L189"></a>func (l *List) Iter() &lt;-chan interface{} {
    <a id="L190"></a>c := make(chan interface{});
    <a id="L191"></a>go l.iterate(c);
    <a id="L192"></a>return c;
<a id="L193"></a>}
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
