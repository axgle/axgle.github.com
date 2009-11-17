<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/image/image.go</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/image/image.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The image package implements a basic 2-D image library.</span>
<a id="L6"></a>package image

<a id="L8"></a><span class="comment">// An Image is a rectangular grid of Colors drawn from a ColorModel.</span>
<a id="L9"></a>type Image interface {
    <a id="L10"></a>ColorModel() ColorModel;
    <a id="L11"></a>Width() int;
    <a id="L12"></a>Height() int;
    <a id="L13"></a><span class="comment">// At(0, 0) returns the upper-left pixel of the grid.</span>
    <a id="L14"></a><span class="comment">// At(Width()-1, Height()-1) returns the lower-right pixel.</span>
    <a id="L15"></a>At(x, y int) Color;
<a id="L16"></a>}

<a id="L18"></a><span class="comment">// An RGBA is an in-memory image backed by a 2-D slice of RGBAColor values.</span>
<a id="L19"></a>type RGBA struct {
    <a id="L20"></a><span class="comment">// The Pixel field&#39;s indices are y first, then x, so that At(x, y) == Pixel[y][x].</span>
    <a id="L21"></a>Pixel [][]RGBAColor;
<a id="L22"></a>}

<a id="L24"></a>func (p *RGBA) ColorModel() ColorModel { return RGBAColorModel }

<a id="L26"></a>func (p *RGBA) Width() int {
    <a id="L27"></a>if len(p.Pixel) == 0 {
        <a id="L28"></a>return 0
    <a id="L29"></a>}
    <a id="L30"></a>return len(p.Pixel[0]);
<a id="L31"></a>}

<a id="L33"></a>func (p *RGBA) Height() int { return len(p.Pixel) }

<a id="L35"></a>func (p *RGBA) At(x, y int) Color { return p.Pixel[y][x] }

<a id="L37"></a>func (p *RGBA) Set(x, y int, c Color) { p.Pixel[y][x] = toRGBAColor(c).(RGBAColor) }

<a id="L39"></a><span class="comment">// NewRGBA returns a new RGBA with the given width and height.</span>
<a id="L40"></a>func NewRGBA(w, h int) *RGBA {
    <a id="L41"></a>pixel := make([][]RGBAColor, h);
    <a id="L42"></a>for y := 0; y &lt; int(h); y++ {
        <a id="L43"></a>pixel[y] = make([]RGBAColor, w)
    <a id="L44"></a>}
    <a id="L45"></a>return &amp;RGBA{pixel};
<a id="L46"></a>}

<a id="L48"></a><span class="comment">// An RGBA64 is an in-memory image backed by a 2-D slice of RGBA64Color values.</span>
<a id="L49"></a>type RGBA64 struct {
    <a id="L50"></a><span class="comment">// The Pixel field&#39;s indices are y first, then x, so that At(x, y) == Pixel[y][x].</span>
    <a id="L51"></a>Pixel [][]RGBA64Color;
<a id="L52"></a>}

<a id="L54"></a>func (p *RGBA64) ColorModel() ColorModel { return RGBA64ColorModel }

<a id="L56"></a>func (p *RGBA64) Width() int {
    <a id="L57"></a>if len(p.Pixel) == 0 {
        <a id="L58"></a>return 0
    <a id="L59"></a>}
    <a id="L60"></a>return len(p.Pixel[0]);
<a id="L61"></a>}

<a id="L63"></a>func (p *RGBA64) Height() int { return len(p.Pixel) }

<a id="L65"></a>func (p *RGBA64) At(x, y int) Color { return p.Pixel[y][x] }

<a id="L67"></a>func (p *RGBA64) Set(x, y int, c Color) { p.Pixel[y][x] = toRGBA64Color(c).(RGBA64Color) }

<a id="L69"></a><span class="comment">// NewRGBA64 returns a new RGBA64 with the given width and height.</span>
<a id="L70"></a>func NewRGBA64(w, h int) *RGBA64 {
    <a id="L71"></a>pixel := make([][]RGBA64Color, h);
    <a id="L72"></a>for y := 0; y &lt; int(h); y++ {
        <a id="L73"></a>pixel[y] = make([]RGBA64Color, w)
    <a id="L74"></a>}
    <a id="L75"></a>return &amp;RGBA64{pixel};
<a id="L76"></a>}

