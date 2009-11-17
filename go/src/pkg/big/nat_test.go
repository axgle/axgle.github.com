<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/big/nat_test.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/big/nat_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package big

<a id="L7"></a>import &#34;testing&#34;

<a id="L9"></a>func TestCmpNN(t *testing.T) {
    <a id="L10"></a><span class="comment">// TODO(gri) write this test - all other tests depends on it</span>
<a id="L11"></a>}


<a id="L14"></a>type funNN func(z, x, y []Word) []Word
<a id="L15"></a>type argNN struct {
    <a id="L16"></a>z, x, y []Word;
<a id="L17"></a>}


<a id="L20"></a>var sumNN = []argNN{
    <a id="L21"></a>argNN{},
    <a id="L22"></a>argNN{[]Word{1}, nil, []Word{1}},
    <a id="L23"></a>argNN{[]Word{1111111110}, []Word{123456789}, []Word{987654321}},
    <a id="L24"></a>argNN{[]Word{0, 0, 0, 1}, nil, []Word{0, 0, 0, 1}},
    <a id="L25"></a>argNN{[]Word{0, 0, 0, 1111111110}, []Word{0, 0, 0, 123456789}, []Word{0, 0, 0, 987654321}},
    <a id="L26"></a>argNN{[]Word{0, 0, 0, 1}, []Word{0, 0, _M}, []Word{0, 0, 1}},
<a id="L27"></a>}


<a id="L30"></a>var prodNN = []argNN{
    <a id="L31"></a>argNN{},
    <a id="L32"></a>argNN{nil, nil, nil},
    <a id="L33"></a>argNN{nil, []Word{991}, nil},
    <a id="L34"></a>argNN{[]Word{991}, []Word{991}, []Word{1}},
    <a id="L35"></a>argNN{[]Word{991 * 991}, []Word{991}, []Word{991}},
    <a id="L36"></a>argNN{[]Word{0, 0, 991 * 991}, []Word{0, 991}, []Word{0, 991}},
    <a id="L37"></a>argNN{[]Word{1 * 991, 2 * 991, 3 * 991, 4 * 991}, []Word{1, 2, 3, 4}, []Word{991}},
    <a id="L38"></a>argNN{[]Word{4, 11, 20, 30, 20, 11, 4}, []Word{1, 2, 3, 4}, []Word{4, 3, 2, 1}},
<a id="L39"></a>}


