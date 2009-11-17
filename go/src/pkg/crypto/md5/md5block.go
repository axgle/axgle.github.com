<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/crypto/md5/md5block.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../doc/style.css">
  <script type="text/javascript" src="../../../../doc/godocs.js"></script>

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
        <a href="../../../../index.html"><img src="../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../index.html">Source files</a></li>

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
  <h1 id="generatedHeader">Source file /src/pkg/crypto/md5/md5block.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// MD5 block step.</span>
<a id="L6"></a><span class="comment">// In its own file so that a faster assembly or C version</span>
<a id="L7"></a><span class="comment">// can be substituted easily.</span>

<a id="L9"></a>package md5

<a id="L11"></a><span class="comment">// table[i] = int((1&lt;&lt;32) * abs(sin(i+1 radians))).</span>
<a id="L12"></a>var table = []uint32{
    <a id="L13"></a><span class="comment">// round 1</span>
    <a id="L14"></a>0xd76aa478,
    <a id="L15"></a>0xe8c7b756,
    <a id="L16"></a>0x242070db,
    <a id="L17"></a>0xc1bdceee,
    <a id="L18"></a>0xf57c0faf,
    <a id="L19"></a>0x4787c62a,
    <a id="L20"></a>0xa8304613,
    <a id="L21"></a>0xfd469501,
    <a id="L22"></a>0x698098d8,
    <a id="L23"></a>0x8b44f7af,
    <a id="L24"></a>0xffff5bb1,
    <a id="L25"></a>0x895cd7be,
    <a id="L26"></a>0x6b901122,
    <a id="L27"></a>0xfd987193,
    <a id="L28"></a>0xa679438e,
    <a id="L29"></a>0x49b40821,

    <a id="L31"></a><span class="comment">// round 2</span>
    <a id="L32"></a>0xf61e2562,
    <a id="L33"></a>0xc040b340,
    <a id="L34"></a>0x265e5a51,
    <a id="L35"></a>0xe9b6c7aa,
    <a id="L36"></a>0xd62f105d,
    <a id="L37"></a>0x2441453,
    <a id="L38"></a>0xd8a1e681,
    <a id="L39"></a>0xe7d3fbc8,
    <a id="L40"></a>0x21e1cde6,
    <a id="L41"></a>0xc33707d6,
    <a id="L42"></a>0xf4d50d87,
    <a id="L43"></a>0x455a14ed,
    <a id="L44"></a>0xa9e3e905,
    <a id="L45"></a>0xfcefa3f8,
    <a id="L46"></a>0x676f02d9,
    <a id="L47"></a>0x8d2a4c8a,

    <a id="L49"></a><span class="comment">// round3</span>
    <a id="L50"></a>0xfffa3942,
    <a id="L51"></a>0x8771f681,
    <a id="L52"></a>0x6d9d6122,
    <a id="L53"></a>0xfde5380c,
    <a id="L54"></a>0xa4beea44,
    <a id="L55"></a>0x4bdecfa9,
    <a id="L56"></a>0xf6bb4b60,
    <a id="L57"></a>0xbebfbc70,
    <a id="L58"></a>0x289b7ec6,
    <a id="L59"></a>0xeaa127fa,
    <a id="L60"></a>0xd4ef3085,
    <a id="L61"></a>0x4881d05,
    <a id="L62"></a>0xd9d4d039,
    <a id="L63"></a>0xe6db99e5,
    <a id="L64"></a>0x1fa27cf8,
    <a id="L65"></a>0xc4ac5665,

    <a id="L67"></a><span class="comment">// round 4</span>
    <a id="L68"></a>0xf4292244,
    <a id="L69"></a>0x432aff97,
    <a id="L70"></a>0xab9423a7,
    <a id="L71"></a>0xfc93a039,
    <a id="L72"></a>0x655b59c3,
    <a id="L73"></a>0x8f0ccc92,
    <a id="L74"></a>0xffeff47d,
    <a id="L75"></a>0x85845dd1,
    <a id="L76"></a>0x6fa87e4f,
    <a id="L77"></a>0xfe2ce6e0,
    <a id="L78"></a>0xa3014314,
    <a id="L79"></a>0x4e0811a1,
    <a id="L80"></a>0xf7537e82,
    <a id="L81"></a>0xbd3af235,
    <a id="L82"></a>0x2ad7d2bb,
    <a id="L83"></a>0xeb86d391,
<a id="L84"></a>}

<a id="L86"></a>var shift1 = []uint{7, 12, 17, 22}
<a id="L87"></a>var shift2 = []uint{5, 9, 14, 20}
<a id="L88"></a>var shift3 = []uint{4, 11, 16, 23}
<a id="L89"></a>var shift4 = []uint{6, 10, 15, 21}

