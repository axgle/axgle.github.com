<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/datafmt/parser.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/datafmt/parser.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package datafmt

<a id="L7"></a>import (
    <a id="L8"></a>&#34;container/vector&#34;;
    <a id="L9"></a>&#34;go/scanner&#34;;
    <a id="L10"></a>&#34;go/token&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;strconv&#34;;
    <a id="L13"></a>&#34;strings&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L17"></a><span class="comment">// Parsing</span>

<a id="L19"></a>type parser struct {
    <a id="L20"></a>scanner.ErrorVector;
    <a id="L21"></a>scanner scanner.Scanner;
    <a id="L22"></a>pos     token.Position; <span class="comment">// token position</span>
    <a id="L23"></a>tok     token.Token;    <span class="comment">// one token look-ahead</span>
    <a id="L24"></a>lit     []byte;         <span class="comment">// token literal</span>

    <a id="L26"></a>packs map[string]string; <span class="comment">// PackageName -&gt; ImportPath</span>
    <a id="L27"></a>rules map[string]expr;   <span class="comment">// RuleName -&gt; Expression</span>
<a id="L28"></a>}


<a id="L31"></a>func (p *parser) next() {
    <a id="L32"></a>p.pos, p.tok, p.lit = p.scanner.Scan();
    <a id="L33"></a>switch p.tok {
    <a id="L34"></a>case token.CHAN, token.FUNC, token.INTERFACE, token.MAP, token.STRUCT:
        <a id="L35"></a><span class="comment">// Go keywords for composite types are type names</span>
        <a id="L36"></a><span class="comment">// returned by reflect. Accept them as identifiers.</span>
        <a id="L37"></a>p.tok = token.IDENT <span class="comment">// p.lit is already set correctly</span>
    <a id="L38"></a>}
<a id="L39"></a>}


<a id="L42"></a>func (p *parser) init(filename string, src []byte) {
    <a id="L43"></a>p.ErrorVector.Init();
    <a id="L44"></a>p.scanner.Init(filename, src, p, scanner.AllowIllegalChars); <span class="comment">// return &#39;@&#39; as token.ILLEGAL w/o error message</span>
    <a id="L45"></a>p.next();                                                    <span class="comment">// initializes pos, tok, lit</span>
    <a id="L46"></a>p.packs = make(map[string]string);
    <a id="L47"></a>p.rules = make(map[string]expr);
<a id="L48"></a>}


<a id="L51"></a>func (p *parser) errorExpected(pos token.Position, msg string) {
    <a id="L52"></a>msg = &#34;expected &#34; + msg;
    <a id="L53"></a>if pos.Offset == p.pos.Offset {
        <a id="L54"></a><span class="comment">// the error happened at the current position;</span>
        <a id="L55"></a><span class="comment">// make the error message more specific</span>
        <a id="L56"></a>msg += &#34;, found &#39;&#34; + p.tok.String() + &#34;&#39;&#34;;
        <a id="L57"></a>if p.tok.IsLiteral() {
            <a id="L58"></a>msg += &#34; &#34; + string(p.lit)
        <a id="L59"></a>}
    <a id="L60"></a>}
    <a id="L61"></a>p.Error(pos, msg);
<a id="L62"></a>}


<a id="L65"></a>func (p *parser) expect(tok token.Token) token.Position {
    <a id="L66"></a>pos := p.pos;
    <a id="L67"></a>if p.tok != tok {
        <a id="L68"></a>p.errorExpected(pos, &#34;&#39;&#34;+tok.String()+&#34;&#39;&#34;)
    <a id="L69"></a>}
    <a id="L70"></a>p.next(); <span class="comment">// make progress in any case</span>
    <a id="L71"></a>return pos;
<a id="L72"></a>}


<a id="L75"></a>func (p *parser) parseIdentifier() string {
    <a id="L76"></a>name := string(p.lit);
    <a id="L77"></a>p.expect(token.IDENT);
    <a id="L78"></a>return name;
<a id="L79"></a>}


<a id="L82"></a>func (p *parser) parseTypeName() (string, bool) {
    <a id="L83"></a>pos := p.pos;
    <a id="L84"></a>name, isIdent := p.parseIdentifier(), true;
    <a id="L85"></a>if p.tok == token.PERIOD {
        <a id="L86"></a><span class="comment">// got a package name, lookup package</span>
        <a id="L87"></a>if importPath, found := p.packs[name]; found {
            <a id="L88"></a>name = importPath
        <a id="L89"></a>} else {
            <a id="L90"></a>p.Error(pos, &#34;package not declared: &#34;+name)
        <a id="L91"></a>}
        <a id="L92"></a>p.next();
        <a id="L93"></a>name, isIdent = name+&#34;.&#34;+p.parseIdentifier(), false;
    <a id="L94"></a>}
    <a id="L95"></a>return name, isIdent;
<a id="L96"></a>}


<a id="L99"></a><span class="comment">// Parses a rule name and returns it. If the rule name is</span>
<a id="L100"></a><span class="comment">// a package-qualified type name, the package name is resolved.</span>
<a id="L101"></a><span class="comment">// The 2nd result value is true iff the rule name consists of a</span>
<a id="L102"></a><span class="comment">// single identifier only (and thus could be a package name).</span>
<a id="L103"></a><span class="comment">//</span>
<a id="L104"></a>func (p *parser) parseRuleName() (string, bool) {
    <a id="L105"></a>name, isIdent := &#34;&#34;, false;
    <a id="L106"></a>switch p.tok {
    <a id="L107"></a>case token.IDENT:
        <a id="L108"></a>name, isIdent = p.parseTypeName()
    <a id="L109"></a>case token.DEFAULT:
        <a id="L110"></a>name = &#34;default&#34;;
        <a id="L111"></a>p.next();
    <a id="L112"></a>case token.QUO:
        <a id="L113"></a>name = &#34;/&#34;;
        <a id="L114"></a>p.next();
    <a id="L115"></a>default:
        <a id="L116"></a>p.errorExpected(p.pos, &#34;rule name&#34;);
        <a id="L117"></a>p.next(); <span class="comment">// make progress in any case</span>
    <a id="L118"></a>}
    <a id="L119"></a>return name, isIdent;
<a id="L120"></a>}


<a id="L123"></a>func (p *parser) parseString() string {
    <a id="L124"></a>s := &#34;&#34;;
    <a id="L125"></a>if p.tok == token.STRING {
        <a id="L126"></a>s, _ = strconv.Unquote(string(p.lit));
        <a id="L127"></a><span class="comment">// Unquote may fail with an error, but only if the scanner found</span>
        <a id="L128"></a><span class="comment">// an illegal string in the first place. In this case the error</span>
        <a id="L129"></a><span class="comment">// has already been reported.</span>
        <a id="L130"></a>p.next();
        <a id="L131"></a>return s;
    <a id="L132"></a>} else {
        <a id="L133"></a>p.expect(token.STRING)
    <a id="L134"></a>}
    <a id="L135"></a>return s;
<a id="L136"></a>}


<a id="L139"></a>func (p *parser) parseLiteral() literal {
    <a id="L140"></a>s := strings.Bytes(p.parseString());

    <a id="L142"></a><span class="comment">// A string literal may contain %-format specifiers. To simplify</span>
    <a id="L143"></a><span class="comment">// and speed up printing of the literal, split it into segments</span>
    <a id="L144"></a><span class="comment">// that start with &#34;%&#34; possibly followed by a last segment that</span>
    <a id="L145"></a><span class="comment">// starts with some other character.</span>
    <a id="L146"></a>var list vector.Vector;
    <a id="L147"></a>list.Init(0);
    <a id="L148"></a>i0 := 0;
    <a id="L149"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L150"></a>if s[i] == &#39;%&#39; &amp;&amp; i+1 &lt; len(s) {
            <a id="L151"></a><span class="comment">// the next segment starts with a % format</span>
            <a id="L152"></a>if i0 &lt; i {
                <a id="L153"></a><span class="comment">// the current segment is not empty, split it off</span>
                <a id="L154"></a>list.Push(s[i0:i]);
                <a id="L155"></a>i0 = i;
            <a id="L156"></a>}
            <a id="L157"></a>i++; <span class="comment">// skip %; let loop skip over char after %</span>
        <a id="L158"></a>}
    <a id="L159"></a>}
    <a id="L160"></a><span class="comment">// the final segment may start with any character</span>
    <a id="L161"></a><span class="comment">// (it is empty iff the string is empty)</span>
    <a id="L162"></a>list.Push(s[i0:len(s)]);

    <a id="L164"></a><span class="comment">// convert list into a literal</span>
    <a id="L165"></a>lit := make(literal, list.Len());
    <a id="L166"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L167"></a>lit[i] = list.At(i).([]byte)
    <a id="L168"></a>}

    <a id="L170"></a>return lit;
<a id="L171"></a>}


<a id="L174"></a>func (p *parser) parseField() expr {
    <a id="L175"></a>var fname string;
    <a id="L176"></a>switch p.tok {
    <a id="L177"></a>case token.ILLEGAL:
        <a id="L178"></a>if string(p.lit) != &#34;@&#34; {
            <a id="L179"></a>return nil
        <a id="L180"></a>}
        <a id="L181"></a>fname = &#34;@&#34;;
        <a id="L182"></a>p.next();
    <a id="L183"></a>case token.MUL:
        <a id="L184"></a>fname = &#34;*&#34;;
        <a id="L185"></a>p.next();
    <a id="L186"></a>case token.IDENT:
        <a id="L187"></a>fname = p.parseIdentifier()
    <a id="L188"></a>default:
        <a id="L189"></a>return nil
    <a id="L190"></a>}

    <a id="L192"></a>var ruleName string;
    <a id="L193"></a>if p.tok == token.COLON {
        <a id="L194"></a>p.next();
        <a id="L195"></a>ruleName, _ = p.parseRuleName();
    <a id="L196"></a>}

    <a id="L198"></a>return &amp;field{fname, ruleName};
<a id="L199"></a>}


