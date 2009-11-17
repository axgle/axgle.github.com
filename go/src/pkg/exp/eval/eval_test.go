<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/eval/eval_test.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/exp/eval/eval_test.go</h1>

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
    <a id="L9"></a>&#34;flag&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;log&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;reflect&#34;;
    <a id="L14"></a>&#34;testing&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// Print each statement or expression before parsing it</span>
<a id="L18"></a>var noisy = false

<a id="L20"></a>func init() { flag.BoolVar(&amp;noisy, &#34;noisy&#34;, false, &#34;chatter during eval tests&#34;) }

<a id="L22"></a><span class="comment">/*</span>
<a id="L23"></a><span class="comment"> * Generic statement/expression test framework</span>
<a id="L24"></a><span class="comment"> */</span>

<a id="L26"></a>type test []job

<a id="L28"></a>type job struct {
    <a id="L29"></a>code  string;
    <a id="L30"></a>cerr  string;
    <a id="L31"></a>rterr string;
    <a id="L32"></a>val   Value;
    <a id="L33"></a>noval bool;
<a id="L34"></a>}

<a id="L36"></a>func runTests(t *testing.T, baseName string, tests []test) {
    <a id="L37"></a>for i, test := range tests {
        <a id="L38"></a>name := fmt.Sprintf(&#34;%s[%d]&#34;, baseName, i);
        <a id="L39"></a>test.run(t, name);
    <a id="L40"></a>}
<a id="L41"></a>}

<a id="L43"></a>func (a test) run(t *testing.T, name string) {
    <a id="L44"></a>w := newTestWorld();
    <a id="L45"></a>for _, j := range a {
        <a id="L46"></a>src := j.code;
        <a id="L47"></a>if noisy {
            <a id="L48"></a>println(&#34;code:&#34;, src)
        <a id="L49"></a>}

        <a id="L51"></a>code, err := w.Compile(src);
        <a id="L52"></a>if err != nil {
            <a id="L53"></a>if j.cerr == &#34;&#34; {
                <a id="L54"></a>t.Errorf(&#34;%s: Compile %s: %v&#34;, name, src, err);
                <a id="L55"></a>break;
            <a id="L56"></a>}
            <a id="L57"></a>if !match(t, err, j.cerr) {
                <a id="L58"></a>t.Errorf(&#34;%s: Compile %s = error %s; want %v&#34;, name, src, err, j.cerr);
                <a id="L59"></a>break;
            <a id="L60"></a>}
            <a id="L61"></a>continue;
        <a id="L62"></a>}
        <a id="L63"></a>if j.cerr != &#34;&#34; {
            <a id="L64"></a>t.Errorf(&#34;%s: Compile %s succeeded; want %s&#34;, name, src, j.cerr);
            <a id="L65"></a>break;
        <a id="L66"></a>}

        <a id="L68"></a>val, err := code.Run();
        <a id="L69"></a>if err != nil {
            <a id="L70"></a>if j.rterr == &#34;&#34; {
                <a id="L71"></a>t.Errorf(&#34;%s: Run %s: %v&#34;, name, src, err);
                <a id="L72"></a>break;
            <a id="L73"></a>}
            <a id="L74"></a>if !match(t, err, j.rterr) {
                <a id="L75"></a>t.Errorf(&#34;%s: Run %s = error %s; want %v&#34;, name, src, err, j.rterr);
                <a id="L76"></a>break;
            <a id="L77"></a>}
            <a id="L78"></a>continue;
        <a id="L79"></a>}
        <a id="L80"></a>if j.rterr != &#34;&#34; {
            <a id="L81"></a>t.Errorf(&#34;%s: Run %s succeeded; want %s&#34;, name, src, j.rterr);
            <a id="L82"></a>break;
        <a id="L83"></a>}

        <a id="L85"></a>if !j.noval &amp;&amp; !reflect.DeepEqual(val, j.val) {
            <a id="L86"></a>t.Errorf(&#34;%s: Run %s = %T(%v) want %T(%v)&#34;, name, src, val, val, j.val, j.val)
        <a id="L87"></a>}
    <a id="L88"></a>}
<a id="L89"></a>}

<a id="L91"></a>func match(t *testing.T, err os.Error, pat string) bool {
    <a id="L92"></a>ok, errstr := testing.MatchString(pat, err.String());
    <a id="L93"></a>if errstr != &#34;&#34; {
        <a id="L94"></a>t.Fatalf(&#34;compile regexp %s: %v&#34;, pat, errstr)
    <a id="L95"></a>}
    <a id="L96"></a>return ok;
<a id="L97"></a>}


<a id="L100"></a><span class="comment">/*</span>
<a id="L101"></a><span class="comment"> * Test constructors</span>
<a id="L102"></a><span class="comment"> */</span>

<a id="L104"></a><span class="comment">// Expression compile error</span>
<a id="L105"></a>func CErr(expr string, cerr string) test { return test([]job{job{code: expr, cerr: cerr}}) }

<a id="L107"></a><span class="comment">// Expression runtime error</span>
<a id="L108"></a>func RErr(expr string, rterr string) test { return test([]job{job{code: expr, rterr: rterr}}) }

<a id="L110"></a><span class="comment">// Expression value</span>
<a id="L111"></a>func Val(expr string, val interface{}) test {
    <a id="L112"></a>return test([]job{job{code: expr, val: toValue(val)}})
<a id="L113"></a>}

<a id="L115"></a><span class="comment">// Statement runs without error</span>
<a id="L116"></a>func Run(stmts string) test { return test([]job{job{code: stmts, noval: true}}) }

<a id="L118"></a><span class="comment">// Two statements without error.</span>
<a id="L119"></a><span class="comment">// TODO(rsc): Should be possible with Run but the parser</span>
<a id="L120"></a><span class="comment">// won&#39;t let us do both top-level and non-top-level statements.</span>
<a id="L121"></a>func Run2(stmt1, stmt2 string) test {
    <a id="L122"></a>return test([]job{job{code: stmt1, noval: true}, job{code: stmt2, noval: true}})
<a id="L123"></a>}

<a id="L125"></a><span class="comment">// Statement runs and test one expression&#39;s value</span>
<a id="L126"></a>func Val1(stmts string, expr1 string, val1 interface{}) test {
    <a id="L127"></a>return test([]job{
        <a id="L128"></a>job{code: stmts, noval: true},
        <a id="L129"></a>job{code: expr1, val: toValue(val1)},
    <a id="L130"></a>})
<a id="L131"></a>}

<a id="L133"></a><span class="comment">// Statement runs and test two expressions&#39; values</span>
<a id="L134"></a>func Val2(stmts string, expr1 string, val1 interface{}, expr2 string, val2 interface{}) test {
    <a id="L135"></a>return test([]job{
        <a id="L136"></a>job{code: stmts, noval: true},
        <a id="L137"></a>job{code: expr1, val: toValue(val1)},
        <a id="L138"></a>job{code: expr2, val: toValue(val2)},
    <a id="L139"></a>})
<a id="L140"></a>}

<a id="L142"></a><span class="comment">/*</span>
<a id="L143"></a><span class="comment"> * Value constructors</span>
<a id="L144"></a><span class="comment"> */</span>

<a id="L146"></a>type vstruct []interface{}

<a id="L148"></a>type varray []interface{}

<a id="L150"></a>type vslice struct {
    <a id="L151"></a>arr      varray;
    <a id="L152"></a>len, cap int;
<a id="L153"></a>}

<a id="L155"></a>func toValue(val interface{}) Value {
    <a id="L156"></a>switch val := val.(type) {
    <a id="L157"></a>case bool:
        <a id="L158"></a>r := boolV(val);
        <a id="L159"></a>return &amp;r;
    <a id="L160"></a>case uint8:
        <a id="L161"></a>r := uint8V(val);
        <a id="L162"></a>return &amp;r;
    <a id="L163"></a>case uint:
        <a id="L164"></a>r := uintV(val);
        <a id="L165"></a>return &amp;r;
    <a id="L166"></a>case int:
        <a id="L167"></a>r := intV(val);
        <a id="L168"></a>return &amp;r;
    <a id="L169"></a>case *bignum.Integer:
        <a id="L170"></a>return &amp;idealIntV{val}
    <a id="L171"></a>case float:
        <a id="L172"></a>r := floatV(val);
        <a id="L173"></a>return &amp;r;
    <a id="L174"></a>case *bignum.Rational:
        <a id="L175"></a>return &amp;idealFloatV{val}
    <a id="L176"></a>case string:
        <a id="L177"></a>r := stringV(val);
        <a id="L178"></a>return &amp;r;
    <a id="L179"></a>case vstruct:
        <a id="L180"></a>elems := make([]Value, len(val));
        <a id="L181"></a>for i, e := range val {
            <a id="L182"></a>elems[i] = toValue(e)
        <a id="L183"></a>}
        <a id="L184"></a>r := structV(elems);
        <a id="L185"></a>return &amp;r;
    <a id="L186"></a>case varray:
        <a id="L187"></a>elems := make([]Value, len(val));
        <a id="L188"></a>for i, e := range val {
            <a id="L189"></a>elems[i] = toValue(e)
        <a id="L190"></a>}
        <a id="L191"></a>r := arrayV(elems);
        <a id="L192"></a>return &amp;r;
    <a id="L193"></a>case vslice:
        <a id="L194"></a>return &amp;sliceV{Slice{toValue(val.arr).(ArrayValue), int64(val.len), int64(val.cap)}}
    <a id="L195"></a>case Func:
        <a id="L196"></a>return &amp;funcV{val}
    <a id="L197"></a>}
    <a id="L198"></a>log.Crashf(&#34;toValue(%T) not implemented&#34;, val);
    <a id="L199"></a>panic();
<a id="L200"></a>}

<a id="L202"></a><span class="comment">/*</span>
<a id="L203"></a><span class="comment"> * Default test scope</span>
<a id="L204"></a><span class="comment"> */</span>

<a id="L206"></a>type testFunc struct{}

<a id="L208"></a>func (*testFunc) NewFrame() *Frame { return &amp;Frame{nil, &amp;[2]Value{}} }

<a id="L210"></a>func (*testFunc) Call(t *Thread) {
    <a id="L211"></a>n := t.f.Vars[0].(IntValue).Get(t);

    <a id="L213"></a>res := n + 1;

    <a id="L215"></a>t.f.Vars[1].(IntValue).Set(t, res);
<a id="L216"></a>}

<a id="L218"></a>type oneTwoFunc struct{}

<a id="L220"></a>func (*oneTwoFunc) NewFrame() *Frame { return &amp;Frame{nil, &amp;[2]Value{}} }

<a id="L222"></a>func (*oneTwoFunc) Call(t *Thread) {
    <a id="L223"></a>t.f.Vars[0].(IntValue).Set(t, 1);
    <a id="L224"></a>t.f.Vars[1].(IntValue).Set(t, 2);
<a id="L225"></a>}

<a id="L227"></a>type voidFunc struct{}

<a id="L229"></a>func (*voidFunc) NewFrame() *Frame { return &amp;Frame{nil, []Value{}} }

<a id="L231"></a>func (*voidFunc) Call(t *Thread) {}

<a id="L233"></a>func newTestWorld() *World {
    <a id="L234"></a>w := NewWorld();

    <a id="L236"></a>def := func(name string, t Type, val interface{}) { w.DefineVar(name, t, toValue(val)) };

    <a id="L238"></a>w.DefineConst(&#34;c&#34;, IdealIntType, toValue(bignum.Int(1)));
    <a id="L239"></a>def(&#34;i&#34;, IntType, 1);
    <a id="L240"></a>def(&#34;i2&#34;, IntType, 2);
    <a id="L241"></a>def(&#34;u&#34;, UintType, uint(1));
    <a id="L242"></a>def(&#34;f&#34;, FloatType, 1.0);
    <a id="L243"></a>def(&#34;s&#34;, StringType, &#34;abc&#34;);
    <a id="L244"></a>def(&#34;t&#34;, NewStructType([]StructField{StructField{&#34;a&#34;, IntType, false}}), vstruct{1});
    <a id="L245"></a>def(&#34;ai&#34;, NewArrayType(2, IntType), varray{1, 2});
    <a id="L246"></a>def(&#34;aai&#34;, NewArrayType(2, NewArrayType(2, IntType)), varray{varray{1, 2}, varray{3, 4}});
    <a id="L247"></a>def(&#34;aai2&#34;, NewArrayType(2, NewArrayType(2, IntType)), varray{varray{5, 6}, varray{7, 8}});
    <a id="L248"></a>def(&#34;fn&#34;, NewFuncType([]Type{IntType}, false, []Type{IntType}), &amp;testFunc{});
    <a id="L249"></a>def(&#34;oneTwo&#34;, NewFuncType([]Type{}, false, []Type{IntType, IntType}), &amp;oneTwoFunc{});
    <a id="L250"></a>def(&#34;void&#34;, NewFuncType([]Type{}, false, []Type{}), &amp;voidFunc{});
    <a id="L251"></a>def(&#34;sli&#34;, NewSliceType(IntType), vslice{varray{1, 2, 3}, 2, 3});

    <a id="L253"></a>return w;
<a id="L254"></a>}
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
