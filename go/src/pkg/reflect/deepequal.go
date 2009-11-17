<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/reflect/deepequal.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/reflect/deepequal.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Deep equality test via reflection</span>

<a id="L7"></a>package reflect


<a id="L10"></a><span class="comment">// During deepValueEqual, must keep track of checks that are</span>
<a id="L11"></a><span class="comment">// in progress.  The comparison algorithm assumes that all</span>
<a id="L12"></a><span class="comment">// checks in progress are true when it reencounters them.</span>
<a id="L13"></a><span class="comment">// Visited are stored in a map indexed by 17 * a1 + a2;</span>
<a id="L14"></a>type visit struct {
    <a id="L15"></a>a1   uintptr;
    <a id="L16"></a>a2   uintptr;
    <a id="L17"></a>typ  Type;
    <a id="L18"></a>next *visit;
<a id="L19"></a>}

<a id="L21"></a><span class="comment">// Tests for deep equality using reflected types. The map argument tracks</span>
<a id="L22"></a><span class="comment">// comparisons that have already been seen, which allows short circuiting on</span>
<a id="L23"></a><span class="comment">// recursive types.</span>
<a id="L24"></a>func deepValueEqual(v1, v2 Value, visited map[uintptr]*visit, depth int) bool {
    <a id="L25"></a>if v1 == nil || v2 == nil {
        <a id="L26"></a>return v1 == v2
    <a id="L27"></a>}
    <a id="L28"></a>if v1.Type() != v2.Type() {
        <a id="L29"></a>return false
    <a id="L30"></a>}

    <a id="L32"></a><span class="comment">// if depth &gt; 10 { panic(&#34;deepValueEqual&#34;) }	// for debugging</span>

    <a id="L34"></a>addr1 := v1.Addr();
    <a id="L35"></a>addr2 := v2.Addr();
    <a id="L36"></a>if addr1 &gt; addr2 {
        <a id="L37"></a><span class="comment">// Canonicalize order to reduce number of entries in visited.</span>
        <a id="L38"></a>addr1, addr2 = addr2, addr1
    <a id="L39"></a>}

    <a id="L41"></a><span class="comment">// Short circuit if references are identical ...</span>
    <a id="L42"></a>if addr1 == addr2 {
        <a id="L43"></a>return true
    <a id="L44"></a>}

    <a id="L46"></a><span class="comment">// ... or already seen</span>
    <a id="L47"></a>h := 17*addr1 + addr2;
    <a id="L48"></a>seen, _ := visited[h];
    <a id="L49"></a>typ := v1.Type();
    <a id="L50"></a>for p := seen; p != nil; p = p.next {
        <a id="L51"></a>if p.a1 == addr1 &amp;&amp; p.a2 == addr2 &amp;&amp; p.typ == typ {
            <a id="L52"></a>return true
        <a id="L53"></a>}
    <a id="L54"></a>}

    <a id="L56"></a><span class="comment">// Remember for later.</span>
    <a id="L57"></a>visited[h] = &amp;visit{addr1, addr2, typ, seen};

    <a id="L59"></a>switch v := v1.(type) {
    <a id="L60"></a>case *ArrayValue:
        <a id="L61"></a>arr1 := v;
        <a id="L62"></a>arr2 := v2.(*ArrayValue);
        <a id="L63"></a>if arr1.Len() != arr2.Len() {
            <a id="L64"></a>return false
        <a id="L65"></a>}
        <a id="L66"></a>for i := 0; i &lt; arr1.Len(); i++ {
            <a id="L67"></a>if !deepValueEqual(arr1.Elem(i), arr2.Elem(i), visited, depth+1) {
                <a id="L68"></a>return false
            <a id="L69"></a>}
        <a id="L70"></a>}
        <a id="L71"></a>return true;
    <a id="L72"></a>case *SliceValue:
        <a id="L73"></a>arr1 := v;
        <a id="L74"></a>arr2 := v2.(*SliceValue);
        <a id="L75"></a>if arr1.Len() != arr2.Len() {
            <a id="L76"></a>return false
        <a id="L77"></a>}
        <a id="L78"></a>for i := 0; i &lt; arr1.Len(); i++ {
            <a id="L79"></a>if !deepValueEqual(arr1.Elem(i), arr2.Elem(i), visited, depth+1) {
                <a id="L80"></a>return false
            <a id="L81"></a>}
        <a id="L82"></a>}
        <a id="L83"></a>return true;
    <a id="L84"></a>case *InterfaceValue:
        <a id="L85"></a>i1 := v.Interface();
        <a id="L86"></a>i2 := v2.Interface();
        <a id="L87"></a>if i1 == nil || i2 == nil {
            <a id="L88"></a>return i1 == i2
        <a id="L89"></a>}
        <a id="L90"></a>return deepValueEqual(NewValue(i1), NewValue(i2), visited, depth+1);
    <a id="L91"></a>case *PtrValue:
        <a id="L92"></a>return deepValueEqual(v.Elem(), v2.(*PtrValue).Elem(), visited, depth+1)
    <a id="L93"></a>case *StructValue:
        <a id="L94"></a>struct1 := v;
        <a id="L95"></a>struct2 := v2.(*StructValue);
        <a id="L96"></a>for i, n := 0, v.NumField(); i &lt; n; i++ {
            <a id="L97"></a>if !deepValueEqual(struct1.Field(i), struct2.Field(i), visited, depth+1) {
                <a id="L98"></a>return false
            <a id="L99"></a>}
        <a id="L100"></a>}
        <a id="L101"></a>return true;
    <a id="L102"></a>case *MapValue:
        <a id="L103"></a>map1 := v;
        <a id="L104"></a>map2 := v2.(*MapValue);
        <a id="L105"></a>if map1.Len() != map2.Len() {
            <a id="L106"></a>return false
        <a id="L107"></a>}
        <a id="L108"></a>for _, k := range map1.Keys() {
            <a id="L109"></a>if !deepValueEqual(map1.Elem(k), map2.Elem(k), visited, depth+1) {
                <a id="L110"></a>return false
            <a id="L111"></a>}
        <a id="L112"></a>}
        <a id="L113"></a>return true;
    <a id="L114"></a>default:
        <a id="L115"></a><span class="comment">// Normal equality suffices</span>
        <a id="L116"></a>return v1.Interface() == v2.Interface()
    <a id="L117"></a>}

    <a id="L119"></a>panic(&#34;Not reached&#34;);
<a id="L120"></a>}

<a id="L122"></a><span class="comment">// DeepEqual tests for deep equality. It uses normal == equality where possible</span>
<a id="L123"></a><span class="comment">// but will scan members of arrays, slices, and fields of structs. It correctly</span>
<a id="L124"></a><span class="comment">// handles recursive types.</span>
<a id="L125"></a>func DeepEqual(a1, a2 interface{}) bool {
    <a id="L126"></a>if a1 == nil || a2 == nil {
        <a id="L127"></a>return a1 == a2
    <a id="L128"></a>}
    <a id="L129"></a>v1 := NewValue(a1);
    <a id="L130"></a>v2 := NewValue(a2);
    <a id="L131"></a>if v1.Type() != v2.Type() {
        <a id="L132"></a>return false
    <a id="L133"></a>}
    <a id="L134"></a>return deepValueEqual(v1, v2, make(map[uintptr]*visit), 0);
<a id="L135"></a>}
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
