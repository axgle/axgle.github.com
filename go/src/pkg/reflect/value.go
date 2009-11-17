<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/reflect/value.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/reflect/value.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package reflect

<a id="L7"></a>import (
    <a id="L8"></a>&#34;runtime&#34;;
    <a id="L9"></a>&#34;unsafe&#34;;
<a id="L10"></a>)

<a id="L12"></a>const ptrSize = uintptr(unsafe.Sizeof((*byte)(nil)))
<a id="L13"></a>const cannotSet = &#34;cannot set value obtained via unexported struct field&#34;

<a id="L15"></a>type addr unsafe.Pointer

<a id="L17"></a><span class="comment">// TODO: This will have to go away when</span>
<a id="L18"></a><span class="comment">// the new gc goes in.</span>
<a id="L19"></a>func memmove(adst, asrc addr, n uintptr) {
    <a id="L20"></a>dst := uintptr(adst);
    <a id="L21"></a>src := uintptr(asrc);
    <a id="L22"></a>switch {
    <a id="L23"></a>case src &lt; dst &amp;&amp; src+n &gt; dst:
        <a id="L24"></a><span class="comment">// byte copy backward</span>
        <a id="L25"></a><span class="comment">// careful: i is unsigned</span>
        <a id="L26"></a>for i := n; i &gt; 0; {
            <a id="L27"></a>i--;
            <a id="L28"></a>*(*byte)(addr(dst + i)) = *(*byte)(addr(src + i));
        <a id="L29"></a>}
    <a id="L30"></a>case (n|src|dst)&amp;(ptrSize-1) != 0:
        <a id="L31"></a><span class="comment">// byte copy forward</span>
        <a id="L32"></a>for i := uintptr(0); i &lt; n; i++ {
            <a id="L33"></a>*(*byte)(addr(dst + i)) = *(*byte)(addr(src + i))
        <a id="L34"></a>}
    <a id="L35"></a>default:
        <a id="L36"></a><span class="comment">// word copy forward</span>
        <a id="L37"></a>for i := uintptr(0); i &lt; n; i += ptrSize {
            <a id="L38"></a>*(*uintptr)(addr(dst + i)) = *(*uintptr)(addr(src + i))
        <a id="L39"></a>}
    <a id="L40"></a>}
<a id="L41"></a>}

<a id="L43"></a><span class="comment">// Value is the common interface to reflection values.</span>
<a id="L44"></a><span class="comment">// The implementations of Value (e.g., ArrayValue, StructValue)</span>
<a id="L45"></a><span class="comment">// have additional type-specific methods.</span>
<a id="L46"></a>type Value interface {
    <a id="L47"></a><span class="comment">// Type returns the value&#39;s type.</span>
    <a id="L48"></a>Type() Type;

    <a id="L50"></a><span class="comment">// Interface returns the value as an interface{}.</span>
    <a id="L51"></a>Interface() interface{};

    <a id="L53"></a><span class="comment">// CanSet returns whether the value can be changed.</span>
    <a id="L54"></a><span class="comment">// Values obtained by the use of non-exported struct fields</span>
    <a id="L55"></a><span class="comment">// can be used in Get but not Set.</span>
    <a id="L56"></a><span class="comment">// If CanSet() returns false, calling the type-specific Set</span>
    <a id="L57"></a><span class="comment">// will cause a crash.</span>
    <a id="L58"></a>CanSet() bool;

    <a id="L60"></a><span class="comment">// SetValue assigns v to the value; v must have the same type as the value.</span>
    <a id="L61"></a>SetValue(v Value);

    <a id="L63"></a><span class="comment">// Addr returns a pointer to the underlying data.</span>
    <a id="L64"></a><span class="comment">// It is for advanced clients that also</span>
    <a id="L65"></a><span class="comment">// import the &#34;unsafe&#34; package.</span>
    <a id="L66"></a>Addr() uintptr;

    <a id="L68"></a><span class="comment">// Method returns a FuncValue corresponding to the value&#39;s i&#39;th method.</span>
    <a id="L69"></a><span class="comment">// The arguments to a Call on the returned FuncValue</span>
    <a id="L70"></a><span class="comment">// should not include a receiver; the FuncValue will use</span>
    <a id="L71"></a><span class="comment">// the value as the receiver.</span>
    <a id="L72"></a>Method(i int) *FuncValue;

    <a id="L74"></a>getAddr() addr;
<a id="L75"></a>}

<a id="L77"></a>type value struct {
    <a id="L78"></a>typ    Type;
    <a id="L79"></a>addr   addr;
    <a id="L80"></a>canSet bool;
<a id="L81"></a>}

<a id="L83"></a>func (v *value) Type() Type { return v.typ }

<a id="L85"></a>func (v *value) Addr() uintptr { return uintptr(v.addr) }

<a id="L87"></a>func (v *value) getAddr() addr { return v.addr }

<a id="L89"></a>func (v *value) Interface() interface{} {
    <a id="L90"></a>if typ, ok := v.typ.(*InterfaceType); ok {
        <a id="L91"></a><span class="comment">// There are two different representations of interface values,</span>
        <a id="L92"></a><span class="comment">// one if the interface type has methods and one if it doesn&#39;t.</span>
        <a id="L93"></a><span class="comment">// These two representations require different expressions</span>
        <a id="L94"></a><span class="comment">// to extract correctly.</span>
        <a id="L95"></a>if typ.NumMethod() == 0 {
            <a id="L96"></a><span class="comment">// Extract as interface value without methods.</span>
            <a id="L97"></a>return *(*interface{})(v.addr)
        <a id="L98"></a>}
        <a id="L99"></a><span class="comment">// Extract from v.addr as interface value with methods.</span>
        <a id="L100"></a>return *(*interface {
            <a id="L101"></a>m();
        <a id="L102"></a>})(v.addr);
    <a id="L103"></a>}
    <a id="L104"></a>return unsafe.Unreflect(v.typ, unsafe.Pointer(v.addr));
<a id="L105"></a>}

<a id="L107"></a>func (v *value) CanSet() bool { return v.canSet }

<a id="L109"></a><span class="comment">/*</span>
<a id="L110"></a><span class="comment"> * basic types</span>
<a id="L111"></a><span class="comment"> */</span>

<a id="L113"></a><span class="comment">// BoolValue represents a bool value.</span>
<a id="L114"></a>type BoolValue struct {
    <a id="L115"></a>value;
<a id="L116"></a>}

<a id="L118"></a><span class="comment">// Get returns the underlying bool value.</span>
<a id="L119"></a>func (v *BoolValue) Get() bool { return *(*bool)(v.addr) }

<a id="L121"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L122"></a>func (v *BoolValue) Set(x bool) {
    <a id="L123"></a>if !v.canSet {
        <a id="L124"></a>panic(cannotSet)
    <a id="L125"></a>}
    <a id="L126"></a>*(*bool)(v.addr) = x;
<a id="L127"></a>}

<a id="L129"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L130"></a>func (v *BoolValue) SetValue(x Value) { v.Set(x.(*BoolValue).Get()) }

<a id="L132"></a><span class="comment">// FloatValue represents a float value.</span>
<a id="L133"></a>type FloatValue struct {
    <a id="L134"></a>value;
<a id="L135"></a>}

<a id="L137"></a><span class="comment">// Get returns the underlying float value.</span>
<a id="L138"></a>func (v *FloatValue) Get() float { return *(*float)(v.addr) }

<a id="L140"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L141"></a>func (v *FloatValue) Set(x float) {
    <a id="L142"></a>if !v.canSet {
        <a id="L143"></a>panic(cannotSet)
    <a id="L144"></a>}
    <a id="L145"></a>*(*float)(v.addr) = x;
<a id="L146"></a>}

<a id="L148"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L149"></a>func (v *FloatValue) SetValue(x Value) { v.Set(x.(*FloatValue).Get()) }

<a id="L151"></a><span class="comment">// Float32Value represents a float32 value.</span>
<a id="L152"></a>type Float32Value struct {
    <a id="L153"></a>value;
<a id="L154"></a>}

<a id="L156"></a><span class="comment">// Get returns the underlying float32 value.</span>
<a id="L157"></a>func (v *Float32Value) Get() float32 { return *(*float32)(v.addr) }

<a id="L159"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L160"></a>func (v *Float32Value) Set(x float32) {
    <a id="L161"></a>if !v.canSet {
        <a id="L162"></a>panic(cannotSet)
    <a id="L163"></a>}
    <a id="L164"></a>*(*float32)(v.addr) = x;
<a id="L165"></a>}

<a id="L167"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L168"></a>func (v *Float32Value) SetValue(x Value) { v.Set(x.(*Float32Value).Get()) }

