<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/sort/sort_test.go</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/sort/sort_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package sort

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;rand&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)


<a id="L14"></a>var ints = [...]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
<a id="L15"></a>var floats = [...]float{74.3, 59.0, 238.2, -784.0, 2.3, 9845.768, -959.7485, 905, 7.8, 7.8}
<a id="L16"></a>var strings = [...]string{&#34;&#34;, &#34;Hello&#34;, &#34;foo&#34;, &#34;bar&#34;, &#34;foo&#34;, &#34;f00&#34;, &#34;%*&amp;^*&amp;^&amp;&#34;, &#34;***&#34;}

<a id="L18"></a>func TestSortIntArray(t *testing.T) {
    <a id="L19"></a>data := ints;
    <a id="L20"></a>a := IntArray(&amp;data);
    <a id="L21"></a>Sort(a);
    <a id="L22"></a>if !IsSorted(a) {
        <a id="L23"></a>t.Errorf(&#34;sorted %v&#34;, ints);
        <a id="L24"></a>t.Errorf(&#34;   got %v&#34;, data);
    <a id="L25"></a>}
<a id="L26"></a>}

<a id="L28"></a>func TestSortFloatArray(t *testing.T) {
    <a id="L29"></a>data := floats;
    <a id="L30"></a>a := FloatArray(&amp;data);
    <a id="L31"></a>Sort(a);
    <a id="L32"></a>if !IsSorted(a) {
        <a id="L33"></a>t.Errorf(&#34;sorted %v&#34;, floats);
        <a id="L34"></a>t.Errorf(&#34;   got %v&#34;, data);
    <a id="L35"></a>}
<a id="L36"></a>}

<a id="L38"></a>func TestSortStringArray(t *testing.T) {
    <a id="L39"></a>data := strings;
    <a id="L40"></a>a := StringArray(&amp;data);
    <a id="L41"></a>Sort(a);
    <a id="L42"></a>if !IsSorted(a) {
        <a id="L43"></a>t.Errorf(&#34;sorted %v&#34;, strings);
        <a id="L44"></a>t.Errorf(&#34;   got %v&#34;, data);
    <a id="L45"></a>}
<a id="L46"></a>}

<a id="L48"></a>func TestSortInts(t *testing.T) {
    <a id="L49"></a>data := ints;
    <a id="L50"></a>SortInts(&amp;data);
    <a id="L51"></a>if !IntsAreSorted(&amp;data) {
        <a id="L52"></a>t.Errorf(&#34;sorted %v&#34;, ints);
        <a id="L53"></a>t.Errorf(&#34;   got %v&#34;, data);
    <a id="L54"></a>}
<a id="L55"></a>}

<a id="L57"></a>func TestSortFloats(t *testing.T) {
    <a id="L58"></a>data := floats;
    <a id="L59"></a>SortFloats(&amp;data);
    <a id="L60"></a>if !FloatsAreSorted(&amp;data) {
        <a id="L61"></a>t.Errorf(&#34;sorted %v&#34;, floats);
        <a id="L62"></a>t.Errorf(&#34;   got %v&#34;, data);
    <a id="L63"></a>}
<a id="L64"></a>}

<a id="L66"></a>func TestSortStrings(t *testing.T) {
    <a id="L67"></a>data := strings;
    <a id="L68"></a>SortStrings(&amp;data);
    <a id="L69"></a>if !StringsAreSorted(&amp;data) {
        <a id="L70"></a>t.Errorf(&#34;sorted %v&#34;, strings);
        <a id="L71"></a>t.Errorf(&#34;   got %v&#34;, data);
    <a id="L72"></a>}
<a id="L73"></a>}

<a id="L75"></a>func TestSortLarge_Random(t *testing.T) {
    <a id="L76"></a>data := make([]int, 1000000);
    <a id="L77"></a>for i := 0; i &lt; len(data); i++ {
        <a id="L78"></a>data[i] = rand.Intn(100)
    <a id="L79"></a>}
    <a id="L80"></a>if IntsAreSorted(data) {
        <a id="L81"></a>t.Fatalf(&#34;terrible rand.rand&#34;)
    <a id="L82"></a>}
    <a id="L83"></a>SortInts(data);
    <a id="L84"></a>if !IntsAreSorted(data) {
        <a id="L85"></a>t.Errorf(&#34;sort didn&#39;t sort - 1M ints&#34;)
    <a id="L86"></a>}
<a id="L87"></a>}

<a id="L89"></a>const (
    <a id="L90"></a>_Sawtooth = iota;
    <a id="L91"></a>_Rand;
    <a id="L92"></a>_Stagger;
    <a id="L93"></a>_Plateau;
    <a id="L94"></a>_Shuffle;
    <a id="L95"></a>_NDist;
<a id="L96"></a>)

<a id="L98"></a>const (
    <a id="L99"></a>_Copy = iota;
    <a id="L100"></a>_Reverse;
    <a id="L101"></a>_ReverseFirstHalf;
    <a id="L102"></a>_ReverseSecondHalf;
    <a id="L103"></a>_Sorted;
    <a id="L104"></a>_Dither;
    <a id="L105"></a>_NMode;
<a id="L106"></a>)

<a id="L108"></a>type testingData struct {
    <a id="L109"></a>desc    string;
    <a id="L110"></a>t       *testing.T;
    <a id="L111"></a>data    []int;
    <a id="L112"></a>maxswap int; <span class="comment">// number of swaps allowed</span>
    <a id="L113"></a>nswap   int;
<a id="L114"></a>}

<a id="L116"></a>func (d *testingData) Len() int           { return len(d.data) }
<a id="L117"></a>func (d *testingData) Less(i, j int) bool { return d.data[i] &lt; d.data[j] }
<a id="L118"></a>func (d *testingData) Swap(i, j int) {
    <a id="L119"></a>if d.nswap &gt;= d.maxswap {
        <a id="L120"></a>d.t.Errorf(&#34;%s: used %d swaps sorting array of %d&#34;, d.desc, d.nswap, len(d.data));
        <a id="L121"></a>d.t.FailNow();
    <a id="L122"></a>}
    <a id="L123"></a>d.nswap++;
    <a id="L124"></a>d.data[i], d.data[j] = d.data[j], d.data[i];
<a id="L125"></a>}

<a id="L127"></a>func lg(n int) int {
    <a id="L128"></a>i := 0;
    <a id="L129"></a>for 1&lt;&lt;uint(i) &lt; n {
        <a id="L130"></a>i++
    <a id="L131"></a>}
    <a id="L132"></a>return i;
<a id="L133"></a>}

<a id="L135"></a>func TestBentleyMcIlroy(t *testing.T) {
    <a id="L136"></a>sizes := []int{100, 1023, 1024, 1025};
    <a id="L137"></a>dists := []string{&#34;sawtooth&#34;, &#34;rand&#34;, &#34;stagger&#34;, &#34;plateau&#34;, &#34;shuffle&#34;};
    <a id="L138"></a>modes := []string{&#34;copy&#34;, &#34;reverse&#34;, &#34;reverse1&#34;, &#34;reverse2&#34;, &#34;sort&#34;, &#34;dither&#34;};
    <a id="L139"></a>var tmp1, tmp2 [1025]int;
    <a id="L140"></a>for ni := 0; ni &lt; len(sizes); ni++ {
        <a id="L141"></a>n := sizes[ni];
        <a id="L142"></a>for m := 1; m &lt; 2*n; m *= 2 {
            <a id="L143"></a>for dist := 0; dist &lt; _NDist; dist++ {
                <a id="L144"></a>j := 0;
                <a id="L145"></a>k := 1;
                <a id="L146"></a>data := tmp1[0:n];
                <a id="L147"></a>for i := 0; i &lt; n; i++ {
                    <a id="L148"></a>switch dist {
                    <a id="L149"></a>case _Sawtooth:
                        <a id="L150"></a>data[i] = i % m
                    <a id="L151"></a>case _Rand:
                        <a id="L152"></a>data[i] = rand.Intn(m)
                    <a id="L153"></a>case _Stagger:
                        <a id="L154"></a>data[i] = (i*m + i) % n
                    <a id="L155"></a>case _Plateau:
                        <a id="L156"></a>data[i] = min(i, m)
                    <a id="L157"></a>case _Shuffle:
                        <a id="L158"></a>if rand.Intn(m) != 0 {
                            <a id="L159"></a>j += 2;
                            <a id="L160"></a>data[i] = j;
                        <a id="L161"></a>} else {
                            <a id="L162"></a>k += 2;
                            <a id="L163"></a>data[i] = k;
                        <a id="L164"></a>}
                    <a id="L165"></a>}
                <a id="L166"></a>}

                <a id="L168"></a>mdata := tmp2[0:n];
                <a id="L169"></a>for mode := 0; mode &lt; _NMode; mode++ {
                    <a id="L170"></a>switch mode {
                    <a id="L171"></a>case _Copy:
                        <a id="L172"></a>for i := 0; i &lt; n; i++ {
                            <a id="L173"></a>mdata[i] = data[i]
                        <a id="L174"></a>}
                    <a id="L175"></a>case _Reverse:
                        <a id="L176"></a>for i := 0; i &lt; n; i++ {
                            <a id="L177"></a>mdata[i] = data[n-i-1]
                        <a id="L178"></a>}
                    <a id="L179"></a>case _ReverseFirstHalf:
                        <a id="L180"></a>for i := 0; i &lt; n/2; i++ {
                            <a id="L181"></a>mdata[i] = data[n/2-i-1]
                        <a id="L182"></a>}
                        <a id="L183"></a>for i := n / 2; i &lt; n; i++ {
                            <a id="L184"></a>mdata[i] = data[i]
                        <a id="L185"></a>}
                    <a id="L186"></a>case _ReverseSecondHalf:
                        <a id="L187"></a>for i := 0; i &lt; n/2; i++ {
                            <a id="L188"></a>mdata[i] = data[i]
                        <a id="L189"></a>}
                        <a id="L190"></a>for i := n / 2; i &lt; n; i++ {
                            <a id="L191"></a>mdata[i] = data[n-(i-n/2)-1]
                        <a id="L192"></a>}
                    <a id="L193"></a>case _Sorted:
                        <a id="L194"></a>for i := 0; i &lt; n; i++ {
                            <a id="L195"></a>mdata[i] = data[i]
                        <a id="L196"></a>}
                        <a id="L197"></a><span class="comment">// SortInts is known to be correct</span>
                        <a id="L198"></a><span class="comment">// because mode Sort runs after mode _Copy.</span>
                        <a id="L199"></a>SortInts(mdata);
                    <a id="L200"></a>case _Dither:
                        <a id="L201"></a>for i := 0; i &lt; n; i++ {
                            <a id="L202"></a>mdata[i] = data[i] + i%5
                        <a id="L203"></a>}
                    <a id="L204"></a>}

                    <a id="L206"></a>desc := fmt.Sprintf(&#34;n=%d m=%d dist=%s mode=%s&#34;, n, m, dists[dist], modes[mode]);
                    <a id="L207"></a>d := &amp;testingData{desc, t, mdata[0:n], n * lg(n) * 12 / 10, 0};
                    <a id="L208"></a>Sort(d);

                    <a id="L210"></a><span class="comment">// If we were testing C qsort, we&#39;d have to make a copy</span>
                    <a id="L211"></a><span class="comment">// of the array and sort it ourselves and then compare</span>
                    <a id="L212"></a><span class="comment">// x against it, to ensure that qsort was only permuting</span>
                    <a id="L213"></a><span class="comment">// the data, not (for example) overwriting it with zeros.</span>
                    <a id="L214"></a><span class="comment">//</span>
                    <a id="L215"></a><span class="comment">// In go, we don&#39;t have to be so paranoid: since the only</span>
                    <a id="L216"></a><span class="comment">// mutating method Sort can call is TestingData.swap,</span>
                    <a id="L217"></a><span class="comment">// it suffices here just to check that the final array is sorted.</span>
                    <a id="L218"></a>if !IntsAreSorted(mdata) {
                        <a id="L219"></a>t.Errorf(&#34;%s: ints not sorted&#34;, desc);
                        <a id="L220"></a>t.Errorf(&#34;\t%v&#34;, mdata);
                        <a id="L221"></a>t.FailNow();
                    <a id="L222"></a>}
                <a id="L223"></a>}
            <a id="L224"></a>}
        <a id="L225"></a>}
    <a id="L226"></a>}
<a id="L227"></a>}
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
