<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/image/png/reader.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/image/png/reader.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The png package implements a PNG image decoder and encoder.</span>
<a id="L6"></a><span class="comment">//</span>
<a id="L7"></a><span class="comment">// The PNG specification is at http://www.libpng.org/pub/png/spec/1.2/PNG-Contents.html</span>
<a id="L8"></a>package png

<a id="L10"></a>import (
    <a id="L11"></a>&#34;compress/zlib&#34;;
    <a id="L12"></a>&#34;hash&#34;;
    <a id="L13"></a>&#34;hash/crc32&#34;;
    <a id="L14"></a>&#34;image&#34;;
    <a id="L15"></a>&#34;io&#34;;
    <a id="L16"></a>&#34;os&#34;;
<a id="L17"></a>)

<a id="L19"></a><span class="comment">// Color type, as per the PNG spec.</span>
<a id="L20"></a>const (
    <a id="L21"></a>ctGrayscale      = 0;
    <a id="L22"></a>ctTrueColor      = 2;
    <a id="L23"></a>ctPaletted       = 3;
    <a id="L24"></a>ctGrayscaleAlpha = 4;
    <a id="L25"></a>ctTrueColorAlpha = 6;
<a id="L26"></a>)

<a id="L28"></a><span class="comment">// Filter type, as per the PNG spec.</span>
<a id="L29"></a>const (
    <a id="L30"></a>ftNone    = 0;
    <a id="L31"></a>ftSub     = 1;
    <a id="L32"></a>ftUp      = 2;
    <a id="L33"></a>ftAverage = 3;
    <a id="L34"></a>ftPaeth   = 4;
    <a id="L35"></a>nFilter   = 5;
<a id="L36"></a>)

<a id="L38"></a><span class="comment">// Decoding stage.</span>
<a id="L39"></a><span class="comment">// The PNG specification says that the IHDR, PLTE (if present), IDAT and IEND</span>
<a id="L40"></a><span class="comment">// chunks must appear in that order. There may be multiple IDAT chunks, and</span>
<a id="L41"></a><span class="comment">// IDAT chunks must be sequential (i.e. they may not have any other chunks</span>
<a id="L42"></a><span class="comment">// between them).</span>
<a id="L43"></a>const (
    <a id="L44"></a>dsStart = iota;
    <a id="L45"></a>dsSeenIHDR;
    <a id="L46"></a>dsSeenPLTE;
    <a id="L47"></a>dsSeenIDAT;
    <a id="L48"></a>dsSeenIEND;
<a id="L49"></a>)

<a id="L51"></a>const pngHeader = &#34;\x89PNG\r\n\x1a\n&#34;

<a id="L53"></a>type decoder struct {
    <a id="L54"></a>width, height int;
    <a id="L55"></a>image         image.Image;
    <a id="L56"></a>colorType     uint8;
    <a id="L57"></a>stage         int;
    <a id="L58"></a>idatWriter    io.WriteCloser;
    <a id="L59"></a>idatDone      chan os.Error;
    <a id="L60"></a>tmp           [3 * 256]byte;
<a id="L61"></a>}

<a id="L63"></a><span class="comment">// A FormatError reports that the input is not a valid PNG.</span>
<a id="L64"></a>type FormatError string

