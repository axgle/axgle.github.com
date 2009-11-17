<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/flate/huffman_code.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/compress/flate/huffman_code.go</h1>

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
    <a id="L8"></a>&#34;math&#34;;
    <a id="L9"></a>&#34;sort&#34;;
<a id="L10"></a>)

<a id="L12"></a>type huffmanEncoder struct {
    <a id="L13"></a>codeBits []uint8;
    <a id="L14"></a>code     []uint16;
<a id="L15"></a>}

<a id="L17"></a>type literalNode struct {
    <a id="L18"></a>literal uint16;
    <a id="L19"></a>freq    int32;
<a id="L20"></a>}

<a id="L22"></a>type chain struct {
    <a id="L23"></a><span class="comment">// The sum of the leaves in this tree</span>
    <a id="L24"></a>freq int32;

    <a id="L26"></a><span class="comment">// The number of literals to the left of this item at this level</span>
    <a id="L27"></a>leafCount int32;

    <a id="L29"></a><span class="comment">// The right child of this chain in the previous level.</span>
    <a id="L30"></a>up  *chain;
<a id="L31"></a>}

<a id="L33"></a>type levelInfo struct {
    <a id="L34"></a><span class="comment">// Our level.  for better printing</span>
    <a id="L35"></a>level int32;

    <a id="L37"></a><span class="comment">// The most recent chain generated for this level</span>
    <a id="L38"></a>lastChain *chain;

    <a id="L40"></a><span class="comment">// The frequency of the next character to add to this level</span>
    <a id="L41"></a>nextCharFreq int32;

    <a id="L43"></a><span class="comment">// The frequency of the next pair (from level below) to add to this level.</span>
    <a id="L44"></a><span class="comment">// Only valid if the &#34;needed&#34; value of the next lower level is 0.</span>
    <a id="L45"></a>nextPairFreq int32;

    <a id="L47"></a><span class="comment">// The number of chains remaining to generate for this level before moving</span>
    <a id="L48"></a><span class="comment">// up to the next level</span>
    <a id="L49"></a>needed int32;

    <a id="L51"></a><span class="comment">// The levelInfo for level+1</span>
    <a id="L52"></a>up  *levelInfo;

    <a id="L54"></a><span class="comment">// The levelInfo for level-1</span>
    <a id="L55"></a>down *levelInfo;
<a id="L56"></a>}

<a id="L58"></a>func maxNode() literalNode { return literalNode{math.MaxUint16, math.MaxInt32} }

<a id="L60"></a>func newHuffmanEncoder(size int) *huffmanEncoder {
    <a id="L61"></a>return &amp;huffmanEncoder{make([]uint8, size), make([]uint16, size)}
<a id="L62"></a>}

<a id="L64"></a><span class="comment">// Generates a HuffmanCode corresponding to the fixed literal table</span>
<a id="L65"></a>func generateFixedLiteralEncoding() *huffmanEncoder {
    <a id="L66"></a>h := newHuffmanEncoder(maxLit);
    <a id="L67"></a>codeBits := h.codeBits;
    <a id="L68"></a>code := h.code;
    <a id="L69"></a>var ch uint16;
    <a id="L70"></a>for ch = 0; ch &lt; maxLit; ch++ {
        <a id="L71"></a>var bits uint16;
        <a id="L72"></a>var size uint8;
        <a id="L73"></a>switch {
        <a id="L74"></a>case ch &lt; 144:
            <a id="L75"></a><span class="comment">// size 8, 000110000  .. 10111111</span>
            <a id="L76"></a>bits = ch + 48;
            <a id="L77"></a>size = 8;
            <a id="L78"></a>break;
        <a id="L79"></a>case ch &lt; 256:
            <a id="L80"></a><span class="comment">// size 9, 110010000 .. 111111111</span>
            <a id="L81"></a>bits = ch + 400 - 144;
            <a id="L82"></a>size = 9;
            <a id="L83"></a>break;
        <a id="L84"></a>case ch &lt; 280:
            <a id="L85"></a><span class="comment">// size 7, 0000000 .. 0010111</span>
            <a id="L86"></a>bits = ch - 256;
            <a id="L87"></a>size = 7;
            <a id="L88"></a>break;
        <a id="L89"></a>default:
            <a id="L90"></a><span class="comment">// size 8, 11000000 .. 11000111</span>
            <a id="L91"></a>bits = ch + 192 - 280;
            <a id="L92"></a>size = 8;
        <a id="L93"></a>}
        <a id="L94"></a>codeBits[ch] = size;
        <a id="L95"></a>code[ch] = reverseBits(bits, size);
    <a id="L96"></a>}
    <a id="L97"></a>return h;
<a id="L98"></a>}

