<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/big/int.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/big/int.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file implements signed multi-precision integers.</span>

<a id="L7"></a>package big

<a id="L9"></a><span class="comment">// An Int represents a signed multi-precision integer.</span>
<a id="L10"></a><span class="comment">// The zero value for an Int represents the value 0.</span>
<a id="L11"></a>type Int struct {
    <a id="L12"></a>neg bool;   <span class="comment">// sign</span>
    <a id="L13"></a>abs []Word; <span class="comment">// absolute value of the integer</span>
<a id="L14"></a>}


<a id="L17"></a><span class="comment">// New allocates and returns a new Int set to x.</span>
<a id="L18"></a>func (z *Int) New(x int64) *Int {
    <a id="L19"></a>z.neg = false;
    <a id="L20"></a>if x &lt; 0 {
        <a id="L21"></a>z.neg = true;
        <a id="L22"></a>x = -x;
    <a id="L23"></a>}
    <a id="L24"></a>z.abs = newN(z.abs, uint64(x));
    <a id="L25"></a>return z;
<a id="L26"></a>}


<a id="L29"></a><span class="comment">// NewInt allocates and returns a new Int set to x.</span>
<a id="L30"></a>func NewInt(x int64) *Int { return new(Int).New(x) }


<a id="L33"></a><span class="comment">// Set sets z to x.</span>
<a id="L34"></a>func (z *Int) Set(x *Int) *Int {
    <a id="L35"></a>z.neg = x.neg;
    <a id="L36"></a>z.abs = setN(z.abs, x.abs);
    <a id="L37"></a>return z;
<a id="L38"></a>}


<a id="L41"></a><span class="comment">// Add computes z = x+y.</span>
<a id="L42"></a>func (z *Int) Add(x, y *Int) *Int {
    <a id="L43"></a>if x.neg == y.neg {
        <a id="L44"></a><span class="comment">// x + y == x + y</span>
        <a id="L45"></a><span class="comment">// (-x) + (-y) == -(x + y)</span>
        <a id="L46"></a>z.neg = x.neg;
        <a id="L47"></a>z.abs = addNN(z.abs, x.abs, y.abs);
    <a id="L48"></a>} else {
        <a id="L49"></a><span class="comment">// x + (-y) == x - y == -(y - x)</span>
        <a id="L50"></a><span class="comment">// (-x) + y == y - x == -(x - y)</span>
        <a id="L51"></a>if cmpNN(x.abs, y.abs) &gt;= 0 {
            <a id="L52"></a>z.neg = x.neg;
            <a id="L53"></a>z.abs = subNN(z.abs, x.abs, y.abs);
        <a id="L54"></a>} else {
            <a id="L55"></a>z.neg = !x.neg;
            <a id="L56"></a>z.abs = subNN(z.abs, y.abs, x.abs);
        <a id="L57"></a>}
    <a id="L58"></a>}
    <a id="L59"></a>if len(z.abs) == 0 {
        <a id="L60"></a>z.neg = false <span class="comment">// 0 has no sign</span>
    <a id="L61"></a>}
    <a id="L62"></a>return z;
<a id="L63"></a>}


<a id="L66"></a><span class="comment">// Sub computes z = x-y.</span>
<a id="L67"></a>func (z *Int) Sub(x, y *Int) *Int {
    <a id="L68"></a>if x.neg != y.neg {
        <a id="L69"></a><span class="comment">// x - (-y) == x + y</span>
        <a id="L70"></a><span class="comment">// (-x) - y == -(x + y)</span>
        <a id="L71"></a>z.neg = x.neg;
        <a id="L72"></a>z.abs = addNN(z.abs, x.abs, y.abs);
    <a id="L73"></a>} else {
        <a id="L74"></a><span class="comment">// x - y == x - y == -(y - x)</span>
        <a id="L75"></a><span class="comment">// (-x) - (-y) == y - x == -(x - y)</span>
        <a id="L76"></a>if cmpNN(x.abs, y.abs) &gt;= 0 {
            <a id="L77"></a>z.neg = x.neg;
            <a id="L78"></a>z.abs = subNN(z.abs, x.abs, y.abs);
        <a id="L79"></a>} else {
            <a id="L80"></a>z.neg = !x.neg;
            <a id="L81"></a>z.abs = subNN(z.abs, y.abs, x.abs);
        <a id="L82"></a>}
    <a id="L83"></a>}
    <a id="L84"></a>if len(z.abs) == 0 {
        <a id="L85"></a>z.neg = false <span class="comment">// 0 has no sign</span>
    <a id="L86"></a>}
    <a id="L87"></a>return z;
<a id="L88"></a>}


