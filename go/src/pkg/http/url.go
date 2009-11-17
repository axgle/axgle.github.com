<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/url.go</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/http/url.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">// Parse URLs (actually URIs, but that seems overly pedantic).</span>
<a id="L6"></a><span class="comment">// RFC 2396</span>

<a id="L8"></a>package http

<a id="L10"></a>import (
    <a id="L11"></a>&#34;os&#34;;
    <a id="L12"></a>&#34;strconv&#34;;
    <a id="L13"></a>&#34;strings&#34;;
<a id="L14"></a>)

<a id="L16"></a><span class="comment">// URLError reports an error and the operation and URL that caused it.</span>
<a id="L17"></a>type URLError struct {
    <a id="L18"></a>Op    string;
    <a id="L19"></a>URL   string;
    <a id="L20"></a>Error os.Error;
<a id="L21"></a>}

<a id="L23"></a>func (e *URLError) String() string { return e.Op + &#34; &#34; + e.URL + &#34;: &#34; + e.Error.String() }

<a id="L25"></a>func ishex(c byte) bool {
    <a id="L26"></a>switch {
    <a id="L27"></a>case &#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39;:
        <a id="L28"></a>return true
    <a id="L29"></a>case &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;f&#39;:
        <a id="L30"></a>return true
    <a id="L31"></a>case &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;F&#39;:
        <a id="L32"></a>return true
    <a id="L33"></a>}
    <a id="L34"></a>return false;
<a id="L35"></a>}

<a id="L37"></a>func unhex(c byte) byte {
    <a id="L38"></a>switch {
    <a id="L39"></a>case &#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39;:
        <a id="L40"></a>return c - &#39;0&#39;
    <a id="L41"></a>case &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;f&#39;:
        <a id="L42"></a>return c - &#39;a&#39; + 10
    <a id="L43"></a>case &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;F&#39;:
        <a id="L44"></a>return c - &#39;A&#39; + 10
    <a id="L45"></a>}
    <a id="L46"></a>return 0;
<a id="L47"></a>}

<a id="L49"></a>type URLEscapeError string

<a id="L51"></a>func (e URLEscapeError) String() string {
    <a id="L52"></a>return &#34;invalid URL escape &#34; + strconv.Quote(string(e))
<a id="L53"></a>}

<a id="L55"></a><span class="comment">// Return true if the specified character should be escaped when appearing in a</span>
<a id="L56"></a><span class="comment">// URL string.</span>
<a id="L57"></a><span class="comment">//</span>
<a id="L58"></a><span class="comment">// TODO: for now, this is a hack; it only flags a few common characters that have</span>
<a id="L59"></a><span class="comment">// special meaning in URLs.  That will get the job done in the common cases.</span>
<a id="L60"></a>func shouldEscape(c byte) bool {
    <a id="L61"></a>switch c {
    <a id="L62"></a>case &#39; &#39;, &#39;?&#39;, &#39;&amp;&#39;, &#39;=&#39;, &#39;#&#39;, &#39;+&#39;, &#39;%&#39;:
        <a id="L63"></a>return true
    <a id="L64"></a>}
    <a id="L65"></a>return false;
<a id="L66"></a>}

<a id="L68"></a><span class="comment">// URLUnescape unescapes a URL-encoded string,</span>
<a id="L69"></a><span class="comment">// converting %AB into the byte 0xAB and &#39;+&#39; into &#39; &#39; (space).</span>
<a id="L70"></a><span class="comment">// It returns an error if any % is not followed</span>
<a id="L71"></a><span class="comment">// by two hexadecimal digits.</span>
<a id="L72"></a>func URLUnescape(s string) (string, os.Error) {
    <a id="L73"></a><span class="comment">// Count %, check that they&#39;re well-formed.</span>
    <a id="L74"></a>n := 0;
    <a id="L75"></a>hasPlus := false;
    <a id="L76"></a>for i := 0; i &lt; len(s); {
        <a id="L77"></a>switch s[i] {
        <a id="L78"></a>case &#39;%&#39;:
            <a id="L79"></a>n++;
            <a id="L80"></a>if i+2 &gt;= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
                <a id="L81"></a>s = s[i:len(s)];
                <a id="L82"></a>if len(s) &gt; 3 {
                    <a id="L83"></a>s = s[0:3]
                <a id="L84"></a>}
                <a id="L85"></a>return &#34;&#34;, URLEscapeError(s);
            <a id="L86"></a>}
            <a id="L87"></a>i += 3;
        <a id="L88"></a>case &#39;+&#39;:
            <a id="L89"></a>hasPlus = true;
            <a id="L90"></a>i++;
        <a id="L91"></a>default:
            <a id="L92"></a>i++
        <a id="L93"></a>}
    <a id="L94"></a>}

    <a id="L96"></a>if n == 0 &amp;&amp; !hasPlus {
        <a id="L97"></a>return s, nil
    <a id="L98"></a>}

    <a id="L100"></a>t := make([]byte, len(s)-2*n);
    <a id="L101"></a>j := 0;
    <a id="L102"></a>for i := 0; i &lt; len(s); {
        <a id="L103"></a>switch s[i] {
        <a id="L104"></a>case &#39;%&#39;:
            <a id="L105"></a>t[j] = unhex(s[i+1])&lt;&lt;4 | unhex(s[i+2]);
            <a id="L106"></a>j++;
            <a id="L107"></a>i += 3;
        <a id="L108"></a>case &#39;+&#39;:
            <a id="L109"></a>t[j] = &#39; &#39;;
            <a id="L110"></a>j++;
            <a id="L111"></a>i++;
        <a id="L112"></a>default:
            <a id="L113"></a>t[j] = s[i];
            <a id="L114"></a>j++;
            <a id="L115"></a>i++;
        <a id="L116"></a>}
    <a id="L117"></a>}
    <a id="L118"></a>return string(t), nil;
<a id="L119"></a>}

