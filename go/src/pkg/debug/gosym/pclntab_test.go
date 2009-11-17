<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/debug/gosym/pclntab_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/debug/gosym/pclntab_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package gosym

<a id="L7"></a>import (
    <a id="L8"></a>&#34;debug/elf&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;testing&#34;;
    <a id="L11"></a>&#34;syscall&#34;;
<a id="L12"></a>)

<a id="L14"></a>func dotest() bool {
    <a id="L15"></a><span class="comment">// For now, only works on ELF platforms.</span>
    <a id="L16"></a>return syscall.OS == &#34;linux&#34; &amp;&amp; os.Getenv(&#34;GOARCH&#34;) == &#34;amd64&#34;
<a id="L17"></a>}

<a id="L19"></a>func getTable(t *testing.T) *Table {
    <a id="L20"></a>f, tab := crack(os.Args[0], t);
    <a id="L21"></a>f.Close();
    <a id="L22"></a>return tab;
<a id="L23"></a>}

<a id="L25"></a>func crack(file string, t *testing.T) (*elf.File, *Table) {
    <a id="L26"></a><span class="comment">// Open self</span>
    <a id="L27"></a>f, err := elf.Open(file);
    <a id="L28"></a>if err != nil {
        <a id="L29"></a>t.Fatal(err)
    <a id="L30"></a>}
    <a id="L31"></a>return parse(file, f, t);
<a id="L32"></a>}

