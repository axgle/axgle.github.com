<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/container/vector/vector_test.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/container/vector/vector_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package vector

<a id="L7"></a>import &#34;testing&#34;
<a id="L8"></a>import &#34;sort&#34;
<a id="L9"></a>import &#34;fmt&#34;


<a id="L12"></a>func TestZeroLen(t *testing.T) {
    <a id="L13"></a>var a *Vector;
    <a id="L14"></a>if a.Len() != 0 {
        <a id="L15"></a>t.Errorf(&#34;A) expected 0, got %d&#34;, a.Len())
    <a id="L16"></a>}
    <a id="L17"></a>a = New(0);
    <a id="L18"></a>if a.Len() != 0 {
        <a id="L19"></a>t.Errorf(&#34;B) expected 0, got %d&#34;, a.Len())
    <a id="L20"></a>}
<a id="L21"></a>}


<a id="L24"></a>func TestInit(t *testing.T) {
    <a id="L25"></a>var a Vector;
    <a id="L26"></a>if a.Init(0).Len() != 0 {
        <a id="L27"></a>t.Error(&#34;A&#34;)
    <a id="L28"></a>}
    <a id="L29"></a>if a.Init(1).Len() != 1 {
        <a id="L30"></a>t.Error(&#34;B&#34;)
    <a id="L31"></a>}
    <a id="L32"></a>if a.Init(10).Len() != 10 {
        <a id="L33"></a>t.Error(&#34;C&#34;)
    <a id="L34"></a>}
<a id="L35"></a>}


<a id="L38"></a>func TestNew(t *testing.T) {
    <a id="L39"></a>if New(0).Len() != 0 {
        <a id="L40"></a>t.Error(&#34;A&#34;)
    <a id="L41"></a>}
    <a id="L42"></a>if New(1).Len() != 1 {
        <a id="L43"></a>t.Error(&#34;B&#34;)
    <a id="L44"></a>}
    <a id="L45"></a>if New(10).Len() != 10 {
        <a id="L46"></a>t.Error(&#34;C&#34;)
    <a id="L47"></a>}
<a id="L48"></a>}


<a id="L51"></a>func val(i int) int { return i*991 - 1234 }


<a id="L54"></a>func TestAccess(t *testing.T) {
    <a id="L55"></a>const n = 100;
    <a id="L56"></a>var a Vector;
    <a id="L57"></a>a.Init(n);
    <a id="L58"></a>for i := 0; i &lt; n; i++ {
        <a id="L59"></a>a.Set(i, val(i))
    <a id="L60"></a>}
    <a id="L61"></a>for i := 0; i &lt; n; i++ {
        <a id="L62"></a>if a.At(i).(int) != val(i) {
            <a id="L63"></a>t.Error(i)
        <a id="L64"></a>}
    <a id="L65"></a>}
<a id="L66"></a>}


<a id="L69"></a>func TestInsertDeleteClear(t *testing.T) {
    <a id="L70"></a>const n = 100;
    <a id="L71"></a>var a Vector;

    <a id="L73"></a>for i := 0; i &lt; n; i++ {
        <a id="L74"></a>if a.Len() != i {
            <a id="L75"></a>t.Errorf(&#34;A) wrong len %d (expected %d)&#34;, a.Len(), i)
        <a id="L76"></a>}
        <a id="L77"></a>a.Insert(0, val(i));
        <a id="L78"></a>if a.Last().(int) != val(0) {
            <a id="L79"></a>t.Error(&#34;B&#34;)
        <a id="L80"></a>}
    <a id="L81"></a>}
    <a id="L82"></a>for i := n - 1; i &gt;= 0; i-- {
        <a id="L83"></a>if a.Last().(int) != val(0) {
            <a id="L84"></a>t.Error(&#34;C&#34;)
        <a id="L85"></a>}
        <a id="L86"></a>if a.At(0).(int) != val(i) {
            <a id="L87"></a>t.Error(&#34;D&#34;)
        <a id="L88"></a>}
        <a id="L89"></a>a.Delete(0);
        <a id="L90"></a>if a.Len() != i {
            <a id="L91"></a>t.Errorf(&#34;E) wrong len %d (expected %d)&#34;, a.Len(), i)
        <a id="L92"></a>}
    <a id="L93"></a>}

    <a id="L95"></a>if a.Len() != 0 {
        <a id="L96"></a>t.Errorf(&#34;F) wrong len %d (expected 0)&#34;, a.Len())
    <a id="L97"></a>}
    <a id="L98"></a>for i := 0; i &lt; n; i++ {
        <a id="L99"></a>a.Push(val(i));
        <a id="L100"></a>if a.Len() != i+1 {
            <a id="L101"></a>t.Errorf(&#34;G) wrong len %d (expected %d)&#34;, a.Len(), i+1)
        <a id="L102"></a>}
        <a id="L103"></a>if a.Last().(int) != val(i) {
            <a id="L104"></a>t.Error(&#34;H&#34;)
        <a id="L105"></a>}
    <a id="L106"></a>}
    <a id="L107"></a>a.Init(0);
    <a id="L108"></a>if a.Len() != 0 {
        <a id="L109"></a>t.Errorf(&#34;I wrong len %d (expected 0)&#34;, a.Len())
    <a id="L110"></a>}

    <a id="L112"></a>const m = 5;
    <a id="L113"></a>for j := 0; j &lt; m; j++ {
        <a id="L114"></a>a.Push(j);
        <a id="L115"></a>for i := 0; i &lt; n; i++ {
            <a id="L116"></a>x := val(i);
            <a id="L117"></a>a.Push(x);
            <a id="L118"></a>if a.Pop().(int) != x {
                <a id="L119"></a>t.Error(&#34;J&#34;)
            <a id="L120"></a>}
            <a id="L121"></a>if a.Len() != j+1 {
                <a id="L122"></a>t.Errorf(&#34;K) wrong len %d (expected %d)&#34;, a.Len(), j+1)
            <a id="L123"></a>}
        <a id="L124"></a>}
    <a id="L125"></a>}
    <a id="L126"></a>if a.Len() != m {
        <a id="L127"></a>t.Errorf(&#34;L) wrong len %d (expected %d)&#34;, a.Len(), m)
    <a id="L128"></a>}
<a id="L129"></a>}


<a id="L132"></a>func verify_slice(t *testing.T, x *Vector, elt, i, j int) {
    <a id="L133"></a>for k := i; k &lt; j; k++ {
        <a id="L134"></a>if x.At(k).(int) != elt {
            <a id="L135"></a>t.Errorf(&#34;M) wrong [%d] element %d (expected %d)&#34;, k, x.At(k).(int), elt)
        <a id="L136"></a>}
    <a id="L137"></a>}

    <a id="L139"></a>s := x.Slice(i, j);
    <a id="L140"></a>for k, n := 0, j-i; k &lt; n; k++ {
        <a id="L141"></a>if s.At(k).(int) != elt {
            <a id="L142"></a>t.Errorf(&#34;N) wrong [%d] element %d (expected %d)&#34;, k, x.At(k).(int), elt)
        <a id="L143"></a>}
    <a id="L144"></a>}
<a id="L145"></a>}


<a id="L148"></a>func verify_pattern(t *testing.T, x *Vector, a, b, c int) {
    <a id="L149"></a>n := a + b + c;
    <a id="L150"></a>if x.Len() != n {
        <a id="L151"></a>t.Errorf(&#34;O) wrong len %d (expected %d)&#34;, x.Len(), n)
    <a id="L152"></a>}
    <a id="L153"></a>verify_slice(t, x, 0, 0, a);
    <a id="L154"></a>verify_slice(t, x, 1, a, a+b);
    <a id="L155"></a>verify_slice(t, x, 0, a+b, n);
<a id="L156"></a>}


<a id="L159"></a>func make_vector(elt, len int) *Vector {
    <a id="L160"></a>x := New(len);
    <a id="L161"></a>for i := 0; i &lt; len; i++ {
        <a id="L162"></a>x.Set(i, elt)
    <a id="L163"></a>}
    <a id="L164"></a>return x;
<a id="L165"></a>}


<a id="L168"></a>func TestInsertVector(t *testing.T) {
    <a id="L169"></a><span class="comment">// 1</span>
    <a id="L170"></a>a := make_vector(0, 0);
    <a id="L171"></a>b := make_vector(1, 10);
    <a id="L172"></a>a.InsertVector(0, b);
    <a id="L173"></a>verify_pattern(t, a, 0, 10, 0);
    <a id="L174"></a><span class="comment">// 2</span>
    <a id="L175"></a>a = make_vector(0, 10);
    <a id="L176"></a>b = make_vector(1, 0);
    <a id="L177"></a>a.InsertVector(5, b);
    <a id="L178"></a>verify_pattern(t, a, 5, 0, 5);
    <a id="L179"></a><span class="comment">// 3</span>
    <a id="L180"></a>a = make_vector(0, 10);
    <a id="L181"></a>b = make_vector(1, 3);
    <a id="L182"></a>a.InsertVector(3, b);
    <a id="L183"></a>verify_pattern(t, a, 3, 3, 7);
    <a id="L184"></a><span class="comment">// 4</span>
    <a id="L185"></a>a = make_vector(0, 10);
    <a id="L186"></a>b = make_vector(1, 1000);
    <a id="L187"></a>a.InsertVector(8, b);
    <a id="L188"></a>verify_pattern(t, a, 8, 1000, 2);
<a id="L189"></a>}


<a id="L192"></a><span class="comment">// This also tests IntVector and StringVector</span>
<a id="L193"></a>func TestSorting(t *testing.T) {
    <a id="L194"></a>const n = 100;

    <a id="L196"></a>a := NewIntVector(n);
    <a id="L197"></a>for i := n - 1; i &gt;= 0; i-- {
        <a id="L198"></a>a.Set(i, n-1-i)
    <a id="L199"></a>}
    <a id="L200"></a>if sort.IsSorted(a) {
        <a id="L201"></a>t.Error(&#34;int vector not sorted&#34;)
    <a id="L202"></a>}

    <a id="L204"></a>b := NewStringVector(n);
    <a id="L205"></a>for i := n - 1; i &gt;= 0; i-- {
        <a id="L206"></a>b.Set(i, fmt.Sprint(n-1-i))
    <a id="L207"></a>}
    <a id="L208"></a>if sort.IsSorted(b) {
        <a id="L209"></a>t.Error(&#34;string vector not sorted&#34;)
    <a id="L210"></a>}
<a id="L211"></a>}


<a id="L214"></a>func TestDo(t *testing.T) {
    <a id="L215"></a>const n = 25;
    <a id="L216"></a>const salt = 17;
    <a id="L217"></a>a := NewIntVector(n);
    <a id="L218"></a>for i := 0; i &lt; n; i++ {
        <a id="L219"></a>a.Set(i, salt*i)
    <a id="L220"></a>}
    <a id="L221"></a>count := 0;
    <a id="L222"></a>a.Do(func(e interface{}) {
        <a id="L223"></a>i := e.(int);
        <a id="L224"></a>if i != count*salt {
            <a id="L225"></a>t.Error(&#34;value at&#34;, count, &#34;should be&#34;, count*salt, &#34;not&#34;, i)
        <a id="L226"></a>}
        <a id="L227"></a>count++;
    <a id="L228"></a>});
    <a id="L229"></a>if count != n {
        <a id="L230"></a>t.Error(&#34;should visit&#34;, n, &#34;values; did visit&#34;, count)
    <a id="L231"></a>}
<a id="L232"></a>}


<a id="L235"></a>func TestIter(t *testing.T) {
    <a id="L236"></a>const Len = 100;
    <a id="L237"></a>x := New(Len);
    <a id="L238"></a>for i := 0; i &lt; Len; i++ {
        <a id="L239"></a>x.Set(i, i*i)
    <a id="L240"></a>}
    <a id="L241"></a>i := 0;
    <a id="L242"></a>for v := range x.Iter() {
        <a id="L243"></a>if v.(int) != i*i {
            <a id="L244"></a>t.Error(&#34;Iter expected&#34;, i*i, &#34;got&#34;, v.(int))
        <a id="L245"></a>}
        <a id="L246"></a>i++;
    <a id="L247"></a>}
    <a id="L248"></a>if i != Len {
        <a id="L249"></a>t.Error(&#34;Iter stopped at&#34;, i, &#34;not&#34;, Len)
    <a id="L250"></a>}
<a id="L251"></a>}
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
