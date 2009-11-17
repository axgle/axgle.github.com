<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/reflect/all_test.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/reflect/all_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package reflect_test

<a id="L7"></a>import (
    <a id="L8"></a>&#34;io&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>. &#34;reflect&#34;;
    <a id="L11"></a>&#34;testing&#34;;
    <a id="L12"></a>&#34;unsafe&#34;;
<a id="L13"></a>)

<a id="L15"></a>type integer int
<a id="L16"></a>type T struct {
    <a id="L17"></a>a   int;
    <a id="L18"></a>b   float64;
    <a id="L19"></a>c   string;
    <a id="L20"></a>d   *int;
<a id="L21"></a>}

<a id="L23"></a>type pair struct {
    <a id="L24"></a>i   interface{};
    <a id="L25"></a>s   string;
<a id="L26"></a>}

<a id="L28"></a>func isDigit(c uint8) bool { return &#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39; }

<a id="L30"></a>func assert(t *testing.T, s, want string) {
    <a id="L31"></a>if s != want {
        <a id="L32"></a>t.Errorf(&#34;have %#q want %#q&#34;, s, want)
    <a id="L33"></a>}
<a id="L34"></a>}

<a id="L36"></a>func typestring(i interface{}) string { return Typeof(i).String() }

<a id="L38"></a>var typeTests = []pair{
    <a id="L39"></a>pair{struct{ x int }{}, &#34;int&#34;},
    <a id="L40"></a>pair{struct{ x int8 }{}, &#34;int8&#34;},
    <a id="L41"></a>pair{struct{ x int16 }{}, &#34;int16&#34;},
    <a id="L42"></a>pair{struct{ x int32 }{}, &#34;int32&#34;},
    <a id="L43"></a>pair{struct{ x int64 }{}, &#34;int64&#34;},
    <a id="L44"></a>pair{struct{ x uint }{}, &#34;uint&#34;},
    <a id="L45"></a>pair{struct{ x uint8 }{}, &#34;uint8&#34;},
    <a id="L46"></a>pair{struct{ x uint16 }{}, &#34;uint16&#34;},
    <a id="L47"></a>pair{struct{ x uint32 }{}, &#34;uint32&#34;},
    <a id="L48"></a>pair{struct{ x uint64 }{}, &#34;uint64&#34;},
    <a id="L49"></a>pair{struct{ x float }{}, &#34;float&#34;},
    <a id="L50"></a>pair{struct{ x float32 }{}, &#34;float32&#34;},
    <a id="L51"></a>pair{struct{ x float64 }{}, &#34;float64&#34;},
    <a id="L52"></a>pair{struct{ x int8 }{}, &#34;int8&#34;},
    <a id="L53"></a>pair{struct{ x (**int8) }{}, &#34;**int8&#34;},
    <a id="L54"></a>pair{struct{ x (**integer) }{}, &#34;**reflect_test.integer&#34;},
    <a id="L55"></a>pair{struct{ x ([32]int32) }{}, &#34;[32]int32&#34;},
    <a id="L56"></a>pair{struct{ x ([]int8) }{}, &#34;[]int8&#34;},
    <a id="L57"></a>pair{struct{ x (map[string]int32) }{}, &#34;map[string] int32&#34;},
    <a id="L58"></a>pair{struct{ x (chan&lt;- string) }{}, &#34;chan&lt;- string&#34;},
    <a id="L59"></a>pair{struct {
        <a id="L60"></a>x struct {
            <a id="L61"></a>c   chan *int32;
            <a id="L62"></a>d   float32;
        <a id="L63"></a>};
    <a id="L64"></a>}{},
        <a id="L65"></a>&#34;struct { c chan *int32; d float32 }&#34;,
    <a id="L66"></a>},
    <a id="L67"></a>pair{struct{ x (func(a int8, b int32)) }{}, &#34;func(int8, int32)&#34;},
    <a id="L68"></a>pair{struct {
        <a id="L69"></a>x struct {
            <a id="L70"></a>c func(chan *integer, *int8);
        <a id="L71"></a>};
    <a id="L72"></a>}{},
        <a id="L73"></a>&#34;struct { c func(chan *reflect_test.integer, *int8) }&#34;,
    <a id="L74"></a>},
    <a id="L75"></a>pair{struct {
        <a id="L76"></a>x struct {
            <a id="L77"></a>a   int8;
            <a id="L78"></a>b   int32;
        <a id="L79"></a>};
    <a id="L80"></a>}{},
        <a id="L81"></a>&#34;struct { a int8; b int32 }&#34;,
    <a id="L82"></a>},
    <a id="L83"></a>pair{struct {
        <a id="L84"></a>x struct {
            <a id="L85"></a>a   int8;
            <a id="L86"></a>b   int8;
            <a id="L87"></a>c   int32;
        <a id="L88"></a>};
    <a id="L89"></a>}{},
        <a id="L90"></a>&#34;struct { a int8; b int8; c int32 }&#34;,
    <a id="L91"></a>},
    <a id="L92"></a>pair{struct {
        <a id="L93"></a>x struct {
            <a id="L94"></a>a   int8;
            <a id="L95"></a>b   int8;
            <a id="L96"></a>c   int8;
            <a id="L97"></a>d   int32;
        <a id="L98"></a>};
    <a id="L99"></a>}{},
        <a id="L100"></a>&#34;struct { a int8; b int8; c int8; d int32 }&#34;,
    <a id="L101"></a>},
    <a id="L102"></a>pair{struct {
        <a id="L103"></a>x struct {
            <a id="L104"></a>a   int8;
            <a id="L105"></a>b   int8;
            <a id="L106"></a>c   int8;
            <a id="L107"></a>d   int8;
            <a id="L108"></a>e   int32;
        <a id="L109"></a>};
    <a id="L110"></a>}{},
        <a id="L111"></a>&#34;struct { a int8; b int8; c int8; d int8; e int32 }&#34;,
    <a id="L112"></a>},
    <a id="L113"></a>pair{struct {
        <a id="L114"></a>x struct {
            <a id="L115"></a>a   int8;
            <a id="L116"></a>b   int8;
            <a id="L117"></a>c   int8;
            <a id="L118"></a>d   int8;
            <a id="L119"></a>e   int8;
            <a id="L120"></a>f   int32;
        <a id="L121"></a>};
    <a id="L122"></a>}{},
        <a id="L123"></a>&#34;struct { a int8; b int8; c int8; d int8; e int8; f int32 }&#34;,
    <a id="L124"></a>},
    <a id="L125"></a>pair{struct {
        <a id="L126"></a>x struct {
            <a id="L127"></a>a int8 &#34;hi there&#34;;
        <a id="L128"></a>};
    <a id="L129"></a>}{},
        <a id="L130"></a>`struct { a int8 &#34;hi there&#34; }`,
    <a id="L131"></a>},
    <a id="L132"></a>pair{struct {
        <a id="L133"></a>x struct {
            <a id="L134"></a>a int8 &#34;hi \x00there\t\n\&#34;\\&#34;;
        <a id="L135"></a>};
    <a id="L136"></a>}{},
        <a id="L137"></a>`struct { a int8 &#34;hi \x00there\t\n\&#34;\\&#34; }`,
    <a id="L138"></a>},
    <a id="L139"></a>pair{struct {
        <a id="L140"></a>x struct {
            <a id="L141"></a>f func(args ...);
        <a id="L142"></a>};
    <a id="L143"></a>}{},
        <a id="L144"></a>&#34;struct { f func(...) }&#34;,
    <a id="L145"></a>},
    <a id="L146"></a>pair{struct {
        <a id="L147"></a>x (interface {
            <a id="L148"></a>a(func(func(int) int) (func(func(int)) int));
            <a id="L149"></a>b();
        <a id="L150"></a>});
    <a id="L151"></a>}{},
        <a id="L152"></a>&#34;interface { a (func(func(int) (int)) (func(func(int)) (int))); b () }&#34;,
    <a id="L153"></a>},
<a id="L154"></a>}

<a id="L156"></a>var valueTests = []pair{
    <a id="L157"></a>pair{(int8)(0), &#34;8&#34;},
    <a id="L158"></a>pair{(int16)(0), &#34;16&#34;},
    <a id="L159"></a>pair{(int32)(0), &#34;32&#34;},
    <a id="L160"></a>pair{(int64)(0), &#34;64&#34;},
    <a id="L161"></a>pair{(uint8)(0), &#34;8&#34;},
    <a id="L162"></a>pair{(uint16)(0), &#34;16&#34;},
    <a id="L163"></a>pair{(uint32)(0), &#34;32&#34;},
    <a id="L164"></a>pair{(uint64)(0), &#34;64&#34;},
    <a id="L165"></a>pair{(float32)(0), &#34;32.1&#34;},
    <a id="L166"></a>pair{(float64)(0), &#34;64.2&#34;},
    <a id="L167"></a>pair{(string)(&#34;&#34;), &#34;stringy cheese&#34;},
    <a id="L168"></a>pair{(bool)(false), &#34;true&#34;},
    <a id="L169"></a>pair{(*int8)(nil), &#34;*int8(0)&#34;},
    <a id="L170"></a>pair{(**int8)(nil), &#34;**int8(0)&#34;},
    <a id="L171"></a>pair{([5]int32){}, &#34;[5]int32{0, 0, 0, 0, 0}&#34;},
    <a id="L172"></a>pair{(**integer)(nil), &#34;**reflect_test.integer(0)&#34;},
    <a id="L173"></a>pair{(map[string]int32)(nil), &#34;map[string] int32{&lt;can&#39;t iterate on maps&gt;}&#34;},
    <a id="L174"></a>pair{(chan&lt;- string)(nil), &#34;chan&lt;- string&#34;},
    <a id="L175"></a>pair{(struct {
        <a id="L176"></a>c   chan *int32;
        <a id="L177"></a>d   float32;
    <a id="L178"></a>}){},
        <a id="L179"></a>&#34;struct { c chan *int32; d float32 }{chan *int32, 0}&#34;,
    <a id="L180"></a>},
    <a id="L181"></a>pair{(func(a int8, b int32))(nil), &#34;func(int8, int32)(0)&#34;},
    <a id="L182"></a>pair{(struct {
        <a id="L183"></a>c func(chan *integer, *int8);
    <a id="L184"></a>}){},
        <a id="L185"></a>&#34;struct { c func(chan *reflect_test.integer, *int8) }{func(chan *reflect_test.integer, *int8)(0)}&#34;,
    <a id="L186"></a>},
    <a id="L187"></a>pair{(struct {
        <a id="L188"></a>a   int8;
        <a id="L189"></a>b   int32;
    <a id="L190"></a>}){},
        <a id="L191"></a>&#34;struct { a int8; b int32 }{0, 0}&#34;,
    <a id="L192"></a>},
    <a id="L193"></a>pair{(struct {
        <a id="L194"></a>a   int8;
        <a id="L195"></a>b   int8;
        <a id="L196"></a>c   int32;
    <a id="L197"></a>}){},
        <a id="L198"></a>&#34;struct { a int8; b int8; c int32 }{0, 0, 0}&#34;,
    <a id="L199"></a>},
<a id="L200"></a>}

<a id="L202"></a>func testType(t *testing.T, i int, typ Type, want string) {
    <a id="L203"></a>s := typ.String();
    <a id="L204"></a>if s != want {
        <a id="L205"></a>t.Errorf(&#34;#%d: have %#q, want %#q&#34;, i, s, want)
    <a id="L206"></a>}
<a id="L207"></a>}

<a id="L209"></a>func TestTypes(t *testing.T) {
    <a id="L210"></a>for i, tt := range typeTests {
        <a id="L211"></a>testType(t, i, NewValue(tt.i).(*StructValue).Field(0).Type(), tt.s)
    <a id="L212"></a>}
<a id="L213"></a>}

<a id="L215"></a>func TestSet(t *testing.T) {
    <a id="L216"></a>for i, tt := range valueTests {
        <a id="L217"></a>v := NewValue(tt.i);
        <a id="L218"></a>switch v := v.(type) {
        <a id="L219"></a>case *IntValue:
            <a id="L220"></a>v.Set(132)
        <a id="L221"></a>case *Int8Value:
            <a id="L222"></a>v.Set(8)
        <a id="L223"></a>case *Int16Value:
            <a id="L224"></a>v.Set(16)
        <a id="L225"></a>case *Int32Value:
            <a id="L226"></a>v.Set(32)
        <a id="L227"></a>case *Int64Value:
            <a id="L228"></a>v.Set(64)
        <a id="L229"></a>case *UintValue:
            <a id="L230"></a>v.Set(132)
        <a id="L231"></a>case *Uint8Value:
            <a id="L232"></a>v.Set(8)
        <a id="L233"></a>case *Uint16Value:
            <a id="L234"></a>v.Set(16)
        <a id="L235"></a>case *Uint32Value:
            <a id="L236"></a>v.Set(32)
        <a id="L237"></a>case *Uint64Value:
            <a id="L238"></a>v.Set(64)
        <a id="L239"></a>case *FloatValue:
            <a id="L240"></a>v.Set(3200.0)
        <a id="L241"></a>case *Float32Value:
            <a id="L242"></a>v.Set(32.1)
        <a id="L243"></a>case *Float64Value:
            <a id="L244"></a>v.Set(64.2)
        <a id="L245"></a>case *StringValue:
            <a id="L246"></a>v.Set(&#34;stringy cheese&#34;)
        <a id="L247"></a>case *BoolValue:
            <a id="L248"></a>v.Set(true)
        <a id="L249"></a>}
        <a id="L250"></a>s := valueToString(v);
        <a id="L251"></a>if s != tt.s {
            <a id="L252"></a>t.Errorf(&#34;#%d: have %#q, want %#q&#34;, i, s, tt.s)
        <a id="L253"></a>}
    <a id="L254"></a>}
<a id="L255"></a>}

<a id="L257"></a>func TestSetValue(t *testing.T) {
    <a id="L258"></a>for i, tt := range valueTests {
        <a id="L259"></a>v := NewValue(tt.i);
        <a id="L260"></a>switch v := v.(type) {
        <a id="L261"></a>case *IntValue:
            <a id="L262"></a>v.SetValue(NewValue(int(132)))
        <a id="L263"></a>case *Int8Value:
            <a id="L264"></a>v.SetValue(NewValue(int8(8)))
        <a id="L265"></a>case *Int16Value:
            <a id="L266"></a>v.SetValue(NewValue(int16(16)))
        <a id="L267"></a>case *Int32Value:
            <a id="L268"></a>v.SetValue(NewValue(int32(32)))
        <a id="L269"></a>case *Int64Value:
            <a id="L270"></a>v.SetValue(NewValue(int64(64)))
        <a id="L271"></a>case *UintValue:
            <a id="L272"></a>v.SetValue(NewValue(uint(132)))
        <a id="L273"></a>case *Uint8Value:
            <a id="L274"></a>v.SetValue(NewValue(uint8(8)))
        <a id="L275"></a>case *Uint16Value:
            <a id="L276"></a>v.SetValue(NewValue(uint16(16)))
        <a id="L277"></a>case *Uint32Value:
            <a id="L278"></a>v.SetValue(NewValue(uint32(32)))
        <a id="L279"></a>case *Uint64Value:
            <a id="L280"></a>v.SetValue(NewValue(uint64(64)))
        <a id="L281"></a>case *FloatValue:
            <a id="L282"></a>v.SetValue(NewValue(float(3200.0)))
        <a id="L283"></a>case *Float32Value:
            <a id="L284"></a>v.SetValue(NewValue(float32(32.1)))
        <a id="L285"></a>case *Float64Value:
            <a id="L286"></a>v.SetValue(NewValue(float64(64.2)))
        <a id="L287"></a>case *StringValue:
            <a id="L288"></a>v.SetValue(NewValue(&#34;stringy cheese&#34;))
        <a id="L289"></a>case *BoolValue:
            <a id="L290"></a>v.SetValue(NewValue(true))
        <a id="L291"></a>}
        <a id="L292"></a>s := valueToString(v);
        <a id="L293"></a>if s != tt.s {
            <a id="L294"></a>t.Errorf(&#34;#%d: have %#q, want %#q&#34;, i, s, tt.s)
        <a id="L295"></a>}
    <a id="L296"></a>}
<a id="L297"></a>}

<a id="L299"></a>var _i = 7

<a id="L301"></a>var valueToStringTests = []pair{
    <a id="L302"></a>pair{123, &#34;123&#34;},
    <a id="L303"></a>pair{123.4, &#34;123.4&#34;},
    <a id="L304"></a>pair{byte(123), &#34;123&#34;},
    <a id="L305"></a>pair{&#34;abc&#34;, &#34;abc&#34;},
    <a id="L306"></a>pair{T{123, 456.75, &#34;hello&#34;, &amp;_i}, &#34;reflect_test.T{123, 456.75, hello, *int(&amp;7)}&#34;},
    <a id="L307"></a>pair{new(chan *T), &#34;*chan *reflect_test.T(&amp;chan *reflect_test.T)&#34;},
    <a id="L308"></a>pair{[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, &#34;[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}&#34;},
    <a id="L309"></a>pair{&amp;[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, &#34;*[10]int(&amp;[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})&#34;},
    <a id="L310"></a>pair{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, &#34;[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}&#34;},
    <a id="L311"></a>pair{&amp;[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, &#34;*[]int(&amp;[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})&#34;},
<a id="L312"></a>}

<a id="L314"></a>func TestValueToString(t *testing.T) {
    <a id="L315"></a>for i, test := range valueToStringTests {
        <a id="L316"></a>s := valueToString(NewValue(test.i));
        <a id="L317"></a>if s != test.s {
            <a id="L318"></a>t.Errorf(&#34;#%d: have %#q, want %#q&#34;, i, s, test.s)
        <a id="L319"></a>}
    <a id="L320"></a>}
<a id="L321"></a>}

<a id="L323"></a>func TestArrayElemSet(t *testing.T) {
    <a id="L324"></a>v := NewValue([10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10});
    <a id="L325"></a>v.(*ArrayValue).Elem(4).(*IntValue).Set(123);
    <a id="L326"></a>s := valueToString(v);
    <a id="L327"></a>const want = &#34;[10]int{1, 2, 3, 4, 123, 6, 7, 8, 9, 10}&#34;;
    <a id="L328"></a>if s != want {
        <a id="L329"></a>t.Errorf(&#34;[10]int: have %#q want %#q&#34;, s, want)
    <a id="L330"></a>}

    <a id="L332"></a>v = NewValue([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10});
    <a id="L333"></a>v.(*SliceValue).Elem(4).(*IntValue).Set(123);
    <a id="L334"></a>s = valueToString(v);
    <a id="L335"></a>const want1 = &#34;[]int{1, 2, 3, 4, 123, 6, 7, 8, 9, 10}&#34;;
    <a id="L336"></a>if s != want1 {
        <a id="L337"></a>t.Errorf(&#34;[]int: have %#q want %#q&#34;, s, want1)
    <a id="L338"></a>}
<a id="L339"></a>}

<a id="L341"></a>func TestPtrPointTo(t *testing.T) {
    <a id="L342"></a>var ip *int32;
    <a id="L343"></a>var i int32 = 1234;
    <a id="L344"></a>vip := NewValue(&amp;ip);
    <a id="L345"></a>vi := NewValue(i);
    <a id="L346"></a>vip.(*PtrValue).Elem().(*PtrValue).PointTo(vi);
    <a id="L347"></a>if *ip != 1234 {
        <a id="L348"></a>t.Errorf(&#34;got %d, want 1234&#34;, *ip)
    <a id="L349"></a>}
<a id="L350"></a>}

<a id="L352"></a>func TestAll(t *testing.T) {
    <a id="L353"></a>testType(t, 1, Typeof((int8)(0)), &#34;int8&#34;);
    <a id="L354"></a>testType(t, 2, Typeof((*int8)(nil)).(*PtrType).Elem(), &#34;int8&#34;);

    <a id="L356"></a>typ := Typeof((*struct {
        <a id="L357"></a>c   chan *int32;
        <a id="L358"></a>d   float32;
    <a id="L359"></a>})(nil));
    <a id="L360"></a>testType(t, 3, typ, &#34;*struct { c chan *int32; d float32 }&#34;);
    <a id="L361"></a>etyp := typ.(*PtrType).Elem();
    <a id="L362"></a>testType(t, 4, etyp, &#34;struct { c chan *int32; d float32 }&#34;);
    <a id="L363"></a>styp := etyp.(*StructType);
    <a id="L364"></a>f := styp.Field(0);
    <a id="L365"></a>testType(t, 5, f.Type, &#34;chan *int32&#34;);

    <a id="L367"></a>f, present := styp.FieldByName(&#34;d&#34;);
    <a id="L368"></a>if !present {
        <a id="L369"></a>t.Errorf(&#34;FieldByName says present field is absent&#34;)
    <a id="L370"></a>}
    <a id="L371"></a>testType(t, 6, f.Type, &#34;float32&#34;);

    <a id="L373"></a>f, present = styp.FieldByName(&#34;absent&#34;);
    <a id="L374"></a>if present {
        <a id="L375"></a>t.Errorf(&#34;FieldByName says absent field is present&#34;)
    <a id="L376"></a>}

    <a id="L378"></a>typ = Typeof([32]int32{});
    <a id="L379"></a>testType(t, 7, typ, &#34;[32]int32&#34;);
    <a id="L380"></a>testType(t, 8, typ.(*ArrayType).Elem(), &#34;int32&#34;);

    <a id="L382"></a>typ = Typeof((map[string]*int32)(nil));
    <a id="L383"></a>testType(t, 9, typ, &#34;map[string] *int32&#34;);
    <a id="L384"></a>mtyp := typ.(*MapType);
    <a id="L385"></a>testType(t, 10, mtyp.Key(), &#34;string&#34;);
    <a id="L386"></a>testType(t, 11, mtyp.Elem(), &#34;*int32&#34;);

    <a id="L388"></a>typ = Typeof((chan&lt;- string)(nil));
    <a id="L389"></a>testType(t, 12, typ, &#34;chan&lt;- string&#34;);
    <a id="L390"></a>testType(t, 13, typ.(*ChanType).Elem(), &#34;string&#34;);

    <a id="L392"></a><span class="comment">// make sure tag strings are not part of element type</span>
    <a id="L393"></a>typ = Typeof(struct {
        <a id="L394"></a>d []uint32 &#34;TAG&#34;;
    <a id="L395"></a>}{}).(*StructType).Field(0).Type;
    <a id="L396"></a>testType(t, 14, typ, &#34;[]uint32&#34;);
<a id="L397"></a>}

<a id="L399"></a>func TestInterfaceGet(t *testing.T) {
    <a id="L400"></a>var inter struct {
        <a id="L401"></a>e interface{};
    <a id="L402"></a>}
    <a id="L403"></a>inter.e = 123.456;
    <a id="L404"></a>v1 := NewValue(&amp;inter);
    <a id="L405"></a>v2 := v1.(*PtrValue).Elem().(*StructValue).Field(0);
    <a id="L406"></a>assert(t, v2.Type().String(), &#34;interface { }&#34;);
    <a id="L407"></a>i2 := v2.(*InterfaceValue).Interface();
    <a id="L408"></a>v3 := NewValue(i2);
    <a id="L409"></a>assert(t, v3.Type().String(), &#34;float&#34;);
<a id="L410"></a>}

<a id="L412"></a>func TestInterfaceValue(t *testing.T) {
    <a id="L413"></a>var inter struct {
        <a id="L414"></a>e interface{};
    <a id="L415"></a>}
    <a id="L416"></a>inter.e = 123.456;
    <a id="L417"></a>v1 := NewValue(&amp;inter);
    <a id="L418"></a>v2 := v1.(*PtrValue).Elem().(*StructValue).Field(0);
    <a id="L419"></a>assert(t, v2.Type().String(), &#34;interface { }&#34;);
    <a id="L420"></a>v3 := v2.(*InterfaceValue).Elem();
    <a id="L421"></a>assert(t, v3.Type().String(), &#34;float&#34;);

    <a id="L423"></a>i3 := v2.Interface();
    <a id="L424"></a>if _, ok := i3.(float); !ok {
        <a id="L425"></a>t.Error(&#34;v2.Interface() did not return float, got &#34;, Typeof(i3))
    <a id="L426"></a>}
<a id="L427"></a>}

<a id="L429"></a>func TestFunctionValue(t *testing.T) {
    <a id="L430"></a>v := NewValue(func() {});
    <a id="L431"></a>if v.Interface() != v.Interface() {
        <a id="L432"></a>t.Fatalf(&#34;TestFunction != itself&#34;)
    <a id="L433"></a>}
    <a id="L434"></a>assert(t, v.Type().String(), &#34;func()&#34;);
<a id="L435"></a>}

<a id="L437"></a>func TestCopyArray(t *testing.T) {
    <a id="L438"></a>a := []int{1, 2, 3, 4, 10, 9, 8, 7};
    <a id="L439"></a>b := []int{11, 22, 33, 44, 1010, 99, 88, 77, 66, 55, 44};
    <a id="L440"></a>c := []int{11, 22, 33, 44, 1010, 99, 88, 77, 66, 55, 44};
    <a id="L441"></a>va := NewValue(&amp;a);
    <a id="L442"></a>vb := NewValue(&amp;b);
    <a id="L443"></a>for i := 0; i &lt; len(b); i++ {
        <a id="L444"></a>if b[i] != c[i] {
            <a id="L445"></a>t.Fatalf(&#34;b != c before test&#34;)
        <a id="L446"></a>}
    <a id="L447"></a>}
    <a id="L448"></a>aa := va.(*PtrValue).Elem().(*SliceValue);
    <a id="L449"></a>ab := vb.(*PtrValue).Elem().(*SliceValue);
    <a id="L450"></a>for tocopy := 1; tocopy &lt;= 7; tocopy++ {
        <a id="L451"></a>aa.SetLen(tocopy);
        <a id="L452"></a>ArrayCopy(ab, aa);
        <a id="L453"></a>aa.SetLen(8);
        <a id="L454"></a>for i := 0; i &lt; tocopy; i++ {
            <a id="L455"></a>if a[i] != b[i] {
                <a id="L456"></a>t.Errorf(&#34;(i) tocopy=%d a[%d]=%d, b[%d]=%d&#34;,
                    <a id="L457"></a>tocopy, i, a[i], i, b[i])
            <a id="L458"></a>}
        <a id="L459"></a>}
        <a id="L460"></a>for i := tocopy; i &lt; len(b); i++ {
            <a id="L461"></a>if b[i] != c[i] {
                <a id="L462"></a>if i &lt; len(a) {
                    <a id="L463"></a>t.Errorf(&#34;(ii) tocopy=%d a[%d]=%d, b[%d]=%d, c[%d]=%d&#34;,
                        <a id="L464"></a>tocopy, i, a[i], i, b[i], i, c[i])
                <a id="L465"></a>} else {
                    <a id="L466"></a>t.Errorf(&#34;(iii) tocopy=%d b[%d]=%d, c[%d]=%d&#34;,
                        <a id="L467"></a>tocopy, i, b[i], i, c[i])
                <a id="L468"></a>}
            <a id="L469"></a>} else {
                <a id="L470"></a>t.Logf(&#34;tocopy=%d elem %d is okay\n&#34;, tocopy, i)
            <a id="L471"></a>}
        <a id="L472"></a>}
    <a id="L473"></a>}
<a id="L474"></a>}

<a id="L476"></a>func TestBigUnnamedStruct(t *testing.T) {
    <a id="L477"></a>b := struct{ a, b, c, d int64 }{1, 2, 3, 4};
    <a id="L478"></a>v := NewValue(b);
    <a id="L479"></a>b1 := v.Interface().(struct {
        <a id="L480"></a>a, b, c, d int64;
    <a id="L481"></a>});
    <a id="L482"></a>if b1.a != b.a || b1.b != b.b || b1.c != b.c || b1.d != b.d {
        <a id="L483"></a>t.Errorf(&#34;NewValue(%v).Interface().(*Big) = %v&#34;, b, b1)
    <a id="L484"></a>}
<a id="L485"></a>}

<a id="L487"></a>type big struct {
    <a id="L488"></a>a, b, c, d, e int64;
<a id="L489"></a>}

<a id="L491"></a>func TestBigStruct(t *testing.T) {
    <a id="L492"></a>b := big{1, 2, 3, 4, 5};
    <a id="L493"></a>v := NewValue(b);
    <a id="L494"></a>b1 := v.Interface().(big);
    <a id="L495"></a>if b1.a != b.a || b1.b != b.b || b1.c != b.c || b1.d != b.d || b1.e != b.e {
        <a id="L496"></a>t.Errorf(&#34;NewValue(%v).Interface().(big) = %v&#34;, b, b1)
    <a id="L497"></a>}
<a id="L498"></a>}

<a id="L500"></a>type Basic struct {
    <a id="L501"></a>x   int;
    <a id="L502"></a>y   float32;
<a id="L503"></a>}

<a id="L505"></a>type NotBasic Basic

<a id="L507"></a>type DeepEqualTest struct {
    <a id="L508"></a>a, b interface{};
    <a id="L509"></a>eq   bool;
<a id="L510"></a>}

<a id="L512"></a>var deepEqualTests = []DeepEqualTest{
    <a id="L513"></a><span class="comment">// Equalities</span>
    <a id="L514"></a>DeepEqualTest{1, 1, true},
    <a id="L515"></a>DeepEqualTest{int32(1), int32(1), true},
    <a id="L516"></a>DeepEqualTest{0.5, 0.5, true},
    <a id="L517"></a>DeepEqualTest{float32(0.5), float32(0.5), true},
    <a id="L518"></a>DeepEqualTest{&#34;hello&#34;, &#34;hello&#34;, true},
    <a id="L519"></a>DeepEqualTest{make([]int, 10), make([]int, 10), true},
    <a id="L520"></a>DeepEqualTest{&amp;[3]int{1, 2, 3}, &amp;[3]int{1, 2, 3}, true},
    <a id="L521"></a>DeepEqualTest{Basic{1, 0.5}, Basic{1, 0.5}, true},
    <a id="L522"></a>DeepEqualTest{os.Error(nil), os.Error(nil), true},
    <a id="L523"></a>DeepEqualTest{map[int]string{1: &#34;one&#34;, 2: &#34;two&#34;}, map[int]string{2: &#34;two&#34;, 1: &#34;one&#34;}, true},

    <a id="L525"></a><span class="comment">// Inequalities</span>
    <a id="L526"></a>DeepEqualTest{1, 2, false},
    <a id="L527"></a>DeepEqualTest{int32(1), int32(2), false},
    <a id="L528"></a>DeepEqualTest{0.5, 0.6, false},
    <a id="L529"></a>DeepEqualTest{float32(0.5), float32(0.6), false},
    <a id="L530"></a>DeepEqualTest{&#34;hello&#34;, &#34;hey&#34;, false},
    <a id="L531"></a>DeepEqualTest{make([]int, 10), make([]int, 11), false},
    <a id="L532"></a>DeepEqualTest{&amp;[3]int{1, 2, 3}, &amp;[3]int{1, 2, 4}, false},
    <a id="L533"></a>DeepEqualTest{Basic{1, 0.5}, Basic{1, 0.6}, false},
    <a id="L534"></a>DeepEqualTest{Basic{1, 0}, Basic{2, 0}, false},
    <a id="L535"></a>DeepEqualTest{map[int]string{1: &#34;one&#34;, 3: &#34;two&#34;}, map[int]string{2: &#34;two&#34;, 1: &#34;one&#34;}, false},
    <a id="L536"></a>DeepEqualTest{map[int]string{1: &#34;one&#34;, 2: &#34;txo&#34;}, map[int]string{2: &#34;two&#34;, 1: &#34;one&#34;}, false},
    <a id="L537"></a>DeepEqualTest{map[int]string{1: &#34;one&#34;}, map[int]string{2: &#34;two&#34;, 1: &#34;one&#34;}, false},
    <a id="L538"></a>DeepEqualTest{map[int]string{2: &#34;two&#34;, 1: &#34;one&#34;}, map[int]string{1: &#34;one&#34;}, false},
    <a id="L539"></a>DeepEqualTest{nil, 1, false},
    <a id="L540"></a>DeepEqualTest{1, nil, false},

    <a id="L542"></a><span class="comment">// Mismatched types</span>
    <a id="L543"></a>DeepEqualTest{1, 1.0, false},
    <a id="L544"></a>DeepEqualTest{int32(1), int64(1), false},
    <a id="L545"></a>DeepEqualTest{0.5, &#34;hello&#34;, false},
    <a id="L546"></a>DeepEqualTest{[]int{1, 2, 3}, [3]int{1, 2, 3}, false},
    <a id="L547"></a>DeepEqualTest{&amp;[3]interface{}{1, 2, 4}, &amp;[3]interface{}{1, 2, &#34;s&#34;}, false},
    <a id="L548"></a>DeepEqualTest{Basic{1, 0.5}, NotBasic{1, 0.5}, false},
    <a id="L549"></a>DeepEqualTest{map[uint]string{1: &#34;one&#34;, 2: &#34;two&#34;}, map[int]string{2: &#34;two&#34;, 1: &#34;one&#34;}, false},
<a id="L550"></a>}

<a id="L552"></a>func TestDeepEqual(t *testing.T) {
    <a id="L553"></a>for _, test := range deepEqualTests {
        <a id="L554"></a>if r := DeepEqual(test.a, test.b); r != test.eq {
            <a id="L555"></a>t.Errorf(&#34;DeepEqual(%v, %v) = %v, want %v&#34;, test.a, test.b, r, test.eq)
        <a id="L556"></a>}
    <a id="L557"></a>}
<a id="L558"></a>}

<a id="L560"></a>func TestTypeof(t *testing.T) {
    <a id="L561"></a>for _, test := range deepEqualTests {
        <a id="L562"></a>v := NewValue(test.a);
        <a id="L563"></a>if v == nil {
            <a id="L564"></a>continue
        <a id="L565"></a>}
        <a id="L566"></a>typ := Typeof(test.a);
        <a id="L567"></a>if typ != v.Type() {
            <a id="L568"></a>t.Errorf(&#34;Typeof(%v) = %v, but NewValue(%v).Type() = %v&#34;, test.a, typ, test.a, v.Type())
        <a id="L569"></a>}
    <a id="L570"></a>}
<a id="L571"></a>}

<a id="L573"></a>type Recursive struct {
    <a id="L574"></a>x   int;
    <a id="L575"></a>r   *Recursive;
<a id="L576"></a>}

<a id="L578"></a>func TestDeepEqualRecursiveStruct(t *testing.T) {
    <a id="L579"></a>a, b := new(Recursive), new(Recursive);
    <a id="L580"></a>*a = Recursive{12, a};
    <a id="L581"></a>*b = Recursive{12, b};
    <a id="L582"></a>if !DeepEqual(a, b) {
        <a id="L583"></a>t.Error(&#34;DeepEqual(recursive same) = false, want true&#34;)
    <a id="L584"></a>}
<a id="L585"></a>}

<a id="L587"></a>type Complex struct {
    <a id="L588"></a>a   int;
    <a id="L589"></a>b   [3]*Complex;
    <a id="L590"></a>c   *string;
    <a id="L591"></a>d   map[float]float;
<a id="L592"></a>}

<a id="L594"></a>func TestDeepEqualComplexStruct(t *testing.T) {
    <a id="L595"></a>m := make(map[float]float);
    <a id="L596"></a>stra, strb := &#34;hello&#34;, &#34;hello&#34;;
    <a id="L597"></a>a, b := new(Complex), new(Complex);
    <a id="L598"></a>*a = Complex{5, [3]*Complex{a, b, a}, &amp;stra, m};
    <a id="L599"></a>*b = Complex{5, [3]*Complex{b, a, a}, &amp;strb, m};
    <a id="L600"></a>if !DeepEqual(a, b) {
        <a id="L601"></a>t.Error(&#34;DeepEqual(complex same) = false, want true&#34;)
    <a id="L602"></a>}
<a id="L603"></a>}

<a id="L605"></a>func TestDeepEqualComplexStructInequality(t *testing.T) {
    <a id="L606"></a>m := make(map[float]float);
    <a id="L607"></a>stra, strb := &#34;hello&#34;, &#34;helloo&#34;; <span class="comment">// Difference is here</span>
    <a id="L608"></a>a, b := new(Complex), new(Complex);
    <a id="L609"></a>*a = Complex{5, [3]*Complex{a, b, a}, &amp;stra, m};
    <a id="L610"></a>*b = Complex{5, [3]*Complex{b, a, a}, &amp;strb, m};
    <a id="L611"></a>if DeepEqual(a, b) {
        <a id="L612"></a>t.Error(&#34;DeepEqual(complex different) = true, want false&#34;)
    <a id="L613"></a>}
<a id="L614"></a>}


<a id="L617"></a>func check2ndField(x interface{}, offs uintptr, t *testing.T) {
    <a id="L618"></a>s := NewValue(x).(*StructValue);
    <a id="L619"></a>f := s.Type().(*StructType).Field(1);
    <a id="L620"></a>if f.Offset != offs {
        <a id="L621"></a>t.Error(&#34;mismatched offsets in structure alignment:&#34;, f.Offset, offs)
    <a id="L622"></a>}
<a id="L623"></a>}

<a id="L625"></a><span class="comment">// Check that structure alignment &amp; offsets viewed through reflect agree with those</span>
<a id="L626"></a><span class="comment">// from the compiler itself.</span>
<a id="L627"></a>func TestAlignment(t *testing.T) {
    <a id="L628"></a>type T1inner struct {
        <a id="L629"></a>a int;
    <a id="L630"></a>}
    <a id="L631"></a>type T1 struct {
        <a id="L632"></a>T1inner;
        <a id="L633"></a>f   int;
    <a id="L634"></a>}
    <a id="L635"></a>type T2inner struct {
        <a id="L636"></a>a, b int;
    <a id="L637"></a>}
    <a id="L638"></a>type T2 struct {
        <a id="L639"></a>T2inner;
        <a id="L640"></a>f   int;
    <a id="L641"></a>}

    <a id="L643"></a>x := T1{T1inner{2}, 17};
    <a id="L644"></a>check2ndField(x, uintptr(unsafe.Pointer(&amp;x.f))-uintptr(unsafe.Pointer(&amp;x)), t);

    <a id="L646"></a>x1 := T2{T2inner{2, 3}, 17};
    <a id="L647"></a>check2ndField(x1, uintptr(unsafe.Pointer(&amp;x1.f))-uintptr(unsafe.Pointer(&amp;x1)), t);
<a id="L648"></a>}

<a id="L650"></a>type IsNiller interface {
    <a id="L651"></a>IsNil() bool;
<a id="L652"></a>}

<a id="L654"></a>func Nil(a interface{}, t *testing.T) {
    <a id="L655"></a>n := NewValue(a).(*StructValue).Field(0).(IsNiller);
    <a id="L656"></a>if !n.IsNil() {
        <a id="L657"></a>t.Errorf(&#34;%v should be nil&#34;, a)
    <a id="L658"></a>}
<a id="L659"></a>}

<a id="L661"></a>func NotNil(a interface{}, t *testing.T) {
    <a id="L662"></a>n := NewValue(a).(*StructValue).Field(0).(IsNiller);
    <a id="L663"></a>if n.IsNil() {
        <a id="L664"></a>t.Errorf(&#34;value of type %v should not be nil&#34;, NewValue(a).Type().String())
    <a id="L665"></a>}
<a id="L666"></a>}

<a id="L668"></a>func TestIsNil(t *testing.T) {
    <a id="L669"></a><span class="comment">// These do not implement IsNil</span>
    <a id="L670"></a>doNotNil := []interface{}{int(0), float32(0), struct{ a int }{}};
    <a id="L671"></a>for _, ts := range doNotNil {
        <a id="L672"></a>ty := Typeof(ts);
        <a id="L673"></a>v := MakeZero(ty);
        <a id="L674"></a>if _, ok := v.(IsNiller); ok {
            <a id="L675"></a>t.Errorf(&#34;%s is nilable; should not be&#34;, ts)
        <a id="L676"></a>}
    <a id="L677"></a>}

    <a id="L679"></a><span class="comment">// These do implement IsNil.</span>
    <a id="L680"></a><span class="comment">// Wrap in extra struct to hide interface type.</span>
    <a id="L681"></a>doNil := []interface{}{
        <a id="L682"></a>struct{ x *int }{},
        <a id="L683"></a>struct{ x interface{} }{},
        <a id="L684"></a>struct{ x map[string]int }{},
        <a id="L685"></a>struct{ x func() bool }{},
        <a id="L686"></a>struct{ x chan int }{},
        <a id="L687"></a>struct{ x []string }{},
    <a id="L688"></a>};
    <a id="L689"></a>for _, ts := range doNil {
        <a id="L690"></a>ty := Typeof(ts).(*StructType).Field(0).Type;
        <a id="L691"></a>v := MakeZero(ty);
        <a id="L692"></a>if _, ok := v.(IsNiller); !ok {
            <a id="L693"></a>t.Errorf(&#34;%s %T is not nilable; should be&#34;, ts, v)
        <a id="L694"></a>}
    <a id="L695"></a>}

    <a id="L697"></a><span class="comment">// Check the implementations</span>
    <a id="L698"></a>var pi struct {
        <a id="L699"></a>x *int;
    <a id="L700"></a>}
    <a id="L701"></a>Nil(pi, t);
    <a id="L702"></a>pi.x = new(int);
    <a id="L703"></a>NotNil(pi, t);

    <a id="L705"></a>var si struct {
        <a id="L706"></a>x []int;
    <a id="L707"></a>}
    <a id="L708"></a>Nil(si, t);
    <a id="L709"></a>si.x = make([]int, 10);
    <a id="L710"></a>NotNil(si, t);

    <a id="L712"></a>var ci struct {
        <a id="L713"></a>x chan int;
    <a id="L714"></a>}
    <a id="L715"></a>Nil(ci, t);
    <a id="L716"></a>ci.x = make(chan int);
    <a id="L717"></a>NotNil(ci, t);

    <a id="L719"></a>var mi struct {
        <a id="L720"></a>x map[int]int;
    <a id="L721"></a>}
    <a id="L722"></a>Nil(mi, t);
    <a id="L723"></a>mi.x = make(map[int]int);
    <a id="L724"></a>NotNil(mi, t);

    <a id="L726"></a>var ii struct {
        <a id="L727"></a>x interface{};
    <a id="L728"></a>}
    <a id="L729"></a>Nil(ii, t);
    <a id="L730"></a>ii.x = 2;
    <a id="L731"></a>NotNil(ii, t);

    <a id="L733"></a>var fi struct {
        <a id="L734"></a>x func(t *testing.T);
    <a id="L735"></a>}
    <a id="L736"></a>Nil(fi, t);
    <a id="L737"></a>fi.x = TestIsNil;
    <a id="L738"></a>NotNil(fi, t);
<a id="L739"></a>}

<a id="L741"></a>func TestInterfaceExtraction(t *testing.T) {
    <a id="L742"></a>var s struct {
        <a id="L743"></a>w io.Writer;
    <a id="L744"></a>}

    <a id="L746"></a>s.w = os.Stdout;
    <a id="L747"></a>v := Indirect(NewValue(&amp;s)).(*StructValue).Field(0).Interface();
    <a id="L748"></a>if v != s.w.(interface{}) {
        <a id="L749"></a>t.Error(&#34;Interface() on interface: &#34;, v, s.w)
    <a id="L750"></a>}
<a id="L751"></a>}

<a id="L753"></a>func TestInterfaceEditing(t *testing.T) {
    <a id="L754"></a><span class="comment">// strings are bigger than one word,</span>
    <a id="L755"></a><span class="comment">// so the interface conversion allocates</span>
    <a id="L756"></a><span class="comment">// memory to hold a string and puts that</span>
    <a id="L757"></a><span class="comment">// pointer in the interface.</span>
    <a id="L758"></a>var i interface{} = &#34;hello&#34;;

    <a id="L760"></a><span class="comment">// if i pass the interface value by value</span>
    <a id="L761"></a><span class="comment">// to NewValue, i should get a fresh copy</span>
    <a id="L762"></a><span class="comment">// of the value.</span>
    <a id="L763"></a>v := NewValue(i);

    <a id="L765"></a><span class="comment">// and setting that copy to &#34;bye&#34; should</span>
    <a id="L766"></a><span class="comment">// not change the value stored in i.</span>
    <a id="L767"></a>v.(*StringValue).Set(&#34;bye&#34;);
    <a id="L768"></a>if i.(string) != &#34;hello&#34; {
        <a id="L769"></a>t.Errorf(`Set(&#34;bye&#34;) changed i to %s`, i.(string))
    <a id="L770"></a>}

    <a id="L772"></a><span class="comment">// the same should be true of smaller items.</span>
    <a id="L773"></a>i = 123;
    <a id="L774"></a>v = NewValue(i);
    <a id="L775"></a>v.(*IntValue).Set(234);
    <a id="L776"></a>if i.(int) != 123 {
        <a id="L777"></a>t.Errorf(&#34;Set(234) changed i to %d&#34;, i.(int))
    <a id="L778"></a>}
<a id="L779"></a>}

<a id="L781"></a>func TestNilPtrValueSub(t *testing.T) {
    <a id="L782"></a>var pi *int;
    <a id="L783"></a>if pv := NewValue(pi).(*PtrValue); pv.Elem() != nil {
        <a id="L784"></a>t.Error(&#34;NewValue((*int)(nil)).(*PtrValue).Elem() != nil&#34;)
    <a id="L785"></a>}
<a id="L786"></a>}

<a id="L788"></a>func TestMap(t *testing.T) {
    <a id="L789"></a>m := map[string]int{&#34;a&#34;: 1, &#34;b&#34;: 2};
    <a id="L790"></a>mv := NewValue(m).(*MapValue);
    <a id="L791"></a>if n := mv.Len(); n != len(m) {
        <a id="L792"></a>t.Errorf(&#34;Len = %d, want %d&#34;, n, len(m))
    <a id="L793"></a>}
    <a id="L794"></a>keys := mv.Keys();
    <a id="L795"></a>i := 0;
    <a id="L796"></a>newmap := MakeMap(mv.Type().(*MapType));
    <a id="L797"></a>for k, v := range m {
        <a id="L798"></a><span class="comment">// Check that returned Keys match keys in range.</span>
        <a id="L799"></a><span class="comment">// These aren&#39;t required to be in the same order,</span>
        <a id="L800"></a><span class="comment">// but they are in this implementation, which makes</span>
        <a id="L801"></a><span class="comment">// the test easier.</span>
        <a id="L802"></a>if i &gt;= len(keys) {
            <a id="L803"></a>t.Errorf(&#34;Missing key #%d %q&#34;, i, k)
        <a id="L804"></a>} else if kv := keys[i].(*StringValue); kv.Get() != k {
            <a id="L805"></a>t.Errorf(&#34;Keys[%d] = %q, want %q&#34;, i, kv.Get(), k)
        <a id="L806"></a>}
        <a id="L807"></a>i++;

        <a id="L809"></a><span class="comment">// Check that value lookup is correct.</span>
        <a id="L810"></a>vv := mv.Elem(NewValue(k));
        <a id="L811"></a>if vi := vv.(*IntValue).Get(); vi != v {
            <a id="L812"></a>t.Errorf(&#34;Key %q: have value %d, want %d&#34;, vi, v)
        <a id="L813"></a>}

        <a id="L815"></a><span class="comment">// Copy into new map.</span>
        <a id="L816"></a>newmap.SetElem(NewValue(k), NewValue(v));
    <a id="L817"></a>}
    <a id="L818"></a>vv := mv.Elem(NewValue(&#34;not-present&#34;));
    <a id="L819"></a>if vv != nil {
        <a id="L820"></a>t.Errorf(&#34;Invalid key: got non-nil value %s&#34;, valueToString(vv))
    <a id="L821"></a>}

    <a id="L823"></a>newm := newmap.Interface().(map[string]int);
    <a id="L824"></a>if len(newm) != len(m) {
        <a id="L825"></a>t.Errorf(&#34;length after copy: newm=%d, m=%d&#34;, newm, m)
    <a id="L826"></a>}

    <a id="L828"></a>for k, v := range newm {
        <a id="L829"></a>mv, ok := m[k];
        <a id="L830"></a>if mv != v {
            <a id="L831"></a>t.Errorf(&#34;newm[%q] = %d, but m[%q] = %d, %v&#34;, k, v, k, mv, ok)
        <a id="L832"></a>}
    <a id="L833"></a>}

    <a id="L835"></a>newmap.SetElem(NewValue(&#34;a&#34;), nil);
    <a id="L836"></a>v, ok := newm[&#34;a&#34;];
    <a id="L837"></a>if ok {
        <a id="L838"></a>t.Errorf(&#34;newm[\&#34;a\&#34;] = %d after delete&#34;, v)
    <a id="L839"></a>}
<a id="L840"></a>}

<a id="L842"></a>func TestChan(t *testing.T) {
    <a id="L843"></a>for loop := 0; loop &lt; 2; loop++ {
        <a id="L844"></a>var c chan int;
        <a id="L845"></a>var cv *ChanValue;

        <a id="L847"></a><span class="comment">// check both ways to allocate channels</span>
        <a id="L848"></a>switch loop {
        <a id="L849"></a>case 1:
            <a id="L850"></a>c = make(chan int, 1);
            <a id="L851"></a>cv = NewValue(c).(*ChanValue);
        <a id="L852"></a>case 0:
            <a id="L853"></a>cv = MakeChan(Typeof(c).(*ChanType), 1);
            <a id="L854"></a>c = cv.Interface().(chan int);
        <a id="L855"></a>}

        <a id="L857"></a><span class="comment">// Send</span>
        <a id="L858"></a>cv.Send(NewValue(2));
        <a id="L859"></a>if i := &lt;-c; i != 2 {
            <a id="L860"></a>t.Errorf(&#34;reflect Send 2, native recv %d&#34;, i)
        <a id="L861"></a>}

        <a id="L863"></a><span class="comment">// Recv</span>
        <a id="L864"></a>c &lt;- 3;
        <a id="L865"></a>if i := cv.Recv().(*IntValue).Get(); i != 3 {
            <a id="L866"></a>t.Errorf(&#34;native send 3, reflect Recv %d&#34;, i)
        <a id="L867"></a>}

        <a id="L869"></a><span class="comment">// TryRecv fail</span>
        <a id="L870"></a>val := cv.TryRecv();
        <a id="L871"></a>if val != nil {
            <a id="L872"></a>t.Errorf(&#34;TryRecv on empty chan: %s&#34;, valueToString(val))
        <a id="L873"></a>}

        <a id="L875"></a><span class="comment">// TryRecv success</span>
        <a id="L876"></a>c &lt;- 4;
        <a id="L877"></a>val = cv.TryRecv();
        <a id="L878"></a>if val == nil {
            <a id="L879"></a>t.Errorf(&#34;TryRecv on ready chan got nil&#34;)
        <a id="L880"></a>} else if i := val.(*IntValue).Get(); i != 4 {
            <a id="L881"></a>t.Errorf(&#34;native send 4, TryRecv %d&#34;, i)
        <a id="L882"></a>}

        <a id="L884"></a><span class="comment">// TrySend fail</span>
        <a id="L885"></a>c &lt;- 100;
        <a id="L886"></a>ok := cv.TrySend(NewValue(5));
        <a id="L887"></a>i := &lt;-c;
        <a id="L888"></a>if ok {
            <a id="L889"></a>t.Errorf(&#34;TrySend on full chan succeeded: value %d&#34;, i)
        <a id="L890"></a>}

        <a id="L892"></a><span class="comment">// TrySend success</span>
        <a id="L893"></a>ok = cv.TrySend(NewValue(6));
        <a id="L894"></a>if !ok {
            <a id="L895"></a>t.Errorf(&#34;TrySend on empty chan failed&#34;)
        <a id="L896"></a>} else {
            <a id="L897"></a>if i = &lt;-c; i != 6 {
                <a id="L898"></a>t.Errorf(&#34;TrySend 6, recv %d&#34;, i)
            <a id="L899"></a>}
        <a id="L900"></a>}

        <a id="L902"></a><span class="comment">// Close</span>
        <a id="L903"></a>c &lt;- 123;
        <a id="L904"></a>cv.Close();
        <a id="L905"></a>if cv.Closed() {
            <a id="L906"></a>t.Errorf(&#34;closed too soon - 1&#34;)
        <a id="L907"></a>}
        <a id="L908"></a>if i := cv.Recv().(*IntValue).Get(); i != 123 {
            <a id="L909"></a>t.Errorf(&#34;send 123 then close; Recv %d&#34;, i)
        <a id="L910"></a>}
        <a id="L911"></a>if cv.Closed() {
            <a id="L912"></a>t.Errorf(&#34;closed too soon - 2&#34;)
        <a id="L913"></a>}
        <a id="L914"></a>if i := cv.Recv().(*IntValue).Get(); i != 0 {
            <a id="L915"></a>t.Errorf(&#34;after close Recv %d&#34;, i)
        <a id="L916"></a>}
        <a id="L917"></a>if !cv.Closed() {
            <a id="L918"></a>t.Errorf(&#34;not closed&#34;)
        <a id="L919"></a>}
    <a id="L920"></a>}

    <a id="L922"></a><span class="comment">// check creation of unbuffered channel</span>
    <a id="L923"></a>var c chan int;
    <a id="L924"></a>cv := MakeChan(Typeof(c).(*ChanType), 0);
    <a id="L925"></a>c = cv.Interface().(chan int);
    <a id="L926"></a>if cv.TrySend(NewValue(7)) {
        <a id="L927"></a>t.Errorf(&#34;TrySend on sync chan succeeded&#34;)
    <a id="L928"></a>}
    <a id="L929"></a>if cv.TryRecv() != nil {
        <a id="L930"></a>t.Errorf(&#34;TryRecv on sync chan succeeded&#34;)
    <a id="L931"></a>}

    <a id="L933"></a><span class="comment">// len/cap</span>
    <a id="L934"></a>cv = MakeChan(Typeof(c).(*ChanType), 10);
    <a id="L935"></a>c = cv.Interface().(chan int);
    <a id="L936"></a>for i := 0; i &lt; 3; i++ {
        <a id="L937"></a>c &lt;- i
    <a id="L938"></a>}
    <a id="L939"></a>if l, m := cv.Len(), cv.Cap(); l != len(c) || m != cap(c) {
        <a id="L940"></a>t.Errorf(&#34;Len/Cap = %d/%d want %d/%d&#34;, l, m, len(c), cap(c))
    <a id="L941"></a>}

<a id="L943"></a>}

<a id="L945"></a><span class="comment">// Difficult test for function call because of</span>
<a id="L946"></a><span class="comment">// implicit padding between arguments.</span>
<a id="L947"></a>func dummy(b byte, c int, d byte) (i byte, j int, k byte) {
    <a id="L948"></a>return b, c, d
<a id="L949"></a>}

<a id="L951"></a>func TestFunc(t *testing.T) {
    <a id="L952"></a>ret := NewValue(dummy).(*FuncValue).Call([]Value{NewValue(byte(10)), NewValue(20), NewValue(byte(30))});
    <a id="L953"></a>if len(ret) != 3 {
        <a id="L954"></a>t.Fatalf(&#34;Call returned %d values, want 3&#34;, len(ret))
    <a id="L955"></a>}

    <a id="L957"></a>i := ret[0].(*Uint8Value).Get();
    <a id="L958"></a>j := ret[1].(*IntValue).Get();
    <a id="L959"></a>k := ret[2].(*Uint8Value).Get();
    <a id="L960"></a>if i != 10 || j != 20 || k != 30 {
        <a id="L961"></a>t.Errorf(&#34;Call returned %d, %d, %d; want 10, 20, 30&#34;, i, j, k)
    <a id="L962"></a>}
<a id="L963"></a>}

<a id="L965"></a>type Point struct {
    <a id="L966"></a>x, y int;
<a id="L967"></a>}

<a id="L969"></a>func (p Point) Dist(scale int) int { return p.x*p.x*scale + p.y*p.y*scale }

<a id="L971"></a>func TestMethod(t *testing.T) {
    <a id="L972"></a><span class="comment">// Non-curried method of type.</span>
    <a id="L973"></a>p := Point{3, 4};
    <a id="L974"></a>i := Typeof(p).Method(0).Func.Call([]Value{NewValue(p), NewValue(10)})[0].(*IntValue).Get();
    <a id="L975"></a>if i != 250 {
        <a id="L976"></a>t.Errorf(&#34;Type Method returned %d; want 250&#34;, i)
    <a id="L977"></a>}

    <a id="L979"></a><span class="comment">// Curried method of value.</span>
    <a id="L980"></a>i = NewValue(p).Method(0).Call([]Value{NewValue(10)})[0].(*IntValue).Get();
    <a id="L981"></a>if i != 250 {
        <a id="L982"></a>t.Errorf(&#34;Value Method returned %d; want 250&#34;, i)
    <a id="L983"></a>}

    <a id="L985"></a><span class="comment">// Curried method of interface value.</span>
    <a id="L986"></a><span class="comment">// Have to wrap interface value in a struct to get at it.</span>
    <a id="L987"></a><span class="comment">// Passing it to NewValue directly would</span>
    <a id="L988"></a><span class="comment">// access the underlying Point, not the interface.</span>
    <a id="L989"></a>var s = struct {
        <a id="L990"></a>x interface {
            <a id="L991"></a>Dist(int) int;
        <a id="L992"></a>};
    <a id="L993"></a>}{p};
    <a id="L994"></a>pv := NewValue(s).(*StructValue).Field(0);
    <a id="L995"></a>i = pv.Method(0).Call([]Value{NewValue(10)})[0].(*IntValue).Get();
    <a id="L996"></a>if i != 250 {
        <a id="L997"></a>t.Errorf(&#34;Interface Method returned %d; want 250&#34;, i)
    <a id="L998"></a>}
<a id="L999"></a>}

<a id="L1001"></a>func TestInterfaceSet(t *testing.T) {
    <a id="L1002"></a>p := &amp;Point{3, 4};

    <a id="L1004"></a>var s struct {
        <a id="L1005"></a>I   interface{};
        <a id="L1006"></a>P   interface {
            <a id="L1007"></a>Dist(int) int;
        <a id="L1008"></a>};
    <a id="L1009"></a>}
    <a id="L1010"></a>sv := NewValue(&amp;s).(*PtrValue).Elem().(*StructValue);
    <a id="L1011"></a>sv.Field(0).(*InterfaceValue).Set(NewValue(p));
    <a id="L1012"></a>if q := s.I.(*Point); q != p {
        <a id="L1013"></a>t.Errorf(&#34;i: have %p want %p&#34;, q, p)
    <a id="L1014"></a>}

    <a id="L1016"></a>pv := sv.Field(1).(*InterfaceValue);
    <a id="L1017"></a>pv.Set(NewValue(p));
    <a id="L1018"></a>if q := s.P.(*Point); q != p {
        <a id="L1019"></a>t.Errorf(&#34;i: have %p want %p&#34;, q, p)
    <a id="L1020"></a>}

    <a id="L1022"></a>i := pv.Method(0).Call([]Value{NewValue(10)})[0].(*IntValue).Get();
    <a id="L1023"></a>if i != 250 {
        <a id="L1024"></a>t.Errorf(&#34;Interface Method returned %d; want 250&#34;, i)
    <a id="L1025"></a>}
<a id="L1026"></a>}

<a id="L1028"></a>type T1 struct {
    <a id="L1029"></a>a   string;
    <a id="L1030"></a>int;
<a id="L1031"></a>}

<a id="L1033"></a>func TestAnonymousFields(t *testing.T) {
    <a id="L1034"></a>var field StructField;
    <a id="L1035"></a>var ok bool;
    <a id="L1036"></a>var t1 T1;
    <a id="L1037"></a>type1 := Typeof(t1).(*StructType);
    <a id="L1038"></a>if field, ok = type1.FieldByName(&#34;int&#34;); !ok {
        <a id="L1039"></a>t.Error(&#34;no field &#39;int&#39;&#34;)
    <a id="L1040"></a>}
    <a id="L1041"></a>if field.Index[0] != 1 {
        <a id="L1042"></a>t.Error(&#34;field index should be 1; is&#34;, field.Index)
    <a id="L1043"></a>}
<a id="L1044"></a>}

<a id="L1046"></a>type FTest struct {
    <a id="L1047"></a>s     interface{};
    <a id="L1048"></a>name  string;
    <a id="L1049"></a>index []int;
    <a id="L1050"></a>value int;
<a id="L1051"></a>}

<a id="L1053"></a>type D1 struct {
    <a id="L1054"></a>d int;
<a id="L1055"></a>}
<a id="L1056"></a>type D2 struct {
    <a id="L1057"></a>d int;
<a id="L1058"></a>}

<a id="L1060"></a>type S0 struct {
    <a id="L1061"></a>a, b, c int;
    <a id="L1062"></a>D1;
    <a id="L1063"></a>D2;
<a id="L1064"></a>}

<a id="L1066"></a>type S1 struct {
    <a id="L1067"></a>b   int;
    <a id="L1068"></a>S0;
<a id="L1069"></a>}

<a id="L1071"></a>type S2 struct {
    <a id="L1072"></a>a   int;
    <a id="L1073"></a>*S1;
<a id="L1074"></a>}

<a id="L1076"></a>type S1x struct {
    <a id="L1077"></a>S1;
<a id="L1078"></a>}

<a id="L1080"></a>type S1y struct {
    <a id="L1081"></a>S1;
<a id="L1082"></a>}

<a id="L1084"></a>type S3 struct {
    <a id="L1085"></a>S1x;
    <a id="L1086"></a>S2;
    <a id="L1087"></a>d, e int;
    <a id="L1088"></a>*S1y;
<a id="L1089"></a>}

<a id="L1091"></a>type S4 struct {
    <a id="L1092"></a>*S4;
    <a id="L1093"></a>a   int;
<a id="L1094"></a>}

<a id="L1096"></a>var fieldTests = []FTest{
    <a id="L1097"></a>FTest{struct{}{}, &#34;&#34;, nil, 0},
    <a id="L1098"></a>FTest{struct{}{}, &#34;foo&#34;, nil, 0},
    <a id="L1099"></a>FTest{S0{a: &#39;a&#39;}, &#34;a&#34;, []int{0}, &#39;a&#39;},
    <a id="L1100"></a>FTest{S0{}, &#34;d&#34;, nil, 0},
    <a id="L1101"></a>FTest{S1{S0: S0{a: &#39;a&#39;}}, &#34;a&#34;, []int{1, 0}, &#39;a&#39;},
    <a id="L1102"></a>FTest{S1{b: &#39;b&#39;}, &#34;b&#34;, []int{0}, &#39;b&#39;},
    <a id="L1103"></a>FTest{S1{}, &#34;S0&#34;, []int{1}, 0},
    <a id="L1104"></a>FTest{S1{S0: S0{c: &#39;c&#39;}}, &#34;c&#34;, []int{1, 2}, &#39;c&#39;},
    <a id="L1105"></a>FTest{S2{a: &#39;a&#39;}, &#34;a&#34;, []int{0}, &#39;a&#39;},
    <a id="L1106"></a>FTest{S2{}, &#34;S1&#34;, []int{1}, 0},
    <a id="L1107"></a>FTest{S2{S1: &amp;S1{b: &#39;b&#39;}}, &#34;b&#34;, []int{1, 0}, &#39;b&#39;},
    <a id="L1108"></a>FTest{S2{S1: &amp;S1{S0: S0{c: &#39;c&#39;}}}, &#34;c&#34;, []int{1, 1, 2}, &#39;c&#39;},
    <a id="L1109"></a>FTest{S2{}, &#34;d&#34;, nil, 0},
    <a id="L1110"></a>FTest{S3{}, &#34;S1&#34;, nil, 0},
    <a id="L1111"></a>FTest{S3{S2: S2{a: &#39;a&#39;}}, &#34;a&#34;, []int{1, 0}, &#39;a&#39;},
    <a id="L1112"></a>FTest{S3{}, &#34;b&#34;, nil, 0},
    <a id="L1113"></a>FTest{S3{d: &#39;d&#39;}, &#34;d&#34;, []int{2}, 0},
    <a id="L1114"></a>FTest{S3{e: &#39;e&#39;}, &#34;e&#34;, []int{3}, &#39;e&#39;},
    <a id="L1115"></a>FTest{S4{a: &#39;a&#39;}, &#34;a&#34;, []int{1}, &#39;a&#39;},
    <a id="L1116"></a>FTest{S4{}, &#34;b&#34;, nil, 0},
<a id="L1117"></a>}

<a id="L1119"></a>func TestFieldByIndex(t *testing.T) {
    <a id="L1120"></a>for _, test := range fieldTests {
        <a id="L1121"></a>s := Typeof(test.s).(*StructType);
        <a id="L1122"></a>f := s.FieldByIndex(test.index);
        <a id="L1123"></a>if f.Name != &#34;&#34; {
            <a id="L1124"></a>if test.index != nil {
                <a id="L1125"></a>if f.Name != test.name {
                    <a id="L1126"></a>t.Errorf(&#34;%s.%s found; want %s&#34;, s.Name(), f.Name, test.name)
                <a id="L1127"></a>}
            <a id="L1128"></a>} else {
                <a id="L1129"></a>t.Errorf(&#34;%s.%s found&#34;, s.Name(), f.Name)
            <a id="L1130"></a>}
        <a id="L1131"></a>} else if len(test.index) &gt; 0 {
            <a id="L1132"></a>t.Errorf(&#34;%s.%s not found&#34;, s.Name(), test.name)
        <a id="L1133"></a>}

        <a id="L1135"></a>if test.value != 0 {
            <a id="L1136"></a>v := NewValue(test.s).(*StructValue).FieldByIndex(test.index);
            <a id="L1137"></a>if v != nil {
                <a id="L1138"></a>if x, ok := v.Interface().(int); ok {
                    <a id="L1139"></a>if x != test.value {
                        <a id="L1140"></a>t.Errorf(&#34;%s%v is %d; want %d&#34;, s.Name(), test.index, x, test.value)
                    <a id="L1141"></a>}
                <a id="L1142"></a>} else {
                    <a id="L1143"></a>t.Errorf(&#34;%s%v value not an int&#34;, s.Name(), test.index)
                <a id="L1144"></a>}
            <a id="L1145"></a>} else {
                <a id="L1146"></a>t.Errorf(&#34;%s%v value not found&#34;, s.Name(), test.index)
            <a id="L1147"></a>}
        <a id="L1148"></a>}
    <a id="L1149"></a>}
<a id="L1150"></a>}

<a id="L1152"></a>func TestFieldByName(t *testing.T) {
    <a id="L1153"></a>for _, test := range fieldTests {
        <a id="L1154"></a>s := Typeof(test.s).(*StructType);
        <a id="L1155"></a>f, found := s.FieldByName(test.name);
        <a id="L1156"></a>if found {
            <a id="L1157"></a>if test.index != nil {
                <a id="L1158"></a><span class="comment">// Verify field depth and index.</span>
                <a id="L1159"></a>if len(f.Index) != len(test.index) {
                    <a id="L1160"></a>t.Errorf(&#34;%s.%s depth %d; want %d&#34;, s.Name(), test.name, len(f.Index), len(test.index))
                <a id="L1161"></a>} else {
                    <a id="L1162"></a>for i, x := range f.Index {
                        <a id="L1163"></a>if x != test.index[i] {
                            <a id="L1164"></a>t.Errorf(&#34;%s.%s.Index[%d] is %d; want %d&#34;, s.Name(), test.name, i, x, test.index[i])
                        <a id="L1165"></a>}
                    <a id="L1166"></a>}
                <a id="L1167"></a>}
            <a id="L1168"></a>} else {
                <a id="L1169"></a>t.Errorf(&#34;%s.%s found&#34;, s.Name(), f.Name)
            <a id="L1170"></a>}
        <a id="L1171"></a>} else if len(test.index) &gt; 0 {
            <a id="L1172"></a>t.Errorf(&#34;%s.%s not found&#34;, s.Name(), test.name)
        <a id="L1173"></a>}

        <a id="L1175"></a>if test.value != 0 {
            <a id="L1176"></a>v := NewValue(test.s).(*StructValue).FieldByName(test.name);
            <a id="L1177"></a>if v != nil {
                <a id="L1178"></a>if x, ok := v.Interface().(int); ok {
                    <a id="L1179"></a>if x != test.value {
                        <a id="L1180"></a>t.Errorf(&#34;%s.%s is %d; want %d&#34;, s.Name(), test.name, x, test.value)
                    <a id="L1181"></a>}
                <a id="L1182"></a>} else {
                    <a id="L1183"></a>t.Errorf(&#34;%s.%s value not an int&#34;, s.Name(), test.name)
                <a id="L1184"></a>}
            <a id="L1185"></a>} else {
                <a id="L1186"></a>t.Errorf(&#34;%s.%s value not found&#34;, s.Name(), test.name)
            <a id="L1187"></a>}
        <a id="L1188"></a>}
    <a id="L1189"></a>}
<a id="L1190"></a>}
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
