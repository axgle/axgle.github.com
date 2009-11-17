<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/request.go</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/http/request.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// HTTP Request reading and parsing.</span>

<a id="L7"></a><span class="comment">// The http package implements parsing of HTTP requests, replies,</span>
<a id="L8"></a><span class="comment">// and URLs and provides an extensible HTTP server and a basic</span>
<a id="L9"></a><span class="comment">// HTTP client.</span>
<a id="L10"></a>package http

<a id="L12"></a>import (
    <a id="L13"></a>&#34;bufio&#34;;
    <a id="L14"></a>&#34;bytes&#34;;
    <a id="L15"></a>&#34;container/vector&#34;;
    <a id="L16"></a>&#34;fmt&#34;;
    <a id="L17"></a>&#34;io&#34;;
    <a id="L18"></a>&#34;os&#34;;
    <a id="L19"></a>&#34;strconv&#34;;
    <a id="L20"></a>&#34;strings&#34;;
<a id="L21"></a>)

<a id="L23"></a>const (
    <a id="L24"></a>maxLineLength  = 1024; <span class="comment">// assumed &lt; bufio.DefaultBufSize</span>
    <a id="L25"></a>maxValueLength = 1024;
    <a id="L26"></a>maxHeaderLines = 1024;
    <a id="L27"></a>chunkSize      = 4 &lt;&lt; 10; <span class="comment">// 4 KB chunks</span>
<a id="L28"></a>)

<a id="L30"></a><span class="comment">// HTTP request parsing errors.</span>
<a id="L31"></a>type ProtocolError struct {
    <a id="L32"></a>os.ErrorString;
<a id="L33"></a>}