<a id="L66"></a>func (e FormatError) String() string { return &#34;invalid PNG format: &#34; + string(e) }

<a id="L68"></a>var chunkOrderError = FormatError(&#34;chunk out of order&#34;)

<a id="L70"></a><span class="comment">// An IDATDecodingError wraps an inner error (such as a ZLIB decoding error) encountered while processing an IDAT chunk.</span>
<a id="L71"></a>type IDATDecodingError struct {
    <a id="L72"></a>Err os.Error;
<a id="L73"></a>}

<a id="L75"></a>func (e IDATDecodingError) String() string { return &#34;IDAT decoding error: &#34; + e.Err.String() }

<a id="L77"></a><span class="comment">// An UnsupportedError reports that the input uses a valid but unimplemented PNG feature.</span>
<a id="L78"></a>type UnsupportedError string

<a id="L80"></a>func (e UnsupportedError) String() string { return &#34;unsupported PNG feature: &#34; + string(e) }

<a id="L82"></a><span class="comment">// Big-endian.</span>
<a id="L83"></a>func parseUint32(b []uint8) uint32 {
    <a id="L84"></a>return uint32(b[0])&lt;&lt;24 | uint32(b[1])&lt;&lt;16 | uint32(b[2])&lt;&lt;8 | uint32(b[3])
<a id="L85"></a>}

<a id="L87"></a>func abs(x int) int {
    <a id="L88"></a>if x &lt; 0 {
        <a id="L89"></a>return -x
    <a id="L90"></a>}
    <a id="L91"></a>return x;
<a id="L92"></a>}

<a id="L94"></a>func min(a, b int) int {
    <a id="L95"></a>if a &lt; b {
        <a id="L96"></a>return a
    <a id="L97"></a>}
    <a id="L98"></a>return b;
<a id="L99"></a>}

<a id="L101"></a>func (d *decoder) parseIHDR(r io.Reader, crc hash.Hash32, length uint32) os.Error {
    <a id="L102"></a>if length != 13 {
        <a id="L103"></a>return FormatError(&#34;bad IHDR length&#34;)
    <a id="L104"></a>}
    <a id="L105"></a>_, err := io.ReadFull(r, d.tmp[0:13]);
    <a id="L106"></a>if err != nil {
        <a id="L107"></a>return err
    <a id="L108"></a>}
    <a id="L109"></a>crc.Write(d.tmp[0:13]);
    <a id="L110"></a>if d.tmp[8] != 8 {
        <a id="L111"></a>return UnsupportedError(&#34;bit depth&#34;)
    <a id="L112"></a>}
    <a id="L113"></a>if d.tmp[10] != 0 || d.tmp[11] != 0 || d.tmp[12] != 0 {
        <a id="L114"></a>return UnsupportedError(&#34;compression, filter or interlace method&#34;)
    <a id="L115"></a>}
    <a id="L116"></a>w := int32(parseUint32(d.tmp[0:4]));
    <a id="L117"></a>h := int32(parseUint32(d.tmp[4:8]));
    <a id="L118"></a>if w &lt; 0 || h &lt; 0 {
        <a id="L119"></a>return FormatError(&#34;negative dimension&#34;)
    <a id="L120"></a>}
    <a id="L121"></a>nPixels := int64(w) * int64(h);
    <a id="L122"></a>if nPixels != int64(int(nPixels)) {
        <a id="L123"></a>return UnsupportedError(&#34;dimension overflow&#34;)
    <a id="L124"></a>}
    <a id="L125"></a>d.colorType = d.tmp[9];
    <a id="L126"></a>switch d.colorType {
    <a id="L127"></a>case ctTrueColor:
        <a id="L128"></a>d.image = image.NewRGBA(int(w), int(h))
    <a id="L129"></a>case ctPaletted:
        <a id="L130"></a>d.image = image.NewPaletted(int(w), int(h), nil)
    <a id="L131"></a>case ctTrueColorAlpha:
        <a id="L132"></a>d.image = image.NewNRGBA(int(w), int(h))
    <a id="L133"></a>default:
        <a id="L134"></a>return UnsupportedError(&#34;color type&#34;)
    <a id="L135"></a>}
    <a id="L136"></a>d.width, d.height = int(w), int(h);
    <a id="L137"></a>return nil;
<a id="L138"></a>}

<a id="L140"></a>func (d *decoder) parsePLTE(r io.Reader, crc hash.Hash32, length uint32) os.Error {
    <a id="L141"></a>np := int(length / 3); <span class="comment">// The number of palette entries.</span>
    <a id="L142"></a>if length%3 != 0 || np &lt;= 0 || np &gt; 256 {
        <a id="L143"></a>return FormatError(&#34;bad PLTE length&#34;)
    <a id="L144"></a>}
    <a id="L145"></a>n, err := io.ReadFull(r, d.tmp[0:3*np]);
    <a id="L146"></a>if err != nil {
        <a id="L147"></a>return err
    <a id="L148"></a>}
    <a id="L149"></a>crc.Write(d.tmp[0:n]);
    <a id="L150"></a>switch d.colorType {
    <a id="L151"></a>case ctPaletted:
        <a id="L152"></a>palette := make([]image.Color, np);
        <a id="L153"></a>for i := 0; i &lt; np; i++ {
            <a id="L154"></a>palette[i] = image.RGBAColor{d.tmp[3*i+0], d.tmp[3*i+1], d.tmp[3*i+2], 0xff}
        <a id="L155"></a>}
        <a id="L156"></a>d.image.(*image.Paletted).Palette = image.PalettedColorModel(palette);
    <a id="L157"></a>case ctTrueColor, ctTrueColorAlpha:
        <a id="L158"></a><span class="comment">// As per the PNG spec, a PLTE chunk is optional (and for practical purposes,</span>
        <a id="L159"></a><span class="comment">// ignorable) for the ctTrueColor and ctTrueColorAlpha color types (section 4.1.2).</span>
        <a id="L160"></a>return nil
    <a id="L161"></a>default:
        <a id="L162"></a>return FormatError(&#34;PLTE, color type mismatch&#34;)
    <a id="L163"></a>}
    <a id="L164"></a>return nil;
<a id="L165"></a>}

<a id="L167"></a><span class="comment">// The Paeth filter function, as per the PNG specification.</span>
<a id="L168"></a>func paeth(a, b, c uint8) uint8 {
    <a id="L169"></a>p := int(a) + int(b) - int(c);
    <a id="L170"></a>pa := abs(p - int(a));
    <a id="L171"></a>pb := abs(p - int(b));
    <a id="L172"></a>pc := abs(p - int(c));
    <a id="L173"></a>if pa &lt;= pb &amp;&amp; pa &lt;= pc {
        <a id="L174"></a>return a
    <a id="L175"></a>} else if pb &lt;= pc {
        <a id="L176"></a>return b
    <a id="L177"></a>}
    <a id="L178"></a>return c;
<a id="L179"></a>}

<a id="L181"></a>func (d *decoder) idatReader(idat io.Reader) os.Error {
    <a id="L182"></a>r, err := zlib.NewInflater(idat);
    <a id="L183"></a>if err != nil {
        <a id="L184"></a>return err
    <a id="L185"></a>}
    <a id="L186"></a>defer r.Close();
    <a id="L187"></a>bpp := 0; <span class="comment">// Bytes per pixel.</span>
    <a id="L188"></a>maxPalette := uint8(0);
    <a id="L189"></a>var (
        <a id="L190"></a>rgba     *image.RGBA;
        <a id="L191"></a>nrgba    *image.NRGBA;
        <a id="L192"></a>paletted *image.Paletted;
    <a id="L193"></a>)
    <a id="L194"></a>switch d.colorType {
    <a id="L195"></a>case ctTrueColor:
        <a id="L196"></a>bpp = 3;
        <a id="L197"></a>rgba = d.image.(*image.RGBA);
    <a id="L198"></a>case ctPaletted:
        <a id="L199"></a>bpp = 1;
        <a id="L200"></a>paletted = d.image.(*image.Paletted);
        <a id="L201"></a>maxPalette = uint8(len(paletted.Palette) - 1);
    <a id="L202"></a>case ctTrueColorAlpha:
        <a id="L203"></a>bpp = 4;
        <a id="L204"></a>nrgba = d.image.(*image.NRGBA);
    <a id="L205"></a>}
    <a id="L206"></a><span class="comment">// cr and pr are the bytes for the current and previous row.</span>
    <a id="L207"></a><span class="comment">// The +1 is for the per-row filter type, which is at cr[0].</span>
    <a id="L208"></a>cr := make([]uint8, 1+bpp*d.width);
    <a id="L209"></a>pr := make([]uint8, 1+bpp*d.width);

    <a id="L211"></a>for y := 0; y &lt; d.height; y++ {
        <a id="L212"></a><span class="comment">// Read the decompressed bytes.</span>
        <a id="L213"></a>_, err := io.ReadFull(r, cr);
        <a id="L214"></a>if err != nil {
            <a id="L215"></a>return err
        <a id="L216"></a>}

        <a id="L218"></a><span class="comment">// Apply the filter.</span>
        <a id="L219"></a>cdat := cr[1:len(cr)];
        <a id="L220"></a>pdat := pr[1:len(pr)];
        <a id="L221"></a>switch cr[0] {
        <a id="L222"></a>case ftNone:
            <a id="L223"></a><span class="comment">// No-op.</span>
        <a id="L224"></a>case ftSub:
            <a id="L225"></a>for i := bpp; i &lt; len(cdat); i++ {
                <a id="L226"></a>cdat[i] += cdat[i-bpp]
            <a id="L227"></a>}
        <a id="L228"></a>case ftUp:
            <a id="L229"></a>for i := 0; i &lt; len(cdat); i++ {
                <a id="L230"></a>cdat[i] += pdat[i]
            <a id="L231"></a>}
        <a id="L232"></a>case ftAverage:
            <a id="L233"></a>for i := 0; i &lt; bpp; i++ {
                <a id="L234"></a>cdat[i] += pdat[i] / 2
            <a id="L235"></a>}
            <a id="L236"></a>for i := bpp; i &lt; len(cdat); i++ {
                <a id="L237"></a>cdat[i] += uint8((int(cdat[i-bpp]) + int(pdat[i])) / 2)
            <a id="L238"></a>}
        <a id="L239"></a>case ftPaeth:
            <a id="L240"></a>for i := 0; i &lt; bpp; i++ {
                <a id="L241"></a>cdat[i] += paeth(0, pdat[i], 0)
            <a id="L242"></a>}
            <a id="L243"></a>for i := bpp; i &lt; len(cdat); i++ {
                <a id="L244"></a>cdat[i] += paeth(cdat[i-bpp], pdat[i], pdat[i-bpp])
            <a id="L245"></a>}
        <a id="L246"></a>default:
            <a id="L247"></a>return FormatError(&#34;bad filter type&#34;)
        <a id="L248"></a>}

        <a id="L250"></a><span class="comment">// Convert from bytes to colors.</span>
        <a id="L251"></a>switch d.colorType {
        <a id="L252"></a>case ctTrueColor:
            <a id="L253"></a>for x := 0; x &lt; d.width; x++ {
                <a id="L254"></a>rgba.Set(x, y, image.RGBAColor{cdat[3*x+0], cdat[3*x+1], cdat[3*x+2], 0xff})
            <a id="L255"></a>}
        <a id="L256"></a>case ctPaletted:
            <a id="L257"></a>for x := 0; x &lt; d.width; x++ {
                <a id="L258"></a>if cdat[x] &gt; maxPalette {
                    <a id="L259"></a>return FormatError(&#34;palette index out of range&#34;)
                <a id="L260"></a>}
                <a id="L261"></a>paletted.SetColorIndex(x, y, cdat[x]);
            <a id="L262"></a>}
        <a id="L263"></a>case ctTrueColorAlpha:
            <a id="L264"></a>for x := 0; x &lt; d.width; x++ {
                <a id="L265"></a>nrgba.Set(x, y, image.NRGBAColor{cdat[4*x+0], cdat[4*x+1], cdat[4*x+2], cdat[4*x+3]})
            <a id="L266"></a>}
        <a id="L267"></a>}

        <a id="L269"></a><span class="comment">// The current row for y is the previous row for y+1.</span>
        <a id="L270"></a>pr, cr = cr, pr;
    <a id="L271"></a>}
    <a id="L272"></a>return nil;
<a id="L273"></a>}

<a id="L275"></a>func (d *decoder) parseIDAT(r io.Reader, crc hash.Hash32, length uint32) os.Error {
    <a id="L276"></a><span class="comment">// There may be more than one IDAT chunk, but their contents must be</span>
    <a id="L277"></a><span class="comment">// treated as if it was one continuous stream (to the zlib decoder).</span>
    <a id="L278"></a><span class="comment">// We bring up an io.Pipe and write the IDAT chunks into the pipe as</span>
    <a id="L279"></a><span class="comment">// we see them, and decode the stream in a separate go-routine, which</span>
    <a id="L280"></a><span class="comment">// signals its completion (successful or not) via a channel.</span>
    <a id="L281"></a>if d.idatWriter == nil {
        <a id="L282"></a>pr, pw := io.Pipe();
        <a id="L283"></a>d.idatWriter = pw;
        <a id="L284"></a>d.idatDone = make(chan os.Error);
        <a id="L285"></a>go func() {
            <a id="L286"></a>err := d.idatReader(pr);
            <a id="L287"></a>if err == os.EOF {
                <a id="L288"></a>err = FormatError(&#34;too little IDAT&#34;)
            <a id="L289"></a>}
            <a id="L290"></a>pr.CloseWithError(FormatError(&#34;too much IDAT&#34;));
            <a id="L291"></a>d.idatDone &lt;- err;
        <a id="L292"></a>}();
    <a id="L293"></a>}
    <a id="L294"></a>var buf [4096]byte;
    <a id="L295"></a>for length &gt; 0 {
        <a id="L296"></a>n, err1 := r.Read(buf[0:min(len(buf), int(length))]);
        <a id="L297"></a><span class="comment">// We delay checking err1. It is possible to get n bytes and an error,</span>
        <a id="L298"></a><span class="comment">// but if the n bytes themselves contain a FormatError, for example, we</span>
        <a id="L299"></a><span class="comment">// want to report that error, and not the one that made the Read stop.</span>
        <a id="L300"></a>n, err2 := d.idatWriter.Write(buf[0:n]);
        <a id="L301"></a>if err2 != nil {
            <a id="L302"></a>return err2
        <a id="L303"></a>}
        <a id="L304"></a>if err1 != nil {
            <a id="L305"></a>return err1
        <a id="L306"></a>}
        <a id="L307"></a>crc.Write(buf[0:n]);
        <a id="L308"></a>length -= uint32(n);
    <a id="L309"></a>}
    <a id="L310"></a>return nil;
<a id="L311"></a>}

<a id="L313"></a>func (d *decoder) parseIEND(r io.Reader, crc hash.Hash32, length uint32) os.Error {
    <a id="L314"></a>if length != 0 {
        <a id="L315"></a>return FormatError(&#34;bad IEND length&#34;)
    <a id="L316"></a>}
    <a id="L317"></a>return nil;
<a id="L318"></a>}

<a id="L320"></a>func (d *decoder) parseChunk(r io.Reader) os.Error {
    <a id="L321"></a><span class="comment">// Read the length.</span>
    <a id="L322"></a>n, err := io.ReadFull(r, d.tmp[0:4]);
    <a id="L323"></a>if err == os.EOF {
        <a id="L324"></a>return io.ErrUnexpectedEOF
    <a id="L325"></a>}
    <a id="L326"></a>if err != nil {
        <a id="L327"></a>return err
    <a id="L328"></a>}
    <a id="L329"></a>length := parseUint32(d.tmp[0:4]);

    <a id="L331"></a><span class="comment">// Read the chunk type.</span>
    <a id="L332"></a>n, err = io.ReadFull(r, d.tmp[0:4]);
    <a id="L333"></a>if err == os.EOF {
        <a id="L334"></a>return io.ErrUnexpectedEOF
    <a id="L335"></a>}
    <a id="L336"></a>if err != nil {
        <a id="L337"></a>return err
    <a id="L338"></a>}
    <a id="L339"></a>crc := crc32.NewIEEE();
    <a id="L340"></a>crc.Write(d.tmp[0:4]);

    <a id="L342"></a><span class="comment">// Read the chunk data.</span>
    <a id="L343"></a>switch string(d.tmp[0:4]) {
    <a id="L344"></a>case &#34;IHDR&#34;:
        <a id="L345"></a>if d.stage != dsStart {
            <a id="L346"></a>return chunkOrderError
        <a id="L347"></a>}
        <a id="L348"></a>d.stage = dsSeenIHDR;
        <a id="L349"></a>err = d.parseIHDR(r, crc, length);
    <a id="L350"></a>case &#34;PLTE&#34;:
        <a id="L351"></a>if d.stage != dsSeenIHDR {
            <a id="L352"></a>return chunkOrderError
        <a id="L353"></a>}
        <a id="L354"></a>d.stage = dsSeenPLTE;
        <a id="L355"></a>err = d.parsePLTE(r, crc, length);
    <a id="L356"></a>case &#34;IDAT&#34;:
        <a id="L357"></a>if d.stage &lt; dsSeenIHDR || d.stage &gt; dsSeenIDAT || (d.colorType == ctPaletted &amp;&amp; d.stage == dsSeenIHDR) {
            <a id="L358"></a>return chunkOrderError
        <a id="L359"></a>}
        <a id="L360"></a>d.stage = dsSeenIDAT;
        <a id="L361"></a>err = d.parseIDAT(r, crc, length);
    <a id="L362"></a>case &#34;IEND&#34;:
        <a id="L363"></a>if d.stage != dsSeenIDAT {
            <a id="L364"></a>return chunkOrderError
        <a id="L365"></a>}
        <a id="L366"></a>d.stage = dsSeenIEND;
        <a id="L367"></a>err = d.parseIEND(r, crc, length);
    <a id="L368"></a>default:
        <a id="L369"></a><span class="comment">// Ignore this chunk (of a known length).</span>
        <a id="L370"></a>var ignored [4096]byte;
        <a id="L371"></a>for length &gt; 0 {
            <a id="L372"></a>n, err = io.ReadFull(r, ignored[0:min(len(ignored), int(length))]);
            <a id="L373"></a>if err != nil {
                <a id="L374"></a>return err
            <a id="L375"></a>}
            <a id="L376"></a>crc.Write(ignored[0:n]);
            <a id="L377"></a>length -= uint32(n);
        <a id="L378"></a>}
    <a id="L379"></a>}
    <a id="L380"></a>if err != nil {
        <a id="L381"></a>return err
    <a id="L382"></a>}

    <a id="L384"></a><span class="comment">// Read the checksum.</span>
    <a id="L385"></a>n, err = io.ReadFull(r, d.tmp[0:4]);
    <a id="L386"></a>if err == os.EOF {
        <a id="L387"></a>return io.ErrUnexpectedEOF
    <a id="L388"></a>}
    <a id="L389"></a>if err != nil {
        <a id="L390"></a>return err
    <a id="L391"></a>}
    <a id="L392"></a>if parseUint32(d.tmp[0:4]) != crc.Sum32() {
        <a id="L393"></a>return FormatError(&#34;invalid checksum&#34;)
    <a id="L394"></a>}
    <a id="L395"></a>return nil;
<a id="L396"></a>}

<a id="L398"></a>func (d *decoder) checkHeader(r io.Reader) os.Error {
    <a id="L399"></a>_, err := io.ReadFull(r, d.tmp[0:8]);
    <a id="L400"></a>if err != nil {
        <a id="L401"></a>return err
    <a id="L402"></a>}
    <a id="L403"></a>if string(d.tmp[0:8]) != pngHeader {
        <a id="L404"></a>return FormatError(&#34;not a PNG file&#34;)
    <a id="L405"></a>}
    <a id="L406"></a>return nil;
<a id="L407"></a>}

<a id="L409"></a><span class="comment">// Decode reads a PNG formatted image from r and returns it as an image.Image.</span>
<a id="L410"></a><span class="comment">// The type of Image returned depends on the PNG contents.</span>
<a id="L411"></a>func Decode(r io.Reader) (image.Image, os.Error) {
    <a id="L412"></a>var d decoder;
    <a id="L413"></a>err := d.checkHeader(r);
    <a id="L414"></a>if err != nil {
        <a id="L415"></a>return nil, err
    <a id="L416"></a>}
    <a id="L417"></a>for d.stage = dsStart; d.stage != dsSeenIEND; {
        <a id="L418"></a>err = d.parseChunk(r);
        <a id="L419"></a>if err != nil {
            <a id="L420"></a>break
        <a id="L421"></a>}
    <a id="L422"></a>}
    <a id="L423"></a>if d.idatWriter != nil {
        <a id="L424"></a>d.idatWriter.Close();
        <a id="L425"></a>err1 := &lt;-d.idatDone;
        <a id="L426"></a>if err == nil {
            <a id="L427"></a>err = err1
        <a id="L428"></a>}
    <a id="L429"></a>}
    <a id="L430"></a>if err != nil {
        <a id="L431"></a>return nil, err
    <a id="L432"></a>}
    <a id="L433"></a>return d.image, nil;
<a id="L434"></a>}
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
