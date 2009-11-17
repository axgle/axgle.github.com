<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/godoc/godoc.go</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/godoc/godoc.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package main

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;flag&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;go/ast&#34;;
    <a id="L12"></a>&#34;go/doc&#34;;
    <a id="L13"></a>&#34;go/parser&#34;;
    <a id="L14"></a>&#34;go/printer&#34;;
    <a id="L15"></a>&#34;go/scanner&#34;;
    <a id="L16"></a>&#34;go/token&#34;;
    <a id="L17"></a>&#34;http&#34;;
    <a id="L18"></a>&#34;io&#34;;
    <a id="L19"></a>&#34;log&#34;;
    <a id="L20"></a>&#34;os&#34;;
    <a id="L21"></a>pathutil &#34;path&#34;;
    <a id="L22"></a>&#34;strings&#34;;
    <a id="L23"></a>&#34;sync&#34;;
    <a id="L24"></a>&#34;template&#34;;
    <a id="L25"></a>&#34;time&#34;;
    <a id="L26"></a>&#34;unicode&#34;;
    <a id="L27"></a>&#34;utf8&#34;;
<a id="L28"></a>)


<a id="L31"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L32"></a><span class="comment">// Support types</span>

<a id="L34"></a><span class="comment">// An RWValue wraps a value and permits mutually exclusive</span>
<a id="L35"></a><span class="comment">// access to it and records the time the value was last set.</span>
<a id="L36"></a>type RWValue struct {
    <a id="L37"></a>mutex     sync.RWMutex;
    <a id="L38"></a>value     interface{};
    <a id="L39"></a>timestamp int64; <span class="comment">// time of last set(), in seconds since epoch</span>
<a id="L40"></a>}


<a id="L43"></a>func (v *RWValue) set(value interface{}) {
    <a id="L44"></a>v.mutex.Lock();
    <a id="L45"></a>v.value = value;
    <a id="L46"></a>v.timestamp = time.Seconds();
    <a id="L47"></a>v.mutex.Unlock();
<a id="L48"></a>}


<a id="L51"></a>func (v *RWValue) get() (interface{}, int64) {
    <a id="L52"></a>v.mutex.RLock();
    <a id="L53"></a>defer v.mutex.RUnlock();
    <a id="L54"></a>return v.value, v.timestamp;
<a id="L55"></a>}


<a id="L58"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L59"></a><span class="comment">// Globals</span>

<a id="L61"></a>type delayTime struct {
    <a id="L62"></a>RWValue;
<a id="L63"></a>}


<a id="L66"></a>func (dt *delayTime) backoff(max int) {
    <a id="L67"></a>dt.mutex.Lock();
    <a id="L68"></a>v := dt.value.(int) * 2;
    <a id="L69"></a>if v &gt; max {
        <a id="L70"></a>v = max
    <a id="L71"></a>}
    <a id="L72"></a>dt.value = v;
    <a id="L73"></a>dt.mutex.Unlock();
<a id="L74"></a>}


