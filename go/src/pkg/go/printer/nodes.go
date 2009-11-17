<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/printer/nodes.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/go/printer/nodes.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file implements printing of AST nodes; specifically</span>
<a id="L6"></a><span class="comment">// expressions, statements, declarations, and files. It uses</span>
<a id="L7"></a><span class="comment">// the print functionality implemented in printer.go.</span>

<a id="L9"></a>package printer

<a id="L11"></a>import (
    <a id="L12"></a>&#34;bytes&#34;;
    <a id="L13"></a>&#34;go/ast&#34;;
    <a id="L14"></a>&#34;go/token&#34;;
<a id="L15"></a>)


<a id="L18"></a><span class="comment">// Disabled formatting - enable eventually and remove the flag.</span>
<a id="L19"></a>const (
    <a id="L20"></a>compositeLitBlank = false;
    <a id="L21"></a>fewerSemis        = true;
    <a id="L22"></a>stringListMode    = exprListMode(0); <span class="comment">// previously: noIndent</span>
<a id="L23"></a>)


<a id="L26"></a><span class="comment">// Other formatting issues:</span>
<a id="L27"></a><span class="comment">// - replacement of expression spacing algorithm with rsc&#39;s algorithm</span>
<a id="L28"></a><span class="comment">// - better comment formatting for /*-style comments at the end of a line (e.g. a declaration)</span>
<a id="L29"></a><span class="comment">//   when the comment spans multiple lines; if such a comment is just two lines, formatting is</span>
<a id="L30"></a><span class="comment">//   not idempotent</span>
<a id="L31"></a><span class="comment">// - formatting of expression lists; especially for string lists (stringListMode)</span>
<a id="L32"></a><span class="comment">// - blank after { and before } in one-line composite literals probably looks better</span>
<a id="L33"></a><span class="comment">// - should use blank instead of tab to separate one-line function bodies from</span>
<a id="L34"></a><span class="comment">//   the function header unless there is a group of consecutive one-liners</span>


<a id="L37"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L38"></a><span class="comment">// Common AST nodes.</span>

<a id="L40"></a><span class="comment">// Print as many newlines as necessary (but at least min and and at most</span>
<a id="L41"></a><span class="comment">// max newlines) to get to the current line. ws is printed before the first</span>
<a id="L42"></a><span class="comment">// line break. If newSection is set, the first line break is printed as</span>
<a id="L43"></a><span class="comment">// formfeed. Returns true if any line break was printed; returns false otherwise.</span>
<a id="L44"></a><span class="comment">//</span>
<a id="L45"></a><span class="comment">// TODO(gri): Reconsider signature (provide position instead of line)</span>
<a id="L46"></a><span class="comment">//</span>
<a id="L47"></a>func (p *printer) linebreak(line, min, max int, ws whiteSpace, newSection bool) (printedBreak bool) {
    <a id="L48"></a>n := line - p.pos.Line;
    <a id="L49"></a>switch {
    <a id="L50"></a>case n &lt; min:
        <a id="L51"></a>n = min
    <a id="L52"></a>case n &gt; max:
        <a id="L53"></a>n = max
    <a id="L54"></a>}
    <a id="L55"></a>if n &gt; 0 {
        <a id="L56"></a>p.print(ws);
        <a id="L57"></a>if newSection {
            <a id="L58"></a>p.print(formfeed);
            <a id="L59"></a>n--;
            <a id="L60"></a>printedBreak = true;
        <a id="L61"></a>}
    <a id="L62"></a>}
    <a id="L63"></a>for ; n &gt; 0; n-- {
        <a id="L64"></a>p.print(newline);
        <a id="L65"></a>printedBreak = true;
    <a id="L66"></a>}
    <a id="L67"></a>return;
<a id="L68"></a>}


<a id="L71"></a><span class="comment">// TODO(gri): The code for printing lead and line comments</span>
<a id="L72"></a><span class="comment">//            should be eliminated in favor of reusing the</span>
<a id="L73"></a><span class="comment">//            comment intersperse mechanism above somehow.</span>

<a id="L75"></a><span class="comment">// Print a list of individual comments.</span>
<a id="L76"></a>func (p *printer) commentList(list []*ast.Comment) {
    <a id="L77"></a>for i, c := range list {
        <a id="L78"></a>t := c.Text;
        <a id="L79"></a><span class="comment">// TODO(gri): this needs to be styled like normal comments</span>
        <a id="L80"></a>p.print(c.Pos(), t);
        <a id="L81"></a>if t[1] == &#39;/&#39; &amp;&amp; i+1 &lt; len(list) {
            <a id="L82"></a><span class="comment">//-style comment which is not at the end; print a newline</span>
            <a id="L83"></a>p.print(newline)
        <a id="L84"></a>}
    <a id="L85"></a>}
<a id="L86"></a>}


<a id="L89"></a><span class="comment">// Print a lead comment followed by a newline.</span>
<a id="L90"></a>func (p *printer) leadComment(d *ast.CommentGroup) {
    <a id="L91"></a><span class="comment">// Ignore the comment if we have comments interspersed (p.comment != nil).</span>
    <a id="L92"></a>if p.comment == nil &amp;&amp; d != nil {
        <a id="L93"></a>p.commentList(d.List);
        <a id="L94"></a>p.print(newline);
    <a id="L95"></a>}
<a id="L96"></a>}


<a id="L99"></a><span class="comment">// Print a tab followed by a line comment.</span>
<a id="L100"></a><span class="comment">// A newline must be printed afterwards since</span>
<a id="L101"></a><span class="comment">// the comment may be a //-style comment.</span>
<a id="L102"></a>func (p *printer) lineComment(d *ast.CommentGroup) {
    <a id="L103"></a><span class="comment">// Ignore the comment if we have comments interspersed (p.comment != nil).</span>
    <a id="L104"></a>if p.comment == nil &amp;&amp; d != nil {
        <a id="L105"></a>p.print(vtab);
        <a id="L106"></a>p.commentList(d.List);
    <a id="L107"></a>}
<a id="L108"></a>}


<a id="L111"></a><span class="comment">// Sets multiLine to true if the identifier list spans multiple lines.</span>
<a id="L112"></a>func (p *printer) identList(list []*ast.Ident, multiLine *bool) {
    <a id="L113"></a><span class="comment">// convert into an expression list so we can re-use exprList formatting</span>
    <a id="L114"></a>xlist := make([]ast.Expr, len(list));
    <a id="L115"></a>for i, x := range list {
        <a id="L116"></a>xlist[i] = x
    <a id="L117"></a>}
    <a id="L118"></a>p.exprList(noPos, xlist, 1, commaSep, multiLine);
<a id="L119"></a>}


<a id="L122"></a><span class="comment">// Sets multiLine to true if the string list spans multiple lines.</span>
<a id="L123"></a>func (p *printer) stringList(list []*ast.BasicLit, multiLine *bool) {
    <a id="L124"></a><span class="comment">// convert into an expression list so we can re-use exprList formatting</span>
    <a id="L125"></a>xlist := make([]ast.Expr, len(list));
    <a id="L126"></a>for i, x := range list {
        <a id="L127"></a>xlist[i] = x
    <a id="L128"></a>}
    <a id="L129"></a>p.exprList(noPos, xlist, 1, stringListMode, multiLine);
<a id="L130"></a>}


<a id="L133"></a>type exprListMode uint

<a id="L135"></a>const (
    <a id="L136"></a>blankStart exprListMode = 1 &lt;&lt; iota; <span class="comment">// print a blank before a non-empty list</span>
    <a id="L137"></a>blankEnd;               <span class="comment">// print a blank after a non-empty list</span>
    <a id="L138"></a>commaSep;               <span class="comment">// elements are separated by commas</span>
    <a id="L139"></a>commaTerm;              <span class="comment">// elements are terminated by comma</span>
    <a id="L140"></a>noIndent;               <span class="comment">// no extra indentation in multi-line lists</span>
<a id="L141"></a>)


