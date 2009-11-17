<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/ebnf/parser.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/ebnf/parser.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ebnf

<a id="L7"></a>import (
    <a id="L8"></a>&#34;container/vector&#34;;
    <a id="L9"></a>&#34;go/scanner&#34;;
    <a id="L10"></a>&#34;go/token&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;strconv&#34;;
<a id="L13"></a>)


<a id="L16"></a>type parser struct {
    <a id="L17"></a>scanner.ErrorVector;
    <a id="L18"></a>scanner scanner.Scanner;
    <a id="L19"></a>pos     token.Position; <span class="comment">// token position</span>
    <a id="L20"></a>tok     token.Token;    <span class="comment">// one token look-ahead</span>
    <a id="L21"></a>lit     []byte;         <span class="comment">// token literal</span>
<a id="L22"></a>}


<a id="L25"></a>func (p *parser) next() {
    <a id="L26"></a>p.pos, p.tok, p.lit = p.scanner.Scan();
    <a id="L27"></a>if p.tok.IsKeyword() {
        <a id="L28"></a><span class="comment">// TODO Should keyword mapping always happen outside scanner?</span>
        <a id="L29"></a><span class="comment">//      Or should there be a flag to scanner to enable keyword mapping?</span>
        <a id="L30"></a>p.tok = token.IDENT
    <a id="L31"></a>}
<a id="L32"></a>}


<a id="L35"></a>func (p *parser) errorExpected(pos token.Position, msg string) {
    <a id="L36"></a>msg = &#34;expected &#34; + msg;
    <a id="L37"></a>if pos.Offset == p.pos.Offset {
        <a id="L38"></a><span class="comment">// the error happened at the current position;</span>
        <a id="L39"></a><span class="comment">// make the error message more specific</span>
        <a id="L40"></a>msg += &#34;, found &#39;&#34; + p.tok.String() + &#34;&#39;&#34;;
        <a id="L41"></a>if p.tok.IsLiteral() {
            <a id="L42"></a>msg += &#34; &#34; + string(p.lit)
        <a id="L43"></a>}
    <a id="L44"></a>}
    <a id="L45"></a>p.Error(pos, msg);
<a id="L46"></a>}


<a id="L49"></a>func (p *parser) expect(tok token.Token) token.Position {
    <a id="L50"></a>pos := p.pos;
    <a id="L51"></a>if p.tok != tok {
        <a id="L52"></a>p.errorExpected(pos, &#34;&#39;&#34;+tok.String()+&#34;&#39;&#34;)
    <a id="L53"></a>}
    <a id="L54"></a>p.next(); <span class="comment">// make progress in any case</span>
    <a id="L55"></a>return pos;
<a id="L56"></a>}


<a id="L59"></a>func (p *parser) parseIdentifier() *Name {
    <a id="L60"></a>pos := p.pos;
    <a id="L61"></a>name := string(p.lit);
    <a id="L62"></a>p.expect(token.IDENT);
    <a id="L63"></a>return &amp;Name{pos, name};
<a id="L64"></a>}


<a id="L67"></a>func (p *parser) parseToken() *Token {
    <a id="L68"></a>pos := p.pos;
    <a id="L69"></a>value := &#34;&#34;;
    <a id="L70"></a>if p.tok == token.STRING {
        <a id="L71"></a>value, _ = strconv.Unquote(string(p.lit));
        <a id="L72"></a><span class="comment">// Unquote may fail with an error, but only if the scanner found</span>
        <a id="L73"></a><span class="comment">// an illegal string in the first place. In this case the error</span>
        <a id="L74"></a><span class="comment">// has already been reported.</span>
        <a id="L75"></a>p.next();
    <a id="L76"></a>} else {
        <a id="L77"></a>p.expect(token.STRING)
    <a id="L78"></a>}
    <a id="L79"></a>return &amp;Token{pos, value};
<a id="L80"></a>}


<a id="L83"></a>func (p *parser) parseTerm() (x Expression) {
    <a id="L84"></a>pos := p.pos;

    <a id="L86"></a>switch p.tok {
    <a id="L87"></a>case token.IDENT:
        <a id="L88"></a>x = p.parseIdentifier()

    <a id="L90"></a>case token.STRING:
        <a id="L91"></a>tok := p.parseToken();
        <a id="L92"></a>x = tok;
        <a id="L93"></a>if p.tok == token.ELLIPSIS {
            <a id="L94"></a>p.next();
            <a id="L95"></a>x = &amp;Range{tok, p.parseToken()};
        <a id="L96"></a>}

    <a id="L98"></a>case token.LPAREN:
        <a id="L99"></a>p.next();
        <a id="L100"></a>x = &amp;Group{pos, p.parseExpression()};
        <a id="L101"></a>p.expect(token.RPAREN);

    <a id="L103"></a>case token.LBRACK:
        <a id="L104"></a>p.next();
        <a id="L105"></a>x = &amp;Option{pos, p.parseExpression()};
        <a id="L106"></a>p.expect(token.RBRACK);

    <a id="L108"></a>case token.LBRACE:
        <a id="L109"></a>p.next();
        <a id="L110"></a>x = &amp;Repetition{pos, p.parseExpression()};
        <a id="L111"></a>p.expect(token.RBRACE);
    <a id="L112"></a>}

    <a id="L114"></a>return x;
<a id="L115"></a>}


<a id="L118"></a>func (p *parser) parseSequence() Expression {
    <a id="L119"></a>var list vector.Vector;
    <a id="L120"></a>list.Init(0);

    <a id="L122"></a>for x := p.parseTerm(); x != nil; x = p.parseTerm() {
        <a id="L123"></a>list.Push(x)
    <a id="L124"></a>}

    <a id="L126"></a><span class="comment">// no need for a sequence if list.Len() &lt; 2</span>
    <a id="L127"></a>switch list.Len() {
    <a id="L128"></a>case 0:
        <a id="L129"></a>return nil
    <a id="L130"></a>case 1:
        <a id="L131"></a>return list.At(0).(Expression)
    <a id="L132"></a>}

    <a id="L134"></a><span class="comment">// convert list into a sequence</span>
    <a id="L135"></a>seq := make(Sequence, list.Len());
    <a id="L136"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L137"></a>seq[i] = list.At(i).(Expression)
    <a id="L138"></a>}
    <a id="L139"></a>return seq;
<a id="L140"></a>}


