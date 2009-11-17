<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/http/url_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/http/url_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package http

<a id="L7"></a>import (
    <a id="L8"></a>&#34;fmt&#34;;
    <a id="L9"></a>&#34;os&#34;;
    <a id="L10"></a>&#34;reflect&#34;;
    <a id="L11"></a>&#34;testing&#34;;
<a id="L12"></a>)

<a id="L14"></a><span class="comment">// TODO(rsc):</span>
<a id="L15"></a><span class="comment">//	test URLUnescape</span>
<a id="L16"></a><span class="comment">//	test URLEscape</span>
<a id="L17"></a><span class="comment">//	test ParseURL</span>

<a id="L19"></a>type URLTest struct {
    <a id="L20"></a>in        string;
    <a id="L21"></a>out       *URL;
    <a id="L22"></a>roundtrip string; <span class="comment">// expected result of reserializing the URL; empty means same as &#34;in&#34;.</span>
<a id="L23"></a>}

<a id="L25"></a>var urltests = []URLTest{
    <a id="L26"></a><span class="comment">// no path</span>
    <a id="L27"></a>URLTest{
        <a id="L28"></a>&#34;http://www.google.com&#34;,
        <a id="L29"></a>&amp;URL{
            <a id="L30"></a>&#34;http://www.google.com&#34;,
            <a id="L31"></a>&#34;http&#34;, &#34;//www.google.com&#34;,
            <a id="L32"></a>&#34;www.google.com&#34;, &#34;&#34;, &#34;www.google.com&#34;,
            <a id="L33"></a>&#34;&#34;, &#34;&#34;, &#34;&#34;,
        <a id="L34"></a>},
        <a id="L35"></a>&#34;&#34;,
    <a id="L36"></a>},
    <a id="L37"></a><span class="comment">// path</span>
    <a id="L38"></a>URLTest{
        <a id="L39"></a>&#34;http://www.google.com/&#34;,
        <a id="L40"></a>&amp;URL{
            <a id="L41"></a>&#34;http://www.google.com/&#34;,
            <a id="L42"></a>&#34;http&#34;, &#34;//www.google.com/&#34;,
            <a id="L43"></a>&#34;www.google.com&#34;, &#34;&#34;, &#34;www.google.com&#34;,
            <a id="L44"></a>&#34;/&#34;, &#34;&#34;, &#34;&#34;,
        <a id="L45"></a>},
        <a id="L46"></a>&#34;&#34;,
    <a id="L47"></a>},
    <a id="L48"></a><span class="comment">// path with hex escaping... note that space roundtrips to +</span>
    <a id="L49"></a>URLTest{
        <a id="L50"></a>&#34;http://www.google.com/file%20one%26two&#34;,
        <a id="L51"></a>&amp;URL{
            <a id="L52"></a>&#34;http://www.google.com/file%20one%26two&#34;,
            <a id="L53"></a>&#34;http&#34;, &#34;//www.google.com/file%20one%26two&#34;,
            <a id="L54"></a>&#34;www.google.com&#34;, &#34;&#34;, &#34;www.google.com&#34;,
            <a id="L55"></a>&#34;/file one&amp;two&#34;, &#34;&#34;, &#34;&#34;,
        <a id="L56"></a>},
        <a id="L57"></a>&#34;http://www.google.com/file+one%26two&#34;,
    <a id="L58"></a>},
    <a id="L59"></a><span class="comment">// user</span>
    <a id="L60"></a>URLTest{
        <a id="L61"></a>&#34;ftp://webmaster@www.google.com/&#34;,
        <a id="L62"></a>&amp;URL{
            <a id="L63"></a>&#34;ftp://webmaster@www.google.com/&#34;,
            <a id="L64"></a>&#34;ftp&#34;, &#34;//webmaster@www.google.com/&#34;,
            <a id="L65"></a>&#34;webmaster@www.google.com&#34;, &#34;webmaster&#34;, &#34;www.google.com&#34;,
            <a id="L66"></a>&#34;/&#34;, &#34;&#34;, &#34;&#34;,
        <a id="L67"></a>},
        <a id="L68"></a>&#34;&#34;,
    <a id="L69"></a>},
    <a id="L70"></a><span class="comment">// escape sequence in username</span>
    <a id="L71"></a>URLTest{
        <a id="L72"></a>&#34;ftp://john%20doe@www.google.com/&#34;,
        <a id="L73"></a>&amp;URL{
            <a id="L74"></a>&#34;ftp://john%20doe@www.google.com/&#34;,
            <a id="L75"></a>&#34;ftp&#34;, &#34;//john%20doe@www.google.com/&#34;,
            <a id="L76"></a>&#34;john doe@www.google.com&#34;, &#34;john doe&#34;, &#34;www.google.com&#34;,
            <a id="L77"></a>&#34;/&#34;, &#34;&#34;, &#34;&#34;,
        <a id="L78"></a>},
        <a id="L79"></a>&#34;ftp://john+doe@www.google.com/&#34;,
    <a id="L80"></a>},
    <a id="L81"></a><span class="comment">// query</span>
    <a id="L82"></a>URLTest{
        <a id="L83"></a>&#34;http://www.google.com/?q=go+language&#34;,
        <a id="L84"></a>&amp;URL{
            <a id="L85"></a>&#34;http://www.google.com/?q=go+language&#34;,
            <a id="L86"></a>&#34;http&#34;, &#34;//www.google.com/?q=go+language&#34;,
            <a id="L87"></a>&#34;www.google.com&#34;, &#34;&#34;, &#34;www.google.com&#34;,
            <a id="L88"></a>&#34;/&#34;, &#34;q=go+language&#34;, &#34;&#34;,
        <a id="L89"></a>},
        <a id="L90"></a>&#34;&#34;,
    <a id="L91"></a>},
    <a id="L92"></a><span class="comment">// query with hex escaping: NOT parsed</span>
    <a id="L93"></a>URLTest{
        <a id="L94"></a>&#34;http://www.google.com/?q=go%20language&#34;,
        <a id="L95"></a>&amp;URL{
            <a id="L96"></a>&#34;http://www.google.com/?q=go%20language&#34;,
            <a id="L97"></a>&#34;http&#34;, &#34;//www.google.com/?q=go%20language&#34;,
            <a id="L98"></a>&#34;www.google.com&#34;, &#34;&#34;, &#34;www.google.com&#34;,
            <a id="L99"></a>&#34;/&#34;, &#34;q=go%20language&#34;, &#34;&#34;,
        <a id="L100"></a>},
        <a id="L101"></a>&#34;&#34;,
    <a id="L102"></a>},
    <a id="L103"></a><span class="comment">// path without /, so no query parsing</span>
    <a id="L104"></a>URLTest{
        <a id="L105"></a>&#34;http:www.google.com/?q=go+language&#34;,
        <a id="L106"></a>&amp;URL{
            <a id="L107"></a>&#34;http:www.google.com/?q=go+language&#34;,
            <a id="L108"></a>&#34;http&#34;, &#34;www.google.com/?q=go+language&#34;,
            <a id="L109"></a>&#34;&#34;, &#34;&#34;, &#34;&#34;,
            <a id="L110"></a>&#34;www.google.com/?q=go language&#34;, &#34;&#34;, &#34;&#34;,
        <a id="L111"></a>},
        <a id="L112"></a>&#34;http:www.google.com/%3fq%3dgo+language&#34;,
    <a id="L113"></a>},
    <a id="L114"></a><span class="comment">// non-authority</span>
    <a id="L115"></a>URLTest{
        <a id="L116"></a>&#34;mailto:/webmaster@golang.org&#34;,
        <a id="L117"></a>&amp;URL{
            <a id="L118"></a>&#34;mailto:/webmaster@golang.org&#34;,
            <a id="L119"></a>&#34;mailto&#34;, &#34;/webmaster@golang.org&#34;,
            <a id="L120"></a>&#34;&#34;, &#34;&#34;, &#34;&#34;,
            <a id="L121"></a>&#34;/webmaster@golang.org&#34;, &#34;&#34;, &#34;&#34;,
        <a id="L122"></a>},
        <a id="L123"></a>&#34;&#34;,
    <a id="L124"></a>},
    <a id="L125"></a><span class="comment">// non-authority</span>
    <a id="L126"></a>URLTest{
        <a id="L127"></a>&#34;mailto:webmaster@golang.org&#34;,
        <a id="L128"></a>&amp;URL{
            <a id="L129"></a>&#34;mailto:webmaster@golang.org&#34;,
            <a id="L130"></a>&#34;mailto&#34;, &#34;webmaster@golang.org&#34;,
            <a id="L131"></a>&#34;&#34;, &#34;&#34;, &#34;&#34;,
            <a id="L132"></a>&#34;webmaster@golang.org&#34;, &#34;&#34;, &#34;&#34;,
        <a id="L133"></a>},
        <a id="L134"></a>&#34;&#34;,
    <a id="L135"></a>},
    <a id="L136"></a><span class="comment">// unescaped :// in query should not create a scheme</span>
    <a id="L137"></a>URLTest{
        <a id="L138"></a>&#34;/foo?query=http://bad&#34;,
        <a id="L139"></a>&amp;URL{
            <a id="L140"></a>&#34;/foo?query=http://bad&#34;,
            <a id="L141"></a>&#34;&#34;, &#34;/foo?query=http://bad&#34;,
            <a id="L142"></a>&#34;&#34;, &#34;&#34;, &#34;&#34;,
            <a id="L143"></a>&#34;/foo&#34;, &#34;query=http://bad&#34;, &#34;&#34;,
        <a id="L144"></a>},
        <a id="L145"></a>&#34;&#34;,
    <a id="L146"></a>},
<a id="L147"></a>}

