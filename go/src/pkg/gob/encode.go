<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/gob/encode.go</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/gob/encode.go</h1>

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
    <a id="L9"></a>&#34;io&#34;;
    <a id="L10"></a>&#34;math&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;reflect&#34;;
    <a id="L13"></a>&#34;unsafe&#34;;
<a id="L14"></a>)

<a id="L16"></a>const uint64Size = unsafe.Sizeof(uint64(0))

<a id="L18"></a><span class="comment">// The global execution state of an instance of the encoder.</span>
<a id="L19"></a><span class="comment">// Field numbers are delta encoded and always increase. The field</span>
<a id="L20"></a><span class="comment">// number is initialized to -1 so 0 comes out as delta(1). A delta of</span>
<a id="L21"></a><span class="comment">// 0 terminates the structure.</span>
<a id="L22"></a>type encoderState struct {
    <a id="L23"></a>b        *bytes.Buffer;
    <a id="L24"></a>err      os.Error;             <span class="comment">// error encountered during encoding;</span>
    <a id="L25"></a>fieldnum int;                  <span class="comment">// the last field number written.</span>
    <a id="L26"></a>buf      [1 + uint64Size]byte; <span class="comment">// buffer used by the encoder; here to avoid allocation.</span>
<a id="L27"></a>}

<a id="L29"></a><span class="comment">// Unsigned integers have a two-state encoding.  If the number is less</span>
<a id="L30"></a><span class="comment">// than 128 (0 through 0x7F), its value is written directly.</span>
<a id="L31"></a><span class="comment">// Otherwise the value is written in big-endian byte order preceded</span>
<a id="L32"></a><span class="comment">// by the byte length, negated.</span>

<a id="L34"></a><span class="comment">// encodeUint writes an encoded unsigned integer to state.b.  Sets state.err.</span>
<a id="L35"></a><span class="comment">// If state.err is already non-nil, it does nothing.</span>
<a id="L36"></a>func encodeUint(state *encoderState, x uint64) {
    <a id="L37"></a>if state.err != nil {
        <a id="L38"></a>return
    <a id="L39"></a>}
    <a id="L40"></a>if x &lt;= 0x7F {
        <a id="L41"></a>state.err = state.b.WriteByte(uint8(x));
        <a id="L42"></a>return;
    <a id="L43"></a>}
    <a id="L44"></a>var n, m int;
    <a id="L45"></a>m = uint64Size;
    <a id="L46"></a>for n = 1; x &gt; 0; n++ {
        <a id="L47"></a>state.buf[m] = uint8(x &amp; 0xFF);
        <a id="L48"></a>x &gt;&gt;= 8;
        <a id="L49"></a>m--;
    <a id="L50"></a>}
    <a id="L51"></a>state.buf[m] = uint8(-(n - 1));
    <a id="L52"></a>n, state.err = state.b.Write(state.buf[m : uint64Size+1]);
<a id="L53"></a>}

<a id="L55"></a><span class="comment">// encodeInt writes an encoded signed integer to state.w.</span>
<a id="L56"></a><span class="comment">// The low bit of the encoding says whether to bit complement the (other bits of the) uint to recover the int.</span>
<a id="L57"></a><span class="comment">// Sets state.err. If state.err is already non-nil, it does nothing.</span>
<a id="L58"></a>func encodeInt(state *encoderState, i int64) {
    <a id="L59"></a>var x uint64;
    <a id="L60"></a>if i &lt; 0 {
        <a id="L61"></a>x = uint64(^i&lt;&lt;1) | 1
    <a id="L62"></a>} else {
        <a id="L63"></a>x = uint64(i &lt;&lt; 1)
    <a id="L64"></a>}
    <a id="L65"></a>encodeUint(state, uint64(x));
<a id="L66"></a>}

<a id="L68"></a>type encOp func(i *encInstr, state *encoderState, p unsafe.Pointer)

