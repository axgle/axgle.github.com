<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/path/path.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/path/path.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// The path package implements utility routines for manipulating</span>
<a id="L6"></a><span class="comment">// slash-separated filename paths.</span>
<a id="L7"></a>package path

<a id="L9"></a>import (
    <a id="L10"></a>&#34;io&#34;;
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;strings&#34;;
<a id="L13"></a>)

<a id="L15"></a><span class="comment">// Clean returns the shortest path name equivalent to path</span>
<a id="L16"></a><span class="comment">// by purely lexical processing.  It applies the following rules</span>
<a id="L17"></a><span class="comment">// iteratively until no further processing can be done:</span>
<a id="L18"></a><span class="comment">//</span>
<a id="L19"></a><span class="comment">//	1. Replace multiple slashes with a single slash.</span>
<a id="L20"></a><span class="comment">//	2. Eliminate each . path name element (the current directory).</span>
<a id="L21"></a><span class="comment">//	3. Eliminate each inner .. path name element (the parent directory)</span>
<a id="L22"></a><span class="comment">//	   along with the non-.. element that precedes it.</span>
<a id="L23"></a><span class="comment">//	4. Eliminate .. elements that begin a rooted path:</span>
<a id="L24"></a><span class="comment">//	   that is, replace &#34;/..&#34; by &#34;/&#34; at the beginning of a path.</span>
<a id="L25"></a><span class="comment">//</span>
<a id="L26"></a><span class="comment">// If the result of this process is an empty string, Clean</span>
<a id="L27"></a><span class="comment">// returns the string &#34;.&#34;.</span>
<a id="L28"></a><span class="comment">//</span>
<a id="L29"></a><span class="comment">// See also Rob Pike, ``Lexical File Names in Plan 9 or</span>
<a id="L30"></a><span class="comment">// Getting Dot-Dot right,&#39;&#39;</span>
<a id="L31"></a><span class="comment">// http://plan9.bell-labs.com/sys/doc/lexnames.html</span>
<a id="L32"></a>func Clean(path string) string {
    <a id="L33"></a>if path == &#34;&#34; {
        <a id="L34"></a>return &#34;.&#34;
    <a id="L35"></a>}

    <a id="L37"></a>rooted := path[0] == &#39;/&#39;;
    <a id="L38"></a>n := len(path);

    <a id="L40"></a><span class="comment">// Invariants:</span>
    <a id="L41"></a><span class="comment">//	reading from path; r is index of next byte to process.</span>
    <a id="L42"></a><span class="comment">//	writing to buf; w is index of next byte to write.</span>
    <a id="L43"></a><span class="comment">//	dotdot is index in buf where .. must stop, either because</span>
    <a id="L44"></a><span class="comment">//		it is the leading slash or it is a leading ../../.. prefix.</span>
    <a id="L45"></a>buf := strings.Bytes(path);
    <a id="L46"></a>r, w, dotdot := 0, 0, 0;
    <a id="L47"></a>if rooted {
        <a id="L48"></a>r, w, dotdot = 1, 1, 1
    <a id="L49"></a>}

    <a id="L51"></a>for r &lt; n {
        <a id="L52"></a>switch {
        <a id="L53"></a>case path[r] == &#39;/&#39;:
            <a id="L54"></a><span class="comment">// empty path element</span>
            <a id="L55"></a>r++
        <a id="L56"></a>case path[r] == &#39;.&#39; &amp;&amp; (r+1 == n || path[r+1] == &#39;/&#39;):
            <a id="L57"></a><span class="comment">// . element</span>
            <a id="L58"></a>r++
        <a id="L59"></a>case path[r] == &#39;.&#39; &amp;&amp; path[r+1] == &#39;.&#39; &amp;&amp; (r+2 == n || path[r+2] == &#39;/&#39;):
            <a id="L60"></a><span class="comment">// .. element: remove to last /</span>
            <a id="L61"></a>r += 2;
            <a id="L62"></a>switch {
            <a id="L63"></a>case w &gt; dotdot:
                <a id="L64"></a><span class="comment">// can backtrack</span>
                <a id="L65"></a>w--;
                <a id="L66"></a>for w &gt; dotdot &amp;&amp; buf[w] != &#39;/&#39; {
                    <a id="L67"></a>w--
                <a id="L68"></a>}
            <a id="L69"></a>case !rooted:
                <a id="L70"></a><span class="comment">// cannot backtrack, but not rooted, so append .. element.</span>
                <a id="L71"></a>if w &gt; 0 {
                    <a id="L72"></a>buf[w] = &#39;/&#39;;
                    <a id="L73"></a>w++;
                <a id="L74"></a>}
                <a id="L75"></a>buf[w] = &#39;.&#39;;
                <a id="L76"></a>w++;
                <a id="L77"></a>buf[w] = &#39;.&#39;;
                <a id="L78"></a>w++;
                <a id="L79"></a>dotdot = w;
            <a id="L80"></a>}
        <a id="L81"></a>default:
            <a id="L82"></a><span class="comment">// real path element.</span>
            <a id="L83"></a><span class="comment">// add slash if needed</span>
            <a id="L84"></a>if rooted &amp;&amp; w != 1 || !rooted &amp;&amp; w != 0 {
                <a id="L85"></a>buf[w] = &#39;/&#39;;
                <a id="L86"></a>w++;
            <a id="L87"></a>}
            <a id="L88"></a><span class="comment">// copy element</span>
            <a id="L89"></a>for ; r &lt; n &amp;&amp; path[r] != &#39;/&#39;; r++ {
                <a id="L90"></a>buf[w] = path[r];
                <a id="L91"></a>w++;
            <a id="L92"></a>}
        <a id="L93"></a>}
    <a id="L94"></a>}

    <a id="L96"></a><span class="comment">// Turn empty string into &#34;.&#34;</span>
    <a id="L97"></a>if w == 0 {
        <a id="L98"></a>buf[w] = &#39;.&#39;;
        <a id="L99"></a>w++;
    <a id="L100"></a>}

    <a id="L102"></a>return string(buf[0:w]);
