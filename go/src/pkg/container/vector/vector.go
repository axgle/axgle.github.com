<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/vector/vector.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/container/vector/vector.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The vector package implements an efficient container for managing</span>
<a id="L6"></a><span class="comment">// linear arrays of elements.  Unlike arrays, vectors can change size dynamically.</span>
<a id="L7"></a>package vector

<a id="L9"></a><span class="comment">// Vector is the container itself.</span>
<a id="L10"></a><span class="comment">// The zero value for Vector is an empty vector ready to use.</span>
<a id="L11"></a>type Vector struct {
    <a id="L12"></a>a []interface{};
<a id="L13"></a>}


<a id="L16"></a>func copy(dst, src []interface{}) {
    <a id="L17"></a>for i, x := range src {
        <a id="L18"></a>dst[i] = x
    <a id="L19"></a>}
<a id="L20"></a>}


<a id="L23"></a><span class="comment">// Insert n elements at position i.</span>
<a id="L24"></a>func expand(a []interface{}, i, n int) []interface{} {
    <a id="L25"></a><span class="comment">// make sure we have enough space</span>
    <a id="L26"></a>len0 := len(a);
    <a id="L27"></a>len1 := len0 + n;
    <a id="L28"></a>if len1 &lt; cap(a) {
        <a id="L29"></a><span class="comment">// enough space - just expand</span>
        <a id="L30"></a>a = a[0:len1]
    <a id="L31"></a>} else {
        <a id="L32"></a><span class="comment">// not enough space - double capacity</span>
        <a id="L33"></a>capb := cap(a) * 2;
        <a id="L34"></a>if capb &lt; len1 {
            <a id="L35"></a><span class="comment">// still not enough - use required length</span>
            <a id="L36"></a>capb = len1
        <a id="L37"></a>}
        <a id="L38"></a><span class="comment">// capb &gt;= len1</span>
        <a id="L39"></a>b := make([]interface{}, len1, capb);
        <a id="L40"></a>copy(b, a);
        <a id="L41"></a>a = b;
    <a id="L42"></a>}

    <a id="L44"></a><span class="comment">// make a hole</span>
    <a id="L45"></a>for j := len0 - 1; j &gt;= i; j-- {
        <a id="L46"></a>a[j+n] = a[j]
    <a id="L47"></a>}
    <a id="L48"></a>return a;
<a id="L49"></a>}


<a id="L52"></a><span class="comment">// Init initializes a new or resized vector.  The initial_len may be &lt;= 0 to</span>
<a id="L53"></a><span class="comment">// request a default length.  If initial_len is shorter than the current</span>
<a id="L54"></a><span class="comment">// length of the Vector, trailing elements of the Vector will be cleared.</span>
<a id="L55"></a>func (p *Vector) Init(initial_len int) *Vector {
    <a id="L56"></a>a := p.a;

    <a id="L58"></a>if cap(a) == 0 || cap(a) &lt; initial_len {
        <a id="L59"></a>n := 8; <span class="comment">// initial capacity</span>
        <a id="L60"></a>if initial_len &gt; n {
            <a id="L61"></a>n = initial_len
        <a id="L62"></a>}
        <a id="L63"></a>a = make([]interface{}, n);
    <a id="L64"></a>} else {
        <a id="L65"></a><span class="comment">// nil out entries</span>
        <a id="L66"></a>for j := len(a) - 1; j &gt;= 0; j-- {
            <a id="L67"></a>a[j] = nil
        <a id="L68"></a>}
    <a id="L69"></a>}

    <a id="L71"></a>p.a = a[0:initial_len];
    <a id="L72"></a>return p;
<a id="L73"></a>}


<a id="L76"></a><span class="comment">// New returns an initialized new Vector with length at least len.</span>
<a id="L77"></a>func New(len int) *Vector { return new(Vector).Init(len) }


<a id="L80"></a><span class="comment">// Len returns the number of elements in the vector.</span>
<a id="L81"></a><span class="comment">// Len is 0 if p == nil.</span>
<a id="L82"></a>func (p *Vector) Len() int {
    <a id="L83"></a>if p == nil {
        <a id="L84"></a>return 0
    <a id="L85"></a>}
    <a id="L86"></a>return len(p.a);
<a id="L87"></a>}


<a id="L90"></a><span class="comment">// At returns the i&#39;th element of the vector.</span>
<a id="L91"></a>func (p *Vector) At(i int) interface{} { return p.a[i] }


<a id="L94"></a><span class="comment">// Set sets the i&#39;th element of the vector to value x.</span>
<a id="L95"></a>func (p *Vector) Set(i int, x interface{}) { p.a[i] = x }


<a id="L98"></a><span class="comment">// Last returns the element in the vector of highest index.</span>
<a id="L99"></a>func (p *Vector) Last() interface{} { return p.a[len(p.a)-1] }


<a id="L102"></a><span class="comment">// Data returns all the elements as a slice.</span>
<a id="L103"></a>func (p *Vector) Data() []interface{} {
    <a id="L104"></a>arr := make([]interface{}, p.Len());
    <a id="L105"></a>for i, v := range p.a {
        <a id="L106"></a>arr[i] = v
    <a id="L107"></a>}
    <a id="L108"></a>return arr;
<a id="L109"></a>}


<a id="L112"></a><span class="comment">// Insert inserts into the vector an element of value x before</span>
<a id="L113"></a><span class="comment">// the current element at index i.</span>
<a id="L114"></a>func (p *Vector) Insert(i int, x interface{}) {
    <a id="L115"></a>p.a = expand(p.a, i, 1);
    <a id="L116"></a>p.a[i] = x;
<a id="L117"></a>}


