<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/math/log.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/math/log.go</h1>

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
<a id="L9"></a><span class="comment">// below are from FreeBSD&#39;s /usr/src/lib/msun/src/e_log.c</span>
<a id="L10"></a><span class="comment">// and came with this notice.  The go code is a simpler</span>
<a id="L11"></a><span class="comment">// version of the original C.</span>
<a id="L12"></a><span class="comment">//</span>
<a id="L13"></a><span class="comment">// ====================================================</span>
<a id="L14"></a><span class="comment">// Copyright (C) 1993 by Sun Microsystems, Inc. All rights reserved.</span>
<a id="L15"></a><span class="comment">//</span>
<a id="L16"></a><span class="comment">// Developed at SunPro, a Sun Microsystems, Inc. business.</span>
<a id="L17"></a><span class="comment">// Permission to use, copy, modify, and distribute this</span>
<a id="L18"></a><span class="comment">// software is freely granted, provided that this notice</span>
<a id="L19"></a><span class="comment">// is preserved.</span>
<a id="L20"></a><span class="comment">// ====================================================</span>
<a id="L21"></a><span class="comment">//</span>
<a id="L22"></a><span class="comment">// __ieee754_log(x)</span>
<a id="L23"></a><span class="comment">// Return the logrithm of x</span>
<a id="L24"></a><span class="comment">//</span>
<a id="L25"></a><span class="comment">// Method :</span>
<a id="L26"></a><span class="comment">//   1. Argument Reduction: find k and f such that</span>
<a id="L27"></a><span class="comment">//			x = 2^k * (1+f),</span>
<a id="L28"></a><span class="comment">//	   where  sqrt(2)/2 &lt; 1+f &lt; sqrt(2) .</span>
<a id="L29"></a><span class="comment">//</span>
<a id="L30"></a><span class="comment">//   2. Approximation of log(1+f).</span>
<a id="L31"></a><span class="comment">//	Let s = f/(2+f) ; based on log(1+f) = log(1+s) - log(1-s)</span>
<a id="L32"></a><span class="comment">//		 = 2s + 2/3 s**3 + 2/5 s**5 + .....,</span>
<a id="L33"></a><span class="comment">//	     	 = 2s + s*R</span>
<a id="L34"></a><span class="comment">//      We use a special Reme algorithm on [0,0.1716] to generate</span>
<a id="L35"></a><span class="comment">//	a polynomial of degree 14 to approximate R.  The maximum error</span>
<a id="L36"></a><span class="comment">//	of this polynomial approximation is bounded by 2**-58.45. In</span>
<a id="L37"></a><span class="comment">//	other words,</span>
<a id="L38"></a><span class="comment">//		        2      4      6      8      10      12      14</span>
<a id="L39"></a><span class="comment">//	    R(z) ~ L1*s +L2*s +L3*s +L4*s +L5*s  +L6*s  +L7*s</span>
<a id="L40"></a><span class="comment">//	(the values of L1 to L7 are listed in the program) and</span>
<a id="L41"></a><span class="comment">//	    |      2          14          |     -58.45</span>
<a id="L42"></a><span class="comment">//	    | L1*s +...+L7*s    -  R(z) | &lt;= 2</span>
<a id="L43"></a><span class="comment">//	    |                             |</span>
<a id="L44"></a><span class="comment">//	Note that 2s = f - s*f = f - hfsq + s*hfsq, where hfsq = f*f/2.</span>
<a id="L45"></a><span class="comment">//	In order to guarantee error in log below 1ulp, we compute log by</span>
<a id="L46"></a><span class="comment">//		log(1+f) = f - s*(f - R)		(if f is not too large)</span>
<a id="L47"></a><span class="comment">//		log(1+f) = f - (hfsq - s*(hfsq+R)).	(better accuracy)</span>
<a id="L48"></a><span class="comment">//</span>
<a id="L49"></a><span class="comment">//	3. Finally,  log(x) = k*Ln2 + log(1+f).</span>
<a id="L50"></a><span class="comment">//			    = k*Ln2_hi+(f-(hfsq-(s*(hfsq+R)+k*Ln2_lo)))</span>
<a id="L51"></a><span class="comment">//	   Here Ln2 is split into two floating point number:</span>
<a id="L52"></a><span class="comment">//			Ln2_hi + Ln2_lo,</span>
<a id="L53"></a><span class="comment">//	   where n*Ln2_hi is always exact for |n| &lt; 2000.</span>
<a id="L54"></a><span class="comment">//</span>
<a id="L55"></a><span class="comment">// Special cases:</span>
<a id="L56"></a><span class="comment">//	log(x) is NaN with signal if x &lt; 0 (including -INF) ;</span>
<a id="L57"></a><span class="comment">//	log(+INF) is +INF; log(0) is -INF with signal;</span>
<a id="L58"></a><span class="comment">//	log(NaN) is that NaN with no signal.</span>
<a id="L59"></a><span class="comment">//</span>
<a id="L60"></a><span class="comment">// Accuracy:</span>
<a id="L61"></a><span class="comment">//	according to an error analysis, the error is always less than</span>
<a id="L62"></a><span class="comment">//	1 ulp (unit in the last place).</span>
<a id="L63"></a><span class="comment">//</span>
<a id="L64"></a><span class="comment">// Constants:</span>
<a id="L65"></a><span class="comment">// The hexadecimal values are the intended ones for the following</span>
<a id="L66"></a><span class="comment">// constants. The decimal values may be used, provided that the</span>
<a id="L67"></a><span class="comment">// compiler will convert from decimal to binary accurately enough</span>
<a id="L68"></a><span class="comment">// to produce the hexadecimal values shown.</span>