<a id="L103"></a>}

<a id="L105"></a><span class="comment">// Split splits path immediately following the final slash,</span>
<a id="L106"></a><span class="comment">// separating it into a directory and file name component.</span>
<a id="L107"></a><span class="comment">// If there is no slash in path, DirFile returns an empty dir and</span>
<a id="L108"></a><span class="comment">// file set to path.</span>
<a id="L109"></a>func Split(path string) (dir, file string) {
    <a id="L110"></a>for i := len(path) - 1; i &gt;= 0; i-- {
        <a id="L111"></a>if path[i] == &#39;/&#39; {
            <a id="L112"></a>return path[0 : i+1], path[i+1 : len(path)]
        <a id="L113"></a>}
    <a id="L114"></a>}
    <a id="L115"></a>return &#34;&#34;, path;
<a id="L116"></a>}

<a id="L118"></a><span class="comment">// Join joins dir and file into a single path, adding a separating</span>
<a id="L119"></a><span class="comment">// slash if necessary.  If dir is empty, it returns file.</span>
<a id="L120"></a>func Join(dir, file string) string {
    <a id="L121"></a>if dir == &#34;&#34; {
        <a id="L122"></a>return file
    <a id="L123"></a>}
    <a id="L124"></a>return Clean(dir + &#34;/&#34; + file);
<a id="L125"></a>}

<a id="L127"></a><span class="comment">// Ext returns the file name extension used by path.</span>
<a id="L128"></a><span class="comment">// The extension is the suffix beginning at the final dot</span>
<a id="L129"></a><span class="comment">// in the final slash-separated element of path;</span>
<a id="L130"></a><span class="comment">// it is empty if there is no dot.</span>
<a id="L131"></a>func Ext(path string) string {
    <a id="L132"></a>for i := len(path) - 1; i &gt;= 0 &amp;&amp; path[i] != &#39;/&#39;; i-- {
        <a id="L133"></a>if path[i] == &#39;.&#39; {
            <a id="L134"></a>return path[i:len(path)]
        <a id="L135"></a>}
    <a id="L136"></a>}
    <a id="L137"></a>return &#34;&#34;;
<a id="L138"></a>}

<a id="L140"></a><span class="comment">// Visitor methods are invoked for corresponding file tree entries</span>
<a id="L141"></a><span class="comment">// visited by Walk. The parameter path is the full path of d relative</span>
<a id="L142"></a><span class="comment">// to root.</span>
<a id="L143"></a>type Visitor interface {
    <a id="L144"></a>VisitDir(path string, d *os.Dir) bool;
    <a id="L145"></a>VisitFile(path string, d *os.Dir);
<a id="L146"></a>}

<a id="L148"></a>func walk(path string, d *os.Dir, v Visitor, errors chan&lt;- os.Error) {
    <a id="L149"></a>if !d.IsDirectory() {
        <a id="L150"></a>v.VisitFile(path, d);
        <a id="L151"></a>return;
    <a id="L152"></a>}

    <a id="L154"></a>if !v.VisitDir(path, d) {
        <a id="L155"></a>return <span class="comment">// skip directory entries</span>
    <a id="L156"></a>}

    <a id="L158"></a>list, err := io.ReadDir(path);
    <a id="L159"></a>if err != nil {
        <a id="L160"></a>if errors != nil {
            <a id="L161"></a>errors &lt;- err
        <a id="L162"></a>}
    <a id="L163"></a>}

    <a id="L165"></a>for _, e := range list {
        <a id="L166"></a>walk(Join(path, e.Name), e, v, errors)
    <a id="L167"></a>}
<a id="L168"></a>}

<a id="L170"></a><span class="comment">// Walk walks the file tree rooted at root, calling v.VisitDir or</span>
<a id="L171"></a><span class="comment">// v.VisitFile for each directory or file in the tree, including root.</span>
<a id="L172"></a><span class="comment">// If v.VisitDir returns false, Walk skips the directory&#39;s entries;</span>
<a id="L173"></a><span class="comment">// otherwise it invokes itself for each directory entry in sorted order.</span>
<a id="L174"></a><span class="comment">// An error reading a directory does not abort the Walk.</span>
<a id="L175"></a><span class="comment">// If errors != nil, Walk sends each directory read error</span>
<a id="L176"></a><span class="comment">// to the channel.  Otherwise Walk discards the error.</span>
<a id="L177"></a>func Walk(root string, v Visitor, errors chan&lt;- os.Error) {
    <a id="L178"></a>d, err := os.Lstat(root);
    <a id="L179"></a>if err != nil {
        <a id="L180"></a>if errors != nil {
            <a id="L181"></a>errors &lt;- err
        <a id="L182"></a>}
        <a id="L183"></a>return; <span class="comment">// can&#39;t progress</span>
    <a id="L184"></a>}
    <a id="L185"></a>walk(root, d, v, errors);
<a id="L186"></a>}
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