<a id="L202"></a>func (p *parser) parseOperand() (x expr) {
    <a id="L203"></a>switch p.tok {
    <a id="L204"></a>case token.STRING:
        <a id="L205"></a>x = p.parseLiteral()

    <a id="L207"></a>case token.LPAREN:
        <a id="L208"></a>p.next();
        <a id="L209"></a>x = p.parseExpression();
        <a id="L210"></a>if p.tok == token.SHR {
            <a id="L211"></a>p.next();
            <a id="L212"></a>x = &amp;group{x, p.parseExpression()};
        <a id="L213"></a>}
        <a id="L214"></a>p.expect(token.RPAREN);

    <a id="L216"></a>case token.LBRACK:
        <a id="L217"></a>p.next();
        <a id="L218"></a>x = &amp;option{p.parseExpression()};
        <a id="L219"></a>p.expect(token.RBRACK);

    <a id="L221"></a>case token.LBRACE:
        <a id="L222"></a>p.next();
        <a id="L223"></a>x = p.parseExpression();
        <a id="L224"></a>var div expr;
        <a id="L225"></a>if p.tok == token.QUO {
            <a id="L226"></a>p.next();
            <a id="L227"></a>div = p.parseExpression();
        <a id="L228"></a>}
        <a id="L229"></a>x = &amp;repetition{x, div};
        <a id="L230"></a>p.expect(token.RBRACE);

    <a id="L232"></a>default:
        <a id="L233"></a>x = p.parseField() <span class="comment">// may be nil</span>
    <a id="L234"></a>}

    <a id="L236"></a>return x;
<a id="L237"></a>}