<a id="L121"></a><span class="comment">// URLEscape converts a string into URL-encoded form.</span>
<a id="L122"></a>func URLEscape(s string) string {
    <a id="L123"></a>spaceCount, hexCount := 0, 0;
    <a id="L124"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L125"></a>c := s[i];
        <a id="L126"></a>if shouldEscape(c) {
            <a id="L127"></a>if c == &#39; &#39; {
                <a id="L128"></a>spaceCount++
            <a id="L129"></a>} else {
                <a id="L130"></a>hexCount++
            <a id="L131"></a>}
        <a id="L132"></a>}
    <a id="L133"></a>}

    <a id="L135"></a>if spaceCount == 0 &amp;&amp; hexCount == 0 {
        <a id="L136"></a>return s
    <a id="L137"></a>}

    <a id="L139"></a>t := make([]byte, len(s)+2*hexCount);
    <a id="L140"></a>j := 0;
    <a id="L141"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L142"></a>switch c := s[i]; {
        <a id="L143"></a>case c == &#39; &#39;:
            <a id="L144"></a>t[j] = &#39;+&#39;;
            <a id="L145"></a>j++;
        <a id="L146"></a>case shouldEscape(c):
            <a id="L147"></a>t[j] = &#39;%&#39;;
            <a id="L148"></a>t[j+1] = &#34;0123456789abcdef&#34;[c&gt;&gt;4];
            <a id="L149"></a>t[j+2] = &#34;0123456789abcdef&#34;[c&amp;15];
            <a id="L150"></a>j += 3;
        <a id="L151"></a>default:
            <a id="L152"></a>t[j] = s[i];
            <a id="L153"></a>j++;
        <a id="L154"></a>}
    <a id="L155"></a>}
    <a id="L156"></a>return string(t);
<a id="L157"></a>}

