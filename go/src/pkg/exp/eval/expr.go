<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/expr.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/expr.go</h1>

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
    <a id="L12"></a>&#34;strconv&#34;;
    <a id="L13"></a>&#34;strings&#34;;
    <a id="L14"></a>&#34;os&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// An expr is the result of compiling an expression.  It stores the</span>
<a id="L18"></a><span class="comment">// type of the expression and its evaluator function.</span>
<a id="L19"></a>type expr struct {
    <a id="L20"></a>*exprInfo;
    <a id="L21"></a>t   Type;

    <a id="L23"></a><span class="comment">// Evaluate this node as the given type.</span>
    <a id="L24"></a>eval interface{};

    <a id="L26"></a><span class="comment">// Map index expressions permit special forms of assignment,</span>
    <a id="L27"></a><span class="comment">// for which we need to know the Map and key.</span>
    <a id="L28"></a>evalMapValue func(t *Thread) (Map, interface{});

    <a id="L30"></a><span class="comment">// Evaluate to the &#34;address of&#34; this value; that is, the</span>
    <a id="L31"></a><span class="comment">// settable Value object.  nil for expressions whose address</span>
    <a id="L32"></a><span class="comment">// cannot be taken.</span>
    <a id="L33"></a>evalAddr func(t *Thread) Value;

    <a id="L35"></a><span class="comment">// Execute this expression as a statement.  Only expressions</span>
    <a id="L36"></a><span class="comment">// that are valid expression statements should set this.</span>
    <a id="L37"></a>exec func(t *Thread);

    <a id="L39"></a><span class="comment">// If this expression is a type, this is its compiled type.</span>
    <a id="L40"></a><span class="comment">// This is only permitted in the function position of a call</span>
    <a id="L41"></a><span class="comment">// expression.  In this case, t should be nil.</span>
    <a id="L42"></a>valType Type;

    <a id="L44"></a><span class="comment">// A short string describing this expression for error</span>
    <a id="L45"></a><span class="comment">// messages.</span>
    <a id="L46"></a>desc string;
<a id="L47"></a>}

<a id="L49"></a><span class="comment">// exprInfo stores information needed to compile any expression node.</span>
<a id="L50"></a><span class="comment">// Each expr also stores its exprInfo so further expressions can be</span>
<a id="L51"></a><span class="comment">// compiled from it.</span>
<a id="L52"></a>type exprInfo struct {
    <a id="L53"></a>*compiler;
    <a id="L54"></a>pos token.Position;
<a id="L55"></a>}

<a id="L57"></a>func (a *exprInfo) newExpr(t Type, desc string) *expr {
    <a id="L58"></a>return &amp;expr{exprInfo: a, t: t, desc: desc}
<a id="L59"></a>}

<a id="L61"></a>func (a *exprInfo) diag(format string, args ...) {
    <a id="L62"></a>a.diagAt(&amp;a.pos, format, args)
<a id="L63"></a>}

