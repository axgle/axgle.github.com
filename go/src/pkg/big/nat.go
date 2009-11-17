<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/big/nat.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/big/nat.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This file contains operations on unsigned multi-precision integers.</span>
<a id="L6"></a><span class="comment">// These are the building blocks for the operations on signed integers</span>
<a id="L7"></a><span class="comment">// and rationals.</span>

<a id="L9"></a><span class="comment">// This package implements multi-precision arithmetic (big numbers).</span>
<a id="L10"></a><span class="comment">// The following numeric types are supported:</span>
<a id="L11"></a><span class="comment">//</span>
<a id="L12"></a><span class="comment">//	- Int	signed integers</span>
<a id="L13"></a><span class="comment">//</span>
<a id="L14"></a><span class="comment">// All methods on Int take the result as the receiver; if it is one</span>
<a id="L15"></a><span class="comment">// of the operands it may be overwritten (and its memory reused).</span>
<a id="L16"></a><span class="comment">// To enable chaining of operations, the result is also returned.</span>
<a id="L17"></a><span class="comment">//</span>
<a id="L18"></a><span class="comment">// If possible, one should use big over bignum as the latter is headed for</span>
<a id="L19"></a><span class="comment">// deprecation.</span>
<a id="L20"></a><span class="comment">//</span>
<a id="L21"></a>package big

<a id="L23"></a>import &#34;rand&#34;

<a id="L25"></a><span class="comment">// An unsigned integer x of the form</span>
<a id="L26"></a><span class="comment">//</span>
<a id="L27"></a><span class="comment">//   x = x[n-1]*_B^(n-1) + x[n-2]*_B^(n-2) + ... + x[1]*_B + x[0]</span>
<a id="L28"></a><span class="comment">//</span>
<a id="L29"></a><span class="comment">// with 0 &lt;= x[i] &lt; _B and 0 &lt;= i &lt; n is stored in a slice of length n,</span>
<a id="L30"></a><span class="comment">// with the digits x[i] as the slice elements.</span>
<a id="L31"></a><span class="comment">//</span>
<a id="L32"></a><span class="comment">// A number is normalized if the slice contains no leading 0 digits.</span>
<a id="L33"></a><span class="comment">// During arithmetic operations, denormalized values may occur but are</span>
<a id="L34"></a><span class="comment">// always normalized before returning the final result. The normalized</span>
<a id="L35"></a><span class="comment">// representation of 0 is the empty or nil slice (length = 0).</span>

<a id="L37"></a><span class="comment">// TODO(gri) - convert these routines into methods for type &#39;nat&#39;</span>
<a id="L38"></a><span class="comment">//           - decide if type &#39;nat&#39; should be exported</span>

<a id="L40"></a>func normN(z []Word) []Word {
    <a id="L41"></a>i := len(z);
    <a id="L42"></a>for i &gt; 0 &amp;&amp; z[i-1] == 0 {
        <a id="L43"></a>i--
    <a id="L44"></a>}
    <a id="L45"></a>z = z[0:i];
    <a id="L46"></a>return z;
<a id="L47"></a>}


<a id="L50"></a>func makeN(z []Word, m int, clear bool) []Word {
    <a id="L51"></a>if len(z) &gt; m {
        <a id="L52"></a>z = z[0:m]; <span class="comment">// reuse z - has at least one extra word for a carry, if any</span>
        <a id="L53"></a>if clear {
            <a id="L54"></a>for i := range z {
                <a id="L55"></a>z[i] = 0
            <a id="L56"></a>}
        <a id="L57"></a>}
        <a id="L58"></a>return z;
    <a id="L59"></a>}

    <a id="L61"></a>c := 4; <span class="comment">// minimum capacity</span>
    <a id="L62"></a>if m &gt; c {
        <a id="L63"></a>c = m
    <a id="L64"></a>}
    <a id="L65"></a>return make([]Word, m, c+1); <span class="comment">// +1: extra word for a carry, if any</span>
<a id="L66"></a>}


<a id="L69"></a>func newN(z []Word, x uint64) []Word {
    <a id="L70"></a>if x == 0 {
        <a id="L71"></a>return makeN(z, 0, false)
    <a id="L72"></a>}

    <a id="L74"></a><span class="comment">// single-digit values</span>
    <a id="L75"></a>if x == uint64(Word(x)) {
        <a id="L76"></a>z = makeN(z, 1, false);
        <a id="L77"></a>z[0] = Word(x);
        <a id="L78"></a>return z;
    <a id="L79"></a>}

    <a id="L81"></a><span class="comment">// compute number of words n required to represent x</span>
    <a id="L82"></a>n := 0;
    <a id="L83"></a>for t := x; t &gt; 0; t &gt;&gt;= _W {
        <a id="L84"></a>n++
    <a id="L85"></a>}

    <a id="L87"></a><span class="comment">// split x into n words</span>
    <a id="L88"></a>z = makeN(z, n, false);
    <a id="L89"></a>for i := 0; i &lt; n; i++ {
        <a id="L90"></a>z[i] = Word(x &amp; _M);
        <a id="L91"></a>x &gt;&gt;= _W;
    <a id="L92"></a>}

    <a id="L94"></a>return z;
<a id="L95"></a>}


