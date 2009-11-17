<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/parser/parser_test.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/go/parser/parser_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package parser

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)


<a id="L13"></a>var illegalInputs = []interface{}{
    <a id="L14"></a>nil,
    <a id="L15"></a>3.14,
    <a id="L16"></a>[]byte(nil),
    <a id="L17"></a>&#34;foo!&#34;,
<a id="L18"></a>}


<a id="L21"></a>func TestParseIllegalInputs(t *testing.T) {
    <a id="L22"></a>for _, src := range illegalInputs {
        <a id="L23"></a>_, err := ParseFile(&#34;&#34;, src, 0);
        <a id="L24"></a>if err == nil {
            <a id="L25"></a>t.Errorf(&#34;ParseFile(%v) should have failed&#34;, src)
        <a id="L26"></a>}
    <a id="L27"></a>}
<a id="L28"></a>}


<a id="L31"></a>var validPrograms = []interface{}{
    <a id="L32"></a>`package main`,
    <a id="L33"></a>`package main import &#34;fmt&#34; func main() { fmt.Println(&#34;Hello, World!&#34;) }`,
<a id="L34"></a>}


<a id="L37"></a>func TestParseValidPrograms(t *testing.T) {
    <a id="L38"></a>for _, src := range validPrograms {
        <a id="L39"></a>_, err := ParseFile(&#34;&#34;, src, 0);
        <a id="L40"></a>if err != nil {
            <a id="L41"></a>t.Errorf(&#34;ParseFile(%q): %v&#34;, src, err)
        <a id="L42"></a>}
    <a id="L43"></a>}
<a id="L44"></a>}


<a id="L47"></a>var validFiles = []string{
    <a id="L48"></a>&#34;parser.go&#34;,
    <a id="L49"></a>&#34;parser_test.go&#34;,
<a id="L50"></a>}


<a id="L53"></a>func TestParse3(t *testing.T) {
    <a id="L54"></a>for _, filename := range validFiles {
        <a id="L55"></a>_, err := ParseFile(filename, nil, 0);
        <a id="L56"></a>if err != nil {
            <a id="L57"></a>t.Errorf(&#34;ParseFile(%s): %v&#34;, filename, err)
        <a id="L58"></a>}
    <a id="L59"></a>}
<a id="L60"></a>}


<a id="L63"></a>func nameFilter(filename string) bool {
    <a id="L64"></a>switch filename {
    <a id="L65"></a>case &#34;parser.go&#34;:
    <a id="L66"></a>case &#34;interface.go&#34;:
    <a id="L67"></a>case &#34;parser_test.go&#34;:
    <a id="L68"></a>default:
        <a id="L69"></a>return false
    <a id="L70"></a>}
    <a id="L71"></a>return true;
<a id="L72"></a>}


<a id="L75"></a>func dirFilter(d *os.Dir) bool { return nameFilter(d.Name) }


<a id="L78"></a>func TestParse4(t *testing.T) {
    <a id="L79"></a>path := &#34;.&#34;;
    <a id="L80"></a>pkg, err := ParsePackage(path, dirFilter, 0);
    <a id="L81"></a>if err != nil {
        <a id="L82"></a>t.Fatalf(&#34;ParsePackage(%s): %v&#34;, path, err)
    <a id="L83"></a>}
    <a id="L84"></a>if pkg.Name != &#34;parser&#34; {
        <a id="L85"></a>t.Errorf(&#34;incorrect package name: %s&#34;, pkg.Name)
    <a id="L86"></a>}
    <a id="L87"></a>for filename, _ := range pkg.Files {
        <a id="L88"></a>if !nameFilter(filename) {
            <a id="L89"></a>t.Errorf(&#34;unexpected package file: %s&#34;, filename)
        <a id="L90"></a>}
    <a id="L91"></a>}
<a id="L92"></a>}
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
