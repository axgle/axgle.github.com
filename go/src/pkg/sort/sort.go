<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/sort/sort.go</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

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
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Source file /src/pkg/sort/sort.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The sort package provides primitives for sorting arrays</span>
<a id="L6"></a><span class="comment">// and user-defined collections.</span>
<a id="L7"></a>package sort

<a id="L9"></a><span class="comment">// A type, typically a collection, that satisfies sort.Interface can be</span>
<a id="L10"></a><span class="comment">// sorted by the routines in this package.  The methods require that the</span>
<a id="L11"></a><span class="comment">// elements of the collection be enumerated by an integer index.</span>
<a id="L12"></a>type Interface interface {
    <a id="L13"></a><span class="comment">// Len is the number of elements in the collection.</span>
    <a id="L14"></a>Len() int;
    <a id="L15"></a><span class="comment">// Less returns whether the element with index i is should sort</span>
    <a id="L16"></a><span class="comment">// before the element with index j.</span>
    <a id="L17"></a>Less(i, j int) bool;
    <a id="L18"></a><span class="comment">// Swap swaps the elements with indexes i and j.</span>
    <a id="L19"></a>Swap(i, j int);
<a id="L20"></a>}

<a id="L22"></a>func min(a, b int) int {
    <a id="L23"></a>if a &lt; b {
        <a id="L24"></a>return a
    <a id="L25"></a>}
    <a id="L26"></a>return b;
<a id="L27"></a>}

<a id="L29"></a><span class="comment">// Insertion sort</span>
<a id="L30"></a>func insertionSort(data Interface, a, b int) {
    <a id="L31"></a>for i := a + 1; i &lt; b; i++ {
        <a id="L32"></a>for j := i; j &gt; a &amp;&amp; data.Less(j, j-1); j-- {
            <a id="L33"></a>data.Swap(j, j-1)
        <a id="L34"></a>}
    <a id="L35"></a>}
<a id="L36"></a>}

<a id="L38"></a><span class="comment">// Quicksort, following Bentley and McIlroy,</span>
<a id="L39"></a><span class="comment">// ``Engineering a Sort Function,&#39;&#39; SP&amp;E November 1993.</span>

<a id="L41"></a><span class="comment">// Move the median of the three values data[a], data[b], data[c] into data[a].</span>
<a id="L42"></a>func medianOfThree(data Interface, a, b, c int) {
    <a id="L43"></a>m0 := b;
    <a id="L44"></a>m1 := a;
    <a id="L45"></a>m2 := c;
    <a id="L46"></a><span class="comment">// bubble sort on 3 elements</span>
    <a id="L47"></a>if data.Less(m1, m0) {
        <a id="L48"></a>data.Swap(m1, m0)
    <a id="L49"></a>}
    <a id="L50"></a>if data.Less(m2, m1) {
        <a id="L51"></a>data.Swap(m2, m1)
    <a id="L52"></a>}
    <a id="L53"></a>if data.Less(m1, m0) {
        <a id="L54"></a>data.Swap(m1, m0)
    <a id="L55"></a>}
    <a id="L56"></a><span class="comment">// now data[m0] &lt;= data[m1] &lt;= data[m2]</span>
<a id="L57"></a>}

<a id="L59"></a>func swapRange(data Interface, a, b, n int) {
    <a id="L60"></a>for i := 0; i &lt; n; i++ {
        <a id="L61"></a>data.Swap(a+i, b+i)
    <a id="L62"></a>}
<a id="L63"></a>}

<a id="L65"></a>func doPivot(data Interface, lo, hi int) (midlo, midhi int) {
    <a id="L66"></a>m := (lo + hi) / 2;
    <a id="L67"></a>if hi-lo &gt; 40 {
        <a id="L68"></a><span class="comment">// Tukey&#39;s ``Ninther,&#39;&#39; median of three medians of three.</span>
        <a id="L69"></a>s := (hi - lo) / 8;
        <a id="L70"></a>medianOfThree(data, lo, lo+s, lo+2*s);
        <a id="L71"></a>medianOfThree(data, m, m-s, m+s);
        <a id="L72"></a>medianOfThree(data, hi-1, hi-1-s, hi-1-2*s);
    <a id="L73"></a>}
    <a id="L74"></a>medianOfThree(data, lo, m, hi-1);

    <a id="L76"></a><span class="comment">// Invariants are:</span>
    <a id="L77"></a><span class="comment">//	data[lo] = pivot (set up by ChoosePivot)</span>
    <a id="L78"></a><span class="comment">//	data[lo &lt;= i &lt; a] = pivot</span>
    <a id="L79"></a><span class="comment">//	data[a &lt;= i &lt; b] &lt; pivot</span>
    <a id="L80"></a><span class="comment">//	data[b &lt;= i &lt; c] is unexamined</span>
    <a id="L81"></a><span class="comment">//	data[c &lt;= i &lt; d] &gt; pivot</span>
    <a id="L82"></a><span class="comment">//	data[d &lt;= i &lt; hi] = pivot</span>
    <a id="L83"></a><span class="comment">//</span>
    <a id="L84"></a><span class="comment">// Once b meets c, can swap the &#34;= pivot&#34; sections</span>
    <a id="L85"></a><span class="comment">// into the middle of the array.</span>
    <a id="L86"></a>pivot := lo;
    <a id="L87"></a>a, b, c, d := lo+1, lo+1, hi, hi;
    <a id="L88"></a>for b &lt; c {
        <a id="L89"></a>if data.Less(b, pivot) { <span class="comment">// data[b] &lt; pivot</span>
            <a id="L90"></a>b++;
            <a id="L91"></a>continue;
        <a id="L92"></a>}
        <a id="L93"></a>if !data.Less(pivot, b) { <span class="comment">// data[b] = pivot</span>
            <a id="L94"></a>data.Swap(a, b);
            <a id="L95"></a>a++;
            <a id="L96"></a>b++;
            <a id="L97"></a>continue;
        <a id="L98"></a>}
        <a id="L99"></a>if data.Less(pivot, c-1) { <span class="comment">// data[c-1] &gt; pivot</span>
            <a id="L100"></a>c--;
            <a id="L101"></a>continue;
        <a id="L102"></a>}
        <a id="L103"></a>if !data.Less(c-1, pivot) { <span class="comment">// data[c-1] = pivot</span>
            <a id="L104"></a>data.Swap(c-1, d-1);
            <a id="L105"></a>c--;
            <a id="L106"></a>d--;
            <a id="L107"></a>continue;
        <a id="L108"></a>}
        <a id="L109"></a><span class="comment">// data[b] &gt; pivot; data[c-1] &lt; pivot</span>
        <a id="L110"></a>data.Swap(b, c-1);
        <a id="L111"></a>b++;
        <a id="L112"></a>c--;
    <a id="L113"></a>}

    <a id="L115"></a>n := min(b-a, a-lo);
    <a id="L116"></a>swapRange(data, lo, b-n, n);

    <a id="L118"></a>n = min(hi-d, d-c);
    <a id="L119"></a>swapRange(data, c, hi-n, n);

    <a id="L121"></a>return lo + b - a, hi - (d - c);
<a id="L122"></a>}

