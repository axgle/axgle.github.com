<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/image/png/writer.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/image/png/writer.go</h1>

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
    <a id="L9"></a>&#34;compress/zlib&#34;;
    <a id="L10"></a>&#34;hash/crc32&#34;;
    <a id="L11"></a>&#34;image&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;os&#34;;
    <a id="L14"></a>&#34;strconv&#34;;
<a id="L15"></a>)

<a id="L17"></a>type encoder struct {
    <a id="L18"></a>w         io.Writer;
    <a id="L19"></a>m         image.Image;
    <a id="L20"></a>colorType uint8;
    <a id="L21"></a>err       os.Error;
    <a id="L22"></a>header    [8]byte;
    <a id="L23"></a>footer    [4]byte;
    <a id="L24"></a>tmp       [3 * 256]byte;
<a id="L25"></a>}

<a id="L27"></a><span class="comment">// Big-endian.</span>
<a id="L28"></a>func writeUint32(b []uint8, u uint32) {
    <a id="L29"></a>b[0] = uint8(u &gt;&gt; 24);
    <a id="L30"></a>b[1] = uint8(u &gt;&gt; 16);
    <a id="L31"></a>b[2] = uint8(u &gt;&gt; 8);
    <a id="L32"></a>b[3] = uint8(u &gt;&gt; 0);
<a id="L33"></a>}

<a id="L35"></a><span class="comment">// Returns whether or not the image is fully opaque.</span>
<a id="L36"></a>func opaque(m image.Image) bool {
    <a id="L37"></a>for y := 0; y &lt; m.Height(); y++ {
        <a id="L38"></a>for x := 0; x &lt; m.Width(); x++ {
            <a id="L39"></a>_, _, _, a := m.At(x, y).RGBA();
            <a id="L40"></a>if a != 0xffffffff {
                <a id="L41"></a>return false
            <a id="L42"></a>}
        <a id="L43"></a>}
    <a id="L44"></a>}
    <a id="L45"></a>return true;
<a id="L46"></a>}

<a id="L48"></a><span class="comment">// The absolute value of a byte interpreted as a signed int8.</span>
<a id="L49"></a>func abs8(d uint8) int {
    <a id="L50"></a>if d &lt; 128 {
        <a id="L51"></a>return int(d)
    <a id="L52"></a>}
    <a id="L53"></a>return 256 - int(d);
<a id="L54"></a>}

