<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/go/scanner/scanner.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/go/scanner/scanner.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// A scanner for Go source text. Takes a []byte as source which can</span>
<a id="L6"></a><span class="comment">// then be tokenized through repeated calls to the Scan function.</span>
<a id="L7"></a><span class="comment">// For a sample use of a scanner, see the implementation of Tokenize.</span>
<a id="L8"></a><span class="comment">//</span>
<a id="L9"></a>package scanner

<a id="L11"></a>import (
    <a id="L12"></a>&#34;bytes&#34;;
    <a id="L13"></a>&#34;go/token&#34;;
    <a id="L14"></a>&#34;strconv&#34;;
    <a id="L15"></a>&#34;unicode&#34;;
    <a id="L16"></a>&#34;utf8&#34;;
<a id="L17"></a>)


<a id="L20"></a><span class="comment">// A Scanner holds the scanner&#39;s internal state while processing</span>
<a id="L21"></a><span class="comment">// a given text.  It can be allocated as part of another data</span>
<a id="L22"></a><span class="comment">// structure but must be initialized via Init before use. For</span>
<a id="L23"></a><span class="comment">// a sample use, see the implementation of Tokenize.</span>
<a id="L24"></a><span class="comment">//</span>
<a id="L25"></a>type Scanner struct {
    <a id="L26"></a><span class="comment">// immutable state</span>
    <a id="L27"></a>src  []byte;       <span class="comment">// source</span>
    <a id="L28"></a>err  ErrorHandler; <span class="comment">// error reporting; or nil</span>
    <a id="L29"></a>mode uint;         <span class="comment">// scanning mode</span>

    <a id="L31"></a><span class="comment">// scanning state</span>
    <a id="L32"></a>pos    token.Position; <span class="comment">// previous reading position (position before ch)</span>
    <a id="L33"></a>offset int;            <span class="comment">// current reading offset (position after ch)</span>
    <a id="L34"></a>ch     int;            <span class="comment">// one char look-ahead</span>

    <a id="L36"></a><span class="comment">// public state - ok to modify</span>
    <a id="L37"></a>ErrorCount int; <span class="comment">// number of errors encountered</span>
<a id="L38"></a>}


<a id="L41"></a><span class="comment">// Read the next Unicode char into S.ch.</span>
<a id="L42"></a><span class="comment">// S.ch &lt; 0 means end-of-file.</span>
<a id="L43"></a><span class="comment">//</span>
<a id="L44"></a>func (S *Scanner) next() {
    <a id="L45"></a>if S.offset &lt; len(S.src) {
        <a id="L46"></a>S.pos.Offset = S.offset;
        <a id="L47"></a>S.pos.Column++;
        <a id="L48"></a>r, w := int(S.src[S.offset]), 1;
        <a id="L49"></a>switch {
        <a id="L50"></a>case r == &#39;\n&#39;:
            <a id="L51"></a>S.pos.Line++;
            <a id="L52"></a>S.pos.Column = 0;
        <a id="L53"></a>case r &gt;= 0x80:
            <a id="L54"></a><span class="comment">// not ASCII</span>
            <a id="L55"></a>r, w = utf8.DecodeRune(S.src[S.offset:len(S.src)])
        <a id="L56"></a>}
        <a id="L57"></a>S.offset += w;
        <a id="L58"></a>S.ch = r;
    <a id="L59"></a>} else {
        <a id="L60"></a>S.pos.Offset = len(S.src);
        <a id="L61"></a>S.ch = -1; <span class="comment">// eof</span>
    <a id="L62"></a>}
<a id="L63"></a>}


<a id="L66"></a><span class="comment">// The mode parameter to the Init function is a set of flags (or 0).</span>
<a id="L67"></a><span class="comment">// They control scanner behavior.</span>
<a id="L68"></a><span class="comment">//</span>
<a id="L69"></a>const (
    <a id="L70"></a>ScanComments       = 1 &lt;&lt; iota; <span class="comment">// return comments as COMMENT tokens</span>
    <a id="L71"></a>AllowIllegalChars; <span class="comment">// do not report an error for illegal chars</span>
<a id="L72"></a>)


