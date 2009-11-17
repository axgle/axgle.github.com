<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/client.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/http/client.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Primitive HTTP client. See RFC 2616.</span>

<a id="L7"></a>package http

<a id="L9"></a>import (
    <a id="L10"></a>&#34;bufio&#34;;
    <a id="L11"></a>&#34;fmt&#34;;
    <a id="L12"></a>&#34;io&#34;;
    <a id="L13"></a>&#34;net&#34;;
    <a id="L14"></a>&#34;os&#34;;
    <a id="L15"></a>&#34;strconv&#34;;
    <a id="L16"></a>&#34;strings&#34;;
<a id="L17"></a>)

<a id="L19"></a><span class="comment">// Response represents the response from an HTTP request.</span>
<a id="L20"></a>type Response struct {
    <a id="L21"></a>Status     string; <span class="comment">// e.g. &#34;200 OK&#34;</span>
    <a id="L22"></a>StatusCode int;    <span class="comment">// e.g. 200</span>

    <a id="L24"></a><span class="comment">// Header maps header keys to values.  If the response had multiple</span>
    <a id="L25"></a><span class="comment">// headers with the same key, they will be concatenated, with comma</span>
    <a id="L26"></a><span class="comment">// delimiters.  (Section 4.2 of RFC 2616 requires that multiple headers</span>
    <a id="L27"></a><span class="comment">// be semantically equivalent to a comma-delimited sequence.)</span>
    <a id="L28"></a><span class="comment">//</span>
    <a id="L29"></a><span class="comment">// Keys in the map are canonicalized (see CanonicalHeaderKey).</span>
    <a id="L30"></a>Header map[string]string;

    <a id="L32"></a><span class="comment">// Stream from which the response body can be read.</span>
    <a id="L33"></a>Body io.ReadCloser;
<a id="L34"></a>}

<a id="L36"></a><span class="comment">// GetHeader returns the value of the response header with the given</span>
<a id="L37"></a><span class="comment">// key, and true.  If there were multiple headers with this key, their</span>
<a id="L38"></a><span class="comment">// values are concatenated, with a comma delimiter.  If there were no</span>
<a id="L39"></a><span class="comment">// response headers with the given key, it returns the empty string and</span>
<a id="L40"></a><span class="comment">// false.  Keys are not case sensitive.</span>
<a id="L41"></a>func (r *Response) GetHeader(key string) (value string) {
    <a id="L42"></a>value, _ = r.Header[CanonicalHeaderKey(key)];
    <a id="L43"></a>return;
<a id="L44"></a>}

<a id="L46"></a><span class="comment">// AddHeader adds a value under the given key.  Keys are not case sensitive.</span>
<a id="L47"></a>func (r *Response) AddHeader(key, value string) {
    <a id="L48"></a>key = CanonicalHeaderKey(key);

    <a id="L50"></a>oldValues, oldValuesPresent := r.Header[key];
    <a id="L51"></a>if oldValuesPresent {
        <a id="L52"></a>r.Header[key] = oldValues + &#34;,&#34; + value
    <a id="L53"></a>} else {
        <a id="L54"></a>r.Header[key] = value
    <a id="L55"></a>}
<a id="L56"></a>}

