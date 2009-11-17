<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/type.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/type.go</h1>

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
    <a id="L9"></a>&#34;go/ast&#34;;
    <a id="L10"></a>&#34;go/token&#34;;
    <a id="L11"></a>&#34;log&#34;;
    <a id="L12"></a>&#34;reflect&#34;;
    <a id="L13"></a>&#34;sort&#34;;
    <a id="L14"></a>&#34;unsafe&#34;; <span class="comment">// For Sizeof</span>
<a id="L15"></a>)


<a id="L18"></a><span class="comment">// XXX(Spec) The type compatibility section is very confusing because</span>
<a id="L19"></a><span class="comment">// it makes it seem like there are three distinct types of</span>
<a id="L20"></a><span class="comment">// compatibility: plain compatibility, assignment compatibility, and</span>
<a id="L21"></a><span class="comment">// comparison compatibility.  As I understand it, there&#39;s really only</span>
<a id="L22"></a><span class="comment">// assignment compatibility and comparison and conversion have some</span>
<a id="L23"></a><span class="comment">// restrictions and have special meaning in some cases where the types</span>
<a id="L24"></a><span class="comment">// are not otherwise assignment compatible.  The comparison</span>
<a id="L25"></a><span class="comment">// compatibility section is almost all about the semantics of</span>
<a id="L26"></a><span class="comment">// comparison, not the type checking of it, so it would make much more</span>
<a id="L27"></a><span class="comment">// sense in the comparison operators section.  The compatibility and</span>
<a id="L28"></a><span class="comment">// assignment compatibility sections should be rolled into one.</span>

<a id="L30"></a>type Type interface {
    <a id="L31"></a><span class="comment">// compat returns whether this type is compatible with another</span>
    <a id="L32"></a><span class="comment">// type.  If conv is false, this is normal compatibility,</span>
    <a id="L33"></a><span class="comment">// where two named types are compatible only if they are the</span>
    <a id="L34"></a><span class="comment">// same named type.  If conv if true, this is conversion</span>
    <a id="L35"></a><span class="comment">// compatibility, where two named types are conversion</span>
    <a id="L36"></a><span class="comment">// compatible if their definitions are conversion compatible.</span>
    <a id="L37"></a><span class="comment">//</span>
    <a id="L38"></a><span class="comment">// TODO(austin) Deal with recursive types</span>
    <a id="L39"></a>compat(o Type, conv bool) bool;
    <a id="L40"></a><span class="comment">// lit returns this type&#39;s literal.  If this is a named type,</span>
    <a id="L41"></a><span class="comment">// this is the unnamed underlying type.  Otherwise, this is an</span>
    <a id="L42"></a><span class="comment">// identity operation.</span>
    <a id="L43"></a>lit() Type;
    <a id="L44"></a><span class="comment">// isBoolean returns true if this is a boolean type.</span>
    <a id="L45"></a>isBoolean() bool;
    <a id="L46"></a><span class="comment">// isInteger returns true if this is an integer type.</span>
    <a id="L47"></a>isInteger() bool;
    <a id="L48"></a><span class="comment">// isFloat returns true if this is a floating type.</span>
    <a id="L49"></a>isFloat() bool;
    <a id="L50"></a><span class="comment">// isIdeal returns true if this is an ideal int or float.</span>
    <a id="L51"></a>isIdeal() bool;
    <a id="L52"></a><span class="comment">// Zero returns a new zero value of this type.</span>
    <a id="L53"></a>Zero() Value;
    <a id="L54"></a><span class="comment">// String returns the string representation of this type.</span>
    <a id="L55"></a>String() string;
    <a id="L56"></a><span class="comment">// The position where this type was defined, if any.</span>
    <a id="L57"></a>Pos() token.Position;
<a id="L58"></a>}

<a id="L60"></a>type BoundedType interface {
    <a id="L61"></a>Type;
    <a id="L62"></a><span class="comment">// minVal returns the smallest value of this type.</span>
    <a id="L63"></a>minVal() *bignum.Rational;
    <a id="L64"></a><span class="comment">// maxVal returns the largest value of this type.</span>
    <a id="L65"></a>maxVal() *bignum.Rational;
<a id="L66"></a>}

