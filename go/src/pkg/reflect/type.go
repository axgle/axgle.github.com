<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/reflect/type.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/reflect/type.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The reflect package implements run-time reflection, allowing a program to</span>
<a id="L6"></a><span class="comment">// manipulate objects with arbitrary types.  The typical use is to take a</span>
<a id="L7"></a><span class="comment">// value with static type interface{} and extract its dynamic type</span>
<a id="L8"></a><span class="comment">// information by calling Typeof, which returns an object with interface</span>
<a id="L9"></a><span class="comment">// type Type.  That contains a pointer to a struct of type *StructType,</span>
<a id="L10"></a><span class="comment">// *IntType, etc. representing the details of the underlying type.  A type</span>
<a id="L11"></a><span class="comment">// switch or type assertion can reveal which.</span>
<a id="L12"></a><span class="comment">//</span>
<a id="L13"></a><span class="comment">// A call to NewValue creates a Value representing the run-time data; it</span>
<a id="L14"></a><span class="comment">// contains a *StructValue, *IntValue, etc.  MakeZero takes a Type and</span>
<a id="L15"></a><span class="comment">// returns a Value representing a zero value for that type.</span>
<a id="L16"></a>package reflect

<a id="L18"></a>import (
    <a id="L19"></a>&#34;runtime&#34;;
    <a id="L20"></a>&#34;strconv&#34;;
    <a id="L21"></a>&#34;unsafe&#34;;
<a id="L22"></a>)

<a id="L24"></a><span class="comment">/*</span>
<a id="L25"></a><span class="comment"> * Copy of data structures from ../runtime/type.go.</span>
<a id="L26"></a><span class="comment"> * For comments, see the ones in that file.</span>
<a id="L27"></a><span class="comment"> *</span>
<a id="L28"></a><span class="comment"> * These data structures are known to the compiler and the runtime.</span>
<a id="L29"></a><span class="comment"> *</span>
<a id="L30"></a><span class="comment"> * Putting these types in runtime instead of reflect means that</span>
<a id="L31"></a><span class="comment"> * reflect doesn&#39;t need to be autolinked into every binary, which</span>
<a id="L32"></a><span class="comment"> * simplifies bootstrapping and package dependencies.</span>
<a id="L33"></a><span class="comment"> * Unfortunately, it also means that reflect needs its own</span>
<a id="L34"></a><span class="comment"> * copy in order to access the private fields.</span>
<a id="L35"></a><span class="comment"> */</span>

<a id="L37"></a>type commonType struct {
    <a id="L38"></a>size       uintptr;
    <a id="L39"></a>hash       uint32;
    <a id="L40"></a>alg        uint8;
    <a id="L41"></a>align      uint8;
    <a id="L42"></a>fieldAlign uint8;
    <a id="L43"></a>string     *string;
    <a id="L44"></a>*uncommonType;
<a id="L45"></a>}

<a id="L47"></a>type method struct {
    <a id="L48"></a>hash    uint32;
    <a id="L49"></a>name    *string;
    <a id="L50"></a>pkgPath *string;
    <a id="L51"></a>typ     *runtime.Type;
    <a id="L52"></a>ifn     unsafe.Pointer;
    <a id="L53"></a>tfn     unsafe.Pointer;
<a id="L54"></a>}

<a id="L56"></a>type uncommonType struct {
    <a id="L57"></a>name    *string;
    <a id="L58"></a>pkgPath *string;
    <a id="L59"></a>methods []method;
<a id="L60"></a>}

<a id="L62"></a><span class="comment">// BoolType represents a boolean type.</span>
<a id="L63"></a>type BoolType struct {
    <a id="L64"></a>commonType;
<a id="L65"></a>}

<a id="L67"></a><span class="comment">// Float32Type represents a float32 type.</span>
<a id="L68"></a>type Float32Type struct {
    <a id="L69"></a>commonType;
<a id="L70"></a>}

<a id="L72"></a><span class="comment">// Float64Type represents a float64 type.</span>
<a id="L73"></a>type Float64Type struct {
    <a id="L74"></a>commonType;
<a id="L75"></a>}

