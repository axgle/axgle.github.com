<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/rand/rand.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/rand/rand.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package rand implements pseudo-random number generators.</span>
<a id="L6"></a>package rand

<a id="L8"></a><span class="comment">// A Source represents a source of uniformly-distributed</span>
<a id="L9"></a><span class="comment">// pseudo-random int64 values in the range [0, 1&lt;&lt;63).</span>
<a id="L10"></a>type Source interface {
    <a id="L11"></a>Int63() int64;
    <a id="L12"></a>Seed(seed int64);
<a id="L13"></a>}

<a id="L15"></a><span class="comment">// NewSource returns a new pseudo-random Source seeded with the given value.</span>
<a id="L16"></a>func NewSource(seed int64) Source {
    <a id="L17"></a>var rng rngSource;
    <a id="L18"></a>rng.Seed(seed);
    <a id="L19"></a>return &amp;rng;
<a id="L20"></a>}

<a id="L22"></a><span class="comment">// A Rand is a source of random numbers.</span>
<a id="L23"></a>type Rand struct {
    <a id="L24"></a>src Source;
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// New returns a new Rand that uses random values from src</span>
<a id="L28"></a><span class="comment">// to generate other random values.</span>
<a id="L29"></a>func New(src Source) *Rand { return &amp;Rand{src} }

<a id="L31"></a><span class="comment">// Seed uses the provided seed value to initialize the generator to a deterministic state.</span>
<a id="L32"></a>func (r *Rand) Seed(seed int64) { r.src.Seed(seed) }

<a id="L34"></a><span class="comment">// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.</span>
<a id="L35"></a>func (r *Rand) Int63() int64 { return r.src.Int63() }

<a id="L37"></a><span class="comment">// Uint32 returns a pseudo-random 32-bit value as a uint32.</span>
<a id="L38"></a>func (r *Rand) Uint32() uint32 { return uint32(r.Int63() &gt;&gt; 31) }

<a id="L40"></a><span class="comment">// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.</span>
<a id="L41"></a>func (r *Rand) Int31() int32 { return int32(r.Int63() &gt;&gt; 32) }

<a id="L43"></a><span class="comment">// Int returns a non-negative pseudo-random int.</span>
<a id="L44"></a>func (r *Rand) Int() int {
    <a id="L45"></a>u := uint(r.Int63());
    <a id="L46"></a>return int(u &lt;&lt; 1 &gt;&gt; 1); <span class="comment">// clear sign bit if int == int32</span>
<a id="L47"></a>}

<a id="L49"></a><span class="comment">// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).</span>
<a id="L50"></a>func (r *Rand) Int63n(n int64) int64 {
    <a id="L51"></a>if n &lt;= 0 {
        <a id="L52"></a>return 0
    <a id="L53"></a>}
    <a id="L54"></a>max := int64((1 &lt;&lt; 63) - 1 - (1&lt;&lt;63)%uint64(n));
    <a id="L55"></a>v := r.Int63();
    <a id="L56"></a>for v &gt; max {
        <a id="L57"></a>v = r.Int63()
    <a id="L58"></a>}
    <a id="L59"></a>return v % n;
<a id="L60"></a>}

<a id="L62"></a><span class="comment">// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).</span>
<a id="L63"></a>func (r *Rand) Int31n(n int32) int32 { return int32(r.Int63n(int64(n))) }

<a id="L65"></a><span class="comment">// Intn returns, as an int, a non-negative pseudo-random number in [0,n).</span>
<a id="L66"></a>func (r *Rand) Intn(n int) int { return int(r.Int63n(int64(n))) }

<a id="L68"></a><span class="comment">// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).</span>
<a id="L69"></a>func (r *Rand) Float64() float64 { return float64(r.Int63()) / (1 &lt;&lt; 63) }

<a id="L71"></a><span class="comment">// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).</span>
<a id="L72"></a>func (r *Rand) Float32() float32 { return float32(r.Float64()) }

<a id="L74"></a><span class="comment">// Float returns, as a float, a pseudo-random number in [0.0,1.0).</span>
<a id="L75"></a>func (r *Rand) Float() float { return float(r.Float64()) }

<a id="L77"></a><span class="comment">// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).</span>
<a id="L78"></a>func (r *Rand) Perm(n int) []int {
    <a id="L79"></a>m := make([]int, n);
    <a id="L80"></a>for i := 0; i &lt; n; i++ {
        <a id="L81"></a>m[i] = i
    <a id="L82"></a>}
    <a id="L83"></a>for i := 0; i &lt; n; i++ {
        <a id="L84"></a>j := r.Intn(i + 1);
        <a id="L85"></a>m[i], m[j] = m[j], m[i];
    <a id="L86"></a>}
    <a id="L87"></a>return m;
<a id="L88"></a>}

<a id="L90"></a><span class="comment">/*</span>
<a id="L91"></a><span class="comment"> * Top-level convenience functions</span>
<a id="L92"></a><span class="comment"> */</span>

<a id="L94"></a>var globalRand = New(NewSource(1))

<a id="L96"></a><span class="comment">// Seed uses the provided seed value to initialize the generator to a deterministic state.</span>
<a id="L97"></a>func Seed(seed int64) { globalRand.Seed(seed) }

<a id="L99"></a><span class="comment">// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.</span>
<a id="L100"></a>func Int63() int64 { return globalRand.Int63() }

<a id="L102"></a><span class="comment">// Uint32 returns a pseudo-random 32-bit value as a uint32.</span>
<a id="L103"></a>func Uint32() uint32 { return globalRand.Uint32() }

<a id="L105"></a><span class="comment">// Int31 returns a non-negative pseudo-random 31-bit integer as an int32.</span>
<a id="L106"></a>func Int31() int32 { return globalRand.Int31() }

<a id="L108"></a><span class="comment">// Int returns a non-negative pseudo-random int.</span>
<a id="L109"></a>func Int() int { return globalRand.Int() }

<a id="L111"></a><span class="comment">// Int63n returns, as an int64, a non-negative pseudo-random number in [0,n).</span>
<a id="L112"></a>func Int63n(n int64) int64 { return globalRand.Int63n(n) }

<a id="L114"></a><span class="comment">// Int31n returns, as an int32, a non-negative pseudo-random number in [0,n).</span>
<a id="L115"></a>func Int31n(n int32) int32 { return globalRand.Int31n(n) }

<a id="L117"></a><span class="comment">// Intn returns, as an int, a non-negative pseudo-random number in [0,n).</span>
<a id="L118"></a>func Intn(n int) int { return globalRand.Intn(n) }

<a id="L120"></a><span class="comment">// Float64 returns, as a float64, a pseudo-random number in [0.0,1.0).</span>
<a id="L121"></a>func Float64() float64 { return globalRand.Float64() }

<a id="L123"></a><span class="comment">// Float32 returns, as a float32, a pseudo-random number in [0.0,1.0).</span>
<a id="L124"></a>func Float32() float32 { return globalRand.Float32() }

<a id="L126"></a><span class="comment">// Float returns, as a float, a pseudo-random number in [0.0,1.0).</span>
<a id="L127"></a>func Float() float { return globalRand.Float() }

<a id="L129"></a><span class="comment">// Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n).</span>
<a id="L130"></a>func Perm(n int) []int { return globalRand.Perm(n) }

<a id="L132"></a><span class="comment">// NormFloat64 returns a normally distributed float64 in the range</span>
<a id="L133"></a><span class="comment">// [-math.MaxFloat64, +math.MaxFloat64] with</span>
<a id="L134"></a><span class="comment">// standard normal distribution (mean = 0, stddev = 1).</span>
<a id="L135"></a><span class="comment">// To produce a different normal distribution, callers can</span>
<a id="L136"></a><span class="comment">// adjust the output using:</span>
<a id="L137"></a><span class="comment">//</span>
<a id="L138"></a><span class="comment">//  sample = NormFloat64() * desiredStdDev + desiredMean</span>
<a id="L139"></a><span class="comment">//</span>
<a id="L140"></a>func NormFloat64() float64 { return globalRand.NormFloat64() }

<a id="L142"></a><span class="comment">// ExpFloat64 returns an exponentially distributed float64 in the range</span>
<a id="L143"></a><span class="comment">// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter</span>
<a id="L144"></a><span class="comment">// (lambda) is 1 and whose mean is 1/lambda (1).</span>
<a id="L145"></a><span class="comment">// To produce a distribution with a different rate parameter,</span>
<a id="L146"></a><span class="comment">// callers can adjust the output using:</span>
<a id="L147"></a><span class="comment">//</span>
<a id="L148"></a><span class="comment">//  sample = ExpFloat64() / desiredRateParameter</span>
<a id="L149"></a><span class="comment">//</span>
<a id="L150"></a>func ExpFloat64() float64 { return globalRand.ExpFloat64() }
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
