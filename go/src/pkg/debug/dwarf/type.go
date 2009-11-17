<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/dwarf/type.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/debug/dwarf/type.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// DWARF type information structures.</span>
<a id="L6"></a><span class="comment">// The format is heavily biased toward C, but for simplicity</span>
<a id="L7"></a><span class="comment">// the String methods use a pseudo-Go syntax.</span>

<a id="L9"></a>package dwarf

<a id="L11"></a>import (
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;strconv&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// A Type conventionally represents a pointer to any of the</span>
<a id="L17"></a><span class="comment">// specific Type structures (CharType, StructType, etc.).</span>
<a id="L18"></a>type Type interface {
    <a id="L19"></a>Common() *CommonType;
    <a id="L20"></a>String() string;
    <a id="L21"></a>Size() int64;
<a id="L22"></a>}

<a id="L24"></a><span class="comment">// A CommonType holds fields common to multiple types.</span>
<a id="L25"></a><span class="comment">// If a field is not known or not applicable for a given type,</span>
<a id="L26"></a><span class="comment">// the zero value is used.</span>
<a id="L27"></a>type CommonType struct {
    <a id="L28"></a>ByteSize int64;  <span class="comment">// size of value of this type, in bytes</span>
    <a id="L29"></a>Name     string; <span class="comment">// name that can be used to refer to type</span>
<a id="L30"></a>}

<a id="L32"></a>func (c *CommonType) Common() *CommonType { return c }

<a id="L34"></a>func (c *CommonType) Size() int64 { return c.ByteSize }

<a id="L36"></a><span class="comment">// Basic types</span>

<a id="L38"></a><span class="comment">// A BasicType holds fields common to all basic types.</span>
<a id="L39"></a>type BasicType struct {
    <a id="L40"></a>CommonType;
    <a id="L41"></a>BitSize   int64;
    <a id="L42"></a>BitOffset int64;
<a id="L43"></a>}

<a id="L45"></a>func (b *BasicType) Basic() *BasicType { return b }

<a id="L47"></a>func (t *BasicType) String() string {
    <a id="L48"></a>if t.Name != &#34;&#34; {
        <a id="L49"></a>return t.Name
    <a id="L50"></a>}
    <a id="L51"></a>return &#34;?&#34;;
<a id="L52"></a>}

<a id="L54"></a><span class="comment">// A CharType represents a signed character type.</span>
<a id="L55"></a>type CharType struct {
    <a id="L56"></a>BasicType;
<a id="L57"></a>}

<a id="L59"></a><span class="comment">// A UcharType represents an unsigned character type.</span>
<a id="L60"></a>type UcharType struct {
    <a id="L61"></a>BasicType;
<a id="L62"></a>}

<a id="L64"></a><span class="comment">// An IntType represents a signed integer type.</span>
<a id="L65"></a>type IntType struct {
    <a id="L66"></a>BasicType;
<a id="L67"></a>}

<a id="L69"></a><span class="comment">// A UintType represents an unsigned integer type.</span>
<a id="L70"></a>type UintType struct {
    <a id="L71"></a>BasicType;
<a id="L72"></a>}

<a id="L74"></a><span class="comment">// A FloatType represents a floating point type.</span>
<a id="L75"></a>type FloatType struct {
    <a id="L76"></a>BasicType;
<a id="L77"></a>}

<a id="L79"></a><span class="comment">// A ComplexType represents a complex floating point type.</span>
<a id="L80"></a>type ComplexType struct {
    <a id="L81"></a>BasicType;
<a id="L82"></a>}

<a id="L84"></a><span class="comment">// A BoolType represents a boolean type.</span>
<a id="L85"></a>type BoolType struct {
    <a id="L86"></a>BasicType;
<a id="L87"></a>}

<a id="L89"></a><span class="comment">// An AddrType represents a machine address type.</span>
<a id="L90"></a>type AddrType struct {
    <a id="L91"></a>BasicType;
<a id="L92"></a>}

<a id="L94"></a><span class="comment">// qualifiers</span>

<a id="L96"></a><span class="comment">// A QualType represents a type that has the C/C++ &#34;const&#34;, &#34;restrict&#34;, or &#34;volatile&#34; qualifier.</span>
<a id="L97"></a>type QualType struct {
    <a id="L98"></a>CommonType;
    <a id="L99"></a>Qual string;
    <a id="L100"></a>Type Type;
<a id="L101"></a>}

