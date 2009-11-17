<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/patch/git.go</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/patch/git.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors.  All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package patch

<a id="L7"></a>import (
    <a id="L8"></a>&#34;bytes&#34;;
    <a id="L9"></a>&#34;compress/zlib&#34;;
    <a id="L10"></a>&#34;crypto/sha1&#34;;
    <a id="L11"></a>&#34;encoding/git85&#34;;
    <a id="L12"></a>&#34;fmt&#34;;
    <a id="L13"></a>&#34;io&#34;;
    <a id="L14"></a>&#34;os&#34;;
<a id="L15"></a>)

<a id="L17"></a>func gitSHA1(data []byte) []byte {
    <a id="L18"></a>if len(data) == 0 {
        <a id="L19"></a><span class="comment">// special case: 0 length is all zeros sum</span>
        <a id="L20"></a>return make([]byte, 20)
    <a id="L21"></a>}
    <a id="L22"></a>h := sha1.New();
    <a id="L23"></a>fmt.Fprintf(h, &#34;blob %d\x00&#34;, len(data));
    <a id="L24"></a>h.Write(data);
    <a id="L25"></a>return h.Sum();
<a id="L26"></a>}

<a id="L28"></a><span class="comment">// BUG(rsc): The Git binary delta format is not implemented, only Git binary literals.</span>

<a id="L30"></a><span class="comment">// GitBinaryLiteral represents a Git binary literal diff.</span>
<a id="L31"></a>type GitBinaryLiteral struct {
    <a id="L32"></a>OldSHA1 []byte; <span class="comment">// if non-empty, the SHA1 hash of the original</span>
    <a id="L33"></a>New     []byte; <span class="comment">// the new contents</span>
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// Apply implements the Diff interface&#39;s Apply method.</span>
<a id="L37"></a>func (d *GitBinaryLiteral) Apply(old []byte) ([]byte, os.Error) {
    <a id="L38"></a>if sum := gitSHA1(old); !bytes.HasPrefix(sum, d.OldSHA1) {
        <a id="L39"></a>return nil, ErrPatchFailure
    <a id="L40"></a>}
    <a id="L41"></a>return d.New, nil;
<a id="L42"></a>}

<a id="L44"></a>func unhex(c byte) uint8 {
    <a id="L45"></a>switch {
    <a id="L46"></a>case &#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39;:
        <a id="L47"></a>return c - &#39;0&#39;
    <a id="L48"></a>case &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;f&#39;:
        <a id="L49"></a>return c - &#39;a&#39; + 10
    <a id="L50"></a>case &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;F&#39;:
        <a id="L51"></a>return c - &#39;A&#39; + 10
    <a id="L52"></a>}
    <a id="L53"></a>return 255;
<a id="L54"></a>}

<a id="L56"></a>func getHex(s []byte) (data []byte, rest []byte) {
    <a id="L57"></a>n := 0;
    <a id="L58"></a>for n &lt; len(s) &amp;&amp; unhex(s[n]) != 255 {
        <a id="L59"></a>n++
    <a id="L60"></a>}
    <a id="L61"></a>n &amp;^= 1; <span class="comment">// Only take an even number of hex digits.</span>
    <a id="L62"></a>data = make([]byte, n/2);
    <a id="L63"></a>for i := range data {
        <a id="L64"></a>data[i] = unhex(s[2*i])&lt;&lt;4 | unhex(s[2*i+1])
    <a id="L65"></a>}
    <a id="L66"></a>rest = s[n:len(s)];
    <a id="L67"></a>return;
<a id="L68"></a>}

<a id="L70"></a><span class="comment">// ParseGitBinary parses raw as a Git binary patch.</span>
<a id="L71"></a>func ParseGitBinary(raw []byte) (Diff, os.Error) {
    <a id="L72"></a>var oldSHA1, newSHA1 []byte;
    <a id="L73"></a>var sawBinary bool;

    <a id="L75"></a>for {
        <a id="L76"></a>var first []byte;
        <a id="L77"></a>first, raw, _ = getLine(raw, 1);
        <a id="L78"></a>first = bytes.TrimSpace(first);
        <a id="L79"></a>if s, ok := skip(first, &#34;index &#34;); ok {
            <a id="L80"></a>oldSHA1, s = getHex(s);
            <a id="L81"></a>if s, ok = skip(s, &#34;..&#34;); !ok {
                <a id="L82"></a>continue
            <a id="L83"></a>}
            <a id="L84"></a>newSHA1, s = getHex(s);
            <a id="L85"></a>continue;
        <a id="L86"></a>}
        <a id="L87"></a>if _, ok := skip(first, &#34;GIT binary patch&#34;); ok {
            <a id="L88"></a>sawBinary = true;
            <a id="L89"></a>continue;
        <a id="L90"></a>}
        <a id="L91"></a>if n, _, ok := atoi(first, &#34;literal &#34;, 10); ok &amp;&amp; sawBinary {
            <a id="L92"></a>data := make([]byte, n);
            <a id="L93"></a>d := git85.NewDecoder(bytes.NewBuffer(raw));
            <a id="L94"></a>z, err := zlib.NewInflater(d);
            <a id="L95"></a>if err != nil {
                <a id="L96"></a>return nil, err
            <a id="L97"></a>}
            <a id="L98"></a>defer z.Close();
            <a id="L99"></a>if _, err = io.ReadFull(z, data); err != nil {
                <a id="L100"></a>if err == os.EOF {
                    <a id="L101"></a>err = io.ErrUnexpectedEOF
                <a id="L102"></a>}
                <a id="L103"></a>return nil, err;
            <a id="L104"></a>}
            <a id="L105"></a>var buf [1]byte;
            <a id="L106"></a>m, err := z.Read(&amp;buf);
            <a id="L107"></a>if m != 0 || err != os.EOF {
                <a id="L108"></a>return nil, os.NewError(&#34;Git binary literal longer than expected&#34;)
            <a id="L109"></a>}

            <a id="L111"></a>if sum := gitSHA1(data); !bytes.HasPrefix(sum, newSHA1) {
                <a id="L112"></a>return nil, os.NewError(&#34;Git binary literal SHA1 mismatch&#34;)
            <a id="L113"></a>}
            <a id="L114"></a>return &amp;GitBinaryLiteral{oldSHA1, data}, nil;
        <a id="L115"></a>}
        <a id="L116"></a>if !sawBinary {
            <a id="L117"></a>return nil, os.NewError(&#34;unexpected Git patch header: &#34; + string(first))
        <a id="L118"></a>}
    <a id="L119"></a>}
    <a id="L120"></a>panic(&#34;unreachable&#34;);
<a id="L121"></a>}
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