<a id="L100"></a>func generateFixedOffsetEncoding() *huffmanEncoder {
    <a id="L101"></a>h := newHuffmanEncoder(30);
    <a id="L102"></a>codeBits := h.codeBits;
    <a id="L103"></a>code := h.code;
    <a id="L104"></a>for ch := uint16(0); ch &lt; 30; ch++ {
        <a id="L105"></a>codeBits[ch] = 5;
        <a id="L106"></a>code[ch] = reverseBits(ch, 5);
    <a id="L107"></a>}
    <a id="L108"></a>return h;
<a id="L109"></a>}

<a id="L111"></a>var fixedLiteralEncoding *huffmanEncoder = generateFixedLiteralEncoding()
<a id="L112"></a>var fixedOffsetEncoding *huffmanEncoder = generateFixedOffsetEncoding()

<a id="L114"></a>func (h *huffmanEncoder) bitLength(freq []int32) int64 {
    <a id="L115"></a>var total int64;
    <a id="L116"></a>for i, f := range freq {
        <a id="L117"></a>if f != 0 {
            <a id="L118"></a>total += int64(f) * int64(h.codeBits[i])
        <a id="L119"></a>}
    <a id="L120"></a>}
    <a id="L121"></a>return total;
<a id="L122"></a>}

<a id="L124"></a><span class="comment">// Generate elements in the chain using an iterative algorithm.</span>
<a id="L125"></a>func (h *huffmanEncoder) generateChains(top *levelInfo, list []literalNode) {
    <a id="L126"></a>n := len(list);
    <a id="L127"></a>list = list[0 : n+1];
    <a id="L128"></a>list[n] = maxNode();

    <a id="L130"></a>l := top;
    <a id="L131"></a>for {
        <a id="L132"></a>if l.nextPairFreq == math.MaxInt32 &amp;&amp; l.nextCharFreq == math.MaxInt32 {
            <a id="L133"></a><span class="comment">// We&#39;ve run out of both leafs and pairs.</span>
            <a id="L134"></a><span class="comment">// End all calculations for this level.</span>
            <a id="L135"></a><span class="comment">// To m sure we never come back to this level or any lower level,</span>
            <a id="L136"></a><span class="comment">// set nextPairFreq impossibly large.</span>
            <a id="L137"></a>l.lastChain = nil;
            <a id="L138"></a>l.needed = 0;
            <a id="L139"></a>l = l.up;
            <a id="L140"></a>l.nextPairFreq = math.MaxInt32;
            <a id="L141"></a>continue;
        <a id="L142"></a>}

        <a id="L144"></a>prevFreq := l.lastChain.freq;
        <a id="L145"></a>if l.nextCharFreq &lt; l.nextPairFreq {
            <a id="L146"></a><span class="comment">// The next item on this row is a leaf node.</span>
            <a id="L147"></a>n := l.lastChain.leafCount + 1;
            <a id="L148"></a>l.lastChain = &amp;chain{l.nextCharFreq, n, l.lastChain.up};
            <a id="L149"></a>l.nextCharFreq = list[n].freq;
        <a id="L150"></a>} else {
            <a id="L151"></a><span class="comment">// The next item on this row is a pair from the previous row.</span>
            <a id="L152"></a><span class="comment">// nextPairFreq isn&#39;t valid until we generate two</span>
            <a id="L153"></a><span class="comment">// more values in the level below</span>
            <a id="L154"></a>l.lastChain = &amp;chain{l.nextPairFreq, l.lastChain.leafCount, l.down.lastChain};
            <a id="L155"></a>l.down.needed = 2;
        <a id="L156"></a>}

        <a id="L158"></a>if l.needed--; l.needed == 0 {
            <a id="L159"></a><span class="comment">// We&#39;ve done everything we need to do for this level.</span>
            <a id="L160"></a><span class="comment">// Continue calculating one level up.  Fill in nextPairFreq</span>
            <a id="L161"></a><span class="comment">// of that level with the sum of the two nodes we&#39;ve just calculated on</span>
            <a id="L162"></a><span class="comment">// this level.</span>
            <a id="L163"></a>up := l.up;
            <a id="L164"></a>if up == nil {
                <a id="L165"></a><span class="comment">// All done!</span>
                <a id="L166"></a>return
            <a id="L167"></a>}
            <a id="L168"></a>up.nextPairFreq = prevFreq + l.lastChain.freq;
            <a id="L169"></a>l = up;
        <a id="L170"></a>} else {
            <a id="L171"></a><span class="comment">// If we stole from below, move down temporarily to replenish it.</span>
            <a id="L172"></a>for l.down.needed &gt; 0 {
                <a id="L173"></a>l = l.down
            <a id="L174"></a>}
        <a id="L175"></a>}
    <a id="L176"></a>}
<a id="L177"></a>}