<a id="L34"></a>func parse(file string, f *elf.File, t *testing.T) (*elf.File, *Table) {
    <a id="L35"></a>symdat, err := f.Section(&#34;.gosymtab&#34;).Data();
    <a id="L36"></a>if err != nil {
        <a id="L37"></a>f.Close();
        <a id="L38"></a>t.Fatalf(&#34;reading %s gosymtab: %v&#34;, file, err);
    <a id="L39"></a>}
    <a id="L40"></a>pclndat, err := f.Section(&#34;.gopclntab&#34;).Data();
    <a id="L41"></a>if err != nil {
        <a id="L42"></a>f.Close();
        <a id="L43"></a>t.Fatalf(&#34;reading %s gopclntab: %v&#34;, file, err);
    <a id="L44"></a>}

    <a id="L46"></a>pcln := NewLineTable(pclndat, f.Section(&#34;.text&#34;).Addr);
    <a id="L47"></a>tab, err := NewTable(symdat, pcln);
    <a id="L48"></a>if err != nil {
        <a id="L49"></a>f.Close();
        <a id="L50"></a>t.Fatalf(&#34;parsing %s gosymtab: %v&#34;, file, err);
    <a id="L51"></a>}

    <a id="L53"></a>return f, tab;
<a id="L54"></a>}

<a id="L56"></a>var goarch = os.Getenv(&#34;O&#34;)

<a id="L58"></a>func TestLineFromAline(t *testing.T) {
    <a id="L59"></a>if !dotest() {
        <a id="L60"></a>return
    <a id="L61"></a>}

    <a id="L63"></a>tab := getTable(t);

    <a id="L65"></a><span class="comment">// Find the sym package</span>
    <a id="L66"></a>pkg := tab.LookupFunc(&#34;gosym.TestLineFromAline&#34;).Obj;
    <a id="L67"></a>if pkg == nil {
        <a id="L68"></a>t.Fatalf(&#34;nil pkg&#34;)
    <a id="L69"></a>}

    <a id="L71"></a><span class="comment">// Walk every absolute line and ensure that we hit every</span>
    <a id="L72"></a><span class="comment">// source line monotonically</span>
    <a id="L73"></a>lastline := make(map[string]int);
    <a id="L74"></a>final := -1;
    <a id="L75"></a>for i := 0; i &lt; 10000; i++ {
        <a id="L76"></a>path, line := pkg.lineFromAline(i);
        <a id="L77"></a><span class="comment">// Check for end of object</span>
        <a id="L78"></a>if path == &#34;&#34; {
            <a id="L79"></a>if final == -1 {
                <a id="L80"></a>final = i - 1
            <a id="L81"></a>}
            <a id="L82"></a>continue;
        <a id="L83"></a>} else if final != -1 {
            <a id="L84"></a>t.Fatalf(&#34;reached end of package at absolute line %d, but absolute line %d mapped to %s:%d&#34;, final, i, path, line)
        <a id="L85"></a>}
        <a id="L86"></a><span class="comment">// It&#39;s okay to see files multiple times (e.g., sys.a)</span>
        <a id="L87"></a>if line == 1 {
            <a id="L88"></a>lastline[path] = 1;
            <a id="L89"></a>continue;
        <a id="L90"></a>}
        <a id="L91"></a><span class="comment">// Check that the is the next line in path</span>
        <a id="L92"></a>ll, ok := lastline[path];
        <a id="L93"></a>if !ok {
            <a id="L94"></a>t.Errorf(&#34;file %s starts on line %d&#34;, path, line)
        <a id="L95"></a>} else if line != ll+1 {
            <a id="L96"></a>t.Errorf(&#34;expected next line of file %s to be %d, got %d&#34;, path, ll+1, line)
        <a id="L97"></a>}
        <a id="L98"></a>lastline[path] = line;
    <a id="L99"></a>}
    <a id="L100"></a>if final == -1 {
        <a id="L101"></a>t.Errorf(&#34;never reached end of object&#34;)
    <a id="L102"></a>}
<a id="L103"></a>}

<a id="L105"></a>func TestLineAline(t *testing.T) {
    <a id="L106"></a>if !dotest() {
        <a id="L107"></a>return
    <a id="L108"></a>}

    <a id="L110"></a>tab := getTable(t);

    <a id="L112"></a>for _, o := range tab.Files {
        <a id="L113"></a><span class="comment">// A source file can appear multiple times in a</span>
        <a id="L114"></a><span class="comment">// object.  alineFromLine will always return alines in</span>
        <a id="L115"></a><span class="comment">// the first file, so track which lines we&#39;ve seen.</span>
        <a id="L116"></a>found := make(map[string]int);
        <a id="L117"></a>for i := 0; i &lt; 1000; i++ {
            <a id="L118"></a>path, line := o.lineFromAline(i);
            <a id="L119"></a>if path == &#34;&#34; {
                <a id="L120"></a>break
            <a id="L121"></a>}

            <a id="L123"></a><span class="comment">// cgo files are full of &#39;Z&#39; symbols, which we don&#39;t handle</span>
            <a id="L124"></a>if len(path) &gt; 4 &amp;&amp; path[len(path)-4:len(path)] == &#34;.cgo&#34; {
                <a id="L125"></a>continue
            <a id="L126"></a>}

            <a id="L128"></a>if minline, ok := found[path]; path != &#34;&#34; &amp;&amp; ok {
                <a id="L129"></a>if minline &gt;= line {
                    <a id="L130"></a><span class="comment">// We&#39;ve already covered this file</span>
                    <a id="L131"></a>continue
                <a id="L132"></a>}
            <a id="L133"></a>}
            <a id="L134"></a>found[path] = line;

            <a id="L136"></a>a, err := o.alineFromLine(path, line);
            <a id="L137"></a>if err != nil {
                <a id="L138"></a>t.Errorf(&#34;absolute line %d in object %s maps to %s:%d, but mapping that back gives error %s&#34;, i, o.Paths[0].Name, path, line, err)
            <a id="L139"></a>} else if a != i {
                <a id="L140"></a>t.Errorf(&#34;absolute line %d in object %s maps to %s:%d, which maps back to absolute line %d\n&#34;, i, o.Paths[0].Name, path, line, a)
            <a id="L141"></a>}
        <a id="L142"></a>}
    <a id="L143"></a>}
<a id="L144"></a>}

<a id="L146"></a><span class="comment">// gotest: if [ &#34;$(uname)-$(uname -m)&#34; = Linux-x86_64 ]; then</span>
<a id="L147"></a><span class="comment">// gotest:    mkdir -p _test &amp;&amp; $AS pclinetest.s &amp;&amp; $LD -E main -l -o _test/pclinetest pclinetest.$O</span>
<a id="L148"></a><span class="comment">// gotest: fi</span>
<a id="L149"></a>func TestPCLine(t *testing.T) {
    <a id="L150"></a>if !dotest() {
        <a id="L151"></a>return
    <a id="L152"></a>}

    <a id="L154"></a>f, tab := crack(&#34;_test/pclinetest&#34;, t);
    <a id="L155"></a>text := f.Section(&#34;.text&#34;);
    <a id="L156"></a>textdat, err := text.Data();
    <a id="L157"></a>if err != nil {
        <a id="L158"></a>t.Fatalf(&#34;reading .text: %v&#34;, err)
    <a id="L159"></a>}

    <a id="L161"></a><span class="comment">// Test PCToLine</span>
    <a id="L162"></a>sym := tab.LookupFunc(&#34;linefrompc&#34;);
    <a id="L163"></a>wantLine := 0;
    <a id="L164"></a>for pc := sym.Entry; pc &lt; sym.End; pc++ {
        <a id="L165"></a>file, line, fn := tab.PCToLine(pc);
        <a id="L166"></a>off := pc - text.Addr; <span class="comment">// TODO(rsc): should not need off; bug in 8g</span>
        <a id="L167"></a>wantLine += int(textdat[off]);
        <a id="L168"></a>if fn == nil {
            <a id="L169"></a>t.Errorf(&#34;failed to get line of PC %#x&#34;, pc)
        <a id="L170"></a>} else if len(file) &lt; 12 || file[len(file)-12:len(file)] != &#34;pclinetest.s&#34; || line != wantLine || fn != sym {
            <a id="L171"></a>t.Errorf(&#34;expected %s:%d (%s) at PC %#x, got %s:%d (%s)&#34;, &#34;pclinetest.s&#34;, wantLine, sym.Name, pc, file, line, fn.Name)
        <a id="L172"></a>}
    <a id="L173"></a>}

    <a id="L175"></a><span class="comment">// Test LineToPC</span>
    <a id="L176"></a>sym = tab.LookupFunc(&#34;pcfromline&#34;);
    <a id="L177"></a>lookupline := -1;
    <a id="L178"></a>wantLine = 0;
    <a id="L179"></a>off := uint64(0); <span class="comment">// TODO(rsc): should not need off; bug in 8g</span>
    <a id="L180"></a>for pc := sym.Value; pc &lt; sym.End; pc += 2 + uint64(textdat[off]) {
        <a id="L181"></a>file, line, fn := tab.PCToLine(pc);
        <a id="L182"></a>off = pc - text.Addr;
        <a id="L183"></a>wantLine += int(textdat[off]);
        <a id="L184"></a>if line != wantLine {
            <a id="L185"></a>t.Errorf(&#34;expected line %d at PC %#x in pcfromline, got %d&#34;, wantLine, pc, line);
            <a id="L186"></a>off = pc + 1 - text.Addr;
            <a id="L187"></a>continue;
        <a id="L188"></a>}
        <a id="L189"></a>if lookupline == -1 {
            <a id="L190"></a>lookupline = line
        <a id="L191"></a>}
        <a id="L192"></a>for ; lookupline &lt;= line; lookupline++ {
            <a id="L193"></a>pc2, fn2, err := tab.LineToPC(file, lookupline);
            <a id="L194"></a>if lookupline != line {
                <a id="L195"></a><span class="comment">// Should be nothing on this line</span>
                <a id="L196"></a>if err == nil {
                    <a id="L197"></a>t.Errorf(&#34;expected no PC at line %d, got %#x (%s)&#34;, lookupline, pc2, fn2.Name)
                <a id="L198"></a>}
            <a id="L199"></a>} else if err != nil {
                <a id="L200"></a>t.Errorf(&#34;failed to get PC of line %d: %s&#34;, lookupline, err)
            <a id="L201"></a>} else if pc != pc2 {
                <a id="L202"></a>t.Errorf(&#34;expected PC %#x (%s) at line %d, got PC %#x (%s)&#34;, pc, fn.Name, line, pc2, fn2.Name)
            <a id="L203"></a>}
        <a id="L204"></a>}
        <a id="L205"></a>off = pc + 1 - text.Addr;
    <a id="L206"></a>}
<a id="L207"></a>}
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