<a id="L240"></a>func (p *parser) parseSequence() expr {
    <a id="L241"></a>var list vector.Vector;
    <a id="L242"></a>list.Init(0);

    <a id="L244"></a>for x := p.parseOperand(); x != nil; x = p.parseOperand() {
        <a id="L245"></a>list.Push(x)
    <a id="L246"></a>}

    <a id="L248"></a><span class="comment">// no need for a sequence if list.Len() &lt; 2</span>
    <a id="L249"></a>switch list.Len() {
    <a id="L250"></a>case 0:
        <a id="L251"></a>return nil
    <a id="L252"></a>case 1:
        <a id="L253"></a>return list.At(0).(expr)
    <a id="L254"></a>}

    <a id="L256"></a><span class="comment">// convert list into a sequence</span>
    <a id="L257"></a>seq := make(sequence, list.Len());
    <a id="L258"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L259"></a>seq[i] = list.At(i).(expr)
    <a id="L260"></a>}
    <a id="L261"></a>return seq;
<a id="L262"></a>}


<a id="L265"></a>func (p *parser) parseExpression() expr {
    <a id="L266"></a>var list vector.Vector;
    <a id="L267"></a>list.Init(0);

    <a id="L269"></a>for {
        <a id="L270"></a>x := p.parseSequence();
        <a id="L271"></a>if x != nil {
            <a id="L272"></a>list.Push(x)
        <a id="L273"></a>}
        <a id="L274"></a>if p.tok != token.OR {
            <a id="L275"></a>break
        <a id="L276"></a>}
        <a id="L277"></a>p.next();
    <a id="L278"></a>}

    <a id="L280"></a><span class="comment">// no need for an alternatives if list.Len() &lt; 2</span>
    <a id="L281"></a>switch list.Len() {
    <a id="L282"></a>case 0:
        <a id="L283"></a>return nil
    <a id="L284"></a>case 1:
        <a id="L285"></a>return list.At(0).(expr)
    <a id="L286"></a>}

    <a id="L288"></a><span class="comment">// convert list into a alternatives</span>
    <a id="L289"></a>alt := make(alternatives, list.Len());
    <a id="L290"></a>for i := 0; i &lt; list.Len(); i++ {
        <a id="L291"></a>alt[i] = list.At(i).(expr)
    <a id="L292"></a>}
    <a id="L293"></a>return alt;
<a id="L294"></a>}