<a id="L159"></a><span class="comment">// A URL represents a parsed URL (technically, a URI reference).</span>
<a id="L160"></a><span class="comment">// The general form represented is:</span>
<a id="L161"></a><span class="comment">//	scheme://[userinfo@]host/path[?query][#fragment]</span>
<a id="L162"></a><span class="comment">// The Raw, RawPath, and RawQuery fields are in &#34;wire format&#34; (special</span>
<a id="L163"></a><span class="comment">// characters must be hex-escaped if not meant to have special meaning).</span>
<a id="L164"></a><span class="comment">// All other fields are logical values; &#39;+&#39; or &#39;%&#39; represent themselves.</span>
<a id="L165"></a><span class="comment">//</span>
<a id="L166"></a><span class="comment">// Note, the reason for using wire format for the query is that it needs</span>
<a id="L167"></a><span class="comment">// to be split into key/value pairs before decoding.</span>
<a id="L168"></a>type URL struct {
    <a id="L169"></a>Raw       string; <span class="comment">// the original string</span>
    <a id="L170"></a>Scheme    string; <span class="comment">// scheme</span>
    <a id="L171"></a>RawPath   string; <span class="comment">// //[userinfo@]host/path[?query][#fragment]</span>
    <a id="L172"></a>Authority string; <span class="comment">// [userinfo@]host</span>
    <a id="L173"></a>Userinfo  string; <span class="comment">// userinfo</span>
    <a id="L174"></a>Host      string; <span class="comment">// host</span>
    <a id="L175"></a>Path      string; <span class="comment">// /path</span>
    <a id="L176"></a>RawQuery  string; <span class="comment">// query</span>
    <a id="L177"></a>Fragment  string; <span class="comment">// fragment</span>
<a id="L178"></a>}

<a id="L180"></a><span class="comment">// Maybe rawurl is of the form scheme:path.</span>
<a id="L181"></a><span class="comment">// (Scheme must be [a-zA-Z][a-zA-Z0-9+-.]*)</span>
<a id="L182"></a><span class="comment">// If so, return scheme, path; else return &#34;&#34;, rawurl.</span>
<a id="L183"></a>func getscheme(rawurl string) (scheme, path string, err os.Error) {
    <a id="L184"></a>for i := 0; i &lt; len(rawurl); i++ {
        <a id="L185"></a>c := rawurl[i];
        <a id="L186"></a>switch {
        <a id="L187"></a>case &#39;a&#39; &lt;= c &amp;&amp; c &lt;= &#39;z&#39; || &#39;A&#39; &lt;= c &amp;&amp; c &lt;= &#39;Z&#39;:
        <a id="L188"></a><span class="comment">// do nothing</span>
        <a id="L189"></a>case &#39;0&#39; &lt;= c &amp;&amp; c &lt;= &#39;9&#39; || c == &#39;+&#39; || c == &#39;-&#39; || c == &#39;.&#39;:
            <a id="L190"></a>if i == 0 {
                <a id="L191"></a>return &#34;&#34;, rawurl, nil
            <a id="L192"></a>}
        <a id="L193"></a>case c == &#39;:&#39;:
            <a id="L194"></a>if i == 0 {
                <a id="L195"></a>return &#34;&#34;, &#34;&#34;, os.ErrorString(&#34;missing protocol scheme&#34;)
            <a id="L196"></a>}
            <a id="L197"></a>return rawurl[0:i], rawurl[i+1 : len(rawurl)], nil;
        <a id="L198"></a>default:
            <a id="L199"></a><span class="comment">// we have encountered an invalid character,</span>
            <a id="L200"></a><span class="comment">// so there is no valid scheme</span>
            <a id="L201"></a>return &#34;&#34;, rawurl, nil
        <a id="L202"></a>}
    <a id="L203"></a>}
    <a id="L204"></a>return &#34;&#34;, rawurl, nil;
<a id="L205"></a>}

