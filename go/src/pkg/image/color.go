<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/image/color.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/image/color.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package image

<a id="L7"></a><span class="comment">// TODO(nigeltao): Think about how floating-point color models work.</span>

<a id="L9"></a><span class="comment">// All Colors can convert themselves, with a possible loss of precision, to 128-bit alpha-premultiplied RGBA.</span>
<a id="L10"></a>type Color interface {
    <a id="L11"></a>RGBA() (r, g, b, a uint32);
<a id="L12"></a>}

<a id="L14"></a><span class="comment">// An RGBAColor represents a traditional 32-bit alpha-premultiplied color, having 8 bits for each of red, green, blue and alpha.</span>
<a id="L15"></a>type RGBAColor struct {
    <a id="L16"></a>R, G, B, A uint8;
<a id="L17"></a>}

<a id="L19"></a>func (c RGBAColor) RGBA() (r, g, b, a uint32) {
    <a id="L20"></a>r = uint32(c.R);
    <a id="L21"></a>r |= r &lt;&lt; 8;
    <a id="L22"></a>r |= r &lt;&lt; 16;
    <a id="L23"></a>g = uint32(c.G);
    <a id="L24"></a>g |= g &lt;&lt; 8;
    <a id="L25"></a>g |= g &lt;&lt; 16;
    <a id="L26"></a>b = uint32(c.B);
    <a id="L27"></a>b |= b &lt;&lt; 8;
    <a id="L28"></a>b |= b &lt;&lt; 16;
    <a id="L29"></a>a = uint32(c.A);
    <a id="L30"></a>a |= a &lt;&lt; 8;
    <a id="L31"></a>a |= a &lt;&lt; 16;
    <a id="L32"></a>return;
<a id="L33"></a>}

<a id="L35"></a><span class="comment">// An RGBA64Color represents a 64-bit alpha-premultiplied color, having 16 bits for each of red, green, blue and alpha.</span>
<a id="L36"></a>type RGBA64Color struct {
    <a id="L37"></a>R, G, B, A uint16;
<a id="L38"></a>}

<a id="L40"></a>func (c RGBA64Color) RGBA() (r, g, b, a uint32) {
    <a id="L41"></a>r = uint32(c.R);
    <a id="L42"></a>r |= r &lt;&lt; 16;
    <a id="L43"></a>g = uint32(c.G);
    <a id="L44"></a>g |= g &lt;&lt; 16;
    <a id="L45"></a>b = uint32(c.B);
    <a id="L46"></a>b |= b &lt;&lt; 16;
    <a id="L47"></a>a = uint32(c.A);
    <a id="L48"></a>a |= a &lt;&lt; 16;
    <a id="L49"></a>return;
<a id="L50"></a>}

<a id="L52"></a><span class="comment">// An NRGBAColor represents a non-alpha-premultiplied 32-bit color.</span>
<a id="L53"></a>type NRGBAColor struct {
    <a id="L54"></a>R, G, B, A uint8;
<a id="L55"></a>}

<a id="L57"></a>func (c NRGBAColor) RGBA() (r, g, b, a uint32) {
    <a id="L58"></a>r = uint32(c.R);
    <a id="L59"></a>r |= r &lt;&lt; 8;
    <a id="L60"></a>r *= uint32(c.A);
    <a id="L61"></a>r /= 0xff;
    <a id="L62"></a>r |= r &lt;&lt; 16;
    <a id="L63"></a>g = uint32(c.G);
    <a id="L64"></a>g |= g &lt;&lt; 8;
    <a id="L65"></a>g *= uint32(c.A);
    <a id="L66"></a>g /= 0xff;
    <a id="L67"></a>g |= g &lt;&lt; 16;
    <a id="L68"></a>b = uint32(c.B);
    <a id="L69"></a>b |= b &lt;&lt; 8;
    <a id="L70"></a>b *= uint32(c.A);
    <a id="L71"></a>b /= 0xff;
    <a id="L72"></a>b |= b &lt;&lt; 16;
    <a id="L73"></a>a = uint32(c.A);
    <a id="L74"></a>a |= a &lt;&lt; 8;
    <a id="L75"></a>a |= a &lt;&lt; 16;
    <a id="L76"></a>return;
<a id="L77"></a>}

