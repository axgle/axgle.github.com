<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/gob/type.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/gob/type.go</h1>

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
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;reflect&#34;;
    <a id="L11"></a>&#34;sync&#34;;
<a id="L12"></a>)

<a id="L14"></a>type kind reflect.Type

<a id="L16"></a><span class="comment">// Reflection types are themselves interface values holding structs</span>
<a id="L17"></a><span class="comment">// describing the type.  Each type has a different struct so that struct can</span>
<a id="L18"></a><span class="comment">// be the kind.  For example, if typ is the reflect type for an int8, typ is</span>
<a id="L19"></a><span class="comment">// a pointer to a reflect.Int8Type struct; if typ is the reflect type for a</span>
<a id="L20"></a><span class="comment">// function, typ is a pointer to a reflect.FuncType struct; we use the type</span>
<a id="L21"></a><span class="comment">// of that pointer as the kind.</span>

<a id="L23"></a><span class="comment">// typeKind returns a reflect.Type representing typ&#39;s kind.  The kind is the</span>
<a id="L24"></a><span class="comment">// general kind of type:</span>
<a id="L25"></a><span class="comment">//	int8, int16, int, uint, float, func, chan, struct, and so on.</span>
<a id="L26"></a><span class="comment">// That is, all struct types have the same kind, all func types have the same</span>
<a id="L27"></a><span class="comment">// kind, all int8 types have the same kind, and so on.</span>
<a id="L28"></a>func typeKind(typ reflect.Type) kind { return kind(reflect.Typeof(typ)) }

<a id="L30"></a><span class="comment">// valueKind returns the kind of the value type</span>
<a id="L31"></a><span class="comment">// stored inside the interface v.</span>
<a id="L32"></a>func valueKind(v interface{}) reflect.Type { return typeKind(reflect.Typeof(v)) }

<a id="L34"></a><span class="comment">// A typeId represents a gob Type as an integer that can be passed on the wire.</span>
<a id="L35"></a><span class="comment">// Internally, typeIds are used as keys to a map to recover the underlying type info.</span>
<a id="L36"></a>type typeId int32

<a id="L38"></a>var nextId typeId       <span class="comment">// incremented for each new type we build</span>
<a id="L39"></a>var typeLock sync.Mutex <span class="comment">// set while building a type</span>

<a id="L41"></a>type gobType interface {
    <a id="L42"></a>id() typeId;
    <a id="L43"></a>setId(id typeId);
    <a id="L44"></a>Name() string;
    <a id="L45"></a>String() string;
    <a id="L46"></a>safeString(seen map[typeId]bool) string;
<a id="L47"></a>}

<a id="L49"></a>var types = make(map[reflect.Type]gobType)
<a id="L50"></a>var idToType = make(map[typeId]gobType)

<a id="L52"></a>func setTypeId(typ gobType) {
    <a id="L53"></a>nextId++;
    <a id="L54"></a>typ.setId(nextId);
    <a id="L55"></a>idToType[nextId] = typ;
<a id="L56"></a>}

<a id="L58"></a>func (t typeId) gobType() gobType {
    <a id="L59"></a>if t == 0 {
        <a id="L60"></a>return nil
    <a id="L61"></a>}
    <a id="L62"></a>return idToType[t];
<a id="L63"></a>}

<a id="L65"></a><span class="comment">// String returns the string representation of the type associated with the typeId.</span>
<a id="L66"></a>func (t typeId) String() string { return t.gobType().String() }

<a id="L68"></a><span class="comment">// Name returns the name of the type associated with the typeId.</span>
<a id="L69"></a>func (t typeId) Name() string { return t.gobType().Name() }

<a id="L71"></a><span class="comment">// Common elements of all types.</span>
<a id="L72"></a>type commonType struct {
    <a id="L73"></a>name string;
    <a id="L74"></a>_id  typeId;
<a id="L75"></a>}

<a id="L77"></a>func (t *commonType) id() typeId { return t._id }

<a id="L79"></a>func (t *commonType) setId(id typeId) { t._id = id }

<a id="L81"></a>func (t *commonType) String() string { return t.name }

