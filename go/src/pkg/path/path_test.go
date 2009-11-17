<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/path/path_test.go</title>

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
	<li>Thu Nov 12 15:58:52 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/path/path_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package path

<a id="L7"></a>import (
    <a id="L8"></a>&#34;os&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a>type CleanTest struct {
    <a id="L13"></a>path, clean string;
<a id="L14"></a>}

<a id="L16"></a>var cleantests = []CleanTest{
    <a id="L17"></a><span class="comment">// Already clean</span>
    <a id="L18"></a>CleanTest{&#34;&#34;, &#34;.&#34;},
    <a id="L19"></a>CleanTest{&#34;abc&#34;, &#34;abc&#34;},
    <a id="L20"></a>CleanTest{&#34;abc/def&#34;, &#34;abc/def&#34;},
    <a id="L21"></a>CleanTest{&#34;a/b/c&#34;, &#34;a/b/c&#34;},
    <a id="L22"></a>CleanTest{&#34;.&#34;, &#34;.&#34;},
    <a id="L23"></a>CleanTest{&#34;..&#34;, &#34;..&#34;},
    <a id="L24"></a>CleanTest{&#34;../..&#34;, &#34;../..&#34;},
    <a id="L25"></a>CleanTest{&#34;../../abc&#34;, &#34;../../abc&#34;},
    <a id="L26"></a>CleanTest{&#34;/abc&#34;, &#34;/abc&#34;},
    <a id="L27"></a>CleanTest{&#34;/&#34;, &#34;/&#34;},

    <a id="L29"></a><span class="comment">// Remove trailing slash</span>
    <a id="L30"></a>CleanTest{&#34;abc/&#34;, &#34;abc&#34;},
    <a id="L31"></a>CleanTest{&#34;abc/def/&#34;, &#34;abc/def&#34;},
    <a id="L32"></a>CleanTest{&#34;a/b/c/&#34;, &#34;a/b/c&#34;},
    <a id="L33"></a>CleanTest{&#34;./&#34;, &#34;.&#34;},
    <a id="L34"></a>CleanTest{&#34;../&#34;, &#34;..&#34;},
    <a id="L35"></a>CleanTest{&#34;../../&#34;, &#34;../..&#34;},
    <a id="L36"></a>CleanTest{&#34;/abc/&#34;, &#34;/abc&#34;},

    <a id="L38"></a><span class="comment">// Remove doubled slash</span>
    <a id="L39"></a>CleanTest{&#34;abc//def//ghi&#34;, &#34;abc/def/ghi&#34;},
    <a id="L40"></a>CleanTest{&#34;//abc&#34;, &#34;/abc&#34;},
    <a id="L41"></a>CleanTest{&#34;///abc&#34;, &#34;/abc&#34;},
    <a id="L42"></a>CleanTest{&#34;//abc//&#34;, &#34;/abc&#34;},
    <a id="L43"></a>CleanTest{&#34;abc//&#34;, &#34;abc&#34;},

    <a id="L45"></a><span class="comment">// Remove . elements</span>
    <a id="L46"></a>CleanTest{&#34;abc/./def&#34;, &#34;abc/def&#34;},
    <a id="L47"></a>CleanTest{&#34;/./abc/def&#34;, &#34;/abc/def&#34;},
    <a id="L48"></a>CleanTest{&#34;abc/.&#34;, &#34;abc&#34;},

    <a id="L50"></a><span class="comment">// Remove .. elements</span>
    <a id="L51"></a>CleanTest{&#34;abc/def/ghi/../jkl&#34;, &#34;abc/def/jkl&#34;},
    <a id="L52"></a>CleanTest{&#34;abc/def/../ghi/../jkl&#34;, &#34;abc/jkl&#34;},
    <a id="L53"></a>CleanTest{&#34;abc/def/..&#34;, &#34;abc&#34;},
    <a id="L54"></a>CleanTest{&#34;abc/def/../..&#34;, &#34;.&#34;},
    <a id="L55"></a>CleanTest{&#34;/abc/def/../..&#34;, &#34;/&#34;},
    <a id="L56"></a>CleanTest{&#34;abc/def/../../..&#34;, &#34;..&#34;},
    <a id="L57"></a>CleanTest{&#34;/abc/def/../../..&#34;, &#34;/&#34;},
    <a id="L58"></a>CleanTest{&#34;abc/def/../../../ghi/jkl/../../../mno&#34;, &#34;../../mno&#34;},

    <a id="L60"></a><span class="comment">// Combinations</span>
    <a id="L61"></a>CleanTest{&#34;abc/./../def&#34;, &#34;def&#34;},
    <a id="L62"></a>CleanTest{&#34;abc//./../def&#34;, &#34;def&#34;},
    <a id="L63"></a>CleanTest{&#34;abc/../../././../def&#34;, &#34;../../def&#34;},
<a id="L64"></a>}

<a id="L66"></a>func TestClean(t *testing.T) {
    <a id="L67"></a>for _, test := range cleantests {
        <a id="L68"></a>if s := Clean(test.path); s != test.clean {
            <a id="L69"></a>t.Errorf(&#34;Clean(%q) = %q, want %q&#34;, test.path, s, test.clean)
        <a id="L70"></a>}
    <a id="L71"></a>}
<a id="L72"></a>}

<a id="L74"></a>type SplitTest struct {
    <a id="L75"></a>path, dir, file string;
<a id="L76"></a>}

<a id="L78"></a>var splittests = []SplitTest{
    <a id="L79"></a>SplitTest{&#34;a/b&#34;, &#34;a/&#34;, &#34;b&#34;},
    <a id="L80"></a>SplitTest{&#34;a/b/&#34;, &#34;a/b/&#34;, &#34;&#34;},
    <a id="L81"></a>SplitTest{&#34;a/&#34;, &#34;a/&#34;, &#34;&#34;},
    <a id="L82"></a>SplitTest{&#34;a&#34;, &#34;&#34;, &#34;a&#34;},
    <a id="L83"></a>SplitTest{&#34;/&#34;, &#34;/&#34;, &#34;&#34;},
<a id="L84"></a>}

<a id="L86"></a>func TestSplit(t *testing.T) {
    <a id="L87"></a>for _, test := range splittests {
        <a id="L88"></a>if d, f := Split(test.path); d != test.dir || f != test.file {
            <a id="L89"></a>t.Errorf(&#34;Split(%q) = %q, %q, want %q, %q&#34;, test.path, d, f, test.dir, test.file)
        <a id="L90"></a>}
    <a id="L91"></a>}
<a id="L92"></a>}

<a id="L94"></a>type JoinTest struct {
    <a id="L95"></a>dir, file, path string;
<a id="L96"></a>}

<a id="L98"></a>var jointests = []JoinTest{
    <a id="L99"></a>JoinTest{&#34;a&#34;, &#34;b&#34;, &#34;a/b&#34;},
    <a id="L100"></a>JoinTest{&#34;a&#34;, &#34;&#34;, &#34;a&#34;},
    <a id="L101"></a>JoinTest{&#34;&#34;, &#34;b&#34;, &#34;b&#34;},
    <a id="L102"></a>JoinTest{&#34;/&#34;, &#34;a&#34;, &#34;/a&#34;},
    <a id="L103"></a>JoinTest{&#34;/&#34;, &#34;&#34;, &#34;/&#34;},
    <a id="L104"></a>JoinTest{&#34;a/&#34;, &#34;b&#34;, &#34;a/b&#34;},
    <a id="L105"></a>JoinTest{&#34;a/&#34;, &#34;&#34;, &#34;a&#34;},
<a id="L106"></a>}

<a id="L108"></a>func TestJoin(t *testing.T) {
    <a id="L109"></a>for _, test := range jointests {
        <a id="L110"></a>if p := Join(test.dir, test.file); p != test.path {
            <a id="L111"></a>t.Errorf(&#34;Join(%q, %q) = %q, want %q&#34;, test.dir, test.file, p, test.path)
        <a id="L112"></a>}
    <a id="L113"></a>}
<a id="L114"></a>}

<a id="L116"></a>type ExtTest struct {
    <a id="L117"></a>path, ext string;
<a id="L118"></a>}

<a id="L120"></a>var exttests = []ExtTest{
    <a id="L121"></a>ExtTest{&#34;path.go&#34;, &#34;.go&#34;},
    <a id="L122"></a>ExtTest{&#34;path.pb.go&#34;, &#34;.go&#34;},
    <a id="L123"></a>ExtTest{&#34;a.dir/b&#34;, &#34;&#34;},
    <a id="L124"></a>ExtTest{&#34;a.dir/b.go&#34;, &#34;.go&#34;},
    <a id="L125"></a>ExtTest{&#34;a.dir/&#34;, &#34;&#34;},
<a id="L126"></a>}

<a id="L128"></a>func TestExt(t *testing.T) {
    <a id="L129"></a>for _, test := range exttests {
        <a id="L130"></a>if x := Ext(test.path); x != test.ext {
            <a id="L131"></a>t.Errorf(&#34;Ext(%q) = %q, want %q&#34;, test.path, x, test.ext)
        <a id="L132"></a>}
    <a id="L133"></a>}
<a id="L134"></a>}

<a id="L136"></a>type Node struct {
    <a id="L137"></a>name    string;
    <a id="L138"></a>entries []*Node; <span class="comment">// nil if the entry is a file</span>
    <a id="L139"></a>mark    int;
<a id="L140"></a>}

<a id="L142"></a>var tree = &amp;Node{
    <a id="L143"></a>&#34;testdata&#34;,
    <a id="L144"></a>[]*Node{
        <a id="L145"></a>&amp;Node{&#34;a&#34;, nil, 0},
        <a id="L146"></a>&amp;Node{&#34;b&#34;, []*Node{}, 0},
        <a id="L147"></a>&amp;Node{&#34;c&#34;, nil, 0},
        <a id="L148"></a>&amp;Node{
            <a id="L149"></a>&#34;d&#34;,
            <a id="L150"></a>[]*Node{
                <a id="L151"></a>&amp;Node{&#34;x&#34;, nil, 0},
                <a id="L152"></a>&amp;Node{&#34;y&#34;, []*Node{}, 0},
                <a id="L153"></a>&amp;Node{
                    <a id="L154"></a>&#34;z&#34;,
                    <a id="L155"></a>[]*Node{
                        <a id="L156"></a>&amp;Node{&#34;u&#34;, nil, 0},
                        <a id="L157"></a>&amp;Node{&#34;v&#34;, nil, 0},
                    <a id="L158"></a>},
                    <a id="L159"></a>0,
                <a id="L160"></a>},
            <a id="L161"></a>},
            <a id="L162"></a>0,
        <a id="L163"></a>},
    <a id="L164"></a>},
    <a id="L165"></a>0,
<a id="L166"></a>}

<a id="L168"></a>func walkTree(n *Node, path string, f func(path string, n *Node)) {
    <a id="L169"></a>f(path, n);
    <a id="L170"></a>for _, e := range n.entries {
        <a id="L171"></a>walkTree(e, Join(path, e.name), f)
    <a id="L172"></a>}
<a id="L173"></a>}

<a id="L175"></a>func makeTree(t *testing.T) {
    <a id="L176"></a>walkTree(tree, tree.name, func(path string, n *Node) {
        <a id="L177"></a>if n.entries == nil {
            <a id="L178"></a>fd, err := os.Open(path, os.O_CREAT, 0660);
            <a id="L179"></a>if err != nil {
                <a id="L180"></a>t.Errorf(&#34;makeTree: %v&#34;, err)
            <a id="L181"></a>}
            <a id="L182"></a>fd.Close();
        <a id="L183"></a>} else {
            <a id="L184"></a>os.Mkdir(path, 0770)
        <a id="L185"></a>}
    <a id="L186"></a>})
<a id="L187"></a>}

<a id="L189"></a>func markTree(n *Node) { walkTree(n, &#34;&#34;, func(path string, n *Node) { n.mark++ }) }

<a id="L191"></a>func checkMarks(t *testing.T) {
    <a id="L192"></a>walkTree(tree, tree.name, func(path string, n *Node) {
        <a id="L193"></a>if n.mark != 1 {
            <a id="L194"></a>t.Errorf(&#34;node %s mark = %d; expected 1&#34;, path, n.mark)
        <a id="L195"></a>}
        <a id="L196"></a>n.mark = 0;
    <a id="L197"></a>})
<a id="L198"></a>}

<a id="L200"></a><span class="comment">// Assumes that each node name is unique. Good enough for a test.</span>
<a id="L201"></a>func mark(name string) {
    <a id="L202"></a>walkTree(tree, tree.name, func(path string, n *Node) {
        <a id="L203"></a>if n.name == name {
            <a id="L204"></a>n.mark++
        <a id="L205"></a>}
    <a id="L206"></a>})
<a id="L207"></a>}

<a id="L209"></a>type TestVisitor struct{}

<a id="L211"></a>func (v *TestVisitor) VisitDir(path string, d *os.Dir) bool {
    <a id="L212"></a>mark(d.Name);
    <a id="L213"></a>return true;
<a id="L214"></a>}

<a id="L216"></a>func (v *TestVisitor) VisitFile(path string, d *os.Dir) {
    <a id="L217"></a>mark(d.Name)
<a id="L218"></a>}

<a id="L220"></a>func TestWalk(t *testing.T) {
    <a id="L221"></a>makeTree(t);

    <a id="L223"></a><span class="comment">// 1) ignore error handling, expect none</span>
    <a id="L224"></a>v := &amp;TestVisitor{};
    <a id="L225"></a>Walk(tree.name, v, nil);
    <a id="L226"></a>checkMarks(t);

    <a id="L228"></a><span class="comment">// 2) handle errors, expect none</span>
    <a id="L229"></a>errors := make(chan os.Error, 64);
    <a id="L230"></a>Walk(tree.name, v, errors);
    <a id="L231"></a>if err, ok := &lt;-errors; ok {
        <a id="L232"></a>t.Errorf(&#34;no error expected, found: s&#34;, err)
    <a id="L233"></a>}
    <a id="L234"></a>checkMarks(t);

    <a id="L236"></a><span class="comment">// introduce 2 errors: chmod top-level directories to 0</span>
    <a id="L237"></a>os.Chmod(Join(tree.name, tree.entries[1].name), 0);
    <a id="L238"></a>os.Chmod(Join(tree.name, tree.entries[3].name), 0);
    <a id="L239"></a><span class="comment">// mark respective subtrees manually</span>
    <a id="L240"></a>markTree(tree.entries[1]);
    <a id="L241"></a>markTree(tree.entries[3]);
    <a id="L242"></a><span class="comment">// correct double-marking of directory itself</span>
    <a id="L243"></a>tree.entries[1].mark--;
    <a id="L244"></a>tree.entries[3].mark--;

    <a id="L246"></a><span class="comment">// 3) handle errors, expect two</span>
    <a id="L247"></a>errors = make(chan os.Error, 64);
    <a id="L248"></a>os.Chmod(Join(tree.name, tree.entries[1].name), 0);
    <a id="L249"></a>Walk(tree.name, v, errors);
    <a id="L250"></a>for i := 1; i &lt;= 2; i++ {
        <a id="L251"></a>if _, ok := &lt;-errors; !ok {
            <a id="L252"></a>t.Errorf(&#34;%d. error expected, none found&#34;, i);
            <a id="L253"></a>break;
        <a id="L254"></a>}
    <a id="L255"></a>}
    <a id="L256"></a>if err, ok := &lt;-errors; ok {
        <a id="L257"></a>t.Errorf(&#34;only two errors expected, found 3rd: %v&#34;, err)
    <a id="L258"></a>}
    <a id="L259"></a><span class="comment">// the inaccessible subtrees were marked manually</span>
    <a id="L260"></a>checkMarks(t);

    <a id="L262"></a><span class="comment">// cleanup</span>
    <a id="L263"></a>os.Chmod(Join(tree.name, tree.entries[1].name), 0770);
    <a id="L264"></a>os.Chmod(Join(tree.name, tree.entries[3].name), 0770);
    <a id="L265"></a>if err := os.RemoveAll(tree.name); err != nil {
        <a id="L266"></a>t.Errorf(&#34;removeTree: %v&#34;, err)
    <a id="L267"></a>}
<a id="L268"></a>}
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
