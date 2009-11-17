<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/gob/decode.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/gob/decode.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package gob

<a id="L7"></a><span class="comment">// TODO(rsc): When garbage collector changes, revisit</span>
<a id="L8"></a><span class="comment">// the allocations in this file that use unsafe.Pointer.</span>

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bytes&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;math&#34;;
    <a id="L14"></a>&#34;os&#34;;
    <a id="L15"></a>&#34;reflect&#34;;
    <a id="L16"></a>&#34;unsafe&#34;;
<a id="L17"></a>)

<a id="L19"></a>var (
    <a id="L20"></a>errBadUint   = os.ErrorString(&#34;gob: encoded unsigned integer out of range&#34;);
    <a id="L21"></a>errBadType   = os.ErrorString(&#34;gob: unknown type id or corrupted data&#34;);
    <a id="L22"></a>errRange     = os.ErrorString(&#34;gob: internal error: field numbers out of bounds&#34;);
    <a id="L23"></a>errNotStruct = os.ErrorString(&#34;gob: TODO: can only handle structs&#34;);
<a id="L24"></a>)

<a id="L26"></a><span class="comment">// The global execution state of an instance of the decoder.</span>
<a id="L27"></a>type decodeState struct {
    <a id="L28"></a>b        *bytes.Buffer;
    <a id="L29"></a>err      os.Error;
    <a id="L30"></a>fieldnum int; <span class="comment">// the last field number read.</span>
    <a id="L31"></a>buf      []byte;
<a id="L32"></a>}

<a id="L34"></a>func newDecodeState(b *bytes.Buffer) *decodeState {
    <a id="L35"></a>d := new(decodeState);
    <a id="L36"></a>d.b = b;
    <a id="L37"></a>d.buf = make([]byte, uint64Size);
    <a id="L38"></a>return d;
<a id="L39"></a>}

<a id="L41"></a>func overflow(name string) os.ErrorString {
    <a id="L42"></a>return os.ErrorString(`value for &#34;` + name + `&#34; out of range`)
<a id="L43"></a>}

<a id="L45"></a><span class="comment">// decodeUintReader reads an encoded unsigned integer from an io.Reader.</span>
<a id="L46"></a><span class="comment">// Used only by the Decoder to read the message length.</span>
<a id="L47"></a>func decodeUintReader(r io.Reader, buf []byte) (x uint64, err os.Error) {
    <a id="L48"></a>_, err = r.Read(buf[0:1]);
    <a id="L49"></a>if err != nil {
        <a id="L50"></a>return
    <a id="L51"></a>}
    <a id="L52"></a>b := buf[0];
    <a id="L53"></a>if b &lt;= 0x7f {
        <a id="L54"></a>return uint64(b), nil
    <a id="L55"></a>}
    <a id="L56"></a>nb := -int(int8(b));
    <a id="L57"></a>if nb &gt; uint64Size {
        <a id="L58"></a>err = errBadUint;
        <a id="L59"></a>return;
    <a id="L60"></a>}
    <a id="L61"></a>var n int;
    <a id="L62"></a>n, err = io.ReadFull(r, buf[0:nb]);
    <a id="L63"></a>if err != nil {
        <a id="L64"></a>if err == os.EOF {
            <a id="L65"></a>err = io.ErrUnexpectedEOF
        <a id="L66"></a>}
        <a id="L67"></a>return;
    <a id="L68"></a>}
    <a id="L69"></a><span class="comment">// Could check that the high byte is zero but it&#39;s not worth it.</span>
    <a id="L70"></a>for i := 0; i &lt; n; i++ {
        <a id="L71"></a>x &lt;&lt;= 8;
        <a id="L72"></a>x |= uint64(buf[i]);
    <a id="L73"></a>}
    <a id="L74"></a>return;
<a id="L75"></a>}

<a id="L77"></a><span class="comment">// decodeUint reads an encoded unsigned integer from state.r.</span>
<a id="L78"></a><span class="comment">// Sets state.err.  If state.err is already non-nil, it does nothing.</span>
<a id="L79"></a><span class="comment">// Does not check for overflow.</span>
<a id="L80"></a>func decodeUint(state *decodeState) (x uint64) {
    <a id="L81"></a>if state.err != nil {
        <a id="L82"></a>return
    <a id="L83"></a>}
    <a id="L84"></a>var b uint8;
    <a id="L85"></a>b, state.err = state.b.ReadByte();
    <a id="L86"></a>if b &lt;= 0x7f { <span class="comment">// includes state.err != nil</span>
        <a id="L87"></a>return uint64(b)
    <a id="L88"></a>}
    <a id="L89"></a>nb := -int(int8(b));
    <a id="L90"></a>if nb &gt; uint64Size {
        <a id="L91"></a>state.err = errBadUint;
        <a id="L92"></a>return;
    <a id="L93"></a>}
    <a id="L94"></a>var n int;
    <a id="L95"></a>n, state.err = state.b.Read(state.buf[0:nb]);
    <a id="L96"></a><span class="comment">// Don&#39;t need to check error; it&#39;s safe to loop regardless.</span>
    <a id="L97"></a><span class="comment">// Could check that the high byte is zero but it&#39;s not worth it.</span>
    <a id="L98"></a>for i := 0; i &lt; n; i++ {
        <a id="L99"></a>x &lt;&lt;= 8;
        <a id="L100"></a>x |= uint64(state.buf[i]);
    <a id="L101"></a>}
    <a id="L102"></a>return x;
<a id="L103"></a>}

<a id="L105"></a><span class="comment">// decodeInt reads an encoded signed integer from state.r.</span>
<a id="L106"></a><span class="comment">// Sets state.err.  If state.err is already non-nil, it does nothing.</span>
<a id="L107"></a><span class="comment">// Does not check for overflow.</span>
<a id="L108"></a>func decodeInt(state *decodeState) int64 {
    <a id="L109"></a>x := decodeUint(state);
    <a id="L110"></a>if state.err != nil {
        <a id="L111"></a>return 0
    <a id="L112"></a>}
    <a id="L113"></a>if x&amp;1 != 0 {
        <a id="L114"></a>return ^int64(x &gt;&gt; 1)
    <a id="L115"></a>}
    <a id="L116"></a>return int64(x &gt;&gt; 1);
<a id="L117"></a>}

<a id="L119"></a>type decOp func(i *decInstr, state *decodeState, p unsafe.Pointer)

<a id="L121"></a><span class="comment">// The &#39;instructions&#39; of the decoding machine</span>
<a id="L122"></a>type decInstr struct {
    <a id="L123"></a>op     decOp;
    <a id="L124"></a>field  int;            <span class="comment">// field number of the wire type</span>
    <a id="L125"></a>indir  int;            <span class="comment">// how many pointer indirections to reach the value in the struct</span>
    <a id="L126"></a>offset uintptr;        <span class="comment">// offset in the structure of the field to encode</span>
    <a id="L127"></a>ovfl   os.ErrorString; <span class="comment">// error message for overflow/underflow (for arrays, of the elements)</span>
<a id="L128"></a>}

<a id="L130"></a><span class="comment">// Since the encoder writes no zeros, if we arrive at a decoder we have</span>
<a id="L131"></a><span class="comment">// a value to extract and store.  The field number has already been read</span>
<a id="L132"></a><span class="comment">// (it&#39;s how we knew to call this decoder).</span>
<a id="L133"></a><span class="comment">// Each decoder is responsible for handling any indirections associated</span>
<a id="L134"></a><span class="comment">// with the data structure.  If any pointer so reached is nil, allocation must</span>
<a id="L135"></a><span class="comment">// be done.</span>

<a id="L137"></a><span class="comment">// Walk the pointer hierarchy, allocating if we find a nil.  Stop one before the end.</span>
<a id="L138"></a>func decIndirect(p unsafe.Pointer, indir int) unsafe.Pointer {
    <a id="L139"></a>for ; indir &gt; 1; indir-- {
        <a id="L140"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L141"></a><span class="comment">// Allocation required</span>
            <a id="L142"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(unsafe.Pointer))
        <a id="L143"></a>}
        <a id="L144"></a>p = *(*unsafe.Pointer)(p);
    <a id="L145"></a>}
    <a id="L146"></a>return p;
<a id="L147"></a>}

<a id="L149"></a>func ignoreUint(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L150"></a>decodeUint(state)
<a id="L151"></a>}

<a id="L153"></a>func decBool(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L154"></a>if i.indir &gt; 0 {
        <a id="L155"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L156"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(bool))
        <a id="L157"></a>}
        <a id="L158"></a>p = *(*unsafe.Pointer)(p);
    <a id="L159"></a>}
    <a id="L160"></a>*(*bool)(p) = decodeInt(state) != 0;
<a id="L161"></a>}

