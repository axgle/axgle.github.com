<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/gen.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/gen.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package main

<a id="L7"></a><span class="comment">// generate operator implementations</span>

<a id="L9"></a>import (
    <a id="L10"></a>&#34;log&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;template&#34;;
<a id="L13"></a>)

<a id="L15"></a>type Op struct {
    <a id="L16"></a>Name        string;
    <a id="L17"></a>Expr        string;
    <a id="L18"></a>Body        string; <span class="comment">// overrides Expr</span>
    <a id="L19"></a>ConstExpr   string;
    <a id="L20"></a>AsRightName string;
    <a id="L21"></a>ReturnType  string;
    <a id="L22"></a>Types       []*Type;
<a id="L23"></a>}

<a id="L25"></a>type Size struct {
    <a id="L26"></a>Bits  int;
    <a id="L27"></a>Sized string;
<a id="L28"></a>}

<a id="L30"></a>type Type struct {
    <a id="L31"></a>Repr      string;
    <a id="L32"></a>Value     string;
    <a id="L33"></a>Native    string;
    <a id="L34"></a>As        string;
    <a id="L35"></a>IsIdeal   bool;
    <a id="L36"></a>HasAssign bool;
    <a id="L37"></a>Sizes     []Size;
<a id="L38"></a>}

<a id="L40"></a>var (
    <a id="L41"></a>boolType = &amp;Type{Repr: &#34;*boolType&#34;, Value: &#34;BoolValue&#34;, Native: &#34;bool&#34;, As: &#34;asBool&#34;};
    <a id="L42"></a>uintType = &amp;Type{Repr: &#34;*uintType&#34;, Value: &#34;UintValue&#34;, Native: &#34;uint64&#34;, As: &#34;asUint&#34;,
        <a id="L43"></a>Sizes: []Size{Size{8, &#34;uint8&#34;}, Size{16, &#34;uint16&#34;}, Size{32, &#34;uint32&#34;}, Size{64, &#34;uint64&#34;}, Size{0, &#34;uint&#34;}},
    <a id="L44"></a>};
    <a id="L45"></a>intType = &amp;Type{Repr: &#34;*intType&#34;, Value: &#34;IntValue&#34;, Native: &#34;int64&#34;, As: &#34;asInt&#34;,
        <a id="L46"></a>Sizes: []Size{Size{8, &#34;int8&#34;}, Size{16, &#34;int16&#34;}, Size{32, &#34;int32&#34;}, Size{64, &#34;int64&#34;}, Size{0, &#34;int&#34;}},
    <a id="L47"></a>};
    <a id="L48"></a>idealIntType = &amp;Type{Repr: &#34;*idealIntType&#34;, Value: &#34;IdealIntValue&#34;, Native: &#34;*bignum.Integer&#34;, As: &#34;asIdealInt&#34;, IsIdeal: true};
    <a id="L49"></a>floatType    = &amp;Type{Repr: &#34;*floatType&#34;, Value: &#34;FloatValue&#34;, Native: &#34;float64&#34;, As: &#34;asFloat&#34;,
        <a id="L50"></a>Sizes: []Size{Size{32, &#34;float32&#34;}, Size{64, &#34;float64&#34;}, Size{0, &#34;float&#34;}},
    <a id="L51"></a>};
    <a id="L52"></a>idealFloatType = &amp;Type{Repr: &#34;*idealFloatType&#34;, Value: &#34;IdealFloatValue&#34;, Native: &#34;*bignum.Rational&#34;, As: &#34;asIdealFloat&#34;, IsIdeal: true};
    <a id="L53"></a>stringType     = &amp;Type{Repr: &#34;*stringType&#34;, Value: &#34;StringValue&#34;, Native: &#34;string&#34;, As: &#34;asString&#34;};
    <a id="L54"></a>arrayType      = &amp;Type{Repr: &#34;*ArrayType&#34;, Value: &#34;ArrayValue&#34;, Native: &#34;ArrayValue&#34;, As: &#34;asArray&#34;, HasAssign: true};
    <a id="L55"></a>structType     = &amp;Type{Repr: &#34;*StructType&#34;, Value: &#34;StructValue&#34;, Native: &#34;StructValue&#34;, As: &#34;asStruct&#34;, HasAssign: true};
    <a id="L56"></a>ptrType        = &amp;Type{Repr: &#34;*PtrType&#34;, Value: &#34;PtrValue&#34;, Native: &#34;Value&#34;, As: &#34;asPtr&#34;};
    <a id="L57"></a>funcType       = &amp;Type{Repr: &#34;*FuncType&#34;, Value: &#34;FuncValue&#34;, Native: &#34;Func&#34;, As: &#34;asFunc&#34;};
    <a id="L58"></a>sliceType      = &amp;Type{Repr: &#34;*SliceType&#34;, Value: &#34;SliceValue&#34;, Native: &#34;Slice&#34;, As: &#34;asSlice&#34;};
    <a id="L59"></a>mapType        = &amp;Type{Repr: &#34;*MapType&#34;, Value: &#34;MapValue&#34;, Native: &#34;Map&#34;, As: &#34;asMap&#34;};

    <a id="L61"></a>all = []*Type{
        <a id="L62"></a>boolType,
        <a id="L63"></a>uintType,
        <a id="L64"></a>intType,
        <a id="L65"></a>idealIntType,
        <a id="L66"></a>floatType,
        <a id="L67"></a>idealFloatType,
        <a id="L68"></a>stringType,
        <a id="L69"></a>arrayType,
        <a id="L70"></a>structType,
        <a id="L71"></a>ptrType,
        <a id="L72"></a>funcType,
        <a id="L73"></a>sliceType,
        <a id="L74"></a>mapType,
    <a id="L75"></a>};
    <a id="L76"></a>bools     = all[0:1];
    <a id="L77"></a>integers  = all[1:4];
    <a id="L78"></a>shiftable = all[1:3];
    <a id="L79"></a>numbers   = all[1:6];
    <a id="L80"></a>addable   = all[1:7];
    <a id="L81"></a>cmpable   = []*Type{
        <a id="L82"></a>boolType,
        <a id="L83"></a>uintType,
        <a id="L84"></a>intType,
        <a id="L85"></a>idealIntType,
        <a id="L86"></a>floatType,
        <a id="L87"></a>idealFloatType,
        <a id="L88"></a>stringType,
        <a id="L89"></a>ptrType,
        <a id="L90"></a>funcType,
        <a id="L91"></a>mapType,
    <a id="L92"></a>};
<a id="L93"></a>)

<a id="L95"></a>var unOps = []Op{
    <a id="L96"></a>Op{Name: &#34;Neg&#34;, Expr: &#34;-v&#34;, ConstExpr: &#34;v.Neg()&#34;, Types: numbers},
    <a id="L97"></a>Op{Name: &#34;Not&#34;, Expr: &#34;!v&#34;, Types: bools},
    <a id="L98"></a>Op{Name: &#34;Xor&#34;, Expr: &#34;^v&#34;, ConstExpr: &#34;v.Neg().Sub(bignum.Int(1))&#34;, Types: integers},
<a id="L99"></a>}

<a id="L101"></a>var binOps = []Op{
    <a id="L102"></a>Op{Name: &#34;Add&#34;, Expr: &#34;l + r&#34;, ConstExpr: &#34;l.Add(r)&#34;, Types: addable},
    <a id="L103"></a>Op{Name: &#34;Sub&#34;, Expr: &#34;l - r&#34;, ConstExpr: &#34;l.Sub(r)&#34;, Types: numbers},
    <a id="L104"></a>Op{Name: &#34;Mul&#34;, Expr: &#34;l * r&#34;, ConstExpr: &#34;l.Mul(r)&#34;, Types: numbers},
    <a id="L105"></a>Op{Name: &#34;Quo&#34;,
        <a id="L106"></a>Body: &#34;if r == 0 { t.Abort(DivByZeroError{}) } ret =  l / r&#34;,
        <a id="L107"></a>ConstExpr: &#34;l.Quo(r)&#34;,
        <a id="L108"></a>Types: numbers,
    <a id="L109"></a>},
    <a id="L110"></a>Op{Name: &#34;Rem&#34;,
        <a id="L111"></a>Body: &#34;if r == 0 { t.Abort(DivByZeroError{}) } ret = l % r&#34;,
        <a id="L112"></a>ConstExpr: &#34;l.Rem(r)&#34;,
        <a id="L113"></a>Types: integers,
    <a id="L114"></a>},
    <a id="L115"></a>Op{Name: &#34;And&#34;, Expr: &#34;l &amp; r&#34;, ConstExpr: &#34;l.And(r)&#34;, Types: integers},
    <a id="L116"></a>Op{Name: &#34;Or&#34;, Expr: &#34;l | r&#34;, ConstExpr: &#34;l.Or(r)&#34;, Types: integers},
    <a id="L117"></a>Op{Name: &#34;Xor&#34;, Expr: &#34;l ^ r&#34;, ConstExpr: &#34;l.Xor(r)&#34;, Types: integers},
    <a id="L118"></a>Op{Name: &#34;AndNot&#34;, Expr: &#34;l &amp;^ r&#34;, ConstExpr: &#34;l.AndNot(r)&#34;, Types: integers},
    <a id="L119"></a>Op{Name: &#34;Shl&#34;, Expr: &#34;l &lt;&lt; r&#34;, ConstExpr: &#34;l.Shl(uint(r.Value()))&#34;,
        <a id="L120"></a>AsRightName: &#34;asUint&#34;, Types: shiftable,
    <a id="L121"></a>},
    <a id="L122"></a>Op{Name: &#34;Shr&#34;, Expr: &#34;l &gt;&gt; r&#34;, ConstExpr: &#34;l.Shr(uint(r.Value()))&#34;,
        <a id="L123"></a>AsRightName: &#34;asUint&#34;, Types: shiftable,
    <a id="L124"></a>},
    <a id="L125"></a>Op{Name: &#34;Lss&#34;, Expr: &#34;l &lt; r&#34;, ConstExpr: &#34;l.Cmp(r) &lt; 0&#34;, ReturnType: &#34;bool&#34;, Types: addable},
    <a id="L126"></a>Op{Name: &#34;Gtr&#34;, Expr: &#34;l &gt; r&#34;, ConstExpr: &#34;l.Cmp(r) &gt; 0&#34;, ReturnType: &#34;bool&#34;, Types: addable},
    <a id="L127"></a>Op{Name: &#34;Leq&#34;, Expr: &#34;l &lt;= r&#34;, ConstExpr: &#34;l.Cmp(r) &lt;= 0&#34;, ReturnType: &#34;bool&#34;, Types: addable},
    <a id="L128"></a>Op{Name: &#34;Geq&#34;, Expr: &#34;l &gt;= r&#34;, ConstExpr: &#34;l.Cmp(r) &gt;= 0&#34;, ReturnType: &#34;bool&#34;, Types: addable},
    <a id="L129"></a>Op{Name: &#34;Eql&#34;, Expr: &#34;l == r&#34;, ConstExpr: &#34;l.Cmp(r) == 0&#34;, ReturnType: &#34;bool&#34;, Types: cmpable},
    <a id="L130"></a>Op{Name: &#34;Neq&#34;, Expr: &#34;l != r&#34;, ConstExpr: &#34;l.Cmp(r) != 0&#34;, ReturnType: &#34;bool&#34;, Types: cmpable},
<a id="L131"></a>}

<a id="L133"></a>type Data struct {
    <a id="L134"></a>UnaryOps  []Op;
    <a id="L135"></a>BinaryOps []Op;
    <a id="L136"></a>Types     []*Type;
<a id="L137"></a>}

<a id="L139"></a>var data = Data{
    <a id="L140"></a>unOps,
    <a id="L141"></a>binOps,
    <a id="L142"></a>all,
<a id="L143"></a>}

<a id="L145"></a>const templateStr = `
// This file is machine generated by gen.go.
// 6g gen.go &amp;&amp; 6l gen.6 &amp;&amp; ./6.out &gt;expr1.go

package eval

import (
	&#34;bignum&#34;;
	&#34;log&#34;;
)

/*
 * &#34;As&#34; functions.  These retrieve evaluator functions from an
 * expr, panicking if the requested evaluator has the wrong type.
 */
«.repeated section Types»
«.section IsIdeal»
func (a *expr) «As»() (func() «Native») {
	return a.eval.(func()(«Native»))
}
«.or»
func (a *expr) «As»() (func(*Thread) «Native») {
	return a.eval.(func(*Thread)(«Native»))
}
«.end»
«.end»
func (a *expr) asMulti() (func(*Thread) []Value) {
	return a.eval.(func(*Thread)[]Value)
}

func (a *expr) asInterface() (func(*Thread) interface{}) {
	switch sf := a.eval.(type) {
«.repeated section Types»
«.section IsIdeal»
	case func()«Native»:
		return func(*Thread) interface{} { return sf() }
«.or»
	case func(t *Thread)«Native»:
		return func(t *Thread) interface{} { return sf(t) }
«.end»
«.end»
	default:
		log.Crashf(&#34;unexpected expression node type %T at %v&#34;, a.eval, a.pos);
	}
	panic();
}

/*
 * Operator generators.
 */

func (a *expr) genConstant(v Value) {
	switch a.t.lit().(type) {
«.repeated section Types»
	case «Repr»:
«.section IsIdeal»
		val := v.(«Value»).Get();
		a.eval = func() «Native» { return val }
«.or»
		a.eval = func(t *Thread) «Native» { return v.(«Value»).Get(t) }
«.end»
«.end»
	default:
		log.Crashf(&#34;unexpected constant type %v at %v&#34;, a.t, a.pos);
	}
}

func (a *expr) genIdentOp(level, index int) {
	a.evalAddr = func(t *Thread) Value { return t.f.Get(level, index) };
	switch a.t.lit().(type) {
«.repeated section Types»
«.section IsIdeal»
«.or»
	case «Repr»:
		a.eval = func(t *Thread) «Native» { return t.f.Get(level, index).(«Value»).Get(t) }
«.end»
«.end»
	default:
		log.Crashf(&#34;unexpected identifier type %v at %v&#34;, a.t, a.pos);
	}
}

func (a *expr) genFuncCall(call func(t *Thread) []Value) {
	a.exec = func(t *Thread) { call(t)};
	switch a.t.lit().(type) {
«.repeated section Types»
«.section IsIdeal»
«.or»
	case «Repr»:
		a.eval = func(t *Thread) «Native» { return call(t)[0].(«Value»).Get(t) }
«.end»
«.end»
	case *MultiType:
		a.eval = func(t *Thread) []Value { return call(t) }
	default:
		log.Crashf(&#34;unexpected result type %v at %v&#34;, a.t, a.pos);
	}
}

func (a *expr) genValue(vf func(*Thread) Value) {
	a.evalAddr = vf;
	switch a.t.lit().(type) {
«.repeated section Types»
«.section IsIdeal»
«.or»
	case «Repr»:
		a.eval = func(t *Thread) «Native» { return vf(t).(«Value»).Get(t) }
«.end»
«.end»
	default:
		log.Crashf(&#34;unexpected result type %v at %v&#34;, a.t, a.pos);
	}
}

«.repeated section UnaryOps»
func (a *expr) genUnaryOp«Name»(v *expr) {
	switch a.t.lit().(type) {
«.repeated section Types»
	case «Repr»:
«.section IsIdeal»
		v := v.«As»()();
		val := «ConstExpr»;
		a.eval = func() «Native» { return val }
«.or»
		vf := v.«As»();
		a.eval = func(t *Thread) «Native» { v := vf(t); return «Expr» }
«.end»
«.end»
	default:
		log.Crashf(&#34;unexpected type %v at %v&#34;, a.t, a.pos);
	}
}

«.end»
func (a *expr) genBinOpLogAnd(l, r *expr) {
	lf := l.asBool();
	rf := r.asBool();
	a.eval = func(t *Thread) bool { return lf(t) &amp;&amp; rf(t) }
}

func (a *expr) genBinOpLogOr(l, r *expr) {
	lf := l.asBool();
	rf := r.asBool();
	a.eval = func(t *Thread) bool { return lf(t) || rf(t) }
}

«.repeated section BinaryOps»
func (a *expr) genBinOp«Name»(l, r *expr) {
	switch t := l.t.lit().(type) {
«.repeated section Types»
	case «Repr»:
	«.section IsIdeal»
		l := l.«As»()();
		r := r.«As»()();
		val := «ConstExpr»;
		«.section ReturnType»
		a.eval = func(t *Thread) «ReturnType» { return val }
		«.or»
		a.eval = func() «Native» { return val }
		«.end»
	«.or»
		lf := l.«As»();
		rf := r.«.section AsRightName»«@»«.or»«As»«.end»();
		«.section ReturnType»
		a.eval = func(t *Thread) «@» {
			l, r := lf(t), rf(t);
			return «Expr»
		}
		«.or»
		«.section Sizes»
		switch t.Bits {
		«.repeated section @»
		case «Bits»:
			a.eval = func(t *Thread) «Native» {
				l, r := lf(t), rf(t);
				var ret «Native»;
				«.section Body»
				«Body»;
				«.or»
				ret = «Expr»;
				«.end»
				return «Native»(«Sized»(ret))
			}
		«.end»
		default:
			log.Crashf(&#34;unexpected size %d in type %v at %v&#34;, t.Bits, t, a.pos);
		}
		«.or»
		a.eval = func(t *Thread) «Native» {
			l, r := lf(t), rf(t);
			return «Expr»
		}
		«.end»
		«.end»
	«.end»
	«.end»
	default:
		log.Crashf(&#34;unexpected type %v at %v&#34;, l.t, a.pos);
	}
}

«.end»
func genAssign(lt Type, r *expr) (func(lv Value, t *Thread)) {
	switch lt.lit().(type) {
«.repeated section Types»
«.section IsIdeal»
«.or»
	case «Repr»:
		rf := r.«As»();
		return func(lv Value, t *Thread) { «.section HasAssign»lv.Assign(t, rf(t))«.or»lv.(«Value»).Set(t, rf(t))«.end» }
«.end»
«.end»
	default:
		log.Crashf(&#34;unexpected left operand type %v at %v&#34;, lt, r.pos);
	}
	panic();
}
`

<a id="L364"></a>func main() {
    <a id="L365"></a>t := template.New(nil);
    <a id="L366"></a>t.SetDelims(&#34;«&#34;, &#34;»&#34;);
    <a id="L367"></a>err := t.Parse(templateStr);
    <a id="L368"></a>if err != nil {
        <a id="L369"></a>log.Exit(err)
    <a id="L370"></a>}
    <a id="L371"></a>err = t.Execute(data, os.Stdout);
    <a id="L372"></a>if err != nil {
        <a id="L373"></a>log.Exit(err)
    <a id="L374"></a>}
<a id="L375"></a>}
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
