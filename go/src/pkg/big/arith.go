<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/big/arith.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/big/arith.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file provides Go implementations of elementary multi-precision</span>
<a id="L6"></a><span class="comment">// arithmetic operations on word vectors. Needed for platforms without</span>
<a id="L7"></a><span class="comment">// assembly implementations of these routines.</span>

<a id="L9"></a>package big

<a id="L11"></a>import &#34;unsafe&#34;

<a id="L13"></a>type Word uintptr

<a id="L15"></a>const (
    <a id="L16"></a>_S    = uintptr(unsafe.Sizeof(Word(0))); <span class="comment">// TODO(gri) should Sizeof return a uintptr?</span>
    <a id="L17"></a>_logW = (0x650 &gt;&gt; _S) &amp; 7;
    <a id="L18"></a>_W    = 1 &lt;&lt; _logW;
    <a id="L19"></a>_B    = 1 &lt;&lt; _W;
    <a id="L20"></a>_M    = _B - 1;
    <a id="L21"></a>_W2   = _W / 2;
    <a id="L22"></a>_B2   = 1 &lt;&lt; _W2;
    <a id="L23"></a>_M2   = _B2 - 1;
<a id="L24"></a>)


<a id="L27"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L28"></a><span class="comment">// Elementary operations on words</span>
<a id="L29"></a><span class="comment">//</span>
<a id="L30"></a><span class="comment">// These operations are used by the vector operations below.</span>

<a id="L32"></a><span class="comment">// z1&lt;&lt;_W + z0 = x+y+c, with c == 0 or 1</span>
<a id="L33"></a>func addWW_g(x, y, c Word) (z1, z0 Word) {
    <a id="L34"></a>yc := y + c;
    <a id="L35"></a>z0 = x + yc;
    <a id="L36"></a>if z0 &lt; x || yc &lt; y {
        <a id="L37"></a>z1 = 1
    <a id="L38"></a>}
    <a id="L39"></a>return;
<a id="L40"></a>}


<a id="L43"></a><span class="comment">// z1&lt;&lt;_W + z0 = x-y-c, with c == 0 or 1</span>
<a id="L44"></a>func subWW_g(x, y, c Word) (z1, z0 Word) {
    <a id="L45"></a>yc := y + c;
    <a id="L46"></a>z0 = x - yc;
    <a id="L47"></a>if z0 &gt; x || yc &lt; y {
        <a id="L48"></a>z1 = 1
    <a id="L49"></a>}
    <a id="L50"></a>return;
<a id="L51"></a>}


