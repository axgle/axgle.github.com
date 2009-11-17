<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/token/token.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/go/token/token.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package defines constants representing the lexical</span>
<a id="L6"></a><span class="comment">// tokens of the Go programming language and basic operations</span>
<a id="L7"></a><span class="comment">// on tokens (printing, predicates).</span>
<a id="L8"></a><span class="comment">//</span>
<a id="L9"></a>package token

<a id="L11"></a>import (
    <a id="L12"></a>&#34;fmt&#34;;
    <a id="L13"></a>&#34;strconv&#34;;
<a id="L14"></a>)


<a id="L17"></a><span class="comment">// Token is the set of lexical tokens of the Go programming language.</span>
<a id="L18"></a>type Token int

<a id="L20"></a><span class="comment">// The list of tokens.</span>
<a id="L21"></a>const (
    <a id="L22"></a><span class="comment">// Special tokens</span>
    <a id="L23"></a>ILLEGAL Token = iota;
    <a id="L24"></a>EOF;
    <a id="L25"></a>COMMENT;

    <a id="L27"></a>literal_beg;
    <a id="L28"></a><span class="comment">// Identifiers and basic type literals</span>
    <a id="L29"></a><span class="comment">// (these tokens stand for classes of literals)</span>
    <a id="L30"></a>IDENT;  <span class="comment">// main</span>
    <a id="L31"></a>INT;    <span class="comment">// 12345</span>
    <a id="L32"></a>FLOAT;  <span class="comment">// 123.45</span>
    <a id="L33"></a>CHAR;   <span class="comment">// &#39;a&#39;</span>
    <a id="L34"></a>STRING; <span class="comment">// &#34;abc&#34;</span>
    <a id="L35"></a>literal_end;

    <a id="L37"></a>operator_beg;
    <a id="L38"></a><span class="comment">// Operators and delimiters</span>
    <a id="L39"></a>ADD; <span class="comment">// +</span>
    <a id="L40"></a>SUB; <span class="comment">// -</span>
    <a id="L41"></a>MUL; <span class="comment">// *</span>
    <a id="L42"></a>QUO; <span class="comment">// /</span>
    <a id="L43"></a>REM; <span class="comment">// %</span>

    <a id="L45"></a>AND;     <span class="comment">// &amp;</span>
    <a id="L46"></a>OR;      <span class="comment">// |</span>
    <a id="L47"></a>XOR;     <span class="comment">// ^</span>
    <a id="L48"></a>SHL;     <span class="comment">// &lt;&lt;</span>
    <a id="L49"></a>SHR;     <span class="comment">// &gt;&gt;</span>
    <a id="L50"></a>AND_NOT; <span class="comment">// &amp;^</span>

    <a id="L52"></a>ADD_ASSIGN; <span class="comment">// +=</span>
    <a id="L53"></a>SUB_ASSIGN; <span class="comment">// -=</span>
    <a id="L54"></a>MUL_ASSIGN; <span class="comment">// *=</span>
    <a id="L55"></a>QUO_ASSIGN; <span class="comment">// /=</span>
    <a id="L56"></a>REM_ASSIGN; <span class="comment">// %=</span>

    <a id="L58"></a>AND_ASSIGN;     <span class="comment">// &amp;=</span>
    <a id="L59"></a>OR_ASSIGN;      <span class="comment">// |=</span>
    <a id="L60"></a>XOR_ASSIGN;     <span class="comment">// ^=</span>
    <a id="L61"></a>SHL_ASSIGN;     <span class="comment">// &lt;&lt;=</span>
    <a id="L62"></a>SHR_ASSIGN;     <span class="comment">// &gt;&gt;=</span>
    <a id="L63"></a>AND_NOT_ASSIGN; <span class="comment">// &amp;^=</span>

    <a id="L65"></a>LAND;  <span class="comment">// &amp;&amp;</span>
    <a id="L66"></a>LOR;   <span class="comment">// ||</span>
    <a id="L67"></a>ARROW; <span class="comment">// &lt;-</span>
    <a id="L68"></a>INC;   <span class="comment">// ++</span>
    <a id="L69"></a>DEC;   <span class="comment">// --</span>

    <a id="L71"></a>EQL;    <span class="comment">// ==</span>
    <a id="L72"></a>LSS;    <span class="comment">// &lt;</span>
    <a id="L73"></a>GTR;    <span class="comment">// &gt;</span>
    <a id="L74"></a>ASSIGN; <span class="comment">// =</span>
    <a id="L75"></a>NOT;    <span class="comment">// !</span>

    <a id="L77"></a>NEQ;      <span class="comment">// !=</span>
    <a id="L78"></a>LEQ;      <span class="comment">// &lt;=</span>
    <a id="L79"></a>GEQ;      <span class="comment">// &gt;=</span>
    <a id="L80"></a>DEFINE;   <span class="comment">// :=</span>
    <a id="L81"></a>ELLIPSIS; <span class="comment">// ...</span>

    <a id="L83"></a>LPAREN; <span class="comment">// (</span>
    <a id="L84"></a>LBRACK; <span class="comment">// [</span>
    <a id="L85"></a>LBRACE; <span class="comment">// {</span>
    <a id="L86"></a>COMMA;  <span class="comment">// ,</span>
    <a id="L87"></a>PERIOD; <span class="comment">// .</span>

    <a id="L89"></a>RPAREN;    <span class="comment">// )</span>
    <a id="L90"></a>RBRACK;    <span class="comment">// ]</span>
    <a id="L91"></a>RBRACE;    <span class="comment">// }</span>
    <a id="L92"></a>SEMICOLON; <span class="comment">// ;</span>
    <a id="L93"></a>COLON;     <span class="comment">// :</span>
    <a id="L94"></a>operator_end;

    <a id="L96"></a>keyword_beg;
    <a id="L97"></a><span class="comment">// Keywords</span>
    <a id="L98"></a>BREAK;
    <a id="L99"></a>CASE;
    <a id="L100"></a>CHAN;
    <a id="L101"></a>CONST;
    <a id="L102"></a>CONTINUE;

    <a id="L104"></a>DEFAULT;
    <a id="L105"></a>DEFER;
    <a id="L106"></a>ELSE;
    <a id="L107"></a>FALLTHROUGH;
    <a id="L108"></a>FOR;

    <a id="L110"></a>FUNC;
    <a id="L111"></a>GO;
    <a id="L112"></a>GOTO;
    <a id="L113"></a>IF;
    <a id="L114"></a>IMPORT;

    <a id="L116"></a>INTERFACE;
    <a id="L117"></a>MAP;
    <a id="L118"></a>PACKAGE;
    <a id="L119"></a>RANGE;
    <a id="L120"></a>RETURN;

    <a id="L122"></a>SELECT;
    <a id="L123"></a>STRUCT;
    <a id="L124"></a>SWITCH;
    <a id="L125"></a>TYPE;
    <a id="L126"></a>VAR;
    <a id="L127"></a>keyword_end;
<a id="L128"></a>)


