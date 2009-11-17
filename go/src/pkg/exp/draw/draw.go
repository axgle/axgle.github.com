<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/draw/draw.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/draw/draw.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package draw provides basic graphics and drawing primitives,</span>
<a id="L6"></a><span class="comment">// in the style of the Plan 9 graphics library</span>
<a id="L7"></a><span class="comment">// (see http://plan9.bell-labs.com/magic/man2html/2/draw)</span>
<a id="L8"></a><span class="comment">// and the X Render extension.</span>
<a id="L9"></a>package draw

<a id="L11"></a><span class="comment">// BUG(rsc): This is a toy library and not ready for production use.</span>

<a id="L13"></a>import &#34;image&#34;

<a id="L15"></a><span class="comment">// A draw.Image is an image.Image with a Set method to change a single pixel.</span>
<a id="L16"></a>type Image interface {
    <a id="L17"></a>image.Image;
    <a id="L18"></a>Set(x, y int, c image.Color);
<a id="L19"></a>}

<a id="L21"></a><span class="comment">// Draw aligns r.Min in dst with pt in src and mask</span>
<a id="L22"></a><span class="comment">// and then replaces the rectangle r in dst with the</span>
<a id="L23"></a><span class="comment">// result of the Porter-Duff compositing operation</span>
<a id="L24"></a><span class="comment">// ``(src in mask) over dst.&#39;&#39;  If mask is nil, the operation</span>
<a id="L25"></a><span class="comment">// simplifies to ``src over dst.&#39;&#39;</span>
<a id="L26"></a><span class="comment">// The implementation is simple and slow.</span>
<a id="L27"></a>func Draw(dst Image, r Rectangle, src, mask image.Image, pt Point) {
    <a id="L28"></a><span class="comment">// Plenty of room for optimizations here.</span>

    <a id="L30"></a>dx, dy := src.Width(), src.Height();
    <a id="L31"></a>if mask != nil {
        <a id="L32"></a>if dx &gt; mask.Width() {
            <a id="L33"></a>dx = mask.Width()
        <a id="L34"></a>}
        <a id="L35"></a>if dy &gt; mask.Width() {
            <a id="L36"></a>dy = mask.Width()
        <a id="L37"></a>}
    <a id="L38"></a>}
    <a id="L39"></a>dx -= pt.X;
    <a id="L40"></a>dy -= pt.Y;
    <a id="L41"></a>if r.Dx() &gt; dx {
        <a id="L42"></a>r.Max.X = r.Min.X + dx
    <a id="L43"></a>}
    <a id="L44"></a>if r.Dy() &gt; dy {
        <a id="L45"></a>r.Max.Y = r.Min.Y + dy
    <a id="L46"></a>}

    <a id="L48"></a>x0, x1, dx := r.Min.X, r.Max.X, 1;
    <a id="L49"></a>y0, y1, dy := r.Min.Y, r.Max.Y, 1;
    <a id="L50"></a>if image.Image(dst) == src &amp;&amp; r.Overlaps(r.Add(pt.Sub(r.Min))) {
        <a id="L51"></a><span class="comment">// Rectangles overlap: process backward?</span>
        <a id="L52"></a>if pt.Y &lt; r.Min.Y || pt.Y == r.Min.Y &amp;&amp; pt.X &lt; r.Min.X {
            <a id="L53"></a>x0, x1, dx = x1-1, x0-1, -1;
            <a id="L54"></a>y0, y1, dy = y1-1, y0-1, -1;
        <a id="L55"></a>}
    <a id="L56"></a>}

    <a id="L58"></a>var out *image.RGBA64Color;
    <a id="L59"></a>for y := y0; y != y1; y += dy {
        <a id="L60"></a>for x := x0; x != x1; x += dx {
            <a id="L61"></a>sx := pt.X + x - r.Min.X;
            <a id="L62"></a>sy := pt.Y + y - r.Min.Y;
            <a id="L63"></a>if mask == nil {
                <a id="L64"></a>dst.Set(x, y, src.At(sx, sy));
                <a id="L65"></a>continue;
            <a id="L66"></a>}
            <a id="L67"></a>_, _, _, ma := mask.At(sx, sy).RGBA();
            <a id="L68"></a>switch ma {
            <a id="L69"></a>case 0:
                <a id="L70"></a>continue
            <a id="L71"></a>case 0xFFFFFFFF:
                <a id="L72"></a>dst.Set(x, y, src.At(sx, sy))
            <a id="L73"></a>default:
                <a id="L74"></a>dr, dg, db, da := dst.At(x, y).RGBA();
                <a id="L75"></a>dr &gt;&gt;= 16;
                <a id="L76"></a>dg &gt;&gt;= 16;
                <a id="L77"></a>db &gt;&gt;= 16;
                <a id="L78"></a>da &gt;&gt;= 16;
                <a id="L79"></a>sr, sg, sb, sa := src.At(sx, sy).RGBA();
                <a id="L80"></a>sr &gt;&gt;= 16;
                <a id="L81"></a>sg &gt;&gt;= 16;
                <a id="L82"></a>sb &gt;&gt;= 16;
                <a id="L83"></a>sa &gt;&gt;= 16;
                <a id="L84"></a>ma &gt;&gt;= 16;
                <a id="L85"></a>const M = 1&lt;&lt;16 - 1;
                <a id="L86"></a>a := sa * ma / M;
                <a id="L87"></a>dr = (dr*(M-a) + sr*ma) / M;
                <a id="L88"></a>dg = (dg*(M-a) + sg*ma) / M;
                <a id="L89"></a>db = (db*(M-a) + sb*ma) / M;
                <a id="L90"></a>da = (da*(M-a) + sa*ma) / M;
                <a id="L91"></a>if out == nil {
                    <a id="L92"></a>out = new(image.RGBA64Color)
                <a id="L93"></a>}
                <a id="L94"></a>out.R = uint16(dr);
                <a id="L95"></a>out.G = uint16(dg);
                <a id="L96"></a>out.B = uint16(db);
                <a id="L97"></a>out.A = uint16(da);
                <a id="L98"></a>dst.Set(x, y, out);
            <a id="L99"></a>}
        <a id="L100"></a>}
    <a id="L101"></a>}
<a id="L102"></a>}

