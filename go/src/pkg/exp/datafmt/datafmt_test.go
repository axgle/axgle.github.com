<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/exp/datafmt/datafmt_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/exp/datafmt/datafmt_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package datafmt

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;strings&#34;;
    <a id="L10"></a>&#34;testing&#34;;
<a id="L11"></a>)


<a id="L14"></a>func parse(t *testing.T, form string, fmap FormatterMap) Format {
    <a id="L15"></a>f, err := Parse(&#34;&#34;, strings.Bytes(form), fmap);
    <a id="L16"></a>if err != nil {
        <a id="L17"></a>t.Errorf(&#34;Parse(%s): %v&#34;, form, err);
        <a id="L18"></a>return nil;
    <a id="L19"></a>}
    <a id="L20"></a>return f;
<a id="L21"></a>}


<a id="L24"></a>func verify(t *testing.T, f Format, expected string, args ...) {
    <a id="L25"></a>if f == nil {
        <a id="L26"></a>return <span class="comment">// allow other tests to run</span>
    <a id="L27"></a>}
    <a id="L28"></a>result := f.Sprint(args);
    <a id="L29"></a>if result != expected {
        <a id="L30"></a>t.Errorf(
            <a id="L31"></a>&#34;result  : `%s`\nexpected: `%s`\n\n&#34;,
            <a id="L32"></a>result, expected)
    <a id="L33"></a>}
<a id="L34"></a>}


<a id="L37"></a>func formatter(s *State, value interface{}, rule_name string) bool {
    <a id="L38"></a>switch rule_name {
    <a id="L39"></a>case &#34;/&#34;:
        <a id="L40"></a>fmt.Fprintf(s, &#34;%d %d %d&#34;, s.Pos().Line, s.LinePos().Column, s.Pos().Column);
        <a id="L41"></a>return true;
    <a id="L42"></a>case &#34;blank&#34;:
        <a id="L43"></a>s.Write([]byte{&#39; &#39;});
        <a id="L44"></a>return true;
    <a id="L45"></a>case &#34;int&#34;:
        <a id="L46"></a>if value.(int)&amp;1 == 0 {
            <a id="L47"></a>fmt.Fprint(s, &#34;even &#34;)
        <a id="L48"></a>} else {
            <a id="L49"></a>fmt.Fprint(s, &#34;odd &#34;)
        <a id="L50"></a>}
        <a id="L51"></a>return true;
    <a id="L52"></a>case &#34;nil&#34;:
        <a id="L53"></a>return false
    <a id="L54"></a>case &#34;testing.T&#34;:
        <a id="L55"></a>s.Write(strings.Bytes(&#34;testing.T&#34;));
        <a id="L56"></a>return true;
    <a id="L57"></a>}
    <a id="L58"></a>panic(&#34;unreachable&#34;);
    <a id="L59"></a>return false;
<a id="L60"></a>}


<a id="L63"></a>func TestCustomFormatters(t *testing.T) {
    <a id="L64"></a>fmap0 := FormatterMap{&#34;/&#34;: formatter};
    <a id="L65"></a>fmap1 := FormatterMap{&#34;int&#34;: formatter, &#34;blank&#34;: formatter, &#34;nil&#34;: formatter};
    <a id="L66"></a>fmap2 := FormatterMap{&#34;testing.T&#34;: formatter};

    <a id="L68"></a>f := parse(t, `int=`, fmap0);
    <a id="L69"></a>verify(t, f, ``, 1, 2, 3);

    <a id="L71"></a>f = parse(t, `int=&#34;#&#34;`, nil);
    <a id="L72"></a>verify(t, f, `###`, 1, 2, 3);

    <a id="L74"></a>f = parse(t, `int=&#34;#&#34;;string=&#34;%s&#34;`, fmap0);
    <a id="L75"></a>verify(t, f, &#34;#1 0 1#1 0 7#1 0 13\n2 0 0foo2 0 8\n&#34;, 1, 2, 3, &#34;\n&#34;, &#34;foo&#34;, &#34;\n&#34;);

    <a id="L77"></a>f = parse(t, ``, fmap1);
    <a id="L78"></a>verify(t, f, `even odd even odd `, 0, 1, 2, 3);

    <a id="L80"></a>f = parse(t, `/ =@:blank; float=&#34;#&#34;`, fmap1);
    <a id="L81"></a>verify(t, f, `# # #`, 0.0, 1.0, 2.0);

    <a id="L83"></a>f = parse(t, `float=@:nil`, fmap1);
    <a id="L84"></a>verify(t, f, ``, 0.0, 1.0, 2.0);

    <a id="L86"></a>f = parse(t, `testing &#34;testing&#34;; ptr=*`, fmap2);
    <a id="L87"></a>verify(t, f, `testing.T`, t);

    <a id="L89"></a><span class="comment">// TODO needs more tests</span>
<a id="L90"></a>}


<a id="L93"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L94"></a><span class="comment">// Formatting of basic and simple composite types</span>

<a id="L96"></a>func check(t *testing.T, form, expected string, args ...) {
    <a id="L97"></a>f := parse(t, form, nil);
    <a id="L98"></a>if f == nil {
        <a id="L99"></a>return <span class="comment">// allow other tests to run</span>
    <a id="L100"></a>}
    <a id="L101"></a>result := f.Sprint(args);
    <a id="L102"></a>if result != expected {
        <a id="L103"></a>t.Errorf(
            <a id="L104"></a>&#34;format  : %s\nresult  : `%s`\nexpected: `%s`\n\n&#34;,
            <a id="L105"></a>form, result, expected)
    <a id="L106"></a>}
<a id="L107"></a>}


<a id="L110"></a>func TestBasicTypes(t *testing.T) {
    <a id="L111"></a>check(t, ``, ``);
    <a id="L112"></a>check(t, `bool=&#34;:%v&#34;`, `:true:false`, true, false);
    <a id="L113"></a>check(t, `int=&#34;%b %d %o 0x%x&#34;`, `101010 42 52 0x2a`, 42);

    <a id="L115"></a>check(t, `int=&#34;%&#34;`, `%`, 42);
    <a id="L116"></a>check(t, `int=&#34;%%&#34;`, `%`, 42);
    <a id="L117"></a>check(t, `int=&#34;**%%**&#34;`, `**%**`, 42);
    <a id="L118"></a>check(t, `int=&#34;%%%%%%&#34;`, `%%%`, 42);
    <a id="L119"></a>check(t, `int=&#34;%%%d%%&#34;`, `%42%`, 42);

    <a id="L121"></a>const i = -42;
    <a id="L122"></a>const is = `-42`;
    <a id="L123"></a>check(t, `int  =&#34;%d&#34;`, is, i);
    <a id="L124"></a>check(t, `int8 =&#34;%d&#34;`, is, int8(i));
    <a id="L125"></a>check(t, `int16=&#34;%d&#34;`, is, int16(i));
    <a id="L126"></a>check(t, `int32=&#34;%d&#34;`, is, int32(i));
    <a id="L127"></a>check(t, `int64=&#34;%d&#34;`, is, int64(i));

    <a id="L129"></a>const u = 42;
    <a id="L130"></a>const us = `42`;
    <a id="L131"></a>check(t, `uint  =&#34;%d&#34;`, us, uint(u));
    <a id="L132"></a>check(t, `uint8 =&#34;%d&#34;`, us, uint8(u));
    <a id="L133"></a>check(t, `uint16=&#34;%d&#34;`, us, uint16(u));
    <a id="L134"></a>check(t, `uint32=&#34;%d&#34;`, us, uint32(u));
    <a id="L135"></a>check(t, `uint64=&#34;%d&#34;`, us, uint64(u));

    <a id="L137"></a>const f = 3.141592;
    <a id="L138"></a>const fs = `3.141592`;
    <a id="L139"></a>check(t, `float  =&#34;%g&#34;`, fs, f);
    <a id="L140"></a>check(t, `float32=&#34;%g&#34;`, fs, float32(f));
    <a id="L141"></a>check(t, `float64=&#34;%g&#34;`, fs, float64(f));
<a id="L142"></a>}


<a id="L145"></a>func TestArrayTypes(t *testing.T) {
    <a id="L146"></a>var a0 [10]int;
    <a id="L147"></a>check(t, `array=&#34;array&#34;;`, `array`, a0);

    <a id="L149"></a>a1 := [...]int{1, 2, 3};
    <a id="L150"></a>check(t, `array=&#34;array&#34;;`, `array`, a1);
    <a id="L151"></a>check(t, `array={*}; int=&#34;%d&#34;;`, `123`, a1);
    <a id="L152"></a>check(t, `array={* / &#34;, &#34;}; int=&#34;%d&#34;;`, `1, 2, 3`, a1);
    <a id="L153"></a>check(t, `array={* / *}; int=&#34;%d&#34;;`, `12233`, a1);

    <a id="L155"></a>a2 := []interface{}{42, &#34;foo&#34;, 3.14};
    <a id="L156"></a>check(t, `array={* / &#34;, &#34;}; interface=*; string=&#34;bar&#34;; default=&#34;%v&#34;;`, `42, bar, 3.14`, a2);
<a id="L157"></a>}


<a id="L160"></a>func TestChanTypes(t *testing.T) {
    <a id="L161"></a>var c0 chan int;
    <a id="L162"></a>check(t, `chan=&#34;chan&#34;`, `chan`, c0);

    <a id="L164"></a>c1 := make(chan int);
    <a id="L165"></a>go func() { c1 &lt;- 42 }();
    <a id="L166"></a>check(t, `chan=&#34;chan&#34;`, `chan`, c1);
    <a id="L167"></a><span class="comment">// check(t, `chan=*`, `42`, c1);  // reflection support for chans incomplete</span>
<a id="L168"></a>}


<a id="L171"></a>func TestFuncTypes(t *testing.T) {
    <a id="L172"></a>var f0 func() int;
    <a id="L173"></a>check(t, `func=&#34;func&#34;`, `func`, f0);

    <a id="L175"></a>f1 := func() int { return 42 };
    <a id="L176"></a>check(t, `func=&#34;func&#34;`, `func`, f1);
    <a id="L177"></a><span class="comment">// check(t, `func=*`, `42`, f1);  // reflection support for funcs incomplete</span>
<a id="L178"></a>}


<a id="L181"></a>func TestInterfaceTypes(t *testing.T) {
    <a id="L182"></a>var i0 interface{}
    <a id="L183"></a>check(t, `interface=&#34;interface&#34;`, `interface`, i0);

    <a id="L185"></a>i0 = &#34;foo&#34;;
    <a id="L186"></a>check(t, `interface=&#34;interface&#34;`, `interface`, i0);
    <a id="L187"></a>check(t, `interface=*; string=&#34;%s&#34;`, `foo`, i0);
<a id="L188"></a>}


<a id="L191"></a>func TestMapTypes(t *testing.T) {
    <a id="L192"></a>var m0 map[string]int;
    <a id="L193"></a>check(t, `map=&#34;map&#34;`, `map`, m0);

    <a id="L195"></a>m1 := map[string]int{};
    <a id="L196"></a>check(t, `map=&#34;map&#34;`, `map`, m1);
    <a id="L197"></a><span class="comment">// check(t, `map=*`, ``, m1);  // reflection support for maps incomplete</span>
<a id="L198"></a>}


<a id="L201"></a>func TestPointerTypes(t *testing.T) {
    <a id="L202"></a>var p0 *int;
    <a id="L203"></a>check(t, `ptr=&#34;ptr&#34;`, `ptr`, p0);
    <a id="L204"></a>check(t, `ptr=*`, ``, p0);
    <a id="L205"></a>check(t, `ptr=*|&#34;nil&#34;`, `nil`, p0);

    <a id="L207"></a>x := 99991;
    <a id="L208"></a>p1 := &amp;x;
    <a id="L209"></a>check(t, `ptr=&#34;ptr&#34;`, `ptr`, p1);
    <a id="L210"></a>check(t, `ptr=*; int=&#34;%d&#34;`, `99991`, p1);
<a id="L211"></a>}


<a id="L214"></a>func TestDefaultRule(t *testing.T) {
    <a id="L215"></a>check(t, `default=&#34;%v&#34;`, `42foo3.14`, 42, &#34;foo&#34;, 3.14);
    <a id="L216"></a>check(t, `default=&#34;%v&#34;; int=&#34;%x&#34;`, `abcdef`, 10, 11, 12, 13, 14, 15);
    <a id="L217"></a>check(t, `default=&#34;%v&#34;; int=&#34;%x&#34;`, `ab**ef`, 10, 11, &#34;**&#34;, 14, 15);
    <a id="L218"></a>check(t, `default=&#34;%x&#34;; int=@:default`, `abcdef`, 10, 11, 12, 13, 14, 15);
<a id="L219"></a>}


<a id="L222"></a>func TestGlobalSeparatorRule(t *testing.T) {
    <a id="L223"></a>check(t, `int=&#34;%d&#34;; / =&#34;-&#34;`, `1-2-3-4`, 1, 2, 3, 4);
    <a id="L224"></a>check(t, `int=&#34;%x%x&#34;; / =&#34;*&#34;`, `aa*aa`, 10, 10);
<a id="L225"></a>}


<a id="L228"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L229"></a><span class="comment">// Formatting of a struct</span>

<a id="L231"></a>type T1 struct {
    <a id="L232"></a>a int;
<a id="L233"></a>}

<a id="L235"></a>const F1 = `datafmt &#34;datafmt&#34;;`
    <a id="L236"></a>`int = &#34;%d&#34;;`
    <a id="L237"></a>`datafmt.T1 = &#34;&lt;&#34; a &#34;&gt;&#34;;`

<a id="L239"></a>func TestStruct1(t *testing.T) { check(t, F1, &#34;&lt;42&gt;&#34;, T1{42}) }


<a id="L242"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L243"></a><span class="comment">// Formatting of a struct with an optional field (ptr)</span>

<a id="L245"></a>type T2 struct {
    <a id="L246"></a>s   string;
    <a id="L247"></a>p   *T1;
<a id="L248"></a>}

<a id="L250"></a>const F2a = F1 +
    <a id="L251"></a>`string = &#34;%s&#34;;`
        <a id="L252"></a>`ptr = *;`
        <a id="L253"></a>`datafmt.T2 = s [&#34;-&#34; p &#34;-&#34;];`

<a id="L255"></a>const F2b = F1 +
    <a id="L256"></a>`string = &#34;%s&#34;;`
        <a id="L257"></a>`ptr = *;`
        <a id="L258"></a>`datafmt.T2 = s (&#34;-&#34; p &#34;-&#34; | &#34;empty&#34;);`

<a id="L260"></a>func TestStruct2(t *testing.T) {
    <a id="L261"></a>check(t, F2a, &#34;foo&#34;, T2{&#34;foo&#34;, nil});
    <a id="L262"></a>check(t, F2a, &#34;bar-&lt;17&gt;-&#34;, T2{&#34;bar&#34;, &amp;T1{17}});
    <a id="L263"></a>check(t, F2b, &#34;fooempty&#34;, T2{&#34;foo&#34;, nil});
<a id="L264"></a>}


<a id="L267"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L268"></a><span class="comment">// Formatting of a struct with a repetitive field (slice)</span>

<a id="L270"></a>type T3 struct {
    <a id="L271"></a>s   string;
    <a id="L272"></a>a   []int;
<a id="L273"></a>}

<a id="L275"></a>const F3a = `datafmt &#34;datafmt&#34;;`
    <a id="L276"></a>`default = &#34;%v&#34;;`
    <a id="L277"></a>`array = *;`
    <a id="L278"></a>`datafmt.T3 = s  {&#34; &#34; a a / &#34;,&#34;};`

<a id="L280"></a>const F3b = `datafmt &#34;datafmt&#34;;`
    <a id="L281"></a>`int = &#34;%d&#34;;`
    <a id="L282"></a>`string = &#34;%s&#34;;`
    <a id="L283"></a>`array = *;`
    <a id="L284"></a>`nil = ;`
    <a id="L285"></a>`empty = *:nil;`
    <a id="L286"></a>`datafmt.T3 = s [a:empty &#34;: &#34; {a / &#34;-&#34;}]`

<a id="L288"></a>func TestStruct3(t *testing.T) {
    <a id="L289"></a>check(t, F3a, &#34;foo&#34;, T3{&#34;foo&#34;, nil});
    <a id="L290"></a>check(t, F3a, &#34;foo 00, 11, 22&#34;, T3{&#34;foo&#34;, []int{0, 1, 2}});
    <a id="L291"></a>check(t, F3b, &#34;bar&#34;, T3{&#34;bar&#34;, nil});
    <a id="L292"></a>check(t, F3b, &#34;bal: 2-3-5&#34;, T3{&#34;bal&#34;, []int{2, 3, 5}});
<a id="L293"></a>}


<a id="L296"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L297"></a><span class="comment">// Formatting of a struct with alternative field</span>

<a id="L299"></a>type T4 struct {
    <a id="L300"></a>x   *int;
    <a id="L301"></a>a   []int;
<a id="L302"></a>}

<a id="L304"></a>const F4a = `datafmt &#34;datafmt&#34;;`
    <a id="L305"></a>`int = &#34;%d&#34;;`
    <a id="L306"></a>`ptr = *;`
    <a id="L307"></a>`array = *;`
    <a id="L308"></a>`nil = ;`
    <a id="L309"></a>`empty = *:nil;`
    <a id="L310"></a>`datafmt.T4 = &#34;&lt;&#34; (x:empty x | &#34;-&#34;) &#34;&gt;&#34; `

<a id="L312"></a>const F4b = `datafmt &#34;datafmt&#34;;`
    <a id="L313"></a>`int = &#34;%d&#34;;`
    <a id="L314"></a>`ptr = *;`
    <a id="L315"></a>`array = *;`
    <a id="L316"></a>`nil = ;`
    <a id="L317"></a>`empty = *:nil;`
    <a id="L318"></a>`datafmt.T4 = &#34;&lt;&#34; (a:empty {a / &#34;, &#34;} | &#34;-&#34;) &#34;&gt;&#34; `

<a id="L320"></a>func TestStruct4(t *testing.T) {
    <a id="L321"></a>x := 7;
    <a id="L322"></a>check(t, F4a, &#34;&lt;-&gt;&#34;, T4{nil, nil});
    <a id="L323"></a>check(t, F4a, &#34;&lt;7&gt;&#34;, T4{&amp;x, nil});
    <a id="L324"></a>check(t, F4b, &#34;&lt;-&gt;&#34;, T4{nil, nil});
    <a id="L325"></a>check(t, F4b, &#34;&lt;2, 3, 7&gt;&#34;, T4{nil, []int{2, 3, 7}});
<a id="L326"></a>}


<a id="L329"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L330"></a><span class="comment">// Formatting a struct (documentation example)</span>

<a id="L332"></a>type Point struct {
    <a id="L333"></a>name string;
    <a id="L334"></a>x, y int;
<a id="L335"></a>}

<a id="L337"></a>const FPoint = `datafmt &#34;datafmt&#34;;`
    <a id="L338"></a>`int = &#34;%d&#34;;`
    <a id="L339"></a>`hexInt = &#34;0x%x&#34;;`
    <a id="L340"></a>`string = &#34;---%s---&#34;;`
    <a id="L341"></a>`datafmt.Point = name &#34;{&#34; x &#34;, &#34; y:hexInt &#34;}&#34;;`

<a id="L343"></a>func TestStructPoint(t *testing.T) {
    <a id="L344"></a>p := Point{&#34;foo&#34;, 3, 15};
    <a id="L345"></a>check(t, FPoint, &#34;---foo---{3, 0xf}&#34;, p);
<a id="L346"></a>}


<a id="L349"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L350"></a><span class="comment">// Formatting a slice (documentation example)</span>

<a id="L352"></a>const FSlice = `int = &#34;%b&#34;;`
    <a id="L353"></a>`array = { * / &#34;, &#34; }`

<a id="L355"></a>func TestSlice(t *testing.T) { check(t, FSlice, &#34;10, 11, 101, 111&#34;, []int{2, 3, 5, 7}) }


<a id="L358"></a><span class="comment">// TODO add more tests</span>
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
