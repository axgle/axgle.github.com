<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/regexp/regexp.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/regexp/regexp.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package regexp implements a simple regular expression library.</span>
<a id="L6"></a><span class="comment">//</span>
<a id="L7"></a><span class="comment">// The syntax of the regular expressions accepted is:</span>
<a id="L8"></a><span class="comment">//</span>
<a id="L9"></a><span class="comment">//	regexp:</span>
<a id="L10"></a><span class="comment">//		concatenation { &#39;|&#39; concatenation }</span>
<a id="L11"></a><span class="comment">//	concatenation:</span>
<a id="L12"></a><span class="comment">//		{ closure }</span>
<a id="L13"></a><span class="comment">//	closure:</span>
<a id="L14"></a><span class="comment">//		term [ &#39;*&#39; | &#39;+&#39; | &#39;?&#39; ]</span>
<a id="L15"></a><span class="comment">//	term:</span>
<a id="L16"></a><span class="comment">//		&#39;^&#39;</span>
<a id="L17"></a><span class="comment">//		&#39;$&#39;</span>
<a id="L18"></a><span class="comment">//		&#39;.&#39;</span>
<a id="L19"></a><span class="comment">//		character</span>
<a id="L20"></a><span class="comment">//		&#39;[&#39; [ &#39;^&#39; ] character-ranges &#39;]&#39;</span>
<a id="L21"></a><span class="comment">//		&#39;(&#39; regexp &#39;)&#39;</span>
<a id="L22"></a><span class="comment">//</span>
<a id="L23"></a>package regexp

<a id="L25"></a>import (
    <a id="L26"></a>&#34;bytes&#34;;
    <a id="L27"></a>&#34;container/vector&#34;;
    <a id="L28"></a>&#34;io&#34;;
    <a id="L29"></a>&#34;os&#34;;
    <a id="L30"></a>&#34;utf8&#34;;
<a id="L31"></a>)

<a id="L33"></a>var debug = false

