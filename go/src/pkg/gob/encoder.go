<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/gob/encoder.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/gob/encoder.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>
<a id="L6"></a><span class="comment">	The gob package manages streams of gobs - binary values exchanged between an</span>
<a id="L7"></a><span class="comment">	Encoder (transmitter) and a Decoder (receiver).  A typical use is transporting</span>
<a id="L8"></a><span class="comment">	arguments and results of remote procedure calls (RPCs) such as those provided by</span>
<a id="L9"></a><span class="comment">	package &#34;rpc&#34;.</span>

<a id="L11"></a><span class="comment">	A stream of gobs is self-describing.  Each data item in the stream is preceded by</span>
<a id="L12"></a><span class="comment">	a specification of its type, expressed in terms of a small set of predefined</span>
<a id="L13"></a><span class="comment">	types.  Pointers are not transmitted, but the things they point to are</span>
<a id="L14"></a><span class="comment">	transmitted; that is, the values are flattened.  Recursive types work fine, but</span>
<a id="L15"></a><span class="comment">	recursive values (data with cycles) are problematic.  This may change.</span>

<a id="L17"></a><span class="comment">	To use gobs, create an Encoder and present it with a series of data items as</span>
<a id="L18"></a><span class="comment">	values or addresses that can be dereferenced to values.  (At the moment, these</span>
<a id="L19"></a><span class="comment">	items must be structs (struct, *struct, **struct etc.), but this may change.) The</span>
<a id="L20"></a><span class="comment">	Encoder makes sure all type information is sent before it is needed.  At the</span>
<a id="L21"></a><span class="comment">	receive side, a Decoder retrieves values from the encoded stream and unpacks them</span>
<a id="L22"></a><span class="comment">	into local variables.</span>

<a id="L24"></a><span class="comment">	The source and destination values/types need not correspond exactly.  For structs,</span>
<a id="L25"></a><span class="comment">	fields (identified by name) that are in the source but absent from the receiving</span>
<a id="L26"></a><span class="comment">	variable will be ignored.  Fields that are in the receiving variable but missing</span>
<a id="L27"></a><span class="comment">	from the transmitted type or value will be ignored in the destination.  If a field</span>
<a id="L28"></a><span class="comment">	with the same name is present in both, their types must be compatible. Both the</span>
<a id="L29"></a><span class="comment">	receiver and transmitter will do all necessary indirection and dereferencing to</span>
<a id="L30"></a><span class="comment">	convert between gobs and actual Go values.  For instance, a gob type that is</span>
<a id="L31"></a><span class="comment">	schematically,</span>

<a id="L33"></a><span class="comment">		struct { a, b int }</span>

<a id="L35"></a><span class="comment">	can be sent from or received into any of these Go types:</span>

<a id="L37"></a><span class="comment">		struct { a, b int }	// the same</span>
<a id="L38"></a><span class="comment">		*struct { a, b int }	// extra indirection of the struct</span>
<a id="L39"></a><span class="comment">		struct { *a, **b int }	// extra indirection of the fields</span>
<a id="L40"></a><span class="comment">		struct { a, b int64 }	// different concrete value type; see below</span>

<a id="L42"></a><span class="comment">	It may also be received into any of these:</span>

<a id="L44"></a><span class="comment">		struct { a, b int }	// the same</span>
<a id="L45"></a><span class="comment">		struct { b, a int }	// ordering doesn&#39;t matter; matching is by name</span>
<a id="L46"></a><span class="comment">		struct { a, b, c int }	// extra field (c) ignored</span>
<a id="L47"></a><span class="comment">		struct { b int }	// missing field (a) ignored; data will be dropped</span>
<a id="L48"></a><span class="comment">		struct { b, c int }	// missing field (a) ignored; extra field (c) ignored.</span>

<a id="L50"></a><span class="comment">	Attempting to receive into these types will draw a decode error:</span>

