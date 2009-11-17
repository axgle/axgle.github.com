<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/cgo/main.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/cgo/main.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Cgo; see gmp.go for an overview.</span>

<a id="L7"></a><span class="comment">// TODO(rsc):</span>
<a id="L8"></a><span class="comment">//	Emit correct line number annotations.</span>
<a id="L9"></a><span class="comment">//	Make 6g understand the annotations.</span>

<a id="L11"></a>package main

<a id="L13"></a>import (
    <a id="L14"></a>&#34;fmt&#34;;
    <a id="L15"></a>&#34;go/ast&#34;;
    <a id="L16"></a>&#34;os&#34;;
<a id="L17"></a>)

<a id="L19"></a>func usage() { fmt.Fprint(os.Stderr, &#34;usage: cgo [compiler options] file.go\n&#34;) }

<a id="L21"></a>var ptrSizeMap = map[string]int64{
    <a id="L22"></a>&#34;386&#34;: 4,
    <a id="L23"></a>&#34;amd64&#34;: 8,
    <a id="L24"></a>&#34;arm&#34;: 4,
<a id="L25"></a>}

<a id="L27"></a>var expandName = map[string]string{
    <a id="L28"></a>&#34;schar&#34;: &#34;signed char&#34;,
    <a id="L29"></a>&#34;uchar&#34;: &#34;unsigned char&#34;,
    <a id="L30"></a>&#34;ushort&#34;: &#34;unsigned short&#34;,
    <a id="L31"></a>&#34;uint&#34;: &#34;unsigned int&#34;,
    <a id="L32"></a>&#34;ulong&#34;: &#34;unsigned long&#34;,
    <a id="L33"></a>&#34;longlong&#34;: &#34;long long&#34;,
    <a id="L34"></a>&#34;ulonglong&#34;: &#34;unsigned long long&#34;,
<a id="L35"></a>}

<a id="L37"></a>func main() {
    <a id="L38"></a>args := os.Args;
    <a id="L39"></a>if len(args) &lt; 2 {
        <a id="L40"></a>usage();
        <a id="L41"></a>os.Exit(2);
    <a id="L42"></a>}
    <a id="L43"></a>gccOptions := args[1 : len(args)-1];
    <a id="L44"></a>input := args[len(args)-1];

    <a id="L46"></a>arch := os.Getenv(&#34;GOARCH&#34;);
    <a id="L47"></a>if arch == &#34;&#34; {
        <a id="L48"></a>fatal(&#34;$GOARCH is not set&#34;)
    <a id="L49"></a>}
    <a id="L50"></a>ptrSize, ok := ptrSizeMap[arch];
    <a id="L51"></a>if !ok {
        <a id="L52"></a>fatal(&#34;unknown architecture %s&#34;, arch)
    <a id="L53"></a>}

    <a id="L55"></a>p := openProg(input);
    <a id="L56"></a>for _, cref := range p.Crefs {
        <a id="L57"></a><span class="comment">// Convert C.ulong to C.unsigned long, etc.</span>
        <a id="L58"></a>if expand, ok := expandName[cref.Name]; ok {
            <a id="L59"></a>cref.Name = expand
        <a id="L60"></a>}
    <a id="L61"></a>}

    <a id="L63"></a>p.PtrSize = ptrSize;
    <a id="L64"></a>p.Preamble = p.Preamble + &#34;\n&#34; + builtinProlog;
    <a id="L65"></a>p.GccOptions = gccOptions;
    <a id="L66"></a>p.loadDebugInfo();
    <a id="L67"></a>p.Vardef = make(map[string]*Type);
    <a id="L68"></a>p.Funcdef = make(map[string]*FuncType);

    <a id="L70"></a>for _, cref := range p.Crefs {
        <a id="L71"></a>switch cref.Context {
        <a id="L72"></a>case &#34;call&#34;:
            <a id="L73"></a>if !cref.TypeName {
                <a id="L74"></a><span class="comment">// Is an actual function call.</span>
                <a id="L75"></a>*cref.Expr = &amp;ast.Ident{Value: &#34;_C_&#34; + cref.Name};
                <a id="L76"></a>p.Funcdef[cref.Name] = cref.FuncType;
                <a id="L77"></a>break;
            <a id="L78"></a>}
            <a id="L79"></a>*cref.Expr = cref.Type.Go;
        <a id="L80"></a>case &#34;expr&#34;:
            <a id="L81"></a>if cref.TypeName {
                <a id="L82"></a>error((*cref.Expr).Pos(), &#34;type C.%s used as expression&#34;, cref.Name)
            <a id="L83"></a>}
            <a id="L84"></a><span class="comment">// Reference to C variable.</span>
            <a id="L85"></a><span class="comment">// We declare a pointer and arrange to have it filled in.</span>
            <a id="L86"></a>*cref.Expr = &amp;ast.StarExpr{X: &amp;ast.Ident{Value: &#34;_C_&#34; + cref.Name}};
            <a id="L87"></a>p.Vardef[cref.Name] = cref.Type;
        <a id="L88"></a>case &#34;type&#34;:
            <a id="L89"></a>if !cref.TypeName {
                <a id="L90"></a>error((*cref.Expr).Pos(), &#34;expression C.%s used as type&#34;, cref.Name)
            <a id="L91"></a>}
            <a id="L92"></a>*cref.Expr = cref.Type.Go;
        <a id="L93"></a>}
    <a id="L94"></a>}
    <a id="L95"></a>if nerrors &gt; 0 {
        <a id="L96"></a>os.Exit(2)
    <a id="L97"></a>}

    <a id="L99"></a>p.PackagePath = p.Package;
    <a id="L100"></a>p.writeOutput(input);
<a id="L101"></a>}
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
