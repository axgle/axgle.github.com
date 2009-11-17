<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/compiler.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/compiler.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package eval

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;go/scanner&#34;;
    <a id="L10"></a>&#34;go/token&#34;;
<a id="L11"></a>)


<a id="L14"></a>type positioned interface {
    <a id="L15"></a>Pos() token.Position;
<a id="L16"></a>}


<a id="L19"></a><span class="comment">// A compiler captures information used throughout an entire</span>
<a id="L20"></a><span class="comment">// compilation.  Currently it includes only the error handler.</span>
<a id="L21"></a><span class="comment">//</span>
<a id="L22"></a><span class="comment">// TODO(austin) This might actually represent package level, in which</span>
<a id="L23"></a><span class="comment">// case it should be package compiler.</span>
<a id="L24"></a>type compiler struct {
    <a id="L25"></a>errors       scanner.ErrorHandler;
    <a id="L26"></a>numErrors    int;
    <a id="L27"></a>silentErrors int;
<a id="L28"></a>}

<a id="L30"></a>func (a *compiler) diagAt(pos positioned, format string, args ...) {
    <a id="L31"></a>a.errors.Error(pos.Pos(), fmt.Sprintf(format, args));
    <a id="L32"></a>a.numErrors++;
<a id="L33"></a>}

<a id="L35"></a>func (a *compiler) numError() int { return a.numErrors + a.silentErrors }

<a id="L37"></a><span class="comment">// The universal scope</span>
<a id="L38"></a>func newUniverse() *Scope {
    <a id="L39"></a>sc := &amp;Scope{nil, 0};
    <a id="L40"></a>sc.block = &amp;block{
        <a id="L41"></a>offset: 0,
        <a id="L42"></a>scope: sc,
        <a id="L43"></a>global: true,
        <a id="L44"></a>defs: make(map[string]Def),
    <a id="L45"></a>};
    <a id="L46"></a>return sc;
<a id="L47"></a>}

<a id="L49"></a>var universe *Scope = newUniverse()


<a id="L52"></a><span class="comment">// TODO(austin) These can all go in stmt.go now</span>
<a id="L53"></a>type label struct {
    <a id="L54"></a>name string;
    <a id="L55"></a>desc string;
    <a id="L56"></a><span class="comment">// The PC goto statements should jump to, or nil if this label</span>
    <a id="L57"></a><span class="comment">// cannot be goto&#39;d (such as an anonymous for loop label).</span>
    <a id="L58"></a>gotoPC *uint;
    <a id="L59"></a><span class="comment">// The PC break statements should jump to, or nil if a break</span>
    <a id="L60"></a><span class="comment">// statement is invalid.</span>
    <a id="L61"></a>breakPC *uint;
    <a id="L62"></a><span class="comment">// The PC continue statements should jump to, or nil if a</span>
    <a id="L63"></a><span class="comment">// continue statement is invalid.</span>
    <a id="L64"></a>continuePC *uint;
    <a id="L65"></a><span class="comment">// The position where this label was resolved.  If it has not</span>
    <a id="L66"></a><span class="comment">// been resolved yet, an invalid position.</span>
    <a id="L67"></a>resolved token.Position;
    <a id="L68"></a><span class="comment">// The position where this label was first jumped to.</span>
    <a id="L69"></a>used token.Position;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// A funcCompiler captures information used throughout the compilation</span>
<a id="L73"></a><span class="comment">// of a single function body.</span>
<a id="L74"></a>type funcCompiler struct {
    <a id="L75"></a>*compiler;
    <a id="L76"></a>fnType *FuncType;
    <a id="L77"></a><span class="comment">// Whether the out variables are named.  This affects what</span>
    <a id="L78"></a><span class="comment">// kinds of return statements are legal.</span>
    <a id="L79"></a>outVarsNamed bool;
    <a id="L80"></a>*codeBuf;
    <a id="L81"></a>flow   *flowBuf;
    <a id="L82"></a>labels map[string]*label;
<a id="L83"></a>}

<a id="L85"></a><span class="comment">// A blockCompiler captures information used throughout the compilation</span>
<a id="L86"></a><span class="comment">// of a single block within a function.</span>
<a id="L87"></a>type blockCompiler struct {
    <a id="L88"></a>*funcCompiler;
    <a id="L89"></a>block *block;
    <a id="L90"></a><span class="comment">// The label of this block, used for finding break and</span>
    <a id="L91"></a><span class="comment">// continue labels.</span>
    <a id="L92"></a>label *label;
    <a id="L93"></a><span class="comment">// The blockCompiler for the block enclosing this one, or nil</span>
    <a id="L94"></a><span class="comment">// for a function-level block.</span>
    <a id="L95"></a>parent *blockCompiler;
<a id="L96"></a>}
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
