<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/rand/normal.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/rand/normal.go</h1>

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
<a id="L9"></a>)

<a id="L11"></a><span class="comment">/*</span>
<a id="L12"></a><span class="comment"> * Normal distribution</span>
<a id="L13"></a><span class="comment"> *</span>
<a id="L14"></a><span class="comment"> * See &#34;The Ziggurat Method for Generating Random Variables&#34;</span>
<a id="L15"></a><span class="comment"> * (Marsaglia &amp; Tsang, 2000)</span>
<a id="L16"></a><span class="comment"> * http://www.jstatsoft.org/v05/i08/paper [pdf]</span>
<a id="L17"></a><span class="comment"> */</span>

<a id="L19"></a>const (
    <a id="L20"></a>rn = 3.442619855899;
<a id="L21"></a>)

<a id="L23"></a>func absInt32(i int32) uint32 {
    <a id="L24"></a>if i &lt; 0 {
        <a id="L25"></a>return uint32(-i)
    <a id="L26"></a>}
    <a id="L27"></a>return uint32(i);
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// NormFloat64 returns a normally distributed float64 in the range</span>
<a id="L31"></a><span class="comment">// [-math.MaxFloat64, +math.MaxFloat64] with</span>
<a id="L32"></a><span class="comment">// standard normal distribution (mean = 0, stddev = 1).</span>
<a id="L33"></a><span class="comment">// To produce a different normal distribution, callers can</span>
<a id="L34"></a><span class="comment">// adjust the output using:</span>
<a id="L35"></a><span class="comment">//</span>
<a id="L36"></a><span class="comment">//  sample = NormFloat64() * desiredStdDev + desiredMean</span>
<a id="L37"></a><span class="comment">//</span>
<a id="L38"></a>func (r *Rand) NormFloat64() float64 {
    <a id="L39"></a>for {
        <a id="L40"></a>j := int32(r.Uint32()); <span class="comment">// Possibly negative</span>
        <a id="L41"></a>i := j &amp; 0x7F;
        <a id="L42"></a>x := float64(j) * float64(wn[i]);
        <a id="L43"></a>if absInt32(j) &lt; kn[i] {
            <a id="L44"></a><span class="comment">// This case should be hit better than 99% of the time.</span>
            <a id="L45"></a>return x
        <a id="L46"></a>}

        <a id="L48"></a>if i == 0 {
            <a id="L49"></a><span class="comment">// This extra work is only required for the base strip.</span>
            <a id="L50"></a>for {
                <a id="L51"></a>x = -math.Log(r.Float64()) * (1.0 / rn);
                <a id="L52"></a>y := -math.Log(r.Float64());
                <a id="L53"></a>if y+y &gt;= x*x {
                    <a id="L54"></a>break
                <a id="L55"></a>}
            <a id="L56"></a>}
            <a id="L57"></a>if j &gt; 0 {
                <a id="L58"></a>return rn + x
            <a id="L59"></a>}
            <a id="L60"></a>return -rn - x;
        <a id="L61"></a>}
        <a id="L62"></a>if fn[i]+float32(r.Float64())*(fn[i-1]-fn[i]) &lt; float32(math.Exp(-.5*x*x)) {
            <a id="L63"></a>return x
        <a id="L64"></a>}
    <a id="L65"></a>}
    <a id="L66"></a>panic(&#34;unreachable&#34;);
<a id="L67"></a>}

<a id="L69"></a>var kn = [128]uint32{
    <a id="L70"></a>0x76ad2212, 0x0, 0x600f1b53, 0x6ce447a6, 0x725b46a2,
    <a id="L71"></a>0x7560051d, 0x774921eb, 0x789a25bd, 0x799045c3, 0x7a4bce5d,
    <a id="L72"></a>0x7adf629f, 0x7b5682a6, 0x7bb8a8c6, 0x7c0ae722, 0x7c50cce7,
    <a id="L73"></a>0x7c8cec5b, 0x7cc12cd6, 0x7ceefed2, 0x7d177e0b, 0x7d3b8883,
    <a id="L74"></a>0x7d5bce6c, 0x7d78dd64, 0x7d932886, 0x7dab0e57, 0x7dc0dd30,
    <a id="L75"></a>0x7dd4d688, 0x7de73185, 0x7df81cea, 0x7e07c0a3, 0x7e163efa,
    <a id="L76"></a>0x7e23b587, 0x7e303dfd, 0x7e3beec2, 0x7e46db77, 0x7e51155d,
    <a id="L77"></a>0x7e5aabb3, 0x7e63abf7, 0x7e6c222c, 0x7e741906, 0x7e7b9a18,
    <a id="L78"></a>0x7e82adfa, 0x7e895c63, 0x7e8fac4b, 0x7e95a3fb, 0x7e9b4924,
    <a id="L79"></a>0x7ea0a0ef, 0x7ea5b00d, 0x7eaa7ac3, 0x7eaf04f3, 0x7eb3522a,
    <a id="L80"></a>0x7eb765a5, 0x7ebb4259, 0x7ebeeafd, 0x7ec2620a, 0x7ec5a9c4,
    <a id="L81"></a>0x7ec8c441, 0x7ecbb365, 0x7ece78ed, 0x7ed11671, 0x7ed38d62,
    <a id="L82"></a>0x7ed5df12, 0x7ed80cb4, 0x7eda175c, 0x7edc0005, 0x7eddc78e,
    <a id="L83"></a>0x7edf6ebf, 0x7ee0f647, 0x7ee25ebe, 0x7ee3a8a9, 0x7ee4d473,
    <a id="L84"></a>0x7ee5e276, 0x7ee6d2f5, 0x7ee7a620, 0x7ee85c10, 0x7ee8f4cd,
    <a id="L85"></a>0x7ee97047, 0x7ee9ce59, 0x7eea0eca, 0x7eea3147, 0x7eea3568,
    <a id="L86"></a>0x7eea1aab, 0x7ee9e071, 0x7ee98602, 0x7ee90a88, 0x7ee86d08,
    <a id="L87"></a>0x7ee7ac6a, 0x7ee6c769, 0x7ee5bc9c, 0x7ee48a67, 0x7ee32efc,
    <a id="L88"></a>0x7ee1a857, 0x7edff42f, 0x7ede0ffa, 0x7edbf8d9, 0x7ed9ab94,
    <a id="L89"></a>0x7ed7248d, 0x7ed45fae, 0x7ed1585c, 0x7ece095f, 0x7eca6ccb,
    <a id="L90"></a>0x7ec67be2, 0x7ec22eee, 0x7ebd7d1a, 0x7eb85c35, 0x7eb2c075,
    <a id="L91"></a>0x7eac9c20, 0x7ea5df27, 0x7e9e769f, 0x7e964c16, 0x7e8d44ba,
    <a id="L92"></a>0x7e834033, 0x7e781728, 0x7e6b9933, 0x7e5d8a1a, 0x7e4d9ded,
    <a id="L93"></a>0x7e3b737a, 0x7e268c2f, 0x7e0e3ff5, 0x7df1aa5d, 0x7dcf8c72,
    <a id="L94"></a>0x7da61a1e, 0x7d72a0fb, 0x7d30e097, 0x7cd9b4ab, 0x7c600f1a,
    <a id="L95"></a>0x7ba90bdc, 0x7a722176, 0x77d664e5,
<a id="L96"></a>}
<a id="L97"></a>var wn = [128]float32{
    <a id="L98"></a>1.7290405e-09, 1.2680929e-10, 1.6897518e-10, 1.9862688e-10,
    <a id="L99"></a>2.2232431e-10, 2.4244937e-10, 2.601613e-10, 2.7611988e-10,
    <a id="L100"></a>2.9073963e-10, 3.042997e-10, 3.1699796e-10, 3.289802e-10,
    <a id="L101"></a>3.4035738e-10, 3.5121603e-10, 3.616251e-10, 3.7164058e-10,
    <a id="L102"></a>3.8130857e-10, 3.9066758e-10, 3.9975012e-10, 4.08584e-10,
    <a id="L103"></a>4.1719309e-10, 4.2559822e-10, 4.338176e-10, 4.418672e-10,
    <a id="L104"></a>4.497613e-10, 4.5751258e-10, 4.651324e-10, 4.7263105e-10,
    <a id="L105"></a>4.8001775e-10, 4.87301e-10, 4.944885e-10, 5.015873e-10,
    <a id="L106"></a>5.0860405e-10, 5.155446e-10, 5.2241467e-10, 5.2921934e-10,
    <a id="L107"></a>5.359635e-10, 5.426517e-10, 5.4928817e-10, 5.5587696e-10,
    <a id="L108"></a>5.624219e-10, 5.6892646e-10, 5.753941e-10, 5.818282e-10,
    <a id="L109"></a>5.882317e-10, 5.946077e-10, 6.00959e-10, 6.072884e-10,
    <a id="L110"></a>6.135985e-10, 6.19892e-10, 6.2617134e-10, 6.3243905e-10,
    <a id="L111"></a>6.386974e-10, 6.449488e-10, 6.511956e-10, 6.5744005e-10,
    <a id="L112"></a>6.6368433e-10, 6.699307e-10, 6.7618144e-10, 6.824387e-10,
    <a id="L113"></a>6.8870465e-10, 6.949815e-10, 7.012715e-10, 7.075768e-10,
    <a id="L114"></a>7.1389966e-10, 7.202424e-10, 7.266073e-10, 7.329966e-10,
    <a id="L115"></a>7.394128e-10, 7.4585826e-10, 7.5233547e-10, 7.58847e-10,
    <a id="L116"></a>7.653954e-10, 7.719835e-10, 7.7861395e-10, 7.852897e-10,
    <a id="L117"></a>7.920138e-10, 7.987892e-10, 8.0561924e-10, 8.125073e-10,
    <a id="L118"></a>8.194569e-10, 8.2647167e-10, 8.3355556e-10, 8.407127e-10,
    <a id="L119"></a>8.479473e-10, 8.55264e-10, 8.6266755e-10, 8.7016316e-10,
    <a id="L120"></a>8.777562e-10, 8.8545243e-10, 8.932582e-10, 9.0117996e-10,
    <a id="L121"></a>9.09225e-10, 9.174008e-10, 9.2571584e-10, 9.341788e-10,
    <a id="L122"></a>9.427997e-10, 9.515889e-10, 9.605579e-10, 9.697193e-10,
    <a id="L123"></a>9.790869e-10, 9.88676e-10, 9.985036e-10, 1.0085882e-09,
    <a id="L124"></a>1.0189509e-09, 1.0296151e-09, 1.0406069e-09, 1.0519566e-09,
    <a id="L125"></a>1.063698e-09, 1.0758702e-09, 1.0885183e-09, 1.1016947e-09,
    <a id="L126"></a>1.1154611e-09, 1.1298902e-09, 1.1450696e-09, 1.1611052e-09,
    <a id="L127"></a>1.1781276e-09, 1.1962995e-09, 1.2158287e-09, 1.2369856e-09,
    <a id="L128"></a>1.2601323e-09, 1.2857697e-09, 1.3146202e-09, 1.347784e-09,
    <a id="L129"></a>1.3870636e-09, 1.4357403e-09, 1.5008659e-09, 1.6030948e-09,
<a id="L130"></a>}
<a id="L131"></a>var fn = [128]float32{
    <a id="L132"></a>1, 0.9635997, 0.9362827, 0.9130436, 0.89228165, 0.87324303,
    <a id="L133"></a>0.8555006, 0.8387836, 0.8229072, 0.8077383, 0.793177,
    <a id="L134"></a>0.7791461, 0.7655842, 0.7524416, 0.73967725, 0.7272569,
    <a id="L135"></a>0.7151515, 0.7033361, 0.69178915, 0.68049186, 0.6694277,
    <a id="L136"></a>0.658582, 0.6479418, 0.63749546, 0.6272325, 0.6171434,
    <a id="L137"></a>0.6072195, 0.5974532, 0.58783704, 0.5783647, 0.56903,
    <a id="L138"></a>0.5598274, 0.5507518, 0.54179835, 0.5329627, 0.52424055,
    <a id="L139"></a>0.5156282, 0.50712204, 0.49871865, 0.49041483, 0.48220766,
    <a id="L140"></a>0.4740943, 0.46607214, 0.4581387, 0.45029163, 0.44252872,
    <a id="L141"></a>0.43484783, 0.427247, 0.41972435, 0.41227803, 0.40490642,
    <a id="L142"></a>0.39760786, 0.3903808, 0.3832238, 0.37613547, 0.36911446,
    <a id="L143"></a>0.3621595, 0.35526937, 0.34844297, 0.34167916, 0.33497685,
    <a id="L144"></a>0.3283351, 0.3217529, 0.3152294, 0.30876362, 0.30235484,
    <a id="L145"></a>0.29600215, 0.28970486, 0.2834622, 0.2772735, 0.27113807,
    <a id="L146"></a>0.2650553, 0.25902456, 0.2530453, 0.24711695, 0.241239,
    <a id="L147"></a>0.23541094, 0.22963232, 0.2239027, 0.21822165, 0.21258877,
    <a id="L148"></a>0.20700371, 0.20146611, 0.19597565, 0.19053204, 0.18513499,
    <a id="L149"></a>0.17978427, 0.17447963, 0.1692209, 0.16400786, 0.15884037,
    <a id="L150"></a>0.15371831, 0.14864157, 0.14361008, 0.13862377, 0.13368265,
    <a id="L151"></a>0.12878671, 0.12393598, 0.119130544, 0.11437051, 0.10965602,
    <a id="L152"></a>0.104987256, 0.10036444, 0.095787846, 0.0912578, 0.08677467,
    <a id="L153"></a>0.0823389, 0.077950984, 0.073611505, 0.06932112, 0.06508058,
    <a id="L154"></a>0.06089077, 0.056752663, 0.0526674, 0.048636295, 0.044660863,
    <a id="L155"></a>0.040742867, 0.03688439, 0.033087887, 0.029356318,
    <a id="L156"></a>0.025693292, 0.022103304, 0.018592102, 0.015167298,
    <a id="L157"></a>0.011839478, 0.008624485, 0.005548995, 0.0026696292,
<a id="L158"></a>}
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
