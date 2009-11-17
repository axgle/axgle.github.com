<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/ebnf/ebnf.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/ebnf/ebnf.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// A library for EBNF grammars. The input is text ([]byte) satisfying</span>
<a id="L6"></a><span class="comment">// the following grammar (represented itself in EBNF):</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a><span class="comment">//	Production  = name &#34;=&#34; Expression &#34;.&#34; .</span>
<a id="L9"></a><span class="comment">//	Expression  = Alternative { &#34;|&#34; Alternative } .</span>
<a id="L10"></a><span class="comment">//	Alternative = Term { Term } .</span>
<a id="L11"></a><span class="comment">//	Term        = name | token [ &#34;...&#34; token ] | Group | Option | Repetition .</span>
<a id="L12"></a><span class="comment">//	Group       = &#34;(&#34; Expression &#34;)&#34; .</span>
<a id="L13"></a><span class="comment">//	Option      = &#34;[&#34; Expression &#34;]&#34; .</span>
<a id="L14"></a><span class="comment">//	Repetition  = &#34;{&#34; Expression &#34;}&#34; .</span>
<a id="L15"></a><span class="comment">//</span>
<a id="L16"></a><span class="comment">// A name is a Go identifier, a token is a Go string, and comments</span>
<a id="L17"></a><span class="comment">// and white space follow the same rules as for the Go language.</span>
<a id="L18"></a><span class="comment">// Production names starting with an uppercase Unicode letter denote</span>
<a id="L19"></a><span class="comment">// non-terminal productions (i.e., productions which allow white-space</span>
<a id="L20"></a><span class="comment">// and comments between tokens); all other production names denote</span>
<a id="L21"></a><span class="comment">// lexical productions.</span>
<a id="L22"></a><span class="comment">//</span>
<a id="L23"></a>package ebnf

<a id="L25"></a>import (
    <a id="L26"></a>&#34;container/vector&#34;;
    <a id="L27"></a>&#34;go/scanner&#34;;
    <a id="L28"></a>&#34;go/token&#34;;
    <a id="L29"></a>&#34;os&#34;;
    <a id="L30"></a>&#34;unicode&#34;;
    <a id="L31"></a>&#34;utf8&#34;;
<a id="L32"></a>)


<a id="L35"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L36"></a><span class="comment">// Internal representation</span>

<a id="L38"></a>type (
    <a id="L39"></a><span class="comment">// An Expression node represents a production expression.</span>
    <a id="L40"></a>Expression interface {
        <a id="L41"></a><span class="comment">// Pos is the position of the first character of the syntactic construct</span>
        <a id="L42"></a>Pos() token.Position;
    <a id="L43"></a>};

    <a id="L45"></a><span class="comment">// An Alternative node represents a non-empty list of alternative expressions.</span>
    <a id="L46"></a>Alternative []Expression; <span class="comment">// x | y | z</span>

    <a id="L48"></a><span class="comment">// A Sequence node represents a non-empty list of sequential expressions.</span>
    <a id="L49"></a>Sequence []Expression; <span class="comment">// x y z</span>

    <a id="L51"></a><span class="comment">// A Name node represents a production name.</span>
    <a id="L52"></a>Name struct {
        <a id="L53"></a>token.Position;
        <a id="L54"></a>String string;
    <a id="L55"></a>};

    <a id="L57"></a><span class="comment">// A Token node represents a literal.</span>
    <a id="L58"></a>Token struct {
        <a id="L59"></a>token.Position;
        <a id="L60"></a>String string;
    <a id="L61"></a>};

    <a id="L63"></a><span class="comment">// A List node represents a range of characters.</span>
    <a id="L64"></a>Range struct {
        <a id="L65"></a>Begin, End *Token; <span class="comment">// begin ... end</span>
    <a id="L66"></a>};

    <a id="L68"></a><span class="comment">// A Group node represents a grouped expression.</span>
    <a id="L69"></a>Group struct {
        <a id="L70"></a>token.Position;
        <a id="L71"></a>Body Expression; <span class="comment">// (body)</span>
    <a id="L72"></a>};

    <a id="L74"></a><span class="comment">// An Option node represents an optional expression.</span>
    <a id="L75"></a>Option struct {
        <a id="L76"></a>token.Position;
        <a id="L77"></a>Body Expression; <span class="comment">// [body]</span>
    <a id="L78"></a>};

    <a id="L80"></a><span class="comment">// A Repetition node represents a repeated expression.</span>
    <a id="L81"></a>Repetition struct {
        <a id="L82"></a>token.Position;
        <a id="L83"></a>Body Expression; <span class="comment">// {body}</span>
    <a id="L84"></a>};

    <a id="L86"></a><span class="comment">// A Production node represents an EBNF production.</span>
    <a id="L87"></a>Production struct {
        <a id="L88"></a>Name *Name;
        <a id="L89"></a>Expr Expression;
    <a id="L90"></a>};

    <a id="L92"></a><span class="comment">// A Grammar is a set of EBNF productions. The map</span>
    <a id="L93"></a><span class="comment">// is indexed by production name.</span>
    <a id="L94"></a><span class="comment">//</span>
    <a id="L95"></a>Grammar map[string]*Production;
<a id="L96"></a>)