<a id="L207"></a><span class="comment">// Maybe s is of the form t c u.</span>
<a id="L208"></a><span class="comment">// If so, return t, c u (or t, u if cutc == true).</span>
<a id="L209"></a><span class="comment">// If not, return s, &#34;&#34;.</span>
<a id="L210"></a>func split(s string, c byte, cutc bool) (string, string) {
    <a id="L211"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L212"></a>if s[i] == c {
            <a id="L213"></a>if cutc {
                <a id="L214"></a>return s[0:i], s[i+1 : len(s)]
            <a id="L215"></a>}
            <a id="L216"></a>return s[0:i], s[i:len(s)];
        <a id="L217"></a>}
    <a id="L218"></a>}
    <a id="L219"></a>return s, &#34;&#34;;
<a id="L220"></a>}

<a id="L222"></a><span class="comment">// TODO(rsc): The BUG comment is supposed to appear in the godoc output</span>
<a id="L223"></a><span class="comment">// in a BUGS section, but that got lost in the transition to godoc.</span>

<a id="L225"></a><span class="comment">// BUG(rsc): ParseURL should canonicalize the path,</span>
<a id="L226"></a><span class="comment">// removing unnecessary . and .. elements.</span>


<a id="L229"></a><span class="comment">// ParseURL parses rawurl into a URL structure.</span>
<a id="L230"></a><span class="comment">// The string rawurl is assumed not to have a #fragment suffix.</span>
<a id="L231"></a><span class="comment">// (Web browsers strip #fragment before sending the URL to a web server.)</span>
<a id="L232"></a>func ParseURL(rawurl string) (url *URL, err os.Error) {
    <a id="L233"></a>if rawurl == &#34;&#34; {
        <a id="L234"></a>err = os.ErrorString(&#34;empty url&#34;);
        <a id="L235"></a>goto Error;
    <a id="L236"></a>}
    <a id="L237"></a>url = new(URL);
    <a id="L238"></a>url.Raw = rawurl;

    <a id="L240"></a><span class="comment">// split off possible leading &#34;http:&#34;, &#34;mailto:&#34;, etc.</span>
    <a id="L241"></a>var path string;
    <a id="L242"></a>if url.Scheme, path, err = getscheme(rawurl); err != nil {
        <a id="L243"></a>goto Error
    <a id="L244"></a>}
    <a id="L245"></a>url.RawPath = path;

    <a id="L247"></a><span class="comment">// RFC 2396: a relative URI (no scheme) has a ?query,</span>
    <a id="L248"></a><span class="comment">// but absolute URIs only have query if path begins with /</span>
    <a id="L249"></a>if url.Scheme == &#34;&#34; || len(path) &gt; 0 &amp;&amp; path[0] == &#39;/&#39; {
        <a id="L250"></a>path, url.RawQuery = split(path, &#39;?&#39;, true)
    <a id="L251"></a>}

    <a id="L253"></a><span class="comment">// Maybe path is //authority/path</span>
    <a id="L254"></a>if len(path) &gt; 2 &amp;&amp; path[0:2] == &#34;//&#34; {
        <a id="L255"></a>url.Authority, path = split(path[2:len(path)], &#39;/&#39;, false)
    <a id="L256"></a>}

    <a id="L258"></a><span class="comment">// If there&#39;s no @, split&#39;s default is wrong.  Check explicitly.</span>
    <a id="L259"></a>if strings.Index(url.Authority, &#34;@&#34;) &lt; 0 {
        <a id="L260"></a>url.Host = url.Authority
    <a id="L261"></a>} else {
        <a id="L262"></a>url.Userinfo, url.Host = split(url.Authority, &#39;@&#39;, true)
    <a id="L263"></a>}

    <a id="L265"></a><span class="comment">// What&#39;s left is the path.</span>
    <a id="L266"></a><span class="comment">// TODO: Canonicalize (remove . and ..)?</span>
    <a id="L267"></a>if url.Path, err = URLUnescape(path); err != nil {
        <a id="L268"></a>goto Error
    <a id="L269"></a>}

    <a id="L271"></a><span class="comment">// Remove escapes from the Authority and Userinfo fields, and verify</span>
    <a id="L272"></a><span class="comment">// that Scheme and Host contain no escapes (that would be illegal).</span>
    <a id="L273"></a>if url.Authority, err = URLUnescape(url.Authority); err != nil {
        <a id="L274"></a>goto Error
    <a id="L275"></a>}
    <a id="L276"></a>if url.Userinfo, err = URLUnescape(url.Userinfo); err != nil {
        <a id="L277"></a>goto Error
    <a id="L278"></a>}
    <a id="L279"></a>if strings.Index(url.Scheme, &#34;%&#34;) &gt;= 0 {
        <a id="L280"></a>err = os.ErrorString(&#34;hexadecimal escape in scheme&#34;);
        <a id="L281"></a>goto Error;
    <a id="L282"></a>}
    <a id="L283"></a>if strings.Index(url.Host, &#34;%&#34;) &gt;= 0 {
        <a id="L284"></a>err = os.ErrorString(&#34;hexadecimal escape in host&#34;);
        <a id="L285"></a>goto Error;
    <a id="L286"></a>}

    <a id="L288"></a>return url, nil;

<a id="L290"></a>Error:
    <a id="L291"></a>return nil, &amp;URLError{&#34;parse&#34;, rawurl, err};

<a id="L293"></a>}

<a id="L295"></a><span class="comment">// ParseURLReference is like ParseURL but allows a trailing #fragment.</span>
<a id="L296"></a>func ParseURLReference(rawurlref string) (url *URL, err os.Error) {
    <a id="L297"></a><span class="comment">// Cut off #frag.</span>
    <a id="L298"></a>rawurl, frag := split(rawurlref, &#39;#&#39;, true);
    <a id="L299"></a>if url, err = ParseURL(rawurl); err != nil {
        <a id="L300"></a>return nil, err
    <a id="L301"></a>}
    <a id="L302"></a>if url.Fragment, err = URLUnescape(frag); err != nil {
        <a id="L303"></a>return nil, &amp;URLError{&#34;parse&#34;, rawurl, err}
    <a id="L304"></a>}
    <a id="L305"></a>return url, nil;
<a id="L306"></a>}

<a id="L308"></a><span class="comment">// String reassembles url into a valid URL string.</span>
<a id="L309"></a><span class="comment">//</span>
<a id="L310"></a><span class="comment">// There are redundant fields stored in the URL structure:</span>
<a id="L311"></a><span class="comment">// the String method consults Scheme, Path, Host, Userinfo,</span>
<a id="L312"></a><span class="comment">// RawQuery, and Fragment, but not Raw, RawPath or Authority.</span>
<a id="L313"></a>func (url *URL) String() string {
    <a id="L314"></a>result := &#34;&#34;;
    <a id="L315"></a>if url.Scheme != &#34;&#34; {
        <a id="L316"></a>result += url.Scheme + &#34;:&#34;
    <a id="L317"></a>}
    <a id="L318"></a>if url.Host != &#34;&#34; || url.Userinfo != &#34;&#34; {
        <a id="L319"></a>result += &#34;//&#34;;
        <a id="L320"></a>if url.Userinfo != &#34;&#34; {
            <a id="L321"></a>result += URLEscape(url.Userinfo) + &#34;@&#34;
        <a id="L322"></a>}
        <a id="L323"></a>result += url.Host;
    <a id="L324"></a>}
    <a id="L325"></a>result += URLEscape(url.Path);
    <a id="L326"></a>if url.RawQuery != &#34;&#34; {
        <a id="L327"></a>result += &#34;?&#34; + url.RawQuery
    <a id="L328"></a>}
    <a id="L329"></a>if url.Fragment != &#34;&#34; {
        <a id="L330"></a>result += &#34;#&#34; + URLEscape(url.Fragment)
    <a id="L331"></a>}
    <a id="L332"></a>return result;
<a id="L333"></a>}
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