<a id="L77"></a><span class="comment">// FloatType represents a float type.</span>
<a id="L78"></a>type FloatType struct {
    <a id="L79"></a>commonType;
<a id="L80"></a>}

<a id="L82"></a><span class="comment">// Int16Type represents an int16 type.</span>
<a id="L83"></a>type Int16Type struct {
    <a id="L84"></a>commonType;
<a id="L85"></a>}

<a id="L87"></a><span class="comment">// Int32Type represents an int32 type.</span>
<a id="L88"></a>type Int32Type struct {
    <a id="L89"></a>commonType;
<a id="L90"></a>}

<a id="L92"></a><span class="comment">// Int64Type represents an int64 type.</span>
<a id="L93"></a>type Int64Type struct {
    <a id="L94"></a>commonType;
<a id="L95"></a>}

<a id="L97"></a><span class="comment">// Int8Type represents an int8 type.</span>
<a id="L98"></a>type Int8Type struct {
    <a id="L99"></a>commonType;
<a id="L100"></a>}

<a id="L102"></a><span class="comment">// IntType represents an int type.</span>
<a id="L103"></a>type IntType struct {
    <a id="L104"></a>commonType;
<a id="L105"></a>}

<a id="L107"></a><span class="comment">// Uint16Type represents a uint16 type.</span>
<a id="L108"></a>type Uint16Type struct {
    <a id="L109"></a>commonType;
<a id="L110"></a>}

<a id="L112"></a><span class="comment">// Uint32Type represents a uint32 type.</span>
<a id="L113"></a>type Uint32Type struct {
    <a id="L114"></a>commonType;
<a id="L115"></a>}

<a id="L117"></a><span class="comment">// Uint64Type represents a uint64 type.</span>
<a id="L118"></a>type Uint64Type struct {
    <a id="L119"></a>commonType;
<a id="L120"></a>}

<a id="L122"></a><span class="comment">// Uint8Type represents a uint8 type.</span>
<a id="L123"></a>type Uint8Type struct {
    <a id="L124"></a>commonType;
<a id="L125"></a>}

<a id="L127"></a><span class="comment">// UintType represents a uint type.</span>
<a id="L128"></a>type UintType struct {
    <a id="L129"></a>commonType;
<a id="L130"></a>}

<a id="L132"></a><span class="comment">// StringType represents a string type.</span>
<a id="L133"></a>type StringType struct {
    <a id="L134"></a>commonType;
<a id="L135"></a>}

<a id="L137"></a><span class="comment">// UintptrType represents a uintptr type.</span>
<a id="L138"></a>type UintptrType struct {
    <a id="L139"></a>commonType;
<a id="L140"></a>}

<a id="L142"></a><span class="comment">// DotDotDotType represents the ... that can</span>
<a id="L143"></a><span class="comment">// be used as the type of the final function parameter.</span>
<a id="L144"></a>type DotDotDotType struct {
    <a id="L145"></a>commonType;
<a id="L146"></a>}

<a id="L148"></a><span class="comment">// UnsafePointerType represents an unsafe.Pointer type.</span>
<a id="L149"></a>type UnsafePointerType struct {
    <a id="L150"></a>commonType;
<a id="L151"></a>}

<a id="L153"></a><span class="comment">// ArrayType represents a fixed array type.</span>
<a id="L154"></a>type ArrayType struct {
    <a id="L155"></a>commonType;
    <a id="L156"></a>elem *runtime.Type;
    <a id="L157"></a>len  uintptr;
<a id="L158"></a>}

<a id="L160"></a><span class="comment">// ChanDir represents a channel type&#39;s direction.</span>
<a id="L161"></a>type ChanDir int

<a id="L163"></a>const (
    <a id="L164"></a>RecvDir ChanDir = 1 &lt;&lt; iota;
    <a id="L165"></a>SendDir;
    <a id="L166"></a>BothDir = RecvDir | SendDir;
<a id="L167"></a>)

<a id="L169"></a><span class="comment">// ChanType represents a channel type.</span>
<a id="L170"></a>type ChanType struct {
    <a id="L171"></a>commonType;
    <a id="L172"></a>elem *runtime.Type;
    <a id="L173"></a>dir  uintptr;
<a id="L174"></a>}

