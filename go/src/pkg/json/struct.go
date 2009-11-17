<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/json/struct.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/json/struct.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Marshalling and unmarshalling of</span>
<a id="L6"></a><span class="comment">// JSON data into Go structs using reflection.</span>

<a id="L8"></a>package json

<a id="L10"></a>import (
    <a id="L11"></a>&#34;reflect&#34;;
    <a id="L12"></a>&#34;strings&#34;;
<a id="L13"></a>)

<a id="L15"></a>type structBuilder struct {
    <a id="L16"></a>val reflect.Value;

    <a id="L18"></a><span class="comment">// if map_ != nil, write val to map_[key] on each change</span>
    <a id="L19"></a>map_ *reflect.MapValue;
    <a id="L20"></a>key  reflect.Value;
<a id="L21"></a>}

<a id="L23"></a>var nobuilder *structBuilder

<a id="L25"></a>func isfloat(v reflect.Value) bool {
    <a id="L26"></a>switch v.(type) {
    <a id="L27"></a>case *reflect.FloatValue, *reflect.Float32Value, *reflect.Float64Value:
        <a id="L28"></a>return true
    <a id="L29"></a>}
    <a id="L30"></a>return false;
<a id="L31"></a>}

<a id="L33"></a>func setfloat(v reflect.Value, f float64) {
    <a id="L34"></a>switch v := v.(type) {
    <a id="L35"></a>case *reflect.FloatValue:
        <a id="L36"></a>v.Set(float(f))
    <a id="L37"></a>case *reflect.Float32Value:
        <a id="L38"></a>v.Set(float32(f))
    <a id="L39"></a>case *reflect.Float64Value:
        <a id="L40"></a>v.Set(float64(f))
    <a id="L41"></a>}
<a id="L42"></a>}

<a id="L44"></a>func setint(v reflect.Value, i int64) {
    <a id="L45"></a>switch v := v.(type) {
    <a id="L46"></a>case *reflect.IntValue:
        <a id="L47"></a>v.Set(int(i))
    <a id="L48"></a>case *reflect.Int8Value:
        <a id="L49"></a>v.Set(int8(i))
    <a id="L50"></a>case *reflect.Int16Value:
        <a id="L51"></a>v.Set(int16(i))
    <a id="L52"></a>case *reflect.Int32Value:
        <a id="L53"></a>v.Set(int32(i))
    <a id="L54"></a>case *reflect.Int64Value:
        <a id="L55"></a>v.Set(int64(i))
    <a id="L56"></a>case *reflect.UintValue:
        <a id="L57"></a>v.Set(uint(i))
    <a id="L58"></a>case *reflect.Uint8Value:
        <a id="L59"></a>v.Set(uint8(i))
    <a id="L60"></a>case *reflect.Uint16Value:
        <a id="L61"></a>v.Set(uint16(i))
    <a id="L62"></a>case *reflect.Uint32Value:
        <a id="L63"></a>v.Set(uint32(i))
    <a id="L64"></a>case *reflect.Uint64Value:
        <a id="L65"></a>v.Set(uint64(i))
    <a id="L66"></a>}
<a id="L67"></a>}

<a id="L69"></a><span class="comment">// If updating b.val is not enough to update the original,</span>
<a id="L70"></a><span class="comment">// copy a changed b.val out to the original.</span>
<a id="L71"></a>func (b *structBuilder) Flush() {
    <a id="L72"></a>if b == nil {
        <a id="L73"></a>return
    <a id="L74"></a>}
    <a id="L75"></a>if b.map_ != nil {
        <a id="L76"></a>b.map_.SetElem(b.key, b.val)
    <a id="L77"></a>}
<a id="L78"></a>}

<a id="L80"></a>func (b *structBuilder) Int64(i int64) {
    <a id="L81"></a>if b == nil {
        <a id="L82"></a>return
    <a id="L83"></a>}
    <a id="L84"></a>v := b.val;
    <a id="L85"></a>if isfloat(v) {
        <a id="L86"></a>setfloat(v, float64(i))
    <a id="L87"></a>} else {
        <a id="L88"></a>setint(v, i)
    <a id="L89"></a>}
<a id="L90"></a>}

