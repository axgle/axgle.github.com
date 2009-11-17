<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/math/all_test.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/math/all_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package math_test

<a id="L7"></a>import (
    <a id="L8"></a>. &#34;math&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>var vf = []float64{
    <a id="L13"></a>4.9790119248836735e+00,
    <a id="L14"></a>7.7388724745781045e+00,
    <a id="L15"></a>-2.7688005719200159e-01,
    <a id="L16"></a>-5.0106036182710749e+00,
    <a id="L17"></a>9.6362937071984173e+00,
    <a id="L18"></a>2.9263772392439646e+00,
    <a id="L19"></a>5.2290834314593066e+00,
    <a id="L20"></a>2.7279399104360102e+00,
    <a id="L21"></a>1.8253080916808550e+00,
    <a id="L22"></a>-8.6859247685756013e+00,
<a id="L23"></a>}
<a id="L24"></a>var asin = []float64{
    <a id="L25"></a>5.2117697218417440e-01,
    <a id="L26"></a>8.8495619865825236e-01,
    <a id="L27"></a>-2.7691544662819413e-02,
    <a id="L28"></a>-5.2482360935268932e-01,
    <a id="L29"></a>1.3002662421166553e+00,
    <a id="L30"></a>2.9698415875871901e-01,
    <a id="L31"></a>5.5025938468083364e-01,
    <a id="L32"></a>2.7629597861677200e-01,
    <a id="L33"></a>1.8355989225745148e-01,
    <a id="L34"></a>-1.0523547536021498e+00,
<a id="L35"></a>}
<a id="L36"></a>var atan = []float64{
    <a id="L37"></a>1.3725902621296217e+00,
    <a id="L38"></a>1.4422906096452980e+00,
    <a id="L39"></a>-2.7011324359471755e-01,
    <a id="L40"></a>-1.3738077684543379e+00,
    <a id="L41"></a>1.4673921193587666e+00,
    <a id="L42"></a>1.2415173565870167e+00,
    <a id="L43"></a>1.3818396865615167e+00,
    <a id="L44"></a>1.2194305844639670e+00,
    <a id="L45"></a>1.0696031952318783e+00,
    <a id="L46"></a>-1.4561721938838085e+00,
<a id="L47"></a>}
<a id="L48"></a>var exp = []float64{
    <a id="L49"></a>1.4533071302642137e+02,
    <a id="L50"></a>2.2958822575694450e+03,
    <a id="L51"></a>7.5814542574851666e-01,
    <a id="L52"></a>6.6668778421791010e-03,
    <a id="L53"></a>1.5310493273896035e+04,
    <a id="L54"></a>1.8659907517999329e+01,
    <a id="L55"></a>1.8662167355098713e+02,
    <a id="L56"></a>1.5301332413189379e+01,
    <a id="L57"></a>6.2047063430646876e+00,
    <a id="L58"></a>1.6894712385826522e-04,
<a id="L59"></a>}
<a id="L60"></a>var floor = []float64{
    <a id="L61"></a>4.0000000000000000e+00,
    <a id="L62"></a>7.0000000000000000e+00,
    <a id="L63"></a>-1.0000000000000000e+00,
    <a id="L64"></a>-6.0000000000000000e+00,
    <a id="L65"></a>9.0000000000000000e+00,
    <a id="L66"></a>2.0000000000000000e+00,
    <a id="L67"></a>5.0000000000000000e+00,
    <a id="L68"></a>2.0000000000000000e+00,
    <a id="L69"></a>1.0000000000000000e+00,
    <a id="L70"></a>-9.0000000000000000e+00,
<a id="L71"></a>}
<a id="L72"></a>var log = []float64{
    <a id="L73"></a>1.6052314626930630e+00,
    <a id="L74"></a>2.0462560018708768e+00,
    <a id="L75"></a>-1.2841708730962657e+00,
    <a id="L76"></a>1.6115563905281544e+00,
    <a id="L77"></a>2.2655365644872018e+00,
    <a id="L78"></a>1.0737652208918380e+00,
    <a id="L79"></a>1.6542360106073545e+00,
    <a id="L80"></a>1.0035467127723465e+00,
    <a id="L81"></a>6.0174879014578053e-01,
    <a id="L82"></a>2.1617038728473527e+00,
<a id="L83"></a>}
<a id="L84"></a>var pow = []float64{
    <a id="L85"></a>9.5282232631648415e+04,
    <a id="L86"></a>5.4811599352999900e+07,
    <a id="L87"></a>5.2859121715894400e-01,
    <a id="L88"></a>9.7587991957286472e-06,
    <a id="L89"></a>4.3280643293460450e+09,
    <a id="L90"></a>8.4406761805034551e+02,
    <a id="L91"></a>1.6946633276191194e+05,
    <a id="L92"></a>5.3449040147551940e+02,
    <a id="L93"></a>6.6881821384514159e+01,
    <a id="L94"></a>2.0609869004248744e-09,
<a id="L95"></a>}
<a id="L96"></a>var sin = []float64{
    <a id="L97"></a>-9.6466616586009283e-01,
    <a id="L98"></a>9.9338225271646543e-01,
    <a id="L99"></a>-2.7335587039794395e-01,
    <a id="L100"></a>9.5586257685042800e-01,
    <a id="L101"></a>-2.0994210667799692e-01,
    <a id="L102"></a>2.1355787807998605e-01,
    <a id="L103"></a>-8.6945689711673619e-01,
    <a id="L104"></a>4.0195666811555783e-01,
    <a id="L105"></a>9.6778633541688000e-01,
    <a id="L106"></a>-6.7344058690503452e-01,
<a id="L107"></a>}
<a id="L108"></a>var sinh = []float64{
    <a id="L109"></a>7.2661916084208533e+01,
    <a id="L110"></a>1.1479409110035194e+03,
    <a id="L111"></a>-2.8043136512812520e-01,
    <a id="L112"></a>-7.4994290911815868e+01,
    <a id="L113"></a>7.6552466042906761e+03,
    <a id="L114"></a>9.3031583421672010e+00,
    <a id="L115"></a>9.3308157558281088e+01,
    <a id="L116"></a>7.6179893137269143e+00,
    <a id="L117"></a>3.0217691805496156e+00,
    <a id="L118"></a>-2.9595057572444951e+03,
<a id="L119"></a>}
<a id="L120"></a>var sqrt = []float64{
    <a id="L121"></a>2.2313699659365484e+00,
    <a id="L122"></a>2.7818829009464263e+00,
    <a id="L123"></a>5.2619393496314792e-01,
    <a id="L124"></a>2.2384377628763938e+00,
    <a id="L125"></a>3.1042380236055380e+00,
    <a id="L126"></a>1.7106657298385224e+00,
    <a id="L127"></a>2.2867189227054791e+00,
    <a id="L128"></a>1.6516476350711160e+00,
    <a id="L129"></a>1.3510396336454586e+00,
    <a id="L130"></a>2.9471892997524950e+00,
<a id="L131"></a>}
<a id="L132"></a>var tan = []float64{
    <a id="L133"></a>-3.6613165650402277e+00,
    <a id="L134"></a>8.6490023264859754e+00,
    <a id="L135"></a>-2.8417941955033615e-01,
    <a id="L136"></a>3.2532901859747287e+00,
    <a id="L137"></a>2.1472756403802937e-01,
    <a id="L138"></a>-2.1860091071106700e-01,
    <a id="L139"></a>-1.7600028178723679e+00,
    <a id="L140"></a>-4.3898089147528178e-01,
    <a id="L141"></a>-3.8438855602011305e+00,
    <a id="L142"></a>9.1098879337768517e-01,
<a id="L143"></a>}
<a id="L144"></a>var tanh = []float64{
    <a id="L145"></a>9.9990531206936328e-01,
    <a id="L146"></a>9.9999962057085307e-01,
    <a id="L147"></a>-2.7001505097318680e-01,
    <a id="L148"></a>-9.9991110943061700e-01,
    <a id="L149"></a>9.9999999146798441e-01,
    <a id="L150"></a>9.9427249436125233e-01,
    <a id="L151"></a>9.9994257600983156e-01,
    <a id="L152"></a>9.9149409509772863e-01,
    <a id="L153"></a>9.4936501296239700e-01,
    <a id="L154"></a>-9.9999994291374019e-01,
<a id="L155"></a>}

