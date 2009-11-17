<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/flate/huffman_bit_writer.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/compress/flate/huffman_bit_writer.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package flate

<a id="L7"></a>import (
    <a id="L8"></a>&#34;io&#34;;
    <a id="L9"></a>&#34;math&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;strconv&#34;;
<a id="L12"></a>)

<a id="L14"></a>const (
    <a id="L15"></a><span class="comment">// The largest offset code.</span>
    <a id="L16"></a>offsetCodeCount = 30;

    <a id="L18"></a><span class="comment">// The largest offset code in the extensions.</span>
    <a id="L19"></a>extendedOffsetCodeCount = 42;

    <a id="L21"></a><span class="comment">// The special code used to mark the end of a block.</span>
    <a id="L22"></a>endBlockMarker = 256;

    <a id="L24"></a><span class="comment">// The first length code.</span>
    <a id="L25"></a>lengthCodesStart = 257;

    <a id="L27"></a><span class="comment">// The number of codegen codes.</span>
    <a id="L28"></a>codegenCodeCount = 19;
    <a id="L29"></a>badCode          = 255;
<a id="L30"></a>)

<a id="L32"></a><span class="comment">// The number of extra bits needed by length code X - LENGTH_CODES_START.</span>
<a id="L33"></a>var lengthExtraBits = []int8{
    <a id="L34"></a><span class="comment">/* 257 */</span> 0, 0, 0,
    <a id="L35"></a><span class="comment">/* 260 */</span> 0, 0, 0, 0, 0, 1, 1, 1, 1, 2,
    <a id="L36"></a><span class="comment">/* 270 */</span> 2, 2, 2, 3, 3, 3, 3, 4, 4, 4,
    <a id="L37"></a><span class="comment">/* 280 */</span> 4, 5, 5, 5, 5, 0,
<a id="L38"></a>}

<a id="L40"></a><span class="comment">// The length indicated by length code X - LENGTH_CODES_START.</span>
<a id="L41"></a>var lengthBase = []uint32{
    <a id="L42"></a>0, 1, 2, 3, 4, 5, 6, 7, 8, 10,
    <a id="L43"></a>12, 14, 16, 20, 24, 28, 32, 40, 48, 56,
    <a id="L44"></a>64, 80, 96, 112, 128, 160, 192, 224, 255,
<a id="L45"></a>}

<a id="L47"></a><span class="comment">// offset code word extra bits.</span>
<a id="L48"></a>var offsetExtraBits = []int8{
    <a id="L49"></a>0, 0, 0, 0, 1, 1, 2, 2, 3, 3,
    <a id="L50"></a>4, 4, 5, 5, 6, 6, 7, 7, 8, 8,
    <a id="L51"></a>9, 9, 10, 10, 11, 11, 12, 12, 13, 13,
    <a id="L52"></a><span class="comment">/* extended window */</span>
    <a id="L53"></a>14, 14, 15, 15, 16, 16, 17, 17, 18, 18, 19, 19, 20, 20,
<a id="L54"></a>}

<a id="L56"></a>var offsetBase = []uint32{
    <a id="L57"></a><span class="comment">/* normal deflate */</span>
    <a id="L58"></a>0x000000, 0x000001, 0x000002, 0x000003, 0x000004,
    <a id="L59"></a>0x000006, 0x000008, 0x00000c, 0x000010, 0x000018,
    <a id="L60"></a>0x000020, 0x000030, 0x000040, 0x000060, 0x000080,
    <a id="L61"></a>0x0000c0, 0x000100, 0x000180, 0x000200, 0x000300,
    <a id="L62"></a>0x000400, 0x000600, 0x000800, 0x000c00, 0x001000,
    <a id="L63"></a>0x001800, 0x002000, 0x003000, 0x004000, 0x006000,

    <a id="L65"></a><span class="comment">/* extended window */</span>
    <a id="L66"></a>0x008000, 0x00c000, 0x010000, 0x018000, 0x020000,
    <a id="L67"></a>0x030000, 0x040000, 0x060000, 0x080000, 0x0c0000,
    <a id="L68"></a>0x100000, 0x180000, 0x200000, 0x300000,
<a id="L69"></a>}

<a id="L71"></a><span class="comment">// The odd order in which the codegen code sizes are written.</span>
<a id="L72"></a>var codegenOrder = []uint32{16, 17, 18, 0, 8, 7, 9, 6, 10, 5, 11, 4, 12, 3, 13, 2, 14, 1, 15}

