<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/json/generic.go</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/json/generic.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Generic representation of JSON objects.</span>

<a id="L7"></a>package json

<a id="L9"></a>import (
    <a id="L10"></a>&#34;container/vector&#34;;
    <a id="L11"></a>&#34;fmt&#34;;
    <a id="L12"></a>&#34;math&#34;;
    <a id="L13"></a>&#34;strconv&#34;;
    <a id="L14"></a>&#34;strings&#34;;
<a id="L15"></a>)

<a id="L17"></a><span class="comment">// Integers identifying the data type in the Json interface.</span>
<a id="L18"></a>const (
    <a id="L19"></a>StringKind = iota;
    <a id="L20"></a>NumberKind;
    <a id="L21"></a>MapKind; <span class="comment">// JSON term is &#34;Object&#34;, but in Go, it&#39;s a map</span>
    <a id="L22"></a>ArrayKind;
    <a id="L23"></a>BoolKind;
    <a id="L24"></a>NullKind;
<a id="L25"></a>)

<a id="L27"></a><span class="comment">// The Json interface is implemented by all JSON objects.</span>
<a id="L28"></a>type Json interface {
    <a id="L29"></a>Kind() int;         <span class="comment">// StringKind, NumberKind, etc.</span>
    <a id="L30"></a>String() string;    <span class="comment">// a string form (any kind)</span>
    <a id="L31"></a>Number() float64;   <span class="comment">// numeric form (NumberKind)</span>
    <a id="L32"></a>Bool() bool;        <span class="comment">// boolean (BoolKind)</span>
    <a id="L33"></a>Get(s string) Json; <span class="comment">// field lookup (MapKind)</span>
    <a id="L34"></a>Elem(i int) Json;   <span class="comment">// element lookup (ArrayKind)</span>
    <a id="L35"></a>Len() int;          <span class="comment">// length (ArrayKind, MapKind)</span>
<a id="L36"></a>}

<a id="L38"></a><span class="comment">// JsonToString returns the textual JSON syntax representation</span>
<a id="L39"></a><span class="comment">// for the JSON object j.</span>
<a id="L40"></a><span class="comment">//</span>
<a id="L41"></a><span class="comment">// JsonToString differs from j.String() in the handling</span>
<a id="L42"></a><span class="comment">// of string objects.  If j represents the string abc,</span>
<a id="L43"></a><span class="comment">// j.String() == `abc`, but JsonToString(j) == `&#34;abc&#34;`.</span>
<a id="L44"></a>func JsonToString(j Json) string {
    <a id="L45"></a>if j == nil {
        <a id="L46"></a>return &#34;null&#34;
    <a id="L47"></a>}
    <a id="L48"></a>if j.Kind() == StringKind {
        <a id="L49"></a>return Quote(j.String())
    <a id="L50"></a>}
    <a id="L51"></a>return j.String();
<a id="L52"></a>}

<a id="L54"></a>type _Null struct{}

<a id="L56"></a><span class="comment">// Null is the JSON object representing the null data object.</span>
<a id="L57"></a>var Null Json = &amp;_Null{}