<a id="L163"></a>func decInt8(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L164"></a>if i.indir &gt; 0 {
        <a id="L165"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L166"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int8))
        <a id="L167"></a>}
        <a id="L168"></a>p = *(*unsafe.Pointer)(p);
    <a id="L169"></a>}
    <a id="L170"></a>v := decodeInt(state);
    <a id="L171"></a>if v &lt; math.MinInt8 || math.MaxInt8 &lt; v {
        <a id="L172"></a>state.err = i.ovfl
    <a id="L173"></a>} else {
        <a id="L174"></a>*(*int8)(p) = int8(v)
    <a id="L175"></a>}
<a id="L176"></a>}

<a id="L178"></a>func decUint8(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L179"></a>if i.indir &gt; 0 {
        <a id="L180"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L181"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(uint8))
        <a id="L182"></a>}
        <a id="L183"></a>p = *(*unsafe.Pointer)(p);
    <a id="L184"></a>}
    <a id="L185"></a>v := decodeUint(state);
    <a id="L186"></a>if math.MaxUint8 &lt; v {
        <a id="L187"></a>state.err = i.ovfl
    <a id="L188"></a>} else {
        <a id="L189"></a>*(*uint8)(p) = uint8(v)
    <a id="L190"></a>}
<a id="L191"></a>}

<a id="L193"></a>func decInt16(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L194"></a>if i.indir &gt; 0 {
        <a id="L195"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L196"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int16))
        <a id="L197"></a>}
        <a id="L198"></a>p = *(*unsafe.Pointer)(p);
    <a id="L199"></a>}
    <a id="L200"></a>v := decodeInt(state);
    <a id="L201"></a>if v &lt; math.MinInt16 || math.MaxInt16 &lt; v {
        <a id="L202"></a>state.err = i.ovfl
    <a id="L203"></a>} else {
        <a id="L204"></a>*(*int16)(p) = int16(v)
    <a id="L205"></a>}
<a id="L206"></a>}

<a id="L208"></a>func decUint16(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L209"></a>if i.indir &gt; 0 {
        <a id="L210"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L211"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(uint16))
        <a id="L212"></a>}
        <a id="L213"></a>p = *(*unsafe.Pointer)(p);
    <a id="L214"></a>}
    <a id="L215"></a>v := decodeUint(state);
    <a id="L216"></a>if math.MaxUint16 &lt; v {
        <a id="L217"></a>state.err = i.ovfl
    <a id="L218"></a>} else {
        <a id="L219"></a>*(*uint16)(p) = uint16(v)
    <a id="L220"></a>}
<a id="L221"></a>}

<a id="L223"></a>func decInt32(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L224"></a>if i.indir &gt; 0 {
        <a id="L225"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L226"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int32))
        <a id="L227"></a>}
        <a id="L228"></a>p = *(*unsafe.Pointer)(p);
    <a id="L229"></a>}
    <a id="L230"></a>v := decodeInt(state);
    <a id="L231"></a>if v &lt; math.MinInt32 || math.MaxInt32 &lt; v {
        <a id="L232"></a>state.err = i.ovfl
    <a id="L233"></a>} else {
        <a id="L234"></a>*(*int32)(p) = int32(v)
    <a id="L235"></a>}
