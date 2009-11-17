<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bignum/rational.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/bignum/rational.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Rational numbers</span>

<a id="L7"></a>package bignum

<a id="L9"></a>import &#34;fmt&#34;


<a id="L12"></a><span class="comment">// Rational represents a quotient a/b of arbitrary precision.</span>
<a id="L13"></a><span class="comment">//</span>
<a id="L14"></a>type Rational struct {
    <a id="L15"></a>a   *Integer; <span class="comment">// numerator</span>
    <a id="L16"></a>b   Natural;  <span class="comment">// denominator</span>
<a id="L17"></a>}


<a id="L20"></a><span class="comment">// MakeRat makes a rational number given a numerator a and a denominator b.</span>
<a id="L21"></a><span class="comment">//</span>
<a id="L22"></a>func MakeRat(a *Integer, b Natural) *Rational {
    <a id="L23"></a>f := a.mant.Gcd(b); <span class="comment">// f &gt; 0</span>
    <a id="L24"></a>if f.Cmp(Nat(1)) != 0 {
        <a id="L25"></a>a = MakeInt(a.sign, a.mant.Div(f));
        <a id="L26"></a>b = b.Div(f);
    <a id="L27"></a>}
    <a id="L28"></a>return &amp;Rational{a, b};
<a id="L29"></a>}


<a id="L32"></a><span class="comment">// Rat creates a small rational number with value a0/b0.</span>
<a id="L33"></a><span class="comment">//</span>
<a id="L34"></a>func Rat(a0 int64, b0 int64) *Rational {
    <a id="L35"></a>a, b := Int(a0), Int(b0);
    <a id="L36"></a>if b.sign {
        <a id="L37"></a>a = a.Neg()
    <a id="L38"></a>}
    <a id="L39"></a>return MakeRat(a, b.mant);
<a id="L40"></a>}


<a id="L43"></a><span class="comment">// Value returns the numerator and denominator of x.</span>
<a id="L44"></a><span class="comment">//</span>
<a id="L45"></a>func (x *Rational) Value() (numerator *Integer, denominator Natural) {
    <a id="L46"></a>return x.a, x.b
<a id="L47"></a>}


<a id="L50"></a><span class="comment">// Predicates</span>

<a id="L52"></a><span class="comment">// IsZero returns true iff x == 0.</span>
<a id="L53"></a><span class="comment">//</span>
<a id="L54"></a>func (x *Rational) IsZero() bool { return x.a.IsZero() }


<a id="L57"></a><span class="comment">// IsNeg returns true iff x &lt; 0.</span>
<a id="L58"></a><span class="comment">//</span>
<a id="L59"></a>func (x *Rational) IsNeg() bool { return x.a.IsNeg() }


<a id="L62"></a><span class="comment">// IsPos returns true iff x &gt; 0.</span>
<a id="L63"></a><span class="comment">//</span>
<a id="L64"></a>func (x *Rational) IsPos() bool { return x.a.IsPos() }


<a id="L67"></a><span class="comment">// IsInt returns true iff x can be written with a denominator 1</span>
<a id="L68"></a><span class="comment">// in the form x == x&#39;/1; i.e., if x is an integer value.</span>
<a id="L69"></a><span class="comment">//</span>
<a id="L70"></a>func (x *Rational) IsInt() bool { return x.b.Cmp(Nat(1)) == 0 }


<a id="L73"></a><span class="comment">// Operations</span>

<a id="L75"></a><span class="comment">// Neg returns the negated value of x.</span>
<a id="L76"></a><span class="comment">//</span>
<a id="L77"></a>func (x *Rational) Neg() *Rational { return MakeRat(x.a.Neg(), x.b) }


<a id="L80"></a><span class="comment">// Add returns the sum x + y.</span>
<a id="L81"></a><span class="comment">//</span>
<a id="L82"></a>func (x *Rational) Add(y *Rational) *Rational {
    <a id="L83"></a>return MakeRat((x.a.MulNat(y.b)).Add(y.a.MulNat(x.b)), x.b.Mul(y.b))
<a id="L84"></a>}


