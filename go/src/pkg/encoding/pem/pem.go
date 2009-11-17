<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/encoding/pem/pem.go</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/encoding/pem/pem.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// This package implements the PEM data encoding, which originated in Privacy</span>
<a id="L6"></a><span class="comment">// Enhanced Mail. The most common use of PEM encoding today is in TLS keys and</span>
<a id="L7"></a><span class="comment">// certificates. See RFC 1421.</span>
<a id="L8"></a>package pem

<a id="L10"></a>import (
    <a id="L11"></a>&#34;bytes&#34;;
    <a id="L12"></a>&#34;encoding/base64&#34;;
    <a id="L13"></a>&#34;strings&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// A Block represents a PEM encoded structure.</span>
<a id="L17"></a><span class="comment">//</span>
<a id="L18"></a><span class="comment">// The encoded form is:</span>
<a id="L19"></a><span class="comment">//    -----BEGIN Type-----</span>
<a id="L20"></a><span class="comment">//    Headers</span>
<a id="L21"></a><span class="comment">//    base64-encoded Bytes</span>
<a id="L22"></a><span class="comment">//    -----END Type-----</span>
<a id="L23"></a><span class="comment">// where Headers is a possibly empty sequence of Key: Value lines.</span>
<a id="L24"></a>type Block struct {
    <a id="L25"></a>Type    string;            <span class="comment">// The type, taken from the preamble (i.e. &#34;RSA PRIVATE KEY&#34;).</span>
    <a id="L26"></a>Headers map[string]string; <span class="comment">// Optional headers.</span>
    <a id="L27"></a>Bytes   []byte;            <span class="comment">// The decoded bytes of the contents. Typically a DER encoded ASN.1 structure.</span>
<a id="L28"></a>}

<a id="L30"></a><span class="comment">// getLine results the first \r\n or \n delineated line from the given byte</span>
<a id="L31"></a><span class="comment">// array. The line does not include the \r\n or \n. The remainder of the byte</span>
<a id="L32"></a><span class="comment">// array (also not including the new line bytes) is also returned and this will</span>
<a id="L33"></a><span class="comment">// always be smaller than the original argument.</span>
<a id="L34"></a>func getLine(data []byte) (line, rest []byte) {
    <a id="L35"></a>i := bytes.Index(data, []byte{&#39;\n&#39;});
    <a id="L36"></a>var j int;
    <a id="L37"></a>if i &lt; 0 {
        <a id="L38"></a>i = len(data);
        <a id="L39"></a>j = i;
    <a id="L40"></a>} else {
        <a id="L41"></a>j = i + 1;
        <a id="L42"></a>if i &gt; 0 &amp;&amp; data[i-1] == &#39;\r&#39; {
            <a id="L43"></a>i--
        <a id="L44"></a>}
    <a id="L45"></a>}
    <a id="L46"></a>return data[0:i], data[j:len(data)];
<a id="L47"></a>}

<a id="L49"></a><span class="comment">// removeWhitespace returns a copy of its input with all spaces, tab and</span>
<a id="L50"></a><span class="comment">// newline characters removed.</span>
<a id="L51"></a>func removeWhitespace(data []byte) []byte {
    <a id="L52"></a>result := make([]byte, len(data));
    <a id="L53"></a>n := 0;

    <a id="L55"></a>for _, b := range data {
        <a id="L56"></a>if b == &#39; &#39; || b == &#39;\t&#39; || b == &#39;\r&#39; || b == &#39;\n&#39; {
            <a id="L57"></a>continue
        <a id="L58"></a>}
        <a id="L59"></a>result[n] = b;
        <a id="L60"></a>n++;
    <a id="L61"></a>}

    <a id="L63"></a>return result[0:n];
<a id="L64"></a>}

<a id="L66"></a>var pemStart = strings.Bytes(&#34;\n-----BEGIN &#34;)
<a id="L67"></a>var pemEnd = strings.Bytes(&#34;\n-----END &#34;)
<a id="L68"></a>var pemEndOfLine = strings.Bytes(&#34;-----&#34;)

<a id="L70"></a><span class="comment">// Decode will find the next PEM formatted block (certificate, private key</span>
<a id="L71"></a><span class="comment">// etc) in the input. It returns that block and the remainder of the input. If</span>
<a id="L72"></a><span class="comment">// no PEM data is found, p is nil and the whole of the input is returned in</span>
<a id="L73"></a><span class="comment">// rest.</span>
<a id="L74"></a>func Decode(data []byte) (p *Block, rest []byte) {
    <a id="L75"></a><span class="comment">// pemStart begins with a newline. However, at the very beginning of</span>
    <a id="L76"></a><span class="comment">// the byte array, we&#39;ll accept the start string without it.</span>
    <a id="L77"></a>rest = data;
    <a id="L78"></a>if bytes.HasPrefix(data, pemStart[1:len(pemStart)]) {
        <a id="L79"></a>rest = rest[len(pemStart)-1 : len(data)]
    <a id="L80"></a>} else if i := bytes.Index(data, pemStart); i &gt;= 0 {
        <a id="L81"></a>rest = rest[i+len(pemStart) : len(data)]
    <a id="L82"></a>} else {
        <a id="L83"></a>return nil, data
    <a id="L84"></a>}

    <a id="L86"></a>typeLine, rest := getLine(rest);
    <a id="L87"></a>if !bytes.HasSuffix(typeLine, pemEndOfLine) {
        <a id="L88"></a>goto Error
    <a id="L89"></a>}
    <a id="L90"></a>typeLine = typeLine[0 : len(typeLine)-len(pemEndOfLine)];

    <a id="L92"></a>p = &amp;Block{
        <a id="L93"></a>Headers: make(map[string]string),
        <a id="L94"></a>Type: string(typeLine),
    <a id="L95"></a>};

    <a id="L97"></a>for {
        <a id="L98"></a><span class="comment">// This loop terminates because getLine&#39;s second result is</span>
        <a id="L99"></a><span class="comment">// always smaller than it&#39;s argument.</span>
        <a id="L100"></a>if len(rest) == 0 {
            <a id="L101"></a>return nil, data
        <a id="L102"></a>}
        <a id="L103"></a>line, next := getLine(rest);

        <a id="L105"></a>i := bytes.Index(line, []byte{&#39;:&#39;});
        <a id="L106"></a>if i == -1 {
            <a id="L107"></a>break
        <a id="L108"></a>}

        <a id="L110"></a><span class="comment">// TODO(agl): need to cope with values that spread across lines.</span>
        <a id="L111"></a>key, val := line[0:i], line[i+1:len(line)];
        <a id="L112"></a>key = bytes.TrimSpace(key);
        <a id="L113"></a>val = bytes.TrimSpace(val);
        <a id="L114"></a>p.Headers[string(key)] = string(val);
        <a id="L115"></a>rest = next;
    <a id="L116"></a>}

    <a id="L118"></a>i := bytes.Index(rest, pemEnd);
    <a id="L119"></a>if i &lt; 0 {
        <a id="L120"></a>goto Error
    <a id="L121"></a>}
    <a id="L122"></a>base64Data := removeWhitespace(rest[0:i]);

    <a id="L124"></a>p.Bytes = make([]byte, base64.StdEncoding.DecodedLen(len(base64Data)));
    <a id="L125"></a>n, err := base64.StdEncoding.Decode(p.Bytes, base64Data);
    <a id="L126"></a>if err != nil {
        <a id="L127"></a>goto Error
    <a id="L128"></a>}
    <a id="L129"></a>p.Bytes = p.Bytes[0:n];

    <a id="L131"></a>_, rest = getLine(rest[i+len(pemEnd) : len(rest)]);

    <a id="L133"></a>return;

<a id="L135"></a>Error:
    <a id="L136"></a><span class="comment">// If we get here then we have rejected a likely looking, but</span>
    <a id="L137"></a><span class="comment">// ultimately invalid PEM block. We need to start over from a new</span>
    <a id="L138"></a><span class="comment">// position.  We have consumed the preamble line and will have consumed</span>
    <a id="L139"></a><span class="comment">// any lines which could be header lines. However, a valid preamble</span>
    <a id="L140"></a><span class="comment">// line is not a valid header line, therefore we cannot have consumed</span>
    <a id="L141"></a><span class="comment">// the preamble line for the any subsequent block. Thus, we will always</span>
    <a id="L142"></a><span class="comment">// find any valid block, no matter what bytes preceed it.</span>
    <a id="L143"></a><span class="comment">//</span>
    <a id="L144"></a><span class="comment">// For example, if the input is</span>
    <a id="L145"></a><span class="comment">//</span>
    <a id="L146"></a><span class="comment">//    -----BEGIN MALFORMED BLOCK-----</span>
    <a id="L147"></a><span class="comment">//    junk that may look like header lines</span>
    <a id="L148"></a><span class="comment">//   or data lines, but no END line</span>
    <a id="L149"></a><span class="comment">//</span>
    <a id="L150"></a><span class="comment">//    -----BEGIN ACTUAL BLOCK-----</span>
    <a id="L151"></a><span class="comment">//    realdata</span>
    <a id="L152"></a><span class="comment">//    -----END ACTUAL BLOCK-----</span>
    <a id="L153"></a><span class="comment">//</span>
    <a id="L154"></a><span class="comment">// we&#39;ve failed to parse using the first BEGIN line</span>
    <a id="L155"></a><span class="comment">// and now will try again, using the second BEGIN line.</span>
    <a id="L156"></a>p, rest = Decode(rest);
    <a id="L157"></a>if p == nil {
        <a id="L158"></a>rest = data
    <a id="L159"></a>}
    <a id="L160"></a>return;
<a id="L161"></a>}
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
