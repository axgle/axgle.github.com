<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/json/generic_test.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/json/generic_test.go</h1>

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
    <a id="L8"></a>&#34;testing&#34;;
<a id="L9"></a>)

<a id="L11"></a>var jsontests = []string{
    <a id="L12"></a>`null`,
    <a id="L13"></a>`true`,
    <a id="L14"></a>`false`,
    <a id="L15"></a>`&#34;abc&#34;`,
    <a id="L16"></a>`123`,
    <a id="L17"></a>`0.1`,
    <a id="L18"></a>`1e-10`,
    <a id="L19"></a>`[]`,
    <a id="L20"></a>`[1,2,3,4]`,
    <a id="L21"></a>`[1,2,&#34;abc&#34;,null,true,false]`,
    <a id="L22"></a>`{}`,
    <a id="L23"></a>`{&#34;a&#34;:1}`,
    <a id="L24"></a>`&#34;q\u0302&#34;`,
<a id="L25"></a>}

<a id="L27"></a>func TestJson(t *testing.T) {
    <a id="L28"></a>for i := 0; i &lt; len(jsontests); i++ {
        <a id="L29"></a>val, ok, errtok := StringToJson(jsontests[i]);
        <a id="L30"></a>if !ok {
            <a id="L31"></a>t.Errorf(&#34;StringToJson(%#q) =&gt; error near %v&#34;, jsontests[i], errtok);
            <a id="L32"></a>continue;
        <a id="L33"></a>}
        <a id="L34"></a>str := JsonToString(val);
        <a id="L35"></a>if str != jsontests[i] {
            <a id="L36"></a>t.Errorf(&#34;JsonToString(StringToJson(%#q)) = %#q&#34;, jsontests[i], str);
            <a id="L37"></a>continue;
        <a id="L38"></a>}
    <a id="L39"></a>}
<a id="L40"></a>}

<a id="L42"></a>func TestJsonMap(t *testing.T) {
    <a id="L43"></a>values := make(map[string]Json);
    <a id="L44"></a>mapstr := &#34;{&#34;;
    <a id="L45"></a>for i := 0; i &lt; len(jsontests); i++ {
        <a id="L46"></a>val, ok, errtok := StringToJson(jsontests[i]);
        <a id="L47"></a>if !ok {
            <a id="L48"></a>t.Errorf(&#34;StringToJson(%#q) =&gt; error near %v&#34;, jsontests[i], errtok)
        <a id="L49"></a>}
        <a id="L50"></a>if i &gt; 0 {
            <a id="L51"></a>mapstr += &#34;,&#34;
        <a id="L52"></a>}
        <a id="L53"></a>values[jsontests[i]] = val;
        <a id="L54"></a>mapstr += Quote(jsontests[i]);
        <a id="L55"></a>mapstr += &#34;:&#34;;
        <a id="L56"></a>mapstr += JsonToString(val);
    <a id="L57"></a>}
    <a id="L58"></a>mapstr += &#34;}&#34;;

    <a id="L60"></a>mapv, ok, errtok := StringToJson(mapstr);
    <a id="L61"></a>if !ok {
        <a id="L62"></a>t.Fatalf(&#34;StringToJson(%#q) =&gt; error near %v&#34;, mapstr, errtok)
    <a id="L63"></a>}
    <a id="L64"></a>if mapv == nil {
        <a id="L65"></a>t.Fatalf(&#34;StringToJson(%#q) =&gt; nil, %v, %v&#34;, mapstr, ok, errtok)
    <a id="L66"></a>}
    <a id="L67"></a>if cnt := mapv.Len(); cnt != len(jsontests) {
        <a id="L68"></a>t.Errorf(&#34;StringToJson(%#q).Len() =&gt; %v, want %v&#34;, mapstr, cnt,
            <a id="L69"></a>len(jsontests))
    <a id="L70"></a>}
    <a id="L71"></a>for k, v := range values {
        <a id="L72"></a>if v1 := mapv.Get(k); !Equal(v1, v) {
            <a id="L73"></a>t.Errorf(&#34;MapTest: Walk(%#q) =&gt; %v, want %v&#34;, k, v1, v)
        <a id="L74"></a>}
    <a id="L75"></a>}
<a id="L76"></a>}
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