<a id="L99"></a>func (x Alternative) Pos() token.Position {
    <a id="L100"></a>return x[0].Pos() <span class="comment">// the parser always generates non-empty Alternative</span>
<a id="L101"></a>}


<a id="L104"></a>func (x Sequence) Pos() token.Position {
    <a id="L105"></a>return x[0].Pos() <span class="comment">// the parser always generates non-empty Sequences</span>
<a id="L106"></a>}


<a id="L109"></a>func (x Range) Pos() token.Position { return x.Begin.Pos() }


<a id="L112"></a>func (p *Production) Pos() token.Position { return p.Name.Pos() }


<a id="L115"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L116"></a><span class="comment">// Grammar verification</span>

<a id="L118"></a>func isLexical(name string) bool {
    <a id="L119"></a>ch, _ := utf8.DecodeRuneInString(name);
    <a id="L120"></a>return !unicode.IsUpper(ch);
<a id="L121"></a>}


<a id="L124"></a>type verifier struct {
    <a id="L125"></a>scanner.ErrorVector;
    <a id="L126"></a>worklist vector.Vector;
    <a id="L127"></a>reached  Grammar; <span class="comment">// set of productions reached from (and including) the root production</span>
    <a id="L128"></a>grammar  Grammar;
<a id="L129"></a>}


<a id="L132"></a>func (v *verifier) push(prod *Production) {
    <a id="L133"></a>name := prod.Name.String;
    <a id="L134"></a>if _, found := v.reached[name]; !found {
        <a id="L135"></a>v.worklist.Push(prod);
        <a id="L136"></a>v.reached[name] = prod;
    <a id="L137"></a>}
<a id="L138"></a>}


