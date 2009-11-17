<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/atof.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/strconv/atof.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// decimal to binary floating point conversion.</span>
<a id="L6"></a><span class="comment">// Algorithm:</span>
<a id="L7"></a><span class="comment">//   1) Store input in multiprecision decimal.</span>
<a id="L8"></a><span class="comment">//   2) Multiply/divide decimal by powers of two until in range [0.5, 1)</span>
<a id="L9"></a><span class="comment">//   3) Multiply by 2^precision and round to get mantissa.</span>

<a id="L11"></a><span class="comment">// The strconv package implements conversions to and from</span>
<a id="L12"></a><span class="comment">// string representations of basic data types.</span>
<a id="L13"></a>package strconv

<a id="L15"></a>import (
    <a id="L16"></a>&#34;math&#34;;
    <a id="L17"></a>&#34;os&#34;;
<a id="L18"></a>)

<a id="L20"></a>var optimize = true <span class="comment">// can change for testing</span>

<a id="L22"></a><span class="comment">// TODO(rsc): Better truncation handling.</span>
<a id="L23"></a>func stringToDecimal(s string) (neg bool, d *decimal, trunc bool, ok bool) {
    <a id="L24"></a>i := 0;

    <a id="L26"></a><span class="comment">// optional sign</span>
    <a id="L27"></a>if i &gt;= len(s) {
        <a id="L28"></a>return
    <a id="L29"></a>}
    <a id="L30"></a>switch {
    <a id="L31"></a>case s[i] == &#39;+&#39;:
        <a id="L32"></a>i++
    <a id="L33"></a>case s[i] == &#39;-&#39;:
        <a id="L34"></a>neg = true;
        <a id="L35"></a>i++;
    <a id="L36"></a>}

    <a id="L38"></a><span class="comment">// digits</span>
    <a id="L39"></a>b := new(decimal);
    <a id="L40"></a>sawdot := false;
    <a id="L41"></a>sawdigits := false;
    <a id="L42"></a>for ; i &lt; len(s); i++ {
        <a id="L43"></a>switch {
        <a id="L44"></a>case s[i] == &#39;.&#39;:
            <a id="L45"></a>if sawdot {
                <a id="L46"></a>return
            <a id="L47"></a>}
            <a id="L48"></a>sawdot = true;
            <a id="L49"></a>b.dp = b.nd;
            <a id="L50"></a>continue;

        <a id="L52"></a>case &#39;0&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;9&#39;:
            <a id="L53"></a>sawdigits = true;
            <a id="L54"></a>if s[i] == &#39;0&#39; &amp;&amp; b.nd == 0 { <span class="comment">// ignore leading zeros</span>
                <a id="L55"></a>b.dp--;
                <a id="L56"></a>continue;
            <a id="L57"></a>}
            <a id="L58"></a>b.d[b.nd] = s[i];
            <a id="L59"></a>b.nd++;
            <a id="L60"></a>continue;
        <a id="L61"></a>}
        <a id="L62"></a>break;
    <a id="L63"></a>}
    <a id="L64"></a>if !sawdigits {
        <a id="L65"></a>return
    <a id="L66"></a>}
    <a id="L67"></a>if !sawdot {
        <a id="L68"></a>b.dp = b.nd
    <a id="L69"></a>}

    <a id="L71"></a><span class="comment">// optional exponent moves decimal point.</span>
    <a id="L72"></a><span class="comment">// if we read a very large, very long number,</span>
    <a id="L73"></a><span class="comment">// just be sure to move the decimal point by</span>
    <a id="L74"></a><span class="comment">// a lot (say, 100000).  it doesn&#39;t matter if it&#39;s</span>
    <a id="L75"></a><span class="comment">// not the exact number.</span>
    <a id="L76"></a>if i &lt; len(s) &amp;&amp; s[i] == &#39;e&#39; {
        <a id="L77"></a>i++;
        <a id="L78"></a>if i &gt;= len(s) {
            <a id="L79"></a>return
        <a id="L80"></a>}
        <a id="L81"></a>esign := 1;
        <a id="L82"></a>if s[i] == &#39;+&#39; {
            <a id="L83"></a>i++
        <a id="L84"></a>} else if s[i] == &#39;-&#39; {
            <a id="L85"></a>i++;
            <a id="L86"></a>esign = -1;
        <a id="L87"></a>}
        <a id="L88"></a>if i &gt;= len(s) || s[i] &lt; &#39;0&#39; || s[i] &gt; &#39;9&#39; {
            <a id="L89"></a>return
        <a id="L90"></a>}
        <a id="L91"></a>e := 0;
        <a id="L92"></a>for ; i &lt; len(s) &amp;&amp; &#39;0&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;9&#39;; i++ {
            <a id="L93"></a>if e &lt; 10000 {
                <a id="L94"></a>e = e*10 + int(s[i]) - &#39;0&#39;
            <a id="L95"></a>}
        <a id="L96"></a>}
        <a id="L97"></a>b.dp += e * esign;
    <a id="L98"></a>}

    <a id="L100"></a>if i != len(s) {
        <a id="L101"></a>return
    <a id="L102"></a>}

    <a id="L104"></a>d = b;
    <a id="L105"></a>ok = true;
    <a id="L106"></a>return;