<a id="L144"></a><span class="comment">// Print a list of expressions. If the list spans multiple</span>
<a id="L145"></a><span class="comment">// source lines, the original line breaks are respected between</span>
<a id="L146"></a><span class="comment">// expressions. Sets multiLine to true if the list spans multiple</span>
<a id="L147"></a><span class="comment">// lines.</span>
<a id="L148"></a>func (p *printer) exprList(prev token.Position, list []ast.Expr, depth int, mode exprListMode, multiLine *bool) {
    <a id="L149"></a>if len(list) == 0 {
        <a id="L150"></a>return
    <a id="L151"></a>}

    <a id="L153"></a>if mode&amp;blankStart != 0 {
        <a id="L154"></a>p.print(blank)
    <a id="L155"></a>}

    <a id="L157"></a><span class="comment">// TODO(gri): endLine may be incorrect as it is really the beginning</span>
    <a id="L158"></a><span class="comment">//            of the last list entry. There may be only one, very long</span>
    <a id="L159"></a><span class="comment">//            entry in which case line == endLine.</span>
    <a id="L160"></a>line := list[0].Pos().Line;
    <a id="L161"></a>endLine := list[len(list)-1].Pos().Line;

    <a id="L163"></a>if prev.IsValid() &amp;&amp; prev.Line == line &amp;&amp; line == endLine {
        <a id="L164"></a><span class="comment">// all list entries on a single line</span>
        <a id="L165"></a>for i, x := range list {
            <a id="L166"></a>if i &gt; 0 {
                <a id="L167"></a>if mode&amp;commaSep != 0 {
                    <a id="L168"></a>p.print(token.COMMA)
                <a id="L169"></a>}
                <a id="L170"></a>p.print(blank);
            <a id="L171"></a>}
            <a id="L172"></a>p.expr0(x, depth, multiLine);
        <a id="L173"></a>}
        <a id="L174"></a>if mode&amp;blankEnd != 0 {
            <a id="L175"></a>p.print(blank)
        <a id="L176"></a>}
        <a id="L177"></a>return;
    <a id="L178"></a>}

    <a id="L180"></a><span class="comment">// list entries span multiple lines;</span>
    <a id="L181"></a><span class="comment">// use source code positions to guide line breaks</span>

    <a id="L183"></a><span class="comment">// don&#39;t add extra indentation if noIndent is set;</span>
    <a id="L184"></a><span class="comment">// i.e., pretend that the first line is already indented</span>
    <a id="L185"></a>ws := ignore;
    <a id="L186"></a>if mode&amp;noIndent == 0 {
        <a id="L187"></a>ws = indent
    <a id="L188"></a>}

    <a id="L190"></a>if prev.IsValid() &amp;&amp; prev.Line &lt; line &amp;&amp; p.linebreak(line, 1, 2, ws, true) {
        <a id="L191"></a>ws = ignore;
        <a id="L192"></a>*multiLine = true;
    <a id="L193"></a>}

    <a id="L195"></a>for i, x := range list {
        <a id="L196"></a>prev := line;
        <a id="L197"></a>line = x.Pos().Line;
        <a id="L198"></a>if i &gt; 0 {
            <a id="L199"></a>if mode&amp;commaSep != 0 {
                <a id="L200"></a>p.print(token.COMMA)
            <a id="L201"></a>}
            <a id="L202"></a>if prev &lt; line {
                <a id="L203"></a>if p.linebreak(line, 1, 2, ws, true) {
                    <a id="L204"></a>ws = ignore;
                    <a id="L205"></a>*multiLine = true;
                <a id="L206"></a>}
            <a id="L207"></a>} else {
                <a id="L208"></a>p.print(blank)
            <a id="L209"></a>}
        <a id="L210"></a>}
        <a id="L211"></a>p.expr0(x, depth, multiLine);
    <a id="L212"></a>}

    <a id="L214"></a>if mode&amp;commaTerm != 0 {
        <a id="L215"></a>p.print(token.COMMA);
        <a id="L216"></a>if ws == ignore &amp;&amp; mode&amp;noIndent == 0 {
            <a id="L217"></a><span class="comment">// unindent if we indented</span>
            <a id="L218"></a>p.print(unindent)
        <a id="L219"></a>}
        <a id="L220"></a>p.print(formfeed); <span class="comment">// terminating comma needs a line break to look good</span>
        <a id="L221"></a>return;
    <a id="L222"></a>}

    <a id="L224"></a>if mode&amp;blankEnd != 0 {
        <a id="L225"></a>p.print(blank)
    <a id="L226"></a>}

    <a id="L228"></a>if ws == ignore &amp;&amp; mode&amp;noIndent == 0 {
        <a id="L229"></a><span class="comment">// unindent if we indented</span>
        <a id="L230"></a>p.print(unindent)
    <a id="L231"></a>}
<a id="L232"></a>}


<a id="L235"></a><span class="comment">// Sets multiLine to true if the the parameter list spans multiple lines.</span>
<a id="L236"></a>func (p *printer) parameters(list []*ast.Field, multiLine *bool) {
    <a id="L237"></a>p.print(token.LPAREN);
    <a id="L238"></a>if len(list) &gt; 0 {
        <a id="L239"></a>for i, par := range list {
            <a id="L240"></a>if i &gt; 0 {
                <a id="L241"></a>p.print(token.COMMA, blank)
            <a id="L242"></a>}
            <a id="L243"></a>if len(par.Names) &gt; 0 {
                <a id="L244"></a>p.identList(par.Names, multiLine);
                <a id="L245"></a>p.print(blank);
            <a id="L246"></a>}
            <a id="L247"></a>p.expr(par.Type, multiLine);
        <a id="L248"></a>}
    <a id="L249"></a>}
    <a id="L250"></a>p.print(token.RPAREN);
<a id="L251"></a>}


<a id="L254"></a><span class="comment">// Returns true if a separating semicolon is optional.</span>
<a id="L255"></a><span class="comment">// Sets multiLine to true if the signature spans multiple lines.</span>
<a id="L256"></a>func (p *printer) signature(params, result []*ast.Field, multiLine *bool) (optSemi bool) {
    <a id="L257"></a>p.parameters(params, multiLine);
    <a id="L258"></a>if result != nil {
        <a id="L259"></a>p.print(blank);

        <a id="L261"></a>if len(result) == 1 &amp;&amp; result[0].Names == nil {
            <a id="L262"></a><span class="comment">// single anonymous result; no ()&#39;s unless it&#39;s a function type</span>
            <a id="L263"></a>f := result[0];
            <a id="L264"></a>if _, isFtyp := f.Type.(*ast.FuncType); !isFtyp {
                <a id="L265"></a>optSemi = p.expr(f.Type, multiLine);
                <a id="L266"></a>return;
            <a id="L267"></a>}
        <a id="L268"></a>}

        <a id="L270"></a>p.parameters(result, multiLine);
    <a id="L271"></a>}
    <a id="L272"></a>return;
<a id="L273"></a>}


<a id="L276"></a>func identListSize(list []*ast.Ident, maxSize int) (size int) {
    <a id="L277"></a>for i, x := range list {
        <a id="L278"></a>if i &gt; 0 {
            <a id="L279"></a>size += 2 <span class="comment">// &#34;, &#34;</span>


        <a id="L282"></a>}
        <a id="L283"></a>size += len(x.Value);
        <a id="L284"></a>if size &gt;= maxSize {
            <a id="L285"></a>break
        <a id="L286"></a>}
    <a id="L287"></a>}
    <a id="L288"></a>return;
<a id="L289"></a>}


<a id="L292"></a>func (p *printer) isOneLineFieldList(list []*ast.Field) bool {
    <a id="L293"></a>if len(list) != 1 {
        <a id="L294"></a>return false <span class="comment">// allow only one field</span>
    <a id="L295"></a>}
    <a id="L296"></a>f := list[0];
    <a id="L297"></a>if f.Tag != nil || f.Comment != nil {
        <a id="L298"></a>return false <span class="comment">// don&#39;t allow tags or comments</span>
    <a id="L299"></a>}
    <a id="L300"></a><span class="comment">// only name(s) and type</span>
    <a id="L301"></a>const maxSize = 30; <span class="comment">// adjust as appropriate, this is an approximate value</span>
    <a id="L302"></a>namesSize := identListSize(f.Names, maxSize);
    <a id="L303"></a>if namesSize &gt; 0 {
        <a id="L304"></a>namesSize = 1 <span class="comment">// blank between names and types</span>


    <a id="L307"></a>}
    <a id="L308"></a>typeSize := p.nodeSize(f.Type, maxSize);
    <a id="L309"></a>return namesSize+typeSize &lt;= maxSize;
<a id="L310"></a>}