<a id="L52"></a><span class="comment">		struct { a int; b uint }	// change of signedness for b</span>
<a id="L53"></a><span class="comment">		struct { a int; b float }	// change of type for b</span>
<a id="L54"></a><span class="comment">		struct { }			// no field names in common</span>
<a id="L55"></a><span class="comment">		struct { c, d int }		// no field names in common</span>

<a id="L57"></a><span class="comment">	Integers are transmitted two ways: arbitrary precision signed integers or</span>
<a id="L58"></a><span class="comment">	arbitrary precision unsigned integers.  There is no int8, int16 etc.</span>
<a id="L59"></a><span class="comment">	discrimination in the gob format; there are only signed and unsigned integers.  As</span>
<a id="L60"></a><span class="comment">	described below, the transmitter sends the value in a variable-length encoding;</span>
<a id="L61"></a><span class="comment">	the receiver accepts the value and stores it in the destination variable.</span>
<a id="L62"></a><span class="comment">	Floating-point numbers are always sent using IEEE-754 64-bit precision (see</span>
<a id="L63"></a><span class="comment">	below).</span>

<a id="L65"></a><span class="comment">	Signed integers may be received into any signed integer variable: int, int16, etc.;</span>
<a id="L66"></a><span class="comment">	unsigned integers may be received into any unsigned integer variable; and floating</span>
<a id="L67"></a><span class="comment">	point values may be received into any floating point variable.  However,</span>
<a id="L68"></a><span class="comment">	the destination variable must be able to represent the value or the decode</span>
<a id="L69"></a><span class="comment">	operation will fail.</span>

<a id="L71"></a><span class="comment">	Structs, arrays and slices are also supported.  Strings and arrays of bytes are</span>
<a id="L72"></a><span class="comment">	supported with a special, efficient representation (see below).</span>

<a id="L74"></a><span class="comment">	Maps are not supported yet, but they will be.  Interfaces, functions, and channels</span>
<a id="L75"></a><span class="comment">	cannot be sent in a gob.  Attempting to encode a value that contains one will</span>
<a id="L76"></a><span class="comment">	fail.</span>

<a id="L78"></a><span class="comment">	The rest of this comment documents the encoding, details that are not important</span>
<a id="L79"></a><span class="comment">	for most users.  Details are presented bottom-up.</span>

<a id="L81"></a><span class="comment">	An unsigned integer is sent one of two ways.  If it is less than 128, it is sent</span>
<a id="L82"></a><span class="comment">	as a byte with that value.  Otherwise it is sent as a minimal-length big-endian</span>
<a id="L83"></a><span class="comment">	(high byte first) byte stream holding the value, preceded by one byte holding the</span>
<a id="L84"></a><span class="comment">	byte count, negated.  Thus 0 is transmitted as (00), 7 is transmitted as (07) and</span>
<a id="L85"></a><span class="comment">	256 is transmitted as (FE 01 00).</span>

<a id="L87"></a><span class="comment">	A boolean is encoded within an unsigned integer: 0 for false, 1 for true.</span>

<a id="L89"></a><span class="comment">	A signed integer, i, is encoded within an unsigned integer, u.  Within u, bits 1</span>
<a id="L90"></a><span class="comment">	upward contain the value; bit 0 says whether they should be complemented upon</span>
<a id="L91"></a><span class="comment">	receipt.  The encode algorithm looks like this:</span>

<a id="L93"></a><span class="comment">		uint u;</span>
<a id="L94"></a><span class="comment">		if i &lt; 0 {</span>
<a id="L95"></a><span class="comment">			u = (^i &lt;&lt; 1) | 1	// complement i, bit 0 is 1</span>
<a id="L96"></a><span class="comment">		} else {</span>
<a id="L97"></a><span class="comment">			u = (i &lt;&lt; 1)	// do not complement i, bit 0 is 0</span>
<a id="L98"></a><span class="comment">		}</span>
<a id="L99"></a><span class="comment">		encodeUnsigned(u)</span>

