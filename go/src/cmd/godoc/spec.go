<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/godoc/spec.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/godoc/spec.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file contains the mechanism to &#34;linkify&#34; html source</span>
<a id="L6"></a><span class="comment">// text containing EBNF sections (as found in go_spec.html).</span>
<a id="L7"></a><span class="comment">// The result is the input source text with the EBNF sections</span>
<a id="L8"></a><span class="comment">// modified such that identifiers are linked to the respective</span>
<a id="L9"></a><span class="comment">// definitions.</span>

<a id="L11"></a>package main

<a id="L13"></a>import (
    <a id="L14"></a>&#34;bytes&#34;;
    <a id="L15"></a>&#34;fmt&#34;;
    <a id="L16"></a>&#34;go/scanner&#34;;
    <a id="L17"></a>&#34;go/token&#34;;
    <a id="L18"></a>&#34;io&#34;;
    <a id="L19"></a>&#34;strings&#34;;
<a id="L20"></a>)


<a id="L23"></a>type ebnfParser struct {
    <a id="L24"></a>out     io.Writer; <span class="comment">// parser output</span>
    <a id="L25"></a>src     []byte;    <span class="comment">// parser source</span>
    <a id="L26"></a>scanner scanner.Scanner;
    <a id="L27"></a>prev    int;            <span class="comment">// offset of previous token</span>
    <a id="L28"></a>pos     token.Position; <span class="comment">// token position</span>
    <a id="L29"></a>tok     token.Token;    <span class="comment">// one token look-ahead</span>
    <a id="L30"></a>lit     []byte;         <span class="comment">// token literal</span>
<a id="L31"></a>}


<a id="L34"></a>func (p *ebnfParser) flush() {
    <a id="L35"></a>p.out.Write(p.src[p.prev:p.pos.Offset]);
    <a id="L36"></a>p.prev = p.pos.Offset;
<a id="L37"></a>}


<a id="L40"></a>func (p *ebnfParser) next() {
    <a id="L41"></a>p.flush();
    <a id="L42"></a>p.pos, p.tok, p.lit = p.scanner.Scan();
    <a id="L43"></a>if p.tok.IsKeyword() {
        <a id="L44"></a><span class="comment">// TODO Should keyword mapping always happen outside scanner?</span>
        <a id="L45"></a><span class="comment">//      Or should there be a flag to scanner to enable keyword mapping?</span>
        <a id="L46"></a>p.tok = token.IDENT
    <a id="L47"></a>}
<a id="L48"></a>}


<a id="L51"></a>func (p *ebnfParser) Error(pos token.Position, msg string) {
    <a id="L52"></a>fmt.Fprintf(p.out, `&lt;span class=&#34;alert&#34;&gt;error: %s&lt;/span&gt;`, msg)
<a id="L53"></a>}


<a id="L56"></a>func (p *ebnfParser) errorExpected(pos token.Position, msg string) {
    <a id="L57"></a>msg = &#34;expected &#34; + msg;
    <a id="L58"></a>if pos.Offset == p.pos.Offset {
        <a id="L59"></a><span class="comment">// the error happened at the current position;</span>
        <a id="L60"></a><span class="comment">// make the error message more specific</span>
        <a id="L61"></a>msg += &#34;, found &#39;&#34; + p.tok.String() + &#34;&#39;&#34;;
        <a id="L62"></a>if p.tok.IsLiteral() {
            <a id="L63"></a>msg += &#34; &#34; + string(p.lit)
        <a id="L64"></a>}
    <a id="L65"></a>}
    <a id="L66"></a>p.Error(pos, msg);
<a id="L67"></a>}


