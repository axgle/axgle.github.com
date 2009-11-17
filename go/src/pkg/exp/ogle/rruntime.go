<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/ogle/rruntime.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/ogle/rruntime.go</h1>

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
    <a id="L10"></a>&#34;reflect&#34;;
<a id="L11"></a>)

<a id="L13"></a><span class="comment">// This file contains remote runtime definitions.  Using reflection,</span>
<a id="L14"></a><span class="comment">// we convert all of these to interpreter types and layout their</span>
<a id="L15"></a><span class="comment">// remote representations using the architecture rules.</span>
<a id="L16"></a><span class="comment">//</span>
<a id="L17"></a><span class="comment">// We could get most of these definitions from our own runtime</span>
<a id="L18"></a><span class="comment">// package; however, some of them differ in convenient ways, some of</span>
<a id="L19"></a><span class="comment">// them are not defined or exported by the runtime, and having our own</span>
<a id="L20"></a><span class="comment">// definitions makes it easy to support multiple remote runtime</span>
<a id="L21"></a><span class="comment">// versions.  This may turn out to be overkill.</span>
<a id="L22"></a><span class="comment">//</span>
<a id="L23"></a><span class="comment">// All of these structures are prefixed with rt1 to indicate the</span>
<a id="L24"></a><span class="comment">// runtime version and to mark them as types used only as templates</span>
<a id="L25"></a><span class="comment">// for remote types.</span>

<a id="L27"></a><span class="comment">/*</span>
<a id="L28"></a><span class="comment"> * Runtime data headers</span>
<a id="L29"></a><span class="comment"> *</span>
<a id="L30"></a><span class="comment"> * See $GOROOT/src/pkg/runtime/runtime.h</span>
<a id="L31"></a><span class="comment"> */</span>

<a id="L33"></a>type rt1String struct {
    <a id="L34"></a>str uintptr;
    <a id="L35"></a>len int;
<a id="L36"></a>}

<a id="L38"></a>type rt1Slice struct {
    <a id="L39"></a>array uintptr;
    <a id="L40"></a>len   int;
    <a id="L41"></a>cap   int;
<a id="L42"></a>}

<a id="L44"></a>type rt1Eface struct {
    <a id="L45"></a>typ uintptr;
    <a id="L46"></a>ptr uintptr;
<a id="L47"></a>}

<a id="L49"></a><span class="comment">/*</span>
<a id="L50"></a><span class="comment"> * Runtime type structures</span>
<a id="L51"></a><span class="comment"> *</span>
<a id="L52"></a><span class="comment"> * See $GOROOT/src/pkg/runtime/type.h and $GOROOT/src/pkg/runtime/type.go</span>
<a id="L53"></a><span class="comment"> */</span>

<a id="L55"></a>type rt1UncommonType struct {
    <a id="L56"></a>name    *string;
    <a id="L57"></a>pkgPath *string;
    <a id="L58"></a><span class="comment">//methods []method;</span>
<a id="L59"></a>}

<a id="L61"></a>type rt1CommonType struct {
    <a id="L62"></a>size                   uintptr;
    <a id="L63"></a>hash                   uint32;
    <a id="L64"></a>alg, align, fieldAlign uint8;
    <a id="L65"></a>string                 *string;
    <a id="L66"></a>uncommonType           *rt1UncommonType;
<a id="L67"></a>}

<a id="L69"></a>type rt1Type struct {
    <a id="L70"></a><span class="comment">// While Type is technically an Eface, treating the</span>
    <a id="L71"></a><span class="comment">// discriminator as an opaque pointer and taking advantage of</span>
    <a id="L72"></a><span class="comment">// the commonType prologue on all Type&#39;s makes type parsing</span>
    <a id="L73"></a><span class="comment">// much simpler.</span>
    <a id="L74"></a>typ uintptr;
    <a id="L75"></a>ptr *rt1CommonType;
<a id="L76"></a>}

<a id="L78"></a>type rt1StructField struct {
    <a id="L79"></a>name    *string;
    <a id="L80"></a>pkgPath *string;
    <a id="L81"></a>typ     *rt1Type;
    <a id="L82"></a>tag     *string;
    <a id="L83"></a>offset  uintptr;
<a id="L84"></a>}

<a id="L86"></a>type rt1StructType struct {
    <a id="L87"></a>rt1CommonType;
    <a id="L88"></a>fields []rt1StructField;
<a id="L89"></a>}

<a id="L91"></a>type rt1PtrType struct {
    <a id="L92"></a>rt1CommonType;
    <a id="L93"></a>elem *rt1Type;
<a id="L94"></a>}

<a id="L96"></a>type rt1SliceType struct {
    <a id="L97"></a>rt1CommonType;
    <a id="L98"></a>elem *rt1Type;
<a id="L99"></a>}