<a id="L54"></a><span class="comment">// z1&lt;&lt;_W + z0 = x*y</span>
<a id="L55"></a>func mulWW_g(x, y Word) (z1, z0 Word) {
    <a id="L56"></a><span class="comment">// Split x and y into 2 halfWords each, multiply</span>
    <a id="L57"></a><span class="comment">// the halfWords separately while avoiding overflow,</span>
    <a id="L58"></a><span class="comment">// and return the product as 2 Words.</span>

    <a id="L60"></a>if x &lt; y {
        <a id="L61"></a>x, y = y, x
    <a id="L62"></a>}

    <a id="L64"></a>if x &lt; _B2 {
        <a id="L65"></a><span class="comment">// y &lt; _B2 because y &lt;= x</span>
        <a id="L66"></a><span class="comment">// sub-digits of x and y are (0, x) and (0, y)</span>
        <a id="L67"></a><span class="comment">// z = z[0] = x*y</span>
        <a id="L68"></a>z0 = x * y;
        <a id="L69"></a>return;
    <a id="L70"></a>}

    <a id="L72"></a>if y &lt; _B2 {
        <a id="L73"></a><span class="comment">// sub-digits of x and y are (x1, x0) and (0, y)</span>
        <a id="L74"></a><span class="comment">// x = (x1*_B2 + x0)</span>
        <a id="L75"></a><span class="comment">// y = (y1*_B2 + y0)</span>
        <a id="L76"></a>x1, x0 := x&gt;&gt;_W2, x&amp;_M2;

        <a id="L78"></a><span class="comment">// x*y = t2*_B2*_B2 + t1*_B2 + t0</span>
        <a id="L79"></a>t0 := x0 * y;
        <a id="L80"></a>t1 := x1 * y;

        <a id="L82"></a><span class="comment">// compute result digits but avoid overflow</span>
        <a id="L83"></a><span class="comment">// z = z[1]*_B + z[0] = x*y</span>
        <a id="L84"></a>z0 = t1&lt;&lt;_W2 + t0;
        <a id="L85"></a>z1 = (t1 + t0&gt;&gt;_W2) &gt;&gt; _W2;
        <a id="L86"></a>return;
    <a id="L87"></a>}

    <a id="L89"></a><span class="comment">// general case</span>
    <a id="L90"></a><span class="comment">// sub-digits of x and y are (x1, x0) and (y1, y0)</span>
    <a id="L91"></a><span class="comment">// x = (x1*_B2 + x0)</span>
    <a id="L92"></a><span class="comment">// y = (y1*_B2 + y0)</span>
    <a id="L93"></a>x1, x0 := x&gt;&gt;_W2, x&amp;_M2;
    <a id="L94"></a>y1, y0 := y&gt;&gt;_W2, y&amp;_M2;

    <a id="L96"></a><span class="comment">// x*y = t2*_B2*_B2 + t1*_B2 + t0</span>
    <a id="L97"></a>t0 := x0 * y0;
    <a id="L98"></a><span class="comment">// t1 := x1*y0 + x0*y1;</span>
    <a id="L99"></a>var c Word;
    <a id="L100"></a>t1 := x1 * y0;
    <a id="L101"></a>t1a := t1;
    <a id="L102"></a>t1 += x0 * y1;
    <a id="L103"></a>if t1 &lt; t1a {
        <a id="L104"></a>c++
    <a id="L105"></a>}
    <a id="L106"></a>t2 := x1*y1 + c*_B2;

    <a id="L108"></a><span class="comment">// compute result digits but avoid overflow</span>
    <a id="L109"></a><span class="comment">// z = z[1]*_B + z[0] = x*y</span>
    <a id="L110"></a><span class="comment">// This may overflow, but that&#39;s ok because we also sum t1 and t0 above</span>
    <a id="L111"></a><span class="comment">// and we take care of the overflow there.</span>
    <a id="L112"></a>z0 = t1&lt;&lt;_W2 + t0;

    <a id="L114"></a><span class="comment">// z1 = t2 + (t1 + t0&gt;&gt;_W2)&gt;&gt;_W2;</span>
    <a id="L115"></a>var c3 Word;
    <a id="L116"></a>z1 = t1 + t0&gt;&gt;_W2;
    <a id="L117"></a>if z1 &lt; t1 {
        <a id="L118"></a>c3++
    <a id="L119"></a>}
    <a id="L120"></a>z1 &gt;&gt;= _W2;
    <a id="L121"></a>z1 += c3 * _B2;
    <a id="L122"></a>z1 += t2;
    <a id="L123"></a>return;
<a id="L124"></a>}