<a id="L313"></a>func (p *printer) fieldList(lbrace token.Position, list []*ast.Field, rbrace token.Position, isIncomplete bool, ctxt exprContext) {
    <a id="L314"></a>if !isIncomplete &amp;&amp; !p.commentBefore(rbrace) {
        <a id="L315"></a><span class="comment">// possibly a one-line struct/interface</span>
        <a id="L316"></a>if len(list) == 0 {
            <a id="L317"></a><span class="comment">// no blank between keyword and {} in this case</span>
            <a id="L318"></a>p.print(lbrace, token.LBRACE, rbrace, token.RBRACE);
            <a id="L319"></a>return;
        <a id="L320"></a>} else if ctxt&amp;(compositeLit|structType) == compositeLit|structType &amp;&amp;
            <a id="L321"></a>p.isOneLineFieldList(list) { <span class="comment">// for now ignore interfaces</span>
            <a id="L322"></a><span class="comment">// small enough - print on one line</span>
            <a id="L323"></a><span class="comment">// (don&#39;t use identList and ignore source line breaks)</span>
            <a id="L324"></a>p.print(lbrace, token.LBRACE, blank);
            <a id="L325"></a>f := list[0];
            <a id="L326"></a>for i, x := range f.Names {
                <a id="L327"></a>if i &gt; 0 {
                    <a id="L328"></a>p.print(token.COMMA, blank)
                <a id="L329"></a>}
                <a id="L330"></a>p.expr(x, ignoreMultiLine);
            <a id="L331"></a>}
            <a id="L332"></a>if len(f.Names) &gt; 0 {
                <a id="L333"></a>p.print(blank)
            <a id="L334"></a>}
            <a id="L335"></a>p.expr(f.Type, ignoreMultiLine);
            <a id="L336"></a>p.print(blank, rbrace, token.RBRACE);
            <a id="L337"></a>return;
        <a id="L338"></a>}
    <a id="L339"></a>}

    <a id="L341"></a><span class="comment">// at least one entry or incomplete</span>
    <a id="L342"></a>p.print(blank, lbrace, token.LBRACE, indent, formfeed);
    <a id="L343"></a>if ctxt&amp;structType != 0 {

        <a id="L345"></a>sep := vtab;
        <a id="L346"></a>if len(list) == 1 {
            <a id="L347"></a>sep = blank
        <a id="L348"></a>}
        <a id="L349"></a>var ml bool;
        <a id="L350"></a>for i, f := range list {
            <a id="L351"></a>if i &gt; 0 {
                <a id="L352"></a>p.linebreak(f.Pos().Line, 1, 2, ignore, ml)
            <a id="L353"></a>}
            <a id="L354"></a>ml = false;
            <a id="L355"></a>extraTabs := 0;
            <a id="L356"></a>p.leadComment(f.Doc);
            <a id="L357"></a>if len(f.Names) &gt; 0 {
                <a id="L358"></a><span class="comment">// named fields</span>
                <a id="L359"></a>p.identList(f.Names, &amp;ml);
                <a id="L360"></a>p.print(sep);
                <a id="L361"></a>p.expr(f.Type, &amp;ml);
                <a id="L362"></a>extraTabs = 1;
            <a id="L363"></a>} else {
                <a id="L364"></a><span class="comment">// anonymous field</span>
                <a id="L365"></a>p.expr(f.Type, &amp;ml);
                <a id="L366"></a>extraTabs = 2;
            <a id="L367"></a>}
            <a id="L368"></a>if f.Tag != nil {
                <a id="L369"></a>if len(f.Names) &gt; 0 &amp;&amp; sep == vtab {
                    <a id="L370"></a>p.print(sep)
                <a id="L371"></a>}
                <a id="L372"></a>p.print(sep);
                <a id="L373"></a>p.expr(&amp;ast.StringList{f.Tag}, &amp;ml);
                <a id="L374"></a>extraTabs = 0;
            <a id="L375"></a>}
            <a id="L376"></a>p.print(token.SEMICOLON);
            <a id="L377"></a>if f.Comment != nil {
                <a id="L378"></a>for ; extraTabs &gt; 0; extraTabs-- {
                    <a id="L379"></a>p.print(vtab)
                <a id="L380"></a>}
                <a id="L381"></a>p.lineComment(f.Comment);
            <a id="L382"></a>}
        <a id="L383"></a>}
        <a id="L384"></a>if isIncomplete {
            <a id="L385"></a>if len(list) &gt; 0 {
                <a id="L386"></a>p.print(formfeed)
            <a id="L387"></a>}
            <a id="L388"></a><span class="comment">// TODO(gri): this needs to be styled like normal comments</span>
            <a id="L389"></a>p.print(&#34;// contains unexported fields&#34;);
        <a id="L390"></a>}

    <a id="L392"></a>} else { <span class="comment">// interface</span>

        <a id="L394"></a>var ml bool;
        <a id="L395"></a>for i, f := range list {
            <a id="L396"></a>if i &gt; 0 {
                <a id="L397"></a>p.linebreak(f.Pos().Line, 1, 2, ignore, ml)
            <a id="L398"></a>}
            <a id="L399"></a>ml = false;
            <a id="L400"></a>p.leadComment(f.Doc);
            <a id="L401"></a>if ftyp, isFtyp := f.Type.(*ast.FuncType); isFtyp {
                <a id="L402"></a><span class="comment">// method</span>
                <a id="L403"></a>p.expr(f.Names[0], &amp;ml);
                <a id="L404"></a>p.signature(ftyp.Params, ftyp.Results, &amp;ml);
            <a id="L405"></a>} else {
                <a id="L406"></a><span class="comment">// embedded interface</span>
                <a id="L407"></a>p.expr(f.Type, &amp;ml)
            <a id="L408"></a>}
            <a id="L409"></a>p.print(token.SEMICOLON);
            <a id="L410"></a>p.lineComment(f.Comment);
        <a id="L411"></a>}
        <a id="L412"></a>if isIncomplete {
            <a id="L413"></a>if len(list) &gt; 0 {
                <a id="L414"></a>p.print(formfeed)
            <a id="L415"></a>}
            <a id="L416"></a><span class="comment">// TODO(gri): this needs to be styled like normal comments</span>
            <a id="L417"></a>p.print(&#34;// contains unexported methods&#34;);
        <a id="L418"></a>}

    <a id="L420"></a>}
    <a id="L421"></a>p.print(unindent, formfeed, rbrace, token.RBRACE);
<a id="L422"></a>}


<a id="L425"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L426"></a><span class="comment">// Expressions</span>

<a id="L428"></a><span class="comment">// exprContext describes the syntactic environment in which an expression node is printed.</span>
<a id="L429"></a>type exprContext uint

<a id="L431"></a>const (
    <a id="L432"></a>compositeLit = 1 &lt;&lt; iota;
    <a id="L433"></a>structType;
<a id="L434"></a>)


<a id="L437"></a>func walkBinary(e *ast.BinaryExpr) (has5, has6 bool, maxProblem int) {
    <a id="L438"></a>switch e.Op.Precedence() {
    <a id="L439"></a>case 5:
        <a id="L440"></a>has5 = true
    <a id="L441"></a>case 6:
        <a id="L442"></a>has6 = true
    <a id="L443"></a>}

    <a id="L445"></a>switch l := e.X.(type) {
    <a id="L446"></a>case *ast.BinaryExpr:
        <a id="L447"></a>h5, h6, mp := walkBinary(l);
        <a id="L448"></a>has5 = has5 || h5;
        <a id="L449"></a>has6 = has6 || h6;
        <a id="L450"></a>if maxProblem &lt; mp {
            <a id="L451"></a>maxProblem = mp
        <a id="L452"></a>}
    <a id="L453"></a>}

    <a id="L455"></a>switch r := e.Y.(type) {
    <a id="L456"></a>case *ast.BinaryExpr:
        <a id="L457"></a>h5, h6, mp := walkBinary(r);
        <a id="L458"></a>has5 = has5 || h5;
        <a id="L459"></a>has6 = has6 || h6;
        <a id="L460"></a>if maxProblem &lt; mp {
            <a id="L461"></a>maxProblem = mp
        <a id="L462"></a>}

    <a id="L464"></a>case *ast.StarExpr:
        <a id="L465"></a>if e.Op.String() == &#34;/&#34; {
            <a id="L466"></a>maxProblem = 6
        <a id="L467"></a>}

    <a id="L469"></a>case *ast.UnaryExpr:
        <a id="L470"></a>switch e.Op.String() + r.Op.String() {
        <a id="L471"></a>case &#34;/*&#34;:
            <a id="L472"></a>maxProblem = 6
        <a id="L473"></a>case &#34;++&#34;, &#34;--&#34;:
            <a id="L474"></a>if maxProblem &lt; 5 {
                <a id="L475"></a>maxProblem = 5
            <a id="L476"></a>}
        <a id="L477"></a>}
    <a id="L478"></a>}
    <a id="L479"></a>return;
<a id="L480"></a>}


<a id="L483"></a>func cutoff(e *ast.BinaryExpr, depth int) int {
    <a id="L484"></a>has5, has6, maxProblem := walkBinary(e);
    <a id="L485"></a>if maxProblem &gt; 0 {
        <a id="L486"></a>return maxProblem + 1
    <a id="L487"></a>}
    <a id="L488"></a>if has5 &amp;&amp; has6 {
        <a id="L489"></a>if depth == 1 {
            <a id="L490"></a>return 6
        <a id="L491"></a>}
        <a id="L492"></a>return 5;
    <a id="L493"></a>}
    <a id="L494"></a>if depth == 1 {
        <a id="L495"></a>return 7
    <a id="L496"></a>}
    <a id="L497"></a>return 5;
<a id="L498"></a>}


<a id="L501"></a>func diffPrec(expr ast.Expr, prec int) int {
    <a id="L502"></a>x, ok := expr.(*ast.BinaryExpr);
    <a id="L503"></a>if !ok || prec != x.Op.Precedence() {
        <a id="L504"></a>return 1
    <a id="L505"></a>}
    <a id="L506"></a>return 0;
<a id="L507"></a>}


<a id="L510"></a><span class="comment">// Format the binary expression: decide the cutoff and then format.</span>
<a id="L511"></a><span class="comment">// Let&#39;s call depth == 1 Normal mode, and depth &gt; 1 Compact mode.</span>
<a id="L512"></a><span class="comment">// (Algorithm suggestion by Russ Cox.)</span>
<a id="L513"></a><span class="comment">//</span>
<a id="L514"></a><span class="comment">// The precedences are:</span>
<a id="L515"></a><span class="comment">//	6             *  /  %  &lt;&lt;  &gt;&gt;  &amp;  &amp;^</span>
<a id="L516"></a><span class="comment">//	5             +  -  |  ^</span>
<a id="L517"></a><span class="comment">//	4             ==  !=  &lt;  &lt;=  &gt;  &gt;=</span>
<a id="L518"></a><span class="comment">//	3             &lt;-</span>
<a id="L519"></a><span class="comment">//	2             &amp;&amp;</span>
<a id="L520"></a><span class="comment">//	1             ||</span>
<a id="L521"></a><span class="comment">//</span>
<a id="L522"></a><span class="comment">// The only decision is whether there will be spaces around levels 5 and 6.</span>
<a id="L523"></a><span class="comment">// There are never spaces at level 7 (unary), and always spaces at levels 4 and below.</span>
<a id="L524"></a><span class="comment">//</span>
<a id="L525"></a><span class="comment">// To choose the cutoff, look at the whole expression but excluding primary</span>
<a id="L526"></a><span class="comment">// expressions (function calls, parenthesized exprs), and apply these rules:</span>
<a id="L527"></a><span class="comment">//</span>
<a id="L528"></a><span class="comment">//	1) If there is a binary operator with a right side unary operand</span>
<a id="L529"></a><span class="comment">//	   that would clash without a space, the cutoff must be (in order):</span>
<a id="L530"></a><span class="comment">//</span>
<a id="L531"></a><span class="comment">//		&amp;^	7</span>
<a id="L532"></a><span class="comment">//		/*	7</span>
<a id="L533"></a><span class="comment">//		++	6</span>
<a id="L534"></a><span class="comment">//		--	6</span>
<a id="L535"></a><span class="comment">//</span>
<a id="L536"></a><span class="comment">//	2) If there is a mix of level 6 and level 5 operators, then the cutoff</span>
<a id="L537"></a><span class="comment">//	   is 6 (use spaces to distinguish precedence) in Normal mode</span>
<a id="L538"></a><span class="comment">//	   and 5 (never use spaces) in Compact mode.</span>
<a id="L539"></a><span class="comment">//</span>
<a id="L540"></a><span class="comment">//	3) If there are no level 5 operators or no level 6 operators, then the</span>
<a id="L541"></a><span class="comment">//	   cutoff is 7 (always use spaces) in Normal mode</span>
<a id="L542"></a><span class="comment">//	   and 5 (never use spaces) in Compact mode.</span>
<a id="L543"></a><span class="comment">//</span>
<a id="L544"></a><span class="comment">// Sets multiLine to true if the binary expression spans multiple lines.</span>
<a id="L545"></a>func (p *printer) binaryExpr(x *ast.BinaryExpr, prec1, cutoff, depth int, multiLine *bool) {
    <a id="L546"></a>prec := x.Op.Precedence();
    <a id="L547"></a>if prec &lt; prec1 {
        <a id="L548"></a><span class="comment">// parenthesis needed</span>
        <a id="L549"></a><span class="comment">// Note: The parser inserts an ast.ParenExpr node; thus this case</span>
        <a id="L550"></a><span class="comment">//       can only occur if the AST is created in a different way.</span>
        <a id="L551"></a>p.print(token.LPAREN);
        <a id="L552"></a>p.expr0(x, depth-1, multiLine); <span class="comment">// parentheses undo one level of depth</span>
        <a id="L553"></a>p.print(token.RPAREN);
        <a id="L554"></a>return;
    <a id="L555"></a>}

    <a id="L557"></a>printBlank := prec &lt; cutoff;

    <a id="L559"></a>ws := indent;
    <a id="L560"></a>p.expr1(x.X, prec, depth+diffPrec(x.X, prec), 0, multiLine);
    <a id="L561"></a>if printBlank {
        <a id="L562"></a>p.print(blank)
    <a id="L563"></a>}
    <a id="L564"></a>xline := p.pos.Line; <span class="comment">// before the operator (it may be on the next line!)</span>
    <a id="L565"></a>yline := x.Y.Pos().Line;
    <a id="L566"></a>p.print(x.OpPos, x.Op);
    <a id="L567"></a>if xline != yline {
        <a id="L568"></a><span class="comment">//println(x.OpPos.String());</span>
        <a id="L569"></a><span class="comment">// at least one line break, but respect an extra empty line</span>
        <a id="L570"></a><span class="comment">// in the source</span>
        <a id="L571"></a>if p.linebreak(yline, 1, 2, ws, true) {
            <a id="L572"></a>ws = ignore;
            <a id="L573"></a>*multiLine = true;
            <a id="L574"></a>printBlank = false; <span class="comment">// no blank after line break</span>
        <a id="L575"></a>}
    <a id="L576"></a>}
    <a id="L577"></a>if printBlank {
        <a id="L578"></a>p.print(blank)
    <a id="L579"></a>}
    <a id="L580"></a>p.expr1(x.Y, prec, depth+1, 0, multiLine);
    <a id="L581"></a>if ws == ignore {
        <a id="L582"></a>p.print(unindent)
    <a id="L583"></a>}
<a id="L584"></a>}


<a id="L587"></a>func isBinary(expr ast.Expr) bool {
    <a id="L588"></a>_, ok := expr.(*ast.BinaryExpr);
    <a id="L589"></a>return ok;
<a id="L590"></a>}


<a id="L593"></a><span class="comment">// Returns true if a separating semicolon is optional.</span>
<a id="L594"></a><span class="comment">// Sets multiLine to true if the expression spans multiple lines.</span>
<a id="L595"></a>func (p *printer) expr1(expr ast.Expr, prec1, depth int, ctxt exprContext, multiLine *bool) (optSemi bool) {
    <a id="L596"></a>p.print(expr.Pos());

    <a id="L598"></a>switch x := expr.(type) {
    <a id="L599"></a>case *ast.BadExpr:
        <a id="L600"></a>p.print(&#34;BadExpr&#34;)

    <a id="L602"></a>case *ast.Ident:
        <a id="L603"></a>p.print(x)

    <a id="L605"></a>case *ast.BinaryExpr:
        <a id="L606"></a>if depth &lt; 1 {
            <a id="L607"></a>p.internalError(&#34;depth &lt; 1:&#34;, depth);
            <a id="L608"></a>depth = 1;
        <a id="L609"></a>}
        <a id="L610"></a>p.binaryExpr(x, prec1, cutoff(x, depth), depth, multiLine);

    <a id="L612"></a>case *ast.KeyValueExpr:
        <a id="L613"></a>p.expr(x.Key, multiLine);
        <a id="L614"></a>p.print(x.Colon, token.COLON, blank);
        <a id="L615"></a>p.expr(x.Value, multiLine);

    <a id="L617"></a>case *ast.StarExpr:
        <a id="L618"></a>p.print(token.MUL);
        <a id="L619"></a>optSemi = p.expr(x.X, multiLine);

    <a id="L621"></a>case *ast.UnaryExpr:
        <a id="L622"></a>const prec = token.UnaryPrec;
        <a id="L623"></a>if prec &lt; prec1 {
            <a id="L624"></a><span class="comment">// parenthesis needed</span>
            <a id="L625"></a>p.print(token.LPAREN);
            <a id="L626"></a>p.expr(x, multiLine);
            <a id="L627"></a>p.print(token.RPAREN);
        <a id="L628"></a>} else {
            <a id="L629"></a><span class="comment">// no parenthesis needed</span>
            <a id="L630"></a>p.print(x.Op);
            <a id="L631"></a>if x.Op == token.RANGE {
                <a id="L632"></a>p.print(blank)
            <a id="L633"></a>}
            <a id="L634"></a>p.expr1(x.X, prec, depth, 0, multiLine);
        <a id="L635"></a>}

    <a id="L637"></a>case *ast.BasicLit:
        <a id="L638"></a>p.print(x)

    <a id="L640"></a>case *ast.StringList:
        <a id="L641"></a>p.stringList(x.Strings, multiLine)

    <a id="L643"></a>case *ast.FuncLit:
        <a id="L644"></a>p.expr(x.Type, multiLine);
        <a id="L645"></a>p.funcBody(x.Body, distance(x.Type.Pos(), p.pos), true, multiLine);

    <a id="L647"></a>case *ast.ParenExpr:
        <a id="L648"></a>p.print(token.LPAREN);
        <a id="L649"></a>p.expr0(x.X, depth-1, multiLine); <span class="comment">// parentheses undo one level of depth</span>
        <a id="L650"></a>p.print(x.Rparen, token.RPAREN);

    <a id="L652"></a>case *ast.SelectorExpr:
        <a id="L653"></a>p.expr1(x.X, token.HighestPrec, depth, 0, multiLine);
        <a id="L654"></a>p.print(token.PERIOD);
        <a id="L655"></a>p.expr1(x.Sel, token.HighestPrec, depth, 0, multiLine);

    <a id="L657"></a>case *ast.TypeAssertExpr:
        <a id="L658"></a>p.expr1(x.X, token.HighestPrec, depth, 0, multiLine);
        <a id="L659"></a>p.print(token.PERIOD, token.LPAREN);
        <a id="L660"></a>if x.Type != nil {
            <a id="L661"></a>p.expr(x.Type, multiLine)
        <a id="L662"></a>} else {
            <a id="L663"></a>p.print(token.TYPE)
        <a id="L664"></a>}
        <a id="L665"></a>p.print(token.RPAREN);

    <a id="L667"></a>case *ast.IndexExpr:
        <a id="L668"></a>p.expr1(x.X, token.HighestPrec, 1, 0, multiLine);
        <a id="L669"></a>p.print(token.LBRACK);
        <a id="L670"></a>p.expr0(x.Index, depth+1, multiLine);
        <a id="L671"></a>if x.End != nil {
            <a id="L672"></a><span class="comment">// blanks around &#34;:&#34; if either side is a binary expression</span>
            <a id="L673"></a>if depth &lt;= 1 &amp;&amp; (isBinary(x.Index) || isBinary(x.End)) {
                <a id="L674"></a>p.print(blank, token.COLON, blank)
            <a id="L675"></a>} else {
                <a id="L676"></a>p.print(token.COLON)
            <a id="L677"></a>}
            <a id="L678"></a>p.expr0(x.End, depth+1, multiLine);
        <a id="L679"></a>}
        <a id="L680"></a>p.print(token.RBRACK);

    <a id="L682"></a>case *ast.CallExpr:
        <a id="L683"></a>if len(x.Args) &gt; 1 {
            <a id="L684"></a>depth++
        <a id="L685"></a>}
        <a id="L686"></a>p.expr1(x.Fun, token.HighestPrec, depth, 0, multiLine);
        <a id="L687"></a>p.print(x.Lparen, token.LPAREN);
        <a id="L688"></a>p.exprList(x.Lparen, x.Args, depth, commaSep, multiLine);
        <a id="L689"></a>p.print(x.Rparen, token.RPAREN);

    <a id="L691"></a>case *ast.CompositeLit:
        <a id="L692"></a>p.expr1(x.Type, token.HighestPrec, depth, compositeLit, multiLine);
        <a id="L693"></a>mode := commaSep | commaTerm;
        <a id="L694"></a>if compositeLitBlank {
            <a id="L695"></a><span class="comment">// add blank padding around composite literal</span>
            <a id="L696"></a><span class="comment">// contents for a less dense look</span>
            <a id="L697"></a>mode |= blankStart | blankEnd;
            <a id="L698"></a>if x.Lbrace.Line &lt; x.Rbrace.Line {
                <a id="L699"></a><span class="comment">// add a blank before the opening { for multi-line composites</span>
                <a id="L700"></a><span class="comment">// TODO(gri): for now this decision is made by looking at the</span>
                <a id="L701"></a><span class="comment">//            source code - it may not be correct if the source</span>
                <a id="L702"></a><span class="comment">//            code was badly misformatted in the first place</span>
                <a id="L703"></a>p.print(blank)
            <a id="L704"></a>}
        <a id="L705"></a>}
        <a id="L706"></a>p.print(x.Lbrace, token.LBRACE);
        <a id="L707"></a>p.exprList(x.Lbrace, x.Elts, 1, mode, multiLine);
        <a id="L708"></a>p.print(x.Rbrace, token.RBRACE);

    <a id="L710"></a>case *ast.Ellipsis:
        <a id="L711"></a>p.print(token.ELLIPSIS)

    <a id="L713"></a>case *ast.ArrayType:
        <a id="L714"></a>p.print(token.LBRACK);
        <a id="L715"></a>if x.Len != nil {
            <a id="L716"></a>p.expr(x.Len, multiLine)
        <a id="L717"></a>}
        <a id="L718"></a>p.print(token.RBRACK);
        <a id="L719"></a>optSemi = p.expr(x.Elt, multiLine);

    <a id="L721"></a>case *ast.StructType:
        <a id="L722"></a>p.print(token.STRUCT);
        <a id="L723"></a>p.fieldList(x.Lbrace, x.Fields, x.Rbrace, x.Incomplete, ctxt|structType);
        <a id="L724"></a>optSemi = true;

    <a id="L726"></a>case *ast.FuncType:
        <a id="L727"></a>p.print(token.FUNC);
        <a id="L728"></a>optSemi = p.signature(x.Params, x.Results, multiLine);

    <a id="L730"></a>case *ast.InterfaceType:
        <a id="L731"></a>p.print(token.INTERFACE);
        <a id="L732"></a>p.fieldList(x.Lbrace, x.Methods, x.Rbrace, x.Incomplete, ctxt);
        <a id="L733"></a>optSemi = true;

    <a id="L735"></a>case *ast.MapType:
        <a id="L736"></a>p.print(token.MAP, token.LBRACK);
        <a id="L737"></a>p.expr(x.Key, multiLine);
        <a id="L738"></a>p.print(token.RBRACK);
        <a id="L739"></a>optSemi = p.expr(x.Value, multiLine);

    <a id="L741"></a>case *ast.ChanType:
        <a id="L742"></a>switch x.Dir {
        <a id="L743"></a>case ast.SEND | ast.RECV:
            <a id="L744"></a>p.print(token.CHAN)
        <a id="L745"></a>case ast.RECV:
            <a id="L746"></a>p.print(token.ARROW, token.CHAN)
        <a id="L747"></a>case ast.SEND:
            <a id="L748"></a>p.print(token.CHAN, token.ARROW)
        <a id="L749"></a>}
        <a id="L750"></a>p.print(blank);
        <a id="L751"></a>optSemi = p.expr(x.Value, multiLine);

    <a id="L753"></a>default:
        <a id="L754"></a>panic(&#34;unreachable&#34;)
    <a id="L755"></a>}

    <a id="L757"></a>return;
<a id="L758"></a>}


<a id="L761"></a>func (p *printer) expr0(x ast.Expr, depth int, multiLine *bool) (optSemi bool) {
    <a id="L762"></a>return p.expr1(x, token.LowestPrec, depth, 0, multiLine)
<a id="L763"></a>}


<a id="L766"></a><span class="comment">// Returns true if a separating semicolon is optional.</span>
<a id="L767"></a><span class="comment">// Sets multiLine to true if the expression spans multiple lines.</span>
<a id="L768"></a>func (p *printer) expr(x ast.Expr, multiLine *bool) (optSemi bool) {
    <a id="L769"></a>const depth = 1;
    <a id="L770"></a>return p.expr1(x, token.LowestPrec, depth, 0, multiLine);
<a id="L771"></a>}


<a id="L774"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L775"></a><span class="comment">// Statements</span>

<a id="L777"></a>const maxStmtNewlines = 2 <span class="comment">// maximum number of newlines between statements</span>

<a id="L779"></a><span class="comment">// Print the statement list indented, but without a newline after the last statement.</span>
<a id="L780"></a><span class="comment">// Extra line breaks between statements in the source are respected but at most one</span>
<a id="L781"></a><span class="comment">// empty line is printed between statements.</span>
<a id="L782"></a>func (p *printer) stmtList(list []ast.Stmt, _indent int) {
    <a id="L783"></a><span class="comment">// TODO(gri): fix _indent code</span>
    <a id="L784"></a>if _indent &gt; 0 {
        <a id="L785"></a>p.print(indent)
    <a id="L786"></a>}
    <a id="L787"></a>var multiLine bool;
    <a id="L788"></a>for i, s := range list {
        <a id="L789"></a><span class="comment">// _indent == 0 only for lists of switch/select case clauses;</span>
        <a id="L790"></a><span class="comment">// in those cases each clause is a new section</span>
        <a id="L791"></a>p.linebreak(s.Pos().Line, 1, maxStmtNewlines, ignore, i == 0 || _indent == 0 || multiLine);
        <a id="L792"></a>multiLine = false;
        <a id="L793"></a>if !p.stmt(s, &amp;multiLine) &amp;&amp; (!fewerSemis || len(list) &gt; 1) {
            <a id="L794"></a>p.print(token.SEMICOLON)
        <a id="L795"></a>}
    <a id="L796"></a>}
    <a id="L797"></a>if _indent &gt; 0 {
        <a id="L798"></a>p.print(unindent)
    <a id="L799"></a>}
<a id="L800"></a>}


<a id="L803"></a><span class="comment">// block prints an *ast.BlockStmt; it always spans at least two lines.</span>
<a id="L804"></a>func (p *printer) block(s *ast.BlockStmt, indent int) {
    <a id="L805"></a>p.print(s.Pos(), token.LBRACE);
    <a id="L806"></a>p.stmtList(s.List, indent);
    <a id="L807"></a>p.linebreak(s.Rbrace.Line, 1, maxStmtNewlines, ignore, true);
    <a id="L808"></a>p.print(s.Rbrace, token.RBRACE);
<a id="L809"></a>}


<a id="L812"></a><span class="comment">// TODO(gri): Decide if this should be used more broadly. The printing code</span>
<a id="L813"></a><span class="comment">//            knows when to insert parentheses for precedence reasons, but</span>
<a id="L814"></a><span class="comment">//            need to be careful to keep them around type expressions.</span>
<a id="L815"></a>func stripParens(x ast.Expr) ast.Expr {
    <a id="L816"></a>if px, hasParens := x.(*ast.ParenExpr); hasParens {
        <a id="L817"></a>return stripParens(px.X)
    <a id="L818"></a>}
    <a id="L819"></a>return x;
<a id="L820"></a>}


<a id="L823"></a>func (p *printer) controlClause(isForStmt bool, init ast.Stmt, expr ast.Expr, post ast.Stmt) {
    <a id="L824"></a>p.print(blank);
    <a id="L825"></a>needsBlank := false;
    <a id="L826"></a>if init == nil &amp;&amp; post == nil {
        <a id="L827"></a><span class="comment">// no semicolons required</span>
        <a id="L828"></a>if expr != nil {
            <a id="L829"></a>p.expr(stripParens(expr), ignoreMultiLine);
            <a id="L830"></a>needsBlank = true;
        <a id="L831"></a>}
    <a id="L832"></a>} else {
        <a id="L833"></a><span class="comment">// all semicolons required</span>
        <a id="L834"></a><span class="comment">// (they are not separators, print them explicitly)</span>
        <a id="L835"></a>if init != nil {
            <a id="L836"></a>p.stmt(init, ignoreMultiLine)
        <a id="L837"></a>}
        <a id="L838"></a>p.print(token.SEMICOLON, blank);
        <a id="L839"></a>if expr != nil {
            <a id="L840"></a>p.expr(stripParens(expr), ignoreMultiLine);
            <a id="L841"></a>needsBlank = true;
        <a id="L842"></a>}
        <a id="L843"></a>if isForStmt {
            <a id="L844"></a>p.print(token.SEMICOLON, blank);
            <a id="L845"></a>needsBlank = false;
            <a id="L846"></a>if post != nil {
                <a id="L847"></a>p.stmt(post, ignoreMultiLine);
                <a id="L848"></a>needsBlank = true;
            <a id="L849"></a>}
        <a id="L850"></a>}
    <a id="L851"></a>}
    <a id="L852"></a>if needsBlank {
        <a id="L853"></a>p.print(blank)
    <a id="L854"></a>}
<a id="L855"></a>}


<a id="L858"></a><span class="comment">// Returns true if a separating semicolon is optional.</span>
<a id="L859"></a><span class="comment">// Sets multiLine to true if the statements spans multiple lines.</span>
<a id="L860"></a>func (p *printer) stmt(stmt ast.Stmt, multiLine *bool) (optSemi bool) {
    <a id="L861"></a>p.print(stmt.Pos());

    <a id="L863"></a>switch s := stmt.(type) {
    <a id="L864"></a>case *ast.BadStmt:
        <a id="L865"></a>p.print(&#34;BadStmt&#34;)

    <a id="L867"></a>case *ast.DeclStmt:
        <a id="L868"></a>p.decl(s.Decl, inStmtList, multiLine);
        <a id="L869"></a>optSemi = true; <span class="comment">// decl prints terminating semicolon if necessary</span>

    <a id="L871"></a>case *ast.EmptyStmt:
        <a id="L872"></a><span class="comment">// nothing to do</span>

    <a id="L874"></a>case *ast.LabeledStmt:
        <a id="L875"></a><span class="comment">// a &#34;correcting&#34; unindent immediately following a line break</span>
        <a id="L876"></a><span class="comment">// is applied before the line break  if there is no comment</span>
        <a id="L877"></a><span class="comment">// between (see writeWhitespace)</span>
        <a id="L878"></a>p.print(unindent);
        <a id="L879"></a>p.expr(s.Label, multiLine);
        <a id="L880"></a>p.print(token.COLON, vtab, indent);
        <a id="L881"></a>p.linebreak(s.Stmt.Pos().Line, 0, 1, ignore, true);
        <a id="L882"></a>optSemi = p.stmt(s.Stmt, multiLine);

    <a id="L884"></a>case *ast.ExprStmt:
        <a id="L885"></a>const depth = 1;
        <a id="L886"></a>p.expr0(s.X, depth, multiLine);

    <a id="L888"></a>case *ast.IncDecStmt:
        <a id="L889"></a>const depth = 1;
        <a id="L890"></a>p.expr0(s.X, depth+1, multiLine);
        <a id="L891"></a>p.print(s.Tok);

    <a id="L893"></a>case *ast.AssignStmt:
        <a id="L894"></a>var depth = 1;
        <a id="L895"></a>if len(s.Lhs) &gt; 1 &amp;&amp; len(s.Rhs) &gt; 1 {
            <a id="L896"></a>depth++
        <a id="L897"></a>}
        <a id="L898"></a>p.exprList(s.Pos(), s.Lhs, depth, commaSep, multiLine);
        <a id="L899"></a>p.print(blank, s.TokPos, s.Tok);
        <a id="L900"></a>p.exprList(s.TokPos, s.Rhs, depth, blankStart|commaSep, multiLine);

    <a id="L902"></a>case *ast.GoStmt:
        <a id="L903"></a>p.print(token.GO, blank);
        <a id="L904"></a>p.expr(s.Call, multiLine);

    <a id="L906"></a>case *ast.DeferStmt:
        <a id="L907"></a>p.print(token.DEFER, blank);
        <a id="L908"></a>p.expr(s.Call, multiLine);

    <a id="L910"></a>case *ast.ReturnStmt:
        <a id="L911"></a>p.print(token.RETURN);
        <a id="L912"></a>if s.Results != nil {
            <a id="L913"></a>p.exprList(s.Pos(), s.Results, 1, blankStart|commaSep, multiLine)
        <a id="L914"></a>}

    <a id="L916"></a>case *ast.BranchStmt:
        <a id="L917"></a>p.print(s.Tok);
        <a id="L918"></a>if s.Label != nil {
            <a id="L919"></a>p.print(blank);
            <a id="L920"></a>p.expr(s.Label, multiLine);
        <a id="L921"></a>}

    <a id="L923"></a>case *ast.BlockStmt:
        <a id="L924"></a>p.block(s, 1);
        <a id="L925"></a>*multiLine = true;
        <a id="L926"></a>optSemi = true;

    <a id="L928"></a>case *ast.IfStmt:
        <a id="L929"></a>p.print(token.IF);
        <a id="L930"></a>p.controlClause(false, s.Init, s.Cond, nil);
        <a id="L931"></a>p.block(s.Body, 1);
        <a id="L932"></a>*multiLine = true;
        <a id="L933"></a>optSemi = true;
        <a id="L934"></a>if s.Else != nil {
            <a id="L935"></a>p.print(blank, token.ELSE, blank);
            <a id="L936"></a>switch s.Else.(type) {
            <a id="L937"></a>case *ast.BlockStmt, *ast.IfStmt:
                <a id="L938"></a>optSemi = p.stmt(s.Else, ignoreMultiLine)
            <a id="L939"></a>default:
                <a id="L940"></a>p.print(token.LBRACE, indent, formfeed);
                <a id="L941"></a>p.stmt(s.Else, ignoreMultiLine);
                <a id="L942"></a>p.print(unindent, formfeed, token.RBRACE);
            <a id="L943"></a>}
        <a id="L944"></a>}

    <a id="L946"></a>case *ast.CaseClause:
        <a id="L947"></a>if s.Values != nil {
            <a id="L948"></a>p.print(token.CASE);
            <a id="L949"></a>p.exprList(s.Pos(), s.Values, 1, blankStart|commaSep, multiLine);
        <a id="L950"></a>} else {
            <a id="L951"></a>p.print(token.DEFAULT)
        <a id="L952"></a>}
        <a id="L953"></a>p.print(s.Colon, token.COLON);
        <a id="L954"></a>p.stmtList(s.Body, 1);
        <a id="L955"></a>optSemi = true; <span class="comment">// &#34;block&#34; without {}&#39;s</span>

    <a id="L957"></a>case *ast.SwitchStmt:
        <a id="L958"></a>p.print(token.SWITCH);
        <a id="L959"></a>p.controlClause(false, s.Init, s.Tag, nil);
        <a id="L960"></a>p.block(s.Body, 0);
        <a id="L961"></a>*multiLine = true;
        <a id="L962"></a>optSemi = true;

    <a id="L964"></a>case *ast.TypeCaseClause:
        <a id="L965"></a>if s.Types != nil {
            <a id="L966"></a>p.print(token.CASE);
            <a id="L967"></a>p.exprList(s.Pos(), s.Types, 1, blankStart|commaSep, multiLine);
        <a id="L968"></a>} else {
            <a id="L969"></a>p.print(token.DEFAULT)
        <a id="L970"></a>}
        <a id="L971"></a>p.print(s.Colon, token.COLON);
        <a id="L972"></a>p.stmtList(s.Body, 1);
        <a id="L973"></a>optSemi = true; <span class="comment">// &#34;block&#34; without {}&#39;s</span>

    <a id="L975"></a>case *ast.TypeSwitchStmt:
        <a id="L976"></a>p.print(token.SWITCH);
        <a id="L977"></a>if s.Init != nil {
            <a id="L978"></a>p.print(blank);
            <a id="L979"></a>p.stmt(s.Init, ignoreMultiLine);
            <a id="L980"></a>p.print(token.SEMICOLON);
        <a id="L981"></a>}
        <a id="L982"></a>p.print(blank);
        <a id="L983"></a>p.stmt(s.Assign, ignoreMultiLine);
        <a id="L984"></a>p.print(blank);
        <a id="L985"></a>p.block(s.Body, 0);
        <a id="L986"></a>*multiLine = true;
        <a id="L987"></a>optSemi = true;

    <a id="L989"></a>case *ast.CommClause:
        <a id="L990"></a>if s.Rhs != nil {
            <a id="L991"></a>p.print(token.CASE, blank);
            <a id="L992"></a>if s.Lhs != nil {
                <a id="L993"></a>p.expr(s.Lhs, multiLine);
                <a id="L994"></a>p.print(blank, s.Tok, blank);
            <a id="L995"></a>}
            <a id="L996"></a>p.expr(s.Rhs, multiLine);
        <a id="L997"></a>} else {
            <a id="L998"></a>p.print(token.DEFAULT)
        <a id="L999"></a>}
        <a id="L1000"></a>p.print(s.Colon, token.COLON);
        <a id="L1001"></a>p.stmtList(s.Body, 1);
        <a id="L1002"></a>optSemi = true; <span class="comment">// &#34;block&#34; without {}&#39;s</span>

    <a id="L1004"></a>case *ast.SelectStmt:
        <a id="L1005"></a>p.print(token.SELECT, blank);
        <a id="L1006"></a>p.block(s.Body, 0);
        <a id="L1007"></a>*multiLine = true;
        <a id="L1008"></a>optSemi = true;

    <a id="L1010"></a>case *ast.ForStmt:
        <a id="L1011"></a>p.print(token.FOR);
        <a id="L1012"></a>p.controlClause(true, s.Init, s.Cond, s.Post);
        <a id="L1013"></a>p.block(s.Body, 1);
        <a id="L1014"></a>*multiLine = true;
        <a id="L1015"></a>optSemi = true;

    <a id="L1017"></a>case *ast.RangeStmt:
        <a id="L1018"></a>p.print(token.FOR, blank);
        <a id="L1019"></a>p.expr(s.Key, multiLine);
        <a id="L1020"></a>if s.Value != nil {
            <a id="L1021"></a>p.print(token.COMMA, blank);
            <a id="L1022"></a>p.expr(s.Value, multiLine);
        <a id="L1023"></a>}
        <a id="L1024"></a>p.print(blank, s.TokPos, s.Tok, blank, token.RANGE, blank);
        <a id="L1025"></a>p.expr(s.X, multiLine);
        <a id="L1026"></a>p.print(blank);
        <a id="L1027"></a>p.block(s.Body, 1);
        <a id="L1028"></a>*multiLine = true;
        <a id="L1029"></a>optSemi = true;

    <a id="L1031"></a>default:
        <a id="L1032"></a>panic(&#34;unreachable&#34;)
    <a id="L1033"></a>}

    <a id="L1035"></a>return;
<a id="L1036"></a>}


<a id="L1039"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L1040"></a><span class="comment">// Declarations</span>

<a id="L1042"></a>type declContext uint

<a id="L1044"></a>const (
    <a id="L1045"></a>atTop declContext = iota;
    <a id="L1046"></a>inGroup;
    <a id="L1047"></a>inStmtList;
<a id="L1048"></a>)

<a id="L1050"></a><span class="comment">// The parameter n is the number of specs in the group; context specifies</span>
<a id="L1051"></a><span class="comment">// the surroundings of the declaration. Separating semicolons are printed</span>
<a id="L1052"></a><span class="comment">// depending on the context. Sets multiLine to true if the spec spans</span>
<a id="L1053"></a><span class="comment">// multiple lines.</span>
<a id="L1054"></a><span class="comment">//</span>
<a id="L1055"></a>func (p *printer) spec(spec ast.Spec, n int, context declContext, multiLine *bool) {
    <a id="L1056"></a>var (
        <a id="L1057"></a>optSemi   bool;              <span class="comment">// true if a semicolon is optional</span>
        <a id="L1058"></a>comment   *ast.CommentGroup; <span class="comment">// a line comment, if any</span>
        <a id="L1059"></a>extraTabs int;               <span class="comment">// number of extra tabs before comment, if any</span>
    <a id="L1060"></a>)

    <a id="L1062"></a>switch s := spec.(type) {
    <a id="L1063"></a>case *ast.ImportSpec:
        <a id="L1064"></a>p.leadComment(s.Doc);
        <a id="L1065"></a>if s.Name != nil {
            <a id="L1066"></a>p.expr(s.Name, multiLine);
            <a id="L1067"></a>p.print(blank);
        <a id="L1068"></a>}
        <a id="L1069"></a>p.expr(&amp;ast.StringList{s.Path}, multiLine);
        <a id="L1070"></a>comment = s.Comment;

    <a id="L1072"></a>case *ast.ValueSpec:
        <a id="L1073"></a>p.leadComment(s.Doc);
        <a id="L1074"></a>p.identList(s.Names, multiLine); <span class="comment">// always present</span>
        <a id="L1075"></a>if n == 1 {
            <a id="L1076"></a>if s.Type != nil {
                <a id="L1077"></a>p.print(blank);
                <a id="L1078"></a>optSemi = p.expr(s.Type, multiLine);
            <a id="L1079"></a>}
            <a id="L1080"></a>if s.Values != nil {
                <a id="L1081"></a>p.print(blank, token.ASSIGN);
                <a id="L1082"></a>p.exprList(noPos, s.Values, 1, blankStart|commaSep, multiLine);
                <a id="L1083"></a>optSemi = false;
            <a id="L1084"></a>}
        <a id="L1085"></a>} else {
            <a id="L1086"></a>extraTabs = 2;
            <a id="L1087"></a>if s.Type != nil || s.Values != nil {
                <a id="L1088"></a>p.print(vtab)
            <a id="L1089"></a>}
            <a id="L1090"></a>if s.Type != nil {
                <a id="L1091"></a>optSemi = p.expr(s.Type, multiLine);
                <a id="L1092"></a>extraTabs = 1;
            <a id="L1093"></a>}
            <a id="L1094"></a>if s.Values != nil {
                <a id="L1095"></a>p.print(vtab);
                <a id="L1096"></a>p.print(token.ASSIGN);
                <a id="L1097"></a>p.exprList(noPos, s.Values, 1, blankStart|commaSep, multiLine);
                <a id="L1098"></a>optSemi = false;
                <a id="L1099"></a>extraTabs = 0;
            <a id="L1100"></a>}
        <a id="L1101"></a>}
        <a id="L1102"></a>comment = s.Comment;

    <a id="L1104"></a>case *ast.TypeSpec:
        <a id="L1105"></a>p.leadComment(s.Doc);
        <a id="L1106"></a>p.expr(s.Name, multiLine);
        <a id="L1107"></a>if n == 1 {
            <a id="L1108"></a>p.print(blank)
        <a id="L1109"></a>} else {
            <a id="L1110"></a>p.print(vtab)
        <a id="L1111"></a>}
        <a id="L1112"></a>optSemi = p.expr(s.Type, multiLine);
        <a id="L1113"></a>comment = s.Comment;

    <a id="L1115"></a>default:
        <a id="L1116"></a>panic(&#34;unreachable&#34;)
    <a id="L1117"></a>}

    <a id="L1119"></a>if context == inGroup || context == inStmtList &amp;&amp; !optSemi {
        <a id="L1120"></a>p.print(token.SEMICOLON)
    <a id="L1121"></a>}

    <a id="L1123"></a>if comment != nil {
        <a id="L1124"></a>for ; extraTabs &gt; 0; extraTabs-- {
            <a id="L1125"></a>p.print(vtab)
        <a id="L1126"></a>}
        <a id="L1127"></a>p.lineComment(comment);
    <a id="L1128"></a>}
<a id="L1129"></a>}


<a id="L1132"></a><span class="comment">// Sets multiLine to true if the declaration spans multiple lines.</span>
<a id="L1133"></a>func (p *printer) genDecl(d *ast.GenDecl, context declContext, multiLine *bool) {
    <a id="L1134"></a>p.leadComment(d.Doc);
    <a id="L1135"></a>p.print(d.Pos(), d.Tok, blank);

    <a id="L1137"></a>if d.Lparen.IsValid() {
        <a id="L1138"></a><span class="comment">// group of parenthesized declarations</span>
        <a id="L1139"></a>p.print(d.Lparen, token.LPAREN);
        <a id="L1140"></a>if len(d.Specs) &gt; 0 {
            <a id="L1141"></a>p.print(indent, formfeed);
            <a id="L1142"></a>var ml bool;
            <a id="L1143"></a>for i, s := range d.Specs {
                <a id="L1144"></a>if i &gt; 0 {
                    <a id="L1145"></a>p.linebreak(s.Pos().Line, 1, 2, ignore, ml)
                <a id="L1146"></a>}
                <a id="L1147"></a>ml = false;
                <a id="L1148"></a>p.spec(s, len(d.Specs), inGroup, &amp;ml);
            <a id="L1149"></a>}
            <a id="L1150"></a>p.print(unindent, formfeed);
            <a id="L1151"></a>*multiLine = true;
        <a id="L1152"></a>}
        <a id="L1153"></a>p.print(d.Rparen, token.RPAREN);

    <a id="L1155"></a>} else {
        <a id="L1156"></a><span class="comment">// single declaration</span>
        <a id="L1157"></a>p.spec(d.Specs[0], 1, context, multiLine)
    <a id="L1158"></a>}
<a id="L1159"></a>}


<a id="L1162"></a><span class="comment">// nodeSize determines the size of n in chars after formatting.</span>
<a id="L1163"></a><span class="comment">// The result is &lt;= maxSize if the node fits on one line with at</span>
<a id="L1164"></a><span class="comment">// most maxSize chars and the formatted output doesn&#39;t contain</span>
<a id="L1165"></a><span class="comment">// any control chars. Otherwise, the result is &gt; maxSize.</span>
<a id="L1166"></a><span class="comment">//</span>
<a id="L1167"></a>func (p *printer) nodeSize(n ast.Node, maxSize int) (size int) {
    <a id="L1168"></a>size = maxSize + 1; <span class="comment">// assume n doesn&#39;t fit</span>
    <a id="L1169"></a><span class="comment">// nodeSize computation must be indendent of particular</span>
    <a id="L1170"></a><span class="comment">// style so that we always get the same decision; print</span>
    <a id="L1171"></a><span class="comment">// in RawFormat</span>
    <a id="L1172"></a>cfg := Config{Mode: RawFormat};
    <a id="L1173"></a>var buf bytes.Buffer;
    <a id="L1174"></a>if _, err := cfg.Fprint(&amp;buf, n); err != nil {
        <a id="L1175"></a>return
    <a id="L1176"></a>}
    <a id="L1177"></a>if buf.Len() &lt;= maxSize {
        <a id="L1178"></a>for _, ch := range buf.Bytes() {
            <a id="L1179"></a>if ch &lt; &#39; &#39; {
                <a id="L1180"></a>return
            <a id="L1181"></a>}
        <a id="L1182"></a>}
        <a id="L1183"></a>size = buf.Len(); <span class="comment">// n fits</span>
    <a id="L1184"></a>}
    <a id="L1185"></a>return;
<a id="L1186"></a>}


<a id="L1189"></a>func (p *printer) isOneLineFunc(b *ast.BlockStmt, headerSize int) bool {
    <a id="L1190"></a>const maxSize = 90; <span class="comment">// adjust as appropriate, this is an approximate value</span>
    <a id="L1191"></a>bodySize := 0;
    <a id="L1192"></a>switch {
    <a id="L1193"></a>case len(b.List) &gt; 1 || p.commentBefore(b.Rbrace):
        <a id="L1194"></a>return false <span class="comment">// too many statements or there is a comment - all bets are off</span>
    <a id="L1195"></a>case len(b.List) == 1:
        <a id="L1196"></a>bodySize = p.nodeSize(b.List[0], maxSize)
    <a id="L1197"></a>}
    <a id="L1198"></a><span class="comment">// require both headers and overall size to be not &#34;too large&#34;</span>
    <a id="L1199"></a>return headerSize &lt;= maxSize/2 &amp;&amp; headerSize+bodySize &lt;= maxSize;
<a id="L1200"></a>}


<a id="L1203"></a><span class="comment">// Sets multiLine to true if the function body spans multiple lines.</span>
<a id="L1204"></a>func (p *printer) funcBody(b *ast.BlockStmt, headerSize int, isLit bool, multiLine *bool) {
    <a id="L1205"></a>if b == nil {
        <a id="L1206"></a>return
    <a id="L1207"></a>}

    <a id="L1209"></a>if p.isOneLineFunc(b, headerSize) {
        <a id="L1210"></a>sep := vtab;
        <a id="L1211"></a>if isLit {
            <a id="L1212"></a>sep = blank
        <a id="L1213"></a>}
        <a id="L1214"></a>if len(b.List) &gt; 0 {
            <a id="L1215"></a>p.print(sep, b.Pos(), token.LBRACE, blank);
            <a id="L1216"></a>p.stmt(b.List[0], ignoreMultiLine);
            <a id="L1217"></a>p.print(blank, b.Rbrace, token.RBRACE);
        <a id="L1218"></a>} else {
            <a id="L1219"></a>p.print(sep, b.Pos(), token.LBRACE, b.Rbrace, token.RBRACE)
        <a id="L1220"></a>}
        <a id="L1221"></a>return;
    <a id="L1222"></a>}

    <a id="L1224"></a>p.print(blank);
    <a id="L1225"></a>p.block(b, 1);
    <a id="L1226"></a>*multiLine = true;
<a id="L1227"></a>}


<a id="L1230"></a><span class="comment">// distance returns the column difference between from and to if both</span>
<a id="L1231"></a><span class="comment">// are on the same line; if they are on different lines (or unknown)</span>
<a id="L1232"></a><span class="comment">// the result is infinity (1&lt;&lt;30).</span>
<a id="L1233"></a>func distance(from, to token.Position) int {
    <a id="L1234"></a>if from.IsValid() &amp;&amp; to.IsValid() &amp;&amp; from.Line == to.Line {
        <a id="L1235"></a>return to.Column - from.Column
    <a id="L1236"></a>}
    <a id="L1237"></a>return 1 &lt;&lt; 30;
<a id="L1238"></a>}


<a id="L1241"></a><span class="comment">// Sets multiLine to true if the declaration spans multiple lines.</span>
<a id="L1242"></a>func (p *printer) funcDecl(d *ast.FuncDecl, multiLine *bool) {
    <a id="L1243"></a>p.leadComment(d.Doc);
    <a id="L1244"></a>p.print(d.Pos(), token.FUNC, blank);
    <a id="L1245"></a>if recv := d.Recv; recv != nil {
        <a id="L1246"></a><span class="comment">// method: print receiver</span>
        <a id="L1247"></a>p.print(token.LPAREN);
        <a id="L1248"></a>if len(recv.Names) &gt; 0 {
            <a id="L1249"></a>p.expr(recv.Names[0], multiLine);
            <a id="L1250"></a>p.print(blank);
        <a id="L1251"></a>}
        <a id="L1252"></a>p.expr(recv.Type, multiLine);
        <a id="L1253"></a>p.print(token.RPAREN, blank);
    <a id="L1254"></a>}
    <a id="L1255"></a>p.expr(d.Name, multiLine);
    <a id="L1256"></a>p.signature(d.Type.Params, d.Type.Results, multiLine);
    <a id="L1257"></a>p.funcBody(d.Body, distance(d.Pos(), p.pos), false, multiLine);
<a id="L1258"></a>}


<a id="L1261"></a><span class="comment">// Sets multiLine to true if the declaration spans multiple lines.</span>
<a id="L1262"></a>func (p *printer) decl(decl ast.Decl, context declContext, multiLine *bool) {
    <a id="L1263"></a>switch d := decl.(type) {
    <a id="L1264"></a>case *ast.BadDecl:
        <a id="L1265"></a>p.print(d.Pos(), &#34;BadDecl&#34;)
    <a id="L1266"></a>case *ast.GenDecl:
        <a id="L1267"></a>p.genDecl(d, context, multiLine)
    <a id="L1268"></a>case *ast.FuncDecl:
        <a id="L1269"></a>p.funcDecl(d, multiLine)
    <a id="L1270"></a>default:
        <a id="L1271"></a>panic(&#34;unreachable&#34;)
    <a id="L1272"></a>}
<a id="L1273"></a>}


<a id="L1276"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L1277"></a><span class="comment">// Files</span>

<a id="L1279"></a>const maxDeclNewlines = 3 <span class="comment">// maximum number of newlines between declarations</span>

<a id="L1281"></a>func declToken(decl ast.Decl) (tok token.Token) {
    <a id="L1282"></a>tok = token.ILLEGAL;
    <a id="L1283"></a>switch d := decl.(type) {
    <a id="L1284"></a>case *ast.GenDecl:
        <a id="L1285"></a>tok = d.Tok
    <a id="L1286"></a>case *ast.FuncDecl:
        <a id="L1287"></a>tok = token.FUNC
    <a id="L1288"></a>}
    <a id="L1289"></a>return;
<a id="L1290"></a>}


<a id="L1293"></a>func (p *printer) file(src *ast.File) {
    <a id="L1294"></a>p.leadComment(src.Doc);
    <a id="L1295"></a>p.print(src.Pos(), token.PACKAGE, blank);
    <a id="L1296"></a>p.expr(src.Name, ignoreMultiLine);

    <a id="L1298"></a>if len(src.Decls) &gt; 0 {
        <a id="L1299"></a>tok := token.ILLEGAL;
        <a id="L1300"></a>for _, d := range src.Decls {
            <a id="L1301"></a>prev := tok;
            <a id="L1302"></a>tok = declToken(d);
            <a id="L1303"></a><span class="comment">// if the declaration token changed (e.g., from CONST to TYPE)</span>
            <a id="L1304"></a><span class="comment">// print an empty line between top-level declarations</span>
            <a id="L1305"></a>min := 1;
            <a id="L1306"></a>if prev != tok {
                <a id="L1307"></a>min = 2
            <a id="L1308"></a>}
            <a id="L1309"></a>p.linebreak(d.Pos().Line, min, maxDeclNewlines, ignore, false);
            <a id="L1310"></a>p.decl(d, atTop, ignoreMultiLine);
        <a id="L1311"></a>}
    <a id="L1312"></a>}

    <a id="L1314"></a>p.print(newline);
<a id="L1315"></a>}
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