<a id="L70"></a><span class="comment">// The &#39;instructions&#39; of the encoding machine</span>
<a id="L71"></a>type encInstr struct {
    <a id="L72"></a>op     encOp;
    <a id="L73"></a>field  int;     <span class="comment">// field number</span>
    <a id="L74"></a>indir  int;     <span class="comment">// how many pointer indirections to reach the value in the struct</span>
    <a id="L75"></a>offset uintptr; <span class="comment">// offset in the structure of the field to encode</span>
<a id="L76"></a>}

<a id="L78"></a><span class="comment">// Emit a field number and update the state to record its value for delta encoding.</span>
<a id="L79"></a><span class="comment">// If the instruction pointer is nil, do nothing</span>
<a id="L80"></a>func (state *encoderState) update(instr *encInstr) {
    <a id="L81"></a>if instr != nil {
        <a id="L82"></a>encodeUint(state, uint64(instr.field-state.fieldnum));
        <a id="L83"></a>state.fieldnum = instr.field;
    <a id="L84"></a>}
<a id="L85"></a>}

<a id="L87"></a><span class="comment">// Each encoder is responsible for handling any indirections associated</span>
<a id="L88"></a><span class="comment">// with the data structure.  If any pointer so reached is nil, no bytes are written.</span>
<a id="L89"></a><span class="comment">// If the data item is zero, no bytes are written.</span>
<a id="L90"></a><span class="comment">// Otherwise, the output (for a scalar) is the field number, as an encoded integer,</span>
<a id="L91"></a><span class="comment">// followed by the field data in its appropriate format.</span>

<a id="L93"></a>func encIndirect(p unsafe.Pointer, indir int) unsafe.Pointer {
    <a id="L94"></a>for ; indir &gt; 0; indir-- {
        <a id="L95"></a>p = *(*unsafe.Pointer)(p);
        <a id="L96"></a>if p == nil {
            <a id="L97"></a>return unsafe.Pointer(nil)
        <a id="L98"></a>}
    <a id="L99"></a>}
    <a id="L100"></a>return p;
<a id="L101"></a>}

<a id="L103"></a>func encBool(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L104"></a>b := *(*bool)(p);
    <a id="L105"></a>if b {
        <a id="L106"></a>state.update(i);
        <a id="L107"></a>encodeUint(state, 1);
    <a id="L108"></a>}
<a id="L109"></a>}

<a id="L111"></a>func encInt(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L112"></a>v := int64(*(*int)(p));
    <a id="L113"></a>if v != 0 {
        <a id="L114"></a>state.update(i);
        <a id="L115"></a>encodeInt(state, v);
    <a id="L116"></a>}
<a id="L117"></a>}

<a id="L119"></a>func encUint(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L120"></a>v := uint64(*(*uint)(p));
    <a id="L121"></a>if v != 0 {
        <a id="L122"></a>state.update(i);
        <a id="L123"></a>encodeUint(state, v);
    <a id="L124"></a>}
<a id="L125"></a>}

<a id="L127"></a>func encInt8(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L128"></a>v := int64(*(*int8)(p));
    <a id="L129"></a>if v != 0 {
        <a id="L130"></a>state.update(i);
        <a id="L131"></a>encodeInt(state, v);
    <a id="L132"></a>}
<a id="L133"></a>}

<a id="L135"></a>func encUint8(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L136"></a>v := uint64(*(*uint8)(p));
    <a id="L137"></a>if v != 0 {
        <a id="L138"></a>state.update(i);
        <a id="L139"></a>encodeUint(state, v);
    <a id="L140"></a>}
<a id="L141"></a>}

<a id="L143"></a>func encInt16(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L144"></a>v := int64(*(*int16)(p));
    <a id="L145"></a>if v != 0 {
        <a id="L146"></a>state.update(i);
        <a id="L147"></a>encodeInt(state, v);
    <a id="L148"></a>}
<a id="L149"></a>}

<a id="L151"></a>func encUint16(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L152"></a>v := uint64(*(*uint16)(p));
    <a id="L153"></a>if v != 0 {
        <a id="L154"></a>state.update(i);
        <a id="L155"></a>encodeUint(state, v);
    <a id="L156"></a>}
<a id="L157"></a>}

