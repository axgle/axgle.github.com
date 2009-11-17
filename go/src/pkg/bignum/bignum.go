<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bignum/bignum.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/bignum/bignum.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// A package for arbitrary precision arithmethic.</span>
<a id="L6"></a><span class="comment">// It implements the following numeric types:</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a><span class="comment">//	- Natural	unsigned integers</span>
<a id="L9"></a><span class="comment">//	- Integer	signed integers</span>
<a id="L10"></a><span class="comment">//	- Rational	rational numbers</span>
<a id="L11"></a><span class="comment">//</span>
<a id="L12"></a><span class="comment">// This package has been designed for ease of use but the functions it provides</span>
<a id="L13"></a><span class="comment">// are likely to be quite slow. It may be deprecated eventually. Use package</span>
<a id="L14"></a><span class="comment">// big instead, if possible.</span>
<a id="L15"></a><span class="comment">//</span>
<a id="L16"></a>package bignum

<a id="L18"></a>import (
    <a id="L19"></a>&#34;fmt&#34;;
<a id="L20"></a>)

<a id="L22"></a><span class="comment">// TODO(gri) Complete the set of in-place operations.</span>

<a id="L24"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L25"></a><span class="comment">// Internal representation</span>
<a id="L26"></a><span class="comment">//</span>
<a id="L27"></a><span class="comment">// A natural number of the form</span>
<a id="L28"></a><span class="comment">//</span>
<a id="L29"></a><span class="comment">//   x = x[n-1]*B^(n-1) + x[n-2]*B^(n-2) + ... + x[1]*B + x[0]</span>
<a id="L30"></a><span class="comment">//</span>
<a id="L31"></a><span class="comment">// with 0 &lt;= x[i] &lt; B and 0 &lt;= i &lt; n is stored in a slice of length n,</span>
<a id="L32"></a><span class="comment">// with the digits x[i] as the slice elements.</span>
<a id="L33"></a><span class="comment">//</span>
<a id="L34"></a><span class="comment">// A natural number is normalized if the slice contains no leading 0 digits.</span>
<a id="L35"></a><span class="comment">// During arithmetic operations, denormalized values may occur but are</span>
<a id="L36"></a><span class="comment">// always normalized before returning the final result. The normalized</span>
<a id="L37"></a><span class="comment">// representation of 0 is the empty slice (length = 0).</span>
<a id="L38"></a><span class="comment">//</span>
<a id="L39"></a><span class="comment">// The operations for all other numeric types are implemented on top of</span>
<a id="L40"></a><span class="comment">// the operations for natural numbers.</span>
<a id="L41"></a><span class="comment">//</span>
<a id="L42"></a><span class="comment">// The base B is chosen as large as possible on a given platform but there</span>
<a id="L43"></a><span class="comment">// are a few constraints besides the size of the largest unsigned integer</span>
<a id="L44"></a><span class="comment">// type available:</span>
<a id="L45"></a><span class="comment">//</span>
<a id="L46"></a><span class="comment">// 1) To improve conversion speed between strings and numbers, the base B</span>
<a id="L47"></a><span class="comment">//    is chosen such that division and multiplication by 10 (for decimal</span>
<a id="L48"></a><span class="comment">//    string representation) can be done without using extended-precision</span>
<a id="L49"></a><span class="comment">//    arithmetic. This makes addition, subtraction, and conversion routines</span>
<a id="L50"></a><span class="comment">//    twice as fast. It requires a ``buffer&#39;&#39; of 4 bits per operand digit.</span>
<a id="L51"></a><span class="comment">//    That is, the size of B must be 4 bits smaller then the size of the</span>
<a id="L52"></a><span class="comment">//    type (digit) in which these operations are performed. Having this</span>
<a id="L53"></a><span class="comment">//    buffer also allows for trivial (single-bit) carry computation in</span>
<a id="L54"></a><span class="comment">//    addition and subtraction (optimization suggested by Ken Thompson).</span>
<a id="L55"></a><span class="comment">//</span>
<a id="L56"></a><span class="comment">// 2) Long division requires extended-precision (2-digit) division per digit.</span>
<a id="L57"></a><span class="comment">//    Instead of sacrificing the largest base type for all other operations,</span>
<a id="L58"></a><span class="comment">//    for division the operands are unpacked into ``half-digits&#39;&#39;, and the</span>
<a id="L59"></a><span class="comment">//    results are packed again. For faster unpacking/packing, the base size</span>
<a id="L60"></a><span class="comment">//    in bits must be even.</span>

<a id="L62"></a>type (
    <a id="L63"></a>digit  uint64;
    <a id="L64"></a>digit2 uint32; <span class="comment">// half-digits for division</span>
<a id="L65"></a>)


<a id="L68"></a>const (
    <a id="L69"></a>logW = 64;          <span class="comment">// word width</span>
    <a id="L70"></a>logH = 4;           <span class="comment">// bits for a hex digit (= small number)</span>
    <a id="L71"></a>logB = logW - logH; <span class="comment">// largest bit-width available</span>

    <a id="L73"></a><span class="comment">// half-digits</span>
    <a id="L74"></a>_W2 = logB / 2; <span class="comment">// width</span>
    <a id="L75"></a>_B2 = 1 &lt;&lt; _W2; <span class="comment">// base</span>
    <a id="L76"></a>_M2 = _B2 - 1;  <span class="comment">// mask</span>

    <a id="L78"></a><span class="comment">// full digits</span>
    <a id="L79"></a>_W  = _W2 * 2; <span class="comment">// width</span>
    <a id="L80"></a>_B  = 1 &lt;&lt; _W; <span class="comment">// base</span>
    <a id="L81"></a>_M  = _B - 1;  <span class="comment">// mask</span>
<a id="L82"></a>)


<a id="L85"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L86"></a><span class="comment">// Support functions</span>