<a id="L42"></a>func TestSetN(t *testing.T) {
    <a id="L43"></a>for _, a := range sumNN {
        <a id="L44"></a>z := setN(nil, a.z);
        <a id="L45"></a>if cmpNN(z, a.z) != 0 {
            <a id="L46"></a>t.Errorf(&#34;got z = %v; want %v&#34;, z, a.z)
        <a id="L47"></a>}
    <a id="L48"></a>}
<a id="L49"></a>}


<a id="L52"></a>func testFunNN(t *testing.T, msg string, f funNN, a argNN) {
    <a id="L53"></a>z := f(nil, a.x, a.y);
    <a id="L54"></a>if cmpNN(z, a.z) != 0 {
        <a id="L55"></a>t.Errorf(&#34;%s%+v\n\tgot z = %v; want %v&#34;, msg, a, z, a.z)
    <a id="L56"></a>}
<a id="L57"></a>}


<a id="L60"></a>func TestFunNN(t *testing.T) {
    <a id="L61"></a>for _, a := range sumNN {
        <a id="L62"></a>arg := a;
        <a id="L63"></a>testFunNN(t, &#34;addNN&#34;, addNN, arg);

        <a id="L65"></a>arg = argNN{a.z, a.y, a.x};
        <a id="L66"></a>testFunNN(t, &#34;addNN symmetric&#34;, addNN, arg);

        <a id="L68"></a>arg = argNN{a.x, a.z, a.y};
        <a id="L69"></a>testFunNN(t, &#34;subNN&#34;, subNN, arg);

        <a id="L71"></a>arg = argNN{a.y, a.z, a.x};
        <a id="L72"></a>testFunNN(t, &#34;subNN symmetric&#34;, subNN, arg);
    <a id="L73"></a>}

    <a id="L75"></a>for _, a := range prodNN {
        <a id="L76"></a>arg := a;
        <a id="L77"></a>testFunNN(t, &#34;mulNN&#34;, mulNN, arg);

        <a id="L79"></a>arg = argNN{a.z, a.y, a.x};
        <a id="L80"></a>testFunNN(t, &#34;mulNN symmetric&#34;, mulNN, arg);
    <a id="L81"></a>}
<a id="L82"></a>}


<a id="L85"></a>type strN struct {
    <a id="L86"></a>x   []Word;
    <a id="L87"></a>b   int;
    <a id="L88"></a>s   string;
<a id="L89"></a>}


<a id="L92"></a>var tabN = []strN{
    <a id="L93"></a>strN{nil, 10, &#34;0&#34;},
    <a id="L94"></a>strN{[]Word{1}, 10, &#34;1&#34;},
    <a id="L95"></a>strN{[]Word{10}, 10, &#34;10&#34;},
    <a id="L96"></a>strN{[]Word{1234567890}, 10, &#34;1234567890&#34;},
<a id="L97"></a>}


<a id="L100"></a>func TestStringN(t *testing.T) {
    <a id="L101"></a>for _, a := range tabN {
        <a id="L102"></a>s := stringN(a.x, a.b);
        <a id="L103"></a>if s != a.s {
            <a id="L104"></a>t.Errorf(&#34;stringN%+v\n\tgot s = %s; want %s&#34;, a, s, a.s)
        <a id="L105"></a>}

        <a id="L107"></a>x, b, n := scanN(nil, a.s, a.b);
        <a id="L108"></a>if cmpNN(x, a.x) != 0 {
            <a id="L109"></a>t.Errorf(&#34;scanN%+v\n\tgot z = %v; want %v&#34;, a, x, a.x)
        <a id="L110"></a>}
        <a id="L111"></a>if b != a.b {
            <a id="L112"></a>t.Errorf(&#34;scanN%+v\n\tgot b = %d; want %d&#34;, a, b, a.b)
        <a id="L113"></a>}
        <a id="L114"></a>if n != len(a.s) {
            <a id="L115"></a>t.Errorf(&#34;scanN%+v\n\tgot n = %d; want %d&#34;, a, n, len(a.s))
        <a id="L116"></a>}
    <a id="L117"></a>}
<a id="L118"></a>}


<a id="L121"></a>func TestLeadingZeroBits(t *testing.T) {
    <a id="L122"></a>var x Word = 1 &lt;&lt; (_W - 1);
    <a id="L123"></a>for i := 0; i &lt;= _W; i++ {
        <a id="L124"></a>if leadingZeroBits(x) != i {
            <a id="L125"></a>t.Errorf(&#34;failed at %x: got %d want %d&#34;, x, leadingZeroBits(x), i)
        <a id="L126"></a>}
        <a id="L127"></a>x &gt;&gt;= 1;
    <a id="L128"></a>}
<a id="L129"></a>}


<a id="L132"></a>type shiftTest struct {
    <a id="L133"></a>in    []Word;
    <a id="L134"></a>shift int;
    <a id="L135"></a>out   []Word;
<a id="L136"></a>}


<a id="L139"></a>var leftShiftTests = []shiftTest{
    <a id="L140"></a>shiftTest{nil, 0, nil},
    <a id="L141"></a>shiftTest{nil, 1, nil},
    <a id="L142"></a>shiftTest{[]Word{0}, 0, []Word{0}},
    <a id="L143"></a>shiftTest{[]Word{1}, 0, []Word{1}},
    <a id="L144"></a>shiftTest{[]Word{1}, 1, []Word{2}},
    <a id="L145"></a>shiftTest{[]Word{1 &lt;&lt; (_W - 1)}, 1, []Word{0}},
    <a id="L146"></a>shiftTest{[]Word{1 &lt;&lt; (_W - 1), 0}, 1, []Word{0, 1}},
<a id="L147"></a>}


<a id="L150"></a>func TestShiftLeft(t *testing.T) {
    <a id="L151"></a>for i, test := range leftShiftTests {
        <a id="L152"></a>dst := make([]Word, len(test.out));
        <a id="L153"></a>shiftLeft(dst, test.in, test.shift);
        <a id="L154"></a>for j, v := range dst {
            <a id="L155"></a>if test.out[j] != v {
                <a id="L156"></a>t.Errorf(&#34;#%d: got: %v want: %v&#34;, i, dst, test.out);
                <a id="L157"></a>break;
            <a id="L158"></a>}
        <a id="L159"></a>}
    <a id="L160"></a>}
<a id="L161"></a>}


<a id="L164"></a>var rightShiftTests = []shiftTest{
    <a id="L165"></a>shiftTest{nil, 0, nil},
    <a id="L166"></a>shiftTest{nil, 1, nil},
    <a id="L167"></a>shiftTest{[]Word{0}, 0, []Word{0}},
    <a id="L168"></a>shiftTest{[]Word{1}, 0, []Word{1}},
    <a id="L169"></a>shiftTest{[]Word{1}, 1, []Word{0}},
    <a id="L170"></a>shiftTest{[]Word{2}, 1, []Word{1}},
    <a id="L171"></a>shiftTest{[]Word{0, 1}, 1, []Word{1 &lt;&lt; (_W - 1), 0}},
    <a id="L172"></a>shiftTest{[]Word{2, 1, 1}, 1, []Word{1&lt;&lt;(_W-1) + 1, 1 &lt;&lt; (_W - 1), 0}},
<a id="L173"></a>}


<a id="L176"></a>func TestShiftRight(t *testing.T) {
    <a id="L177"></a>for i, test := range rightShiftTests {
        <a id="L178"></a>dst := make([]Word, len(test.out));
        <a id="L179"></a>shiftRight(dst, test.in, test.shift);
        <a id="L180"></a>for j, v := range dst {
            <a id="L181"></a>if test.out[j] != v {
                <a id="L182"></a>t.Errorf(&#34;#%d: got: %v want: %v&#34;, i, dst, test.out);
                <a id="L183"></a>break;
            <a id="L184"></a>}
        <a id="L185"></a>}
    <a id="L186"></a>}
<a id="L187"></a>}


<a id="L190"></a>type modNWTest struct {
    <a id="L191"></a>in       string;
    <a id="L192"></a>dividend string;
    <a id="L193"></a>out      string;
<a id="L194"></a>}


<a id="L197"></a>var modNWTests32 = []modNWTest{
    <a id="L198"></a>modNWTest{&#34;23492635982634928349238759823742&#34;, &#34;252341&#34;, &#34;220170&#34;},
<a id="L199"></a>}


<a id="L202"></a>var modNWTests64 = []modNWTest{
    <a id="L203"></a>modNWTest{&#34;6527895462947293856291561095690465243862946&#34;, &#34;524326975699234&#34;, &#34;375066989628668&#34;},
<a id="L204"></a>}


<a id="L207"></a>func runModNWTests(t *testing.T, tests []modNWTest) {
    <a id="L208"></a>for i, test := range tests {
        <a id="L209"></a>in, _ := new(Int).SetString(test.in, 10);
        <a id="L210"></a>d, _ := new(Int).SetString(test.dividend, 10);
        <a id="L211"></a>out, _ := new(Int).SetString(test.out, 10);

        <a id="L213"></a>r := modNW(in.abs, d.abs[0]);
        <a id="L214"></a>if r != out.abs[0] {
            <a id="L215"></a>t.Errorf(&#34;#%d failed: got %s want %s\n&#34;, i, r, out)
        <a id="L216"></a>}
    <a id="L217"></a>}
<a id="L218"></a>}


<a id="L221"></a>func TestModNW(t *testing.T) {
    <a id="L222"></a>if _W &gt;= 32 {
        <a id="L223"></a>runModNWTests(t, modNWTests32)
    <a id="L224"></a>}
    <a id="L225"></a>if _W &gt;= 64 {
        <a id="L226"></a>runModNWTests(t, modNWTests32)
    <a id="L227"></a>}
<a id="L228"></a>}


<a id="L231"></a>func TestTrailingZeroBits(t *testing.T) {
    <a id="L232"></a>var x Word;
    <a id="L233"></a>x--;
    <a id="L234"></a>for i := 0; i &lt; _W; i++ {
        <a id="L235"></a>if trailingZeroBits(x) != i {
            <a id="L236"></a>t.Errorf(&#34;Failed at step %d: x: %x got: %d\n&#34;, i, x, trailingZeroBits(x))
        <a id="L237"></a>}
        <a id="L238"></a>x &lt;&lt;= 1;
    <a id="L239"></a>}
<a id="L240"></a>}


<a id="L243"></a>type expNNNTest struct {
    <a id="L244"></a>x, y, m string;
    <a id="L245"></a>out     string;
<a id="L246"></a>}


<a id="L249"></a>var expNNNTests = []expNNNTest{
    <a id="L250"></a>expNNNTest{&#34;0x8000000000000000&#34;, &#34;2&#34;, &#34;&#34;, &#34;0x40000000000000000000000000000000&#34;},
    <a id="L251"></a>expNNNTest{&#34;0x8000000000000000&#34;, &#34;2&#34;, &#34;6719&#34;, &#34;4944&#34;},
    <a id="L252"></a>expNNNTest{&#34;0x8000000000000000&#34;, &#34;3&#34;, &#34;6719&#34;, &#34;5447&#34;},
    <a id="L253"></a>expNNNTest{&#34;0x8000000000000000&#34;, &#34;1000&#34;, &#34;6719&#34;, &#34;1603&#34;},
    <a id="L254"></a>expNNNTest{&#34;0x8000000000000000&#34;, &#34;1000000&#34;, &#34;6719&#34;, &#34;3199&#34;},
    <a id="L255"></a>expNNNTest{
        <a id="L256"></a>&#34;2938462938472983472983659726349017249287491026512746239764525612965293865296239471239874193284792387498274256129746192347&#34;,
        <a id="L257"></a>&#34;298472983472983471903246121093472394872319615612417471234712061&#34;,
        <a id="L258"></a>&#34;29834729834729834729347290846729561262544958723956495615629569234729836259263598127342374289365912465901365498236492183464&#34;,
        <a id="L259"></a>&#34;23537740700184054162508175125554701713153216681790245129157191391322321508055833908509185839069455749219131480588829346291&#34;,
    <a id="L260"></a>},
<a id="L261"></a>}


<a id="L264"></a>func TestExpNNN(t *testing.T) {
    <a id="L265"></a>for i, test := range expNNNTests {
        <a id="L266"></a>x, _, _ := scanN(nil, test.x, 0);
        <a id="L267"></a>y, _, _ := scanN(nil, test.y, 0);
        <a id="L268"></a>out, _, _ := scanN(nil, test.out, 0);

        <a id="L270"></a>var m []Word;

        <a id="L272"></a>if len(test.m) &gt; 0 {
            <a id="L273"></a>m, _, _ = scanN(nil, test.m, 0)
        <a id="L274"></a>}

        <a id="L276"></a>z := expNNN(nil, x, y, m);
        <a id="L277"></a>if cmpNN(z, out) != 0 {
            <a id="L278"></a>t.Errorf(&#34;#%d got %v want %v&#34;, i, z, out)
        <a id="L279"></a>}
    <a id="L280"></a>}
<a id="L281"></a>}
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
