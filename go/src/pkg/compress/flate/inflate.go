<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/flate/inflate.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/compress/flate/inflate.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The flate package implements the DEFLATE compressed data</span>
<a id="L6"></a><span class="comment">// format, described in RFC 1951.  The gzip and zlib packages</span>
<a id="L7"></a><span class="comment">// implement access to DEFLATE-based file formats.</span>
<a id="L8"></a>package flate

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bufio&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;os&#34;;
    <a id="L14"></a>&#34;strconv&#34;;
<a id="L15"></a>)

<a id="L17"></a>const (
    <a id="L18"></a>maxCodeLen = 16;    <span class="comment">// max length of Huffman code</span>
    <a id="L19"></a>maxHist    = 32768; <span class="comment">// max history required</span>
    <a id="L20"></a>maxLit     = 286;
    <a id="L21"></a>maxDist    = 32;
    <a id="L22"></a>numCodes   = 19; <span class="comment">// number of codes in Huffman meta-code</span>
<a id="L23"></a>)

<a id="L25"></a><span class="comment">// A CorruptInputError reports the presence of corrupt input at a given offset.</span>
<a id="L26"></a>type CorruptInputError int64

<a id="L28"></a>func (e CorruptInputError) String() string {
    <a id="L29"></a>return &#34;flate: corrupt input before offset &#34; + strconv.Itoa64(int64(e))
<a id="L30"></a>}

<a id="L32"></a><span class="comment">// An InternalError reports an error in the flate code itself.</span>
<a id="L33"></a>type InternalError string