<a id="L91"></a><span class="comment">// Mul computes z = x*y.</span>
<a id="L92"></a>func (z *Int) Mul(x, y *Int) *Int {
    <a id="L93"></a><span class="comment">// x * y == x * y</span>
    <a id="L94"></a><span class="comment">// x * (-y) == -(x * y)</span>
    <a id="L95"></a><span class="comment">// (-x) * y == -(x * y)</span>
    <a id="L96"></a><span class="comment">// (-x) * (-y) == x * y</span>
    <a id="L97"></a>z.abs = mulNN(z.abs, x.abs, y.abs);
    <a id="L98"></a>z.neg = len(z.abs) &gt; 0 &amp;&amp; x.neg != y.neg; <span class="comment">// 0 has no sign</span>
    <a id="L99"></a>return z;
<a id="L100"></a>}


<a id="L103"></a><span class="comment">// Div calculates q = (x-r)/y where 0 &lt;= r &lt; y. The receiver is set to q.</span>
<a id="L104"></a>func (z *Int) Div(x, y *Int) (q, r *Int) {
    <a id="L105"></a>q = z;
    <a id="L106"></a>r = new(Int);
    <a id="L107"></a>div(q, r, x, y);
    <a id="L108"></a>return;
<a id="L109"></a>}


<a id="L112"></a><span class="comment">// Mod calculates q = (x-r)/y and returns r.</span>
<a id="L113"></a>func (z *Int) Mod(x, y *Int) (r *Int) {
    <a id="L114"></a>q := new(Int);
    <a id="L115"></a>r = z;
    <a id="L116"></a>div(q, r, x, y);
    <a id="L117"></a>return;
<a id="L118"></a>}


<a id="L121"></a>func div(q, r, x, y *Int) {
    <a id="L122"></a>q.neg = x.neg != y.neg;
    <a id="L123"></a>r.neg = x.neg;
    <a id="L124"></a>q.abs, r.abs = divNN(q.abs, r.abs, x.abs, y.abs);
    <a id="L125"></a>return;
<a id="L126"></a>}


<a id="L129"></a><span class="comment">// Neg computes z = -x.</span>
<a id="L130"></a>func (z *Int) Neg(x *Int) *Int {
    <a id="L131"></a>z.abs = setN(z.abs, x.abs);
    <a id="L132"></a>z.neg = len(z.abs) &gt; 0 &amp;&amp; !x.neg; <span class="comment">// 0 has no sign</span>
    <a id="L133"></a>return z;
<a id="L134"></a>}


<a id="L137"></a><span class="comment">// Cmp compares x and y. The result is</span>
<a id="L138"></a><span class="comment">//</span>
<a id="L139"></a><span class="comment">//   -1 if x &lt;  y</span>
<a id="L140"></a><span class="comment">//    0 if x == y</span>
<a id="L141"></a><span class="comment">//   +1 if x &gt;  y</span>
<a id="L142"></a><span class="comment">//</span>
<a id="L143"></a>func (x *Int) Cmp(y *Int) (r int) {
    <a id="L144"></a><span class="comment">// x cmp y == x cmp y</span>
    <a id="L145"></a><span class="comment">// x cmp (-y) == x</span>
    <a id="L146"></a><span class="comment">// (-x) cmp y == y</span>
    <a id="L147"></a><span class="comment">// (-x) cmp (-y) == -(x cmp y)</span>
    <a id="L148"></a>switch {
    <a id="L149"></a>case x.neg == y.neg:
        <a id="L150"></a>r = cmpNN(x.abs, y.abs);
        <a id="L151"></a>if x.neg {
            <a id="L152"></a>r = -r
        <a id="L153"></a>}
    <a id="L154"></a>case x.neg:
        <a id="L155"></a>r = -1
    <a id="L156"></a>default:
        <a id="L157"></a>r = 1
    <a id="L158"></a>}
    <a id="L159"></a>return;
<a id="L160"></a>}


<a id="L163"></a>func (z *Int) String() string {
    <a id="L164"></a>s := &#34;&#34;;
    <a id="L165"></a>if z.neg {
        <a id="L166"></a>s = &#34;-&#34;
    <a id="L167"></a>}
    <a id="L168"></a>return s + stringN(z.abs, 10);
<a id="L169"></a>}


