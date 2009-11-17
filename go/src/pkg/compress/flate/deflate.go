<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/compress/flate/deflate.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/compress/flate/deflate.go</h1>

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
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;math&#34;;
    <a id="L11"></a>&#34;os&#34;;
<a id="L12"></a>)

<a id="L14"></a>const (
    <a id="L15"></a>NoCompression        = 0;
    <a id="L16"></a>BestSpeed            = 1;
    <a id="L17"></a>fastCompression      = 3;
    <a id="L18"></a>BestCompression      = 9;
    <a id="L19"></a>DefaultCompression   = -1;
    <a id="L20"></a>logMaxOffsetSize     = 15;  <span class="comment">// Standard DEFLATE</span>
    <a id="L21"></a>wideLogMaxOffsetSize = 22;  <span class="comment">// Wide DEFLATE</span>
    <a id="L22"></a>minMatchLength       = 3;   <span class="comment">// The smallest match that the deflater looks for</span>
    <a id="L23"></a>maxMatchLength       = 258; <span class="comment">// The longest match for the deflater</span>
    <a id="L24"></a>minOffsetSize        = 1;   <span class="comment">// The shortest offset that makes any sence</span>

    <a id="L26"></a><span class="comment">// The maximum number of tokens we put into a single flat block, just too</span>
    <a id="L27"></a><span class="comment">// stop things from getting too large.</span>
    <a id="L28"></a>maxFlateBlockTokens = 1 &lt;&lt; 14;
    <a id="L29"></a>maxStoreBlockSize   = 65535;
    <a id="L30"></a>hashBits            = 15;
    <a id="L31"></a>hashSize            = 1 &lt;&lt; hashBits;
    <a id="L32"></a>hashMask            = (1 &lt;&lt; hashBits) - 1;
    <a id="L33"></a>hashShift           = (hashBits + minMatchLength - 1) / minMatchLength;
<a id="L34"></a>)

<a id="L36"></a>type syncPipeReader struct {
    <a id="L37"></a>*io.PipeReader;
    <a id="L38"></a>closeChan chan bool;
<a id="L39"></a>}

<a id="L41"></a>func (sr *syncPipeReader) CloseWithError(err os.Error) os.Error {
    <a id="L42"></a>retErr := sr.PipeReader.CloseWithError(err);
    <a id="L43"></a>sr.closeChan &lt;- true; <span class="comment">// finish writer close</span>
    <a id="L44"></a>return retErr;
<a id="L45"></a>}

<a id="L47"></a>type syncPipeWriter struct {
    <a id="L48"></a>*io.PipeWriter;
    <a id="L49"></a>closeChan chan bool;
<a id="L50"></a>}

<a id="L52"></a>type compressionLevel struct {
    <a id="L53"></a>good, lazy, nice, chain, fastSkipHashing int;
<a id="L54"></a>}

<a id="L56"></a>var levels = []compressionLevel{
    <a id="L57"></a>compressionLevel{}, <span class="comment">// 0</span>
    <a id="L58"></a><span class="comment">// For levels 1-3 we don&#39;t bother trying with lazy matches</span>
    <a id="L59"></a>compressionLevel{3, 0, 8, 4, 4},
    <a id="L60"></a>compressionLevel{3, 0, 16, 8, 5},
    <a id="L61"></a>compressionLevel{3, 0, 32, 32, 6},
    <a id="L62"></a><span class="comment">// Levels 4-9 use increasingly more lazy matching</span>
    <a id="L63"></a><span class="comment">// and increasingly stringent conditions for &#34;good enough&#34;.</span>
    <a id="L64"></a>compressionLevel{4, 4, 16, 16, math.MaxInt32},
    <a id="L65"></a>compressionLevel{8, 16, 32, 32, math.MaxInt32},
    <a id="L66"></a>compressionLevel{8, 16, 128, 128, math.MaxInt32},
    <a id="L67"></a>compressionLevel{8, 32, 128, 256, math.MaxInt32},
    <a id="L68"></a>compressionLevel{32, 128, 258, 1024, math.MaxInt32},
    <a id="L69"></a>compressionLevel{32, 258, 258, 4096, math.MaxInt32},
<a id="L70"></a>}