<a id="L159"></a>func encInt32(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L160"></a>v := int64(*(*int32)(p));
    <a id="L161"></a>if v != 0 {
        <a id="L162"></a>state.update(i);
        <a id="L163"></a>encodeInt(state, v);
    <a id="L164"></a>}
<a id="L165"></a>}

<a id="L167"></a>func encUint32(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L168"></a>v := uint64(*(*uint32)(p));
    <a id="L169"></a>if v != 0 {
        <a id="L170"></a>state.update(i);
        <a id="L171"></a>encodeUint(state, v);
    <a id="L172"></a>}
<a id="L173"></a>}

<a id="L175"></a>func encInt64(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L176"></a>v := *(*int64)(p);
    <a id="L177"></a>if v != 0 {
        <a id="L178"></a>state.update(i);
        <a id="L179"></a>encodeInt(state, v);
    <a id="L180"></a>}
<a id="L181"></a>}

<a id="L183"></a>func encUint64(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L184"></a>v := *(*uint64)(p);
    <a id="L185"></a>if v != 0 {
        <a id="L186"></a>state.update(i);
        <a id="L187"></a>encodeUint(state, v);
    <a id="L188"></a>}
<a id="L189"></a>}

<a id="L191"></a>func encUintptr(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L192"></a>v := uint64(*(*uintptr)(p));
    <a id="L193"></a>if v != 0 {
        <a id="L194"></a>state.update(i);
        <a id="L195"></a>encodeUint(state, v);
    <a id="L196"></a>}
<a id="L197"></a>}

<a id="L199"></a><span class="comment">// Floating-point numbers are transmitted as uint64s holding the bits</span>
<a id="L200"></a><span class="comment">// of the underlying representation.  They are sent byte-reversed, with</span>
<a id="L201"></a><span class="comment">// the exponent end coming out first, so integer floating point numbers</span>
<a id="L202"></a><span class="comment">// (for example) transmit more compactly.  This routine does the</span>
<a id="L203"></a><span class="comment">// swizzling.</span>
<a id="L204"></a>func floatBits(f float64) uint64 {
    <a id="L205"></a>u := math.Float64bits(f);
    <a id="L206"></a>var v uint64;
    <a id="L207"></a>for i := 0; i &lt; 8; i++ {
        <a id="L208"></a>v &lt;&lt;= 8;
        <a id="L209"></a>v |= u &amp; 0xFF;
        <a id="L210"></a>u &gt;&gt;= 8;
    <a id="L211"></a>}
    <a id="L212"></a>return v;
<a id="L213"></a>}

<a id="L215"></a>func encFloat(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L216"></a>f := float(*(*float)(p));
    <a id="L217"></a>if f != 0 {
        <a id="L218"></a>v := floatBits(float64(f));
        <a id="L219"></a>state.update(i);
        <a id="L220"></a>encodeUint(state, v);
    <a id="L221"></a>}
<a id="L222"></a>}

<a id="L224"></a>func encFloat32(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L225"></a>f := float32(*(*float32)(p));
    <a id="L226"></a>if f != 0 {
        <a id="L227"></a>v := floatBits(float64(f));
        <a id="L228"></a>state.update(i);
        <a id="L229"></a>encodeUint(state, v);
    <a id="L230"></a>}
<a id="L231"></a>}

<a id="L233"></a>func encFloat64(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L234"></a>f := *(*float64)(p);
    <a id="L235"></a>if f != 0 {
        <a id="L236"></a>state.update(i);
        <a id="L237"></a>v := floatBits(f);
        <a id="L238"></a>encodeUint(state, v);
    <a id="L239"></a>}
<a id="L240"></a>}

<a id="L242"></a><span class="comment">// Byte arrays are encoded as an unsigned count followed by the raw bytes.</span>
<a id="L243"></a>func encUint8Array(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L244"></a>b := *(*[]byte)(p);
    <a id="L245"></a>if len(b) &gt; 0 {
        <a id="L246"></a>state.update(i);
        <a id="L247"></a>encodeUint(state, uint64(len(b)));
        <a id="L248"></a>state.b.Write(b);
    <a id="L249"></a>}
<a id="L250"></a>}