<a id="L172"></a><span class="comment">// SetString sets z to the value of s, interpreted in the given base.</span>
<a id="L173"></a><span class="comment">// If base is 0 then SetString attempts to detect the base by at the prefix of</span>
<a id="L174"></a><span class="comment">// s. &#39;0x&#39; implies base 16, &#39;0&#39; implies base 8. Otherwise base 10 is assumed.</span>
<a id="L175"></a>func (z *Int) SetString(s string, base int) (*Int, bool) {
    <a id="L176"></a>var scanned int;

    <a id="L178"></a>if base == 1 || base &gt; 16 {
        <a id="L179"></a>goto Error
    <a id="L180"></a>}

    <a id="L182"></a>if len(s) == 0 {
        <a id="L183"></a>goto Error
    <a id="L184"></a>}

    <a id="L186"></a>if s[0] == &#39;-&#39; {
        <a id="L187"></a>z.neg = true;
        <a id="L188"></a>s = s[1:len(s)];
    <a id="L189"></a>} else {
        <a id="L190"></a>z.neg = false
    <a id="L191"></a>}

    <a id="L193"></a>z.abs, _, scanned = scanN(z.abs, s, base);
    <a id="L194"></a>if scanned != len(s) {
        <a id="L195"></a>goto Error
    <a id="L196"></a>}

    <a id="L198"></a>return z, true;

<a id="L200"></a>Error:
    <a id="L201"></a>z.neg = false;
    <a id="L202"></a>z.abs = nil;
    <a id="L203"></a>return nil, false;
<a id="L204"></a>}


<a id="L207"></a><span class="comment">// SetBytes interprets b as the bytes of a big-endian, unsigned integer and</span>
<a id="L208"></a><span class="comment">// sets x to that value.</span>
<a id="L209"></a>func (z *Int) SetBytes(b []byte) *Int {
    <a id="L210"></a>s := int(_S);
    <a id="L211"></a>z.abs = makeN(z.abs, (len(b)+s-1)/s, false);
    <a id="L212"></a>z.neg = false;

    <a id="L214"></a>j := 0;
    <a id="L215"></a>for len(b) &gt;= s {
        <a id="L216"></a>var w Word;

        <a id="L218"></a>for i := s; i &gt; 0; i-- {
            <a id="L219"></a>w &lt;&lt;= 8;
            <a id="L220"></a>w |= Word(b[len(b)-i]);
        <a id="L221"></a>}

        <a id="L223"></a>z.abs[j] = w;
        <a id="L224"></a>j++;
        <a id="L225"></a>b = b[0 : len(b)-s];
    <a id="L226"></a>}

    <a id="L228"></a>if len(b) &gt; 0 {
        <a id="L229"></a>var w Word;

        <a id="L231"></a>for i := len(b); i &gt; 0; i-- {
            <a id="L232"></a>w &lt;&lt;= 8;
            <a id="L233"></a>w |= Word(b[len(b)-i]);
        <a id="L234"></a>}

        <a id="L236"></a>z.abs[j] = w;
    <a id="L237"></a>}

    <a id="L239"></a>z.abs = normN(z.abs);

    <a id="L241"></a>return z;
<a id="L242"></a>}


<a id="L245"></a><span class="comment">// Bytes returns the absolute value of x as a big-endian byte array.</span>
<a id="L246"></a>func (z *Int) Bytes() []byte {
    <a id="L247"></a>s := int(_S);
    <a id="L248"></a>b := make([]byte, len(z.abs)*s);

    <a id="L250"></a>for i, w := range z.abs {
        <a id="L251"></a>wordBytes := b[(len(z.abs)-i-1)*s : (len(z.abs)-i)*s];
        <a id="L252"></a>for j := s - 1; j &gt;= 0; j-- {
            <a id="L253"></a>wordBytes[j] = byte(w);
            <a id="L254"></a>w &gt;&gt;= 8;
        <a id="L255"></a>}
    <a id="L256"></a>}

    <a id="L258"></a>i := 0;
    <a id="L259"></a>for i &lt; len(b) &amp;&amp; b[i] == 0 {
        <a id="L260"></a>i++
    <a id="L261"></a>}

    <a id="L263"></a>return b[i:len(b)];
<a id="L264"></a>}


<a id="L267"></a><span class="comment">// Len returns the length of the absolute value of x in bits. Zero is</span>
<a id="L268"></a><span class="comment">// considered to have a length of one.</span>
<a id="L269"></a>func (z *Int) Len() int {
    <a id="L270"></a>if len(z.abs) == 0 {
        <a id="L271"></a>return 0
    <a id="L272"></a>}

    <a id="L274"></a>return len(z.abs)*_W - int(leadingZeros(z.abs[len(z.abs)-1]));
<a id="L275"></a>}


