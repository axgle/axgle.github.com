<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/ebnflint/ebnflint.go</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/cmd/ebnflint/ebnflint.go</h1>

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
    <a id="L9"></a>&#34;ebnf&#34;;
    <a id="L10"></a>&#34;flag&#34;;
    <a id="L11"></a>&#34;fmt&#34;;
    <a id="L12"></a>&#34;go/scanner&#34;;
    <a id="L13"></a>&#34;io&#34;;
    <a id="L14"></a>&#34;os&#34;;
    <a id="L15"></a>&#34;path&#34;;
    <a id="L16"></a>&#34;strings&#34;;
<a id="L17"></a>)


<a id="L20"></a>var start = flag.String(&#34;start&#34;, &#34;Start&#34;, &#34;name of start production&#34;)


<a id="L23"></a>func usage() {
    <a id="L24"></a>fmt.Fprintf(os.Stderr, &#34;usage: ebnflint [flags] [filename]\n&#34;);
    <a id="L25"></a>flag.PrintDefaults();
    <a id="L26"></a>os.Exit(1);
<a id="L27"></a>}


<a id="L30"></a><span class="comment">// Markers around EBNF sections in .html files</span>
<a id="L31"></a>var (
    <a id="L32"></a>open  = strings.Bytes(`&lt;pre class=&#34;ebnf&#34;&gt;`);
    <a id="L33"></a>close = strings.Bytes(`&lt;/pre&gt;`);
<a id="L34"></a>)


<a id="L37"></a>func extractEBNF(src []byte) []byte {
    <a id="L38"></a>var buf bytes.Buffer;

    <a id="L40"></a>for {
        <a id="L41"></a><span class="comment">// i = beginning of EBNF text</span>
        <a id="L42"></a>i := bytes.Index(src, open);
        <a id="L43"></a>if i &lt; 0 {
            <a id="L44"></a>break <span class="comment">// no EBNF found - we are done</span>
        <a id="L45"></a>}
        <a id="L46"></a>i += len(open);

        <a id="L48"></a><span class="comment">// write as many newlines as found in the excluded text</span>
        <a id="L49"></a><span class="comment">// to maintain correct line numbers in error messages</span>
        <a id="L50"></a>for _, ch := range src[0:i] {
            <a id="L51"></a>if ch == &#39;\n&#39; {
                <a id="L52"></a>buf.WriteByte(&#39;\n&#39;)
            <a id="L53"></a>}
        <a id="L54"></a>}

        <a id="L56"></a><span class="comment">// j = end of EBNF text (or end of source)</span>
        <a id="L57"></a>j := bytes.Index(src[i:len(src)], close); <span class="comment">// close marker</span>
        <a id="L58"></a>if j &lt; 0 {
            <a id="L59"></a>j = len(src) - i
        <a id="L60"></a>}
        <a id="L61"></a>j += i;

        <a id="L63"></a><span class="comment">// copy EBNF text</span>
        <a id="L64"></a>buf.Write(src[i:j]);

        <a id="L66"></a><span class="comment">// advance</span>
        <a id="L67"></a>src = src[j:len(src)];
    <a id="L68"></a>}

    <a id="L70"></a>return buf.Bytes();
<a id="L71"></a>}


<a id="L74"></a>func main() {
    <a id="L75"></a>flag.Parse();

    <a id="L77"></a>var filename string;
    <a id="L78"></a>switch flag.NArg() {
    <a id="L79"></a>case 0:
        <a id="L80"></a>filename = &#34;/dev/stdin&#34;
    <a id="L81"></a>case 1:
        <a id="L82"></a>filename = flag.Arg(0)
    <a id="L83"></a>default:
        <a id="L84"></a>usage()
    <a id="L85"></a>}

    <a id="L87"></a>src, err := io.ReadFile(filename);
    <a id="L88"></a>if err != nil {
        <a id="L89"></a>scanner.PrintError(os.Stderr, err)
    <a id="L90"></a>}

    <a id="L92"></a>if path.Ext(filename) == &#34;.html&#34; {
        <a id="L93"></a>src = extractEBNF(src)
    <a id="L94"></a>}

    <a id="L96"></a>grammar, err := ebnf.Parse(filename, src);
    <a id="L97"></a>if err != nil {
        <a id="L98"></a>scanner.PrintError(os.Stderr, err)
    <a id="L99"></a>}

    <a id="L101"></a>if err = ebnf.Verify(grammar, *start); err != nil {
        <a id="L102"></a>scanner.PrintError(os.Stderr, err)
    <a id="L103"></a>}
<a id="L104"></a>}
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