<a id="L72"></a>func (sw *syncPipeWriter) Close() os.Error {
    <a id="L73"></a>err := sw.PipeWriter.Close();
    <a id="L74"></a>&lt;-sw.closeChan; <span class="comment">// wait for reader close</span>
    <a id="L75"></a>return err;
<a id="L76"></a>}

<a id="L78"></a>func syncPipe() (*syncPipeReader, *syncPipeWriter) {
    <a id="L79"></a>r, w := io.Pipe();
    <a id="L80"></a>sr := &amp;syncPipeReader{r, make(chan bool, 1)};
    <a id="L81"></a>sw := &amp;syncPipeWriter{w, sr.closeChan};
    <a id="L82"></a>return sr, sw;
<a id="L83"></a>}

<a id="L85"></a>type deflater struct {
    <a id="L86"></a>level         int;
    <a id="L87"></a>logWindowSize uint;
    <a id="L88"></a>w             *huffmanBitWriter;
    <a id="L89"></a>r             io.Reader;
    <a id="L90"></a><span class="comment">// (1 &lt;&lt; logWindowSize) - 1.</span>
    <a id="L91"></a>windowMask int;

    <a id="L93"></a><span class="comment">// hashHead[hashValue] contains the largest inputIndex with the specified hash value</span>
    <a id="L94"></a>hashHead []int;

    <a id="L96"></a><span class="comment">// If hashHead[hashValue] is within the current window, then</span>
    <a id="L97"></a><span class="comment">// hashPrev[hashHead[hashValue] &amp; windowMask] contains the previous index</span>
    <a id="L98"></a><span class="comment">// with the same hash value.</span>
    <a id="L99"></a>hashPrev []int;

    <a id="L101"></a><span class="comment">// If we find a match of length &gt;= niceMatch, then we don&#39;t bother searching</span>
    <a id="L102"></a><span class="comment">// any further.</span>
    <a id="L103"></a>niceMatch int;

    <a id="L105"></a><span class="comment">// If we find a match of length &gt;= goodMatch, we only do a half-hearted</span>
    <a id="L106"></a><span class="comment">// effort at doing lazy matching starting at the next character</span>
    <a id="L107"></a>goodMatch int;

    <a id="L109"></a><span class="comment">// The maximum number of chains we look at when finding a match</span>
    <a id="L110"></a>maxChainLength int;

    <a id="L112"></a><span class="comment">// The sliding window we use for matching</span>
    <a id="L113"></a>window []byte;

    <a id="L115"></a><span class="comment">// The index just past the last valid character</span>
    <a id="L116"></a>windowEnd int;

    <a id="L118"></a><span class="comment">// index in &#34;window&#34; at which current block starts</span>
    <a id="L119"></a>blockStart int;
<a id="L120"></a>}

<a id="L122"></a>func (d *deflater) flush() os.Error {
    <a id="L123"></a>d.w.flush();
    <a id="L124"></a>return d.w.err;
<a id="L125"></a>}

<a id="L127"></a>func (d *deflater) fillWindow(index int) (int, os.Error) {
    <a id="L128"></a>wSize := d.windowMask + 1;
    <a id="L129"></a>if index &gt;= wSize+wSize-(minMatchLength+maxMatchLength) {
        <a id="L130"></a><span class="comment">// shift the window by wSize</span>
        <a id="L131"></a>bytes.Copy(d.window, d.window[wSize:2*wSize]);
        <a id="L132"></a>index -= wSize;
        <a id="L133"></a>d.windowEnd -= wSize;
        <a id="L134"></a>if d.blockStart &gt;= wSize {
            <a id="L135"></a>d.blockStart -= wSize
        <a id="L136"></a>} else {
            <a id="L137"></a>d.blockStart = math.MaxInt32
        <a id="L138"></a>}
        <a id="L139"></a>for i, h := range d.hashHead {
            <a id="L140"></a>d.hashHead[i] = max(h-wSize, -1)
        <a id="L141"></a>}
        <a id="L142"></a>for i, h := range d.hashPrev {
            <a id="L143"></a>d.hashPrev[i] = max(h-wSize, -1)
        <a id="L144"></a>}
    <a id="L145"></a>}
    <a id="L146"></a>var count int;
    <a id="L147"></a>var err os.Error;
    <a id="L148"></a>count, err = io.ReadAtLeast(d.r, d.window[d.windowEnd:len(d.window)], 1);
    <a id="L149"></a>d.windowEnd += count;
    <a id="L150"></a>if err == os.EOF {
        <a id="L151"></a>return index, nil
    <a id="L152"></a>}
    <a id="L153"></a>return index, err;
<a id="L154"></a>}