<a id="L252"></a><span class="comment">// Strings are encoded as an unsigned count followed by the raw bytes.</span>
<a id="L253"></a>func encString(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L254"></a>s := *(*string)(p);
    <a id="L255"></a>if len(s) &gt; 0 {
        <a id="L256"></a>state.update(i);
        <a id="L257"></a>encodeUint(state, uint64(len(s)));
        <a id="L258"></a>io.WriteString(state.b, s);
    <a id="L259"></a>}
<a id="L260"></a>}

<a id="L262"></a><span class="comment">// The end of a struct is marked by a delta field number of 0.</span>
<a id="L263"></a>func encStructTerminator(i *encInstr, state *encoderState, p unsafe.Pointer) {
    <a id="L264"></a>encodeUint(state, 0)
<a id="L265"></a>}

<a id="L267"></a><span class="comment">// Execution engine</span>

<a id="L269"></a><span class="comment">// The encoder engine is an array of instructions indexed by field number of the encoding</span>
<a id="L270"></a><span class="comment">// data, typically a struct.  It is executed top to bottom, walking the struct.</span>
<a id="L271"></a>type encEngine struct {
    <a id="L272"></a>instr []encInstr;
<a id="L273"></a>}

<a id="L275"></a>func encodeStruct(engine *encEngine, b *bytes.Buffer, basep uintptr) os.Error {
    <a id="L276"></a>state := new(encoderState);
    <a id="L277"></a>state.b = b;
    <a id="L278"></a>state.fieldnum = -1;
    <a id="L279"></a>for i := 0; i &lt; len(engine.instr); i++ {
        <a id="L280"></a>instr := &amp;engine.instr[i];
        <a id="L281"></a>p := unsafe.Pointer(basep + instr.offset);
        <a id="L282"></a>if instr.indir &gt; 0 {
            <a id="L283"></a>if p = encIndirect(p, instr.indir); p == nil {
                <a id="L284"></a>continue
            <a id="L285"></a>}
        <a id="L286"></a>}
        <a id="L287"></a>instr.op(instr, state, p);
        <a id="L288"></a>if state.err != nil {
            <a id="L289"></a>break
        <a id="L290"></a>}
    <a id="L291"></a>}
    <a id="L292"></a>return state.err;
<a id="L293"></a>}

