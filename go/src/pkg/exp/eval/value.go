<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/value.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/value.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package eval

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bignum&#34;;
    <a id="L9"></a>&#34;fmt&#34;;
<a id="L10"></a>)

<a id="L12"></a>type Value interface {
    <a id="L13"></a>String() string;
    <a id="L14"></a><span class="comment">// Assign copies another value into this one.  It should</span>
    <a id="L15"></a><span class="comment">// assume that the other value satisfies the same specific</span>
    <a id="L16"></a><span class="comment">// value interface (BoolValue, etc.), but must not assume</span>
    <a id="L17"></a><span class="comment">// anything about its specific type.</span>
    <a id="L18"></a>Assign(t *Thread, o Value);
<a id="L19"></a>}

<a id="L21"></a>type BoolValue interface {
    <a id="L22"></a>Value;
    <a id="L23"></a>Get(*Thread) bool;
    <a id="L24"></a>Set(*Thread, bool);
<a id="L25"></a>}

<a id="L27"></a>type UintValue interface {
    <a id="L28"></a>Value;
    <a id="L29"></a>Get(*Thread) uint64;
    <a id="L30"></a>Set(*Thread, uint64);
<a id="L31"></a>}

<a id="L33"></a>type IntValue interface {
    <a id="L34"></a>Value;
    <a id="L35"></a>Get(*Thread) int64;
    <a id="L36"></a>Set(*Thread, int64);
<a id="L37"></a>}

<a id="L39"></a><span class="comment">// TODO(austin) IdealIntValue and IdealFloatValue should not exist</span>
<a id="L40"></a><span class="comment">// because ideals are not l-values.</span>
<a id="L41"></a>type IdealIntValue interface {
    <a id="L42"></a>Value;
    <a id="L43"></a>Get() *bignum.Integer;
<a id="L44"></a>}

<a id="L46"></a>type FloatValue interface {
    <a id="L47"></a>Value;
    <a id="L48"></a>Get(*Thread) float64;
    <a id="L49"></a>Set(*Thread, float64);
<a id="L50"></a>}

<a id="L52"></a>type IdealFloatValue interface {
    <a id="L53"></a>Value;
    <a id="L54"></a>Get() *bignum.Rational;
<a id="L55"></a>}

<a id="L57"></a>type StringValue interface {
    <a id="L58"></a>Value;
    <a id="L59"></a>Get(*Thread) string;
    <a id="L60"></a>Set(*Thread, string);
<a id="L61"></a>}

<a id="L63"></a>type ArrayValue interface {
    <a id="L64"></a>Value;
    <a id="L65"></a><span class="comment">// TODO(austin) Get() is here for uniformity, but is</span>
    <a id="L66"></a><span class="comment">// completely useless.  If a lot of other types have similarly</span>
    <a id="L67"></a><span class="comment">// useless Get methods, just special-case these uses.</span>
    <a id="L68"></a>Get(*Thread) ArrayValue;
    <a id="L69"></a>Elem(*Thread, int64) Value;
    <a id="L70"></a><span class="comment">// Sub returns an ArrayValue backed by the same array that</span>
    <a id="L71"></a><span class="comment">// starts from element i and has length len.</span>
    <a id="L72"></a>Sub(i int64, len int64) ArrayValue;
<a id="L73"></a>}

<a id="L75"></a>type StructValue interface {
    <a id="L76"></a>Value;
    <a id="L77"></a><span class="comment">// TODO(austin) This is another useless Get()</span>
    <a id="L78"></a>Get(*Thread) StructValue;
    <a id="L79"></a>Field(*Thread, int) Value;
<a id="L80"></a>}

<a id="L82"></a>type PtrValue interface {
    <a id="L83"></a>Value;
    <a id="L84"></a>Get(*Thread) Value;
    <a id="L85"></a>Set(*Thread, Value);
<a id="L86"></a>}

<a id="L88"></a>type Func interface {
    <a id="L89"></a>NewFrame() *Frame;
    <a id="L90"></a>Call(*Thread);
<a id="L91"></a>}