<a id="L101"></a>type rt1ArrayType struct {
    <a id="L102"></a>rt1CommonType;
    <a id="L103"></a>elem *rt1Type;
    <a id="L104"></a>len  uintptr;
<a id="L105"></a>}

<a id="L107"></a><span class="comment">/*</span>
<a id="L108"></a><span class="comment"> * Runtime scheduler structures</span>
<a id="L109"></a><span class="comment"> *</span>
<a id="L110"></a><span class="comment"> * See $GOROOT/src/pkg/runtime/runtime.h</span>
<a id="L111"></a><span class="comment"> */</span>

<a id="L113"></a><span class="comment">// Fields beginning with _ are only for padding</span>

<a id="L115"></a>type rt1Stktop struct {
    <a id="L116"></a>stackguard uintptr;
    <a id="L117"></a>stackbase  *rt1Stktop;
    <a id="L118"></a>gobuf      rt1Gobuf;
    <a id="L119"></a>_args      uint32;
    <a id="L120"></a>_fp        uintptr;
<a id="L121"></a>}

<a id="L123"></a>type rt1Gobuf struct {
    <a id="L124"></a>sp  uintptr;
    <a id="L125"></a>pc  uintptr;
    <a id="L126"></a>g   *rt1G;
    <a id="L127"></a>r0  uintptr;
<a id="L128"></a>}

<a id="L130"></a>type rt1G struct {
    <a id="L131"></a>_stackguard uintptr;
    <a id="L132"></a>stackbase   *rt1Stktop;
    <a id="L133"></a>_defer      uintptr;
    <a id="L134"></a>sched       rt1Gobuf;
    <a id="L135"></a>_stack0     uintptr;
    <a id="L136"></a>_entry      uintptr;
    <a id="L137"></a>alllink     *rt1G;
    <a id="L138"></a>_param      uintptr;
    <a id="L139"></a>status      int16;
    <a id="L140"></a><span class="comment">// Incomplete</span>
<a id="L141"></a>}

<a id="L143"></a>var rt1GStatus = runtimeGStatus{
    <a id="L144"></a>Gidle: 0,
    <a id="L145"></a>Grunnable: 1,
    <a id="L146"></a>Grunning: 2,
    <a id="L147"></a>Gsyscall: 3,
    <a id="L148"></a>Gwaiting: 4,
    <a id="L149"></a>Gmoribund: 5,
    <a id="L150"></a>Gdead: 6,
<a id="L151"></a>}

<a id="L153"></a><span class="comment">// runtimeIndexes stores the indexes of fields in the runtime</span>
<a id="L154"></a><span class="comment">// structures.  It is filled in using reflection, so the name of the</span>
<a id="L155"></a><span class="comment">// fields must match the names of the remoteType&#39;s in runtimeValues</span>
<a id="L156"></a><span class="comment">// exactly and the names of the index fields must be the capitalized</span>
<a id="L157"></a><span class="comment">// version of the names of the fields in the runtime structures above.</span>
<a id="L158"></a>type runtimeIndexes struct {
    <a id="L159"></a>String struct {
        <a id="L160"></a>Str, Len int;
    <a id="L161"></a>};
    <a id="L162"></a>Slice struct {
        <a id="L163"></a>Array, Len, Cap int;
    <a id="L164"></a>};
    <a id="L165"></a>Eface struct {
        <a id="L166"></a>Typ, Ptr int;
    <a id="L167"></a>};

    <a id="L169"></a>UncommonType struct {
        <a id="L170"></a>Name, PkgPath int;
    <a id="L171"></a>};
    <a id="L172"></a>CommonType struct {
        <a id="L173"></a>Size, Hash, Alg, Align, FieldAlign, String, UncommonType int;
    <a id="L174"></a>};
    <a id="L175"></a>Type struct {
        <a id="L176"></a>Typ, Ptr int;
    <a id="L177"></a>};
    <a id="L178"></a>StructField struct {
        <a id="L179"></a>Name, PkgPath, Typ, Tag, Offset int;
    <a id="L180"></a>};
    <a id="L181"></a>StructType struct {
        <a id="L182"></a>Fields int;
    <a id="L183"></a>};
    <a id="L184"></a>PtrType struct {
        <a id="L185"></a>Elem int;
    <a id="L186"></a>};
    <a id="L187"></a>SliceType struct {
        <a id="L188"></a>Elem int;
    <a id="L189"></a>};
    <a id="L190"></a>ArrayType struct {
        <a id="L191"></a>Elem, Len int;
    <a id="L192"></a>};

    <a id="L194"></a>Stktop struct {
        <a id="L195"></a>Stackguard, Stackbase, Gobuf int;
    <a id="L196"></a>};
    <a id="L197"></a>Gobuf struct {
        <a id="L198"></a>Sp, Pc, G int;
    <a id="L199"></a>};
    <a id="L200"></a>G   struct {
        <a id="L201"></a>Stackbase, Sched, Status, Alllink int;
    <a id="L202"></a>};
<a id="L203"></a>}