<a id="L98"></a>func setN(z, x []Word) []Word {
    <a id="L99"></a>z = makeN(z, len(x), false);
    <a id="L100"></a>for i, d := range x {
        <a id="L101"></a>z[i] = d
    <a id="L102"></a>}
    <a id="L103"></a>return z;
<a id="L104"></a>}


<a id="L107"></a>func addNN(z, x, y []Word) []Word {
    <a id="L108"></a>m := len(x);
    <a id="L109"></a>n := len(y);

    <a id="L111"></a>switch {
    <a id="L112"></a>case m &lt; n:
        <a id="L113"></a>return addNN(z, y, x)
    <a id="L114"></a>case m == 0:
        <a id="L115"></a><span class="comment">// n == 0 because m &gt;= n; result is 0</span>
        <a id="L116"></a>return makeN(z, 0, false)
    <a id="L117"></a>case n == 0:
        <a id="L118"></a><span class="comment">// result is x</span>
        <a id="L119"></a>return setN(z, x)
    <a id="L120"></a>}
    <a id="L121"></a><span class="comment">// m &gt; 0</span>

    <a id="L123"></a>z = makeN(z, m, false);
    <a id="L124"></a>c := addVV(&amp;z[0], &amp;x[0], &amp;y[0], n);
    <a id="L125"></a>if m &gt; n {
        <a id="L126"></a>c = addVW(&amp;z[n], &amp;x[n], c, m-n)
    <a id="L127"></a>}
    <a id="L128"></a>if c &gt; 0 {
        <a id="L129"></a>z = z[0 : m+1];
        <a id="L130"></a>z[m] = c;
    <a id="L131"></a>}

    <a id="L133"></a>return z;
<a id="L134"></a>}


