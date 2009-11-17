<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/stmt_test.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/stmt_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package eval

<a id="L7"></a>import &#34;testing&#34;

<a id="L9"></a>var atLeastOneDecl = &#34;at least one new variable must be declared&#34;

<a id="L11"></a>var stmtTests = []test{
    <a id="L12"></a><span class="comment">// Short declarations</span>
    <a id="L13"></a>Val1(&#34;x := i&#34;, &#34;x&#34;, 1),
    <a id="L14"></a>Val1(&#34;x := f&#34;, &#34;x&#34;, 1.0),
    <a id="L15"></a><span class="comment">// Type defaulting</span>
    <a id="L16"></a>Val1(&#34;a := 42&#34;, &#34;a&#34;, 42),
    <a id="L17"></a>Val1(&#34;a := 1.0&#34;, &#34;a&#34;, 1.0),
    <a id="L18"></a><span class="comment">// Parallel assignment</span>
    <a id="L19"></a>Val2(&#34;a, b := 1, 2&#34;, &#34;a&#34;, 1, &#34;b&#34;, 2),
    <a id="L20"></a>Val2(&#34;a, i := 1, 2&#34;, &#34;a&#34;, 1, &#34;i&#34;, 2),
    <a id="L21"></a>CErr(&#34;a, i := 1, f&#34;, opTypes),
    <a id="L22"></a><span class="comment">// TODO(austin) The parser produces an error message for this</span>
    <a id="L23"></a><span class="comment">// one that&#39;s inconsistent with the errors I give for other</span>
    <a id="L24"></a><span class="comment">// things</span>
    <a id="L25"></a><span class="comment">//CErr(&#34;a, b := 1, 2, 3&#34;, &#34;too many&#34;),</span>
    <a id="L26"></a>CErr(&#34;a, b := 1, 2, 3&#34;, &#34;arity&#34;),
    <a id="L27"></a>CErr(&#34;a := 1, 2&#34;, &#34;too many&#34;),
    <a id="L28"></a>CErr(&#34;a, b := 1&#34;, &#34;not enough&#34;),
    <a id="L29"></a><span class="comment">// Mixed declarations</span>
    <a id="L30"></a>CErr(&#34;i := 1&#34;, atLeastOneDecl),
    <a id="L31"></a>CErr(&#34;i, u := 1, 2&#34;, atLeastOneDecl),
    <a id="L32"></a>Val2(&#34;i, x := 2, f&#34;, &#34;i&#34;, 2, &#34;x&#34;, 1.0),
    <a id="L33"></a><span class="comment">// Various errors</span>
    <a id="L34"></a>CErr(&#34;1 := 2&#34;, &#34;left side of := must be a name&#34;),
    <a id="L35"></a>CErr(&#34;c, a := 1, 1&#34;, &#34;cannot assign&#34;),
    <a id="L36"></a><span class="comment">// Unpacking</span>
    <a id="L37"></a>Val2(&#34;x, y := oneTwo()&#34;, &#34;x&#34;, 1, &#34;y&#34;, 2),
    <a id="L38"></a>CErr(&#34;x := oneTwo()&#34;, &#34;too many&#34;),
    <a id="L39"></a>CErr(&#34;x, y, z := oneTwo()&#34;, &#34;not enough&#34;),
    <a id="L40"></a>CErr(&#34;x, y := oneTwo(), 2&#34;, &#34;multi-valued&#34;),
    <a id="L41"></a>CErr(&#34;x := oneTwo()+2&#34;, opTypes),
    <a id="L42"></a><span class="comment">// TOOD(austin) This error message is weird</span>
    <a id="L43"></a>CErr(&#34;x := void()&#34;, &#34;not enough&#34;),
    <a id="L44"></a><span class="comment">// Placeholders</span>
    <a id="L45"></a>CErr(&#34;x := 1+\&#34;x\&#34;; i=x+1&#34;, opTypes),

    <a id="L47"></a><span class="comment">// Assignment</span>
    <a id="L48"></a>Val1(&#34;i = 2&#34;, &#34;i&#34;, 2),
    <a id="L49"></a>Val1(&#34;(i) = 2&#34;, &#34;i&#34;, 2),
    <a id="L50"></a>CErr(&#34;1 = 2&#34;, &#34;cannot assign&#34;),
    <a id="L51"></a>CErr(&#34;1-1 = 2&#34;, &#34;- expression&#34;),
    <a id="L52"></a>Val1(&#34;i = 2.0&#34;, &#34;i&#34;, 2),
    <a id="L53"></a>CErr(&#34;i = 2.2&#34;, constantTruncated),
    <a id="L54"></a>CErr(&#34;u = -2&#34;, constantUnderflows),
    <a id="L55"></a>CErr(&#34;i = f&#34;, opTypes),
    <a id="L56"></a>CErr(&#34;i, u = 0, f&#34;, opTypes),
    <a id="L57"></a>CErr(&#34;i, u = 0, f&#34;, &#34;value 2&#34;),
    <a id="L58"></a>Val2(&#34;i, i2 = i2, i&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 1),
    <a id="L59"></a>CErr(&#34;c = 1&#34;, &#34;cannot assign&#34;),

    <a id="L61"></a>Val1(&#34;x := &amp;i; *x = 2&#34;, &#34;i&#34;, 2),

    <a id="L63"></a>Val1(&#34;ai[0] = 42&#34;, &#34;ai&#34;, varray{42, 2}),
    <a id="L64"></a>Val1(&#34;aai[1] = ai; ai[0] = 42&#34;, &#34;aai&#34;, varray{varray{1, 2}, varray{1, 2}}),
    <a id="L65"></a>Val1(&#34;aai = aai2&#34;, &#34;aai&#34;, varray{varray{5, 6}, varray{7, 8}}),

    <a id="L67"></a><span class="comment">// Assignment conversions</span>
    <a id="L68"></a>Run(&#34;var sl []int; sl = &amp;ai&#34;),
    <a id="L69"></a>CErr(&#34;type ST []int; type AT *[2]int; var x AT = &amp;ai; var y ST = x&#34;, opTypes),
    <a id="L70"></a>Run(&#34;type ST []int; var y ST = &amp;ai&#34;),
    <a id="L71"></a>Run(&#34;type AT *[2]int; var x AT = &amp;ai; var y []int = x&#34;),

    <a id="L73"></a><span class="comment">// Op-assignment</span>
    <a id="L74"></a>Val1(&#34;i += 2&#34;, &#34;i&#34;, 3),
    <a id="L75"></a>Val(&#34;i&#34;, 1),
    <a id="L76"></a>Val1(&#34;f += 2&#34;, &#34;f&#34;, 3.0),
    <a id="L77"></a>CErr(&#34;2 += 2&#34;, &#34;cannot assign&#34;),
    <a id="L78"></a>CErr(&#34;i, j += 2&#34;, &#34;cannot be combined&#34;),
    <a id="L79"></a>CErr(&#34;i += 2, 3&#34;, &#34;cannot be combined&#34;),
    <a id="L80"></a>Val2(&#34;s2 := s; s += \&#34;def\&#34;&#34;, &#34;s2&#34;, &#34;abc&#34;, &#34;s&#34;, &#34;abcdef&#34;),
    <a id="L81"></a>CErr(&#34;s += 1&#34;, opTypes),
    <a id="L82"></a><span class="comment">// Single evaluation</span>
    <a id="L83"></a>Val2(&#34;ai[func()int{i+=1;return 0}()] *= 3; i2 = ai[0]&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 3),

    <a id="L85"></a><span class="comment">// Type declarations</span>
    <a id="L86"></a><span class="comment">// Identifiers</span>
    <a id="L87"></a>Run(&#34;type T int&#34;),
    <a id="L88"></a>CErr(&#34;type T x&#34;, &#34;undefined&#34;),
    <a id="L89"></a>CErr(&#34;type T c&#34;, &#34;constant&#34;),
    <a id="L90"></a>CErr(&#34;type T i&#34;, &#34;variable&#34;),
    <a id="L91"></a>CErr(&#34;type T T&#34;, &#34;recursive&#34;),
    <a id="L92"></a>CErr(&#34;type T x; type U T; var v U; v = 1&#34;, &#34;undefined&#34;),
    <a id="L93"></a><span class="comment">// Pointer types</span>
    <a id="L94"></a>Run(&#34;type T *int&#34;),
    <a id="L95"></a>Run(&#34;type T *T&#34;),
    <a id="L96"></a><span class="comment">// Array types</span>
    <a id="L97"></a>Run(&#34;type T [5]int&#34;),
    <a id="L98"></a>Run(&#34;type T [c+42/2]int&#34;),
    <a id="L99"></a>Run(&#34;type T [2.0]int&#34;),
    <a id="L100"></a>CErr(&#34;type T [i]int&#34;, &#34;constant expression&#34;),
    <a id="L101"></a>CErr(&#34;type T [2.5]int&#34;, constantTruncated),
    <a id="L102"></a>CErr(&#34;type T [-1]int&#34;, &#34;negative&#34;),
    <a id="L103"></a>CErr(&#34;type T [2]T&#34;, &#34;recursive&#34;),
    <a id="L104"></a><span class="comment">// Struct types</span>
    <a id="L105"></a>Run(&#34;type T struct { a int; b int }&#34;),
    <a id="L106"></a>Run(&#34;type T struct { a int; int }&#34;),
    <a id="L107"></a>Run(&#34;type T struct { x *T }&#34;),
    <a id="L108"></a>Run(&#34;type T int; type U struct { T }&#34;),
    <a id="L109"></a>CErr(&#34;type T *int; type U struct { T }&#34;, &#34;embedded.*pointer&#34;),
    <a id="L110"></a>CErr(&#34;type T *struct { T }&#34;, &#34;embedded.*pointer&#34;),
    <a id="L111"></a>CErr(&#34;type T struct { a int; a int }&#34;, &#34; a .*redeclared.*:1:17&#34;),
    <a id="L112"></a>CErr(&#34;type T struct { int; int }&#34;, &#34;int .*redeclared.*:1:17&#34;),
    <a id="L113"></a>CErr(&#34;type T struct { int int; int }&#34;, &#34;int .*redeclared.*:1:17&#34;),
    <a id="L114"></a>Run(&#34;type T struct { x *struct { T } }&#34;),
    <a id="L115"></a>CErr(&#34;type T struct { x struct { T } }&#34;, &#34;recursive&#34;),
    <a id="L116"></a>CErr(&#34;type T struct { x }; type U struct { T }&#34;, &#34;undefined&#34;),
    <a id="L117"></a><span class="comment">// Function types</span>
    <a id="L118"></a>Run(&#34;type T func()&#34;),
    <a id="L119"></a>Run(&#34;type T func(a, b int) int&#34;),
    <a id="L120"></a>Run(&#34;type T func(a, b int) (x int, y int)&#34;),
    <a id="L121"></a>Run(&#34;type T func(a, a int) (a int, a int)&#34;),
    <a id="L122"></a>Run(&#34;type T func(a, b int) (x, y int)&#34;),
    <a id="L123"></a>Run(&#34;type T func(int, int) (int, int)&#34;),
    <a id="L124"></a>CErr(&#34;type T func(x); type U T&#34;, &#34;undefined&#34;),
    <a id="L125"></a>CErr(&#34;type T func(a T)&#34;, &#34;recursive&#34;),
    <a id="L126"></a><span class="comment">// Interface types</span>
    <a id="L127"></a>Run(&#34;type T interface {x(a, b int) int}&#34;),
    <a id="L128"></a>Run(&#34;type T interface {x(a, b int) int}; type U interface {T; y(c int)}&#34;),
    <a id="L129"></a>CErr(&#34;type T interface {x(a int); x()}&#34;, &#34;method x redeclared&#34;),
    <a id="L130"></a>CErr(&#34;type T interface {x()}; type U interface {T; x()}&#34;, &#34;method x redeclared&#34;),
    <a id="L131"></a>CErr(&#34;type T int; type U interface {T}&#34;, &#34;embedded type&#34;),
    <a id="L132"></a><span class="comment">// Parens</span>
    <a id="L133"></a>Run(&#34;type T (int)&#34;),

    <a id="L135"></a><span class="comment">// Variable declarations</span>
    <a id="L136"></a>Val2(&#34;var x int&#34;, &#34;i&#34;, 1, &#34;x&#34;, 0),
    <a id="L137"></a>Val1(&#34;var x = 1&#34;, &#34;x&#34;, 1),
    <a id="L138"></a>Val1(&#34;var x = 1.0&#34;, &#34;x&#34;, 1.0),
    <a id="L139"></a>Val1(&#34;var x int = 1.0&#34;, &#34;x&#34;, 1),
    <a id="L140"></a><span class="comment">// Placeholders</span>
    <a id="L141"></a>CErr(&#34;var x foo; x = 1&#34;, &#34;undefined&#34;),
    <a id="L142"></a>CErr(&#34;var x foo = 1; x = 1&#34;, &#34;undefined&#34;),
    <a id="L143"></a><span class="comment">// Redeclaration</span>
    <a id="L144"></a>CErr(&#34;var i, x int&#34;, &#34; i .*redeclared&#34;),
    <a id="L145"></a>CErr(&#34;var x int; var x int&#34;, &#34; x .*redeclared.*:1:5&#34;),

    <a id="L147"></a><span class="comment">// Expression statements</span>
    <a id="L148"></a>CErr(&#34;x := func(){ 1-1 }&#34;, &#34;expression statement&#34;),
    <a id="L149"></a>CErr(&#34;x := func(){ 1-1 }&#34;, &#34;- expression&#34;),
    <a id="L150"></a>Val1(&#34;fn(2)&#34;, &#34;i&#34;, 1),

    <a id="L152"></a><span class="comment">// IncDec statements</span>
    <a id="L153"></a>Val1(&#34;i++&#34;, &#34;i&#34;, 2),
    <a id="L154"></a>Val1(&#34;i--&#34;, &#34;i&#34;, 0),
    <a id="L155"></a>Val1(&#34;u++&#34;, &#34;u&#34;, uint(2)),
    <a id="L156"></a>Val1(&#34;u--&#34;, &#34;u&#34;, uint(0)),
    <a id="L157"></a>Val1(&#34;f++&#34;, &#34;f&#34;, 2.0),
    <a id="L158"></a>Val1(&#34;f--&#34;, &#34;f&#34;, 0.0),
    <a id="L159"></a><span class="comment">// Single evaluation</span>
    <a id="L160"></a>Val2(&#34;ai[func()int{i+=1;return 0}()]++; i2 = ai[0]&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 2),
    <a id="L161"></a><span class="comment">// Operand types</span>
    <a id="L162"></a>CErr(&#34;s++&#34;, opTypes),
    <a id="L163"></a>CErr(&#34;s++&#34;, &#34;&#39;\\+\\+&#39;&#34;),
    <a id="L164"></a>CErr(&#34;2++&#34;, &#34;cannot assign&#34;),
    <a id="L165"></a>CErr(&#34;c++&#34;, &#34;cannot assign&#34;),

    <a id="L167"></a><span class="comment">// Function scoping</span>
    <a id="L168"></a>Val1(&#34;fn1 := func() { i=2 }; fn1()&#34;, &#34;i&#34;, 2),
    <a id="L169"></a>Val1(&#34;fn1 := func() { i:=2 }; fn1()&#34;, &#34;i&#34;, 1),
    <a id="L170"></a>Val2(&#34;fn1 := func() int { i=2; i:=3; i=4; return i }; x := fn1()&#34;, &#34;i&#34;, 2, &#34;x&#34;, 4),

    <a id="L172"></a><span class="comment">// Basic returns</span>
    <a id="L173"></a>CErr(&#34;fn1 := func() int {}&#34;, &#34;return&#34;),
    <a id="L174"></a>Run(&#34;fn1 := func() {}&#34;),
    <a id="L175"></a>CErr(&#34;fn1 := func() (r int) {}&#34;, &#34;return&#34;),
    <a id="L176"></a>Val1(&#34;fn1 := func() (r int) {return}; i = fn1()&#34;, &#34;i&#34;, 0),
    <a id="L177"></a>Val1(&#34;fn1 := func() (r int) {r = 2; return}; i = fn1()&#34;, &#34;i&#34;, 2),
    <a id="L178"></a>Val1(&#34;fn1 := func() (r int) {return 2}; i = fn1()&#34;, &#34;i&#34;, 2),
    <a id="L179"></a>Val1(&#34;fn1 := func(int) int {return 2}; i = fn1(1)&#34;, &#34;i&#34;, 2),

    <a id="L181"></a><span class="comment">// Multi-valued returns</span>
    <a id="L182"></a>Val2(&#34;fn1 := func() (bool, int) {return true, 2}; x, y := fn1()&#34;, &#34;x&#34;, true, &#34;y&#34;, 2),
    <a id="L183"></a>CErr(&#34;fn1 := func() int {return}&#34;, &#34;not enough values&#34;),
    <a id="L184"></a>CErr(&#34;fn1 := func() int {return 1,2}&#34;, &#34;too many values&#34;),
    <a id="L185"></a>CErr(&#34;fn1 := func() {return 1}&#34;, &#34;too many values&#34;),
    <a id="L186"></a>CErr(&#34;fn1 := func() (int,int,int) {return 1,2}&#34;, &#34;not enough values&#34;),
    <a id="L187"></a>Val2(&#34;fn1 := func() (int, int) {return oneTwo()}; x, y := fn1()&#34;, &#34;x&#34;, 1, &#34;y&#34;, 2),
    <a id="L188"></a>CErr(&#34;fn1 := func() int {return oneTwo()}&#34;, &#34;too many values&#34;),
    <a id="L189"></a>CErr(&#34;fn1 := func() (int,int,int) {return oneTwo()}&#34;, &#34;not enough values&#34;),
    <a id="L190"></a>Val1(&#34;fn1 := func(x,y int) int {return x+y}; x := fn1(oneTwo())&#34;, &#34;x&#34;, 3),

    <a id="L192"></a><span class="comment">// Return control flow</span>
    <a id="L193"></a>Val2(&#34;fn1 := func(x *int) bool { *x = 2; return true; *x = 3; }; x := fn1(&amp;i)&#34;, &#34;i&#34;, 2, &#34;x&#34;, true),

    <a id="L195"></a><span class="comment">// Break/continue/goto/fallthrough</span>
    <a id="L196"></a>CErr(&#34;break&#34;, &#34;outside&#34;),
    <a id="L197"></a>CErr(&#34;break foo&#34;, &#34;break.*foo.*not defined&#34;),
    <a id="L198"></a>CErr(&#34;continue&#34;, &#34;outside&#34;),
    <a id="L199"></a>CErr(&#34;continue foo&#34;, &#34;continue.*foo.*not defined&#34;),
    <a id="L200"></a>CErr(&#34;fallthrough&#34;, &#34;outside&#34;),
    <a id="L201"></a>CErr(&#34;goto foo&#34;, &#34;foo.*not defined&#34;),
    <a id="L202"></a>CErr(&#34; foo: foo:;&#34;, &#34;foo.*redeclared.*:1:2&#34;),
    <a id="L203"></a>Val1(&#34;i+=2; goto L; i+=4; L: i+=8&#34;, &#34;i&#34;, 1+2+8),
    <a id="L204"></a><span class="comment">// Return checking</span>
    <a id="L205"></a>CErr(&#34;fn1 := func() int { goto L; return 1; L: }&#34;, &#34;return&#34;),
    <a id="L206"></a>Run(&#34;fn1 := func() int { L: goto L; i = 2 }&#34;),
    <a id="L207"></a>Run(&#34;fn1 := func() int { return 1; L: goto L }&#34;),
    <a id="L208"></a><span class="comment">// Scope checking</span>
    <a id="L209"></a>Run(&#34;fn1 := func() { { L: x:=1 } goto L }&#34;),
    <a id="L210"></a>CErr(&#34;fn1 := func() { { x:=1; L: } goto L }&#34;, &#34;into scope&#34;),
    <a id="L211"></a>CErr(&#34;fn1 := func() { goto L; x:=1; L: }&#34;, &#34;into scope&#34;),
    <a id="L212"></a>Run(&#34;fn1 := func() { goto L; { L: x:=1 } }&#34;),
    <a id="L213"></a>CErr(&#34;fn1 := func() { goto L; { x:=1; L: } }&#34;, &#34;into scope&#34;),

    <a id="L215"></a><span class="comment">// Blocks</span>
    <a id="L216"></a>CErr(&#34;fn1 := func() int {{}}&#34;, &#34;return&#34;),
    <a id="L217"></a>Val1(&#34;fn1 := func() bool { { return true } }; b := fn1()&#34;, &#34;b&#34;, true),

    <a id="L219"></a><span class="comment">// If</span>
    <a id="L220"></a>Val2(&#34;if true { i = 2 } else { i = 3 }; i2 = 4&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 4),
    <a id="L221"></a>Val2(&#34;if false { i = 2 } else { i = 3 }; i2 = 4&#34;, &#34;i&#34;, 3, &#34;i2&#34;, 4),
    <a id="L222"></a>Val2(&#34;if i == i2 { i = 2 } else { i = 3 }; i2 = 4&#34;, &#34;i&#34;, 3, &#34;i2&#34;, 4),
    <a id="L223"></a><span class="comment">// Omit optional parts</span>
    <a id="L224"></a>Val2(&#34;if { i = 2 } else { i = 3 }; i2 = 4&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 4),
    <a id="L225"></a>Val2(&#34;if true { i = 2 }; i2 = 4&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 4),
    <a id="L226"></a>Val2(&#34;if false { i = 2 }; i2 = 4&#34;, &#34;i&#34;, 1, &#34;i2&#34;, 4),
    <a id="L227"></a><span class="comment">// Init</span>
    <a id="L228"></a>Val2(&#34;if x := true; x { i = 2 } else { i = 3 }; i2 = 4&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 4),
    <a id="L229"></a>Val2(&#34;if x := false; x { i = 2 } else { i = 3 }; i2 = 4&#34;, &#34;i&#34;, 3, &#34;i2&#34;, 4),
    <a id="L230"></a><span class="comment">// Statement else</span>
    <a id="L231"></a>Val2(&#34;if true { i = 2 } else i = 3; i2 = 4&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 4),
    <a id="L232"></a>Val2(&#34;if false { i = 2 } else i = 3; i2 = 4&#34;, &#34;i&#34;, 3, &#34;i2&#34;, 4),
    <a id="L233"></a><span class="comment">// Scoping</span>
    <a id="L234"></a>Val2(&#34;if true { i := 2 } else { i := 3 }; i2 = i&#34;, &#34;i&#34;, 1, &#34;i2&#34;, 1),
    <a id="L235"></a>Val2(&#34;if false { i := 2 } else { i := 3 }; i2 = i&#34;, &#34;i&#34;, 1, &#34;i2&#34;, 1),
    <a id="L236"></a>Val2(&#34;if false { i := 2 } else i := 3; i2 = i&#34;, &#34;i&#34;, 1, &#34;i2&#34;, 1),
    <a id="L237"></a>CErr(&#34;if true { x := 2 }; x = 4&#34;, undefined),
    <a id="L238"></a>Val2(&#34;if i := 2; true { i2 = i; i := 3 }&#34;, &#34;i&#34;, 1, &#34;i2&#34;, 2),
    <a id="L239"></a>Val2(&#34;if i := 2; false {} else { i2 = i; i := 3 }&#34;, &#34;i&#34;, 1, &#34;i2&#34;, 2),
    <a id="L240"></a><span class="comment">// Return checking</span>
    <a id="L241"></a>Run(&#34;fn1 := func() int { if true { return 1 } else { return 2 } }&#34;),
    <a id="L242"></a>Run(&#34;fn1 := func() int { if true { return 1 } else return 2 }&#34;),
    <a id="L243"></a>CErr(&#34;fn1 := func() int { if true { return 1 } else { } }&#34;, &#34;return&#34;),
    <a id="L244"></a>CErr(&#34;fn1 := func() int { if true { } else { return 1 } }&#34;, &#34;return&#34;),
    <a id="L245"></a>CErr(&#34;fn1 := func() int { if true { } else return 1 }&#34;, &#34;return&#34;),
    <a id="L246"></a>CErr(&#34;fn1 := func() int { if true { } else { } }&#34;, &#34;return&#34;),
    <a id="L247"></a>CErr(&#34;fn1 := func() int { if true { return 1 } }&#34;, &#34;return&#34;),
    <a id="L248"></a>CErr(&#34;fn1 := func() int { if true { } }&#34;, &#34;return&#34;),
    <a id="L249"></a>Run(&#34;fn1 := func() int { if true { }; return 1 }&#34;),
    <a id="L250"></a>CErr(&#34;fn1 := func() int { if { } }&#34;, &#34;return&#34;),
    <a id="L251"></a>CErr(&#34;fn1 := func() int { if { } else { return 2 } }&#34;, &#34;return&#34;),
    <a id="L252"></a>Run(&#34;fn1 := func() int { if { return 1 } }&#34;),
    <a id="L253"></a>Run(&#34;fn1 := func() int { if { return 1 } else { } }&#34;),
    <a id="L254"></a>Run(&#34;fn1 := func() int { if { return 1 } else { } }&#34;),

    <a id="L256"></a><span class="comment">// Switch</span>
    <a id="L257"></a>Val1(&#34;switch { case false: i += 2; case true: i += 4; default: i += 8 }&#34;, &#34;i&#34;, 1+4),
    <a id="L258"></a>Val1(&#34;switch { default: i += 2; case false: i += 4; case true: i += 8 }&#34;, &#34;i&#34;, 1+8),
    <a id="L259"></a>CErr(&#34;switch { default: i += 2; default: i += 4 }&#34;, &#34;more than one&#34;),
    <a id="L260"></a>Val1(&#34;switch false { case false: i += 2; case true: i += 4; default: i += 8 }&#34;, &#34;i&#34;, 1+2),
    <a id="L261"></a>CErr(&#34;switch s { case 1: }&#34;, opTypes),
    <a id="L262"></a>CErr(&#34;switch ai { case ai: i += 2 }&#34;, opTypes),
    <a id="L263"></a>Val1(&#34;switch 1.0 { case 1: i += 2; case 2: i += 4 }&#34;, &#34;i&#34;, 1+2),
    <a id="L264"></a>Val1(&#34;switch 1.5 { case 1: i += 2; case 2: i += 4 }&#34;, &#34;i&#34;, 1),
    <a id="L265"></a>CErr(&#34;switch oneTwo() {}&#34;, &#34;multi-valued expression&#34;),
    <a id="L266"></a>Val1(&#34;switch 2 { case 1: i += 2; fallthrough; case 2: i += 4; fallthrough; case 3: i += 8; fallthrough }&#34;, &#34;i&#34;, 1+4+8),
    <a id="L267"></a>Val1(&#34;switch 5 { case 1: i += 2; fallthrough; default: i += 4; fallthrough; case 2: i += 8; fallthrough; case 3: i += 16; fallthrough }&#34;, &#34;i&#34;, 1+4+8+16),
    <a id="L268"></a>CErr(&#34;switch { case true: fallthrough; i += 2 }&#34;, &#34;final statement&#34;),
    <a id="L269"></a>Val1(&#34;switch { case true: i += 2; fallthrough; ; ; case false: i += 4 }&#34;, &#34;i&#34;, 1+2+4),
    <a id="L270"></a>Val1(&#34;switch 2 { case 0, 1: i += 2; case 2, 3: i += 4 }&#34;, &#34;i&#34;, 1+4),
    <a id="L271"></a>Val2(&#34;switch func()int{i2++;return 5}() { case 1, 2: i += 2; case 4, 5: i += 4 }&#34;, &#34;i&#34;, 1+4, &#34;i2&#34;, 3),
    <a id="L272"></a>Run(&#34;switch i { case i: }&#34;),
    <a id="L273"></a><span class="comment">// TODO(austin) Why doesn&#39;t this fail?</span>
    <a id="L274"></a><span class="comment">//CErr(&#34;case 1:&#34;, &#34;XXX&#34;),</span>

    <a id="L276"></a><span class="comment">// For</span>
    <a id="L277"></a>Val2(&#34;for x := 1; x &lt; 5; x++ { i+=x }; i2 = 4&#34;, &#34;i&#34;, 11, &#34;i2&#34;, 4),
    <a id="L278"></a>Val2(&#34;for x := 1; x &lt; 5; x++ { i+=x; break; i++ }; i2 = 4&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 4),
    <a id="L279"></a>Val2(&#34;for x := 1; x &lt; 5; x++ { i+=x; continue; i++ }; i2 = 4&#34;, &#34;i&#34;, 11, &#34;i2&#34;, 4),
    <a id="L280"></a>Val2(&#34;for i = 2; false; i = 3 { i = 4 }; i2 = 4&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 4),
    <a id="L281"></a>Val2(&#34;for i &lt; 5 { i++ }; i2 = 4&#34;, &#34;i&#34;, 5, &#34;i2&#34;, 4),
    <a id="L282"></a>Val2(&#34;for i &lt; 0 { i++ }; i2 = 4&#34;, &#34;i&#34;, 1, &#34;i2&#34;, 4),
    <a id="L283"></a><span class="comment">// Scoping</span>
    <a id="L284"></a>Val2(&#34;for i := 2; true; { i2 = i; i := 3; break }&#34;, &#34;i&#34;, 1, &#34;i2&#34;, 2),
    <a id="L285"></a><span class="comment">// Labeled break/continue</span>
    <a id="L286"></a>Val1(&#34;L1: for { L2: for { i+=2; break L1; i+=4 } i+=8 }&#34;, &#34;i&#34;, 1+2),
    <a id="L287"></a>Val1(&#34;L1: for { L2: for { i+=2; break L2; i+=4 } i+=8; break; i+=16 }&#34;, &#34;i&#34;, 1+2+8),
    <a id="L288"></a>CErr(&#34;L1: { for { break L1 } }&#34;, &#34;break.*not defined&#34;),
    <a id="L289"></a>CErr(&#34;L1: for {} for { break L1 }&#34;, &#34;break.*not defined&#34;),
    <a id="L290"></a>CErr(&#34;L1:; for { break L1 }&#34;, &#34;break.*not defined&#34;),
    <a id="L291"></a>Val2(&#34;L1: for i = 0; i &lt; 2; i++ { L2: for { i2++; continue L1; i2++ } }&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 4),
    <a id="L292"></a>CErr(&#34;L1: { for { continue L1 } }&#34;, &#34;continue.*not defined&#34;),
    <a id="L293"></a>CErr(&#34;L1:; for { continue L1 }&#34;, &#34;continue.*not defined&#34;),
    <a id="L294"></a><span class="comment">// Return checking</span>
    <a id="L295"></a>Run(&#34;fn1 := func() int{ for {} }&#34;),
    <a id="L296"></a>CErr(&#34;fn1 := func() int{ for true {} }&#34;, &#34;return&#34;),
    <a id="L297"></a>CErr(&#34;fn1 := func() int{ for true {return 1} }&#34;, &#34;return&#34;),
    <a id="L298"></a>CErr(&#34;fn1 := func() int{ for {break} }&#34;, &#34;return&#34;),
    <a id="L299"></a>Run(&#34;fn1 := func() int{ for { for {break} } }&#34;),
    <a id="L300"></a>CErr(&#34;fn1 := func() int{ L1: for { for {break L1} } }&#34;, &#34;return&#34;),
    <a id="L301"></a>Run(&#34;fn1 := func() int{ for true {} return 1 }&#34;),

    <a id="L303"></a><span class="comment">// Selectors</span>
    <a id="L304"></a>Val1(&#34;var x struct { a int; b int }; x.a = 42; i = x.a&#34;, &#34;i&#34;, 42),
    <a id="L305"></a>Val1(&#34;type T struct { x int }; var y struct { T }; y.x = 42; i = y.x&#34;, &#34;i&#34;, 42),
    <a id="L306"></a>Val2(&#34;type T struct { x int }; var y struct { T; x int }; y.x = 42; i = y.x; i2 = y.T.x&#34;, &#34;i&#34;, 42, &#34;i2&#34;, 0),
    <a id="L307"></a>Run(&#34;type T struct { x int }; var y struct { *T }; a := func(){i=y.x}&#34;),
    <a id="L308"></a>CErr(&#34;type T struct { x int }; var x T; x.y = 42&#34;, &#34;no field&#34;),
    <a id="L309"></a>CErr(&#34;type T struct { x int }; type U struct { x int }; var y struct { T; U }; y.x = 42&#34;, &#34;ambiguous.*\tT\\.x\n\tU\\.x&#34;),
    <a id="L310"></a>CErr(&#34;type T struct { *T }; var x T; x.foo&#34;, &#34;no field&#34;),

    <a id="L312"></a>Val1(&#34;fib := func(int) int{return 0;}; fib = func(v int) int { if v &lt; 2 { return 1 } return fib(v-1)+fib(v-2) }; i = fib(20)&#34;, &#34;i&#34;, 10946),

    <a id="L314"></a><span class="comment">// Make slice</span>
    <a id="L315"></a>Val2(&#34;x := make([]int, 2); x[0] = 42; i, i2 = x[0], x[1]&#34;, &#34;i&#34;, 42, &#34;i2&#34;, 0),
    <a id="L316"></a>Val2(&#34;x := make([]int, 2); x[1] = 42; i, i2 = x[0], x[1]&#34;, &#34;i&#34;, 0, &#34;i2&#34;, 42),
    <a id="L317"></a>RErr(&#34;x := make([]int, 2); x[-i] = 42&#34;, &#34;negative index&#34;),
    <a id="L318"></a>RErr(&#34;x := make([]int, 2); x[2] = 42&#34;, &#34;index 2 exceeds&#34;),
    <a id="L319"></a>Val2(&#34;x := make([]int, 2, 3); i, i2 = len(x), cap(x)&#34;, &#34;i&#34;, 2, &#34;i2&#34;, 3),
    <a id="L320"></a>Val2(&#34;x := make([]int, 3, 2); i, i2 = len(x), cap(x)&#34;, &#34;i&#34;, 3, &#34;i2&#34;, 3),
    <a id="L321"></a>RErr(&#34;x := make([]int, -i)&#34;, &#34;negative length&#34;),
    <a id="L322"></a>RErr(&#34;x := make([]int, 2, -i)&#34;, &#34;negative capacity&#34;),
    <a id="L323"></a>RErr(&#34;x := make([]int, 2, 3); x[2] = 42&#34;, &#34;index 2 exceeds&#34;),
    <a id="L324"></a>CErr(&#34;x := make([]int, 2, 3, 4)&#34;, &#34;too many&#34;),
    <a id="L325"></a>CErr(&#34;x := make([]int)&#34;, &#34;not enough&#34;),

    <a id="L327"></a><span class="comment">// TODO(austin) Test make map</span>

    <a id="L329"></a><span class="comment">// Maps</span>
    <a id="L330"></a>Val1(&#34;x := make(map[int] int); x[1] = 42; i = x[1]&#34;, &#34;i&#34;, 42),
    <a id="L331"></a>Val2(&#34;x := make(map[int] int); x[1] = 42; i, y := x[1]&#34;, &#34;i&#34;, 42, &#34;y&#34;, true),
    <a id="L332"></a>Val2(&#34;x := make(map[int] int); x[1] = 42; i, y := x[2]&#34;, &#34;i&#34;, 0, &#34;y&#34;, false),
    <a id="L333"></a><span class="comment">// Not implemented</span>
    <a id="L334"></a><span class="comment">//Val1(&#34;x := make(map[int] int); x[1] = 42, true; i = x[1]&#34;, &#34;i&#34;, 42),</span>
    <a id="L335"></a><span class="comment">//Val2(&#34;x := make(map[int] int); x[1] = 42; x[1] = 42, false; i, y := x[1]&#34;, &#34;i&#34;, 0, &#34;y&#34;, false),</span>
    <a id="L336"></a>Run(&#34;var x int; a := make(map[int] int); a[0], x = 1, 2&#34;),
    <a id="L337"></a>CErr(&#34;x := make(map[int] int); (func(a,b int){})(x[0])&#34;, &#34;not enough&#34;),
    <a id="L338"></a>CErr(&#34;x := make(map[int] int); x[1] = oneTwo()&#34;, &#34;too many&#34;),
    <a id="L339"></a>RErr(&#34;x := make(map[int] int); i = x[1]&#34;, &#34;key &#39;1&#39; not found&#34;),

    <a id="L341"></a><span class="comment">// Functions</span>
    <a id="L342"></a>Val2(&#34;func fib(n int) int { if n &lt;= 2 { return n } return fib(n-1) + fib(n-2) }&#34;, &#34;fib(4)&#34;, 5, &#34;fib(10)&#34;, 89),
    <a id="L343"></a>Run(&#34;func f1(){}&#34;),
    <a id="L344"></a>Run2(&#34;func f1(){}&#34;, &#34;f1()&#34;),
<a id="L345"></a>}

<a id="L347"></a>func TestStmt(t *testing.T) { runTests(t, &#34;stmtTests&#34;, stmtTests) }
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
