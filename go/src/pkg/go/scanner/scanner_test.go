<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/scanner/scanner_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/go/scanner/scanner_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package scanner

<a id="L7"></a>import (
    <a id="L8"></a>&#34;go/token&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;strings&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)


<a id="L15"></a>const <span class="comment">/* class */</span> (
    <a id="L16"></a>special = iota;
    <a id="L17"></a>literal;
    <a id="L18"></a>operator;
    <a id="L19"></a>keyword;
<a id="L20"></a>)


<a id="L23"></a>func tokenclass(tok token.Token) int {
    <a id="L24"></a>switch {
    <a id="L25"></a>case tok.IsLiteral():
        <a id="L26"></a>return literal
    <a id="L27"></a>case tok.IsOperator():
        <a id="L28"></a>return operator
    <a id="L29"></a>case tok.IsKeyword():
        <a id="L30"></a>return keyword
    <a id="L31"></a>}
    <a id="L32"></a>return special;
<a id="L33"></a>}


<a id="L36"></a>type elt struct {
    <a id="L37"></a>tok   token.Token;
    <a id="L38"></a>lit   string;
    <a id="L39"></a>class int;
<a id="L40"></a>}


<a id="L43"></a>var tokens = [...]elt{
    <a id="L44"></a><span class="comment">// Special tokens</span>
    <a id="L45"></a>elt{token.COMMENT, &#34;/* a comment */&#34;, special},
    <a id="L46"></a>elt{token.COMMENT, &#34;// a comment \n&#34;, special},

    <a id="L48"></a><span class="comment">// Identifiers and basic type literals</span>
    <a id="L49"></a>elt{token.IDENT, &#34;foobar&#34;, literal},
    <a id="L50"></a>elt{token.IDENT, &#34;a۰۱۸&#34;, literal},
    <a id="L51"></a>elt{token.IDENT, &#34;foo६४&#34;, literal},
    <a id="L52"></a>elt{token.IDENT, &#34;bar９８７６&#34;, literal},
    <a id="L53"></a>elt{token.INT, &#34;0&#34;, literal},
    <a id="L54"></a>elt{token.INT, &#34;01234567&#34;, literal},
    <a id="L55"></a>elt{token.INT, &#34;0xcafebabe&#34;, literal},
    <a id="L56"></a>elt{token.FLOAT, &#34;0.&#34;, literal},
    <a id="L57"></a>elt{token.FLOAT, &#34;.0&#34;, literal},
    <a id="L58"></a>elt{token.FLOAT, &#34;3.14159265&#34;, literal},
    <a id="L59"></a>elt{token.FLOAT, &#34;1e0&#34;, literal},
    <a id="L60"></a>elt{token.FLOAT, &#34;1e+100&#34;, literal},
    <a id="L61"></a>elt{token.FLOAT, &#34;1e-100&#34;, literal},
    <a id="L62"></a>elt{token.FLOAT, &#34;2.71828e-1000&#34;, literal},
    <a id="L63"></a>elt{token.CHAR, &#34;&#39;a&#39;&#34;, literal},
    <a id="L64"></a>elt{token.CHAR, &#34;&#39;\\000&#39;&#34;, literal},
    <a id="L65"></a>elt{token.CHAR, &#34;&#39;\\xFF&#39;&#34;, literal},
    <a id="L66"></a>elt{token.CHAR, &#34;&#39;\\uff16&#39;&#34;, literal},
    <a id="L67"></a>elt{token.CHAR, &#34;&#39;\\U0000ff16&#39;&#34;, literal},
    <a id="L68"></a>elt{token.STRING, &#34;`foobar`&#34;, literal},
    <a id="L69"></a>elt{token.STRING, &#34;`&#34; `foo
	                        bar`
        <a id="L71"></a>&#34;`&#34;,
        <a id="L72"></a>literal,
    <a id="L73"></a>},

    <a id="L75"></a><span class="comment">// Operators and delimitors</span>
    <a id="L76"></a>elt{token.ADD, &#34;+&#34;, operator},
    <a id="L77"></a>elt{token.SUB, &#34;-&#34;, operator},
    <a id="L78"></a>elt{token.MUL, &#34;*&#34;, operator},
    <a id="L79"></a>elt{token.QUO, &#34;/&#34;, operator},
    <a id="L80"></a>elt{token.REM, &#34;%&#34;, operator},

    <a id="L82"></a>elt{token.AND, &#34;&amp;&#34;, operator},
    <a id="L83"></a>elt{token.OR, &#34;|&#34;, operator},
    <a id="L84"></a>elt{token.XOR, &#34;^&#34;, operator},
    <a id="L85"></a>elt{token.SHL, &#34;&lt;&lt;&#34;, operator},
    <a id="L86"></a>elt{token.SHR, &#34;&gt;&gt;&#34;, operator},
    <a id="L87"></a>elt{token.AND_NOT, &#34;&amp;^&#34;, operator},

    <a id="L89"></a>elt{token.ADD_ASSIGN, &#34;+=&#34;, operator},
    <a id="L90"></a>elt{token.SUB_ASSIGN, &#34;-=&#34;, operator},
    <a id="L91"></a>elt{token.MUL_ASSIGN, &#34;*=&#34;, operator},
    <a id="L92"></a>elt{token.QUO_ASSIGN, &#34;/=&#34;, operator},
    <a id="L93"></a>elt{token.REM_ASSIGN, &#34;%=&#34;, operator},

    <a id="L95"></a>elt{token.AND_ASSIGN, &#34;&amp;=&#34;, operator},
    <a id="L96"></a>elt{token.OR_ASSIGN, &#34;|=&#34;, operator},
    <a id="L97"></a>elt{token.XOR_ASSIGN, &#34;^=&#34;, operator},
    <a id="L98"></a>elt{token.SHL_ASSIGN, &#34;&lt;&lt;=&#34;, operator},
    <a id="L99"></a>elt{token.SHR_ASSIGN, &#34;&gt;&gt;=&#34;, operator},
    <a id="L100"></a>elt{token.AND_NOT_ASSIGN, &#34;&amp;^=&#34;, operator},

    <a id="L102"></a>elt{token.LAND, &#34;&amp;&amp;&#34;, operator},
    <a id="L103"></a>elt{token.LOR, &#34;||&#34;, operator},
    <a id="L104"></a>elt{token.ARROW, &#34;&lt;-&#34;, operator},
    <a id="L105"></a>elt{token.INC, &#34;++&#34;, operator},
    <a id="L106"></a>elt{token.DEC, &#34;--&#34;, operator},

    <a id="L108"></a>elt{token.EQL, &#34;==&#34;, operator},
    <a id="L109"></a>elt{token.LSS, &#34;&lt;&#34;, operator},
    <a id="L110"></a>elt{token.GTR, &#34;&gt;&#34;, operator},
    <a id="L111"></a>elt{token.ASSIGN, &#34;=&#34;, operator},
    <a id="L112"></a>elt{token.NOT, &#34;!&#34;, operator},

    <a id="L114"></a>elt{token.NEQ, &#34;!=&#34;, operator},
    <a id="L115"></a>elt{token.LEQ, &#34;&lt;=&#34;, operator},
    <a id="L116"></a>elt{token.GEQ, &#34;&gt;=&#34;, operator},
    <a id="L117"></a>elt{token.DEFINE, &#34;:=&#34;, operator},
    <a id="L118"></a>elt{token.ELLIPSIS, &#34;...&#34;, operator},

    <a id="L120"></a>elt{token.LPAREN, &#34;(&#34;, operator},
    <a id="L121"></a>elt{token.LBRACK, &#34;[&#34;, operator},
    <a id="L122"></a>elt{token.LBRACE, &#34;{&#34;, operator},
    <a id="L123"></a>elt{token.COMMA, &#34;,&#34;, operator},
    <a id="L124"></a>elt{token.PERIOD, &#34;.&#34;, operator},

    <a id="L126"></a>elt{token.RPAREN, &#34;)&#34;, operator},
    <a id="L127"></a>elt{token.RBRACK, &#34;]&#34;, operator},
    <a id="L128"></a>elt{token.RBRACE, &#34;}&#34;, operator},
    <a id="L129"></a>elt{token.SEMICOLON, &#34;;&#34;, operator},
    <a id="L130"></a>elt{token.COLON, &#34;:&#34;, operator},

    <a id="L132"></a><span class="comment">// Keywords</span>
    <a id="L133"></a>elt{token.BREAK, &#34;break&#34;, keyword},
    <a id="L134"></a>elt{token.CASE, &#34;case&#34;, keyword},
    <a id="L135"></a>elt{token.CHAN, &#34;chan&#34;, keyword},
    <a id="L136"></a>elt{token.CONST, &#34;const&#34;, keyword},
    <a id="L137"></a>elt{token.CONTINUE, &#34;continue&#34;, keyword},

    <a id="L139"></a>elt{token.DEFAULT, &#34;default&#34;, keyword},
    <a id="L140"></a>elt{token.DEFER, &#34;defer&#34;, keyword},
    <a id="L141"></a>elt{token.ELSE, &#34;else&#34;, keyword},
    <a id="L142"></a>elt{token.FALLTHROUGH, &#34;fallthrough&#34;, keyword},
    <a id="L143"></a>elt{token.FOR, &#34;for&#34;, keyword},

    <a id="L145"></a>elt{token.FUNC, &#34;func&#34;, keyword},
    <a id="L146"></a>elt{token.GO, &#34;go&#34;, keyword},
    <a id="L147"></a>elt{token.GOTO, &#34;goto&#34;, keyword},
    <a id="L148"></a>elt{token.IF, &#34;if&#34;, keyword},
    <a id="L149"></a>elt{token.IMPORT, &#34;import&#34;, keyword},

    <a id="L151"></a>elt{token.INTERFACE, &#34;interface&#34;, keyword},
    <a id="L152"></a>elt{token.MAP, &#34;map&#34;, keyword},
    <a id="L153"></a>elt{token.PACKAGE, &#34;package&#34;, keyword},
    <a id="L154"></a>elt{token.RANGE, &#34;range&#34;, keyword},
    <a id="L155"></a>elt{token.RETURN, &#34;return&#34;, keyword},

    <a id="L157"></a>elt{token.SELECT, &#34;select&#34;, keyword},
    <a id="L158"></a>elt{token.STRUCT, &#34;struct&#34;, keyword},
    <a id="L159"></a>elt{token.SWITCH, &#34;switch&#34;, keyword},
    <a id="L160"></a>elt{token.TYPE, &#34;type&#34;, keyword},
    <a id="L161"></a>elt{token.VAR, &#34;var&#34;, keyword},
<a id="L162"></a>}


<a id="L165"></a>const whitespace = &#34;  \t  \n\n\n&#34; <span class="comment">// to separate tokens</span>

<a id="L167"></a>type TestErrorHandler struct {
    <a id="L168"></a>t *testing.T;
<a id="L169"></a>}

<a id="L171"></a>func (h *TestErrorHandler) Error(pos token.Position, msg string) {
    <a id="L172"></a>h.t.Errorf(&#34;Error() called (msg = %s)&#34;, msg)
<a id="L173"></a>}


<a id="L176"></a>func NewlineCount(s string) int {
    <a id="L177"></a>n := 0;
    <a id="L178"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L179"></a>if s[i] == &#39;\n&#39; {
            <a id="L180"></a>n++
        <a id="L181"></a>}
    <a id="L182"></a>}
    <a id="L183"></a>return n;
<a id="L184"></a>}


<a id="L187"></a>func checkPos(t *testing.T, lit string, pos, expected token.Position) {
    <a id="L188"></a>if pos.Filename != expected.Filename {
        <a id="L189"></a>t.Errorf(&#34;bad filename for %s: got %s, expected %s&#34;, lit, pos.Filename, expected.Filename)
    <a id="L190"></a>}
    <a id="L191"></a>if pos.Offset != expected.Offset {
        <a id="L192"></a>t.Errorf(&#34;bad position for %s: got %d, expected %d&#34;, lit, pos.Offset, expected.Offset)
    <a id="L193"></a>}
    <a id="L194"></a>if pos.Line != expected.Line {
        <a id="L195"></a>t.Errorf(&#34;bad line for %s: got %d, expected %d&#34;, lit, pos.Line, expected.Line)
    <a id="L196"></a>}
    <a id="L197"></a>if pos.Column != expected.Column {
        <a id="L198"></a>t.Errorf(&#34;bad column for %s: got %d, expected %d&#34;, lit, pos.Column, expected.Column)
    <a id="L199"></a>}
<a id="L200"></a>}


<a id="L203"></a><span class="comment">// Verify that calling Scan() provides the correct results.</span>
<a id="L204"></a>func TestScan(t *testing.T) {
    <a id="L205"></a><span class="comment">// make source</span>
    <a id="L206"></a>var src string;
    <a id="L207"></a>for _, e := range tokens {
        <a id="L208"></a>src += e.lit + whitespace
    <a id="L209"></a>}
    <a id="L210"></a>whitespace_linecount := NewlineCount(whitespace);

    <a id="L212"></a><span class="comment">// verify scan</span>
    <a id="L213"></a>index := 0;
    <a id="L214"></a>epos := token.Position{&#34;&#34;, 0, 1, 1};
    <a id="L215"></a>nerrors := Tokenize(&#34;&#34;, strings.Bytes(src), &amp;TestErrorHandler{t}, ScanComments,
        <a id="L216"></a>func(pos token.Position, tok token.Token, litb []byte) bool {
            <a id="L217"></a>e := elt{token.EOF, &#34;&#34;, special};
            <a id="L218"></a>if index &lt; len(tokens) {
                <a id="L219"></a>e = tokens[index]
            <a id="L220"></a>}
            <a id="L221"></a>lit := string(litb);
            <a id="L222"></a>if tok == token.EOF {
                <a id="L223"></a>lit = &#34;&lt;EOF&gt;&#34;;
                <a id="L224"></a>epos.Column = 0;
            <a id="L225"></a>}
            <a id="L226"></a>checkPos(t, lit, pos, epos);
            <a id="L227"></a>if tok != e.tok {
                <a id="L228"></a>t.Errorf(&#34;bad token for %s: got %s, expected %s&#34;, lit, tok.String(), e.tok.String())
            <a id="L229"></a>}
            <a id="L230"></a>if e.tok.IsLiteral() &amp;&amp; lit != e.lit {
                <a id="L231"></a>t.Errorf(&#34;bad literal for %s: got %s, expected %s&#34;, lit, lit, e.lit)
            <a id="L232"></a>}
            <a id="L233"></a>if tokenclass(tok) != e.class {
                <a id="L234"></a>t.Errorf(&#34;bad class for %s: got %d, expected %d&#34;, lit, tokenclass(tok), e.class)
            <a id="L235"></a>}
            <a id="L236"></a>epos.Offset += len(lit) + len(whitespace);
            <a id="L237"></a>epos.Line += NewlineCount(lit) + whitespace_linecount;
            <a id="L238"></a>if tok == token.COMMENT &amp;&amp; litb[1] == &#39;/&#39; {
                <a id="L239"></a><span class="comment">// correct for unaccounted &#39;/n&#39; in //-style comment</span>
                <a id="L240"></a>epos.Offset++;
                <a id="L241"></a>epos.Line++;
            <a id="L242"></a>}
            <a id="L243"></a>index++;
            <a id="L244"></a>return tok != token.EOF;
        <a id="L245"></a>});
    <a id="L246"></a>if nerrors != 0 {
        <a id="L247"></a>t.Errorf(&#34;found %d errors&#34;, nerrors)
    <a id="L248"></a>}
<a id="L249"></a>}


<a id="L252"></a>type seg struct {
    <a id="L253"></a>srcline  string; <span class="comment">// a line of source text</span>
    <a id="L254"></a>filename string; <span class="comment">// filename for current token</span>
    <a id="L255"></a>line     int;    <span class="comment">// line number for current token</span>
<a id="L256"></a>}


<a id="L259"></a>var segments = []seg{
    <a id="L260"></a><span class="comment">// exactly one token per line since the test consumes one token per segment</span>
    <a id="L261"></a>seg{&#34;  line1&#34;, &#34;TestLineComments&#34;, 1},
    <a id="L262"></a>seg{&#34;\nline2&#34;, &#34;TestLineComments&#34;, 2},
    <a id="L263"></a>seg{&#34;\nline3  //line File1.go:100&#34;, &#34;TestLineComments&#34;, 3}, <span class="comment">// bad line comment, ignored</span>
    <a id="L264"></a>seg{&#34;\nline4&#34;, &#34;TestLineComments&#34;, 4},
    <a id="L265"></a>seg{&#34;\n//line File1.go:100\n  line100&#34;, &#34;File1.go&#34;, 100},
    <a id="L266"></a>seg{&#34;\n//line File2.go:200\n  line200&#34;, &#34;File2.go&#34;, 200},
    <a id="L267"></a>seg{&#34;\n//line :1\n  line1&#34;, &#34;&#34;, 1},
    <a id="L268"></a>seg{&#34;\n//line foo:42\n  line42&#34;, &#34;foo&#34;, 42},
    <a id="L269"></a>seg{&#34;\n //line foo:42\n  line44&#34;, &#34;foo&#34;, 44}, <span class="comment">// bad line comment, ignored</span>
    <a id="L270"></a>seg{&#34;\n//line foo 42\n  line46&#34;, &#34;foo&#34;, 46}, <span class="comment">// bad line comment, ignored</span>
    <a id="L271"></a>seg{&#34;\n//line foo:42 extra text\n  line48&#34;, &#34;foo&#34;, 48}, <span class="comment">// bad line comment, ignored</span>
    <a id="L272"></a>seg{&#34;\n//line foo:42\n  line42&#34;, &#34;foo&#34;, 42},
    <a id="L273"></a>seg{&#34;\n//line foo:42\n  line42&#34;, &#34;foo&#34;, 42},
    <a id="L274"></a>seg{&#34;\n//line File1.go:100\n  line100&#34;, &#34;File1.go&#34;, 100},
<a id="L275"></a>}


<a id="L278"></a><span class="comment">// Verify that comments of the form &#34;//line filename:line&#34; are interpreted correctly.</span>
<a id="L279"></a>func TestLineComments(t *testing.T) {
    <a id="L280"></a><span class="comment">// make source</span>
    <a id="L281"></a>var src string;
    <a id="L282"></a>for _, e := range segments {
        <a id="L283"></a>src += e.srcline
    <a id="L284"></a>}

    <a id="L286"></a><span class="comment">// verify scan</span>
    <a id="L287"></a>var S Scanner;
    <a id="L288"></a>S.Init(&#34;TestLineComments&#34;, strings.Bytes(src), nil, 0);
    <a id="L289"></a>for _, s := range segments {
        <a id="L290"></a>pos, _, lit := S.Scan();
        <a id="L291"></a>checkPos(t, string(lit), pos, token.Position{s.filename, pos.Offset, s.line, pos.Column});
    <a id="L292"></a>}

    <a id="L294"></a>if S.ErrorCount != 0 {
        <a id="L295"></a>t.Errorf(&#34;found %d errors&#34;, S.ErrorCount)
    <a id="L296"></a>}
<a id="L297"></a>}


<a id="L300"></a><span class="comment">// Verify that initializing the same scanner more then once works correctly.</span>
<a id="L301"></a>func TestInit(t *testing.T) {
    <a id="L302"></a>var s Scanner;

    <a id="L304"></a><span class="comment">// 1st init</span>
    <a id="L305"></a>s.Init(&#34;&#34;, strings.Bytes(&#34;if true { }&#34;), nil, 0);
    <a id="L306"></a>s.Scan();              <span class="comment">// if</span>
    <a id="L307"></a>s.Scan();              <span class="comment">// true</span>
    <a id="L308"></a>_, tok, _ := s.Scan(); <span class="comment">// {</span>
    <a id="L309"></a>if tok != token.LBRACE {
        <a id="L310"></a>t.Errorf(&#34;bad token: got %s, expected %s&#34;, tok.String(), token.LBRACE)
    <a id="L311"></a>}

    <a id="L313"></a><span class="comment">// 2nd init</span>
    <a id="L314"></a>s.Init(&#34;&#34;, strings.Bytes(&#34;go true { ]&#34;), nil, 0);
    <a id="L315"></a>_, tok, _ = s.Scan(); <span class="comment">// go</span>
    <a id="L316"></a>if tok != token.GO {
        <a id="L317"></a>t.Errorf(&#34;bad token: got %s, expected %s&#34;, tok.String(), token.GO)
    <a id="L318"></a>}

    <a id="L320"></a>if s.ErrorCount != 0 {
        <a id="L321"></a>t.Errorf(&#34;found %d errors&#34;, s.ErrorCount)
    <a id="L322"></a>}
<a id="L323"></a>}


<a id="L326"></a>func TestIllegalChars(t *testing.T) {
    <a id="L327"></a>var s Scanner;

    <a id="L329"></a>const src = &#34;*?*$*@*&#34;;
    <a id="L330"></a>s.Init(&#34;&#34;, strings.Bytes(src), &amp;TestErrorHandler{t}, AllowIllegalChars);
    <a id="L331"></a>for offs, ch := range src {
        <a id="L332"></a>pos, tok, lit := s.Scan();
        <a id="L333"></a>if pos.Offset != offs {
            <a id="L334"></a>t.Errorf(&#34;bad position for %s: got %d, expected %d&#34;, string(lit), pos.Offset, offs)
        <a id="L335"></a>}
        <a id="L336"></a>if tok == token.ILLEGAL &amp;&amp; string(lit) != string(ch) {
            <a id="L337"></a>t.Errorf(&#34;bad token: got %s, expected %s&#34;, string(lit), string(ch))
        <a id="L338"></a>}
    <a id="L339"></a>}

    <a id="L341"></a>if s.ErrorCount != 0 {
        <a id="L342"></a>t.Errorf(&#34;found %d errors&#34;, s.ErrorCount)
    <a id="L343"></a>}
<a id="L344"></a>}


<a id="L347"></a>func TestStdErrorHander(t *testing.T) {
    <a id="L348"></a>const src = &#34;@\n&#34; <span class="comment">// illegal character, cause an error</span>
        <a id="L349"></a>&#34;@ @\n&#34; <span class="comment">// two errors on the same line</span>
        <a id="L350"></a>&#34;//line File2:20\n&#34;
        <a id="L351"></a>&#34;@\n&#34; <span class="comment">// different file, but same line</span>
        <a id="L352"></a>&#34;//line File2:1\n&#34;
        <a id="L353"></a>&#34;@ @\n&#34; <span class="comment">// same file, decreasing line number</span>
        <a id="L354"></a>&#34;//line File1:1\n&#34;
        <a id="L355"></a>&#34;@ @ @&#34;; <span class="comment">// original file, line 1 again</span>


    <a id="L358"></a>v := NewErrorVector();
    <a id="L359"></a>nerrors := Tokenize(&#34;File1&#34;, strings.Bytes(src), v, 0,
        <a id="L360"></a>func(pos token.Position, tok token.Token, litb []byte) bool {
            <a id="L361"></a>return tok != token.EOF
        <a id="L362"></a>});

    <a id="L364"></a>list := v.GetErrorList(Raw);
    <a id="L365"></a>if len(list) != 9 {
        <a id="L366"></a>t.Errorf(&#34;found %d raw errors, expected 9&#34;, len(list));
        <a id="L367"></a>PrintError(os.Stderr, list);
    <a id="L368"></a>}

    <a id="L370"></a>list = v.GetErrorList(Sorted);
    <a id="L371"></a>if len(list) != 9 {
        <a id="L372"></a>t.Errorf(&#34;found %d sorted errors, expected 9&#34;, len(list));
        <a id="L373"></a>PrintError(os.Stderr, list);
    <a id="L374"></a>}

    <a id="L376"></a>list = v.GetErrorList(NoMultiples);
    <a id="L377"></a>if len(list) != 4 {
        <a id="L378"></a>t.Errorf(&#34;found %d one-per-line errors, expected 4&#34;, len(list));
        <a id="L379"></a>PrintError(os.Stderr, list);
    <a id="L380"></a>}

    <a id="L382"></a>if v.ErrorCount() != nerrors {
        <a id="L383"></a>t.Errorf(&#34;found %d errors, expected %d&#34;, v.ErrorCount(), nerrors)
    <a id="L384"></a>}
<a id="L385"></a>}
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