<a id="L176"></a><span class="comment">// FuncType represents a function type.</span>
<a id="L177"></a>type FuncType struct {
    <a id="L178"></a>commonType;
    <a id="L179"></a>in  []*runtime.Type;
    <a id="L180"></a>out []*runtime.Type;
<a id="L181"></a>}

<a id="L183"></a><span class="comment">// Method on interface type</span>
<a id="L184"></a>type imethod struct {
    <a id="L185"></a>hash    uint32;
    <a id="L186"></a>perm    uint32;
    <a id="L187"></a>name    *string;
    <a id="L188"></a>pkgPath *string;
    <a id="L189"></a>typ     *runtime.Type;
<a id="L190"></a>}

<a id="L192"></a><span class="comment">// InterfaceType represents an interface type.</span>
<a id="L193"></a>type InterfaceType struct {
    <a id="L194"></a>commonType;
    <a id="L195"></a>methods []imethod;
<a id="L196"></a>}

<a id="L198"></a><span class="comment">// MapType represents a map type.</span>
<a id="L199"></a>type MapType struct {
    <a id="L200"></a>commonType;
    <a id="L201"></a>key  *runtime.Type;
    <a id="L202"></a>elem *runtime.Type;
<a id="L203"></a>}

<a id="L205"></a><span class="comment">// PtrType represents a pointer type.</span>
<a id="L206"></a>type PtrType struct {
    <a id="L207"></a>commonType;
    <a id="L208"></a>elem *runtime.Type;
<a id="L209"></a>}

<a id="L211"></a><span class="comment">// SliceType represents a slice type.</span>
<a id="L212"></a>type SliceType struct {
    <a id="L213"></a>commonType;
    <a id="L214"></a>elem *runtime.Type;
<a id="L215"></a>}

<a id="L217"></a><span class="comment">// Struct field</span>
<a id="L218"></a>type structField struct {
    <a id="L219"></a>name    *string;
    <a id="L220"></a>pkgPath *string;
    <a id="L221"></a>typ     *runtime.Type;
    <a id="L222"></a>tag     *string;
    <a id="L223"></a>offset  uintptr;
<a id="L224"></a>}

<a id="L226"></a><span class="comment">// StructType represents a struct type.</span>
<a id="L227"></a>type StructType struct {
    <a id="L228"></a>commonType;
    <a id="L229"></a>fields []structField;
<a id="L230"></a>}


<a id="L233"></a><span class="comment">/*</span>
<a id="L234"></a><span class="comment"> * The compiler knows the exact layout of all the data structures above.</span>
<a id="L235"></a><span class="comment"> * The compiler does not know about the data structures and methods below.</span>
<a id="L236"></a><span class="comment"> */</span>

<a id="L238"></a><span class="comment">// Method represents a single method.</span>
<a id="L239"></a>type Method struct {
    <a id="L240"></a>PkgPath string; <span class="comment">// empty for uppercase Name</span>
    <a id="L241"></a>Name    string;
    <a id="L242"></a>Type    *FuncType;
    <a id="L243"></a>Func    *FuncValue;
<a id="L244"></a>}