<a id="L156"></a>func (d *deflater) writeBlock(tokens []token, index int, eof bool) os.Error {
    <a id="L157"></a>if index &gt; 0 || eof {
        <a id="L158"></a>var window []byte;
        <a id="L159"></a>if d.blockStart &lt;= index {
            <a id="L160"></a>window = d.window[d.blockStart:index]
        <a id="L161"></a>}
        <a id="L162"></a>d.blockStart = index;
        <a id="L163"></a>d.w.writeBlock(tokens, eof, window);
        <a id="L164"></a>return d.w.err;
    <a id="L165"></a>}
    <a id="L166"></a>return nil;
<a id="L167"></a>}

<a id="L169"></a><span class="comment">// Try to find a match starting at index whose length is greater than prevSize.</span>
<a id="L170"></a><span class="comment">// We only look at chainCount possibilities before giving up.</span>
<a id="L171"></a>func (d *deflater) findMatch(pos int, prevHead int, prevLength int, lookahead int) (length, offset int, ok bool) {
    <a id="L172"></a>win := d.window[0 : pos+min(maxMatchLength, lookahead)];

    <a id="L174"></a><span class="comment">// We quit when we get a match that&#39;s at least nice long</span>
    <a id="L175"></a>nice := min(d.niceMatch, len(win)-pos);

    <a id="L177"></a><span class="comment">// If we&#39;ve got a match that&#39;s good enough, only look in 1/4 the chain.</span>
    <a id="L178"></a>tries := d.maxChainLength;
    <a id="L179"></a>length = prevLength;
    <a id="L180"></a>if length &gt;= d.goodMatch {
        <a id="L181"></a>tries &gt;&gt;= 2
    <a id="L182"></a>}

    <a id="L184"></a>w0 := win[pos];
    <a id="L185"></a>w1 := win[pos+1];
    <a id="L186"></a>wEnd := win[pos+length];
    <a id="L187"></a>minIndex := pos - (d.windowMask + 1);

    <a id="L189"></a>for i := prevHead; tries &gt; 0; tries-- {
        <a id="L190"></a>if w0 == win[i] &amp;&amp; w1 == win[i+1] &amp;&amp; wEnd == win[i+length] {
            <a id="L191"></a><span class="comment">// The hash function ensures that if win[i] and win[i+1] match, win[i+2] matches</span>

            <a id="L193"></a>n := 3;
            <a id="L194"></a>for pos+n &lt; len(win) &amp;&amp; win[i+n] == win[pos+n] {
                <a id="L195"></a>n++
            <a id="L196"></a>}
            <a id="L197"></a>if n &gt; length &amp;&amp; (n &gt; 3 || pos-i &lt;= 4096) {
                <a id="L198"></a>length = n;
                <a id="L199"></a>offset = pos - i;
                <a id="L200"></a>ok = true;
                <a id="L201"></a>if n &gt;= nice {
                    <a id="L202"></a><span class="comment">// The match is good enough that we don&#39;t try to find a better one.</span>
                    <a id="L203"></a>break
                <a id="L204"></a>}
                <a id="L205"></a>wEnd = win[pos+n];
            <a id="L206"></a>}
        <a id="L207"></a>}
        <a id="L208"></a>if i == minIndex {
            <a id="L209"></a><span class="comment">// hashPrev[i &amp; windowMask] has already been overwritten, so stop now.</span>
            <a id="L210"></a>break
        <a id="L211"></a>}
        <a id="L212"></a>if i = d.hashPrev[i&amp;d.windowMask]; i &lt; minIndex || i &lt; 0 {
            <a id="L213"></a>break
        <a id="L214"></a>}
    <a id="L215"></a>}
    <a id="L216"></a>return;
<a id="L217"></a>}

<a id="L219"></a>func (d *deflater) writeStoredBlock(buf []byte) os.Error {
    <a id="L220"></a>if d.w.writeStoredHeader(len(buf), false); d.w.err != nil {
        <a id="L221"></a>return d.w.err
    <a id="L222"></a>}
    <a id="L223"></a>d.w.writeBytes(buf);
    <a id="L224"></a>return d.w.err;
<a id="L225"></a>}

