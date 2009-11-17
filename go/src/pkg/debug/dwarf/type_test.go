<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/dwarf/type_test.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/debug/dwarf/type_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package dwarf_test

<a id="L7"></a>import (
    <a id="L8"></a>. &#34;debug/dwarf&#34;;
    <a id="L9"></a>&#34;debug/elf&#34;;
    <a id="L10"></a>&#34;debug/macho&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a>var typedefTests = map[string]string{
    <a id="L15"></a>&#34;t_ptr_volatile_int&#34;: &#34;*volatile int&#34;,
    <a id="L16"></a>&#34;t_ptr_const_char&#34;: &#34;*const char&#34;,
    <a id="L17"></a>&#34;t_long&#34;: &#34;long int&#34;,
    <a id="L18"></a>&#34;t_ushort&#34;: &#34;short unsigned int&#34;,
    <a id="L19"></a>&#34;t_func_int_of_float_double&#34;: &#34;func(float, double) int&#34;,
    <a id="L20"></a>&#34;t_ptr_func_int_of_float_double&#34;: &#34;*func(float, double) int&#34;,
    <a id="L21"></a>&#34;t_func_ptr_int_of_char_schar_uchar&#34;: &#34;func(char, signed char, unsigned char) *int&#34;,
    <a id="L22"></a>&#34;t_func_void_of_char&#34;: &#34;func(char) void&#34;,
    <a id="L23"></a>&#34;t_func_void_of_void&#34;: &#34;func() void&#34;,
    <a id="L24"></a>&#34;t_func_void_of_ptr_char_dots&#34;: &#34;func(*char, ...) void&#34;,
    <a id="L25"></a>&#34;t_my_struct&#34;: &#34;struct my_struct {vi volatile int@0; x char@4 : 1@7; y int@4 : 4@27; array [40]long long int@8}&#34;,
    <a id="L26"></a>&#34;t_my_union&#34;: &#34;union my_union {vi volatile int@0; x char@0 : 1@7; y int@0 : 4@28; array [40]long long int@0}&#34;,
    <a id="L27"></a>&#34;t_my_enum&#34;: &#34;enum my_enum {e1=1; e2=2; e3=-5; e4=1000000000000000}&#34;,
    <a id="L28"></a>&#34;t_my_list&#34;: &#34;struct list {val short int@0; next *t_my_list@8}&#34;,
    <a id="L29"></a>&#34;t_my_tree&#34;: &#34;struct tree {left *struct tree@0; right *struct tree@8; val long long unsigned int@16}&#34;,
<a id="L30"></a>}

<a id="L32"></a>func elfData(t *testing.T, name string) *Data {
    <a id="L33"></a>f, err := elf.Open(name);
    <a id="L34"></a>if err != nil {
        <a id="L35"></a>t.Fatal(err)
    <a id="L36"></a>}

    <a id="L38"></a>d, err := f.DWARF();
    <a id="L39"></a>if err != nil {
        <a id="L40"></a>t.Fatal(err)
    <a id="L41"></a>}
    <a id="L42"></a>return d;
<a id="L43"></a>}

<a id="L45"></a>func machoData(t *testing.T, name string) *Data {
    <a id="L46"></a>f, err := macho.Open(name);
    <a id="L47"></a>if err != nil {
        <a id="L48"></a>t.Fatal(err)
    <a id="L49"></a>}

    <a id="L51"></a>d, err := f.DWARF();
    <a id="L52"></a>if err != nil {
        <a id="L53"></a>t.Fatal(err)
    <a id="L54"></a>}
    <a id="L55"></a>return d;
<a id="L56"></a>}


<a id="L59"></a>func TestTypedefsELF(t *testing.T) { testTypedefs(t, elfData(t, &#34;testdata/typedef.elf&#34;)) }

<a id="L61"></a>func TestTypedefsMachO(t *testing.T) {
    <a id="L62"></a>testTypedefs(t, machoData(t, &#34;testdata/typedef.macho&#34;))
<a id="L63"></a>}

<a id="L65"></a>func testTypedefs(t *testing.T, d *Data) {
    <a id="L66"></a>r := d.Reader();
    <a id="L67"></a>seen := make(map[string]bool);
    <a id="L68"></a>for {
        <a id="L69"></a>e, err := r.Next();
        <a id="L70"></a>if err != nil {
            <a id="L71"></a>t.Fatal(&#34;r.Next:&#34;, err)
        <a id="L72"></a>}
        <a id="L73"></a>if e == nil {
            <a id="L74"></a>break
        <a id="L75"></a>}
        <a id="L76"></a>if e.Tag == TagTypedef {
            <a id="L77"></a>typ, err := d.Type(e.Offset);
            <a id="L78"></a>if err != nil {
                <a id="L79"></a>t.Fatal(&#34;d.Type:&#34;, err)
            <a id="L80"></a>}
            <a id="L81"></a>t1 := typ.(*TypedefType);
            <a id="L82"></a>var typstr string;
            <a id="L83"></a>if ts, ok := t1.Type.(*StructType); ok {
                <a id="L84"></a>typstr = ts.Defn()
            <a id="L85"></a>} else {
                <a id="L86"></a>typstr = t1.Type.String()
            <a id="L87"></a>}

            <a id="L89"></a>if want, ok := typedefTests[t1.Name]; ok {
                <a id="L90"></a>if _, ok := seen[t1.Name]; ok {
                    <a id="L91"></a>t.Errorf(&#34;multiple definitions for %s&#34;, t1.Name)
                <a id="L92"></a>}
                <a id="L93"></a>seen[t1.Name] = true;
                <a id="L94"></a>if typstr != want {
                    <a id="L95"></a>t.Errorf(&#34;%s:\n\thave %s\n\twant %s&#34;, t1.Name, typstr, want)
                <a id="L96"></a>}
            <a id="L97"></a>}
        <a id="L98"></a>}
        <a id="L99"></a>if e.Tag != TagCompileUnit {
            <a id="L100"></a>r.SkipChildren()
        <a id="L101"></a>}
    <a id="L102"></a>}

    <a id="L104"></a>for k := range typedefTests {
        <a id="L105"></a>if _, ok := seen[k]; !ok {
            <a id="L106"></a>t.Errorf(&#34;missing %s&#34;, k)
        <a id="L107"></a>}
    <a id="L108"></a>}
<a id="L109"></a>}
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
