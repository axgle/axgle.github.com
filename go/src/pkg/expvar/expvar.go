<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/expvar/expvar.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/expvar/expvar.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The expvar package provides a standardized interface to public variables,</span>
<a id="L6"></a><span class="comment">// such as operation counters in servers. It exposes these variables via</span>
<a id="L7"></a><span class="comment">// HTTP at /debug/vars in JSON format.</span>
<a id="L8"></a>package expvar

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bytes&#34;;
    <a id="L12"></a>&#34;fmt&#34;;
    <a id="L13"></a>&#34;http&#34;;
    <a id="L14"></a>&#34;log&#34;;
    <a id="L15"></a>&#34;strconv&#34;;
    <a id="L16"></a>&#34;sync&#34;;
<a id="L17"></a>)

<a id="L19"></a><span class="comment">// Var is an abstract type for all exported variables.</span>
<a id="L20"></a>type Var interface {
    <a id="L21"></a>String() string;
<a id="L22"></a>}

<a id="L24"></a><span class="comment">// Int is a 64-bit integer variable, and satisfies the Var interface.</span>
<a id="L25"></a>type Int struct {
    <a id="L26"></a>i   int64;
    <a id="L27"></a>mu  sync.Mutex;
<a id="L28"></a>}

<a id="L30"></a>func (v *Int) String() string { return strconv.Itoa64(v.i) }

<a id="L32"></a>func (v *Int) Add(delta int64) {
    <a id="L33"></a>v.mu.Lock();
    <a id="L34"></a>defer v.mu.Unlock();
    <a id="L35"></a>v.i += delta;
<a id="L36"></a>}

<a id="L38"></a><span class="comment">// Map is a string-to-Var map variable, and satisfies the Var interface.</span>
<a id="L39"></a>type Map struct {
    <a id="L40"></a>m   map[string]Var;
    <a id="L41"></a>mu  sync.Mutex;
<a id="L42"></a>}

<a id="L44"></a><span class="comment">// KeyValue represents a single entry in a Map.</span>
<a id="L45"></a>type KeyValue struct {
    <a id="L46"></a>Key   string;
    <a id="L47"></a>Value Var;
<a id="L48"></a>}