<a id="L75"></a><span class="comment">// Init prepares the scanner S to tokenize the text src. Calls to Scan</span>
<a id="L76"></a><span class="comment">// will use the error handler err if they encounter a syntax error and</span>
<a id="L77"></a><span class="comment">// err is not nil. Also, for each error encountered, the Scanner field</span>
<a id="L78"></a><span class="comment">// ErrorCount is incremented by one. The filename parameter is used as</span>
<a id="L79"></a><span class="comment">// filename in the token.Position returned by Scan for each token. The</span>
<a id="L80"></a><span class="comment">// mode parameter determines how comments and illegal characters are</span>
<a id="L81"></a><span class="comment">// handled.</span>
<a id="L82"></a><span class="comment">//</span>
<a id="L83"></a>func (S *Scanner) Init(filename string, src []byte, err ErrorHandler, mode uint) {
    <a id="L84"></a><span class="comment">// Explicitly initialize all fields since a scanner may be reused.</span>
    <a id="L85"></a>S.src = src;
    <a id="L86"></a>S.err = err;
    <a id="L87"></a>S.mode = mode;
    <a id="L88"></a>S.pos = token.Position{filename, 0, 1, 0};
    <a id="L89"></a>S.offset = 0;
    <a id="L90"></a>S.ErrorCount = 0;
    <a id="L91"></a>S.next();
<a id="L92"></a>}


<a id="L95"></a>func charString(ch int) string {
    <a id="L96"></a>var s string;
    <a id="L97"></a>switch ch {
    <a id="L98"></a>case -1:
        <a id="L99"></a>return `EOF`
    <a id="L100"></a>case &#39;\a&#39;:
        <a id="L101"></a>s = `\a`
    <a id="L102"></a>case &#39;\b&#39;:
        <a id="L103"></a>s = `\b`
    <a id="L104"></a>case &#39;\f&#39;:
        <a id="L105"></a>s = `\f`
    <a id="L106"></a>case &#39;\n&#39;:
        <a id="L107"></a>s = `\n`
    <a id="L108"></a>case &#39;\r&#39;:
        <a id="L109"></a>s = `\r`
    <a id="L110"></a>case &#39;\t&#39;:
        <a id="L111"></a>s = `\t`
    <a id="L112"></a>case &#39;\v&#39;:
        <a id="L113"></a>s = `\v`
    <a id="L114"></a>case &#39;\\&#39;:
        <a id="L115"></a>s = `\\`
    <a id="L116"></a>case &#39;\&#39;&#39;:
        <a id="L117"></a>s = `\&#39;`
    <a id="L118"></a>default:
        <a id="L119"></a>s = string(ch)
    <a id="L120"></a>}
    <a id="L121"></a>return &#34;&#39;&#34; + s + &#34;&#39; (U+&#34; + strconv.Itob(ch, 16) + &#34;)&#34;;
<a id="L122"></a>}


<a id="L125"></a>func (S *Scanner) error(pos token.Position, msg string) {
    <a id="L126"></a>if S.err != nil {
        <a id="L127"></a>S.err.Error(pos, msg)
    <a id="L128"></a>}
    <a id="L129"></a>S.ErrorCount++;
<a id="L130"></a>}