<a id="L205"></a><span class="comment">// Values of G status codes</span>
<a id="L206"></a>type runtimeGStatus struct {
    <a id="L207"></a>Gidle, Grunnable, Grunning, Gsyscall, Gwaiting, Gmoribund, Gdead int64;
<a id="L208"></a>}

<a id="L210"></a><span class="comment">// runtimeValues stores the types and values that correspond to those</span>
<a id="L211"></a><span class="comment">// in the remote runtime package.</span>
<a id="L212"></a>type runtimeValues struct {
    <a id="L213"></a><span class="comment">// Runtime data headers</span>
    <a id="L214"></a>String, Slice, Eface *remoteType;
    <a id="L215"></a><span class="comment">// Runtime type structures</span>
    <a id="L216"></a>Type, CommonType, UncommonType, StructField, StructType, PtrType,
        <a id="L217"></a>ArrayType, SliceType *remoteType;
    <a id="L218"></a><span class="comment">// Runtime scheduler structures</span>
    <a id="L219"></a>Stktop, Gobuf, G *remoteType;
    <a id="L220"></a><span class="comment">// Addresses of *runtime.XType types.  These are the</span>
    <a id="L221"></a><span class="comment">// discriminators on the runtime.Type interface.  We use local</span>
    <a id="L222"></a><span class="comment">// reflection to fill these in from the remote symbol table,</span>
    <a id="L223"></a><span class="comment">// so the names must match the runtime names.</span>
    <a id="L224"></a>PBoolType,
        <a id="L225"></a>PUint8Type, PUint16Type, PUint32Type, PUint64Type, PUintType, PUintptrType,
        <a id="L226"></a>PInt8Type, PInt16Type, PInt32Type, PInt64Type, PIntType,
        <a id="L227"></a>PFloat32Type, PFloat64Type, PFloatType,
        <a id="L228"></a>PArrayType, PStringType, PStructType, PPtrType, PFuncType,
        <a id="L229"></a>PInterfaceType, PSliceType, PMapType, PChanType,
        <a id="L230"></a>PDotDotDotType, PUnsafePointerType proc.Word;
    <a id="L231"></a><span class="comment">// G status values</span>
    <a id="L232"></a>runtimeGStatus;
<a id="L233"></a>}

<a id="L235"></a><span class="comment">// fillRuntimeIndexes fills a runtimeIndexes structure will the field</span>
<a id="L236"></a><span class="comment">// indexes gathered from the remoteTypes recorded in a runtimeValues</span>
<a id="L237"></a><span class="comment">// structure.</span>
<a id="L238"></a>func fillRuntimeIndexes(runtime *runtimeValues, out *runtimeIndexes) {
    <a id="L239"></a>outv := reflect.Indirect(reflect.NewValue(out)).(*reflect.StructValue);
    <a id="L240"></a>outt := outv.Type().(*reflect.StructType);
    <a id="L241"></a>runtimev := reflect.Indirect(reflect.NewValue(runtime)).(*reflect.StructValue);

    <a id="L243"></a><span class="comment">// out contains fields corresponding to each runtime type</span>
    <a id="L244"></a>for i := 0; i &lt; outt.NumField(); i++ {
        <a id="L245"></a><span class="comment">// Find the interpreter type for this runtime type</span>
        <a id="L246"></a>name := outt.Field(i).Name;
        <a id="L247"></a>et := runtimev.FieldByName(name).Interface().(*remoteType).Type.(*eval.StructType);

        <a id="L249"></a><span class="comment">// Get the field indexes of the interpreter struct type</span>
        <a id="L250"></a>indexes := make(map[string]int, len(et.Elems));
        <a id="L251"></a>for j, f := range et.Elems {
            <a id="L252"></a>if f.Anonymous {
                <a id="L253"></a>continue
            <a id="L254"></a>}
            <a id="L255"></a>name := f.Name;
            <a id="L256"></a>if name[0] &gt;= &#39;a&#39; &amp;&amp; name[0] &lt;= &#39;z&#39; {
                <a id="L257"></a>name = string(name[0]+&#39;A&#39;-&#39;a&#39;) + name[1:len(name)]
            <a id="L258"></a>}
            <a id="L259"></a>indexes[name] = j;
        <a id="L260"></a>}

        <a id="L262"></a><span class="comment">// Fill this field of out</span>
        <a id="L263"></a>outStructv := outv.Field(i).(*reflect.StructValue);
        <a id="L264"></a>outStructt := outStructv.Type().(*reflect.StructType);
        <a id="L265"></a>for j := 0; j &lt; outStructt.NumField(); j++ {
            <a id="L266"></a>f := outStructv.Field(j).(*reflect.IntValue);
            <a id="L267"></a>name := outStructt.Field(j).Name;
            <a id="L268"></a>f.Set(indexes[name]);
        <a id="L269"></a>}
    <a id="L270"></a>}
<a id="L271"></a>}
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