<a id="L131"></a><span class="comment">// At the moment we have no array literal syntax that lets us describe</span>
<a id="L132"></a><span class="comment">// the index for each element - use a map for now to make sure they are</span>
<a id="L133"></a><span class="comment">// in sync.</span>
<a id="L134"></a>var tokens = map[Token]string{
    <a id="L135"></a>ILLEGAL: &#34;ILLEGAL&#34;,

    <a id="L137"></a>EOF: &#34;EOF&#34;,
    <a id="L138"></a>COMMENT: &#34;COMMENT&#34;,

    <a id="L140"></a>IDENT: &#34;IDENT&#34;,
    <a id="L141"></a>INT: &#34;INT&#34;,
    <a id="L142"></a>FLOAT: &#34;FLOAT&#34;,
    <a id="L143"></a>CHAR: &#34;CHAR&#34;,
    <a id="L144"></a>STRING: &#34;STRING&#34;,

    <a id="L146"></a>ADD: &#34;+&#34;,
    <a id="L147"></a>SUB: &#34;-&#34;,
    <a id="L148"></a>MUL: &#34;*&#34;,
    <a id="L149"></a>QUO: &#34;/&#34;,
    <a id="L150"></a>REM: &#34;%&#34;,

    <a id="L152"></a>AND: &#34;&amp;&#34;,
    <a id="L153"></a>OR: &#34;|&#34;,
    <a id="L154"></a>XOR: &#34;^&#34;,
    <a id="L155"></a>SHL: &#34;&lt;&lt;&#34;,
    <a id="L156"></a>SHR: &#34;&gt;&gt;&#34;,
    <a id="L157"></a>AND_NOT: &#34;&amp;^&#34;,

    <a id="L159"></a>ADD_ASSIGN: &#34;+=&#34;,
    <a id="L160"></a>SUB_ASSIGN: &#34;-=&#34;,
    <a id="L161"></a>MUL_ASSIGN: &#34;*=&#34;,
    <a id="L162"></a>QUO_ASSIGN: &#34;/=&#34;,
    <a id="L163"></a>REM_ASSIGN: &#34;%=&#34;,

    <a id="L165"></a>AND_ASSIGN: &#34;&amp;=&#34;,
    <a id="L166"></a>OR_ASSIGN: &#34;|=&#34;,
    <a id="L167"></a>XOR_ASSIGN: &#34;^=&#34;,
    <a id="L168"></a>SHL_ASSIGN: &#34;&lt;&lt;=&#34;,
    <a id="L169"></a>SHR_ASSIGN: &#34;&gt;&gt;=&#34;,
    <a id="L170"></a>AND_NOT_ASSIGN: &#34;&amp;^=&#34;,

    <a id="L172"></a>LAND: &#34;&amp;&amp;&#34;,
    <a id="L173"></a>LOR: &#34;||&#34;,
    <a id="L174"></a>ARROW: &#34;&lt;-&#34;,
    <a id="L175"></a>INC: &#34;++&#34;,
    <a id="L176"></a>DEC: &#34;--&#34;,

    <a id="L178"></a>EQL: &#34;==&#34;,
    <a id="L179"></a>LSS: &#34;&lt;&#34;,
    <a id="L180"></a>GTR: &#34;&gt;&#34;,
    <a id="L181"></a>ASSIGN: &#34;=&#34;,
    <a id="L182"></a>NOT: &#34;!&#34;,

    <a id="L184"></a>NEQ: &#34;!=&#34;,
    <a id="L185"></a>LEQ: &#34;&lt;=&#34;,
    <a id="L186"></a>GEQ: &#34;&gt;=&#34;,
    <a id="L187"></a>DEFINE: &#34;:=&#34;,
    <a id="L188"></a>ELLIPSIS: &#34;...&#34;,

    <a id="L190"></a>LPAREN: &#34;(&#34;,
    <a id="L191"></a>LBRACK: &#34;[&#34;,
    <a id="L192"></a>LBRACE: &#34;{&#34;,
    <a id="L193"></a>COMMA: &#34;,&#34;,
    <a id="L194"></a>PERIOD: &#34;.&#34;,

    <a id="L196"></a>RPAREN: &#34;)&#34;,
    <a id="L197"></a>RBRACK: &#34;]&#34;,
    <a id="L198"></a>RBRACE: &#34;}&#34;,
    <a id="L199"></a>SEMICOLON: &#34;;&#34;,
    <a id="L200"></a>COLON: &#34;:&#34;,

    <a id="L202"></a>BREAK: &#34;break&#34;,
    <a id="L203"></a>CASE: &#34;case&#34;,
    <a id="L204"></a>CHAN: &#34;chan&#34;,
    <a id="L205"></a>CONST: &#34;const&#34;,
    <a id="L206"></a>CONTINUE: &#34;continue&#34;,

    <a id="L208"></a>DEFAULT: &#34;default&#34;,
    <a id="L209"></a>DEFER: &#34;defer&#34;,
    <a id="L210"></a>ELSE: &#34;else&#34;,
    <a id="L211"></a>FALLTHROUGH: &#34;fallthrough&#34;,
    <a id="L212"></a>FOR: &#34;for&#34;,

    <a id="L214"></a>FUNC: &#34;func&#34;,
    <a id="L215"></a>GO: &#34;go&#34;,
    <a id="L216"></a>GOTO: &#34;goto&#34;,
    <a id="L217"></a>IF: &#34;if&#34;,
    <a id="L218"></a>IMPORT: &#34;import&#34;,

    <a id="L220"></a>INTERFACE: &#34;interface&#34;,
    <a id="L221"></a>MAP: &#34;map&#34;,
    <a id="L222"></a>PACKAGE: &#34;package&#34;,
    <a id="L223"></a>RANGE: &#34;range&#34;,
    <a id="L224"></a>RETURN: &#34;return&#34;,

    <a id="L226"></a>SELECT: &#34;select&#34;,
    <a id="L227"></a>STRUCT: &#34;struct&#34;,
    <a id="L228"></a>SWITCH: &#34;switch&#34;,
    <a id="L229"></a>TYPE: &#34;type&#34;,
    <a id="L230"></a>VAR: &#34;var&#34;,
<a id="L231"></a>}