<a id="L133"></a>func (S *Scanner) expect(ch int) {
    <a id="L134"></a>if S.ch != ch {
        <a id="L135"></a>S.error(S.pos, &#34;expected &#34;+charString(ch)+&#34;, found &#34;+charString(S.ch))
    <a id="L136"></a>}
    <a id="L137"></a>S.next(); <span class="comment">// always make progress</span>
<a id="L138"></a>}


<a id="L141"></a>var prefix = []byte{&#39;l&#39;, &#39;i&#39;, &#39;n&#39;, &#39;e&#39;, &#39; &#39;} <span class="comment">// &#34;line &#34;</span>

<a id="L143"></a>func (S *Scanner) scanComment(pos token.Position) {
    <a id="L144"></a><span class="comment">// first &#39;/&#39; already consumed</span>

    <a id="L146"></a>if S.ch == &#39;/&#39; {
        <a id="L147"></a><span class="comment">//-style comment</span>
        <a id="L148"></a>for S.ch &gt;= 0 {
            <a id="L149"></a>S.next();
            <a id="L150"></a>if S.ch == &#39;\n&#39; {
                <a id="L151"></a><span class="comment">// &#39;\n&#39; is not part of the comment</span>
                <a id="L152"></a><span class="comment">// (the comment ends on the same line where it started)</span>
                <a id="L153"></a>if pos.Column == 1 {
                    <a id="L154"></a>text := S.src[pos.Offset+2 : S.pos.Offset];
                    <a id="L155"></a>if bytes.HasPrefix(text, prefix) {
                        <a id="L156"></a><span class="comment">// comment starts at beginning of line with &#34;//line &#34;;</span>
                        <a id="L157"></a><span class="comment">// get filename and line number, if any</span>
                        <a id="L158"></a>i := bytes.Index(text, []byte{&#39;:&#39;});
                        <a id="L159"></a>if i &gt;= 0 {
                            <a id="L160"></a>if line, err := strconv.Atoi(string(text[i+1 : len(text)])); err == nil &amp;&amp; line &gt; 0 {
                                <a id="L161"></a><span class="comment">// valid //line filename:line comment;</span>
                                <a id="L162"></a><span class="comment">// update scanner position</span>
                                <a id="L163"></a>S.pos.Filename = string(text[len(prefix):i]);
                                <a id="L164"></a>S.pos.Line = line;
                            <a id="L165"></a>}
                        <a id="L166"></a>}
                    <a id="L167"></a>}
                <a id="L168"></a>}
                <a id="L169"></a>return;
            <a id="L170"></a>}
        <a id="L171"></a>}

    <a id="L173"></a>} else {
        <a id="L174"></a><span class="comment">/*-style comment */</span>
        <a id="L175"></a>S.expect(&#39;*&#39;);
        <a id="L176"></a>for S.ch &gt;= 0 {
            <a id="L177"></a>ch := S.ch;
            <a id="L178"></a>S.next();
            <a id="L179"></a>if ch == &#39;*&#39; &amp;&amp; S.ch == &#39;/&#39; {
                <a id="L180"></a>S.next();
                <a id="L181"></a>return;
            <a id="L182"></a>}
        <a id="L183"></a>}
    <a id="L184"></a>}

    <a id="L186"></a>S.error(pos, &#34;comment not terminated&#34;);
<a id="L187"></a>}


<a id="L190"></a>func isLetter(ch int) bool {
    <a id="L191"></a>return &#39;a&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;z&#39; || &#39;A&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;Z&#39; || ch == &#39;_&#39; || ch &gt;= 0x80 &amp;&amp; unicode.IsLetter(ch)
<a id="L192"></a>}


<a id="L195"></a>func isDigit(ch int) bool {
    <a id="L196"></a>return &#39;0&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;9&#39; || ch &gt;= 0x80 &amp;&amp; unicode.IsDigit(ch)
<a id="L197"></a>}


<a id="L200"></a>func (S *Scanner) scanIdentifier() token.Token {
    <a id="L201"></a>pos := S.pos.Offset;
    <a id="L202"></a>for isLetter(S.ch) || isDigit(S.ch) {
        <a id="L203"></a>S.next()
    <a id="L204"></a>}
    <a id="L205"></a>return token.Lookup(S.src[pos:S.pos.Offset]);
<a id="L206"></a>}


<a id="L209"></a>func digitVal(ch int) int {
    <a id="L210"></a>switch {
    <a id="L211"></a>case &#39;0&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;9&#39;:
        <a id="L212"></a>return ch - &#39;0&#39;
    <a id="L213"></a>case &#39;a&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;f&#39;:
        <a id="L214"></a>return ch - &#39;a&#39; + 10
    <a id="L215"></a>case &#39;A&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;F&#39;:
        <a id="L216"></a>return ch - &#39;A&#39; + 10
    <a id="L217"></a>}
    <a id="L218"></a>return 16; <span class="comment">// larger than any legal digit val</span>
<a id="L219"></a>}


<a id="L222"></a>func (S *Scanner) scanMantissa(base int) {
    <a id="L223"></a>for digitVal(S.ch) &lt; base {
        <a id="L224"></a>S.next()
    <a id="L225"></a>}
<a id="L226"></a>}


<a id="L229"></a>func (S *Scanner) scanNumber(seen_decimal_point bool) token.Token {
    <a id="L230"></a>tok := token.INT;

    <a id="L232"></a>if seen_decimal_point {
        <a id="L233"></a>tok = token.FLOAT;
        <a id="L234"></a>S.scanMantissa(10);
        <a id="L235"></a>goto exponent;
    <a id="L236"></a>}

    <a id="L238"></a>if S.ch == &#39;0&#39; {
        <a id="L239"></a><span class="comment">// int or float</span>
        <a id="L240"></a>S.next();
        <a id="L241"></a>if S.ch == &#39;x&#39; || S.ch == &#39;X&#39; {
            <a id="L242"></a><span class="comment">// hexadecimal int</span>
            <a id="L243"></a>S.next();
            <a id="L244"></a>S.scanMantissa(16);
        <a id="L245"></a>} else {
            <a id="L246"></a><span class="comment">// octal int or float</span>
            <a id="L247"></a>S.scanMantissa(8);
            <a id="L248"></a>if digitVal(S.ch) &lt; 10 || S.ch == &#39;.&#39; || S.ch == &#39;e&#39; || S.ch == &#39;E&#39; {
                <a id="L249"></a><span class="comment">// float</span>
                <a id="L250"></a>tok = token.FLOAT;
                <a id="L251"></a>goto mantissa;
            <a id="L252"></a>}
            <a id="L253"></a><span class="comment">// octal int</span>
        <a id="L254"></a>}
        <a id="L255"></a>goto exit;
    <a id="L256"></a>}

<a id="L258"></a>mantissa:
    <a id="L259"></a><span class="comment">// decimal int or float</span>
    <a id="L260"></a>S.scanMantissa(10);

    <a id="L262"></a>if S.ch == &#39;.&#39; {
        <a id="L263"></a><span class="comment">// float</span>
        <a id="L264"></a>tok = token.FLOAT;
        <a id="L265"></a>S.next();
        <a id="L266"></a>S.scanMantissa(10);
    <a id="L267"></a>}

