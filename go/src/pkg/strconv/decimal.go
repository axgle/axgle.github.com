<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/decimal.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strconv/decimal.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Multiprecision decimal numbers.</span>
<a id="L6"></a><span class="comment">// For floating-point formatting only; not general purpose.</span>
<a id="L7"></a><span class="comment">// Only operations are assign and (binary) left/right shift.</span>
<a id="L8"></a><span class="comment">// Can do binary floating point in multiprecision decimal precisely</span>
<a id="L9"></a><span class="comment">// because 2 divides 10; cannot do decimal floating point</span>
<a id="L10"></a><span class="comment">// in multiprecision binary precisely.</span>

<a id="L12"></a>package strconv

<a id="L14"></a>import &#34;bytes&#34;

<a id="L16"></a>type decimal struct {
    <a id="L17"></a><span class="comment">// TODO(rsc): Can make d[] a bit smaller and add</span>
    <a id="L18"></a><span class="comment">// truncated bool;</span>
    <a id="L19"></a>d   [2000]byte; <span class="comment">// digits</span>
    <a id="L20"></a>nd  int;        <span class="comment">// number of digits used</span>
    <a id="L21"></a>dp  int;        <span class="comment">// decimal point</span>
<a id="L22"></a>}

<a id="L24"></a>func (a *decimal) String() string {
    <a id="L25"></a>n := 10 + a.nd;
    <a id="L26"></a>if a.dp &gt; 0 {
        <a id="L27"></a>n += a.dp
    <a id="L28"></a>}
    <a id="L29"></a>if a.dp &lt; 0 {
        <a id="L30"></a>n += -a.dp
    <a id="L31"></a>}

    <a id="L33"></a>buf := make([]byte, n);
    <a id="L34"></a>w := 0;
    <a id="L35"></a>switch {
    <a id="L36"></a>case a.nd == 0:
        <a id="L37"></a>return &#34;0&#34;

    <a id="L39"></a>case a.dp &lt;= 0:
        <a id="L40"></a><span class="comment">// zeros fill space between decimal point and digits</span>
        <a id="L41"></a>buf[w] = &#39;0&#39;;
        <a id="L42"></a>w++;
        <a id="L43"></a>buf[w] = &#39;.&#39;;
        <a id="L44"></a>w++;
        <a id="L45"></a>w += digitZero(buf[w : w+-a.dp]);
        <a id="L46"></a>w += bytes.Copy(buf[w:w+a.nd], a.d[0:a.nd]);

    <a id="L48"></a>case a.dp &lt; a.nd:
        <a id="L49"></a><span class="comment">// decimal point in middle of digits</span>
        <a id="L50"></a>w += bytes.Copy(buf[w:w+a.dp], a.d[0:a.dp]);
        <a id="L51"></a>buf[w] = &#39;.&#39;;
        <a id="L52"></a>w++;
        <a id="L53"></a>w += bytes.Copy(buf[w:w+a.nd-a.dp], a.d[a.dp:a.nd]);

    <a id="L55"></a>default:
        <a id="L56"></a><span class="comment">// zeros fill space between digits and decimal point</span>
        <a id="L57"></a>w += bytes.Copy(buf[w:w+a.nd], a.d[0:a.nd]);
        <a id="L58"></a>w += digitZero(buf[w : w+a.dp-a.nd]);
    <a id="L59"></a>}
    <a id="L60"></a>return string(buf[0:w]);
<a id="L61"></a>}

<a id="L63"></a>func copy(dst []byte, src []byte) int {
    <a id="L64"></a>for i := 0; i &lt; len(dst); i++ {
        <a id="L65"></a>dst[i] = src[i]
    <a id="L66"></a>}
    <a id="L67"></a>return len(dst);
<a id="L68"></a>}

<a id="L70"></a>func digitZero(dst []byte) int {
    <a id="L71"></a>for i := 0; i &lt; len(dst); i++ {
        <a id="L72"></a>dst[i] = &#39;0&#39;
    <a id="L73"></a>}
    <a id="L74"></a>return len(dst);
<a id="L75"></a>}

<a id="L77"></a><span class="comment">// trim trailing zeros from number.</span>
<a id="L78"></a><span class="comment">// (They are meaningless; the decimal point is tracked</span>
<a id="L79"></a><span class="comment">// independent of the number of digits.)</span>
<a id="L80"></a>func trim(a *decimal) {
    <a id="L81"></a>for a.nd &gt; 0 &amp;&amp; a.d[a.nd-1] == &#39;0&#39; {
        <a id="L82"></a>a.nd--
    <a id="L83"></a>}
    <a id="L84"></a>if a.nd == 0 {
        <a id="L85"></a>a.dp = 0
    <a id="L86"></a>}
