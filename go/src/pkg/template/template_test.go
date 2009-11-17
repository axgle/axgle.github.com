<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/template/template_test.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/template/template_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package template

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;container/vector&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;testing&#34;;
<a id="L13"></a>)

<a id="L15"></a>type Test struct {
    <a id="L16"></a>in, out, err string;
<a id="L17"></a>}

<a id="L19"></a>type T struct {
    <a id="L20"></a>item  string;
    <a id="L21"></a>value string;
<a id="L22"></a>}

<a id="L24"></a>type S struct {
    <a id="L25"></a>header        string;
    <a id="L26"></a>integer       int;
    <a id="L27"></a>raw           string;
    <a id="L28"></a>innerT        T;
    <a id="L29"></a>innerPointerT *T;
    <a id="L30"></a>data          []T;
    <a id="L31"></a>pdata         []*T;
    <a id="L32"></a>empty         []*T;
    <a id="L33"></a>emptystring   string;
    <a id="L34"></a>null          []*T;
    <a id="L35"></a>vec           *vector.Vector;
    <a id="L36"></a>true          bool;
    <a id="L37"></a>false         bool;
<a id="L38"></a>}

<a id="L40"></a>var t1 = T{&#34;ItemNumber1&#34;, &#34;ValueNumber1&#34;}
<a id="L41"></a>var t2 = T{&#34;ItemNumber2&#34;, &#34;ValueNumber2&#34;}

<a id="L43"></a>func uppercase(v interface{}) string {
    <a id="L44"></a>s := v.(string);
    <a id="L45"></a>t := &#34;&#34;;
    <a id="L46"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L47"></a>c := s[i];
        <a id="L48"></a>if &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;z&#39; {
            <a id="L49"></a>c = c + &#39;A&#39; - &#39;a&#39;
        <a id="L50"></a>}
        <a id="L51"></a>t += string(c);
    <a id="L52"></a>}
    <a id="L53"></a>return t;
<a id="L54"></a>}

<a id="L56"></a>func plus1(v interface{}) string {
    <a id="L57"></a>i := v.(int);
    <a id="L58"></a>return fmt.Sprint(i + 1);
<a id="L59"></a>}

<a id="L61"></a>func writer(f func(interface{}) string) (func(io.Writer, interface{}, string)) {
    <a id="L62"></a>return func(w io.Writer, v interface{}, format string) {
        <a id="L63"></a>io.WriteString(w, f(v))
    <a id="L64"></a>}
<a id="L65"></a>}


<a id="L68"></a>var formatters = FormatterMap{
    <a id="L69"></a>&#34;uppercase&#34;: writer(uppercase),
    <a id="L70"></a>&#34;+1&#34;: writer(plus1),
<a id="L71"></a>}