<a id="L101"></a><span class="comment">	The low bit is therefore analogous to a sign bit, but making it the complement bit</span>
<a id="L102"></a><span class="comment">	instead guarantees that the largest negative integer is not a special case.  For</span>
<a id="L103"></a><span class="comment">	example, -129=^128=(^256&gt;&gt;1) encodes as (01 82).</span>

<a id="L105"></a><span class="comment">	Floating-point numbers are always sent as a representation of a float64 value.</span>
<a id="L106"></a><span class="comment">	That value is converted to a uint64 using math.Float64bits.  The uint64 is then</span>
<a id="L107"></a><span class="comment">	byte-reversed and sent as a regular unsigned integer.  The byte-reversal means the</span>
<a id="L108"></a><span class="comment">	exponent and high-precision part of the mantissa go first.  Since the low bits are</span>
<a id="L109"></a><span class="comment">	often zero, this can save encoding bytes.  For instance, 17.0 is encoded in only</span>
<a id="L110"></a><span class="comment">	two bytes (40 e2).</span>

<a id="L112"></a><span class="comment">	Strings and slices of bytes are sent as an unsigned count followed by that many</span>
<a id="L113"></a><span class="comment">	uninterpreted bytes of the value.</span>

<a id="L115"></a><span class="comment">	All other slices and arrays are sent as an unsigned count followed by that many</span>
<a id="L116"></a><span class="comment">	elements using the standard gob encoding for their type, recursively.</span>

<a id="L118"></a><span class="comment">	Structs are sent as a sequence of (field number, field value) pairs.  The field</span>
<a id="L119"></a><span class="comment">	value is sent using the standard gob encoding for its type, recursively.  If a</span>
<a id="L120"></a><span class="comment">	field has the zero value for its type, it is omitted from the transmission.  The</span>
<a id="L121"></a><span class="comment">	field number is defined by the type of the encoded struct: the first field of the</span>
<a id="L122"></a><span class="comment">	encoded type is field 0, the second is field 1, etc.  When encoding a value, the</span>
<a id="L123"></a><span class="comment">	field numbers are delta encoded for efficiency and the fields are always sent in</span>
<a id="L124"></a><span class="comment">	order of increasing field number; the deltas are therefore unsigned.  The</span>
<a id="L125"></a><span class="comment">	initialization for the delta encoding sets the field number to -1, so an unsigned</span>
<a id="L126"></a><span class="comment">	integer field 0 with value 7 is transmitted as unsigned delta = 1, unsigned value</span>
<a id="L127"></a><span class="comment">	= 7 or (81 87).  Finally, after all the fields have been sent a terminating mark</span>
<a id="L128"></a><span class="comment">	denotes the end of the struct.  That mark is a delta=0 value, which has</span>
<a id="L129"></a><span class="comment">	representation (80).</span>

<a id="L131"></a><span class="comment">	The representation of types is described below.  When a type is defined on a given</span>
<a id="L132"></a><span class="comment">	connection between an Encoder and Decoder, it is assigned a signed integer type</span>
<a id="L133"></a><span class="comment">	id.  When Encoder.Encode(v) is called, it makes sure there is an id assigned for</span>
<a id="L134"></a><span class="comment">	the type of v and all its elements and then it sends the pair (typeid, encoded-v)</span>
<a id="L135"></a><span class="comment">	where typeid is the type id of the encoded type of v and encoded-v is the gob</span>
<a id="L136"></a><span class="comment">	encoding of the value v.</span>

<a id="L138"></a><span class="comment">	To define a type, the encoder chooses an unused, positive type id and sends the</span>
<a id="L139"></a><span class="comment">	pair (-type id, encoded-type) where encoded-type is the gob encoding of a wireType</span>
<a id="L140"></a><span class="comment">	description, constructed from these types:</span>

