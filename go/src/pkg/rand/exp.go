<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/rand/exp.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/rand/exp.go</h1>

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
<a id="L12"></a><span class="comment"> * Exponential distribution</span>
<a id="L13"></a><span class="comment"> *</span>
<a id="L14"></a><span class="comment"> * See &#34;The Ziggurat Method for Generating Random Variables&#34;</span>
<a id="L15"></a><span class="comment"> * (Marsaglia &amp; Tsang, 2000)</span>
<a id="L16"></a><span class="comment"> * http://www.jstatsoft.org/v05/i08/paper [pdf]</span>
<a id="L17"></a><span class="comment"> */</span>

<a id="L19"></a>const (
    <a id="L20"></a>re = 7.69711747013104972;
<a id="L21"></a>)

<a id="L23"></a><span class="comment">// ExpFloat64 returns an exponentially distributed float64 in the range</span>
<a id="L24"></a><span class="comment">// (0, +math.MaxFloat64] with an exponential distribution whose rate parameter</span>
<a id="L25"></a><span class="comment">// (lambda) is 1 and whose mean is 1/lambda (1).</span>
<a id="L26"></a><span class="comment">// To produce a distribution with a different rate parameter,</span>
<a id="L27"></a><span class="comment">// callers can adjust the output using:</span>
<a id="L28"></a><span class="comment">//</span>
<a id="L29"></a><span class="comment">//  sample = ExpFloat64() / desiredRateParameter</span>
<a id="L30"></a><span class="comment">//</span>
<a id="L31"></a>func (r *Rand) ExpFloat64() float64 {
    <a id="L32"></a>for {
        <a id="L33"></a>j := r.Uint32();
        <a id="L34"></a>i := j &amp; 0xFF;
        <a id="L35"></a>x := float64(j) * float64(we[i]);
        <a id="L36"></a>if j &lt; ke[i] {
            <a id="L37"></a>return x
        <a id="L38"></a>}
        <a id="L39"></a>if i == 0 {
            <a id="L40"></a>return re - math.Log(r.Float64())
        <a id="L41"></a>}
        <a id="L42"></a>if fe[i]+float32(r.Float64())*(fe[i-1]-fe[i]) &lt; float32(math.Exp(-x)) {
            <a id="L43"></a>return x
        <a id="L44"></a>}
    <a id="L45"></a>}
    <a id="L46"></a>panic(&#34;unreachable&#34;);
<a id="L47"></a>}

<a id="L49"></a>var ke = [256]uint32{
    <a id="L50"></a>0xe290a139, 0x0, 0x9beadebc, 0xc377ac71, 0xd4ddb990,
    <a id="L51"></a>0xde893fb8, 0xe4a8e87c, 0xe8dff16a, 0xebf2deab, 0xee49a6e8,
    <a id="L52"></a>0xf0204efd, 0xf19bdb8e, 0xf2d458bb, 0xf3da104b, 0xf4b86d78,
    <a id="L53"></a>0xf577ad8a, 0xf61de83d, 0xf6afb784, 0xf730a573, 0xf7a37651,
    <a id="L54"></a>0xf80a5bb6, 0xf867189d, 0xf8bb1b4f, 0xf9079062, 0xf94d70ca,
    <a id="L55"></a>0xf98d8c7d, 0xf9c8928a, 0xf9ff175b, 0xfa319996, 0xfa6085f8,
    <a id="L56"></a>0xfa8c3a62, 0xfab5084e, 0xfadb36c8, 0xfaff0410, 0xfb20a6ea,
    <a id="L57"></a>0xfb404fb4, 0xfb5e2951, 0xfb7a59e9, 0xfb95038c, 0xfbae44ba,
    <a id="L58"></a>0xfbc638d8, 0xfbdcf892, 0xfbf29a30, 0xfc0731df, 0xfc1ad1ed,
    <a id="L59"></a>0xfc2d8b02, 0xfc3f6c4d, 0xfc5083ac, 0xfc60ddd1, 0xfc708662,
    <a id="L60"></a>0xfc7f8810, 0xfc8decb4, 0xfc9bbd62, 0xfca9027c, 0xfcb5c3c3,
    <a id="L61"></a>0xfcc20864, 0xfccdd70a, 0xfcd935e3, 0xfce42ab0, 0xfceebace,
    <a id="L62"></a>0xfcf8eb3b, 0xfd02c0a0, 0xfd0c3f59, 0xfd156b7b, 0xfd1e48d6,
    <a id="L63"></a>0xfd26daff, 0xfd2f2552, 0xfd372af7, 0xfd3eeee5, 0xfd4673e7,
    <a id="L64"></a>0xfd4dbc9e, 0xfd54cb85, 0xfd5ba2f2, 0xfd62451b, 0xfd68b415,
    <a id="L65"></a>0xfd6ef1da, 0xfd750047, 0xfd7ae120, 0xfd809612, 0xfd8620b4,
    <a id="L66"></a>0xfd8b8285, 0xfd90bcf5, 0xfd95d15e, 0xfd9ac10b, 0xfd9f8d36,
    <a id="L67"></a>0xfda43708, 0xfda8bf9e, 0xfdad2806, 0xfdb17141, 0xfdb59c46,
    <a id="L68"></a>0xfdb9a9fd, 0xfdbd9b46, 0xfdc170f6, 0xfdc52bd8, 0xfdc8ccac,
    <a id="L69"></a>0xfdcc542d, 0xfdcfc30b, 0xfdd319ef, 0xfdd6597a, 0xfdd98245,
    <a id="L70"></a>0xfddc94e5, 0xfddf91e6, 0xfde279ce, 0xfde54d1f, 0xfde80c52,
    <a id="L71"></a>0xfdeab7de, 0xfded5034, 0xfdefd5be, 0xfdf248e3, 0xfdf4aa06,
    <a id="L72"></a>0xfdf6f984, 0xfdf937b6, 0xfdfb64f4, 0xfdfd818d, 0xfdff8dd0,
    <a id="L73"></a>0xfe018a08, 0xfe03767a, 0xfe05536c, 0xfe07211c, 0xfe08dfc9,
    <a id="L74"></a>0xfe0a8fab, 0xfe0c30fb, 0xfe0dc3ec, 0xfe0f48b1, 0xfe10bf76,
    <a id="L75"></a>0xfe122869, 0xfe1383b4, 0xfe14d17c, 0xfe1611e7, 0xfe174516,
    <a id="L76"></a>0xfe186b2a, 0xfe19843e, 0xfe1a9070, 0xfe1b8fd6, 0xfe1c8289,
    <a id="L77"></a>0xfe1d689b, 0xfe1e4220, 0xfe1f0f26, 0xfe1fcfbc, 0xfe2083ed,
    <a id="L78"></a>0xfe212bc3, 0xfe21c745, 0xfe225678, 0xfe22d95f, 0xfe234ffb,
    <a id="L79"></a>0xfe23ba4a, 0xfe241849, 0xfe2469f2, 0xfe24af3c, 0xfe24e81e,
    <a id="L80"></a>0xfe25148b, 0xfe253474, 0xfe2547c7, 0xfe254e70, 0xfe25485a,
    <a id="L81"></a>0xfe25356a, 0xfe251586, 0xfe24e88f, 0xfe24ae64, 0xfe2466e1,
    <a id="L82"></a>0xfe2411df, 0xfe23af34, 0xfe233eb4, 0xfe22c02c, 0xfe22336b,
    <a id="L83"></a>0xfe219838, 0xfe20ee58, 0xfe20358c, 0xfe1f6d92, 0xfe1e9621,
    <a id="L84"></a>0xfe1daef0, 0xfe1cb7ac, 0xfe1bb002, 0xfe1a9798, 0xfe196e0d,
    <a id="L85"></a>0xfe1832fd, 0xfe16e5fe, 0xfe15869d, 0xfe141464, 0xfe128ed3,
    <a id="L86"></a>0xfe10f565, 0xfe0f478c, 0xfe0d84b1, 0xfe0bac36, 0xfe09bd73,
    <a id="L87"></a>0xfe07b7b5, 0xfe059a40, 0xfe03644c, 0xfe011504, 0xfdfeab88,
    <a id="L88"></a>0xfdfc26e9, 0xfdf98629, 0xfdf6c83b, 0xfdf3ec01, 0xfdf0f04a,
    <a id="L89"></a>0xfdedd3d1, 0xfdea953d, 0xfde7331e, 0xfde3abe9, 0xfddffdfb,
    <a id="L90"></a>0xfddc2791, 0xfdd826cd, 0xfdd3f9a8, 0xfdcf9dfc, 0xfdcb1176,
    <a id="L91"></a>0xfdc65198, 0xfdc15bb3, 0xfdbc2ce2, 0xfdb6c206, 0xfdb117be,
    <a id="L92"></a>0xfdab2a63, 0xfda4f5fd, 0xfd9e7640, 0xfd97a67a, 0xfd908192,
    <a id="L93"></a>0xfd8901f2, 0xfd812182, 0xfd78d98e, 0xfd7022bb, 0xfd66f4ed,
    <a id="L94"></a>0xfd5d4732, 0xfd530f9c, 0xfd48432b, 0xfd3cd59a, 0xfd30b936,
    <a id="L95"></a>0xfd23dea4, 0xfd16349e, 0xfd07a7a3, 0xfcf8219b, 0xfce7895b,
    <a id="L96"></a>0xfcd5c220, 0xfcc2aadb, 0xfcae1d5e, 0xfc97ed4e, 0xfc7fe6d4,
    <a id="L97"></a>0xfc65ccf3, 0xfc495762, 0xfc2a2fc8, 0xfc07ee19, 0xfbe213c1,
    <a id="L98"></a>0xfbb8051a, 0xfb890078, 0xfb5411a5, 0xfb180005, 0xfad33482,
    <a id="L99"></a>0xfa839276, 0xfa263b32, 0xf9b72d1c, 0xf930a1a2, 0xf889f023,
    <a id="L100"></a>0xf7b577d2, 0xf69c650c, 0xf51530f0, 0xf2cb0e3c, 0xeeefb15d,
    <a id="L101"></a>0xe6da6ecf,
<a id="L102"></a>}
<a id="L103"></a>var we = [256]float32{
    <a id="L104"></a>2.0249555e-09, 1.486674e-11, 2.4409617e-11, 3.1968806e-11,
    <a id="L105"></a>3.844677e-11, 4.4228204e-11, 4.9516443e-11, 5.443359e-11,
    <a id="L106"></a>5.905944e-11, 6.344942e-11, 6.7643814e-11, 7.1672945e-11,
    <a id="L107"></a>7.556032e-11, 7.932458e-11, 8.298079e-11, 8.654132e-11,
    <a id="L108"></a>9.0016515e-11, 9.3415074e-11, 9.674443e-11, 1.0001099e-10,
    <a id="L109"></a>1.03220314e-10, 1.06377254e-10, 1.09486115e-10, 1.1255068e-10,
    <a id="L110"></a>1.1557435e-10, 1.1856015e-10, 1.2151083e-10, 1.2442886e-10,
    <a id="L111"></a>1.2731648e-10, 1.3017575e-10, 1.3300853e-10, 1.3581657e-10,
    <a id="L112"></a>1.3860142e-10, 1.4136457e-10, 1.4410738e-10, 1.4683108e-10,
    <a id="L113"></a>1.4953687e-10, 1.5222583e-10, 1.54899e-10, 1.5755733e-10,
    <a id="L114"></a>1.6020171e-10, 1.6283301e-10, 1.6545203e-10, 1.6805951e-10,
    <a id="L115"></a>1.7065617e-10, 1.732427e-10, 1.7581973e-10, 1.7838787e-10,
    <a id="L116"></a>1.8094774e-10, 1.8349985e-10, 1.8604476e-10, 1.8858298e-10,
    <a id="L117"></a>1.9111498e-10, 1.9364126e-10, 1.9616223e-10, 1.9867835e-10,
    <a id="L118"></a>2.0119004e-10, 2.0369768e-10, 2.0620168e-10, 2.087024e-10,
    <a id="L119"></a>2.1120022e-10, 2.136955e-10, 2.1618855e-10, 2.1867974e-10,
    <a id="L120"></a>2.2116936e-10, 2.2365775e-10, 2.261452e-10, 2.2863202e-10,
    <a id="L121"></a>2.311185e-10, 2.3360494e-10, 2.360916e-10, 2.3857874e-10,
    <a id="L122"></a>2.4106667e-10, 2.4355562e-10, 2.4604588e-10, 2.485377e-10,
    <a id="L123"></a>2.5103128e-10, 2.5352695e-10, 2.560249e-10, 2.585254e-10,
    <a id="L124"></a>2.6102867e-10, 2.6353494e-10, 2.6604446e-10, 2.6855745e-10,
    <a id="L125"></a>2.7107416e-10, 2.7359479e-10, 2.761196e-10, 2.7864877e-10,
    <a id="L126"></a>2.8118255e-10, 2.8372119e-10, 2.8626485e-10, 2.888138e-10,
    <a id="L127"></a>2.9136826e-10, 2.939284e-10, 2.9649452e-10, 2.9906677e-10,
    <a id="L128"></a>3.016454e-10, 3.0423064e-10, 3.0682268e-10, 3.0942177e-10,
    <a id="L129"></a>3.1202813e-10, 3.1464195e-10, 3.1726352e-10, 3.19893e-10,
    <a id="L130"></a>3.2253064e-10, 3.251767e-10, 3.2783135e-10, 3.3049485e-10,
    <a id="L131"></a>3.3316744e-10, 3.3584938e-10, 3.3854083e-10, 3.4124212e-10,
    <a id="L132"></a>3.4395342e-10, 3.46675e-10, 3.4940711e-10, 3.5215003e-10,
    <a id="L133"></a>3.5490397e-10, 3.5766917e-10, 3.6044595e-10, 3.6323455e-10,
    <a id="L134"></a>3.660352e-10, 3.6884823e-10, 3.7167386e-10, 3.745124e-10,
    <a id="L135"></a>3.773641e-10, 3.802293e-10, 3.8310827e-10, 3.860013e-10,
    <a id="L136"></a>3.8890866e-10, 3.918307e-10, 3.9476775e-10, 3.9772008e-10,
    <a id="L137"></a>4.0068804e-10, 4.0367196e-10, 4.0667217e-10, 4.09689e-10,
    <a id="L138"></a>4.1272286e-10, 4.1577405e-10, 4.1884296e-10, 4.2192994e-10,
    <a id="L139"></a>4.250354e-10, 4.281597e-10, 4.313033e-10, 4.3446652e-10,
    <a id="L140"></a>4.3764986e-10, 4.408537e-10, 4.4407847e-10, 4.4732465e-10,
    <a id="L141"></a>4.5059267e-10, 4.5388301e-10, 4.571962e-10, 4.6053267e-10,
    <a id="L142"></a>4.6389292e-10, 4.6727755e-10, 4.70687e-10, 4.741219e-10,
    <a id="L143"></a>4.7758275e-10, 4.810702e-10, 4.845848e-10, 4.8812715e-10,
    <a id="L144"></a>4.9169796e-10, 4.9529775e-10, 4.989273e-10, 5.0258725e-10,
    <a id="L145"></a>5.0627835e-10, 5.100013e-10, 5.1375687e-10, 5.1754584e-10,
    <a id="L146"></a>5.21369e-10, 5.2522725e-10, 5.2912136e-10, 5.330522e-10,
    <a id="L147"></a>5.370208e-10, 5.4102806e-10, 5.45075e-10, 5.491625e-10,
    <a id="L148"></a>5.532918e-10, 5.5746385e-10, 5.616799e-10, 5.6594107e-10,
    <a id="L149"></a>5.7024857e-10, 5.746037e-10, 5.7900773e-10, 5.834621e-10,
    <a id="L150"></a>5.8796823e-10, 5.925276e-10, 5.971417e-10, 6.018122e-10,
    <a id="L151"></a>6.065408e-10, 6.113292e-10, 6.1617933e-10, 6.2109295e-10,
    <a id="L152"></a>6.260722e-10, 6.3111916e-10, 6.3623595e-10, 6.4142497e-10,
    <a id="L153"></a>6.4668854e-10, 6.5202926e-10, 6.5744976e-10, 6.6295286e-10,
    <a id="L154"></a>6.6854156e-10, 6.742188e-10, 6.79988e-10, 6.858526e-10,
    <a id="L155"></a>6.9181616e-10, 6.978826e-10, 7.04056e-10, 7.103407e-10,
    <a id="L156"></a>7.167412e-10, 7.2326256e-10, 7.2990985e-10, 7.366886e-10,
    <a id="L157"></a>7.4360473e-10, 7.5066453e-10, 7.5787476e-10, 7.6524265e-10,
    <a id="L158"></a>7.7277595e-10, 7.80483e-10, 7.883728e-10, 7.9645507e-10,
    <a id="L159"></a>8.047402e-10, 8.1323964e-10, 8.219657e-10, 8.309319e-10,
    <a id="L160"></a>8.401528e-10, 8.496445e-10, 8.594247e-10, 8.6951274e-10,
    <a id="L161"></a>8.799301e-10, 8.9070046e-10, 9.018503e-10, 9.134092e-10,
    <a id="L162"></a>9.254101e-10, 9.378904e-10, 9.508923e-10, 9.644638e-10,
    <a id="L163"></a>9.786603e-10, 9.935448e-10, 1.0091913e-09, 1.025686e-09,
    <a id="L164"></a>1.0431306e-09, 1.0616465e-09, 1.08138e-09, 1.1025096e-09,
    <a id="L165"></a>1.1252564e-09, 1.1498986e-09, 1.1767932e-09, 1.206409e-09,
    <a id="L166"></a>1.2393786e-09, 1.276585e-09, 1.3193139e-09, 1.3695435e-09,
    <a id="L167"></a>1.4305498e-09, 1.508365e-09, 1.6160854e-09, 1.7921248e-09,
<a id="L168"></a>}
<a id="L169"></a>var fe = [256]float32{
    <a id="L170"></a>1, 0.9381437, 0.90046996, 0.87170434, 0.8477855, 0.8269933,
    <a id="L171"></a>0.8084217, 0.7915276, 0.77595687, 0.7614634, 0.7478686,
    <a id="L172"></a>0.7350381, 0.72286767, 0.71127474, 0.70019263, 0.6895665,
    <a id="L173"></a>0.67935055, 0.6695063, 0.66000086, 0.65080583, 0.6418967,
    <a id="L174"></a>0.63325197, 0.6248527, 0.6166822, 0.60872537, 0.60096896,
    <a id="L175"></a>0.5934009, 0.58601034, 0.5787874, 0.57172304, 0.5648092,
    <a id="L176"></a>0.5580383, 0.5514034, 0.5448982, 0.5385169, 0.53225386,
    <a id="L177"></a>0.5261042, 0.52006316, 0.5141264, 0.50828975, 0.5025495,
    <a id="L178"></a>0.496902, 0.49134386, 0.485872, 0.48048335, 0.4751752,
    <a id="L179"></a>0.46994483, 0.46478975, 0.45970762, 0.45469615, 0.44975325,
    <a id="L180"></a>0.44487688, 0.44006512, 0.43531612, 0.43062815, 0.42599955,
    <a id="L181"></a>0.42142874, 0.4169142, 0.41245446, 0.40804818, 0.403694,
    <a id="L182"></a>0.3993907, 0.39513698, 0.39093173, 0.38677382, 0.38266218,
    <a id="L183"></a>0.37859577, 0.37457356, 0.37059465, 0.3666581, 0.362763,
    <a id="L184"></a>0.35890847, 0.35509375, 0.351318, 0.3475805, 0.34388044,
    <a id="L185"></a>0.34021714, 0.3365899, 0.33299807, 0.32944095, 0.32591796,
    <a id="L186"></a>0.3224285, 0.3189719, 0.31554767, 0.31215525, 0.30879408,
    <a id="L187"></a>0.3054636, 0.3021634, 0.29889292, 0.2956517, 0.29243928,
    <a id="L188"></a>0.28925523, 0.28609908, 0.28297043, 0.27986884, 0.27679393,
    <a id="L189"></a>0.2737453, 0.2707226, 0.2677254, 0.26475343, 0.26180625,
    <a id="L190"></a>0.25888354, 0.25598502, 0.2531103, 0.25025907, 0.24743107,
    <a id="L191"></a>0.24462597, 0.24184346, 0.23908329, 0.23634516, 0.23362878,
    <a id="L192"></a>0.23093392, 0.2282603, 0.22560766, 0.22297576, 0.22036438,
    <a id="L193"></a>0.21777324, 0.21520215, 0.21265087, 0.21011916, 0.20760682,
    <a id="L194"></a>0.20511365, 0.20263945, 0.20018397, 0.19774707, 0.19532852,
    <a id="L195"></a>0.19292815, 0.19054577, 0.1881812, 0.18583426, 0.18350479,
    <a id="L196"></a>0.1811926, 0.17889754, 0.17661946, 0.17435817, 0.17211354,
    <a id="L197"></a>0.1698854, 0.16767362, 0.16547804, 0.16329853, 0.16113494,
    <a id="L198"></a>0.15898713, 0.15685499, 0.15473837, 0.15263714, 0.15055119,
    <a id="L199"></a>0.14848037, 0.14642459, 0.14438373, 0.14235765, 0.14034624,
    <a id="L200"></a>0.13834943, 0.13636707, 0.13439907, 0.13244532, 0.13050574,
    <a id="L201"></a>0.1285802, 0.12666863, 0.12477092, 0.12288698, 0.12101672,
    <a id="L202"></a>0.119160056, 0.1173169, 0.115487166, 0.11367077, 0.11186763,
    <a id="L203"></a>0.11007768, 0.10830083, 0.10653701, 0.10478614, 0.10304816,
    <a id="L204"></a>0.101323, 0.09961058, 0.09791085, 0.09622374, 0.09454919,
    <a id="L205"></a>0.09288713, 0.091237515, 0.08960028, 0.087975375, 0.08636274,
    <a id="L206"></a>0.08476233, 0.083174095, 0.081597984, 0.08003395, 0.07848195,
    <a id="L207"></a>0.076941945, 0.07541389, 0.07389775, 0.072393484, 0.07090106,
    <a id="L208"></a>0.069420435, 0.06795159, 0.066494495, 0.06504912, 0.063615434,
    <a id="L209"></a>0.062193416, 0.060783047, 0.059384305, 0.057997175,
    <a id="L210"></a>0.05662164, 0.05525769, 0.053905312, 0.052564494, 0.051235236,
    <a id="L211"></a>0.049917534, 0.048611384, 0.047316793, 0.046033762, 0.0447623,
    <a id="L212"></a>0.043502413, 0.042254124, 0.041017443, 0.039792392,
    <a id="L213"></a>0.038578995, 0.037377283, 0.036187284, 0.035009038,
    <a id="L214"></a>0.033842582, 0.032687962, 0.031545233, 0.030414443, 0.02929566,
    <a id="L215"></a>0.02818895, 0.027094385, 0.026012046, 0.024942026, 0.023884421,
    <a id="L216"></a>0.022839336, 0.021806888, 0.020787204, 0.019780423, 0.0187867,
    <a id="L217"></a>0.0178062, 0.016839107, 0.015885621, 0.014945968, 0.014020392,
    <a id="L218"></a>0.013109165, 0.012212592, 0.011331013, 0.01046481, 0.009614414,
    <a id="L219"></a>0.008780315, 0.007963077, 0.0071633533, 0.006381906,
    <a id="L220"></a>0.0056196423, 0.0048776558, 0.004157295, 0.0034602648,
    <a id="L221"></a>0.0027887989, 0.0021459677, 0.0015362998, 0.0009672693,
    <a id="L222"></a>0.00045413437,
<a id="L223"></a>}
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