<a id="L87"></a>}

<a id="L89"></a><span class="comment">// Assign v to a.</span>
<a id="L90"></a>func (a *decimal) Assign(v uint64) {
    <a id="L91"></a>var buf [50]byte;

    <a id="L93"></a><span class="comment">// Write reversed decimal in buf.</span>
    <a id="L94"></a>n := 0;
    <a id="L95"></a>for v &gt; 0 {
        <a id="L96"></a>v1 := v / 10;
        <a id="L97"></a>v -= 10 * v1;
        <a id="L98"></a>buf[n] = byte(v + &#39;0&#39;);
        <a id="L99"></a>n++;
        <a id="L100"></a>v = v1;
    <a id="L101"></a>}

    <a id="L103"></a><span class="comment">// Reverse again to produce forward decimal in a.d.</span>
    <a id="L104"></a>a.nd = 0;
    <a id="L105"></a>for n--; n &gt;= 0; n-- {
        <a id="L106"></a>a.d[a.nd] = buf[n];
        <a id="L107"></a>a.nd++;
    <a id="L108"></a>}
    <a id="L109"></a>a.dp = a.nd;
    <a id="L110"></a>trim(a);
<a id="L111"></a>}

<a id="L113"></a>func newDecimal(i uint64) *decimal {
    <a id="L114"></a>a := new(decimal);
    <a id="L115"></a>a.Assign(i);
    <a id="L116"></a>return a;
<a id="L117"></a>}

<a id="L119"></a><span class="comment">// Maximum shift that we can do in one pass without overflow.</span>
<a id="L120"></a><span class="comment">// Signed int has 31 bits, and we have to be able to accomodate 9&lt;&lt;k.</span>
<a id="L121"></a>const maxShift = 27

<a id="L123"></a><span class="comment">// Binary shift right (* 2) by k bits.  k &lt;= maxShift to avoid overflow.</span>
<a id="L124"></a>func rightShift(a *decimal, k uint) {
    <a id="L125"></a>r := 0; <span class="comment">// read pointer</span>
    <a id="L126"></a>w := 0; <span class="comment">// write pointer</span>

    <a id="L128"></a><span class="comment">// Pick up enough leading digits to cover first shift.</span>
    <a id="L129"></a>n := 0;
    <a id="L130"></a>for ; n&gt;&gt;k == 0; r++ {
        <a id="L131"></a>if r &gt;= a.nd {
            <a id="L132"></a>if n == 0 {
                <a id="L133"></a><span class="comment">// a == 0; shouldn&#39;t get here, but handle anyway.</span>
                <a id="L134"></a>a.nd = 0;
                <a id="L135"></a>return;
            <a id="L136"></a>}
            <a id="L137"></a>for n&gt;&gt;k == 0 {
                <a id="L138"></a>n = n * 10;
                <a id="L139"></a>r++;
            <a id="L140"></a>}
            <a id="L141"></a>break;
        <a id="L142"></a>}
        <a id="L143"></a>c := int(a.d[r]);
        <a id="L144"></a>n = n*10 + c - &#39;0&#39;;
    <a id="L145"></a>}
    <a id="L146"></a>a.dp -= r - 1;

    <a id="L148"></a><span class="comment">// Pick up a digit, put down a digit.</span>
    <a id="L149"></a>for ; r &lt; a.nd; r++ {
        <a id="L150"></a>c := int(a.d[r]);
        <a id="L151"></a>dig := n &gt;&gt; k;
        <a id="L152"></a>n -= dig &lt;&lt; k;
        <a id="L153"></a>a.d[w] = byte(dig + &#39;0&#39;);
        <a id="L154"></a>w++;
        <a id="L155"></a>n = n*10 + c - &#39;0&#39;;
    <a id="L156"></a>}

    <a id="L158"></a><span class="comment">// Put down extra digits.</span>
    <a id="L159"></a>for n &gt; 0 {
        <a id="L160"></a>dig := n &gt;&gt; k;
        <a id="L161"></a>n -= dig &lt;&lt; k;
        <a id="L162"></a>a.d[w] = byte(dig + &#39;0&#39;);
        <a id="L163"></a>w++;
        <a id="L164"></a>n = n * 10;
    <a id="L165"></a>}

    <a id="L167"></a>a.nd = w;
    <a id="L168"></a>trim(a);
<a id="L169"></a>}

<a id="L171"></a><span class="comment">// Cheat sheet for left shift: table indexed by shift count giving</span>
<a id="L172"></a><span class="comment">// number of new digits that will be introduced by that shift.</span>
<a id="L173"></a><span class="comment">//</span>
<a id="L174"></a><span class="comment">// For example, leftcheats[4] = {2, &#34;625&#34;}.  That means that</span>
<a id="L175"></a><span class="comment">// if we are shifting by 4 (multiplying by 16), it will add 2 digits</span>
<a id="L176"></a><span class="comment">// when the string prefix is &#34;625&#34; through &#34;999&#34;, and one fewer digit</span>
<a id="L177"></a><span class="comment">// if the string prefix is &#34;000&#34; through &#34;624&#34;.</span>
<a id="L178"></a><span class="comment">//</span>
<a id="L179"></a><span class="comment">// Credit for this trick goes to Ken.</span>

<a id="L181"></a>type leftCheat struct {
    <a id="L182"></a>delta  int;    <span class="comment">// number of new digits</span>
    <a id="L183"></a>cutoff string; <span class="comment">//   minus one digit if original &lt; a.</span>
<a id="L184"></a>}

<a id="L186"></a>var leftcheats = []leftCheat{
    <a id="L187"></a><span class="comment">// Leading digits of 1/2^i = 5^i.</span>
    <a id="L188"></a><span class="comment">// 5^23 is not an exact 64-bit floating point number,</span>
    <a id="L189"></a><span class="comment">// so have to use bc for the math.</span>
    <a id="L190"></a><span class="comment">/*</span>
    <a id="L191"></a><span class="comment">	seq 27 | sed &#39;s/^/5^/&#39; | bc |</span>
    <a id="L192"></a><span class="comment">	awk &#39;BEGIN{ print &#34;\tleftCheat{ 0, \&#34;\&#34; },&#34; }</span>
    <a id="L193"></a><span class="comment">	{</span>
    <a id="L194"></a><span class="comment">		log2 = log(2)/log(10)</span>
    <a id="L195"></a><span class="comment">		printf(&#34;\tleftCheat{ %d, \&#34;%s\&#34; },\t// * %d\n&#34;,</span>
    <a id="L196"></a><span class="comment">			int(log2*NR+1), $0, 2**NR)</span>
    <a id="L197"></a><span class="comment">	}&#39;</span>
    <a id="L198"></a><span class="comment">*/</span>
    <a id="L199"></a>leftCheat{0, &#34;&#34;},
    <a id="L200"></a>leftCheat{1, &#34;5&#34;}, <span class="comment">// * 2</span>
    <a id="L201"></a>leftCheat{1, &#34;25&#34;}, <span class="comment">// * 4</span>
    <a id="L202"></a>leftCheat{1, &#34;125&#34;}, <span class="comment">// * 8</span>
    <a id="L203"></a>leftCheat{2, &#34;625&#34;}, <span class="comment">// * 16</span>
    <a id="L204"></a>leftCheat{2, &#34;3125&#34;}, <span class="comment">// * 32</span>
    <a id="L205"></a>leftCheat{2, &#34;15625&#34;}, <span class="comment">// * 64</span>
    <a id="L206"></a>leftCheat{3, &#34;78125&#34;}, <span class="comment">// * 128</span>
    <a id="L207"></a>leftCheat{3, &#34;390625&#34;}, <span class="comment">// * 256</span>
    <a id="L208"></a>leftCheat{3, &#34;1953125&#34;}, <span class="comment">// * 512</span>
    <a id="L209"></a>leftCheat{4, &#34;9765625&#34;}, <span class="comment">// * 1024</span>
    <a id="L210"></a>leftCheat{4, &#34;48828125&#34;}, <span class="comment">// * 2048</span>
    <a id="L211"></a>leftCheat{4, &#34;244140625&#34;}, <span class="comment">// * 4096</span>
    <a id="L212"></a>leftCheat{4, &#34;1220703125&#34;}, <span class="comment">// * 8192</span>
    <a id="L213"></a>leftCheat{5, &#34;6103515625&#34;}, <span class="comment">// * 16384</span>
    <a id="L214"></a>leftCheat{5, &#34;30517578125&#34;}, <span class="comment">// * 32768</span>
    <a id="L215"></a>leftCheat{5, &#34;152587890625&#34;}, <span class="comment">// * 65536</span>
    <a id="L216"></a>leftCheat{6, &#34;762939453125&#34;}, <span class="comment">// * 131072</span>
    <a id="L217"></a>leftCheat{6, &#34;3814697265625&#34;}, <span class="comment">// * 262144</span>
    <a id="L218"></a>leftCheat{6, &#34;19073486328125&#34;}, <span class="comment">// * 524288</span>
    <a id="L219"></a>leftCheat{7, &#34;95367431640625&#34;}, <span class="comment">// * 1048576</span>
    <a id="L220"></a>leftCheat{7, &#34;476837158203125&#34;}, <span class="comment">// * 2097152</span>
    <a id="L221"></a>leftCheat{7, &#34;2384185791015625&#34;}, <span class="comment">// * 4194304</span>
    <a id="L222"></a>leftCheat{7, &#34;11920928955078125&#34;}, <span class="comment">// * 8388608</span>
    <a id="L223"></a>leftCheat{8, &#34;59604644775390625&#34;}, <span class="comment">// * 16777216</span>
    <a id="L224"></a>leftCheat{8, &#34;298023223876953125&#34;}, <span class="comment">// * 33554432</span>
    <a id="L225"></a>leftCheat{8, &#34;1490116119384765625&#34;}, <span class="comment">// * 67108864</span>
    <a id="L226"></a>leftCheat{9, &#34;7450580596923828125&#34;}, <span class="comment">// * 134217728</span>
<a id="L227"></a>}

<a id="L229"></a><span class="comment">// Is the leading prefix of b lexicographically less than s?</span>
<a id="L230"></a>func prefixIsLessThan(b []byte, s string) bool {
    <a id="L231"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L232"></a>if i &gt;= len(b) {
            <a id="L233"></a>return true
        <a id="L234"></a>}
        <a id="L235"></a>if b[i] != s[i] {
            <a id="L236"></a>return b[i] &lt; s[i]
        <a id="L237"></a>}
    <a id="L238"></a>}
    <a id="L239"></a>return false;
<a id="L240"></a>}

<a id="L242"></a><span class="comment">// Binary shift left (/ 2) by k bits.  k &lt;= maxShift to avoid overflow.</span>
<a id="L243"></a>func leftShift(a *decimal, k uint) {
    <a id="L244"></a>delta := leftcheats[k].delta;
    <a id="L245"></a>if prefixIsLessThan(a.d[0:a.nd], leftcheats[k].cutoff) {
        <a id="L246"></a>delta--
    <a id="L247"></a>}

    <a id="L249"></a>r := a.nd;         <span class="comment">// read index</span>
    <a id="L250"></a>w := a.nd + delta; <span class="comment">// write index</span>
    <a id="L251"></a>n := 0;

    <a id="L253"></a><span class="comment">// Pick up a digit, put down a digit.</span>
    <a id="L254"></a>for r--; r &gt;= 0; r-- {
        <a id="L255"></a>n += (int(a.d[r]) - &#39;0&#39;) &lt;&lt; k;
        <a id="L256"></a>quo := n / 10;
        <a id="L257"></a>rem := n - 10*quo;
        <a id="L258"></a>w--;
        <a id="L259"></a>a.d[w] = byte(rem + &#39;0&#39;);
        <a id="L260"></a>n = quo;
    <a id="L261"></a>}

    <a id="L263"></a><span class="comment">// Put down extra digits.</span>
    <a id="L264"></a>for n &gt; 0 {
        <a id="L265"></a>quo := n / 10;
        <a id="L266"></a>rem := n - 10*quo;
        <a id="L267"></a>w--;
        <a id="L268"></a>a.d[w] = byte(rem + &#39;0&#39;);
        <a id="L269"></a>n = quo;
    <a id="L270"></a>}

    <a id="L272"></a>a.nd += delta;
    <a id="L273"></a>a.dp += delta;
    <a id="L274"></a>trim(a);
<a id="L275"></a>}

<a id="L277"></a><span class="comment">// Binary shift left (k &gt; 0) or right (k &lt; 0).</span>
<a id="L278"></a><span class="comment">// Returns receiver for convenience.</span>
<a id="L279"></a>func (a *decimal) Shift(k int) *decimal {
    <a id="L280"></a>switch {
    <a id="L281"></a>case a.nd == 0:
        <a id="L282"></a><span class="comment">// nothing to do: a == 0</span>
    <a id="L283"></a>case k &gt; 0:
        <a id="L284"></a>for k &gt; maxShift {
            <a id="L285"></a>leftShift(a, maxShift);
            <a id="L286"></a>k -= maxShift;
        <a id="L287"></a>}
        <a id="L288"></a>leftShift(a, uint(k));
    <a id="L289"></a>case k &lt; 0:
        <a id="L290"></a>for k &lt; -maxShift {
            <a id="L291"></a>rightShift(a, maxShift);
            <a id="L292"></a>k += maxShift;
        <a id="L293"></a>}
        <a id="L294"></a>rightShift(a, uint(-k));
    <a id="L295"></a>}
    <a id="L296"></a>return a;
<a id="L297"></a>}

<a id="L299"></a><span class="comment">// If we chop a at nd digits, should we round up?</span>
<a id="L300"></a>func shouldRoundUp(a *decimal, nd int) bool {
    <a id="L301"></a>if nd &lt;= 0 || nd &gt;= a.nd {
        <a id="L302"></a>return false
    <a id="L303"></a>}
    <a id="L304"></a>if a.d[nd] == &#39;5&#39; &amp;&amp; nd+1 == a.nd { <span class="comment">// exactly halfway - round to even</span>
        <a id="L305"></a>return (a.d[nd-1]-&#39;0&#39;)%2 != 0
    <a id="L306"></a>}
    <a id="L307"></a><span class="comment">// not halfway - digit tells all</span>
    <a id="L308"></a>return a.d[nd] &gt;= &#39;5&#39;;
<a id="L309"></a>}

<a id="L311"></a><span class="comment">// Round a to nd digits (or fewer).</span>
<a id="L312"></a><span class="comment">// Returns receiver for convenience.</span>
<a id="L313"></a>func (a *decimal) Round(nd int) *decimal {
    <a id="L314"></a>if nd &lt;= 0 || nd &gt;= a.nd {
        <a id="L315"></a>return a
    <a id="L316"></a>}
    <a id="L317"></a>if shouldRoundUp(a, nd) {
        <a id="L318"></a>return a.RoundUp(nd)
    <a id="L319"></a>}
    <a id="L320"></a>return a.RoundDown(nd);
<a id="L321"></a>}

<a id="L323"></a><span class="comment">// Round a down to nd digits (or fewer).</span>
<a id="L324"></a><span class="comment">// Returns receiver for convenience.</span>
<a id="L325"></a>func (a *decimal) RoundDown(nd int) *decimal {
    <a id="L326"></a>if nd &lt;= 0 || nd &gt;= a.nd {
        <a id="L327"></a>return a
    <a id="L328"></a>}
    <a id="L329"></a>a.nd = nd;
    <a id="L330"></a>trim(a);
    <a id="L331"></a>return a;
<a id="L332"></a>}

<a id="L334"></a><span class="comment">// Round a up to nd digits (or fewer).</span>
<a id="L335"></a><span class="comment">// Returns receiver for convenience.</span>
<a id="L336"></a>func (a *decimal) RoundUp(nd int) *decimal {
    <a id="L337"></a>if nd &lt;= 0 || nd &gt;= a.nd {
        <a id="L338"></a>return a
    <a id="L339"></a>}

    <a id="L341"></a><span class="comment">// round up</span>
    <a id="L342"></a>for i := nd - 1; i &gt;= 0; i-- {
        <a id="L343"></a>c := a.d[i];
        <a id="L344"></a>if c &lt; &#39;9&#39; { <span class="comment">// can stop after this digit</span>
            <a id="L345"></a>a.d[i]++;
            <a id="L346"></a>a.nd = i + 1;
            <a id="L347"></a>return a;
        <a id="L348"></a>}
    <a id="L349"></a>}

    <a id="L351"></a><span class="comment">// Number is all 9s.</span>
    <a id="L352"></a><span class="comment">// Change to single 1 with adjusted decimal point.</span>
    <a id="L353"></a>a.d[0] = &#39;1&#39;;
    <a id="L354"></a>a.nd = 1;
    <a id="L355"></a>a.dp++;
    <a id="L356"></a>return a;
<a id="L357"></a>}

<a id="L359"></a><span class="comment">// Extract integer part, rounded appropriately.</span>
<a id="L360"></a><span class="comment">// No guarantees about overflow.</span>
<a id="L361"></a>func (a *decimal) RoundedInteger() uint64 {
    <a id="L362"></a>if a.dp &gt; 20 {
        <a id="L363"></a>return 0xFFFFFFFFFFFFFFFF
    <a id="L364"></a>}
    <a id="L365"></a>var i int;
    <a id="L366"></a>n := uint64(0);
    <a id="L367"></a>for i = 0; i &lt; a.dp &amp;&amp; i &lt; a.nd; i++ {
        <a id="L368"></a>n = n*10 + uint64(a.d[i]-&#39;0&#39;)
    <a id="L369"></a>}
    <a id="L370"></a>for ; i &lt; a.dp; i++ {
        <a id="L371"></a>n *= 10
    <a id="L372"></a>}
    <a id="L373"></a>if shouldRoundUp(a, a.dp) {
        <a id="L374"></a>n++
    <a id="L375"></a>}
    <a id="L376"></a>return n;
<a id="L377"></a>}
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
