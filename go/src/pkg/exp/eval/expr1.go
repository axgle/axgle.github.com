<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/expr1.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/expr1.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// This file is machine generated by gen.go.</span>
<a id="L2"></a><span class="comment">// 6g gen.go &amp;&amp; 6l gen.6 &amp;&amp; ./6.out &gt;expr1.go</span>

<a id="L4"></a>package eval

<a id="L6"></a>import (
    <a id="L7"></a>&#34;bignum&#34;;
    <a id="L8"></a>&#34;log&#34;;
<a id="L9"></a>)

<a id="L11"></a><span class="comment">/*</span>
<a id="L12"></a><span class="comment"> * &#34;As&#34; functions.  These retrieve evaluator functions from an</span>
<a id="L13"></a><span class="comment"> * expr, panicking if the requested evaluator has the wrong type.</span>
<a id="L14"></a><span class="comment"> */</span>
<a id="L15"></a>func (a *expr) asBool() (func(*Thread) bool) { return a.eval.(func(*Thread) bool) }
<a id="L16"></a>func (a *expr) asUint() (func(*Thread) uint64) {
    <a id="L17"></a>return a.eval.(func(*Thread) uint64)
<a id="L18"></a>}
<a id="L19"></a>func (a *expr) asInt() (func(*Thread) int64) { return a.eval.(func(*Thread) int64) }
<a id="L20"></a>func (a *expr) asIdealInt() (func() *bignum.Integer) {
    <a id="L21"></a>return a.eval.(func() *bignum.Integer)
<a id="L22"></a>}
<a id="L23"></a>func (a *expr) asFloat() (func(*Thread) float64) {
    <a id="L24"></a>return a.eval.(func(*Thread) float64)
<a id="L25"></a>}
<a id="L26"></a>func (a *expr) asIdealFloat() (func() *bignum.Rational) {
    <a id="L27"></a>return a.eval.(func() *bignum.Rational)
<a id="L28"></a>}
<a id="L29"></a>func (a *expr) asString() (func(*Thread) string) {
    <a id="L30"></a>return a.eval.(func(*Thread) string)
<a id="L31"></a>}
<a id="L32"></a>func (a *expr) asArray() (func(*Thread) ArrayValue) {
    <a id="L33"></a>return a.eval.(func(*Thread) ArrayValue)
<a id="L34"></a>}
<a id="L35"></a>func (a *expr) asStruct() (func(*Thread) StructValue) {
    <a id="L36"></a>return a.eval.(func(*Thread) StructValue)
<a id="L37"></a>}
<a id="L38"></a>func (a *expr) asPtr() (func(*Thread) Value) { return a.eval.(func(*Thread) Value) }
<a id="L39"></a>func (a *expr) asFunc() (func(*Thread) Func) { return a.eval.(func(*Thread) Func) }
<a id="L40"></a>func (a *expr) asSlice() (func(*Thread) Slice) {
    <a id="L41"></a>return a.eval.(func(*Thread) Slice)
<a id="L42"></a>}
<a id="L43"></a>func (a *expr) asMap() (func(*Thread) Map) { return a.eval.(func(*Thread) Map) }
<a id="L44"></a>func (a *expr) asMulti() (func(*Thread) []Value) {
    <a id="L45"></a>return a.eval.(func(*Thread) []Value)
<a id="L46"></a>}

