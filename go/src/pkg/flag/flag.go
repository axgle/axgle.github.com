<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/flag/flag.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/flag/flag.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>
<a id="L6"></a><span class="comment">	The flag package implements command-line flag parsing.</span>

<a id="L8"></a><span class="comment">	Usage:</span>

<a id="L10"></a><span class="comment">	1) Define flags using flag.String(), Bool(), Int(), etc. Example:</span>
<a id="L11"></a><span class="comment">		import &#34;flag&#34;</span>
<a id="L12"></a><span class="comment">		var ip *int = flag.Int(&#34;flagname&#34;, 1234, &#34;help message for flagname&#34;)</span>
<a id="L13"></a><span class="comment">	If you like, you can bind the flag to a variable using the Var() functions.</span>
<a id="L14"></a><span class="comment">		var flagvar int</span>
<a id="L15"></a><span class="comment">		func init() {</span>
<a id="L16"></a><span class="comment">			flag.IntVar(&amp;flagvar, &#34;flagname&#34;, 1234, &#34;help message for flagname&#34;)</span>
<a id="L17"></a><span class="comment">		}</span>

<a id="L19"></a><span class="comment">	2) After all flags are defined, call</span>
<a id="L20"></a><span class="comment">		flag.Parse()</span>
<a id="L21"></a><span class="comment">	to parse the command line into the defined flags.</span>

<a id="L23"></a><span class="comment">	3) Flags may then be used directly. If you&#39;re using the flags themselves,</span>
<a id="L24"></a><span class="comment">	they are all pointers; if you bind to variables, they&#39;re values.</span>
<a id="L25"></a><span class="comment">		fmt.Println(&#34;ip has value &#34;, *ip);</span>
<a id="L26"></a><span class="comment">		fmt.Println(&#34;flagvar has value &#34;, flagvar);</span>

<a id="L28"></a><span class="comment">	4) After parsing, flag.Arg(i) is the i&#39;th argument after the flags.</span>
<a id="L29"></a><span class="comment">	Args are indexed from 0 up to flag.NArg().</span>

<a id="L31"></a><span class="comment">	Command line flag syntax:</span>
<a id="L32"></a><span class="comment">		-flag</span>
<a id="L33"></a><span class="comment">		-flag=x</span>
<a id="L34"></a><span class="comment">		-flag x</span>
<a id="L35"></a><span class="comment">	One or two minus signs may be used; they are equivalent.</span>

<a id="L37"></a><span class="comment">	Flag parsing stops just before the first non-flag argument</span>
<a id="L38"></a><span class="comment">	(&#34;-&#34; is a non-flag argument) or after the terminator &#34;--&#34;.</span>

<a id="L40"></a><span class="comment">	Integer flags accept 1234, 0664, 0x1234 and may be negative.</span>
<a id="L41"></a><span class="comment">	Boolean flags may be 1, 0, t, f, true, false, TRUE, FALSE, True, False.</span>
<a id="L42"></a><span class="comment">*/</span>
<a id="L43"></a>package flag

<a id="L45"></a>import (
    <a id="L46"></a>&#34;fmt&#34;;
    <a id="L47"></a>&#34;os&#34;;
    <a id="L48"></a>&#34;strconv&#34;;
<a id="L49"></a>)

<a id="L51"></a><span class="comment">// TODO(r): BUG: atob belongs elsewhere</span>
<a id="L52"></a>func atob(str string) (value bool, ok bool) {
    <a id="L53"></a>switch str {
    <a id="L54"></a>case &#34;1&#34;, &#34;t&#34;, &#34;T&#34;, &#34;true&#34;, &#34;TRUE&#34;, &#34;True&#34;:
        <a id="L55"></a>return true, true
    <a id="L56"></a>case &#34;0&#34;, &#34;f&#34;, &#34;F&#34;, &#34;false&#34;, &#34;FALSE&#34;, &#34;False&#34;:
        <a id="L57"></a>return false, true
    <a id="L58"></a>}
    <a id="L59"></a>return false, false;
<a id="L60"></a>}

<a id="L62"></a><span class="comment">// -- Bool Value</span>
<a id="L63"></a>type boolValue struct {
    <a id="L64"></a>p *bool;
<a id="L65"></a>}

<a id="L67"></a>func newBoolValue(val bool, p *bool) *boolValue {
    <a id="L68"></a>*p = val;
    <a id="L69"></a>return &amp;boolValue{p};
<a id="L70"></a>}

<a id="L72"></a>func (b *boolValue) set(s string) bool {
    <a id="L73"></a>v, ok := atob(s);
    <a id="L74"></a>*b.p = v;
    <a id="L75"></a>return ok;
<a id="L76"></a>}

<a id="L78"></a>func (b *boolValue) String() string { return fmt.Sprintf(&#34;%v&#34;, *b.p) }

<a id="L80"></a><span class="comment">// -- Int Value</span>
<a id="L81"></a>type intValue struct {
    <a id="L82"></a>p *int;
<a id="L83"></a>}

<a id="L85"></a>func newIntValue(val int, p *int) *intValue {
    <a id="L86"></a>*p = val;
    <a id="L87"></a>return &amp;intValue{p};
<a id="L88"></a>}

<a id="L90"></a>func (i *intValue) set(s string) bool {
    <a id="L91"></a>v, err := strconv.Atoi(s);
    <a id="L92"></a>*i.p = int(v);
    <a id="L93"></a>return err == nil;
<a id="L94"></a>}

<a id="L96"></a>func (i *intValue) String() string { return fmt.Sprintf(&#34;%v&#34;, *i.p) }

<a id="L98"></a><span class="comment">// -- Int64 Value</span>
<a id="L99"></a>type int64Value struct {
    <a id="L100"></a>p *int64;
<a id="L101"></a>}

<a id="L103"></a>func newInt64Value(val int64, p *int64) *int64Value {
    <a id="L104"></a>*p = val;
    <a id="L105"></a>return &amp;int64Value{p};
<a id="L106"></a>}

<a id="L108"></a>func (i *int64Value) set(s string) bool {
    <a id="L109"></a>v, err := strconv.Atoi64(s);
    <a id="L110"></a>*i.p = v;
    <a id="L111"></a>return err == nil;
<a id="L112"></a>}

<a id="L114"></a>func (i *int64Value) String() string { return fmt.Sprintf(&#34;%v&#34;, *i.p) }

<a id="L116"></a><span class="comment">// -- Uint Value</span>
<a id="L117"></a>type uintValue struct {
    <a id="L118"></a>p *uint;
<a id="L119"></a>}

<a id="L121"></a>func newUintValue(val uint, p *uint) *uintValue {
    <a id="L122"></a>*p = val;
    <a id="L123"></a>return &amp;uintValue{p};
<a id="L124"></a>}

<a id="L126"></a>func (i *uintValue) set(s string) bool {
    <a id="L127"></a>v, err := strconv.Atoui(s);
    <a id="L128"></a>*i.p = uint(v);
    <a id="L129"></a>return err == nil;
<a id="L130"></a>}

<a id="L132"></a>func (i *uintValue) String() string { return fmt.Sprintf(&#34;%v&#34;, *i.p) }

<a id="L134"></a><span class="comment">// -- uint64 Value</span>
<a id="L135"></a>type uint64Value struct {
    <a id="L136"></a>p *uint64;
<a id="L137"></a>}

<a id="L139"></a>func newUint64Value(val uint64, p *uint64) *uint64Value {
    <a id="L140"></a>*p = val;
    <a id="L141"></a>return &amp;uint64Value{p};
<a id="L142"></a>}

<a id="L144"></a>func (i *uint64Value) set(s string) bool {
    <a id="L145"></a>v, err := strconv.Atoui64(s);
    <a id="L146"></a>*i.p = uint64(v);
    <a id="L147"></a>return err == nil;
<a id="L148"></a>}

<a id="L150"></a>func (i *uint64Value) String() string { return fmt.Sprintf(&#34;%v&#34;, *i.p) }

<a id="L152"></a><span class="comment">// -- string Value</span>
<a id="L153"></a>type stringValue struct {
    <a id="L154"></a>p *string;
<a id="L155"></a>}

<a id="L157"></a>func newStringValue(val string, p *string) *stringValue {
    <a id="L158"></a>*p = val;
    <a id="L159"></a>return &amp;stringValue{p};
<a id="L160"></a>}

<a id="L162"></a>func (s *stringValue) set(val string) bool {
    <a id="L163"></a>*s.p = val;
    <a id="L164"></a>return true;
<a id="L165"></a>}

<a id="L167"></a>func (s *stringValue) String() string { return fmt.Sprintf(&#34;%s&#34;, *s.p) }

<a id="L169"></a><span class="comment">// -- Float Value</span>
<a id="L170"></a>type floatValue struct {
    <a id="L171"></a>p *float;
<a id="L172"></a>}

<a id="L174"></a>func newFloatValue(val float, p *float) *floatValue {
    <a id="L175"></a>*p = val;
    <a id="L176"></a>return &amp;floatValue{p};
<a id="L177"></a>}

<a id="L179"></a>func (f *floatValue) set(s string) bool {
    <a id="L180"></a>v, err := strconv.Atof(s);
    <a id="L181"></a>*f.p = v;
    <a id="L182"></a>return err == nil;
<a id="L183"></a>}

<a id="L185"></a>func (f *floatValue) String() string { return fmt.Sprintf(&#34;%v&#34;, *f.p) }

<a id="L187"></a><span class="comment">// -- Float64 Value</span>
<a id="L188"></a>type float64Value struct {
    <a id="L189"></a>p *float64;
<a id="L190"></a>}

<a id="L192"></a>func newFloat64Value(val float64, p *float64) *float64Value {
    <a id="L193"></a>*p = val;
    <a id="L194"></a>return &amp;float64Value{p};
<a id="L195"></a>}

<a id="L197"></a>func (f *float64Value) set(s string) bool {
    <a id="L198"></a>v, err := strconv.Atof64(s);
    <a id="L199"></a>*f.p = v;
    <a id="L200"></a>return err == nil;
<a id="L201"></a>}

<a id="L203"></a>func (f *float64Value) String() string { return fmt.Sprintf(&#34;%v&#34;, *f.p) }

<a id="L205"></a><span class="comment">// FlagValue is the interface to the dynamic value stored in a flag.</span>
<a id="L206"></a><span class="comment">// (The default value is represented as a string.)</span>
<a id="L207"></a>type FlagValue interface {
    <a id="L208"></a>String() string;
    <a id="L209"></a>set(string) bool;
<a id="L210"></a>}

<a id="L212"></a><span class="comment">// A Flag represents the state of a flag.</span>
<a id="L213"></a>type Flag struct {
    <a id="L214"></a>Name     string;    <span class="comment">// name as it appears on command line</span>
    <a id="L215"></a>Usage    string;    <span class="comment">// help message</span>
    <a id="L216"></a>Value    FlagValue; <span class="comment">// value as set</span>
    <a id="L217"></a>DefValue string;    <span class="comment">// default value (as text); for usage message</span>
<a id="L218"></a>}

<a id="L220"></a>type allFlags struct {
    <a id="L221"></a>actual    map[string]*Flag;
    <a id="L222"></a>formal    map[string]*Flag;
    <a id="L223"></a>first_arg int; <span class="comment">// 0 is the program name, 1 is first arg</span>
<a id="L224"></a>}

<a id="L226"></a>var flags *allFlags = &amp;allFlags{make(map[string]*Flag), make(map[string]*Flag), 1}

<a id="L228"></a><span class="comment">// VisitAll visits the flags, calling fn for each. It visits all flags, even those not set.</span>
<a id="L229"></a>func VisitAll(fn func(*Flag)) {
    <a id="L230"></a>for _, f := range flags.formal {
        <a id="L231"></a>fn(f)
    <a id="L232"></a>}
<a id="L233"></a>}

<a id="L235"></a><span class="comment">// Visit visits the flags, calling fn for each. It visits only those flags that have been set.</span>
<a id="L236"></a>func Visit(fn func(*Flag)) {
    <a id="L237"></a>for _, f := range flags.actual {
        <a id="L238"></a>fn(f)
    <a id="L239"></a>}
<a id="L240"></a>}

<a id="L242"></a><span class="comment">// Lookup returns the Flag structure of the named flag, returning nil if none exists.</span>
<a id="L243"></a>func Lookup(name string) *Flag {
    <a id="L244"></a>f, ok := flags.formal[name];
    <a id="L245"></a>if !ok {
        <a id="L246"></a>return nil
    <a id="L247"></a>}
    <a id="L248"></a>return f;
<a id="L249"></a>}

<a id="L251"></a><span class="comment">// Set sets the value of the named flag.  It returns true if the set succeeded; false if</span>
<a id="L252"></a><span class="comment">// there is no such flag defined.</span>
<a id="L253"></a>func Set(name, value string) bool {
    <a id="L254"></a>f, ok := flags.formal[name];
    <a id="L255"></a>if !ok {
        <a id="L256"></a>return false
    <a id="L257"></a>}
    <a id="L258"></a>ok = f.Value.set(value);
    <a id="L259"></a>if !ok {
        <a id="L260"></a>return false
    <a id="L261"></a>}
    <a id="L262"></a>flags.actual[name] = f;
    <a id="L263"></a>return true;
<a id="L264"></a>}

<a id="L266"></a><span class="comment">// PrintDefaults prints to standard error the default values of all defined flags.</span>
<a id="L267"></a>func PrintDefaults() {
    <a id="L268"></a>VisitAll(func(f *Flag) {
        <a id="L269"></a>format := &#34;  -%s=%s: %s\n&#34;;
        <a id="L270"></a>if _, ok := f.Value.(*stringValue); ok {
            <a id="L271"></a><span class="comment">// put quotes on the value</span>
            <a id="L272"></a>format = &#34;  -%s=%q: %s\n&#34;
        <a id="L273"></a>}
        <a id="L274"></a>fmt.Fprintf(os.Stderr, format, f.Name, f.DefValue, f.Usage);
    <a id="L275"></a>})
<a id="L276"></a>}

<a id="L278"></a><span class="comment">// Usage prints to standard error a default usage message documenting all defined flags.</span>
<a id="L279"></a><span class="comment">// The function is a variable that may be changed to point to a custom function.</span>
<a id="L280"></a>var Usage = func() {
    <a id="L281"></a>fmt.Fprintf(os.Stderr, &#34;Usage of %s:\n&#34;, os.Args[0]);
    <a id="L282"></a>PrintDefaults();
<a id="L283"></a>}

<a id="L285"></a>func NFlag() int { return len(flags.actual) }

<a id="L287"></a><span class="comment">// Arg returns the i&#39;th command-line argument.  Arg(0) is the first remaining argument</span>
<a id="L288"></a><span class="comment">// after flags have been processed.</span>
<a id="L289"></a>func Arg(i int) string {
    <a id="L290"></a>i += flags.first_arg;
    <a id="L291"></a>if i &lt; 0 || i &gt;= len(os.Args) {
        <a id="L292"></a>return &#34;&#34;
    <a id="L293"></a>}
    <a id="L294"></a>return os.Args[i];
<a id="L295"></a>}

<a id="L297"></a><span class="comment">// NArg is the number of arguments remaining after flags have been processed.</span>
<a id="L298"></a>func NArg() int { return len(os.Args) - flags.first_arg }

<a id="L300"></a><span class="comment">// Args returns the non-flag command-line arguments.</span>
<a id="L301"></a>func Args() []string { return os.Args[flags.first_arg:len(os.Args)] }

<a id="L303"></a>func add(name string, value FlagValue, usage string) {
    <a id="L304"></a><span class="comment">// Remember the default value as a string; it won&#39;t change.</span>
    <a id="L305"></a>f := &amp;Flag{name, usage, value, value.String()};
    <a id="L306"></a>_, alreadythere := flags.formal[name];
    <a id="L307"></a>if alreadythere {
        <a id="L308"></a>fmt.Fprintln(os.Stderr, &#34;flag redefined:&#34;, name);
        <a id="L309"></a>panic(&#34;flag redefinition&#34;); <span class="comment">// Happens only if flags are declared with identical names</span>
    <a id="L310"></a>}
    <a id="L311"></a>flags.formal[name] = f;
<a id="L312"></a>}

<a id="L314"></a><span class="comment">// BoolVar defines a bool flag with specified name, default value, and usage string.</span>
<a id="L315"></a><span class="comment">// The argument p points to a bool variable in which to store the value of the flag.</span>
<a id="L316"></a>func BoolVar(p *bool, name string, value bool, usage string) {
    <a id="L317"></a>add(name, newBoolValue(value, p), usage)
<a id="L318"></a>}

<a id="L320"></a><span class="comment">// Bool defines a bool flag with specified name, default value, and usage string.</span>
<a id="L321"></a><span class="comment">// The return value is the address of a bool variable that stores the value of the flag.</span>
<a id="L322"></a>func Bool(name string, value bool, usage string) *bool {
    <a id="L323"></a>p := new(bool);
    <a id="L324"></a>BoolVar(p, name, value, usage);
    <a id="L325"></a>return p;
<a id="L326"></a>}

<a id="L328"></a><span class="comment">// IntVar defines an int flag with specified name, default value, and usage string.</span>
<a id="L329"></a><span class="comment">// The argument p points to an int variable in which to store the value of the flag.</span>
<a id="L330"></a>func IntVar(p *int, name string, value int, usage string) {
    <a id="L331"></a>add(name, newIntValue(value, p), usage)
<a id="L332"></a>}

<a id="L334"></a><span class="comment">// Int defines an int flag with specified name, default value, and usage string.</span>
<a id="L335"></a><span class="comment">// The return value is the address of an int variable that stores the value of the flag.</span>
<a id="L336"></a>func Int(name string, value int, usage string) *int {
    <a id="L337"></a>p := new(int);
    <a id="L338"></a>IntVar(p, name, value, usage);
    <a id="L339"></a>return p;
<a id="L340"></a>}

<a id="L342"></a><span class="comment">// Int64Var defines an int64 flag with specified name, default value, and usage string.</span>
<a id="L343"></a><span class="comment">// The argument p points to an int64 variable in which to store the value of the flag.</span>
<a id="L344"></a>func Int64Var(p *int64, name string, value int64, usage string) {
    <a id="L345"></a>add(name, newInt64Value(value, p), usage)
<a id="L346"></a>}

<a id="L348"></a><span class="comment">// Int64 defines an int64 flag with specified name, default value, and usage string.</span>
<a id="L349"></a><span class="comment">// The return value is the address of an int64 variable that stores the value of the flag.</span>
<a id="L350"></a>func Int64(name string, value int64, usage string) *int64 {
    <a id="L351"></a>p := new(int64);
    <a id="L352"></a>Int64Var(p, name, value, usage);
    <a id="L353"></a>return p;
<a id="L354"></a>}

<a id="L356"></a><span class="comment">// UintVar defines a uint flag with specified name, default value, and usage string.</span>
<a id="L357"></a><span class="comment">// The argument p points to a uint variable in which to store the value of the flag.</span>
<a id="L358"></a>func UintVar(p *uint, name string, value uint, usage string) {
    <a id="L359"></a>add(name, newUintValue(value, p), usage)
<a id="L360"></a>}

<a id="L362"></a><span class="comment">// Uint defines a uint flag with specified name, default value, and usage string.</span>
<a id="L363"></a><span class="comment">// The return value is the address of a uint variable that stores the value of the flag.</span>
<a id="L364"></a>func Uint(name string, value uint, usage string) *uint {
    <a id="L365"></a>p := new(uint);
    <a id="L366"></a>UintVar(p, name, value, usage);
    <a id="L367"></a>return p;
<a id="L368"></a>}

<a id="L370"></a><span class="comment">// Uint64Var defines a uint64 flag with specified name, default value, and usage string.</span>
<a id="L371"></a><span class="comment">// The argument p points to a uint64 variable in which to store the value of the flag.</span>
<a id="L372"></a>func Uint64Var(p *uint64, name string, value uint64, usage string) {
    <a id="L373"></a>add(name, newUint64Value(value, p), usage)
<a id="L374"></a>}

<a id="L376"></a><span class="comment">// Uint64 defines a uint64 flag with specified name, default value, and usage string.</span>
<a id="L377"></a><span class="comment">// The return value is the address of a uint64 variable that stores the value of the flag.</span>
<a id="L378"></a>func Uint64(name string, value uint64, usage string) *uint64 {
    <a id="L379"></a>p := new(uint64);
    <a id="L380"></a>Uint64Var(p, name, value, usage);
    <a id="L381"></a>return p;
<a id="L382"></a>}

<a id="L384"></a><span class="comment">// StringVar defines a string flag with specified name, default value, and usage string.</span>
<a id="L385"></a><span class="comment">// The argument p points to a string variable in which to store the value of the flag.</span>
<a id="L386"></a>func StringVar(p *string, name, value string, usage string) {
    <a id="L387"></a>add(name, newStringValue(value, p), usage)
<a id="L388"></a>}

<a id="L390"></a><span class="comment">// String defines a string flag with specified name, default value, and usage string.</span>
<a id="L391"></a><span class="comment">// The return value is the address of a string variable that stores the value of the flag.</span>
<a id="L392"></a>func String(name, value string, usage string) *string {
    <a id="L393"></a>p := new(string);
    <a id="L394"></a>StringVar(p, name, value, usage);
    <a id="L395"></a>return p;
<a id="L396"></a>}

<a id="L398"></a><span class="comment">// FloatVar defines a float flag with specified name, default value, and usage string.</span>
<a id="L399"></a><span class="comment">// The argument p points to a float variable in which to store the value of the flag.</span>
<a id="L400"></a>func FloatVar(p *float, name string, value float, usage string) {
    <a id="L401"></a>add(name, newFloatValue(value, p), usage)
<a id="L402"></a>}

<a id="L404"></a><span class="comment">// Float defines a float flag with specified name, default value, and usage string.</span>
<a id="L405"></a><span class="comment">// The return value is the address of a float variable that stores the value of the flag.</span>
<a id="L406"></a>func Float(name string, value float, usage string) *float {
    <a id="L407"></a>p := new(float);
    <a id="L408"></a>FloatVar(p, name, value, usage);
    <a id="L409"></a>return p;
<a id="L410"></a>}

<a id="L412"></a><span class="comment">// Float64Var defines a float64 flag with specified name, default value, and usage string.</span>
<a id="L413"></a><span class="comment">// The argument p points to a float64 variable in which to store the value of the flag.</span>
<a id="L414"></a>func Float64Var(p *float64, name string, value float64, usage string) {
    <a id="L415"></a>add(name, newFloat64Value(value, p), usage)
<a id="L416"></a>}

<a id="L418"></a><span class="comment">// Float64 defines a float64 flag with specified name, default value, and usage string.</span>
<a id="L419"></a><span class="comment">// The return value is the address of a float64 variable that stores the value of the flag.</span>
<a id="L420"></a>func Float64(name string, value float64, usage string) *float64 {
    <a id="L421"></a>p := new(float64);
    <a id="L422"></a>Float64Var(p, name, value, usage);
    <a id="L423"></a>return p;
<a id="L424"></a>}


<a id="L427"></a>func (f *allFlags) parseOne(index int) (ok bool, next int) {
    <a id="L428"></a>s := os.Args[index];
    <a id="L429"></a>f.first_arg = index; <span class="comment">// until proven otherwise</span>
    <a id="L430"></a>if len(s) == 0 {
        <a id="L431"></a>return false, -1
    <a id="L432"></a>}
    <a id="L433"></a>if s[0] != &#39;-&#39; {
        <a id="L434"></a>return false, -1
    <a id="L435"></a>}
    <a id="L436"></a>num_minuses := 1;
    <a id="L437"></a>if len(s) == 1 {
        <a id="L438"></a>return false, index
    <a id="L439"></a>}
    <a id="L440"></a>if s[1] == &#39;-&#39; {
        <a id="L441"></a>num_minuses++;
        <a id="L442"></a>if len(s) == 2 { <span class="comment">// &#34;--&#34; terminates the flags</span>
            <a id="L443"></a>return false, index + 1
        <a id="L444"></a>}
    <a id="L445"></a>}
    <a id="L446"></a>name := s[num_minuses:len(s)];
    <a id="L447"></a>if len(name) == 0 || name[0] == &#39;-&#39; || name[0] == &#39;=&#39; {
        <a id="L448"></a>fmt.Fprintln(os.Stderr, &#34;bad flag syntax:&#34;, s);
        <a id="L449"></a>Usage();
        <a id="L450"></a>os.Exit(2);
    <a id="L451"></a>}

    <a id="L453"></a><span class="comment">// it&#39;s a flag. does it have an argument?</span>
    <a id="L454"></a>has_value := false;
    <a id="L455"></a>value := &#34;&#34;;
    <a id="L456"></a>for i := 1; i &lt; len(name); i++ { <span class="comment">// equals cannot be first</span>
        <a id="L457"></a>if name[i] == &#39;=&#39; {
            <a id="L458"></a>value = name[i+1 : len(name)];
            <a id="L459"></a>has_value = true;
            <a id="L460"></a>name = name[0:i];
            <a id="L461"></a>break;
        <a id="L462"></a>}
    <a id="L463"></a>}
    <a id="L464"></a>flag, alreadythere := flags.actual[name];
    <a id="L465"></a>if alreadythere {
        <a id="L466"></a>fmt.Fprintf(os.Stderr, &#34;flag specified twice: -%s\n&#34;, name);
        <a id="L467"></a>Usage();
        <a id="L468"></a>os.Exit(2);
    <a id="L469"></a>}
    <a id="L470"></a>m := flags.formal;
    <a id="L471"></a>flag, alreadythere = m[name]; <span class="comment">// BUG</span>
    <a id="L472"></a>if !alreadythere {
        <a id="L473"></a>fmt.Fprintf(os.Stderr, &#34;flag provided but not defined: -%s\n&#34;, name);
        <a id="L474"></a>Usage();
        <a id="L475"></a>os.Exit(2);
    <a id="L476"></a>}
    <a id="L477"></a>if f, ok := flag.Value.(*boolValue); ok { <span class="comment">// special case: doesn&#39;t need an arg</span>
        <a id="L478"></a>if has_value {
            <a id="L479"></a>if !f.set(value) {
                <a id="L480"></a>fmt.Fprintf(os.Stderr, &#34;invalid boolean value %t for flag: -%s\n&#34;, value, name);
                <a id="L481"></a>Usage();
                <a id="L482"></a>os.Exit(2);
            <a id="L483"></a>}
        <a id="L484"></a>} else {
            <a id="L485"></a>f.set(&#34;true&#34;)
        <a id="L486"></a>}
    <a id="L487"></a>} else {
        <a id="L488"></a><span class="comment">// It must have a value, which might be the next argument.</span>
        <a id="L489"></a>if !has_value &amp;&amp; index &lt; len(os.Args)-1 {
            <a id="L490"></a><span class="comment">// value is the next arg</span>
            <a id="L491"></a>has_value = true;
            <a id="L492"></a>index++;
            <a id="L493"></a>value = os.Args[index];
        <a id="L494"></a>}
        <a id="L495"></a>if !has_value {
            <a id="L496"></a>fmt.Fprintf(os.Stderr, &#34;flag needs an argument: -%s\n&#34;, name);
            <a id="L497"></a>Usage();
            <a id="L498"></a>os.Exit(2);
        <a id="L499"></a>}
        <a id="L500"></a>ok = flag.Value.set(value);
        <a id="L501"></a>if !ok {
            <a id="L502"></a>fmt.Fprintf(os.Stderr, &#34;invalid value %s for flag: -%s\n&#34;, value, name);
            <a id="L503"></a>Usage();
            <a id="L504"></a>os.Exit(2);
        <a id="L505"></a>}
    <a id="L506"></a>}
    <a id="L507"></a>flags.actual[name] = flag;
    <a id="L508"></a>return true, index + 1;
<a id="L509"></a>}

<a id="L511"></a><span class="comment">// Parse parses the command-line flags.  Must be called after all flags are defined</span>
<a id="L512"></a><span class="comment">// and before any are accessed by the program.</span>
<a id="L513"></a>func Parse() {
    <a id="L514"></a>for i := 1; i &lt; len(os.Args); {
        <a id="L515"></a>ok, next := flags.parseOne(i);
        <a id="L516"></a>if next &gt; 0 {
            <a id="L517"></a>flags.first_arg = next;
            <a id="L518"></a>i = next;
        <a id="L519"></a>}
        <a id="L520"></a>if !ok {
            <a id="L521"></a>break
        <a id="L522"></a>}
    <a id="L523"></a>}
<a id="L524"></a>}
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
