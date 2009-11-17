<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/runtime/type.go</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/runtime/type.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>
<a id="L6"></a><span class="comment"> * Runtime type representation.</span>
<a id="L7"></a><span class="comment"> *</span>
<a id="L8"></a><span class="comment"> * The following files know the exact layout of these</span>
<a id="L9"></a><span class="comment"> * data structures and must be kept in sync with this file:</span>
<a id="L10"></a><span class="comment"> *</span>
<a id="L11"></a><span class="comment"> *	../../cmd/gc/reflect.c</span>
<a id="L12"></a><span class="comment"> *	../reflect/type.go</span>
<a id="L13"></a><span class="comment"> *	type.h</span>
<a id="L14"></a><span class="comment"> */</span>

<a id="L16"></a>package runtime

<a id="L18"></a>import &#34;unsafe&#34;

<a id="L20"></a><span class="comment">// The compiler can only construct empty interface values at</span>
<a id="L21"></a><span class="comment">// compile time; non-empty interface values get created</span>
<a id="L22"></a><span class="comment">// during initialization.  Type is an empty interface</span>
<a id="L23"></a><span class="comment">// so that the compiler can lay out references as data.</span>
<a id="L24"></a>type Type interface{}

<a id="L26"></a><span class="comment">// All types begin with a few common fields needed for</span>
<a id="L27"></a><span class="comment">// the interface runtime.</span>
<a id="L28"></a>type commonType struct {
    <a id="L29"></a>size           uintptr; <span class="comment">// size in bytes</span>
    <a id="L30"></a>hash           uint32;  <span class="comment">// hash of type; avoids computation in hash tables</span>
    <a id="L31"></a>alg            uint8;   <span class="comment">// algorithm for copy+hash+cmp (../runtime/runtime.h:/AMEM)</span>
    <a id="L32"></a>align          uint8;   <span class="comment">// alignment of variable with this type</span>
    <a id="L33"></a>fieldAlign     uint8;   <span class="comment">// alignment of struct field with this type</span>
    <a id="L34"></a>string         *string; <span class="comment">// string form; unnecessary  but undeniably useful</span>
    <a id="L35"></a>*uncommonType;          <span class="comment">// (relatively) uncommon fields</span>
<a id="L36"></a>}

<a id="L38"></a><span class="comment">// Method on non-interface type</span>
<a id="L39"></a>type method struct {
    <a id="L40"></a>hash    uint32;         <span class="comment">// hash of name + pkg + typ</span>
    <a id="L41"></a>name    *string;        <span class="comment">// name of method</span>
    <a id="L42"></a>pkgPath *string;        <span class="comment">// nil for exported Names; otherwise import path</span>
    <a id="L43"></a>typ     *Type;          <span class="comment">// .(*FuncType) underneath</span>
    <a id="L44"></a>ifn     unsafe.Pointer; <span class="comment">// fn used in interface call (one-word receiver)</span>
    <a id="L45"></a>tfn     unsafe.Pointer; <span class="comment">// fn used for normal method call</span>
<a id="L46"></a>}

<a id="L48"></a><span class="comment">// uncommonType is present only for types with names or methods</span>
<a id="L49"></a><span class="comment">// (if T is a named type, the uncommonTypes for T and *T have methods).</span>
<a id="L50"></a><span class="comment">// Using a pointer to this struct reduces the overall size required</span>
<a id="L51"></a><span class="comment">// to describe an unnamed type with no methods.</span>
<a id="L52"></a>type uncommonType struct {
    <a id="L53"></a>name    *string;  <span class="comment">// name of type</span>
    <a id="L54"></a>pkgPath *string;  <span class="comment">// import path; nil for built-in types like int, string</span>
    <a id="L55"></a>methods []method; <span class="comment">// methods associated with type</span>
<a id="L56"></a>}

<a id="L58"></a><span class="comment">// BoolType represents a boolean type.</span>
<a id="L59"></a>type BoolType commonType

<a id="L61"></a><span class="comment">// Float32Type represents a float32 type.</span>
<a id="L62"></a>type Float32Type commonType

<a id="L64"></a><span class="comment">// Float64Type represents a float64 type.</span>
<a id="L65"></a>type Float64Type commonType

<a id="L67"></a><span class="comment">// FloatType represents a float type.</span>
<a id="L68"></a>type FloatType commonType

<a id="L70"></a><span class="comment">// Int16Type represents an int16 type.</span>
<a id="L71"></a>type Int16Type commonType

<a id="L73"></a><span class="comment">// Int32Type represents an int32 type.</span>
<a id="L74"></a>type Int32Type commonType

<a id="L76"></a><span class="comment">// Int64Type represents an int64 type.</span>
<a id="L77"></a>type Int64Type commonType

<a id="L79"></a><span class="comment">// Int8Type represents an int8 type.</span>
<a id="L80"></a>type Int8Type commonType

<a id="L82"></a><span class="comment">// IntType represents an int type.</span>
<a id="L83"></a>type IntType commonType

<a id="L85"></a><span class="comment">// Uint16Type represents a uint16 type.</span>
<a id="L86"></a>type Uint16Type commonType

<a id="L88"></a><span class="comment">// Uint32Type represents a uint32 type.</span>
<a id="L89"></a>type Uint32Type commonType

<a id="L91"></a><span class="comment">// Uint64Type represents a uint64 type.</span>
<a id="L92"></a>type Uint64Type commonType

<a id="L94"></a><span class="comment">// Uint8Type represents a uint8 type.</span>
<a id="L95"></a>type Uint8Type commonType

<a id="L97"></a><span class="comment">// UintType represents a uint type.</span>
<a id="L98"></a>type UintType commonType

<a id="L100"></a><span class="comment">// StringType represents a string type.</span>
<a id="L101"></a>type StringType commonType

