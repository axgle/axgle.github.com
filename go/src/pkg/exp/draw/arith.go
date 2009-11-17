<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/draw/arith.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/draw/arith.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package draw

<a id="L7"></a><span class="comment">// A Point is an X, Y coordinate pair.</span>
<a id="L8"></a>type Point struct {
    <a id="L9"></a>X, Y int;
<a id="L10"></a>}

<a id="L12"></a><span class="comment">// ZP is the zero Point.</span>
<a id="L13"></a>var ZP Point

<a id="L15"></a><span class="comment">// A Rectangle contains the Points with Min.X &lt;= X &lt; Max.X, Min.Y &lt;= Y &lt; Max.Y.</span>
<a id="L16"></a>type Rectangle struct {
    <a id="L17"></a>Min, Max Point;
<a id="L18"></a>}

<a id="L20"></a><span class="comment">// ZR is the zero Rectangle.</span>
<a id="L21"></a>var ZR Rectangle

<a id="L23"></a><span class="comment">// Pt is shorthand for Point{X, Y}.</span>
<a id="L24"></a>func Pt(X, Y int) Point { return Point{X, Y} }

<a id="L26"></a><span class="comment">// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}.</span>
<a id="L27"></a>func Rect(x0, y0, x1, y1 int) Rectangle { return Rectangle{Point{x0, y0}, Point{x1, y1}} }

<a id="L29"></a><span class="comment">// Rpt is shorthand for Rectangle{min, max}.</span>
<a id="L30"></a>func Rpt(min, max Point) Rectangle { return Rectangle{min, max} }

<a id="L32"></a><span class="comment">// Add returns the sum of p and q: Pt(p.X+q.X, p.Y+q.Y).</span>
<a id="L33"></a>func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }

<a id="L35"></a><span class="comment">// Sub returns the difference of p and q: Pt(p.X-q.X, p.Y-q.Y).</span>
<a id="L36"></a>func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

<a id="L38"></a><span class="comment">// Mul returns p scaled by k: Pt(p.X*k p.Y*k).</span>
<a id="L39"></a>func (p Point) Mul(k int) Point { return Point{p.X * k, p.Y * k} }

<a id="L41"></a><span class="comment">// Div returns p divided by k: Pt(p.X/k, p.Y/k).</span>
<a id="L42"></a>func (p Point) Div(k int) Point { return Point{p.X / k, p.Y / k} }

<a id="L44"></a><span class="comment">// Eq returns true if p and q are equal.</span>
<a id="L45"></a>func (p Point) Eq(q Point) bool { return p.X == q.X &amp;&amp; p.Y == q.Y }

<a id="L47"></a><span class="comment">// Inset returns the rectangle r inset by n: Rect(r.Min.X+n, r.Min.Y+n, r.Max.X-n, r.Max.Y-n).</span>
<a id="L48"></a>func (r Rectangle) Inset(n int) Rectangle {
    <a id="L49"></a>return Rectangle{Point{r.Min.X + n, r.Min.Y + n}, Point{r.Max.X - n, r.Max.Y - n}}
<a id="L50"></a>}

<a id="L52"></a><span class="comment">// Add returns the rectangle r translated by p: Rpt(r.Min.Add(p), r.Max.Add(p)).</span>
<a id="L53"></a>func (r Rectangle) Add(p Point) Rectangle { return Rectangle{r.Min.Add(p), r.Max.Add(p)} }

<a id="L55"></a><span class="comment">// Sub returns the rectangle r translated by -p: Rpt(r.Min.Sub(p), r.Max.Sub(p)).</span>
<a id="L56"></a>func (r Rectangle) Sub(p Point) Rectangle { return Rectangle{r.Min.Sub(p), r.Max.Sub(p)} }

<a id="L58"></a><span class="comment">// Canon returns a canonical version of r: the returned rectangle</span>
<a id="L59"></a><span class="comment">// has Min.X &lt;= Max.X and Min.Y &lt;= Max.Y.</span>
<a id="L60"></a>func (r Rectangle) Canon() Rectangle {
    <a id="L61"></a>if r.Max.X &lt; r.Min.X {
        <a id="L62"></a>r.Max.X = r.Min.X
    <a id="L63"></a>}
    <a id="L64"></a>if r.Max.Y &lt; r.Min.Y {
        <a id="L65"></a>r.Max.Y = r.Min.Y
    <a id="L66"></a>}
    <a id="L67"></a>return r;
<a id="L68"></a>}