<a id="L88"></a>func assert(p bool) {
    <a id="L89"></a>if !p {
        <a id="L90"></a>panic(&#34;assert failed&#34;)
    <a id="L91"></a>}
<a id="L92"></a>}


<a id="L95"></a>func isSmall(x digit) bool { return x &lt; 1&lt;&lt;logH }


<a id="L98"></a><span class="comment">// For debugging. Keep around.</span>
<a id="L99"></a><span class="comment">/*</span>
<a id="L100"></a><span class="comment">func dump(x Natural) {</span>
<a id="L101"></a><span class="comment">	print(&#34;[&#34;, len(x), &#34;]&#34;);</span>
<a id="L102"></a><span class="comment">	for i := len(x) - 1; i &gt;= 0; i-- {</span>
<a id="L103"></a><span class="comment">		print(&#34; &#34;, x[i]);</span>
<a id="L104"></a><span class="comment">	}</span>
<a id="L105"></a><span class="comment">	println();</span>
<a id="L106"></a><span class="comment">}</span>
<a id="L107"></a><span class="comment">*/</span>


<a id="L110"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L111"></a><span class="comment">// Natural numbers</span>

<a id="L113"></a><span class="comment">// Natural represents an unsigned integer value of arbitrary precision.</span>
<a id="L114"></a><span class="comment">//</span>
<a id="L115"></a>type Natural []digit


<a id="L118"></a><span class="comment">// Nat creates a small natural number with value x.</span>
<a id="L119"></a><span class="comment">//</span>
<a id="L120"></a>func Nat(x uint64) Natural {
    <a id="L121"></a>if x == 0 {
        <a id="L122"></a>return nil <span class="comment">// len == 0</span>
    <a id="L123"></a>}

    <a id="L125"></a><span class="comment">// single-digit values</span>
    <a id="L126"></a><span class="comment">// (note: cannot re-use preallocated values because</span>
    <a id="L127"></a><span class="comment">//        the in-place operations may overwrite them)</span>
    <a id="L128"></a>if x &lt; _B {
        <a id="L129"></a>return Natural{digit(x)}
    <a id="L130"></a>}

    <a id="L132"></a><span class="comment">// compute number of digits required to represent x</span>
    <a id="L133"></a><span class="comment">// (this is usually 1 or 2, but the algorithm works</span>
    <a id="L134"></a><span class="comment">// for any base)</span>
    <a id="L135"></a>n := 0;
    <a id="L136"></a>for t := x; t &gt; 0; t &gt;&gt;= _W {
        <a id="L137"></a>n++
    <a id="L138"></a>}

    <a id="L140"></a><span class="comment">// split x into digits</span>
    <a id="L141"></a>z := make(Natural, n);
    <a id="L142"></a>for i := 0; i &lt; n; i++ {
        <a id="L143"></a>z[i] = digit(x &amp; _M);
        <a id="L144"></a>x &gt;&gt;= _W;
    <a id="L145"></a>}

    <a id="L147"></a>return z;
<a id="L148"></a>}


<a id="L151"></a><span class="comment">// Value returns the lowest 64bits of x.</span>
<a id="L152"></a><span class="comment">//</span>
<a id="L153"></a>func (x Natural) Value() uint64 {
    <a id="L154"></a><span class="comment">// single-digit values</span>
    <a id="L155"></a>n := len(x);
    <a id="L156"></a>switch n {
    <a id="L157"></a>case 0:
        <a id="L158"></a>return 0
    <a id="L159"></a>case 1:
        <a id="L160"></a>return uint64(x[0])
    <a id="L161"></a>}

    <a id="L163"></a><span class="comment">// multi-digit values</span>
    <a id="L164"></a><span class="comment">// (this is usually 1 or 2, but the algorithm works</span>
    <a id="L165"></a><span class="comment">// for any base)</span>
    <a id="L166"></a>z := uint64(0);
    <a id="L167"></a>s := uint(0);
    <a id="L168"></a>for i := 0; i &lt; n &amp;&amp; s &lt; 64; i++ {
        <a id="L169"></a>z += uint64(x[i]) &lt;&lt; s;
        <a id="L170"></a>s += _W;
    <a id="L171"></a>}

    <a id="L173"></a>return z;
<a id="L174"></a>}


<a id="L177"></a><span class="comment">// Predicates</span>

<a id="L179"></a><span class="comment">// IsEven returns true iff x is divisible by 2.</span>
<a id="L180"></a><span class="comment">//</span>
<a id="L181"></a>func (x Natural) IsEven() bool { return len(x) == 0 || x[0]&amp;1 == 0 }


<a id="L184"></a><span class="comment">// IsOdd returns true iff x is not divisible by 2.</span>
<a id="L185"></a><span class="comment">//</span>
<a id="L186"></a>func (x Natural) IsOdd() bool { return len(x) &gt; 0 &amp;&amp; x[0]&amp;1 != 0 }


<a id="L189"></a><span class="comment">// IsZero returns true iff x == 0.</span>
<a id="L190"></a><span class="comment">//</span>
<a id="L191"></a>func (x Natural) IsZero() bool { return len(x) == 0 }


<a id="L194"></a><span class="comment">// Operations</span>
<a id="L195"></a><span class="comment">//</span>
<a id="L196"></a><span class="comment">// Naming conventions</span>
<a id="L197"></a><span class="comment">//</span>
<a id="L198"></a><span class="comment">// c      carry</span>
<a id="L199"></a><span class="comment">// x, y   operands</span>
<a id="L200"></a><span class="comment">// z      result</span>
<a id="L201"></a><span class="comment">// n, m   len(x), len(y)</span>

<a id="L203"></a>func normalize(x Natural) Natural {
    <a id="L204"></a>n := len(x);
    <a id="L205"></a>for n &gt; 0 &amp;&amp; x[n-1] == 0 {
        <a id="L206"></a>n--
    <a id="L207"></a>}
    <a id="L208"></a>return x[0:n];
<a id="L209"></a>}


<a id="L212"></a><span class="comment">// nalloc returns a Natural of n digits. If z is large</span>
<a id="L213"></a><span class="comment">// enough, z is resized and returned. Otherwise, a new</span>
<a id="L214"></a><span class="comment">// Natural is allocated.</span>
<a id="L215"></a><span class="comment">//</span>
<a id="L216"></a>func nalloc(z Natural, n int) Natural {
    <a id="L217"></a>size := n;
    <a id="L218"></a>if size &lt;= 0 {
        <a id="L219"></a>size = 4
    <a id="L220"></a>}
    <a id="L221"></a>if size &lt;= cap(z) {
        <a id="L222"></a>return z[0:n]
    <a id="L223"></a>}
    <a id="L224"></a>return make(Natural, n, size);
<a id="L225"></a>}


<a id="L228"></a><span class="comment">// Nadd sets *zp to the sum x + y.</span>
<a id="L229"></a><span class="comment">// *zp may be x or y.</span>
<a id="L230"></a><span class="comment">//</span>
<a id="L231"></a>func Nadd(zp *Natural, x, y Natural) {
    <a id="L232"></a>n := len(x);
    <a id="L233"></a>m := len(y);
    <a id="L234"></a>if n &lt; m {
        <a id="L235"></a>Nadd(zp, y, x);
        <a id="L236"></a>return;
    <a id="L237"></a>}

    <a id="L239"></a>z := nalloc(*zp, n+1);
    <a id="L240"></a>c := digit(0);
    <a id="L241"></a>i := 0;
    <a id="L242"></a>for i &lt; m {
        <a id="L243"></a>t := c + x[i] + y[i];
        <a id="L244"></a>c, z[i] = t&gt;&gt;_W, t&amp;_M;
        <a id="L245"></a>i++;
    <a id="L246"></a>}
    <a id="L247"></a>for i &lt; n {
        <a id="L248"></a>t := c + x[i];
        <a id="L249"></a>c, z[i] = t&gt;&gt;_W, t&amp;_M;
        <a id="L250"></a>i++;
    <a id="L251"></a>}
    <a id="L252"></a>if c != 0 {
        <a id="L253"></a>z[i] = c;
        <a id="L254"></a>i++;
    <a id="L255"></a>}
    <a id="L256"></a>*zp = z[0:i];
<a id="L257"></a>}


<a id="L260"></a><span class="comment">// Add returns the sum z = x + y.</span>
<a id="L261"></a><span class="comment">//</span>
<a id="L262"></a>func (x Natural) Add(y Natural) Natural {
    <a id="L263"></a>var z Natural;
    <a id="L264"></a>Nadd(&amp;z, x, y);
    <a id="L265"></a>return z;
<a id="L266"></a>}


<a id="L269"></a><span class="comment">// Nsub sets *zp to the difference x - y for x &gt;= y.</span>
<a id="L270"></a><span class="comment">// If x &lt; y, an underflow run-time error occurs (use Cmp to test if x &gt;= y).</span>
<a id="L271"></a><span class="comment">// *zp may be x or y.</span>
<a id="L272"></a><span class="comment">//</span>
<a id="L273"></a>func Nsub(zp *Natural, x, y Natural) {
    <a id="L274"></a>n := len(x);
    <a id="L275"></a>m := len(y);
    <a id="L276"></a>if n &lt; m {
        <a id="L277"></a>panic(&#34;underflow&#34;)
    <a id="L278"></a>}

    <a id="L280"></a>z := nalloc(*zp, n);
    <a id="L281"></a>c := digit(0);
    <a id="L282"></a>i := 0;
    <a id="L283"></a>for i &lt; m {
        <a id="L284"></a>t := c + x[i] - y[i];
        <a id="L285"></a>c, z[i] = digit(int64(t)&gt;&gt;_W), t&amp;_M; <span class="comment">// requires arithmetic shift!</span>
        <a id="L286"></a>i++;
    <a id="L287"></a>}
    <a id="L288"></a>for i &lt; n {
        <a id="L289"></a>t := c + x[i];
        <a id="L290"></a>c, z[i] = digit(int64(t)&gt;&gt;_W), t&amp;_M; <span class="comment">// requires arithmetic shift!</span>
        <a id="L291"></a>i++;
    <a id="L292"></a>}
    <a id="L293"></a>if int64(c) &lt; 0 {
        <a id="L294"></a>panic(&#34;underflow&#34;)
    <a id="L295"></a>}
    <a id="L296"></a>*zp = normalize(z);
<a id="L297"></a>}


<a id="L300"></a><span class="comment">// Sub returns the difference x - y for x &gt;= y.</span>
<a id="L301"></a><span class="comment">// If x &lt; y, an underflow run-time error occurs (use Cmp to test if x &gt;= y).</span>
<a id="L302"></a><span class="comment">//</span>
<a id="L303"></a>func (x Natural) Sub(y Natural) Natural {
    <a id="L304"></a>var z Natural;
    <a id="L305"></a>Nsub(&amp;z, x, y);
    <a id="L306"></a>return z;
<a id="L307"></a>}


<a id="L310"></a><span class="comment">// Returns z1 = (x*y + c) div B, z0 = (x*y + c) mod B.</span>
<a id="L311"></a><span class="comment">//</span>
<a id="L312"></a>func muladd11(x, y, c digit) (digit, digit) {
    <a id="L313"></a>z1, z0 := MulAdd128(uint64(x), uint64(y), uint64(c));
    <a id="L314"></a>return digit(z1&lt;&lt;(64-logB) | z0&gt;&gt;logB), digit(z0 &amp; _M);
<a id="L315"></a>}


<a id="L318"></a>func mul1(z, x Natural, y digit) (c digit) {
    <a id="L319"></a>for i := 0; i &lt; len(x); i++ {
        <a id="L320"></a>c, z[i] = muladd11(x[i], y, c)
    <a id="L321"></a>}
    <a id="L322"></a>return;
<a id="L323"></a>}


<a id="L326"></a><span class="comment">// Nscale sets *z to the scaled value (*z) * d.</span>
<a id="L327"></a><span class="comment">//</span>
<a id="L328"></a>func Nscale(z *Natural, d uint64) {
    <a id="L329"></a>switch {
    <a id="L330"></a>case d == 0:
        <a id="L331"></a>*z = Nat(0);
        <a id="L332"></a>return;
    <a id="L333"></a>case d == 1:
        <a id="L334"></a>return
    <a id="L335"></a>case d &gt;= _B:
        <a id="L336"></a>*z = z.Mul1(d);
        <a id="L337"></a>return;
    <a id="L338"></a>}

    <a id="L340"></a>c := mul1(*z, *z, digit(d));

    <a id="L342"></a>if c != 0 {
        <a id="L343"></a>n := len(*z);
        <a id="L344"></a>if n &gt;= cap(*z) {
            <a id="L345"></a>zz := make(Natural, n+1);
            <a id="L346"></a>for i, d := range *z {
                <a id="L347"></a>zz[i] = d
            <a id="L348"></a>}
            <a id="L349"></a>*z = zz;
        <a id="L350"></a>} else {
            <a id="L351"></a>*z = (*z)[0 : n+1]
        <a id="L352"></a>}
        <a id="L353"></a>(*z)[n] = c;
    <a id="L354"></a>}
<a id="L355"></a>}


<a id="L358"></a><span class="comment">// Computes x = x*d + c for small d&#39;s.</span>
<a id="L359"></a><span class="comment">//</span>
<a id="L360"></a>func muladd1(x Natural, d, c digit) Natural {
    <a id="L361"></a>assert(isSmall(d-1) &amp;&amp; isSmall(c));
    <a id="L362"></a>n := len(x);
    <a id="L363"></a>z := make(Natural, n+1);

    <a id="L365"></a>for i := 0; i &lt; n; i++ {
        <a id="L366"></a>t := c + x[i]*d;
        <a id="L367"></a>c, z[i] = t&gt;&gt;_W, t&amp;_M;
    <a id="L368"></a>}
    <a id="L369"></a>z[n] = c;

    <a id="L371"></a>return normalize(z);
<a id="L372"></a>}


<a id="L375"></a><span class="comment">// Mul1 returns the product x * d.</span>
<a id="L376"></a><span class="comment">//</span>
<a id="L377"></a>func (x Natural) Mul1(d uint64) Natural {
    <a id="L378"></a>switch {
    <a id="L379"></a>case d == 0:
        <a id="L380"></a>return Nat(0)
    <a id="L381"></a>case d == 1:
        <a id="L382"></a>return x
    <a id="L383"></a>case isSmall(digit(d)):
        <a id="L384"></a>muladd1(x, digit(d), 0)
    <a id="L385"></a>case d &gt;= _B:
        <a id="L386"></a>return x.Mul(Nat(d))
    <a id="L387"></a>}

    <a id="L389"></a>z := make(Natural, len(x)+1);
    <a id="L390"></a>c := mul1(z, x, digit(d));
    <a id="L391"></a>z[len(x)] = c;
    <a id="L392"></a>return normalize(z);
<a id="L393"></a>}


<a id="L396"></a><span class="comment">// Mul returns the product x * y.</span>
<a id="L397"></a><span class="comment">//</span>
<a id="L398"></a>func (x Natural) Mul(y Natural) Natural {
    <a id="L399"></a>n := len(x);
    <a id="L400"></a>m := len(y);
    <a id="L401"></a>if n &lt; m {
        <a id="L402"></a>return y.Mul(x)
    <a id="L403"></a>}

    <a id="L405"></a>if m == 0 {
        <a id="L406"></a>return Nat(0)
    <a id="L407"></a>}

    <a id="L409"></a>if m == 1 &amp;&amp; y[0] &lt; _B {
        <a id="L410"></a>return x.Mul1(uint64(y[0]))
    <a id="L411"></a>}

    <a id="L413"></a>z := make(Natural, n+m);
    <a id="L414"></a>for j := 0; j &lt; m; j++ {
        <a id="L415"></a>d := y[j];
        <a id="L416"></a>if d != 0 {
            <a id="L417"></a>c := digit(0);
            <a id="L418"></a>for i := 0; i &lt; n; i++ {
                <a id="L419"></a>c, z[i+j] = muladd11(x[i], d, z[i+j]+c)
            <a id="L420"></a>}
            <a id="L421"></a>z[n+j] = c;
        <a id="L422"></a>}
    <a id="L423"></a>}

    <a id="L425"></a>return normalize(z);
<a id="L426"></a>}


<a id="L429"></a><span class="comment">// DivMod needs multi-precision division, which is not available if digit</span>
<a id="L430"></a><span class="comment">// is already using the largest uint size. Instead, unpack each operand</span>
<a id="L431"></a><span class="comment">// into operands with twice as many digits of half the size (digit2), do</span>
<a id="L432"></a><span class="comment">// DivMod, and then pack the results again.</span>

<a id="L434"></a>func unpack(x Natural) []digit2 {
    <a id="L435"></a>n := len(x);
    <a id="L436"></a>z := make([]digit2, n*2+1); <span class="comment">// add space for extra digit (used by DivMod)</span>
    <a id="L437"></a>for i := 0; i &lt; n; i++ {
        <a id="L438"></a>t := x[i];
        <a id="L439"></a>z[i*2] = digit2(t &amp; _M2);
        <a id="L440"></a>z[i*2+1] = digit2(t &gt;&gt; _W2 &amp; _M2);
    <a id="L441"></a>}

    <a id="L443"></a><span class="comment">// normalize result</span>
    <a id="L444"></a>k := 2 * n;
    <a id="L445"></a>for k &gt; 0 &amp;&amp; z[k-1] == 0 {
        <a id="L446"></a>k--
    <a id="L447"></a>}
    <a id="L448"></a>return z[0:k]; <span class="comment">// trim leading 0&#39;s</span>
<a id="L449"></a>}


<a id="L452"></a>func pack(x []digit2) Natural {
    <a id="L453"></a>n := (len(x) + 1) / 2;
    <a id="L454"></a>z := make(Natural, n);
    <a id="L455"></a>if len(x)&amp;1 == 1 {
        <a id="L456"></a><span class="comment">// handle odd len(x)</span>
        <a id="L457"></a>n--;
        <a id="L458"></a>z[n] = digit(x[n*2]);
    <a id="L459"></a>}
    <a id="L460"></a>for i := 0; i &lt; n; i++ {
        <a id="L461"></a>z[i] = digit(x[i*2+1])&lt;&lt;_W2 | digit(x[i*2])
    <a id="L462"></a>}
    <a id="L463"></a>return normalize(z);
<a id="L464"></a>}


<a id="L467"></a>func mul21(z, x []digit2, y digit2) digit2 {
    <a id="L468"></a>c := digit(0);
    <a id="L469"></a>f := digit(y);
    <a id="L470"></a>for i := 0; i &lt; len(x); i++ {
        <a id="L471"></a>t := c + digit(x[i])*f;
        <a id="L472"></a>c, z[i] = t&gt;&gt;_W2, digit2(t&amp;_M2);
    <a id="L473"></a>}
    <a id="L474"></a>return digit2(c);
<a id="L475"></a>}


<a id="L478"></a>func div21(z, x []digit2, y digit2) digit2 {
    <a id="L479"></a>c := digit(0);
    <a id="L480"></a>d := digit(y);
    <a id="L481"></a>for i := len(x) - 1; i &gt;= 0; i-- {
        <a id="L482"></a>t := c&lt;&lt;_W2 + digit(x[i]);
        <a id="L483"></a>c, z[i] = t%d, digit2(t/d);
    <a id="L484"></a>}
    <a id="L485"></a>return digit2(c);
<a id="L486"></a>}


<a id="L489"></a><span class="comment">// divmod returns q and r with x = y*q + r and 0 &lt;= r &lt; y.</span>
<a id="L490"></a><span class="comment">// x and y are destroyed in the process.</span>
<a id="L491"></a><span class="comment">//</span>
<a id="L492"></a><span class="comment">// The algorithm used here is based on 1). 2) describes the same algorithm</span>
<a id="L493"></a><span class="comment">// in C. A discussion and summary of the relevant theorems can be found in</span>
<a id="L494"></a><span class="comment">// 3). 3) also describes an easier way to obtain the trial digit - however</span>
<a id="L495"></a><span class="comment">// it relies on tripple-precision arithmetic which is why Knuth&#39;s method is</span>
<a id="L496"></a><span class="comment">// used here.</span>
<a id="L497"></a><span class="comment">//</span>
<a id="L498"></a><span class="comment">// 1) D. Knuth, The Art of Computer Programming. Volume 2. Seminumerical</span>
<a id="L499"></a><span class="comment">//    Algorithms. Addison-Wesley, Reading, 1969.</span>
<a id="L500"></a><span class="comment">//    (Algorithm D, Sec. 4.3.1)</span>
<a id="L501"></a><span class="comment">//</span>
<a id="L502"></a><span class="comment">// 2) Henry S. Warren, Jr., Hacker&#39;s Delight. Addison-Wesley, 2003.</span>
<a id="L503"></a><span class="comment">//    (9-2 Multiword Division, p.140ff)</span>
<a id="L504"></a><span class="comment">//</span>
<a id="L505"></a><span class="comment">// 3) P. Brinch Hansen, ``Multiple-length division revisited: A tour of the</span>
<a id="L506"></a><span class="comment">//    minefield&#39;&#39;. Software - Practice and Experience 24, (June 1994),</span>
<a id="L507"></a><span class="comment">//    579-601. John Wiley &amp; Sons, Ltd.</span>

<a id="L509"></a>func divmod(x, y []digit2) ([]digit2, []digit2) {
    <a id="L510"></a>n := len(x);
    <a id="L511"></a>m := len(y);
    <a id="L512"></a>if m == 0 {
        <a id="L513"></a>panic(&#34;division by zero&#34;)
    <a id="L514"></a>}
    <a id="L515"></a>assert(n+1 &lt;= cap(x)); <span class="comment">// space for one extra digit</span>
    <a id="L516"></a>x = x[0 : n+1];
    <a id="L517"></a>assert(x[n] == 0);

    <a id="L519"></a>if m == 1 {
        <a id="L520"></a><span class="comment">// division by single digit</span>
        <a id="L521"></a><span class="comment">// result is shifted left by 1 in place!</span>
        <a id="L522"></a>x[0] = div21(x[1:n+1], x[0:n], y[0])

    <a id="L524"></a>} else if m &gt; n {
        <a id="L525"></a><span class="comment">// y &gt; x =&gt; quotient = 0, remainder = x</span>
        <a id="L526"></a><span class="comment">// TODO in this case we shouldn&#39;t even unpack x and y</span>
        <a id="L527"></a>m = n

    <a id="L529"></a>} else {
        <a id="L530"></a><span class="comment">// general case</span>
        <a id="L531"></a>assert(2 &lt;= m &amp;&amp; m &lt;= n);

        <a id="L533"></a><span class="comment">// normalize x and y</span>
        <a id="L534"></a><span class="comment">// TODO Instead of multiplying, it would be sufficient to</span>
        <a id="L535"></a><span class="comment">//      shift y such that the normalization condition is</span>
        <a id="L536"></a><span class="comment">//      satisfied (as done in Hacker&#39;s Delight).</span>
        <a id="L537"></a>f := _B2 / (digit(y[m-1]) + 1);
        <a id="L538"></a>if f != 1 {
            <a id="L539"></a>mul21(x, x, digit2(f));
            <a id="L540"></a>mul21(y, y, digit2(f));
        <a id="L541"></a>}
        <a id="L542"></a>assert(_B2/2 &lt;= y[m-1] &amp;&amp; y[m-1] &lt; _B2); <span class="comment">// incorrect scaling</span>

        <a id="L544"></a>y1, y2 := digit(y[m-1]), digit(y[m-2]);
        <a id="L545"></a>for i := n - m; i &gt;= 0; i-- {
            <a id="L546"></a>k := i + m;

            <a id="L548"></a><span class="comment">// compute trial digit (Knuth)</span>
            <a id="L549"></a>var q digit;
            <a id="L550"></a>{
                <a id="L551"></a>x0, x1, x2 := digit(x[k]), digit(x[k-1]), digit(x[k-2]);
                <a id="L552"></a>if x0 != y1 {
                    <a id="L553"></a>q = (x0&lt;&lt;_W2 + x1) / y1
                <a id="L554"></a>} else {
                    <a id="L555"></a>q = _B2 - 1
                <a id="L556"></a>}
                <a id="L557"></a>for y2*q &gt; (x0&lt;&lt;_W2+x1-y1*q)&lt;&lt;_W2+x2 {
                    <a id="L558"></a>q--
                <a id="L559"></a>}
            <a id="L560"></a>}

            <a id="L562"></a><span class="comment">// subtract y*q</span>
            <a id="L563"></a>c := digit(0);
            <a id="L564"></a>for j := 0; j &lt; m; j++ {
                <a id="L565"></a>t := c + digit(x[i+j]) - digit(y[j])*q;
                <a id="L566"></a>c, x[i+j] = digit(int64(t)&gt;&gt;_W2), digit2(t&amp;_M2); <span class="comment">// requires arithmetic shift!</span>
            <a id="L567"></a>}

            <a id="L569"></a><span class="comment">// correct if trial digit was too large</span>
            <a id="L570"></a>if c+digit(x[k]) != 0 {
                <a id="L571"></a><span class="comment">// add y</span>
                <a id="L572"></a>c := digit(0);
                <a id="L573"></a>for j := 0; j &lt; m; j++ {
                    <a id="L574"></a>t := c + digit(x[i+j]) + digit(y[j]);
                    <a id="L575"></a>c, x[i+j] = t&gt;&gt;_W2, digit2(t&amp;_M2);
                <a id="L576"></a>}
                <a id="L577"></a>assert(c+digit(x[k]) == 0);
                <a id="L578"></a><span class="comment">// correct trial digit</span>
                <a id="L579"></a>q--;
            <a id="L580"></a>}

            <a id="L582"></a>x[k] = digit2(q);
        <a id="L583"></a>}

        <a id="L585"></a><span class="comment">// undo normalization for remainder</span>
        <a id="L586"></a>if f != 1 {
            <a id="L587"></a>c := div21(x[0:m], x[0:m], digit2(f));
            <a id="L588"></a>assert(c == 0);
        <a id="L589"></a>}
    <a id="L590"></a>}

    <a id="L592"></a>return x[m : n+1], x[0:m];
<a id="L593"></a>}


<a id="L596"></a><span class="comment">// Div returns the quotient q = x / y for y &gt; 0,</span>
<a id="L597"></a><span class="comment">// with x = y*q + r and 0 &lt;= r &lt; y.</span>
<a id="L598"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L599"></a><span class="comment">//</span>
<a id="L600"></a>func (x Natural) Div(y Natural) Natural {
    <a id="L601"></a>q, _ := divmod(unpack(x), unpack(y));
    <a id="L602"></a>return pack(q);
<a id="L603"></a>}


<a id="L606"></a><span class="comment">// Mod returns the modulus r of the division x / y for y &gt; 0,</span>
<a id="L607"></a><span class="comment">// with x = y*q + r and 0 &lt;= r &lt; y.</span>
<a id="L608"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L609"></a><span class="comment">//</span>
<a id="L610"></a>func (x Natural) Mod(y Natural) Natural {
    <a id="L611"></a>_, r := divmod(unpack(x), unpack(y));
    <a id="L612"></a>return pack(r);
<a id="L613"></a>}


<a id="L616"></a><span class="comment">// DivMod returns the pair (x.Div(y), x.Mod(y)) for y &gt; 0.</span>
<a id="L617"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L618"></a><span class="comment">//</span>
<a id="L619"></a>func (x Natural) DivMod(y Natural) (Natural, Natural) {
    <a id="L620"></a>q, r := divmod(unpack(x), unpack(y));
    <a id="L621"></a>return pack(q), pack(r);
<a id="L622"></a>}


<a id="L625"></a>func shl(z, x Natural, s uint) digit {
    <a id="L626"></a>assert(s &lt;= _W);
    <a id="L627"></a>n := len(x);
    <a id="L628"></a>c := digit(0);
    <a id="L629"></a>for i := 0; i &lt; n; i++ {
        <a id="L630"></a>c, z[i] = x[i]&gt;&gt;(_W-s), x[i]&lt;&lt;s&amp;_M|c
    <a id="L631"></a>}
    <a id="L632"></a>return c;
<a id="L633"></a>}


<a id="L636"></a><span class="comment">// Shl implements ``shift left&#39;&#39; x &lt;&lt; s. It returns x * 2^s.</span>
<a id="L637"></a><span class="comment">//</span>
<a id="L638"></a>func (x Natural) Shl(s uint) Natural {
    <a id="L639"></a>n := uint(len(x));
    <a id="L640"></a>m := n + s/_W;
    <a id="L641"></a>z := make(Natural, m+1);

    <a id="L643"></a>z[m] = shl(z[m-n:m], x, s%_W);

    <a id="L645"></a>return normalize(z);
<a id="L646"></a>}


<a id="L649"></a>func shr(z, x Natural, s uint) digit {
    <a id="L650"></a>assert(s &lt;= _W);
    <a id="L651"></a>n := len(x);
    <a id="L652"></a>c := digit(0);
    <a id="L653"></a>for i := n - 1; i &gt;= 0; i-- {
        <a id="L654"></a>c, z[i] = x[i]&lt;&lt;(_W-s)&amp;_M, x[i]&gt;&gt;s|c
    <a id="L655"></a>}
    <a id="L656"></a>return c;
<a id="L657"></a>}


<a id="L660"></a><span class="comment">// Shr implements ``shift right&#39;&#39; x &gt;&gt; s. It returns x / 2^s.</span>
<a id="L661"></a><span class="comment">//</span>
<a id="L662"></a>func (x Natural) Shr(s uint) Natural {
    <a id="L663"></a>n := uint(len(x));
    <a id="L664"></a>m := n - s/_W;
    <a id="L665"></a>if m &gt; n { <span class="comment">// check for underflow</span>
        <a id="L666"></a>m = 0
    <a id="L667"></a>}
    <a id="L668"></a>z := make(Natural, m);

    <a id="L670"></a>shr(z, x[n-m:n], s%_W);

    <a id="L672"></a>return normalize(z);
<a id="L673"></a>}


<a id="L676"></a><span class="comment">// And returns the ``bitwise and&#39;&#39; x &amp; y for the 2&#39;s-complement representation of x and y.</span>
<a id="L677"></a><span class="comment">//</span>
<a id="L678"></a>func (x Natural) And(y Natural) Natural {
    <a id="L679"></a>n := len(x);
    <a id="L680"></a>m := len(y);
    <a id="L681"></a>if n &lt; m {
        <a id="L682"></a>return y.And(x)
    <a id="L683"></a>}

    <a id="L685"></a>z := make(Natural, m);
    <a id="L686"></a>for i := 0; i &lt; m; i++ {
        <a id="L687"></a>z[i] = x[i] &amp; y[i]
    <a id="L688"></a>}
    <a id="L689"></a><span class="comment">// upper bits are 0</span>

    <a id="L691"></a>return normalize(z);
<a id="L692"></a>}


<a id="L695"></a>func copy(z, x Natural) {
    <a id="L696"></a>for i, e := range x {
        <a id="L697"></a>z[i] = e
    <a id="L698"></a>}
<a id="L699"></a>}


<a id="L702"></a><span class="comment">// AndNot returns the ``bitwise clear&#39;&#39; x &amp;^ y for the 2&#39;s-complement representation of x and y.</span>
<a id="L703"></a><span class="comment">//</span>
<a id="L704"></a>func (x Natural) AndNot(y Natural) Natural {
    <a id="L705"></a>n := len(x);
    <a id="L706"></a>m := len(y);
    <a id="L707"></a>if n &lt; m {
        <a id="L708"></a>m = n
    <a id="L709"></a>}

    <a id="L711"></a>z := make(Natural, n);
    <a id="L712"></a>for i := 0; i &lt; m; i++ {
        <a id="L713"></a>z[i] = x[i] &amp;^ y[i]
    <a id="L714"></a>}
    <a id="L715"></a>copy(z[m:n], x[m:n]);

    <a id="L717"></a>return normalize(z);
<a id="L718"></a>}


<a id="L721"></a><span class="comment">// Or returns the ``bitwise or&#39;&#39; x | y for the 2&#39;s-complement representation of x and y.</span>
<a id="L722"></a><span class="comment">//</span>
<a id="L723"></a>func (x Natural) Or(y Natural) Natural {
    <a id="L724"></a>n := len(x);
    <a id="L725"></a>m := len(y);
    <a id="L726"></a>if n &lt; m {
        <a id="L727"></a>return y.Or(x)
    <a id="L728"></a>}

    <a id="L730"></a>z := make(Natural, n);
    <a id="L731"></a>for i := 0; i &lt; m; i++ {
        <a id="L732"></a>z[i] = x[i] | y[i]
    <a id="L733"></a>}
    <a id="L734"></a>copy(z[m:n], x[m:n]);

    <a id="L736"></a>return z;
<a id="L737"></a>}


<a id="L740"></a><span class="comment">// Xor returns the ``bitwise exclusive or&#39;&#39; x ^ y for the 2&#39;s-complement representation of x and y.</span>
<a id="L741"></a><span class="comment">//</span>
<a id="L742"></a>func (x Natural) Xor(y Natural) Natural {
    <a id="L743"></a>n := len(x);
    <a id="L744"></a>m := len(y);
    <a id="L745"></a>if n &lt; m {
        <a id="L746"></a>return y.Xor(x)
    <a id="L747"></a>}

    <a id="L749"></a>z := make(Natural, n);
    <a id="L750"></a>for i := 0; i &lt; m; i++ {
        <a id="L751"></a>z[i] = x[i] ^ y[i]
    <a id="L752"></a>}
    <a id="L753"></a>copy(z[m:n], x[m:n]);

    <a id="L755"></a>return normalize(z);
<a id="L756"></a>}


<a id="L759"></a><span class="comment">// Cmp compares x and y. The result is an int value</span>
<a id="L760"></a><span class="comment">//</span>
<a id="L761"></a><span class="comment">//   &lt;  0 if x &lt;  y</span>
<a id="L762"></a><span class="comment">//   == 0 if x == y</span>
<a id="L763"></a><span class="comment">//   &gt;  0 if x &gt;  y</span>
<a id="L764"></a><span class="comment">//</span>
<a id="L765"></a>func (x Natural) Cmp(y Natural) int {
    <a id="L766"></a>n := len(x);
    <a id="L767"></a>m := len(y);

    <a id="L769"></a>if n != m || n == 0 {
        <a id="L770"></a>return n - m
    <a id="L771"></a>}

    <a id="L773"></a>i := n - 1;
    <a id="L774"></a>for i &gt; 0 &amp;&amp; x[i] == y[i] {
        <a id="L775"></a>i--
    <a id="L776"></a>}

    <a id="L778"></a>d := 0;
    <a id="L779"></a>switch {
    <a id="L780"></a>case x[i] &lt; y[i]:
        <a id="L781"></a>d = -1
    <a id="L782"></a>case x[i] &gt; y[i]:
        <a id="L783"></a>d = 1
    <a id="L784"></a>}

    <a id="L786"></a>return d;
<a id="L787"></a>}


<a id="L790"></a><span class="comment">// log2 computes the binary logarithm of x for x &gt; 0.</span>
<a id="L791"></a><span class="comment">// The result is the integer n for which 2^n &lt;= x &lt; 2^(n+1).</span>
<a id="L792"></a><span class="comment">// If x == 0 a run-time error occurs.</span>
<a id="L793"></a><span class="comment">//</span>
<a id="L794"></a>func log2(x uint64) uint {
    <a id="L795"></a>assert(x &gt; 0);
    <a id="L796"></a>n := uint(0);
    <a id="L797"></a>for x &gt; 0 {
        <a id="L798"></a>x &gt;&gt;= 1;
        <a id="L799"></a>n++;
    <a id="L800"></a>}
    <a id="L801"></a>return n - 1;
<a id="L802"></a>}


<a id="L805"></a><span class="comment">// Log2 computes the binary logarithm of x for x &gt; 0.</span>
<a id="L806"></a><span class="comment">// The result is the integer n for which 2^n &lt;= x &lt; 2^(n+1).</span>
<a id="L807"></a><span class="comment">// If x == 0 a run-time error occurs.</span>
<a id="L808"></a><span class="comment">//</span>
<a id="L809"></a>func (x Natural) Log2() uint {
    <a id="L810"></a>n := len(x);
    <a id="L811"></a>if n &gt; 0 {
        <a id="L812"></a>return (uint(n)-1)*_W + log2(uint64(x[n-1]))
    <a id="L813"></a>}
    <a id="L814"></a>panic(&#34;Log2(0)&#34;);
<a id="L815"></a>}


<a id="L818"></a><span class="comment">// Computes x = x div d in place (modifies x) for small d&#39;s.</span>
<a id="L819"></a><span class="comment">// Returns updated x and x mod d.</span>
<a id="L820"></a><span class="comment">//</span>
<a id="L821"></a>func divmod1(x Natural, d digit) (Natural, digit) {
    <a id="L822"></a>assert(0 &lt; d &amp;&amp; isSmall(d-1));

    <a id="L824"></a>c := digit(0);
    <a id="L825"></a>for i := len(x) - 1; i &gt;= 0; i-- {
        <a id="L826"></a>t := c&lt;&lt;_W + x[i];
        <a id="L827"></a>c, x[i] = t%d, t/d;
    <a id="L828"></a>}

    <a id="L830"></a>return normalize(x), c;
<a id="L831"></a>}


<a id="L834"></a><span class="comment">// ToString converts x to a string for a given base, with 2 &lt;= base &lt;= 16.</span>
<a id="L835"></a><span class="comment">//</span>
<a id="L836"></a>func (x Natural) ToString(base uint) string {
    <a id="L837"></a>if len(x) == 0 {
        <a id="L838"></a>return &#34;0&#34;
    <a id="L839"></a>}

    <a id="L841"></a><span class="comment">// allocate buffer for conversion</span>
    <a id="L842"></a>assert(2 &lt;= base &amp;&amp; base &lt;= 16);
    <a id="L843"></a>n := (x.Log2()+1)/log2(uint64(base)) + 1; <span class="comment">// +1: round up</span>
    <a id="L844"></a>s := make([]byte, n);

    <a id="L846"></a><span class="comment">// don&#39;t destroy x</span>
    <a id="L847"></a>t := make(Natural, len(x));
    <a id="L848"></a>copy(t, x);

    <a id="L850"></a><span class="comment">// convert</span>
    <a id="L851"></a>i := n;
    <a id="L852"></a>for !t.IsZero() {
        <a id="L853"></a>i--;
        <a id="L854"></a>var d digit;
        <a id="L855"></a>t, d = divmod1(t, digit(base));
        <a id="L856"></a>s[i] = &#34;0123456789abcdef&#34;[d];
    <a id="L857"></a>}

    <a id="L859"></a>return string(s[i:n]);
<a id="L860"></a>}


<a id="L863"></a><span class="comment">// String converts x to its decimal string representation.</span>
<a id="L864"></a><span class="comment">// x.String() is the same as x.ToString(10).</span>
<a id="L865"></a><span class="comment">//</span>
<a id="L866"></a>func (x Natural) String() string { return x.ToString(10) }


<a id="L869"></a>func fmtbase(c int) uint {
    <a id="L870"></a>switch c {
    <a id="L871"></a>case &#39;b&#39;:
        <a id="L872"></a>return 2
    <a id="L873"></a>case &#39;o&#39;:
        <a id="L874"></a>return 8
    <a id="L875"></a>case &#39;x&#39;:
        <a id="L876"></a>return 16
    <a id="L877"></a>}
    <a id="L878"></a>return 10;
<a id="L879"></a>}


<a id="L882"></a><span class="comment">// Format is a support routine for fmt.Formatter. It accepts</span>
<a id="L883"></a><span class="comment">// the formats &#39;b&#39; (binary), &#39;o&#39; (octal), and &#39;x&#39; (hexadecimal).</span>
<a id="L884"></a><span class="comment">//</span>
<a id="L885"></a>func (x Natural) Format(h fmt.State, c int) { fmt.Fprintf(h, &#34;%s&#34;, x.ToString(fmtbase(c))) }


<a id="L888"></a>func hexvalue(ch byte) uint {
    <a id="L889"></a>d := uint(1 &lt;&lt; logH);
    <a id="L890"></a>switch {
    <a id="L891"></a>case &#39;0&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;9&#39;:
        <a id="L892"></a>d = uint(ch - &#39;0&#39;)
    <a id="L893"></a>case &#39;a&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;f&#39;:
        <a id="L894"></a>d = uint(ch-&#39;a&#39;) + 10
    <a id="L895"></a>case &#39;A&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;F&#39;:
        <a id="L896"></a>d = uint(ch-&#39;A&#39;) + 10
    <a id="L897"></a>}
    <a id="L898"></a>return d;
<a id="L899"></a>}


<a id="L902"></a><span class="comment">// NatFromString returns the natural number corresponding to the</span>
<a id="L903"></a><span class="comment">// longest possible prefix of s representing a natural number in a</span>
<a id="L904"></a><span class="comment">// given conversion base, the actual conversion base used, and the</span>
<a id="L905"></a><span class="comment">// prefix length. The syntax of natural numbers follows the syntax</span>
<a id="L906"></a><span class="comment">// of unsigned integer literals in Go.</span>
<a id="L907"></a><span class="comment">//</span>
<a id="L908"></a><span class="comment">// If the base argument is 0, the string prefix determines the actual</span>
<a id="L909"></a><span class="comment">// conversion base. A prefix of ``0x&#39;&#39; or ``0X&#39;&#39; selects base 16; the</span>
<a id="L910"></a><span class="comment">// ``0&#39;&#39; prefix selects base 8. Otherwise the selected base is 10.</span>
<a id="L911"></a><span class="comment">//</span>
<a id="L912"></a>func NatFromString(s string, base uint) (Natural, uint, int) {
    <a id="L913"></a><span class="comment">// determine base if necessary</span>
    <a id="L914"></a>i, n := 0, len(s);
    <a id="L915"></a>if base == 0 {
        <a id="L916"></a>base = 10;
        <a id="L917"></a>if n &gt; 0 &amp;&amp; s[0] == &#39;0&#39; {
            <a id="L918"></a>if n &gt; 1 &amp;&amp; (s[1] == &#39;x&#39; || s[1] == &#39;X&#39;) {
                <a id="L919"></a>base, i = 16, 2
            <a id="L920"></a>} else {
                <a id="L921"></a>base, i = 8, 1
            <a id="L922"></a>}
        <a id="L923"></a>}
    <a id="L924"></a>}

    <a id="L926"></a><span class="comment">// convert string</span>
    <a id="L927"></a>assert(2 &lt;= base &amp;&amp; base &lt;= 16);
    <a id="L928"></a>x := Nat(0);
    <a id="L929"></a>for ; i &lt; n; i++ {
        <a id="L930"></a>d := hexvalue(s[i]);
        <a id="L931"></a>if d &lt; base {
            <a id="L932"></a>x = muladd1(x, digit(base), digit(d))
        <a id="L933"></a>} else {
            <a id="L934"></a>break
        <a id="L935"></a>}
    <a id="L936"></a>}

    <a id="L938"></a>return x, base, i;
<a id="L939"></a>}


<a id="L942"></a><span class="comment">// Natural number functions</span>

<a id="L944"></a>func pop1(x digit) uint {
    <a id="L945"></a>n := uint(0);
    <a id="L946"></a>for x != 0 {
        <a id="L947"></a>x &amp;= x - 1;
        <a id="L948"></a>n++;
    <a id="L949"></a>}
    <a id="L950"></a>return n;
<a id="L951"></a>}


<a id="L954"></a><span class="comment">// Pop computes the ``population count&#39;&#39; of (the number of 1 bits in) x.</span>
<a id="L955"></a><span class="comment">//</span>
<a id="L956"></a>func (x Natural) Pop() uint {
    <a id="L957"></a>n := uint(0);
    <a id="L958"></a>for i := len(x) - 1; i &gt;= 0; i-- {
        <a id="L959"></a>n += pop1(x[i])
    <a id="L960"></a>}
    <a id="L961"></a>return n;
<a id="L962"></a>}


<a id="L965"></a><span class="comment">// Pow computes x to the power of n.</span>
<a id="L966"></a><span class="comment">//</span>
<a id="L967"></a>func (xp Natural) Pow(n uint) Natural {
    <a id="L968"></a>z := Nat(1);
    <a id="L969"></a>x := xp;
    <a id="L970"></a>for n &gt; 0 {
        <a id="L971"></a><span class="comment">// z * x^n == x^n0</span>
        <a id="L972"></a>if n&amp;1 == 1 {
            <a id="L973"></a>z = z.Mul(x)
        <a id="L974"></a>}
        <a id="L975"></a>x, n = x.Mul(x), n/2;
    <a id="L976"></a>}
    <a id="L977"></a>return z;
<a id="L978"></a>}


<a id="L981"></a><span class="comment">// MulRange computes the product of all the unsigned integers</span>
<a id="L982"></a><span class="comment">// in the range [a, b] inclusively.</span>
<a id="L983"></a><span class="comment">//</span>
<a id="L984"></a>func MulRange(a, b uint) Natural {
    <a id="L985"></a>switch {
    <a id="L986"></a>case a &gt; b:
        <a id="L987"></a>return Nat(1)
    <a id="L988"></a>case a == b:
        <a id="L989"></a>return Nat(uint64(a))
    <a id="L990"></a>case a+1 == b:
        <a id="L991"></a>return Nat(uint64(a)).Mul(Nat(uint64(b)))
    <a id="L992"></a>}
    <a id="L993"></a>m := (a + b) &gt;&gt; 1;
    <a id="L994"></a>assert(a &lt;= m &amp;&amp; m &lt; b);
    <a id="L995"></a>return MulRange(a, m).Mul(MulRange(m+1, b));
<a id="L996"></a>}


<a id="L999"></a><span class="comment">// Fact computes the factorial of n (Fact(n) == MulRange(2, n)).</span>
<a id="L1000"></a><span class="comment">//</span>
<a id="L1001"></a>func Fact(n uint) Natural {
    <a id="L1002"></a><span class="comment">// Using MulRange() instead of the basic for-loop</span>
    <a id="L1003"></a><span class="comment">// lead to faster factorial computation.</span>
    <a id="L1004"></a>return MulRange(2, n)
<a id="L1005"></a>}


<a id="L1008"></a><span class="comment">// Binomial computes the binomial coefficient of (n, k).</span>
<a id="L1009"></a><span class="comment">//</span>
<a id="L1010"></a>func Binomial(n, k uint) Natural { return MulRange(n-k+1, n).Div(MulRange(1, k)) }


<a id="L1013"></a><span class="comment">// Gcd computes the gcd of x and y.</span>
<a id="L1014"></a><span class="comment">//</span>
<a id="L1015"></a>func (x Natural) Gcd(y Natural) Natural {
    <a id="L1016"></a><span class="comment">// Euclidean algorithm.</span>
    <a id="L1017"></a>a, b := x, y;
    <a id="L1018"></a>for !b.IsZero() {
        <a id="L1019"></a>a, b = b, a.Mod(b)
    <a id="L1020"></a>}
    <a id="L1021"></a>return a;
<a id="L1022"></a>}
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
