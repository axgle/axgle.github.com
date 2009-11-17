<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/rand/rand_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/rand/rand_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package rand

<a id="L7"></a>import (
    <a id="L8"></a>&#34;math&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>const (
    <a id="L15"></a>numTestSamples = 10000;
<a id="L16"></a>)

<a id="L18"></a>type statsResults struct {
    <a id="L19"></a>mean        float64;
    <a id="L20"></a>stddev      float64;
    <a id="L21"></a>closeEnough float64;
    <a id="L22"></a>maxError    float64;
<a id="L23"></a>}

<a id="L25"></a>func max(a, b float64) float64 {
    <a id="L26"></a>if a &gt; b {
        <a id="L27"></a>return a
    <a id="L28"></a>}
    <a id="L29"></a>return b;
<a id="L30"></a>}

<a id="L32"></a>func nearEqual(a, b, closeEnough, maxError float64) bool {
    <a id="L33"></a>absDiff := math.Fabs(a - b);
    <a id="L34"></a>if absDiff &lt; closeEnough { <span class="comment">// Necessary when one value is zero and one value is close to zero.</span>
        <a id="L35"></a>return true
    <a id="L36"></a>}
    <a id="L37"></a>return absDiff/max(math.Fabs(a), math.Fabs(b)) &lt; maxError;
<a id="L38"></a>}

<a id="L40"></a>var testSeeds = []int64{1, 1754801282, 1698661970, 1550503961}