<a id="L269"></a>exponent:
    <a id="L270"></a>if S.ch == &#39;e&#39; || S.ch == &#39;E&#39; {
        <a id="L271"></a><span class="comment">// float</span>
        <a id="L272"></a>tok = token.FLOAT;
        <a id="L273"></a>S.next();
        <a id="L274"></a>if S.ch == &#39;-&#39; || S.ch == &#39;+&#39; {
            <a id="L275"></a>S.next()
        <a id="L276"></a>}
        <a id="L277"></a>S.scanMantissa(10);
    <a id="L278"></a>}

<a id="L280"></a>exit:
    <a id="L281"></a>return tok;
<a id="L282"></a>}


<a id="L285"></a>func (S *Scanner) scanDigits(base, length int) {
    <a id="L286"></a>for length &gt; 0 &amp;&amp; digitVal(S.ch) &lt; base {
        <a id="L287"></a>S.next();
        <a id="L288"></a>length--;
    <a id="L289"></a>}
    <a id="L290"></a>if length &gt; 0 {
        <a id="L291"></a>S.error(S.pos, &#34;illegal char escape&#34;)
    <a id="L292"></a>}
<a id="L293"></a>}


<a id="L296"></a>func (S *Scanner) scanEscape(quote int) {
    <a id="L297"></a>pos := S.pos;
    <a id="L298"></a>ch := S.ch;
    <a id="L299"></a>S.next();
    <a id="L300"></a>switch ch {
    <a id="L301"></a>case &#39;a&#39;, &#39;b&#39;, &#39;f&#39;, &#39;n&#39;, &#39;r&#39;, &#39;t&#39;, &#39;v&#39;, &#39;\\&#39;, quote:
    <a id="L302"></a><span class="comment">// nothing to do</span>
    <a id="L303"></a>case &#39;0&#39;, &#39;1&#39;, &#39;2&#39;, &#39;3&#39;, &#39;4&#39;, &#39;5&#39;, &#39;6&#39;, &#39;7&#39;:
        <a id="L304"></a>S.scanDigits(8, 3-1) <span class="comment">// 1 char read already</span>
    <a id="L305"></a>case &#39;x&#39;:
        <a id="L306"></a>S.scanDigits(16, 2)
    <a id="L307"></a>case &#39;u&#39;:
        <a id="L308"></a>S.scanDigits(16, 4)
    <a id="L309"></a>case &#39;U&#39;:
        <a id="L310"></a>S.scanDigits(16, 8)
    <a id="L311"></a>default:
        <a id="L312"></a>S.error(pos, &#34;illegal char escape&#34;)
    <a id="L313"></a>}
<a id="L314"></a>}


<a id="L317"></a>func (S *Scanner) scanChar(pos token.Position) {
    <a id="L318"></a><span class="comment">// &#39;\&#39;&#39; already consumed</span>

    <a id="L320"></a>n := 0;
    <a id="L321"></a>for S.ch != &#39;\&#39;&#39; {
        <a id="L322"></a>ch := S.ch;
        <a id="L323"></a>n++;
        <a id="L324"></a>S.next();
        <a id="L325"></a>if ch == &#39;\n&#39; || ch &lt; 0 {
            <a id="L326"></a>S.error(pos, &#34;character literal not terminated&#34;);
            <a id="L327"></a>n = 1;
            <a id="L328"></a>break;
        <a id="L329"></a>}
        <a id="L330"></a>if ch == &#39;\\&#39; {
            <a id="L331"></a>S.scanEscape(&#39;\&#39;&#39;)
        <a id="L332"></a>}
    <a id="L333"></a>}

    <a id="L335"></a>S.next();

    <a id="L337"></a>if n != 1 {
        <a id="L338"></a>S.error(pos, &#34;illegal character literal&#34;)
    <a id="L339"></a>}
<a id="L340"></a>}


<a id="L343"></a>func (S *Scanner) scanString(pos token.Position) {
    <a id="L344"></a><span class="comment">// &#39;&#34;&#39; already consumed</span>

    <a id="L346"></a>for S.ch != &#39;&#34;&#39; {
        <a id="L347"></a>ch := S.ch;
        <a id="L348"></a>S.next();
        <a id="L349"></a>if ch == &#39;\n&#39; || ch &lt; 0 {
            <a id="L350"></a>S.error(pos, &#34;string not terminated&#34;);
            <a id="L351"></a>break;
        <a id="L352"></a>}
        <a id="L353"></a>if ch == &#39;\\&#39; {
            <a id="L354"></a>S.scanEscape(&#39;&#34;&#39;)
        <a id="L355"></a>}
    <a id="L356"></a>}

    <a id="L358"></a>S.next();
<a id="L359"></a>}


<a id="L362"></a>func (S *Scanner) scanRawString(pos token.Position) {
    <a id="L363"></a><span class="comment">// &#39;`&#39; already consumed</span>

    <a id="L365"></a>for S.ch != &#39;`&#39; {
        <a id="L366"></a>ch := S.ch;
        <a id="L367"></a>S.next();
        <a id="L368"></a>if ch &lt; 0 {
            <a id="L369"></a>S.error(pos, &#34;string not terminated&#34;);
            <a id="L370"></a>break;
        <a id="L371"></a>}
    <a id="L372"></a>}

    <a id="L374"></a>S.next();
<a id="L375"></a>}


<a id="L378"></a><span class="comment">// Helper functions for scanning multi-byte tokens such as &gt;&gt; += &gt;&gt;= .</span>
<a id="L379"></a><span class="comment">// Different routines recognize different length tok_i based on matches</span>
<a id="L380"></a><span class="comment">// of ch_i. If a token ends in &#39;=&#39;, the result is tok1 or tok3</span>
<a id="L381"></a><span class="comment">// respectively. Otherwise, the result is tok0 if there was no other</span>
<a id="L382"></a><span class="comment">// matching character, or tok2 if the matching character was ch2.</span>

<a id="L384"></a>func (S *Scanner) switch2(tok0, tok1 token.Token) token.Token {
    <a id="L385"></a>if S.ch == &#39;=&#39; {
        <a id="L386"></a>S.next();
        <a id="L387"></a>return tok1;
    <a id="L388"></a>}
    <a id="L389"></a>return tok0;
<a id="L390"></a>}


