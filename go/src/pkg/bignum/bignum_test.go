<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/bignum/bignum_test.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/bignum/bignum_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package bignum

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>const (
    <a id="L13"></a>sa  = &#34;991&#34;;
    <a id="L14"></a>sb  = &#34;2432902008176640000&#34;; <span class="comment">// 20!</span>
    <a id="L15"></a>sc  = &#34;933262154439441526816992388562667004907159682643816214685929&#34;
        <a id="L16"></a>&#34;638952175999932299156089414639761565182862536979208272237582&#34;
        <a id="L17"></a>&#34;51185210916864000000000000000000000000&#34;; <span class="comment">// 100!</span>
    <a id="L18"></a>sp  = &#34;170141183460469231731687303715884105727&#34;; <span class="comment">// prime</span>
<a id="L19"></a>)

<a id="L21"></a>func natFromString(s string, base uint, slen *int) Natural {
    <a id="L22"></a>x, _, len := NatFromString(s, base);
    <a id="L23"></a>if slen != nil {
        <a id="L24"></a>*slen = len
    <a id="L25"></a>}
    <a id="L26"></a>return x;
<a id="L27"></a>}


<a id="L30"></a>func intFromString(s string, base uint, slen *int) *Integer {
    <a id="L31"></a>x, _, len := IntFromString(s, base);
    <a id="L32"></a>if slen != nil {
        <a id="L33"></a>*slen = len
    <a id="L34"></a>}
    <a id="L35"></a>return x;
<a id="L36"></a>}


<a id="L39"></a>func ratFromString(s string, base uint, slen *int) *Rational {
    <a id="L40"></a>x, _, len := RatFromString(s, base);
    <a id="L41"></a>if slen != nil {
        <a id="L42"></a>*slen = len
    <a id="L43"></a>}
    <a id="L44"></a>return x;
<a id="L45"></a>}


<a id="L48"></a>var (
    <a id="L49"></a>nat_zero = Nat(0);
    <a id="L50"></a>nat_one  = Nat(1);
    <a id="L51"></a>nat_two  = Nat(2);
    <a id="L52"></a>a        = natFromString(sa, 10, nil);
    <a id="L53"></a>b        = natFromString(sb, 10, nil);
    <a id="L54"></a>c        = natFromString(sc, 10, nil);
    <a id="L55"></a>p        = natFromString(sp, 10, nil);
    <a id="L56"></a>int_zero = Int(0);
    <a id="L57"></a>int_one  = Int(1);
    <a id="L58"></a>int_two  = Int(2);
    <a id="L59"></a>ip       = intFromString(sp, 10, nil);
    <a id="L60"></a>rat_zero = Rat(0, 1);
    <a id="L61"></a>rat_half = Rat(1, 2);
    <a id="L62"></a>rat_one  = Rat(1, 1);
    <a id="L63"></a>rat_two  = Rat(2, 1);
<a id="L64"></a>)


<a id="L67"></a>var test_msg string
<a id="L68"></a>var tester *testing.T