<a id="L92"></a>func (b *structBuilder) Uint64(i uint64) {
    <a id="L93"></a>if b == nil {
        <a id="L94"></a>return
    <a id="L95"></a>}
    <a id="L96"></a>v := b.val;
    <a id="L97"></a>if isfloat(v) {
        <a id="L98"></a>setfloat(v, float64(i))
    <a id="L99"></a>} else {
        <a id="L100"></a>setint(v, int64(i))
    <a id="L101"></a>}
<a id="L102"></a>}

<a id="L104"></a>func (b *structBuilder) Float64(f float64) {
    <a id="L105"></a>if b == nil {
        <a id="L106"></a>return
    <a id="L107"></a>}
    <a id="L108"></a>v := b.val;
    <a id="L109"></a>if isfloat(v) {
        <a id="L110"></a>setfloat(v, f)
    <a id="L111"></a>} else {
        <a id="L112"></a>setint(v, int64(f))
    <a id="L113"></a>}
<a id="L114"></a>}

<a id="L116"></a>func (b *structBuilder) Null() {}

<a id="L118"></a>func (b *structBuilder) String(s string) {
    <a id="L119"></a>if b == nil {
        <a id="L120"></a>return
    <a id="L121"></a>}
    <a id="L122"></a>if v, ok := b.val.(*reflect.StringValue); ok {
        <a id="L123"></a>v.Set(s)
    <a id="L124"></a>}
<a id="L125"></a>}

<a id="L127"></a>func (b *structBuilder) Bool(tf bool) {
    <a id="L128"></a>if b == nil {
        <a id="L129"></a>return
    <a id="L130"></a>}
    <a id="L131"></a>if v, ok := b.val.(*reflect.BoolValue); ok {
        <a id="L132"></a>v.Set(tf)
    <a id="L133"></a>}
<a id="L134"></a>}

<a id="L136"></a>func (b *structBuilder) Array() {
    <a id="L137"></a>if b == nil {
        <a id="L138"></a>return
    <a id="L139"></a>}
    <a id="L140"></a>if v, ok := b.val.(*reflect.SliceValue); ok {
        <a id="L141"></a>if v.IsNil() {
            <a id="L142"></a>v.Set(reflect.MakeSlice(v.Type().(*reflect.SliceType), 0, 8))
        <a id="L143"></a>}
    <a id="L144"></a>}
<a id="L145"></a>}

<a id="L147"></a>func (b *structBuilder) Elem(i int) Builder {
    <a id="L148"></a>if b == nil || i &lt; 0 {
        <a id="L149"></a>return nobuilder
    <a id="L150"></a>}
    <a id="L151"></a>switch v := b.val.(type) {
    <a id="L152"></a>case *reflect.ArrayValue:
        <a id="L153"></a>if i &lt; v.Len() {
            <a id="L154"></a>return &amp;structBuilder{val: v.Elem(i)}
        <a id="L155"></a>}
    <a id="L156"></a>case *reflect.SliceValue:
        <a id="L157"></a>if i &gt; v.Cap() {
            <a id="L158"></a>n := v.Cap();
            <a id="L159"></a>if n &lt; 8 {
                <a id="L160"></a>n = 8
            <a id="L161"></a>}
            <a id="L162"></a>for n &lt;= i {
                <a id="L163"></a>n *= 2
            <a id="L164"></a>}
            <a id="L165"></a>nv := reflect.MakeSlice(v.Type().(*reflect.SliceType), v.Len(), n);
            <a id="L166"></a>reflect.ArrayCopy(nv, v);
            <a id="L167"></a>v.Set(nv);
        <a id="L168"></a>}
        <a id="L169"></a>if v.Len() &lt;= i &amp;&amp; i &lt; v.Cap() {
            <a id="L170"></a>v.SetLen(i + 1)
        <a id="L171"></a>}
        <a id="L172"></a>if i &lt; v.Len() {
            <a id="L173"></a>return &amp;structBuilder{val: v.Elem(i)}
        <a id="L174"></a>}
    <a id="L175"></a>}
    <a id="L176"></a>return nobuilder;
<a id="L177"></a>}

