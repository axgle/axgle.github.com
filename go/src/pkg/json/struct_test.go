<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/json/struct_test.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/json/struct_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package json

<a id="L7"></a>import (
    <a id="L8"></a>&#34;reflect&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>type myStruct struct {
    <a id="L13"></a>T            bool;
    <a id="L14"></a>F            bool;
    <a id="L15"></a>S            string;
    <a id="L16"></a>I8           int8;
    <a id="L17"></a>I16          int16;
    <a id="L18"></a>I32          int32;
    <a id="L19"></a>I64          int64;
    <a id="L20"></a>U8           uint8;
    <a id="L21"></a>U16          uint16;
    <a id="L22"></a>U32          uint32;
    <a id="L23"></a>U64          uint64;
    <a id="L24"></a>I            int;
    <a id="L25"></a>U            uint;
    <a id="L26"></a>Fl           float;
    <a id="L27"></a>Fl32         float32;
    <a id="L28"></a>Fl64         float64;
    <a id="L29"></a>A            []string;
    <a id="L30"></a>My           *myStruct;
    <a id="L31"></a>Map          map[string][]int;
    <a id="L32"></a>MapStruct    map[string]myStruct;
    <a id="L33"></a>MapPtrStruct map[string]*myStruct;
<a id="L34"></a>}

<a id="L36"></a>const encoded = `{&#34;t&#34;:true,&#34;f&#34;:false,&#34;s&#34;:&#34;abc&#34;,&#34;i8&#34;:1,&#34;i16&#34;:2,&#34;i32&#34;:3,&#34;i64&#34;:4,`
    <a id="L37"></a>` &#34;u8&#34;:5,&#34;u16&#34;:6,&#34;u32&#34;:7,&#34;u64&#34;:8,`
    <a id="L38"></a>` &#34;i&#34;:-9,&#34;u&#34;:10,&#34;bogusfield&#34;:&#34;should be ignored&#34;,`
    <a id="L39"></a>` &#34;fl&#34;:11.5,&#34;fl32&#34;:12.25,&#34;fl64&#34;:13.75,`
    <a id="L40"></a>` &#34;a&#34;:[&#34;x&#34;,&#34;y&#34;,&#34;z&#34;],&#34;my&#34;:{&#34;s&#34;:&#34;subguy&#34;},`
    <a id="L41"></a>`&#34;map&#34;:{&#34;k1&#34;:[1,2,3],&#34;k2&#34;:[],&#34;k3&#34;:[3,4]},`
    <a id="L42"></a>`&#34;mapstruct&#34;:{&#34;m1&#34;:{&#34;u8&#34;:8}},`
    <a id="L43"></a>`&#34;mapptrstruct&#34;:{&#34;m1&#34;:{&#34;u8&#34;:8}}}`

<a id="L45"></a>var decodedMap = map[string][]int{
    <a id="L46"></a>&#34;k1&#34;: []int{1, 2, 3},
    <a id="L47"></a>&#34;k2&#34;: []int{},
    <a id="L48"></a>&#34;k3&#34;: []int{3, 4},
<a id="L49"></a>}

<a id="L51"></a>var decodedMapStruct = map[string]myStruct{
    <a id="L52"></a>&#34;m1&#34;: myStruct{U8: 8},
<a id="L53"></a>}

<a id="L55"></a>var decodedMapPtrStruct = map[string]*myStruct{
    <a id="L56"></a>&#34;m1&#34;: &amp;myStruct{U8: 8},
<a id="L57"></a>}

<a id="L59"></a>func check(t *testing.T, ok bool, name string, v interface{}) {
    <a id="L60"></a>if !ok {
        <a id="L61"></a>t.Errorf(&#34;%s = %v (BAD)&#34;, name, v)
    <a id="L62"></a>} else {
        <a id="L63"></a>t.Logf(&#34;%s = %v (good)&#34;, name, v)
    <a id="L64"></a>}
<a id="L65"></a>}

<a id="L67"></a>func TestUnmarshal(t *testing.T) {
    <a id="L68"></a>var m myStruct;
    <a id="L69"></a>m.F = true;
    <a id="L70"></a>ok, errtok := Unmarshal(encoded, &amp;m);
    <a id="L71"></a>if !ok {
        <a id="L72"></a>t.Fatalf(&#34;Unmarshal failed near %s&#34;, errtok)
    <a id="L73"></a>}
    <a id="L74"></a>check(t, m.T == true, &#34;t&#34;, m.T);
    <a id="L75"></a>check(t, m.F == false, &#34;f&#34;, m.F);
    <a id="L76"></a>check(t, m.S == &#34;abc&#34;, &#34;s&#34;, m.S);
    <a id="L77"></a>check(t, m.I8 == 1, &#34;i8&#34;, m.I8);
    <a id="L78"></a>check(t, m.I16 == 2, &#34;i16&#34;, m.I16);
    <a id="L79"></a>check(t, m.I32 == 3, &#34;i32&#34;, m.I32);
    <a id="L80"></a>check(t, m.I64 == 4, &#34;i64&#34;, m.I64);
    <a id="L81"></a>check(t, m.U8 == 5, &#34;u8&#34;, m.U8);
    <a id="L82"></a>check(t, m.U16 == 6, &#34;u16&#34;, m.U16);
    <a id="L83"></a>check(t, m.U32 == 7, &#34;u32&#34;, m.U32);
    <a id="L84"></a>check(t, m.U64 == 8, &#34;u64&#34;, m.U64);
    <a id="L85"></a>check(t, m.I == -9, &#34;i&#34;, m.I);
    <a id="L86"></a>check(t, m.U == 10, &#34;u&#34;, m.U);
    <a id="L87"></a>check(t, m.Fl == 11.5, &#34;fl&#34;, m.Fl);
    <a id="L88"></a>check(t, m.Fl32 == 12.25, &#34;fl32&#34;, m.Fl32);
    <a id="L89"></a>check(t, m.Fl64 == 13.75, &#34;fl64&#34;, m.Fl64);
    <a id="L90"></a>check(t, m.A != nil, &#34;a&#34;, m.A);
    <a id="L91"></a>if m.A != nil {
        <a id="L92"></a>check(t, m.A[0] == &#34;x&#34;, &#34;a[0]&#34;, m.A[0]);
        <a id="L93"></a>check(t, m.A[1] == &#34;y&#34;, &#34;a[1]&#34;, m.A[1]);
        <a id="L94"></a>check(t, m.A[2] == &#34;z&#34;, &#34;a[2]&#34;, m.A[2]);
    <a id="L95"></a>}
    <a id="L96"></a>check(t, m.My != nil, &#34;my&#34;, m.My);
    <a id="L97"></a>if m.My != nil {
        <a id="L98"></a>check(t, m.My.S == &#34;subguy&#34;, &#34;my.s&#34;, m.My.S)
    <a id="L99"></a>}
    <a id="L100"></a>check(t, reflect.DeepEqual(m.Map, decodedMap), &#34;map&#34;, m.Map);
    <a id="L101"></a>check(t, reflect.DeepEqual(m.MapStruct, decodedMapStruct), &#34;mapstruct&#34;, m.MapStruct);
    <a id="L102"></a>check(t, reflect.DeepEqual(m.MapPtrStruct, decodedMapPtrStruct), &#34;mapptrstruct&#34;, m.MapPtrStruct);
<a id="L103"></a>}
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