<a id="L234"></a><span class="comment">// String returns the string corresponding to the token tok.</span>
<a id="L235"></a><span class="comment">// For operators, delimiters, and keywords the string is the actual</span>
<a id="L236"></a><span class="comment">// token character sequence (e.g., for the token ADD, the string is</span>
<a id="L237"></a><span class="comment">// &#34;+&#34;). For all other tokens the string corresponds to the token</span>
<a id="L238"></a><span class="comment">// constant name (e.g. for the token IDENT, the string is &#34;IDENT&#34;).</span>
<a id="L239"></a><span class="comment">//</span>
<a id="L240"></a>func (tok Token) String() string {
    <a id="L241"></a>if str, exists := tokens[tok]; exists {
        <a id="L242"></a>return str
    <a id="L243"></a>}
    <a id="L244"></a>return &#34;token(&#34; + strconv.Itoa(int(tok)) + &#34;)&#34;;
<a id="L245"></a>}


<a id="L248"></a><span class="comment">// A set of constants for precedence-based expression parsing.</span>
<a id="L249"></a><span class="comment">// Non-operators have lowest precedence, followed by operators</span>
<a id="L250"></a><span class="comment">// starting with precedence 1 up to unary operators. The highest</span>
<a id="L251"></a><span class="comment">// precedence corresponds serves as &#34;catch-all&#34; precedence for</span>
<a id="L252"></a><span class="comment">// selector, indexing, and other operator and delimiter tokens.</span>
<a id="L253"></a><span class="comment">//</span>
<a id="L254"></a>const (
    <a id="L255"></a>LowestPrec  = 0; <span class="comment">// non-operators</span>
    <a id="L256"></a>UnaryPrec   = 7;
    <a id="L257"></a>HighestPrec = 8;
<a id="L258"></a>)


