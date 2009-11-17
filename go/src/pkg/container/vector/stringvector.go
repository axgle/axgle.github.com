<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/vector/stringvector.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/container/vector/stringvector.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package vector

<a id="L7"></a><span class="comment">// StringVector is a specialization of Vector that hides the wrapping of Elements around strings.</span>
<a id="L8"></a>type StringVector struct {
    <a id="L9"></a>Vector;
<a id="L10"></a>}


<a id="L13"></a><span class="comment">// Init initializes a new or resized vector.  The initial length may be &lt;= 0 to</span>
<a id="L14"></a><span class="comment">// request a default length.  If initial_len is shorter than the current</span>
<a id="L15"></a><span class="comment">// length of the StringVector, trailing elements of the StringVector will be cleared.</span>
<a id="L16"></a>func (p *StringVector) Init(len int) *StringVector {
    <a id="L17"></a>p.Vector.Init(len);
    <a id="L18"></a>return p;
<a id="L19"></a>}


<a id="L22"></a><span class="comment">// NewStringVector returns an initialized new StringVector with length at least len.</span>
<a id="L23"></a>func NewStringVector(len int) *StringVector { return new(StringVector).Init(len) }


<a id="L26"></a><span class="comment">// At returns the i&#39;th element of the vector.</span>
<a id="L27"></a>func (p *StringVector) At(i int) string { return p.Vector.At(i).(string) }


<a id="L30"></a><span class="comment">// Set sets the i&#39;th element of the vector to value x.</span>
<a id="L31"></a>func (p *StringVector) Set(i int, x string) { p.a[i] = x }


<a id="L34"></a><span class="comment">// Last returns the element in the vector of highest index.</span>
<a id="L35"></a>func (p *StringVector) Last() string { return p.Vector.Last().(string) }


<a id="L38"></a><span class="comment">// Data returns all the elements as a slice.</span>
<a id="L39"></a>func (p *StringVector) Data() []string {
    <a id="L40"></a>arr := make([]string, p.Len());
    <a id="L41"></a>for i, v := range p.a {
        <a id="L42"></a>arr[i] = v.(string)
    <a id="L43"></a>}
    <a id="L44"></a>return arr;
<a id="L45"></a>}


<a id="L48"></a><span class="comment">// Insert inserts into the vector an element of value x before</span>
<a id="L49"></a><span class="comment">// the current element at index i.</span>
<a id="L50"></a>func (p *StringVector) Insert(i int, x string) {
    <a id="L51"></a>p.Vector.Insert(i, x)
<a id="L52"></a>}


<a id="L55"></a><span class="comment">// InsertVector inserts into the vector the contents of the Vector</span>
<a id="L56"></a><span class="comment">// x such that the 0th element of x appears at index i after insertion.</span>
<a id="L57"></a>func (p *StringVector) InsertVector(i int, x *StringVector) {
    <a id="L58"></a>p.Vector.InsertVector(i, &amp;x.Vector)
<a id="L59"></a>}


<a id="L62"></a><span class="comment">// Slice returns a new StringVector by slicing the old one to extract slice [i:j].</span>
<a id="L63"></a><span class="comment">// The elements are copied. The original vector is unchanged.</span>
<a id="L64"></a>func (p *StringVector) Slice(i, j int) *StringVector {
    <a id="L65"></a>return &amp;StringVector{*p.Vector.Slice(i, j)}
<a id="L66"></a>}


<a id="L69"></a><span class="comment">// Push appends x to the end of the vector.</span>
<a id="L70"></a>func (p *StringVector) Push(x string) { p.Vector.Push(x) }


<a id="L73"></a><span class="comment">// Pop deletes and returns the last element of the vector.</span>
<a id="L74"></a>func (p *StringVector) Pop() string { return p.Vector.Pop().(string) }


<a id="L77"></a><span class="comment">// AppendVector appends the entire StringVector x to the end of this vector.</span>
<a id="L78"></a>func (p *StringVector) AppendVector(x *StringVector) {
    <a id="L79"></a>p.Vector.InsertVector(len(p.a), &amp;x.Vector)
<a id="L80"></a>}


<a id="L83"></a><span class="comment">// sort.Interface support</span>
<a id="L84"></a><span class="comment">// Less returns a boolean denoting whether the i&#39;th element is less than the j&#39;th element.</span>
<a id="L85"></a>func (p *StringVector) Less(i, j int) bool { return p.At(i) &lt; p.At(j) }


<a id="L88"></a><span class="comment">// Iterate over all elements; driver for range</span>
<a id="L89"></a>func (p *StringVector) iterate(c chan&lt;- string) {
    <a id="L90"></a>for _, v := range p.a {
        <a id="L91"></a>c &lt;- v.(string)
    <a id="L92"></a>}
    <a id="L93"></a>close(c);
<a id="L94"></a>}


<a id="L97"></a><span class="comment">// Channel iterator for range.</span>
<a id="L98"></a>func (p *StringVector) Iter() &lt;-chan string {
    <a id="L99"></a>c := make(chan string);
    <a id="L100"></a>go p.iterate(c);
    <a id="L101"></a>return c;
<a id="L102"></a>}
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
