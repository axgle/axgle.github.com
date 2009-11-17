<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/image/png/reader_test.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/image/png/reader_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package png

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bufio&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
    <a id="L10"></a>&#34;image&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;testing&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// The go PNG library currently supports only a subset of the full PNG specification.</span>
<a id="L17"></a><span class="comment">// In particular, bit depths other than 8 are not supported, and neither are grayscale images.</span>
<a id="L18"></a>var filenames = []string{
    <a id="L19"></a><span class="comment">//&#34;basn0g01&#34;,	// bit depth is not 8</span>
    <a id="L20"></a><span class="comment">//&#34;basn0g02&#34;,	// bit depth is not 8</span>
    <a id="L21"></a><span class="comment">//&#34;basn0g04&#34;,	// bit depth is not 8</span>
    <a id="L22"></a><span class="comment">//&#34;basn0g08&#34;,	// grayscale color model</span>
    <a id="L23"></a><span class="comment">//&#34;basn0g16&#34;,	// bit depth is not 8</span>
    <a id="L24"></a>&#34;basn2c08&#34;,
    <a id="L25"></a><span class="comment">//&#34;basn2c16&#34;,	// bit depth is not 8</span>
    <a id="L26"></a><span class="comment">//&#34;basn3p01&#34;,	// bit depth is not 8</span>
    <a id="L27"></a><span class="comment">//&#34;basn3p02&#34;,	// bit depth is not 8</span>
    <a id="L28"></a><span class="comment">//&#34;basn3p04&#34;,	// bit depth is not 8</span>
    <a id="L29"></a>&#34;basn3p08&#34;,
    <a id="L30"></a><span class="comment">//&#34;basn4a08&#34;,	// grayscale color model</span>
    <a id="L31"></a><span class="comment">//&#34;basn4a16&#34;,	// bit depth is not 8</span>
    <a id="L32"></a>&#34;basn6a08&#34;,
    <a id="L33"></a><span class="comment">//&#34;basn6a16&#34;,	// bit depth is not 8</span>
<a id="L34"></a>}

<a id="L36"></a>func readPng(filename string) (image.Image, os.Error) {
    <a id="L37"></a>f, err := os.Open(filename, os.O_RDONLY, 0444);
    <a id="L38"></a>if err != nil {
        <a id="L39"></a>return nil, err
    <a id="L40"></a>}
    <a id="L41"></a>defer f.Close();
    <a id="L42"></a>return Decode(f);
<a id="L43"></a>}