<a id="L91"></a>func _Block(dig *digest, p []byte) int {
    <a id="L92"></a>a := dig.s[0];
    <a id="L93"></a>b := dig.s[1];
    <a id="L94"></a>c := dig.s[2];
    <a id="L95"></a>d := dig.s[3];
    <a id="L96"></a>n := 0;
    <a id="L97"></a>var X [16]uint32;
    <a id="L98"></a>for len(p) &gt;= _Chunk {
        <a id="L99"></a>aa, bb, cc, dd := a, b, c, d;

        <a id="L101"></a>for i := 0; i &lt; 16; i++ {
            <a id="L102"></a>j := i * 4;
            <a id="L103"></a>X[i] = uint32(p[j]) | uint32(p[j+1])&lt;&lt;8 | uint32(p[j+2])&lt;&lt;16 | uint32(p[j+3])&lt;&lt;24;
        <a id="L104"></a>}

        <a id="L106"></a><span class="comment">// If this needs to be made faster in the future,</span>
        <a id="L107"></a><span class="comment">// the usual trick is to unroll each of these</span>
        <a id="L108"></a><span class="comment">// loops by a factor of 4; that lets you replace</span>
        <a id="L109"></a><span class="comment">// the shift[] lookups with constants and,</span>
        <a id="L110"></a><span class="comment">// with suitable variable renaming in each</span>
        <a id="L111"></a><span class="comment">// unrolled body, delete the a, b, c, d = d, a, b, c</span>
        <a id="L112"></a><span class="comment">// (or you can let the optimizer do the renaming).</span>

        <a id="L114"></a><span class="comment">// Round 1.</span>
        <a id="L115"></a>for i := 0; i &lt; 16; i++ {
            <a id="L116"></a>x := i;
            <a id="L117"></a>t := i;
            <a id="L118"></a>s := shift1[i%4];
            <a id="L119"></a>f := ((c ^ d) &amp; b) ^ d;
            <a id="L120"></a>a += f + X[x] + table[t];
            <a id="L121"></a>a = a&lt;&lt;s | a&gt;&gt;(32-s);
            <a id="L122"></a>a += b;
            <a id="L123"></a>a, b, c, d = d, a, b, c;
        <a id="L124"></a>}

        <a id="L126"></a><span class="comment">// Round 2.</span>
        <a id="L127"></a>for i := 0; i &lt; 16; i++ {
            <a id="L128"></a>x := (1 + 5*i) % 16;
            <a id="L129"></a>t := 16 + i;
            <a id="L130"></a>s := shift2[i%4];
            <a id="L131"></a>g := ((b ^ c) &amp; d) ^ c;
            <a id="L132"></a>a += g + X[x] + table[t];
            <a id="L133"></a>a = a&lt;&lt;s | a&gt;&gt;(32-s);
            <a id="L134"></a>a += b;
            <a id="L135"></a>a, b, c, d = d, a, b, c;
        <a id="L136"></a>}

        <a id="L138"></a><span class="comment">// Round 3.</span>
        <a id="L139"></a>for i := 0; i &lt; 16; i++ {
            <a id="L140"></a>x := (5 + 3*i) % 16;
            <a id="L141"></a>t := 32 + i;
            <a id="L142"></a>s := shift3[i%4];
            <a id="L143"></a>h := b ^ c ^ d;
            <a id="L144"></a>a += h + X[x] + table[t];
            <a id="L145"></a>a = a&lt;&lt;s | a&gt;&gt;(32-s);
            <a id="L146"></a>a += b;
            <a id="L147"></a>a, b, c, d = d, a, b, c;
        <a id="L148"></a>}

        <a id="L150"></a><span class="comment">// Round 4.</span>
        <a id="L151"></a>for i := 0; i &lt; 16; i++ {
            <a id="L152"></a>x := (7 * i) % 16;
            <a id="L153"></a>s := shift4[i%4];
            <a id="L154"></a>t := 48 + i;
            <a id="L155"></a>ii := c ^ (b | ^d);
            <a id="L156"></a>a += ii + X[x] + table[t];
            <a id="L157"></a>a = a&lt;&lt;s | a&gt;&gt;(32-s);
            <a id="L158"></a>a += b;
            <a id="L159"></a>a, b, c, d = d, a, b, c;
        <a id="L160"></a>}

        <a id="L162"></a>a += aa;
        <a id="L163"></a>b += bb;
        <a id="L164"></a>c += cc;
        <a id="L165"></a>d += dd;

        <a id="L167"></a>p = p[_Chunk:len(p)];
        <a id="L168"></a>n += _Chunk;
    <a id="L169"></a>}

    <a id="L171"></a>dig.s[0] = a;
    <a id="L172"></a>dig.s[1] = b;
    <a id="L173"></a>dig.s[2] = c;
    <a id="L174"></a>dig.s[3] = d;
    <a id="L175"></a>return n;
<a id="L176"></a>}
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