<a id="L170"></a><span class="comment">// Float64Value represents a float64 value.</span>
<a id="L171"></a>type Float64Value struct {
    <a id="L172"></a>value;
<a id="L173"></a>}

<a id="L175"></a><span class="comment">// Get returns the underlying float64 value.</span>
<a id="L176"></a>func (v *Float64Value) Get() float64 { return *(*float64)(v.addr) }

<a id="L178"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L179"></a>func (v *Float64Value) Set(x float64) {
    <a id="L180"></a>if !v.canSet {
        <a id="L181"></a>panic(cannotSet)
    <a id="L182"></a>}
    <a id="L183"></a>*(*float64)(v.addr) = x;
<a id="L184"></a>}

<a id="L186"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L187"></a>func (v *Float64Value) SetValue(x Value) { v.Set(x.(*Float64Value).Get()) }

<a id="L189"></a><span class="comment">// IntValue represents an int value.</span>
<a id="L190"></a>type IntValue struct {
    <a id="L191"></a>value;
<a id="L192"></a>}

<a id="L194"></a><span class="comment">// Get returns the underlying int value.</span>
<a id="L195"></a>func (v *IntValue) Get() int { return *(*int)(v.addr) }

<a id="L197"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L198"></a>func (v *IntValue) Set(x int) {
    <a id="L199"></a>if !v.canSet {
        <a id="L200"></a>panic(cannotSet)
    <a id="L201"></a>}
    <a id="L202"></a>*(*int)(v.addr) = x;
<a id="L203"></a>}

<a id="L205"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L206"></a>func (v *IntValue) SetValue(x Value) { v.Set(x.(*IntValue).Get()) }

<a id="L208"></a><span class="comment">// Int8Value represents an int8 value.</span>
<a id="L209"></a>type Int8Value struct {
    <a id="L210"></a>value;
<a id="L211"></a>}

<a id="L213"></a><span class="comment">// Get returns the underlying int8 value.</span>
<a id="L214"></a>func (v *Int8Value) Get() int8 { return *(*int8)(v.addr) }

<a id="L216"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L217"></a>func (v *Int8Value) Set(x int8) {
    <a id="L218"></a>if !v.canSet {
        <a id="L219"></a>panic(cannotSet)
    <a id="L220"></a>}
    <a id="L221"></a>*(*int8)(v.addr) = x;
<a id="L222"></a>}

<a id="L224"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L225"></a>func (v *Int8Value) SetValue(x Value) { v.Set(x.(*Int8Value).Get()) }

<a id="L227"></a><span class="comment">// Int16Value represents an int16 value.</span>
<a id="L228"></a>type Int16Value struct {
    <a id="L229"></a>value;
<a id="L230"></a>}

<a id="L232"></a><span class="comment">// Get returns the underlying int16 value.</span>
<a id="L233"></a>func (v *Int16Value) Get() int16 { return *(*int16)(v.addr) }

<a id="L235"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L236"></a>func (v *Int16Value) Set(x int16) {
    <a id="L237"></a>if !v.canSet {
        <a id="L238"></a>panic(cannotSet)
    <a id="L239"></a>}
    <a id="L240"></a>*(*int16)(v.addr) = x;
<a id="L241"></a>}

<a id="L243"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L244"></a>func (v *Int16Value) SetValue(x Value) { v.Set(x.(*Int16Value).Get()) }

<a id="L246"></a><span class="comment">// Int32Value represents an int32 value.</span>
<a id="L247"></a>type Int32Value struct {
    <a id="L248"></a>value;
<a id="L249"></a>}

<a id="L251"></a><span class="comment">// Get returns the underlying int32 value.</span>
<a id="L252"></a>func (v *Int32Value) Get() int32 { return *(*int32)(v.addr) }

<a id="L254"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L255"></a>func (v *Int32Value) Set(x int32) {
    <a id="L256"></a>if !v.canSet {
        <a id="L257"></a>panic(cannotSet)
    <a id="L258"></a>}
    <a id="L259"></a>*(*int32)(v.addr) = x;
<a id="L260"></a>}

<a id="L262"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L263"></a>func (v *Int32Value) SetValue(x Value) { v.Set(x.(*Int32Value).Get()) }

<a id="L265"></a><span class="comment">// Int64Value represents an int64 value.</span>
<a id="L266"></a>type Int64Value struct {
    <a id="L267"></a>value;
<a id="L268"></a>}

<a id="L270"></a><span class="comment">// Get returns the underlying int64 value.</span>
<a id="L271"></a>func (v *Int64Value) Get() int64 { return *(*int64)(v.addr) }

<a id="L273"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L274"></a>func (v *Int64Value) Set(x int64) {
    <a id="L275"></a>if !v.canSet {
        <a id="L276"></a>panic(cannotSet)
    <a id="L277"></a>}
    <a id="L278"></a>*(*int64)(v.addr) = x;
<a id="L279"></a>}

<a id="L281"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L282"></a>func (v *Int64Value) SetValue(x Value) { v.Set(x.(*Int64Value).Get()) }

<a id="L284"></a><span class="comment">// StringValue represents a string value.</span>
<a id="L285"></a>type StringValue struct {
    <a id="L286"></a>value;
<a id="L287"></a>}

<a id="L289"></a><span class="comment">// Get returns the underlying string value.</span>
<a id="L290"></a>func (v *StringValue) Get() string { return *(*string)(v.addr) }

<a id="L292"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L293"></a>func (v *StringValue) Set(x string) {
    <a id="L294"></a>if !v.canSet {
        <a id="L295"></a>panic(cannotSet)
    <a id="L296"></a>}
    <a id="L297"></a>*(*string)(v.addr) = x;
<a id="L298"></a>}

<a id="L300"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L301"></a>func (v *StringValue) SetValue(x Value) { v.Set(x.(*StringValue).Get()) }

<a id="L303"></a><span class="comment">// UintValue represents a uint value.</span>
<a id="L304"></a>type UintValue struct {
    <a id="L305"></a>value;
<a id="L306"></a>}

<a id="L308"></a><span class="comment">// Get returns the underlying uint value.</span>
<a id="L309"></a>func (v *UintValue) Get() uint { return *(*uint)(v.addr) }

<a id="L311"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L312"></a>func (v *UintValue) Set(x uint) {
    <a id="L313"></a>if !v.canSet {
        <a id="L314"></a>panic(cannotSet)
    <a id="L315"></a>}
    <a id="L316"></a>*(*uint)(v.addr) = x;
<a id="L317"></a>}

<a id="L319"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L320"></a>func (v *UintValue) SetValue(x Value) { v.Set(x.(*UintValue).Get()) }

<a id="L322"></a><span class="comment">// Uint8Value represents a uint8 value.</span>
<a id="L323"></a>type Uint8Value struct {
    <a id="L324"></a>value;
<a id="L325"></a>}

<a id="L327"></a><span class="comment">// Get returns the underlying uint8 value.</span>
<a id="L328"></a>func (v *Uint8Value) Get() uint8 { return *(*uint8)(v.addr) }

<a id="L330"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L331"></a>func (v *Uint8Value) Set(x uint8) {
    <a id="L332"></a>if !v.canSet {
        <a id="L333"></a>panic(cannotSet)
    <a id="L334"></a>}
    <a id="L335"></a>*(*uint8)(v.addr) = x;
<a id="L336"></a>}

<a id="L338"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L339"></a>func (v *Uint8Value) SetValue(x Value) { v.Set(x.(*Uint8Value).Get()) }

<a id="L341"></a><span class="comment">// Uint16Value represents a uint16 value.</span>
<a id="L342"></a>type Uint16Value struct {
    <a id="L343"></a>value;
<a id="L344"></a>}

<a id="L346"></a><span class="comment">// Get returns the underlying uint16 value.</span>
<a id="L347"></a>func (v *Uint16Value) Get() uint16 { return *(*uint16)(v.addr) }

<a id="L349"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L350"></a>func (v *Uint16Value) Set(x uint16) {
    <a id="L351"></a>if !v.canSet {
        <a id="L352"></a>panic(cannotSet)
    <a id="L353"></a>}
    <a id="L354"></a>*(*uint16)(v.addr) = x;
<a id="L355"></a>}

<a id="L357"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L358"></a>func (v *Uint16Value) SetValue(x Value) { v.Set(x.(*Uint16Value).Get()) }

<a id="L360"></a><span class="comment">// Uint32Value represents a uint32 value.</span>
<a id="L361"></a>type Uint32Value struct {
    <a id="L362"></a>value;
<a id="L363"></a>}

<a id="L365"></a><span class="comment">// Get returns the underlying uint32 value.</span>
<a id="L366"></a>func (v *Uint32Value) Get() uint32 { return *(*uint32)(v.addr) }