<a id="L93"></a>type FuncValue interface {
    <a id="L94"></a>Value;
    <a id="L95"></a>Get(*Thread) Func;
    <a id="L96"></a>Set(*Thread, Func);
<a id="L97"></a>}

<a id="L99"></a>type Interface struct {
    <a id="L100"></a>Type  Type;
    <a id="L101"></a>Value Value;
<a id="L102"></a>}

<a id="L104"></a>type InterfaceValue interface {
    <a id="L105"></a>Value;
    <a id="L106"></a>Get(*Thread) Interface;
    <a id="L107"></a>Set(*Thread, Interface);
<a id="L108"></a>}

<a id="L110"></a>type Slice struct {
    <a id="L111"></a>Base     ArrayValue;
    <a id="L112"></a>Len, Cap int64;
<a id="L113"></a>}

<a id="L115"></a>type SliceValue interface {
    <a id="L116"></a>Value;
    <a id="L117"></a>Get(*Thread) Slice;
    <a id="L118"></a>Set(*Thread, Slice);
<a id="L119"></a>}

<a id="L121"></a>type Map interface {
    <a id="L122"></a>Len(*Thread) int64;
    <a id="L123"></a><span class="comment">// Retrieve an element from the map, returning nil if it does</span>
    <a id="L124"></a><span class="comment">// not exist.</span>
    <a id="L125"></a>Elem(t *Thread, key interface{}) Value;
    <a id="L126"></a><span class="comment">// Set an entry in the map.  If val is nil, delete the entry.</span>
    <a id="L127"></a>SetElem(t *Thread, key interface{}, val Value);
    <a id="L128"></a><span class="comment">// TODO(austin)  Perhaps there should be an iterator interface instead.</span>
    <a id="L129"></a>Iter(func(key interface{}, val Value) bool);
<a id="L130"></a>}

<a id="L132"></a>type MapValue interface {
    <a id="L133"></a>Value;
    <a id="L134"></a>Get(*Thread) Map;
    <a id="L135"></a>Set(*Thread, Map);
<a id="L136"></a>}

<a id="L138"></a><span class="comment">/*</span>
<a id="L139"></a><span class="comment"> * Bool</span>
<a id="L140"></a><span class="comment"> */</span>

<a id="L142"></a>type boolV bool

<a id="L144"></a>func (v *boolV) String() string { return fmt.Sprint(*v) }

<a id="L146"></a>func (v *boolV) Assign(t *Thread, o Value) { *v = boolV(o.(BoolValue).Get(t)) }

<a id="L148"></a>func (v *boolV) Get(*Thread) bool { return bool(*v) }

<a id="L150"></a>func (v *boolV) Set(t *Thread, x bool) { *v = boolV(x) }

<a id="L152"></a><span class="comment">/*</span>
<a id="L153"></a><span class="comment"> * Uint</span>
<a id="L154"></a><span class="comment"> */</span>

<a id="L156"></a>type uint8V uint8

<a id="L158"></a>func (v *uint8V) String() string { return fmt.Sprint(*v) }

<a id="L160"></a>func (v *uint8V) Assign(t *Thread, o Value) { *v = uint8V(o.(UintValue).Get(t)) }

<a id="L162"></a>func (v *uint8V) Get(*Thread) uint64 { return uint64(*v) }

<a id="L164"></a>func (v *uint8V) Set(t *Thread, x uint64) { *v = uint8V(x) }

<a id="L166"></a>type uint16V uint16

<a id="L168"></a>func (v *uint16V) String() string { return fmt.Sprint(*v) }

<a id="L170"></a>func (v *uint16V) Assign(t *Thread, o Value) { *v = uint16V(o.(UintValue).Get(t)) }

<a id="L172"></a>func (v *uint16V) Get(*Thread) uint64 { return uint64(*v) }

<a id="L174"></a>func (v *uint16V) Set(t *Thread, x uint64) { *v = uint16V(x) }