<a id="L142"></a><span class="comment">		type wireType struct {</span>
<a id="L143"></a><span class="comment">			s	structType;</span>
<a id="L144"></a><span class="comment">		}</span>
<a id="L145"></a><span class="comment">		type fieldType struct {</span>
<a id="L146"></a><span class="comment">			name	string;	// the name of the field.</span>
<a id="L147"></a><span class="comment">			id	int;	// the type id of the field, which must be already defined</span>
<a id="L148"></a><span class="comment">		}</span>
<a id="L149"></a><span class="comment">		type commonType {</span>
<a id="L150"></a><span class="comment">			name	string;	// the name of the struct type</span>
<a id="L151"></a><span class="comment">			id	int;	// the id of the type, repeated for so it&#39;s inside the type</span>
<a id="L152"></a><span class="comment">		}</span>
<a id="L153"></a><span class="comment">		type structType struct {</span>
<a id="L154"></a><span class="comment">			commonType;</span>
<a id="L155"></a><span class="comment">			field	[]fieldType;	// the fields of the struct.</span>
<a id="L156"></a><span class="comment">		}</span>

<a id="L158"></a><span class="comment">	If there are nested type ids, the types for all inner type ids must be defined</span>
<a id="L159"></a><span class="comment">	before the top-level type id is used to describe an encoded-v.</span>

<a id="L161"></a><span class="comment">	For simplicity in setup, the connection is defined to understand these types a</span>
<a id="L162"></a><span class="comment">	priori, as well as the basic gob types int, uint, etc.  Their ids are:</span>

<a id="L164"></a><span class="comment">		bool		1</span>
<a id="L165"></a><span class="comment">		int		2</span>
<a id="L166"></a><span class="comment">		uint		3</span>
<a id="L167"></a><span class="comment">		float		4</span>
<a id="L168"></a><span class="comment">		[]byte		5</span>
<a id="L169"></a><span class="comment">		string		6</span>
<a id="L170"></a><span class="comment">		wireType	7</span>
<a id="L171"></a><span class="comment">		structType	8</span>
<a id="L172"></a><span class="comment">		commonType	9</span>
<a id="L173"></a><span class="comment">		fieldType	10</span>

<a id="L175"></a><span class="comment">	In summary, a gob stream looks like</span>

<a id="L177"></a><span class="comment">		((-type id, encoding of a wireType)* (type id, encoding of a value))*</span>

<a id="L179"></a><span class="comment">	where * signifies zero or more repetitions and the type id of a value must</span>
<a id="L180"></a><span class="comment">	be predefined or be defined before the value in the stream.</span>
<a id="L181"></a><span class="comment">*/</span>
<a id="L182"></a>package gob

<a id="L184"></a>import (
    <a id="L185"></a>&#34;bytes&#34;;
    <a id="L186"></a>&#34;io&#34;;
    <a id="L187"></a>&#34;os&#34;;
    <a id="L188"></a>&#34;reflect&#34;;
    <a id="L189"></a>&#34;sync&#34;;
<a id="L190"></a>)

<a id="L192"></a><span class="comment">// An Encoder manages the transmission of type and data information to the</span>
<a id="L193"></a><span class="comment">// other side of a connection.</span>
<a id="L194"></a>type Encoder struct {
    <a id="L195"></a>mutex      sync.Mutex;              <span class="comment">// each item must be sent atomically</span>
    <a id="L196"></a>w          io.Writer;               <span class="comment">// where to send the data</span>
    <a id="L197"></a>sent       map[reflect.Type]typeId; <span class="comment">// which types we&#39;ve already sent</span>
    <a id="L198"></a>state      *encoderState;           <span class="comment">// so we can encode integers, strings directly</span>
    <a id="L199"></a>countState *encoderState;           <span class="comment">// stage for writing counts</span>
    <a id="L200"></a>buf        []byte;                  <span class="comment">// for collecting the output.</span>
<a id="L201"></a>}