<a id="L246"></a><span class="comment">// Type is the runtime representation of a Go type.</span>
<a id="L247"></a><span class="comment">// Every type implements the methods listed here.</span>
<a id="L248"></a><span class="comment">// Some types implement additional interfaces;</span>
<a id="L249"></a><span class="comment">// use a type switch to find out what kind of type a Type is.</span>
<a id="L250"></a><span class="comment">// Each type in a program has a unique Type, so == on Types</span>
<a id="L251"></a><span class="comment">// corresponds to Go&#39;s type equality.</span>
<a id="L252"></a>type Type interface {
    <a id="L253"></a><span class="comment">// PkgPath returns the type&#39;s package path.</span>
    <a id="L254"></a><span class="comment">// The package path is a full package import path like &#34;container/vector&#34;.</span>
    <a id="L255"></a>PkgPath() string;

    <a id="L257"></a><span class="comment">// Name returns the type&#39;s name within its package.</span>
    <a id="L258"></a>Name() string;

    <a id="L260"></a><span class="comment">// String returns a string representation of the type.</span>
    <a id="L261"></a><span class="comment">// The string representation may use shortened package names</span>
    <a id="L262"></a><span class="comment">// (e.g., vector instead of &#34;container/vector&#34;) and is not</span>
    <a id="L263"></a><span class="comment">// guaranteed to be unique among types.  To test for equality,</span>
    <a id="L264"></a><span class="comment">// compare the Types directly.</span>
    <a id="L265"></a>String() string;

    <a id="L267"></a><span class="comment">// Size returns the number of bytes needed to store</span>
    <a id="L268"></a><span class="comment">// a value of the given type; it is analogous to unsafe.Sizeof.</span>
    <a id="L269"></a>Size() uintptr;

    <a id="L271"></a><span class="comment">// Align returns the alignment of a value of this type</span>
    <a id="L272"></a><span class="comment">// when allocated in memory.</span>
    <a id="L273"></a>Align() int;

    <a id="L275"></a><span class="comment">// FieldAlign returns the alignment of a value of this type</span>
    <a id="L276"></a><span class="comment">// when used as a field in a struct.</span>
    <a id="L277"></a>FieldAlign() int;

    <a id="L279"></a><span class="comment">// For non-interface types, Method returns the i&#39;th method with receiver T.</span>
    <a id="L280"></a><span class="comment">// For interface types, Method returns the i&#39;th method in the interface.</span>
    <a id="L281"></a><span class="comment">// NumMethod returns the number of such methods.</span>
    <a id="L282"></a>Method(int) Method;
    <a id="L283"></a>NumMethod() int;
    <a id="L284"></a>uncommon() *uncommonType;
<a id="L285"></a>}

<a id="L287"></a>func (t *uncommonType) uncommon() *uncommonType {
    <a id="L288"></a>return t
<a id="L289"></a>}

<a id="L291"></a>func (t *uncommonType) PkgPath() string {
    <a id="L292"></a>if t == nil || t.pkgPath == nil {
        <a id="L293"></a>return &#34;&#34;
    <a id="L294"></a>}
    <a id="L295"></a>return *t.pkgPath;
<a id="L296"></a>}

<a id="L298"></a>func (t *uncommonType) Name() string {
    <a id="L299"></a>if t == nil || t.name == nil {
        <a id="L300"></a>return &#34;&#34;
    <a id="L301"></a>}
    <a id="L302"></a>return *t.name;
<a id="L303"></a>}

<a id="L305"></a>func (t *commonType) String() string { return *t.string }

<a id="L307"></a>func (t *commonType) Size() uintptr { return t.size }

<a id="L309"></a>func (t *commonType) Align() int { return int(t.align) }

<a id="L311"></a>func (t *commonType) FieldAlign() int { return int(t.fieldAlign) }

<a id="L313"></a>func (t *uncommonType) Method(i int) (m Method) {
    <a id="L314"></a>if t == nil || i &lt; 0 || i &gt;= len(t.methods) {
        <a id="L315"></a>return
    <a id="L316"></a>}
    <a id="L317"></a>p := &amp;t.methods[i];
    <a id="L318"></a>if p.name != nil {
        <a id="L319"></a>m.Name = *p.name
    <a id="L320"></a>}
    <a id="L321"></a>if p.pkgPath != nil {
        <a id="L322"></a>m.PkgPath = *p.pkgPath
    <a id="L323"></a>}
    <a id="L324"></a>m.Type = toType(*p.typ).(*FuncType);
    <a id="L325"></a>fn := p.tfn;
    <a id="L326"></a>m.Func = newFuncValue(m.Type, addr(&amp;fn), true);
    <a id="L327"></a>return;
<a id="L328"></a>}

<a id="L330"></a>func (t *uncommonType) NumMethod() int {
    <a id="L331"></a>if t == nil {
        <a id="L332"></a>return 0
    <a id="L333"></a>}
    <a id="L334"></a>return len(t.methods);
<a id="L335"></a>}

<a id="L337"></a><span class="comment">// TODO(rsc): 6g supplies these, but they are not</span>
<a id="L338"></a><span class="comment">// as efficient as they could be: they have commonType</span>
<a id="L339"></a><span class="comment">// as the receiver instead of *commonType.</span>
<a id="L340"></a>func (t *commonType) NumMethod() int { return t.uncommonType.NumMethod() }