<a id="L70"></a>func (p *ebnfParser) expect(tok token.Token) token.Position {
    <a id="L71"></a>pos := p.pos;
    <a id="L72"></a>if p.tok != tok {
        <a id="L73"></a>p.errorExpected(pos, &#34;&#39;&#34;+tok.String()+&#34;&#39;&#34;)
    <a id="L74"></a>}
    <a id="L75"></a>p.next(); <span class="comment">// make progress in any case</span>
    <a id="L76"></a>return pos;
<a id="L77"></a>}


<a id="L80"></a>func (p *ebnfParser) parseIdentifier(def bool) {
    <a id="L81"></a>name := string(p.lit);
    <a id="L82"></a>p.expect(token.IDENT);
    <a id="L83"></a>if def {
        <a id="L84"></a>fmt.Fprintf(p.out, `&lt;a id=&#34;%s&#34;&gt;%s&lt;/a&gt;`, name, name)
    <a id="L85"></a>} else {
        <a id="L86"></a>fmt.Fprintf(p.out, `&lt;a href=&#34;#%s&#34; class=&#34;noline&#34;&gt;%s&lt;/a&gt;`, name, name)
    <a id="L87"></a>}
    <a id="L88"></a>p.prev += len(name); <span class="comment">// skip identifier when calling flush</span>
<a id="L89"></a>}


<a id="L92"></a>func (p *ebnfParser) parseTerm() bool {
    <a id="L93"></a>switch p.tok {
    <a id="L94"></a>case token.IDENT:
        <a id="L95"></a>p.parseIdentifier(false)

    <a id="L97"></a>case token.STRING:
        <a id="L98"></a>p.next();
        <a id="L99"></a>if p.tok == token.ELLIPSIS {
            <a id="L100"></a>p.next();
            <a id="L101"></a>p.expect(token.STRING);
        <a id="L102"></a>}

    <a id="L104"></a>case token.LPAREN:
        <a id="L105"></a>p.next();
        <a id="L106"></a>p.parseExpression();
        <a id="L107"></a>p.expect(token.RPAREN);

    <a id="L109"></a>case token.LBRACK:
        <a id="L110"></a>p.next();
        <a id="L111"></a>p.parseExpression();
        <a id="L112"></a>p.expect(token.RBRACK);

    <a id="L114"></a>case token.LBRACE:
        <a id="L115"></a>p.next();
        <a id="L116"></a>p.parseExpression();
        <a id="L117"></a>p.expect(token.RBRACE);

    <a id="L119"></a>default:
        <a id="L120"></a>return false
    <a id="L121"></a>}

    <a id="L123"></a>return true;
<a id="L124"></a>}


<a id="L127"></a>func (p *ebnfParser) parseSequence() {
    <a id="L128"></a>for p.parseTerm() {
    <a id="L129"></a>}
<a id="L130"></a>}


<a id="L133"></a>func (p *ebnfParser) parseExpression() {
    <a id="L134"></a>for {
        <a id="L135"></a>p.parseSequence();
        <a id="L136"></a>if p.tok != token.OR {
            <a id="L137"></a>break
        <a id="L138"></a>}
        <a id="L139"></a>p.next();
    <a id="L140"></a>}
<a id="L141"></a>}


<a id="L144"></a>func (p *ebnfParser) parseProduction() {
    <a id="L145"></a>p.parseIdentifier(true);
    <a id="L146"></a>p.expect(token.ASSIGN);
    <a id="L147"></a>p.parseExpression();
    <a id="L148"></a>p.expect(token.PERIOD);
<a id="L149"></a>}


<a id="L152"></a>func (p *ebnfParser) parse(out io.Writer, src []byte) {
    <a id="L153"></a><span class="comment">// initialize ebnfParser</span>
    <a id="L154"></a>p.out = out;
    <a id="L155"></a>p.src = src;
    <a id="L156"></a>p.scanner.Init(&#34;&#34;, src, p, 0);
    <a id="L157"></a>p.next(); <span class="comment">// initializes pos, tok, lit</span>

    <a id="L159"></a><span class="comment">// process source</span>
    <a id="L160"></a>for p.tok != token.EOF {
        <a id="L161"></a>p.parseProduction()
    <a id="L162"></a>}
    <a id="L163"></a>p.flush();
<a id="L164"></a>}


<a id="L167"></a><span class="comment">// Markers around EBNF sections</span>
<a id="L168"></a>var (
    <a id="L169"></a>openTag  = strings.Bytes(`&lt;pre class=&#34;ebnf&#34;&gt;`);
    <a id="L170"></a>closeTag = strings.Bytes(`&lt;/pre&gt;`);
<a id="L171"></a>)


<a id="L174"></a>func linkify(out io.Writer, src []byte) {
    <a id="L175"></a>for len(src) &gt; 0 {
        <a id="L176"></a>n := len(src);

        <a id="L178"></a><span class="comment">// i: beginning of EBNF text (or end of source)</span>
        <a id="L179"></a>i := bytes.Index(src, openTag);
        <a id="L180"></a>if i &lt; 0 {
            <a id="L181"></a>i = n - len(openTag)
        <a id="L182"></a>}
        <a id="L183"></a>i += len(openTag);

        <a id="L185"></a><span class="comment">// j: end of EBNF text (or end of source)</span>
        <a id="L186"></a>j := bytes.Index(src[i:n], closeTag); <span class="comment">// close marker</span>
        <a id="L187"></a>if j &lt; 0 {
            <a id="L188"></a>j = n - i
        <a id="L189"></a>}
        <a id="L190"></a>j += i;

        <a id="L192"></a><span class="comment">// write text before EBNF</span>
        <a id="L193"></a>out.Write(src[0:i]);
        <a id="L194"></a><span class="comment">// parse and write EBNF</span>
        <a id="L195"></a>var p ebnfParser;
        <a id="L196"></a>p.parse(out, src[i:j]);

        <a id="L198"></a><span class="comment">// advance</span>
        <a id="L199"></a>src = src[j:n];
    <a id="L200"></a>}
<a id="L201"></a>}
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