<a id="L124"></a>func quickSort(data Interface, a, b int) {
    <a id="L125"></a>if b-a &gt; 7 {
        <a id="L126"></a>mlo, mhi := doPivot(data, a, b);
        <a id="L127"></a>quickSort(data, a, mlo);
        <a id="L128"></a>quickSort(data, mhi, b);
    <a id="L129"></a>} else if b-a &gt; 1 {
        <a id="L130"></a>insertionSort(data, a, b)
    <a id="L131"></a>}
<a id="L132"></a>}

<a id="L134"></a>func Sort(data Interface) { quickSort(data, 0, data.Len()) }


<a id="L137"></a>func IsSorted(data Interface) bool {
    <a id="L138"></a>n := data.Len();
    <a id="L139"></a>for i := n - 1; i &gt; 0; i-- {
        <a id="L140"></a>if data.Less(i, i-1) {
            <a id="L141"></a>return false
        <a id="L142"></a>}
    <a id="L143"></a>}
    <a id="L144"></a>return true;
<a id="L145"></a>}


<a id="L148"></a><span class="comment">// Convenience types for common cases</span>

<a id="L150"></a><span class="comment">// IntArray attaches the methods of Interface to []int, sorting in increasing order.</span>
<a id="L151"></a>type IntArray []int

<a id="L153"></a>func (p IntArray) Len() int           { return len(p) }
<a id="L154"></a>func (p IntArray) Less(i, j int) bool { return p[i] &lt; p[j] }
<a id="L155"></a>func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

<a id="L157"></a><span class="comment">// Sort is a convenience method.</span>
<a id="L158"></a>func (p IntArray) Sort() { Sort(p) }


<a id="L161"></a><span class="comment">// FloatArray attaches the methods of Interface to []float, sorting in increasing order.</span>
<a id="L162"></a>type FloatArray []float

<a id="L164"></a>func (p FloatArray) Len() int           { return len(p) }
<a id="L165"></a>func (p FloatArray) Less(i, j int) bool { return p[i] &lt; p[j] }
<a id="L166"></a>func (p FloatArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

<a id="L168"></a><span class="comment">// Sort is a convenience method.</span>
<a id="L169"></a>func (p FloatArray) Sort() { Sort(p) }


<a id="L172"></a><span class="comment">// StringArray attaches the methods of Interface to []string, sorting in increasing order.</span>
<a id="L173"></a>type StringArray []string

<a id="L175"></a>func (p StringArray) Len() int           { return len(p) }
<a id="L176"></a>func (p StringArray) Less(i, j int) bool { return p[i] &lt; p[j] }
<a id="L177"></a>func (p StringArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

<a id="L179"></a><span class="comment">// Sort is a convenience method.</span>
<a id="L180"></a>func (p StringArray) Sort() { Sort(p) }


<a id="L183"></a><span class="comment">// Convenience wrappers for common cases</span>

<a id="L185"></a><span class="comment">// SortInts sorts an array of ints in increasing order.</span>
<a id="L186"></a>func SortInts(a []int) { Sort(IntArray(a)) }
<a id="L187"></a><span class="comment">// SortFloats sorts an array of floats in increasing order.</span>
<a id="L188"></a>func SortFloats(a []float) { Sort(FloatArray(a)) }
<a id="L189"></a><span class="comment">// SortStrings sorts an array of strings in increasing order.</span>
<a id="L190"></a>func SortStrings(a []string) { Sort(StringArray(a)) }


<a id="L193"></a><span class="comment">// IntsAreSorted tests whether an array of ints is sorted in increasing order.</span>
<a id="L194"></a>func IntsAreSorted(a []int) bool { return IsSorted(IntArray(a)) }
<a id="L195"></a><span class="comment">// FloatsAreSorted tests whether an array of floats is sorted in increasing order.</span>
<a id="L196"></a>func FloatsAreSorted(a []float) bool { return IsSorted(FloatArray(a)) }
<a id="L197"></a><span class="comment">// StringsAreSorted tests whether an array of strings is sorted in increasing order.</span>
<a id="L198"></a>func StringsAreSorted(a []string) bool { return IsSorted(StringArray(a)) }
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
