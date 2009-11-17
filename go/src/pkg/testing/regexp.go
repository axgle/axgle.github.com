<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/testing/regexp.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/testing/regexp.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The testing package implements a simple regular expression library.</span>
<a id="L6"></a><span class="comment">// It is a reduced version of the regular expression package suitable</span>
<a id="L7"></a><span class="comment">// for use in tests; it avoids many dependencies.</span>
<a id="L8"></a><span class="comment">//</span>
<a id="L9"></a><span class="comment">// The syntax of the regular expressions accepted is:</span>
<a id="L10"></a><span class="comment">//</span>
<a id="L11"></a><span class="comment">//	regexp:</span>
<a id="L12"></a><span class="comment">//		concatenation { &#39;|&#39; concatenation }</span>
<a id="L13"></a><span class="comment">//	concatenation:</span>
<a id="L14"></a><span class="comment">//		{ closure }</span>
<a id="L15"></a><span class="comment">//	closure:</span>
<a id="L16"></a><span class="comment">//		term [ &#39;*&#39; | &#39;+&#39; | &#39;?&#39; ]</span>
<a id="L17"></a><span class="comment">//	term:</span>
<a id="L18"></a><span class="comment">//		&#39;^&#39;</span>
<a id="L19"></a><span class="comment">//		&#39;$&#39;</span>
<a id="L20"></a><span class="comment">//		&#39;.&#39;</span>
<a id="L21"></a><span class="comment">//		character</span>
<a id="L22"></a><span class="comment">//		&#39;[&#39; [ &#39;^&#39; ] character-ranges &#39;]&#39;</span>
<a id="L23"></a><span class="comment">//		&#39;(&#39; regexp &#39;)&#39;</span>
<a id="L24"></a><span class="comment">//</span>

<a id="L26"></a>package testing

<a id="L28"></a>import (
    <a id="L29"></a>&#34;utf8&#34;;
<a id="L30"></a>)

<a id="L32"></a>var debug = false

<a id="L34"></a><span class="comment">// Error codes returned by failures to parse an expression.</span>
<a id="L35"></a>var (
    <a id="L36"></a>ErrInternal            = &#34;internal error&#34;;
    <a id="L37"></a>ErrUnmatchedLpar       = &#34;unmatched &#39;&#39;&#34;;
    <a id="L38"></a>ErrUnmatchedRpar       = &#34;unmatched &#39;&#39;&#34;;
    <a id="L39"></a>ErrUnmatchedLbkt       = &#34;unmatched &#39;[&#39;&#34;;
    <a id="L40"></a>ErrUnmatchedRbkt       = &#34;unmatched &#39;]&#39;&#34;;
    <a id="L41"></a>ErrBadRange            = &#34;bad range in character class&#34;;
    <a id="L42"></a>ErrExtraneousBackslash = &#34;extraneous backslash&#34;;
    <a id="L43"></a>ErrBadClosure          = &#34;repeated closure **, ++, etc.&#34;;
    <a id="L44"></a>ErrBareClosure         = &#34;closure applies to nothing&#34;;
    <a id="L45"></a>ErrBadBackslash        = &#34;illegal backslash escape&#34;;
<a id="L46"></a>)

<a id="L48"></a><span class="comment">// An instruction executed by the NFA</span>
<a id="L49"></a>type instr interface {
    <a id="L50"></a>kind() int;   <span class="comment">// the type of this instruction: _CHAR, _ANY, etc.</span>
    <a id="L51"></a>next() instr; <span class="comment">// the instruction to execute after this one</span>
    <a id="L52"></a>setNext(i instr);
    <a id="L53"></a>index() int;
    <a id="L54"></a>setIndex(i int);
    <a id="L55"></a>print();
<a id="L56"></a>}

<a id="L58"></a><span class="comment">// Fields and methods common to all instructions</span>
<a id="L59"></a>type common struct {
    <a id="L60"></a>_next  instr;
    <a id="L61"></a>_index int;
<a id="L62"></a>}

<a id="L64"></a>func (c *common) next() instr     { return c._next }
<a id="L65"></a>func (c *common) setNext(i instr) { c._next = i }
<a id="L66"></a>func (c *common) index() int      { return c._index }
<a id="L67"></a>func (c *common) setIndex(i int)  { c._index = i }

<a id="L69"></a><span class="comment">// The representation of a compiled regular expression.</span>
<a id="L70"></a><span class="comment">// The public interface is entirely through methods.</span>
<a id="L71"></a>type Regexp struct {
    <a id="L72"></a>expr  string; <span class="comment">// the original expression</span>
    <a id="L73"></a>inst  []instr;
    <a id="L74"></a>start instr;
    <a id="L75"></a>nbra  int; <span class="comment">// number of brackets in expression, for subexpressions</span>
<a id="L76"></a>}

<a id="L78"></a>const (
    <a id="L79"></a>_START =   <span class="comment">// beginning of program</span>
    <a id="L80"></a>iota;
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
    <a id="L147"></a><span class="comment">// stored pairwise: [a-z] is (a,z); x is (x,x):</span>
    <a id="L148"></a>ranges []int;
<a id="L149"></a>}

<a id="L151"></a>func (cclass *_CharClass) kind() int { return _CHARCLASS }