<a id="L74"></a>type huffmanBitWriter struct {
    <a id="L75"></a>w   io.Writer;
    <a id="L76"></a><span class="comment">// Data waiting to be written is bytes[0:nbytes]</span>
    <a id="L77"></a><span class="comment">// and then the low nbits of bits.</span>
    <a id="L78"></a>bits            uint32;
    <a id="L79"></a>nbits           uint32;
    <a id="L80"></a>bytes           [64]byte;
    <a id="L81"></a>nbytes          int;
    <a id="L82"></a>literalFreq     []int32;
    <a id="L83"></a>offsetFreq      []int32;
    <a id="L84"></a>codegen         []uint8;
    <a id="L85"></a>codegenFreq     []int32;
    <a id="L86"></a>literalEncoding *huffmanEncoder;
    <a id="L87"></a>offsetEncoding  *huffmanEncoder;
    <a id="L88"></a>codegenEncoding *huffmanEncoder;
    <a id="L89"></a>err             os.Error;
<a id="L90"></a>}

<a id="L92"></a>type WrongValueError struct {
    <a id="L93"></a>name  string;
    <a id="L94"></a>from  int32;
    <a id="L95"></a>to    int32;
    <a id="L96"></a>value int32;
<a id="L97"></a>}

<a id="L99"></a>func newHuffmanBitWriter(w io.Writer) *huffmanBitWriter {
    <a id="L100"></a>return &amp;huffmanBitWriter{
        <a id="L101"></a>w: w,
        <a id="L102"></a>literalFreq: make([]int32, maxLit),
        <a id="L103"></a>offsetFreq: make([]int32, extendedOffsetCodeCount),
        <a id="L104"></a>codegen: make([]uint8, maxLit+extendedOffsetCodeCount+1),
        <a id="L105"></a>codegenFreq: make([]int32, codegenCodeCount),
        <a id="L106"></a>literalEncoding: newHuffmanEncoder(maxLit),
        <a id="L107"></a>offsetEncoding: newHuffmanEncoder(extendedOffsetCodeCount),
        <a id="L108"></a>codegenEncoding: newHuffmanEncoder(codegenCodeCount),
    <a id="L109"></a>}
<a id="L110"></a>}

<a id="L112"></a>func (err WrongValueError) String() string {
    <a id="L113"></a>return &#34;huffmanBitWriter: &#34; + err.name + &#34; should belong to [&#34; + strconv.Itoa64(int64(err.from)) + &#34;;&#34; +
        <a id="L114"></a>strconv.Itoa64(int64(err.to)) + &#34;] but actual value is &#34; + strconv.Itoa64(int64(err.value))
<a id="L115"></a>}

<a id="L117"></a>func (w *huffmanBitWriter) flushBits() {
    <a id="L118"></a>if w.err != nil {
        <a id="L119"></a>w.nbits = 0;
        <a id="L120"></a>return;
    <a id="L121"></a>}
    <a id="L122"></a>bits := w.bits;
    <a id="L123"></a>w.bits &gt;&gt;= 16;
    <a id="L124"></a>w.nbits -= 16;
    <a id="L125"></a>n := w.nbytes;
    <a id="L126"></a>w.bytes[n] = byte(bits);
    <a id="L127"></a>w.bytes[n+1] = byte(bits &gt;&gt; 8);
    <a id="L128"></a>if n += 2; n &gt;= len(w.bytes) {
        <a id="L129"></a>_, w.err = w.w.Write(&amp;w.bytes);
        <a id="L130"></a>n = 0;
    <a id="L131"></a>}
    <a id="L132"></a>w.nbytes = n;
<a id="L133"></a>}

<a id="L135"></a>func (w *huffmanBitWriter) flush() {
    <a id="L136"></a>if w.err != nil {
        <a id="L137"></a>w.nbits = 0;
        <a id="L138"></a>return;
    <a id="L139"></a>}
    <a id="L140"></a>n := w.nbytes;
    <a id="L141"></a>if w.nbits &gt; 8 {
        <a id="L142"></a>w.bytes[n] = byte(w.bits);
        <a id="L143"></a>w.bits &gt;&gt;= 8;
        <a id="L144"></a>w.nbits -= 8;
        <a id="L145"></a>n++;
    <a id="L146"></a>}
    <a id="L147"></a>if w.nbits &gt; 0 {
        <a id="L148"></a>w.bytes[n] = byte(w.bits);
        <a id="L149"></a>w.nbits = 0;
        <a id="L150"></a>n++;
    <a id="L151"></a>}
    <a id="L152"></a>w.bits = 0;
    <a id="L153"></a>_, w.err = w.w.Write(w.bytes[0:n]);
    <a id="L154"></a>w.nbytes = 0;
<a id="L155"></a>}