<a id="L107"></a>}

<a id="L109"></a><span class="comment">// decimal power of ten to binary power of two.</span>
<a id="L110"></a>var powtab = []int{1, 3, 6, 9, 13, 16, 19, 23, 26}

<a id="L112"></a>func decimalToFloatBits(neg bool, d *decimal, trunc bool, flt *floatInfo) (b uint64, overflow bool) {
    <a id="L113"></a>var exp int;
    <a id="L114"></a>var mant uint64;

    <a id="L116"></a><span class="comment">// Zero is always a special case.</span>
    <a id="L117"></a>if d.nd == 0 {
        <a id="L118"></a>mant = 0;
        <a id="L119"></a>exp = flt.bias;
        <a id="L120"></a>goto out;
    <a id="L121"></a>}

    <a id="L123"></a><span class="comment">// Obvious overflow/underflow.</span>
    <a id="L124"></a><span class="comment">// These bounds are for 64-bit floats.</span>
    <a id="L125"></a><span class="comment">// Will have to change if we want to support 80-bit floats in the future.</span>
    <a id="L126"></a>if d.dp &gt; 310 {
        <a id="L127"></a>goto overflow
    <a id="L128"></a>}
    <a id="L129"></a>if d.dp &lt; -330 {
        <a id="L130"></a><span class="comment">// zero</span>
        <a id="L131"></a>mant = 0;
        <a id="L132"></a>exp = flt.bias;
        <a id="L133"></a>goto out;
    <a id="L134"></a>}

    <a id="L136"></a><span class="comment">// Scale by powers of two until in range [0.5, 1.0)</span>
    <a id="L137"></a>exp = 0;
    <a id="L138"></a>for d.dp &gt; 0 {
        <a id="L139"></a>var n int;
        <a id="L140"></a>if d.dp &gt;= len(powtab) {
            <a id="L141"></a>n = 27
        <a id="L142"></a>} else {
            <a id="L143"></a>n = powtab[d.dp]
        <a id="L144"></a>}
        <a id="L145"></a>d.Shift(-n);
        <a id="L146"></a>exp += n;
    <a id="L147"></a>}
    <a id="L148"></a>for d.dp &lt; 0 || d.dp == 0 &amp;&amp; d.d[0] &lt; &#39;5&#39; {
        <a id="L149"></a>var n int;
        <a id="L150"></a>if -d.dp &gt;= len(powtab) {
            <a id="L151"></a>n = 27
        <a id="L152"></a>} else {
            <a id="L153"></a>n = powtab[-d.dp]
        <a id="L154"></a>}
        <a id="L155"></a>d.Shift(n);
        <a id="L156"></a>exp -= n;
    <a id="L157"></a>}

    <a id="L159"></a><span class="comment">// Our range is [0.5,1) but floating point range is [1,2).</span>
    <a id="L160"></a>exp--;

    <a id="L162"></a><span class="comment">// Minimum representable exponent is flt.bias+1.</span>
    <a id="L163"></a><span class="comment">// If the exponent is smaller, move it up and</span>
    <a id="L164"></a><span class="comment">// adjust d accordingly.</span>
    <a id="L165"></a>if exp &lt; flt.bias+1 {
        <a id="L166"></a>n := flt.bias + 1 - exp;
        <a id="L167"></a>d.Shift(-n);
        <a id="L168"></a>exp += n;
    <a id="L169"></a>}

    <a id="L171"></a>if exp-flt.bias &gt;= 1&lt;&lt;flt.expbits-1 {
        <a id="L172"></a>goto overflow
    <a id="L173"></a>}

    <a id="L175"></a><span class="comment">// Extract 1+flt.mantbits bits.</span>
    <a id="L176"></a>mant = d.Shift(int(1 + flt.mantbits)).RoundedInteger();

    <a id="L178"></a><span class="comment">// Rounding might have added a bit; shift down.</span>
    <a id="L179"></a>if mant == 2&lt;&lt;flt.mantbits {
        <a id="L180"></a>mant &gt;&gt;= 1;
        <a id="L181"></a>exp++;
        <a id="L182"></a>if exp-flt.bias &gt;= 1&lt;&lt;flt.expbits-1 {
            <a id="L183"></a>goto overflow
        <a id="L184"></a>}
    <a id="L185"></a>}

    <a id="L187"></a><span class="comment">// Denormalized?</span>
    <a id="L188"></a>if mant&amp;(1&lt;&lt;flt.mantbits) == 0 {
        <a id="L189"></a>exp = flt.bias
    <a id="L190"></a>}
    <a id="L191"></a>goto out;

