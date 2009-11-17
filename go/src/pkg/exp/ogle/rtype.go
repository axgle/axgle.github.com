<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/rtype.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/rtype.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package ogle

<a id="L7"></a>import (
    <a id="L8"></a>&#34;debug/proc&#34;;
    <a id="L9"></a>&#34;exp/eval&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;log&#34;;
<a id="L12"></a>)

<a id="L14"></a>const debugParseRemoteType = false

<a id="L16"></a><span class="comment">// A remoteType is the local representation of a type in a remote process.</span>
<a id="L17"></a>type remoteType struct {
    <a id="L18"></a>eval.Type;
    <a id="L19"></a><span class="comment">// The size of values of this type in bytes.</span>
    <a id="L20"></a>size int;
    <a id="L21"></a><span class="comment">// The field alignment of this type.  Only used for</span>
    <a id="L22"></a><span class="comment">// manually-constructed types.</span>
    <a id="L23"></a>fieldAlign int;
    <a id="L24"></a><span class="comment">// The maker function to turn a remote address of a value of</span>
    <a id="L25"></a><span class="comment">// this type into an interpreter Value.</span>
    <a id="L26"></a>mk  maker;
<a id="L27"></a>}

<a id="L29"></a>var manualTypes = make(map[Arch]map[eval.Type]*remoteType)