<a id="L48"></a>func (a *expr) asInterface() (func(*Thread) interface{}) {
    <a id="L49"></a>switch sf := a.eval.(type) {
    <a id="L50"></a>case func(t *Thread) bool:
        <a id="L51"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L52"></a>case func(t *Thread) uint64:
        <a id="L53"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L54"></a>case func(t *Thread) int64:
        <a id="L55"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L56"></a>case func() *bignum.Integer:
        <a id="L57"></a>return func(*Thread) interface{} { return sf() }
    <a id="L58"></a>case func(t *Thread) float64:
        <a id="L59"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L60"></a>case func() *bignum.Rational:
        <a id="L61"></a>return func(*Thread) interface{} { return sf() }
    <a id="L62"></a>case func(t *Thread) string:
        <a id="L63"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L64"></a>case func(t *Thread) ArrayValue:
        <a id="L65"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L66"></a>case func(t *Thread) StructValue:
        <a id="L67"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L68"></a>case func(t *Thread) Value:
        <a id="L69"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L70"></a>case func(t *Thread) Func:
        <a id="L71"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L72"></a>case func(t *Thread) Slice:
        <a id="L73"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L74"></a>case func(t *Thread) Map:
        <a id="L75"></a>return func(t *Thread) interface{} { return sf(t) }
    <a id="L76"></a>default:
        <a id="L77"></a>log.Crashf(&#34;unexpected expression node type %T at %v&#34;, a.eval, a.pos)
    <a id="L78"></a>}
    <a id="L79"></a>panic();
<a id="L80"></a>}

<a id="L82"></a><span class="comment">/*</span>
<a id="L83"></a><span class="comment"> * Operator generators.</span>
<a id="L84"></a><span class="comment"> */</span>

<a id="L86"></a>func (a *expr) genConstant(v Value) {
    <a id="L87"></a>switch a.t.lit().(type) {
    <a id="L88"></a>case *boolType:
        <a id="L89"></a>a.eval = func(t *Thread) bool { return v.(BoolValue).Get(t) }
    <a id="L90"></a>case *uintType:
        <a id="L91"></a>a.eval = func(t *Thread) uint64 { return v.(UintValue).Get(t) }
    <a id="L92"></a>case *intType:
        <a id="L93"></a>a.eval = func(t *Thread) int64 { return v.(IntValue).Get(t) }
    <a id="L94"></a>case *idealIntType:
        <a id="L95"></a>val := v.(IdealIntValue).Get();
        <a id="L96"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L97"></a>case *floatType:
        <a id="L98"></a>a.eval = func(t *Thread) float64 { return v.(FloatValue).Get(t) }
    <a id="L99"></a>case *idealFloatType:
        <a id="L100"></a>val := v.(IdealFloatValue).Get();
        <a id="L101"></a>a.eval = func() *bignum.Rational { return val };
    <a id="L102"></a>case *stringType:
        <a id="L103"></a>a.eval = func(t *Thread) string { return v.(StringValue).Get(t) }
    <a id="L104"></a>case *ArrayType:
        <a id="L105"></a>a.eval = func(t *Thread) ArrayValue { return v.(ArrayValue).Get(t) }
    <a id="L106"></a>case *StructType:
        <a id="L107"></a>a.eval = func(t *Thread) StructValue { return v.(StructValue).Get(t) }
    <a id="L108"></a>case *PtrType:
        <a id="L109"></a>a.eval = func(t *Thread) Value { return v.(PtrValue).Get(t) }
    <a id="L110"></a>case *FuncType:
        <a id="L111"></a>a.eval = func(t *Thread) Func { return v.(FuncValue).Get(t) }
    <a id="L112"></a>case *SliceType:
        <a id="L113"></a>a.eval = func(t *Thread) Slice { return v.(SliceValue).Get(t) }
    <a id="L114"></a>case *MapType:
        <a id="L115"></a>a.eval = func(t *Thread) Map { return v.(MapValue).Get(t) }
    <a id="L116"></a>default:
        <a id="L117"></a>log.Crashf(&#34;unexpected constant type %v at %v&#34;, a.t, a.pos)
    <a id="L118"></a>}
<a id="L119"></a>}

<a id="L121"></a>func (a *expr) genIdentOp(level, index int) {
    <a id="L122"></a>a.evalAddr = func(t *Thread) Value { return t.f.Get(level, index) };
    <a id="L123"></a>switch a.t.lit().(type) {
    <a id="L124"></a>case *boolType:
        <a id="L125"></a>a.eval = func(t *Thread) bool { return t.f.Get(level, index).(BoolValue).Get(t) }
    <a id="L126"></a>case *uintType:
        <a id="L127"></a>a.eval = func(t *Thread) uint64 { return t.f.Get(level, index).(UintValue).Get(t) }
    <a id="L128"></a>case *intType:
        <a id="L129"></a>a.eval = func(t *Thread) int64 { return t.f.Get(level, index).(IntValue).Get(t) }
    <a id="L130"></a>case *floatType:
        <a id="L131"></a>a.eval = func(t *Thread) float64 { return t.f.Get(level, index).(FloatValue).Get(t) }
    <a id="L132"></a>case *stringType:
        <a id="L133"></a>a.eval = func(t *Thread) string { return t.f.Get(level, index).(StringValue).Get(t) }
    <a id="L134"></a>case *ArrayType:
        <a id="L135"></a>a.eval = func(t *Thread) ArrayValue { return t.f.Get(level, index).(ArrayValue).Get(t) }
    <a id="L136"></a>case *StructType:
        <a id="L137"></a>a.eval = func(t *Thread) StructValue { return t.f.Get(level, index).(StructValue).Get(t) }
    <a id="L138"></a>case *PtrType:
        <a id="L139"></a>a.eval = func(t *Thread) Value { return t.f.Get(level, index).(PtrValue).Get(t) }
    <a id="L140"></a>case *FuncType:
        <a id="L141"></a>a.eval = func(t *Thread) Func { return t.f.Get(level, index).(FuncValue).Get(t) }
    <a id="L142"></a>case *SliceType:
        <a id="L143"></a>a.eval = func(t *Thread) Slice { return t.f.Get(level, index).(SliceValue).Get(t) }
    <a id="L144"></a>case *MapType:
        <a id="L145"></a>a.eval = func(t *Thread) Map { return t.f.Get(level, index).(MapValue).Get(t) }
    <a id="L146"></a>default:
        <a id="L147"></a>log.Crashf(&#34;unexpected identifier type %v at %v&#34;, a.t, a.pos)
    <a id="L148"></a>}
<a id="L149"></a>}

<a id="L151"></a>func (a *expr) genFuncCall(call func(t *Thread) []Value) {
    <a id="L152"></a>a.exec = func(t *Thread) { call(t) };
    <a id="L153"></a>switch a.t.lit().(type) {
    <a id="L154"></a>case *boolType:
        <a id="L155"></a>a.eval = func(t *Thread) bool { return call(t)[0].(BoolValue).Get(t) }
    <a id="L156"></a>case *uintType:
        <a id="L157"></a>a.eval = func(t *Thread) uint64 { return call(t)[0].(UintValue).Get(t) }
    <a id="L158"></a>case *intType:
        <a id="L159"></a>a.eval = func(t *Thread) int64 { return call(t)[0].(IntValue).Get(t) }
    <a id="L160"></a>case *floatType:
        <a id="L161"></a>a.eval = func(t *Thread) float64 { return call(t)[0].(FloatValue).Get(t) }
    <a id="L162"></a>case *stringType:
        <a id="L163"></a>a.eval = func(t *Thread) string { return call(t)[0].(StringValue).Get(t) }
    <a id="L164"></a>case *ArrayType:
        <a id="L165"></a>a.eval = func(t *Thread) ArrayValue { return call(t)[0].(ArrayValue).Get(t) }
    <a id="L166"></a>case *StructType:
        <a id="L167"></a>a.eval = func(t *Thread) StructValue { return call(t)[0].(StructValue).Get(t) }
    <a id="L168"></a>case *PtrType:
        <a id="L169"></a>a.eval = func(t *Thread) Value { return call(t)[0].(PtrValue).Get(t) }
    <a id="L170"></a>case *FuncType:
        <a id="L171"></a>a.eval = func(t *Thread) Func { return call(t)[0].(FuncValue).Get(t) }
    <a id="L172"></a>case *SliceType:
        <a id="L173"></a>a.eval = func(t *Thread) Slice { return call(t)[0].(SliceValue).Get(t) }
    <a id="L174"></a>case *MapType:
        <a id="L175"></a>a.eval = func(t *Thread) Map { return call(t)[0].(MapValue).Get(t) }
    <a id="L176"></a>case *MultiType:
        <a id="L177"></a>a.eval = func(t *Thread) []Value { return call(t) }
    <a id="L178"></a>default:
        <a id="L179"></a>log.Crashf(&#34;unexpected result type %v at %v&#34;, a.t, a.pos)
    <a id="L180"></a>}
<a id="L181"></a>}

<a id="L183"></a>func (a *expr) genValue(vf func(*Thread) Value) {
    <a id="L184"></a>a.evalAddr = vf;
    <a id="L185"></a>switch a.t.lit().(type) {
    <a id="L186"></a>case *boolType:
        <a id="L187"></a>a.eval = func(t *Thread) bool { return vf(t).(BoolValue).Get(t) }
    <a id="L188"></a>case *uintType:
        <a id="L189"></a>a.eval = func(t *Thread) uint64 { return vf(t).(UintValue).Get(t) }
    <a id="L190"></a>case *intType:
        <a id="L191"></a>a.eval = func(t *Thread) int64 { return vf(t).(IntValue).Get(t) }
    <a id="L192"></a>case *floatType:
        <a id="L193"></a>a.eval = func(t *Thread) float64 { return vf(t).(FloatValue).Get(t) }
    <a id="L194"></a>case *stringType:
        <a id="L195"></a>a.eval = func(t *Thread) string { return vf(t).(StringValue).Get(t) }
    <a id="L196"></a>case *ArrayType:
        <a id="L197"></a>a.eval = func(t *Thread) ArrayValue { return vf(t).(ArrayValue).Get(t) }
    <a id="L198"></a>case *StructType:
        <a id="L199"></a>a.eval = func(t *Thread) StructValue { return vf(t).(StructValue).Get(t) }
    <a id="L200"></a>case *PtrType:
        <a id="L201"></a>a.eval = func(t *Thread) Value { return vf(t).(PtrValue).Get(t) }
    <a id="L202"></a>case *FuncType:
        <a id="L203"></a>a.eval = func(t *Thread) Func { return vf(t).(FuncValue).Get(t) }
    <a id="L204"></a>case *SliceType:
        <a id="L205"></a>a.eval = func(t *Thread) Slice { return vf(t).(SliceValue).Get(t) }
    <a id="L206"></a>case *MapType:
        <a id="L207"></a>a.eval = func(t *Thread) Map { return vf(t).(MapValue).Get(t) }
    <a id="L208"></a>default:
        <a id="L209"></a>log.Crashf(&#34;unexpected result type %v at %v&#34;, a.t, a.pos)
    <a id="L210"></a>}
<a id="L211"></a>}

<a id="L213"></a>func (a *expr) genUnaryOpNeg(v *expr) {
    <a id="L214"></a>switch a.t.lit().(type) {
    <a id="L215"></a>case *uintType:
        <a id="L216"></a>vf := v.asUint();
        <a id="L217"></a>a.eval = func(t *Thread) uint64 {
            <a id="L218"></a>v := vf(t);
            <a id="L219"></a>return -v;
        <a id="L220"></a>};
    <a id="L221"></a>case *intType:
        <a id="L222"></a>vf := v.asInt();
        <a id="L223"></a>a.eval = func(t *Thread) int64 {
            <a id="L224"></a>v := vf(t);
            <a id="L225"></a>return -v;
        <a id="L226"></a>};
    <a id="L227"></a>case *idealIntType:
        <a id="L228"></a>v := v.asIdealInt()();
        <a id="L229"></a>val := v.Neg();
        <a id="L230"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L231"></a>case *floatType:
        <a id="L232"></a>vf := v.asFloat();
        <a id="L233"></a>a.eval = func(t *Thread) float64 {
            <a id="L234"></a>v := vf(t);
            <a id="L235"></a>return -v;
        <a id="L236"></a>};
    <a id="L237"></a>case *idealFloatType:
        <a id="L238"></a>v := v.asIdealFloat()();
        <a id="L239"></a>val := v.Neg();
        <a id="L240"></a>a.eval = func() *bignum.Rational { return val };
    <a id="L241"></a>default:
        <a id="L242"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, a.t, a.pos)
    <a id="L243"></a>}
<a id="L244"></a>}

<a id="L246"></a>func (a *expr) genUnaryOpNot(v *expr) {
    <a id="L247"></a>switch a.t.lit().(type) {
    <a id="L248"></a>case *boolType:
        <a id="L249"></a>vf := v.asBool();
        <a id="L250"></a>a.eval = func(t *Thread) bool {
            <a id="L251"></a>v := vf(t);
            <a id="L252"></a>return !v;
        <a id="L253"></a>};
    <a id="L254"></a>default:
        <a id="L255"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, a.t, a.pos)
    <a id="L256"></a>}
<a id="L257"></a>}

<a id="L259"></a>func (a *expr) genUnaryOpXor(v *expr) {
    <a id="L260"></a>switch a.t.lit().(type) {
    <a id="L261"></a>case *uintType:
        <a id="L262"></a>vf := v.asUint();
        <a id="L263"></a>a.eval = func(t *Thread) uint64 {
            <a id="L264"></a>v := vf(t);
            <a id="L265"></a>return ^v;
        <a id="L266"></a>};
    <a id="L267"></a>case *intType:
        <a id="L268"></a>vf := v.asInt();
        <a id="L269"></a>a.eval = func(t *Thread) int64 {
            <a id="L270"></a>v := vf(t);
            <a id="L271"></a>return ^v;
        <a id="L272"></a>};
    <a id="L273"></a>case *idealIntType:
        <a id="L274"></a>v := v.asIdealInt()();
        <a id="L275"></a>val := v.Neg().Sub(bignum.Int(1));
        <a id="L276"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L277"></a>default:
        <a id="L278"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, a.t, a.pos)
    <a id="L279"></a>}
<a id="L280"></a>}

<a id="L282"></a>func (a *expr) genBinOpLogAnd(l, r *expr) {
    <a id="L283"></a>lf := l.asBool();
    <a id="L284"></a>rf := r.asBool();
    <a id="L285"></a>a.eval = func(t *Thread) bool { return lf(t) &amp;&amp; rf(t) };
<a id="L286"></a>}

<a id="L288"></a>func (a *expr) genBinOpLogOr(l, r *expr) {
    <a id="L289"></a>lf := l.asBool();
    <a id="L290"></a>rf := r.asBool();
    <a id="L291"></a>a.eval = func(t *Thread) bool { return lf(t) || rf(t) };
<a id="L292"></a>}

<a id="L294"></a>func (a *expr) genBinOpAdd(l, r *expr) {
    <a id="L295"></a>switch t := l.t.lit().(type) {
    <a id="L296"></a>case *uintType:
        <a id="L297"></a>lf := l.asUint();
        <a id="L298"></a>rf := r.asUint();
        <a id="L299"></a>switch t.Bits {
        <a id="L300"></a>case 8:
            <a id="L301"></a>a.eval = func(t *Thread) uint64 {
                <a id="L302"></a>l, r := lf(t), rf(t);
                <a id="L303"></a>var ret uint64;
                <a id="L304"></a>ret = l + r;
                <a id="L305"></a>return uint64(uint8(ret));
            <a id="L306"></a>}
        <a id="L307"></a>case 16:
            <a id="L308"></a>a.eval = func(t *Thread) uint64 {
                <a id="L309"></a>l, r := lf(t), rf(t);
                <a id="L310"></a>var ret uint64;
                <a id="L311"></a>ret = l + r;
                <a id="L312"></a>return uint64(uint16(ret));
            <a id="L313"></a>}
        <a id="L314"></a>case 32:
            <a id="L315"></a>a.eval = func(t *Thread) uint64 {
                <a id="L316"></a>l, r := lf(t), rf(t);
                <a id="L317"></a>var ret uint64;
                <a id="L318"></a>ret = l + r;
                <a id="L319"></a>return uint64(uint32(ret));
            <a id="L320"></a>}
        <a id="L321"></a>case 64:
            <a id="L322"></a>a.eval = func(t *Thread) uint64 {
                <a id="L323"></a>l, r := lf(t), rf(t);
                <a id="L324"></a>var ret uint64;
                <a id="L325"></a>ret = l + r;
                <a id="L326"></a>return uint64(uint64(ret));
            <a id="L327"></a>}
        <a id="L328"></a>case 0:
            <a id="L329"></a>a.eval = func(t *Thread) uint64 {
                <a id="L330"></a>l, r := lf(t), rf(t);
                <a id="L331"></a>var ret uint64;
                <a id="L332"></a>ret = l + r;
                <a id="L333"></a>return uint64(uint(ret));
            <a id="L334"></a>}
        <a id="L335"></a>default:
            <a id="L336"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L337"></a>}
    <a id="L338"></a>case *intType:
        <a id="L339"></a>lf := l.asInt();
        <a id="L340"></a>rf := r.asInt();
        <a id="L341"></a>switch t.Bits {
        <a id="L342"></a>case 8:
            <a id="L343"></a>a.eval = func(t *Thread) int64 {
                <a id="L344"></a>l, r := lf(t), rf(t);
                <a id="L345"></a>var ret int64;
                <a id="L346"></a>ret = l + r;
                <a id="L347"></a>return int64(int8(ret));
            <a id="L348"></a>}
        <a id="L349"></a>case 16:
            <a id="L350"></a>a.eval = func(t *Thread) int64 {
                <a id="L351"></a>l, r := lf(t), rf(t);
                <a id="L352"></a>var ret int64;
                <a id="L353"></a>ret = l + r;
                <a id="L354"></a>return int64(int16(ret));
            <a id="L355"></a>}
        <a id="L356"></a>case 32:
            <a id="L357"></a>a.eval = func(t *Thread) int64 {
                <a id="L358"></a>l, r := lf(t), rf(t);
                <a id="L359"></a>var ret int64;
                <a id="L360"></a>ret = l + r;
                <a id="L361"></a>return int64(int32(ret));
            <a id="L362"></a>}
        <a id="L363"></a>case 64:
            <a id="L364"></a>a.eval = func(t *Thread) int64 {
                <a id="L365"></a>l, r := lf(t), rf(t);
                <a id="L366"></a>var ret int64;
                <a id="L367"></a>ret = l + r;
                <a id="L368"></a>return int64(int64(ret));
            <a id="L369"></a>}
        <a id="L370"></a>case 0:
            <a id="L371"></a>a.eval = func(t *Thread) int64 {
                <a id="L372"></a>l, r := lf(t), rf(t);
                <a id="L373"></a>var ret int64;
                <a id="L374"></a>ret = l + r;
                <a id="L375"></a>return int64(int(ret));
            <a id="L376"></a>}
        <a id="L377"></a>default:
            <a id="L378"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L379"></a>}
    <a id="L380"></a>case *idealIntType:
        <a id="L381"></a>l := l.asIdealInt()();
        <a id="L382"></a>r := r.asIdealInt()();
        <a id="L383"></a>val := l.Add(r);
        <a id="L384"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L385"></a>case *floatType:
        <a id="L386"></a>lf := l.asFloat();
        <a id="L387"></a>rf := r.asFloat();
        <a id="L388"></a>switch t.Bits {
        <a id="L389"></a>case 32:
            <a id="L390"></a>a.eval = func(t *Thread) float64 {
                <a id="L391"></a>l, r := lf(t), rf(t);
                <a id="L392"></a>var ret float64;
                <a id="L393"></a>ret = l + r;
                <a id="L394"></a>return float64(float32(ret));
            <a id="L395"></a>}
        <a id="L396"></a>case 64:
            <a id="L397"></a>a.eval = func(t *Thread) float64 {
                <a id="L398"></a>l, r := lf(t), rf(t);
                <a id="L399"></a>var ret float64;
                <a id="L400"></a>ret = l + r;
                <a id="L401"></a>return float64(float64(ret));
            <a id="L402"></a>}
        <a id="L403"></a>case 0:
            <a id="L404"></a>a.eval = func(t *Thread) float64 {
                <a id="L405"></a>l, r := lf(t), rf(t);
                <a id="L406"></a>var ret float64;
                <a id="L407"></a>ret = l + r;
                <a id="L408"></a>return float64(float(ret));
            <a id="L409"></a>}
        <a id="L410"></a>default:
            <a id="L411"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L412"></a>}
    <a id="L413"></a>case *idealFloatType:
        <a id="L414"></a>l := l.asIdealFloat()();
        <a id="L415"></a>r := r.asIdealFloat()();
        <a id="L416"></a>val := l.Add(r);
        <a id="L417"></a>a.eval = func() *bignum.Rational { return val };
    <a id="L418"></a>case *stringType:
        <a id="L419"></a>lf := l.asString();
        <a id="L420"></a>rf := r.asString();
        <a id="L421"></a>a.eval = func(t *Thread) string {
            <a id="L422"></a>l, r := lf(t), rf(t);
            <a id="L423"></a>return l + r;
        <a id="L424"></a>};
    <a id="L425"></a>default:
        <a id="L426"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L427"></a>}
<a id="L428"></a>}

<a id="L430"></a>func (a *expr) genBinOpSub(l, r *expr) {
    <a id="L431"></a>switch t := l.t.lit().(type) {
    <a id="L432"></a>case *uintType:
        <a id="L433"></a>lf := l.asUint();
        <a id="L434"></a>rf := r.asUint();
        <a id="L435"></a>switch t.Bits {
        <a id="L436"></a>case 8:
            <a id="L437"></a>a.eval = func(t *Thread) uint64 {
                <a id="L438"></a>l, r := lf(t), rf(t);
                <a id="L439"></a>var ret uint64;
                <a id="L440"></a>ret = l - r;
                <a id="L441"></a>return uint64(uint8(ret));
            <a id="L442"></a>}
        <a id="L443"></a>case 16:
            <a id="L444"></a>a.eval = func(t *Thread) uint64 {
                <a id="L445"></a>l, r := lf(t), rf(t);
                <a id="L446"></a>var ret uint64;
                <a id="L447"></a>ret = l - r;
                <a id="L448"></a>return uint64(uint16(ret));
            <a id="L449"></a>}
        <a id="L450"></a>case 32:
            <a id="L451"></a>a.eval = func(t *Thread) uint64 {
                <a id="L452"></a>l, r := lf(t), rf(t);
                <a id="L453"></a>var ret uint64;
                <a id="L454"></a>ret = l - r;
                <a id="L455"></a>return uint64(uint32(ret));
            <a id="L456"></a>}
        <a id="L457"></a>case 64:
            <a id="L458"></a>a.eval = func(t *Thread) uint64 {
                <a id="L459"></a>l, r := lf(t), rf(t);
                <a id="L460"></a>var ret uint64;
                <a id="L461"></a>ret = l - r;
                <a id="L462"></a>return uint64(uint64(ret));
            <a id="L463"></a>}
        <a id="L464"></a>case 0:
            <a id="L465"></a>a.eval = func(t *Thread) uint64 {
                <a id="L466"></a>l, r := lf(t), rf(t);
                <a id="L467"></a>var ret uint64;
                <a id="L468"></a>ret = l - r;
                <a id="L469"></a>return uint64(uint(ret));
            <a id="L470"></a>}
        <a id="L471"></a>default:
            <a id="L472"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L473"></a>}
    <a id="L474"></a>case *intType:
        <a id="L475"></a>lf := l.asInt();
        <a id="L476"></a>rf := r.asInt();
        <a id="L477"></a>switch t.Bits {
        <a id="L478"></a>case 8:
            <a id="L479"></a>a.eval = func(t *Thread) int64 {
                <a id="L480"></a>l, r := lf(t), rf(t);
                <a id="L481"></a>var ret int64;
                <a id="L482"></a>ret = l - r;
                <a id="L483"></a>return int64(int8(ret));
            <a id="L484"></a>}
        <a id="L485"></a>case 16:
            <a id="L486"></a>a.eval = func(t *Thread) int64 {
                <a id="L487"></a>l, r := lf(t), rf(t);
                <a id="L488"></a>var ret int64;
                <a id="L489"></a>ret = l - r;
                <a id="L490"></a>return int64(int16(ret));
            <a id="L491"></a>}
        <a id="L492"></a>case 32:
            <a id="L493"></a>a.eval = func(t *Thread) int64 {
                <a id="L494"></a>l, r := lf(t), rf(t);
                <a id="L495"></a>var ret int64;
                <a id="L496"></a>ret = l - r;
                <a id="L497"></a>return int64(int32(ret));
            <a id="L498"></a>}
        <a id="L499"></a>case 64:
            <a id="L500"></a>a.eval = func(t *Thread) int64 {
                <a id="L501"></a>l, r := lf(t), rf(t);
                <a id="L502"></a>var ret int64;
                <a id="L503"></a>ret = l - r;
                <a id="L504"></a>return int64(int64(ret));
            <a id="L505"></a>}
        <a id="L506"></a>case 0:
            <a id="L507"></a>a.eval = func(t *Thread) int64 {
                <a id="L508"></a>l, r := lf(t), rf(t);
                <a id="L509"></a>var ret int64;
                <a id="L510"></a>ret = l - r;
                <a id="L511"></a>return int64(int(ret));
            <a id="L512"></a>}
        <a id="L513"></a>default:
            <a id="L514"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L515"></a>}
    <a id="L516"></a>case *idealIntType:
        <a id="L517"></a>l := l.asIdealInt()();
        <a id="L518"></a>r := r.asIdealInt()();
        <a id="L519"></a>val := l.Sub(r);
        <a id="L520"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L521"></a>case *floatType:
        <a id="L522"></a>lf := l.asFloat();
        <a id="L523"></a>rf := r.asFloat();
        <a id="L524"></a>switch t.Bits {
        <a id="L525"></a>case 32:
            <a id="L526"></a>a.eval = func(t *Thread) float64 {
                <a id="L527"></a>l, r := lf(t), rf(t);
                <a id="L528"></a>var ret float64;
                <a id="L529"></a>ret = l - r;
                <a id="L530"></a>return float64(float32(ret));
            <a id="L531"></a>}
        <a id="L532"></a>case 64:
            <a id="L533"></a>a.eval = func(t *Thread) float64 {
                <a id="L534"></a>l, r := lf(t), rf(t);
                <a id="L535"></a>var ret float64;
                <a id="L536"></a>ret = l - r;
                <a id="L537"></a>return float64(float64(ret));
            <a id="L538"></a>}
        <a id="L539"></a>case 0:
            <a id="L540"></a>a.eval = func(t *Thread) float64 {
                <a id="L541"></a>l, r := lf(t), rf(t);
                <a id="L542"></a>var ret float64;
                <a id="L543"></a>ret = l - r;
                <a id="L544"></a>return float64(float(ret));
            <a id="L545"></a>}
        <a id="L546"></a>default:
            <a id="L547"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L548"></a>}
    <a id="L549"></a>case *idealFloatType:
        <a id="L550"></a>l := l.asIdealFloat()();
        <a id="L551"></a>r := r.asIdealFloat()();
        <a id="L552"></a>val := l.Sub(r);
        <a id="L553"></a>a.eval = func() *bignum.Rational { return val };
    <a id="L554"></a>default:
        <a id="L555"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L556"></a>}
<a id="L557"></a>}

<a id="L559"></a>func (a *expr) genBinOpMul(l, r *expr) {
    <a id="L560"></a>switch t := l.t.lit().(type) {
    <a id="L561"></a>case *uintType:
        <a id="L562"></a>lf := l.asUint();
        <a id="L563"></a>rf := r.asUint();
        <a id="L564"></a>switch t.Bits {
        <a id="L565"></a>case 8:
            <a id="L566"></a>a.eval = func(t *Thread) uint64 {
                <a id="L567"></a>l, r := lf(t), rf(t);
                <a id="L568"></a>var ret uint64;
                <a id="L569"></a>ret = l * r;
                <a id="L570"></a>return uint64(uint8(ret));
            <a id="L571"></a>}
        <a id="L572"></a>case 16:
            <a id="L573"></a>a.eval = func(t *Thread) uint64 {
                <a id="L574"></a>l, r := lf(t), rf(t);
                <a id="L575"></a>var ret uint64;
                <a id="L576"></a>ret = l * r;
                <a id="L577"></a>return uint64(uint16(ret));
            <a id="L578"></a>}
        <a id="L579"></a>case 32:
            <a id="L580"></a>a.eval = func(t *Thread) uint64 {
                <a id="L581"></a>l, r := lf(t), rf(t);
                <a id="L582"></a>var ret uint64;
                <a id="L583"></a>ret = l * r;
                <a id="L584"></a>return uint64(uint32(ret));
            <a id="L585"></a>}
        <a id="L586"></a>case 64:
            <a id="L587"></a>a.eval = func(t *Thread) uint64 {
                <a id="L588"></a>l, r := lf(t), rf(t);
                <a id="L589"></a>var ret uint64;
                <a id="L590"></a>ret = l * r;
                <a id="L591"></a>return uint64(uint64(ret));
            <a id="L592"></a>}
        <a id="L593"></a>case 0:
            <a id="L594"></a>a.eval = func(t *Thread) uint64 {
                <a id="L595"></a>l, r := lf(t), rf(t);
                <a id="L596"></a>var ret uint64;
                <a id="L597"></a>ret = l * r;
                <a id="L598"></a>return uint64(uint(ret));
            <a id="L599"></a>}
        <a id="L600"></a>default:
            <a id="L601"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L602"></a>}
    <a id="L603"></a>case *intType:
        <a id="L604"></a>lf := l.asInt();
        <a id="L605"></a>rf := r.asInt();
        <a id="L606"></a>switch t.Bits {
        <a id="L607"></a>case 8:
            <a id="L608"></a>a.eval = func(t *Thread) int64 {
                <a id="L609"></a>l, r := lf(t), rf(t);
                <a id="L610"></a>var ret int64;
                <a id="L611"></a>ret = l * r;
                <a id="L612"></a>return int64(int8(ret));
            <a id="L613"></a>}
        <a id="L614"></a>case 16:
            <a id="L615"></a>a.eval = func(t *Thread) int64 {
                <a id="L616"></a>l, r := lf(t), rf(t);
                <a id="L617"></a>var ret int64;
                <a id="L618"></a>ret = l * r;
                <a id="L619"></a>return int64(int16(ret));
            <a id="L620"></a>}
        <a id="L621"></a>case 32:
            <a id="L622"></a>a.eval = func(t *Thread) int64 {
                <a id="L623"></a>l, r := lf(t), rf(t);
                <a id="L624"></a>var ret int64;
                <a id="L625"></a>ret = l * r;
                <a id="L626"></a>return int64(int32(ret));
            <a id="L627"></a>}
        <a id="L628"></a>case 64:
            <a id="L629"></a>a.eval = func(t *Thread) int64 {
                <a id="L630"></a>l, r := lf(t), rf(t);
                <a id="L631"></a>var ret int64;
                <a id="L632"></a>ret = l * r;
                <a id="L633"></a>return int64(int64(ret));
            <a id="L634"></a>}
        <a id="L635"></a>case 0:
            <a id="L636"></a>a.eval = func(t *Thread) int64 {
                <a id="L637"></a>l, r := lf(t), rf(t);
                <a id="L638"></a>var ret int64;
                <a id="L639"></a>ret = l * r;
                <a id="L640"></a>return int64(int(ret));
            <a id="L641"></a>}
        <a id="L642"></a>default:
            <a id="L643"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L644"></a>}
    <a id="L645"></a>case *idealIntType:
        <a id="L646"></a>l := l.asIdealInt()();
        <a id="L647"></a>r := r.asIdealInt()();
        <a id="L648"></a>val := l.Mul(r);
        <a id="L649"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L650"></a>case *floatType:
        <a id="L651"></a>lf := l.asFloat();
        <a id="L652"></a>rf := r.asFloat();
        <a id="L653"></a>switch t.Bits {
        <a id="L654"></a>case 32:
            <a id="L655"></a>a.eval = func(t *Thread) float64 {
                <a id="L656"></a>l, r := lf(t), rf(t);
                <a id="L657"></a>var ret float64;
                <a id="L658"></a>ret = l * r;
                <a id="L659"></a>return float64(float32(ret));
            <a id="L660"></a>}
        <a id="L661"></a>case 64:
            <a id="L662"></a>a.eval = func(t *Thread) float64 {
                <a id="L663"></a>l, r := lf(t), rf(t);
                <a id="L664"></a>var ret float64;
                <a id="L665"></a>ret = l * r;
                <a id="L666"></a>return float64(float64(ret));
            <a id="L667"></a>}
        <a id="L668"></a>case 0:
            <a id="L669"></a>a.eval = func(t *Thread) float64 {
                <a id="L670"></a>l, r := lf(t), rf(t);
                <a id="L671"></a>var ret float64;
                <a id="L672"></a>ret = l * r;
                <a id="L673"></a>return float64(float(ret));
            <a id="L674"></a>}
        <a id="L675"></a>default:
            <a id="L676"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L677"></a>}
    <a id="L678"></a>case *idealFloatType:
        <a id="L679"></a>l := l.asIdealFloat()();
        <a id="L680"></a>r := r.asIdealFloat()();
        <a id="L681"></a>val := l.Mul(r);
        <a id="L682"></a>a.eval = func() *bignum.Rational { return val };
    <a id="L683"></a>default:
        <a id="L684"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L685"></a>}
<a id="L686"></a>}

<a id="L688"></a>func (a *expr) genBinOpQuo(l, r *expr) {
    <a id="L689"></a>switch t := l.t.lit().(type) {
    <a id="L690"></a>case *uintType:
        <a id="L691"></a>lf := l.asUint();
        <a id="L692"></a>rf := r.asUint();
        <a id="L693"></a>switch t.Bits {
        <a id="L694"></a>case 8:
            <a id="L695"></a>a.eval = func(t *Thread) uint64 {
                <a id="L696"></a>l, r := lf(t), rf(t);
                <a id="L697"></a>var ret uint64;
                <a id="L698"></a>if r == 0 {
                    <a id="L699"></a>t.Abort(DivByZeroError{})
                <a id="L700"></a>}
                <a id="L701"></a>ret = l / r;
                <a id="L702"></a>return uint64(uint8(ret));
            <a id="L703"></a>}
        <a id="L704"></a>case 16:
            <a id="L705"></a>a.eval = func(t *Thread) uint64 {
                <a id="L706"></a>l, r := lf(t), rf(t);
                <a id="L707"></a>var ret uint64;
                <a id="L708"></a>if r == 0 {
                    <a id="L709"></a>t.Abort(DivByZeroError{})
                <a id="L710"></a>}
                <a id="L711"></a>ret = l / r;
                <a id="L712"></a>return uint64(uint16(ret));
            <a id="L713"></a>}
        <a id="L714"></a>case 32:
            <a id="L715"></a>a.eval = func(t *Thread) uint64 {
                <a id="L716"></a>l, r := lf(t), rf(t);
                <a id="L717"></a>var ret uint64;
                <a id="L718"></a>if r == 0 {
                    <a id="L719"></a>t.Abort(DivByZeroError{})
                <a id="L720"></a>}
                <a id="L721"></a>ret = l / r;
                <a id="L722"></a>return uint64(uint32(ret));
            <a id="L723"></a>}
        <a id="L724"></a>case 64:
            <a id="L725"></a>a.eval = func(t *Thread) uint64 {
                <a id="L726"></a>l, r := lf(t), rf(t);
                <a id="L727"></a>var ret uint64;
                <a id="L728"></a>if r == 0 {
                    <a id="L729"></a>t.Abort(DivByZeroError{})
                <a id="L730"></a>}
                <a id="L731"></a>ret = l / r;
                <a id="L732"></a>return uint64(uint64(ret));
            <a id="L733"></a>}
        <a id="L734"></a>case 0:
            <a id="L735"></a>a.eval = func(t *Thread) uint64 {
                <a id="L736"></a>l, r := lf(t), rf(t);
                <a id="L737"></a>var ret uint64;
                <a id="L738"></a>if r == 0 {
                    <a id="L739"></a>t.Abort(DivByZeroError{})
                <a id="L740"></a>}
                <a id="L741"></a>ret = l / r;
                <a id="L742"></a>return uint64(uint(ret));
            <a id="L743"></a>}
        <a id="L744"></a>default:
            <a id="L745"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L746"></a>}
    <a id="L747"></a>case *intType:
        <a id="L748"></a>lf := l.asInt();
        <a id="L749"></a>rf := r.asInt();
        <a id="L750"></a>switch t.Bits {
        <a id="L751"></a>case 8:
            <a id="L752"></a>a.eval = func(t *Thread) int64 {
                <a id="L753"></a>l, r := lf(t), rf(t);
                <a id="L754"></a>var ret int64;
                <a id="L755"></a>if r == 0 {
                    <a id="L756"></a>t.Abort(DivByZeroError{})
                <a id="L757"></a>}
                <a id="L758"></a>ret = l / r;
                <a id="L759"></a>return int64(int8(ret));
            <a id="L760"></a>}
        <a id="L761"></a>case 16:
            <a id="L762"></a>a.eval = func(t *Thread) int64 {
                <a id="L763"></a>l, r := lf(t), rf(t);
                <a id="L764"></a>var ret int64;
                <a id="L765"></a>if r == 0 {
                    <a id="L766"></a>t.Abort(DivByZeroError{})
                <a id="L767"></a>}
                <a id="L768"></a>ret = l / r;
                <a id="L769"></a>return int64(int16(ret));
            <a id="L770"></a>}
        <a id="L771"></a>case 32:
            <a id="L772"></a>a.eval = func(t *Thread) int64 {
                <a id="L773"></a>l, r := lf(t), rf(t);
                <a id="L774"></a>var ret int64;
                <a id="L775"></a>if r == 0 {
                    <a id="L776"></a>t.Abort(DivByZeroError{})
                <a id="L777"></a>}
                <a id="L778"></a>ret = l / r;
                <a id="L779"></a>return int64(int32(ret));
            <a id="L780"></a>}
        <a id="L781"></a>case 64:
            <a id="L782"></a>a.eval = func(t *Thread) int64 {
                <a id="L783"></a>l, r := lf(t), rf(t);
                <a id="L784"></a>var ret int64;
                <a id="L785"></a>if r == 0 {
                    <a id="L786"></a>t.Abort(DivByZeroError{})
                <a id="L787"></a>}
                <a id="L788"></a>ret = l / r;
                <a id="L789"></a>return int64(int64(ret));
            <a id="L790"></a>}
        <a id="L791"></a>case 0:
            <a id="L792"></a>a.eval = func(t *Thread) int64 {
                <a id="L793"></a>l, r := lf(t), rf(t);
                <a id="L794"></a>var ret int64;
                <a id="L795"></a>if r == 0 {
                    <a id="L796"></a>t.Abort(DivByZeroError{})
                <a id="L797"></a>}
                <a id="L798"></a>ret = l / r;
                <a id="L799"></a>return int64(int(ret));
            <a id="L800"></a>}
        <a id="L801"></a>default:
            <a id="L802"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L803"></a>}
    <a id="L804"></a>case *idealIntType:
        <a id="L805"></a>l := l.asIdealInt()();
        <a id="L806"></a>r := r.asIdealInt()();
        <a id="L807"></a>val := l.Quo(r);
        <a id="L808"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L809"></a>case *floatType:
        <a id="L810"></a>lf := l.asFloat();
        <a id="L811"></a>rf := r.asFloat();
        <a id="L812"></a>switch t.Bits {
        <a id="L813"></a>case 32:
            <a id="L814"></a>a.eval = func(t *Thread) float64 {
                <a id="L815"></a>l, r := lf(t), rf(t);
                <a id="L816"></a>var ret float64;
                <a id="L817"></a>if r == 0 {
                    <a id="L818"></a>t.Abort(DivByZeroError{})
                <a id="L819"></a>}
                <a id="L820"></a>ret = l / r;
                <a id="L821"></a>return float64(float32(ret));
            <a id="L822"></a>}
        <a id="L823"></a>case 64:
            <a id="L824"></a>a.eval = func(t *Thread) float64 {
                <a id="L825"></a>l, r := lf(t), rf(t);
                <a id="L826"></a>var ret float64;
                <a id="L827"></a>if r == 0 {
                    <a id="L828"></a>t.Abort(DivByZeroError{})
                <a id="L829"></a>}
                <a id="L830"></a>ret = l / r;
                <a id="L831"></a>return float64(float64(ret));
            <a id="L832"></a>}
        <a id="L833"></a>case 0:
            <a id="L834"></a>a.eval = func(t *Thread) float64 {
                <a id="L835"></a>l, r := lf(t), rf(t);
                <a id="L836"></a>var ret float64;
                <a id="L837"></a>if r == 0 {
                    <a id="L838"></a>t.Abort(DivByZeroError{})
                <a id="L839"></a>}
                <a id="L840"></a>ret = l / r;
                <a id="L841"></a>return float64(float(ret));
            <a id="L842"></a>}
        <a id="L843"></a>default:
            <a id="L844"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L845"></a>}
    <a id="L846"></a>case *idealFloatType:
        <a id="L847"></a>l := l.asIdealFloat()();
        <a id="L848"></a>r := r.asIdealFloat()();
        <a id="L849"></a>val := l.Quo(r);
        <a id="L850"></a>a.eval = func() *bignum.Rational { return val };
    <a id="L851"></a>default:
        <a id="L852"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L853"></a>}
<a id="L854"></a>}

<a id="L856"></a>func (a *expr) genBinOpRem(l, r *expr) {
    <a id="L857"></a>switch t := l.t.lit().(type) {
    <a id="L858"></a>case *uintType:
        <a id="L859"></a>lf := l.asUint();
        <a id="L860"></a>rf := r.asUint();
        <a id="L861"></a>switch t.Bits {
        <a id="L862"></a>case 8:
            <a id="L863"></a>a.eval = func(t *Thread) uint64 {
                <a id="L864"></a>l, r := lf(t), rf(t);
                <a id="L865"></a>var ret uint64;
                <a id="L866"></a>if r == 0 {
                    <a id="L867"></a>t.Abort(DivByZeroError{})
                <a id="L868"></a>}
                <a id="L869"></a>ret = l % r;
                <a id="L870"></a>return uint64(uint8(ret));
            <a id="L871"></a>}
        <a id="L872"></a>case 16:
            <a id="L873"></a>a.eval = func(t *Thread) uint64 {
                <a id="L874"></a>l, r := lf(t), rf(t);
                <a id="L875"></a>var ret uint64;
                <a id="L876"></a>if r == 0 {
                    <a id="L877"></a>t.Abort(DivByZeroError{})
                <a id="L878"></a>}
                <a id="L879"></a>ret = l % r;
                <a id="L880"></a>return uint64(uint16(ret));
            <a id="L881"></a>}
        <a id="L882"></a>case 32:
            <a id="L883"></a>a.eval = func(t *Thread) uint64 {
                <a id="L884"></a>l, r := lf(t), rf(t);
                <a id="L885"></a>var ret uint64;
                <a id="L886"></a>if r == 0 {
                    <a id="L887"></a>t.Abort(DivByZeroError{})
                <a id="L888"></a>}
                <a id="L889"></a>ret = l % r;
                <a id="L890"></a>return uint64(uint32(ret));
            <a id="L891"></a>}
        <a id="L892"></a>case 64:
            <a id="L893"></a>a.eval = func(t *Thread) uint64 {
                <a id="L894"></a>l, r := lf(t), rf(t);
                <a id="L895"></a>var ret uint64;
                <a id="L896"></a>if r == 0 {
                    <a id="L897"></a>t.Abort(DivByZeroError{})
                <a id="L898"></a>}
                <a id="L899"></a>ret = l % r;
                <a id="L900"></a>return uint64(uint64(ret));
            <a id="L901"></a>}
        <a id="L902"></a>case 0:
            <a id="L903"></a>a.eval = func(t *Thread) uint64 {
                <a id="L904"></a>l, r := lf(t), rf(t);
                <a id="L905"></a>var ret uint64;
                <a id="L906"></a>if r == 0 {
                    <a id="L907"></a>t.Abort(DivByZeroError{})
                <a id="L908"></a>}
                <a id="L909"></a>ret = l % r;
                <a id="L910"></a>return uint64(uint(ret));
            <a id="L911"></a>}
        <a id="L912"></a>default:
            <a id="L913"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L914"></a>}
    <a id="L915"></a>case *intType:
        <a id="L916"></a>lf := l.asInt();
        <a id="L917"></a>rf := r.asInt();
        <a id="L918"></a>switch t.Bits {
        <a id="L919"></a>case 8:
            <a id="L920"></a>a.eval = func(t *Thread) int64 {
                <a id="L921"></a>l, r := lf(t), rf(t);
                <a id="L922"></a>var ret int64;
                <a id="L923"></a>if r == 0 {
                    <a id="L924"></a>t.Abort(DivByZeroError{})
                <a id="L925"></a>}
                <a id="L926"></a>ret = l % r;
                <a id="L927"></a>return int64(int8(ret));
            <a id="L928"></a>}
        <a id="L929"></a>case 16:
            <a id="L930"></a>a.eval = func(t *Thread) int64 {
                <a id="L931"></a>l, r := lf(t), rf(t);
                <a id="L932"></a>var ret int64;
                <a id="L933"></a>if r == 0 {
                    <a id="L934"></a>t.Abort(DivByZeroError{})
                <a id="L935"></a>}
                <a id="L936"></a>ret = l % r;
                <a id="L937"></a>return int64(int16(ret));
            <a id="L938"></a>}
        <a id="L939"></a>case 32:
            <a id="L940"></a>a.eval = func(t *Thread) int64 {
                <a id="L941"></a>l, r := lf(t), rf(t);
                <a id="L942"></a>var ret int64;
                <a id="L943"></a>if r == 0 {
                    <a id="L944"></a>t.Abort(DivByZeroError{})
                <a id="L945"></a>}
                <a id="L946"></a>ret = l % r;
                <a id="L947"></a>return int64(int32(ret));
            <a id="L948"></a>}
        <a id="L949"></a>case 64:
            <a id="L950"></a>a.eval = func(t *Thread) int64 {
                <a id="L951"></a>l, r := lf(t), rf(t);
                <a id="L952"></a>var ret int64;
                <a id="L953"></a>if r == 0 {
                    <a id="L954"></a>t.Abort(DivByZeroError{})
                <a id="L955"></a>}
                <a id="L956"></a>ret = l % r;
                <a id="L957"></a>return int64(int64(ret));
            <a id="L958"></a>}
        <a id="L959"></a>case 0:
            <a id="L960"></a>a.eval = func(t *Thread) int64 {
                <a id="L961"></a>l, r := lf(t), rf(t);
                <a id="L962"></a>var ret int64;
                <a id="L963"></a>if r == 0 {
                    <a id="L964"></a>t.Abort(DivByZeroError{})
                <a id="L965"></a>}
                <a id="L966"></a>ret = l % r;
                <a id="L967"></a>return int64(int(ret));
            <a id="L968"></a>}
        <a id="L969"></a>default:
            <a id="L970"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L971"></a>}
    <a id="L972"></a>case *idealIntType:
        <a id="L973"></a>l := l.asIdealInt()();
        <a id="L974"></a>r := r.asIdealInt()();
        <a id="L975"></a>val := l.Rem(r);
        <a id="L976"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L977"></a>default:
        <a id="L978"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L979"></a>}
<a id="L980"></a>}

<a id="L982"></a>func (a *expr) genBinOpAnd(l, r *expr) {
    <a id="L983"></a>switch t := l.t.lit().(type) {
    <a id="L984"></a>case *uintType:
        <a id="L985"></a>lf := l.asUint();
        <a id="L986"></a>rf := r.asUint();
        <a id="L987"></a>switch t.Bits {
        <a id="L988"></a>case 8:
            <a id="L989"></a>a.eval = func(t *Thread) uint64 {
                <a id="L990"></a>l, r := lf(t), rf(t);
                <a id="L991"></a>var ret uint64;
                <a id="L992"></a>ret = l &amp; r;
                <a id="L993"></a>return uint64(uint8(ret));
            <a id="L994"></a>}
        <a id="L995"></a>case 16:
            <a id="L996"></a>a.eval = func(t *Thread) uint64 {
                <a id="L997"></a>l, r := lf(t), rf(t);
                <a id="L998"></a>var ret uint64;
                <a id="L999"></a>ret = l &amp; r;
                <a id="L1000"></a>return uint64(uint16(ret));
            <a id="L1001"></a>}
        <a id="L1002"></a>case 32:
            <a id="L1003"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1004"></a>l, r := lf(t), rf(t);
                <a id="L1005"></a>var ret uint64;
                <a id="L1006"></a>ret = l &amp; r;
                <a id="L1007"></a>return uint64(uint32(ret));
            <a id="L1008"></a>}
        <a id="L1009"></a>case 64:
            <a id="L1010"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1011"></a>l, r := lf(t), rf(t);
                <a id="L1012"></a>var ret uint64;
                <a id="L1013"></a>ret = l &amp; r;
                <a id="L1014"></a>return uint64(uint64(ret));
            <a id="L1015"></a>}
        <a id="L1016"></a>case 0:
            <a id="L1017"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1018"></a>l, r := lf(t), rf(t);
                <a id="L1019"></a>var ret uint64;
                <a id="L1020"></a>ret = l &amp; r;
                <a id="L1021"></a>return uint64(uint(ret));
            <a id="L1022"></a>}
        <a id="L1023"></a>default:
            <a id="L1024"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1025"></a>}
    <a id="L1026"></a>case *intType:
        <a id="L1027"></a>lf := l.asInt();
        <a id="L1028"></a>rf := r.asInt();
        <a id="L1029"></a>switch t.Bits {
        <a id="L1030"></a>case 8:
            <a id="L1031"></a>a.eval = func(t *Thread) int64 {
                <a id="L1032"></a>l, r := lf(t), rf(t);
                <a id="L1033"></a>var ret int64;
                <a id="L1034"></a>ret = l &amp; r;
                <a id="L1035"></a>return int64(int8(ret));
            <a id="L1036"></a>}
        <a id="L1037"></a>case 16:
            <a id="L1038"></a>a.eval = func(t *Thread) int64 {
                <a id="L1039"></a>l, r := lf(t), rf(t);
                <a id="L1040"></a>var ret int64;
                <a id="L1041"></a>ret = l &amp; r;
                <a id="L1042"></a>return int64(int16(ret));
            <a id="L1043"></a>}
        <a id="L1044"></a>case 32:
            <a id="L1045"></a>a.eval = func(t *Thread) int64 {
                <a id="L1046"></a>l, r := lf(t), rf(t);
                <a id="L1047"></a>var ret int64;
                <a id="L1048"></a>ret = l &amp; r;
                <a id="L1049"></a>return int64(int32(ret));
            <a id="L1050"></a>}
        <a id="L1051"></a>case 64:
            <a id="L1052"></a>a.eval = func(t *Thread) int64 {
                <a id="L1053"></a>l, r := lf(t), rf(t);
                <a id="L1054"></a>var ret int64;
                <a id="L1055"></a>ret = l &amp; r;
                <a id="L1056"></a>return int64(int64(ret));
            <a id="L1057"></a>}
        <a id="L1058"></a>case 0:
            <a id="L1059"></a>a.eval = func(t *Thread) int64 {
                <a id="L1060"></a>l, r := lf(t), rf(t);
                <a id="L1061"></a>var ret int64;
                <a id="L1062"></a>ret = l &amp; r;
                <a id="L1063"></a>return int64(int(ret));
            <a id="L1064"></a>}
        <a id="L1065"></a>default:
            <a id="L1066"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1067"></a>}
    <a id="L1068"></a>case *idealIntType:
        <a id="L1069"></a>l := l.asIdealInt()();
        <a id="L1070"></a>r := r.asIdealInt()();
        <a id="L1071"></a>val := l.And(r);
        <a id="L1072"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L1073"></a>default:
        <a id="L1074"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1075"></a>}
<a id="L1076"></a>}

<a id="L1078"></a>func (a *expr) genBinOpOr(l, r *expr) {
    <a id="L1079"></a>switch t := l.t.lit().(type) {
    <a id="L1080"></a>case *uintType:
        <a id="L1081"></a>lf := l.asUint();
        <a id="L1082"></a>rf := r.asUint();
        <a id="L1083"></a>switch t.Bits {
        <a id="L1084"></a>case 8:
            <a id="L1085"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1086"></a>l, r := lf(t), rf(t);
                <a id="L1087"></a>var ret uint64;
                <a id="L1088"></a>ret = l | r;
                <a id="L1089"></a>return uint64(uint8(ret));
            <a id="L1090"></a>}
        <a id="L1091"></a>case 16:
            <a id="L1092"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1093"></a>l, r := lf(t), rf(t);
                <a id="L1094"></a>var ret uint64;
                <a id="L1095"></a>ret = l | r;
                <a id="L1096"></a>return uint64(uint16(ret));
            <a id="L1097"></a>}
        <a id="L1098"></a>case 32:
            <a id="L1099"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1100"></a>l, r := lf(t), rf(t);
                <a id="L1101"></a>var ret uint64;
                <a id="L1102"></a>ret = l | r;
                <a id="L1103"></a>return uint64(uint32(ret));
            <a id="L1104"></a>}
        <a id="L1105"></a>case 64:
            <a id="L1106"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1107"></a>l, r := lf(t), rf(t);
                <a id="L1108"></a>var ret uint64;
                <a id="L1109"></a>ret = l | r;
                <a id="L1110"></a>return uint64(uint64(ret));
            <a id="L1111"></a>}
        <a id="L1112"></a>case 0:
            <a id="L1113"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1114"></a>l, r := lf(t), rf(t);
                <a id="L1115"></a>var ret uint64;
                <a id="L1116"></a>ret = l | r;
                <a id="L1117"></a>return uint64(uint(ret));
            <a id="L1118"></a>}
        <a id="L1119"></a>default:
            <a id="L1120"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1121"></a>}
    <a id="L1122"></a>case *intType:
        <a id="L1123"></a>lf := l.asInt();
        <a id="L1124"></a>rf := r.asInt();
        <a id="L1125"></a>switch t.Bits {
        <a id="L1126"></a>case 8:
            <a id="L1127"></a>a.eval = func(t *Thread) int64 {
                <a id="L1128"></a>l, r := lf(t), rf(t);
                <a id="L1129"></a>var ret int64;
                <a id="L1130"></a>ret = l | r;
                <a id="L1131"></a>return int64(int8(ret));
            <a id="L1132"></a>}
        <a id="L1133"></a>case 16:
            <a id="L1134"></a>a.eval = func(t *Thread) int64 {
                <a id="L1135"></a>l, r := lf(t), rf(t);
                <a id="L1136"></a>var ret int64;
                <a id="L1137"></a>ret = l | r;
                <a id="L1138"></a>return int64(int16(ret));
            <a id="L1139"></a>}
        <a id="L1140"></a>case 32:
            <a id="L1141"></a>a.eval = func(t *Thread) int64 {
                <a id="L1142"></a>l, r := lf(t), rf(t);
                <a id="L1143"></a>var ret int64;
                <a id="L1144"></a>ret = l | r;
                <a id="L1145"></a>return int64(int32(ret));
            <a id="L1146"></a>}
        <a id="L1147"></a>case 64:
            <a id="L1148"></a>a.eval = func(t *Thread) int64 {
                <a id="L1149"></a>l, r := lf(t), rf(t);
                <a id="L1150"></a>var ret int64;
                <a id="L1151"></a>ret = l | r;
                <a id="L1152"></a>return int64(int64(ret));
            <a id="L1153"></a>}
        <a id="L1154"></a>case 0:
            <a id="L1155"></a>a.eval = func(t *Thread) int64 {
                <a id="L1156"></a>l, r := lf(t), rf(t);
                <a id="L1157"></a>var ret int64;
                <a id="L1158"></a>ret = l | r;
                <a id="L1159"></a>return int64(int(ret));
            <a id="L1160"></a>}
        <a id="L1161"></a>default:
            <a id="L1162"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1163"></a>}
    <a id="L1164"></a>case *idealIntType:
        <a id="L1165"></a>l := l.asIdealInt()();
        <a id="L1166"></a>r := r.asIdealInt()();
        <a id="L1167"></a>val := l.Or(r);
        <a id="L1168"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L1169"></a>default:
        <a id="L1170"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1171"></a>}
<a id="L1172"></a>}

<a id="L1174"></a>func (a *expr) genBinOpXor(l, r *expr) {
    <a id="L1175"></a>switch t := l.t.lit().(type) {
    <a id="L1176"></a>case *uintType:
        <a id="L1177"></a>lf := l.asUint();
        <a id="L1178"></a>rf := r.asUint();
        <a id="L1179"></a>switch t.Bits {
        <a id="L1180"></a>case 8:
            <a id="L1181"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1182"></a>l, r := lf(t), rf(t);
                <a id="L1183"></a>var ret uint64;
                <a id="L1184"></a>ret = l ^ r;
                <a id="L1185"></a>return uint64(uint8(ret));
            <a id="L1186"></a>}
        <a id="L1187"></a>case 16:
            <a id="L1188"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1189"></a>l, r := lf(t), rf(t);
                <a id="L1190"></a>var ret uint64;
                <a id="L1191"></a>ret = l ^ r;
                <a id="L1192"></a>return uint64(uint16(ret));
            <a id="L1193"></a>}
        <a id="L1194"></a>case 32:
            <a id="L1195"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1196"></a>l, r := lf(t), rf(t);
                <a id="L1197"></a>var ret uint64;
                <a id="L1198"></a>ret = l ^ r;
                <a id="L1199"></a>return uint64(uint32(ret));
            <a id="L1200"></a>}
        <a id="L1201"></a>case 64:
            <a id="L1202"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1203"></a>l, r := lf(t), rf(t);
                <a id="L1204"></a>var ret uint64;
                <a id="L1205"></a>ret = l ^ r;
                <a id="L1206"></a>return uint64(uint64(ret));
            <a id="L1207"></a>}
        <a id="L1208"></a>case 0:
            <a id="L1209"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1210"></a>l, r := lf(t), rf(t);
                <a id="L1211"></a>var ret uint64;
                <a id="L1212"></a>ret = l ^ r;
                <a id="L1213"></a>return uint64(uint(ret));
            <a id="L1214"></a>}
        <a id="L1215"></a>default:
            <a id="L1216"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1217"></a>}
    <a id="L1218"></a>case *intType:
        <a id="L1219"></a>lf := l.asInt();
        <a id="L1220"></a>rf := r.asInt();
        <a id="L1221"></a>switch t.Bits {
        <a id="L1222"></a>case 8:
            <a id="L1223"></a>a.eval = func(t *Thread) int64 {
                <a id="L1224"></a>l, r := lf(t), rf(t);
                <a id="L1225"></a>var ret int64;
                <a id="L1226"></a>ret = l ^ r;
                <a id="L1227"></a>return int64(int8(ret));
            <a id="L1228"></a>}
        <a id="L1229"></a>case 16:
            <a id="L1230"></a>a.eval = func(t *Thread) int64 {
                <a id="L1231"></a>l, r := lf(t), rf(t);
                <a id="L1232"></a>var ret int64;
                <a id="L1233"></a>ret = l ^ r;
                <a id="L1234"></a>return int64(int16(ret));
            <a id="L1235"></a>}
        <a id="L1236"></a>case 32:
            <a id="L1237"></a>a.eval = func(t *Thread) int64 {
                <a id="L1238"></a>l, r := lf(t), rf(t);
                <a id="L1239"></a>var ret int64;
                <a id="L1240"></a>ret = l ^ r;
                <a id="L1241"></a>return int64(int32(ret));
            <a id="L1242"></a>}
        <a id="L1243"></a>case 64:
            <a id="L1244"></a>a.eval = func(t *Thread) int64 {
                <a id="L1245"></a>l, r := lf(t), rf(t);
                <a id="L1246"></a>var ret int64;
                <a id="L1247"></a>ret = l ^ r;
                <a id="L1248"></a>return int64(int64(ret));
            <a id="L1249"></a>}
        <a id="L1250"></a>case 0:
            <a id="L1251"></a>a.eval = func(t *Thread) int64 {
                <a id="L1252"></a>l, r := lf(t), rf(t);
                <a id="L1253"></a>var ret int64;
                <a id="L1254"></a>ret = l ^ r;
                <a id="L1255"></a>return int64(int(ret));
            <a id="L1256"></a>}
        <a id="L1257"></a>default:
            <a id="L1258"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1259"></a>}
    <a id="L1260"></a>case *idealIntType:
        <a id="L1261"></a>l := l.asIdealInt()();
        <a id="L1262"></a>r := r.asIdealInt()();
        <a id="L1263"></a>val := l.Xor(r);
        <a id="L1264"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L1265"></a>default:
        <a id="L1266"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1267"></a>}
<a id="L1268"></a>}

<a id="L1270"></a>func (a *expr) genBinOpAndNot(l, r *expr) {
    <a id="L1271"></a>switch t := l.t.lit().(type) {
    <a id="L1272"></a>case *uintType:
        <a id="L1273"></a>lf := l.asUint();
        <a id="L1274"></a>rf := r.asUint();
        <a id="L1275"></a>switch t.Bits {
        <a id="L1276"></a>case 8:
            <a id="L1277"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1278"></a>l, r := lf(t), rf(t);
                <a id="L1279"></a>var ret uint64;
                <a id="L1280"></a>ret = l &amp;^ r;
                <a id="L1281"></a>return uint64(uint8(ret));
            <a id="L1282"></a>}
        <a id="L1283"></a>case 16:
            <a id="L1284"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1285"></a>l, r := lf(t), rf(t);
                <a id="L1286"></a>var ret uint64;
                <a id="L1287"></a>ret = l &amp;^ r;
                <a id="L1288"></a>return uint64(uint16(ret));
            <a id="L1289"></a>}
        <a id="L1290"></a>case 32:
            <a id="L1291"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1292"></a>l, r := lf(t), rf(t);
                <a id="L1293"></a>var ret uint64;
                <a id="L1294"></a>ret = l &amp;^ r;
                <a id="L1295"></a>return uint64(uint32(ret));
            <a id="L1296"></a>}
        <a id="L1297"></a>case 64:
            <a id="L1298"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1299"></a>l, r := lf(t), rf(t);
                <a id="L1300"></a>var ret uint64;
                <a id="L1301"></a>ret = l &amp;^ r;
                <a id="L1302"></a>return uint64(uint64(ret));
            <a id="L1303"></a>}
        <a id="L1304"></a>case 0:
            <a id="L1305"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1306"></a>l, r := lf(t), rf(t);
                <a id="L1307"></a>var ret uint64;
                <a id="L1308"></a>ret = l &amp;^ r;
                <a id="L1309"></a>return uint64(uint(ret));
            <a id="L1310"></a>}
        <a id="L1311"></a>default:
            <a id="L1312"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1313"></a>}
    <a id="L1314"></a>case *intType:
        <a id="L1315"></a>lf := l.asInt();
        <a id="L1316"></a>rf := r.asInt();
        <a id="L1317"></a>switch t.Bits {
        <a id="L1318"></a>case 8:
            <a id="L1319"></a>a.eval = func(t *Thread) int64 {
                <a id="L1320"></a>l, r := lf(t), rf(t);
                <a id="L1321"></a>var ret int64;
                <a id="L1322"></a>ret = l &amp;^ r;
                <a id="L1323"></a>return int64(int8(ret));
            <a id="L1324"></a>}
        <a id="L1325"></a>case 16:
            <a id="L1326"></a>a.eval = func(t *Thread) int64 {
                <a id="L1327"></a>l, r := lf(t), rf(t);
                <a id="L1328"></a>var ret int64;
                <a id="L1329"></a>ret = l &amp;^ r;
                <a id="L1330"></a>return int64(int16(ret));
            <a id="L1331"></a>}
        <a id="L1332"></a>case 32:
            <a id="L1333"></a>a.eval = func(t *Thread) int64 {
                <a id="L1334"></a>l, r := lf(t), rf(t);
                <a id="L1335"></a>var ret int64;
                <a id="L1336"></a>ret = l &amp;^ r;
                <a id="L1337"></a>return int64(int32(ret));
            <a id="L1338"></a>}
        <a id="L1339"></a>case 64:
            <a id="L1340"></a>a.eval = func(t *Thread) int64 {
                <a id="L1341"></a>l, r := lf(t), rf(t);
                <a id="L1342"></a>var ret int64;
                <a id="L1343"></a>ret = l &amp;^ r;
                <a id="L1344"></a>return int64(int64(ret));
            <a id="L1345"></a>}
        <a id="L1346"></a>case 0:
            <a id="L1347"></a>a.eval = func(t *Thread) int64 {
                <a id="L1348"></a>l, r := lf(t), rf(t);
                <a id="L1349"></a>var ret int64;
                <a id="L1350"></a>ret = l &amp;^ r;
                <a id="L1351"></a>return int64(int(ret));
            <a id="L1352"></a>}
        <a id="L1353"></a>default:
            <a id="L1354"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1355"></a>}
    <a id="L1356"></a>case *idealIntType:
        <a id="L1357"></a>l := l.asIdealInt()();
        <a id="L1358"></a>r := r.asIdealInt()();
        <a id="L1359"></a>val := l.AndNot(r);
        <a id="L1360"></a>a.eval = func() *bignum.Integer { return val };
    <a id="L1361"></a>default:
        <a id="L1362"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1363"></a>}
<a id="L1364"></a>}

<a id="L1366"></a>func (a *expr) genBinOpShl(l, r *expr) {
    <a id="L1367"></a>switch t := l.t.lit().(type) {
    <a id="L1368"></a>case *uintType:
        <a id="L1369"></a>lf := l.asUint();
        <a id="L1370"></a>rf := r.asUint();
        <a id="L1371"></a>switch t.Bits {
        <a id="L1372"></a>case 8:
            <a id="L1373"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1374"></a>l, r := lf(t), rf(t);
                <a id="L1375"></a>var ret uint64;
                <a id="L1376"></a>ret = l &lt;&lt; r;
                <a id="L1377"></a>return uint64(uint8(ret));
            <a id="L1378"></a>}
        <a id="L1379"></a>case 16:
            <a id="L1380"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1381"></a>l, r := lf(t), rf(t);
                <a id="L1382"></a>var ret uint64;
                <a id="L1383"></a>ret = l &lt;&lt; r;
                <a id="L1384"></a>return uint64(uint16(ret));
            <a id="L1385"></a>}
        <a id="L1386"></a>case 32:
            <a id="L1387"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1388"></a>l, r := lf(t), rf(t);
                <a id="L1389"></a>var ret uint64;
                <a id="L1390"></a>ret = l &lt;&lt; r;
                <a id="L1391"></a>return uint64(uint32(ret));
            <a id="L1392"></a>}
        <a id="L1393"></a>case 64:
            <a id="L1394"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1395"></a>l, r := lf(t), rf(t);
                <a id="L1396"></a>var ret uint64;
                <a id="L1397"></a>ret = l &lt;&lt; r;
                <a id="L1398"></a>return uint64(uint64(ret));
            <a id="L1399"></a>}
        <a id="L1400"></a>case 0:
            <a id="L1401"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1402"></a>l, r := lf(t), rf(t);
                <a id="L1403"></a>var ret uint64;
                <a id="L1404"></a>ret = l &lt;&lt; r;
                <a id="L1405"></a>return uint64(uint(ret));
            <a id="L1406"></a>}
        <a id="L1407"></a>default:
            <a id="L1408"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1409"></a>}
    <a id="L1410"></a>case *intType:
        <a id="L1411"></a>lf := l.asInt();
        <a id="L1412"></a>rf := r.asUint();
        <a id="L1413"></a>switch t.Bits {
        <a id="L1414"></a>case 8:
            <a id="L1415"></a>a.eval = func(t *Thread) int64 {
                <a id="L1416"></a>l, r := lf(t), rf(t);
                <a id="L1417"></a>var ret int64;
                <a id="L1418"></a>ret = l &lt;&lt; r;
                <a id="L1419"></a>return int64(int8(ret));
            <a id="L1420"></a>}
        <a id="L1421"></a>case 16:
            <a id="L1422"></a>a.eval = func(t *Thread) int64 {
                <a id="L1423"></a>l, r := lf(t), rf(t);
                <a id="L1424"></a>var ret int64;
                <a id="L1425"></a>ret = l &lt;&lt; r;
                <a id="L1426"></a>return int64(int16(ret));
            <a id="L1427"></a>}
        <a id="L1428"></a>case 32:
            <a id="L1429"></a>a.eval = func(t *Thread) int64 {
                <a id="L1430"></a>l, r := lf(t), rf(t);
                <a id="L1431"></a>var ret int64;
                <a id="L1432"></a>ret = l &lt;&lt; r;
                <a id="L1433"></a>return int64(int32(ret));
            <a id="L1434"></a>}
        <a id="L1435"></a>case 64:
            <a id="L1436"></a>a.eval = func(t *Thread) int64 {
                <a id="L1437"></a>l, r := lf(t), rf(t);
                <a id="L1438"></a>var ret int64;
                <a id="L1439"></a>ret = l &lt;&lt; r;
                <a id="L1440"></a>return int64(int64(ret));
            <a id="L1441"></a>}
        <a id="L1442"></a>case 0:
            <a id="L1443"></a>a.eval = func(t *Thread) int64 {
                <a id="L1444"></a>l, r := lf(t), rf(t);
                <a id="L1445"></a>var ret int64;
                <a id="L1446"></a>ret = l &lt;&lt; r;
                <a id="L1447"></a>return int64(int(ret));
            <a id="L1448"></a>}
        <a id="L1449"></a>default:
            <a id="L1450"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1451"></a>}
    <a id="L1452"></a>default:
        <a id="L1453"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1454"></a>}
<a id="L1455"></a>}

<a id="L1457"></a>func (a *expr) genBinOpShr(l, r *expr) {
    <a id="L1458"></a>switch t := l.t.lit().(type) {
    <a id="L1459"></a>case *uintType:
        <a id="L1460"></a>lf := l.asUint();
        <a id="L1461"></a>rf := r.asUint();
        <a id="L1462"></a>switch t.Bits {
        <a id="L1463"></a>case 8:
            <a id="L1464"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1465"></a>l, r := lf(t), rf(t);
                <a id="L1466"></a>var ret uint64;
                <a id="L1467"></a>ret = l &gt;&gt; r;
                <a id="L1468"></a>return uint64(uint8(ret));
            <a id="L1469"></a>}
        <a id="L1470"></a>case 16:
            <a id="L1471"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1472"></a>l, r := lf(t), rf(t);
                <a id="L1473"></a>var ret uint64;
                <a id="L1474"></a>ret = l &gt;&gt; r;
                <a id="L1475"></a>return uint64(uint16(ret));
            <a id="L1476"></a>}
        <a id="L1477"></a>case 32:
            <a id="L1478"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1479"></a>l, r := lf(t), rf(t);
                <a id="L1480"></a>var ret uint64;
                <a id="L1481"></a>ret = l &gt;&gt; r;
                <a id="L1482"></a>return uint64(uint32(ret));
            <a id="L1483"></a>}
        <a id="L1484"></a>case 64:
            <a id="L1485"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1486"></a>l, r := lf(t), rf(t);
                <a id="L1487"></a>var ret uint64;
                <a id="L1488"></a>ret = l &gt;&gt; r;
                <a id="L1489"></a>return uint64(uint64(ret));
            <a id="L1490"></a>}
        <a id="L1491"></a>case 0:
            <a id="L1492"></a>a.eval = func(t *Thread) uint64 {
                <a id="L1493"></a>l, r := lf(t), rf(t);
                <a id="L1494"></a>var ret uint64;
                <a id="L1495"></a>ret = l &gt;&gt; r;
                <a id="L1496"></a>return uint64(uint(ret));
            <a id="L1497"></a>}
        <a id="L1498"></a>default:
            <a id="L1499"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1500"></a>}
    <a id="L1501"></a>case *intType:
        <a id="L1502"></a>lf := l.asInt();
        <a id="L1503"></a>rf := r.asUint();
        <a id="L1504"></a>switch t.Bits {
        <a id="L1505"></a>case 8:
            <a id="L1506"></a>a.eval = func(t *Thread) int64 {
                <a id="L1507"></a>l, r := lf(t), rf(t);
                <a id="L1508"></a>var ret int64;
                <a id="L1509"></a>ret = l &gt;&gt; r;
                <a id="L1510"></a>return int64(int8(ret));
            <a id="L1511"></a>}
        <a id="L1512"></a>case 16:
            <a id="L1513"></a>a.eval = func(t *Thread) int64 {
                <a id="L1514"></a>l, r := lf(t), rf(t);
                <a id="L1515"></a>var ret int64;
                <a id="L1516"></a>ret = l &gt;&gt; r;
                <a id="L1517"></a>return int64(int16(ret));
            <a id="L1518"></a>}
        <a id="L1519"></a>case 32:
            <a id="L1520"></a>a.eval = func(t *Thread) int64 {
                <a id="L1521"></a>l, r := lf(t), rf(t);
                <a id="L1522"></a>var ret int64;
                <a id="L1523"></a>ret = l &gt;&gt; r;
                <a id="L1524"></a>return int64(int32(ret));
            <a id="L1525"></a>}
        <a id="L1526"></a>case 64:
            <a id="L1527"></a>a.eval = func(t *Thread) int64 {
                <a id="L1528"></a>l, r := lf(t), rf(t);
                <a id="L1529"></a>var ret int64;
                <a id="L1530"></a>ret = l &gt;&gt; r;
                <a id="L1531"></a>return int64(int64(ret));
            <a id="L1532"></a>}
        <a id="L1533"></a>case 0:
            <a id="L1534"></a>a.eval = func(t *Thread) int64 {
                <a id="L1535"></a>l, r := lf(t), rf(t);
                <a id="L1536"></a>var ret int64;
                <a id="L1537"></a>ret = l &gt;&gt; r;
                <a id="L1538"></a>return int64(int(ret));
            <a id="L1539"></a>}
        <a id="L1540"></a>default:
            <a id="L1541"></a>log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos)
        <a id="L1542"></a>}
    <a id="L1543"></a>default:
        <a id="L1544"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1545"></a>}
<a id="L1546"></a>}

<a id="L1548"></a>func (a *expr) genBinOpLss(l, r *expr) {
    <a id="L1549"></a>switch t := l.t.lit().(type) {
    <a id="L1550"></a>case *uintType:
        <a id="L1551"></a>lf := l.asUint();
        <a id="L1552"></a>rf := r.asUint();
        <a id="L1553"></a>a.eval = func(t *Thread) bool {
            <a id="L1554"></a>l, r := lf(t), rf(t);
            <a id="L1555"></a>return l &lt; r;
        <a id="L1556"></a>};
    <a id="L1557"></a>case *intType:
        <a id="L1558"></a>lf := l.asInt();
        <a id="L1559"></a>rf := r.asInt();
        <a id="L1560"></a>a.eval = func(t *Thread) bool {
            <a id="L1561"></a>l, r := lf(t), rf(t);
            <a id="L1562"></a>return l &lt; r;
        <a id="L1563"></a>};
    <a id="L1564"></a>case *idealIntType:
        <a id="L1565"></a>l := l.asIdealInt()();
        <a id="L1566"></a>r := r.asIdealInt()();
        <a id="L1567"></a>val := l.Cmp(r) &lt; 0;
        <a id="L1568"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1569"></a>case *floatType:
        <a id="L1570"></a>lf := l.asFloat();
        <a id="L1571"></a>rf := r.asFloat();
        <a id="L1572"></a>a.eval = func(t *Thread) bool {
            <a id="L1573"></a>l, r := lf(t), rf(t);
            <a id="L1574"></a>return l &lt; r;
        <a id="L1575"></a>};
    <a id="L1576"></a>case *idealFloatType:
        <a id="L1577"></a>l := l.asIdealFloat()();
        <a id="L1578"></a>r := r.asIdealFloat()();
        <a id="L1579"></a>val := l.Cmp(r) &lt; 0;
        <a id="L1580"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1581"></a>case *stringType:
        <a id="L1582"></a>lf := l.asString();
        <a id="L1583"></a>rf := r.asString();
        <a id="L1584"></a>a.eval = func(t *Thread) bool {
            <a id="L1585"></a>l, r := lf(t), rf(t);
            <a id="L1586"></a>return l &lt; r;
        <a id="L1587"></a>};
    <a id="L1588"></a>default:
        <a id="L1589"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1590"></a>}
<a id="L1591"></a>}

<a id="L1593"></a>func (a *expr) genBinOpGtr(l, r *expr) {
    <a id="L1594"></a>switch t := l.t.lit().(type) {
    <a id="L1595"></a>case *uintType:
        <a id="L1596"></a>lf := l.asUint();
        <a id="L1597"></a>rf := r.asUint();
        <a id="L1598"></a>a.eval = func(t *Thread) bool {
            <a id="L1599"></a>l, r := lf(t), rf(t);
            <a id="L1600"></a>return l &gt; r;
        <a id="L1601"></a>};
    <a id="L1602"></a>case *intType:
        <a id="L1603"></a>lf := l.asInt();
        <a id="L1604"></a>rf := r.asInt();
        <a id="L1605"></a>a.eval = func(t *Thread) bool {
            <a id="L1606"></a>l, r := lf(t), rf(t);
            <a id="L1607"></a>return l &gt; r;
        <a id="L1608"></a>};
    <a id="L1609"></a>case *idealIntType:
        <a id="L1610"></a>l := l.asIdealInt()();
        <a id="L1611"></a>r := r.asIdealInt()();
        <a id="L1612"></a>val := l.Cmp(r) &gt; 0;
        <a id="L1613"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1614"></a>case *floatType:
        <a id="L1615"></a>lf := l.asFloat();
        <a id="L1616"></a>rf := r.asFloat();
        <a id="L1617"></a>a.eval = func(t *Thread) bool {
            <a id="L1618"></a>l, r := lf(t), rf(t);
            <a id="L1619"></a>return l &gt; r;
        <a id="L1620"></a>};
    <a id="L1621"></a>case *idealFloatType:
        <a id="L1622"></a>l := l.asIdealFloat()();
        <a id="L1623"></a>r := r.asIdealFloat()();
        <a id="L1624"></a>val := l.Cmp(r) &gt; 0;
        <a id="L1625"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1626"></a>case *stringType:
        <a id="L1627"></a>lf := l.asString();
        <a id="L1628"></a>rf := r.asString();
        <a id="L1629"></a>a.eval = func(t *Thread) bool {
            <a id="L1630"></a>l, r := lf(t), rf(t);
            <a id="L1631"></a>return l &gt; r;
        <a id="L1632"></a>};
    <a id="L1633"></a>default:
        <a id="L1634"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1635"></a>}
<a id="L1636"></a>}

<a id="L1638"></a>func (a *expr) genBinOpLeq(l, r *expr) {
    <a id="L1639"></a>switch t := l.t.lit().(type) {
    <a id="L1640"></a>case *uintType:
        <a id="L1641"></a>lf := l.asUint();
        <a id="L1642"></a>rf := r.asUint();
        <a id="L1643"></a>a.eval = func(t *Thread) bool {
            <a id="L1644"></a>l, r := lf(t), rf(t);
            <a id="L1645"></a>return l &lt;= r;
        <a id="L1646"></a>};
    <a id="L1647"></a>case *intType:
        <a id="L1648"></a>lf := l.asInt();
        <a id="L1649"></a>rf := r.asInt();
        <a id="L1650"></a>a.eval = func(t *Thread) bool {
            <a id="L1651"></a>l, r := lf(t), rf(t);
            <a id="L1652"></a>return l &lt;= r;
        <a id="L1653"></a>};
    <a id="L1654"></a>case *idealIntType:
        <a id="L1655"></a>l := l.asIdealInt()();
        <a id="L1656"></a>r := r.asIdealInt()();
        <a id="L1657"></a>val := l.Cmp(r) &lt;= 0;
        <a id="L1658"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1659"></a>case *floatType:
        <a id="L1660"></a>lf := l.asFloat();
        <a id="L1661"></a>rf := r.asFloat();
        <a id="L1662"></a>a.eval = func(t *Thread) bool {
            <a id="L1663"></a>l, r := lf(t), rf(t);
            <a id="L1664"></a>return l &lt;= r;
        <a id="L1665"></a>};
    <a id="L1666"></a>case *idealFloatType:
        <a id="L1667"></a>l := l.asIdealFloat()();
        <a id="L1668"></a>r := r.asIdealFloat()();
        <a id="L1669"></a>val := l.Cmp(r) &lt;= 0;
        <a id="L1670"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1671"></a>case *stringType:
        <a id="L1672"></a>lf := l.asString();
        <a id="L1673"></a>rf := r.asString();
        <a id="L1674"></a>a.eval = func(t *Thread) bool {
            <a id="L1675"></a>l, r := lf(t), rf(t);
            <a id="L1676"></a>return l &lt;= r;
        <a id="L1677"></a>};
    <a id="L1678"></a>default:
        <a id="L1679"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1680"></a>}
<a id="L1681"></a>}

<a id="L1683"></a>func (a *expr) genBinOpGeq(l, r *expr) {
    <a id="L1684"></a>switch t := l.t.lit().(type) {
    <a id="L1685"></a>case *uintType:
        <a id="L1686"></a>lf := l.asUint();
        <a id="L1687"></a>rf := r.asUint();
        <a id="L1688"></a>a.eval = func(t *Thread) bool {
            <a id="L1689"></a>l, r := lf(t), rf(t);
            <a id="L1690"></a>return l &gt;= r;
        <a id="L1691"></a>};
    <a id="L1692"></a>case *intType:
        <a id="L1693"></a>lf := l.asInt();
        <a id="L1694"></a>rf := r.asInt();
        <a id="L1695"></a>a.eval = func(t *Thread) bool {
            <a id="L1696"></a>l, r := lf(t), rf(t);
            <a id="L1697"></a>return l &gt;= r;
        <a id="L1698"></a>};
    <a id="L1699"></a>case *idealIntType:
        <a id="L1700"></a>l := l.asIdealInt()();
        <a id="L1701"></a>r := r.asIdealInt()();
        <a id="L1702"></a>val := l.Cmp(r) &gt;= 0;
        <a id="L1703"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1704"></a>case *floatType:
        <a id="L1705"></a>lf := l.asFloat();
        <a id="L1706"></a>rf := r.asFloat();
        <a id="L1707"></a>a.eval = func(t *Thread) bool {
            <a id="L1708"></a>l, r := lf(t), rf(t);
            <a id="L1709"></a>return l &gt;= r;
        <a id="L1710"></a>};
    <a id="L1711"></a>case *idealFloatType:
        <a id="L1712"></a>l := l.asIdealFloat()();
        <a id="L1713"></a>r := r.asIdealFloat()();
        <a id="L1714"></a>val := l.Cmp(r) &gt;= 0;
        <a id="L1715"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1716"></a>case *stringType:
        <a id="L1717"></a>lf := l.asString();
        <a id="L1718"></a>rf := r.asString();
        <a id="L1719"></a>a.eval = func(t *Thread) bool {
            <a id="L1720"></a>l, r := lf(t), rf(t);
            <a id="L1721"></a>return l &gt;= r;
        <a id="L1722"></a>};
    <a id="L1723"></a>default:
        <a id="L1724"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1725"></a>}
<a id="L1726"></a>}

<a id="L1728"></a>func (a *expr) genBinOpEql(l, r *expr) {
    <a id="L1729"></a>switch t := l.t.lit().(type) {
    <a id="L1730"></a>case *boolType:
        <a id="L1731"></a>lf := l.asBool();
        <a id="L1732"></a>rf := r.asBool();
        <a id="L1733"></a>a.eval = func(t *Thread) bool {
            <a id="L1734"></a>l, r := lf(t), rf(t);
            <a id="L1735"></a>return l == r;
        <a id="L1736"></a>};
    <a id="L1737"></a>case *uintType:
        <a id="L1738"></a>lf := l.asUint();
        <a id="L1739"></a>rf := r.asUint();
        <a id="L1740"></a>a.eval = func(t *Thread) bool {
            <a id="L1741"></a>l, r := lf(t), rf(t);
            <a id="L1742"></a>return l == r;
        <a id="L1743"></a>};
    <a id="L1744"></a>case *intType:
        <a id="L1745"></a>lf := l.asInt();
        <a id="L1746"></a>rf := r.asInt();
        <a id="L1747"></a>a.eval = func(t *Thread) bool {
            <a id="L1748"></a>l, r := lf(t), rf(t);
            <a id="L1749"></a>return l == r;
        <a id="L1750"></a>};
    <a id="L1751"></a>case *idealIntType:
        <a id="L1752"></a>l := l.asIdealInt()();
        <a id="L1753"></a>r := r.asIdealInt()();
        <a id="L1754"></a>val := l.Cmp(r) == 0;
        <a id="L1755"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1756"></a>case *floatType:
        <a id="L1757"></a>lf := l.asFloat();
        <a id="L1758"></a>rf := r.asFloat();
        <a id="L1759"></a>a.eval = func(t *Thread) bool {
            <a id="L1760"></a>l, r := lf(t), rf(t);
            <a id="L1761"></a>return l == r;
        <a id="L1762"></a>};
    <a id="L1763"></a>case *idealFloatType:
        <a id="L1764"></a>l := l.asIdealFloat()();
        <a id="L1765"></a>r := r.asIdealFloat()();
        <a id="L1766"></a>val := l.Cmp(r) == 0;
        <a id="L1767"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1768"></a>case *stringType:
        <a id="L1769"></a>lf := l.asString();
        <a id="L1770"></a>rf := r.asString();
        <a id="L1771"></a>a.eval = func(t *Thread) bool {
            <a id="L1772"></a>l, r := lf(t), rf(t);
            <a id="L1773"></a>return l == r;
        <a id="L1774"></a>};
    <a id="L1775"></a>case *PtrType:
        <a id="L1776"></a>lf := l.asPtr();
        <a id="L1777"></a>rf := r.asPtr();
        <a id="L1778"></a>a.eval = func(t *Thread) bool {
            <a id="L1779"></a>l, r := lf(t), rf(t);
            <a id="L1780"></a>return l == r;
        <a id="L1781"></a>};
    <a id="L1782"></a>case *FuncType:
        <a id="L1783"></a>lf := l.asFunc();
        <a id="L1784"></a>rf := r.asFunc();
        <a id="L1785"></a>a.eval = func(t *Thread) bool {
            <a id="L1786"></a>l, r := lf(t), rf(t);
            <a id="L1787"></a>return l == r;
        <a id="L1788"></a>};
    <a id="L1789"></a>case *MapType:
        <a id="L1790"></a>lf := l.asMap();
        <a id="L1791"></a>rf := r.asMap();
        <a id="L1792"></a>a.eval = func(t *Thread) bool {
            <a id="L1793"></a>l, r := lf(t), rf(t);
            <a id="L1794"></a>return l == r;
        <a id="L1795"></a>};
    <a id="L1796"></a>default:
        <a id="L1797"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1798"></a>}
<a id="L1799"></a>}

<a id="L1801"></a>func (a *expr) genBinOpNeq(l, r *expr) {
    <a id="L1802"></a>switch t := l.t.lit().(type) {
    <a id="L1803"></a>case *boolType:
        <a id="L1804"></a>lf := l.asBool();
        <a id="L1805"></a>rf := r.asBool();
        <a id="L1806"></a>a.eval = func(t *Thread) bool {
            <a id="L1807"></a>l, r := lf(t), rf(t);
            <a id="L1808"></a>return l != r;
        <a id="L1809"></a>};
    <a id="L1810"></a>case *uintType:
        <a id="L1811"></a>lf := l.asUint();
        <a id="L1812"></a>rf := r.asUint();
        <a id="L1813"></a>a.eval = func(t *Thread) bool {
            <a id="L1814"></a>l, r := lf(t), rf(t);
            <a id="L1815"></a>return l != r;
        <a id="L1816"></a>};
    <a id="L1817"></a>case *intType:
        <a id="L1818"></a>lf := l.asInt();
        <a id="L1819"></a>rf := r.asInt();
        <a id="L1820"></a>a.eval = func(t *Thread) bool {
            <a id="L1821"></a>l, r := lf(t), rf(t);
            <a id="L1822"></a>return l != r;
        <a id="L1823"></a>};
    <a id="L1824"></a>case *idealIntType:
        <a id="L1825"></a>l := l.asIdealInt()();
        <a id="L1826"></a>r := r.asIdealInt()();
        <a id="L1827"></a>val := l.Cmp(r) != 0;
        <a id="L1828"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1829"></a>case *floatType:
        <a id="L1830"></a>lf := l.asFloat();
        <a id="L1831"></a>rf := r.asFloat();
        <a id="L1832"></a>a.eval = func(t *Thread) bool {
            <a id="L1833"></a>l, r := lf(t), rf(t);
            <a id="L1834"></a>return l != r;
        <a id="L1835"></a>};
    <a id="L1836"></a>case *idealFloatType:
        <a id="L1837"></a>l := l.asIdealFloat()();
        <a id="L1838"></a>r := r.asIdealFloat()();
        <a id="L1839"></a>val := l.Cmp(r) != 0;
        <a id="L1840"></a>a.eval = func(t *Thread) bool { return val };
    <a id="L1841"></a>case *stringType:
        <a id="L1842"></a>lf := l.asString();
        <a id="L1843"></a>rf := r.asString();
        <a id="L1844"></a>a.eval = func(t *Thread) bool {
            <a id="L1845"></a>l, r := lf(t), rf(t);
            <a id="L1846"></a>return l != r;
        <a id="L1847"></a>};
    <a id="L1848"></a>case *PtrType:
        <a id="L1849"></a>lf := l.asPtr();
        <a id="L1850"></a>rf := r.asPtr();
        <a id="L1851"></a>a.eval = func(t *Thread) bool {
            <a id="L1852"></a>l, r := lf(t), rf(t);
            <a id="L1853"></a>return l != r;
        <a id="L1854"></a>};
    <a id="L1855"></a>case *FuncType:
        <a id="L1856"></a>lf := l.asFunc();
        <a id="L1857"></a>rf := r.asFunc();
        <a id="L1858"></a>a.eval = func(t *Thread) bool {
            <a id="L1859"></a>l, r := lf(t), rf(t);
            <a id="L1860"></a>return l != r;
        <a id="L1861"></a>};
    <a id="L1862"></a>case *MapType:
        <a id="L1863"></a>lf := l.asMap();
        <a id="L1864"></a>rf := r.asMap();
        <a id="L1865"></a>a.eval = func(t *Thread) bool {
            <a id="L1866"></a>l, r := lf(t), rf(t);
            <a id="L1867"></a>return l != r;
        <a id="L1868"></a>};
    <a id="L1869"></a>default:
        <a id="L1870"></a>log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos)
    <a id="L1871"></a>}
<a id="L1872"></a>}

<a id="L1874"></a>func genAssign(lt Type, r *expr) (func(lv Value, t *Thread)) {
    <a id="L1875"></a>switch lt.lit().(type) {
    <a id="L1876"></a>case *boolType:
        <a id="L1877"></a>rf := r.asBool();
        <a id="L1878"></a>return func(lv Value, t *Thread) { lv.(BoolValue).Set(t, rf(t)) };
    <a id="L1879"></a>case *uintType:
        <a id="L1880"></a>rf := r.asUint();
        <a id="L1881"></a>return func(lv Value, t *Thread) { lv.(UintValue).Set(t, rf(t)) };
    <a id="L1882"></a>case *intType:
        <a id="L1883"></a>rf := r.asInt();
        <a id="L1884"></a>return func(lv Value, t *Thread) { lv.(IntValue).Set(t, rf(t)) };
    <a id="L1885"></a>case *floatType:
        <a id="L1886"></a>rf := r.asFloat();
        <a id="L1887"></a>return func(lv Value, t *Thread) { lv.(FloatValue).Set(t, rf(t)) };
    <a id="L1888"></a>case *stringType:
        <a id="L1889"></a>rf := r.asString();
        <a id="L1890"></a>return func(lv Value, t *Thread) { lv.(StringValue).Set(t, rf(t)) };
    <a id="L1891"></a>case *ArrayType:
        <a id="L1892"></a>rf := r.asArray();
        <a id="L1893"></a>return func(lv Value, t *Thread) { lv.Assign(t, rf(t)) };
    <a id="L1894"></a>case *StructType:
        <a id="L1895"></a>rf := r.asStruct();
        <a id="L1896"></a>return func(lv Value, t *Thread) { lv.Assign(t, rf(t)) };
    <a id="L1897"></a>case *PtrType:
        <a id="L1898"></a>rf := r.asPtr();
        <a id="L1899"></a>return func(lv Value, t *Thread) { lv.(PtrValue).Set(t, rf(t)) };
    <a id="L1900"></a>case *FuncType:
        <a id="L1901"></a>rf := r.asFunc();
        <a id="L1902"></a>return func(lv Value, t *Thread) { lv.(FuncValue).Set(t, rf(t)) };
    <a id="L1903"></a>case *SliceType:
        <a id="L1904"></a>rf := r.asSlice();
        <a id="L1905"></a>return func(lv Value, t *Thread) { lv.(SliceValue).Set(t, rf(t)) };
    <a id="L1906"></a>case *MapType:
        <a id="L1907"></a>rf := r.asMap();
        <a id="L1908"></a>return func(lv Value, t *Thread) { lv.(MapValue).Set(t, rf(t)) };
    <a id="L1909"></a>default:
        <a id="L1910"></a>log.Crashf(&#34;unexpected left operand type %v at %v&#34;, lt, r.pos)
    <a id="L1911"></a>}
    <a id="L1912"></a>panic();
<a id="L1913"></a>}
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