<a id="L103"></a><span class="comment">// UintptrType represents a uintptr type.</span>
<a id="L104"></a>type UintptrType commonType

<a id="L106"></a><span class="comment">// DotDotDotType represents the ... that can</span>
<a id="L107"></a><span class="comment">// be used as the type of the final function parameter.</span>
<a id="L108"></a>type DotDotDotType commonType

<a id="L110"></a><span class="comment">// UnsafePointerType represents an unsafe.Pointer type.</span>
<a id="L111"></a>type UnsafePointerType commonType

<a id="L113"></a><span class="comment">// ArrayType represents a fixed array type.</span>
<a id="L114"></a>type ArrayType struct {
    <a id="L115"></a>commonType;
    <a id="L116"></a>elem *Type; <span class="comment">// array element type</span>
    <a id="L117"></a>len  uintptr;
<a id="L118"></a>}

<a id="L120"></a><span class="comment">// SliceType represents a slice type.</span>
<a id="L121"></a>type SliceType struct {
    <a id="L122"></a>commonType;
    <a id="L123"></a>elem *Type; <span class="comment">// slice element type</span>
<a id="L124"></a>}

<a id="L126"></a><span class="comment">// ChanDir represents a channel type&#39;s direction.</span>
<a id="L127"></a>type ChanDir int

<a id="L129"></a>const (
    <a id="L130"></a>RecvDir  ChanDir = 1 &lt;&lt; iota; <span class="comment">// &lt;-chan</span>
    <a id="L131"></a>SendDir;         <span class="comment">// chan&lt;-</span>
    <a id="L132"></a>BothDir          = RecvDir | SendDir; <span class="comment">// chan</span>
<a id="L133"></a>)

<a id="L135"></a><span class="comment">// ChanType represents a channel type.</span>
<a id="L136"></a>type ChanType struct {
    <a id="L137"></a>commonType;
    <a id="L138"></a>elem *Type;   <span class="comment">// channel element type</span>
    <a id="L139"></a>dir  uintptr; <span class="comment">// channel direction (ChanDir)</span>
<a id="L140"></a>}

<a id="L142"></a><span class="comment">// FuncType represents a function type.</span>
<a id="L143"></a>type FuncType struct {
    <a id="L144"></a>commonType;
    <a id="L145"></a>in  []*Type; <span class="comment">// input parameter types</span>
    <a id="L146"></a>out []*Type; <span class="comment">// output parameter types</span>
<a id="L147"></a>}

<a id="L149"></a><span class="comment">// Method on interface type</span>
<a id="L150"></a>type imethod struct {
    <a id="L151"></a>hash    uint32;  <span class="comment">// hash of name + pkg + typ; same hash as method</span>
    <a id="L152"></a>perm    uint32;  <span class="comment">// index of function pointer in interface map</span>
    <a id="L153"></a>name    *string; <span class="comment">// name of method</span>
    <a id="L154"></a>pkgPath *string; <span class="comment">// nil for exported Names; otherwise import path</span>
    <a id="L155"></a>typ     *Type;   <span class="comment">// .(*FuncType) underneath</span>
<a id="L156"></a>}

<a id="L158"></a><span class="comment">// InterfaceType represents an interface type.</span>
<a id="L159"></a>type InterfaceType struct {
    <a id="L160"></a>commonType;
    <a id="L161"></a>methods []imethod; <span class="comment">// sorted by hash</span>
<a id="L162"></a>}

<a id="L164"></a><span class="comment">// MapType represents a map type.</span>
<a id="L165"></a>type MapType struct {
    <a id="L166"></a>commonType;
    <a id="L167"></a>key  *Type; <span class="comment">// map key type</span>
    <a id="L168"></a>elem *Type; <span class="comment">// map element (value) type</span>
<a id="L169"></a>}

<a id="L171"></a><span class="comment">// PtrType represents a pointer type.</span>
<a id="L172"></a>type PtrType struct {
    <a id="L173"></a>commonType;
    <a id="L174"></a>elem *Type; <span class="comment">// pointer element (pointed at) type</span>
<a id="L175"></a>}

<a id="L177"></a><span class="comment">// Struct field</span>
<a id="L178"></a>type structField struct {
    <a id="L179"></a>name    *string; <span class="comment">// nil for embedded fields</span>
    <a id="L180"></a>pkgPath *string; <span class="comment">// nil for exported Names; otherwise import path</span>
    <a id="L181"></a>typ     *Type;   <span class="comment">// type of field</span>
    <a id="L182"></a>tag     *string; <span class="comment">// nil if no tag</span>
    <a id="L183"></a>offset  uintptr; <span class="comment">// byte offset of field within struct</span>
<a id="L184"></a>}

<a id="L186"></a><span class="comment">// StructType represents a struct type.</span>
<a id="L187"></a>type StructType struct {
    <a id="L188"></a>commonType;
    <a id="L189"></a>fields []structField; <span class="comment">// sorted by offset</span>
<a id="L190"></a>}

<a id="L192"></a><span class="comment">/*</span>
<a id="L193"></a><span class="comment"> * Must match iface.c:/Itab and compilers.</span>
<a id="L194"></a><span class="comment"> */</span>
<a id="L195"></a>type Itable struct {
    <a id="L196"></a>Itype  *Type; <span class="comment">// (*tab.inter).(*InterfaceType) is the interface type</span>
    <a id="L197"></a>Type   *Type;
    <a id="L198"></a>link   *Itable;
    <a id="L199"></a>bad    int32;
    <a id="L200"></a>unused int32;
    <a id="L201"></a>Fn     [100000]uintptr; <span class="comment">// bigger than we&#39;ll ever see</span>
<a id="L202"></a>}
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