<a id="L87"></a><span class="comment">// Sub returns the difference x - y.</span>
<a id="L88"></a><span class="comment">//</span>
<a id="L89"></a>func (x *Rational) Sub(y *Rational) *Rational {
    <a id="L90"></a>return MakeRat((x.a.MulNat(y.b)).Sub(y.a.MulNat(x.b)), x.b.Mul(y.b))
<a id="L91"></a>}


<a id="L94"></a><span class="comment">// Mul returns the product x * y.</span>
<a id="L95"></a><span class="comment">//</span>
<a id="L96"></a>func (x *Rational) Mul(y *Rational) *Rational { return MakeRat(x.a.Mul(y.a), x.b.Mul(y.b)) }


<a id="L99"></a><span class="comment">// Quo returns the quotient x / y for y != 0.</span>
<a id="L100"></a><span class="comment">// If y == 0, a division-by-zero run-time error occurs.</span>
<a id="L101"></a><span class="comment">//</span>
<a id="L102"></a>func (x *Rational) Quo(y *Rational) *Rational {
    <a id="L103"></a>a := x.a.MulNat(y.b);
    <a id="L104"></a>b := y.a.MulNat(x.b);
    <a id="L105"></a>if b.IsNeg() {
        <a id="L106"></a>a = a.Neg()
    <a id="L107"></a>}
    <a id="L108"></a>return MakeRat(a, b.mant);
<a id="L109"></a>}


<a id="L112"></a><span class="comment">// Cmp compares x and y. The result is an int value</span>
<a id="L113"></a><span class="comment">//</span>
<a id="L114"></a><span class="comment">//   &lt;  0 if x &lt;  y</span>
<a id="L115"></a><span class="comment">//   == 0 if x == y</span>
<a id="L116"></a><span class="comment">//   &gt;  0 if x &gt;  y</span>
<a id="L117"></a><span class="comment">//</span>
<a id="L118"></a>func (x *Rational) Cmp(y *Rational) int { return (x.a.MulNat(y.b)).Cmp(y.a.MulNat(x.b)) }


<a id="L121"></a><span class="comment">// ToString converts x to a string for a given base, with 2 &lt;= base &lt;= 16.</span>
<a id="L122"></a><span class="comment">// The string representation is of the form &#34;n&#34; if x is an integer; otherwise</span>
<a id="L123"></a><span class="comment">// it is of form &#34;n/d&#34;.</span>
<a id="L124"></a><span class="comment">//</span>
<a id="L125"></a>func (x *Rational) ToString(base uint) string {
    <a id="L126"></a>s := x.a.ToString(base);
    <a id="L127"></a>if !x.IsInt() {
        <a id="L128"></a>s += &#34;/&#34; + x.b.ToString(base)
    <a id="L129"></a>}
    <a id="L130"></a>return s;
<a id="L131"></a>}


<a id="L134"></a><span class="comment">// String converts x to its decimal string representation.</span>
<a id="L135"></a><span class="comment">// x.String() is the same as x.ToString(10).</span>
<a id="L136"></a><span class="comment">//</span>
<a id="L137"></a>func (x *Rational) String() string { return x.ToString(10) }


