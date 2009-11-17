<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/fs.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/http/fs.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// HTTP file system request handler</span>

<a id="L7"></a>package http

<a id="L9"></a>import (
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;io&#34;;
    <a id="L12"></a>&#34;os&#34;;
    <a id="L13"></a>&#34;path&#34;;
    <a id="L14"></a>&#34;strings&#34;;
    <a id="L15"></a>&#34;utf8&#34;;
<a id="L16"></a>)

<a id="L18"></a><span class="comment">// TODO this should be in a mime package somewhere</span>
<a id="L19"></a>var contentByExt = map[string]string{
    <a id="L20"></a>&#34;.css&#34;: &#34;text/css&#34;,
    <a id="L21"></a>&#34;.gif&#34;: &#34;image/gif&#34;,
    <a id="L22"></a>&#34;.html&#34;: &#34;text/html; charset=utf-8&#34;,
    <a id="L23"></a>&#34;.jpg&#34;: &#34;image/jpeg&#34;,
    <a id="L24"></a>&#34;.js&#34;: &#34;application/x-javascript&#34;,
    <a id="L25"></a>&#34;.pdf&#34;: &#34;application/pdf&#34;,
    <a id="L26"></a>&#34;.png&#34;: &#34;image/png&#34;,
<a id="L27"></a>}

<a id="L29"></a><span class="comment">// Heuristic: b is text if it is valid UTF-8 and doesn&#39;t</span>
<a id="L30"></a><span class="comment">// contain any unprintable ASCII or Unicode characters.</span>
<a id="L31"></a>func isText(b []byte) bool {
    <a id="L32"></a>for len(b) &gt; 0 &amp;&amp; utf8.FullRune(b) {
        <a id="L33"></a>rune, size := utf8.DecodeRune(b);
        <a id="L34"></a>if size == 1 &amp;&amp; rune == utf8.RuneError {
            <a id="L35"></a><span class="comment">// decoding error</span>
            <a id="L36"></a>return false
        <a id="L37"></a>}
        <a id="L38"></a>if 0x80 &lt;= rune &amp;&amp; rune &lt;= 0x9F {
            <a id="L39"></a>return false
        <a id="L40"></a>}
        <a id="L41"></a>if rune &lt; &#39; &#39; {
            <a id="L42"></a>switch rune {
            <a id="L43"></a>case &#39;\n&#39;, &#39;\r&#39;, &#39;\t&#39;:
                <a id="L44"></a><span class="comment">// okay</span>
            <a id="L45"></a>default:
                <a id="L46"></a><span class="comment">// binary garbage</span>
                <a id="L47"></a>return false
            <a id="L48"></a>}
        <a id="L49"></a>}
        <a id="L50"></a>b = b[size:len(b)];
    <a id="L51"></a>}
    <a id="L52"></a>return true;
<a id="L53"></a>}

