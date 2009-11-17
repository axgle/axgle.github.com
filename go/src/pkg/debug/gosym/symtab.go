<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/gosym/symtab.go</title>

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
	<li>Thu Nov 12 15:46:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/gosym/symtab.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Package gosym implements access to the Go symbol</span>
<a id="L6"></a><span class="comment">// and line number tables embedded in Go binaries generated</span>
<a id="L7"></a><span class="comment">// by the gc compilers.</span>
<a id="L8"></a>package gosym

<a id="L10"></a><span class="comment">// The table format is a variant of the format used in Plan 9&#39;s a.out</span>
<a id="L11"></a><span class="comment">// format, documented at http://plan9.bell-labs.com/magic/man2html/6/a.out.</span>
<a id="L12"></a><span class="comment">// The best reference for the differences between the Plan 9 format</span>
<a id="L13"></a><span class="comment">// and the Go format is the runtime source, specifically ../../runtime/symtab.c.</span>

<a id="L15"></a>import (
    <a id="L16"></a>&#34;encoding/binary&#34;;
    <a id="L17"></a>&#34;fmt&#34;;
    <a id="L18"></a>&#34;os&#34;;
    <a id="L19"></a>&#34;strconv&#34;;
    <a id="L20"></a>&#34;strings&#34;;
<a id="L21"></a>)

<a id="L23"></a><span class="comment">/*</span>
<a id="L24"></a><span class="comment"> * Symbols</span>
<a id="L25"></a><span class="comment"> */</span>

<a id="L27"></a><span class="comment">// A Sym represents a single symbol table entry.</span>
<a id="L28"></a>type Sym struct {
    <a id="L29"></a>Value  uint64;
    <a id="L30"></a>Type   byte;
    <a id="L31"></a>Name   string;
    <a id="L32"></a>GoType uint64;
    <a id="L33"></a><span class="comment">// If this symbol if a function symbol, the corresponding Func</span>
    <a id="L34"></a>Func *Func;
<a id="L35"></a>}