<a id="L50"></a>func (v *Map) String() string {
    <a id="L51"></a>v.mu.Lock();
    <a id="L52"></a>defer v.mu.Unlock();
    <a id="L53"></a>b := new(bytes.Buffer);
    <a id="L54"></a>fmt.Fprintf(b, &#34;{&#34;);
    <a id="L55"></a>first := true;
    <a id="L56"></a>for key, val := range v.m {
        <a id="L57"></a>if !first {
            <a id="L58"></a>fmt.Fprintf(b, &#34;, &#34;)
        <a id="L59"></a>}
        <a id="L60"></a>fmt.Fprintf(b, &#34;\&#34;%s\&#34;: %v&#34;, key, val.String());
        <a id="L61"></a>first = false;
    <a id="L62"></a>}
    <a id="L63"></a>fmt.Fprintf(b, &#34;}&#34;);
    <a id="L64"></a>return b.String();
<a id="L65"></a>}

<a id="L67"></a>func (v *Map) Init() *Map {
    <a id="L68"></a>v.m = make(map[string]Var);
    <a id="L69"></a>return v;
<a id="L70"></a>}

<a id="L72"></a>func (v *Map) Get(key string) Var {
    <a id="L73"></a>v.mu.Lock();
    <a id="L74"></a>defer v.mu.Unlock();
    <a id="L75"></a>if av, ok := v.m[key]; ok {
        <a id="L76"></a>return av
    <a id="L77"></a>}
    <a id="L78"></a>return nil;
<a id="L79"></a>}

<a id="L81"></a>func (v *Map) Set(key string, av Var) {
    <a id="L82"></a>v.mu.Lock();
    <a id="L83"></a>defer v.mu.Unlock();
    <a id="L84"></a>v.m[key] = av;
<a id="L85"></a>}

<a id="L87"></a>func (v *Map) Add(key string, delta int64) {
    <a id="L88"></a>v.mu.Lock();
    <a id="L89"></a>defer v.mu.Unlock();
    <a id="L90"></a>av, ok := v.m[key];
    <a id="L91"></a>if !ok {
        <a id="L92"></a>av = new(Int);
        <a id="L93"></a>v.m[key] = av;
    <a id="L94"></a>}

    <a id="L96"></a><span class="comment">// Add to Int; ignore otherwise.</span>
    <a id="L97"></a>if iv, ok := av.(*Int); ok {
        <a id="L98"></a>iv.Add(delta)
    <a id="L99"></a>}
<a id="L100"></a>}

<a id="L102"></a><span class="comment">// TODO(rsc): Make sure map access in separate thread is safe.</span>
<a id="L103"></a>func (v *Map) iterate(c chan&lt;- KeyValue) {
    <a id="L104"></a>for k, v := range v.m {
        <a id="L105"></a>c &lt;- KeyValue{k, v}
    <a id="L106"></a>}
    <a id="L107"></a>close(c);
<a id="L108"></a>}

<a id="L110"></a>func (v *Map) Iter() &lt;-chan KeyValue {
    <a id="L111"></a>c := make(chan KeyValue);
    <a id="L112"></a>go v.iterate(c);
    <a id="L113"></a>return c;
<a id="L114"></a>}

<a id="L116"></a><span class="comment">// String is a string variable, and satisfies the Var interface.</span>
<a id="L117"></a>type String struct {
    <a id="L118"></a>s string;
<a id="L119"></a>}

<a id="L121"></a>func (v *String) String() string { return strconv.Quote(v.s) }

<a id="L123"></a>func (v *String) Set(value string) { v.s = value }

<a id="L125"></a><span class="comment">// IntFunc wraps a func() int64 to create a value that satisfies the Var interface.</span>
<a id="L126"></a><span class="comment">// The function will be called each time the Var is evaluated.</span>
<a id="L127"></a>type IntFunc func() int64

<a id="L129"></a>func (v IntFunc) String() string { return strconv.Itoa64(v()) }


<a id="L132"></a><span class="comment">// All published variables.</span>
<a id="L133"></a>var vars map[string]Var = make(map[string]Var)
<a id="L134"></a>var mutex sync.Mutex

<a id="L136"></a><span class="comment">// Publish declares an named exported variable. This should be called from a</span>
<a id="L137"></a><span class="comment">// package&#39;s init function when it creates its Vars. If the name is already</span>
<a id="L138"></a><span class="comment">// registered then this will log.Crash.</span>
<a id="L139"></a>func Publish(name string, v Var) {
    <a id="L140"></a>mutex.Lock();
    <a id="L141"></a>defer mutex.Unlock();
    <a id="L142"></a>if _, existing := vars[name]; existing {
        <a id="L143"></a>log.Crash(&#34;Reuse of exported var name:&#34;, name)
    <a id="L144"></a>}
    <a id="L145"></a>vars[name] = v;
<a id="L146"></a>}

<a id="L148"></a><span class="comment">// Get retrieves a named exported variable.</span>
<a id="L149"></a>func Get(name string) Var {
    <a id="L150"></a>if v, ok := vars[name]; ok {
        <a id="L151"></a>return v
    <a id="L152"></a>}
    <a id="L153"></a>return nil;
<a id="L154"></a>}

<a id="L156"></a><span class="comment">// RemoveAll removes all exported variables.</span>
<a id="L157"></a><span class="comment">// This is for tests; don&#39;t call this on a real server.</span>
<a id="L158"></a>func RemoveAll() {
    <a id="L159"></a>mutex.Lock();
    <a id="L160"></a>defer mutex.Unlock();
    <a id="L161"></a>vars = make(map[string]Var);
<a id="L162"></a>}

<a id="L164"></a><span class="comment">// Convenience functions for creating new exported variables.</span>

<a id="L166"></a>func NewInt(name string) *Int {
    <a id="L167"></a>v := new(Int);
    <a id="L168"></a>Publish(name, v);
    <a id="L169"></a>return v;
<a id="L170"></a>}

<a id="L172"></a>func NewMap(name string) *Map {
    <a id="L173"></a>v := new(Map).Init();
    <a id="L174"></a>Publish(name, v);
    <a id="L175"></a>return v;
<a id="L176"></a>}

<a id="L178"></a>func NewString(name string) *String {
    <a id="L179"></a>v := new(String);
    <a id="L180"></a>Publish(name, v);
    <a id="L181"></a>return v;
<a id="L182"></a>}

<a id="L184"></a><span class="comment">// TODO(rsc): Make sure map access in separate thread is safe.</span>
<a id="L185"></a>func iterate(c chan&lt;- KeyValue) {
    <a id="L186"></a>for k, v := range vars {
        <a id="L187"></a>c &lt;- KeyValue{k, v}
    <a id="L188"></a>}
    <a id="L189"></a>close(c);
<a id="L190"></a>}

<a id="L192"></a>func Iter() &lt;-chan KeyValue {
    <a id="L193"></a>c := make(chan KeyValue);
    <a id="L194"></a>go iterate(c);
    <a id="L195"></a>return c;
<a id="L196"></a>}

<a id="L198"></a>func expvarHandler(c *http.Conn, req *http.Request) {
    <a id="L199"></a>c.SetHeader(&#34;content-type&#34;, &#34;application/json; charset=utf-8&#34;);
    <a id="L200"></a>fmt.Fprintf(c, &#34;{\n&#34;);
    <a id="L201"></a>first := true;
    <a id="L202"></a>for name, value := range vars {
        <a id="L203"></a>if !first {
            <a id="L204"></a>fmt.Fprintf(c, &#34;,\n&#34;)
        <a id="L205"></a>}
        <a id="L206"></a>first = false;
        <a id="L207"></a>fmt.Fprintf(c, &#34;  %q: %s&#34;, name, value);
    <a id="L208"></a>}
    <a id="L209"></a>fmt.Fprintf(c, &#34;\n}\n&#34;);
<a id="L210"></a>}

<a id="L212"></a>func init() { http.Handle(&#34;/debug/vars&#34;, http.HandlerFunc(expvarHandler)) }
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