<a id="L157"></a>func tolerance(a, b, e float64) bool {
    <a id="L158"></a>d := a - b;
    <a id="L159"></a>if d &lt; 0 {
        <a id="L160"></a>d = -d
    <a id="L161"></a>}

    <a id="L163"></a>if a != 0 {
        <a id="L164"></a>e = e * a;
        <a id="L165"></a>if e &lt; 0 {
            <a id="L166"></a>e = -e
        <a id="L167"></a>}
    <a id="L168"></a>}
    <a id="L169"></a>return d &lt; e;
<a id="L170"></a>}
<a id="L171"></a>func close(a, b float64) bool     { return tolerance(a, b, 1e-14) }
<a id="L172"></a>func veryclose(a, b float64) bool { return tolerance(a, b, 4e-16) }

<a id="L174"></a>func TestAsin(t *testing.T) {
    <a id="L175"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L176"></a>if f := Asin(vf[i] / 10); !veryclose(asin[i], f) {
            <a id="L177"></a>t.Errorf(&#34;Asin(%g) = %g, want %g\n&#34;, vf[i]/10, f, asin[i])
        <a id="L178"></a>}
    <a id="L179"></a>}
<a id="L180"></a>}

<a id="L182"></a>func TestAtan(t *testing.T) {
    <a id="L183"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L184"></a>if f := Atan(vf[i]); !veryclose(atan[i], f) {
            <a id="L185"></a>t.Errorf(&#34;Atan(%g) = %g, want %g\n&#34;, vf[i], f, atan[i])
        <a id="L186"></a>}
    <a id="L187"></a>}