<a id="L143"></a>func (p *parser) parseExpression() Expression {
    <a id="L144"></a>var list vector.Vector;
    <a id="L145"></a>list.Init(0);

    <a id="L147"></a>for {
        <a id="L148"></a>x := p.parseSequence();
        <a id="L149"></a>if x != nil {
            <a id="L150"></a>list.Push(x)
        <a id="L151"></a>}
        <a id="L152"></a>if p.tok != token.OR {
            <a id="L153"></a>break
        <a id="L154"></a>}
        <a id="L155"></a>p.next();
    <a id="L156"></a>}

    <a id="L158"></a><span class="comment">// no need for an Alternative node if list.Len() &lt; 2</span>
    <a id="L159"></a>switch list.Len() {
    <a id="L160"></a>case 0:
        <a id="L161"></a>return nil
    <a id="L162"></a>case 1:
        <a id="L163"></a>return list.At(0).(Expression)
    <a id="L164"></a>}

    <a id="L166"></a><span class="comment">// convert list into an Alternative node</span>
    <a id="L167"></a>alt := make(Alternative, list.Len());
    <a id="L168"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L169"></a>alt[i] = list.At(i).(Expression)
    <a id="L170"></a>}
    <a id="L171"></a>return alt;
<a id="L172"></a>}


<a id="L175"></a>func (p *parser) parseProduction() *Production {
    <a id="L176"></a>name := p.parseIdentifier();
    <a id="L177"></a>p.expect(token.ASSIGN);
    <a id="L178"></a>expr := p.parseExpression();
    <a id="L179"></a>p.expect(token.PERIOD);
    <a id="L180"></a>return &amp;Production{name, expr};
<a id="L181"></a>}


<a id="L184"></a>func (p *parser) parse(filename string, src []byte) Grammar {
    <a id="L185"></a><span class="comment">// initialize parser</span>
    <a id="L186"></a>p.ErrorVector.Init();
    <a id="L187"></a>p.scanner.Init(filename, src, p, 0);
    <a id="L188"></a>p.next(); <span class="comment">// initializes pos, tok, lit</span>

    <a id="L190"></a>grammar := make(Grammar);
    <a id="L191"></a>for p.tok != token.EOF {
        <a id="L192"></a>prod := p.parseProduction();
        <a id="L193"></a>name := prod.Name.String;
        <a id="L194"></a>if _, found := grammar[name]; !found {
            <a id="L195"></a>grammar[name] = prod
        <a id="L196"></a>} else {
            <a id="L197"></a>p.Error(prod.Pos(), name+&#34; declared already&#34;)
        <a id="L198"></a>}
    <a id="L199"></a>}

    <a id="L201"></a>return grammar;
<a id="L202"></a>}


<a id="L205"></a><span class="comment">// Parse parses a set of EBNF productions from source src.</span>
<a id="L206"></a><span class="comment">// It returns a set of productions. Errors are reported</span>
<a id="L207"></a><span class="comment">// for incorrect syntax and if a production is declared</span>
<a id="L208"></a><span class="comment">// more than once.</span>
<a id="L209"></a><span class="comment">//</span>
<a id="L210"></a>func Parse(filename string, src []byte) (Grammar, os.Error) {
    <a id="L211"></a>var p parser;
    <a id="L212"></a>grammar := p.parse(filename, src);
    <a id="L213"></a>return grammar, p.GetError(scanner.Sorted);
<a id="L214"></a>}
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
