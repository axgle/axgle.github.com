<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/draw/color.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/draw/color.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package draw

<a id="L7"></a>import &#34;image&#34;

<a id="L9"></a><span class="comment">// A Color represents a color with 8-bit R, G, B, and A values,</span>
<a id="L10"></a><span class="comment">// packed into a uint32—0xRRGGBBAA—so that comparison</span>
<a id="L11"></a><span class="comment">// is defined on colors.</span>
<a id="L12"></a><span class="comment">// Color implements image.Color.</span>
<a id="L13"></a><span class="comment">// Color also implements image.Image: it is a</span>
<a id="L14"></a><span class="comment">// 10⁹x10⁹-pixel image of uniform color.</span>
<a id="L15"></a>type Color uint32

<a id="L17"></a><span class="comment">// Check that Color implements image.Color and image.Image</span>
<a id="L18"></a>var _ image.Color = Black
<a id="L19"></a>var _ image.Image = Black

<a id="L21"></a>var (
    <a id="L22"></a>Opaque        Color = 0xFFFFFFFF;
    <a id="L23"></a>Transparent   Color = 0x00000000;
    <a id="L24"></a>Black         Color = 0x000000FF;
    <a id="L25"></a>White         Color = 0xFFFFFFFF;
    <a id="L26"></a>Red           Color = 0xFF0000FF;
    <a id="L27"></a>Green         Color = 0x00FF00FF;
    <a id="L28"></a>Blue          Color = 0x0000FFFF;
    <a id="L29"></a>Cyan          Color = 0x00FFFFFF;
    <a id="L30"></a>Magenta       Color = 0xFF00FFFF;
    <a id="L31"></a>Yellow        Color = 0xFFFF00FF;
    <a id="L32"></a>PaleYellow    Color = 0xFFFFAAFF;
    <a id="L33"></a>DarkYellow    Color = 0xEEEE9EFF;
    <a id="L34"></a>DarkGreen     Color = 0x448844FF;
    <a id="L35"></a>PaleGreen     Color = 0xAAFFAAFF;
    <a id="L36"></a>MedGreen      Color = 0x88CC88FF;
    <a id="L37"></a>DarkBlue      Color = 0x000055FF;
    <a id="L38"></a>PaleBlueGreen Color = 0xAAFFFFFF;
    <a id="L39"></a>PaleBlue      Color = 0x0000BBFF;
    <a id="L40"></a>BlueGreen     Color = 0x008888FF;
    <a id="L41"></a>GreyGreen     Color = 0x55AAAAFF;
    <a id="L42"></a>PaleGreyGreen Color = 0x9EEEEEFF;
    <a id="L43"></a>YellowGreen   Color = 0x99994CFF;
    <a id="L44"></a>MedBlue       Color = 0x000099FF;
    <a id="L45"></a>GreyBlue      Color = 0x005DBBFF;
    <a id="L46"></a>PaleGreyBlue  Color = 0x4993DDFF;
    <a id="L47"></a>PurpleBlue    Color = 0x8888CCFF;
<a id="L48"></a>)

<a id="L50"></a>func (c Color) RGBA() (r, g, b, a uint32) {
    <a id="L51"></a>x := uint32(c);
    <a id="L52"></a>r, g, b, a = x&gt;&gt;24, (x&gt;&gt;16)&amp;0xFF, (x&gt;&gt;8)&amp;0xFF, x&amp;0xFF;
    <a id="L53"></a>r |= r &lt;&lt; 8;
    <a id="L54"></a>r |= r &lt;&lt; 16;
    <a id="L55"></a>g |= g &lt;&lt; 8;
    <a id="L56"></a>g |= g &lt;&lt; 16;
    <a id="L57"></a>b |= b &lt;&lt; 8;
    <a id="L58"></a>b |= b &lt;&lt; 16;
    <a id="L59"></a>a |= a &lt;&lt; 8;
    <a id="L60"></a>a |= a &lt;&lt; 16;
    <a id="L61"></a>return;
<a id="L62"></a>}

<a id="L64"></a><span class="comment">// SetAlpha returns the color obtained by changing</span>
<a id="L65"></a><span class="comment">// c&#39;s alpha value to a and scaling r, g, and b appropriately.</span>
<a id="L66"></a>func (c Color) SetAlpha(a uint8) Color {
    <a id="L67"></a>r, g, b, oa := c&gt;&gt;24, (c&gt;&gt;16)&amp;0xFF, (c&gt;&gt;8)&amp;0xFF, c&amp;0xFF;
    <a id="L68"></a>if oa == 0 {
        <a id="L69"></a>return 0
    <a id="L70"></a>}
    <a id="L71"></a>r = r * Color(a) / oa;
    <a id="L72"></a>if r &lt; 0 {
        <a id="L73"></a>r = 0
    <a id="L74"></a>}
    <a id="L75"></a>if r &gt; 0xFF {
        <a id="L76"></a>r = 0xFF
    <a id="L77"></a>}
    <a id="L78"></a>g = g * Color(a) / oa;
    <a id="L79"></a>if g &lt; 0 {
        <a id="L80"></a>g = 0
    <a id="L81"></a>}
    <a id="L82"></a>if g &gt; 0xFF {
        <a id="L83"></a>g = 0xFF
    <a id="L84"></a>}
    <a id="L85"></a>b = b * Color(a) / oa;
    <a id="L86"></a>if b &lt; 0 {
        <a id="L87"></a>b = 0
    <a id="L88"></a>}
    <a id="L89"></a>if b &gt; 0xFF {
        <a id="L90"></a>b = 0xFF
    <a id="L91"></a>}
    <a id="L92"></a>return r&lt;&lt;24 | g&lt;&lt;16 | b&lt;&lt;8 | Color(a);
<a id="L93"></a>}

<a id="L95"></a>func (c Color) Width() int { return 1e9 }

<a id="L97"></a>func (c Color) Height() int { return 1e9 }

<a id="L99"></a>func (c Color) At(x, y int) image.Color { return c }

<a id="L101"></a>func toColor(color image.Color) image.Color {
    <a id="L102"></a>if c, ok := color.(Color); ok {
        <a id="L103"></a>return c
    <a id="L104"></a>}
    <a id="L105"></a>r, g, b, a := color.RGBA();
    <a id="L106"></a>return Color(r&gt;&gt;24&lt;&lt;24 | g&gt;&gt;24&lt;&lt;16 | b&gt;&gt;24&lt;&lt;8 | a&gt;&gt;24);
<a id="L107"></a>}

<a id="L109"></a>func (c Color) ColorModel() image.ColorModel { return image.ColorModelFunc(toColor) }
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