<a id="L179"></a>func (b *structBuilder) Map() {
    <a id="L180"></a>if b == nil {
        <a id="L181"></a>return
    <a id="L182"></a>}
    <a id="L183"></a>if v, ok := b.val.(*reflect.PtrValue); ok &amp;&amp; v.IsNil() {
        <a id="L184"></a>if v.IsNil() {
            <a id="L185"></a>v.PointTo(reflect.MakeZero(v.Type().(*reflect.PtrType).Elem()));
            <a id="L186"></a>b.Flush();
        <a id="L187"></a>}
        <a id="L188"></a>b.map_ = nil;
        <a id="L189"></a>b.val = v.Elem();
    <a id="L190"></a>}
    <a id="L191"></a>if v, ok := b.val.(*reflect.MapValue); ok &amp;&amp; v.IsNil() {
        <a id="L192"></a>v.Set(reflect.MakeMap(v.Type().(*reflect.MapType)))
    <a id="L193"></a>}
<a id="L194"></a>}

<a id="L196"></a>func (b *structBuilder) Key(k string) Builder {
    <a id="L197"></a>if b == nil {
        <a id="L198"></a>return nobuilder
    <a id="L199"></a>}
    <a id="L200"></a>switch v := reflect.Indirect(b.val).(type) {
    <a id="L201"></a>case *reflect.StructValue:
        <a id="L202"></a>t := v.Type().(*reflect.StructType);
        <a id="L203"></a><span class="comment">// Case-insensitive field lookup.</span>
        <a id="L204"></a>k = strings.ToLower(k);
        <a id="L205"></a>for i := 0; i &lt; t.NumField(); i++ {
            <a id="L206"></a>if strings.ToLower(t.Field(i).Name) == k {
                <a id="L207"></a>return &amp;structBuilder{val: v.Field(i)}
            <a id="L208"></a>}
        <a id="L209"></a>}
    <a id="L210"></a>case *reflect.MapValue:
        <a id="L211"></a>t := v.Type().(*reflect.MapType);
        <a id="L212"></a>if t.Key() != reflect.Typeof(k) {
            <a id="L213"></a>break
        <a id="L214"></a>}
        <a id="L215"></a>key := reflect.NewValue(k);
        <a id="L216"></a>elem := v.Elem(key);
        <a id="L217"></a>if elem == nil {
            <a id="L218"></a>v.SetElem(key, reflect.MakeZero(t.Elem()));
            <a id="L219"></a>elem = v.Elem(key);
        <a id="L220"></a>}
        <a id="L221"></a>return &amp;structBuilder{val: elem, map_: v, key: key};
    <a id="L222"></a>}
    <a id="L223"></a>return nobuilder;
<a id="L224"></a>}

