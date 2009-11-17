<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/strconv/ftoa.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/strconv/ftoa.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Binary to decimal floating point conversion.</span>
<a id="L6"></a><span class="comment">// Algorithm:</span>
<a id="L7"></a><span class="comment">//   1) store mantissa in multiprecision decimal</span>
<a id="L8"></a><span class="comment">//   2) shift decimal by exponent</span>
<a id="L9"></a><span class="comment">//   3) read digits out &amp; format</span>

<a id="L11"></a>package strconv

<a id="L13"></a>import &#34;math&#34;

<a id="L15"></a><span class="comment">// TODO: move elsewhere?</span>
<a id="L16"></a>type floatInfo struct {
    <a id="L17"></a>mantbits uint;
    <a id="L18"></a>expbits  uint;
    <a id="L19"></a>bias     int;
<a id="L20"></a>}

<a id="L22"></a>var float32info = floatInfo{23, 8, -127}
<a id="L23"></a>var float64info = floatInfo{52, 11, -1023}

<a id="L25"></a>func floatsize() int {
    <a id="L26"></a><span class="comment">// Figure out whether float is float32 or float64.</span>
    <a id="L27"></a><span class="comment">// 1e-35 is representable in both, but 1e-70</span>
    <a id="L28"></a><span class="comment">// is too small for a float32.</span>
    <a id="L29"></a>var f float = 1e-35;
    <a id="L30"></a>if f*f == 0 {
        <a id="L31"></a>return 32
    <a id="L32"></a>}
    <a id="L33"></a>return 64;
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// Floatsize gives the size of the float type, either 32 or 64.</span>
<a id="L37"></a>var FloatSize = floatsize()

<a id="L39"></a><span class="comment">// Ftoa32 converts the 32-bit floating-point number f to a string,</span>
<a id="L40"></a><span class="comment">// according to the format fmt and precision prec.</span>
<a id="L41"></a><span class="comment">//</span>
<a id="L42"></a><span class="comment">// The format fmt is one of</span>
<a id="L43"></a><span class="comment">// &#39;b&#39; (-ddddp±ddd, a binary exponent),</span>
<a id="L44"></a><span class="comment">// &#39;e&#39; (-d.dddde±dd, a decimal exponent),</span>
<a id="L45"></a><span class="comment">// &#39;f&#39; (-ddd.dddd, no exponent), or</span>
<a id="L46"></a><span class="comment">// &#39;g&#39; (&#39;e&#39; for large exponents, &#39;f&#39; otherwise).</span>
<a id="L47"></a><span class="comment">//</span>
<a id="L48"></a><span class="comment">// The precision prec controls the number of digits</span>
<a id="L49"></a><span class="comment">// (excluding the exponent) printed by the &#39;e&#39;, &#39;f&#39;, and &#39;g&#39; formats.</span>
<a id="L50"></a><span class="comment">// For &#39;e&#39; and &#39;f&#39; it is the number of digits after the decimal point.</span>
<a id="L51"></a><span class="comment">// For &#39;g&#39; it is the total number of digits.</span>
<a id="L52"></a><span class="comment">// The special precision -1 uses the smallest number of digits</span>
<a id="L53"></a><span class="comment">// necessary such that Atof32 will return f exactly.</span>
<a id="L54"></a><span class="comment">//</span>
<a id="L55"></a><span class="comment">// Ftoa32(f) is not the same as Ftoa64(float32(f)),</span>
<a id="L56"></a><span class="comment">// because correct rounding and the number of digits</span>
<a id="L57"></a><span class="comment">// needed to identify f depend on the precision of the representation.</span>
<a id="L58"></a>func Ftoa32(f float32, fmt byte, prec int) string {
    <a id="L59"></a>return genericFtoa(uint64(math.Float32bits(f)), fmt, prec, &amp;float32info)
<a id="L60"></a>}

<a id="L62"></a><span class="comment">// Ftoa64 is like Ftoa32 but converts a 64-bit floating-point number.</span>
<a id="L63"></a>func Ftoa64(f float64, fmt byte, prec int) string {
    <a id="L64"></a>return genericFtoa(math.Float64bits(f), fmt, prec, &amp;float64info)
<a id="L65"></a>}

<a id="L67"></a><span class="comment">// Ftoa behaves as Ftoa32 or Ftoa64, depending on the size of the float type.</span>
<a id="L68"></a>func Ftoa(f float, fmt byte, prec int) string {
    <a id="L69"></a>if FloatSize == 32 {
        <a id="L70"></a>return Ftoa32(float32(f), fmt, prec)
    <a id="L71"></a>}
    <a id="L72"></a>return Ftoa64(float64(f), fmt, prec);
<a id="L73"></a>}

<a id="L75"></a>func genericFtoa(bits uint64, fmt byte, prec int, flt *floatInfo) string {
    <a id="L76"></a>neg := bits&gt;&gt;flt.expbits&gt;&gt;flt.mantbits != 0;
    <a id="L77"></a>exp := int(bits&gt;&gt;flt.mantbits) &amp; (1&lt;&lt;flt.expbits - 1);
    <a id="L78"></a>mant := bits &amp; (uint64(1)&lt;&lt;flt.mantbits - 1);

    <a id="L80"></a>switch exp {
    <a id="L81"></a>case 1&lt;&lt;flt.expbits - 1:
        <a id="L82"></a><span class="comment">// Inf, NaN</span>
        <a id="L83"></a>if mant != 0 {
            <a id="L84"></a>return &#34;NaN&#34;
        <a id="L85"></a>}
        <a id="L86"></a>if neg {
            <a id="L87"></a>return &#34;-Inf&#34;
        <a id="L88"></a>}
        <a id="L89"></a>return &#34;+Inf&#34;;

    <a id="L91"></a>case 0:
        <a id="L92"></a><span class="comment">// denormalized</span>
        <a id="L93"></a>exp++

    <a id="L95"></a>default:
        <a id="L96"></a><span class="comment">// add implicit top bit</span>
        <a id="L97"></a>mant |= uint64(1) &lt;&lt; flt.mantbits
    <a id="L98"></a>}
    <a id="L99"></a>exp += flt.bias;

    <a id="L101"></a><span class="comment">// Pick off easy binary format.</span>
    <a id="L102"></a>if fmt == &#39;b&#39; {
        <a id="L103"></a>return fmtB(neg, mant, exp, flt)
    <a id="L104"></a>}

    <a id="L106"></a><span class="comment">// Create exact decimal representation.</span>
    <a id="L107"></a><span class="comment">// The shift is exp - flt.mantbits because mant is a 1-bit integer</span>
    <a id="L108"></a><span class="comment">// followed by a flt.mantbits fraction, and we are treating it as</span>
    <a id="L109"></a><span class="comment">// a 1+flt.mantbits-bit integer.</span>
    <a id="L110"></a>d := newDecimal(mant).Shift(exp - int(flt.mantbits));

    <a id="L112"></a><span class="comment">// Round appropriately.</span>
    <a id="L113"></a><span class="comment">// Negative precision means &#34;only as much as needed to be exact.&#34;</span>
    <a id="L114"></a>shortest := false;
    <a id="L115"></a>if prec &lt; 0 {
        <a id="L116"></a>shortest = true;
        <a id="L117"></a>roundShortest(d, mant, exp, flt);
        <a id="L118"></a>switch fmt {
        <a id="L119"></a>case &#39;e&#39;, &#39;E&#39;:
            <a id="L120"></a>prec = d.nd - 1
        <a id="L121"></a>case &#39;f&#39;:
            <a id="L122"></a>prec = max(d.nd-d.dp, 0)
        <a id="L123"></a>case &#39;g&#39;, &#39;G&#39;:
            <a id="L124"></a>prec = d.nd
        <a id="L125"></a>}
    <a id="L126"></a>} else {
        <a id="L127"></a>switch fmt {
        <a id="L128"></a>case &#39;e&#39;, &#39;E&#39;:
            <a id="L129"></a>d.Round(prec + 1)
        <a id="L130"></a>case &#39;f&#39;:
            <a id="L131"></a>d.Round(d.dp + prec)
        <a id="L132"></a>case &#39;g&#39;, &#39;G&#39;:
            <a id="L133"></a>if prec == 0 {
                <a id="L134"></a>prec = 1
            <a id="L135"></a>}
            <a id="L136"></a>d.Round(prec);
        <a id="L137"></a>}
    <a id="L138"></a>}

    <a id="L140"></a>switch fmt {
    <a id="L141"></a>case &#39;e&#39;, &#39;E&#39;:
        <a id="L142"></a>return fmtE(neg, d, prec, fmt)
    <a id="L143"></a>case &#39;f&#39;:
        <a id="L144"></a>return fmtF(neg, d, prec)
    <a id="L145"></a>case &#39;g&#39;, &#39;G&#39;:
        <a id="L146"></a><span class="comment">// trailing zeros are removed.</span>
        <a id="L147"></a>if prec &gt; d.nd {
            <a id="L148"></a>prec = d.nd
        <a id="L149"></a>}
        <a id="L150"></a><span class="comment">// %e is used if the exponent from the conversion</span>
        <a id="L151"></a><span class="comment">// is less than -4 or greater than or equal to the precision.</span>
        <a id="L152"></a><span class="comment">// if precision was the shortest possible, use precision 6 for this decision.</span>
        <a id="L153"></a>eprec := prec;
        <a id="L154"></a>if shortest {
            <a id="L155"></a>eprec = 6
        <a id="L156"></a>}
        <a id="L157"></a>exp := d.dp - 1;
        <a id="L158"></a>if exp &lt; -4 || exp &gt;= eprec {
            <a id="L159"></a>return fmtE(neg, d, prec-1, fmt+&#39;e&#39;-&#39;g&#39;)
        <a id="L160"></a>}
        <a id="L161"></a>return fmtF(neg, d, max(prec-d.dp, 0));
    <a id="L162"></a>}

    <a id="L164"></a>return &#34;%&#34; + string(fmt);
<a id="L165"></a>}

<a id="L167"></a><span class="comment">// Round d (= mant * 2^exp) to the shortest number of digits</span>
<a id="L168"></a><span class="comment">// that will let the original floating point value be precisely</span>
<a id="L169"></a><span class="comment">// reconstructed.  Size is original floating point size (64 or 32).</span>
<a id="L170"></a>func roundShortest(d *decimal, mant uint64, exp int, flt *floatInfo) {
    <a id="L171"></a><span class="comment">// If mantissa is zero, the number is zero; stop now.</span>
    <a id="L172"></a>if mant == 0 {
        <a id="L173"></a>d.nd = 0;
        <a id="L174"></a>return;
    <a id="L175"></a>}

    <a id="L177"></a><span class="comment">// TODO(rsc): Unless exp == minexp, if the number of digits in d</span>
    <a id="L178"></a><span class="comment">// is less than 17, it seems likely that it would be</span>
    <a id="L179"></a><span class="comment">// the shortest possible number already.  So maybe we can</span>
    <a id="L180"></a><span class="comment">// bail out without doing the extra multiprecision math here.</span>

    <a id="L182"></a><span class="comment">// Compute upper and lower such that any decimal number</span>
    <a id="L183"></a><span class="comment">// between upper and lower (possibly inclusive)</span>
    <a id="L184"></a><span class="comment">// will round to the original floating point number.</span>

    <a id="L186"></a><span class="comment">// d = mant &lt;&lt; (exp - mantbits)</span>
    <a id="L187"></a><span class="comment">// Next highest floating point number is mant+1 &lt;&lt; exp-mantbits.</span>
    <a id="L188"></a><span class="comment">// Our upper bound is halfway inbetween, mant*2+1 &lt;&lt; exp-mantbits-1.</span>
    <a id="L189"></a>upper := newDecimal(mant*2 + 1).Shift(exp - int(flt.mantbits) - 1);

    <a id="L191"></a><span class="comment">// d = mant &lt;&lt; (exp - mantbits)</span>
    <a id="L192"></a><span class="comment">// Next lowest floating point number is mant-1 &lt;&lt; exp-mantbits,</span>
    <a id="L193"></a><span class="comment">// unless mant-1 drops the significant bit and exp is not the minimum exp,</span>
    <a id="L194"></a><span class="comment">// in which case the next lowest is mant*2-1 &lt;&lt; exp-mantbits-1.</span>
    <a id="L195"></a><span class="comment">// Either way, call it mantlo &lt;&lt; explo-mantbits.</span>
    <a id="L196"></a><span class="comment">// Our lower bound is halfway inbetween, mantlo*2+1 &lt;&lt; explo-mantbits-1.</span>
    <a id="L197"></a>minexp := flt.bias + 1; <span class="comment">// minimum possible exponent</span>
    <a id="L198"></a>var mantlo uint64;
    <a id="L199"></a>var explo int;
    <a id="L200"></a>if mant &gt; 1&lt;&lt;flt.mantbits || exp == minexp {
        <a id="L201"></a>mantlo = mant - 1;
        <a id="L202"></a>explo = exp;
    <a id="L203"></a>} else {
        <a id="L204"></a>mantlo = mant*2 - 1;
        <a id="L205"></a>explo = exp - 1;
    <a id="L206"></a>}
    <a id="L207"></a>lower := newDecimal(mantlo*2 + 1).Shift(explo - int(flt.mantbits) - 1);

    <a id="L209"></a><span class="comment">// The upper and lower bounds are possible outputs only if</span>
    <a id="L210"></a><span class="comment">// the original mantissa is even, so that IEEE round-to-even</span>
    <a id="L211"></a><span class="comment">// would round to the original mantissa and not the neighbors.</span>
    <a id="L212"></a>inclusive := mant%2 == 0;

    <a id="L214"></a><span class="comment">// Now we can figure out the minimum number of digits required.</span>
    <a id="L215"></a><span class="comment">// Walk along until d has distinguished itself from upper and lower.</span>
    <a id="L216"></a>for i := 0; i &lt; d.nd; i++ {
        <a id="L217"></a>var l, m, u byte; <span class="comment">// lower, middle, upper digits</span>
        <a id="L218"></a>if i &lt; lower.nd {
            <a id="L219"></a>l = lower.d[i]
        <a id="L220"></a>} else {
            <a id="L221"></a>l = &#39;0&#39;
        <a id="L222"></a>}
        <a id="L223"></a>m = d.d[i];
        <a id="L224"></a>if i &lt; upper.nd {
            <a id="L225"></a>u = upper.d[i]
        <a id="L226"></a>} else {
            <a id="L227"></a>u = &#39;0&#39;
        <a id="L228"></a>}

        <a id="L230"></a><span class="comment">// Okay to round down (truncate) if lower has a different digit</span>
        <a id="L231"></a><span class="comment">// or if lower is inclusive and is exactly the result of rounding down.</span>
        <a id="L232"></a>okdown := l != m || (inclusive &amp;&amp; l == m &amp;&amp; i+1 == lower.nd);

        <a id="L234"></a><span class="comment">// Okay to round up if upper has a different digit and</span>
        <a id="L235"></a><span class="comment">// either upper is inclusive or upper is bigger than the result of rounding up.</span>
        <a id="L236"></a>okup := m != u &amp;&amp; (inclusive || i+1 &lt; upper.nd);

        <a id="L238"></a><span class="comment">// If it&#39;s okay to do either, then round to the nearest one.</span>
        <a id="L239"></a><span class="comment">// If it&#39;s okay to do only one, do it.</span>
        <a id="L240"></a>switch {
        <a id="L241"></a>case okdown &amp;&amp; okup:
            <a id="L242"></a>d.Round(i + 1);
            <a id="L243"></a>return;
        <a id="L244"></a>case okdown:
            <a id="L245"></a>d.RoundDown(i + 1);
            <a id="L246"></a>return;
        <a id="L247"></a>case okup:
            <a id="L248"></a>d.RoundUp(i + 1);
            <a id="L249"></a>return;
        <a id="L250"></a>}
    <a id="L251"></a>}
<a id="L252"></a>}

<a id="L254"></a><span class="comment">// %e: -d.ddddde±dd</span>
<a id="L255"></a>func fmtE(neg bool, d *decimal, prec int, fmt byte) string {
    <a id="L256"></a>buf := make([]byte, 3+max(prec, 0)+30); <span class="comment">// &#34;-0.&#34; + prec digits + exp</span>
    <a id="L257"></a>w := 0;                                 <span class="comment">// write index</span>

    <a id="L259"></a><span class="comment">// sign</span>
    <a id="L260"></a>if neg {
        <a id="L261"></a>buf[w] = &#39;-&#39;;
        <a id="L262"></a>w++;
    <a id="L263"></a>}

    <a id="L265"></a><span class="comment">// first digit</span>
    <a id="L266"></a>if d.nd == 0 {
        <a id="L267"></a>buf[w] = &#39;0&#39;
    <a id="L268"></a>} else {
        <a id="L269"></a>buf[w] = d.d[0]
    <a id="L270"></a>}
    <a id="L271"></a>w++;

    <a id="L273"></a><span class="comment">// .moredigits</span>
    <a id="L274"></a>if prec &gt; 0 {
        <a id="L275"></a>buf[w] = &#39;.&#39;;
        <a id="L276"></a>w++;
        <a id="L277"></a>for i := 0; i &lt; prec; i++ {
            <a id="L278"></a>if 1+i &lt; d.nd {
                <a id="L279"></a>buf[w] = d.d[1+i]
            <a id="L280"></a>} else {
                <a id="L281"></a>buf[w] = &#39;0&#39;
            <a id="L282"></a>}
            <a id="L283"></a>w++;
        <a id="L284"></a>}
    <a id="L285"></a>}

    <a id="L287"></a><span class="comment">// e±</span>
    <a id="L288"></a>buf[w] = fmt;
    <a id="L289"></a>w++;
    <a id="L290"></a>exp := d.dp - 1;
    <a id="L291"></a>if d.nd == 0 { <span class="comment">// special case: 0 has exponent 0</span>
        <a id="L292"></a>exp = 0
    <a id="L293"></a>}
    <a id="L294"></a>if exp &lt; 0 {
        <a id="L295"></a>buf[w] = &#39;-&#39;;
        <a id="L296"></a>exp = -exp;
    <a id="L297"></a>} else {
        <a id="L298"></a>buf[w] = &#39;+&#39;
    <a id="L299"></a>}
    <a id="L300"></a>w++;

    <a id="L302"></a><span class="comment">// dddd</span>
    <a id="L303"></a><span class="comment">// count digits</span>
    <a id="L304"></a>n := 0;
    <a id="L305"></a>for e := exp; e &gt; 0; e /= 10 {
        <a id="L306"></a>n++
    <a id="L307"></a>}
    <a id="L308"></a><span class="comment">// leading zeros</span>
    <a id="L309"></a>for i := n; i &lt; 2; i++ {
        <a id="L310"></a>buf[w] = &#39;0&#39;;
        <a id="L311"></a>w++;
    <a id="L312"></a>}
    <a id="L313"></a><span class="comment">// digits</span>
    <a id="L314"></a>w += n;
    <a id="L315"></a>n = 0;
    <a id="L316"></a>for e := exp; e &gt; 0; e /= 10 {
        <a id="L317"></a>n++;
        <a id="L318"></a>buf[w-n] = byte(e%10 + &#39;0&#39;);
    <a id="L319"></a>}

    <a id="L321"></a>return string(buf[0:w]);
<a id="L322"></a>}

<a id="L324"></a><span class="comment">// %f: -ddddddd.ddddd</span>
<a id="L325"></a>func fmtF(neg bool, d *decimal, prec int) string {
    <a id="L326"></a>buf := make([]byte, 1+max(d.dp, 1)+1+max(prec, 0));
    <a id="L327"></a>w := 0;

    <a id="L329"></a><span class="comment">// sign</span>
    <a id="L330"></a>if neg {
        <a id="L331"></a>buf[w] = &#39;-&#39;;
        <a id="L332"></a>w++;
    <a id="L333"></a>}

    <a id="L335"></a><span class="comment">// integer, padded with zeros as needed.</span>
    <a id="L336"></a>if d.dp &gt; 0 {
        <a id="L337"></a>var i int;
        <a id="L338"></a>for i = 0; i &lt; d.dp &amp;&amp; i &lt; d.nd; i++ {
            <a id="L339"></a>buf[w] = d.d[i];
            <a id="L340"></a>w++;
        <a id="L341"></a>}
        <a id="L342"></a>for ; i &lt; d.dp; i++ {
            <a id="L343"></a>buf[w] = &#39;0&#39;;
            <a id="L344"></a>w++;
        <a id="L345"></a>}
    <a id="L346"></a>} else {
        <a id="L347"></a>buf[w] = &#39;0&#39;;
        <a id="L348"></a>w++;
    <a id="L349"></a>}

    <a id="L351"></a><span class="comment">// fraction</span>
    <a id="L352"></a>if prec &gt; 0 {
        <a id="L353"></a>buf[w] = &#39;.&#39;;
        <a id="L354"></a>w++;
        <a id="L355"></a>for i := 0; i &lt; prec; i++ {
            <a id="L356"></a>if d.dp+i &lt; 0 || d.dp+i &gt;= d.nd {
                <a id="L357"></a>buf[w] = &#39;0&#39;
            <a id="L358"></a>} else {
                <a id="L359"></a>buf[w] = d.d[d.dp+i]
            <a id="L360"></a>}
            <a id="L361"></a>w++;
        <a id="L362"></a>}
    <a id="L363"></a>}

    <a id="L365"></a>return string(buf[0:w]);
<a id="L366"></a>}

<a id="L368"></a><span class="comment">// %b: -ddddddddp+ddd</span>
<a id="L369"></a>func fmtB(neg bool, mant uint64, exp int, flt *floatInfo) string {
    <a id="L370"></a>var buf [50]byte;
    <a id="L371"></a>w := len(buf);
    <a id="L372"></a>exp -= int(flt.mantbits);
    <a id="L373"></a>esign := byte(&#39;+&#39;);
    <a id="L374"></a>if exp &lt; 0 {
        <a id="L375"></a>esign = &#39;-&#39;;
        <a id="L376"></a>exp = -exp;
    <a id="L377"></a>}
    <a id="L378"></a>n := 0;
    <a id="L379"></a>for exp &gt; 0 || n &lt; 1 {
        <a id="L380"></a>n++;
        <a id="L381"></a>w--;
        <a id="L382"></a>buf[w] = byte(exp%10 + &#39;0&#39;);
        <a id="L383"></a>exp /= 10;
    <a id="L384"></a>}
    <a id="L385"></a>w--;
    <a id="L386"></a>buf[w] = esign;
    <a id="L387"></a>w--;
    <a id="L388"></a>buf[w] = &#39;p&#39;;
    <a id="L389"></a>n = 0;
    <a id="L390"></a>for mant &gt; 0 || n &lt; 1 {
        <a id="L391"></a>n++;
        <a id="L392"></a>w--;
        <a id="L393"></a>buf[w] = byte(mant%10 + &#39;0&#39;);
        <a id="L394"></a>mant /= 10;
    <a id="L395"></a>}
    <a id="L396"></a>if neg {
        <a id="L397"></a>w--;
        <a id="L398"></a>buf[w] = &#39;-&#39;;
    <a id="L399"></a>}
    <a id="L400"></a>return string(buf[w:len(buf)]);
<a id="L401"></a>}

<a id="L403"></a>func max(a, b int) int {
    <a id="L404"></a>if a &gt; b {
        <a id="L405"></a>return a
    <a id="L406"></a>}
    <a id="L407"></a>return b;
<a id="L408"></a>}
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
