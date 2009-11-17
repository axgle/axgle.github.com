<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bignum/nrdiv_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/bignum/nrdiv_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file implements Newton-Raphson division and uses</span>
<a id="L6"></a><span class="comment">// it as an additional test case for bignum.</span>
<a id="L7"></a><span class="comment">//</span>
<a id="L8"></a><span class="comment">// Division of x/y is achieved by computing r = 1/y to</span>
<a id="L9"></a><span class="comment">// obtain the quotient q = x*r = x*(1/y) = x/y. The</span>
<a id="L10"></a><span class="comment">// reciprocal r is the solution for f(x) = 1/x - y and</span>
<a id="L11"></a><span class="comment">// the solution is approximated through iteration. The</span>
<a id="L12"></a><span class="comment">// iteration does not require division.</span>

<a id="L14"></a>package bignum

<a id="L16"></a>import &#34;testing&#34;


<a id="L19"></a><span class="comment">// An fpNat is a Natural scaled by a power of two</span>
<a id="L20"></a><span class="comment">// (an unsigned floating point representation). The</span>
<a id="L21"></a><span class="comment">// value of an fpNat x is x.m * 2^x.e .</span>
<a id="L22"></a><span class="comment">//</span>
<a id="L23"></a>type fpNat struct {
    <a id="L24"></a>m   Natural;
    <a id="L25"></a>e   int;
<a id="L26"></a>}


<a id="L29"></a><span class="comment">// sub computes x - y.</span>
<a id="L30"></a>func (x fpNat) sub(y fpNat) fpNat {
    <a id="L31"></a>switch d := x.e - y.e; {
    <a id="L32"></a>case d &lt; 0:
        <a id="L33"></a>return fpNat{x.m.Sub(y.m.Shl(uint(-d))), x.e}
    <a id="L34"></a>case d &gt; 0:
        <a id="L35"></a>return fpNat{x.m.Shl(uint(d)).Sub(y.m), y.e}
    <a id="L36"></a>}
    <a id="L37"></a>return fpNat{x.m.Sub(y.m), x.e};
<a id="L38"></a>}


<a id="L41"></a><span class="comment">// mul2 computes x*2.</span>
<a id="L42"></a>func (x fpNat) mul2() fpNat { return fpNat{x.m, x.e + 1} }


<a id="L45"></a><span class="comment">// mul computes x*y.</span>
<a id="L46"></a>func (x fpNat) mul(y fpNat) fpNat { return fpNat{x.m.Mul(y.m), x.e + y.e} }


<a id="L49"></a><span class="comment">// mant computes the (possibly truncated) Natural representation</span>
<a id="L50"></a><span class="comment">// of an fpNat x.</span>
<a id="L51"></a><span class="comment">//</span>
<a id="L52"></a>func (x fpNat) mant() Natural {
    <a id="L53"></a>switch {
    <a id="L54"></a>case x.e &gt; 0:
        <a id="L55"></a>return x.m.Shl(uint(x.e))
    <a id="L56"></a>case x.e &lt; 0:
        <a id="L57"></a>return x.m.Shr(uint(-x.e))
    <a id="L58"></a>}
    <a id="L59"></a>return x.m;
<a id="L60"></a>}