<a id="L176"></a>type uint32V uint32

<a id="L178"></a>func (v *uint32V) String() string { return fmt.Sprint(*v) }

<a id="L180"></a>func (v *uint32V) Assign(t *Thread, o Value) { *v = uint32V(o.(UintValue).Get(t)) }

<a id="L182"></a>func (v *uint32V) Get(*Thread) uint64 { return uint64(*v) }

<a id="L184"></a>func (v *uint32V) Set(t *Thread, x uint64) { *v = uint32V(x) }

<a id="L186"></a>type uint64V uint64

<a id="L188"></a>func (v *uint64V) String() string { return fmt.Sprint(*v) }

<a id="L190"></a>func (v *uint64V) Assign(t *Thread, o Value) { *v = uint64V(o.(UintValue).Get(t)) }

<a id="L192"></a>func (v *uint64V) Get(*Thread) uint64 { return uint64(*v) }

<a id="L194"></a>func (v *uint64V) Set(t *Thread, x uint64) { *v = uint64V(x) }

<a id="L196"></a>type uintV uint

<a id="L198"></a>func (v *uintV) String() string { return fmt.Sprint(*v) }

<a id="L200"></a>func (v *uintV) Assign(t *Thread, o Value) { *v = uintV(o.(UintValue).Get(t)) }

<a id="L202"></a>func (v *uintV) Get(*Thread) uint64 { return uint64(*v) }

<a id="L204"></a>func (v *uintV) Set(t *Thread, x uint64) { *v = uintV(x) }

<a id="L206"></a>type uintptrV uintptr

<a id="L208"></a>func (v *uintptrV) String() string { return fmt.Sprint(*v) }

<a id="L210"></a>func (v *uintptrV) Assign(t *Thread, o Value) { *v = uintptrV(o.(UintValue).Get(t)) }

<a id="L212"></a>func (v *uintptrV) Get(*Thread) uint64 { return uint64(*v) }

<a id="L214"></a>func (v *uintptrV) Set(t *Thread, x uint64) { *v = uintptrV(x) }

<a id="L216"></a><span class="comment">/*</span>
<a id="L217"></a><span class="comment"> * Int</span>
<a id="L218"></a><span class="comment"> */</span>

<a id="L220"></a>type int8V int8

<a id="L222"></a>func (v *int8V) String() string { return fmt.Sprint(*v) }

<a id="L224"></a>func (v *int8V) Assign(t *Thread, o Value) { *v = int8V(o.(IntValue).Get(t)) }

<a id="L226"></a>func (v *int8V) Get(*Thread) int64 { return int64(*v) }

<a id="L228"></a>func (v *int8V) Set(t *Thread, x int64) { *v = int8V(x) }

<a id="L230"></a>type int16V int16

<a id="L232"></a>func (v *int16V) String() string { return fmt.Sprint(*v) }

<a id="L234"></a>func (v *int16V) Assign(t *Thread, o Value) { *v = int16V(o.(IntValue).Get(t)) }

<a id="L236"></a>func (v *int16V) Get(*Thread) int64 { return int64(*v) }

<a id="L238"></a>func (v *int16V) Set(t *Thread, x int64) { *v = int16V(x) }

<a id="L240"></a>type int32V int32

<a id="L242"></a>func (v *int32V) String() string { return fmt.Sprint(*v) }

<a id="L244"></a>func (v *int32V) Assign(t *Thread, o Value) { *v = int32V(o.(IntValue).Get(t)) }

<a id="L246"></a>func (v *int32V) Get(*Thread) int64 { return int64(*v) }

<a id="L248"></a>func (v *int32V) Set(t *Thread, x int64) { *v = int32V(x) }

<a id="L250"></a>type int64V int64

<a id="L252"></a>func (v *int64V) String() string { return fmt.Sprint(*v) }

<a id="L254"></a>func (v *int64V) Assign(t *Thread, o Value) { *v = int64V(o.(IntValue).Get(t)) }