<a id="L83"></a>func (t *commonType) safeString(seen map[typeId]bool) string {
    <a id="L84"></a>return t.name
<a id="L85"></a>}

<a id="L87"></a>func (t *commonType) Name() string { return t.name }

<a id="L89"></a><span class="comment">// Create and check predefined types</span>
<a id="L90"></a><span class="comment">// The string for tBytes is &#34;bytes&#34; not &#34;[]byte&#34; to signify its specialness.</span>

<a id="L92"></a>var tBool = bootstrapType(&#34;bool&#34;, false, 1)
<a id="L93"></a>var tInt = bootstrapType(&#34;int&#34;, int(0), 2)
<a id="L94"></a>var tUint = bootstrapType(&#34;uint&#34;, uint(0), 3)
<a id="L95"></a>var tFloat = bootstrapType(&#34;float&#34;, float64(0), 4)
<a id="L96"></a>var tBytes = bootstrapType(&#34;bytes&#34;, make([]byte, 0), 5)
<a id="L97"></a>var tString = bootstrapType(&#34;string&#34;, &#34;&#34;, 6)

<a id="L99"></a><span class="comment">// Predefined because it&#39;s needed by the Decoder</span>
<a id="L100"></a>var tWireType = getTypeInfoNoError(reflect.Typeof(wireType{})).id

<a id="L102"></a>func init() {
    <a id="L103"></a>checkId(7, tWireType);
    <a id="L104"></a>checkId(8, getTypeInfoNoError(reflect.Typeof(structType{})).id);
    <a id="L105"></a>checkId(9, getTypeInfoNoError(reflect.Typeof(commonType{})).id);
    <a id="L106"></a>checkId(10, getTypeInfoNoError(reflect.Typeof(fieldType{})).id);
<a id="L107"></a>}

<a id="L109"></a><span class="comment">// Array type</span>
<a id="L110"></a>type arrayType struct {
    <a id="L111"></a>commonType;
    <a id="L112"></a>Elem typeId;
    <a id="L113"></a>Len  int;
<a id="L114"></a>}

<a id="L116"></a>func newArrayType(name string, elem gobType, length int) *arrayType {
    <a id="L117"></a>a := &amp;arrayType{commonType{name: name}, elem.id(), length};
    <a id="L118"></a>setTypeId(a);
    <a id="L119"></a>return a;
<a id="L120"></a>}

<a id="L122"></a>func (a *arrayType) safeString(seen map[typeId]bool) string {
    <a id="L123"></a>if _, ok := seen[a._id]; ok {
        <a id="L124"></a>return a.name
    <a id="L125"></a>}
    <a id="L126"></a>seen[a._id] = true;
    <a id="L127"></a>return fmt.Sprintf(&#34;[%d]%s&#34;, a.Len, a.Elem.gobType().safeString(seen));
<a id="L128"></a>}

<a id="L130"></a>func (a *arrayType) String() string { return a.safeString(make(map[typeId]bool)) }

<a id="L132"></a><span class="comment">// Slice type</span>
<a id="L133"></a>type sliceType struct {
    <a id="L134"></a>commonType;
    <a id="L135"></a>Elem typeId;
<a id="L136"></a>}

<a id="L138"></a>func newSliceType(name string, elem gobType) *sliceType {
    <a id="L139"></a>s := &amp;sliceType{commonType{name: name}, elem.id()};
    <a id="L140"></a>setTypeId(s);
    <a id="L141"></a>return s;
<a id="L142"></a>}

<a id="L144"></a>func (s *sliceType) safeString(seen map[typeId]bool) string {
    <a id="L145"></a>if _, ok := seen[s._id]; ok {
        <a id="L146"></a>return s.name
    <a id="L147"></a>}
    <a id="L148"></a>seen[s._id] = true;
    <a id="L149"></a>return fmt.Sprintf(&#34;[]%s&#34;, s.Elem.gobType().safeString(seen));
<a id="L150"></a>}

<a id="L152"></a>func (s *sliceType) String() string { return s.safeString(make(map[typeId]bool)) }

<a id="L154"></a><span class="comment">// Struct type</span>
<a id="L155"></a>type fieldType struct {
    <a id="L156"></a>name string;
    <a id="L157"></a>id   typeId;
<a id="L158"></a>}

<a id="L160"></a>type structType struct {
    <a id="L161"></a>commonType;
    <a id="L162"></a>field []*fieldType;
<a id="L163"></a>}

<a id="L165"></a>func (s *structType) safeString(seen map[typeId]bool) string {
    <a id="L166"></a>if s == nil {
        <a id="L167"></a>return &#34;&lt;nil&gt;&#34;
    <a id="L168"></a>}
    <a id="L169"></a>if _, ok := seen[s._id]; ok {
        <a id="L170"></a>return s.name
    <a id="L171"></a>}
    <a id="L172"></a>seen[s._id] = true;
    <a id="L173"></a>str := s.name + &#34; = struct { &#34;;
    <a id="L174"></a>for _, f := range s.field {
        <a id="L175"></a>str += fmt.Sprintf(&#34;%s %s; &#34;, f.name, f.id.gobType().safeString(seen))
    <a id="L176"></a>}
    <a id="L177"></a>str += &#34;}&#34;;
    <a id="L178"></a>return str;
<a id="L179"></a>}

<a id="L181"></a>func (s *structType) String() string { return s.safeString(make(map[typeId]bool)) }

<a id="L183"></a>func newStructType(name string) *structType {
    <a id="L184"></a>s := &amp;structType{commonType{name: name}, nil};
    <a id="L185"></a>setTypeId(s);
    <a id="L186"></a>return s;
<a id="L187"></a>}

<a id="L189"></a><span class="comment">// Step through the indirections on a type to discover the base type.</span>
<a id="L190"></a><span class="comment">// Return the number of indirections.</span>
<a id="L191"></a>func indirect(t reflect.Type) (rt reflect.Type, count int) {
    <a id="L192"></a>rt = t;
    <a id="L193"></a>for {
        <a id="L194"></a>pt, ok := rt.(*reflect.PtrType);
        <a id="L195"></a>if !ok {
            <a id="L196"></a>break
        <a id="L197"></a>}
        <a id="L198"></a>rt = pt.Elem();
        <a id="L199"></a>count++;
    <a id="L200"></a>}
    <a id="L201"></a>return;
<a id="L202"></a>}

<a id="L204"></a>func newTypeObject(name string, rt reflect.Type) (gobType, os.Error) {
    <a id="L205"></a>switch t := rt.(type) {
    <a id="L206"></a><span class="comment">// All basic types are easy: they are predefined.</span>
    <a id="L207"></a>case *reflect.BoolType:
        <a id="L208"></a>return tBool.gobType(), nil

    <a id="L210"></a>case *reflect.IntType:
        <a id="L211"></a>return tInt.gobType(), nil
    <a id="L212"></a>case *reflect.Int8Type:
        <a id="L213"></a>return tInt.gobType(), nil
    <a id="L214"></a>case *reflect.Int16Type:
        <a id="L215"></a>return tInt.gobType(), nil
    <a id="L216"></a>case *reflect.Int32Type:
        <a id="L217"></a>return tInt.gobType(), nil
    <a id="L218"></a>case *reflect.Int64Type:
        <a id="L219"></a>return tInt.gobType(), nil

    <a id="L221"></a>case *reflect.UintType:
        <a id="L222"></a>return tUint.gobType(), nil
    <a id="L223"></a>case *reflect.Uint8Type:
        <a id="L224"></a>return tUint.gobType(), nil
    <a id="L225"></a>case *reflect.Uint16Type:
        <a id="L226"></a>return tUint.gobType(), nil
    <a id="L227"></a>case *reflect.Uint32Type:
        <a id="L228"></a>return tUint.gobType(), nil
    <a id="L229"></a>case *reflect.Uint64Type:
        <a id="L230"></a>return tUint.gobType(), nil
    <a id="L231"></a>case *reflect.UintptrType:
        <a id="L232"></a>return tUint.gobType(), nil

    <a id="L234"></a>case *reflect.FloatType:
        <a id="L235"></a>return tFloat.gobType(), nil
    <a id="L236"></a>case *reflect.Float32Type:
        <a id="L237"></a>return tFloat.gobType(), nil
    <a id="L238"></a>case *reflect.Float64Type:
        <a id="L239"></a>return tFloat.gobType(), nil

    <a id="L241"></a>case *reflect.StringType:
        <a id="L242"></a>return tString.gobType(), nil

    <a id="L244"></a>case *reflect.ArrayType:
        <a id="L245"></a>gt, err := getType(&#34;&#34;, t.Elem());
        <a id="L246"></a>if err != nil {
            <a id="L247"></a>return nil, err
        <a id="L248"></a>}
        <a id="L249"></a>return newArrayType(name, gt, t.Len()), nil;

    <a id="L251"></a>case *reflect.SliceType:
        <a id="L252"></a><span class="comment">// []byte == []uint8 is a special case</span>
        <a id="L253"></a>if _, ok := t.Elem().(*reflect.Uint8Type); ok {
            <a id="L254"></a>return tBytes.gobType(), nil
        <a id="L255"></a>}
        <a id="L256"></a>gt, err := getType(t.Elem().Name(), t.Elem());
        <a id="L257"></a>if err != nil {
            <a id="L258"></a>return nil, err
        <a id="L259"></a>}
        <a id="L260"></a>return newSliceType(name, gt), nil;

    <a id="L262"></a>case *reflect.StructType:
        <a id="L263"></a><span class="comment">// Install the struct type itself before the fields so recursive</span>
        <a id="L264"></a><span class="comment">// structures can be constructed safely.</span>
        <a id="L265"></a>strType := newStructType(name);
        <a id="L266"></a>types[rt] = strType;
        <a id="L267"></a>idToType[strType.id()] = strType;
        <a id="L268"></a>field := make([]*fieldType, t.NumField());
        <a id="L269"></a>for i := 0; i &lt; t.NumField(); i++ {
            <a id="L270"></a>f := t.Field(i);
            <a id="L271"></a>typ, _ := indirect(f.Type);
            <a id="L272"></a>tname := typ.Name();
            <a id="L273"></a>if tname == &#34;&#34; {
                <a id="L274"></a>tname = f.Type.String()
            <a id="L275"></a>}
            <a id="L276"></a>gt, err := getType(tname, f.Type);
            <a id="L277"></a>if err != nil {
                <a id="L278"></a>return nil, err
            <a id="L279"></a>}
            <a id="L280"></a>field[i] = &amp;fieldType{f.Name, gt.id()};
        <a id="L281"></a>}
        <a id="L282"></a>strType.field = field;
        <a id="L283"></a>return strType, nil;

    <a id="L285"></a>default:
        <a id="L286"></a>return nil, os.ErrorString(&#34;gob NewTypeObject can&#39;t handle type: &#34; + rt.String())
    <a id="L287"></a>}
    <a id="L288"></a>return nil, nil;
<a id="L289"></a>}

<a id="L291"></a><span class="comment">// getType returns the Gob type describing the given reflect.Type.</span>
<a id="L292"></a><span class="comment">// typeLock must be held.</span>
<a id="L293"></a>func getType(name string, rt reflect.Type) (gobType, os.Error) {
    <a id="L294"></a><span class="comment">// Flatten the data structure by collapsing out pointers</span>
    <a id="L295"></a>for {
        <a id="L296"></a>pt, ok := rt.(*reflect.PtrType);
        <a id="L297"></a>if !ok {
            <a id="L298"></a>break
        <a id="L299"></a>}
        <a id="L300"></a>rt = pt.Elem();
    <a id="L301"></a>}
    <a id="L302"></a>typ, present := types[rt];
    <a id="L303"></a>if present {
        <a id="L304"></a>return typ, nil
    <a id="L305"></a>}
    <a id="L306"></a>typ, err := newTypeObject(name, rt);
    <a id="L307"></a>if err == nil {
        <a id="L308"></a>types[rt] = typ
    <a id="L309"></a>}
    <a id="L310"></a>return typ, err;
<a id="L311"></a>}

<a id="L313"></a>func checkId(want, got typeId) {
    <a id="L314"></a>if want != got {
        <a id="L315"></a>panicln(&#34;bootstrap type wrong id:&#34;, got.Name(), got, &#34;not&#34;, want)
    <a id="L316"></a>}
<a id="L317"></a>}

<a id="L319"></a><span class="comment">// used for building the basic types; called only from init()</span>
<a id="L320"></a>func bootstrapType(name string, e interface{}, expect typeId) typeId {
    <a id="L321"></a>rt := reflect.Typeof(e);
    <a id="L322"></a>_, present := types[rt];
    <a id="L323"></a>if present {
        <a id="L324"></a>panicln(&#34;bootstrap type already present:&#34;, name)
    <a id="L325"></a>}
    <a id="L326"></a>typ := &amp;commonType{name: name};
    <a id="L327"></a>types[rt] = typ;
    <a id="L328"></a>setTypeId(typ);
    <a id="L329"></a>checkId(expect, nextId);
    <a id="L330"></a>return nextId;
<a id="L331"></a>}

<a id="L333"></a><span class="comment">// Representation of the information we send and receive about this type.</span>
<a id="L334"></a><span class="comment">// Each value we send is preceded by its type definition: an encoded int.</span>
<a id="L335"></a><span class="comment">// However, the very first time we send the value, we first send the pair</span>
<a id="L336"></a><span class="comment">// (-id, wireType).</span>
<a id="L337"></a><span class="comment">// For bootstrapping purposes, we assume that the recipient knows how</span>
<a id="L338"></a><span class="comment">// to decode a wireType; it is exactly the wireType struct here, interpreted</span>
<a id="L339"></a><span class="comment">// using the gob rules for sending a structure, except that we assume the</span>
<a id="L340"></a><span class="comment">// ids for wireType and structType are known.  The relevant pieces</span>
<a id="L341"></a><span class="comment">// are built in encode.go&#39;s init() function.</span>

<a id="L343"></a>type wireType struct {
    <a id="L344"></a>s *structType;
<a id="L345"></a>}

<a id="L347"></a>func (w *wireType) name() string {
    <a id="L348"></a><span class="comment">// generalize once we can have non-struct types on the wire.</span>
    <a id="L349"></a>return w.s.name
<a id="L350"></a>}

<a id="L352"></a>type typeInfo struct {
    <a id="L353"></a>id      typeId;
    <a id="L354"></a>encoder *encEngine;
    <a id="L355"></a>wire    *wireType;
<a id="L356"></a>}

<a id="L358"></a>var typeInfoMap = make(map[reflect.Type]*typeInfo) <span class="comment">// protected by typeLock</span>

<a id="L360"></a><span class="comment">// The reflection type must have all its indirections processed out.</span>
<a id="L361"></a><span class="comment">// typeLock must be held.</span>
<a id="L362"></a>func getTypeInfo(rt reflect.Type) (*typeInfo, os.Error) {
    <a id="L363"></a>if _, ok := rt.(*reflect.PtrType); ok {
        <a id="L364"></a>panicln(&#34;pointer type in getTypeInfo:&#34;, rt.String())
    <a id="L365"></a>}
    <a id="L366"></a>info, ok := typeInfoMap[rt];
    <a id="L367"></a>if !ok {
        <a id="L368"></a>info = new(typeInfo);
        <a id="L369"></a>name := rt.Name();
        <a id="L370"></a>gt, err := getType(name, rt);
        <a id="L371"></a>if err != nil {
            <a id="L372"></a>return nil, err
        <a id="L373"></a>}
        <a id="L374"></a>info.id = gt.id();
        <a id="L375"></a><span class="comment">// assume it&#39;s a struct type</span>
        <a id="L376"></a>info.wire = &amp;wireType{info.id.gobType().(*structType)};
        <a id="L377"></a>typeInfoMap[rt] = info;
    <a id="L378"></a>}
    <a id="L379"></a>return info, nil;
<a id="L380"></a>}

<a id="L382"></a><span class="comment">// Called only when a panic is acceptable and unexpected.</span>
<a id="L383"></a>func getTypeInfoNoError(rt reflect.Type) *typeInfo {
    <a id="L384"></a>t, err := getTypeInfo(rt);
    <a id="L385"></a>if err != nil {
        <a id="L386"></a>panicln(&#34;getTypeInfo:&#34;, err.String())
    <a id="L387"></a>}
    <a id="L388"></a>return t;
<a id="L389"></a>}
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
