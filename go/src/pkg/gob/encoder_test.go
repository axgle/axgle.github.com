<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/gob/encoder_test.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/gob/encoder_test.go</h1>

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
    <a id="L10"></a>&#34;os&#34;;
    <a id="L11"></a>&#34;reflect&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a>type ET2 struct {
    <a id="L16"></a>x string;
<a id="L17"></a>}

<a id="L19"></a>type ET1 struct {
    <a id="L20"></a>a    int;
    <a id="L21"></a>et2  *ET2;
    <a id="L22"></a>next *ET1;
<a id="L23"></a>}

<a id="L25"></a><span class="comment">// Like ET1 but with a different name for a field</span>
<a id="L26"></a>type ET3 struct {
    <a id="L27"></a>a             int;
    <a id="L28"></a>et2           *ET2;
    <a id="L29"></a>differentNext *ET1;
<a id="L30"></a>}

<a id="L32"></a><span class="comment">// Like ET1 but with a different type for a field</span>
<a id="L33"></a>type ET4 struct {
    <a id="L34"></a>a    int;
    <a id="L35"></a>et2  *ET1;
    <a id="L36"></a>next int;
<a id="L37"></a>}

<a id="L39"></a>func TestBasicEncoder(t *testing.T) {
    <a id="L40"></a>b := new(bytes.Buffer);
    <a id="L41"></a>enc := NewEncoder(b);
    <a id="L42"></a>et1 := new(ET1);
    <a id="L43"></a>et1.a = 7;
    <a id="L44"></a>et1.et2 = new(ET2);
    <a id="L45"></a>enc.Encode(et1);
    <a id="L46"></a>if enc.state.err != nil {
        <a id="L47"></a>t.Error(&#34;encoder fail:&#34;, enc.state.err)
    <a id="L48"></a>}

    <a id="L50"></a><span class="comment">// Decode the result by hand to verify;</span>
    <a id="L51"></a>state := newDecodeState(b);
    <a id="L52"></a><span class="comment">// The output should be:</span>
    <a id="L53"></a><span class="comment">// 0) The length, 38.</span>
    <a id="L54"></a>length := decodeUint(state);
    <a id="L55"></a>if length != 38 {
        <a id="L56"></a>t.Fatal(&#34;0. expected length 38; got&#34;, length)
    <a id="L57"></a>}
    <a id="L58"></a><span class="comment">// 1) -7: the type id of ET1</span>
    <a id="L59"></a>id1 := decodeInt(state);
    <a id="L60"></a>if id1 &gt;= 0 {
        <a id="L61"></a>t.Fatal(&#34;expected ET1 negative id; got&#34;, id1)
    <a id="L62"></a>}
    <a id="L63"></a><span class="comment">// 2) The wireType for ET1</span>
    <a id="L64"></a>wire1 := new(wireType);
    <a id="L65"></a>err := decode(b, tWireType, wire1);
    <a id="L66"></a>if err != nil {
        <a id="L67"></a>t.Fatal(&#34;error decoding ET1 type:&#34;, err)
    <a id="L68"></a>}
    <a id="L69"></a>info := getTypeInfoNoError(reflect.Typeof(ET1{}));
    <a id="L70"></a>trueWire1 := &amp;wireType{s: info.id.gobType().(*structType)};
    <a id="L71"></a>if !reflect.DeepEqual(wire1, trueWire1) {
        <a id="L72"></a>t.Fatalf(&#34;invalid wireType for ET1: expected %+v; got %+v\n&#34;, *trueWire1, *wire1)
    <a id="L73"></a>}
    <a id="L74"></a><span class="comment">// 3) The length, 21.</span>
    <a id="L75"></a>length = decodeUint(state);
    <a id="L76"></a>if length != 21 {
        <a id="L77"></a>t.Fatal(&#34;3. expected length 21; got&#34;, length)
    <a id="L78"></a>}
    <a id="L79"></a><span class="comment">// 4) -8: the type id of ET2</span>
    <a id="L80"></a>id2 := decodeInt(state);
    <a id="L81"></a>if id2 &gt;= 0 {
        <a id="L82"></a>t.Fatal(&#34;expected ET2 negative id; got&#34;, id2)
    <a id="L83"></a>}
    <a id="L84"></a><span class="comment">// 5) The wireType for ET2</span>
    <a id="L85"></a>wire2 := new(wireType);
    <a id="L86"></a>err = decode(b, tWireType, wire2);
    <a id="L87"></a>if err != nil {
        <a id="L88"></a>t.Fatal(&#34;error decoding ET2 type:&#34;, err)
    <a id="L89"></a>}
    <a id="L90"></a>info = getTypeInfoNoError(reflect.Typeof(ET2{}));
    <a id="L91"></a>trueWire2 := &amp;wireType{s: info.id.gobType().(*structType)};
    <a id="L92"></a>if !reflect.DeepEqual(wire2, trueWire2) {
        <a id="L93"></a>t.Fatalf(&#34;invalid wireType for ET2: expected %+v; got %+v\n&#34;, *trueWire2, *wire2)
    <a id="L94"></a>}
    <a id="L95"></a><span class="comment">// 6) The length, 6.</span>
    <a id="L96"></a>length = decodeUint(state);
    <a id="L97"></a>if length != 6 {
        <a id="L98"></a>t.Fatal(&#34;6. expected length 6; got&#34;, length)
    <a id="L99"></a>}
    <a id="L100"></a><span class="comment">// 7) The type id for the et1 value</span>
    <a id="L101"></a>newId1 := decodeInt(state);
    <a id="L102"></a>if newId1 != -id1 {
        <a id="L103"></a>t.Fatal(&#34;expected Et1 id&#34;, -id1, &#34;got&#34;, newId1)
    <a id="L104"></a>}
    <a id="L105"></a><span class="comment">// 8) The value of et1</span>
    <a id="L106"></a>newEt1 := new(ET1);
    <a id="L107"></a>et1Id := getTypeInfoNoError(reflect.Typeof(*newEt1)).id;
    <a id="L108"></a>err = decode(b, et1Id, newEt1);
    <a id="L109"></a>if err != nil {
        <a id="L110"></a>t.Fatal(&#34;error decoding ET1 value:&#34;, err)
    <a id="L111"></a>}
    <a id="L112"></a>if !reflect.DeepEqual(et1, newEt1) {
        <a id="L113"></a>t.Fatalf(&#34;invalid data for et1: expected %+v; got %+v\n&#34;, *et1, *newEt1)
    <a id="L114"></a>}
    <a id="L115"></a><span class="comment">// 9) EOF</span>
    <a id="L116"></a>if b.Len() != 0 {
        <a id="L117"></a>t.Error(&#34;not at eof;&#34;, b.Len(), &#34;bytes left&#34;)
    <a id="L118"></a>}

    <a id="L120"></a><span class="comment">// Now do it again. This time we should see only the type id and value.</span>
    <a id="L121"></a>b.Reset();
    <a id="L122"></a>enc.Encode(et1);
    <a id="L123"></a>if enc.state.err != nil {
        <a id="L124"></a>t.Error(&#34;2nd round: encoder fail:&#34;, enc.state.err)
    <a id="L125"></a>}
    <a id="L126"></a><span class="comment">// The length.</span>
    <a id="L127"></a>length = decodeUint(state);
    <a id="L128"></a>if length != 6 {
        <a id="L129"></a>t.Fatal(&#34;6. expected length 6; got&#34;, length)
    <a id="L130"></a>}
    <a id="L131"></a><span class="comment">// 5a) The type id for the et1 value</span>
    <a id="L132"></a>newId1 = decodeInt(state);
    <a id="L133"></a>if newId1 != -id1 {
        <a id="L134"></a>t.Fatal(&#34;2nd round: expected Et1 id&#34;, -id1, &#34;got&#34;, newId1)
    <a id="L135"></a>}
    <a id="L136"></a><span class="comment">// 6a) The value of et1</span>
    <a id="L137"></a>newEt1 = new(ET1);
    <a id="L138"></a>err = decode(b, et1Id, newEt1);
    <a id="L139"></a>if err != nil {
        <a id="L140"></a>t.Fatal(&#34;2nd round: error decoding ET1 value:&#34;, err)
    <a id="L141"></a>}
    <a id="L142"></a>if !reflect.DeepEqual(et1, newEt1) {
        <a id="L143"></a>t.Fatalf(&#34;2nd round: invalid data for et1: expected %+v; got %+v\n&#34;, *et1, *newEt1)
    <a id="L144"></a>}
    <a id="L145"></a><span class="comment">// 7a) EOF</span>
    <a id="L146"></a>if b.Len() != 0 {
        <a id="L147"></a>t.Error(&#34;2nd round: not at eof;&#34;, b.Len(), &#34;bytes left&#34;)
    <a id="L148"></a>}
<a id="L149"></a>}

<a id="L151"></a>func TestEncoderDecoder(t *testing.T) {
    <a id="L152"></a>b := new(bytes.Buffer);
    <a id="L153"></a>enc := NewEncoder(b);
    <a id="L154"></a>et1 := new(ET1);
    <a id="L155"></a>et1.a = 7;
    <a id="L156"></a>et1.et2 = new(ET2);
    <a id="L157"></a>enc.Encode(et1);
    <a id="L158"></a>if enc.state.err != nil {
        <a id="L159"></a>t.Error(&#34;encoder fail:&#34;, enc.state.err)
    <a id="L160"></a>}
    <a id="L161"></a>dec := NewDecoder(b);
    <a id="L162"></a>newEt1 := new(ET1);
    <a id="L163"></a>dec.Decode(newEt1);
    <a id="L164"></a>if dec.state.err != nil {
        <a id="L165"></a>t.Fatal(&#34;error decoding ET1:&#34;, dec.state.err)
    <a id="L166"></a>}

    <a id="L168"></a>if !reflect.DeepEqual(et1, newEt1) {
        <a id="L169"></a>t.Fatalf(&#34;invalid data for et1: expected %+v; got %+v\n&#34;, *et1, *newEt1)
    <a id="L170"></a>}
    <a id="L171"></a>if b.Len() != 0 {
        <a id="L172"></a>t.Error(&#34;not at eof;&#34;, b.Len(), &#34;bytes left&#34;)
    <a id="L173"></a>}

    <a id="L175"></a>enc.Encode(et1);
    <a id="L176"></a>newEt1 = new(ET1);
    <a id="L177"></a>dec.Decode(newEt1);
    <a id="L178"></a>if dec.state.err != nil {
        <a id="L179"></a>t.Fatal(&#34;round 2: error decoding ET1:&#34;, dec.state.err)
    <a id="L180"></a>}
    <a id="L181"></a>if !reflect.DeepEqual(et1, newEt1) {
        <a id="L182"></a>t.Fatalf(&#34;round 2: invalid data for et1: expected %+v; got %+v\n&#34;, *et1, *newEt1)
    <a id="L183"></a>}
    <a id="L184"></a>if b.Len() != 0 {
        <a id="L185"></a>t.Error(&#34;round 2: not at eof;&#34;, b.Len(), &#34;bytes left&#34;)
    <a id="L186"></a>}

    <a id="L188"></a><span class="comment">// Now test with a running encoder/decoder pair that we recognize a type mismatch.</span>
    <a id="L189"></a>enc.Encode(et1);
    <a id="L190"></a>if enc.state.err != nil {
        <a id="L191"></a>t.Error(&#34;round 3: encoder fail:&#34;, enc.state.err)
    <a id="L192"></a>}
    <a id="L193"></a>newEt2 := new(ET2);
    <a id="L194"></a>dec.Decode(newEt2);
    <a id="L195"></a>if dec.state.err == nil {
        <a id="L196"></a>t.Fatal(&#34;round 3: expected `bad type&#39; error decoding ET2&#34;)
    <a id="L197"></a>}
<a id="L198"></a>}

<a id="L200"></a><span class="comment">// Run one value through the encoder/decoder, but use the wrong type.</span>
<a id="L201"></a><span class="comment">// Input is always an ET1; we compare it to whatever is under &#39;e&#39;.</span>
<a id="L202"></a>func badTypeCheck(e interface{}, shouldFail bool, msg string, t *testing.T) {
    <a id="L203"></a>b := new(bytes.Buffer);
    <a id="L204"></a>enc := NewEncoder(b);
    <a id="L205"></a>et1 := new(ET1);
    <a id="L206"></a>et1.a = 7;
    <a id="L207"></a>et1.et2 = new(ET2);
    <a id="L208"></a>enc.Encode(et1);
    <a id="L209"></a>if enc.state.err != nil {
        <a id="L210"></a>t.Error(&#34;encoder fail:&#34;, enc.state.err)
    <a id="L211"></a>}
    <a id="L212"></a>dec := NewDecoder(b);
    <a id="L213"></a>dec.Decode(e);
    <a id="L214"></a>if shouldFail &amp;&amp; (dec.state.err == nil) {
        <a id="L215"></a>t.Error(&#34;expected error for&#34;, msg)
    <a id="L216"></a>}
    <a id="L217"></a>if !shouldFail &amp;&amp; (dec.state.err != nil) {
        <a id="L218"></a>t.Error(&#34;unexpected error for&#34;, msg)
    <a id="L219"></a>}
<a id="L220"></a>}

<a id="L222"></a><span class="comment">// Test that we recognize a bad type the first time.</span>
<a id="L223"></a>func TestWrongTypeDecoder(t *testing.T) {
    <a id="L224"></a>badTypeCheck(new(ET2), true, &#34;no fields in common&#34;, t);
    <a id="L225"></a>badTypeCheck(new(ET3), false, &#34;different name of field&#34;, t);
    <a id="L226"></a>badTypeCheck(new(ET4), true, &#34;different type of field&#34;, t);
<a id="L227"></a>}

<a id="L229"></a>func corruptDataCheck(s string, err os.Error, t *testing.T) {
    <a id="L230"></a>b := bytes.NewBufferString(s);
    <a id="L231"></a>dec := NewDecoder(b);
    <a id="L232"></a>dec.Decode(new(ET2));
    <a id="L233"></a>if dec.state.err != err {
        <a id="L234"></a>t.Error(&#34;expected error&#34;, err, &#34;got&#34;, dec.state.err)
    <a id="L235"></a>}
<a id="L236"></a>}

<a id="L238"></a><span class="comment">// Check that we survive bad data.</span>
<a id="L239"></a>func TestBadData(t *testing.T) {
    <a id="L240"></a>corruptDataCheck(&#34;&#34;, os.EOF, t);
    <a id="L241"></a>corruptDataCheck(&#34;\x7Fhi&#34;, io.ErrUnexpectedEOF, t);
    <a id="L242"></a>corruptDataCheck(&#34;\x03now is the time for all good men&#34;, errBadType, t);
<a id="L243"></a>}

<a id="L245"></a><span class="comment">// Types not supported by the Encoder (only structs work at the top level).</span>
<a id="L246"></a><span class="comment">// Basic types work implicitly.</span>
<a id="L247"></a>var unsupportedValues = []interface{}{
    <a id="L248"></a>3,
    <a id="L249"></a>&#34;hi&#34;,
    <a id="L250"></a>7.2,
    <a id="L251"></a>[]int{1, 2, 3},
    <a id="L252"></a>[3]int{1, 2, 3},
    <a id="L253"></a>make(chan int),
    <a id="L254"></a>func(a int) bool { return true },
    <a id="L255"></a>make(map[string]int),
    <a id="L256"></a>new(interface{}),
<a id="L257"></a>}

<a id="L259"></a>func TestUnsupported(t *testing.T) {
    <a id="L260"></a>var b bytes.Buffer;
    <a id="L261"></a>enc := NewEncoder(&amp;b);
    <a id="L262"></a>for _, v := range unsupportedValues {
        <a id="L263"></a>err := enc.Encode(v);
        <a id="L264"></a>if err == nil {
            <a id="L265"></a>t.Errorf(&#34;expected error for %T; got none&#34;, v)
        <a id="L266"></a>}
    <a id="L267"></a>}
<a id="L268"></a>}

<a id="L270"></a>func encAndDec(in, out interface{}) os.Error {
    <a id="L271"></a>b := new(bytes.Buffer);
    <a id="L272"></a>enc := NewEncoder(b);
    <a id="L273"></a>enc.Encode(in);
    <a id="L274"></a>if enc.state.err != nil {
        <a id="L275"></a>return enc.state.err
    <a id="L276"></a>}
    <a id="L277"></a>dec := NewDecoder(b);
    <a id="L278"></a>dec.Decode(out);
    <a id="L279"></a>if dec.state.err != nil {
        <a id="L280"></a>return dec.state.err
    <a id="L281"></a>}
    <a id="L282"></a>return nil;
<a id="L283"></a>}

<a id="L285"></a>func TestTypeToPtrType(t *testing.T) {
    <a id="L286"></a><span class="comment">// Encode a T, decode a *T</span>
    <a id="L287"></a>type Type0 struct {
        <a id="L288"></a>a int;
    <a id="L289"></a>}
    <a id="L290"></a>t0 := Type0{7};
    <a id="L291"></a>t0p := (*Type0)(nil);
    <a id="L292"></a>if err := encAndDec(t0, t0p); err != nil {
        <a id="L293"></a>t.Error(err)
    <a id="L294"></a>}
<a id="L295"></a>}

<a id="L297"></a>func TestPtrTypeToType(t *testing.T) {
    <a id="L298"></a><span class="comment">// Encode a *T, decode a T</span>
    <a id="L299"></a>type Type1 struct {
        <a id="L300"></a>a uint;
    <a id="L301"></a>}
    <a id="L302"></a>t1p := &amp;Type1{17};
    <a id="L303"></a>var t1 Type1;
    <a id="L304"></a>if err := encAndDec(t1, t1p); err != nil {
        <a id="L305"></a>t.Error(err)
    <a id="L306"></a>}
<a id="L307"></a>}

<a id="L309"></a>func TestTypeToPtrPtrPtrPtrType(t *testing.T) {
    <a id="L310"></a><span class="comment">// Encode a *T, decode a T</span>
    <a id="L311"></a>type Type2 struct {
        <a id="L312"></a>a ****float;
    <a id="L313"></a>}
    <a id="L314"></a>t2 := Type2{};
    <a id="L315"></a>t2.a = new(***float);
    <a id="L316"></a>*t2.a = new(**float);
    <a id="L317"></a>**t2.a = new(*float);
    <a id="L318"></a>***t2.a = new(float);
    <a id="L319"></a>****t2.a = 27.4;
    <a id="L320"></a>t2pppp := new(***Type2);
    <a id="L321"></a>if err := encAndDec(t2, t2pppp); err != nil {
        <a id="L322"></a>t.Error(err)
    <a id="L323"></a>}
    <a id="L324"></a>if ****(****t2pppp).a != ****t2.a {
        <a id="L325"></a>t.Errorf(&#34;wrong value after decode: %g not %g&#34;, ****(****t2pppp).a, ****t2.a)
    <a id="L326"></a>}
<a id="L327"></a>}
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