<a id="L342"></a>func (t *commonType) Method(i int) (m Method) { return t.uncommonType.Method(i) }

<a id="L344"></a>func (t *commonType) PkgPath() string { return t.uncommonType.PkgPath() }

<a id="L346"></a>func (t *commonType) Name() string { return t.uncommonType.Name() }

<a id="L348"></a><span class="comment">// Len returns the number of elements in the array.</span>
<a id="L349"></a>func (t *ArrayType) Len() int { return int(t.len) }

<a id="L351"></a><span class="comment">// Elem returns the type of the array&#39;s elements.</span>
<a id="L352"></a>func (t *ArrayType) Elem() Type { return toType(*t.elem) }

<a id="L354"></a><span class="comment">// Dir returns the channel direction.</span>
<a id="L355"></a>func (t *ChanType) Dir() ChanDir { return ChanDir(t.dir) }

<a id="L357"></a><span class="comment">// Elem returns the channel&#39;s element type.</span>
<a id="L358"></a>func (t *ChanType) Elem() Type { return toType(*t.elem) }

<a id="L360"></a>func (d ChanDir) String() string {
    <a id="L361"></a>switch d {
    <a id="L362"></a>case SendDir:
        <a id="L363"></a>return &#34;chan&lt;-&#34;
    <a id="L364"></a>case RecvDir:
        <a id="L365"></a>return &#34;&lt;-chan&#34;
    <a id="L366"></a>case BothDir:
        <a id="L367"></a>return &#34;chan&#34;
    <a id="L368"></a>}
    <a id="L369"></a>return &#34;ChanDir&#34; + strconv.Itoa(int(d));
<a id="L370"></a>}

<a id="L372"></a><span class="comment">// In returns the type of the i&#39;th function input parameter.</span>
<a id="L373"></a>func (t *FuncType) In(i int) Type {
    <a id="L374"></a>if i &lt; 0 || i &gt;= len(t.in) {
        <a id="L375"></a>return nil
    <a id="L376"></a>}
    <a id="L377"></a>return toType(*t.in[i]);
<a id="L378"></a>}

<a id="L380"></a><span class="comment">// NumIn returns the number of input parameters.</span>
<a id="L381"></a>func (t *FuncType) NumIn() int { return len(t.in) }

<a id="L383"></a><span class="comment">// Out returns the type of the i&#39;th function output parameter.</span>
<a id="L384"></a>func (t *FuncType) Out(i int) Type {
    <a id="L385"></a>if i &lt; 0 || i &gt;= len(t.out) {
        <a id="L386"></a>return nil
    <a id="L387"></a>}
    <a id="L388"></a>return toType(*t.out[i]);
<a id="L389"></a>}

<a id="L391"></a><span class="comment">// NumOut returns the number of function output parameters.</span>
<a id="L392"></a>func (t *FuncType) NumOut() int { return len(t.out) }

<a id="L394"></a><span class="comment">// Method returns the i&#39;th interface method.</span>
<a id="L395"></a>func (t *InterfaceType) Method(i int) (m Method) {
    <a id="L396"></a>if i &lt; 0 || i &gt;= len(t.methods) {
        <a id="L397"></a>return
    <a id="L398"></a>}
    <a id="L399"></a>p := &amp;t.methods[i];
    <a id="L400"></a>m.Name = *p.name;
    <a id="L401"></a>if p.pkgPath != nil {
        <a id="L402"></a>m.PkgPath = *p.pkgPath
    <a id="L403"></a>}
    <a id="L404"></a>m.Type = toType(*p.typ).(*FuncType);
    <a id="L405"></a>return;
<a id="L406"></a>}

<a id="L408"></a><span class="comment">// NumMethod returns the number of interface methods.</span>
<a id="L409"></a>func (t *InterfaceType) NumMethod() int { return len(t.methods) }

<a id="L411"></a><span class="comment">// Key returns the map key type.</span>
<a id="L412"></a>func (t *MapType) Key() Type { return toType(*t.key) }

<a id="L414"></a><span class="comment">// Elem returns the map element type.</span>
<a id="L415"></a>func (t *MapType) Elem() Type { return toType(*t.elem) }