<a id="L256"></a>func (v *int64V) Get(*Thread) int64 { return int64(*v) }

<a id="L258"></a>func (v *int64V) Set(t *Thread, x int64) { *v = int64V(x) }

<a id="L260"></a>type intV int

<a id="L262"></a>func (v *intV) String() string { return fmt.Sprint(*v) }

<a id="L264"></a>func (v *intV) Assign(t *Thread, o Value) { *v = intV(o.(IntValue).Get(t)) }

<a id="L266"></a>func (v *intV) Get(*Thread) int64 { return int64(*v) }

<a id="L268"></a>func (v *intV) Set(t *Thread, x int64) { *v = intV(x) }

<a id="L270"></a><span class="comment">/*</span>
<a id="L271"></a><span class="comment"> * Ideal int</span>
<a id="L272"></a><span class="comment"> */</span>

<a id="L274"></a>type idealIntV struct {
    <a id="L275"></a>V *bignum.Integer;
<a id="L276"></a>}

<a id="L278"></a>func (v *idealIntV) String() string { return v.V.String() }

<a id="L280"></a>func (v *idealIntV) Assign(t *Thread, o Value) {
    <a id="L281"></a>v.V = o.(IdealIntValue).Get()
<a id="L282"></a>}

<a id="L284"></a>func (v *idealIntV) Get() *bignum.Integer { return v.V }

<a id="L286"></a><span class="comment">/*</span>
<a id="L287"></a><span class="comment"> * Float</span>
<a id="L288"></a><span class="comment"> */</span>

<a id="L290"></a>type float32V float32

<a id="L292"></a>func (v *float32V) String() string { return fmt.Sprint(*v) }

<a id="L294"></a>func (v *float32V) Assign(t *Thread, o Value) { *v = float32V(o.(FloatValue).Get(t)) }

<a id="L296"></a>func (v *float32V) Get(*Thread) float64 { return float64(*v) }

<a id="L298"></a>func (v *float32V) Set(t *Thread, x float64) { *v = float32V(x) }

<a id="L300"></a>type float64V float64

<a id="L302"></a>func (v *float64V) String() string { return fmt.Sprint(*v) }

<a id="L304"></a>func (v *float64V) Assign(t *Thread, o Value) { *v = float64V(o.(FloatValue).Get(t)) }

<a id="L306"></a>func (v *float64V) Get(*Thread) float64 { return float64(*v) }

<a id="L308"></a>func (v *float64V) Set(t *Thread, x float64) { *v = float64V(x) }

<a id="L310"></a>type floatV float

<a id="L312"></a>func (v *floatV) String() string { return fmt.Sprint(*v) }

<a id="L314"></a>func (v *floatV) Assign(t *Thread, o Value) { *v = floatV(o.(FloatValue).Get(t)) }

<a id="L316"></a>func (v *floatV) Get(*Thread) float64 { return float64(*v) }

<a id="L318"></a>func (v *floatV) Set(t *Thread, x float64) { *v = floatV(x) }

<a id="L320"></a><span class="comment">/*</span>
<a id="L321"></a><span class="comment"> * Ideal float</span>
<a id="L322"></a><span class="comment"> */</span>

<a id="L324"></a>type idealFloatV struct {
    <a id="L325"></a>V *bignum.Rational;
<a id="L326"></a>}

<a id="L328"></a>func (v *idealFloatV) String() string { return ratToString(v.V) }

<a id="L330"></a>func (v *idealFloatV) Assign(t *Thread, o Value) {
    <a id="L331"></a>v.V = o.(IdealFloatValue).Get()
<a id="L332"></a>}

<a id="L334"></a>func (v *idealFloatV) Get() *bignum.Rational { return v.V }

<a id="L336"></a><span class="comment">/*</span>
<a id="L337"></a><span class="comment"> * String</span>
<a id="L338"></a><span class="comment"> */</span>

<a id="L340"></a>type stringV string

<a id="L342"></a>func (v *stringV) String() string { return fmt.Sprint(*v) }