<a id="L70"></a><span class="comment">// Log returns the natural logarithm of x.</span>
<a id="L71"></a><span class="comment">//</span>
<a id="L72"></a><span class="comment">// Special cases are:</span>
<a id="L73"></a><span class="comment">//	Log(+Inf) = +Inf</span>
<a id="L74"></a><span class="comment">//	Log(0) = -Inf</span>
<a id="L75"></a><span class="comment">//	Log(x &lt; 0) = NaN</span>
<a id="L76"></a><span class="comment">//	Log(NaN) = NaN</span>
<a id="L77"></a>func Log(x float64) float64 {
    <a id="L78"></a>const (
        <a id="L79"></a>Ln2Hi = 6.93147180369123816490e-01; <span class="comment">/* 3fe62e42 fee00000 */</span>
        <a id="L80"></a>Ln2Lo = 1.90821492927058770002e-10; <span class="comment">/* 3dea39ef 35793c76 */</span>
        <a id="L81"></a>L1    = 6.666666666666735130e-01;   <span class="comment">/* 3FE55555 55555593 */</span>
        <a id="L82"></a>L2    = 3.999999999940941908e-01;   <span class="comment">/* 3FD99999 9997FA04 */</span>
        <a id="L83"></a>L3    = 2.857142874366239149e-01;   <span class="comment">/* 3FD24924 94229359 */</span>
        <a id="L84"></a>L4    = 2.222219843214978396e-01;   <span class="comment">/* 3FCC71C5 1D8E78AF */</span>
        <a id="L85"></a>L5    = 1.818357216161805012e-01;   <span class="comment">/* 3FC74664 96CB03DE */</span>
        <a id="L86"></a>L6    = 1.531383769920937332e-01;   <span class="comment">/* 3FC39A09 D078C69F */</span>
        <a id="L87"></a>L7    = 1.479819860511658591e-01;   <span class="comment">/* 3FC2F112 DF3E5244 */</span>
    <a id="L88"></a>)

    <a id="L90"></a><span class="comment">// special cases</span>
    <a id="L91"></a>switch {
    <a id="L92"></a>case IsNaN(x) || IsInf(x, 1):
        <a id="L93"></a>return x
    <a id="L94"></a>case x &lt; 0:
        <a id="L95"></a>return NaN()
    <a id="L96"></a>case x == 0:
        <a id="L97"></a>return Inf(-1)
    <a id="L98"></a>}

    <a id="L100"></a><span class="comment">// reduce</span>
    <a id="L101"></a>f1, ki := Frexp(x);
    <a id="L102"></a>if f1 &lt; Sqrt2/2 {
        <a id="L103"></a>f1 *= 2;
        <a id="L104"></a>ki--;
    <a id="L105"></a>}
    <a id="L106"></a>f := f1 - 1;
    <a id="L107"></a>k := float64(ki);

    <a id="L109"></a><span class="comment">// compute</span>
    <a id="L110"></a>s := f / (2 + f);
    <a id="L111"></a>s2 := s * s;
    <a id="L112"></a>s4 := s2 * s2;
    <a id="L113"></a>t1 := s2 * (L1 + s4*(L3+s4*(L5+s4*L7)));
    <a id="L114"></a>t2 := s4 * (L2 + s4*(L4+s4*L6));
    <a id="L115"></a>R := t1 + t2;
    <a id="L116"></a>hfsq := 0.5 * f * f;
    <a id="L117"></a>return k*Ln2Hi - ((hfsq - (s*(hfsq+R) + k*Ln2Lo)) - f);
<a id="L118"></a>}

<a id="L120"></a><span class="comment">// Log10 returns the decimal logarithm of x.</span>
<a id="L121"></a><span class="comment">// The special cases are the same as for Log.</span>
<a id="L122"></a>func Log10(x float64) float64 {
    <a id="L123"></a>if x &lt;= 0 {
        <a id="L124"></a>return NaN()
    <a id="L125"></a>}
    <a id="L126"></a>return Log(x) * (1 / Ln10);
<a id="L127"></a>}
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
