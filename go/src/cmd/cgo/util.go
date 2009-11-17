<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/cgo/util.go</title>

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
  <h1 id="generatedHeader">Source file /src/cmd/cgo/util.go</h1>

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
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;exec&#34;;
    <a id="L10"></a>&#34;fmt&#34;;
    <a id="L11"></a>&#34;go/token&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;os&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// A ByteReaderAt implements io.ReadAt using a slice of bytes.</span>
<a id="L17"></a>type ByteReaderAt []byte

<a id="L19"></a>func (r ByteReaderAt) ReadAt(p []byte, off int64) (n int, err os.Error) {
    <a id="L20"></a>if off &gt;= int64(len(r)) || off &lt; 0 {
        <a id="L21"></a>return 0, os.EOF
    <a id="L22"></a>}
    <a id="L23"></a>return bytes.Copy(p, r[off:len(r)]), nil;
<a id="L24"></a>}

<a id="L26"></a><span class="comment">// run runs the command argv, feeding in stdin on standard input.</span>
<a id="L27"></a><span class="comment">// It returns the output to standard output and standard error.</span>
<a id="L28"></a><span class="comment">// ok indicates whether the command exited successfully.</span>
<a id="L29"></a>func run(stdin []byte, argv []string) (stdout, stderr []byte, ok bool) {
    <a id="L30"></a>cmd, err := exec.LookPath(argv[0]);
    <a id="L31"></a>if err != nil {
        <a id="L32"></a>fatal(&#34;exec %s: %s&#34;, argv[0], err)
    <a id="L33"></a>}
    <a id="L34"></a>r0, w0, err := os.Pipe();
    <a id="L35"></a>if err != nil {
        <a id="L36"></a>fatal(&#34;%s&#34;, err)
    <a id="L37"></a>}
    <a id="L38"></a>r1, w1, err := os.Pipe();
    <a id="L39"></a>if err != nil {
        <a id="L40"></a>fatal(&#34;%s&#34;, err)
    <a id="L41"></a>}
    <a id="L42"></a>r2, w2, err := os.Pipe();
    <a id="L43"></a>if err != nil {
        <a id="L44"></a>fatal(&#34;%s&#34;, err)
    <a id="L45"></a>}
    <a id="L46"></a>pid, err := os.ForkExec(cmd, argv, os.Environ(), &#34;&#34;, []*os.File{r0, w1, w2});
    <a id="L47"></a>if err != nil {
        <a id="L48"></a>fatal(&#34;%s&#34;, err)
    <a id="L49"></a>}
    <a id="L50"></a>r0.Close();
    <a id="L51"></a>w1.Close();
    <a id="L52"></a>w2.Close();
    <a id="L53"></a>c := make(chan bool);
    <a id="L54"></a>go func() {
        <a id="L55"></a>w0.Write(stdin);
        <a id="L56"></a>w0.Close();
        <a id="L57"></a>c &lt;- true;
    <a id="L58"></a>}();
    <a id="L59"></a>var xstdout []byte; <span class="comment">// TODO(rsc): delete after 6g can take address of out parameter</span>
    <a id="L60"></a>go func() {
        <a id="L61"></a>xstdout, _ = io.ReadAll(r1);
        <a id="L62"></a>r1.Close();
        <a id="L63"></a>c &lt;- true;
    <a id="L64"></a>}();
    <a id="L65"></a>stderr, _ = io.ReadAll(r2);
    <a id="L66"></a>r2.Close();
    <a id="L67"></a>&lt;-c;
    <a id="L68"></a>&lt;-c;
    <a id="L69"></a>stdout = xstdout;

    <a id="L71"></a>w, err := os.Wait(pid, 0);
    <a id="L72"></a>if err != nil {
        <a id="L73"></a>fatal(&#34;%s&#34;, err)
    <a id="L74"></a>}
    <a id="L75"></a>ok = w.Exited() &amp;&amp; w.ExitStatus() == 0;
    <a id="L76"></a>return;
<a id="L77"></a>}

<a id="L79"></a><span class="comment">// Die with an error message.</span>
<a id="L80"></a>func fatal(msg string, args ...) {
    <a id="L81"></a>fmt.Fprintf(os.Stderr, msg+&#34;\n&#34;, args);
    <a id="L82"></a>os.Exit(2);
<a id="L83"></a>}

<a id="L85"></a>var nerrors int
<a id="L86"></a>var noPos token.Position

<a id="L88"></a>func error(pos token.Position, msg string, args ...) {
    <a id="L89"></a>nerrors++;
    <a id="L90"></a>if pos.IsValid() {
        <a id="L91"></a>fmt.Fprintf(os.Stderr, &#34;%s: &#34;, pos)
    <a id="L92"></a>}
    <a id="L93"></a>fmt.Fprintf(os.Stderr, msg, args);
    <a id="L94"></a>fmt.Fprintf(os.Stderr, &#34;\n&#34;);
<a id="L95"></a>}
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