<a id="L344"></a>func (v *stringV) Assign(t *Thread, o Value) { *v = stringV(o.(StringValue).Get(t)) }

<a id="L346"></a>func (v *stringV) Get(*Thread) string { return string(*v) }

<a id="L348"></a>func (v *stringV) Set(t *Thread, x string) { *v = stringV(x) }

<a id="L350"></a><span class="comment">/*</span>
<a id="L351"></a><span class="comment"> * Array</span>
<a id="L352"></a><span class="comment"> */</span>

<a id="L354"></a>type arrayV []Value

<a id="L356"></a>func (v *arrayV) String() string {
    <a id="L357"></a>res := &#34;{&#34;;
    <a id="L358"></a>for i, e := range *v {
        <a id="L359"></a>if i &gt; 0 {
            <a id="L360"></a>res += &#34;, &#34;
        <a id="L361"></a>}
        <a id="L362"></a>res += e.String();
    <a id="L363"></a>}
    <a id="L364"></a>return res + &#34;}&#34;;
<a id="L365"></a>}

<a id="L367"></a>func (v *arrayV) Assign(t *Thread, o Value) {
    <a id="L368"></a>oa := o.(ArrayValue);
    <a id="L369"></a>l := int64(len(*v));
    <a id="L370"></a>for i := int64(0); i &lt; l; i++ {
        <a id="L371"></a>(*v)[i].Assign(t, oa.Elem(t, i))
    <a id="L372"></a>}
<a id="L373"></a>}

<a id="L375"></a>func (v *arrayV) Get(*Thread) ArrayValue { return v }

<a id="L377"></a>func (v *arrayV) Elem(t *Thread, i int64) Value {
    <a id="L378"></a>return (*v)[i]
<a id="L379"></a>}

<a id="L381"></a>func (v *arrayV) Sub(i int64, len int64) ArrayValue {
    <a id="L382"></a>res := (*v)[i : i+len];
    <a id="L383"></a>return &amp;res;
<a id="L384"></a>}

<a id="L386"></a><span class="comment">/*</span>
<a id="L387"></a><span class="comment"> * Struct</span>
<a id="L388"></a><span class="comment"> */</span>

<a id="L390"></a>type structV []Value

<a id="L392"></a><span class="comment">// TODO(austin) Should these methods (and arrayV&#39;s) be on structV</span>
<a id="L393"></a><span class="comment">// instead of *structV?</span>
<a id="L394"></a>func (v *structV) String() string {
    <a id="L395"></a>res := &#34;{&#34;;
    <a id="L396"></a>for i, v := range *v {
        <a id="L397"></a>if i &gt; 0 {
            <a id="L398"></a>res += &#34;, &#34;
        <a id="L399"></a>}
        <a id="L400"></a>res += v.String();
    <a id="L401"></a>}
    <a id="L402"></a>return res + &#34;}&#34;;
<a id="L403"></a>}

<a id="L405"></a>func (v *structV) Assign(t *Thread, o Value) {
    <a id="L406"></a>oa := o.(StructValue);
    <a id="L407"></a>l := len(*v);
    <a id="L408"></a>for i := 0; i &lt; l; i++ {
        <a id="L409"></a>(*v)[i].Assign(t, oa.Field(t, i))
    <a id="L410"></a>}
<a id="L411"></a>}

<a id="L413"></a>func (v *structV) Get(*Thread) StructValue { return v }

<a id="L415"></a>func (v *structV) Field(t *Thread, i int) Value {
    <a id="L416"></a>return (*v)[i]
<a id="L417"></a>}

<a id="L419"></a><span class="comment">/*</span>
<a id="L420"></a><span class="comment"> * Pointer</span>
<a id="L421"></a><span class="comment"> */</span>

<a id="L423"></a>type ptrV struct {
    <a id="L424"></a><span class="comment">// nil if the pointer is nil</span>
    <a id="L425"></a>target Value;
<a id="L426"></a>}