<a id="L153"></a>func (cclass *_CharClass) print() {
    <a id="L154"></a>print(&#34;charclass&#34;);
    <a id="L155"></a>if cclass.negate {
        <a id="L156"></a>print(&#34; (negated)&#34;)
    <a id="L157"></a>}
    <a id="L158"></a>for i := 0; i &lt; len(cclass.ranges); i += 2 {
        <a id="L159"></a>l := cclass.ranges[i];
        <a id="L160"></a>r := cclass.ranges[i+1];
        <a id="L161"></a>if l == r {
            <a id="L162"></a>print(&#34; [&#34;, string(l), &#34;]&#34;)
        <a id="L163"></a>} else {
            <a id="L164"></a>print(&#34; [&#34;, string(l), &#34;-&#34;, string(r), &#34;]&#34;)
        <a id="L165"></a>}
    <a id="L166"></a>}
<a id="L167"></a>}

<a id="L169"></a>func (cclass *_CharClass) addRange(a, b int) {
    <a id="L170"></a><span class="comment">// range is a through b inclusive</span>
    <a id="L171"></a>n := len(cclass.ranges);
    <a id="L172"></a>if n &gt;= cap(cclass.ranges) {
        <a id="L173"></a>nr := make([]int, n, 2*n);
        <a id="L174"></a>for i, j := range nr {
            <a id="L175"></a>nr[i] = j
        <a id="L176"></a>}
        <a id="L177"></a>cclass.ranges = nr;
    <a id="L178"></a>}
    <a id="L179"></a>cclass.ranges = cclass.ranges[0 : n+2];
    <a id="L180"></a>cclass.ranges[n] = a;
    <a id="L181"></a>n++;
    <a id="L182"></a>cclass.ranges[n] = b;
    <a id="L183"></a>n++;
<a id="L184"></a>}

<a id="L186"></a>func (cclass *_CharClass) matches(c int) bool {
    <a id="L187"></a>for i := 0; i &lt; len(cclass.ranges); i = i + 2 {
        <a id="L188"></a>min := cclass.ranges[i];
        <a id="L189"></a>max := cclass.ranges[i+1];
        <a id="L190"></a>if min &lt;= c &amp;&amp; c &lt;= max {
            <a id="L191"></a>return !cclass.negate
        <a id="L192"></a>}
    <a id="L193"></a>}
    <a id="L194"></a>return cclass.negate;
<a id="L195"></a>}

<a id="L197"></a>func newCharClass() *_CharClass {
    <a id="L198"></a>c := new(_CharClass);
    <a id="L199"></a>c.ranges = make([]int, 0, 20);
    <a id="L200"></a>return c;
<a id="L201"></a>}

<a id="L203"></a><span class="comment">// --- ANY any character</span>
<a id="L204"></a>type _Any struct {
    <a id="L205"></a>common;
<a id="L206"></a>}

<a id="L208"></a>func (any *_Any) kind() int { return _ANY }
<a id="L209"></a>func (any *_Any) print()    { print(&#34;any&#34;) }

<a id="L211"></a><span class="comment">// --- NOTNL any character but newline</span>
<a id="L212"></a>type _NotNl struct {
    <a id="L213"></a>common;
<a id="L214"></a>}

<a id="L216"></a>func (notnl *_NotNl) kind() int { return _NOTNL }
<a id="L217"></a>func (notnl *_NotNl) print()    { print(&#34;notnl&#34;) }

<a id="L219"></a><span class="comment">// --- BRA parenthesized expression</span>
<a id="L220"></a>type _Bra struct {
    <a id="L221"></a>common;
    <a id="L222"></a>n   int; <span class="comment">// subexpression number</span>
<a id="L223"></a>}

<a id="L225"></a>func (bra *_Bra) kind() int { return _BRA }
<a id="L226"></a>func (bra *_Bra) print()    { print(&#34;bra&#34;, bra.n) }

<a id="L228"></a><span class="comment">// --- EBRA end of parenthesized expression</span>
<a id="L229"></a>type _Ebra struct {
    <a id="L230"></a>common;
    <a id="L231"></a>n   int; <span class="comment">// subexpression number</span>
<a id="L232"></a>}

<a id="L234"></a>func (ebra *_Ebra) kind() int { return _EBRA }
<a id="L235"></a>func (ebra *_Ebra) print()    { print(&#34;ebra &#34;, ebra.n) }

<a id="L237"></a><span class="comment">// --- ALT alternation</span>
<a id="L238"></a>type _Alt struct {
    <a id="L239"></a>common;
    <a id="L240"></a>left instr; <span class="comment">// other branch</span>
<a id="L241"></a>}

<a id="L243"></a>func (alt *_Alt) kind() int { return _ALT }
<a id="L244"></a>func (alt *_Alt) print()    { print(&#34;alt(&#34;, alt.left.index(), &#34;)&#34;) }

<a id="L246"></a><span class="comment">// --- NOP no operation</span>
<a id="L247"></a>type _Nop struct {
    <a id="L248"></a>common;
<a id="L249"></a>}

<a id="L251"></a>func (nop *_Nop) kind() int { return _NOP }
<a id="L252"></a>func (nop *_Nop) print()    { print(&#34;nop&#34;) }

<a id="L254"></a>func (re *Regexp) add(i instr) instr {
    <a id="L255"></a>n := len(re.inst);
    <a id="L256"></a>i.setIndex(len(re.inst));
    <a id="L257"></a>if n &gt;= cap(re.inst) {
        <a id="L258"></a>ni := make([]instr, n, 2*n);
        <a id="L259"></a>for i, j := range re.inst {
            <a id="L260"></a>ni[i] = j
        <a id="L261"></a>}
        <a id="L262"></a>re.inst = ni;
    <a id="L263"></a>}
    <a id="L264"></a>re.inst = re.inst[0 : n+1];
    <a id="L265"></a>re.inst[n] = i;
    <a id="L266"></a>return i;
<a id="L267"></a>}

<a id="L269"></a>type parser struct {
    <a id="L270"></a>re    *Regexp;
    <a id="L271"></a>error string;
    <a id="L272"></a>nlpar int; <span class="comment">// number of unclosed lpars</span>
    <a id="L273"></a>pos   int;
    <a id="L274"></a>ch    int;
<a id="L275"></a>}

<a id="L277"></a>const endOfFile = -1

<a id="L279"></a>func (p *parser) c() int { return p.ch }

<a id="L281"></a>func (p *parser) nextc() int {
    <a id="L282"></a>if p.pos &gt;= len(p.re.expr) {
        <a id="L283"></a>p.ch = endOfFile
    <a id="L284"></a>} else {
        <a id="L285"></a>c, w := utf8.DecodeRuneInString(p.re.expr[p.pos:len(p.re.expr)]);
        <a id="L286"></a>p.ch = c;
        <a id="L287"></a>p.pos += w;
    <a id="L288"></a>}
    <a id="L289"></a>return p.ch;
<a id="L290"></a>}

<a id="L292"></a>func newParser(re *Regexp) *parser {
    <a id="L293"></a>p := new(parser);
    <a id="L294"></a>p.re = re;
    <a id="L295"></a>p.nextc(); <span class="comment">// load p.ch</span>
    <a id="L296"></a>return p;
<a id="L297"></a>}

<a id="L299"></a>func special(c int) bool {
    <a id="L300"></a>s := `\.+*?()|[]^$`;
    <a id="L301"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L302"></a>if c == int(s[i]) {
            <a id="L303"></a>return true
        <a id="L304"></a>}
    <a id="L305"></a>}
    <a id="L306"></a>return false;
<a id="L307"></a>}

<a id="L309"></a>func specialcclass(c int) bool {
    <a id="L310"></a>s := `\-[]`;
    <a id="L311"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L312"></a>if c == int(s[i]) {
            <a id="L313"></a>return true
        <a id="L314"></a>}
    <a id="L315"></a>}
    <a id="L316"></a>return false;
<a id="L317"></a>}

<a id="L319"></a>func (p *parser) charClass() instr {
    <a id="L320"></a>cc := newCharClass();
    <a id="L321"></a>if p.c() == &#39;^&#39; {
        <a id="L322"></a>cc.negate = true;
        <a id="L323"></a>p.nextc();
    <a id="L324"></a>}
    <a id="L325"></a>left := -1;
    <a id="L326"></a>for {
        <a id="L327"></a>switch c := p.c(); c {
        <a id="L328"></a>case &#39;]&#39;, endOfFile:
            <a id="L329"></a>if left &gt;= 0 {
                <a id="L330"></a>p.error = ErrBadRange;
                <a id="L331"></a>return nil;
            <a id="L332"></a>}
            <a id="L333"></a><span class="comment">// Is it [^\n]?</span>
            <a id="L334"></a>if cc.negate &amp;&amp; len(cc.ranges) == 2 &amp;&amp;
                <a id="L335"></a>cc.ranges[0] == &#39;\n&#39; &amp;&amp; cc.ranges[1] == &#39;\n&#39; {
                <a id="L336"></a>nl := new(_NotNl);
                <a id="L337"></a>p.re.add(nl);
                <a id="L338"></a>return nl;
            <a id="L339"></a>}
            <a id="L340"></a>p.re.add(cc);
            <a id="L341"></a>return cc;
        <a id="L342"></a>case &#39;-&#39;: <span class="comment">// do this before backslash processing</span>
            <a id="L343"></a>p.error = ErrBadRange;
            <a id="L344"></a>return nil;
        <a id="L345"></a>case &#39;\\&#39;:
            <a id="L346"></a>c = p.nextc();
            <a id="L347"></a>switch {
            <a id="L348"></a>case c == endOfFile:
                <a id="L349"></a>p.error = ErrExtraneousBackslash;
                <a id="L350"></a>return nil;
            <a id="L351"></a>case c == &#39;n&#39;:
                <a id="L352"></a>c = &#39;\n&#39;
            <a id="L353"></a>case specialcclass(c):
            <a id="L354"></a><span class="comment">// c is as delivered</span>
            <a id="L355"></a>default:
                <a id="L356"></a>p.error = ErrBadBackslash;
                <a id="L357"></a>return nil;
            <a id="L358"></a>}
            <a id="L359"></a>fallthrough;
        <a id="L360"></a>default:
            <a id="L361"></a>p.nextc();
            <a id="L362"></a>switch {
            <a id="L363"></a>case left &lt; 0: <span class="comment">// first of pair</span>
                <a id="L364"></a>if p.c() == &#39;-&#39; { <span class="comment">// range</span>
                    <a id="L365"></a>p.nextc();
                    <a id="L366"></a>left = c;
                <a id="L367"></a>} else { <span class="comment">// single char</span>
                    <a id="L368"></a>cc.addRange(c, c)
                <a id="L369"></a>}
            <a id="L370"></a>case left &lt;= c: <span class="comment">// second of pair</span>
                <a id="L371"></a>cc.addRange(left, c);
                <a id="L372"></a>left = -1;
            <a id="L373"></a>default:
                <a id="L374"></a>p.error = ErrBadRange;
                <a id="L375"></a>return nil;
            <a id="L376"></a>}
        <a id="L377"></a>}
    <a id="L378"></a>}
    <a id="L379"></a>return nil;
<a id="L380"></a>}

<a id="L382"></a>func (p *parser) term() (start, end instr) {
    <a id="L383"></a><span class="comment">// term() is the leaf of the recursion, so it&#39;s sufficient to pick off the</span>
    <a id="L384"></a><span class="comment">// error state here for early exit.</span>
    <a id="L385"></a><span class="comment">// The other functions (closure(), concatenation() etc.) assume</span>
    <a id="L386"></a><span class="comment">// it&#39;s safe to recur to here.</span>
    <a id="L387"></a>if p.error != &#34;&#34; {
        <a id="L388"></a>return
    <a id="L389"></a>}
    <a id="L390"></a>switch c := p.c(); c {
    <a id="L391"></a>case &#39;|&#39;, endOfFile:
        <a id="L392"></a>return nil, nil
    <a id="L393"></a>case &#39;*&#39;, &#39;+&#39;:
        <a id="L394"></a>p.error = ErrBareClosure;
        <a id="L395"></a>return;
    <a id="L396"></a>case &#39;)&#39;:
        <a id="L397"></a>if p.nlpar == 0 {
            <a id="L398"></a>p.error = ErrUnmatchedRpar;
            <a id="L399"></a>return;
        <a id="L400"></a>}
        <a id="L401"></a>return nil, nil;
    <a id="L402"></a>case &#39;]&#39;:
        <a id="L403"></a>p.error = ErrUnmatchedRbkt;
        <a id="L404"></a>return;
    <a id="L405"></a>case &#39;^&#39;:
        <a id="L406"></a>p.nextc();
        <a id="L407"></a>start = p.re.add(new(_Bot));
        <a id="L408"></a>return start, start;
    <a id="L409"></a>case &#39;$&#39;:
        <a id="L410"></a>p.nextc();
        <a id="L411"></a>start = p.re.add(new(_Eot));
        <a id="L412"></a>return start, start;
    <a id="L413"></a>case &#39;.&#39;:
        <a id="L414"></a>p.nextc();
        <a id="L415"></a>start = p.re.add(new(_Any));
        <a id="L416"></a>return start, start;
    <a id="L417"></a>case &#39;[&#39;:
        <a id="L418"></a>p.nextc();
        <a id="L419"></a>start = p.charClass();
        <a id="L420"></a>if p.error != &#34;&#34; {
            <a id="L421"></a>return
        <a id="L422"></a>}
        <a id="L423"></a>if p.c() != &#39;]&#39; {
            <a id="L424"></a>p.error = ErrUnmatchedLbkt;
            <a id="L425"></a>return;
        <a id="L426"></a>}
        <a id="L427"></a>p.nextc();
        <a id="L428"></a>return start, start;
    <a id="L429"></a>case &#39;(&#39;:
        <a id="L430"></a>p.nextc();
        <a id="L431"></a>p.nlpar++;
        <a id="L432"></a>p.re.nbra++; <span class="comment">// increment first so first subexpr is \1</span>
        <a id="L433"></a>nbra := p.re.nbra;
        <a id="L434"></a>start, end = p.regexp();
        <a id="L435"></a>if p.c() != &#39;)&#39; {
            <a id="L436"></a>p.error = ErrUnmatchedLpar;
            <a id="L437"></a>return;
        <a id="L438"></a>}
        <a id="L439"></a>p.nlpar--;
        <a id="L440"></a>p.nextc();
        <a id="L441"></a>bra := new(_Bra);
        <a id="L442"></a>p.re.add(bra);
        <a id="L443"></a>ebra := new(_Ebra);
        <a id="L444"></a>p.re.add(ebra);
        <a id="L445"></a>bra.n = nbra;
        <a id="L446"></a>ebra.n = nbra;
        <a id="L447"></a>if start == nil {
            <a id="L448"></a>if end == nil {
                <a id="L449"></a>p.error = ErrInternal;
                <a id="L450"></a>return;
            <a id="L451"></a>}
            <a id="L452"></a>start = ebra;
        <a id="L453"></a>} else {
            <a id="L454"></a>end.setNext(ebra)
        <a id="L455"></a>}
        <a id="L456"></a>bra.setNext(start);
        <a id="L457"></a>return bra, ebra;
    <a id="L458"></a>case &#39;\\&#39;:
        <a id="L459"></a>c = p.nextc();
        <a id="L460"></a>switch {
        <a id="L461"></a>case c == endOfFile:
            <a id="L462"></a>p.error = ErrExtraneousBackslash;
            <a id="L463"></a>return;
        <a id="L464"></a>case c == &#39;n&#39;:
            <a id="L465"></a>c = &#39;\n&#39;
        <a id="L466"></a>case special(c):
        <a id="L467"></a><span class="comment">// c is as delivered</span>
        <a id="L468"></a>default:
            <a id="L469"></a>p.error = ErrBadBackslash;
            <a id="L470"></a>return;
        <a id="L471"></a>}
        <a id="L472"></a>fallthrough;
    <a id="L473"></a>default:
        <a id="L474"></a>p.nextc();
        <a id="L475"></a>start = newChar(c);
        <a id="L476"></a>p.re.add(start);
        <a id="L477"></a>return start, start;
    <a id="L478"></a>}
    <a id="L479"></a>panic(&#34;unreachable&#34;);
<a id="L480"></a>}

<a id="L482"></a>func (p *parser) closure() (start, end instr) {
    <a id="L483"></a>start, end = p.term();
    <a id="L484"></a>if start == nil || p.error != &#34;&#34; {
        <a id="L485"></a>return
    <a id="L486"></a>}
    <a id="L487"></a>switch p.c() {
    <a id="L488"></a>case &#39;*&#39;:
        <a id="L489"></a><span class="comment">// (start,end)*:</span>
        <a id="L490"></a>alt := new(_Alt);
        <a id="L491"></a>p.re.add(alt);
        <a id="L492"></a>end.setNext(alt); <span class="comment">// after end, do alt</span>
        <a id="L493"></a>alt.left = start; <span class="comment">// alternate brach: return to start</span>
        <a id="L494"></a>start = alt;      <span class="comment">// alt becomes new (start, end)</span>
        <a id="L495"></a>end = alt;
    <a id="L496"></a>case &#39;+&#39;:
        <a id="L497"></a><span class="comment">// (start,end)+:</span>
        <a id="L498"></a>alt := new(_Alt);
        <a id="L499"></a>p.re.add(alt);
        <a id="L500"></a>end.setNext(alt); <span class="comment">// after end, do alt</span>
        <a id="L501"></a>alt.left = start; <span class="comment">// alternate brach: return to start</span>
        <a id="L502"></a>end = alt;        <span class="comment">// start is unchanged; end is alt</span>
    <a id="L503"></a>case &#39;?&#39;:
        <a id="L504"></a><span class="comment">// (start,end)?:</span>
        <a id="L505"></a>alt := new(_Alt);
        <a id="L506"></a>p.re.add(alt);
        <a id="L507"></a>nop := new(_Nop);
        <a id="L508"></a>p.re.add(nop);
        <a id="L509"></a>alt.left = start; <span class="comment">// alternate branch is start</span>
        <a id="L510"></a>alt.setNext(nop); <span class="comment">// follow on to nop</span>
        <a id="L511"></a>end.setNext(nop); <span class="comment">// after end, go to nop</span>
        <a id="L512"></a>start = alt;      <span class="comment">// start is now alt</span>
        <a id="L513"></a>end = nop;        <span class="comment">// end is nop pointed to by both branches</span>
    <a id="L514"></a>default:
        <a id="L515"></a>return
    <a id="L516"></a>}
    <a id="L517"></a>switch p.nextc() {
    <a id="L518"></a>case &#39;*&#39;, &#39;+&#39;, &#39;?&#39;:
        <a id="L519"></a>p.error = ErrBadClosure
    <a id="L520"></a>}
    <a id="L521"></a>return;
<a id="L522"></a>}

<a id="L524"></a>func (p *parser) concatenation() (start, end instr) {
    <a id="L525"></a>for {
        <a id="L526"></a>nstart, nend := p.closure();
        <a id="L527"></a>if p.error != &#34;&#34; {
            <a id="L528"></a>return
        <a id="L529"></a>}
        <a id="L530"></a>switch {
        <a id="L531"></a>case nstart == nil: <span class="comment">// end of this concatenation</span>
            <a id="L532"></a>if start == nil { <span class="comment">// this is the empty string</span>
                <a id="L533"></a>nop := p.re.add(new(_Nop));
                <a id="L534"></a>return nop, nop;
            <a id="L535"></a>}
            <a id="L536"></a>return;
        <a id="L537"></a>case start == nil: <span class="comment">// this is first element of concatenation</span>
            <a id="L538"></a>start, end = nstart, nend
        <a id="L539"></a>default:
            <a id="L540"></a>end.setNext(nstart);
            <a id="L541"></a>end = nend;
        <a id="L542"></a>}
    <a id="L543"></a>}
    <a id="L544"></a>panic(&#34;unreachable&#34;);
<a id="L545"></a>}

<a id="L547"></a>func (p *parser) regexp() (start, end instr) {
    <a id="L548"></a>start, end = p.concatenation();
    <a id="L549"></a>if p.error != &#34;&#34; {
        <a id="L550"></a>return
    <a id="L551"></a>}
    <a id="L552"></a>for {
        <a id="L553"></a>switch p.c() {
        <a id="L554"></a>default:
            <a id="L555"></a>return
        <a id="L556"></a>case &#39;|&#39;:
            <a id="L557"></a>p.nextc();
            <a id="L558"></a>nstart, nend := p.concatenation();
            <a id="L559"></a>if p.error != &#34;&#34; {
                <a id="L560"></a>return
            <a id="L561"></a>}
            <a id="L562"></a>alt := new(_Alt);
            <a id="L563"></a>p.re.add(alt);
            <a id="L564"></a>alt.left = start;
            <a id="L565"></a>alt.setNext(nstart);
            <a id="L566"></a>nop := new(_Nop);
            <a id="L567"></a>p.re.add(nop);
            <a id="L568"></a>end.setNext(nop);
            <a id="L569"></a>nend.setNext(nop);
            <a id="L570"></a>start, end = alt, nop;
        <a id="L571"></a>}
    <a id="L572"></a>}
    <a id="L573"></a>panic(&#34;unreachable&#34;);
<a id="L574"></a>}

<a id="L576"></a>func unNop(i instr) instr {
    <a id="L577"></a>for i.kind() == _NOP {
        <a id="L578"></a>i = i.next()
    <a id="L579"></a>}
    <a id="L580"></a>return i;
<a id="L581"></a>}

<a id="L583"></a>func (re *Regexp) eliminateNops() {
    <a id="L584"></a>for i := 0; i &lt; len(re.inst); i++ {
        <a id="L585"></a>inst := re.inst[i];
        <a id="L586"></a>if inst.kind() == _END {
            <a id="L587"></a>continue
        <a id="L588"></a>}
        <a id="L589"></a>inst.setNext(unNop(inst.next()));
        <a id="L590"></a>if inst.kind() == _ALT {
            <a id="L591"></a>alt := inst.(*_Alt);
            <a id="L592"></a>alt.left = unNop(alt.left);
        <a id="L593"></a>}
    <a id="L594"></a>}
<a id="L595"></a>}

<a id="L597"></a>func (re *Regexp) doParse() string {
    <a id="L598"></a>p := newParser(re);
    <a id="L599"></a>start := new(_Start);
    <a id="L600"></a>re.add(start);
    <a id="L601"></a>s, e := p.regexp();
    <a id="L602"></a>if p.error != &#34;&#34; {
        <a id="L603"></a>return p.error
    <a id="L604"></a>}
    <a id="L605"></a>start.setNext(s);
    <a id="L606"></a>re.start = start;
    <a id="L607"></a>e.setNext(re.add(new(_End)));
    <a id="L608"></a>re.eliminateNops();
    <a id="L609"></a>return p.error;
<a id="L610"></a>}

<a id="L612"></a><span class="comment">// CompileRegexp parses a regular expression and returns, if successful, a Regexp</span>
<a id="L613"></a><span class="comment">// object that can be used to match against text.</span>
<a id="L614"></a>func CompileRegexp(str string) (regexp *Regexp, error string) {
    <a id="L615"></a>regexp = new(Regexp);
    <a id="L616"></a>regexp.expr = str;
    <a id="L617"></a>regexp.inst = make([]instr, 0, 20);
    <a id="L618"></a>error = regexp.doParse();
    <a id="L619"></a>return;
<a id="L620"></a>}

<a id="L622"></a><span class="comment">// MustCompileRegexp is like CompileRegexp but panics if the expression cannot be parsed.</span>
<a id="L623"></a><span class="comment">// It simplifies safe initialization of global variables holding compiled regular</span>
<a id="L624"></a><span class="comment">// expressions.</span>
<a id="L625"></a>func MustCompile(str string) *Regexp {
    <a id="L626"></a>regexp, error := CompileRegexp(str);
    <a id="L627"></a>if error != &#34;&#34; {
        <a id="L628"></a>panicln(`regexp: compiling &#34;`, str, `&#34;: `, error)
    <a id="L629"></a>}
    <a id="L630"></a>return regexp;
<a id="L631"></a>}

<a id="L633"></a>type state struct {
    <a id="L634"></a>inst  instr; <span class="comment">// next instruction to execute</span>
    <a id="L635"></a>match []int; <span class="comment">// pairs of bracketing submatches. 0th is start,end</span>
<a id="L636"></a>}

<a id="L638"></a><span class="comment">// Append new state to to-do list.  Leftmost-longest wins so avoid</span>
<a id="L639"></a><span class="comment">// adding a state that&#39;s already active.</span>
<a id="L640"></a>func addState(s []state, inst instr, match []int) []state {
    <a id="L641"></a>index := inst.index();
    <a id="L642"></a>l := len(s);
    <a id="L643"></a>pos := match[0];
    <a id="L644"></a><span class="comment">// TODO: Once the state is a vector and we can do insert, have inputs always</span>
    <a id="L645"></a><span class="comment">// go in order correctly and this &#34;earlier&#34; test is never necessary,</span>
    <a id="L646"></a>for i := 0; i &lt; l; i++ {
        <a id="L647"></a>if s[i].inst.index() == index &amp;&amp; <span class="comment">// same instruction</span>
            <a id="L648"></a>s[i].match[0] &lt; pos { <span class="comment">// earlier match already going; lefmost wins</span>
            <a id="L649"></a>return s
        <a id="L650"></a>}
    <a id="L651"></a>}
    <a id="L652"></a>if l == cap(s) {
        <a id="L653"></a>s1 := make([]state, 2*l)[0:l];
        <a id="L654"></a>for i := 0; i &lt; l; i++ {
            <a id="L655"></a>s1[i] = s[i]
        <a id="L656"></a>}
        <a id="L657"></a>s = s1;
    <a id="L658"></a>}
    <a id="L659"></a>s = s[0 : l+1];
    <a id="L660"></a>s[l].inst = inst;
    <a id="L661"></a>s[l].match = match;
    <a id="L662"></a>return s;
<a id="L663"></a>}

<a id="L665"></a><span class="comment">// Accepts either string or bytes - the logic is identical either way.</span>
<a id="L666"></a><span class="comment">// If bytes == nil, scan str.</span>
<a id="L667"></a>func (re *Regexp) doExecute(str string, bytes []byte, pos int) []int {
    <a id="L668"></a>var s [2][]state; <span class="comment">// TODO: use a vector when state values (not ptrs) can be vector elements</span>
    <a id="L669"></a>s[0] = make([]state, 10)[0:0];
    <a id="L670"></a>s[1] = make([]state, 10)[0:0];
    <a id="L671"></a>in, out := 0, 1;
    <a id="L672"></a>var final state;
    <a id="L673"></a>found := false;
    <a id="L674"></a>end := len(str);
    <a id="L675"></a>if bytes != nil {
        <a id="L676"></a>end = len(bytes)
    <a id="L677"></a>}
    <a id="L678"></a>for pos &lt;= end {
        <a id="L679"></a>if !found {
            <a id="L680"></a><span class="comment">// prime the pump if we haven&#39;t seen a match yet</span>
            <a id="L681"></a>match := make([]int, 2*(re.nbra+1));
            <a id="L682"></a>for i := 0; i &lt; len(match); i++ {
                <a id="L683"></a>match[i] = -1 <span class="comment">// no match seen; catches cases like &#34;a(b)?c&#34; on &#34;ac&#34;</span>
            <a id="L684"></a>}
            <a id="L685"></a>match[0] = pos;
            <a id="L686"></a>s[out] = addState(s[out], re.start.next(), match);
        <a id="L687"></a>}
        <a id="L688"></a>in, out = out, in;    <span class="comment">// old out state is new in state</span>
        <a id="L689"></a>s[out] = s[out][0:0]; <span class="comment">// clear out state</span>
        <a id="L690"></a>if len(s[in]) == 0 {
            <a id="L691"></a><span class="comment">// machine has completed</span>
            <a id="L692"></a>break
        <a id="L693"></a>}
        <a id="L694"></a>charwidth := 1;
        <a id="L695"></a>c := endOfFile;
        <a id="L696"></a>if pos &lt; end {
            <a id="L697"></a>if bytes == nil {
                <a id="L698"></a>c, charwidth = utf8.DecodeRuneInString(str[pos:end])
            <a id="L699"></a>} else {
                <a id="L700"></a>c, charwidth = utf8.DecodeRune(bytes[pos:end])
            <a id="L701"></a>}
        <a id="L702"></a>}
        <a id="L703"></a>for i := 0; i &lt; len(s[in]); i++ {
            <a id="L704"></a>st := s[in][i];
            <a id="L705"></a>switch s[in][i].inst.kind() {
            <a id="L706"></a>case _BOT:
                <a id="L707"></a>if pos == 0 {
                    <a id="L708"></a>s[in] = addState(s[in], st.inst.next(), st.match)
                <a id="L709"></a>}
            <a id="L710"></a>case _EOT:
                <a id="L711"></a>if pos == end {
                    <a id="L712"></a>s[in] = addState(s[in], st.inst.next(), st.match)
                <a id="L713"></a>}
            <a id="L714"></a>case _CHAR:
                <a id="L715"></a>if c == st.inst.(*_Char).char {
                    <a id="L716"></a>s[out] = addState(s[out], st.inst.next(), st.match)
                <a id="L717"></a>}
            <a id="L718"></a>case _CHARCLASS:
                <a id="L719"></a>if st.inst.(*_CharClass).matches(c) {
                    <a id="L720"></a>s[out] = addState(s[out], st.inst.next(), st.match)
                <a id="L721"></a>}
            <a id="L722"></a>case _ANY:
                <a id="L723"></a>if c != endOfFile {
                    <a id="L724"></a>s[out] = addState(s[out], st.inst.next(), st.match)
                <a id="L725"></a>}
            <a id="L726"></a>case _NOTNL:
                <a id="L727"></a>if c != endOfFile &amp;&amp; c != &#39;\n&#39; {
                    <a id="L728"></a>s[out] = addState(s[out], st.inst.next(), st.match)
                <a id="L729"></a>}
            <a id="L730"></a>case _BRA:
                <a id="L731"></a>n := st.inst.(*_Bra).n;
                <a id="L732"></a>st.match[2*n] = pos;
                <a id="L733"></a>s[in] = addState(s[in], st.inst.next(), st.match);
            <a id="L734"></a>case _EBRA:
                <a id="L735"></a>n := st.inst.(*_Ebra).n;
                <a id="L736"></a>st.match[2*n+1] = pos;
                <a id="L737"></a>s[in] = addState(s[in], st.inst.next(), st.match);
            <a id="L738"></a>case _ALT:
                <a id="L739"></a>s[in] = addState(s[in], st.inst.(*_Alt).left, st.match);
                <a id="L740"></a><span class="comment">// give other branch a copy of this match vector</span>
                <a id="L741"></a>s1 := make([]int, 2*(re.nbra+1));
                <a id="L742"></a>for i := 0; i &lt; len(s1); i++ {
                    <a id="L743"></a>s1[i] = st.match[i]
                <a id="L744"></a>}
                <a id="L745"></a>s[in] = addState(s[in], st.inst.next(), s1);
            <a id="L746"></a>case _END:
                <a id="L747"></a><span class="comment">// choose leftmost longest</span>
                <a id="L748"></a>if !found || <span class="comment">// first</span>
                    <a id="L749"></a>st.match[0] &lt; final.match[0] || <span class="comment">// leftmost</span>
                    <a id="L750"></a>(st.match[0] == final.match[0] &amp;&amp; pos &gt; final.match[1]) { <span class="comment">// longest</span>
                    <a id="L751"></a>final = st;
                    <a id="L752"></a>final.match[1] = pos;
                <a id="L753"></a>}
                <a id="L754"></a>found = true;
            <a id="L755"></a>default:
                <a id="L756"></a>st.inst.print();
                <a id="L757"></a>panic(&#34;unknown instruction in execute&#34;);
            <a id="L758"></a>}
        <a id="L759"></a>}
        <a id="L760"></a>pos += charwidth;
    <a id="L761"></a>}
    <a id="L762"></a>return final.match;
<a id="L763"></a>}


<a id="L766"></a><span class="comment">// ExecuteString matches the Regexp against the string s.</span>
<a id="L767"></a><span class="comment">// The return value is an array of integers, in pairs, identifying the positions of</span>
<a id="L768"></a><span class="comment">// substrings matched by the expression.</span>
<a id="L769"></a><span class="comment">//    s[a[0]:a[1]] is the substring matched by the entire expression.</span>
<a id="L770"></a><span class="comment">//    s[a[2*i]:a[2*i+1]] for i &gt; 0 is the substring matched by the ith parenthesized subexpression.</span>
<a id="L771"></a><span class="comment">// A negative value means the subexpression did not match any element of the string.</span>
<a id="L772"></a><span class="comment">// An empty array means &#34;no match&#34;.</span>
<a id="L773"></a>func (re *Regexp) ExecuteString(s string) (a []int) {
    <a id="L774"></a>return re.doExecute(s, nil, 0)
<a id="L775"></a>}


<a id="L778"></a><span class="comment">// Execute matches the Regexp against the byte slice b.</span>
<a id="L779"></a><span class="comment">// The return value is an array of integers, in pairs, identifying the positions of</span>
<a id="L780"></a><span class="comment">// subslices matched by the expression.</span>
<a id="L781"></a><span class="comment">//    b[a[0]:a[1]] is the subslice matched by the entire expression.</span>
<a id="L782"></a><span class="comment">//    b[a[2*i]:a[2*i+1]] for i &gt; 0 is the subslice matched by the ith parenthesized subexpression.</span>
<a id="L783"></a><span class="comment">// A negative value means the subexpression did not match any element of the slice.</span>
<a id="L784"></a><span class="comment">// An empty array means &#34;no match&#34;.</span>
<a id="L785"></a>func (re *Regexp) Execute(b []byte) (a []int) { return re.doExecute(&#34;&#34;, b, 0) }


<a id="L788"></a><span class="comment">// MatchString returns whether the Regexp matches the string s.</span>
<a id="L789"></a><span class="comment">// The return value is a boolean: true for match, false for no match.</span>
<a id="L790"></a>func (re *Regexp) MatchString(s string) bool { return len(re.doExecute(s, nil, 0)) &gt; 0 }


<a id="L793"></a><span class="comment">// Match returns whether the Regexp matches the byte slice b.</span>
<a id="L794"></a><span class="comment">// The return value is a boolean: true for match, false for no match.</span>
<a id="L795"></a>func (re *Regexp) Match(b []byte) bool { return len(re.doExecute(&#34;&#34;, b, 0)) &gt; 0 }


<a id="L798"></a><span class="comment">// MatchStrings matches the Regexp against the string s.</span>
<a id="L799"></a><span class="comment">// The return value is an array of strings matched by the expression.</span>
<a id="L800"></a><span class="comment">//    a[0] is the substring matched by the entire expression.</span>
<a id="L801"></a><span class="comment">//    a[i] for i &gt; 0 is the substring matched by the ith parenthesized subexpression.</span>
<a id="L802"></a><span class="comment">// An empty array means ``no match&#39;&#39;.</span>
<a id="L803"></a>func (re *Regexp) MatchStrings(s string) (a []string) {
    <a id="L804"></a>r := re.doExecute(s, nil, 0);
    <a id="L805"></a>if r == nil {
        <a id="L806"></a>return nil
    <a id="L807"></a>}
    <a id="L808"></a>a = make([]string, len(r)/2);
    <a id="L809"></a>for i := 0; i &lt; len(r); i += 2 {
        <a id="L810"></a>if r[i] != -1 { <span class="comment">// -1 means no match for this subexpression</span>
            <a id="L811"></a>a[i/2] = s[r[i]:r[i+1]]
        <a id="L812"></a>}
    <a id="L813"></a>}
    <a id="L814"></a>return;
<a id="L815"></a>}

<a id="L817"></a><span class="comment">// MatchSlices matches the Regexp against the byte slice b.</span>
<a id="L818"></a><span class="comment">// The return value is an array of subslices matched by the expression.</span>
<a id="L819"></a><span class="comment">//    a[0] is the subslice matched by the entire expression.</span>
<a id="L820"></a><span class="comment">//    a[i] for i &gt; 0 is the subslice matched by the ith parenthesized subexpression.</span>
<a id="L821"></a><span class="comment">// An empty array means ``no match&#39;&#39;.</span>
<a id="L822"></a>func (re *Regexp) MatchSlices(b []byte) (a [][]byte) {
    <a id="L823"></a>r := re.doExecute(&#34;&#34;, b, 0);
    <a id="L824"></a>if r == nil {
        <a id="L825"></a>return nil
    <a id="L826"></a>}
    <a id="L827"></a>a = make([][]byte, len(r)/2);
    <a id="L828"></a>for i := 0; i &lt; len(r); i += 2 {
        <a id="L829"></a>if r[i] != -1 { <span class="comment">// -1 means no match for this subexpression</span>
            <a id="L830"></a>a[i/2] = b[r[i]:r[i+1]]
        <a id="L831"></a>}
    <a id="L832"></a>}
    <a id="L833"></a>return;
<a id="L834"></a>}

<a id="L836"></a><span class="comment">// MatchString checks whether a textual regular expression</span>
<a id="L837"></a><span class="comment">// matches a string.  More complicated queries need</span>
<a id="L838"></a><span class="comment">// to use Compile and the full Regexp interface.</span>
<a id="L839"></a>func MatchString(pattern string, s string) (matched bool, error string) {
    <a id="L840"></a>re, err := CompileRegexp(pattern);
    <a id="L841"></a>if err != &#34;&#34; {
        <a id="L842"></a>return false, err
    <a id="L843"></a>}
    <a id="L844"></a>return re.MatchString(s), &#34;&#34;;
<a id="L845"></a>}

<a id="L847"></a><span class="comment">// Match checks whether a textual regular expression</span>
<a id="L848"></a><span class="comment">// matches a byte slice.  More complicated queries need</span>
<a id="L849"></a><span class="comment">// to use Compile and the full Regexp interface.</span>
<a id="L850"></a>func Match(pattern string, b []byte) (matched bool, error string) {
    <a id="L851"></a>re, err := CompileRegexp(pattern);
    <a id="L852"></a>if err != &#34;&#34; {
        <a id="L853"></a>return false, err
    <a id="L854"></a>}
    <a id="L855"></a>return re.Match(b), &#34;&#34;;
<a id="L856"></a>}
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
