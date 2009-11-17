<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/bridge.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/bridge.go</h1>

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
    <a id="L8"></a>&#34;log&#34;;
    <a id="L9"></a>&#34;go/token&#34;;
    <a id="L10"></a>&#34;reflect&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">/*</span>
<a id="L14"></a><span class="comment"> * Type bridging</span>
<a id="L15"></a><span class="comment"> */</span>

<a id="L17"></a>var (
    <a id="L18"></a>evalTypes   = make(map[reflect.Type]Type);
    <a id="L19"></a>nativeTypes = make(map[Type]reflect.Type);
<a id="L20"></a>)

<a id="L22"></a><span class="comment">// TypeFromNative converts a regular Go type into a the corresponding</span>
<a id="L23"></a><span class="comment">// interpreter Type.</span>
<a id="L24"></a>func TypeFromNative(t reflect.Type) Type {
    <a id="L25"></a>if et, ok := evalTypes[t]; ok {
        <a id="L26"></a>return et
    <a id="L27"></a>}

    <a id="L29"></a>var nt *NamedType;
    <a id="L30"></a>if t.Name() != &#34;&#34; {
        <a id="L31"></a>name := t.PkgPath() + &#34;·&#34; + t.Name();
        <a id="L32"></a>nt = &amp;NamedType{token.Position{}, name, nil, true, make(map[string]Method)};
        <a id="L33"></a>evalTypes[t] = nt;
    <a id="L34"></a>}

    <a id="L36"></a>var et Type;
    <a id="L37"></a>switch t := t.(type) {
    <a id="L38"></a>case *reflect.BoolType:
        <a id="L39"></a>et = BoolType
    <a id="L40"></a>case *reflect.Float32Type:
        <a id="L41"></a>et = Float32Type
    <a id="L42"></a>case *reflect.Float64Type:
        <a id="L43"></a>et = Float64Type
    <a id="L44"></a>case *reflect.FloatType:
        <a id="L45"></a>et = FloatType
    <a id="L46"></a>case *reflect.Int16Type:
        <a id="L47"></a>et = Int16Type
    <a id="L48"></a>case *reflect.Int32Type:
        <a id="L49"></a>et = Int32Type
    <a id="L50"></a>case *reflect.Int64Type:
        <a id="L51"></a>et = Int64Type
    <a id="L52"></a>case *reflect.Int8Type:
        <a id="L53"></a>et = Int8Type
    <a id="L54"></a>case *reflect.IntType:
        <a id="L55"></a>et = IntType
    <a id="L56"></a>case *reflect.StringType:
        <a id="L57"></a>et = StringType
    <a id="L58"></a>case *reflect.Uint16Type:
        <a id="L59"></a>et = Uint16Type
    <a id="L60"></a>case *reflect.Uint32Type:
        <a id="L61"></a>et = Uint32Type
    <a id="L62"></a>case *reflect.Uint64Type:
        <a id="L63"></a>et = Uint64Type
    <a id="L64"></a>case *reflect.Uint8Type:
        <a id="L65"></a>et = Uint8Type
    <a id="L66"></a>case *reflect.UintType:
        <a id="L67"></a>et = UintType
    <a id="L68"></a>case *reflect.UintptrType:
        <a id="L69"></a>et = UintptrType

    <a id="L71"></a>case *reflect.ArrayType:
        <a id="L72"></a>et = NewArrayType(int64(t.Len()), TypeFromNative(t.Elem()))
    <a id="L73"></a>case *reflect.ChanType:
        <a id="L74"></a>log.Crashf(&#34;%T not implemented&#34;, t)
    <a id="L75"></a>case *reflect.FuncType:
        <a id="L76"></a>nin := t.NumIn();
        <a id="L77"></a><span class="comment">// Variadic functions have DotDotDotType at the end</span>
        <a id="L78"></a>varidic := false;
        <a id="L79"></a>if nin &gt; 0 {
            <a id="L80"></a>if _, ok := t.In(nin - 1).(*reflect.DotDotDotType); ok {
                <a id="L81"></a>varidic = true;
                <a id="L82"></a>nin--;
            <a id="L83"></a>}
        <a id="L84"></a>}
        <a id="L85"></a>in := make([]Type, nin);
        <a id="L86"></a>for i := range in {
            <a id="L87"></a>in[i] = TypeFromNative(t.In(i))
        <a id="L88"></a>}
        <a id="L89"></a>out := make([]Type, t.NumOut());
        <a id="L90"></a>for i := range out {
            <a id="L91"></a>out[i] = TypeFromNative(t.Out(i))
        <a id="L92"></a>}
        <a id="L93"></a>et = NewFuncType(in, varidic, out);
    <a id="L94"></a>case *reflect.InterfaceType:
        <a id="L95"></a>log.Crashf(&#34;%T not implemented&#34;, t)
    <a id="L96"></a>case *reflect.MapType:
        <a id="L97"></a>log.Crashf(&#34;%T not implemented&#34;, t)
    <a id="L98"></a>case *reflect.PtrType:
        <a id="L99"></a>et = NewPtrType(TypeFromNative(t.Elem()))
    <a id="L100"></a>case *reflect.SliceType:
        <a id="L101"></a>et = NewSliceType(TypeFromNative(t.Elem()))
    <a id="L102"></a>case *reflect.StructType:
        <a id="L103"></a>n := t.NumField();
        <a id="L104"></a>fields := make([]StructField, n);
        <a id="L105"></a>for i := 0; i &lt; n; i++ {
            <a id="L106"></a>sf := t.Field(i);
            <a id="L107"></a><span class="comment">// TODO(austin) What to do about private fields?</span>
            <a id="L108"></a>fields[i].Name = sf.Name;
            <a id="L109"></a>fields[i].Type = TypeFromNative(sf.Type);
            <a id="L110"></a>fields[i].Anonymous = sf.Anonymous;
        <a id="L111"></a>}
        <a id="L112"></a>et = NewStructType(fields);
    <a id="L113"></a>case *reflect.UnsafePointerType:
        <a id="L114"></a>log.Crashf(&#34;%T not implemented&#34;, t)
    <a id="L115"></a>default:
        <a id="L116"></a>log.Crashf(&#34;unexpected reflect.Type: %T&#34;, t)
    <a id="L117"></a>}

    <a id="L119"></a>if nt != nil {
        <a id="L120"></a>if _, ok := et.(*NamedType); !ok {
            <a id="L121"></a>nt.Complete(et);
            <a id="L122"></a>et = nt;
        <a id="L123"></a>}
    <a id="L124"></a>}

    <a id="L126"></a>nativeTypes[et] = t;
    <a id="L127"></a>evalTypes[t] = et;

    <a id="L129"></a>return et;
<a id="L130"></a>}

<a id="L132"></a><span class="comment">// TypeOfNative returns the interpreter Type of a regular Go value.</span>
<a id="L133"></a>func TypeOfNative(v interface{}) Type { return TypeFromNative(reflect.Typeof(v)) }

<a id="L135"></a><span class="comment">/*</span>
<a id="L136"></a><span class="comment"> * Function bridging</span>
<a id="L137"></a><span class="comment"> */</span>

<a id="L139"></a>type nativeFunc struct {
    <a id="L140"></a>fn      func(*Thread, []Value, []Value);
    <a id="L141"></a>in, out int;
<a id="L142"></a>}

<a id="L144"></a>func (f *nativeFunc) NewFrame() *Frame {
    <a id="L145"></a>vars := make([]Value, f.in+f.out);
    <a id="L146"></a>return &amp;Frame{nil, vars};
<a id="L147"></a>}

<a id="L149"></a>func (f *nativeFunc) Call(t *Thread) { f.fn(t, t.f.Vars[0:f.in], t.f.Vars[f.in:f.in+f.out]) }

<a id="L151"></a><span class="comment">// FuncFromNative creates an interpreter function from a native</span>
<a id="L152"></a><span class="comment">// function that takes its in and out arguments as slices of</span>
<a id="L153"></a><span class="comment">// interpreter Value&#39;s.  While somewhat inconvenient, this avoids</span>
<a id="L154"></a><span class="comment">// value marshalling.</span>
<a id="L155"></a>func FuncFromNative(fn func(*Thread, []Value, []Value), t *FuncType) FuncValue {
    <a id="L156"></a>return &amp;funcV{&amp;nativeFunc{fn, len(t.In), len(t.Out)}}
<a id="L157"></a>}

<a id="L159"></a><span class="comment">// FuncFromNativeTyped is like FuncFromNative, but constructs the</span>
<a id="L160"></a><span class="comment">// function type from a function pointer using reflection.  Typically,</span>
<a id="L161"></a><span class="comment">// the type will be given as a nil pointer to a function with the</span>
<a id="L162"></a><span class="comment">// desired signature.</span>
<a id="L163"></a>func FuncFromNativeTyped(fn func(*Thread, []Value, []Value), t interface{}) (*FuncType, FuncValue) {
    <a id="L164"></a>ft := TypeOfNative(t).(*FuncType);
    <a id="L165"></a>return ft, FuncFromNative(fn, ft);
<a id="L166"></a>}
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
