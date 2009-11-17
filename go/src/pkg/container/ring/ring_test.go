<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/ring/ring_test.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/container/ring/ring_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ring

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)


<a id="L13"></a><span class="comment">// For debugging - keep around.</span>
<a id="L14"></a>func dump(r *Ring) {
    <a id="L15"></a>if r == nil {
        <a id="L16"></a>fmt.Println(&#34;empty&#34;);
        <a id="L17"></a>return;
    <a id="L18"></a>}
    <a id="L19"></a>i, n := 0, r.Len();
    <a id="L20"></a>for p := r; i &lt; n; p = p.next {
        <a id="L21"></a>fmt.Printf(&#34;%4d: %p = {&lt;- %p | %p -&gt;}\n&#34;, i, p, p.prev, p.next);
        <a id="L22"></a>i++;
    <a id="L23"></a>}
    <a id="L24"></a>fmt.Println();
<a id="L25"></a>}


<a id="L28"></a>func verify(t *testing.T, r *Ring, N int, sum int) {
    <a id="L29"></a><span class="comment">// Len</span>
    <a id="L30"></a>n := r.Len();
    <a id="L31"></a>if n != N {
        <a id="L32"></a>t.Errorf(&#34;r.Len() == %d; expected %d&#34;, n, N)
    <a id="L33"></a>}

    <a id="L35"></a><span class="comment">// iteration</span>
    <a id="L36"></a>n = 0;
    <a id="L37"></a>s := 0;
    <a id="L38"></a>for p := range r.Iter() {
        <a id="L39"></a>n++;
        <a id="L40"></a>if p != nil {
            <a id="L41"></a>s += p.(int)
        <a id="L42"></a>}
    <a id="L43"></a>}
    <a id="L44"></a>if n != N {
        <a id="L45"></a>t.Errorf(&#34;number of forward iterations == %d; expected %d&#34;, n, N)
    <a id="L46"></a>}
    <a id="L47"></a>if sum &gt;= 0 &amp;&amp; s != sum {
        <a id="L48"></a>t.Errorf(&#34;forward ring sum = %d; expected %d&#34;, s, sum)
    <a id="L49"></a>}

    <a id="L51"></a>if r == nil {
        <a id="L52"></a>return
    <a id="L53"></a>}

    <a id="L55"></a><span class="comment">// connections</span>
    <a id="L56"></a>if r.next != nil {
        <a id="L57"></a>var p *Ring; <span class="comment">// previous element</span>
        <a id="L58"></a>for q := r; p == nil || q != r; q = q.next {
            <a id="L59"></a>if p != nil &amp;&amp; p != q.prev {
                <a id="L60"></a>t.Errorf(&#34;prev = %p, expected q.prev = %p\n&#34;, p, q.prev)
            <a id="L61"></a>}
            <a id="L62"></a>p = q;
        <a id="L63"></a>}
        <a id="L64"></a>if p != r.prev {
            <a id="L65"></a>t.Errorf(&#34;prev = %p, expected r.prev = %p\n&#34;, p, r.prev)
        <a id="L66"></a>}
    <a id="L67"></a>}

    <a id="L69"></a><span class="comment">// Next, Prev</span>
    <a id="L70"></a>if r.Next() != r.next {
        <a id="L71"></a>t.Errorf(&#34;r.Next() != r.next&#34;)
    <a id="L72"></a>}
    <a id="L73"></a>if r.Prev() != r.prev {
        <a id="L74"></a>t.Errorf(&#34;r.Prev() != r.prev&#34;)
    <a id="L75"></a>}

    <a id="L77"></a><span class="comment">// Move</span>
    <a id="L78"></a>if r.Move(0) != r {
        <a id="L79"></a>t.Errorf(&#34;r.Move(0) != r&#34;)
    <a id="L80"></a>}
    <a id="L81"></a>if r.Move(N) != r {
        <a id="L82"></a>t.Errorf(&#34;r.Move(%d) != r&#34;, N)
    <a id="L83"></a>}
    <a id="L84"></a>if r.Move(-N) != r {
        <a id="L85"></a>t.Errorf(&#34;r.Move(%d) != r&#34;, -N)
    <a id="L86"></a>}
    <a id="L87"></a>for i := 0; i &lt; 10; i++ {
        <a id="L88"></a>ni := N + i;
        <a id="L89"></a>mi := ni % N;
        <a id="L90"></a>if r.Move(ni) != r.Move(mi) {
            <a id="L91"></a>t.Errorf(&#34;r.Move(%d) != r.Move(%d)&#34;, ni, mi)
        <a id="L92"></a>}
        <a id="L93"></a>if r.Move(-ni) != r.Move(-mi) {
            <a id="L94"></a>t.Errorf(&#34;r.Move(%d) != r.Move(%d)&#34;, -ni, -mi)
        <a id="L95"></a>}
    <a id="L96"></a>}
<a id="L97"></a>}


<a id="L100"></a>func TestCornerCases(t *testing.T) {
    <a id="L101"></a>var (
        <a id="L102"></a>r0  *Ring;
        <a id="L103"></a>r1  Ring;
    <a id="L104"></a>)
    <a id="L105"></a><span class="comment">// Basics</span>
    <a id="L106"></a>verify(t, r0, 0, 0);
    <a id="L107"></a>verify(t, &amp;r1, 1, 0);
    <a id="L108"></a><span class="comment">// Insert</span>
    <a id="L109"></a>r1.Link(r0);
    <a id="L110"></a>verify(t, r0, 0, 0);
    <a id="L111"></a>verify(t, &amp;r1, 1, 0);
    <a id="L112"></a><span class="comment">// Insert</span>
    <a id="L113"></a>r1.Link(r0);
    <a id="L114"></a>verify(t, r0, 0, 0);
    <a id="L115"></a>verify(t, &amp;r1, 1, 0);
    <a id="L116"></a><span class="comment">// Unlink</span>
    <a id="L117"></a>r1.Unlink(0);
    <a id="L118"></a>verify(t, &amp;r1, 1, 0);
<a id="L119"></a>}


<a id="L122"></a>func makeN(n int) *Ring {
    <a id="L123"></a>r := New(n);
    <a id="L124"></a>for i := 1; i &lt;= n; i++ {
        <a id="L125"></a>r.Value = i;
        <a id="L126"></a>r = r.Next();
    <a id="L127"></a>}
    <a id="L128"></a>return r;
<a id="L129"></a>}


<a id="L132"></a>func sum(r *Ring) int {
    <a id="L133"></a>s := 0;
    <a id="L134"></a>for p := range r.Iter() {
        <a id="L135"></a>s += p.(int)
    <a id="L136"></a>}
    <a id="L137"></a>return s;
<a id="L138"></a>}


<a id="L141"></a>func sumN(n int) int { return (n*n + n) / 2 }


<a id="L144"></a>func TestNew(t *testing.T) {
    <a id="L145"></a>for i := 0; i &lt; 10; i++ {
        <a id="L146"></a>r := New(i);
        <a id="L147"></a>verify(t, r, i, -1);
    <a id="L148"></a>}
    <a id="L149"></a>for i := 0; i &lt; 10; i++ {
        <a id="L150"></a>r := makeN(i);
        <a id="L151"></a>verify(t, r, i, sumN(i));
    <a id="L152"></a>}
<a id="L153"></a>}


<a id="L156"></a>func TestLink1(t *testing.T) {
    <a id="L157"></a>r1a := makeN(1);
    <a id="L158"></a>var r1b Ring;
    <a id="L159"></a>r2a := r1a.Link(&amp;r1b);
    <a id="L160"></a>verify(t, r2a, 2, 1);
    <a id="L161"></a>if r2a != r1a {
        <a id="L162"></a>t.Errorf(&#34;a) 2-element link failed&#34;)
    <a id="L163"></a>}

    <a id="L165"></a>r2b := r2a.Link(r2a.Next());
    <a id="L166"></a>verify(t, r2b, 2, 1);
    <a id="L167"></a>if r2b != r2a.Next() {
        <a id="L168"></a>t.Errorf(&#34;b) 2-element link failed&#34;)
    <a id="L169"></a>}

    <a id="L171"></a>r1c := r2b.Link(r2b);
    <a id="L172"></a>verify(t, r1c, 1, 1);
    <a id="L173"></a>verify(t, r2b, 1, 0);
<a id="L174"></a>}


<a id="L177"></a>func TestLink2(t *testing.T) {
    <a id="L178"></a>var r0 *Ring;
    <a id="L179"></a>r1a := &amp;Ring{Value: 42};
    <a id="L180"></a>r1b := &amp;Ring{Value: 77};
    <a id="L181"></a>r10 := makeN(10);

    <a id="L183"></a>r1a.Link(r0);
    <a id="L184"></a>verify(t, r1a, 1, 42);

    <a id="L186"></a>r1a.Link(r1b);
    <a id="L187"></a>verify(t, r1a, 2, 42+77);

    <a id="L189"></a>r10.Link(r0);
    <a id="L190"></a>verify(t, r10, 10, sumN(10));

    <a id="L192"></a>r10.Link(r1a);
    <a id="L193"></a>verify(t, r10, 12, sumN(10)+42+77);
<a id="L194"></a>}


<a id="L197"></a>func TestLink3(t *testing.T) {
    <a id="L198"></a>var r Ring;
    <a id="L199"></a>n := 1;
    <a id="L200"></a>for i := 1; i &lt; 100; i++ {
        <a id="L201"></a>n += i;
        <a id="L202"></a>verify(t, r.Link(New(i)), n, -1);
    <a id="L203"></a>}
<a id="L204"></a>}


<a id="L207"></a>func TestUnlink(t *testing.T) {
    <a id="L208"></a>r10 := makeN(10);
    <a id="L209"></a>s10 := r10.Move(6);

    <a id="L211"></a>sum10 := sumN(10);

    <a id="L213"></a>verify(t, r10, 10, sum10);
    <a id="L214"></a>verify(t, s10, 10, sum10);

    <a id="L216"></a>r0 := r10.Unlink(0);
    <a id="L217"></a>verify(t, r0, 0, 0);

    <a id="L219"></a>r1 := r10.Unlink(1);
    <a id="L220"></a>verify(t, r1, 1, 2);
    <a id="L221"></a>verify(t, r10, 9, sum10-2);

    <a id="L223"></a>r9 := r10.Unlink(9);
    <a id="L224"></a>verify(t, r9, 9, sum10-2);
    <a id="L225"></a>verify(t, r10, 9, sum10-2);
<a id="L226"></a>}


<a id="L229"></a>func TestLinkUnlink(t *testing.T) {
    <a id="L230"></a>for i := 1; i &lt; 4; i++ {
        <a id="L231"></a>ri := New(i);
        <a id="L232"></a>for j := 0; j &lt; i; j++ {
            <a id="L233"></a>rj := ri.Unlink(j);
            <a id="L234"></a>verify(t, rj, j, -1);
            <a id="L235"></a>verify(t, ri, i-j, -1);
            <a id="L236"></a>ri.Link(rj);
            <a id="L237"></a>verify(t, ri, i, -1);
        <a id="L238"></a>}
    <a id="L239"></a>}
<a id="L240"></a>}
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