<a id="L37"></a><span class="comment">// Static returns whether this symbol is static (not visible outside its file).</span>
<a id="L38"></a>func (s *Sym) Static() bool { return s.Type &gt;= &#39;a&#39; }

<a id="L40"></a><span class="comment">// PackageName returns the package part of the symbol name,</span>
<a id="L41"></a><span class="comment">// or the empty string if there is none.</span>
<a id="L42"></a>func (s *Sym) PackageName() string {
    <a id="L43"></a>if i := strings.Index(s.Name, &#34;.&#34;); i != -1 {
        <a id="L44"></a>return s.Name[0:i]
    <a id="L45"></a>}
    <a id="L46"></a>return &#34;&#34;;
<a id="L47"></a>}

<a id="L49"></a><span class="comment">// ReceiverName returns the receiver type name of this symbol,</span>
<a id="L50"></a><span class="comment">// or the empty string if there is none.</span>
<a id="L51"></a>func (s *Sym) ReceiverName() string {
    <a id="L52"></a>l := strings.Index(s.Name, &#34;.&#34;);
    <a id="L53"></a>r := strings.LastIndex(s.Name, &#34;.&#34;);
    <a id="L54"></a>if l == -1 || r == -1 || l == r {
        <a id="L55"></a>return &#34;&#34;
    <a id="L56"></a>}
    <a id="L57"></a>return s.Name[l+1 : r];
<a id="L58"></a>}

<a id="L60"></a><span class="comment">// BaseName returns the symbol name without the package or receiver name.</span>
<a id="L61"></a>func (s *Sym) BaseName() string {
    <a id="L62"></a>if i := strings.LastIndex(s.Name, &#34;.&#34;); i != -1 {
        <a id="L63"></a>return s.Name[i+1 : len(s.Name)]
    <a id="L64"></a>}
    <a id="L65"></a>return s.Name;
<a id="L66"></a>}

<a id="L68"></a><span class="comment">// A Func collects information about a single function.</span>
<a id="L69"></a>type Func struct {
    <a id="L70"></a>Entry uint64;
    <a id="L71"></a>*Sym;
    <a id="L72"></a>End       uint64;
    <a id="L73"></a>Params    []*Sym;
    <a id="L74"></a>Locals    []*Sym;
    <a id="L75"></a>FrameSize int;
    <a id="L76"></a>LineTable *LineTable;
    <a id="L77"></a>Obj       *Obj;
<a id="L78"></a>}

<a id="L80"></a><span class="comment">// An Obj represents a single object file.</span>
<a id="L81"></a>type Obj struct {
    <a id="L82"></a>Funcs []Func;
    <a id="L83"></a>Paths []Sym;
<a id="L84"></a>}

<a id="L86"></a><span class="comment">/*</span>
<a id="L87"></a><span class="comment"> * Symbol tables</span>
<a id="L88"></a><span class="comment"> */</span>

<a id="L90"></a><span class="comment">// Table represents a Go symbol table.  It stores all of the</span>
<a id="L91"></a><span class="comment">// symbols decoded from the program and provides methods to translate</span>
<a id="L92"></a><span class="comment">// between symbols, names, and addresses.</span>
<a id="L93"></a>type Table struct {
    <a id="L94"></a>Syms  []Sym;
    <a id="L95"></a>Funcs []Func;
    <a id="L96"></a>Files map[string]*Obj;
    <a id="L97"></a>Objs  []Obj;
    <a id="L98"></a><span class="comment">//	textEnd uint64;</span>
<a id="L99"></a>}

<a id="L101"></a>type sym struct {
    <a id="L102"></a>value  uint32;
    <a id="L103"></a>gotype uint32;
    <a id="L104"></a>typ    byte;
    <a id="L105"></a>name   []byte;
<a id="L106"></a>}

<a id="L108"></a>func walksymtab(data []byte, fn func(sym) os.Error) os.Error {
    <a id="L109"></a>var s sym;
    <a id="L110"></a>p := data;
    <a id="L111"></a>for len(p) &gt;= 6 {
        <a id="L112"></a>s.value = binary.BigEndian.Uint32(p[0:4]);
        <a id="L113"></a>typ := p[4];
        <a id="L114"></a>if typ&amp;0x80 == 0 {
            <a id="L115"></a>return &amp;DecodingError{len(data) - len(p) + 4, &#34;bad symbol type&#34;, typ}
        <a id="L116"></a>}
        <a id="L117"></a>typ &amp;^= 0x80;
        <a id="L118"></a>s.typ = typ;
        <a id="L119"></a>p = p[5:len(p)];
        <a id="L120"></a>var i int;
        <a id="L121"></a>var nnul int;
        <a id="L122"></a>for i = 0; i &lt; len(p); i++ {
            <a id="L123"></a>if p[i] == 0 {
                <a id="L124"></a>nnul = 1;
                <a id="L125"></a>break;
            <a id="L126"></a>}
        <a id="L127"></a>}
        <a id="L128"></a>switch typ {
        <a id="L129"></a>case &#39;z&#39;, &#39;Z&#39;:
            <a id="L130"></a>p = p[i+nnul : len(p)];
            <a id="L131"></a>for i = 0; i+2 &lt;= len(p); i += 2 {
                <a id="L132"></a>if p[i] == 0 &amp;&amp; p[i+1] == 0 {
                    <a id="L133"></a>nnul = 2;
                    <a id="L134"></a>break;
                <a id="L135"></a>}
            <a id="L136"></a>}
        <a id="L137"></a>}
        <a id="L138"></a>if i+nnul+4 &gt; len(p) {
            <a id="L139"></a>return &amp;DecodingError{len(data), &#34;unexpected EOF&#34;, nil}
        <a id="L140"></a>}
        <a id="L141"></a>s.name = p[0:i];
        <a id="L142"></a>i += nnul;
        <a id="L143"></a>s.gotype = binary.BigEndian.Uint32(p[i : i+4]);
        <a id="L144"></a>p = p[i+4 : len(p)];
        <a id="L145"></a>fn(s);
    <a id="L146"></a>}
    <a id="L147"></a>return nil;
<a id="L148"></a>}

<a id="L150"></a><span class="comment">// NewTable decodes the Go symbol table in data,</span>
<a id="L151"></a><span class="comment">// returning an in-memory representation.</span>
<a id="L152"></a>func NewTable(symtab []byte, pcln *LineTable) (*Table, os.Error) {
    <a id="L153"></a>var n int;
    <a id="L154"></a>err := walksymtab(symtab, func(s sym) os.Error {
        <a id="L155"></a>n++;
        <a id="L156"></a>return nil;
    <a id="L157"></a>});
    <a id="L158"></a>if err != nil {
        <a id="L159"></a>return nil, err
    <a id="L160"></a>}

    <a id="L162"></a>var t Table;
    <a id="L163"></a>fname := make(map[uint16]string);
    <a id="L164"></a>t.Syms = make([]Sym, 0, n);
    <a id="L165"></a>nf := 0;
    <a id="L166"></a>nz := 0;
    <a id="L167"></a>lasttyp := uint8(0);
    <a id="L168"></a>err = walksymtab(symtab, func(s sym) os.Error {
        <a id="L169"></a>n := len(t.Syms);
        <a id="L170"></a>t.Syms = t.Syms[0 : n+1];
        <a id="L171"></a>ts := &amp;t.Syms[n];
        <a id="L172"></a>ts.Type = s.typ;
        <a id="L173"></a>ts.Value = uint64(s.value);
        <a id="L174"></a>ts.GoType = uint64(s.gotype);
        <a id="L175"></a>switch s.typ {
        <a id="L176"></a>default:
            <a id="L177"></a><span class="comment">// rewrite name to use . instead of · (c2 b7)</span>
            <a id="L178"></a>w := 0;
            <a id="L179"></a>b := s.name;
            <a id="L180"></a>for i := 0; i &lt; len(b); i++ {
                <a id="L181"></a>if b[i] == 0xc2 &amp;&amp; i+1 &lt; len(b) &amp;&amp; b[i+1] == 0xb7 {
                    <a id="L182"></a>i++;
                    <a id="L183"></a>b[i] = &#39;.&#39;;
                <a id="L184"></a>}
                <a id="L185"></a>b[w] = b[i];
                <a id="L186"></a>w++;
            <a id="L187"></a>}
            <a id="L188"></a>ts.Name = string(s.name[0:w]);
        <a id="L189"></a>case &#39;z&#39;, &#39;Z&#39;:
            <a id="L190"></a>if lasttyp != &#39;z&#39; &amp;&amp; lasttyp != &#39;Z&#39; {
                <a id="L191"></a>nz++
            <a id="L192"></a>}
            <a id="L193"></a>for i := 0; i &lt; len(s.name); i += 2 {
                <a id="L194"></a>eltIdx := binary.BigEndian.Uint16(s.name[i : i+2]);
                <a id="L195"></a>elt, ok := fname[eltIdx];
                <a id="L196"></a>if !ok {
                    <a id="L197"></a>return &amp;DecodingError{-1, &#34;bad filename code&#34;, eltIdx}
                <a id="L198"></a>}
                <a id="L199"></a>if n := len(ts.Name); n &gt; 0 &amp;&amp; ts.Name[n-1] != &#39;/&#39; {
                    <a id="L200"></a>ts.Name += &#34;/&#34;
                <a id="L201"></a>}
                <a id="L202"></a>ts.Name += elt;
            <a id="L203"></a>}
        <a id="L204"></a>}
        <a id="L205"></a>switch s.typ {
        <a id="L206"></a>case &#39;T&#39;, &#39;t&#39;, &#39;L&#39;, &#39;l&#39;:
            <a id="L207"></a>nf++
        <a id="L208"></a>case &#39;f&#39;:
            <a id="L209"></a>fname[uint16(s.value)] = ts.Name
        <a id="L210"></a>}
        <a id="L211"></a>lasttyp = s.typ;
        <a id="L212"></a>return nil;
    <a id="L213"></a>});
    <a id="L214"></a>if err != nil {
        <a id="L215"></a>return nil, err
    <a id="L216"></a>}

    <a id="L218"></a>t.Funcs = make([]Func, 0, nf);
    <a id="L219"></a>t.Objs = make([]Obj, 0, nz);
    <a id="L220"></a>t.Files = make(map[string]*Obj);

    <a id="L222"></a><span class="comment">// Count text symbols and attach frame sizes, parameters, and</span>
    <a id="L223"></a><span class="comment">// locals to them.  Also, find object file boundaries.</span>
    <a id="L224"></a>var obj *Obj;
    <a id="L225"></a>lastf := 0;
    <a id="L226"></a>for i := 0; i &lt; len(t.Syms); i++ {
        <a id="L227"></a>sym := &amp;t.Syms[i];
        <a id="L228"></a>switch sym.Type {
        <a id="L229"></a>case &#39;Z&#39;, &#39;z&#39;: <span class="comment">// path symbol</span>
            <a id="L230"></a><span class="comment">// Finish the current object</span>
            <a id="L231"></a>if obj != nil {
                <a id="L232"></a>obj.Funcs = t.Funcs[lastf:len(t.Funcs)]
            <a id="L233"></a>}
            <a id="L234"></a>lastf = len(t.Funcs);

            <a id="L236"></a><span class="comment">// Start new object</span>
            <a id="L237"></a>n := len(t.Objs);
            <a id="L238"></a>t.Objs = t.Objs[0 : n+1];
            <a id="L239"></a>obj = &amp;t.Objs[n];

            <a id="L241"></a><span class="comment">// Count &amp; copy path symbols</span>
            <a id="L242"></a>var end int;
            <a id="L243"></a>for end = i + 1; end &lt; len(t.Syms); end++ {
                <a id="L244"></a>if c := t.Syms[end].Type; c != &#39;Z&#39; &amp;&amp; c != &#39;z&#39; {
                    <a id="L245"></a>break
                <a id="L246"></a>}
            <a id="L247"></a>}
            <a id="L248"></a>obj.Paths = t.Syms[i:end];
            <a id="L249"></a>i = end - 1; <span class="comment">// loop will i++</span>

            <a id="L251"></a><span class="comment">// Record file names</span>
            <a id="L252"></a>depth := 0;
            <a id="L253"></a>for j := range obj.Paths {
                <a id="L254"></a>s := &amp;obj.Paths[j];
                <a id="L255"></a>if s.Name == &#34;&#34; {
                    <a id="L256"></a>depth--
                <a id="L257"></a>} else {
                    <a id="L258"></a>if depth == 0 {
                        <a id="L259"></a>t.Files[s.Name] = obj
                    <a id="L260"></a>}
                    <a id="L261"></a>depth++;
                <a id="L262"></a>}
            <a id="L263"></a>}

        <a id="L265"></a>case &#39;T&#39;, &#39;t&#39;, &#39;L&#39;, &#39;l&#39;: <span class="comment">// text symbol</span>
            <a id="L266"></a>if n := len(t.Funcs); n &gt; 0 {
                <a id="L267"></a>t.Funcs[n-1].End = sym.Value
            <a id="L268"></a>}
            <a id="L269"></a>if sym.Name == &#34;etext&#34; {
                <a id="L270"></a>continue
            <a id="L271"></a>}

            <a id="L273"></a><span class="comment">// Count parameter and local (auto) syms</span>
            <a id="L274"></a>var np, na int;
            <a id="L275"></a>var end int;
        <a id="L276"></a>countloop:
            <a id="L277"></a>for end = i + 1; end &lt; len(t.Syms); end++ {
                <a id="L278"></a>switch t.Syms[end].Type {
                <a id="L279"></a>case &#39;T&#39;, &#39;t&#39;, &#39;L&#39;, &#39;l&#39;, &#39;Z&#39;, &#39;z&#39;:
                    <a id="L280"></a>break countloop
                <a id="L281"></a>case &#39;p&#39;:
                    <a id="L282"></a>np++
                <a id="L283"></a>case &#39;a&#39;:
                    <a id="L284"></a>na++
                <a id="L285"></a>}
            <a id="L286"></a>}

            <a id="L288"></a><span class="comment">// Fill in the function symbol</span>
            <a id="L289"></a>n := len(t.Funcs);
            <a id="L290"></a>t.Funcs = t.Funcs[0 : n+1];
            <a id="L291"></a>fn := &amp;t.Funcs[n];
            <a id="L292"></a>sym.Func = fn;
            <a id="L293"></a>fn.Params = make([]*Sym, 0, np);
            <a id="L294"></a>fn.Locals = make([]*Sym, 0, na);
            <a id="L295"></a>fn.Sym = sym;
            <a id="L296"></a>fn.Entry = sym.Value;
            <a id="L297"></a>fn.Obj = obj;
            <a id="L298"></a>if pcln != nil {
                <a id="L299"></a>fn.LineTable = pcln.slice(fn.Entry);
                <a id="L300"></a>pcln = fn.LineTable;
            <a id="L301"></a>}
            <a id="L302"></a>for j := i; j &lt; end; j++ {
                <a id="L303"></a>s := &amp;t.Syms[j];
                <a id="L304"></a>switch s.Type {
                <a id="L305"></a>case &#39;m&#39;:
                    <a id="L306"></a>fn.FrameSize = int(s.Value)
                <a id="L307"></a>case &#39;p&#39;:
                    <a id="L308"></a>n := len(fn.Params);
                    <a id="L309"></a>fn.Params = fn.Params[0 : n+1];
                    <a id="L310"></a>fn.Params[n] = s;
                <a id="L311"></a>case &#39;a&#39;:
                    <a id="L312"></a>n := len(fn.Locals);
                    <a id="L313"></a>fn.Locals = fn.Locals[0 : n+1];
                    <a id="L314"></a>fn.Locals[n] = s;
                <a id="L315"></a>}
            <a id="L316"></a>}
            <a id="L317"></a>i = end - 1; <span class="comment">// loop will i++</span>
        <a id="L318"></a>}
    <a id="L319"></a>}
    <a id="L320"></a>if obj != nil {
        <a id="L321"></a>obj.Funcs = t.Funcs[lastf:len(t.Funcs)]
    <a id="L322"></a>}
    <a id="L323"></a>return &amp;t, nil;
<a id="L324"></a>}

<a id="L326"></a><span class="comment">// PCToFunc returns the function containing the program counter pc,</span>
<a id="L327"></a><span class="comment">// or nil if there is no such function.</span>
<a id="L328"></a>func (t *Table) PCToFunc(pc uint64) *Func {
    <a id="L329"></a>funcs := t.Funcs;
    <a id="L330"></a>for len(funcs) &gt; 0 {
        <a id="L331"></a>m := len(funcs) / 2;
        <a id="L332"></a>fn := &amp;funcs[m];
        <a id="L333"></a>switch {
        <a id="L334"></a>case pc &lt; fn.Entry:
            <a id="L335"></a>funcs = funcs[0:m]
        <a id="L336"></a>case fn.Entry &lt;= pc &amp;&amp; pc &lt; fn.End:
            <a id="L337"></a>return fn
        <a id="L338"></a>default:
            <a id="L339"></a>funcs = funcs[m+1 : len(funcs)]
        <a id="L340"></a>}
    <a id="L341"></a>}
    <a id="L342"></a>return nil;
<a id="L343"></a>}

<a id="L345"></a><span class="comment">// PCToLine looks up line number information for a program counter.</span>
<a id="L346"></a><span class="comment">// If there is no information, it returns fn == nil.</span>
<a id="L347"></a>func (t *Table) PCToLine(pc uint64) (file string, line int, fn *Func) {
    <a id="L348"></a>if fn = t.PCToFunc(pc); fn == nil {
        <a id="L349"></a>return
    <a id="L350"></a>}
    <a id="L351"></a>file, line = fn.Obj.lineFromAline(fn.LineTable.PCToLine(pc));
    <a id="L352"></a>return;
<a id="L353"></a>}

<a id="L355"></a><span class="comment">// LineToPC looks up the first program counter on the given line in</span>
<a id="L356"></a><span class="comment">// the named file.  Returns UnknownPathError or UnknownLineError if</span>
<a id="L357"></a><span class="comment">// there is an error looking up this line.</span>
<a id="L358"></a>func (t *Table) LineToPC(file string, line int) (pc uint64, fn *Func, err os.Error) {
    <a id="L359"></a>obj, ok := t.Files[file];
    <a id="L360"></a>if !ok {
        <a id="L361"></a>return 0, nil, UnknownFileError(file)
    <a id="L362"></a>}
    <a id="L363"></a>abs, err := obj.alineFromLine(file, line);
    <a id="L364"></a>if err != nil {
        <a id="L365"></a>return
    <a id="L366"></a>}
    <a id="L367"></a>for i := range obj.Funcs {
        <a id="L368"></a>f := &amp;obj.Funcs[i];
        <a id="L369"></a>pc := f.LineTable.LineToPC(abs, f.End);
        <a id="L370"></a>if pc != 0 {
            <a id="L371"></a>return pc, f, nil
        <a id="L372"></a>}
    <a id="L373"></a>}
    <a id="L374"></a>return 0, nil, &amp;UnknownLineError{file, line};
<a id="L375"></a>}

<a id="L377"></a><span class="comment">// LookupSym returns the text, data, or bss symbol with the given name,</span>
<a id="L378"></a><span class="comment">// or nil if no such symbol is found.</span>
<a id="L379"></a>func (t *Table) LookupSym(name string) *Sym {
    <a id="L380"></a><span class="comment">// TODO(austin) Maybe make a map</span>
    <a id="L381"></a>for i := range t.Syms {
        <a id="L382"></a>s := &amp;t.Syms[i];
        <a id="L383"></a>switch s.Type {
        <a id="L384"></a>case &#39;T&#39;, &#39;t&#39;, &#39;L&#39;, &#39;l&#39;, &#39;D&#39;, &#39;d&#39;, &#39;B&#39;, &#39;b&#39;:
            <a id="L385"></a>if s.Name == name {
                <a id="L386"></a>return s
            <a id="L387"></a>}
        <a id="L388"></a>}
    <a id="L389"></a>}
    <a id="L390"></a>return nil;
<a id="L391"></a>}

<a id="L393"></a><span class="comment">// LookupFunc returns the text, data, or bss symbol with the given name,</span>
<a id="L394"></a><span class="comment">// or nil if no such symbol is found.</span>
<a id="L395"></a>func (t *Table) LookupFunc(name string) *Func {
    <a id="L396"></a>for i := range t.Funcs {
        <a id="L397"></a>f := &amp;t.Funcs[i];
        <a id="L398"></a>if f.Sym.Name == name {
            <a id="L399"></a>return f
        <a id="L400"></a>}
    <a id="L401"></a>}
    <a id="L402"></a>return nil;
<a id="L403"></a>}

<a id="L405"></a><span class="comment">// SymByAddr returns the text, data, or bss symbol starting at the given address.</span>
<a id="L406"></a><span class="comment">// TODO(rsc): Allow lookup by any address within the symbol.</span>
<a id="L407"></a>func (t *Table) SymByAddr(addr uint64) *Sym {
    <a id="L408"></a><span class="comment">// TODO(austin) Maybe make a map</span>
    <a id="L409"></a>for i := range t.Syms {
        <a id="L410"></a>s := &amp;t.Syms[i];
        <a id="L411"></a>switch s.Type {
        <a id="L412"></a>case &#39;T&#39;, &#39;t&#39;, &#39;L&#39;, &#39;l&#39;, &#39;D&#39;, &#39;d&#39;, &#39;B&#39;, &#39;b&#39;:
            <a id="L413"></a>if s.Value == addr {
                <a id="L414"></a>return s
            <a id="L415"></a>}
        <a id="L416"></a>}
    <a id="L417"></a>}
    <a id="L418"></a>return nil;
<a id="L419"></a>}

<a id="L421"></a><span class="comment">/*</span>
<a id="L422"></a><span class="comment"> * Object files</span>
<a id="L423"></a><span class="comment"> */</span>

<a id="L425"></a>func (o *Obj) lineFromAline(aline int) (string, int) {
    <a id="L426"></a>type stackEnt struct {
        <a id="L427"></a>path   string;
        <a id="L428"></a>start  int;
        <a id="L429"></a>offset int;
        <a id="L430"></a>prev   *stackEnt;
    <a id="L431"></a>}

    <a id="L433"></a>noPath := &amp;stackEnt{&#34;&#34;, 0, 0, nil};
    <a id="L434"></a>tos := noPath;

    <a id="L436"></a><span class="comment">// TODO(austin) I have no idea how &#39;Z&#39; symbols work, except</span>
    <a id="L437"></a><span class="comment">// that they pop the stack.</span>
<a id="L438"></a>pathloop:
    <a id="L439"></a>for _, s := range o.Paths {
        <a id="L440"></a>val := int(s.Value);
        <a id="L441"></a>switch {
        <a id="L442"></a>case val &gt; aline:
            <a id="L443"></a>break pathloop

        <a id="L445"></a>case val == 1:
            <a id="L446"></a><span class="comment">// Start a new stack</span>
            <a id="L447"></a>tos = &amp;stackEnt{s.Name, val, 0, noPath}

        <a id="L449"></a>case s.Name == &#34;&#34;:
            <a id="L450"></a><span class="comment">// Pop</span>
            <a id="L451"></a>if tos == noPath {
                <a id="L452"></a>return &#34;&lt;malformed symbol table&gt;&#34;, 0
            <a id="L453"></a>}
            <a id="L454"></a>tos.prev.offset += val - tos.start;
            <a id="L455"></a>tos = tos.prev;

        <a id="L457"></a>default:
            <a id="L458"></a><span class="comment">// Push</span>
            <a id="L459"></a>tos = &amp;stackEnt{s.Name, val, 0, tos}
        <a id="L460"></a>}
    <a id="L461"></a>}

    <a id="L463"></a>if tos == noPath {
        <a id="L464"></a>return &#34;&#34;, 0
    <a id="L465"></a>}
    <a id="L466"></a>return tos.path, aline - tos.start - tos.offset + 1;
<a id="L467"></a>}

<a id="L469"></a>func (o *Obj) alineFromLine(path string, line int) (int, os.Error) {
    <a id="L470"></a>if line &lt; 1 {
        <a id="L471"></a>return 0, &amp;UnknownLineError{path, line}
    <a id="L472"></a>}

    <a id="L474"></a>for i, s := range o.Paths {
        <a id="L475"></a><span class="comment">// Find this path</span>
        <a id="L476"></a>if s.Name != path {
            <a id="L477"></a>continue
        <a id="L478"></a>}

        <a id="L480"></a><span class="comment">// Find this line at this stack level</span>
        <a id="L481"></a>depth := 0;
        <a id="L482"></a>var incstart int;
        <a id="L483"></a>line += int(s.Value);
    <a id="L484"></a>pathloop:
        <a id="L485"></a>for _, s := range o.Paths[i:len(o.Paths)] {
            <a id="L486"></a>val := int(s.Value);
            <a id="L487"></a>switch {
            <a id="L488"></a>case depth == 1 &amp;&amp; val &gt;= line:
                <a id="L489"></a>return line - 1, nil

            <a id="L491"></a>case s.Name == &#34;&#34;:
                <a id="L492"></a>depth--;
                <a id="L493"></a>if depth == 0 {
                    <a id="L494"></a>break pathloop
                <a id="L495"></a>} else if depth == 1 {
                    <a id="L496"></a>line += val - incstart
                <a id="L497"></a>}

            <a id="L499"></a>default:
                <a id="L500"></a>if depth == 1 {
                    <a id="L501"></a>incstart = val
                <a id="L502"></a>}
                <a id="L503"></a>depth++;
            <a id="L504"></a>}
        <a id="L505"></a>}
        <a id="L506"></a>return 0, &amp;UnknownLineError{path, line};
    <a id="L507"></a>}
    <a id="L508"></a>return 0, UnknownFileError(path);
<a id="L509"></a>}

<a id="L511"></a><span class="comment">/*</span>
<a id="L512"></a><span class="comment"> * Errors</span>
<a id="L513"></a><span class="comment"> */</span>

<a id="L515"></a><span class="comment">// UnknownFileError represents a failure to find the specific file in</span>
<a id="L516"></a><span class="comment">// the symbol table.</span>
<a id="L517"></a>type UnknownFileError string

<a id="L519"></a>func (e UnknownFileError) String() string { return &#34;unknown file: &#34; + string(e) }

<a id="L521"></a><span class="comment">// UnknownLineError represents a failure to map a line to a program</span>
<a id="L522"></a><span class="comment">// counter, either because the line is beyond the bounds of the file</span>
<a id="L523"></a><span class="comment">// or because there is no code on the given line.</span>
<a id="L524"></a>type UnknownLineError struct {
    <a id="L525"></a>File string;
    <a id="L526"></a>Line int;
<a id="L527"></a>}

<a id="L529"></a>func (e *UnknownLineError) String() string {
    <a id="L530"></a>return &#34;no code at &#34; + e.File + &#34;:&#34; + strconv.Itoa(e.Line)
<a id="L531"></a>}

<a id="L533"></a><span class="comment">// DecodingError represents an error during the decoding of</span>
<a id="L534"></a><span class="comment">// the symbol table.</span>
<a id="L535"></a>type DecodingError struct {
    <a id="L536"></a>off int;
    <a id="L537"></a>msg string;
    <a id="L538"></a>val interface{};
<a id="L539"></a>}

<a id="L541"></a>func (e *DecodingError) String() string {
    <a id="L542"></a>msg := e.msg;
    <a id="L543"></a>if e.val != nil {
        <a id="L544"></a>msg += fmt.Sprintf(&#34; &#39;%v&#39;&#34;, e.val)
    <a id="L545"></a>}
    <a id="L546"></a>msg += fmt.Sprintf(&#34; at byte %#x&#34;, e.off);
    <a id="L547"></a>return msg;
<a id="L548"></a>}
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