<a id="L149"></a>var urlnofragtests = []URLTest{
    <a id="L150"></a>URLTest{
        <a id="L151"></a>&#34;http://www.google.com/?q=go+language#foo&#34;,
        <a id="L152"></a>&amp;URL{
            <a id="L153"></a>&#34;http://www.google.com/?q=go+language#foo&#34;,
            <a id="L154"></a>&#34;http&#34;, &#34;//www.google.com/?q=go+language#foo&#34;,
            <a id="L155"></a>&#34;www.google.com&#34;, &#34;&#34;, &#34;www.google.com&#34;,
            <a id="L156"></a>&#34;/&#34;, &#34;q=go+language#foo&#34;, &#34;&#34;,
        <a id="L157"></a>},
        <a id="L158"></a>&#34;&#34;,
    <a id="L159"></a>},
<a id="L160"></a>}

<a id="L162"></a>var urlfragtests = []URLTest{
    <a id="L163"></a>URLTest{
        <a id="L164"></a>&#34;http://www.google.com/?q=go+language#foo&#34;,
        <a id="L165"></a>&amp;URL{
            <a id="L166"></a>&#34;http://www.google.com/?q=go+language&#34;,
            <a id="L167"></a>&#34;http&#34;, &#34;//www.google.com/?q=go+language&#34;,
            <a id="L168"></a>&#34;www.google.com&#34;, &#34;&#34;, &#34;www.google.com&#34;,
            <a id="L169"></a>&#34;/&#34;, &#34;q=go+language&#34;, &#34;foo&#34;,
        <a id="L170"></a>},
        <a id="L171"></a>&#34;&#34;,
    <a id="L172"></a>},
    <a id="L173"></a>URLTest{
        <a id="L174"></a>&#34;http://www.google.com/?q=go+language#foo%26bar&#34;,
        <a id="L175"></a>&amp;URL{
            <a id="L176"></a>&#34;http://www.google.com/?q=go+language&#34;,
            <a id="L177"></a>&#34;http&#34;, &#34;//www.google.com/?q=go+language&#34;,
            <a id="L178"></a>&#34;www.google.com&#34;, &#34;&#34;, &#34;www.google.com&#34;,
            <a id="L179"></a>&#34;/&#34;, &#34;q=go+language&#34;, &#34;foo&amp;bar&#34;,
        <a id="L180"></a>},
        <a id="L181"></a>&#34;&#34;,
    <a id="L182"></a>},
<a id="L183"></a>}