<a id="L127"></a><span class="comment">// z1&lt;&lt;_W + z0 = x*y + c</span>
<a id="L128"></a>func mulAddWWW_g(x, y, c Word) (z1, z0 Word) {
    <a id="L129"></a><span class="comment">// Split x and y into 2 halfWords each, multiply</span>
    <a id="L130"></a><span class="comment">// the halfWords separately while avoiding overflow,</span>
    <a id="L131"></a><span class="comment">// and return the product as 2 Words.</span>

    <a id="L133"></a><span class="comment">// TODO(gri) Should implement special cases for faster execution.</span>

    <a id="L135"></a><span class="comment">// general case</span>
    <a id="L136"></a><span class="comment">// sub-digits of x, y, and c are (x1, x0), (y1, y0), (c1, c0)</span>
    <a id="L137"></a><span class="comment">// x = (x1*_B2 + x0)</span>
    <a id="L138"></a><span class="comment">// y = (y1*_B2 + y0)</span>
    <a id="L139"></a>x1, x0 := x&gt;&gt;_W2, x&amp;_M2;
    <a id="L140"></a>y1, y0 := y&gt;&gt;_W2, y&amp;_M2;
    <a id="L141"></a>c1, c0 := c&gt;&gt;_W2, c&amp;_M2;

    <a id="L143"></a><span class="comment">// x*y + c = t2*_B2*_B2 + t1*_B2 + t0</span>
    <a id="L144"></a><span class="comment">// (1&lt;&lt;32-1)^2 == 1&lt;&lt;64 - 1&lt;&lt;33 + 1, so there&#39;s space to add c0 in here.</span>
    <a id="L145"></a>t0 := x0*y0 + c0;

    <a id="L147"></a><span class="comment">// t1 := x1*y0 + x0*y1 + c1;</span>
    <a id="L148"></a>var c2 Word; <span class="comment">// extra carry</span>
    <a id="L149"></a>t1 := x1*y0 + c1;
    <a id="L150"></a>t1a := t1;
    <a id="L151"></a>t1 += x0 * y1;
    <a id="L152"></a>if t1 &lt; t1a { <span class="comment">// If the number got smaller then we overflowed.</span>
        <a id="L153"></a>c2++
    <a id="L154"></a>}

    <a id="L156"></a>t2 := x1*y1 + c2*_B2;

    <a id="L158"></a><span class="comment">// compute result digits but avoid overflow</span>
    <a id="L159"></a><span class="comment">// z = z[1]*_B + z[0] = x*y</span>
    <a id="L160"></a><span class="comment">// z0 = t1&lt;&lt;_W2 + t0;</span>
    <a id="L161"></a><span class="comment">// This may overflow, but that&#39;s ok because we also sum t1 and t0 below</span>
    <a id="L162"></a><span class="comment">// and we take care of the overflow there.</span>
    <a id="L163"></a>z0 = t1&lt;&lt;_W2 + t0;

    <a id="L165"></a>var c3 Word;
    <a id="L166"></a>z1 = t1 + t0&gt;&gt;_W2;
    <a id="L167"></a>if z1 &lt; t1 {
        <a id="L168"></a>c3++
    <a id="L169"></a>}
    <a id="L170"></a>z1 &gt;&gt;= _W2;
    <a id="L171"></a>z1 += t2 + c3*_B2;

    <a id="L173"></a>return;
<a id="L174"></a>}


<a id="L177"></a><span class="comment">// q = (x1&lt;&lt;_W + x0 - r)/y</span>
<a id="L178"></a><span class="comment">// The most significant bit of y must be 1.</span>
<a id="L179"></a>func divStep(x1, x0, y Word) (q, r Word) {
    <a id="L180"></a>d1, d0 := y&gt;&gt;_W2, y&amp;_M2;
    <a id="L181"></a>q1, r1 := x1/d1, x1%d1;
    <a id="L182"></a>m := q1 * d0;
    <a id="L183"></a>r1 = r1*_B2 | x0&gt;&gt;_W2;
    <a id="L184"></a>if r1 &lt; m {
        <a id="L185"></a>q1--;
        <a id="L186"></a>r1 += y;
        <a id="L187"></a>if r1 &gt;= y &amp;&amp; r1 &lt; m {
            <a id="L188"></a>q1--;
            <a id="L189"></a>r1 += y;
        <a id="L190"></a>}
    <a id="L191"></a>}
    <a id="L192"></a>r1 -= m;

    <a id="L194"></a>r0 := r1 % d1;
    <a id="L195"></a>q0 := r1 / d1;
    <a id="L196"></a>m = q0 * d0;
    <a id="L197"></a>r0 = r0*_B2 | x0&amp;_M2;
    <a id="L198"></a>if r0 &lt; m {
        <a id="L199"></a>q0--;
        <a id="L200"></a>r0 += y;
        <a id="L201"></a>if r0 &gt;= y &amp;&amp; r0 &lt; m {
            <a id="L202"></a>q0--;
            <a id="L203"></a>r0 += y;
        <a id="L204"></a>}
    <a id="L205"></a>}
    <a id="L206"></a>r0 -= m;

    <a id="L208"></a>q = q1*_B2 | q0;
    <a id="L209"></a>r = r0;
    <a id="L210"></a>return;
<a id="L211"></a>}


<a id="L214"></a><span class="comment">// Number of leading zeros in x.</span>
<a id="L215"></a>func leadingZeros(x Word) (n uint) {
    <a id="L216"></a>if x == 0 {
        <a id="L217"></a>return _W
    <a id="L218"></a>}
    <a id="L219"></a>for x&amp;(1&lt;&lt;(_W-1)) == 0 {
        <a id="L220"></a>n++;
        <a id="L221"></a>x &lt;&lt;= 1;
    <a id="L222"></a>}
    <a id="L223"></a>return;
<a id="L224"></a>}


