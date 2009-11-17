<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/cgo/out.go</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/cgo/out.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package main

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;go/ast&#34;;
    <a id="L10"></a>&#34;go/printer&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;strings&#34;;
<a id="L13"></a>)

<a id="L15"></a>func creat(name string) *os.File {
    <a id="L16"></a>f, err := os.Open(name, os.O_WRONLY|os.O_CREAT|os.O_TRUNC, 0666);
    <a id="L17"></a>if err != nil {
        <a id="L18"></a>fatal(&#34;%s&#34;, err)
    <a id="L19"></a>}
    <a id="L20"></a>return f;
<a id="L21"></a>}

<a id="L23"></a><span class="comment">// writeOutput creates output files to be compiled by 6g, 6c, and gcc.</span>
<a id="L24"></a><span class="comment">// (The comments here say 6g and 6c but the code applies to the 8 and 5 tools too.)</span>
<a id="L25"></a>func (p *Prog) writeOutput(srcfile string) {
    <a id="L26"></a>pkgroot := os.Getenv(&#34;GOROOT&#34;) + &#34;/pkg/&#34; + os.Getenv(&#34;GOOS&#34;) + &#34;_&#34; + os.Getenv(&#34;GOARCH&#34;);

    <a id="L28"></a>base := srcfile;
    <a id="L29"></a>if strings.HasSuffix(base, &#34;.go&#34;) {
        <a id="L30"></a>base = base[0 : len(base)-3]
    <a id="L31"></a>}
    <a id="L32"></a>fgo1 := creat(base + &#34;.cgo1.go&#34;);
    <a id="L33"></a>fgo2 := creat(base + &#34;.cgo2.go&#34;);
    <a id="L34"></a>fc := creat(base + &#34;.cgo3.c&#34;);
    <a id="L35"></a>fgcc := creat(base + &#34;.cgo4.c&#34;);

    <a id="L37"></a><span class="comment">// Write Go output: Go input with rewrites of C.xxx to _C_xxx.</span>
    <a id="L38"></a>fmt.Fprintf(fgo1, &#34;// Created by cgo - DO NOT EDIT\n&#34;);
    <a id="L39"></a>fmt.Fprintf(fgo1, &#34;//line %s:1\n&#34;, srcfile);
    <a id="L40"></a>printer.Fprint(fgo1, p.AST);

    <a id="L42"></a><span class="comment">// Write second Go output: definitions of _C_xxx.</span>
    <a id="L43"></a><span class="comment">// In a separate file so that the import of &#34;unsafe&#34; does not</span>
    <a id="L44"></a><span class="comment">// pollute the original file.</span>
    <a id="L45"></a>fmt.Fprintf(fgo2, &#34;// Created by cgo - DO NOT EDIT\n&#34;);
    <a id="L46"></a>fmt.Fprintf(fgo2, &#34;package %s\n\n&#34;, p.Package);
    <a id="L47"></a>fmt.Fprintf(fgo2, &#34;import \&#34;unsafe\&#34;\n\n&#34;);
    <a id="L48"></a>fmt.Fprintf(fgo2, &#34;type _ unsafe.Pointer\n\n&#34;);

    <a id="L50"></a>for name, def := range p.Typedef {
        <a id="L51"></a>fmt.Fprintf(fgo2, &#34;type %s &#34;, name);
        <a id="L52"></a>printer.Fprint(fgo2, def);
        <a id="L53"></a>fmt.Fprintf(fgo2, &#34;\n&#34;);
    <a id="L54"></a>}
    <a id="L55"></a>fmt.Fprintf(fgo2, &#34;type _C_void [0]byte\n&#34;);

    <a id="L57"></a><span class="comment">// While we process the vars and funcs, also write 6c and gcc output.</span>
    <a id="L58"></a><span class="comment">// Gcc output starts with the preamble.</span>
    <a id="L59"></a>fmt.Fprintf(fgcc, &#34;%s\n&#34;, p.Preamble);
    <a id="L60"></a>fmt.Fprintf(fgcc, &#34;%s\n&#34;, gccProlog);

    <a id="L62"></a>fmt.Fprintf(fc, cProlog, pkgroot, pkgroot, pkgroot, pkgroot, p.Package, p.Package);

    <a id="L64"></a>for name, def := range p.Vardef {
        <a id="L65"></a>fmt.Fprintf(fc, &#34;#pragma dynld %s·_C_%s %s \&#34;%s/%s_%s.so\&#34;\n&#34;, p.Package, name, name, pkgroot, p.PackagePath, base);
        <a id="L66"></a>fmt.Fprintf(fgo2, &#34;var _C_%s &#34;, name);
        <a id="L67"></a>printer.Fprint(fgo2, &amp;ast.StarExpr{X: def.Go});
        <a id="L68"></a>fmt.Fprintf(fgo2, &#34;\n&#34;);
    <a id="L69"></a>}
    <a id="L70"></a>fmt.Fprintf(fc, &#34;\n&#34;);

    <a id="L72"></a>for name, def := range p.Funcdef {
        <a id="L73"></a><span class="comment">// Go func declaration.</span>
        <a id="L74"></a>d := &amp;ast.FuncDecl{
            <a id="L75"></a>Name: &amp;ast.Ident{Value: &#34;_C_&#34; + name},
            <a id="L76"></a>Type: def.Go,
        <a id="L77"></a>};
        <a id="L78"></a>printer.Fprint(fgo2, d);
        <a id="L79"></a>fmt.Fprintf(fgo2, &#34;\n&#34;);

        <a id="L81"></a>if name == &#34;CString&#34; || name == &#34;GoString&#34; {
            <a id="L82"></a><span class="comment">// The builtins are already defined in the C prolog.</span>
            <a id="L83"></a>continue
        <a id="L84"></a>}

        <a id="L86"></a><span class="comment">// Construct a gcc struct matching the 6c argument frame.</span>
        <a id="L87"></a><span class="comment">// Assumes that in gcc, char is 1 byte, short 2 bytes, int 4 bytes, long long 8 bytes.</span>
        <a id="L88"></a><span class="comment">// These assumptions are checked by the gccProlog.</span>
        <a id="L89"></a><span class="comment">// Also assumes that 6c convention is to word-align the</span>
        <a id="L90"></a><span class="comment">// input and output parameters.</span>
        <a id="L91"></a>structType := &#34;struct {\n&#34;;
        <a id="L92"></a>off := int64(0);
        <a id="L93"></a>npad := 0;
        <a id="L94"></a>for i, t := range def.Params {
            <a id="L95"></a>if off%t.Align != 0 {
                <a id="L96"></a>pad := t.Align - off%t.Align;
                <a id="L97"></a>structType += fmt.Sprintf(&#34;\t\tchar __pad%d[%d];\n&#34;, npad, pad);
                <a id="L98"></a>off += pad;
                <a id="L99"></a>npad++;
            <a id="L100"></a>}
            <a id="L101"></a>structType += fmt.Sprintf(&#34;\t\t%s p%d;\n&#34;, t.C, i);
            <a id="L102"></a>off += t.Size;
        <a id="L103"></a>}
        <a id="L104"></a>if off%p.PtrSize != 0 {
            <a id="L105"></a>pad := p.PtrSize - off%p.PtrSize;
            <a id="L106"></a>structType += fmt.Sprintf(&#34;\t\tchar __pad%d[%d];\n&#34;, npad, pad);
            <a id="L107"></a>off += pad;
            <a id="L108"></a>npad++;
        <a id="L109"></a>}
        <a id="L110"></a>if t := def.Result; t != nil {
            <a id="L111"></a>if off%t.Align != 0 {
                <a id="L112"></a>pad := t.Align - off%t.Align;
                <a id="L113"></a>structType += fmt.Sprintf(&#34;\t\tchar __pad%d[%d];\n&#34;, npad, pad);
                <a id="L114"></a>off += pad;
                <a id="L115"></a>npad++;
            <a id="L116"></a>}
            <a id="L117"></a>structType += fmt.Sprintf(&#34;\t\t%s r;\n&#34;, t.C);
            <a id="L118"></a>off += t.Size;
        <a id="L119"></a>}
        <a id="L120"></a>if off%p.PtrSize != 0 {
            <a id="L121"></a>pad := p.PtrSize - off%p.PtrSize;
            <a id="L122"></a>structType += fmt.Sprintf(&#34;\t\tchar __pad%d[%d];\n&#34;, npad, pad);
            <a id="L123"></a>off += pad;
            <a id="L124"></a>npad++;
        <a id="L125"></a>}
        <a id="L126"></a>if len(def.Params) == 0 &amp;&amp; def.Result == nil {
            <a id="L127"></a>structType += &#34;\t\tchar unused;\n&#34;; <span class="comment">// avoid empty struct</span>
            <a id="L128"></a>off++;
        <a id="L129"></a>}
        <a id="L130"></a>structType += &#34;\t}&#34;;
        <a id="L131"></a>argSize := off;

        <a id="L133"></a><span class="comment">// C wrapper calls into gcc, passing a pointer to the argument frame.</span>
        <a id="L134"></a><span class="comment">// Also emit #pragma to get a pointer to the gcc wrapper.</span>
        <a id="L135"></a>fmt.Fprintf(fc, &#34;#pragma dynld _cgo_%s _cgo_%s \&#34;%s/%s_%s.so\&#34;\n&#34;, name, name, pkgroot, p.PackagePath, base);
        <a id="L136"></a>fmt.Fprintf(fc, &#34;void (*_cgo_%s)(void*);\n&#34;, name);
        <a id="L137"></a>fmt.Fprintf(fc, &#34;\n&#34;);
        <a id="L138"></a>fmt.Fprintf(fc, &#34;void\n&#34;);
        <a id="L139"></a>fmt.Fprintf(fc, &#34;%s·_C_%s(struct{uint8 x[%d];}p)\n&#34;, p.Package, name, argSize);
        <a id="L140"></a>fmt.Fprintf(fc, &#34;{\n&#34;);
        <a id="L141"></a>fmt.Fprintf(fc, &#34;\tcgocall(_cgo_%s, &amp;p);\n&#34;, name);
        <a id="L142"></a>fmt.Fprintf(fc, &#34;}\n&#34;);
        <a id="L143"></a>fmt.Fprintf(fc, &#34;\n&#34;);

        <a id="L145"></a><span class="comment">// Gcc wrapper unpacks the C argument struct</span>
        <a id="L146"></a><span class="comment">// and calls the actual C function.</span>
        <a id="L147"></a>fmt.Fprintf(fgcc, &#34;void\n&#34;);
        <a id="L148"></a>fmt.Fprintf(fgcc, &#34;_cgo_%s(void *v)\n&#34;, name);
        <a id="L149"></a>fmt.Fprintf(fgcc, &#34;{\n&#34;);
        <a id="L150"></a>fmt.Fprintf(fgcc, &#34;\t%s *a = v;\n&#34;, structType);
        <a id="L151"></a>fmt.Fprintf(fgcc, &#34;\t&#34;);
        <a id="L152"></a>if def.Result != nil {
            <a id="L153"></a>fmt.Fprintf(fgcc, &#34;a-&gt;r = &#34;)
        <a id="L154"></a>}
        <a id="L155"></a>fmt.Fprintf(fgcc, &#34;%s(&#34;, name);
        <a id="L156"></a>for i := range def.Params {
            <a id="L157"></a>if i &gt; 0 {
                <a id="L158"></a>fmt.Fprintf(fgcc, &#34;, &#34;)
            <a id="L159"></a>}
            <a id="L160"></a>fmt.Fprintf(fgcc, &#34;a-&gt;p%d&#34;, i);
        <a id="L161"></a>}
        <a id="L162"></a>fmt.Fprintf(fgcc, &#34;);\n&#34;);
        <a id="L163"></a>fmt.Fprintf(fgcc, &#34;}\n&#34;);
        <a id="L164"></a>fmt.Fprintf(fgcc, &#34;\n&#34;);
    <a id="L165"></a>}

    <a id="L167"></a>fgo1.Close();
    <a id="L168"></a>fgo2.Close();
    <a id="L169"></a>fc.Close();
    <a id="L170"></a>fgcc.Close();
<a id="L171"></a>}

<a id="L173"></a>const gccProlog = `
// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention &#34;name&#34;.
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)
`

<a id="L191"></a>const builtinProlog = `
typedef struct { char *p; int n; } _GoString_;
_GoString_ GoString(char *p);
char *CString(_GoString_);
`

<a id="L197"></a>const cProlog = `
#include &#34;runtime.h&#34;
#include &#34;cgocall.h&#34;

#pragma dynld initcgo initcgo &#34;%s/libcgo.so&#34;
#pragma dynld libcgo_thread_start libcgo_thread_start &#34;%s/libcgo.so&#34;
#pragma dynld _cgo_malloc _cgo_malloc &#34;%s/libcgo.so&#34;
#pragma dynld _cgo_free free &#34;%s/libcgo.so&#34;

void
%s·_C_GoString(int8 *p, String s)
{
	s = gostring((byte*)p);
	FLUSH(&amp;s);
}

void
%s·_C_CString(String s, int8 *p)
{
	p = cmalloc(s.len+1);
	mcpy((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&amp;p);
}
`
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