<a id="L73"></a>var tests = []*Test{
    <a id="L74"></a><span class="comment">// Simple</span>
    <a id="L75"></a>&amp;Test{&#34;&#34;, &#34;&#34;, &#34;&#34;},
    <a id="L76"></a>&amp;Test{&#34;abc\ndef\n&#34;, &#34;abc\ndef\n&#34;, &#34;&#34;},
    <a id="L77"></a>&amp;Test{&#34; {.meta-left}   \n&#34;, &#34;{&#34;, &#34;&#34;},
    <a id="L78"></a>&amp;Test{&#34; {.meta-right}   \n&#34;, &#34;}&#34;, &#34;&#34;},
    <a id="L79"></a>&amp;Test{&#34; {.space}   \n&#34;, &#34; &#34;, &#34;&#34;},
    <a id="L80"></a>&amp;Test{&#34; {.tab}   \n&#34;, &#34;\t&#34;, &#34;&#34;},
    <a id="L81"></a>&amp;Test{&#34;     {#comment}   \n&#34;, &#34;&#34;, &#34;&#34;},

    <a id="L83"></a><span class="comment">// Variables at top level</span>
    <a id="L84"></a>&amp;Test{
        <a id="L85"></a>in: &#34;{header}={integer}\n&#34;,

        <a id="L87"></a>out: &#34;Header=77\n&#34;,
    <a id="L88"></a>},

    <a id="L90"></a><span class="comment">// Section</span>
    <a id="L91"></a>&amp;Test{
        <a id="L92"></a>in: &#34;{.section data }\n&#34;
            <a id="L93"></a>&#34;some text for the section\n&#34;
            <a id="L94"></a>&#34;{.end}\n&#34;,

        <a id="L96"></a>out: &#34;some text for the section\n&#34;,
    <a id="L97"></a>},
    <a id="L98"></a>&amp;Test{
        <a id="L99"></a>in: &#34;{.section data }\n&#34;
            <a id="L100"></a>&#34;{header}={integer}\n&#34;
            <a id="L101"></a>&#34;{.end}\n&#34;,

        <a id="L103"></a>out: &#34;Header=77\n&#34;,
    <a id="L104"></a>},
    <a id="L105"></a>&amp;Test{
        <a id="L106"></a>in: &#34;{.section pdata }\n&#34;
            <a id="L107"></a>&#34;{header}={integer}\n&#34;
            <a id="L108"></a>&#34;{.end}\n&#34;,

        <a id="L110"></a>out: &#34;Header=77\n&#34;,
    <a id="L111"></a>},
    <a id="L112"></a>&amp;Test{
        <a id="L113"></a>in: &#34;{.section pdata }\n&#34;
            <a id="L114"></a>&#34;data present\n&#34;
            <a id="L115"></a>&#34;{.or}\n&#34;
            <a id="L116"></a>&#34;data not present\n&#34;
            <a id="L117"></a>&#34;{.end}\n&#34;,

        <a id="L119"></a>out: &#34;data present\n&#34;,
    <a id="L120"></a>},
    <a id="L121"></a>&amp;Test{
        <a id="L122"></a>in: &#34;{.section empty }\n&#34;
            <a id="L123"></a>&#34;data present\n&#34;
            <a id="L124"></a>&#34;{.or}\n&#34;
            <a id="L125"></a>&#34;data not present\n&#34;
            <a id="L126"></a>&#34;{.end}\n&#34;,

        <a id="L128"></a>out: &#34;data not present\n&#34;,
    <a id="L129"></a>},
    <a id="L130"></a>&amp;Test{
        <a id="L131"></a>in: &#34;{.section null }\n&#34;
            <a id="L132"></a>&#34;data present\n&#34;
            <a id="L133"></a>&#34;{.or}\n&#34;
            <a id="L134"></a>&#34;data not present\n&#34;
            <a id="L135"></a>&#34;{.end}\n&#34;,

        <a id="L137"></a>out: &#34;data not present\n&#34;,
    <a id="L138"></a>},
    <a id="L139"></a>&amp;Test{
        <a id="L140"></a>in: &#34;{.section pdata }\n&#34;
            <a id="L141"></a>&#34;{header}={integer}\n&#34;
            <a id="L142"></a>&#34;{.section @ }\n&#34;
            <a id="L143"></a>&#34;{header}={integer}\n&#34;
            <a id="L144"></a>&#34;{.end}\n&#34;
            <a id="L145"></a>&#34;{.end}\n&#34;,

        <a id="L147"></a>out: &#34;Header=77\n&#34;
            <a id="L148"></a>&#34;Header=77\n&#34;,
    <a id="L149"></a>},
    <a id="L150"></a>&amp;Test{
        <a id="L151"></a>in: &#34;{.section data}{.end} {header}\n&#34;,

        <a id="L153"></a>out: &#34; Header\n&#34;,
    <a id="L154"></a>},

    <a id="L156"></a><span class="comment">// Repeated</span>
    <a id="L157"></a>&amp;Test{
        <a id="L158"></a>in: &#34;{.section pdata }\n&#34;
            <a id="L159"></a>&#34;{.repeated section @ }\n&#34;
            <a id="L160"></a>&#34;{item}={value}\n&#34;
            <a id="L161"></a>&#34;{.end}\n&#34;
            <a id="L162"></a>&#34;{.end}\n&#34;,

        <a id="L164"></a>out: &#34;ItemNumber1=ValueNumber1\n&#34;
            <a id="L165"></a>&#34;ItemNumber2=ValueNumber2\n&#34;,
    <a id="L166"></a>},
    <a id="L167"></a>&amp;Test{
        <a id="L168"></a>in: &#34;{.section pdata }\n&#34;
            <a id="L169"></a>&#34;{.repeated section @ }\n&#34;
            <a id="L170"></a>&#34;{item}={value}\n&#34;
            <a id="L171"></a>&#34;{.or}\n&#34;
            <a id="L172"></a>&#34;this should not appear\n&#34;
            <a id="L173"></a>&#34;{.end}\n&#34;
            <a id="L174"></a>&#34;{.end}\n&#34;,

        <a id="L176"></a>out: &#34;ItemNumber1=ValueNumber1\n&#34;
            <a id="L177"></a>&#34;ItemNumber2=ValueNumber2\n&#34;,
    <a id="L178"></a>},
    <a id="L179"></a>&amp;Test{
        <a id="L180"></a>in: &#34;{.section @ }\n&#34;
            <a id="L181"></a>&#34;{.repeated section empty }\n&#34;
            <a id="L182"></a>&#34;{item}={value}\n&#34;
            <a id="L183"></a>&#34;{.or}\n&#34;
            <a id="L184"></a>&#34;this should appear: empty field\n&#34;
            <a id="L185"></a>&#34;{.end}\n&#34;
            <a id="L186"></a>&#34;{.end}\n&#34;,

        <a id="L188"></a>out: &#34;this should appear: empty field\n&#34;,
    <a id="L189"></a>},
    <a id="L190"></a>&amp;Test{
        <a id="L191"></a>in: &#34;{.repeated section pdata }\n&#34;
            <a id="L192"></a>&#34;{item}\n&#34;
            <a id="L193"></a>&#34;{.alternates with}\n&#34;
            <a id="L194"></a>&#34;is\nover\nmultiple\nlines\n&#34;
            <a id="L195"></a>&#34;{.end}\n&#34;,

        <a id="L197"></a>out: &#34;ItemNumber1\n&#34;
            <a id="L198"></a>&#34;is\nover\nmultiple\nlines\n&#34;
            <a id="L199"></a>&#34;ItemNumber2\n&#34;,
    <a id="L200"></a>},
    <a id="L201"></a>&amp;Test{
        <a id="L202"></a>in: &#34;{.section pdata }\n&#34;
            <a id="L203"></a>&#34;{.repeated section @ }\n&#34;
            <a id="L204"></a>&#34;{item}={value}\n&#34;
            <a id="L205"></a>&#34;{.alternates with}DIVIDER\n&#34;
            <a id="L206"></a>&#34;{.or}\n&#34;
            <a id="L207"></a>&#34;this should not appear\n&#34;
            <a id="L208"></a>&#34;{.end}\n&#34;
            <a id="L209"></a>&#34;{.end}\n&#34;,

        <a id="L211"></a>out: &#34;ItemNumber1=ValueNumber1\n&#34;
            <a id="L212"></a>&#34;DIVIDER\n&#34;
            <a id="L213"></a>&#34;ItemNumber2=ValueNumber2\n&#34;,
    <a id="L214"></a>},
    <a id="L215"></a>&amp;Test{
        <a id="L216"></a>in: &#34;{.repeated section vec }\n&#34;
            <a id="L217"></a>&#34;{@}\n&#34;
            <a id="L218"></a>&#34;{.end}\n&#34;,

        <a id="L220"></a>out: &#34;elt1\n&#34;
            <a id="L221"></a>&#34;elt2\n&#34;,
    <a id="L222"></a>},
    <a id="L223"></a>&amp;Test{
        <a id="L224"></a>in: &#34;{.repeated section integer}{.end}&#34;,

        <a id="L226"></a>err: &#34;line 1: .repeated: cannot repeat integer (type int)&#34;,
    <a id="L227"></a>},

    <a id="L229"></a><span class="comment">// Nested names</span>
    <a id="L230"></a>&amp;Test{
        <a id="L231"></a>in: &#34;{.section @ }\n&#34;
            <a id="L232"></a>&#34;{innerT.item}={innerT.value}\n&#34;
            <a id="L233"></a>&#34;{.end}&#34;,

        <a id="L235"></a>out: &#34;ItemNumber1=ValueNumber1\n&#34;,
    <a id="L236"></a>},
    <a id="L237"></a>&amp;Test{
        <a id="L238"></a>in: &#34;{.section @ }\n&#34;
            <a id="L239"></a>&#34;{innerT.item}={.section innerT}{.section value}{@}{.end}{.end}\n&#34;
            <a id="L240"></a>&#34;{.end}&#34;,

        <a id="L242"></a>out: &#34;ItemNumber1=ValueNumber1\n&#34;,
    <a id="L243"></a>},


    <a id="L246"></a><span class="comment">// Formatters</span>
    <a id="L247"></a>&amp;Test{
        <a id="L248"></a>in: &#34;{.section pdata }\n&#34;
            <a id="L249"></a>&#34;{header|uppercase}={integer|+1}\n&#34;
            <a id="L250"></a>&#34;{header|html}={integer|str}\n&#34;
            <a id="L251"></a>&#34;{.end}\n&#34;,

        <a id="L253"></a>out: &#34;HEADER=78\n&#34;
            <a id="L254"></a>&#34;Header=77\n&#34;,
    <a id="L255"></a>},

    <a id="L257"></a>&amp;Test{
        <a id="L258"></a>in: &#34;{raw}\n&#34;
            <a id="L259"></a>&#34;{raw|html}\n&#34;,

        <a id="L261"></a>out: &#34;&amp;&lt;&gt;!@ #$%^\n&#34;
            <a id="L262"></a>&#34;&amp;amp;&amp;lt;&amp;gt;!@ #$%^\n&#34;,
    <a id="L263"></a>},

    <a id="L265"></a>&amp;Test{
        <a id="L266"></a>in: &#34;{.section emptystring}emptystring{.end}\n&#34;
            <a id="L267"></a>&#34;{.section header}header{.end}\n&#34;,

        <a id="L269"></a>out: &#34;\nheader\n&#34;,
    <a id="L270"></a>},

    <a id="L272"></a>&amp;Test{
        <a id="L273"></a>in: &#34;{.section true}1{.or}2{.end}\n&#34;
            <a id="L274"></a>&#34;{.section false}3{.or}4{.end}\n&#34;,

        <a id="L276"></a>out: &#34;1\n4\n&#34;,
    <a id="L277"></a>},
<a id="L278"></a>}

<a id="L280"></a>func TestAll(t *testing.T) {
    <a id="L281"></a>s := new(S);
    <a id="L282"></a><span class="comment">// initialized by hand for clarity.</span>
    <a id="L283"></a>s.header = &#34;Header&#34;;
    <a id="L284"></a>s.integer = 77;
    <a id="L285"></a>s.raw = &#34;&amp;&lt;&gt;!@ #$%^&#34;;
    <a id="L286"></a>s.innerT = t1;
    <a id="L287"></a>s.data = []T{t1, t2};
    <a id="L288"></a>s.pdata = []*T{&amp;t1, &amp;t2};
    <a id="L289"></a>s.empty = []*T{};
    <a id="L290"></a>s.null = nil;
    <a id="L291"></a>s.vec = vector.New(0);
    <a id="L292"></a>s.vec.Push(&#34;elt1&#34;);
    <a id="L293"></a>s.vec.Push(&#34;elt2&#34;);
    <a id="L294"></a>s.true = true;
    <a id="L295"></a>s.false = false;

    <a id="L297"></a>var buf bytes.Buffer;
    <a id="L298"></a>for _, test := range tests {
        <a id="L299"></a>buf.Reset();
        <a id="L300"></a>tmpl, err := Parse(test.in, formatters);
        <a id="L301"></a>if err != nil {
            <a id="L302"></a>t.Error(&#34;unexpected parse error:&#34;, err);
            <a id="L303"></a>continue;
        <a id="L304"></a>}
        <a id="L305"></a>err = tmpl.Execute(s, &amp;buf);
        <a id="L306"></a>if test.err == &#34;&#34; {
            <a id="L307"></a>if err != nil {
                <a id="L308"></a>t.Error(&#34;unexpected execute error:&#34;, err)
            <a id="L309"></a>}
        <a id="L310"></a>} else {
            <a id="L311"></a>if err == nil || err.String() != test.err {
                <a id="L312"></a>t.Errorf(&#34;expected execute error %q, got %q&#34;, test.err, err.String())
            <a id="L313"></a>}
        <a id="L314"></a>}
        <a id="L315"></a>if buf.String() != test.out {
            <a id="L316"></a>t.Errorf(&#34;for %q: expected %q got %q&#34;, test.in, test.out, buf.String())
        <a id="L317"></a>}
    <a id="L318"></a>}
<a id="L319"></a>}

<a id="L321"></a>func TestStringDriverType(t *testing.T) {
    <a id="L322"></a>tmpl, err := Parse(&#34;template: {@}&#34;, nil);
    <a id="L323"></a>if err != nil {
        <a id="L324"></a>t.Error(&#34;unexpected parse error:&#34;, err)
    <a id="L325"></a>}
    <a id="L326"></a>var b bytes.Buffer;
    <a id="L327"></a>err = tmpl.Execute(&#34;hello&#34;, &amp;b);
    <a id="L328"></a>if err != nil {
        <a id="L329"></a>t.Error(&#34;unexpected execute error:&#34;, err)
    <a id="L330"></a>}
    <a id="L331"></a>s := b.String();
    <a id="L332"></a>if s != &#34;template: hello&#34; {
        <a id="L333"></a>t.Errorf(&#34;failed passing string as data: expected %q got %q&#34;, &#34;template: hello&#34;, s)
    <a id="L334"></a>}
<a id="L335"></a>}

<a id="L337"></a>func TestTwice(t *testing.T) {
    <a id="L338"></a>tmpl, err := Parse(&#34;template: {@}&#34;, nil);
    <a id="L339"></a>if err != nil {
        <a id="L340"></a>t.Error(&#34;unexpected parse error:&#34;, err)
    <a id="L341"></a>}
    <a id="L342"></a>var b bytes.Buffer;
    <a id="L343"></a>err = tmpl.Execute(&#34;hello&#34;, &amp;b);
    <a id="L344"></a>if err != nil {
        <a id="L345"></a>t.Error(&#34;unexpected parse error:&#34;, err)
    <a id="L346"></a>}
    <a id="L347"></a>s := b.String();
    <a id="L348"></a>text := &#34;template: hello&#34;;
    <a id="L349"></a>if s != text {
        <a id="L350"></a>t.Errorf(&#34;failed passing string as data: expected %q got %q&#34;, text, s)
    <a id="L351"></a>}
    <a id="L352"></a>err = tmpl.Execute(&#34;hello&#34;, &amp;b);
    <a id="L353"></a>if err != nil {
        <a id="L354"></a>t.Error(&#34;unexpected parse error:&#34;, err)
    <a id="L355"></a>}
    <a id="L356"></a>s = b.String();
    <a id="L357"></a>text += text;
    <a id="L358"></a>if s != text {
        <a id="L359"></a>t.Errorf(&#34;failed passing string as data: expected %q got %q&#34;, text, s)
    <a id="L360"></a>}
<a id="L361"></a>}

<a id="L363"></a>func TestCustomDelims(t *testing.T) {
    <a id="L364"></a><span class="comment">// try various lengths.  zero should catch error.</span>
    <a id="L365"></a>for i := 0; i &lt; 7; i++ {
        <a id="L366"></a>for j := 0; j &lt; 7; j++ {
            <a id="L367"></a>tmpl := New(nil);
            <a id="L368"></a><span class="comment">// first two chars deliberately the same to test equal left and right delims</span>
            <a id="L369"></a>ldelim := &#34;$!#$%^&amp;&#34;[0:i];
            <a id="L370"></a>rdelim := &#34;$*&amp;^%$!&#34;[0:j];
            <a id="L371"></a>tmpl.SetDelims(ldelim, rdelim);
            <a id="L372"></a><span class="comment">// if braces, this would be template: {@}{.meta-left}{.meta-right}</span>
            <a id="L373"></a>text := &#34;template: &#34; +
                <a id="L374"></a>ldelim + &#34;@&#34; + rdelim +
                <a id="L375"></a>ldelim + &#34;.meta-left&#34; + rdelim +
                <a id="L376"></a>ldelim + &#34;.meta-right&#34; + rdelim;
            <a id="L377"></a>err := tmpl.Parse(text);
            <a id="L378"></a>if err != nil {
                <a id="L379"></a>if i == 0 || j == 0 { <span class="comment">// expected</span>
                    <a id="L380"></a>continue
                <a id="L381"></a>}
                <a id="L382"></a>t.Error(&#34;unexpected parse error:&#34;, err);
            <a id="L383"></a>} else if i == 0 || j == 0 {
                <a id="L384"></a>t.Errorf(&#34;expected parse error for empty delimiter: %d %d %q %q&#34;, i, j, ldelim, rdelim);
                <a id="L385"></a>continue;
            <a id="L386"></a>}
            <a id="L387"></a>var b bytes.Buffer;
            <a id="L388"></a>err = tmpl.Execute(&#34;hello&#34;, &amp;b);
            <a id="L389"></a>s := b.String();
            <a id="L390"></a>if s != &#34;template: hello&#34;+ldelim+rdelim {
                <a id="L391"></a>t.Errorf(&#34;failed delim check(%q %q) %q got %q&#34;, ldelim, rdelim, text, s)
            <a id="L392"></a>}
        <a id="L393"></a>}
    <a id="L394"></a>}
<a id="L395"></a>}

<a id="L397"></a><span class="comment">// Test that a variable evaluates to the field itself and does not further indirection</span>
<a id="L398"></a>func TestVarIndirection(t *testing.T) {
    <a id="L399"></a>s := new(S);
    <a id="L400"></a><span class="comment">// initialized by hand for clarity.</span>
    <a id="L401"></a>s.innerPointerT = &amp;t1;

    <a id="L403"></a>var buf bytes.Buffer;
    <a id="L404"></a>input := &#34;{.section @}{innerPointerT}{.end}&#34;;
    <a id="L405"></a>tmpl, err := Parse(input, nil);
    <a id="L406"></a>if err != nil {
        <a id="L407"></a>t.Fatal(&#34;unexpected parse error:&#34;, err)
    <a id="L408"></a>}
    <a id="L409"></a>err = tmpl.Execute(s, &amp;buf);
    <a id="L410"></a>if err != nil {
        <a id="L411"></a>t.Fatal(&#34;unexpected execute error:&#34;, err)
    <a id="L412"></a>}
    <a id="L413"></a>expect := fmt.Sprintf(&#34;%v&#34;, &amp;t1); <span class="comment">// output should be hex address of t1</span>
    <a id="L414"></a>if buf.String() != expect {
        <a id="L415"></a>t.Errorf(&#34;for %q: expected %q got %q&#34;, input, expect, buf.String())
    <a id="L416"></a>}
<a id="L417"></a>}
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
