<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/godoc/snippet.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/godoc/snippet.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file contains the infrastructure to create a code</span>
<a id="L6"></a><span class="comment">// snippet for search results.</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a><span class="comment">// Note: At the moment, this only creates HTML snippets.</span>

<a id="L10"></a>package main

<a id="L12"></a>import (
    <a id="L13"></a>&#34;bytes&#34;;
    <a id="L14"></a>&#34;go/ast&#34;;
    <a id="L15"></a>&#34;go/printer&#34;;
    <a id="L16"></a>&#34;fmt&#34;;
    <a id="L17"></a>&#34;strings&#34;;
<a id="L18"></a>)


<a id="L21"></a>type Snippet struct {
    <a id="L22"></a>Line int;
    <a id="L23"></a>Text string;
<a id="L24"></a>}


<a id="L27"></a>type snippetStyler struct {
    <a id="L28"></a>Styler;               <span class="comment">// defined in godoc.go</span>
    <a id="L29"></a>highlight *ast.Ident; <span class="comment">// identifier to highlight</span>
<a id="L30"></a>}


<a id="L33"></a>func (s *snippetStyler) LineTag(line int) (text []uint8, tag printer.HTMLTag) {
    <a id="L34"></a>return <span class="comment">// no LineTag for snippets</span>
<a id="L35"></a>}


<a id="L38"></a>func (s *snippetStyler) Ident(id *ast.Ident) (text []byte, tag printer.HTMLTag) {
    <a id="L39"></a>text = strings.Bytes(id.Value);
    <a id="L40"></a>if s.highlight == id {
        <a id="L41"></a>tag = printer.HTMLTag{&#34;&lt;span class=highlight&gt;&#34;, &#34;&lt;/span&gt;&#34;}
    <a id="L42"></a>}
    <a id="L43"></a>return;
<a id="L44"></a>}


<a id="L47"></a>func newSnippet(decl ast.Decl, id *ast.Ident) *Snippet {
    <a id="L48"></a>var buf bytes.Buffer;
    <a id="L49"></a>writeNode(&amp;buf, decl, true, &amp;snippetStyler{highlight: id});
    <a id="L50"></a>return &amp;Snippet{id.Pos().Line, buf.String()};
<a id="L51"></a>}


<a id="L54"></a>func findSpec(list []ast.Spec, id *ast.Ident) ast.Spec {
    <a id="L55"></a>for _, spec := range list {
        <a id="L56"></a>switch s := spec.(type) {
        <a id="L57"></a>case *ast.ImportSpec:
            <a id="L58"></a>if s.Name == id {
                <a id="L59"></a>return s
            <a id="L60"></a>}
        <a id="L61"></a>case *ast.ValueSpec:
            <a id="L62"></a>for _, n := range s.Names {
                <a id="L63"></a>if n == id {
                    <a id="L64"></a>return s
                <a id="L65"></a>}
            <a id="L66"></a>}
        <a id="L67"></a>case *ast.TypeSpec:
            <a id="L68"></a>if s.Name == id {
                <a id="L69"></a>return s
            <a id="L70"></a>}
        <a id="L71"></a>}
    <a id="L72"></a>}
    <a id="L73"></a>return nil;
<a id="L74"></a>}


<a id="L77"></a>func genSnippet(d *ast.GenDecl, id *ast.Ident) *Snippet {
    <a id="L78"></a>s := findSpec(d.Specs, id);
    <a id="L79"></a>if s == nil {
        <a id="L80"></a>return nil <span class="comment">//  declaration doesn&#39;t contain id - exit gracefully</span>
    <a id="L81"></a>}

    <a id="L83"></a><span class="comment">// only use the spec containing the id for the snippet</span>
    <a id="L84"></a>dd := &amp;ast.GenDecl{d.Doc, d.Position, d.Tok, d.Lparen, []ast.Spec{s}, d.Rparen};

    <a id="L86"></a>return newSnippet(dd, id);
<a id="L87"></a>}


<a id="L90"></a>func funcSnippet(d *ast.FuncDecl, id *ast.Ident) *Snippet {
    <a id="L91"></a>if d.Name != id {
        <a id="L92"></a>return nil <span class="comment">//  declaration doesn&#39;t contain id - exit gracefully</span>
    <a id="L93"></a>}

    <a id="L95"></a><span class="comment">// only use the function signature for the snippet</span>
    <a id="L96"></a>dd := &amp;ast.FuncDecl{d.Doc, d.Recv, d.Name, d.Type, nil};

    <a id="L98"></a>return newSnippet(dd, id);
<a id="L99"></a>}


<a id="L102"></a><span class="comment">// NewSnippet creates a text snippet from a declaration decl containing an</span>
<a id="L103"></a><span class="comment">// identifier id. Parts of the declaration not containing the identifier</span>
<a id="L104"></a><span class="comment">// may be removed for a more compact snippet.</span>
<a id="L105"></a><span class="comment">//</span>
<a id="L106"></a>func NewSnippet(decl ast.Decl, id *ast.Ident) (s *Snippet) {
    <a id="L107"></a>switch d := decl.(type) {
    <a id="L108"></a>case *ast.GenDecl:
        <a id="L109"></a>s = genSnippet(d, id)
    <a id="L110"></a>case *ast.FuncDecl:
        <a id="L111"></a>s = funcSnippet(d, id)
    <a id="L112"></a>}

    <a id="L114"></a><span class="comment">// handle failure gracefully</span>
    <a id="L115"></a>if s == nil {
        <a id="L116"></a>s = &amp;Snippet{
            <a id="L117"></a>id.Pos().Line,
            <a id="L118"></a>fmt.Sprintf(`could not generate a snippet for &lt;span class=&#34;highlight&#34;&gt;%s&lt;/span&gt;`, id.Value),
        <a id="L119"></a>}
    <a id="L120"></a>}
    <a id="L121"></a>return;
<a id="L122"></a>}
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
