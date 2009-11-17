<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/gob/codec_test.go</title>

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
	<li>Thu Nov 12 15:57:42 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/gob/codec_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package gob

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;math&#34;;
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;reflect&#34;;
    <a id="L12"></a>&#34;strings&#34;;
    <a id="L13"></a>&#34;testing&#34;;
    <a id="L14"></a>&#34;unsafe&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// Guarantee encoding format by comparing some encodings to hand-written values</span>
<a id="L18"></a>type EncodeT struct {
    <a id="L19"></a>x   uint64;
    <a id="L20"></a>b   []byte;
<a id="L21"></a>}

<a id="L23"></a>var encodeT = []EncodeT{
    <a id="L24"></a>EncodeT{0x00, []byte{0x00}},
    <a id="L25"></a>EncodeT{0x0F, []byte{0x0F}},
    <a id="L26"></a>EncodeT{0xFF, []byte{0xFF, 0xFF}},
    <a id="L27"></a>EncodeT{0xFFFF, []byte{0xFE, 0xFF, 0xFF}},
    <a id="L28"></a>EncodeT{0xFFFFFF, []byte{0xFD, 0xFF, 0xFF, 0xFF}},
    <a id="L29"></a>EncodeT{0xFFFFFFFF, []byte{0xFC, 0xFF, 0xFF, 0xFF, 0xFF}},
    <a id="L30"></a>EncodeT{0xFFFFFFFFFF, []byte{0xFB, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
    <a id="L31"></a>EncodeT{0xFFFFFFFFFFFF, []byte{0xFA, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
    <a id="L32"></a>EncodeT{0xFFFFFFFFFFFFFF, []byte{0xF9, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
    <a id="L33"></a>EncodeT{0xFFFFFFFFFFFFFFFF, []byte{0xF8, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
    <a id="L34"></a>EncodeT{0x1111, []byte{0xFE, 0x11, 0x11}},
    <a id="L35"></a>EncodeT{0x1111111111111111, []byte{0xF8, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}},
    <a id="L36"></a>EncodeT{0x8888888888888888, []byte{0xF8, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88, 0x88}},
    <a id="L37"></a>EncodeT{1 &lt;&lt; 63, []byte{0xF8, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}},
<a id="L38"></a>}


<a id="L41"></a><span class="comment">// Test basic encode/decode routines for unsigned integers</span>
<a id="L42"></a>func TestUintCodec(t *testing.T) {
    <a id="L43"></a>b := new(bytes.Buffer);
    <a id="L44"></a>encState := new(encoderState);
    <a id="L45"></a>encState.b = b;
    <a id="L46"></a>for _, tt := range encodeT {
        <a id="L47"></a>b.Reset();
        <a id="L48"></a>encodeUint(encState, tt.x);
        <a id="L49"></a>if encState.err != nil {
            <a id="L50"></a>t.Error(&#34;encodeUint:&#34;, tt.x, encState.err)
        <a id="L51"></a>}
        <a id="L52"></a>if !bytes.Equal(tt.b, b.Bytes()) {
            <a id="L53"></a>t.Errorf(&#34;encodeUint: %#x encode: expected % x got % x&#34;, tt.x, tt.b, b.Bytes())
        <a id="L54"></a>}
    <a id="L55"></a>}
    <a id="L56"></a>decState := newDecodeState(b);
    <a id="L57"></a>for u := uint64(0); ; u = (u + 1) * 7 {
        <a id="L58"></a>b.Reset();
        <a id="L59"></a>encodeUint(encState, u);
        <a id="L60"></a>if encState.err != nil {
            <a id="L61"></a>t.Error(&#34;encodeUint:&#34;, u, encState.err)
        <a id="L62"></a>}
        <a id="L63"></a>v := decodeUint(decState);
        <a id="L64"></a>if decState.err != nil {
            <a id="L65"></a>t.Error(&#34;DecodeUint:&#34;, u, decState.err)
        <a id="L66"></a>}
        <a id="L67"></a>if u != v {
            <a id="L68"></a>t.Errorf(&#34;Encode/Decode: sent %#x received %#x\n&#34;, u, v)
        <a id="L69"></a>}
        <a id="L70"></a>if u&amp;(1&lt;&lt;63) != 0 {
            <a id="L71"></a>break
        <a id="L72"></a>}
    <a id="L73"></a>}
<a id="L74"></a>}

<a id="L76"></a>func verifyInt(i int64, t *testing.T) {
    <a id="L77"></a>var b = new(bytes.Buffer);
    <a id="L78"></a>encState := new(encoderState);
    <a id="L79"></a>encState.b = b;
    <a id="L80"></a>encodeInt(encState, i);
    <a id="L81"></a>if encState.err != nil {
        <a id="L82"></a>t.Error(&#34;encodeInt:&#34;, i, encState.err)
    <a id="L83"></a>}
    <a id="L84"></a>decState := newDecodeState(b);
    <a id="L85"></a>decState.buf = make([]byte, 8);
    <a id="L86"></a>j := decodeInt(decState);
    <a id="L87"></a>if decState.err != nil {
        <a id="L88"></a>t.Error(&#34;DecodeInt:&#34;, i, decState.err)
    <a id="L89"></a>}
    <a id="L90"></a>if i != j {
        <a id="L91"></a>t.Errorf(&#34;Encode/Decode: sent %#x received %#x\n&#34;, uint64(i), uint64(j))
    <a id="L92"></a>}
<a id="L93"></a>}

<a id="L95"></a><span class="comment">// Test basic encode/decode routines for signed integers</span>
<a id="L96"></a>func TestIntCodec(t *testing.T) {
    <a id="L97"></a>for u := uint64(0); ; u = (u + 1) * 7 {
        <a id="L98"></a><span class="comment">// Do positive and negative values</span>
        <a id="L99"></a>i := int64(u);
        <a id="L100"></a>verifyInt(i, t);
        <a id="L101"></a>verifyInt(-i, t);
        <a id="L102"></a>verifyInt(^i, t);
        <a id="L103"></a>if u&amp;(1&lt;&lt;63) != 0 {
            <a id="L104"></a>break
        <a id="L105"></a>}
    <a id="L106"></a>}
    <a id="L107"></a>verifyInt(-1&lt;&lt;63, t); <span class="comment">// a tricky case</span>
<a id="L108"></a>}

<a id="L110"></a><span class="comment">// The result of encoding a true boolean with field number 7</span>
<a id="L111"></a>var boolResult = []byte{0x07, 0x01}
<a id="L112"></a><span class="comment">// The result of encoding a number 17 with field number 7</span>
<a id="L113"></a>var signedResult = []byte{0x07, 2 * 17}
<a id="L114"></a>var unsignedResult = []byte{0x07, 17}
<a id="L115"></a>var floatResult = []byte{0x07, 0xFE, 0x31, 0x40}
<a id="L116"></a><span class="comment">// The result of encoding &#34;hello&#34; with field number 6</span>
<a id="L117"></a>var bytesResult = []byte{0x07, 0x05, &#39;h&#39;, &#39;e&#39;, &#39;l&#39;, &#39;l&#39;, &#39;o&#39;}

<a id="L119"></a>func newencoderState(b *bytes.Buffer) *encoderState {
    <a id="L120"></a>b.Reset();
    <a id="L121"></a>state := new(encoderState);
    <a id="L122"></a>state.b = b;
    <a id="L123"></a>state.fieldnum = -1;
    <a id="L124"></a>return state;
<a id="L125"></a>}

<a id="L127"></a><span class="comment">// Test instruction execution for encoding.</span>
<a id="L128"></a><span class="comment">// Do not run the machine yet; instead do individual instructions crafted by hand.</span>
<a id="L129"></a>func TestScalarEncInstructions(t *testing.T) {
    <a id="L130"></a>var b = new(bytes.Buffer);

    <a id="L132"></a><span class="comment">// bool</span>
    <a id="L133"></a>{
        <a id="L134"></a>data := struct{ a bool }{true};
        <a id="L135"></a>instr := &amp;encInstr{encBool, 6, 0, 0};
        <a id="L136"></a>state := newencoderState(b);
        <a id="L137"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L138"></a>if !bytes.Equal(boolResult, b.Bytes()) {
            <a id="L139"></a>t.Errorf(&#34;bool enc instructions: expected % x got % x&#34;, boolResult, b.Bytes())
        <a id="L140"></a>}
    <a id="L141"></a>}

    <a id="L143"></a><span class="comment">// int</span>
    <a id="L144"></a>{
        <a id="L145"></a>b.Reset();
        <a id="L146"></a>data := struct{ a int }{17};
        <a id="L147"></a>instr := &amp;encInstr{encInt, 6, 0, 0};
        <a id="L148"></a>state := newencoderState(b);
        <a id="L149"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L150"></a>if !bytes.Equal(signedResult, b.Bytes()) {
            <a id="L151"></a>t.Errorf(&#34;int enc instructions: expected % x got % x&#34;, signedResult, b.Bytes())
        <a id="L152"></a>}
    <a id="L153"></a>}

    <a id="L155"></a><span class="comment">// uint</span>
    <a id="L156"></a>{
        <a id="L157"></a>b.Reset();
        <a id="L158"></a>data := struct{ a uint }{17};
        <a id="L159"></a>instr := &amp;encInstr{encUint, 6, 0, 0};
        <a id="L160"></a>state := newencoderState(b);
        <a id="L161"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L162"></a>if !bytes.Equal(unsignedResult, b.Bytes()) {
            <a id="L163"></a>t.Errorf(&#34;uint enc instructions: expected % x got % x&#34;, unsignedResult, b.Bytes())
        <a id="L164"></a>}
    <a id="L165"></a>}

    <a id="L167"></a><span class="comment">// int8</span>
    <a id="L168"></a>{
        <a id="L169"></a>b.Reset();
        <a id="L170"></a>data := struct{ a int8 }{17};
        <a id="L171"></a>instr := &amp;encInstr{encInt8, 6, 0, 0};
        <a id="L172"></a>state := newencoderState(b);
        <a id="L173"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L174"></a>if !bytes.Equal(signedResult, b.Bytes()) {
            <a id="L175"></a>t.Errorf(&#34;int8 enc instructions: expected % x got % x&#34;, signedResult, b.Bytes())
        <a id="L176"></a>}
    <a id="L177"></a>}

    <a id="L179"></a><span class="comment">// uint8</span>
    <a id="L180"></a>{
        <a id="L181"></a>b.Reset();
        <a id="L182"></a>data := struct{ a uint8 }{17};
        <a id="L183"></a>instr := &amp;encInstr{encUint8, 6, 0, 0};
        <a id="L184"></a>state := newencoderState(b);
        <a id="L185"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L186"></a>if !bytes.Equal(unsignedResult, b.Bytes()) {
            <a id="L187"></a>t.Errorf(&#34;uint8 enc instructions: expected % x got % x&#34;, unsignedResult, b.Bytes())
        <a id="L188"></a>}
    <a id="L189"></a>}

    <a id="L191"></a><span class="comment">// int16</span>
    <a id="L192"></a>{
        <a id="L193"></a>b.Reset();
        <a id="L194"></a>data := struct{ a int16 }{17};
        <a id="L195"></a>instr := &amp;encInstr{encInt16, 6, 0, 0};
        <a id="L196"></a>state := newencoderState(b);
        <a id="L197"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L198"></a>if !bytes.Equal(signedResult, b.Bytes()) {
            <a id="L199"></a>t.Errorf(&#34;int16 enc instructions: expected % x got % x&#34;, signedResult, b.Bytes())
        <a id="L200"></a>}
    <a id="L201"></a>}

    <a id="L203"></a><span class="comment">// uint16</span>
    <a id="L204"></a>{
        <a id="L205"></a>b.Reset();
        <a id="L206"></a>data := struct{ a uint16 }{17};
        <a id="L207"></a>instr := &amp;encInstr{encUint16, 6, 0, 0};
        <a id="L208"></a>state := newencoderState(b);
        <a id="L209"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L210"></a>if !bytes.Equal(unsignedResult, b.Bytes()) {
            <a id="L211"></a>t.Errorf(&#34;uint16 enc instructions: expected % x got % x&#34;, unsignedResult, b.Bytes())
        <a id="L212"></a>}
    <a id="L213"></a>}

    <a id="L215"></a><span class="comment">// int32</span>
    <a id="L216"></a>{
        <a id="L217"></a>b.Reset();
        <a id="L218"></a>data := struct{ a int32 }{17};
        <a id="L219"></a>instr := &amp;encInstr{encInt32, 6, 0, 0};
        <a id="L220"></a>state := newencoderState(b);
        <a id="L221"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L222"></a>if !bytes.Equal(signedResult, b.Bytes()) {
            <a id="L223"></a>t.Errorf(&#34;int32 enc instructions: expected % x got % x&#34;, signedResult, b.Bytes())
        <a id="L224"></a>}
    <a id="L225"></a>}

    <a id="L227"></a><span class="comment">// uint32</span>
    <a id="L228"></a>{
        <a id="L229"></a>b.Reset();
        <a id="L230"></a>data := struct{ a uint32 }{17};
        <a id="L231"></a>instr := &amp;encInstr{encUint32, 6, 0, 0};
        <a id="L232"></a>state := newencoderState(b);
        <a id="L233"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L234"></a>if !bytes.Equal(unsignedResult, b.Bytes()) {
            <a id="L235"></a>t.Errorf(&#34;uint32 enc instructions: expected % x got % x&#34;, unsignedResult, b.Bytes())
        <a id="L236"></a>}
    <a id="L237"></a>}

    <a id="L239"></a><span class="comment">// int64</span>
    <a id="L240"></a>{
        <a id="L241"></a>b.Reset();
        <a id="L242"></a>data := struct{ a int64 }{17};
        <a id="L243"></a>instr := &amp;encInstr{encInt64, 6, 0, 0};
        <a id="L244"></a>state := newencoderState(b);
        <a id="L245"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L246"></a>if !bytes.Equal(signedResult, b.Bytes()) {
            <a id="L247"></a>t.Errorf(&#34;int64 enc instructions: expected % x got % x&#34;, signedResult, b.Bytes())
        <a id="L248"></a>}
    <a id="L249"></a>}

    <a id="L251"></a><span class="comment">// uint64</span>
    <a id="L252"></a>{
        <a id="L253"></a>b.Reset();
        <a id="L254"></a>data := struct{ a uint64 }{17};
        <a id="L255"></a>instr := &amp;encInstr{encUint64, 6, 0, 0};
        <a id="L256"></a>state := newencoderState(b);
        <a id="L257"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L258"></a>if !bytes.Equal(unsignedResult, b.Bytes()) {
            <a id="L259"></a>t.Errorf(&#34;uint64 enc instructions: expected % x got % x&#34;, unsignedResult, b.Bytes())
        <a id="L260"></a>}
    <a id="L261"></a>}

    <a id="L263"></a><span class="comment">// float</span>
    <a id="L264"></a>{
        <a id="L265"></a>b.Reset();
        <a id="L266"></a>data := struct{ a float }{17};
        <a id="L267"></a>instr := &amp;encInstr{encFloat, 6, 0, 0};
        <a id="L268"></a>state := newencoderState(b);
        <a id="L269"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L270"></a>if !bytes.Equal(floatResult, b.Bytes()) {
            <a id="L271"></a>t.Errorf(&#34;float enc instructions: expected % x got % x&#34;, floatResult, b.Bytes())
        <a id="L272"></a>}
    <a id="L273"></a>}

    <a id="L275"></a><span class="comment">// float32</span>
    <a id="L276"></a>{
        <a id="L277"></a>b.Reset();
        <a id="L278"></a>data := struct{ a float32 }{17};
        <a id="L279"></a>instr := &amp;encInstr{encFloat32, 6, 0, 0};
        <a id="L280"></a>state := newencoderState(b);
        <a id="L281"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L282"></a>if !bytes.Equal(floatResult, b.Bytes()) {
            <a id="L283"></a>t.Errorf(&#34;float32 enc instructions: expected % x got % x&#34;, floatResult, b.Bytes())
        <a id="L284"></a>}
    <a id="L285"></a>}

    <a id="L287"></a><span class="comment">// float64</span>
    <a id="L288"></a>{
        <a id="L289"></a>b.Reset();
        <a id="L290"></a>data := struct{ a float64 }{17};
        <a id="L291"></a>instr := &amp;encInstr{encFloat64, 6, 0, 0};
        <a id="L292"></a>state := newencoderState(b);
        <a id="L293"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L294"></a>if !bytes.Equal(floatResult, b.Bytes()) {
            <a id="L295"></a>t.Errorf(&#34;float64 enc instructions: expected % x got % x&#34;, floatResult, b.Bytes())
        <a id="L296"></a>}
    <a id="L297"></a>}

    <a id="L299"></a><span class="comment">// bytes == []uint8</span>
    <a id="L300"></a>{
        <a id="L301"></a>b.Reset();
        <a id="L302"></a>data := struct{ a []byte }{strings.Bytes(&#34;hello&#34;)};
        <a id="L303"></a>instr := &amp;encInstr{encUint8Array, 6, 0, 0};
        <a id="L304"></a>state := newencoderState(b);
        <a id="L305"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L306"></a>if !bytes.Equal(bytesResult, b.Bytes()) {
            <a id="L307"></a>t.Errorf(&#34;bytes enc instructions: expected % x got % x&#34;, bytesResult, b.Bytes())
        <a id="L308"></a>}
    <a id="L309"></a>}

    <a id="L311"></a><span class="comment">// string</span>
    <a id="L312"></a>{
        <a id="L313"></a>b.Reset();
        <a id="L314"></a>data := struct{ a string }{&#34;hello&#34;};
        <a id="L315"></a>instr := &amp;encInstr{encString, 6, 0, 0};
        <a id="L316"></a>state := newencoderState(b);
        <a id="L317"></a>instr.op(instr, state, unsafe.Pointer(&amp;data));
        <a id="L318"></a>if !bytes.Equal(bytesResult, b.Bytes()) {
            <a id="L319"></a>t.Errorf(&#34;string enc instructions: expected % x got % x&#34;, bytesResult, b.Bytes())
        <a id="L320"></a>}
    <a id="L321"></a>}
<a id="L322"></a>}

<a id="L324"></a>func execDec(typ string, instr *decInstr, state *decodeState, t *testing.T, p unsafe.Pointer) {
    <a id="L325"></a>v := int(decodeUint(state));
    <a id="L326"></a>if state.err != nil {
        <a id="L327"></a>t.Fatalf(&#34;decoding %s field: %v&#34;, typ, state.err)
    <a id="L328"></a>}
    <a id="L329"></a>if v+state.fieldnum != 6 {
        <a id="L330"></a>t.Fatalf(&#34;decoding field number %d, got %d&#34;, 6, v+state.fieldnum)
    <a id="L331"></a>}
    <a id="L332"></a>instr.op(instr, state, decIndirect(p, instr.indir));
    <a id="L333"></a>state.fieldnum = 6;
<a id="L334"></a>}

<a id="L336"></a>func newDecodeStateFromData(data []byte) *decodeState {
    <a id="L337"></a>state := newDecodeState(bytes.NewBuffer(data));
    <a id="L338"></a>state.fieldnum = -1;
    <a id="L339"></a>return state;
<a id="L340"></a>}

<a id="L342"></a><span class="comment">// Test instruction execution for decoding.</span>
<a id="L343"></a><span class="comment">// Do not run the machine yet; instead do individual instructions crafted by hand.</span>
<a id="L344"></a>func TestScalarDecInstructions(t *testing.T) {
    <a id="L345"></a>ovfl := os.ErrorString(&#34;overflow&#34;);

    <a id="L347"></a><span class="comment">// bool</span>
    <a id="L348"></a>{
        <a id="L349"></a>var data struct {
            <a id="L350"></a>a bool;
        <a id="L351"></a>}
        <a id="L352"></a>instr := &amp;decInstr{decBool, 6, 0, 0, ovfl};
        <a id="L353"></a>state := newDecodeStateFromData(boolResult);
        <a id="L354"></a>execDec(&#34;bool&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L355"></a>if data.a != true {
            <a id="L356"></a>t.Errorf(&#34;bool a = %v not true&#34;, data.a)
        <a id="L357"></a>}
    <a id="L358"></a>}
    <a id="L359"></a><span class="comment">// int</span>
    <a id="L360"></a>{
        <a id="L361"></a>var data struct {
            <a id="L362"></a>a int;
        <a id="L363"></a>}
        <a id="L364"></a>instr := &amp;decInstr{decOpMap[valueKind(data.a)], 6, 0, 0, ovfl};
        <a id="L365"></a>state := newDecodeStateFromData(signedResult);
        <a id="L366"></a>execDec(&#34;int&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L367"></a>if data.a != 17 {
            <a id="L368"></a>t.Errorf(&#34;int a = %v not 17&#34;, data.a)
        <a id="L369"></a>}
    <a id="L370"></a>}

    <a id="L372"></a><span class="comment">// uint</span>
    <a id="L373"></a>{
        <a id="L374"></a>var data struct {
            <a id="L375"></a>a uint;
        <a id="L376"></a>}
        <a id="L377"></a>instr := &amp;decInstr{decOpMap[valueKind(data.a)], 6, 0, 0, ovfl};
        <a id="L378"></a>state := newDecodeStateFromData(unsignedResult);
        <a id="L379"></a>execDec(&#34;uint&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L380"></a>if data.a != 17 {
            <a id="L381"></a>t.Errorf(&#34;uint a = %v not 17&#34;, data.a)
        <a id="L382"></a>}
    <a id="L383"></a>}

    <a id="L385"></a><span class="comment">// int8</span>
    <a id="L386"></a>{
        <a id="L387"></a>var data struct {
            <a id="L388"></a>a int8;
        <a id="L389"></a>}
        <a id="L390"></a>instr := &amp;decInstr{decInt8, 6, 0, 0, ovfl};
        <a id="L391"></a>state := newDecodeStateFromData(signedResult);
        <a id="L392"></a>execDec(&#34;int8&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L393"></a>if data.a != 17 {
            <a id="L394"></a>t.Errorf(&#34;int8 a = %v not 17&#34;, data.a)
        <a id="L395"></a>}
    <a id="L396"></a>}

    <a id="L398"></a><span class="comment">// uint8</span>
    <a id="L399"></a>{
        <a id="L400"></a>var data struct {
            <a id="L401"></a>a uint8;
        <a id="L402"></a>}
        <a id="L403"></a>instr := &amp;decInstr{decUint8, 6, 0, 0, ovfl};
        <a id="L404"></a>state := newDecodeStateFromData(unsignedResult);
        <a id="L405"></a>execDec(&#34;uint8&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L406"></a>if data.a != 17 {
            <a id="L407"></a>t.Errorf(&#34;uint8 a = %v not 17&#34;, data.a)
        <a id="L408"></a>}
    <a id="L409"></a>}

    <a id="L411"></a><span class="comment">// int16</span>
    <a id="L412"></a>{
        <a id="L413"></a>var data struct {
            <a id="L414"></a>a int16;
        <a id="L415"></a>}
        <a id="L416"></a>instr := &amp;decInstr{decInt16, 6, 0, 0, ovfl};
        <a id="L417"></a>state := newDecodeStateFromData(signedResult);
        <a id="L418"></a>execDec(&#34;int16&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L419"></a>if data.a != 17 {
            <a id="L420"></a>t.Errorf(&#34;int16 a = %v not 17&#34;, data.a)
        <a id="L421"></a>}
    <a id="L422"></a>}

    <a id="L424"></a><span class="comment">// uint16</span>
    <a id="L425"></a>{
        <a id="L426"></a>var data struct {
            <a id="L427"></a>a uint16;
        <a id="L428"></a>}
        <a id="L429"></a>instr := &amp;decInstr{decUint16, 6, 0, 0, ovfl};
        <a id="L430"></a>state := newDecodeStateFromData(unsignedResult);
        <a id="L431"></a>execDec(&#34;uint16&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L432"></a>if data.a != 17 {
            <a id="L433"></a>t.Errorf(&#34;uint16 a = %v not 17&#34;, data.a)
        <a id="L434"></a>}
    <a id="L435"></a>}

    <a id="L437"></a><span class="comment">// int32</span>
    <a id="L438"></a>{
        <a id="L439"></a>var data struct {
            <a id="L440"></a>a int32;
        <a id="L441"></a>}
        <a id="L442"></a>instr := &amp;decInstr{decInt32, 6, 0, 0, ovfl};
        <a id="L443"></a>state := newDecodeStateFromData(signedResult);
        <a id="L444"></a>execDec(&#34;int32&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L445"></a>if data.a != 17 {
            <a id="L446"></a>t.Errorf(&#34;int32 a = %v not 17&#34;, data.a)
        <a id="L447"></a>}
    <a id="L448"></a>}

    <a id="L450"></a><span class="comment">// uint32</span>
    <a id="L451"></a>{
        <a id="L452"></a>var data struct {
            <a id="L453"></a>a uint32;
        <a id="L454"></a>}
        <a id="L455"></a>instr := &amp;decInstr{decUint32, 6, 0, 0, ovfl};
        <a id="L456"></a>state := newDecodeStateFromData(unsignedResult);
        <a id="L457"></a>execDec(&#34;uint32&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L458"></a>if data.a != 17 {
            <a id="L459"></a>t.Errorf(&#34;uint32 a = %v not 17&#34;, data.a)
        <a id="L460"></a>}
    <a id="L461"></a>}

    <a id="L463"></a><span class="comment">// uintptr</span>
    <a id="L464"></a>{
        <a id="L465"></a>var data struct {
            <a id="L466"></a>a uintptr;
        <a id="L467"></a>}
        <a id="L468"></a>instr := &amp;decInstr{decOpMap[valueKind(data.a)], 6, 0, 0, ovfl};
        <a id="L469"></a>state := newDecodeStateFromData(unsignedResult);
        <a id="L470"></a>execDec(&#34;uintptr&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L471"></a>if data.a != 17 {
            <a id="L472"></a>t.Errorf(&#34;uintptr a = %v not 17&#34;, data.a)
        <a id="L473"></a>}
    <a id="L474"></a>}

    <a id="L476"></a><span class="comment">// int64</span>
    <a id="L477"></a>{
        <a id="L478"></a>var data struct {
            <a id="L479"></a>a int64;
        <a id="L480"></a>}
        <a id="L481"></a>instr := &amp;decInstr{decInt64, 6, 0, 0, ovfl};
        <a id="L482"></a>state := newDecodeStateFromData(signedResult);
        <a id="L483"></a>execDec(&#34;int64&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L484"></a>if data.a != 17 {
            <a id="L485"></a>t.Errorf(&#34;int64 a = %v not 17&#34;, data.a)
        <a id="L486"></a>}
    <a id="L487"></a>}

    <a id="L489"></a><span class="comment">// uint64</span>
    <a id="L490"></a>{
        <a id="L491"></a>var data struct {
            <a id="L492"></a>a uint64;
        <a id="L493"></a>}
        <a id="L494"></a>instr := &amp;decInstr{decUint64, 6, 0, 0, ovfl};
        <a id="L495"></a>state := newDecodeStateFromData(unsignedResult);
        <a id="L496"></a>execDec(&#34;uint64&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L497"></a>if data.a != 17 {
            <a id="L498"></a>t.Errorf(&#34;uint64 a = %v not 17&#34;, data.a)
        <a id="L499"></a>}
    <a id="L500"></a>}

    <a id="L502"></a><span class="comment">// float</span>
    <a id="L503"></a>{
        <a id="L504"></a>var data struct {
            <a id="L505"></a>a float;
        <a id="L506"></a>}
        <a id="L507"></a>instr := &amp;decInstr{decOpMap[valueKind(data.a)], 6, 0, 0, ovfl};
        <a id="L508"></a>state := newDecodeStateFromData(floatResult);
        <a id="L509"></a>execDec(&#34;float&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L510"></a>if data.a != 17 {
            <a id="L511"></a>t.Errorf(&#34;float a = %v not 17&#34;, data.a)
        <a id="L512"></a>}
    <a id="L513"></a>}

    <a id="L515"></a><span class="comment">// float32</span>
    <a id="L516"></a>{
        <a id="L517"></a>var data struct {
            <a id="L518"></a>a float32;
        <a id="L519"></a>}
        <a id="L520"></a>instr := &amp;decInstr{decFloat32, 6, 0, 0, ovfl};
        <a id="L521"></a>state := newDecodeStateFromData(floatResult);
        <a id="L522"></a>execDec(&#34;float32&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L523"></a>if data.a != 17 {
            <a id="L524"></a>t.Errorf(&#34;float32 a = %v not 17&#34;, data.a)
        <a id="L525"></a>}
    <a id="L526"></a>}

    <a id="L528"></a><span class="comment">// float64</span>
    <a id="L529"></a>{
        <a id="L530"></a>var data struct {
            <a id="L531"></a>a float64;
        <a id="L532"></a>}
        <a id="L533"></a>instr := &amp;decInstr{decFloat64, 6, 0, 0, ovfl};
        <a id="L534"></a>state := newDecodeStateFromData(floatResult);
        <a id="L535"></a>execDec(&#34;float64&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L536"></a>if data.a != 17 {
            <a id="L537"></a>t.Errorf(&#34;float64 a = %v not 17&#34;, data.a)
        <a id="L538"></a>}
    <a id="L539"></a>}

    <a id="L541"></a><span class="comment">// bytes == []uint8</span>
    <a id="L542"></a>{
        <a id="L543"></a>var data struct {
            <a id="L544"></a>a []byte;
        <a id="L545"></a>}
        <a id="L546"></a>instr := &amp;decInstr{decUint8Array, 6, 0, 0, ovfl};
        <a id="L547"></a>state := newDecodeStateFromData(bytesResult);
        <a id="L548"></a>execDec(&#34;bytes&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L549"></a>if string(data.a) != &#34;hello&#34; {
            <a id="L550"></a>t.Errorf(`bytes a = %q not &#34;hello&#34;`, string(data.a))
        <a id="L551"></a>}
    <a id="L552"></a>}

    <a id="L554"></a><span class="comment">// string</span>
    <a id="L555"></a>{
        <a id="L556"></a>var data struct {
            <a id="L557"></a>a string;
        <a id="L558"></a>}
        <a id="L559"></a>instr := &amp;decInstr{decString, 6, 0, 0, ovfl};
        <a id="L560"></a>state := newDecodeStateFromData(bytesResult);
        <a id="L561"></a>execDec(&#34;bytes&#34;, instr, state, t, unsafe.Pointer(&amp;data));
        <a id="L562"></a>if data.a != &#34;hello&#34; {
            <a id="L563"></a>t.Errorf(`bytes a = %q not &#34;hello&#34;`, data.a)
        <a id="L564"></a>}
    <a id="L565"></a>}
<a id="L566"></a>}

<a id="L568"></a>func TestEndToEnd(t *testing.T) {
    <a id="L569"></a>type T2 struct {
        <a id="L570"></a>t string;
    <a id="L571"></a>}
    <a id="L572"></a>s1 := &#34;string1&#34;;
    <a id="L573"></a>s2 := &#34;string2&#34;;
    <a id="L574"></a>type T1 struct {
        <a id="L575"></a>a, b, c int;
        <a id="L576"></a>n       *[3]float;
        <a id="L577"></a>strs    *[2]string;
        <a id="L578"></a>int64s  *[]int64;
        <a id="L579"></a>s       string;
        <a id="L580"></a>y       []byte;
        <a id="L581"></a>t       *T2;
    <a id="L582"></a>}
    <a id="L583"></a>t1 := &amp;T1{
        <a id="L584"></a>a: 17,
        <a id="L585"></a>b: 18,
        <a id="L586"></a>c: -5,
        <a id="L587"></a>n: &amp;[3]float{1.5, 2.5, 3.5},
        <a id="L588"></a>strs: &amp;[2]string{s1, s2},
        <a id="L589"></a>int64s: &amp;[]int64{77, 89, 123412342134},
        <a id="L590"></a>s: &#34;Now is the time&#34;,
        <a id="L591"></a>y: strings.Bytes(&#34;hello, sailor&#34;),
        <a id="L592"></a>t: &amp;T2{&#34;this is T2&#34;},
    <a id="L593"></a>};
    <a id="L594"></a>b := new(bytes.Buffer);
    <a id="L595"></a>encode(b, t1);
    <a id="L596"></a>var _t1 T1;
    <a id="L597"></a>decode(b, getTypeInfoNoError(reflect.Typeof(_t1)).id, &amp;_t1);
    <a id="L598"></a>if !reflect.DeepEqual(t1, &amp;_t1) {
        <a id="L599"></a>t.Errorf(&#34;encode expected %v got %v&#34;, *t1, _t1)
    <a id="L600"></a>}
<a id="L601"></a>}

<a id="L603"></a>func TestOverflow(t *testing.T) {
    <a id="L604"></a>type inputT struct {
        <a id="L605"></a>maxi int64;
        <a id="L606"></a>mini int64;
        <a id="L607"></a>maxu uint64;
        <a id="L608"></a>maxf float64;
        <a id="L609"></a>minf float64;
    <a id="L610"></a>}
    <a id="L611"></a>var it inputT;
    <a id="L612"></a>var err os.Error;
    <a id="L613"></a>id := getTypeInfoNoError(reflect.Typeof(it)).id;
    <a id="L614"></a>b := new(bytes.Buffer);

    <a id="L616"></a><span class="comment">// int8</span>
    <a id="L617"></a>b.Reset();
    <a id="L618"></a>it = inputT{
        <a id="L619"></a>maxi: math.MaxInt8 + 1,
    <a id="L620"></a>};
    <a id="L621"></a>type outi8 struct {
        <a id="L622"></a>maxi int8;
        <a id="L623"></a>mini int8;
    <a id="L624"></a>}
    <a id="L625"></a>var o1 outi8;
    <a id="L626"></a>encode(b, it);
    <a id="L627"></a>err = decode(b, id, &amp;o1);
    <a id="L628"></a>if err == nil || err.String() != `value for &#34;maxi&#34; out of range` {
        <a id="L629"></a>t.Error(&#34;wrong overflow error for int8:&#34;, err)
    <a id="L630"></a>}
    <a id="L631"></a>it = inputT{
        <a id="L632"></a>mini: math.MinInt8 - 1,
    <a id="L633"></a>};
    <a id="L634"></a>b.Reset();
    <a id="L635"></a>encode(b, it);
    <a id="L636"></a>err = decode(b, id, &amp;o1);
    <a id="L637"></a>if err == nil || err.String() != `value for &#34;mini&#34; out of range` {
        <a id="L638"></a>t.Error(&#34;wrong underflow error for int8:&#34;, err)
    <a id="L639"></a>}

    <a id="L641"></a><span class="comment">// int16</span>
    <a id="L642"></a>b.Reset();
    <a id="L643"></a>it = inputT{
        <a id="L644"></a>maxi: math.MaxInt16 + 1,
    <a id="L645"></a>};
    <a id="L646"></a>type outi16 struct {
        <a id="L647"></a>maxi int16;
        <a id="L648"></a>mini int16;
    <a id="L649"></a>}
    <a id="L650"></a>var o2 outi16;
    <a id="L651"></a>encode(b, it);
    <a id="L652"></a>err = decode(b, id, &amp;o2);
    <a id="L653"></a>if err == nil || err.String() != `value for &#34;maxi&#34; out of range` {
        <a id="L654"></a>t.Error(&#34;wrong overflow error for int16:&#34;, err)
    <a id="L655"></a>}
    <a id="L656"></a>it = inputT{
        <a id="L657"></a>mini: math.MinInt16 - 1,
    <a id="L658"></a>};
    <a id="L659"></a>b.Reset();
    <a id="L660"></a>encode(b, it);
    <a id="L661"></a>err = decode(b, id, &amp;o2);
    <a id="L662"></a>if err == nil || err.String() != `value for &#34;mini&#34; out of range` {
        <a id="L663"></a>t.Error(&#34;wrong underflow error for int16:&#34;, err)
    <a id="L664"></a>}

    <a id="L666"></a><span class="comment">// int32</span>
    <a id="L667"></a>b.Reset();
    <a id="L668"></a>it = inputT{
        <a id="L669"></a>maxi: math.MaxInt32 + 1,
    <a id="L670"></a>};
    <a id="L671"></a>type outi32 struct {
        <a id="L672"></a>maxi int32;
        <a id="L673"></a>mini int32;
    <a id="L674"></a>}
    <a id="L675"></a>var o3 outi32;
    <a id="L676"></a>encode(b, it);
    <a id="L677"></a>err = decode(b, id, &amp;o3);
    <a id="L678"></a>if err == nil || err.String() != `value for &#34;maxi&#34; out of range` {
        <a id="L679"></a>t.Error(&#34;wrong overflow error for int32:&#34;, err)
    <a id="L680"></a>}
    <a id="L681"></a>it = inputT{
        <a id="L682"></a>mini: math.MinInt32 - 1,
    <a id="L683"></a>};
    <a id="L684"></a>b.Reset();
    <a id="L685"></a>encode(b, it);
    <a id="L686"></a>err = decode(b, id, &amp;o3);
    <a id="L687"></a>if err == nil || err.String() != `value for &#34;mini&#34; out of range` {
        <a id="L688"></a>t.Error(&#34;wrong underflow error for int32:&#34;, err)
    <a id="L689"></a>}

    <a id="L691"></a><span class="comment">// uint8</span>
    <a id="L692"></a>b.Reset();
    <a id="L693"></a>it = inputT{
        <a id="L694"></a>maxu: math.MaxUint8 + 1,
    <a id="L695"></a>};
    <a id="L696"></a>type outu8 struct {
        <a id="L697"></a>maxu uint8;
    <a id="L698"></a>}
    <a id="L699"></a>var o4 outu8;
    <a id="L700"></a>encode(b, it);
    <a id="L701"></a>err = decode(b, id, &amp;o4);
    <a id="L702"></a>if err == nil || err.String() != `value for &#34;maxu&#34; out of range` {
        <a id="L703"></a>t.Error(&#34;wrong overflow error for uint8:&#34;, err)
    <a id="L704"></a>}

    <a id="L706"></a><span class="comment">// uint16</span>
    <a id="L707"></a>b.Reset();
    <a id="L708"></a>it = inputT{
        <a id="L709"></a>maxu: math.MaxUint16 + 1,
    <a id="L710"></a>};
    <a id="L711"></a>type outu16 struct {
        <a id="L712"></a>maxu uint16;
    <a id="L713"></a>}
    <a id="L714"></a>var o5 outu16;
    <a id="L715"></a>encode(b, it);
    <a id="L716"></a>err = decode(b, id, &amp;o5);
    <a id="L717"></a>if err == nil || err.String() != `value for &#34;maxu&#34; out of range` {
        <a id="L718"></a>t.Error(&#34;wrong overflow error for uint16:&#34;, err)
    <a id="L719"></a>}

    <a id="L721"></a><span class="comment">// uint32</span>
    <a id="L722"></a>b.Reset();
    <a id="L723"></a>it = inputT{
        <a id="L724"></a>maxu: math.MaxUint32 + 1,
    <a id="L725"></a>};
    <a id="L726"></a>type outu32 struct {
        <a id="L727"></a>maxu uint32;
    <a id="L728"></a>}
    <a id="L729"></a>var o6 outu32;
    <a id="L730"></a>encode(b, it);
    <a id="L731"></a>err = decode(b, id, &amp;o6);
    <a id="L732"></a>if err == nil || err.String() != `value for &#34;maxu&#34; out of range` {
        <a id="L733"></a>t.Error(&#34;wrong overflow error for uint32:&#34;, err)
    <a id="L734"></a>}

    <a id="L736"></a><span class="comment">// float32</span>
    <a id="L737"></a>b.Reset();
    <a id="L738"></a>it = inputT{
        <a id="L739"></a>maxf: math.MaxFloat32 * 2,
    <a id="L740"></a>};
    <a id="L741"></a>type outf32 struct {
        <a id="L742"></a>maxf float32;
        <a id="L743"></a>minf float32;
    <a id="L744"></a>}
    <a id="L745"></a>var o7 outf32;
    <a id="L746"></a>encode(b, it);
    <a id="L747"></a>err = decode(b, id, &amp;o7);
    <a id="L748"></a>if err == nil || err.String() != `value for &#34;maxf&#34; out of range` {
        <a id="L749"></a>t.Error(&#34;wrong overflow error for float32:&#34;, err)
    <a id="L750"></a>}
<a id="L751"></a>}


<a id="L754"></a>func TestNesting(t *testing.T) {
    <a id="L755"></a>type RT struct {
        <a id="L756"></a>a    string;
        <a id="L757"></a>next *RT;
    <a id="L758"></a>}
    <a id="L759"></a>rt := new(RT);
    <a id="L760"></a>rt.a = &#34;level1&#34;;
    <a id="L761"></a>rt.next = new(RT);
    <a id="L762"></a>rt.next.a = &#34;level2&#34;;
    <a id="L763"></a>b := new(bytes.Buffer);
    <a id="L764"></a>encode(b, rt);
    <a id="L765"></a>var drt RT;
    <a id="L766"></a>decode(b, getTypeInfoNoError(reflect.Typeof(drt)).id, &amp;drt);
    <a id="L767"></a>if drt.a != rt.a {
        <a id="L768"></a>t.Errorf(&#34;nesting: encode expected %v got %v&#34;, *rt, drt)
    <a id="L769"></a>}
    <a id="L770"></a>if drt.next == nil {
        <a id="L771"></a>t.Errorf(&#34;nesting: recursion failed&#34;)
    <a id="L772"></a>}
    <a id="L773"></a>if drt.next.a != rt.next.a {
        <a id="L774"></a>t.Errorf(&#34;nesting: encode expected %v got %v&#34;, *rt.next, *drt.next)
    <a id="L775"></a>}
<a id="L776"></a>}

<a id="L778"></a><span class="comment">// These three structures have the same data with different indirections</span>
<a id="L779"></a>type T0 struct {
    <a id="L780"></a>a   int;
    <a id="L781"></a>b   int;
    <a id="L782"></a>c   int;
    <a id="L783"></a>d   int;
<a id="L784"></a>}
<a id="L785"></a>type T1 struct {
    <a id="L786"></a>a   int;
    <a id="L787"></a>b   *int;
    <a id="L788"></a>c   **int;
    <a id="L789"></a>d   ***int;
<a id="L790"></a>}
<a id="L791"></a>type T2 struct {
    <a id="L792"></a>a   ***int;
    <a id="L793"></a>b   **int;
    <a id="L794"></a>c   *int;
    <a id="L795"></a>d   int;
<a id="L796"></a>}

<a id="L798"></a>func TestAutoIndirection(t *testing.T) {
    <a id="L799"></a><span class="comment">// First transfer t1 into t0</span>
    <a id="L800"></a>var t1 T1;
    <a id="L801"></a>t1.a = 17;
    <a id="L802"></a>t1.b = new(int);
    <a id="L803"></a>*t1.b = 177;
    <a id="L804"></a>t1.c = new(*int);
    <a id="L805"></a>*t1.c = new(int);
    <a id="L806"></a>**t1.c = 1777;
    <a id="L807"></a>t1.d = new(**int);
    <a id="L808"></a>*t1.d = new(*int);
    <a id="L809"></a>**t1.d = new(int);
    <a id="L810"></a>***t1.d = 17777;
    <a id="L811"></a>b := new(bytes.Buffer);
    <a id="L812"></a>encode(b, t1);
    <a id="L813"></a>var t0 T0;
    <a id="L814"></a>t0Id := getTypeInfoNoError(reflect.Typeof(t0)).id;
    <a id="L815"></a>decode(b, t0Id, &amp;t0);
    <a id="L816"></a>if t0.a != 17 || t0.b != 177 || t0.c != 1777 || t0.d != 17777 {
        <a id="L817"></a>t.Errorf(&#34;t1-&gt;t0: expected {17 177 1777 17777}; got %v&#34;, t0)
    <a id="L818"></a>}

    <a id="L820"></a><span class="comment">// Now transfer t2 into t0</span>
    <a id="L821"></a>var t2 T2;
    <a id="L822"></a>t2.d = 17777;
    <a id="L823"></a>t2.c = new(int);
    <a id="L824"></a>*t2.c = 1777;
    <a id="L825"></a>t2.b = new(*int);
    <a id="L826"></a>*t2.b = new(int);
    <a id="L827"></a>**t2.b = 177;
    <a id="L828"></a>t2.a = new(**int);
    <a id="L829"></a>*t2.a = new(*int);
    <a id="L830"></a>**t2.a = new(int);
    <a id="L831"></a>***t2.a = 17;
    <a id="L832"></a>b.Reset();
    <a id="L833"></a>encode(b, t2);
    <a id="L834"></a>t0 = T0{};
    <a id="L835"></a>decode(b, t0Id, &amp;t0);
    <a id="L836"></a>if t0.a != 17 || t0.b != 177 || t0.c != 1777 || t0.d != 17777 {
        <a id="L837"></a>t.Errorf(&#34;t2-&gt;t0 expected {17 177 1777 17777}; got %v&#34;, t0)
    <a id="L838"></a>}

    <a id="L840"></a><span class="comment">// Now transfer t0 into t1</span>
    <a id="L841"></a>t0 = T0{17, 177, 1777, 17777};
    <a id="L842"></a>b.Reset();
    <a id="L843"></a>encode(b, t0);
    <a id="L844"></a>t1 = T1{};
    <a id="L845"></a>t1Id := getTypeInfoNoError(reflect.Typeof(t1)).id;
    <a id="L846"></a>decode(b, t1Id, &amp;t1);
    <a id="L847"></a>if t1.a != 17 || *t1.b != 177 || **t1.c != 1777 || ***t1.d != 17777 {
        <a id="L848"></a>t.Errorf(&#34;t0-&gt;t1 expected {17 177 1777 17777}; got {%d %d %d %d}&#34;, t1.a, *t1.b, **t1.c, ***t1.d)
    <a id="L849"></a>}

    <a id="L851"></a><span class="comment">// Now transfer t0 into t2</span>
    <a id="L852"></a>b.Reset();
    <a id="L853"></a>encode(b, t0);
    <a id="L854"></a>t2 = T2{};
    <a id="L855"></a>t2Id := getTypeInfoNoError(reflect.Typeof(t2)).id;
    <a id="L856"></a>decode(b, t2Id, &amp;t2);
    <a id="L857"></a>if ***t2.a != 17 || **t2.b != 177 || *t2.c != 1777 || t2.d != 17777 {
        <a id="L858"></a>t.Errorf(&#34;t0-&gt;t2 expected {17 177 1777 17777}; got {%d %d %d %d}&#34;, ***t2.a, **t2.b, *t2.c, t2.d)
    <a id="L859"></a>}

    <a id="L861"></a><span class="comment">// Now do t2 again but without pre-allocated pointers.</span>
    <a id="L862"></a>b.Reset();
    <a id="L863"></a>encode(b, t0);
    <a id="L864"></a>***t2.a = 0;
    <a id="L865"></a>**t2.b = 0;
    <a id="L866"></a>*t2.c = 0;
    <a id="L867"></a>t2.d = 0;
    <a id="L868"></a>decode(b, t2Id, &amp;t2);
    <a id="L869"></a>if ***t2.a != 17 || **t2.b != 177 || *t2.c != 1777 || t2.d != 17777 {
        <a id="L870"></a>t.Errorf(&#34;t0-&gt;t2 expected {17 177 1777 17777}; got {%d %d %d %d}&#34;, ***t2.a, **t2.b, *t2.c, t2.d)
    <a id="L871"></a>}
<a id="L872"></a>}

<a id="L874"></a>type RT0 struct {
    <a id="L875"></a>a   int;
    <a id="L876"></a>b   string;
    <a id="L877"></a>c   float;
<a id="L878"></a>}
<a id="L879"></a>type RT1 struct {
    <a id="L880"></a>c      float;
    <a id="L881"></a>b      string;
    <a id="L882"></a>a      int;
    <a id="L883"></a>notSet string;
<a id="L884"></a>}

<a id="L886"></a>func TestReorderedFields(t *testing.T) {
    <a id="L887"></a>var rt0 RT0;
    <a id="L888"></a>rt0.a = 17;
    <a id="L889"></a>rt0.b = &#34;hello&#34;;
    <a id="L890"></a>rt0.c = 3.14159;
    <a id="L891"></a>b := new(bytes.Buffer);
    <a id="L892"></a>encode(b, rt0);
    <a id="L893"></a>rt0Id := getTypeInfoNoError(reflect.Typeof(rt0)).id;
    <a id="L894"></a>var rt1 RT1;
    <a id="L895"></a><span class="comment">// Wire type is RT0, local type is RT1.</span>
    <a id="L896"></a>decode(b, rt0Id, &amp;rt1);
    <a id="L897"></a>if rt0.a != rt1.a || rt0.b != rt1.b || rt0.c != rt1.c {
        <a id="L898"></a>t.Errorf(&#34;rt1-&gt;rt0: expected %v; got %v&#34;, rt0, rt1)
    <a id="L899"></a>}
<a id="L900"></a>}

<a id="L902"></a><span class="comment">// Like an RT0 but with fields we&#39;ll ignore on the decode side.</span>
<a id="L903"></a>type IT0 struct {
    <a id="L904"></a>a        int64;
    <a id="L905"></a>b        string;
    <a id="L906"></a>ignore_d []int;
    <a id="L907"></a>ignore_e [3]float;
    <a id="L908"></a>ignore_f bool;
    <a id="L909"></a>ignore_g string;
    <a id="L910"></a>ignore_h []byte;
    <a id="L911"></a>ignore_i *RT1;
    <a id="L912"></a>c        float;
<a id="L913"></a>}

<a id="L915"></a>func TestIgnoredFields(t *testing.T) {
    <a id="L916"></a>var it0 IT0;
    <a id="L917"></a>it0.a = 17;
    <a id="L918"></a>it0.b = &#34;hello&#34;;
    <a id="L919"></a>it0.c = 3.14159;
    <a id="L920"></a>it0.ignore_d = []int{1, 2, 3};
    <a id="L921"></a>it0.ignore_e[0] = 1.0;
    <a id="L922"></a>it0.ignore_e[1] = 2.0;
    <a id="L923"></a>it0.ignore_e[2] = 3.0;
    <a id="L924"></a>it0.ignore_f = true;
    <a id="L925"></a>it0.ignore_g = &#34;pay no attention&#34;;
    <a id="L926"></a>it0.ignore_h = strings.Bytes(&#34;to the curtain&#34;);
    <a id="L927"></a>it0.ignore_i = &amp;RT1{3.1, &#34;hi&#34;, 7, &#34;hello&#34;};

    <a id="L929"></a>b := new(bytes.Buffer);
    <a id="L930"></a>encode(b, it0);
    <a id="L931"></a>rt0Id := getTypeInfoNoError(reflect.Typeof(it0)).id;
    <a id="L932"></a>var rt1 RT1;
    <a id="L933"></a><span class="comment">// Wire type is IT0, local type is RT1.</span>
    <a id="L934"></a>err := decode(b, rt0Id, &amp;rt1);
    <a id="L935"></a>if err != nil {
        <a id="L936"></a>t.Error(&#34;error: &#34;, err)
    <a id="L937"></a>}
    <a id="L938"></a>if int(it0.a) != rt1.a || it0.b != rt1.b || it0.c != rt1.c {
        <a id="L939"></a>t.Errorf(&#34;rt1-&gt;rt0: expected %v; got %v&#34;, it0, rt1)
    <a id="L940"></a>}
<a id="L941"></a>}

<a id="L943"></a>type Bad0 struct {
    <a id="L944"></a>inter interface{};
    <a id="L945"></a>c     float;
<a id="L946"></a>}

<a id="L948"></a>func TestInvalidField(t *testing.T) {
    <a id="L949"></a>var bad0 Bad0;
    <a id="L950"></a>bad0.inter = 17;
    <a id="L951"></a>b := new(bytes.Buffer);
    <a id="L952"></a>err := encode(b, &amp;bad0);
    <a id="L953"></a>if err == nil {
        <a id="L954"></a>t.Error(&#34;expected error; got none&#34;)
    <a id="L955"></a>} else if strings.Index(err.String(), &#34;interface&#34;) &lt; 0 {
        <a id="L956"></a>t.Error(&#34;expected type error; got&#34;, err)
    <a id="L957"></a>}
<a id="L958"></a>}
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