<a id="L55"></a>func dirList(c *Conn, f *os.File) {
    <a id="L56"></a>fmt.Fprintf(c, &#34;&lt;pre&gt;\n&#34;);
    <a id="L57"></a>for {
        <a id="L58"></a>dirs, err := f.Readdir(100);
        <a id="L59"></a>if err != nil || len(dirs) == 0 {
            <a id="L60"></a>break
        <a id="L61"></a>}
        <a id="L62"></a>for _, d := range dirs {
            <a id="L63"></a>name := d.Name;
            <a id="L64"></a>if d.IsDirectory() {
                <a id="L65"></a>name += &#34;/&#34;
            <a id="L66"></a>}
            <a id="L67"></a><span class="comment">// TODO htmlescape</span>
            <a id="L68"></a>fmt.Fprintf(c, &#34;&lt;a href=\&#34;%s\&#34;&gt;%s&lt;/a&gt;\n&#34;, name, name);
        <a id="L69"></a>}
    <a id="L70"></a>}
    <a id="L71"></a>fmt.Fprintf(c, &#34;&lt;/pre&gt;\n&#34;);
<a id="L72"></a>}


<a id="L75"></a>func serveFileInternal(c *Conn, r *Request, name string, redirect bool) {
    <a id="L76"></a>const indexPage = &#34;/index.html&#34;;

    <a id="L78"></a><span class="comment">// redirect to strip off any index.html</span>
    <a id="L79"></a>n := len(name) - len(indexPage);
    <a id="L80"></a>if n &gt;= 0 &amp;&amp; name[n:len(name)] == indexPage {
        <a id="L81"></a>Redirect(c, name[0:n+1], StatusMovedPermanently);
        <a id="L82"></a>return;
    <a id="L83"></a>}

    <a id="L85"></a>f, err := os.Open(name, os.O_RDONLY, 0);
    <a id="L86"></a>if err != nil {
        <a id="L87"></a><span class="comment">// TODO expose actual error?</span>
        <a id="L88"></a>NotFound(c, r);
        <a id="L89"></a>return;
    <a id="L90"></a>}
    <a id="L91"></a>defer f.Close();

    <a id="L93"></a>d, err1 := f.Stat();
    <a id="L94"></a>if err1 != nil {
        <a id="L95"></a><span class="comment">// TODO expose actual error?</span>
        <a id="L96"></a>NotFound(c, r);
        <a id="L97"></a>return;
    <a id="L98"></a>}

    <a id="L100"></a>if redirect {
        <a id="L101"></a><span class="comment">// redirect to canonical path: / at end of directory url</span>
        <a id="L102"></a><span class="comment">// r.URL.Path always begins with /</span>
        <a id="L103"></a>url := r.URL.Path;
        <a id="L104"></a>if d.IsDirectory() {
            <a id="L105"></a>if url[len(url)-1] != &#39;/&#39; {
                <a id="L106"></a>Redirect(c, url+&#34;/&#34;, StatusMovedPermanently);
                <a id="L107"></a>return;
            <a id="L108"></a>}
        <a id="L109"></a>} else {
            <a id="L110"></a>if url[len(url)-1] == &#39;/&#39; {
                <a id="L111"></a>Redirect(c, url[0:len(url)-1], StatusMovedPermanently);
                <a id="L112"></a>return;
            <a id="L113"></a>}
        <a id="L114"></a>}
    <a id="L115"></a>}

    <a id="L117"></a><span class="comment">// use contents of index.html for directory, if present</span>
    <a id="L118"></a>if d.IsDirectory() {
        <a id="L119"></a>index := name + indexPage;
        <a id="L120"></a>ff, err := os.Open(index, os.O_RDONLY, 0);
        <a id="L121"></a>if err == nil {
            <a id="L122"></a>defer ff.Close();
            <a id="L123"></a>dd, err := ff.Stat();
            <a id="L124"></a>if err == nil {
                <a id="L125"></a>name = index;
                <a id="L126"></a>d = dd;
                <a id="L127"></a>f = ff;
            <a id="L128"></a>}
        <a id="L129"></a>}
    <a id="L130"></a>}

    <a id="L132"></a>if d.IsDirectory() {
        <a id="L133"></a>dirList(c, f);
        <a id="L134"></a>return;
    <a id="L135"></a>}

    <a id="L137"></a><span class="comment">// serve file</span>
    <a id="L138"></a><span class="comment">// use extension to find content type.</span>
    <a id="L139"></a>ext := path.Ext(name);
    <a id="L140"></a>if ctype, ok := contentByExt[ext]; ok {
        <a id="L141"></a>c.SetHeader(&#34;Content-Type&#34;, ctype)
    <a id="L142"></a>} else {
        <a id="L143"></a><span class="comment">// read first chunk to decide between utf-8 text and binary</span>
        <a id="L144"></a>var buf [1024]byte;
        <a id="L145"></a>n, _ := io.ReadFull(f, &amp;buf);
        <a id="L146"></a>b := buf[0:n];
        <a id="L147"></a>if isText(b) {
            <a id="L148"></a>c.SetHeader(&#34;Content-Type&#34;, &#34;text-plain; charset=utf-8&#34;)
        <a id="L149"></a>} else {
            <a id="L150"></a>c.SetHeader(&#34;Content-Type&#34;, &#34;application/octet-stream&#34;) <span class="comment">// generic binary</span>
        <a id="L151"></a>}
        <a id="L152"></a>c.Write(b);
    <a id="L153"></a>}
    <a id="L154"></a>io.Copy(c, f);
<a id="L155"></a>}

<a id="L157"></a><span class="comment">// ServeFile replies to the request with the contents of the named file or directory.</span>
<a id="L158"></a>func ServeFile(c *Conn, r *Request, name string) {
    <a id="L159"></a>serveFileInternal(c, r, name, false)
<a id="L160"></a>}

<a id="L162"></a>type fileHandler struct {
    <a id="L163"></a>root   string;
    <a id="L164"></a>prefix string;
<a id="L165"></a>}

<a id="L167"></a><span class="comment">// FileServer returns a handler that serves HTTP requests</span>
<a id="L168"></a><span class="comment">// with the contents of the file system rooted at root.</span>
<a id="L169"></a><span class="comment">// It strips prefix from the incoming requests before</span>
<a id="L170"></a><span class="comment">// looking up the file name in the file system.</span>
<a id="L171"></a>func FileServer(root, prefix string) Handler { return &amp;fileHandler{root, prefix} }

<a id="L173"></a>func (f *fileHandler) ServeHTTP(c *Conn, r *Request) {
    <a id="L174"></a>path := r.URL.Path;
    <a id="L175"></a>if !strings.HasPrefix(path, f.prefix) {
        <a id="L176"></a>NotFound(c, r);
        <a id="L177"></a>return;
    <a id="L178"></a>}
    <a id="L179"></a>path = path[len(f.prefix):len(path)];
    <a id="L180"></a>serveFileInternal(c, r, f.root+&#34;/&#34;+path, true);
<a id="L181"></a>}
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