<a id="L428"></a>func (v *ptrV) String() string {
    <a id="L429"></a>if v.target == nil {
        <a id="L430"></a>return &#34;&lt;nil&gt;&#34;
    <a id="L431"></a>}
    <a id="L432"></a>return &#34;&amp;&#34; + v.target.String();
<a id="L433"></a>}

<a id="L435"></a>func (v *ptrV) Assign(t *Thread, o Value) { v.target = o.(PtrValue).Get(t) }

<a id="L437"></a>func (v *ptrV) Get(*Thread) Value { return v.target }

<a id="L439"></a>func (v *ptrV) Set(t *Thread, x Value) { v.target = x }

<a id="L441"></a><span class="comment">/*</span>
<a id="L442"></a><span class="comment"> * Functions</span>
<a id="L443"></a><span class="comment"> */</span>

<a id="L445"></a>type funcV struct {
    <a id="L446"></a>target Func;
<a id="L447"></a>}

<a id="L449"></a>func (v *funcV) String() string {
    <a id="L450"></a><span class="comment">// TODO(austin) Rob wants to see the definition</span>
    <a id="L451"></a>return &#34;func {...}&#34;
<a id="L452"></a>}

<a id="L454"></a>func (v *funcV) Assign(t *Thread, o Value) { v.target = o.(FuncValue).Get(t) }

<a id="L456"></a>func (v *funcV) Get(*Thread) Func { return v.target }

<a id="L458"></a>func (v *funcV) Set(t *Thread, x Func) { v.target = x }

<a id="L460"></a><span class="comment">/*</span>
<a id="L461"></a><span class="comment"> * Interfaces</span>
<a id="L462"></a><span class="comment"> */</span>

<a id="L464"></a>type interfaceV struct {
    <a id="L465"></a>Interface;
<a id="L466"></a>}

<a id="L468"></a>func (v *interfaceV) String() string {
    <a id="L469"></a>if v.Type == nil || v.Value == nil {
        <a id="L470"></a>return &#34;&lt;nil&gt;&#34;
    <a id="L471"></a>}
    <a id="L472"></a>return v.Value.String();
<a id="L473"></a>}

<a id="L475"></a>func (v *interfaceV) Assign(t *Thread, o Value) {
    <a id="L476"></a>v.Interface = o.(InterfaceValue).Get(t)
<a id="L477"></a>}

<a id="L479"></a>func (v *interfaceV) Get(*Thread) Interface { return v.Interface }

<a id="L481"></a>func (v *interfaceV) Set(t *Thread, x Interface) {
    <a id="L482"></a>v.Interface = x
<a id="L483"></a>}

<a id="L485"></a><span class="comment">/*</span>
<a id="L486"></a><span class="comment"> * Slices</span>
<a id="L487"></a><span class="comment"> */</span>

<a id="L489"></a>type sliceV struct {
    <a id="L490"></a>Slice;
<a id="L491"></a>}

<a id="L493"></a>func (v *sliceV) String() string {
    <a id="L494"></a>if v.Base == nil {
        <a id="L495"></a>return &#34;&lt;nil&gt;&#34;
    <a id="L496"></a>}
    <a id="L497"></a>return v.Base.Sub(0, v.Len).String();
<a id="L498"></a>}

<a id="L500"></a>func (v *sliceV) Assign(t *Thread, o Value) { v.Slice = o.(SliceValue).Get(t) }

<a id="L502"></a>func (v *sliceV) Get(*Thread) Slice { return v.Slice }

<a id="L504"></a>func (v *sliceV) Set(t *Thread, x Slice) { v.Slice = x }

<a id="L506"></a><span class="comment">/*</span>
<a id="L507"></a><span class="comment"> * Maps</span>
<a id="L508"></a><span class="comment"> */</span>

<a id="L510"></a>type mapV struct {
    <a id="L511"></a>target Map;
<a id="L512"></a>}