<a id="L35"></a>func (e InternalError) String() string { return &#34;flate: internal error: &#34; + string(e) }

<a id="L37"></a><span class="comment">// A ReadError reports an error encountered while reading input.</span>
<a id="L38"></a>type ReadError struct {
    <a id="L39"></a>Offset int64;    <span class="comment">// byte offset where error occurred</span>
    <a id="L40"></a>Error  os.Error; <span class="comment">// error returned by underlying Read</span>
<a id="L41"></a>}

<a id="L43"></a>func (e *ReadError) String() string {
    <a id="L44"></a>return &#34;flate: read error at offset &#34; + strconv.Itoa64(e.Offset) + &#34;: &#34; + e.Error.String()
<a id="L45"></a>}

<a id="L47"></a><span class="comment">// A WriteError reports an error encountered while writing output.</span>
<a id="L48"></a>type WriteError struct {
    <a id="L49"></a>Offset int64;    <span class="comment">// byte offset where error occurred</span>
    <a id="L50"></a>Error  os.Error; <span class="comment">// error returned by underlying Read</span>
<a id="L51"></a>}

<a id="L53"></a>func (e *WriteError) String() string {
    <a id="L54"></a>return &#34;flate: write error at offset &#34; + strconv.Itoa64(e.Offset) + &#34;: &#34; + e.Error.String()
<a id="L55"></a>}

<a id="L57"></a><span class="comment">// Huffman decoder is based on</span>
<a id="L58"></a><span class="comment">// J. Brian Connell, ``A Huffman-Shannon-Fano Code,&#39;&#39;</span>
<a id="L59"></a><span class="comment">// Proceedings of the IEEE, 61(7) (July 1973), pp 1046-1047.</span>
<a id="L60"></a>type huffmanDecoder struct {
    <a id="L61"></a><span class="comment">// min, max code length</span>
    <a id="L62"></a>min, max int;

    <a id="L64"></a><span class="comment">// limit[i] = largest code word of length i</span>
    <a id="L65"></a><span class="comment">// Given code v of length n,</span>
    <a id="L66"></a><span class="comment">// need more bits if v &gt; limit[n].</span>
    <a id="L67"></a>limit [maxCodeLen + 1]int;

    <a id="L69"></a><span class="comment">// base[i] = smallest code word of length i - seq number</span>
    <a id="L70"></a>base [maxCodeLen + 1]int;

    <a id="L72"></a><span class="comment">// codes[seq number] = output code.</span>
    <a id="L73"></a><span class="comment">// Given code v of length n, value is</span>
    <a id="L74"></a><span class="comment">// codes[v - base[n]].</span>
    <a id="L75"></a>codes []int;
<a id="L76"></a>}

<a id="L78"></a><span class="comment">// Initialize Huffman decoding tables from array of code lengths.</span>
<a id="L79"></a>func (h *huffmanDecoder) init(bits []int) bool {
    <a id="L80"></a><span class="comment">// TODO(rsc): Return false sometimes.</span>

    <a id="L82"></a><span class="comment">// Count number of codes of each length,</span>
    <a id="L83"></a><span class="comment">// compute min and max length.</span>
    <a id="L84"></a>var count [maxCodeLen + 1]int;
    <a id="L85"></a>var min, max int;
    <a id="L86"></a>for _, n := range bits {
        <a id="L87"></a>if n == 0 {
            <a id="L88"></a>continue
        <a id="L89"></a>}
        <a id="L90"></a>if min == 0 || n &lt; min {
            <a id="L91"></a>min = n
        <a id="L92"></a>}
        <a id="L93"></a>if n &gt; max {
            <a id="L94"></a>max = n
        <a id="L95"></a>}
        <a id="L96"></a>count[n]++;
    <a id="L97"></a>}
    <a id="L98"></a>if max == 0 {
        <a id="L99"></a>return false
    <a id="L100"></a>}

    <a id="L102"></a>h.min = min;
    <a id="L103"></a>h.max = max;


    <a id="L106"></a><span class="comment">// For each code range, compute</span>
    <a id="L107"></a><span class="comment">// nextcode (first code of that length),</span>
    <a id="L108"></a><span class="comment">// limit (last code of that length), and</span>
    <a id="L109"></a><span class="comment">// base (offset from first code to sequence number).</span>
    <a id="L110"></a>code := 0;
    <a id="L111"></a>seq := 0;
    <a id="L112"></a>var nextcode [maxCodeLen]int;
    <a id="L113"></a>for i := min; i &lt;= max; i++ {
        <a id="L114"></a>n := count[i];
        <a id="L115"></a>nextcode[i] = code;
        <a id="L116"></a>h.base[i] = code - seq;
        <a id="L117"></a>code += n;
        <a id="L118"></a>seq += n;
        <a id="L119"></a>h.limit[i] = code - 1;
        <a id="L120"></a>code &lt;&lt;= 1;
    <a id="L121"></a>}

    <a id="L123"></a><span class="comment">// Make array mapping sequence numbers to codes.</span>
    <a id="L124"></a>if len(h.codes) &lt; len(bits) {
        <a id="L125"></a>h.codes = make([]int, len(bits))
    <a id="L126"></a>}
    <a id="L127"></a>for i, n := range bits {
        <a id="L128"></a>if n == 0 {
            <a id="L129"></a>continue
        <a id="L130"></a>}
        <a id="L131"></a>code := nextcode[n];
        <a id="L132"></a>nextcode[n]++;
        <a id="L133"></a>seq := code - h.base[n];
        <a id="L134"></a>h.codes[seq] = i;
    <a id="L135"></a>}
    <a id="L136"></a>return true;
<a id="L137"></a>}

<a id="L139"></a><span class="comment">// Hard-coded Huffman tables for DEFLATE algorithm.</span>
<a id="L140"></a><span class="comment">// See RFC 1951, section 3.2.6.</span>
<a id="L141"></a>var fixedHuffmanDecoder = huffmanDecoder{
    <a id="L142"></a>7, 9,
    <a id="L143"></a>[maxCodeLen + 1]int{7: 23, 199, 511},
    <a id="L144"></a>[maxCodeLen + 1]int{7: 0, 24, 224},
    <a id="L145"></a>[]int{
        <a id="L146"></a><span class="comment">// length 7: 256-279</span>
        <a id="L147"></a>256, 257, 258, 259, 260, 261, 262,
        <a id="L148"></a>263, 264, 265, 266, 267, 268, 269,
        <a id="L149"></a>270, 271, 272, 273, 274, 275, 276,
        <a id="L150"></a>277, 278, 279,

        <a id="L152"></a><span class="comment">// length 8: 0-143</span>
        <a id="L153"></a>0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
        <a id="L154"></a>12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
        <a id="L155"></a>22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
        <a id="L156"></a>32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
        <a id="L157"></a>42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
        <a id="L158"></a>52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
        <a id="L159"></a>62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
        <a id="L160"></a>72, 73, 74, 75, 76, 77, 78, 79, 80, 81,
        <a id="L161"></a>82, 83, 84, 85, 86, 87, 88, 89, 90, 91,
        <a id="L162"></a>92, 93, 94, 95, 96, 97, 98, 99, 100,
        <a id="L163"></a>101, 102, 103, 104, 105, 106, 107, 108,
        <a id="L164"></a>109, 110, 111, 112, 113, 114, 115, 116,
        <a id="L165"></a>117, 118, 119, 120, 121, 122, 123, 124,
        <a id="L166"></a>125, 126, 127, 128, 129, 130, 131, 132,
        <a id="L167"></a>133, 134, 135, 136, 137, 138, 139, 140,
        <a id="L168"></a>141, 142, 143,

        <a id="L170"></a><span class="comment">// length 8: 280-287</span>
        <a id="L171"></a>280, 281, 282, 283, 284, 285, 286, 287,

        <a id="L173"></a><span class="comment">// length 9: 144-255</span>
        <a id="L174"></a>144, 145, 146, 147, 148, 149, 150, 151,
        <a id="L175"></a>152, 153, 154, 155, 156, 157, 158, 159,
        <a id="L176"></a>160, 161, 162, 163, 164, 165, 166, 167,
        <a id="L177"></a>168, 169, 170, 171, 172, 173, 174, 175,
        <a id="L178"></a>176, 177, 178, 179, 180, 181, 182, 183,
        <a id="L179"></a>184, 185, 186, 187, 188, 189, 190, 191,
        <a id="L180"></a>192, 193, 194, 195, 196, 197, 198, 199,
        <a id="L181"></a>200, 201, 202, 203, 204, 205, 206, 207,
        <a id="L182"></a>208, 209, 210, 211, 212, 213, 214, 215,
        <a id="L183"></a>216, 217, 218, 219, 220, 221, 222, 223,
        <a id="L184"></a>224, 225, 226, 227, 228, 229, 230, 231,
        <a id="L185"></a>232, 233, 234, 235, 236, 237, 238, 239,
        <a id="L186"></a>240, 241, 242, 243, 244, 245, 246, 247,
        <a id="L187"></a>248, 249, 250, 251, 252, 253, 254, 255,
    <a id="L188"></a>},
<a id="L189"></a>}

<a id="L191"></a><span class="comment">// The actual read interface needed by NewInflater.</span>
<a id="L192"></a><span class="comment">// If the passed in io.Reader does not also have ReadByte,</span>
<a id="L193"></a><span class="comment">// the NewInflater will introduce its own buffering.</span>
<a id="L194"></a>type Reader interface {
    <a id="L195"></a>io.Reader;
    <a id="L196"></a>ReadByte() (c byte, err os.Error);
<a id="L197"></a>}

<a id="L199"></a><span class="comment">// Inflate state.</span>
<a id="L200"></a>type inflater struct {
    <a id="L201"></a><span class="comment">// Input/output sources.</span>
    <a id="L202"></a>r       Reader;
    <a id="L203"></a>w       io.Writer;
    <a id="L204"></a>roffset int64;
    <a id="L205"></a>woffset int64;

    <a id="L207"></a><span class="comment">// Input bits, in top of b.</span>
    <a id="L208"></a>b   uint32;
    <a id="L209"></a>nb  uint;

    <a id="L211"></a><span class="comment">// Huffman decoders for literal/length, distance.</span>
    <a id="L212"></a>h1, h2 huffmanDecoder;

    <a id="L214"></a><span class="comment">// Length arrays used to define Huffman codes.</span>
    <a id="L215"></a>bits     [maxLit + maxDist]int;
    <a id="L216"></a>codebits [numCodes]int;

    <a id="L218"></a><span class="comment">// Output history, buffer.</span>
    <a id="L219"></a>hist  [maxHist]byte;
    <a id="L220"></a>hp    int;  <span class="comment">// current output position in buffer</span>
    <a id="L221"></a>hfull bool; <span class="comment">// buffer has filled at least once</span>

    <a id="L223"></a><span class="comment">// Temporary buffer (avoids repeated allocation).</span>
    <a id="L224"></a>buf [4]byte;
<a id="L225"></a>}

<a id="L227"></a>func (f *inflater) inflate() (err os.Error) {
    <a id="L228"></a>final := false;
    <a id="L229"></a>for err == nil &amp;&amp; !final {
        <a id="L230"></a>for f.nb &lt; 1+2 {
            <a id="L231"></a>if err = f.moreBits(); err != nil {
                <a id="L232"></a>return
            <a id="L233"></a>}
        <a id="L234"></a>}
        <a id="L235"></a>final = f.b&amp;1 == 1;
        <a id="L236"></a>f.b &gt;&gt;= 1;
        <a id="L237"></a>typ := f.b &amp; 3;
        <a id="L238"></a>f.b &gt;&gt;= 2;
        <a id="L239"></a>f.nb -= 1 + 2;
        <a id="L240"></a>switch typ {
        <a id="L241"></a>case 0:
            <a id="L242"></a>err = f.dataBlock()
        <a id="L243"></a>case 1:
            <a id="L244"></a><span class="comment">// compressed, fixed Huffman tables</span>
            <a id="L245"></a>err = f.decodeBlock(&amp;fixedHuffmanDecoder, nil)
        <a id="L246"></a>case 2:
            <a id="L247"></a><span class="comment">// compressed, dynamic Huffman tables</span>
            <a id="L248"></a>if err = f.readHuffman(); err == nil {
                <a id="L249"></a>err = f.decodeBlock(&amp;f.h1, &amp;f.h2)
            <a id="L250"></a>}
        <a id="L251"></a>default:
            <a id="L252"></a><span class="comment">// 3 is reserved.</span>
            <a id="L253"></a>err = CorruptInputError(f.roffset)
        <a id="L254"></a>}
    <a id="L255"></a>}
    <a id="L256"></a>return;
<a id="L257"></a>}

<a id="L259"></a><span class="comment">// RFC 1951 section 3.2.7.</span>
<a id="L260"></a><span class="comment">// Compression with dynamic Huffman codes</span>

<a id="L262"></a>var codeOrder = [...]int{16, 17, 18, 0, 8, 7, 9, 6, 10, 5, 11, 4, 12, 3, 13, 2, 14, 1, 15}

<a id="L264"></a>func (f *inflater) readHuffman() os.Error {
    <a id="L265"></a><span class="comment">// HLIT[5], HDIST[5], HCLEN[4].</span>
    <a id="L266"></a>for f.nb &lt; 5+5+4 {
        <a id="L267"></a>if err := f.moreBits(); err != nil {
            <a id="L268"></a>return err
        <a id="L269"></a>}
    <a id="L270"></a>}
    <a id="L271"></a>nlit := int(f.b&amp;0x1F) + 257;
    <a id="L272"></a>f.b &gt;&gt;= 5;
    <a id="L273"></a>ndist := int(f.b&amp;0x1F) + 1;
    <a id="L274"></a>f.b &gt;&gt;= 5;
    <a id="L275"></a>nclen := int(f.b&amp;0xF) + 4;
    <a id="L276"></a>f.b &gt;&gt;= 4;
    <a id="L277"></a>f.nb -= 5 + 5 + 4;

    <a id="L279"></a><span class="comment">// (HCLEN+4)*3 bits: code lengths in the magic codeOrder order.</span>
    <a id="L280"></a>for i := 0; i &lt; nclen; i++ {
        <a id="L281"></a>for f.nb &lt; 3 {
            <a id="L282"></a>if err := f.moreBits(); err != nil {
                <a id="L283"></a>return err
            <a id="L284"></a>}
        <a id="L285"></a>}
        <a id="L286"></a>f.codebits[codeOrder[i]] = int(f.b &amp; 0x7);
        <a id="L287"></a>f.b &gt;&gt;= 3;
        <a id="L288"></a>f.nb -= 3;
    <a id="L289"></a>}
    <a id="L290"></a>for i := nclen; i &lt; len(codeOrder); i++ {
        <a id="L291"></a>f.codebits[codeOrder[i]] = 0
    <a id="L292"></a>}
    <a id="L293"></a>if !f.h1.init(&amp;f.codebits) {
        <a id="L294"></a>return CorruptInputError(f.roffset)
    <a id="L295"></a>}

    <a id="L297"></a><span class="comment">// HLIT + 257 code lengths, HDIST + 1 code lengths,</span>
    <a id="L298"></a><span class="comment">// using the code length Huffman code.</span>
    <a id="L299"></a>for i, n := 0, nlit+ndist; i &lt; n; {
        <a id="L300"></a>x, err := f.huffSym(&amp;f.h1);
        <a id="L301"></a>if err != nil {
            <a id="L302"></a>return err
        <a id="L303"></a>}
        <a id="L304"></a>if x &lt; 16 {
            <a id="L305"></a><span class="comment">// Actual length.</span>
            <a id="L306"></a>f.bits[i] = x;
            <a id="L307"></a>i++;
            <a id="L308"></a>continue;
        <a id="L309"></a>}
        <a id="L310"></a><span class="comment">// Repeat previous length or zero.</span>
        <a id="L311"></a>var rep int;
        <a id="L312"></a>var nb uint;
        <a id="L313"></a>var b int;
        <a id="L314"></a>switch x {
        <a id="L315"></a>default:
            <a id="L316"></a>return InternalError(&#34;unexpected length code&#34;)
        <a id="L317"></a>case 16:
            <a id="L318"></a>rep = 3;
            <a id="L319"></a>nb = 2;
            <a id="L320"></a>if i == 0 {
                <a id="L321"></a>return CorruptInputError(f.roffset)
            <a id="L322"></a>}
            <a id="L323"></a>b = f.bits[i-1];
        <a id="L324"></a>case 17:
            <a id="L325"></a>rep = 3;
            <a id="L326"></a>nb = 3;
            <a id="L327"></a>b = 0;
        <a id="L328"></a>case 18:
            <a id="L329"></a>rep = 11;
            <a id="L330"></a>nb = 7;
            <a id="L331"></a>b = 0;
        <a id="L332"></a>}
        <a id="L333"></a>for f.nb &lt; nb {
            <a id="L334"></a>if err := f.moreBits(); err != nil {
                <a id="L335"></a>return err
            <a id="L336"></a>}
        <a id="L337"></a>}
        <a id="L338"></a>rep += int(f.b &amp; uint32(1&lt;&lt;nb-1));
        <a id="L339"></a>f.b &gt;&gt;= nb;
        <a id="L340"></a>f.nb -= nb;
        <a id="L341"></a>if i+rep &gt; n {
            <a id="L342"></a>return CorruptInputError(f.roffset)
        <a id="L343"></a>}
        <a id="L344"></a>for j := 0; j &lt; rep; j++ {
            <a id="L345"></a>f.bits[i] = b;
            <a id="L346"></a>i++;
        <a id="L347"></a>}
    <a id="L348"></a>}

    <a id="L350"></a>if !f.h1.init(f.bits[0:nlit]) || !f.h2.init(f.bits[nlit:nlit+ndist]) {
        <a id="L351"></a>return CorruptInputError(f.roffset)
    <a id="L352"></a>}

    <a id="L354"></a>return nil;
<a id="L355"></a>}

<a id="L357"></a><span class="comment">// Decode a single Huffman block from f.</span>
<a id="L358"></a><span class="comment">// hl and hd are the Huffman states for the lit/length values</span>
<a id="L359"></a><span class="comment">// and the distance values, respectively.  If hd == nil, using the</span>
<a id="L360"></a><span class="comment">// fixed distance encoding associated with fixed Huffman blocks.</span>
<a id="L361"></a>func (f *inflater) decodeBlock(hl, hd *huffmanDecoder) os.Error {
    <a id="L362"></a>for {
        <a id="L363"></a>v, err := f.huffSym(hl);
        <a id="L364"></a>if err != nil {
            <a id="L365"></a>return err
        <a id="L366"></a>}
        <a id="L367"></a>var n uint; <span class="comment">// number of bits extra</span>
        <a id="L368"></a>var length int;
        <a id="L369"></a>switch {
        <a id="L370"></a>case v &lt; 256:
            <a id="L371"></a>f.hist[f.hp] = byte(v);
            <a id="L372"></a>f.hp++;
            <a id="L373"></a>if f.hp == len(f.hist) {
                <a id="L374"></a>if err = f.flush(); err != nil {
                    <a id="L375"></a>return err
                <a id="L376"></a>}
            <a id="L377"></a>}
            <a id="L378"></a>continue;
        <a id="L379"></a>case v == 256:
            <a id="L380"></a>return nil
        <a id="L381"></a><span class="comment">// otherwise, reference to older data</span>
        <a id="L382"></a>case v &lt; 265:
            <a id="L383"></a>length = v - (257 - 3);
            <a id="L384"></a>n = 0;
        <a id="L385"></a>case v &lt; 269:
            <a id="L386"></a>length = v*2 - (265*2 - 11);
            <a id="L387"></a>n = 1;
        <a id="L388"></a>case v &lt; 273:
            <a id="L389"></a>length = v*4 - (269*4 - 19);
            <a id="L390"></a>n = 2;
        <a id="L391"></a>case v &lt; 277:
            <a id="L392"></a>length = v*8 - (273*8 - 35);
            <a id="L393"></a>n = 3;
        <a id="L394"></a>case v &lt; 281:
            <a id="L395"></a>length = v*16 - (277*16 - 67);
            <a id="L396"></a>n = 4;
        <a id="L397"></a>case v &lt; 285:
            <a id="L398"></a>length = v*32 - (281*32 - 131);
            <a id="L399"></a>n = 5;
        <a id="L400"></a>default:
            <a id="L401"></a>length = 258;
            <a id="L402"></a>n = 0;
        <a id="L403"></a>}
        <a id="L404"></a>if n &gt; 0 {
            <a id="L405"></a>for f.nb &lt; n {
                <a id="L406"></a>if err = f.moreBits(); err != nil {
                    <a id="L407"></a>return err
                <a id="L408"></a>}
            <a id="L409"></a>}
            <a id="L410"></a>length += int(f.b &amp; uint32(1&lt;&lt;n-1));
            <a id="L411"></a>f.b &gt;&gt;= n;
            <a id="L412"></a>f.nb -= n;
        <a id="L413"></a>}

        <a id="L415"></a>var dist int;
        <a id="L416"></a>if hd == nil {
            <a id="L417"></a>for f.nb &lt; 5 {
                <a id="L418"></a>if err = f.moreBits(); err != nil {
                    <a id="L419"></a>return err
                <a id="L420"></a>}
            <a id="L421"></a>}
            <a id="L422"></a>dist = int(reverseByte[(f.b&amp;0x1F)&lt;&lt;3]);
            <a id="L423"></a>f.b &gt;&gt;= 5;
            <a id="L424"></a>f.nb -= 5;
        <a id="L425"></a>} else {
            <a id="L426"></a>if dist, err = f.huffSym(hd); err != nil {
                <a id="L427"></a>return err
            <a id="L428"></a>}
        <a id="L429"></a>}

        <a id="L431"></a>switch {
        <a id="L432"></a>case dist &lt; 4:
            <a id="L433"></a>dist++
        <a id="L434"></a>case dist &gt;= 30:
            <a id="L435"></a>return CorruptInputError(f.roffset)
        <a id="L436"></a>default:
            <a id="L437"></a>nb := uint(dist-2) &gt;&gt; 1;
            <a id="L438"></a><span class="comment">// have 1 bit in bottom of dist, need nb more.</span>
            <a id="L439"></a>extra := (dist &amp; 1) &lt;&lt; nb;
            <a id="L440"></a>for f.nb &lt; nb {
                <a id="L441"></a>if err = f.moreBits(); err != nil {
                    <a id="L442"></a>return err
                <a id="L443"></a>}
            <a id="L444"></a>}
            <a id="L445"></a>extra |= int(f.b &amp; uint32(1&lt;&lt;nb-1));
            <a id="L446"></a>f.b &gt;&gt;= nb;
            <a id="L447"></a>f.nb -= nb;
            <a id="L448"></a>dist = 1&lt;&lt;(nb+1) + 1 + extra;
        <a id="L449"></a>}

        <a id="L451"></a><span class="comment">// Copy history[-dist:-dist+length] into output.</span>
        <a id="L452"></a>if dist &gt; len(f.hist) {
            <a id="L453"></a>return InternalError(&#34;bad history distance&#34;)
        <a id="L454"></a>}

        <a id="L456"></a><span class="comment">// No check on length; encoding can be prescient.</span>
        <a id="L457"></a>if !f.hfull &amp;&amp; dist &gt; f.hp {
            <a id="L458"></a>return CorruptInputError(f.roffset)
        <a id="L459"></a>}

        <a id="L461"></a>p := f.hp - dist;
        <a id="L462"></a>if p &lt; 0 {
            <a id="L463"></a>p += len(f.hist)
        <a id="L464"></a>}
        <a id="L465"></a>for i := 0; i &lt; length; i++ {
            <a id="L466"></a>f.hist[f.hp] = f.hist[p];
            <a id="L467"></a>f.hp++;
            <a id="L468"></a>p++;
            <a id="L469"></a>if f.hp == len(f.hist) {
                <a id="L470"></a>if err = f.flush(); err != nil {
                    <a id="L471"></a>return err
                <a id="L472"></a>}
            <a id="L473"></a>}
            <a id="L474"></a>if p == len(f.hist) {
                <a id="L475"></a>p = 0
            <a id="L476"></a>}
        <a id="L477"></a>}
    <a id="L478"></a>}
    <a id="L479"></a>panic(&#34;unreached&#34;);
<a id="L480"></a>}

<a id="L482"></a><span class="comment">// Copy a single uncompressed data block from input to output.</span>
<a id="L483"></a>func (f *inflater) dataBlock() os.Error {
    <a id="L484"></a><span class="comment">// Uncompressed.</span>
    <a id="L485"></a><span class="comment">// Discard current half-byte.</span>
    <a id="L486"></a>f.nb = 0;
    <a id="L487"></a>f.b = 0;

    <a id="L489"></a><span class="comment">// Length then ones-complement of length.</span>
    <a id="L490"></a>nr, err := io.ReadFull(f.r, f.buf[0:4]);
    <a id="L491"></a>f.roffset += int64(nr);
    <a id="L492"></a>if err != nil {
        <a id="L493"></a>return &amp;ReadError{f.roffset, err}
    <a id="L494"></a>}
    <a id="L495"></a>n := int(f.buf[0]) | int(f.buf[1])&lt;&lt;8;
    <a id="L496"></a>nn := int(f.buf[2]) | int(f.buf[3])&lt;&lt;8;
    <a id="L497"></a>if uint16(nn) != uint16(^n) {
        <a id="L498"></a>return CorruptInputError(f.roffset)
    <a id="L499"></a>}

    <a id="L501"></a><span class="comment">// Read len bytes into history,</span>
    <a id="L502"></a><span class="comment">// writing as history fills.</span>
    <a id="L503"></a>for n &gt; 0 {
        <a id="L504"></a>m := len(f.hist) - f.hp;
        <a id="L505"></a>if m &gt; n {
            <a id="L506"></a>m = n
        <a id="L507"></a>}
        <a id="L508"></a>m, err := io.ReadFull(f.r, f.hist[f.hp:f.hp+m]);
        <a id="L509"></a>f.roffset += int64(m);
        <a id="L510"></a>if err != nil {
            <a id="L511"></a>return &amp;ReadError{f.roffset, err}
        <a id="L512"></a>}
        <a id="L513"></a>n -= m;
        <a id="L514"></a>f.hp += m;
        <a id="L515"></a>if f.hp == len(f.hist) {
            <a id="L516"></a>if err = f.flush(); err != nil {
                <a id="L517"></a>return err
            <a id="L518"></a>}
        <a id="L519"></a>}
    <a id="L520"></a>}
    <a id="L521"></a>return nil;
<a id="L522"></a>}

<a id="L524"></a>func (f *inflater) moreBits() os.Error {
    <a id="L525"></a>c, err := f.r.ReadByte();
    <a id="L526"></a>if err != nil {
        <a id="L527"></a>if err == os.EOF {
            <a id="L528"></a>err = io.ErrUnexpectedEOF
        <a id="L529"></a>}
        <a id="L530"></a>return err;
    <a id="L531"></a>}
    <a id="L532"></a>f.roffset++;
    <a id="L533"></a>f.b |= uint32(c) &lt;&lt; f.nb;
    <a id="L534"></a>f.nb += 8;
    <a id="L535"></a>return nil;
<a id="L536"></a>}

<a id="L538"></a><span class="comment">// Read the next Huffman-encoded symbol from f according to h.</span>
<a id="L539"></a>func (f *inflater) huffSym(h *huffmanDecoder) (int, os.Error) {
    <a id="L540"></a>for n := uint(h.min); n &lt;= uint(h.max); n++ {
        <a id="L541"></a>lim := h.limit[n];
        <a id="L542"></a>if lim == -1 {
            <a id="L543"></a>continue
        <a id="L544"></a>}
        <a id="L545"></a>for f.nb &lt; n {
            <a id="L546"></a>if err := f.moreBits(); err != nil {
                <a id="L547"></a>return 0, err
            <a id="L548"></a>}
        <a id="L549"></a>}
        <a id="L550"></a>v := int(f.b &amp; uint32(1&lt;&lt;n-1));
        <a id="L551"></a>v &lt;&lt;= 16 - n;
        <a id="L552"></a>v = int(reverseByte[v&gt;&gt;8]) | int(reverseByte[v&amp;0xFF])&lt;&lt;8; <span class="comment">// reverse bits</span>
        <a id="L553"></a>if v &lt;= lim {
            <a id="L554"></a>f.b &gt;&gt;= n;
            <a id="L555"></a>f.nb -= n;
            <a id="L556"></a>return h.codes[v-h.base[n]], nil;
        <a id="L557"></a>}
    <a id="L558"></a>}
    <a id="L559"></a>return 0, CorruptInputError(f.roffset);
<a id="L560"></a>}

<a id="L562"></a><span class="comment">// Flush any buffered output to the underlying writer.</span>
<a id="L563"></a>func (f *inflater) flush() os.Error {
    <a id="L564"></a>if f.hp == 0 {
        <a id="L565"></a>return nil
    <a id="L566"></a>}
    <a id="L567"></a>n, err := f.w.Write(f.hist[0:f.hp]);
    <a id="L568"></a>if n != f.hp &amp;&amp; err == nil {
        <a id="L569"></a>err = io.ErrShortWrite
    <a id="L570"></a>}
    <a id="L571"></a>if err != nil {
        <a id="L572"></a>return &amp;WriteError{f.woffset, err}
    <a id="L573"></a>}
    <a id="L574"></a>f.woffset += int64(f.hp);
    <a id="L575"></a>f.hp = 0;
    <a id="L576"></a>f.hfull = true;
    <a id="L577"></a>return nil;
<a id="L578"></a>}

<a id="L580"></a>func makeReader(r io.Reader) Reader {
    <a id="L581"></a>if rr, ok := r.(Reader); ok {
        <a id="L582"></a>return rr
    <a id="L583"></a>}
    <a id="L584"></a>return bufio.NewReader(r);
<a id="L585"></a>}

<a id="L587"></a><span class="comment">// Inflate reads DEFLATE-compressed data from r and writes</span>
<a id="L588"></a><span class="comment">// the uncompressed data to w.</span>
<a id="L589"></a>func (f *inflater) inflater(r io.Reader, w io.Writer) os.Error {
    <a id="L590"></a>f.r = makeReader(r);
    <a id="L591"></a>f.w = w;
    <a id="L592"></a>f.woffset = 0;
    <a id="L593"></a>if err := f.inflate(); err != nil {
        <a id="L594"></a>return err
    <a id="L595"></a>}
    <a id="L596"></a>if err := f.flush(); err != nil {
        <a id="L597"></a>return err
    <a id="L598"></a>}
    <a id="L599"></a>return nil;
<a id="L600"></a>}

<a id="L602"></a><span class="comment">// NewInflater returns a new ReadCloser that can be used</span>
<a id="L603"></a><span class="comment">// to read the uncompressed version of r.  It is the caller&#39;s</span>
<a id="L604"></a><span class="comment">// responsibility to call Close on the ReadCloser when</span>
<a id="L605"></a><span class="comment">// finished reading.</span>
<a id="L606"></a>func NewInflater(r io.Reader) io.ReadCloser {
    <a id="L607"></a>var f inflater;
    <a id="L608"></a>pr, pw := io.Pipe();
    <a id="L609"></a>go func() { pw.CloseWithError(f.inflater(r, pw)) }();
    <a id="L610"></a>return pr;
<a id="L611"></a>}
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