<a id="L103"></a>func (t *QualType) String() string { return t.Qual + &#34; &#34; + t.Type.String() }

<a id="L105"></a>func (t *QualType) Size() int64 { return t.Type.Size() }

<a id="L107"></a><span class="comment">// An ArrayType represents a fixed size array type.</span>
<a id="L108"></a>type ArrayType struct {
    <a id="L109"></a>CommonType;
    <a id="L110"></a>Type          Type;
    <a id="L111"></a>StrideBitSize int64; <span class="comment">// if &gt; 0, number of bits to hold each element</span>
    <a id="L112"></a>Count         int64; <span class="comment">// if == -1, an incomplete array, like char x[].</span>
<a id="L113"></a>}

<a id="L115"></a>func (t *ArrayType) String() string {
    <a id="L116"></a>return &#34;[&#34; + strconv.Itoa64(t.Count) + &#34;]&#34; + t.Type.String()
<a id="L117"></a>}

<a id="L119"></a>func (t *ArrayType) Size() int64 { return t.Count * t.Type.Size() }

<a id="L121"></a><span class="comment">// A VoidType represents the C void type.</span>
<a id="L122"></a>type VoidType struct {
    <a id="L123"></a>CommonType;
<a id="L124"></a>}

<a id="L126"></a>func (t *VoidType) String() string { return &#34;void&#34; }

<a id="L128"></a><span class="comment">// A PtrType represents a pointer type.</span>
<a id="L129"></a>type PtrType struct {
    <a id="L130"></a>CommonType;
    <a id="L131"></a>Type Type;
<a id="L132"></a>}

<a id="L134"></a>func (t *PtrType) String() string { return &#34;*&#34; + t.Type.String() }

<a id="L136"></a><span class="comment">// A StructType represents a struct, union, or C++ class type.</span>
<a id="L137"></a>type StructType struct {
    <a id="L138"></a>CommonType;
    <a id="L139"></a>StructName string;
    <a id="L140"></a>Kind       string; <span class="comment">// &#34;struct&#34;, &#34;union&#34;, or &#34;class&#34;.</span>
    <a id="L141"></a>Field      []*StructField;
    <a id="L142"></a>Incomplete bool; <span class="comment">// if true, struct, union, class is declared but not defined</span>
<a id="L143"></a>}

<a id="L145"></a><span class="comment">// A StructField represents a field in a struct, union, or C++ class type.</span>
<a id="L146"></a>type StructField struct {
    <a id="L147"></a>Name       string;
    <a id="L148"></a>Type       Type;
    <a id="L149"></a>ByteOffset int64;
    <a id="L150"></a>ByteSize   int64;
    <a id="L151"></a>BitOffset  int64; <span class="comment">// within the ByteSize bytes at ByteOffset</span>
    <a id="L152"></a>BitSize    int64; <span class="comment">// zero if not a bit field</span>
<a id="L153"></a>}

<a id="L155"></a>func (t *StructType) String() string {
    <a id="L156"></a>if t.StructName != &#34;&#34; {
        <a id="L157"></a>return t.Kind + &#34; &#34; + t.StructName
    <a id="L158"></a>}
    <a id="L159"></a>return t.Defn();
<a id="L160"></a>}

<a id="L162"></a>func (t *StructType) Defn() string {
    <a id="L163"></a>s := t.Kind;
    <a id="L164"></a>if t.StructName != &#34;&#34; {
        <a id="L165"></a>s += &#34; &#34; + t.StructName
    <a id="L166"></a>}
    <a id="L167"></a>if t.Incomplete {
        <a id="L168"></a>s += &#34; /*incomplete*/&#34;;
        <a id="L169"></a>return s;
    <a id="L170"></a>}
    <a id="L171"></a>s += &#34; {&#34;;
    <a id="L172"></a>for i, f := range t.Field {
        <a id="L173"></a>if i &gt; 0 {
            <a id="L174"></a>s += &#34;; &#34;
        <a id="L175"></a>}
        <a id="L176"></a>s += f.Name + &#34; &#34; + f.Type.String();
        <a id="L177"></a>s += &#34;@&#34; + strconv.Itoa64(f.ByteOffset);
        <a id="L178"></a>if f.BitSize &gt; 0 {
            <a id="L179"></a>s += &#34; : &#34; + strconv.Itoa64(f.BitSize);
            <a id="L180"></a>s += &#34;@&#34; + strconv.Itoa64(f.BitOffset);
        <a id="L181"></a>}
    <a id="L182"></a>}
    <a id="L183"></a>s += &#34;}&#34;;
    <a id="L184"></a>return s;
<a id="L185"></a>}

<a id="L187"></a><span class="comment">// An EnumType represents an enumerated type.</span>
<a id="L188"></a><span class="comment">// The only indication of its native integer type is its ByteSize</span>
<a id="L189"></a><span class="comment">// (inside CommonType).</span>
<a id="L190"></a>type EnumType struct {
    <a id="L191"></a>CommonType;
    <a id="L192"></a>EnumName string;
    <a id="L193"></a>Val      []*EnumValue;
<a id="L194"></a>}

<a id="L196"></a><span class="comment">// An EnumValue represents a single enumeration value.</span>
<a id="L197"></a>type EnumValue struct {
    <a id="L198"></a>Name string;
    <a id="L199"></a>Val  int64;
<a id="L200"></a>}

<a id="L202"></a>func (t *EnumType) String() string {
    <a id="L203"></a>s := &#34;enum&#34;;
    <a id="L204"></a>if t.EnumName != &#34;&#34; {
        <a id="L205"></a>s += &#34; &#34; + t.EnumName
    <a id="L206"></a>}
    <a id="L207"></a>s += &#34; {&#34;;
    <a id="L208"></a>for i, v := range t.Val {
        <a id="L209"></a>if i &gt; 0 {
            <a id="L210"></a>s += &#34;; &#34;
        <a id="L211"></a>}
        <a id="L212"></a>s += v.Name + &#34;=&#34; + strconv.Itoa64(v.Val);
    <a id="L213"></a>}
    <a id="L214"></a>s += &#34;}&#34;;
    <a id="L215"></a>return s;
<a id="L216"></a>}