<a id="L417"></a><span class="comment">// Elem returns the pointer element type.</span>
<a id="L418"></a>func (t *PtrType) Elem() Type { return toType(*t.elem) }

<a id="L420"></a><span class="comment">// Elem returns the type of the slice&#39;s elements.</span>
<a id="L421"></a>func (t *SliceType) Elem() Type { return toType(*t.elem) }

<a id="L423"></a>type StructField struct {
    <a id="L424"></a>PkgPath   string; <span class="comment">// empty for uppercase Name</span>
    <a id="L425"></a>Name      string;
    <a id="L426"></a>Type      Type;
    <a id="L427"></a>Tag       string;
    <a id="L428"></a>Offset    uintptr;
    <a id="L429"></a>Index     []int;
    <a id="L430"></a>Anonymous bool;
<a id="L431"></a>}

<a id="L433"></a><span class="comment">// Field returns the i&#39;th struct field.</span>
<a id="L434"></a>func (t *StructType) Field(i int) (f StructField) {
    <a id="L435"></a>if i &lt; 0 || i &gt;= len(t.fields) {
        <a id="L436"></a>return
    <a id="L437"></a>}
    <a id="L438"></a>p := t.fields[i];
    <a id="L439"></a>f.Type = toType(*p.typ);
    <a id="L440"></a>if p.name != nil {
        <a id="L441"></a>f.Name = *p.name
    <a id="L442"></a>} else {
        <a id="L443"></a>t := f.Type;
        <a id="L444"></a>if pt, ok := t.(*PtrType); ok {
            <a id="L445"></a>t = pt.Elem()
        <a id="L446"></a>}
        <a id="L447"></a>f.Name = t.Name();
        <a id="L448"></a>f.Anonymous = true;
    <a id="L449"></a>}
    <a id="L450"></a>if p.pkgPath != nil {
        <a id="L451"></a>f.PkgPath = *p.pkgPath
    <a id="L452"></a>}
    <a id="L453"></a>if p.tag != nil {
        <a id="L454"></a>f.Tag = *p.tag
    <a id="L455"></a>}
    <a id="L456"></a>f.Offset = p.offset;
    <a id="L457"></a>f.Index = []int{i};
    <a id="L458"></a>return;
<a id="L459"></a>}

<a id="L461"></a><span class="comment">// TODO(gri): Should there be an error/bool indicator if the index</span>
<a id="L462"></a><span class="comment">//            is wrong for FieldByIndex?</span>

<a id="L464"></a><span class="comment">// FieldByIndex returns the nested field corresponding to index.</span>
<a id="L465"></a>func (t *StructType) FieldByIndex(index []int) (f StructField) {
    <a id="L466"></a>for i, x := range index {
        <a id="L467"></a>if i &gt; 0 {
            <a id="L468"></a>ft := f.Type;
            <a id="L469"></a>if pt, ok := ft.(*PtrType); ok {
                <a id="L470"></a>ft = pt.Elem()
            <a id="L471"></a>}
            <a id="L472"></a>if st, ok := ft.(*StructType); ok {
                <a id="L473"></a>t = st
            <a id="L474"></a>} else {
                <a id="L475"></a>var f0 StructField;
                <a id="L476"></a>f = f0;
                <a id="L477"></a>return;
            <a id="L478"></a>}
        <a id="L479"></a>}
        <a id="L480"></a>f = t.Field(x);
    <a id="L481"></a>}
    <a id="L482"></a>return;
<a id="L483"></a>}

<a id="L485"></a>const inf = 1 &lt;&lt; 30 <span class="comment">// infinity - no struct has that many nesting levels</span>

