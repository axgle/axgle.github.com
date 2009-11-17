<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/binary/binary.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/encoding/binary/binary.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements translation between</span>
<a id="L6"></a><span class="comment">// unsigned integer values and byte sequences.</span>
<a id="L7"></a>package binary

<a id="L9"></a>import (
    <a id="L10"></a>&#34;math&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;reflect&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// A ByteOrder specifies how to convert byte sequences into</span>
<a id="L17"></a><span class="comment">// 16-, 32-, or 64-bit unsigned integers.</span>
<a id="L18"></a>type ByteOrder interface {
    <a id="L19"></a>Uint16(b []byte) uint16;
    <a id="L20"></a>Uint32(b []byte) uint32;
    <a id="L21"></a>Uint64(b []byte) uint64;
    <a id="L22"></a>PutUint16([]byte, uint16);
    <a id="L23"></a>PutUint32([]byte, uint32);
    <a id="L24"></a>PutUint64([]byte, uint64);
    <a id="L25"></a>String() string;
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// This is byte instead of struct{} so that it can be compared,</span>
<a id="L29"></a><span class="comment">// allowing, e.g., order == binary.LittleEndian.</span>
<a id="L30"></a>type unused byte

<a id="L32"></a>var LittleEndian ByteOrder = littleEndian(0)
<a id="L33"></a>var BigEndian ByteOrder = bigEndian(0)

<a id="L35"></a>type littleEndian unused

<a id="L37"></a>func (littleEndian) Uint16(b []byte) uint16 { return uint16(b[0]) | uint16(b[1])&lt;&lt;8 }

<a id="L39"></a>func (littleEndian) PutUint16(b []byte, v uint16) {
    <a id="L40"></a>b[0] = byte(v);
    <a id="L41"></a>b[1] = byte(v &gt;&gt; 8);
<a id="L42"></a>}

<a id="L44"></a>func (littleEndian) Uint32(b []byte) uint32 {
    <a id="L45"></a>return uint32(b[0]) | uint32(b[1])&lt;&lt;8 | uint32(b[2])&lt;&lt;16 | uint32(b[3])&lt;&lt;24
<a id="L46"></a>}

<a id="L48"></a>func (littleEndian) PutUint32(b []byte, v uint32) {
    <a id="L49"></a>b[0] = byte(v);
    <a id="L50"></a>b[1] = byte(v &gt;&gt; 8);
    <a id="L51"></a>b[2] = byte(v &gt;&gt; 16);
    <a id="L52"></a>b[3] = byte(v &gt;&gt; 24);
<a id="L53"></a>}

<a id="L55"></a>func (littleEndian) Uint64(b []byte) uint64 {
    <a id="L56"></a>return uint64(b[0]) | uint64(b[1])&lt;&lt;8 | uint64(b[2])&lt;&lt;16 | uint64(b[3])&lt;&lt;24 |
        <a id="L57"></a>uint64(b[4])&lt;&lt;32 | uint64(b[5])&lt;&lt;40 | uint64(b[6])&lt;&lt;48 | uint64(b[7])&lt;&lt;56
<a id="L58"></a>}

<a id="L60"></a>func (littleEndian) PutUint64(b []byte, v uint64) {
    <a id="L61"></a>b[0] = byte(v);
    <a id="L62"></a>b[1] = byte(v &gt;&gt; 8);
    <a id="L63"></a>b[2] = byte(v &gt;&gt; 16);
    <a id="L64"></a>b[3] = byte(v &gt;&gt; 24);
    <a id="L65"></a>b[4] = byte(v &gt;&gt; 32);
    <a id="L66"></a>b[5] = byte(v &gt;&gt; 40);
    <a id="L67"></a>b[6] = byte(v &gt;&gt; 48);
    <a id="L68"></a>b[7] = byte(v &gt;&gt; 56);
<a id="L69"></a>}

<a id="L71"></a>func (littleEndian) String() string { return &#34;LittleEndian&#34; }

<a id="L73"></a>func (littleEndian) GoString() string { return &#34;binary.LittleEndian&#34; }

<a id="L75"></a>type bigEndian unused

<a id="L77"></a>func (bigEndian) Uint16(b []byte) uint16 { return uint16(b[1]) | uint16(b[0])&lt;&lt;8 }

<a id="L79"></a>func (bigEndian) PutUint16(b []byte, v uint16) {
    <a id="L80"></a>b[0] = byte(v &gt;&gt; 8);
    <a id="L81"></a>b[1] = byte(v);
<a id="L82"></a>}

<a id="L84"></a>func (bigEndian) Uint32(b []byte) uint32 {
    <a id="L85"></a>return uint32(b[3]) | uint32(b[2])&lt;&lt;8 | uint32(b[1])&lt;&lt;16 | uint32(b[0])&lt;&lt;24
<a id="L86"></a>}

<a id="L88"></a>func (bigEndian) PutUint32(b []byte, v uint32) {
    <a id="L89"></a>b[0] = byte(v &gt;&gt; 24);
    <a id="L90"></a>b[1] = byte(v &gt;&gt; 16);
    <a id="L91"></a>b[2] = byte(v &gt;&gt; 8);
    <a id="L92"></a>b[3] = byte(v);
<a id="L93"></a>}

<a id="L95"></a>func (bigEndian) Uint64(b []byte) uint64 {
    <a id="L96"></a>return uint64(b[7]) | uint64(b[6])&lt;&lt;8 | uint64(b[5])&lt;&lt;16 | uint64(b[4])&lt;&lt;24 |
        <a id="L97"></a>uint64(b[3])&lt;&lt;32 | uint64(b[2])&lt;&lt;40 | uint64(b[1])&lt;&lt;48 | uint64(b[0])&lt;&lt;56
<a id="L98"></a>}

<a id="L100"></a>func (bigEndian) PutUint64(b []byte, v uint64) {
    <a id="L101"></a>b[0] = byte(v &gt;&gt; 56);
    <a id="L102"></a>b[1] = byte(v &gt;&gt; 48);
    <a id="L103"></a>b[2] = byte(v &gt;&gt; 40);
    <a id="L104"></a>b[3] = byte(v &gt;&gt; 32);
    <a id="L105"></a>b[4] = byte(v &gt;&gt; 24);
    <a id="L106"></a>b[5] = byte(v &gt;&gt; 16);
    <a id="L107"></a>b[6] = byte(v &gt;&gt; 8);
    <a id="L108"></a>b[7] = byte(v);
<a id="L109"></a>}

<a id="L111"></a>func (bigEndian) String() string { return &#34;BigEndian&#34; }

<a id="L113"></a>func (bigEndian) GoString() string { return &#34;binary.BigEndian&#34; }

<a id="L115"></a><span class="comment">// Read reads structured binary data from r into data.</span>
<a id="L116"></a><span class="comment">// Data must be a pointer to a fixed-size value.</span>
<a id="L117"></a><span class="comment">// A fixed-size value is either a fixed-size integer</span>
<a id="L118"></a><span class="comment">// (int8, uint8, int16, uint16, ...) or an array or struct</span>
<a id="L119"></a><span class="comment">// containing only fixed-size values.  Bytes read from</span>
<a id="L120"></a><span class="comment">// r are decoded using order and written to successive</span>
<a id="L121"></a><span class="comment">// fields of the data.</span>
<a id="L122"></a>func Read(r io.Reader, order ByteOrder, data interface{}) os.Error {
    <a id="L123"></a>v := reflect.NewValue(data).(*reflect.PtrValue).Elem();
    <a id="L124"></a>size := sizeof(v.Type());
    <a id="L125"></a>if size &lt; 0 {
        <a id="L126"></a>return os.NewError(&#34;binary.Read: invalid type &#34; + v.Type().String())
    <a id="L127"></a>}
    <a id="L128"></a>d := &amp;decoder{order: order, buf: make([]byte, size)};
    <a id="L129"></a>if _, err := io.ReadFull(r, d.buf); err != nil {
        <a id="L130"></a>return err
    <a id="L131"></a>}
    <a id="L132"></a>d.value(v);
    <a id="L133"></a>return nil;
<a id="L134"></a>}

<a id="L136"></a>func sizeof(t reflect.Type) int {
    <a id="L137"></a>switch t := t.(type) {
    <a id="L138"></a>case *reflect.ArrayType:
        <a id="L139"></a>n := sizeof(t.Elem());
        <a id="L140"></a>if n &lt; 0 {
            <a id="L141"></a>return -1
        <a id="L142"></a>}
        <a id="L143"></a>return t.Len() * n;

    <a id="L145"></a>case *reflect.StructType:
        <a id="L146"></a>sum := 0;
        <a id="L147"></a>for i, n := 0, t.NumField(); i &lt; n; i++ {
            <a id="L148"></a>s := sizeof(t.Field(i).Type);
            <a id="L149"></a>if s &lt; 0 {
                <a id="L150"></a>return -1
            <a id="L151"></a>}
            <a id="L152"></a>sum += s;
        <a id="L153"></a>}
        <a id="L154"></a>return sum;

    <a id="L156"></a>case *reflect.Uint8Type:
        <a id="L157"></a>return 1
    <a id="L158"></a>case *reflect.Uint16Type:
        <a id="L159"></a>return 2
    <a id="L160"></a>case *reflect.Uint32Type:
        <a id="L161"></a>return 4
    <a id="L162"></a>case *reflect.Uint64Type:
        <a id="L163"></a>return 8
    <a id="L164"></a>case *reflect.Int8Type:
        <a id="L165"></a>return 1
    <a id="L166"></a>case *reflect.Int16Type:
        <a id="L167"></a>return 2
    <a id="L168"></a>case *reflect.Int32Type:
        <a id="L169"></a>return 4
    <a id="L170"></a>case *reflect.Int64Type:
        <a id="L171"></a>return 8
    <a id="L172"></a>case *reflect.Float32Type:
        <a id="L173"></a>return 4
    <a id="L174"></a>case *reflect.Float64Type:
        <a id="L175"></a>return 8
    <a id="L176"></a>}
    <a id="L177"></a>return -1;
<a id="L178"></a>}

<a id="L180"></a>type decoder struct {
    <a id="L181"></a>order ByteOrder;
    <a id="L182"></a>buf   []byte;
<a id="L183"></a>}

<a id="L185"></a>func (d *decoder) uint8() uint8 {
    <a id="L186"></a>x := d.buf[0];
    <a id="L187"></a>d.buf = d.buf[1:len(d.buf)];
    <a id="L188"></a>return x;
<a id="L189"></a>}

<a id="L191"></a>func (d *decoder) uint16() uint16 {
    <a id="L192"></a>x := d.order.Uint16(d.buf[0:2]);
    <a id="L193"></a>d.buf = d.buf[2:len(d.buf)];
    <a id="L194"></a>return x;
<a id="L195"></a>}

<a id="L197"></a>func (d *decoder) uint32() uint32 {
    <a id="L198"></a>x := d.order.Uint32(d.buf[0:4]);
    <a id="L199"></a>d.buf = d.buf[4:len(d.buf)];
    <a id="L200"></a>return x;
<a id="L201"></a>}

<a id="L203"></a>func (d *decoder) uint64() uint64 {
    <a id="L204"></a>x := d.order.Uint64(d.buf[0:8]);
    <a id="L205"></a>d.buf = d.buf[8:len(d.buf)];
    <a id="L206"></a>return x;
<a id="L207"></a>}

<a id="L209"></a>func (d *decoder) int8() int8 { return int8(d.uint8()) }

<a id="L211"></a>func (d *decoder) int16() int16 { return int16(d.uint16()) }

<a id="L213"></a>func (d *decoder) int32() int32 { return int32(d.uint32()) }

<a id="L215"></a>func (d *decoder) int64() int64 { return int64(d.uint64()) }

<a id="L217"></a>func (d *decoder) value(v reflect.Value) {
    <a id="L218"></a>switch v := v.(type) {
    <a id="L219"></a>case *reflect.ArrayValue:
        <a id="L220"></a>l := v.Len();
        <a id="L221"></a>for i := 0; i &lt; l; i++ {
            <a id="L222"></a>d.value(v.Elem(i))
        <a id="L223"></a>}
    <a id="L224"></a>case *reflect.StructValue:
        <a id="L225"></a>l := v.NumField();
        <a id="L226"></a>for i := 0; i &lt; l; i++ {
            <a id="L227"></a>d.value(v.Field(i))
        <a id="L228"></a>}

    <a id="L230"></a>case *reflect.Uint8Value:
        <a id="L231"></a>v.Set(d.uint8())
    <a id="L232"></a>case *reflect.Uint16Value:
        <a id="L233"></a>v.Set(d.uint16())
    <a id="L234"></a>case *reflect.Uint32Value:
        <a id="L235"></a>v.Set(d.uint32())
    <a id="L236"></a>case *reflect.Uint64Value:
        <a id="L237"></a>v.Set(d.uint64())
    <a id="L238"></a>case *reflect.Int8Value:
        <a id="L239"></a>v.Set(d.int8())
    <a id="L240"></a>case *reflect.Int16Value:
        <a id="L241"></a>v.Set(d.int16())
    <a id="L242"></a>case *reflect.Int32Value:
        <a id="L243"></a>v.Set(d.int32())
    <a id="L244"></a>case *reflect.Int64Value:
        <a id="L245"></a>v.Set(d.int64())
    <a id="L246"></a>case *reflect.Float32Value:
        <a id="L247"></a>v.Set(math.Float32frombits(d.uint32()))
    <a id="L248"></a>case *reflect.Float64Value:
        <a id="L249"></a>v.Set(math.Float64frombits(d.uint64()))
    <a id="L250"></a>}
<a id="L251"></a>}
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