<a id="L368"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L369"></a>func (v *Uint32Value) Set(x uint32) {
    <a id="L370"></a>if !v.canSet {
        <a id="L371"></a>panic(cannotSet)
    <a id="L372"></a>}
    <a id="L373"></a>*(*uint32)(v.addr) = x;
<a id="L374"></a>}

<a id="L376"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L377"></a>func (v *Uint32Value) SetValue(x Value) { v.Set(x.(*Uint32Value).Get()) }

<a id="L379"></a><span class="comment">// Uint64Value represents a uint64 value.</span>
<a id="L380"></a>type Uint64Value struct {
    <a id="L381"></a>value;
<a id="L382"></a>}

<a id="L384"></a><span class="comment">// Get returns the underlying uint64 value.</span>
<a id="L385"></a>func (v *Uint64Value) Get() uint64 { return *(*uint64)(v.addr) }

<a id="L387"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L388"></a>func (v *Uint64Value) Set(x uint64) {
    <a id="L389"></a>if !v.canSet {
        <a id="L390"></a>panic(cannotSet)
    <a id="L391"></a>}
    <a id="L392"></a>*(*uint64)(v.addr) = x;
<a id="L393"></a>}

<a id="L395"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L396"></a>func (v *Uint64Value) SetValue(x Value) { v.Set(x.(*Uint64Value).Get()) }

<a id="L398"></a><span class="comment">// UintptrValue represents a uintptr value.</span>
<a id="L399"></a>type UintptrValue struct {
    <a id="L400"></a>value;
<a id="L401"></a>}

<a id="L403"></a><span class="comment">// Get returns the underlying uintptr value.</span>
<a id="L404"></a>func (v *UintptrValue) Get() uintptr { return *(*uintptr)(v.addr) }

<a id="L406"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L407"></a>func (v *UintptrValue) Set(x uintptr) {
    <a id="L408"></a>if !v.canSet {
        <a id="L409"></a>panic(cannotSet)
    <a id="L410"></a>}
    <a id="L411"></a>*(*uintptr)(v.addr) = x;
<a id="L412"></a>}

<a id="L414"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L415"></a>func (v *UintptrValue) SetValue(x Value) { v.Set(x.(*UintptrValue).Get()) }

<a id="L417"></a><span class="comment">// UnsafePointerValue represents an unsafe.Pointer value.</span>
<a id="L418"></a>type UnsafePointerValue struct {
    <a id="L419"></a>value;
<a id="L420"></a>}

<a id="L422"></a><span class="comment">// Get returns the underlying uintptr value.</span>
<a id="L423"></a><span class="comment">// Get returns uintptr, not unsafe.Pointer, so that</span>
<a id="L424"></a><span class="comment">// programs that do not import &#34;unsafe&#34; cannot</span>
<a id="L425"></a><span class="comment">// obtain a value of unsafe.Pointer type from &#34;reflect&#34;.</span>
<a id="L426"></a>func (v *UnsafePointerValue) Get() uintptr { return uintptr(*(*unsafe.Pointer)(v.addr)) }

<a id="L428"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L429"></a>func (v *UnsafePointerValue) Set(x unsafe.Pointer) {
    <a id="L430"></a>if !v.canSet {
        <a id="L431"></a>panic(cannotSet)
    <a id="L432"></a>}
    <a id="L433"></a>*(*unsafe.Pointer)(v.addr) = x;
<a id="L434"></a>}

<a id="L436"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L437"></a>func (v *UnsafePointerValue) SetValue(x Value) {
    <a id="L438"></a>v.Set(unsafe.Pointer(x.(*UnsafePointerValue).Get()))
<a id="L439"></a>}