<a id="L188"></a>}

<a id="L190"></a>func TestExp(t *testing.T) {
    <a id="L191"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L192"></a>if f := Exp(vf[i]); !veryclose(exp[i], f) {
            <a id="L193"></a>t.Errorf(&#34;Exp(%g) = %g, want %g\n&#34;, vf[i], f, exp[i])
        <a id="L194"></a>}
    <a id="L195"></a>}
<a id="L196"></a>}

<a id="L198"></a>func TestFloor(t *testing.T) {
    <a id="L199"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L200"></a>if f := Floor(vf[i]); floor[i] != f {
            <a id="L201"></a>t.Errorf(&#34;Floor(%g) = %g, want %g\n&#34;, vf[i], f, floor[i])
        <a id="L202"></a>}
    <a id="L203"></a>}
<a id="L204"></a>}

<a id="L206"></a>func TestLog(t *testing.T) {
    <a id="L207"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L208"></a>a := Fabs(vf[i]);
        <a id="L209"></a>if f := Log(a); log[i] != f {
            <a id="L210"></a>t.Errorf(&#34;Log(%g) = %g, want %g\n&#34;, a, f, log[i])
        <a id="L211"></a>}
    <a id="L212"></a>}
    <a id="L213"></a>if f := Log(10); f != Ln10 {
        <a id="L214"></a>t.Errorf(&#34;Log(%g) = %g, want %g\n&#34;, 10, f, Ln10)
    <a id="L215"></a>}
<a id="L216"></a>}

<a id="L218"></a>func TestPow(t *testing.T) {
    <a id="L219"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L220"></a>if f := Pow(10, vf[i]); !close(pow[i], f) {
            <a id="L221"></a>t.Errorf(&#34;Pow(10, %.17g) = %.17g, want %.17g\n&#34;, vf[i], f, pow[i])
        <a id="L222"></a>}
    <a id="L223"></a>}
<a id="L224"></a>}

<a id="L226"></a>func TestSin(t *testing.T) {
    <a id="L227"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L228"></a>if f := Sin(vf[i]); !close(sin[i], f) {
            <a id="L229"></a>t.Errorf(&#34;Sin(%g) = %g, want %g\n&#34;, vf[i], f, sin[i])
        <a id="L230"></a>}
    <a id="L231"></a>}
<a id="L232"></a>}

<a id="L234"></a>func TestSinh(t *testing.T) {
    <a id="L235"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L236"></a>if f := Sinh(vf[i]); !veryclose(sinh[i], f) {
            <a id="L237"></a>t.Errorf(&#34;Sinh(%g) = %g, want %g\n&#34;, vf[i], f, sinh[i])
        <a id="L238"></a>}
    <a id="L239"></a>}
<a id="L240"></a>}

<a id="L242"></a>func TestSqrt(t *testing.T) {
    <a id="L243"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L244"></a>a := Fabs(vf[i]);
        <a id="L245"></a>if f := Sqrt(a); !veryclose(sqrt[i], f) {
            <a id="L246"></a>t.Errorf(&#34;Sqrt(%g) = %g, want %g\n&#34;, a, f, floor[i])
        <a id="L247"></a>}
    <a id="L248"></a>}
<a id="L249"></a>}

<a id="L251"></a>func TestTan(t *testing.T) {
    <a id="L252"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L253"></a>if f := Tan(vf[i]); !close(tan[i], f) {
            <a id="L254"></a>t.Errorf(&#34;Tan(%g) = %g, want %g\n&#34;, vf[i], f, tan[i])
        <a id="L255"></a>}
    <a id="L256"></a>}
<a id="L257"></a>}

<a id="L259"></a>func TestTanh(t *testing.T) {
    <a id="L260"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L261"></a>if f := Tanh(vf[i]); !veryclose(tanh[i], f) {
            <a id="L262"></a>t.Errorf(&#34;Tanh(%g) = %g, want %g\n&#34;, vf[i], f, tanh[i])
        <a id="L263"></a>}
    <a id="L264"></a>}
<a id="L265"></a>}

<a id="L267"></a>func TestHypot(t *testing.T) {
    <a id="L268"></a>for i := 0; i &lt; len(vf); i++ {
        <a id="L269"></a>a := Fabs(tanh[i] * Sqrt(2));
        <a id="L270"></a>if f := Hypot(tanh[i], tanh[i]); !veryclose(a, f) {
            <a id="L271"></a>t.Errorf(&#34;Hypot(%g, %g) = %g, want %g\n&#34;, tanh[i], tanh[i], f, a)
        <a id="L272"></a>}
    <a id="L273"></a>}
<a id="L274"></a>}
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