<a id="L35"></a>var (
    <a id="L36"></a>ErrLineTooLong   = &amp;ProtocolError{&#34;header line too long&#34;};
    <a id="L37"></a>ErrHeaderTooLong = &amp;ProtocolError{&#34;header too long&#34;};
    <a id="L38"></a>ErrShortBody     = &amp;ProtocolError{&#34;entity body too short&#34;};
<a id="L39"></a>)

<a id="L41"></a>type badStringError struct {
    <a id="L42"></a>what string;
    <a id="L43"></a>str  string;
<a id="L44"></a>}

<a id="L46"></a>func (e *badStringError) String() string { return fmt.Sprintf(&#34;%s %q&#34;, e.what, e.str) }

<a id="L48"></a><span class="comment">// A Request represents a parsed HTTP request header.</span>
<a id="L49"></a>type Request struct {
    <a id="L50"></a>Method     string; <span class="comment">// GET, POST, PUT, etc.</span>
    <a id="L51"></a>RawURL     string; <span class="comment">// The raw URL given in the request.</span>
    <a id="L52"></a>URL        *URL;   <span class="comment">// Parsed URL.</span>
    <a id="L53"></a>Proto      string; <span class="comment">// &#34;HTTP/1.0&#34;</span>
    <a id="L54"></a>ProtoMajor int;    <span class="comment">// 1</span>
    <a id="L55"></a>ProtoMinor int;    <span class="comment">// 0</span>

    <a id="L57"></a><span class="comment">// A header mapping request lines to their values.</span>
    <a id="L58"></a><span class="comment">// If the header says</span>
    <a id="L59"></a><span class="comment">//</span>
    <a id="L60"></a><span class="comment">//	Accept-Language: en-us</span>
    <a id="L61"></a><span class="comment">//	accept-encoding: gzip, deflate</span>
    <a id="L62"></a><span class="comment">//	Connection: keep-alive</span>
    <a id="L63"></a><span class="comment">//</span>
    <a id="L64"></a><span class="comment">// then</span>
    <a id="L65"></a><span class="comment">//</span>
    <a id="L66"></a><span class="comment">//	Header = map[string]string{</span>
    <a id="L67"></a><span class="comment">//		&#34;Accept-Encoding&#34;: &#34;en-us&#34;,</span>
    <a id="L68"></a><span class="comment">//		&#34;Accept-Language&#34;: &#34;gzip, deflate&#34;,</span>
    <a id="L69"></a><span class="comment">//		&#34;Connection&#34;: &#34;keep-alive&#34;</span>
    <a id="L70"></a><span class="comment">//	}</span>
    <a id="L71"></a><span class="comment">//</span>
    <a id="L72"></a><span class="comment">// HTTP defines that header names are case-insensitive.</span>
    <a id="L73"></a><span class="comment">// The request parser implements this by canonicalizing the</span>
    <a id="L74"></a><span class="comment">// name, making the first character and any characters</span>
    <a id="L75"></a><span class="comment">// following a hyphen uppercase and the rest lowercase.</span>
    <a id="L76"></a>Header map[string]string;

    <a id="L78"></a><span class="comment">// The message body.</span>
    <a id="L79"></a>Body io.Reader;

    <a id="L81"></a><span class="comment">// Whether to close the connection after replying to this request.</span>
    <a id="L82"></a>Close bool;

    <a id="L84"></a><span class="comment">// The host on which the URL is sought.</span>
    <a id="L85"></a><span class="comment">// Per RFC 2616, this is either the value of the Host: header</span>
    <a id="L86"></a><span class="comment">// or the host name given in the URL itself.</span>
    <a id="L87"></a>Host string;

    <a id="L89"></a><span class="comment">// The referring URL, if sent in the request.</span>
    <a id="L90"></a><span class="comment">//</span>
    <a id="L91"></a><span class="comment">// Referer is misspelled as in the request itself,</span>
    <a id="L92"></a><span class="comment">// a mistake from the earliest days of HTTP.</span>
    <a id="L93"></a><span class="comment">// This value can also be fetched from the Header map</span>
    <a id="L94"></a><span class="comment">// as Header[&#34;Referer&#34;]; the benefit of making it</span>
    <a id="L95"></a><span class="comment">// available as a structure field is that the compiler</span>
    <a id="L96"></a><span class="comment">// can diagnose programs that use the alternate</span>
    <a id="L97"></a><span class="comment">// (correct English) spelling req.Referrer but cannot</span>
    <a id="L98"></a><span class="comment">// diagnose programs that use Header[&#34;Referrer&#34;].</span>
    <a id="L99"></a>Referer string;

    <a id="L101"></a><span class="comment">// The User-Agent: header string, if sent in the request.</span>
    <a id="L102"></a>UserAgent string;

    <a id="L104"></a><span class="comment">// The parsed form. Only available after ParseForm is called.</span>
    <a id="L105"></a>Form map[string][]string;
<a id="L106"></a>}

<a id="L108"></a><span class="comment">// ProtoAtLeast returns whether the HTTP protocol used</span>
<a id="L109"></a><span class="comment">// in the request is at least major.minor.</span>
<a id="L110"></a>func (r *Request) ProtoAtLeast(major, minor int) bool {
    <a id="L111"></a>return r.ProtoMajor &gt; major ||
        <a id="L112"></a>r.ProtoMajor == major &amp;&amp; r.ProtoMinor &gt;= minor
<a id="L113"></a>}

<a id="L115"></a><span class="comment">// Return value if nonempty, def otherwise.</span>
<a id="L116"></a>func valueOrDefault(value, def string) string {
    <a id="L117"></a>if value != &#34;&#34; {
        <a id="L118"></a>return value
    <a id="L119"></a>}
    <a id="L120"></a>return def;
<a id="L121"></a>}

<a id="L123"></a>const defaultUserAgent = &#34;Go http package&#34;

<a id="L125"></a><span class="comment">// Write writes an HTTP/1.1 request -- header and body -- in wire format.</span>
<a id="L126"></a><span class="comment">// This method consults the following fields of req:</span>
<a id="L127"></a><span class="comment">//	URL</span>
<a id="L128"></a><span class="comment">//	Method (defaults to &#34;GET&#34;)</span>
<a id="L129"></a><span class="comment">//	UserAgent (defaults to defaultUserAgent)</span>
<a id="L130"></a><span class="comment">//	Referer</span>
<a id="L131"></a><span class="comment">//	Header</span>
<a id="L132"></a><span class="comment">//	Body</span>
<a id="L133"></a><span class="comment">//</span>
<a id="L134"></a><span class="comment">// If Body is present, &#34;Transfer-Encoding: chunked&#34; is forced as a header.</span>
<a id="L135"></a>func (req *Request) Write(w io.Writer) os.Error {
    <a id="L136"></a>uri := URLEscape(req.URL.Path);
    <a id="L137"></a>if req.URL.RawQuery != &#34;&#34; {
        <a id="L138"></a>uri += &#34;?&#34; + req.URL.RawQuery
    <a id="L139"></a>}

    <a id="L141"></a>fmt.Fprintf(w, &#34;%s %s HTTP/1.1\r\n&#34;, valueOrDefault(req.Method, &#34;GET&#34;), uri);
    <a id="L142"></a>fmt.Fprintf(w, &#34;Host: %s\r\n&#34;, req.URL.Host);
    <a id="L143"></a>fmt.Fprintf(w, &#34;User-Agent: %s\r\n&#34;, valueOrDefault(req.UserAgent, defaultUserAgent));

    <a id="L145"></a>if req.Referer != &#34;&#34; {
        <a id="L146"></a>fmt.Fprintf(w, &#34;Referer: %s\r\n&#34;, req.Referer)
    <a id="L147"></a>}

    <a id="L149"></a>if req.Body != nil {
        <a id="L150"></a><span class="comment">// Force chunked encoding</span>
        <a id="L151"></a>req.Header[&#34;Transfer-Encoding&#34;] = &#34;chunked&#34;
    <a id="L152"></a>}

    <a id="L154"></a><span class="comment">// TODO: split long values?  (If so, should share code with Conn.Write)</span>
    <a id="L155"></a><span class="comment">// TODO: if Header includes values for Host, User-Agent, or Referer, this</span>
    <a id="L156"></a><span class="comment">// may conflict with the User-Agent or Referer headers we add manually.</span>
    <a id="L157"></a><span class="comment">// One solution would be to remove the Host, UserAgent, and Referer fields</span>
    <a id="L158"></a><span class="comment">// from Request, and introduce Request methods along the lines of</span>
    <a id="L159"></a><span class="comment">// Response.{GetHeader,AddHeader} and string constants for &#34;Host&#34;,</span>
    <a id="L160"></a><span class="comment">// &#34;User-Agent&#34; and &#34;Referer&#34;.</span>
    <a id="L161"></a>for k, v := range req.Header {
        <a id="L162"></a>io.WriteString(w, k+&#34;: &#34;+v+&#34;\r\n&#34;)
    <a id="L163"></a>}

    <a id="L165"></a>io.WriteString(w, &#34;\r\n&#34;);

    <a id="L167"></a>if req.Body != nil {
        <a id="L168"></a>buf := make([]byte, chunkSize);
    <a id="L169"></a>Loop:
        <a id="L170"></a>for {
            <a id="L171"></a>var nr, nw int;
            <a id="L172"></a>var er, ew os.Error;
            <a id="L173"></a>if nr, er = req.Body.Read(buf); nr &gt; 0 {
                <a id="L174"></a>if er == nil || er == os.EOF {
                    <a id="L175"></a>fmt.Fprintf(w, &#34;%x\r\n&#34;, nr);
                    <a id="L176"></a>nw, ew = w.Write(buf[0:nr]);
                    <a id="L177"></a>fmt.Fprint(w, &#34;\r\n&#34;);
                <a id="L178"></a>}
            <a id="L179"></a>}
            <a id="L180"></a>switch {
            <a id="L181"></a>case er != nil:
                <a id="L182"></a>if er == os.EOF {
                    <a id="L183"></a>break Loop
                <a id="L184"></a>}
                <a id="L185"></a>return er;
            <a id="L186"></a>case ew != nil:
                <a id="L187"></a>return ew
            <a id="L188"></a>case nw &lt; nr:
                <a id="L189"></a>return io.ErrShortWrite
            <a id="L190"></a>}
        <a id="L191"></a>}
        <a id="L192"></a><span class="comment">// last-chunk CRLF</span>
        <a id="L193"></a>fmt.Fprint(w, &#34;0\r\n\r\n&#34;);
    <a id="L194"></a>}

    <a id="L196"></a>return nil;
<a id="L197"></a>}

<a id="L199"></a><span class="comment">// Read a line of bytes (up to \n) from b.</span>
<a id="L200"></a><span class="comment">// Give up if the line exceeds maxLineLength.</span>
<a id="L201"></a><span class="comment">// The returned bytes are a pointer into storage in</span>
<a id="L202"></a><span class="comment">// the bufio, so they are only valid until the next bufio read.</span>
<a id="L203"></a>func readLineBytes(b *bufio.Reader) (p []byte, err os.Error) {
    <a id="L204"></a>if p, err = b.ReadSlice(&#39;\n&#39;); err != nil {
        <a id="L205"></a><span class="comment">// We always know when EOF is coming.</span>
        <a id="L206"></a><span class="comment">// If the caller asked for a line, there should be a line.</span>
        <a id="L207"></a>if err == os.EOF {
            <a id="L208"></a>err = io.ErrUnexpectedEOF
        <a id="L209"></a>}
        <a id="L210"></a>return nil, err;
    <a id="L211"></a>}
    <a id="L212"></a>if len(p) &gt;= maxLineLength {
        <a id="L213"></a>return nil, ErrLineTooLong
    <a id="L214"></a>}

    <a id="L216"></a><span class="comment">// Chop off trailing white space.</span>
    <a id="L217"></a>var i int;
    <a id="L218"></a>for i = len(p); i &gt; 0; i-- {
        <a id="L219"></a>if c := p[i-1]; c != &#39; &#39; &amp;&amp; c != &#39;\r&#39; &amp;&amp; c != &#39;\t&#39; &amp;&amp; c != &#39;\n&#39; {
            <a id="L220"></a>break
        <a id="L221"></a>}
    <a id="L222"></a>}
    <a id="L223"></a>return p[0:i], nil;
<a id="L224"></a>}

<a id="L226"></a><span class="comment">// readLineBytes, but convert the bytes into a string.</span>
<a id="L227"></a>func readLine(b *bufio.Reader) (s string, err os.Error) {
    <a id="L228"></a>p, e := readLineBytes(b);
    <a id="L229"></a>if e != nil {
        <a id="L230"></a>return &#34;&#34;, e
    <a id="L231"></a>}
    <a id="L232"></a>return string(p), nil;
<a id="L233"></a>}

<a id="L235"></a>var colon = []byte{&#39;:&#39;}

<a id="L237"></a><span class="comment">// Read a key/value pair from b.</span>
<a id="L238"></a><span class="comment">// A key/value has the form Key: Value\r\n</span>
<a id="L239"></a><span class="comment">// and the Value can continue on multiple lines if each continuation line</span>
<a id="L240"></a><span class="comment">// starts with a space.</span>
<a id="L241"></a>func readKeyValue(b *bufio.Reader) (key, value string, err os.Error) {
    <a id="L242"></a>line, e := readLineBytes(b);
    <a id="L243"></a>if e != nil {
        <a id="L244"></a>return &#34;&#34;, &#34;&#34;, e
    <a id="L245"></a>}
    <a id="L246"></a>if len(line) == 0 {
        <a id="L247"></a>return &#34;&#34;, &#34;&#34;, nil
    <a id="L248"></a>}

    <a id="L250"></a><span class="comment">// Scan first line for colon.</span>
    <a id="L251"></a>i := bytes.Index(line, colon);
    <a id="L252"></a>if i &lt; 0 {
        <a id="L253"></a>goto Malformed
    <a id="L254"></a>}

    <a id="L256"></a>key = string(line[0:i]);
    <a id="L257"></a>if strings.Index(key, &#34; &#34;) &gt;= 0 {
        <a id="L258"></a><span class="comment">// Key field has space - no good.</span>
        <a id="L259"></a>goto Malformed
    <a id="L260"></a>}

    <a id="L262"></a><span class="comment">// Skip initial space before value.</span>
    <a id="L263"></a>for i++; i &lt; len(line); i++ {
        <a id="L264"></a>if line[i] != &#39; &#39; {
            <a id="L265"></a>break
        <a id="L266"></a>}
    <a id="L267"></a>}
    <a id="L268"></a>value = string(line[i:len(line)]);

    <a id="L270"></a><span class="comment">// Look for extension lines, which must begin with space.</span>
    <a id="L271"></a>for {
        <a id="L272"></a>c, e := b.ReadByte();
        <a id="L273"></a>if c != &#39; &#39; {
            <a id="L274"></a>if e != os.EOF {
                <a id="L275"></a>b.UnreadByte()
            <a id="L276"></a>}
            <a id="L277"></a>break;
        <a id="L278"></a>}

        <a id="L280"></a><span class="comment">// Eat leading space.</span>
        <a id="L281"></a>for c == &#39; &#39; {
            <a id="L282"></a>if c, e = b.ReadByte(); e != nil {
                <a id="L283"></a>if e == os.EOF {
                    <a id="L284"></a>e = io.ErrUnexpectedEOF
                <a id="L285"></a>}
                <a id="L286"></a>return &#34;&#34;, &#34;&#34;, e;
            <a id="L287"></a>}
        <a id="L288"></a>}
        <a id="L289"></a>b.UnreadByte();

        <a id="L291"></a><span class="comment">// Read the rest of the line and add to value.</span>
        <a id="L292"></a>if line, e = readLineBytes(b); e != nil {
            <a id="L293"></a>return &#34;&#34;, &#34;&#34;, e
        <a id="L294"></a>}
        <a id="L295"></a>value += &#34; &#34; + string(line);

        <a id="L297"></a>if len(value) &gt;= maxValueLength {
            <a id="L298"></a>return &#34;&#34;, &#34;&#34;, &amp;badStringError{&#34;value too long for key&#34;, key}
        <a id="L299"></a>}
    <a id="L300"></a>}
    <a id="L301"></a>return key, value, nil;

<a id="L303"></a>Malformed:
    <a id="L304"></a>return &#34;&#34;, &#34;&#34;, &amp;badStringError{&#34;malformed header line&#34;, string(line)};
<a id="L305"></a>}

<a id="L307"></a><span class="comment">// Convert decimal at s[i:len(s)] to integer,</span>
<a id="L308"></a><span class="comment">// returning value, string position where the digits stopped,</span>
<a id="L309"></a><span class="comment">// and whether there was a valid number (digits, not too big).</span>
<a id="L310"></a>func atoi(s string, i int) (n, i1 int, ok bool) {
    <a id="L311"></a>const Big = 1000000;
    <a id="L312"></a>if i &gt;= len(s) || s[i] &lt; &#39;0&#39; || s[i] &gt; &#39;9&#39; {
        <a id="L313"></a>return 0, 0, false
    <a id="L314"></a>}
    <a id="L315"></a>n = 0;
    <a id="L316"></a>for ; i &lt; len(s) &amp;&amp; &#39;0&#39; &lt;= s[i] &amp;&amp; s[i] &lt;= &#39;9&#39;; i++ {
        <a id="L317"></a>n = n*10 + int(s[i]-&#39;0&#39;);
        <a id="L318"></a>if n &gt; Big {
            <a id="L319"></a>return 0, 0, false
        <a id="L320"></a>}
    <a id="L321"></a>}
    <a id="L322"></a>return n, i, true;
<a id="L323"></a>}

<a id="L325"></a><span class="comment">// Parse HTTP version: &#34;HTTP/1.2&#34; -&gt; (1, 2, true).</span>
<a id="L326"></a>func parseHTTPVersion(vers string) (int, int, bool) {
    <a id="L327"></a>if vers[0:5] != &#34;HTTP/&#34; {
        <a id="L328"></a>return 0, 0, false
    <a id="L329"></a>}
    <a id="L330"></a>major, i, ok := atoi(vers, 5);
    <a id="L331"></a>if !ok || i &gt;= len(vers) || vers[i] != &#39;.&#39; {
        <a id="L332"></a>return 0, 0, false
    <a id="L333"></a>}
    <a id="L334"></a>var minor int;
    <a id="L335"></a>minor, i, ok = atoi(vers, i+1);
    <a id="L336"></a>if !ok || i != len(vers) {
        <a id="L337"></a>return 0, 0, false
    <a id="L338"></a>}
    <a id="L339"></a>return major, minor, true;
<a id="L340"></a>}

<a id="L342"></a>var cmap = make(map[string]string)

<a id="L344"></a><span class="comment">// CanonicalHeaderKey returns the canonical format of the</span>
<a id="L345"></a><span class="comment">// HTTP header key s.  The canonicalization converts the first</span>
<a id="L346"></a><span class="comment">// letter and any letter following a hyphen to upper case;</span>
<a id="L347"></a><span class="comment">// the rest are converted to lowercase.  For example, the</span>
<a id="L348"></a><span class="comment">// canonical key for &#34;accept-encoding&#34; is &#34;Accept-Encoding&#34;.</span>
<a id="L349"></a>func CanonicalHeaderKey(s string) string {
    <a id="L350"></a>if t, ok := cmap[s]; ok {
        <a id="L351"></a>return t
    <a id="L352"></a>}

    <a id="L354"></a><span class="comment">// canonicalize: first letter upper case</span>
    <a id="L355"></a><span class="comment">// and upper case after each dash.</span>
    <a id="L356"></a><span class="comment">// (Host, User-Agent, If-Modified-Since).</span>
    <a id="L357"></a><span class="comment">// HTTP headers are ASCII only, so no Unicode issues.</span>
    <a id="L358"></a>a := strings.Bytes(s);
    <a id="L359"></a>upper := true;
    <a id="L360"></a>for i, v := range a {
        <a id="L361"></a>if upper &amp;&amp; &#39;a&#39; &lt;= v &amp;&amp; v &lt;= &#39;z&#39; {
            <a id="L362"></a>a[i] = v + &#39;A&#39; - &#39;a&#39;
        <a id="L363"></a>}
        <a id="L364"></a>if !upper &amp;&amp; &#39;A&#39; &lt;= v &amp;&amp; v &lt;= &#39;Z&#39; {
            <a id="L365"></a>a[i] = v + &#39;a&#39; - &#39;A&#39;
        <a id="L366"></a>}
        <a id="L367"></a>upper = false;
        <a id="L368"></a>if v == &#39;-&#39; {
            <a id="L369"></a>upper = true
        <a id="L370"></a>}
    <a id="L371"></a>}
    <a id="L372"></a>t := string(a);
    <a id="L373"></a>cmap[s] = t;
    <a id="L374"></a>return t;
<a id="L375"></a>}

<a id="L377"></a>type chunkedReader struct {
    <a id="L378"></a>r   *bufio.Reader;
    <a id="L379"></a>n   uint64; <span class="comment">// unread bytes in chunk</span>
    <a id="L380"></a>err os.Error;
<a id="L381"></a>}

<a id="L383"></a>func newChunkedReader(r *bufio.Reader) *chunkedReader {
    <a id="L384"></a>return &amp;chunkedReader{r: r}
<a id="L385"></a>}

<a id="L387"></a>func (cr *chunkedReader) beginChunk() {
    <a id="L388"></a><span class="comment">// chunk-size CRLF</span>
    <a id="L389"></a>var line string;
    <a id="L390"></a>line, cr.err = readLine(cr.r);
    <a id="L391"></a>if cr.err != nil {
        <a id="L392"></a>return
    <a id="L393"></a>}
    <a id="L394"></a>cr.n, cr.err = strconv.Btoui64(line, 16);
    <a id="L395"></a>if cr.err != nil {
        <a id="L396"></a>return
    <a id="L397"></a>}
    <a id="L398"></a>if cr.n == 0 {
        <a id="L399"></a><span class="comment">// trailer CRLF</span>
        <a id="L400"></a>for {
            <a id="L401"></a>line, cr.err = readLine(cr.r);
            <a id="L402"></a>if cr.err != nil {
                <a id="L403"></a>return
            <a id="L404"></a>}
            <a id="L405"></a>if line == &#34;&#34; {
                <a id="L406"></a>break
            <a id="L407"></a>}
        <a id="L408"></a>}
        <a id="L409"></a>cr.err = os.EOF;
    <a id="L410"></a>}
<a id="L411"></a>}

<a id="L413"></a>func (cr *chunkedReader) Read(b []uint8) (n int, err os.Error) {
    <a id="L414"></a>if cr.err != nil {
        <a id="L415"></a>return 0, cr.err
    <a id="L416"></a>}
    <a id="L417"></a>if cr.n == 0 {
        <a id="L418"></a>cr.beginChunk();
        <a id="L419"></a>if cr.err != nil {
            <a id="L420"></a>return 0, cr.err
        <a id="L421"></a>}
    <a id="L422"></a>}
    <a id="L423"></a>if uint64(len(b)) &gt; cr.n {
        <a id="L424"></a>b = b[0:cr.n]
    <a id="L425"></a>}
    <a id="L426"></a>n, cr.err = cr.r.Read(b);
    <a id="L427"></a>cr.n -= uint64(n);
    <a id="L428"></a>if cr.n == 0 &amp;&amp; cr.err == nil {
        <a id="L429"></a><span class="comment">// end of chunk (CRLF)</span>
        <a id="L430"></a>b := make([]byte, 2);
        <a id="L431"></a>if _, cr.err = io.ReadFull(cr.r, b); cr.err == nil {
            <a id="L432"></a>if b[0] != &#39;\r&#39; || b[1] != &#39;\n&#39; {
                <a id="L433"></a>cr.err = os.NewError(&#34;malformed chunked encoding&#34;)
            <a id="L434"></a>}
        <a id="L435"></a>}
    <a id="L436"></a>}
    <a id="L437"></a>return n, cr.err;
<a id="L438"></a>}

<a id="L440"></a><span class="comment">// ReadRequest reads and parses a request from b.</span>
<a id="L441"></a>func ReadRequest(b *bufio.Reader) (req *Request, err os.Error) {
    <a id="L442"></a>req = new(Request);

    <a id="L444"></a><span class="comment">// First line: GET /index.html HTTP/1.0</span>
    <a id="L445"></a>var s string;
    <a id="L446"></a>if s, err = readLine(b); err != nil {
        <a id="L447"></a>return nil, err
    <a id="L448"></a>}

    <a id="L450"></a>var f []string;
    <a id="L451"></a>if f = strings.Split(s, &#34; &#34;, 3); len(f) &lt; 3 {
        <a id="L452"></a>return nil, &amp;badStringError{&#34;malformed HTTP request&#34;, s}
    <a id="L453"></a>}
    <a id="L454"></a>req.Method, req.RawURL, req.Proto = f[0], f[1], f[2];
    <a id="L455"></a>var ok bool;
    <a id="L456"></a>if req.ProtoMajor, req.ProtoMinor, ok = parseHTTPVersion(req.Proto); !ok {
        <a id="L457"></a>return nil, &amp;badStringError{&#34;malformed HTTP version&#34;, req.Proto}
    <a id="L458"></a>}

    <a id="L460"></a>if req.URL, err = ParseURL(req.RawURL); err != nil {
        <a id="L461"></a>return nil, err
    <a id="L462"></a>}

    <a id="L464"></a><span class="comment">// Subsequent lines: Key: value.</span>
    <a id="L465"></a>nheader := 0;
    <a id="L466"></a>req.Header = make(map[string]string);
    <a id="L467"></a>for {
        <a id="L468"></a>var key, value string;
        <a id="L469"></a>if key, value, err = readKeyValue(b); err != nil {
            <a id="L470"></a>return nil, err
        <a id="L471"></a>}
        <a id="L472"></a>if key == &#34;&#34; {
            <a id="L473"></a>break
        <a id="L474"></a>}
        <a id="L475"></a>if nheader++; nheader &gt;= maxHeaderLines {
            <a id="L476"></a>return nil, ErrHeaderTooLong
        <a id="L477"></a>}

        <a id="L479"></a>key = CanonicalHeaderKey(key);

        <a id="L481"></a><span class="comment">// RFC 2616 says that if you send the same header key</span>
        <a id="L482"></a><span class="comment">// multiple times, it has to be semantically equivalent</span>
        <a id="L483"></a><span class="comment">// to concatenating the values separated by commas.</span>
        <a id="L484"></a>oldvalue, present := req.Header[key];
        <a id="L485"></a>if present {
            <a id="L486"></a>req.Header[key] = oldvalue + &#34;,&#34; + value
        <a id="L487"></a>} else {
            <a id="L488"></a>req.Header[key] = value
        <a id="L489"></a>}
    <a id="L490"></a>}

    <a id="L492"></a><span class="comment">// RFC2616: Must treat</span>
    <a id="L493"></a><span class="comment">//	GET /index.html HTTP/1.1</span>
    <a id="L494"></a><span class="comment">//	Host: www.google.com</span>
    <a id="L495"></a><span class="comment">// and</span>
    <a id="L496"></a><span class="comment">//	GET http://www.google.com/index.html HTTP/1.1</span>
    <a id="L497"></a><span class="comment">//	Host: doesntmatter</span>
    <a id="L498"></a><span class="comment">// the same.  In the second case, any Host line is ignored.</span>
    <a id="L499"></a>if v, present := req.Header[&#34;Host&#34;]; present &amp;&amp; req.URL.Host == &#34;&#34; {
        <a id="L500"></a>req.Host = v
    <a id="L501"></a>}

    <a id="L503"></a><span class="comment">// RFC2616: Should treat</span>
    <a id="L504"></a><span class="comment">//	Pragma: no-cache</span>
    <a id="L505"></a><span class="comment">// like</span>
    <a id="L506"></a><span class="comment">//	Cache-Control: no-cache</span>
    <a id="L507"></a>if v, present := req.Header[&#34;Pragma&#34;]; present &amp;&amp; v == &#34;no-cache&#34; {
        <a id="L508"></a>if _, presentcc := req.Header[&#34;Cache-Control&#34;]; !presentcc {
            <a id="L509"></a>req.Header[&#34;Cache-Control&#34;] = &#34;no-cache&#34;
        <a id="L510"></a>}
    <a id="L511"></a>}

    <a id="L513"></a><span class="comment">// Determine whether to hang up after sending the reply.</span>
    <a id="L514"></a>if req.ProtoMajor &lt; 1 || (req.ProtoMajor == 1 &amp;&amp; req.ProtoMinor &lt; 1) {
        <a id="L515"></a>req.Close = true
    <a id="L516"></a>} else if v, present := req.Header[&#34;Connection&#34;]; present {
        <a id="L517"></a><span class="comment">// TODO: Should split on commas, toss surrounding white space,</span>
        <a id="L518"></a><span class="comment">// and check each field.</span>
        <a id="L519"></a>if v == &#34;close&#34; {
            <a id="L520"></a>req.Close = true
        <a id="L521"></a>}
    <a id="L522"></a>}

    <a id="L524"></a><span class="comment">// Pull out useful fields as a convenience to clients.</span>
    <a id="L525"></a>if v, present := req.Header[&#34;Referer&#34;]; present {
        <a id="L526"></a>req.Referer = v
    <a id="L527"></a>}
    <a id="L528"></a>if v, present := req.Header[&#34;User-Agent&#34;]; present {
        <a id="L529"></a>req.UserAgent = v
    <a id="L530"></a>}

    <a id="L532"></a><span class="comment">// TODO: Parse specific header values:</span>
    <a id="L533"></a><span class="comment">//	Accept</span>
    <a id="L534"></a><span class="comment">//	Accept-Encoding</span>
    <a id="L535"></a><span class="comment">//	Accept-Language</span>
    <a id="L536"></a><span class="comment">//	Authorization</span>
    <a id="L537"></a><span class="comment">//	Cache-Control</span>
    <a id="L538"></a><span class="comment">//	Connection</span>
    <a id="L539"></a><span class="comment">//	Date</span>
    <a id="L540"></a><span class="comment">//	Expect</span>
    <a id="L541"></a><span class="comment">//	From</span>
    <a id="L542"></a><span class="comment">//	If-Match</span>
    <a id="L543"></a><span class="comment">//	If-Modified-Since</span>
    <a id="L544"></a><span class="comment">//	If-None-Match</span>
    <a id="L545"></a><span class="comment">//	If-Range</span>
    <a id="L546"></a><span class="comment">//	If-Unmodified-Since</span>
    <a id="L547"></a><span class="comment">//	Max-Forwards</span>
    <a id="L548"></a><span class="comment">//	Proxy-Authorization</span>
    <a id="L549"></a><span class="comment">//	Referer [sic]</span>
    <a id="L550"></a><span class="comment">//	TE (transfer-codings)</span>
    <a id="L551"></a><span class="comment">//	Trailer</span>
    <a id="L552"></a><span class="comment">//	Transfer-Encoding</span>
    <a id="L553"></a><span class="comment">//	Upgrade</span>
    <a id="L554"></a><span class="comment">//	User-Agent</span>
    <a id="L555"></a><span class="comment">//	Via</span>
    <a id="L556"></a><span class="comment">//	Warning</span>

    <a id="L558"></a><span class="comment">// A message body exists when either Content-Length or Transfer-Encoding</span>
    <a id="L559"></a><span class="comment">// headers are present. Transfer-Encoding trumps Content-Length.</span>
    <a id="L560"></a>if v, present := req.Header[&#34;Transfer-Encoding&#34;]; present &amp;&amp; v == &#34;chunked&#34; {
        <a id="L561"></a>req.Body = newChunkedReader(b)
    <a id="L562"></a>} else if v, present := req.Header[&#34;Content-Length&#34;]; present {
        <a id="L563"></a>length, err := strconv.Btoui64(v, 10);
        <a id="L564"></a>if err != nil {
            <a id="L565"></a>return nil, &amp;badStringError{&#34;invalid Content-Length&#34;, v}
        <a id="L566"></a>}
        <a id="L567"></a><span class="comment">// TODO: limit the Content-Length. This is an easy DoS vector.</span>
        <a id="L568"></a>raw := make([]byte, length);
        <a id="L569"></a>n, err := b.Read(raw);
        <a id="L570"></a>if err != nil || uint64(n) &lt; length {
            <a id="L571"></a>return nil, ErrShortBody
        <a id="L572"></a>}
        <a id="L573"></a>req.Body = bytes.NewBuffer(raw);
    <a id="L574"></a>}

    <a id="L576"></a>return req, nil;
<a id="L577"></a>}

<a id="L579"></a>func parseForm(query string) (m map[string][]string, err os.Error) {
    <a id="L580"></a>data := make(map[string]*vector.StringVector);
    <a id="L581"></a>for _, kv := range strings.Split(query, &#34;&amp;&#34;, 0) {
        <a id="L582"></a>kvPair := strings.Split(kv, &#34;=&#34;, 2);

        <a id="L584"></a>var key, value string;
        <a id="L585"></a>var e os.Error;
        <a id="L586"></a>key, e = URLUnescape(kvPair[0]);
        <a id="L587"></a>if e == nil &amp;&amp; len(kvPair) &gt; 1 {
            <a id="L588"></a>value, e = URLUnescape(kvPair[1])
        <a id="L589"></a>}
        <a id="L590"></a>if e != nil {
            <a id="L591"></a>err = e
        <a id="L592"></a>}

        <a id="L594"></a>vec, ok := data[key];
        <a id="L595"></a>if !ok {
            <a id="L596"></a>vec = vector.NewStringVector(0);
            <a id="L597"></a>data[key] = vec;
        <a id="L598"></a>}
        <a id="L599"></a>vec.Push(value);
    <a id="L600"></a>}

    <a id="L602"></a>m = make(map[string][]string);
    <a id="L603"></a>for k, vec := range data {
        <a id="L604"></a>m[k] = vec.Data()
    <a id="L605"></a>}

    <a id="L607"></a>return;
<a id="L608"></a>}

<a id="L610"></a><span class="comment">// ParseForm parses the request body as a form for POST requests, or the raw query for GET requests.</span>
<a id="L611"></a><span class="comment">// It is idempotent.</span>
<a id="L612"></a>func (r *Request) ParseForm() (err os.Error) {
    <a id="L613"></a>if r.Form != nil {
        <a id="L614"></a>return
    <a id="L615"></a>}

    <a id="L617"></a>var query string;

    <a id="L619"></a>switch r.Method {
    <a id="L620"></a>case &#34;GET&#34;:
        <a id="L621"></a>query = r.URL.RawQuery
    <a id="L622"></a>case &#34;POST&#34;:
        <a id="L623"></a>if r.Body == nil {
            <a id="L624"></a>return os.ErrorString(&#34;missing form body&#34;)
        <a id="L625"></a>}
        <a id="L626"></a>ct, _ := r.Header[&#34;Content-Type&#34;];
        <a id="L627"></a>switch strings.Split(ct, &#34;;&#34;, 2)[0] {
        <a id="L628"></a>case &#34;text/plain&#34;, &#34;application/x-www-form-urlencoded&#34;, &#34;&#34;:
            <a id="L629"></a>var b []byte;
            <a id="L630"></a>if b, err = io.ReadAll(r.Body); err != nil {
                <a id="L631"></a>return
            <a id="L632"></a>}
            <a id="L633"></a>query = string(b);
        <a id="L634"></a><span class="comment">// TODO(dsymonds): Handle multipart/form-data</span>
        <a id="L635"></a>default:
            <a id="L636"></a>return &amp;badStringError{&#34;unknown Content-Type&#34;, ct}
        <a id="L637"></a>}
    <a id="L638"></a>}
    <a id="L639"></a>r.Form, err = parseForm(query);
    <a id="L640"></a>return;
<a id="L641"></a>}

<a id="L643"></a><span class="comment">// FormValue returns the first value for the named component of the query.</span>
<a id="L644"></a><span class="comment">// FormValue calls ParseForm if necessary.</span>
<a id="L645"></a>func (r *Request) FormValue(key string) string {
    <a id="L646"></a>if r.Form == nil {
        <a id="L647"></a>r.ParseForm()
    <a id="L648"></a>}
    <a id="L649"></a>if vs, ok := r.Form[key]; ok &amp;&amp; len(vs) &gt; 0 {
        <a id="L650"></a>return vs[0]
    <a id="L651"></a>}
    <a id="L652"></a>return &#34;&#34;;
<a id="L653"></a>}
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