<a id="L261"></a><span class="comment">// Precedence returns the operator precedence of the binary</span>
<a id="L262"></a><span class="comment">// operator op. If op is not a binary operator, the result</span>
<a id="L263"></a><span class="comment">// is LowestPrecedence.</span>
<a id="L264"></a><span class="comment">//</span>
<a id="L265"></a>func (op Token) Precedence() int {
    <a id="L266"></a>switch op {
    <a id="L267"></a>case LOR:
        <a id="L268"></a>return 1
    <a id="L269"></a>case LAND:
        <a id="L270"></a>return 2
    <a id="L271"></a>case ARROW:
        <a id="L272"></a>return 3
    <a id="L273"></a>case EQL, NEQ, LSS, LEQ, GTR, GEQ:
        <a id="L274"></a>return 4
    <a id="L275"></a>case ADD, SUB, OR, XOR:
        <a id="L276"></a>return 5
    <a id="L277"></a>case MUL, QUO, REM, SHL, SHR, AND, AND_NOT:
        <a id="L278"></a>return 6
    <a id="L279"></a>}
    <a id="L280"></a>return LowestPrec;
<a id="L281"></a>}


<a id="L284"></a>var keywords map[string]Token

<a id="L286"></a>func init() {
    <a id="L287"></a>keywords = make(map[string]Token);
    <a id="L288"></a>for i := keyword_beg + 1; i &lt; keyword_end; i++ {
        <a id="L289"></a>keywords[tokens[i]] = i
    <a id="L290"></a>}
<a id="L291"></a>}


