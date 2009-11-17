<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/big/int_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/big/int_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package big

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;encoding/hex&#34;;
    <a id="L10"></a>&#34;testing&#34;;
    <a id="L11"></a>&#34;testing/quick&#34;;
<a id="L12"></a>)

<a id="L14"></a>func newZ(x int64) *Int {
    <a id="L15"></a>var z Int;
    <a id="L16"></a>return z.New(x);
<a id="L17"></a>}


<a id="L20"></a>type funZZ func(z, x, y *Int) *Int
<a id="L21"></a>type argZZ struct {
    <a id="L22"></a>z, x, y *Int;
<a id="L23"></a>}


<a id="L26"></a>var sumZZ = []argZZ{
    <a id="L27"></a>argZZ{newZ(0), newZ(0), newZ(0)},
    <a id="L28"></a>argZZ{newZ(1), newZ(1), newZ(0)},
    <a id="L29"></a>argZZ{newZ(1111111110), newZ(123456789), newZ(987654321)},
    <a id="L30"></a>argZZ{newZ(-1), newZ(-1), newZ(0)},
    <a id="L31"></a>argZZ{newZ(864197532), newZ(-123456789), newZ(987654321)},
    <a id="L32"></a>argZZ{newZ(-1111111110), newZ(-123456789), newZ(-987654321)},
<a id="L33"></a>}


<a id="L36"></a>var prodZZ = []argZZ{
    <a id="L37"></a>argZZ{newZ(0), newZ(0), newZ(0)},
    <a id="L38"></a>argZZ{newZ(0), newZ(1), newZ(0)},
    <a id="L39"></a>argZZ{newZ(1), newZ(1), newZ(1)},
    <a id="L40"></a>argZZ{newZ(-991 * 991), newZ(991), newZ(-991)},
    <a id="L41"></a><span class="comment">// TODO(gri) add larger products</span>
<a id="L42"></a>}


