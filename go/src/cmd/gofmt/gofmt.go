<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/gofmt/gofmt.go</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/gofmt/gofmt.go</h1>

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
    <a id="L11"></a>&#34;go/parser&#34;;
    <a id="L12"></a>&#34;go/printer&#34;;
    <a id="L13"></a>&#34;go/scanner&#34;;
    <a id="L14"></a>&#34;io&#34;;
    <a id="L15"></a>&#34;os&#34;;
    <a id="L16"></a>pathutil &#34;path&#34;;
    <a id="L17"></a>&#34;strings&#34;;
<a id="L18"></a>)


<a id="L21"></a>var (
    <a id="L22"></a><span class="comment">// main operation modes</span>
    <a id="L23"></a>list  = flag.Bool(&#34;l&#34;, false, &#34;list files whose formatting differs from gofmt&#39;s&#34;);
    <a id="L24"></a>write = flag.Bool(&#34;w&#34;, false, &#34;write result to (source) file instead of stdout&#34;);

    <a id="L26"></a><span class="comment">// debugging support</span>
    <a id="L27"></a>comments = flag.Bool(&#34;comments&#34;, true, &#34;print comments&#34;);
    <a id="L28"></a>trace    = flag.Bool(&#34;trace&#34;, false, &#34;print parse trace&#34;);

    <a id="L30"></a><span class="comment">// layout control</span>
    <a id="L31"></a>align     = flag.Bool(&#34;align&#34;, true, &#34;align columns&#34;);
    <a id="L32"></a>tabwidth  = flag.Int(&#34;tabwidth&#34;, 8, &#34;tab width&#34;);
    <a id="L33"></a>usespaces = flag.Bool(&#34;spaces&#34;, false, &#34;align with spaces instead of tabs&#34;);
<a id="L34"></a>)


<a id="L37"></a>var exitCode = 0

<a id="L39"></a>func report(err os.Error) {
    <a id="L40"></a>scanner.PrintError(os.Stderr, err);
    <a id="L41"></a>exitCode = 2;
<a id="L42"></a>}


<a id="L45"></a>func usage() {
    <a id="L46"></a>fmt.Fprintf(os.Stderr, &#34;usage: gofmt [flags] [path ...]\n&#34;);
    <a id="L47"></a>flag.PrintDefaults();
    <a id="L48"></a>os.Exit(2);
<a id="L49"></a>}


<a id="L52"></a>func parserMode() uint {
    <a id="L53"></a>mode := uint(0);
    <a id="L54"></a>if *comments {
        <a id="L55"></a>mode |= parser.ParseComments
    <a id="L56"></a>}
    <a id="L57"></a>if *trace {
        <a id="L58"></a>mode |= parser.Trace
    <a id="L59"></a>}
    <a id="L60"></a>return mode;
<a id="L61"></a>}


<a id="L64"></a>func printerMode() uint {
    <a id="L65"></a>mode := uint(0);
    <a id="L66"></a>if !*align {
        <a id="L67"></a>mode |= printer.RawFormat
    <a id="L68"></a>}
    <a id="L69"></a>if *usespaces {
        <a id="L70"></a>mode |= printer.UseSpaces
    <a id="L71"></a>}
    <a id="L72"></a>return mode;
<a id="L73"></a>}


<a id="L76"></a>func isGoFile(d *os.Dir) bool {
    <a id="L77"></a><span class="comment">// ignore non-Go files</span>
    <a id="L78"></a>return d.IsRegular() &amp;&amp; !strings.HasPrefix(d.Name, &#34;.&#34;) &amp;&amp; strings.HasSuffix(d.Name, &#34;.go&#34;)
<a id="L79"></a>}


<a id="L82"></a>func processFile(filename string) os.Error {
    <a id="L83"></a>src, err := io.ReadFile(filename);
    <a id="L84"></a>if err != nil {
        <a id="L85"></a>return err
    <a id="L86"></a>}

    <a id="L88"></a>file, err := parser.ParseFile(filename, src, parserMode());
    <a id="L89"></a>if err != nil {
        <a id="L90"></a>return err
    <a id="L91"></a>}

    <a id="L93"></a>var res bytes.Buffer;
    <a id="L94"></a>_, err = (&amp;printer.Config{printerMode(), *tabwidth, nil}).Fprint(&amp;res, file);
    <a id="L95"></a>if err != nil {
        <a id="L96"></a>return err
    <a id="L97"></a>}

    <a id="L99"></a>if bytes.Compare(src, res.Bytes()) != 0 {
        <a id="L100"></a><span class="comment">// formatting has changed</span>
        <a id="L101"></a>if *list {
            <a id="L102"></a>fmt.Fprintln(os.Stdout, filename)
        <a id="L103"></a>}
        <a id="L104"></a>if *write {
            <a id="L105"></a>err = io.WriteFile(filename, res.Bytes(), 0);
            <a id="L106"></a>if err != nil {
                <a id="L107"></a>return err
            <a id="L108"></a>}
        <a id="L109"></a>}
    <a id="L110"></a>}

    <a id="L112"></a>if !*list &amp;&amp; !*write {
        <a id="L113"></a>_, err = os.Stdout.Write(res.Bytes())
    <a id="L114"></a>}

    <a id="L116"></a>return err;
<a id="L117"></a>}


<a id="L120"></a>type fileVisitor chan os.Error

<a id="L122"></a>func (v fileVisitor) VisitDir(path string, d *os.Dir) bool {
    <a id="L123"></a>return true
<a id="L124"></a>}


<a id="L127"></a>func (v fileVisitor) VisitFile(path string, d *os.Dir) {
    <a id="L128"></a>if isGoFile(d) {
        <a id="L129"></a>v &lt;- nil; <span class="comment">// synchronize error handler</span>
        <a id="L130"></a>if err := processFile(path); err != nil {
            <a id="L131"></a>v &lt;- err
        <a id="L132"></a>}
    <a id="L133"></a>}
<a id="L134"></a>}


<a id="L137"></a>func walkDir(path string) {
    <a id="L138"></a><span class="comment">// start an error handler</span>
    <a id="L139"></a>v := make(fileVisitor);
    <a id="L140"></a>go func() {
        <a id="L141"></a>for err := range v {
            <a id="L142"></a>if err != nil {
                <a id="L143"></a>report(err)
            <a id="L144"></a>}
        <a id="L145"></a>}
    <a id="L146"></a>}();
    <a id="L147"></a><span class="comment">// walk the tree</span>
    <a id="L148"></a>pathutil.Walk(path, v, v);
    <a id="L149"></a>close(v);
<a id="L150"></a>}


<a id="L153"></a>func main() {
    <a id="L154"></a>flag.Usage = usage;
    <a id="L155"></a>flag.Parse();
    <a id="L156"></a>if *tabwidth &lt; 0 {
        <a id="L157"></a>fmt.Fprintf(os.Stderr, &#34;negative tabwidth %d\n&#34;, *tabwidth);
        <a id="L158"></a>os.Exit(2);
    <a id="L159"></a>}

    <a id="L161"></a>if flag.NArg() == 0 {
        <a id="L162"></a>if err := processFile(&#34;/dev/stdin&#34;); err != nil {
            <a id="L163"></a>report(err)
        <a id="L164"></a>}
    <a id="L165"></a>}

    <a id="L167"></a>for i := 0; i &lt; flag.NArg(); i++ {
        <a id="L168"></a>path := flag.Arg(i);
        <a id="L169"></a>switch dir, err := os.Stat(path); {
        <a id="L170"></a>case err != nil:
            <a id="L171"></a>report(err)
        <a id="L172"></a>case dir.IsRegular():
            <a id="L173"></a>if err := processFile(path); err != nil {
                <a id="L174"></a>report(err)
            <a id="L175"></a>}
        <a id="L176"></a>case dir.IsDirectory():
            <a id="L177"></a>walkDir(path)
        <a id="L178"></a>}
    <a id="L179"></a>}

    <a id="L181"></a>os.Exit(exitCode);
<a id="L182"></a>}
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