<a id="L45"></a><span class="comment">// An approximation of the sng command-line tool.</span>
<a id="L46"></a>func sng(w io.WriteCloser, filename string, png image.Image) {
    <a id="L47"></a>defer w.Close();
    <a id="L48"></a><span class="comment">// For now, the go PNG parser only reads bitdepths of 8.</span>
    <a id="L49"></a>bitdepth := 8;

    <a id="L51"></a><span class="comment">// Write the filename and IHDR.</span>
    <a id="L52"></a>io.WriteString(w, &#34;#SNG: from &#34;+filename+&#34;.png\nIHDR {\n&#34;);
    <a id="L53"></a>fmt.Fprintf(w, &#34;    width: %d; height: %d; bitdepth: %d;\n&#34;, png.Width(), png.Height(), bitdepth);
    <a id="L54"></a>cm := png.ColorModel();
    <a id="L55"></a>var paletted *image.Paletted;
    <a id="L56"></a>cpm, _ := cm.(image.PalettedColorModel);
    <a id="L57"></a>switch {
    <a id="L58"></a>case cm == image.RGBAColorModel:
        <a id="L59"></a>io.WriteString(w, &#34;    using color;\n&#34;)
    <a id="L60"></a>case cm == image.NRGBAColorModel:
        <a id="L61"></a>io.WriteString(w, &#34;    using color alpha;\n&#34;)
    <a id="L62"></a>case cpm != nil:
        <a id="L63"></a>io.WriteString(w, &#34;    using color palette;\n&#34;);
        <a id="L64"></a>paletted = png.(*image.Paletted);
    <a id="L65"></a>default:
        <a id="L66"></a>io.WriteString(w, &#34;unknown PNG decoder color model\n&#34;)
    <a id="L67"></a>}
    <a id="L68"></a>io.WriteString(w, &#34;}\n&#34;);

    <a id="L70"></a><span class="comment">// We fake a gAMA output. The test files have a gAMA chunk but the go PNG parser ignores it</span>
    <a id="L71"></a><span class="comment">// (the PNG spec section 11.3 says &#34;Ancillary chunks may be ignored by a decoder&#34;).</span>
    <a id="L72"></a>io.WriteString(w, &#34;gAMA {1.0000}\n&#34;);

    <a id="L74"></a><span class="comment">// Write the PLTE (if applicable).</span>
    <a id="L75"></a>if cpm != nil {
        <a id="L76"></a>io.WriteString(w, &#34;PLTE {\n&#34;);
        <a id="L77"></a>for i := 0; i &lt; len(cpm); i++ {
            <a id="L78"></a>r, g, b, _ := cpm[i].RGBA();
            <a id="L79"></a>r &gt;&gt;= 24;
            <a id="L80"></a>g &gt;&gt;= 24;
            <a id="L81"></a>b &gt;&gt;= 24;
            <a id="L82"></a>fmt.Fprintf(w, &#34;    (%3d,%3d,%3d)     # rgb = (0x%02x,0x%02x,0x%02x)\n&#34;, r, g, b, r, g, b);
        <a id="L83"></a>}
        <a id="L84"></a>io.WriteString(w, &#34;}\n&#34;);
    <a id="L85"></a>}

    <a id="L87"></a><span class="comment">// Write the IMAGE.</span>
    <a id="L88"></a>io.WriteString(w, &#34;IMAGE {\n    pixels hex\n&#34;);
    <a id="L89"></a>for y := 0; y &lt; png.Height(); y++ {
        <a id="L90"></a>switch {
        <a id="L91"></a>case cm == image.RGBAColorModel:
            <a id="L92"></a>for x := 0; x &lt; png.Width(); x++ {
                <a id="L93"></a>rgba := png.At(x, y).(image.RGBAColor);
                <a id="L94"></a>fmt.Fprintf(w, &#34;%02x%02x%02x &#34;, rgba.R, rgba.G, rgba.B);
            <a id="L95"></a>}
        <a id="L96"></a>case cm == image.NRGBAColorModel:
            <a id="L97"></a>for x := 0; x &lt; png.Width(); x++ {
                <a id="L98"></a>nrgba := png.At(x, y).(image.NRGBAColor);
                <a id="L99"></a>fmt.Fprintf(w, &#34;%02x%02x%02x%02x &#34;, nrgba.R, nrgba.G, nrgba.B, nrgba.A);
            <a id="L100"></a>}
        <a id="L101"></a>case cpm != nil:
            <a id="L102"></a>for x := 0; x &lt; png.Width(); x++ {
                <a id="L103"></a>fmt.Fprintf(w, &#34;%02x&#34;, paletted.ColorIndexAt(x, y))
            <a id="L104"></a>}
        <a id="L105"></a>}
        <a id="L106"></a>io.WriteString(w, &#34;\n&#34;);
    <a id="L107"></a>}
    <a id="L108"></a>io.WriteString(w, &#34;}\n&#34;);
<a id="L109"></a>}

<a id="L111"></a>func TestReader(t *testing.T) {
    <a id="L112"></a>for _, fn := range filenames {
        <a id="L113"></a><span class="comment">// Read the .png file.</span>
        <a id="L114"></a>image, err := readPng(&#34;testdata/pngsuite/&#34; + fn + &#34;.png&#34;);
        <a id="L115"></a>if err != nil {
            <a id="L116"></a>t.Error(fn, err);
            <a id="L117"></a>continue;
        <a id="L118"></a>}
        <a id="L119"></a>piper, pipew := io.Pipe();
        <a id="L120"></a>pb := bufio.NewReader(piper);
        <a id="L121"></a>go sng(pipew, fn, image);
        <a id="L122"></a>defer piper.Close();

        <a id="L124"></a><span class="comment">// Read the .sng file.</span>
        <a id="L125"></a>sf, err := os.Open(&#34;testdata/pngsuite/&#34;+fn+&#34;.sng&#34;, os.O_RDONLY, 0444);
        <a id="L126"></a>if err != nil {
            <a id="L127"></a>t.Error(fn, err);
            <a id="L128"></a>continue;
        <a id="L129"></a>}
        <a id="L130"></a>defer sf.Close();
        <a id="L131"></a>sb := bufio.NewReader(sf);
        <a id="L132"></a>if err != nil {
            <a id="L133"></a>t.Error(fn, err);
            <a id="L134"></a>continue;
        <a id="L135"></a>}

        <a id="L137"></a><span class="comment">// Compare the two, in SNG format, line by line.</span>
        <a id="L138"></a>for {
            <a id="L139"></a>ps, perr := pb.ReadString(&#39;\n&#39;);
            <a id="L140"></a>ss, serr := sb.ReadString(&#39;\n&#39;);
            <a id="L141"></a>if perr == os.EOF &amp;&amp; serr == os.EOF {
                <a id="L142"></a>break
            <a id="L143"></a>}
            <a id="L144"></a>if perr != nil {
                <a id="L145"></a>t.Error(fn, perr);
                <a id="L146"></a>break;
            <a id="L147"></a>}
            <a id="L148"></a>if serr != nil {
                <a id="L149"></a>t.Error(fn, serr);
                <a id="L150"></a>break;
            <a id="L151"></a>}
            <a id="L152"></a>if ps != ss {
                <a id="L153"></a>t.Errorf(&#34;%s: Mismatch\n%sversus\n%s\n&#34;, fn, ps, ss);
                <a id="L154"></a>break;
            <a id="L155"></a>}
        <a id="L156"></a>}
    <a id="L157"></a>}
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
