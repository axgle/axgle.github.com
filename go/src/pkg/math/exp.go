<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/math/exp.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/math/exp.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package math


<a id="L8"></a><span class="comment">// The original C code, the long comment, and the constants</span>
<a id="L9"></a><span class="comment">// below are from FreeBSD&#39;s /usr/src/lib/msun/src/e_exp.c</span>
<a id="L10"></a><span class="comment">// and came with this notice.  The go code is a simplified</span>
<a id="L11"></a><span class="comment">// version of the original C.</span>
<a id="L12"></a><span class="comment">//</span>
<a id="L13"></a><span class="comment">// ====================================================</span>
<a id="L14"></a><span class="comment">// Copyright (C) 2004 by Sun Microsystems, Inc. All rights reserved.</span>
<a id="L15"></a><span class="comment">//</span>
<a id="L16"></a><span class="comment">// Permission to use, copy, modify, and distribute this</span>
<a id="L17"></a><span class="comment">// software is freely granted, provided that this notice</span>
<a id="L18"></a><span class="comment">// is preserved.</span>
<a id="L19"></a><span class="comment">// ====================================================</span>
<a id="L20"></a><span class="comment">//</span>
<a id="L21"></a><span class="comment">//</span>
<a id="L22"></a><span class="comment">// exp(x)</span>
<a id="L23"></a><span class="comment">// Returns the exponential of x.</span>
<a id="L24"></a><span class="comment">//</span>
<a id="L25"></a><span class="comment">// Method</span>
<a id="L26"></a><span class="comment">//   1. Argument reduction:</span>
<a id="L27"></a><span class="comment">//      Reduce x to an r so that |r| &lt;= 0.5*ln2 ~ 0.34658.</span>
<a id="L28"></a><span class="comment">//      Given x, find r and integer k such that</span>
<a id="L29"></a><span class="comment">//</span>
<a id="L30"></a><span class="comment">//               x = k*ln2 + r,  |r| &lt;= 0.5*ln2.</span>
<a id="L31"></a><span class="comment">//</span>
<a id="L32"></a><span class="comment">//      Here r will be represented as r = hi-lo for better</span>
<a id="L33"></a><span class="comment">//      accuracy.</span>
<a id="L34"></a><span class="comment">//</span>
<a id="L35"></a><span class="comment">//   2. Approximation of exp(r) by a special rational function on</span>
<a id="L36"></a><span class="comment">//      the interval [0,0.34658]:</span>
<a id="L37"></a><span class="comment">//      Write</span>
<a id="L38"></a><span class="comment">//          R(r**2) = r*(exp(r)+1)/(exp(r)-1) = 2 + r*r/6 - r**4/360 + ...</span>
<a id="L39"></a><span class="comment">//      We use a special Remes algorithm on [0,0.34658] to generate</span>
<a id="L40"></a><span class="comment">//      a polynomial of degree 5 to approximate R. The maximum error</span>
<a id="L41"></a><span class="comment">//      of this polynomial approximation is bounded by 2**-59. In</span>
<a id="L42"></a><span class="comment">//      other words,</span>
<a id="L43"></a><span class="comment">//          R(z) ~ 2.0 + P1*z + P2*z**2 + P3*z**3 + P4*z**4 + P5*z**5</span>
<a id="L44"></a><span class="comment">//      (where z=r*r, and the values of P1 to P5 are listed below)</span>
<a id="L45"></a><span class="comment">//      and</span>
<a id="L46"></a><span class="comment">//          |                  5          |     -59</span>
<a id="L47"></a><span class="comment">//          | 2.0+P1*z+...+P5*z   -  R(z) | &lt;= 2</span>
<a id="L48"></a><span class="comment">//          |                             |</span>
<a id="L49"></a><span class="comment">//      The computation of exp(r) thus becomes</span>
<a id="L50"></a><span class="comment">//                             2*r</span>
<a id="L51"></a><span class="comment">//              exp(r) = 1 + -------</span>
<a id="L52"></a><span class="comment">//                            R - r</span>
<a id="L53"></a><span class="comment">//                                 r*R1(r)</span>
<a id="L54"></a><span class="comment">//                     = 1 + r + ----------- (for better accuracy)</span>
<a id="L55"></a><span class="comment">//                                2 - R1(r)</span>
<a id="L56"></a><span class="comment">//      where</span>
<a id="L57"></a><span class="comment">//                               2       4             10</span>
<a id="L58"></a><span class="comment">//              R1(r) = r - (P1*r  + P2*r  + ... + P5*r   ).</span>
<a id="L59"></a><span class="comment">//</span>
<a id="L60"></a><span class="comment">//   3. Scale back to obtain exp(x):</span>
<a id="L61"></a><span class="comment">//      From step 1, we have</span>
<a id="L62"></a><span class="comment">//         exp(x) = 2^k * exp(r)</span>
<a id="L63"></a><span class="comment">//</span>
<a id="L64"></a><span class="comment">// Special cases:</span>
<a id="L65"></a><span class="comment">//      exp(INF) is INF, exp(NaN) is NaN;</span>
<a id="L66"></a><span class="comment">//      exp(-INF) is 0, and</span>
<a id="L67"></a><span class="comment">//      for finite argument, only exp(0)=1 is exact.</span>
<a id="L68"></a><span class="comment">//</span>
<a id="L69"></a><span class="comment">// Accuracy:</span>
<a id="L70"></a><span class="comment">//      according to an error analysis, the error is always less than</span>
<a id="L71"></a><span class="comment">//      1 ulp (unit in the last place).</span>
<a id="L72"></a><span class="comment">//</span>
<a id="L73"></a><span class="comment">// Misc. info.</span>
<a id="L74"></a><span class="comment">//      For IEEE double</span>
<a id="L75"></a><span class="comment">//          if x &gt;  7.09782712893383973096e+02 then exp(x) overflow</span>
<a id="L76"></a><span class="comment">//          if x &lt; -7.45133219101941108420e+02 then exp(x) underflow</span>
<a id="L77"></a><span class="comment">//</span>
<a id="L78"></a><span class="comment">// Constants:</span>
<a id="L79"></a><span class="comment">// The hexadecimal values are the intended ones for the following</span>
<a id="L80"></a><span class="comment">// constants. The decimal values may be used, provided that the</span>
<a id="L81"></a><span class="comment">// compiler will convert from decimal to binary accurately enough</span>
<a id="L82"></a><span class="comment">// to produce the hexadecimal values shown.</span>