<a id="L218"></a><span class="comment">// A FuncType represents a function type.</span>
<a id="L219"></a>type FuncType struct {
    <a id="L220"></a>CommonType;
    <a id="L221"></a>ReturnType Type;
    <a id="L222"></a>ParamType  []Type;
<a id="L223"></a>}

<a id="L225"></a>func (t *FuncType) String() string {
    <a id="L226"></a>s := &#34;func(&#34;;
    <a id="L227"></a>for i, t := range t.ParamType {
        <a id="L228"></a>if i &gt; 0 {
            <a id="L229"></a>s += &#34;, &#34;
        <a id="L230"></a>}
        <a id="L231"></a>s += t.String();
    <a id="L232"></a>}
    <a id="L233"></a>s += &#34;)&#34;;
    <a id="L234"></a>if t.ReturnType != nil {
        <a id="L235"></a>s += &#34; &#34; + t.ReturnType.String()
    <a id="L236"></a>}
    <a id="L237"></a>return s;
<a id="L238"></a>}

<a id="L240"></a><span class="comment">// A DotDotDotType represents the variadic ... function parameter.</span>
<a id="L241"></a>type DotDotDotType struct {
    <a id="L242"></a>CommonType;
<a id="L243"></a>}

<a id="L245"></a>func (t *DotDotDotType) String() string { return &#34;...&#34; }

<a id="L247"></a><span class="comment">// A TypedefType represents a named type.</span>
<a id="L248"></a>type TypedefType struct {
    <a id="L249"></a>CommonType;
    <a id="L250"></a>Type Type;
<a id="L251"></a>}

<a id="L253"></a>func (t *TypedefType) String() string { return t.Name }

<a id="L255"></a>func (t *TypedefType) Size() int64 { return t.Type.Size() }

<a id="L257"></a>func (d *Data) Type(off Offset) (Type, os.Error) {
    <a id="L258"></a>if t, ok := d.typeCache[off]; ok {
        <a id="L259"></a>return t, nil
    <a id="L260"></a>}

    <a id="L262"></a>r := d.Reader();
    <a id="L263"></a>r.Seek(off);
    <a id="L264"></a>e, err := r.Next();
    <a id="L265"></a>if err != nil {
        <a id="L266"></a>return nil, err
    <a id="L267"></a>}
    <a id="L268"></a>if e == nil || e.Offset != off {
        <a id="L269"></a>return nil, DecodeError{&#34;info&#34;, off, &#34;no type at offset&#34;}
    <a id="L270"></a>}

    <a id="L272"></a><span class="comment">// Parse type from Entry.</span>
    <a id="L273"></a><span class="comment">// Must always set d.typeCache[off] before calling</span>
    <a id="L274"></a><span class="comment">// d.Type recursively, to handle circular types correctly.</span>
    <a id="L275"></a>var typ Type;

    <a id="L277"></a><span class="comment">// Get next child; set err if error happens.</span>
    <a id="L278"></a>next := func() *Entry {
        <a id="L279"></a>if !e.Children {
            <a id="L280"></a>return nil
        <a id="L281"></a>}
        <a id="L282"></a>kid, err1 := r.Next();
        <a id="L283"></a>if err1 != nil {
            <a id="L284"></a>err = err1;
            <a id="L285"></a>return nil;
        <a id="L286"></a>}
        <a id="L287"></a>if kid == nil {
            <a id="L288"></a>err = DecodeError{&#34;info&#34;, r.b.off, &#34;unexpected end of DWARF entries&#34;};
            <a id="L289"></a>return nil;
        <a id="L290"></a>}
        <a id="L291"></a>if kid.Tag == 0 {
            <a id="L292"></a>return nil
        <a id="L293"></a>}
        <a id="L294"></a>return kid;
    <a id="L295"></a>};

    <a id="L297"></a><span class="comment">// Get Type referred to by Entry&#39;s AttrType field.</span>
    <a id="L298"></a><span class="comment">// Set err if error happens.  Not having a type is an error.</span>
    <a id="L299"></a>typeOf := func(e *Entry) Type {
        <a id="L300"></a>toff, ok := e.Val(AttrType).(Offset);
        <a id="L301"></a>if !ok {
            <a id="L302"></a><span class="comment">// It appears that no Type means &#34;void&#34;.</span>
            <a id="L303"></a>return new(VoidType)
        <a id="L304"></a>}
        <a id="L305"></a>var t Type;
        <a id="L306"></a>if t, err = d.Type(toff); err != nil {
            <a id="L307"></a>return nil
        <a id="L308"></a>}
        <a id="L309"></a>return t;
    <a id="L310"></a>};

    <a id="L312"></a>switch e.Tag {
    <a id="L313"></a>case TagArrayType:
        <a id="L314"></a><span class="comment">// Multi-dimensional array.  (DWARF v2 §5.4)</span>
        <a id="L315"></a><span class="comment">// Attributes:</span>
        <a id="L316"></a><span class="comment">//	AttrType:subtype [required]</span>
        <a id="L317"></a><span class="comment">//	AttrStrideSize: size in bits of each element of the array</span>
        <a id="L318"></a><span class="comment">//	AttrByteSize: size of entire array</span>
        <a id="L319"></a><span class="comment">// Children:</span>
        <a id="L320"></a><span class="comment">//	TagSubrangeType or TagEnumerationType giving one dimension.</span>
        <a id="L321"></a><span class="comment">//	dimensions are in left to right order.</span>
        <a id="L322"></a>t := new(ArrayType);
        <a id="L323"></a>typ = t;
        <a id="L324"></a>d.typeCache[off] = t;
        <a id="L325"></a>if t.Type = typeOf(e); err != nil {
            <a id="L326"></a>goto Error
        <a id="L327"></a>}
        <a id="L328"></a>t.StrideBitSize, _ = e.Val(AttrStrideSize).(int64);

        <a id="L330"></a><span class="comment">// Accumulate dimensions,</span>
        <a id="L331"></a>ndim := 0;
        <a id="L332"></a>for kid := next(); kid != nil; kid = next() {
            <a id="L333"></a><span class="comment">// TODO(rsc): Can also be TagEnumerationType</span>
            <a id="L334"></a><span class="comment">// but haven&#39;t seen that in the wild yet.</span>
            <a id="L335"></a>switch kid.Tag {
            <a id="L336"></a>case TagSubrangeType:
                <a id="L337"></a>max, ok := kid.Val(AttrUpperBound).(int64);
                <a id="L338"></a>if !ok {
                    <a id="L339"></a>max = -2 <span class="comment">// Count == -1, as in x[].</span>
                <a id="L340"></a>}
                <a id="L341"></a>if ndim == 0 {
                    <a id="L342"></a>t.Count = max + 1
                <a id="L343"></a>} else {
                    <a id="L344"></a><span class="comment">// Multidimensional array.</span>
                    <a id="L345"></a><span class="comment">// Create new array type underneath this one.</span>
                    <a id="L346"></a>t.Type = &amp;ArrayType{Type: t.Type, Count: max + 1}
                <a id="L347"></a>}
                <a id="L348"></a>ndim++;
            <a id="L349"></a>case TagEnumerationType:
                <a id="L350"></a>err = DecodeError{&#34;info&#34;, kid.Offset, &#34;cannot handle enumeration type as array bound&#34;};
                <a id="L351"></a>goto Error;
            <a id="L352"></a>}
        <a id="L353"></a>}
        <a id="L354"></a>if ndim == 0 {
            <a id="L355"></a>err = DecodeError{&#34;info&#34;, e.Offset, &#34;missing dimension for array&#34;};
            <a id="L356"></a>goto Error;
        <a id="L357"></a>}

    <a id="L359"></a>case TagBaseType:
        <a id="L360"></a><span class="comment">// Basic type.  (DWARF v2 §5.1)</span>
        <a id="L361"></a><span class="comment">// Attributes:</span>
        <a id="L362"></a><span class="comment">//	AttrName: name of base type in programming language of the compilation unit [required]</span>
        <a id="L363"></a><span class="comment">//	AttrEncoding: encoding value for type (encFloat etc) [required]</span>
        <a id="L364"></a><span class="comment">//	AttrByteSize: size of type in bytes [required]</span>
        <a id="L365"></a><span class="comment">//	AttrBitOffset: for sub-byte types, size in bits</span>
        <a id="L366"></a><span class="comment">//	AttrBitSize: for sub-byte types, bit offset of high order bit in the AttrByteSize bytes</span>
        <a id="L367"></a>name, _ := e.Val(AttrName).(string);
        <a id="L368"></a>enc, ok := e.Val(AttrEncoding).(int64);
        <a id="L369"></a>if !ok {
            <a id="L370"></a>err = DecodeError{&#34;info&#34;, e.Offset, &#34;missing encoding attribute for &#34; + name};
            <a id="L371"></a>goto Error;
        <a id="L372"></a>}
        <a id="L373"></a>switch enc {
        <a id="L374"></a>default:
            <a id="L375"></a>err = DecodeError{&#34;info&#34;, e.Offset, &#34;unrecognized encoding attribute value&#34;};
            <a id="L376"></a>goto Error;

        <a id="L378"></a>case encAddress:
            <a id="L379"></a>typ = new(AddrType)
        <a id="L380"></a>case encBoolean:
            <a id="L381"></a>typ = new(BoolType)
        <a id="L382"></a>case encComplexFloat:
            <a id="L383"></a>typ = new(ComplexType)
        <a id="L384"></a>case encFloat:
            <a id="L385"></a>typ = new(FloatType)
        <a id="L386"></a>case encSigned:
            <a id="L387"></a>typ = new(IntType)
        <a id="L388"></a>case encUnsigned:
            <a id="L389"></a>typ = new(UintType)
        <a id="L390"></a>case encSignedChar:
            <a id="L391"></a>typ = new(CharType)
        <a id="L392"></a>case encUnsignedChar:
            <a id="L393"></a>typ = new(UcharType)
        <a id="L394"></a>}
        <a id="L395"></a>d.typeCache[off] = typ;
        <a id="L396"></a>t := typ.(interface {
            <a id="L397"></a>Basic() *BasicType;
        <a id="L398"></a>}).Basic();
        <a id="L399"></a>t.Name = name;
        <a id="L400"></a>t.BitSize, _ = e.Val(AttrBitSize).(int64);
        <a id="L401"></a>t.BitOffset, _ = e.Val(AttrBitOffset).(int64);

    <a id="L403"></a>case TagClassType, TagStructType, TagUnionType:
        <a id="L404"></a><span class="comment">// Structure, union, or class type.  (DWARF v2 §5.5)</span>
        <a id="L405"></a><span class="comment">// Attributes:</span>
        <a id="L406"></a><span class="comment">//	AttrName: name of struct, union, or class</span>
        <a id="L407"></a><span class="comment">//	AttrByteSize: byte size [required]</span>
        <a id="L408"></a><span class="comment">//	AttrDeclaration: if true, struct/union/class is incomplete</span>
        <a id="L409"></a><span class="comment">// Children:</span>
        <a id="L410"></a><span class="comment">//	TagMember to describe one member.</span>
        <a id="L411"></a><span class="comment">//		AttrName: name of member [required]</span>
        <a id="L412"></a><span class="comment">//		AttrType: type of member [required]</span>
        <a id="L413"></a><span class="comment">//		AttrByteSize: size in bytes</span>
        <a id="L414"></a><span class="comment">//		AttrBitOffset: bit offset within bytes for bit fields</span>
        <a id="L415"></a><span class="comment">//		AttrBitSize: bit size for bit fields</span>
        <a id="L416"></a><span class="comment">//		AttrDataMemberLoc: location within struct [required for struct, class]</span>
        <a id="L417"></a><span class="comment">// There is much more to handle C++, all ignored for now.</span>
        <a id="L418"></a>t := new(StructType);
        <a id="L419"></a>typ = t;
        <a id="L420"></a>d.typeCache[off] = t;
        <a id="L421"></a>switch e.Tag {
        <a id="L422"></a>case TagClassType:
            <a id="L423"></a>t.Kind = &#34;class&#34;
        <a id="L424"></a>case TagStructType:
            <a id="L425"></a>t.Kind = &#34;struct&#34;
        <a id="L426"></a>case TagUnionType:
            <a id="L427"></a>t.Kind = &#34;union&#34;
        <a id="L428"></a>}
        <a id="L429"></a>t.StructName, _ = e.Val(AttrName).(string);
        <a id="L430"></a>t.Incomplete = e.Val(AttrDeclaration) != nil;
        <a id="L431"></a>t.Field = make([]*StructField, 0, 8);
        <a id="L432"></a>for kid := next(); kid != nil; kid = next() {
            <a id="L433"></a>if kid.Tag == TagMember {
                <a id="L434"></a>f := new(StructField);
                <a id="L435"></a>if f.Type = typeOf(kid); err != nil {
                    <a id="L436"></a>goto Error
                <a id="L437"></a>}
                <a id="L438"></a>if loc, ok := kid.Val(AttrDataMemberLoc).([]byte); ok {
                    <a id="L439"></a>b := makeBuf(d, &#34;location&#34;, 0, loc, d.addrsize);
                    <a id="L440"></a>if b.uint8() != opPlusUconst {
                        <a id="L441"></a>err = DecodeError{&#34;info&#34;, kid.Offset, &#34;unexpected opcode&#34;};
                        <a id="L442"></a>goto Error;
                    <a id="L443"></a>}
                    <a id="L444"></a>f.ByteOffset = int64(b.uint());
                    <a id="L445"></a>if b.err != nil {
                        <a id="L446"></a>err = b.err;
                        <a id="L447"></a>goto Error;
                    <a id="L448"></a>}
                <a id="L449"></a>}
                <a id="L450"></a>f.Name, _ = kid.Val(AttrName).(string);
                <a id="L451"></a>f.ByteSize, _ = kid.Val(AttrByteSize).(int64);
                <a id="L452"></a>f.BitOffset, _ = kid.Val(AttrBitOffset).(int64);
                <a id="L453"></a>f.BitSize, _ = kid.Val(AttrBitSize).(int64);
                <a id="L454"></a>n := len(t.Field);
                <a id="L455"></a>if n &gt;= cap(t.Field) {
                    <a id="L456"></a>fld := make([]*StructField, n, n*2);
                    <a id="L457"></a>for i, f := range t.Field {
                        <a id="L458"></a>fld[i] = f
                    <a id="L459"></a>}
                    <a id="L460"></a>t.Field = fld;
                <a id="L461"></a>}
                <a id="L462"></a>t.Field = t.Field[0 : n+1];
                <a id="L463"></a>t.Field[n] = f;
            <a id="L464"></a>}
        <a id="L465"></a>}

    <a id="L467"></a>case TagConstType, TagVolatileType, TagRestrictType:
        <a id="L468"></a><span class="comment">// Type modifier (DWARF v2 §5.2)</span>
        <a id="L469"></a><span class="comment">// Attributes:</span>
        <a id="L470"></a><span class="comment">//	AttrType: subtype</span>
        <a id="L471"></a>t := new(QualType);
        <a id="L472"></a>typ = t;
        <a id="L473"></a>d.typeCache[off] = t;
        <a id="L474"></a>if t.Type = typeOf(e); err != nil {
            <a id="L475"></a>goto Error
        <a id="L476"></a>}
        <a id="L477"></a>switch e.Tag {
        <a id="L478"></a>case TagConstType:
            <a id="L479"></a>t.Qual = &#34;const&#34;
        <a id="L480"></a>case TagRestrictType:
            <a id="L481"></a>t.Qual = &#34;restrict&#34;
        <a id="L482"></a>case TagVolatileType:
            <a id="L483"></a>t.Qual = &#34;volatile&#34;
        <a id="L484"></a>}

    <a id="L486"></a>case TagEnumerationType:
        <a id="L487"></a><span class="comment">// Enumeration type (DWARF v2 §5.6)</span>
        <a id="L488"></a><span class="comment">// Attributes:</span>
        <a id="L489"></a><span class="comment">//	AttrName: enum name if any</span>
        <a id="L490"></a><span class="comment">//	AttrByteSize: bytes required to represent largest value</span>
        <a id="L491"></a><span class="comment">// Children:</span>
        <a id="L492"></a><span class="comment">//	TagEnumerator:</span>
        <a id="L493"></a><span class="comment">//		AttrName: name of constant</span>
        <a id="L494"></a><span class="comment">//		AttrConstValue: value of constant</span>
        <a id="L495"></a>t := new(EnumType);
        <a id="L496"></a>typ = t;
        <a id="L497"></a>d.typeCache[off] = t;
        <a id="L498"></a>t.EnumName, _ = e.Val(AttrName).(string);
        <a id="L499"></a>t.Val = make([]*EnumValue, 0, 8);
        <a id="L500"></a>for kid := next(); kid != nil; kid = next() {
            <a id="L501"></a>if kid.Tag == TagEnumerator {
                <a id="L502"></a>f := new(EnumValue);
                <a id="L503"></a>f.Name, _ = kid.Val(AttrName).(string);
                <a id="L504"></a>f.Val, _ = kid.Val(AttrConstValue).(int64);
                <a id="L505"></a>n := len(t.Val);
                <a id="L506"></a>if n &gt;= cap(t.Val) {
                    <a id="L507"></a>val := make([]*EnumValue, n, n*2);
                    <a id="L508"></a>for i, f := range t.Val {
                        <a id="L509"></a>val[i] = f
                    <a id="L510"></a>}
                    <a id="L511"></a>t.Val = val;
                <a id="L512"></a>}
                <a id="L513"></a>t.Val = t.Val[0 : n+1];
                <a id="L514"></a>t.Val[n] = f;
            <a id="L515"></a>}
        <a id="L516"></a>}

    <a id="L518"></a>case TagPointerType:
        <a id="L519"></a><span class="comment">// Type modifier (DWARF v2 §5.2)</span>
        <a id="L520"></a><span class="comment">// Attributes:</span>
        <a id="L521"></a><span class="comment">//	AttrType: subtype [not required!  void* has no AttrType]</span>
        <a id="L522"></a><span class="comment">//	AttrAddrClass: address class [ignored]</span>
        <a id="L523"></a>t := new(PtrType);
        <a id="L524"></a>typ = t;
        <a id="L525"></a>d.typeCache[off] = t;
        <a id="L526"></a>if e.Val(AttrType) == nil {
            <a id="L527"></a>t.Type = &amp;VoidType{};
            <a id="L528"></a>break;
        <a id="L529"></a>}
        <a id="L530"></a>t.Type = typeOf(e);

    <a id="L532"></a>case TagSubroutineType:
        <a id="L533"></a><span class="comment">// Subroutine type.  (DWARF v2 §5.7)</span>
        <a id="L534"></a><span class="comment">// Attributes:</span>
        <a id="L535"></a><span class="comment">//	AttrType: type of return value if any</span>
        <a id="L536"></a><span class="comment">//	AttrName: possible name of type [ignored]</span>
        <a id="L537"></a><span class="comment">//	AttrPrototyped: whether used ANSI C prototye [ignored]</span>
        <a id="L538"></a><span class="comment">// Children:</span>
        <a id="L539"></a><span class="comment">//	TagFormalParameter: typed parameter</span>
        <a id="L540"></a><span class="comment">//		AttrType: type of parameter</span>
        <a id="L541"></a><span class="comment">//	TagUnspecifiedParameter: final ...</span>
        <a id="L542"></a>t := new(FuncType);
        <a id="L543"></a>typ = t;
        <a id="L544"></a>d.typeCache[off] = t;
        <a id="L545"></a>if t.ReturnType = typeOf(e); err != nil {
            <a id="L546"></a>goto Error
        <a id="L547"></a>}
        <a id="L548"></a>t.ParamType = make([]Type, 0, 8);
        <a id="L549"></a>for kid := next(); kid != nil; kid = next() {
            <a id="L550"></a>var tkid Type;
            <a id="L551"></a>switch kid.Tag {
            <a id="L552"></a>default:
                <a id="L553"></a>continue
            <a id="L554"></a>case TagFormalParameter:
                <a id="L555"></a>if tkid = typeOf(kid); err != nil {
                    <a id="L556"></a>goto Error
                <a id="L557"></a>}
            <a id="L558"></a>case TagUnspecifiedParameters:
                <a id="L559"></a>tkid = &amp;DotDotDotType{}
            <a id="L560"></a>}
            <a id="L561"></a>n := len(t.ParamType);
            <a id="L562"></a>if n &gt;= cap(t.ParamType) {
                <a id="L563"></a>param := make([]Type, n, n*2);
                <a id="L564"></a>for i, t := range t.ParamType {
                    <a id="L565"></a>param[i] = t
                <a id="L566"></a>}
                <a id="L567"></a>t.ParamType = param;
            <a id="L568"></a>}
            <a id="L569"></a>t.ParamType = t.ParamType[0 : n+1];
            <a id="L570"></a>t.ParamType[n] = tkid;
        <a id="L571"></a>}

    <a id="L573"></a>case TagTypedef:
        <a id="L574"></a><span class="comment">// Typedef (DWARF v2 §5.3)</span>
        <a id="L575"></a><span class="comment">// Attributes:</span>
        <a id="L576"></a><span class="comment">//	AttrName: name [required]</span>
        <a id="L577"></a><span class="comment">//	AttrType: type definition [required]</span>
        <a id="L578"></a>t := new(TypedefType);
        <a id="L579"></a>typ = t;
        <a id="L580"></a>d.typeCache[off] = t;
        <a id="L581"></a>t.Name, _ = e.Val(AttrName).(string);
        <a id="L582"></a>t.Type = typeOf(e);
    <a id="L583"></a>}

    <a id="L585"></a>if err != nil {
        <a id="L586"></a>goto Error
    <a id="L587"></a>}

    <a id="L589"></a>b, ok := e.Val(AttrByteSize).(int64);
    <a id="L590"></a>if !ok {
        <a id="L591"></a>b = -1
    <a id="L592"></a>}
    <a id="L593"></a>typ.Common().ByteSize = b;

    <a id="L595"></a>return typ, nil;

<a id="L597"></a>Error:
    <a id="L598"></a><span class="comment">// If the parse fails, take the type out of the cache</span>
    <a id="L599"></a><span class="comment">// so that the next call with this offset doesn&#39;t hit</span>
    <a id="L600"></a><span class="comment">// the cache and return success.</span>
    <a id="L601"></a>d.typeCache[off] = nil, false;
    <a id="L602"></a>return nil, err;
<a id="L603"></a>}
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