<a id="L295"></a>func encodeArray(b *bytes.Buffer, p uintptr, op encOp, elemWid uintptr, length int, elemIndir int) os.Error {
    <a id="L296"></a>state := new(encoderState);
    <a id="L297"></a>state.b = b;
    <a id="L298"></a>state.fieldnum = -1;
    <a id="L299"></a>encodeUint(state, uint64(length));
    <a id="L300"></a>for i := 0; i &lt; length &amp;&amp; state.err == nil; i++ {
        <a id="L301"></a>elemp := p;
        <a id="L302"></a>up := unsafe.Pointer(elemp);
        <a id="L303"></a>if elemIndir &gt; 0 {
            <a id="L304"></a>if up = encIndirect(up, elemIndir); up == nil {
                <a id="L305"></a>state.err = os.ErrorString(&#34;gob: encodeArray: nil element&#34;);
                <a id="L306"></a>break;
            <a id="L307"></a>}
            <a id="L308"></a>elemp = uintptr(up);
        <a id="L309"></a>}
        <a id="L310"></a>op(nil, state, unsafe.Pointer(elemp));
        <a id="L311"></a>p += uintptr(elemWid);
    <a id="L312"></a>}
    <a id="L313"></a>return state.err;
<a id="L314"></a>}

<a id="L316"></a>var encOpMap = map[reflect.Type]encOp{
    <a id="L317"></a>valueKind(false): encBool,
    <a id="L318"></a>valueKind(int(0)): encInt,
    <a id="L319"></a>valueKind(int8(0)): encInt8,
    <a id="L320"></a>valueKind(int16(0)): encInt16,
    <a id="L321"></a>valueKind(int32(0)): encInt32,
    <a id="L322"></a>valueKind(int64(0)): encInt64,
    <a id="L323"></a>valueKind(uint(0)): encUint,
    <a id="L324"></a>valueKind(uint8(0)): encUint8,
    <a id="L325"></a>valueKind(uint16(0)): encUint16,
    <a id="L326"></a>valueKind(uint32(0)): encUint32,
    <a id="L327"></a>valueKind(uint64(0)): encUint64,
    <a id="L328"></a>valueKind(uintptr(0)): encUintptr,
    <a id="L329"></a>valueKind(float(0)): encFloat,
    <a id="L330"></a>valueKind(float32(0)): encFloat32,
    <a id="L331"></a>valueKind(float64(0)): encFloat64,
    <a id="L332"></a>valueKind(&#34;x&#34;): encString,
<a id="L333"></a>}

<a id="L335"></a><span class="comment">// Return the encoding op for the base type under rt and</span>
<a id="L336"></a><span class="comment">// the indirection count to reach it.</span>
<a id="L337"></a>func encOpFor(rt reflect.Type) (encOp, int, os.Error) {
    <a id="L338"></a>typ, indir := indirect(rt);
    <a id="L339"></a>op, ok := encOpMap[reflect.Typeof(typ)];
    <a id="L340"></a>if !ok {
        <a id="L341"></a>typ, _ := indirect(rt);
        <a id="L342"></a><span class="comment">// Special cases</span>
        <a id="L343"></a>switch t := typ.(type) {
        <a id="L344"></a>case *reflect.SliceType:
            <a id="L345"></a>if _, ok := t.Elem().(*reflect.Uint8Type); ok {
                <a id="L346"></a>op = encUint8Array;
                <a id="L347"></a>break;
            <a id="L348"></a>}
            <a id="L349"></a><span class="comment">// Slices have a header; we decode it to find the underlying array.</span>
            <a id="L350"></a>elemOp, indir, err := encOpFor(t.Elem());
            <a id="L351"></a>if err != nil {
                <a id="L352"></a>return nil, 0, err
            <a id="L353"></a>}
            <a id="L354"></a>op = func(i *encInstr, state *encoderState, p unsafe.Pointer) {
                <a id="L355"></a>slice := (*reflect.SliceHeader)(p);
                <a id="L356"></a>if slice.Len == 0 {
                    <a id="L357"></a>return
                <a id="L358"></a>}
                <a id="L359"></a>state.update(i);
                <a id="L360"></a>state.err = encodeArray(state.b, slice.Data, elemOp, t.Elem().Size(), int(slice.Len), indir);
            <a id="L361"></a>};
        <a id="L362"></a>case *reflect.ArrayType:
            <a id="L363"></a><span class="comment">// True arrays have size in the type.</span>
            <a id="L364"></a>elemOp, indir, err := encOpFor(t.Elem());
            <a id="L365"></a>if err != nil {
                <a id="L366"></a>return nil, 0, err
            <a id="L367"></a>}
            <a id="L368"></a>op = func(i *encInstr, state *encoderState, p unsafe.Pointer) {
                <a id="L369"></a>state.update(i);
                <a id="L370"></a>state.err = encodeArray(state.b, uintptr(p), elemOp, t.Elem().Size(), t.Len(), indir);
            <a id="L371"></a>};
        <a id="L372"></a>case *reflect.StructType:
            <a id="L373"></a><span class="comment">// Generate a closure that calls out to the engine for the nested type.</span>
            <a id="L374"></a>_, err := getEncEngine(typ);
            <a id="L375"></a>if err != nil {
                <a id="L376"></a>return nil, 0, err
            <a id="L377"></a>}
            <a id="L378"></a>info := getTypeInfoNoError(typ);
            <a id="L379"></a>op = func(i *encInstr, state *encoderState, p unsafe.Pointer) {
                <a id="L380"></a>state.update(i);
                <a id="L381"></a><span class="comment">// indirect through info to delay evaluation for recursive structs</span>
                <a id="L382"></a>state.err = encodeStruct(info.encoder, state.b, uintptr(p));
            <a id="L383"></a>};
        <a id="L384"></a>}
    <a id="L385"></a>}
    <a id="L386"></a>if op == nil {
        <a id="L387"></a>return op, indir, os.ErrorString(&#34;gob enc: can&#39;t happen: encode type&#34; + rt.String())
    <a id="L388"></a>}
    <a id="L389"></a>return op, indir, nil;
<a id="L390"></a>}

<a id="L392"></a><span class="comment">// The local Type was compiled from the actual value, so we know it&#39;s compatible.</span>
<a id="L393"></a>func compileEnc(rt reflect.Type) (*encEngine, os.Error) {
    <a id="L394"></a>srt, ok := rt.(*reflect.StructType);
    <a id="L395"></a>if !ok {
        <a id="L396"></a>panicln(&#34;can&#39;t happen: non-struct&#34;)
    <a id="L397"></a>}
    <a id="L398"></a>engine := new(encEngine);
    <a id="L399"></a>engine.instr = make([]encInstr, srt.NumField()+1); <span class="comment">// +1 for terminator</span>
    <a id="L400"></a>for fieldnum := 0; fieldnum &lt; srt.NumField(); fieldnum++ {
        <a id="L401"></a>f := srt.Field(fieldnum);
        <a id="L402"></a>op, indir, err := encOpFor(f.Type);
        <a id="L403"></a>if err != nil {
            <a id="L404"></a>return nil, err
        <a id="L405"></a>}
        <a id="L406"></a>engine.instr[fieldnum] = encInstr{op, fieldnum, indir, uintptr(f.Offset)};
    <a id="L407"></a>}
    <a id="L408"></a>engine.instr[srt.NumField()] = encInstr{encStructTerminator, 0, 0, 0};
    <a id="L409"></a>return engine, nil;
<a id="L410"></a>}

<a id="L412"></a><span class="comment">// typeLock must be held (or we&#39;re in initialization and guaranteed single-threaded).</span>
<a id="L413"></a><span class="comment">// The reflection type must have all its indirections processed out.</span>
<a id="L414"></a>func getEncEngine(rt reflect.Type) (*encEngine, os.Error) {
    <a id="L415"></a>info, err := getTypeInfo(rt);
    <a id="L416"></a>if err != nil {
        <a id="L417"></a>return nil, err
    <a id="L418"></a>}
    <a id="L419"></a>if info.encoder == nil {
        <a id="L420"></a><span class="comment">// mark this engine as underway before compiling to handle recursive types.</span>
        <a id="L421"></a>info.encoder = new(encEngine);
        <a id="L422"></a>info.encoder, err = compileEnc(rt);
    <a id="L423"></a>}
    <a id="L424"></a>return info.encoder, err;
<a id="L425"></a>}

<a id="L427"></a>func encode(b *bytes.Buffer, e interface{}) os.Error {
    <a id="L428"></a><span class="comment">// Dereference down to the underlying object.</span>
    <a id="L429"></a>rt, indir := indirect(reflect.Typeof(e));
    <a id="L430"></a>v := reflect.NewValue(e);
    <a id="L431"></a>for i := 0; i &lt; indir; i++ {
        <a id="L432"></a>v = reflect.Indirect(v)
    <a id="L433"></a>}
    <a id="L434"></a>if _, ok := v.(*reflect.StructValue); !ok {
        <a id="L435"></a>return os.ErrorString(&#34;gob: encode can&#39;t handle &#34; + v.Type().String())
    <a id="L436"></a>}
    <a id="L437"></a>typeLock.Lock();
    <a id="L438"></a>engine, err := getEncEngine(rt);
    <a id="L439"></a>typeLock.Unlock();
    <a id="L440"></a>if err != nil {
        <a id="L441"></a>return err
    <a id="L442"></a>}
    <a id="L443"></a>return encodeStruct(engine, b, v.Addr());
<a id="L444"></a>}
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