<a id="L227"></a>func (d *deflater) storedDeflate() os.Error {
    <a id="L228"></a>buf := make([]byte, maxStoreBlockSize);
    <a id="L229"></a>for {
        <a id="L230"></a>n, err := d.r.Read(buf);
        <a id="L231"></a>if n &gt; 0 {
            <a id="L232"></a>if err := d.writeStoredBlock(buf[0:n]); err != nil {
                <a id="L233"></a>return err
            <a id="L234"></a>}
        <a id="L235"></a>}
        <a id="L236"></a>if err != nil {
            <a id="L237"></a>if err == os.EOF {
                <a id="L238"></a>break
            <a id="L239"></a>}
            <a id="L240"></a>return err;
        <a id="L241"></a>}
    <a id="L242"></a>}
    <a id="L243"></a>return nil;
<a id="L244"></a>}

<a id="L246"></a>func (d *deflater) doDeflate() (err os.Error) {
    <a id="L247"></a><span class="comment">// init</span>
    <a id="L248"></a>d.windowMask = 1&lt;&lt;d.logWindowSize - 1;
    <a id="L249"></a>d.hashHead = make([]int, hashSize);
    <a id="L250"></a>d.hashPrev = make([]int, 1&lt;&lt;d.logWindowSize);
    <a id="L251"></a>d.window = make([]byte, 2&lt;&lt;d.logWindowSize);
    <a id="L252"></a>fillInts(d.hashHead, -1);
    <a id="L253"></a>tokens := make([]token, maxFlateBlockTokens, maxFlateBlockTokens+1);
    <a id="L254"></a>l := levels[d.level];
    <a id="L255"></a>d.goodMatch = l.good;
    <a id="L256"></a>d.niceMatch = l.nice;
    <a id="L257"></a>d.maxChainLength = l.chain;
    <a id="L258"></a>lazyMatch := l.lazy;
    <a id="L259"></a>length := minMatchLength - 1;
    <a id="L260"></a>offset := 0;
    <a id="L261"></a>byteAvailable := false;
    <a id="L262"></a>isFastDeflate := l.fastSkipHashing != 0;
    <a id="L263"></a>index := 0;
    <a id="L264"></a><span class="comment">// run</span>
    <a id="L265"></a>if index, err = d.fillWindow(index); err != nil {
        <a id="L266"></a>return
    <a id="L267"></a>}
    <a id="L268"></a>maxOffset := d.windowMask + 1; <span class="comment">// (1 &lt;&lt; logWindowSize);</span>
    <a id="L269"></a><span class="comment">// only need to change when you refill the window</span>
    <a id="L270"></a>windowEnd := d.windowEnd;
    <a id="L271"></a>maxInsertIndex := windowEnd - (minMatchLength - 1);
    <a id="L272"></a>ti := 0;

    <a id="L274"></a>hash := int(0);
    <a id="L275"></a>if index &lt; maxInsertIndex {
        <a id="L276"></a>hash = int(d.window[index])&lt;&lt;hashShift + int(d.window[index+1])
    <a id="L277"></a>}
    <a id="L278"></a>chainHead := -1;
    <a id="L279"></a>for {
        <a id="L280"></a>if index &gt; windowEnd {
            <a id="L281"></a>panic(&#34;index &gt; windowEnd&#34;)
        <a id="L282"></a>}
        <a id="L283"></a>lookahead := windowEnd - index;
        <a id="L284"></a>if lookahead &lt; minMatchLength+maxMatchLength {
            <a id="L285"></a>if index, err = d.fillWindow(index); err != nil {
                <a id="L286"></a>return
            <a id="L287"></a>}
            <a id="L288"></a>windowEnd = d.windowEnd;
            <a id="L289"></a>if index &gt; windowEnd {
                <a id="L290"></a>panic(&#34;index &gt; windowEnd&#34;)
            <a id="L291"></a>}
            <a id="L292"></a>maxInsertIndex = windowEnd - (minMatchLength - 1);
            <a id="L293"></a>lookahead = windowEnd - index;
            <a id="L294"></a>if lookahead == 0 {
                <a id="L295"></a>break
            <a id="L296"></a>}
        <a id="L297"></a>}
        <a id="L298"></a>if index &lt; maxInsertIndex {
            <a id="L299"></a><span class="comment">// Update the hash</span>
            <a id="L300"></a>hash = (hash&lt;&lt;hashShift + int(d.window[index+2])) &amp; hashMask;
            <a id="L301"></a>chainHead = d.hashHead[hash];
            <a id="L302"></a>d.hashPrev[index&amp;d.windowMask] = chainHead;
            <a id="L303"></a>d.hashHead[hash] = index;
        <a id="L304"></a>}
        <a id="L305"></a>prevLength := length;
        <a id="L306"></a>prevOffset := offset;
        <a id="L307"></a>minIndex := max(index-maxOffset, 0);
        <a id="L308"></a>length = minMatchLength - 1;
        <a id="L309"></a>offset = 0;

        <a id="L311"></a>if chainHead &gt;= minIndex &amp;&amp;
            <a id="L312"></a>(isFastDeflate &amp;&amp; lookahead &gt; minMatchLength-1 ||
                <a id="L313"></a>!isFastDeflate &amp;&amp; lookahead &gt; prevLength &amp;&amp; prevLength &lt; lazyMatch) {
            <a id="L314"></a>if newLength, newOffset, ok := d.findMatch(index, chainHead, minMatchLength-1, lookahead); ok {
                <a id="L315"></a>length = newLength;
                <a id="L316"></a>offset = newOffset;
            <a id="L317"></a>}
        <a id="L318"></a>}
        <a id="L319"></a>if isFastDeflate &amp;&amp; length &gt;= minMatchLength ||
            <a id="L320"></a>!isFastDeflate &amp;&amp; prevLength &gt;= minMatchLength &amp;&amp; length &lt;= prevLength {
            <a id="L321"></a><span class="comment">// There was a match at the previous step, and the current match is</span>
            <a id="L322"></a><span class="comment">// not better. Output the previous match.</span>
            <a id="L323"></a>if isFastDeflate {
                <a id="L324"></a>tokens[ti] = matchToken(uint32(length-minMatchLength), uint32(offset-minOffsetSize))
            <a id="L325"></a>} else {
                <a id="L326"></a>tokens[ti] = matchToken(uint32(prevLength-minMatchLength), uint32(prevOffset-minOffsetSize))
            <a id="L327"></a>}
            <a id="L328"></a>ti++;
            <a id="L329"></a><span class="comment">// Insert in the hash table all strings up to the end of the match.</span>
            <a id="L330"></a><span class="comment">// index and index-1 are already inserted. If there is not enough</span>
            <a id="L331"></a><span class="comment">// lookahead, the last two strings are not inserted into the hash</span>
            <a id="L332"></a><span class="comment">// table.</span>
            <a id="L333"></a>if length &lt;= l.fastSkipHashing {
                <a id="L334"></a>var newIndex int;
                <a id="L335"></a>if isFastDeflate {
                    <a id="L336"></a>newIndex = index + length
                <a id="L337"></a>} else {
                    <a id="L338"></a>newIndex = prevLength - 1
                <a id="L339"></a>}
                <a id="L340"></a>for index++; index &lt; newIndex; index++ {
                    <a id="L341"></a>if index &lt; maxInsertIndex {
                        <a id="L342"></a>hash = (hash&lt;&lt;hashShift + int(d.window[index+2])) &amp; hashMask;
                        <a id="L343"></a><span class="comment">// Get previous value with the same hash.</span>
                        <a id="L344"></a><span class="comment">// Our chain should point to the previous value.</span>
                        <a id="L345"></a>d.hashPrev[index&amp;d.windowMask] = d.hashHead[hash];
                        <a id="L346"></a><span class="comment">// Set the head of the hash chain to us.</span>
                        <a id="L347"></a>d.hashHead[hash] = index;
                    <a id="L348"></a>}
                <a id="L349"></a>}
                <a id="L350"></a>if !isFastDeflate {
                    <a id="L351"></a>byteAvailable = false;
                    <a id="L352"></a>length = minMatchLength - 1;
                <a id="L353"></a>}
            <a id="L354"></a>} else {
                <a id="L355"></a><span class="comment">// For matches this long, we don&#39;t bother inserting each individual</span>
                <a id="L356"></a><span class="comment">// item into the table.</span>
                <a id="L357"></a>index += length;
                <a id="L358"></a>hash = (int(d.window[index])&lt;&lt;hashShift + int(d.window[index+1]));
            <a id="L359"></a>}
            <a id="L360"></a>if ti == maxFlateBlockTokens {
                <a id="L361"></a><span class="comment">// The block includes the current character</span>
                <a id="L362"></a>if err = d.writeBlock(tokens, index, false); err != nil {
                    <a id="L363"></a>return
                <a id="L364"></a>}
                <a id="L365"></a>ti = 0;
            <a id="L366"></a>}
        <a id="L367"></a>} else {
            <a id="L368"></a>if isFastDeflate || byteAvailable {
                <a id="L369"></a>i := index - 1;
                <a id="L370"></a>if isFastDeflate {
                    <a id="L371"></a>i = index
                <a id="L372"></a>}
                <a id="L373"></a>tokens[ti] = literalToken(uint32(d.window[i]) &amp; 0xFF);
                <a id="L374"></a>ti++;
                <a id="L375"></a>if ti == maxFlateBlockTokens {
                    <a id="L376"></a>if err = d.writeBlock(tokens, i+1, false); err != nil {
                        <a id="L377"></a>return
                    <a id="L378"></a>}
                    <a id="L379"></a>ti = 0;
                <a id="L380"></a>}
            <a id="L381"></a>}
            <a id="L382"></a>index++;
            <a id="L383"></a>if !isFastDeflate {
                <a id="L384"></a>byteAvailable = true
            <a id="L385"></a>}
        <a id="L386"></a>}

    <a id="L388"></a>}
    <a id="L389"></a>if byteAvailable {
        <a id="L390"></a><span class="comment">// There is still one pending token that needs to be flushed</span>
        <a id="L391"></a>tokens[ti] = literalToken(uint32(d.window[index-1]) &amp; 0xFF);
        <a id="L392"></a>ti++;
    <a id="L393"></a>}

    <a id="L395"></a>if ti &gt; 0 {
        <a id="L396"></a>if err = d.writeBlock(tokens[0:ti], index, false); err != nil {
            <a id="L397"></a>return
        <a id="L398"></a>}
    <a id="L399"></a>}
    <a id="L400"></a>return;
<a id="L401"></a>}

<a id="L403"></a>func (d *deflater) deflater(r io.Reader, w io.Writer, level int, logWindowSize uint) (err os.Error) {
    <a id="L404"></a>d.r = r;
    <a id="L405"></a>d.w = newHuffmanBitWriter(w);
    <a id="L406"></a>d.level = level;
    <a id="L407"></a>d.logWindowSize = logWindowSize;

    <a id="L409"></a>switch {
    <a id="L410"></a>case level == NoCompression:
        <a id="L411"></a>err = d.storedDeflate()
    <a id="L412"></a>case level == DefaultCompression:
        <a id="L413"></a>d.level = 6;
        <a id="L414"></a>fallthrough;
    <a id="L415"></a>case 1 &lt;= level &amp;&amp; level &lt;= 9:
        <a id="L416"></a>err = d.doDeflate()
    <a id="L417"></a>default:
        <a id="L418"></a>return WrongValueError{&#34;level&#34;, 0, 9, int32(level)}
    <a id="L419"></a>}

    <a id="L421"></a>if err != nil {
        <a id="L422"></a>return err
    <a id="L423"></a>}
    <a id="L424"></a>if d.w.writeStoredHeader(0, true); d.w.err != nil {
        <a id="L425"></a>return d.w.err
    <a id="L426"></a>}
    <a id="L427"></a>return d.flush();
<a id="L428"></a>}

<a id="L430"></a>func newDeflater(w io.Writer, level int, logWindowSize uint) io.WriteCloser {
    <a id="L431"></a>var d deflater;
    <a id="L432"></a>pr, pw := syncPipe();
    <a id="L433"></a>go func() {
        <a id="L434"></a>err := d.deflater(pr, w, level, logWindowSize);
        <a id="L435"></a>pr.CloseWithError(err);
    <a id="L436"></a>}();
    <a id="L437"></a>return pw;
<a id="L438"></a>}

<a id="L440"></a>func NewDeflater(w io.Writer, level int) io.WriteCloser {
    <a id="L441"></a>return newDeflater(w, level, logMaxOffsetSize)
<a id="L442"></a>}
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
