<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/cgo/gcc.go</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/cgo/gcc.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Annotate Crefs in Prog with C types by parsing gcc debug output.</span>
<a id="L6"></a><span class="comment">// Conversion of debug output to Go types.</span>

<a id="L8"></a>package main

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bytes&#34;;
    <a id="L12"></a>&#34;debug/dwarf&#34;;
    <a id="L13"></a>&#34;debug/elf&#34;;
    <a id="L14"></a>&#34;debug/macho&#34;;
    <a id="L15"></a>&#34;fmt&#34;;
    <a id="L16"></a>&#34;go/ast&#34;;
    <a id="L17"></a>&#34;go/token&#34;;
    <a id="L18"></a>&#34;os&#34;;
    <a id="L19"></a>&#34;strconv&#34;;
    <a id="L20"></a>&#34;strings&#34;;
<a id="L21"></a>)

<a id="L23"></a>func (p *Prog) loadDebugInfo() {
    <a id="L24"></a><span class="comment">// Construct a slice of unique names from p.Crefs.</span>
    <a id="L25"></a>m := make(map[string]int);
    <a id="L26"></a>for _, c := range p.Crefs {
        <a id="L27"></a>m[c.Name] = -1
    <a id="L28"></a>}
    <a id="L29"></a>names := make([]string, 0, len(m));
    <a id="L30"></a>for name, _ := range m {
        <a id="L31"></a>i := len(names);
        <a id="L32"></a>names = names[0 : i+1];
        <a id="L33"></a>names[i] = name;
        <a id="L34"></a>m[name] = i;
    <a id="L35"></a>}

    <a id="L37"></a><span class="comment">// Coerce gcc into telling us whether each name is</span>
    <a id="L38"></a><span class="comment">// a type, a value, or undeclared.  We compile a function</span>
    <a id="L39"></a><span class="comment">// containing the line:</span>
    <a id="L40"></a><span class="comment">//	name;</span>
    <a id="L41"></a><span class="comment">// If name is a type, gcc will print:</span>
    <a id="L42"></a><span class="comment">//	x.c:2: warning: useless type name in empty declaration</span>
    <a id="L43"></a><span class="comment">// If name is a value, gcc will print</span>
    <a id="L44"></a><span class="comment">//	x.c:2: warning: statement with no effect</span>
    <a id="L45"></a><span class="comment">// If name is undeclared, gcc will print</span>
    <a id="L46"></a><span class="comment">//	x.c:2: error: &#39;name&#39; undeclared (first use in this function)</span>
    <a id="L47"></a><span class="comment">// A line number directive causes the line number to</span>
    <a id="L48"></a><span class="comment">// correspond to the index in the names array.</span>
    <a id="L49"></a>var b bytes.Buffer;
    <a id="L50"></a>b.WriteString(p.Preamble);
    <a id="L51"></a>b.WriteString(&#34;void f(void) {\n&#34;);
    <a id="L52"></a>b.WriteString(&#34;#line 0 \&#34;cgo-test\&#34;\n&#34;);
    <a id="L53"></a>for _, n := range names {
        <a id="L54"></a>b.WriteString(n);
        <a id="L55"></a>b.WriteString(&#34;;\n&#34;);
    <a id="L56"></a>}
    <a id="L57"></a>b.WriteString(&#34;}\n&#34;);

    <a id="L59"></a>kind := make(map[string]string);
    <a id="L60"></a>_, stderr := p.gccDebug(b.Bytes());
    <a id="L61"></a>if stderr == &#34;&#34; {
        <a id="L62"></a>fatal(&#34;gcc produced no output&#34;)
    <a id="L63"></a>}
    <a id="L64"></a>for _, line := range strings.Split(stderr, &#34;\n&#34;, 0) {
        <a id="L65"></a>if len(line) &lt; 9 || line[0:9] != &#34;cgo-test:&#34; {
            <a id="L66"></a>continue
        <a id="L67"></a>}
        <a id="L68"></a>line = line[9:len(line)];
        <a id="L69"></a>colon := strings.Index(line, &#34;:&#34;);
        <a id="L70"></a>if colon &lt; 0 {
            <a id="L71"></a>continue
        <a id="L72"></a>}
        <a id="L73"></a>i, err := strconv.Atoi(line[0:colon]);
        <a id="L74"></a>if err != nil {
            <a id="L75"></a>continue
        <a id="L76"></a>}
        <a id="L77"></a>what := &#34;&#34;;
        <a id="L78"></a>switch {
        <a id="L79"></a>default:
            <a id="L80"></a>continue
        <a id="L81"></a>case strings.Index(line, &#34;: useless type name in empty declaration&#34;) &gt;= 0:
            <a id="L82"></a>what = &#34;type&#34;
        <a id="L83"></a>case strings.Index(line, &#34;: statement with no effect&#34;) &gt;= 0:
            <a id="L84"></a>what = &#34;value&#34;
        <a id="L85"></a>case strings.Index(line, &#34;undeclared&#34;) &gt;= 0:
            <a id="L86"></a>what = &#34;error&#34;
        <a id="L87"></a>}
        <a id="L88"></a>if old, ok := kind[names[i]]; ok &amp;&amp; old != what {
            <a id="L89"></a>error(noPos, &#34;inconsistent gcc output about C.%s&#34;, names[i])
        <a id="L90"></a>}
        <a id="L91"></a>kind[names[i]] = what;
    <a id="L92"></a>}
    <a id="L93"></a>for _, n := range names {
        <a id="L94"></a>if _, ok := kind[n]; !ok {
            <a id="L95"></a>error(noPos, &#34;could not determine kind of name for C.%s&#34;, n)
        <a id="L96"></a>}
    <a id="L97"></a>}

    <a id="L99"></a>if nerrors &gt; 0 {
        <a id="L100"></a>fatal(&#34;failed to interpret gcc output:\n%s&#34;, stderr)
    <a id="L101"></a>}

    <a id="L103"></a><span class="comment">// Extract the types from the DWARF section of an object</span>
    <a id="L104"></a><span class="comment">// from a well-formed C program.  Gcc only generates DWARF info</span>
    <a id="L105"></a><span class="comment">// for symbols in the object file, so it is not enough to print the</span>
    <a id="L106"></a><span class="comment">// preamble and hope the symbols we care about will be there.</span>
    <a id="L107"></a><span class="comment">// Instead, emit</span>
    <a id="L108"></a><span class="comment">//	typeof(names[i]) *__cgo__i;</span>
    <a id="L109"></a><span class="comment">// for each entry in names and then dereference the type we</span>
    <a id="L110"></a><span class="comment">// learn for __cgo__i.</span>
    <a id="L111"></a>b.Reset();
    <a id="L112"></a>b.WriteString(p.Preamble);
    <a id="L113"></a>for i, n := range names {
        <a id="L114"></a>fmt.Fprintf(&amp;b, &#34;typeof(%s) *__cgo__%d;\n&#34;, n, i)
    <a id="L115"></a>}
    <a id="L116"></a>d, stderr := p.gccDebug(b.Bytes());
    <a id="L117"></a>if d == nil {
        <a id="L118"></a>fatal(&#34;gcc failed:\n%s\non input:\n%s&#34;, stderr, b.Bytes())
    <a id="L119"></a>}

    <a id="L121"></a><span class="comment">// Scan DWARF info for top-level TagVariable entries with AttrName __cgo__i.</span>
    <a id="L122"></a>types := make([]dwarf.Type, len(names));
    <a id="L123"></a>r := d.Reader();
    <a id="L124"></a>for {
        <a id="L125"></a>e, err := r.Next();
        <a id="L126"></a>if err != nil {
            <a id="L127"></a>fatal(&#34;reading DWARF entry: %s&#34;, err)
        <a id="L128"></a>}
        <a id="L129"></a>if e == nil {
            <a id="L130"></a>break
        <a id="L131"></a>}
        <a id="L132"></a>if e.Tag != dwarf.TagVariable {
            <a id="L133"></a>goto Continue
        <a id="L134"></a>}
        <a id="L135"></a>name, _ := e.Val(dwarf.AttrName).(string);
        <a id="L136"></a>typOff, _ := e.Val(dwarf.AttrType).(dwarf.Offset);
        <a id="L137"></a>if name == &#34;&#34; || typOff == 0 {
            <a id="L138"></a>fatal(&#34;malformed DWARF TagVariable entry&#34;)
        <a id="L139"></a>}
        <a id="L140"></a>if !strings.HasPrefix(name, &#34;__cgo__&#34;) {
            <a id="L141"></a>goto Continue
        <a id="L142"></a>}
        <a id="L143"></a>typ, err := d.Type(typOff);
        <a id="L144"></a>if err != nil {
            <a id="L145"></a>fatal(&#34;loading DWARF type: %s&#34;, err)
        <a id="L146"></a>}
        <a id="L147"></a>t, ok := typ.(*dwarf.PtrType);
        <a id="L148"></a>if !ok || t == nil {
            <a id="L149"></a>fatal(&#34;internal error: %s has non-pointer type&#34;, name)
        <a id="L150"></a>}
        <a id="L151"></a>i, err := strconv.Atoi(name[7:len(name)]);
        <a id="L152"></a>if err != nil {
            <a id="L153"></a>fatal(&#34;malformed __cgo__ name: %s&#34;, name)
        <a id="L154"></a>}
        <a id="L155"></a>types[i] = t.Type;

    <a id="L157"></a>Continue:
        <a id="L158"></a>if e.Tag != dwarf.TagCompileUnit {
            <a id="L159"></a>r.SkipChildren()
        <a id="L160"></a>}
    <a id="L161"></a>}

    <a id="L163"></a><span class="comment">// Record types and typedef information in Crefs.</span>
    <a id="L164"></a>var conv typeConv;
    <a id="L165"></a>conv.Init(p.PtrSize);
    <a id="L166"></a>for _, c := range p.Crefs {
        <a id="L167"></a>i := m[c.Name];
        <a id="L168"></a>c.TypeName = kind[c.Name] == &#34;type&#34;;
        <a id="L169"></a>f, fok := types[i].(*dwarf.FuncType);
        <a id="L170"></a>if c.Context == &#34;call&#34; &amp;&amp; !c.TypeName &amp;&amp; fok {
            <a id="L171"></a>c.FuncType = conv.FuncType(f)
        <a id="L172"></a>} else {
            <a id="L173"></a>c.Type = conv.Type(types[i])
        <a id="L174"></a>}
    <a id="L175"></a>}
    <a id="L176"></a>p.Typedef = conv.typedef;
<a id="L177"></a>}

<a id="L179"></a>func concat(a, b []string) []string {
    <a id="L180"></a>c := make([]string, len(a)+len(b));
    <a id="L181"></a>for i, s := range a {
        <a id="L182"></a>c[i] = s
    <a id="L183"></a>}
    <a id="L184"></a>for i, s := range b {
        <a id="L185"></a>c[i+len(a)] = s
    <a id="L186"></a>}
    <a id="L187"></a>return c;
<a id="L188"></a>}

<a id="L190"></a><span class="comment">// gccDebug runs gcc -gdwarf-2 over the C program stdin and</span>
<a id="L191"></a><span class="comment">// returns the corresponding DWARF data and any messages</span>
<a id="L192"></a><span class="comment">// printed to standard error.</span>
<a id="L193"></a>func (p *Prog) gccDebug(stdin []byte) (*dwarf.Data, string) {
    <a id="L194"></a>machine := &#34;-m32&#34;;
    <a id="L195"></a>if p.PtrSize == 8 {
        <a id="L196"></a>machine = &#34;-m64&#34;
    <a id="L197"></a>}

    <a id="L199"></a>tmp := &#34;_cgo_.o&#34;;
    <a id="L200"></a>base := []string{
        <a id="L201"></a>&#34;gcc&#34;,
        <a id="L202"></a>machine,
        <a id="L203"></a>&#34;-Wall&#34;, <span class="comment">// many warnings</span>
        <a id="L204"></a>&#34;-Werror&#34;, <span class="comment">// warnings are errors</span>
        <a id="L205"></a>&#34;-o&#34; + tmp, <span class="comment">// write object to tmp</span>
        <a id="L206"></a>&#34;-gdwarf-2&#34;, <span class="comment">// generate DWARF v2 debugging symbols</span>
        <a id="L207"></a>&#34;-c&#34;, <span class="comment">// do not link</span>
        <a id="L208"></a>&#34;-xc&#34;, <span class="comment">// input language is C</span>
        <a id="L209"></a>&#34;-&#34;, <span class="comment">// read input from standard input</span>
    <a id="L210"></a>};
    <a id="L211"></a>_, stderr, ok := run(stdin, concat(base, p.GccOptions));
    <a id="L212"></a>if !ok {
        <a id="L213"></a>return nil, string(stderr)
    <a id="L214"></a>}

    <a id="L216"></a><span class="comment">// Try to parse f as ELF and Mach-O and hope one works.</span>
    <a id="L217"></a>var f interface {
        <a id="L218"></a>DWARF() (*dwarf.Data, os.Error);
    <a id="L219"></a>}
    <a id="L220"></a>var err os.Error;
    <a id="L221"></a>if f, err = elf.Open(tmp); err != nil {
        <a id="L222"></a>if f, err = macho.Open(tmp); err != nil {
            <a id="L223"></a>fatal(&#34;cannot parse gcc output %s as ELF or Mach-O object&#34;, tmp)
        <a id="L224"></a>}
    <a id="L225"></a>}

    <a id="L227"></a>d, err := f.DWARF();
    <a id="L228"></a>if err != nil {
        <a id="L229"></a>fatal(&#34;cannot load DWARF debug information from %s: %s&#34;, tmp, err)
    <a id="L230"></a>}
    <a id="L231"></a>return d, &#34;&#34;;
<a id="L232"></a>}

<a id="L234"></a><span class="comment">// A typeConv is a translator from dwarf types to Go types</span>
<a id="L235"></a><span class="comment">// with equivalent memory layout.</span>
<a id="L236"></a>type typeConv struct {
    <a id="L237"></a><span class="comment">// Cache of already-translated or in-progress types.</span>
    <a id="L238"></a>m       map[dwarf.Type]*Type;
    <a id="L239"></a>typedef map[string]ast.Expr;

    <a id="L241"></a><span class="comment">// Predeclared types.</span>
    <a id="L242"></a>byte                                   ast.Expr; <span class="comment">// denotes padding</span>
    <a id="L243"></a>int8, int16, int32, int64              ast.Expr;
    <a id="L244"></a>uint8, uint16, uint32, uint64, uintptr ast.Expr;
    <a id="L245"></a>float32, float64                       ast.Expr;
    <a id="L246"></a>void                                   ast.Expr;
    <a id="L247"></a>unsafePointer                          ast.Expr;
    <a id="L248"></a>string                                 ast.Expr;

    <a id="L250"></a>ptrSize int64;

    <a id="L252"></a>tagGen int;
<a id="L253"></a>}

<a id="L255"></a>func (c *typeConv) Init(ptrSize int64) {
    <a id="L256"></a>c.ptrSize = ptrSize;
    <a id="L257"></a>c.m = make(map[dwarf.Type]*Type);
    <a id="L258"></a>c.typedef = make(map[string]ast.Expr);
    <a id="L259"></a>c.byte = c.Ident(&#34;byte&#34;);
    <a id="L260"></a>c.int8 = c.Ident(&#34;int8&#34;);
    <a id="L261"></a>c.int16 = c.Ident(&#34;int16&#34;);
    <a id="L262"></a>c.int32 = c.Ident(&#34;int32&#34;);
    <a id="L263"></a>c.int64 = c.Ident(&#34;int64&#34;);
    <a id="L264"></a>c.uint8 = c.Ident(&#34;uint8&#34;);
    <a id="L265"></a>c.uint16 = c.Ident(&#34;uint16&#34;);
    <a id="L266"></a>c.uint32 = c.Ident(&#34;uint32&#34;);
    <a id="L267"></a>c.uint64 = c.Ident(&#34;uint64&#34;);
    <a id="L268"></a>c.uintptr = c.Ident(&#34;uintptr&#34;);
    <a id="L269"></a>c.float32 = c.Ident(&#34;float32&#34;);
    <a id="L270"></a>c.float64 = c.Ident(&#34;float64&#34;);
    <a id="L271"></a>c.unsafePointer = c.Ident(&#34;unsafe.Pointer&#34;);
    <a id="L272"></a>c.void = c.Ident(&#34;void&#34;);
    <a id="L273"></a>c.string = c.Ident(&#34;string&#34;);
<a id="L274"></a>}

<a id="L276"></a><span class="comment">// base strips away qualifiers and typedefs to get the underlying type</span>
<a id="L277"></a>func base(dt dwarf.Type) dwarf.Type {
    <a id="L278"></a>for {
        <a id="L279"></a>if d, ok := dt.(*dwarf.QualType); ok {
            <a id="L280"></a>dt = d.Type;
            <a id="L281"></a>continue;
        <a id="L282"></a>}
        <a id="L283"></a>if d, ok := dt.(*dwarf.TypedefType); ok {
            <a id="L284"></a>dt = d.Type;
            <a id="L285"></a>continue;
        <a id="L286"></a>}
        <a id="L287"></a>break;
    <a id="L288"></a>}
    <a id="L289"></a>return dt;
<a id="L290"></a>}

<a id="L292"></a><span class="comment">// Map from dwarf text names to aliases we use in package &#34;C&#34;.</span>
<a id="L293"></a>var cnameMap = map[string]string{
    <a id="L294"></a>&#34;long int&#34;: &#34;long&#34;,
    <a id="L295"></a>&#34;long unsigned int&#34;: &#34;ulong&#34;,
    <a id="L296"></a>&#34;unsigned int&#34;: &#34;uint&#34;,
    <a id="L297"></a>&#34;short unsigned int&#34;: &#34;ushort&#34;,
    <a id="L298"></a>&#34;short int&#34;: &#34;short&#34;,
    <a id="L299"></a>&#34;long long int&#34;: &#34;longlong&#34;,
    <a id="L300"></a>&#34;long long unsigned int&#34;: &#34;ulonglong&#34;,
    <a id="L301"></a>&#34;signed char&#34;: &#34;schar&#34;,
<a id="L302"></a>}

<a id="L304"></a><span class="comment">// Type returns a *Type with the same memory layout as</span>
<a id="L305"></a><span class="comment">// dtype when used as the type of a variable or a struct field.</span>
<a id="L306"></a>func (c *typeConv) Type(dtype dwarf.Type) *Type {
    <a id="L307"></a>if t, ok := c.m[dtype]; ok {
        <a id="L308"></a>if t.Go == nil {
            <a id="L309"></a>fatal(&#34;type conversion loop at %s&#34;, dtype)
        <a id="L310"></a>}
        <a id="L311"></a>return t;
    <a id="L312"></a>}

    <a id="L314"></a>t := new(Type);
    <a id="L315"></a>t.Size = dtype.Size();
    <a id="L316"></a>t.Align = -1;
    <a id="L317"></a>t.C = dtype.Common().Name;
    <a id="L318"></a>if t.Size &lt; 0 {
        <a id="L319"></a>fatal(&#34;dwarf.Type %s reports unknown size&#34;, dtype)
    <a id="L320"></a>}

    <a id="L322"></a>c.m[dtype] = t;
    <a id="L323"></a>switch dt := dtype.(type) {
    <a id="L324"></a>default:
        <a id="L325"></a>fatal(&#34;unexpected type: %s&#34;, dtype)

    <a id="L327"></a>case *dwarf.AddrType:
        <a id="L328"></a>if t.Size != c.ptrSize {
            <a id="L329"></a>fatal(&#34;unexpected: %d-byte address type - %s&#34;, t.Size, dtype)
        <a id="L330"></a>}
        <a id="L331"></a>t.Go = c.uintptr;
        <a id="L332"></a>t.Align = t.Size;

    <a id="L334"></a>case *dwarf.ArrayType:
        <a id="L335"></a>if dt.StrideBitSize &gt; 0 {
            <a id="L336"></a><span class="comment">// Cannot represent bit-sized elements in Go.</span>
            <a id="L337"></a>t.Go = c.Opaque(t.Size);
            <a id="L338"></a>break;
        <a id="L339"></a>}
        <a id="L340"></a>gt := &amp;ast.ArrayType{
            <a id="L341"></a>Len: c.intExpr(dt.Count),
        <a id="L342"></a>};
        <a id="L343"></a>t.Go = gt; <span class="comment">// publish before recursive call</span>
        <a id="L344"></a>sub := c.Type(dt.Type);
        <a id="L345"></a>t.Align = sub.Align;
        <a id="L346"></a>gt.Elt = sub.Go;
        <a id="L347"></a>t.C = fmt.Sprintf(&#34;typeof(%s[%d])&#34;, sub.C, dt.Count);

    <a id="L349"></a>case *dwarf.CharType:
        <a id="L350"></a>if t.Size != 1 {
            <a id="L351"></a>fatal(&#34;unexpected: %d-byte char type - %s&#34;, t.Size, dtype)
        <a id="L352"></a>}
        <a id="L353"></a>t.Go = c.int8;
        <a id="L354"></a>t.Align = 1;

    <a id="L356"></a>case *dwarf.EnumType:
        <a id="L357"></a>switch t.Size {
        <a id="L358"></a>default:
            <a id="L359"></a>fatal(&#34;unexpected: %d-byte enum type - %s&#34;, t.Size, dtype)
        <a id="L360"></a>case 1:
            <a id="L361"></a>t.Go = c.uint8
        <a id="L362"></a>case 2:
            <a id="L363"></a>t.Go = c.uint16
        <a id="L364"></a>case 4:
            <a id="L365"></a>t.Go = c.uint32
        <a id="L366"></a>case 8:
            <a id="L367"></a>t.Go = c.uint64
        <a id="L368"></a>}
        <a id="L369"></a>if t.Align = t.Size; t.Align &gt;= c.ptrSize {
            <a id="L370"></a>t.Align = c.ptrSize
        <a id="L371"></a>}
        <a id="L372"></a>t.C = &#34;enum &#34; + dt.EnumName;

    <a id="L374"></a>case *dwarf.FloatType:
        <a id="L375"></a>switch t.Size {
        <a id="L376"></a>default:
            <a id="L377"></a>fatal(&#34;unexpected: %d-byte float type - %s&#34;, t.Size, dtype)
        <a id="L378"></a>case 4:
            <a id="L379"></a>t.Go = c.float32
        <a id="L380"></a>case 8:
            <a id="L381"></a>t.Go = c.float64
        <a id="L382"></a>}
        <a id="L383"></a>if t.Align = t.Size; t.Align &gt;= c.ptrSize {
            <a id="L384"></a>t.Align = c.ptrSize
        <a id="L385"></a>}

    <a id="L387"></a>case *dwarf.FuncType:
        <a id="L388"></a><span class="comment">// No attempt at translation: would enable calls</span>
        <a id="L389"></a><span class="comment">// directly between worlds, but we need to moderate those.</span>
        <a id="L390"></a>t.Go = c.uintptr;
        <a id="L391"></a>t.Align = c.ptrSize;

    <a id="L393"></a>case *dwarf.IntType:
        <a id="L394"></a>if dt.BitSize &gt; 0 {
            <a id="L395"></a>fatal(&#34;unexpected: %d-bit int type - %s&#34;, dt.BitSize, dtype)
        <a id="L396"></a>}
        <a id="L397"></a>switch t.Size {
        <a id="L398"></a>default:
            <a id="L399"></a>fatal(&#34;unexpected: %d-byte int type - %s&#34;, t.Size, dtype)
        <a id="L400"></a>case 1:
            <a id="L401"></a>t.Go = c.int8
        <a id="L402"></a>case 2:
            <a id="L403"></a>t.Go = c.int16
        <a id="L404"></a>case 4:
            <a id="L405"></a>t.Go = c.int32
        <a id="L406"></a>case 8:
            <a id="L407"></a>t.Go = c.int64
        <a id="L408"></a>}
        <a id="L409"></a>if t.Align = t.Size; t.Align &gt;= c.ptrSize {
            <a id="L410"></a>t.Align = c.ptrSize
        <a id="L411"></a>}

    <a id="L413"></a>case *dwarf.PtrType:
        <a id="L414"></a>t.Align = c.ptrSize;

        <a id="L416"></a><span class="comment">// Translate void* as unsafe.Pointer</span>
        <a id="L417"></a>if _, ok := base(dt.Type).(*dwarf.VoidType); ok {
            <a id="L418"></a>t.Go = c.unsafePointer;
            <a id="L419"></a>t.C = &#34;void*&#34;;
            <a id="L420"></a>break;
        <a id="L421"></a>}

        <a id="L423"></a>gt := &amp;ast.StarExpr{};
        <a id="L424"></a>t.Go = gt; <span class="comment">// publish before recursive call</span>
        <a id="L425"></a>sub := c.Type(dt.Type);
        <a id="L426"></a>gt.X = sub.Go;
        <a id="L427"></a>t.C = sub.C + &#34;*&#34;;

    <a id="L429"></a>case *dwarf.QualType:
        <a id="L430"></a><span class="comment">// Ignore qualifier.</span>
        <a id="L431"></a>t = c.Type(dt.Type);
        <a id="L432"></a>c.m[dtype] = t;
        <a id="L433"></a>return t;

    <a id="L435"></a>case *dwarf.StructType:
        <a id="L436"></a><span class="comment">// Convert to Go struct, being careful about alignment.</span>
        <a id="L437"></a><span class="comment">// Have to give it a name to simulate C &#34;struct foo&#34; references.</span>
        <a id="L438"></a>tag := dt.StructName;
        <a id="L439"></a>if tag == &#34;&#34; {
            <a id="L440"></a>tag = &#34;__&#34; + strconv.Itoa(c.tagGen);
            <a id="L441"></a>c.tagGen++;
        <a id="L442"></a>} else if t.C == &#34;&#34; {
            <a id="L443"></a>t.C = dt.Kind + &#34; &#34; + tag
        <a id="L444"></a>}
        <a id="L445"></a>name := c.Ident(&#34;_C&#34; + dt.Kind + &#34;_&#34; + tag);
        <a id="L446"></a>t.Go = name; <span class="comment">// publish before recursive calls</span>
        <a id="L447"></a>switch dt.Kind {
        <a id="L448"></a>case &#34;union&#34;, &#34;class&#34;:
            <a id="L449"></a>c.typedef[name.Value] = c.Opaque(t.Size);
            <a id="L450"></a>if t.C == &#34;&#34; {
                <a id="L451"></a>t.C = fmt.Sprintf(&#34;typeof(unsigned char[%d])&#34;, t.Size)
            <a id="L452"></a>}
        <a id="L453"></a>case &#34;struct&#34;:
            <a id="L454"></a>g, csyntax, align := c.Struct(dt);
            <a id="L455"></a>if t.C == &#34;&#34; {
                <a id="L456"></a>t.C = csyntax
            <a id="L457"></a>}
            <a id="L458"></a>t.Align = align;
            <a id="L459"></a>c.typedef[name.Value] = g;
        <a id="L460"></a>}

    <a id="L462"></a>case *dwarf.TypedefType:
        <a id="L463"></a><span class="comment">// Record typedef for printing.</span>
        <a id="L464"></a>if dt.Name == &#34;_GoString_&#34; {
            <a id="L465"></a><span class="comment">// Special C name for Go string type.</span>
            <a id="L466"></a><span class="comment">// Knows string layout used by compilers: pointer plus length,</span>
            <a id="L467"></a><span class="comment">// which rounds up to 2 pointers after alignment.</span>
            <a id="L468"></a>t.Go = c.string;
            <a id="L469"></a>t.Size = c.ptrSize * 2;
            <a id="L470"></a>t.Align = c.ptrSize;
            <a id="L471"></a>break;
        <a id="L472"></a>}
        <a id="L473"></a>name := c.Ident(&#34;_C_&#34; + dt.Name);
        <a id="L474"></a>t.Go = name; <span class="comment">// publish before recursive call</span>
        <a id="L475"></a>sub := c.Type(dt.Type);
        <a id="L476"></a>t.Size = sub.Size;
        <a id="L477"></a>t.Align = sub.Align;
        <a id="L478"></a>if _, ok := c.typedef[name.Value]; !ok {
            <a id="L479"></a>c.typedef[name.Value] = sub.Go
        <a id="L480"></a>}

    <a id="L482"></a>case *dwarf.UcharType:
        <a id="L483"></a>if t.Size != 1 {
            <a id="L484"></a>fatal(&#34;unexpected: %d-byte uchar type - %s&#34;, t.Size, dtype)
        <a id="L485"></a>}
        <a id="L486"></a>t.Go = c.uint8;
        <a id="L487"></a>t.Align = 1;

    <a id="L489"></a>case *dwarf.UintType:
        <a id="L490"></a>if dt.BitSize &gt; 0 {
            <a id="L491"></a>fatal(&#34;unexpected: %d-bit uint type - %s&#34;, dt.BitSize, dtype)
        <a id="L492"></a>}
        <a id="L493"></a>switch t.Size {
        <a id="L494"></a>default:
            <a id="L495"></a>fatal(&#34;unexpected: %d-byte uint type - %s&#34;, t.Size, dtype)
        <a id="L496"></a>case 1:
            <a id="L497"></a>t.Go = c.uint8
        <a id="L498"></a>case 2:
            <a id="L499"></a>t.Go = c.uint16
        <a id="L500"></a>case 4:
            <a id="L501"></a>t.Go = c.uint32
        <a id="L502"></a>case 8:
            <a id="L503"></a>t.Go = c.uint64
        <a id="L504"></a>}
        <a id="L505"></a>if t.Align = t.Size; t.Align &gt;= c.ptrSize {
            <a id="L506"></a>t.Align = c.ptrSize
        <a id="L507"></a>}

    <a id="L509"></a>case *dwarf.VoidType:
        <a id="L510"></a>t.Go = c.void;
        <a id="L511"></a>t.C = &#34;void&#34;;
    <a id="L512"></a>}

    <a id="L514"></a>switch dtype.(type) {
    <a id="L515"></a>case *dwarf.AddrType, *dwarf.CharType, *dwarf.IntType, *dwarf.FloatType, *dwarf.UcharType, *dwarf.UintType:
        <a id="L516"></a>s := dtype.Common().Name;
        <a id="L517"></a>if s != &#34;&#34; {
            <a id="L518"></a>if ss, ok := cnameMap[s]; ok {
                <a id="L519"></a>s = ss
            <a id="L520"></a>}
            <a id="L521"></a>s = strings.Join(strings.Split(s, &#34; &#34;, 0), &#34;&#34;); <span class="comment">// strip spaces</span>
            <a id="L522"></a>name := c.Ident(&#34;_C_&#34; + s);
            <a id="L523"></a>c.typedef[name.Value] = t.Go;
            <a id="L524"></a>t.Go = name;
        <a id="L525"></a>}
    <a id="L526"></a>}

    <a id="L528"></a>if t.C == &#34;&#34; {
        <a id="L529"></a>fatal(&#34;internal error: did not create C name for %s&#34;, dtype)
    <a id="L530"></a>}

    <a id="L532"></a>return t;
<a id="L533"></a>}

<a id="L535"></a><span class="comment">// FuncArg returns a Go type with the same memory layout as</span>
<a id="L536"></a><span class="comment">// dtype when used as the type of a C function argument.</span>
<a id="L537"></a>func (c *typeConv) FuncArg(dtype dwarf.Type) *Type {
    <a id="L538"></a>t := c.Type(dtype);
    <a id="L539"></a>switch dt := dtype.(type) {
    <a id="L540"></a>case *dwarf.ArrayType:
        <a id="L541"></a><span class="comment">// Arrays are passed implicitly as pointers in C.</span>
        <a id="L542"></a><span class="comment">// In Go, we must be explicit.</span>
        <a id="L543"></a>return &amp;Type{
            <a id="L544"></a>Size: c.ptrSize,
            <a id="L545"></a>Align: c.ptrSize,
            <a id="L546"></a>Go: &amp;ast.StarExpr{X: t.Go},
            <a id="L547"></a>C: t.C + &#34;*&#34;,
        <a id="L548"></a>}
    <a id="L549"></a>case *dwarf.TypedefType:
        <a id="L550"></a><span class="comment">// C has much more relaxed rules than Go for</span>
        <a id="L551"></a><span class="comment">// implicit type conversions.  When the parameter</span>
        <a id="L552"></a><span class="comment">// is type T defined as *X, simulate a little of the</span>
        <a id="L553"></a><span class="comment">// laxness of C by making the argument *X instead of T.</span>
        <a id="L554"></a>if ptr, ok := base(dt.Type).(*dwarf.PtrType); ok {
            <a id="L555"></a>return c.Type(ptr)
        <a id="L556"></a>}
    <a id="L557"></a>}
    <a id="L558"></a>return t;
<a id="L559"></a>}

<a id="L561"></a><span class="comment">// FuncType returns the Go type analogous to dtype.</span>
<a id="L562"></a><span class="comment">// There is no guarantee about matching memory layout.</span>
<a id="L563"></a>func (c *typeConv) FuncType(dtype *dwarf.FuncType) *FuncType {
    <a id="L564"></a>p := make([]*Type, len(dtype.ParamType));
    <a id="L565"></a>gp := make([]*ast.Field, len(dtype.ParamType));
    <a id="L566"></a>for i, f := range dtype.ParamType {
        <a id="L567"></a>p[i] = c.FuncArg(f);
        <a id="L568"></a>gp[i] = &amp;ast.Field{Type: p[i].Go};
    <a id="L569"></a>}
    <a id="L570"></a>var r *Type;
    <a id="L571"></a>var gr []*ast.Field;
    <a id="L572"></a>if _, ok := dtype.ReturnType.(*dwarf.VoidType); !ok &amp;&amp; dtype.ReturnType != nil {
        <a id="L573"></a>r = c.Type(dtype.ReturnType);
        <a id="L574"></a>gr = []*ast.Field{&amp;ast.Field{Type: r.Go}};
    <a id="L575"></a>}
    <a id="L576"></a>return &amp;FuncType{
        <a id="L577"></a>Params: p,
        <a id="L578"></a>Result: r,
        <a id="L579"></a>Go: &amp;ast.FuncType{
            <a id="L580"></a>Params: gp,
            <a id="L581"></a>Results: gr,
        <a id="L582"></a>},
    <a id="L583"></a>};
<a id="L584"></a>}

<a id="L586"></a><span class="comment">// Identifier</span>
<a id="L587"></a>func (c *typeConv) Ident(s string) *ast.Ident { return &amp;ast.Ident{Value: s} }

<a id="L589"></a><span class="comment">// Opaque type of n bytes.</span>
<a id="L590"></a>func (c *typeConv) Opaque(n int64) ast.Expr {
    <a id="L591"></a>return &amp;ast.ArrayType{
        <a id="L592"></a>Len: c.intExpr(n),
        <a id="L593"></a>Elt: c.byte,
    <a id="L594"></a>}
<a id="L595"></a>}

<a id="L597"></a><span class="comment">// Expr for integer n.</span>
<a id="L598"></a>func (c *typeConv) intExpr(n int64) ast.Expr {
    <a id="L599"></a>return &amp;ast.BasicLit{
        <a id="L600"></a>Kind: token.INT,
        <a id="L601"></a>Value: strings.Bytes(strconv.Itoa64(n)),
    <a id="L602"></a>}
<a id="L603"></a>}

<a id="L605"></a><span class="comment">// Add padding of given size to fld.</span>
<a id="L606"></a>func (c *typeConv) pad(fld []*ast.Field, size int64) []*ast.Field {
    <a id="L607"></a>n := len(fld);
    <a id="L608"></a>fld = fld[0 : n+1];
    <a id="L609"></a>fld[n] = &amp;ast.Field{Names: []*ast.Ident{c.Ident(&#34;_&#34;)}, Type: c.Opaque(size)};
    <a id="L610"></a>return fld;
<a id="L611"></a>}

<a id="L613"></a><span class="comment">// Struct conversion</span>
<a id="L614"></a>func (c *typeConv) Struct(dt *dwarf.StructType) (expr *ast.StructType, csyntax string, align int64) {
    <a id="L615"></a>csyntax = &#34;struct { &#34;;
    <a id="L616"></a>fld := make([]*ast.Field, 0, 2*len(dt.Field)+1); <span class="comment">// enough for padding around every field</span>
    <a id="L617"></a>off := int64(0);
    <a id="L618"></a>for _, f := range dt.Field {
        <a id="L619"></a>if f.ByteOffset &gt; off {
            <a id="L620"></a>fld = c.pad(fld, f.ByteOffset-off);
            <a id="L621"></a>off = f.ByteOffset;
        <a id="L622"></a>}
        <a id="L623"></a>t := c.Type(f.Type);
        <a id="L624"></a>n := len(fld);
        <a id="L625"></a>fld = fld[0 : n+1];
        <a id="L626"></a>fld[n] = &amp;ast.Field{Names: []*ast.Ident{c.Ident(f.Name)}, Type: t.Go};
        <a id="L627"></a>off += t.Size;
        <a id="L628"></a>csyntax += t.C + &#34; &#34; + f.Name + &#34;; &#34;;
        <a id="L629"></a>if t.Align &gt; align {
            <a id="L630"></a>align = t.Align
        <a id="L631"></a>}
    <a id="L632"></a>}
    <a id="L633"></a>if off &lt; dt.ByteSize {
        <a id="L634"></a>fld = c.pad(fld, dt.ByteSize-off);
        <a id="L635"></a>off = dt.ByteSize;
    <a id="L636"></a>}
    <a id="L637"></a>if off != dt.ByteSize {
        <a id="L638"></a>fatal(&#34;struct size calculation error&#34;)
    <a id="L639"></a>}
    <a id="L640"></a>csyntax += &#34;}&#34;;
    <a id="L641"></a>expr = &amp;ast.StructType{Fields: fld};
    <a id="L642"></a>return;
<a id="L643"></a>}
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