<a id="L227"></a><span class="comment">// q = (x1&lt;&lt;_W + x0 - r)/y</span>
<a id="L228"></a>func divWW_g(x1, x0, y Word) (q, r Word) {
    <a id="L229"></a>if x1 == 0 {
        <a id="L230"></a>q, r = x0/y, x0%y;
        <a id="L231"></a>return;
    <a id="L232"></a>}

    <a id="L234"></a>var q0, q1 Word;
    <a id="L235"></a>z := leadingZeros(y);
    <a id="L236"></a>if y &gt; x1 {
        <a id="L237"></a>if z != 0 {
            <a id="L238"></a>y &lt;&lt;= z;
            <a id="L239"></a>x1 = (x1 &lt;&lt; z) | (x0 &gt;&gt; (_W - z));
            <a id="L240"></a>x0 &lt;&lt;= z;
        <a id="L241"></a>}
        <a id="L242"></a>q0, x0 = divStep(x1, x0, y);
        <a id="L243"></a>q1 = 0;
    <a id="L244"></a>} else {
        <a id="L245"></a>if z == 0 {
            <a id="L246"></a>x1 -= y;
            <a id="L247"></a>q1 = 1;
        <a id="L248"></a>} else {
            <a id="L249"></a>z1 := _W - z;
            <a id="L250"></a>y &lt;&lt;= z;
            <a id="L251"></a>x2 := x1 &gt;&gt; z1;
            <a id="L252"></a>x1 = (x1 &lt;&lt; z) | (x0 &gt;&gt; z1);
            <a id="L253"></a>x0 &lt;&lt;= z;
            <a id="L254"></a>q1, x1 = divStep(x2, x1, y);
        <a id="L255"></a>}

        <a id="L257"></a>q0, x0 = divStep(x1, x0, y);
    <a id="L258"></a>}

    <a id="L260"></a>r = x0 &gt;&gt; z;

    <a id="L262"></a>if q1 != 0 {
        <a id="L263"></a>panic(&#34;div out of range&#34;)
    <a id="L264"></a>}

    <a id="L266"></a>return q0, r;
<a id="L267"></a>}


<a id="L270"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L271"></a><span class="comment">// Elementary operations on vectors</span>

<a id="L273"></a><span class="comment">// All higher-level functions use these elementary vector operations.</span>
<a id="L274"></a><span class="comment">// The function pointers f are initialized with default implementations</span>
<a id="L275"></a><span class="comment">// f_g, written in Go for portability. The corresponding assembly routines</span>
<a id="L276"></a><span class="comment">// f_s should be installed if they exist.</span>
<a id="L277"></a>var (
    <a id="L278"></a><span class="comment">// addVV sets z and returns c such that z+c = x+y.</span>
    <a id="L279"></a>addVV func(z, x, y *Word, n int) (c Word) = addVV_g;

    <a id="L281"></a><span class="comment">// subVV sets z and returns c such that z-c = x-y.</span>
    <a id="L282"></a>subVV func(z, x, y *Word, n int) (c Word) = subVV_g;

    <a id="L284"></a><span class="comment">// addVW sets z and returns c such that z+c = x-y.</span>
    <a id="L285"></a>addVW func(z, x *Word, y Word, n int) (c Word) = addVW_g;

    <a id="L287"></a><span class="comment">// subVW sets z and returns c such that z-c = x-y.</span>
    <a id="L288"></a>subVW func(z, x *Word, y Word, n int) (c Word) = subVW_g;

    <a id="L290"></a><span class="comment">// mulAddVWW sets z and returns c such that z+c = x*y + r.</span>
    <a id="L291"></a>mulAddVWW func(z, x *Word, y, r Word, n int) (c Word) = mulAddVWW_g;

    <a id="L293"></a><span class="comment">// addMulVVW sets z and returns c such that z+c = z + x*y.</span>
    <a id="L294"></a>addMulVVW func(z, x *Word, y Word, n int) (c Word) = addMulVVW_g;

    <a id="L296"></a><span class="comment">// divWVW sets z and returns r such that z-r = (xn&lt;&lt;(n*_W) + x) / y.</span>
    <a id="L297"></a>divWVW func(z *Word, xn Word, x *Word, y Word, n int) (r Word) = divWVW_g;
<a id="L298"></a>)


<a id="L301"></a><span class="comment">// UseAsm returns true if the assembly routines are enabled.</span>
<a id="L302"></a>func useAsm() bool

<a id="L304"></a>func init() {
    <a id="L305"></a>if useAsm() {
        <a id="L306"></a><span class="comment">// Install assembly routines.</span>
        <a id="L307"></a>addVV = addVV_s;
        <a id="L308"></a>subVV = subVV_s;
        <a id="L309"></a>addVW = addVW_s;
        <a id="L310"></a>subVW = subVW_s;
        <a id="L311"></a>mulAddVWW = mulAddVWW_s;
        <a id="L312"></a>addMulVVW = addMulVVW_s;
        <a id="L313"></a>divWVW = divWVW_s;
    <a id="L314"></a>}
<a id="L315"></a>}


<a id="L318"></a>func (p *Word) at(i int) *Word {
    <a id="L319"></a>return (*Word)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)*_S))