<a id="L441"></a>func typesMustMatch(t1, t2 Type) {
    <a id="L442"></a>if t1 != t2 {
        <a id="L443"></a>panicln(&#34;type mismatch:&#34;, t1.String(), &#34;!=&#34;, t2.String())
    <a id="L444"></a>}
<a id="L445"></a>}

<a id="L447"></a><span class="comment">/*</span>
<a id="L448"></a><span class="comment"> * array</span>
<a id="L449"></a><span class="comment"> */</span>

<a id="L451"></a><span class="comment">// ArrayOrSliceValue is the common interface</span>
<a id="L452"></a><span class="comment">// implemented by both ArrayValue and SliceValue.</span>
<a id="L453"></a>type ArrayOrSliceValue interface {
    <a id="L454"></a>Value;
    <a id="L455"></a>Len() int;
    <a id="L456"></a>Cap() int;
    <a id="L457"></a>Elem(i int) Value;
    <a id="L458"></a>addr() addr;
<a id="L459"></a>}

<a id="L461"></a><span class="comment">// ArrayCopy copies the contents of src into dst until either</span>
<a id="L462"></a><span class="comment">// dst has been filled or src has been exhausted.</span>
<a id="L463"></a><span class="comment">// It returns the number of elements copied.</span>
<a id="L464"></a><span class="comment">// The arrays dst and src must have the same element type.</span>
<a id="L465"></a>func ArrayCopy(dst, src ArrayOrSliceValue) int {
    <a id="L466"></a><span class="comment">// TODO: This will have to move into the runtime</span>
    <a id="L467"></a><span class="comment">// once the real gc goes in.</span>
    <a id="L468"></a>de := dst.Type().(ArrayOrSliceType).Elem();
    <a id="L469"></a>se := src.Type().(ArrayOrSliceType).Elem();
    <a id="L470"></a>typesMustMatch(de, se);
    <a id="L471"></a>n := dst.Len();
    <a id="L472"></a>if xn := src.Len(); n &gt; xn {
        <a id="L473"></a>n = xn
    <a id="L474"></a>}
    <a id="L475"></a>memmove(dst.addr(), src.addr(), uintptr(n)*de.Size());
    <a id="L476"></a>return n;
<a id="L477"></a>}

<a id="L479"></a><span class="comment">// An ArrayValue represents an array.</span>
<a id="L480"></a>type ArrayValue struct {
    <a id="L481"></a>value;
<a id="L482"></a>}

<a id="L484"></a><span class="comment">// Len returns the length of the array.</span>
<a id="L485"></a>func (v *ArrayValue) Len() int { return v.typ.(*ArrayType).Len() }

<a id="L487"></a><span class="comment">// Cap returns the capacity of the array (equal to Len()).</span>
<a id="L488"></a>func (v *ArrayValue) Cap() int { return v.typ.(*ArrayType).Len() }

<a id="L490"></a><span class="comment">// addr returns the base address of the data in the array.</span>
<a id="L491"></a>func (v *ArrayValue) addr() addr { return v.value.addr }

<a id="L493"></a><span class="comment">// Set assigns x to v.</span>
<a id="L494"></a><span class="comment">// The new value x must have the same type as v.</span>
<a id="L495"></a>func (v *ArrayValue) Set(x *ArrayValue) {
    <a id="L496"></a>if !v.canSet {
        <a id="L497"></a>panic(cannotSet)
    <a id="L498"></a>}
    <a id="L499"></a>typesMustMatch(v.typ, x.typ);
    <a id="L500"></a>ArrayCopy(v, x);
<a id="L501"></a>}

<a id="L503"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L504"></a>func (v *ArrayValue) SetValue(x Value) { v.Set(x.(*ArrayValue)) }

<a id="L506"></a><span class="comment">// Elem returns the i&#39;th element of v.</span>
<a id="L507"></a>func (v *ArrayValue) Elem(i int) Value {
    <a id="L508"></a>typ := v.typ.(*ArrayType).Elem();
    <a id="L509"></a>n := v.Len();
    <a id="L510"></a>if i &lt; 0 || i &gt;= n {
        <a id="L511"></a>panic(&#34;index&#34;, i, &#34;in array len&#34;, n)
    <a id="L512"></a>}
    <a id="L513"></a>p := addr(uintptr(v.addr()) + uintptr(i)*typ.Size());
    <a id="L514"></a>return newValue(typ, p, v.canSet);
<a id="L515"></a>}

<a id="L517"></a><span class="comment">/*</span>
<a id="L518"></a><span class="comment"> * slice</span>
<a id="L519"></a><span class="comment"> */</span>

<a id="L521"></a><span class="comment">// runtime representation of slice</span>
<a id="L522"></a>type SliceHeader struct {
    <a id="L523"></a>Data uintptr;
    <a id="L524"></a>Len  int;
    <a id="L525"></a>Cap  int;
<a id="L526"></a>}

<a id="L528"></a><span class="comment">// A SliceValue represents a slice.</span>
<a id="L529"></a>type SliceValue struct {
    <a id="L530"></a>value;
<a id="L531"></a>}

<a id="L533"></a>func (v *SliceValue) slice() *SliceHeader { return (*SliceHeader)(v.value.addr) }

<a id="L535"></a><span class="comment">// IsNil returns whether v is a nil slice.</span>
<a id="L536"></a>func (v *SliceValue) IsNil() bool { return v.slice().Data == 0 }

<a id="L538"></a><span class="comment">// Len returns the length of the slice.</span>
<a id="L539"></a>func (v *SliceValue) Len() int { return int(v.slice().Len) }

<a id="L541"></a><span class="comment">// Cap returns the capacity of the slice.</span>
<a id="L542"></a>func (v *SliceValue) Cap() int { return int(v.slice().Cap) }

<a id="L544"></a><span class="comment">// addr returns the base address of the data in the slice.</span>
<a id="L545"></a>func (v *SliceValue) addr() addr { return addr(v.slice().Data) }

<a id="L547"></a><span class="comment">// SetLen changes the length of v.</span>
<a id="L548"></a><span class="comment">// The new length n must be between 0 and the capacity, inclusive.</span>
<a id="L549"></a>func (v *SliceValue) SetLen(n int) {
    <a id="L550"></a>s := v.slice();
    <a id="L551"></a>if n &lt; 0 || n &gt; int(s.Cap) {
        <a id="L552"></a>panicln(&#34;SetLen&#34;, n, &#34;with capacity&#34;, s.Cap)
    <a id="L553"></a>}
    <a id="L554"></a>s.Len = n;
<a id="L555"></a>}

<a id="L557"></a><span class="comment">// Set assigns x to v.</span>
<a id="L558"></a><span class="comment">// The new value x must have the same type as v.</span>
<a id="L559"></a>func (v *SliceValue) Set(x *SliceValue) {
    <a id="L560"></a>if !v.canSet {
        <a id="L561"></a>panic(cannotSet)
    <a id="L562"></a>}
    <a id="L563"></a>typesMustMatch(v.typ, x.typ);
    <a id="L564"></a>*v.slice() = *x.slice();
<a id="L565"></a>}

<a id="L567"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L568"></a>func (v *SliceValue) SetValue(x Value) { v.Set(x.(*SliceValue)) }

<a id="L570"></a><span class="comment">// Slice returns a sub-slice of the slice v.</span>
<a id="L571"></a>func (v *SliceValue) Slice(beg, end int) *SliceValue {
    <a id="L572"></a>cap := v.Cap();
    <a id="L573"></a>if beg &lt; 0 || end &lt; beg || end &gt; cap {
        <a id="L574"></a>panic(&#34;slice bounds [&#34;, beg, &#34;:&#34;, end, &#34;] with capacity &#34;, cap)
    <a id="L575"></a>}
    <a id="L576"></a>typ := v.typ.(*SliceType);
    <a id="L577"></a>s := new(SliceHeader);
    <a id="L578"></a>s.Data = uintptr(v.addr()) + uintptr(beg)*typ.Elem().Size();
    <a id="L579"></a>s.Len = end - beg;
    <a id="L580"></a>s.Cap = cap - beg;
    <a id="L581"></a>return newValue(typ, addr(s), v.canSet).(*SliceValue);
<a id="L582"></a>}

<a id="L584"></a><span class="comment">// Elem returns the i&#39;th element of v.</span>
<a id="L585"></a>func (v *SliceValue) Elem(i int) Value {
    <a id="L586"></a>typ := v.typ.(*SliceType).Elem();
    <a id="L587"></a>n := v.Len();
    <a id="L588"></a>if i &lt; 0 || i &gt;= n {
        <a id="L589"></a>panicln(&#34;index&#34;, i, &#34;in array of length&#34;, n)
    <a id="L590"></a>}
    <a id="L591"></a>p := addr(uintptr(v.addr()) + uintptr(i)*typ.Size());
    <a id="L592"></a>return newValue(typ, p, v.canSet);
<a id="L593"></a>}

<a id="L595"></a><span class="comment">// MakeSlice creates a new zero-initialized slice value</span>
<a id="L596"></a><span class="comment">// for the specified slice type, length, and capacity.</span>
<a id="L597"></a>func MakeSlice(typ *SliceType, len, cap int) *SliceValue {
    <a id="L598"></a>s := new(SliceHeader);
    <a id="L599"></a>size := typ.Elem().Size() * uintptr(cap);
    <a id="L600"></a>if size == 0 {
        <a id="L601"></a>size = 1
    <a id="L602"></a>}
    <a id="L603"></a>data := make([]uint8, size);
    <a id="L604"></a>s.Data = uintptr(addr(&amp;data[0]));
    <a id="L605"></a>s.Len = len;
    <a id="L606"></a>s.Cap = cap;
    <a id="L607"></a>return newValue(typ, addr(s), true).(*SliceValue);
<a id="L608"></a>}

<a id="L610"></a><span class="comment">/*</span>
<a id="L611"></a><span class="comment"> * chan</span>
<a id="L612"></a><span class="comment"> */</span>

<a id="L614"></a><span class="comment">// A ChanValue represents a chan.</span>
<a id="L615"></a>type ChanValue struct {
    <a id="L616"></a>value;
<a id="L617"></a>}

<a id="L619"></a><span class="comment">// IsNil returns whether v is a nil channel.</span>
<a id="L620"></a>func (v *ChanValue) IsNil() bool { return *(*uintptr)(v.addr) == 0 }

<a id="L622"></a><span class="comment">// Set assigns x to v.</span>
<a id="L623"></a><span class="comment">// The new value x must have the same type as v.</span>
<a id="L624"></a>func (v *ChanValue) Set(x *ChanValue) {
    <a id="L625"></a>if !v.canSet {
        <a id="L626"></a>panic(cannotSet)
    <a id="L627"></a>}
    <a id="L628"></a>typesMustMatch(v.typ, x.typ);
    <a id="L629"></a>*(*uintptr)(v.addr) = *(*uintptr)(x.addr);
<a id="L630"></a>}

<a id="L632"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L633"></a>func (v *ChanValue) SetValue(x Value) { v.Set(x.(*ChanValue)) }

<a id="L635"></a><span class="comment">// Get returns the uintptr value of v.</span>
<a id="L636"></a><span class="comment">// It is mainly useful for printing.</span>
<a id="L637"></a>func (v *ChanValue) Get() uintptr { return *(*uintptr)(v.addr) }

<a id="L639"></a><span class="comment">// implemented in ../pkg/runtime/reflect.cgo</span>
<a id="L640"></a>func makechan(typ *runtime.ChanType, size uint32) (ch *byte)
<a id="L641"></a>func chansend(ch, val *byte, pres *bool)
<a id="L642"></a>func chanrecv(ch, val *byte, pres *bool)
<a id="L643"></a>func chanclosed(ch *byte) bool
<a id="L644"></a>func chanclose(ch *byte)
<a id="L645"></a>func chanlen(ch *byte) int32
<a id="L646"></a>func chancap(ch *byte) int32

<a id="L648"></a><span class="comment">// Closed returns the result of closed(c) on the underlying channel.</span>
<a id="L649"></a>func (v *ChanValue) Closed() bool {
    <a id="L650"></a>ch := *(**byte)(v.addr);
    <a id="L651"></a>return chanclosed(ch);
<a id="L652"></a>}

<a id="L654"></a><span class="comment">// Close closes the channel.</span>
<a id="L655"></a>func (v *ChanValue) Close() {
    <a id="L656"></a>ch := *(**byte)(v.addr);
    <a id="L657"></a>chanclose(ch);
<a id="L658"></a>}

<a id="L660"></a>func (v *ChanValue) Len() int {
    <a id="L661"></a>ch := *(**byte)(v.addr);
    <a id="L662"></a>return int(chanlen(ch));
<a id="L663"></a>}

<a id="L665"></a>func (v *ChanValue) Cap() int {
    <a id="L666"></a>ch := *(**byte)(v.addr);
    <a id="L667"></a>return int(chancap(ch));
<a id="L668"></a>}

<a id="L670"></a><span class="comment">// internal send; non-blocking if b != nil</span>
<a id="L671"></a>func (v *ChanValue) send(x Value, b *bool) {
    <a id="L672"></a>t := v.Type().(*ChanType);
    <a id="L673"></a>if t.Dir()&amp;SendDir == 0 {
        <a id="L674"></a>panic(&#34;send on recv-only channel&#34;)
    <a id="L675"></a>}
    <a id="L676"></a>typesMustMatch(t.Elem(), x.Type());
    <a id="L677"></a>ch := *(**byte)(v.addr);
    <a id="L678"></a>chansend(ch, (*byte)(x.getAddr()), b);
<a id="L679"></a>}

<a id="L681"></a><span class="comment">// internal recv; non-blocking if b != nil</span>
<a id="L682"></a>func (v *ChanValue) recv(b *bool) Value {
    <a id="L683"></a>t := v.Type().(*ChanType);
    <a id="L684"></a>if t.Dir()&amp;RecvDir == 0 {
        <a id="L685"></a>panic(&#34;recv on send-only channel&#34;)
    <a id="L686"></a>}
    <a id="L687"></a>ch := *(**byte)(v.addr);
    <a id="L688"></a>x := MakeZero(t.Elem());
    <a id="L689"></a>chanrecv(ch, (*byte)(x.getAddr()), b);
    <a id="L690"></a>return x;
<a id="L691"></a>}

<a id="L693"></a><span class="comment">// Send sends x on the channel v.</span>
<a id="L694"></a>func (v *ChanValue) Send(x Value) { v.send(x, nil) }

<a id="L696"></a><span class="comment">// Recv receives and returns a value from the channel v.</span>
<a id="L697"></a>func (v *ChanValue) Recv() Value { return v.recv(nil) }

<a id="L699"></a><span class="comment">// TrySend attempts to sends x on the channel v but will not block.</span>
<a id="L700"></a><span class="comment">// It returns true if the value was sent, false otherwise.</span>
<a id="L701"></a>func (v *ChanValue) TrySend(x Value) bool {
    <a id="L702"></a>var ok bool;
    <a id="L703"></a>v.send(x, &amp;ok);
    <a id="L704"></a>return ok;
<a id="L705"></a>}

<a id="L707"></a><span class="comment">// TryRecv attempts to receive a value from the channel v but will not block.</span>
<a id="L708"></a><span class="comment">// It returns the value if one is received, nil otherwise.</span>
<a id="L709"></a>func (v *ChanValue) TryRecv() Value {
    <a id="L710"></a>var ok bool;
    <a id="L711"></a>x := v.recv(&amp;ok);
    <a id="L712"></a>if !ok {
        <a id="L713"></a>return nil
    <a id="L714"></a>}
    <a id="L715"></a>return x;
<a id="L716"></a>}

<a id="L718"></a><span class="comment">// MakeChan creates a new channel with the specified type and buffer size.</span>
<a id="L719"></a>func MakeChan(typ *ChanType, buffer int) *ChanValue {
    <a id="L720"></a>if buffer &lt; 0 {
        <a id="L721"></a>panic(&#34;MakeChan: negative buffer size&#34;)
    <a id="L722"></a>}
    <a id="L723"></a>if typ.Dir() != BothDir {
        <a id="L724"></a>panic(&#34;MakeChan: unidirectional channel type&#34;)
    <a id="L725"></a>}
    <a id="L726"></a>v := MakeZero(typ).(*ChanValue);
    <a id="L727"></a>*(**byte)(v.addr) = makechan((*runtime.ChanType)(unsafe.Pointer(typ)), uint32(buffer));
    <a id="L728"></a>return v;
<a id="L729"></a>}

<a id="L731"></a><span class="comment">/*</span>
<a id="L732"></a><span class="comment"> * func</span>
<a id="L733"></a><span class="comment"> */</span>

<a id="L735"></a><span class="comment">// A FuncValue represents a function value.</span>
<a id="L736"></a>type FuncValue struct {
    <a id="L737"></a>value;
    <a id="L738"></a>first       *value;
    <a id="L739"></a>isInterface bool;
<a id="L740"></a>}

<a id="L742"></a><span class="comment">// IsNil returns whether v is a nil function.</span>
<a id="L743"></a>func (v *FuncValue) IsNil() bool { return *(*uintptr)(v.addr) == 0 }

<a id="L745"></a><span class="comment">// Get returns the uintptr value of v.</span>
<a id="L746"></a><span class="comment">// It is mainly useful for printing.</span>
<a id="L747"></a>func (v *FuncValue) Get() uintptr { return *(*uintptr)(v.addr) }

<a id="L749"></a><span class="comment">// Set assigns x to v.</span>
<a id="L750"></a><span class="comment">// The new value x must have the same type as v.</span>
<a id="L751"></a>func (v *FuncValue) Set(x *FuncValue) {
    <a id="L752"></a>if !v.canSet {
        <a id="L753"></a>panic(cannotSet)
    <a id="L754"></a>}
    <a id="L755"></a>typesMustMatch(v.typ, x.typ);
    <a id="L756"></a>*(*uintptr)(v.addr) = *(*uintptr)(x.addr);
<a id="L757"></a>}

<a id="L759"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L760"></a>func (v *FuncValue) SetValue(x Value) { v.Set(x.(*FuncValue)) }

<a id="L762"></a><span class="comment">// Method returns a FuncValue corresponding to v&#39;s i&#39;th method.</span>
<a id="L763"></a><span class="comment">// The arguments to a Call on the returned FuncValue</span>
<a id="L764"></a><span class="comment">// should not include a receiver; the FuncValue will use v</span>
<a id="L765"></a><span class="comment">// as the receiver.</span>
<a id="L766"></a>func (v *value) Method(i int) *FuncValue {
    <a id="L767"></a>t := v.Type().uncommon();
    <a id="L768"></a>if t == nil || i &lt; 0 || i &gt;= len(t.methods) {
        <a id="L769"></a>return nil
    <a id="L770"></a>}
    <a id="L771"></a>p := &amp;t.methods[i];
    <a id="L772"></a>fn := p.tfn;
    <a id="L773"></a>fv := &amp;FuncValue{value: value{toType(*p.typ), addr(&amp;fn), true}, first: v, isInterface: false};
    <a id="L774"></a>return fv;
<a id="L775"></a>}

<a id="L777"></a><span class="comment">// implemented in ../pkg/runtime/*/asm.s</span>
<a id="L778"></a>func call(fn, arg *byte, n uint32)

<a id="L780"></a>type tiny struct {
    <a id="L781"></a>b byte;
<a id="L782"></a>}

<a id="L784"></a><span class="comment">// Call calls the function v with input parameters in.</span>
<a id="L785"></a><span class="comment">// It returns the function&#39;s output parameters as Values.</span>
<a id="L786"></a>func (fv *FuncValue) Call(in []Value) []Value {
    <a id="L787"></a>var structAlign = Typeof((*tiny)(nil)).(*PtrType).Elem().Size();

    <a id="L789"></a>t := fv.Type().(*FuncType);
    <a id="L790"></a>nin := len(in);
    <a id="L791"></a>if fv.first != nil &amp;&amp; !fv.isInterface {
        <a id="L792"></a>nin++
    <a id="L793"></a>}
    <a id="L794"></a>if nin != t.NumIn() {
        <a id="L795"></a>panic(&#34;FuncValue: wrong argument count&#34;)
    <a id="L796"></a>}
    <a id="L797"></a>nout := t.NumOut();

    <a id="L799"></a><span class="comment">// Compute arg size &amp; allocate.</span>
    <a id="L800"></a><span class="comment">// This computation is 6g/8g-dependent</span>
    <a id="L801"></a><span class="comment">// and probably wrong for gccgo, but so</span>
    <a id="L802"></a><span class="comment">// is most of this function.</span>
    <a id="L803"></a>size := uintptr(0);
    <a id="L804"></a>if fv.isInterface {
        <a id="L805"></a><span class="comment">// extra word for interface value</span>
        <a id="L806"></a>size += ptrSize
    <a id="L807"></a>}
    <a id="L808"></a>for i := 0; i &lt; nin; i++ {
        <a id="L809"></a>tv := t.In(i);
        <a id="L810"></a>a := uintptr(tv.Align());
        <a id="L811"></a>size = (size + a - 1) &amp;^ (a - 1);
        <a id="L812"></a>size += tv.Size();
    <a id="L813"></a>}
    <a id="L814"></a>size = (size + structAlign - 1) &amp;^ (structAlign - 1);
    <a id="L815"></a>for i := 0; i &lt; nout; i++ {
        <a id="L816"></a>tv := t.Out(i);
        <a id="L817"></a>a := uintptr(tv.Align());
        <a id="L818"></a>size = (size + a - 1) &amp;^ (a - 1);
        <a id="L819"></a>size += tv.Size();
    <a id="L820"></a>}

    <a id="L822"></a><span class="comment">// size must be &gt; 0 in order for &amp;args[0] to be valid.</span>
    <a id="L823"></a><span class="comment">// the argument copying is going to round it up to</span>
    <a id="L824"></a><span class="comment">// a multiple of 8 anyway, so make it 8 to begin with.</span>
    <a id="L825"></a>if size &lt; 8 {
        <a id="L826"></a>size = 8
    <a id="L827"></a>}
    <a id="L828"></a>args := make([]byte, size);
    <a id="L829"></a>ptr := uintptr(unsafe.Pointer(&amp;args[0]));

    <a id="L831"></a><span class="comment">// Copy into args.</span>
    <a id="L832"></a><span class="comment">//</span>
    <a id="L833"></a><span class="comment">// TODO(rsc): revisit when reference counting happens.</span>
    <a id="L834"></a><span class="comment">// This one may be fine.  The values are holding up the</span>
    <a id="L835"></a><span class="comment">// references for us, so maybe this can be treated</span>
    <a id="L836"></a><span class="comment">// like any stack-to-stack copy.</span>
    <a id="L837"></a>off := uintptr(0);
    <a id="L838"></a>delta := 0;
    <a id="L839"></a>if v := fv.first; v != nil {
        <a id="L840"></a><span class="comment">// Hard-wired first argument.</span>
        <a id="L841"></a>if fv.isInterface {
            <a id="L842"></a><span class="comment">// v is a single uninterpreted word</span>
            <a id="L843"></a>memmove(addr(ptr), v.getAddr(), ptrSize);
            <a id="L844"></a>off = ptrSize;
        <a id="L845"></a>} else {
            <a id="L846"></a><span class="comment">// v is a real value</span>
            <a id="L847"></a>tv := v.Type();
            <a id="L848"></a>typesMustMatch(t.In(0), tv);
            <a id="L849"></a>n := tv.Size();
            <a id="L850"></a>memmove(addr(ptr), v.getAddr(), n);
            <a id="L851"></a>off = n;
            <a id="L852"></a>delta = 1;
        <a id="L853"></a>}
    <a id="L854"></a>}
    <a id="L855"></a>for i, v := range in {
        <a id="L856"></a>tv := v.Type();
        <a id="L857"></a>typesMustMatch(t.In(i+delta), tv);
        <a id="L858"></a>a := uintptr(tv.Align());
        <a id="L859"></a>off = (off + a - 1) &amp;^ (a - 1);
        <a id="L860"></a>n := tv.Size();
        <a id="L861"></a>memmove(addr(ptr+off), v.getAddr(), n);
        <a id="L862"></a>off += n;
    <a id="L863"></a>}
    <a id="L864"></a>off = (off + structAlign - 1) &amp;^ (structAlign - 1);

    <a id="L866"></a><span class="comment">// Call</span>
    <a id="L867"></a>call(*(**byte)(fv.addr), (*byte)(addr(ptr)), uint32(size));

    <a id="L869"></a><span class="comment">// Copy return values out of args.</span>
    <a id="L870"></a><span class="comment">//</span>
    <a id="L871"></a><span class="comment">// TODO(rsc): revisit like above.</span>
    <a id="L872"></a>ret := make([]Value, nout);
    <a id="L873"></a>for i := 0; i &lt; nout; i++ {
        <a id="L874"></a>tv := t.Out(i);
        <a id="L875"></a>a := uintptr(tv.Align());
        <a id="L876"></a>off = (off + a - 1) &amp;^ (a - 1);
        <a id="L877"></a>v := MakeZero(tv);
        <a id="L878"></a>n := tv.Size();
        <a id="L879"></a>memmove(v.getAddr(), addr(ptr+off), n);
        <a id="L880"></a>ret[i] = v;
        <a id="L881"></a>off += n;
    <a id="L882"></a>}

    <a id="L884"></a>return ret;
<a id="L885"></a>}

<a id="L887"></a><span class="comment">/*</span>
<a id="L888"></a><span class="comment"> * interface</span>
<a id="L889"></a><span class="comment"> */</span>

<a id="L891"></a><span class="comment">// An InterfaceValue represents an interface value.</span>
<a id="L892"></a>type InterfaceValue struct {
    <a id="L893"></a>value;
<a id="L894"></a>}

<a id="L896"></a><span class="comment">// No Get because v.Interface() is available.</span>

<a id="L898"></a><span class="comment">// IsNil returns whether v is a nil interface value.</span>
<a id="L899"></a>func (v *InterfaceValue) IsNil() bool { return v.Interface() == nil }

<a id="L901"></a><span class="comment">// Elem returns the concrete value stored in the interface value v.</span>
<a id="L902"></a>func (v *InterfaceValue) Elem() Value { return NewValue(v.Interface()) }

<a id="L904"></a><span class="comment">// ../runtime/reflect.cgo</span>
<a id="L905"></a>func setiface(typ *InterfaceType, x *interface{}, addr addr)

<a id="L907"></a><span class="comment">// Set assigns x to v.</span>
<a id="L908"></a>func (v *InterfaceValue) Set(x Value) {
    <a id="L909"></a>i := x.Interface();
    <a id="L910"></a>if !v.canSet {
        <a id="L911"></a>panic(cannotSet)
    <a id="L912"></a>}
    <a id="L913"></a><span class="comment">// Two different representations; see comment in Get.</span>
    <a id="L914"></a><span class="comment">// Empty interface is easy.</span>
    <a id="L915"></a>t := v.typ.(*InterfaceType);
    <a id="L916"></a>if t.NumMethod() == 0 {
        <a id="L917"></a>*(*interface{})(v.addr) = i;
        <a id="L918"></a>return;
    <a id="L919"></a>}

    <a id="L921"></a><span class="comment">// Non-empty interface requires a runtime check.</span>
    <a id="L922"></a>setiface(t, &amp;i, v.addr);
<a id="L923"></a>}

<a id="L925"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L926"></a>func (v *InterfaceValue) SetValue(x Value) { v.Set(x) }

<a id="L928"></a><span class="comment">// Method returns a FuncValue corresponding to v&#39;s i&#39;th method.</span>
<a id="L929"></a><span class="comment">// The arguments to a Call on the returned FuncValue</span>
<a id="L930"></a><span class="comment">// should not include a receiver; the FuncValue will use v</span>
<a id="L931"></a><span class="comment">// as the receiver.</span>
<a id="L932"></a>func (v *InterfaceValue) Method(i int) *FuncValue {
    <a id="L933"></a>t := v.Type().(*InterfaceType);
    <a id="L934"></a>if t == nil || i &lt; 0 || i &gt;= len(t.methods) {
        <a id="L935"></a>return nil
    <a id="L936"></a>}
    <a id="L937"></a>p := &amp;t.methods[i];

    <a id="L939"></a><span class="comment">// Interface is two words: itable, data.</span>
    <a id="L940"></a>tab := *(**runtime.Itable)(v.addr);
    <a id="L941"></a>data := &amp;value{Typeof((*byte)(nil)), addr(uintptr(v.addr) + ptrSize), true};

    <a id="L943"></a><span class="comment">// Function pointer is at p.perm in the table.</span>
    <a id="L944"></a>fn := tab.Fn[p.perm];
    <a id="L945"></a>fv := &amp;FuncValue{value: value{toType(*p.typ), addr(&amp;fn), true}, first: data, isInterface: true};
    <a id="L946"></a>return fv;
<a id="L947"></a>}

<a id="L949"></a><span class="comment">/*</span>
<a id="L950"></a><span class="comment"> * map</span>
<a id="L951"></a><span class="comment"> */</span>

<a id="L953"></a><span class="comment">// A MapValue represents a map value.</span>
<a id="L954"></a>type MapValue struct {
    <a id="L955"></a>value;
<a id="L956"></a>}

<a id="L958"></a><span class="comment">// IsNil returns whether v is a nil map value.</span>
<a id="L959"></a>func (v *MapValue) IsNil() bool { return *(*uintptr)(v.addr) == 0 }

<a id="L961"></a><span class="comment">// Set assigns x to v.</span>
<a id="L962"></a><span class="comment">// The new value x must have the same type as v.</span>
<a id="L963"></a>func (v *MapValue) Set(x *MapValue) {
    <a id="L964"></a>if !v.canSet {
        <a id="L965"></a>panic(cannotSet)
    <a id="L966"></a>}
    <a id="L967"></a>typesMustMatch(v.typ, x.typ);
    <a id="L968"></a>*(*uintptr)(v.addr) = *(*uintptr)(x.addr);
<a id="L969"></a>}

<a id="L971"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L972"></a>func (v *MapValue) SetValue(x Value) { v.Set(x.(*MapValue)) }

<a id="L974"></a><span class="comment">// implemented in ../pkg/runtime/reflect.cgo</span>
<a id="L975"></a>func mapaccess(m, key, val *byte) bool
<a id="L976"></a>func mapassign(m, key, val *byte)
<a id="L977"></a>func maplen(m *byte) int32
<a id="L978"></a>func mapiterinit(m *byte) *byte
<a id="L979"></a>func mapiternext(it *byte)
<a id="L980"></a>func mapiterkey(it *byte, key *byte) bool
<a id="L981"></a>func makemap(t *runtime.MapType) *byte

<a id="L983"></a><span class="comment">// Elem returns the value associated with key in the map v.</span>
<a id="L984"></a><span class="comment">// It returns nil if key is not found in the map.</span>
<a id="L985"></a>func (v *MapValue) Elem(key Value) Value {
    <a id="L986"></a>t := v.Type().(*MapType);
    <a id="L987"></a>typesMustMatch(t.Key(), key.Type());
    <a id="L988"></a>m := *(**byte)(v.addr);
    <a id="L989"></a>if m == nil {
        <a id="L990"></a>return nil
    <a id="L991"></a>}
    <a id="L992"></a>newval := MakeZero(t.Elem());
    <a id="L993"></a>if !mapaccess(m, (*byte)(key.getAddr()), (*byte)(newval.getAddr())) {
        <a id="L994"></a>return nil
    <a id="L995"></a>}
    <a id="L996"></a>return newval;
<a id="L997"></a>}

<a id="L999"></a><span class="comment">// SetElem sets the value associated with key in the map v to val.</span>
<a id="L1000"></a><span class="comment">// If val is nil, Put deletes the key from map.</span>
<a id="L1001"></a>func (v *MapValue) SetElem(key, val Value) {
    <a id="L1002"></a>t := v.Type().(*MapType);
    <a id="L1003"></a>typesMustMatch(t.Key(), key.Type());
    <a id="L1004"></a>var vaddr *byte;
    <a id="L1005"></a>if val != nil {
        <a id="L1006"></a>typesMustMatch(t.Elem(), val.Type());
        <a id="L1007"></a>vaddr = (*byte)(val.getAddr());
    <a id="L1008"></a>}
    <a id="L1009"></a>m := *(**byte)(v.addr);
    <a id="L1010"></a>mapassign(m, (*byte)(key.getAddr()), vaddr);
<a id="L1011"></a>}

<a id="L1013"></a><span class="comment">// Len returns the number of keys in the map v.</span>
<a id="L1014"></a>func (v *MapValue) Len() int {
    <a id="L1015"></a>m := *(**byte)(v.addr);
    <a id="L1016"></a>if m == nil {
        <a id="L1017"></a>return 0
    <a id="L1018"></a>}
    <a id="L1019"></a>return int(maplen(m));
<a id="L1020"></a>}

<a id="L1022"></a><span class="comment">// Keys returns a slice containing all the keys present in the map,</span>
<a id="L1023"></a><span class="comment">// in unspecified order.</span>
<a id="L1024"></a>func (v *MapValue) Keys() []Value {
    <a id="L1025"></a>tk := v.Type().(*MapType).Key();
    <a id="L1026"></a>m := *(**byte)(v.addr);
    <a id="L1027"></a>mlen := int32(0);
    <a id="L1028"></a>if m != nil {
        <a id="L1029"></a>mlen = maplen(m)
    <a id="L1030"></a>}
    <a id="L1031"></a>it := mapiterinit(m);
    <a id="L1032"></a>a := make([]Value, mlen);
    <a id="L1033"></a>var i int;
    <a id="L1034"></a>for i = 0; i &lt; len(a); i++ {
        <a id="L1035"></a>k := MakeZero(tk);
        <a id="L1036"></a>if !mapiterkey(it, (*byte)(k.getAddr())) {
            <a id="L1037"></a>break
        <a id="L1038"></a>}
        <a id="L1039"></a>a[i] = k;
        <a id="L1040"></a>mapiternext(it);
    <a id="L1041"></a>}
    <a id="L1042"></a>return a[0:i];
<a id="L1043"></a>}

<a id="L1045"></a><span class="comment">// MakeMap creates a new map of the specified type.</span>
<a id="L1046"></a>func MakeMap(typ *MapType) *MapValue {
    <a id="L1047"></a>v := MakeZero(typ).(*MapValue);
    <a id="L1048"></a>*(**byte)(v.addr) = makemap((*runtime.MapType)(unsafe.Pointer(typ)));
    <a id="L1049"></a>return v;
<a id="L1050"></a>}

<a id="L1052"></a><span class="comment">/*</span>
<a id="L1053"></a><span class="comment"> * ptr</span>
<a id="L1054"></a><span class="comment"> */</span>

<a id="L1056"></a><span class="comment">// A PtrValue represents a pointer.</span>
<a id="L1057"></a>type PtrValue struct {
    <a id="L1058"></a>value;
<a id="L1059"></a>}

<a id="L1061"></a><span class="comment">// IsNil returns whether v is a nil pointer.</span>
<a id="L1062"></a>func (v *PtrValue) IsNil() bool { return *(*uintptr)(v.addr) == 0 }

<a id="L1064"></a><span class="comment">// Get returns the uintptr value of v.</span>
<a id="L1065"></a><span class="comment">// It is mainly useful for printing.</span>
<a id="L1066"></a>func (v *PtrValue) Get() uintptr { return *(*uintptr)(v.addr) }

<a id="L1068"></a><span class="comment">// Set assigns x to v.</span>
<a id="L1069"></a><span class="comment">// The new value x must have the same type as v.</span>
<a id="L1070"></a>func (v *PtrValue) Set(x *PtrValue) {
    <a id="L1071"></a>if !v.canSet {
        <a id="L1072"></a>panic(cannotSet)
    <a id="L1073"></a>}
    <a id="L1074"></a>typesMustMatch(v.typ, x.typ);
    <a id="L1075"></a><span class="comment">// TODO: This will have to move into the runtime</span>
    <a id="L1076"></a><span class="comment">// once the new gc goes in</span>
    <a id="L1077"></a>*(*uintptr)(v.addr) = *(*uintptr)(x.addr);
<a id="L1078"></a>}

<a id="L1080"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L1081"></a>func (v *PtrValue) SetValue(x Value) { v.Set(x.(*PtrValue)) }

<a id="L1083"></a><span class="comment">// PointTo changes v to point to x.</span>
<a id="L1084"></a>func (v *PtrValue) PointTo(x Value) {
    <a id="L1085"></a>if !x.CanSet() {
        <a id="L1086"></a>panic(&#34;cannot set x; cannot point to x&#34;)
    <a id="L1087"></a>}
    <a id="L1088"></a>typesMustMatch(v.typ.(*PtrType).Elem(), x.Type());
    <a id="L1089"></a><span class="comment">// TODO: This will have to move into the runtime</span>
    <a id="L1090"></a><span class="comment">// once the new gc goes in.</span>
    <a id="L1091"></a>*(*uintptr)(v.addr) = x.Addr();
<a id="L1092"></a>}

<a id="L1094"></a><span class="comment">// Elem returns the value that v points to.</span>
<a id="L1095"></a><span class="comment">// If v is a nil pointer, Elem returns a nil Value.</span>
<a id="L1096"></a>func (v *PtrValue) Elem() Value {
    <a id="L1097"></a>if v.IsNil() {
        <a id="L1098"></a>return nil
    <a id="L1099"></a>}
    <a id="L1100"></a>return newValue(v.typ.(*PtrType).Elem(), *(*addr)(v.addr), v.canSet);
<a id="L1101"></a>}

<a id="L1103"></a><span class="comment">// Indirect returns the value that v points to.</span>
<a id="L1104"></a><span class="comment">// If v is a nil pointer, Indirect returns a nil Value.</span>
<a id="L1105"></a><span class="comment">// If v is not a pointer, Indirect returns v.</span>
<a id="L1106"></a>func Indirect(v Value) Value {
    <a id="L1107"></a>if pv, ok := v.(*PtrValue); ok {
        <a id="L1108"></a>return pv.Elem()
    <a id="L1109"></a>}
    <a id="L1110"></a>return v;
<a id="L1111"></a>}

<a id="L1113"></a><span class="comment">/*</span>
<a id="L1114"></a><span class="comment"> * struct</span>
<a id="L1115"></a><span class="comment"> */</span>

<a id="L1117"></a><span class="comment">// A StructValue represents a struct value.</span>
<a id="L1118"></a>type StructValue struct {
    <a id="L1119"></a>value;
<a id="L1120"></a>}

<a id="L1122"></a><span class="comment">// Set assigns x to v.</span>
<a id="L1123"></a><span class="comment">// The new value x must have the same type as v.</span>
<a id="L1124"></a>func (v *StructValue) Set(x *StructValue) {
    <a id="L1125"></a><span class="comment">// TODO: This will have to move into the runtime</span>
    <a id="L1126"></a><span class="comment">// once the gc goes in.</span>
    <a id="L1127"></a>if !v.canSet {
        <a id="L1128"></a>panic(cannotSet)
    <a id="L1129"></a>}
    <a id="L1130"></a>typesMustMatch(v.typ, x.typ);
    <a id="L1131"></a>memmove(v.addr, x.addr, v.typ.Size());
<a id="L1132"></a>}

<a id="L1134"></a><span class="comment">// Set sets v to the value x.</span>
<a id="L1135"></a>func (v *StructValue) SetValue(x Value) { v.Set(x.(*StructValue)) }

<a id="L1137"></a><span class="comment">// Field returns the i&#39;th field of the struct.</span>
<a id="L1138"></a>func (v *StructValue) Field(i int) Value {
    <a id="L1139"></a>t := v.typ.(*StructType);
    <a id="L1140"></a>if i &lt; 0 || i &gt;= t.NumField() {
        <a id="L1141"></a>return nil
    <a id="L1142"></a>}
    <a id="L1143"></a>f := t.Field(i);
    <a id="L1144"></a>return newValue(f.Type, addr(uintptr(v.addr)+f.Offset), v.canSet &amp;&amp; f.PkgPath == &#34;&#34;);
<a id="L1145"></a>}

<a id="L1147"></a><span class="comment">// FieldByIndex returns the nested field corresponding to index.</span>
<a id="L1148"></a>func (t *StructValue) FieldByIndex(index []int) (v Value) {
    <a id="L1149"></a>v = t;
    <a id="L1150"></a>for i, x := range index {
        <a id="L1151"></a>if i &gt; 0 {
            <a id="L1152"></a>if p, ok := v.(*PtrValue); ok {
                <a id="L1153"></a>v = p.Elem()
            <a id="L1154"></a>}
            <a id="L1155"></a>if s, ok := v.(*StructValue); ok {
                <a id="L1156"></a>t = s
            <a id="L1157"></a>} else {
                <a id="L1158"></a>v = nil;
                <a id="L1159"></a>return;
            <a id="L1160"></a>}
        <a id="L1161"></a>}
        <a id="L1162"></a>v = t.Field(x);
    <a id="L1163"></a>}
    <a id="L1164"></a>return;
<a id="L1165"></a>}

<a id="L1167"></a><span class="comment">// FieldByName returns the struct field with the given name.</span>
<a id="L1168"></a><span class="comment">// The result is nil if no field was found.</span>
<a id="L1169"></a>func (t *StructValue) FieldByName(name string) Value {
    <a id="L1170"></a>if f, ok := t.Type().(*StructType).FieldByName(name); ok {
        <a id="L1171"></a>return t.FieldByIndex(f.Index)
    <a id="L1172"></a>}
    <a id="L1173"></a>return nil;
<a id="L1174"></a>}

<a id="L1176"></a><span class="comment">// NumField returns the number of fields in the struct.</span>
<a id="L1177"></a>func (v *StructValue) NumField() int { return v.typ.(*StructType).NumField() }

<a id="L1179"></a><span class="comment">/*</span>
<a id="L1180"></a><span class="comment"> * constructors</span>
<a id="L1181"></a><span class="comment"> */</span>

<a id="L1183"></a><span class="comment">// NewValue returns a new Value initialized to the concrete value</span>
<a id="L1184"></a><span class="comment">// stored in the interface i.  NewValue(nil) returns nil.</span>
<a id="L1185"></a>func NewValue(i interface{}) Value {
    <a id="L1186"></a>if i == nil {
        <a id="L1187"></a>return nil
    <a id="L1188"></a>}
    <a id="L1189"></a>t, a := unsafe.Reflect(i);
    <a id="L1190"></a>return newValue(toType(t), addr(a), true);
<a id="L1191"></a>}


<a id="L1194"></a>func newFuncValue(typ Type, addr addr, canSet bool) *FuncValue {
    <a id="L1195"></a>return &amp;FuncValue{value: value{typ, addr, canSet}}
<a id="L1196"></a>}

<a id="L1198"></a>func newValue(typ Type, addr addr, canSet bool) Value {
    <a id="L1199"></a><span class="comment">// FuncValue has a different layout;</span>
    <a id="L1200"></a><span class="comment">// it needs a extra space for the fixed receivers.</span>
    <a id="L1201"></a>if _, ok := typ.(*FuncType); ok {
        <a id="L1202"></a>return newFuncValue(typ, addr, canSet)
    <a id="L1203"></a>}

    <a id="L1205"></a><span class="comment">// All values have same memory layout;</span>
    <a id="L1206"></a><span class="comment">// build once and convert.</span>
    <a id="L1207"></a>v := &amp;struct{ value }{value{typ, addr, canSet}};
    <a id="L1208"></a>switch typ.(type) {
    <a id="L1209"></a>case *ArrayType:
        <a id="L1210"></a><span class="comment">// TODO(rsc): Something must prevent</span>
        <a id="L1211"></a><span class="comment">// clients of the package from doing</span>
        <a id="L1212"></a><span class="comment">// this same kind of cast.</span>
        <a id="L1213"></a><span class="comment">// We should be allowed because</span>
        <a id="L1214"></a><span class="comment">// they&#39;re our types.</span>
        <a id="L1215"></a><span class="comment">// Something about implicit assignment</span>
        <a id="L1216"></a><span class="comment">// to struct fields.</span>
        <a id="L1217"></a>return (*ArrayValue)(v)
    <a id="L1218"></a>case *BoolType:
        <a id="L1219"></a>return (*BoolValue)(v)
    <a id="L1220"></a>case *ChanType:
        <a id="L1221"></a>return (*ChanValue)(v)
    <a id="L1222"></a>case *FloatType:
        <a id="L1223"></a>return (*FloatValue)(v)
    <a id="L1224"></a>case *Float32Type:
        <a id="L1225"></a>return (*Float32Value)(v)
    <a id="L1226"></a>case *Float64Type:
        <a id="L1227"></a>return (*Float64Value)(v)
    <a id="L1228"></a>case *IntType:
        <a id="L1229"></a>return (*IntValue)(v)
    <a id="L1230"></a>case *Int8Type:
        <a id="L1231"></a>return (*Int8Value)(v)
    <a id="L1232"></a>case *Int16Type:
        <a id="L1233"></a>return (*Int16Value)(v)
    <a id="L1234"></a>case *Int32Type:
        <a id="L1235"></a>return (*Int32Value)(v)
    <a id="L1236"></a>case *Int64Type:
        <a id="L1237"></a>return (*Int64Value)(v)
    <a id="L1238"></a>case *InterfaceType:
        <a id="L1239"></a>return (*InterfaceValue)(v)
    <a id="L1240"></a>case *MapType:
        <a id="L1241"></a>return (*MapValue)(v)
    <a id="L1242"></a>case *PtrType:
        <a id="L1243"></a>return (*PtrValue)(v)
    <a id="L1244"></a>case *SliceType:
        <a id="L1245"></a>return (*SliceValue)(v)
    <a id="L1246"></a>case *StringType:
        <a id="L1247"></a>return (*StringValue)(v)
    <a id="L1248"></a>case *StructType:
        <a id="L1249"></a>return (*StructValue)(v)
    <a id="L1250"></a>case *UintType:
        <a id="L1251"></a>return (*UintValue)(v)
    <a id="L1252"></a>case *Uint8Type:
        <a id="L1253"></a>return (*Uint8Value)(v)
    <a id="L1254"></a>case *Uint16Type:
        <a id="L1255"></a>return (*Uint16Value)(v)
    <a id="L1256"></a>case *Uint32Type:
        <a id="L1257"></a>return (*Uint32Value)(v)
    <a id="L1258"></a>case *Uint64Type:
        <a id="L1259"></a>return (*Uint64Value)(v)
    <a id="L1260"></a>case *UintptrType:
        <a id="L1261"></a>return (*UintptrValue)(v)
    <a id="L1262"></a>case *UnsafePointerType:
        <a id="L1263"></a>return (*UnsafePointerValue)(v)
    <a id="L1264"></a>}
    <a id="L1265"></a>panicln(&#34;newValue&#34;, typ.String());
<a id="L1266"></a>}

<a id="L1268"></a><span class="comment">// MakeZero returns a zero Value for the specified Type.</span>
<a id="L1269"></a>func MakeZero(typ Type) Value {
    <a id="L1270"></a><span class="comment">// TODO: this will have to move into</span>
    <a id="L1271"></a><span class="comment">// the runtime proper in order to play nicely</span>
    <a id="L1272"></a><span class="comment">// with the garbage collector.</span>
    <a id="L1273"></a>size := typ.Size();
    <a id="L1274"></a>if size == 0 {
        <a id="L1275"></a>size = 1
    <a id="L1276"></a>}
    <a id="L1277"></a>data := make([]uint8, size);
    <a id="L1278"></a>return newValue(typ, addr(&amp;data[0]), true);
<a id="L1279"></a>}
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