<a id="L79"></a><span class="comment">// An NRGBA64Color represents a non-alpha-premultiplied 64-bit color, having 16 bits for each of red, green, blue and alpha.</span>
<a id="L80"></a>type NRGBA64Color struct {
    <a id="L81"></a>R, G, B, A uint16;
<a id="L82"></a>}

<a id="L84"></a>func (c NRGBA64Color) RGBA() (r, g, b, a uint32) {
    <a id="L85"></a>r = uint32(c.R);
    <a id="L86"></a>r *= uint32(c.A);
    <a id="L87"></a>r /= 0xffff;
    <a id="L88"></a>r |= r &lt;&lt; 16;
    <a id="L89"></a>g = uint32(c.G);
    <a id="L90"></a>g *= uint32(c.A);
    <a id="L91"></a>g /= 0xffff;
    <a id="L92"></a>g |= g &lt;&lt; 16;
    <a id="L93"></a>b = uint32(c.B);
    <a id="L94"></a>b *= uint32(c.A);
    <a id="L95"></a>b /= 0xffff;
    <a id="L96"></a>b |= b &lt;&lt; 16;
    <a id="L97"></a>a = uint32(c.A);
    <a id="L98"></a>a |= a &lt;&lt; 8;
    <a id="L99"></a>a |= a &lt;&lt; 16;
    <a id="L100"></a>return;
<a id="L101"></a>}

<a id="L103"></a><span class="comment">// A ColorModel can convert foreign Colors, with a possible loss of precision, to a Color</span>
<a id="L104"></a><span class="comment">// from its own color model.</span>
<a id="L105"></a>type ColorModel interface {
    <a id="L106"></a>Convert(c Color) Color;
<a id="L107"></a>}

<a id="L109"></a><span class="comment">// The ColorModelFunc type is an adapter to allow the use of an ordinary</span>
<a id="L110"></a><span class="comment">// color conversion function as a ColorModel.  If f is such a function,</span>
<a id="L111"></a><span class="comment">// ColorModelFunc(f) is a ColorModel object that invokes f to implement</span>
<a id="L112"></a><span class="comment">// the conversion.</span>
<a id="L113"></a>type ColorModelFunc func(Color) Color

<a id="L115"></a>func (f ColorModelFunc) Convert(c Color) Color {
    <a id="L116"></a>return f(c)
<a id="L117"></a>}

<a id="L119"></a>func toRGBAColor(c Color) Color {
    <a id="L120"></a>if _, ok := c.(RGBAColor); ok { <span class="comment">// no-op conversion</span>
        <a id="L121"></a>return c
    <a id="L122"></a>}
    <a id="L123"></a>r, g, b, a := c.RGBA();
    <a id="L124"></a>return RGBAColor{uint8(r &gt;&gt; 24), uint8(g &gt;&gt; 24), uint8(b &gt;&gt; 24), uint8(a &gt;&gt; 24)};
<a id="L125"></a>}

<a id="L127"></a>func toRGBA64Color(c Color) Color {
    <a id="L128"></a>if _, ok := c.(RGBA64Color); ok { <span class="comment">// no-op conversion</span>
        <a id="L129"></a>return c
    <a id="L130"></a>}
    <a id="L131"></a>r, g, b, a := c.RGBA();
    <a id="L132"></a>return RGBA64Color{uint16(r &gt;&gt; 16), uint16(g &gt;&gt; 16), uint16(b &gt;&gt; 16), uint16(a &gt;&gt; 16)};
<a id="L133"></a>}