<a id="L320"></a>}


<a id="L323"></a>func addVV_s(z, x, y *Word, n int) (c Word)
<a id="L324"></a>func addVV_g(z, x, y *Word, n int) (c Word) {
    <a id="L325"></a>for i := 0; i &lt; n; i++ {
        <a id="L326"></a>c, *z.at(i) = addWW_g(*x.at(i), *y.at(i), c)
    <a id="L327"></a>}
    <a id="L328"></a>return;
<a id="L329"></a>}


<a id="L332"></a>func subVV_s(z, x, y *Word, n int) (c Word)
<a id="L333"></a>func subVV_g(z, x, y *Word, n int) (c Word) {
    <a id="L334"></a>for i := 0; i &lt; n; i++ {
        <a id="L335"></a>c, *z.at(i) = subWW_g(*x.at(i), *y.at(i), c)
    <a id="L336"></a>}
    <a id="L337"></a>return;
<a id="L338"></a>}


<a id="L341"></a>func addVW_s(z, x *Word, y Word, n int) (c Word)
<a id="L342"></a>func addVW_g(z, x *Word, y Word, n int) (c Word) {
    <a id="L343"></a>c = y;
    <a id="L344"></a>for i := 0; i &lt; n; i++ {
        <a id="L345"></a>c, *z.at(i) = addWW_g(*x.at(i), c, 0)
    <a id="L346"></a>}
    <a id="L347"></a>return;
<a id="L348"></a>}


<a id="L351"></a>func subVW_s(z, x *Word, y Word, n int) (c Word)
<a id="L352"></a>func subVW_g(z, x *Word, y Word, n int) (c Word) {
    <a id="L353"></a>c = y;
    <a id="L354"></a>for i := 0; i &lt; n; i++ {
        <a id="L355"></a>c, *z.at(i) = subWW_g(*x.at(i), c, 0)
    <a id="L356"></a>}
    <a id="L357"></a>return;
<a id="L358"></a>}


<a id="L361"></a>func mulAddVWW_s(z, x *Word, y, r Word, n int) (c Word)
<a id="L362"></a>func mulAddVWW_g(z, x *Word, y, r Word, n int) (c Word) {
    <a id="L363"></a>c = r;
    <a id="L364"></a>for i := 0; i &lt; n; i++ {
        <a id="L365"></a>c, *z.at(i) = mulAddWWW_g(*x.at(i), y, c)
    <a id="L366"></a>}
    <a id="L367"></a>return;
<a id="L368"></a>}


<a id="L371"></a>func addMulVVW_s(z, x *Word, y Word, n int) (c Word)
<a id="L372"></a>func addMulVVW_g(z, x *Word, y Word, n int) (c Word) {
    <a id="L373"></a>for i := 0; i &lt; n; i++ {
        <a id="L374"></a>z1, z0 := mulAddWWW_g(*x.at(i), y, *z.at(i));
        <a id="L375"></a>c, *z.at(i) = addWW_g(z0, c, 0);
        <a id="L376"></a>c += z1;
    <a id="L377"></a>}
    <a id="L378"></a>return;
<a id="L379"></a>}


<a id="L382"></a>func divWVW_s(z *Word, xn Word, x *Word, y Word, n int) (r Word)
<a id="L383"></a>func divWVW_g(z *Word, xn Word, x *Word, y Word, n int) (r Word) {
    <a id="L384"></a>r = xn;
    <a id="L385"></a>for i := n - 1; i &gt;= 0; i-- {
        <a id="L386"></a>*z.at(i), r = divWW_g(r, *x.at(i), y)
    <a id="L387"></a>}
    <a id="L388"></a>return;
<a id="L389"></a>}
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