<a id="L294"></a><span class="comment">// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).</span>
<a id="L295"></a><span class="comment">//</span>
<a id="L296"></a>func Lookup(ident []byte) Token {
    <a id="L297"></a><span class="comment">// TODO Maps with []byte key are illegal because []byte does not</span>
    <a id="L298"></a><span class="comment">//      support == . Should find a more efficient solution eventually.</span>
    <a id="L299"></a>if tok, is_keyword := keywords[string(ident)]; is_keyword {
        <a id="L300"></a>return tok
    <a id="L301"></a>}
    <a id="L302"></a>return IDENT;
<a id="L303"></a>}


<a id="L306"></a><span class="comment">// Predicates</span>

<a id="L308"></a><span class="comment">// IsLiteral returns true for tokens corresponding to identifiers</span>
<a id="L309"></a><span class="comment">// and basic type literals; returns false otherwise.</span>
<a id="L310"></a><span class="comment">//</span>
<a id="L311"></a>func (tok Token) IsLiteral() bool { return literal_beg &lt; tok &amp;&amp; tok &lt; literal_end }

<a id="L313"></a><span class="comment">// IsOperator returns true for tokens corresponding to operators and</span>
<a id="L314"></a><span class="comment">// delimiters; returns false otherwise.</span>
<a id="L315"></a><span class="comment">//</span>
<a id="L316"></a>func (tok Token) IsOperator() bool { return operator_beg &lt; tok &amp;&amp; tok &lt; operator_end }

<a id="L318"></a><span class="comment">// IsKeyword returns true for tokens corresponding to keywords;</span>
<a id="L319"></a><span class="comment">// returns false otherwise.</span>
<a id="L320"></a><span class="comment">//</span>
<a id="L321"></a>func (tok Token) IsKeyword() bool { return keyword_beg &lt; tok &amp;&amp; tok &lt; keyword_end }


<a id="L324"></a><span class="comment">// Token source positions are represented by a Position value.</span>
<a id="L325"></a><span class="comment">// A Position is valid if the line number is &gt; 0.</span>
<a id="L326"></a><span class="comment">//</span>
<a id="L327"></a>type Position struct {
    <a id="L328"></a>Filename string; <span class="comment">// filename, if any</span>
    <a id="L329"></a>Offset   int;    <span class="comment">// byte offset, starting at 0</span>
    <a id="L330"></a>Line     int;    <span class="comment">// line number, starting at 1</span>
    <a id="L331"></a>Column   int;    <span class="comment">// column number, starting at 1 (character count)</span>
<a id="L332"></a>}


<a id="L335"></a><span class="comment">// Pos is an accessor method for anonymous Position fields.</span>
<a id="L336"></a><span class="comment">// It returns its receiver.</span>
<a id="L337"></a><span class="comment">//</span>
<a id="L338"></a>func (pos *Position) Pos() Position { return *pos }


<a id="L341"></a><span class="comment">// IsValid returns true if the position is valid.</span>
<a id="L342"></a>func (pos *Position) IsValid() bool { return pos.Line &gt; 0 }


<a id="L345"></a>func (pos Position) String() string {
    <a id="L346"></a>s := pos.Filename;
    <a id="L347"></a>if pos.IsValid() {
        <a id="L348"></a>if s != &#34;&#34; {
            <a id="L349"></a>s += &#34;:&#34;
        <a id="L350"></a>}
        <a id="L351"></a>s += fmt.Sprintf(&#34;%d:%d&#34;, pos.Line, pos.Column);
    <a id="L352"></a>}
    <a id="L353"></a>if s == &#34;&#34; {
        <a id="L354"></a>s = &#34;???&#34;
    <a id="L355"></a>}
    <a id="L356"></a>return s;
<a id="L357"></a>}
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