<a id="L226"></a><span class="comment">// Unmarshal parses the JSON syntax string s and fills in</span>
<a id="L227"></a><span class="comment">// an arbitrary struct or array pointed at by val.</span>
<a id="L228"></a><span class="comment">// It uses the reflect package to assign to fields</span>
<a id="L229"></a><span class="comment">// and arrays embedded in val.  Well-formed data that does not fit</span>
<a id="L230"></a><span class="comment">// into the struct is discarded.</span>
<a id="L231"></a><span class="comment">//</span>
<a id="L232"></a><span class="comment">// For example, given these definitions:</span>
<a id="L233"></a><span class="comment">//</span>
<a id="L234"></a><span class="comment">//	type Email struct {</span>
<a id="L235"></a><span class="comment">//		Where string;</span>
<a id="L236"></a><span class="comment">//		Addr string;</span>
<a id="L237"></a><span class="comment">//	}</span>
<a id="L238"></a><span class="comment">//</span>
<a id="L239"></a><span class="comment">//	type Result struct {</span>
<a id="L240"></a><span class="comment">//		Name string;</span>
<a id="L241"></a><span class="comment">//		Phone string;</span>
<a id="L242"></a><span class="comment">//		Email []Email</span>
<a id="L243"></a><span class="comment">//	}</span>
<a id="L244"></a><span class="comment">//</span>
<a id="L245"></a><span class="comment">//	var r = Result{ &#34;name&#34;, &#34;phone&#34;, nil }</span>
<a id="L246"></a><span class="comment">//</span>
<a id="L247"></a><span class="comment">// unmarshalling the JSON syntax string</span>
<a id="L248"></a><span class="comment">//</span>
<a id="L249"></a><span class="comment">//	{</span>
<a id="L250"></a><span class="comment">//	  &#34;email&#34;: [</span>
<a id="L251"></a><span class="comment">//	    {</span>
<a id="L252"></a><span class="comment">//	      &#34;where&#34;: &#34;home&#34;,</span>
<a id="L253"></a><span class="comment">//	      &#34;addr&#34;: &#34;gre@example.com&#34;</span>
<a id="L254"></a><span class="comment">//	    },</span>
<a id="L255"></a><span class="comment">//	    {</span>
<a id="L256"></a><span class="comment">//	      &#34;where&#34;: &#34;work&#34;,</span>
<a id="L257"></a><span class="comment">//	      &#34;addr&#34;: &#34;gre@work.com&#34;</span>
<a id="L258"></a><span class="comment">//	    }</span>
<a id="L259"></a><span class="comment">//	  ],</span>
<a id="L260"></a><span class="comment">//	  &#34;name&#34;: &#34;Grace R. Emlin&#34;,</span>
<a id="L261"></a><span class="comment">//	  &#34;address&#34;: &#34;123 Main Street&#34;</span>
<a id="L262"></a><span class="comment">//	}</span>
<a id="L263"></a><span class="comment">//</span>
<a id="L264"></a><span class="comment">// via Unmarshal(s, &amp;r) is equivalent to assigning</span>
<a id="L265"></a><span class="comment">//</span>
<a id="L266"></a><span class="comment">//	r = Result{</span>
<a id="L267"></a><span class="comment">//		&#34;Grace R. Emlin&#34;,	// name</span>
<a id="L268"></a><span class="comment">//		&#34;phone&#34;,		// no phone given</span>
<a id="L269"></a><span class="comment">//		[]Email{</span>
<a id="L270"></a><span class="comment">//			Email{ &#34;home&#34;, &#34;gre@example.com&#34; },</span>
<a id="L271"></a><span class="comment">//			Email{ &#34;work&#34;, &#34;gre@work.com&#34; }</span>
<a id="L272"></a><span class="comment">//		}</span>
<a id="L273"></a><span class="comment">//	}</span>
<a id="L274"></a><span class="comment">//</span>
<a id="L275"></a><span class="comment">// Note that the field r.Phone has not been modified and</span>
<a id="L276"></a><span class="comment">// that the JSON field &#34;address&#34; was discarded.</span>
<a id="L277"></a><span class="comment">//</span>
<a id="L278"></a><span class="comment">// Because Unmarshal uses the reflect package, it can only</span>
<a id="L279"></a><span class="comment">// assign to upper case fields.  Unmarshal uses a case-insensitive</span>
<a id="L280"></a><span class="comment">// comparison to match JSON field names to struct field names.</span>
<a id="L281"></a><span class="comment">//</span>
<a id="L282"></a><span class="comment">// On success, Unmarshal returns with ok set to true.</span>
<a id="L283"></a><span class="comment">// On a syntax error, it returns with ok set to false and errtok</span>
<a id="L284"></a><span class="comment">// set to the offending token.</span>
<a id="L285"></a>func Unmarshal(s string, val interface{}) (ok bool, errtok string) {
    <a id="L286"></a>b := &amp;structBuilder{val: reflect.NewValue(val)};
    <a id="L287"></a>ok, _, errtok = Parse(s, b);
    <a id="L288"></a>if !ok {
        <a id="L289"></a>return false, errtok
    <a id="L290"></a>}
    <a id="L291"></a>return true, &#34;&#34;;
<a id="L292"></a>}
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