<a id="L77"></a>var (
    <a id="L78"></a>verbose = flag.Bool(&#34;v&#34;, false, &#34;verbose mode&#34;);

    <a id="L80"></a><span class="comment">// file system roots</span>
    <a id="L81"></a>goroot   string;
    <a id="L82"></a>cmdroot  = flag.String(&#34;cmdroot&#34;, &#34;src/cmd&#34;, &#34;root command source directory (if unrooted, relative to goroot)&#34;);
    <a id="L83"></a>pkgroot  = flag.String(&#34;pkgroot&#34;, &#34;src/pkg&#34;, &#34;root package source directory (if unrooted, relative to goroot)&#34;);
    <a id="L84"></a>tmplroot = flag.String(&#34;tmplroot&#34;, &#34;lib/godoc&#34;, &#34;root template directory (if unrooted, relative to goroot)&#34;);

    <a id="L86"></a><span class="comment">// layout control</span>
    <a id="L87"></a>tabwidth = flag.Int(&#34;tabwidth&#34;, 4, &#34;tab width&#34;);
<a id="L88"></a>)


<a id="L91"></a>var fsTree RWValue <span class="comment">// *Directory tree of packages, updated with each sync</span>


<a id="L94"></a>func init() {
    <a id="L95"></a>goroot = os.Getenv(&#34;GOROOT&#34;);
    <a id="L96"></a>if goroot == &#34;&#34; {
        <a id="L97"></a>goroot = pathutil.Join(os.Getenv(&#34;HOME&#34;), &#34;go&#34;)
    <a id="L98"></a>}
    <a id="L99"></a>flag.StringVar(&amp;goroot, &#34;goroot&#34;, goroot, &#34;Go root directory&#34;);
<a id="L100"></a>}


<a id="L103"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L104"></a><span class="comment">// Predicates and small utility functions</span>

<a id="L106"></a>func isGoFile(dir *os.Dir) bool {
    <a id="L107"></a>return dir.IsRegular() &amp;&amp;
        <a id="L108"></a>!strings.HasPrefix(dir.Name, &#34;.&#34;) &amp;&amp; <span class="comment">// ignore .files</span>
        <a id="L109"></a>pathutil.Ext(dir.Name) == &#34;.go&#34;
<a id="L110"></a>}


<a id="L113"></a>func isPkgFile(dir *os.Dir) bool {
    <a id="L114"></a>return isGoFile(dir) &amp;&amp;
        <a id="L115"></a>!strings.HasSuffix(dir.Name, &#34;_test.go&#34;) <span class="comment">// ignore test files</span>
<a id="L116"></a>}


<a id="L119"></a>func isPkgDir(dir *os.Dir) bool {
    <a id="L120"></a>return dir.IsDirectory() &amp;&amp; len(dir.Name) &gt; 0 &amp;&amp; dir.Name[0] != &#39;_&#39;
<a id="L121"></a>}


<a id="L124"></a>func pkgName(filename string) string {
    <a id="L125"></a>file, err := parse(filename, parser.PackageClauseOnly);
    <a id="L126"></a>if err != nil || file == nil {
        <a id="L127"></a>return &#34;&#34;
    <a id="L128"></a>}
    <a id="L129"></a>return file.Name.Value;
<a id="L130"></a>}


<a id="L133"></a>func htmlEscape(s string) string {
    <a id="L134"></a>var buf bytes.Buffer;
    <a id="L135"></a>template.HTMLEscape(&amp;buf, strings.Bytes(s));
    <a id="L136"></a>return buf.String();
<a id="L137"></a>}


<a id="L140"></a>func firstSentence(s string) string {
    <a id="L141"></a>i := -1; <span class="comment">// index+1 of first period</span>
    <a id="L142"></a>j := -1; <span class="comment">// index+1 of first period that is followed by white space</span>
    <a id="L143"></a>prev := &#39;A&#39;;
    <a id="L144"></a>for k, ch := range s {
        <a id="L145"></a>k1 := k + 1;
        <a id="L146"></a>if ch == &#39;.&#39; {
            <a id="L147"></a>if i &lt; 0 {
                <a id="L148"></a>i = k1 <span class="comment">// first period</span>
            <a id="L149"></a>}
            <a id="L150"></a>if k1 &lt; len(s) &amp;&amp; s[k1] &lt;= &#39; &#39; {
                <a id="L151"></a>if j &lt; 0 {
                    <a id="L152"></a>j = k1 <span class="comment">// first period followed by white space</span>
                <a id="L153"></a>}
                <a id="L154"></a>if !unicode.IsUpper(prev) {
                    <a id="L155"></a>j = k1;
                    <a id="L156"></a>break;
                <a id="L157"></a>}
            <a id="L158"></a>}
        <a id="L159"></a>}
        <a id="L160"></a>prev = ch;
    <a id="L161"></a>}

    <a id="L163"></a>if j &lt; 0 {
        <a id="L164"></a><span class="comment">// use the next best period</span>
        <a id="L165"></a>j = i;
        <a id="L166"></a>if j &lt; 0 {
            <a id="L167"></a><span class="comment">// no period at all, use the entire string</span>
            <a id="L168"></a>j = len(s)
        <a id="L169"></a>}
    <a id="L170"></a>}

    <a id="L172"></a>return s[0:j];
<a id="L173"></a>}


<a id="L176"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L177"></a><span class="comment">// Package directories</span>

<a id="L179"></a>type Directory struct {
    <a id="L180"></a>Depth int;
    <a id="L181"></a>Path  string; <span class="comment">// includes Name</span>
    <a id="L182"></a>Name  string;
    <a id="L183"></a>Text  string;       <span class="comment">// package documentation, if any</span>
    <a id="L184"></a>Dirs  []*Directory; <span class="comment">// subdirectories</span>
<a id="L185"></a>}


<a id="L188"></a>func newDirTree(path, name string, depth, maxDepth int) *Directory {
    <a id="L189"></a>if depth &gt;= maxDepth {
        <a id="L190"></a><span class="comment">// return a dummy directory so that the parent directory</span>
        <a id="L191"></a><span class="comment">// doesn&#39;t get discarded just because we reached the max</span>
        <a id="L192"></a><span class="comment">// directory depth</span>
        <a id="L193"></a>return &amp;Directory{depth, path, name, &#34;&#34;, nil}
    <a id="L194"></a>}

    <a id="L196"></a>list, _ := io.ReadDir(path); <span class="comment">// ignore errors</span>

    <a id="L198"></a><span class="comment">// determine number of subdirectories and package files</span>
    <a id="L199"></a>ndirs := 0;
    <a id="L200"></a>nfiles := 0;
    <a id="L201"></a>text := &#34;&#34;;
    <a id="L202"></a>for _, d := range list {
        <a id="L203"></a>switch {
        <a id="L204"></a>case isPkgDir(d):
            <a id="L205"></a>ndirs++
        <a id="L206"></a>case isPkgFile(d):
            <a id="L207"></a>nfiles++;
            <a id="L208"></a>if text == &#34;&#34; {
                <a id="L209"></a><span class="comment">// no package documentation yet; take the first found</span>
                <a id="L210"></a>file, err := parser.ParseFile(pathutil.Join(path, d.Name), nil,
                    <a id="L211"></a>parser.ParseComments|parser.PackageClauseOnly);
                <a id="L212"></a>if err == nil &amp;&amp;
                    <a id="L213"></a><span class="comment">// Also accept fakePkgName, so we get synopses for commmands.</span>
                    <a id="L214"></a><span class="comment">// Note: This may lead to incorrect results if there is a</span>
                    <a id="L215"></a><span class="comment">// (left-over) &#34;documentation&#34; package somewhere in a package</span>
                    <a id="L216"></a><span class="comment">// directory of different name, but this is very unlikely and</span>
                    <a id="L217"></a><span class="comment">// against current conventions.</span>
                    <a id="L218"></a>(file.Name.Value == name || file.Name.Value == fakePkgName) &amp;&amp;
                    <a id="L219"></a>file.Doc != nil {
                    <a id="L220"></a><span class="comment">// found documentation; extract a synopsys</span>
                    <a id="L221"></a>text = firstSentence(doc.CommentText(file.Doc))
                <a id="L222"></a>}
            <a id="L223"></a>}
        <a id="L224"></a>}
    <a id="L225"></a>}

    <a id="L227"></a><span class="comment">// create subdirectory tree</span>
    <a id="L228"></a>var dirs []*Directory;
    <a id="L229"></a>if ndirs &gt; 0 {
        <a id="L230"></a>dirs = make([]*Directory, ndirs);
        <a id="L231"></a>i := 0;
        <a id="L232"></a>for _, d := range list {
            <a id="L233"></a>if isPkgDir(d) {
                <a id="L234"></a>dd := newDirTree(pathutil.Join(path, d.Name), d.Name, depth+1, maxDepth);
                <a id="L235"></a>if dd != nil {
                    <a id="L236"></a>dirs[i] = dd;
                    <a id="L237"></a>i++;
                <a id="L238"></a>}
            <a id="L239"></a>}
        <a id="L240"></a>}
        <a id="L241"></a>dirs = dirs[0:i];
    <a id="L242"></a>}

    <a id="L244"></a><span class="comment">// if there are no package files and no subdirectories</span>
    <a id="L245"></a><span class="comment">// (with package files), ignore the directory</span>
    <a id="L246"></a>if nfiles == 0 &amp;&amp; len(dirs) == 0 {
        <a id="L247"></a>return nil
    <a id="L248"></a>}

    <a id="L250"></a>return &amp;Directory{depth, path, name, text, dirs};
<a id="L251"></a>}


<a id="L254"></a><span class="comment">// newDirectory creates a new package directory tree with at most maxDepth</span>
<a id="L255"></a><span class="comment">// levels, anchored at root which is relative to goroot. The result tree</span>
<a id="L256"></a><span class="comment">// only contains directories that contain package files or that contain</span>
<a id="L257"></a><span class="comment">// subdirectories containing package files (transitively).</span>
<a id="L258"></a><span class="comment">//</span>
<a id="L259"></a>func newDirectory(root string, maxDepth int) *Directory {
    <a id="L260"></a>d, err := os.Lstat(root);
    <a id="L261"></a>if err != nil || !isPkgDir(d) {
        <a id="L262"></a>return nil
    <a id="L263"></a>}
    <a id="L264"></a>return newDirTree(root, d.Name, 0, maxDepth);
<a id="L265"></a>}


<a id="L268"></a>func (dir *Directory) walk(c chan&lt;- *Directory, skipRoot bool) {
    <a id="L269"></a>if dir != nil {
        <a id="L270"></a>if !skipRoot {
            <a id="L271"></a>c &lt;- dir
        <a id="L272"></a>}
        <a id="L273"></a>for _, d := range dir.Dirs {
            <a id="L274"></a>d.walk(c, false)
        <a id="L275"></a>}
    <a id="L276"></a>}
<a id="L277"></a>}


<a id="L280"></a>func (dir *Directory) iter(skipRoot bool) &lt;-chan *Directory {
    <a id="L281"></a>c := make(chan *Directory);
    <a id="L282"></a>go func() {
        <a id="L283"></a>dir.walk(c, skipRoot);
        <a id="L284"></a>close(c);
    <a id="L285"></a>}();
    <a id="L286"></a>return c;
<a id="L287"></a>}


<a id="L290"></a><span class="comment">// lookup looks for the *Directory for a given path, relative to dir.</span>
<a id="L291"></a>func (dir *Directory) lookup(path string) *Directory {
    <a id="L292"></a>path = pathutil.Clean(path); <span class="comment">// no trailing &#39;/&#39;</span>

    <a id="L294"></a>if dir == nil || path == &#34;&#34; || path == &#34;.&#34; {
        <a id="L295"></a>return dir
    <a id="L296"></a>}

    <a id="L298"></a>dpath, dname := pathutil.Split(path);
    <a id="L299"></a>if dpath == &#34;&#34; {
        <a id="L300"></a><span class="comment">// directory-local name</span>
        <a id="L301"></a>for _, d := range dir.Dirs {
            <a id="L302"></a>if dname == d.Name {
                <a id="L303"></a>return d
            <a id="L304"></a>}
        <a id="L305"></a>}
        <a id="L306"></a>return nil;
    <a id="L307"></a>}

    <a id="L309"></a>return dir.lookup(dpath).lookup(dname);
<a id="L310"></a>}


<a id="L313"></a><span class="comment">// DirEntry describes a directory entry. The Depth and Height values</span>
<a id="L314"></a><span class="comment">// are useful for presenting an entry in an indented fashion.</span>
<a id="L315"></a><span class="comment">//</span>
<a id="L316"></a>type DirEntry struct {
    <a id="L317"></a>Depth    int;    <span class="comment">// &gt;= 0</span>
    <a id="L318"></a>Height   int;    <span class="comment">// = DirList.MaxHeight - Depth, &gt; 0</span>
    <a id="L319"></a>Path     string; <span class="comment">// includes Name, relative to DirList root</span>
    <a id="L320"></a>Name     string;
    <a id="L321"></a>Synopsis string;
<a id="L322"></a>}


<a id="L325"></a>type DirList struct {
    <a id="L326"></a>MaxHeight int; <span class="comment">// directory tree height, &gt; 0</span>
    <a id="L327"></a>List      []DirEntry;
<a id="L328"></a>}


<a id="L331"></a><span class="comment">// listing creates a (linear) directory listing from a directory tree.</span>
<a id="L332"></a><span class="comment">// If skipRoot is set, the root directory itself is excluded from the list.</span>
<a id="L333"></a><span class="comment">//</span>
<a id="L334"></a>func (root *Directory) listing(skipRoot bool) *DirList {
    <a id="L335"></a>if root == nil {
        <a id="L336"></a>return nil
    <a id="L337"></a>}

    <a id="L339"></a><span class="comment">// determine number of entries n and maximum height</span>
    <a id="L340"></a>n := 0;
    <a id="L341"></a>minDepth := 1 &lt;&lt; 30; <span class="comment">// infinity</span>
    <a id="L342"></a>maxDepth := 0;
    <a id="L343"></a>for d := range root.iter(skipRoot) {
        <a id="L344"></a>n++;
        <a id="L345"></a>if minDepth &gt; d.Depth {
            <a id="L346"></a>minDepth = d.Depth
        <a id="L347"></a>}
        <a id="L348"></a>if maxDepth &lt; d.Depth {
            <a id="L349"></a>maxDepth = d.Depth
        <a id="L350"></a>}
    <a id="L351"></a>}
    <a id="L352"></a>maxHeight := maxDepth - minDepth + 1;

    <a id="L354"></a>if n == 0 {
        <a id="L355"></a>return nil
    <a id="L356"></a>}

    <a id="L358"></a><span class="comment">// create list</span>
    <a id="L359"></a>list := make([]DirEntry, n);
    <a id="L360"></a>i := 0;
    <a id="L361"></a>for d := range root.iter(skipRoot) {
        <a id="L362"></a>p := &amp;list[i];
        <a id="L363"></a>p.Depth = d.Depth - minDepth;
        <a id="L364"></a>p.Height = maxHeight - p.Depth;
        <a id="L365"></a><span class="comment">// the path is relative to root.Path - remove the root.Path</span>
        <a id="L366"></a><span class="comment">// prefix (the prefix should always be present but avoid</span>
        <a id="L367"></a><span class="comment">// crashes and check)</span>
        <a id="L368"></a>path := d.Path;
        <a id="L369"></a>if strings.HasPrefix(d.Path, root.Path) {
            <a id="L370"></a>path = d.Path[len(root.Path):len(d.Path)]
        <a id="L371"></a>}
        <a id="L372"></a><span class="comment">// remove trailing &#39;/&#39; if any - path must be relative</span>
        <a id="L373"></a>if len(path) &gt; 0 &amp;&amp; path[0] == &#39;/&#39; {
            <a id="L374"></a>path = path[1:len(path)]
        <a id="L375"></a>}
        <a id="L376"></a>p.Path = path;
        <a id="L377"></a>p.Name = d.Name;
        <a id="L378"></a>p.Synopsis = d.Text;
        <a id="L379"></a>i++;
    <a id="L380"></a>}

    <a id="L382"></a>return &amp;DirList{maxHeight, list};
<a id="L383"></a>}


<a id="L386"></a>func listing(dirs []*os.Dir) *DirList {
    <a id="L387"></a>list := make([]DirEntry, len(dirs)+1);
    <a id="L388"></a>list[0] = DirEntry{0, 1, &#34;..&#34;, &#34;..&#34;, &#34;&#34;};
    <a id="L389"></a>for i, d := range dirs {
        <a id="L390"></a>p := &amp;list[i+1];
        <a id="L391"></a>p.Depth = 0;
        <a id="L392"></a>p.Height = 1;
        <a id="L393"></a>p.Path = d.Name;
        <a id="L394"></a>p.Name = d.Name;
    <a id="L395"></a>}
    <a id="L396"></a>return &amp;DirList{1, list};
<a id="L397"></a>}


<a id="L400"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L401"></a><span class="comment">// Parsing</span>

<a id="L403"></a><span class="comment">// A single error in the parsed file.</span>
<a id="L404"></a>type parseError struct {
    <a id="L405"></a>src  []byte; <span class="comment">// source before error</span>
    <a id="L406"></a>line int;    <span class="comment">// line number of error</span>
    <a id="L407"></a>msg  string; <span class="comment">// error message</span>
<a id="L408"></a>}


<a id="L411"></a><span class="comment">// All the errors in the parsed file, plus surrounding source code.</span>
<a id="L412"></a><span class="comment">// Each error has a slice giving the source text preceding it</span>
<a id="L413"></a><span class="comment">// (starting where the last error occurred).  The final element in list[]</span>
<a id="L414"></a><span class="comment">// has msg = &#34;&#34;, to give the remainder of the source code.</span>
<a id="L415"></a><span class="comment">// This data structure is handed to the templates parseerror.txt and parseerror.html.</span>
<a id="L416"></a><span class="comment">//</span>
<a id="L417"></a>type parseErrors struct {
    <a id="L418"></a>filename string;       <span class="comment">// path to file</span>
    <a id="L419"></a>list     []parseError; <span class="comment">// the errors</span>
    <a id="L420"></a>src      []byte;       <span class="comment">// the file&#39;s entire source code</span>
<a id="L421"></a>}


<a id="L424"></a><span class="comment">// Parses a file (path) and returns the corresponding AST and</span>
<a id="L425"></a><span class="comment">// a sorted list (by file position) of errors, if any.</span>
<a id="L426"></a><span class="comment">//</span>
<a id="L427"></a>func parse(path string, mode uint) (*ast.File, *parseErrors) {
    <a id="L428"></a>src, err := io.ReadFile(path);
    <a id="L429"></a>if err != nil {
        <a id="L430"></a>log.Stderrf(&#34;%v&#34;, err);
        <a id="L431"></a>errs := []parseError{parseError{nil, 0, err.String()}};
        <a id="L432"></a>return nil, &amp;parseErrors{path, errs, nil};
    <a id="L433"></a>}

    <a id="L435"></a>prog, err := parser.ParseFile(path, src, mode);
    <a id="L436"></a>if err != nil {
        <a id="L437"></a>var errs []parseError;
        <a id="L438"></a>if errors, ok := err.(scanner.ErrorList); ok {
            <a id="L439"></a><span class="comment">// convert error list (already sorted)</span>
            <a id="L440"></a><span class="comment">// TODO(gri) If the file contains //line comments, the errors</span>
            <a id="L441"></a><span class="comment">//           may not be sorted in increasing file offset value</span>
            <a id="L442"></a><span class="comment">//           which will lead to incorrect output.</span>
            <a id="L443"></a>errs = make([]parseError, len(errors)+1); <span class="comment">// +1 for final fragment of source</span>
            <a id="L444"></a>offs := 0;
            <a id="L445"></a>for i, r := range errors {
                <a id="L446"></a><span class="comment">// Should always be true, but check for robustness.</span>
                <a id="L447"></a>if 0 &lt;= r.Pos.Offset &amp;&amp; r.Pos.Offset &lt;= len(src) {
                    <a id="L448"></a>errs[i].src = src[offs:r.Pos.Offset];
                    <a id="L449"></a>offs = r.Pos.Offset;
                <a id="L450"></a>}
                <a id="L451"></a>errs[i].line = r.Pos.Line;
                <a id="L452"></a>errs[i].msg = r.Msg;
            <a id="L453"></a>}
            <a id="L454"></a>errs[len(errors)].src = src[offs:len(src)];
        <a id="L455"></a>} else {
            <a id="L456"></a><span class="comment">// single error of unspecified type</span>
            <a id="L457"></a>errs = make([]parseError, 2);
            <a id="L458"></a>errs[0] = parseError{[]byte{}, 0, err.String()};
            <a id="L459"></a>errs[1].src = src;
        <a id="L460"></a>}
        <a id="L461"></a>return nil, &amp;parseErrors{path, errs, src};
    <a id="L462"></a>}

    <a id="L464"></a>return prog, nil;
<a id="L465"></a>}


<a id="L468"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L469"></a><span class="comment">// HTML formatting support</span>

<a id="L471"></a><span class="comment">// Styler implements a printer.Styler.</span>
<a id="L472"></a>type Styler struct {
    <a id="L473"></a>highlight string;
<a id="L474"></a>}


<a id="L477"></a><span class="comment">// Use the defaultStyler when there is no specific styler.</span>
<a id="L478"></a>var defaultStyler Styler


<a id="L481"></a>func (s *Styler) LineTag(line int) (text []byte, tag printer.HTMLTag) {
    <a id="L482"></a>tag = printer.HTMLTag{fmt.Sprintf(`&lt;a id=&#34;L%d&#34;&gt;`, line), &#34;&lt;/a&gt;&#34;};
    <a id="L483"></a>return;
<a id="L484"></a>}


<a id="L487"></a>func (s *Styler) Comment(c *ast.Comment, line []byte) (text []byte, tag printer.HTMLTag) {
    <a id="L488"></a>text = line;
    <a id="L489"></a><span class="comment">// minimal syntax-coloring of comments for now - people will want more</span>
    <a id="L490"></a><span class="comment">// (don&#39;t do anything more until there&#39;s a button to turn it on/off)</span>
    <a id="L491"></a>tag = printer.HTMLTag{`&lt;span class=&#34;comment&#34;&gt;`, &#34;&lt;/span&gt;&#34;};
    <a id="L492"></a>return;
<a id="L493"></a>}


<a id="L496"></a>func (s *Styler) BasicLit(x *ast.BasicLit) (text []byte, tag printer.HTMLTag) {
    <a id="L497"></a>text = x.Value;
    <a id="L498"></a>return;
<a id="L499"></a>}


<a id="L502"></a>func (s *Styler) Ident(id *ast.Ident) (text []byte, tag printer.HTMLTag) {
    <a id="L503"></a>text = strings.Bytes(id.Value);
    <a id="L504"></a>if s.highlight == id.Value {
        <a id="L505"></a>tag = printer.HTMLTag{&#34;&lt;span class=highlight&gt;&#34;, &#34;&lt;/span&gt;&#34;}
    <a id="L506"></a>}
    <a id="L507"></a>return;
<a id="L508"></a>}


<a id="L511"></a>func (s *Styler) Token(tok token.Token) (text []byte, tag printer.HTMLTag) {
    <a id="L512"></a>text = strings.Bytes(tok.String());
    <a id="L513"></a>return;
<a id="L514"></a>}


<a id="L517"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L518"></a><span class="comment">// Templates</span>

<a id="L520"></a><span class="comment">// Write an AST-node to w; optionally html-escaped.</span>
<a id="L521"></a>func writeNode(w io.Writer, node interface{}, html bool, styler printer.Styler) {
    <a id="L522"></a>mode := printer.UseSpaces;
    <a id="L523"></a>if html {
        <a id="L524"></a>mode |= printer.GenHTML
    <a id="L525"></a>}
    <a id="L526"></a>(&amp;printer.Config{mode, *tabwidth, styler}).Fprint(w, node);
<a id="L527"></a>}


<a id="L530"></a><span class="comment">// Write text to w; optionally html-escaped.</span>
<a id="L531"></a>func writeText(w io.Writer, text []byte, html bool) {
    <a id="L532"></a>if html {
        <a id="L533"></a>template.HTMLEscape(w, text);
        <a id="L534"></a>return;
    <a id="L535"></a>}
    <a id="L536"></a>w.Write(text);
<a id="L537"></a>}


<a id="L540"></a><span class="comment">// Write anything to w; optionally html-escaped.</span>
<a id="L541"></a>func writeAny(w io.Writer, x interface{}, html bool) {
    <a id="L542"></a>switch v := x.(type) {
    <a id="L543"></a>case []byte:
        <a id="L544"></a>writeText(w, v, html)
    <a id="L545"></a>case string:
        <a id="L546"></a>writeText(w, strings.Bytes(v), html)
    <a id="L547"></a>case ast.Decl:
        <a id="L548"></a>writeNode(w, v, html, &amp;defaultStyler)
    <a id="L549"></a>case ast.Expr:
        <a id="L550"></a>writeNode(w, v, html, &amp;defaultStyler)
    <a id="L551"></a>default:
        <a id="L552"></a>if html {
            <a id="L553"></a>var buf bytes.Buffer;
            <a id="L554"></a>fmt.Fprint(&amp;buf, x);
            <a id="L555"></a>writeText(w, buf.Bytes(), true);
        <a id="L556"></a>} else {
            <a id="L557"></a>fmt.Fprint(w, x)
        <a id="L558"></a>}
    <a id="L559"></a>}
<a id="L560"></a>}


<a id="L563"></a><span class="comment">// Template formatter for &#34;html&#34; format.</span>
<a id="L564"></a>func htmlFmt(w io.Writer, x interface{}, format string) {
    <a id="L565"></a>writeAny(w, x, true)
<a id="L566"></a>}


<a id="L569"></a><span class="comment">// Template formatter for &#34;html-comment&#34; format.</span>
<a id="L570"></a>func htmlCommentFmt(w io.Writer, x interface{}, format string) {
    <a id="L571"></a>var buf bytes.Buffer;
    <a id="L572"></a>writeAny(&amp;buf, x, false);
    <a id="L573"></a>doc.ToHTML(w, buf.Bytes()); <span class="comment">// does html-escaping</span>
<a id="L574"></a>}


<a id="L577"></a><span class="comment">// Template formatter for &#34;&#34; (default) format.</span>
<a id="L578"></a>func textFmt(w io.Writer, x interface{}, format string) {
    <a id="L579"></a>writeAny(w, x, false)
<a id="L580"></a>}


<a id="L583"></a>func removePrefix(s, prefix string) string {
    <a id="L584"></a>if strings.HasPrefix(s, prefix) {
        <a id="L585"></a>return s[len(prefix):len(s)]
    <a id="L586"></a>}
    <a id="L587"></a>return s;
<a id="L588"></a>}


<a id="L591"></a><span class="comment">// Template formatter for &#34;path&#34; format.</span>
<a id="L592"></a>func pathFmt(w io.Writer, x interface{}, format string) {
    <a id="L593"></a><span class="comment">// TODO(gri): Need to find a better solution for this.</span>
    <a id="L594"></a><span class="comment">//            This will not work correctly if *cmdroot</span>
    <a id="L595"></a><span class="comment">//            or *pkgroot change.</span>
    <a id="L596"></a>writeAny(w, removePrefix(x.(string), &#34;src&#34;), true)
<a id="L597"></a>}


<a id="L600"></a><span class="comment">// Template formatter for &#34;link&#34; format.</span>
<a id="L601"></a>func linkFmt(w io.Writer, x interface{}, format string) {
    <a id="L602"></a>type Positioner interface {
        <a id="L603"></a>Pos() token.Position;
    <a id="L604"></a>}
    <a id="L605"></a>if node, ok := x.(Positioner); ok {
        <a id="L606"></a>pos := node.Pos();
        <a id="L607"></a>if pos.IsValid() {
            <a id="L608"></a><span class="comment">// line id&#39;s in html-printed source are of the</span>
            <a id="L609"></a><span class="comment">// form &#34;L%d&#34; where %d stands for the line number</span>
            <a id="L610"></a>fmt.Fprintf(w, &#34;/%s#L%d&#34;, htmlEscape(pos.Filename), pos.Line)
        <a id="L611"></a>}
    <a id="L612"></a>}
<a id="L613"></a>}


<a id="L616"></a><span class="comment">// The strings in infoKinds must be properly html-escaped.</span>
<a id="L617"></a>var infoKinds = [nKinds]string{
    <a id="L618"></a>PackageClause: &#34;package&amp;nbsp;clause&#34;,
    <a id="L619"></a>ImportDecl: &#34;import&amp;nbsp;decl&#34;,
    <a id="L620"></a>ConstDecl: &#34;const&amp;nbsp;decl&#34;,
    <a id="L621"></a>TypeDecl: &#34;type&amp;nbsp;decl&#34;,
    <a id="L622"></a>VarDecl: &#34;var&amp;nbsp;decl&#34;,
    <a id="L623"></a>FuncDecl: &#34;func&amp;nbsp;decl&#34;,
    <a id="L624"></a>MethodDecl: &#34;method&amp;nbsp;decl&#34;,
    <a id="L625"></a>Use: &#34;use&#34;,
<a id="L626"></a>}


<a id="L629"></a><span class="comment">// Template formatter for &#34;infoKind&#34; format.</span>
<a id="L630"></a>func infoKindFmt(w io.Writer, x interface{}, format string) {
    <a id="L631"></a>fmt.Fprintf(w, infoKinds[x.(SpotKind)]) <span class="comment">// infoKind entries are html-escaped</span>
<a id="L632"></a>}


<a id="L635"></a><span class="comment">// Template formatter for &#34;infoLine&#34; format.</span>
<a id="L636"></a>func infoLineFmt(w io.Writer, x interface{}, format string) {
    <a id="L637"></a>info := x.(SpotInfo);
    <a id="L638"></a>line := info.Lori();
    <a id="L639"></a>if info.IsIndex() {
        <a id="L640"></a>index, _ := searchIndex.get();
        <a id="L641"></a>line = index.(*Index).Snippet(line).Line;
    <a id="L642"></a>}
    <a id="L643"></a>fmt.Fprintf(w, &#34;%d&#34;, line);
<a id="L644"></a>}


<a id="L647"></a><span class="comment">// Template formatter for &#34;infoSnippet&#34; format.</span>
<a id="L648"></a>func infoSnippetFmt(w io.Writer, x interface{}, format string) {
    <a id="L649"></a>info := x.(SpotInfo);
    <a id="L650"></a>text := `&lt;span class=&#34;alert&#34;&gt;no snippet text available&lt;/span&gt;`;
    <a id="L651"></a>if info.IsIndex() {
        <a id="L652"></a>index, _ := searchIndex.get();
        <a id="L653"></a><span class="comment">// no escaping of snippet text needed;</span>
        <a id="L654"></a><span class="comment">// snippet text is escaped when generated</span>
        <a id="L655"></a>text = index.(*Index).Snippet(info.Lori()).Text;
    <a id="L656"></a>}
    <a id="L657"></a>fmt.Fprint(w, text);
<a id="L658"></a>}


<a id="L661"></a><span class="comment">// Template formatter for &#34;padding&#34; format.</span>
<a id="L662"></a>func paddingFmt(w io.Writer, x interface{}, format string) {
    <a id="L663"></a>for i := x.(int); i &gt; 0; i-- {
        <a id="L664"></a>fmt.Fprint(w, `&lt;td width=&#34;25&#34;&gt;&lt;/td&gt;`)
    <a id="L665"></a>}
<a id="L666"></a>}


<a id="L669"></a><span class="comment">// Template formatter for &#34;time&#34; format.</span>
<a id="L670"></a>func timeFmt(w io.Writer, x interface{}, format string) {
    <a id="L671"></a><span class="comment">// note: os.Dir.Mtime_ns is in uint64 in ns!</span>
    <a id="L672"></a>template.HTMLEscape(w, strings.Bytes(time.SecondsToLocalTime(int64(x.(uint64)/1e9)).String()))
<a id="L673"></a>}


<a id="L676"></a>var fmap = template.FormatterMap{
    <a id="L677"></a>&#34;&#34;: textFmt,
    <a id="L678"></a>&#34;html&#34;: htmlFmt,
    <a id="L679"></a>&#34;html-comment&#34;: htmlCommentFmt,
    <a id="L680"></a>&#34;path&#34;: pathFmt,
    <a id="L681"></a>&#34;link&#34;: linkFmt,
    <a id="L682"></a>&#34;infoKind&#34;: infoKindFmt,
    <a id="L683"></a>&#34;infoLine&#34;: infoLineFmt,
    <a id="L684"></a>&#34;infoSnippet&#34;: infoSnippetFmt,
    <a id="L685"></a>&#34;padding&#34;: paddingFmt,
    <a id="L686"></a>&#34;time&#34;: timeFmt,
<a id="L687"></a>}


<a id="L690"></a>func readTemplate(name string) *template.Template {
    <a id="L691"></a>path := pathutil.Join(*tmplroot, name);
    <a id="L692"></a>data, err := io.ReadFile(path);
    <a id="L693"></a>if err != nil {
        <a id="L694"></a>log.Exitf(&#34;ReadFile %s: %v&#34;, path, err)
    <a id="L695"></a>}
    <a id="L696"></a>t, err := template.Parse(string(data), fmap);
    <a id="L697"></a>if err != nil {
        <a id="L698"></a>log.Exitf(&#34;%s: %v&#34;, name, err)
    <a id="L699"></a>}
    <a id="L700"></a>return t;
<a id="L701"></a>}


<a id="L704"></a>var (
    <a id="L705"></a>dirlistHTML,
        <a id="L706"></a>godocHTML,
        <a id="L707"></a>packageHTML,
        <a id="L708"></a>packageText,
        <a id="L709"></a>parseerrorHTML,
        <a id="L710"></a>parseerrorText,
        <a id="L711"></a>searchHTML *template.Template;
<a id="L712"></a>)

<a id="L714"></a>func readTemplates() {
    <a id="L715"></a><span class="comment">// have to delay until after flags processing,</span>
    <a id="L716"></a><span class="comment">// so that main has chdir&#39;ed to goroot.</span>
    <a id="L717"></a>dirlistHTML = readTemplate(&#34;dirlist.html&#34;);
    <a id="L718"></a>godocHTML = readTemplate(&#34;godoc.html&#34;);
    <a id="L719"></a>packageHTML = readTemplate(&#34;package.html&#34;);
    <a id="L720"></a>packageText = readTemplate(&#34;package.txt&#34;);
    <a id="L721"></a>parseerrorHTML = readTemplate(&#34;parseerror.html&#34;);
    <a id="L722"></a>parseerrorText = readTemplate(&#34;parseerror.txt&#34;);
    <a id="L723"></a>searchHTML = readTemplate(&#34;search.html&#34;);
<a id="L724"></a>}


<a id="L727"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L728"></a><span class="comment">// Generic HTML wrapper</span>

<a id="L730"></a>func servePage(c *http.Conn, title, query string, content []byte) {
    <a id="L731"></a>type Data struct {
        <a id="L732"></a>Title     string;
        <a id="L733"></a>Timestamp uint64; <span class="comment">// int64 to be compatible with os.Dir.Mtime_ns</span>
        <a id="L734"></a>Query     string;
        <a id="L735"></a>Content   []byte;
    <a id="L736"></a>}

    <a id="L738"></a>_, ts := fsTree.get();
    <a id="L739"></a>d := Data{
        <a id="L740"></a>Title: title,
        <a id="L741"></a>Timestamp: uint64(ts) * 1e9, <span class="comment">// timestamp in ns</span>
        <a id="L742"></a>Query: query,
        <a id="L743"></a>Content: content,
    <a id="L744"></a>};

    <a id="L746"></a>if err := godocHTML.Execute(&amp;d, c); err != nil {
        <a id="L747"></a>log.Stderrf(&#34;godocHTML.Execute: %s&#34;, err)
    <a id="L748"></a>}
<a id="L749"></a>}


<a id="L752"></a>func serveText(c *http.Conn, text []byte) {
    <a id="L753"></a>c.SetHeader(&#34;content-type&#34;, &#34;text/plain; charset=utf-8&#34;);
    <a id="L754"></a>c.Write(text);
<a id="L755"></a>}


<a id="L758"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L759"></a><span class="comment">// Files</span>

<a id="L761"></a>var (
    <a id="L762"></a>tagBegin = strings.Bytes(&#34;&lt;!--&#34;);
    <a id="L763"></a>tagEnd   = strings.Bytes(&#34;--&gt;&#34;);
<a id="L764"></a>)

<a id="L766"></a><span class="comment">// commentText returns the text of the first HTML comment in src.</span>
<a id="L767"></a>func commentText(src []byte) (text string) {
    <a id="L768"></a>i := bytes.Index(src, tagBegin);
    <a id="L769"></a>j := bytes.Index(src, tagEnd);
    <a id="L770"></a>if i &gt;= 0 &amp;&amp; j &gt;= i+len(tagBegin) {
        <a id="L771"></a>text = string(bytes.TrimSpace(src[i+len(tagBegin) : j]))
    <a id="L772"></a>}
    <a id="L773"></a>return;
<a id="L774"></a>}


<a id="L777"></a>func serveHTMLDoc(c *http.Conn, r *http.Request, path string) {
    <a id="L778"></a><span class="comment">// get HTML body contents</span>
    <a id="L779"></a>src, err := io.ReadFile(path);
    <a id="L780"></a>if err != nil {
        <a id="L781"></a>log.Stderrf(&#34;%v&#34;, err);
        <a id="L782"></a>http.NotFound(c, r);
        <a id="L783"></a>return;
    <a id="L784"></a>}

    <a id="L786"></a><span class="comment">// if it&#39;s the language spec, add tags to EBNF productions</span>
    <a id="L787"></a>if strings.HasSuffix(path, &#34;go_spec.html&#34;) {
        <a id="L788"></a>var buf bytes.Buffer;
        <a id="L789"></a>linkify(&amp;buf, src);
        <a id="L790"></a>src = buf.Bytes();
    <a id="L791"></a>}

    <a id="L793"></a>title := commentText(src);
    <a id="L794"></a>servePage(c, title, &#34;&#34;, src);
<a id="L795"></a>}


<a id="L798"></a>func serveParseErrors(c *http.Conn, errors *parseErrors) {
    <a id="L799"></a><span class="comment">// format errors</span>
    <a id="L800"></a>var buf bytes.Buffer;
    <a id="L801"></a>if err := parseerrorHTML.Execute(errors, &amp;buf); err != nil {
        <a id="L802"></a>log.Stderrf(&#34;parseerrorHTML.Execute: %s&#34;, err)
    <a id="L803"></a>}
    <a id="L804"></a>servePage(c, &#34;Parse errors in source file &#34;+errors.filename, &#34;&#34;, buf.Bytes());
<a id="L805"></a>}


<a id="L808"></a>func serveGoSource(c *http.Conn, r *http.Request, path string, styler printer.Styler) {
    <a id="L809"></a>prog, errors := parse(path, parser.ParseComments);
    <a id="L810"></a>if errors != nil {
        <a id="L811"></a>serveParseErrors(c, errors);
        <a id="L812"></a>return;
    <a id="L813"></a>}

    <a id="L815"></a>var buf bytes.Buffer;
    <a id="L816"></a>fmt.Fprintln(&amp;buf, &#34;&lt;pre&gt;&#34;);
    <a id="L817"></a>writeNode(&amp;buf, prog, true, styler);
    <a id="L818"></a>fmt.Fprintln(&amp;buf, &#34;&lt;/pre&gt;&#34;);

    <a id="L820"></a>servePage(c, &#34;Source file &#34;+r.URL.Path, &#34;&#34;, buf.Bytes());
<a id="L821"></a>}


<a id="L824"></a>func redirect(c *http.Conn, r *http.Request) (redirected bool) {
    <a id="L825"></a>if canonical := pathutil.Clean(r.URL.Path) + &#34;/&#34;; r.URL.Path != canonical {
        <a id="L826"></a>http.Redirect(c, canonical, http.StatusMovedPermanently);
        <a id="L827"></a>redirected = true;
    <a id="L828"></a>}
    <a id="L829"></a>return;
<a id="L830"></a>}


<a id="L833"></a><span class="comment">// TODO(gri): Should have a mapping from extension to handler, eventually.</span>

<a id="L835"></a><span class="comment">// textExt[x] is true if the extension x indicates a text file, and false otherwise.</span>
<a id="L836"></a>var textExt = map[string]bool{
    <a id="L837"></a>&#34;.css&#34;: false, <span class="comment">// must be served raw</span>
    <a id="L838"></a>&#34;.js&#34;: false, <span class="comment">// must be served raw</span>
<a id="L839"></a>}


<a id="L842"></a>func isTextFile(path string) bool {
    <a id="L843"></a><span class="comment">// if the extension is known, use it for decision making</span>
    <a id="L844"></a>if isText, found := textExt[pathutil.Ext(path)]; found {
        <a id="L845"></a>return isText
    <a id="L846"></a>}

    <a id="L848"></a><span class="comment">// the extension is not known; read an initial chunk of</span>
    <a id="L849"></a><span class="comment">// file and check if it looks like correct UTF-8; if it</span>
    <a id="L850"></a><span class="comment">// does, it&#39;s probably a text file</span>
    <a id="L851"></a>f, err := os.Open(path, os.O_RDONLY, 0);
    <a id="L852"></a>if err != nil {
        <a id="L853"></a>return false
    <a id="L854"></a>}

    <a id="L856"></a>var buf [1024]byte;
    <a id="L857"></a>n, err := f.Read(&amp;buf);
    <a id="L858"></a>if err != nil {
        <a id="L859"></a>return false
    <a id="L860"></a>}

    <a id="L862"></a>s := string(buf[0:n]);
    <a id="L863"></a>n -= utf8.UTFMax; <span class="comment">// make sure there&#39;s enough bytes for a complete unicode char</span>
    <a id="L864"></a>for i, c := range s {
        <a id="L865"></a>if i &gt; n {
            <a id="L866"></a>break
        <a id="L867"></a>}
        <a id="L868"></a>if c == 0xFFFD || c &lt; &#39; &#39; &amp;&amp; c != &#39;\n&#39; &amp;&amp; c != &#39;\t&#39; {
            <a id="L869"></a><span class="comment">// decoding error or control character - not a text file</span>
            <a id="L870"></a>return false
        <a id="L871"></a>}
    <a id="L872"></a>}

    <a id="L874"></a><span class="comment">// likely a text file</span>
    <a id="L875"></a>return true;
<a id="L876"></a>}


<a id="L879"></a>func serveTextFile(c *http.Conn, r *http.Request, path string) {
    <a id="L880"></a>src, err := io.ReadFile(path);
    <a id="L881"></a>if err != nil {
        <a id="L882"></a>log.Stderrf(&#34;serveTextFile: %s&#34;, err)
    <a id="L883"></a>}

    <a id="L885"></a>var buf bytes.Buffer;
    <a id="L886"></a>fmt.Fprintln(&amp;buf, &#34;&lt;pre&gt;&#34;);
    <a id="L887"></a>template.HTMLEscape(&amp;buf, src);
    <a id="L888"></a>fmt.Fprintln(&amp;buf, &#34;&lt;/pre&gt;&#34;);

    <a id="L890"></a>servePage(c, &#34;Text file &#34;+path, &#34;&#34;, buf.Bytes());
<a id="L891"></a>}


<a id="L894"></a>func serveDirectory(c *http.Conn, r *http.Request, path string) {
    <a id="L895"></a>if redirect(c, r) {
        <a id="L896"></a>return
    <a id="L897"></a>}

    <a id="L899"></a>list, err := io.ReadDir(path);
    <a id="L900"></a>if err != nil {
        <a id="L901"></a>http.NotFound(c, r);
        <a id="L902"></a>return;
    <a id="L903"></a>}

    <a id="L905"></a>var buf bytes.Buffer;
    <a id="L906"></a>if err := dirlistHTML.Execute(list, &amp;buf); err != nil {
        <a id="L907"></a>log.Stderrf(&#34;dirlistHTML.Execute: %s&#34;, err)
    <a id="L908"></a>}

    <a id="L910"></a>servePage(c, &#34;Directory &#34;+path, &#34;&#34;, buf.Bytes());
<a id="L911"></a>}


<a id="L914"></a>var fileServer = http.FileServer(&#34;.&#34;, &#34;&#34;)

<a id="L916"></a>func serveFile(c *http.Conn, r *http.Request) {
    <a id="L917"></a>path := pathutil.Join(&#34;.&#34;, r.URL.Path);

    <a id="L919"></a><span class="comment">// pick off special cases and hand the rest to the standard file server</span>
    <a id="L920"></a>switch ext := pathutil.Ext(path); {
    <a id="L921"></a>case r.URL.Path == &#34;/&#34;:
        <a id="L922"></a>serveHTMLDoc(c, r, &#34;doc/root.html&#34;);
        <a id="L923"></a>return;

    <a id="L925"></a>case r.URL.Path == &#34;/doc/root.html&#34;:
        <a id="L926"></a><span class="comment">// hide landing page from its real name</span>
        <a id="L927"></a>http.NotFound(c, r);
        <a id="L928"></a>return;

    <a id="L930"></a>case ext == &#34;.html&#34;:
        <a id="L931"></a>serveHTMLDoc(c, r, path);
        <a id="L932"></a>return;

    <a id="L934"></a>case ext == &#34;.go&#34;:
        <a id="L935"></a>serveGoSource(c, r, path, &amp;Styler{highlight: r.FormValue(&#34;h&#34;)});
        <a id="L936"></a>return;
    <a id="L937"></a>}

    <a id="L939"></a>dir, err := os.Lstat(path);
    <a id="L940"></a>if err != nil {
        <a id="L941"></a>http.NotFound(c, r);
        <a id="L942"></a>return;
    <a id="L943"></a>}

    <a id="L945"></a>if dir != nil &amp;&amp; dir.IsDirectory() {
        <a id="L946"></a>serveDirectory(c, r, path);
        <a id="L947"></a>return;
    <a id="L948"></a>}

    <a id="L950"></a>if isTextFile(path) {
        <a id="L951"></a>serveTextFile(c, r, path);
        <a id="L952"></a>return;
    <a id="L953"></a>}

    <a id="L955"></a>fileServer.ServeHTTP(c, r);
<a id="L956"></a>}


<a id="L959"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L960"></a><span class="comment">// Packages</span>

<a id="L962"></a><span class="comment">// Package name used for commands that have non-identifier names.</span>
<a id="L963"></a>const fakePkgName = &#34;documentation&#34;


<a id="L966"></a>type PageInfo struct {
    <a id="L967"></a>PDoc  *doc.PackageDoc; <span class="comment">// nil if no package found</span>
    <a id="L968"></a>Dirs  *DirList;        <span class="comment">// nil if no directory information found</span>
    <a id="L969"></a>IsPkg bool;            <span class="comment">// false if this is not documenting a real package</span>
<a id="L970"></a>}


<a id="L973"></a>type httpHandler struct {
    <a id="L974"></a>pattern string; <span class="comment">// url pattern; e.g. &#34;/pkg/&#34;</span>
    <a id="L975"></a>fsRoot  string; <span class="comment">// file system root to which the pattern is mapped</span>
    <a id="L976"></a>isPkg   bool;   <span class="comment">// true if this handler serves real package documentation (as opposed to command documentation)</span>
<a id="L977"></a>}


<a id="L980"></a><span class="comment">// getPageInfo returns the PageInfo for a given package directory.</span>
<a id="L981"></a><span class="comment">// If there is no corresponding package in the directory,</span>
<a id="L982"></a><span class="comment">// PageInfo.PDoc is nil. If there are no subdirectories,</span>
<a id="L983"></a><span class="comment">// PageInfo.Dirs is nil.</span>
<a id="L984"></a><span class="comment">//</span>
<a id="L985"></a>func (h *httpHandler) getPageInfo(path string) PageInfo {
    <a id="L986"></a><span class="comment">// the path is relative to h.fsroot</span>
    <a id="L987"></a>dirname := pathutil.Join(h.fsRoot, path);

    <a id="L989"></a><span class="comment">// the package name is the directory name within its parent</span>
    <a id="L990"></a><span class="comment">// (use dirname instead of path because dirname is clean; i.e. has no trailing &#39;/&#39;)</span>
    <a id="L991"></a>_, pkgname := pathutil.Split(dirname);

    <a id="L993"></a><span class="comment">// filter function to select the desired .go files</span>
    <a id="L994"></a>filter := func(d *os.Dir) bool {
        <a id="L995"></a>if isPkgFile(d) {
            <a id="L996"></a><span class="comment">// Some directories contain main packages: Only accept</span>
            <a id="L997"></a><span class="comment">// files that belong to the expected package so that</span>
            <a id="L998"></a><span class="comment">// parser.ParsePackage doesn&#39;t return &#34;multiple packages</span>
            <a id="L999"></a><span class="comment">// found&#34; errors.</span>
            <a id="L1000"></a><span class="comment">// Additionally, accept the special package name</span>
            <a id="L1001"></a><span class="comment">// fakePkgName if we are looking at cmd documentation.</span>
            <a id="L1002"></a>name := pkgName(dirname + &#34;/&#34; + d.Name);
            <a id="L1003"></a>return name == pkgname || h.fsRoot == *cmdroot &amp;&amp; name == fakePkgName;
        <a id="L1004"></a>}
        <a id="L1005"></a>return false;
    <a id="L1006"></a>};

    <a id="L1008"></a><span class="comment">// get package AST</span>
    <a id="L1009"></a>pkg, err := parser.ParsePackage(dirname, filter, parser.ParseComments);
    <a id="L1010"></a>if err != nil {
        <a id="L1011"></a><span class="comment">// TODO: parse errors should be shown instead of an empty directory</span>
        <a id="L1012"></a>log.Stderrf(&#34;parser.parsePackage: %s&#34;, err)
    <a id="L1013"></a>}

    <a id="L1015"></a><span class="comment">// compute package documentation</span>
    <a id="L1016"></a>var pdoc *doc.PackageDoc;
    <a id="L1017"></a>if pkg != nil {
        <a id="L1018"></a>ast.PackageExports(pkg);
        <a id="L1019"></a>pdoc = doc.NewPackageDoc(pkg, pathutil.Clean(path)); <span class="comment">// no trailing &#39;/&#39; in importpath</span>
    <a id="L1020"></a>}

    <a id="L1022"></a><span class="comment">// get directory information</span>
    <a id="L1023"></a>var dir *Directory;
    <a id="L1024"></a>if tree, _ := fsTree.get(); tree != nil {
        <a id="L1025"></a><span class="comment">// directory tree is present; lookup respective directory</span>
        <a id="L1026"></a><span class="comment">// (may still fail if the file system was updated and the</span>
        <a id="L1027"></a><span class="comment">// new directory tree has not yet beet computed)</span>
        <a id="L1028"></a>dir = tree.(*Directory).lookup(dirname)
    <a id="L1029"></a>} else {
        <a id="L1030"></a><span class="comment">// no directory tree present (either early after startup</span>
        <a id="L1031"></a><span class="comment">// or command-line mode); compute one level for this page</span>
        <a id="L1032"></a>dir = newDirectory(dirname, 1)
    <a id="L1033"></a>}

    <a id="L1035"></a>return PageInfo{pdoc, dir.listing(true), h.isPkg};
<a id="L1036"></a>}


<a id="L1039"></a>func (h *httpHandler) ServeHTTP(c *http.Conn, r *http.Request) {
    <a id="L1040"></a>if redirect(c, r) {
        <a id="L1041"></a>return
    <a id="L1042"></a>}

    <a id="L1044"></a>path := r.URL.Path;
    <a id="L1045"></a>path = path[len(h.pattern):len(path)];
    <a id="L1046"></a>info := h.getPageInfo(path);

    <a id="L1048"></a>var buf bytes.Buffer;
    <a id="L1049"></a>if r.FormValue(&#34;f&#34;) == &#34;text&#34; {
        <a id="L1050"></a>if err := packageText.Execute(info, &amp;buf); err != nil {
            <a id="L1051"></a>log.Stderrf(&#34;packageText.Execute: %s&#34;, err)
        <a id="L1052"></a>}
        <a id="L1053"></a>serveText(c, buf.Bytes());
        <a id="L1054"></a>return;
    <a id="L1055"></a>}

    <a id="L1057"></a>if err := packageHTML.Execute(info, &amp;buf); err != nil {
        <a id="L1058"></a>log.Stderrf(&#34;packageHTML.Execute: %s&#34;, err)
    <a id="L1059"></a>}

    <a id="L1061"></a>if path == &#34;&#34; {
        <a id="L1062"></a>path = &#34;.&#34; <span class="comment">// don&#39;t display an empty path</span>
    <a id="L1063"></a>}
    <a id="L1064"></a>title := &#34;Directory &#34; + path;
    <a id="L1065"></a>if info.PDoc != nil {
        <a id="L1066"></a>switch {
        <a id="L1067"></a>case h.isPkg:
            <a id="L1068"></a>title = &#34;Package &#34; + info.PDoc.PackageName
        <a id="L1069"></a>case info.PDoc.PackageName == fakePkgName:
            <a id="L1070"></a><span class="comment">// assume that the directory name is the command name</span>
            <a id="L1071"></a>_, pkgname := pathutil.Split(pathutil.Clean(path));
            <a id="L1072"></a>title = &#34;Command &#34; + pkgname;
        <a id="L1073"></a>default:
            <a id="L1074"></a>title = &#34;Command &#34; + info.PDoc.PackageName
        <a id="L1075"></a>}
    <a id="L1076"></a>}

    <a id="L1078"></a>servePage(c, title, &#34;&#34;, buf.Bytes());
<a id="L1079"></a>}


<a id="L1082"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L1083"></a><span class="comment">// Search</span>

<a id="L1085"></a>var searchIndex RWValue

<a id="L1087"></a>type SearchResult struct {
    <a id="L1088"></a>Query    string;
    <a id="L1089"></a>Hit      *LookupResult;
    <a id="L1090"></a>Alt      *AltWords;
    <a id="L1091"></a>Illegal  bool;
    <a id="L1092"></a>Accurate bool;
<a id="L1093"></a>}

<a id="L1095"></a>func search(c *http.Conn, r *http.Request) {
    <a id="L1096"></a>query := r.FormValue(&#34;q&#34;);
    <a id="L1097"></a>var result SearchResult;

    <a id="L1099"></a>if index, timestamp := searchIndex.get(); index != nil {
        <a id="L1100"></a>result.Query = query;
        <a id="L1101"></a>result.Hit, result.Alt, result.Illegal = index.(*Index).Lookup(query);
        <a id="L1102"></a>_, ts := fsTree.get();
        <a id="L1103"></a>result.Accurate = timestamp &gt;= ts;
    <a id="L1104"></a>}

    <a id="L1106"></a>var buf bytes.Buffer;
    <a id="L1107"></a>if err := searchHTML.Execute(result, &amp;buf); err != nil {
        <a id="L1108"></a>log.Stderrf(&#34;searchHTML.Execute: %s&#34;, err)
    <a id="L1109"></a>}

    <a id="L1111"></a>var title string;
    <a id="L1112"></a>if result.Hit != nil {
        <a id="L1113"></a>title = fmt.Sprintf(`Results for query %q`, query)
    <a id="L1114"></a>} else {
        <a id="L1115"></a>title = fmt.Sprintf(`No results found for query %q`, query)
    <a id="L1116"></a>}

    <a id="L1118"></a>servePage(c, title, query, buf.Bytes());
<a id="L1119"></a>}


<a id="L1122"></a><span class="comment">// ----------------------------------------------------------------------------</span>
<a id="L1123"></a><span class="comment">// Server</span>

<a id="L1125"></a>var (
    <a id="L1126"></a>cmdHandler = httpHandler{&#34;/cmd/&#34;, *cmdroot, false};
    <a id="L1127"></a>pkgHandler = httpHandler{&#34;/pkg/&#34;, *pkgroot, true};
<a id="L1128"></a>)


<a id="L1131"></a>func registerPublicHandlers(mux *http.ServeMux) {
    <a id="L1132"></a>mux.Handle(cmdHandler.pattern, &amp;cmdHandler);
    <a id="L1133"></a>mux.Handle(pkgHandler.pattern, &amp;pkgHandler);
    <a id="L1134"></a>mux.Handle(&#34;/search&#34;, http.HandlerFunc(search));
    <a id="L1135"></a>mux.Handle(&#34;/&#34;, http.HandlerFunc(serveFile));
<a id="L1136"></a>}


<a id="L1139"></a><span class="comment">// Indexing goroutine.</span>
<a id="L1140"></a>func indexer() {
    <a id="L1141"></a>for {
        <a id="L1142"></a>_, ts := fsTree.get();
        <a id="L1143"></a>if _, timestamp := searchIndex.get(); timestamp &lt; ts {
            <a id="L1144"></a><span class="comment">// index possibly out of date - make a new one</span>
            <a id="L1145"></a><span class="comment">// (could use a channel to send an explicit signal</span>
            <a id="L1146"></a><span class="comment">// from the sync goroutine, but this solution is</span>
            <a id="L1147"></a><span class="comment">// more decoupled, trivial, and works well enough)</span>
            <a id="L1148"></a>start := time.Nanoseconds();
            <a id="L1149"></a>index := NewIndex(&#34;.&#34;);
            <a id="L1150"></a>stop := time.Nanoseconds();
            <a id="L1151"></a>searchIndex.set(index);
            <a id="L1152"></a>if *verbose {
                <a id="L1153"></a>secs := float64((stop-start)/1e6) / 1e3;
                <a id="L1154"></a>nwords, nspots := index.Size();
                <a id="L1155"></a>log.Stderrf(&#34;index updated (%gs, %d unique words, %d spots)&#34;, secs, nwords, nspots);
            <a id="L1156"></a>}
        <a id="L1157"></a>}
        <a id="L1158"></a>time.Sleep(1 * 60e9); <span class="comment">// try once a minute</span>
    <a id="L1159"></a>}
<a id="L1160"></a>}
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