<a id="L236"></a>}

<a id="L238"></a>func decUint32(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L239"></a>if i.indir &gt; 0 {
        <a id="L240"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L241"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(uint32))
        <a id="L242"></a>}
        <a id="L243"></a>p = *(*unsafe.Pointer)(p);
    <a id="L244"></a>}
    <a id="L245"></a>v := decodeUint(state);
    <a id="L246"></a>if math.MaxUint32 &lt; v {
        <a id="L247"></a>state.err = i.ovfl
    <a id="L248"></a>} else {
        <a id="L249"></a>*(*uint32)(p) = uint32(v)
    <a id="L250"></a>}
<a id="L251"></a>}

<a id="L253"></a>func decInt64(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L254"></a>if i.indir &gt; 0 {
        <a id="L255"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L256"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(int64))
        <a id="L257"></a>}
        <a id="L258"></a>p = *(*unsafe.Pointer)(p);
    <a id="L259"></a>}
    <a id="L260"></a>*(*int64)(p) = int64(decodeInt(state));
<a id="L261"></a>}

<a id="L263"></a>func decUint64(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L264"></a>if i.indir &gt; 0 {
        <a id="L265"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L266"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(uint64))
        <a id="L267"></a>}
        <a id="L268"></a>p = *(*unsafe.Pointer)(p);
    <a id="L269"></a>}
    <a id="L270"></a>*(*uint64)(p) = uint64(decodeUint(state));
<a id="L271"></a>}

<a id="L273"></a><span class="comment">// Floating-point numbers are transmitted as uint64s holding the bits</span>
<a id="L274"></a><span class="comment">// of the underlying representation.  They are sent byte-reversed, with</span>
<a id="L275"></a><span class="comment">// the exponent end coming out first, so integer floating point numbers</span>
<a id="L276"></a><span class="comment">// (for example) transmit more compactly.  This routine does the</span>
<a id="L277"></a><span class="comment">// unswizzling.</span>
<a id="L278"></a>func floatFromBits(u uint64) float64 {
    <a id="L279"></a>var v uint64;
    <a id="L280"></a>for i := 0; i &lt; 8; i++ {
        <a id="L281"></a>v &lt;&lt;= 8;
        <a id="L282"></a>v |= u &amp; 0xFF;
        <a id="L283"></a>u &gt;&gt;= 8;
    <a id="L284"></a>}
    <a id="L285"></a>return math.Float64frombits(v);
<a id="L286"></a>}

<a id="L288"></a>func decFloat32(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L289"></a>if i.indir &gt; 0 {
        <a id="L290"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L291"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(float32))
        <a id="L292"></a>}
        <a id="L293"></a>p = *(*unsafe.Pointer)(p);
    <a id="L294"></a>}
    <a id="L295"></a>v := floatFromBits(decodeUint(state));
    <a id="L296"></a>av := v;
    <a id="L297"></a>if av &lt; 0 {
        <a id="L298"></a>av = -av
    <a id="L299"></a>}
    <a id="L300"></a>if math.MaxFloat32 &lt; av { <span class="comment">// underflow is OK</span>
        <a id="L301"></a>state.err = i.ovfl
    <a id="L302"></a>} else {
        <a id="L303"></a>*(*float32)(p) = float32(v)
    <a id="L304"></a>}
<a id="L305"></a>}

<a id="L307"></a>func decFloat64(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L308"></a>if i.indir &gt; 0 {
        <a id="L309"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L310"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new(float64))
        <a id="L311"></a>}
        <a id="L312"></a>p = *(*unsafe.Pointer)(p);
    <a id="L313"></a>}
    <a id="L314"></a>*(*float64)(p) = floatFromBits(uint64(decodeUint(state)));
<a id="L315"></a>}

<a id="L317"></a><span class="comment">// uint8 arrays are encoded as an unsigned count followed by the raw bytes.</span>
<a id="L318"></a>func decUint8Array(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L319"></a>if i.indir &gt; 0 {
        <a id="L320"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L321"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new([]uint8))
        <a id="L322"></a>}
        <a id="L323"></a>p = *(*unsafe.Pointer)(p);
    <a id="L324"></a>}
    <a id="L325"></a>b := make([]uint8, decodeUint(state));
    <a id="L326"></a>state.b.Read(b);
    <a id="L327"></a>*(*[]uint8)(p) = b;
<a id="L328"></a>}

<a id="L330"></a><span class="comment">// Strings are encoded as an unsigned count followed by the raw bytes.</span>
<a id="L331"></a>func decString(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L332"></a>if i.indir &gt; 0 {
        <a id="L333"></a>if *(*unsafe.Pointer)(p) == nil {
            <a id="L334"></a>*(*unsafe.Pointer)(p) = unsafe.Pointer(new([]byte))
        <a id="L335"></a>}
        <a id="L336"></a>p = *(*unsafe.Pointer)(p);
    <a id="L337"></a>}
    <a id="L338"></a>b := make([]byte, decodeUint(state));
    <a id="L339"></a>state.b.Read(b);
    <a id="L340"></a>*(*string)(p) = string(b);
<a id="L341"></a>}

<a id="L343"></a>func ignoreUint8Array(i *decInstr, state *decodeState, p unsafe.Pointer) {
    <a id="L344"></a>b := make([]byte, decodeUint(state));
    <a id="L345"></a>state.b.Read(b);
<a id="L346"></a>}

<a id="L348"></a><span class="comment">// Execution engine</span>

<a id="L350"></a><span class="comment">// The encoder engine is an array of instructions indexed by field number of the incoming</span>
<a id="L351"></a><span class="comment">// data.  It is executed with random access according to field number.</span>
<a id="L352"></a>type decEngine struct {
    <a id="L353"></a>instr    []decInstr;
    <a id="L354"></a>numInstr int; <span class="comment">// the number of active instructions</span>
<a id="L355"></a>}