<a id="L104"></a><span class="comment">// Border aligns r.Min in dst with sp in src and then replaces pixels</span>
<a id="L105"></a><span class="comment">// in a w-pixel border around r in dst with the result of the Porter-Duff compositing</span>
<a id="L106"></a><span class="comment">// operation ``src over dst.&#39;&#39;  If w is positive, the border extends w pixels inside r.</span>
<a id="L107"></a><span class="comment">// If w is negative, the border extends w pixels outside r.</span>
<a id="L108"></a>func Border(dst Image, r Rectangle, w int, src image.Image, sp Point) {
    <a id="L109"></a>i := w;
    <a id="L110"></a>if i &gt; 0 {
        <a id="L111"></a><span class="comment">// inside r</span>
        <a id="L112"></a>Draw(dst, Rect(r.Min.X, r.Min.Y, r.Max.X, r.Min.Y+i), src, nil, sp);                          <span class="comment">// top</span>
        <a id="L113"></a>Draw(dst, Rect(r.Min.X, r.Min.Y+i, r.Min.X+i, r.Max.Y-i), src, nil, sp.Add(Pt(0, i)));        <span class="comment">// left</span>
        <a id="L114"></a>Draw(dst, Rect(r.Max.X-i, r.Min.Y+i, r.Max.X, r.Max.Y-i), src, nil, sp.Add(Pt(r.Dx()-i, i))); <span class="comment">// right</span>
        <a id="L115"></a>Draw(dst, Rect(r.Min.X, r.Max.Y-i, r.Max.X, r.Max.Y), src, nil, sp.Add(Pt(0, r.Dy()-i)));     <span class="comment">// bottom</span>
        <a id="L116"></a>return;
    <a id="L117"></a>}

    <a id="L119"></a><span class="comment">// outside r;</span>
    <a id="L120"></a>i = -i;
    <a id="L121"></a>Draw(dst, Rect(r.Min.X-i, r.Min.Y-i, r.Max.X+i, r.Min.Y), src, nil, sp.Add(Pt(-i, -i))); <span class="comment">// top</span>
    <a id="L122"></a>Draw(dst, Rect(r.Min.X-i, r.Min.Y, r.Min.X, r.Max.Y), src, nil, sp.Add(Pt(-i, 0)));      <span class="comment">// left</span>
    <a id="L123"></a>Draw(dst, Rect(r.Max.X, r.Min.Y, r.Max.X+i, r.Max.Y), src, nil, sp.Add(Pt(r.Dx(), 0)));  <span class="comment">// right</span>
    <a id="L124"></a>Draw(dst, Rect(r.Min.X-i, r.Max.Y, r.Max.X+i, r.Max.Y+i), src, nil, sp.Add(Pt(-i, 0)));  <span class="comment">// bottom</span>
<a id="L125"></a>}
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