<a id="L45"></a>func TestSetZ(t *testing.T) {
    <a id="L46"></a>for _, a := range sumZZ {
        <a id="L47"></a>var z Int;
        <a id="L48"></a>z.Set(a.z);
        <a id="L49"></a>if (&amp;z).Cmp(a.z) != 0 {
            <a id="L50"></a>t.Errorf(&#34;got z = %v; want %v&#34;, z, a.z)
        <a id="L51"></a>}
    <a id="L52"></a>}
<a id="L53"></a>}


<a id="L56"></a>func testFunZZ(t *testing.T, msg string, f funZZ, a argZZ) {
    <a id="L57"></a>var z Int;
    <a id="L58"></a>f(&amp;z, a.x, a.y);
    <a id="L59"></a>if (&amp;z).Cmp(a.z) != 0 {
        <a id="L60"></a>t.Errorf(&#34;%s%+v\n\tgot z = %v; want %v&#34;, msg, a, &amp;z, a.z)
    <a id="L61"></a>}
<a id="L62"></a>}


<a id="L65"></a>func TestSumZZ(t *testing.T) {
    <a id="L66"></a>AddZZ := func(z, x, y *Int) *Int { return z.Add(x, y) };
    <a id="L67"></a>SubZZ := func(z, x, y *Int) *Int { return z.Sub(x, y) };
    <a id="L68"></a>for _, a := range sumZZ {
        <a id="L69"></a>arg := a;
        <a id="L70"></a>testFunZZ(t, &#34;AddZZ&#34;, AddZZ, arg);

        <a id="L72"></a>arg = argZZ{a.z, a.y, a.x};
        <a id="L73"></a>testFunZZ(t, &#34;AddZZ symmetric&#34;, AddZZ, arg);

        <a id="L75"></a>arg = argZZ{a.x, a.z, a.y};
        <a id="L76"></a>testFunZZ(t, &#34;SubZZ&#34;, SubZZ, arg);

        <a id="L78"></a>arg = argZZ{a.y, a.z, a.x};
        <a id="L79"></a>testFunZZ(t, &#34;SubZZ symmetric&#34;, SubZZ, arg);
    <a id="L80"></a>}
<a id="L81"></a>}


<a id="L84"></a>func TestProdZZ(t *testing.T) {
    <a id="L85"></a>MulZZ := func(z, x, y *Int) *Int { return z.Mul(x, y) };
    <a id="L86"></a>for _, a := range prodZZ {
        <a id="L87"></a>arg := a;
        <a id="L88"></a>testFunZZ(t, &#34;MulZZ&#34;, MulZZ, arg);

        <a id="L90"></a>arg = argZZ{a.z, a.y, a.x};
        <a id="L91"></a>testFunZZ(t, &#34;MulZZ symmetric&#34;, MulZZ, arg);
    <a id="L92"></a>}
<a id="L93"></a>}


<a id="L96"></a>var facts = map[int]string{
    <a id="L97"></a>0: &#34;1&#34;,
    <a id="L98"></a>1: &#34;1&#34;,
    <a id="L99"></a>2: &#34;2&#34;,
    <a id="L100"></a>10: &#34;3628800&#34;,
    <a id="L101"></a>20: &#34;2432902008176640000&#34;,
    <a id="L102"></a>100: &#34;933262154439441526816992388562667004907159682643816214685929&#34;
        <a id="L103"></a>&#34;638952175999932299156089414639761565182862536979208272237582&#34;
        <a id="L104"></a>&#34;51185210916864000000000000000000000000&#34;,
<a id="L105"></a>}


<a id="L108"></a>func fact(n int) *Int {
    <a id="L109"></a>var z Int;
    <a id="L110"></a>z.New(1);
    <a id="L111"></a>for i := 2; i &lt;= n; i++ {
        <a id="L112"></a>var t Int;
        <a id="L113"></a>t.New(int64(i));
        <a id="L114"></a>z.Mul(&amp;z, &amp;t);
    <a id="L115"></a>}
    <a id="L116"></a>return &amp;z;
<a id="L117"></a>}


<a id="L120"></a>func TestFact(t *testing.T) {
    <a id="L121"></a>for n, s := range facts {
        <a id="L122"></a>f := fact(n).String();
        <a id="L123"></a>if f != s {
            <a id="L124"></a>t.Errorf(&#34;%d! = %s; want %s&#34;, n, f, s)
        <a id="L125"></a>}
    <a id="L126"></a>}
<a id="L127"></a>}


<a id="L130"></a>type fromStringTest struct {
    <a id="L131"></a>in   string;
    <a id="L132"></a>base int;
    <a id="L133"></a>out  int64;
    <a id="L134"></a>ok   bool;
<a id="L135"></a>}


<a id="L138"></a>var fromStringTests = []fromStringTest{
    <a id="L139"></a>fromStringTest{in: &#34;&#34;, ok: false},
    <a id="L140"></a>fromStringTest{in: &#34;a&#34;, ok: false},
    <a id="L141"></a>fromStringTest{in: &#34;z&#34;, ok: false},
    <a id="L142"></a>fromStringTest{in: &#34;+&#34;, ok: false},
    <a id="L143"></a>fromStringTest{&#34;0&#34;, 0, 0, true},
    <a id="L144"></a>fromStringTest{&#34;0&#34;, 10, 0, true},
    <a id="L145"></a>fromStringTest{&#34;0&#34;, 16, 0, true},
    <a id="L146"></a>fromStringTest{&#34;10&#34;, 0, 10, true},
    <a id="L147"></a>fromStringTest{&#34;10&#34;, 10, 10, true},
    <a id="L148"></a>fromStringTest{&#34;10&#34;, 16, 16, true},
    <a id="L149"></a>fromStringTest{&#34;-10&#34;, 16, -16, true},
    <a id="L150"></a>fromStringTest{in: &#34;0x&#34;, ok: false},
    <a id="L151"></a>fromStringTest{&#34;0x10&#34;, 0, 16, true},
    <a id="L152"></a>fromStringTest{in: &#34;0x10&#34;, base: 16, ok: false},
    <a id="L153"></a>fromStringTest{&#34;-0x10&#34;, 0, -16, true},
<a id="L154"></a>}


<a id="L157"></a>func TestSetString(t *testing.T) {
    <a id="L158"></a>for i, test := range fromStringTests {
        <a id="L159"></a>n, ok := new(Int).SetString(test.in, test.base);
        <a id="L160"></a>if ok != test.ok {
            <a id="L161"></a>t.Errorf(&#34;#%d (input &#39;%s&#39;) ok incorrect (should be %t)&#34;, i, test.in, test.ok);
            <a id="L162"></a>continue;
        <a id="L163"></a>}
        <a id="L164"></a>if !ok {
            <a id="L165"></a>continue
        <a id="L166"></a>}

        <a id="L168"></a>if n.Cmp(new(Int).New(test.out)) != 0 {
            <a id="L169"></a>t.Errorf(&#34;#%d (input &#39;%s&#39;) got: %s want: %d\n&#34;, i, test.in, n, test.out)
        <a id="L170"></a>}
    <a id="L171"></a>}
<a id="L172"></a>}


<a id="L175"></a>type divSignsTest struct {
    <a id="L176"></a>x, y int64;
    <a id="L177"></a>q, r int64;
<a id="L178"></a>}


<a id="L181"></a><span class="comment">// These examples taken from the Go Language Spec, section &#34;Arithmetic operators&#34;</span>
<a id="L182"></a>var divSignsTests = []divSignsTest{
    <a id="L183"></a>divSignsTest{5, 3, 1, 2},
    <a id="L184"></a>divSignsTest{-5, 3, -1, -2},
    <a id="L185"></a>divSignsTest{5, -3, -1, 2},
    <a id="L186"></a>divSignsTest{-5, -3, 1, -2},
    <a id="L187"></a>divSignsTest{1, 2, 0, 1},
<a id="L188"></a>}


<a id="L191"></a>func TestDivSigns(t *testing.T) {
    <a id="L192"></a>for i, test := range divSignsTests {
        <a id="L193"></a>x := new(Int).New(test.x);
        <a id="L194"></a>y := new(Int).New(test.y);
        <a id="L195"></a>q, r := new(Int).Div(x, y);
        <a id="L196"></a>expectedQ := new(Int).New(test.q);
        <a id="L197"></a>expectedR := new(Int).New(test.r);

        <a id="L199"></a>if q.Cmp(expectedQ) != 0 || r.Cmp(expectedR) != 0 {
            <a id="L200"></a>t.Errorf(&#34;#%d: got (%s, %s) want (%s, %s)&#34;, i, q, r, expectedQ, expectedR)
        <a id="L201"></a>}
    <a id="L202"></a>}
<a id="L203"></a>}


<a id="L206"></a>func checkSetBytes(b []byte) bool {
    <a id="L207"></a>hex1 := hex.EncodeToString(new(Int).SetBytes(b).Bytes());
    <a id="L208"></a>hex2 := hex.EncodeToString(b);

    <a id="L210"></a>for len(hex1) &lt; len(hex2) {
        <a id="L211"></a>hex1 = &#34;0&#34; + hex1
    <a id="L212"></a>}

    <a id="L214"></a>for len(hex1) &gt; len(hex2) {
        <a id="L215"></a>hex2 = &#34;0&#34; + hex2
    <a id="L216"></a>}

    <a id="L218"></a>return hex1 == hex2;
<a id="L219"></a>}


<a id="L222"></a>func TestSetBytes(t *testing.T) {
    <a id="L223"></a>err := quick.Check(checkSetBytes, nil);
    <a id="L224"></a>if err != nil {
        <a id="L225"></a>t.Error(err)
    <a id="L226"></a>}
<a id="L227"></a>}


<a id="L230"></a>func checkBytes(b []byte) bool {
    <a id="L231"></a>b2 := new(Int).SetBytes(b).Bytes();
    <a id="L232"></a>return bytes.Compare(b, b2) == 0;
<a id="L233"></a>}


<a id="L236"></a>func TestBytes(t *testing.T) {
    <a id="L237"></a>err := quick.Check(checkSetBytes, nil);
    <a id="L238"></a>if err != nil {
        <a id="L239"></a>t.Error(err)
    <a id="L240"></a>}
<a id="L241"></a>}


<a id="L244"></a>func checkDiv(x, y []byte) bool {
    <a id="L245"></a>u := new(Int).SetBytes(x);
    <a id="L246"></a>v := new(Int).SetBytes(y);

    <a id="L248"></a>if len(v.abs) == 0 {
        <a id="L249"></a>return true
    <a id="L250"></a>}

    <a id="L252"></a>q, r := new(Int).Div(u, v);

    <a id="L254"></a>if r.Cmp(v) &gt;= 0 {
        <a id="L255"></a>return false
    <a id="L256"></a>}

    <a id="L258"></a>uprime := new(Int).Set(q);
    <a id="L259"></a>uprime.Mul(uprime, v);
    <a id="L260"></a>uprime.Add(uprime, r);

    <a id="L262"></a>return uprime.Cmp(u) == 0;
<a id="L263"></a>}


<a id="L266"></a>type divTest struct {
    <a id="L267"></a>x, y string;
    <a id="L268"></a>q, r string;
<a id="L269"></a>}


<a id="L272"></a>var divTests = []divTest{
    <a id="L273"></a>divTest{
        <a id="L274"></a>&#34;476217953993950760840509444250624797097991362735329973741718102894495832294430498335824897858659711275234906400899559094370964723884706254265559534144986498357&#34;,
        <a id="L275"></a>&#34;9353930466774385905609975137998169297361893554149986716853295022578535724979483772383667534691121982974895531435241089241440253066816724367338287092081996&#34;,
        <a id="L276"></a>&#34;50911&#34;,
        <a id="L277"></a>&#34;1&#34;,
    <a id="L278"></a>},
    <a id="L279"></a>divTest{
        <a id="L280"></a>&#34;11510768301994997771168&#34;,
        <a id="L281"></a>&#34;1328165573307167369775&#34;,
        <a id="L282"></a>&#34;8&#34;,
        <a id="L283"></a>&#34;885443715537658812968&#34;,
    <a id="L284"></a>},
<a id="L285"></a>}


<a id="L288"></a>func TestDiv(t *testing.T) {
    <a id="L289"></a>err := quick.Check(checkDiv, nil);
    <a id="L290"></a>if err != nil {
        <a id="L291"></a>t.Error(err)
    <a id="L292"></a>}

    <a id="L294"></a>for i, test := range divTests {
        <a id="L295"></a>x, _ := new(Int).SetString(test.x, 10);
        <a id="L296"></a>y, _ := new(Int).SetString(test.y, 10);
        <a id="L297"></a>expectedQ, _ := new(Int).SetString(test.q, 10);
        <a id="L298"></a>expectedR, _ := new(Int).SetString(test.r, 10);

        <a id="L300"></a>q, r := new(Int).Div(x, y);

        <a id="L302"></a>if q.Cmp(expectedQ) != 0 || r.Cmp(expectedR) != 0 {
            <a id="L303"></a>t.Errorf(&#34;#%d got (%s, %s) want (%s, %s)&#34;, i, q, r, expectedQ, expectedR)
        <a id="L304"></a>}
    <a id="L305"></a>}
<a id="L306"></a>}


<a id="L309"></a>func TestDivStepD6(t *testing.T) {
    <a id="L310"></a><span class="comment">// See Knuth, Volume 2, section 4.3.1, exercise 21. This code exercises</span>
    <a id="L311"></a><span class="comment">// a code path which only triggers 1 in 10^{-19} cases.</span>

    <a id="L313"></a>u := &amp;Int{false, []Word{0, 0, 1 + 1&lt;&lt;(_W-1), _M ^ (1 &lt;&lt; (_W - 1))}};
    <a id="L314"></a>v := &amp;Int{false, []Word{5, 2 + 1&lt;&lt;(_W-1), 1 &lt;&lt; (_W - 1)}};

    <a id="L316"></a>q, r := new(Int).Div(u, v);
    <a id="L317"></a>const expectedQ64 = &#34;18446744073709551613&#34;;
    <a id="L318"></a>const expectedR64 = &#34;3138550867693340382088035895064302439801311770021610913807&#34;;
    <a id="L319"></a>const expectedQ32 = &#34;4294967293&#34;;
    <a id="L320"></a>const expectedR32 = &#34;39614081266355540837921718287&#34;;
    <a id="L321"></a>if q.String() != expectedQ64 &amp;&amp; q.String() != expectedQ32 ||
        <a id="L322"></a>r.String() != expectedR64 &amp;&amp; r.String() != expectedR32 {
        <a id="L323"></a>t.Errorf(&#34;got (%s, %s) want (%s, %s) or (%s, %s)&#34;, q, r, expectedQ64, expectedR64, expectedQ32, expectedR32)
    <a id="L324"></a>}
<a id="L325"></a>}


<a id="L328"></a>type lenTest struct {
    <a id="L329"></a>in  string;
    <a id="L330"></a>out int;
<a id="L331"></a>}


<a id="L334"></a>var lenTests = []lenTest{
    <a id="L335"></a>lenTest{&#34;0&#34;, 0},
    <a id="L336"></a>lenTest{&#34;1&#34;, 1},
    <a id="L337"></a>lenTest{&#34;2&#34;, 2},
    <a id="L338"></a>lenTest{&#34;4&#34;, 3},
    <a id="L339"></a>lenTest{&#34;0x8000&#34;, 16},
    <a id="L340"></a>lenTest{&#34;0x80000000&#34;, 32},
    <a id="L341"></a>lenTest{&#34;0x800000000000&#34;, 48},
    <a id="L342"></a>lenTest{&#34;0x8000000000000000&#34;, 64},
    <a id="L343"></a>lenTest{&#34;0x80000000000000000000&#34;, 80},
<a id="L344"></a>}


<a id="L347"></a>func TestLen(t *testing.T) {
    <a id="L348"></a>for i, test := range lenTests {
        <a id="L349"></a>n, ok := new(Int).SetString(test.in, 0);
        <a id="L350"></a>if !ok {
            <a id="L351"></a>t.Errorf(&#34;#%d test input invalid: %s&#34;, i, test.in);
            <a id="L352"></a>continue;
        <a id="L353"></a>}

        <a id="L355"></a>if n.Len() != test.out {
            <a id="L356"></a>t.Errorf(&#34;#%d got %d want %d\n&#34;, i, n.Len(), test.out)
        <a id="L357"></a>}
    <a id="L358"></a>}
<a id="L359"></a>}


<a id="L362"></a>type expTest struct {
    <a id="L363"></a>x, y, m string;
    <a id="L364"></a>out     string;
<a id="L365"></a>}


<a id="L368"></a>var expTests = []expTest{
    <a id="L369"></a>expTest{&#34;5&#34;, &#34;0&#34;, &#34;&#34;, &#34;1&#34;},
    <a id="L370"></a>expTest{&#34;-5&#34;, &#34;0&#34;, &#34;&#34;, &#34;-1&#34;},
    <a id="L371"></a>expTest{&#34;5&#34;, &#34;1&#34;, &#34;&#34;, &#34;5&#34;},
    <a id="L372"></a>expTest{&#34;-5&#34;, &#34;1&#34;, &#34;&#34;, &#34;-5&#34;},
    <a id="L373"></a>expTest{&#34;5&#34;, &#34;2&#34;, &#34;&#34;, &#34;25&#34;},
    <a id="L374"></a>expTest{&#34;1&#34;, &#34;65537&#34;, &#34;2&#34;, &#34;1&#34;},
    <a id="L375"></a>expTest{&#34;0x8000000000000000&#34;, &#34;2&#34;, &#34;&#34;, &#34;0x40000000000000000000000000000000&#34;},
    <a id="L376"></a>expTest{&#34;0x8000000000000000&#34;, &#34;2&#34;, &#34;6719&#34;, &#34;4944&#34;},
    <a id="L377"></a>expTest{&#34;0x8000000000000000&#34;, &#34;3&#34;, &#34;6719&#34;, &#34;5447&#34;},
    <a id="L378"></a>expTest{&#34;0x8000000000000000&#34;, &#34;1000&#34;, &#34;6719&#34;, &#34;1603&#34;},
    <a id="L379"></a>expTest{&#34;0x8000000000000000&#34;, &#34;1000000&#34;, &#34;6719&#34;, &#34;3199&#34;},
    <a id="L380"></a>expTest{
        <a id="L381"></a>&#34;2938462938472983472983659726349017249287491026512746239764525612965293865296239471239874193284792387498274256129746192347&#34;,
        <a id="L382"></a>&#34;298472983472983471903246121093472394872319615612417471234712061&#34;,
        <a id="L383"></a>&#34;29834729834729834729347290846729561262544958723956495615629569234729836259263598127342374289365912465901365498236492183464&#34;,
        <a id="L384"></a>&#34;23537740700184054162508175125554701713153216681790245129157191391322321508055833908509185839069455749219131480588829346291&#34;,
    <a id="L385"></a>},
<a id="L386"></a>}


<a id="L389"></a>func TestExp(t *testing.T) {
    <a id="L390"></a>for i, test := range expTests {
        <a id="L391"></a>x, ok1 := new(Int).SetString(test.x, 0);
        <a id="L392"></a>y, ok2 := new(Int).SetString(test.y, 0);
        <a id="L393"></a>out, ok3 := new(Int).SetString(test.out, 0);

        <a id="L395"></a>var ok4 bool;
        <a id="L396"></a>var m *Int;

        <a id="L398"></a>if len(test.m) == 0 {
            <a id="L399"></a>m, ok4 = nil, true
        <a id="L400"></a>} else {
            <a id="L401"></a>m, ok4 = new(Int).SetString(test.m, 0)
        <a id="L402"></a>}

        <a id="L404"></a>if !ok1 || !ok2 || !ok3 || !ok4 {
            <a id="L405"></a>t.Errorf(&#34;#%d error in input&#34;, i);
            <a id="L406"></a>continue;
        <a id="L407"></a>}

        <a id="L409"></a>z := new(Int).Exp(x, y, m);
        <a id="L410"></a>if z.Cmp(out) != 0 {
            <a id="L411"></a>t.Errorf(&#34;#%d got %s want %s&#34;, i, z, out)
        <a id="L412"></a>}
    <a id="L413"></a>}
<a id="L414"></a>}


<a id="L417"></a>func checkGcd(aBytes, bBytes []byte) bool {
    <a id="L418"></a>a := new(Int).SetBytes(aBytes);
    <a id="L419"></a>b := new(Int).SetBytes(bBytes);

    <a id="L421"></a>x := new(Int);
    <a id="L422"></a>y := new(Int);
    <a id="L423"></a>d := new(Int);

    <a id="L425"></a>GcdInt(d, x, y, a, b);
    <a id="L426"></a>x.Mul(x, a);
    <a id="L427"></a>y.Mul(y, b);
    <a id="L428"></a>x.Add(x, y);

    <a id="L430"></a>return x.Cmp(d) == 0;
<a id="L431"></a>}


<a id="L434"></a>type gcdTest struct {
    <a id="L435"></a>a, b    int64;
    <a id="L436"></a>d, x, y int64;
<a id="L437"></a>}


<a id="L440"></a>var gcdTests = []gcdTest{
    <a id="L441"></a>gcdTest{120, 23, 1, -9, 47},
<a id="L442"></a>}


<a id="L445"></a>func TestGcd(t *testing.T) {
    <a id="L446"></a>for i, test := range gcdTests {
        <a id="L447"></a>a := new(Int).New(test.a);
        <a id="L448"></a>b := new(Int).New(test.b);

        <a id="L450"></a>x := new(Int);
        <a id="L451"></a>y := new(Int);
        <a id="L452"></a>d := new(Int);

        <a id="L454"></a>expectedX := new(Int).New(test.x);
        <a id="L455"></a>expectedY := new(Int).New(test.y);
        <a id="L456"></a>expectedD := new(Int).New(test.d);

        <a id="L458"></a>GcdInt(d, x, y, a, b);

        <a id="L460"></a>if expectedX.Cmp(x) != 0 ||
            <a id="L461"></a>expectedY.Cmp(y) != 0 ||
            <a id="L462"></a>expectedD.Cmp(d) != 0 {
            <a id="L463"></a>t.Errorf(&#34;#%d got (%s %s %s) want (%s %s %s)&#34;, i, x, y, d, expectedX, expectedY, expectedD)
        <a id="L464"></a>}
    <a id="L465"></a>}

    <a id="L467"></a>quick.Check(checkGcd, nil);
<a id="L468"></a>}


<a id="L471"></a>var primes = []string{
    <a id="L472"></a>&#34;2&#34;,
    <a id="L473"></a>&#34;3&#34;,
    <a id="L474"></a>&#34;5&#34;,
    <a id="L475"></a>&#34;7&#34;,
    <a id="L476"></a>&#34;11&#34;,
    <a id="L477"></a>&#34;98920366548084643601728869055592650835572950932266967461790948584315647051443&#34;,
    <a id="L478"></a>&#34;94560208308847015747498523884063394671606671904944666360068158221458669711639&#34;,
    <a id="L479"></a><span class="comment">// http://primes.utm.edu/lists/small/small3.html</span>
    <a id="L480"></a>&#34;449417999055441493994709297093108513015373787049558499205492347871729927573118262811508386655998299074566974373711472560655026288668094291699357843464363003144674940345912431129144354948751003607115263071543163&#34;,
    <a id="L481"></a>&#34;230975859993204150666423538988557839555560243929065415434980904258310530753006723857139742334640122533598517597674807096648905501653461687601339782814316124971547968912893214002992086353183070342498989426570593&#34;,
    <a id="L482"></a>&#34;5521712099665906221540423207019333379125265462121169655563495403888449493493629943498064604536961775110765377745550377067893607246020694972959780839151452457728855382113555867743022746090187341871655890805971735385789993&#34;,
    <a id="L483"></a>&#34;203956878356401977405765866929034577280193993314348263094772646453283062722701277632936616063144088173312372882677123879538709400158306567338328279154499698366071906766440037074217117805690872792848149112022286332144876183376326512083574821647933992961249917319836219304274280243803104015000563790123&#34;,
<a id="L484"></a>}


<a id="L487"></a>var composites = []string{
    <a id="L488"></a>&#34;21284175091214687912771199898307297748211672914763848041968395774954376176754&#34;,
    <a id="L489"></a>&#34;6084766654921918907427900243509372380954290099172559290432744450051395395951&#34;,
    <a id="L490"></a>&#34;84594350493221918389213352992032324280367711247940675652888030554255915464401&#34;,
    <a id="L491"></a>&#34;82793403787388584738507275144194252681&#34;,
<a id="L492"></a>}


<a id="L495"></a>func TestProbablyPrime(t *testing.T) {
    <a id="L496"></a>for i, s := range primes {
        <a id="L497"></a>p, _ := new(Int).SetString(s, 10);
        <a id="L498"></a>if !ProbablyPrime(p, 20) {
            <a id="L499"></a>t.Errorf(&#34;#%d prime found to be non-prime&#34;, i)
        <a id="L500"></a>}
    <a id="L501"></a>}

    <a id="L503"></a>for i, s := range composites {
        <a id="L504"></a>c, _ := new(Int).SetString(s, 10);
        <a id="L505"></a>if ProbablyPrime(c, 20) {
            <a id="L506"></a>t.Errorf(&#34;#%d composite found to be prime&#34;, i)
        <a id="L507"></a>}
    <a id="L508"></a>}
<a id="L509"></a>}


<a id="L512"></a>type rshTest struct {
    <a id="L513"></a>in    string;
    <a id="L514"></a>shift int;
    <a id="L515"></a>out   string;
<a id="L516"></a>}


<a id="L519"></a>var rshTests = []rshTest{
    <a id="L520"></a>rshTest{&#34;0&#34;, 0, &#34;0&#34;},
    <a id="L521"></a>rshTest{&#34;0&#34;, 1, &#34;0&#34;},
    <a id="L522"></a>rshTest{&#34;0&#34;, 2, &#34;0&#34;},
    <a id="L523"></a>rshTest{&#34;1&#34;, 0, &#34;1&#34;},
    <a id="L524"></a>rshTest{&#34;1&#34;, 1, &#34;0&#34;},
    <a id="L525"></a>rshTest{&#34;1&#34;, 2, &#34;0&#34;},
    <a id="L526"></a>rshTest{&#34;2&#34;, 0, &#34;2&#34;},
    <a id="L527"></a>rshTest{&#34;2&#34;, 1, &#34;1&#34;},
    <a id="L528"></a>rshTest{&#34;2&#34;, 2, &#34;0&#34;},
    <a id="L529"></a>rshTest{&#34;4294967296&#34;, 0, &#34;4294967296&#34;},
    <a id="L530"></a>rshTest{&#34;4294967296&#34;, 1, &#34;2147483648&#34;},
    <a id="L531"></a>rshTest{&#34;4294967296&#34;, 2, &#34;1073741824&#34;},
    <a id="L532"></a>rshTest{&#34;18446744073709551616&#34;, 0, &#34;18446744073709551616&#34;},
    <a id="L533"></a>rshTest{&#34;18446744073709551616&#34;, 1, &#34;9223372036854775808&#34;},
    <a id="L534"></a>rshTest{&#34;18446744073709551616&#34;, 2, &#34;4611686018427387904&#34;},
    <a id="L535"></a>rshTest{&#34;18446744073709551616&#34;, 64, &#34;1&#34;},
    <a id="L536"></a>rshTest{&#34;340282366920938463463374607431768211456&#34;, 64, &#34;18446744073709551616&#34;},
    <a id="L537"></a>rshTest{&#34;340282366920938463463374607431768211456&#34;, 128, &#34;1&#34;},
<a id="L538"></a>}


<a id="L541"></a>func TestRsh(t *testing.T) {
    <a id="L542"></a>for i, test := range rshTests {
        <a id="L543"></a>in, _ := new(Int).SetString(test.in, 10);
        <a id="L544"></a>expected, _ := new(Int).SetString(test.out, 10);
        <a id="L545"></a>out := new(Int).Rsh(in, test.shift);

        <a id="L547"></a>if out.Cmp(expected) != 0 {
            <a id="L548"></a>t.Errorf(&#34;#%d got %s want %s&#34;, i, out, expected)
        <a id="L549"></a>}
    <a id="L550"></a>}
<a id="L551"></a>}
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