<a id="L193"></a>overflow:
    <a id="L194"></a><span class="comment">// ±Inf</span>
    <a id="L195"></a>mant = 0;
    <a id="L196"></a>exp = 1&lt;&lt;flt.expbits - 1 + flt.bias;
    <a id="L197"></a>overflow = true;

<a id="L199"></a>out:
    <a id="L200"></a><span class="comment">// Assemble bits.</span>
    <a id="L201"></a>bits := mant &amp; (uint64(1)&lt;&lt;flt.mantbits - 1);
    <a id="L202"></a>bits |= uint64((exp-flt.bias)&amp;(1&lt;&lt;flt.expbits-1)) &lt;&lt; flt.mantbits;
    <a id="L203"></a>if neg {
        <a id="L204"></a>bits |= 1 &lt;&lt; flt.mantbits &lt;&lt; flt.expbits
    <a id="L205"></a>}
    <a id="L206"></a>return bits, overflow;
<a id="L207"></a>}

<a id="L209"></a><span class="comment">// Compute exact floating-point integer from d&#39;s digits.</span>
<a id="L210"></a><span class="comment">// Caller is responsible for avoiding overflow.</span>
<a id="L211"></a>func decimalAtof64Int(neg bool, d *decimal) float64 {
    <a id="L212"></a>f := float64(0);
    <a id="L213"></a>for i := 0; i &lt; d.nd; i++ {
        <a id="L214"></a>f = f*10 + float64(d.d[i]-&#39;0&#39;)
    <a id="L215"></a>}
    <a id="L216"></a>if neg {
        <a id="L217"></a>f *= -1 <span class="comment">// BUG work around 6g f = -f.</span>
    <a id="L218"></a>}
    <a id="L219"></a>return f;
<a id="L220"></a>}

<a id="L222"></a>func decimalAtof32Int(neg bool, d *decimal) float32 {
    <a id="L223"></a>f := float32(0);
    <a id="L224"></a>for i := 0; i &lt; d.nd; i++ {
        <a id="L225"></a>f = f*10 + float32(d.d[i]-&#39;0&#39;)
    <a id="L226"></a>}
    <a id="L227"></a>if neg {
        <a id="L228"></a>f *= -1 <span class="comment">// BUG work around 6g f = -f.</span>
    <a id="L229"></a>}
    <a id="L230"></a>return f;
<a id="L231"></a>}

<a id="L233"></a><span class="comment">// Exact powers of 10.</span>
<a id="L234"></a>var float64pow10 = []float64{
    <a id="L235"></a>1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9,
    <a id="L236"></a>1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18, 1e19,
    <a id="L237"></a>1e20, 1e21, 1e22,
<a id="L238"></a>}
<a id="L239"></a>var float32pow10 = []float32{1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9, 1e10}

<a id="L241"></a><span class="comment">// If possible to convert decimal d to 64-bit float f exactly,</span>
<a id="L242"></a><span class="comment">// entirely in floating-point math, do so, avoiding the expense of decimalToFloatBits.</span>
<a id="L243"></a><span class="comment">// Three common cases:</span>
<a id="L244"></a><span class="comment">//	value is exact integer</span>
<a id="L245"></a><span class="comment">//	value is exact integer * exact power of ten</span>
<a id="L246"></a><span class="comment">//	value is exact integer / exact power of ten</span>
<a id="L247"></a><span class="comment">// These all produce potentially inexact but correctly rounded answers.</span>
<a id="L248"></a>func decimalAtof64(neg bool, d *decimal, trunc bool) (f float64, ok bool) {
    <a id="L249"></a><span class="comment">// Exact integers are &lt;= 10^15.</span>
    <a id="L250"></a><span class="comment">// Exact powers of ten are &lt;= 10^22.</span>
    <a id="L251"></a>if d.nd &gt; 15 {
        <a id="L252"></a>return
    <a id="L253"></a>}
    <a id="L254"></a>switch {
    <a id="L255"></a>case d.dp == d.nd: <span class="comment">// int</span>
        <a id="L256"></a>f := decimalAtof64Int(neg, d);
        <a id="L257"></a>return f, true;

    <a id="L259"></a>case d.dp &gt; d.nd &amp;&amp; d.dp &lt;= 15+22: <span class="comment">// int * 10^k</span>
        <a id="L260"></a>f := decimalAtof64Int(neg, d);
        <a id="L261"></a>k := d.dp - d.nd;
        <a id="L262"></a><span class="comment">// If exponent is big but number of digits is not,</span>
        <a id="L263"></a><span class="comment">// can move a few zeros into the integer part.</span>
        <a id="L264"></a>if k &gt; 22 {
            <a id="L265"></a>f *= float64pow10[k-22];
            <a id="L266"></a>k = 22;
        <a id="L267"></a>}
        <a id="L268"></a>return f * float64pow10[k], true;

    <a id="L270"></a>case d.dp &lt; d.nd &amp;&amp; d.nd-d.dp &lt;= 22: <span class="comment">// int / 10^k</span>
        <a id="L271"></a>f := decimalAtof64Int(neg, d);
        <a id="L272"></a>return f / float64pow10[d.nd-d.dp], true;
    <a id="L273"></a>}
    <a id="L274"></a>return;
<a id="L275"></a>}

<a id="L277"></a><span class="comment">// If possible to convert decimal d to 32-bit float f exactly,</span>
<a id="L278"></a><span class="comment">// entirely in floating-point math, do so, avoiding the machinery above.</span>
<a id="L279"></a>func decimalAtof32(neg bool, d *decimal, trunc bool) (f float32, ok bool) {
    <a id="L280"></a><span class="comment">// Exact integers are &lt;= 10^7.</span>
    <a id="L281"></a><span class="comment">// Exact powers of ten are &lt;= 10^10.</span>
    <a id="L282"></a>if d.nd &gt; 7 {
        <a id="L283"></a>return
    <a id="L284"></a>}
    <a id="L285"></a>switch {
    <a id="L286"></a>case d.dp == d.nd: <span class="comment">// int</span>
        <a id="L287"></a>f := decimalAtof32Int(neg, d);
        <a id="L288"></a>return f, true;

    <a id="L290"></a>case d.dp &gt; d.nd &amp;&amp; d.dp &lt;= 7+10: <span class="comment">// int * 10^k</span>
        <a id="L291"></a>f := decimalAtof32Int(neg, d);
        <a id="L292"></a>k := d.dp - d.nd;
        <a id="L293"></a><span class="comment">// If exponent is big but number of digits is not,</span>
        <a id="L294"></a><span class="comment">// can move a few zeros into the integer part.</span>
        <a id="L295"></a>if k &gt; 10 {
            <a id="L296"></a>f *= float32pow10[k-10];
            <a id="L297"></a>k = 10;
        <a id="L298"></a>}
        <a id="L299"></a>return f * float32pow10[k], true;

    <a id="L301"></a>case d.dp &lt; d.nd &amp;&amp; d.nd-d.dp &lt;= 10: <span class="comment">// int / 10^k</span>
        <a id="L302"></a>f := decimalAtof32Int(neg, d);
        <a id="L303"></a>return f / float32pow10[d.nd-d.dp], true;
    <a id="L304"></a>}
    <a id="L305"></a>return;
<a id="L306"></a>}

<a id="L308"></a><span class="comment">// Atof32 converts the string s to a 32-bit floating-point number.</span>
<a id="L309"></a><span class="comment">//</span>
<a id="L310"></a><span class="comment">// If s is well-formed and near a valid floating point number,</span>
<a id="L311"></a><span class="comment">// Atof32 returns the nearest floating point number rounded</span>
<a id="L312"></a><span class="comment">// using IEEE754 unbiased rounding.</span>
<a id="L313"></a><span class="comment">//</span>
<a id="L314"></a><span class="comment">// The errors that Atof32 returns have concrete type *NumError</span>
<a id="L315"></a><span class="comment">// and include err.Num = s.</span>
<a id="L316"></a><span class="comment">//</span>
<a id="L317"></a><span class="comment">// If s is not syntactically well-formed, Atof32 returns err.Error = os.EINVAL.</span>
<a id="L318"></a><span class="comment">//</span>
<a id="L319"></a><span class="comment">// If s is syntactically well-formed but is more than 1/2 ULP</span>
<a id="L320"></a><span class="comment">// away from the largest floating point number of the given size,</span>
<a id="L321"></a><span class="comment">// Atof32 returns f = ±Inf, err.Error = os.ERANGE.</span>
<a id="L322"></a>func Atof32(s string) (f float32, err os.Error) {
    <a id="L323"></a>neg, d, trunc, ok := stringToDecimal(s);
    <a id="L324"></a>if !ok {
        <a id="L325"></a>return 0, &amp;NumError{s, os.EINVAL}
    <a id="L326"></a>}
    <a id="L327"></a>if optimize {
        <a id="L328"></a>if f, ok := decimalAtof32(neg, d, trunc); ok {
            <a id="L329"></a>return f, nil
        <a id="L330"></a>}
    <a id="L331"></a>}
    <a id="L332"></a>b, ovf := decimalToFloatBits(neg, d, trunc, &amp;float32info);
    <a id="L333"></a>f = math.Float32frombits(uint32(b));
    <a id="L334"></a>if ovf {
        <a id="L335"></a>err = &amp;NumError{s, os.ERANGE}
    <a id="L336"></a>}
    <a id="L337"></a>return f, err;
<a id="L338"></a>}

<a id="L340"></a><span class="comment">// Atof64 converts the string s to a 64-bit floating-point number.</span>
<a id="L341"></a><span class="comment">// Except for the type of its result, its definition is the same as that</span>
<a id="L342"></a><span class="comment">// of Atof32.</span>
<a id="L343"></a>func Atof64(s string) (f float64, err os.Error) {
    <a id="L344"></a>neg, d, trunc, ok := stringToDecimal(s);
    <a id="L345"></a>if !ok {
        <a id="L346"></a>return 0, &amp;NumError{s, os.EINVAL}
    <a id="L347"></a>}
    <a id="L348"></a>if optimize {
        <a id="L349"></a>if f, ok := decimalAtof64(neg, d, trunc); ok {
            <a id="L350"></a>return f, nil
        <a id="L351"></a>}
    <a id="L352"></a>}
    <a id="L353"></a>b, ovf := decimalToFloatBits(neg, d, trunc, &amp;float64info);
    <a id="L354"></a>f = math.Float64frombits(b);
    <a id="L355"></a>if ovf {
        <a id="L356"></a>err = &amp;NumError{s, os.ERANGE}
    <a id="L357"></a>}
    <a id="L358"></a>return f, err;
<a id="L359"></a>}

<a id="L361"></a><span class="comment">// Atof is like Atof32 or Atof64, depending on the size of float.</span>
<a id="L362"></a>func Atof(s string) (f float, err os.Error) {
    <a id="L363"></a>if FloatSize == 32 {
        <a id="L364"></a>f1, err1 := Atof32(s);
        <a id="L365"></a>return float(f1), err1;
    <a id="L366"></a>}
    <a id="L367"></a>f1, err1 := Atof64(s);
    <a id="L368"></a>return float(f1), err1;
<a id="L369"></a>}
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