<a id="L59"></a>func (*_Null) Kind() int         { return NullKind }
<a id="L60"></a>func (*_Null) String() string    { return &#34;null&#34; }
<a id="L61"></a>func (*_Null) Number() float64   { return 0 }
<a id="L62"></a>func (*_Null) Bool() bool        { return false }
<a id="L63"></a>func (*_Null) Get(s string) Json { return Null }
<a id="L64"></a>func (*_Null) Elem(int) Json     { return Null }
<a id="L65"></a>func (*_Null) Len() int          { return 0 }

<a id="L67"></a>type _String struct {
    <a id="L68"></a>s   string;
    <a id="L69"></a>_Null;
<a id="L70"></a>}

<a id="L72"></a>func (j *_String) Kind() int      { return StringKind }
<a id="L73"></a>func (j *_String) String() string { return j.s }

<a id="L75"></a>type _Number struct {
    <a id="L76"></a>f   float64;
    <a id="L77"></a>_Null;
<a id="L78"></a>}

<a id="L80"></a>func (j *_Number) Kind() int       { return NumberKind }
<a id="L81"></a>func (j *_Number) Number() float64 { return j.f }
<a id="L82"></a>func (j *_Number) String() string {
    <a id="L83"></a>if math.Floor(j.f) == j.f {
        <a id="L84"></a>return fmt.Sprintf(&#34;%.0f&#34;, j.f)
    <a id="L85"></a>}
    <a id="L86"></a>return fmt.Sprintf(&#34;%g&#34;, j.f);
<a id="L87"></a>}

<a id="L89"></a>type _Array struct {
    <a id="L90"></a>a   *vector.Vector;
    <a id="L91"></a>_Null;
<a id="L92"></a>}

<a id="L94"></a>func (j *_Array) Kind() int { return ArrayKind }
<a id="L95"></a>func (j *_Array) Len() int  { return j.a.Len() }
<a id="L96"></a>func (j *_Array) Elem(i int) Json {
    <a id="L97"></a>if i &lt; 0 || i &gt;= j.a.Len() {
        <a id="L98"></a>return Null
    <a id="L99"></a>}
    <a id="L100"></a>return j.a.At(i).(Json);
<a id="L101"></a>}
<a id="L102"></a>func (j *_Array) String() string {
    <a id="L103"></a>s := &#34;[&#34;;
    <a id="L104"></a>for i := 0; i &lt; j.a.Len(); i++ {
        <a id="L105"></a>if i &gt; 0 {
            <a id="L106"></a>s += &#34;,&#34;
        <a id="L107"></a>}
        <a id="L108"></a>s += JsonToString(j.a.At(i).(Json));
    <a id="L109"></a>}
    <a id="L110"></a>s += &#34;]&#34;;
    <a id="L111"></a>return s;
<a id="L112"></a>}

<a id="L114"></a>type _Bool struct {
    <a id="L115"></a>b   bool;
    <a id="L116"></a>_Null;
<a id="L117"></a>}

<a id="L119"></a>func (j *_Bool) Kind() int  { return BoolKind }
<a id="L120"></a>func (j *_Bool) Bool() bool { return j.b }
<a id="L121"></a>func (j *_Bool) String() string {
    <a id="L122"></a>if j.b {
        <a id="L123"></a>return &#34;true&#34;
    <a id="L124"></a>}
    <a id="L125"></a>return &#34;false&#34;;
<a id="L126"></a>}

<a id="L128"></a>type _Map struct {
    <a id="L129"></a>m   map[string]Json;
    <a id="L130"></a>_Null;
<a id="L131"></a>}

<a id="L133"></a>func (j *_Map) Kind() int { return MapKind }
<a id="L134"></a>func (j *_Map) Len() int  { return len(j.m) }
<a id="L135"></a>func (j *_Map) Get(s string) Json {
    <a id="L136"></a>if j.m == nil {
        <a id="L137"></a>return Null
    <a id="L138"></a>}
    <a id="L139"></a>v, ok := j.m[s];
    <a id="L140"></a>if !ok {
        <a id="L141"></a>return Null
    <a id="L142"></a>}
    <a id="L143"></a>return v;
<a id="L144"></a>}
<a id="L145"></a>func (j *_Map) String() string {
    <a id="L146"></a>s := &#34;{&#34;;
    <a id="L147"></a>first := true;
    <a id="L148"></a>for k, v := range j.m {
        <a id="L149"></a>if first {
            <a id="L150"></a>first = false
        <a id="L151"></a>} else {
            <a id="L152"></a>s += &#34;,&#34;
        <a id="L153"></a>}
        <a id="L154"></a>s += Quote(k);
        <a id="L155"></a>s += &#34;:&#34;;
        <a id="L156"></a>s += JsonToString(v);
    <a id="L157"></a>}
    <a id="L158"></a>s += &#34;}&#34;;
    <a id="L159"></a>return s;
<a id="L160"></a>}

<a id="L162"></a><span class="comment">// Walk evaluates path relative to the JSON object j.</span>
<a id="L163"></a><span class="comment">// Path is taken as a sequence of slash-separated field names</span>
<a id="L164"></a><span class="comment">// or numbers that can be used to index into JSON map and</span>
<a id="L165"></a><span class="comment">// array objects.</span>
<a id="L166"></a><span class="comment">//</span>
<a id="L167"></a><span class="comment">// For example, if j is the JSON object for</span>
<a id="L168"></a><span class="comment">// {&#34;abc&#34;: [true, false]}, then Walk(j, &#34;abc/1&#34;) returns the</span>
<a id="L169"></a><span class="comment">// JSON object for true.</span>
<a id="L170"></a>func Walk(j Json, path string) Json {
    <a id="L171"></a>for len(path) &gt; 0 {
        <a id="L172"></a>var elem string;
        <a id="L173"></a>if i := strings.Index(path, &#34;/&#34;); i &gt;= 0 {
            <a id="L174"></a>elem = path[0:i];
            <a id="L175"></a>path = path[i+1 : len(path)];
        <a id="L176"></a>} else {
            <a id="L177"></a>elem = path;
            <a id="L178"></a>path = &#34;&#34;;
        <a id="L179"></a>}
        <a id="L180"></a>switch j.Kind() {
        <a id="L181"></a>case ArrayKind:
            <a id="L182"></a>indx, err := strconv.Atoi(elem);
            <a id="L183"></a>if err != nil {
                <a id="L184"></a>return Null
            <a id="L185"></a>}
            <a id="L186"></a>j = j.Elem(indx);
        <a id="L187"></a>case MapKind:
            <a id="L188"></a>j = j.Get(elem)
        <a id="L189"></a>default:
            <a id="L190"></a>return Null
        <a id="L191"></a>}
    <a id="L192"></a>}
    <a id="L193"></a>return j;
<a id="L194"></a>}

<a id="L196"></a><span class="comment">// Equal returns whether a and b are indistinguishable JSON objects.</span>
<a id="L197"></a>func Equal(a, b Json) bool {
    <a id="L198"></a>switch {
    <a id="L199"></a>case a == nil &amp;&amp; b == nil:
        <a id="L200"></a>return true
    <a id="L201"></a>case a == nil || b == nil:
        <a id="L202"></a>return false
    <a id="L203"></a>case a.Kind() != b.Kind():
        <a id="L204"></a>return false
    <a id="L205"></a>}

    <a id="L207"></a>switch a.Kind() {
    <a id="L208"></a>case NullKind:
        <a id="L209"></a>return true
    <a id="L210"></a>case StringKind:
        <a id="L211"></a>return a.String() == b.String()
    <a id="L212"></a>case NumberKind:
        <a id="L213"></a>return a.Number() == b.Number()
    <a id="L214"></a>case BoolKind:
        <a id="L215"></a>return a.Bool() == b.Bool()
    <a id="L216"></a>case ArrayKind:
        <a id="L217"></a>if a.Len() != b.Len() {
            <a id="L218"></a>return false
        <a id="L219"></a>}
        <a id="L220"></a>for i := 0; i &lt; a.Len(); i++ {
            <a id="L221"></a>if !Equal(a.Elem(i), b.Elem(i)) {
                <a id="L222"></a>return false
            <a id="L223"></a>}
        <a id="L224"></a>}
        <a id="L225"></a>return true;
    <a id="L226"></a>case MapKind:
        <a id="L227"></a>m := a.(*_Map).m;
        <a id="L228"></a>if len(m) != len(b.(*_Map).m) {
            <a id="L229"></a>return false
        <a id="L230"></a>}
        <a id="L231"></a>for k, v := range m {
            <a id="L232"></a>if !Equal(v, b.Get(k)) {
                <a id="L233"></a>return false
            <a id="L234"></a>}
        <a id="L235"></a>}
        <a id="L236"></a>return true;
    <a id="L237"></a>}

    <a id="L239"></a><span class="comment">// invalid kind</span>
    <a id="L240"></a>return false;
<a id="L241"></a>}


<a id="L244"></a><span class="comment">// Parse builder for JSON objects.</span>

<a id="L246"></a>type _JsonBuilder struct {
    <a id="L247"></a><span class="comment">// either writing to *ptr</span>
    <a id="L248"></a>ptr *Json;

    <a id="L250"></a><span class="comment">// or to a[i] (can&#39;t set ptr = &amp;a[i])</span>
    <a id="L251"></a>a   *vector.Vector;
    <a id="L252"></a>i   int;

    <a id="L254"></a><span class="comment">// or to m[k] (can&#39;t set ptr = &amp;m[k])</span>
    <a id="L255"></a>m   map[string]Json;
    <a id="L256"></a>k   string;
<a id="L257"></a>}

<a id="L259"></a>func (b *_JsonBuilder) Put(j Json) {
    <a id="L260"></a>switch {
    <a id="L261"></a>case b.ptr != nil:
        <a id="L262"></a>*b.ptr = j
    <a id="L263"></a>case b.a != nil:
        <a id="L264"></a>b.a.Set(b.i, j)
    <a id="L265"></a>case b.m != nil:
        <a id="L266"></a>b.m[b.k] = j
    <a id="L267"></a>}
<a id="L268"></a>}

<a id="L270"></a>func (b *_JsonBuilder) Get() Json {
    <a id="L271"></a>switch {
    <a id="L272"></a>case b.ptr != nil:
        <a id="L273"></a>return *b.ptr
    <a id="L274"></a>case b.a != nil:
        <a id="L275"></a>return b.a.At(b.i).(Json)
    <a id="L276"></a>case b.m != nil:
        <a id="L277"></a>return b.m[b.k]
    <a id="L278"></a>}
    <a id="L279"></a>return nil;
<a id="L280"></a>}

<a id="L282"></a>func (b *_JsonBuilder) Float64(f float64) { b.Put(&amp;_Number{f, _Null{}}) }

<a id="L284"></a>func (b *_JsonBuilder) Int64(i int64) { b.Float64(float64(i)) }

<a id="L286"></a>func (b *_JsonBuilder) Uint64(i uint64) { b.Float64(float64(i)) }

<a id="L288"></a>func (b *_JsonBuilder) Bool(tf bool) { b.Put(&amp;_Bool{tf, _Null{}}) }

<a id="L290"></a>func (b *_JsonBuilder) Null() { b.Put(Null) }

<a id="L292"></a>func (b *_JsonBuilder) String(s string) { b.Put(&amp;_String{s, _Null{}}) }


<a id="L295"></a>func (b *_JsonBuilder) Array() { b.Put(&amp;_Array{vector.New(0), _Null{}}) }

<a id="L297"></a>func (b *_JsonBuilder) Map() { b.Put(&amp;_Map{make(map[string]Json), _Null{}}) }

<a id="L299"></a>func (b *_JsonBuilder) Elem(i int) Builder {
    <a id="L300"></a>bb := new(_JsonBuilder);
    <a id="L301"></a>bb.a = b.Get().(*_Array).a;
    <a id="L302"></a>bb.i = i;
    <a id="L303"></a>for i &gt;= bb.a.Len() {
        <a id="L304"></a>bb.a.Push(Null)
    <a id="L305"></a>}
    <a id="L306"></a>return bb;
<a id="L307"></a>}

<a id="L309"></a>func (b *_JsonBuilder) Key(k string) Builder {
    <a id="L310"></a>bb := new(_JsonBuilder);
    <a id="L311"></a>bb.m = b.Get().(*_Map).m;
    <a id="L312"></a>bb.k = k;
    <a id="L313"></a>bb.m[k] = Null;
    <a id="L314"></a>return bb;
<a id="L315"></a>}

<a id="L317"></a>func (b *_JsonBuilder) Flush() {}

<a id="L319"></a><span class="comment">// StringToJson parses the string s as a JSON-syntax string</span>
<a id="L320"></a><span class="comment">// and returns the generic JSON object representation.</span>
<a id="L321"></a><span class="comment">// On success, StringToJson returns with ok set to true and errtok empty.</span>
<a id="L322"></a><span class="comment">// If StringToJson encounters a syntax error, it returns with</span>
<a id="L323"></a><span class="comment">// ok set to false and errtok set to a fragment of the offending syntax.</span>
<a id="L324"></a>func StringToJson(s string) (json Json, ok bool, errtok string) {
    <a id="L325"></a>var j Json;
    <a id="L326"></a>b := new(_JsonBuilder);
    <a id="L327"></a>b.ptr = &amp;j;
    <a id="L328"></a>ok, _, errtok = Parse(s, b);
    <a id="L329"></a>if !ok {
        <a id="L330"></a>return nil, false, errtok
    <a id="L331"></a>}
    <a id="L332"></a>return j, true, &#34;&#34;;
<a id="L333"></a>}

<a id="L335"></a><span class="comment">// BUG(rsc): StringToJson should return an os.Error instead of a bool.</span>
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