<a id="L56"></a>func (e *encoder) writeChunk(b []byte, name string) {
    <a id="L57"></a>if e.err != nil {
        <a id="L58"></a>return
    <a id="L59"></a>}
    <a id="L60"></a>n := uint32(len(b));
    <a id="L61"></a>if int(n) != len(b) {
        <a id="L62"></a>e.err = UnsupportedError(name + &#34; chunk is too large: &#34; + strconv.Itoa(len(b)));
        <a id="L63"></a>return;
    <a id="L64"></a>}
    <a id="L65"></a>writeUint32(e.header[0:4], n);
    <a id="L66"></a>e.header[4] = name[0];
    <a id="L67"></a>e.header[5] = name[1];
    <a id="L68"></a>e.header[6] = name[2];
    <a id="L69"></a>e.header[7] = name[3];
    <a id="L70"></a>crc := crc32.NewIEEE();
    <a id="L71"></a>crc.Write(e.header[4:8]);
    <a id="L72"></a>crc.Write(b);
    <a id="L73"></a>writeUint32(e.footer[0:4], crc.Sum32());

    <a id="L75"></a>_, e.err = e.w.Write(e.header[0:8]);
    <a id="L76"></a>if e.err != nil {
        <a id="L77"></a>return
    <a id="L78"></a>}
    <a id="L79"></a>_, e.err = e.w.Write(b);
    <a id="L80"></a>if e.err != nil {
        <a id="L81"></a>return
    <a id="L82"></a>}
    <a id="L83"></a>_, e.err = e.w.Write(e.footer[0:4]);
<a id="L84"></a>}

<a id="L86"></a>func (e *encoder) writeIHDR() {
    <a id="L87"></a>writeUint32(e.tmp[0:4], uint32(e.m.Width()));
    <a id="L88"></a>writeUint32(e.tmp[4:8], uint32(e.m.Height()));
    <a id="L89"></a>e.tmp[8] = 8; <span class="comment">// bit depth</span>
    <a id="L90"></a>e.tmp[9] = e.colorType;
    <a id="L91"></a>e.tmp[10] = 0; <span class="comment">// default compression method</span>
    <a id="L92"></a>e.tmp[11] = 0; <span class="comment">// default filter method</span>
    <a id="L93"></a>e.tmp[12] = 0; <span class="comment">// non-interlaced</span>
    <a id="L94"></a>e.writeChunk(e.tmp[0:13], &#34;IHDR&#34;);
<a id="L95"></a>}

<a id="L97"></a>func (e *encoder) writePLTE(p image.PalettedColorModel) {
    <a id="L98"></a>if len(p) &lt; 1 || len(p) &gt; 256 {
        <a id="L99"></a>e.err = FormatError(&#34;bad palette length: &#34; + strconv.Itoa(len(p)));
        <a id="L100"></a>return;
    <a id="L101"></a>}
    <a id="L102"></a>for i := 0; i &lt; len(p); i++ {
        <a id="L103"></a>r, g, b, a := p[i].RGBA();
        <a id="L104"></a>if a != 0xffffffff {
            <a id="L105"></a>e.err = UnsupportedError(&#34;non-opaque palette color&#34;);
            <a id="L106"></a>return;
        <a id="L107"></a>}
        <a id="L108"></a>e.tmp[3*i+0] = uint8(r &gt;&gt; 24);
        <a id="L109"></a>e.tmp[3*i+1] = uint8(g &gt;&gt; 24);
        <a id="L110"></a>e.tmp[3*i+2] = uint8(b &gt;&gt; 24);
    <a id="L111"></a>}
    <a id="L112"></a>e.writeChunk(e.tmp[0:3*len(p)], &#34;PLTE&#34;);
<a id="L113"></a>}

<a id="L115"></a><span class="comment">// An encoder is an io.Writer that satisfies writes by writing PNG IDAT chunks,</span>
<a id="L116"></a><span class="comment">// including an 8-byte header and 4-byte CRC checksum per Write call. Such calls</span>
<a id="L117"></a><span class="comment">// should be relatively infrequent, since writeIDATs uses a bufio.Writer.</span>
<a id="L118"></a><span class="comment">//</span>
<a id="L119"></a><span class="comment">// This method should only be called from writeIDATs (via writeImage).</span>
<a id="L120"></a><span class="comment">// No other code should treat an encoder as an io.Writer.</span>
<a id="L121"></a><span class="comment">//</span>
<a id="L122"></a><span class="comment">// Note that, because the zlib deflater may involve an io.Pipe, e.Write calls may</span>
<a id="L123"></a><span class="comment">// occur on a separate go-routine than the e.writeIDATs call, and care should be</span>
<a id="L124"></a><span class="comment">// taken that e&#39;s state (such as its tmp buffer) is not modified concurrently.</span>
<a id="L125"></a>func (e *encoder) Write(b []byte) (int, os.Error) {
    <a id="L126"></a>e.writeChunk(b, &#34;IDAT&#34;);
    <a id="L127"></a>if e.err != nil {
        <a id="L128"></a>return 0, e.err
    <a id="L129"></a>}
    <a id="L130"></a>return len(b), nil;
<a id="L131"></a>}

<a id="L133"></a><span class="comment">// Chooses the filter to use for encoding the current row, and applies it.</span>
<a id="L134"></a><span class="comment">// The return value is the index of the filter and also of the row in cr that has had it applied.</span>
<a id="L135"></a>func filter(cr [][]byte, pr []byte, bpp int) int {
    <a id="L136"></a><span class="comment">// We try all five filter types, and pick the one that minimizes the sum of absolute differences.</span>
    <a id="L137"></a><span class="comment">// This is the same heuristic that libpng uses, although the filters are attempted in order of</span>
    <a id="L138"></a><span class="comment">// estimated most likely to be minimal (ftUp, ftPaeth, ftNone, ftSub, ftAverage), rather than</span>
    <a id="L139"></a><span class="comment">// in their enumeration order (ftNone, ftSub, ftUp, ftAverage, ftPaeth).</span>
    <a id="L140"></a>cdat0 := cr[0][1:len(cr[0])];
    <a id="L141"></a>cdat1 := cr[1][1:len(cr[1])];
    <a id="L142"></a>cdat2 := cr[2][1:len(cr[2])];
    <a id="L143"></a>cdat3 := cr[3][1:len(cr[3])];
    <a id="L144"></a>cdat4 := cr[4][1:len(cr[4])];
    <a id="L145"></a>pdat := pr[1:len(pr)];
    <a id="L146"></a>n := len(cdat0);

    <a id="L148"></a><span class="comment">// The up filter.</span>
    <a id="L149"></a>sum := 0;
    <a id="L150"></a>for i := 0; i &lt; n; i++ {
        <a id="L151"></a>cdat2[i] = cdat0[i] - pdat[i];
        <a id="L152"></a>sum += abs8(cdat2[i]);
    <a id="L153"></a>}
    <a id="L154"></a>best := sum;
    <a id="L155"></a>filter := ftUp;

    <a id="L157"></a><span class="comment">// The Paeth filter.</span>
    <a id="L158"></a>sum = 0;
    <a id="L159"></a>for i := 0; i &lt; bpp; i++ {
        <a id="L160"></a>cdat4[i] = cdat0[i] - paeth(0, pdat[i], 0);
        <a id="L161"></a>sum += abs8(cdat4[i]);
    <a id="L162"></a>}
    <a id="L163"></a>for i := bpp; i &lt; n; i++ {
        <a id="L164"></a>cdat4[i] = cdat0[i] - paeth(cdat0[i-bpp], pdat[i], pdat[i-bpp]);
        <a id="L165"></a>sum += abs8(cdat4[i]);
        <a id="L166"></a>if sum &gt;= best {
            <a id="L167"></a>break
        <a id="L168"></a>}
    <a id="L169"></a>}
    <a id="L170"></a>if sum &lt; best {
        <a id="L171"></a>best = sum;
        <a id="L172"></a>filter = ftPaeth;
    <a id="L173"></a>}

    <a id="L175"></a><span class="comment">// The none filter.</span>
    <a id="L176"></a>sum = 0;
    <a id="L177"></a>for i := 0; i &lt; n; i++ {
        <a id="L178"></a>sum += abs8(cdat0[i]);
        <a id="L179"></a>if sum &gt;= best {
            <a id="L180"></a>break
        <a id="L181"></a>}
    <a id="L182"></a>}
    <a id="L183"></a>if sum &lt; best {
        <a id="L184"></a>best = sum;
        <a id="L185"></a>filter = ftNone;
    <a id="L186"></a>}

    <a id="L188"></a><span class="comment">// The sub filter.</span>
    <a id="L189"></a>sum = 0;
    <a id="L190"></a>for i := 0; i &lt; bpp; i++ {
        <a id="L191"></a>cdat1[i] = cdat0[i];
        <a id="L192"></a>sum += abs8(cdat1[i]);
    <a id="L193"></a>}
    <a id="L194"></a>for i := bpp; i &lt; n; i++ {
        <a id="L195"></a>cdat1[i] = cdat0[i] - cdat0[i-bpp];
        <a id="L196"></a>sum += abs8(cdat1[i]);
        <a id="L197"></a>if sum &gt;= best {
            <a id="L198"></a>break
        <a id="L199"></a>}
    <a id="L200"></a>}
    <a id="L201"></a>if sum &lt; best {
        <a id="L202"></a>best = sum;
        <a id="L203"></a>filter = ftSub;
    <a id="L204"></a>}

    <a id="L206"></a><span class="comment">// The average filter.</span>
    <a id="L207"></a>sum = 0;
    <a id="L208"></a>for i := 0; i &lt; bpp; i++ {
        <a id="L209"></a>cdat3[i] = cdat0[i] - pdat[i]/2;
        <a id="L210"></a>sum += abs8(cdat3[i]);
    <a id="L211"></a>}
    <a id="L212"></a>for i := bpp; i &lt; n; i++ {
        <a id="L213"></a>cdat3[i] = cdat0[i] - uint8((int(cdat0[i-bpp])+int(pdat[i]))/2);
        <a id="L214"></a>sum += abs8(cdat3[i]);
        <a id="L215"></a>if sum &gt;= best {
            <a id="L216"></a>break
        <a id="L217"></a>}
    <a id="L218"></a>}
    <a id="L219"></a>if sum &lt; best {
        <a id="L220"></a>best = sum;
        <a id="L221"></a>filter = ftAverage;
    <a id="L222"></a>}

    <a id="L224"></a>return filter;
<a id="L225"></a>}

<a id="L227"></a>func writeImage(w io.Writer, m image.Image, ct uint8) os.Error {
    <a id="L228"></a>zw, err := zlib.NewDeflater(w);
    <a id="L229"></a>if err != nil {
        <a id="L230"></a>return err
    <a id="L231"></a>}
    <a id="L232"></a>defer zw.Close();

    <a id="L234"></a>bpp := 0; <span class="comment">// Bytes per pixel.</span>
    <a id="L235"></a>var paletted *image.Paletted;
    <a id="L236"></a>switch ct {
    <a id="L237"></a>case ctTrueColor:
        <a id="L238"></a>bpp = 3
    <a id="L239"></a>case ctPaletted:
        <a id="L240"></a>bpp = 1;
        <a id="L241"></a>paletted = m.(*image.Paletted);
    <a id="L242"></a>case ctTrueColorAlpha:
        <a id="L243"></a>bpp = 4
    <a id="L244"></a>}
    <a id="L245"></a><span class="comment">// cr[*] and pr are the bytes for the current and previous row.</span>
    <a id="L246"></a><span class="comment">// cr[0] is unfiltered (or equivalently, filtered with the ftNone filter).</span>
    <a id="L247"></a><span class="comment">// cr[ft], for non-zero filter types ft, are buffers for transforming cr[0] under the</span>
    <a id="L248"></a><span class="comment">// other PNG filter types. These buffers are allocated once and re-used for each row.</span>
    <a id="L249"></a><span class="comment">// The +1 is for the per-row filter type, which is at cr[*][0].</span>
    <a id="L250"></a>var cr [nFilter][]uint8;
    <a id="L251"></a>for i := 0; i &lt; len(cr); i++ {
        <a id="L252"></a>cr[i] = make([]uint8, 1+bpp*m.Width());
        <a id="L253"></a>cr[i][0] = uint8(i);
    <a id="L254"></a>}
    <a id="L255"></a>pr := make([]uint8, 1+bpp*m.Width());

    <a id="L257"></a>for y := 0; y &lt; m.Height(); y++ {
        <a id="L258"></a><span class="comment">// Convert from colors to bytes.</span>
        <a id="L259"></a>switch ct {
        <a id="L260"></a>case ctTrueColor:
            <a id="L261"></a>for x := 0; x &lt; m.Width(); x++ {
                <a id="L262"></a><span class="comment">// We have previously verified that the alpha value is fully opaque.</span>
                <a id="L263"></a>r, g, b, _ := m.At(x, y).RGBA();
                <a id="L264"></a>cr[0][3*x+1] = uint8(r &gt;&gt; 24);
                <a id="L265"></a>cr[0][3*x+2] = uint8(g &gt;&gt; 24);
                <a id="L266"></a>cr[0][3*x+3] = uint8(b &gt;&gt; 24);
            <a id="L267"></a>}
        <a id="L268"></a>case ctPaletted:
            <a id="L269"></a>for x := 0; x &lt; m.Width(); x++ {
                <a id="L270"></a>cr[0][x+1] = paletted.ColorIndexAt(x, y)
            <a id="L271"></a>}
        <a id="L272"></a>case ctTrueColorAlpha:
            <a id="L273"></a><span class="comment">// Convert from image.Image (which is alpha-premultiplied) to PNG&#39;s non-alpha-premultiplied.</span>
            <a id="L274"></a>for x := 0; x &lt; m.Width(); x++ {
                <a id="L275"></a>c := image.NRGBAColorModel.Convert(m.At(x, y)).(image.NRGBAColor);
                <a id="L276"></a>cr[0][4*x+1] = c.R;
                <a id="L277"></a>cr[0][4*x+2] = c.G;
                <a id="L278"></a>cr[0][4*x+3] = c.B;
                <a id="L279"></a>cr[0][4*x+4] = c.A;
            <a id="L280"></a>}
        <a id="L281"></a>}

        <a id="L283"></a><span class="comment">// Apply the filter.</span>
        <a id="L284"></a>f := filter(cr[0:nFilter], pr, bpp);

        <a id="L286"></a><span class="comment">// Write the compressed bytes.</span>
        <a id="L287"></a>_, err = zw.Write(cr[f]);
        <a id="L288"></a>if err != nil {
            <a id="L289"></a>return err
        <a id="L290"></a>}

        <a id="L292"></a><span class="comment">// The current row for y is the previous row for y+1.</span>
        <a id="L293"></a>pr, cr[0] = cr[0], pr;
    <a id="L294"></a>}
    <a id="L295"></a>return nil;
<a id="L296"></a>}

<a id="L298"></a><span class="comment">// Write the actual image data to one or more IDAT chunks.</span>
<a id="L299"></a>func (e *encoder) writeIDATs() {
    <a id="L300"></a>if e.err != nil {
        <a id="L301"></a>return
    <a id="L302"></a>}
    <a id="L303"></a>var bw *bufio.Writer;
    <a id="L304"></a>bw, e.err = bufio.NewWriterSize(e, 1&lt;&lt;15);
    <a id="L305"></a>if e.err != nil {
        <a id="L306"></a>return
    <a id="L307"></a>}
    <a id="L308"></a>e.err = writeImage(bw, e.m, e.colorType);
    <a id="L309"></a>if e.err != nil {
        <a id="L310"></a>return
    <a id="L311"></a>}
    <a id="L312"></a>e.err = bw.Flush();
<a id="L313"></a>}

<a id="L315"></a>func (e *encoder) writeIEND() { e.writeChunk(e.tmp[0:0], &#34;IEND&#34;) }

<a id="L317"></a><span class="comment">// Encode writes the Image m to w in PNG format. Any Image may be encoded, but</span>
<a id="L318"></a><span class="comment">// images that are not image.NRGBA might be encoded lossily.</span>
<a id="L319"></a>func Encode(w io.Writer, m image.Image) os.Error {
    <a id="L320"></a><span class="comment">// Obviously, negative widths and heights are invalid. Furthermore, the PNG</span>
    <a id="L321"></a><span class="comment">// spec section 11.2.2 says that zero is invalid. Excessively large images are</span>
    <a id="L322"></a><span class="comment">// also rejected.</span>
    <a id="L323"></a>mw, mh := int64(m.Width()), int64(m.Height());
    <a id="L324"></a>if mw &lt;= 0 || mh &lt;= 0 || mw &gt;= 1&lt;&lt;32 || mh &gt;= 1&lt;&lt;32 {
        <a id="L325"></a>return FormatError(&#34;invalid image size: &#34; + strconv.Itoa64(mw) + &#34;x&#34; + strconv.Itoa64(mw))
    <a id="L326"></a>}

    <a id="L328"></a>var e encoder;
    <a id="L329"></a>e.w = w;
    <a id="L330"></a>e.m = m;
    <a id="L331"></a>e.colorType = uint8(ctTrueColorAlpha);
    <a id="L332"></a>pal, _ := m.(*image.Paletted);
    <a id="L333"></a>if pal != nil {
        <a id="L334"></a>e.colorType = ctPaletted
    <a id="L335"></a>} else if opaque(m) {
        <a id="L336"></a>e.colorType = ctTrueColor
    <a id="L337"></a>}

    <a id="L339"></a>_, e.err = io.WriteString(w, pngHeader);
    <a id="L340"></a>e.writeIHDR();
    <a id="L341"></a>if pal != nil {
        <a id="L342"></a>e.writePLTE(pal.Palette)
    <a id="L343"></a>}
    <a id="L344"></a>e.writeIDATs();
    <a id="L345"></a>e.writeIEND();
    <a id="L346"></a>return e.err;
<a id="L347"></a>}
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