<a id="L35"></a><span class="comment">// Error codes returned by failures to parse an expression.</span>
<a id="L36"></a>var (
    <a id="L37"></a>ErrInternal            = os.NewError(&#34;internal error&#34;);
    <a id="L38"></a>ErrUnmatchedLpar       = os.NewError(&#34;unmatched &#39;(&#39;&#34;);
    <a id="L39"></a>ErrUnmatchedRpar       = os.NewError(&#34;unmatched &#39;)&#39;&#34;);
    <a id="L40"></a>ErrUnmatchedLbkt       = os.NewError(&#34;unmatched &#39;[&#39;&#34;);
    <a id="L41"></a>ErrUnmatchedRbkt       = os.NewError(&#34;unmatched &#39;]&#39;&#34;);
    <a id="L42"></a>ErrBadRange            = os.NewError(&#34;bad range in character class&#34;);
    <a id="L43"></a>ErrExtraneousBackslash = os.NewError(&#34;extraneous backslash&#34;);
    <a id="L44"></a>ErrBadClosure          = os.NewError(&#34;repeated closure (**, ++, etc.)&#34;);
    <a id="L45"></a>ErrBareClosure         = os.NewError(&#34;closure applies to nothing&#34;);
    <a id="L46"></a>ErrBadBackslash        = os.NewError(&#34;illegal backslash escape&#34;);
<a id="L47"></a>)

<a id="L49"></a><span class="comment">// An instruction executed by the NFA</span>
<a id="L50"></a>type instr interface {
    <a id="L51"></a>kind() int;   <span class="comment">// the type of this instruction: _CHAR, _ANY, etc.</span>
    <a id="L52"></a>next() instr; <span class="comment">// the instruction to execute after this one</span>
    <a id="L53"></a>setNext(i instr);
    <a id="L54"></a>index() int;
    <a id="L55"></a>setIndex(i int);
    <a id="L56"></a>print();
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// Fields and methods common to all instructions</span>
<a id="L60"></a>type common struct {
    <a id="L61"></a>_next  instr;
    <a id="L62"></a>_index int;
<a id="L63"></a>}

<a id="L65"></a>func (c *common) next() instr     { return c._next }
<a id="L66"></a>func (c *common) setNext(i instr) { c._next = i }
<a id="L67"></a>func (c *common) index() int      { return c._index }
<a id="L68"></a>func (c *common) setIndex(i int)  { c._index = i }

<a id="L70"></a><span class="comment">// Regexp is the representation of a compiled regular expression.</span>
<a id="L71"></a><span class="comment">// The public interface is entirely through methods.</span>
<a id="L72"></a>type Regexp struct {
    <a id="L73"></a>expr  string; <span class="comment">// the original expression</span>
    <a id="L74"></a>inst  *vector.Vector;
    <a id="L75"></a>start instr;
    <a id="L76"></a>nbra  int; <span class="comment">// number of brackets in expression, for subexpressions</span>
<a id="L77"></a>}

<a id="L79"></a>const (
    <a id="L80"></a>_START      = iota; <span class="comment">// beginning of program</span>
    <a id="L81"></a>_END;       <span class="comment">// end of program: success</span>
    <a id="L82"></a>_BOT;       <span class="comment">// &#39;^&#39; beginning of text</span>
    <a id="L83"></a>_EOT;       <span class="comment">// &#39;$&#39; end of text</span>
    <a id="L84"></a>_CHAR;      <span class="comment">// &#39;a&#39; regular character</span>
    <a id="L85"></a>_CHARCLASS; <span class="comment">// [a-z] character class</span>
    <a id="L86"></a>_ANY;       <span class="comment">// &#39;.&#39; any character including newline</span>
    <a id="L87"></a>_NOTNL;     <span class="comment">// [^\n] special case: any character but newline</span>
    <a id="L88"></a>_BRA;       <span class="comment">// &#39;(&#39; parenthesized expression</span>
    <a id="L89"></a>_EBRA;      <span class="comment">// &#39;)&#39;; end of &#39;(&#39; parenthesized expression</span>
    <a id="L90"></a>_ALT;       <span class="comment">// &#39;|&#39; alternation</span>
    <a id="L91"></a>_NOP;       <span class="comment">// do nothing; makes it easy to link without patching</span>
<a id="L92"></a>)

<a id="L94"></a><span class="comment">// --- START start of program</span>
<a id="L95"></a>type _Start struct {
    <a id="L96"></a>common;
<a id="L97"></a>}

<a id="L99"></a>func (start *_Start) kind() int { return _START }
<a id="L100"></a>func (start *_Start) print()    { print(&#34;start&#34;) }

<a id="L102"></a><span class="comment">// --- END end of program</span>
<a id="L103"></a>type _End struct {
    <a id="L104"></a>common;
<a id="L105"></a>}

<a id="L107"></a>func (end *_End) kind() int { return _END }
<a id="L108"></a>func (end *_End) print()    { print(&#34;end&#34;) }

<a id="L110"></a><span class="comment">// --- BOT beginning of text</span>
<a id="L111"></a>type _Bot struct {
    <a id="L112"></a>common;
<a id="L113"></a>}

<a id="L115"></a>func (bot *_Bot) kind() int { return _BOT }
<a id="L116"></a>func (bot *_Bot) print()    { print(&#34;bot&#34;) }

<a id="L118"></a><span class="comment">// --- EOT end of text</span>
<a id="L119"></a>type _Eot struct {
    <a id="L120"></a>common;
<a id="L121"></a>}

<a id="L123"></a>func (eot *_Eot) kind() int { return _EOT }
<a id="L124"></a>func (eot *_Eot) print()    { print(&#34;eot&#34;) }

<a id="L126"></a><span class="comment">// --- CHAR a regular character</span>
<a id="L127"></a>type _Char struct {
    <a id="L128"></a>common;
    <a id="L129"></a>char int;
<a id="L130"></a>}

<a id="L132"></a>func (char *_Char) kind() int { return _CHAR }
<a id="L133"></a>func (char *_Char) print()    { print(&#34;char &#34;, string(char.char)) }

<a id="L135"></a>func newChar(char int) *_Char {
    <a id="L136"></a>c := new(_Char);
    <a id="L137"></a>c.char = char;
    <a id="L138"></a>return c;
<a id="L139"></a>}

<a id="L141"></a><span class="comment">// --- CHARCLASS [a-z]</span>

<a id="L143"></a>type _CharClass struct {
    <a id="L144"></a>common;
    <a id="L145"></a>char   int;
    <a id="L146"></a>negate bool; <span class="comment">// is character class negated? ([^a-z])</span>
    <a id="L147"></a><span class="comment">// vector of int, stored pairwise: [a-z] is (a,z); x is (x,x):</span>
    <a id="L148"></a>ranges *vector.IntVector;
<a id="L149"></a>}

<a id="L151"></a>func (cclass *_CharClass) kind() int { return _CHARCLASS }

<a id="L153"></a>func (cclass *_CharClass) print() {
    <a id="L154"></a>print(&#34;charclass&#34;);
    <a id="L155"></a>if cclass.negate {
        <a id="L156"></a>print(&#34; (negated)&#34;)
    <a id="L157"></a>}
    <a id="L158"></a>for i := 0; i &lt; cclass.ranges.Len(); i += 2 {
        <a id="L159"></a>l := cclass.ranges.At(i);
        <a id="L160"></a>r := cclass.ranges.At(i + 1);
        <a id="L161"></a>if l == r {
            <a id="L162"></a>print(&#34; [&#34;, string(l), &#34;]&#34;)
        <a id="L163"></a>} else {
            <a id="L164"></a>print(&#34; [&#34;, string(l), &#34;-&#34;, string(r), &#34;]&#34;)
        <a id="L165"></a>}
    <a id="L166"></a>}
<a id="L167"></a>}

<a id="L169"></a>func (cclass *_CharClass) addRange(a, b int) {
    <a id="L170"></a><span class="comment">// range is a through b inclusive</span>
    <a id="L171"></a>cclass.ranges.Push(a);
    <a id="L172"></a>cclass.ranges.Push(b);
<a id="L173"></a>}

<a id="L175"></a>func (cclass *_CharClass) matches(c int) bool {
    <a id="L176"></a>for i := 0; i &lt; cclass.ranges.Len(); i = i + 2 {
        <a id="L177"></a>min := cclass.ranges.At(i);
        <a id="L178"></a>max := cclass.ranges.At(i + 1);
        <a id="L179"></a>if min &lt;= c &amp;&amp; c &lt;= max {
            <a id="L180"></a>return !cclass.negate
        <a id="L181"></a>}
    <a id="L182"></a>}
    <a id="L183"></a>return cclass.negate;
<a id="L184"></a>}

<a id="L186"></a>func newCharClass() *_CharClass {
    <a id="L187"></a>c := new(_CharClass);
    <a id="L188"></a>c.ranges = vector.NewIntVector(0);
    <a id="L189"></a>return c;
<a id="L190"></a>}

<a id="L192"></a><span class="comment">// --- ANY any character</span>
<a id="L193"></a>type _Any struct {
    <a id="L194"></a>common;
<a id="L195"></a>}

<a id="L197"></a>func (any *_Any) kind() int { return _ANY }
<a id="L198"></a>func (any *_Any) print()    { print(&#34;any&#34;) }

<a id="L200"></a><span class="comment">// --- NOTNL any character but newline</span>
<a id="L201"></a>type _NotNl struct {
    <a id="L202"></a>common;
<a id="L203"></a>}

<a id="L205"></a>func (notnl *_NotNl) kind() int { return _NOTNL }
<a id="L206"></a>func (notnl *_NotNl) print()    { print(&#34;notnl&#34;) }

<a id="L208"></a><span class="comment">// --- BRA parenthesized expression</span>
<a id="L209"></a>type _Bra struct {
    <a id="L210"></a>common;
    <a id="L211"></a>n   int; <span class="comment">// subexpression number</span>
<a id="L212"></a>}

<a id="L214"></a>func (bra *_Bra) kind() int { return _BRA }
<a id="L215"></a>func (bra *_Bra) print()    { print(&#34;bra&#34;, bra.n) }

<a id="L217"></a><span class="comment">// --- EBRA end of parenthesized expression</span>
<a id="L218"></a>type _Ebra struct {
    <a id="L219"></a>common;
    <a id="L220"></a>n   int; <span class="comment">// subexpression number</span>
<a id="L221"></a>}

<a id="L223"></a>func (ebra *_Ebra) kind() int { return _EBRA }
<a id="L224"></a>func (ebra *_Ebra) print()    { print(&#34;ebra &#34;, ebra.n) }

<a id="L226"></a><span class="comment">// --- ALT alternation</span>
<a id="L227"></a>type _Alt struct {
    <a id="L228"></a>common;
    <a id="L229"></a>left instr; <span class="comment">// other branch</span>
<a id="L230"></a>}

<a id="L232"></a>func (alt *_Alt) kind() int { return _ALT }
<a id="L233"></a>func (alt *_Alt) print()    { print(&#34;alt(&#34;, alt.left.index(), &#34;)&#34;) }

<a id="L235"></a><span class="comment">// --- NOP no operation</span>
<a id="L236"></a>type _Nop struct {
    <a id="L237"></a>common;
<a id="L238"></a>}

<a id="L240"></a>func (nop *_Nop) kind() int { return _NOP }
<a id="L241"></a>func (nop *_Nop) print()    { print(&#34;nop&#34;) }

<a id="L243"></a>func (re *Regexp) add(i instr) instr {
    <a id="L244"></a>i.setIndex(re.inst.Len());
    <a id="L245"></a>re.inst.Push(i);
    <a id="L246"></a>return i;
<a id="L247"></a>}

<a id="L249"></a>type parser struct {
    <a id="L250"></a>re    *Regexp;
    <a id="L251"></a>error os.Error;
    <a id="L252"></a>nlpar int; <span class="comment">// number of unclosed lpars</span>
    <a id="L253"></a>pos   int;
    <a id="L254"></a>ch    int;
<a id="L255"></a>}

<a id="L257"></a>const endOfFile = -1

<a id="L259"></a>func (p *parser) c() int { return p.ch }

<a id="L261"></a>func (p *parser) nextc() int {
    <a id="L262"></a>if p.pos &gt;= len(p.re.expr) {
        <a id="L263"></a>p.ch = endOfFile
    <a id="L264"></a>} else {
        <a id="L265"></a>c, w := utf8.DecodeRuneInString(p.re.expr[p.pos:len(p.re.expr)]);
        <a id="L266"></a>p.ch = c;
        <a id="L267"></a>p.pos += w;
    <a id="L268"></a>}
    <a id="L269"></a>return p.ch;
<a id="L270"></a>}

<a id="L272"></a>func newParser(re *Regexp) *parser {
    <a id="L273"></a>p := new(parser);
    <a id="L274"></a>p.re = re;
    <a id="L275"></a>p.nextc(); <span class="comment">// load p.ch</span>
    <a id="L276"></a>return p;
<a id="L277"></a>}

<a id="L279"></a>func special(c int) bool {
    <a id="L280"></a>s := `\.+*?()|[]^$`;
    <a id="L281"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L282"></a>if c == int(s[i]) {
            <a id="L283"></a>return true
        <a id="L284"></a>}
    <a id="L285"></a>}
    <a id="L286"></a>return false;
<a id="L287"></a>}

<a id="L289"></a>func specialcclass(c int) bool {
    <a id="L290"></a>s := `\-[]`;
    <a id="L291"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L292"></a>if c == int(s[i]) {
            <a id="L293"></a>return true
        <a id="L294"></a>}
    <a id="L295"></a>}
    <a id="L296"></a>return false;
<a id="L297"></a>}

<a id="L299"></a>func (p *parser) charClass() instr {
    <a id="L300"></a>cc := newCharClass();
    <a id="L301"></a>if p.c() == &#39;^&#39; {
        <a id="L302"></a>cc.negate = true;
        <a id="L303"></a>p.nextc();
    <a id="L304"></a>}
    <a id="L305"></a>left := -1;
    <a id="L306"></a>for {
        <a id="L307"></a>switch c := p.c(); c {
        <a id="L308"></a>case &#39;]&#39;, endOfFile:
            <a id="L309"></a>if left &gt;= 0 {
                <a id="L310"></a>p.error = ErrBadRange;
                <a id="L311"></a>return nil;
            <a id="L312"></a>}
            <a id="L313"></a><span class="comment">// Is it [^\n]?</span>
            <a id="L314"></a>if cc.negate &amp;&amp; cc.ranges.Len() == 2 &amp;&amp;
                <a id="L315"></a>cc.ranges.At(0) == &#39;\n&#39; &amp;&amp; cc.ranges.At(1) == &#39;\n&#39; {
                <a id="L316"></a>nl := new(_NotNl);
                <a id="L317"></a>p.re.add(nl);
                <a id="L318"></a>return nl;
            <a id="L319"></a>}
            <a id="L320"></a>p.re.add(cc);
            <a id="L321"></a>return cc;
        <a id="L322"></a>case &#39;-&#39;: <span class="comment">// do this before backslash processing</span>
            <a id="L323"></a>p.error = ErrBadRange;
            <a id="L324"></a>return nil;
        <a id="L325"></a>case &#39;\\&#39;:
            <a id="L326"></a>c = p.nextc();
            <a id="L327"></a>switch {
            <a id="L328"></a>case c == endOfFile:
                <a id="L329"></a>p.error = ErrExtraneousBackslash;
                <a id="L330"></a>return nil;
            <a id="L331"></a>case c == &#39;n&#39;:
                <a id="L332"></a>c = &#39;\n&#39;
            <a id="L333"></a>case specialcclass(c):
                <a id="L334"></a><span class="comment">// c is as delivered</span>
            <a id="L335"></a>default:
                <a id="L336"></a>p.error = ErrBadBackslash;
                <a id="L337"></a>return nil;
            <a id="L338"></a>}
            <a id="L339"></a>fallthrough;
        <a id="L340"></a>default:
            <a id="L341"></a>p.nextc();
            <a id="L342"></a>switch {
            <a id="L343"></a>case left &lt; 0: <span class="comment">// first of pair</span>
                <a id="L344"></a>if p.c() == &#39;-&#39; { <span class="comment">// range</span>
                    <a id="L345"></a>p.nextc();
                    <a id="L346"></a>left = c;
                <a id="L347"></a>} else { <span class="comment">// single char</span>
                    <a id="L348"></a>cc.addRange(c, c)
                <a id="L349"></a>}
            <a id="L350"></a>case left &lt;= c: <span class="comment">// second of pair</span>
                <a id="L351"></a>cc.addRange(left, c);
                <a id="L352"></a>left = -1;
            <a id="L353"></a>default:
                <a id="L354"></a>p.error = ErrBadRange;
                <a id="L355"></a>return nil;
            <a id="L356"></a>}
        <a id="L357"></a>}
    <a id="L358"></a>}
    <a id="L359"></a>return nil;
<a id="L360"></a>}

<a id="L362"></a>func (p *parser) term() (start, end instr) {
    <a id="L363"></a><span class="comment">// term() is the leaf of the recursion, so it&#39;s sufficient to pick off the</span>
    <a id="L364"></a><span class="comment">// error state here for early exit.</span>
    <a id="L365"></a><span class="comment">// The other functions (closure(), concatenation() etc.) assume</span>
    <a id="L366"></a><span class="comment">// it&#39;s safe to recur to here.</span>
    <a id="L367"></a>if p.error != nil {
        <a id="L368"></a>return
    <a id="L369"></a>}
    <a id="L370"></a>switch c := p.c(); c {
    <a id="L371"></a>case &#39;|&#39;, endOfFile:
        <a id="L372"></a>return nil, nil
    <a id="L373"></a>case &#39;*&#39;, &#39;+&#39;:
        <a id="L374"></a>p.error = ErrBareClosure;
        <a id="L375"></a>return;
    <a id="L376"></a>case &#39;)&#39;:
        <a id="L377"></a>if p.nlpar == 0 {
            <a id="L378"></a>p.error = ErrUnmatchedRpar;
            <a id="L379"></a>return;
        <a id="L380"></a>}
        <a id="L381"></a>return nil, nil;
    <a id="L382"></a>case &#39;]&#39;:
        <a id="L383"></a>p.error = ErrUnmatchedRbkt;
        <a id="L384"></a>return;
    <a id="L385"></a>case &#39;^&#39;:
        <a id="L386"></a>p.nextc();
        <a id="L387"></a>start = p.re.add(new(_Bot));
        <a id="L388"></a>return start, start;
    <a id="L389"></a>case &#39;$&#39;:
        <a id="L390"></a>p.nextc();
        <a id="L391"></a>start = p.re.add(new(_Eot));
        <a id="L392"></a>return start, start;
    <a id="L393"></a>case &#39;.&#39;:
        <a id="L394"></a>p.nextc();
        <a id="L395"></a>start = p.re.add(new(_Any));
        <a id="L396"></a>return start, start;
    <a id="L397"></a>case &#39;[&#39;:
        <a id="L398"></a>p.nextc();
        <a id="L399"></a>start = p.charClass();
        <a id="L400"></a>if p.error != nil {
            <a id="L401"></a>return
        <a id="L402"></a>}
        <a id="L403"></a>if p.c() != &#39;]&#39; {
            <a id="L404"></a>p.error = ErrUnmatchedLbkt;
            <a id="L405"></a>return;
        <a id="L406"></a>}
        <a id="L407"></a>p.nextc();
        <a id="L408"></a>return start, start;
    <a id="L409"></a>case &#39;(&#39;:
        <a id="L410"></a>p.nextc();
        <a id="L411"></a>p.nlpar++;
        <a id="L412"></a>p.re.nbra++; <span class="comment">// increment first so first subexpr is \1</span>
        <a id="L413"></a>nbra := p.re.nbra;
        <a id="L414"></a>start, end = p.regexp();
        <a id="L415"></a>if p.c() != &#39;)&#39; {
            <a id="L416"></a>p.error = ErrUnmatchedLpar;
            <a id="L417"></a>return;
        <a id="L418"></a>}
        <a id="L419"></a>p.nlpar--;
        <a id="L420"></a>p.nextc();
        <a id="L421"></a>bra := new(_Bra);
        <a id="L422"></a>p.re.add(bra);
        <a id="L423"></a>ebra := new(_Ebra);
        <a id="L424"></a>p.re.add(ebra);
        <a id="L425"></a>bra.n = nbra;
        <a id="L426"></a>ebra.n = nbra;
        <a id="L427"></a>if start == nil {
            <a id="L428"></a>if end == nil {
                <a id="L429"></a>p.error = ErrInternal;
                <a id="L430"></a>return;
            <a id="L431"></a>}
            <a id="L432"></a>start = ebra;
        <a id="L433"></a>} else {
            <a id="L434"></a>end.setNext(ebra)
        <a id="L435"></a>}
        <a id="L436"></a>bra.setNext(start);
        <a id="L437"></a>return bra, ebra;
    <a id="L438"></a>case &#39;\\&#39;:
        <a id="L439"></a>c = p.nextc();
        <a id="L440"></a>switch {
        <a id="L441"></a>case c == endOfFile:
            <a id="L442"></a>p.error = ErrExtraneousBackslash;
            <a id="L443"></a>return;
        <a id="L444"></a>case c == &#39;n&#39;:
            <a id="L445"></a>c = &#39;\n&#39;
        <a id="L446"></a>case special(c):
            <a id="L447"></a><span class="comment">// c is as delivered</span>
        <a id="L448"></a>default:
            <a id="L449"></a>p.error = ErrBadBackslash;
            <a id="L450"></a>return;
        <a id="L451"></a>}
        <a id="L452"></a>fallthrough;
    <a id="L453"></a>default:
        <a id="L454"></a>p.nextc();
        <a id="L455"></a>start = newChar(c);
        <a id="L456"></a>p.re.add(start);
        <a id="L457"></a>return start, start;
    <a id="L458"></a>}
    <a id="L459"></a>panic(&#34;unreachable&#34;);
<a id="L460"></a>}

<a id="L462"></a>func (p *parser) closure() (start, end instr) {
    <a id="L463"></a>start, end = p.term();
    <a id="L464"></a>if start == nil || p.error != nil {
        <a id="L465"></a>return
    <a id="L466"></a>}
    <a id="L467"></a>switch p.c() {
    <a id="L468"></a>case &#39;*&#39;:
        <a id="L469"></a><span class="comment">// (start,end)*:</span>
        <a id="L470"></a>alt := new(_Alt);
        <a id="L471"></a>p.re.add(alt);
        <a id="L472"></a>end.setNext(alt); <span class="comment">// after end, do alt</span>
        <a id="L473"></a>alt.left = start; <span class="comment">// alternate brach: return to start</span>
        <a id="L474"></a>start = alt;      <span class="comment">// alt becomes new (start, end)</span>
        <a id="L475"></a>end = alt;
    <a id="L476"></a>case &#39;+&#39;:
        <a id="L477"></a><span class="comment">// (start,end)+:</span>
        <a id="L478"></a>alt := new(_Alt);
        <a id="L479"></a>p.re.add(alt);
        <a id="L480"></a>end.setNext(alt); <span class="comment">// after end, do alt</span>
        <a id="L481"></a>alt.left = start; <span class="comment">// alternate brach: return to start</span>
        <a id="L482"></a>end = alt;        <span class="comment">// start is unchanged; end is alt</span>
    <a id="L483"></a>case &#39;?&#39;:
        <a id="L484"></a><span class="comment">// (start,end)?:</span>
        <a id="L485"></a>alt := new(_Alt);
        <a id="L486"></a>p.re.add(alt);
        <a id="L487"></a>nop := new(_Nop);
        <a id="L488"></a>p.re.add(nop);
        <a id="L489"></a>alt.left = start; <span class="comment">// alternate branch is start</span>
        <a id="L490"></a>alt.setNext(nop); <span class="comment">// follow on to nop</span>
        <a id="L491"></a>end.setNext(nop); <span class="comment">// after end, go to nop</span>
        <a id="L492"></a>start = alt;      <span class="comment">// start is now alt</span>
        <a id="L493"></a>end = nop;        <span class="comment">// end is nop pointed to by both branches</span>
    <a id="L494"></a>default:
        <a id="L495"></a>return
    <a id="L496"></a>}
    <a id="L497"></a>switch p.nextc() {
    <a id="L498"></a>case &#39;*&#39;, &#39;+&#39;, &#39;?&#39;:
        <a id="L499"></a>p.error = ErrBadClosure
    <a id="L500"></a>}
    <a id="L501"></a>return;
<a id="L502"></a>}

<a id="L504"></a>func (p *parser) concatenation() (start, end instr) {
    <a id="L505"></a>for {
        <a id="L506"></a>nstart, nend := p.closure();
        <a id="L507"></a>if p.error != nil {
            <a id="L508"></a>return
        <a id="L509"></a>}
        <a id="L510"></a>switch {
        <a id="L511"></a>case nstart == nil: <span class="comment">// end of this concatenation</span>
            <a id="L512"></a>if start == nil { <span class="comment">// this is the empty string</span>
                <a id="L513"></a>nop := p.re.add(new(_Nop));
                <a id="L514"></a>return nop, nop;
            <a id="L515"></a>}
            <a id="L516"></a>return;
        <a id="L517"></a>case start == nil: <span class="comment">// this is first element of concatenation</span>
            <a id="L518"></a>start, end = nstart, nend
        <a id="L519"></a>default:
            <a id="L520"></a>end.setNext(nstart);
            <a id="L521"></a>end = nend;
        <a id="L522"></a>}
    <a id="L523"></a>}
    <a id="L524"></a>panic(&#34;unreachable&#34;);
<a id="L525"></a>}

<a id="L527"></a>func (p *parser) regexp() (start, end instr) {
    <a id="L528"></a>start, end = p.concatenation();
    <a id="L529"></a>if p.error != nil {
        <a id="L530"></a>return
    <a id="L531"></a>}
    <a id="L532"></a>for {
        <a id="L533"></a>switch p.c() {
        <a id="L534"></a>default:
            <a id="L535"></a>return
        <a id="L536"></a>case &#39;|&#39;:
            <a id="L537"></a>p.nextc();
            <a id="L538"></a>nstart, nend := p.concatenation();
            <a id="L539"></a>if p.error != nil {
                <a id="L540"></a>return
            <a id="L541"></a>}
            <a id="L542"></a>alt := new(_Alt);
            <a id="L543"></a>p.re.add(alt);
            <a id="L544"></a>alt.left = start;
            <a id="L545"></a>alt.setNext(nstart);
            <a id="L546"></a>nop := new(_Nop);
            <a id="L547"></a>p.re.add(nop);
            <a id="L548"></a>end.setNext(nop);
            <a id="L549"></a>nend.setNext(nop);
            <a id="L550"></a>start, end = alt, nop;
        <a id="L551"></a>}
    <a id="L552"></a>}
    <a id="L553"></a>panic(&#34;unreachable&#34;);
<a id="L554"></a>}

<a id="L556"></a>func unNop(i instr) instr {
    <a id="L557"></a>for i.kind() == _NOP {
        <a id="L558"></a>i = i.next()
    <a id="L559"></a>}
    <a id="L560"></a>return i;
<a id="L561"></a>}

<a id="L563"></a>func (re *Regexp) eliminateNops() {
    <a id="L564"></a>for i := 0; i &lt; re.inst.Len(); i++ {
        <a id="L565"></a>inst := re.inst.At(i).(instr);
        <a id="L566"></a>if inst.kind() == _END {
            <a id="L567"></a>continue
        <a id="L568"></a>}
        <a id="L569"></a>inst.setNext(unNop(inst.next()));
        <a id="L570"></a>if inst.kind() == _ALT {
            <a id="L571"></a>alt := inst.(*_Alt);
            <a id="L572"></a>alt.left = unNop(alt.left);
        <a id="L573"></a>}
    <a id="L574"></a>}
<a id="L575"></a>}

<a id="L577"></a>func (re *Regexp) dump() {
    <a id="L578"></a>for i := 0; i &lt; re.inst.Len(); i++ {
        <a id="L579"></a>inst := re.inst.At(i).(instr);
        <a id="L580"></a>print(inst.index(), &#34;: &#34;);
        <a id="L581"></a>inst.print();
        <a id="L582"></a>if inst.kind() != _END {
            <a id="L583"></a>print(&#34; -&gt; &#34;, inst.next().index())
        <a id="L584"></a>}
        <a id="L585"></a>print(&#34;\n&#34;);
    <a id="L586"></a>}
<a id="L587"></a>}

<a id="L589"></a>func (re *Regexp) doParse() os.Error {
    <a id="L590"></a>p := newParser(re);
    <a id="L591"></a>start := new(_Start);
    <a id="L592"></a>re.add(start);
    <a id="L593"></a>s, e := p.regexp();
    <a id="L594"></a>if p.error != nil {
        <a id="L595"></a>return p.error
    <a id="L596"></a>}
    <a id="L597"></a>start.setNext(s);
    <a id="L598"></a>re.start = start;
    <a id="L599"></a>e.setNext(re.add(new(_End)));

    <a id="L601"></a>if debug {
        <a id="L602"></a>re.dump();
        <a id="L603"></a>println();
    <a id="L604"></a>}

    <a id="L606"></a>re.eliminateNops();
    <a id="L607"></a>if debug {
        <a id="L608"></a>re.dump();
        <a id="L609"></a>println();
    <a id="L610"></a>}
    <a id="L611"></a>return p.error;
<a id="L612"></a>}

<a id="L614"></a><span class="comment">// Compile parses a regular expression and returns, if successful, a Regexp</span>
<a id="L615"></a><span class="comment">// object that can be used to match against text.</span>
<a id="L616"></a>func Compile(str string) (regexp *Regexp, error os.Error) {
    <a id="L617"></a>regexp = new(Regexp);
    <a id="L618"></a>regexp.expr = str;
    <a id="L619"></a>regexp.inst = vector.New(0);
    <a id="L620"></a>error = regexp.doParse();
    <a id="L621"></a>return;
<a id="L622"></a>}

<a id="L624"></a><span class="comment">// MustCompile is like Compile but panics if the expression cannot be parsed.</span>
<a id="L625"></a><span class="comment">// It simplifies safe initialization of global variables holding compiled regular</span>
<a id="L626"></a><span class="comment">// expressions.</span>
<a id="L627"></a>func MustCompile(str string) *Regexp {
    <a id="L628"></a>regexp, error := Compile(str);
    <a id="L629"></a>if error != nil {
        <a id="L630"></a>panicln(`regexp: compiling &#34;`, str, `&#34;: `, error.String())
    <a id="L631"></a>}
    <a id="L632"></a>return regexp;
<a id="L633"></a>}

<a id="L635"></a>type state struct {
    <a id="L636"></a>inst  instr; <span class="comment">// next instruction to execute</span>
    <a id="L637"></a>match []int; <span class="comment">// pairs of bracketing submatches. 0th is start,end</span>
<a id="L638"></a>}

<a id="L640"></a><span class="comment">// Append new state to to-do list.  Leftmost-longest wins so avoid</span>
<a id="L641"></a><span class="comment">// adding a state that&#39;s already active.</span>
<a id="L642"></a>func addState(s []state, inst instr, match []int) []state {
    <a id="L643"></a>index := inst.index();
    <a id="L644"></a>l := len(s);
    <a id="L645"></a>pos := match[0];
    <a id="L646"></a><span class="comment">// TODO: Once the state is a vector and we can do insert, have inputs always</span>
    <a id="L647"></a><span class="comment">// go in order correctly and this &#34;earlier&#34; test is never necessary,</span>
    <a id="L648"></a>for i := 0; i &lt; l; i++ {
        <a id="L649"></a>if s[i].inst.index() == index &amp;&amp; <span class="comment">// same instruction</span>
            <a id="L650"></a>s[i].match[0] &lt; pos { <span class="comment">// earlier match already going; lefmost wins</span>
            <a id="L651"></a>return s
        <a id="L652"></a>}
    <a id="L653"></a>}
    <a id="L654"></a>if l == cap(s) {
        <a id="L655"></a>s1 := make([]state, 2*l)[0:l];
        <a id="L656"></a>for i := 0; i &lt; l; i++ {
            <a id="L657"></a>s1[i] = s[i]
        <a id="L658"></a>}
        <a id="L659"></a>s = s1;
    <a id="L660"></a>}
    <a id="L661"></a>s = s[0 : l+1];
    <a id="L662"></a>s[l].inst = inst;
    <a id="L663"></a>s[l].match = match;
    <a id="L664"></a>return s;
<a id="L665"></a>}

<a id="L667"></a><span class="comment">// Accepts either string or bytes - the logic is identical either way.</span>
<a id="L668"></a><span class="comment">// If bytes == nil, scan str.</span>
<a id="L669"></a>func (re *Regexp) doExecute(str string, bytes []byte, pos int) []int {
    <a id="L670"></a>var s [2][]state; <span class="comment">// TODO: use a vector when state values (not ptrs) can be vector elements</span>
    <a id="L671"></a>s[0] = make([]state, 10)[0:0];
    <a id="L672"></a>s[1] = make([]state, 10)[0:0];
    <a id="L673"></a>in, out := 0, 1;
    <a id="L674"></a>var final state;
    <a id="L675"></a>found := false;
    <a id="L676"></a>end := len(str);
    <a id="L677"></a>if bytes != nil {
        <a id="L678"></a>end = len(bytes)
    <a id="L679"></a>}
    <a id="L680"></a>for pos &lt;= end {
        <a id="L681"></a>if !found {
            <a id="L682"></a><span class="comment">// prime the pump if we haven&#39;t seen a match yet</span>
            <a id="L683"></a>match := make([]int, 2*(re.nbra+1));
            <a id="L684"></a>for i := 0; i &lt; len(match); i++ {
                <a id="L685"></a>match[i] = -1 <span class="comment">// no match seen; catches cases like &#34;a(b)?c&#34; on &#34;ac&#34;</span>
            <a id="L686"></a>}
            <a id="L687"></a>match[0] = pos;
            <a id="L688"></a>s[out] = addState(s[out], re.start.next(), match);
        <a id="L689"></a>}
        <a id="L690"></a>in, out = out, in;    <span class="comment">// old out state is new in state</span>
        <a id="L691"></a>s[out] = s[out][0:0]; <span class="comment">// clear out state</span>
        <a id="L692"></a>if len(s[in]) == 0 {
            <a id="L693"></a><span class="comment">// machine has completed</span>
            <a id="L694"></a>break
        <a id="L695"></a>}
        <a id="L696"></a>charwidth := 1;
        <a id="L697"></a>c := endOfFile;
        <a id="L698"></a>if pos &lt; end {
            <a id="L699"></a>if bytes == nil {
                <a id="L700"></a>c, charwidth = utf8.DecodeRuneInString(str[pos:end])
            <a id="L701"></a>} else {
                <a id="L702"></a>c, charwidth = utf8.DecodeRune(bytes[pos:end])
            <a id="L703"></a>}
        <a id="L704"></a>}
        <a id="L705"></a>for i := 0; i &lt; len(s[in]); i++ {
            <a id="L706"></a>st := s[in][i];
            <a id="L707"></a>switch s[in][i].inst.kind() {
            <a id="L708"></a>case _BOT:
                <a id="L709"></a>if pos == 0 {
                    <a id="L710"></a>s[in] = addState(s[in], st.inst.next(), st.match)
                <a id="L711"></a>}
            <a id="L712"></a>case _EOT:
                <a id="L713"></a>if pos == end {
                    <a id="L714"></a>s[in] = addState(s[in], st.inst.next(), st.match)
                <a id="L715"></a>}
            <a id="L716"></a>case _CHAR:
                <a id="L717"></a>if c == st.inst.(*_Char).char {
                    <a id="L718"></a>s[out] = addState(s[out], st.inst.next(), st.match)
                <a id="L719"></a>}
            <a id="L720"></a>case _CHARCLASS:
                <a id="L721"></a>if st.inst.(*_CharClass).matches(c) {
                    <a id="L722"></a>s[out] = addState(s[out], st.inst.next(), st.match)
                <a id="L723"></a>}
            <a id="L724"></a>case _ANY:
                <a id="L725"></a>if c != endOfFile {
                    <a id="L726"></a>s[out] = addState(s[out], st.inst.next(), st.match)
                <a id="L727"></a>}
            <a id="L728"></a>case _NOTNL:
                <a id="L729"></a>if c != endOfFile &amp;&amp; c != &#39;\n&#39; {
                    <a id="L730"></a>s[out] = addState(s[out], st.inst.next(), st.match)
                <a id="L731"></a>}
            <a id="L732"></a>case _BRA:
                <a id="L733"></a>n := st.inst.(*_Bra).n;
                <a id="L734"></a>st.match[2*n] = pos;
                <a id="L735"></a>s[in] = addState(s[in], st.inst.next(), st.match);
            <a id="L736"></a>case _EBRA:
                <a id="L737"></a>n := st.inst.(*_Ebra).n;
                <a id="L738"></a>st.match[2*n+1] = pos;
                <a id="L739"></a>s[in] = addState(s[in], st.inst.next(), st.match);
            <a id="L740"></a>case _ALT:
                <a id="L741"></a>s[in] = addState(s[in], st.inst.(*_Alt).left, st.match);
                <a id="L742"></a><span class="comment">// give other branch a copy of this match vector</span>
                <a id="L743"></a>s1 := make([]int, 2*(re.nbra+1));
                <a id="L744"></a>for i := 0; i &lt; len(s1); i++ {
                    <a id="L745"></a>s1[i] = st.match[i]
                <a id="L746"></a>}
                <a id="L747"></a>s[in] = addState(s[in], st.inst.next(), s1);
            <a id="L748"></a>case _END:
                <a id="L749"></a><span class="comment">// choose leftmost longest</span>
                <a id="L750"></a>if !found || <span class="comment">// first</span>
                    <a id="L751"></a>st.match[0] &lt; final.match[0] || <span class="comment">// leftmost</span>
                    <a id="L752"></a>(st.match[0] == final.match[0] &amp;&amp; pos &gt; final.match[1]) { <span class="comment">// longest</span>
                    <a id="L753"></a>final = st;
                    <a id="L754"></a>final.match[1] = pos;
                <a id="L755"></a>}
                <a id="L756"></a>found = true;
            <a id="L757"></a>default:
                <a id="L758"></a>st.inst.print();
                <a id="L759"></a>panic(&#34;unknown instruction in execute&#34;);
            <a id="L760"></a>}
        <a id="L761"></a>}
        <a id="L762"></a>pos += charwidth;
    <a id="L763"></a>}
    <a id="L764"></a>return final.match;
<a id="L765"></a>}


<a id="L768"></a><span class="comment">// ExecuteString matches the Regexp against the string s.</span>
<a id="L769"></a><span class="comment">// The return value is an array of integers, in pairs, identifying the positions of</span>
<a id="L770"></a><span class="comment">// substrings matched by the expression.</span>
<a id="L771"></a><span class="comment">//    s[a[0]:a[1]] is the substring matched by the entire expression.</span>
<a id="L772"></a><span class="comment">//    s[a[2*i]:a[2*i+1]] for i &gt; 0 is the substring matched by the ith parenthesized subexpression.</span>
<a id="L773"></a><span class="comment">// A negative value means the subexpression did not match any element of the string.</span>
<a id="L774"></a><span class="comment">// An empty array means &#34;no match&#34;.</span>
<a id="L775"></a>func (re *Regexp) ExecuteString(s string) (a []int) {
    <a id="L776"></a>return re.doExecute(s, nil, 0)
<a id="L777"></a>}


<a id="L780"></a><span class="comment">// Execute matches the Regexp against the byte slice b.</span>
<a id="L781"></a><span class="comment">// The return value is an array of integers, in pairs, identifying the positions of</span>
<a id="L782"></a><span class="comment">// subslices matched by the expression.</span>
<a id="L783"></a><span class="comment">//    b[a[0]:a[1]] is the subslice matched by the entire expression.</span>
<a id="L784"></a><span class="comment">//    b[a[2*i]:a[2*i+1]] for i &gt; 0 is the subslice matched by the ith parenthesized subexpression.</span>
<a id="L785"></a><span class="comment">// A negative value means the subexpression did not match any element of the slice.</span>
<a id="L786"></a><span class="comment">// An empty array means &#34;no match&#34;.</span>
<a id="L787"></a>func (re *Regexp) Execute(b []byte) (a []int) { return re.doExecute(&#34;&#34;, b, 0) }


<a id="L790"></a><span class="comment">// MatchString returns whether the Regexp matches the string s.</span>
<a id="L791"></a><span class="comment">// The return value is a boolean: true for match, false for no match.</span>
<a id="L792"></a>func (re *Regexp) MatchString(s string) bool { return len(re.doExecute(s, nil, 0)) &gt; 0 }


<a id="L795"></a><span class="comment">// Match returns whether the Regexp matches the byte slice b.</span>
<a id="L796"></a><span class="comment">// The return value is a boolean: true for match, false for no match.</span>
<a id="L797"></a>func (re *Regexp) Match(b []byte) bool { return len(re.doExecute(&#34;&#34;, b, 0)) &gt; 0 }


<a id="L800"></a><span class="comment">// MatchStrings matches the Regexp against the string s.</span>
<a id="L801"></a><span class="comment">// The return value is an array of strings matched by the expression.</span>
<a id="L802"></a><span class="comment">//    a[0] is the substring matched by the entire expression.</span>
<a id="L803"></a><span class="comment">//    a[i] for i &gt; 0 is the substring matched by the ith parenthesized subexpression.</span>
<a id="L804"></a><span class="comment">// An empty array means ``no match&#39;&#39;.</span>
<a id="L805"></a>func (re *Regexp) MatchStrings(s string) (a []string) {
    <a id="L806"></a>r := re.doExecute(s, nil, 0);
    <a id="L807"></a>if r == nil {
        <a id="L808"></a>return nil
    <a id="L809"></a>}
    <a id="L810"></a>a = make([]string, len(r)/2);
    <a id="L811"></a>for i := 0; i &lt; len(r); i += 2 {
        <a id="L812"></a>if r[i] != -1 { <span class="comment">// -1 means no match for this subexpression</span>
            <a id="L813"></a>a[i/2] = s[r[i]:r[i+1]]
        <a id="L814"></a>}
    <a id="L815"></a>}
    <a id="L816"></a>return;
<a id="L817"></a>}

<a id="L819"></a><span class="comment">// MatchSlices matches the Regexp against the byte slice b.</span>
<a id="L820"></a><span class="comment">// The return value is an array of subslices matched by the expression.</span>
<a id="L821"></a><span class="comment">//    a[0] is the subslice matched by the entire expression.</span>
<a id="L822"></a><span class="comment">//    a[i] for i &gt; 0 is the subslice matched by the ith parenthesized subexpression.</span>
<a id="L823"></a><span class="comment">// An empty array means ``no match&#39;&#39;.</span>
<a id="L824"></a>func (re *Regexp) MatchSlices(b []byte) (a [][]byte) {
    <a id="L825"></a>r := re.doExecute(&#34;&#34;, b, 0);
    <a id="L826"></a>if r == nil {
        <a id="L827"></a>return nil
    <a id="L828"></a>}
    <a id="L829"></a>a = make([][]byte, len(r)/2);
    <a id="L830"></a>for i := 0; i &lt; len(r); i += 2 {
        <a id="L831"></a>if r[i] != -1 { <span class="comment">// -1 means no match for this subexpression</span>
            <a id="L832"></a>a[i/2] = b[r[i]:r[i+1]]
        <a id="L833"></a>}
    <a id="L834"></a>}
    <a id="L835"></a>return;
<a id="L836"></a>}

<a id="L838"></a><span class="comment">// MatchString checks whether a textual regular expression</span>
<a id="L839"></a><span class="comment">// matches a string.  More complicated queries need</span>
<a id="L840"></a><span class="comment">// to use Compile and the full Regexp interface.</span>
<a id="L841"></a>func MatchString(pattern string, s string) (matched bool, error os.Error) {
    <a id="L842"></a>re, err := Compile(pattern);
    <a id="L843"></a>if err != nil {
        <a id="L844"></a>return false, err
    <a id="L845"></a>}
    <a id="L846"></a>return re.MatchString(s), nil;
<a id="L847"></a>}

<a id="L849"></a><span class="comment">// Match checks whether a textual regular expression</span>
<a id="L850"></a><span class="comment">// matches a byte slice.  More complicated queries need</span>
<a id="L851"></a><span class="comment">// to use Compile and the full Regexp interface.</span>
<a id="L852"></a>func Match(pattern string, b []byte) (matched bool, error os.Error) {
    <a id="L853"></a>re, err := Compile(pattern);
    <a id="L854"></a>if err != nil {
        <a id="L855"></a>return false, err
    <a id="L856"></a>}
    <a id="L857"></a>return re.Match(b), nil;
<a id="L858"></a>}

<a id="L860"></a><span class="comment">// ReplaceAllString returns a copy of src in which all matches for the Regexp</span>
<a id="L861"></a><span class="comment">// have been replaced by repl.  No support is provided for expressions</span>
<a id="L862"></a><span class="comment">// (e.g. \1 or $1) in the replacement string.</span>
<a id="L863"></a>func (re *Regexp) ReplaceAllString(src, repl string) string {
    <a id="L864"></a>lastMatchEnd := 0; <span class="comment">// end position of the most recent match</span>
    <a id="L865"></a>searchPos := 0;    <span class="comment">// position where we next look for a match</span>
    <a id="L866"></a>buf := new(bytes.Buffer);
    <a id="L867"></a>for searchPos &lt;= len(src) {
        <a id="L868"></a>a := re.doExecute(src, nil, searchPos);
        <a id="L869"></a>if len(a) == 0 {
            <a id="L870"></a>break <span class="comment">// no more matches</span>
        <a id="L871"></a>}

        <a id="L873"></a><span class="comment">// Copy the unmatched characters before this match.</span>
        <a id="L874"></a>io.WriteString(buf, src[lastMatchEnd:a[0]]);

        <a id="L876"></a><span class="comment">// Now insert a copy of the replacement string, but not for a</span>
        <a id="L877"></a><span class="comment">// match of the empty string immediately after another match.</span>
        <a id="L878"></a><span class="comment">// (Otherwise, we get double replacement for patterns that</span>
        <a id="L879"></a><span class="comment">// match both empty and nonempty strings.)</span>
        <a id="L880"></a>if a[1] &gt; lastMatchEnd || a[0] == 0 {
            <a id="L881"></a>io.WriteString(buf, repl)
        <a id="L882"></a>}
        <a id="L883"></a>lastMatchEnd = a[1];

        <a id="L885"></a><span class="comment">// Advance past this match; always advance at least one character.</span>
        <a id="L886"></a>_, width := utf8.DecodeRuneInString(src[searchPos:len(src)]);
        <a id="L887"></a>if searchPos+width &gt; a[1] {
            <a id="L888"></a>searchPos += width
        <a id="L889"></a>} else if searchPos+1 &gt; a[1] {
            <a id="L890"></a><span class="comment">// This clause is only needed at the end of the input</span>
            <a id="L891"></a><span class="comment">// string.  In that case, DecodeRuneInString returns width=0.</span>
            <a id="L892"></a>searchPos++
        <a id="L893"></a>} else {
            <a id="L894"></a>searchPos = a[1]
        <a id="L895"></a>}
    <a id="L896"></a>}

    <a id="L898"></a><span class="comment">// Copy the unmatched characters after the last match.</span>
    <a id="L899"></a>io.WriteString(buf, src[lastMatchEnd:len(src)]);

    <a id="L901"></a>return buf.String();
<a id="L902"></a>}

<a id="L904"></a><span class="comment">// ReplaceAll returns a copy of src in which all matches for the Regexp</span>
<a id="L905"></a><span class="comment">// have been replaced by repl.  No support is provided for expressions</span>
<a id="L906"></a><span class="comment">// (e.g. \1 or $1) in the replacement text.</span>
<a id="L907"></a>func (re *Regexp) ReplaceAll(src, repl []byte) []byte {
    <a id="L908"></a>lastMatchEnd := 0; <span class="comment">// end position of the most recent match</span>
    <a id="L909"></a>searchPos := 0;    <span class="comment">// position where we next look for a match</span>
    <a id="L910"></a>buf := new(bytes.Buffer);
    <a id="L911"></a>for searchPos &lt;= len(src) {
        <a id="L912"></a>a := re.doExecute(&#34;&#34;, src, searchPos);
        <a id="L913"></a>if len(a) == 0 {
            <a id="L914"></a>break <span class="comment">// no more matches</span>
        <a id="L915"></a>}

        <a id="L917"></a><span class="comment">// Copy the unmatched characters before this match.</span>
        <a id="L918"></a>buf.Write(src[lastMatchEnd:a[0]]);

        <a id="L920"></a><span class="comment">// Now insert a copy of the replacement string, but not for a</span>
        <a id="L921"></a><span class="comment">// match of the empty string immediately after another match.</span>
        <a id="L922"></a><span class="comment">// (Otherwise, we get double replacement for patterns that</span>
        <a id="L923"></a><span class="comment">// match both empty and nonempty strings.)</span>
        <a id="L924"></a>if a[1] &gt; lastMatchEnd || a[0] == 0 {
            <a id="L925"></a>buf.Write(repl)
        <a id="L926"></a>}
        <a id="L927"></a>lastMatchEnd = a[1];

        <a id="L929"></a><span class="comment">// Advance past this match; always advance at least one character.</span>
        <a id="L930"></a>_, width := utf8.DecodeRune(src[searchPos:len(src)]);
        <a id="L931"></a>if searchPos+width &gt; a[1] {
            <a id="L932"></a>searchPos += width
        <a id="L933"></a>} else if searchPos+1 &gt; a[1] {
            <a id="L934"></a><span class="comment">// This clause is only needed at the end of the input</span>
            <a id="L935"></a><span class="comment">// string.  In that case, DecodeRuneInString returns width=0.</span>
            <a id="L936"></a>searchPos++
        <a id="L937"></a>} else {
            <a id="L938"></a>searchPos = a[1]
        <a id="L939"></a>}
    <a id="L940"></a>}

    <a id="L942"></a><span class="comment">// Copy the unmatched characters after the last match.</span>
    <a id="L943"></a>buf.Write(src[lastMatchEnd:len(src)]);

    <a id="L945"></a>return buf.Bytes();
<a id="L946"></a>}

<a id="L948"></a><span class="comment">// QuoteMeta returns a string that quotes all regular expression metacharacters</span>
<a id="L949"></a><span class="comment">// inside the argument text; the returned string is a regular expression matching</span>
<a id="L950"></a><span class="comment">// the literal text.  For example, QuoteMeta(`[foo]`) returns `\[foo\]`.</span>
<a id="L951"></a>func QuoteMeta(s string) string {
    <a id="L952"></a>b := make([]byte, 2*len(s));

    <a id="L954"></a><span class="comment">// A byte loop is correct because all metacharacters are ASCII.</span>
    <a id="L955"></a>j := 0;
    <a id="L956"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L957"></a>if special(int(s[i])) {
            <a id="L958"></a>b[j] = &#39;\\&#39;;
            <a id="L959"></a>j++;
        <a id="L960"></a>}
        <a id="L961"></a>b[j] = s[i];
        <a id="L962"></a>j++;
    <a id="L963"></a>}
    <a id="L964"></a>return string(b[0:j]);
<a id="L965"></a>}

<a id="L967"></a><span class="comment">// Find matches in slice b if b is non-nil, otherwise find matches in string s.</span>
<a id="L968"></a>func (re *Regexp) allMatches(s string, b []byte, n int, deliver func(int, int)) {
    <a id="L969"></a>var end int;
    <a id="L970"></a>if b == nil {
        <a id="L971"></a>end = len(s)
    <a id="L972"></a>} else {
        <a id="L973"></a>end = len(b)
    <a id="L974"></a>}

    <a id="L976"></a>for pos, i, prevMatchEnd := 0, 0, -1; i &lt; n &amp;&amp; pos &lt;= end; {
        <a id="L977"></a>matches := re.doExecute(s, b, pos);
        <a id="L978"></a>if len(matches) == 0 {
            <a id="L979"></a>break
        <a id="L980"></a>}

        <a id="L982"></a>accept := true;
        <a id="L983"></a>if matches[1] == pos {
            <a id="L984"></a><span class="comment">// We&#39;ve found an empty match.</span>
            <a id="L985"></a>if matches[0] == prevMatchEnd {
                <a id="L986"></a><span class="comment">// We don&#39;t allow an empty match right</span>
                <a id="L987"></a><span class="comment">// after a previous match, so ignore it.</span>
                <a id="L988"></a>accept = false
            <a id="L989"></a>}
            <a id="L990"></a>var width int;
            <a id="L991"></a>if b == nil {
                <a id="L992"></a>_, width = utf8.DecodeRuneInString(s[pos:end])
            <a id="L993"></a>} else {
                <a id="L994"></a>_, width = utf8.DecodeRune(b[pos:end])
            <a id="L995"></a>}
            <a id="L996"></a>if width &gt; 0 {
                <a id="L997"></a>pos += width
            <a id="L998"></a>} else {
                <a id="L999"></a>pos = end + 1
            <a id="L1000"></a>}
        <a id="L1001"></a>} else {
            <a id="L1002"></a>pos = matches[1]
        <a id="L1003"></a>}
        <a id="L1004"></a>prevMatchEnd = matches[1];

        <a id="L1006"></a>if accept {
            <a id="L1007"></a>deliver(matches[0], matches[1]);
            <a id="L1008"></a>i++;
        <a id="L1009"></a>}
    <a id="L1010"></a>}
<a id="L1011"></a>}

<a id="L1013"></a><span class="comment">// AllMatches slices the byte slice b into substrings that are successive</span>
<a id="L1014"></a><span class="comment">// matches of the Regexp within b. If n &gt; 0, the function returns at most n</span>
<a id="L1015"></a><span class="comment">// matches. Text that does not match the expression will be skipped. Empty</span>
<a id="L1016"></a><span class="comment">// matches abutting a preceding match are ignored. The function returns a slice</span>
<a id="L1017"></a><span class="comment">// containing the matching substrings.</span>
<a id="L1018"></a>func (re *Regexp) AllMatches(b []byte, n int) [][]byte {
    <a id="L1019"></a>if n &lt;= 0 {
        <a id="L1020"></a>n = len(b) + 1
    <a id="L1021"></a>}
    <a id="L1022"></a>result := make([][]byte, n);
    <a id="L1023"></a>i := 0;
    <a id="L1024"></a>re.allMatches(&#34;&#34;, b, n, func(start, end int) {
        <a id="L1025"></a>result[i] = b[start:end];
        <a id="L1026"></a>i++;
    <a id="L1027"></a>});
    <a id="L1028"></a>return result[0:i];
<a id="L1029"></a>}

<a id="L1031"></a><span class="comment">// AllMatchesString slices the string s into substrings that are successive</span>
<a id="L1032"></a><span class="comment">// matches of the Regexp within s. If n &gt; 0, the function returns at most n</span>
<a id="L1033"></a><span class="comment">// matches. Text that does not match the expression will be skipped. Empty</span>
<a id="L1034"></a><span class="comment">// matches abutting a preceding match are ignored. The function returns a slice</span>
<a id="L1035"></a><span class="comment">// containing the matching substrings.</span>
<a id="L1036"></a>func (re *Regexp) AllMatchesString(s string, n int) []string {
    <a id="L1037"></a>if n &lt;= 0 {
        <a id="L1038"></a>n = len(s) + 1
    <a id="L1039"></a>}
    <a id="L1040"></a>result := make([]string, n);
    <a id="L1041"></a>i := 0;
    <a id="L1042"></a>re.allMatches(s, nil, n, func(start, end int) {
        <a id="L1043"></a>result[i] = s[start:end];
        <a id="L1044"></a>i++;
    <a id="L1045"></a>});
    <a id="L1046"></a>return result[0:i];
<a id="L1047"></a>}

<a id="L1049"></a><span class="comment">// AllMatchesIter slices the byte slice b into substrings that are successive</span>
<a id="L1050"></a><span class="comment">// matches of the Regexp within b. If n &gt; 0, the function returns at most n</span>
<a id="L1051"></a><span class="comment">// matches. Text that does not match the expression will be skipped. Empty</span>
<a id="L1052"></a><span class="comment">// matches abutting a preceding match are ignored. The function returns a</span>
<a id="L1053"></a><span class="comment">// channel that iterates over the matching substrings.</span>
<a id="L1054"></a>func (re *Regexp) AllMatchesIter(b []byte, n int) &lt;-chan []byte {
    <a id="L1055"></a>if n &lt;= 0 {
        <a id="L1056"></a>n = len(b) + 1
    <a id="L1057"></a>}
    <a id="L1058"></a>c := make(chan []byte, 10);
    <a id="L1059"></a>go func() {
        <a id="L1060"></a>re.allMatches(&#34;&#34;, b, n, func(start, end int) { c &lt;- b[start:end] });
        <a id="L1061"></a>close(c);
    <a id="L1062"></a>}();
    <a id="L1063"></a>return c;
<a id="L1064"></a>}

<a id="L1066"></a><span class="comment">// AllMatchesStringIter slices the string s into substrings that are successive</span>
<a id="L1067"></a><span class="comment">// matches of the Regexp within s. If n &gt; 0, the function returns at most n</span>
<a id="L1068"></a><span class="comment">// matches. Text that does not match the expression will be skipped. Empty</span>
<a id="L1069"></a><span class="comment">// matches abutting a preceding match are ignored. The function returns a</span>
<a id="L1070"></a><span class="comment">// channel that iterates over the matching substrings.</span>
<a id="L1071"></a>func (re *Regexp) AllMatchesStringIter(s string, n int) &lt;-chan string {
    <a id="L1072"></a>if n &lt;= 0 {
        <a id="L1073"></a>n = len(s) + 1
    <a id="L1074"></a>}
    <a id="L1075"></a>c := make(chan string, 10);
    <a id="L1076"></a>go func() {
        <a id="L1077"></a>re.allMatches(s, nil, n, func(start, end int) { c &lt;- s[start:end] });
        <a id="L1078"></a>close(c);
    <a id="L1079"></a>}();
    <a id="L1080"></a>return c;
<a id="L1081"></a>}
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
