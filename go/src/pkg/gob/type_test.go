<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/gob/type_test.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/gob/type_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package gob

<a id="L7"></a>import (
    <a id="L8"></a>&#34;reflect&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>type typeT struct {
    <a id="L13"></a>id  typeId;
    <a id="L14"></a>str string;
<a id="L15"></a>}

<a id="L17"></a>var basicTypes = []typeT{
    <a id="L18"></a>typeT{tBool, &#34;bool&#34;},
    <a id="L19"></a>typeT{tInt, &#34;int&#34;},
    <a id="L20"></a>typeT{tUint, &#34;uint&#34;},
    <a id="L21"></a>typeT{tFloat, &#34;float&#34;},
    <a id="L22"></a>typeT{tBytes, &#34;bytes&#34;},
    <a id="L23"></a>typeT{tString, &#34;string&#34;},
<a id="L24"></a>}

<a id="L26"></a>func getTypeUnlocked(name string, rt reflect.Type) gobType {
    <a id="L27"></a>typeLock.Lock();
    <a id="L28"></a>defer typeLock.Unlock();
    <a id="L29"></a>t, err := getType(name, rt);
    <a id="L30"></a>if err != nil {
        <a id="L31"></a>panicln(&#34;getTypeUnlocked:&#34;, err.String())
    <a id="L32"></a>}
    <a id="L33"></a>return t;
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// Sanity checks</span>
<a id="L37"></a>func TestBasic(t *testing.T) {
    <a id="L38"></a>for _, tt := range basicTypes {
        <a id="L39"></a>if tt.id.String() != tt.str {
            <a id="L40"></a>t.Errorf(&#34;checkType: expected %q got %s&#34;, tt.str, tt.id.String())
        <a id="L41"></a>}
        <a id="L42"></a>if tt.id == 0 {
            <a id="L43"></a>t.Errorf(&#34;id for %q is zero&#34;, tt.str)
        <a id="L44"></a>}
    <a id="L45"></a>}
<a id="L46"></a>}

<a id="L48"></a><span class="comment">// Reregister some basic types to check registration is idempotent.</span>
<a id="L49"></a>func TestReregistration(t *testing.T) {
    <a id="L50"></a>newtyp := getTypeUnlocked(&#34;int&#34;, reflect.Typeof(int(0)));
    <a id="L51"></a>if newtyp != tInt.gobType() {
        <a id="L52"></a>t.Errorf(&#34;reregistration of %s got new type&#34;, newtyp.String())
    <a id="L53"></a>}
    <a id="L54"></a>newtyp = getTypeUnlocked(&#34;uint&#34;, reflect.Typeof(uint(0)));
    <a id="L55"></a>if newtyp != tUint.gobType() {
        <a id="L56"></a>t.Errorf(&#34;reregistration of %s got new type&#34;, newtyp.String())
    <a id="L57"></a>}
    <a id="L58"></a>newtyp = getTypeUnlocked(&#34;string&#34;, reflect.Typeof(&#34;hello&#34;));
    <a id="L59"></a>if newtyp != tString.gobType() {
        <a id="L60"></a>t.Errorf(&#34;reregistration of %s got new type&#34;, newtyp.String())
    <a id="L61"></a>}
<a id="L62"></a>}

<a id="L64"></a>func TestArrayType(t *testing.T) {
    <a id="L65"></a>var a3 [3]int;
    <a id="L66"></a>a3int := getTypeUnlocked(&#34;foo&#34;, reflect.Typeof(a3));
    <a id="L67"></a>newa3int := getTypeUnlocked(&#34;bar&#34;, reflect.Typeof(a3));
    <a id="L68"></a>if a3int != newa3int {
        <a id="L69"></a>t.Errorf(&#34;second registration of [3]int creates new type&#34;)
    <a id="L70"></a>}
    <a id="L71"></a>var a4 [4]int;
    <a id="L72"></a>a4int := getTypeUnlocked(&#34;goo&#34;, reflect.Typeof(a4));
    <a id="L73"></a>if a3int == a4int {
        <a id="L74"></a>t.Errorf(&#34;registration of [3]int creates same type as [4]int&#34;)
    <a id="L75"></a>}
    <a id="L76"></a>var b3 [3]bool;
    <a id="L77"></a>a3bool := getTypeUnlocked(&#34;&#34;, reflect.Typeof(b3));
    <a id="L78"></a>if a3int == a3bool {
        <a id="L79"></a>t.Errorf(&#34;registration of [3]bool creates same type as [3]int&#34;)
    <a id="L80"></a>}
    <a id="L81"></a>str := a3bool.String();
    <a id="L82"></a>expected := &#34;[3]bool&#34;;
    <a id="L83"></a>if str != expected {
        <a id="L84"></a>t.Errorf(&#34;array printed as %q; expected %q&#34;, str, expected)
    <a id="L85"></a>}
<a id="L86"></a>}

<a id="L88"></a>func TestSliceType(t *testing.T) {
    <a id="L89"></a>var s []int;
    <a id="L90"></a>sint := getTypeUnlocked(&#34;slice&#34;, reflect.Typeof(s));
    <a id="L91"></a>var news []int;
    <a id="L92"></a>newsint := getTypeUnlocked(&#34;slice1&#34;, reflect.Typeof(news));
    <a id="L93"></a>if sint != newsint {
        <a id="L94"></a>t.Errorf(&#34;second registration of []int creates new type&#34;)
    <a id="L95"></a>}
    <a id="L96"></a>var b []bool;
    <a id="L97"></a>sbool := getTypeUnlocked(&#34;&#34;, reflect.Typeof(b));
    <a id="L98"></a>if sbool == sint {
        <a id="L99"></a>t.Errorf(&#34;registration of []bool creates same type as []int&#34;)
    <a id="L100"></a>}
    <a id="L101"></a>str := sbool.String();
    <a id="L102"></a>expected := &#34;[]bool&#34;;
    <a id="L103"></a>if str != expected {
        <a id="L104"></a>t.Errorf(&#34;slice printed as %q; expected %q&#34;, str, expected)
    <a id="L105"></a>}
<a id="L106"></a>}

<a id="L108"></a>type Bar struct {
    <a id="L109"></a>x string;
<a id="L110"></a>}

<a id="L112"></a><span class="comment">// This structure has pointers and refers to itself, making it a good test case.</span>
<a id="L113"></a>type Foo struct {
    <a id="L114"></a>a   int;
    <a id="L115"></a>b   int32; <span class="comment">// will become int</span>
    <a id="L116"></a>c   string;
    <a id="L117"></a>d   []byte;
    <a id="L118"></a>e   *float;      <span class="comment">// will become float</span>
    <a id="L119"></a>f   ****float64; <span class="comment">// will become float</span>
    <a id="L120"></a>g   *Bar;
    <a id="L121"></a>h   *Bar; <span class="comment">// should not interpolate the definition of Bar again</span>
    <a id="L122"></a>i   *Foo; <span class="comment">// will not explode</span>
<a id="L123"></a>}

<a id="L125"></a>func TestStructType(t *testing.T) {
    <a id="L126"></a>sstruct := getTypeUnlocked(&#34;Foo&#34;, reflect.Typeof(Foo{}));
    <a id="L127"></a>str := sstruct.String();
    <a id="L128"></a><span class="comment">// If we can print it correctly, we built it correctly.</span>
    <a id="L129"></a>expected := &#34;Foo = struct { a int; b int; c string; d bytes; e float; f float; g Bar = struct { x string; }; h Bar; i Foo; }&#34;;
    <a id="L130"></a>if str != expected {
        <a id="L131"></a>t.Errorf(&#34;struct printed as %q; expected %q&#34;, str, expected)
    <a id="L132"></a>}
<a id="L133"></a>}
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