<a id="L393"></a>func (S *Scanner) switch3(tok0, tok1 token.Token, ch2 int, tok2 token.Token) token.Token {
    <a id="L394"></a>if S.ch == &#39;=&#39; {
        <a id="L395"></a>S.next();
        <a id="L396"></a>return tok1;
    <a id="L397"></a>}
    <a id="L398"></a>if S.ch == ch2 {
        <a id="L399"></a>S.next();
        <a id="L400"></a>return tok2;
    <a id="L401"></a>}
    <a id="L402"></a>return tok0;
<a id="L403"></a>}


<a id="L406"></a>func (S *Scanner) switch4(tok0, tok1 token.Token, ch2 int, tok2, tok3 token.Token) token.Token {
    <a id="L407"></a>if S.ch == &#39;=&#39; {
        <a id="L408"></a>S.next();
        <a id="L409"></a>return tok1;
    <a id="L410"></a>}
    <a id="L411"></a>if S.ch == ch2 {
        <a id="L412"></a>S.next();
        <a id="L413"></a>if S.ch == &#39;=&#39; {
            <a id="L414"></a>S.next();
            <a id="L415"></a>return tok3;
        <a id="L416"></a>}
        <a id="L417"></a>return tok2;
    <a id="L418"></a>}
    <a id="L419"></a>return tok0;