<a id="L84"></a><span class="comment">// Exp returns e^x, the base-e exponential of x.</span>
<a id="L85"></a><span class="comment">//</span>
<a id="L86"></a><span class="comment">// Special cases are:</span>
<a id="L87"></a><span class="comment">//	Exp(+Inf) = +Inf</span>
<a id="L88"></a><span class="comment">//	Exp(NaN) = NaN</span>
<a id="L89"></a><span class="comment">// Very large values overflow to -Inf or +Inf.</span>
<a id="L90"></a><span class="comment">// Very small values underflow to 1.</span>
<a id="L91"></a>func Exp(x float64) float64 {
    <a id="L92"></a>const (
        <a id="L93"></a>Ln2Hi = 6.93147180369123816490e-01;
        <a id="L94"></a>Ln2Lo = 1.90821492927058770002e-10;
        <a id="L95"></a>Log2e = 1.44269504088896338700e+00;
        <a id="L96"></a>P1    = 1.66666666666666019037e-01;  <span class="comment">/* 0x3FC55555; 0x5555553E */</span>
        <a id="L97"></a>P2    = -2.77777777770155933842e-03; <span class="comment">/* 0xBF66C16C; 0x16BEBD93 */</span>
        <a id="L98"></a>P3    = 6.61375632143793436117e-05;  <span class="comment">/* 0x3F11566A; 0xAF25DE2C */</span>
        <a id="L99"></a>P4    = -1.65339022054652515390e-06; <span class="comment">/* 0xBEBBBD41; 0xC5D26BF1 */</span>
        <a id="L100"></a>P5    = 4.13813679705723846039e-08;  <span class="comment">/* 0x3E663769; 0x72BEA4D0 */</span>

        <a id="L102"></a>Overflow  = 7.09782712893383973096e+02;
        <a id="L103"></a>Underflow = -7.45133219101941108420e+02;
        <a id="L104"></a>NearZero  = 1.0 / (1 &lt;&lt; 28); <span class="comment">// 2^-28</span>
    <a id="L105"></a>)

    <a id="L107"></a><span class="comment">// special cases</span>
    <a id="L108"></a>switch {
    <a id="L109"></a>case IsNaN(x) || IsInf(x, 1):
        <a id="L110"></a>return x
    <a id="L111"></a>case IsInf(x, -1):
        <a id="L112"></a>return 0
    <a id="L113"></a>case x &gt; Overflow:
        <a id="L114"></a>return Inf(1)
    <a id="L115"></a>case x &lt; Underflow:
        <a id="L116"></a>return 0
    <a id="L117"></a>case -NearZero &lt; x &amp;&amp; x &lt; NearZero:
        <a id="L118"></a>return 1
    <a id="L119"></a>}

    <a id="L121"></a><span class="comment">// reduce; computed as r = hi - lo for extra precision.</span>
    <a id="L122"></a>var k int;
    <a id="L123"></a>switch {
    <a id="L124"></a>case x &lt; 0:
        <a id="L125"></a>k = int(Log2e*x - 0.5)
    <a id="L126"></a>case x &gt; 0:
        <a id="L127"></a>k = int(Log2e*x + 0.5)
    <a id="L128"></a>}
    <a id="L129"></a>hi := x - float64(k)*Ln2Hi;
    <a id="L130"></a>lo := float64(k) * Ln2Lo;
    <a id="L131"></a>r := hi - lo;

    <a id="L133"></a><span class="comment">// compute</span>
    <a id="L134"></a>t := r * r;
    <a id="L135"></a>c := r - t*(P1+t*(P2+t*(P3+t*(P4+t*P5))));
    <a id="L136"></a>y := 1 - ((lo - (r*c)/(2-c)) - hi);
    <a id="L137"></a><span class="comment">// TODO(rsc): make sure Ldexp can handle boundary k</span>
    <a id="L138"></a>return Ldexp(y, k);
<a id="L139"></a>}
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