<a id="L487"></a>func (t *StructType) fieldByName(name string, mark map[*StructType]bool, depth int) (ff StructField, fd int) {
    <a id="L488"></a>fd = inf; <span class="comment">// field depth</span>

    <a id="L490"></a>if _, marked := mark[t]; marked {
        <a id="L491"></a><span class="comment">// Struct already seen.</span>
        <a id="L492"></a>return
    <a id="L493"></a>}
    <a id="L494"></a>mark[t] = true;

    <a id="L496"></a>var fi int; <span class="comment">// field index</span>
    <a id="L497"></a>n := 0;     <span class="comment">// number of matching fields at depth fd</span>
<a id="L498"></a>L:  for i, _ := range t.fields {
        <a id="L499"></a>f := t.Field(i);
        <a id="L500"></a>d := inf;
        <a id="L501"></a>switch {
        <a id="L502"></a>case f.Name == name:
            <a id="L503"></a><span class="comment">// Matching top-level field.</span>
            <a id="L504"></a>d = depth
        <a id="L505"></a>case f.Anonymous:
            <a id="L506"></a>ft := f.Type;
            <a id="L507"></a>if pt, ok := ft.(*PtrType); ok {
                <a id="L508"></a>ft = pt.Elem()
            <a id="L509"></a>}
            <a id="L510"></a>switch {
            <a id="L511"></a>case ft.Name() == name:
                <a id="L512"></a><span class="comment">// Matching anonymous top-level field.</span>
                <a id="L513"></a>d = depth
            <a id="L514"></a>case fd &gt; depth:
                <a id="L515"></a><span class="comment">// No top-level field yet; look inside nested structs.</span>
                <a id="L516"></a>if st, ok := ft.(*StructType); ok {
                    <a id="L517"></a>f, d = st.fieldByName(name, mark, depth+1)
                <a id="L518"></a>}
            <a id="L519"></a>}
        <a id="L520"></a>}

        <a id="L522"></a>switch {
        <a id="L523"></a>case d &lt; fd:
            <a id="L524"></a><span class="comment">// Found field at shallower depth.</span>
            <a id="L525"></a>ff, fi, fd = f, i, d;
            <a id="L526"></a>n = 1;
        <a id="L527"></a>case d == fd:
            <a id="L528"></a><span class="comment">// More than one matching field at the same depth (or d, fd == inf).</span>
            <a id="L529"></a><span class="comment">// Same as no field found at this depth.</span>
            <a id="L530"></a>n++;
            <a id="L531"></a>if d == depth {
                <a id="L532"></a><span class="comment">// Impossible to find a field at lower depth.</span>
                <a id="L533"></a>break L
            <a id="L534"></a>}
        <a id="L535"></a>}
    <a id="L536"></a>}

    <a id="L538"></a>if n == 1 {
        <a id="L539"></a><span class="comment">// Found matching field.</span>
        <a id="L540"></a>if len(ff.Index) &lt;= depth {
            <a id="L541"></a>ff.Index = make([]int, depth+1)
        <a id="L542"></a>}
        <a id="L543"></a>ff.Index[depth] = fi;
    <a id="L544"></a>} else {
        <a id="L545"></a><span class="comment">// None or more than one matching field found.</span>
        <a id="L546"></a>fd = inf
    <a id="L547"></a>}

    <a id="L549"></a>mark[t] = false, false;
    <a id="L550"></a>return;
<a id="L551"></a>}

<a id="L553"></a><span class="comment">// FieldByName returns the struct field with the given name</span>
<a id="L554"></a><span class="comment">// and a boolean to indicate if the field was found.</span>
<a id="L555"></a>func (t *StructType) FieldByName(name string) (f StructField, present bool) {
    <a id="L556"></a>if ff, fd := t.fieldByName(name, make(map[*StructType]bool), 0); fd &lt; inf {
        <a id="L557"></a>ff.Index = ff.Index[0 : fd+1];
        <a id="L558"></a>f, present = ff, true;
    <a id="L559"></a>}
    <a id="L560"></a>return;
<a id="L561"></a>}

<a id="L563"></a><span class="comment">// NumField returns the number of struct fields.</span>
<a id="L564"></a>func (t *StructType) NumField() int { return len(t.fields) }