<a id="L420"></a>}


<a id="L423"></a><span class="comment">// Scan scans the next token and returns the token position pos,</span>
<a id="L424"></a><span class="comment">// the token tok, and the literal text lit corresponding to the</span>
<a id="L425"></a><span class="comment">// token. The source end is indicated by token.EOF.</span>
<a id="L426"></a><span class="comment">//</span>
<a id="L427"></a><span class="comment">// For more tolerant parsing, Scan will return a valid token if</span>
<a id="L428"></a><span class="comment">// possible even if a syntax error was encountered. Thus, even</span>
<a id="L429"></a><span class="comment">// if the resulting token sequence contains no illegal tokens,</span>
<a id="L430"></a><span class="comment">// a client may not assume that no error occurred. Instead it</span>
<a id="L431"></a><span class="comment">// must check the scanner&#39;s ErrorCount or the number of calls</span>
<a id="L432"></a><span class="comment">// of the error handler, if there was one installed.</span>
<a id="L433"></a><span class="comment">//</span>
<a id="L434"></a>func (S *Scanner) Scan() (pos token.Position, tok token.Token, lit []byte) {
<a id="L435"></a>scan_again:
    <a id="L436"></a><span class="comment">// skip white space</span>
    <a id="L437"></a>for S.ch == &#39; &#39; || S.ch == &#39;\t&#39; || S.ch == &#39;\n&#39; || S.ch == &#39;\r&#39; {
        <a id="L438"></a>S.next()
    <a id="L439"></a>}

    <a id="L441"></a><span class="comment">// current token start</span>
    <a id="L442"></a>pos, tok = S.pos, token.ILLEGAL;

    <a id="L444"></a><span class="comment">// determine token value</span>
    <a id="L445"></a>switch ch := S.ch; {
    <a id="L446"></a>case isLetter(ch):
        <a id="L447"></a>tok = S.scanIdentifier()
    <a id="L448"></a>case digitVal(ch) &lt; 10:
        <a id="L449"></a>tok = S.scanNumber(false)
    <a id="L450"></a>default:
        <a id="L451"></a>S.next(); <span class="comment">// always make progress</span>
        <a id="L452"></a>switch ch {
        <a id="L453"></a>case -1:
            <a id="L454"></a>tok = token.EOF
        <a id="L455"></a>case &#39;&#34;&#39;:
            <a id="L456"></a>tok = token.STRING;
            <a id="L457"></a>S.scanString(pos);
        <a id="L458"></a>case &#39;\&#39;&#39;:
            <a id="L459"></a>tok = token.CHAR;
            <a id="L460"></a>S.scanChar(pos);
        <a id="L461"></a>case &#39;`&#39;:
            <a id="L462"></a>tok = token.STRING;
            <a id="L463"></a>S.scanRawString(pos);
        <a id="L464"></a>case &#39;:&#39;:
            <a id="L465"></a>tok = S.switch2(token.COLON, token.DEFINE)
        <a id="L466"></a>case &#39;.&#39;:
            <a id="L467"></a>if digitVal(S.ch) &lt; 10 {
                <a id="L468"></a>tok = S.scanNumber(true)
            <a id="L469"></a>} else if S.ch == &#39;.&#39; {
                <a id="L470"></a>S.next();
                <a id="L471"></a>if S.ch == &#39;.&#39; {
                    <a id="L472"></a>S.next();
                    <a id="L473"></a>tok = token.ELLIPSIS;
                <a id="L474"></a>}
            <a id="L475"></a>} else {
                <a id="L476"></a>tok = token.PERIOD
            <a id="L477"></a>}
        <a id="L478"></a>case &#39;,&#39;:
            <a id="L479"></a>tok = token.COMMA
        <a id="L480"></a>case &#39;;&#39;:
            <a id="L481"></a>tok = token.SEMICOLON
        <a id="L482"></a>case &#39;(&#39;:
            <a id="L483"></a>tok = token.LPAREN
        <a id="L484"></a>case &#39;)&#39;:
            <a id="L485"></a>tok = token.RPAREN
        <a id="L486"></a>case &#39;[&#39;:
            <a id="L487"></a>tok = token.LBRACK
        <a id="L488"></a>case &#39;]&#39;:
            <a id="L489"></a>tok = token.RBRACK
        <a id="L490"></a>case &#39;{&#39;:
            <a id="L491"></a>tok = token.LBRACE
        <a id="L492"></a>case &#39;}&#39;:
            <a id="L493"></a>tok = token.RBRACE
        <a id="L494"></a>case &#39;+&#39;:
            <a id="L495"></a>tok = S.switch3(token.ADD, token.ADD_ASSIGN, &#39;+&#39;, token.INC)
        <a id="L496"></a>case &#39;-&#39;:
            <a id="L497"></a>tok = S.switch3(token.SUB, token.SUB_ASSIGN, &#39;-&#39;, token.DEC)
        <a id="L498"></a>case &#39;*&#39;:
            <a id="L499"></a>tok = S.switch2(token.MUL, token.MUL_ASSIGN)
        <a id="L500"></a>case &#39;/&#39;:
            <a id="L501"></a>if S.ch == &#39;/&#39; || S.ch == &#39;*&#39; {
                <a id="L502"></a>S.scanComment(pos);
                <a id="L503"></a>tok = token.COMMENT;
                <a id="L504"></a>if S.mode&amp;ScanComments == 0 {
                    <a id="L505"></a>goto scan_again
                <a id="L506"></a>}
            <a id="L507"></a>} else {
                <a id="L508"></a>tok = S.switch2(token.QUO, token.QUO_ASSIGN)
            <a id="L509"></a>}
        <a id="L510"></a>case &#39;%&#39;:
            <a id="L511"></a>tok = S.switch2(token.REM, token.REM_ASSIGN)
        <a id="L512"></a>case &#39;^&#39;:
            <a id="L513"></a>tok = S.switch2(token.XOR, token.XOR_ASSIGN)
        <a id="L514"></a>case &#39;&lt;&#39;:
            <a id="L515"></a>if S.ch == &#39;-&#39; {
                <a id="L516"></a>S.next();
                <a id="L517"></a>tok = token.ARROW;
            <a id="L518"></a>} else {
                <a id="L519"></a>tok = S.switch4(token.LSS, token.LEQ, &#39;&lt;&#39;, token.SHL, token.SHL_ASSIGN)
            <a id="L520"></a>}
        <a id="L521"></a>case &#39;&gt;&#39;:
            <a id="L522"></a>tok = S.switch4(token.GTR, token.GEQ, &#39;&gt;&#39;, token.SHR, token.SHR_ASSIGN)
        <a id="L523"></a>case &#39;=&#39;:
            <a id="L524"></a>tok = S.switch2(token.ASSIGN, token.EQL)
        <a id="L525"></a>case &#39;!&#39;:
            <a id="L526"></a>tok = S.switch2(token.NOT, token.NEQ)
        <a id="L527"></a>case &#39;&amp;&#39;:
            <a id="L528"></a>if S.ch == &#39;^&#39; {
                <a id="L529"></a>S.next();
                <a id="L530"></a>tok = S.switch2(token.AND_NOT, token.AND_NOT_ASSIGN);
            <a id="L531"></a>} else {
                <a id="L532"></a>tok = S.switch3(token.AND, token.AND_ASSIGN, &#39;&amp;&#39;, token.LAND)
            <a id="L533"></a>}
        <a id="L534"></a>case &#39;|&#39;:
            <a id="L535"></a>tok = S.switch3(token.OR, token.OR_ASSIGN, &#39;|&#39;, token.LOR)
        <a id="L536"></a>default:
            <a id="L537"></a>if S.mode&amp;AllowIllegalChars == 0 {
                <a id="L538"></a>S.error(pos, &#34;illegal character &#34;+charString(ch))
            <a id="L539"></a>}
        <a id="L540"></a>}
    <a id="L541"></a>}

    <a id="L543"></a>return pos, tok, S.src[pos.Offset:S.pos.Offset];
<a id="L544"></a>}


<a id="L547"></a><span class="comment">// Tokenize calls a function f with the token position, token value, and token</span>
<a id="L548"></a><span class="comment">// text for each token in the source src. The other parameters have the same</span>
<a id="L549"></a><span class="comment">// meaning as for the Init function. Tokenize keeps scanning until f returns</span>
<a id="L550"></a><span class="comment">// false (usually when the token value is token.EOF). The result is the number</span>
<a id="L551"></a><span class="comment">// of errors encountered.</span>
<a id="L552"></a><span class="comment">//</span>
<a id="L553"></a>func Tokenize(filename string, src []byte, err ErrorHandler, mode uint, f func(pos token.Position, tok token.Token, lit []byte) bool) int {
    <a id="L554"></a>var s Scanner;
    <a id="L555"></a>s.Init(filename, src, err, mode);
    <a id="L556"></a>for f(s.Scan()) {
        <a id="L557"></a><span class="comment">// action happens in f</span>
    <a id="L558"></a>}
    <a id="L559"></a>return s.ErrorCount;
<a id="L560"></a>}
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