<a id="L78"></a><span class="comment">// A NRGBA is an in-memory image backed by a 2-D slice of NRGBAColor values.</span>
<a id="L79"></a>type NRGBA struct {
    <a id="L80"></a><span class="comment">// The Pixel field&#39;s indices are y first, then x, so that At(x, y) == Pixel[y][x].</span>
    <a id="L81"></a>Pixel [][]NRGBAColor;
<a id="L82"></a>}

<a id="L84"></a>func (p *NRGBA) ColorModel() ColorModel { return NRGBAColorModel }

<a id="L86"></a>func (p *NRGBA) Width() int {
    <a id="L87"></a>if len(p.Pixel) == 0 {
        <a id="L88"></a>return 0
    <a id="L89"></a>}
    <a id="L90"></a>return len(p.Pixel[0]);
<a id="L91"></a>}

<a id="L93"></a>func (p *NRGBA) Height() int { return len(p.Pixel) }

<a id="L95"></a>func (p *NRGBA) At(x, y int) Color { return p.Pixel[y][x] }

<a id="L97"></a>func (p *NRGBA) Set(x, y int, c Color) { p.Pixel[y][x] = toNRGBAColor(c).(NRGBAColor) }

<a id="L99"></a><span class="comment">// NewNRGBA returns a new NRGBA with the given width and height.</span>
<a id="L100"></a>func NewNRGBA(w, h int) *NRGBA {
    <a id="L101"></a>pixel := make([][]NRGBAColor, h);
    <a id="L102"></a>for y := 0; y &lt; int(h); y++ {
        <a id="L103"></a>pixel[y] = make([]NRGBAColor, w)
    <a id="L104"></a>}
    <a id="L105"></a>return &amp;NRGBA{pixel};
<a id="L106"></a>}

<a id="L108"></a><span class="comment">// A NRGBA64 is an in-memory image backed by a 2-D slice of NRGBA64Color values.</span>
<a id="L109"></a>type NRGBA64 struct {
    <a id="L110"></a><span class="comment">// The Pixel field&#39;s indices are y first, then x, so that At(x, y) == Pixel[y][x].</span>
    <a id="L111"></a>Pixel [][]NRGBA64Color;
<a id="L112"></a>}

<a id="L114"></a>func (p *NRGBA64) ColorModel() ColorModel { return NRGBA64ColorModel }

<a id="L116"></a>func (p *NRGBA64) Width() int {
    <a id="L117"></a>if len(p.Pixel) == 0 {
        <a id="L118"></a>return 0
    <a id="L119"></a>}
    <a id="L120"></a>return len(p.Pixel[0]);
<a id="L121"></a>}

<a id="L123"></a>func (p *NRGBA64) Height() int { return len(p.Pixel) }

<a id="L125"></a>func (p *NRGBA64) At(x, y int) Color { return p.Pixel[y][x] }

<a id="L127"></a>func (p *NRGBA64) Set(x, y int, c Color) { p.Pixel[y][x] = toNRGBA64Color(c).(NRGBA64Color) }

<a id="L129"></a><span class="comment">// NewNRGBA64 returns a new NRGBA64 with the given width and height.</span>
<a id="L130"></a>func NewNRGBA64(w, h int) *NRGBA64 {
    <a id="L131"></a>pixel := make([][]NRGBA64Color, h);
    <a id="L132"></a>for y := 0; y &lt; int(h); y++ {
        <a id="L133"></a>pixel[y] = make([]NRGBA64Color, w)
    <a id="L134"></a>}
    <a id="L135"></a>return &amp;NRGBA64{pixel};
<a id="L136"></a>}

<a id="L138"></a><span class="comment">// A PalettedColorModel represents a fixed palette of colors.</span>
<a id="L139"></a>type PalettedColorModel []Color

<a id="L141"></a>func diff(a, b uint32) uint32 {
    <a id="L142"></a>if a &gt; b {
        <a id="L143"></a>return a - b
    <a id="L144"></a>}
    <a id="L145"></a>return b - a;
<a id="L146"></a>}