<a id="L141"></a>func (v *verifier) verifyChar(x *Token) int {
    <a id="L142"></a>s := x.String;
    <a id="L143"></a>if utf8.RuneCountInString(s) != 1 {
        <a id="L144"></a>v.Error(x.Pos(), &#34;single char expected, found &#34;+s);
        <a id="L145"></a>return 0;
    <a id="L146"></a>}
    <a id="L147"></a>ch, _ := utf8.DecodeRuneInString(s);
    <a id="L148"></a>return ch;
<a id="L149"></a>}


<a id="L152"></a>func (v *verifier) verifyExpr(expr Expression, lexical bool) {
    <a id="L153"></a>switch x := expr.(type) {
    <a id="L154"></a>case nil:
        <a id="L155"></a><span class="comment">// empty expression</span>
    <a id="L156"></a>case Alternative:
        <a id="L157"></a>for _, e := range x {
            <a id="L158"></a>v.verifyExpr(e, lexical)
        <a id="L159"></a>}
    <a id="L160"></a>case Sequence:
        <a id="L161"></a>for _, e := range x {
            <a id="L162"></a>v.verifyExpr(e, lexical)
        <a id="L163"></a>}
    <a id="L164"></a>case *Name:
        <a id="L165"></a><span class="comment">// a production with this name must exist;</span>
        <a id="L166"></a><span class="comment">// add it to the worklist if not yet processed</span>
        <a id="L167"></a>if prod, found := v.grammar[x.String]; found {
            <a id="L168"></a>v.push(prod)
        <a id="L169"></a>} else {
            <a id="L170"></a>v.Error(x.Pos(), &#34;missing production &#34;+x.String)
        <a id="L171"></a>}
        <a id="L172"></a><span class="comment">// within a lexical production references</span>
        <a id="L173"></a><span class="comment">// to non-lexical productions are invalid</span>
        <a id="L174"></a>if lexical &amp;&amp; !isLexical(x.String) {
            <a id="L175"></a>v.Error(x.Pos(), &#34;reference to non-lexical production &#34;+x.String)
        <a id="L176"></a>}
    <a id="L177"></a>case *Token:
        <a id="L178"></a><span class="comment">// nothing to do for now</span>
    <a id="L179"></a>case *Range:
        <a id="L180"></a>i := v.verifyChar(x.Begin);
        <a id="L181"></a>j := v.verifyChar(x.End);
        <a id="L182"></a>if i &gt;= j {
            <a id="L183"></a>v.Error(x.Pos(), &#34;decreasing character range&#34;)
        <a id="L184"></a>}
    <a id="L185"></a>case *Group:
        <a id="L186"></a>v.verifyExpr(x.Body, lexical)
    <a id="L187"></a>case *Option:
        <a id="L188"></a>v.verifyExpr(x.Body, lexical)
    <a id="L189"></a>case *Repetition:
        <a id="L190"></a>v.verifyExpr(x.Body, lexical)
    <a id="L191"></a>default:
        <a id="L192"></a>panic(&#34;unreachable&#34;)
    <a id="L193"></a>}
<a id="L194"></a>}


<a id="L197"></a>func (v *verifier) verify(grammar Grammar, start string) {
    <a id="L198"></a><span class="comment">// find root production</span>
    <a id="L199"></a>root, found := grammar[start];
    <a id="L200"></a>if !found {
        <a id="L201"></a>var noPos token.Position;
        <a id="L202"></a>v.Error(noPos, &#34;no start production &#34;+start);
        <a id="L203"></a>return;
    <a id="L204"></a>}

    <a id="L206"></a><span class="comment">// initialize verifier</span>
    <a id="L207"></a>v.ErrorVector.Init();
    <a id="L208"></a>v.worklist.Init(0);
    <a id="L209"></a>v.reached = make(Grammar);
    <a id="L210"></a>v.grammar = grammar;

    <a id="L212"></a><span class="comment">// work through the worklist</span>
    <a id="L213"></a>v.push(root);
    <a id="L214"></a>for v.worklist.Len() &gt; 0 {
        <a id="L215"></a>prod := v.worklist.Pop().(*Production);
        <a id="L216"></a>v.verifyExpr(prod.Expr, isLexical(prod.Name.String));
    <a id="L217"></a>}

    <a id="L219"></a><span class="comment">// check if all productions were reached</span>
    <a id="L220"></a>if len(v.reached) &lt; len(v.grammar) {
        <a id="L221"></a>for name, prod := range v.grammar {
            <a id="L222"></a>if _, found := v.reached[name]; !found {
                <a id="L223"></a>v.Error(prod.Pos(), name+&#34; is unreachable&#34;)
            <a id="L224"></a>}
        <a id="L225"></a>}
    <a id="L226"></a>}
<a id="L227"></a>}


<a id="L230"></a><span class="comment">// Verify checks that:</span>
<a id="L231"></a><span class="comment">//	- all productions used are defined</span>
<a id="L232"></a><span class="comment">//	- all productions defined are used when beginning at start</span>
<a id="L233"></a><span class="comment">//	- lexical productions refer only to other lexical productions</span>
<a id="L234"></a><span class="comment">//</span>
<a id="L235"></a>func Verify(grammar Grammar, start string) os.Error {
    <a id="L236"></a>var v verifier;
    <a id="L237"></a>v.verify(grammar, start);
    <a id="L238"></a>return v.GetError(scanner.Sorted);
<a id="L239"></a>}
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
