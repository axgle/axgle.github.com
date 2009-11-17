<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/parser/parser.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/go/parser/parser.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// A parser for Go source files. Input may be provided in a variety of</span>
<a id="L6"></a><span class="comment">// forms (see the various Parse* functions); the output is an abstract</span>
<a id="L7"></a><span class="comment">// syntax tree (AST) representing the Go source. The parser is invoked</span>
<a id="L8"></a><span class="comment">// through one of the Parse* functions.</span>
<a id="L9"></a><span class="comment">//</span>
<a id="L10"></a>package parser

<a id="L12"></a>import (
    <a id="L13"></a>&#34;container/vector&#34;;
    <a id="L14"></a>&#34;fmt&#34;;
    <a id="L15"></a>&#34;go/ast&#34;;
    <a id="L16"></a>&#34;go/scanner&#34;;
    <a id="L17"></a>&#34;go/token&#34;;
<a id="L18"></a>)


<a id="L21"></a><span class="comment">// noPos is used when there is no corresponding source position for a token.</span>
<a id="L22"></a>var noPos token.Position


<a id="L25"></a><span class="comment">// The mode parameter to the Parse* functions is a set of flags (or 0).</span>
<a id="L26"></a><span class="comment">// They control the amount of source code parsed and other optional</span>
<a id="L27"></a><span class="comment">// parser functionality.</span>
<a id="L28"></a><span class="comment">//</span>
<a id="L29"></a>const (
    <a id="L30"></a>PackageClauseOnly uint = 1 &lt;&lt; iota; <span class="comment">// parsing stops after package clause</span>
    <a id="L31"></a>ImportsOnly;           <span class="comment">// parsing stops after import declarations</span>
    <a id="L32"></a>ParseComments;         <span class="comment">// parse comments and add them to AST</span>
    <a id="L33"></a>Trace;                 <span class="comment">// print a trace of parsed productions</span>
<a id="L34"></a>)


<a id="L37"></a><span class="comment">// The parser structure holds the parser&#39;s internal state.</span>
<a id="L38"></a>type parser struct {
    <a id="L39"></a>scanner.ErrorVector;
    <a id="L40"></a>scanner scanner.Scanner;

    <a id="L42"></a><span class="comment">// Tracing/debugging</span>
    <a id="L43"></a>mode   uint; <span class="comment">// parsing mode</span>
    <a id="L44"></a>trace  bool; <span class="comment">// == (mode &amp; Trace != 0)</span>
    <a id="L45"></a>indent uint; <span class="comment">// indentation used for tracing output</span>

    <a id="L47"></a><span class="comment">// Comments</span>
    <a id="L48"></a>comments    *ast.CommentGroup; <span class="comment">// list of collected comments</span>
    <a id="L49"></a>lastComment *ast.CommentGroup; <span class="comment">// last comment in the comments list</span>
    <a id="L50"></a>leadComment *ast.CommentGroup; <span class="comment">// the last lead comment</span>
    <a id="L51"></a>lineComment *ast.CommentGroup; <span class="comment">// the last line comment</span>

    <a id="L53"></a><span class="comment">// Next token</span>
    <a id="L54"></a>pos token.Position; <span class="comment">// token position</span>
    <a id="L55"></a>tok token.Token;    <span class="comment">// one token look-ahead</span>
    <a id="L56"></a>lit []byte;         <span class="comment">// token literal</span>

    <a id="L58"></a><span class="comment">// Non-syntactic parser control</span>
    <a id="L59"></a>optSemi bool; <span class="comment">// true if semicolon separator is optional in statement list</span>
    <a id="L60"></a>exprLev int;  <span class="comment">// &lt; 0: in control clause, &gt;= 0: in expression</span>

    <a id="L62"></a><span class="comment">// Scopes</span>
    <a id="L63"></a>pkgScope  *ast.Scope;
    <a id="L64"></a>fileScope *ast.Scope;
    <a id="L65"></a>topScope  *ast.Scope;
<a id="L66"></a>}


<a id="L69"></a><span class="comment">// scannerMode returns the scanner mode bits given the parser&#39;s mode bits.</span>
<a id="L70"></a>func scannerMode(mode uint) uint {
    <a id="L71"></a>if mode&amp;ParseComments != 0 {
        <a id="L72"></a>return scanner.ScanComments
    <a id="L73"></a>}
    <a id="L74"></a>return 0;
<a id="L75"></a>}


<a id="L78"></a>func (p *parser) init(filename string, src []byte, mode uint) {
    <a id="L79"></a>p.ErrorVector.Init();
    <a id="L80"></a>p.scanner.Init(filename, src, p, scannerMode(mode));
    <a id="L81"></a>p.mode = mode;
    <a id="L82"></a>p.trace = mode&amp;Trace != 0; <span class="comment">// for convenience (p.trace is used frequently)</span>
    <a id="L83"></a>p.next();
<a id="L84"></a>}


<a id="L87"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L88"></a><span class="comment">// Parsing support</span>