<a id="L68"></a>var universePos = token.Position{&#34;&lt;universe&gt;&#34;, 0, 0, 0}

<a id="L70"></a><span class="comment">/*</span>
<a id="L71"></a><span class="comment"> * Type array maps.  These are used to memoize composite types.</span>
<a id="L72"></a><span class="comment"> */</span>

<a id="L74"></a>type typeArrayMapEntry struct {
    <a id="L75"></a>key  []Type;
    <a id="L76"></a>v    interface{};
    <a id="L77"></a>next *typeArrayMapEntry;
<a id="L78"></a>}

<a id="L80"></a>type typeArrayMap map[uintptr]*typeArrayMapEntry

<a id="L82"></a>func hashTypeArray(key []Type) uintptr {
    <a id="L83"></a>hash := uintptr(0);
    <a id="L84"></a>for _, t := range key {
        <a id="L85"></a>hash = hash * 33;
        <a id="L86"></a>if t == nil {
            <a id="L87"></a>continue
        <a id="L88"></a>}
        <a id="L89"></a>addr := reflect.NewValue(t).(*reflect.PtrValue).Get();
        <a id="L90"></a>hash ^= addr;
    <a id="L91"></a>}
    <a id="L92"></a>return hash;
<a id="L93"></a>}

<a id="L95"></a>func newTypeArrayMap() typeArrayMap { return make(map[uintptr]*typeArrayMapEntry) }

<a id="L97"></a>func (m typeArrayMap) Get(key []Type) interface{} {
    <a id="L98"></a>ent, ok := m[hashTypeArray(key)];
    <a id="L99"></a>if !ok {
        <a id="L100"></a>return nil
    <a id="L101"></a>}

<a id="L103"></a>nextEnt:
    <a id="L104"></a>for ; ent != nil; ent = ent.next {
        <a id="L105"></a>if len(key) != len(ent.key) {
            <a id="L106"></a>continue
        <a id="L107"></a>}
        <a id="L108"></a>for i := 0; i &lt; len(key); i++ {
            <a id="L109"></a>if key[i] != ent.key[i] {
                <a id="L110"></a>continue nextEnt
            <a id="L111"></a>}
        <a id="L112"></a>}
        <a id="L113"></a><span class="comment">// Found it</span>
        <a id="L114"></a>return ent.v;
    <a id="L115"></a>}

    <a id="L117"></a>return nil;
<a id="L118"></a>}

<a id="L120"></a>func (m typeArrayMap) Put(key []Type, v interface{}) interface{} {
    <a id="L121"></a>hash := hashTypeArray(key);
    <a id="L122"></a>ent, _ := m[hash];

    <a id="L124"></a>new := &amp;typeArrayMapEntry{key, v, ent};
    <a id="L125"></a>m[hash] = new;
    <a id="L126"></a>return v;
<a id="L127"></a>}

<a id="L129"></a><span class="comment">/*</span>
<a id="L130"></a><span class="comment"> * Common type</span>
<a id="L131"></a><span class="comment"> */</span>

<a id="L133"></a>type commonType struct{}

<a id="L135"></a>func (commonType) isBoolean() bool { return false }

<a id="L137"></a>func (commonType) isInteger() bool { return false }

<a id="L139"></a>func (commonType) isFloat() bool { return false }

<a id="L141"></a>func (commonType) isIdeal() bool { return false }

<a id="L143"></a>func (commonType) Pos() token.Position { return token.Position{} }

<a id="L145"></a><span class="comment">/*</span>
<a id="L146"></a><span class="comment"> * Bool</span>
<a id="L147"></a><span class="comment"> */</span>

<a id="L149"></a>type boolType struct {
    <a id="L150"></a>commonType;
<a id="L151"></a>}

<a id="L153"></a>var BoolType = universe.DefineType(&#34;bool&#34;, universePos, &amp;boolType{})

<a id="L155"></a>func (t *boolType) compat(o Type, conv bool) bool {
    <a id="L156"></a>_, ok := o.lit().(*boolType);
    <a id="L157"></a>return ok;
<a id="L158"></a>}

<a id="L160"></a>func (t *boolType) lit() Type { return t }

<a id="L162"></a>func (t *boolType) isBoolean() bool { return true }

<a id="L164"></a>func (boolType) String() string {
    <a id="L165"></a><span class="comment">// Use angle brackets as a convention for printing the</span>
    <a id="L166"></a><span class="comment">// underlying, unnamed type.  This should only show up in</span>
    <a id="L167"></a><span class="comment">// debug output.</span>
    <a id="L168"></a>return &#34;&lt;bool&gt;&#34;
<a id="L169"></a>}

<a id="L171"></a>func (t *boolType) Zero() Value {
    <a id="L172"></a>res := boolV(false);
    <a id="L173"></a>return &amp;res;
<a id="L174"></a>}

<a id="L176"></a><span class="comment">/*</span>
<a id="L177"></a><span class="comment"> * Uint</span>
<a id="L178"></a><span class="comment"> */</span>

<a id="L180"></a>type uintType struct {
    <a id="L181"></a>commonType;

    <a id="L183"></a><span class="comment">// 0 for architecture-dependent types</span>
    <a id="L184"></a>Bits uint;
    <a id="L185"></a><span class="comment">// true for uintptr, false for all others</span>
    <a id="L186"></a>Ptr  bool;
    <a id="L187"></a>name string;
<a id="L188"></a>}

<a id="L190"></a>var (
    <a id="L191"></a>Uint8Type  = universe.DefineType(&#34;uint8&#34;, universePos, &amp;uintType{commonType{}, 8, false, &#34;uint8&#34;});
    <a id="L192"></a>Uint16Type = universe.DefineType(&#34;uint16&#34;, universePos, &amp;uintType{commonType{}, 16, false, &#34;uint16&#34;});
    <a id="L193"></a>Uint32Type = universe.DefineType(&#34;uint32&#34;, universePos, &amp;uintType{commonType{}, 32, false, &#34;uint32&#34;});
    <a id="L194"></a>Uint64Type = universe.DefineType(&#34;uint64&#34;, universePos, &amp;uintType{commonType{}, 64, false, &#34;uint64&#34;});

    <a id="L196"></a>UintType    = universe.DefineType(&#34;uint&#34;, universePos, &amp;uintType{commonType{}, 0, false, &#34;uint&#34;});
    <a id="L197"></a>UintptrType = universe.DefineType(&#34;uintptr&#34;, universePos, &amp;uintType{commonType{}, 0, true, &#34;uintptr&#34;});
<a id="L198"></a>)

<a id="L200"></a>func (t *uintType) compat(o Type, conv bool) bool {
    <a id="L201"></a>t2, ok := o.lit().(*uintType);
    <a id="L202"></a>return ok &amp;&amp; t == t2;
    <a id="L203"></a>;
<a id="L204"></a>}

<a id="L206"></a>func (t *uintType) lit() Type { return t }

<a id="L208"></a>func (t *uintType) isInteger() bool { return true }

<a id="L210"></a>func (t *uintType) String() string { return &#34;&lt;&#34; + t.name + &#34;&gt;&#34; }

<a id="L212"></a>func (t *uintType) Zero() Value {
    <a id="L213"></a>switch t.Bits {
    <a id="L214"></a>case 0:
        <a id="L215"></a>if t.Ptr {
            <a id="L216"></a>res := uintptrV(0);
            <a id="L217"></a>return &amp;res;
        <a id="L218"></a>} else {
            <a id="L219"></a>res := uintV(0);
            <a id="L220"></a>return &amp;res;
        <a id="L221"></a>}
    <a id="L222"></a>case 8:
        <a id="L223"></a>res := uint8V(0);
        <a id="L224"></a>return &amp;res;
    <a id="L225"></a>case 16:
        <a id="L226"></a>res := uint16V(0);
        <a id="L227"></a>return &amp;res;
    <a id="L228"></a>case 32:
        <a id="L229"></a>res := uint32V(0);
        <a id="L230"></a>return &amp;res;
    <a id="L231"></a>case 64:
        <a id="L232"></a>res := uint64V(0);
        <a id="L233"></a>return &amp;res;
    <a id="L234"></a>}
    <a id="L235"></a>panic(&#34;unexpected uint bit count: &#34;, t.Bits);
<a id="L236"></a>}

<a id="L238"></a>func (t *uintType) minVal() *bignum.Rational { return bignum.Rat(0, 1) }

<a id="L240"></a>func (t *uintType) maxVal() *bignum.Rational {
    <a id="L241"></a>bits := t.Bits;
    <a id="L242"></a>if bits == 0 {
        <a id="L243"></a>if t.Ptr {
            <a id="L244"></a>bits = uint(8 * unsafe.Sizeof(uintptr(0)))
        <a id="L245"></a>} else {
            <a id="L246"></a>bits = uint(8 * unsafe.Sizeof(uint(0)))
        <a id="L247"></a>}
    <a id="L248"></a>}
    <a id="L249"></a>return bignum.MakeRat(bignum.Int(1).Shl(bits).Add(bignum.Int(-1)), bignum.Nat(1));
<a id="L250"></a>}

<a id="L252"></a><span class="comment">/*</span>
<a id="L253"></a><span class="comment"> * Int</span>
<a id="L254"></a><span class="comment"> */</span>

<a id="L256"></a>type intType struct {
    <a id="L257"></a>commonType;

    <a id="L259"></a><span class="comment">// XXX(Spec) Numeric types: &#34;There is also a set of</span>
    <a id="L260"></a><span class="comment">// architecture-independent basic numeric types whose size</span>
    <a id="L261"></a><span class="comment">// depends on the architecture.&#34;  Should that be</span>
    <a id="L262"></a><span class="comment">// architecture-dependent?</span>

    <a id="L264"></a><span class="comment">// 0 for architecture-dependent types</span>
    <a id="L265"></a>Bits uint;
    <a id="L266"></a>name string;
<a id="L267"></a>}

<a id="L269"></a>var (
    <a id="L270"></a>Int8Type  = universe.DefineType(&#34;int8&#34;, universePos, &amp;intType{commonType{}, 8, &#34;int8&#34;});
    <a id="L271"></a>Int16Type = universe.DefineType(&#34;int16&#34;, universePos, &amp;intType{commonType{}, 16, &#34;int16&#34;});
    <a id="L272"></a>Int32Type = universe.DefineType(&#34;int32&#34;, universePos, &amp;intType{commonType{}, 32, &#34;int32&#34;});
    <a id="L273"></a>Int64Type = universe.DefineType(&#34;int64&#34;, universePos, &amp;intType{commonType{}, 64, &#34;int64&#34;});

    <a id="L275"></a>IntType = universe.DefineType(&#34;int&#34;, universePos, &amp;intType{commonType{}, 0, &#34;int&#34;});
<a id="L276"></a>)

<a id="L278"></a>func (t *intType) compat(o Type, conv bool) bool {
    <a id="L279"></a>t2, ok := o.lit().(*intType);
    <a id="L280"></a>return ok &amp;&amp; t == t2;
<a id="L281"></a>}

<a id="L283"></a>func (t *intType) lit() Type { return t }

<a id="L285"></a>func (t *intType) isInteger() bool { return true }

<a id="L287"></a>func (t *intType) String() string { return &#34;&lt;&#34; + t.name + &#34;&gt;&#34; }

<a id="L289"></a>func (t *intType) Zero() Value {
    <a id="L290"></a>switch t.Bits {
    <a id="L291"></a>case 8:
        <a id="L292"></a>res := int8V(0);
        <a id="L293"></a>return &amp;res;
    <a id="L294"></a>case 16:
        <a id="L295"></a>res := int16V(0);
        <a id="L296"></a>return &amp;res;
    <a id="L297"></a>case 32:
        <a id="L298"></a>res := int32V(0);
        <a id="L299"></a>return &amp;res;
    <a id="L300"></a>case 64:
        <a id="L301"></a>res := int64V(0);
        <a id="L302"></a>return &amp;res;

    <a id="L304"></a>case 0:
        <a id="L305"></a>res := intV(0);
        <a id="L306"></a>return &amp;res;
    <a id="L307"></a>}
    <a id="L308"></a>panic(&#34;unexpected int bit count: &#34;, t.Bits);
<a id="L309"></a>}

<a id="L311"></a>func (t *intType) minVal() *bignum.Rational {
    <a id="L312"></a>bits := t.Bits;
    <a id="L313"></a>if bits == 0 {
        <a id="L314"></a>bits = uint(8 * unsafe.Sizeof(int(0)))
    <a id="L315"></a>}
    <a id="L316"></a>return bignum.MakeRat(bignum.Int(-1).Shl(bits-1), bignum.Nat(1));
<a id="L317"></a>}

<a id="L319"></a>func (t *intType) maxVal() *bignum.Rational {
    <a id="L320"></a>bits := t.Bits;
    <a id="L321"></a>if bits == 0 {
        <a id="L322"></a>bits = uint(8 * unsafe.Sizeof(int(0)))
    <a id="L323"></a>}
    <a id="L324"></a>return bignum.MakeRat(bignum.Int(1).Shl(bits-1).Add(bignum.Int(-1)), bignum.Nat(1));
<a id="L325"></a>}

<a id="L327"></a><span class="comment">/*</span>
<a id="L328"></a><span class="comment"> * Ideal int</span>
<a id="L329"></a><span class="comment"> */</span>

<a id="L331"></a>type idealIntType struct {
    <a id="L332"></a>commonType;
<a id="L333"></a>}

<a id="L335"></a>var IdealIntType Type = &amp;idealIntType{}

<a id="L337"></a>func (t *idealIntType) compat(o Type, conv bool) bool {
    <a id="L338"></a>_, ok := o.lit().(*idealIntType);
    <a id="L339"></a>return ok;
<a id="L340"></a>}

<a id="L342"></a>func (t *idealIntType) lit() Type { return t }

<a id="L344"></a>func (t *idealIntType) isInteger() bool { return true }

<a id="L346"></a>func (t *idealIntType) isIdeal() bool { return true }

<a id="L348"></a>func (t *idealIntType) String() string { return &#34;ideal integer&#34; }

<a id="L350"></a>func (t *idealIntType) Zero() Value { return &amp;idealIntV{bignum.Int(0)} }

<a id="L352"></a><span class="comment">/*</span>
<a id="L353"></a><span class="comment"> * Float</span>
<a id="L354"></a><span class="comment"> */</span>

<a id="L356"></a>type floatType struct {
    <a id="L357"></a>commonType;

    <a id="L359"></a><span class="comment">// 0 for architecture-dependent type</span>
    <a id="L360"></a>Bits uint;

    <a id="L362"></a>name string;
<a id="L363"></a>}

<a id="L365"></a>var (
    <a id="L366"></a>Float32Type = universe.DefineType(&#34;float32&#34;, universePos, &amp;floatType{commonType{}, 32, &#34;float32&#34;});
    <a id="L367"></a>Float64Type = universe.DefineType(&#34;float64&#34;, universePos, &amp;floatType{commonType{}, 64, &#34;float64&#34;});
    <a id="L368"></a>FloatType   = universe.DefineType(&#34;float&#34;, universePos, &amp;floatType{commonType{}, 0, &#34;float&#34;});
<a id="L369"></a>)

<a id="L371"></a>func (t *floatType) compat(o Type, conv bool) bool {
    <a id="L372"></a>t2, ok := o.lit().(*floatType);
    <a id="L373"></a>return ok &amp;&amp; t == t2;
<a id="L374"></a>}

<a id="L376"></a>func (t *floatType) lit() Type { return t }

<a id="L378"></a>func (t *floatType) isFloat() bool { return true }

<a id="L380"></a>func (t *floatType) String() string { return &#34;&lt;&#34; + t.name + &#34;&gt;&#34; }

<a id="L382"></a>func (t *floatType) Zero() Value {
    <a id="L383"></a>switch t.Bits {
    <a id="L384"></a>case 32:
        <a id="L385"></a>res := float32V(0);
        <a id="L386"></a>return &amp;res;
    <a id="L387"></a>case 64:
        <a id="L388"></a>res := float64V(0);
        <a id="L389"></a>return &amp;res;
    <a id="L390"></a>case 0:
        <a id="L391"></a>res := floatV(0);
        <a id="L392"></a>return &amp;res;
    <a id="L393"></a>}
    <a id="L394"></a>panic(&#34;unexpected float bit count: &#34;, t.Bits);
<a id="L395"></a>}

<a id="L397"></a>var maxFloat32Val = bignum.MakeRat(bignum.Int(0xffffff).Shl(127-23), bignum.Nat(1))
<a id="L398"></a>var maxFloat64Val = bignum.MakeRat(bignum.Int(0x1fffffffffffff).Shl(1023-52), bignum.Nat(1))
<a id="L399"></a>var minFloat32Val = maxFloat32Val.Neg()
<a id="L400"></a>var minFloat64Val = maxFloat64Val.Neg()

<a id="L402"></a>func (t *floatType) minVal() *bignum.Rational {
    <a id="L403"></a>bits := t.Bits;
    <a id="L404"></a>if bits == 0 {
        <a id="L405"></a>bits = uint(8 * unsafe.Sizeof(float(0)))
    <a id="L406"></a>}
    <a id="L407"></a>switch bits {
    <a id="L408"></a>case 32:
        <a id="L409"></a>return minFloat32Val
    <a id="L410"></a>case 64:
        <a id="L411"></a>return minFloat64Val
    <a id="L412"></a>}
    <a id="L413"></a>log.Crashf(&#34;unexpected floating point bit count: %d&#34;, bits);
    <a id="L414"></a>panic();
<a id="L415"></a>}

<a id="L417"></a>func (t *floatType) maxVal() *bignum.Rational {
    <a id="L418"></a>bits := t.Bits;
    <a id="L419"></a>if bits == 0 {
        <a id="L420"></a>bits = uint(8 * unsafe.Sizeof(float(0)))
    <a id="L421"></a>}
    <a id="L422"></a>switch bits {
    <a id="L423"></a>case 32:
        <a id="L424"></a>return maxFloat32Val
    <a id="L425"></a>case 64:
        <a id="L426"></a>return maxFloat64Val
    <a id="L427"></a>}
    <a id="L428"></a>log.Crashf(&#34;unexpected floating point bit count: %d&#34;, bits);
    <a id="L429"></a>panic();
<a id="L430"></a>}

<a id="L432"></a><span class="comment">/*</span>
<a id="L433"></a><span class="comment"> * Ideal float</span>
<a id="L434"></a><span class="comment"> */</span>

<a id="L436"></a>type idealFloatType struct {
    <a id="L437"></a>commonType;
<a id="L438"></a>}

<a id="L440"></a>var IdealFloatType Type = &amp;idealFloatType{}

<a id="L442"></a>func (t *idealFloatType) compat(o Type, conv bool) bool {
    <a id="L443"></a>_, ok := o.lit().(*idealFloatType);
    <a id="L444"></a>return ok;
<a id="L445"></a>}

<a id="L447"></a>func (t *idealFloatType) lit() Type { return t }

<a id="L449"></a>func (t *idealFloatType) isFloat() bool { return true }

<a id="L451"></a>func (t *idealFloatType) isIdeal() bool { return true }

<a id="L453"></a>func (t *idealFloatType) String() string { return &#34;ideal float&#34; }

<a id="L455"></a>func (t *idealFloatType) Zero() Value { return &amp;idealFloatV{bignum.Rat(1, 0)} }

<a id="L457"></a><span class="comment">/*</span>
<a id="L458"></a><span class="comment"> * String</span>
<a id="L459"></a><span class="comment"> */</span>

<a id="L461"></a>type stringType struct {
    <a id="L462"></a>commonType;
<a id="L463"></a>}

<a id="L465"></a>var StringType = universe.DefineType(&#34;string&#34;, universePos, &amp;stringType{})

<a id="L467"></a>func (t *stringType) compat(o Type, conv bool) bool {
    <a id="L468"></a>_, ok := o.lit().(*stringType);
    <a id="L469"></a>return ok;
<a id="L470"></a>}

<a id="L472"></a>func (t *stringType) lit() Type { return t }

<a id="L474"></a>func (t *stringType) String() string { return &#34;&lt;string&gt;&#34; }

<a id="L476"></a>func (t *stringType) Zero() Value {
    <a id="L477"></a>res := stringV(&#34;&#34;);
    <a id="L478"></a>return &amp;res;
<a id="L479"></a>}

<a id="L481"></a><span class="comment">/*</span>
<a id="L482"></a><span class="comment"> * Array</span>
<a id="L483"></a><span class="comment"> */</span>

<a id="L485"></a>type ArrayType struct {
    <a id="L486"></a>commonType;
    <a id="L487"></a>Len  int64;
    <a id="L488"></a>Elem Type;
<a id="L489"></a>}

<a id="L491"></a>var arrayTypes = make(map[int64]map[Type]*ArrayType)

<a id="L493"></a><span class="comment">// Two array types are identical if they have identical element types</span>
<a id="L494"></a><span class="comment">// and the same array length.</span>

<a id="L496"></a>func NewArrayType(len int64, elem Type) *ArrayType {
    <a id="L497"></a>ts, ok := arrayTypes[len];
    <a id="L498"></a>if !ok {
        <a id="L499"></a>ts = make(map[Type]*ArrayType);
        <a id="L500"></a>arrayTypes[len] = ts;
    <a id="L501"></a>}
    <a id="L502"></a>t, ok := ts[elem];
    <a id="L503"></a>if !ok {
        <a id="L504"></a>t = &amp;ArrayType{commonType{}, len, elem};
        <a id="L505"></a>ts[elem] = t;
    <a id="L506"></a>}
    <a id="L507"></a>return t;
<a id="L508"></a>}

<a id="L510"></a>func (t *ArrayType) compat(o Type, conv bool) bool {
    <a id="L511"></a>t2, ok := o.lit().(*ArrayType);
    <a id="L512"></a>if !ok {
        <a id="L513"></a>return false
    <a id="L514"></a>}
    <a id="L515"></a>return t.Len == t2.Len &amp;&amp; t.Elem.compat(t2.Elem, conv);
<a id="L516"></a>}

<a id="L518"></a>func (t *ArrayType) lit() Type { return t }

<a id="L520"></a>func (t *ArrayType) String() string { return &#34;[]&#34; + t.Elem.String() }

<a id="L522"></a>func (t *ArrayType) Zero() Value {
    <a id="L523"></a>res := arrayV(make([]Value, t.Len));
    <a id="L524"></a><span class="comment">// TODO(austin) It&#39;s unfortunate that each element is</span>
    <a id="L525"></a><span class="comment">// separately heap allocated.  We could add ZeroArray to</span>
    <a id="L526"></a><span class="comment">// everything, though that doesn&#39;t help with multidimensional</span>
    <a id="L527"></a><span class="comment">// arrays.  Or we could do something unsafe.  We&#39;ll have this</span>
    <a id="L528"></a><span class="comment">// same problem with structs.</span>
    <a id="L529"></a>for i := int64(0); i &lt; t.Len; i++ {
        <a id="L530"></a>res[i] = t.Elem.Zero()
    <a id="L531"></a>}
    <a id="L532"></a>return &amp;res;
<a id="L533"></a>}

<a id="L535"></a><span class="comment">/*</span>
<a id="L536"></a><span class="comment"> * Struct</span>
<a id="L537"></a><span class="comment"> */</span>

<a id="L539"></a>type StructField struct {
    <a id="L540"></a>Name      string;
    <a id="L541"></a>Type      Type;
    <a id="L542"></a>Anonymous bool;
<a id="L543"></a>}

<a id="L545"></a>type StructType struct {
    <a id="L546"></a>commonType;
    <a id="L547"></a>Elems []StructField;
<a id="L548"></a>}

<a id="L550"></a>var structTypes = newTypeArrayMap()

<a id="L552"></a><span class="comment">// Two struct types are identical if they have the same sequence of</span>
<a id="L553"></a><span class="comment">// fields, and if corresponding fields have the same names and</span>
<a id="L554"></a><span class="comment">// identical types. Two anonymous fields are considered to have the</span>
<a id="L555"></a><span class="comment">// same name.</span>

<a id="L557"></a>func NewStructType(fields []StructField) *StructType {
    <a id="L558"></a><span class="comment">// Start by looking up just the types</span>
    <a id="L559"></a>fts := make([]Type, len(fields));
    <a id="L560"></a>for i, f := range fields {
        <a id="L561"></a>fts[i] = f.Type
    <a id="L562"></a>}
    <a id="L563"></a>tMapI := structTypes.Get(fts);
    <a id="L564"></a>if tMapI == nil {
        <a id="L565"></a>tMapI = structTypes.Put(fts, make(map[string]*StructType))
    <a id="L566"></a>}
    <a id="L567"></a>tMap := tMapI.(map[string]*StructType);

    <a id="L569"></a><span class="comment">// Construct key for field names</span>
    <a id="L570"></a>key := &#34;&#34;;
    <a id="L571"></a>for _, f := range fields {
        <a id="L572"></a><span class="comment">// XXX(Spec) It&#39;s not clear if struct { T } and struct</span>
        <a id="L573"></a><span class="comment">// { T T } are either identical or compatible.  The</span>
        <a id="L574"></a><span class="comment">// &#34;Struct Types&#34; section says that the name of that</span>
        <a id="L575"></a><span class="comment">// field is &#34;T&#34;, which suggests that they are</span>
        <a id="L576"></a><span class="comment">// identical, but it really means that it&#39;s the name</span>
        <a id="L577"></a><span class="comment">// for the purpose of selector expressions and nothing</span>
        <a id="L578"></a><span class="comment">// else.  We decided that they should be neither</span>
        <a id="L579"></a><span class="comment">// identical or compatible.</span>
        <a id="L580"></a>if f.Anonymous {
            <a id="L581"></a>key += &#34;!&#34;
        <a id="L582"></a>}
        <a id="L583"></a>key += f.Name + &#34; &#34;;
    <a id="L584"></a>}

    <a id="L586"></a><span class="comment">// XXX(Spec) Do the tags also have to be identical for the</span>
    <a id="L587"></a><span class="comment">// types to be identical?  I certainly hope so, because</span>
    <a id="L588"></a><span class="comment">// otherwise, this is the only case where two distinct type</span>
    <a id="L589"></a><span class="comment">// objects can represent identical types.</span>

    <a id="L591"></a>t, ok := tMap[key];
    <a id="L592"></a>if !ok {
        <a id="L593"></a><span class="comment">// Create new struct type</span>
        <a id="L594"></a>t = &amp;StructType{commonType{}, fields};
        <a id="L595"></a>tMap[key] = t;
    <a id="L596"></a>}
    <a id="L597"></a>return t;
<a id="L598"></a>}

<a id="L600"></a>func (t *StructType) compat(o Type, conv bool) bool {
    <a id="L601"></a>t2, ok := o.lit().(*StructType);
    <a id="L602"></a>if !ok {
        <a id="L603"></a>return false
    <a id="L604"></a>}
    <a id="L605"></a>if len(t.Elems) != len(t2.Elems) {
        <a id="L606"></a>return false
    <a id="L607"></a>}
    <a id="L608"></a>for i, e := range t.Elems {
        <a id="L609"></a>e2 := t2.Elems[i];
        <a id="L610"></a><span class="comment">// XXX(Spec) An anonymous and a non-anonymous field</span>
        <a id="L611"></a><span class="comment">// are neither identical nor compatible.</span>
        <a id="L612"></a>if e.Anonymous != e2.Anonymous ||
            <a id="L613"></a>(!e.Anonymous &amp;&amp; e.Name != e2.Name) ||
            <a id="L614"></a>!e.Type.compat(e2.Type, conv) {
            <a id="L615"></a>return false
        <a id="L616"></a>}
    <a id="L617"></a>}
    <a id="L618"></a>return true;
<a id="L619"></a>}

<a id="L621"></a>func (t *StructType) lit() Type { return t }

<a id="L623"></a>func (t *StructType) String() string {
    <a id="L624"></a>s := &#34;struct {&#34;;
    <a id="L625"></a>for i, f := range t.Elems {
        <a id="L626"></a>if i &gt; 0 {
            <a id="L627"></a>s += &#34;; &#34;
        <a id="L628"></a>}
        <a id="L629"></a>if !f.Anonymous {
            <a id="L630"></a>s += f.Name + &#34; &#34;
        <a id="L631"></a>}
        <a id="L632"></a>s += f.Type.String();
    <a id="L633"></a>}
    <a id="L634"></a>return s + &#34;}&#34;;
<a id="L635"></a>}

<a id="L637"></a>func (t *StructType) Zero() Value {
    <a id="L638"></a>res := structV(make([]Value, len(t.Elems)));
    <a id="L639"></a>for i, f := range t.Elems {
        <a id="L640"></a>res[i] = f.Type.Zero()
    <a id="L641"></a>}
    <a id="L642"></a>return &amp;res;
<a id="L643"></a>}

<a id="L645"></a><span class="comment">/*</span>
<a id="L646"></a><span class="comment"> * Pointer</span>
<a id="L647"></a><span class="comment"> */</span>

<a id="L649"></a>type PtrType struct {
    <a id="L650"></a>commonType;
    <a id="L651"></a>Elem Type;
<a id="L652"></a>}

<a id="L654"></a>var ptrTypes = make(map[Type]*PtrType)

<a id="L656"></a><span class="comment">// Two pointer types are identical if they have identical base types.</span>

<a id="L658"></a>func NewPtrType(elem Type) *PtrType {
    <a id="L659"></a>t, ok := ptrTypes[elem];
    <a id="L660"></a>if !ok {
        <a id="L661"></a>t = &amp;PtrType{commonType{}, elem};
        <a id="L662"></a>ptrTypes[elem] = t;
    <a id="L663"></a>}
    <a id="L664"></a>return t;
<a id="L665"></a>}

<a id="L667"></a>func (t *PtrType) compat(o Type, conv bool) bool {
    <a id="L668"></a>t2, ok := o.lit().(*PtrType);
    <a id="L669"></a>if !ok {
        <a id="L670"></a>return false
    <a id="L671"></a>}
    <a id="L672"></a>return t.Elem.compat(t2.Elem, conv);
<a id="L673"></a>}

<a id="L675"></a>func (t *PtrType) lit() Type { return t }

<a id="L677"></a>func (t *PtrType) String() string { return &#34;*&#34; + t.Elem.String() }

<a id="L679"></a>func (t *PtrType) Zero() Value { return &amp;ptrV{nil} }

<a id="L681"></a><span class="comment">/*</span>
<a id="L682"></a><span class="comment"> * Function</span>
<a id="L683"></a><span class="comment"> */</span>

<a id="L685"></a>type FuncType struct {
    <a id="L686"></a>commonType;
    <a id="L687"></a><span class="comment">// TODO(austin) Separate receiver Type for methods?</span>
    <a id="L688"></a>In       []Type;
    <a id="L689"></a>Variadic bool;
    <a id="L690"></a>Out      []Type;
    <a id="L691"></a>builtin  string;
<a id="L692"></a>}

<a id="L694"></a>var funcTypes = newTypeArrayMap()
<a id="L695"></a>var variadicFuncTypes = newTypeArrayMap()

<a id="L697"></a><span class="comment">// Create singleton function types for magic built-in functions</span>
<a id="L698"></a>var (
    <a id="L699"></a>capType     = &amp;FuncType{builtin: &#34;cap&#34;};
    <a id="L700"></a>closeType   = &amp;FuncType{builtin: &#34;close&#34;};
    <a id="L701"></a>closedType  = &amp;FuncType{builtin: &#34;closed&#34;};
    <a id="L702"></a>lenType     = &amp;FuncType{builtin: &#34;len&#34;};
    <a id="L703"></a>makeType    = &amp;FuncType{builtin: &#34;make&#34;};
    <a id="L704"></a>newType     = &amp;FuncType{builtin: &#34;new&#34;};
    <a id="L705"></a>panicType   = &amp;FuncType{builtin: &#34;panic&#34;};
    <a id="L706"></a>paniclnType = &amp;FuncType{builtin: &#34;panicln&#34;};
    <a id="L707"></a>printType   = &amp;FuncType{builtin: &#34;print&#34;};
    <a id="L708"></a>printlnType = &amp;FuncType{builtin: &#34;println&#34;};
<a id="L709"></a>)

<a id="L711"></a><span class="comment">// Two function types are identical if they have the same number of</span>
<a id="L712"></a><span class="comment">// parameters and result values and if corresponding parameter and</span>
<a id="L713"></a><span class="comment">// result types are identical. All &#34;...&#34; parameters have identical</span>
<a id="L714"></a><span class="comment">// type. Parameter and result names are not required to match.</span>

<a id="L716"></a>func NewFuncType(in []Type, variadic bool, out []Type) *FuncType {
    <a id="L717"></a>inMap := funcTypes;
    <a id="L718"></a>if variadic {
        <a id="L719"></a>inMap = variadicFuncTypes
    <a id="L720"></a>}

    <a id="L722"></a>outMapI := inMap.Get(in);
    <a id="L723"></a>if outMapI == nil {
        <a id="L724"></a>outMapI = inMap.Put(in, newTypeArrayMap())
    <a id="L725"></a>}
    <a id="L726"></a>outMap := outMapI.(typeArrayMap);

    <a id="L728"></a>tI := outMap.Get(out);
    <a id="L729"></a>if tI != nil {
        <a id="L730"></a>return tI.(*FuncType)
    <a id="L731"></a>}

    <a id="L733"></a>t := &amp;FuncType{commonType{}, in, variadic, out, &#34;&#34;};
    <a id="L734"></a>outMap.Put(out, t);
    <a id="L735"></a>return t;
<a id="L736"></a>}

<a id="L738"></a>func (t *FuncType) compat(o Type, conv bool) bool {
    <a id="L739"></a>t2, ok := o.lit().(*FuncType);
    <a id="L740"></a>if !ok {
        <a id="L741"></a>return false
    <a id="L742"></a>}
    <a id="L743"></a>if len(t.In) != len(t2.In) || t.Variadic != t2.Variadic || len(t.Out) != len(t2.Out) {
        <a id="L744"></a>return false
    <a id="L745"></a>}
    <a id="L746"></a>for i := range t.In {
        <a id="L747"></a>if !t.In[i].compat(t2.In[i], conv) {
            <a id="L748"></a>return false
        <a id="L749"></a>}
    <a id="L750"></a>}
    <a id="L751"></a>for i := range t.Out {
        <a id="L752"></a>if !t.Out[i].compat(t2.Out[i], conv) {
            <a id="L753"></a>return false
        <a id="L754"></a>}
    <a id="L755"></a>}
    <a id="L756"></a>return true;
<a id="L757"></a>}

<a id="L759"></a>func (t *FuncType) lit() Type { return t }

<a id="L761"></a>func typeListString(ts []Type, ns []*ast.Ident) string {
    <a id="L762"></a>s := &#34;&#34;;
    <a id="L763"></a>for i, t := range ts {
        <a id="L764"></a>if i &gt; 0 {
            <a id="L765"></a>s += &#34;, &#34;
        <a id="L766"></a>}
        <a id="L767"></a>if ns != nil &amp;&amp; ns[i] != nil {
            <a id="L768"></a>s += ns[i].Value + &#34; &#34;
        <a id="L769"></a>}
        <a id="L770"></a>if t == nil {
            <a id="L771"></a><span class="comment">// Some places use nil types to represent errors</span>
            <a id="L772"></a>s += &#34;&lt;none&gt;&#34;
        <a id="L773"></a>} else {
            <a id="L774"></a>s += t.String()
        <a id="L775"></a>}
    <a id="L776"></a>}
    <a id="L777"></a>return s;
<a id="L778"></a>}

<a id="L780"></a>func (t *FuncType) String() string {
    <a id="L781"></a>if t.builtin != &#34;&#34; {
        <a id="L782"></a>return &#34;built-in function &#34; + t.builtin
    <a id="L783"></a>}
    <a id="L784"></a>args := typeListString(t.In, nil);
    <a id="L785"></a>if t.Variadic {
        <a id="L786"></a>if len(args) &gt; 0 {
            <a id="L787"></a>args += &#34;, &#34;
        <a id="L788"></a>}
        <a id="L789"></a>args += &#34;...&#34;;
    <a id="L790"></a>}
    <a id="L791"></a>s := &#34;func(&#34; + args + &#34;)&#34;;
    <a id="L792"></a>if len(t.Out) &gt; 0 {
        <a id="L793"></a>s += &#34; (&#34; + typeListString(t.Out, nil) + &#34;)&#34;
    <a id="L794"></a>}
    <a id="L795"></a>return s;
<a id="L796"></a>}

<a id="L798"></a>func (t *FuncType) Zero() Value { return &amp;funcV{nil} }

<a id="L800"></a>type FuncDecl struct {
    <a id="L801"></a>Type *FuncType;
    <a id="L802"></a>Name *ast.Ident; <span class="comment">// nil for function literals</span>
    <a id="L803"></a><span class="comment">// InNames will be one longer than Type.In if this function is</span>
    <a id="L804"></a><span class="comment">// variadic.</span>
    <a id="L805"></a>InNames  []*ast.Ident;
    <a id="L806"></a>OutNames []*ast.Ident;
<a id="L807"></a>}

<a id="L809"></a>func (t *FuncDecl) String() string {
    <a id="L810"></a>s := &#34;func&#34;;
    <a id="L811"></a>if t.Name != nil {
        <a id="L812"></a>s += &#34; &#34; + t.Name.Value
    <a id="L813"></a>}
    <a id="L814"></a>s += funcTypeString(t.Type, t.InNames, t.OutNames);
    <a id="L815"></a>return s;
<a id="L816"></a>}

<a id="L818"></a>func funcTypeString(ft *FuncType, ins []*ast.Ident, outs []*ast.Ident) string {
    <a id="L819"></a>s := &#34;(&#34;;
    <a id="L820"></a>s += typeListString(ft.In, ins);
    <a id="L821"></a>if ft.Variadic {
        <a id="L822"></a>if len(ft.In) &gt; 0 {
            <a id="L823"></a>s += &#34;, &#34;
        <a id="L824"></a>}
        <a id="L825"></a>s += &#34;...&#34;;
    <a id="L826"></a>}
    <a id="L827"></a>s += &#34;)&#34;;
    <a id="L828"></a>if len(ft.Out) &gt; 0 {
        <a id="L829"></a>s += &#34; (&#34; + typeListString(ft.Out, outs) + &#34;)&#34;
    <a id="L830"></a>}
    <a id="L831"></a>return s;
<a id="L832"></a>}

<a id="L834"></a><span class="comment">/*</span>
<a id="L835"></a><span class="comment"> * Interface</span>
<a id="L836"></a><span class="comment"> */</span>

<a id="L838"></a><span class="comment">// TODO(austin) Interface values, types, and type compilation are</span>
<a id="L839"></a><span class="comment">// implemented, but none of the type checking or semantics of</span>
<a id="L840"></a><span class="comment">// interfaces are.</span>

<a id="L842"></a>type InterfaceType struct {
    <a id="L843"></a>commonType;
    <a id="L844"></a><span class="comment">// TODO(austin) This should be a map from names to</span>
    <a id="L845"></a><span class="comment">// *FuncType&#39;s.  We only need the sorted list for generating</span>
    <a id="L846"></a><span class="comment">// the type map key.  It&#39;s detrimental for everything else.</span>
    <a id="L847"></a>methods []IMethod;
<a id="L848"></a>}

<a id="L850"></a>type IMethod struct {
    <a id="L851"></a>Name string;
    <a id="L852"></a>Type *FuncType;
<a id="L853"></a>}

<a id="L855"></a>var interfaceTypes = newTypeArrayMap()

<a id="L857"></a>func NewInterfaceType(methods []IMethod, embeds []*InterfaceType) *InterfaceType {
    <a id="L858"></a><span class="comment">// Count methods of embedded interfaces</span>
    <a id="L859"></a>nMethods := len(methods);
    <a id="L860"></a>for _, e := range embeds {
        <a id="L861"></a>nMethods += len(e.methods)
    <a id="L862"></a>}

    <a id="L864"></a><span class="comment">// Combine methods</span>
    <a id="L865"></a>allMethods := make([]IMethod, nMethods);
    <a id="L866"></a>for i, m := range methods {
        <a id="L867"></a>allMethods[i] = m
    <a id="L868"></a>}
    <a id="L869"></a>n := len(methods);
    <a id="L870"></a>for _, e := range embeds {
        <a id="L871"></a>for _, m := range e.methods {
            <a id="L872"></a>allMethods[n] = m;
            <a id="L873"></a>n++;
        <a id="L874"></a>}
    <a id="L875"></a>}

    <a id="L877"></a><span class="comment">// Sort methods</span>
    <a id="L878"></a>sort.Sort(iMethodSorter(allMethods));

    <a id="L880"></a>mts := make([]Type, len(allMethods));
    <a id="L881"></a>for i, m := range methods {
        <a id="L882"></a>mts[i] = m.Type
    <a id="L883"></a>}
    <a id="L884"></a>tMapI := interfaceTypes.Get(mts);
    <a id="L885"></a>if tMapI == nil {
        <a id="L886"></a>tMapI = interfaceTypes.Put(mts, make(map[string]*InterfaceType))
    <a id="L887"></a>}
    <a id="L888"></a>tMap := tMapI.(map[string]*InterfaceType);

    <a id="L890"></a>key := &#34;&#34;;
    <a id="L891"></a>for _, m := range allMethods {
        <a id="L892"></a>key += m.Name + &#34; &#34;
    <a id="L893"></a>}

    <a id="L895"></a>t, ok := tMap[key];
    <a id="L896"></a>if !ok {
        <a id="L897"></a>t = &amp;InterfaceType{commonType{}, allMethods};
        <a id="L898"></a>tMap[key] = t;
    <a id="L899"></a>}
    <a id="L900"></a>return t;
<a id="L901"></a>}

<a id="L903"></a>type iMethodSorter []IMethod

<a id="L905"></a>func (s iMethodSorter) Less(a, b int) bool { return s[a].Name &lt; s[b].Name }

<a id="L907"></a>func (s iMethodSorter) Swap(a, b int) { s[a], s[b] = s[b], s[a] }

<a id="L909"></a>func (s iMethodSorter) Len() int { return len(s) }

<a id="L911"></a>func (t *InterfaceType) compat(o Type, conv bool) bool {
    <a id="L912"></a>t2, ok := o.lit().(*InterfaceType);
    <a id="L913"></a>if !ok {
        <a id="L914"></a>return false
    <a id="L915"></a>}
    <a id="L916"></a>if len(t.methods) != len(t2.methods) {
        <a id="L917"></a>return false
    <a id="L918"></a>}
    <a id="L919"></a>for i, e := range t.methods {
        <a id="L920"></a>e2 := t2.methods[i];
        <a id="L921"></a>if e.Name != e2.Name || !e.Type.compat(e2.Type, conv) {
            <a id="L922"></a>return false
        <a id="L923"></a>}
    <a id="L924"></a>}
    <a id="L925"></a>return true;
<a id="L926"></a>}

<a id="L928"></a>func (t *InterfaceType) lit() Type { return t }

<a id="L930"></a>func (t *InterfaceType) String() string {
    <a id="L931"></a><span class="comment">// TODO(austin) Instead of showing embedded interfaces, this</span>
    <a id="L932"></a><span class="comment">// shows their methods.</span>
    <a id="L933"></a>s := &#34;interface {&#34;;
    <a id="L934"></a>for i, m := range t.methods {
        <a id="L935"></a>if i &gt; 0 {
            <a id="L936"></a>s += &#34;; &#34;
        <a id="L937"></a>}
        <a id="L938"></a>s += m.Name + funcTypeString(m.Type, nil, nil);
    <a id="L939"></a>}
    <a id="L940"></a>return s + &#34;}&#34;;
<a id="L941"></a>}

<a id="L943"></a><span class="comment">// implementedBy tests if o implements t, returning nil, true if it does.</span>
<a id="L944"></a><span class="comment">// Otherwise, it returns a method of t that o is missing and false.</span>
<a id="L945"></a>func (t *InterfaceType) implementedBy(o Type) (*IMethod, bool) {
    <a id="L946"></a>if len(t.methods) == 0 {
        <a id="L947"></a>return nil, true
    <a id="L948"></a>}

    <a id="L950"></a><span class="comment">// The methods of a named interface types are those of the</span>
    <a id="L951"></a><span class="comment">// underlying type.</span>
    <a id="L952"></a>if it, ok := o.lit().(*InterfaceType); ok {
        <a id="L953"></a>o = it
    <a id="L954"></a>}

    <a id="L956"></a><span class="comment">// XXX(Spec) Interface types: &#34;A type implements any interface</span>
    <a id="L957"></a><span class="comment">// comprising any subset of its methods&#34; It&#39;s unclear if</span>
    <a id="L958"></a><span class="comment">// methods must have identical or compatible types.  6g</span>
    <a id="L959"></a><span class="comment">// requires identical types.</span>

    <a id="L961"></a>switch o := o.(type) {
    <a id="L962"></a>case *NamedType:
        <a id="L963"></a>for _, tm := range t.methods {
            <a id="L964"></a>sm, ok := o.methods[tm.Name];
            <a id="L965"></a>if !ok || sm.decl.Type != tm.Type {
                <a id="L966"></a>return &amp;tm, false
            <a id="L967"></a>}
        <a id="L968"></a>}
        <a id="L969"></a>return nil, true;

    <a id="L971"></a>case *InterfaceType:
        <a id="L972"></a>var ti, oi int;
        <a id="L973"></a>for ti &lt; len(t.methods) &amp;&amp; oi &lt; len(o.methods) {
            <a id="L974"></a>tm, om := &amp;t.methods[ti], &amp;o.methods[oi];
            <a id="L975"></a>switch {
            <a id="L976"></a>case tm.Name == om.Name:
                <a id="L977"></a>if tm.Type != om.Type {
                    <a id="L978"></a>return tm, false
                <a id="L979"></a>}
                <a id="L980"></a>ti++;
                <a id="L981"></a>oi++;
            <a id="L982"></a>case tm.Name &gt; om.Name:
                <a id="L983"></a>oi++
            <a id="L984"></a>default:
                <a id="L985"></a>return tm, false
            <a id="L986"></a>}
        <a id="L987"></a>}
        <a id="L988"></a>if ti &lt; len(t.methods) {
            <a id="L989"></a>return &amp;t.methods[ti], false
        <a id="L990"></a>}
        <a id="L991"></a>return nil, true;
    <a id="L992"></a>}

    <a id="L994"></a>return &amp;t.methods[0], false;
<a id="L995"></a>}

<a id="L997"></a>func (t *InterfaceType) Zero() Value { return &amp;interfaceV{} }

<a id="L999"></a><span class="comment">/*</span>
<a id="L1000"></a><span class="comment"> * Slice</span>
<a id="L1001"></a><span class="comment"> */</span>

<a id="L1003"></a>type SliceType struct {
    <a id="L1004"></a>commonType;
    <a id="L1005"></a>Elem Type;
<a id="L1006"></a>}

<a id="L1008"></a>var sliceTypes = make(map[Type]*SliceType)

<a id="L1010"></a><span class="comment">// Two slice types are identical if they have identical element types.</span>

<a id="L1012"></a>func NewSliceType(elem Type) *SliceType {
    <a id="L1013"></a>t, ok := sliceTypes[elem];
    <a id="L1014"></a>if !ok {
        <a id="L1015"></a>t = &amp;SliceType{commonType{}, elem};
        <a id="L1016"></a>sliceTypes[elem] = t;
    <a id="L1017"></a>}
    <a id="L1018"></a>return t;
<a id="L1019"></a>}

<a id="L1021"></a>func (t *SliceType) compat(o Type, conv bool) bool {
    <a id="L1022"></a>t2, ok := o.lit().(*SliceType);
    <a id="L1023"></a>if !ok {
        <a id="L1024"></a>return false
    <a id="L1025"></a>}
    <a id="L1026"></a>return t.Elem.compat(t2.Elem, conv);
<a id="L1027"></a>}

<a id="L1029"></a>func (t *SliceType) lit() Type { return t }

<a id="L1031"></a>func (t *SliceType) String() string { return &#34;[]&#34; + t.Elem.String() }

<a id="L1033"></a>func (t *SliceType) Zero() Value {
    <a id="L1034"></a><span class="comment">// The value of an uninitialized slice is nil. The length and</span>
    <a id="L1035"></a><span class="comment">// capacity of a nil slice are 0.</span>
    <a id="L1036"></a>return &amp;sliceV{Slice{nil, 0, 0}}
<a id="L1037"></a>}

<a id="L1039"></a><span class="comment">/*</span>
<a id="L1040"></a><span class="comment"> * Map type</span>
<a id="L1041"></a><span class="comment"> */</span>

<a id="L1043"></a>type MapType struct {
    <a id="L1044"></a>commonType;
    <a id="L1045"></a>Key  Type;
    <a id="L1046"></a>Elem Type;
<a id="L1047"></a>}

<a id="L1049"></a>var mapTypes = make(map[Type]map[Type]*MapType)

<a id="L1051"></a>func NewMapType(key Type, elem Type) *MapType {
    <a id="L1052"></a>ts, ok := mapTypes[key];
    <a id="L1053"></a>if !ok {
        <a id="L1054"></a>ts = make(map[Type]*MapType);
        <a id="L1055"></a>mapTypes[key] = ts;
    <a id="L1056"></a>}
    <a id="L1057"></a>t, ok := ts[elem];
    <a id="L1058"></a>if !ok {
        <a id="L1059"></a>t = &amp;MapType{commonType{}, key, elem};
        <a id="L1060"></a>ts[elem] = t;
    <a id="L1061"></a>}
    <a id="L1062"></a>return t;
<a id="L1063"></a>}

<a id="L1065"></a>func (t *MapType) compat(o Type, conv bool) bool {
    <a id="L1066"></a>t2, ok := o.lit().(*MapType);
    <a id="L1067"></a>if !ok {
        <a id="L1068"></a>return false
    <a id="L1069"></a>}
    <a id="L1070"></a>return t.Elem.compat(t2.Elem, conv) &amp;&amp; t.Key.compat(t2.Key, conv);
<a id="L1071"></a>}

<a id="L1073"></a>func (t *MapType) lit() Type { return t }

<a id="L1075"></a>func (t *MapType) String() string { return &#34;map[&#34; + t.Key.String() + &#34;] &#34; + t.Elem.String() }

<a id="L1077"></a>func (t *MapType) Zero() Value {
    <a id="L1078"></a><span class="comment">// The value of an uninitialized map is nil.</span>
    <a id="L1079"></a>return &amp;mapV{nil}
<a id="L1080"></a>}

<a id="L1082"></a><span class="comment">/*</span>
<a id="L1083"></a><span class="comment">type ChanType struct {</span>
<a id="L1084"></a><span class="comment">	// TODO(austin)</span>
<a id="L1085"></a><span class="comment">}</span>
<a id="L1086"></a><span class="comment">*/</span>

<a id="L1088"></a><span class="comment">/*</span>
<a id="L1089"></a><span class="comment"> * Named types</span>
<a id="L1090"></a><span class="comment"> */</span>

<a id="L1092"></a>type Method struct {
    <a id="L1093"></a>decl *FuncDecl;
    <a id="L1094"></a>fn   Func;
<a id="L1095"></a>}

<a id="L1097"></a>type NamedType struct {
    <a id="L1098"></a>token.Position;
    <a id="L1099"></a>Name string;
    <a id="L1100"></a><span class="comment">// Underlying type.  If incomplete is true, this will be nil.</span>
    <a id="L1101"></a><span class="comment">// If incomplete is false and this is still nil, then this is</span>
    <a id="L1102"></a><span class="comment">// a placeholder type representing an error.</span>
    <a id="L1103"></a>Def Type;
    <a id="L1104"></a><span class="comment">// True while this type is being defined.</span>
    <a id="L1105"></a>incomplete bool;
    <a id="L1106"></a>methods    map[string]Method;
<a id="L1107"></a>}

<a id="L1109"></a><span class="comment">// TODO(austin) This is temporarily needed by the debugger&#39;s remote</span>
<a id="L1110"></a><span class="comment">// type parser.  This should only be possible with block.DefineType.</span>
<a id="L1111"></a>func NewNamedType(name string) *NamedType {
    <a id="L1112"></a>return &amp;NamedType{token.Position{}, name, nil, true, make(map[string]Method)}
<a id="L1113"></a>}

<a id="L1115"></a>func (t *NamedType) Complete(def Type) {
    <a id="L1116"></a>if !t.incomplete {
        <a id="L1117"></a>log.Crashf(&#34;cannot complete already completed NamedType %+v&#34;, *t)
    <a id="L1118"></a>}
    <a id="L1119"></a><span class="comment">// We strip the name from def because multiple levels of</span>
    <a id="L1120"></a><span class="comment">// naming are useless.</span>
    <a id="L1121"></a>if ndef, ok := def.(*NamedType); ok {
        <a id="L1122"></a>def = ndef.Def
    <a id="L1123"></a>}
    <a id="L1124"></a>t.Def = def;
    <a id="L1125"></a>t.incomplete = false;
<a id="L1126"></a>}

<a id="L1128"></a>func (t *NamedType) compat(o Type, conv bool) bool {
    <a id="L1129"></a>t2, ok := o.(*NamedType);
    <a id="L1130"></a>if ok {
        <a id="L1131"></a>if conv {
            <a id="L1132"></a><span class="comment">// Two named types are conversion compatible</span>
            <a id="L1133"></a><span class="comment">// if their literals are conversion</span>
            <a id="L1134"></a><span class="comment">// compatible.</span>
            <a id="L1135"></a>return t.Def.compat(t2.Def, conv)
        <a id="L1136"></a>} else {
            <a id="L1137"></a><span class="comment">// Two named types are compatible if their</span>
            <a id="L1138"></a><span class="comment">// type names originate in the same type</span>
            <a id="L1139"></a><span class="comment">// declaration.</span>
            <a id="L1140"></a>return t == t2
        <a id="L1141"></a>}
    <a id="L1142"></a>}
    <a id="L1143"></a><span class="comment">// A named and an unnamed type are compatible if the</span>
    <a id="L1144"></a><span class="comment">// respective type literals are compatible.</span>
    <a id="L1145"></a>return o.compat(t.Def, conv);
<a id="L1146"></a>}

<a id="L1148"></a>func (t *NamedType) lit() Type { return t.Def.lit() }

<a id="L1150"></a>func (t *NamedType) isBoolean() bool { return t.Def.isBoolean() }

<a id="L1152"></a>func (t *NamedType) isInteger() bool { return t.Def.isInteger() }

<a id="L1154"></a>func (t *NamedType) isFloat() bool { return t.Def.isFloat() }

<a id="L1156"></a>func (t *NamedType) isIdeal() bool { return false }

<a id="L1158"></a>func (t *NamedType) String() string { return t.Name }

<a id="L1160"></a>func (t *NamedType) Zero() Value { return t.Def.Zero() }

<a id="L1162"></a><span class="comment">/*</span>
<a id="L1163"></a><span class="comment"> * Multi-valued type</span>
<a id="L1164"></a><span class="comment"> */</span>

<a id="L1166"></a><span class="comment">// MultiType is a special type used for multi-valued expressions, akin</span>
<a id="L1167"></a><span class="comment">// to a tuple type.  It&#39;s not generally accessible within the</span>
<a id="L1168"></a><span class="comment">// language.</span>
<a id="L1169"></a>type MultiType struct {
    <a id="L1170"></a>commonType;
    <a id="L1171"></a>Elems []Type;
<a id="L1172"></a>}

<a id="L1174"></a>var multiTypes = newTypeArrayMap()

<a id="L1176"></a>func NewMultiType(elems []Type) *MultiType {
    <a id="L1177"></a>if t := multiTypes.Get(elems); t != nil {
        <a id="L1178"></a>return t.(*MultiType)
    <a id="L1179"></a>}

    <a id="L1181"></a>t := &amp;MultiType{commonType{}, elems};
    <a id="L1182"></a>multiTypes.Put(elems, t);
    <a id="L1183"></a>return t;
<a id="L1184"></a>}

<a id="L1186"></a>func (t *MultiType) compat(o Type, conv bool) bool {
    <a id="L1187"></a>t2, ok := o.lit().(*MultiType);
    <a id="L1188"></a>if !ok {
        <a id="L1189"></a>return false
    <a id="L1190"></a>}
    <a id="L1191"></a>if len(t.Elems) != len(t2.Elems) {
        <a id="L1192"></a>return false
    <a id="L1193"></a>}
    <a id="L1194"></a>for i := range t.Elems {
        <a id="L1195"></a>if !t.Elems[i].compat(t2.Elems[i], conv) {
            <a id="L1196"></a>return false
        <a id="L1197"></a>}
    <a id="L1198"></a>}
    <a id="L1199"></a>return true;
<a id="L1200"></a>}

<a id="L1202"></a>var EmptyType Type = NewMultiType([]Type{})

<a id="L1204"></a>func (t *MultiType) lit() Type { return t }

<a id="L1206"></a>func (t *MultiType) String() string {
    <a id="L1207"></a>if len(t.Elems) == 0 {
        <a id="L1208"></a>return &#34;&lt;none&gt;&#34;
    <a id="L1209"></a>}
    <a id="L1210"></a>return typeListString(t.Elems, nil);
<a id="L1211"></a>}

<a id="L1213"></a>func (t *MultiType) Zero() Value {
    <a id="L1214"></a>res := make([]Value, len(t.Elems));
    <a id="L1215"></a>for i, t := range t.Elems {
        <a id="L1216"></a>res[i] = t.Zero()
    <a id="L1217"></a>}
    <a id="L1218"></a>return multiV(res);
<a id="L1219"></a>}

<a id="L1221"></a><span class="comment">/*</span>
<a id="L1222"></a><span class="comment"> * Initialize the universe</span>
<a id="L1223"></a><span class="comment"> */</span>

<a id="L1225"></a>func init() {
    <a id="L1226"></a><span class="comment">// To avoid portability issues all numeric types are distinct</span>
    <a id="L1227"></a><span class="comment">// except byte, which is an alias for uint8.</span>

    <a id="L1229"></a><span class="comment">// Make byte an alias for the named type uint8.  Type aliases</span>
    <a id="L1230"></a><span class="comment">// are otherwise impossible in Go, so just hack it here.</span>
    <a id="L1231"></a>universe.defs[&#34;byte&#34;] = universe.defs[&#34;uint8&#34;];

    <a id="L1233"></a><span class="comment">// Built-in functions</span>
    <a id="L1234"></a>universe.DefineConst(&#34;cap&#34;, universePos, capType, nil);
    <a id="L1235"></a>universe.DefineConst(&#34;close&#34;, universePos, closeType, nil);
    <a id="L1236"></a>universe.DefineConst(&#34;closed&#34;, universePos, closedType, nil);
    <a id="L1237"></a>universe.DefineConst(&#34;len&#34;, universePos, lenType, nil);
    <a id="L1238"></a>universe.DefineConst(&#34;make&#34;, universePos, makeType, nil);
    <a id="L1239"></a>universe.DefineConst(&#34;new&#34;, universePos, newType, nil);
    <a id="L1240"></a>universe.DefineConst(&#34;panic&#34;, universePos, panicType, nil);
    <a id="L1241"></a>universe.DefineConst(&#34;panicln&#34;, universePos, paniclnType, nil);
    <a id="L1242"></a>universe.DefineConst(&#34;print&#34;, universePos, printType, nil);
    <a id="L1243"></a>universe.DefineConst(&#34;println&#34;, universePos, printlnType, nil);
<a id="L1244"></a>}
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