<a id="L278"></a><span class="comment">// Exp sets z = x**y mod m. If m is nil, z = x**y.</span>
<a id="L279"></a><span class="comment">// See Knuth, volume 2, section 4.6.3.</span>
<a id="L280"></a>func (z *Int) Exp(x, y, m *Int) *Int {
    <a id="L281"></a>if y.neg || len(y.abs) == 0 {
        <a id="L282"></a>z.New(1);
        <a id="L283"></a>z.neg = x.neg;
        <a id="L284"></a>return z;
    <a id="L285"></a>}

    <a id="L287"></a>var mWords []Word;
    <a id="L288"></a>if m != nil {
        <a id="L289"></a>mWords = m.abs
    <a id="L290"></a>}

    <a id="L292"></a>z.abs = expNNN(z.abs, x.abs, y.abs, mWords);
    <a id="L293"></a>z.neg = x.neg &amp;&amp; y.abs[0]&amp;1 == 1;
    <a id="L294"></a>return z;
<a id="L295"></a>}


<a id="L298"></a><span class="comment">// GcdInt sets d to the greatest common divisor of a and b, which must be</span>
<a id="L299"></a><span class="comment">// positive numbers.</span>
<a id="L300"></a><span class="comment">// If x and y are not nil, GcdInt sets x and y such that d = a*x + b*y.</span>
<a id="L301"></a><span class="comment">// If either a or b is not positive, GcdInt sets d = x = y = 0.</span>
<a id="L302"></a>func GcdInt(d, x, y, a, b *Int) {
    <a id="L303"></a>if a.neg || b.neg {
        <a id="L304"></a>d.New(0);
        <a id="L305"></a>if x != nil {
            <a id="L306"></a>x.New(0)
        <a id="L307"></a>}
        <a id="L308"></a>if y != nil {
            <a id="L309"></a>y.New(0)
        <a id="L310"></a>}
        <a id="L311"></a>return;
    <a id="L312"></a>}

    <a id="L314"></a>A := new(Int).Set(a);
    <a id="L315"></a>B := new(Int).Set(b);

    <a id="L317"></a>X := new(Int);
    <a id="L318"></a>Y := new(Int).New(1);

    <a id="L320"></a>lastX := new(Int).New(1);
    <a id="L321"></a>lastY := new(Int);

    <a id="L323"></a>q := new(Int);
    <a id="L324"></a>temp := new(Int);

    <a id="L326"></a>for len(B.abs) &gt; 0 {
        <a id="L327"></a>q, r := q.Div(A, B);

        <a id="L329"></a>A, B = B, r;

        <a id="L331"></a>temp.Set(X);
        <a id="L332"></a>X.Mul(X, q);
        <a id="L333"></a>X.neg = !X.neg;
        <a id="L334"></a>X.Add(X, lastX);
        <a id="L335"></a>lastX.Set(temp);

        <a id="L337"></a>temp.Set(Y);
        <a id="L338"></a>Y.Mul(Y, q);
        <a id="L339"></a>Y.neg = !Y.neg;
        <a id="L340"></a>Y.Add(Y, lastY);
        <a id="L341"></a>lastY.Set(temp);
    <a id="L342"></a>}

    <a id="L344"></a>if x != nil {
        <a id="L345"></a>*x = *lastX
    <a id="L346"></a>}

    <a id="L348"></a>if y != nil {
        <a id="L349"></a>*y = *lastY
    <a id="L350"></a>}

    <a id="L352"></a>*d = *A;
<a id="L353"></a>}


<a id="L356"></a><span class="comment">// ProbablyPrime performs n Miller-Rabin tests to check whether z is prime.</span>
<a id="L357"></a><span class="comment">// If it returns true, z is prime with probability 1 - 1/4^n.</span>
<a id="L358"></a><span class="comment">// If it returns false, z is not prime.</span>
<a id="L359"></a>func ProbablyPrime(z *Int, reps int) bool { return !z.neg &amp;&amp; probablyPrime(z.abs, reps) }


<a id="L362"></a><span class="comment">// Rsh sets z = x &gt;&gt; s and returns z.</span>
<a id="L363"></a>func (z *Int) Rsh(x *Int, n int) *Int {
    <a id="L364"></a>removedWords := n / _W;
    <a id="L365"></a>z.abs = makeN(z.abs, len(x.abs)-removedWords, false);
    <a id="L366"></a>z.neg = x.neg;
    <a id="L367"></a>shiftRight(z.abs, x.abs[removedWords:len(x.abs)], n%_W);
    <a id="L368"></a>z.abs = normN(z.abs);
    <a id="L369"></a>return z;
<a id="L370"></a>}
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