<a id="L135"></a>func toNRGBAColor(c Color) Color {
    <a id="L136"></a>if _, ok := c.(NRGBAColor); ok { <span class="comment">// no-op conversion</span>
        <a id="L137"></a>return c
    <a id="L138"></a>}
    <a id="L139"></a>r, g, b, a := c.RGBA();
    <a id="L140"></a>a &gt;&gt;= 16;
    <a id="L141"></a>if a == 0xffff {
        <a id="L142"></a>return NRGBAColor{uint8(r &gt;&gt; 24), uint8(g &gt;&gt; 24), uint8(b &gt;&gt; 24), 0xff}
    <a id="L143"></a>}
    <a id="L144"></a>if a == 0 {
        <a id="L145"></a>return NRGBAColor{0, 0, 0, 0}
    <a id="L146"></a>}
    <a id="L147"></a>r &gt;&gt;= 16;
    <a id="L148"></a>g &gt;&gt;= 16;
    <a id="L149"></a>b &gt;&gt;= 16;
    <a id="L150"></a><span class="comment">// Since Color.RGBA returns a alpha-premultiplied color, we should have r &lt;= a &amp;&amp; g &lt;= a &amp;&amp; b &lt;= a.</span>
    <a id="L151"></a>r = (r * 0xffff) / a;
    <a id="L152"></a>g = (g * 0xffff) / a;
    <a id="L153"></a>b = (b * 0xffff) / a;
    <a id="L154"></a>return NRGBAColor{uint8(r &gt;&gt; 8), uint8(g &gt;&gt; 8), uint8(b &gt;&gt; 8), uint8(a &gt;&gt; 8)};
<a id="L155"></a>}

<a id="L157"></a>func toNRGBA64Color(c Color) Color {
    <a id="L158"></a>if _, ok := c.(NRGBA64Color); ok { <span class="comment">// no-op conversion</span>
        <a id="L159"></a>return c
    <a id="L160"></a>}
    <a id="L161"></a>r, g, b, a := c.RGBA();
    <a id="L162"></a>a &gt;&gt;= 16;
    <a id="L163"></a>r &gt;&gt;= 16;
    <a id="L164"></a>g &gt;&gt;= 16;
    <a id="L165"></a>b &gt;&gt;= 16;
    <a id="L166"></a>if a == 0xffff {
        <a id="L167"></a>return NRGBA64Color{uint16(r), uint16(g), uint16(b), 0xffff}
    <a id="L168"></a>}
    <a id="L169"></a>if a == 0 {
        <a id="L170"></a>return NRGBA64Color{0, 0, 0, 0}
    <a id="L171"></a>}
    <a id="L172"></a><span class="comment">// Since Color.RGBA returns a alpha-premultiplied color, we should have r &lt;= a &amp;&amp; g &lt;= a &amp;&amp; b &lt;= a.</span>
    <a id="L173"></a>r = (r * 0xffff) / a;
    <a id="L174"></a>g = (g * 0xffff) / a;
    <a id="L175"></a>b = (b * 0xffff) / a;
    <a id="L176"></a>return NRGBA64Color{uint16(r), uint16(g), uint16(b), uint16(a)};
<a id="L177"></a>}

<a id="L179"></a><span class="comment">// The ColorModel associated with RGBAColor.</span>
<a id="L180"></a>var RGBAColorModel ColorModel = ColorModelFunc(toRGBAColor)

<a id="L182"></a><span class="comment">// The ColorModel associated with RGBA64Color.</span>
<a id="L183"></a>var RGBA64ColorModel ColorModel = ColorModelFunc(toRGBA64Color)

<a id="L185"></a><span class="comment">// The ColorModel associated with NRGBAColor.</span>
<a id="L186"></a>var NRGBAColorModel ColorModel = ColorModelFunc(toNRGBAColor)

<a id="L188"></a><span class="comment">// The ColorModel associated with NRGBA64Color.</span>
<a id="L189"></a>var NRGBA64ColorModel ColorModel = ColorModelFunc(toNRGBA64Color)
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