<a id="L514"></a>func (v *mapV) String() string {
    <a id="L515"></a>if v.target == nil {
        <a id="L516"></a>return &#34;&lt;nil&gt;&#34;
    <a id="L517"></a>}
    <a id="L518"></a>res := &#34;map[&#34;;
    <a id="L519"></a>i := 0;
    <a id="L520"></a>v.target.Iter(func(key interface{}, val Value) bool {
        <a id="L521"></a>if i &gt; 0 {
            <a id="L522"></a>res += &#34;, &#34;
        <a id="L523"></a>}
        <a id="L524"></a>i++;
        <a id="L525"></a>res += fmt.Sprint(key) + &#34;:&#34; + val.String();
        <a id="L526"></a>return true;
    <a id="L527"></a>});
    <a id="L528"></a>return res + &#34;]&#34;;
<a id="L529"></a>}

<a id="L531"></a>func (v *mapV) Assign(t *Thread, o Value) { v.target = o.(MapValue).Get(t) }

<a id="L533"></a>func (v *mapV) Get(*Thread) Map { return v.target }

<a id="L535"></a>func (v *mapV) Set(t *Thread, x Map) { v.target = x }

<a id="L537"></a>type evalMap map[interface{}]Value

<a id="L539"></a>func (m evalMap) Len(t *Thread) int64 { return int64(len(m)) }

<a id="L541"></a>func (m evalMap) Elem(t *Thread, key interface{}) Value {
    <a id="L542"></a>if v, ok := m[key]; ok {
        <a id="L543"></a>return v
    <a id="L544"></a>}
    <a id="L545"></a>return nil;
<a id="L546"></a>}

<a id="L548"></a>func (m evalMap) SetElem(t *Thread, key interface{}, val Value) {
    <a id="L549"></a>if val == nil {
        <a id="L550"></a>m[key] = nil, false
    <a id="L551"></a>} else {
        <a id="L552"></a>m[key] = val
    <a id="L553"></a>}
<a id="L554"></a>}

<a id="L556"></a>func (m evalMap) Iter(cb func(key interface{}, val Value) bool) {
    <a id="L557"></a>for k, v := range m {
        <a id="L558"></a>if !cb(k, v) {
            <a id="L559"></a>break
        <a id="L560"></a>}
    <a id="L561"></a>}
<a id="L562"></a>}

<a id="L564"></a><span class="comment">/*</span>
<a id="L565"></a><span class="comment"> * Multi-values</span>
<a id="L566"></a><span class="comment"> */</span>

<a id="L568"></a>type multiV []Value

<a id="L570"></a>func (v multiV) String() string {
    <a id="L571"></a>res := &#34;(&#34;;
    <a id="L572"></a>for i, v := range v {
        <a id="L573"></a>if i &gt; 0 {
            <a id="L574"></a>res += &#34;, &#34;
        <a id="L575"></a>}
        <a id="L576"></a>res += v.String();
    <a id="L577"></a>}
    <a id="L578"></a>return res + &#34;)&#34;;
<a id="L579"></a>}

<a id="L581"></a>func (v multiV) Assign(t *Thread, o Value) {
    <a id="L582"></a>omv := o.(multiV);
    <a id="L583"></a>for i := range v {
        <a id="L584"></a>v[i].Assign(t, omv[i])
    <a id="L585"></a>}
<a id="L586"></a>}

<a id="L588"></a><span class="comment">/*</span>
<a id="L589"></a><span class="comment"> * Universal constants</span>
<a id="L590"></a><span class="comment"> */</span>

<a id="L592"></a><span class="comment">// TODO(austin) Nothing complains if I accidentally define init with</span>
<a id="L593"></a><span class="comment">// arguments.  Is this intentional?</span>
<a id="L594"></a>func init() {
    <a id="L595"></a>s := universe;

    <a id="L597"></a>true := boolV(true);
    <a id="L598"></a>s.DefineConst(&#34;true&#34;, universePos, BoolType, &amp;true);
    <a id="L599"></a>false := boolV(false);
    <a id="L600"></a>s.DefineConst(&#34;false&#34;, universePos, BoolType, &amp;false);
<a id="L601"></a>}
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