<a id="L63"></a><span class="comment">// nrDivEst computes an estimate of the quotient q = x0/y0 and returns q.</span>
<a id="L64"></a><span class="comment">// q may be too small (usually by 1).</span>
<a id="L65"></a><span class="comment">//</span>
<a id="L66"></a>func nrDivEst(x0, y0 Natural) Natural {
    <a id="L67"></a>if y0.IsZero() {
        <a id="L68"></a>panic(&#34;division by zero&#34;);
        <a id="L69"></a>return nil;
    <a id="L70"></a>}
    <a id="L71"></a><span class="comment">// y0 &gt; 0</span>

    <a id="L73"></a>if y0.Cmp(Nat(1)) == 0 {
        <a id="L74"></a>return x0
    <a id="L75"></a>}
    <a id="L76"></a><span class="comment">// y0 &gt; 1</span>

    <a id="L78"></a>switch d := x0.Cmp(y0); {
    <a id="L79"></a>case d &lt; 0:
        <a id="L80"></a>return Nat(0)
    <a id="L81"></a>case d == 0:
        <a id="L82"></a>return Nat(1)
    <a id="L83"></a>}
    <a id="L84"></a><span class="comment">// x0 &gt; y0 &gt; 1</span>

    <a id="L86"></a><span class="comment">// Determine maximum result length.</span>
    <a id="L87"></a>maxLen := int(x0.Log2() - y0.Log2() + 1);

    <a id="L89"></a><span class="comment">// In the following, each number x is represented</span>
    <a id="L90"></a><span class="comment">// as a mantissa x.m and an exponent x.e such that</span>
    <a id="L91"></a><span class="comment">// x = xm * 2^x.e.</span>
    <a id="L92"></a>x := fpNat{x0, 0};
    <a id="L93"></a>y := fpNat{y0, 0};

    <a id="L95"></a><span class="comment">// Determine a scale factor f = 2^e such that</span>
    <a id="L96"></a><span class="comment">// 0.5 &lt;= y/f == y*(2^-e) &lt; 1.0</span>
    <a id="L97"></a><span class="comment">// and scale y accordingly.</span>
    <a id="L98"></a>e := int(y.m.Log2()) + 1;
    <a id="L99"></a>y.e -= e;

    <a id="L101"></a><span class="comment">// t1</span>
    <a id="L102"></a>var c = 2.9142;
    <a id="L103"></a>const n = 14;
    <a id="L104"></a>t1 := fpNat{Nat(uint64(c * (1 &lt;&lt; n))), -n};

    <a id="L106"></a><span class="comment">// Compute initial value r0 for the reciprocal of y/f.</span>
    <a id="L107"></a><span class="comment">// r0 = t1 - 2*y</span>
    <a id="L108"></a>r := t1.sub(y.mul2());
    <a id="L109"></a>two := fpNat{Nat(2), 0};

    <a id="L111"></a><span class="comment">// Newton-Raphson iteration</span>
    <a id="L112"></a>p := Nat(0);
    <a id="L113"></a>for i := 0; ; i++ {
        <a id="L114"></a><span class="comment">// check if we are done</span>
        <a id="L115"></a><span class="comment">// TODO: Need to come up with a better test here</span>
        <a id="L116"></a><span class="comment">//       as it will reduce computation time significantly.</span>
        <a id="L117"></a><span class="comment">// q = x*r/f</span>
        <a id="L118"></a>q := x.mul(r);
        <a id="L119"></a>q.e -= e;
        <a id="L120"></a>res := q.mant();
        <a id="L121"></a>if res.Cmp(p) == 0 {
            <a id="L122"></a>return res
        <a id="L123"></a>}
        <a id="L124"></a>p = res;

        <a id="L126"></a><span class="comment">// r&#39; = r*(2 - y*r)</span>
        <a id="L127"></a>r = r.mul(two.sub(y.mul(r)));

        <a id="L129"></a><span class="comment">// reduce mantissa size</span>
        <a id="L130"></a><span class="comment">// TODO: Find smaller bound as it will reduce</span>
        <a id="L131"></a><span class="comment">//       computation time massively.</span>
        <a id="L132"></a>d := int(r.m.Log2()+1) - maxLen;
        <a id="L133"></a>if d &gt; 0 {
            <a id="L134"></a>r = fpNat{r.m.Shr(uint(d)), r.e + d}
        <a id="L135"></a>}
    <a id="L136"></a>}

    <a id="L138"></a>panic(&#34;unreachable&#34;);
    <a id="L139"></a>return nil;
<a id="L140"></a>}


<a id="L143"></a>func nrdiv(x, y Natural) (q, r Natural) {
    <a id="L144"></a>q = nrDivEst(x, y);
    <a id="L145"></a>r = x.Sub(y.Mul(q));
    <a id="L146"></a><span class="comment">// if r is too large, correct q and r</span>
    <a id="L147"></a><span class="comment">// (usually one iteration)</span>
    <a id="L148"></a>for r.Cmp(y) &gt;= 0 {
        <a id="L149"></a>q = q.Add(Nat(1));
        <a id="L150"></a>r = r.Sub(y);
    <a id="L151"></a>}
    <a id="L152"></a>return;
<a id="L153"></a>}


<a id="L156"></a>func div(t *testing.T, x, y Natural) {
    <a id="L157"></a>q, r := nrdiv(x, y);
    <a id="L158"></a>qx, rx := x.DivMod(y);
    <a id="L159"></a>if q.Cmp(qx) != 0 {
        <a id="L160"></a>t.Errorf(&#34;x = %s, y = %s, got q = %s, want q = %s&#34;, x, y, q, qx)
    <a id="L161"></a>}
    <a id="L162"></a>if r.Cmp(rx) != 0 {
        <a id="L163"></a>t.Errorf(&#34;x = %s, y = %s, got r = %s, want r = %s&#34;, x, y, r, rx)
    <a id="L164"></a>}
<a id="L165"></a>}


<a id="L168"></a>func idiv(t *testing.T, x0, y0 uint64) { div(t, Nat(x0), Nat(y0)) }


<a id="L171"></a>func TestNRDiv(t *testing.T) {
    <a id="L172"></a>idiv(t, 17, 18);
    <a id="L173"></a>idiv(t, 17, 17);
    <a id="L174"></a>idiv(t, 17, 1);
    <a id="L175"></a>idiv(t, 17, 16);
    <a id="L176"></a>idiv(t, 17, 10);
    <a id="L177"></a>idiv(t, 17, 9);
    <a id="L178"></a>idiv(t, 17, 8);
    <a id="L179"></a>idiv(t, 17, 5);
    <a id="L180"></a>idiv(t, 17, 3);
    <a id="L181"></a>idiv(t, 1025, 512);
    <a id="L182"></a>idiv(t, 7489595, 2);
    <a id="L183"></a>idiv(t, 5404679459, 78495);
    <a id="L184"></a>idiv(t, 7484890589595, 7484890589594);
    <a id="L185"></a>div(t, Fact(100), Fact(91));
    <a id="L186"></a>div(t, Fact(1000), Fact(991));
    <a id="L187"></a><span class="comment">//div(t, Fact(10000), Fact(9991));  // takes too long - disabled for now</span>
<a id="L188"></a>}
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