<a id="L90"></a>func (p *parser) printTrace(a ...) {
    <a id="L91"></a>const dots = &#34;. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . &#34;
        <a id="L92"></a>&#34;. . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . &#34;;
    <a id="L93"></a>const n = uint(len(dots));
    <a id="L94"></a>fmt.Printf(&#34;%5d:%3d: &#34;, p.pos.Line, p.pos.Column);
    <a id="L95"></a>i := 2 * p.indent;
    <a id="L96"></a>for ; i &gt; n; i -= n {
        <a id="L97"></a>fmt.Print(dots)
    <a id="L98"></a>}
    <a id="L99"></a>fmt.Print(dots[0:i]);
    <a id="L100"></a>fmt.Println(a);
<a id="L101"></a>}


<a id="L104"></a>func trace(p *parser, msg string) *parser {
    <a id="L105"></a>p.printTrace(msg, &#34;(&#34;);
    <a id="L106"></a>p.indent++;
    <a id="L107"></a>return p;
<a id="L108"></a>}


<a id="L111"></a><span class="comment">// Usage pattern: defer un(trace(p, &#34;...&#34;));</span>
<a id="L112"></a>func un(p *parser) {
    <a id="L113"></a>p.indent--;
    <a id="L114"></a>p.printTrace(&#34;)&#34;);
<a id="L115"></a>}


<a id="L118"></a><span class="comment">// Advance to the next token.</span>
<a id="L119"></a>func (p *parser) next0() {
    <a id="L120"></a><span class="comment">// Because of one-token look-ahead, print the previous token</span>
    <a id="L121"></a><span class="comment">// when tracing as it provides a more readable output. The</span>
    <a id="L122"></a><span class="comment">// very first token (p.pos.Line == 0) is not initialized (it</span>
    <a id="L123"></a><span class="comment">// is token.ILLEGAL), so don&#39;t print it .</span>
    <a id="L124"></a>if p.trace &amp;&amp; p.pos.Line &gt; 0 {
        <a id="L125"></a>s := p.tok.String();
        <a id="L126"></a>switch {
        <a id="L127"></a>case p.tok.IsLiteral():
            <a id="L128"></a>p.printTrace(s, string(p.lit))
        <a id="L129"></a>case p.tok.IsOperator(), p.tok.IsKeyword():
            <a id="L130"></a>p.printTrace(&#34;\&#34;&#34; + s + &#34;\&#34;&#34;)
        <a id="L131"></a>default:
            <a id="L132"></a>p.printTrace(s)
        <a id="L133"></a>}
    <a id="L134"></a>}

    <a id="L136"></a>p.pos, p.tok, p.lit = p.scanner.Scan();
    <a id="L137"></a>p.optSemi = false;
<a id="L138"></a>}


<a id="L141"></a><span class="comment">// Consume a comment and return it and the line on which it ends.</span>
<a id="L142"></a>func (p *parser) consumeComment() (comment *ast.Comment, endline int) {
    <a id="L143"></a><span class="comment">// /*-style comments may end on a different line than where they start.</span>
    <a id="L144"></a><span class="comment">// Scan the comment for &#39;\n&#39; chars and adjust endline accordingly.</span>
    <a id="L145"></a>endline = p.pos.Line;
    <a id="L146"></a>if p.lit[1] == &#39;*&#39; {
        <a id="L147"></a>for _, b := range p.lit {
            <a id="L148"></a>if b == &#39;\n&#39; {
                <a id="L149"></a>endline++
            <a id="L150"></a>}
        <a id="L151"></a>}
    <a id="L152"></a>}

    <a id="L154"></a>comment = &amp;ast.Comment{p.pos, p.lit};
    <a id="L155"></a>p.next0();

    <a id="L157"></a>return;
<a id="L158"></a>}


<a id="L161"></a><span class="comment">// Consume a group of adjacent comments, add it to the parser&#39;s</span>
<a id="L162"></a><span class="comment">// comments list, and return the line of which the last comment</span>
<a id="L163"></a><span class="comment">// in the group ends. An empty line or non-comment token terminates</span>
<a id="L164"></a><span class="comment">// a comment group.</span>
<a id="L165"></a><span class="comment">//</span>
<a id="L166"></a>func (p *parser) consumeCommentGroup() int {
    <a id="L167"></a>list := vector.New(0);
    <a id="L168"></a>endline := p.pos.Line;
    <a id="L169"></a>for p.tok == token.COMMENT &amp;&amp; endline+1 &gt;= p.pos.Line {
        <a id="L170"></a>var comment *ast.Comment;
        <a id="L171"></a>comment, endline = p.consumeComment();
        <a id="L172"></a>list.Push(comment);
    <a id="L173"></a>}

    <a id="L175"></a><span class="comment">// convert list</span>
    <a id="L176"></a>group := make([]*ast.Comment, list.Len());
    <a id="L177"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L178"></a>group[i] = list.At(i).(*ast.Comment)
    <a id="L179"></a>}

    <a id="L181"></a><span class="comment">// add comment group to the comments list</span>
    <a id="L182"></a>g := &amp;ast.CommentGroup{group, nil};
    <a id="L183"></a>if p.lastComment != nil {
        <a id="L184"></a>p.lastComment.Next = g
    <a id="L185"></a>} else {
        <a id="L186"></a>p.comments = g
    <a id="L187"></a>}
    <a id="L188"></a>p.lastComment = g;

    <a id="L190"></a>return endline;
<a id="L191"></a>}


<a id="L194"></a><span class="comment">// Advance to the next non-comment token. In the process, collect</span>
<a id="L195"></a><span class="comment">// any comment groups encountered, and remember the last lead and</span>
<a id="L196"></a><span class="comment">// and line comments.</span>
<a id="L197"></a><span class="comment">//</span>
<a id="L198"></a><span class="comment">// A lead comment is a comment group that starts and ends in a</span>
<a id="L199"></a><span class="comment">// line without any other tokens and that is followed by a non-comment</span>
<a id="L200"></a><span class="comment">// token on the line immediately after the comment group.</span>
<a id="L201"></a><span class="comment">//</span>
<a id="L202"></a><span class="comment">// A line comment is a comment group that follows a non-comment</span>
<a id="L203"></a><span class="comment">// token on the same line, and that has no tokens after it on the line</span>
<a id="L204"></a><span class="comment">// where it ends.</span>
<a id="L205"></a><span class="comment">//</span>
<a id="L206"></a><span class="comment">// Lead and line comments may be considered documentation that is</span>
<a id="L207"></a><span class="comment">// stored in the AST.</span>
<a id="L208"></a><span class="comment">//</span>
<a id="L209"></a>func (p *parser) next() {
    <a id="L210"></a>p.leadComment = nil;
    <a id="L211"></a>p.lineComment = nil;
    <a id="L212"></a>line := p.pos.Line; <span class="comment">// current line</span>
    <a id="L213"></a>p.next0();

    <a id="L215"></a>if p.tok == token.COMMENT {
        <a id="L216"></a>if p.pos.Line == line {
            <a id="L217"></a><span class="comment">// The comment is on same line as previous token; it</span>
            <a id="L218"></a><span class="comment">// cannot be a lead comment but may be a line comment.</span>
            <a id="L219"></a>endline := p.consumeCommentGroup();
            <a id="L220"></a>if p.pos.Line != endline {
                <a id="L221"></a><span class="comment">// The next token is on a different line, thus</span>
                <a id="L222"></a><span class="comment">// the last comment group is a line comment.</span>
                <a id="L223"></a>p.lineComment = p.lastComment
            <a id="L224"></a>}
        <a id="L225"></a>}

        <a id="L227"></a><span class="comment">// consume successor comments, if any</span>
        <a id="L228"></a>endline := -1;
        <a id="L229"></a>for p.tok == token.COMMENT {
            <a id="L230"></a>endline = p.consumeCommentGroup()
        <a id="L231"></a>}

        <a id="L233"></a>if endline &gt;= 0 &amp;&amp; endline+1 == p.pos.Line {
            <a id="L234"></a><span class="comment">// The next token is following on the line immediately after the</span>
            <a id="L235"></a><span class="comment">// comment group, thus the last comment group is a lead comment.</span>
            <a id="L236"></a>p.leadComment = p.lastComment
        <a id="L237"></a>}
    <a id="L238"></a>}
<a id="L239"></a>}


<a id="L242"></a>func (p *parser) errorExpected(pos token.Position, msg string) {
    <a id="L243"></a>msg = &#34;expected &#34; + msg;
    <a id="L244"></a>if pos.Offset == p.pos.Offset {
        <a id="L245"></a><span class="comment">// the error happened at the current position;</span>
        <a id="L246"></a><span class="comment">// make the error message more specific</span>
        <a id="L247"></a>msg += &#34;, found &#39;&#34; + p.tok.String() + &#34;&#39;&#34;;
        <a id="L248"></a>if p.tok.IsLiteral() {
            <a id="L249"></a>msg += &#34; &#34; + string(p.lit)
        <a id="L250"></a>}
    <a id="L251"></a>}
    <a id="L252"></a>p.Error(pos, msg);
<a id="L253"></a>}


<a id="L256"></a>func (p *parser) expect(tok token.Token) token.Position {
    <a id="L257"></a>pos := p.pos;
    <a id="L258"></a>if p.tok != tok {
        <a id="L259"></a>p.errorExpected(pos, &#34;&#39;&#34;+tok.String()+&#34;&#39;&#34;)
    <a id="L260"></a>}
    <a id="L261"></a>p.next(); <span class="comment">// make progress in any case</span>
    <a id="L262"></a>return pos;
<a id="L263"></a>}


<a id="L266"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L267"></a><span class="comment">// Scope support</span>

<a id="L269"></a>func openScope(p *parser) *parser {
    <a id="L270"></a>p.topScope = ast.NewScope(p.topScope);
    <a id="L271"></a>return p;
<a id="L272"></a>}


<a id="L275"></a><span class="comment">// Usage pattern: defer close(openScope(p));</span>
<a id="L276"></a>func close(p *parser) { p.topScope = p.topScope.Outer }


<a id="L279"></a>func (p *parser) declare(ident *ast.Ident) {
    <a id="L280"></a>if !p.topScope.Declare(ident) {
        <a id="L281"></a>p.Error(p.pos, &#34;&#39;&#34;+ident.Value+&#34;&#39; declared already&#34;)
    <a id="L282"></a>}
<a id="L283"></a>}


<a id="L286"></a>func (p *parser) declareList(idents []*ast.Ident) {
    <a id="L287"></a>for _, ident := range idents {
        <a id="L288"></a>p.declare(ident)
    <a id="L289"></a>}
<a id="L290"></a>}


<a id="L293"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L294"></a><span class="comment">// Common productions</span>

<a id="L296"></a>func (p *parser) parseIdent() *ast.Ident {
    <a id="L297"></a>if p.tok == token.IDENT {
        <a id="L298"></a>x := &amp;ast.Ident{p.pos, string(p.lit)};
        <a id="L299"></a>p.next();
        <a id="L300"></a>return x;
    <a id="L301"></a>}
    <a id="L302"></a>p.expect(token.IDENT); <span class="comment">// use expect() error handling</span>
    <a id="L303"></a>return &amp;ast.Ident{p.pos, &#34;&#34;};
<a id="L304"></a>}


<a id="L307"></a>func (p *parser) parseIdentList() []*ast.Ident {
    <a id="L308"></a>if p.trace {
        <a id="L309"></a>defer un(trace(p, &#34;IdentList&#34;))
    <a id="L310"></a>}

    <a id="L312"></a>list := vector.New(0);
    <a id="L313"></a>list.Push(p.parseIdent());
    <a id="L314"></a>for p.tok == token.COMMA {
        <a id="L315"></a>p.next();
        <a id="L316"></a>list.Push(p.parseIdent());
    <a id="L317"></a>}

    <a id="L319"></a><span class="comment">// convert vector</span>
    <a id="L320"></a>idents := make([]*ast.Ident, list.Len());
    <a id="L321"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L322"></a>idents[i] = list.At(i).(*ast.Ident)
    <a id="L323"></a>}

    <a id="L325"></a>return idents;
<a id="L326"></a>}


<a id="L329"></a>func (p *parser) parseExprList() []ast.Expr {
    <a id="L330"></a>if p.trace {
        <a id="L331"></a>defer un(trace(p, &#34;ExpressionList&#34;))
    <a id="L332"></a>}

    <a id="L334"></a>list := vector.New(0);
    <a id="L335"></a>list.Push(p.parseExpr());
    <a id="L336"></a>for p.tok == token.COMMA {
        <a id="L337"></a>p.next();
        <a id="L338"></a>list.Push(p.parseExpr());
    <a id="L339"></a>}

    <a id="L341"></a><span class="comment">// convert list</span>
    <a id="L342"></a>exprs := make([]ast.Expr, list.Len());
    <a id="L343"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L344"></a>exprs[i] = list.At(i).(ast.Expr)
    <a id="L345"></a>}

    <a id="L347"></a>return exprs;
<a id="L348"></a>}


<a id="L351"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L352"></a><span class="comment">// Types</span>

<a id="L354"></a>func (p *parser) parseType() ast.Expr {
    <a id="L355"></a>if p.trace {
        <a id="L356"></a>defer un(trace(p, &#34;Type&#34;))
    <a id="L357"></a>}

    <a id="L359"></a>typ := p.tryType();

    <a id="L361"></a>if typ == nil {
        <a id="L362"></a>p.errorExpected(p.pos, &#34;type&#34;);
        <a id="L363"></a>p.next(); <span class="comment">// make progress</span>
        <a id="L364"></a>return &amp;ast.BadExpr{p.pos};
    <a id="L365"></a>}

    <a id="L367"></a>return typ;
<a id="L368"></a>}


<a id="L371"></a>func (p *parser) parseQualifiedIdent() ast.Expr {
    <a id="L372"></a>if p.trace {
        <a id="L373"></a>defer un(trace(p, &#34;QualifiedIdent&#34;))
    <a id="L374"></a>}

    <a id="L376"></a>var x ast.Expr = p.parseIdent();
    <a id="L377"></a>if p.tok == token.PERIOD {
        <a id="L378"></a><span class="comment">// first identifier is a package identifier</span>
        <a id="L379"></a>p.next();
        <a id="L380"></a>sel := p.parseIdent();
        <a id="L381"></a>x = &amp;ast.SelectorExpr{x, sel};
    <a id="L382"></a>}
    <a id="L383"></a>return x;
<a id="L384"></a>}


<a id="L387"></a>func (p *parser) parseTypeName() ast.Expr {
    <a id="L388"></a>if p.trace {
        <a id="L389"></a>defer un(trace(p, &#34;TypeName&#34;))
    <a id="L390"></a>}

    <a id="L392"></a>return p.parseQualifiedIdent();
<a id="L393"></a>}


<a id="L396"></a>func (p *parser) parseArrayType(ellipsisOk bool) ast.Expr {
    <a id="L397"></a>if p.trace {
        <a id="L398"></a>defer un(trace(p, &#34;ArrayType&#34;))
    <a id="L399"></a>}

    <a id="L401"></a>lbrack := p.expect(token.LBRACK);
    <a id="L402"></a>var len ast.Expr;
    <a id="L403"></a>if ellipsisOk &amp;&amp; p.tok == token.ELLIPSIS {
        <a id="L404"></a>len = &amp;ast.Ellipsis{p.pos};
        <a id="L405"></a>p.next();
    <a id="L406"></a>} else if p.tok != token.RBRACK {
        <a id="L407"></a>len = p.parseExpr()
    <a id="L408"></a>}
    <a id="L409"></a>p.expect(token.RBRACK);
    <a id="L410"></a>elt := p.parseType();

    <a id="L412"></a>return &amp;ast.ArrayType{lbrack, len, elt};
<a id="L413"></a>}


<a id="L416"></a>func (p *parser) makeIdentList(list *vector.Vector) []*ast.Ident {
    <a id="L417"></a>idents := make([]*ast.Ident, list.Len());
    <a id="L418"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L419"></a>ident, isIdent := list.At(i).(*ast.Ident);
        <a id="L420"></a>if !isIdent {
            <a id="L421"></a>pos := list.At(i).(ast.Expr).Pos();
            <a id="L422"></a>p.errorExpected(pos, &#34;identifier&#34;);
            <a id="L423"></a>idents[i] = &amp;ast.Ident{pos, &#34;&#34;};
        <a id="L424"></a>}
        <a id="L425"></a>idents[i] = ident;
    <a id="L426"></a>}
    <a id="L427"></a>return idents;
<a id="L428"></a>}


<a id="L431"></a>func (p *parser) parseFieldDecl() *ast.Field {
    <a id="L432"></a>if p.trace {
        <a id="L433"></a>defer un(trace(p, &#34;FieldDecl&#34;))
    <a id="L434"></a>}

    <a id="L436"></a>doc := p.leadComment;

    <a id="L438"></a><span class="comment">// a list of identifiers looks like a list of type names</span>
    <a id="L439"></a>list := vector.New(0);
    <a id="L440"></a>for {
        <a id="L441"></a><span class="comment">// TODO(gri): do not allow ()&#39;s here</span>
        <a id="L442"></a>list.Push(p.parseType());
        <a id="L443"></a>if p.tok == token.COMMA {
            <a id="L444"></a>p.next()
        <a id="L445"></a>} else {
            <a id="L446"></a>break
        <a id="L447"></a>}
    <a id="L448"></a>}

    <a id="L450"></a><span class="comment">// if we had a list of identifiers, it must be followed by a type</span>
    <a id="L451"></a>typ := p.tryType();

    <a id="L453"></a><span class="comment">// optional tag</span>
    <a id="L454"></a>var tag []*ast.BasicLit;
    <a id="L455"></a>if p.tok == token.STRING {
        <a id="L456"></a>tag = p.parseStringList(nil)
    <a id="L457"></a>}

    <a id="L459"></a><span class="comment">// analyze case</span>
    <a id="L460"></a>var idents []*ast.Ident;
    <a id="L461"></a>if typ != nil {
        <a id="L462"></a><span class="comment">// IdentifierList Type</span>
        <a id="L463"></a>idents = p.makeIdentList(list)
    <a id="L464"></a>} else {
        <a id="L465"></a><span class="comment">// Type (anonymous field)</span>
        <a id="L466"></a>if list.Len() == 1 {
            <a id="L467"></a><span class="comment">// TODO(gri): check that this looks like a type</span>
            <a id="L468"></a>typ = list.At(0).(ast.Expr)
        <a id="L469"></a>} else {
            <a id="L470"></a>p.errorExpected(p.pos, &#34;anonymous field&#34;);
            <a id="L471"></a>typ = &amp;ast.BadExpr{p.pos};
        <a id="L472"></a>}
    <a id="L473"></a>}

    <a id="L475"></a>return &amp;ast.Field{doc, idents, typ, tag, nil};
<a id="L476"></a>}


<a id="L479"></a>func (p *parser) parseStructType() *ast.StructType {
    <a id="L480"></a>if p.trace {
        <a id="L481"></a>defer un(trace(p, &#34;StructType&#34;))
    <a id="L482"></a>}

    <a id="L484"></a>pos := p.expect(token.STRUCT);
    <a id="L485"></a>lbrace := p.expect(token.LBRACE);
    <a id="L486"></a>list := vector.New(0);
    <a id="L487"></a>for p.tok == token.IDENT || p.tok == token.MUL {
        <a id="L488"></a>f := p.parseFieldDecl();
        <a id="L489"></a>if p.tok != token.RBRACE {
            <a id="L490"></a>p.expect(token.SEMICOLON)
        <a id="L491"></a>}
        <a id="L492"></a>f.Comment = p.lineComment;
        <a id="L493"></a>list.Push(f);
    <a id="L494"></a>}
    <a id="L495"></a>rbrace := p.expect(token.RBRACE);
    <a id="L496"></a>p.optSemi = true;

    <a id="L498"></a><span class="comment">// convert vector</span>
    <a id="L499"></a>fields := make([]*ast.Field, list.Len());
    <a id="L500"></a>for i := list.Len() - 1; i &gt;= 0; i-- {
        <a id="L501"></a>fields[i] = list.At(i).(*ast.Field)
    <a id="L502"></a>}

    <a id="L504"></a>return &amp;ast.StructType{pos, lbrace, fields, rbrace, false};
<a id="L505"></a>}


<a id="L508"></a>func (p *parser) parsePointerType() *ast.StarExpr {
    <a id="L509"></a>if p.trace {
        <a id="L510"></a>defer un(trace(p, &#34;PointerType&#34;))
    <a id="L511"></a>}

    <a id="L513"></a>star := p.expect(token.MUL);
    <a id="L514"></a>base := p.parseType();

    <a id="L516"></a>return &amp;ast.StarExpr{star, base};
<a id="L517"></a>}


<a id="L520"></a>func (p *parser) tryParameterType(ellipsisOk bool) ast.Expr {
    <a id="L521"></a>if ellipsisOk &amp;&amp; p.tok == token.ELLIPSIS {
        <a id="L522"></a>pos := p.pos;
        <a id="L523"></a>p.next();
        <a id="L524"></a>if p.tok != token.RPAREN {
            <a id="L525"></a><span class="comment">// &#34;...&#34; always must be at the very end of a parameter list</span>
            <a id="L526"></a>p.Error(pos, &#34;expected type, found &#39;...&#39;&#34;)
        <a id="L527"></a>}
        <a id="L528"></a>return &amp;ast.Ellipsis{pos};
    <a id="L529"></a>}
    <a id="L530"></a>return p.tryType();
<a id="L531"></a>}


<a id="L534"></a>func (p *parser) parseParameterType(ellipsisOk bool) ast.Expr {
    <a id="L535"></a>typ := p.tryParameterType(ellipsisOk);
    <a id="L536"></a>if typ == nil {
        <a id="L537"></a>p.errorExpected(p.pos, &#34;type&#34;);
        <a id="L538"></a>p.next(); <span class="comment">// make progress</span>
        <a id="L539"></a>typ = &amp;ast.BadExpr{p.pos};
    <a id="L540"></a>}
    <a id="L541"></a>return typ;
<a id="L542"></a>}


<a id="L545"></a>func (p *parser) parseParameterDecl(ellipsisOk bool) (*vector.Vector, ast.Expr) {
    <a id="L546"></a>if p.trace {
        <a id="L547"></a>defer un(trace(p, &#34;ParameterDecl&#34;))
    <a id="L548"></a>}

    <a id="L550"></a><span class="comment">// a list of identifiers looks like a list of type names</span>
    <a id="L551"></a>list := vector.New(0);
    <a id="L552"></a>for {
        <a id="L553"></a><span class="comment">// TODO(gri): do not allow ()&#39;s here</span>
        <a id="L554"></a>list.Push(p.parseParameterType(ellipsisOk));
        <a id="L555"></a>if p.tok == token.COMMA {
            <a id="L556"></a>p.next()
        <a id="L557"></a>} else {
            <a id="L558"></a>break
        <a id="L559"></a>}
    <a id="L560"></a>}

    <a id="L562"></a><span class="comment">// if we had a list of identifiers, it must be followed by a type</span>
    <a id="L563"></a>typ := p.tryParameterType(ellipsisOk);

    <a id="L565"></a>return list, typ;
<a id="L566"></a>}


<a id="L569"></a>func (p *parser) parseParameterList(ellipsisOk bool) []*ast.Field {
    <a id="L570"></a>if p.trace {
        <a id="L571"></a>defer un(trace(p, &#34;ParameterList&#34;))
    <a id="L572"></a>}

    <a id="L574"></a>list, typ := p.parseParameterDecl(ellipsisOk);
    <a id="L575"></a>if typ != nil {
        <a id="L576"></a><span class="comment">// IdentifierList Type</span>
        <a id="L577"></a>idents := p.makeIdentList(list);
        <a id="L578"></a>list.Init(0);
        <a id="L579"></a>list.Push(&amp;ast.Field{nil, idents, typ, nil, nil});

        <a id="L581"></a>for p.tok == token.COMMA {
            <a id="L582"></a>p.next();
            <a id="L583"></a>idents := p.parseIdentList();
            <a id="L584"></a>typ := p.parseParameterType(ellipsisOk);
            <a id="L585"></a>list.Push(&amp;ast.Field{nil, idents, typ, nil, nil});
        <a id="L586"></a>}

    <a id="L588"></a>} else {
        <a id="L589"></a><span class="comment">// Type { &#34;,&#34; Type } (anonymous parameters)</span>
        <a id="L590"></a><span class="comment">// convert list of types into list of *Param</span>
        <a id="L591"></a>for i := 0; i &lt; list.Len(); i++ {
            <a id="L592"></a>list.Set(i, &amp;ast.Field{Type: list.At(i).(ast.Expr)})
        <a id="L593"></a>}
    <a id="L594"></a>}

    <a id="L596"></a><span class="comment">// convert list</span>
    <a id="L597"></a>params := make([]*ast.Field, list.Len());
    <a id="L598"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L599"></a>params[i] = list.At(i).(*ast.Field)
    <a id="L600"></a>}

    <a id="L602"></a>return params;
<a id="L603"></a>}


<a id="L606"></a>func (p *parser) parseParameters(ellipsisOk bool) []*ast.Field {
    <a id="L607"></a>if p.trace {
        <a id="L608"></a>defer un(trace(p, &#34;Parameters&#34;))
    <a id="L609"></a>}

    <a id="L611"></a>var params []*ast.Field;
    <a id="L612"></a>p.expect(token.LPAREN);
    <a id="L613"></a>if p.tok != token.RPAREN {
        <a id="L614"></a>params = p.parseParameterList(ellipsisOk)
    <a id="L615"></a>}
    <a id="L616"></a>p.expect(token.RPAREN);

    <a id="L618"></a>return params;
<a id="L619"></a>}


<a id="L622"></a>func (p *parser) parseResult() []*ast.Field {
    <a id="L623"></a>if p.trace {
        <a id="L624"></a>defer un(trace(p, &#34;Result&#34;))
    <a id="L625"></a>}

    <a id="L627"></a>var results []*ast.Field;
    <a id="L628"></a>if p.tok == token.LPAREN {
        <a id="L629"></a>results = p.parseParameters(false)
    <a id="L630"></a>} else if p.tok != token.FUNC {
        <a id="L631"></a>typ := p.tryType();
        <a id="L632"></a>if typ != nil {
            <a id="L633"></a>results = make([]*ast.Field, 1);
            <a id="L634"></a>results[0] = &amp;ast.Field{Type: typ};
        <a id="L635"></a>}
    <a id="L636"></a>}

    <a id="L638"></a>return results;
<a id="L639"></a>}


<a id="L642"></a>func (p *parser) parseSignature() (params []*ast.Field, results []*ast.Field) {
    <a id="L643"></a>if p.trace {
        <a id="L644"></a>defer un(trace(p, &#34;Signature&#34;))
    <a id="L645"></a>}

    <a id="L647"></a>params = p.parseParameters(true);
    <a id="L648"></a>results = p.parseResult();

    <a id="L650"></a>return;
<a id="L651"></a>}


<a id="L654"></a>func (p *parser) parseFuncType() *ast.FuncType {
    <a id="L655"></a>if p.trace {
        <a id="L656"></a>defer un(trace(p, &#34;FuncType&#34;))
    <a id="L657"></a>}

    <a id="L659"></a>pos := p.expect(token.FUNC);
    <a id="L660"></a>params, results := p.parseSignature();

    <a id="L662"></a>return &amp;ast.FuncType{pos, params, results};
<a id="L663"></a>}


<a id="L666"></a>func (p *parser) parseMethodSpec() *ast.Field {
    <a id="L667"></a>if p.trace {
        <a id="L668"></a>defer un(trace(p, &#34;MethodSpec&#34;))
    <a id="L669"></a>}

    <a id="L671"></a>doc := p.leadComment;
    <a id="L672"></a>var idents []*ast.Ident;
    <a id="L673"></a>var typ ast.Expr;
    <a id="L674"></a>x := p.parseQualifiedIdent();
    <a id="L675"></a>if ident, isIdent := x.(*ast.Ident); isIdent &amp;&amp; p.tok == token.LPAREN {
        <a id="L676"></a><span class="comment">// method</span>
        <a id="L677"></a>idents = []*ast.Ident{ident};
        <a id="L678"></a>params, results := p.parseSignature();
        <a id="L679"></a>typ = &amp;ast.FuncType{noPos, params, results};
    <a id="L680"></a>} else {
        <a id="L681"></a><span class="comment">// embedded interface</span>
        <a id="L682"></a>typ = x
    <a id="L683"></a>}

    <a id="L685"></a>return &amp;ast.Field{doc, idents, typ, nil, nil};
<a id="L686"></a>}


<a id="L689"></a>func (p *parser) parseInterfaceType() *ast.InterfaceType {
    <a id="L690"></a>if p.trace {
        <a id="L691"></a>defer un(trace(p, &#34;InterfaceType&#34;))
    <a id="L692"></a>}

    <a id="L694"></a>pos := p.expect(token.INTERFACE);
    <a id="L695"></a>lbrace := p.expect(token.LBRACE);
    <a id="L696"></a>list := vector.New(0);
    <a id="L697"></a>for p.tok == token.IDENT {
        <a id="L698"></a>m := p.parseMethodSpec();
        <a id="L699"></a>if p.tok != token.RBRACE {
            <a id="L700"></a>p.expect(token.SEMICOLON)
        <a id="L701"></a>}
        <a id="L702"></a>m.Comment = p.lineComment;
        <a id="L703"></a>list.Push(m);
    <a id="L704"></a>}
    <a id="L705"></a>rbrace := p.expect(token.RBRACE);
    <a id="L706"></a>p.optSemi = true;

    <a id="L708"></a><span class="comment">// convert vector</span>
    <a id="L709"></a>methods := make([]*ast.Field, list.Len());
    <a id="L710"></a>for i := list.Len() - 1; i &gt;= 0; i-- {
        <a id="L711"></a>methods[i] = list.At(i).(*ast.Field)
    <a id="L712"></a>}

    <a id="L714"></a>return &amp;ast.InterfaceType{pos, lbrace, methods, rbrace, false};
<a id="L715"></a>}


<a id="L718"></a>func (p *parser) parseMapType() *ast.MapType {
    <a id="L719"></a>if p.trace {
        <a id="L720"></a>defer un(trace(p, &#34;MapType&#34;))
    <a id="L721"></a>}

    <a id="L723"></a>pos := p.expect(token.MAP);
    <a id="L724"></a>p.expect(token.LBRACK);
    <a id="L725"></a>key := p.parseType();
    <a id="L726"></a>p.expect(token.RBRACK);
    <a id="L727"></a>value := p.parseType();

    <a id="L729"></a>return &amp;ast.MapType{pos, key, value};
<a id="L730"></a>}


<a id="L733"></a>func (p *parser) parseChanType() *ast.ChanType {
    <a id="L734"></a>if p.trace {
        <a id="L735"></a>defer un(trace(p, &#34;ChanType&#34;))
    <a id="L736"></a>}

    <a id="L738"></a>pos := p.pos;
    <a id="L739"></a>dir := ast.SEND | ast.RECV;
    <a id="L740"></a>if p.tok == token.CHAN {
        <a id="L741"></a>p.next();
        <a id="L742"></a>if p.tok == token.ARROW {
            <a id="L743"></a>p.next();
            <a id="L744"></a>dir = ast.SEND;
        <a id="L745"></a>}
    <a id="L746"></a>} else {
        <a id="L747"></a>p.expect(token.ARROW);
        <a id="L748"></a>p.expect(token.CHAN);
        <a id="L749"></a>dir = ast.RECV;
    <a id="L750"></a>}
    <a id="L751"></a>value := p.parseType();

    <a id="L753"></a>return &amp;ast.ChanType{pos, dir, value};
<a id="L754"></a>}


<a id="L757"></a>func (p *parser) tryRawType(ellipsisOk bool) ast.Expr {
    <a id="L758"></a>switch p.tok {
    <a id="L759"></a>case token.IDENT:
        <a id="L760"></a>return p.parseTypeName()
    <a id="L761"></a>case token.LBRACK:
        <a id="L762"></a>return p.parseArrayType(ellipsisOk)
    <a id="L763"></a>case token.STRUCT:
        <a id="L764"></a>return p.parseStructType()
    <a id="L765"></a>case token.MUL:
        <a id="L766"></a>return p.parsePointerType()
    <a id="L767"></a>case token.FUNC:
        <a id="L768"></a>return p.parseFuncType()
    <a id="L769"></a>case token.INTERFACE:
        <a id="L770"></a>return p.parseInterfaceType()
    <a id="L771"></a>case token.MAP:
        <a id="L772"></a>return p.parseMapType()
    <a id="L773"></a>case token.CHAN, token.ARROW:
        <a id="L774"></a>return p.parseChanType()
    <a id="L775"></a>case token.LPAREN:
        <a id="L776"></a>lparen := p.pos;
        <a id="L777"></a>p.next();
        <a id="L778"></a>typ := p.parseType();
        <a id="L779"></a>rparen := p.expect(token.RPAREN);
        <a id="L780"></a>return &amp;ast.ParenExpr{lparen, typ, rparen};
    <a id="L781"></a>}

    <a id="L783"></a><span class="comment">// no type found</span>
    <a id="L784"></a>return nil;
<a id="L785"></a>}


<a id="L788"></a>func (p *parser) tryType() ast.Expr { return p.tryRawType(false) }


<a id="L791"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L792"></a><span class="comment">// Blocks</span>

<a id="L794"></a>func makeStmtList(list *vector.Vector) []ast.Stmt {
    <a id="L795"></a>stats := make([]ast.Stmt, list.Len());
    <a id="L796"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L797"></a>stats[i] = list.At(i).(ast.Stmt)
    <a id="L798"></a>}
    <a id="L799"></a>return stats;
<a id="L800"></a>}


<a id="L803"></a>func (p *parser) parseStmtList() []ast.Stmt {
    <a id="L804"></a>if p.trace {
        <a id="L805"></a>defer un(trace(p, &#34;StatementList&#34;))
    <a id="L806"></a>}

    <a id="L808"></a>list := vector.New(0);
    <a id="L809"></a>expectSemi := false;
    <a id="L810"></a>for p.tok != token.CASE &amp;&amp; p.tok != token.DEFAULT &amp;&amp; p.tok != token.RBRACE &amp;&amp; p.tok != token.EOF {
        <a id="L811"></a>if expectSemi {
            <a id="L812"></a>p.expect(token.SEMICOLON);
            <a id="L813"></a>expectSemi = false;
        <a id="L814"></a>}
        <a id="L815"></a>list.Push(p.parseStmt());
        <a id="L816"></a>if p.tok == token.SEMICOLON {
            <a id="L817"></a>p.next()
        <a id="L818"></a>} else if p.optSemi {
            <a id="L819"></a>p.optSemi = false <span class="comment">// &#34;consume&#34; optional semicolon</span>
        <a id="L820"></a>} else {
            <a id="L821"></a>expectSemi = true
        <a id="L822"></a>}
    <a id="L823"></a>}

    <a id="L825"></a>return makeStmtList(list);
<a id="L826"></a>}


<a id="L829"></a>func (p *parser) parseBlockStmt(idents []*ast.Ident) *ast.BlockStmt {
    <a id="L830"></a>if p.trace {
        <a id="L831"></a>defer un(trace(p, &#34;BlockStmt&#34;))
    <a id="L832"></a>}

    <a id="L834"></a>defer close(openScope(p));

    <a id="L836"></a>lbrace := p.expect(token.LBRACE);
    <a id="L837"></a>list := p.parseStmtList();
    <a id="L838"></a>rbrace := p.expect(token.RBRACE);
    <a id="L839"></a>p.optSemi = true;

    <a id="L841"></a>return &amp;ast.BlockStmt{lbrace, list, rbrace};
<a id="L842"></a>}


<a id="L845"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L846"></a><span class="comment">// Expressions</span>

<a id="L848"></a>func (p *parser) parseStringList(x *ast.BasicLit) []*ast.BasicLit {
    <a id="L849"></a>if p.trace {
        <a id="L850"></a>defer un(trace(p, &#34;StringList&#34;))
    <a id="L851"></a>}

    <a id="L853"></a>list := vector.New(0);
    <a id="L854"></a>if x != nil {
        <a id="L855"></a>list.Push(x)
    <a id="L856"></a>}

    <a id="L858"></a>for p.tok == token.STRING {
        <a id="L859"></a>list.Push(&amp;ast.BasicLit{p.pos, token.STRING, p.lit});
        <a id="L860"></a>p.next();
    <a id="L861"></a>}

    <a id="L863"></a><span class="comment">// convert list</span>
    <a id="L864"></a>strings := make([]*ast.BasicLit, list.Len());
    <a id="L865"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L866"></a>strings[i] = list.At(i).(*ast.BasicLit)
    <a id="L867"></a>}

    <a id="L869"></a>return strings;
<a id="L870"></a>}


<a id="L873"></a>func (p *parser) parseFuncTypeOrLit() ast.Expr {
    <a id="L874"></a>if p.trace {
        <a id="L875"></a>defer un(trace(p, &#34;FuncTypeOrLit&#34;))
    <a id="L876"></a>}

    <a id="L878"></a>typ := p.parseFuncType();
    <a id="L879"></a>if p.tok != token.LBRACE {
        <a id="L880"></a><span class="comment">// function type only</span>
        <a id="L881"></a>return typ
    <a id="L882"></a>}

    <a id="L884"></a>p.exprLev++;
    <a id="L885"></a>body := p.parseBlockStmt(nil);
    <a id="L886"></a>p.optSemi = false; <span class="comment">// function body requires separating &#34;;&#34;</span>
    <a id="L887"></a>p.exprLev--;

    <a id="L889"></a>return &amp;ast.FuncLit{typ, body};
<a id="L890"></a>}


<a id="L893"></a><span class="comment">// parseOperand may return an expression or a raw type (incl. array</span>
<a id="L894"></a><span class="comment">// types of the form [...]T. Callers must verify the result.</span>
<a id="L895"></a><span class="comment">//</span>
<a id="L896"></a>func (p *parser) parseOperand() ast.Expr {
    <a id="L897"></a>if p.trace {
        <a id="L898"></a>defer un(trace(p, &#34;Operand&#34;))
    <a id="L899"></a>}

    <a id="L901"></a>switch p.tok {
    <a id="L902"></a>case token.IDENT:
        <a id="L903"></a>return p.parseIdent()

    <a id="L905"></a>case token.INT, token.FLOAT, token.CHAR, token.STRING:
        <a id="L906"></a>x := &amp;ast.BasicLit{p.pos, p.tok, p.lit};
        <a id="L907"></a>p.next();
        <a id="L908"></a>if p.tok == token.STRING &amp;&amp; p.tok == token.STRING {
            <a id="L909"></a>return &amp;ast.StringList{p.parseStringList(x)}
        <a id="L910"></a>}
        <a id="L911"></a>return x;

    <a id="L913"></a>case token.LPAREN:
        <a id="L914"></a>lparen := p.pos;
        <a id="L915"></a>p.next();
        <a id="L916"></a>p.exprLev++;
        <a id="L917"></a>x := p.parseExpr();
        <a id="L918"></a>p.exprLev--;
        <a id="L919"></a>rparen := p.expect(token.RPAREN);
        <a id="L920"></a>return &amp;ast.ParenExpr{lparen, x, rparen};

    <a id="L922"></a>case token.FUNC:
        <a id="L923"></a>return p.parseFuncTypeOrLit()

    <a id="L925"></a>default:
        <a id="L926"></a>t := p.tryRawType(true); <span class="comment">// could be type for composite literal or conversion</span>
        <a id="L927"></a>if t != nil {
            <a id="L928"></a>return t
        <a id="L929"></a>}
    <a id="L930"></a>}

    <a id="L932"></a>p.errorExpected(p.pos, &#34;operand&#34;);
    <a id="L933"></a>p.next(); <span class="comment">// make progress</span>
    <a id="L934"></a>return &amp;ast.BadExpr{p.pos};
<a id="L935"></a>}


<a id="L938"></a>func (p *parser) parseSelectorOrTypeAssertion(x ast.Expr) ast.Expr {
    <a id="L939"></a>if p.trace {
        <a id="L940"></a>defer un(trace(p, &#34;SelectorOrTypeAssertion&#34;))
    <a id="L941"></a>}

    <a id="L943"></a>p.expect(token.PERIOD);
    <a id="L944"></a>if p.tok == token.IDENT {
        <a id="L945"></a><span class="comment">// selector</span>
        <a id="L946"></a>sel := p.parseIdent();
        <a id="L947"></a>return &amp;ast.SelectorExpr{x, sel};
    <a id="L948"></a>}

    <a id="L950"></a><span class="comment">// type assertion</span>
    <a id="L951"></a>p.expect(token.LPAREN);
    <a id="L952"></a>var typ ast.Expr;
    <a id="L953"></a>if p.tok == token.TYPE {
        <a id="L954"></a><span class="comment">// type switch: typ == nil</span>
        <a id="L955"></a>p.next()
    <a id="L956"></a>} else {
        <a id="L957"></a>typ = p.parseType()
    <a id="L958"></a>}
    <a id="L959"></a>p.expect(token.RPAREN);

    <a id="L961"></a>return &amp;ast.TypeAssertExpr{x, typ};
<a id="L962"></a>}


<a id="L965"></a>func (p *parser) parseIndex(x ast.Expr) ast.Expr {
    <a id="L966"></a>if p.trace {
        <a id="L967"></a>defer un(trace(p, &#34;Index&#34;))
    <a id="L968"></a>}

    <a id="L970"></a>p.expect(token.LBRACK);
    <a id="L971"></a>p.exprLev++;
    <a id="L972"></a>begin := p.parseExpr();
    <a id="L973"></a>var end ast.Expr;
    <a id="L974"></a>if p.tok == token.COLON {
        <a id="L975"></a>p.next();
        <a id="L976"></a>end = p.parseExpr();
    <a id="L977"></a>}
    <a id="L978"></a>p.exprLev--;
    <a id="L979"></a>p.expect(token.RBRACK);

    <a id="L981"></a>return &amp;ast.IndexExpr{x, begin, end};
<a id="L982"></a>}


<a id="L985"></a>func (p *parser) parseCallOrConversion(fun ast.Expr) *ast.CallExpr {
    <a id="L986"></a>if p.trace {
        <a id="L987"></a>defer un(trace(p, &#34;CallOrConversion&#34;))
    <a id="L988"></a>}

    <a id="L990"></a>lparen := p.expect(token.LPAREN);
    <a id="L991"></a>var args []ast.Expr;
    <a id="L992"></a>if p.tok != token.RPAREN {
        <a id="L993"></a>args = p.parseExprList()
    <a id="L994"></a>}
    <a id="L995"></a>rparen := p.expect(token.RPAREN);

    <a id="L997"></a>return &amp;ast.CallExpr{fun, lparen, args, rparen};
<a id="L998"></a>}


<a id="L1001"></a>func (p *parser) parseElement() ast.Expr {
    <a id="L1002"></a>if p.trace {
        <a id="L1003"></a>defer un(trace(p, &#34;Element&#34;))
    <a id="L1004"></a>}

    <a id="L1006"></a>x := p.parseExpr();
    <a id="L1007"></a>if p.tok == token.COLON {
        <a id="L1008"></a>colon := p.pos;
        <a id="L1009"></a>p.next();
        <a id="L1010"></a>x = &amp;ast.KeyValueExpr{x, colon, p.parseExpr()};
    <a id="L1011"></a>}

    <a id="L1013"></a>return x;
<a id="L1014"></a>}


<a id="L1017"></a>func (p *parser) parseElementList() []ast.Expr {
    <a id="L1018"></a>if p.trace {
        <a id="L1019"></a>defer un(trace(p, &#34;ElementList&#34;))
    <a id="L1020"></a>}

    <a id="L1022"></a>list := vector.New(0);
    <a id="L1023"></a>for p.tok != token.RBRACE &amp;&amp; p.tok != token.EOF {
        <a id="L1024"></a>list.Push(p.parseElement());
        <a id="L1025"></a>if p.tok == token.COMMA {
            <a id="L1026"></a>p.next()
        <a id="L1027"></a>} else {
            <a id="L1028"></a>break
        <a id="L1029"></a>}
    <a id="L1030"></a>}

    <a id="L1032"></a><span class="comment">// convert list</span>
    <a id="L1033"></a>elts := make([]ast.Expr, list.Len());
    <a id="L1034"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L1035"></a>elts[i] = list.At(i).(ast.Expr)
    <a id="L1036"></a>}

    <a id="L1038"></a>return elts;
<a id="L1039"></a>}


<a id="L1042"></a>func (p *parser) parseCompositeLit(typ ast.Expr) ast.Expr {
    <a id="L1043"></a>if p.trace {
        <a id="L1044"></a>defer un(trace(p, &#34;CompositeLit&#34;))
    <a id="L1045"></a>}

    <a id="L1047"></a>lbrace := p.expect(token.LBRACE);
    <a id="L1048"></a>var elts []ast.Expr;
    <a id="L1049"></a>if p.tok != token.RBRACE {
        <a id="L1050"></a>elts = p.parseElementList()
    <a id="L1051"></a>}
    <a id="L1052"></a>rbrace := p.expect(token.RBRACE);
    <a id="L1053"></a>return &amp;ast.CompositeLit{typ, lbrace, elts, rbrace};
<a id="L1054"></a>}


<a id="L1057"></a><span class="comment">// TODO(gri): Consider different approach to checking syntax after parsing:</span>
<a id="L1058"></a><span class="comment">//            Provide a arguments (set of flags) to parsing functions</span>
<a id="L1059"></a><span class="comment">//            restricting what they are supposed to accept depending</span>
<a id="L1060"></a><span class="comment">//            on context.</span>

<a id="L1062"></a><span class="comment">// checkExpr checks that x is an expression (and not a type).</span>
<a id="L1063"></a>func (p *parser) checkExpr(x ast.Expr) ast.Expr {
    <a id="L1064"></a><span class="comment">// TODO(gri): should provide predicate in AST nodes</span>
    <a id="L1065"></a>switch t := x.(type) {
    <a id="L1066"></a>case *ast.BadExpr:
    <a id="L1067"></a>case *ast.Ident:
    <a id="L1068"></a>case *ast.BasicLit:
    <a id="L1069"></a>case *ast.StringList:
    <a id="L1070"></a>case *ast.FuncLit:
    <a id="L1071"></a>case *ast.CompositeLit:
    <a id="L1072"></a>case *ast.ParenExpr:
    <a id="L1073"></a>case *ast.SelectorExpr:
    <a id="L1074"></a>case *ast.IndexExpr:
    <a id="L1075"></a>case *ast.TypeAssertExpr:
        <a id="L1076"></a>if t.Type == nil {
            <a id="L1077"></a><span class="comment">// the form X.(type) is only allowed in type switch expressions</span>
            <a id="L1078"></a>p.errorExpected(x.Pos(), &#34;expression&#34;);
            <a id="L1079"></a>x = &amp;ast.BadExpr{x.Pos()};
        <a id="L1080"></a>}
    <a id="L1081"></a>case *ast.CallExpr:
    <a id="L1082"></a>case *ast.StarExpr:
    <a id="L1083"></a>case *ast.UnaryExpr:
        <a id="L1084"></a>if t.Op == token.RANGE {
            <a id="L1085"></a><span class="comment">// the range operator is only allowed at the top of a for statement</span>
            <a id="L1086"></a>p.errorExpected(x.Pos(), &#34;expression&#34;);
            <a id="L1087"></a>x = &amp;ast.BadExpr{x.Pos()};
        <a id="L1088"></a>}
    <a id="L1089"></a>case *ast.BinaryExpr:
    <a id="L1090"></a>default:
        <a id="L1091"></a><span class="comment">// all other nodes are not proper expressions</span>
        <a id="L1092"></a>p.errorExpected(x.Pos(), &#34;expression&#34;);
        <a id="L1093"></a>x = &amp;ast.BadExpr{x.Pos()};
    <a id="L1094"></a>}
    <a id="L1095"></a>return x;
<a id="L1096"></a>}


<a id="L1099"></a><span class="comment">// isTypeName returns true iff x is type name.</span>
<a id="L1100"></a>func isTypeName(x ast.Expr) bool {
    <a id="L1101"></a><span class="comment">// TODO(gri): should provide predicate in AST nodes</span>
    <a id="L1102"></a>switch t := x.(type) {
    <a id="L1103"></a>case *ast.BadExpr:
    <a id="L1104"></a>case *ast.Ident:
    <a id="L1105"></a>case *ast.ParenExpr:
        <a id="L1106"></a>return isTypeName(t.X) <span class="comment">// TODO(gri): should (TypeName) be illegal?</span>
    <a id="L1107"></a>case *ast.SelectorExpr:
        <a id="L1108"></a>return isTypeName(t.X)
    <a id="L1109"></a>default:
        <a id="L1110"></a>return false <span class="comment">// all other nodes are not type names</span>
    <a id="L1111"></a>}
    <a id="L1112"></a>return true;
<a id="L1113"></a>}


<a id="L1116"></a><span class="comment">// isCompositeLitType returns true iff x is a legal composite literal type.</span>
<a id="L1117"></a>func isCompositeLitType(x ast.Expr) bool {
    <a id="L1118"></a><span class="comment">// TODO(gri): should provide predicate in AST nodes</span>
    <a id="L1119"></a>switch t := x.(type) {
    <a id="L1120"></a>case *ast.BadExpr:
    <a id="L1121"></a>case *ast.Ident:
    <a id="L1122"></a>case *ast.ParenExpr:
        <a id="L1123"></a>return isCompositeLitType(t.X)
    <a id="L1124"></a>case *ast.SelectorExpr:
        <a id="L1125"></a>return isTypeName(t.X)
    <a id="L1126"></a>case *ast.ArrayType:
    <a id="L1127"></a>case *ast.StructType:
    <a id="L1128"></a>case *ast.MapType:
    <a id="L1129"></a>default:
        <a id="L1130"></a>return false <span class="comment">// all other nodes are not legal composite literal types</span>
    <a id="L1131"></a>}
    <a id="L1132"></a>return true;
<a id="L1133"></a>}


<a id="L1136"></a><span class="comment">// checkExprOrType checks that x is an expression or a type</span>
<a id="L1137"></a><span class="comment">// (and not a raw type such as [...]T).</span>
<a id="L1138"></a><span class="comment">//</span>
<a id="L1139"></a>func (p *parser) checkExprOrType(x ast.Expr) ast.Expr {
    <a id="L1140"></a><span class="comment">// TODO(gri): should provide predicate in AST nodes</span>
    <a id="L1141"></a>switch t := x.(type) {
    <a id="L1142"></a>case *ast.UnaryExpr:
        <a id="L1143"></a>if t.Op == token.RANGE {
            <a id="L1144"></a><span class="comment">// the range operator is only allowed at the top of a for statement</span>
            <a id="L1145"></a>p.errorExpected(x.Pos(), &#34;expression&#34;);
            <a id="L1146"></a>x = &amp;ast.BadExpr{x.Pos()};
        <a id="L1147"></a>}
    <a id="L1148"></a>case *ast.ArrayType:
        <a id="L1149"></a>if len, isEllipsis := t.Len.(*ast.Ellipsis); isEllipsis {
            <a id="L1150"></a>p.Error(len.Pos(), &#34;expected array length, found &#39;...&#39;&#34;);
            <a id="L1151"></a>x = &amp;ast.BadExpr{x.Pos()};
        <a id="L1152"></a>}
    <a id="L1153"></a>}

    <a id="L1155"></a><span class="comment">// all other nodes are expressions or types</span>
    <a id="L1156"></a>return x;
<a id="L1157"></a>}


<a id="L1160"></a>func (p *parser) parsePrimaryExpr() ast.Expr {
    <a id="L1161"></a>if p.trace {
        <a id="L1162"></a>defer un(trace(p, &#34;PrimaryExpr&#34;))
    <a id="L1163"></a>}

    <a id="L1165"></a>x := p.parseOperand();
<a id="L1166"></a>L:  for {
        <a id="L1167"></a>switch p.tok {
        <a id="L1168"></a>case token.PERIOD:
            <a id="L1169"></a>x = p.parseSelectorOrTypeAssertion(p.checkExpr(x))
        <a id="L1170"></a>case token.LBRACK:
            <a id="L1171"></a>x = p.parseIndex(p.checkExpr(x))
        <a id="L1172"></a>case token.LPAREN:
            <a id="L1173"></a>x = p.parseCallOrConversion(p.checkExprOrType(x))
        <a id="L1174"></a>case token.LBRACE:
            <a id="L1175"></a>if isCompositeLitType(x) &amp;&amp; (p.exprLev &gt;= 0 || !isTypeName(x)) {
                <a id="L1176"></a>x = p.parseCompositeLit(x)
            <a id="L1177"></a>} else {
                <a id="L1178"></a>break L
            <a id="L1179"></a>}
        <a id="L1180"></a>default:
            <a id="L1181"></a>break L
        <a id="L1182"></a>}
    <a id="L1183"></a>}

    <a id="L1185"></a>return x;
<a id="L1186"></a>}


<a id="L1189"></a>func (p *parser) parseUnaryExpr() ast.Expr {
    <a id="L1190"></a>if p.trace {
        <a id="L1191"></a>defer un(trace(p, &#34;UnaryExpr&#34;))
    <a id="L1192"></a>}

    <a id="L1194"></a>switch p.tok {
    <a id="L1195"></a>case token.ADD, token.SUB, token.NOT, token.XOR, token.ARROW, token.AND, token.RANGE:
        <a id="L1196"></a>pos, op := p.pos, p.tok;
        <a id="L1197"></a>p.next();
        <a id="L1198"></a>x := p.parseUnaryExpr();
        <a id="L1199"></a>return &amp;ast.UnaryExpr{pos, op, p.checkExpr(x)};

    <a id="L1201"></a>case token.MUL:
        <a id="L1202"></a><span class="comment">// unary &#34;*&#34; expression or pointer type</span>
        <a id="L1203"></a>pos := p.pos;
        <a id="L1204"></a>p.next();
        <a id="L1205"></a>x := p.parseUnaryExpr();
        <a id="L1206"></a>return &amp;ast.StarExpr{pos, p.checkExprOrType(x)};
    <a id="L1207"></a>}

    <a id="L1209"></a>return p.parsePrimaryExpr();
<a id="L1210"></a>}


<a id="L1213"></a>func (p *parser) parseBinaryExpr(prec1 int) ast.Expr {
    <a id="L1214"></a>if p.trace {
        <a id="L1215"></a>defer un(trace(p, &#34;BinaryExpr&#34;))
    <a id="L1216"></a>}

    <a id="L1218"></a>x := p.parseUnaryExpr();
    <a id="L1219"></a>for prec := p.tok.Precedence(); prec &gt;= prec1; prec-- {
        <a id="L1220"></a>for p.tok.Precedence() == prec {
            <a id="L1221"></a>pos, op := p.pos, p.tok;
            <a id="L1222"></a>p.next();
            <a id="L1223"></a>y := p.parseBinaryExpr(prec + 1);
            <a id="L1224"></a>x = &amp;ast.BinaryExpr{p.checkExpr(x), pos, op, p.checkExpr(y)};
        <a id="L1225"></a>}
    <a id="L1226"></a>}

    <a id="L1228"></a>return x;
<a id="L1229"></a>}


<a id="L1232"></a><span class="comment">// TODO(gri): parseExpr may return a type or even a raw type ([..]int) -</span>
<a id="L1233"></a><span class="comment">//            should reject when a type/raw type is obviously not allowed</span>
<a id="L1234"></a>func (p *parser) parseExpr() ast.Expr {
    <a id="L1235"></a>if p.trace {
        <a id="L1236"></a>defer un(trace(p, &#34;Expression&#34;))
    <a id="L1237"></a>}

    <a id="L1239"></a>return p.parseBinaryExpr(token.LowestPrec + 1);
<a id="L1240"></a>}


<a id="L1243"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L1244"></a><span class="comment">// Statements</span>


<a id="L1247"></a>func (p *parser) parseSimpleStmt(labelOk bool) ast.Stmt {
    <a id="L1248"></a>if p.trace {
        <a id="L1249"></a>defer un(trace(p, &#34;SimpleStmt&#34;))
    <a id="L1250"></a>}

    <a id="L1252"></a>x := p.parseExprList();

    <a id="L1254"></a>switch p.tok {
    <a id="L1255"></a>case token.COLON:
        <a id="L1256"></a><span class="comment">// labeled statement</span>
        <a id="L1257"></a>p.next();
        <a id="L1258"></a>if labelOk &amp;&amp; len(x) == 1 {
            <a id="L1259"></a>if label, isIdent := x[0].(*ast.Ident); isIdent {
                <a id="L1260"></a>return &amp;ast.LabeledStmt{label, p.parseStmt()}
            <a id="L1261"></a>}
        <a id="L1262"></a>}
        <a id="L1263"></a>p.Error(x[0].Pos(), &#34;illegal label declaration&#34;);
        <a id="L1264"></a>return &amp;ast.BadStmt{x[0].Pos()};

    <a id="L1266"></a>case
        <a id="L1267"></a>token.DEFINE, token.ASSIGN, token.ADD_ASSIGN,
        <a id="L1268"></a>token.SUB_ASSIGN, token.MUL_ASSIGN, token.QUO_ASSIGN,
        <a id="L1269"></a>token.REM_ASSIGN, token.AND_ASSIGN, token.OR_ASSIGN,
        <a id="L1270"></a>token.XOR_ASSIGN, token.SHL_ASSIGN, token.SHR_ASSIGN, token.AND_NOT_ASSIGN:
        <a id="L1271"></a><span class="comment">// assignment statement</span>
        <a id="L1272"></a>pos, tok := p.pos, p.tok;
        <a id="L1273"></a>p.next();
        <a id="L1274"></a>y := p.parseExprList();
        <a id="L1275"></a>if len(x) &gt; 1 &amp;&amp; len(y) &gt; 1 &amp;&amp; len(x) != len(y) {
            <a id="L1276"></a>p.Error(x[0].Pos(), &#34;arity of lhs doesn&#39;t match rhs&#34;)
        <a id="L1277"></a>}
        <a id="L1278"></a>return &amp;ast.AssignStmt{x, pos, tok, y};
    <a id="L1279"></a>}

    <a id="L1281"></a>if len(x) &gt; 1 {
        <a id="L1282"></a>p.Error(x[0].Pos(), &#34;only one expression allowed&#34;)
        <a id="L1283"></a><span class="comment">// continue with first expression</span>
    <a id="L1284"></a>}

    <a id="L1286"></a>if p.tok == token.INC || p.tok == token.DEC {
        <a id="L1287"></a><span class="comment">// increment or decrement</span>
        <a id="L1288"></a>s := &amp;ast.IncDecStmt{x[0], p.tok};
        <a id="L1289"></a>p.next(); <span class="comment">// consume &#34;++&#34; or &#34;--&#34;</span>
        <a id="L1290"></a>return s;
    <a id="L1291"></a>}

    <a id="L1293"></a><span class="comment">// expression</span>
    <a id="L1294"></a>return &amp;ast.ExprStmt{x[0]};
<a id="L1295"></a>}


<a id="L1298"></a>func (p *parser) parseCallExpr() *ast.CallExpr {
    <a id="L1299"></a>x := p.parseExpr();
    <a id="L1300"></a>if call, isCall := x.(*ast.CallExpr); isCall {
        <a id="L1301"></a>return call
    <a id="L1302"></a>}
    <a id="L1303"></a>p.errorExpected(x.Pos(), &#34;function/method call&#34;);
    <a id="L1304"></a>return nil;
<a id="L1305"></a>}


<a id="L1308"></a>func (p *parser) parseGoStmt() ast.Stmt {
    <a id="L1309"></a>if p.trace {
        <a id="L1310"></a>defer un(trace(p, &#34;GoStmt&#34;))
    <a id="L1311"></a>}

    <a id="L1313"></a>pos := p.expect(token.GO);
    <a id="L1314"></a>call := p.parseCallExpr();
    <a id="L1315"></a>if call != nil {
        <a id="L1316"></a>return &amp;ast.GoStmt{pos, call}
    <a id="L1317"></a>}
    <a id="L1318"></a>return &amp;ast.BadStmt{pos};
<a id="L1319"></a>}


<a id="L1322"></a>func (p *parser) parseDeferStmt() ast.Stmt {
    <a id="L1323"></a>if p.trace {
        <a id="L1324"></a>defer un(trace(p, &#34;DeferStmt&#34;))
    <a id="L1325"></a>}

    <a id="L1327"></a>pos := p.expect(token.DEFER);
    <a id="L1328"></a>call := p.parseCallExpr();
    <a id="L1329"></a>if call != nil {
        <a id="L1330"></a>return &amp;ast.DeferStmt{pos, call}
    <a id="L1331"></a>}
    <a id="L1332"></a>return &amp;ast.BadStmt{pos};
<a id="L1333"></a>}


<a id="L1336"></a>func (p *parser) parseReturnStmt() *ast.ReturnStmt {
    <a id="L1337"></a>if p.trace {
        <a id="L1338"></a>defer un(trace(p, &#34;ReturnStmt&#34;))
    <a id="L1339"></a>}

    <a id="L1341"></a>pos := p.pos;
    <a id="L1342"></a>p.expect(token.RETURN);
    <a id="L1343"></a>var x []ast.Expr;
    <a id="L1344"></a>if p.tok != token.SEMICOLON &amp;&amp; p.tok != token.CASE &amp;&amp; p.tok != token.DEFAULT &amp;&amp; p.tok != token.RBRACE {
        <a id="L1345"></a>x = p.parseExprList()
    <a id="L1346"></a>}

    <a id="L1348"></a>return &amp;ast.ReturnStmt{pos, x};
<a id="L1349"></a>}


<a id="L1352"></a>func (p *parser) parseBranchStmt(tok token.Token) *ast.BranchStmt {
    <a id="L1353"></a>if p.trace {
        <a id="L1354"></a>defer un(trace(p, &#34;BranchStmt&#34;))
    <a id="L1355"></a>}

    <a id="L1357"></a>s := &amp;ast.BranchStmt{p.pos, tok, nil};
    <a id="L1358"></a>p.expect(tok);
    <a id="L1359"></a>if tok != token.FALLTHROUGH &amp;&amp; p.tok == token.IDENT {
        <a id="L1360"></a>s.Label = p.parseIdent()
    <a id="L1361"></a>}

    <a id="L1363"></a>return s;
<a id="L1364"></a>}


<a id="L1367"></a>func (p *parser) makeExpr(s ast.Stmt) ast.Expr {
    <a id="L1368"></a>if s == nil {
        <a id="L1369"></a>return nil
    <a id="L1370"></a>}
    <a id="L1371"></a>if es, isExpr := s.(*ast.ExprStmt); isExpr {
        <a id="L1372"></a>return p.checkExpr(es.X)
    <a id="L1373"></a>}
    <a id="L1374"></a>p.Error(s.Pos(), &#34;expected condition, found simple statement&#34;);
    <a id="L1375"></a>return &amp;ast.BadExpr{s.Pos()};
<a id="L1376"></a>}


<a id="L1379"></a>func (p *parser) parseControlClause(isForStmt bool) (s1, s2, s3 ast.Stmt) {
    <a id="L1380"></a>if p.tok != token.LBRACE {
        <a id="L1381"></a>prevLev := p.exprLev;
        <a id="L1382"></a>p.exprLev = -1;

        <a id="L1384"></a>if p.tok != token.SEMICOLON {
            <a id="L1385"></a>s1 = p.parseSimpleStmt(false)
        <a id="L1386"></a>}
        <a id="L1387"></a>if p.tok == token.SEMICOLON {
            <a id="L1388"></a>p.next();
            <a id="L1389"></a>if p.tok != token.LBRACE &amp;&amp; p.tok != token.SEMICOLON {
                <a id="L1390"></a>s2 = p.parseSimpleStmt(false)
            <a id="L1391"></a>}
            <a id="L1392"></a>if isForStmt {
                <a id="L1393"></a><span class="comment">// for statements have a 3rd section</span>
                <a id="L1394"></a>p.expect(token.SEMICOLON);
                <a id="L1395"></a>if p.tok != token.LBRACE {
                    <a id="L1396"></a>s3 = p.parseSimpleStmt(false)
                <a id="L1397"></a>}
            <a id="L1398"></a>}
        <a id="L1399"></a>} else {
            <a id="L1400"></a>s1, s2 = nil, s1
        <a id="L1401"></a>}

        <a id="L1403"></a>p.exprLev = prevLev;
    <a id="L1404"></a>}

    <a id="L1406"></a>return s1, s2, s3;
<a id="L1407"></a>}


<a id="L1410"></a>func (p *parser) parseIfStmt() *ast.IfStmt {
    <a id="L1411"></a>if p.trace {
        <a id="L1412"></a>defer un(trace(p, &#34;IfStmt&#34;))
    <a id="L1413"></a>}

    <a id="L1415"></a><span class="comment">// IfStmt block</span>
    <a id="L1416"></a>defer close(openScope(p));

    <a id="L1418"></a>pos := p.expect(token.IF);
    <a id="L1419"></a>s1, s2, _ := p.parseControlClause(false);
    <a id="L1420"></a>body := p.parseBlockStmt(nil);
    <a id="L1421"></a>var else_ ast.Stmt;
    <a id="L1422"></a>if p.tok == token.ELSE {
        <a id="L1423"></a>p.next();
        <a id="L1424"></a>else_ = p.parseStmt();
    <a id="L1425"></a>}

    <a id="L1427"></a>return &amp;ast.IfStmt{pos, s1, p.makeExpr(s2), body, else_};
<a id="L1428"></a>}


<a id="L1431"></a>func (p *parser) parseCaseClause() *ast.CaseClause {
    <a id="L1432"></a>if p.trace {
        <a id="L1433"></a>defer un(trace(p, &#34;CaseClause&#34;))
    <a id="L1434"></a>}

    <a id="L1436"></a><span class="comment">// CaseClause block</span>
    <a id="L1437"></a>defer close(openScope(p));

    <a id="L1439"></a><span class="comment">// SwitchCase</span>
    <a id="L1440"></a>pos := p.pos;
    <a id="L1441"></a>var x []ast.Expr;
    <a id="L1442"></a>if p.tok == token.CASE {
        <a id="L1443"></a>p.next();
        <a id="L1444"></a>x = p.parseExprList();
    <a id="L1445"></a>} else {
        <a id="L1446"></a>p.expect(token.DEFAULT)
    <a id="L1447"></a>}

    <a id="L1449"></a>colon := p.expect(token.COLON);
    <a id="L1450"></a>body := p.parseStmtList();

    <a id="L1452"></a>return &amp;ast.CaseClause{pos, x, colon, body};
<a id="L1453"></a>}


<a id="L1456"></a>func (p *parser) parseTypeList() []ast.Expr {
    <a id="L1457"></a>if p.trace {
        <a id="L1458"></a>defer un(trace(p, &#34;TypeList&#34;))
    <a id="L1459"></a>}

    <a id="L1461"></a>list := vector.New(0);
    <a id="L1462"></a>list.Push(p.parseType());
    <a id="L1463"></a>for p.tok == token.COMMA {
        <a id="L1464"></a>p.next();
        <a id="L1465"></a>list.Push(p.parseType());
    <a id="L1466"></a>}

    <a id="L1468"></a><span class="comment">// convert list</span>
    <a id="L1469"></a>exprs := make([]ast.Expr, list.Len());
    <a id="L1470"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L1471"></a>exprs[i] = list.At(i).(ast.Expr)
    <a id="L1472"></a>}

    <a id="L1474"></a>return exprs;
<a id="L1475"></a>}


<a id="L1478"></a>func (p *parser) parseTypeCaseClause() *ast.TypeCaseClause {
    <a id="L1479"></a>if p.trace {
        <a id="L1480"></a>defer un(trace(p, &#34;TypeCaseClause&#34;))
    <a id="L1481"></a>}

    <a id="L1483"></a><span class="comment">// TypeCaseClause block</span>
    <a id="L1484"></a>defer close(openScope(p));

    <a id="L1486"></a><span class="comment">// TypeSwitchCase</span>
    <a id="L1487"></a>pos := p.pos;
    <a id="L1488"></a>var types []ast.Expr;
    <a id="L1489"></a>if p.tok == token.CASE {
        <a id="L1490"></a>p.next();
        <a id="L1491"></a>types = p.parseTypeList();
    <a id="L1492"></a>} else {
        <a id="L1493"></a>p.expect(token.DEFAULT)
    <a id="L1494"></a>}

    <a id="L1496"></a>colon := p.expect(token.COLON);
    <a id="L1497"></a>body := p.parseStmtList();

    <a id="L1499"></a>return &amp;ast.TypeCaseClause{pos, types, colon, body};
<a id="L1500"></a>}


<a id="L1503"></a>func isExprSwitch(s ast.Stmt) bool {
    <a id="L1504"></a>if s == nil {
        <a id="L1505"></a>return true
    <a id="L1506"></a>}
    <a id="L1507"></a>if e, ok := s.(*ast.ExprStmt); ok {
        <a id="L1508"></a>if a, ok := e.X.(*ast.TypeAssertExpr); ok {
            <a id="L1509"></a>return a.Type != nil <span class="comment">// regular type assertion</span>
        <a id="L1510"></a>}
        <a id="L1511"></a>return true;
    <a id="L1512"></a>}
    <a id="L1513"></a>return false;
<a id="L1514"></a>}


<a id="L1517"></a>func (p *parser) parseSwitchStmt() ast.Stmt {
    <a id="L1518"></a>if p.trace {
        <a id="L1519"></a>defer un(trace(p, &#34;SwitchStmt&#34;))
    <a id="L1520"></a>}

    <a id="L1522"></a><span class="comment">// SwitchStmt block</span>
    <a id="L1523"></a>defer close(openScope(p));

    <a id="L1525"></a>pos := p.expect(token.SWITCH);
    <a id="L1526"></a>s1, s2, _ := p.parseControlClause(false);

    <a id="L1528"></a>if isExprSwitch(s2) {
        <a id="L1529"></a>lbrace := p.expect(token.LBRACE);
        <a id="L1530"></a>cases := vector.New(0);
        <a id="L1531"></a>for p.tok == token.CASE || p.tok == token.DEFAULT {
            <a id="L1532"></a>cases.Push(p.parseCaseClause())
        <a id="L1533"></a>}
        <a id="L1534"></a>rbrace := p.expect(token.RBRACE);
        <a id="L1535"></a>p.optSemi = true;
        <a id="L1536"></a>body := &amp;ast.BlockStmt{lbrace, makeStmtList(cases), rbrace};
        <a id="L1537"></a>return &amp;ast.SwitchStmt{pos, s1, p.makeExpr(s2), body};
    <a id="L1538"></a>}

    <a id="L1540"></a><span class="comment">// type switch</span>
    <a id="L1541"></a><span class="comment">// TODO(gri): do all the checks!</span>
    <a id="L1542"></a>lbrace := p.expect(token.LBRACE);
    <a id="L1543"></a>cases := vector.New(0);
    <a id="L1544"></a>for p.tok == token.CASE || p.tok == token.DEFAULT {
        <a id="L1545"></a>cases.Push(p.parseTypeCaseClause())
    <a id="L1546"></a>}
    <a id="L1547"></a>rbrace := p.expect(token.RBRACE);
    <a id="L1548"></a>p.optSemi = true;
    <a id="L1549"></a>body := &amp;ast.BlockStmt{lbrace, makeStmtList(cases), rbrace};
    <a id="L1550"></a>return &amp;ast.TypeSwitchStmt{pos, s1, s2, body};
<a id="L1551"></a>}


<a id="L1554"></a>func (p *parser) parseCommClause() *ast.CommClause {
    <a id="L1555"></a>if p.trace {
        <a id="L1556"></a>defer un(trace(p, &#34;CommClause&#34;))
    <a id="L1557"></a>}

    <a id="L1559"></a><span class="comment">// CommClause block</span>
    <a id="L1560"></a>defer close(openScope(p));

    <a id="L1562"></a><span class="comment">// CommCase</span>
    <a id="L1563"></a>pos := p.pos;
    <a id="L1564"></a>var tok token.Token;
    <a id="L1565"></a>var lhs, rhs ast.Expr;
    <a id="L1566"></a>if p.tok == token.CASE {
        <a id="L1567"></a>p.next();
        <a id="L1568"></a>if p.tok == token.ARROW {
            <a id="L1569"></a><span class="comment">// RecvExpr without assignment</span>
            <a id="L1570"></a>rhs = p.parseExpr()
        <a id="L1571"></a>} else {
            <a id="L1572"></a><span class="comment">// SendExpr or RecvExpr</span>
            <a id="L1573"></a>rhs = p.parseExpr();
            <a id="L1574"></a>if p.tok == token.ASSIGN || p.tok == token.DEFINE {
                <a id="L1575"></a><span class="comment">// RecvExpr with assignment</span>
                <a id="L1576"></a>tok = p.tok;
                <a id="L1577"></a>p.next();
                <a id="L1578"></a>lhs = rhs;
                <a id="L1579"></a>if p.tok == token.ARROW {
                    <a id="L1580"></a>rhs = p.parseExpr()
                <a id="L1581"></a>} else {
                    <a id="L1582"></a>p.expect(token.ARROW) <span class="comment">// use expect() error handling</span>
                <a id="L1583"></a>}
            <a id="L1584"></a>}
            <a id="L1585"></a><span class="comment">// else SendExpr</span>
        <a id="L1586"></a>}
    <a id="L1587"></a>} else {
        <a id="L1588"></a>p.expect(token.DEFAULT)
    <a id="L1589"></a>}

    <a id="L1591"></a>colon := p.expect(token.COLON);
    <a id="L1592"></a>body := p.parseStmtList();

    <a id="L1594"></a>return &amp;ast.CommClause{pos, tok, lhs, rhs, colon, body};
<a id="L1595"></a>}


<a id="L1598"></a>func (p *parser) parseSelectStmt() *ast.SelectStmt {
    <a id="L1599"></a>if p.trace {
        <a id="L1600"></a>defer un(trace(p, &#34;SelectStmt&#34;))
    <a id="L1601"></a>}

    <a id="L1603"></a>pos := p.expect(token.SELECT);
    <a id="L1604"></a>lbrace := p.expect(token.LBRACE);
    <a id="L1605"></a>cases := vector.New(0);
    <a id="L1606"></a>for p.tok == token.CASE || p.tok == token.DEFAULT {
        <a id="L1607"></a>cases.Push(p.parseCommClause())
    <a id="L1608"></a>}
    <a id="L1609"></a>rbrace := p.expect(token.RBRACE);
    <a id="L1610"></a>p.optSemi = true;
    <a id="L1611"></a>body := &amp;ast.BlockStmt{lbrace, makeStmtList(cases), rbrace};

    <a id="L1613"></a>return &amp;ast.SelectStmt{pos, body};
<a id="L1614"></a>}


<a id="L1617"></a>func (p *parser) parseForStmt() ast.Stmt {
    <a id="L1618"></a>if p.trace {
        <a id="L1619"></a>defer un(trace(p, &#34;ForStmt&#34;))
    <a id="L1620"></a>}

    <a id="L1622"></a><span class="comment">// ForStmt block</span>
    <a id="L1623"></a>defer close(openScope(p));

    <a id="L1625"></a>pos := p.expect(token.FOR);
    <a id="L1626"></a>s1, s2, s3 := p.parseControlClause(true);
    <a id="L1627"></a>body := p.parseBlockStmt(nil);

    <a id="L1629"></a>if as, isAssign := s2.(*ast.AssignStmt); isAssign {
        <a id="L1630"></a><span class="comment">// possibly a for statement with a range clause; check assignment operator</span>
        <a id="L1631"></a>if as.Tok != token.ASSIGN &amp;&amp; as.Tok != token.DEFINE {
            <a id="L1632"></a>p.errorExpected(as.TokPos, &#34;&#39;=&#39; or &#39;:=&#39;&#34;);
            <a id="L1633"></a>return &amp;ast.BadStmt{pos};
        <a id="L1634"></a>}
        <a id="L1635"></a><span class="comment">// check lhs</span>
        <a id="L1636"></a>var key, value ast.Expr;
        <a id="L1637"></a>switch len(as.Lhs) {
        <a id="L1638"></a>case 2:
            <a id="L1639"></a>value = as.Lhs[1];
            <a id="L1640"></a>fallthrough;
        <a id="L1641"></a>case 1:
            <a id="L1642"></a>key = as.Lhs[0]
        <a id="L1643"></a>default:
            <a id="L1644"></a>p.errorExpected(as.Lhs[0].Pos(), &#34;1 or 2 expressions&#34;);
            <a id="L1645"></a>return &amp;ast.BadStmt{pos};
        <a id="L1646"></a>}
        <a id="L1647"></a><span class="comment">// check rhs</span>
        <a id="L1648"></a>if len(as.Rhs) != 1 {
            <a id="L1649"></a>p.errorExpected(as.Rhs[0].Pos(), &#34;1 expressions&#34;);
            <a id="L1650"></a>return &amp;ast.BadStmt{pos};
        <a id="L1651"></a>}
        <a id="L1652"></a>if rhs, isUnary := as.Rhs[0].(*ast.UnaryExpr); isUnary &amp;&amp; rhs.Op == token.RANGE {
            <a id="L1653"></a><span class="comment">// rhs is range expression; check lhs</span>
            <a id="L1654"></a>return &amp;ast.RangeStmt{pos, key, value, as.TokPos, as.Tok, rhs.X, body}
        <a id="L1655"></a>} else {
            <a id="L1656"></a>p.errorExpected(s2.Pos(), &#34;range clause&#34;);
            <a id="L1657"></a>return &amp;ast.BadStmt{pos};
        <a id="L1658"></a>}
    <a id="L1659"></a>} else {
        <a id="L1660"></a><span class="comment">// regular for statement</span>
        <a id="L1661"></a>return &amp;ast.ForStmt{pos, s1, p.makeExpr(s2), s3, body}
    <a id="L1662"></a>}

    <a id="L1664"></a>panic(); <span class="comment">// unreachable</span>
    <a id="L1665"></a>return nil;
<a id="L1666"></a>}


<a id="L1669"></a>func (p *parser) parseStmt() ast.Stmt {
    <a id="L1670"></a>if p.trace {
        <a id="L1671"></a>defer un(trace(p, &#34;Statement&#34;))
    <a id="L1672"></a>}

    <a id="L1674"></a>switch p.tok {
    <a id="L1675"></a>case token.CONST, token.TYPE, token.VAR:
        <a id="L1676"></a>decl, _ := p.parseDecl(false); <span class="comment">// do not consume trailing semicolon</span>
        <a id="L1677"></a>return &amp;ast.DeclStmt{decl};
    <a id="L1678"></a>case
        <a id="L1679"></a><span class="comment">// tokens that may start a top-level expression</span>
        <a id="L1680"></a>token.IDENT, token.INT, token.FLOAT, token.CHAR, token.STRING, token.FUNC, token.LPAREN, <span class="comment">// operand</span>
        <a id="L1681"></a>token.LBRACK, token.STRUCT, <span class="comment">// composite type</span>
        <a id="L1682"></a>token.MUL, token.AND, token.ARROW, token.ADD, token.SUB, token.XOR: <span class="comment">// unary operators</span>
        <a id="L1683"></a>return p.parseSimpleStmt(true)
    <a id="L1684"></a>case token.GO:
        <a id="L1685"></a>return p.parseGoStmt()
    <a id="L1686"></a>case token.DEFER:
        <a id="L1687"></a>return p.parseDeferStmt()
    <a id="L1688"></a>case token.RETURN:
        <a id="L1689"></a>return p.parseReturnStmt()
    <a id="L1690"></a>case token.BREAK, token.CONTINUE, token.GOTO, token.FALLTHROUGH:
        <a id="L1691"></a>return p.parseBranchStmt(p.tok)
    <a id="L1692"></a>case token.LBRACE:
        <a id="L1693"></a>return p.parseBlockStmt(nil)
    <a id="L1694"></a>case token.IF:
        <a id="L1695"></a>return p.parseIfStmt()
    <a id="L1696"></a>case token.SWITCH:
        <a id="L1697"></a>return p.parseSwitchStmt()
    <a id="L1698"></a>case token.SELECT:
        <a id="L1699"></a>return p.parseSelectStmt()
    <a id="L1700"></a>case token.FOR:
        <a id="L1701"></a>return p.parseForStmt()
    <a id="L1702"></a>case token.SEMICOLON, token.RBRACE:
        <a id="L1703"></a><span class="comment">// don&#39;t consume the &#34;;&#34;, it is the separator following the empty statement</span>
        <a id="L1704"></a>return &amp;ast.EmptyStmt{p.pos}
    <a id="L1705"></a>}

    <a id="L1707"></a><span class="comment">// no statement found</span>
    <a id="L1708"></a>p.errorExpected(p.pos, &#34;statement&#34;);
    <a id="L1709"></a>p.next(); <span class="comment">// make progress</span>
    <a id="L1710"></a>return &amp;ast.BadStmt{p.pos};
<a id="L1711"></a>}


<a id="L1714"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L1715"></a><span class="comment">// Declarations</span>

<a id="L1717"></a>type parseSpecFunction func(p *parser, doc *ast.CommentGroup, getSemi bool) (spec ast.Spec, gotSemi bool)


<a id="L1720"></a><span class="comment">// Consume semicolon if there is one and getSemi is set, and get any line comment.</span>
<a id="L1721"></a><span class="comment">// Return the comment if any and indicate if a semicolon was consumed.</span>
<a id="L1722"></a><span class="comment">//</span>
<a id="L1723"></a>func (p *parser) parseComment(getSemi bool) (comment *ast.CommentGroup, gotSemi bool) {
    <a id="L1724"></a>if getSemi &amp;&amp; p.tok == token.SEMICOLON {
        <a id="L1725"></a>p.next();
        <a id="L1726"></a>gotSemi = true;
    <a id="L1727"></a>}
    <a id="L1728"></a>return p.lineComment, gotSemi;
<a id="L1729"></a>}


<a id="L1732"></a>func parseImportSpec(p *parser, doc *ast.CommentGroup, getSemi bool) (spec ast.Spec, gotSemi bool) {
    <a id="L1733"></a>if p.trace {
        <a id="L1734"></a>defer un(trace(p, &#34;ImportSpec&#34;))
    <a id="L1735"></a>}

    <a id="L1737"></a>var ident *ast.Ident;
    <a id="L1738"></a>if p.tok == token.PERIOD {
        <a id="L1739"></a>ident = &amp;ast.Ident{p.pos, &#34;.&#34;};
        <a id="L1740"></a>p.next();
    <a id="L1741"></a>} else if p.tok == token.IDENT {
        <a id="L1742"></a>ident = p.parseIdent()
    <a id="L1743"></a>}

    <a id="L1745"></a>var path []*ast.BasicLit;
    <a id="L1746"></a>if p.tok == token.STRING {
        <a id="L1747"></a>path = p.parseStringList(nil)
    <a id="L1748"></a>} else {
        <a id="L1749"></a>p.expect(token.STRING) <span class="comment">// use expect() error handling</span>
    <a id="L1750"></a>}

    <a id="L1752"></a>comment, gotSemi := p.parseComment(getSemi);

    <a id="L1754"></a>return &amp;ast.ImportSpec{doc, ident, path, comment}, gotSemi;
<a id="L1755"></a>}


<a id="L1758"></a>func parseConstSpec(p *parser, doc *ast.CommentGroup, getSemi bool) (spec ast.Spec, gotSemi bool) {
    <a id="L1759"></a>if p.trace {
        <a id="L1760"></a>defer un(trace(p, &#34;ConstSpec&#34;))
    <a id="L1761"></a>}

    <a id="L1763"></a>idents := p.parseIdentList();
    <a id="L1764"></a>typ := p.tryType();
    <a id="L1765"></a>var values []ast.Expr;
    <a id="L1766"></a>if typ != nil || p.tok == token.ASSIGN {
        <a id="L1767"></a>p.expect(token.ASSIGN);
        <a id="L1768"></a>values = p.parseExprList();
    <a id="L1769"></a>}
    <a id="L1770"></a>comment, gotSemi := p.parseComment(getSemi);

    <a id="L1772"></a>return &amp;ast.ValueSpec{doc, idents, typ, values, comment}, gotSemi;
<a id="L1773"></a>}


<a id="L1776"></a>func parseTypeSpec(p *parser, doc *ast.CommentGroup, getSemi bool) (spec ast.Spec, gotSemi bool) {
    <a id="L1777"></a>if p.trace {
        <a id="L1778"></a>defer un(trace(p, &#34;TypeSpec&#34;))
    <a id="L1779"></a>}

    <a id="L1781"></a>ident := p.parseIdent();
    <a id="L1782"></a>typ := p.parseType();
    <a id="L1783"></a>comment, gotSemi := p.parseComment(getSemi);

    <a id="L1785"></a>return &amp;ast.TypeSpec{doc, ident, typ, comment}, gotSemi;
<a id="L1786"></a>}


<a id="L1789"></a>func parseVarSpec(p *parser, doc *ast.CommentGroup, getSemi bool) (spec ast.Spec, gotSemi bool) {
    <a id="L1790"></a>if p.trace {
        <a id="L1791"></a>defer un(trace(p, &#34;VarSpec&#34;))
    <a id="L1792"></a>}

    <a id="L1794"></a>idents := p.parseIdentList();
    <a id="L1795"></a>typ := p.tryType();
    <a id="L1796"></a>var values []ast.Expr;
    <a id="L1797"></a>if typ == nil || p.tok == token.ASSIGN {
        <a id="L1798"></a>p.expect(token.ASSIGN);
        <a id="L1799"></a>values = p.parseExprList();
    <a id="L1800"></a>}
    <a id="L1801"></a>comment, gotSemi := p.parseComment(getSemi);

    <a id="L1803"></a>return &amp;ast.ValueSpec{doc, idents, typ, values, comment}, gotSemi;
<a id="L1804"></a>}


<a id="L1807"></a>func (p *parser) parseGenDecl(keyword token.Token, f parseSpecFunction, getSemi bool) (decl *ast.GenDecl, gotSemi bool) {
    <a id="L1808"></a>if p.trace {
        <a id="L1809"></a>defer un(trace(p, keyword.String()+&#34;Decl&#34;))
    <a id="L1810"></a>}

    <a id="L1812"></a>doc := p.leadComment;
    <a id="L1813"></a>pos := p.expect(keyword);
    <a id="L1814"></a>var lparen, rparen token.Position;
    <a id="L1815"></a>list := vector.New(0);
    <a id="L1816"></a>if p.tok == token.LPAREN {
        <a id="L1817"></a>lparen = p.pos;
        <a id="L1818"></a>p.next();
        <a id="L1819"></a>for p.tok != token.RPAREN &amp;&amp; p.tok != token.EOF {
            <a id="L1820"></a>doc := p.leadComment;
            <a id="L1821"></a>spec, semi := f(p, doc, true); <span class="comment">// consume semicolon if any</span>
            <a id="L1822"></a>list.Push(spec);
            <a id="L1823"></a>if !semi {
                <a id="L1824"></a>break
            <a id="L1825"></a>}
        <a id="L1826"></a>}
        <a id="L1827"></a>rparen = p.expect(token.RPAREN);

        <a id="L1829"></a>if getSemi &amp;&amp; p.tok == token.SEMICOLON {
            <a id="L1830"></a>p.next();
            <a id="L1831"></a>gotSemi = true;
        <a id="L1832"></a>} else {
            <a id="L1833"></a>p.optSemi = true
        <a id="L1834"></a>}
    <a id="L1835"></a>} else {
        <a id="L1836"></a>spec, semi := f(p, nil, getSemi);
        <a id="L1837"></a>list.Push(spec);
        <a id="L1838"></a>gotSemi = semi;
    <a id="L1839"></a>}

    <a id="L1841"></a><span class="comment">// convert vector</span>
    <a id="L1842"></a>specs := make([]ast.Spec, list.Len());
    <a id="L1843"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L1844"></a>specs[i] = list.At(i).(ast.Spec)
    <a id="L1845"></a>}

    <a id="L1847"></a>return &amp;ast.GenDecl{doc, pos, keyword, lparen, specs, rparen}, gotSemi;
<a id="L1848"></a>}


<a id="L1851"></a>func (p *parser) parseReceiver() *ast.Field {
    <a id="L1852"></a>if p.trace {
        <a id="L1853"></a>defer un(trace(p, &#34;Receiver&#34;))
    <a id="L1854"></a>}

    <a id="L1856"></a>pos := p.pos;
    <a id="L1857"></a>par := p.parseParameters(false);

    <a id="L1859"></a><span class="comment">// must have exactly one receiver</span>
    <a id="L1860"></a>if len(par) != 1 || len(par) == 1 &amp;&amp; len(par[0].Names) &gt; 1 {
        <a id="L1861"></a>p.errorExpected(pos, &#34;exactly one receiver&#34;);
        <a id="L1862"></a>return &amp;ast.Field{Type: &amp;ast.BadExpr{noPos}};
    <a id="L1863"></a>}

    <a id="L1865"></a>recv := par[0];

    <a id="L1867"></a><span class="comment">// recv type must be TypeName or *TypeName</span>
    <a id="L1868"></a>base := recv.Type;
    <a id="L1869"></a>if ptr, isPtr := base.(*ast.StarExpr); isPtr {
        <a id="L1870"></a>base = ptr.X
    <a id="L1871"></a>}
    <a id="L1872"></a>if !isTypeName(base) {
        <a id="L1873"></a>p.errorExpected(base.Pos(), &#34;type name&#34;)
    <a id="L1874"></a>}

    <a id="L1876"></a>return recv;
<a id="L1877"></a>}


<a id="L1880"></a>func (p *parser) parseFunctionDecl() *ast.FuncDecl {
    <a id="L1881"></a>if p.trace {
        <a id="L1882"></a>defer un(trace(p, &#34;FunctionDecl&#34;))
    <a id="L1883"></a>}

    <a id="L1885"></a>doc := p.leadComment;
    <a id="L1886"></a>pos := p.expect(token.FUNC);

    <a id="L1888"></a>var recv *ast.Field;
    <a id="L1889"></a>if p.tok == token.LPAREN {
        <a id="L1890"></a>recv = p.parseReceiver()
    <a id="L1891"></a>}

    <a id="L1893"></a>ident := p.parseIdent();
    <a id="L1894"></a>params, results := p.parseSignature();

    <a id="L1896"></a>var body *ast.BlockStmt;
    <a id="L1897"></a>if p.tok == token.LBRACE {
        <a id="L1898"></a>body = p.parseBlockStmt(nil)
    <a id="L1899"></a>}

    <a id="L1901"></a>return &amp;ast.FuncDecl{doc, recv, ident, &amp;ast.FuncType{pos, params, results}, body};
<a id="L1902"></a>}


<a id="L1905"></a>func (p *parser) parseDecl(getSemi bool) (decl ast.Decl, gotSemi bool) {
    <a id="L1906"></a>if p.trace {
        <a id="L1907"></a>defer un(trace(p, &#34;Declaration&#34;))
    <a id="L1908"></a>}

    <a id="L1910"></a>var f parseSpecFunction;
    <a id="L1911"></a>switch p.tok {
    <a id="L1912"></a>case token.CONST:
        <a id="L1913"></a>f = parseConstSpec

    <a id="L1915"></a>case token.TYPE:
        <a id="L1916"></a>f = parseTypeSpec

    <a id="L1918"></a>case token.VAR:
        <a id="L1919"></a>f = parseVarSpec

    <a id="L1921"></a>case token.FUNC:
        <a id="L1922"></a>decl = p.parseFunctionDecl();
        <a id="L1923"></a>_, gotSemi := p.parseComment(getSemi);
        <a id="L1924"></a>return decl, gotSemi;

    <a id="L1926"></a>default:
        <a id="L1927"></a>pos := p.pos;
        <a id="L1928"></a>p.errorExpected(pos, &#34;declaration&#34;);
        <a id="L1929"></a>decl = &amp;ast.BadDecl{pos};
        <a id="L1930"></a>gotSemi = getSemi &amp;&amp; p.tok == token.SEMICOLON;
        <a id="L1931"></a>p.next(); <span class="comment">// make progress in any case</span>
        <a id="L1932"></a>return decl, gotSemi;
    <a id="L1933"></a>}

    <a id="L1935"></a>return p.parseGenDecl(p.tok, f, getSemi);
<a id="L1936"></a>}


<a id="L1939"></a>func (p *parser) parseDeclList() []ast.Decl {
    <a id="L1940"></a>if p.trace {
        <a id="L1941"></a>defer un(trace(p, &#34;DeclList&#34;))
    <a id="L1942"></a>}

    <a id="L1944"></a>list := vector.New(0);
    <a id="L1945"></a>for p.tok != token.EOF {
        <a id="L1946"></a>decl, _ := p.parseDecl(true); <span class="comment">// consume optional semicolon</span>
        <a id="L1947"></a>list.Push(decl);
    <a id="L1948"></a>}

    <a id="L1950"></a><span class="comment">// convert vector</span>
    <a id="L1951"></a>decls := make([]ast.Decl, list.Len());
    <a id="L1952"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L1953"></a>decls[i] = list.At(i).(ast.Decl)
    <a id="L1954"></a>}

    <a id="L1956"></a>return decls;
<a id="L1957"></a>}


<a id="L1960"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L1961"></a><span class="comment">// Source files</span>

<a id="L1963"></a>func (p *parser) parseFile() *ast.File {
    <a id="L1964"></a>if p.trace {
        <a id="L1965"></a>defer un(trace(p, &#34;File&#34;))
    <a id="L1966"></a>}

    <a id="L1968"></a><span class="comment">// file block</span>
    <a id="L1969"></a>defer close(openScope(p));

    <a id="L1971"></a><span class="comment">// package clause</span>
    <a id="L1972"></a>doc := p.leadComment;
    <a id="L1973"></a>pos := p.expect(token.PACKAGE);
    <a id="L1974"></a>ident := p.parseIdent();
    <a id="L1975"></a>var decls []ast.Decl;

    <a id="L1977"></a><span class="comment">// Don&#39;t bother parsing the rest if we had errors already.</span>
    <a id="L1978"></a><span class="comment">// Likely not a Go source file at all.</span>

    <a id="L1980"></a>if p.ErrorCount() == 0 &amp;&amp; p.mode&amp;PackageClauseOnly == 0 {
        <a id="L1981"></a><span class="comment">// import decls</span>
        <a id="L1982"></a>list := vector.New(0);
        <a id="L1983"></a>for p.tok == token.IMPORT {
            <a id="L1984"></a>decl, _ := p.parseGenDecl(token.IMPORT, parseImportSpec, true); <span class="comment">// consume optional semicolon</span>
            <a id="L1985"></a>list.Push(decl);
        <a id="L1986"></a>}

        <a id="L1988"></a>if p.mode&amp;ImportsOnly == 0 {
            <a id="L1989"></a><span class="comment">// rest of package body</span>
            <a id="L1990"></a>for p.tok != token.EOF {
                <a id="L1991"></a>decl, _ := p.parseDecl(true); <span class="comment">// consume optional semicolon</span>
                <a id="L1992"></a>list.Push(decl);
            <a id="L1993"></a>}
        <a id="L1994"></a>}

        <a id="L1996"></a><span class="comment">// convert declaration list</span>
        <a id="L1997"></a>decls = make([]ast.Decl, list.Len());
        <a id="L1998"></a>for i := 0; i &lt; list.Len(); i++ {
            <a id="L1999"></a>decls[i] = list.At(i).(ast.Decl)
        <a id="L2000"></a>}
    <a id="L2001"></a>}

    <a id="L2003"></a>return &amp;ast.File{doc, pos, ident, decls, p.comments};
<a id="L2004"></a>}
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