<a id="L566"></a><span class="comment">// Convert runtime type to reflect type.</span>
<a id="L567"></a><span class="comment">// Same memory layouts, different method sets.</span>
<a id="L568"></a>func toType(i interface{}) Type {
    <a id="L569"></a>switch v := i.(type) {
    <a id="L570"></a>case *runtime.BoolType:
        <a id="L571"></a>return (*BoolType)(unsafe.Pointer(v))
    <a id="L572"></a>case *runtime.DotDotDotType:
        <a id="L573"></a>return (*DotDotDotType)(unsafe.Pointer(v))
    <a id="L574"></a>case *runtime.FloatType:
        <a id="L575"></a>return (*FloatType)(unsafe.Pointer(v))
    <a id="L576"></a>case *runtime.Float32Type:
        <a id="L577"></a>return (*Float32Type)(unsafe.Pointer(v))
    <a id="L578"></a>case *runtime.Float64Type:
        <a id="L579"></a>return (*Float64Type)(unsafe.Pointer(v))
    <a id="L580"></a>case *runtime.IntType:
        <a id="L581"></a>return (*IntType)(unsafe.Pointer(v))
    <a id="L582"></a>case *runtime.Int8Type:
        <a id="L583"></a>return (*Int8Type)(unsafe.Pointer(v))
    <a id="L584"></a>case *runtime.Int16Type:
        <a id="L585"></a>return (*Int16Type)(unsafe.Pointer(v))
    <a id="L586"></a>case *runtime.Int32Type:
        <a id="L587"></a>return (*Int32Type)(unsafe.Pointer(v))
    <a id="L588"></a>case *runtime.Int64Type:
        <a id="L589"></a>return (*Int64Type)(unsafe.Pointer(v))
    <a id="L590"></a>case *runtime.StringType:
        <a id="L591"></a>return (*StringType)(unsafe.Pointer(v))
    <a id="L592"></a>case *runtime.UintType:
        <a id="L593"></a>return (*UintType)(unsafe.Pointer(v))
    <a id="L594"></a>case *runtime.Uint8Type:
        <a id="L595"></a>return (*Uint8Type)(unsafe.Pointer(v))
    <a id="L596"></a>case *runtime.Uint16Type:
        <a id="L597"></a>return (*Uint16Type)(unsafe.Pointer(v))
    <a id="L598"></a>case *runtime.Uint32Type:
        <a id="L599"></a>return (*Uint32Type)(unsafe.Pointer(v))
    <a id="L600"></a>case *runtime.Uint64Type:
        <a id="L601"></a>return (*Uint64Type)(unsafe.Pointer(v))
    <a id="L602"></a>case *runtime.UintptrType:
        <a id="L603"></a>return (*UintptrType)(unsafe.Pointer(v))
    <a id="L604"></a>case *runtime.UnsafePointerType:
        <a id="L605"></a>return (*UnsafePointerType)(unsafe.Pointer(v))
    <a id="L606"></a>case *runtime.ArrayType:
        <a id="L607"></a>return (*ArrayType)(unsafe.Pointer(v))
    <a id="L608"></a>case *runtime.ChanType:
        <a id="L609"></a>return (*ChanType)(unsafe.Pointer(v))
    <a id="L610"></a>case *runtime.FuncType:
        <a id="L611"></a>return (*FuncType)(unsafe.Pointer(v))
    <a id="L612"></a>case *runtime.InterfaceType:
        <a id="L613"></a>return (*InterfaceType)(unsafe.Pointer(v))
    <a id="L614"></a>case *runtime.MapType:
        <a id="L615"></a>return (*MapType)(unsafe.Pointer(v))
    <a id="L616"></a>case *runtime.PtrType:
        <a id="L617"></a>return (*PtrType)(unsafe.Pointer(v))
    <a id="L618"></a>case *runtime.SliceType:
        <a id="L619"></a>return (*SliceType)(unsafe.Pointer(v))
    <a id="L620"></a>case *runtime.StructType:
        <a id="L621"></a>return (*StructType)(unsafe.Pointer(v))
    <a id="L622"></a>}
    <a id="L623"></a>panicln(&#34;toType&#34;, i);
<a id="L624"></a>}

<a id="L626"></a><span class="comment">// ArrayOrSliceType is the common interface implemented</span>
<a id="L627"></a><span class="comment">// by both ArrayType and SliceType.</span>
<a id="L628"></a>type ArrayOrSliceType interface {
    <a id="L629"></a>Type;
    <a id="L630"></a>Elem() Type;
<a id="L631"></a>}

<a id="L633"></a><span class="comment">// Typeof returns the reflection Type of the value in the interface{}.</span>
<a id="L634"></a>func Typeof(i interface{}) Type { return toType(unsafe.Typeof(i)) }
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