<a id="L157"></a>func (w *huffmanBitWriter) writeBits(b, nb int32) {
    <a id="L158"></a>w.bits |= uint32(b) &lt;&lt; w.nbits;
    <a id="L159"></a>if w.nbits += uint32(nb); w.nbits &gt;= 16 {
        <a id="L160"></a>w.flushBits()
    <a id="L161"></a>}
<a id="L162"></a>}

<a id="L164"></a>func (w *huffmanBitWriter) writeBytes(bytes []byte) {
    <a id="L165"></a>if w.err != nil {
        <a id="L166"></a>return
    <a id="L167"></a>}
    <a id="L168"></a>n := w.nbytes;
    <a id="L169"></a>if w.nbits == 8 {
        <a id="L170"></a>w.bytes[n] = byte(w.bits);
        <a id="L171"></a>w.nbits = 0;
        <a id="L172"></a>n++;
    <a id="L173"></a>}
    <a id="L174"></a>if w.nbits != 0 {
        <a id="L175"></a>w.err = InternalError(&#34;writeBytes with unfinished bits&#34;);
        <a id="L176"></a>return;
    <a id="L177"></a>}
    <a id="L178"></a>if n != 0 {
        <a id="L179"></a>_, w.err = w.w.Write(w.bytes[0:n]);
        <a id="L180"></a>if w.err != nil {
            <a id="L181"></a>return
        <a id="L182"></a>}
    <a id="L183"></a>}
    <a id="L184"></a>w.nbytes = 0;
    <a id="L185"></a>_, w.err = w.w.Write(bytes);
<a id="L186"></a>}

<a id="L188"></a><span class="comment">// RFC 1951 3.2.7 specifies a special run-length encoding for specifiying</span>
<a id="L189"></a><span class="comment">// the literal and offset lengths arrays (which are concatenated into a single</span>
<a id="L190"></a><span class="comment">// array).  This method generates that run-length encoding.</span>
<a id="L191"></a><span class="comment">//</span>
<a id="L192"></a><span class="comment">// The result is written into the codegen array, and the frequencies</span>
<a id="L193"></a><span class="comment">// of each code is written into the codegenFreq array.</span>
<a id="L194"></a><span class="comment">// Codes 0-15 are single byte codes. Codes 16-18 are followed by additional</span>
<a id="L195"></a><span class="comment">// information.  Code badCode is an end marker</span>
<a id="L196"></a><span class="comment">//</span>
<a id="L197"></a><span class="comment">//  numLiterals      The number of literals in literalEncoding</span>
<a id="L198"></a><span class="comment">//  numOffsets       The number of offsets in offsetEncoding</span>
<a id="L199"></a>func (w *huffmanBitWriter) generateCodegen(numLiterals int, numOffsets int) {
    <a id="L200"></a>fillInt32s(w.codegenFreq, 0);
    <a id="L201"></a><span class="comment">// Note that we are using codegen both as a temporary variable for holding</span>
    <a id="L202"></a><span class="comment">// a copy of the frequencies, and as the place where we put the result.</span>
    <a id="L203"></a><span class="comment">// This is fine because the output is always shorter than the input used</span>
    <a id="L204"></a><span class="comment">// so far.</span>
    <a id="L205"></a>codegen := w.codegen; <span class="comment">// cache</span>
    <a id="L206"></a><span class="comment">// Copy the concatenated code sizes to codegen.  Put a marker at the end.</span>
    <a id="L207"></a>copyUint8s(codegen[0:numLiterals], w.literalEncoding.codeBits);
    <a id="L208"></a>copyUint8s(codegen[numLiterals:numLiterals+numOffsets], w.offsetEncoding.codeBits);
    <a id="L209"></a>codegen[numLiterals+numOffsets] = badCode;

    <a id="L211"></a>size := codegen[0];
    <a id="L212"></a>count := 1;
    <a id="L213"></a>outIndex := 0;
    <a id="L214"></a>for inIndex := 1; size != badCode; inIndex++ {
        <a id="L215"></a><span class="comment">// INVARIANT: We have seen &#34;count&#34; copies of size that have not yet</span>
        <a id="L216"></a><span class="comment">// had output generated for them.</span>
        <a id="L217"></a>nextSize := codegen[inIndex];
        <a id="L218"></a>if nextSize == size {
            <a id="L219"></a>count++;
            <a id="L220"></a>continue;
        <a id="L221"></a>}
        <a id="L222"></a><span class="comment">// We need to generate codegen indicating &#34;count&#34; of size.</span>
        <a id="L223"></a>if size != 0 {
            <a id="L224"></a>codegen[outIndex] = size;
            <a id="L225"></a>outIndex++;
            <a id="L226"></a>w.codegenFreq[size]++;
            <a id="L227"></a>count--;
            <a id="L228"></a>for count &gt;= 3 {
                <a id="L229"></a>n := min(count, 6);
                <a id="L230"></a>codegen[outIndex] = 16;
                <a id="L231"></a>outIndex++;
                <a id="L232"></a>codegen[outIndex] = uint8(n - 3);
                <a id="L233"></a>outIndex++;
                <a id="L234"></a>w.codegenFreq[16]++;
                <a id="L235"></a>count -= n;
            <a id="L236"></a>}
        <a id="L237"></a>} else {
            <a id="L238"></a>for count &gt;= 11 {
                <a id="L239"></a>n := min(count, 138);
                <a id="L240"></a>codegen[outIndex] = 18;
                <a id="L241"></a>outIndex++;
                <a id="L242"></a>codegen[outIndex] = uint8(n - 11);
                <a id="L243"></a>outIndex++;
                <a id="L244"></a>w.codegenFreq[18]++;
                <a id="L245"></a>count -= n;
            <a id="L246"></a>}
            <a id="L247"></a>if count &gt;= 3 {
                <a id="L248"></a><span class="comment">// count &gt;= 3 &amp;&amp; count &lt;= 10</span>
                <a id="L249"></a>codegen[outIndex] = 17;
                <a id="L250"></a>outIndex++;
                <a id="L251"></a>codegen[outIndex] = uint8(count - 3);
                <a id="L252"></a>outIndex++;
                <a id="L253"></a>w.codegenFreq[17]++;
                <a id="L254"></a>count = 0;
            <a id="L255"></a>}
        <a id="L256"></a>}
        <a id="L257"></a>count--;
        <a id="L258"></a>for ; count &gt;= 0; count-- {
            <a id="L259"></a>codegen[outIndex] = size;
            <a id="L260"></a>outIndex++;
            <a id="L261"></a>w.codegenFreq[size]++;
        <a id="L262"></a>}
        <a id="L263"></a><span class="comment">// Set up invariant for next time through the loop.</span>
        <a id="L264"></a>size = nextSize;
        <a id="L265"></a>count = 1;
    <a id="L266"></a>}
    <a id="L267"></a><span class="comment">// Marker indicating the end of the codegen.</span>
    <a id="L268"></a>codegen[outIndex] = badCode;
<a id="L269"></a>}

<a id="L271"></a>func (w *huffmanBitWriter) writeCode(code *huffmanEncoder, literal uint32) {
    <a id="L272"></a>if w.err != nil {
        <a id="L273"></a>return
    <a id="L274"></a>}
    <a id="L275"></a>w.writeBits(int32(code.code[literal]), int32(code.codeBits[literal]));
<a id="L276"></a>}

<a id="L278"></a><span class="comment">// Write the header of a dynamic Huffman block to the output stream.</span>
<a id="L279"></a><span class="comment">//</span>
<a id="L280"></a><span class="comment">//  numLiterals  The number of literals specified in codegen</span>
<a id="L281"></a><span class="comment">//  numOffsets   The number of offsets specified in codegen</span>
<a id="L282"></a><span class="comment">//  numCodegens  Tne number of codegens used in codegen</span>
<a id="L283"></a>func (w *huffmanBitWriter) writeDynamicHeader(numLiterals int, numOffsets int, numCodegens int, isEof bool) {
    <a id="L284"></a>if w.err != nil {
        <a id="L285"></a>return
    <a id="L286"></a>}
    <a id="L287"></a>var firstBits int32 = 4;
    <a id="L288"></a>if isEof {
        <a id="L289"></a>firstBits = 5
    <a id="L290"></a>}
    <a id="L291"></a>w.writeBits(firstBits, 3);
    <a id="L292"></a>w.writeBits(int32(numLiterals-257), 5);
    <a id="L293"></a>if numOffsets &gt; offsetCodeCount {
        <a id="L294"></a><span class="comment">// Extended version of deflater</span>
        <a id="L295"></a>w.writeBits(int32(offsetCodeCount+((numOffsets-(1+offsetCodeCount))&gt;&gt;3)), 5);
        <a id="L296"></a>w.writeBits(int32((numOffsets-(1+offsetCodeCount))&amp;0x7), 3);
    <a id="L297"></a>} else {
        <a id="L298"></a>w.writeBits(int32(numOffsets-1), 5)
    <a id="L299"></a>}
    <a id="L300"></a>w.writeBits(int32(numCodegens-4), 4);

    <a id="L302"></a>for i := 0; i &lt; numCodegens; i++ {
        <a id="L303"></a>value := w.codegenEncoding.codeBits[codegenOrder[i]];
        <a id="L304"></a>w.writeBits(int32(value), 3);
    <a id="L305"></a>}

    <a id="L307"></a>i := 0;
    <a id="L308"></a>for {
        <a id="L309"></a>var codeWord int = int(w.codegen[i]);
        <a id="L310"></a>i++;
        <a id="L311"></a>if codeWord == badCode {
            <a id="L312"></a>break
        <a id="L313"></a>}
        <a id="L314"></a><span class="comment">// The low byte contains the actual code to generate.</span>
        <a id="L315"></a>w.writeCode(w.codegenEncoding, uint32(codeWord));

        <a id="L317"></a>switch codeWord {
        <a id="L318"></a>case 16:
            <a id="L319"></a>w.writeBits(int32(w.codegen[i]), 2);
            <a id="L320"></a>i++;
            <a id="L321"></a>break;
        <a id="L322"></a>case 17:
            <a id="L323"></a>w.writeBits(int32(w.codegen[i]), 3);
            <a id="L324"></a>i++;
            <a id="L325"></a>break;
        <a id="L326"></a>case 18:
            <a id="L327"></a>w.writeBits(int32(w.codegen[i]), 7);
            <a id="L328"></a>i++;
            <a id="L329"></a>break;
        <a id="L330"></a>}
    <a id="L331"></a>}
<a id="L332"></a>}

<a id="L334"></a>func (w *huffmanBitWriter) writeStoredHeader(length int, isEof bool) {
    <a id="L335"></a>if w.err != nil {
        <a id="L336"></a>return
    <a id="L337"></a>}
    <a id="L338"></a>var flag int32;
    <a id="L339"></a>if isEof {
        <a id="L340"></a>flag = 1
    <a id="L341"></a>}
    <a id="L342"></a>w.writeBits(flag, 3);
    <a id="L343"></a>w.flush();
    <a id="L344"></a>w.writeBits(int32(length), 16);
    <a id="L345"></a>w.writeBits(int32(^uint16(length)), 16);
<a id="L346"></a>}

<a id="L348"></a>func (w *huffmanBitWriter) writeFixedHeader(isEof bool) {
    <a id="L349"></a>if w.err != nil {
        <a id="L350"></a>return
    <a id="L351"></a>}
    <a id="L352"></a><span class="comment">// Indicate that we are a fixed Huffman block</span>
    <a id="L353"></a>var value int32 = 2;
    <a id="L354"></a>if isEof {
        <a id="L355"></a>value = 3
    <a id="L356"></a>}
    <a id="L357"></a>w.writeBits(value, 3);
<a id="L358"></a>}

<a id="L360"></a>func (w *huffmanBitWriter) writeBlock(tokens []token, eof bool, input []byte) {
    <a id="L361"></a>if w.err != nil {
        <a id="L362"></a>return
    <a id="L363"></a>}
    <a id="L364"></a>fillInt32s(w.literalFreq, 0);
    <a id="L365"></a>fillInt32s(w.offsetFreq, 0);

    <a id="L367"></a>n := len(tokens);
    <a id="L368"></a>tokens = tokens[0 : n+1];
    <a id="L369"></a>tokens[n] = endBlockMarker;

    <a id="L371"></a>totalLength := -1; <span class="comment">// Subtract 1 for endBlock.</span>
    <a id="L372"></a>for _, t := range tokens {
        <a id="L373"></a>switch t.typ() {
        <a id="L374"></a>case literalType:
            <a id="L375"></a>w.literalFreq[t.literal()]++;
            <a id="L376"></a>totalLength++;
            <a id="L377"></a>break;
        <a id="L378"></a>case matchType:
            <a id="L379"></a>length := t.length();
            <a id="L380"></a>offset := t.offset();
            <a id="L381"></a>totalLength += int(length + 3);
            <a id="L382"></a>w.literalFreq[lengthCodesStart+lengthCode(length)]++;
            <a id="L383"></a>w.offsetFreq[offsetCode(offset)]++;
            <a id="L384"></a>break;
        <a id="L385"></a>}
    <a id="L386"></a>}
    <a id="L387"></a>w.literalEncoding.generate(w.literalFreq, 15);
    <a id="L388"></a>w.offsetEncoding.generate(w.offsetFreq, 15);

    <a id="L390"></a><span class="comment">// get the number of literals</span>
    <a id="L391"></a>numLiterals := len(w.literalFreq);
    <a id="L392"></a>for w.literalFreq[numLiterals-1] == 0 {
        <a id="L393"></a>numLiterals--
    <a id="L394"></a>}
    <a id="L395"></a><span class="comment">// get the number of offsets</span>
    <a id="L396"></a>numOffsets := len(w.offsetFreq);
    <a id="L397"></a>for numOffsets &gt; 1 &amp;&amp; w.offsetFreq[numOffsets-1] == 0 {
        <a id="L398"></a>numOffsets--
    <a id="L399"></a>}
    <a id="L400"></a>storedBytes := 0;
    <a id="L401"></a>if input != nil {
        <a id="L402"></a>storedBytes = len(input)
    <a id="L403"></a>}
    <a id="L404"></a>var extraBits int64;
    <a id="L405"></a>var storedSize int64;
    <a id="L406"></a>if storedBytes &lt;= maxStoreBlockSize &amp;&amp; input != nil {
        <a id="L407"></a>storedSize = int64((storedBytes + 5) * 8);
        <a id="L408"></a><span class="comment">// We only bother calculating the costs of the extra bits required by</span>
        <a id="L409"></a><span class="comment">// the length of offset fields (which will be the same for both fixed</span>
        <a id="L410"></a><span class="comment">// and dynamic encoding), if we need to compare those two encodings</span>
        <a id="L411"></a><span class="comment">// against stored encoding.</span>
        <a id="L412"></a>for lengthCode := lengthCodesStart + 8; lengthCode &lt; numLiterals; lengthCode++ {
            <a id="L413"></a><span class="comment">// First eight length codes have extra size = 0.</span>
            <a id="L414"></a>extraBits += int64(w.literalFreq[lengthCode]) * int64(lengthExtraBits[lengthCode-lengthCodesStart])
        <a id="L415"></a>}
        <a id="L416"></a>for offsetCode := 4; offsetCode &lt; numOffsets; offsetCode++ {
            <a id="L417"></a><span class="comment">// First four offset codes have extra size = 0.</span>
            <a id="L418"></a>extraBits += int64(w.offsetFreq[offsetCode]) * int64(offsetExtraBits[offsetCode])
        <a id="L419"></a>}
    <a id="L420"></a>} else {
        <a id="L421"></a>storedSize = math.MaxInt32
    <a id="L422"></a>}

    <a id="L424"></a><span class="comment">// Figure out which generates smaller code, fixed Huffman, dynamic</span>
    <a id="L425"></a><span class="comment">// Huffman, or just storing the data.</span>
    <a id="L426"></a>var fixedSize int64 = math.MaxInt64;
    <a id="L427"></a>if numOffsets &lt;= offsetCodeCount {
        <a id="L428"></a>fixedSize = int64(3) +
            <a id="L429"></a>fixedLiteralEncoding.bitLength(w.literalFreq) +
            <a id="L430"></a>fixedOffsetEncoding.bitLength(w.offsetFreq) +
            <a id="L431"></a>extraBits
    <a id="L432"></a>}
    <a id="L433"></a><span class="comment">// Generate codegen and codegenFrequencies, which indicates how to encode</span>
    <a id="L434"></a><span class="comment">// the literalEncoding and the offsetEncoding.</span>
    <a id="L435"></a>w.generateCodegen(numLiterals, numOffsets);
    <a id="L436"></a>w.codegenEncoding.generate(w.codegenFreq, 7);
    <a id="L437"></a>numCodegens := len(w.codegenFreq);
    <a id="L438"></a>for numCodegens &gt; 4 &amp;&amp; w.codegenFreq[codegenOrder[numCodegens-1]] == 0 {
        <a id="L439"></a>numCodegens--
    <a id="L440"></a>}
    <a id="L441"></a>extensionSummand := 0;
    <a id="L442"></a>if numOffsets &gt; offsetCodeCount {
        <a id="L443"></a>extensionSummand = 3
    <a id="L444"></a>}
    <a id="L445"></a>dynamicHeader := int64(3+5+5+4+(3*numCodegens)) +
        <a id="L446"></a><span class="comment">// Following line is an extension.</span>
        <a id="L447"></a>int64(extensionSummand) +
        <a id="L448"></a>w.codegenEncoding.bitLength(w.codegenFreq) +
        <a id="L449"></a>int64(extraBits) +
        <a id="L450"></a>int64(w.codegenFreq[16]*2) +
        <a id="L451"></a>int64(w.codegenFreq[17]*3) +
        <a id="L452"></a>int64(w.codegenFreq[18]*7);
    <a id="L453"></a>dynamicSize := dynamicHeader +
        <a id="L454"></a>w.literalEncoding.bitLength(w.literalFreq) +
        <a id="L455"></a>w.offsetEncoding.bitLength(w.offsetFreq);

    <a id="L457"></a>if storedSize &lt; fixedSize &amp;&amp; storedSize &lt; dynamicSize {
        <a id="L458"></a>w.writeStoredHeader(storedBytes, eof);
        <a id="L459"></a>w.writeBytes(input[0:storedBytes]);
        <a id="L460"></a>return;
    <a id="L461"></a>}
    <a id="L462"></a>var literalEncoding *huffmanEncoder;
    <a id="L463"></a>var offsetEncoding *huffmanEncoder;

    <a id="L465"></a>if fixedSize &lt;= dynamicSize {
        <a id="L466"></a>w.writeFixedHeader(eof);
        <a id="L467"></a>literalEncoding = fixedLiteralEncoding;
        <a id="L468"></a>offsetEncoding = fixedOffsetEncoding;
    <a id="L469"></a>} else {
        <a id="L470"></a><span class="comment">// Write the header.</span>
        <a id="L471"></a>w.writeDynamicHeader(numLiterals, numOffsets, numCodegens, eof);
        <a id="L472"></a>literalEncoding = w.literalEncoding;
        <a id="L473"></a>offsetEncoding = w.offsetEncoding;
    <a id="L474"></a>}

    <a id="L476"></a><span class="comment">// Write the tokens.</span>
    <a id="L477"></a>for _, t := range tokens {
        <a id="L478"></a>switch t.typ() {
        <a id="L479"></a>case literalType:
            <a id="L480"></a>w.writeCode(literalEncoding, t.literal());
            <a id="L481"></a>break;
        <a id="L482"></a>case matchType:
            <a id="L483"></a><span class="comment">// Write the length</span>
            <a id="L484"></a>length := t.length();
            <a id="L485"></a>lengthCode := lengthCode(length);
            <a id="L486"></a>w.writeCode(literalEncoding, lengthCode+lengthCodesStart);
            <a id="L487"></a>extraLengthBits := int32(lengthExtraBits[lengthCode]);
            <a id="L488"></a>if extraLengthBits &gt; 0 {
                <a id="L489"></a>extraLength := int32(length - lengthBase[lengthCode]);
                <a id="L490"></a>w.writeBits(extraLength, extraLengthBits);
            <a id="L491"></a>}
            <a id="L492"></a><span class="comment">// Write the offset</span>
            <a id="L493"></a>offset := t.offset();
            <a id="L494"></a>offsetCode := offsetCode(offset);
            <a id="L495"></a>w.writeCode(offsetEncoding, offsetCode);
            <a id="L496"></a>extraOffsetBits := int32(offsetExtraBits[offsetCode]);
            <a id="L497"></a>if extraOffsetBits &gt; 0 {
                <a id="L498"></a>extraOffset := int32(offset - offsetBase[offsetCode]);
                <a id="L499"></a>w.writeBits(extraOffset, extraOffsetBits);
            <a id="L500"></a>}
            <a id="L501"></a>break;
        <a id="L502"></a>default:
            <a id="L503"></a>panic(&#34;unknown token type: &#34; + string(t))
        <a id="L504"></a>}
    <a id="L505"></a>}
<a id="L506"></a>}
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