<a id="L42"></a><span class="comment">// checkSimilarDistribution returns success if the mean and stddev of the</span>
<a id="L43"></a><span class="comment">// two statsResults are similar.</span>
<a id="L44"></a>func (this *statsResults) checkSimilarDistribution(expected *statsResults) os.Error {
    <a id="L45"></a>if !nearEqual(this.mean, expected.mean, expected.closeEnough, expected.maxError) {
        <a id="L46"></a>s := fmt.Sprintf(&#34;mean %v != %v (allowed error %v, %v)&#34;, this.mean, expected.mean, expected.closeEnough, expected.maxError);
        <a id="L47"></a>fmt.Println(s);
        <a id="L48"></a>return os.ErrorString(s);
    <a id="L49"></a>}
    <a id="L50"></a>if !nearEqual(this.stddev, expected.stddev, 0, expected.maxError) {
        <a id="L51"></a>s := fmt.Sprintf(&#34;stddev %v != %v (allowed error %v, %v)&#34;, this.stddev, expected.stddev, expected.closeEnough, expected.maxError);
        <a id="L52"></a>fmt.Println(s);
        <a id="L53"></a>return os.ErrorString(s);
    <a id="L54"></a>}
    <a id="L55"></a>return nil;
<a id="L56"></a>}

<a id="L58"></a>func getStatsResults(samples []float64) *statsResults {
    <a id="L59"></a>res := new(statsResults);
    <a id="L60"></a>var sum float64;
    <a id="L61"></a>for i := range samples {
        <a id="L62"></a>sum += samples[i]
    <a id="L63"></a>}
    <a id="L64"></a>res.mean = sum / float64(len(samples));
    <a id="L65"></a>var devsum float64;
    <a id="L66"></a>for i := range samples {
        <a id="L67"></a>devsum += math.Pow(samples[i]-res.mean, 2)
    <a id="L68"></a>}
    <a id="L69"></a>res.stddev = math.Sqrt(devsum / float64(len(samples)));
    <a id="L70"></a>return res;
<a id="L71"></a>}

<a id="L73"></a>func checkSampleDistribution(t *testing.T, samples []float64, expected *statsResults) {
    <a id="L74"></a>actual := getStatsResults(samples);
    <a id="L75"></a>err := actual.checkSimilarDistribution(expected);
    <a id="L76"></a>if err != nil {
        <a id="L77"></a>t.Errorf(err.String())
    <a id="L78"></a>}
<a id="L79"></a>}

<a id="L81"></a>func checkSampleSliceDistributions(t *testing.T, samples []float64, nslices int, expected *statsResults) {
    <a id="L82"></a>chunk := len(samples) / nslices;
    <a id="L83"></a>for i := 0; i &lt; nslices; i++ {
        <a id="L84"></a>low := i * chunk;
        <a id="L85"></a>var high int;
        <a id="L86"></a>if i == nslices-1 {
            <a id="L87"></a>high = len(samples) - 1
        <a id="L88"></a>} else {
            <a id="L89"></a>high = (i + 1) * chunk
        <a id="L90"></a>}
        <a id="L91"></a>checkSampleDistribution(t, samples[low:high], expected);
    <a id="L92"></a>}
<a id="L93"></a>}

<a id="L95"></a><span class="comment">//</span>
<a id="L96"></a><span class="comment">// Normal distribution tests</span>
<a id="L97"></a><span class="comment">//</span>

<a id="L99"></a>func generateNormalSamples(nsamples int, mean, stddev float64, seed int64) []float64 {
    <a id="L100"></a>r := New(NewSource(seed));
    <a id="L101"></a>samples := make([]float64, nsamples);
    <a id="L102"></a>for i := range samples {
        <a id="L103"></a>samples[i] = r.NormFloat64()*stddev + mean
    <a id="L104"></a>}
    <a id="L105"></a>return samples;
<a id="L106"></a>}

<a id="L108"></a>func testNormalDistribution(t *testing.T, nsamples int, mean, stddev float64, seed int64) {
    <a id="L109"></a><span class="comment">//fmt.Printf(&#34;testing nsamples=%v mean=%v stddev=%v seed=%v\n&#34;, nsamples, mean, stddev, seed);</span>

    <a id="L111"></a>samples := generateNormalSamples(nsamples, mean, stddev, seed);
    <a id="L112"></a>errorScale := max(1.0, stddev); <span class="comment">// Error scales with stddev</span>
    <a id="L113"></a>expected := &amp;statsResults{mean, stddev, 0.10 * errorScale, 0.08 * errorScale};

    <a id="L115"></a><span class="comment">// Make sure that the entire set matches the expected distribution.</span>
    <a id="L116"></a>checkSampleDistribution(t, samples, expected);

    <a id="L118"></a><span class="comment">// Make sure that each half of the set matches the expected distribution.</span>
    <a id="L119"></a>checkSampleSliceDistributions(t, samples, 2, expected);

    <a id="L121"></a><span class="comment">// Make sure that each 7th of the set matches the expected distribution.</span>
    <a id="L122"></a>checkSampleSliceDistributions(t, samples, 7, expected);
<a id="L123"></a>}

<a id="L125"></a><span class="comment">// Actual tests</span>

<a id="L127"></a>func TestStandardNormalValues(t *testing.T) {
    <a id="L128"></a>for _, seed := range testSeeds {
        <a id="L129"></a>testNormalDistribution(t, numTestSamples, 0, 1, seed)
    <a id="L130"></a>}
<a id="L131"></a>}

<a id="L133"></a>func TestNonStandardNormalValues(t *testing.T) {
    <a id="L134"></a>for sd := float64(0.5); sd &lt; 1000; sd *= 2 {
        <a id="L135"></a>for m := float64(0.5); m &lt; 1000; m *= 2 {
            <a id="L136"></a>for _, seed := range testSeeds {
                <a id="L137"></a>testNormalDistribution(t, numTestSamples, m, sd, seed)
            <a id="L138"></a>}
        <a id="L139"></a>}
    <a id="L140"></a>}
<a id="L141"></a>}

<a id="L143"></a><span class="comment">//</span>
<a id="L144"></a><span class="comment">// Exponential distribution tests</span>
<a id="L145"></a><span class="comment">//</span>

<a id="L147"></a>func generateExponentialSamples(nsamples int, rate float64, seed int64) []float64 {
    <a id="L148"></a>r := New(NewSource(seed));
    <a id="L149"></a>samples := make([]float64, nsamples);
    <a id="L150"></a>for i := range samples {
        <a id="L151"></a>samples[i] = r.ExpFloat64() / rate
    <a id="L152"></a>}
    <a id="L153"></a>return samples;
<a id="L154"></a>}

<a id="L156"></a>func testExponentialDistribution(t *testing.T, nsamples int, rate float64, seed int64) {
    <a id="L157"></a><span class="comment">//fmt.Printf(&#34;testing nsamples=%v rate=%v seed=%v\n&#34;, nsamples, rate, seed);</span>

    <a id="L159"></a>mean := 1 / rate;
    <a id="L160"></a>stddev := mean;

    <a id="L162"></a>samples := generateExponentialSamples(nsamples, rate, seed);
    <a id="L163"></a>errorScale := max(1.0, 1/rate); <span class="comment">// Error scales with the inverse of the rate</span>
    <a id="L164"></a>expected := &amp;statsResults{mean, stddev, 0.10 * errorScale, 0.20 * errorScale};

    <a id="L166"></a><span class="comment">// Make sure that the entire set matches the expected distribution.</span>
    <a id="L167"></a>checkSampleDistribution(t, samples, expected);

    <a id="L169"></a><span class="comment">// Make sure that each half of the set matches the expected distribution.</span>
    <a id="L170"></a>checkSampleSliceDistributions(t, samples, 2, expected);

    <a id="L172"></a><span class="comment">// Make sure that each 7th of the set matches the expected distribution.</span>
    <a id="L173"></a>checkSampleSliceDistributions(t, samples, 7, expected);
<a id="L174"></a>}

<a id="L176"></a><span class="comment">// Actual tests</span>

<a id="L178"></a>func TestStandardExponentialValues(t *testing.T) {
    <a id="L179"></a>for _, seed := range testSeeds {
        <a id="L180"></a>testExponentialDistribution(t, numTestSamples, 1, seed)
    <a id="L181"></a>}
<a id="L182"></a>}

<a id="L184"></a>func TestNonStandardExponentialValues(t *testing.T) {
    <a id="L185"></a>for rate := float64(0.05); rate &lt; 10; rate *= 2 {
        <a id="L186"></a>for _, seed := range testSeeds {
            <a id="L187"></a>testExponentialDistribution(t, numTestSamples, rate, seed)
        <a id="L188"></a>}
    <a id="L189"></a>}
<a id="L190"></a>}

<a id="L192"></a><span class="comment">//</span>
<a id="L193"></a><span class="comment">// Table generation tests</span>
<a id="L194"></a><span class="comment">//</span>

<a id="L196"></a>func initNorm() (testKn []uint32, testWn, testFn []float32) {
    <a id="L197"></a>const m1 = 1 &lt;&lt; 31;
    <a id="L198"></a>var (
        <a id="L199"></a>dn  float64 = rn;
        <a id="L200"></a>tn          = dn;
        <a id="L201"></a>vn  float64 = 9.91256303526217e-3;
    <a id="L202"></a>)

    <a id="L204"></a>testKn = make([]uint32, 128);
    <a id="L205"></a>testWn = make([]float32, 128);
    <a id="L206"></a>testFn = make([]float32, 128);

    <a id="L208"></a>q := vn / math.Exp(-0.5*dn*dn);
    <a id="L209"></a>testKn[0] = uint32((dn / q) * m1);
    <a id="L210"></a>testKn[1] = 0;
    <a id="L211"></a>testWn[0] = float32(q / m1);
    <a id="L212"></a>testWn[127] = float32(dn / m1);
    <a id="L213"></a>testFn[0] = 1.0;
    <a id="L214"></a>testFn[127] = float32(math.Exp(-0.5 * dn * dn));
    <a id="L215"></a>for i := 126; i &gt;= 1; i-- {
        <a id="L216"></a>dn = math.Sqrt(-2.0 * math.Log(vn/dn+math.Exp(-0.5*dn*dn)));
        <a id="L217"></a>testKn[i+1] = uint32((dn / tn) * m1);
        <a id="L218"></a>tn = dn;
        <a id="L219"></a>testFn[i] = float32(math.Exp(-0.5 * dn * dn));
        <a id="L220"></a>testWn[i] = float32(dn / m1);
    <a id="L221"></a>}
    <a id="L222"></a>return;
<a id="L223"></a>}

<a id="L225"></a>func initExp() (testKe []uint32, testWe, testFe []float32) {
    <a id="L226"></a>const m2 = 1 &lt;&lt; 32;
    <a id="L227"></a>var (
        <a id="L228"></a>de  float64 = re;
        <a id="L229"></a>te          = de;
        <a id="L230"></a>ve  float64 = 3.9496598225815571993e-3;
    <a id="L231"></a>)

    <a id="L233"></a>testKe = make([]uint32, 256);
    <a id="L234"></a>testWe = make([]float32, 256);
    <a id="L235"></a>testFe = make([]float32, 256);

    <a id="L237"></a>q := ve / math.Exp(-de);
    <a id="L238"></a>testKe[0] = uint32((de / q) * m2);
    <a id="L239"></a>testKe[1] = 0;
    <a id="L240"></a>testWe[0] = float32(q / m2);
    <a id="L241"></a>testWe[255] = float32(de / m2);
    <a id="L242"></a>testFe[0] = 1.0;
    <a id="L243"></a>testFe[255] = float32(math.Exp(-de));
    <a id="L244"></a>for i := 254; i &gt;= 1; i-- {
        <a id="L245"></a>de = -math.Log(ve/de + math.Exp(-de));
        <a id="L246"></a>testKe[i+1] = uint32((de / te) * m2);
        <a id="L247"></a>te = de;
        <a id="L248"></a>testFe[i] = float32(math.Exp(-de));
        <a id="L249"></a>testWe[i] = float32(de / m2);
    <a id="L250"></a>}
    <a id="L251"></a>return;
<a id="L252"></a>}

<a id="L254"></a><span class="comment">// compareUint32Slices returns the first index where the two slices</span>
<a id="L255"></a><span class="comment">// disagree, or &lt;0 if the lengths are the same and all elements</span>
<a id="L256"></a><span class="comment">// are identical.</span>
<a id="L257"></a>func compareUint32Slices(s1, s2 []uint32) int {
    <a id="L258"></a>if len(s1) != len(s2) {
        <a id="L259"></a>if len(s1) &gt; len(s2) {
            <a id="L260"></a>return len(s2) + 1
        <a id="L261"></a>}
        <a id="L262"></a>return len(s1) + 1;
    <a id="L263"></a>}
    <a id="L264"></a>for i := range s1 {
        <a id="L265"></a>if s1[i] != s2[i] {
            <a id="L266"></a>return i
        <a id="L267"></a>}
    <a id="L268"></a>}
    <a id="L269"></a>return -1;
<a id="L270"></a>}

<a id="L272"></a><span class="comment">// compareFloat32Slices returns the first index where the two slices</span>
<a id="L273"></a><span class="comment">// disagree, or &lt;0 if the lengths are the same and all elements</span>
<a id="L274"></a><span class="comment">// are identical.</span>
<a id="L275"></a>func compareFloat32Slices(s1, s2 []float32) int {
    <a id="L276"></a>if len(s1) != len(s2) {
        <a id="L277"></a>if len(s1) &gt; len(s2) {
            <a id="L278"></a>return len(s2) + 1
        <a id="L279"></a>}
        <a id="L280"></a>return len(s1) + 1;
    <a id="L281"></a>}
    <a id="L282"></a>for i := range s1 {
        <a id="L283"></a>if !nearEqual(float64(s1[i]), float64(s2[i]), 0, 1e-7) {
            <a id="L284"></a>return i
        <a id="L285"></a>}
    <a id="L286"></a>}
    <a id="L287"></a>return -1;
<a id="L288"></a>}

<a id="L290"></a>func TestNormTables(t *testing.T) {
    <a id="L291"></a>testKn, testWn, testFn := initNorm();
    <a id="L292"></a>if i := compareUint32Slices(kn[0:len(kn)], testKn); i &gt;= 0 {
        <a id="L293"></a>t.Errorf(&#34;kn disagrees at index %v; %v != %v\n&#34;, i, kn[i], testKn[i])
    <a id="L294"></a>}
    <a id="L295"></a>if i := compareFloat32Slices(wn[0:len(wn)], testWn); i &gt;= 0 {
        <a id="L296"></a>t.Errorf(&#34;wn disagrees at index %v; %v != %v\n&#34;, i, wn[i], testWn[i])
    <a id="L297"></a>}
    <a id="L298"></a>if i := compareFloat32Slices(fn[0:len(fn)], testFn); i &gt;= 0 {
        <a id="L299"></a>t.Errorf(&#34;fn disagrees at index %v; %v != %v\n&#34;, i, fn[i], testFn[i])
    <a id="L300"></a>}
<a id="L301"></a>}

<a id="L303"></a>func TestExpTables(t *testing.T) {
    <a id="L304"></a>testKe, testWe, testFe := initExp();
    <a id="L305"></a>if i := compareUint32Slices(ke[0:len(ke)], testKe); i &gt;= 0 {
        <a id="L306"></a>t.Errorf(&#34;ke disagrees at index %v; %v != %v\n&#34;, i, ke[i], testKe[i])
    <a id="L307"></a>}
    <a id="L308"></a>if i := compareFloat32Slices(we[0:len(we)], testWe); i &gt;= 0 {
        <a id="L309"></a>t.Errorf(&#34;we disagrees at index %v; %v != %v\n&#34;, i, we[i], testWe[i])
    <a id="L310"></a>}
    <a id="L311"></a>if i := compareFloat32Slices(fe[0:len(fe)], testFe); i &gt;= 0 {
        <a id="L312"></a>t.Errorf(&#34;fe disagrees at index %v; %v != %v\n&#34;, i, fe[i], testFe[i])
    <a id="L313"></a>}
<a id="L314"></a>}
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