<a id="L140"></a><span class="comment">// Format is a support routine for fmt.Formatter. It accepts</span>
<a id="L141"></a><span class="comment">// the formats &#39;b&#39; (binary), &#39;o&#39; (octal), and &#39;x&#39; (hexadecimal).</span>
<a id="L142"></a><span class="comment">//</span>
<a id="L143"></a>func (x *Rational) Format(h fmt.State, c int) { fmt.Fprintf(h, &#34;%s&#34;, x.ToString(fmtbase(c))) }


<a id="L146"></a><span class="comment">// RatFromString returns the rational number corresponding to the</span>
<a id="L147"></a><span class="comment">// longest possible prefix of s representing a rational number in a</span>
<a id="L148"></a><span class="comment">// given conversion base, the actual conversion base used, and the</span>
<a id="L149"></a><span class="comment">// prefix length. The syntax of a rational number is:</span>
<a id="L150"></a><span class="comment">//</span>
<a id="L151"></a><span class="comment">//	rational = mantissa [exponent] .</span>
<a id="L152"></a><span class="comment">//	mantissa = integer (&#39;/&#39; natural | &#39;.&#39; natural) .</span>
<a id="L153"></a><span class="comment">//	exponent = (&#39;e&#39;|&#39;E&#39;) integer .</span>
<a id="L154"></a><span class="comment">//</span>
<a id="L155"></a><span class="comment">// If the base argument is 0, the string prefix determines the actual</span>
<a id="L156"></a><span class="comment">// conversion base for the mantissa. A prefix of ``0x&#39;&#39; or ``0X&#39;&#39; selects</span>
<a id="L157"></a><span class="comment">// base 16; the ``0&#39;&#39; prefix selects base 8. Otherwise the selected base is 10.</span>
<a id="L158"></a><span class="comment">// If the mantissa is represented via a division, both the numerator and</span>
<a id="L159"></a><span class="comment">// denominator may have different base prefixes; in that case the base of</span>
<a id="L160"></a><span class="comment">// of the numerator is returned. If the mantissa contains a decimal point,</span>
<a id="L161"></a><span class="comment">// the base for the fractional part is the same as for the part before the</span>
<a id="L162"></a><span class="comment">// decimal point and the fractional part does not accept a base prefix.</span>
<a id="L163"></a><span class="comment">// The base for the exponent is always 10.</span>
<a id="L164"></a><span class="comment">//</span>
<a id="L165"></a>func RatFromString(s string, base uint) (*Rational, uint, int) {
    <a id="L166"></a><span class="comment">// read numerator</span>
    <a id="L167"></a>a, abase, alen := IntFromString(s, base);
    <a id="L168"></a>b := Nat(1);

    <a id="L170"></a><span class="comment">// read denominator or fraction, if any</span>
    <a id="L171"></a>var blen int;
    <a id="L172"></a>if alen &lt; len(s) {
        <a id="L173"></a>ch := s[alen];
        <a id="L174"></a>if ch == &#39;/&#39; {
            <a id="L175"></a>alen++;
            <a id="L176"></a>b, base, blen = NatFromString(s[alen:len(s)], base);
        <a id="L177"></a>} else if ch == &#39;.&#39; {
            <a id="L178"></a>alen++;
            <a id="L179"></a>b, base, blen = NatFromString(s[alen:len(s)], abase);
            <a id="L180"></a>assert(base == abase);
            <a id="L181"></a>f := Nat(uint64(base)).Pow(uint(blen));
            <a id="L182"></a>a = MakeInt(a.sign, a.mant.Mul(f).Add(b));
            <a id="L183"></a>b = f;
        <a id="L184"></a>}
    <a id="L185"></a>}

    <a id="L187"></a><span class="comment">// read exponent, if any</span>
    <a id="L188"></a>rlen := alen + blen;
    <a id="L189"></a>if rlen &lt; len(s) {
        <a id="L190"></a>ch := s[rlen];
        <a id="L191"></a>if ch == &#39;e&#39; || ch == &#39;E&#39; {
            <a id="L192"></a>rlen++;
            <a id="L193"></a>e, _, elen := IntFromString(s[rlen:len(s)], 10);
            <a id="L194"></a>rlen += elen;
            <a id="L195"></a>m := Nat(10).Pow(uint(e.mant.Value()));
            <a id="L196"></a>if e.sign {
                <a id="L197"></a>b = b.Mul(m)
            <a id="L198"></a>} else {
                <a id="L199"></a>a = a.MulNat(m)
            <a id="L200"></a>}
        <a id="L201"></a>}
    <a id="L202"></a>}

    <a id="L204"></a>return MakeRat(a, b), base, rlen;
<a id="L205"></a>}
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