<a id="L31"></a><span class="comment">// newManualType constructs a remote type from an interpreter Type</span>
<a id="L32"></a><span class="comment">// using the size and alignment properties of the given architecture.</span>
<a id="L33"></a><span class="comment">// Most types are parsed directly out of the remote process, but to do</span>
<a id="L34"></a><span class="comment">// so we need to layout the structures that describe those types ourselves.</span>
<a id="L35"></a>func newManualType(t eval.Type, arch Arch) *remoteType {
    <a id="L36"></a>if nt, ok := t.(*eval.NamedType); ok {
        <a id="L37"></a>t = nt.Def
    <a id="L38"></a>}

    <a id="L40"></a><span class="comment">// Get the type map for this architecture</span>
    <a id="L41"></a>typeMap, _ := manualTypes[arch];
    <a id="L42"></a>if typeMap == nil {
        <a id="L43"></a>typeMap = make(map[eval.Type]*remoteType);
        <a id="L44"></a>manualTypes[arch] = typeMap;

        <a id="L46"></a><span class="comment">// Construct basic types for this architecture</span>
        <a id="L47"></a>basicType := func(t eval.Type, mk maker, size int, fieldAlign int) {
            <a id="L48"></a>t = t.(*eval.NamedType).Def;
            <a id="L49"></a>if fieldAlign == 0 {
                <a id="L50"></a>fieldAlign = size
            <a id="L51"></a>}
            <a id="L52"></a>typeMap[t] = &amp;remoteType{t, size, fieldAlign, mk};
        <a id="L53"></a>};
        <a id="L54"></a>basicType(eval.Uint8Type, mkUint8, 1, 0);
        <a id="L55"></a>basicType(eval.Uint32Type, mkUint32, 4, 0);
        <a id="L56"></a>basicType(eval.UintptrType, mkUintptr, arch.PtrSize(), 0);
        <a id="L57"></a>basicType(eval.Int16Type, mkInt16, 2, 0);
        <a id="L58"></a>basicType(eval.Int32Type, mkInt32, 4, 0);
        <a id="L59"></a>basicType(eval.IntType, mkInt, arch.IntSize(), 0);
        <a id="L60"></a>basicType(eval.StringType, mkString, arch.PtrSize()+arch.IntSize(), arch.PtrSize());
    <a id="L61"></a>}

    <a id="L63"></a>if rt, ok := typeMap[t]; ok {
        <a id="L64"></a>return rt
    <a id="L65"></a>}

    <a id="L67"></a>var rt *remoteType;
    <a id="L68"></a>switch t := t.(type) {
    <a id="L69"></a>case *eval.PtrType:
        <a id="L70"></a>var elem *remoteType;
        <a id="L71"></a>mk := func(r remote) eval.Value { return remotePtr{r, elem} };
        <a id="L72"></a>rt = &amp;remoteType{t, arch.PtrSize(), arch.PtrSize(), mk};
        <a id="L73"></a><span class="comment">// Construct the element type after registering the</span>
        <a id="L74"></a><span class="comment">// type to break cycles.</span>
        <a id="L75"></a>typeMap[eval.Type(t)] = rt;
        <a id="L76"></a>elem = newManualType(t.Elem, arch);

    <a id="L78"></a>case *eval.ArrayType:
        <a id="L79"></a>elem := newManualType(t.Elem, arch);
        <a id="L80"></a>mk := func(r remote) eval.Value { return remoteArray{r, t.Len, elem} };
        <a id="L81"></a>rt = &amp;remoteType{t, elem.size * int(t.Len), elem.fieldAlign, mk};

    <a id="L83"></a>case *eval.SliceType:
        <a id="L84"></a>elem := newManualType(t.Elem, arch);
        <a id="L85"></a>mk := func(r remote) eval.Value { return remoteSlice{r, elem} };
        <a id="L86"></a>rt = &amp;remoteType{t, arch.PtrSize() + 2*arch.IntSize(), arch.PtrSize(), mk};

    <a id="L88"></a>case *eval.StructType:
        <a id="L89"></a>layout := make([]remoteStructField, len(t.Elems));
        <a id="L90"></a>offset := 0;
        <a id="L91"></a>fieldAlign := 0;
        <a id="L92"></a>for i, f := range t.Elems {
            <a id="L93"></a>elem := newManualType(f.Type, arch);
            <a id="L94"></a>if fieldAlign == 0 {
                <a id="L95"></a>fieldAlign = elem.fieldAlign
            <a id="L96"></a>}
            <a id="L97"></a>offset = arch.Align(offset, elem.fieldAlign);
            <a id="L98"></a>layout[i].offset = offset;
            <a id="L99"></a>layout[i].fieldType = elem;
            <a id="L100"></a>offset += elem.size;
        <a id="L101"></a>}
        <a id="L102"></a>mk := func(r remote) eval.Value { return remoteStruct{r, layout} };
        <a id="L103"></a>rt = &amp;remoteType{t, offset, fieldAlign, mk};

    <a id="L105"></a>default:
        <a id="L106"></a>log.Crashf(&#34;cannot manually construct type %T&#34;, t)
    <a id="L107"></a>}

    <a id="L109"></a>typeMap[t] = rt;
    <a id="L110"></a>return rt;
<a id="L111"></a>}

<a id="L113"></a>var prtIndent = &#34;&#34;

<a id="L115"></a><span class="comment">// parseRemoteType parses a Type structure in a remote process to</span>
<a id="L116"></a><span class="comment">// construct the corresponding interpreter type and remote type.</span>
<a id="L117"></a>func parseRemoteType(a aborter, rs remoteStruct) *remoteType {
    <a id="L118"></a>addr := rs.addr().base;
    <a id="L119"></a>p := rs.addr().p;

    <a id="L121"></a><span class="comment">// We deal with circular types by discovering cycles at</span>
    <a id="L122"></a><span class="comment">// NamedTypes.  If a type cycles back to something other than</span>
    <a id="L123"></a><span class="comment">// a named type, we&#39;re guaranteed that there will be a named</span>
    <a id="L124"></a><span class="comment">// type somewhere in that cycle.  Thus, we continue down,</span>
    <a id="L125"></a><span class="comment">// re-parsing types until we reach the named type in the</span>
    <a id="L126"></a><span class="comment">// cycle.  In order to still create one remoteType per remote</span>
    <a id="L127"></a><span class="comment">// type, we insert an empty remoteType in the type map the</span>
    <a id="L128"></a><span class="comment">// first time we encounter the type and re-use that structure</span>
    <a id="L129"></a><span class="comment">// the second time we encounter it.</span>

    <a id="L131"></a>rt, ok := p.types[addr];
    <a id="L132"></a>if ok &amp;&amp; rt.Type != nil {
        <a id="L133"></a>return rt
    <a id="L134"></a>} else if !ok {
        <a id="L135"></a>rt = &amp;remoteType{};
        <a id="L136"></a>p.types[addr] = rt;
    <a id="L137"></a>}

    <a id="L139"></a>if debugParseRemoteType {
        <a id="L140"></a>sym := p.syms.SymByAddr(uint64(addr));
        <a id="L141"></a>name := &#34;&lt;unknown&gt;&#34;;
        <a id="L142"></a>if sym != nil {
            <a id="L143"></a>name = sym.Name
        <a id="L144"></a>}
        <a id="L145"></a>log.Stderrf(&#34;%sParsing type at %#x (%s)&#34;, prtIndent, addr, name);
        <a id="L146"></a>prtIndent += &#34; &#34;;
        <a id="L147"></a>defer func() { prtIndent = prtIndent[0 : len(prtIndent)-1] }();
    <a id="L148"></a>}

    <a id="L150"></a><span class="comment">// Get Type header</span>
    <a id="L151"></a>itype := proc.Word(rs.field(p.f.Type.Typ).(remoteUint).aGet(a));
    <a id="L152"></a>typ := rs.field(p.f.Type.Ptr).(remotePtr).aGet(a).(remoteStruct);

    <a id="L154"></a><span class="comment">// Is this a named type?</span>
    <a id="L155"></a>var nt *eval.NamedType;
    <a id="L156"></a>uncommon := typ.field(p.f.CommonType.UncommonType).(remotePtr).aGet(a);
    <a id="L157"></a>if uncommon != nil {
        <a id="L158"></a>name := uncommon.(remoteStruct).field(p.f.UncommonType.Name).(remotePtr).aGet(a);
        <a id="L159"></a>if name != nil {
            <a id="L160"></a><span class="comment">// TODO(austin) Declare type in appropriate remote package</span>
            <a id="L161"></a>nt = eval.NewNamedType(name.(remoteString).aGet(a));
            <a id="L162"></a>rt.Type = nt;
        <a id="L163"></a>}
    <a id="L164"></a>}

    <a id="L166"></a><span class="comment">// Create type</span>
    <a id="L167"></a>var t eval.Type;
    <a id="L168"></a>var mk maker;
    <a id="L169"></a>switch itype {
    <a id="L170"></a>case p.runtime.PBoolType:
        <a id="L171"></a>t = eval.BoolType;
        <a id="L172"></a>mk = mkBool;
    <a id="L173"></a>case p.runtime.PUint8Type:
        <a id="L174"></a>t = eval.Uint8Type;
        <a id="L175"></a>mk = mkUint8;
    <a id="L176"></a>case p.runtime.PUint16Type:
        <a id="L177"></a>t = eval.Uint16Type;
        <a id="L178"></a>mk = mkUint16;
    <a id="L179"></a>case p.runtime.PUint32Type:
        <a id="L180"></a>t = eval.Uint32Type;
        <a id="L181"></a>mk = mkUint32;
    <a id="L182"></a>case p.runtime.PUint64Type:
        <a id="L183"></a>t = eval.Uint64Type;
        <a id="L184"></a>mk = mkUint64;
    <a id="L185"></a>case p.runtime.PUintType:
        <a id="L186"></a>t = eval.UintType;
        <a id="L187"></a>mk = mkUint;
    <a id="L188"></a>case p.runtime.PUintptrType:
        <a id="L189"></a>t = eval.UintptrType;
        <a id="L190"></a>mk = mkUintptr;
    <a id="L191"></a>case p.runtime.PInt8Type:
        <a id="L192"></a>t = eval.Int8Type;
        <a id="L193"></a>mk = mkInt8;
    <a id="L194"></a>case p.runtime.PInt16Type:
        <a id="L195"></a>t = eval.Int16Type;
        <a id="L196"></a>mk = mkInt16;
    <a id="L197"></a>case p.runtime.PInt32Type:
        <a id="L198"></a>t = eval.Int32Type;
        <a id="L199"></a>mk = mkInt32;
    <a id="L200"></a>case p.runtime.PInt64Type:
        <a id="L201"></a>t = eval.Int64Type;
        <a id="L202"></a>mk = mkInt64;
    <a id="L203"></a>case p.runtime.PIntType:
        <a id="L204"></a>t = eval.IntType;
        <a id="L205"></a>mk = mkInt;
    <a id="L206"></a>case p.runtime.PFloat32Type:
        <a id="L207"></a>t = eval.Float32Type;
        <a id="L208"></a>mk = mkFloat32;
    <a id="L209"></a>case p.runtime.PFloat64Type:
        <a id="L210"></a>t = eval.Float64Type;
        <a id="L211"></a>mk = mkFloat64;
    <a id="L212"></a>case p.runtime.PFloatType:
        <a id="L213"></a>t = eval.FloatType;
        <a id="L214"></a>mk = mkFloat;
    <a id="L215"></a>case p.runtime.PStringType:
        <a id="L216"></a>t = eval.StringType;
        <a id="L217"></a>mk = mkString;

    <a id="L219"></a>case p.runtime.PArrayType:
        <a id="L220"></a><span class="comment">// Cast to an ArrayType</span>
        <a id="L221"></a>typ := p.runtime.ArrayType.mk(typ.addr()).(remoteStruct);
        <a id="L222"></a>len := int64(typ.field(p.f.ArrayType.Len).(remoteUint).aGet(a));
        <a id="L223"></a>elem := parseRemoteType(a, typ.field(p.f.ArrayType.Elem).(remotePtr).aGet(a).(remoteStruct));
        <a id="L224"></a>t = eval.NewArrayType(len, elem.Type);
        <a id="L225"></a>mk = func(r remote) eval.Value { return remoteArray{r, len, elem} };

    <a id="L227"></a>case p.runtime.PStructType:
        <a id="L228"></a><span class="comment">// Cast to a StructType</span>
        <a id="L229"></a>typ := p.runtime.StructType.mk(typ.addr()).(remoteStruct);
        <a id="L230"></a>fs := typ.field(p.f.StructType.Fields).(remoteSlice).aGet(a);

        <a id="L232"></a>fields := make([]eval.StructField, fs.Len);
        <a id="L233"></a>layout := make([]remoteStructField, fs.Len);
        <a id="L234"></a>for i := range fields {
            <a id="L235"></a>f := fs.Base.(remoteArray).elem(int64(i)).(remoteStruct);
            <a id="L236"></a>elemrs := f.field(p.f.StructField.Typ).(remotePtr).aGet(a).(remoteStruct);
            <a id="L237"></a>elem := parseRemoteType(a, elemrs);
            <a id="L238"></a>fields[i].Type = elem.Type;
            <a id="L239"></a>name := f.field(p.f.StructField.Name).(remotePtr).aGet(a);
            <a id="L240"></a>if name == nil {
                <a id="L241"></a>fields[i].Anonymous = true
            <a id="L242"></a>} else {
                <a id="L243"></a>fields[i].Name = name.(remoteString).aGet(a)
            <a id="L244"></a>}
            <a id="L245"></a>layout[i].offset = int(f.field(p.f.StructField.Offset).(remoteUint).aGet(a));
            <a id="L246"></a>layout[i].fieldType = elem;
        <a id="L247"></a>}

        <a id="L249"></a>t = eval.NewStructType(fields);
        <a id="L250"></a>mk = func(r remote) eval.Value { return remoteStruct{r, layout} };

    <a id="L252"></a>case p.runtime.PPtrType:
        <a id="L253"></a><span class="comment">// Cast to a PtrType</span>
        <a id="L254"></a>typ := p.runtime.PtrType.mk(typ.addr()).(remoteStruct);
        <a id="L255"></a>elem := parseRemoteType(a, typ.field(p.f.PtrType.Elem).(remotePtr).aGet(a).(remoteStruct));
        <a id="L256"></a>t = eval.NewPtrType(elem.Type);
        <a id="L257"></a>mk = func(r remote) eval.Value { return remotePtr{r, elem} };

    <a id="L259"></a>case p.runtime.PSliceType:
        <a id="L260"></a><span class="comment">// Cast to a SliceType</span>
        <a id="L261"></a>typ := p.runtime.SliceType.mk(typ.addr()).(remoteStruct);
        <a id="L262"></a>elem := parseRemoteType(a, typ.field(p.f.SliceType.Elem).(remotePtr).aGet(a).(remoteStruct));
        <a id="L263"></a>t = eval.NewSliceType(elem.Type);
        <a id="L264"></a>mk = func(r remote) eval.Value { return remoteSlice{r, elem} };

    <a id="L266"></a>case p.runtime.PMapType, p.runtime.PChanType, p.runtime.PFuncType, p.runtime.PInterfaceType, p.runtime.PUnsafePointerType, p.runtime.PDotDotDotType:
        <a id="L267"></a><span class="comment">// TODO(austin)</span>
        <a id="L268"></a>t = eval.UintptrType;
        <a id="L269"></a>mk = mkUintptr;

    <a id="L271"></a>default:
        <a id="L272"></a>sym := p.syms.SymByAddr(uint64(itype));
        <a id="L273"></a>name := &#34;&lt;unknown symbol&gt;&#34;;
        <a id="L274"></a>if sym != nil {
            <a id="L275"></a>name = sym.Name
        <a id="L276"></a>}
        <a id="L277"></a>err := fmt.Sprintf(&#34;runtime type at %#x has unexpected type %#x (%s)&#34;, addr, itype, name);
        <a id="L278"></a>a.Abort(FormatError(err));
    <a id="L279"></a>}

    <a id="L281"></a><span class="comment">// Fill in the remote type</span>
    <a id="L282"></a>if nt != nil {
        <a id="L283"></a>nt.Complete(t)
    <a id="L284"></a>} else {
        <a id="L285"></a>rt.Type = t
    <a id="L286"></a>}
    <a id="L287"></a>rt.size = int(typ.field(p.f.CommonType.Size).(remoteUint).aGet(a));
    <a id="L288"></a>rt.mk = mk;

    <a id="L290"></a>return rt;
<a id="L291"></a>}
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