<a id="L357"></a>func decodeStruct(engine *decEngine, rtyp *reflect.StructType, b *bytes.Buffer, p uintptr, indir int) os.Error {
    <a id="L358"></a>if indir &gt; 0 {
        <a id="L359"></a>up := unsafe.Pointer(p);
        <a id="L360"></a>if indir &gt; 1 {
            <a id="L361"></a>up = decIndirect(up, indir)
        <a id="L362"></a>}
        <a id="L363"></a>if *(*unsafe.Pointer)(up) == nil {
            <a id="L364"></a><span class="comment">// Allocate object by making a slice of bytes and recording the</span>
            <a id="L365"></a><span class="comment">// address of the beginning of the array. TODO(rsc).</span>
            <a id="L366"></a>b := make([]byte, rtyp.Size());
            <a id="L367"></a>*(*unsafe.Pointer)(up) = unsafe.Pointer(&amp;b[0]);
        <a id="L368"></a>}
        <a id="L369"></a>p = *(*uintptr)(up);
    <a id="L370"></a>}
    <a id="L371"></a>state := newDecodeState(b);
    <a id="L372"></a>state.fieldnum = -1;
    <a id="L373"></a>basep := p;
    <a id="L374"></a>for state.err == nil {
        <a id="L375"></a>delta := int(decodeUint(state));
        <a id="L376"></a>if delta &lt; 0 {
            <a id="L377"></a>state.err = os.ErrorString(&#34;gob decode: corrupted data: negative delta&#34;);
            <a id="L378"></a>break;
        <a id="L379"></a>}
        <a id="L380"></a>if state.err != nil || delta == 0 { <span class="comment">// struct terminator is zero delta fieldnum</span>
            <a id="L381"></a>break
        <a id="L382"></a>}
        <a id="L383"></a>fieldnum := state.fieldnum + delta;
        <a id="L384"></a>if fieldnum &gt;= len(engine.instr) {
            <a id="L385"></a>state.err = errRange;
            <a id="L386"></a>break;
        <a id="L387"></a>}
        <a id="L388"></a>instr := &amp;engine.instr[fieldnum];
        <a id="L389"></a>p := unsafe.Pointer(basep + instr.offset);
        <a id="L390"></a>if instr.indir &gt; 1 {
            <a id="L391"></a>p = decIndirect(p, instr.indir)
        <a id="L392"></a>}
        <a id="L393"></a>instr.op(instr, state, p);
        <a id="L394"></a>state.fieldnum = fieldnum;
    <a id="L395"></a>}
    <a id="L396"></a>return state.err;
<a id="L397"></a>}

<a id="L399"></a>func ignoreStruct(engine *decEngine, b *bytes.Buffer) os.Error {
    <a id="L400"></a>state := newDecodeState(b);
    <a id="L401"></a>state.fieldnum = -1;
    <a id="L402"></a>for state.err == nil {
        <a id="L403"></a>delta := int(decodeUint(state));
        <a id="L404"></a>if delta &lt; 0 {
            <a id="L405"></a>state.err = os.ErrorString(&#34;gob ignore decode: corrupted data: negative delta&#34;);
            <a id="L406"></a>break;
        <a id="L407"></a>}
        <a id="L408"></a>if state.err != nil || delta == 0 { <span class="comment">// struct terminator is zero delta fieldnum</span>
            <a id="L409"></a>break
        <a id="L410"></a>}
        <a id="L411"></a>fieldnum := state.fieldnum + delta;
        <a id="L412"></a>if fieldnum &gt;= len(engine.instr) {
            <a id="L413"></a>state.err = errRange;
            <a id="L414"></a>break;
        <a id="L415"></a>}
        <a id="L416"></a>instr := &amp;engine.instr[fieldnum];
        <a id="L417"></a>instr.op(instr, state, unsafe.Pointer(nil));
        <a id="L418"></a>state.fieldnum = fieldnum;
    <a id="L419"></a>}
    <a id="L420"></a>return state.err;
<a id="L421"></a>}

<a id="L423"></a>func decodeArrayHelper(state *decodeState, p uintptr, elemOp decOp, elemWid uintptr, length, elemIndir int, ovfl os.ErrorString) os.Error {
    <a id="L424"></a>instr := &amp;decInstr{elemOp, 0, elemIndir, 0, ovfl};
    <a id="L425"></a>for i := 0; i &lt; length &amp;&amp; state.err == nil; i++ {
        <a id="L426"></a>up := unsafe.Pointer(p);
        <a id="L427"></a>if elemIndir &gt; 1 {
            <a id="L428"></a>up = decIndirect(up, elemIndir)
        <a id="L429"></a>}
        <a id="L430"></a>elemOp(instr, state, up);
        <a id="L431"></a>p += uintptr(elemWid);
    <a id="L432"></a>}
    <a id="L433"></a>return state.err;
<a id="L434"></a>}

<a id="L436"></a>func decodeArray(atyp *reflect.ArrayType, state *decodeState, p uintptr, elemOp decOp, elemWid uintptr, length, indir, elemIndir int, ovfl os.ErrorString) os.Error {
    <a id="L437"></a>if indir &gt; 0 {
        <a id="L438"></a>up := unsafe.Pointer(p);
        <a id="L439"></a>if *(*unsafe.Pointer)(up) == nil {
            <a id="L440"></a><span class="comment">// Allocate the array by making a slice of bytes of the correct size</span>
            <a id="L441"></a><span class="comment">// and taking the address of the beginning of the array. TODO(rsc).</span>
            <a id="L442"></a>b := make([]byte, atyp.Size());
            <a id="L443"></a>*(**byte)(up) = &amp;b[0];
        <a id="L444"></a>}
        <a id="L445"></a>p = *(*uintptr)(up);
    <a id="L446"></a>}
    <a id="L447"></a>if n := decodeUint(state); n != uint64(length) {
        <a id="L448"></a>return os.ErrorString(&#34;gob: length mismatch in decodeArray&#34;)
    <a id="L449"></a>}
    <a id="L450"></a>return decodeArrayHelper(state, p, elemOp, elemWid, length, elemIndir, ovfl);
<a id="L451"></a>}

<a id="L453"></a>func ignoreArrayHelper(state *decodeState, elemOp decOp, length int) os.Error {
    <a id="L454"></a>instr := &amp;decInstr{elemOp, 0, 0, 0, os.ErrorString(&#34;no error&#34;)};
    <a id="L455"></a>for i := 0; i &lt; length &amp;&amp; state.err == nil; i++ {
        <a id="L456"></a>elemOp(instr, state, nil)
    <a id="L457"></a>}
    <a id="L458"></a>return state.err;
<a id="L459"></a>}

<a id="L461"></a>func ignoreArray(state *decodeState, elemOp decOp, length int) os.Error {
    <a id="L462"></a>if n := decodeUint(state); n != uint64(length) {
        <a id="L463"></a>return os.ErrorString(&#34;gob: length mismatch in ignoreArray&#34;)
    <a id="L464"></a>}
    <a id="L465"></a>return ignoreArrayHelper(state, elemOp, length);
<a id="L466"></a>}

<a id="L468"></a>func decodeSlice(atyp *reflect.SliceType, state *decodeState, p uintptr, elemOp decOp, elemWid uintptr, indir, elemIndir int, ovfl os.ErrorString) os.Error {
    <a id="L469"></a>length := uintptr(decodeUint(state));
    <a id="L470"></a>if indir &gt; 0 {
        <a id="L471"></a>up := unsafe.Pointer(p);
        <a id="L472"></a>if *(*unsafe.Pointer)(up) == nil {
            <a id="L473"></a><span class="comment">// Allocate the slice header.</span>
            <a id="L474"></a>*(*unsafe.Pointer)(up) = unsafe.Pointer(new(reflect.SliceHeader))
        <a id="L475"></a>}
        <a id="L476"></a>p = *(*uintptr)(up);
    <a id="L477"></a>}
    <a id="L478"></a><span class="comment">// Allocate storage for the slice elements, that is, the underlying array.</span>
    <a id="L479"></a>data := make([]byte, length*atyp.Elem().Size());
    <a id="L480"></a><span class="comment">// Always write a header at p.</span>
    <a id="L481"></a>hdrp := (*reflect.SliceHeader)(unsafe.Pointer(p));
    <a id="L482"></a>hdrp.Data = uintptr(unsafe.Pointer(&amp;data[0]));
    <a id="L483"></a>hdrp.Len = int(length);
    <a id="L484"></a>hdrp.Cap = int(length);
    <a id="L485"></a>return decodeArrayHelper(state, hdrp.Data, elemOp, elemWid, int(length), elemIndir, ovfl);
<a id="L486"></a>}

<a id="L488"></a>func ignoreSlice(state *decodeState, elemOp decOp) os.Error {
    <a id="L489"></a>return ignoreArrayHelper(state, elemOp, int(decodeUint(state)))
<a id="L490"></a>}

<a id="L492"></a>var decOpMap = map[reflect.Type]decOp{
    <a id="L493"></a>valueKind(false): decBool,
    <a id="L494"></a>valueKind(int8(0)): decInt8,
    <a id="L495"></a>valueKind(int16(0)): decInt16,
    <a id="L496"></a>valueKind(int32(0)): decInt32,
    <a id="L497"></a>valueKind(int64(0)): decInt64,
    <a id="L498"></a>valueKind(uint8(0)): decUint8,
    <a id="L499"></a>valueKind(uint16(0)): decUint16,
    <a id="L500"></a>valueKind(uint32(0)): decUint32,
    <a id="L501"></a>valueKind(uint64(0)): decUint64,
    <a id="L502"></a>valueKind(float32(0)): decFloat32,
    <a id="L503"></a>valueKind(float64(0)): decFloat64,
    <a id="L504"></a>valueKind(&#34;x&#34;): decString,
<a id="L505"></a>}

<a id="L507"></a>var decIgnoreOpMap = map[typeId]decOp{
    <a id="L508"></a>tBool: ignoreUint,
    <a id="L509"></a>tInt: ignoreUint,
    <a id="L510"></a>tUint: ignoreUint,
    <a id="L511"></a>tFloat: ignoreUint,
    <a id="L512"></a>tBytes: ignoreUint8Array,
    <a id="L513"></a>tString: ignoreUint8Array,
<a id="L514"></a>}

<a id="L516"></a><span class="comment">// Return the decoding op for the base type under rt and</span>
<a id="L517"></a><span class="comment">// the indirection count to reach it.</span>
<a id="L518"></a>func decOpFor(wireId typeId, rt reflect.Type, name string) (decOp, int, os.Error) {
    <a id="L519"></a>typ, indir := indirect(rt);
    <a id="L520"></a>op, ok := decOpMap[reflect.Typeof(typ)];
    <a id="L521"></a>if !ok {
        <a id="L522"></a><span class="comment">// Special cases</span>
        <a id="L523"></a>switch t := typ.(type) {
        <a id="L524"></a>case *reflect.SliceType:
            <a id="L525"></a>name = &#34;element of &#34; + name;
            <a id="L526"></a>if _, ok := t.Elem().(*reflect.Uint8Type); ok {
                <a id="L527"></a>op = decUint8Array;
                <a id="L528"></a>break;
            <a id="L529"></a>}
            <a id="L530"></a>elemId := wireId.gobType().(*sliceType).Elem;
            <a id="L531"></a>elemOp, elemIndir, err := decOpFor(elemId, t.Elem(), name);
            <a id="L532"></a>if err != nil {
                <a id="L533"></a>return nil, 0, err
            <a id="L534"></a>}
            <a id="L535"></a>ovfl := overflow(name);
            <a id="L536"></a>op = func(i *decInstr, state *decodeState, p unsafe.Pointer) {
                <a id="L537"></a>state.err = decodeSlice(t, state, uintptr(p), elemOp, t.Elem().Size(), i.indir, elemIndir, ovfl)
            <a id="L538"></a>};

        <a id="L540"></a>case *reflect.ArrayType:
            <a id="L541"></a>name = &#34;element of &#34; + name;
            <a id="L542"></a>elemId := wireId.gobType().(*arrayType).Elem;
            <a id="L543"></a>elemOp, elemIndir, err := decOpFor(elemId, t.Elem(), name);
            <a id="L544"></a>if err != nil {
                <a id="L545"></a>return nil, 0, err
            <a id="L546"></a>}
            <a id="L547"></a>ovfl := overflow(name);
            <a id="L548"></a>op = func(i *decInstr, state *decodeState, p unsafe.Pointer) {
                <a id="L549"></a>state.err = decodeArray(t, state, uintptr(p), elemOp, t.Elem().Size(), t.Len(), i.indir, elemIndir, ovfl)
            <a id="L550"></a>};

        <a id="L552"></a>case *reflect.StructType:
            <a id="L553"></a><span class="comment">// Generate a closure that calls out to the engine for the nested type.</span>
            <a id="L554"></a>enginePtr, err := getDecEnginePtr(wireId, typ);
            <a id="L555"></a>if err != nil {
                <a id="L556"></a>return nil, 0, err
            <a id="L557"></a>}
            <a id="L558"></a>op = func(i *decInstr, state *decodeState, p unsafe.Pointer) {
                <a id="L559"></a><span class="comment">// indirect through enginePtr to delay evaluation for recursive structs</span>
                <a id="L560"></a>state.err = decodeStruct(*enginePtr, t, state.b, uintptr(p), i.indir)
            <a id="L561"></a>};
        <a id="L562"></a>}
    <a id="L563"></a>}
    <a id="L564"></a>if op == nil {
        <a id="L565"></a>return nil, 0, os.ErrorString(&#34;gob: decode can&#39;t handle type &#34; + rt.String())
    <a id="L566"></a>}
    <a id="L567"></a>return op, indir, nil;
<a id="L568"></a>}

<a id="L570"></a><span class="comment">// Return the decoding op for a field that has no destination.</span>
<a id="L571"></a>func decIgnoreOpFor(wireId typeId) (decOp, os.Error) {
    <a id="L572"></a>op, ok := decIgnoreOpMap[wireId];
    <a id="L573"></a>if !ok {
        <a id="L574"></a><span class="comment">// Special cases</span>
        <a id="L575"></a>switch t := wireId.gobType().(type) {
        <a id="L576"></a>case *sliceType:
            <a id="L577"></a>elemId := wireId.gobType().(*sliceType).Elem;
            <a id="L578"></a>elemOp, err := decIgnoreOpFor(elemId);
            <a id="L579"></a>if err != nil {
                <a id="L580"></a>return nil, err
            <a id="L581"></a>}
            <a id="L582"></a>op = func(i *decInstr, state *decodeState, p unsafe.Pointer) {
                <a id="L583"></a>state.err = ignoreSlice(state, elemOp)
            <a id="L584"></a>};

        <a id="L586"></a>case *arrayType:
            <a id="L587"></a>elemId := wireId.gobType().(*arrayType).Elem;
            <a id="L588"></a>elemOp, err := decIgnoreOpFor(elemId);
            <a id="L589"></a>if err != nil {
                <a id="L590"></a>return nil, err
            <a id="L591"></a>}
            <a id="L592"></a>op = func(i *decInstr, state *decodeState, p unsafe.Pointer) {
                <a id="L593"></a>state.err = ignoreArray(state, elemOp, t.Len)
            <a id="L594"></a>};

        <a id="L596"></a>case *structType:
            <a id="L597"></a><span class="comment">// Generate a closure that calls out to the engine for the nested type.</span>
            <a id="L598"></a>enginePtr, err := getIgnoreEnginePtr(wireId);
            <a id="L599"></a>if err != nil {
                <a id="L600"></a>return nil, err
            <a id="L601"></a>}
            <a id="L602"></a>op = func(i *decInstr, state *decodeState, p unsafe.Pointer) {
                <a id="L603"></a><span class="comment">// indirect through enginePtr to delay evaluation for recursive structs</span>
                <a id="L604"></a>state.err = ignoreStruct(*enginePtr, state.b)
            <a id="L605"></a>};
        <a id="L606"></a>}
    <a id="L607"></a>}
    <a id="L608"></a>if op == nil {
        <a id="L609"></a>return nil, os.ErrorString(&#34;ignore can&#39;t handle type &#34; + wireId.String())
    <a id="L610"></a>}
    <a id="L611"></a>return op, nil;
<a id="L612"></a>}

<a id="L614"></a><span class="comment">// Are these two gob Types compatible?</span>
<a id="L615"></a><span class="comment">// Answers the question for basic types, arrays, and slices.</span>
<a id="L616"></a><span class="comment">// Structs are considered ok; fields will be checked later.</span>
<a id="L617"></a>func compatibleType(fr reflect.Type, fw typeId) bool {
    <a id="L618"></a>for {
        <a id="L619"></a>if pt, ok := fr.(*reflect.PtrType); ok {
            <a id="L620"></a>fr = pt.Elem();
            <a id="L621"></a>continue;
        <a id="L622"></a>}
        <a id="L623"></a>break;
    <a id="L624"></a>}
    <a id="L625"></a>switch t := fr.(type) {
    <a id="L626"></a>default:
        <a id="L627"></a><span class="comment">// interface, map, chan, etc: cannot handle.</span>
        <a id="L628"></a>return false
    <a id="L629"></a>case *reflect.BoolType:
        <a id="L630"></a>return fw == tBool
    <a id="L631"></a>case *reflect.IntType:
        <a id="L632"></a>return fw == tInt
    <a id="L633"></a>case *reflect.Int8Type:
        <a id="L634"></a>return fw == tInt
    <a id="L635"></a>case *reflect.Int16Type:
        <a id="L636"></a>return fw == tInt
    <a id="L637"></a>case *reflect.Int32Type:
        <a id="L638"></a>return fw == tInt
    <a id="L639"></a>case *reflect.Int64Type:
        <a id="L640"></a>return fw == tInt
    <a id="L641"></a>case *reflect.UintType:
        <a id="L642"></a>return fw == tUint
    <a id="L643"></a>case *reflect.Uint8Type:
        <a id="L644"></a>return fw == tUint
    <a id="L645"></a>case *reflect.Uint16Type:
        <a id="L646"></a>return fw == tUint
    <a id="L647"></a>case *reflect.Uint32Type:
        <a id="L648"></a>return fw == tUint
    <a id="L649"></a>case *reflect.Uint64Type:
        <a id="L650"></a>return fw == tUint
    <a id="L651"></a>case *reflect.UintptrType:
        <a id="L652"></a>return fw == tUint
    <a id="L653"></a>case *reflect.FloatType:
        <a id="L654"></a>return fw == tFloat
    <a id="L655"></a>case *reflect.Float32Type:
        <a id="L656"></a>return fw == tFloat
    <a id="L657"></a>case *reflect.Float64Type:
        <a id="L658"></a>return fw == tFloat
    <a id="L659"></a>case *reflect.StringType:
        <a id="L660"></a>return fw == tString
    <a id="L661"></a>case *reflect.ArrayType:
        <a id="L662"></a>aw, ok := fw.gobType().(*arrayType);
        <a id="L663"></a>return ok &amp;&amp; t.Len() == aw.Len &amp;&amp; compatibleType(t.Elem(), aw.Elem);
    <a id="L664"></a>case *reflect.SliceType:
        <a id="L665"></a><span class="comment">// Is it an array of bytes?</span>
        <a id="L666"></a>et := t.Elem();
        <a id="L667"></a>if _, ok := et.(*reflect.Uint8Type); ok {
            <a id="L668"></a>return fw == tBytes
        <a id="L669"></a>}
        <a id="L670"></a>sw, ok := fw.gobType().(*sliceType);
        <a id="L671"></a>elem, _ := indirect(t.Elem());
        <a id="L672"></a>return ok &amp;&amp; compatibleType(elem, sw.Elem);
    <a id="L673"></a>case *reflect.StructType:
        <a id="L674"></a>return true
    <a id="L675"></a>}
    <a id="L676"></a>return true;
<a id="L677"></a>}

<a id="L679"></a>func compileDec(wireId typeId, rt reflect.Type) (engine *decEngine, err os.Error) {
    <a id="L680"></a>srt, ok1 := rt.(*reflect.StructType);
    <a id="L681"></a>wireStruct, ok2 := wireId.gobType().(*structType);
    <a id="L682"></a>if !ok1 || !ok2 {
        <a id="L683"></a>return nil, errNotStruct
    <a id="L684"></a>}
    <a id="L685"></a>engine = new(decEngine);
    <a id="L686"></a>engine.instr = make([]decInstr, len(wireStruct.field));
    <a id="L687"></a><span class="comment">// Loop over the fields of the wire type.</span>
    <a id="L688"></a>for fieldnum := 0; fieldnum &lt; len(wireStruct.field); fieldnum++ {
        <a id="L689"></a>wireField := wireStruct.field[fieldnum];
        <a id="L690"></a><span class="comment">// Find the field of the local type with the same name.</span>
        <a id="L691"></a>localField, present := srt.FieldByName(wireField.name);
        <a id="L692"></a>ovfl := overflow(wireField.name);
        <a id="L693"></a><span class="comment">// TODO(r): anonymous names</span>
        <a id="L694"></a>if !present {
            <a id="L695"></a>op, err := decIgnoreOpFor(wireField.id);
            <a id="L696"></a>if err != nil {
                <a id="L697"></a>return nil, err
            <a id="L698"></a>}
            <a id="L699"></a>engine.instr[fieldnum] = decInstr{op, fieldnum, 0, 0, ovfl};
            <a id="L700"></a>continue;
        <a id="L701"></a>}
        <a id="L702"></a>if !compatibleType(localField.Type, wireField.id) {
            <a id="L703"></a>details := &#34; (&#34; + wireField.id.String() + &#34; incompatible with &#34; + localField.Type.String() + &#34;) in type &#34; + wireId.Name();
            <a id="L704"></a>return nil, os.ErrorString(&#34;gob: wrong type for field &#34; + wireField.name + details);
        <a id="L705"></a>}
        <a id="L706"></a>op, indir, err := decOpFor(wireField.id, localField.Type, localField.Name);
        <a id="L707"></a>if err != nil {
            <a id="L708"></a>return nil, err
        <a id="L709"></a>}
        <a id="L710"></a>engine.instr[fieldnum] = decInstr{op, fieldnum, indir, uintptr(localField.Offset), ovfl};
        <a id="L711"></a>engine.numInstr++;
    <a id="L712"></a>}
    <a id="L713"></a>return;
<a id="L714"></a>}

<a id="L716"></a>var decoderCache = make(map[reflect.Type]map[typeId]**decEngine)
<a id="L717"></a>var ignorerCache = make(map[typeId]**decEngine)

<a id="L719"></a><span class="comment">// typeLock must be held.</span>
<a id="L720"></a>func getDecEnginePtr(wireId typeId, rt reflect.Type) (enginePtr **decEngine, err os.Error) {
    <a id="L721"></a>decoderMap, ok := decoderCache[rt];
    <a id="L722"></a>if !ok {
        <a id="L723"></a>decoderMap = make(map[typeId]**decEngine);
        <a id="L724"></a>decoderCache[rt] = decoderMap;
    <a id="L725"></a>}
    <a id="L726"></a>if enginePtr, ok = decoderMap[wireId]; !ok {
        <a id="L727"></a><span class="comment">// To handle recursive types, mark this engine as underway before compiling.</span>
        <a id="L728"></a>enginePtr = new(*decEngine);
        <a id="L729"></a>decoderMap[wireId] = enginePtr;
        <a id="L730"></a>*enginePtr, err = compileDec(wireId, rt);
        <a id="L731"></a>if err != nil {
            <a id="L732"></a>decoderMap[wireId] = nil, false
        <a id="L733"></a>}
    <a id="L734"></a>}
    <a id="L735"></a>return;
<a id="L736"></a>}

<a id="L738"></a><span class="comment">// When ignoring data, in effect we compile it into this type</span>
<a id="L739"></a>type emptyStruct struct{}

<a id="L741"></a>var emptyStructType = reflect.Typeof(emptyStruct{})

<a id="L743"></a><span class="comment">// typeLock must be held.</span>
<a id="L744"></a>func getIgnoreEnginePtr(wireId typeId) (enginePtr **decEngine, err os.Error) {
    <a id="L745"></a>var ok bool;
    <a id="L746"></a>if enginePtr, ok = ignorerCache[wireId]; !ok {
        <a id="L747"></a><span class="comment">// To handle recursive types, mark this engine as underway before compiling.</span>
        <a id="L748"></a>enginePtr = new(*decEngine);
        <a id="L749"></a>ignorerCache[wireId] = enginePtr;
        <a id="L750"></a>*enginePtr, err = compileDec(wireId, emptyStructType);
        <a id="L751"></a>if err != nil {
            <a id="L752"></a>ignorerCache[wireId] = nil, false
        <a id="L753"></a>}
    <a id="L754"></a>}
    <a id="L755"></a>return;
<a id="L756"></a>}

<a id="L758"></a>func decode(b *bytes.Buffer, wireId typeId, e interface{}) os.Error {
    <a id="L759"></a><span class="comment">// Dereference down to the underlying struct type.</span>
    <a id="L760"></a>rt, indir := indirect(reflect.Typeof(e));
    <a id="L761"></a>st, ok := rt.(*reflect.StructType);
    <a id="L762"></a>if !ok {
        <a id="L763"></a>return os.ErrorString(&#34;gob: decode can&#39;t handle &#34; + rt.String())
    <a id="L764"></a>}
    <a id="L765"></a>typeLock.Lock();
    <a id="L766"></a>if _, ok := idToType[wireId]; !ok {
        <a id="L767"></a>typeLock.Unlock();
        <a id="L768"></a>return errBadType;
    <a id="L769"></a>}
    <a id="L770"></a>enginePtr, err := getDecEnginePtr(wireId, rt);
    <a id="L771"></a>typeLock.Unlock();
    <a id="L772"></a>if err != nil {
        <a id="L773"></a>return err
    <a id="L774"></a>}
    <a id="L775"></a>engine := *enginePtr;
    <a id="L776"></a>if engine.numInstr == 0 &amp;&amp; st.NumField() &gt; 0 &amp;&amp; len(wireId.gobType().(*structType).field) &gt; 0 {
        <a id="L777"></a>name := rt.Name();
        <a id="L778"></a>return os.ErrorString(&#34;gob: type mismatch: no fields matched compiling decoder for &#34; + name);
    <a id="L779"></a>}
    <a id="L780"></a>return decodeStruct(engine, st, b, uintptr(reflect.NewValue(e).Addr()), indir);
<a id="L781"></a>}

<a id="L783"></a>func init() {
    <a id="L784"></a><span class="comment">// We assume that the size of float is sufficient to tell us whether it is</span>
    <a id="L785"></a><span class="comment">// equivalent to float32 or to float64.   This is very unlikely to be wrong.</span>
    <a id="L786"></a>var op decOp;
    <a id="L787"></a>switch unsafe.Sizeof(float(0)) {
    <a id="L788"></a>case unsafe.Sizeof(float32(0)):
        <a id="L789"></a>op = decFloat32
    <a id="L790"></a>case unsafe.Sizeof(float64(0)):
        <a id="L791"></a>op = decFloat64
    <a id="L792"></a>default:
        <a id="L793"></a>panic(&#34;gob: unknown size of float&#34;, unsafe.Sizeof(float(0)))
    <a id="L794"></a>}
    <a id="L795"></a>decOpMap[valueKind(float(0))] = op;

    <a id="L797"></a><span class="comment">// A similar assumption about int and uint.  Also assume int and uint have the same size.</span>
    <a id="L798"></a>var uop decOp;
    <a id="L799"></a>switch unsafe.Sizeof(int(0)) {
    <a id="L800"></a>case unsafe.Sizeof(int32(0)):
        <a id="L801"></a>op = decInt32;
        <a id="L802"></a>uop = decUint32;
    <a id="L803"></a>case unsafe.Sizeof(int64(0)):
        <a id="L804"></a>op = decInt64;
        <a id="L805"></a>uop = decUint64;
    <a id="L806"></a>default:
        <a id="L807"></a>panic(&#34;gob: unknown size of int/uint&#34;, unsafe.Sizeof(int(0)))
    <a id="L808"></a>}
    <a id="L809"></a>decOpMap[valueKind(int(0))] = op;
    <a id="L810"></a>decOpMap[valueKind(uint(0))] = uop;

    <a id="L812"></a><span class="comment">// Finally uintptr</span>
    <a id="L813"></a>switch unsafe.Sizeof(uintptr(0)) {
    <a id="L814"></a>case unsafe.Sizeof(uint32(0)):
        <a id="L815"></a>uop = decUint32
    <a id="L816"></a>case unsafe.Sizeof(uint64(0)):
        <a id="L817"></a>uop = decUint64
    <a id="L818"></a>default:
        <a id="L819"></a>panic(&#34;gob: unknown size of uintptr&#34;, unsafe.Sizeof(uintptr(0)))
    <a id="L820"></a>}
    <a id="L821"></a>decOpMap[valueKind(uintptr(0))] = uop;
<a id="L822"></a>}
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