<a id="L65"></a>func (a *exprInfo) diagOpType(op token.Token, vt Type) {
    <a id="L66"></a>a.diag(&#34;illegal operand type for &#39;%v&#39; operator\n\t%v&#34;, op, vt)
<a id="L67"></a>}

<a id="L69"></a>func (a *exprInfo) diagOpTypes(op token.Token, lt Type, rt Type) {
    <a id="L70"></a>a.diag(&#34;illegal operand types for &#39;%v&#39; operator\n\t%v\n\t%v&#34;, op, lt, rt)
<a id="L71"></a>}

<a id="L73"></a><span class="comment">/*</span>
<a id="L74"></a><span class="comment"> * Common expression manipulations</span>
<a id="L75"></a><span class="comment"> */</span>

<a id="L77"></a><span class="comment">// a.convertTo(t) converts the value of the analyzed expression a,</span>
<a id="L78"></a><span class="comment">// which must be a constant, ideal number, to a new analyzed</span>
<a id="L79"></a><span class="comment">// expression with a constant value of type t.</span>
<a id="L80"></a><span class="comment">//</span>
<a id="L81"></a><span class="comment">// TODO(austin) Rename to resolveIdeal or something?</span>
<a id="L82"></a>func (a *expr) convertTo(t Type) *expr {
    <a id="L83"></a>if !a.t.isIdeal() {
        <a id="L84"></a>log.Crashf(&#34;attempted to convert from %v, expected ideal&#34;, a.t)
    <a id="L85"></a>}

    <a id="L87"></a>var rat *bignum.Rational;

    <a id="L89"></a><span class="comment">// XXX(Spec)  The spec says &#34;It is erroneous&#34;.</span>
    <a id="L90"></a><span class="comment">//</span>
    <a id="L91"></a><span class="comment">// It is an error to assign a value with a non-zero fractional</span>
    <a id="L92"></a><span class="comment">// part to an integer, or if the assignment would overflow or</span>
    <a id="L93"></a><span class="comment">// underflow, or in general if the value cannot be represented</span>
    <a id="L94"></a><span class="comment">// by the type of the variable.</span>
    <a id="L95"></a>switch a.t {
    <a id="L96"></a>case IdealFloatType:
        <a id="L97"></a>rat = a.asIdealFloat()();
        <a id="L98"></a>if t.isInteger() &amp;&amp; !rat.IsInt() {
            <a id="L99"></a>a.diag(&#34;constant %v truncated to integer&#34;, ratToString(rat));
            <a id="L100"></a>return nil;
        <a id="L101"></a>}
    <a id="L102"></a>case IdealIntType:
        <a id="L103"></a>i := a.asIdealInt()();
        <a id="L104"></a>rat = bignum.MakeRat(i, bignum.Nat(1));
    <a id="L105"></a>default:
        <a id="L106"></a>log.Crashf(&#34;unexpected ideal type %v&#34;, a.t)
    <a id="L107"></a>}

    <a id="L109"></a><span class="comment">// Check bounds</span>
    <a id="L110"></a>if t, ok := t.lit().(BoundedType); ok {
        <a id="L111"></a>if rat.Cmp(t.minVal()) &lt; 0 {
            <a id="L112"></a>a.diag(&#34;constant %v underflows %v&#34;, ratToString(rat), t);
            <a id="L113"></a>return nil;
        <a id="L114"></a>}
        <a id="L115"></a>if rat.Cmp(t.maxVal()) &gt; 0 {
            <a id="L116"></a>a.diag(&#34;constant %v overflows %v&#34;, ratToString(rat), t);
            <a id="L117"></a>return nil;
        <a id="L118"></a>}
    <a id="L119"></a>}

    <a id="L121"></a><span class="comment">// Convert rat to type t.</span>
    <a id="L122"></a>res := a.newExpr(t, a.desc);
    <a id="L123"></a>switch t := t.lit().(type) {
    <a id="L124"></a>case *uintType:
        <a id="L125"></a>n, d := rat.Value();
        <a id="L126"></a>f := n.Quo(bignum.MakeInt(false, d));
        <a id="L127"></a>v := f.Abs().Value();
        <a id="L128"></a>res.eval = func(*Thread) uint64 { return v };
    <a id="L129"></a>case *intType:
        <a id="L130"></a>n, d := rat.Value();
        <a id="L131"></a>f := n.Quo(bignum.MakeInt(false, d));
        <a id="L132"></a>v := f.Value();
        <a id="L133"></a>res.eval = func(*Thread) int64 { return v };
    <a id="L134"></a>case *idealIntType:
        <a id="L135"></a>n, d := rat.Value();
        <a id="L136"></a>f := n.Quo(bignum.MakeInt(false, d));
        <a id="L137"></a>res.eval = func() *bignum.Integer { return f };
    <a id="L138"></a>case *floatType:
        <a id="L139"></a>n, d := rat.Value();
        <a id="L140"></a>v := float64(n.Value()) / float64(d.Value());
        <a id="L141"></a>res.eval = func(*Thread) float64 { return v };
    <a id="L142"></a>case *idealFloatType:
        <a id="L143"></a>res.eval = func() *bignum.Rational { return rat }
    <a id="L144"></a>default:
        <a id="L145"></a>log.Crashf(&#34;cannot convert to type %T&#34;, t)
    <a id="L146"></a>}

    <a id="L148"></a>return res;
<a id="L149"></a>}

<a id="L151"></a><span class="comment">// convertToInt converts this expression to an integer, if possible,</span>
<a id="L152"></a><span class="comment">// or produces an error if not.  This accepts ideal ints, uints, and</span>
<a id="L153"></a><span class="comment">// ints.  If max is not -1, produces an error if possible if the value</span>
<a id="L154"></a><span class="comment">// exceeds max.  If negErr is not &#34;&#34;, produces an error if possible if</span>
<a id="L155"></a><span class="comment">// the value is negative.</span>
<a id="L156"></a>func (a *expr) convertToInt(max int64, negErr string, errOp string) *expr {
    <a id="L157"></a>switch a.t.lit().(type) {
    <a id="L158"></a>case *idealIntType:
        <a id="L159"></a>val := a.asIdealInt()();
        <a id="L160"></a>if negErr != &#34;&#34; &amp;&amp; val.IsNeg() {
            <a id="L161"></a>a.diag(&#34;negative %s: %s&#34;, negErr, val);
            <a id="L162"></a>return nil;
        <a id="L163"></a>}
        <a id="L164"></a>bound := max;
        <a id="L165"></a>if negErr == &#34;slice&#34; {
            <a id="L166"></a>bound++
        <a id="L167"></a>}
        <a id="L168"></a>if max != -1 &amp;&amp; val.Cmp(bignum.Int(bound)) &gt;= 0 {
            <a id="L169"></a>a.diag(&#34;index %s exceeds length %d&#34;, val, max);
            <a id="L170"></a>return nil;
        <a id="L171"></a>}
        <a id="L172"></a>return a.convertTo(IntType);

    <a id="L174"></a>case *uintType:
        <a id="L175"></a><span class="comment">// Convert to int</span>
        <a id="L176"></a>na := a.newExpr(IntType, a.desc);
        <a id="L177"></a>af := a.asUint();
        <a id="L178"></a>na.eval = func(t *Thread) int64 { return int64(af(t)) };
        <a id="L179"></a>return na;

    <a id="L181"></a>case *intType:
        <a id="L182"></a><span class="comment">// Good as is</span>
        <a id="L183"></a>return a
    <a id="L184"></a>}

    <a id="L186"></a>a.diag(&#34;illegal operand type for %s\n\t%v&#34;, errOp, a.t);
    <a id="L187"></a>return nil;
<a id="L188"></a>}

<a id="L190"></a><span class="comment">// derefArray returns an expression of array type if the given</span>
<a id="L191"></a><span class="comment">// expression is a *array type.  Otherwise, returns the given</span>
<a id="L192"></a><span class="comment">// expression.</span>
<a id="L193"></a>func (a *expr) derefArray() *expr {
    <a id="L194"></a>if pt, ok := a.t.lit().(*PtrType); ok {
        <a id="L195"></a>if _, ok := pt.Elem.lit().(*ArrayType); ok {
            <a id="L196"></a>deref := a.compileStarExpr(a);
            <a id="L197"></a>if deref == nil {
                <a id="L198"></a>log.Crashf(&#34;failed to dereference *array&#34;)
            <a id="L199"></a>}
            <a id="L200"></a>return deref;
        <a id="L201"></a>}
    <a id="L202"></a>}
    <a id="L203"></a>return a;
<a id="L204"></a>}

<a id="L206"></a><span class="comment">/*</span>
<a id="L207"></a><span class="comment"> * Assignments</span>
<a id="L208"></a><span class="comment"> */</span>

<a id="L210"></a><span class="comment">// An assignCompiler compiles assignment operations.  Anything other</span>
<a id="L211"></a><span class="comment">// than short declarations should use the compileAssign wrapper.</span>
<a id="L212"></a><span class="comment">//</span>
<a id="L213"></a><span class="comment">// There are three valid types of assignment:</span>
<a id="L214"></a><span class="comment">// 1) T = T</span>
<a id="L215"></a><span class="comment">//    Assigning a single expression with single-valued type to a</span>
<a id="L216"></a><span class="comment">//    single-valued type.</span>
<a id="L217"></a><span class="comment">// 2) MT = T, T, ...</span>
<a id="L218"></a><span class="comment">//    Assigning multiple expressions with single-valued types to a</span>
<a id="L219"></a><span class="comment">//    multi-valued type.</span>
<a id="L220"></a><span class="comment">// 3) MT = MT</span>
<a id="L221"></a><span class="comment">//    Assigning a single expression with multi-valued type to a</span>
<a id="L222"></a><span class="comment">//    multi-valued type.</span>
<a id="L223"></a>type assignCompiler struct {
    <a id="L224"></a>*compiler;
    <a id="L225"></a>pos token.Position;
    <a id="L226"></a><span class="comment">// The RHS expressions.  This may include nil&#39;s for</span>
    <a id="L227"></a><span class="comment">// expressions that failed to compile.</span>
    <a id="L228"></a>rs  []*expr;
    <a id="L229"></a><span class="comment">// The (possibly unary) MultiType of the RHS.</span>
    <a id="L230"></a>rmt *MultiType;
    <a id="L231"></a><span class="comment">// Whether this is an unpack assignment (case 3).</span>
    <a id="L232"></a>isUnpack bool;
    <a id="L233"></a><span class="comment">// Whether map special assignment forms are allowed.</span>
    <a id="L234"></a>allowMap bool;
    <a id="L235"></a><span class="comment">// Whether this is a &#34;r, ok = a[x]&#34; assignment.</span>
    <a id="L236"></a>isMapUnpack bool;
    <a id="L237"></a><span class="comment">// The operation name to use in error messages, such as</span>
    <a id="L238"></a><span class="comment">// &#34;assignment&#34; or &#34;function call&#34;.</span>
    <a id="L239"></a>errOp string;
    <a id="L240"></a><span class="comment">// The name to use for positions in error messages, such as</span>
    <a id="L241"></a><span class="comment">// &#34;argument&#34;.</span>
    <a id="L242"></a>errPosName string;
<a id="L243"></a>}

<a id="L245"></a><span class="comment">// Type check the RHS of an assignment, returning a new assignCompiler</span>
<a id="L246"></a><span class="comment">// and indicating if the type check succeeded.  This always returns an</span>
<a id="L247"></a><span class="comment">// assignCompiler with rmt set, but if type checking fails, slots in</span>
<a id="L248"></a><span class="comment">// the MultiType may be nil.  If rs contains nil&#39;s, type checking will</span>
<a id="L249"></a><span class="comment">// fail and these expressions given a nil type.</span>
<a id="L250"></a>func (a *compiler) checkAssign(pos token.Position, rs []*expr, errOp, errPosName string) (*assignCompiler, bool) {
    <a id="L251"></a>c := &amp;assignCompiler{
        <a id="L252"></a>compiler: a,
        <a id="L253"></a>pos: pos,
        <a id="L254"></a>rs: rs,
        <a id="L255"></a>errOp: errOp,
        <a id="L256"></a>errPosName: errPosName,
    <a id="L257"></a>};

    <a id="L259"></a><span class="comment">// Is this an unpack?</span>
    <a id="L260"></a>if len(rs) == 1 &amp;&amp; rs[0] != nil {
        <a id="L261"></a>if rmt, isUnpack := rs[0].t.(*MultiType); isUnpack {
            <a id="L262"></a>c.rmt = rmt;
            <a id="L263"></a>c.isUnpack = true;
            <a id="L264"></a>return c, true;
        <a id="L265"></a>}
    <a id="L266"></a>}

    <a id="L268"></a><span class="comment">// Create MultiType for RHS and check that all RHS expressions</span>
    <a id="L269"></a><span class="comment">// are single-valued.</span>
    <a id="L270"></a>rts := make([]Type, len(rs));
    <a id="L271"></a>ok := true;
    <a id="L272"></a>for i, r := range rs {
        <a id="L273"></a>if r == nil {
            <a id="L274"></a>ok = false;
            <a id="L275"></a>continue;
        <a id="L276"></a>}

        <a id="L278"></a>if _, isMT := r.t.(*MultiType); isMT {
            <a id="L279"></a>r.diag(&#34;multi-valued expression not allowed in %s&#34;, errOp);
            <a id="L280"></a>ok = false;
            <a id="L281"></a>continue;
        <a id="L282"></a>}

        <a id="L284"></a>rts[i] = r.t;
    <a id="L285"></a>}

    <a id="L287"></a>c.rmt = NewMultiType(rts);
    <a id="L288"></a>return c, ok;
<a id="L289"></a>}

<a id="L291"></a>func (a *assignCompiler) allowMapForms(nls int) {
    <a id="L292"></a>a.allowMap = true;

    <a id="L294"></a><span class="comment">// Update unpacking info if this is r, ok = a[x]</span>
    <a id="L295"></a>if nls == 2 &amp;&amp; len(a.rs) == 1 &amp;&amp; a.rs[0] != nil &amp;&amp; a.rs[0].evalMapValue != nil {
        <a id="L296"></a>a.isUnpack = true;
        <a id="L297"></a>a.rmt = NewMultiType([]Type{a.rs[0].t, BoolType});
        <a id="L298"></a>a.isMapUnpack = true;
    <a id="L299"></a>}
<a id="L300"></a>}

<a id="L302"></a><span class="comment">// compile type checks and compiles an assignment operation, returning</span>
<a id="L303"></a><span class="comment">// a function that expects an l-value and the frame in which to</span>
<a id="L304"></a><span class="comment">// evaluate the RHS expressions.  The l-value must have exactly the</span>
<a id="L305"></a><span class="comment">// type given by lt.  Returns nil if type checking fails.</span>
<a id="L306"></a>func (a *assignCompiler) compile(b *block, lt Type) (func(Value, *Thread)) {
    <a id="L307"></a>lmt, isMT := lt.(*MultiType);
    <a id="L308"></a>rmt, isUnpack := a.rmt, a.isUnpack;

    <a id="L310"></a><span class="comment">// Create unary MultiType for single LHS</span>
    <a id="L311"></a>if !isMT {
        <a id="L312"></a>lmt = NewMultiType([]Type{lt})
    <a id="L313"></a>}

    <a id="L315"></a><span class="comment">// Check that the assignment count matches</span>
    <a id="L316"></a>lcount := len(lmt.Elems);
    <a id="L317"></a>rcount := len(rmt.Elems);
    <a id="L318"></a>if lcount != rcount {
        <a id="L319"></a>msg := &#34;not enough&#34;;
        <a id="L320"></a>pos := a.pos;
        <a id="L321"></a>if rcount &gt; lcount {
            <a id="L322"></a>msg = &#34;too many&#34;;
            <a id="L323"></a>if lcount &gt; 0 {
                <a id="L324"></a>pos = a.rs[lcount-1].pos
            <a id="L325"></a>}
        <a id="L326"></a>}
        <a id="L327"></a>a.diagAt(&amp;pos, &#34;%s %ss for %s\n\t%s\n\t%s&#34;, msg, a.errPosName, a.errOp, lt, rmt);
        <a id="L328"></a>return nil;
    <a id="L329"></a>}

    <a id="L331"></a>bad := false;

    <a id="L333"></a><span class="comment">// If this is an unpack, create a temporary to store the</span>
    <a id="L334"></a><span class="comment">// multi-value and replace the RHS with expressions to pull</span>
    <a id="L335"></a><span class="comment">// out values from the temporary.  Technically, this is only</span>
    <a id="L336"></a><span class="comment">// necessary when we need to perform assignment conversions.</span>
    <a id="L337"></a>var effect func(*Thread);
    <a id="L338"></a>if isUnpack {
        <a id="L339"></a><span class="comment">// This leaks a slot, but is definitely safe.</span>
        <a id="L340"></a>temp := b.DefineTemp(a.rmt);
        <a id="L341"></a>tempIdx := temp.Index;
        <a id="L342"></a>if tempIdx &lt; 0 {
            <a id="L343"></a>panicln(&#34;tempidx&#34;, tempIdx)
        <a id="L344"></a>}
        <a id="L345"></a>if a.isMapUnpack {
            <a id="L346"></a>rf := a.rs[0].evalMapValue;
            <a id="L347"></a>vt := a.rmt.Elems[0];
            <a id="L348"></a>effect = func(t *Thread) {
                <a id="L349"></a>m, k := rf(t);
                <a id="L350"></a>v := m.Elem(t, k);
                <a id="L351"></a>found := boolV(true);
                <a id="L352"></a>if v == nil {
                    <a id="L353"></a>found = boolV(false);
                    <a id="L354"></a>v = vt.Zero();
                <a id="L355"></a>}
                <a id="L356"></a>t.f.Vars[tempIdx] = multiV([]Value{v, &amp;found});
            <a id="L357"></a>};
        <a id="L358"></a>} else {
            <a id="L359"></a>rf := a.rs[0].asMulti();
            <a id="L360"></a>effect = func(t *Thread) { t.f.Vars[tempIdx] = multiV(rf(t)) };
        <a id="L361"></a>}
        <a id="L362"></a>orig := a.rs[0];
        <a id="L363"></a>a.rs = make([]*expr, len(a.rmt.Elems));
        <a id="L364"></a>for i, t := range a.rmt.Elems {
            <a id="L365"></a>if t.isIdeal() {
                <a id="L366"></a>log.Crashf(&#34;Right side of unpack contains ideal: %s&#34;, rmt)
            <a id="L367"></a>}
            <a id="L368"></a>a.rs[i] = orig.newExpr(t, orig.desc);
            <a id="L369"></a>index := i;
            <a id="L370"></a>a.rs[i].genValue(func(t *Thread) Value { return t.f.Vars[tempIdx].(multiV)[index] });
        <a id="L371"></a>}
    <a id="L372"></a>}
    <a id="L373"></a><span class="comment">// Now len(a.rs) == len(a.rmt) and we&#39;ve reduced any unpacking</span>
    <a id="L374"></a><span class="comment">// to multi-assignment.</span>

    <a id="L376"></a><span class="comment">// TODO(austin) Deal with assignment special cases.</span>

    <a id="L378"></a><span class="comment">// Values of any type may always be assigned to variables of</span>
    <a id="L379"></a><span class="comment">// compatible static type.</span>
    <a id="L380"></a>for i, lt := range lmt.Elems {
        <a id="L381"></a>rt := rmt.Elems[i];

        <a id="L383"></a><span class="comment">// When [an ideal is] (used in an expression) assigned</span>
        <a id="L384"></a><span class="comment">// to a variable or typed constant, the destination</span>
        <a id="L385"></a><span class="comment">// must be able to represent the assigned value.</span>
        <a id="L386"></a>if rt.isIdeal() {
            <a id="L387"></a>a.rs[i] = a.rs[i].convertTo(lmt.Elems[i]);
            <a id="L388"></a>if a.rs[i] == nil {
                <a id="L389"></a>bad = true;
                <a id="L390"></a>continue;
            <a id="L391"></a>}
            <a id="L392"></a>rt = a.rs[i].t;
        <a id="L393"></a>}

        <a id="L395"></a><span class="comment">// A pointer p to an array can be assigned to a slice</span>
        <a id="L396"></a><span class="comment">// variable v with compatible element type if the type</span>
        <a id="L397"></a><span class="comment">// of p or v is unnamed.</span>
        <a id="L398"></a>if rpt, ok := rt.lit().(*PtrType); ok {
            <a id="L399"></a>if at, ok := rpt.Elem.lit().(*ArrayType); ok {
                <a id="L400"></a>if lst, ok := lt.lit().(*SliceType); ok {
                    <a id="L401"></a>if lst.Elem.compat(at.Elem, false) &amp;&amp; (rt.lit() == Type(rt) || lt.lit() == Type(lt)) {
                        <a id="L402"></a>rf := a.rs[i].asPtr();
                        <a id="L403"></a>a.rs[i] = a.rs[i].newExpr(lt, a.rs[i].desc);
                        <a id="L404"></a>len := at.Len;
                        <a id="L405"></a>a.rs[i].eval = func(t *Thread) Slice { return Slice{rf(t).(ArrayValue), len, len} };
                        <a id="L406"></a>rt = a.rs[i].t;
                    <a id="L407"></a>}
                <a id="L408"></a>}
            <a id="L409"></a>}
        <a id="L410"></a>}

        <a id="L412"></a>if !lt.compat(rt, false) {
            <a id="L413"></a>if len(a.rs) == 1 {
                <a id="L414"></a>a.rs[0].diag(&#34;illegal operand types for %s\n\t%v\n\t%v&#34;, a.errOp, lt, rt)
            <a id="L415"></a>} else {
                <a id="L416"></a>a.rs[i].diag(&#34;illegal operand types in %s %d of %s\n\t%v\n\t%v&#34;, a.errPosName, i+1, a.errOp, lt, rt)
            <a id="L417"></a>}
            <a id="L418"></a>bad = true;
        <a id="L419"></a>}
    <a id="L420"></a>}
    <a id="L421"></a>if bad {
        <a id="L422"></a>return nil
    <a id="L423"></a>}

    <a id="L425"></a><span class="comment">// Compile</span>
    <a id="L426"></a>if !isMT {
        <a id="L427"></a><span class="comment">// Case 1</span>
        <a id="L428"></a>return genAssign(lt, a.rs[0])
    <a id="L429"></a>}
    <a id="L430"></a><span class="comment">// Case 2 or 3</span>
    <a id="L431"></a>as := make([]func(lv Value, t *Thread), len(a.rs));
    <a id="L432"></a>for i, r := range a.rs {
        <a id="L433"></a>as[i] = genAssign(lmt.Elems[i], r)
    <a id="L434"></a>}
    <a id="L435"></a>return func(lv Value, t *Thread) {
        <a id="L436"></a>if effect != nil {
            <a id="L437"></a>effect(t)
        <a id="L438"></a>}
        <a id="L439"></a>lmv := lv.(multiV);
        <a id="L440"></a>for i, a := range as {
            <a id="L441"></a>a(lmv[i], t)
        <a id="L442"></a>}
    <a id="L443"></a>};
<a id="L444"></a>}

<a id="L446"></a><span class="comment">// compileAssign compiles an assignment operation without the full</span>
<a id="L447"></a><span class="comment">// generality of an assignCompiler.  See assignCompiler for a</span>
<a id="L448"></a><span class="comment">// description of the arguments.</span>
<a id="L449"></a>func (a *compiler) compileAssign(pos token.Position, b *block, lt Type, rs []*expr, errOp, errPosName string) (func(Value, *Thread)) {
    <a id="L450"></a>ac, ok := a.checkAssign(pos, rs, errOp, errPosName);
    <a id="L451"></a>if !ok {
        <a id="L452"></a>return nil
    <a id="L453"></a>}
    <a id="L454"></a>return ac.compile(b, lt);
<a id="L455"></a>}

<a id="L457"></a><span class="comment">/*</span>
<a id="L458"></a><span class="comment"> * Expression compiler</span>
<a id="L459"></a><span class="comment"> */</span>

<a id="L461"></a><span class="comment">// An exprCompiler stores information used throughout the compilation</span>
<a id="L462"></a><span class="comment">// of a single expression.  It does not embed funcCompiler because</span>
<a id="L463"></a><span class="comment">// expressions can appear at top level.</span>
<a id="L464"></a>type exprCompiler struct {
    <a id="L465"></a>*compiler;
    <a id="L466"></a><span class="comment">// The block this expression is being compiled in.</span>
    <a id="L467"></a>block *block;
    <a id="L468"></a><span class="comment">// Whether this expression is used in a constant context.</span>
    <a id="L469"></a>constant bool;
<a id="L470"></a>}

<a id="L472"></a><span class="comment">// compile compiles an expression AST.  callCtx should be true if this</span>
<a id="L473"></a><span class="comment">// AST is in the function position of a function call node; it allows</span>
<a id="L474"></a><span class="comment">// the returned expression to be a type or a built-in function (which</span>
<a id="L475"></a><span class="comment">// otherwise result in errors).</span>
<a id="L476"></a>func (a *exprCompiler) compile(x ast.Expr, callCtx bool) *expr {
    <a id="L477"></a>ei := &amp;exprInfo{a.compiler, x.Pos()};

    <a id="L479"></a>switch x := x.(type) {
    <a id="L480"></a><span class="comment">// Literals</span>
    <a id="L481"></a>case *ast.BasicLit:
        <a id="L482"></a>switch x.Kind {
        <a id="L483"></a>case token.INT:
            <a id="L484"></a>return ei.compileIntLit(string(x.Value))
        <a id="L485"></a>case token.FLOAT:
            <a id="L486"></a>return ei.compileFloatLit(string(x.Value))
        <a id="L487"></a>case token.CHAR:
            <a id="L488"></a>return ei.compileCharLit(string(x.Value))
        <a id="L489"></a>case token.STRING:
            <a id="L490"></a>return ei.compileStringLit(string(x.Value))
        <a id="L491"></a>default:
            <a id="L492"></a>log.Crashf(&#34;unexpected basic literal type %v&#34;, x.Kind)
        <a id="L493"></a>}

    <a id="L495"></a>case *ast.CompositeLit:
        <a id="L496"></a>goto notimpl

    <a id="L498"></a>case *ast.FuncLit:
        <a id="L499"></a>decl := ei.compileFuncType(a.block, x.Type);
        <a id="L500"></a>if decl == nil {
            <a id="L501"></a><span class="comment">// TODO(austin) Try compiling the body,</span>
            <a id="L502"></a><span class="comment">// perhaps with dummy argument definitions</span>
            <a id="L503"></a>return nil
        <a id="L504"></a>}
        <a id="L505"></a>fn := ei.compileFunc(a.block, decl, x.Body);
        <a id="L506"></a>if fn == nil {
            <a id="L507"></a>return nil
        <a id="L508"></a>}
        <a id="L509"></a>if a.constant {
            <a id="L510"></a>a.diagAt(x, &#34;function literal used in constant expression&#34;);
            <a id="L511"></a>return nil;
        <a id="L512"></a>}
        <a id="L513"></a>return ei.compileFuncLit(decl, fn);

    <a id="L515"></a><span class="comment">// Types</span>
    <a id="L516"></a>case *ast.ArrayType:
        <a id="L517"></a><span class="comment">// TODO(austin) Use a multi-type case</span>
        <a id="L518"></a>goto typeexpr

    <a id="L520"></a>case *ast.ChanType:
        <a id="L521"></a>goto typeexpr

    <a id="L523"></a>case *ast.Ellipsis:
        <a id="L524"></a>goto typeexpr

    <a id="L526"></a>case *ast.FuncType:
        <a id="L527"></a>goto typeexpr

    <a id="L529"></a>case *ast.InterfaceType:
        <a id="L530"></a>goto typeexpr

    <a id="L532"></a>case *ast.MapType:
        <a id="L533"></a>goto typeexpr

    <a id="L535"></a><span class="comment">// Remaining expressions</span>
    <a id="L536"></a>case *ast.BadExpr:
        <a id="L537"></a><span class="comment">// Error already reported by parser</span>
        <a id="L538"></a>a.silentErrors++;
        <a id="L539"></a>return nil;

    <a id="L541"></a>case *ast.BinaryExpr:
        <a id="L542"></a>l, r := a.compile(x.X, false), a.compile(x.Y, false);
        <a id="L543"></a>if l == nil || r == nil {
            <a id="L544"></a>return nil
        <a id="L545"></a>}
        <a id="L546"></a>return ei.compileBinaryExpr(x.Op, l, r);

    <a id="L548"></a>case *ast.CallExpr:
        <a id="L549"></a>l := a.compile(x.Fun, true);
        <a id="L550"></a>args := make([]*expr, len(x.Args));
        <a id="L551"></a>bad := false;
        <a id="L552"></a>for i, arg := range x.Args {
            <a id="L553"></a>if i == 0 &amp;&amp; l != nil &amp;&amp; (l.t == Type(makeType) || l.t == Type(newType)) {
                <a id="L554"></a>argei := &amp;exprInfo{a.compiler, arg.Pos()};
                <a id="L555"></a>args[i] = argei.exprFromType(a.compileType(a.block, arg));
            <a id="L556"></a>} else {
                <a id="L557"></a>args[i] = a.compile(arg, false)
            <a id="L558"></a>}
            <a id="L559"></a>if args[i] == nil {
                <a id="L560"></a>bad = true
            <a id="L561"></a>}
        <a id="L562"></a>}
        <a id="L563"></a>if bad || l == nil {
            <a id="L564"></a>return nil
        <a id="L565"></a>}
        <a id="L566"></a>if a.constant {
            <a id="L567"></a>a.diagAt(x, &#34;function call in constant context&#34;);
            <a id="L568"></a>return nil;
        <a id="L569"></a>}

        <a id="L571"></a>if l.valType != nil {
            <a id="L572"></a>a.diagAt(x, &#34;type conversions not implemented&#34;);
            <a id="L573"></a>return nil;
        <a id="L574"></a>} else if ft, ok := l.t.(*FuncType); ok &amp;&amp; ft.builtin != &#34;&#34; {
            <a id="L575"></a>return ei.compileBuiltinCallExpr(a.block, ft, args)
        <a id="L576"></a>} else {
            <a id="L577"></a>return ei.compileCallExpr(a.block, l, args)
        <a id="L578"></a>}

    <a id="L580"></a>case *ast.Ident:
        <a id="L581"></a>return ei.compileIdent(a.block, a.constant, callCtx, x.Value)

    <a id="L583"></a>case *ast.IndexExpr:
        <a id="L584"></a>if x.End != nil {
            <a id="L585"></a>arr := a.compile(x.X, false);
            <a id="L586"></a>lo := a.compile(x.Index, false);
            <a id="L587"></a>hi := a.compile(x.End, false);
            <a id="L588"></a>if arr == nil || lo == nil || hi == nil {
                <a id="L589"></a>return nil
            <a id="L590"></a>}
            <a id="L591"></a>return ei.compileSliceExpr(arr, lo, hi);
        <a id="L592"></a>}
        <a id="L593"></a>l, r := a.compile(x.X, false), a.compile(x.Index, false);
        <a id="L594"></a>if l == nil || r == nil {
            <a id="L595"></a>return nil
        <a id="L596"></a>}
        <a id="L597"></a>return ei.compileIndexExpr(l, r);

    <a id="L599"></a>case *ast.KeyValueExpr:
        <a id="L600"></a>goto notimpl

    <a id="L602"></a>case *ast.ParenExpr:
        <a id="L603"></a>return a.compile(x.X, callCtx)

    <a id="L605"></a>case *ast.SelectorExpr:
        <a id="L606"></a>v := a.compile(x.X, false);
        <a id="L607"></a>if v == nil {
            <a id="L608"></a>return nil
        <a id="L609"></a>}
        <a id="L610"></a>return ei.compileSelectorExpr(v, x.Sel.Value);

    <a id="L612"></a>case *ast.StarExpr:
        <a id="L613"></a><span class="comment">// We pass down our call context because this could be</span>
        <a id="L614"></a><span class="comment">// a pointer type (and thus a type conversion)</span>
        <a id="L615"></a>v := a.compile(x.X, callCtx);
        <a id="L616"></a>if v == nil {
            <a id="L617"></a>return nil
        <a id="L618"></a>}
        <a id="L619"></a>if v.valType != nil {
            <a id="L620"></a><span class="comment">// Turns out this was a pointer type, not a dereference</span>
            <a id="L621"></a>return ei.exprFromType(NewPtrType(v.valType))
        <a id="L622"></a>}
        <a id="L623"></a>return ei.compileStarExpr(v);

    <a id="L625"></a>case *ast.StringList:
        <a id="L626"></a>strings := make([]*expr, len(x.Strings));
        <a id="L627"></a>bad := false;
        <a id="L628"></a>for i, s := range x.Strings {
            <a id="L629"></a>strings[i] = a.compile(s, false);
            <a id="L630"></a>if strings[i] == nil {
                <a id="L631"></a>bad = true
            <a id="L632"></a>}
        <a id="L633"></a>}
        <a id="L634"></a>if bad {
            <a id="L635"></a>return nil
        <a id="L636"></a>}
        <a id="L637"></a>return ei.compileStringList(strings);

    <a id="L639"></a>case *ast.StructType:
        <a id="L640"></a>goto notimpl

    <a id="L642"></a>case *ast.TypeAssertExpr:
        <a id="L643"></a>goto notimpl

    <a id="L645"></a>case *ast.UnaryExpr:
        <a id="L646"></a>v := a.compile(x.X, false);
        <a id="L647"></a>if v == nil {
            <a id="L648"></a>return nil
        <a id="L649"></a>}
        <a id="L650"></a>return ei.compileUnaryExpr(x.Op, v);
    <a id="L651"></a>}
    <a id="L652"></a>log.Crashf(&#34;unexpected ast node type %T&#34;, x);
    <a id="L653"></a>panic();

<a id="L655"></a>typeexpr:
    <a id="L656"></a>if !callCtx {
        <a id="L657"></a>a.diagAt(x, &#34;type used as expression&#34;);
        <a id="L658"></a>return nil;
    <a id="L659"></a>}
    <a id="L660"></a>return ei.exprFromType(a.compileType(a.block, x));

<a id="L662"></a>notimpl:
    <a id="L663"></a>a.diagAt(x, &#34;%T expression node not implemented&#34;, x);
    <a id="L664"></a>return nil;
<a id="L665"></a>}

<a id="L667"></a>func (a *exprInfo) exprFromType(t Type) *expr {
    <a id="L668"></a>if t == nil {
        <a id="L669"></a>return nil
    <a id="L670"></a>}
    <a id="L671"></a>expr := a.newExpr(nil, &#34;type&#34;);
    <a id="L672"></a>expr.valType = t;
    <a id="L673"></a>return expr;
<a id="L674"></a>}

<a id="L676"></a>func (a *exprInfo) compileIdent(b *block, constant bool, callCtx bool, name string) *expr {
    <a id="L677"></a>bl, level, def := b.Lookup(name);
    <a id="L678"></a>if def == nil {
        <a id="L679"></a>a.diag(&#34;%s: undefined&#34;, name);
        <a id="L680"></a>return nil;
    <a id="L681"></a>}
    <a id="L682"></a>switch def := def.(type) {
    <a id="L683"></a>case *Constant:
        <a id="L684"></a>expr := a.newExpr(def.Type, &#34;constant&#34;);
        <a id="L685"></a>if ft, ok := def.Type.(*FuncType); ok &amp;&amp; ft.builtin != &#34;&#34; {
            <a id="L686"></a><span class="comment">// XXX(Spec) I don&#39;t think anything says that</span>
            <a id="L687"></a><span class="comment">// built-in functions can&#39;t be used as values.</span>
            <a id="L688"></a>if !callCtx {
                <a id="L689"></a>a.diag(&#34;built-in function %s cannot be used as a value&#34;, ft.builtin);
                <a id="L690"></a>return nil;
            <a id="L691"></a>}
            <a id="L692"></a><span class="comment">// Otherwise, we leave the evaluators empty</span>
            <a id="L693"></a><span class="comment">// because this is handled specially</span>
        <a id="L694"></a>} else {
            <a id="L695"></a>expr.genConstant(def.Value)
        <a id="L696"></a>}
        <a id="L697"></a>return expr;
    <a id="L698"></a>case *Variable:
        <a id="L699"></a>if constant {
            <a id="L700"></a>a.diag(&#34;variable %s used in constant expression&#34;, name);
            <a id="L701"></a>return nil;
        <a id="L702"></a>}
        <a id="L703"></a>if bl.global {
            <a id="L704"></a>return a.compileGlobalVariable(def)
        <a id="L705"></a>}
        <a id="L706"></a>return a.compileVariable(level, def);
    <a id="L707"></a>case Type:
        <a id="L708"></a>if callCtx {
            <a id="L709"></a>return a.exprFromType(def)
        <a id="L710"></a>}
        <a id="L711"></a>a.diag(&#34;type %v used as expression&#34;, name);
        <a id="L712"></a>return nil;
    <a id="L713"></a>}
    <a id="L714"></a>log.Crashf(&#34;name %s has unknown type %T&#34;, name, def);
    <a id="L715"></a>panic();
<a id="L716"></a>}

<a id="L718"></a>func (a *exprInfo) compileVariable(level int, v *Variable) *expr {
    <a id="L719"></a>if v.Type == nil {
        <a id="L720"></a><span class="comment">// Placeholder definition from an earlier error</span>
        <a id="L721"></a>a.silentErrors++;
        <a id="L722"></a>return nil;
    <a id="L723"></a>}
    <a id="L724"></a>expr := a.newExpr(v.Type, &#34;variable&#34;);
    <a id="L725"></a>expr.genIdentOp(level, v.Index);
    <a id="L726"></a>return expr;
<a id="L727"></a>}

<a id="L729"></a>func (a *exprInfo) compileGlobalVariable(v *Variable) *expr {
    <a id="L730"></a>if v.Type == nil {
        <a id="L731"></a><span class="comment">// Placeholder definition from an earlier error</span>
        <a id="L732"></a>a.silentErrors++;
        <a id="L733"></a>return nil;
    <a id="L734"></a>}
    <a id="L735"></a>if v.Init == nil {
        <a id="L736"></a>v.Init = v.Type.Zero()
    <a id="L737"></a>}
    <a id="L738"></a>expr := a.newExpr(v.Type, &#34;variable&#34;);
    <a id="L739"></a>val := v.Init;
    <a id="L740"></a>expr.genValue(func(t *Thread) Value { return val });
    <a id="L741"></a>return expr;
<a id="L742"></a>}

<a id="L744"></a>func (a *exprInfo) compileIdealInt(i *bignum.Integer, desc string) *expr {
    <a id="L745"></a>expr := a.newExpr(IdealIntType, desc);
    <a id="L746"></a>expr.eval = func() *bignum.Integer { return i };
    <a id="L747"></a>return expr;
<a id="L748"></a>}

<a id="L750"></a>func (a *exprInfo) compileIntLit(lit string) *expr {
    <a id="L751"></a>i, _, _ := bignum.IntFromString(lit, 0);
    <a id="L752"></a>return a.compileIdealInt(i, &#34;integer literal&#34;);
<a id="L753"></a>}

<a id="L755"></a>func (a *exprInfo) compileCharLit(lit string) *expr {
    <a id="L756"></a>if lit[0] != &#39;\&#39;&#39; {
        <a id="L757"></a><span class="comment">// Caught by parser</span>
        <a id="L758"></a>a.silentErrors++;
        <a id="L759"></a>return nil;
    <a id="L760"></a>}
    <a id="L761"></a>v, _, tail, err := strconv.UnquoteChar(lit[1:len(lit)], &#39;\&#39;&#39;);
    <a id="L762"></a>if err != nil || tail != &#34;&#39;&#34; {
        <a id="L763"></a><span class="comment">// Caught by parser</span>
        <a id="L764"></a>a.silentErrors++;
        <a id="L765"></a>return nil;
    <a id="L766"></a>}
    <a id="L767"></a>return a.compileIdealInt(bignum.Int(int64(v)), &#34;character literal&#34;);
<a id="L768"></a>}

<a id="L770"></a>func (a *exprInfo) compileFloatLit(lit string) *expr {
    <a id="L771"></a>f, _, n := bignum.RatFromString(lit, 0);
    <a id="L772"></a>if n != len(lit) {
        <a id="L773"></a>log.Crashf(&#34;malformed float literal %s at %v passed parser&#34;, lit, a.pos)
    <a id="L774"></a>}
    <a id="L775"></a>expr := a.newExpr(IdealFloatType, &#34;float literal&#34;);
    <a id="L776"></a>expr.eval = func() *bignum.Rational { return f };
    <a id="L777"></a>return expr;
<a id="L778"></a>}

<a id="L780"></a>func (a *exprInfo) compileString(s string) *expr {
    <a id="L781"></a><span class="comment">// Ideal strings don&#39;t have a named type but they are</span>
    <a id="L782"></a><span class="comment">// compatible with type string.</span>

    <a id="L784"></a><span class="comment">// TODO(austin) Use unnamed string type.</span>
    <a id="L785"></a>expr := a.newExpr(StringType, &#34;string literal&#34;);
    <a id="L786"></a>expr.eval = func(*Thread) string { return s };
    <a id="L787"></a>return expr;
<a id="L788"></a>}

<a id="L790"></a>func (a *exprInfo) compileStringLit(lit string) *expr {
    <a id="L791"></a>s, err := strconv.Unquote(lit);
    <a id="L792"></a>if err != nil {
        <a id="L793"></a>a.diag(&#34;illegal string literal, %v&#34;, err);
        <a id="L794"></a>return nil;
    <a id="L795"></a>}
    <a id="L796"></a>return a.compileString(s);
<a id="L797"></a>}

<a id="L799"></a>func (a *exprInfo) compileStringList(list []*expr) *expr {
    <a id="L800"></a>ss := make([]string, len(list));
    <a id="L801"></a>for i, s := range list {
        <a id="L802"></a>ss[i] = s.asString()(nil)
    <a id="L803"></a>}
    <a id="L804"></a>return a.compileString(strings.Join(ss, &#34;&#34;));
<a id="L805"></a>}

<a id="L807"></a>func (a *exprInfo) compileFuncLit(decl *FuncDecl, fn func(*Thread) Func) *expr {
    <a id="L808"></a>expr := a.newExpr(decl.Type, &#34;function literal&#34;);
    <a id="L809"></a>expr.eval = fn;
    <a id="L810"></a>return expr;
<a id="L811"></a>}

<a id="L813"></a>func (a *exprInfo) compileSelectorExpr(v *expr, name string) *expr {
    <a id="L814"></a><span class="comment">// mark marks a field that matches the selector name.  It</span>
    <a id="L815"></a><span class="comment">// tracks the best depth found so far and whether more than</span>
    <a id="L816"></a><span class="comment">// one field has been found at that depth.</span>
    <a id="L817"></a>bestDepth := -1;
    <a id="L818"></a>ambig := false;
    <a id="L819"></a>amberr := &#34;&#34;;
    <a id="L820"></a>mark := func(depth int, pathName string) {
        <a id="L821"></a>switch {
        <a id="L822"></a>case bestDepth == -1 || depth &lt; bestDepth:
            <a id="L823"></a>bestDepth = depth;
            <a id="L824"></a>ambig = false;
            <a id="L825"></a>amberr = &#34;&#34;;

        <a id="L827"></a>case depth == bestDepth:
            <a id="L828"></a>ambig = true

        <a id="L830"></a>default:
            <a id="L831"></a>log.Crashf(&#34;Marked field at depth %d, but already found one at depth %d&#34;, depth, bestDepth)
        <a id="L832"></a>}
        <a id="L833"></a>amberr += &#34;\n\t&#34; + pathName[1:len(pathName)];
    <a id="L834"></a>};

    <a id="L836"></a>visited := make(map[Type]bool);

    <a id="L838"></a><span class="comment">// find recursively searches for the named field, starting at</span>
    <a id="L839"></a><span class="comment">// type t.  If it finds the named field, it returns a function</span>
    <a id="L840"></a><span class="comment">// which takes an expr that represents a value of type &#39;t&#39; and</span>
    <a id="L841"></a><span class="comment">// returns an expr that retrieves the named field.  We delay</span>
    <a id="L842"></a><span class="comment">// expr construction to avoid producing lots of useless expr&#39;s</span>
    <a id="L843"></a><span class="comment">// as we search.</span>
    <a id="L844"></a><span class="comment">//</span>
    <a id="L845"></a><span class="comment">// TODO(austin) Now that the expression compiler works on</span>
    <a id="L846"></a><span class="comment">// semantic values instead of AST&#39;s, there should be a much</span>
    <a id="L847"></a><span class="comment">// better way of doing this.</span>
    <a id="L848"></a>var find func(Type, int, string) (func(*expr) *expr);
    <a id="L849"></a>find = func(t Type, depth int, pathName string) (func(*expr) *expr) {
        <a id="L850"></a><span class="comment">// Don&#39;t bother looking if we&#39;ve found something shallower</span>
        <a id="L851"></a>if bestDepth != -1 &amp;&amp; bestDepth &lt; depth {
            <a id="L852"></a>return nil
        <a id="L853"></a>}

        <a id="L855"></a><span class="comment">// Don&#39;t check the same type twice and avoid loops</span>
        <a id="L856"></a>if _, ok := visited[t]; ok {
            <a id="L857"></a>return nil
        <a id="L858"></a>}
        <a id="L859"></a>visited[t] = true;

        <a id="L861"></a><span class="comment">// Implicit dereference</span>
        <a id="L862"></a>deref := false;
        <a id="L863"></a>if ti, ok := t.(*PtrType); ok {
            <a id="L864"></a>deref = true;
            <a id="L865"></a>t = ti.Elem;
        <a id="L866"></a>}

        <a id="L868"></a><span class="comment">// If it&#39;s a named type, look for methods</span>
        <a id="L869"></a>if ti, ok := t.(*NamedType); ok {
            <a id="L870"></a>_, ok := ti.methods[name];
            <a id="L871"></a>if ok {
                <a id="L872"></a>mark(depth, pathName+&#34;.&#34;+name);
                <a id="L873"></a>log.Crash(&#34;Methods not implemented&#34;);
            <a id="L874"></a>}
            <a id="L875"></a>t = ti.Def;
        <a id="L876"></a>}

        <a id="L878"></a><span class="comment">// If it&#39;s a struct type, check fields and embedded types</span>
        <a id="L879"></a>var builder func(*expr) *expr;
        <a id="L880"></a>if t, ok := t.(*StructType); ok {
            <a id="L881"></a>for i, f := range t.Elems {
                <a id="L882"></a>var sub func(*expr) *expr;
                <a id="L883"></a>switch {
                <a id="L884"></a>case f.Name == name:
                    <a id="L885"></a>mark(depth, pathName+&#34;.&#34;+name);
                    <a id="L886"></a>sub = func(e *expr) *expr { return e };

                <a id="L888"></a>case f.Anonymous:
                    <a id="L889"></a>sub = find(f.Type, depth+1, pathName+&#34;.&#34;+f.Name);
                    <a id="L890"></a>if sub == nil {
                        <a id="L891"></a>continue
                    <a id="L892"></a>}

                <a id="L894"></a>default:
                    <a id="L895"></a>continue
                <a id="L896"></a>}

                <a id="L898"></a><span class="comment">// We found something.  Create a</span>
                <a id="L899"></a><span class="comment">// builder for accessing this field.</span>
                <a id="L900"></a>ft := f.Type;
                <a id="L901"></a>index := i;
                <a id="L902"></a>builder = func(parent *expr) *expr {
                    <a id="L903"></a>if deref {
                        <a id="L904"></a>parent = a.compileStarExpr(parent)
                    <a id="L905"></a>}
                    <a id="L906"></a>expr := a.newExpr(ft, &#34;selector expression&#34;);
                    <a id="L907"></a>pf := parent.asStruct();
                    <a id="L908"></a>evalAddr := func(t *Thread) Value { return pf(t).Field(t, index) };
                    <a id="L909"></a>expr.genValue(evalAddr);
                    <a id="L910"></a>return sub(expr);
                <a id="L911"></a>};
            <a id="L912"></a>}
        <a id="L913"></a>}

        <a id="L915"></a>return builder;
    <a id="L916"></a>};

    <a id="L918"></a>builder := find(v.t, 0, &#34;&#34;);
    <a id="L919"></a>if builder == nil {
        <a id="L920"></a>a.diag(&#34;type %v has no field or method %s&#34;, v.t, name);
        <a id="L921"></a>return nil;
    <a id="L922"></a>}
    <a id="L923"></a>if ambig {
        <a id="L924"></a>a.diag(&#34;field %s is ambiguous in type %v%s&#34;, name, v.t, amberr);
        <a id="L925"></a>return nil;
    <a id="L926"></a>}

    <a id="L928"></a>return builder(v);
<a id="L929"></a>}

<a id="L931"></a>func (a *exprInfo) compileSliceExpr(arr, lo, hi *expr) *expr {
    <a id="L932"></a><span class="comment">// Type check object</span>
    <a id="L933"></a>arr = arr.derefArray();

    <a id="L935"></a>var at Type;
    <a id="L936"></a>var maxIndex int64 = -1;

    <a id="L938"></a>switch lt := arr.t.lit().(type) {
    <a id="L939"></a>case *ArrayType:
        <a id="L940"></a>at = NewSliceType(lt.Elem);
        <a id="L941"></a>maxIndex = lt.Len;

    <a id="L943"></a>case *SliceType:
        <a id="L944"></a>at = lt

    <a id="L946"></a>case *stringType:
        <a id="L947"></a>at = lt

    <a id="L949"></a>default:
        <a id="L950"></a>a.diag(&#34;cannot slice %v&#34;, arr.t);
        <a id="L951"></a>return nil;
    <a id="L952"></a>}

    <a id="L954"></a><span class="comment">// Type check index and convert to int</span>
    <a id="L955"></a><span class="comment">// XXX(Spec) It&#39;s unclear if ideal floats with no</span>
    <a id="L956"></a><span class="comment">// fractional part are allowed here.  6g allows it.  I</span>
    <a id="L957"></a><span class="comment">// believe that&#39;s wrong.</span>
    <a id="L958"></a>lo = lo.convertToInt(maxIndex, &#34;slice&#34;, &#34;slice&#34;);
    <a id="L959"></a>hi = hi.convertToInt(maxIndex, &#34;slice&#34;, &#34;slice&#34;);
    <a id="L960"></a>if lo == nil || hi == nil {
        <a id="L961"></a>return nil
    <a id="L962"></a>}

    <a id="L964"></a>expr := a.newExpr(at, &#34;slice expression&#34;);

    <a id="L966"></a><span class="comment">// Compile</span>
    <a id="L967"></a>lof := lo.asInt();
    <a id="L968"></a>hif := hi.asInt();
    <a id="L969"></a>switch lt := arr.t.lit().(type) {
    <a id="L970"></a>case *ArrayType:
        <a id="L971"></a>arrf := arr.asArray();
        <a id="L972"></a>bound := lt.Len;
        <a id="L973"></a>expr.eval = func(t *Thread) Slice {
            <a id="L974"></a>arr, lo, hi := arrf(t), lof(t), hif(t);
            <a id="L975"></a>if lo &gt; hi || hi &gt; bound || lo &lt; 0 {
                <a id="L976"></a>t.Abort(SliceError{lo, hi, bound})
            <a id="L977"></a>}
            <a id="L978"></a>return Slice{arr.Sub(lo, bound-lo), hi - lo, bound - lo};
        <a id="L979"></a>};

    <a id="L981"></a>case *SliceType:
        <a id="L982"></a>arrf := arr.asSlice();
        <a id="L983"></a>expr.eval = func(t *Thread) Slice {
            <a id="L984"></a>arr, lo, hi := arrf(t), lof(t), hif(t);
            <a id="L985"></a>if lo &gt; hi || hi &gt; arr.Cap || lo &lt; 0 {
                <a id="L986"></a>t.Abort(SliceError{lo, hi, arr.Cap})
            <a id="L987"></a>}
            <a id="L988"></a>return Slice{arr.Base.Sub(lo, arr.Cap-lo), hi - lo, arr.Cap - lo};
        <a id="L989"></a>};

    <a id="L991"></a>case *stringType:
        <a id="L992"></a>arrf := arr.asString();
        <a id="L993"></a><span class="comment">// TODO(austin) This pulls over the whole string in a</span>
        <a id="L994"></a><span class="comment">// remote setting, instead of creating a substring backed</span>
        <a id="L995"></a><span class="comment">// by remote memory.</span>
        <a id="L996"></a>expr.eval = func(t *Thread) string {
            <a id="L997"></a>arr, lo, hi := arrf(t), lof(t), hif(t);
            <a id="L998"></a>if lo &gt; hi || hi &gt; int64(len(arr)) || lo &lt; 0 {
                <a id="L999"></a>t.Abort(SliceError{lo, hi, int64(len(arr))})
            <a id="L1000"></a>}
            <a id="L1001"></a>return arr[lo:hi];
        <a id="L1002"></a>};

    <a id="L1004"></a>default:
        <a id="L1005"></a>log.Crashf(&#34;unexpected left operand type %T&#34;, arr.t.lit())
    <a id="L1006"></a>}

    <a id="L1008"></a>return expr;
<a id="L1009"></a>}

<a id="L1011"></a>func (a *exprInfo) compileIndexExpr(l, r *expr) *expr {
    <a id="L1012"></a><span class="comment">// Type check object</span>
    <a id="L1013"></a>l = l.derefArray();

    <a id="L1015"></a>var at Type;
    <a id="L1016"></a>intIndex := false;
    <a id="L1017"></a>var maxIndex int64 = -1;

    <a id="L1019"></a>switch lt := l.t.lit().(type) {
    <a id="L1020"></a>case *ArrayType:
        <a id="L1021"></a>at = lt.Elem;
        <a id="L1022"></a>intIndex = true;
        <a id="L1023"></a>maxIndex = lt.Len;

    <a id="L1025"></a>case *SliceType:
        <a id="L1026"></a>at = lt.Elem;
        <a id="L1027"></a>intIndex = true;

    <a id="L1029"></a>case *stringType:
        <a id="L1030"></a>at = Uint8Type;
        <a id="L1031"></a>intIndex = true;

    <a id="L1033"></a>case *MapType:
        <a id="L1034"></a>at = lt.Elem;
        <a id="L1035"></a>if r.t.isIdeal() {
            <a id="L1036"></a>r = r.convertTo(lt.Key);
            <a id="L1037"></a>if r == nil {
                <a id="L1038"></a>return nil
            <a id="L1039"></a>}
        <a id="L1040"></a>}
        <a id="L1041"></a>if !lt.Key.compat(r.t, false) {
            <a id="L1042"></a>a.diag(&#34;cannot use %s as index into %s&#34;, r.t, lt);
            <a id="L1043"></a>return nil;
        <a id="L1044"></a>}

    <a id="L1046"></a>default:
        <a id="L1047"></a>a.diag(&#34;cannot index into %v&#34;, l.t);
        <a id="L1048"></a>return nil;
    <a id="L1049"></a>}

    <a id="L1051"></a><span class="comment">// Type check index and convert to int if necessary</span>
    <a id="L1052"></a>if intIndex {
        <a id="L1053"></a><span class="comment">// XXX(Spec) It&#39;s unclear if ideal floats with no</span>
        <a id="L1054"></a><span class="comment">// fractional part are allowed here.  6g allows it.  I</span>
        <a id="L1055"></a><span class="comment">// believe that&#39;s wrong.</span>
        <a id="L1056"></a>r = r.convertToInt(maxIndex, &#34;index&#34;, &#34;index&#34;);
        <a id="L1057"></a>if r == nil {
            <a id="L1058"></a>return nil
        <a id="L1059"></a>}
    <a id="L1060"></a>}

    <a id="L1062"></a>expr := a.newExpr(at, &#34;index expression&#34;);

    <a id="L1064"></a><span class="comment">// Compile</span>
    <a id="L1065"></a>switch lt := l.t.lit().(type) {
    <a id="L1066"></a>case *ArrayType:
        <a id="L1067"></a>lf := l.asArray();
        <a id="L1068"></a>rf := r.asInt();
        <a id="L1069"></a>bound := lt.Len;
        <a id="L1070"></a>expr.genValue(func(t *Thread) Value {
            <a id="L1071"></a>l, r := lf(t), rf(t);
            <a id="L1072"></a>if r &lt; 0 || r &gt;= bound {
                <a id="L1073"></a>t.Abort(IndexError{r, bound})
            <a id="L1074"></a>}
            <a id="L1075"></a>return l.Elem(t, r);
        <a id="L1076"></a>});

    <a id="L1078"></a>case *SliceType:
        <a id="L1079"></a>lf := l.asSlice();
        <a id="L1080"></a>rf := r.asInt();
        <a id="L1081"></a>expr.genValue(func(t *Thread) Value {
            <a id="L1082"></a>l, r := lf(t), rf(t);
            <a id="L1083"></a>if l.Base == nil {
                <a id="L1084"></a>t.Abort(NilPointerError{})
            <a id="L1085"></a>}
            <a id="L1086"></a>if r &lt; 0 || r &gt;= l.Len {
                <a id="L1087"></a>t.Abort(IndexError{r, l.Len})
            <a id="L1088"></a>}
            <a id="L1089"></a>return l.Base.Elem(t, r);
        <a id="L1090"></a>});

    <a id="L1092"></a>case *stringType:
        <a id="L1093"></a>lf := l.asString();
        <a id="L1094"></a>rf := r.asInt();
        <a id="L1095"></a><span class="comment">// TODO(austin) This pulls over the whole string in a</span>
        <a id="L1096"></a><span class="comment">// remote setting, instead of just the one character.</span>
        <a id="L1097"></a>expr.eval = func(t *Thread) uint64 {
            <a id="L1098"></a>l, r := lf(t), rf(t);
            <a id="L1099"></a>if r &lt; 0 || r &gt;= int64(len(l)) {
                <a id="L1100"></a>t.Abort(IndexError{r, int64(len(l))})
            <a id="L1101"></a>}
            <a id="L1102"></a>return uint64(l[r]);
        <a id="L1103"></a>};

    <a id="L1105"></a>case *MapType:
        <a id="L1106"></a>lf := l.asMap();
        <a id="L1107"></a>rf := r.asInterface();
        <a id="L1108"></a>expr.genValue(func(t *Thread) Value {
            <a id="L1109"></a>m := lf(t);
            <a id="L1110"></a>k := rf(t);
            <a id="L1111"></a>if m == nil {
                <a id="L1112"></a>t.Abort(NilPointerError{})
            <a id="L1113"></a>}
            <a id="L1114"></a>e := m.Elem(t, k);
            <a id="L1115"></a>if e == nil {
                <a id="L1116"></a>t.Abort(KeyError{k})
            <a id="L1117"></a>}
            <a id="L1118"></a>return e;
        <a id="L1119"></a>});
        <a id="L1120"></a><span class="comment">// genValue makes things addressable, but map values</span>
        <a id="L1121"></a><span class="comment">// aren&#39;t addressable.</span>
        <a id="L1122"></a>expr.evalAddr = nil;
        <a id="L1123"></a>expr.evalMapValue = func(t *Thread) (Map, interface{}) {
            <a id="L1124"></a><span class="comment">// TODO(austin) Key check?  nil check?</span>
            <a id="L1125"></a>return lf(t), rf(t)
        <a id="L1126"></a>};

    <a id="L1128"></a>default:
        <a id="L1129"></a>log.Crashf(&#34;unexpected left operand type %T&#34;, l.t.lit())
    <a id="L1130"></a>}

    <a id="L1132"></a>return expr;
<a id="L1133"></a>}

<a id="L1135"></a>func (a *exprInfo) compileCallExpr(b *block, l *expr, as []*expr) *expr {
    <a id="L1136"></a><span class="comment">// TODO(austin) Variadic functions.</span>

    <a id="L1138"></a><span class="comment">// Type check</span>

    <a id="L1140"></a><span class="comment">// XXX(Spec) Calling a named function type is okay.  I really</span>
    <a id="L1141"></a><span class="comment">// think there needs to be a general discussion of named</span>
    <a id="L1142"></a><span class="comment">// types.  A named type creates a new, distinct type, but the</span>
    <a id="L1143"></a><span class="comment">// type of that type is still whatever it&#39;s defined to.  Thus,</span>
    <a id="L1144"></a><span class="comment">// in &#34;type Foo int&#34;, Foo is still an integer type and in</span>
    <a id="L1145"></a><span class="comment">// &#34;type Foo func()&#34;, Foo is a function type.</span>
    <a id="L1146"></a>lt, ok := l.t.lit().(*FuncType);
    <a id="L1147"></a>if !ok {
        <a id="L1148"></a>a.diag(&#34;cannot call non-function type %v&#34;, l.t);
        <a id="L1149"></a>return nil;
    <a id="L1150"></a>}

    <a id="L1152"></a><span class="comment">// The arguments must be single-valued expressions assignment</span>
    <a id="L1153"></a><span class="comment">// compatible with the parameters of F.</span>
    <a id="L1154"></a><span class="comment">//</span>
    <a id="L1155"></a><span class="comment">// XXX(Spec) The spec is wrong.  It can also be a single</span>
    <a id="L1156"></a><span class="comment">// multi-valued expression.</span>
    <a id="L1157"></a>nin := len(lt.In);
    <a id="L1158"></a>assign := a.compileAssign(a.pos, b, NewMultiType(lt.In), as, &#34;function call&#34;, &#34;argument&#34;);
    <a id="L1159"></a>if assign == nil {
        <a id="L1160"></a>return nil
    <a id="L1161"></a>}

    <a id="L1163"></a>var t Type;
    <a id="L1164"></a>nout := len(lt.Out);
    <a id="L1165"></a>switch nout {
    <a id="L1166"></a>case 0:
        <a id="L1167"></a>t = EmptyType
    <a id="L1168"></a>case 1:
        <a id="L1169"></a>t = lt.Out[0]
    <a id="L1170"></a>default:
        <a id="L1171"></a>t = NewMultiType(lt.Out)
    <a id="L1172"></a>}
    <a id="L1173"></a>expr := a.newExpr(t, &#34;function call&#34;);

    <a id="L1175"></a><span class="comment">// Gather argument and out types to initialize frame variables</span>
    <a id="L1176"></a>vts := make([]Type, nin+nout);
    <a id="L1177"></a>for i, t := range lt.In {
        <a id="L1178"></a>vts[i] = t
    <a id="L1179"></a>}
    <a id="L1180"></a>for i, t := range lt.Out {
        <a id="L1181"></a>vts[i+nin] = t
    <a id="L1182"></a>}

    <a id="L1184"></a><span class="comment">// Compile</span>
    <a id="L1185"></a>lf := l.asFunc();
    <a id="L1186"></a>call := func(t *Thread) []Value {
        <a id="L1187"></a>fun := lf(t);
        <a id="L1188"></a>fr := fun.NewFrame();
        <a id="L1189"></a>for i, t := range vts {
            <a id="L1190"></a>fr.Vars[i] = t.Zero()
        <a id="L1191"></a>}
        <a id="L1192"></a>assign(multiV(fr.Vars[0:nin]), t);
        <a id="L1193"></a>oldf := t.f;
        <a id="L1194"></a>t.f = fr;
        <a id="L1195"></a>fun.Call(t);
        <a id="L1196"></a>t.f = oldf;
        <a id="L1197"></a>return fr.Vars[nin : nin+nout];
    <a id="L1198"></a>};
    <a id="L1199"></a>expr.genFuncCall(call);

    <a id="L1201"></a>return expr;
<a id="L1202"></a>}

<a id="L1204"></a>func (a *exprInfo) compileBuiltinCallExpr(b *block, ft *FuncType, as []*expr) *expr {
    <a id="L1205"></a>checkCount := func(min, max int) bool {
        <a id="L1206"></a>if len(as) &lt; min {
            <a id="L1207"></a>a.diag(&#34;not enough arguments to %s&#34;, ft.builtin);
            <a id="L1208"></a>return false;
        <a id="L1209"></a>} else if len(as) &gt; max {
            <a id="L1210"></a>a.diag(&#34;too many arguments to %s&#34;, ft.builtin);
            <a id="L1211"></a>return false;
        <a id="L1212"></a>}
        <a id="L1213"></a>return true;
    <a id="L1214"></a>};

    <a id="L1216"></a>switch ft {
    <a id="L1217"></a>case capType:
        <a id="L1218"></a>if !checkCount(1, 1) {
            <a id="L1219"></a>return nil
        <a id="L1220"></a>}
        <a id="L1221"></a>arg := as[0].derefArray();
        <a id="L1222"></a>expr := a.newExpr(IntType, &#34;function call&#34;);
        <a id="L1223"></a>switch t := arg.t.lit().(type) {
        <a id="L1224"></a>case *ArrayType:
            <a id="L1225"></a><span class="comment">// TODO(austin) It would be nice if this could</span>
            <a id="L1226"></a><span class="comment">// be a constant int.</span>
            <a id="L1227"></a>v := t.Len;
            <a id="L1228"></a>expr.eval = func(t *Thread) int64 { return v };

        <a id="L1230"></a>case *SliceType:
            <a id="L1231"></a>vf := arg.asSlice();
            <a id="L1232"></a>expr.eval = func(t *Thread) int64 { return vf(t).Cap };

        <a id="L1234"></a><span class="comment">//case *ChanType:</span>

        <a id="L1236"></a>default:
            <a id="L1237"></a>a.diag(&#34;illegal argument type for cap function\n\t%v&#34;, arg.t);
            <a id="L1238"></a>return nil;
        <a id="L1239"></a>}
        <a id="L1240"></a>return expr;

    <a id="L1242"></a>case lenType:
        <a id="L1243"></a>if !checkCount(1, 1) {
            <a id="L1244"></a>return nil
        <a id="L1245"></a>}
        <a id="L1246"></a>arg := as[0].derefArray();
        <a id="L1247"></a>expr := a.newExpr(IntType, &#34;function call&#34;);
        <a id="L1248"></a>switch t := arg.t.lit().(type) {
        <a id="L1249"></a>case *stringType:
            <a id="L1250"></a>vf := arg.asString();
            <a id="L1251"></a>expr.eval = func(t *Thread) int64 { return int64(len(vf(t))) };

        <a id="L1253"></a>case *ArrayType:
            <a id="L1254"></a><span class="comment">// TODO(austin) It would be nice if this could</span>
            <a id="L1255"></a><span class="comment">// be a constant int.</span>
            <a id="L1256"></a>v := t.Len;
            <a id="L1257"></a>expr.eval = func(t *Thread) int64 { return v };

        <a id="L1259"></a>case *SliceType:
            <a id="L1260"></a>vf := arg.asSlice();
            <a id="L1261"></a>expr.eval = func(t *Thread) int64 { return vf(t).Len };

        <a id="L1263"></a>case *MapType:
            <a id="L1264"></a>vf := arg.asMap();
            <a id="L1265"></a>expr.eval = func(t *Thread) int64 {
                <a id="L1266"></a><span class="comment">// XXX(Spec) What&#39;s the len of an</span>
                <a id="L1267"></a><span class="comment">// uninitialized map?</span>
                <a id="L1268"></a>m := vf(t);
                <a id="L1269"></a>if m == nil {
                    <a id="L1270"></a>return 0
                <a id="L1271"></a>}
                <a id="L1272"></a>return m.Len(t);
            <a id="L1273"></a>};

        <a id="L1275"></a><span class="comment">//case *ChanType:</span>

        <a id="L1277"></a>default:
            <a id="L1278"></a>a.diag(&#34;illegal argument type for len function\n\t%v&#34;, arg.t);
            <a id="L1279"></a>return nil;
        <a id="L1280"></a>}
        <a id="L1281"></a>return expr;

    <a id="L1283"></a>case makeType:
        <a id="L1284"></a>if !checkCount(1, 3) {
            <a id="L1285"></a>return nil
        <a id="L1286"></a>}
        <a id="L1287"></a><span class="comment">// XXX(Spec) What are the types of the</span>
        <a id="L1288"></a><span class="comment">// arguments?  Do they have to be ints?  6g</span>
        <a id="L1289"></a><span class="comment">// accepts any integral type.</span>
        <a id="L1290"></a>var lenexpr, capexpr *expr;
        <a id="L1291"></a>var lenf, capf func(*Thread) int64;
        <a id="L1292"></a>if len(as) &gt; 1 {
            <a id="L1293"></a>lenexpr = as[1].convertToInt(-1, &#34;length&#34;, &#34;make function&#34;);
            <a id="L1294"></a>if lenexpr == nil {
                <a id="L1295"></a>return nil
            <a id="L1296"></a>}
            <a id="L1297"></a>lenf = lenexpr.asInt();
        <a id="L1298"></a>}
        <a id="L1299"></a>if len(as) &gt; 2 {
            <a id="L1300"></a>capexpr = as[2].convertToInt(-1, &#34;capacity&#34;, &#34;make function&#34;);
            <a id="L1301"></a>if capexpr == nil {
                <a id="L1302"></a>return nil
            <a id="L1303"></a>}
            <a id="L1304"></a>capf = capexpr.asInt();
        <a id="L1305"></a>}

        <a id="L1307"></a>switch t := as[0].valType.lit().(type) {
        <a id="L1308"></a>case *SliceType:
            <a id="L1309"></a><span class="comment">// A new, initialized slice value for a given</span>
            <a id="L1310"></a><span class="comment">// element type T is made using the built-in</span>
            <a id="L1311"></a><span class="comment">// function make, which takes a slice type and</span>
            <a id="L1312"></a><span class="comment">// parameters specifying the length and</span>
            <a id="L1313"></a><span class="comment">// optionally the capacity.</span>
            <a id="L1314"></a>if !checkCount(2, 3) {
                <a id="L1315"></a>return nil
            <a id="L1316"></a>}
            <a id="L1317"></a>et := t.Elem;
            <a id="L1318"></a>expr := a.newExpr(t, &#34;function call&#34;);
            <a id="L1319"></a>expr.eval = func(t *Thread) Slice {
                <a id="L1320"></a>l := lenf(t);
                <a id="L1321"></a><span class="comment">// XXX(Spec) What if len or cap is</span>
                <a id="L1322"></a><span class="comment">// negative?  The runtime panics.</span>
                <a id="L1323"></a>if l &lt; 0 {
                    <a id="L1324"></a>t.Abort(NegativeLengthError{l})
                <a id="L1325"></a>}
                <a id="L1326"></a>c := l;
                <a id="L1327"></a>if capf != nil {
                    <a id="L1328"></a>c = capf(t);
                    <a id="L1329"></a>if c &lt; 0 {
                        <a id="L1330"></a>t.Abort(NegativeCapacityError{c})
                    <a id="L1331"></a>}
                    <a id="L1332"></a><span class="comment">// XXX(Spec) What happens if</span>
                    <a id="L1333"></a><span class="comment">// len &gt; cap?  The runtime</span>
                    <a id="L1334"></a><span class="comment">// sets cap to len.</span>
                    <a id="L1335"></a>if l &gt; c {
                        <a id="L1336"></a>c = l
                    <a id="L1337"></a>}
                <a id="L1338"></a>}
                <a id="L1339"></a>base := arrayV(make([]Value, c));
                <a id="L1340"></a>for i := int64(0); i &lt; c; i++ {
                    <a id="L1341"></a>base[i] = et.Zero()
                <a id="L1342"></a>}
                <a id="L1343"></a>return Slice{&amp;base, l, c};
            <a id="L1344"></a>};
            <a id="L1345"></a>return expr;

        <a id="L1347"></a>case *MapType:
            <a id="L1348"></a><span class="comment">// A new, empty map value is made using the</span>
            <a id="L1349"></a><span class="comment">// built-in function make, which takes the map</span>
            <a id="L1350"></a><span class="comment">// type and an optional capacity hint as</span>
            <a id="L1351"></a><span class="comment">// arguments.</span>
            <a id="L1352"></a>if !checkCount(1, 2) {
                <a id="L1353"></a>return nil
            <a id="L1354"></a>}
            <a id="L1355"></a>expr := a.newExpr(t, &#34;function call&#34;);
            <a id="L1356"></a>expr.eval = func(t *Thread) Map {
                <a id="L1357"></a>if lenf == nil {
                    <a id="L1358"></a>return make(evalMap)
                <a id="L1359"></a>}
                <a id="L1360"></a>l := lenf(t);
                <a id="L1361"></a>return make(evalMap, l);
            <a id="L1362"></a>};
            <a id="L1363"></a>return expr;

        <a id="L1365"></a><span class="comment">//case *ChanType:</span>

        <a id="L1367"></a>default:
            <a id="L1368"></a>a.diag(&#34;illegal argument type for make function\n\t%v&#34;, as[0].valType);
            <a id="L1369"></a>return nil;
        <a id="L1370"></a>}

    <a id="L1372"></a>case closeType, closedType:
        <a id="L1373"></a>a.diag(&#34;built-in function %s not implemented&#34;, ft.builtin);
        <a id="L1374"></a>return nil;

    <a id="L1376"></a>case newType:
        <a id="L1377"></a>if !checkCount(1, 1) {
            <a id="L1378"></a>return nil
        <a id="L1379"></a>}

        <a id="L1381"></a>t := as[0].valType;
        <a id="L1382"></a>expr := a.newExpr(NewPtrType(t), &#34;new&#34;);
        <a id="L1383"></a>expr.eval = func(*Thread) Value { return t.Zero() };
        <a id="L1384"></a>return expr;

    <a id="L1386"></a>case panicType, paniclnType, printType, printlnType:
        <a id="L1387"></a>evals := make([]func(*Thread) interface{}, len(as));
        <a id="L1388"></a>for i, x := range as {
            <a id="L1389"></a>evals[i] = x.asInterface()
        <a id="L1390"></a>}
        <a id="L1391"></a>spaces := ft == paniclnType || ft == printlnType;
        <a id="L1392"></a>newline := ft != printType;
        <a id="L1393"></a>printer := func(t *Thread) {
            <a id="L1394"></a>for i, eval := range evals {
                <a id="L1395"></a>if i &gt; 0 &amp;&amp; spaces {
                    <a id="L1396"></a>print(&#34; &#34;)
                <a id="L1397"></a>}
                <a id="L1398"></a>v := eval(t);
                <a id="L1399"></a>type stringer interface {
                    <a id="L1400"></a>String() string;
                <a id="L1401"></a>}
                <a id="L1402"></a>switch v1 := v.(type) {
                <a id="L1403"></a>case bool:
                    <a id="L1404"></a>print(v1)
                <a id="L1405"></a>case uint64:
                    <a id="L1406"></a>print(v1)
                <a id="L1407"></a>case int64:
                    <a id="L1408"></a>print(v1)
                <a id="L1409"></a>case float64:
                    <a id="L1410"></a>print(v1)
                <a id="L1411"></a>case string:
                    <a id="L1412"></a>print(v1)
                <a id="L1413"></a>case stringer:
                    <a id="L1414"></a>print(v1.String())
                <a id="L1415"></a>default:
                    <a id="L1416"></a>print(&#34;???&#34;)
                <a id="L1417"></a>}
            <a id="L1418"></a>}
            <a id="L1419"></a>if newline {
                <a id="L1420"></a>print(&#34;\n&#34;)
            <a id="L1421"></a>}
        <a id="L1422"></a>};
        <a id="L1423"></a>expr := a.newExpr(EmptyType, &#34;print&#34;);
        <a id="L1424"></a>expr.exec = printer;
        <a id="L1425"></a>if ft == panicType || ft == paniclnType {
            <a id="L1426"></a>expr.exec = func(t *Thread) {
                <a id="L1427"></a>printer(t);
                <a id="L1428"></a>t.Abort(os.NewError(&#34;panic&#34;));
            <a id="L1429"></a>}
        <a id="L1430"></a>}
        <a id="L1431"></a>return expr;
    <a id="L1432"></a>}

    <a id="L1434"></a>log.Crashf(&#34;unexpected built-in function &#39;%s&#39;&#34;, ft.builtin);
    <a id="L1435"></a>panic();
<a id="L1436"></a>}

<a id="L1438"></a>func (a *exprInfo) compileStarExpr(v *expr) *expr {
    <a id="L1439"></a>switch vt := v.t.lit().(type) {
    <a id="L1440"></a>case *PtrType:
        <a id="L1441"></a>expr := a.newExpr(vt.Elem, &#34;indirect expression&#34;);
        <a id="L1442"></a>vf := v.asPtr();
        <a id="L1443"></a>expr.genValue(func(t *Thread) Value {
            <a id="L1444"></a>v := vf(t);
            <a id="L1445"></a>if v == nil {
                <a id="L1446"></a>t.Abort(NilPointerError{})
            <a id="L1447"></a>}
            <a id="L1448"></a>return v;
        <a id="L1449"></a>});
        <a id="L1450"></a>return expr;
    <a id="L1451"></a>}

    <a id="L1453"></a>a.diagOpType(token.MUL, v.t);
    <a id="L1454"></a>return nil;
<a id="L1455"></a>}

<a id="L1457"></a>var unaryOpDescs = make(map[token.Token]string)

<a id="L1459"></a>func (a *exprInfo) compileUnaryExpr(op token.Token, v *expr) *expr {
    <a id="L1460"></a><span class="comment">// Type check</span>
    <a id="L1461"></a>var t Type;
    <a id="L1462"></a>switch op {
    <a id="L1463"></a>case token.ADD, token.SUB:
        <a id="L1464"></a>if !v.t.isInteger() &amp;&amp; !v.t.isFloat() {
            <a id="L1465"></a>a.diagOpType(op, v.t);
            <a id="L1466"></a>return nil;
        <a id="L1467"></a>}
        <a id="L1468"></a>t = v.t;

    <a id="L1470"></a>case token.NOT:
        <a id="L1471"></a>if !v.t.isBoolean() {
            <a id="L1472"></a>a.diagOpType(op, v.t);
            <a id="L1473"></a>return nil;
        <a id="L1474"></a>}
        <a id="L1475"></a>t = BoolType;

    <a id="L1477"></a>case token.XOR:
        <a id="L1478"></a>if !v.t.isInteger() {
            <a id="L1479"></a>a.diagOpType(op, v.t);
            <a id="L1480"></a>return nil;
        <a id="L1481"></a>}
        <a id="L1482"></a>t = v.t;

    <a id="L1484"></a>case token.AND:
        <a id="L1485"></a><span class="comment">// The unary prefix address-of operator &amp; generates</span>
        <a id="L1486"></a><span class="comment">// the address of its operand, which must be a</span>
        <a id="L1487"></a><span class="comment">// variable, pointer indirection, field selector, or</span>
        <a id="L1488"></a><span class="comment">// array or slice indexing operation.</span>
        <a id="L1489"></a>if v.evalAddr == nil {
            <a id="L1490"></a>a.diag(&#34;cannot take the address of %s&#34;, v.desc);
            <a id="L1491"></a>return nil;
        <a id="L1492"></a>}

        <a id="L1494"></a><span class="comment">// TODO(austin) Implement &#34;It is illegal to take the</span>
        <a id="L1495"></a><span class="comment">// address of a function result variable&#34; once I have</span>
        <a id="L1496"></a><span class="comment">// function result variables.</span>

        <a id="L1498"></a>t = NewPtrType(v.t);

    <a id="L1500"></a>case token.ARROW:
        <a id="L1501"></a>log.Crashf(&#34;Unary op %v not implemented&#34;, op)

    <a id="L1503"></a>default:
        <a id="L1504"></a>log.Crashf(&#34;unknown unary operator %v&#34;, op)
    <a id="L1505"></a>}

    <a id="L1507"></a>desc, ok := unaryOpDescs[op];
    <a id="L1508"></a>if !ok {
        <a id="L1509"></a>desc = &#34;unary &#34; + op.String() + &#34; expression&#34;;
        <a id="L1510"></a>unaryOpDescs[op] = desc;
    <a id="L1511"></a>}

    <a id="L1513"></a><span class="comment">// Compile</span>
    <a id="L1514"></a>expr := a.newExpr(t, desc);
    <a id="L1515"></a>switch op {
    <a id="L1516"></a>case token.ADD:
        <a id="L1517"></a><span class="comment">// Just compile it out</span>
        <a id="L1518"></a>expr = v;
        <a id="L1519"></a>expr.desc = desc;

    <a id="L1521"></a>case token.SUB:
        <a id="L1522"></a>expr.genUnaryOpNeg(v)

    <a id="L1524"></a>case token.NOT:
        <a id="L1525"></a>expr.genUnaryOpNot(v)

    <a id="L1527"></a>case token.XOR:
        <a id="L1528"></a>expr.genUnaryOpXor(v)

    <a id="L1530"></a>case token.AND:
        <a id="L1531"></a>vf := v.evalAddr;
        <a id="L1532"></a>expr.eval = func(t *Thread) Value { return vf(t) };

    <a id="L1534"></a>default:
        <a id="L1535"></a>log.Crashf(&#34;Compilation of unary op %v not implemented&#34;, op)
    <a id="L1536"></a>}

    <a id="L1538"></a>return expr;
<a id="L1539"></a>}

<a id="L1541"></a>var binOpDescs = make(map[token.Token]string)

<a id="L1543"></a>func (a *exprInfo) compileBinaryExpr(op token.Token, l, r *expr) *expr {
    <a id="L1544"></a><span class="comment">// Save the original types of l.t and r.t for error messages.</span>
    <a id="L1545"></a>origlt := l.t;
    <a id="L1546"></a>origrt := r.t;

    <a id="L1548"></a><span class="comment">// XXX(Spec) What is the exact definition of a &#34;named type&#34;?</span>

    <a id="L1550"></a><span class="comment">// XXX(Spec) Arithmetic operators: &#34;Integer types&#34; apparently</span>
    <a id="L1551"></a><span class="comment">// means all types compatible with basic integer types, though</span>
    <a id="L1552"></a><span class="comment">// this is never explained.  Likewise for float types, etc.</span>
    <a id="L1553"></a><span class="comment">// This relates to the missing explanation of named types.</span>

    <a id="L1555"></a><span class="comment">// XXX(Spec) Operators: &#34;If both operands are ideal numbers,</span>
    <a id="L1556"></a><span class="comment">// the conversion is to ideal floats if one of the operands is</span>
    <a id="L1557"></a><span class="comment">// an ideal float (relevant for / and %).&#34;  How is that</span>
    <a id="L1558"></a><span class="comment">// relevant only for / and %?  If I add an ideal int and an</span>
    <a id="L1559"></a><span class="comment">// ideal float, I get an ideal float.</span>

    <a id="L1561"></a>if op != token.SHL &amp;&amp; op != token.SHR {
        <a id="L1562"></a><span class="comment">// Except in shift expressions, if one operand has</span>
        <a id="L1563"></a><span class="comment">// numeric type and the other operand is an ideal</span>
        <a id="L1564"></a><span class="comment">// number, the ideal number is converted to match the</span>
        <a id="L1565"></a><span class="comment">// type of the other operand.</span>
        <a id="L1566"></a>if (l.t.isInteger() || l.t.isFloat()) &amp;&amp; !l.t.isIdeal() &amp;&amp; r.t.isIdeal() {
            <a id="L1567"></a>r = r.convertTo(l.t)
        <a id="L1568"></a>} else if (r.t.isInteger() || r.t.isFloat()) &amp;&amp; !r.t.isIdeal() &amp;&amp; l.t.isIdeal() {
            <a id="L1569"></a>l = l.convertTo(r.t)
        <a id="L1570"></a>}
        <a id="L1571"></a>if l == nil || r == nil {
            <a id="L1572"></a>return nil
        <a id="L1573"></a>}

        <a id="L1575"></a><span class="comment">// Except in shift expressions, if both operands are</span>
        <a id="L1576"></a><span class="comment">// ideal numbers and one is an ideal float, the other</span>
        <a id="L1577"></a><span class="comment">// is converted to ideal float.</span>
        <a id="L1578"></a>if l.t.isIdeal() &amp;&amp; r.t.isIdeal() {
            <a id="L1579"></a>if l.t.isInteger() &amp;&amp; r.t.isFloat() {
                <a id="L1580"></a>l = l.convertTo(r.t)
            <a id="L1581"></a>} else if l.t.isFloat() &amp;&amp; r.t.isInteger() {
                <a id="L1582"></a>r = r.convertTo(l.t)
            <a id="L1583"></a>}
            <a id="L1584"></a>if l == nil || r == nil {
                <a id="L1585"></a>return nil
            <a id="L1586"></a>}
        <a id="L1587"></a>}
    <a id="L1588"></a>}

    <a id="L1590"></a><span class="comment">// Useful type predicates</span>
    <a id="L1591"></a><span class="comment">// TODO(austin) CL 33668 mandates identical types except for comparisons.</span>
    <a id="L1592"></a>compat := func() bool { return l.t.compat(r.t, false) };
    <a id="L1593"></a>integers := func() bool { return l.t.isInteger() &amp;&amp; r.t.isInteger() };
    <a id="L1594"></a>floats := func() bool { return l.t.isFloat() &amp;&amp; r.t.isFloat() };
    <a id="L1595"></a>strings := func() bool {
        <a id="L1596"></a><span class="comment">// TODO(austin) Deal with named types</span>
        <a id="L1597"></a>return l.t == StringType &amp;&amp; r.t == StringType
    <a id="L1598"></a>};
    <a id="L1599"></a>booleans := func() bool { return l.t.isBoolean() &amp;&amp; r.t.isBoolean() };

    <a id="L1601"></a><span class="comment">// Type check</span>
    <a id="L1602"></a>var t Type;
    <a id="L1603"></a>switch op {
    <a id="L1604"></a>case token.ADD:
        <a id="L1605"></a>if !compat() || (!integers() &amp;&amp; !floats() &amp;&amp; !strings()) {
            <a id="L1606"></a>a.diagOpTypes(op, origlt, origrt);
            <a id="L1607"></a>return nil;
        <a id="L1608"></a>}
        <a id="L1609"></a>t = l.t;

    <a id="L1611"></a>case token.SUB, token.MUL, token.QUO:
        <a id="L1612"></a>if !compat() || (!integers() &amp;&amp; !floats()) {
            <a id="L1613"></a>a.diagOpTypes(op, origlt, origrt);
            <a id="L1614"></a>return nil;
        <a id="L1615"></a>}
        <a id="L1616"></a>t = l.t;

    <a id="L1618"></a>case token.REM, token.AND, token.OR, token.XOR, token.AND_NOT:
        <a id="L1619"></a>if !compat() || !integers() {
            <a id="L1620"></a>a.diagOpTypes(op, origlt, origrt);
            <a id="L1621"></a>return nil;
        <a id="L1622"></a>}
        <a id="L1623"></a>t = l.t;

    <a id="L1625"></a>case token.SHL, token.SHR:
        <a id="L1626"></a><span class="comment">// XXX(Spec) Is it okay for the right operand to be an</span>
        <a id="L1627"></a><span class="comment">// ideal float with no fractional part?  &#34;The right</span>
        <a id="L1628"></a><span class="comment">// operand in a shift operation must be always be of</span>
        <a id="L1629"></a><span class="comment">// unsigned integer type or an ideal number that can</span>
        <a id="L1630"></a><span class="comment">// be safely converted into an unsigned integer type</span>
        <a id="L1631"></a><span class="comment">// (§Arithmetic operators)&#34; suggests so and 6g agrees.</span>

        <a id="L1633"></a>if !l.t.isInteger() || !(r.t.isInteger() || r.t.isIdeal()) {
            <a id="L1634"></a>a.diagOpTypes(op, origlt, origrt);
            <a id="L1635"></a>return nil;
        <a id="L1636"></a>}

        <a id="L1638"></a><span class="comment">// The right operand in a shift operation must be</span>
        <a id="L1639"></a><span class="comment">// always be of unsigned integer type or an ideal</span>
        <a id="L1640"></a><span class="comment">// number that can be safely converted into an</span>
        <a id="L1641"></a><span class="comment">// unsigned integer type.</span>
        <a id="L1642"></a>if r.t.isIdeal() {
            <a id="L1643"></a>r2 := r.convertTo(UintType);
            <a id="L1644"></a>if r2 == nil {
                <a id="L1645"></a>return nil
            <a id="L1646"></a>}

            <a id="L1648"></a><span class="comment">// If the left operand is not ideal, convert</span>
            <a id="L1649"></a><span class="comment">// the right to not ideal.</span>
            <a id="L1650"></a>if !l.t.isIdeal() {
                <a id="L1651"></a>r = r2
            <a id="L1652"></a>}

            <a id="L1654"></a><span class="comment">// If both are ideal, but the right side isn&#39;t</span>
            <a id="L1655"></a><span class="comment">// an ideal int, convert it to simplify things.</span>
            <a id="L1656"></a>if l.t.isIdeal() &amp;&amp; !r.t.isInteger() {
                <a id="L1657"></a>r = r.convertTo(IdealIntType);
                <a id="L1658"></a>if r == nil {
                    <a id="L1659"></a>log.Crashf(&#34;conversion to uintType succeeded, but conversion to idealIntType failed&#34;)
                <a id="L1660"></a>}
            <a id="L1661"></a>}
        <a id="L1662"></a>} else if _, ok := r.t.lit().(*uintType); !ok {
            <a id="L1663"></a>a.diag(&#34;right operand of shift must be unsigned&#34;);
            <a id="L1664"></a>return nil;
        <a id="L1665"></a>}

        <a id="L1667"></a>if l.t.isIdeal() &amp;&amp; !r.t.isIdeal() {
            <a id="L1668"></a><span class="comment">// XXX(Spec) What is the meaning of &#34;ideal &gt;&gt;</span>
            <a id="L1669"></a><span class="comment">// non-ideal&#34;?  Russ says the ideal should be</span>
            <a id="L1670"></a><span class="comment">// converted to an int.  6g propagates the</span>
            <a id="L1671"></a><span class="comment">// type down from assignments as a hint.</span>

            <a id="L1673"></a>l = l.convertTo(IntType);
            <a id="L1674"></a>if l == nil {
                <a id="L1675"></a>return nil
            <a id="L1676"></a>}
        <a id="L1677"></a>}

        <a id="L1679"></a><span class="comment">// At this point, we should have one of three cases:</span>
        <a id="L1680"></a><span class="comment">// 1) uint SHIFT uint</span>
        <a id="L1681"></a><span class="comment">// 2) int SHIFT uint</span>
        <a id="L1682"></a><span class="comment">// 3) ideal int SHIFT ideal int</span>

        <a id="L1684"></a>t = l.t;

    <a id="L1686"></a>case token.LOR, token.LAND:
        <a id="L1687"></a>if !booleans() {
            <a id="L1688"></a>return nil
        <a id="L1689"></a>}
        <a id="L1690"></a><span class="comment">// XXX(Spec) There&#39;s no mention of *which* boolean</span>
        <a id="L1691"></a><span class="comment">// type the logical operators return.  From poking at</span>
        <a id="L1692"></a><span class="comment">// 6g, it appears to be the named boolean type, NOT</span>
        <a id="L1693"></a><span class="comment">// the type of the left operand, and NOT an unnamed</span>
        <a id="L1694"></a><span class="comment">// boolean type.</span>

        <a id="L1696"></a>t = BoolType;

    <a id="L1698"></a>case token.ARROW:
        <a id="L1699"></a><span class="comment">// The operands in channel sends differ in type: one</span>
        <a id="L1700"></a><span class="comment">// is always a channel and the other is a variable or</span>
        <a id="L1701"></a><span class="comment">// value of the channel&#39;s element type.</span>
        <a id="L1702"></a>log.Crash(&#34;Binary op &lt;- not implemented&#34;);
        <a id="L1703"></a>t = BoolType;

    <a id="L1705"></a>case token.LSS, token.GTR, token.LEQ, token.GEQ:
        <a id="L1706"></a><span class="comment">// XXX(Spec) It&#39;s really unclear what types which</span>
        <a id="L1707"></a><span class="comment">// comparison operators apply to.  I feel like the</span>
        <a id="L1708"></a><span class="comment">// text is trying to paint a Venn diagram for me,</span>
        <a id="L1709"></a><span class="comment">// which it&#39;s really pretty simple: &lt;, &lt;=, &gt;, &gt;= apply</span>
        <a id="L1710"></a><span class="comment">// only to numeric types and strings.  == and != apply</span>
        <a id="L1711"></a><span class="comment">// to everything except arrays and structs, and there</span>
        <a id="L1712"></a><span class="comment">// are some restrictions on when it applies to slices.</span>

        <a id="L1714"></a>if !compat() || (!integers() &amp;&amp; !floats() &amp;&amp; !strings()) {
            <a id="L1715"></a>a.diagOpTypes(op, origlt, origrt);
            <a id="L1716"></a>return nil;
        <a id="L1717"></a>}
        <a id="L1718"></a>t = BoolType;

    <a id="L1720"></a>case token.EQL, token.NEQ:
        <a id="L1721"></a><span class="comment">// XXX(Spec) The rules for type checking comparison</span>
        <a id="L1722"></a><span class="comment">// operators are spread across three places that all</span>
        <a id="L1723"></a><span class="comment">// partially overlap with each other: the Comparison</span>
        <a id="L1724"></a><span class="comment">// Compatibility section, the Operators section, and</span>
        <a id="L1725"></a><span class="comment">// the Comparison Operators section.  The Operators</span>
        <a id="L1726"></a><span class="comment">// section should just say that operators require</span>
        <a id="L1727"></a><span class="comment">// identical types (as it does currently) except that</span>
        <a id="L1728"></a><span class="comment">// there a few special cases for comparison, which are</span>
        <a id="L1729"></a><span class="comment">// described in section X.  Currently it includes just</span>
        <a id="L1730"></a><span class="comment">// one of the four special cases.  The Comparison</span>
        <a id="L1731"></a><span class="comment">// Compatibility section and the Comparison Operators</span>
        <a id="L1732"></a><span class="comment">// section should either be merged, or at least the</span>
        <a id="L1733"></a><span class="comment">// Comparison Compatibility section should be</span>
        <a id="L1734"></a><span class="comment">// exclusively about type checking and the Comparison</span>
        <a id="L1735"></a><span class="comment">// Operators section should be exclusively about</span>
        <a id="L1736"></a><span class="comment">// semantics.</span>

        <a id="L1738"></a><span class="comment">// XXX(Spec) Comparison operators: &#34;All comparison</span>
        <a id="L1739"></a><span class="comment">// operators apply to basic types except bools.&#34;  This</span>
        <a id="L1740"></a><span class="comment">// is very difficult to parse.  It&#39;s explained much</span>
        <a id="L1741"></a><span class="comment">// better in the Comparison Compatibility section.</span>

        <a id="L1743"></a><span class="comment">// XXX(Spec) Comparison compatibility: &#34;Function</span>
        <a id="L1744"></a><span class="comment">// values are equal if they refer to the same</span>
        <a id="L1745"></a><span class="comment">// function.&#34; is rather vague.  It should probably be</span>
        <a id="L1746"></a><span class="comment">// similar to the way the rule for map values is</span>
        <a id="L1747"></a><span class="comment">// written: Function values are equal if they were</span>
        <a id="L1748"></a><span class="comment">// created by the same execution of a function literal</span>
        <a id="L1749"></a><span class="comment">// or refer to the same function declaration.  This is</span>
        <a id="L1750"></a><span class="comment">// *almost* but not quite waht 6g implements.  If a</span>
        <a id="L1751"></a><span class="comment">// function literals does not capture any variables,</span>
        <a id="L1752"></a><span class="comment">// then multiple executions of it will result in the</span>
        <a id="L1753"></a><span class="comment">// same closure.  Russ says he&#39;ll change that.</span>

        <a id="L1755"></a><span class="comment">// TODO(austin) Deal with remaining special cases</span>

        <a id="L1757"></a>if !compat() {
            <a id="L1758"></a>a.diagOpTypes(op, origlt, origrt);
            <a id="L1759"></a>return nil;
        <a id="L1760"></a>}
        <a id="L1761"></a><span class="comment">// Arrays and structs may not be compared to anything.</span>
        <a id="L1762"></a>switch l.t.(type) {
        <a id="L1763"></a>case *ArrayType, *StructType:
            <a id="L1764"></a>a.diagOpTypes(op, origlt, origrt);
            <a id="L1765"></a>return nil;
        <a id="L1766"></a>}
        <a id="L1767"></a>t = BoolType;

    <a id="L1769"></a>default:
        <a id="L1770"></a>log.Crashf(&#34;unknown binary operator %v&#34;, op)
    <a id="L1771"></a>}

    <a id="L1773"></a>desc, ok := binOpDescs[op];
    <a id="L1774"></a>if !ok {
        <a id="L1775"></a>desc = op.String() + &#34; expression&#34;;
        <a id="L1776"></a>binOpDescs[op] = desc;
    <a id="L1777"></a>}

    <a id="L1779"></a><span class="comment">// Check for ideal divide by zero</span>
    <a id="L1780"></a>switch op {
    <a id="L1781"></a>case token.QUO, token.REM:
        <a id="L1782"></a>if r.t.isIdeal() {
            <a id="L1783"></a>if (r.t.isInteger() &amp;&amp; r.asIdealInt()().IsZero()) ||
                <a id="L1784"></a>(r.t.isFloat() &amp;&amp; r.asIdealFloat()().IsZero()) {
                <a id="L1785"></a>a.diag(&#34;divide by zero&#34;);
                <a id="L1786"></a>return nil;
            <a id="L1787"></a>}
        <a id="L1788"></a>}
    <a id="L1789"></a>}

    <a id="L1791"></a><span class="comment">// Compile</span>
    <a id="L1792"></a>expr := a.newExpr(t, desc);
    <a id="L1793"></a>switch op {
    <a id="L1794"></a>case token.ADD:
        <a id="L1795"></a>expr.genBinOpAdd(l, r)

    <a id="L1797"></a>case token.SUB:
        <a id="L1798"></a>expr.genBinOpSub(l, r)

    <a id="L1800"></a>case token.MUL:
        <a id="L1801"></a>expr.genBinOpMul(l, r)

    <a id="L1803"></a>case token.QUO:
        <a id="L1804"></a>expr.genBinOpQuo(l, r)

    <a id="L1806"></a>case token.REM:
        <a id="L1807"></a>expr.genBinOpRem(l, r)

    <a id="L1809"></a>case token.AND:
        <a id="L1810"></a>expr.genBinOpAnd(l, r)

    <a id="L1812"></a>case token.OR:
        <a id="L1813"></a>expr.genBinOpOr(l, r)

    <a id="L1815"></a>case token.XOR:
        <a id="L1816"></a>expr.genBinOpXor(l, r)

    <a id="L1818"></a>case token.AND_NOT:
        <a id="L1819"></a>expr.genBinOpAndNot(l, r)

    <a id="L1821"></a>case token.SHL:
        <a id="L1822"></a>if l.t.isIdeal() {
            <a id="L1823"></a>lv := l.asIdealInt()();
            <a id="L1824"></a>rv := r.asIdealInt()();
            <a id="L1825"></a>const maxShift = 99999;
            <a id="L1826"></a>if rv.Cmp(bignum.Int(maxShift)) &gt; 0 {
                <a id="L1827"></a>a.diag(&#34;left shift by %v; exceeds implementation limit of %v&#34;, rv, maxShift);
                <a id="L1828"></a>expr.t = nil;
                <a id="L1829"></a>return nil;
            <a id="L1830"></a>}
            <a id="L1831"></a>val := lv.Shl(uint(rv.Value()));
            <a id="L1832"></a>expr.eval = func() *bignum.Integer { return val };
        <a id="L1833"></a>} else {
            <a id="L1834"></a>expr.genBinOpShl(l, r)
        <a id="L1835"></a>}

    <a id="L1837"></a>case token.SHR:
        <a id="L1838"></a>if l.t.isIdeal() {
            <a id="L1839"></a>lv := l.asIdealInt()();
            <a id="L1840"></a>rv := r.asIdealInt()();
            <a id="L1841"></a>val := lv.Shr(uint(rv.Value()));
            <a id="L1842"></a>expr.eval = func() *bignum.Integer { return val };
        <a id="L1843"></a>} else {
            <a id="L1844"></a>expr.genBinOpShr(l, r)
        <a id="L1845"></a>}

    <a id="L1847"></a>case token.LSS:
        <a id="L1848"></a>expr.genBinOpLss(l, r)

    <a id="L1850"></a>case token.GTR:
        <a id="L1851"></a>expr.genBinOpGtr(l, r)

    <a id="L1853"></a>case token.LEQ:
        <a id="L1854"></a>expr.genBinOpLeq(l, r)

    <a id="L1856"></a>case token.GEQ:
        <a id="L1857"></a>expr.genBinOpGeq(l, r)

    <a id="L1859"></a>case token.EQL:
        <a id="L1860"></a>expr.genBinOpEql(l, r)

    <a id="L1862"></a>case token.NEQ:
        <a id="L1863"></a>expr.genBinOpNeq(l, r)

    <a id="L1865"></a>case token.LAND:
        <a id="L1866"></a>expr.genBinOpLogAnd(l, r)

    <a id="L1868"></a>case token.LOR:
        <a id="L1869"></a>expr.genBinOpLogOr(l, r)

    <a id="L1871"></a>default:
        <a id="L1872"></a>log.Crashf(&#34;Compilation of binary op %v not implemented&#34;, op)
    <a id="L1873"></a>}

    <a id="L1875"></a>return expr;
<a id="L1876"></a>}

<a id="L1878"></a><span class="comment">// TODO(austin) This is a hack to eliminate a circular dependency</span>
<a id="L1879"></a><span class="comment">// between type.go and expr.go</span>
<a id="L1880"></a>func (a *compiler) compileArrayLen(b *block, expr ast.Expr) (int64, bool) {
    <a id="L1881"></a>lenExpr := a.compileExpr(b, true, expr);
    <a id="L1882"></a>if lenExpr == nil {
        <a id="L1883"></a>return 0, false
    <a id="L1884"></a>}

    <a id="L1886"></a><span class="comment">// XXX(Spec) Are ideal floats with no fractional part okay?</span>
    <a id="L1887"></a>if lenExpr.t.isIdeal() {
        <a id="L1888"></a>lenExpr = lenExpr.convertTo(IntType);
        <a id="L1889"></a>if lenExpr == nil {
            <a id="L1890"></a>return 0, false
        <a id="L1891"></a>}
    <a id="L1892"></a>}

    <a id="L1894"></a>if !lenExpr.t.isInteger() {
        <a id="L1895"></a>a.diagAt(expr, &#34;array size must be an integer&#34;);
        <a id="L1896"></a>return 0, false;
    <a id="L1897"></a>}

    <a id="L1899"></a>switch lenExpr.t.lit().(type) {
    <a id="L1900"></a>case *intType:
        <a id="L1901"></a>return lenExpr.asInt()(nil), true
    <a id="L1902"></a>case *uintType:
        <a id="L1903"></a>return int64(lenExpr.asUint()(nil)), true
    <a id="L1904"></a>}
    <a id="L1905"></a>log.Crashf(&#34;unexpected integer type %T&#34;, lenExpr.t);
    <a id="L1906"></a>return 0, false;
<a id="L1907"></a>}

<a id="L1909"></a>func (a *compiler) compileExpr(b *block, constant bool, expr ast.Expr) *expr {
    <a id="L1910"></a>ec := &amp;exprCompiler{a, b, constant};
    <a id="L1911"></a>nerr := a.numError();
    <a id="L1912"></a>e := ec.compile(expr, false);
    <a id="L1913"></a>if e == nil &amp;&amp; nerr == a.numError() {
        <a id="L1914"></a>log.Crashf(&#34;expression compilation failed without reporting errors&#34;)
    <a id="L1915"></a>}
    <a id="L1916"></a>return e;
<a id="L1917"></a>}

<a id="L1919"></a><span class="comment">// extractEffect separates out any effects that the expression may</span>
<a id="L1920"></a><span class="comment">// have, returning a function that will perform those effects and a</span>
<a id="L1921"></a><span class="comment">// new exprCompiler that is guaranteed to be side-effect free.  These</span>
<a id="L1922"></a><span class="comment">// are the moral equivalents of &#34;temp := expr&#34; and &#34;temp&#34; (or &#34;temp :=</span>
<a id="L1923"></a><span class="comment">// &amp;expr&#34; and &#34;*temp&#34; for addressable exprs).  Because this creates a</span>
<a id="L1924"></a><span class="comment">// temporary variable, the caller should create a temporary block for</span>
<a id="L1925"></a><span class="comment">// the compilation of this expression and the evaluation of the</span>
<a id="L1926"></a><span class="comment">// results.</span>
<a id="L1927"></a>func (a *expr) extractEffect(b *block, errOp string) (func(*Thread), *expr) {
    <a id="L1928"></a><span class="comment">// Create &#34;&amp;a&#34; if a is addressable</span>
    <a id="L1929"></a>rhs := a;
    <a id="L1930"></a>if a.evalAddr != nil {
        <a id="L1931"></a>rhs = a.compileUnaryExpr(token.AND, rhs)
    <a id="L1932"></a>}

    <a id="L1934"></a><span class="comment">// Create temp</span>
    <a id="L1935"></a>ac, ok := a.checkAssign(a.pos, []*expr{rhs}, errOp, &#34;&#34;);
    <a id="L1936"></a>if !ok {
        <a id="L1937"></a>return nil, nil
    <a id="L1938"></a>}
    <a id="L1939"></a>if len(ac.rmt.Elems) != 1 {
        <a id="L1940"></a>a.diag(&#34;multi-valued expression not allowed in %s&#34;, errOp);
        <a id="L1941"></a>return nil, nil;
    <a id="L1942"></a>}
    <a id="L1943"></a>tempType := ac.rmt.Elems[0];
    <a id="L1944"></a>if tempType.isIdeal() {
        <a id="L1945"></a><span class="comment">// It&#39;s too bad we have to duplicate this rule.</span>
        <a id="L1946"></a>switch {
        <a id="L1947"></a>case tempType.isInteger():
            <a id="L1948"></a>tempType = IntType
        <a id="L1949"></a>case tempType.isFloat():
            <a id="L1950"></a>tempType = FloatType
        <a id="L1951"></a>default:
            <a id="L1952"></a>log.Crashf(&#34;unexpected ideal type %v&#34;, tempType)
        <a id="L1953"></a>}
    <a id="L1954"></a>}
    <a id="L1955"></a>temp := b.DefineTemp(tempType);
    <a id="L1956"></a>tempIdx := temp.Index;

    <a id="L1958"></a><span class="comment">// Create &#34;temp := rhs&#34;</span>
    <a id="L1959"></a>assign := ac.compile(b, tempType);
    <a id="L1960"></a>if assign == nil {
        <a id="L1961"></a>log.Crashf(&#34;compileAssign type check failed&#34;)
    <a id="L1962"></a>}

    <a id="L1964"></a>effect := func(t *Thread) {
        <a id="L1965"></a>tempVal := tempType.Zero();
        <a id="L1966"></a>t.f.Vars[tempIdx] = tempVal;
        <a id="L1967"></a>assign(tempVal, t);
    <a id="L1968"></a>};

    <a id="L1970"></a><span class="comment">// Generate &#34;temp&#34; or &#34;*temp&#34;</span>
    <a id="L1971"></a>getTemp := a.compileVariable(0, temp);
    <a id="L1972"></a>if a.evalAddr == nil {
        <a id="L1973"></a>return effect, getTemp
    <a id="L1974"></a>}

    <a id="L1976"></a>deref := a.compileStarExpr(getTemp);
    <a id="L1977"></a>if deref == nil {
        <a id="L1978"></a>return nil, nil
    <a id="L1979"></a>}
    <a id="L1980"></a>return effect, deref;
<a id="L1981"></a>}
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