<a id="L70"></a><span class="comment">// Overlaps returns true if r and r1 cross; that is, it returns true if they share any point.</span>
<a id="L71"></a>func (r Rectangle) Overlaps(r1 Rectangle) bool {
    <a id="L72"></a>return r.Min.X &lt; r1.Max.X &amp;&amp; r1.Min.X &lt; r.Max.X &amp;&amp;
        <a id="L73"></a>r.Min.Y &lt; r1.Max.Y &amp;&amp; r1.Min.Y &lt; r.Max.Y
<a id="L74"></a>}

<a id="L76"></a><span class="comment">// Empty retruns true if r contains no points.</span>
<a id="L77"></a>func (r Rectangle) Empty() bool { return r.Max.X &lt;= r.Min.X || r.Max.Y &lt;= r.Min.Y }

<a id="L79"></a><span class="comment">// InRect returns true if all the points in r are also in r1.</span>
<a id="L80"></a>func (r Rectangle) In(r1 Rectangle) bool {
    <a id="L81"></a>if r.Empty() {
        <a id="L82"></a>return true
    <a id="L83"></a>}
    <a id="L84"></a>if r1.Empty() {
        <a id="L85"></a>return false
    <a id="L86"></a>}
    <a id="L87"></a>return r1.Min.X &lt;= r.Min.X &amp;&amp; r.Max.X &lt;= r1.Max.X &amp;&amp;
        <a id="L88"></a>r1.Min.Y &lt;= r.Min.Y &amp;&amp; r.Max.Y &lt;= r1.Max.Y;
<a id="L89"></a>}

<a id="L91"></a><span class="comment">// Combine returns the smallest rectangle containing all points from r and from r1.</span>
<a id="L92"></a>func (r Rectangle) Combine(r1 Rectangle) Rectangle {
    <a id="L93"></a>if r.Empty() {
        <a id="L94"></a>return r1
    <a id="L95"></a>}
    <a id="L96"></a>if r1.Empty() {
        <a id="L97"></a>return r
    <a id="L98"></a>}
    <a id="L99"></a>if r.Min.X &gt; r1.Min.X {
        <a id="L100"></a>r.Min.X = r1.Min.X
    <a id="L101"></a>}
    <a id="L102"></a>if r.Min.Y &gt; r1.Min.Y {
        <a id="L103"></a>r.Min.Y = r1.Min.Y
    <a id="L104"></a>}
    <a id="L105"></a>if r.Max.X &lt; r1.Max.X {
        <a id="L106"></a>r.Max.X = r1.Max.X
    <a id="L107"></a>}
    <a id="L108"></a>if r.Max.Y &lt; r1.Max.Y {
        <a id="L109"></a>r.Max.Y = r1.Max.Y
    <a id="L110"></a>}
    <a id="L111"></a>return r;
<a id="L112"></a>}

<a id="L114"></a><span class="comment">// Clip returns the largest rectangle containing only points shared by r and r1.</span>
<a id="L115"></a>func (r Rectangle) Clip(r1 Rectangle) Rectangle {
    <a id="L116"></a>if r.Empty() {
        <a id="L117"></a>return r
    <a id="L118"></a>}
    <a id="L119"></a>if r1.Empty() {
        <a id="L120"></a>return r1
    <a id="L121"></a>}
    <a id="L122"></a>if r.Min.X &lt; r1.Min.X {
        <a id="L123"></a>r.Min.X = r1.Min.X
    <a id="L124"></a>}
    <a id="L125"></a>if r.Min.Y &lt; r1.Min.Y {
        <a id="L126"></a>r.Min.Y = r1.Min.Y
    <a id="L127"></a>}
    <a id="L128"></a>if r.Max.X &gt; r1.Max.X {
        <a id="L129"></a>r.Max.X = r1.Max.X
    <a id="L130"></a>}
    <a id="L131"></a>if r.Max.Y &gt; r1.Max.Y {
        <a id="L132"></a>r.Max.Y = r1.Max.Y
    <a id="L133"></a>}
    <a id="L134"></a>return r;
<a id="L135"></a>}

<a id="L137"></a><span class="comment">// Dx returns the width of the rectangle r: r.Max.X - r.Min.X.</span>
<a id="L138"></a>func (r Rectangle) Dx() int { return r.Max.X - r.Min.X }

<a id="L140"></a><span class="comment">// Dy returns the width of the rectangle r: r.Max.Y - r.Min.Y.</span>
<a id="L141"></a>func (r Rectangle) Dy() int { return r.Max.Y - r.Min.Y }
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