<a id="L70"></a>func test(n uint, b bool) {
    <a id="L71"></a>if !b {
        <a id="L72"></a>tester.Fatalf(&#34;TEST failed: %s (%d)&#34;, test_msg, n)
    <a id="L73"></a>}
<a id="L74"></a>}


<a id="L77"></a>func nat_eq(n uint, x, y Natural) {
    <a id="L78"></a>if x.Cmp(y) != 0 {
        <a id="L79"></a>tester.Fatalf(&#34;TEST failed: %s (%d)\nx = %v\ny = %v&#34;, test_msg, n, &amp;x, &amp;y)
    <a id="L80"></a>}
<a id="L81"></a>}


<a id="L84"></a>func int_eq(n uint, x, y *Integer) {
    <a id="L85"></a>if x.Cmp(y) != 0 {
        <a id="L86"></a>tester.Fatalf(&#34;TEST failed: %s (%d)\nx = %v\ny = %v&#34;, test_msg, n, x, y)
    <a id="L87"></a>}
<a id="L88"></a>}


<a id="L91"></a>func rat_eq(n uint, x, y *Rational) {
    <a id="L92"></a>if x.Cmp(y) != 0 {
        <a id="L93"></a>tester.Fatalf(&#34;TEST failed: %s (%d)\nx = %v\ny = %v&#34;, test_msg, n, x, y)
    <a id="L94"></a>}
<a id="L95"></a>}


<a id="L98"></a>func TestNatConv(t *testing.T) {
    <a id="L99"></a>tester = t;
    <a id="L100"></a>test_msg = &#34;NatConvA&#34;;
    <a id="L101"></a>type entry1 struct {
        <a id="L102"></a>x   uint64;
        <a id="L103"></a>s   string;
    <a id="L104"></a>}
    <a id="L105"></a>tab := []entry1{
        <a id="L106"></a>entry1{0, &#34;0&#34;},
        <a id="L107"></a>entry1{255, &#34;255&#34;},
        <a id="L108"></a>entry1{65535, &#34;65535&#34;},
        <a id="L109"></a>entry1{4294967295, &#34;4294967295&#34;},
        <a id="L110"></a>entry1{18446744073709551615, &#34;18446744073709551615&#34;},
    <a id="L111"></a>};
    <a id="L112"></a>for i, e := range tab {
        <a id="L113"></a>test(100+uint(i), Nat(e.x).String() == e.s);
        <a id="L114"></a>test(200+uint(i), natFromString(e.s, 0, nil).Value() == e.x);
    <a id="L115"></a>}

    <a id="L117"></a>test_msg = &#34;NatConvB&#34;;
    <a id="L118"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L119"></a>test(i, Nat(uint64(i)).String() == fmt.Sprintf(&#34;%d&#34;, i))
    <a id="L120"></a>}

    <a id="L122"></a>test_msg = &#34;NatConvC&#34;;
    <a id="L123"></a>z := uint64(7);
    <a id="L124"></a>for i := uint(0); i &lt;= 64; i++ {
        <a id="L125"></a>test(i, Nat(z).Value() == z);
        <a id="L126"></a>z &lt;&lt;= 1;
    <a id="L127"></a>}

    <a id="L129"></a>test_msg = &#34;NatConvD&#34;;
    <a id="L130"></a>nat_eq(0, a, Nat(991));
    <a id="L131"></a>nat_eq(1, b, Fact(20));
    <a id="L132"></a>nat_eq(2, c, Fact(100));
    <a id="L133"></a>test(3, a.String() == sa);
    <a id="L134"></a>test(4, b.String() == sb);
    <a id="L135"></a>test(5, c.String() == sc);

    <a id="L137"></a>test_msg = &#34;NatConvE&#34;;
    <a id="L138"></a>var slen int;
    <a id="L139"></a>nat_eq(10, natFromString(&#34;0&#34;, 0, nil), nat_zero);
    <a id="L140"></a>nat_eq(11, natFromString(&#34;123&#34;, 0, nil), Nat(123));
    <a id="L141"></a>nat_eq(12, natFromString(&#34;077&#34;, 0, nil), Nat(7*8+7));
    <a id="L142"></a>nat_eq(13, natFromString(&#34;0x1f&#34;, 0, nil), Nat(1*16+15));
    <a id="L143"></a>nat_eq(14, natFromString(&#34;0x1fg&#34;, 0, &amp;slen), Nat(1*16+15));
    <a id="L144"></a>test(4, slen == 4);

    <a id="L146"></a>test_msg = &#34;NatConvF&#34;;
    <a id="L147"></a>tmp := c.Mul(c);
    <a id="L148"></a>for base := uint(2); base &lt;= 16; base++ {
        <a id="L149"></a>nat_eq(base, natFromString(tmp.ToString(base), base, nil), tmp)
    <a id="L150"></a>}

    <a id="L152"></a>test_msg = &#34;NatConvG&#34;;
    <a id="L153"></a>x := Nat(100);
    <a id="L154"></a>y, _, _ := NatFromString(fmt.Sprintf(&#34;%b&#34;, &amp;x), 2);
    <a id="L155"></a>nat_eq(100, y, x);
<a id="L156"></a>}


<a id="L159"></a>func abs(x int64) uint64 {
    <a id="L160"></a>if x &lt; 0 {
        <a id="L161"></a>x = -x
    <a id="L162"></a>}
    <a id="L163"></a>return uint64(x);
<a id="L164"></a>}


<a id="L167"></a>func TestIntConv(t *testing.T) {
    <a id="L168"></a>tester = t;
    <a id="L169"></a>test_msg = &#34;IntConvA&#34;;
    <a id="L170"></a>type entry2 struct {
        <a id="L171"></a>x   int64;
        <a id="L172"></a>s   string;
    <a id="L173"></a>}
    <a id="L174"></a>tab := []entry2{
        <a id="L175"></a>entry2{0, &#34;0&#34;},
        <a id="L176"></a>entry2{-128, &#34;-128&#34;},
        <a id="L177"></a>entry2{127, &#34;127&#34;},
        <a id="L178"></a>entry2{-32768, &#34;-32768&#34;},
        <a id="L179"></a>entry2{32767, &#34;32767&#34;},
        <a id="L180"></a>entry2{-2147483648, &#34;-2147483648&#34;},
        <a id="L181"></a>entry2{2147483647, &#34;2147483647&#34;},
        <a id="L182"></a>entry2{-9223372036854775808, &#34;-9223372036854775808&#34;},
        <a id="L183"></a>entry2{9223372036854775807, &#34;9223372036854775807&#34;},
    <a id="L184"></a>};
    <a id="L185"></a>for i, e := range tab {
        <a id="L186"></a>test(100+uint(i), Int(e.x).String() == e.s);
        <a id="L187"></a>test(200+uint(i), intFromString(e.s, 0, nil).Value() == e.x);
        <a id="L188"></a>test(300+uint(i), Int(e.x).Abs().Value() == abs(e.x));
    <a id="L189"></a>}

    <a id="L191"></a>test_msg = &#34;IntConvB&#34;;
    <a id="L192"></a>var slen int;
    <a id="L193"></a>int_eq(0, intFromString(&#34;0&#34;, 0, nil), int_zero);
    <a id="L194"></a>int_eq(1, intFromString(&#34;-0&#34;, 0, nil), int_zero);
    <a id="L195"></a>int_eq(2, intFromString(&#34;123&#34;, 0, nil), Int(123));
    <a id="L196"></a>int_eq(3, intFromString(&#34;-123&#34;, 0, nil), Int(-123));
    <a id="L197"></a>int_eq(4, intFromString(&#34;077&#34;, 0, nil), Int(7*8+7));
    <a id="L198"></a>int_eq(5, intFromString(&#34;-077&#34;, 0, nil), Int(-(7*8 + 7)));
    <a id="L199"></a>int_eq(6, intFromString(&#34;0x1f&#34;, 0, nil), Int(1*16+15));
    <a id="L200"></a>int_eq(7, intFromString(&#34;-0x1f&#34;, 0, &amp;slen), Int(-(1*16 + 15)));
    <a id="L201"></a>test(7, slen == 5);
    <a id="L202"></a>int_eq(8, intFromString(&#34;+0x1f&#34;, 0, &amp;slen), Int(+(1*16 + 15)));
    <a id="L203"></a>test(8, slen == 5);
    <a id="L204"></a>int_eq(9, intFromString(&#34;0x1fg&#34;, 0, &amp;slen), Int(1*16+15));
    <a id="L205"></a>test(9, slen == 4);
    <a id="L206"></a>int_eq(10, intFromString(&#34;-0x1fg&#34;, 0, &amp;slen), Int(-(1*16 + 15)));
    <a id="L207"></a>test(10, slen == 5);
<a id="L208"></a>}


<a id="L211"></a>func TestRatConv(t *testing.T) {
    <a id="L212"></a>tester = t;
    <a id="L213"></a>test_msg = &#34;RatConv&#34;;
    <a id="L214"></a>var slen int;
    <a id="L215"></a>rat_eq(0, ratFromString(&#34;0&#34;, 0, nil), rat_zero);
    <a id="L216"></a>rat_eq(1, ratFromString(&#34;0/1&#34;, 0, nil), rat_zero);
    <a id="L217"></a>rat_eq(2, ratFromString(&#34;0/01&#34;, 0, nil), rat_zero);
    <a id="L218"></a>rat_eq(3, ratFromString(&#34;0x14/10&#34;, 0, &amp;slen), rat_two);
    <a id="L219"></a>test(4, slen == 7);
    <a id="L220"></a>rat_eq(5, ratFromString(&#34;0.&#34;, 0, nil), rat_zero);
    <a id="L221"></a>rat_eq(6, ratFromString(&#34;0.001f&#34;, 10, nil), Rat(1, 1000));
    <a id="L222"></a>rat_eq(7, ratFromString(&#34;.1&#34;, 0, nil), Rat(1, 10));
    <a id="L223"></a>rat_eq(8, ratFromString(&#34;10101.0101&#34;, 2, nil), Rat(0x155, 1&lt;&lt;4));
    <a id="L224"></a>rat_eq(9, ratFromString(&#34;-0003.145926&#34;, 10, &amp;slen), Rat(-3145926, 1000000));
    <a id="L225"></a>test(10, slen == 12);
    <a id="L226"></a>rat_eq(11, ratFromString(&#34;1e2&#34;, 0, nil), Rat(100, 1));
    <a id="L227"></a>rat_eq(12, ratFromString(&#34;1e-2&#34;, 0, nil), Rat(1, 100));
    <a id="L228"></a>rat_eq(13, ratFromString(&#34;1.1e2&#34;, 0, nil), Rat(110, 1));
    <a id="L229"></a>rat_eq(14, ratFromString(&#34;.1e2x&#34;, 0, &amp;slen), Rat(10, 1));
    <a id="L230"></a>test(15, slen == 4);
<a id="L231"></a>}


<a id="L234"></a>func add(x, y Natural) Natural {
    <a id="L235"></a>z1 := x.Add(y);
    <a id="L236"></a>z2 := y.Add(x);
    <a id="L237"></a>if z1.Cmp(z2) != 0 {
        <a id="L238"></a>tester.Fatalf(&#34;addition not symmetric:\n\tx = %v\n\ty = %t&#34;, x, y)
    <a id="L239"></a>}
    <a id="L240"></a>return z1;
<a id="L241"></a>}


<a id="L244"></a>func sum(n uint64, scale Natural) Natural {
    <a id="L245"></a>s := nat_zero;
    <a id="L246"></a>for ; n &gt; 0; n-- {
        <a id="L247"></a>s = add(s, Nat(n).Mul(scale))
    <a id="L248"></a>}
    <a id="L249"></a>return s;
<a id="L250"></a>}


<a id="L253"></a>func TestNatAdd(t *testing.T) {
    <a id="L254"></a>tester = t;
    <a id="L255"></a>test_msg = &#34;NatAddA&#34;;
    <a id="L256"></a>nat_eq(0, add(nat_zero, nat_zero), nat_zero);
    <a id="L257"></a>nat_eq(1, add(nat_zero, c), c);

    <a id="L259"></a>test_msg = &#34;NatAddB&#34;;
    <a id="L260"></a>for i := uint64(0); i &lt; 100; i++ {
        <a id="L261"></a>t := Nat(i);
        <a id="L262"></a>nat_eq(uint(i), sum(i, c), t.Mul(t).Add(t).Shr(1).Mul(c));
    <a id="L263"></a>}
<a id="L264"></a>}


<a id="L267"></a>func mul(x, y Natural) Natural {
    <a id="L268"></a>z1 := x.Mul(y);
    <a id="L269"></a>z2 := y.Mul(x);
    <a id="L270"></a>if z1.Cmp(z2) != 0 {
        <a id="L271"></a>tester.Fatalf(&#34;multiplication not symmetric:\n\tx = %v\n\ty = %t&#34;, x, y)
    <a id="L272"></a>}
    <a id="L273"></a>if !x.IsZero() &amp;&amp; z1.Div(x).Cmp(y) != 0 {
        <a id="L274"></a>tester.Fatalf(&#34;multiplication/division not inverse (A):\n\tx = %v\n\ty = %t&#34;, x, y)
    <a id="L275"></a>}
    <a id="L276"></a>if !y.IsZero() &amp;&amp; z1.Div(y).Cmp(x) != 0 {
        <a id="L277"></a>tester.Fatalf(&#34;multiplication/division not inverse (B):\n\tx = %v\n\ty = %t&#34;, x, y)
    <a id="L278"></a>}
    <a id="L279"></a>return z1;
<a id="L280"></a>}


<a id="L283"></a>func TestNatSub(t *testing.T) {
    <a id="L284"></a>tester = t;
    <a id="L285"></a>test_msg = &#34;NatSubA&#34;;
    <a id="L286"></a>nat_eq(0, nat_zero.Sub(nat_zero), nat_zero);
    <a id="L287"></a>nat_eq(1, c.Sub(nat_zero), c);

    <a id="L289"></a>test_msg = &#34;NatSubB&#34;;
    <a id="L290"></a>for i := uint64(0); i &lt; 100; i++ {
        <a id="L291"></a>t := sum(i, c);
        <a id="L292"></a>for j := uint64(0); j &lt;= i; j++ {
            <a id="L293"></a>t = t.Sub(mul(Nat(j), c))
        <a id="L294"></a>}
        <a id="L295"></a>nat_eq(uint(i), t, nat_zero);
    <a id="L296"></a>}
<a id="L297"></a>}


<a id="L300"></a>func TestNatMul(t *testing.T) {
    <a id="L301"></a>tester = t;
    <a id="L302"></a>test_msg = &#34;NatMulA&#34;;
    <a id="L303"></a>nat_eq(0, mul(c, nat_zero), nat_zero);
    <a id="L304"></a>nat_eq(1, mul(c, nat_one), c);

    <a id="L306"></a>test_msg = &#34;NatMulB&#34;;
    <a id="L307"></a>nat_eq(0, b.Mul(MulRange(0, 100)), nat_zero);
    <a id="L308"></a>nat_eq(1, b.Mul(MulRange(21, 100)), c);

    <a id="L310"></a>test_msg = &#34;NatMulC&#34;;
    <a id="L311"></a>const n = 100;
    <a id="L312"></a>p := b.Mul(c).Shl(n);
    <a id="L313"></a>for i := uint(0); i &lt; n; i++ {
        <a id="L314"></a>nat_eq(i, mul(b.Shl(i), c.Shl(n-i)), p)
    <a id="L315"></a>}
<a id="L316"></a>}


<a id="L319"></a>func TestNatDiv(t *testing.T) {
    <a id="L320"></a>tester = t;
    <a id="L321"></a>test_msg = &#34;NatDivA&#34;;
    <a id="L322"></a>nat_eq(0, c.Div(nat_one), c);
    <a id="L323"></a>nat_eq(1, c.Div(Nat(100)), Fact(99));
    <a id="L324"></a>nat_eq(2, b.Div(c), nat_zero);
    <a id="L325"></a>nat_eq(4, nat_one.Shl(100).Div(nat_one.Shl(90)), nat_one.Shl(10));
    <a id="L326"></a>nat_eq(5, c.Div(b), MulRange(21, 100));

    <a id="L328"></a>test_msg = &#34;NatDivB&#34;;
    <a id="L329"></a>const n = 100;
    <a id="L330"></a>p := Fact(n);
    <a id="L331"></a>for i := uint(0); i &lt; n; i++ {
        <a id="L332"></a>nat_eq(100+i, p.Div(MulRange(1, i)), MulRange(i+1, n))
    <a id="L333"></a>}
<a id="L334"></a>}


<a id="L337"></a>func TestIntQuoRem(t *testing.T) {
    <a id="L338"></a>tester = t;
    <a id="L339"></a>test_msg = &#34;IntQuoRem&#34;;
    <a id="L340"></a>type T struct {
        <a id="L341"></a>x, y, q, r int64;
    <a id="L342"></a>}
    <a id="L343"></a>a := []T{
        <a id="L344"></a>T{+8, +3, +2, +2},
        <a id="L345"></a>T{+8, -3, -2, +2},
        <a id="L346"></a>T{-8, +3, -2, -2},
        <a id="L347"></a>T{-8, -3, +2, -2},
        <a id="L348"></a>T{+1, +2, 0, +1},
        <a id="L349"></a>T{+1, -2, 0, +1},
        <a id="L350"></a>T{-1, +2, 0, -1},
        <a id="L351"></a>T{-1, -2, 0, -1},
    <a id="L352"></a>};
    <a id="L353"></a>for i := uint(0); i &lt; uint(len(a)); i++ {
        <a id="L354"></a>e := &amp;a[i];
        <a id="L355"></a>x, y := Int(e.x).Mul(ip), Int(e.y).Mul(ip);
        <a id="L356"></a>q, r := Int(e.q), Int(e.r).Mul(ip);
        <a id="L357"></a>qq, rr := x.QuoRem(y);
        <a id="L358"></a>int_eq(4*i+0, x.Quo(y), q);
        <a id="L359"></a>int_eq(4*i+1, x.Rem(y), r);
        <a id="L360"></a>int_eq(4*i+2, qq, q);
        <a id="L361"></a>int_eq(4*i+3, rr, r);
    <a id="L362"></a>}
<a id="L363"></a>}


<a id="L366"></a>func TestIntDivMod(t *testing.T) {
    <a id="L367"></a>tester = t;
    <a id="L368"></a>test_msg = &#34;IntDivMod&#34;;
    <a id="L369"></a>type T struct {
        <a id="L370"></a>x, y, q, r int64;
    <a id="L371"></a>}
    <a id="L372"></a>a := []T{
        <a id="L373"></a>T{+8, +3, +2, +2},
        <a id="L374"></a>T{+8, -3, -2, +2},
        <a id="L375"></a>T{-8, +3, -3, +1},
        <a id="L376"></a>T{-8, -3, +3, +1},
        <a id="L377"></a>T{+1, +2, 0, +1},
        <a id="L378"></a>T{+1, -2, 0, +1},
        <a id="L379"></a>T{-1, +2, -1, +1},
        <a id="L380"></a>T{-1, -2, +1, +1},
    <a id="L381"></a>};
    <a id="L382"></a>for i := uint(0); i &lt; uint(len(a)); i++ {
        <a id="L383"></a>e := &amp;a[i];
        <a id="L384"></a>x, y := Int(e.x).Mul(ip), Int(e.y).Mul(ip);
        <a id="L385"></a>q, r := Int(e.q), Int(e.r).Mul(ip);
        <a id="L386"></a>qq, rr := x.DivMod(y);
        <a id="L387"></a>int_eq(4*i+0, x.Div(y), q);
        <a id="L388"></a>int_eq(4*i+1, x.Mod(y), r);
        <a id="L389"></a>int_eq(4*i+2, qq, q);
        <a id="L390"></a>int_eq(4*i+3, rr, r);
    <a id="L391"></a>}
<a id="L392"></a>}


<a id="L395"></a>func TestNatMod(t *testing.T) {
    <a id="L396"></a>tester = t;
    <a id="L397"></a>test_msg = &#34;NatModA&#34;;
    <a id="L398"></a>for i := uint(0); ; i++ {
        <a id="L399"></a>d := nat_one.Shl(i);
        <a id="L400"></a>if d.Cmp(c) &lt; 0 {
            <a id="L401"></a>nat_eq(i, c.Add(d).Mod(c), d)
        <a id="L402"></a>} else {
            <a id="L403"></a>nat_eq(i, c.Add(d).Div(c), nat_two);
            <a id="L404"></a>nat_eq(i, c.Add(d).Mod(c), d.Sub(c));
            <a id="L405"></a>break;
        <a id="L406"></a>}
    <a id="L407"></a>}
<a id="L408"></a>}


<a id="L411"></a>func TestNatShift(t *testing.T) {
    <a id="L412"></a>tester = t;
    <a id="L413"></a>test_msg = &#34;NatShift1L&#34;;
    <a id="L414"></a>test(0, b.Shl(0).Cmp(b) == 0);
    <a id="L415"></a>test(1, c.Shl(1).Cmp(c) &gt; 0);

    <a id="L417"></a>test_msg = &#34;NatShift1R&#34;;
    <a id="L418"></a>test(3, b.Shr(0).Cmp(b) == 0);
    <a id="L419"></a>test(4, c.Shr(1).Cmp(c) &lt; 0);

    <a id="L421"></a>test_msg = &#34;NatShift2&#34;;
    <a id="L422"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L423"></a>test(i, c.Shl(i).Shr(i).Cmp(c) == 0)
    <a id="L424"></a>}

    <a id="L426"></a>test_msg = &#34;NatShift3L&#34;;
    <a id="L427"></a>{
        <a id="L428"></a>const m = 3;
        <a id="L429"></a>p := b;
        <a id="L430"></a>f := Nat(1 &lt;&lt; m);
        <a id="L431"></a>for i := uint(0); i &lt; 100; i++ {
            <a id="L432"></a>nat_eq(i, b.Shl(i*m), p);
            <a id="L433"></a>p = mul(p, f);
        <a id="L434"></a>}
    <a id="L435"></a>}

    <a id="L437"></a>test_msg = &#34;NatShift3R&#34;;
    <a id="L438"></a>{
        <a id="L439"></a>p := c;
        <a id="L440"></a>for i := uint(0); !p.IsZero(); i++ {
            <a id="L441"></a>nat_eq(i, c.Shr(i), p);
            <a id="L442"></a>p = p.Shr(1);
        <a id="L443"></a>}
    <a id="L444"></a>}
<a id="L445"></a>}


<a id="L448"></a>func TestIntShift(t *testing.T) {
    <a id="L449"></a>tester = t;
    <a id="L450"></a>test_msg = &#34;IntShift1L&#34;;
    <a id="L451"></a>test(0, ip.Shl(0).Cmp(ip) == 0);
    <a id="L452"></a>test(1, ip.Shl(1).Cmp(ip) &gt; 0);

    <a id="L454"></a>test_msg = &#34;IntShift1R&#34;;
    <a id="L455"></a>test(0, ip.Shr(0).Cmp(ip) == 0);
    <a id="L456"></a>test(1, ip.Shr(1).Cmp(ip) &lt; 0);

    <a id="L458"></a>test_msg = &#34;IntShift2&#34;;
    <a id="L459"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L460"></a>test(i, ip.Shl(i).Shr(i).Cmp(ip) == 0)
    <a id="L461"></a>}

    <a id="L463"></a>test_msg = &#34;IntShift3L&#34;;
    <a id="L464"></a>{
        <a id="L465"></a>const m = 3;
        <a id="L466"></a>p := ip;
        <a id="L467"></a>f := Int(1 &lt;&lt; m);
        <a id="L468"></a>for i := uint(0); i &lt; 100; i++ {
            <a id="L469"></a>int_eq(i, ip.Shl(i*m), p);
            <a id="L470"></a>p = p.Mul(f);
        <a id="L471"></a>}
    <a id="L472"></a>}

    <a id="L474"></a>test_msg = &#34;IntShift3R&#34;;
    <a id="L475"></a>{
        <a id="L476"></a>p := ip;
        <a id="L477"></a>for i := uint(0); p.IsPos(); i++ {
            <a id="L478"></a>int_eq(i, ip.Shr(i), p);
            <a id="L479"></a>p = p.Shr(1);
        <a id="L480"></a>}
    <a id="L481"></a>}

    <a id="L483"></a>test_msg = &#34;IntShift4R&#34;;
    <a id="L484"></a>int_eq(0, Int(-43).Shr(1), Int(-43&gt;&gt;1));
    <a id="L485"></a>int_eq(0, Int(-1024).Shr(100), Int(-1));
    <a id="L486"></a>int_eq(1, ip.Neg().Shr(10), ip.Neg().Div(Int(1).Shl(10)));
<a id="L487"></a>}


<a id="L490"></a>func TestNatBitOps(t *testing.T) {
    <a id="L491"></a>tester = t;

    <a id="L493"></a>x := uint64(0xf08e6f56bd8c3941);
    <a id="L494"></a>y := uint64(0x3984ef67834bc);

    <a id="L496"></a>bx := Nat(x);
    <a id="L497"></a>by := Nat(y);

    <a id="L499"></a>test_msg = &#34;NatAnd&#34;;
    <a id="L500"></a>bz := Nat(x &amp; y);
    <a id="L501"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L502"></a>nat_eq(i, bx.Shl(i).And(by.Shl(i)), bz.Shl(i))
    <a id="L503"></a>}

    <a id="L505"></a>test_msg = &#34;NatAndNot&#34;;
    <a id="L506"></a>bz = Nat(x &amp;^ y);
    <a id="L507"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L508"></a>nat_eq(i, bx.Shl(i).AndNot(by.Shl(i)), bz.Shl(i))
    <a id="L509"></a>}

    <a id="L511"></a>test_msg = &#34;NatOr&#34;;
    <a id="L512"></a>bz = Nat(x | y);
    <a id="L513"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L514"></a>nat_eq(i, bx.Shl(i).Or(by.Shl(i)), bz.Shl(i))
    <a id="L515"></a>}

    <a id="L517"></a>test_msg = &#34;NatXor&#34;;
    <a id="L518"></a>bz = Nat(x ^ y);
    <a id="L519"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L520"></a>nat_eq(i, bx.Shl(i).Xor(by.Shl(i)), bz.Shl(i))
    <a id="L521"></a>}
<a id="L522"></a>}


<a id="L525"></a>func TestIntBitOps1(t *testing.T) {
    <a id="L526"></a>tester = t;
    <a id="L527"></a>test_msg = &#34;IntBitOps1&#34;;
    <a id="L528"></a>type T struct {
        <a id="L529"></a>x, y int64;
    <a id="L530"></a>}
    <a id="L531"></a>a := []T{
        <a id="L532"></a>T{+7, +3},
        <a id="L533"></a>T{+7, -3},
        <a id="L534"></a>T{-7, +3},
        <a id="L535"></a>T{-7, -3},
    <a id="L536"></a>};
    <a id="L537"></a>for i := uint(0); i &lt; uint(len(a)); i++ {
        <a id="L538"></a>e := &amp;a[i];
        <a id="L539"></a>int_eq(4*i+0, Int(e.x).And(Int(e.y)), Int(e.x&amp;e.y));
        <a id="L540"></a>int_eq(4*i+1, Int(e.x).AndNot(Int(e.y)), Int(e.x&amp;^e.y));
        <a id="L541"></a>int_eq(4*i+2, Int(e.x).Or(Int(e.y)), Int(e.x|e.y));
        <a id="L542"></a>int_eq(4*i+3, Int(e.x).Xor(Int(e.y)), Int(e.x^e.y));
    <a id="L543"></a>}
<a id="L544"></a>}


<a id="L547"></a>func TestIntBitOps2(t *testing.T) {
    <a id="L548"></a>tester = t;

    <a id="L550"></a>test_msg = &#34;IntNot&#34;;
    <a id="L551"></a>int_eq(0, Int(-2).Not(), Int(1));
    <a id="L552"></a>int_eq(0, Int(-1).Not(), Int(0));
    <a id="L553"></a>int_eq(0, Int(0).Not(), Int(-1));
    <a id="L554"></a>int_eq(0, Int(1).Not(), Int(-2));
    <a id="L555"></a>int_eq(0, Int(2).Not(), Int(-3));

    <a id="L557"></a>test_msg = &#34;IntAnd&#34;;
    <a id="L558"></a>for x := int64(-15); x &lt; 5; x++ {
        <a id="L559"></a>bx := Int(x);
        <a id="L560"></a>for y := int64(-5); y &lt; 15; y++ {
            <a id="L561"></a>by := Int(y);
            <a id="L562"></a>for i := uint(50); i &lt; 70; i++ { <span class="comment">// shift across 64bit boundary</span>
                <a id="L563"></a>int_eq(i, bx.Shl(i).And(by.Shl(i)), Int(x&amp;y).Shl(i))
            <a id="L564"></a>}
        <a id="L565"></a>}
    <a id="L566"></a>}

    <a id="L568"></a>test_msg = &#34;IntAndNot&#34;;
    <a id="L569"></a>for x := int64(-15); x &lt; 5; x++ {
        <a id="L570"></a>bx := Int(x);
        <a id="L571"></a>for y := int64(-5); y &lt; 15; y++ {
            <a id="L572"></a>by := Int(y);
            <a id="L573"></a>for i := uint(50); i &lt; 70; i++ { <span class="comment">// shift across 64bit boundary</span>
                <a id="L574"></a>int_eq(2*i+0, bx.Shl(i).AndNot(by.Shl(i)), Int(x&amp;^y).Shl(i));
                <a id="L575"></a>int_eq(2*i+1, bx.Shl(i).And(by.Shl(i).Not()), Int(x&amp;^y).Shl(i));
            <a id="L576"></a>}
        <a id="L577"></a>}
    <a id="L578"></a>}

    <a id="L580"></a>test_msg = &#34;IntOr&#34;;
    <a id="L581"></a>for x := int64(-15); x &lt; 5; x++ {
        <a id="L582"></a>bx := Int(x);
        <a id="L583"></a>for y := int64(-5); y &lt; 15; y++ {
            <a id="L584"></a>by := Int(y);
            <a id="L585"></a>for i := uint(50); i &lt; 70; i++ { <span class="comment">// shift across 64bit boundary</span>
                <a id="L586"></a>int_eq(i, bx.Shl(i).Or(by.Shl(i)), Int(x|y).Shl(i))
            <a id="L587"></a>}
        <a id="L588"></a>}
    <a id="L589"></a>}

    <a id="L591"></a>test_msg = &#34;IntXor&#34;;
    <a id="L592"></a>for x := int64(-15); x &lt; 5; x++ {
        <a id="L593"></a>bx := Int(x);
        <a id="L594"></a>for y := int64(-5); y &lt; 15; y++ {
            <a id="L595"></a>by := Int(y);
            <a id="L596"></a>for i := uint(50); i &lt; 70; i++ { <span class="comment">// shift across 64bit boundary</span>
                <a id="L597"></a>int_eq(i, bx.Shl(i).Xor(by.Shl(i)), Int(x^y).Shl(i))
            <a id="L598"></a>}
        <a id="L599"></a>}
    <a id="L600"></a>}
<a id="L601"></a>}


<a id="L604"></a>func TestNatCmp(t *testing.T) {
    <a id="L605"></a>tester = t;
    <a id="L606"></a>test_msg = &#34;NatCmp&#34;;
    <a id="L607"></a>test(0, a.Cmp(a) == 0);
    <a id="L608"></a>test(1, a.Cmp(b) &lt; 0);
    <a id="L609"></a>test(2, b.Cmp(a) &gt; 0);
    <a id="L610"></a>test(3, a.Cmp(c) &lt; 0);
    <a id="L611"></a>d := c.Add(b);
    <a id="L612"></a>test(4, c.Cmp(d) &lt; 0);
    <a id="L613"></a>test(5, d.Cmp(c) &gt; 0);
<a id="L614"></a>}


<a id="L617"></a>func TestNatLog2(t *testing.T) {
    <a id="L618"></a>tester = t;
    <a id="L619"></a>test_msg = &#34;NatLog2A&#34;;
    <a id="L620"></a>test(0, nat_one.Log2() == 0);
    <a id="L621"></a>test(1, nat_two.Log2() == 1);
    <a id="L622"></a>test(2, Nat(3).Log2() == 1);
    <a id="L623"></a>test(3, Nat(4).Log2() == 2);

    <a id="L625"></a>test_msg = &#34;NatLog2B&#34;;
    <a id="L626"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L627"></a>test(i, nat_one.Shl(i).Log2() == i)
    <a id="L628"></a>}
<a id="L629"></a>}


<a id="L632"></a>func TestNatGcd(t *testing.T) {
    <a id="L633"></a>tester = t;
    <a id="L634"></a>test_msg = &#34;NatGcdA&#34;;
    <a id="L635"></a>f := Nat(99991);
    <a id="L636"></a>nat_eq(0, b.Mul(f).Gcd(c.Mul(f)), MulRange(1, 20).Mul(f));
<a id="L637"></a>}


<a id="L640"></a>func TestNatPow(t *testing.T) {
    <a id="L641"></a>tester = t;
    <a id="L642"></a>test_msg = &#34;NatPowA&#34;;
    <a id="L643"></a>nat_eq(0, nat_two.Pow(0), nat_one);

    <a id="L645"></a>test_msg = &#34;NatPowB&#34;;
    <a id="L646"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L647"></a>nat_eq(i, nat_two.Pow(i), nat_one.Shl(i))
    <a id="L648"></a>}
<a id="L649"></a>}


<a id="L652"></a>func TestNatPop(t *testing.T) {
    <a id="L653"></a>tester = t;
    <a id="L654"></a>test_msg = &#34;NatPopA&#34;;
    <a id="L655"></a>test(0, nat_zero.Pop() == 0);
    <a id="L656"></a>test(1, nat_one.Pop() == 1);
    <a id="L657"></a>test(2, Nat(10).Pop() == 2);
    <a id="L658"></a>test(3, Nat(30).Pop() == 4);
    <a id="L659"></a>test(4, Nat(0x1248f).Shl(33).Pop() == 8);

    <a id="L661"></a>test_msg = &#34;NatPopB&#34;;
    <a id="L662"></a>for i := uint(0); i &lt; 100; i++ {
        <a id="L663"></a>test(i, nat_one.Shl(i).Sub(nat_one).Pop() == i)
    <a id="L664"></a>}
<a id="L665"></a>}
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
