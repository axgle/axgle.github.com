<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/ast/scope.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/go/ast/scope.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ast

<a id="L7"></a><span class="comment">// A Scope maintains the set of identifiers visible</span>
<a id="L8"></a><span class="comment">// in the scope and a link to the immediately surrounding</span>
<a id="L9"></a><span class="comment">// (outer) scope.</span>
<a id="L10"></a><span class="comment">//</span>
<a id="L11"></a><span class="comment">//	NOTE: WORK IN PROGRESS</span>
<a id="L12"></a><span class="comment">//</span>
<a id="L13"></a>type Scope struct {
    <a id="L14"></a>Outer *Scope;
    <a id="L15"></a>Names map[string]*Ident;
<a id="L16"></a>}


<a id="L19"></a><span class="comment">// NewScope creates a new scope nested in the outer scope.</span>
<a id="L20"></a>func NewScope(outer *Scope) *Scope { return &amp;Scope{outer, make(map[string]*Ident)} }


<a id="L23"></a><span class="comment">// Declare inserts an identifier into the scope s. If the</span>
<a id="L24"></a><span class="comment">// declaration succeeds, the result is true, if the identifier</span>
<a id="L25"></a><span class="comment">// exists already in the scope, the result is false.</span>
<a id="L26"></a><span class="comment">//</span>
<a id="L27"></a>func (s *Scope) Declare(ident *Ident) bool {
    <a id="L28"></a>if _, found := s.Names[ident.Value]; found {
        <a id="L29"></a>return false
    <a id="L30"></a>}
    <a id="L31"></a>s.Names[ident.Value] = ident;
    <a id="L32"></a>return true;
<a id="L33"></a>}


<a id="L36"></a><span class="comment">// Lookup looks up an identifier in the current scope chain.</span>
<a id="L37"></a><span class="comment">// If the identifier is found, it is returned; otherwise the</span>
<a id="L38"></a><span class="comment">// result is nil.</span>
<a id="L39"></a><span class="comment">//</span>
<a id="L40"></a>func (s *Scope) Lookup(name string) *Ident {
    <a id="L41"></a>for ; s != nil; s = s.Outer {
        <a id="L42"></a>if ident, found := s.Names[name]; found {
            <a id="L43"></a>return ident
        <a id="L44"></a>}
    <a id="L45"></a>}
    <a id="L46"></a>return nil;
<a id="L47"></a>}


<a id="L50"></a><span class="comment">// TODO(gri) Uncomment once this code is needed.</span>
<a id="L51"></a><span class="comment">/*</span>
<a id="L52"></a><span class="comment">var Universe = Scope {</span>
<a id="L53"></a><span class="comment">	Names: map[string]*Ident {</span>
<a id="L54"></a><span class="comment">		// basic types</span>
<a id="L55"></a><span class="comment">		&#34;bool&#34;: nil,</span>
<a id="L56"></a><span class="comment">		&#34;byte&#34;: nil,</span>
<a id="L57"></a><span class="comment">		&#34;int8&#34;: nil,</span>
<a id="L58"></a><span class="comment">		&#34;int16&#34;: nil,</span>
<a id="L59"></a><span class="comment">		&#34;int32&#34;: nil,</span>
<a id="L60"></a><span class="comment">		&#34;int64&#34;: nil,</span>
<a id="L61"></a><span class="comment">		&#34;uint8&#34;: nil,</span>
<a id="L62"></a><span class="comment">		&#34;uint16&#34;: nil,</span>
<a id="L63"></a><span class="comment">		&#34;uint32&#34;: nil,</span>
<a id="L64"></a><span class="comment">		&#34;uint64&#34;: nil,</span>
<a id="L65"></a><span class="comment">		&#34;float32&#34;: nil,</span>
<a id="L66"></a><span class="comment">		&#34;float64&#34;: nil,</span>
<a id="L67"></a><span class="comment">		&#34;string&#34;: nil,</span>

<a id="L69"></a><span class="comment">		// convenience types</span>
<a id="L70"></a><span class="comment">		&#34;int&#34;: nil,</span>
<a id="L71"></a><span class="comment">		&#34;uint&#34;: nil,</span>
<a id="L72"></a><span class="comment">		&#34;uintptr&#34;: nil,</span>
<a id="L73"></a><span class="comment">		&#34;float&#34;: nil,</span>

<a id="L75"></a><span class="comment">		// constants</span>
<a id="L76"></a><span class="comment">		&#34;false&#34;: nil,</span>
<a id="L77"></a><span class="comment">		&#34;true&#34;: nil,</span>
<a id="L78"></a><span class="comment">		&#34;iota&#34;: nil,</span>
<a id="L79"></a><span class="comment">		&#34;nil&#34;: nil,</span>

<a id="L81"></a><span class="comment">		// functions</span>
<a id="L82"></a><span class="comment">		&#34;cap&#34;: nil,</span>
<a id="L83"></a><span class="comment">		&#34;len&#34;: nil,</span>
<a id="L84"></a><span class="comment">		&#34;new&#34;: nil,</span>
<a id="L85"></a><span class="comment">		&#34;make&#34;: nil,</span>
<a id="L86"></a><span class="comment">		&#34;panic&#34;: nil,</span>
<a id="L87"></a><span class="comment">		&#34;panicln&#34;: nil,</span>
<a id="L88"></a><span class="comment">		&#34;print&#34;: nil,</span>
<a id="L89"></a><span class="comment">		&#34;println&#34;: nil,</span>
<a id="L90"></a><span class="comment">	}</span>
<a id="L91"></a><span class="comment">}</span>
<a id="L92"></a><span class="comment">*/</span></pre>

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
