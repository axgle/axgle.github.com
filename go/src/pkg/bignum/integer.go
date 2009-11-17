<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bignum/integer.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/bignum/integer.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Integer numbers</span>
<a id="L6"></a><span class="comment">//</span>
<a id="L7"></a><span class="comment">// Integers are normalized if the mantissa is normalized and the sign is</span>
<a id="L8"></a><span class="comment">// false for mant == 0. Use MakeInt to create normalized Integers.</span>

<a id="L10"></a>package bignum

<a id="L12"></a>import (
    <a id="L13"></a>&#34;fmt&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// TODO(gri) Complete the set of in-place operations.</span>

<a id="L18"></a><span class="comment">// Integer represents a signed integer value of arbitrary precision.</span>
<a id="L19"></a><span class="comment">//</span>
<a id="L20"></a>type Integer struct {
    <a id="L21"></a>sign bool;
    <a id="L22"></a>mant Natural;
<a id="L23"></a>}


<a id="L26"></a><span class="comment">// MakeInt makes an integer given a sign and a mantissa.</span>
<a id="L27"></a><span class="comment">// The number is positive (&gt;= 0) if sign is false or the</span>
<a id="L28"></a><span class="comment">// mantissa is zero; it is negative otherwise.</span>
<a id="L29"></a><span class="comment">//</span>
<a id="L30"></a>func MakeInt(sign bool, mant Natural) *Integer {
    <a id="L31"></a>if mant.IsZero() {
        <a id="L32"></a>sign = false <span class="comment">// normalize</span>
    <a id="L33"></a>}
    <a id="L34"></a>return &amp;Integer{sign, mant};
<a id="L35"></a>}


<a id="L38"></a><span class="comment">// Int creates a small integer with value x.</span>
<a id="L39"></a><span class="comment">//</span>
<a id="L40"></a>func Int(x int64) *Integer {
    <a id="L41"></a>var ux uint64;
    <a id="L42"></a>if x &lt; 0 {
        <a id="L43"></a><span class="comment">// For the most negative x, -x == x, and</span>
        <a id="L44"></a><span class="comment">// the bit pattern has the correct value.</span>
        <a id="L45"></a>ux = uint64(-x)
    <a id="L46"></a>} else {
        <a id="L47"></a>ux = uint64(x)
    <a id="L48"></a>}
    <a id="L49"></a>return MakeInt(x &lt; 0, Nat(ux));
<a id="L50"></a>}


<a id="L53"></a><span class="comment">// Value returns the value of x, if x fits into an int64;</span>
<a id="L54"></a><span class="comment">// otherwise the result is undefined.</span>
<a id="L55"></a><span class="comment">//</span>
<a id="L56"></a>func (x *Integer) Value() int64 {
    <a id="L57"></a>z := int64(x.mant.Value());
    <a id="L58"></a>if x.sign {
        <a id="L59"></a>z = -z
    <a id="L60"></a>}
    <a id="L61"></a>return z;
<a id="L62"></a>}


<a id="L65"></a><span class="comment">// Abs returns the absolute value of x.</span>
<a id="L66"></a><span class="comment">//</span>
<a id="L67"></a>func (x *Integer) Abs() Natural { return x.mant }


<a id="L70"></a><span class="comment">// Predicates</span>

<a id="L72"></a><span class="comment">// IsEven returns true iff x is divisible by 2.</span>
<a id="L73"></a><span class="comment">//</span>
<a id="L74"></a>func (x *Integer) IsEven() bool { return x.mant.IsEven() }


<a id="L77"></a><span class="comment">// IsOdd returns true iff x is not divisible by 2.</span>
<a id="L78"></a><span class="comment">//</span>
<a id="L79"></a>func (x *Integer) IsOdd() bool { return x.mant.IsOdd() }


<a id="L82"></a><span class="comment">// IsZero returns true iff x == 0.</span>
<a id="L83"></a><span class="comment">//</span>
<a id="L84"></a>func (x *Integer) IsZero() bool { return x.mant.IsZero() }


<a id="L87"></a><span class="comment">// IsNeg returns true iff x &lt; 0.</span>
<a id="L88"></a><span class="comment">//</span>
<a id="L89"></a>func (x *Integer) IsNeg() bool { return x.sign &amp;&amp; !x.mant.IsZero() }


<a id="L92"></a><span class="comment">// IsPos returns true iff x &gt;= 0.</span>
<a id="L93"></a><span class="comment">//</span>
<a id="L94"></a>func (x *Integer) IsPos() bool { return !x.sign &amp;&amp; !x.mant.IsZero() }


<a id="L97"></a><span class="comment">// Operations</span>

<a id="L99"></a><span class="comment">// Neg returns the negated value of x.</span>
<a id="L100"></a><span class="comment">//</span>
<a id="L101"></a>func (x *Integer) Neg() *Integer { return MakeInt(!x.sign, x.mant) }


<a id="L104"></a><span class="comment">// Iadd sets z to the sum x + y.</span>
<a id="L105"></a><span class="comment">// z must exist and may be x or y.</span>
<a id="L106"></a><span class="comment">//</span>
<a id="L107"></a>func Iadd(z, x, y *Integer) {
    <a id="L108"></a>if x.sign == y.sign {
        <a id="L109"></a><span class="comment">// x + y == x + y</span>
        <a id="L110"></a><span class="comment">// (-x) + (-y) == -(x + y)</span>
        <a id="L111"></a>z.sign = x.sign;
        <a id="L112"></a>Nadd(&amp;z.mant, x.mant, y.mant);
    <a id="L113"></a>} else {
        <a id="L114"></a><span class="comment">// x + (-y) == x - y == -(y - x)</span>
        <a id="L115"></a><span class="comment">// (-x) + y == y - x == -(x - y)</span>
        <a id="L116"></a>if x.mant.Cmp(y.mant) &gt;= 0 {
            <a id="L117"></a>z.sign = x.sign;
            <a id="L118"></a>Nsub(&amp;z.mant, x.mant, y.mant);
        <a id="L119"></a>} else {
            <a id="L120"></a>z.sign = !x.sign;
            <a id="L121"></a>Nsub(&amp;z.mant, y.mant, x.mant);
        <a id="L122"></a>}
    <a id="L123"></a>}
<a id="L124"></a>}


<a id="L127"></a><span class="comment">// Add returns the sum x + y.</span>
<a id="L128"></a><span class="comment">//</span>
<a id="L129"></a>func (x *Integer) Add(y *Integer) *Integer {
    <a id="L130"></a>var z Integer;
    <a id="L131"></a>Iadd(&amp;z, x, y);
    <a id="L132"></a>return &amp;z;
<a id="L133"></a>}


<a id="L136"></a>func Isub(z, x, y *Integer) {
    <a id="L137"></a>if x.sign != y.sign {
        <a id="L138"></a><span class="comment">// x - (-y) == x + y</span>
        <a id="L139"></a><span class="comment">// (-x) - y == -(x + y)</span>
        <a id="L140"></a>z.sign = x.sign;
        <a id="L141"></a>Nadd(&amp;z.mant, x.mant, y.mant);
    <a id="L142"></a>} else {
        <a id="L143"></a><span class="comment">// x - y == x - y == -(y - x)</span>
        <a id="L144"></a><span class="comment">// (-x) - (-y) == y - x == -(x - y)</span>
        <a id="L145"></a>if x.mant.Cmp(y.mant) &gt;= 0 {
            <a id="L146"></a>z.sign = x.sign;
            <a id="L147"></a>Nsub(&amp;z.mant, x.mant, y.mant);
        <a id="L148"></a>} else {
            <a id="L149"></a>z.sign = !x.sign;
            <a id="L150"></a>Nsub(&amp;z.mant, y.mant, x.mant);
        <a id="L151"></a>}
    <a id="L152"></a>}
<a id="L153"></a>}


<a id="L156"></a><span class="comment">// Sub returns the difference x - y.</span>
<a id="L157"></a><span class="comment">//</span>
<a id="L158"></a>func (x *Integer) Sub(y *Integer) *Integer {
    <a id="L159"></a>var z Integer;
    <a id="L160"></a>Isub(&amp;z, x, y);
    <a id="L161"></a>return &amp;z;
<a id="L162"></a>}


<a id="L165"></a><span class="comment">// Nscale sets *z to the scaled value (*z) * d.</span>
<a id="L166"></a><span class="comment">//</span>
<a id="L167"></a>func Iscale(z *Integer, d int64) {
    <a id="L168"></a>f := uint64(d);
    <a id="L169"></a>if d &lt; 0 {
        <a id="L170"></a>f = uint64(-d)
    <a id="L171"></a>}
    <a id="L172"></a>z.sign = z.sign != (d &lt; 0);
    <a id="L173"></a>Nscale(&amp;z.mant, f);
<a id="L174"></a>}


<a id="L177"></a><span class="comment">// Mul1 returns the product x * d.</span>
<a id="L178"></a><span class="comment">//</span>
<a id="L179"></a>func (x *Integer) Mul1(d int64) *Integer {
    <a id="L180"></a>f := uint64(d);
    <a id="L181"></a>if d &lt; 0 {
        <a id="L182"></a>f = uint64(-d)
    <a id="L183"></a>}
    <a id="L184"></a>return MakeInt(x.sign != (d &lt; 0), x.mant.Mul1(f));
<a id="L185"></a>}


<a id="L188"></a><span class="comment">// Mul returns the product x * y.</span>
<a id="L189"></a><span class="comment">//</span>
<a id="L190"></a>func (x *Integer) Mul(y *Integer) *Integer {
    <a id="L191"></a><span class="comment">// x * y == x * y</span>
    <a id="L192"></a><span class="comment">// x * (-y) == -(x * y)</span>
    <a id="L193"></a><span class="comment">// (-x) * y == -(x * y)</span>
    <a id="L194"></a><span class="comment">// (-x) * (-y) == x * y</span>
    <a id="L195"></a>return MakeInt(x.sign != y.sign, x.mant.Mul(y.mant))
<a id="L196"></a>}


<a id="L199"></a><span class="comment">// MulNat returns the product x * y, where y is a (unsigned) natural number.</span>
<a id="L200"></a><span class="comment">//</span>
<a id="L201"></a>func (x *Integer) MulNat(y Natural) *Integer {
    <a id="L202"></a><span class="comment">// x * y == x * y</span>
    <a id="L203"></a><span class="comment">// (-x) * y == -(x * y)</span>
    <a id="L204"></a>return MakeInt(x.sign, x.mant.Mul(y))
<a id="L205"></a>}


<a id="L208"></a><span class="comment">// Quo returns the quotient q = x / y for y != 0.</span>
<a id="L209"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L210"></a><span class="comment">//</span>
<a id="L211"></a><span class="comment">// Quo and Rem implement T-division and modulus (like C99):</span>
<a id="L212"></a><span class="comment">//</span>
<a id="L213"></a><span class="comment">//   q = x.Quo(y) = trunc(x/y)  (truncation towards zero)</span>
<a id="L214"></a><span class="comment">//   r = x.Rem(y) = x - y*q</span>
<a id="L215"></a><span class="comment">//</span>
<a id="L216"></a><span class="comment">// (Daan Leijen, ``Division and Modulus for Computer Scientists&#39;&#39;.)</span>
<a id="L217"></a><span class="comment">//</span>
<a id="L218"></a>func (x *Integer) Quo(y *Integer) *Integer {
    <a id="L219"></a><span class="comment">// x / y == x / y</span>
    <a id="L220"></a><span class="comment">// x / (-y) == -(x / y)</span>
    <a id="L221"></a><span class="comment">// (-x) / y == -(x / y)</span>
    <a id="L222"></a><span class="comment">// (-x) / (-y) == x / y</span>
    <a id="L223"></a>return MakeInt(x.sign != y.sign, x.mant.Div(y.mant))
<a id="L224"></a>}


<a id="L227"></a><span class="comment">// Rem returns the remainder r of the division x / y for y != 0,</span>
<a id="L228"></a><span class="comment">// with r = x - y*x.Quo(y). Unless r is zero, its sign corresponds</span>
<a id="L229"></a><span class="comment">// to the sign of x.</span>
<a id="L230"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L231"></a><span class="comment">//</span>
<a id="L232"></a>func (x *Integer) Rem(y *Integer) *Integer {
    <a id="L233"></a><span class="comment">// x % y == x % y</span>
    <a id="L234"></a><span class="comment">// x % (-y) == x % y</span>
    <a id="L235"></a><span class="comment">// (-x) % y == -(x % y)</span>
    <a id="L236"></a><span class="comment">// (-x) % (-y) == -(x % y)</span>
    <a id="L237"></a>return MakeInt(x.sign, x.mant.Mod(y.mant))
<a id="L238"></a>}


<a id="L241"></a><span class="comment">// QuoRem returns the pair (x.Quo(y), x.Rem(y)) for y != 0.</span>
<a id="L242"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L243"></a><span class="comment">//</span>
<a id="L244"></a>func (x *Integer) QuoRem(y *Integer) (*Integer, *Integer) {
    <a id="L245"></a>q, r := x.mant.DivMod(y.mant);
    <a id="L246"></a>return MakeInt(x.sign != y.sign, q), MakeInt(x.sign, r);
<a id="L247"></a>}


<a id="L250"></a><span class="comment">// Div returns the quotient q = x / y for y != 0.</span>
<a id="L251"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L252"></a><span class="comment">//</span>
<a id="L253"></a><span class="comment">// Div and Mod implement Euclidian division and modulus:</span>
<a id="L254"></a><span class="comment">//</span>
<a id="L255"></a><span class="comment">//   q = x.Div(y)</span>
<a id="L256"></a><span class="comment">//   r = x.Mod(y) with: 0 &lt;= r &lt; |q| and: y = x*q + r</span>
<a id="L257"></a><span class="comment">//</span>
<a id="L258"></a><span class="comment">// (Raymond T. Boute, ``The Euclidian definition of the functions</span>
<a id="L259"></a><span class="comment">// div and mod&#39;&#39;. ACM Transactions on Programming Languages and</span>
<a id="L260"></a><span class="comment">// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.</span>
<a id="L261"></a><span class="comment">// ACM press.)</span>
<a id="L262"></a><span class="comment">//</span>
<a id="L263"></a>func (x *Integer) Div(y *Integer) *Integer {
    <a id="L264"></a>q, r := x.QuoRem(y);
    <a id="L265"></a>if r.IsNeg() {
        <a id="L266"></a>if y.IsPos() {
            <a id="L267"></a>q = q.Sub(Int(1))
        <a id="L268"></a>} else {
            <a id="L269"></a>q = q.Add(Int(1))
        <a id="L270"></a>}
    <a id="L271"></a>}
    <a id="L272"></a>return q;
<a id="L273"></a>}


<a id="L276"></a><span class="comment">// Mod returns the modulus r of the division x / y for y != 0,</span>
<a id="L277"></a><span class="comment">// with r = x - y*x.Div(y). r is always positive.</span>
<a id="L278"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L279"></a><span class="comment">//</span>
<a id="L280"></a>func (x *Integer) Mod(y *Integer) *Integer {
    <a id="L281"></a>r := x.Rem(y);
    <a id="L282"></a>if r.IsNeg() {
        <a id="L283"></a>if y.IsPos() {
            <a id="L284"></a>r = r.Add(y)
        <a id="L285"></a>} else {
            <a id="L286"></a>r = r.Sub(y)
        <a id="L287"></a>}
    <a id="L288"></a>}
    <a id="L289"></a>return r;
<a id="L290"></a>}


<a id="L293"></a><span class="comment">// DivMod returns the pair (x.Div(y), x.Mod(y)).</span>
<a id="L294"></a><span class="comment">//</span>
<a id="L295"></a>func (x *Integer) DivMod(y *Integer) (*Integer, *Integer) {
    <a id="L296"></a>q, r := x.QuoRem(y);
    <a id="L297"></a>if r.IsNeg() {
        <a id="L298"></a>if y.IsPos() {
            <a id="L299"></a>q = q.Sub(Int(1));
            <a id="L300"></a>r = r.Add(y);
        <a id="L301"></a>} else {
            <a id="L302"></a>q = q.Add(Int(1));
            <a id="L303"></a>r = r.Sub(y);
        <a id="L304"></a>}
    <a id="L305"></a>}
    <a id="L306"></a>return q, r;
<a id="L307"></a>}


<a id="L310"></a><span class="comment">// Shl implements ``shift left&#39;&#39; x &lt;&lt; s. It returns x * 2^s.</span>
<a id="L311"></a><span class="comment">//</span>
<a id="L312"></a>func (x *Integer) Shl(s uint) *Integer { return MakeInt(x.sign, x.mant.Shl(s)) }


<a id="L315"></a><span class="comment">// The bitwise operations on integers are defined on the 2&#39;s-complement</span>
<a id="L316"></a><span class="comment">// representation of integers. From</span>
<a id="L317"></a><span class="comment">//</span>
<a id="L318"></a><span class="comment">//   -x == ^x + 1  (1)  2&#39;s complement representation</span>
<a id="L319"></a><span class="comment">//</span>
<a id="L320"></a><span class="comment">// follows:</span>
<a id="L321"></a><span class="comment">//</span>
<a id="L322"></a><span class="comment">//   -(x) == ^(x) + 1</span>
<a id="L323"></a><span class="comment">//   -(-x) == ^(-x) + 1</span>
<a id="L324"></a><span class="comment">//   x-1 == ^(-x)</span>
<a id="L325"></a><span class="comment">//   ^(x-1) == -x  (2)</span>
<a id="L326"></a><span class="comment">//</span>
<a id="L327"></a><span class="comment">// Using (1) and (2), operations on negative integers of the form -x are</span>
<a id="L328"></a><span class="comment">// converted to operations on negated positive integers of the form ~(x-1).</span>


<a id="L331"></a><span class="comment">// Shr implements ``shift right&#39;&#39; x &gt;&gt; s. It returns x / 2^s.</span>
<a id="L332"></a><span class="comment">//</span>
<a id="L333"></a>func (x *Integer) Shr(s uint) *Integer {
    <a id="L334"></a>if x.sign {
        <a id="L335"></a><span class="comment">// (-x) &gt;&gt; s == ^(x-1) &gt;&gt; s == ^((x-1) &gt;&gt; s) == -(((x-1) &gt;&gt; s) + 1)</span>
        <a id="L336"></a>return MakeInt(true, x.mant.Sub(Nat(1)).Shr(s).Add(Nat(1)))
    <a id="L337"></a>}

    <a id="L339"></a>return MakeInt(false, x.mant.Shr(s));
<a id="L340"></a>}


<a id="L343"></a><span class="comment">// Not returns the ``bitwise not&#39;&#39; ^x for the 2&#39;s-complement representation of x.</span>
<a id="L344"></a>func (x *Integer) Not() *Integer {
    <a id="L345"></a>if x.sign {
        <a id="L346"></a><span class="comment">// ^(-x) == ^(^(x-1)) == x-1</span>
        <a id="L347"></a>return MakeInt(false, x.mant.Sub(Nat(1)))
    <a id="L348"></a>}

    <a id="L350"></a><span class="comment">// ^x == -x-1 == -(x+1)</span>
    <a id="L351"></a>return MakeInt(true, x.mant.Add(Nat(1)));
<a id="L352"></a>}


<a id="L355"></a><span class="comment">// And returns the ``bitwise and&#39;&#39; x &amp; y for the 2&#39;s-complement representation of x and y.</span>
<a id="L356"></a><span class="comment">//</span>
<a id="L357"></a>func (x *Integer) And(y *Integer) *Integer {
    <a id="L358"></a>if x.sign == y.sign {
        <a id="L359"></a>if x.sign {
            <a id="L360"></a><span class="comment">// (-x) &amp; (-y) == ^(x-1) &amp; ^(y-1) == ^((x-1) | (y-1)) == -(((x-1) | (y-1)) + 1)</span>
            <a id="L361"></a>return MakeInt(true, x.mant.Sub(Nat(1)).Or(y.mant.Sub(Nat(1))).Add(Nat(1)))
        <a id="L362"></a>}

        <a id="L364"></a><span class="comment">// x &amp; y == x &amp; y</span>
        <a id="L365"></a>return MakeInt(false, x.mant.And(y.mant));
    <a id="L366"></a>}

    <a id="L368"></a><span class="comment">// x.sign != y.sign</span>
    <a id="L369"></a>if x.sign {
        <a id="L370"></a>x, y = y, x <span class="comment">// &amp; is symmetric</span>
    <a id="L371"></a>}

    <a id="L373"></a><span class="comment">// x &amp; (-y) == x &amp; ^(y-1) == x &amp;^ (y-1)</span>
    <a id="L374"></a>return MakeInt(false, x.mant.AndNot(y.mant.Sub(Nat(1))));
<a id="L375"></a>}


<a id="L378"></a><span class="comment">// AndNot returns the ``bitwise clear&#39;&#39; x &amp;^ y for the 2&#39;s-complement representation of x and y.</span>
<a id="L379"></a><span class="comment">//</span>
<a id="L380"></a>func (x *Integer) AndNot(y *Integer) *Integer {
    <a id="L381"></a>if x.sign == y.sign {
        <a id="L382"></a>if x.sign {
            <a id="L383"></a><span class="comment">// (-x) &amp;^ (-y) == ^(x-1) &amp;^ ^(y-1) == ^(x-1) &amp; (y-1) == (y-1) &amp;^ (x-1)</span>
            <a id="L384"></a>return MakeInt(false, y.mant.Sub(Nat(1)).AndNot(x.mant.Sub(Nat(1))))
        <a id="L385"></a>}

        <a id="L387"></a><span class="comment">// x &amp;^ y == x &amp;^ y</span>
        <a id="L388"></a>return MakeInt(false, x.mant.AndNot(y.mant));
    <a id="L389"></a>}

    <a id="L391"></a>if x.sign {
        <a id="L392"></a><span class="comment">// (-x) &amp;^ y == ^(x-1) &amp;^ y == ^(x-1) &amp; ^y == ^((x-1) | y) == -(((x-1) | y) + 1)</span>
        <a id="L393"></a>return MakeInt(true, x.mant.Sub(Nat(1)).Or(y.mant).Add(Nat(1)))
    <a id="L394"></a>}

    <a id="L396"></a><span class="comment">// x &amp;^ (-y) == x &amp;^ ^(y-1) == x &amp; (y-1)</span>
    <a id="L397"></a>return MakeInt(false, x.mant.And(y.mant.Sub(Nat(1))));
<a id="L398"></a>}


<a id="L401"></a><span class="comment">// Or returns the ``bitwise or&#39;&#39; x | y for the 2&#39;s-complement representation of x and y.</span>
<a id="L402"></a><span class="comment">//</span>
<a id="L403"></a>func (x *Integer) Or(y *Integer) *Integer {
    <a id="L404"></a>if x.sign == y.sign {
        <a id="L405"></a>if x.sign {
            <a id="L406"></a><span class="comment">// (-x) | (-y) == ^(x-1) | ^(y-1) == ^((x-1) &amp; (y-1)) == -(((x-1) &amp; (y-1)) + 1)</span>
            <a id="L407"></a>return MakeInt(true, x.mant.Sub(Nat(1)).And(y.mant.Sub(Nat(1))).Add(Nat(1)))
        <a id="L408"></a>}

        <a id="L410"></a><span class="comment">// x | y == x | y</span>
        <a id="L411"></a>return MakeInt(false, x.mant.Or(y.mant));
    <a id="L412"></a>}

    <a id="L414"></a><span class="comment">// x.sign != y.sign</span>
    <a id="L415"></a>if x.sign {
        <a id="L416"></a>x, y = y, x <span class="comment">// | or symmetric</span>
    <a id="L417"></a>}

    <a id="L419"></a><span class="comment">// x | (-y) == x | ^(y-1) == ^((y-1) &amp;^ x) == -(^((y-1) &amp;^ x) + 1)</span>
    <a id="L420"></a>return MakeInt(true, y.mant.Sub(Nat(1)).AndNot(x.mant).Add(Nat(1)));
<a id="L421"></a>}


<a id="L424"></a><span class="comment">// Xor returns the ``bitwise xor&#39;&#39; x | y for the 2&#39;s-complement representation of x and y.</span>
<a id="L425"></a><span class="comment">//</span>
<a id="L426"></a>func (x *Integer) Xor(y *Integer) *Integer {
    <a id="L427"></a>if x.sign == y.sign {
        <a id="L428"></a>if x.sign {
            <a id="L429"></a><span class="comment">// (-x) ^ (-y) == ^(x-1) ^ ^(y-1) == (x-1) ^ (y-1)</span>
            <a id="L430"></a>return MakeInt(false, x.mant.Sub(Nat(1)).Xor(y.mant.Sub(Nat(1))))
        <a id="L431"></a>}

        <a id="L433"></a><span class="comment">// x ^ y == x ^ y</span>
        <a id="L434"></a>return MakeInt(false, x.mant.Xor(y.mant));
    <a id="L435"></a>}

    <a id="L437"></a><span class="comment">// x.sign != y.sign</span>
    <a id="L438"></a>if x.sign {
        <a id="L439"></a>x, y = y, x <span class="comment">// ^ is symmetric</span>
    <a id="L440"></a>}

    <a id="L442"></a><span class="comment">// x ^ (-y) == x ^ ^(y-1) == ^(x ^ (y-1)) == -((x ^ (y-1)) + 1)</span>
    <a id="L443"></a>return MakeInt(true, x.mant.Xor(y.mant.Sub(Nat(1))).Add(Nat(1)));
<a id="L444"></a>}


<a id="L447"></a><span class="comment">// Cmp compares x and y. The result is an int value that is</span>
<a id="L448"></a><span class="comment">//</span>
<a id="L449"></a><span class="comment">//   &lt;  0 if x &lt;  y</span>
<a id="L450"></a><span class="comment">//   == 0 if x == y</span>
<a id="L451"></a><span class="comment">//   &gt;  0 if x &gt;  y</span>
<a id="L452"></a><span class="comment">//</span>
<a id="L453"></a>func (x *Integer) Cmp(y *Integer) int {
    <a id="L454"></a><span class="comment">// x cmp y == x cmp y</span>
    <a id="L455"></a><span class="comment">// x cmp (-y) == x</span>
    <a id="L456"></a><span class="comment">// (-x) cmp y == y</span>
    <a id="L457"></a><span class="comment">// (-x) cmp (-y) == -(x cmp y)</span>
    <a id="L458"></a>var r int;
    <a id="L459"></a>switch {
    <a id="L460"></a>case x.sign == y.sign:
        <a id="L461"></a>r = x.mant.Cmp(y.mant);
        <a id="L462"></a>if x.sign {
            <a id="L463"></a>r = -r
        <a id="L464"></a>}
    <a id="L465"></a>case x.sign:
        <a id="L466"></a>r = -1
    <a id="L467"></a>case y.sign:
        <a id="L468"></a>r = 1
    <a id="L469"></a>}
    <a id="L470"></a>return r;
<a id="L471"></a>}


<a id="L474"></a><span class="comment">// ToString converts x to a string for a given base, with 2 &lt;= base &lt;= 16.</span>
<a id="L475"></a><span class="comment">//</span>
<a id="L476"></a>func (x *Integer) ToString(base uint) string {
    <a id="L477"></a>if x.mant.IsZero() {
        <a id="L478"></a>return &#34;0&#34;
    <a id="L479"></a>}
    <a id="L480"></a>var s string;
    <a id="L481"></a>if x.sign {
        <a id="L482"></a>s = &#34;-&#34;
    <a id="L483"></a>}
    <a id="L484"></a>return s + x.mant.ToString(base);
<a id="L485"></a>}


<a id="L488"></a><span class="comment">// String converts x to its decimal string representation.</span>
<a id="L489"></a><span class="comment">// x.String() is the same as x.ToString(10).</span>
<a id="L490"></a><span class="comment">//</span>
<a id="L491"></a>func (x *Integer) String() string { return x.ToString(10) }


<a id="L494"></a><span class="comment">// Format is a support routine for fmt.Formatter. It accepts</span>
<a id="L495"></a><span class="comment">// the formats &#39;b&#39; (binary), &#39;o&#39; (octal), and &#39;x&#39; (hexadecimal).</span>
<a id="L496"></a><span class="comment">//</span>
<a id="L497"></a>func (x *Integer) Format(h fmt.State, c int) { fmt.Fprintf(h, &#34;%s&#34;, x.ToString(fmtbase(c))) }


<a id="L500"></a><span class="comment">// IntFromString returns the integer corresponding to the</span>
<a id="L501"></a><span class="comment">// longest possible prefix of s representing an integer in a</span>
<a id="L502"></a><span class="comment">// given conversion base, the actual conversion base used, and</span>
<a id="L503"></a><span class="comment">// the prefix length. The syntax of integers follows the syntax</span>
<a id="L504"></a><span class="comment">// of signed integer literals in Go.</span>
<a id="L505"></a><span class="comment">//</span>
<a id="L506"></a><span class="comment">// If the base argument is 0, the string prefix determines the actual</span>
<a id="L507"></a><span class="comment">// conversion base. A prefix of ``0x&#39;&#39; or ``0X&#39;&#39; selects base 16; the</span>
<a id="L508"></a><span class="comment">// ``0&#39;&#39; prefix selects base 8. Otherwise the selected base is 10.</span>
<a id="L509"></a><span class="comment">//</span>
<a id="L510"></a>func IntFromString(s string, base uint) (*Integer, uint, int) {
    <a id="L511"></a><span class="comment">// skip sign, if any</span>
    <a id="L512"></a>i0 := 0;
    <a id="L513"></a>if len(s) &gt; 0 &amp;&amp; (s[0] == &#39;-&#39; || s[0] == &#39;+&#39;) {
        <a id="L514"></a>i0 = 1
    <a id="L515"></a>}

    <a id="L517"></a>mant, base, slen := NatFromString(s[i0:len(s)], base);

    <a id="L519"></a>return MakeInt(i0 &gt; 0 &amp;&amp; s[0] == &#39;-&#39;, mant), base, i0 + slen;
<a id="L520"></a>}
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