<a id="L179"></a><span class="comment">// Return the number of literals assigned to each bit size in the Huffman encoding</span>
<a id="L180"></a><span class="comment">//</span>
<a id="L181"></a><span class="comment">// This method is only called when list.length &gt;= 3</span>
<a id="L182"></a><span class="comment">// The cases of 0, 1, and 2 literals are handled by special case code.</span>
<a id="L183"></a><span class="comment">//</span>
<a id="L184"></a><span class="comment">// list  An array of the literals with non-zero frequencies</span>
<a id="L185"></a><span class="comment">//             and their associated frequencies.  The array is in order of increasing</span>
<a id="L186"></a><span class="comment">//             frequency, and has as its last element a special element with frequency</span>
<a id="L187"></a><span class="comment">//             MaxInt32</span>
<a id="L188"></a><span class="comment">// maxBits     The maximum number of bits that should be used to encode any literal.</span>
<a id="L189"></a><span class="comment">// return      An integer array in which array[i] indicates the number of literals</span>
<a id="L190"></a><span class="comment">//             that should be encoded in i bits.</span>
<a id="L191"></a>func (h *huffmanEncoder) bitCounts(list []literalNode, maxBits int32) []int32 {
    <a id="L192"></a>n := int32(len(list));
    <a id="L193"></a>list = list[0 : n+1];
    <a id="L194"></a>list[n] = maxNode();

    <a id="L196"></a><span class="comment">// The tree can&#39;t have greater depth than n - 1, no matter what.  This</span>
    <a id="L197"></a><span class="comment">// saves a little bit of work in some small cases</span>
    <a id="L198"></a>maxBits = minInt32(maxBits, n-1);

    <a id="L200"></a><span class="comment">// Create information about each of the levels.</span>
    <a id="L201"></a><span class="comment">// A bogus &#34;Level 0&#34; whose sole purpose is so that</span>
    <a id="L202"></a><span class="comment">// level1.prev.needed==0.  This makes level1.nextPairFreq</span>
    <a id="L203"></a><span class="comment">// be a legitimate value that never gets chosen.</span>
    <a id="L204"></a>top := &amp;levelInfo{needed: 0};
    <a id="L205"></a>chain2 := &amp;chain{list[1].freq, 2, new(chain)};
    <a id="L206"></a>for level := int32(1); level &lt;= maxBits; level++ {
        <a id="L207"></a><span class="comment">// For every level, the first two items are the first two characters.</span>
        <a id="L208"></a><span class="comment">// We initialize the levels as if we had already figured this out.</span>
        <a id="L209"></a>top = &amp;levelInfo{
            <a id="L210"></a>level: level,
            <a id="L211"></a>lastChain: chain2,
            <a id="L212"></a>nextCharFreq: list[2].freq,
            <a id="L213"></a>nextPairFreq: list[0].freq + list[1].freq,
            <a id="L214"></a>down: top,
        <a id="L215"></a>};
        <a id="L216"></a>top.down.up = top;
        <a id="L217"></a>if level == 1 {
            <a id="L218"></a>top.nextPairFreq = math.MaxInt32
        <a id="L219"></a>}
    <a id="L220"></a>}

    <a id="L222"></a><span class="comment">// We need a total of 2*n - 2 items at top level and have already generated 2.</span>
    <a id="L223"></a>top.needed = 2*n - 4;

    <a id="L225"></a>l := top;
    <a id="L226"></a>for {
        <a id="L227"></a>if l.nextPairFreq == math.MaxInt32 &amp;&amp; l.nextCharFreq == math.MaxInt32 {
            <a id="L228"></a><span class="comment">// We&#39;ve run out of both leafs and pairs.</span>
            <a id="L229"></a><span class="comment">// End all calculations for this level.</span>
            <a id="L230"></a><span class="comment">// To m sure we never come back to this level or any lower level,</span>
            <a id="L231"></a><span class="comment">// set nextPairFreq impossibly large.</span>
            <a id="L232"></a>l.lastChain = nil;
            <a id="L233"></a>l.needed = 0;
            <a id="L234"></a>l = l.up;
            <a id="L235"></a>l.nextPairFreq = math.MaxInt32;
            <a id="L236"></a>continue;
        <a id="L237"></a>}

        <a id="L239"></a>prevFreq := l.lastChain.freq;
        <a id="L240"></a>if l.nextCharFreq &lt; l.nextPairFreq {
            <a id="L241"></a><span class="comment">// The next item on this row is a leaf node.</span>
            <a id="L242"></a>n := l.lastChain.leafCount + 1;
            <a id="L243"></a>l.lastChain = &amp;chain{l.nextCharFreq, n, l.lastChain.up};
            <a id="L244"></a>l.nextCharFreq = list[n].freq;
        <a id="L245"></a>} else {
            <a id="L246"></a><span class="comment">// The next item on this row is a pair from the previous row.</span>
            <a id="L247"></a><span class="comment">// nextPairFreq isn&#39;t valid until we generate two</span>
            <a id="L248"></a><span class="comment">// more values in the level below</span>
            <a id="L249"></a>l.lastChain = &amp;chain{l.nextPairFreq, l.lastChain.leafCount, l.down.lastChain};
            <a id="L250"></a>l.down.needed = 2;
        <a id="L251"></a>}

        <a id="L253"></a>if l.needed--; l.needed == 0 {
            <a id="L254"></a><span class="comment">// We&#39;ve done everything we need to do for this level.</span>
            <a id="L255"></a><span class="comment">// Continue calculating one level up.  Fill in nextPairFreq</span>
            <a id="L256"></a><span class="comment">// of that level with the sum of the two nodes we&#39;ve just calculated on</span>
            <a id="L257"></a><span class="comment">// this level.</span>
            <a id="L258"></a>up := l.up;
            <a id="L259"></a>if up == nil {
                <a id="L260"></a><span class="comment">// All done!</span>
                <a id="L261"></a>break
            <a id="L262"></a>}
            <a id="L263"></a>up.nextPairFreq = prevFreq + l.lastChain.freq;
            <a id="L264"></a>l = up;
        <a id="L265"></a>} else {
            <a id="L266"></a><span class="comment">// If we stole from below, move down temporarily to replenish it.</span>
            <a id="L267"></a>for l.down.needed &gt; 0 {
                <a id="L268"></a>l = l.down
            <a id="L269"></a>}
        <a id="L270"></a>}
    <a id="L271"></a>}


    <a id="L274"></a><span class="comment">// Somethings is wrong if at the end, the top level is null or hasn&#39;t used</span>
    <a id="L275"></a><span class="comment">// all of the leaves.</span>
    <a id="L276"></a>if top.lastChain.leafCount != n {
        <a id="L277"></a>panic(&#34;top.lastChain.leafCount != n&#34;)
    <a id="L278"></a>}

    <a id="L280"></a>bitCount := make([]int32, maxBits+1);
    <a id="L281"></a>bits := 1;
    <a id="L282"></a>for chain := top.lastChain; chain.up != nil; chain = chain.up {
        <a id="L283"></a><span class="comment">// chain.leafCount gives the number of literals requiring at least &#34;bits&#34;</span>
        <a id="L284"></a><span class="comment">// bits to encode.</span>
        <a id="L285"></a>bitCount[bits] = chain.leafCount - chain.up.leafCount;
        <a id="L286"></a>bits++;
    <a id="L287"></a>}
    <a id="L288"></a>return bitCount;
<a id="L289"></a>}

<a id="L291"></a><span class="comment">// Look at the leaves and assign them a bit count and an encoding as specified</span>
<a id="L292"></a><span class="comment">// in RFC 1951 3.2.2</span>
<a id="L293"></a>func (h *huffmanEncoder) assignEncodingAndSize(bitCount []int32, list []literalNode) {
    <a id="L294"></a>code := uint16(0);
    <a id="L295"></a>for n, bits := range bitCount {
        <a id="L296"></a>code &lt;&lt;= 1;
        <a id="L297"></a>if n == 0 || bits == 0 {
            <a id="L298"></a>continue
        <a id="L299"></a>}
        <a id="L300"></a><span class="comment">// The literals list[len(list)-bits] .. list[len(list)-bits]</span>
        <a id="L301"></a><span class="comment">// are encoded using &#34;bits&#34; bits, and get the values</span>
        <a id="L302"></a><span class="comment">// code, code + 1, ....  The code values are</span>
        <a id="L303"></a><span class="comment">// assigned in literal order (not frequency order).</span>
        <a id="L304"></a>chunk := list[len(list)-int(bits) : len(list)];
        <a id="L305"></a>sortByLiteral(chunk);
        <a id="L306"></a>for _, node := range chunk {
            <a id="L307"></a>h.codeBits[node.literal] = uint8(n);
            <a id="L308"></a>h.code[node.literal] = reverseBits(code, uint8(n));
            <a id="L309"></a>code++;
        <a id="L310"></a>}
        <a id="L311"></a>list = list[0 : len(list)-int(bits)];
    <a id="L312"></a>}
<a id="L313"></a>}

<a id="L315"></a><span class="comment">// Update this Huffman Code object to be the minimum code for the specified frequency count.</span>
<a id="L316"></a><span class="comment">//</span>
<a id="L317"></a><span class="comment">// freq  An array of frequencies, in which frequency[i] gives the frequency of literal i.</span>
<a id="L318"></a><span class="comment">// maxBits  The maximum number of bits to use for any literal.</span>
<a id="L319"></a>func (h *huffmanEncoder) generate(freq []int32, maxBits int32) {
    <a id="L320"></a>list := make([]literalNode, len(freq)+1);
    <a id="L321"></a><span class="comment">// Number of non-zero literals</span>
    <a id="L322"></a>count := 0;
    <a id="L323"></a><span class="comment">// Set list to be the set of all non-zero literals and their frequencies</span>
    <a id="L324"></a>for i, f := range freq {
        <a id="L325"></a>if f != 0 {
            <a id="L326"></a>list[count] = literalNode{uint16(i), f};
            <a id="L327"></a>count++;
        <a id="L328"></a>} else {
            <a id="L329"></a>h.codeBits[i] = 0
        <a id="L330"></a>}
    <a id="L331"></a>}
    <a id="L332"></a><span class="comment">// If freq[] is shorter than codeBits[], fill rest of codeBits[] with zeros</span>
    <a id="L333"></a>h.codeBits = h.codeBits[0:len(freq)];
    <a id="L334"></a>list = list[0:count];
    <a id="L335"></a>if count &lt;= 2 {
        <a id="L336"></a><span class="comment">// Handle the small cases here, because they are awkward for the general case code.  With</span>
        <a id="L337"></a><span class="comment">// two or fewer literals, everything has bit length 1.</span>
        <a id="L338"></a>for i, node := range list {
            <a id="L339"></a><span class="comment">// &#34;list&#34; is in order of increasing literal value.</span>
            <a id="L340"></a>h.codeBits[node.literal] = 1;
            <a id="L341"></a>h.code[node.literal] = uint16(i);
        <a id="L342"></a>}
        <a id="L343"></a>return;
    <a id="L344"></a>}
    <a id="L345"></a>sortByFreq(list);

    <a id="L347"></a><span class="comment">// Get the number of literals for each bit count</span>
    <a id="L348"></a>bitCount := h.bitCounts(list, maxBits);
    <a id="L349"></a><span class="comment">// And do the assignment</span>
    <a id="L350"></a>h.assignEncodingAndSize(bitCount, list);
<a id="L351"></a>}

<a id="L353"></a>type literalNodeSorter struct {
    <a id="L354"></a>a    []literalNode;
    <a id="L355"></a>less func(i, j int) bool;
<a id="L356"></a>}

<a id="L358"></a>func (s literalNodeSorter) Len() int { return len(s.a) }

<a id="L360"></a>func (s literalNodeSorter) Less(i, j int) bool {
    <a id="L361"></a>return s.less(i, j)
<a id="L362"></a>}

<a id="L364"></a>func (s literalNodeSorter) Swap(i, j int) { s.a[i], s.a[j] = s.a[j], s.a[i] }

<a id="L366"></a>func sortByFreq(a []literalNode) {
    <a id="L367"></a>s := &amp;literalNodeSorter{a, func(i, j int) bool { return a[i].freq &lt; a[j].freq }};
    <a id="L368"></a>sort.Sort(s);
<a id="L369"></a>}

<a id="L371"></a>func sortByLiteral(a []literalNode) {
    <a id="L372"></a>s := &amp;literalNodeSorter{a, func(i, j int) bool { return a[i].literal &lt; a[j].literal }};
    <a id="L373"></a>sort.Sort(s);
<a id="L374"></a>}
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