<a id="L203"></a><span class="comment">// NewEncoder returns a new encoder that will transmit on the io.Writer.</span>
<a id="L204"></a>func NewEncoder(w io.Writer) *Encoder {
    <a id="L205"></a>enc := new(Encoder);
    <a id="L206"></a>enc.w = w;
    <a id="L207"></a>enc.sent = make(map[reflect.Type]typeId);
    <a id="L208"></a>enc.state = new(encoderState);
    <a id="L209"></a>enc.state.b = new(bytes.Buffer); <span class="comment">// the rest isn&#39;t important; all we need is buffer and writer</span>
    <a id="L210"></a>enc.countState = new(encoderState);
    <a id="L211"></a>enc.countState.b = new(bytes.Buffer); <span class="comment">// the rest isn&#39;t important; all we need is buffer and writer</span>
    <a id="L212"></a>return enc;
<a id="L213"></a>}

<a id="L215"></a>func (enc *Encoder) badType(rt reflect.Type) {
    <a id="L216"></a>enc.state.err = os.ErrorString(&#34;gob: can&#39;t encode type &#34; + rt.String())
<a id="L217"></a>}

<a id="L219"></a><span class="comment">// Send the data item preceded by a unsigned count of its length.</span>
<a id="L220"></a>func (enc *Encoder) send() {
    <a id="L221"></a><span class="comment">// Encode the length.</span>
    <a id="L222"></a>encodeUint(enc.countState, uint64(enc.state.b.Len()));
    <a id="L223"></a><span class="comment">// Build the buffer.</span>
    <a id="L224"></a>countLen := enc.countState.b.Len();
    <a id="L225"></a>total := countLen + enc.state.b.Len();
    <a id="L226"></a>if total &gt; len(enc.buf) {
        <a id="L227"></a>enc.buf = make([]byte, total+1000) <span class="comment">// extra for growth</span>
    <a id="L228"></a>}
    <a id="L229"></a><span class="comment">// Place the length before the data.</span>
    <a id="L230"></a><span class="comment">// TODO(r): avoid the extra copy here.</span>
    <a id="L231"></a>enc.countState.b.Read(enc.buf[0:countLen]);
    <a id="L232"></a><span class="comment">// Now the data.</span>
    <a id="L233"></a>enc.state.b.Read(enc.buf[countLen:total]);
    <a id="L234"></a><span class="comment">// Write the data.</span>
    <a id="L235"></a>enc.w.Write(enc.buf[0:total]);
<a id="L236"></a>}

<a id="L238"></a>func (enc *Encoder) sendType(origt reflect.Type, topLevel bool) {
    <a id="L239"></a><span class="comment">// Drill down to the base type.</span>
    <a id="L240"></a>rt, _ := indirect(origt);

    <a id="L242"></a><span class="comment">// We only send structs - everything else is basic or an error</span>
    <a id="L243"></a>switch rt.(type) {
    <a id="L244"></a>default:
        <a id="L245"></a><span class="comment">// Basic types do not need to be described, but if this is a top-level</span>
        <a id="L246"></a><span class="comment">// type, it&#39;s a user error, at least for now.</span>
        <a id="L247"></a>if topLevel {
            <a id="L248"></a>enc.badType(rt)
        <a id="L249"></a>}
        <a id="L250"></a>return;
    <a id="L251"></a>case *reflect.StructType:
        <a id="L252"></a><span class="comment">// Structs do need to be described.</span>
        <a id="L253"></a>break
    <a id="L254"></a>case *reflect.ChanType, *reflect.FuncType, *reflect.MapType, *reflect.InterfaceType:
        <a id="L255"></a><span class="comment">// Probably a bad field in a struct.</span>
        <a id="L256"></a>enc.badType(rt);
        <a id="L257"></a>return;
    <a id="L258"></a>case *reflect.ArrayType, *reflect.SliceType:
        <a id="L259"></a><span class="comment">// Array and slice types are not sent, only their element types.</span>
        <a id="L260"></a><span class="comment">// If we see one here it&#39;s user error; probably a bad top-level value.</span>
        <a id="L261"></a>enc.badType(rt);
        <a id="L262"></a>return;
    <a id="L263"></a>}

    <a id="L265"></a><span class="comment">// Have we already sent this type?  This time we ask about the base type.</span>
    <a id="L266"></a>if _, alreadySent := enc.sent[rt]; alreadySent {
        <a id="L267"></a>return
    <a id="L268"></a>}

    <a id="L270"></a><span class="comment">// Need to send it.</span>
    <a id="L271"></a>typeLock.Lock();
    <a id="L272"></a>info, err := getTypeInfo(rt);
    <a id="L273"></a>typeLock.Unlock();
    <a id="L274"></a>if err != nil {
        <a id="L275"></a>enc.state.err = err;
        <a id="L276"></a>return;
    <a id="L277"></a>}
    <a id="L278"></a><span class="comment">// Send the pair (-id, type)</span>
    <a id="L279"></a><span class="comment">// Id:</span>
    <a id="L280"></a>encodeInt(enc.state, -int64(info.id));
    <a id="L281"></a><span class="comment">// Type:</span>
    <a id="L282"></a>encode(enc.state.b, info.wire);
    <a id="L283"></a>enc.send();

    <a id="L285"></a><span class="comment">// Remember we&#39;ve sent this type.</span>
    <a id="L286"></a>enc.sent[rt] = info.id;
    <a id="L287"></a><span class="comment">// Remember we&#39;ve sent the top-level, possibly indirect type too.</span>
    <a id="L288"></a>enc.sent[origt] = info.id;
    <a id="L289"></a><span class="comment">// Now send the inner types</span>
    <a id="L290"></a>st := rt.(*reflect.StructType);
    <a id="L291"></a>for i := 0; i &lt; st.NumField(); i++ {
        <a id="L292"></a>enc.sendType(st.Field(i).Type, false)
    <a id="L293"></a>}
    <a id="L294"></a>return;
<a id="L295"></a>}

<a id="L297"></a><span class="comment">// Encode transmits the data item represented by the empty interface value,</span>
<a id="L298"></a><span class="comment">// guaranteeing that all necessary type information has been transmitted first.</span>
<a id="L299"></a>func (enc *Encoder) Encode(e interface{}) os.Error {
    <a id="L300"></a>if enc.state.b.Len() &gt; 0 || enc.countState.b.Len() &gt; 0 {
        <a id="L301"></a>panicln(&#34;Encoder: buffer not empty&#34;)
    <a id="L302"></a>}
    <a id="L303"></a>rt, _ := indirect(reflect.Typeof(e));

    <a id="L305"></a><span class="comment">// Make sure we&#39;re single-threaded through here.</span>
    <a id="L306"></a>enc.mutex.Lock();
    <a id="L307"></a>defer enc.mutex.Unlock();

    <a id="L309"></a><span class="comment">// Make sure the type is known to the other side.</span>
    <a id="L310"></a><span class="comment">// First, have we already sent this type?</span>
    <a id="L311"></a>if _, alreadySent := enc.sent[rt]; !alreadySent {
        <a id="L312"></a><span class="comment">// No, so send it.</span>
        <a id="L313"></a>enc.sendType(rt, true);
        <a id="L314"></a>if enc.state.err != nil {
            <a id="L315"></a>enc.state.b.Reset();
            <a id="L316"></a>enc.countState.b.Reset();
            <a id="L317"></a>return enc.state.err;
        <a id="L318"></a>}
    <a id="L319"></a>}

    <a id="L321"></a><span class="comment">// Identify the type of this top-level value.</span>
    <a id="L322"></a>encodeInt(enc.state, int64(enc.sent[rt]));

    <a id="L324"></a><span class="comment">// Encode the object.</span>
    <a id="L325"></a>encode(enc.state.b, e);
    <a id="L326"></a>enc.send();

    <a id="L328"></a>return enc.state.err;
<a id="L329"></a>}
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
