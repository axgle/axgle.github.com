<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/nacl/av/image.go</title>

  <link rel="stylesheet" type="text/css" href="../../../../../doc/style.css">
  <script type="text/javascript" src="../../../../../doc/godocs.js"></script>

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
        <a href="../../../../../index.html"><img src="../../../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../../../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/nacl/av/image.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package av

<a id="L7"></a>import (
    <a id="L8"></a>&#34;image&#34;;
<a id="L9"></a>)

<a id="L11"></a><span class="comment">// Native Client image format:</span>
<a id="L12"></a><span class="comment">// a single linear array of 32-bit ARGB as packed uint32s.</span>

<a id="L14"></a><span class="comment">// An Image represents a Native Client frame buffer.</span>
<a id="L15"></a><span class="comment">// The pixels in the image can be accessed as a single</span>
<a id="L16"></a><span class="comment">// linear slice or as a two-dimensional slice of slices.</span>
<a id="L17"></a><span class="comment">// Image implements image.Image.</span>
<a id="L18"></a>type Image struct {
    <a id="L19"></a>Linear []Color;
    <a id="L20"></a>Pixel  [][]Color;
<a id="L21"></a>}

<a id="L23"></a>var _ image.Image = (*Image)(nil)

<a id="L25"></a>func (m *Image) ColorModel() image.ColorModel { return ColorModel }

<a id="L27"></a>func (m *Image) Width() int {
    <a id="L28"></a>if len(m.Pixel) == 0 {
        <a id="L29"></a>return 0
    <a id="L30"></a>}
    <a id="L31"></a>return len(m.Pixel[0]);
<a id="L32"></a>}

<a id="L34"></a>func (m *Image) Height() int { return len(m.Pixel) }

<a id="L36"></a>func (m *Image) At(x, y int) image.Color { return m.Pixel[y][x] }

<a id="L38"></a>func (m *Image) Set(x, y int, color image.Color) {
    <a id="L39"></a>if c, ok := color.(Color); ok {
        <a id="L40"></a>m.Pixel[y][x] = c
    <a id="L41"></a>}
    <a id="L42"></a>m.Pixel[y][x] = makeColor(color.RGBA());
<a id="L43"></a>}

<a id="L45"></a>func newImage(dx, dy int, linear []Color) *Image {
    <a id="L46"></a>if linear == nil {
        <a id="L47"></a>linear = make([]Color, dx*dy)
    <a id="L48"></a>}
    <a id="L49"></a>pix := make([][]Color, dy);
    <a id="L50"></a>for i := range pix {
        <a id="L51"></a>pix[i] = linear[dx*i : dx*(i+1)]
    <a id="L52"></a>}
    <a id="L53"></a>return &amp;Image{linear, pix};
<a id="L54"></a>}

<a id="L56"></a><span class="comment">// A Color represents a Native Client color value,</span>
<a id="L57"></a><span class="comment">// a 32-bit R, G, B, A value packed as 0xAARRGGBB.</span>
<a id="L58"></a>type Color uint32

<a id="L60"></a>func (p Color) RGBA() (r, g, b, a uint32) {
    <a id="L61"></a>x := uint32(p);
    <a id="L62"></a>a = x &gt;&gt; 24;
    <a id="L63"></a>a |= a &lt;&lt; 8;
    <a id="L64"></a>a |= a &lt;&lt; 16;
    <a id="L65"></a>r = (x &gt;&gt; 16) &amp; 0xFF;
    <a id="L66"></a>r |= r &lt;&lt; 8;
    <a id="L67"></a>r |= r &lt;&lt; 16;
    <a id="L68"></a>g = (x &gt;&gt; 8) &amp; 0xFF;
    <a id="L69"></a>g |= g &lt;&lt; 8;
    <a id="L70"></a>g |= g &lt;&lt; 16;
    <a id="L71"></a>b = x &amp; 0xFF;
    <a id="L72"></a>b |= b &lt;&lt; 8;
    <a id="L73"></a>b |= b &lt;&lt; 16;
    <a id="L74"></a>return;
<a id="L75"></a>}

<a id="L77"></a>func makeColor(r, g, b, a uint32) Color {
    <a id="L78"></a>return Color(a&gt;&gt;24&lt;&lt;24 | r&gt;&gt;24&lt;&lt;16 | g&gt;&gt;24&lt;&lt;8 | b&gt;&gt;24)
<a id="L79"></a>}

<a id="L81"></a>func toColor(color image.Color) image.Color {
    <a id="L82"></a>if c, ok := color.(Color); ok {
        <a id="L83"></a>return c
    <a id="L84"></a>}
    <a id="L85"></a>return makeColor(color.RGBA());
<a id="L86"></a>}

<a id="L88"></a><span class="comment">// ColorModel is the color model corresponding to the Native Client Color.</span>
<a id="L89"></a>var ColorModel = image.ColorModelFunc(toColor)
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