<a id="L297"></a>func (p *parser) parseFormat() {
    <a id="L298"></a>for p.tok != token.EOF {
        <a id="L299"></a>pos := p.pos;

        <a id="L301"></a>name, isIdent := p.parseRuleName();
        <a id="L302"></a>switch p.tok {
        <a id="L303"></a>case token.STRING:
            <a id="L304"></a><span class="comment">// package declaration</span>
            <a id="L305"></a>importPath := p.parseString();

            <a id="L307"></a><span class="comment">// add package declaration</span>
            <a id="L308"></a>if !isIdent {
                <a id="L309"></a>p.Error(pos, &#34;illegal package name: &#34;+name)
            <a id="L310"></a>} else if _, found := p.packs[name]; !found {
                <a id="L311"></a>p.packs[name] = importPath
            <a id="L312"></a>} else {
                <a id="L313"></a>p.Error(pos, &#34;package already declared: &#34;+name)
            <a id="L314"></a>}

        <a id="L316"></a>case token.ASSIGN:
            <a id="L317"></a><span class="comment">// format rule</span>
            <a id="L318"></a>p.next();
            <a id="L319"></a>x := p.parseExpression();

            <a id="L321"></a><span class="comment">// add rule</span>
            <a id="L322"></a>if _, found := p.rules[name]; !found {
                <a id="L323"></a>p.rules[name] = x
            <a id="L324"></a>} else {
                <a id="L325"></a>p.Error(pos, &#34;format rule already declared: &#34;+name)
            <a id="L326"></a>}

        <a id="L328"></a>default:
            <a id="L329"></a>p.errorExpected(p.pos, &#34;package declaration or format rule&#34;);
            <a id="L330"></a>p.next(); <span class="comment">// make progress in any case</span>
        <a id="L331"></a>}

        <a id="L333"></a>if p.tok == token.SEMICOLON {
            <a id="L334"></a>p.next()
        <a id="L335"></a>} else {
            <a id="L336"></a>break
        <a id="L337"></a>}
    <a id="L338"></a>}
    <a id="L339"></a>p.expect(token.EOF);
<a id="L340"></a>}


<a id="L343"></a>func remap(p *parser, name string) string {
    <a id="L344"></a>i := strings.Index(name, &#34;.&#34;);
    <a id="L345"></a>if i &gt;= 0 {
        <a id="L346"></a>packageName, suffix := name[0:i], name[i:len(name)];
        <a id="L347"></a><span class="comment">// lookup package</span>
        <a id="L348"></a>if importPath, found := p.packs[packageName]; found {
            <a id="L349"></a>name = importPath + suffix
        <a id="L350"></a>} else {
            <a id="L351"></a>var invalidPos token.Position;
            <a id="L352"></a>p.Error(invalidPos, &#34;package not declared: &#34;+packageName);
        <a id="L353"></a>}
    <a id="L354"></a>}
    <a id="L355"></a>return name;
<a id="L356"></a>}


<a id="L359"></a><span class="comment">// Parse parses a set of format productions from source src. Custom</span>
<a id="L360"></a><span class="comment">// formatters may be provided via a map of formatter functions. If</span>
<a id="L361"></a><span class="comment">// there are no errors, the result is a Format and the error is nil.</span>
<a id="L362"></a><span class="comment">// Otherwise the format is nil and a non-empty ErrorList is returned.</span>
<a id="L363"></a><span class="comment">//</span>
<a id="L364"></a>func Parse(filename string, src []byte, fmap FormatterMap) (Format, os.Error) {
    <a id="L365"></a><span class="comment">// parse source</span>
    <a id="L366"></a>var p parser;
    <a id="L367"></a>p.init(filename, src);
    <a id="L368"></a>p.parseFormat();

    <a id="L370"></a><span class="comment">// add custom formatters, if any</span>
    <a id="L371"></a>for name, form := range fmap {
        <a id="L372"></a>name = remap(&amp;p, name);
        <a id="L373"></a>if _, found := p.rules[name]; !found {
            <a id="L374"></a>p.rules[name] = &amp;custom{name, form}
        <a id="L375"></a>} else {
            <a id="L376"></a>var invalidPos token.Position;
            <a id="L377"></a>p.Error(invalidPos, &#34;formatter already declared: &#34;+name);
        <a id="L378"></a>}
    <a id="L379"></a>}

    <a id="L381"></a>return p.rules, p.GetError(scanner.NoMultiples);
<a id="L382"></a>}
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