<a id="L58"></a><span class="comment">// Given a string of the form &#34;host&#34;, &#34;host:port&#34;, or &#34;[ipv6::address]:port&#34;,</span>
<a id="L59"></a><span class="comment">// return true if the string includes a port.</span>
<a id="L60"></a>func hasPort(s string) bool { return strings.LastIndex(s, &#34;:&#34;) &gt; strings.LastIndex(s, &#34;]&#34;) }

<a id="L62"></a><span class="comment">// Used in Send to implement io.ReadCloser by bundling together the</span>
<a id="L63"></a><span class="comment">// io.BufReader through which we read the response, and the underlying</span>
<a id="L64"></a><span class="comment">// network connection.</span>
<a id="L65"></a>type readClose struct {
    <a id="L66"></a>io.Reader;
    <a id="L67"></a>io.Closer;
<a id="L68"></a>}

<a id="L70"></a><span class="comment">// ReadResponse reads and returns an HTTP response from r.</span>
<a id="L71"></a>func ReadResponse(r *bufio.Reader) (*Response, os.Error) {
    <a id="L72"></a>resp := new(Response);

    <a id="L74"></a><span class="comment">// Parse the first line of the response.</span>
    <a id="L75"></a>resp.Header = make(map[string]string);

    <a id="L77"></a>line, err := readLine(r);
    <a id="L78"></a>if err != nil {
        <a id="L79"></a>return nil, err
    <a id="L80"></a>}
    <a id="L81"></a>f := strings.Split(line, &#34; &#34;, 3);
    <a id="L82"></a>if len(f) &lt; 3 {
        <a id="L83"></a>return nil, &amp;badStringError{&#34;malformed HTTP response&#34;, line}
    <a id="L84"></a>}
    <a id="L85"></a>resp.Status = f[1] + &#34; &#34; + f[2];
    <a id="L86"></a>resp.StatusCode, err = strconv.Atoi(f[1]);
    <a id="L87"></a>if err != nil {
        <a id="L88"></a>return nil, &amp;badStringError{&#34;malformed HTTP status code&#34;, f[1]}
    <a id="L89"></a>}

    <a id="L91"></a><span class="comment">// Parse the response headers.</span>
    <a id="L92"></a>for {
        <a id="L93"></a>key, value, err := readKeyValue(r);
        <a id="L94"></a>if err != nil {
            <a id="L95"></a>return nil, err
        <a id="L96"></a>}
        <a id="L97"></a>if key == &#34;&#34; {
            <a id="L98"></a>break <span class="comment">// end of response header</span>
        <a id="L99"></a>}
        <a id="L100"></a>resp.AddHeader(key, value);
    <a id="L101"></a>}

    <a id="L103"></a>return resp, nil;
<a id="L104"></a>}


<a id="L107"></a><span class="comment">// Send issues an HTTP request.  Caller should close resp.Body when done reading it.</span>
<a id="L108"></a><span class="comment">//</span>
<a id="L109"></a><span class="comment">// TODO: support persistent connections (multiple requests on a single connection).</span>
<a id="L110"></a><span class="comment">// send() method is nonpublic because, when we refactor the code for persistent</span>
<a id="L111"></a><span class="comment">// connections, it may no longer make sense to have a method with this signature.</span>
<a id="L112"></a>func send(req *Request) (resp *Response, err os.Error) {
    <a id="L113"></a>if req.URL.Scheme != &#34;http&#34; {
        <a id="L114"></a>return nil, &amp;badStringError{&#34;unsupported protocol scheme&#34;, req.URL.Scheme}
    <a id="L115"></a>}

    <a id="L117"></a>addr := req.URL.Host;
    <a id="L118"></a>if !hasPort(addr) {
        <a id="L119"></a>addr += &#34;:http&#34;
    <a id="L120"></a>}
    <a id="L121"></a>conn, err := net.Dial(&#34;tcp&#34;, &#34;&#34;, addr);
    <a id="L122"></a>if err != nil {
        <a id="L123"></a>return nil, err
    <a id="L124"></a>}

    <a id="L126"></a>err = req.Write(conn);
    <a id="L127"></a>if err != nil {
        <a id="L128"></a>conn.Close();
        <a id="L129"></a>return nil, err;
    <a id="L130"></a>}

    <a id="L132"></a>reader := bufio.NewReader(conn);
    <a id="L133"></a>resp, err = ReadResponse(reader);
    <a id="L134"></a>if err != nil {
        <a id="L135"></a>conn.Close();
        <a id="L136"></a>return nil, err;
    <a id="L137"></a>}

    <a id="L139"></a>r := io.Reader(reader);
    <a id="L140"></a>if v := resp.GetHeader(&#34;Transfer-Encoding&#34;); v == &#34;chunked&#34; {
        <a id="L141"></a>r = newChunkedReader(reader)
    <a id="L142"></a>} else if v := resp.GetHeader(&#34;Content-Length&#34;); v != &#34;&#34; {
        <a id="L143"></a>n, err := strconv.Atoi64(v);
        <a id="L144"></a>if err != nil {
            <a id="L145"></a>return nil, &amp;badStringError{&#34;invalid Content-Length&#34;, v}
        <a id="L146"></a>}
        <a id="L147"></a>r = io.LimitReader(r, n);
    <a id="L148"></a>}
    <a id="L149"></a>resp.Body = readClose{r, conn};

    <a id="L151"></a>return;
<a id="L152"></a>}

<a id="L154"></a><span class="comment">// True if the specified HTTP status code is one for which the Get utility should</span>
<a id="L155"></a><span class="comment">// automatically redirect.</span>
<a id="L156"></a>func shouldRedirect(statusCode int) bool {
    <a id="L157"></a>switch statusCode {
    <a id="L158"></a>case StatusMovedPermanently, StatusFound, StatusSeeOther, StatusTemporaryRedirect:
        <a id="L159"></a>return true
    <a id="L160"></a>}
    <a id="L161"></a>return false;
<a id="L162"></a>}

<a id="L164"></a><span class="comment">// Get issues a GET to the specified URL.  If the response is one of the following</span>
<a id="L165"></a><span class="comment">// redirect codes, it follows the redirect, up to a maximum of 10 redirects:</span>
<a id="L166"></a><span class="comment">//</span>
<a id="L167"></a><span class="comment">//    301 (Moved Permanently)</span>
<a id="L168"></a><span class="comment">//    302 (Found)</span>
<a id="L169"></a><span class="comment">//    303 (See Other)</span>
<a id="L170"></a><span class="comment">//    307 (Temporary Redirect)</span>
<a id="L171"></a><span class="comment">//</span>
<a id="L172"></a><span class="comment">// finalURL is the URL from which the response was fetched -- identical to the input</span>
<a id="L173"></a><span class="comment">// URL unless redirects were followed.</span>
<a id="L174"></a><span class="comment">//</span>
<a id="L175"></a><span class="comment">// Caller should close r.Body when done reading it.</span>
<a id="L176"></a>func Get(url string) (r *Response, finalURL string, err os.Error) {
    <a id="L177"></a><span class="comment">// TODO: if/when we add cookie support, the redirected request shouldn&#39;t</span>
    <a id="L178"></a><span class="comment">// necessarily supply the same cookies as the original.</span>
    <a id="L179"></a><span class="comment">// TODO: set referrer header on redirects.</span>
    <a id="L180"></a>for redirect := 0; ; redirect++ {
        <a id="L181"></a>if redirect &gt;= 10 {
            <a id="L182"></a>err = os.ErrorString(&#34;stopped after 10 redirects&#34;);
            <a id="L183"></a>break;
        <a id="L184"></a>}

        <a id="L186"></a>var req Request;
        <a id="L187"></a>if req.URL, err = ParseURL(url); err != nil {
            <a id="L188"></a>break
        <a id="L189"></a>}
        <a id="L190"></a>if r, err = send(&amp;req); err != nil {
            <a id="L191"></a>break
        <a id="L192"></a>}
        <a id="L193"></a>if shouldRedirect(r.StatusCode) {
            <a id="L194"></a>r.Body.Close();
            <a id="L195"></a>if url = r.GetHeader(&#34;Location&#34;); url == &#34;&#34; {
                <a id="L196"></a>err = os.ErrorString(fmt.Sprintf(&#34;%d response missing Location header&#34;, r.StatusCode));
                <a id="L197"></a>break;
            <a id="L198"></a>}
            <a id="L199"></a>continue;
        <a id="L200"></a>}
        <a id="L201"></a>finalURL = url;
        <a id="L202"></a>return;
    <a id="L203"></a>}

    <a id="L205"></a>err = &amp;URLError{&#34;Get&#34;, url, err};
    <a id="L206"></a>return;
<a id="L207"></a>}


<a id="L210"></a><span class="comment">// Post issues a POST to the specified URL.</span>
<a id="L211"></a><span class="comment">//</span>
<a id="L212"></a><span class="comment">// Caller should close r.Body when done reading it.</span>
<a id="L213"></a>func Post(url string, bodyType string, body io.Reader) (r *Response, err os.Error) {
    <a id="L214"></a>var req Request;
    <a id="L215"></a>req.Method = &#34;POST&#34;;
    <a id="L216"></a>req.Body = body;
    <a id="L217"></a>req.Header = map[string]string{
        <a id="L218"></a>&#34;Content-Type&#34;: bodyType,
        <a id="L219"></a>&#34;Transfer-Encoding&#34;: &#34;chunked&#34;,
    <a id="L220"></a>};

    <a id="L222"></a>req.URL, err = ParseURL(url);
    <a id="L223"></a>if err != nil {
        <a id="L224"></a>return nil, err
    <a id="L225"></a>}

    <a id="L227"></a>return send(&amp;req);
<a id="L228"></a>}
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