<a id="L137"></a>func subNN(z, x, y []Word) []Word {
    <a id="L138"></a>m := len(x);
    <a id="L139"></a>n := len(y);

    <a id="L141"></a>switch {
    <a id="L142"></a>case m &lt; n:
        <a id="L143"></a>panic(&#34;underflow&#34;)
    <a id="L144"></a>case m == 0:
        <a id="L145"></a><span class="comment">// n == 0 because m &gt;= n; result is 0</span>
        <a id="L146"></a>return makeN(z, 0, false)
    <a id="L147"></a>case n == 0:
        <a id="L148"></a><span class="comment">// result is x</span>
        <a id="L149"></a>return setN(z, x)
    <a id="L150"></a>}
    <a id="L151"></a><span class="comment">// m &gt; 0</span>

    <a id="L153"></a>z = makeN(z, m, false);
    <a id="L154"></a>c := subVV(&amp;z[0], &amp;x[0], &amp;y[0], n);
    <a id="L155"></a>if m &gt; n {
        <a id="L156"></a>c = subVW(&amp;z[n], &amp;x[n], c, m-n)
    <a id="L157"></a>}
    <a id="L158"></a>if c != 0 {
        <a id="L159"></a>panic(&#34;underflow&#34;)
    <a id="L160"></a>}
    <a id="L161"></a>z = normN(z);

    <a id="L163"></a>return z;
<a id="L164"></a>}


<a id="L167"></a>func cmpNN(x, y []Word) (r int) {
    <a id="L168"></a>m := len(x);
    <a id="L169"></a>n := len(y);
    <a id="L170"></a>if m != n || m == 0 {
        <a id="L171"></a>switch {
        <a id="L172"></a>case m &lt; n:
            <a id="L173"></a>r = -1
        <a id="L174"></a>case m &gt; n:
            <a id="L175"></a>r = 1
        <a id="L176"></a>}
        <a id="L177"></a>return;
    <a id="L178"></a>}

    <a id="L180"></a>i := m - 1;
    <a id="L181"></a>for i &gt; 0 &amp;&amp; x[i] == y[i] {
        <a id="L182"></a>i--
    <a id="L183"></a>}

    <a id="L185"></a>switch {
    <a id="L186"></a>case x[i] &lt; y[i]:
        <a id="L187"></a>r = -1
    <a id="L188"></a>case x[i] &gt; y[i]:
        <a id="L189"></a>r = 1
    <a id="L190"></a>}
    <a id="L191"></a>return;
<a id="L192"></a>}


<a id="L195"></a>func mulAddNWW(z, x []Word, y, r Word) []Word {
    <a id="L196"></a>m := len(x);
    <a id="L197"></a>if m == 0 || y == 0 {
        <a id="L198"></a>return newN(z, uint64(r)) <span class="comment">// result is r</span>
    <a id="L199"></a>}
    <a id="L200"></a><span class="comment">// m &gt; 0</span>

    <a id="L202"></a>z = makeN(z, m, false);
    <a id="L203"></a>c := mulAddVWW(&amp;z[0], &amp;x[0], y, r, m);
    <a id="L204"></a>if c &gt; 0 {
        <a id="L205"></a>z = z[0 : m+1];
        <a id="L206"></a>z[m] = c;
    <a id="L207"></a>}

    <a id="L209"></a>return z;
<a id="L210"></a>}


<a id="L213"></a>func mulNN(z, x, y []Word) []Word {
    <a id="L214"></a>m := len(x);
    <a id="L215"></a>n := len(y);

    <a id="L217"></a>switch {
    <a id="L218"></a>case m &lt; n:
        <a id="L219"></a>return mulNN(z, y, x)
    <a id="L220"></a>case m == 0 || n == 0:
        <a id="L221"></a>return makeN(z, 0, false)
    <a id="L222"></a>case n == 1:
        <a id="L223"></a>return mulAddNWW(z, x, y[0], 0)
    <a id="L224"></a>}
    <a id="L225"></a><span class="comment">// m &gt;= n &amp;&amp; m &gt; 1 &amp;&amp; n &gt; 1</span>

    <a id="L227"></a>z = makeN(z, m+n, true);
    <a id="L228"></a>if &amp;z[0] == &amp;x[0] || &amp;z[0] == &amp;y[0] {
        <a id="L229"></a>z = makeN(nil, m+n, true) <span class="comment">// z is an alias for x or y - cannot reuse</span>
    <a id="L230"></a>}
    <a id="L231"></a>for i := 0; i &lt; n; i++ {
        <a id="L232"></a>if f := y[i]; f != 0 {
            <a id="L233"></a>z[m+i] = addMulVVW(&amp;z[i], &amp;x[0], f, m)
        <a id="L234"></a>}
    <a id="L235"></a>}
    <a id="L236"></a>z = normN(z);

    <a id="L238"></a>return z;
<a id="L239"></a>}


<a id="L242"></a><span class="comment">// q = (x-r)/y, with 0 &lt;= r &lt; y</span>
<a id="L243"></a>func divNW(z, x []Word, y Word) (q []Word, r Word) {
    <a id="L244"></a>m := len(x);
    <a id="L245"></a>switch {
    <a id="L246"></a>case y == 0:
        <a id="L247"></a>panic(&#34;division by zero&#34;)
    <a id="L248"></a>case y == 1:
        <a id="L249"></a>q = setN(z, x); <span class="comment">// result is x</span>
        <a id="L250"></a>return;
    <a id="L251"></a>case m == 0:
        <a id="L252"></a>q = setN(z, nil); <span class="comment">// result is 0</span>
        <a id="L253"></a>return;
    <a id="L254"></a>}
    <a id="L255"></a><span class="comment">// m &gt; 0</span>
    <a id="L256"></a>z = makeN(z, m, false);
    <a id="L257"></a>r = divWVW(&amp;z[0], 0, &amp;x[0], y, m);
    <a id="L258"></a>q = normN(z);
    <a id="L259"></a>return;
<a id="L260"></a>}


<a id="L263"></a>func divNN(z, z2, u, v []Word) (q, r []Word) {
    <a id="L264"></a>if len(v) == 0 {
        <a id="L265"></a>panic(&#34;Divide by zero undefined&#34;)
    <a id="L266"></a>}

    <a id="L268"></a>if cmpNN(u, v) &lt; 0 {
        <a id="L269"></a>q = makeN(z, 0, false);
        <a id="L270"></a>r = setN(z2, u);
        <a id="L271"></a>return;
    <a id="L272"></a>}

    <a id="L274"></a>if len(v) == 1 {
        <a id="L275"></a>var rprime Word;
        <a id="L276"></a>q, rprime = divNW(z, u, v[0]);
        <a id="L277"></a>if rprime &gt; 0 {
            <a id="L278"></a>r = makeN(z2, 1, false);
            <a id="L279"></a>r[0] = rprime;
        <a id="L280"></a>} else {
            <a id="L281"></a>r = makeN(z2, 0, false)
        <a id="L282"></a>}
        <a id="L283"></a>return;
    <a id="L284"></a>}

    <a id="L286"></a>q, r = divLargeNN(z, z2, u, v);
    <a id="L287"></a>return;
<a id="L288"></a>}


<a id="L291"></a><span class="comment">// q = (uIn-r)/v, with 0 &lt;= r &lt; y</span>
<a id="L292"></a><span class="comment">// See Knuth, Volume 2, section 4.3.1, Algorithm D.</span>
<a id="L293"></a><span class="comment">// Preconditions:</span>
<a id="L294"></a><span class="comment">//    len(v) &gt;= 2</span>
<a id="L295"></a><span class="comment">//    len(uIn) &gt;= len(v)</span>
<a id="L296"></a>func divLargeNN(z, z2, uIn, v []Word) (q, r []Word) {
    <a id="L297"></a>n := len(v);
    <a id="L298"></a>m := len(uIn) - len(v);

    <a id="L300"></a>u := makeN(z2, len(uIn)+1, false);
    <a id="L301"></a>qhatv := make([]Word, len(v)+1);
    <a id="L302"></a>q = makeN(z, m+1, false);

    <a id="L304"></a><span class="comment">// D1.</span>
    <a id="L305"></a>shift := leadingZeroBits(v[n-1]);
    <a id="L306"></a>shiftLeft(v, v, shift);
    <a id="L307"></a>shiftLeft(u, uIn, shift);
    <a id="L308"></a>u[len(uIn)] = uIn[len(uIn)-1] &gt;&gt; (_W - uint(shift));

    <a id="L310"></a><span class="comment">// D2.</span>
    <a id="L311"></a>for j := m; j &gt;= 0; j-- {
        <a id="L312"></a><span class="comment">// D3.</span>
        <a id="L313"></a>qhat, rhat := divWW_g(u[j+n], u[j+n-1], v[n-1]);

        <a id="L315"></a><span class="comment">// x1 | x2 = q̂v_{n-2}</span>
        <a id="L316"></a>x1, x2 := mulWW_g(qhat, v[n-2]);
        <a id="L317"></a><span class="comment">// test if q̂v_{n-2} &gt; br̂ + u_{j+n-2}</span>
        <a id="L318"></a>for greaterThan(x1, x2, rhat, u[j+n-2]) {
            <a id="L319"></a>qhat--;
            <a id="L320"></a>prevRhat := rhat;
            <a id="L321"></a>rhat += v[n-1];
            <a id="L322"></a><span class="comment">// v[n-1] &gt;= 0, so this tests for overflow.</span>
            <a id="L323"></a>if rhat &lt; prevRhat {
                <a id="L324"></a>break
            <a id="L325"></a>}
            <a id="L326"></a>x1, x2 = mulWW_g(qhat, v[n-2]);
        <a id="L327"></a>}

        <a id="L329"></a><span class="comment">// D4.</span>
        <a id="L330"></a>qhatv[len(v)] = mulAddVWW(&amp;qhatv[0], &amp;v[0], qhat, 0, len(v));

        <a id="L332"></a>c := subVV(&amp;u[j], &amp;u[j], &amp;qhatv[0], len(qhatv));
        <a id="L333"></a>if c != 0 {
            <a id="L334"></a>c := addVV(&amp;u[j], &amp;u[j], &amp;v[0], len(v));
            <a id="L335"></a>u[j+len(v)] += c;
            <a id="L336"></a>qhat--;
        <a id="L337"></a>}

        <a id="L339"></a>q[j] = qhat;
    <a id="L340"></a>}

    <a id="L342"></a>q = normN(q);
    <a id="L343"></a>shiftRight(u, u, shift);
    <a id="L344"></a>shiftRight(v, v, shift);
    <a id="L345"></a>r = normN(u);

    <a id="L347"></a>return q, r;
<a id="L348"></a>}


<a id="L351"></a><span class="comment">// log2 computes the integer binary logarithm of x.</span>
<a id="L352"></a><span class="comment">// The result is the integer n for which 2^n &lt;= x &lt; 2^(n+1).</span>
<a id="L353"></a><span class="comment">// If x == 0, the result is -1.</span>
<a id="L354"></a>func log2(x Word) int {
    <a id="L355"></a>n := 0;
    <a id="L356"></a>for ; x &gt; 0; x &gt;&gt;= 1 {
        <a id="L357"></a>n++
    <a id="L358"></a>}
    <a id="L359"></a>return n - 1;
<a id="L360"></a>}


<a id="L363"></a><span class="comment">// log2N computes the integer binary logarithm of x.</span>
<a id="L364"></a><span class="comment">// The result is the integer n for which 2^n &lt;= x &lt; 2^(n+1).</span>
<a id="L365"></a><span class="comment">// If x == 0, the result is -1.</span>
<a id="L366"></a>func log2N(x []Word) int {
    <a id="L367"></a>m := len(x);
    <a id="L368"></a>if m &gt; 0 {
        <a id="L369"></a>return (m-1)*_W + log2(x[m-1])
    <a id="L370"></a>}
    <a id="L371"></a>return -1;
<a id="L372"></a>}


<a id="L375"></a>func hexValue(ch byte) int {
    <a id="L376"></a>var d byte;
    <a id="L377"></a>switch {
    <a id="L378"></a>case &#39;0&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;9&#39;:
        <a id="L379"></a>d = ch - &#39;0&#39;
    <a id="L380"></a>case &#39;a&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;f&#39;:
        <a id="L381"></a>d = ch - &#39;a&#39; + 10
    <a id="L382"></a>case &#39;A&#39; &lt;= ch &amp;&amp; ch &lt;= &#39;F&#39;:
        <a id="L383"></a>d = ch - &#39;A&#39; + 10
    <a id="L384"></a>default:
        <a id="L385"></a>return -1
    <a id="L386"></a>}
    <a id="L387"></a>return int(d);
<a id="L388"></a>}


<a id="L391"></a><span class="comment">// scanN returns the natural number corresponding to the</span>
<a id="L392"></a><span class="comment">// longest possible prefix of s representing a natural number in a</span>
<a id="L393"></a><span class="comment">// given conversion base, the actual conversion base used, and the</span>
<a id="L394"></a><span class="comment">// prefix length. The syntax of natural numbers follows the syntax</span>
<a id="L395"></a><span class="comment">// of unsigned integer literals in Go.</span>
<a id="L396"></a><span class="comment">//</span>
<a id="L397"></a><span class="comment">// If the base argument is 0, the string prefix determines the actual</span>
<a id="L398"></a><span class="comment">// conversion base. A prefix of ``0x&#39;&#39; or ``0X&#39;&#39; selects base 16; the</span>
<a id="L399"></a><span class="comment">// ``0&#39;&#39; prefix selects base 8. Otherwise the selected base is 10.</span>
<a id="L400"></a><span class="comment">//</span>
<a id="L401"></a>func scanN(z []Word, s string, base int) ([]Word, int, int) {
    <a id="L402"></a><span class="comment">// determine base if necessary</span>
    <a id="L403"></a>i, n := 0, len(s);
    <a id="L404"></a>if base == 0 {
        <a id="L405"></a>base = 10;
        <a id="L406"></a>if n &gt; 0 &amp;&amp; s[0] == &#39;0&#39; {
            <a id="L407"></a>if n &gt; 1 &amp;&amp; (s[1] == &#39;x&#39; || s[1] == &#39;X&#39;) {
                <a id="L408"></a>if n == 2 {
                    <a id="L409"></a><span class="comment">// Reject a string which is just &#39;0x&#39; as nonsense.</span>
                    <a id="L410"></a>return nil, 0, 0
                <a id="L411"></a>}
                <a id="L412"></a>base, i = 16, 2;
            <a id="L413"></a>} else {
                <a id="L414"></a>base, i = 8, 1
            <a id="L415"></a>}
        <a id="L416"></a>}
    <a id="L417"></a>}
    <a id="L418"></a>if base &lt; 2 || 16 &lt; base {
        <a id="L419"></a>panic(&#34;illegal base&#34;)
    <a id="L420"></a>}

    <a id="L422"></a><span class="comment">// convert string</span>
    <a id="L423"></a>z = makeN(z, len(z), false);
    <a id="L424"></a>for ; i &lt; n; i++ {
        <a id="L425"></a>d := hexValue(s[i]);
        <a id="L426"></a>if 0 &lt;= d &amp;&amp; d &lt; base {
            <a id="L427"></a>z = mulAddNWW(z, z, Word(base), Word(d))
        <a id="L428"></a>} else {
            <a id="L429"></a>break
        <a id="L430"></a>}
    <a id="L431"></a>}

    <a id="L433"></a>return z, base, i;
<a id="L434"></a>}


<a id="L437"></a><span class="comment">// string converts x to a string for a given base, with 2 &lt;= base &lt;= 16.</span>
<a id="L438"></a><span class="comment">// TODO(gri) in the style of the other routines, perhaps this should take</span>
<a id="L439"></a><span class="comment">//           a []byte buffer and return it</span>
<a id="L440"></a>func stringN(x []Word, base int) string {
    <a id="L441"></a>if base &lt; 2 || 16 &lt; base {
        <a id="L442"></a>panic(&#34;illegal base&#34;)
    <a id="L443"></a>}

    <a id="L445"></a>if len(x) == 0 {
        <a id="L446"></a>return &#34;0&#34;
    <a id="L447"></a>}

    <a id="L449"></a><span class="comment">// allocate buffer for conversion</span>
    <a id="L450"></a>i := (log2N(x)+1)/log2(Word(base)) + 1; <span class="comment">// +1: round up</span>
    <a id="L451"></a>s := make([]byte, i);

    <a id="L453"></a><span class="comment">// don&#39;t destroy x</span>
    <a id="L454"></a>q := setN(nil, x);

    <a id="L456"></a><span class="comment">// convert</span>
    <a id="L457"></a>for len(q) &gt; 0 {
        <a id="L458"></a>i--;
        <a id="L459"></a>var r Word;
        <a id="L460"></a>q, r = divNW(q, q, Word(base));
        <a id="L461"></a>s[i] = &#34;0123456789abcdef&#34;[r];
    <a id="L462"></a>}

    <a id="L464"></a>return string(s[i:len(s)]);
<a id="L465"></a>}


<a id="L468"></a><span class="comment">// leadingZeroBits returns the number of leading zero bits in x.</span>
<a id="L469"></a>func leadingZeroBits(x Word) int {
    <a id="L470"></a>c := 0;
    <a id="L471"></a>if x &lt; 1&lt;&lt;(_W/2) {
        <a id="L472"></a>x &lt;&lt;= _W / 2;
        <a id="L473"></a>c = _W / 2;
    <a id="L474"></a>}

    <a id="L476"></a>for i := 0; x != 0; i++ {
        <a id="L477"></a>if x&amp;(1&lt;&lt;(_W-1)) != 0 {
            <a id="L478"></a>return i + c
        <a id="L479"></a>}
        <a id="L480"></a>x &lt;&lt;= 1;
    <a id="L481"></a>}

    <a id="L483"></a>return _W;
<a id="L484"></a>}

<a id="L486"></a>const deBruijn32 = 0x077CB531

<a id="L488"></a>var deBruijn32Lookup = []byte{
    <a id="L489"></a>0, 1, 28, 2, 29, 14, 24, 3, 30, 22, 20, 15, 25, 17, 4, 8,
    <a id="L490"></a>31, 27, 13, 23, 21, 19, 16, 7, 26, 12, 18, 6, 11, 5, 10, 9,
<a id="L491"></a>}

<a id="L493"></a>const deBruijn64 = 0x03f79d71b4ca8b09

<a id="L495"></a>var deBruijn64Lookup = []byte{
    <a id="L496"></a>0, 1, 56, 2, 57, 49, 28, 3, 61, 58, 42, 50, 38, 29, 17, 4,
    <a id="L497"></a>62, 47, 59, 36, 45, 43, 51, 22, 53, 39, 33, 30, 24, 18, 12, 5,
    <a id="L498"></a>63, 55, 48, 27, 60, 41, 37, 16, 46, 35, 44, 21, 52, 32, 23, 11,
    <a id="L499"></a>54, 26, 40, 15, 34, 20, 31, 10, 25, 14, 19, 9, 13, 8, 7, 6,
<a id="L500"></a>}

<a id="L502"></a><span class="comment">// trailingZeroBits returns the number of consecutive zero bits on the right</span>
<a id="L503"></a><span class="comment">// side of the given Word.</span>
<a id="L504"></a><span class="comment">// See Knuth, volume 4, section 7.3.1</span>
<a id="L505"></a>func trailingZeroBits(x Word) int {
    <a id="L506"></a><span class="comment">// x &amp; -x leaves only the right-most bit set in the word. Let k be the</span>
    <a id="L507"></a><span class="comment">// index of that bit. Since only a single bit is set, the value is two</span>
    <a id="L508"></a><span class="comment">// to the power of k. Multipling by a power of two is equivalent to</span>
    <a id="L509"></a><span class="comment">// left shifting, in this case by k bits.  The de Bruijn constant is</span>
    <a id="L510"></a><span class="comment">// such that all six bit, consecutive substrings are distinct.</span>
    <a id="L511"></a><span class="comment">// Therefore, if we have a left shifted version of this constant we can</span>
    <a id="L512"></a><span class="comment">// find by how many bits it was shifted by looking at which six bit</span>
    <a id="L513"></a><span class="comment">// substring ended up at the top of the word.</span>
    <a id="L514"></a>switch _W {
    <a id="L515"></a>case 32:
        <a id="L516"></a>return int(deBruijn32Lookup[((x&amp;-x)*deBruijn32)&gt;&gt;27])
    <a id="L517"></a>case 64:
        <a id="L518"></a>return int(deBruijn64Lookup[((x&amp;-x)*(deBruijn64&amp;_M))&gt;&gt;58])
    <a id="L519"></a>default:
        <a id="L520"></a>panic(&#34;Unknown word size&#34;)
    <a id="L521"></a>}

    <a id="L523"></a>return 0;
<a id="L524"></a>}


<a id="L527"></a>func shiftLeft(dst, src []Word, n int) {
    <a id="L528"></a>if len(src) == 0 {
        <a id="L529"></a>return
    <a id="L530"></a>}

    <a id="L532"></a>ñ := _W - uint(n);
    <a id="L533"></a>for i := len(src) - 1; i &gt;= 1; i-- {
        <a id="L534"></a>dst[i] = src[i] &lt;&lt; uint(n);
        <a id="L535"></a>dst[i] |= src[i-1] &gt;&gt; ñ;
    <a id="L536"></a>}
    <a id="L537"></a>dst[0] = src[0] &lt;&lt; uint(n);
<a id="L538"></a>}


<a id="L541"></a>func shiftRight(dst, src []Word, n int) {
    <a id="L542"></a>if len(src) == 0 {
        <a id="L543"></a>return
    <a id="L544"></a>}

    <a id="L546"></a>ñ := _W - uint(n);
    <a id="L547"></a>for i := 0; i &lt; len(src)-1; i++ {
        <a id="L548"></a>dst[i] = src[i] &gt;&gt; uint(n);
        <a id="L549"></a>dst[i] |= src[i+1] &lt;&lt; ñ;
    <a id="L550"></a>}
    <a id="L551"></a>dst[len(src)-1] = src[len(src)-1] &gt;&gt; uint(n);
<a id="L552"></a>}


<a id="L555"></a><span class="comment">// greaterThan returns true iff (x1&lt;&lt;_W + x2) &gt; (y1&lt;&lt;_W + y2)</span>
<a id="L556"></a>func greaterThan(x1, x2, y1, y2 Word) bool { return x1 &gt; y1 || x1 == y1 &amp;&amp; x2 &gt; y2 }


<a id="L559"></a><span class="comment">// modNW returns x % d.</span>
<a id="L560"></a>func modNW(x []Word, d Word) (r Word) {
    <a id="L561"></a><span class="comment">// TODO(agl): we don&#39;t actually need to store the q value.</span>
    <a id="L562"></a>q := makeN(nil, len(x), false);
    <a id="L563"></a>return divWVW(&amp;q[0], 0, &amp;x[0], d, len(x));
<a id="L564"></a>}


<a id="L567"></a><span class="comment">// powersOfTwoDecompose finds q and k such that q * 1&lt;&lt;k = n and q is odd.</span>
<a id="L568"></a>func powersOfTwoDecompose(n []Word) (q []Word, k Word) {
    <a id="L569"></a>if len(n) == 0 {
        <a id="L570"></a>return n, 0
    <a id="L571"></a>}

    <a id="L573"></a>zeroWords := 0;
    <a id="L574"></a>for n[zeroWords] == 0 {
        <a id="L575"></a>zeroWords++
    <a id="L576"></a>}
    <a id="L577"></a><span class="comment">// One of the words must be non-zero by invariant, therefore</span>
    <a id="L578"></a><span class="comment">// zeroWords &lt; len(n).</span>
    <a id="L579"></a>x := trailingZeroBits(n[zeroWords]);

    <a id="L581"></a>q = makeN(nil, len(n)-zeroWords, false);
    <a id="L582"></a>shiftRight(q, n[zeroWords:len(n)], x);

    <a id="L584"></a>k = Word(_W*zeroWords + x);
    <a id="L585"></a>return;
<a id="L586"></a>}


<a id="L589"></a><span class="comment">// randomN creates a random integer in [0..limit), using the space in z if</span>
<a id="L590"></a><span class="comment">// possible. n is the bit length of limit.</span>
<a id="L591"></a>func randomN(z []Word, rand *rand.Rand, limit []Word, n int) []Word {
    <a id="L592"></a>bitLengthOfMSW := uint(n % _W);
    <a id="L593"></a>mask := Word((1 &lt;&lt; bitLengthOfMSW) - 1);
    <a id="L594"></a>z = makeN(z, len(limit), false);

    <a id="L596"></a>for {
        <a id="L597"></a>for i := range z {
            <a id="L598"></a>switch _W {
            <a id="L599"></a>case 32:
                <a id="L600"></a>z[i] = Word(rand.Uint32())
            <a id="L601"></a>case 64:
                <a id="L602"></a>z[i] = Word(rand.Uint32()) | Word(rand.Uint32())&lt;&lt;32
            <a id="L603"></a>}
        <a id="L604"></a>}

        <a id="L606"></a>z[len(limit)-1] &amp;= mask;

        <a id="L608"></a>if cmpNN(z, limit) &lt; 0 {
            <a id="L609"></a>break
        <a id="L610"></a>}
    <a id="L611"></a>}

    <a id="L613"></a>return z;
<a id="L614"></a>}


<a id="L617"></a><span class="comment">// If m != nil, expNNN calculates x**y mod m. Otherwise it calculates x**y. It</span>
<a id="L618"></a><span class="comment">// reuses the storage of z if possible.</span>
<a id="L619"></a>func expNNN(z, x, y, m []Word) []Word {
    <a id="L620"></a>if len(y) == 0 {
        <a id="L621"></a>z = makeN(z, 1, false);
        <a id="L622"></a>z[0] = 1;
        <a id="L623"></a>return z;
    <a id="L624"></a>}

    <a id="L626"></a>if m != nil {
        <a id="L627"></a><span class="comment">// We likely end up being as long as the modulus.</span>
        <a id="L628"></a>z = makeN(z, len(m), false)
    <a id="L629"></a>}
    <a id="L630"></a>z = setN(z, x);
    <a id="L631"></a>v := y[len(y)-1];
    <a id="L632"></a><span class="comment">// It&#39;s invalid for the most significant word to be zero, therefore we</span>
    <a id="L633"></a><span class="comment">// will find a one bit.</span>
    <a id="L634"></a>shift := leadingZeros(v) + 1;
    <a id="L635"></a>v &lt;&lt;= shift;
    <a id="L636"></a>var q []Word;

    <a id="L638"></a>const mask = 1 &lt;&lt; (_W - 1);

    <a id="L640"></a><span class="comment">// We walk through the bits of the exponent one by one. Each time we</span>
    <a id="L641"></a><span class="comment">// see a bit, we square, thus doubling the power. If the bit is a one,</span>
    <a id="L642"></a><span class="comment">// we also multiply by x, thus adding one to the power.</span>

    <a id="L644"></a>w := _W - int(shift);
    <a id="L645"></a>for j := 0; j &lt; w; j++ {
        <a id="L646"></a>z = mulNN(z, z, z);

        <a id="L648"></a>if v&amp;mask != 0 {
            <a id="L649"></a>z = mulNN(z, z, x)
        <a id="L650"></a>}

        <a id="L652"></a>if m != nil {
            <a id="L653"></a>q, z = divNN(q, z, z, m)
        <a id="L654"></a>}

        <a id="L656"></a>v &lt;&lt;= 1;
    <a id="L657"></a>}

    <a id="L659"></a>for i := len(y) - 2; i &gt;= 0; i-- {
        <a id="L660"></a>v = y[i];

        <a id="L662"></a>for j := 0; j &lt; _W; j++ {
            <a id="L663"></a>z = mulNN(z, z, z);

            <a id="L665"></a>if v&amp;mask != 0 {
                <a id="L666"></a>z = mulNN(z, z, x)
            <a id="L667"></a>}

            <a id="L669"></a>if m != nil {
                <a id="L670"></a>q, z = divNN(q, z, z, m)
            <a id="L671"></a>}

            <a id="L673"></a>v &lt;&lt;= 1;
        <a id="L674"></a>}
    <a id="L675"></a>}

    <a id="L677"></a>return z;
<a id="L678"></a>}


<a id="L681"></a><span class="comment">// lenN returns the bit length of z.</span>
<a id="L682"></a>func lenN(z []Word) int {
    <a id="L683"></a>if len(z) == 0 {
        <a id="L684"></a>return 0
    <a id="L685"></a>}

    <a id="L687"></a>return (len(z)-1)*_W + (_W - leadingZeroBits(z[len(z)-1]));
<a id="L688"></a>}


<a id="L691"></a>const (
    <a id="L692"></a>primesProduct32 = 0xC0CFD797;         <span class="comment">// Π {p ∈ primes, 2 &lt; p &lt;= 29}</span>
    <a id="L693"></a>primesProduct64 = 0xE221F97C30E94E1D; <span class="comment">// Π {p ∈ primes, 2 &lt; p &lt;= 53}</span>
<a id="L694"></a>)

<a id="L696"></a>var bigOne = []Word{1}
<a id="L697"></a>var bigTwo = []Word{2}

<a id="L699"></a><span class="comment">// ProbablyPrime performs n Miller-Rabin tests to check whether n is prime.</span>
<a id="L700"></a><span class="comment">// If it returns true, n is prime with probability 1 - 1/4^n.</span>
<a id="L701"></a><span class="comment">// If it returns false, n is not prime.</span>
<a id="L702"></a>func probablyPrime(n []Word, reps int) bool {
    <a id="L703"></a>if len(n) == 0 {
        <a id="L704"></a>return false
    <a id="L705"></a>}

    <a id="L707"></a>if len(n) == 1 {
        <a id="L708"></a>if n[0]%2 == 0 {
            <a id="L709"></a>return n[0] == 2
        <a id="L710"></a>}

        <a id="L712"></a><span class="comment">// We have to exclude these cases because we reject all</span>
        <a id="L713"></a><span class="comment">// multiples of these numbers below.</span>
        <a id="L714"></a>if n[0] == 3 || n[0] == 5 || n[0] == 7 || n[0] == 11 ||
            <a id="L715"></a>n[0] == 13 || n[0] == 17 || n[0] == 19 || n[0] == 23 ||
            <a id="L716"></a>n[0] == 29 || n[0] == 31 || n[0] == 37 || n[0] == 41 ||
            <a id="L717"></a>n[0] == 43 || n[0] == 47 || n[0] == 53 {
            <a id="L718"></a>return true
        <a id="L719"></a>}
    <a id="L720"></a>}

    <a id="L722"></a>var r Word;
    <a id="L723"></a>switch _W {
    <a id="L724"></a>case 32:
        <a id="L725"></a>r = modNW(n, primesProduct32)
    <a id="L726"></a>case 64:
        <a id="L727"></a>r = modNW(n, primesProduct64&amp;_M)
    <a id="L728"></a>default:
        <a id="L729"></a>panic(&#34;Unknown word size&#34;)
    <a id="L730"></a>}

    <a id="L732"></a>if r%3 == 0 || r%5 == 0 || r%7 == 0 || r%11 == 0 ||
        <a id="L733"></a>r%13 == 0 || r%17 == 0 || r%19 == 0 || r%23 == 0 || r%29 == 0 {
        <a id="L734"></a>return false
    <a id="L735"></a>}

    <a id="L737"></a>if _W == 64 &amp;&amp; (r%31 == 0 || r%37 == 0 || r%41 == 0 ||
        <a id="L738"></a>r%43 == 0 || r%47 == 0 || r%53 == 0) {
        <a id="L739"></a>return false
    <a id="L740"></a>}

    <a id="L742"></a>nm1 := subNN(nil, n, bigOne);
    <a id="L743"></a><span class="comment">// 1&lt;&lt;k * q = nm1;</span>
    <a id="L744"></a>q, k := powersOfTwoDecompose(nm1);

    <a id="L746"></a>nm3 := subNN(nil, nm1, bigTwo);
    <a id="L747"></a>rand := rand.New(rand.NewSource(int64(n[0])));

    <a id="L749"></a>var x, y, quotient []Word;
    <a id="L750"></a>nm3Len := lenN(nm3);

<a id="L752"></a>NextRandom:
    <a id="L753"></a>for i := 0; i &lt; reps; i++ {
        <a id="L754"></a>x = randomN(x, rand, nm3, nm3Len);
        <a id="L755"></a>addNN(x, x, bigTwo);
        <a id="L756"></a>y = expNNN(y, x, q, n);
        <a id="L757"></a>if cmpNN(y, bigOne) == 0 || cmpNN(y, nm1) == 0 {
            <a id="L758"></a>continue
        <a id="L759"></a>}
        <a id="L760"></a>for j := Word(1); j &lt; k; j++ {
            <a id="L761"></a>y = mulNN(y, y, y);
            <a id="L762"></a>quotient, y = divNN(quotient, y, y, n);
            <a id="L763"></a>if cmpNN(y, nm1) == 0 {
                <a id="L764"></a>continue NextRandom
            <a id="L765"></a>}
            <a id="L766"></a>if cmpNN(y, bigOne) == 0 {
                <a id="L767"></a>return false
            <a id="L768"></a>}
        <a id="L769"></a>}
        <a id="L770"></a>return false;
    <a id="L771"></a>}

    <a id="L773"></a>return true;
<a id="L774"></a>}
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