<a id="L185"></a><span class="comment">// more useful string for debugging than fmt&#39;s struct printer</span>
<a id="L186"></a>func ufmt(u *URL) string {
    <a id="L187"></a>return fmt.Sprintf(&#34;%q, %q, %q, %q, %q, %q, %q, %q, %q&#34;,
        <a id="L188"></a>u.Raw, u.Scheme, u.RawPath, u.Authority, u.Userinfo,
        <a id="L189"></a>u.Host, u.Path, u.RawQuery, u.Fragment)
<a id="L190"></a>}

<a id="L192"></a>func DoTest(t *testing.T, parse func(string) (*URL, os.Error), name string, tests []URLTest) {
    <a id="L193"></a>for _, tt := range tests {
        <a id="L194"></a>u, err := parse(tt.in);
        <a id="L195"></a>if err != nil {
            <a id="L196"></a>t.Errorf(&#34;%s(%q) returned error %s&#34;, name, tt.in, err);
            <a id="L197"></a>continue;
        <a id="L198"></a>}
        <a id="L199"></a>if !reflect.DeepEqual(u, tt.out) {
            <a id="L200"></a>t.Errorf(&#34;%s(%q):\n\thave %v\n\twant %v\n&#34;,
                <a id="L201"></a>name, tt.in, ufmt(u), ufmt(tt.out))
        <a id="L202"></a>}
    <a id="L203"></a>}
<a id="L204"></a>}

<a id="L206"></a>func TestParseURL(t *testing.T) {
    <a id="L207"></a>DoTest(t, ParseURL, &#34;ParseURL&#34;, urltests);
    <a id="L208"></a>DoTest(t, ParseURL, &#34;ParseURL&#34;, urlnofragtests);
<a id="L209"></a>}

<a id="L211"></a>func TestParseURLReference(t *testing.T) {
    <a id="L212"></a>DoTest(t, ParseURLReference, &#34;ParseURLReference&#34;, urltests);
    <a id="L213"></a>DoTest(t, ParseURLReference, &#34;ParseURLReference&#34;, urlfragtests);
<a id="L214"></a>}

<a id="L216"></a>func DoTestString(t *testing.T, parse func(string) (*URL, os.Error), name string, tests []URLTest) {
    <a id="L217"></a>for _, tt := range tests {
        <a id="L218"></a>u, err := parse(tt.in);
        <a id="L219"></a>if err != nil {
            <a id="L220"></a>t.Errorf(&#34;%s(%q) returned error %s&#34;, name, tt.in, err);
            <a id="L221"></a>continue;
        <a id="L222"></a>}
        <a id="L223"></a>s := u.String();
        <a id="L224"></a>expected := tt.in;
        <a id="L225"></a>if len(tt.roundtrip) &gt; 0 {
            <a id="L226"></a>expected = tt.roundtrip
        <a id="L227"></a>}
        <a id="L228"></a>if s != expected {
            <a id="L229"></a>t.Errorf(&#34;%s(%q).String() == %q (expected %q)&#34;, name, tt.in, s, expected)
        <a id="L230"></a>}
    <a id="L231"></a>}
<a id="L232"></a>}

<a id="L234"></a>func TestURLString(t *testing.T) {
    <a id="L235"></a>DoTestString(t, ParseURL, &#34;ParseURL&#34;, urltests);
    <a id="L236"></a>DoTestString(t, ParseURL, &#34;ParseURL&#34;, urlfragtests);
    <a id="L237"></a>DoTestString(t, ParseURL, &#34;ParseURL&#34;, urlnofragtests);
    <a id="L238"></a>DoTestString(t, ParseURLReference, &#34;ParseURLReference&#34;, urltests);
    <a id="L239"></a>DoTestString(t, ParseURLReference, &#34;ParseURLReference&#34;, urlfragtests);
    <a id="L240"></a>DoTestString(t, ParseURLReference, &#34;ParseURLReference&#34;, urlnofragtests);
<a id="L241"></a>}

<a id="L243"></a>type URLEscapeTest struct {
    <a id="L244"></a>in  string;
    <a id="L245"></a>out string;
    <a id="L246"></a>err os.Error;
<a id="L247"></a>}

<a id="L249"></a>var unescapeTests = []URLEscapeTest{
    <a id="L250"></a>URLEscapeTest{
        <a id="L251"></a>&#34;&#34;,
        <a id="L252"></a>&#34;&#34;,
        <a id="L253"></a>nil,
    <a id="L254"></a>},
    <a id="L255"></a>URLEscapeTest{
        <a id="L256"></a>&#34;abc&#34;,
        <a id="L257"></a>&#34;abc&#34;,
        <a id="L258"></a>nil,
    <a id="L259"></a>},
    <a id="L260"></a>URLEscapeTest{
        <a id="L261"></a>&#34;1%41&#34;,
        <a id="L262"></a>&#34;1A&#34;,
        <a id="L263"></a>nil,
    <a id="L264"></a>},
    <a id="L265"></a>URLEscapeTest{
        <a id="L266"></a>&#34;1%41%42%43&#34;,
        <a id="L267"></a>&#34;1ABC&#34;,
        <a id="L268"></a>nil,
    <a id="L269"></a>},
    <a id="L270"></a>URLEscapeTest{
        <a id="L271"></a>&#34;%4a&#34;,
        <a id="L272"></a>&#34;J&#34;,
        <a id="L273"></a>nil,
    <a id="L274"></a>},
    <a id="L275"></a>URLEscapeTest{
        <a id="L276"></a>&#34;%6F&#34;,
        <a id="L277"></a>&#34;o&#34;,
        <a id="L278"></a>nil,
    <a id="L279"></a>},
    <a id="L280"></a>URLEscapeTest{
        <a id="L281"></a>&#34;%&#34;, <span class="comment">// not enough characters after %</span>
        <a id="L282"></a>&#34;&#34;,
        <a id="L283"></a>URLEscapeError(&#34;%&#34;),
    <a id="L284"></a>},
    <a id="L285"></a>URLEscapeTest{
        <a id="L286"></a>&#34;%a&#34;, <span class="comment">// not enough characters after %</span>
        <a id="L287"></a>&#34;&#34;,
        <a id="L288"></a>URLEscapeError(&#34;%a&#34;),
    <a id="L289"></a>},
    <a id="L290"></a>URLEscapeTest{
        <a id="L291"></a>&#34;%1&#34;, <span class="comment">// not enough characters after %</span>
        <a id="L292"></a>&#34;&#34;,
        <a id="L293"></a>URLEscapeError(&#34;%1&#34;),
    <a id="L294"></a>},
    <a id="L295"></a>URLEscapeTest{
        <a id="L296"></a>&#34;123%45%6&#34;, <span class="comment">// not enough characters after %</span>
        <a id="L297"></a>&#34;&#34;,
        <a id="L298"></a>URLEscapeError(&#34;%6&#34;),
    <a id="L299"></a>},
    <a id="L300"></a>URLEscapeTest{
        <a id="L301"></a>&#34;%zzzzz&#34;, <span class="comment">// invalid hex digits</span>
        <a id="L302"></a>&#34;&#34;,
        <a id="L303"></a>URLEscapeError(&#34;%zz&#34;),
    <a id="L304"></a>},
<a id="L305"></a>}

<a id="L307"></a>func TestURLUnescape(t *testing.T) {
    <a id="L308"></a>for _, tt := range unescapeTests {
        <a id="L309"></a>actual, err := URLUnescape(tt.in);
        <a id="L310"></a>if actual != tt.out || (err != nil) != (tt.err != nil) {
            <a id="L311"></a>t.Errorf(&#34;URLUnescape(%q) = %q, %s; want %q, %s&#34;, tt.in, actual, err, tt.out, tt.err)
        <a id="L312"></a>}
    <a id="L313"></a>}
<a id="L314"></a>}

<a id="L316"></a>var escapeTests = []URLEscapeTest{
    <a id="L317"></a>URLEscapeTest{
        <a id="L318"></a>&#34;&#34;,
        <a id="L319"></a>&#34;&#34;,
        <a id="L320"></a>nil,
    <a id="L321"></a>},
    <a id="L322"></a>URLEscapeTest{
        <a id="L323"></a>&#34;abc&#34;,
        <a id="L324"></a>&#34;abc&#34;,
        <a id="L325"></a>nil,
    <a id="L326"></a>},
    <a id="L327"></a>URLEscapeTest{
        <a id="L328"></a>&#34;one two&#34;,
        <a id="L329"></a>&#34;one+two&#34;,
        <a id="L330"></a>nil,
    <a id="L331"></a>},
    <a id="L332"></a>URLEscapeTest{
        <a id="L333"></a>&#34;10%&#34;,
        <a id="L334"></a>&#34;10%25&#34;,
        <a id="L335"></a>nil,
    <a id="L336"></a>},
    <a id="L337"></a>URLEscapeTest{
        <a id="L338"></a>&#34; ?&amp;=#+%!&#34;,
        <a id="L339"></a>&#34;+%3f%26%3d%23%2b%25!&#34;,
        <a id="L340"></a>nil,
    <a id="L341"></a>},
<a id="L342"></a>}

<a id="L344"></a>func TestURLEscape(t *testing.T) {
    <a id="L345"></a>for _, tt := range escapeTests {
        <a id="L346"></a>actual := URLEscape(tt.in);
        <a id="L347"></a>if tt.out != actual {
            <a id="L348"></a>t.Errorf(&#34;URLEscape(%q) = %q, want %q&#34;, tt.in, actual, tt.out)
        <a id="L349"></a>}

        <a id="L351"></a><span class="comment">// for bonus points, verify that escape:unescape is an identity.</span>
        <a id="L352"></a>roundtrip, err := URLUnescape(actual);
        <a id="L353"></a>if roundtrip != tt.in || err != nil {
            <a id="L354"></a>t.Errorf(&#34;URLUnescape(%q) = %q, %s; want %q, %s&#34;, actual, roundtrip, err, tt.in, &#34;[no error]&#34;)
        <a id="L355"></a>}
    <a id="L356"></a>}
<a id="L357"></a>}
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