<a id="L120"></a><span class="comment">// Delete deletes the i&#39;th element of the vector.  The gap is closed so the old</span>
<a id="L121"></a><span class="comment">// element at index i+1 has index i afterwards.</span>
<a id="L122"></a>func (p *Vector) Delete(i int) {
    <a id="L123"></a>a := p.a;
    <a id="L124"></a>n := len(a);

    <a id="L126"></a>copy(a[i:n-1], a[i+1:n]);
    <a id="L127"></a>a[n-1] = nil; <span class="comment">// support GC, nil out entry</span>
    <a id="L128"></a>p.a = a[0 : n-1];
<a id="L129"></a>}


<a id="L132"></a><span class="comment">// InsertVector inserts into the vector the contents of the Vector</span>
<a id="L133"></a><span class="comment">// x such that the 0th element of x appears at index i after insertion.</span>
<a id="L134"></a>func (p *Vector) InsertVector(i int, x *Vector) {
    <a id="L135"></a>p.a = expand(p.a, i, len(x.a));
    <a id="L136"></a>copy(p.a[i:i+len(x.a)], x.a);
<a id="L137"></a>}


<a id="L140"></a><span class="comment">// Cut deletes elements i through j-1, inclusive.</span>
<a id="L141"></a>func (p *Vector) Cut(i, j int) {
    <a id="L142"></a>a := p.a;
    <a id="L143"></a>n := len(a);
    <a id="L144"></a>m := n - (j - i);

    <a id="L146"></a>copy(a[i:m], a[j:n]);
    <a id="L147"></a>for k := m; k &lt; n; k++ {
        <a id="L148"></a>a[k] = nil <span class="comment">// support GC, nil out entries</span>
    <a id="L149"></a>}

    <a id="L151"></a>p.a = a[0:m];
<a id="L152"></a>}


<a id="L155"></a><span class="comment">// Slice returns a new Vector by slicing the old one to extract slice [i:j].</span>
<a id="L156"></a><span class="comment">// The elements are copied. The original vector is unchanged.</span>
<a id="L157"></a>func (p *Vector) Slice(i, j int) *Vector {
    <a id="L158"></a>s := New(j - i); <span class="comment">// will fail in Init() if j &lt; j</span>
    <a id="L159"></a>copy(s.a, p.a[i:j]);
    <a id="L160"></a>return s;
<a id="L161"></a>}


<a id="L164"></a><span class="comment">// Do calls function f for each element of the vector, in order.</span>
<a id="L165"></a><span class="comment">// The function should not change the indexing of the vector underfoot.</span>
<a id="L166"></a>func (p *Vector) Do(f func(elem interface{})) {
    <a id="L167"></a>for i := 0; i &lt; len(p.a); i++ {
        <a id="L168"></a>f(p.a[i]) <span class="comment">// not too safe if f changes the Vector</span>
    <a id="L169"></a>}
<a id="L170"></a>}


<a id="L173"></a><span class="comment">// Convenience wrappers</span>

<a id="L175"></a><span class="comment">// Push appends x to the end of the vector.</span>
<a id="L176"></a>func (p *Vector) Push(x interface{}) { p.Insert(len(p.a), x) }


<a id="L179"></a><span class="comment">// Pop deletes the last element of the vector.</span>
<a id="L180"></a>func (p *Vector) Pop() interface{} {
    <a id="L181"></a>i := len(p.a) - 1;
    <a id="L182"></a>x := p.a[i];
    <a id="L183"></a>p.a[i] = nil; <span class="comment">// support GC, nil out entry</span>
    <a id="L184"></a>p.a = p.a[0:i];
    <a id="L185"></a>return x;
<a id="L186"></a>}


<a id="L189"></a><span class="comment">// AppendVector appends the entire Vector x to the end of this vector.</span>
<a id="L190"></a>func (p *Vector) AppendVector(x *Vector) { p.InsertVector(len(p.a), x) }


<a id="L193"></a><span class="comment">// Partial sort.Interface support</span>

<a id="L195"></a><span class="comment">// LessInterface provides partial support of the sort.Interface.</span>
<a id="L196"></a>type LessInterface interface {
    <a id="L197"></a>Less(y interface{}) bool;
<a id="L198"></a>}


<a id="L201"></a><span class="comment">// Less returns a boolean denoting whether the i&#39;th element is less than the j&#39;th element.</span>
<a id="L202"></a>func (p *Vector) Less(i, j int) bool { return p.a[i].(LessInterface).Less(p.a[j]) }


<a id="L205"></a><span class="comment">// Swap exchanges the elements at indexes i and j.</span>
<a id="L206"></a>func (p *Vector) Swap(i, j int) {
    <a id="L207"></a>a := p.a;
    <a id="L208"></a>a[i], a[j] = a[j], a[i];
<a id="L209"></a>}


<a id="L212"></a><span class="comment">// Iterate over all elements; driver for range</span>
<a id="L213"></a>func (p *Vector) iterate(c chan&lt;- interface{}) {
    <a id="L214"></a>for _, v := range p.a {
        <a id="L215"></a>c &lt;- v
    <a id="L216"></a>}
    <a id="L217"></a>close(c);
<a id="L218"></a>}


<a id="L221"></a><span class="comment">// Channel iterator for range.</span>
<a id="L222"></a>func (p *Vector) Iter() &lt;-chan interface{} {
    <a id="L223"></a>c := make(chan interface{});
    <a id="L224"></a>go p.iterate(c);
    <a id="L225"></a>return c;
<a id="L226"></a>}
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