<a id="L148"></a><span class="comment">// Convert returns the palette color closest to c in Euclidean R,G,B space.</span>
<a id="L149"></a>func (p PalettedColorModel) Convert(c Color) Color {
    <a id="L150"></a>if len(p) == 0 {
        <a id="L151"></a>return nil
    <a id="L152"></a>}
    <a id="L153"></a><span class="comment">// TODO(nigeltao): Revisit the &#34;pick the palette color which minimizes sum-squared-difference&#34;</span>
    <a id="L154"></a><span class="comment">// algorithm when the premultiplied vs unpremultiplied issue is resolved.</span>
    <a id="L155"></a><span class="comment">// Currently, we only compare the R, G and B values, and ignore A.</span>
    <a id="L156"></a>cr, cg, cb, _ := c.RGBA();
    <a id="L157"></a><span class="comment">// Shift by 17 bits to avoid potential uint32 overflow in sum-squared-difference.</span>
    <a id="L158"></a>cr &gt;&gt;= 17;
    <a id="L159"></a>cg &gt;&gt;= 17;
    <a id="L160"></a>cb &gt;&gt;= 17;
    <a id="L161"></a>result := Color(nil);
    <a id="L162"></a>bestSSD := uint32(1&lt;&lt;32 - 1);
    <a id="L163"></a>for _, v := range p {
        <a id="L164"></a>vr, vg, vb, _ := v.RGBA();
        <a id="L165"></a>vr &gt;&gt;= 17;
        <a id="L166"></a>vg &gt;&gt;= 17;
        <a id="L167"></a>vb &gt;&gt;= 17;
        <a id="L168"></a>dr, dg, db := diff(cr, vr), diff(cg, vg), diff(cb, vb);
        <a id="L169"></a>ssd := (dr * dr) + (dg * dg) + (db * db);
        <a id="L170"></a>if ssd &lt; bestSSD {
            <a id="L171"></a>bestSSD = ssd;
            <a id="L172"></a>result = v;
        <a id="L173"></a>}
    <a id="L174"></a>}
    <a id="L175"></a>return result;
<a id="L176"></a>}

<a id="L178"></a><span class="comment">// A Paletted is an in-memory image backed by a 2-D slice of uint8 values and a PalettedColorModel.</span>
<a id="L179"></a>type Paletted struct {
    <a id="L180"></a><span class="comment">// The Pixel field&#39;s indices are y first, then x, so that At(x, y) == Palette[Pixel[y][x]].</span>
    <a id="L181"></a>Pixel   [][]uint8;
    <a id="L182"></a>Palette PalettedColorModel;
<a id="L183"></a>}

<a id="L185"></a>func (p *Paletted) ColorModel() ColorModel { return p.Palette }

<a id="L187"></a>func (p *Paletted) Width() int {
    <a id="L188"></a>if len(p.Pixel) == 0 {
        <a id="L189"></a>return 0
    <a id="L190"></a>}
    <a id="L191"></a>return len(p.Pixel[0]);
<a id="L192"></a>}

<a id="L194"></a>func (p *Paletted) Height() int { return len(p.Pixel) }

<a id="L196"></a>func (p *Paletted) At(x, y int) Color { return p.Palette[p.Pixel[y][x]] }

<a id="L198"></a>func (p *Paletted) ColorIndexAt(x, y int) uint8 {
    <a id="L199"></a>return p.Pixel[y][x]
<a id="L200"></a>}

<a id="L202"></a>func (p *Paletted) SetColorIndex(x, y int, index uint8) {
    <a id="L203"></a>p.Pixel[y][x] = index
<a id="L204"></a>}

<a id="L206"></a><span class="comment">// NewPaletted returns a new Paletted with the given width, height and palette.</span>
<a id="L207"></a>func NewPaletted(w, h int, m PalettedColorModel) *Paletted {
    <a id="L208"></a>pixel := make([][]uint8, h);
    <a id="L209"></a>for y := 0; y &lt; int(h); y++ {
        <a id="L210"></a>pixel[y] = make([]uint8, w)
    <a id="L211"></a>}
    <a id="L212"></a>return &amp;Paletted{pixel, m};
<a id="L213"></a>}
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
